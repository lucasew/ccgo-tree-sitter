// Code generated for linux/386 by 'ccgo -ignore-unsupported-alignment -ignore-unsupported-atomic-sizes -ignore-vector-functions -std=gnu11 -O0 -Dfunc=func_token -Dinterface=interface_token -Dselect=select_token -Dchan=chan_token -Dgo=go_token -Dmap=map_token -Dpackage=package_token -Dtype=type_token -Dvar=var_token -Dimport=import_token -Ddefer=defer_token -Dfallthrough=fallthrough_token -Drange=range_token -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-yuck/src -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-yuck -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party -I /home/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/include -I /home/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/src combined.c -o grammar.go', DO NOT EDIT.

//go:build linux && 386

package grammar_yuck

import (
	"reflect"
	"unsafe"

	"modernc.org/libc"
)

var _ reflect.Type
var _ unsafe.Pointer

const ALIAS_COUNT = 2
const EXIT_FAILURE = 1
const EXIT_SUCCESS = 0
const EXTERNAL_TOKEN_COUNT = 3
const FIELD_COUNT = 8
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
const LANGUAGE_VERSION = 15
const LARGE_STATE_COUNT = 2
const MAX_ALIAS_SEQUENCE_LENGTH = 7
const MAX_RESERVED_WORD_SET_SIZE = 0
const PRODUCTION_ID_COUNT = 7
const PTRDIFF_MAX = "INT32_MAX"
const PTRDIFF_MIN = "INT32_MIN"
const RAND_MAX = 0x7fffffff
const SIG_ATOMIC_MAX = "INT32_MAX"
const SIG_ATOMIC_MIN = "INT32_MIN"
const SIZE_MAX = "UINT32_MAX"
const STATE_COUNT = 127
const SUPERTYPE_COUNT = 2
const SYMBOL_COUNT = 80
const TOKEN_COUNT = 47
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

type wint_t = uint32

type wctype_t = uint32

type locale_t = uintptr

type wctrans_t = uintptr

type TokenType = int32

const UNESCAPED_SINGLE_QUOTE_STRING_FRAGMENT = 0
const UNESCAPED_DOUBLE_QUOTE_STRING_FRAGMENT = 1
const UNESCAPED_BACKTICK_STRING_FRAGMENT = 2

func tree_sitter_yuck_external_scanner_create(tls *libc.TLS) (r uintptr) {
	return libc.UintptrFromInt32(0)
}

func tree_sitter_yuck_external_scanner_destroy(tls *libc.TLS, payload uintptr) {
}

func tree_sitter_yuck_external_scanner_serialize(tls *libc.TLS, payload uintptr, buffer uintptr) (r uint32) {
	return uint32(0)
}

func tree_sitter_yuck_external_scanner_deserialize(tls *libc.TLS, payload uintptr, buffer uintptr, length uint32) {
}

func advance(tls *libc.TLS, lexer uintptr) {
	(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(false1 != 0))
}

func skip(tls *libc.TLS, lexer uintptr) {
	(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(true1 != 0))
}

func scan_string_literal_fragment(tls *libc.TLS, lexer uintptr, quote int32_t) (r uint8) {
	var has_content uint8
	var next int32_t
	_, _ = has_content, next
	has_content = libc.BoolUint8(false1 != 0)
	for {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		next = (*TSLexer)(unsafe.Pointer(lexer)).Flookahead
		if next == quote {
			return has_content
		} else {
			if next == int32('\000') {
				return libc.BoolUint8(false1 != 0)
			} else {
				if next == int32('$') {
					advance(tls, lexer)
					if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('{') {
						return has_content
					}
				} else {
					if next == int32('\\') {
						return has_content
					} else {
						advance(tls, lexer)
					}
				}
			}
		}
		goto _1
	_1:
		;
		has_content = libc.BoolUint8(true1 != 0)
	}
	return r
}

func tree_sitter_yuck_external_scanner_scan(tls *libc.TLS, payload uintptr, lexer uintptr, valid_symbols uintptr) (r uint8) {
	if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(UNESCAPED_DOUBLE_QUOTE_STRING_FRAGMENT))) != 0 {
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(UNESCAPED_DOUBLE_QUOTE_STRING_FRAGMENT)
		return scan_string_literal_fragment(tls, lexer, int32('"'))
	}
	if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(UNESCAPED_SINGLE_QUOTE_STRING_FRAGMENT))) != 0 {
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(UNESCAPED_SINGLE_QUOTE_STRING_FRAGMENT)
		return scan_string_literal_fragment(tls, lexer, int32('\''))
	}
	if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(UNESCAPED_BACKTICK_STRING_FRAGMENT))) != 0 {
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(UNESCAPED_BACKTICK_STRING_FRAGMENT)
		return scan_string_literal_fragment(tls, lexer, int32('`'))
	}
	return libc.BoolUint8(false1 != 0)
}

/* Automatically @generated by tree-sitter */

type ts_symbol_identifiers = int32

const sym_symbol = 1
const anon_sym_LPAREN = 2
const anon_sym_for = 3
const anon_sym_in = 4
const anon_sym_RPAREN = 5
const anon_sym_LBRACK = 6
const anon_sym_RBRACK = 7
const sym_keyword = 8
const sym_integer = 9
const sym_float = 10
const anon_sym_true = 11
const anon_sym_false = 12
const anon_sym_DQUOTE = 13
const anon_sym_SQUOTE = 14
const anon_sym_BQUOTE = 15
const anon_sym_DOLLAR_LBRACE = 16
const anon_sym_RBRACE = 17
const aux_sym__escape_sequence_token1 = 18
const sym_escape_sequence = 19
const anon_sym_LBRACE = 20
const anon_sym_COMMA = 21
const anon_sym_COLON = 22
const anon_sym_QMARK_DOT = 23
const anon_sym_DOT = 24
const anon_sym_PLUS = 25
const anon_sym_DASH = 26
const anon_sym_STAR = 27
const anon_sym_SLASH = 28
const anon_sym_PERCENT = 29
const anon_sym_AMP_AMP = 30
const anon_sym_PIPE_PIPE = 31
const anon_sym_EQ_EQ = 32
const anon_sym_BANG_EQ = 33
const anon_sym_EQ_TILDE = 34
const anon_sym_GT_EQ = 35
const anon_sym_LT_EQ = 36
const anon_sym_GT = 37
const anon_sym_LT = 38
const anon_sym_QMARK_COLON = 39
const anon_sym_BANG = 40
const anon_sym_QMARK = 41
const sym_ident = 42
const sym_comment = 43
const sym__unescaped_single_quote_string_fragment = 44
const sym__unescaped_double_quote_string_fragment = 45
const sym__unescaped_backtick_string_fragment = 46
const sym_source_file = 47
const sym_ast_block = 48
const sym_loop_widget = 49
const sym_list = 50
const sym_array = 51
const sym_literal = 52
const sym_number = 53
const sym_boolean = 54
const sym_string = 55
const sym_string_interpolation = 56
const sym__escape_sequence = 57
const sym_expr = 58
const sym_simplexpr = 59
const sym_json_array = 60
const sym_json_object = 61
const sym_json_access = 62
const sym_json_safe_access = 63
const sym_json_dot_access = 64
const sym_json_safe_dot_access = 65
const sym_function_call = 66
const sym_binary_expression = 67
const sym_unary_expression = 68
const sym_ternary_expression = 69
const sym_parenthesized_expression = 70
const aux_sym_source_file_repeat1 = 71
const aux_sym_string_repeat1 = 72
const aux_sym_string_repeat2 = 73
const aux_sym_string_repeat3 = 74
const aux_sym_string_repeat4 = 75
const aux_sym_string_repeat5 = 76
const aux_sym_string_repeat6 = 77
const aux_sym_json_array_repeat1 = 78
const aux_sym_json_object_repeat1 = 79
const alias_sym_index = 80
const alias_sym_string_fragment = 81

var ts_symbol_names = [82]uintptr{
	0:  __ccgo_ts,
	1:  __ccgo_ts + 4,
	2:  __ccgo_ts + 11,
	3:  __ccgo_ts + 13,
	4:  __ccgo_ts + 17,
	5:  __ccgo_ts + 20,
	6:  __ccgo_ts + 22,
	7:  __ccgo_ts + 24,
	8:  __ccgo_ts + 26,
	9:  __ccgo_ts + 34,
	10: __ccgo_ts + 42,
	11: __ccgo_ts + 48,
	12: __ccgo_ts + 53,
	13: __ccgo_ts + 59,
	14: __ccgo_ts + 61,
	15: __ccgo_ts + 63,
	16: __ccgo_ts + 65,
	17: __ccgo_ts + 68,
	18: __ccgo_ts + 70,
	19: __ccgo_ts + 94,
	20: __ccgo_ts + 110,
	21: __ccgo_ts + 112,
	22: __ccgo_ts + 114,
	23: __ccgo_ts + 116,
	24: __ccgo_ts + 119,
	25: __ccgo_ts + 121,
	26: __ccgo_ts + 123,
	27: __ccgo_ts + 125,
	28: __ccgo_ts + 127,
	29: __ccgo_ts + 129,
	30: __ccgo_ts + 131,
	31: __ccgo_ts + 134,
	32: __ccgo_ts + 137,
	33: __ccgo_ts + 140,
	34: __ccgo_ts + 143,
	35: __ccgo_ts + 146,
	36: __ccgo_ts + 149,
	37: __ccgo_ts + 152,
	38: __ccgo_ts + 154,
	39: __ccgo_ts + 156,
	40: __ccgo_ts + 159,
	41: __ccgo_ts + 161,
	42: __ccgo_ts + 163,
	43: __ccgo_ts + 169,
	44: __ccgo_ts + 177,
	45: __ccgo_ts + 217,
	46: __ccgo_ts + 257,
	47: __ccgo_ts + 293,
	48: __ccgo_ts + 305,
	49: __ccgo_ts + 315,
	50: __ccgo_ts + 327,
	51: __ccgo_ts + 332,
	52: __ccgo_ts + 338,
	53: __ccgo_ts + 346,
	54: __ccgo_ts + 353,
	55: __ccgo_ts + 361,
	56: __ccgo_ts + 368,
	57: __ccgo_ts + 389,
	58: __ccgo_ts + 406,
	59: __ccgo_ts + 411,
	60: __ccgo_ts + 421,
	61: __ccgo_ts + 432,
	62: __ccgo_ts + 444,
	63: __ccgo_ts + 456,
	64: __ccgo_ts + 473,
	65: __ccgo_ts + 489,
	66: __ccgo_ts + 510,
	67: __ccgo_ts + 524,
	68: __ccgo_ts + 542,
	69: __ccgo_ts + 559,
	70: __ccgo_ts + 578,
	71: __ccgo_ts + 603,
	72: __ccgo_ts + 623,
	73: __ccgo_ts + 638,
	74: __ccgo_ts + 653,
	75: __ccgo_ts + 668,
	76: __ccgo_ts + 683,
	77: __ccgo_ts + 698,
	78: __ccgo_ts + 713,
	79: __ccgo_ts + 732,
	80: __ccgo_ts + 752,
	81: __ccgo_ts + 758,
}

var ts_symbol_map = [82]TSSymbol{
	1:  uint16(sym_symbol),
	2:  uint16(anon_sym_LPAREN),
	3:  uint16(anon_sym_for),
	4:  uint16(anon_sym_in),
	5:  uint16(anon_sym_RPAREN),
	6:  uint16(anon_sym_LBRACK),
	7:  uint16(anon_sym_RBRACK),
	8:  uint16(sym_keyword),
	9:  uint16(sym_integer),
	10: uint16(sym_float),
	11: uint16(anon_sym_true),
	12: uint16(anon_sym_false),
	13: uint16(anon_sym_DQUOTE),
	14: uint16(anon_sym_SQUOTE),
	15: uint16(anon_sym_BQUOTE),
	16: uint16(anon_sym_DOLLAR_LBRACE),
	17: uint16(anon_sym_RBRACE),
	18: uint16(aux_sym__escape_sequence_token1),
	19: uint16(sym_escape_sequence),
	20: uint16(anon_sym_LBRACE),
	21: uint16(anon_sym_COMMA),
	22: uint16(anon_sym_COLON),
	23: uint16(anon_sym_QMARK_DOT),
	24: uint16(anon_sym_DOT),
	25: uint16(anon_sym_PLUS),
	26: uint16(anon_sym_DASH),
	27: uint16(anon_sym_STAR),
	28: uint16(anon_sym_SLASH),
	29: uint16(anon_sym_PERCENT),
	30: uint16(anon_sym_AMP_AMP),
	31: uint16(anon_sym_PIPE_PIPE),
	32: uint16(anon_sym_EQ_EQ),
	33: uint16(anon_sym_BANG_EQ),
	34: uint16(anon_sym_EQ_TILDE),
	35: uint16(anon_sym_GT_EQ),
	36: uint16(anon_sym_LT_EQ),
	37: uint16(anon_sym_GT),
	38: uint16(anon_sym_LT),
	39: uint16(anon_sym_QMARK_COLON),
	40: uint16(anon_sym_BANG),
	41: uint16(anon_sym_QMARK),
	42: uint16(sym_ident),
	43: uint16(sym_comment),
	44: uint16(sym__unescaped_single_quote_string_fragment),
	45: uint16(sym__unescaped_double_quote_string_fragment),
	46: uint16(sym__unescaped_backtick_string_fragment),
	47: uint16(sym_source_file),
	48: uint16(sym_ast_block),
	49: uint16(sym_loop_widget),
	50: uint16(sym_list),
	51: uint16(sym_array),
	52: uint16(sym_literal),
	53: uint16(sym_number),
	54: uint16(sym_boolean),
	55: uint16(sym_string),
	56: uint16(sym_string_interpolation),
	57: uint16(sym__escape_sequence),
	58: uint16(sym_expr),
	59: uint16(sym_simplexpr),
	60: uint16(sym_json_array),
	61: uint16(sym_json_object),
	62: uint16(sym_json_access),
	63: uint16(sym_json_safe_access),
	64: uint16(sym_json_dot_access),
	65: uint16(sym_json_safe_dot_access),
	66: uint16(sym_function_call),
	67: uint16(sym_binary_expression),
	68: uint16(sym_unary_expression),
	69: uint16(sym_ternary_expression),
	70: uint16(sym_parenthesized_expression),
	71: uint16(aux_sym_source_file_repeat1),
	72: uint16(aux_sym_string_repeat1),
	73: uint16(aux_sym_string_repeat2),
	74: uint16(aux_sym_string_repeat3),
	75: uint16(aux_sym_string_repeat4),
	76: uint16(aux_sym_string_repeat5),
	77: uint16(aux_sym_string_repeat6),
	78: uint16(aux_sym_json_array_repeat1),
	79: uint16(aux_sym_json_object_repeat1),
	80: uint16(alias_sym_index),
	81: uint16(alias_sym_string_fragment),
}

var ts_symbol_metadata = [82]TSSymbolMetadata{
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
	18: {},
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
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	47: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	48: {
		Fnamed:     libc.BoolUint8(true1 != 0),
		Fsupertype: libc.BoolUint8(true1 != 0),
	},
	49: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
		Fnamed:     libc.BoolUint8(true1 != 0),
		Fsupertype: libc.BoolUint8(true1 != 0),
	},
	53: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
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
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	62: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	63: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
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
	71: {},
	72: {},
	73: {},
	74: {},
	75: {},
	76: {},
	77: {},
	78: {},
	79: {},
	80: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	81: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
}

type ts_field_identifiers = int32

const field_alternative = 1
const field_argument = 2
const field_condition = 3
const field_consequence = 4
const field_left = 5
const field_name = 6
const field_operator = 7
const field_right = 8

var ts_field_names = [9]uintptr{
	0: libc.UintptrFromInt32(0),
	1: __ccgo_ts + 774,
	2: __ccgo_ts + 786,
	3: __ccgo_ts + 795,
	4: __ccgo_ts + 805,
	5: __ccgo_ts + 817,
	6: __ccgo_ts + 822,
	7: __ccgo_ts + 827,
	8: __ccgo_ts + 836,
}

var ts_field_map_slices = [7]TSMapSlice{
	2: {
		Flength: uint16(2),
	},
	3: {
		Findex:  uint16(2),
		Flength: uint16(1),
	},
	5: {
		Findex:  uint16(3),
		Flength: uint16(3),
	},
	6: {
		Findex:  uint16(6),
		Flength: uint16(3),
	},
}

var ts_field_map_entries = [9]TSFieldMapEntry{
	0: {
		Ffield_id:    uint16(field_argument),
		Fchild_index: uint8(1),
	},
	1: {
		Ffield_id: uint16(field_operator),
	},
	2: {
		Ffield_id: uint16(field_name),
	},
	3: {
		Ffield_id: uint16(field_left),
	},
	4: {
		Ffield_id:    uint16(field_operator),
		Fchild_index: uint8(1),
	},
	5: {
		Ffield_id:    uint16(field_right),
		Fchild_index: uint8(2),
	},
	6: {
		Ffield_id:    uint16(field_alternative),
		Fchild_index: uint8(4),
	},
	7: {
		Ffield_id: uint16(field_condition),
	},
	8: {
		Ffield_id:    uint16(field_consequence),
		Fchild_index: uint8(2),
	},
}

var ts_alias_sequences = [7][7]TSSymbol{
	0: {},
	1: {
		0: uint16(alias_sym_string_fragment),
	},
	4: {
		2: uint16(alias_sym_index),
	},
}

var ts_non_terminal_alias_map = [13]uint16_t{
	0:  uint16(aux_sym_string_repeat1),
	1:  uint16(2),
	2:  uint16(aux_sym_string_repeat1),
	3:  uint16(alias_sym_string_fragment),
	4:  uint16(aux_sym_string_repeat3),
	5:  uint16(2),
	6:  uint16(aux_sym_string_repeat3),
	7:  uint16(alias_sym_string_fragment),
	8:  uint16(aux_sym_string_repeat5),
	9:  uint16(2),
	10: uint16(aux_sym_string_repeat5),
	11: uint16(alias_sym_string_fragment),
}

var ts_primary_state_ids = [127]TSStateId{
	1:   uint16(1),
	2:   uint16(2),
	3:   uint16(3),
	4:   uint16(4),
	5:   uint16(5),
	6:   uint16(5),
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
	23:  uint16(5),
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
	54:  uint16(54),
	55:  uint16(55),
	56:  uint16(56),
	57:  uint16(57),
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
	69:  uint16(66),
	70:  uint16(70),
	71:  uint16(71),
	72:  uint16(72),
	73:  uint16(73),
	74:  uint16(74),
	75:  uint16(66),
	76:  uint16(49),
	77:  uint16(77),
	78:  uint16(51),
	79:  uint16(79),
	80:  uint16(80),
	81:  uint16(81),
	82:  uint16(53),
	83:  uint16(83),
	84:  uint16(84),
	85:  uint16(50),
	86:  uint16(52),
	87:  uint16(87),
	88:  uint16(88),
	89:  uint16(89),
	90:  uint16(90),
	91:  uint16(91),
	92:  uint16(92),
	93:  uint16(93),
	94:  uint16(94),
	95:  uint16(93),
	96:  uint16(92),
	97:  uint16(97),
	98:  uint16(94),
	99:  uint16(90),
	100: uint16(89),
	101: uint16(97),
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
	112: uint16(111),
	113: uint16(113),
	114: uint16(114),
	115: uint16(111),
	116: uint16(116),
	117: uint16(117),
	118: uint16(118),
	119: uint16(119),
	120: uint16(120),
	121: uint16(121),
	122: uint16(122),
	123: uint16(123),
	124: uint16(124),
	125: uint16(125),
	126: uint16(126),
}

var ts_supertype_symbols = [2]TSSymbol{
	0: uint16(sym_ast_block),
	1: uint16(sym_literal),
}

var ts_supertype_map_slices = [53]TSMapSlice{
	48: {
		Flength: uint16(8),
	},
	52: {
		Findex:  uint16(8),
		Flength: uint16(2),
	},
}

var ts_supertype_map_entries = [10]TSSymbol{
	0: uint16(sym_array),
	1: uint16(sym_expr),
	2: uint16(sym_keyword),
	3: uint16(sym_list),
	4: uint16(sym_literal),
	5: uint16(sym_loop_widget),
	6: uint16(sym_string),
	7: uint16(sym_symbol),
	8: uint16(sym_boolean),
	9: uint16(sym_number),
}

func ts_lex(tls *libc.TLS, lexer uintptr, state TSStateId) (r uint8) {
	var eof, result, skip uint8
	var i, i1, i2, i3, i4 uint32_t
	var lookahead int32_t
	_, _, _, _, _, _, _, _, _ = eof, i, i1, i2, i3, i4, lookahead, result, skip
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
			state = uint16(20)
			goto next_state
		}
		i = uint32(0)
		for {
			if !(i < libc.Uint32FromInt64(116)/libc.Uint32FromInt64(2)) {
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
			state = uint16(18)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(34)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('_') <= lookahead && lookahead <= int32('z') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(1):
		i1 = uint32(0)
		for {
			if !(i1 < libc.Uint32FromInt64(60)/libc.Uint32FromInt64(2)) {
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
			state = uint16(34)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('_') <= lookahead && lookahead <= int32('z') {
			state = uint16(92)
			goto next_state
		}
		return result
	case int32(2):
		i2 = uint32(0)
		for {
			if !(i2 < libc.Uint32FromInt64(84)/libc.Uint32FromInt64(2)) {
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
			state = uint16(2)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(92)
			goto next_state
		}
		return result
	case int32(3):
		if lookahead == int32('"') {
			state = uint16(40)
			goto next_state
		}
		if lookahead == int32('\'') {
			state = uint16(41)
			goto next_state
		}
		if lookahead == int32(';') {
			state = uint16(93)
			goto next_state
		}
		if lookahead == int32('`') {
			state = uint16(42)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(49)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(3)
			goto next_state
		}
		if lookahead == int32('!') || lookahead == int32('*') || lookahead == int32('+') || int32('-') <= lookahead && lookahead <= int32('/') || lookahead == int32('<') || lookahead == int32('>') || lookahead == int32('?') || int32('A') <= lookahead && lookahead <= int32('Z') || int32('_') <= lookahead && lookahead <= int32('z') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead == int32('&') {
			state = uint16(65)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead == int32('=') {
			state = uint16(67)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(70)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead == int32('=') {
			state = uint16(68)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead == int32('u') {
			state = uint16(9)
			goto next_state
		}
		if lookahead == int32('x') {
			state = uint16(16)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('7') {
			state = uint16(46)
			goto next_state
		}
		if lookahead == int32('"') || lookahead == int32('\'') || lookahead == int32('?') || lookahead == int32('\\') || lookahead == int32('a') || lookahead == int32('b') || lookahead == int32('f') || lookahead == int32('n') || lookahead == int32('r') || int32('t') <= lookahead && lookahead <= int32('v') {
			state = uint16(47)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(45)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead == int32('{') {
			state = uint16(43)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead == int32('{') {
			state = uint16(15)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(13)
			goto next_state
		}
		return result
	case int32(10):
		if lookahead == int32('|') {
			state = uint16(66)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead == int32('}') {
			state = uint16(47)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(11)
			goto next_state
		}
		return result
	case int32(12):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(35)
			goto next_state
		}
		return result
	case int32(13):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(16)
			goto next_state
		}
		return result
	case int32(14):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(47)
			goto next_state
		}
		return result
	case int32(15):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(11)
			goto next_state
		}
		return result
	case int32(16):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(14)
			goto next_state
		}
		return result
	case int32(17):
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32(')') && lookahead != int32(']') && lookahead != int32('}') {
			state = uint16(25)
			goto next_state
		}
		return result
	case int32(18):
		if eof != 0 {
			state = uint16(20)
			goto next_state
		}
		i3 = uint32(0)
		for {
			if !(i3 < libc.Uint32FromInt64(112)/libc.Uint32FromInt64(2)) {
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
			state = uint16(18)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(34)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('_') <= lookahead && lookahead <= int32('z') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(19):
		if eof != 0 {
			state = uint16(20)
			goto next_state
		}
		i4 = uint32(0)
		for {
			if !(i4 < libc.Uint32FromInt64(48)/libc.Uint32FromInt64(2)) {
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
			state = uint16(19)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(34)
			goto next_state
		}
		if lookahead == int32('!') || lookahead == int32('*') || lookahead == int32('+') || int32('-') <= lookahead && lookahead <= int32('<') || lookahead == int32('>') || lookahead == int32('?') || int32('A') <= lookahead && lookahead <= int32('Z') || int32('_') <= lookahead && lookahead <= int32('z') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(20):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(ts_builtin_sym_end)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(21):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(22):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(23):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(24):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(25):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_keyword)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32(')') && lookahead != int32(']') && lookahead != int32('}') {
			state = uint16(25)
			goto next_state
		}
		return result
	case int32(26):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('a') {
			state = uint16(29)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32('(') && lookahead != int32(')') && lookahead != int32('[') && lookahead != int32(']') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(27):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(37)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32('(') && lookahead != int32(')') && lookahead != int32('[') && lookahead != int32(']') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(28):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(39)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32('(') && lookahead != int32(')') && lookahead != int32('[') && lookahead != int32(']') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(29):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('l') {
			state = uint16(31)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32('(') && lookahead != int32(')') && lookahead != int32('[') && lookahead != int32(']') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(30):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(32)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32('(') && lookahead != int32(')') && lookahead != int32('[') && lookahead != int32(']') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(31):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('s') {
			state = uint16(28)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32('(') && lookahead != int32(')') && lookahead != int32('[') && lookahead != int32(']') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(32):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('u') {
			state = uint16(27)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32('(') && lookahead != int32(')') && lookahead != int32('[') && lookahead != int32(']') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(33):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32('(') && lookahead != int32(')') && lookahead != int32('[') && lookahead != int32(']') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(34):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_integer)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(12)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(34)
			goto next_state
		}
		return result
	case int32(35):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_float)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(35)
			goto next_state
		}
		return result
	case int32(36):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_true)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(92)
			goto next_state
		}
		return result
	case int32(37):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_true)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32('(') && lookahead != int32(')') && lookahead != int32('[') && lookahead != int32(']') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(38):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_false)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(92)
			goto next_state
		}
		return result
	case int32(39):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_false)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32('(') && lookahead != int32(')') && lookahead != int32('[') && lookahead != int32(']') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(40):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(41):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(42):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(43):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOLLAR_LBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(44):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(45):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__escape_sequence_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(46):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__escape_sequence_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('7') {
			state = uint16(48)
			goto next_state
		}
		return result
	case int32(47):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_escape_sequence)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(48):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_escape_sequence)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('7') {
			state = uint16(47)
			goto next_state
		}
		return result
	case int32(49):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(50):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COMMA)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(51):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COLON)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(52):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_QMARK_DOT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(53):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_QMARK_DOT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32('(') && lookahead != int32(')') && lookahead != int32('[') && lookahead != int32(']') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(54):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(55):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32('(') && lookahead != int32(')') && lookahead != int32('[') && lookahead != int32(']') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(56):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PLUS)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(57):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PLUS)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32('(') && lookahead != int32(')') && lookahead != int32('[') && lookahead != int32(']') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(58):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(59):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32('(') && lookahead != int32(')') && lookahead != int32('[') && lookahead != int32(']') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(60):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STAR)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(61):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STAR)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32('(') && lookahead != int32(')') && lookahead != int32('[') && lookahead != int32(']') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(62):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SLASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(63):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SLASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32('(') && lookahead != int32(')') && lookahead != int32('[') && lookahead != int32(']') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(64):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PERCENT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(65):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AMP_AMP)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(66):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PIPE_PIPE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(67):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EQ_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(68):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BANG_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(69):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BANG_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32('(') && lookahead != int32(')') && lookahead != int32('[') && lookahead != int32(']') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(70):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EQ_TILDE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(71):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_GT_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(72):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_GT_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32('(') && lookahead != int32(')') && lookahead != int32('[') && lookahead != int32(']') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(73):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(74):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32('(') && lookahead != int32(')') && lookahead != int32('[') && lookahead != int32(']') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(75):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_GT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('=') {
			state = uint16(72)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32('(') && lookahead != int32(')') && lookahead != int32('[') && lookahead != int32(']') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(76):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_GT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('=') {
			state = uint16(71)
			goto next_state
		}
		return result
	case int32(77):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('=') {
			state = uint16(74)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32('(') && lookahead != int32(')') && lookahead != int32('[') && lookahead != int32(']') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(78):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('=') {
			state = uint16(73)
			goto next_state
		}
		return result
	case int32(79):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_QMARK_COLON)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(80):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_QMARK_COLON)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32('(') && lookahead != int32(')') && lookahead != int32('[') && lookahead != int32(']') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(81):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BANG)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(82):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BANG)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('=') {
			state = uint16(69)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32('(') && lookahead != int32(')') && lookahead != int32('[') && lookahead != int32(']') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(83):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_QMARK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(53)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(80)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32('(') && lookahead != int32(')') && lookahead != int32('[') && lookahead != int32(']') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(84):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_QMARK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(52)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(79)
			goto next_state
		}
		return result
	case int32(85):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_ident)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('a') {
			state = uint16(88)
			goto next_state
		}
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('b') <= lookahead && lookahead <= int32('z') {
			state = uint16(92)
			goto next_state
		}
		return result
	case int32(86):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_ident)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(36)
			goto next_state
		}
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(92)
			goto next_state
		}
		return result
	case int32(87):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_ident)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(38)
			goto next_state
		}
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(92)
			goto next_state
		}
		return result
	case int32(88):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_ident)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('l') {
			state = uint16(90)
			goto next_state
		}
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(92)
			goto next_state
		}
		return result
	case int32(89):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_ident)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(91)
			goto next_state
		}
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(92)
			goto next_state
		}
		return result
	case int32(90):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_ident)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('s') {
			state = uint16(87)
			goto next_state
		}
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(92)
			goto next_state
		}
		return result
	case int32(91):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_ident)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('u') {
			state = uint16(86)
			goto next_state
		}
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(92)
			goto next_state
		}
		return result
	case int32(92):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_ident)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(92)
			goto next_state
		}
		return result
	case int32(93):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(93)
			goto next_state
		}
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var map_token = [58]uint16_t{
	0:  uint16('!'),
	1:  uint16(82),
	2:  uint16('"'),
	3:  uint16(40),
	4:  uint16('$'),
	5:  uint16(8),
	6:  uint16('%'),
	7:  uint16(64),
	8:  uint16('&'),
	9:  uint16(4),
	10: uint16('\''),
	11: uint16(41),
	12: uint16('('),
	13: uint16(21),
	14: uint16(')'),
	15: uint16(22),
	16: uint16('*'),
	17: uint16(61),
	18: uint16('+'),
	19: uint16(57),
	20: uint16(','),
	21: uint16(50),
	22: uint16('-'),
	23: uint16(59),
	24: uint16('.'),
	25: uint16(55),
	26: uint16('/'),
	27: uint16(63),
	28: uint16(':'),
	29: uint16(51),
	30: uint16(';'),
	31: uint16(93),
	32: uint16('<'),
	33: uint16(77),
	34: uint16('='),
	35: uint16(5),
	36: uint16('>'),
	37: uint16(75),
	38: uint16('?'),
	39: uint16(83),
	40: uint16('['),
	41: uint16(23),
	42: uint16('\\'),
	43: uint16(7),
	44: uint16(']'),
	45: uint16(24),
	46: uint16('`'),
	47: uint16(42),
	48: uint16('f'),
	49: uint16(26),
	50: uint16('t'),
	51: uint16(30),
	52: uint16('{'),
	53: uint16(49),
	54: uint16('|'),
	55: uint16(10),
	56: uint16('}'),
	57: uint16(44),
}

