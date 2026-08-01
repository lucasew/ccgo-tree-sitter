// Code generated for linux/386 by 'ccgo -ignore-unsupported-alignment -ignore-unsupported-atomic-sizes -ignore-vector-functions -std=gnu11 -O0 -Dfunc=func_token -Dinterface=interface_token -Dselect=select_token -Dchan=chan_token -Dgo=go_token -Dmap=map_token -Dpackage=package_token -Dtype=type_token -Dvar=var_token -Dimport=import_token -Ddefer=defer_token -Dfallthrough=fallthrough_token -Drange=range_token -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-toml/src -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-toml -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party -I /home/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/include -I /home/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/src combined.c -o grammar.go', DO NOT EDIT.

//go:build linux && 386

package grammar_toml

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
const EXTERNAL_TOKEN_COUNT = 5
const FIELD_COUNT = 0
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
const LANGUAGE_VERSION = 13
const LARGE_STATE_COUNT = 2
const MAX_ALIAS_SEQUENCE_LENGTH = 8
const PRODUCTION_ID_COUNT = 2
const PTRDIFF_MAX = "INT32_MAX"
const PTRDIFF_MIN = "INT32_MIN"
const RAND_MAX = 0x7fffffff
const SIG_ATOMIC_MAX = "INT32_MAX"
const SIG_ATOMIC_MIN = "INT32_MIN"
const SIZE_MAX = "UINT32_MAX"
const STATE_COUNT = 152
const SYMBOL_COUNT = 66
const TOKEN_COUNT = 40
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

/*
 *  Lexer Macros
 */

/*
 *  Parse Table Macros
 */

type TokenType = int32

const LINE_ENDING_OR_EOF = 0
const MULTILINE_BASIC_STRING_CONTENT = 1
const MULTILINE_BASIC_STRING_END = 2
const MULTILINE_LITERAL_STRING_CONTENT = 3
const MULTILINE_LITERAL_STRING_END = 4

func tree_sitter_toml_external_scanner_create(tls *libc.TLS) (r uintptr) {
	return libc.UintptrFromInt32(0)
}

func tree_sitter_toml_external_scanner_destroy(tls *libc.TLS, payload uintptr) {
}

func tree_sitter_toml_external_scanner_serialize(tls *libc.TLS, payload uintptr, buffer uintptr) (r uint32) {
	return uint32(0)
}

func tree_sitter_toml_external_scanner_deserialize(tls *libc.TLS, payload uintptr, buffer uintptr, length uint32) {
}

func tree_sitter_toml_external_scanner_scan_multiline_string_end(tls *libc.TLS, lexer uintptr, valid_symbols uintptr, delimiter int32_t, content_symbol TokenType, end_symbol TokenType) (r uint8) {
	if !(*(*uint8)(unsafe.Pointer(valid_symbols + uintptr(end_symbol))) != 0) || (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != delimiter {
		return libc.BoolUint8(false1 != 0)
	}
	(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(false1 != 0))
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
	if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != delimiter {
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = libc.Uint16FromInt32(content_symbol)
		return libc.BoolUint8(true1 != 0)
	}
	(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(false1 != 0))
	if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != delimiter {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = libc.Uint16FromInt32(content_symbol)
		return libc.BoolUint8(true1 != 0)
	}
	(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(false1 != 0))
	if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != delimiter {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = libc.Uint16FromInt32(end_symbol)
		return libc.BoolUint8(true1 != 0)
	}
	(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = libc.Uint16FromInt32(content_symbol)
	return libc.BoolUint8(true1 != 0)
}

func tree_sitter_toml_external_scanner_scan(tls *libc.TLS, payload uintptr, lexer uintptr, valid_symbols uintptr) (r uint8) {
	if tree_sitter_toml_external_scanner_scan_multiline_string_end(tls, lexer, valid_symbols, int32('"'), int32(MULTILINE_BASIC_STRING_CONTENT), int32(MULTILINE_BASIC_STRING_END)) != 0 || tree_sitter_toml_external_scanner_scan_multiline_string_end(tls, lexer, valid_symbols, int32('\''), int32(MULTILINE_LITERAL_STRING_CONTENT), int32(MULTILINE_LITERAL_STRING_END)) != 0 {
		return libc.BoolUint8(true1 != 0)
	}
	if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(LINE_ENDING_OR_EOF))) != 0 {
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(LINE_ENDING_OR_EOF)
		for (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32(' ') || (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('\t') {
			(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(true1 != 0))
		}
		if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == 0 || (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('\n') {
			return libc.BoolUint8(true1 != 0)
		}
		if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('\r') {
			(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(true1 != 0))
			if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('\n') {
				return libc.BoolUint8(true1 != 0)
			}
		}
	}
	return libc.BoolUint8(false1 != 0)
}

const aux_sym_document_token1 = 1
const sym_comment = 2
const anon_sym_LBRACK = 3
const anon_sym_RBRACK = 4
const anon_sym_LBRACK_LBRACK = 5
const anon_sym_RBRACK_RBRACK = 6
const anon_sym_EQ = 7
const anon_sym_DOT = 8
const sym_bare_key = 9
const anon_sym_DQUOTE = 10
const aux_sym__basic_string_token1 = 11
const anon_sym_DQUOTE2 = 12
const anon_sym_DQUOTE_DQUOTE_DQUOTE = 13
const aux_sym__multiline_basic_string_token1 = 14
const sym_escape_sequence = 15
const sym__escape_line_ending = 16
const anon_sym_SQUOTE = 17
const aux_sym__literal_string_token1 = 18
const anon_sym_SQUOTE2 = 19
const anon_sym_SQUOTE_SQUOTE_SQUOTE = 20
const aux_sym_integer_token1 = 21
const aux_sym_integer_token2 = 22
const aux_sym_integer_token3 = 23
const aux_sym_integer_token4 = 24
const aux_sym_float_token1 = 25
const aux_sym_float_token2 = 26
const sym_boolean = 27
const sym_offset_date_time = 28
const sym_local_date_time = 29
const sym_local_date = 30
const sym_local_time = 31
const anon_sym_COMMA = 32
const anon_sym_LBRACE = 33
const anon_sym_RBRACE = 34
const sym__line_ending_or_eof = 35
const sym__multiline_basic_string_content = 36
const sym__multiline_basic_string_end = 37
const sym__multiline_literal_string_content = 38
const sym__multiline_literal_string_end = 39
const sym_document = 40
const sym_table = 41
const sym_table_array_element = 42
const sym_pair = 43
const sym__inline_pair = 44
const sym__key = 45
const sym_dotted_key = 46
const sym_quoted_key = 47
const sym__inline_value = 48
const sym_string = 49
const sym__basic_string = 50
const sym__multiline_basic_string = 51
const sym__literal_string = 52
const sym__multiline_literal_string = 53
const sym_integer = 54
const sym_float = 55
const sym_array = 56
const sym_inline_table = 57
const aux_sym_document_repeat1 = 58
const aux_sym_document_repeat2 = 59
const aux_sym__basic_string_repeat1 = 60
const aux_sym__multiline_basic_string_repeat1 = 61
const aux_sym__multiline_literal_string_repeat1 = 62
const aux_sym_array_repeat1 = 63
const aux_sym_array_repeat2 = 64
const aux_sym_inline_table_repeat1 = 65

var ts_symbol_names = [66]uintptr{
	0:  __ccgo_ts,
	1:  __ccgo_ts + 4,
	2:  __ccgo_ts + 20,
	3:  __ccgo_ts + 28,
	4:  __ccgo_ts + 30,
	5:  __ccgo_ts + 32,
	6:  __ccgo_ts + 35,
	7:  __ccgo_ts + 38,
	8:  __ccgo_ts + 40,
	9:  __ccgo_ts + 42,
	10: __ccgo_ts + 51,
	11: __ccgo_ts + 53,
	12: __ccgo_ts + 51,
	13: __ccgo_ts + 74,
	14: __ccgo_ts + 78,
	15: __ccgo_ts + 109,
	16: __ccgo_ts + 109,
	17: __ccgo_ts + 125,
	18: __ccgo_ts + 127,
	19: __ccgo_ts + 125,
	20: __ccgo_ts + 150,
	21: __ccgo_ts + 154,
	22: __ccgo_ts + 169,
	23: __ccgo_ts + 184,
	24: __ccgo_ts + 199,
	25: __ccgo_ts + 214,
	26: __ccgo_ts + 227,
	27: __ccgo_ts + 240,
	28: __ccgo_ts + 248,
	29: __ccgo_ts + 265,
	30: __ccgo_ts + 281,
	31: __ccgo_ts + 292,
	32: __ccgo_ts + 303,
	33: __ccgo_ts + 305,
	34: __ccgo_ts + 307,
	35: __ccgo_ts + 309,
	36: __ccgo_ts + 329,
	37: __ccgo_ts + 361,
	38: __ccgo_ts + 389,
	39: __ccgo_ts + 423,
	40: __ccgo_ts + 453,
	41: __ccgo_ts + 462,
	42: __ccgo_ts + 468,
	43: __ccgo_ts + 488,
	44: __ccgo_ts + 493,
	45: __ccgo_ts + 506,
	46: __ccgo_ts + 511,
	47: __ccgo_ts + 522,
	48: __ccgo_ts + 533,
	49: __ccgo_ts + 547,
	50: __ccgo_ts + 554,
	51: __ccgo_ts + 568,
	52: __ccgo_ts + 592,
	53: __ccgo_ts + 608,
	54: __ccgo_ts + 634,
	55: __ccgo_ts + 642,
	56: __ccgo_ts + 648,
	57: __ccgo_ts + 654,
	58: __ccgo_ts + 667,
	59: __ccgo_ts + 684,
	60: __ccgo_ts + 701,
	61: __ccgo_ts + 723,
	62: __ccgo_ts + 755,
	63: __ccgo_ts + 789,
	64: __ccgo_ts + 803,
	65: __ccgo_ts + 817,
}

var ts_symbol_map = [66]TSSymbol{
	1:  uint16(aux_sym_document_token1),
	2:  uint16(sym_comment),
	3:  uint16(anon_sym_LBRACK),
	4:  uint16(anon_sym_RBRACK),
	5:  uint16(anon_sym_LBRACK_LBRACK),
	6:  uint16(anon_sym_RBRACK_RBRACK),
	7:  uint16(anon_sym_EQ),
	8:  uint16(anon_sym_DOT),
	9:  uint16(sym_bare_key),
	10: uint16(anon_sym_DQUOTE),
	11: uint16(aux_sym__basic_string_token1),
	12: uint16(anon_sym_DQUOTE),
	13: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	14: uint16(aux_sym__multiline_basic_string_token1),
	15: uint16(sym_escape_sequence),
	16: uint16(sym_escape_sequence),
	17: uint16(anon_sym_SQUOTE),
	18: uint16(aux_sym__literal_string_token1),
	19: uint16(anon_sym_SQUOTE),
	20: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	21: uint16(aux_sym_integer_token1),
	22: uint16(aux_sym_integer_token2),
	23: uint16(aux_sym_integer_token3),
	24: uint16(aux_sym_integer_token4),
	25: uint16(aux_sym_float_token1),
	26: uint16(aux_sym_float_token2),
	27: uint16(sym_boolean),
	28: uint16(sym_offset_date_time),
	29: uint16(sym_local_date_time),
	30: uint16(sym_local_date),
	31: uint16(sym_local_time),
	32: uint16(anon_sym_COMMA),
	33: uint16(anon_sym_LBRACE),
	34: uint16(anon_sym_RBRACE),
	35: uint16(sym__line_ending_or_eof),
	36: uint16(sym__multiline_basic_string_content),
	37: uint16(sym__multiline_basic_string_end),
	38: uint16(sym__multiline_literal_string_content),
	39: uint16(sym__multiline_literal_string_end),
	40: uint16(sym_document),
	41: uint16(sym_table),
	42: uint16(sym_table_array_element),
	43: uint16(sym_pair),
	44: uint16(sym__inline_pair),
	45: uint16(sym__key),
	46: uint16(sym_dotted_key),
	47: uint16(sym_quoted_key),
	48: uint16(sym__inline_value),
	49: uint16(sym_string),
	50: uint16(sym__basic_string),
	51: uint16(sym__multiline_basic_string),
	52: uint16(sym__literal_string),
	53: uint16(sym__multiline_literal_string),
	54: uint16(sym_integer),
	55: uint16(sym_float),
	56: uint16(sym_array),
	57: uint16(sym_inline_table),
	58: uint16(aux_sym_document_repeat1),
	59: uint16(aux_sym_document_repeat2),
	60: uint16(aux_sym__basic_string_repeat1),
	61: uint16(aux_sym__multiline_basic_string_repeat1),
	62: uint16(aux_sym__multiline_literal_string_repeat1),
	63: uint16(aux_sym_array_repeat1),
	64: uint16(aux_sym_array_repeat2),
	65: uint16(aux_sym_inline_table_repeat1),
}

var ts_symbol_metadata = [66]TSSymbolMetadata{
	0: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	1: {},
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
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	10: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	11: {},
	12: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	13: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	14: {},
	15: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	16: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	17: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	18: {},
	19: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	20: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	21: {},
	22: {},
	23: {},
	24: {},
	25: {},
	26: {},
	27: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
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
	},
	33: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	34: {
		Fvisible: libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	39: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	40: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	41: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	45: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	46: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	47: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	48: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	49: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	50: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	51: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	52: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	53: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	54: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	55: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	56: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	57: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	58: {},
	59: {},
	60: {},
	61: {},
	62: {},
	63: {},
	64: {},
	65: {},
}

var ts_alias_sequences = [2][8]TSSymbol{
	0: {},
	1: {
		1: uint16(sym_pair),
	},
}

var ts_non_terminal_alias_map = [5]uint16_t{
	0: uint16(sym__inline_pair),
	1: uint16(2),
	2: uint16(sym__inline_pair),
	3: uint16(sym_pair),
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
			state = uint16(77)
			goto next_state
		}
		if lookahead == int32('\n') {
			state = uint16(129)
			goto next_state
		}
		if lookahead == int32('\r') {
			state = uint16(1)
			goto next_state
		}
		if lookahead == int32('"') {
			state = uint16(127)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(79)
			goto next_state
		}
		if lookahead == int32('\'') {
			state = uint16(137)
			goto next_state
		}
		if lookahead == int32('+') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(161)
			goto next_state
		}
		if lookahead == int32('-') {
			state = uint16(95)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(86)
			goto next_state
		}
		if lookahead == int32('0') {
			state = uint16(92)
			goto next_state
		}
		if lookahead == int32('1') {
			state = uint16(90)
			goto next_state
		}
		if lookahead == int32('2') {
			state = uint16(89)
			goto next_state
		}
		if lookahead == int32('=') {
			state = uint16(85)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(81)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(5)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(82)
			goto next_state
		}
		if lookahead == int32('f') {
			state = uint16(102)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(108)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(103)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(109)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(162)
			goto next_state
		}
		if lookahead == int32('}') {
			state = uint16(163)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(75)
			goto next_state
		}
		if int32('3') <= lookahead && lookahead <= int32('9') {
			state = uint16(91)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(1):
		if lookahead == int32('\n') {
			state = uint16(129)
			goto next_state
		}
		return result
	case int32(2):
		if lookahead == int32('\n') {
			state = uint16(129)
			goto next_state
		}
		if lookahead == int32('\r') {
			state = uint16(1)
			goto next_state
		}
		if lookahead == int32('"') {
			state = uint16(126)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(125)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(5)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			state = uint16(124)
			goto next_state
		}
		if lookahead != 0 && lookahead > int32(31) && lookahead != int32(127) {
			state = uint16(125)
			goto next_state
		}
		return result
	case int32(3):
		if lookahead == int32('\n') {
			state = uint16(129)
			goto next_state
		}
		if lookahead == int32('\r') {
			state = uint16(1)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(135)
			goto next_state
		}
		if lookahead == int32('\'') {
			state = uint16(136)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			state = uint16(134)
			goto next_state
		}
		if lookahead != 0 && lookahead > int32(31) && lookahead != int32(127) {
			state = uint16(135)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead == int32('\n') {
			state = uint16(131)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead == int32('\n') {
			state = uint16(131)
			goto next_state
		}
		if lookahead == int32('\r') {
			state = uint16(4)
			goto next_state
		}
		if lookahead == int32('U') {
			state = uint16(74)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(70)
			goto next_state
		}
		if lookahead == int32('"') || lookahead == int32('\\') || lookahead == int32('b') || lookahead == int32('f') || lookahead == int32('n') || lookahead == int32('r') || lookahead == int32('t') {
			state = uint16(130)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead == int32('\n') {
			state = uint16(78)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead == int32('\n') {
			state = uint16(78)
			goto next_state
		}
		if lookahead == int32('\r') {
			state = uint16(6)
			goto next_state
		}
		if lookahead == int32('"') {
			state = uint16(123)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(79)
			goto next_state
		}
		if lookahead == int32('\'') {
			state = uint16(133)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(161)
			goto next_state
		}
		if lookahead == int32('0') {
			state = uint16(144)
			goto next_state
		}
		if lookahead == int32('1') {
			state = uint16(142)
			goto next_state
		}
		if lookahead == int32('2') {
			state = uint16(141)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(80)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(82)
			goto next_state
		}
		if lookahead == int32('f') {
			state = uint16(30)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(35)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(31)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(37)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(162)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(7)
			goto next_state
		}
		if int32('+') <= lookahead && lookahead <= int32('-') {
			state = uint16(17)
			goto next_state
		}
		if int32('3') <= lookahead && lookahead <= int32('9') {
			state = uint16(143)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead == int32('"') {
			state = uint16(128)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead == int32('#') {
			state = uint16(79)
			goto next_state
		}
		if lookahead == int32('\'') {
			state = uint16(136)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(86)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(29)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(10)
			goto next_state
		}
		return result
	case int32(10):
		if lookahead == int32('#') {
			state = uint16(79)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(86)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(29)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(10)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead == int32('\'') {
			state = uint16(138)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead == int32('-') {
			state = uint16(18)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(46)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(13)
			goto next_state
		}
		return result
	case int32(13):
		if lookahead == int32('-') {
			state = uint16(18)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(13)
			goto next_state
		}
		return result
	case int32(14):
		if lookahead == int32('-') {
			state = uint16(20)
			goto next_state
		}
		return result
	case int32(15):
		if lookahead == int32('0') {
			state = uint16(139)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(35)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(31)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(147)
			goto next_state
		}
		return result
	case int32(16):
		if lookahead == int32('0') {
			state = uint16(159)
			goto next_state
		}
		return result
	case int32(17):
		if lookahead == int32('0') {
			state = uint16(146)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(35)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(31)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(145)
			goto next_state
		}
		return result
	case int32(18):
		if lookahead == int32('0') {
			state = uint16(50)
			goto next_state
		}
		if lookahead == int32('1') {
			state = uint16(43)
			goto next_state
		}
		return result
	case int32(19):
		if lookahead == int32('0') {
			state = uint16(156)
			goto next_state
		}
		return result
	case int32(20):
		if lookahead == int32('0') {
			state = uint16(51)
			goto next_state
		}
		if lookahead == int32('3') {
			state = uint16(42)
			goto next_state
		}
		if lookahead == int32('1') || lookahead == int32('2') {
			state = uint16(59)
			goto next_state
		}
		return result
	case int32(21):
		if lookahead == int32('2') {
			state = uint16(44)
			goto next_state
		}
		if lookahead == int32('0') || lookahead == int32('1') {
			state = uint16(63)
			goto next_state
		}
		return result
	case int32(22):
		if lookahead == int32('2') {
			state = uint16(45)
			goto next_state
		}
		if lookahead == int32('0') || lookahead == int32('1') {
			state = uint16(64)
			goto next_state
		}
		return result
	case int32(23):
		if lookahead == int32('6') {
			state = uint16(16)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('5') {
			state = uint16(54)
			goto next_state
		}
		return result
	case int32(24):
		if lookahead == int32('6') {
			state = uint16(19)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('5') {
			state = uint16(60)
			goto next_state
		}
		return result
	case int32(25):
		if lookahead == int32(':') {
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(26):
		if lookahead == int32(':') {
			state = uint16(47)
			goto next_state
		}
		return result
	case int32(27):
		if lookahead == int32(':') {
			state = uint16(24)
			goto next_state
		}
		return result
	case int32(28):
		if lookahead == int32(':') {
			state = uint16(48)
			goto next_state
		}
		return result
	case int32(29):
		if lookahead == int32(']') {
			state = uint16(84)
			goto next_state
		}
		return result
	case int32(30):
		if lookahead == int32('a') {
			state = uint16(34)
			goto next_state
		}
		return result
	case int32(31):
		if lookahead == int32('a') {
			state = uint16(36)
			goto next_state
		}
		return result
	case int32(32):
		if lookahead == int32('e') {
			state = uint16(154)
			goto next_state
		}
		return result
	case int32(33):
		if lookahead == int32('f') {
			state = uint16(153)
			goto next_state
		}
		return result
	case int32(34):
		if lookahead == int32('l') {
			state = uint16(38)
			goto next_state
		}
		return result
	case int32(35):
		if lookahead == int32('n') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(36):
		if lookahead == int32('n') {
			state = uint16(153)
			goto next_state
		}
		return result
	case int32(37):
		if lookahead == int32('r') {
			state = uint16(39)
			goto next_state
		}
		return result
	case int32(38):
		if lookahead == int32('s') {
			state = uint16(32)
			goto next_state
		}
		return result
	case int32(39):
		if lookahead == int32('u') {
			state = uint16(32)
			goto next_state
		}
		return result
	case int32(40):
		if lookahead == int32('+') || lookahead == int32('-') {
			state = uint16(58)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(152)
			goto next_state
		}
		return result
	case int32(41):
		if lookahead == int32('0') || lookahead == int32('1') {
			state = uint16(150)
			goto next_state
		}
		return result
	case int32(42):
		if lookahead == int32('0') || lookahead == int32('1') {
			state = uint16(158)
			goto next_state
		}
		return result
	case int32(43):
		if int32('0') <= lookahead && lookahead <= int32('2') {
			state = uint16(14)
			goto next_state
		}
		return result
	case int32(44):
		if int32('0') <= lookahead && lookahead <= int32('3') {
			state = uint16(28)
			goto next_state
		}
		return result
	case int32(45):
		if int32('0') <= lookahead && lookahead <= int32('3') {
			state = uint16(26)
			goto next_state
		}
		return result
	case int32(46):
		if int32('0') <= lookahead && lookahead <= int32('5') {
			state = uint16(53)
			goto next_state
		}
		return result
	case int32(47):
		if int32('0') <= lookahead && lookahead <= int32('5') {
			state = uint16(61)
			goto next_state
		}
		return result
	case int32(48):
		if int32('0') <= lookahead && lookahead <= int32('5') {
			state = uint16(65)
			goto next_state
		}
		return result
	case int32(49):
		if int32('0') <= lookahead && lookahead <= int32('7') {
			state = uint16(149)
			goto next_state
		}
		return result
	case int32(50):
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(14)
			goto next_state
		}
		return result
	case int32(51):
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(158)
			goto next_state
		}
		return result
	case int32(52):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(147)
			goto next_state
		}
		return result
	case int32(53):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(25)
			goto next_state
		}
		return result
	case int32(54):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(159)
			goto next_state
		}
		return result
	case int32(55):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(160)
			goto next_state
		}
		return result
	case int32(56):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(145)
			goto next_state
		}
		return result
	case int32(57):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(151)
			goto next_state
		}
		return result
	case int32(58):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(152)
			goto next_state
		}
		return result
	case int32(59):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(158)
			goto next_state
		}
		return result
	case int32(60):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(156)
			goto next_state
		}
		return result
	case int32(61):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(155)
			goto next_state
		}
		return result
	case int32(62):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(63):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(28)
			goto next_state
		}
		return result
	case int32(64):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(26)
			goto next_state
		}
		return result
	case int32(65):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(27)
			goto next_state
		}
		return result
	case int32(66):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(130)
			goto next_state
		}
		return result
	case int32(67):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(148)
			goto next_state
		}
		return result
	case int32(68):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(66)
			goto next_state
		}
		return result
	case int32(69):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(68)
			goto next_state
		}
		return result
	case int32(70):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(71):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(70)
			goto next_state
		}
		return result
	case int32(72):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(71)
			goto next_state
		}
		return result
	case int32(73):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(72)
			goto next_state
		}
		return result
	case int32(74):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(73)
			goto next_state
		}
		return result
	case int32(75):
		if eof != 0 {
			state = uint16(77)
			goto next_state
		}
		if lookahead == int32('\n') {
			state = uint16(78)
			goto next_state
		}
		if lookahead == int32('\r') {
			state = uint16(6)
			goto next_state
		}
		if lookahead == int32('"') {
			state = uint16(123)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(79)
			goto next_state
		}
		if lookahead == int32('\'') {
			state = uint16(133)
			goto next_state
		}
		if lookahead == int32('+') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(161)
			goto next_state
		}
		if lookahead == int32('-') {
			state = uint16(95)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(86)
			goto next_state
		}
		if lookahead == int32('0') {
			state = uint16(92)
			goto next_state
		}
		if lookahead == int32('1') {
			state = uint16(90)
			goto next_state
		}
		if lookahead == int32('2') {
			state = uint16(89)
			goto next_state
		}
		if lookahead == int32('=') {
			state = uint16(85)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(81)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(82)
			goto next_state
		}
		if lookahead == int32('f') {
			state = uint16(102)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(108)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(103)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(109)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(162)
			goto next_state
		}
		if lookahead == int32('}') {
			state = uint16(163)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(75)
			goto next_state
		}
		if int32('3') <= lookahead && lookahead <= int32('9') {
			state = uint16(91)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(76):
		if eof != 0 {
			state = uint16(77)
			goto next_state
		}
		if lookahead == int32('\n') {
			state = uint16(78)
			goto next_state
		}
		if lookahead == int32('\r') {
			state = uint16(6)
			goto next_state
		}
		if lookahead == int32('"') {
			state = uint16(122)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(79)
			goto next_state
		}
		if lookahead == int32('\'') {
			state = uint16(132)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(161)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(86)
			goto next_state
		}
		if lookahead == int32('=') {
			state = uint16(85)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(81)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(82)
			goto next_state
		}
		if lookahead == int32('}') {
			state = uint16(163)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(76)
			goto next_state
		}
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(77):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(ts_builtin_sym_end)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(78):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_document_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(79):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead > int32(8) && (lookahead < int32('\n') || int32(31) < lookahead) && lookahead != int32(127) {
			state = uint16(79)
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
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('[') {
			state = uint16(83)
			goto next_state
		}
		return result
	case int32(82):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(83):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACK_LBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(84):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACK_RBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(85):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(86):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(87):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_bare_key)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') {
			state = uint16(96)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(46)
			goto next_state
		}
		if lookahead == int32('_') {
			state = uint16(119)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(91)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(88):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_bare_key)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') {
			state = uint16(96)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(46)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(93)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(89):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_bare_key)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') {
			state = uint16(96)
			goto next_state
		}
		if lookahead == int32('_') {
			state = uint16(119)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('3') {
			state = uint16(87)
			goto next_state
		}
		if int32('4') <= lookahead && lookahead <= int32('9') {
			state = uint16(91)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(90):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_bare_key)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') {
			state = uint16(96)
			goto next_state
		}
		if lookahead == int32('_') {
			state = uint16(119)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(87)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(91):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_bare_key)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') {
			state = uint16(96)
			goto next_state
		}
		if lookahead == int32('_') {
			state = uint16(119)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(91)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(92):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_bare_key)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') {
			state = uint16(96)
			goto next_state
		}
		if lookahead == int32('b') {
			state = uint16(113)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(115)
			goto next_state
		}
		if lookahead == int32('x') {
			state = uint16(120)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(88)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(93):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_bare_key)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') {
			state = uint16(96)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(93)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(94):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_bare_key)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') {
			state = uint16(97)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(95):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_bare_key)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('0') {
			state = uint16(121)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(108)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(103)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(101)
			goto next_state
		}
		if lookahead == int32('-') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(96):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_bare_key)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('0') {
			state = uint16(117)
			goto next_state
		}
		if lookahead == int32('1') {
			state = uint16(114)
			goto next_state
		}
		if lookahead == int32('-') || int32('2') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(97):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_bare_key)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('0') {
			state = uint16(116)
			goto next_state
		}
		if lookahead == int32('3') {
			state = uint16(112)
			goto next_state
		}
		if lookahead == int32('1') || lookahead == int32('2') {
			state = uint16(118)
			goto next_state
		}
		if lookahead == int32('-') || int32('4') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(98):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_bare_key)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(113)
			goto next_state
		}
		if lookahead == int32('0') || lookahead == int32('1') {
			state = uint16(98)
			goto next_state
		}
		if lookahead == int32('-') || int32('2') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(99):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_bare_key)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(115)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('7') {
			state = uint16(99)
			goto next_state
		}
		if lookahead == int32('-') || lookahead == int32('8') || lookahead == int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(100):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_bare_key)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(120)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(100)
			goto next_state
		}
		if lookahead == int32('-') || int32('G') <= lookahead && lookahead <= int32('Z') || int32('g') <= lookahead && lookahead <= int32('z') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(101):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_bare_key)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(119)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(101)
			goto next_state
		}
		if lookahead == int32('-') || int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(102):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_bare_key)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('a') {
			state = uint16(106)
			goto next_state
		}
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('b') <= lookahead && lookahead <= int32('z') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(103):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_bare_key)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('a') {
			state = uint16(107)
			goto next_state
		}
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('b') <= lookahead && lookahead <= int32('z') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(104):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_bare_key)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(121)
			goto next_state
		}
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(105):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_bare_key)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('f') {
			state = uint16(121)
			goto next_state
		}
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(106):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_bare_key)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('l') {
			state = uint16(110)
			goto next_state
		}
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(107):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_bare_key)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('n') {
			state = uint16(121)
			goto next_state
		}
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(108):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_bare_key)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('n') {
			state = uint16(105)
			goto next_state
		}
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(109):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_bare_key)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(111)
			goto next_state
		}
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(110):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_bare_key)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('s') {
			state = uint16(104)
			goto next_state
		}
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(111):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_bare_key)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('u') {
			state = uint16(104)
			goto next_state
		}
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(112):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_bare_key)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('0') || lookahead == int32('1') {
			state = uint16(121)
			goto next_state
		}
		if lookahead == int32('-') || int32('2') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(113):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_bare_key)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('0') || lookahead == int32('1') {
			state = uint16(98)
			goto next_state
		}
		if lookahead == int32('-') || int32('2') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(114):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_bare_key)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('2') {
			state = uint16(94)
			goto next_state
		}
		if lookahead == int32('-') || int32('3') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(115):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_bare_key)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('7') {
			state = uint16(99)
			goto next_state
		}
		if lookahead == int32('-') || lookahead == int32('8') || lookahead == int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(116):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_bare_key)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(121)
			goto next_state
		}
		if lookahead == int32('-') || lookahead == int32('0') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(117):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_bare_key)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(94)
			goto next_state
		}
		if lookahead == int32('-') || lookahead == int32('0') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(118):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_bare_key)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(121)
			goto next_state
		}
		if lookahead == int32('-') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(119):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_bare_key)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(101)
			goto next_state
		}
		if lookahead == int32('-') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(120):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_bare_key)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(100)
			goto next_state
		}
		if lookahead == int32('-') || int32('G') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('g') <= lookahead && lookahead <= int32('z') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(121):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_bare_key)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(122):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(123):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('"') {
			state = uint16(8)
			goto next_state
		}
		return result
	case int32(124):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__basic_string_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(125)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			state = uint16(124)
			goto next_state
		}
		if lookahead != 0 && lookahead > int32(31) && lookahead != int32('"') && lookahead != int32('\\') && lookahead != int32(127) {
			state = uint16(125)
			goto next_state
		}
		return result
	case int32(125):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__basic_string_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead > int32(8) && (lookahead < int32('\n') || int32(31) < lookahead) && lookahead != int32('"') && lookahead != int32('\\') && lookahead != int32(127) {
			state = uint16(125)
			goto next_state
		}
		return result
	case int32(126):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DQUOTE2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(127):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DQUOTE2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('"') {
			state = uint16(8)
			goto next_state
		}
		return result
	case int32(128):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(129):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__multiline_basic_string_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(130):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_escape_sequence)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(131):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__escape_line_ending)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(132):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(133):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\'') {
			state = uint16(11)
			goto next_state
		}
		return result
	case int32(134):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__literal_string_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(135)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			state = uint16(134)
			goto next_state
		}
		if lookahead != 0 && lookahead > int32(31) && lookahead != int32('\'') && lookahead != int32(127) {
			state = uint16(135)
			goto next_state
		}
		return result
	case int32(135):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__literal_string_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead > int32(8) && (lookahead < int32('\n') || int32(31) < lookahead) && lookahead != int32('\'') && lookahead != int32(127) {
			state = uint16(135)
			goto next_state
		}
		return result
	case int32(136):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SQUOTE2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(137):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SQUOTE2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\'') {
			state = uint16(11)
			goto next_state
		}
		return result
	case int32(138):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(139):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_integer_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(140):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_integer_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') {
			state = uint16(18)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(57)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(46)
			goto next_state
		}
		if lookahead == int32('_') {
			state = uint16(56)
			goto next_state
		}
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(40)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(143)
			goto next_state
		}
		return result
	case int32(141):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_integer_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') {
			state = uint16(18)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(57)
			goto next_state
		}
		if lookahead == int32('_') {
			state = uint16(56)
			goto next_state
		}
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(40)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('3') {
			state = uint16(140)
			goto next_state
		}
		if int32('4') <= lookahead && lookahead <= int32('9') {
			state = uint16(143)
			goto next_state
		}
		return result
	case int32(142):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_integer_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') {
			state = uint16(18)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(57)
			goto next_state
		}
		if lookahead == int32('_') {
			state = uint16(56)
			goto next_state
		}
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(40)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(143):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_integer_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') {
			state = uint16(18)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(57)
			goto next_state
		}
		if lookahead == int32('_') {
			state = uint16(56)
			goto next_state
		}
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(40)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(143)
			goto next_state
		}
		return result
	case int32(144):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_integer_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') {
			state = uint16(18)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(57)
			goto next_state
		}
		if lookahead == int32('b') {
			state = uint16(41)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(49)
			goto next_state
		}
		if lookahead == int32('x') {
			state = uint16(67)
			goto next_state
		}
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(40)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(12)
			goto next_state
		}
		return result
	case int32(145):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_integer_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(57)
			goto next_state
		}
		if lookahead == int32('_') {
			state = uint16(56)
			goto next_state
		}
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(40)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(145)
			goto next_state
		}
		return result
	case int32(146):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_integer_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(57)
			goto next_state
		}
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(40)
			goto next_state
		}
		return result
	case int32(147):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_integer_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(52)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(147)
			goto next_state
		}
		return result
	case int32(148):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_integer_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(67)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(148)
			goto next_state
		}
		return result
	case int32(149):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_integer_token3)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(49)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('7') {
			state = uint16(149)
			goto next_state
		}
		return result
	case int32(150):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_integer_token4)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(41)
			goto next_state
		}
		if lookahead == int32('0') || lookahead == int32('1') {
			state = uint16(150)
			goto next_state
		}
		return result
	case int32(151):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_float_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(57)
			goto next_state
		}
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(40)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(151)
			goto next_state
		}
		return result
	case int32(152):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_float_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(58)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(152)
			goto next_state
		}
		return result
	case int32(153):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_float_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(154):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_boolean)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(155):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_offset_date_time)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(156):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_local_date_time)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(62)
			goto next_state
		}
		if lookahead == int32('+') || lookahead == int32('-') {
			state = uint16(22)
			goto next_state
		}
		if lookahead == int32('Z') || lookahead == int32('z') {
			state = uint16(155)
			goto next_state
		}
		return result
	case int32(157):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_local_date_time)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('+') || lookahead == int32('-') {
			state = uint16(22)
			goto next_state
		}
		if lookahead == int32('Z') || lookahead == int32('z') {
			state = uint16(155)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(158):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_local_date)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(' ') || lookahead == int32('T') || lookahead == int32('t') {
			state = uint16(21)
			goto next_state
		}
		return result
	case int32(159):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_local_time)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(55)
			goto next_state
		}
		return result
	case int32(160):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_local_time)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(160)
			goto next_state
		}
		return result
	case int32(161):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COMMA)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(162):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(163):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var ts_lex_modes = [152]TSLexMode{
	0: {
		Fexternal_lex_state: uint16(1),
	},
	1: {
		Flex_state: uint16(76),
	},
	2: {
		Flex_state: uint16(7),
	},
	3: {
		Flex_state: uint16(7),
	},
	4: {
		Flex_state: uint16(7),
	},
	5: {
		Flex_state: uint16(7),
	},
	6: {
		Flex_state: uint16(7),
	},
	7: {
		Flex_state: uint16(7),
	},
	8: {
		Flex_state: uint16(7),
	},
	9: {
		Flex_state: uint16(7),
	},
	10: {
		Flex_state: uint16(7),
	},
	11: {
		Flex_state: uint16(7),
	},
	12: {
		Flex_state: uint16(7),
	},
	13: {
		Flex_state: uint16(7),
	},
	14: {
		Flex_state: uint16(7),
	},
	15: {
		Flex_state: uint16(7),
	},
	16: {
		Flex_state: uint16(7),
	},
	17: {
		Flex_state: uint16(7),
	},
	18: {
		Flex_state: uint16(7),
	},
	19: {
		Flex_state: uint16(7),
	},
	20: {
		Flex_state: uint16(7),
	},
	21: {
		Flex_state: uint16(7),
	},
	22: {
		Flex_state: uint16(7),
	},
	23: {
		Flex_state: uint16(7),
	},
	24: {
		Flex_state: uint16(7),
	},
	25: {
		Flex_state: uint16(7),
	},
	26: {
		Flex_state: uint16(7),
	},
	27: {
		Flex_state: uint16(76),
	},
	28: {
		Flex_state: uint16(76),
	},
	29: {
		Flex_state: uint16(76),
	},
	30: {
		Flex_state: uint16(76),
	},
	31: {
		Flex_state: uint16(76),
	},
	32: {
		Flex_state: uint16(76),
	},
	33: {
		Flex_state: uint16(76),
	},
	34: {
		Flex_state: uint16(76),
	},
	35: {
		Flex_state: uint16(76),
	},
	36: {
		Flex_state: uint16(76),
	},
	37: {
		Flex_state: uint16(76),
	},
	38: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(2),
	},
	39: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(2),
	},
	40: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(2),
	},
	41: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(2),
	},
	42: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(2),
	},
	43: {
		Flex_state: uint16(76),
	},
	44: {
		Flex_state: uint16(76),
	},
	45: {
		Flex_state: uint16(76),
	},
	46: {
		Flex_state: uint16(76),
	},
	47: {},
	48: {},
	49: {
		Flex_state: uint16(76),
	},
	50: {
		Flex_state: uint16(76),
	},
	51: {
		Flex_state: uint16(76),
	},
	52: {},
	53: {
		Flex_state: uint16(76),
	},
	54: {
		Flex_state: uint16(76),
	},
	55: {
		Flex_state: uint16(76),
	},
	56: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(3),
	},
	57: {
		Flex_state: uint16(76),
	},
	58: {
		Flex_state: uint16(76),
	},
	59: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(3),
	},
	60: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(3),
	},
	61: {
		Flex_state: uint16(76),
	},
	62: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(3),
	},
	63: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(3),
	},
	64: {
		Flex_state: uint16(76),
	},
	65: {
		Flex_state: uint16(76),
	},
	66: {
		Flex_state: uint16(76),
	},
	67: {
		Flex_state: uint16(76),
	},
	68: {
		Flex_state: uint16(76),
	},
	69: {
		Flex_state: uint16(2),
	},
	70: {
		Flex_state: uint16(2),
	},
	71: {
		Flex_state: uint16(2),
	},
	72: {
		Flex_state: uint16(2),
	},
	73: {
		Flex_state: uint16(2),
	},
	74: {
		Flex_state: uint16(76),
	},
	75: {
		Flex_state: uint16(2),
	},
	76: {
		Flex_state: uint16(2),
	},
	77: {
		Flex_state: uint16(76),
	},
	78: {
		Flex_state: uint16(76),
	},
	79: {
		Flex_state: uint16(76),
	},
	80: {
		Flex_state: uint16(76),
	},
	81: {
		Flex_state: uint16(76),
	},
	82: {
		Flex_state: uint16(76),
	},
	83: {
		Flex_state: uint16(76),
	},
	84: {
		Flex_state: uint16(76),
	},
	85: {
		Flex_state: uint16(76),
	},
	86: {
		Flex_state: uint16(76),
	},
	87: {
		Flex_state: uint16(76),
	},
	88: {
		Flex_state: uint16(76),
	},
	89: {
		Flex_state: uint16(76),
	},
	90: {
		Flex_state: uint16(76),
	},
	91: {
		Flex_state: uint16(76),
	},
	92: {
		Flex_state: uint16(76),
	},
	93: {
		Flex_state: uint16(76),
	},
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
	108: {
		Flex_state: uint16(9),
	},
	109: {},
	110: {},
	111: {
		Flex_state: uint16(9),
	},
	112: {},
	113: {
		Flex_state: uint16(3),
	},
	114: {},
	115: {
		Flex_state: uint16(9),
	},
	116: {
		Flex_state: uint16(9),
	},
	117: {
		Flex_state: uint16(9),
	},
	118: {
		Flex_state: uint16(9),
	},
	119: {
		Flex_state: uint16(3),
	},
	120: {
		Flex_state: uint16(3),
	},
	121: {},
	122: {
		Flex_state: uint16(9),
	},
	123: {
		Fexternal_lex_state: uint16(4),
	},
	124: {
		Fexternal_lex_state: uint16(4),
	},
	125: {
		Fexternal_lex_state: uint16(4),
	},
	126: {
		Fexternal_lex_state: uint16(4),
	},
	127: {
		Fexternal_lex_state: uint16(4),
	},
	128: {
		Fexternal_lex_state: uint16(4),
	},
	129: {
		Fexternal_lex_state: uint16(4),
	},
	130: {
		Fexternal_lex_state: uint16(4),
	},
	131: {
		Fexternal_lex_state: uint16(4),
	},
	132: {
		Fexternal_lex_state: uint16(4),
	},
	133: {
		Fexternal_lex_state: uint16(4),
	},
	134: {
		Fexternal_lex_state: uint16(4),
	},
	135: {
		Flex_state: uint16(9),
	},
	136: {
		Fexternal_lex_state: uint16(4),
	},
	137: {
		Fexternal_lex_state: uint16(4),
	},
	138: {
		Fexternal_lex_state: uint16(4),
	},
	139: {
		Fexternal_lex_state: uint16(4),
	},
	140: {
		Fexternal_lex_state: uint16(4),
	},
	141: {
		Fexternal_lex_state: uint16(4),
	},
	142: {
		Fexternal_lex_state: uint16(4),
	},
	143: {
		Fexternal_lex_state: uint16(4),
	},
	144: {
		Fexternal_lex_state: uint16(4),
	},
	145: {},
	146: {
		Flex_state: uint16(9),
	},
	147: {
		Fexternal_lex_state: uint16(4),
	},
	148: {
		Fexternal_lex_state: uint16(4),
	},
	149: {
		Fexternal_lex_state: uint16(4),
	},
	150: {
		Flex_state: uint16(9),
	},
	151: {
		Fexternal_lex_state: uint16(4),
	},
}