var map_token1 = [30]uint16_t{
	0:  uint16('!'),
	1:  uint16(81),
	2:  uint16('"'),
	3:  uint16(40),
	4:  uint16('\''),
	5:  uint16(41),
	6:  uint16('('),
	7:  uint16(21),
	8:  uint16(')'),
	9:  uint16(22),
	10: uint16('+'),
	11: uint16(56),
	12: uint16('-'),
	13: uint16(58),
	14: uint16(';'),
	15: uint16(93),
	16: uint16('['),
	17: uint16(23),
	18: uint16(']'),
	19: uint16(24),
	20: uint16('`'),
	21: uint16(42),
	22: uint16('f'),
	23: uint16(85),
	24: uint16('t'),
	25: uint16(89),
	26: uint16('{'),
	27: uint16(49),
	28: uint16('}'),
	29: uint16(44),
}

var map_token2 = [42]uint16_t{
	0:  uint16('!'),
	1:  uint16(6),
	2:  uint16('%'),
	3:  uint16(64),
	4:  uint16('&'),
	5:  uint16(4),
	6:  uint16('('),
	7:  uint16(21),
	8:  uint16(')'),
	9:  uint16(22),
	10: uint16('*'),
	11: uint16(60),
	12: uint16('+'),
	13: uint16(56),
	14: uint16(','),
	15: uint16(50),
	16: uint16('-'),
	17: uint16(58),
	18: uint16('.'),
	19: uint16(54),
	20: uint16('/'),
	21: uint16(62),
	22: uint16(':'),
	23: uint16(51),
	24: uint16(';'),
	25: uint16(93),
	26: uint16('<'),
	27: uint16(78),
	28: uint16('='),
	29: uint16(5),
	30: uint16('>'),
	31: uint16(76),
	32: uint16('?'),
	33: uint16(84),
	34: uint16('['),
	35: uint16(23),
	36: uint16(']'),
	37: uint16(24),
	38: uint16('|'),
	39: uint16(10),
	40: uint16('}'),
	41: uint16(44),
}

var map_token3 = [56]uint16_t{
	0:  uint16('!'),
	1:  uint16(82),
	2:  uint16('"'),
	3:  uint16(40),
	4:  uint16('$'),
	5:  uint16(8),
	6:  uint16('%'),
	7:  uint16(64),
	8:  uint16('&'),
	9:  uint16(4),
	10: uint16('\''),
	11: uint16(41),
	12: uint16('('),
	13: uint16(21),
	14: uint16(')'),
	15: uint16(22),
	16: uint16('*'),
	17: uint16(61),
	18: uint16('+'),
	19: uint16(57),
	20: uint16(','),
	21: uint16(50),
	22: uint16('-'),
	23: uint16(59),
	24: uint16('.'),
	25: uint16(55),
	26: uint16('/'),
	27: uint16(63),
	28: uint16(':'),
	29: uint16(51),
	30: uint16(';'),
	31: uint16(93),
	32: uint16('<'),
	33: uint16(77),
	34: uint16('='),
	35: uint16(5),
	36: uint16('>'),
	37: uint16(75),
	38: uint16('?'),
	39: uint16(83),
	40: uint16('['),
	41: uint16(23),
	42: uint16(']'),
	43: uint16(24),
	44: uint16('`'),
	45: uint16(42),
	46: uint16('f'),
	47: uint16(26),
	48: uint16('t'),
	49: uint16(30),
	50: uint16('{'),
	51: uint16(49),
	52: uint16('|'),
	53: uint16(10),
	54: uint16('}'),
	55: uint16(44),
}

var map_token4 = [24]uint16_t{
	0:  uint16('"'),
	1:  uint16(40),
	2:  uint16('\''),
	3:  uint16(41),
	4:  uint16('('),
	5:  uint16(21),
	6:  uint16(')'),
	7:  uint16(22),
	8:  uint16(':'),
	9:  uint16(17),
	10: uint16(';'),
	11: uint16(93),
	12: uint16('['),
	13: uint16(23),
	14: uint16(']'),
	15: uint16(24),
	16: uint16('`'),
	17: uint16(42),
	18: uint16('f'),
	19: uint16(26),
	20: uint16('t'),
	21: uint16(30),
	22: uint16('{'),
	23: uint16(49),
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
		if lookahead == int32('f') {
			state = uint16(1)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(2)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(0)
			goto next_state
		}
		return result
	case int32(1):
		if lookahead == int32('o') {
			state = uint16(3)
			goto next_state
		}
		return result
	case int32(2):
		if lookahead == int32('n') {
			state = uint16(4)
			goto next_state
		}
		return result
	case int32(3):
		if lookahead == int32('r') {
			state = uint16(5)
			goto next_state
		}
		return result
	case int32(4):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_in)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(5):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_for)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var ts_lex_modes = [127]TSLexerMode{
	0: {
		Fexternal_lex_state: uint16(1),
	},
	1: {
		Flex_state: uint16(19),
	},
	2: {
		Flex_state: uint16(1),
	},
	3: {
		Flex_state: uint16(1),
	},
	4: {
		Flex_state: uint16(1),
	},
	5: {
		Flex_state: uint16(1),
	},
	6: {
		Flex_state: uint16(1),
	},
	7: {
		Flex_state: uint16(1),
	},
	8: {
		Flex_state: uint16(1),
	},
	9: {
		Flex_state: uint16(1),
	},
	10: {
		Flex_state: uint16(1),
	},
	11: {
		Flex_state: uint16(1),
	},
	12: {
		Flex_state: uint16(1),
	},
	13: {
		Flex_state: uint16(1),
	},
	14: {
		Flex_state: uint16(1),
	},
	15: {
		Flex_state: uint16(1),
	},
	16: {
		Flex_state: uint16(1),
	},
	17: {
		Flex_state: uint16(1),
	},
	18: {
		Flex_state: uint16(1),
	},
	19: {
		Flex_state: uint16(1),
	},
	20: {
		Flex_state: uint16(1),
	},
	21: {
		Flex_state: uint16(1),
	},
	22: {
		Flex_state: uint16(1),
	},
	23: {
		Flex_state: uint16(1),
	},
	24: {
		Flex_state: uint16(1),
	},
	25: {
		Flex_state: uint16(19),
	},
	26: {
		Flex_state: uint16(2),
	},
	27: {
		Flex_state: uint16(2),
	},
	28: {
		Flex_state: uint16(2),
	},
	29: {
		Flex_state: uint16(2),
	},
	30: {
		Flex_state: uint16(19),
	},
	31: {
		Flex_state: uint16(2),
	},
	32: {
		Flex_state: uint16(2),
	},
	33: {
		Flex_state: uint16(2),
	},
	34: {
		Flex_state: uint16(2),
	},
	35: {
		Flex_state: uint16(2),
	},
	36: {
		Flex_state: uint16(2),
	},
	37: {
		Flex_state: uint16(2),
	},
	38: {
		Flex_state: uint16(2),
	},
	39: {
		Flex_state: uint16(2),
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
		Flex_state: uint16(2),
	},
	50: {
		Flex_state: uint16(2),
	},
	51: {
		Flex_state: uint16(2),
	},
	52: {
		Flex_state: uint16(2),
	},
	53: {
		Flex_state: uint16(2),
	},
	54: {
		Flex_state: uint16(2),
	},
	55: {
		Flex_state: uint16(2),
	},
	56: {
		Flex_state: uint16(19),
	},
	57: {
		Flex_state: uint16(19),
	},
	58: {
		Flex_state: uint16(19),
	},
	59: {
		Flex_state: uint16(19),
	},
	60: {
		Flex_state: uint16(2),
	},
	61: {
		Flex_state: uint16(2),
	},
	62: {
		Flex_state: uint16(2),
	},
	63: {
		Flex_state: uint16(2),
	},
	64: {
		Flex_state: uint16(2),
	},
	65: {
		Flex_state: uint16(19),
	},
	66: {
		Flex_state: uint16(2),
	},
	67: {
		Flex_state: uint16(2),
	},
	68: {
		Flex_state: uint16(2),
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
		Flex_state: uint16(2),
	},
	75: {
		Flex_state: uint16(2),
	},
	76: {
		Flex_state: uint16(19),
	},
	77: {
		Flex_state: uint16(19),
	},
	78: {
		Flex_state: uint16(19),
	},
	79: {
		Flex_state: uint16(19),
	},
	80: {
		Flex_state: uint16(19),
	},
	81: {
		Flex_state: uint16(19),
	},
	82: {
		Flex_state: uint16(19),
	},
	83: {
		Flex_state: uint16(19),
	},
	84: {
		Flex_state: uint16(19),
	},
	85: {
		Flex_state: uint16(19),
	},
	86: {
		Flex_state: uint16(19),
	},
	87: {
		Flex_state: uint16(19),
	},
	88: {
		Fexternal_lex_state: uint16(2),
	},
	89: {
		Fexternal_lex_state: uint16(3),
	},
	90: {
		Fexternal_lex_state: uint16(2),
	},
	91: {
		Fexternal_lex_state: uint16(3),
	},
	92: {
		Fexternal_lex_state: uint16(2),
	},
	93: {
		Fexternal_lex_state: uint16(4),
	},
	94: {
		Fexternal_lex_state: uint16(4),
	},
	95: {
		Fexternal_lex_state: uint16(4),
	},
	96: {
		Fexternal_lex_state: uint16(2),
	},
	97: {
		Fexternal_lex_state: uint16(3),
	},
	98: {
		Fexternal_lex_state: uint16(4),
	},
	99: {
		Fexternal_lex_state: uint16(2),
	},
	100: {
		Fexternal_lex_state: uint16(3),
	},
	101: {
		Fexternal_lex_state: uint16(3),
	},
	102: {
		Fexternal_lex_state: uint16(4),
	},
	103: {
		Fexternal_lex_state: uint16(3),
	},
	104: {
		Flex_state: uint16(3),
	},
	105: {
		Fexternal_lex_state: uint16(3),
	},
	106: {
		Fexternal_lex_state: uint16(4),
	},
	107: {
		Fexternal_lex_state: uint16(2),
	},
	108: {
		Fexternal_lex_state: uint16(2),
	},
	109: {
		Fexternal_lex_state: uint16(4),
	},
	110: {
		Fexternal_lex_state: uint16(3),
	},
	111: {
		Fexternal_lex_state: uint16(2),
	},
	112: {
		Fexternal_lex_state: uint16(4),
	},
	113: {
		Fexternal_lex_state: uint16(2),
	},
	114: {
		Fexternal_lex_state: uint16(4),
	},
	115: {
		Fexternal_lex_state: uint16(3),
	},
	116: {},
	117: {},
	118: {},
	119: {},
	120: {},
	121: {
		Flex_state: uint16(2),
	},
	122: {
		Flex_state: uint16(3),
	},
	123: {
		Flex_state: uint16(3),
	},
	124: {},
	125: {
		Flex_state: uint16(2),
	},
	126: {},
}

var ts_parse_table = [2][80]uint16_t{
	0: {
		0:  uint16(1),
		1:  uint16(1),
		2:  uint16(1),
		3:  uint16(1),
		4:  uint16(1),
		5:  uint16(1),
		6:  uint16(1),
		7:  uint16(1),
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
		43: uint16(3),
		44: uint16(1),
		45: uint16(1),
		46: uint16(1),
	},
	1: {
		0:  uint16(5),
		1:  uint16(7),
		2:  uint16(9),
		6:  uint16(11),
		8:  uint16(13),
		9:  uint16(15),
		10: uint16(17),
		11: uint16(19),
		12: uint16(19),
		13: uint16(21),
		14: uint16(23),
		15: uint16(25),
		20: uint16(27),
		43: uint16(3),
		47: uint16(124),
		48: uint16(57),
		49: uint16(77),
		50: uint16(77),
		51: uint16(77),
		52: uint16(77),
		53: uint16(78),
		54: uint16(78),
		55: uint16(77),
		58: uint16(77),
		71: uint16(57),
	},
}