const ts_external_token__line_ending_or_eof = 0
const ts_external_token__multiline_basic_string_content = 1
const ts_external_token__multiline_basic_string_end = 2
const ts_external_token__multiline_literal_string_content = 3
const ts_external_token__multiline_literal_string_end = 4

var ts_external_scanner_symbol_map = [5]TSSymbol{
	0: uint16(sym__line_ending_or_eof),
	1: uint16(sym__multiline_basic_string_content),
	2: uint16(sym__multiline_basic_string_end),
	3: uint16(sym__multiline_literal_string_content),
	4: uint16(sym__multiline_literal_string_end),
}

var ts_external_scanner_states = [5][5]uint8{
	1: {
		0: libc.BoolUint8(true1 != 0),
		1: libc.BoolUint8(true1 != 0),
		2: libc.BoolUint8(true1 != 0),
		3: libc.BoolUint8(true1 != 0),
		4: libc.BoolUint8(true1 != 0),
	},
	2: {
		1: libc.BoolUint8(true1 != 0),
		2: libc.BoolUint8(true1 != 0),
	},
	3: {
		3: libc.BoolUint8(true1 != 0),
		4: libc.BoolUint8(true1 != 0),
	},
	4: {
		0: libc.BoolUint8(true1 != 0),
	},
}

var ts_parse_table = [2][66]uint16_t{
	0: {
		0:  uint16(1),
		1:  uint16(1),
		2:  uint16(3),
		3:  uint16(1),
		4:  uint16(1),
		5:  uint16(1),
		7:  uint16(1),
		8:  uint16(1),
		9:  uint16(1),
		10: uint16(1),
		12: uint16(1),
		13: uint16(1),
		14: uint16(1),
		15: uint16(1),
		16: uint16(1),
		17: uint16(1),
		19: uint16(1),
		20: uint16(1),
		21: uint16(1),
		22: uint16(1),
		23: uint16(1),
		24: uint16(1),
		26: uint16(1),
		27: uint16(1),
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
	},
	1: {
		0:  uint16(5),
		1:  uint16(7),
		2:  uint16(3),
		3:  uint16(9),
		5:  uint16(11),
		9:  uint16(13),
		10: uint16(15),
		17: uint16(17),
		40: uint16(145),
		41: uint16(52),
		42: uint16(52),
		43: uint16(27),
		44: uint16(144),
		45: uint16(121),
		46: uint16(121),
		47: uint16(121),
		50: uint16(100),
		52: uint16(100),
		58: uint16(27),
		59: uint16(52),
	},
}