var ts_small_parse_table = [4624]uint16_t{
	0:    uint16(16),
	1:    uint16(3),
	2:    uint16(1),
	3:    uint16(sym_comment),
	4:    uint16(29),
	5:    uint16(1),
	6:    uint16(anon_sym_LPAREN),
	7:    uint16(31),
	8:    uint16(1),
	9:    uint16(anon_sym_LBRACK),
	10:   uint16(33),
	11:   uint16(1),
	12:   uint16(anon_sym_RBRACK),
	13:   uint16(35),
	14:   uint16(1),
	15:   uint16(sym_integer),
	16:   uint16(37),
	17:   uint16(1),
	18:   uint16(sym_float),
	19:   uint16(41),
	20:   uint16(1),
	21:   uint16(anon_sym_DQUOTE),
	22:   uint16(43),
	23:   uint16(1),
	24:   uint16(anon_sym_SQUOTE),
	25:   uint16(45),
	26:   uint16(1),
	27:   uint16(anon_sym_BQUOTE),
	28:   uint16(47),
	29:   uint16(1),
	30:   uint16(anon_sym_LBRACE),
	31:   uint16(51),
	32:   uint16(1),
	33:   uint16(sym_ident),
	34:   uint16(63),
	35:   uint16(1),
	36:   uint16(sym_simplexpr),
	37:   uint16(39),
	38:   uint16(2),
	39:   uint16(anon_sym_true),
	40:   uint16(anon_sym_false),
	41:   uint16(51),
	42:   uint16(2),
	43:   uint16(sym_number),
	44:   uint16(sym_boolean),
	45:   uint16(49),
	46:   uint16(3),
	47:   uint16(anon_sym_PLUS),
	48:   uint16(anon_sym_DASH),
	49:   uint16(anon_sym_BANG),
	50:   uint16(43),
	51:   uint16(13),
	52:   uint16(sym_literal),
	53:   uint16(sym_string),
	54:   uint16(sym_json_array),
	55:   uint16(sym_json_object),
	56:   uint16(sym_json_access),
	57:   uint16(sym_json_safe_access),
	58:   uint16(sym_json_dot_access),
	59:   uint16(sym_json_safe_dot_access),
	60:   uint16(sym_function_call),
	61:   uint16(sym_binary_expression),
	62:   uint16(sym_unary_expression),
	63:   uint16(sym_ternary_expression),
	64:   uint16(sym_parenthesized_expression),
	65:   uint16(16),
	66:   uint16(3),
	67:   uint16(1),
	68:   uint16(sym_comment),
	69:   uint16(29),
	70:   uint16(1),
	71:   uint16(anon_sym_LPAREN),
	72:   uint16(31),
	73:   uint16(1),
	74:   uint16(anon_sym_LBRACK),
	75:   uint16(35),
	76:   uint16(1),
	77:   uint16(sym_integer),
	78:   uint16(37),
	79:   uint16(1),
	80:   uint16(sym_float),
	81:   uint16(41),
	82:   uint16(1),
	83:   uint16(anon_sym_DQUOTE),
	84:   uint16(43),
	85:   uint16(1),
	86:   uint16(anon_sym_SQUOTE),
	87:   uint16(45),
	88:   uint16(1),
	89:   uint16(anon_sym_BQUOTE),
	90:   uint16(47),
	91:   uint16(1),
	92:   uint16(anon_sym_LBRACE),
	93:   uint16(51),
	94:   uint16(1),
	95:   uint16(sym_ident),
	96:   uint16(53),
	97:   uint16(1),
	98:   uint16(anon_sym_RBRACE),
	99:   uint16(74),
	100:  uint16(1),
	101:  uint16(sym_simplexpr),
	102:  uint16(39),
	103:  uint16(2),
	104:  uint16(anon_sym_true),
	105:  uint16(anon_sym_false),
	106:  uint16(51),
	107:  uint16(2),
	108:  uint16(sym_number),
	109:  uint16(sym_boolean),
	110:  uint16(49),
	111:  uint16(3),
	112:  uint16(anon_sym_PLUS),
	113:  uint16(anon_sym_DASH),
	114:  uint16(anon_sym_BANG),
	115:  uint16(43),
	116:  uint16(13),
	117:  uint16(sym_literal),
	118:  uint16(sym_string),
	119:  uint16(sym_json_array),
	120:  uint16(sym_json_object),
	121:  uint16(sym_json_access),
	122:  uint16(sym_json_safe_access),
	123:  uint16(sym_json_dot_access),
	124:  uint16(sym_json_safe_dot_access),
	125:  uint16(sym_function_call),
	126:  uint16(sym_binary_expression),
	127:  uint16(sym_unary_expression),
	128:  uint16(sym_ternary_expression),
	129:  uint16(sym_parenthesized_expression),
	130:  uint16(16),
	131:  uint16(3),
	132:  uint16(1),
	133:  uint16(sym_comment),
	134:  uint16(29),
	135:  uint16(1),
	136:  uint16(anon_sym_LPAREN),
	137:  uint16(31),
	138:  uint16(1),
	139:  uint16(anon_sym_LBRACK),
	140:  uint16(35),
	141:  uint16(1),
	142:  uint16(sym_integer),
	143:  uint16(37),
	144:  uint16(1),
	145:  uint16(sym_float),
	146:  uint16(41),
	147:  uint16(1),
	148:  uint16(anon_sym_DQUOTE),
	149:  uint16(43),
	150:  uint16(1),
	151:  uint16(anon_sym_SQUOTE),
	152:  uint16(45),
	153:  uint16(1),
	154:  uint16(anon_sym_BQUOTE),
	155:  uint16(47),
	156:  uint16(1),
	157:  uint16(anon_sym_LBRACE),
	158:  uint16(51),
	159:  uint16(1),
	160:  uint16(sym_ident),
	161:  uint16(55),
	162:  uint16(1),
	163:  uint16(anon_sym_RPAREN),
	164:  uint16(60),
	165:  uint16(1),
	166:  uint16(sym_simplexpr),
	167:  uint16(39),
	168:  uint16(2),
	169:  uint16(anon_sym_true),
	170:  uint16(anon_sym_false),
	171:  uint16(51),
	172:  uint16(2),
	173:  uint16(sym_number),
	174:  uint16(sym_boolean),
	175:  uint16(49),
	176:  uint16(3),
	177:  uint16(anon_sym_PLUS),
	178:  uint16(anon_sym_DASH),
	179:  uint16(anon_sym_BANG),
	180:  uint16(43),
	181:  uint16(13),
	182:  uint16(sym_literal),
	183:  uint16(sym_string),
	184:  uint16(sym_json_array),
	185:  uint16(sym_json_object),
	186:  uint16(sym_json_access),
	187:  uint16(sym_json_safe_access),
	188:  uint16(sym_json_dot_access),
	189:  uint16(sym_json_safe_dot_access),
	190:  uint16(sym_function_call),
	191:  uint16(sym_binary_expression),
	192:  uint16(sym_unary_expression),
	193:  uint16(sym_ternary_expression),
	194:  uint16(sym_parenthesized_expression),
	195:  uint16(15),
	196:  uint16(3),
	197:  uint16(1),
	198:  uint16(sym_comment),
	199:  uint16(29),
	200:  uint16(1),
	201:  uint16(anon_sym_LPAREN),
	202:  uint16(31),
	203:  uint16(1),
	204:  uint16(anon_sym_LBRACK),
	205:  uint16(35),
	206:  uint16(1),
	207:  uint16(sym_integer),
	208:  uint16(37),
	209:  uint16(1),
	210:  uint16(sym_float),
	211:  uint16(41),
	212:  uint16(1),
	213:  uint16(anon_sym_DQUOTE),
	214:  uint16(43),
	215:  uint16(1),
	216:  uint16(anon_sym_SQUOTE),
	217:  uint16(45),
	218:  uint16(1),
	219:  uint16(anon_sym_BQUOTE),
	220:  uint16(47),
	221:  uint16(1),
	222:  uint16(anon_sym_LBRACE),
	223:  uint16(51),
	224:  uint16(1),
	225:  uint16(sym_ident),
	226:  uint16(69),
	227:  uint16(1),
	228:  uint16(sym_simplexpr),
	229:  uint16(39),
	230:  uint16(2),
	231:  uint16(anon_sym_true),
	232:  uint16(anon_sym_false),
	233:  uint16(51),
	234:  uint16(2),
	235:  uint16(sym_number),
	236:  uint16(sym_boolean),
	237:  uint16(49),
	238:  uint16(3),
	239:  uint16(anon_sym_PLUS),
	240:  uint16(anon_sym_DASH),
	241:  uint16(anon_sym_BANG),
	242:  uint16(43),
	243:  uint16(13),
	244:  uint16(sym_literal),
	245:  uint16(sym_string),
	246:  uint16(sym_json_array),
	247:  uint16(sym_json_object),
	248:  uint16(sym_json_access),
	249:  uint16(sym_json_safe_access),
	250:  uint16(sym_json_dot_access),
	251:  uint16(sym_json_safe_dot_access),
	252:  uint16(sym_function_call),
	253:  uint16(sym_binary_expression),
	254:  uint16(sym_unary_expression),
	255:  uint16(sym_ternary_expression),
	256:  uint16(sym_parenthesized_expression),
	257:  uint16(15),
	258:  uint16(3),
	259:  uint16(1),
	260:  uint16(sym_comment),
	261:  uint16(29),
	262:  uint16(1),
	263:  uint16(anon_sym_LPAREN),
	264:  uint16(31),
	265:  uint16(1),
	266:  uint16(anon_sym_LBRACK),
	267:  uint16(35),
	268:  uint16(1),
	269:  uint16(sym_integer),
	270:  uint16(37),
	271:  uint16(1),
	272:  uint16(sym_float),
	273:  uint16(41),
	274:  uint16(1),
	275:  uint16(anon_sym_DQUOTE),
	276:  uint16(43),
	277:  uint16(1),
	278:  uint16(anon_sym_SQUOTE),
	279:  uint16(45),
	280:  uint16(1),
	281:  uint16(anon_sym_BQUOTE),
	282:  uint16(47),
	283:  uint16(1),
	284:  uint16(anon_sym_LBRACE),
	285:  uint16(51),
	286:  uint16(1),
	287:  uint16(sym_ident),
	288:  uint16(66),
	289:  uint16(1),
	290:  uint16(sym_simplexpr),
	291:  uint16(39),
	292:  uint16(2),
	293:  uint16(anon_sym_true),
	294:  uint16(anon_sym_false),
	295:  uint16(51),
	296:  uint16(2),
	297:  uint16(sym_number),
	298:  uint16(sym_boolean),
	299:  uint16(49),
	300:  uint16(3),
	301:  uint16(anon_sym_PLUS),
	302:  uint16(anon_sym_DASH),
	303:  uint16(anon_sym_BANG),
	304:  uint16(43),
	305:  uint16(13),
	306:  uint16(sym_literal),
	307:  uint16(sym_string),
	308:  uint16(sym_json_array),
	309:  uint16(sym_json_object),
	310:  uint16(sym_json_access),
	311:  uint16(sym_json_safe_access),
	312:  uint16(sym_json_dot_access),
	313:  uint16(sym_json_safe_dot_access),
	314:  uint16(sym_function_call),
	315:  uint16(sym_binary_expression),
	316:  uint16(sym_unary_expression),
	317:  uint16(sym_ternary_expression),
	318:  uint16(sym_parenthesized_expression),
	319:  uint16(15),
	320:  uint16(3),
	321:  uint16(1),
	322:  uint16(sym_comment),
	323:  uint16(29),
	324:  uint16(1),
	325:  uint16(anon_sym_LPAREN),
	326:  uint16(31),
	327:  uint16(1),
	328:  uint16(anon_sym_LBRACK),
	329:  uint16(35),
	330:  uint16(1),
	331:  uint16(sym_integer),
	332:  uint16(37),
	333:  uint16(1),
	334:  uint16(sym_float),
	335:  uint16(41),
	336:  uint16(1),
	337:  uint16(anon_sym_DQUOTE),
	338:  uint16(43),
	339:  uint16(1),
	340:  uint16(anon_sym_SQUOTE),
	341:  uint16(45),
	342:  uint16(1),
	343:  uint16(anon_sym_BQUOTE),
	344:  uint16(47),
	345:  uint16(1),
	346:  uint16(anon_sym_LBRACE),
	347:  uint16(51),
	348:  uint16(1),
	349:  uint16(sym_ident),
	350:  uint16(72),
	351:  uint16(1),
	352:  uint16(sym_simplexpr),
	353:  uint16(39),
	354:  uint16(2),
	355:  uint16(anon_sym_true),
	356:  uint16(anon_sym_false),
	357:  uint16(51),
	358:  uint16(2),
	359:  uint16(sym_number),
	360:  uint16(sym_boolean),
	361:  uint16(49),
	362:  uint16(3),
	363:  uint16(anon_sym_PLUS),
	364:  uint16(anon_sym_DASH),
	365:  uint16(anon_sym_BANG),
	366:  uint16(43),
	367:  uint16(13),
	368:  uint16(sym_literal),
	369:  uint16(sym_string),
	370:  uint16(sym_json_array),
	371:  uint16(sym_json_object),
	372:  uint16(sym_json_access),
	373:  uint16(sym_json_safe_access),
	374:  uint16(sym_json_dot_access),
	375:  uint16(sym_json_safe_dot_access),
	376:  uint16(sym_function_call),
	377:  uint16(sym_binary_expression),
	378:  uint16(sym_unary_expression),
	379:  uint16(sym_ternary_expression),
	380:  uint16(sym_parenthesized_expression),
	381:  uint16(15),
	382:  uint16(3),
	383:  uint16(1),
	384:  uint16(sym_comment),
	385:  uint16(29),
	386:  uint16(1),
	387:  uint16(anon_sym_LPAREN),
	388:  uint16(31),
	389:  uint16(1),
	390:  uint16(anon_sym_LBRACK),
	391:  uint16(35),
	392:  uint16(1),
	393:  uint16(sym_integer),
	394:  uint16(37),
	395:  uint16(1),
	396:  uint16(sym_float),
	397:  uint16(41),
	398:  uint16(1),
	399:  uint16(anon_sym_DQUOTE),
	400:  uint16(43),
	401:  uint16(1),
	402:  uint16(anon_sym_SQUOTE),
	403:  uint16(45),
	404:  uint16(1),
	405:  uint16(anon_sym_BQUOTE),
	406:  uint16(47),
	407:  uint16(1),
	408:  uint16(anon_sym_LBRACE),
	409:  uint16(51),
	410:  uint16(1),
	411:  uint16(sym_ident),
	412:  uint16(68),
	413:  uint16(1),
	414:  uint16(sym_simplexpr),
	415:  uint16(39),
	416:  uint16(2),
	417:  uint16(anon_sym_true),
	418:  uint16(anon_sym_false),
	419:  uint16(51),
	420:  uint16(2),
	421:  uint16(sym_number),
	422:  uint16(sym_boolean),
	423:  uint16(49),
	424:  uint16(3),
	425:  uint16(anon_sym_PLUS),
	426:  uint16(anon_sym_DASH),
	427:  uint16(anon_sym_BANG),
	428:  uint16(43),
	429:  uint16(13),
	430:  uint16(sym_literal),
	431:  uint16(sym_string),
	432:  uint16(sym_json_array),
	433:  uint16(sym_json_object),
	434:  uint16(sym_json_access),
	435:  uint16(sym_json_safe_access),
	436:  uint16(sym_json_dot_access),
	437:  uint16(sym_json_safe_dot_access),
	438:  uint16(sym_function_call),
	439:  uint16(sym_binary_expression),
	440:  uint16(sym_unary_expression),
	441:  uint16(sym_ternary_expression),
	442:  uint16(sym_parenthesized_expression),
	443:  uint16(15),
	444:  uint16(3),
	445:  uint16(1),
	446:  uint16(sym_comment),
	447:  uint16(29),
	448:  uint16(1),
	449:  uint16(anon_sym_LPAREN),
	450:  uint16(31),
	451:  uint16(1),
	452:  uint16(anon_sym_LBRACK),
	453:  uint16(35),
	454:  uint16(1),
	455:  uint16(sym_integer),
	456:  uint16(37),
	457:  uint16(1),
	458:  uint16(sym_float),
	459:  uint16(41),
	460:  uint16(1),
	461:  uint16(anon_sym_DQUOTE),
	462:  uint16(43),
	463:  uint16(1),
	464:  uint16(anon_sym_SQUOTE),
	465:  uint16(45),
	466:  uint16(1),
	467:  uint16(anon_sym_BQUOTE),
	468:  uint16(47),
	469:  uint16(1),
	470:  uint16(anon_sym_LBRACE),
	471:  uint16(51),
	472:  uint16(1),
	473:  uint16(sym_ident),
	474:  uint16(41),
	475:  uint16(1),
	476:  uint16(sym_simplexpr),
	477:  uint16(39),
	478:  uint16(2),
	479:  uint16(anon_sym_true),
	480:  uint16(anon_sym_false),
	481:  uint16(51),
	482:  uint16(2),
	483:  uint16(sym_number),
	484:  uint16(sym_boolean),
	485:  uint16(49),
	486:  uint16(3),
	487:  uint16(anon_sym_PLUS),
	488:  uint16(anon_sym_DASH),
	489:  uint16(anon_sym_BANG),
	490:  uint16(43),
	491:  uint16(13),
	492:  uint16(sym_literal),
	493:  uint16(sym_string),
	494:  uint16(sym_json_array),
	495:  uint16(sym_json_object),
	496:  uint16(sym_json_access),
	497:  uint16(sym_json_safe_access),
	498:  uint16(sym_json_dot_access),
	499:  uint16(sym_json_safe_dot_access),
	500:  uint16(sym_function_call),
	501:  uint16(sym_binary_expression),
	502:  uint16(sym_unary_expression),
	503:  uint16(sym_ternary_expression),
	504:  uint16(sym_parenthesized_expression),
	505:  uint16(15),
	506:  uint16(3),
	507:  uint16(1),
	508:  uint16(sym_comment),
	509:  uint16(29),
	510:  uint16(1),
	511:  uint16(anon_sym_LPAREN),
	512:  uint16(31),
	513:  uint16(1),
	514:  uint16(anon_sym_LBRACK),
	515:  uint16(35),
	516:  uint16(1),
	517:  uint16(sym_integer),
	518:  uint16(37),
	519:  uint16(1),
	520:  uint16(sym_float),
	521:  uint16(41),
	522:  uint16(1),
	523:  uint16(anon_sym_DQUOTE),
	524:  uint16(43),
	525:  uint16(1),
	526:  uint16(anon_sym_SQUOTE),
	527:  uint16(45),
	528:  uint16(1),
	529:  uint16(anon_sym_BQUOTE),
	530:  uint16(47),
	531:  uint16(1),
	532:  uint16(anon_sym_LBRACE),
	533:  uint16(51),
	534:  uint16(1),
	535:  uint16(sym_ident),
	536:  uint16(67),
	537:  uint16(1),
	538:  uint16(sym_simplexpr),
	539:  uint16(39),
	540:  uint16(2),
	541:  uint16(anon_sym_true),
	542:  uint16(anon_sym_false),
	543:  uint16(51),
	544:  uint16(2),
	545:  uint16(sym_number),
	546:  uint16(sym_boolean),
	547:  uint16(49),
	548:  uint16(3),
	549:  uint16(anon_sym_PLUS),
	550:  uint16(anon_sym_DASH),
	551:  uint16(anon_sym_BANG),
	552:  uint16(43),
	553:  uint16(13),
	554:  uint16(sym_literal),
	555:  uint16(sym_string),
	556:  uint16(sym_json_array),
	557:  uint16(sym_json_object),
	558:  uint16(sym_json_access),
	559:  uint16(sym_json_safe_access),
	560:  uint16(sym_json_dot_access),
	561:  uint16(sym_json_safe_dot_access),
	562:  uint16(sym_function_call),
	563:  uint16(sym_binary_expression),
	564:  uint16(sym_unary_expression),
	565:  uint16(sym_ternary_expression),
	566:  uint16(sym_parenthesized_expression),
	567:  uint16(15),
	568:  uint16(3),
	569:  uint16(1),
	570:  uint16(sym_comment),
	571:  uint16(29),
	572:  uint16(1),
	573:  uint16(anon_sym_LPAREN),
	574:  uint16(31),
	575:  uint16(1),
	576:  uint16(anon_sym_LBRACK),
	577:  uint16(35),
	578:  uint16(1),
	579:  uint16(sym_integer),
	580:  uint16(37),
	581:  uint16(1),
	582:  uint16(sym_float),
	583:  uint16(41),
	584:  uint16(1),
	585:  uint16(anon_sym_DQUOTE),
	586:  uint16(43),
	587:  uint16(1),
	588:  uint16(anon_sym_SQUOTE),
	589:  uint16(45),
	590:  uint16(1),
	591:  uint16(anon_sym_BQUOTE),
	592:  uint16(47),
	593:  uint16(1),
	594:  uint16(anon_sym_LBRACE),
	595:  uint16(51),
	596:  uint16(1),
	597:  uint16(sym_ident),
	598:  uint16(33),
	599:  uint16(1),
	600:  uint16(sym_simplexpr),
	601:  uint16(39),
	602:  uint16(2),
	603:  uint16(anon_sym_true),
	604:  uint16(anon_sym_false),
	605:  uint16(51),
	606:  uint16(2),
	607:  uint16(sym_number),
	608:  uint16(sym_boolean),
	609:  uint16(49),
	610:  uint16(3),
	611:  uint16(anon_sym_PLUS),
	612:  uint16(anon_sym_DASH),
	613:  uint16(anon_sym_BANG),
	614:  uint16(43),
	615:  uint16(13),
	616:  uint16(sym_literal),
	617:  uint16(sym_string),
	618:  uint16(sym_json_array),
	619:  uint16(sym_json_object),
	620:  uint16(sym_json_access),
	621:  uint16(sym_json_safe_access),
	622:  uint16(sym_json_dot_access),
	623:  uint16(sym_json_safe_dot_access),
	624:  uint16(sym_function_call),
	625:  uint16(sym_binary_expression),
	626:  uint16(sym_unary_expression),
	627:  uint16(sym_ternary_expression),
	628:  uint16(sym_parenthesized_expression),
	629:  uint16(15),
	630:  uint16(3),
	631:  uint16(1),
	632:  uint16(sym_comment),
	633:  uint16(29),
	634:  uint16(1),
	635:  uint16(anon_sym_LPAREN),
	636:  uint16(31),
	637:  uint16(1),
	638:  uint16(anon_sym_LBRACK),
	639:  uint16(35),
	640:  uint16(1),
	641:  uint16(sym_integer),
	642:  uint16(37),
	643:  uint16(1),
	644:  uint16(sym_float),
	645:  uint16(41),
	646:  uint16(1),
	647:  uint16(anon_sym_DQUOTE),
	648:  uint16(43),
	649:  uint16(1),
	650:  uint16(anon_sym_SQUOTE),
	651:  uint16(45),
	652:  uint16(1),
	653:  uint16(anon_sym_BQUOTE),
	654:  uint16(47),
	655:  uint16(1),
	656:  uint16(anon_sym_LBRACE),
	657:  uint16(51),
	658:  uint16(1),
	659:  uint16(sym_ident),
	660:  uint16(34),
	661:  uint16(1),
	662:  uint16(sym_simplexpr),
	663:  uint16(39),
	664:  uint16(2),
	665:  uint16(anon_sym_true),
	666:  uint16(anon_sym_false),
	667:  uint16(51),
	668:  uint16(2),
	669:  uint16(sym_number),
	670:  uint16(sym_boolean),
	671:  uint16(49),
	672:  uint16(3),
	673:  uint16(anon_sym_PLUS),
	674:  uint16(anon_sym_DASH),
	675:  uint16(anon_sym_BANG),
	676:  uint16(43),
	677:  uint16(13),
	678:  uint16(sym_literal),
	679:  uint16(sym_string),
	680:  uint16(sym_json_array),
	681:  uint16(sym_json_object),
	682:  uint16(sym_json_access),
	683:  uint16(sym_json_safe_access),
	684:  uint16(sym_json_dot_access),
	685:  uint16(sym_json_safe_dot_access),
	686:  uint16(sym_function_call),
	687:  uint16(sym_binary_expression),
	688:  uint16(sym_unary_expression),
	689:  uint16(sym_ternary_expression),
	690:  uint16(sym_parenthesized_expression),
	691:  uint16(15),
	692:  uint16(3),
	693:  uint16(1),
	694:  uint16(sym_comment),
	695:  uint16(29),
	696:  uint16(1),
	697:  uint16(anon_sym_LPAREN),
	698:  uint16(31),
	699:  uint16(1),
	700:  uint16(anon_sym_LBRACK),
	701:  uint16(35),
	702:  uint16(1),
	703:  uint16(sym_integer),
	704:  uint16(37),
	705:  uint16(1),
	706:  uint16(sym_float),
	707:  uint16(41),
	708:  uint16(1),
	709:  uint16(anon_sym_DQUOTE),
	710:  uint16(43),
	711:  uint16(1),
	712:  uint16(anon_sym_SQUOTE),
	713:  uint16(45),
	714:  uint16(1),
	715:  uint16(anon_sym_BQUOTE),
	716:  uint16(47),
	717:  uint16(1),
	718:  uint16(anon_sym_LBRACE),
	719:  uint16(51),
	720:  uint16(1),
	721:  uint16(sym_ident),
	722:  uint16(35),
	723:  uint16(1),
	724:  uint16(sym_simplexpr),
	725:  uint16(39),
	726:  uint16(2),
	727:  uint16(anon_sym_true),
	728:  uint16(anon_sym_false),
	729:  uint16(51),
	730:  uint16(2),
	731:  uint16(sym_number),
	732:  uint16(sym_boolean),
	733:  uint16(49),
	734:  uint16(3),
	735:  uint16(anon_sym_PLUS),
	736:  uint16(anon_sym_DASH),
	737:  uint16(anon_sym_BANG),
	738:  uint16(43),
	739:  uint16(13),
	740:  uint16(sym_literal),
	741:  uint16(sym_string),
	742:  uint16(sym_json_array),
	743:  uint16(sym_json_object),
	744:  uint16(sym_json_access),
	745:  uint16(sym_json_safe_access),
	746:  uint16(sym_json_dot_access),
	747:  uint16(sym_json_safe_dot_access),
	748:  uint16(sym_function_call),
	749:  uint16(sym_binary_expression),
	750:  uint16(sym_unary_expression),
	751:  uint16(sym_ternary_expression),
	752:  uint16(sym_parenthesized_expression),
	753:  uint16(15),
	754:  uint16(3),
	755:  uint16(1),
	756:  uint16(sym_comment),
	757:  uint16(29),
	758:  uint16(1),
	759:  uint16(anon_sym_LPAREN),
	760:  uint16(31),
	761:  uint16(1),
	762:  uint16(anon_sym_LBRACK),
	763:  uint16(35),
	764:  uint16(1),
	765:  uint16(sym_integer),
	766:  uint16(37),
	767:  uint16(1),
	768:  uint16(sym_float),
	769:  uint16(41),
	770:  uint16(1),
	771:  uint16(anon_sym_DQUOTE),
	772:  uint16(43),
	773:  uint16(1),
	774:  uint16(anon_sym_SQUOTE),
	775:  uint16(45),
	776:  uint16(1),
	777:  uint16(anon_sym_BQUOTE),
	778:  uint16(47),
	779:  uint16(1),
	780:  uint16(anon_sym_LBRACE),
	781:  uint16(51),
	782:  uint16(1),
	783:  uint16(sym_ident),
	784:  uint16(36),
	785:  uint16(1),
	786:  uint16(sym_simplexpr),
	787:  uint16(39),
	788:  uint16(2),
	789:  uint16(anon_sym_true),
	790:  uint16(anon_sym_false),
	791:  uint16(51),
	792:  uint16(2),
	793:  uint16(sym_number),
	794:  uint16(sym_boolean),
	795:  uint16(49),
	796:  uint16(3),
	797:  uint16(anon_sym_PLUS),
	798:  uint16(anon_sym_DASH),
	799:  uint16(anon_sym_BANG),
	800:  uint16(43),
	801:  uint16(13),
	802:  uint16(sym_literal),
	803:  uint16(sym_string),
	804:  uint16(sym_json_array),
	805:  uint16(sym_json_object),
	806:  uint16(sym_json_access),
	807:  uint16(sym_json_safe_access),
	808:  uint16(sym_json_dot_access),
	809:  uint16(sym_json_safe_dot_access),
	810:  uint16(sym_function_call),
	811:  uint16(sym_binary_expression),
	812:  uint16(sym_unary_expression),
	813:  uint16(sym_ternary_expression),
	814:  uint16(sym_parenthesized_expression),
	815:  uint16(15),
	816:  uint16(3),
	817:  uint16(1),
	818:  uint16(sym_comment),
	819:  uint16(29),
	820:  uint16(1),
	821:  uint16(anon_sym_LPAREN),
	822:  uint16(31),
	823:  uint16(1),
	824:  uint16(anon_sym_LBRACK),
	825:  uint16(35),
	826:  uint16(1),
	827:  uint16(sym_integer),
	828:  uint16(37),
	829:  uint16(1),
	830:  uint16(sym_float),
	831:  uint16(41),
	832:  uint16(1),
	833:  uint16(anon_sym_DQUOTE),
	834:  uint16(43),
	835:  uint16(1),
	836:  uint16(anon_sym_SQUOTE),
	837:  uint16(45),
	838:  uint16(1),
	839:  uint16(anon_sym_BQUOTE),
	840:  uint16(47),
	841:  uint16(1),
	842:  uint16(anon_sym_LBRACE),
	843:  uint16(51),
	844:  uint16(1),
	845:  uint16(sym_ident),
	846:  uint16(27),
	847:  uint16(1),
	848:  uint16(sym_simplexpr),
	849:  uint16(39),
	850:  uint16(2),
	851:  uint16(anon_sym_true),
	852:  uint16(anon_sym_false),
	853:  uint16(51),
	854:  uint16(2),
	855:  uint16(sym_number),
	856:  uint16(sym_boolean),
	857:  uint16(49),
	858:  uint16(3),
	859:  uint16(anon_sym_PLUS),
	860:  uint16(anon_sym_DASH),
	861:  uint16(anon_sym_BANG),
	862:  uint16(43),
	863:  uint16(13),
	864:  uint16(sym_literal),
	865:  uint16(sym_string),
	866:  uint16(sym_json_array),
	867:  uint16(sym_json_object),
	868:  uint16(sym_json_access),
	869:  uint16(sym_json_safe_access),
	870:  uint16(sym_json_dot_access),
	871:  uint16(sym_json_safe_dot_access),
	872:  uint16(sym_function_call),
	873:  uint16(sym_binary_expression),
	874:  uint16(sym_unary_expression),
	875:  uint16(sym_ternary_expression),
	876:  uint16(sym_parenthesized_expression),
	877:  uint16(15),
	878:  uint16(3),
	879:  uint16(1),
	880:  uint16(sym_comment),
	881:  uint16(29),
	882:  uint16(1),
	883:  uint16(anon_sym_LPAREN),
	884:  uint16(31),
	885:  uint16(1),
	886:  uint16(anon_sym_LBRACK),
	887:  uint16(35),
	888:  uint16(1),
	889:  uint16(sym_integer),
	890:  uint16(37),
	891:  uint16(1),
	892:  uint16(sym_float),
	893:  uint16(41),
	894:  uint16(1),
	895:  uint16(anon_sym_DQUOTE),
	896:  uint16(43),
	897:  uint16(1),
	898:  uint16(anon_sym_SQUOTE),
	899:  uint16(45),
	900:  uint16(1),
	901:  uint16(anon_sym_BQUOTE),
	902:  uint16(47),
	903:  uint16(1),
	904:  uint16(anon_sym_LBRACE),
	905:  uint16(51),
	906:  uint16(1),
	907:  uint16(sym_ident),
	908:  uint16(37),
	909:  uint16(1),
	910:  uint16(sym_simplexpr),
	911:  uint16(39),
	912:  uint16(2),
	913:  uint16(anon_sym_true),
	914:  uint16(anon_sym_false),
	915:  uint16(51),
	916:  uint16(2),
	917:  uint16(sym_number),
	918:  uint16(sym_boolean),
	919:  uint16(49),
	920:  uint16(3),
	921:  uint16(anon_sym_PLUS),
	922:  uint16(anon_sym_DASH),
	923:  uint16(anon_sym_BANG),
	924:  uint16(43),
	925:  uint16(13),
	926:  uint16(sym_literal),
	927:  uint16(sym_string),
	928:  uint16(sym_json_array),
	929:  uint16(sym_json_object),
	930:  uint16(sym_json_access),
	931:  uint16(sym_json_safe_access),
	932:  uint16(sym_json_dot_access),
	933:  uint16(sym_json_safe_dot_access),
	934:  uint16(sym_function_call),
	935:  uint16(sym_binary_expression),
	936:  uint16(sym_unary_expression),
	937:  uint16(sym_ternary_expression),
	938:  uint16(sym_parenthesized_expression),
	939:  uint16(15),
	940:  uint16(3),
	941:  uint16(1),
	942:  uint16(sym_comment),
	943:  uint16(29),
	944:  uint16(1),
	945:  uint16(anon_sym_LPAREN),
	946:  uint16(31),
	947:  uint16(1),
	948:  uint16(anon_sym_LBRACK),
	949:  uint16(35),
	950:  uint16(1),
	951:  uint16(sym_integer),
	952:  uint16(37),
	953:  uint16(1),
	954:  uint16(sym_float),
	955:  uint16(41),
	956:  uint16(1),
	957:  uint16(anon_sym_DQUOTE),
	958:  uint16(43),
	959:  uint16(1),
	960:  uint16(anon_sym_SQUOTE),
	961:  uint16(45),
	962:  uint16(1),
	963:  uint16(anon_sym_BQUOTE),
	964:  uint16(47),
	965:  uint16(1),
	966:  uint16(anon_sym_LBRACE),
	967:  uint16(51),
	968:  uint16(1),
	969:  uint16(sym_ident),
	970:  uint16(70),
	971:  uint16(1),
	972:  uint16(sym_simplexpr),
	973:  uint16(39),
	974:  uint16(2),
	975:  uint16(anon_sym_true),
	976:  uint16(anon_sym_false),
	977:  uint16(51),
	978:  uint16(2),
	979:  uint16(sym_number),
	980:  uint16(sym_boolean),
	981:  uint16(49),
	982:  uint16(3),
	983:  uint16(anon_sym_PLUS),
	984:  uint16(anon_sym_DASH),
	985:  uint16(anon_sym_BANG),
	986:  uint16(43),
	987:  uint16(13),
	988:  uint16(sym_literal),
	989:  uint16(sym_string),
	990:  uint16(sym_json_array),
	991:  uint16(sym_json_object),
	992:  uint16(sym_json_access),
	993:  uint16(sym_json_safe_access),
	994:  uint16(sym_json_dot_access),
	995:  uint16(sym_json_safe_dot_access),
	996:  uint16(sym_function_call),
	997:  uint16(sym_binary_expression),
	998:  uint16(sym_unary_expression),
	999:  uint16(sym_ternary_expression),
	1000: uint16(sym_parenthesized_expression),
	1001: uint16(15),
	1002: uint16(3),
	1003: uint16(1),
	1004: uint16(sym_comment),
	1005: uint16(29),
	1006: uint16(1),
	1007: uint16(anon_sym_LPAREN),
	1008: uint16(31),
	1009: uint16(1),
	1010: uint16(anon_sym_LBRACK),
	1011: uint16(35),
	1012: uint16(1),
	1013: uint16(sym_integer),
	1014: uint16(37),
	1015: uint16(1),
	1016: uint16(sym_float),
	1017: uint16(41),
	1018: uint16(1),
	1019: uint16(anon_sym_DQUOTE),
	1020: uint16(43),
	1021: uint16(1),
	1022: uint16(anon_sym_SQUOTE),
	1023: uint16(45),
	1024: uint16(1),
	1025: uint16(anon_sym_BQUOTE),
	1026: uint16(47),
	1027: uint16(1),
	1028: uint16(anon_sym_LBRACE),
	1029: uint16(51),
	1030: uint16(1),
	1031: uint16(sym_ident),
	1032: uint16(62),
	1033: uint16(1),
	1034: uint16(sym_simplexpr),
	1035: uint16(39),
	1036: uint16(2),
	1037: uint16(anon_sym_true),
	1038: uint16(anon_sym_false),
	1039: uint16(51),
	1040: uint16(2),
	1041: uint16(sym_number),
	1042: uint16(sym_boolean),
	1043: uint16(49),
	1044: uint16(3),
	1045: uint16(anon_sym_PLUS),
	1046: uint16(anon_sym_DASH),
	1047: uint16(anon_sym_BANG),
	1048: uint16(43),
	1049: uint16(13),
	1050: uint16(sym_literal),
	1051: uint16(sym_string),
	1052: uint16(sym_json_array),
	1053: uint16(sym_json_object),
	1054: uint16(sym_json_access),
	1055: uint16(sym_json_safe_access),
	1056: uint16(sym_json_dot_access),
	1057: uint16(sym_json_safe_dot_access),
	1058: uint16(sym_function_call),
	1059: uint16(sym_binary_expression),
	1060: uint16(sym_unary_expression),
	1061: uint16(sym_ternary_expression),
	1062: uint16(sym_parenthesized_expression),
	1063: uint16(15),
	1064: uint16(3),
	1065: uint16(1),
	1066: uint16(sym_comment),
	1067: uint16(29),
	1068: uint16(1),
	1069: uint16(anon_sym_LPAREN),
	1070: uint16(31),
	1071: uint16(1),
	1072: uint16(anon_sym_LBRACK),
	1073: uint16(35),
	1074: uint16(1),
	1075: uint16(sym_integer),
	1076: uint16(37),
	1077: uint16(1),
	1078: uint16(sym_float),
	1079: uint16(41),
	1080: uint16(1),
	1081: uint16(anon_sym_DQUOTE),
	1082: uint16(43),
	1083: uint16(1),
	1084: uint16(anon_sym_SQUOTE),
	1085: uint16(45),
	1086: uint16(1),
	1087: uint16(anon_sym_BQUOTE),
	1088: uint16(47),
	1089: uint16(1),
	1090: uint16(anon_sym_LBRACE),
	1091: uint16(51),
	1092: uint16(1),
	1093: uint16(sym_ident),
	1094: uint16(61),
	1095: uint16(1),
	1096: uint16(sym_simplexpr),
	1097: uint16(39),
	1098: uint16(2),
	1099: uint16(anon_sym_true),
	1100: uint16(anon_sym_false),
	1101: uint16(51),
	1102: uint16(2),
	1103: uint16(sym_number),
	1104: uint16(sym_boolean),
	1105: uint16(49),
	1106: uint16(3),
	1107: uint16(anon_sym_PLUS),
	1108: uint16(anon_sym_DASH),
	1109: uint16(anon_sym_BANG),
	1110: uint16(43),
	1111: uint16(13),
	1112: uint16(sym_literal),
	1113: uint16(sym_string),
	1114: uint16(sym_json_array),
	1115: uint16(sym_json_object),
	1116: uint16(sym_json_access),
	1117: uint16(sym_json_safe_access),
	1118: uint16(sym_json_dot_access),
	1119: uint16(sym_json_safe_dot_access),
	1120: uint16(sym_function_call),
	1121: uint16(sym_binary_expression),
	1122: uint16(sym_unary_expression),
	1123: uint16(sym_ternary_expression),
	1124: uint16(sym_parenthesized_expression),
	1125: uint16(15),
	1126: uint16(3),
	1127: uint16(1),
	1128: uint16(sym_comment),
	1129: uint16(29),
	1130: uint16(1),
	1131: uint16(anon_sym_LPAREN),
	1132: uint16(31),
	1133: uint16(1),
	1134: uint16(anon_sym_LBRACK),
	1135: uint16(35),
	1136: uint16(1),
	1137: uint16(sym_integer),
	1138: uint16(37),
	1139: uint16(1),
	1140: uint16(sym_float),
	1141: uint16(41),
	1142: uint16(1),
	1143: uint16(anon_sym_DQUOTE),
	1144: uint16(43),
	1145: uint16(1),
	1146: uint16(anon_sym_SQUOTE),
	1147: uint16(45),
	1148: uint16(1),
	1149: uint16(anon_sym_BQUOTE),
	1150: uint16(47),
	1151: uint16(1),
	1152: uint16(anon_sym_LBRACE),
	1153: uint16(51),
	1154: uint16(1),
	1155: uint16(sym_ident),
	1156: uint16(46),
	1157: uint16(1),
	1158: uint16(sym_simplexpr),
	1159: uint16(39),
	1160: uint16(2),
	1161: uint16(anon_sym_true),
	1162: uint16(anon_sym_false),
	1163: uint16(51),
	1164: uint16(2),
	1165: uint16(sym_number),
	1166: uint16(sym_boolean),
	1167: uint16(49),
	1168: uint16(3),
	1169: uint16(anon_sym_PLUS),
	1170: uint16(anon_sym_DASH),
	1171: uint16(anon_sym_BANG),
	1172: uint16(43),
	1173: uint16(13),
	1174: uint16(sym_literal),
	1175: uint16(sym_string),
	1176: uint16(sym_json_array),
	1177: uint16(sym_json_object),
	1178: uint16(sym_json_access),
	1179: uint16(sym_json_safe_access),
	1180: uint16(sym_json_dot_access),
	1181: uint16(sym_json_safe_dot_access),
	1182: uint16(sym_function_call),
	1183: uint16(sym_binary_expression),
	1184: uint16(sym_unary_expression),
	1185: uint16(sym_ternary_expression),
	1186: uint16(sym_parenthesized_expression),
	1187: uint16(15),
	1188: uint16(3),
	1189: uint16(1),
	1190: uint16(sym_comment),
	1191: uint16(29),
	1192: uint16(1),
	1193: uint16(anon_sym_LPAREN),
	1194: uint16(31),
	1195: uint16(1),
	1196: uint16(anon_sym_LBRACK),
	1197: uint16(35),
	1198: uint16(1),
	1199: uint16(sym_integer),
	1200: uint16(37),
	1201: uint16(1),
	1202: uint16(sym_float),
	1203: uint16(41),
	1204: uint16(1),
	1205: uint16(anon_sym_DQUOTE),
	1206: uint16(43),
	1207: uint16(1),
	1208: uint16(anon_sym_SQUOTE),
	1209: uint16(45),
	1210: uint16(1),
	1211: uint16(anon_sym_BQUOTE),
	1212: uint16(47),
	1213: uint16(1),
	1214: uint16(anon_sym_LBRACE),
	1215: uint16(51),
	1216: uint16(1),
	1217: uint16(sym_ident),
	1218: uint16(73),
	1219: uint16(1),
	1220: uint16(sym_simplexpr),
	1221: uint16(39),
	1222: uint16(2),
	1223: uint16(anon_sym_true),
	1224: uint16(anon_sym_false),
	1225: uint16(51),
	1226: uint16(2),
	1227: uint16(sym_number),
	1228: uint16(sym_boolean),
	1229: uint16(49),
	1230: uint16(3),
	1231: uint16(anon_sym_PLUS),
	1232: uint16(anon_sym_DASH),
	1233: uint16(anon_sym_BANG),
	1234: uint16(43),
	1235: uint16(13),
	1236: uint16(sym_literal),
	1237: uint16(sym_string),
	1238: uint16(sym_json_array),
	1239: uint16(sym_json_object),
	1240: uint16(sym_json_access),
	1241: uint16(sym_json_safe_access),
	1242: uint16(sym_json_dot_access),
	1243: uint16(sym_json_safe_dot_access),
	1244: uint16(sym_function_call),
	1245: uint16(sym_binary_expression),
	1246: uint16(sym_unary_expression),
	1247: uint16(sym_ternary_expression),
	1248: uint16(sym_parenthesized_expression),
	1249: uint16(15),
	1250: uint16(3),
	1251: uint16(1),
	1252: uint16(sym_comment),
	1253: uint16(29),
	1254: uint16(1),
	1255: uint16(anon_sym_LPAREN),
	1256: uint16(31),
	1257: uint16(1),
	1258: uint16(anon_sym_LBRACK),
	1259: uint16(35),
	1260: uint16(1),
	1261: uint16(sym_integer),
	1262: uint16(37),
	1263: uint16(1),
	1264: uint16(sym_float),
	1265: uint16(41),
	1266: uint16(1),
	1267: uint16(anon_sym_DQUOTE),
	1268: uint16(43),
	1269: uint16(1),
	1270: uint16(anon_sym_SQUOTE),
	1271: uint16(45),
	1272: uint16(1),
	1273: uint16(anon_sym_BQUOTE),
	1274: uint16(47),
	1275: uint16(1),
	1276: uint16(anon_sym_LBRACE),
	1277: uint16(51),
	1278: uint16(1),
	1279: uint16(sym_ident),
	1280: uint16(64),
	1281: uint16(1),
	1282: uint16(sym_simplexpr),
	1283: uint16(39),
	1284: uint16(2),
	1285: uint16(anon_sym_true),
	1286: uint16(anon_sym_false),
	1287: uint16(51),
	1288: uint16(2),
	1289: uint16(sym_number),
	1290: uint16(sym_boolean),
	1291: uint16(49),
	1292: uint16(3),
	1293: uint16(anon_sym_PLUS),
	1294: uint16(anon_sym_DASH),
	1295: uint16(anon_sym_BANG),
	1296: uint16(43),
	1297: uint16(13),
	1298: uint16(sym_literal),
	1299: uint16(sym_string),
	1300: uint16(sym_json_array),
	1301: uint16(sym_json_object),
	1302: uint16(sym_json_access),
	1303: uint16(sym_json_safe_access),
	1304: uint16(sym_json_dot_access),
	1305: uint16(sym_json_safe_dot_access),
	1306: uint16(sym_function_call),
	1307: uint16(sym_binary_expression),
	1308: uint16(sym_unary_expression),
	1309: uint16(sym_ternary_expression),
	1310: uint16(sym_parenthesized_expression),
	1311: uint16(15),
	1312: uint16(3),
	1313: uint16(1),
	1314: uint16(sym_comment),
	1315: uint16(29),
	1316: uint16(1),
	1317: uint16(anon_sym_LPAREN),
	1318: uint16(31),
	1319: uint16(1),
	1320: uint16(anon_sym_LBRACK),
	1321: uint16(35),
	1322: uint16(1),
	1323: uint16(sym_integer),
	1324: uint16(37),
	1325: uint16(1),
	1326: uint16(sym_float),
	1327: uint16(41),
	1328: uint16(1),
	1329: uint16(anon_sym_DQUOTE),
	1330: uint16(43),
	1331: uint16(1),
	1332: uint16(anon_sym_SQUOTE),
	1333: uint16(45),
	1334: uint16(1),
	1335: uint16(anon_sym_BQUOTE),
	1336: uint16(47),
	1337: uint16(1),
	1338: uint16(anon_sym_LBRACE),
	1339: uint16(51),
	1340: uint16(1),
	1341: uint16(sym_ident),
	1342: uint16(75),
	1343: uint16(1),
	1344: uint16(sym_simplexpr),
	1345: uint16(39),
	1346: uint16(2),
	1347: uint16(anon_sym_true),
	1348: uint16(anon_sym_false),
	1349: uint16(51),
	1350: uint16(2),
	1351: uint16(sym_number),
	1352: uint16(sym_boolean),
	1353: uint16(49),
	1354: uint16(3),
	1355: uint16(anon_sym_PLUS),
	1356: uint16(anon_sym_DASH),
	1357: uint16(anon_sym_BANG),
	1358: uint16(43),
	1359: uint16(13),
	1360: uint16(sym_literal),
	1361: uint16(sym_string),
	1362: uint16(sym_json_array),
	1363: uint16(sym_json_object),
	1364: uint16(sym_json_access),
	1365: uint16(sym_json_safe_access),
	1366: uint16(sym_json_dot_access),
	1367: uint16(sym_json_safe_dot_access),
	1368: uint16(sym_function_call),
	1369: uint16(sym_binary_expression),
	1370: uint16(sym_unary_expression),
	1371: uint16(sym_ternary_expression),
	1372: uint16(sym_parenthesized_expression),
	1373: uint16(15),
	1374: uint16(3),
	1375: uint16(1),
	1376: uint16(sym_comment),
	1377: uint16(29),
	1378: uint16(1),
	1379: uint16(anon_sym_LPAREN),
	1380: uint16(31),
	1381: uint16(1),
	1382: uint16(anon_sym_LBRACK),
	1383: uint16(35),
	1384: uint16(1),
	1385: uint16(sym_integer),
	1386: uint16(37),
	1387: uint16(1),
	1388: uint16(sym_float),
	1389: uint16(41),
	1390: uint16(1),
	1391: uint16(anon_sym_DQUOTE),
	1392: uint16(43),
	1393: uint16(1),
	1394: uint16(anon_sym_SQUOTE),
	1395: uint16(45),
	1396: uint16(1),
	1397: uint16(anon_sym_BQUOTE),
	1398: uint16(47),
	1399: uint16(1),
	1400: uint16(anon_sym_LBRACE),
	1401: uint16(51),
	1402: uint16(1),
	1403: uint16(sym_ident),
	1404: uint16(71),
	1405: uint16(1),
	1406: uint16(sym_simplexpr),
	1407: uint16(39),
	1408: uint16(2),
	1409: uint16(anon_sym_true),
	1410: uint16(anon_sym_false),
	1411: uint16(51),
	1412: uint16(2),
	1413: uint16(sym_number),
	1414: uint16(sym_boolean),
	1415: uint16(49),
	1416: uint16(3),
	1417: uint16(anon_sym_PLUS),
	1418: uint16(anon_sym_DASH),
	1419: uint16(anon_sym_BANG),
	1420: uint16(43),
	1421: uint16(13),
	1422: uint16(sym_literal),
	1423: uint16(sym_string),
	1424: uint16(sym_json_array),
	1425: uint16(sym_json_object),
	1426: uint16(sym_json_access),
	1427: uint16(sym_json_safe_access),
	1428: uint16(sym_json_dot_access),
	1429: uint16(sym_json_safe_dot_access),
	1430: uint16(sym_function_call),
	1431: uint16(sym_binary_expression),
	1432: uint16(sym_unary_expression),
	1433: uint16(sym_ternary_expression),
	1434: uint16(sym_parenthesized_expression),
	1435: uint16(16),
	1436: uint16(3),
	1437: uint16(1),
	1438: uint16(sym_comment),
	1439: uint16(59),
	1440: uint16(1),
	1441: uint16(sym_symbol),
	1442: uint16(62),
	1443: uint16(1),
	1444: uint16(anon_sym_LPAREN),
	1445: uint16(65),
	1446: uint16(1),
	1447: uint16(anon_sym_LBRACK),
	1448: uint16(68),
	1449: uint16(1),
	1450: uint16(sym_keyword),
	1451: uint16(71),
	1452: uint16(1),
	1453: uint16(sym_integer),
	1454: uint16(74),
	1455: uint16(1),
	1456: uint16(sym_float),
	1457: uint16(80),
	1458: uint16(1),
	1459: uint16(anon_sym_DQUOTE),
	1460: uint16(83),
	1461: uint16(1),
	1462: uint16(anon_sym_SQUOTE),
	1463: uint16(86),
	1464: uint16(1),
	1465: uint16(anon_sym_BQUOTE),
	1466: uint16(89),
	1467: uint16(1),
	1468: uint16(anon_sym_LBRACE),
	1469: uint16(77),
	1470: uint16(2),
	1471: uint16(anon_sym_true),
	1472: uint16(anon_sym_false),
	1473: uint16(25),
	1474: uint16(2),
	1475: uint16(sym_ast_block),
	1476: uint16(aux_sym_source_file_repeat1),
	1477: uint16(78),
	1478: uint16(2),
	1479: uint16(sym_number),
	1480: uint16(sym_boolean),
	1481: uint16(57),
	1482: uint16(3),
	1484: uint16(anon_sym_RPAREN),
	1485: uint16(anon_sym_RBRACK),
	1486: uint16(77),
	1487: uint16(6),
	1488: uint16(sym_loop_widget),
	1489: uint16(sym_list),
	1490: uint16(sym_array),
	1491: uint16(sym_literal),
	1492: uint16(sym_string),
	1493: uint16(sym_expr),
	1494: uint16(4),
	1495: uint16(3),
	1496: uint16(1),
	1497: uint16(sym_comment),
	1498: uint16(92),
	1499: uint16(1),
	1500: uint16(anon_sym_LPAREN),
	1501: uint16(96),
	1502: uint16(3),
	1503: uint16(anon_sym_GT),
	1504: uint16(anon_sym_LT),
	1505: uint16(anon_sym_QMARK),
	1506: uint16(94),
	1507: uint16(21),
	1508: uint16(anon_sym_RPAREN),
	1509: uint16(anon_sym_LBRACK),
	1510: uint16(anon_sym_RBRACK),
	1511: uint16(anon_sym_RBRACE),
	1512: uint16(anon_sym_COMMA),
	1513: uint16(anon_sym_COLON),
	1514: uint16(anon_sym_QMARK_DOT),
	1515: uint16(anon_sym_DOT),
	1516: uint16(anon_sym_PLUS),
	1517: uint16(anon_sym_DASH),
	1518: uint16(anon_sym_STAR),
	1519: uint16(anon_sym_SLASH),
	1520: uint16(anon_sym_PERCENT),
	1521: uint16(anon_sym_AMP_AMP),
	1522: uint16(anon_sym_PIPE_PIPE),
	1523: uint16(anon_sym_EQ_EQ),
	1524: uint16(anon_sym_BANG_EQ),
	1525: uint16(anon_sym_EQ_TILDE),
	1526: uint16(anon_sym_GT_EQ),
	1527: uint16(anon_sym_LT_EQ),
	1528: uint16(anon_sym_QMARK_COLON),
	1529: uint16(10),
	1530: uint16(3),
	1531: uint16(1),
	1532: uint16(sym_comment),
	1533: uint16(100),
	1534: uint16(1),
	1535: uint16(anon_sym_LBRACK),
	1536: uint16(102),
	1537: uint16(1),
	1538: uint16(anon_sym_QMARK_DOT),
	1539: uint16(104),
	1540: uint16(1),
	1541: uint16(anon_sym_DOT),
	1542: uint16(114),
	1543: uint16(1),
	1544: uint16(anon_sym_QMARK),
	1545: uint16(106),
	1546: uint16(2),
	1547: uint16(anon_sym_PLUS),
	1548: uint16(anon_sym_DASH),
	1549: uint16(110),
	1550: uint16(2),
	1551: uint16(anon_sym_GT_EQ),
	1552: uint16(anon_sym_LT_EQ),
	1553: uint16(112),
	1554: uint16(2),
	1555: uint16(anon_sym_GT),
	1556: uint16(anon_sym_LT),
	1557: uint16(108),
	1558: uint16(3),
	1559: uint16(anon_sym_STAR),
	1560: uint16(anon_sym_SLASH),
	1561: uint16(anon_sym_PERCENT),
	1562: uint16(98),
	1563: uint16(11),
	1564: uint16(anon_sym_RPAREN),
	1565: uint16(anon_sym_RBRACK),
	1566: uint16(anon_sym_RBRACE),
	1567: uint16(anon_sym_COMMA),
	1568: uint16(anon_sym_COLON),
	1569: uint16(anon_sym_AMP_AMP),
	1570: uint16(anon_sym_PIPE_PIPE),
	1571: uint16(anon_sym_EQ_EQ),
	1572: uint16(anon_sym_BANG_EQ),
	1573: uint16(anon_sym_EQ_TILDE),
	1574: uint16(anon_sym_QMARK_COLON),
	1575: uint16(3),
	1576: uint16(3),
	1577: uint16(1),
	1578: uint16(sym_comment),
	1579: uint16(118),
	1580: uint16(3),
	1581: uint16(anon_sym_GT),
	1582: uint16(anon_sym_LT),
	1583: uint16(anon_sym_QMARK),
	1584: uint16(116),
	1585: uint16(21),
	1586: uint16(anon_sym_RPAREN),
	1587: uint16(anon_sym_LBRACK),
	1588: uint16(anon_sym_RBRACK),
	1589: uint16(anon_sym_RBRACE),
	1590: uint16(anon_sym_COMMA),
	1591: uint16(anon_sym_COLON),
	1592: uint16(anon_sym_QMARK_DOT),
	1593: uint16(anon_sym_DOT),
	1594: uint16(anon_sym_PLUS),
	1595: uint16(anon_sym_DASH),
	1596: uint16(anon_sym_STAR),
	1597: uint16(anon_sym_SLASH),
	1598: uint16(anon_sym_PERCENT),
	1599: uint16(anon_sym_AMP_AMP),
	1600: uint16(anon_sym_PIPE_PIPE),
	1601: uint16(anon_sym_EQ_EQ),
	1602: uint16(anon_sym_BANG_EQ),
	1603: uint16(anon_sym_EQ_TILDE),
	1604: uint16(anon_sym_GT_EQ),
	1605: uint16(anon_sym_LT_EQ),
	1606: uint16(anon_sym_QMARK_COLON),
	1607: uint16(3),
	1608: uint16(3),
	1609: uint16(1),
	1610: uint16(sym_comment),
	1611: uint16(122),
	1612: uint16(3),
	1613: uint16(anon_sym_GT),
	1614: uint16(anon_sym_LT),
	1615: uint16(anon_sym_QMARK),
	1616: uint16(120),
	1617: uint16(21),
	1618: uint16(anon_sym_RPAREN),
	1619: uint16(anon_sym_LBRACK),
	1620: uint16(anon_sym_RBRACK),
	1621: uint16(anon_sym_RBRACE),
	1622: uint16(anon_sym_COMMA),
	1623: uint16(anon_sym_COLON),
	1624: uint16(anon_sym_QMARK_DOT),
	1625: uint16(anon_sym_DOT),
	1626: uint16(anon_sym_PLUS),
	1627: uint16(anon_sym_DASH),
	1628: uint16(anon_sym_STAR),
	1629: uint16(anon_sym_SLASH),
	1630: uint16(anon_sym_PERCENT),
	1631: uint16(anon_sym_AMP_AMP),
	1632: uint16(anon_sym_PIPE_PIPE),
	1633: uint16(anon_sym_EQ_EQ),
	1634: uint16(anon_sym_BANG_EQ),
	1635: uint16(anon_sym_EQ_TILDE),
	1636: uint16(anon_sym_GT_EQ),
	1637: uint16(anon_sym_LT_EQ),
	1638: uint16(anon_sym_QMARK_COLON),
	1639: uint16(17),
	1640: uint16(3),
	1641: uint16(1),
	1642: uint16(sym_comment),
	1643: uint16(7),
	1644: uint16(1),
	1645: uint16(sym_symbol),
	1646: uint16(9),
	1647: uint16(1),
	1648: uint16(anon_sym_LPAREN),
	1649: uint16(11),
	1650: uint16(1),
	1651: uint16(anon_sym_LBRACK),
	1652: uint16(13),
	1653: uint16(1),
	1654: uint16(sym_keyword),
	1655: uint16(15),
	1656: uint16(1),
	1657: uint16(sym_integer),
	1658: uint16(17),
	1659: uint16(1),
	1660: uint16(sym_float),
	1661: uint16(21),
	1662: uint16(1),
	1663: uint16(anon_sym_DQUOTE),
	1664: uint16(23),
	1665: uint16(1),
	1666: uint16(anon_sym_SQUOTE),
	1667: uint16(25),
	1668: uint16(1),
	1669: uint16(anon_sym_BQUOTE),
	1670: uint16(27),
	1671: uint16(1),
	1672: uint16(anon_sym_LBRACE),
	1673: uint16(124),
	1674: uint16(1),
	1675: uint16(anon_sym_for),
	1676: uint16(126),
	1677: uint16(1),
	1678: uint16(anon_sym_RPAREN),
	1679: uint16(19),
	1680: uint16(2),
	1681: uint16(anon_sym_true),
	1682: uint16(anon_sym_false),
	1683: uint16(59),
	1684: uint16(2),
	1685: uint16(sym_ast_block),
	1686: uint16(aux_sym_source_file_repeat1),
	1687: uint16(78),
	1688: uint16(2),
	1689: uint16(sym_number),
	1690: uint16(sym_boolean),
	1691: uint16(77),
	1692: uint16(6),
	1693: uint16(sym_loop_widget),
	1694: uint16(sym_list),
	1695: uint16(sym_array),
	1696: uint16(sym_literal),
	1697: uint16(sym_string),
	1698: uint16(sym_expr),
	1699: uint16(3),
	1700: uint16(3),
	1701: uint16(1),
	1702: uint16(sym_comment),
	1703: uint16(130),
	1704: uint16(3),
	1705: uint16(anon_sym_GT),
	1706: uint16(anon_sym_LT),
	1707: uint16(anon_sym_QMARK),
	1708: uint16(128),
	1709: uint16(21),
	1710: uint16(anon_sym_RPAREN),
	1711: uint16(anon_sym_LBRACK),
	1712: uint16(anon_sym_RBRACK),
	1713: uint16(anon_sym_RBRACE),
	1714: uint16(anon_sym_COMMA),
	1715: uint16(anon_sym_COLON),
	1716: uint16(anon_sym_QMARK_DOT),
	1717: uint16(anon_sym_DOT),
	1718: uint16(anon_sym_PLUS),
	1719: uint16(anon_sym_DASH),
	1720: uint16(anon_sym_STAR),
	1721: uint16(anon_sym_SLASH),
	1722: uint16(anon_sym_PERCENT),
	1723: uint16(anon_sym_AMP_AMP),
	1724: uint16(anon_sym_PIPE_PIPE),
	1725: uint16(anon_sym_EQ_EQ),
	1726: uint16(anon_sym_BANG_EQ),
	1727: uint16(anon_sym_EQ_TILDE),
	1728: uint16(anon_sym_GT_EQ),
	1729: uint16(anon_sym_LT_EQ),
	1730: uint16(anon_sym_QMARK_COLON),
	1731: uint16(3),
	1732: uint16(3),
	1733: uint16(1),
	1734: uint16(sym_comment),
	1735: uint16(134),
	1736: uint16(3),
	1737: uint16(anon_sym_GT),
	1738: uint16(anon_sym_LT),
	1739: uint16(anon_sym_QMARK),
	1740: uint16(132),
	1741: uint16(21),
	1742: uint16(anon_sym_RPAREN),
	1743: uint16(anon_sym_LBRACK),
	1744: uint16(anon_sym_RBRACK),
	1745: uint16(anon_sym_RBRACE),
	1746: uint16(anon_sym_COMMA),
	1747: uint16(anon_sym_COLON),
	1748: uint16(anon_sym_QMARK_DOT),
	1749: uint16(anon_sym_DOT),
	1750: uint16(anon_sym_PLUS),
	1751: uint16(anon_sym_DASH),
	1752: uint16(anon_sym_STAR),
	1753: uint16(anon_sym_SLASH),
	1754: uint16(anon_sym_PERCENT),
	1755: uint16(anon_sym_AMP_AMP),
	1756: uint16(anon_sym_PIPE_PIPE),
	1757: uint16(anon_sym_EQ_EQ),
	1758: uint16(anon_sym_BANG_EQ),
	1759: uint16(anon_sym_EQ_TILDE),
	1760: uint16(anon_sym_GT_EQ),
	1761: uint16(anon_sym_LT_EQ),
	1762: uint16(anon_sym_QMARK_COLON),
	1763: uint16(7),
	1764: uint16(3),
	1765: uint16(1),
	1766: uint16(sym_comment),
	1767: uint16(100),
	1768: uint16(1),
	1769: uint16(anon_sym_LBRACK),
	1770: uint16(102),
	1771: uint16(1),
	1772: uint16(anon_sym_QMARK_DOT),
	1773: uint16(104),
	1774: uint16(1),
	1775: uint16(anon_sym_DOT),
	1776: uint16(108),
	1777: uint16(3),
	1778: uint16(anon_sym_STAR),
	1779: uint16(anon_sym_SLASH),
	1780: uint16(anon_sym_PERCENT),
	1781: uint16(114),
	1782: uint16(3),
	1783: uint16(anon_sym_GT),
	1784: uint16(anon_sym_LT),
	1785: uint16(anon_sym_QMARK),
	1786: uint16(98),
	1787: uint16(15),
	1788: uint16(anon_sym_RPAREN),
	1789: uint16(anon_sym_RBRACK),
	1790: uint16(anon_sym_RBRACE),
	1791: uint16(anon_sym_COMMA),
	1792: uint16(anon_sym_COLON),
	1793: uint16(anon_sym_PLUS),
	1794: uint16(anon_sym_DASH),
	1795: uint16(anon_sym_AMP_AMP),
	1796: uint16(anon_sym_PIPE_PIPE),
	1797: uint16(anon_sym_EQ_EQ),
	1798: uint16(anon_sym_BANG_EQ),
	1799: uint16(anon_sym_EQ_TILDE),
	1800: uint16(anon_sym_GT_EQ),
	1801: uint16(anon_sym_LT_EQ),
	1802: uint16(anon_sym_QMARK_COLON),
	1803: uint16(6),
	1804: uint16(3),
	1805: uint16(1),
	1806: uint16(sym_comment),
	1807: uint16(100),
	1808: uint16(1),
	1809: uint16(anon_sym_LBRACK),
	1810: uint16(102),
	1811: uint16(1),
	1812: uint16(anon_sym_QMARK_DOT),
	1813: uint16(104),
	1814: uint16(1),
	1815: uint16(anon_sym_DOT),
	1816: uint16(114),
	1817: uint16(3),
	1818: uint16(anon_sym_GT),
	1819: uint16(anon_sym_LT),
	1820: uint16(anon_sym_QMARK),
	1821: uint16(98),
	1822: uint16(18),
	1823: uint16(anon_sym_RPAREN),
	1824: uint16(anon_sym_RBRACK),
	1825: uint16(anon_sym_RBRACE),
	1826: uint16(anon_sym_COMMA),
	1827: uint16(anon_sym_COLON),
	1828: uint16(anon_sym_PLUS),
	1829: uint16(anon_sym_DASH),
	1830: uint16(anon_sym_STAR),
	1831: uint16(anon_sym_SLASH),
	1832: uint16(anon_sym_PERCENT),
	1833: uint16(anon_sym_AMP_AMP),
	1834: uint16(anon_sym_PIPE_PIPE),
	1835: uint16(anon_sym_EQ_EQ),
	1836: uint16(anon_sym_BANG_EQ),
	1837: uint16(anon_sym_EQ_TILDE),
	1838: uint16(anon_sym_GT_EQ),
	1839: uint16(anon_sym_LT_EQ),
	1840: uint16(anon_sym_QMARK_COLON),
	1841: uint16(11),
	1842: uint16(3),
	1843: uint16(1),
	1844: uint16(sym_comment),
	1845: uint16(100),
	1846: uint16(1),
	1847: uint16(anon_sym_LBRACK),
	1848: uint16(102),
	1849: uint16(1),
	1850: uint16(anon_sym_QMARK_DOT),
	1851: uint16(104),
	1852: uint16(1),
	1853: uint16(anon_sym_DOT),
	1854: uint16(114),
	1855: uint16(1),
	1856: uint16(anon_sym_QMARK),
	1857: uint16(106),
	1858: uint16(2),
	1859: uint16(anon_sym_PLUS),
	1860: uint16(anon_sym_DASH),
	1861: uint16(110),
	1862: uint16(2),
	1863: uint16(anon_sym_GT_EQ),
	1864: uint16(anon_sym_LT_EQ),
	1865: uint16(112),
	1866: uint16(2),
	1867: uint16(anon_sym_GT),
	1868: uint16(anon_sym_LT),
	1869: uint16(108),
	1870: uint16(3),
	1871: uint16(anon_sym_STAR),
	1872: uint16(anon_sym_SLASH),
	1873: uint16(anon_sym_PERCENT),
	1874: uint16(136),
	1875: uint16(3),
	1876: uint16(anon_sym_EQ_EQ),
	1877: uint16(anon_sym_BANG_EQ),
	1878: uint16(anon_sym_EQ_TILDE),
	1879: uint16(98),
	1880: uint16(8),
	1881: uint16(anon_sym_RPAREN),
	1882: uint16(anon_sym_RBRACK),
	1883: uint16(anon_sym_RBRACE),
	1884: uint16(anon_sym_COMMA),
	1885: uint16(anon_sym_COLON),
	1886: uint16(anon_sym_AMP_AMP),
	1887: uint16(anon_sym_PIPE_PIPE),
	1888: uint16(anon_sym_QMARK_COLON),
	1889: uint16(12),
	1890: uint16(3),
	1891: uint16(1),
	1892: uint16(sym_comment),
	1893: uint16(100),
	1894: uint16(1),
	1895: uint16(anon_sym_LBRACK),
	1896: uint16(102),
	1897: uint16(1),
	1898: uint16(anon_sym_QMARK_DOT),
	1899: uint16(104),
	1900: uint16(1),
	1901: uint16(anon_sym_DOT),
	1902: uint16(114),
	1903: uint16(1),
	1904: uint16(anon_sym_QMARK),
	1905: uint16(106),
	1906: uint16(2),
	1907: uint16(anon_sym_PLUS),
	1908: uint16(anon_sym_DASH),
	1909: uint16(110),
	1910: uint16(2),
	1911: uint16(anon_sym_GT_EQ),
	1912: uint16(anon_sym_LT_EQ),
	1913: uint16(112),
	1914: uint16(2),
	1915: uint16(anon_sym_GT),
	1916: uint16(anon_sym_LT),
	1917: uint16(138),
	1918: uint16(2),
	1919: uint16(anon_sym_AMP_AMP),
	1920: uint16(anon_sym_QMARK_COLON),
	1921: uint16(108),
	1922: uint16(3),
	1923: uint16(anon_sym_STAR),
	1924: uint16(anon_sym_SLASH),
	1925: uint16(anon_sym_PERCENT),
	1926: uint16(136),
	1927: uint16(3),
	1928: uint16(anon_sym_EQ_EQ),
	1929: uint16(anon_sym_BANG_EQ),
	1930: uint16(anon_sym_EQ_TILDE),
	1931: uint16(98),
	1932: uint16(6),
	1933: uint16(anon_sym_RPAREN),
	1934: uint16(anon_sym_RBRACK),
	1935: uint16(anon_sym_RBRACE),
	1936: uint16(anon_sym_COMMA),
	1937: uint16(anon_sym_COLON),
	1938: uint16(anon_sym_PIPE_PIPE),
	1939: uint16(8),
	1940: uint16(3),
	1941: uint16(1),
	1942: uint16(sym_comment),
	1943: uint16(100),
	1944: uint16(1),
	1945: uint16(anon_sym_LBRACK),
	1946: uint16(102),
	1947: uint16(1),
	1948: uint16(anon_sym_QMARK_DOT),
	1949: uint16(104),
	1950: uint16(1),
	1951: uint16(anon_sym_DOT),
	1952: uint16(106),
	1953: uint16(2),
	1954: uint16(anon_sym_PLUS),
	1955: uint16(anon_sym_DASH),
	1956: uint16(108),
	1957: uint16(3),
	1958: uint16(anon_sym_STAR),
	1959: uint16(anon_sym_SLASH),
	1960: uint16(anon_sym_PERCENT),
	1961: uint16(114),
	1962: uint16(3),
	1963: uint16(anon_sym_GT),
	1964: uint16(anon_sym_LT),
	1965: uint16(anon_sym_QMARK),
	1966: uint16(98),
	1967: uint16(13),
	1968: uint16(anon_sym_RPAREN),
	1969: uint16(anon_sym_RBRACK),
	1970: uint16(anon_sym_RBRACE),
	1971: uint16(anon_sym_COMMA),
	1972: uint16(anon_sym_COLON),
	1973: uint16(anon_sym_AMP_AMP),
	1974: uint16(anon_sym_PIPE_PIPE),
	1975: uint16(anon_sym_EQ_EQ),
	1976: uint16(anon_sym_BANG_EQ),
	1977: uint16(anon_sym_EQ_TILDE),
	1978: uint16(anon_sym_GT_EQ),
	1979: uint16(anon_sym_LT_EQ),
	1980: uint16(anon_sym_QMARK_COLON),
	1981: uint16(3),
	1982: uint16(3),
	1983: uint16(1),
	1984: uint16(sym_comment),
	1985: uint16(142),
	1986: uint16(3),
	1987: uint16(anon_sym_GT),
	1988: uint16(anon_sym_LT),
	1989: uint16(anon_sym_QMARK),
	1990: uint16(140),
	1991: uint16(21),
	1992: uint16(anon_sym_RPAREN),
	1993: uint16(anon_sym_LBRACK),
	1994: uint16(anon_sym_RBRACK),
	1995: uint16(anon_sym_RBRACE),
	1996: uint16(anon_sym_COMMA),
	1997: uint16(anon_sym_COLON),
	1998: uint16(anon_sym_QMARK_DOT),
	1999: uint16(anon_sym_DOT),
	2000: uint16(anon_sym_PLUS),
	2001: uint16(anon_sym_DASH),
	2002: uint16(anon_sym_STAR),
	2003: uint16(anon_sym_SLASH),
	2004: uint16(anon_sym_PERCENT),
	2005: uint16(anon_sym_AMP_AMP),
	2006: uint16(anon_sym_PIPE_PIPE),
	2007: uint16(anon_sym_EQ_EQ),
	2008: uint16(anon_sym_BANG_EQ),
	2009: uint16(anon_sym_EQ_TILDE),
	2010: uint16(anon_sym_GT_EQ),
	2011: uint16(anon_sym_LT_EQ),
	2012: uint16(anon_sym_QMARK_COLON),
	2013: uint16(3),
	2014: uint16(3),
	2015: uint16(1),
	2016: uint16(sym_comment),
	2017: uint16(146),
	2018: uint16(3),
	2019: uint16(anon_sym_GT),
	2020: uint16(anon_sym_LT),
	2021: uint16(anon_sym_QMARK),
	2022: uint16(144),
	2023: uint16(21),
	2024: uint16(anon_sym_RPAREN),
	2025: uint16(anon_sym_LBRACK),
	2026: uint16(anon_sym_RBRACK),
	2027: uint16(anon_sym_RBRACE),
	2028: uint16(anon_sym_COMMA),
	2029: uint16(anon_sym_COLON),
	2030: uint16(anon_sym_QMARK_DOT),
	2031: uint16(anon_sym_DOT),
	2032: uint16(anon_sym_PLUS),
	2033: uint16(anon_sym_DASH),
	2034: uint16(anon_sym_STAR),
	2035: uint16(anon_sym_SLASH),
	2036: uint16(anon_sym_PERCENT),
	2037: uint16(anon_sym_AMP_AMP),
	2038: uint16(anon_sym_PIPE_PIPE),
	2039: uint16(anon_sym_EQ_EQ),
	2040: uint16(anon_sym_BANG_EQ),
	2041: uint16(anon_sym_EQ_TILDE),
	2042: uint16(anon_sym_GT_EQ),
	2043: uint16(anon_sym_LT_EQ),
	2044: uint16(anon_sym_QMARK_COLON),
	2045: uint16(3),
	2046: uint16(3),
	2047: uint16(1),
	2048: uint16(sym_comment),
	2049: uint16(150),
	2050: uint16(3),
	2051: uint16(anon_sym_GT),
	2052: uint16(anon_sym_LT),
	2053: uint16(anon_sym_QMARK),
	2054: uint16(148),
	2055: uint16(21),
	2056: uint16(anon_sym_RPAREN),
	2057: uint16(anon_sym_LBRACK),
	2058: uint16(anon_sym_RBRACK),
	2059: uint16(anon_sym_RBRACE),
	2060: uint16(anon_sym_COMMA),
	2061: uint16(anon_sym_COLON),
	2062: uint16(anon_sym_QMARK_DOT),
	2063: uint16(anon_sym_DOT),
	2064: uint16(anon_sym_PLUS),
	2065: uint16(anon_sym_DASH),
	2066: uint16(anon_sym_STAR),
	2067: uint16(anon_sym_SLASH),
	2068: uint16(anon_sym_PERCENT),
	2069: uint16(anon_sym_AMP_AMP),
	2070: uint16(anon_sym_PIPE_PIPE),
	2071: uint16(anon_sym_EQ_EQ),
	2072: uint16(anon_sym_BANG_EQ),
	2073: uint16(anon_sym_EQ_TILDE),
	2074: uint16(anon_sym_GT_EQ),
	2075: uint16(anon_sym_LT_EQ),
	2076: uint16(anon_sym_QMARK_COLON),
	2077: uint16(6),
	2078: uint16(3),
	2079: uint16(1),
	2080: uint16(sym_comment),
	2081: uint16(100),
	2082: uint16(1),
	2083: uint16(anon_sym_LBRACK),
	2084: uint16(102),
	2085: uint16(1),
	2086: uint16(anon_sym_QMARK_DOT),
	2087: uint16(104),
	2088: uint16(1),
	2089: uint16(anon_sym_DOT),
	2090: uint16(154),
	2091: uint16(3),
	2092: uint16(anon_sym_GT),
	2093: uint16(anon_sym_LT),
	2094: uint16(anon_sym_QMARK),
	2095: uint16(152),
	2096: uint16(18),
	2097: uint16(anon_sym_RPAREN),
	2098: uint16(anon_sym_RBRACK),
	2099: uint16(anon_sym_RBRACE),
	2100: uint16(anon_sym_COMMA),
	2101: uint16(anon_sym_COLON),
	2102: uint16(anon_sym_PLUS),
	2103: uint16(anon_sym_DASH),
	2104: uint16(anon_sym_STAR),
	2105: uint16(anon_sym_SLASH),
	2106: uint16(anon_sym_PERCENT),
	2107: uint16(anon_sym_AMP_AMP),
	2108: uint16(anon_sym_PIPE_PIPE),
	2109: uint16(anon_sym_EQ_EQ),
	2110: uint16(anon_sym_BANG_EQ),
	2111: uint16(anon_sym_EQ_TILDE),
	2112: uint16(anon_sym_GT_EQ),
	2113: uint16(anon_sym_LT_EQ),
	2114: uint16(anon_sym_QMARK_COLON),
	2115: uint16(3),
	2116: uint16(3),
	2117: uint16(1),
	2118: uint16(sym_comment),
	2119: uint16(158),
	2120: uint16(3),
	2121: uint16(anon_sym_GT),
	2122: uint16(anon_sym_LT),
	2123: uint16(anon_sym_QMARK),
	2124: uint16(156),
	2125: uint16(21),
	2126: uint16(anon_sym_RPAREN),
	2127: uint16(anon_sym_LBRACK),
	2128: uint16(anon_sym_RBRACK),
	2129: uint16(anon_sym_RBRACE),
	2130: uint16(anon_sym_COMMA),
	2131: uint16(anon_sym_COLON),
	2132: uint16(anon_sym_QMARK_DOT),
	2133: uint16(anon_sym_DOT),
	2134: uint16(anon_sym_PLUS),
	2135: uint16(anon_sym_DASH),
	2136: uint16(anon_sym_STAR),
	2137: uint16(anon_sym_SLASH),
	2138: uint16(anon_sym_PERCENT),
	2139: uint16(anon_sym_AMP_AMP),
	2140: uint16(anon_sym_PIPE_PIPE),
	2141: uint16(anon_sym_EQ_EQ),
	2142: uint16(anon_sym_BANG_EQ),
	2143: uint16(anon_sym_EQ_TILDE),
	2144: uint16(anon_sym_GT_EQ),
	2145: uint16(anon_sym_LT_EQ),
	2146: uint16(anon_sym_QMARK_COLON),
	2147: uint16(3),
	2148: uint16(3),
	2149: uint16(1),
	2150: uint16(sym_comment),
	2151: uint16(96),
	2152: uint16(3),
	2153: uint16(anon_sym_GT),
	2154: uint16(anon_sym_LT),
	2155: uint16(anon_sym_QMARK),
	2156: uint16(94),
	2157: uint16(21),
	2158: uint16(anon_sym_RPAREN),
	2159: uint16(anon_sym_LBRACK),
	2160: uint16(anon_sym_RBRACK),
	2161: uint16(anon_sym_RBRACE),
	2162: uint16(anon_sym_COMMA),
	2163: uint16(anon_sym_COLON),
	2164: uint16(anon_sym_QMARK_DOT),
	2165: uint16(anon_sym_DOT),
	2166: uint16(anon_sym_PLUS),
	2167: uint16(anon_sym_DASH),
	2168: uint16(anon_sym_STAR),
	2169: uint16(anon_sym_SLASH),
	2170: uint16(anon_sym_PERCENT),
	2171: uint16(anon_sym_AMP_AMP),
	2172: uint16(anon_sym_PIPE_PIPE),
	2173: uint16(anon_sym_EQ_EQ),
	2174: uint16(anon_sym_BANG_EQ),
	2175: uint16(anon_sym_EQ_TILDE),
	2176: uint16(anon_sym_GT_EQ),
	2177: uint16(anon_sym_LT_EQ),
	2178: uint16(anon_sym_QMARK_COLON),
	2179: uint16(3),
	2180: uint16(3),
	2181: uint16(1),
	2182: uint16(sym_comment),
	2183: uint16(162),
	2184: uint16(3),
	2185: uint16(anon_sym_GT),
	2186: uint16(anon_sym_LT),
	2187: uint16(anon_sym_QMARK),
	2188: uint16(160),
	2189: uint16(21),
	2190: uint16(anon_sym_RPAREN),
	2191: uint16(anon_sym_LBRACK),
	2192: uint16(anon_sym_RBRACK),
	2193: uint16(anon_sym_RBRACE),
	2194: uint16(anon_sym_COMMA),
	2195: uint16(anon_sym_COLON),
	2196: uint16(anon_sym_QMARK_DOT),
	2197: uint16(anon_sym_DOT),
	2198: uint16(anon_sym_PLUS),
	2199: uint16(anon_sym_DASH),
	2200: uint16(anon_sym_STAR),
	2201: uint16(anon_sym_SLASH),
	2202: uint16(anon_sym_PERCENT),
	2203: uint16(anon_sym_AMP_AMP),
	2204: uint16(anon_sym_PIPE_PIPE),
	2205: uint16(anon_sym_EQ_EQ),
	2206: uint16(anon_sym_BANG_EQ),
	2207: uint16(anon_sym_EQ_TILDE),
	2208: uint16(anon_sym_GT_EQ),
	2209: uint16(anon_sym_LT_EQ),
	2210: uint16(anon_sym_QMARK_COLON),
	2211: uint16(3),
	2212: uint16(3),
	2213: uint16(1),
	2214: uint16(sym_comment),
	2215: uint16(166),
	2216: uint16(3),
	2217: uint16(anon_sym_GT),
	2218: uint16(anon_sym_LT),
	2219: uint16(anon_sym_QMARK),
	2220: uint16(164),
	2221: uint16(21),
	2222: uint16(anon_sym_RPAREN),
	2223: uint16(anon_sym_LBRACK),
	2224: uint16(anon_sym_RBRACK),
	2225: uint16(anon_sym_RBRACE),
	2226: uint16(anon_sym_COMMA),
	2227: uint16(anon_sym_COLON),
	2228: uint16(anon_sym_QMARK_DOT),
	2229: uint16(anon_sym_DOT),
	2230: uint16(anon_sym_PLUS),
	2231: uint16(anon_sym_DASH),
	2232: uint16(anon_sym_STAR),
	2233: uint16(anon_sym_SLASH),
	2234: uint16(anon_sym_PERCENT),
	2235: uint16(anon_sym_AMP_AMP),
	2236: uint16(anon_sym_PIPE_PIPE),
	2237: uint16(anon_sym_EQ_EQ),
	2238: uint16(anon_sym_BANG_EQ),
	2239: uint16(anon_sym_EQ_TILDE),
	2240: uint16(anon_sym_GT_EQ),
	2241: uint16(anon_sym_LT_EQ),
	2242: uint16(anon_sym_QMARK_COLON),
	2243: uint16(13),
	2244: uint16(3),
	2245: uint16(1),
	2246: uint16(sym_comment),
	2247: uint16(100),
	2248: uint16(1),
	2249: uint16(anon_sym_LBRACK),
	2250: uint16(102),
	2251: uint16(1),
	2252: uint16(anon_sym_QMARK_DOT),
	2253: uint16(104),
	2254: uint16(1),
	2255: uint16(anon_sym_DOT),
	2256: uint16(170),
	2257: uint16(1),
	2258: uint16(anon_sym_PIPE_PIPE),
	2259: uint16(172),
	2260: uint16(1),
	2261: uint16(anon_sym_QMARK),
	2262: uint16(106),
	2263: uint16(2),
	2264: uint16(anon_sym_PLUS),
	2265: uint16(anon_sym_DASH),
	2266: uint16(110),
	2267: uint16(2),
	2268: uint16(anon_sym_GT_EQ),
	2269: uint16(anon_sym_LT_EQ),
	2270: uint16(112),
	2271: uint16(2),
	2272: uint16(anon_sym_GT),
	2273: uint16(anon_sym_LT),
	2274: uint16(138),
	2275: uint16(2),
	2276: uint16(anon_sym_AMP_AMP),
	2277: uint16(anon_sym_QMARK_COLON),
	2278: uint16(108),
	2279: uint16(3),
	2280: uint16(anon_sym_STAR),
	2281: uint16(anon_sym_SLASH),
	2282: uint16(anon_sym_PERCENT),
	2283: uint16(136),
	2284: uint16(3),
	2285: uint16(anon_sym_EQ_EQ),
	2286: uint16(anon_sym_BANG_EQ),
	2287: uint16(anon_sym_EQ_TILDE),
	2288: uint16(168),
	2289: uint16(5),
	2290: uint16(anon_sym_RPAREN),
	2291: uint16(anon_sym_RBRACK),
	2292: uint16(anon_sym_RBRACE),
	2293: uint16(anon_sym_COMMA),
	2294: uint16(anon_sym_COLON),
	2295: uint16(3),
	2296: uint16(3),
	2297: uint16(1),
	2298: uint16(sym_comment),
	2299: uint16(176),
	2300: uint16(3),
	2301: uint16(anon_sym_GT),
	2302: uint16(anon_sym_LT),
	2303: uint16(anon_sym_QMARK),
	2304: uint16(174),
	2305: uint16(21),
	2306: uint16(anon_sym_RPAREN),
	2307: uint16(anon_sym_LBRACK),
	2308: uint16(anon_sym_RBRACK),
	2309: uint16(anon_sym_RBRACE),
	2310: uint16(anon_sym_COMMA),
	2311: uint16(anon_sym_COLON),
	2312: uint16(anon_sym_QMARK_DOT),
	2313: uint16(anon_sym_DOT),
	2314: uint16(anon_sym_PLUS),
	2315: uint16(anon_sym_DASH),
	2316: uint16(anon_sym_STAR),
	2317: uint16(anon_sym_SLASH),
	2318: uint16(anon_sym_PERCENT),
	2319: uint16(anon_sym_AMP_AMP),
	2320: uint16(anon_sym_PIPE_PIPE),
	2321: uint16(anon_sym_EQ_EQ),
	2322: uint16(anon_sym_BANG_EQ),
	2323: uint16(anon_sym_EQ_TILDE),
	2324: uint16(anon_sym_GT_EQ),
	2325: uint16(anon_sym_LT_EQ),
	2326: uint16(anon_sym_QMARK_COLON),
	2327: uint16(3),
	2328: uint16(3),
	2329: uint16(1),
	2330: uint16(sym_comment),
	2331: uint16(180),
	2332: uint16(3),
	2333: uint16(anon_sym_GT),
	2334: uint16(anon_sym_LT),
	2335: uint16(anon_sym_QMARK),
	2336: uint16(178),
	2337: uint16(21),
	2338: uint16(anon_sym_RPAREN),
	2339: uint16(anon_sym_LBRACK),
	2340: uint16(anon_sym_RBRACK),
	2341: uint16(anon_sym_RBRACE),
	2342: uint16(anon_sym_COMMA),
	2343: uint16(anon_sym_COLON),
	2344: uint16(anon_sym_QMARK_DOT),
	2345: uint16(anon_sym_DOT),
	2346: uint16(anon_sym_PLUS),
	2347: uint16(anon_sym_DASH),
	2348: uint16(anon_sym_STAR),
	2349: uint16(anon_sym_SLASH),
	2350: uint16(anon_sym_PERCENT),
	2351: uint16(anon_sym_AMP_AMP),
	2352: uint16(anon_sym_PIPE_PIPE),
	2353: uint16(anon_sym_EQ_EQ),
	2354: uint16(anon_sym_BANG_EQ),
	2355: uint16(anon_sym_EQ_TILDE),
	2356: uint16(anon_sym_GT_EQ),
	2357: uint16(anon_sym_LT_EQ),
	2358: uint16(anon_sym_QMARK_COLON),
	2359: uint16(3),
	2360: uint16(3),
	2361: uint16(1),
	2362: uint16(sym_comment),
	2363: uint16(184),
	2364: uint16(3),
	2365: uint16(anon_sym_GT),
	2366: uint16(anon_sym_LT),
	2367: uint16(anon_sym_QMARK),
	2368: uint16(182),
	2369: uint16(21),
	2370: uint16(anon_sym_RPAREN),
	2371: uint16(anon_sym_LBRACK),
	2372: uint16(anon_sym_RBRACK),
	2373: uint16(anon_sym_RBRACE),
	2374: uint16(anon_sym_COMMA),
	2375: uint16(anon_sym_COLON),
	2376: uint16(anon_sym_QMARK_DOT),
	2377: uint16(anon_sym_DOT),
	2378: uint16(anon_sym_PLUS),
	2379: uint16(anon_sym_DASH),
	2380: uint16(anon_sym_STAR),
	2381: uint16(anon_sym_SLASH),
	2382: uint16(anon_sym_PERCENT),
	2383: uint16(anon_sym_AMP_AMP),
	2384: uint16(anon_sym_PIPE_PIPE),
	2385: uint16(anon_sym_EQ_EQ),
	2386: uint16(anon_sym_BANG_EQ),
	2387: uint16(anon_sym_EQ_TILDE),
	2388: uint16(anon_sym_GT_EQ),
	2389: uint16(anon_sym_LT_EQ),
	2390: uint16(anon_sym_QMARK_COLON),
	2391: uint16(3),
	2392: uint16(3),
	2393: uint16(1),
	2394: uint16(sym_comment),
	2395: uint16(188),
	2396: uint16(3),
	2397: uint16(anon_sym_GT),
	2398: uint16(anon_sym_LT),
	2399: uint16(anon_sym_QMARK),
	2400: uint16(186),
	2401: uint16(21),
	2402: uint16(anon_sym_RPAREN),
	2403: uint16(anon_sym_LBRACK),
	2404: uint16(anon_sym_RBRACK),
	2405: uint16(anon_sym_RBRACE),
	2406: uint16(anon_sym_COMMA),
	2407: uint16(anon_sym_COLON),
	2408: uint16(anon_sym_QMARK_DOT),
	2409: uint16(anon_sym_DOT),
	2410: uint16(anon_sym_PLUS),
	2411: uint16(anon_sym_DASH),
	2412: uint16(anon_sym_STAR),
	2413: uint16(anon_sym_SLASH),
	2414: uint16(anon_sym_PERCENT),
	2415: uint16(anon_sym_AMP_AMP),
	2416: uint16(anon_sym_PIPE_PIPE),
	2417: uint16(anon_sym_EQ_EQ),
	2418: uint16(anon_sym_BANG_EQ),
	2419: uint16(anon_sym_EQ_TILDE),
	2420: uint16(anon_sym_GT_EQ),
	2421: uint16(anon_sym_LT_EQ),
	2422: uint16(anon_sym_QMARK_COLON),
	2423: uint16(3),
	2424: uint16(3),
	2425: uint16(1),
	2426: uint16(sym_comment),
	2427: uint16(192),
	2428: uint16(3),
	2429: uint16(anon_sym_GT),
	2430: uint16(anon_sym_LT),
	2431: uint16(anon_sym_QMARK),
	2432: uint16(190),
	2433: uint16(21),
	2434: uint16(anon_sym_RPAREN),
	2435: uint16(anon_sym_LBRACK),
	2436: uint16(anon_sym_RBRACK),
	2437: uint16(anon_sym_RBRACE),
	2438: uint16(anon_sym_COMMA),
	2439: uint16(anon_sym_COLON),
	2440: uint16(anon_sym_QMARK_DOT),
	2441: uint16(anon_sym_DOT),
	2442: uint16(anon_sym_PLUS),
	2443: uint16(anon_sym_DASH),
	2444: uint16(anon_sym_STAR),
	2445: uint16(anon_sym_SLASH),
	2446: uint16(anon_sym_PERCENT),
	2447: uint16(anon_sym_AMP_AMP),
	2448: uint16(anon_sym_PIPE_PIPE),
	2449: uint16(anon_sym_EQ_EQ),
	2450: uint16(anon_sym_BANG_EQ),
	2451: uint16(anon_sym_EQ_TILDE),
	2452: uint16(anon_sym_GT_EQ),
	2453: uint16(anon_sym_LT_EQ),
	2454: uint16(anon_sym_QMARK_COLON),
	2455: uint16(3),
	2456: uint16(3),
	2457: uint16(1),
	2458: uint16(sym_comment),
	2459: uint16(196),
	2460: uint16(3),
	2461: uint16(anon_sym_GT),
	2462: uint16(anon_sym_LT),
	2463: uint16(anon_sym_QMARK),
	2464: uint16(194),
	2465: uint16(21),
	2466: uint16(anon_sym_RPAREN),
	2467: uint16(anon_sym_LBRACK),
	2468: uint16(anon_sym_RBRACK),
	2469: uint16(anon_sym_RBRACE),
	2470: uint16(anon_sym_COMMA),
	2471: uint16(anon_sym_COLON),
	2472: uint16(anon_sym_QMARK_DOT),
	2473: uint16(anon_sym_DOT),
	2474: uint16(anon_sym_PLUS),
	2475: uint16(anon_sym_DASH),
	2476: uint16(anon_sym_STAR),
	2477: uint16(anon_sym_SLASH),
	2478: uint16(anon_sym_PERCENT),
	2479: uint16(anon_sym_AMP_AMP),
	2480: uint16(anon_sym_PIPE_PIPE),
	2481: uint16(anon_sym_EQ_EQ),
	2482: uint16(anon_sym_BANG_EQ),
	2483: uint16(anon_sym_EQ_TILDE),
	2484: uint16(anon_sym_GT_EQ),
	2485: uint16(anon_sym_LT_EQ),
	2486: uint16(anon_sym_QMARK_COLON),
	2487: uint16(3),
	2488: uint16(3),
	2489: uint16(1),
	2490: uint16(sym_comment),
	2491: uint16(200),
	2492: uint16(3),
	2493: uint16(anon_sym_GT),
	2494: uint16(anon_sym_LT),
	2495: uint16(anon_sym_QMARK),
	2496: uint16(198),
	2497: uint16(21),
	2498: uint16(anon_sym_RPAREN),
	2499: uint16(anon_sym_LBRACK),
	2500: uint16(anon_sym_RBRACK),
	2501: uint16(anon_sym_RBRACE),
	2502: uint16(anon_sym_COMMA),
	2503: uint16(anon_sym_COLON),
	2504: uint16(anon_sym_QMARK_DOT),
	2505: uint16(anon_sym_DOT),
	2506: uint16(anon_sym_PLUS),
	2507: uint16(anon_sym_DASH),
	2508: uint16(anon_sym_STAR),
	2509: uint16(anon_sym_SLASH),
	2510: uint16(anon_sym_PERCENT),
	2511: uint16(anon_sym_AMP_AMP),
	2512: uint16(anon_sym_PIPE_PIPE),
	2513: uint16(anon_sym_EQ_EQ),
	2514: uint16(anon_sym_BANG_EQ),
	2515: uint16(anon_sym_EQ_TILDE),
	2516: uint16(anon_sym_GT_EQ),
	2517: uint16(anon_sym_LT_EQ),
	2518: uint16(anon_sym_QMARK_COLON),
	2519: uint16(3),
	2520: uint16(3),
	2521: uint16(1),
	2522: uint16(sym_comment),
	2523: uint16(204),
	2524: uint16(3),
	2525: uint16(anon_sym_GT),
	2526: uint16(anon_sym_LT),
	2527: uint16(anon_sym_QMARK),
	2528: uint16(202),
	2529: uint16(21),
	2530: uint16(anon_sym_RPAREN),
	2531: uint16(anon_sym_LBRACK),
	2532: uint16(anon_sym_RBRACK),
	2533: uint16(anon_sym_RBRACE),
	2534: uint16(anon_sym_COMMA),
	2535: uint16(anon_sym_COLON),
	2536: uint16(anon_sym_QMARK_DOT),
	2537: uint16(anon_sym_DOT),
	2538: uint16(anon_sym_PLUS),
	2539: uint16(anon_sym_DASH),
	2540: uint16(anon_sym_STAR),
	2541: uint16(anon_sym_SLASH),
	2542: uint16(anon_sym_PERCENT),
	2543: uint16(anon_sym_AMP_AMP),
	2544: uint16(anon_sym_PIPE_PIPE),
	2545: uint16(anon_sym_EQ_EQ),
	2546: uint16(anon_sym_BANG_EQ),
	2547: uint16(anon_sym_EQ_TILDE),
	2548: uint16(anon_sym_GT_EQ),
	2549: uint16(anon_sym_LT_EQ),
	2550: uint16(anon_sym_QMARK_COLON),
	2551: uint16(3),
	2552: uint16(3),
	2553: uint16(1),
	2554: uint16(sym_comment),
	2555: uint16(208),
	2556: uint16(3),
	2557: uint16(anon_sym_GT),
	2558: uint16(anon_sym_LT),
	2559: uint16(anon_sym_QMARK),
	2560: uint16(206),
	2561: uint16(21),
	2562: uint16(anon_sym_RPAREN),
	2563: uint16(anon_sym_LBRACK),
	2564: uint16(anon_sym_RBRACK),
	2565: uint16(anon_sym_RBRACE),
	2566: uint16(anon_sym_COMMA),
	2567: uint16(anon_sym_COLON),
	2568: uint16(anon_sym_QMARK_DOT),
	2569: uint16(anon_sym_DOT),
	2570: uint16(anon_sym_PLUS),
	2571: uint16(anon_sym_DASH),
	2572: uint16(anon_sym_STAR),
	2573: uint16(anon_sym_SLASH),
	2574: uint16(anon_sym_PERCENT),
	2575: uint16(anon_sym_AMP_AMP),
	2576: uint16(anon_sym_PIPE_PIPE),
	2577: uint16(anon_sym_EQ_EQ),
	2578: uint16(anon_sym_BANG_EQ),
	2579: uint16(anon_sym_EQ_TILDE),
	2580: uint16(anon_sym_GT_EQ),
	2581: uint16(anon_sym_LT_EQ),
	2582: uint16(anon_sym_QMARK_COLON),
	2583: uint16(16),
	2584: uint16(3),
	2585: uint16(1),
	2586: uint16(sym_comment),
	2587: uint16(7),
	2588: uint16(1),
	2589: uint16(sym_symbol),
	2590: uint16(9),
	2591: uint16(1),
	2592: uint16(anon_sym_LPAREN),
	2593: uint16(11),
	2594: uint16(1),
	2595: uint16(anon_sym_LBRACK),
	2596: uint16(13),
	2597: uint16(1),
	2598: uint16(sym_keyword),
	2599: uint16(15),
	2600: uint16(1),
	2601: uint16(sym_integer),
	2602: uint16(17),
	2603: uint16(1),
	2604: uint16(sym_float),
	2605: uint16(21),
	2606: uint16(1),
	2607: uint16(anon_sym_DQUOTE),
	2608: uint16(23),
	2609: uint16(1),
	2610: uint16(anon_sym_SQUOTE),
	2611: uint16(25),
	2612: uint16(1),
	2613: uint16(anon_sym_BQUOTE),
	2614: uint16(27),
	2615: uint16(1),
	2616: uint16(anon_sym_LBRACE),
	2617: uint16(210),
	2618: uint16(1),
	2619: uint16(anon_sym_RBRACK),
	2620: uint16(19),
	2621: uint16(2),
	2622: uint16(anon_sym_true),
	2623: uint16(anon_sym_false),
	2624: uint16(25),
	2625: uint16(2),
	2626: uint16(sym_ast_block),
	2627: uint16(aux_sym_source_file_repeat1),
	2628: uint16(78),
	2629: uint16(2),
	2630: uint16(sym_number),
	2631: uint16(sym_boolean),
	2632: uint16(77),
	2633: uint16(6),
	2634: uint16(sym_loop_widget),
	2635: uint16(sym_list),
	2636: uint16(sym_array),
	2637: uint16(sym_literal),
	2638: uint16(sym_string),
	2639: uint16(sym_expr),
	2640: uint16(16),
	2641: uint16(3),
	2642: uint16(1),
	2643: uint16(sym_comment),
	2644: uint16(7),
	2645: uint16(1),
	2646: uint16(sym_symbol),
	2647: uint16(9),
	2648: uint16(1),
	2649: uint16(anon_sym_LPAREN),
	2650: uint16(11),
	2651: uint16(1),
	2652: uint16(anon_sym_LBRACK),
	2653: uint16(13),
	2654: uint16(1),
	2655: uint16(sym_keyword),
	2656: uint16(15),
	2657: uint16(1),
	2658: uint16(sym_integer),
	2659: uint16(17),
	2660: uint16(1),
	2661: uint16(sym_float),
	2662: uint16(21),
	2663: uint16(1),
	2664: uint16(anon_sym_DQUOTE),
	2665: uint16(23),
	2666: uint16(1),
	2667: uint16(anon_sym_SQUOTE),
	2668: uint16(25),
	2669: uint16(1),
	2670: uint16(anon_sym_BQUOTE),
	2671: uint16(27),
	2672: uint16(1),
	2673: uint16(anon_sym_LBRACE),
	2674: uint16(212),
	2675: uint16(1),
	2677: uint16(19),
	2678: uint16(2),
	2679: uint16(anon_sym_true),
	2680: uint16(anon_sym_false),
	2681: uint16(25),
	2682: uint16(2),
	2683: uint16(sym_ast_block),
	2684: uint16(aux_sym_source_file_repeat1),
	2685: uint16(78),
	2686: uint16(2),
	2687: uint16(sym_number),
	2688: uint16(sym_boolean),
	2689: uint16(77),
	2690: uint16(6),
	2691: uint16(sym_loop_widget),
	2692: uint16(sym_list),
	2693: uint16(sym_array),
	2694: uint16(sym_literal),
	2695: uint16(sym_string),
	2696: uint16(sym_expr),
	2697: uint16(16),
	2698: uint16(3),
	2699: uint16(1),
	2700: uint16(sym_comment),
	2701: uint16(7),
	2702: uint16(1),
	2703: uint16(sym_symbol),
	2704: uint16(9),
	2705: uint16(1),
	2706: uint16(anon_sym_LPAREN),
	2707: uint16(11),
	2708: uint16(1),
	2709: uint16(anon_sym_LBRACK),
	2710: uint16(13),
	2711: uint16(1),
	2712: uint16(sym_keyword),
	2713: uint16(15),
	2714: uint16(1),
	2715: uint16(sym_integer),
	2716: uint16(17),
	2717: uint16(1),
	2718: uint16(sym_float),
	2719: uint16(21),
	2720: uint16(1),
	2721: uint16(anon_sym_DQUOTE),
	2722: uint16(23),
	2723: uint16(1),
	2724: uint16(anon_sym_SQUOTE),
	2725: uint16(25),
	2726: uint16(1),
	2727: uint16(anon_sym_BQUOTE),
	2728: uint16(27),
	2729: uint16(1),
	2730: uint16(anon_sym_LBRACE),
	2731: uint16(214),
	2732: uint16(1),
	2733: uint16(anon_sym_RBRACK),
	2734: uint16(19),
	2735: uint16(2),
	2736: uint16(anon_sym_true),
	2737: uint16(anon_sym_false),
	2738: uint16(56),
	2739: uint16(2),
	2740: uint16(sym_ast_block),
	2741: uint16(aux_sym_source_file_repeat1),
	2742: uint16(78),
	2743: uint16(2),
	2744: uint16(sym_number),
	2745: uint16(sym_boolean),
	2746: uint16(77),
	2747: uint16(6),
	2748: uint16(sym_loop_widget),
	2749: uint16(sym_list),
	2750: uint16(sym_array),
	2751: uint16(sym_literal),
	2752: uint16(sym_string),
	2753: uint16(sym_expr),
	2754: uint16(16),
	2755: uint16(3),
	2756: uint16(1),
	2757: uint16(sym_comment),
	2758: uint16(7),
	2759: uint16(1),
	2760: uint16(sym_symbol),
	2761: uint16(9),
	2762: uint16(1),
	2763: uint16(anon_sym_LPAREN),
	2764: uint16(11),
	2765: uint16(1),
	2766: uint16(anon_sym_LBRACK),
	2767: uint16(13),
	2768: uint16(1),
	2769: uint16(sym_keyword),
	2770: uint16(15),
	2771: uint16(1),
	2772: uint16(sym_integer),
	2773: uint16(17),
	2774: uint16(1),
	2775: uint16(sym_float),
	2776: uint16(21),
	2777: uint16(1),
	2778: uint16(anon_sym_DQUOTE),
	2779: uint16(23),
	2780: uint16(1),
	2781: uint16(anon_sym_SQUOTE),
	2782: uint16(25),
	2783: uint16(1),
	2784: uint16(anon_sym_BQUOTE),
	2785: uint16(27),
	2786: uint16(1),
	2787: uint16(anon_sym_LBRACE),
	2788: uint16(216),
	2789: uint16(1),
	2790: uint16(anon_sym_RPAREN),
	2791: uint16(19),
	2792: uint16(2),
	2793: uint16(anon_sym_true),
	2794: uint16(anon_sym_false),
	2795: uint16(25),
	2796: uint16(2),
	2797: uint16(sym_ast_block),
	2798: uint16(aux_sym_source_file_repeat1),
	2799: uint16(78),
	2800: uint16(2),
	2801: uint16(sym_number),
	2802: uint16(sym_boolean),
	2803: uint16(77),
	2804: uint16(6),
	2805: uint16(sym_loop_widget),
	2806: uint16(sym_list),
	2807: uint16(sym_array),
	2808: uint16(sym_literal),
	2809: uint16(sym_string),
	2810: uint16(sym_expr),
	2811: uint16(15),
	2812: uint16(3),
	2813: uint16(1),
	2814: uint16(sym_comment),
	2815: uint16(100),
	2816: uint16(1),
	2817: uint16(anon_sym_LBRACK),
	2818: uint16(102),
	2819: uint16(1),
	2820: uint16(anon_sym_QMARK_DOT),
	2821: uint16(104),
	2822: uint16(1),
	2823: uint16(anon_sym_DOT),
	2824: uint16(170),
	2825: uint16(1),
	2826: uint16(anon_sym_PIPE_PIPE),
	2827: uint16(172),
	2828: uint16(1),
	2829: uint16(anon_sym_QMARK),
	2830: uint16(218),
	2831: uint16(1),
	2832: uint16(anon_sym_RPAREN),
	2833: uint16(220),
	2834: uint16(1),
	2835: uint16(anon_sym_COMMA),
	2836: uint16(119),
	2837: uint16(1),
	2838: uint16(aux_sym_json_array_repeat1),
	2839: uint16(106),
	2840: uint16(2),
	2841: uint16(anon_sym_PLUS),
	2842: uint16(anon_sym_DASH),
	2843: uint16(110),
	2844: uint16(2),
	2845: uint16(anon_sym_GT_EQ),
	2846: uint16(anon_sym_LT_EQ),
	2847: uint16(112),
	2848: uint16(2),
	2849: uint16(anon_sym_GT),
	2850: uint16(anon_sym_LT),
	2851: uint16(138),
	2852: uint16(2),
	2853: uint16(anon_sym_AMP_AMP),
	2854: uint16(anon_sym_QMARK_COLON),
	2855: uint16(108),
	2856: uint16(3),
	2857: uint16(anon_sym_STAR),
	2858: uint16(anon_sym_SLASH),
	2859: uint16(anon_sym_PERCENT),
	2860: uint16(136),
	2861: uint16(3),
	2862: uint16(anon_sym_EQ_EQ),
	2863: uint16(anon_sym_BANG_EQ),
	2864: uint16(anon_sym_EQ_TILDE),
	2865: uint16(15),
	2866: uint16(3),
	2867: uint16(1),
	2868: uint16(sym_comment),
	2869: uint16(100),
	2870: uint16(1),
	2871: uint16(anon_sym_LBRACK),
	2872: uint16(102),
	2873: uint16(1),
	2874: uint16(anon_sym_QMARK_DOT),
	2875: uint16(104),
	2876: uint16(1),
	2877: uint16(anon_sym_DOT),
	2878: uint16(170),
	2879: uint16(1),
	2880: uint16(anon_sym_PIPE_PIPE),
	2881: uint16(172),
	2882: uint16(1),
	2883: uint16(anon_sym_QMARK),
	2884: uint16(222),
	2885: uint16(1),
	2886: uint16(anon_sym_RBRACE),
	2887: uint16(224),
	2888: uint16(1),
	2889: uint16(anon_sym_COMMA),
	2890: uint16(118),
	2891: uint16(1),
	2892: uint16(aux_sym_json_object_repeat1),
	2893: uint16(106),
	2894: uint16(2),
	2895: uint16(anon_sym_PLUS),
	2896: uint16(anon_sym_DASH),
	2897: uint16(110),
	2898: uint16(2),
	2899: uint16(anon_sym_GT_EQ),
	2900: uint16(anon_sym_LT_EQ),
	2901: uint16(112),
	2902: uint16(2),
	2903: uint16(anon_sym_GT),
	2904: uint16(anon_sym_LT),
	2905: uint16(138),
	2906: uint16(2),
	2907: uint16(anon_sym_AMP_AMP),
	2908: uint16(anon_sym_QMARK_COLON),
	2909: uint16(108),
	2910: uint16(3),
	2911: uint16(anon_sym_STAR),
	2912: uint16(anon_sym_SLASH),
	2913: uint16(anon_sym_PERCENT),
	2914: uint16(136),
	2915: uint16(3),
	2916: uint16(anon_sym_EQ_EQ),
	2917: uint16(anon_sym_BANG_EQ),
	2918: uint16(anon_sym_EQ_TILDE),
	2919: uint16(13),
	2920: uint16(3),
	2921: uint16(1),
	2922: uint16(sym_comment),
	2923: uint16(100),
	2924: uint16(1),
	2925: uint16(anon_sym_LBRACK),
	2926: uint16(102),
	2927: uint16(1),
	2928: uint16(anon_sym_QMARK_DOT),
	2929: uint16(104),
	2930: uint16(1),
	2931: uint16(anon_sym_DOT),
	2932: uint16(170),
	2933: uint16(1),
	2934: uint16(anon_sym_PIPE_PIPE),
	2935: uint16(172),
	2936: uint16(1),
	2937: uint16(anon_sym_QMARK),
	2938: uint16(106),
	2939: uint16(2),
	2940: uint16(anon_sym_PLUS),
	2941: uint16(anon_sym_DASH),
	2942: uint16(110),
	2943: uint16(2),
	2944: uint16(anon_sym_GT_EQ),
	2945: uint16(anon_sym_LT_EQ),
	2946: uint16(112),
	2947: uint16(2),
	2948: uint16(anon_sym_GT),
	2949: uint16(anon_sym_LT),
	2950: uint16(138),
	2951: uint16(2),
	2952: uint16(anon_sym_AMP_AMP),
	2953: uint16(anon_sym_QMARK_COLON),
	2954: uint16(108),
	2955: uint16(3),
	2956: uint16(anon_sym_STAR),
	2957: uint16(anon_sym_SLASH),
	2958: uint16(anon_sym_PERCENT),
	2959: uint16(136),
	2960: uint16(3),
	2961: uint16(anon_sym_EQ_EQ),
	2962: uint16(anon_sym_BANG_EQ),
	2963: uint16(anon_sym_EQ_TILDE),
	2964: uint16(226),
	2965: uint16(3),
	2966: uint16(anon_sym_RPAREN),
	2967: uint16(anon_sym_RBRACK),
	2968: uint16(anon_sym_COMMA),
	2969: uint16(15),
	2970: uint16(3),
	2971: uint16(1),
	2972: uint16(sym_comment),
	2973: uint16(100),
	2974: uint16(1),
	2975: uint16(anon_sym_LBRACK),
	2976: uint16(102),
	2977: uint16(1),
	2978: uint16(anon_sym_QMARK_DOT),
	2979: uint16(104),
	2980: uint16(1),
	2981: uint16(anon_sym_DOT),
	2982: uint16(170),
	2983: uint16(1),
	2984: uint16(anon_sym_PIPE_PIPE),
	2985: uint16(172),
	2986: uint16(1),
	2987: uint16(anon_sym_QMARK),
	2988: uint16(220),
	2989: uint16(1),
	2990: uint16(anon_sym_COMMA),
	2991: uint16(228),
	2992: uint16(1),
	2993: uint16(anon_sym_RBRACK),
	2994: uint16(120),
	2995: uint16(1),
	2996: uint16(aux_sym_json_array_repeat1),
	2997: uint16(106),
	2998: uint16(2),
	2999: uint16(anon_sym_PLUS),
	3000: uint16(anon_sym_DASH),
	3001: uint16(110),
	3002: uint16(2),
	3003: uint16(anon_sym_GT_EQ),
	3004: uint16(anon_sym_LT_EQ),
	3005: uint16(112),
	3006: uint16(2),
	3007: uint16(anon_sym_GT),
	3008: uint16(anon_sym_LT),
	3009: uint16(138),
	3010: uint16(2),
	3011: uint16(anon_sym_AMP_AMP),
	3012: uint16(anon_sym_QMARK_COLON),
	3013: uint16(108),
	3014: uint16(3),
	3015: uint16(anon_sym_STAR),
	3016: uint16(anon_sym_SLASH),
	3017: uint16(anon_sym_PERCENT),
	3018: uint16(136),
	3019: uint16(3),
	3020: uint16(anon_sym_EQ_EQ),
	3021: uint16(anon_sym_BANG_EQ),
	3022: uint16(anon_sym_EQ_TILDE),
	3023: uint16(13),
	3024: uint16(3),
	3025: uint16(1),
	3026: uint16(sym_comment),
	3027: uint16(100),
	3028: uint16(1),
	3029: uint16(anon_sym_LBRACK),
	3030: uint16(102),
	3031: uint16(1),
	3032: uint16(anon_sym_QMARK_DOT),
	3033: uint16(104),
	3034: uint16(1),
	3035: uint16(anon_sym_DOT),
	3036: uint16(170),
	3037: uint16(1),
	3038: uint16(anon_sym_PIPE_PIPE),
	3039: uint16(172),
	3040: uint16(1),
	3041: uint16(anon_sym_QMARK),
	3042: uint16(106),
	3043: uint16(2),
	3044: uint16(anon_sym_PLUS),
	3045: uint16(anon_sym_DASH),
	3046: uint16(110),
	3047: uint16(2),
	3048: uint16(anon_sym_GT_EQ),
	3049: uint16(anon_sym_LT_EQ),
	3050: uint16(112),
	3051: uint16(2),
	3052: uint16(anon_sym_GT),
	3053: uint16(anon_sym_LT),
	3054: uint16(138),
	3055: uint16(2),
	3056: uint16(anon_sym_AMP_AMP),
	3057: uint16(anon_sym_QMARK_COLON),
	3058: uint16(230),
	3059: uint16(2),
	3060: uint16(anon_sym_RBRACE),
	3061: uint16(anon_sym_COMMA),
	3062: uint16(108),
	3063: uint16(3),
	3064: uint16(anon_sym_STAR),
	3065: uint16(anon_sym_SLASH),
	3066: uint16(anon_sym_PERCENT),
	3067: uint16(136),
	3068: uint16(3),
	3069: uint16(anon_sym_EQ_EQ),
	3070: uint16(anon_sym_BANG_EQ),
	3071: uint16(anon_sym_EQ_TILDE),
	3072: uint16(15),
	3073: uint16(3),
	3074: uint16(1),
	3075: uint16(sym_comment),
	3076: uint16(7),
	3077: uint16(1),
	3078: uint16(sym_symbol),
	3079: uint16(9),
	3080: uint16(1),
	3081: uint16(anon_sym_LPAREN),
	3082: uint16(11),
	3083: uint16(1),
	3084: uint16(anon_sym_LBRACK),
	3085: uint16(13),
	3086: uint16(1),
	3087: uint16(sym_keyword),
	3088: uint16(15),
	3089: uint16(1),
	3090: uint16(sym_integer),
	3091: uint16(17),
	3092: uint16(1),
	3093: uint16(sym_float),
	3094: uint16(21),
	3095: uint16(1),
	3096: uint16(anon_sym_DQUOTE),
	3097: uint16(23),
	3098: uint16(1),
	3099: uint16(anon_sym_SQUOTE),
	3100: uint16(25),
	3101: uint16(1),
	3102: uint16(anon_sym_BQUOTE),
	3103: uint16(27),
	3104: uint16(1),
	3105: uint16(anon_sym_LBRACE),
	3106: uint16(126),
	3107: uint16(1),
	3108: uint16(sym_ast_block),
	3109: uint16(19),
	3110: uint16(2),
	3111: uint16(anon_sym_true),
	3112: uint16(anon_sym_false),
	3113: uint16(78),
	3114: uint16(2),
	3115: uint16(sym_number),
	3116: uint16(sym_boolean),
	3117: uint16(77),
	3118: uint16(6),
	3119: uint16(sym_loop_widget),
	3120: uint16(sym_list),
	3121: uint16(sym_array),
	3122: uint16(sym_literal),
	3123: uint16(sym_string),
	3124: uint16(sym_expr),
	3125: uint16(13),
	3126: uint16(3),
	3127: uint16(1),
	3128: uint16(sym_comment),
	3129: uint16(100),
	3130: uint16(1),
	3131: uint16(anon_sym_LBRACK),
	3132: uint16(102),
	3133: uint16(1),
	3134: uint16(anon_sym_QMARK_DOT),
	3135: uint16(104),
	3136: uint16(1),
	3137: uint16(anon_sym_DOT),
	3138: uint16(170),
	3139: uint16(1),
	3140: uint16(anon_sym_PIPE_PIPE),
	3141: uint16(172),
	3142: uint16(1),
	3143: uint16(anon_sym_QMARK),
	3144: uint16(232),
	3145: uint16(1),
	3146: uint16(anon_sym_RBRACE),
	3147: uint16(106),
	3148: uint16(2),
	3149: uint16(anon_sym_PLUS),
	3150: uint16(anon_sym_DASH),
	3151: uint16(110),
	3152: uint16(2),
	3153: uint16(anon_sym_GT_EQ),
	3154: uint16(anon_sym_LT_EQ),
	3155: uint16(112),
	3156: uint16(2),
	3157: uint16(anon_sym_GT),
	3158: uint16(anon_sym_LT),
	3159: uint16(138),
	3160: uint16(2),
	3161: uint16(anon_sym_AMP_AMP),
	3162: uint16(anon_sym_QMARK_COLON),
	3163: uint16(108),
	3164: uint16(3),
	3165: uint16(anon_sym_STAR),
	3166: uint16(anon_sym_SLASH),
	3167: uint16(anon_sym_PERCENT),
	3168: uint16(136),
	3169: uint16(3),
	3170: uint16(anon_sym_EQ_EQ),
	3171: uint16(anon_sym_BANG_EQ),
	3172: uint16(anon_sym_EQ_TILDE),
	3173: uint16(13),
	3174: uint16(3),
	3175: uint16(1),
	3176: uint16(sym_comment),
	3177: uint16(100),
	3178: uint16(1),
	3179: uint16(anon_sym_LBRACK),
	3180: uint16(102),
	3181: uint16(1),
	3182: uint16(anon_sym_QMARK_DOT),
	3183: uint16(104),
	3184: uint16(1),
	3185: uint16(anon_sym_DOT),
	3186: uint16(170),
	3187: uint16(1),
	3188: uint16(anon_sym_PIPE_PIPE),
	3189: uint16(172),
	3190: uint16(1),
	3191: uint16(anon_sym_QMARK),
	3192: uint16(234),
	3193: uint16(1),
	3194: uint16(anon_sym_RBRACK),
	3195: uint16(106),
	3196: uint16(2),
	3197: uint16(anon_sym_PLUS),
	3198: uint16(anon_sym_DASH),
	3199: uint16(110),
	3200: uint16(2),
	3201: uint16(anon_sym_GT_EQ),
	3202: uint16(anon_sym_LT_EQ),
	3203: uint16(112),
	3204: uint16(2),
	3205: uint16(anon_sym_GT),
	3206: uint16(anon_sym_LT),
	3207: uint16(138),
	3208: uint16(2),
	3209: uint16(anon_sym_AMP_AMP),
	3210: uint16(anon_sym_QMARK_COLON),
	3211: uint16(108),
	3212: uint16(3),
	3213: uint16(anon_sym_STAR),
	3214: uint16(anon_sym_SLASH),
	3215: uint16(anon_sym_PERCENT),
	3216: uint16(136),
	3217: uint16(3),
	3218: uint16(anon_sym_EQ_EQ),
	3219: uint16(anon_sym_BANG_EQ),
	3220: uint16(anon_sym_EQ_TILDE),
	3221: uint16(13),
	3222: uint16(3),
	3223: uint16(1),
	3224: uint16(sym_comment),
	3225: uint16(100),
	3226: uint16(1),
	3227: uint16(anon_sym_LBRACK),
	3228: uint16(102),
	3229: uint16(1),
	3230: uint16(anon_sym_QMARK_DOT),
	3231: uint16(104),
	3232: uint16(1),
	3233: uint16(anon_sym_DOT),
	3234: uint16(170),
	3235: uint16(1),
	3236: uint16(anon_sym_PIPE_PIPE),
	3237: uint16(172),
	3238: uint16(1),
	3239: uint16(anon_sym_QMARK),
	3240: uint16(236),
	3241: uint16(1),
	3242: uint16(anon_sym_RBRACE),
	3243: uint16(106),
	3244: uint16(2),
	3245: uint16(anon_sym_PLUS),
	3246: uint16(anon_sym_DASH),
	3247: uint16(110),
	3248: uint16(2),
	3249: uint16(anon_sym_GT_EQ),
	3250: uint16(anon_sym_LT_EQ),
	3251: uint16(112),
	3252: uint16(2),
	3253: uint16(anon_sym_GT),
	3254: uint16(anon_sym_LT),
	3255: uint16(138),
	3256: uint16(2),
	3257: uint16(anon_sym_AMP_AMP),
	3258: uint16(anon_sym_QMARK_COLON),
	3259: uint16(108),
	3260: uint16(3),
	3261: uint16(anon_sym_STAR),
	3262: uint16(anon_sym_SLASH),
	3263: uint16(anon_sym_PERCENT),
	3264: uint16(136),
	3265: uint16(3),
	3266: uint16(anon_sym_EQ_EQ),
	3267: uint16(anon_sym_BANG_EQ),
	3268: uint16(anon_sym_EQ_TILDE),
	3269: uint16(13),
	3270: uint16(3),
	3271: uint16(1),
	3272: uint16(sym_comment),
	3273: uint16(100),
	3274: uint16(1),
	3275: uint16(anon_sym_LBRACK),
	3276: uint16(102),
	3277: uint16(1),
	3278: uint16(anon_sym_QMARK_DOT),
	3279: uint16(104),
	3280: uint16(1),
	3281: uint16(anon_sym_DOT),
	3282: uint16(170),
	3283: uint16(1),
	3284: uint16(anon_sym_PIPE_PIPE),
	3285: uint16(172),
	3286: uint16(1),
	3287: uint16(anon_sym_QMARK),
	3288: uint16(238),
	3289: uint16(1),
	3290: uint16(anon_sym_RBRACE),
	3291: uint16(106),
	3292: uint16(2),
	3293: uint16(anon_sym_PLUS),
	3294: uint16(anon_sym_DASH),
	3295: uint16(110),
	3296: uint16(2),
	3297: uint16(anon_sym_GT_EQ),
	3298: uint16(anon_sym_LT_EQ),
	3299: uint16(112),
	3300: uint16(2),
	3301: uint16(anon_sym_GT),
	3302: uint16(anon_sym_LT),
	3303: uint16(138),
	3304: uint16(2),
	3305: uint16(anon_sym_AMP_AMP),
	3306: uint16(anon_sym_QMARK_COLON),
	3307: uint16(108),
	3308: uint16(3),
	3309: uint16(anon_sym_STAR),
	3310: uint16(anon_sym_SLASH),
	3311: uint16(anon_sym_PERCENT),
	3312: uint16(136),
	3313: uint16(3),
	3314: uint16(anon_sym_EQ_EQ),
	3315: uint16(anon_sym_BANG_EQ),
	3316: uint16(anon_sym_EQ_TILDE),
	3317: uint16(13),
	3318: uint16(3),
	3319: uint16(1),
	3320: uint16(sym_comment),
	3321: uint16(100),
	3322: uint16(1),
	3323: uint16(anon_sym_LBRACK),
	3324: uint16(102),
	3325: uint16(1),
	3326: uint16(anon_sym_QMARK_DOT),
	3327: uint16(104),
	3328: uint16(1),
	3329: uint16(anon_sym_DOT),
	3330: uint16(170),
	3331: uint16(1),
	3332: uint16(anon_sym_PIPE_PIPE),
	3333: uint16(172),
	3334: uint16(1),
	3335: uint16(anon_sym_QMARK),
	3336: uint16(240),
	3337: uint16(1),
	3338: uint16(anon_sym_COLON),
	3339: uint16(106),
	3340: uint16(2),
	3341: uint16(anon_sym_PLUS),
	3342: uint16(anon_sym_DASH),
	3343: uint16(110),
	3344: uint16(2),
	3345: uint16(anon_sym_GT_EQ),
	3346: uint16(anon_sym_LT_EQ),
	3347: uint16(112),
	3348: uint16(2),
	3349: uint16(anon_sym_GT),
	3350: uint16(anon_sym_LT),
	3351: uint16(138),
	3352: uint16(2),
	3353: uint16(anon_sym_AMP_AMP),
	3354: uint16(anon_sym_QMARK_COLON),
	3355: uint16(108),
	3356: uint16(3),
	3357: uint16(anon_sym_STAR),
	3358: uint16(anon_sym_SLASH),
	3359: uint16(anon_sym_PERCENT),
	3360: uint16(136),
	3361: uint16(3),
	3362: uint16(anon_sym_EQ_EQ),
	3363: uint16(anon_sym_BANG_EQ),
	3364: uint16(anon_sym_EQ_TILDE),
	3365: uint16(13),
	3366: uint16(3),
	3367: uint16(1),
	3368: uint16(sym_comment),
	3369: uint16(100),
	3370: uint16(1),
	3371: uint16(anon_sym_LBRACK),
	3372: uint16(102),
	3373: uint16(1),
	3374: uint16(anon_sym_QMARK_DOT),
	3375: uint16(104),
	3376: uint16(1),
	3377: uint16(anon_sym_DOT),
	3378: uint16(170),
	3379: uint16(1),
	3380: uint16(anon_sym_PIPE_PIPE),
	3381: uint16(172),
	3382: uint16(1),
	3383: uint16(anon_sym_QMARK),
	3384: uint16(242),
	3385: uint16(1),
	3386: uint16(anon_sym_RBRACK),
	3387: uint16(106),
	3388: uint16(2),
	3389: uint16(anon_sym_PLUS),
	3390: uint16(anon_sym_DASH),
	3391: uint16(110),
	3392: uint16(2),
	3393: uint16(anon_sym_GT_EQ),
	3394: uint16(anon_sym_LT_EQ),
	3395: uint16(112),
	3396: uint16(2),
	3397: uint16(anon_sym_GT),
	3398: uint16(anon_sym_LT),
	3399: uint16(138),
	3400: uint16(2),
	3401: uint16(anon_sym_AMP_AMP),
	3402: uint16(anon_sym_QMARK_COLON),
	3403: uint16(108),
	3404: uint16(3),
	3405: uint16(anon_sym_STAR),
	3406: uint16(anon_sym_SLASH),
	3407: uint16(anon_sym_PERCENT),
	3408: uint16(136),
	3409: uint16(3),
	3410: uint16(anon_sym_EQ_EQ),
	3411: uint16(anon_sym_BANG_EQ),
	3412: uint16(anon_sym_EQ_TILDE),
	3413: uint16(13),
	3414: uint16(3),
	3415: uint16(1),
	3416: uint16(sym_comment),
	3417: uint16(100),
	3418: uint16(1),
	3419: uint16(anon_sym_LBRACK),
	3420: uint16(102),
	3421: uint16(1),
	3422: uint16(anon_sym_QMARK_DOT),
	3423: uint16(104),
	3424: uint16(1),
	3425: uint16(anon_sym_DOT),
	3426: uint16(170),
	3427: uint16(1),
	3428: uint16(anon_sym_PIPE_PIPE),
	3429: uint16(172),
	3430: uint16(1),
	3431: uint16(anon_sym_QMARK),
	3432: uint16(244),
	3433: uint16(1),
	3434: uint16(anon_sym_RPAREN),
	3435: uint16(106),
	3436: uint16(2),
	3437: uint16(anon_sym_PLUS),
	3438: uint16(anon_sym_DASH),
	3439: uint16(110),
	3440: uint16(2),
	3441: uint16(anon_sym_GT_EQ),
	3442: uint16(anon_sym_LT_EQ),
	3443: uint16(112),
	3444: uint16(2),
	3445: uint16(anon_sym_GT),
	3446: uint16(anon_sym_LT),
	3447: uint16(138),
	3448: uint16(2),
	3449: uint16(anon_sym_AMP_AMP),
	3450: uint16(anon_sym_QMARK_COLON),
	3451: uint16(108),
	3452: uint16(3),
	3453: uint16(anon_sym_STAR),
	3454: uint16(anon_sym_SLASH),
	3455: uint16(anon_sym_PERCENT),
	3456: uint16(136),
	3457: uint16(3),
	3458: uint16(anon_sym_EQ_EQ),
	3459: uint16(anon_sym_BANG_EQ),
	3460: uint16(anon_sym_EQ_TILDE),
	3461: uint16(13),
	3462: uint16(3),
	3463: uint16(1),
	3464: uint16(sym_comment),
	3465: uint16(100),
	3466: uint16(1),
	3467: uint16(anon_sym_LBRACK),
	3468: uint16(102),
	3469: uint16(1),
	3470: uint16(anon_sym_QMARK_DOT),
	3471: uint16(104),
	3472: uint16(1),
	3473: uint16(anon_sym_DOT),
	3474: uint16(170),
	3475: uint16(1),
	3476: uint16(anon_sym_PIPE_PIPE),
	3477: uint16(172),
	3478: uint16(1),
	3479: uint16(anon_sym_QMARK),
	3480: uint16(246),
	3481: uint16(1),
	3482: uint16(anon_sym_COLON),
	3483: uint16(106),
	3484: uint16(2),
	3485: uint16(anon_sym_PLUS),
	3486: uint16(anon_sym_DASH),
	3487: uint16(110),
	3488: uint16(2),
	3489: uint16(anon_sym_GT_EQ),
	3490: uint16(anon_sym_LT_EQ),
	3491: uint16(112),
	3492: uint16(2),
	3493: uint16(anon_sym_GT),
	3494: uint16(anon_sym_LT),
	3495: uint16(138),
	3496: uint16(2),
	3497: uint16(anon_sym_AMP_AMP),
	3498: uint16(anon_sym_QMARK_COLON),
	3499: uint16(108),
	3500: uint16(3),
	3501: uint16(anon_sym_STAR),
	3502: uint16(anon_sym_SLASH),
	3503: uint16(anon_sym_PERCENT),
	3504: uint16(136),
	3505: uint16(3),
	3506: uint16(anon_sym_EQ_EQ),
	3507: uint16(anon_sym_BANG_EQ),
	3508: uint16(anon_sym_EQ_TILDE),
	3509: uint16(13),
	3510: uint16(3),
	3511: uint16(1),
	3512: uint16(sym_comment),
	3513: uint16(100),
	3514: uint16(1),
	3515: uint16(anon_sym_LBRACK),
	3516: uint16(102),
	3517: uint16(1),
	3518: uint16(anon_sym_QMARK_DOT),
	3519: uint16(104),
	3520: uint16(1),
	3521: uint16(anon_sym_DOT),
	3522: uint16(170),
	3523: uint16(1),
	3524: uint16(anon_sym_PIPE_PIPE),
	3525: uint16(172),
	3526: uint16(1),
	3527: uint16(anon_sym_QMARK),
	3528: uint16(248),
	3529: uint16(1),
	3530: uint16(anon_sym_COLON),
	3531: uint16(106),
	3532: uint16(2),
	3533: uint16(anon_sym_PLUS),
	3534: uint16(anon_sym_DASH),
	3535: uint16(110),
	3536: uint16(2),
	3537: uint16(anon_sym_GT_EQ),
	3538: uint16(anon_sym_LT_EQ),
	3539: uint16(112),
	3540: uint16(2),
	3541: uint16(anon_sym_GT),
	3542: uint16(anon_sym_LT),
	3543: uint16(138),
	3544: uint16(2),
	3545: uint16(anon_sym_AMP_AMP),
	3546: uint16(anon_sym_QMARK_COLON),
	3547: uint16(108),
	3548: uint16(3),
	3549: uint16(anon_sym_STAR),
	3550: uint16(anon_sym_SLASH),
	3551: uint16(anon_sym_PERCENT),
	3552: uint16(136),
	3553: uint16(3),
	3554: uint16(anon_sym_EQ_EQ),
	3555: uint16(anon_sym_BANG_EQ),
	3556: uint16(anon_sym_EQ_TILDE),
	3557: uint16(13),
	3558: uint16(3),
	3559: uint16(1),
	3560: uint16(sym_comment),
	3561: uint16(100),
	3562: uint16(1),
	3563: uint16(anon_sym_LBRACK),
	3564: uint16(102),
	3565: uint16(1),
	3566: uint16(anon_sym_QMARK_DOT),
	3567: uint16(104),
	3568: uint16(1),
	3569: uint16(anon_sym_DOT),
	3570: uint16(170),
	3571: uint16(1),
	3572: uint16(anon_sym_PIPE_PIPE),
	3573: uint16(172),
	3574: uint16(1),
	3575: uint16(anon_sym_QMARK),
	3576: uint16(250),
	3577: uint16(1),
	3578: uint16(anon_sym_RBRACE),
	3579: uint16(106),
	3580: uint16(2),
	3581: uint16(anon_sym_PLUS),
	3582: uint16(anon_sym_DASH),
	3583: uint16(110),
	3584: uint16(2),
	3585: uint16(anon_sym_GT_EQ),
	3586: uint16(anon_sym_LT_EQ),
	3587: uint16(112),
	3588: uint16(2),
	3589: uint16(anon_sym_GT),
	3590: uint16(anon_sym_LT),
	3591: uint16(138),
	3592: uint16(2),
	3593: uint16(anon_sym_AMP_AMP),
	3594: uint16(anon_sym_QMARK_COLON),
	3595: uint16(108),
	3596: uint16(3),
	3597: uint16(anon_sym_STAR),
	3598: uint16(anon_sym_SLASH),
	3599: uint16(anon_sym_PERCENT),
	3600: uint16(136),
	3601: uint16(3),
	3602: uint16(anon_sym_EQ_EQ),
	3603: uint16(anon_sym_BANG_EQ),
	3604: uint16(anon_sym_EQ_TILDE),
	3605: uint16(3),
	3606: uint16(3),
	3607: uint16(1),
	3608: uint16(sym_comment),
	3609: uint16(184),
	3610: uint16(4),
	3611: uint16(sym_symbol),
	3612: uint16(sym_integer),
	3613: uint16(anon_sym_true),
	3614: uint16(anon_sym_false),
	3615: uint16(182),
	3616: uint16(11),
	3618: uint16(anon_sym_LPAREN),
	3619: uint16(anon_sym_RPAREN),
	3620: uint16(anon_sym_LBRACK),
	3621: uint16(anon_sym_RBRACK),
	3622: uint16(sym_keyword),
	3623: uint16(sym_float),
	3624: uint16(anon_sym_DQUOTE),
	3625: uint16(anon_sym_SQUOTE),
	3626: uint16(anon_sym_BQUOTE),
	3627: uint16(anon_sym_LBRACE),
	3628: uint16(3),
	3629: uint16(3),
	3630: uint16(1),
	3631: uint16(sym_comment),
	3632: uint16(254),
	3633: uint16(4),
	3634: uint16(sym_symbol),
	3635: uint16(sym_integer),
	3636: uint16(anon_sym_true),
	3637: uint16(anon_sym_false),
	3638: uint16(252),
	3639: uint16(11),
	3641: uint16(anon_sym_LPAREN),
	3642: uint16(anon_sym_RPAREN),
	3643: uint16(anon_sym_LBRACK),
	3644: uint16(anon_sym_RBRACK),
	3645: uint16(sym_keyword),
	3646: uint16(sym_float),
	3647: uint16(anon_sym_DQUOTE),
	3648: uint16(anon_sym_SQUOTE),
	3649: uint16(anon_sym_BQUOTE),
	3650: uint16(anon_sym_LBRACE),
	3651: uint16(3),
	3652: uint16(3),
	3653: uint16(1),
	3654: uint16(sym_comment),
	3655: uint16(192),
	3656: uint16(4),
	3657: uint16(sym_symbol),
	3658: uint16(sym_integer),
	3659: uint16(anon_sym_true),
	3660: uint16(anon_sym_false),
	3661: uint16(190),
	3662: uint16(11),
	3664: uint16(anon_sym_LPAREN),
	3665: uint16(anon_sym_RPAREN),
	3666: uint16(anon_sym_LBRACK),
	3667: uint16(anon_sym_RBRACK),
	3668: uint16(sym_keyword),
	3669: uint16(sym_float),
	3670: uint16(anon_sym_DQUOTE),
	3671: uint16(anon_sym_SQUOTE),
	3672: uint16(anon_sym_BQUOTE),
	3673: uint16(anon_sym_LBRACE),
	3674: uint16(3),
	3675: uint16(3),
	3676: uint16(1),
	3677: uint16(sym_comment),
	3678: uint16(258),
	3679: uint16(4),
	3680: uint16(sym_symbol),
	3681: uint16(sym_integer),
	3682: uint16(anon_sym_true),
	3683: uint16(anon_sym_false),
	3684: uint16(256),
	3685: uint16(11),
	3687: uint16(anon_sym_LPAREN),
	3688: uint16(anon_sym_RPAREN),
	3689: uint16(anon_sym_LBRACK),
	3690: uint16(anon_sym_RBRACK),
	3691: uint16(sym_keyword),
	3692: uint16(sym_float),
	3693: uint16(anon_sym_DQUOTE),
	3694: uint16(anon_sym_SQUOTE),
	3695: uint16(anon_sym_BQUOTE),
	3696: uint16(anon_sym_LBRACE),
	3697: uint16(3),
	3698: uint16(3),
	3699: uint16(1),
	3700: uint16(sym_comment),
	3701: uint16(262),
	3702: uint16(4),
	3703: uint16(sym_symbol),
	3704: uint16(sym_integer),
	3705: uint16(anon_sym_true),
	3706: uint16(anon_sym_false),
	3707: uint16(260),
	3708: uint16(11),
	3710: uint16(anon_sym_LPAREN),
	3711: uint16(anon_sym_RPAREN),
	3712: uint16(anon_sym_LBRACK),
	3713: uint16(anon_sym_RBRACK),
	3714: uint16(sym_keyword),
	3715: uint16(sym_float),
	3716: uint16(anon_sym_DQUOTE),
	3717: uint16(anon_sym_SQUOTE),
	3718: uint16(anon_sym_BQUOTE),
	3719: uint16(anon_sym_LBRACE),
	3720: uint16(3),
	3721: uint16(3),
	3722: uint16(1),
	3723: uint16(sym_comment),
	3724: uint16(266),
	3725: uint16(4),
	3726: uint16(sym_symbol),
	3727: uint16(sym_integer),
	3728: uint16(anon_sym_true),
	3729: uint16(anon_sym_false),
	3730: uint16(264),
	3731: uint16(11),
	3733: uint16(anon_sym_LPAREN),
	3734: uint16(anon_sym_RPAREN),
	3735: uint16(anon_sym_LBRACK),
	3736: uint16(anon_sym_RBRACK),
	3737: uint16(sym_keyword),
	3738: uint16(sym_float),
	3739: uint16(anon_sym_DQUOTE),
	3740: uint16(anon_sym_SQUOTE),
	3741: uint16(anon_sym_BQUOTE),
	3742: uint16(anon_sym_LBRACE),
	3743: uint16(3),
	3744: uint16(3),
	3745: uint16(1),
	3746: uint16(sym_comment),
	3747: uint16(200),
	3748: uint16(4),
	3749: uint16(sym_symbol),
	3750: uint16(sym_integer),
	3751: uint16(anon_sym_true),
	3752: uint16(anon_sym_false),
	3753: uint16(198),
	3754: uint16(11),
	3756: uint16(anon_sym_LPAREN),
	3757: uint16(anon_sym_RPAREN),
	3758: uint16(anon_sym_LBRACK),
	3759: uint16(anon_sym_RBRACK),
	3760: uint16(sym_keyword),
	3761: uint16(sym_float),
	3762: uint16(anon_sym_DQUOTE),
	3763: uint16(anon_sym_SQUOTE),
	3764: uint16(anon_sym_BQUOTE),
	3765: uint16(anon_sym_LBRACE),
	3766: uint16(3),
	3767: uint16(3),
	3768: uint16(1),
	3769: uint16(sym_comment),
	3770: uint16(270),
	3771: uint16(4),
	3772: uint16(sym_symbol),
	3773: uint16(sym_integer),
	3774: uint16(anon_sym_true),
	3775: uint16(anon_sym_false),
	3776: uint16(268),
	3777: uint16(11),
	3779: uint16(anon_sym_LPAREN),
	3780: uint16(anon_sym_RPAREN),
	3781: uint16(anon_sym_LBRACK),
	3782: uint16(anon_sym_RBRACK),
	3783: uint16(sym_keyword),
	3784: uint16(sym_float),
	3785: uint16(anon_sym_DQUOTE),
	3786: uint16(anon_sym_SQUOTE),
	3787: uint16(anon_sym_BQUOTE),
	3788: uint16(anon_sym_LBRACE),
	3789: uint16(3),
	3790: uint16(3),
	3791: uint16(1),
	3792: uint16(sym_comment),
	3793: uint16(274),
	3794: uint16(4),
	3795: uint16(sym_symbol),
	3796: uint16(sym_integer),
	3797: uint16(anon_sym_true),
	3798: uint16(anon_sym_false),
	3799: uint16(272),
	3800: uint16(11),
	3802: uint16(anon_sym_LPAREN),
	3803: uint16(anon_sym_RPAREN),
	3804: uint16(anon_sym_LBRACK),
	3805: uint16(anon_sym_RBRACK),
	3806: uint16(sym_keyword),
	3807: uint16(sym_float),
	3808: uint16(anon_sym_DQUOTE),
	3809: uint16(anon_sym_SQUOTE),
	3810: uint16(anon_sym_BQUOTE),
	3811: uint16(anon_sym_LBRACE),
	3812: uint16(3),
	3813: uint16(3),
	3814: uint16(1),
	3815: uint16(sym_comment),
	3816: uint16(188),
	3817: uint16(4),
	3818: uint16(sym_symbol),
	3819: uint16(sym_integer),
	3820: uint16(anon_sym_true),
	3821: uint16(anon_sym_false),
	3822: uint16(186),
	3823: uint16(11),
	3825: uint16(anon_sym_LPAREN),
	3826: uint16(anon_sym_RPAREN),
	3827: uint16(anon_sym_LBRACK),
	3828: uint16(anon_sym_RBRACK),
	3829: uint16(sym_keyword),
	3830: uint16(sym_float),
	3831: uint16(anon_sym_DQUOTE),
	3832: uint16(anon_sym_SQUOTE),
	3833: uint16(anon_sym_BQUOTE),
	3834: uint16(anon_sym_LBRACE),
	3835: uint16(3),
	3836: uint16(3),
	3837: uint16(1),
	3838: uint16(sym_comment),
	3839: uint16(196),
	3840: uint16(4),
	3841: uint16(sym_symbol),
	3842: uint16(sym_integer),
	3843: uint16(anon_sym_true),
	3844: uint16(anon_sym_false),
	3845: uint16(194),
	3846: uint16(11),
	3848: uint16(anon_sym_LPAREN),
	3849: uint16(anon_sym_RPAREN),
	3850: uint16(anon_sym_LBRACK),
	3851: uint16(anon_sym_RBRACK),
	3852: uint16(sym_keyword),
	3853: uint16(sym_float),
	3854: uint16(anon_sym_DQUOTE),
	3855: uint16(anon_sym_SQUOTE),
	3856: uint16(anon_sym_BQUOTE),
	3857: uint16(anon_sym_LBRACE),
	3858: uint16(3),
	3859: uint16(3),
	3860: uint16(1),
	3861: uint16(sym_comment),
	3862: uint16(278),
	3863: uint16(4),
	3864: uint16(sym_symbol),
	3865: uint16(sym_integer),
	3866: uint16(anon_sym_true),
	3867: uint16(anon_sym_false),
	3868: uint16(276),
	3869: uint16(11),
	3871: uint16(anon_sym_LPAREN),
	3872: uint16(anon_sym_RPAREN),
	3873: uint16(anon_sym_LBRACK),
	3874: uint16(anon_sym_RBRACK),
	3875: uint16(sym_keyword),
	3876: uint16(sym_float),
	3877: uint16(anon_sym_DQUOTE),
	3878: uint16(anon_sym_SQUOTE),
	3879: uint16(anon_sym_BQUOTE),
	3880: uint16(anon_sym_LBRACE),
	3881: uint16(8),
	3882: uint16(3),
	3883: uint16(1),
	3884: uint16(sym_comment),
	3885: uint16(280),
	3886: uint16(1),
	3887: uint16(anon_sym_SQUOTE),
	3888: uint16(282),
	3889: uint16(1),
	3890: uint16(anon_sym_DOLLAR_LBRACE),
	3891: uint16(288),
	3892: uint16(1),
	3893: uint16(sym__unescaped_single_quote_string_fragment),
	3894: uint16(108),
	3895: uint16(1),
	3896: uint16(aux_sym_string_repeat3),
	3897: uint16(113),
	3898: uint16(1),
	3899: uint16(sym__escape_sequence),
	3900: uint16(285),
	3901: uint16(2),
	3902: uint16(aux_sym__escape_sequence_token1),
	3903: uint16(sym_escape_sequence),
	3904: uint16(88),
	3905: uint16(2),
	3906: uint16(sym_string_interpolation),
	3907: uint16(aux_sym_string_repeat4),
	3908: uint16(8),
	3909: uint16(3),
	3910: uint16(1),
	3911: uint16(sym_comment),
	3912: uint16(291),
	3913: uint16(1),
	3914: uint16(anon_sym_BQUOTE),
	3915: uint16(293),
	3916: uint16(1),
	3917: uint16(anon_sym_DOLLAR_LBRACE),
	3918: uint16(297),
	3919: uint16(1),
	3920: uint16(sym__unescaped_backtick_string_fragment),
	3921: uint16(105),
	3922: uint16(1),
	3923: uint16(aux_sym_string_repeat5),
	3924: uint16(110),
	3925: uint16(1),
	3926: uint16(sym__escape_sequence),
	3927: uint16(295),
	3928: uint16(2),
	3929: uint16(aux_sym__escape_sequence_token1),
	3930: uint16(sym_escape_sequence),
	3931: uint16(91),
	3932: uint16(2),
	3933: uint16(sym_string_interpolation),
	3934: uint16(aux_sym_string_repeat6),
	3935: uint16(8),
	3936: uint16(3),
	3937: uint16(1),
	3938: uint16(sym_comment),
	3939: uint16(291),
	3940: uint16(1),
	3941: uint16(anon_sym_SQUOTE),
	3942: uint16(299),
	3943: uint16(1),
	3944: uint16(anon_sym_DOLLAR_LBRACE),
	3945: uint16(303),
	3946: uint16(1),
	3947: uint16(sym__unescaped_single_quote_string_fragment),
	3948: uint16(108),
	3949: uint16(1),
	3950: uint16(aux_sym_string_repeat3),
	3951: uint16(113),
	3952: uint16(1),
	3953: uint16(sym__escape_sequence),
	3954: uint16(301),
	3955: uint16(2),
	3956: uint16(aux_sym__escape_sequence_token1),
	3957: uint16(sym_escape_sequence),
	3958: uint16(88),
	3959: uint16(2),
	3960: uint16(sym_string_interpolation),
	3961: uint16(aux_sym_string_repeat4),
	3962: uint16(8),
	3963: uint16(3),
	3964: uint16(1),
	3965: uint16(sym_comment),
	3966: uint16(305),
	3967: uint16(1),
	3968: uint16(anon_sym_BQUOTE),
	3969: uint16(307),
	3970: uint16(1),
	3971: uint16(anon_sym_DOLLAR_LBRACE),
	3972: uint16(313),
	3973: uint16(1),
	3974: uint16(sym__unescaped_backtick_string_fragment),
	3975: uint16(105),
	3976: uint16(1),
	3977: uint16(aux_sym_string_repeat5),
	3978: uint16(110),
	3979: uint16(1),
	3980: uint16(sym__escape_sequence),
	3981: uint16(310),
	3982: uint16(2),
	3983: uint16(aux_sym__escape_sequence_token1),
	3984: uint16(sym_escape_sequence),
	3985: uint16(91),
	3986: uint16(2),
	3987: uint16(sym_string_interpolation),
	3988: uint16(aux_sym_string_repeat6),
	3989: uint16(8),
	3990: uint16(3),
	3991: uint16(1),
	3992: uint16(sym_comment),
	3993: uint16(299),
	3994: uint16(1),
	3995: uint16(anon_sym_DOLLAR_LBRACE),
	3996: uint16(303),
	3997: uint16(1),
	3998: uint16(sym__unescaped_single_quote_string_fragment),
	3999: uint16(316),
	4000: uint16(1),
	4001: uint16(anon_sym_SQUOTE),
	4002: uint16(108),
	4003: uint16(1),
	4004: uint16(aux_sym_string_repeat3),
	4005: uint16(113),
	4006: uint16(1),
	4007: uint16(sym__escape_sequence),
	4008: uint16(301),
	4009: uint16(2),
	4010: uint16(aux_sym__escape_sequence_token1),
	4011: uint16(sym_escape_sequence),
	4012: uint16(90),
	4013: uint16(2),
	4014: uint16(sym_string_interpolation),
	4015: uint16(aux_sym_string_repeat4),
	4016: uint16(8),
	4017: uint16(3),
	4018: uint16(1),
	4019: uint16(sym_comment),
	4020: uint16(316),
	4021: uint16(1),
	4022: uint16(anon_sym_DQUOTE),
	4023: uint16(318),
	4024: uint16(1),
	4025: uint16(anon_sym_DOLLAR_LBRACE),
	4026: uint16(322),
	4027: uint16(1),
	4028: uint16(sym__unescaped_double_quote_string_fragment),
	4029: uint16(106),
	4030: uint16(1),
	4031: uint16(aux_sym_string_repeat1),
	4032: uint16(114),
	4033: uint16(1),
	4034: uint16(sym__escape_sequence),
	4035: uint16(320),
	4036: uint16(2),
	4037: uint16(aux_sym__escape_sequence_token1),
	4038: uint16(sym_escape_sequence),
	4039: uint16(94),
	4040: uint16(2),
	4041: uint16(sym_string_interpolation),
	4042: uint16(aux_sym_string_repeat2),
	4043: uint16(8),
	4044: uint16(3),
	4045: uint16(1),
	4046: uint16(sym_comment),
	4047: uint16(291),
	4048: uint16(1),
	4049: uint16(anon_sym_DQUOTE),
	4050: uint16(318),
	4051: uint16(1),
	4052: uint16(anon_sym_DOLLAR_LBRACE),
	4053: uint16(322),
	4054: uint16(1),
	4055: uint16(sym__unescaped_double_quote_string_fragment),
	4056: uint16(106),
	4057: uint16(1),
	4058: uint16(aux_sym_string_repeat1),
	4059: uint16(114),
	4060: uint16(1),
	4061: uint16(sym__escape_sequence),
	4062: uint16(320),
	4063: uint16(2),
	4064: uint16(aux_sym__escape_sequence_token1),
	4065: uint16(sym_escape_sequence),
	4066: uint16(102),
	4067: uint16(2),
	4068: uint16(sym_string_interpolation),
	4069: uint16(aux_sym_string_repeat2),
	4070: uint16(8),
	4071: uint16(3),
	4072: uint16(1),
	4073: uint16(sym_comment),
	4074: uint16(318),
	4075: uint16(1),
	4076: uint16(anon_sym_DOLLAR_LBRACE),
	4077: uint16(322),
	4078: uint16(1),
	4079: uint16(sym__unescaped_double_quote_string_fragment),
	4080: uint16(324),
	4081: uint16(1),
	4082: uint16(anon_sym_DQUOTE),
	4083: uint16(106),
	4084: uint16(1),
	4085: uint16(aux_sym_string_repeat1),
	4086: uint16(114),
	4087: uint16(1),
	4088: uint16(sym__escape_sequence),
	4089: uint16(320),
	4090: uint16(2),
	4091: uint16(aux_sym__escape_sequence_token1),
	4092: uint16(sym_escape_sequence),
	4093: uint16(98),
	4094: uint16(2),
	4095: uint16(sym_string_interpolation),
	4096: uint16(aux_sym_string_repeat2),
	4097: uint16(8),
	4098: uint16(3),
	4099: uint16(1),
	4100: uint16(sym_comment),
	4101: uint16(299),
	4102: uint16(1),
	4103: uint16(anon_sym_DOLLAR_LBRACE),
	4104: uint16(303),
	4105: uint16(1),
	4106: uint16(sym__unescaped_single_quote_string_fragment),
	4107: uint16(324),
	4108: uint16(1),
	4109: uint16(anon_sym_SQUOTE),
	4110: uint16(108),
	4111: uint16(1),
	4112: uint16(aux_sym_string_repeat3),
	4113: uint16(113),
	4114: uint16(1),
	4115: uint16(sym__escape_sequence),
	4116: uint16(301),
	4117: uint16(2),
	4118: uint16(aux_sym__escape_sequence_token1),
	4119: uint16(sym_escape_sequence),
	4120: uint16(99),
	4121: uint16(2),
	4122: uint16(sym_string_interpolation),
	4123: uint16(aux_sym_string_repeat4),
	4124: uint16(8),
	4125: uint16(3),
	4126: uint16(1),
	4127: uint16(sym_comment),
	4128: uint16(293),
	4129: uint16(1),
	4130: uint16(anon_sym_DOLLAR_LBRACE),
	4131: uint16(297),
	4132: uint16(1),
	4133: uint16(sym__unescaped_backtick_string_fragment),
	4134: uint16(324),
	4135: uint16(1),
	4136: uint16(anon_sym_BQUOTE),
	4137: uint16(105),
	4138: uint16(1),
	4139: uint16(aux_sym_string_repeat5),
	4140: uint16(110),
	4141: uint16(1),
	4142: uint16(sym__escape_sequence),
	4143: uint16(295),
	4144: uint16(2),
	4145: uint16(aux_sym__escape_sequence_token1),
	4146: uint16(sym_escape_sequence),
	4147: uint16(100),
	4148: uint16(2),
	4149: uint16(sym_string_interpolation),
	4150: uint16(aux_sym_string_repeat6),
	4151: uint16(8),
	4152: uint16(3),
	4153: uint16(1),
	4154: uint16(sym_comment),
	4155: uint16(318),
	4156: uint16(1),
	4157: uint16(anon_sym_DOLLAR_LBRACE),
	4158: uint16(322),
	4159: uint16(1),
	4160: uint16(sym__unescaped_double_quote_string_fragment),
	4161: uint16(326),
	4162: uint16(1),
	4163: uint16(anon_sym_DQUOTE),
	4164: uint16(106),
	4165: uint16(1),
	4166: uint16(aux_sym_string_repeat1),
	4167: uint16(114),
	4168: uint16(1),
	4169: uint16(sym__escape_sequence),
	4170: uint16(320),
	4171: uint16(2),
	4172: uint16(aux_sym__escape_sequence_token1),
	4173: uint16(sym_escape_sequence),
	4174: uint16(102),
	4175: uint16(2),
	4176: uint16(sym_string_interpolation),
	4177: uint16(aux_sym_string_repeat2),
	4178: uint16(8),
	4179: uint16(3),
	4180: uint16(1),
	4181: uint16(sym_comment),
	4182: uint16(299),
	4183: uint16(1),
	4184: uint16(anon_sym_DOLLAR_LBRACE),
	4185: uint16(303),
	4186: uint16(1),
	4187: uint16(sym__unescaped_single_quote_string_fragment),
	4188: uint16(326),
	4189: uint16(1),
	4190: uint16(anon_sym_SQUOTE),
	4191: uint16(108),
	4192: uint16(1),
	4193: uint16(aux_sym_string_repeat3),
	4194: uint16(113),
	4195: uint16(1),
	4196: uint16(sym__escape_sequence),
	4197: uint16(301),
	4198: uint16(2),
	4199: uint16(aux_sym__escape_sequence_token1),
	4200: uint16(sym_escape_sequence),
	4201: uint16(88),
	4202: uint16(2),
	4203: uint16(sym_string_interpolation),
	4204: uint16(aux_sym_string_repeat4),
	4205: uint16(8),
	4206: uint16(3),
	4207: uint16(1),
	4208: uint16(sym_comment),
	4209: uint16(293),
	4210: uint16(1),
	4211: uint16(anon_sym_DOLLAR_LBRACE),
	4212: uint16(297),
	4213: uint16(1),
	4214: uint16(sym__unescaped_backtick_string_fragment),
	4215: uint16(326),
	4216: uint16(1),
	4217: uint16(anon_sym_BQUOTE),
	4218: uint16(105),
	4219: uint16(1),
	4220: uint16(aux_sym_string_repeat5),
	4221: uint16(110),
	4222: uint16(1),
	4223: uint16(sym__escape_sequence),
	4224: uint16(295),
	4225: uint16(2),
	4226: uint16(aux_sym__escape_sequence_token1),
	4227: uint16(sym_escape_sequence),
	4228: uint16(91),
	4229: uint16(2),
	4230: uint16(sym_string_interpolation),
	4231: uint16(aux_sym_string_repeat6),
	4232: uint16(8),
	4233: uint16(3),
	4234: uint16(1),
	4235: uint16(sym_comment),
	4236: uint16(293),
	4237: uint16(1),
	4238: uint16(anon_sym_DOLLAR_LBRACE),
	4239: uint16(297),
	4240: uint16(1),
	4241: uint16(sym__unescaped_backtick_string_fragment),
	4242: uint16(316),
	4243: uint16(1),
	4244: uint16(anon_sym_BQUOTE),
	4245: uint16(105),
	4246: uint16(1),
	4247: uint16(aux_sym_string_repeat5),
	4248: uint16(110),
	4249: uint16(1),
	4250: uint16(sym__escape_sequence),
	4251: uint16(295),
	4252: uint16(2),
	4253: uint16(aux_sym__escape_sequence_token1),
	4254: uint16(sym_escape_sequence),
	4255: uint16(89),
	4256: uint16(2),
	4257: uint16(sym_string_interpolation),
	4258: uint16(aux_sym_string_repeat6),
	4259: uint16(8),
	4260: uint16(3),
	4261: uint16(1),
	4262: uint16(sym_comment),
	4263: uint16(328),
	4264: uint16(1),
	4265: uint16(anon_sym_DQUOTE),
	4266: uint16(330),
	4267: uint16(1),
	4268: uint16(anon_sym_DOLLAR_LBRACE),
	4269: uint16(336),
	4270: uint16(1),
	4271: uint16(sym__unescaped_double_quote_string_fragment),
	4272: uint16(106),
	4273: uint16(1),
	4274: uint16(aux_sym_string_repeat1),
	4275: uint16(114),
	4276: uint16(1),
	4277: uint16(sym__escape_sequence),
	4278: uint16(333),
	4279: uint16(2),
	4280: uint16(aux_sym__escape_sequence_token1),
	4281: uint16(sym_escape_sequence),
	4282: uint16(102),
	4283: uint16(2),
	4284: uint16(sym_string_interpolation),
	4285: uint16(aux_sym_string_repeat2),
	4286: uint16(6),
	4287: uint16(3),
	4288: uint16(1),
	4289: uint16(sym_comment),
	4290: uint16(344),
	4291: uint16(1),
	4292: uint16(sym__unescaped_backtick_string_fragment),
	4293: uint16(103),
	4294: uint16(1),
	4295: uint16(aux_sym_string_repeat5),
	4296: uint16(110),
	4297: uint16(1),
	4298: uint16(sym__escape_sequence),
	4299: uint16(339),
	4300: uint16(2),
	4301: uint16(anon_sym_BQUOTE),
	4302: uint16(anon_sym_DOLLAR_LBRACE),
	4303: uint16(341),
	4304: uint16(2),
	4305: uint16(aux_sym__escape_sequence_token1),
	4306: uint16(sym_escape_sequence),
	4307: uint16(7),
	4308: uint16(3),
	4309: uint16(1),
	4310: uint16(sym_comment),
	4311: uint16(21),
	4312: uint16(1),
	4313: uint16(anon_sym_DQUOTE),
	4314: uint16(23),
	4315: uint16(1),
	4316: uint16(anon_sym_SQUOTE),
	4317: uint16(25),
	4318: uint16(1),
	4319: uint16(anon_sym_BQUOTE),
	4320: uint16(27),
	4321: uint16(1),
	4322: uint16(anon_sym_LBRACE),
	4323: uint16(347),
	4324: uint16(1),
	4325: uint16(sym_symbol),
	4326: uint16(65),
	4327: uint16(2),
	4328: uint16(sym_string),
	4329: uint16(sym_expr),
	4330: uint16(6),
	4331: uint16(3),
	4332: uint16(1),
	4333: uint16(sym_comment),
	4334: uint16(297),
	4335: uint16(1),
	4336: uint16(sym__unescaped_backtick_string_fragment),
	4337: uint16(103),
	4338: uint16(1),
	4339: uint16(aux_sym_string_repeat5),
	4340: uint16(110),
	4341: uint16(1),
	4342: uint16(sym__escape_sequence),
	4343: uint16(295),
	4344: uint16(2),
	4345: uint16(aux_sym__escape_sequence_token1),
	4346: uint16(sym_escape_sequence),
	4347: uint16(349),
	4348: uint16(2),
	4349: uint16(anon_sym_BQUOTE),
	4350: uint16(anon_sym_DOLLAR_LBRACE),
	4351: uint16(6),
	4352: uint16(3),
	4353: uint16(1),
	4354: uint16(sym_comment),
	4355: uint16(322),
	4356: uint16(1),
	4357: uint16(sym__unescaped_double_quote_string_fragment),
	4358: uint16(109),
	4359: uint16(1),
	4360: uint16(aux_sym_string_repeat1),
	4361: uint16(114),
	4362: uint16(1),
	4363: uint16(sym__escape_sequence),
	4364: uint16(320),
	4365: uint16(2),
	4366: uint16(aux_sym__escape_sequence_token1),
	4367: uint16(sym_escape_sequence),
	4368: uint16(351),
	4369: uint16(2),
	4370: uint16(anon_sym_DQUOTE),
	4371: uint16(anon_sym_DOLLAR_LBRACE),
	4372: uint16(6),
	4373: uint16(3),
	4374: uint16(1),
	4375: uint16(sym_comment),
	4376: uint16(358),
	4377: uint16(1),
	4378: uint16(sym__unescaped_single_quote_string_fragment),
	4379: uint16(107),
	4380: uint16(1),
	4381: uint16(aux_sym_string_repeat3),
	4382: uint16(113),
	4383: uint16(1),
	4384: uint16(sym__escape_sequence),
	4385: uint16(353),
	4386: uint16(2),
	4387: uint16(anon_sym_SQUOTE),
	4388: uint16(anon_sym_DOLLAR_LBRACE),
	4389: uint16(355),
	4390: uint16(2),
	4391: uint16(aux_sym__escape_sequence_token1),
	4392: uint16(sym_escape_sequence),
	4393: uint16(6),
	4394: uint16(3),
	4395: uint16(1),
	4396: uint16(sym_comment),
	4397: uint16(303),
	4398: uint16(1),
	4399: uint16(sym__unescaped_single_quote_string_fragment),
	4400: uint16(107),
	4401: uint16(1),
	4402: uint16(aux_sym_string_repeat3),
	4403: uint16(113),
	4404: uint16(1),
	4405: uint16(sym__escape_sequence),
	4406: uint16(301),
	4407: uint16(2),
	4408: uint16(aux_sym__escape_sequence_token1),
	4409: uint16(sym_escape_sequence),
	4410: uint16(361),
	4411: uint16(2),
	4412: uint16(anon_sym_SQUOTE),
	4413: uint16(anon_sym_DOLLAR_LBRACE),
	4414: uint16(6),
	4415: uint16(3),
	4416: uint16(1),
	4417: uint16(sym_comment),
	4418: uint16(368),
	4419: uint16(1),
	4420: uint16(sym__unescaped_double_quote_string_fragment),
	4421: uint16(109),
	4422: uint16(1),
	4423: uint16(aux_sym_string_repeat1),
	4424: uint16(114),
	4425: uint16(1),
	4426: uint16(sym__escape_sequence),
	4427: uint16(363),
	4428: uint16(2),
	4429: uint16(anon_sym_DQUOTE),
	4430: uint16(anon_sym_DOLLAR_LBRACE),
	4431: uint16(365),
	4432: uint16(2),
	4433: uint16(aux_sym__escape_sequence_token1),
	4434: uint16(sym_escape_sequence),
	4435: uint16(3),
	4436: uint16(3),
	4437: uint16(1),
	4438: uint16(sym_comment),
	4439: uint16(373),
	4440: uint16(2),
	4441: uint16(aux_sym__escape_sequence_token1),
	4442: uint16(sym_escape_sequence),
	4443: uint16(371),
	4444: uint16(3),
	4445: uint16(sym__unescaped_backtick_string_fragment),
	4446: uint16(anon_sym_BQUOTE),
	4447: uint16(anon_sym_DOLLAR_LBRACE),
	4448: uint16(3),
	4449: uint16(3),
	4450: uint16(1),
	4451: uint16(sym_comment),
	4452: uint16(377),
	4453: uint16(2),
	4454: uint16(aux_sym__escape_sequence_token1),
	4455: uint16(sym_escape_sequence),
	4456: uint16(375),
	4457: uint16(3),
	4458: uint16(sym__unescaped_single_quote_string_fragment),
	4459: uint16(anon_sym_SQUOTE),
	4460: uint16(anon_sym_DOLLAR_LBRACE),
	4461: uint16(3),
	4462: uint16(3),
	4463: uint16(1),
	4464: uint16(sym_comment),
	4465: uint16(377),
	4466: uint16(2),
	4467: uint16(aux_sym__escape_sequence_token1),
	4468: uint16(sym_escape_sequence),
	4469: uint16(375),
	4470: uint16(3),
	4471: uint16(sym__unescaped_double_quote_string_fragment),
	4472: uint16(anon_sym_DQUOTE),
	4473: uint16(anon_sym_DOLLAR_LBRACE),
	4474: uint16(3),
	4475: uint16(3),
	4476: uint16(1),
	4477: uint16(sym_comment),
	4478: uint16(381),
	4479: uint16(2),
	4480: uint16(aux_sym__escape_sequence_token1),
	4481: uint16(sym_escape_sequence),
	4482: uint16(379),
	4483: uint16(3),
	4484: uint16(sym__unescaped_single_quote_string_fragment),
	4485: uint16(anon_sym_SQUOTE),
	4486: uint16(anon_sym_DOLLAR_LBRACE),
	4487: uint16(3),
	4488: uint16(3),
	4489: uint16(1),
	4490: uint16(sym_comment),
	4491: uint16(385),
	4492: uint16(2),
	4493: uint16(aux_sym__escape_sequence_token1),
	4494: uint16(sym_escape_sequence),
	4495: uint16(383),
	4496: uint16(3),
	4497: uint16(sym__unescaped_double_quote_string_fragment),
	4498: uint16(anon_sym_DQUOTE),
	4499: uint16(anon_sym_DOLLAR_LBRACE),
	4500: uint16(3),
	4501: uint16(3),
	4502: uint16(1),
	4503: uint16(sym_comment),
	4504: uint16(377),
	4505: uint16(2),
	4506: uint16(aux_sym__escape_sequence_token1),
	4507: uint16(sym_escape_sequence),
	4508: uint16(375),
	4509: uint16(3),
	4510: uint16(sym__unescaped_backtick_string_fragment),
	4511: uint16(anon_sym_BQUOTE),
	4512: uint16(anon_sym_DOLLAR_LBRACE),
	4513: uint16(4),
	4514: uint16(3),
	4515: uint16(1),
	4516: uint16(sym_comment),
	4517: uint16(387),
	4518: uint16(1),
	4519: uint16(anon_sym_COMMA),
	4520: uint16(116),
	4521: uint16(1),
	4522: uint16(aux_sym_json_array_repeat1),
	4523: uint16(226),
	4524: uint16(2),
	4525: uint16(anon_sym_RPAREN),
	4526: uint16(anon_sym_RBRACK),
	4527: uint16(4),
	4528: uint16(3),
	4529: uint16(1),
	4530: uint16(sym_comment),
	4531: uint16(390),
	4532: uint16(1),
	4533: uint16(anon_sym_RBRACE),
	4534: uint16(392),
	4535: uint16(1),
	4536: uint16(anon_sym_COMMA),
	4537: uint16(117),
	4538: uint16(1),
	4539: uint16(aux_sym_json_object_repeat1),
	4540: uint16(4),
	4541: uint16(3),
	4542: uint16(1),
	4543: uint16(sym_comment),
	4544: uint16(224),
	4545: uint16(1),
	4546: uint16(anon_sym_COMMA),
	4547: uint16(395),
	4548: uint16(1),
	4549: uint16(anon_sym_RBRACE),
	4550: uint16(117),
	4551: uint16(1),
	4552: uint16(aux_sym_json_object_repeat1),
	4553: uint16(4),
	4554: uint16(3),
	4555: uint16(1),
	4556: uint16(sym_comment),
	4557: uint16(220),
	4558: uint16(1),
	4559: uint16(anon_sym_COMMA),
	4560: uint16(397),
	4561: uint16(1),
	4562: uint16(anon_sym_RPAREN),
	4563: uint16(116),
	4564: uint16(1),
	4565: uint16(aux_sym_json_array_repeat1),
	4566: uint16(4),
	4567: uint16(3),
	4568: uint16(1),
	4569: uint16(sym_comment),
	4570: uint16(220),
	4571: uint16(1),
	4572: uint16(anon_sym_COMMA),
	4573: uint16(399),
	4574: uint16(1),
	4575: uint16(anon_sym_RBRACK),
	4576: uint16(116),
	4577: uint16(1),
	4578: uint16(aux_sym_json_array_repeat1),
	4579: uint16(3),
	4580: uint16(3),
	4581: uint16(1),
	4582: uint16(sym_comment),
	4583: uint16(401),
	4584: uint16(1),
	4585: uint16(anon_sym_LBRACK),
	4586: uint16(403),
	4587: uint16(1),
	4588: uint16(sym_ident),
	4589: uint16(2),
	4590: uint16(3),
	4591: uint16(1),
	4592: uint16(sym_comment),
	4593: uint16(405),
	4594: uint16(1),
	4595: uint16(sym_symbol),
	4596: uint16(2),
	4597: uint16(3),
	4598: uint16(1),
	4599: uint16(sym_comment),
	4600: uint16(407),
	4601: uint16(1),
	4602: uint16(anon_sym_in),
	4603: uint16(2),
	4604: uint16(3),
	4605: uint16(1),
	4606: uint16(sym_comment),
	4607: uint16(409),
	4608: uint16(1),
	4610: uint16(2),
	4611: uint16(3),
	4612: uint16(1),
	4613: uint16(sym_comment),
	4614: uint16(411),
	4615: uint16(1),
	4616: uint16(sym_ident),
	4617: uint16(2),
	4618: uint16(3),
	4619: uint16(1),
	4620: uint16(sym_comment),
	4621: uint16(413),
	4622: uint16(1),
	4623: uint16(anon_sym_RPAREN),
}

var ts_small_parse_table_map = [125]uint32_t{
	1:   uint32(65),
	2:   uint32(130),
	3:   uint32(195),
	4:   uint32(257),
	5:   uint32(319),
	6:   uint32(381),
	7:   uint32(443),
	8:   uint32(505),
	9:   uint32(567),
	10:  uint32(629),
	11:  uint32(691),
	12:  uint32(753),
	13:  uint32(815),
	14:  uint32(877),
	15:  uint32(939),
	16:  uint32(1001),
	17:  uint32(1063),
	18:  uint32(1125),
	19:  uint32(1187),
	20:  uint32(1249),
	21:  uint32(1311),
	22:  uint32(1373),
	23:  uint32(1435),
	24:  uint32(1494),
	25:  uint32(1529),
	26:  uint32(1575),
	27:  uint32(1607),
	28:  uint32(1639),
	29:  uint32(1699),
	30:  uint32(1731),
	31:  uint32(1763),
	32:  uint32(1803),
	33:  uint32(1841),
	34:  uint32(1889),
	35:  uint32(1939),
	36:  uint32(1981),
	37:  uint32(2013),
	38:  uint32(2045),
	39:  uint32(2077),
	40:  uint32(2115),
	41:  uint32(2147),
	42:  uint32(2179),
	43:  uint32(2211),
	44:  uint32(2243),
	45:  uint32(2295),
	46:  uint32(2327),
	47:  uint32(2359),
	48:  uint32(2391),
	49:  uint32(2423),
	50:  uint32(2455),
	51:  uint32(2487),
	52:  uint32(2519),
	53:  uint32(2551),
	54:  uint32(2583),
	55:  uint32(2640),
	56:  uint32(2697),
	57:  uint32(2754),
	58:  uint32(2811),
	59:  uint32(2865),
	60:  uint32(2919),
	61:  uint32(2969),
	62:  uint32(3023),
	63:  uint32(3072),
	64:  uint32(3125),
	65:  uint32(3173),
	66:  uint32(3221),
	67:  uint32(3269),
	68:  uint32(3317),
	69:  uint32(3365),
	70:  uint32(3413),
	71:  uint32(3461),
	72:  uint32(3509),
	73:  uint32(3557),
	74:  uint32(3605),
	75:  uint32(3628),
	76:  uint32(3651),
	77:  uint32(3674),
	78:  uint32(3697),
	79:  uint32(3720),
	80:  uint32(3743),
	81:  uint32(3766),
	82:  uint32(3789),
	83:  uint32(3812),
	84:  uint32(3835),
	85:  uint32(3858),
	86:  uint32(3881),
	87:  uint32(3908),
	88:  uint32(3935),
	89:  uint32(3962),
	90:  uint32(3989),
	91:  uint32(4016),
	92:  uint32(4043),
	93:  uint32(4070),
	94:  uint32(4097),
	95:  uint32(4124),
	96:  uint32(4151),
	97:  uint32(4178),
	98:  uint32(4205),
	99:  uint32(4232),
	100: uint32(4259),
	101: uint32(4286),
	102: uint32(4307),
	103: uint32(4330),
	104: uint32(4351),
	105: uint32(4372),
	106: uint32(4393),
	107: uint32(4414),
	108: uint32(4435),
	109: uint32(4448),
	110: uint32(4461),
	111: uint32(4474),
	112: uint32(4487),
	113: uint32(4500),
	114: uint32(4513),
	115: uint32(4527),
	116: uint32(4540),
	117: uint32(4553),
	118: uint32(4566),
	119: uint32(4579),
	120: uint32(4589),
	121: uint32(4596),
	122: uint32(4603),
	123: uint32(4610),
	124: uint32(4617),
}