var ts_small_parse_table = [3265]uint16_t{
	0:    uint16(17),
	1:    uint16(3),
	2:    uint16(1),
	3:    uint16(sym_comment),
	4:    uint16(19),
	5:    uint16(1),
	6:    uint16(aux_sym_document_token1),
	7:    uint16(21),
	8:    uint16(1),
	9:    uint16(anon_sym_LBRACK),
	10:   uint16(23),
	11:   uint16(1),
	12:   uint16(anon_sym_RBRACK),
	13:   uint16(25),
	14:   uint16(1),
	15:   uint16(anon_sym_DQUOTE),
	16:   uint16(27),
	17:   uint16(1),
	18:   uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	19:   uint16(29),
	20:   uint16(1),
	21:   uint16(anon_sym_SQUOTE),
	22:   uint16(31),
	23:   uint16(1),
	24:   uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	25:   uint16(33),
	26:   uint16(1),
	27:   uint16(aux_sym_integer_token1),
	28:   uint16(43),
	29:   uint16(1),
	30:   uint16(anon_sym_LBRACE),
	31:   uint16(26),
	32:   uint16(1),
	33:   uint16(aux_sym_array_repeat1),
	34:   uint16(37),
	35:   uint16(2),
	36:   uint16(aux_sym_float_token1),
	37:   uint16(aux_sym_float_token2),
	38:   uint16(41),
	39:   uint16(2),
	40:   uint16(sym_local_date_time),
	41:   uint16(sym_local_date),
	42:   uint16(35),
	43:   uint16(3),
	44:   uint16(aux_sym_integer_token2),
	45:   uint16(aux_sym_integer_token3),
	46:   uint16(aux_sym_integer_token4),
	47:   uint16(39),
	48:   uint16(3),
	49:   uint16(sym_boolean),
	50:   uint16(sym_offset_date_time),
	51:   uint16(sym_local_time),
	52:   uint16(86),
	53:   uint16(4),
	54:   uint16(sym__basic_string),
	55:   uint16(sym__multiline_basic_string),
	56:   uint16(sym__literal_string),
	57:   uint16(sym__multiline_literal_string),
	58:   uint16(55),
	59:   uint16(6),
	60:   uint16(sym__inline_value),
	61:   uint16(sym_string),
	62:   uint16(sym_integer),
	63:   uint16(sym_float),
	64:   uint16(sym_array),
	65:   uint16(sym_inline_table),
	66:   uint16(17),
	67:   uint16(3),
	68:   uint16(1),
	69:   uint16(sym_comment),
	70:   uint16(19),
	71:   uint16(1),
	72:   uint16(aux_sym_document_token1),
	73:   uint16(21),
	74:   uint16(1),
	75:   uint16(anon_sym_LBRACK),
	76:   uint16(25),
	77:   uint16(1),
	78:   uint16(anon_sym_DQUOTE),
	79:   uint16(27),
	80:   uint16(1),
	81:   uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	82:   uint16(29),
	83:   uint16(1),
	84:   uint16(anon_sym_SQUOTE),
	85:   uint16(31),
	86:   uint16(1),
	87:   uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	88:   uint16(33),
	89:   uint16(1),
	90:   uint16(aux_sym_integer_token1),
	91:   uint16(43),
	92:   uint16(1),
	93:   uint16(anon_sym_LBRACE),
	94:   uint16(45),
	95:   uint16(1),
	96:   uint16(anon_sym_RBRACK),
	97:   uint16(26),
	98:   uint16(1),
	99:   uint16(aux_sym_array_repeat1),
	100:  uint16(37),
	101:  uint16(2),
	102:  uint16(aux_sym_float_token1),
	103:  uint16(aux_sym_float_token2),
	104:  uint16(49),
	105:  uint16(2),
	106:  uint16(sym_local_date_time),
	107:  uint16(sym_local_date),
	108:  uint16(35),
	109:  uint16(3),
	110:  uint16(aux_sym_integer_token2),
	111:  uint16(aux_sym_integer_token3),
	112:  uint16(aux_sym_integer_token4),
	113:  uint16(47),
	114:  uint16(3),
	115:  uint16(sym_boolean),
	116:  uint16(sym_offset_date_time),
	117:  uint16(sym_local_time),
	118:  uint16(86),
	119:  uint16(4),
	120:  uint16(sym__basic_string),
	121:  uint16(sym__multiline_basic_string),
	122:  uint16(sym__literal_string),
	123:  uint16(sym__multiline_literal_string),
	124:  uint16(66),
	125:  uint16(6),
	126:  uint16(sym__inline_value),
	127:  uint16(sym_string),
	128:  uint16(sym_integer),
	129:  uint16(sym_float),
	130:  uint16(sym_array),
	131:  uint16(sym_inline_table),
	132:  uint16(17),
	133:  uint16(3),
	134:  uint16(1),
	135:  uint16(sym_comment),
	136:  uint16(21),
	137:  uint16(1),
	138:  uint16(anon_sym_LBRACK),
	139:  uint16(25),
	140:  uint16(1),
	141:  uint16(anon_sym_DQUOTE),
	142:  uint16(27),
	143:  uint16(1),
	144:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	145:  uint16(29),
	146:  uint16(1),
	147:  uint16(anon_sym_SQUOTE),
	148:  uint16(31),
	149:  uint16(1),
	150:  uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	151:  uint16(33),
	152:  uint16(1),
	153:  uint16(aux_sym_integer_token1),
	154:  uint16(43),
	155:  uint16(1),
	156:  uint16(anon_sym_LBRACE),
	157:  uint16(51),
	158:  uint16(1),
	159:  uint16(aux_sym_document_token1),
	160:  uint16(53),
	161:  uint16(1),
	162:  uint16(anon_sym_RBRACK),
	163:  uint16(15),
	164:  uint16(1),
	165:  uint16(aux_sym_array_repeat1),
	166:  uint16(37),
	167:  uint16(2),
	168:  uint16(aux_sym_float_token1),
	169:  uint16(aux_sym_float_token2),
	170:  uint16(57),
	171:  uint16(2),
	172:  uint16(sym_local_date_time),
	173:  uint16(sym_local_date),
	174:  uint16(35),
	175:  uint16(3),
	176:  uint16(aux_sym_integer_token2),
	177:  uint16(aux_sym_integer_token3),
	178:  uint16(aux_sym_integer_token4),
	179:  uint16(55),
	180:  uint16(3),
	181:  uint16(sym_boolean),
	182:  uint16(sym_offset_date_time),
	183:  uint16(sym_local_time),
	184:  uint16(86),
	185:  uint16(4),
	186:  uint16(sym__basic_string),
	187:  uint16(sym__multiline_basic_string),
	188:  uint16(sym__literal_string),
	189:  uint16(sym__multiline_literal_string),
	190:  uint16(68),
	191:  uint16(6),
	192:  uint16(sym__inline_value),
	193:  uint16(sym_string),
	194:  uint16(sym_integer),
	195:  uint16(sym_float),
	196:  uint16(sym_array),
	197:  uint16(sym_inline_table),
	198:  uint16(17),
	199:  uint16(3),
	200:  uint16(1),
	201:  uint16(sym_comment),
	202:  uint16(21),
	203:  uint16(1),
	204:  uint16(anon_sym_LBRACK),
	205:  uint16(25),
	206:  uint16(1),
	207:  uint16(anon_sym_DQUOTE),
	208:  uint16(27),
	209:  uint16(1),
	210:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	211:  uint16(29),
	212:  uint16(1),
	213:  uint16(anon_sym_SQUOTE),
	214:  uint16(31),
	215:  uint16(1),
	216:  uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	217:  uint16(33),
	218:  uint16(1),
	219:  uint16(aux_sym_integer_token1),
	220:  uint16(43),
	221:  uint16(1),
	222:  uint16(anon_sym_LBRACE),
	223:  uint16(59),
	224:  uint16(1),
	225:  uint16(aux_sym_document_token1),
	226:  uint16(61),
	227:  uint16(1),
	228:  uint16(anon_sym_RBRACK),
	229:  uint16(14),
	230:  uint16(1),
	231:  uint16(aux_sym_array_repeat1),
	232:  uint16(37),
	233:  uint16(2),
	234:  uint16(aux_sym_float_token1),
	235:  uint16(aux_sym_float_token2),
	236:  uint16(57),
	237:  uint16(2),
	238:  uint16(sym_local_date_time),
	239:  uint16(sym_local_date),
	240:  uint16(35),
	241:  uint16(3),
	242:  uint16(aux_sym_integer_token2),
	243:  uint16(aux_sym_integer_token3),
	244:  uint16(aux_sym_integer_token4),
	245:  uint16(55),
	246:  uint16(3),
	247:  uint16(sym_boolean),
	248:  uint16(sym_offset_date_time),
	249:  uint16(sym_local_time),
	250:  uint16(86),
	251:  uint16(4),
	252:  uint16(sym__basic_string),
	253:  uint16(sym__multiline_basic_string),
	254:  uint16(sym__literal_string),
	255:  uint16(sym__multiline_literal_string),
	256:  uint16(68),
	257:  uint16(6),
	258:  uint16(sym__inline_value),
	259:  uint16(sym_string),
	260:  uint16(sym_integer),
	261:  uint16(sym_float),
	262:  uint16(sym_array),
	263:  uint16(sym_inline_table),
	264:  uint16(17),
	265:  uint16(3),
	266:  uint16(1),
	267:  uint16(sym_comment),
	268:  uint16(19),
	269:  uint16(1),
	270:  uint16(aux_sym_document_token1),
	271:  uint16(21),
	272:  uint16(1),
	273:  uint16(anon_sym_LBRACK),
	274:  uint16(25),
	275:  uint16(1),
	276:  uint16(anon_sym_DQUOTE),
	277:  uint16(27),
	278:  uint16(1),
	279:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	280:  uint16(29),
	281:  uint16(1),
	282:  uint16(anon_sym_SQUOTE),
	283:  uint16(31),
	284:  uint16(1),
	285:  uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	286:  uint16(33),
	287:  uint16(1),
	288:  uint16(aux_sym_integer_token1),
	289:  uint16(43),
	290:  uint16(1),
	291:  uint16(anon_sym_LBRACE),
	292:  uint16(61),
	293:  uint16(1),
	294:  uint16(anon_sym_RBRACK),
	295:  uint16(26),
	296:  uint16(1),
	297:  uint16(aux_sym_array_repeat1),
	298:  uint16(37),
	299:  uint16(2),
	300:  uint16(aux_sym_float_token1),
	301:  uint16(aux_sym_float_token2),
	302:  uint16(49),
	303:  uint16(2),
	304:  uint16(sym_local_date_time),
	305:  uint16(sym_local_date),
	306:  uint16(35),
	307:  uint16(3),
	308:  uint16(aux_sym_integer_token2),
	309:  uint16(aux_sym_integer_token3),
	310:  uint16(aux_sym_integer_token4),
	311:  uint16(47),
	312:  uint16(3),
	313:  uint16(sym_boolean),
	314:  uint16(sym_offset_date_time),
	315:  uint16(sym_local_time),
	316:  uint16(86),
	317:  uint16(4),
	318:  uint16(sym__basic_string),
	319:  uint16(sym__multiline_basic_string),
	320:  uint16(sym__literal_string),
	321:  uint16(sym__multiline_literal_string),
	322:  uint16(66),
	323:  uint16(6),
	324:  uint16(sym__inline_value),
	325:  uint16(sym_string),
	326:  uint16(sym_integer),
	327:  uint16(sym_float),
	328:  uint16(sym_array),
	329:  uint16(sym_inline_table),
	330:  uint16(17),
	331:  uint16(3),
	332:  uint16(1),
	333:  uint16(sym_comment),
	334:  uint16(21),
	335:  uint16(1),
	336:  uint16(anon_sym_LBRACK),
	337:  uint16(25),
	338:  uint16(1),
	339:  uint16(anon_sym_DQUOTE),
	340:  uint16(27),
	341:  uint16(1),
	342:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	343:  uint16(29),
	344:  uint16(1),
	345:  uint16(anon_sym_SQUOTE),
	346:  uint16(31),
	347:  uint16(1),
	348:  uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	349:  uint16(33),
	350:  uint16(1),
	351:  uint16(aux_sym_integer_token1),
	352:  uint16(43),
	353:  uint16(1),
	354:  uint16(anon_sym_LBRACE),
	355:  uint16(63),
	356:  uint16(1),
	357:  uint16(aux_sym_document_token1),
	358:  uint16(65),
	359:  uint16(1),
	360:  uint16(anon_sym_RBRACK),
	361:  uint16(6),
	362:  uint16(1),
	363:  uint16(aux_sym_array_repeat1),
	364:  uint16(37),
	365:  uint16(2),
	366:  uint16(aux_sym_float_token1),
	367:  uint16(aux_sym_float_token2),
	368:  uint16(57),
	369:  uint16(2),
	370:  uint16(sym_local_date_time),
	371:  uint16(sym_local_date),
	372:  uint16(35),
	373:  uint16(3),
	374:  uint16(aux_sym_integer_token2),
	375:  uint16(aux_sym_integer_token3),
	376:  uint16(aux_sym_integer_token4),
	377:  uint16(55),
	378:  uint16(3),
	379:  uint16(sym_boolean),
	380:  uint16(sym_offset_date_time),
	381:  uint16(sym_local_time),
	382:  uint16(86),
	383:  uint16(4),
	384:  uint16(sym__basic_string),
	385:  uint16(sym__multiline_basic_string),
	386:  uint16(sym__literal_string),
	387:  uint16(sym__multiline_literal_string),
	388:  uint16(68),
	389:  uint16(6),
	390:  uint16(sym__inline_value),
	391:  uint16(sym_string),
	392:  uint16(sym_integer),
	393:  uint16(sym_float),
	394:  uint16(sym_array),
	395:  uint16(sym_inline_table),
	396:  uint16(17),
	397:  uint16(3),
	398:  uint16(1),
	399:  uint16(sym_comment),
	400:  uint16(19),
	401:  uint16(1),
	402:  uint16(aux_sym_document_token1),
	403:  uint16(21),
	404:  uint16(1),
	405:  uint16(anon_sym_LBRACK),
	406:  uint16(25),
	407:  uint16(1),
	408:  uint16(anon_sym_DQUOTE),
	409:  uint16(27),
	410:  uint16(1),
	411:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	412:  uint16(29),
	413:  uint16(1),
	414:  uint16(anon_sym_SQUOTE),
	415:  uint16(31),
	416:  uint16(1),
	417:  uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	418:  uint16(33),
	419:  uint16(1),
	420:  uint16(aux_sym_integer_token1),
	421:  uint16(43),
	422:  uint16(1),
	423:  uint16(anon_sym_LBRACE),
	424:  uint16(65),
	425:  uint16(1),
	426:  uint16(anon_sym_RBRACK),
	427:  uint16(26),
	428:  uint16(1),
	429:  uint16(aux_sym_array_repeat1),
	430:  uint16(37),
	431:  uint16(2),
	432:  uint16(aux_sym_float_token1),
	433:  uint16(aux_sym_float_token2),
	434:  uint16(49),
	435:  uint16(2),
	436:  uint16(sym_local_date_time),
	437:  uint16(sym_local_date),
	438:  uint16(35),
	439:  uint16(3),
	440:  uint16(aux_sym_integer_token2),
	441:  uint16(aux_sym_integer_token3),
	442:  uint16(aux_sym_integer_token4),
	443:  uint16(47),
	444:  uint16(3),
	445:  uint16(sym_boolean),
	446:  uint16(sym_offset_date_time),
	447:  uint16(sym_local_time),
	448:  uint16(86),
	449:  uint16(4),
	450:  uint16(sym__basic_string),
	451:  uint16(sym__multiline_basic_string),
	452:  uint16(sym__literal_string),
	453:  uint16(sym__multiline_literal_string),
	454:  uint16(66),
	455:  uint16(6),
	456:  uint16(sym__inline_value),
	457:  uint16(sym_string),
	458:  uint16(sym_integer),
	459:  uint16(sym_float),
	460:  uint16(sym_array),
	461:  uint16(sym_inline_table),
	462:  uint16(17),
	463:  uint16(3),
	464:  uint16(1),
	465:  uint16(sym_comment),
	466:  uint16(21),
	467:  uint16(1),
	468:  uint16(anon_sym_LBRACK),
	469:  uint16(25),
	470:  uint16(1),
	471:  uint16(anon_sym_DQUOTE),
	472:  uint16(27),
	473:  uint16(1),
	474:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	475:  uint16(29),
	476:  uint16(1),
	477:  uint16(anon_sym_SQUOTE),
	478:  uint16(31),
	479:  uint16(1),
	480:  uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	481:  uint16(33),
	482:  uint16(1),
	483:  uint16(aux_sym_integer_token1),
	484:  uint16(43),
	485:  uint16(1),
	486:  uint16(anon_sym_LBRACE),
	487:  uint16(67),
	488:  uint16(1),
	489:  uint16(aux_sym_document_token1),
	490:  uint16(69),
	491:  uint16(1),
	492:  uint16(anon_sym_RBRACK),
	493:  uint16(10),
	494:  uint16(1),
	495:  uint16(aux_sym_array_repeat1),
	496:  uint16(37),
	497:  uint16(2),
	498:  uint16(aux_sym_float_token1),
	499:  uint16(aux_sym_float_token2),
	500:  uint16(57),
	501:  uint16(2),
	502:  uint16(sym_local_date_time),
	503:  uint16(sym_local_date),
	504:  uint16(35),
	505:  uint16(3),
	506:  uint16(aux_sym_integer_token2),
	507:  uint16(aux_sym_integer_token3),
	508:  uint16(aux_sym_integer_token4),
	509:  uint16(55),
	510:  uint16(3),
	511:  uint16(sym_boolean),
	512:  uint16(sym_offset_date_time),
	513:  uint16(sym_local_time),
	514:  uint16(86),
	515:  uint16(4),
	516:  uint16(sym__basic_string),
	517:  uint16(sym__multiline_basic_string),
	518:  uint16(sym__literal_string),
	519:  uint16(sym__multiline_literal_string),
	520:  uint16(68),
	521:  uint16(6),
	522:  uint16(sym__inline_value),
	523:  uint16(sym_string),
	524:  uint16(sym_integer),
	525:  uint16(sym_float),
	526:  uint16(sym_array),
	527:  uint16(sym_inline_table),
	528:  uint16(17),
	529:  uint16(3),
	530:  uint16(1),
	531:  uint16(sym_comment),
	532:  uint16(19),
	533:  uint16(1),
	534:  uint16(aux_sym_document_token1),
	535:  uint16(21),
	536:  uint16(1),
	537:  uint16(anon_sym_LBRACK),
	538:  uint16(25),
	539:  uint16(1),
	540:  uint16(anon_sym_DQUOTE),
	541:  uint16(27),
	542:  uint16(1),
	543:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	544:  uint16(29),
	545:  uint16(1),
	546:  uint16(anon_sym_SQUOTE),
	547:  uint16(31),
	548:  uint16(1),
	549:  uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	550:  uint16(33),
	551:  uint16(1),
	552:  uint16(aux_sym_integer_token1),
	553:  uint16(43),
	554:  uint16(1),
	555:  uint16(anon_sym_LBRACE),
	556:  uint16(53),
	557:  uint16(1),
	558:  uint16(anon_sym_RBRACK),
	559:  uint16(26),
	560:  uint16(1),
	561:  uint16(aux_sym_array_repeat1),
	562:  uint16(37),
	563:  uint16(2),
	564:  uint16(aux_sym_float_token1),
	565:  uint16(aux_sym_float_token2),
	566:  uint16(49),
	567:  uint16(2),
	568:  uint16(sym_local_date_time),
	569:  uint16(sym_local_date),
	570:  uint16(35),
	571:  uint16(3),
	572:  uint16(aux_sym_integer_token2),
	573:  uint16(aux_sym_integer_token3),
	574:  uint16(aux_sym_integer_token4),
	575:  uint16(47),
	576:  uint16(3),
	577:  uint16(sym_boolean),
	578:  uint16(sym_offset_date_time),
	579:  uint16(sym_local_time),
	580:  uint16(86),
	581:  uint16(4),
	582:  uint16(sym__basic_string),
	583:  uint16(sym__multiline_basic_string),
	584:  uint16(sym__literal_string),
	585:  uint16(sym__multiline_literal_string),
	586:  uint16(66),
	587:  uint16(6),
	588:  uint16(sym__inline_value),
	589:  uint16(sym_string),
	590:  uint16(sym_integer),
	591:  uint16(sym_float),
	592:  uint16(sym_array),
	593:  uint16(sym_inline_table),
	594:  uint16(17),
	595:  uint16(3),
	596:  uint16(1),
	597:  uint16(sym_comment),
	598:  uint16(21),
	599:  uint16(1),
	600:  uint16(anon_sym_LBRACK),
	601:  uint16(25),
	602:  uint16(1),
	603:  uint16(anon_sym_DQUOTE),
	604:  uint16(27),
	605:  uint16(1),
	606:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	607:  uint16(29),
	608:  uint16(1),
	609:  uint16(anon_sym_SQUOTE),
	610:  uint16(31),
	611:  uint16(1),
	612:  uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	613:  uint16(33),
	614:  uint16(1),
	615:  uint16(aux_sym_integer_token1),
	616:  uint16(43),
	617:  uint16(1),
	618:  uint16(anon_sym_LBRACE),
	619:  uint16(71),
	620:  uint16(1),
	621:  uint16(aux_sym_document_token1),
	622:  uint16(73),
	623:  uint16(1),
	624:  uint16(anon_sym_RBRACK),
	625:  uint16(18),
	626:  uint16(1),
	627:  uint16(aux_sym_array_repeat1),
	628:  uint16(37),
	629:  uint16(2),
	630:  uint16(aux_sym_float_token1),
	631:  uint16(aux_sym_float_token2),
	632:  uint16(57),
	633:  uint16(2),
	634:  uint16(sym_local_date_time),
	635:  uint16(sym_local_date),
	636:  uint16(35),
	637:  uint16(3),
	638:  uint16(aux_sym_integer_token2),
	639:  uint16(aux_sym_integer_token3),
	640:  uint16(aux_sym_integer_token4),
	641:  uint16(55),
	642:  uint16(3),
	643:  uint16(sym_boolean),
	644:  uint16(sym_offset_date_time),
	645:  uint16(sym_local_time),
	646:  uint16(86),
	647:  uint16(4),
	648:  uint16(sym__basic_string),
	649:  uint16(sym__multiline_basic_string),
	650:  uint16(sym__literal_string),
	651:  uint16(sym__multiline_literal_string),
	652:  uint16(68),
	653:  uint16(6),
	654:  uint16(sym__inline_value),
	655:  uint16(sym_string),
	656:  uint16(sym_integer),
	657:  uint16(sym_float),
	658:  uint16(sym_array),
	659:  uint16(sym_inline_table),
	660:  uint16(17),
	661:  uint16(3),
	662:  uint16(1),
	663:  uint16(sym_comment),
	664:  uint16(21),
	665:  uint16(1),
	666:  uint16(anon_sym_LBRACK),
	667:  uint16(25),
	668:  uint16(1),
	669:  uint16(anon_sym_DQUOTE),
	670:  uint16(27),
	671:  uint16(1),
	672:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	673:  uint16(29),
	674:  uint16(1),
	675:  uint16(anon_sym_SQUOTE),
	676:  uint16(31),
	677:  uint16(1),
	678:  uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	679:  uint16(33),
	680:  uint16(1),
	681:  uint16(aux_sym_integer_token1),
	682:  uint16(43),
	683:  uint16(1),
	684:  uint16(anon_sym_LBRACE),
	685:  uint16(75),
	686:  uint16(1),
	687:  uint16(aux_sym_document_token1),
	688:  uint16(77),
	689:  uint16(1),
	690:  uint16(anon_sym_RBRACK),
	691:  uint16(8),
	692:  uint16(1),
	693:  uint16(aux_sym_array_repeat1),
	694:  uint16(37),
	695:  uint16(2),
	696:  uint16(aux_sym_float_token1),
	697:  uint16(aux_sym_float_token2),
	698:  uint16(57),
	699:  uint16(2),
	700:  uint16(sym_local_date_time),
	701:  uint16(sym_local_date),
	702:  uint16(35),
	703:  uint16(3),
	704:  uint16(aux_sym_integer_token2),
	705:  uint16(aux_sym_integer_token3),
	706:  uint16(aux_sym_integer_token4),
	707:  uint16(55),
	708:  uint16(3),
	709:  uint16(sym_boolean),
	710:  uint16(sym_offset_date_time),
	711:  uint16(sym_local_time),
	712:  uint16(86),
	713:  uint16(4),
	714:  uint16(sym__basic_string),
	715:  uint16(sym__multiline_basic_string),
	716:  uint16(sym__literal_string),
	717:  uint16(sym__multiline_literal_string),
	718:  uint16(68),
	719:  uint16(6),
	720:  uint16(sym__inline_value),
	721:  uint16(sym_string),
	722:  uint16(sym_integer),
	723:  uint16(sym_float),
	724:  uint16(sym_array),
	725:  uint16(sym_inline_table),
	726:  uint16(17),
	727:  uint16(3),
	728:  uint16(1),
	729:  uint16(sym_comment),
	730:  uint16(19),
	731:  uint16(1),
	732:  uint16(aux_sym_document_token1),
	733:  uint16(21),
	734:  uint16(1),
	735:  uint16(anon_sym_LBRACK),
	736:  uint16(25),
	737:  uint16(1),
	738:  uint16(anon_sym_DQUOTE),
	739:  uint16(27),
	740:  uint16(1),
	741:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	742:  uint16(29),
	743:  uint16(1),
	744:  uint16(anon_sym_SQUOTE),
	745:  uint16(31),
	746:  uint16(1),
	747:  uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	748:  uint16(33),
	749:  uint16(1),
	750:  uint16(aux_sym_integer_token1),
	751:  uint16(43),
	752:  uint16(1),
	753:  uint16(anon_sym_LBRACE),
	754:  uint16(77),
	755:  uint16(1),
	756:  uint16(anon_sym_RBRACK),
	757:  uint16(26),
	758:  uint16(1),
	759:  uint16(aux_sym_array_repeat1),
	760:  uint16(37),
	761:  uint16(2),
	762:  uint16(aux_sym_float_token1),
	763:  uint16(aux_sym_float_token2),
	764:  uint16(49),
	765:  uint16(2),
	766:  uint16(sym_local_date_time),
	767:  uint16(sym_local_date),
	768:  uint16(35),
	769:  uint16(3),
	770:  uint16(aux_sym_integer_token2),
	771:  uint16(aux_sym_integer_token3),
	772:  uint16(aux_sym_integer_token4),
	773:  uint16(47),
	774:  uint16(3),
	775:  uint16(sym_boolean),
	776:  uint16(sym_offset_date_time),
	777:  uint16(sym_local_time),
	778:  uint16(86),
	779:  uint16(4),
	780:  uint16(sym__basic_string),
	781:  uint16(sym__multiline_basic_string),
	782:  uint16(sym__literal_string),
	783:  uint16(sym__multiline_literal_string),
	784:  uint16(66),
	785:  uint16(6),
	786:  uint16(sym__inline_value),
	787:  uint16(sym_string),
	788:  uint16(sym_integer),
	789:  uint16(sym_float),
	790:  uint16(sym_array),
	791:  uint16(sym_inline_table),
	792:  uint16(17),
	793:  uint16(3),
	794:  uint16(1),
	795:  uint16(sym_comment),
	796:  uint16(19),
	797:  uint16(1),
	798:  uint16(aux_sym_document_token1),
	799:  uint16(21),
	800:  uint16(1),
	801:  uint16(anon_sym_LBRACK),
	802:  uint16(25),
	803:  uint16(1),
	804:  uint16(anon_sym_DQUOTE),
	805:  uint16(27),
	806:  uint16(1),
	807:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	808:  uint16(29),
	809:  uint16(1),
	810:  uint16(anon_sym_SQUOTE),
	811:  uint16(31),
	812:  uint16(1),
	813:  uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	814:  uint16(33),
	815:  uint16(1),
	816:  uint16(aux_sym_integer_token1),
	817:  uint16(43),
	818:  uint16(1),
	819:  uint16(anon_sym_LBRACE),
	820:  uint16(79),
	821:  uint16(1),
	822:  uint16(anon_sym_RBRACK),
	823:  uint16(26),
	824:  uint16(1),
	825:  uint16(aux_sym_array_repeat1),
	826:  uint16(37),
	827:  uint16(2),
	828:  uint16(aux_sym_float_token1),
	829:  uint16(aux_sym_float_token2),
	830:  uint16(49),
	831:  uint16(2),
	832:  uint16(sym_local_date_time),
	833:  uint16(sym_local_date),
	834:  uint16(35),
	835:  uint16(3),
	836:  uint16(aux_sym_integer_token2),
	837:  uint16(aux_sym_integer_token3),
	838:  uint16(aux_sym_integer_token4),
	839:  uint16(47),
	840:  uint16(3),
	841:  uint16(sym_boolean),
	842:  uint16(sym_offset_date_time),
	843:  uint16(sym_local_time),
	844:  uint16(86),
	845:  uint16(4),
	846:  uint16(sym__basic_string),
	847:  uint16(sym__multiline_basic_string),
	848:  uint16(sym__literal_string),
	849:  uint16(sym__multiline_literal_string),
	850:  uint16(66),
	851:  uint16(6),
	852:  uint16(sym__inline_value),
	853:  uint16(sym_string),
	854:  uint16(sym_integer),
	855:  uint16(sym_float),
	856:  uint16(sym_array),
	857:  uint16(sym_inline_table),
	858:  uint16(17),
	859:  uint16(3),
	860:  uint16(1),
	861:  uint16(sym_comment),
	862:  uint16(19),
	863:  uint16(1),
	864:  uint16(aux_sym_document_token1),
	865:  uint16(21),
	866:  uint16(1),
	867:  uint16(anon_sym_LBRACK),
	868:  uint16(25),
	869:  uint16(1),
	870:  uint16(anon_sym_DQUOTE),
	871:  uint16(27),
	872:  uint16(1),
	873:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	874:  uint16(29),
	875:  uint16(1),
	876:  uint16(anon_sym_SQUOTE),
	877:  uint16(31),
	878:  uint16(1),
	879:  uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	880:  uint16(33),
	881:  uint16(1),
	882:  uint16(aux_sym_integer_token1),
	883:  uint16(43),
	884:  uint16(1),
	885:  uint16(anon_sym_LBRACE),
	886:  uint16(81),
	887:  uint16(1),
	888:  uint16(anon_sym_RBRACK),
	889:  uint16(26),
	890:  uint16(1),
	891:  uint16(aux_sym_array_repeat1),
	892:  uint16(37),
	893:  uint16(2),
	894:  uint16(aux_sym_float_token1),
	895:  uint16(aux_sym_float_token2),
	896:  uint16(49),
	897:  uint16(2),
	898:  uint16(sym_local_date_time),
	899:  uint16(sym_local_date),
	900:  uint16(35),
	901:  uint16(3),
	902:  uint16(aux_sym_integer_token2),
	903:  uint16(aux_sym_integer_token3),
	904:  uint16(aux_sym_integer_token4),
	905:  uint16(47),
	906:  uint16(3),
	907:  uint16(sym_boolean),
	908:  uint16(sym_offset_date_time),
	909:  uint16(sym_local_time),
	910:  uint16(86),
	911:  uint16(4),
	912:  uint16(sym__basic_string),
	913:  uint16(sym__multiline_basic_string),
	914:  uint16(sym__literal_string),
	915:  uint16(sym__multiline_literal_string),
	916:  uint16(66),
	917:  uint16(6),
	918:  uint16(sym__inline_value),
	919:  uint16(sym_string),
	920:  uint16(sym_integer),
	921:  uint16(sym_float),
	922:  uint16(sym_array),
	923:  uint16(sym_inline_table),
	924:  uint16(17),
	925:  uint16(3),
	926:  uint16(1),
	927:  uint16(sym_comment),
	928:  uint16(21),
	929:  uint16(1),
	930:  uint16(anon_sym_LBRACK),
	931:  uint16(25),
	932:  uint16(1),
	933:  uint16(anon_sym_DQUOTE),
	934:  uint16(27),
	935:  uint16(1),
	936:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	937:  uint16(29),
	938:  uint16(1),
	939:  uint16(anon_sym_SQUOTE),
	940:  uint16(31),
	941:  uint16(1),
	942:  uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	943:  uint16(33),
	944:  uint16(1),
	945:  uint16(aux_sym_integer_token1),
	946:  uint16(43),
	947:  uint16(1),
	948:  uint16(anon_sym_LBRACE),
	949:  uint16(83),
	950:  uint16(1),
	951:  uint16(aux_sym_document_token1),
	952:  uint16(85),
	953:  uint16(1),
	954:  uint16(anon_sym_RBRACK),
	955:  uint16(2),
	956:  uint16(1),
	957:  uint16(aux_sym_array_repeat1),
	958:  uint16(37),
	959:  uint16(2),
	960:  uint16(aux_sym_float_token1),
	961:  uint16(aux_sym_float_token2),
	962:  uint16(89),
	963:  uint16(2),
	964:  uint16(sym_local_date_time),
	965:  uint16(sym_local_date),
	966:  uint16(35),
	967:  uint16(3),
	968:  uint16(aux_sym_integer_token2),
	969:  uint16(aux_sym_integer_token3),
	970:  uint16(aux_sym_integer_token4),
	971:  uint16(87),
	972:  uint16(3),
	973:  uint16(sym_boolean),
	974:  uint16(sym_offset_date_time),
	975:  uint16(sym_local_time),
	976:  uint16(86),
	977:  uint16(4),
	978:  uint16(sym__basic_string),
	979:  uint16(sym__multiline_basic_string),
	980:  uint16(sym__literal_string),
	981:  uint16(sym__multiline_literal_string),
	982:  uint16(54),
	983:  uint16(6),
	984:  uint16(sym__inline_value),
	985:  uint16(sym_string),
	986:  uint16(sym_integer),
	987:  uint16(sym_float),
	988:  uint16(sym_array),
	989:  uint16(sym_inline_table),
	990:  uint16(17),
	991:  uint16(3),
	992:  uint16(1),
	993:  uint16(sym_comment),
	994:  uint16(21),
	995:  uint16(1),
	996:  uint16(anon_sym_LBRACK),
	997:  uint16(25),
	998:  uint16(1),
	999:  uint16(anon_sym_DQUOTE),
	1000: uint16(27),
	1001: uint16(1),
	1002: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1003: uint16(29),
	1004: uint16(1),
	1005: uint16(anon_sym_SQUOTE),
	1006: uint16(31),
	1007: uint16(1),
	1008: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	1009: uint16(33),
	1010: uint16(1),
	1011: uint16(aux_sym_integer_token1),
	1012: uint16(43),
	1013: uint16(1),
	1014: uint16(anon_sym_LBRACE),
	1015: uint16(81),
	1016: uint16(1),
	1017: uint16(anon_sym_RBRACK),
	1018: uint16(91),
	1019: uint16(1),
	1020: uint16(aux_sym_document_token1),
	1021: uint16(3),
	1022: uint16(1),
	1023: uint16(aux_sym_array_repeat1),
	1024: uint16(37),
	1025: uint16(2),
	1026: uint16(aux_sym_float_token1),
	1027: uint16(aux_sym_float_token2),
	1028: uint16(57),
	1029: uint16(2),
	1030: uint16(sym_local_date_time),
	1031: uint16(sym_local_date),
	1032: uint16(35),
	1033: uint16(3),
	1034: uint16(aux_sym_integer_token2),
	1035: uint16(aux_sym_integer_token3),
	1036: uint16(aux_sym_integer_token4),
	1037: uint16(55),
	1038: uint16(3),
	1039: uint16(sym_boolean),
	1040: uint16(sym_offset_date_time),
	1041: uint16(sym_local_time),
	1042: uint16(86),
	1043: uint16(4),
	1044: uint16(sym__basic_string),
	1045: uint16(sym__multiline_basic_string),
	1046: uint16(sym__literal_string),
	1047: uint16(sym__multiline_literal_string),
	1048: uint16(68),
	1049: uint16(6),
	1050: uint16(sym__inline_value),
	1051: uint16(sym_string),
	1052: uint16(sym_integer),
	1053: uint16(sym_float),
	1054: uint16(sym_array),
	1055: uint16(sym_inline_table),
	1056: uint16(17),
	1057: uint16(3),
	1058: uint16(1),
	1059: uint16(sym_comment),
	1060: uint16(19),
	1061: uint16(1),
	1062: uint16(aux_sym_document_token1),
	1063: uint16(21),
	1064: uint16(1),
	1065: uint16(anon_sym_LBRACK),
	1066: uint16(25),
	1067: uint16(1),
	1068: uint16(anon_sym_DQUOTE),
	1069: uint16(27),
	1070: uint16(1),
	1071: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1072: uint16(29),
	1073: uint16(1),
	1074: uint16(anon_sym_SQUOTE),
	1075: uint16(31),
	1076: uint16(1),
	1077: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	1078: uint16(33),
	1079: uint16(1),
	1080: uint16(aux_sym_integer_token1),
	1081: uint16(43),
	1082: uint16(1),
	1083: uint16(anon_sym_LBRACE),
	1084: uint16(69),
	1085: uint16(1),
	1086: uint16(anon_sym_RBRACK),
	1087: uint16(26),
	1088: uint16(1),
	1089: uint16(aux_sym_array_repeat1),
	1090: uint16(37),
	1091: uint16(2),
	1092: uint16(aux_sym_float_token1),
	1093: uint16(aux_sym_float_token2),
	1094: uint16(49),
	1095: uint16(2),
	1096: uint16(sym_local_date_time),
	1097: uint16(sym_local_date),
	1098: uint16(35),
	1099: uint16(3),
	1100: uint16(aux_sym_integer_token2),
	1101: uint16(aux_sym_integer_token3),
	1102: uint16(aux_sym_integer_token4),
	1103: uint16(47),
	1104: uint16(3),
	1105: uint16(sym_boolean),
	1106: uint16(sym_offset_date_time),
	1107: uint16(sym_local_time),
	1108: uint16(86),
	1109: uint16(4),
	1110: uint16(sym__basic_string),
	1111: uint16(sym__multiline_basic_string),
	1112: uint16(sym__literal_string),
	1113: uint16(sym__multiline_literal_string),
	1114: uint16(66),
	1115: uint16(6),
	1116: uint16(sym__inline_value),
	1117: uint16(sym_string),
	1118: uint16(sym_integer),
	1119: uint16(sym_float),
	1120: uint16(sym_array),
	1121: uint16(sym_inline_table),
	1122: uint16(17),
	1123: uint16(3),
	1124: uint16(1),
	1125: uint16(sym_comment),
	1126: uint16(21),
	1127: uint16(1),
	1128: uint16(anon_sym_LBRACK),
	1129: uint16(25),
	1130: uint16(1),
	1131: uint16(anon_sym_DQUOTE),
	1132: uint16(27),
	1133: uint16(1),
	1134: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1135: uint16(29),
	1136: uint16(1),
	1137: uint16(anon_sym_SQUOTE),
	1138: uint16(31),
	1139: uint16(1),
	1140: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	1141: uint16(33),
	1142: uint16(1),
	1143: uint16(aux_sym_integer_token1),
	1144: uint16(43),
	1145: uint16(1),
	1146: uint16(anon_sym_LBRACE),
	1147: uint16(93),
	1148: uint16(1),
	1149: uint16(aux_sym_document_token1),
	1150: uint16(95),
	1151: uint16(1),
	1152: uint16(anon_sym_RBRACK),
	1153: uint16(13),
	1154: uint16(1),
	1155: uint16(aux_sym_array_repeat1),
	1156: uint16(37),
	1157: uint16(2),
	1158: uint16(aux_sym_float_token1),
	1159: uint16(aux_sym_float_token2),
	1160: uint16(57),
	1161: uint16(2),
	1162: uint16(sym_local_date_time),
	1163: uint16(sym_local_date),
	1164: uint16(35),
	1165: uint16(3),
	1166: uint16(aux_sym_integer_token2),
	1167: uint16(aux_sym_integer_token3),
	1168: uint16(aux_sym_integer_token4),
	1169: uint16(55),
	1170: uint16(3),
	1171: uint16(sym_boolean),
	1172: uint16(sym_offset_date_time),
	1173: uint16(sym_local_time),
	1174: uint16(86),
	1175: uint16(4),
	1176: uint16(sym__basic_string),
	1177: uint16(sym__multiline_basic_string),
	1178: uint16(sym__literal_string),
	1179: uint16(sym__multiline_literal_string),
	1180: uint16(68),
	1181: uint16(6),
	1182: uint16(sym__inline_value),
	1183: uint16(sym_string),
	1184: uint16(sym_integer),
	1185: uint16(sym_float),
	1186: uint16(sym_array),
	1187: uint16(sym_inline_table),
	1188: uint16(17),
	1189: uint16(3),
	1190: uint16(1),
	1191: uint16(sym_comment),
	1192: uint16(21),
	1193: uint16(1),
	1194: uint16(anon_sym_LBRACK),
	1195: uint16(25),
	1196: uint16(1),
	1197: uint16(anon_sym_DQUOTE),
	1198: uint16(27),
	1199: uint16(1),
	1200: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1201: uint16(29),
	1202: uint16(1),
	1203: uint16(anon_sym_SQUOTE),
	1204: uint16(31),
	1205: uint16(1),
	1206: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	1207: uint16(33),
	1208: uint16(1),
	1209: uint16(aux_sym_integer_token1),
	1210: uint16(43),
	1211: uint16(1),
	1212: uint16(anon_sym_LBRACE),
	1213: uint16(97),
	1214: uint16(1),
	1215: uint16(aux_sym_document_token1),
	1216: uint16(99),
	1217: uint16(1),
	1218: uint16(anon_sym_RBRACK),
	1219: uint16(21),
	1220: uint16(1),
	1221: uint16(aux_sym_array_repeat1),
	1222: uint16(37),
	1223: uint16(2),
	1224: uint16(aux_sym_float_token1),
	1225: uint16(aux_sym_float_token2),
	1226: uint16(103),
	1227: uint16(2),
	1228: uint16(sym_local_date_time),
	1229: uint16(sym_local_date),
	1230: uint16(35),
	1231: uint16(3),
	1232: uint16(aux_sym_integer_token2),
	1233: uint16(aux_sym_integer_token3),
	1234: uint16(aux_sym_integer_token4),
	1235: uint16(101),
	1236: uint16(3),
	1237: uint16(sym_boolean),
	1238: uint16(sym_offset_date_time),
	1239: uint16(sym_local_time),
	1240: uint16(86),
	1241: uint16(4),
	1242: uint16(sym__basic_string),
	1243: uint16(sym__multiline_basic_string),
	1244: uint16(sym__literal_string),
	1245: uint16(sym__multiline_literal_string),
	1246: uint16(61),
	1247: uint16(6),
	1248: uint16(sym__inline_value),
	1249: uint16(sym_string),
	1250: uint16(sym_integer),
	1251: uint16(sym_float),
	1252: uint16(sym_array),
	1253: uint16(sym_inline_table),
	1254: uint16(17),
	1255: uint16(3),
	1256: uint16(1),
	1257: uint16(sym_comment),
	1258: uint16(19),
	1259: uint16(1),
	1260: uint16(aux_sym_document_token1),
	1261: uint16(21),
	1262: uint16(1),
	1263: uint16(anon_sym_LBRACK),
	1264: uint16(25),
	1265: uint16(1),
	1266: uint16(anon_sym_DQUOTE),
	1267: uint16(27),
	1268: uint16(1),
	1269: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1270: uint16(29),
	1271: uint16(1),
	1272: uint16(anon_sym_SQUOTE),
	1273: uint16(31),
	1274: uint16(1),
	1275: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	1276: uint16(33),
	1277: uint16(1),
	1278: uint16(aux_sym_integer_token1),
	1279: uint16(43),
	1280: uint16(1),
	1281: uint16(anon_sym_LBRACE),
	1282: uint16(105),
	1283: uint16(1),
	1284: uint16(anon_sym_RBRACK),
	1285: uint16(26),
	1286: uint16(1),
	1287: uint16(aux_sym_array_repeat1),
	1288: uint16(37),
	1289: uint16(2),
	1290: uint16(aux_sym_float_token1),
	1291: uint16(aux_sym_float_token2),
	1292: uint16(109),
	1293: uint16(2),
	1294: uint16(sym_local_date_time),
	1295: uint16(sym_local_date),
	1296: uint16(35),
	1297: uint16(3),
	1298: uint16(aux_sym_integer_token2),
	1299: uint16(aux_sym_integer_token3),
	1300: uint16(aux_sym_integer_token4),
	1301: uint16(107),
	1302: uint16(3),
	1303: uint16(sym_boolean),
	1304: uint16(sym_offset_date_time),
	1305: uint16(sym_local_time),
	1306: uint16(86),
	1307: uint16(4),
	1308: uint16(sym__basic_string),
	1309: uint16(sym__multiline_basic_string),
	1310: uint16(sym__literal_string),
	1311: uint16(sym__multiline_literal_string),
	1312: uint16(53),
	1313: uint16(6),
	1314: uint16(sym__inline_value),
	1315: uint16(sym_string),
	1316: uint16(sym_integer),
	1317: uint16(sym_float),
	1318: uint16(sym_array),
	1319: uint16(sym_inline_table),
	1320: uint16(16),
	1321: uint16(3),
	1322: uint16(1),
	1323: uint16(sym_comment),
	1324: uint16(21),
	1325: uint16(1),
	1326: uint16(anon_sym_LBRACK),
	1327: uint16(25),
	1328: uint16(1),
	1329: uint16(anon_sym_DQUOTE),
	1330: uint16(27),
	1331: uint16(1),
	1332: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1333: uint16(29),
	1334: uint16(1),
	1335: uint16(anon_sym_SQUOTE),
	1336: uint16(31),
	1337: uint16(1),
	1338: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	1339: uint16(33),
	1340: uint16(1),
	1341: uint16(aux_sym_integer_token1),
	1342: uint16(43),
	1343: uint16(1),
	1344: uint16(anon_sym_LBRACE),
	1345: uint16(111),
	1346: uint16(1),
	1347: uint16(aux_sym_document_token1),
	1348: uint16(23),
	1349: uint16(1),
	1350: uint16(aux_sym_array_repeat1),
	1351: uint16(37),
	1352: uint16(2),
	1353: uint16(aux_sym_float_token1),
	1354: uint16(aux_sym_float_token2),
	1355: uint16(57),
	1356: uint16(2),
	1357: uint16(sym_local_date_time),
	1358: uint16(sym_local_date),
	1359: uint16(35),
	1360: uint16(3),
	1361: uint16(aux_sym_integer_token2),
	1362: uint16(aux_sym_integer_token3),
	1363: uint16(aux_sym_integer_token4),
	1364: uint16(55),
	1365: uint16(3),
	1366: uint16(sym_boolean),
	1367: uint16(sym_offset_date_time),
	1368: uint16(sym_local_time),
	1369: uint16(86),
	1370: uint16(4),
	1371: uint16(sym__basic_string),
	1372: uint16(sym__multiline_basic_string),
	1373: uint16(sym__literal_string),
	1374: uint16(sym__multiline_literal_string),
	1375: uint16(68),
	1376: uint16(6),
	1377: uint16(sym__inline_value),
	1378: uint16(sym_string),
	1379: uint16(sym_integer),
	1380: uint16(sym_float),
	1381: uint16(sym_array),
	1382: uint16(sym_inline_table),
	1383: uint16(16),
	1384: uint16(3),
	1385: uint16(1),
	1386: uint16(sym_comment),
	1387: uint16(19),
	1388: uint16(1),
	1389: uint16(aux_sym_document_token1),
	1390: uint16(21),
	1391: uint16(1),
	1392: uint16(anon_sym_LBRACK),
	1393: uint16(25),
	1394: uint16(1),
	1395: uint16(anon_sym_DQUOTE),
	1396: uint16(27),
	1397: uint16(1),
	1398: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1399: uint16(29),
	1400: uint16(1),
	1401: uint16(anon_sym_SQUOTE),
	1402: uint16(31),
	1403: uint16(1),
	1404: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	1405: uint16(33),
	1406: uint16(1),
	1407: uint16(aux_sym_integer_token1),
	1408: uint16(43),
	1409: uint16(1),
	1410: uint16(anon_sym_LBRACE),
	1411: uint16(26),
	1412: uint16(1),
	1413: uint16(aux_sym_array_repeat1),
	1414: uint16(37),
	1415: uint16(2),
	1416: uint16(aux_sym_float_token1),
	1417: uint16(aux_sym_float_token2),
	1418: uint16(49),
	1419: uint16(2),
	1420: uint16(sym_local_date_time),
	1421: uint16(sym_local_date),
	1422: uint16(35),
	1423: uint16(3),
	1424: uint16(aux_sym_integer_token2),
	1425: uint16(aux_sym_integer_token3),
	1426: uint16(aux_sym_integer_token4),
	1427: uint16(47),
	1428: uint16(3),
	1429: uint16(sym_boolean),
	1430: uint16(sym_offset_date_time),
	1431: uint16(sym_local_time),
	1432: uint16(86),
	1433: uint16(4),
	1434: uint16(sym__basic_string),
	1435: uint16(sym__multiline_basic_string),
	1436: uint16(sym__literal_string),
	1437: uint16(sym__multiline_literal_string),
	1438: uint16(66),
	1439: uint16(6),
	1440: uint16(sym__inline_value),
	1441: uint16(sym_string),
	1442: uint16(sym_integer),
	1443: uint16(sym_float),
	1444: uint16(sym_array),
	1445: uint16(sym_inline_table),
	1446: uint16(14),
	1447: uint16(3),
	1448: uint16(1),
	1449: uint16(sym_comment),
	1450: uint16(21),
	1451: uint16(1),
	1452: uint16(anon_sym_LBRACK),
	1453: uint16(25),
	1454: uint16(1),
	1455: uint16(anon_sym_DQUOTE),
	1456: uint16(27),
	1457: uint16(1),
	1458: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1459: uint16(29),
	1460: uint16(1),
	1461: uint16(anon_sym_SQUOTE),
	1462: uint16(31),
	1463: uint16(1),
	1464: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	1465: uint16(33),
	1466: uint16(1),
	1467: uint16(aux_sym_integer_token1),
	1468: uint16(43),
	1469: uint16(1),
	1470: uint16(anon_sym_LBRACE),
	1471: uint16(37),
	1472: uint16(2),
	1473: uint16(aux_sym_float_token1),
	1474: uint16(aux_sym_float_token2),
	1475: uint16(115),
	1476: uint16(2),
	1477: uint16(sym_local_date_time),
	1478: uint16(sym_local_date),
	1479: uint16(35),
	1480: uint16(3),
	1481: uint16(aux_sym_integer_token2),
	1482: uint16(aux_sym_integer_token3),
	1483: uint16(aux_sym_integer_token4),
	1484: uint16(113),
	1485: uint16(3),
	1486: uint16(sym_boolean),
	1487: uint16(sym_offset_date_time),
	1488: uint16(sym_local_time),
	1489: uint16(86),
	1490: uint16(4),
	1491: uint16(sym__basic_string),
	1492: uint16(sym__multiline_basic_string),
	1493: uint16(sym__literal_string),
	1494: uint16(sym__multiline_literal_string),
	1495: uint16(110),
	1496: uint16(6),
	1497: uint16(sym__inline_value),
	1498: uint16(sym_string),
	1499: uint16(sym_integer),
	1500: uint16(sym_float),
	1501: uint16(sym_array),
	1502: uint16(sym_inline_table),
	1503: uint16(14),
	1504: uint16(3),
	1505: uint16(1),
	1506: uint16(sym_comment),
	1507: uint16(117),
	1508: uint16(1),
	1509: uint16(anon_sym_LBRACK),
	1510: uint16(119),
	1511: uint16(1),
	1512: uint16(anon_sym_DQUOTE),
	1513: uint16(121),
	1514: uint16(1),
	1515: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1516: uint16(123),
	1517: uint16(1),
	1518: uint16(anon_sym_SQUOTE),
	1519: uint16(125),
	1520: uint16(1),
	1521: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	1522: uint16(127),
	1523: uint16(1),
	1524: uint16(aux_sym_integer_token1),
	1525: uint16(137),
	1526: uint16(1),
	1527: uint16(anon_sym_LBRACE),
	1528: uint16(131),
	1529: uint16(2),
	1530: uint16(aux_sym_float_token1),
	1531: uint16(aux_sym_float_token2),
	1532: uint16(135),
	1533: uint16(2),
	1534: uint16(sym_local_date_time),
	1535: uint16(sym_local_date),
	1536: uint16(129),
	1537: uint16(3),
	1538: uint16(aux_sym_integer_token2),
	1539: uint16(aux_sym_integer_token3),
	1540: uint16(aux_sym_integer_token4),
	1541: uint16(133),
	1542: uint16(3),
	1543: uint16(sym_boolean),
	1544: uint16(sym_offset_date_time),
	1545: uint16(sym_local_time),
	1546: uint16(126),
	1547: uint16(4),
	1548: uint16(sym__basic_string),
	1549: uint16(sym__multiline_basic_string),
	1550: uint16(sym__literal_string),
	1551: uint16(sym__multiline_literal_string),
	1552: uint16(127),
	1553: uint16(6),
	1554: uint16(sym__inline_value),
	1555: uint16(sym_string),
	1556: uint16(sym_integer),
	1557: uint16(sym_float),
	1558: uint16(sym_array),
	1559: uint16(sym_inline_table),
	1560: uint16(5),
	1561: uint16(3),
	1562: uint16(1),
	1563: uint16(sym_comment),
	1564: uint16(139),
	1565: uint16(1),
	1566: uint16(aux_sym_document_token1),
	1567: uint16(26),
	1568: uint16(1),
	1569: uint16(aux_sym_array_repeat1),
	1570: uint16(144),
	1571: uint16(5),
	1572: uint16(anon_sym_DQUOTE),
	1573: uint16(anon_sym_SQUOTE),
	1574: uint16(aux_sym_integer_token1),
	1575: uint16(sym_local_date_time),
	1576: uint16(sym_local_date),
	1577: uint16(142),
	1578: uint16(14),
	1579: uint16(anon_sym_LBRACK),
	1580: uint16(anon_sym_RBRACK),
	1581: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1582: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	1583: uint16(aux_sym_integer_token2),
	1584: uint16(aux_sym_integer_token3),
	1585: uint16(aux_sym_integer_token4),
	1586: uint16(aux_sym_float_token1),
	1587: uint16(aux_sym_float_token2),
	1588: uint16(sym_boolean),
	1589: uint16(sym_offset_date_time),
	1590: uint16(sym_local_time),
	1591: uint16(anon_sym_COMMA),
	1592: uint16(anon_sym_LBRACE),
	1593: uint16(13),
	1594: uint16(3),
	1595: uint16(1),
	1596: uint16(sym_comment),
	1597: uint16(9),
	1598: uint16(1),
	1599: uint16(anon_sym_LBRACK),
	1600: uint16(11),
	1601: uint16(1),
	1602: uint16(anon_sym_LBRACK_LBRACK),
	1603: uint16(13),
	1604: uint16(1),
	1605: uint16(sym_bare_key),
	1606: uint16(15),
	1607: uint16(1),
	1608: uint16(anon_sym_DQUOTE),
	1609: uint16(17),
	1610: uint16(1),
	1611: uint16(anon_sym_SQUOTE),
	1612: uint16(146),
	1613: uint16(1),
	1615: uint16(148),
	1616: uint16(1),
	1617: uint16(aux_sym_document_token1),
	1618: uint16(144),
	1619: uint16(1),
	1620: uint16(sym__inline_pair),
	1621: uint16(32),
	1622: uint16(2),
	1623: uint16(sym_pair),
	1624: uint16(aux_sym_document_repeat1),
	1625: uint16(100),
	1626: uint16(2),
	1627: uint16(sym__basic_string),
	1628: uint16(sym__literal_string),
	1629: uint16(48),
	1630: uint16(3),
	1631: uint16(sym_table),
	1632: uint16(sym_table_array_element),
	1633: uint16(aux_sym_document_repeat2),
	1634: uint16(121),
	1635: uint16(3),
	1636: uint16(sym__key),
	1637: uint16(sym_dotted_key),
	1638: uint16(sym_quoted_key),
	1639: uint16(11),
	1640: uint16(3),
	1641: uint16(1),
	1642: uint16(sym_comment),
	1643: uint16(13),
	1644: uint16(1),
	1645: uint16(sym_bare_key),
	1646: uint16(15),
	1647: uint16(1),
	1648: uint16(anon_sym_DQUOTE),
	1649: uint16(17),
	1650: uint16(1),
	1651: uint16(anon_sym_SQUOTE),
	1652: uint16(152),
	1653: uint16(1),
	1654: uint16(aux_sym_document_token1),
	1655: uint16(154),
	1656: uint16(1),
	1657: uint16(anon_sym_LBRACK),
	1658: uint16(144),
	1659: uint16(1),
	1660: uint16(sym__inline_pair),
	1661: uint16(150),
	1662: uint16(2),
	1664: uint16(anon_sym_LBRACK_LBRACK),
	1665: uint16(31),
	1666: uint16(2),
	1667: uint16(sym_pair),
	1668: uint16(aux_sym_document_repeat1),
	1669: uint16(100),
	1670: uint16(2),
	1671: uint16(sym__basic_string),
	1672: uint16(sym__literal_string),
	1673: uint16(121),
	1674: uint16(3),
	1675: uint16(sym__key),
	1676: uint16(sym_dotted_key),
	1677: uint16(sym_quoted_key),
	1678: uint16(11),
	1679: uint16(3),
	1680: uint16(1),
	1681: uint16(sym_comment),
	1682: uint16(13),
	1683: uint16(1),
	1684: uint16(sym_bare_key),
	1685: uint16(15),
	1686: uint16(1),
	1687: uint16(anon_sym_DQUOTE),
	1688: uint16(17),
	1689: uint16(1),
	1690: uint16(anon_sym_SQUOTE),
	1691: uint16(158),
	1692: uint16(1),
	1693: uint16(aux_sym_document_token1),
	1694: uint16(160),
	1695: uint16(1),
	1696: uint16(anon_sym_LBRACK),
	1697: uint16(144),
	1698: uint16(1),
	1699: uint16(sym__inline_pair),
	1700: uint16(156),
	1701: uint16(2),
	1703: uint16(anon_sym_LBRACK_LBRACK),
	1704: uint16(30),
	1705: uint16(2),
	1706: uint16(sym_pair),
	1707: uint16(aux_sym_document_repeat1),
	1708: uint16(100),
	1709: uint16(2),
	1710: uint16(sym__basic_string),
	1711: uint16(sym__literal_string),
	1712: uint16(121),
	1713: uint16(3),
	1714: uint16(sym__key),
	1715: uint16(sym_dotted_key),
	1716: uint16(sym_quoted_key),
	1717: uint16(11),
	1718: uint16(3),
	1719: uint16(1),
	1720: uint16(sym_comment),
	1721: uint16(13),
	1722: uint16(1),
	1723: uint16(sym_bare_key),
	1724: uint16(15),
	1725: uint16(1),
	1726: uint16(anon_sym_DQUOTE),
	1727: uint16(17),
	1728: uint16(1),
	1729: uint16(anon_sym_SQUOTE),
	1730: uint16(148),
	1731: uint16(1),
	1732: uint16(aux_sym_document_token1),
	1733: uint16(164),
	1734: uint16(1),
	1735: uint16(anon_sym_LBRACK),
	1736: uint16(144),
	1737: uint16(1),
	1738: uint16(sym__inline_pair),
	1739: uint16(162),
	1740: uint16(2),
	1742: uint16(anon_sym_LBRACK_LBRACK),
	1743: uint16(32),
	1744: uint16(2),
	1745: uint16(sym_pair),
	1746: uint16(aux_sym_document_repeat1),
	1747: uint16(100),
	1748: uint16(2),
	1749: uint16(sym__basic_string),
	1750: uint16(sym__literal_string),
	1751: uint16(121),
	1752: uint16(3),
	1753: uint16(sym__key),
	1754: uint16(sym_dotted_key),
	1755: uint16(sym_quoted_key),
	1756: uint16(11),
	1757: uint16(3),
	1758: uint16(1),
	1759: uint16(sym_comment),
	1760: uint16(13),
	1761: uint16(1),
	1762: uint16(sym_bare_key),
	1763: uint16(15),
	1764: uint16(1),
	1765: uint16(anon_sym_DQUOTE),
	1766: uint16(17),
	1767: uint16(1),
	1768: uint16(anon_sym_SQUOTE),
	1769: uint16(148),
	1770: uint16(1),
	1771: uint16(aux_sym_document_token1),
	1772: uint16(168),
	1773: uint16(1),
	1774: uint16(anon_sym_LBRACK),
	1775: uint16(144),
	1776: uint16(1),
	1777: uint16(sym__inline_pair),
	1778: uint16(166),
	1779: uint16(2),
	1781: uint16(anon_sym_LBRACK_LBRACK),
	1782: uint16(32),
	1783: uint16(2),
	1784: uint16(sym_pair),
	1785: uint16(aux_sym_document_repeat1),
	1786: uint16(100),
	1787: uint16(2),
	1788: uint16(sym__basic_string),
	1789: uint16(sym__literal_string),
	1790: uint16(121),
	1791: uint16(3),
	1792: uint16(sym__key),
	1793: uint16(sym_dotted_key),
	1794: uint16(sym_quoted_key),
	1795: uint16(11),
	1796: uint16(3),
	1797: uint16(1),
	1798: uint16(sym_comment),
	1799: uint16(172),
	1800: uint16(1),
	1801: uint16(aux_sym_document_token1),
	1802: uint16(175),
	1803: uint16(1),
	1804: uint16(anon_sym_LBRACK),
	1805: uint16(177),
	1806: uint16(1),
	1807: uint16(sym_bare_key),
	1808: uint16(180),
	1809: uint16(1),
	1810: uint16(anon_sym_DQUOTE),
	1811: uint16(183),
	1812: uint16(1),
	1813: uint16(anon_sym_SQUOTE),
	1814: uint16(144),
	1815: uint16(1),
	1816: uint16(sym__inline_pair),
	1817: uint16(170),
	1818: uint16(2),
	1820: uint16(anon_sym_LBRACK_LBRACK),
	1821: uint16(32),
	1822: uint16(2),
	1823: uint16(sym_pair),
	1824: uint16(aux_sym_document_repeat1),
	1825: uint16(100),
	1826: uint16(2),
	1827: uint16(sym__basic_string),
	1828: uint16(sym__literal_string),
	1829: uint16(121),
	1830: uint16(3),
	1831: uint16(sym__key),
	1832: uint16(sym_dotted_key),
	1833: uint16(sym_quoted_key),
	1834: uint16(8),
	1835: uint16(3),
	1836: uint16(1),
	1837: uint16(sym_comment),
	1838: uint16(15),
	1839: uint16(1),
	1840: uint16(anon_sym_DQUOTE),
	1841: uint16(17),
	1842: uint16(1),
	1843: uint16(anon_sym_SQUOTE),
	1844: uint16(186),
	1845: uint16(1),
	1846: uint16(sym_bare_key),
	1847: uint16(188),
	1848: uint16(1),
	1849: uint16(anon_sym_RBRACE),
	1850: uint16(94),
	1851: uint16(1),
	1852: uint16(sym__inline_pair),
	1853: uint16(100),
	1854: uint16(2),
	1855: uint16(sym__basic_string),
	1856: uint16(sym__literal_string),
	1857: uint16(112),
	1858: uint16(3),
	1859: uint16(sym__key),
	1860: uint16(sym_dotted_key),
	1861: uint16(sym_quoted_key),
	1862: uint16(8),
	1863: uint16(3),
	1864: uint16(1),
	1865: uint16(sym_comment),
	1866: uint16(15),
	1867: uint16(1),
	1868: uint16(anon_sym_DQUOTE),
	1869: uint16(17),
	1870: uint16(1),
	1871: uint16(anon_sym_SQUOTE),
	1872: uint16(186),
	1873: uint16(1),
	1874: uint16(sym_bare_key),
	1875: uint16(190),
	1876: uint16(1),
	1877: uint16(anon_sym_RBRACE),
	1878: uint16(99),
	1879: uint16(1),
	1880: uint16(sym__inline_pair),
	1881: uint16(100),
	1882: uint16(2),
	1883: uint16(sym__basic_string),
	1884: uint16(sym__literal_string),
	1885: uint16(112),
	1886: uint16(3),
	1887: uint16(sym__key),
	1888: uint16(sym_dotted_key),
	1889: uint16(sym_quoted_key),
	1890: uint16(7),
	1891: uint16(3),
	1892: uint16(1),
	1893: uint16(sym_comment),
	1894: uint16(15),
	1895: uint16(1),
	1896: uint16(anon_sym_DQUOTE),
	1897: uint16(17),
	1898: uint16(1),
	1899: uint16(anon_sym_SQUOTE),
	1900: uint16(186),
	1901: uint16(1),
	1902: uint16(sym_bare_key),
	1903: uint16(114),
	1904: uint16(1),
	1905: uint16(sym__inline_pair),
	1906: uint16(100),
	1907: uint16(2),
	1908: uint16(sym__basic_string),
	1909: uint16(sym__literal_string),
	1910: uint16(112),
	1911: uint16(3),
	1912: uint16(sym__key),
	1913: uint16(sym_dotted_key),
	1914: uint16(sym_quoted_key),
	1915: uint16(6),
	1916: uint16(3),
	1917: uint16(1),
	1918: uint16(sym_comment),
	1919: uint16(15),
	1920: uint16(1),
	1921: uint16(anon_sym_DQUOTE),
	1922: uint16(17),
	1923: uint16(1),
	1924: uint16(anon_sym_SQUOTE),
	1925: uint16(192),
	1926: uint16(1),
	1927: uint16(sym_bare_key),
	1928: uint16(100),
	1929: uint16(2),
	1930: uint16(sym__basic_string),
	1931: uint16(sym__literal_string),
	1932: uint16(109),
	1933: uint16(3),
	1934: uint16(sym__key),
	1935: uint16(sym_dotted_key),
	1936: uint16(sym_quoted_key),
	1937: uint16(6),
	1938: uint16(3),
	1939: uint16(1),
	1940: uint16(sym_comment),
	1941: uint16(194),
	1942: uint16(1),
	1943: uint16(sym_bare_key),
	1944: uint16(196),
	1945: uint16(1),
	1946: uint16(anon_sym_DQUOTE),
	1947: uint16(198),
	1948: uint16(1),
	1949: uint16(anon_sym_SQUOTE),
	1950: uint16(122),
	1951: uint16(2),
	1952: uint16(sym__basic_string),
	1953: uint16(sym__literal_string),
	1954: uint16(108),
	1955: uint16(3),
	1956: uint16(sym__key),
	1957: uint16(sym_dotted_key),
	1958: uint16(sym_quoted_key),
	1959: uint16(4),
	1960: uint16(200),
	1961: uint16(1),
	1962: uint16(sym_comment),
	1963: uint16(204),
	1964: uint16(1),
	1965: uint16(sym__multiline_basic_string_end),
	1966: uint16(39),
	1967: uint16(1),
	1968: uint16(aux_sym__multiline_basic_string_repeat1),
	1969: uint16(202),
	1970: uint16(5),
	1971: uint16(sym__multiline_basic_string_content),
	1972: uint16(aux_sym__basic_string_token1),
	1973: uint16(aux_sym__multiline_basic_string_token1),
	1974: uint16(sym_escape_sequence),
	1975: uint16(sym__escape_line_ending),
	1976: uint16(4),
	1977: uint16(200),
	1978: uint16(1),
	1979: uint16(sym_comment),
	1980: uint16(209),
	1981: uint16(1),
	1982: uint16(sym__multiline_basic_string_end),
	1983: uint16(39),
	1984: uint16(1),
	1985: uint16(aux_sym__multiline_basic_string_repeat1),
	1986: uint16(206),
	1987: uint16(5),
	1988: uint16(sym__multiline_basic_string_content),
	1989: uint16(aux_sym__basic_string_token1),
	1990: uint16(aux_sym__multiline_basic_string_token1),
	1991: uint16(sym_escape_sequence),
	1992: uint16(sym__escape_line_ending),
	1993: uint16(4),
	1994: uint16(200),
	1995: uint16(1),
	1996: uint16(sym_comment),
	1997: uint16(213),
	1998: uint16(1),
	1999: uint16(sym__multiline_basic_string_end),
	2000: uint16(38),
	2001: uint16(1),
	2002: uint16(aux_sym__multiline_basic_string_repeat1),
	2003: uint16(211),
	2004: uint16(5),
	2005: uint16(sym__multiline_basic_string_content),
	2006: uint16(aux_sym__basic_string_token1),
	2007: uint16(aux_sym__multiline_basic_string_token1),
	2008: uint16(sym_escape_sequence),
	2009: uint16(sym__escape_line_ending),
	2010: uint16(4),
	2011: uint16(200),
	2012: uint16(1),
	2013: uint16(sym_comment),
	2014: uint16(217),
	2015: uint16(1),
	2016: uint16(sym__multiline_basic_string_end),
	2017: uint16(42),
	2018: uint16(1),
	2019: uint16(aux_sym__multiline_basic_string_repeat1),
	2020: uint16(215),
	2021: uint16(5),
	2022: uint16(sym__multiline_basic_string_content),
	2023: uint16(aux_sym__basic_string_token1),
	2024: uint16(aux_sym__multiline_basic_string_token1),
	2025: uint16(sym_escape_sequence),
	2026: uint16(sym__escape_line_ending),
	2027: uint16(4),
	2028: uint16(200),
	2029: uint16(1),
	2030: uint16(sym_comment),
	2031: uint16(219),
	2032: uint16(1),
	2033: uint16(sym__multiline_basic_string_end),
	2034: uint16(39),
	2035: uint16(1),
	2036: uint16(aux_sym__multiline_basic_string_repeat1),
	2037: uint16(202),
	2038: uint16(5),
	2039: uint16(sym__multiline_basic_string_content),
	2040: uint16(aux_sym__basic_string_token1),
	2041: uint16(aux_sym__multiline_basic_string_token1),
	2042: uint16(sym_escape_sequence),
	2043: uint16(sym__escape_line_ending),
	2044: uint16(6),
	2045: uint16(3),
	2046: uint16(1),
	2047: uint16(sym_comment),
	2048: uint16(196),
	2049: uint16(1),
	2050: uint16(anon_sym_DQUOTE),
	2051: uint16(198),
	2052: uint16(1),
	2053: uint16(anon_sym_SQUOTE),
	2054: uint16(221),
	2055: uint16(1),
	2056: uint16(sym_bare_key),
	2057: uint16(111),
	2058: uint16(2),
	2059: uint16(sym__key),
	2060: uint16(sym_quoted_key),
	2061: uint16(122),
	2062: uint16(2),
	2063: uint16(sym__basic_string),
	2064: uint16(sym__literal_string),
	2065: uint16(6),
	2066: uint16(3),
	2067: uint16(1),
	2068: uint16(sym_comment),
	2069: uint16(15),
	2070: uint16(1),
	2071: uint16(anon_sym_DQUOTE),
	2072: uint16(17),
	2073: uint16(1),
	2074: uint16(anon_sym_SQUOTE),
	2075: uint16(223),
	2076: uint16(1),
	2077: uint16(sym_bare_key),
	2078: uint16(100),
	2079: uint16(2),
	2080: uint16(sym__basic_string),
	2081: uint16(sym__literal_string),
	2082: uint16(105),
	2083: uint16(2),
	2084: uint16(sym__key),
	2085: uint16(sym_quoted_key),
	2086: uint16(3),
	2087: uint16(3),
	2088: uint16(1),
	2089: uint16(sym_comment),
	2090: uint16(227),
	2091: uint16(1),
	2092: uint16(anon_sym_LBRACK),
	2093: uint16(225),
	2094: uint16(6),
	2096: uint16(aux_sym_document_token1),
	2097: uint16(anon_sym_LBRACK_LBRACK),
	2098: uint16(sym_bare_key),
	2099: uint16(anon_sym_DQUOTE),
	2100: uint16(anon_sym_SQUOTE),
	2101: uint16(2),
	2102: uint16(3),
	2103: uint16(1),
	2104: uint16(sym_comment),
	2105: uint16(229),
	2106: uint16(6),
	2107: uint16(aux_sym_document_token1),
	2108: uint16(anon_sym_RBRACK),
	2109: uint16(anon_sym_EQ),
	2110: uint16(anon_sym_DOT),
	2111: uint16(anon_sym_COMMA),
	2112: uint16(anon_sym_RBRACE),
	2113: uint16(5),
	2114: uint16(3),
	2115: uint16(1),
	2116: uint16(sym_comment),
	2117: uint16(231),
	2118: uint16(1),
	2120: uint16(233),
	2121: uint16(1),
	2122: uint16(anon_sym_LBRACK),
	2123: uint16(236),
	2124: uint16(1),
	2125: uint16(anon_sym_LBRACK_LBRACK),
	2126: uint16(47),
	2127: uint16(3),
	2128: uint16(sym_table),
	2129: uint16(sym_table_array_element),
	2130: uint16(aux_sym_document_repeat2),
	2131: uint16(5),
	2132: uint16(3),
	2133: uint16(1),
	2134: uint16(sym_comment),
	2135: uint16(9),
	2136: uint16(1),
	2137: uint16(anon_sym_LBRACK),
	2138: uint16(11),
	2139: uint16(1),
	2140: uint16(anon_sym_LBRACK_LBRACK),
	2141: uint16(239),
	2142: uint16(1),
	2144: uint16(47),
	2145: uint16(3),
	2146: uint16(sym_table),
	2147: uint16(sym_table_array_element),
	2148: uint16(aux_sym_document_repeat2),
	2149: uint16(2),
	2150: uint16(3),
	2151: uint16(1),
	2152: uint16(sym_comment),
	2153: uint16(241),
	2154: uint16(6),
	2155: uint16(aux_sym_document_token1),
	2156: uint16(anon_sym_RBRACK),
	2157: uint16(anon_sym_EQ),
	2158: uint16(anon_sym_DOT),
	2159: uint16(anon_sym_COMMA),
	2160: uint16(anon_sym_RBRACE),
	2161: uint16(2),
	2162: uint16(3),
	2163: uint16(1),
	2164: uint16(sym_comment),
	2165: uint16(243),
	2166: uint16(6),
	2167: uint16(aux_sym_document_token1),
	2168: uint16(anon_sym_RBRACK),
	2169: uint16(anon_sym_EQ),
	2170: uint16(anon_sym_DOT),
	2171: uint16(anon_sym_COMMA),
	2172: uint16(anon_sym_RBRACE),
	2173: uint16(2),
	2174: uint16(3),
	2175: uint16(1),
	2176: uint16(sym_comment),
	2177: uint16(245),
	2178: uint16(6),
	2179: uint16(aux_sym_document_token1),
	2180: uint16(anon_sym_RBRACK),
	2181: uint16(anon_sym_EQ),
	2182: uint16(anon_sym_DOT),
	2183: uint16(anon_sym_COMMA),
	2184: uint16(anon_sym_RBRACE),
	2185: uint16(5),
	2186: uint16(3),
	2187: uint16(1),
	2188: uint16(sym_comment),
	2189: uint16(9),
	2190: uint16(1),
	2191: uint16(anon_sym_LBRACK),
	2192: uint16(11),
	2193: uint16(1),
	2194: uint16(anon_sym_LBRACK_LBRACK),
	2195: uint16(146),
	2196: uint16(1),
	2198: uint16(47),
	2199: uint16(3),
	2200: uint16(sym_table),
	2201: uint16(sym_table_array_element),
	2202: uint16(aux_sym_document_repeat2),
	2203: uint16(6),
	2204: uint16(3),
	2205: uint16(1),
	2206: uint16(sym_comment),
	2207: uint16(95),
	2208: uint16(1),
	2209: uint16(anon_sym_RBRACK),
	2210: uint16(247),
	2211: uint16(1),
	2212: uint16(aux_sym_document_token1),
	2213: uint16(249),
	2214: uint16(1),
	2215: uint16(anon_sym_COMMA),
	2216: uint16(58),
	2217: uint16(1),
	2218: uint16(aux_sym_array_repeat1),
	2219: uint16(101),
	2220: uint16(1),
	2221: uint16(aux_sym_array_repeat2),
	2222: uint16(6),
	2223: uint16(3),
	2224: uint16(1),
	2225: uint16(sym_comment),
	2226: uint16(23),
	2227: uint16(1),
	2228: uint16(anon_sym_RBRACK),
	2229: uint16(251),
	2230: uint16(1),
	2231: uint16(aux_sym_document_token1),
	2232: uint16(253),
	2233: uint16(1),
	2234: uint16(anon_sym_COMMA),
	2235: uint16(57),
	2236: uint16(1),
	2237: uint16(aux_sym_array_repeat1),
	2238: uint16(96),
	2239: uint16(1),
	2240: uint16(aux_sym_array_repeat2),
	2241: uint16(6),
	2242: uint16(3),
	2243: uint16(1),
	2244: uint16(sym_comment),
	2245: uint16(73),
	2246: uint16(1),
	2247: uint16(anon_sym_RBRACK),
	2248: uint16(255),
	2249: uint16(1),
	2250: uint16(aux_sym_document_token1),
	2251: uint16(257),
	2252: uint16(1),
	2253: uint16(anon_sym_COMMA),
	2254: uint16(65),
	2255: uint16(1),
	2256: uint16(aux_sym_array_repeat1),
	2257: uint16(106),
	2258: uint16(1),
	2259: uint16(aux_sym_array_repeat2),
	2260: uint16(4),
	2261: uint16(200),
	2262: uint16(1),
	2263: uint16(sym_comment),
	2264: uint16(261),
	2265: uint16(1),
	2266: uint16(sym__multiline_literal_string_end),
	2267: uint16(60),
	2268: uint16(1),
	2269: uint16(aux_sym__multiline_literal_string_repeat1),
	2270: uint16(259),
	2271: uint16(3),
	2272: uint16(sym__multiline_literal_string_content),
	2273: uint16(aux_sym__multiline_basic_string_token1),
	2274: uint16(aux_sym__literal_string_token1),
	2275: uint16(6),
	2276: uint16(3),
	2277: uint16(1),
	2278: uint16(sym_comment),
	2279: uint16(19),
	2280: uint16(1),
	2281: uint16(aux_sym_document_token1),
	2282: uint16(73),
	2283: uint16(1),
	2284: uint16(anon_sym_RBRACK),
	2285: uint16(257),
	2286: uint16(1),
	2287: uint16(anon_sym_COMMA),
	2288: uint16(26),
	2289: uint16(1),
	2290: uint16(aux_sym_array_repeat1),
	2291: uint16(106),
	2292: uint16(1),
	2293: uint16(aux_sym_array_repeat2),
	2294: uint16(6),
	2295: uint16(3),
	2296: uint16(1),
	2297: uint16(sym_comment),
	2298: uint16(19),
	2299: uint16(1),
	2300: uint16(aux_sym_document_token1),
	2301: uint16(77),
	2302: uint16(1),
	2303: uint16(anon_sym_RBRACK),
	2304: uint16(263),
	2305: uint16(1),
	2306: uint16(anon_sym_COMMA),
	2307: uint16(26),
	2308: uint16(1),
	2309: uint16(aux_sym_array_repeat1),
	2310: uint16(98),
	2311: uint16(1),
	2312: uint16(aux_sym_array_repeat2),
	2313: uint16(4),
	2314: uint16(200),
	2315: uint16(1),
	2316: uint16(sym_comment),
	2317: uint16(267),
	2318: uint16(1),
	2319: uint16(sym__multiline_literal_string_end),
	2320: uint16(63),
	2321: uint16(1),
	2322: uint16(aux_sym__multiline_literal_string_repeat1),
	2323: uint16(265),
	2324: uint16(3),
	2325: uint16(sym__multiline_literal_string_content),
	2326: uint16(aux_sym__multiline_basic_string_token1),
	2327: uint16(aux_sym__literal_string_token1),
	2328: uint16(4),
	2329: uint16(200),
	2330: uint16(1),
	2331: uint16(sym_comment),
	2332: uint16(272),
	2333: uint16(1),
	2334: uint16(sym__multiline_literal_string_end),
	2335: uint16(60),
	2336: uint16(1),
	2337: uint16(aux_sym__multiline_literal_string_repeat1),
	2338: uint16(269),
	2339: uint16(3),
	2340: uint16(sym__multiline_literal_string_content),
	2341: uint16(aux_sym__multiline_basic_string_token1),
	2342: uint16(aux_sym__literal_string_token1),
	2343: uint16(6),
	2344: uint16(3),
	2345: uint16(1),
	2346: uint16(sym_comment),
	2347: uint16(105),
	2348: uint16(1),
	2349: uint16(anon_sym_RBRACK),
	2350: uint16(274),
	2351: uint16(1),
	2352: uint16(aux_sym_document_token1),
	2353: uint16(276),
	2354: uint16(1),
	2355: uint16(anon_sym_COMMA),
	2356: uint16(64),
	2357: uint16(1),
	2358: uint16(aux_sym_array_repeat1),
	2359: uint16(103),
	2360: uint16(1),
	2361: uint16(aux_sym_array_repeat2),
	2362: uint16(4),
	2363: uint16(200),
	2364: uint16(1),
	2365: uint16(sym_comment),
	2366: uint16(280),
	2367: uint16(1),
	2368: uint16(sym__multiline_literal_string_end),
	2369: uint16(56),
	2370: uint16(1),
	2371: uint16(aux_sym__multiline_literal_string_repeat1),
	2372: uint16(278),
	2373: uint16(3),
	2374: uint16(sym__multiline_literal_string_content),
	2375: uint16(aux_sym__multiline_basic_string_token1),
	2376: uint16(aux_sym__literal_string_token1),
	2377: uint16(4),
	2378: uint16(200),
	2379: uint16(1),
	2380: uint16(sym_comment),
	2381: uint16(282),
	2382: uint16(1),
	2383: uint16(sym__multiline_literal_string_end),
	2384: uint16(60),
	2385: uint16(1),
	2386: uint16(aux_sym__multiline_literal_string_repeat1),
	2387: uint16(259),
	2388: uint16(3),
	2389: uint16(sym__multiline_literal_string_content),
	2390: uint16(aux_sym__multiline_basic_string_token1),
	2391: uint16(aux_sym__literal_string_token1),
	2392: uint16(6),
	2393: uint16(3),
	2394: uint16(1),
	2395: uint16(sym_comment),
	2396: uint16(19),
	2397: uint16(1),
	2398: uint16(aux_sym_document_token1),
	2399: uint16(95),
	2400: uint16(1),
	2401: uint16(anon_sym_RBRACK),
	2402: uint16(249),
	2403: uint16(1),
	2404: uint16(anon_sym_COMMA),
	2405: uint16(26),
	2406: uint16(1),
	2407: uint16(aux_sym_array_repeat1),
	2408: uint16(101),
	2409: uint16(1),
	2410: uint16(aux_sym_array_repeat2),
	2411: uint16(6),
	2412: uint16(3),
	2413: uint16(1),
	2414: uint16(sym_comment),
	2415: uint16(19),
	2416: uint16(1),
	2417: uint16(aux_sym_document_token1),
	2418: uint16(69),
	2419: uint16(1),
	2420: uint16(anon_sym_RBRACK),
	2421: uint16(284),
	2422: uint16(1),
	2423: uint16(anon_sym_COMMA),
	2424: uint16(26),
	2425: uint16(1),
	2426: uint16(aux_sym_array_repeat1),
	2427: uint16(95),
	2428: uint16(1),
	2429: uint16(aux_sym_array_repeat2),
	2430: uint16(4),
	2431: uint16(3),
	2432: uint16(1),
	2433: uint16(sym_comment),
	2434: uint16(286),
	2435: uint16(1),
	2436: uint16(aux_sym_document_token1),
	2437: uint16(81),
	2438: uint16(1),
	2439: uint16(aux_sym_array_repeat1),
	2440: uint16(288),
	2441: uint16(2),
	2442: uint16(anon_sym_RBRACK),
	2443: uint16(anon_sym_COMMA),
	2444: uint16(2),
	2445: uint16(3),
	2446: uint16(1),
	2447: uint16(sym_comment),
	2448: uint16(290),
	2449: uint16(4),
	2450: uint16(aux_sym_document_token1),
	2451: uint16(anon_sym_RBRACK),
	2452: uint16(anon_sym_COMMA),
	2453: uint16(anon_sym_RBRACE),
	2454: uint16(4),
	2455: uint16(3),
	2456: uint16(1),
	2457: uint16(sym_comment),
	2458: uint16(292),
	2459: uint16(1),
	2460: uint16(aux_sym_document_token1),
	2461: uint16(74),
	2462: uint16(1),
	2463: uint16(aux_sym_array_repeat1),
	2464: uint16(294),
	2465: uint16(2),
	2466: uint16(anon_sym_RBRACK),
	2467: uint16(anon_sym_COMMA),
	2468: uint16(4),
	2469: uint16(200),
	2470: uint16(1),
	2471: uint16(sym_comment),
	2472: uint16(298),
	2473: uint16(1),
	2474: uint16(anon_sym_DQUOTE2),
	2475: uint16(72),
	2476: uint16(1),
	2477: uint16(aux_sym__basic_string_repeat1),
	2478: uint16(296),
	2479: uint16(2),
	2480: uint16(aux_sym__basic_string_token1),
	2481: uint16(sym_escape_sequence),
	2482: uint16(4),
	2483: uint16(200),
	2484: uint16(1),
	2485: uint16(sym_comment),
	2486: uint16(302),
	2487: uint16(1),
	2488: uint16(anon_sym_DQUOTE2),
	2489: uint16(69),
	2490: uint16(1),
	2491: uint16(aux_sym__basic_string_repeat1),
	2492: uint16(300),
	2493: uint16(2),
	2494: uint16(aux_sym__basic_string_token1),
	2495: uint16(sym_escape_sequence),
	2496: uint16(4),
	2497: uint16(200),
	2498: uint16(1),
	2499: uint16(sym_comment),
	2500: uint16(304),
	2501: uint16(1),
	2502: uint16(anon_sym_DQUOTE2),
	2503: uint16(72),
	2504: uint16(1),
	2505: uint16(aux_sym__basic_string_repeat1),
	2506: uint16(296),
	2507: uint16(2),
	2508: uint16(aux_sym__basic_string_token1),
	2509: uint16(sym_escape_sequence),
	2510: uint16(4),
	2511: uint16(200),
	2512: uint16(1),
	2513: uint16(sym_comment),
	2514: uint16(309),
	2515: uint16(1),
	2516: uint16(anon_sym_DQUOTE2),
	2517: uint16(72),
	2518: uint16(1),
	2519: uint16(aux_sym__basic_string_repeat1),
	2520: uint16(306),
	2521: uint16(2),
	2522: uint16(aux_sym__basic_string_token1),
	2523: uint16(sym_escape_sequence),
	2524: uint16(4),
	2525: uint16(200),
	2526: uint16(1),
	2527: uint16(sym_comment),
	2528: uint16(311),
	2529: uint16(1),
	2530: uint16(anon_sym_DQUOTE2),
	2531: uint16(72),
	2532: uint16(1),
	2533: uint16(aux_sym__basic_string_repeat1),
	2534: uint16(296),
	2535: uint16(2),
	2536: uint16(aux_sym__basic_string_token1),
	2537: uint16(sym_escape_sequence),
	2538: uint16(4),
	2539: uint16(3),
	2540: uint16(1),
	2541: uint16(sym_comment),
	2542: uint16(19),
	2543: uint16(1),
	2544: uint16(aux_sym_document_token1),
	2545: uint16(26),
	2546: uint16(1),
	2547: uint16(aux_sym_array_repeat1),
	2548: uint16(288),
	2549: uint16(2),
	2550: uint16(anon_sym_RBRACK),
	2551: uint16(anon_sym_COMMA),
	2552: uint16(4),
	2553: uint16(200),
	2554: uint16(1),
	2555: uint16(sym_comment),
	2556: uint16(315),
	2557: uint16(1),
	2558: uint16(anon_sym_DQUOTE2),
	2559: uint16(73),
	2560: uint16(1),
	2561: uint16(aux_sym__basic_string_repeat1),
	2562: uint16(313),
	2563: uint16(2),
	2564: uint16(aux_sym__basic_string_token1),
	2565: uint16(sym_escape_sequence),
	2566: uint16(4),
	2567: uint16(200),
	2568: uint16(1),
	2569: uint16(sym_comment),
	2570: uint16(319),
	2571: uint16(1),
	2572: uint16(anon_sym_DQUOTE2),
	2573: uint16(71),
	2574: uint16(1),
	2575: uint16(aux_sym__basic_string_repeat1),
	2576: uint16(317),
	2577: uint16(2),
	2578: uint16(aux_sym__basic_string_token1),
	2579: uint16(sym_escape_sequence),
	2580: uint16(2),
	2581: uint16(3),
	2582: uint16(1),
	2583: uint16(sym_comment),
	2584: uint16(321),
	2585: uint16(4),
	2586: uint16(aux_sym_document_token1),
	2587: uint16(anon_sym_RBRACK),
	2588: uint16(anon_sym_COMMA),
	2589: uint16(anon_sym_RBRACE),
	2590: uint16(2),
	2591: uint16(3),
	2592: uint16(1),
	2593: uint16(sym_comment),
	2594: uint16(323),
	2595: uint16(4),
	2596: uint16(aux_sym_document_token1),
	2597: uint16(anon_sym_RBRACK),
	2598: uint16(anon_sym_COMMA),
	2599: uint16(anon_sym_RBRACE),
	2600: uint16(2),
	2601: uint16(3),
	2602: uint16(1),
	2603: uint16(sym_comment),
	2604: uint16(325),
	2605: uint16(4),
	2606: uint16(aux_sym_document_token1),
	2607: uint16(anon_sym_RBRACK),
	2608: uint16(anon_sym_COMMA),
	2609: uint16(anon_sym_RBRACE),
	2610: uint16(2),
	2611: uint16(3),
	2612: uint16(1),
	2613: uint16(sym_comment),
	2614: uint16(327),
	2615: uint16(4),
	2616: uint16(aux_sym_document_token1),
	2617: uint16(anon_sym_RBRACK),
	2618: uint16(anon_sym_COMMA),
	2619: uint16(anon_sym_RBRACE),
	2620: uint16(4),
	2621: uint16(3),
	2622: uint16(1),
	2623: uint16(sym_comment),
	2624: uint16(19),
	2625: uint16(1),
	2626: uint16(aux_sym_document_token1),
	2627: uint16(26),
	2628: uint16(1),
	2629: uint16(aux_sym_array_repeat1),
	2630: uint16(329),
	2631: uint16(2),
	2632: uint16(anon_sym_RBRACK),
	2633: uint16(anon_sym_COMMA),
	2634: uint16(2),
	2635: uint16(3),
	2636: uint16(1),
	2637: uint16(sym_comment),
	2638: uint16(331),
	2639: uint16(4),
	2640: uint16(aux_sym_document_token1),
	2641: uint16(anon_sym_RBRACK),
	2642: uint16(anon_sym_COMMA),
	2643: uint16(anon_sym_RBRACE),
	2644: uint16(2),
	2645: uint16(3),
	2646: uint16(1),
	2647: uint16(sym_comment),
	2648: uint16(333),
	2649: uint16(4),
	2650: uint16(aux_sym_document_token1),
	2651: uint16(anon_sym_RBRACK),
	2652: uint16(anon_sym_COMMA),
	2653: uint16(anon_sym_RBRACE),
	2654: uint16(2),
	2655: uint16(3),
	2656: uint16(1),
	2657: uint16(sym_comment),
	2658: uint16(335),
	2659: uint16(4),
	2660: uint16(aux_sym_document_token1),
	2661: uint16(anon_sym_RBRACK),
	2662: uint16(anon_sym_COMMA),
	2663: uint16(anon_sym_RBRACE),
	2664: uint16(2),
	2665: uint16(3),
	2666: uint16(1),
	2667: uint16(sym_comment),
	2668: uint16(337),
	2669: uint16(4),
	2670: uint16(aux_sym_document_token1),
	2671: uint16(anon_sym_RBRACK),
	2672: uint16(anon_sym_COMMA),
	2673: uint16(anon_sym_RBRACE),
	2674: uint16(2),
	2675: uint16(3),
	2676: uint16(1),
	2677: uint16(sym_comment),
	2678: uint16(339),
	2679: uint16(4),
	2680: uint16(aux_sym_document_token1),
	2681: uint16(anon_sym_RBRACK),
	2682: uint16(anon_sym_COMMA),
	2683: uint16(anon_sym_RBRACE),
	2684: uint16(2),
	2685: uint16(3),
	2686: uint16(1),
	2687: uint16(sym_comment),
	2688: uint16(341),
	2689: uint16(4),
	2690: uint16(aux_sym_document_token1),
	2691: uint16(anon_sym_RBRACK),
	2692: uint16(anon_sym_COMMA),
	2693: uint16(anon_sym_RBRACE),
	2694: uint16(2),
	2695: uint16(3),
	2696: uint16(1),
	2697: uint16(sym_comment),
	2698: uint16(343),
	2699: uint16(4),
	2700: uint16(aux_sym_document_token1),
	2701: uint16(anon_sym_RBRACK),
	2702: uint16(anon_sym_COMMA),
	2703: uint16(anon_sym_RBRACE),
	2704: uint16(2),
	2705: uint16(3),
	2706: uint16(1),
	2707: uint16(sym_comment),
	2708: uint16(345),
	2709: uint16(4),
	2710: uint16(aux_sym_document_token1),
	2711: uint16(anon_sym_RBRACK),
	2712: uint16(anon_sym_COMMA),
	2713: uint16(anon_sym_RBRACE),
	2714: uint16(2),
	2715: uint16(3),
	2716: uint16(1),
	2717: uint16(sym_comment),
	2718: uint16(347),
	2719: uint16(4),
	2720: uint16(aux_sym_document_token1),
	2721: uint16(anon_sym_RBRACK),
	2722: uint16(anon_sym_COMMA),
	2723: uint16(anon_sym_RBRACE),
	2724: uint16(2),
	2725: uint16(3),
	2726: uint16(1),
	2727: uint16(sym_comment),
	2728: uint16(349),
	2729: uint16(4),
	2730: uint16(aux_sym_document_token1),
	2731: uint16(anon_sym_RBRACK),
	2732: uint16(anon_sym_COMMA),
	2733: uint16(anon_sym_RBRACE),
	2734: uint16(2),
	2735: uint16(3),
	2736: uint16(1),
	2737: uint16(sym_comment),
	2738: uint16(351),
	2739: uint16(4),
	2740: uint16(aux_sym_document_token1),
	2741: uint16(anon_sym_RBRACK),
	2742: uint16(anon_sym_COMMA),
	2743: uint16(anon_sym_RBRACE),
	2744: uint16(2),
	2745: uint16(3),
	2746: uint16(1),
	2747: uint16(sym_comment),
	2748: uint16(353),
	2749: uint16(4),
	2750: uint16(aux_sym_document_token1),
	2751: uint16(anon_sym_RBRACK),
	2752: uint16(anon_sym_COMMA),
	2753: uint16(anon_sym_RBRACE),
	2754: uint16(4),
	2755: uint16(3),
	2756: uint16(1),
	2757: uint16(sym_comment),
	2758: uint16(355),
	2759: uint16(1),
	2760: uint16(anon_sym_COMMA),
	2761: uint16(357),
	2762: uint16(1),
	2763: uint16(anon_sym_RBRACE),
	2764: uint16(102),
	2765: uint16(1),
	2766: uint16(aux_sym_inline_table_repeat1),
	2767: uint16(4),
	2768: uint16(3),
	2769: uint16(1),
	2770: uint16(sym_comment),
	2771: uint16(53),
	2772: uint16(1),
	2773: uint16(anon_sym_RBRACK),
	2774: uint16(359),
	2775: uint16(1),
	2776: uint16(anon_sym_COMMA),
	2777: uint16(97),
	2778: uint16(1),
	2779: uint16(aux_sym_array_repeat2),
	2780: uint16(4),
	2781: uint16(3),
	2782: uint16(1),
	2783: uint16(sym_comment),
	2784: uint16(73),
	2785: uint16(1),
	2786: uint16(anon_sym_RBRACK),
	2787: uint16(257),
	2788: uint16(1),
	2789: uint16(anon_sym_COMMA),
	2790: uint16(97),
	2791: uint16(1),
	2792: uint16(aux_sym_array_repeat2),
	2793: uint16(4),
	2794: uint16(3),
	2795: uint16(1),
	2796: uint16(sym_comment),
	2797: uint16(294),
	2798: uint16(1),
	2799: uint16(anon_sym_RBRACK),
	2800: uint16(361),
	2801: uint16(1),
	2802: uint16(anon_sym_COMMA),
	2803: uint16(97),
	2804: uint16(1),
	2805: uint16(aux_sym_array_repeat2),
	2806: uint16(4),
	2807: uint16(3),
	2808: uint16(1),
	2809: uint16(sym_comment),
	2810: uint16(65),
	2811: uint16(1),
	2812: uint16(anon_sym_RBRACK),
	2813: uint16(364),
	2814: uint16(1),
	2815: uint16(anon_sym_COMMA),
	2816: uint16(97),
	2817: uint16(1),
	2818: uint16(aux_sym_array_repeat2),
	2819: uint16(4),
	2820: uint16(3),
	2821: uint16(1),
	2822: uint16(sym_comment),
	2823: uint16(355),
	2824: uint16(1),
	2825: uint16(anon_sym_COMMA),
	2826: uint16(366),
	2827: uint16(1),
	2828: uint16(anon_sym_RBRACE),
	2829: uint16(107),
	2830: uint16(1),
	2831: uint16(aux_sym_inline_table_repeat1),
	2832: uint16(2),
	2833: uint16(3),
	2834: uint16(1),
	2835: uint16(sym_comment),
	2836: uint16(368),
	2837: uint16(3),
	2838: uint16(anon_sym_RBRACK),
	2839: uint16(anon_sym_EQ),
	2840: uint16(anon_sym_DOT),
	2841: uint16(4),
	2842: uint16(3),
	2843: uint16(1),
	2844: uint16(sym_comment),
	2845: uint16(77),
	2846: uint16(1),
	2847: uint16(anon_sym_RBRACK),
	2848: uint16(263),
	2849: uint16(1),
	2850: uint16(anon_sym_COMMA),
	2851: uint16(97),
	2852: uint16(1),
	2853: uint16(aux_sym_array_repeat2),
	2854: uint16(4),
	2855: uint16(3),
	2856: uint16(1),
	2857: uint16(sym_comment),
	2858: uint16(355),
	2859: uint16(1),
	2860: uint16(anon_sym_COMMA),
	2861: uint16(370),
	2862: uint16(1),
	2863: uint16(anon_sym_RBRACE),
	2864: uint16(104),
	2865: uint16(1),
	2866: uint16(aux_sym_inline_table_repeat1),
	2867: uint16(4),
	2868: uint16(3),
	2869: uint16(1),
	2870: uint16(sym_comment),
	2871: uint16(95),
	2872: uint16(1),
	2873: uint16(anon_sym_RBRACK),
	2874: uint16(249),
	2875: uint16(1),
	2876: uint16(anon_sym_COMMA),
	2877: uint16(97),
	2878: uint16(1),
	2879: uint16(aux_sym_array_repeat2),
	2880: uint16(4),
	2881: uint16(3),
	2882: uint16(1),
	2883: uint16(sym_comment),
	2884: uint16(372),
	2885: uint16(1),
	2886: uint16(anon_sym_COMMA),
	2887: uint16(375),
	2888: uint16(1),
	2889: uint16(anon_sym_RBRACE),
	2890: uint16(104),
	2891: uint16(1),
	2892: uint16(aux_sym_inline_table_repeat1),
	2893: uint16(2),
	2894: uint16(3),
	2895: uint16(1),
	2896: uint16(sym_comment),
	2897: uint16(377),
	2898: uint16(3),
	2899: uint16(anon_sym_RBRACK),
	2900: uint16(anon_sym_EQ),
	2901: uint16(anon_sym_DOT),
	2902: uint16(4),
	2903: uint16(3),
	2904: uint16(1),
	2905: uint16(sym_comment),
	2906: uint16(69),
	2907: uint16(1),
	2908: uint16(anon_sym_RBRACK),
	2909: uint16(284),
	2910: uint16(1),
	2911: uint16(anon_sym_COMMA),
	2912: uint16(97),
	2913: uint16(1),
	2914: uint16(aux_sym_array_repeat2),
	2915: uint16(4),
	2916: uint16(3),
	2917: uint16(1),
	2918: uint16(sym_comment),
	2919: uint16(355),
	2920: uint16(1),
	2921: uint16(anon_sym_COMMA),
	2922: uint16(379),
	2923: uint16(1),
	2924: uint16(anon_sym_RBRACE),
	2925: uint16(104),
	2926: uint16(1),
	2927: uint16(aux_sym_inline_table_repeat1),
	2928: uint16(3),
	2929: uint16(3),
	2930: uint16(1),
	2931: uint16(sym_comment),
	2932: uint16(381),
	2933: uint16(1),
	2934: uint16(anon_sym_RBRACK_RBRACK),
	2935: uint16(383),
	2936: uint16(1),
	2937: uint16(anon_sym_DOT),
	2938: uint16(3),
	2939: uint16(3),
	2940: uint16(1),
	2941: uint16(sym_comment),
	2942: uint16(385),
	2943: uint16(1),
	2944: uint16(anon_sym_RBRACK),
	2945: uint16(387),
	2946: uint16(1),
	2947: uint16(anon_sym_DOT),
	2948: uint16(2),
	2949: uint16(3),
	2950: uint16(1),
	2951: uint16(sym_comment),
	2952: uint16(389),
	2953: uint16(2),
	2954: uint16(anon_sym_COMMA),
	2955: uint16(anon_sym_RBRACE),
	2956: uint16(2),
	2957: uint16(3),
	2958: uint16(1),
	2959: uint16(sym_comment),
	2960: uint16(377),
	2961: uint16(2),
	2962: uint16(anon_sym_RBRACK_RBRACK),
	2963: uint16(anon_sym_DOT),
	2964: uint16(3),
	2965: uint16(3),
	2966: uint16(1),
	2967: uint16(sym_comment),
	2968: uint16(387),
	2969: uint16(1),
	2970: uint16(anon_sym_DOT),
	2971: uint16(391),
	2972: uint16(1),
	2973: uint16(anon_sym_EQ),
	2974: uint16(3),
	2975: uint16(200),
	2976: uint16(1),
	2977: uint16(sym_comment),
	2978: uint16(393),
	2979: uint16(1),
	2980: uint16(aux_sym__literal_string_token1),
	2981: uint16(395),
	2982: uint16(1),
	2983: uint16(anon_sym_SQUOTE2),
	2984: uint16(2),
	2985: uint16(3),
	2986: uint16(1),
	2987: uint16(sym_comment),
	2988: uint16(397),
	2989: uint16(2),
	2990: uint16(anon_sym_COMMA),
	2991: uint16(anon_sym_RBRACE),
	2992: uint16(2),
	2993: uint16(3),
	2994: uint16(1),
	2995: uint16(sym_comment),
	2996: uint16(245),
	2997: uint16(2),
	2998: uint16(anon_sym_RBRACK_RBRACK),
	2999: uint16(anon_sym_DOT),
	3000: uint16(2),
	3001: uint16(3),
	3002: uint16(1),
	3003: uint16(sym_comment),
	3004: uint16(243),
	3005: uint16(2),
	3006: uint16(anon_sym_RBRACK_RBRACK),
	3007: uint16(anon_sym_DOT),
	3008: uint16(2),
	3009: uint16(3),
	3010: uint16(1),
	3011: uint16(sym_comment),
	3012: uint16(229),
	3013: uint16(2),
	3014: uint16(anon_sym_RBRACK_RBRACK),
	3015: uint16(anon_sym_DOT),
	3016: uint16(2),
	3017: uint16(3),
	3018: uint16(1),
	3019: uint16(sym_comment),
	3020: uint16(241),
	3021: uint16(2),
	3022: uint16(anon_sym_RBRACK_RBRACK),
	3023: uint16(anon_sym_DOT),
	3024: uint16(3),
	3025: uint16(200),
	3026: uint16(1),
	3027: uint16(sym_comment),
	3028: uint16(399),
	3029: uint16(1),
	3030: uint16(aux_sym__literal_string_token1),
	3031: uint16(401),
	3032: uint16(1),
	3033: uint16(anon_sym_SQUOTE2),
	3034: uint16(3),
	3035: uint16(200),
	3036: uint16(1),
	3037: uint16(sym_comment),
	3038: uint16(403),
	3039: uint16(1),
	3040: uint16(aux_sym__literal_string_token1),
	3041: uint16(405),
	3042: uint16(1),
	3043: uint16(anon_sym_SQUOTE2),
	3044: uint16(3),
	3045: uint16(3),
	3046: uint16(1),
	3047: uint16(sym_comment),
	3048: uint16(387),
	3049: uint16(1),
	3050: uint16(anon_sym_DOT),
	3051: uint16(407),
	3052: uint16(1),
	3053: uint16(anon_sym_EQ),
	3054: uint16(2),
	3055: uint16(3),
	3056: uint16(1),
	3057: uint16(sym_comment),
	3058: uint16(368),
	3059: uint16(2),
	3060: uint16(anon_sym_RBRACK_RBRACK),
	3061: uint16(anon_sym_DOT),
	3062: uint16(2),
	3063: uint16(3),
	3064: uint16(1),
	3065: uint16(sym_comment),
	3066: uint16(333),
	3067: uint16(1),
	3068: uint16(sym__line_ending_or_eof),
	3069: uint16(2),
	3070: uint16(3),
	3071: uint16(1),
	3072: uint16(sym_comment),
	3073: uint16(353),
	3074: uint16(1),
	3075: uint16(sym__line_ending_or_eof),
	3076: uint16(2),
	3077: uint16(3),
	3078: uint16(1),
	3079: uint16(sym_comment),
	3080: uint16(351),
	3081: uint16(1),
	3082: uint16(sym__line_ending_or_eof),
	3083: uint16(2),
	3084: uint16(3),
	3085: uint16(1),
	3086: uint16(sym_comment),
	3087: uint16(339),
	3088: uint16(1),
	3089: uint16(sym__line_ending_or_eof),
	3090: uint16(2),
	3091: uint16(3),
	3092: uint16(1),
	3093: uint16(sym_comment),
	3094: uint16(389),
	3095: uint16(1),
	3096: uint16(sym__line_ending_or_eof),
	3097: uint16(2),
	3098: uint16(3),
	3099: uint16(1),
	3100: uint16(sym_comment),
	3101: uint16(335),
	3102: uint16(1),
	3103: uint16(sym__line_ending_or_eof),
	3104: uint16(2),
	3105: uint16(3),
	3106: uint16(1),
	3107: uint16(sym_comment),
	3108: uint16(241),
	3109: uint16(1),
	3110: uint16(sym__line_ending_or_eof),
	3111: uint16(2),
	3112: uint16(3),
	3113: uint16(1),
	3114: uint16(sym_comment),
	3115: uint16(229),
	3116: uint16(1),
	3117: uint16(sym__line_ending_or_eof),
	3118: uint16(2),
	3119: uint16(3),
	3120: uint16(1),
	3121: uint16(sym_comment),
	3122: uint16(341),
	3123: uint16(1),
	3124: uint16(sym__line_ending_or_eof),
	3125: uint16(2),
	3126: uint16(3),
	3127: uint16(1),
	3128: uint16(sym_comment),
	3129: uint16(409),
	3130: uint16(1),
	3131: uint16(sym__line_ending_or_eof),
	3132: uint16(2),
	3133: uint16(3),
	3134: uint16(1),
	3135: uint16(sym_comment),
	3136: uint16(411),
	3137: uint16(1),
	3138: uint16(sym__line_ending_or_eof),
	3139: uint16(2),
	3140: uint16(3),
	3141: uint16(1),
	3142: uint16(sym_comment),
	3143: uint16(349),
	3144: uint16(1),
	3145: uint16(sym__line_ending_or_eof),
	3146: uint16(2),
	3147: uint16(3),
	3148: uint16(1),
	3149: uint16(sym_comment),
	3150: uint16(413),
	3151: uint16(1),
	3152: uint16(anon_sym_SQUOTE2),
	3153: uint16(2),
	3154: uint16(3),
	3155: uint16(1),
	3156: uint16(sym_comment),
	3157: uint16(243),
	3158: uint16(1),
	3159: uint16(sym__line_ending_or_eof),
	3160: uint16(2),
	3161: uint16(3),
	3162: uint16(1),
	3163: uint16(sym_comment),
	3164: uint16(331),
	3165: uint16(1),
	3166: uint16(sym__line_ending_or_eof),
	3167: uint16(2),
	3168: uint16(3),
	3169: uint16(1),
	3170: uint16(sym_comment),
	3171: uint16(343),
	3172: uint16(1),
	3173: uint16(sym__line_ending_or_eof),
	3174: uint16(2),
	3175: uint16(3),
	3176: uint16(1),
	3177: uint16(sym_comment),
	3178: uint16(245),
	3179: uint16(1),
	3180: uint16(sym__line_ending_or_eof),
	3181: uint16(2),
	3182: uint16(3),
	3183: uint16(1),
	3184: uint16(sym_comment),
	3185: uint16(325),
	3186: uint16(1),
	3187: uint16(sym__line_ending_or_eof),
	3188: uint16(2),
	3189: uint16(3),
	3190: uint16(1),
	3191: uint16(sym_comment),
	3192: uint16(345),
	3193: uint16(1),
	3194: uint16(sym__line_ending_or_eof),
	3195: uint16(2),
	3196: uint16(3),
	3197: uint16(1),
	3198: uint16(sym_comment),
	3199: uint16(347),
	3200: uint16(1),
	3201: uint16(sym__line_ending_or_eof),
	3202: uint16(2),
	3203: uint16(3),
	3204: uint16(1),
	3205: uint16(sym_comment),
	3206: uint16(321),
	3207: uint16(1),
	3208: uint16(sym__line_ending_or_eof),
	3209: uint16(2),
	3210: uint16(3),
	3211: uint16(1),
	3212: uint16(sym_comment),
	3213: uint16(415),
	3214: uint16(1),
	3215: uint16(sym__line_ending_or_eof),
	3216: uint16(2),
	3217: uint16(3),
	3218: uint16(1),
	3219: uint16(sym_comment),
	3220: uint16(417),
	3221: uint16(1),
	3223: uint16(2),
	3224: uint16(3),
	3225: uint16(1),
	3226: uint16(sym_comment),
	3227: uint16(419),
	3228: uint16(1),
	3229: uint16(anon_sym_SQUOTE2),
	3230: uint16(2),
	3231: uint16(3),
	3232: uint16(1),
	3233: uint16(sym_comment),
	3234: uint16(327),
	3235: uint16(1),
	3236: uint16(sym__line_ending_or_eof),
	3237: uint16(2),
	3238: uint16(3),
	3239: uint16(1),
	3240: uint16(sym_comment),
	3241: uint16(337),
	3242: uint16(1),
	3243: uint16(sym__line_ending_or_eof),
	3244: uint16(2),
	3245: uint16(3),
	3246: uint16(1),
	3247: uint16(sym_comment),
	3248: uint16(323),
	3249: uint16(1),
	3250: uint16(sym__line_ending_or_eof),
	3251: uint16(2),
	3252: uint16(3),
	3253: uint16(1),
	3254: uint16(sym_comment),
	3255: uint16(421),
	3256: uint16(1),
	3257: uint16(anon_sym_SQUOTE2),
	3258: uint16(2),
	3259: uint16(3),
	3260: uint16(1),
	3261: uint16(sym_comment),
	3262: uint16(290),
	3263: uint16(1),
	3264: uint16(sym__line_ending_or_eof),
}

var ts_small_parse_table_map = [150]uint32_t{
	1:   uint32(66),
	2:   uint32(132),
	3:   uint32(198),
	4:   uint32(264),
	5:   uint32(330),
	6:   uint32(396),
	7:   uint32(462),
	8:   uint32(528),
	9:   uint32(594),
	10:  uint32(660),
	11:  uint32(726),
	12:  uint32(792),
	13:  uint32(858),
	14:  uint32(924),
	15:  uint32(990),
	16:  uint32(1056),
	17:  uint32(1122),
	18:  uint32(1188),
	19:  uint32(1254),
	20:  uint32(1320),
	21:  uint32(1383),
	22:  uint32(1446),
	23:  uint32(1503),
	24:  uint32(1560),
	25:  uint32(1593),
	26:  uint32(1639),
	27:  uint32(1678),
	28:  uint32(1717),
	29:  uint32(1756),
	30:  uint32(1795),
	31:  uint32(1834),
	32:  uint32(1862),
	33:  uint32(1890),
	34:  uint32(1915),
	35:  uint32(1937),
	36:  uint32(1959),
	37:  uint32(1976),
	38:  uint32(1993),
	39:  uint32(2010),
	40:  uint32(2027),
	41:  uint32(2044),
	42:  uint32(2065),
	43:  uint32(2086),
	44:  uint32(2101),
	45:  uint32(2113),
	46:  uint32(2131),
	47:  uint32(2149),
	48:  uint32(2161),
	49:  uint32(2173),
	50:  uint32(2185),
	51:  uint32(2203),
	52:  uint32(2222),
	53:  uint32(2241),
	54:  uint32(2260),
	55:  uint32(2275),
	56:  uint32(2294),
	57:  uint32(2313),
	58:  uint32(2328),
	59:  uint32(2343),
	60:  uint32(2362),
	61:  uint32(2377),
	62:  uint32(2392),
	63:  uint32(2411),
	64:  uint32(2430),
	65:  uint32(2444),
	66:  uint32(2454),
	67:  uint32(2468),
	68:  uint32(2482),
	69:  uint32(2496),
	70:  uint32(2510),
	71:  uint32(2524),
	72:  uint32(2538),
	73:  uint32(2552),
	74:  uint32(2566),
	75:  uint32(2580),
	76:  uint32(2590),
	77:  uint32(2600),
	78:  uint32(2610),
	79:  uint32(2620),
	80:  uint32(2634),
	81:  uint32(2644),
	82:  uint32(2654),
	83:  uint32(2664),
	84:  uint32(2674),
	85:  uint32(2684),
	86:  uint32(2694),
	87:  uint32(2704),
	88:  uint32(2714),
	89:  uint32(2724),
	90:  uint32(2734),
	91:  uint32(2744),
	92:  uint32(2754),
	93:  uint32(2767),
	94:  uint32(2780),
	95:  uint32(2793),
	96:  uint32(2806),
	97:  uint32(2819),
	98:  uint32(2832),
	99:  uint32(2841),
	100: uint32(2854),
	101: uint32(2867),
	102: uint32(2880),
	103: uint32(2893),
	104: uint32(2902),
	105: uint32(2915),
	106: uint32(2928),
	107: uint32(2938),
	108: uint32(2948),
	109: uint32(2956),
	110: uint32(2964),
	111: uint32(2974),
	112: uint32(2984),
	113: uint32(2992),
	114: uint32(3000),
	115: uint32(3008),
	116: uint32(3016),
	117: uint32(3024),
	118: uint32(3034),
	119: uint32(3044),
	120: uint32(3054),
	121: uint32(3062),
	122: uint32(3069),
	123: uint32(3076),
	124: uint32(3083),
	125: uint32(3090),
	126: uint32(3097),
	127: uint32(3104),
	128: uint32(3111),
	129: uint32(3118),
	130: uint32(3125),
	131: uint32(3132),
	132: uint32(3139),
	133: uint32(3146),
	134: uint32(3153),
	135: uint32(3160),
	136: uint32(3167),
	137: uint32(3174),
	138: uint32(3181),
	139: uint32(3188),
	140: uint32(3195),
	141: uint32(3202),
	142: uint32(3209),
	143: uint32(3216),
	144: uint32(3223),
	145: uint32(3230),
	146: uint32(3237),
	147: uint32(3244),
	148: uint32(3251),
	149: uint32(3258),
}