var ts_parse_actions = [415]TSParseActionEntry{
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(30)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(58)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(77)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(76)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(76)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(85)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(93)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(92)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(101)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(8)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(7)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(2)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(54)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(49)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(49)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(95)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(96)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(97)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(3)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(9)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(26)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(55)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(29)),
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
		Fcount: uint8(2),
	}})),
	60: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(77)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_source_file_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(30)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_source_file_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(58)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	69: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(77)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	72: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(76)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	75: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(76)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_source_file_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(85)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_source_file_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(93)),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_source_file_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(92)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_source_file_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(101)),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_source_file_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(8)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(4)),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_simplexpr),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_simplexpr),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_binary_expression),
		Fproduction_id: uint16(5),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(10)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(121)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(125)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(12)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(16)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_binary_expression),
		Fproduction_id: uint16(5),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_json_array),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_json_array),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_function_call),
		Fproduction_id: uint16(3),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_function_call),
		Fproduction_id: uint16(3),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(122)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(79)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_json_safe_dot_access),
		Fproduction_id: uint16(4),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_json_safe_dot_access),
		Fproduction_id: uint16(4),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_json_dot_access),
		Fproduction_id: uint16(4),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_json_dot_access),
		Fproduction_id: uint16(4),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(13)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_json_array),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_json_array),
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
		Fsymbol:        uint16(sym_function_call),
		Fproduction_id: uint16(3),
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
		Fsymbol:        uint16(sym_function_call),
		Fproduction_id: uint16(3),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_json_access),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_json_access),
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
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_unary_expression),
		Fproduction_id: uint16(2),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_unary_expression),
		Fproduction_id: uint16(2),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_json_object),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_json_object),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_function_call),
		Fproduction_id: uint16(3),
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
		Fcount: uint8(1),
	}})),
	163: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_function_call),
		Fproduction_id: uint16(3),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_json_safe_access),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_json_safe_access),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_ternary_expression),
		Fproduction_id: uint16(6),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(17)),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_json_object),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_json_object),
	})))),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_parenthesized_expression),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	183: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_number),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_number),
	})))),
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
		Fcount: uint8(1),
	}})),
	189: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_literal),
	})))),
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
		Fcount: uint8(1),
	}})),
	193: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_literal),
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
		Fsymbol:      uint16(sym_string),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_string),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_json_array),
	})))),
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
		Fcount: uint8(1),
	}})),
	205: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_json_array),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_json_object),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_json_object),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(83)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(80)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(39)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(18)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(42)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(21)),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_json_array_repeat1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(28)),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(aux_sym_json_object_repeat1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(115)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(112)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(20)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(45)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(48)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(22)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(19)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_ast_block),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_ast_block),
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
		Fsymbol:      uint16(sym_list),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_list),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_list),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_list),
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
		Fsymbol:      uint16(sym_array),
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
		Fsymbol:      uint16(sym_array),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_array),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_array),
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
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_loop_widget),
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
		Fcount: uint8(1),
	}})),
	275: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_loop_widget),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_expr),
	})))),
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
		Fcount: uint8(1),
	}})),
	279: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_expr),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_string_repeat4),
	})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_string_repeat4),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(23)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	286: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_string_repeat4),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(113)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_string_repeat4),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(113)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(82)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(110)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(110)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(23)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(113)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(113)),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_string_repeat6),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_string_repeat6),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(6)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	311: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_string_repeat6),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(110)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_string_repeat6),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(110)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(86)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(5)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(114)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(114)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(52)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(53)),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_string_repeat2),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_string_repeat2),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(5)),
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
		Fsymbol:      uint16(aux_sym_string_repeat2),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(114)),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_string_repeat2),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(114)),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_string_repeat5),
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
		Fcount: uint8(2),
	}})),
	342: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_string_repeat5),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(110)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_string_repeat5),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(110)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	348: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(65)),
	}})))),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(aux_sym_string_repeat6),
		Fproduction_id: uint16(1),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(aux_sym_string_repeat2),
		Fproduction_id: uint16(1),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_string_repeat3),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_string_repeat3),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(113)),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_string_repeat3),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(113)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	362: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(aux_sym_string_repeat4),
		Fproduction_id: uint16(1),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_string_repeat1),
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
		Fcount: uint8(2),
	}})),
	366: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(114)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_string_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(114)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_string_repeat5),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_string_repeat5),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_string_interpolation),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_string_interpolation),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_string_repeat3),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_string_repeat3),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	384: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_string_repeat1),
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
		Fcount: uint8(1),
	}})),
	386: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_string_repeat1),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	388: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_json_array_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(18)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_json_object_repeat1),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	393: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_json_object_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(21)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(47)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(44)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(38)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(24)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(31)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(123)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(104)),
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
	410: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f uint8_t
		_ [7]byte
	}{f: uint8(TSParseActionTypeAccept)})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(32)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(84)),
	}})))),
}