var ts_parse_actions = [423]TSParseActionEntry{
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
		Fsymbol:     uint16(sym_document),
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
		Fstate: uint16(27),
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
		Fstate: uint16(36),
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
		Fstate: uint16(37),
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
		Fstate: uint16(121),
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
		Fstate: uint16(76),
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
		Fstate: uint16(120),
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
		Fstate: uint16(26),
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
		Fstate: uint16(20),
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
		Fstate: uint16(125),
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
		Fstate: uint16(76),
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
		Fstate: uint16(40),
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
		Fstate: uint16(120),
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
		Fstate: uint16(59),
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
		Fstate: uint16(83),
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
		Fstate: uint16(83),
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
		Fstate: uint16(84),
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
		Fstate: uint16(55),
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
		Fstate: uint16(55),
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
		Fstate: uint16(33),
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
		Fstate: uint16(143),
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
		Fstate: uint16(66),
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
		Fstate: uint16(66),
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
		Fstate: uint16(15),
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
		Fstate: uint16(140),
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
		Fstate: uint16(68),
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
		Fcount: uint8(1),
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
		Fstate: uint16(68),
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
		Fstate: uint16(14),
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
		Fstate: uint16(78),
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
		Fstate: uint16(6),
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
		Fstate: uint16(79),
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
		Fstate: uint16(10),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(18),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(131),
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
		Fstate: uint16(8),
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
		Fstate: uint16(80),
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
		Fstate: uint16(77),
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
		Fstate: uint16(149),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(138),
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
		Fstate: uint16(54),
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
		Fcount: uint8(1),
	}})),
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
		Fstate: uint16(54),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(3),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(87),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(21),
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
		Fstate: uint16(88),
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
		Fstate: uint16(61),
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
		Fstate: uint16(61),
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
		Fstate: uint16(92),
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
		Fstate: uint16(53),
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
		Fstate: uint16(53),
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
		Fstate: uint16(23),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
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
		Fstate: uint16(110),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
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
		Fstate: uint16(70),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(41),
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
		Fcount: uint8(1),
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
		Fstate: uint16(113),
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
		Fstate: uint16(62),
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
		Fcount: uint8(1),
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
		Fstate: uint16(123),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(128),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	134: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
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
		Fstate: uint16(127),
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
		Fstate: uint16(34),
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
		Fsymbol:      uint16(aux_sym_array_repeat1),
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
		Fstate:      uint16(26),
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
		Fsymbol:      uint16(aux_sym_array_repeat1),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(aux_sym_array_repeat1),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_document),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_table),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_table),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_table_array_element),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_table_array_element),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_table_array_element),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_table_array_element),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_table),
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
		Fcount: uint8(1),
	}})),
	169: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_table),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_document_repeat1),
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
		Fsymbol:      uint16(aux_sym_document_repeat1),
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
		Fstate:      uint16(32),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(aux_sym_document_repeat1),
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
		Fsymbol:      uint16(aux_sym_document_repeat1),
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
		Fstate:      uint16(121),
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
		Fsymbol:      uint16(aux_sym_document_repeat1),
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
		Fstate:      uint16(76),
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
		Fsymbol:      uint16(aux_sym_document_repeat1),
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
		Fstate:      uint16(120),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(91),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(134),
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
		Fstate: uint16(109),
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
		Fstate: uint16(108),
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
		Fstate: uint16(75),
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
		Fstate: uint16(119),
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
		Fcount: uint8(1),
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
		Fextra: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(39),
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
		Fstate: uint16(93),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__multiline_basic_string_repeat1),
	})))),
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
		Fstate:      uint16(39),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__multiline_basic_string_repeat1),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(38),
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
		Fstate: uint16(89),
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
		Fstate: uint16(42),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(111),
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
		Fstate: uint16(105),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_pair),
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
		Fsymbol:      uint16(sym_pair),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__basic_string),
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
		Fcount:    uint8(1),
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
		Fsymbol:      uint16(aux_sym_document_repeat2),
	})))),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_document_repeat2),
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
		Fstate:      uint16(36),
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
		Fsymbol:      uint16(aux_sym_document_repeat2),
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
		Fstate:      uint16(37),
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
		Fsymbol:      uint16(sym_document),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	242: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__literal_string),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__literal_string),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__basic_string),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(58),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(12),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(57),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(11),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(65),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(60),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(151),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(7),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(63),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(90),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym__multiline_literal_string_repeat1),
	})))),
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
		Fstate:      uint16(60),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__multiline_literal_string_repeat1),
	})))),
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
		Fstate: uint16(64),
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
		Fstate: uint16(19),
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
		Fstate: uint16(56),
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
		Fstate: uint16(142),
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
		Fstate: uint16(67),
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
		Fstate: uint16(4),
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
		Fstate: uint16(81),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(aux_sym_array_repeat2),
	})))),
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
		Fsymbol:      uint16(sym__multiline_literal_string),
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
		Fstate: uint16(74),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_array_repeat2),
	})))),
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
		Fstate: uint16(72),
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
		Fstate: uint16(130),
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
		Fstate: uint16(69),
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
		Fstate: uint16(139),
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
		Fstate: uint16(46),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	307: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__basic_string_repeat1),
	})))),
	308: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	310: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__basic_string_repeat1),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	314: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(71),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	320: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	322: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(8),
		Fsymbol:      uint16(sym_array),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_array),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	326: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_array),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	328: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	330: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(aux_sym_array_repeat2),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_inline_table),
		Fproduction_id: uint16(1),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_integer),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_float),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_inline_table),
		Fproduction_id: uint16(1),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_string),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	342: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__multiline_basic_string),
	})))),
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
		Fsymbol:      uint16(sym__multiline_literal_string),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(sym_inline_table),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_array),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	354: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__multiline_basic_string),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	358: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	362: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_array_repeat2),
	})))),
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
		Fstate:      uint16(22),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(5),
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
		Fstate: uint16(148),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	369: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_quoted_key),
	})))),
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
		Fstate: uint16(82),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	373: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_inline_table_repeat1),
	})))),
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
		Fstate:      uint16(35),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	376: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_inline_table_repeat1),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	378: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_dotted_key),
	})))),
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
		Fstate: uint16(137),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(43),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(133),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(44),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__inline_pair),
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
		Fstate: uint16(24),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(136),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_inline_table_repeat1),
		Fproduction_id: uint16(1),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(146),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(116),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(135),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(50),
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
		Fstate: uint16(25),
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
		Fstate: uint16(29),
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
		Fstate: uint16(28),
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
		Fstate: uint16(49),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	418: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f uint8_t
		_ [7]byte
	}{f: uint8(TSParseActionTypeAccept)})),
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
		Fstate: uint16(118),
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
		Fstate: uint16(129),
	}})))),
}

func tree_sitter_toml(tls *libc.TLS) (r uintptr) {
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
}

func init() {
	p := unsafe.Pointer(&language)
	*(*uintptr)(unsafe.Add(p, 92)) = __ccgo_fp(ts_lex)
	*(*uintptr)(unsafe.Add(p, 112)) = __ccgo_fp(tree_sitter_toml_external_scanner_create)
	*(*uintptr)(unsafe.Add(p, 116)) = __ccgo_fp(tree_sitter_toml_external_scanner_destroy)
	*(*uintptr)(unsafe.Add(p, 120)) = __ccgo_fp(tree_sitter_toml_external_scanner_scan)
	*(*uintptr)(unsafe.Add(p, 124)) = __ccgo_fp(tree_sitter_toml_external_scanner_serialize)
	*(*uintptr)(unsafe.Add(p, 128)) = __ccgo_fp(tree_sitter_toml_external_scanner_deserialize)
}

func __ccgo_fp(f interface{}) uintptr {
	type iface [2]uintptr
	return (*iface)(unsafe.Pointer(&f))[1]
}

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = "end\x00document_token1\x00comment\x00[\x00]\x00[[\x00]]\x00=\x00.\x00bare_key\x00\"\x00_basic_string_token1\x00\"\"\"\x00_multiline_basic_string_token1\x00escape_sequence\x00'\x00_literal_string_token1\x00'''\x00integer_token1\x00integer_token2\x00integer_token3\x00integer_token4\x00float_token1\x00float_token2\x00boolean\x00offset_date_time\x00local_date_time\x00local_date\x00local_time\x00,\x00{\x00}\x00_line_ending_or_eof\x00_multiline_basic_string_content\x00_multiline_basic_string_end\x00_multiline_literal_string_content\x00_multiline_literal_string_end\x00document\x00table\x00table_array_element\x00pair\x00_inline_pair\x00_key\x00dotted_key\x00quoted_key\x00_inline_value\x00string\x00_basic_string\x00_multiline_basic_string\x00_literal_string\x00_multiline_literal_string\x00integer\x00float\x00array\x00inline_table\x00document_repeat1\x00document_repeat2\x00_basic_string_repeat1\x00_multiline_basic_string_repeat1\x00_multiline_literal_string_repeat1\x00array_repeat1\x00array_repeat2\x00inline_table_repeat1\x00"