type ts_external_scanner_symbol_identifiers = int32

const ts_external_token__unescaped_single_quote_string_fragment = 0
const ts_external_token__unescaped_double_quote_string_fragment = 1
const ts_external_token__unescaped_backtick_string_fragment = 2

var ts_external_scanner_symbol_map = [3]TSSymbol{
	0: uint16(sym__unescaped_single_quote_string_fragment),
	1: uint16(sym__unescaped_double_quote_string_fragment),
	2: uint16(sym__unescaped_backtick_string_fragment),
}

var ts_external_scanner_states = [5][3]uint8{
	1: {
		0: libc.BoolUint8(true1 != 0),
		1: libc.BoolUint8(true1 != 0),
		2: libc.BoolUint8(true1 != 0),
	},
	2: {
		0: libc.BoolUint8(true1 != 0),
	},
	3: {
		2: libc.BoolUint8(true1 != 0),
	},
	4: {
		1: libc.BoolUint8(true1 != 0),
	},
}

func tree_sitter_yuck(tls *libc.TLS) (r uintptr) {
	return uintptr(unsafe.Pointer(&language))
}

var language = TSLanguage{
	Fabi_version:               uint32(LANGUAGE_VERSION),
	Fsymbol_count:              uint32(SYMBOL_COUNT),
	Falias_count:               uint32(ALIAS_COUNT),
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
	Fkeyword_capture_token:     uint16(sym_symbol),
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
	Fprimary_state_ids:     uintptr(unsafe.Pointer(&ts_primary_state_ids)),
	Fname:                  __ccgo_ts + 842,
	Fsupertype_count:       uint32(SUPERTYPE_COUNT),
	Fsupertype_symbols:     uintptr(unsafe.Pointer(&ts_supertype_symbols)),
	Fsupertype_map_slices:  uintptr(unsafe.Pointer(&ts_supertype_map_slices)),
	Fsupertype_map_entries: uintptr(unsafe.Pointer(&ts_supertype_map_entries)),
	Fmetadata: TSLanguageMetadata{
		Fpatch_version: uint8(2),
	},
}

func init() {
	p := unsafe.Pointer(&language)
	*(*uintptr)(unsafe.Add(p, 92)) = __ccgo_fp(ts_lex)
	*(*uintptr)(unsafe.Add(p, 96)) = __ccgo_fp(ts_lex_keywords)
	*(*uintptr)(unsafe.Add(p, 112)) = __ccgo_fp(tree_sitter_yuck_external_scanner_create)
	*(*uintptr)(unsafe.Add(p, 116)) = __ccgo_fp(tree_sitter_yuck_external_scanner_destroy)
	*(*uintptr)(unsafe.Add(p, 120)) = __ccgo_fp(tree_sitter_yuck_external_scanner_scan)
	*(*uintptr)(unsafe.Add(p, 124)) = __ccgo_fp(tree_sitter_yuck_external_scanner_serialize)
	*(*uintptr)(unsafe.Add(p, 128)) = __ccgo_fp(tree_sitter_yuck_external_scanner_deserialize)
}

func __ccgo_fp(f interface{}) uintptr {
	type iface [2]uintptr
	return (*iface)(unsafe.Pointer(&f))[1]
}

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = "end\x00symbol\x00(\x00for\x00in\x00)\x00[\x00]\x00keyword\x00integer\x00float\x00true\x00false\x00\"\x00'\x00`\x00${\x00}\x00_escape_sequence_token1\x00escape_sequence\x00{\x00,\x00:\x00?.\x00.\x00+\x00-\x00*\x00/\x00%\x00&&\x00||\x00==\x00!=\x00=~\x00>=\x00<=\x00>\x00<\x00?:\x00!\x00?\x00ident\x00comment\x00_unescaped_single_quote_string_fragment\x00_unescaped_double_quote_string_fragment\x00_unescaped_backtick_string_fragment\x00source_file\x00ast_block\x00loop_widget\x00list\x00array\x00literal\x00number\x00boolean\x00string\x00string_interpolation\x00_escape_sequence\x00expr\x00simplexpr\x00json_array\x00json_object\x00json_access\x00json_safe_access\x00json_dot_access\x00json_safe_dot_access\x00function_call\x00binary_expression\x00unary_expression\x00ternary_expression\x00parenthesized_expression\x00source_file_repeat1\x00string_repeat1\x00string_repeat2\x00string_repeat3\x00string_repeat4\x00string_repeat5\x00string_repeat6\x00json_array_repeat1\x00json_object_repeat1\x00index\x00string_fragment\x00alternative\x00argument\x00condition\x00consequence\x00left\x00name\x00operator\x00right\x00yuck\x00"
