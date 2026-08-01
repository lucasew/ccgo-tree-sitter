// Code generated for linux/386 by 'ccgo -ignore-unsupported-alignment -ignore-unsupported-atomic-sizes -ignore-vector-functions -std=gnu11 -O0 -Dfunc=func_token -Dinterface=interface_token -Dselect=select_token -Dchan=chan_token -Dgo=go_token -Dmap=map_token -Dpackage=package_token -Dtype=type_token -Dvar=var_token -Dimport=import_token -Ddefer=defer_token -Dfallthrough=fallthrough_token -Drange=range_token -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-udev/src -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-udev -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party -I /home/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/include -I /home/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/src /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-udev/src/parser.c -o grammar.go', DO NOT EDIT.

//go:build linux && 386

package grammar_udev

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
const FIELD_COUNT = 1
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
const MAX_ALIAS_SEQUENCE_LENGTH = 6
const MAX_RESERVED_WORD_SET_SIZE = 0
const PRODUCTION_ID_COUNT = 4
const PTRDIFF_MAX = "INT32_MAX"
const PTRDIFF_MIN = "INT32_MIN"
const RAND_MAX = 0x7fffffff
const SIG_ATOMIC_MAX = "INT32_MAX"
const SIG_ATOMIC_MIN = "INT32_MIN"
const SIZE_MAX = "UINT32_MAX"
const STATE_COUNT = 167
const SUPERTYPE_COUNT = 0
const SYMBOL_COUNT = 115
const TOKEN_COUNT = 93
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

/*
 *  Lexer Macros
 */

/*
 *  Parse Table Macros
 */

type ts_symbol_identifiers = int32

const aux_sym_rules_token1 = 1
const anon_sym_COMMA = 2
const anon_sym_ACTION = 3
const anon_sym_DEVPATH = 4
const anon_sym_KERNEL = 5
const anon_sym_KERNELS = 6
const anon_sym_NAME = 7
const anon_sym_SYMLINK = 8
const anon_sym_SUBSYSTEM = 9
const anon_sym_SUBSYSTEMS = 10
const anon_sym_DRIVER = 11
const anon_sym_DRIVERS = 12
const anon_sym_ATTR = 13
const anon_sym_LBRACE = 14
const anon_sym_RBRACE = 15
const anon_sym_ATTRS = 16
const anon_sym_SYSCTL = 17
const anon_sym_ENV = 18
const anon_sym_CONST = 19
const anon_sym_TAG = 20
const anon_sym_TAGS = 21
const anon_sym_TEST = 22
const anon_sym_PROGRAM = 23
const anon_sym_RESULT = 24
const anon_sym_OWNER = 25
const anon_sym_GROUP = 26
const anon_sym_MODE = 27
const anon_sym_SECLABEL = 28
const anon_sym_RUN = 29
const anon_sym_LABEL = 30
const anon_sym_GOTO = 31
const anon_sym_IMPORT = 32
const anon_sym_OPTIONS = 33
const sym_system_const = 34
const sym_run_type = 35
const sym_import_type = 36
const aux_sym_attribute_token1 = 37
const aux_sym_env_var_token1 = 38
const sym_seclabel = 39
const sym_octal = 40
const sym_number = 41
const sym_match_op = 42
const sym_assignment_op = 43
const anon_sym_DQUOTE = 44
const anon_sym_DQUOTE2 = 45
const anon_sym_e = 46
const aux_sym_content_token1 = 47
const anon_sym_BSLASH_DQUOTE = 48
const anon_sym_STAR = 49
const anon_sym_QMARK = 50
const anon_sym_PIPE = 51
const aux_sym_pattern_token1 = 52
const aux_sym_c_escape_token1 = 53
const aux_sym_c_escape_token2 = 54
const aux_sym_c_escape_token3 = 55
const aux_sym_c_escape_token4 = 56
const aux_sym_c_escape_token5 = 57
const anon_sym_PERCENTk = 58
const anon_sym_PERCENTn = 59
const anon_sym_PERCENTp = 60
const anon_sym_PERCENTb = 61
const anon_sym_PERCENTs = 62
const anon_sym_PERCENTE = 63
const anon_sym_PERCENTM = 64
const anon_sym_PERCENTm = 65
const anon_sym_PERCENTc = 66
const anon_sym_PLUS = 67
const anon_sym_RBRACE2 = 68
const anon_sym_PERCENTP = 69
const anon_sym_PERCENTr = 70
const anon_sym_PERCENTS = 71
const anon_sym_PERCENTN = 72
const anon_sym_PERCENT_PERCENT = 73
const anon_sym_DOLLARkernel = 74
const anon_sym_DOLLARnumber = 75
const anon_sym_DOLLARdevpath = 76
const anon_sym_DOLLARid = 77
const anon_sym_DOLLARdriver = 78
const anon_sym_DOLLARattr = 79
const anon_sym_DOLLARenv = 80
const anon_sym_DOLLARmajor = 81
const anon_sym_DOLLARminor = 82
const anon_sym_DOLLARresult = 83
const anon_sym_DOLLARparent = 84
const anon_sym_DOLLARname = 85
const anon_sym_DOLLARlinks = 86
const anon_sym_DOLLARroot = 87
const anon_sym_DOLLARsys = 88
const anon_sym_DOLLARdevnode = 89
const anon_sym_DOLLAR_DOLLAR = 90
const sym_linebreak = 91
const sym_comment = 92
const sym_rules = 93
const sym_rule = 94
const sym_match = 95
const sym_assignment = 96
const sym_attribute = 97
const sym_env_var = 98
const sym_kernel_param = 99
const sym_value = 100
const sym__sub_value = 101
const sym_content = 102
const aux_sym__sub_content = 103
const aux_sym__c_content = 104
const aux_sym__sub_c_content = 105
const sym_pattern = 106
const sym_c_escape = 107
const sym_fmt_sub = 108
const sym_var_sub = 109
const aux_sym_rules_repeat1 = 110
const aux_sym_rule_repeat1 = 111
const aux_sym_rule_repeat2 = 112
const aux_sym_attribute_repeat1 = 113
const aux_sym_content_repeat1 = 114

var ts_symbol_names = [115]uintptr{
	0:   __ccgo_ts,
	1:   __ccgo_ts + 4,
	2:   __ccgo_ts + 17,
	3:   __ccgo_ts + 19,
	4:   __ccgo_ts + 26,
	5:   __ccgo_ts + 34,
	6:   __ccgo_ts + 41,
	7:   __ccgo_ts + 49,
	8:   __ccgo_ts + 54,
	9:   __ccgo_ts + 62,
	10:  __ccgo_ts + 72,
	11:  __ccgo_ts + 83,
	12:  __ccgo_ts + 90,
	13:  __ccgo_ts + 98,
	14:  __ccgo_ts + 103,
	15:  __ccgo_ts + 105,
	16:  __ccgo_ts + 107,
	17:  __ccgo_ts + 113,
	18:  __ccgo_ts + 120,
	19:  __ccgo_ts + 124,
	20:  __ccgo_ts + 130,
	21:  __ccgo_ts + 134,
	22:  __ccgo_ts + 139,
	23:  __ccgo_ts + 144,
	24:  __ccgo_ts + 152,
	25:  __ccgo_ts + 159,
	26:  __ccgo_ts + 165,
	27:  __ccgo_ts + 171,
	28:  __ccgo_ts + 176,
	29:  __ccgo_ts + 185,
	30:  __ccgo_ts + 189,
	31:  __ccgo_ts + 195,
	32:  __ccgo_ts + 200,
	33:  __ccgo_ts + 207,
	34:  __ccgo_ts + 215,
	35:  __ccgo_ts + 228,
	36:  __ccgo_ts + 237,
	37:  __ccgo_ts + 249,
	38:  __ccgo_ts + 266,
	39:  __ccgo_ts + 281,
	40:  __ccgo_ts + 290,
	41:  __ccgo_ts + 296,
	42:  __ccgo_ts + 303,
	43:  __ccgo_ts + 312,
	44:  __ccgo_ts + 326,
	45:  __ccgo_ts + 326,
	46:  __ccgo_ts + 328,
	47:  __ccgo_ts + 330,
	48:  __ccgo_ts + 345,
	49:  __ccgo_ts + 348,
	50:  __ccgo_ts + 350,
	51:  __ccgo_ts + 352,
	52:  __ccgo_ts + 354,
	53:  __ccgo_ts + 369,
	54:  __ccgo_ts + 385,
	55:  __ccgo_ts + 401,
	56:  __ccgo_ts + 417,
	57:  __ccgo_ts + 433,
	58:  __ccgo_ts + 449,
	59:  __ccgo_ts + 452,
	60:  __ccgo_ts + 455,
	61:  __ccgo_ts + 458,
	62:  __ccgo_ts + 461,
	63:  __ccgo_ts + 464,
	64:  __ccgo_ts + 467,
	65:  __ccgo_ts + 470,
	66:  __ccgo_ts + 473,
	67:  __ccgo_ts + 476,
	68:  __ccgo_ts + 105,
	69:  __ccgo_ts + 478,
	70:  __ccgo_ts + 481,
	71:  __ccgo_ts + 484,
	72:  __ccgo_ts + 487,
	73:  __ccgo_ts + 490,
	74:  __ccgo_ts + 493,
	75:  __ccgo_ts + 501,
	76:  __ccgo_ts + 509,
	77:  __ccgo_ts + 518,
	78:  __ccgo_ts + 522,
	79:  __ccgo_ts + 530,
	80:  __ccgo_ts + 536,
	81:  __ccgo_ts + 541,
	82:  __ccgo_ts + 548,
	83:  __ccgo_ts + 555,
	84:  __ccgo_ts + 563,
	85:  __ccgo_ts + 571,
	86:  __ccgo_ts + 577,
	87:  __ccgo_ts + 584,
	88:  __ccgo_ts + 590,
	89:  __ccgo_ts + 595,
	90:  __ccgo_ts + 604,
	91:  __ccgo_ts + 607,
	92:  __ccgo_ts + 617,
	93:  __ccgo_ts + 625,
	94:  __ccgo_ts + 631,
	95:  __ccgo_ts + 636,
	96:  __ccgo_ts + 642,
	97:  __ccgo_ts + 653,
	98:  __ccgo_ts + 663,
	99:  __ccgo_ts + 671,
	100: __ccgo_ts + 684,
	101: __ccgo_ts + 684,
	102: __ccgo_ts + 690,
	103: __ccgo_ts + 698,
	104: __ccgo_ts + 711,
	105: __ccgo_ts + 722,
	106: __ccgo_ts + 737,
	107: __ccgo_ts + 745,
	108: __ccgo_ts + 754,
	109: __ccgo_ts + 762,
	110: __ccgo_ts + 770,
	111: __ccgo_ts + 784,
	112: __ccgo_ts + 797,
	113: __ccgo_ts + 810,
	114: __ccgo_ts + 828,
}

var ts_symbol_map = [115]TSSymbol{
	1:   uint16(aux_sym_rules_token1),
	2:   uint16(anon_sym_COMMA),
	3:   uint16(anon_sym_ACTION),
	4:   uint16(anon_sym_DEVPATH),
	5:   uint16(anon_sym_KERNEL),
	6:   uint16(anon_sym_KERNELS),
	7:   uint16(anon_sym_NAME),
	8:   uint16(anon_sym_SYMLINK),
	9:   uint16(anon_sym_SUBSYSTEM),
	10:  uint16(anon_sym_SUBSYSTEMS),
	11:  uint16(anon_sym_DRIVER),
	12:  uint16(anon_sym_DRIVERS),
	13:  uint16(anon_sym_ATTR),
	14:  uint16(anon_sym_LBRACE),
	15:  uint16(anon_sym_RBRACE),
	16:  uint16(anon_sym_ATTRS),
	17:  uint16(anon_sym_SYSCTL),
	18:  uint16(anon_sym_ENV),
	19:  uint16(anon_sym_CONST),
	20:  uint16(anon_sym_TAG),
	21:  uint16(anon_sym_TAGS),
	22:  uint16(anon_sym_TEST),
	23:  uint16(anon_sym_PROGRAM),
	24:  uint16(anon_sym_RESULT),
	25:  uint16(anon_sym_OWNER),
	26:  uint16(anon_sym_GROUP),
	27:  uint16(anon_sym_MODE),
	28:  uint16(anon_sym_SECLABEL),
	29:  uint16(anon_sym_RUN),
	30:  uint16(anon_sym_LABEL),
	31:  uint16(anon_sym_GOTO),
	32:  uint16(anon_sym_IMPORT),
	33:  uint16(anon_sym_OPTIONS),
	34:  uint16(sym_system_const),
	35:  uint16(sym_run_type),
	36:  uint16(sym_import_type),
	37:  uint16(aux_sym_attribute_token1),
	38:  uint16(aux_sym_env_var_token1),
	39:  uint16(sym_seclabel),
	40:  uint16(sym_octal),
	41:  uint16(sym_number),
	42:  uint16(sym_match_op),
	43:  uint16(sym_assignment_op),
	44:  uint16(anon_sym_DQUOTE),
	45:  uint16(anon_sym_DQUOTE),
	46:  uint16(anon_sym_e),
	47:  uint16(aux_sym_content_token1),
	48:  uint16(anon_sym_BSLASH_DQUOTE),
	49:  uint16(anon_sym_STAR),
	50:  uint16(anon_sym_QMARK),
	51:  uint16(anon_sym_PIPE),
	52:  uint16(aux_sym_pattern_token1),
	53:  uint16(aux_sym_c_escape_token1),
	54:  uint16(aux_sym_c_escape_token2),
	55:  uint16(aux_sym_c_escape_token3),
	56:  uint16(aux_sym_c_escape_token4),
	57:  uint16(aux_sym_c_escape_token5),
	58:  uint16(anon_sym_PERCENTk),
	59:  uint16(anon_sym_PERCENTn),
	60:  uint16(anon_sym_PERCENTp),
	61:  uint16(anon_sym_PERCENTb),
	62:  uint16(anon_sym_PERCENTs),
	63:  uint16(anon_sym_PERCENTE),
	64:  uint16(anon_sym_PERCENTM),
	65:  uint16(anon_sym_PERCENTm),
	66:  uint16(anon_sym_PERCENTc),
	67:  uint16(anon_sym_PLUS),
	68:  uint16(anon_sym_RBRACE),
	69:  uint16(anon_sym_PERCENTP),
	70:  uint16(anon_sym_PERCENTr),
	71:  uint16(anon_sym_PERCENTS),
	72:  uint16(anon_sym_PERCENTN),
	73:  uint16(anon_sym_PERCENT_PERCENT),
	74:  uint16(anon_sym_DOLLARkernel),
	75:  uint16(anon_sym_DOLLARnumber),
	76:  uint16(anon_sym_DOLLARdevpath),
	77:  uint16(anon_sym_DOLLARid),
	78:  uint16(anon_sym_DOLLARdriver),
	79:  uint16(anon_sym_DOLLARattr),
	80:  uint16(anon_sym_DOLLARenv),
	81:  uint16(anon_sym_DOLLARmajor),
	82:  uint16(anon_sym_DOLLARminor),
	83:  uint16(anon_sym_DOLLARresult),
	84:  uint16(anon_sym_DOLLARparent),
	85:  uint16(anon_sym_DOLLARname),
	86:  uint16(anon_sym_DOLLARlinks),
	87:  uint16(anon_sym_DOLLARroot),
	88:  uint16(anon_sym_DOLLARsys),
	89:  uint16(anon_sym_DOLLARdevnode),
	90:  uint16(anon_sym_DOLLAR_DOLLAR),
	91:  uint16(sym_linebreak),
	92:  uint16(sym_comment),
	93:  uint16(sym_rules),
	94:  uint16(sym_rule),
	95:  uint16(sym_match),
	96:  uint16(sym_assignment),
	97:  uint16(sym_attribute),
	98:  uint16(sym_env_var),
	99:  uint16(sym_kernel_param),
	100: uint16(sym_value),
	101: uint16(sym_value),
	102: uint16(sym_content),
	103: uint16(aux_sym__sub_content),
	104: uint16(aux_sym__c_content),
	105: uint16(aux_sym__sub_c_content),
	106: uint16(sym_pattern),
	107: uint16(sym_c_escape),
	108: uint16(sym_fmt_sub),
	109: uint16(sym_var_sub),
	110: uint16(aux_sym_rules_repeat1),
	111: uint16(aux_sym_rule_repeat1),
	112: uint16(aux_sym_rule_repeat2),
	113: uint16(aux_sym_attribute_repeat1),
	114: uint16(aux_sym_content_repeat1),
}

var ts_symbol_metadata = [115]TSSymbolMetadata{
	0: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	1: {},
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
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	35: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	36: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	37: {},
	38: {},
	39: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	45: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	46: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	47: {},
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
	52: {},
	53: {},
	54: {},
	55: {},
	56: {},
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
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	98: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
	103: {},
	104: {},
	105: {},
	106: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	107: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	108: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	109: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	110: {},
	111: {},
	112: {},
	113: {},
	114: {},
}

type ts_field_identifiers = int32

const field_key = 1

var ts_field_names = [2]uintptr{
	0: libc.UintptrFromInt32(0),
	1: __ccgo_ts + 844,
}

var ts_field_map_slices = [4]TSMapSlice{
	1: {
		Flength: uint16(1),
	},
}

var ts_field_map_entries = [1]TSFieldMapEntry{
	0: {
		Ffield_id: uint16(field_key),
	},
}

var ts_alias_sequences = [4][6]TSSymbol{
	0: {},
	2: {
		1: uint16(sym_content),
	},
	3: {
		2: uint16(sym_content),
	},
}

var ts_non_terminal_alias_map = [13]uint16_t{
	0:  uint16(aux_sym__sub_content),
	1:  uint16(2),
	2:  uint16(aux_sym__sub_content),
	3:  uint16(sym_content),
	4:  uint16(aux_sym__c_content),
	5:  uint16(2),
	6:  uint16(aux_sym__c_content),
	7:  uint16(sym_content),
	8:  uint16(aux_sym__sub_c_content),
	9:  uint16(2),
	10: uint16(aux_sym__sub_c_content),
	11: uint16(sym_content),
}

var ts_primary_state_ids = [167]TSStateId{
	1:   uint16(1),
	2:   uint16(2),
	3:   uint16(3),
	4:   uint16(4),
	5:   uint16(5),
	6:   uint16(6),
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
	19:  uint16(6),
	20:  uint16(5),
	21:  uint16(21),
	22:  uint16(12),
	23:  uint16(14),
	24:  uint16(15),
	25:  uint16(16),
	26:  uint16(18),
	27:  uint16(17),
	28:  uint16(11),
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
	42:  uint16(36),
	43:  uint16(38),
	44:  uint16(36),
	45:  uint16(5),
	46:  uint16(17),
	47:  uint16(16),
	48:  uint16(18),
	49:  uint16(15),
	50:  uint16(50),
	51:  uint16(51),
	52:  uint16(52),
	53:  uint16(13),
	54:  uint16(54),
	55:  uint16(18),
	56:  uint16(56),
	57:  uint16(57),
	58:  uint16(58),
	59:  uint16(18),
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
	73:  uint16(73),
	74:  uint16(74),
	75:  uint16(75),
	76:  uint16(76),
	77:  uint16(77),
	78:  uint16(78),
	79:  uint16(79),
	80:  uint16(80),
	81:  uint16(81),
	82:  uint16(82),
	83:  uint16(83),
	84:  uint16(84),
	85:  uint16(85),
	86:  uint16(86),
	87:  uint16(87),
	88:  uint16(88),
	89:  uint16(75),
	90:  uint16(90),
	91:  uint16(91),
	92:  uint16(92),
	93:  uint16(93),
	94:  uint16(94),
	95:  uint16(95),
	96:  uint16(96),
	97:  uint16(97),
	98:  uint16(85),
	99:  uint16(90),
	100: uint16(85),
	101: uint16(101),
	102: uint16(75),
	103: uint16(103),
	104: uint16(82),
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
	120: uint16(120),
	121: uint16(121),
	122: uint16(122),
	123: uint16(123),
	124: uint16(124),
	125: uint16(125),
	126: uint16(126),
	127: uint16(127),
	128: uint16(128),
	129: uint16(129),
	130: uint16(130),
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
	144: uint16(112),
	145: uint16(145),
	146: uint16(117),
	147: uint16(122),
	148: uint16(142),
	149: uint16(149),
	150: uint16(117),
	151: uint16(151),
	152: uint16(142),
	153: uint16(110),
	154: uint16(154),
	155: uint16(155),
	156: uint16(139),
	157: uint16(157),
	158: uint16(158),
	159: uint16(110),
	160: uint16(107),
	161: uint16(123),
	162: uint16(135),
	163: uint16(158),
	164: uint16(107),
	165: uint16(123),
	166: uint16(166),
}

var aux_sym_c_escape_token1_character_set_1 = [9]TSCharacterRange{
	0: {
		Fstart: int32('\''),
		Fend:   int32('\''),
	},
	1: {
		Fstart: int32('?'),
		Fend:   int32('?'),
	},
	2: {
		Fstart: int32('\\'),
		Fend:   int32('\\'),
	},
	3: {
		Fstart: int32('a'),
		Fend:   int32('b'),
	},
	4: {
		Fstart: int32('e'),
		Fend:   int32('f'),
	},
	5: {
		Fstart: int32('n'),
		Fend:   int32('n'),
	},
	6: {
		Fstart: int32('r'),
		Fend:   int32('r'),
	},
	7: {
		Fstart: int32('t'),
		Fend:   int32('t'),
	},
	8: {
		Fstart: int32('v'),
		Fend:   int32('v'),
	},
}

func ts_lex(tls *libc.TLS, lexer uintptr, state TSStateId) (r uint8) {
	var eof, result, skip, v4 uint8
	var half_size, i, i1, i10, i11, i12, i13, i14, i15, i16, i17, i2, i3, i4, i5, i6, i7, i8, i9, index, mid_index, size uint32_t
	var lookahead1, v3 int32_t
	var range_token, range_token1, v2 uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = eof, half_size, i, i1, i10, i11, i12, i13, i14, i15, i16, i17, i2, i3, i4, i5, i6, i7, i8, i9, index, lookahead1, mid_index, range_token, range_token1, result, size, skip, v2, v3, v4
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
			state = uint16(239)
			goto next_state
		}
		i = uint32(0)
		for {
			if !(i < libc.Uint32FromInt64(76)/libc.Uint32FromInt64(2)) {
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
		if lookahead1 == int32('\t') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(236)
			goto next_state
		}
		if int32('.') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(276)
			goto next_state
		}
		return result
	case int32(1):
		if lookahead1 == int32('\n') {
			state = uint16(344)
			goto next_state
		}
		return result
	case int32(2):
		if lookahead1 == int32('\n') {
			state = uint16(344)
			goto next_state
		}
		if lookahead1 == int32('"') {
			state = uint16(298)
			goto next_state
		}
		if lookahead1 == int32('U') {
			state = uint16(235)
			goto next_state
		}
		if lookahead1 == int32('u') {
			state = uint16(230)
			goto next_state
		}
		if lookahead1 == int32('x') {
			state = uint16(225)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(306)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&aux_sym_c_escape_token1_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(9) - index
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
			state = uint16(303)
			goto next_state
		}
		return result
	case int32(3):
		if lookahead1 == int32('!') {
			state = uint16(116)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32('!') && lookahead1 != int32('"') && lookahead1 != int32('[') && lookahead1 != int32(']') {
			state = uint16(116)
			goto next_state
		}
		return result
	case int32(4):
		i1 = uint32(0)
		for {
			if !(i1 < libc.Uint32FromInt64(44)/libc.Uint32FromInt64(2)) {
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
		if lookahead1 != 0 {
			state = uint16(288)
			goto next_state
		}
		return result
	case int32(5):
		i2 = uint32(0)
		for {
			if !(i2 < libc.Uint32FromInt64(40)/libc.Uint32FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token2[i2]) == lookahead1 {
				state = map_token2[i2+uint32(1)]
				goto next_state
			}
			goto _7
		_7:
			;
			i2 = i2 + uint32(2)
		}
		if lookahead1 != 0 {
			state = uint16(288)
			goto next_state
		}
		return result
	case int32(6):
		i3 = uint32(0)
		for {
			if !(i3 < libc.Uint32FromInt64(44)/libc.Uint32FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token3[i3]) == lookahead1 {
				state = map_token3[i3+uint32(1)]
				goto next_state
			}
			goto _8
		_8:
			;
			i3 = i3 + uint32(2)
		}
		if lookahead1 != 0 {
			state = uint16(288)
			goto next_state
		}
		return result
	case int32(7):
		i4 = uint32(0)
		for {
			if !(i4 < libc.Uint32FromInt64(40)/libc.Uint32FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token4[i4]) == lookahead1 {
				state = map_token4[i4+uint32(1)]
				goto next_state
			}
			goto _9
		_9:
			;
			i4 = i4 + uint32(2)
		}
		if lookahead1 != 0 {
			state = uint16(288)
			goto next_state
		}
		return result
	case int32(8):
		i5 = uint32(0)
		for {
			if !(i5 < libc.Uint32FromInt64(40)/libc.Uint32FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token5[i5]) == lookahead1 {
				state = map_token5[i5+uint32(1)]
				goto next_state
			}
			goto _10
		_10:
			;
			i5 = i5 + uint32(2)
		}
		if lookahead1 == int32('\t') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(12)
			goto next_state
		}
		if int32('.') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(276)
			goto next_state
		}
		return result
	case int32(9):
		i6 = uint32(0)
		for {
			if !(i6 < libc.Uint32FromInt64(32)/libc.Uint32FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token6[i6]) == lookahead1 {
				state = map_token6[i6+uint32(1)]
				goto next_state
			}
			goto _11
		_11:
			;
			i6 = i6 + uint32(2)
		}
		if lookahead1 != 0 {
			state = uint16(288)
			goto next_state
		}
		return result
	case int32(10):
		i7 = uint32(0)
		for {
			if !(i7 < libc.Uint32FromInt64(32)/libc.Uint32FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token7[i7]) == lookahead1 {
				state = map_token7[i7+uint32(1)]
				goto next_state
			}
			goto _12
		_12:
			;
			i7 = i7 + uint32(2)
		}
		if lookahead1 != 0 {
			state = uint16(288)
			goto next_state
		}
		return result
	case int32(11):
		i8 = uint32(0)
		for {
			if !(i8 < libc.Uint32FromInt64(48)/libc.Uint32FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token8[i8]) == lookahead1 {
				state = map_token8[i8+uint32(1)]
				goto next_state
			}
			goto _13
		_13:
			;
			i8 = i8 + uint32(2)
		}
		return result
	case int32(12):
		i9 = uint32(0)
		for {
			if !(i9 < libc.Uint32FromInt64(32)/libc.Uint32FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token9[i9]) == lookahead1 {
				state = map_token9[i9+uint32(1)]
				goto next_state
			}
			goto _14
		_14:
			;
			i9 = i9 + uint32(2)
		}
		if lookahead1 == int32('\t') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(12)
			goto next_state
		}
		if int32('.') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(276)
			goto next_state
		}
		return result
	case int32(13):
		i10 = uint32(0)
		for {
			if !(i10 < libc.Uint32FromInt64(56)/libc.Uint32FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token10[i10]) == lookahead1 {
				state = map_token10[i10+uint32(1)]
				goto next_state
			}
			goto _15
		_15:
			;
			i10 = i10 + uint32(2)
		}
		return result
	case int32(14):
		if lookahead1 == int32('+') {
			state = uint16(319)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(1)
			goto next_state
		}
		if lookahead1 == int32('}') {
			state = uint16(321)
			goto next_state
		}
		if lookahead1 == int32('\t') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(15)
			goto next_state
		}
		if lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(277)
			goto next_state
		}
		return result
	case int32(15):
		if lookahead1 == int32('+') {
			state = uint16(319)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(1)
			goto next_state
		}
		if lookahead1 == int32('\t') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(15)
			goto next_state
		}
		if lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(277)
			goto next_state
		}
		return result
	case int32(16):
		if lookahead1 == int32('0') {
			state = uint16(220)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(1)
			goto next_state
		}
		if lookahead1 == int32('\t') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(16)
			goto next_state
		}
		if int32('1') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(221)
			goto next_state
		}
		if int32('A') <= lookahead1 && lookahead1 <= int32('Z') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(278)
			goto next_state
		}
		return result
	case int32(17):
		if lookahead1 == int32('=') {
			state = uint16(282)
			goto next_state
		}
		return result
	case int32(18):
		if lookahead1 == int32('=') {
			state = uint16(283)
			goto next_state
		}
		return result
	case int32(19):
		if lookahead1 == int32('A') {
			state = uint16(25)
			goto next_state
		}
		return result
	case int32(20):
		if lookahead1 == int32('A') {
			state = uint16(44)
			goto next_state
		}
		if lookahead1 == int32('E') {
			state = uint16(95)
			goto next_state
		}
		return result
	case int32(21):
		if lookahead1 == int32('A') {
			state = uint16(63)
			goto next_state
		}
		return result
	case int32(22):
		if lookahead1 == int32('A') {
			state = uint16(60)
			goto next_state
		}
		return result
	case int32(23):
		if lookahead1 == int32('A') {
			state = uint16(100)
			goto next_state
		}
		return result
	case int32(24):
		if lookahead1 == int32('A') {
			state = uint16(27)
			goto next_state
		}
		return result
	case int32(25):
		if lookahead1 == int32('B') {
			state = uint16(38)
			goto next_state
		}
		return result
	case int32(26):
		if lookahead1 == int32('B') {
			state = uint16(92)
			goto next_state
		}
		return result
	case int32(27):
		if lookahead1 == int32('B') {
			state = uint16(43)
			goto next_state
		}
		return result
	case int32(28):
		if lookahead1 == int32('C') {
			state = uint16(103)
			goto next_state
		}
		if lookahead1 == int32('T') {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(29):
		if lookahead1 == int32('C') {
			state = uint16(58)
			goto next_state
		}
		return result
	case int32(30):
		if lookahead1 == int32('C') {
			state = uint16(107)
			goto next_state
		}
		return result
	case int32(31):
		if lookahead1 == int32('D') {
			state = uint16(35)
			goto next_state
		}
		return result
	case int32(32):
		if lookahead1 == int32('E') {
			state = uint16(112)
			goto next_state
		}
		if lookahead1 == int32('R') {
			state = uint16(47)
			goto next_state
		}
		return result
	case int32(33):
		if lookahead1 == int32('E') {
			state = uint16(94)
			goto next_state
		}
		if lookahead1 == int32('U') {
			state = uint16(64)
			goto next_state
		}
		return result
	case int32(34):
		if lookahead1 == int32('E') {
			state = uint16(29)
			goto next_state
		}
		if lookahead1 == int32('U') {
			state = uint16(26)
			goto next_state
		}
		if lookahead1 == int32('Y') {
			state = uint16(62)
			goto next_state
		}
		return result
	case int32(35):
		if lookahead1 == int32('E') {
			state = uint16(266)
			goto next_state
		}
		return result
	case int32(36):
		if lookahead1 == int32('E') {
			state = uint16(246)
			goto next_state
		}
		return result
	case int32(37):
		if lookahead1 == int32('E') {
			state = uint16(90)
			goto next_state
		}
		return result
	case int32(38):
		if lookahead1 == int32('E') {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(39):
		if lookahead1 == int32('E') {
			state = uint16(61)
			goto next_state
		}
		return result
	case int32(40):
		if lookahead1 == int32('E') {
			state = uint16(86)
			goto next_state
		}
		return result
	case int32(41):
		if lookahead1 == int32('E') {
			state = uint16(53)
			goto next_state
		}
		return result
	case int32(42):
		if lookahead1 == int32('E') {
			state = uint16(87)
			goto next_state
		}
		return result
	case int32(43):
		if lookahead1 == int32('E') {
			state = uint16(55)
			goto next_state
		}
		return result
	case int32(44):
		if lookahead1 == int32('G') {
			state = uint16(259)
			goto next_state
		}
		return result
	case int32(45):
		if lookahead1 == int32('G') {
			state = uint16(89)
			goto next_state
		}
		return result
	case int32(46):
		if lookahead1 == int32('H') {
			state = uint16(243)
			goto next_state
		}
		return result
	case int32(47):
		if lookahead1 == int32('I') {
			state = uint16(113)
			goto next_state
		}
		return result
	case int32(48):
		if lookahead1 == int32('I') {
			state = uint16(66)
			goto next_state
		}
		return result
	case int32(49):
		if lookahead1 == int32('I') {
			state = uint16(79)
			goto next_state
		}
		return result
	case int32(50):
		if lookahead1 == int32('I') {
			state = uint16(80)
			goto next_state
		}
		return result
	case int32(51):
		if lookahead1 == int32('K') {
			state = uint16(247)
			goto next_state
		}
		return result
	case int32(52):
		if lookahead1 == int32('L') {
			state = uint16(269)
			goto next_state
		}
		return result
	case int32(53):
		if lookahead1 == int32('L') {
			state = uint16(244)
			goto next_state
		}
		return result
	case int32(54):
		if lookahead1 == int32('L') {
			state = uint16(256)
			goto next_state
		}
		return result
	case int32(55):
		if lookahead1 == int32('L') {
			state = uint16(267)
			goto next_state
		}
		return result
	case int32(56):
		if lookahead1 == int32('L') {
			state = uint16(48)
			goto next_state
		}
		return result
	case int32(57):
		if lookahead1 == int32('L') {
			state = uint16(102)
			goto next_state
		}
		return result
	case int32(58):
		if lookahead1 == int32('L') {
			state = uint16(24)
			goto next_state
		}
		return result
	case int32(59):
		if lookahead1 == int32('M') {
			state = uint16(83)
			goto next_state
		}
		return result
	case int32(60):
		if lookahead1 == int32('M') {
			state = uint16(262)
			goto next_state
		}
		return result
	case int32(61):
		if lookahead1 == int32('M') {
			state = uint16(248)
			goto next_state
		}
		return result
	case int32(62):
		if lookahead1 == int32('M') {
			state = uint16(56)
			goto next_state
		}
		if lookahead1 == int32('S') {
			state = uint16(30)
			goto next_state
		}
		return result
	case int32(63):
		if lookahead1 == int32('M') {
			state = uint16(36)
			goto next_state
		}
		return result
	case int32(64):
		if lookahead1 == int32('N') {
			state = uint16(268)
			goto next_state
		}
		return result
	case int32(65):
		if lookahead1 == int32('N') {
			state = uint16(242)
			goto next_state
		}
		return result
	case int32(66):
		if lookahead1 == int32('N') {
			state = uint16(51)
			goto next_state
		}
		return result
	case int32(67):
		if lookahead1 == int32('N') {
			state = uint16(111)
			goto next_state
		}
		return result
	case int32(68):
		if lookahead1 == int32('N') {
			state = uint16(93)
			goto next_state
		}
		return result
	case int32(69):
		if lookahead1 == int32('N') {
			state = uint16(40)
			goto next_state
		}
		return result
	case int32(70):
		if lookahead1 == int32('N') {
			state = uint16(41)
			goto next_state
		}
		return result
	case int32(71):
		if lookahead1 == int32('N') {
			state = uint16(96)
			goto next_state
		}
		return result
	case int32(72):
		if lookahead1 == int32('O') {
			state = uint16(31)
			goto next_state
		}
		return result
	case int32(73):
		if lookahead1 == int32('O') {
			state = uint16(109)
			goto next_state
		}
		return result
	case int32(74):
		if lookahead1 == int32('O') {
			state = uint16(270)
			goto next_state
		}
		return result
	case int32(75):
		if lookahead1 == int32('O') {
			state = uint16(71)
			goto next_state
		}
		return result
	case int32(76):
		if lookahead1 == int32('O') {
			state = uint16(45)
			goto next_state
		}
		return result
	case int32(77):
		if lookahead1 == int32('O') {
			state = uint16(105)
			goto next_state
		}
		if lookahead1 == int32('R') {
			state = uint16(73)
			goto next_state
		}
		return result
	case int32(78):
		if lookahead1 == int32('O') {
			state = uint16(91)
			goto next_state
		}
		return result
	case int32(79):
		if lookahead1 == int32('O') {
			state = uint16(65)
			goto next_state
		}
		return result
	case int32(80):
		if lookahead1 == int32('O') {
			state = uint16(68)
			goto next_state
		}
		return result
	case int32(81):
		if lookahead1 == int32('P') {
			state = uint16(265)
			goto next_state
		}
		return result
	case int32(82):
		if lookahead1 == int32('P') {
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(83):
		if lookahead1 == int32('P') {
			state = uint16(78)
			goto next_state
		}
		return result
	case int32(84):
		if lookahead1 == int32('P') {
			state = uint16(108)
			goto next_state
		}
		if lookahead1 == int32('W') {
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(85):
		if lookahead1 == int32('R') {
			state = uint16(252)
			goto next_state
		}
		return result
	case int32(86):
		if lookahead1 == int32('R') {
			state = uint16(264)
			goto next_state
		}
		return result
	case int32(87):
		if lookahead1 == int32('R') {
			state = uint16(250)
			goto next_state
		}
		return result
	case int32(88):
		if lookahead1 == int32('R') {
			state = uint16(76)
			goto next_state
		}
		return result
	case int32(89):
		if lookahead1 == int32('R') {
			state = uint16(22)
			goto next_state
		}
		return result
	case int32(90):
		if lookahead1 == int32('R') {
			state = uint16(70)
			goto next_state
		}
		return result
	case int32(91):
		if lookahead1 == int32('R') {
			state = uint16(101)
			goto next_state
		}
		return result
	case int32(92):
		if lookahead1 == int32('S') {
			state = uint16(114)
			goto next_state
		}
		return result
	case int32(93):
		if lookahead1 == int32('S') {
			state = uint16(272)
			goto next_state
		}
		return result
	case int32(94):
		if lookahead1 == int32('S') {
			state = uint16(110)
			goto next_state
		}
		return result
	case int32(95):
		if lookahead1 == int32('S') {
			state = uint16(98)
			goto next_state
		}
		return result
	case int32(96):
		if lookahead1 == int32('S') {
			state = uint16(99)
			goto next_state
		}
		return result
	case int32(97):
		if lookahead1 == int32('S') {
			state = uint16(106)
			goto next_state
		}
		return result
	case int32(98):
		if lookahead1 == int32('T') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(99):
		if lookahead1 == int32('T') {
			state = uint16(258)
			goto next_state
		}
		return result
	case int32(100):
		if lookahead1 == int32('T') {
			state = uint16(46)
			goto next_state
		}
		return result
	case int32(101):
		if lookahead1 == int32('T') {
			state = uint16(271)
			goto next_state
		}
		return result
	case int32(102):
		if lookahead1 == int32('T') {
			state = uint16(263)
			goto next_state
		}
		return result
	case int32(103):
		if lookahead1 == int32('T') {
			state = uint16(49)
			goto next_state
		}
		return result
	case int32(104):
		if lookahead1 == int32('T') {
			state = uint16(85)
			goto next_state
		}
		return result
	case int32(105):
		if lookahead1 == int32('T') {
			state = uint16(74)
			goto next_state
		}
		return result
	case int32(106):
		if lookahead1 == int32('T') {
			state = uint16(39)
			goto next_state
		}
		return result
	case int32(107):
		if lookahead1 == int32('T') {
			state = uint16(54)
			goto next_state
		}
		return result
	case int32(108):
		if lookahead1 == int32('T') {
			state = uint16(50)
			goto next_state
		}
		return result
	case int32(109):
		if lookahead1 == int32('U') {
			state = uint16(81)
			goto next_state
		}
		return result
	case int32(110):
		if lookahead1 == int32('U') {
			state = uint16(57)
			goto next_state
		}
		return result
	case int32(111):
		if lookahead1 == int32('V') {
			state = uint16(257)
			goto next_state
		}
		return result
	case int32(112):
		if lookahead1 == int32('V') {
			state = uint16(82)
			goto next_state
		}
		return result
	case int32(113):
		if lookahead1 == int32('V') {
			state = uint16(42)
			goto next_state
		}
		return result
	case int32(114):
		if lookahead1 == int32('Y') {
			state = uint16(97)
			goto next_state
		}
		return result
	case int32(115):
		if lookahead1 == int32('\\') {
			state = uint16(1)
			goto next_state
		}
		if lookahead1 == int32('b') {
			state = uint16(212)
			goto next_state
		}
		if lookahead1 == int32('c') {
			state = uint16(167)
			goto next_state
		}
		if lookahead1 == int32('d') {
			state = uint16(124)
			goto next_state
		}
		if lookahead1 == int32('f') {
			state = uint16(147)
			goto next_state
		}
		if lookahead1 == int32('p') {
			state = uint16(123)
			goto next_state
		}
		if lookahead1 == int32('\t') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(115)
			goto next_state
		}
		return result
	case int32(116):
		if lookahead1 == int32(']') {
			state = uint16(302)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32('"') && lookahead1 != int32('[') {
			state = uint16(116)
			goto next_state
		}
		return result
	case int32(117):
		if lookahead1 == int32('a') {
			state = uint16(154)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(177)
			goto next_state
		}
		return result
	case int32(118):
		if lookahead1 == int32('a') {
			state = uint16(166)
			goto next_state
		}
		if lookahead1 == int32('u') {
			state = uint16(162)
			goto next_state
		}
		return result
	case int32(119):
		if lookahead1 == int32('a') {
			state = uint16(190)
			goto next_state
		}
		return result
	case int32(120):
		if lookahead1 == int32('a') {
			state = uint16(201)
			goto next_state
		}
		return result
	case int32(121):
		if lookahead1 == int32('a') {
			state = uint16(164)
			goto next_state
		}
		return result
	case int32(122):
		if lookahead1 == int32('a') {
			state = uint16(165)
			goto next_state
		}
		return result
	case int32(123):
		if lookahead1 == int32('a') {
			state = uint16(195)
			goto next_state
		}
		if lookahead1 == int32('r') {
			state = uint16(183)
			goto next_state
		}
		return result
	case int32(124):
		if lookahead1 == int32('b') {
			state = uint16(275)
			goto next_state
		}
		return result
	case int32(125):
		if lookahead1 == int32('b') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(126):
		if lookahead1 == int32('c') {
			state = uint16(144)
			goto next_state
		}
		return result
	case int32(127):
		if lookahead1 == int32('d') {
			state = uint16(330)
			goto next_state
		}
		return result
	case int32(128):
		if lookahead1 == int32('d') {
			state = uint16(160)
			goto next_state
		}
		return result
	case int32(129):
		if lookahead1 == int32('d') {
			state = uint16(134)
			goto next_state
		}
		return result
	case int32(130):
		if lookahead1 == int32('e') {
			state = uint16(213)
			goto next_state
		}
		if lookahead1 == int32('r') {
			state = uint16(146)
			goto next_state
		}
		return result
	case int32(131):
		if lookahead1 == int32('e') {
			state = uint16(192)
			goto next_state
		}
		return result
	case int32(132):
		if lookahead1 == int32('e') {
			state = uint16(197)
			goto next_state
		}
		if lookahead1 == int32('o') {
			state = uint16(180)
			goto next_state
		}
		return result
	case int32(133):
		if lookahead1 == int32('e') {
			state = uint16(338)
			goto next_state
		}
		return result
	case int32(134):
		if lookahead1 == int32('e') {
			state = uint16(342)
			goto next_state
		}
		return result
	case int32(135):
		if lookahead1 == int32('e') {
			state = uint16(275)
			goto next_state
		}
		return result
	case int32(136):
		if lookahead1 == int32('e') {
			state = uint16(156)
			goto next_state
		}
		return result
	case int32(137):
		if lookahead1 == int32('e') {
			state = uint16(174)
			goto next_state
		}
		return result
	case int32(138):
		if lookahead1 == int32('e') {
			state = uint16(187)
			goto next_state
		}
		return result
	case int32(139):
		if lookahead1 == int32('e') {
			state = uint16(176)
			goto next_state
		}
		return result
	case int32(140):
		if lookahead1 == int32('e') {
			state = uint16(188)
			goto next_state
		}
		return result
	case int32(141):
		if lookahead1 == int32('g') {
			state = uint16(191)
			goto next_state
		}
		return result
	case int32(142):
		if lookahead1 == int32('g') {
			state = uint16(196)
			goto next_state
		}
		return result
	case int32(143):
		if lookahead1 == int32('h') {
			state = uint16(329)
			goto next_state
		}
		return result
	case int32(144):
		if lookahead1 == int32('h') {
			state = uint16(273)
			goto next_state
		}
		return result
	case int32(145):
		if lookahead1 == int32('i') {
			state = uint16(168)
			goto next_state
		}
		return result
	case int32(146):
		if lookahead1 == int32('i') {
			state = uint16(216)
			goto next_state
		}
		return result
	case int32(147):
		if lookahead1 == int32('i') {
			state = uint16(158)
			goto next_state
		}
		return result
	case int32(148):
		if lookahead1 == int32('i') {
			state = uint16(169)
			goto next_state
		}
		return result
	case int32(149):
		if lookahead1 == int32('i') {
			state = uint16(170)
			goto next_state
		}
		return result
	case int32(150):
		if lookahead1 == int32('i') {
			state = uint16(194)
			goto next_state
		}
		return result
	case int32(151):
		if lookahead1 == int32('i') {
			state = uint16(159)
			goto next_state
		}
		return result
	case int32(152):
		if lookahead1 == int32('i') {
			state = uint16(175)
			goto next_state
		}
		return result
	case int32(153):
		if lookahead1 == int32('i') {
			state = uint16(161)
			goto next_state
		}
		return result
	case int32(154):
		if lookahead1 == int32('j') {
			state = uint16(181)
			goto next_state
		}
		return result
	case int32(155):
		if lookahead1 == int32('k') {
			state = uint16(199)
			goto next_state
		}
		return result
	case int32(156):
		if lookahead1 == int32('l') {
			state = uint16(327)
			goto next_state
		}
		return result
	case int32(157):
		if lookahead1 == int32('l') {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(158):
		if lookahead1 == int32('l') {
			state = uint16(135)
			goto next_state
		}
		return result
	case int32(159):
		if lookahead1 == int32('l') {
			state = uint16(208)
			goto next_state
		}
		return result
	case int32(160):
		if lookahead1 == int32('l') {
			state = uint16(152)
			goto next_state
		}
		return result
	case int32(161):
		if lookahead1 == int32('l') {
			state = uint16(209)
			goto next_state
		}
		return result
	case int32(162):
		if lookahead1 == int32('m') {
			state = uint16(125)
			goto next_state
		}
		return result
	case int32(163):
		if lookahead1 == int32('m') {
			state = uint16(273)
			goto next_state
		}
		return result
	case int32(164):
		if lookahead1 == int32('m') {
			state = uint16(274)
			goto next_state
		}
		return result
	case int32(165):
		if lookahead1 == int32('m') {
			state = uint16(275)
			goto next_state
		}
		return result
	case int32(166):
		if lookahead1 == int32('m') {
			state = uint16(133)
			goto next_state
		}
		return result
	case int32(167):
		if lookahead1 == int32('m') {
			state = uint16(128)
			goto next_state
		}
		return result
	case int32(168):
		if lookahead1 == int32('n') {
			state = uint16(155)
			goto next_state
		}
		return result
	case int32(169):
		if lookahead1 == int32('n') {
			state = uint16(274)
			goto next_state
		}
		return result
	case int32(170):
		if lookahead1 == int32('n') {
			state = uint16(275)
			goto next_state
		}
		return result
	case int32(171):
		if lookahead1 == int32('n') {
			state = uint16(214)
			goto next_state
		}
		return result
	case int32(172):
		if lookahead1 == int32('n') {
			state = uint16(179)
			goto next_state
		}
		if lookahead1 == int32('p') {
			state = uint16(120)
			goto next_state
		}
		return result
	case int32(173):
		if lookahead1 == int32('n') {
			state = uint16(136)
			goto next_state
		}
		return result
	case int32(174):
		if lookahead1 == int32('n') {
			state = uint16(202)
			goto next_state
		}
		return result
	case int32(175):
		if lookahead1 == int32('n') {
			state = uint16(135)
			goto next_state
		}
		return result
	case int32(176):
		if lookahead1 == int32('n') {
			state = uint16(205)
			goto next_state
		}
		return result
	case int32(177):
		if lookahead1 == int32('n') {
			state = uint16(182)
			goto next_state
		}
		return result
	case int32(178):
		if lookahead1 == int32('o') {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(179):
		if lookahead1 == int32('o') {
			state = uint16(129)
			goto next_state
		}
		return result
	case int32(180):
		if lookahead1 == int32('o') {
			state = uint16(200)
			goto next_state
		}
		return result
	case int32(181):
		if lookahead1 == int32('o') {
			state = uint16(185)
			goto next_state
		}
		return result
	case int32(182):
		if lookahead1 == int32('o') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(183):
		if lookahead1 == int32('o') {
			state = uint16(142)
			goto next_state
		}
		return result
	case int32(184):
		if lookahead1 == int32('r') {
			state = uint16(332)
			goto next_state
		}
		return result
	case int32(185):
		if lookahead1 == int32('r') {
			state = uint16(334)
			goto next_state
		}
		return result
	case int32(186):
		if lookahead1 == int32('r') {
			state = uint16(335)
			goto next_state
		}
		return result
	case int32(187):
		if lookahead1 == int32('r') {
			state = uint16(331)
			goto next_state
		}
		return result
	case int32(188):
		if lookahead1 == int32('r') {
			state = uint16(328)
			goto next_state
		}
		return result
	case int32(189):
		if lookahead1 == int32('r') {
			state = uint16(126)
			goto next_state
		}
		return result
	case int32(190):
		if lookahead1 == int32('r') {
			state = uint16(137)
			goto next_state
		}
		return result
	case int32(191):
		if lookahead1 == int32('r') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(192):
		if lookahead1 == int32('r') {
			state = uint16(173)
			goto next_state
		}
		return result
	case int32(193):
		if lookahead1 == int32('r') {
			state = uint16(178)
			goto next_state
		}
		return result
	case int32(194):
		if lookahead1 == int32('r') {
			state = uint16(204)
			goto next_state
		}
		return result
	case int32(195):
		if lookahead1 == int32('r') {
			state = uint16(139)
			goto next_state
		}
		return result
	case int32(196):
		if lookahead1 == int32('r') {
			state = uint16(122)
			goto next_state
		}
		return result
	case int32(197):
		if lookahead1 == int32('s') {
			state = uint16(210)
			goto next_state
		}
		return result
	case int32(198):
		if lookahead1 == int32('s') {
			state = uint16(341)
			goto next_state
		}
		return result
	case int32(199):
		if lookahead1 == int32('s') {
			state = uint16(339)
			goto next_state
		}
		return result
	case int32(200):
		if lookahead1 == int32('t') {
			state = uint16(340)
			goto next_state
		}
		return result
	case int32(201):
		if lookahead1 == int32('t') {
			state = uint16(143)
			goto next_state
		}
		return result
	case int32(202):
		if lookahead1 == int32('t') {
			state = uint16(337)
			goto next_state
		}
		return result
	case int32(203):
		if lookahead1 == int32('t') {
			state = uint16(336)
			goto next_state
		}
		return result
	case int32(204):
		if lookahead1 == int32('t') {
			state = uint16(273)
			goto next_state
		}
		return result
	case int32(205):
		if lookahead1 == int32('t') {
			state = uint16(275)
			goto next_state
		}
		return result
	case int32(206):
		if lookahead1 == int32('t') {
			state = uint16(207)
			goto next_state
		}
		return result
	case int32(207):
		if lookahead1 == int32('t') {
			state = uint16(184)
			goto next_state
		}
		return result
	case int32(208):
		if lookahead1 == int32('t') {
			state = uint16(148)
			goto next_state
		}
		return result
	case int32(209):
		if lookahead1 == int32('t') {
			state = uint16(149)
			goto next_state
		}
		return result
	case int32(210):
		if lookahead1 == int32('u') {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(211):
		if lookahead1 == int32('u') {
			state = uint16(151)
			goto next_state
		}
		return result
	case int32(212):
		if lookahead1 == int32('u') {
			state = uint16(153)
			goto next_state
		}
		return result
	case int32(213):
		if lookahead1 == int32('v') {
			state = uint16(172)
			goto next_state
		}
		return result
	case int32(214):
		if lookahead1 == int32('v') {
			state = uint16(333)
			goto next_state
		}
		return result
	case int32(215):
		if lookahead1 == int32('v') {
			state = uint16(163)
			goto next_state
		}
		return result
	case int32(216):
		if lookahead1 == int32('v') {
			state = uint16(138)
			goto next_state
		}
		return result
	case int32(217):
		if lookahead1 == int32('y') {
			state = uint16(198)
			goto next_state
		}
		return result
	case int32(218):
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(280)
			goto next_state
		}
		return result
	case int32(219):
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(279)
			goto next_state
		}
		return result
	case int32(220):
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(218)
			goto next_state
		}
		return result
	case int32(221):
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(219)
			goto next_state
		}
		return result
	case int32(222):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(307)
			goto next_state
		}
		return result
	case int32(223):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(308)
			goto next_state
		}
		return result
	case int32(224):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(309)
			goto next_state
		}
		return result
	case int32(225):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(222)
			goto next_state
		}
		return result
	case int32(226):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(223)
			goto next_state
		}
		return result
	case int32(227):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(224)
			goto next_state
		}
		return result
	case int32(228):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(226)
			goto next_state
		}
		return result
	case int32(229):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(227)
			goto next_state
		}
		return result
	case int32(230):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(228)
			goto next_state
		}
		return result
	case int32(231):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(229)
			goto next_state
		}
		return result
	case int32(232):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(231)
			goto next_state
		}
		return result
	case int32(233):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(232)
			goto next_state
		}
		return result
	case int32(234):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(233)
			goto next_state
		}
		return result
	case int32(235):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(234)
			goto next_state
		}
		return result
	case int32(236):
		if eof != 0 {
			state = uint16(239)
			goto next_state
		}
		i11 = uint32(0)
		for {
			if !(i11 < libc.Uint32FromInt64(72)/libc.Uint32FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token11[i11]) == lookahead1 {
				state = map_token11[i11+uint32(1)]
				goto next_state
			}
			goto _16
		_16:
			;
			i11 = i11 + uint32(2)
		}
		if lookahead1 == int32('\t') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(236)
			goto next_state
		}
		if int32('.') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(276)
			goto next_state
		}
		return result
	case int32(237):
		if eof != 0 {
			state = uint16(239)
			goto next_state
		}
		i12 = uint32(0)
		for {
			if !(i12 < libc.Uint32FromInt64(128)/libc.Uint32FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token12[i12]) == lookahead1 {
				state = map_token12[i12+uint32(1)]
				goto next_state
			}
			goto _17
		_17:
			;
			i12 = i12 + uint32(2)
		}
		if lookahead1 == int32('\t') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(238)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(281)
			goto next_state
		}
		return result
	case int32(238):
		if eof != 0 {
			state = uint16(239)
			goto next_state
		}
		i13 = uint32(0)
		for {
			if !(i13 < libc.Uint32FromInt64(124)/libc.Uint32FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token13[i13]) == lookahead1 {
				state = map_token13[i13+uint32(1)]
				goto next_state
			}
			goto _18
		_18:
			;
			i13 = i13 + uint32(2)
		}
		if lookahead1 == int32('\t') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(238)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(281)
			goto next_state
		}
		return result
	case int32(239):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(ts_builtin_sym_end)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(240):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_rules_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(241):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COMMA)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(242):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ACTION)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(243):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DEVPATH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(244):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_KERNEL)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('S') {
			state = uint16(245)
			goto next_state
		}
		return result
	case int32(245):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_KERNELS)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(246):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_NAME)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(247):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SYMLINK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(248):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SUBSYSTEM)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('S') {
			state = uint16(249)
			goto next_state
		}
		return result
	case int32(249):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SUBSYSTEMS)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(250):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DRIVER)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('S') {
			state = uint16(251)
			goto next_state
		}
		return result
	case int32(251):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DRIVERS)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(252):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ATTR)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('S') {
			state = uint16(255)
			goto next_state
		}
		return result
	case int32(253):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(254):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(255):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ATTRS)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(256):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SYSCTL)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(257):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ENV)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(258):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_CONST)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(259):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_TAG)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('S') {
			state = uint16(260)
			goto next_state
		}
		return result
	case int32(260):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_TAGS)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(261):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_TEST)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(262):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PROGRAM)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(263):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RESULT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(264):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_OWNER)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(265):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_GROUP)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(266):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_MODE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(267):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SECLABEL)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(268):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RUN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(269):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LABEL)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(270):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_GOTO)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(271):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_IMPORT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(272):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_OPTIONS)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(273):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_system_const)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(274):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_run_type)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(275):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_import_type)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(276):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_attribute_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(277):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_env_var_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(277)
			goto next_state
		}
		return result
	case int32(278):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_seclabel)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('A') <= lookahead1 && lookahead1 <= int32('Z') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(278)
			goto next_state
		}
		return result
	case int32(279):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_octal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(280):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_octal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(279)
			goto next_state
		}
		return result
	case int32(281):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(281)
			goto next_state
		}
		return result
	case int32(282):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_match_op)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(283):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_assignment_op)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(284):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_assignment_op)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('=') {
			state = uint16(282)
			goto next_state
		}
		return result
	case int32(285):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(286):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DQUOTE2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(287):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_e)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(288):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_content_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(289):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_content_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\n') {
			state = uint16(344)
			goto next_state
		}
		if lookahead1 == int32('"') {
			state = uint16(298)
			goto next_state
		}
		return result
	case int32(290):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_content_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\n') {
			state = uint16(344)
			goto next_state
		}
		if lookahead1 == int32('U') {
			state = uint16(235)
			goto next_state
		}
		if lookahead1 == int32('u') {
			state = uint16(230)
			goto next_state
		}
		if lookahead1 == int32('x') {
			state = uint16(225)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(306)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&aux_sym_c_escape_token1_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(9) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _22
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _22
	_22:
		if v4 != 0 {
			state = uint16(303)
			goto next_state
		}
		return result
	case int32(291):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_content_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('!') {
			state = uint16(116)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32('!') && lookahead1 != int32('"') && lookahead1 != int32('[') && lookahead1 != int32(']') {
			state = uint16(116)
			goto next_state
		}
		return result
	case int32(292):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_content_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i14 = uint32(0)
		for {
			if !(i14 < libc.Uint32FromInt64(48)/libc.Uint32FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token14[i14]) == lookahead1 {
				state = map_token14[i14+uint32(1)]
				goto next_state
			}
			goto _23
		_23:
			;
			i14 = i14 + uint32(2)
		}
		return result
	case int32(293):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_content_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i15 = uint32(0)
		for {
			if !(i15 < libc.Uint32FromInt64(36)/libc.Uint32FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token15[i15]) == lookahead1 {
				state = map_token15[i15+uint32(1)]
				goto next_state
			}
			goto _24
		_24:
			;
			i15 = i15 + uint32(2)
		}
		if lookahead1 != 0 && lookahead1 != int32('"') {
			state = uint16(288)
			goto next_state
		}
		return result
	case int32(294):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_content_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i16 = uint32(0)
		for {
			if !(i16 < libc.Uint32FromInt64(36)/libc.Uint32FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token16[i16]) == lookahead1 {
				state = map_token16[i16+uint32(1)]
				goto next_state
			}
			goto _25
		_25:
			;
			i16 = i16 + uint32(2)
		}
		if lookahead1 != 0 && lookahead1 != int32('"') {
			state = uint16(288)
			goto next_state
		}
		return result
	case int32(295):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_content_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i17 = uint32(0)
		for {
			if !(i17 < libc.Uint32FromInt64(56)/libc.Uint32FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token17[i17]) == lookahead1 {
				state = map_token17[i17+uint32(1)]
				goto next_state
			}
			goto _26
		_26:
			;
			i17 = i17 + uint32(2)
		}
		return result
	case int32(296):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_content_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('*') {
			state = uint16(299)
			goto next_state
		}
		if lookahead1 == int32('?') {
			state = uint16(300)
			goto next_state
		}
		if lookahead1 == int32('[') {
			state = uint16(291)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(290)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(301)
			goto next_state
		}
		if lookahead1 == int32('\t') || lookahead1 == int32(' ') {
			state = uint16(296)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32('"') {
			state = uint16(288)
			goto next_state
		}
		return result
	case int32(297):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_content_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('*') {
			state = uint16(299)
			goto next_state
		}
		if lookahead1 == int32('?') {
			state = uint16(300)
			goto next_state
		}
		if lookahead1 == int32('[') {
			state = uint16(291)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(289)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(301)
			goto next_state
		}
		if lookahead1 == int32('\t') || lookahead1 == int32(' ') {
			state = uint16(297)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32('"') {
			state = uint16(288)
			goto next_state
		}
		return result
	case int32(298):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BSLASH_DQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(299):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STAR)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(300):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_QMARK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(301):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PIPE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(302):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_pattern_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(303):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_c_escape_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(304):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_c_escape_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(305):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_c_escape_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(304)
			goto next_state
		}
		return result
	case int32(306):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_c_escape_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(305)
			goto next_state
		}
		return result
	case int32(307):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_c_escape_token3)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(308):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_c_escape_token4)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(309):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_c_escape_token5)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(310):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PERCENTk)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(311):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PERCENTn)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(312):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PERCENTp)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(313):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PERCENTb)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(314):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PERCENTs)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(315):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PERCENTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(316):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PERCENTM)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(317):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PERCENTm)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(318):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PERCENTc)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(319):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PLUS)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(320):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PLUS)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('=') {
			state = uint16(283)
			goto next_state
		}
		return result
	case int32(321):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACE2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(322):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PERCENTP)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(323):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PERCENTr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(324):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PERCENTS)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(325):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PERCENTN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(326):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PERCENT_PERCENT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(327):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOLLARkernel)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(328):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOLLARnumber)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(329):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOLLARdevpath)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(330):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOLLARid)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(331):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOLLARdriver)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(332):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOLLARattr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(333):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOLLARenv)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(334):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOLLARmajor)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(335):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOLLARminor)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(336):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOLLARresult)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(337):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOLLARparent)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(338):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOLLARname)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(339):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOLLARlinks)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(340):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOLLARroot)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(341):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOLLARsys)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(342):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOLLARdevnode)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(343):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOLLAR_DOLLAR)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(344):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_linebreak)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(345):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 != 0 && lookahead1 != int32('\n') {
			state = uint16(345)
			goto next_state
		}
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var map_token = [38]uint16_t{
	0:  uint16('\n'),
	1:  uint16(240),
	2:  uint16('!'),
	3:  uint16(17),
	4:  uint16('"'),
	5:  uint16(286),
	6:  uint16('#'),
	7:  uint16(345),
	8:  uint16('$'),
	9:  uint16(11),
	10: uint16('%'),
	11: uint16(13),
	12: uint16('*'),
	13: uint16(299),
	14: uint16('+'),
	15: uint16(320),
	16: uint16(','),
	17: uint16(241),
	18: uint16('-'),
	19: uint16(18),
	20: uint16(':'),
	21: uint16(18),
	22: uint16('='),
	23: uint16(284),
	24: uint16('?'),
	25: uint16(300),
	26: uint16('['),
	27: uint16(3),
	28: uint16('\\'),
	29: uint16(2),
	30: uint16('e'),
	31: uint16(287),
	32: uint16('{'),
	33: uint16(253),
	34: uint16('|'),
	35: uint16(301),
	36: uint16('}'),
	37: uint16(321),
}

var map_token1 = [22]uint16_t{
	0:  uint16('"'),
	1:  uint16(286),
	2:  uint16('$'),
	3:  uint16(292),
	4:  uint16('%'),
	5:  uint16(295),
	6:  uint16('*'),
	7:  uint16(299),
	8:  uint16('?'),
	9:  uint16(300),
	10: uint16('['),
	11: uint16(291),
	12: uint16('\\'),
	13: uint16(290),
	14: uint16('{'),
	15: uint16(253),
	16: uint16('|'),
	17: uint16(301),
	18: uint16('\t'),
	19: uint16(293),
	20: uint16(' '),
	21: uint16(293),
}

var map_token2 = [20]uint16_t{
	0:  uint16('"'),
	1:  uint16(286),
	2:  uint16('$'),
	3:  uint16(292),
	4:  uint16('%'),
	5:  uint16(295),
	6:  uint16('*'),
	7:  uint16(299),
	8:  uint16('?'),
	9:  uint16(300),
	10: uint16('['),
	11: uint16(291),
	12: uint16('\\'),
	13: uint16(290),
	14: uint16('|'),
	15: uint16(301),
	16: uint16('\t'),
	17: uint16(293),
	18: uint16(' '),
	19: uint16(293),
}

var map_token3 = [22]uint16_t{
	0:  uint16('"'),
	1:  uint16(286),
	2:  uint16('$'),
	3:  uint16(292),
	4:  uint16('%'),
	5:  uint16(295),
	6:  uint16('*'),
	7:  uint16(299),
	8:  uint16('?'),
	9:  uint16(300),
	10: uint16('['),
	11: uint16(291),
	12: uint16('\\'),
	13: uint16(289),
	14: uint16('{'),
	15: uint16(253),
	16: uint16('|'),
	17: uint16(301),
	18: uint16('\t'),
	19: uint16(294),
	20: uint16(' '),
	21: uint16(294),
}

var map_token4 = [20]uint16_t{
	0:  uint16('"'),
	1:  uint16(286),
	2:  uint16('$'),
	3:  uint16(292),
	4:  uint16('%'),
	5:  uint16(295),
	6:  uint16('*'),
	7:  uint16(299),
	8:  uint16('?'),
	9:  uint16(300),
	10: uint16('['),
	11: uint16(291),
	12: uint16('\\'),
	13: uint16(289),
	14: uint16('|'),
	15: uint16(301),
	16: uint16('\t'),
	17: uint16(294),
	18: uint16(' '),
	19: uint16(294),
}

var map_token5 = [20]uint16_t{
	0:  uint16('"'),
	1:  uint16(286),
	2:  uint16('%'),
	3:  uint16(13),
	4:  uint16('*'),
	5:  uint16(299),
	6:  uint16('+'),
	7:  uint16(319),
	8:  uint16('?'),
	9:  uint16(300),
	10: uint16('['),
	11: uint16(3),
	12: uint16('\\'),
	13: uint16(1),
	14: uint16('{'),
	15: uint16(253),
	16: uint16('|'),
	17: uint16(301),
	18: uint16('}'),
	19: uint16(254),
}

var map_token6 = [16]uint16_t{
	0:  uint16('"'),
	1:  uint16(286),
	2:  uint16('*'),
	3:  uint16(299),
	4:  uint16('?'),
	5:  uint16(300),
	6:  uint16('['),
	7:  uint16(291),
	8:  uint16('\\'),
	9:  uint16(290),
	10: uint16('|'),
	11: uint16(301),
	12: uint16('\t'),
	13: uint16(296),
	14: uint16(' '),
	15: uint16(296),
}

var map_token7 = [16]uint16_t{
	0:  uint16('"'),
	1:  uint16(286),
	2:  uint16('*'),
	3:  uint16(299),
	4:  uint16('?'),
	5:  uint16(300),
	6:  uint16('['),
	7:  uint16(291),
	8:  uint16('\\'),
	9:  uint16(289),
	10: uint16('|'),
	11: uint16(301),
	12: uint16('\t'),
	13: uint16(297),
	14: uint16(' '),
	15: uint16(297),
}

var map_token8 = [24]uint16_t{
	0:  uint16('$'),
	1:  uint16(343),
	2:  uint16('a'),
	3:  uint16(206),
	4:  uint16('d'),
	5:  uint16(130),
	6:  uint16('e'),
	7:  uint16(171),
	8:  uint16('i'),
	9:  uint16(127),
	10: uint16('k'),
	11: uint16(131),
	12: uint16('l'),
	13: uint16(145),
	14: uint16('m'),
	15: uint16(117),
	16: uint16('n'),
	17: uint16(118),
	18: uint16('p'),
	19: uint16(119),
	20: uint16('r'),
	21: uint16(132),
	22: uint16('s'),
	23: uint16(217),
}

var map_token9 = [16]uint16_t{
	0:  uint16('%'),
	1:  uint16(13),
	2:  uint16('*'),
	3:  uint16(299),
	4:  uint16('+'),
	5:  uint16(319),
	6:  uint16('?'),
	7:  uint16(300),
	8:  uint16('['),
	9:  uint16(3),
	10: uint16('\\'),
	11: uint16(1),
	12: uint16('|'),
	13: uint16(301),
	14: uint16('}'),
	15: uint16(254),
}

var map_token10 = [28]uint16_t{
	0:  uint16('%'),
	1:  uint16(326),
	2:  uint16('E'),
	3:  uint16(315),
	4:  uint16('M'),
	5:  uint16(316),
	6:  uint16('N'),
	7:  uint16(325),
	8:  uint16('P'),
	9:  uint16(322),
	10: uint16('S'),
	11: uint16(324),
	12: uint16('b'),
	13: uint16(313),
	14: uint16('c'),
	15: uint16(318),
	16: uint16('k'),
	17: uint16(310),
	18: uint16('m'),
	19: uint16(317),
	20: uint16('n'),
	21: uint16(311),
	22: uint16('p'),
	23: uint16(312),
	24: uint16('r'),
	25: uint16(323),
	26: uint16('s'),
	27: uint16(314),
}

var map_token11 = [36]uint16_t{
	0:  uint16('\n'),
	1:  uint16(240),
	2:  uint16('!'),
	3:  uint16(17),
	4:  uint16('"'),
	5:  uint16(285),
	6:  uint16('#'),
	7:  uint16(345),
	8:  uint16('$'),
	9:  uint16(11),
	10: uint16('%'),
	11: uint16(13),
	12: uint16('*'),
	13: uint16(299),
	14: uint16('+'),
	15: uint16(320),
	16: uint16(','),
	17: uint16(241),
	18: uint16('-'),
	19: uint16(18),
	20: uint16(':'),
	21: uint16(18),
	22: uint16('='),
	23: uint16(284),
	24: uint16('?'),
	25: uint16(300),
	26: uint16('['),
	27: uint16(3),
	28: uint16('\\'),
	29: uint16(2),
	30: uint16('e'),
	31: uint16(287),
	32: uint16('|'),
	33: uint16(301),
	34: uint16('}'),
	35: uint16(254),
}

var map_token12 = [64]uint16_t{
	0:  uint16('\n'),
	1:  uint16(240),
	2:  uint16('!'),
	3:  uint16(17),
	4:  uint16('"'),
	5:  uint16(285),
	6:  uint16('#'),
	7:  uint16(345),
	8:  uint16('+'),
	9:  uint16(18),
	10: uint16('-'),
	11: uint16(18),
	12: uint16(':'),
	13: uint16(18),
	14: uint16('='),
	15: uint16(284),
	16: uint16('A'),
	17: uint16(28),
	18: uint16('C'),
	19: uint16(75),
	20: uint16('D'),
	21: uint16(32),
	22: uint16('E'),
	23: uint16(67),
	24: uint16('G'),
	25: uint16(77),
	26: uint16('I'),
	27: uint16(59),
	28: uint16('K'),
	29: uint16(37),
	30: uint16('L'),
	31: uint16(19),
	32: uint16('M'),
	33: uint16(72),
	34: uint16('N'),
	35: uint16(21),
	36: uint16('O'),
	37: uint16(84),
	38: uint16('P'),
	39: uint16(88),
	40: uint16('R'),
	41: uint16(33),
	42: uint16('S'),
	43: uint16(34),
	44: uint16('T'),
	45: uint16(20),
	46: uint16('\\'),
	47: uint16(1),
	48: uint16('a'),
	49: uint16(189),
	50: uint16('b'),
	51: uint16(211),
	52: uint16('c'),
	53: uint16(215),
	54: uint16('e'),
	55: uint16(287),
	56: uint16('p'),
	57: uint16(193),
	58: uint16('v'),
	59: uint16(150),
	60: uint16('{'),
	61: uint16(253),
	62: uint16('}'),
	63: uint16(254),
}

var map_token13 = [62]uint16_t{
	0:  uint16('\n'),
	1:  uint16(240),
	2:  uint16('!'),
	3:  uint16(17),
	4:  uint16('"'),
	5:  uint16(285),
	6:  uint16('#'),
	7:  uint16(345),
	8:  uint16('+'),
	9:  uint16(18),
	10: uint16('-'),
	11: uint16(18),
	12: uint16(':'),
	13: uint16(18),
	14: uint16('='),
	15: uint16(284),
	16: uint16('A'),
	17: uint16(28),
	18: uint16('C'),
	19: uint16(75),
	20: uint16('D'),
	21: uint16(32),
	22: uint16('E'),
	23: uint16(67),
	24: uint16('G'),
	25: uint16(77),
	26: uint16('I'),
	27: uint16(59),
	28: uint16('K'),
	29: uint16(37),
	30: uint16('L'),
	31: uint16(19),
	32: uint16('M'),
	33: uint16(72),
	34: uint16('N'),
	35: uint16(21),
	36: uint16('O'),
	37: uint16(84),
	38: uint16('P'),
	39: uint16(88),
	40: uint16('R'),
	41: uint16(33),
	42: uint16('S'),
	43: uint16(34),
	44: uint16('T'),
	45: uint16(20),
	46: uint16('\\'),
	47: uint16(1),
	48: uint16('a'),
	49: uint16(189),
	50: uint16('b'),
	51: uint16(211),
	52: uint16('c'),
	53: uint16(215),
	54: uint16('e'),
	55: uint16(287),
	56: uint16('p'),
	57: uint16(193),
	58: uint16('v'),
	59: uint16(150),
	60: uint16('}'),
	61: uint16(254),
}

var map_token14 = [24]uint16_t{
	0:  uint16('$'),
	1:  uint16(343),
	2:  uint16('a'),
	3:  uint16(206),
	4:  uint16('d'),
	5:  uint16(130),
	6:  uint16('e'),
	7:  uint16(171),
	8:  uint16('i'),
	9:  uint16(127),
	10: uint16('k'),
	11: uint16(131),
	12: uint16('l'),
	13: uint16(145),
	14: uint16('m'),
	15: uint16(117),
	16: uint16('n'),
	17: uint16(118),
	18: uint16('p'),
	19: uint16(119),
	20: uint16('r'),
	21: uint16(132),
	22: uint16('s'),
	23: uint16(217),
}

var map_token15 = [18]uint16_t{
	0:  uint16('$'),
	1:  uint16(292),
	2:  uint16('%'),
	3:  uint16(295),
	4:  uint16('*'),
	5:  uint16(299),
	6:  uint16('?'),
	7:  uint16(300),
	8:  uint16('['),
	9:  uint16(291),
	10: uint16('\\'),
	11: uint16(290),
	12: uint16('|'),
	13: uint16(301),
	14: uint16('\t'),
	15: uint16(293),
	16: uint16(' '),
	17: uint16(293),
}

var map_token16 = [18]uint16_t{
	0:  uint16('$'),
	1:  uint16(292),
	2:  uint16('%'),
	3:  uint16(295),
	4:  uint16('*'),
	5:  uint16(299),
	6:  uint16('?'),
	7:  uint16(300),
	8:  uint16('['),
	9:  uint16(291),
	10: uint16('\\'),
	11: uint16(289),
	12: uint16('|'),
	13: uint16(301),
	14: uint16('\t'),
	15: uint16(294),
	16: uint16(' '),
	17: uint16(294),
}

var map_token17 = [28]uint16_t{
	0:  uint16('%'),
	1:  uint16(326),
	2:  uint16('E'),
	3:  uint16(315),
	4:  uint16('M'),
	5:  uint16(316),
	6:  uint16('N'),
	7:  uint16(325),
	8:  uint16('P'),
	9:  uint16(322),
	10: uint16('S'),
	11: uint16(324),
	12: uint16('b'),
	13: uint16(313),
	14: uint16('c'),
	15: uint16(318),
	16: uint16('k'),
	17: uint16(310),
	18: uint16('m'),
	19: uint16(317),
	20: uint16('n'),
	21: uint16(311),
	22: uint16('p'),
	23: uint16(312),
	24: uint16('r'),
	25: uint16(323),
	26: uint16('s'),
	27: uint16(314),
}

var ts_lex_modes = [167]TSLexerMode{
	0: {},
	1: {
		Flex_state: uint16(237),
	},
	2: {
		Flex_state: uint16(5),
	},
	3: {
		Flex_state: uint16(5),
	},
	4: {
		Flex_state: uint16(5),
	},
	5: {
		Flex_state: uint16(4),
	},
	6: {
		Flex_state: uint16(4),
	},
	7: {
		Flex_state: uint16(7),
	},
	8: {
		Flex_state: uint16(5),
	},
	9: {
		Flex_state: uint16(7),
	},
	10: {
		Flex_state: uint16(7),
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
		Flex_state: uint16(5),
	},
	16: {
		Flex_state: uint16(5),
	},
	17: {
		Flex_state: uint16(5),
	},
	18: {
		Flex_state: uint16(5),
	},
	19: {
		Flex_state: uint16(6),
	},
	20: {
		Flex_state: uint16(6),
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
		Flex_state: uint16(7),
	},
	28: {
		Flex_state: uint16(7),
	},
	29: {
		Flex_state: uint16(237),
	},
	30: {
		Flex_state: uint16(237),
	},
	31: {
		Flex_state: uint16(237),
	},
	32: {
		Flex_state: uint16(237),
	},
	33: {
		Flex_state: uint16(237),
	},
	34: {
		Flex_state: uint16(237),
	},
	35: {
		Flex_state: uint16(237),
	},
	36: {
		Flex_state: uint16(8),
	},
	37: {
		Flex_state: uint16(8),
	},
	38: {
		Flex_state: uint16(8),
	},
	39: {
		Flex_state: uint16(8),
	},
	40: {
		Flex_state: uint16(8),
	},
	41: {
		Flex_state: uint16(8),
	},
	42: {
		Flex_state: uint16(8),
	},
	43: {
		Flex_state: uint16(8),
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
		Flex_state: uint16(8),
	},
	49: {
		Flex_state: uint16(8),
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
		Flex_state: uint16(10),
	},
	57: {
		Flex_state: uint16(10),
	},
	58: {
		Flex_state: uint16(10),
	},
	59: {
		Flex_state: uint16(10),
	},
	60: {},
	61: {},
	62: {},
	63: {
		Flex_state: uint16(237),
	},
	64: {
		Flex_state: uint16(237),
	},
	65: {
		Flex_state: uint16(237),
	},
	66: {
		Flex_state: uint16(237),
	},
	67: {},
	68: {
		Flex_state: uint16(237),
	},
	69: {
		Flex_state: uint16(237),
	},
	70: {},
	71: {
		Flex_state: uint16(237),
	},
	72: {
		Flex_state: uint16(237),
	},
	73: {
		Flex_state: uint16(237),
	},
	74: {
		Flex_state: uint16(237),
	},
	75: {
		Flex_state: uint16(14),
	},
	76: {
		Flex_state: uint16(237),
	},
	77: {
		Flex_state: uint16(237),
	},
	78: {
		Flex_state: uint16(237),
	},
	79: {
		Flex_state: uint16(14),
	},
	80: {
		Flex_state: uint16(14),
	},
	81: {},
	82: {
		Flex_state: uint16(14),
	},
	83: {},
	84: {},
	85: {
		Flex_state: uint16(14),
	},
	86: {},
	87: {
		Flex_state: uint16(237),
	},
	88: {},
	89: {
		Flex_state: uint16(14),
	},
	90: {
		Flex_state: uint16(8),
	},
	91: {},
	92: {},
	93: {},
	94: {},
	95: {
		Flex_state: uint16(237),
	},
	96: {},
	97: {},
	98: {
		Flex_state: uint16(14),
	},
	99: {
		Flex_state: uint16(8),
	},
	100: {
		Flex_state: uint16(14),
	},
	101: {
		Flex_state: uint16(14),
	},
	102: {
		Flex_state: uint16(14),
	},
	103: {},
	104: {
		Flex_state: uint16(14),
	},
	105: {
		Flex_state: uint16(14),
	},
	106: {},
	107: {},
	108: {
		Flex_state: uint16(8),
	},
	109: {
		Flex_state: uint16(237),
	},
	110: {
		Flex_state: uint16(237),
	},
	111: {
		Flex_state: uint16(16),
	},
	112: {
		Flex_state: uint16(237),
	},
	113: {
		Flex_state: uint16(237),
	},
	114: {},
	115: {},
	116: {},
	117: {
		Flex_state: uint16(14),
	},
	118: {},
	119: {
		Flex_state: uint16(237),
	},
	120: {},
	121: {
		Flex_state: uint16(237),
	},
	122: {
		Flex_state: uint16(237),
	},
	123: {},
	124: {},
	125: {
		Flex_state: uint16(237),
	},
	126: {},
	127: {
		Flex_state: uint16(237),
	},
	128: {
		Flex_state: uint16(237),
	},
	129: {
		Flex_state: uint16(237),
	},
	130: {
		Flex_state: uint16(237),
	},
	131: {
		Flex_state: uint16(237),
	},
	132: {
		Flex_state: uint16(237),
	},
	133: {},
	134: {},
	135: {},
	136: {},
	137: {},
	138: {},
	139: {
		Flex_state: uint16(237),
	},
	140: {
		Flex_state: uint16(8),
	},
	141: {},
	142: {
		Flex_state: uint16(237),
	},
	143: {},
	144: {
		Flex_state: uint16(237),
	},
	145: {
		Flex_state: uint16(16),
	},
	146: {
		Flex_state: uint16(14),
	},
	147: {
		Flex_state: uint16(237),
	},
	148: {
		Flex_state: uint16(237),
	},
	149: {},
	150: {
		Flex_state: uint16(14),
	},
	151: {},
	152: {
		Flex_state: uint16(237),
	},
	153: {
		Flex_state: uint16(237),
	},
	154: {
		Flex_state: uint16(237),
	},
	155: {
		Flex_state: uint16(8),
	},
	156: {
		Flex_state: uint16(237),
	},
	157: {
		Flex_state: uint16(237),
	},
	158: {},
	159: {
		Flex_state: uint16(237),
	},
	160: {},
	161: {},
	162: {},
	163: {},
	164: {},
	165: {},
	166: {
		Flex_state: uint16(115),
	},
}

var ts_parse_table = [2][115]uint16_t{
	0: {
		0:  uint16(1),
		1:  uint16(1),
		2:  uint16(1),
		14: uint16(1),
		15: uint16(1),
		37: uint16(1),
		42: uint16(1),
		43: uint16(1),
		44: uint16(1),
		45: uint16(1),
		46: uint16(1),
		48: uint16(1),
		49: uint16(1),
		50: uint16(1),
		51: uint16(1),
		52: uint16(1),
		53: uint16(1),
		54: uint16(1),
		55: uint16(1),
		56: uint16(1),
		57: uint16(1),
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
		76: uint16(1),
		77: uint16(1),
		78: uint16(1),
		79: uint16(1),
		80: uint16(1),
		81: uint16(1),
		82: uint16(1),
		83: uint16(1),
		84: uint16(1),
		85: uint16(1),
		86: uint16(1),
		87: uint16(1),
		88: uint16(1),
		89: uint16(1),
		90: uint16(1),
		91: uint16(3),
		92: uint16(1),
	},
	1: {
		0:   uint16(5),
		1:   uint16(7),
		3:   uint16(9),
		4:   uint16(9),
		5:   uint16(11),
		6:   uint16(9),
		7:   uint16(13),
		8:   uint16(13),
		9:   uint16(11),
		10:  uint16(9),
		11:  uint16(11),
		12:  uint16(9),
		13:  uint16(15),
		16:  uint16(17),
		17:  uint16(19),
		18:  uint16(21),
		19:  uint16(23),
		20:  uint16(25),
		21:  uint16(9),
		22:  uint16(27),
		23:  uint16(29),
		24:  uint16(9),
		25:  uint16(31),
		26:  uint16(31),
		27:  uint16(31),
		28:  uint16(33),
		29:  uint16(35),
		30:  uint16(37),
		31:  uint16(37),
		32:  uint16(39),
		33:  uint16(37),
		91:  uint16(3),
		92:  uint16(41),
		93:  uint16(149),
		94:  uint16(143),
		95:  uint16(151),
		96:  uint16(70),
		110: uint16(30),
		111: uint16(31),
	},
}

var ts_small_parse_table = [3587]uint16_t{
	0:    uint16(15),
	1:    uint16(43),
	2:    uint16(1),
	3:    uint16(anon_sym_DQUOTE2),
	4:    uint16(45),
	5:    uint16(1),
	6:    uint16(aux_sym_content_token1),
	7:    uint16(53),
	8:    uint16(1),
	9:    uint16(anon_sym_PERCENTs),
	10:   uint16(55),
	11:   uint16(1),
	12:   uint16(anon_sym_PERCENTE),
	13:   uint16(57),
	14:   uint16(1),
	15:   uint16(anon_sym_PERCENTc),
	16:   uint16(61),
	17:   uint16(1),
	18:   uint16(anon_sym_DOLLARattr),
	19:   uint16(63),
	20:   uint16(1),
	21:   uint16(anon_sym_DOLLARenv),
	22:   uint16(65),
	23:   uint16(1),
	24:   uint16(anon_sym_DOLLARresult),
	25:   uint16(67),
	26:   uint16(1),
	27:   uint16(sym_linebreak),
	28:   uint16(3),
	29:   uint16(1),
	30:   uint16(aux_sym__sub_c_content),
	31:   uint16(47),
	32:   uint16(4),
	33:   uint16(anon_sym_STAR),
	34:   uint16(anon_sym_QMARK),
	35:   uint16(anon_sym_PIPE),
	36:   uint16(aux_sym_pattern_token1),
	37:   uint16(8),
	38:   uint16(4),
	39:   uint16(sym_pattern),
	40:   uint16(sym_c_escape),
	41:   uint16(sym_fmt_sub),
	42:   uint16(sym_var_sub),
	43:   uint16(49),
	44:   uint16(5),
	45:   uint16(aux_sym_c_escape_token1),
	46:   uint16(aux_sym_c_escape_token2),
	47:   uint16(aux_sym_c_escape_token3),
	48:   uint16(aux_sym_c_escape_token4),
	49:   uint16(aux_sym_c_escape_token5),
	50:   uint16(51),
	51:   uint16(11),
	52:   uint16(anon_sym_PERCENTk),
	53:   uint16(anon_sym_PERCENTn),
	54:   uint16(anon_sym_PERCENTp),
	55:   uint16(anon_sym_PERCENTb),
	56:   uint16(anon_sym_PERCENTM),
	57:   uint16(anon_sym_PERCENTm),
	58:   uint16(anon_sym_PERCENTP),
	59:   uint16(anon_sym_PERCENTr),
	60:   uint16(anon_sym_PERCENTS),
	61:   uint16(anon_sym_PERCENTN),
	62:   uint16(anon_sym_PERCENT_PERCENT),
	63:   uint16(59),
	64:   uint16(14),
	65:   uint16(anon_sym_DOLLARkernel),
	66:   uint16(anon_sym_DOLLARnumber),
	67:   uint16(anon_sym_DOLLARdevpath),
	68:   uint16(anon_sym_DOLLARid),
	69:   uint16(anon_sym_DOLLARdriver),
	70:   uint16(anon_sym_DOLLARmajor),
	71:   uint16(anon_sym_DOLLARminor),
	72:   uint16(anon_sym_DOLLARparent),
	73:   uint16(anon_sym_DOLLARname),
	74:   uint16(anon_sym_DOLLARlinks),
	75:   uint16(anon_sym_DOLLARroot),
	76:   uint16(anon_sym_DOLLARsys),
	77:   uint16(anon_sym_DOLLARdevnode),
	78:   uint16(anon_sym_DOLLAR_DOLLAR),
	79:   uint16(15),
	80:   uint16(45),
	81:   uint16(1),
	82:   uint16(aux_sym_content_token1),
	83:   uint16(53),
	84:   uint16(1),
	85:   uint16(anon_sym_PERCENTs),
	86:   uint16(55),
	87:   uint16(1),
	88:   uint16(anon_sym_PERCENTE),
	89:   uint16(57),
	90:   uint16(1),
	91:   uint16(anon_sym_PERCENTc),
	92:   uint16(61),
	93:   uint16(1),
	94:   uint16(anon_sym_DOLLARattr),
	95:   uint16(63),
	96:   uint16(1),
	97:   uint16(anon_sym_DOLLARenv),
	98:   uint16(65),
	99:   uint16(1),
	100:  uint16(anon_sym_DOLLARresult),
	101:  uint16(67),
	102:  uint16(1),
	103:  uint16(sym_linebreak),
	104:  uint16(69),
	105:  uint16(1),
	106:  uint16(anon_sym_DQUOTE2),
	107:  uint16(4),
	108:  uint16(1),
	109:  uint16(aux_sym__sub_c_content),
	110:  uint16(47),
	111:  uint16(4),
	112:  uint16(anon_sym_STAR),
	113:  uint16(anon_sym_QMARK),
	114:  uint16(anon_sym_PIPE),
	115:  uint16(aux_sym_pattern_token1),
	116:  uint16(8),
	117:  uint16(4),
	118:  uint16(sym_pattern),
	119:  uint16(sym_c_escape),
	120:  uint16(sym_fmt_sub),
	121:  uint16(sym_var_sub),
	122:  uint16(49),
	123:  uint16(5),
	124:  uint16(aux_sym_c_escape_token1),
	125:  uint16(aux_sym_c_escape_token2),
	126:  uint16(aux_sym_c_escape_token3),
	127:  uint16(aux_sym_c_escape_token4),
	128:  uint16(aux_sym_c_escape_token5),
	129:  uint16(51),
	130:  uint16(11),
	131:  uint16(anon_sym_PERCENTk),
	132:  uint16(anon_sym_PERCENTn),
	133:  uint16(anon_sym_PERCENTp),
	134:  uint16(anon_sym_PERCENTb),
	135:  uint16(anon_sym_PERCENTM),
	136:  uint16(anon_sym_PERCENTm),
	137:  uint16(anon_sym_PERCENTP),
	138:  uint16(anon_sym_PERCENTr),
	139:  uint16(anon_sym_PERCENTS),
	140:  uint16(anon_sym_PERCENTN),
	141:  uint16(anon_sym_PERCENT_PERCENT),
	142:  uint16(59),
	143:  uint16(14),
	144:  uint16(anon_sym_DOLLARkernel),
	145:  uint16(anon_sym_DOLLARnumber),
	146:  uint16(anon_sym_DOLLARdevpath),
	147:  uint16(anon_sym_DOLLARid),
	148:  uint16(anon_sym_DOLLARdriver),
	149:  uint16(anon_sym_DOLLARmajor),
	150:  uint16(anon_sym_DOLLARminor),
	151:  uint16(anon_sym_DOLLARparent),
	152:  uint16(anon_sym_DOLLARname),
	153:  uint16(anon_sym_DOLLARlinks),
	154:  uint16(anon_sym_DOLLARroot),
	155:  uint16(anon_sym_DOLLARsys),
	156:  uint16(anon_sym_DOLLARdevnode),
	157:  uint16(anon_sym_DOLLAR_DOLLAR),
	158:  uint16(15),
	159:  uint16(67),
	160:  uint16(1),
	161:  uint16(sym_linebreak),
	162:  uint16(71),
	163:  uint16(1),
	164:  uint16(anon_sym_DQUOTE2),
	165:  uint16(73),
	166:  uint16(1),
	167:  uint16(aux_sym_content_token1),
	168:  uint16(85),
	169:  uint16(1),
	170:  uint16(anon_sym_PERCENTs),
	171:  uint16(88),
	172:  uint16(1),
	173:  uint16(anon_sym_PERCENTE),
	174:  uint16(91),
	175:  uint16(1),
	176:  uint16(anon_sym_PERCENTc),
	177:  uint16(97),
	178:  uint16(1),
	179:  uint16(anon_sym_DOLLARattr),
	180:  uint16(100),
	181:  uint16(1),
	182:  uint16(anon_sym_DOLLARenv),
	183:  uint16(103),
	184:  uint16(1),
	185:  uint16(anon_sym_DOLLARresult),
	186:  uint16(4),
	187:  uint16(1),
	188:  uint16(aux_sym__sub_c_content),
	189:  uint16(76),
	190:  uint16(4),
	191:  uint16(anon_sym_STAR),
	192:  uint16(anon_sym_QMARK),
	193:  uint16(anon_sym_PIPE),
	194:  uint16(aux_sym_pattern_token1),
	195:  uint16(8),
	196:  uint16(4),
	197:  uint16(sym_pattern),
	198:  uint16(sym_c_escape),
	199:  uint16(sym_fmt_sub),
	200:  uint16(sym_var_sub),
	201:  uint16(79),
	202:  uint16(5),
	203:  uint16(aux_sym_c_escape_token1),
	204:  uint16(aux_sym_c_escape_token2),
	205:  uint16(aux_sym_c_escape_token3),
	206:  uint16(aux_sym_c_escape_token4),
	207:  uint16(aux_sym_c_escape_token5),
	208:  uint16(82),
	209:  uint16(11),
	210:  uint16(anon_sym_PERCENTk),
	211:  uint16(anon_sym_PERCENTn),
	212:  uint16(anon_sym_PERCENTp),
	213:  uint16(anon_sym_PERCENTb),
	214:  uint16(anon_sym_PERCENTM),
	215:  uint16(anon_sym_PERCENTm),
	216:  uint16(anon_sym_PERCENTP),
	217:  uint16(anon_sym_PERCENTr),
	218:  uint16(anon_sym_PERCENTS),
	219:  uint16(anon_sym_PERCENTN),
	220:  uint16(anon_sym_PERCENT_PERCENT),
	221:  uint16(94),
	222:  uint16(14),
	223:  uint16(anon_sym_DOLLARkernel),
	224:  uint16(anon_sym_DOLLARnumber),
	225:  uint16(anon_sym_DOLLARdevpath),
	226:  uint16(anon_sym_DOLLARid),
	227:  uint16(anon_sym_DOLLARdriver),
	228:  uint16(anon_sym_DOLLARmajor),
	229:  uint16(anon_sym_DOLLARminor),
	230:  uint16(anon_sym_DOLLARparent),
	231:  uint16(anon_sym_DOLLARname),
	232:  uint16(anon_sym_DOLLARlinks),
	233:  uint16(anon_sym_DOLLARroot),
	234:  uint16(anon_sym_DOLLARsys),
	235:  uint16(anon_sym_DOLLARdevnode),
	236:  uint16(anon_sym_DOLLAR_DOLLAR),
	237:  uint16(4),
	238:  uint16(67),
	239:  uint16(1),
	240:  uint16(sym_linebreak),
	241:  uint16(106),
	242:  uint16(1),
	243:  uint16(anon_sym_LBRACE),
	244:  uint16(108),
	245:  uint16(1),
	246:  uint16(anon_sym_DQUOTE2),
	247:  uint16(110),
	248:  uint16(41),
	249:  uint16(aux_sym_content_token1),
	250:  uint16(anon_sym_STAR),
	251:  uint16(anon_sym_QMARK),
	252:  uint16(anon_sym_PIPE),
	253:  uint16(aux_sym_pattern_token1),
	254:  uint16(aux_sym_c_escape_token1),
	255:  uint16(aux_sym_c_escape_token2),
	256:  uint16(aux_sym_c_escape_token3),
	257:  uint16(aux_sym_c_escape_token4),
	258:  uint16(aux_sym_c_escape_token5),
	259:  uint16(anon_sym_PERCENTk),
	260:  uint16(anon_sym_PERCENTn),
	261:  uint16(anon_sym_PERCENTp),
	262:  uint16(anon_sym_PERCENTb),
	263:  uint16(anon_sym_PERCENTs),
	264:  uint16(anon_sym_PERCENTE),
	265:  uint16(anon_sym_PERCENTM),
	266:  uint16(anon_sym_PERCENTm),
	267:  uint16(anon_sym_PERCENTc),
	268:  uint16(anon_sym_PERCENTP),
	269:  uint16(anon_sym_PERCENTr),
	270:  uint16(anon_sym_PERCENTS),
	271:  uint16(anon_sym_PERCENTN),
	272:  uint16(anon_sym_PERCENT_PERCENT),
	273:  uint16(anon_sym_DOLLARkernel),
	274:  uint16(anon_sym_DOLLARnumber),
	275:  uint16(anon_sym_DOLLARdevpath),
	276:  uint16(anon_sym_DOLLARid),
	277:  uint16(anon_sym_DOLLARdriver),
	278:  uint16(anon_sym_DOLLARattr),
	279:  uint16(anon_sym_DOLLARenv),
	280:  uint16(anon_sym_DOLLARmajor),
	281:  uint16(anon_sym_DOLLARminor),
	282:  uint16(anon_sym_DOLLARresult),
	283:  uint16(anon_sym_DOLLARparent),
	284:  uint16(anon_sym_DOLLARname),
	285:  uint16(anon_sym_DOLLARlinks),
	286:  uint16(anon_sym_DOLLARroot),
	287:  uint16(anon_sym_DOLLARsys),
	288:  uint16(anon_sym_DOLLARdevnode),
	289:  uint16(anon_sym_DOLLAR_DOLLAR),
	290:  uint16(4),
	291:  uint16(67),
	292:  uint16(1),
	293:  uint16(sym_linebreak),
	294:  uint16(112),
	295:  uint16(1),
	296:  uint16(anon_sym_LBRACE),
	297:  uint16(114),
	298:  uint16(1),
	299:  uint16(anon_sym_DQUOTE2),
	300:  uint16(116),
	301:  uint16(41),
	302:  uint16(aux_sym_content_token1),
	303:  uint16(anon_sym_STAR),
	304:  uint16(anon_sym_QMARK),
	305:  uint16(anon_sym_PIPE),
	306:  uint16(aux_sym_pattern_token1),
	307:  uint16(aux_sym_c_escape_token1),
	308:  uint16(aux_sym_c_escape_token2),
	309:  uint16(aux_sym_c_escape_token3),
	310:  uint16(aux_sym_c_escape_token4),
	311:  uint16(aux_sym_c_escape_token5),
	312:  uint16(anon_sym_PERCENTk),
	313:  uint16(anon_sym_PERCENTn),
	314:  uint16(anon_sym_PERCENTp),
	315:  uint16(anon_sym_PERCENTb),
	316:  uint16(anon_sym_PERCENTs),
	317:  uint16(anon_sym_PERCENTE),
	318:  uint16(anon_sym_PERCENTM),
	319:  uint16(anon_sym_PERCENTm),
	320:  uint16(anon_sym_PERCENTc),
	321:  uint16(anon_sym_PERCENTP),
	322:  uint16(anon_sym_PERCENTr),
	323:  uint16(anon_sym_PERCENTS),
	324:  uint16(anon_sym_PERCENTN),
	325:  uint16(anon_sym_PERCENT_PERCENT),
	326:  uint16(anon_sym_DOLLARkernel),
	327:  uint16(anon_sym_DOLLARnumber),
	328:  uint16(anon_sym_DOLLARdevpath),
	329:  uint16(anon_sym_DOLLARid),
	330:  uint16(anon_sym_DOLLARdriver),
	331:  uint16(anon_sym_DOLLARattr),
	332:  uint16(anon_sym_DOLLARenv),
	333:  uint16(anon_sym_DOLLARmajor),
	334:  uint16(anon_sym_DOLLARminor),
	335:  uint16(anon_sym_DOLLARresult),
	336:  uint16(anon_sym_DOLLARparent),
	337:  uint16(anon_sym_DOLLARname),
	338:  uint16(anon_sym_DOLLARlinks),
	339:  uint16(anon_sym_DOLLARroot),
	340:  uint16(anon_sym_DOLLARsys),
	341:  uint16(anon_sym_DOLLARdevnode),
	342:  uint16(anon_sym_DOLLAR_DOLLAR),
	343:  uint16(14),
	344:  uint16(67),
	345:  uint16(1),
	346:  uint16(sym_linebreak),
	347:  uint16(118),
	348:  uint16(1),
	349:  uint16(anon_sym_DQUOTE2),
	350:  uint16(129),
	351:  uint16(1),
	352:  uint16(anon_sym_PERCENTs),
	353:  uint16(132),
	354:  uint16(1),
	355:  uint16(anon_sym_PERCENTE),
	356:  uint16(135),
	357:  uint16(1),
	358:  uint16(anon_sym_PERCENTc),
	359:  uint16(141),
	360:  uint16(1),
	361:  uint16(anon_sym_DOLLARattr),
	362:  uint16(144),
	363:  uint16(1),
	364:  uint16(anon_sym_DOLLARenv),
	365:  uint16(147),
	366:  uint16(1),
	367:  uint16(anon_sym_DOLLARresult),
	368:  uint16(7),
	369:  uint16(1),
	370:  uint16(aux_sym__sub_content),
	371:  uint16(120),
	372:  uint16(2),
	373:  uint16(aux_sym_content_token1),
	374:  uint16(anon_sym_BSLASH_DQUOTE),
	375:  uint16(21),
	376:  uint16(3),
	377:  uint16(sym_pattern),
	378:  uint16(sym_fmt_sub),
	379:  uint16(sym_var_sub),
	380:  uint16(123),
	381:  uint16(4),
	382:  uint16(anon_sym_STAR),
	383:  uint16(anon_sym_QMARK),
	384:  uint16(anon_sym_PIPE),
	385:  uint16(aux_sym_pattern_token1),
	386:  uint16(126),
	387:  uint16(11),
	388:  uint16(anon_sym_PERCENTk),
	389:  uint16(anon_sym_PERCENTn),
	390:  uint16(anon_sym_PERCENTp),
	391:  uint16(anon_sym_PERCENTb),
	392:  uint16(anon_sym_PERCENTM),
	393:  uint16(anon_sym_PERCENTm),
	394:  uint16(anon_sym_PERCENTP),
	395:  uint16(anon_sym_PERCENTr),
	396:  uint16(anon_sym_PERCENTS),
	397:  uint16(anon_sym_PERCENTN),
	398:  uint16(anon_sym_PERCENT_PERCENT),
	399:  uint16(138),
	400:  uint16(14),
	401:  uint16(anon_sym_DOLLARkernel),
	402:  uint16(anon_sym_DOLLARnumber),
	403:  uint16(anon_sym_DOLLARdevpath),
	404:  uint16(anon_sym_DOLLARid),
	405:  uint16(anon_sym_DOLLARdriver),
	406:  uint16(anon_sym_DOLLARmajor),
	407:  uint16(anon_sym_DOLLARminor),
	408:  uint16(anon_sym_DOLLARparent),
	409:  uint16(anon_sym_DOLLARname),
	410:  uint16(anon_sym_DOLLARlinks),
	411:  uint16(anon_sym_DOLLARroot),
	412:  uint16(anon_sym_DOLLARsys),
	413:  uint16(anon_sym_DOLLARdevnode),
	414:  uint16(anon_sym_DOLLAR_DOLLAR),
	415:  uint16(3),
	416:  uint16(67),
	417:  uint16(1),
	418:  uint16(sym_linebreak),
	419:  uint16(150),
	420:  uint16(1),
	421:  uint16(anon_sym_DQUOTE2),
	422:  uint16(152),
	423:  uint16(41),
	424:  uint16(aux_sym_content_token1),
	425:  uint16(anon_sym_STAR),
	426:  uint16(anon_sym_QMARK),
	427:  uint16(anon_sym_PIPE),
	428:  uint16(aux_sym_pattern_token1),
	429:  uint16(aux_sym_c_escape_token1),
	430:  uint16(aux_sym_c_escape_token2),
	431:  uint16(aux_sym_c_escape_token3),
	432:  uint16(aux_sym_c_escape_token4),
	433:  uint16(aux_sym_c_escape_token5),
	434:  uint16(anon_sym_PERCENTk),
	435:  uint16(anon_sym_PERCENTn),
	436:  uint16(anon_sym_PERCENTp),
	437:  uint16(anon_sym_PERCENTb),
	438:  uint16(anon_sym_PERCENTs),
	439:  uint16(anon_sym_PERCENTE),
	440:  uint16(anon_sym_PERCENTM),
	441:  uint16(anon_sym_PERCENTm),
	442:  uint16(anon_sym_PERCENTc),
	443:  uint16(anon_sym_PERCENTP),
	444:  uint16(anon_sym_PERCENTr),
	445:  uint16(anon_sym_PERCENTS),
	446:  uint16(anon_sym_PERCENTN),
	447:  uint16(anon_sym_PERCENT_PERCENT),
	448:  uint16(anon_sym_DOLLARkernel),
	449:  uint16(anon_sym_DOLLARnumber),
	450:  uint16(anon_sym_DOLLARdevpath),
	451:  uint16(anon_sym_DOLLARid),
	452:  uint16(anon_sym_DOLLARdriver),
	453:  uint16(anon_sym_DOLLARattr),
	454:  uint16(anon_sym_DOLLARenv),
	455:  uint16(anon_sym_DOLLARmajor),
	456:  uint16(anon_sym_DOLLARminor),
	457:  uint16(anon_sym_DOLLARresult),
	458:  uint16(anon_sym_DOLLARparent),
	459:  uint16(anon_sym_DOLLARname),
	460:  uint16(anon_sym_DOLLARlinks),
	461:  uint16(anon_sym_DOLLARroot),
	462:  uint16(anon_sym_DOLLARsys),
	463:  uint16(anon_sym_DOLLARdevnode),
	464:  uint16(anon_sym_DOLLAR_DOLLAR),
	465:  uint16(14),
	466:  uint16(67),
	467:  uint16(1),
	468:  uint16(sym_linebreak),
	469:  uint16(154),
	470:  uint16(1),
	471:  uint16(anon_sym_DQUOTE2),
	472:  uint16(162),
	473:  uint16(1),
	474:  uint16(anon_sym_PERCENTs),
	475:  uint16(164),
	476:  uint16(1),
	477:  uint16(anon_sym_PERCENTE),
	478:  uint16(166),
	479:  uint16(1),
	480:  uint16(anon_sym_PERCENTc),
	481:  uint16(170),
	482:  uint16(1),
	483:  uint16(anon_sym_DOLLARattr),
	484:  uint16(172),
	485:  uint16(1),
	486:  uint16(anon_sym_DOLLARenv),
	487:  uint16(174),
	488:  uint16(1),
	489:  uint16(anon_sym_DOLLARresult),
	490:  uint16(7),
	491:  uint16(1),
	492:  uint16(aux_sym__sub_content),
	493:  uint16(156),
	494:  uint16(2),
	495:  uint16(aux_sym_content_token1),
	496:  uint16(anon_sym_BSLASH_DQUOTE),
	497:  uint16(21),
	498:  uint16(3),
	499:  uint16(sym_pattern),
	500:  uint16(sym_fmt_sub),
	501:  uint16(sym_var_sub),
	502:  uint16(158),
	503:  uint16(4),
	504:  uint16(anon_sym_STAR),
	505:  uint16(anon_sym_QMARK),
	506:  uint16(anon_sym_PIPE),
	507:  uint16(aux_sym_pattern_token1),
	508:  uint16(160),
	509:  uint16(11),
	510:  uint16(anon_sym_PERCENTk),
	511:  uint16(anon_sym_PERCENTn),
	512:  uint16(anon_sym_PERCENTp),
	513:  uint16(anon_sym_PERCENTb),
	514:  uint16(anon_sym_PERCENTM),
	515:  uint16(anon_sym_PERCENTm),
	516:  uint16(anon_sym_PERCENTP),
	517:  uint16(anon_sym_PERCENTr),
	518:  uint16(anon_sym_PERCENTS),
	519:  uint16(anon_sym_PERCENTN),
	520:  uint16(anon_sym_PERCENT_PERCENT),
	521:  uint16(168),
	522:  uint16(14),
	523:  uint16(anon_sym_DOLLARkernel),
	524:  uint16(anon_sym_DOLLARnumber),
	525:  uint16(anon_sym_DOLLARdevpath),
	526:  uint16(anon_sym_DOLLARid),
	527:  uint16(anon_sym_DOLLARdriver),
	528:  uint16(anon_sym_DOLLARmajor),
	529:  uint16(anon_sym_DOLLARminor),
	530:  uint16(anon_sym_DOLLARparent),
	531:  uint16(anon_sym_DOLLARname),
	532:  uint16(anon_sym_DOLLARlinks),
	533:  uint16(anon_sym_DOLLARroot),
	534:  uint16(anon_sym_DOLLARsys),
	535:  uint16(anon_sym_DOLLARdevnode),
	536:  uint16(anon_sym_DOLLAR_DOLLAR),
	537:  uint16(14),
	538:  uint16(67),
	539:  uint16(1),
	540:  uint16(sym_linebreak),
	541:  uint16(162),
	542:  uint16(1),
	543:  uint16(anon_sym_PERCENTs),
	544:  uint16(164),
	545:  uint16(1),
	546:  uint16(anon_sym_PERCENTE),
	547:  uint16(166),
	548:  uint16(1),
	549:  uint16(anon_sym_PERCENTc),
	550:  uint16(170),
	551:  uint16(1),
	552:  uint16(anon_sym_DOLLARattr),
	553:  uint16(172),
	554:  uint16(1),
	555:  uint16(anon_sym_DOLLARenv),
	556:  uint16(174),
	557:  uint16(1),
	558:  uint16(anon_sym_DOLLARresult),
	559:  uint16(176),
	560:  uint16(1),
	561:  uint16(anon_sym_DQUOTE2),
	562:  uint16(9),
	563:  uint16(1),
	564:  uint16(aux_sym__sub_content),
	565:  uint16(156),
	566:  uint16(2),
	567:  uint16(aux_sym_content_token1),
	568:  uint16(anon_sym_BSLASH_DQUOTE),
	569:  uint16(21),
	570:  uint16(3),
	571:  uint16(sym_pattern),
	572:  uint16(sym_fmt_sub),
	573:  uint16(sym_var_sub),
	574:  uint16(158),
	575:  uint16(4),
	576:  uint16(anon_sym_STAR),
	577:  uint16(anon_sym_QMARK),
	578:  uint16(anon_sym_PIPE),
	579:  uint16(aux_sym_pattern_token1),
	580:  uint16(160),
	581:  uint16(11),
	582:  uint16(anon_sym_PERCENTk),
	583:  uint16(anon_sym_PERCENTn),
	584:  uint16(anon_sym_PERCENTp),
	585:  uint16(anon_sym_PERCENTb),
	586:  uint16(anon_sym_PERCENTM),
	587:  uint16(anon_sym_PERCENTm),
	588:  uint16(anon_sym_PERCENTP),
	589:  uint16(anon_sym_PERCENTr),
	590:  uint16(anon_sym_PERCENTS),
	591:  uint16(anon_sym_PERCENTN),
	592:  uint16(anon_sym_PERCENT_PERCENT),
	593:  uint16(168),
	594:  uint16(14),
	595:  uint16(anon_sym_DOLLARkernel),
	596:  uint16(anon_sym_DOLLARnumber),
	597:  uint16(anon_sym_DOLLARdevpath),
	598:  uint16(anon_sym_DOLLARid),
	599:  uint16(anon_sym_DOLLARdriver),
	600:  uint16(anon_sym_DOLLARmajor),
	601:  uint16(anon_sym_DOLLARminor),
	602:  uint16(anon_sym_DOLLARparent),
	603:  uint16(anon_sym_DOLLARname),
	604:  uint16(anon_sym_DOLLARlinks),
	605:  uint16(anon_sym_DOLLARroot),
	606:  uint16(anon_sym_DOLLARsys),
	607:  uint16(anon_sym_DOLLARdevnode),
	608:  uint16(anon_sym_DOLLAR_DOLLAR),
	609:  uint16(3),
	610:  uint16(67),
	611:  uint16(1),
	612:  uint16(sym_linebreak),
	613:  uint16(178),
	614:  uint16(1),
	615:  uint16(anon_sym_DQUOTE2),
	616:  uint16(180),
	617:  uint16(41),
	618:  uint16(aux_sym_content_token1),
	619:  uint16(anon_sym_STAR),
	620:  uint16(anon_sym_QMARK),
	621:  uint16(anon_sym_PIPE),
	622:  uint16(aux_sym_pattern_token1),
	623:  uint16(aux_sym_c_escape_token1),
	624:  uint16(aux_sym_c_escape_token2),
	625:  uint16(aux_sym_c_escape_token3),
	626:  uint16(aux_sym_c_escape_token4),
	627:  uint16(aux_sym_c_escape_token5),
	628:  uint16(anon_sym_PERCENTk),
	629:  uint16(anon_sym_PERCENTn),
	630:  uint16(anon_sym_PERCENTp),
	631:  uint16(anon_sym_PERCENTb),
	632:  uint16(anon_sym_PERCENTs),
	633:  uint16(anon_sym_PERCENTE),
	634:  uint16(anon_sym_PERCENTM),
	635:  uint16(anon_sym_PERCENTm),
	636:  uint16(anon_sym_PERCENTc),
	637:  uint16(anon_sym_PERCENTP),
	638:  uint16(anon_sym_PERCENTr),
	639:  uint16(anon_sym_PERCENTS),
	640:  uint16(anon_sym_PERCENTN),
	641:  uint16(anon_sym_PERCENT_PERCENT),
	642:  uint16(anon_sym_DOLLARkernel),
	643:  uint16(anon_sym_DOLLARnumber),
	644:  uint16(anon_sym_DOLLARdevpath),
	645:  uint16(anon_sym_DOLLARid),
	646:  uint16(anon_sym_DOLLARdriver),
	647:  uint16(anon_sym_DOLLARattr),
	648:  uint16(anon_sym_DOLLARenv),
	649:  uint16(anon_sym_DOLLARmajor),
	650:  uint16(anon_sym_DOLLARminor),
	651:  uint16(anon_sym_DOLLARresult),
	652:  uint16(anon_sym_DOLLARparent),
	653:  uint16(anon_sym_DOLLARname),
	654:  uint16(anon_sym_DOLLARlinks),
	655:  uint16(anon_sym_DOLLARroot),
	656:  uint16(anon_sym_DOLLARsys),
	657:  uint16(anon_sym_DOLLARdevnode),
	658:  uint16(anon_sym_DOLLAR_DOLLAR),
	659:  uint16(3),
	660:  uint16(67),
	661:  uint16(1),
	662:  uint16(sym_linebreak),
	663:  uint16(114),
	664:  uint16(1),
	665:  uint16(anon_sym_DQUOTE2),
	666:  uint16(116),
	667:  uint16(41),
	668:  uint16(aux_sym_content_token1),
	669:  uint16(anon_sym_STAR),
	670:  uint16(anon_sym_QMARK),
	671:  uint16(anon_sym_PIPE),
	672:  uint16(aux_sym_pattern_token1),
	673:  uint16(aux_sym_c_escape_token1),
	674:  uint16(aux_sym_c_escape_token2),
	675:  uint16(aux_sym_c_escape_token3),
	676:  uint16(aux_sym_c_escape_token4),
	677:  uint16(aux_sym_c_escape_token5),
	678:  uint16(anon_sym_PERCENTk),
	679:  uint16(anon_sym_PERCENTn),
	680:  uint16(anon_sym_PERCENTp),
	681:  uint16(anon_sym_PERCENTb),
	682:  uint16(anon_sym_PERCENTs),
	683:  uint16(anon_sym_PERCENTE),
	684:  uint16(anon_sym_PERCENTM),
	685:  uint16(anon_sym_PERCENTm),
	686:  uint16(anon_sym_PERCENTc),
	687:  uint16(anon_sym_PERCENTP),
	688:  uint16(anon_sym_PERCENTr),
	689:  uint16(anon_sym_PERCENTS),
	690:  uint16(anon_sym_PERCENTN),
	691:  uint16(anon_sym_PERCENT_PERCENT),
	692:  uint16(anon_sym_DOLLARkernel),
	693:  uint16(anon_sym_DOLLARnumber),
	694:  uint16(anon_sym_DOLLARdevpath),
	695:  uint16(anon_sym_DOLLARid),
	696:  uint16(anon_sym_DOLLARdriver),
	697:  uint16(anon_sym_DOLLARattr),
	698:  uint16(anon_sym_DOLLARenv),
	699:  uint16(anon_sym_DOLLARmajor),
	700:  uint16(anon_sym_DOLLARminor),
	701:  uint16(anon_sym_DOLLARresult),
	702:  uint16(anon_sym_DOLLARparent),
	703:  uint16(anon_sym_DOLLARname),
	704:  uint16(anon_sym_DOLLARlinks),
	705:  uint16(anon_sym_DOLLARroot),
	706:  uint16(anon_sym_DOLLARsys),
	707:  uint16(anon_sym_DOLLARdevnode),
	708:  uint16(anon_sym_DOLLAR_DOLLAR),
	709:  uint16(3),
	710:  uint16(67),
	711:  uint16(1),
	712:  uint16(sym_linebreak),
	713:  uint16(182),
	714:  uint16(1),
	715:  uint16(anon_sym_DQUOTE2),
	716:  uint16(184),
	717:  uint16(41),
	718:  uint16(aux_sym_content_token1),
	719:  uint16(anon_sym_STAR),
	720:  uint16(anon_sym_QMARK),
	721:  uint16(anon_sym_PIPE),
	722:  uint16(aux_sym_pattern_token1),
	723:  uint16(aux_sym_c_escape_token1),
	724:  uint16(aux_sym_c_escape_token2),
	725:  uint16(aux_sym_c_escape_token3),
	726:  uint16(aux_sym_c_escape_token4),
	727:  uint16(aux_sym_c_escape_token5),
	728:  uint16(anon_sym_PERCENTk),
	729:  uint16(anon_sym_PERCENTn),
	730:  uint16(anon_sym_PERCENTp),
	731:  uint16(anon_sym_PERCENTb),
	732:  uint16(anon_sym_PERCENTs),
	733:  uint16(anon_sym_PERCENTE),
	734:  uint16(anon_sym_PERCENTM),
	735:  uint16(anon_sym_PERCENTm),
	736:  uint16(anon_sym_PERCENTc),
	737:  uint16(anon_sym_PERCENTP),
	738:  uint16(anon_sym_PERCENTr),
	739:  uint16(anon_sym_PERCENTS),
	740:  uint16(anon_sym_PERCENTN),
	741:  uint16(anon_sym_PERCENT_PERCENT),
	742:  uint16(anon_sym_DOLLARkernel),
	743:  uint16(anon_sym_DOLLARnumber),
	744:  uint16(anon_sym_DOLLARdevpath),
	745:  uint16(anon_sym_DOLLARid),
	746:  uint16(anon_sym_DOLLARdriver),
	747:  uint16(anon_sym_DOLLARattr),
	748:  uint16(anon_sym_DOLLARenv),
	749:  uint16(anon_sym_DOLLARmajor),
	750:  uint16(anon_sym_DOLLARminor),
	751:  uint16(anon_sym_DOLLARresult),
	752:  uint16(anon_sym_DOLLARparent),
	753:  uint16(anon_sym_DOLLARname),
	754:  uint16(anon_sym_DOLLARlinks),
	755:  uint16(anon_sym_DOLLARroot),
	756:  uint16(anon_sym_DOLLARsys),
	757:  uint16(anon_sym_DOLLARdevnode),
	758:  uint16(anon_sym_DOLLAR_DOLLAR),
	759:  uint16(3),
	760:  uint16(67),
	761:  uint16(1),
	762:  uint16(sym_linebreak),
	763:  uint16(186),
	764:  uint16(1),
	765:  uint16(anon_sym_DQUOTE2),
	766:  uint16(188),
	767:  uint16(41),
	768:  uint16(aux_sym_content_token1),
	769:  uint16(anon_sym_STAR),
	770:  uint16(anon_sym_QMARK),
	771:  uint16(anon_sym_PIPE),
	772:  uint16(aux_sym_pattern_token1),
	773:  uint16(aux_sym_c_escape_token1),
	774:  uint16(aux_sym_c_escape_token2),
	775:  uint16(aux_sym_c_escape_token3),
	776:  uint16(aux_sym_c_escape_token4),
	777:  uint16(aux_sym_c_escape_token5),
	778:  uint16(anon_sym_PERCENTk),
	779:  uint16(anon_sym_PERCENTn),
	780:  uint16(anon_sym_PERCENTp),
	781:  uint16(anon_sym_PERCENTb),
	782:  uint16(anon_sym_PERCENTs),
	783:  uint16(anon_sym_PERCENTE),
	784:  uint16(anon_sym_PERCENTM),
	785:  uint16(anon_sym_PERCENTm),
	786:  uint16(anon_sym_PERCENTc),
	787:  uint16(anon_sym_PERCENTP),
	788:  uint16(anon_sym_PERCENTr),
	789:  uint16(anon_sym_PERCENTS),
	790:  uint16(anon_sym_PERCENTN),
	791:  uint16(anon_sym_PERCENT_PERCENT),
	792:  uint16(anon_sym_DOLLARkernel),
	793:  uint16(anon_sym_DOLLARnumber),
	794:  uint16(anon_sym_DOLLARdevpath),
	795:  uint16(anon_sym_DOLLARid),
	796:  uint16(anon_sym_DOLLARdriver),
	797:  uint16(anon_sym_DOLLARattr),
	798:  uint16(anon_sym_DOLLARenv),
	799:  uint16(anon_sym_DOLLARmajor),
	800:  uint16(anon_sym_DOLLARminor),
	801:  uint16(anon_sym_DOLLARresult),
	802:  uint16(anon_sym_DOLLARparent),
	803:  uint16(anon_sym_DOLLARname),
	804:  uint16(anon_sym_DOLLARlinks),
	805:  uint16(anon_sym_DOLLARroot),
	806:  uint16(anon_sym_DOLLARsys),
	807:  uint16(anon_sym_DOLLARdevnode),
	808:  uint16(anon_sym_DOLLAR_DOLLAR),
	809:  uint16(3),
	810:  uint16(67),
	811:  uint16(1),
	812:  uint16(sym_linebreak),
	813:  uint16(108),
	814:  uint16(1),
	815:  uint16(anon_sym_DQUOTE2),
	816:  uint16(110),
	817:  uint16(41),
	818:  uint16(aux_sym_content_token1),
	819:  uint16(anon_sym_STAR),
	820:  uint16(anon_sym_QMARK),
	821:  uint16(anon_sym_PIPE),
	822:  uint16(aux_sym_pattern_token1),
	823:  uint16(aux_sym_c_escape_token1),
	824:  uint16(aux_sym_c_escape_token2),
	825:  uint16(aux_sym_c_escape_token3),
	826:  uint16(aux_sym_c_escape_token4),
	827:  uint16(aux_sym_c_escape_token5),
	828:  uint16(anon_sym_PERCENTk),
	829:  uint16(anon_sym_PERCENTn),
	830:  uint16(anon_sym_PERCENTp),
	831:  uint16(anon_sym_PERCENTb),
	832:  uint16(anon_sym_PERCENTs),
	833:  uint16(anon_sym_PERCENTE),
	834:  uint16(anon_sym_PERCENTM),
	835:  uint16(anon_sym_PERCENTm),
	836:  uint16(anon_sym_PERCENTc),
	837:  uint16(anon_sym_PERCENTP),
	838:  uint16(anon_sym_PERCENTr),
	839:  uint16(anon_sym_PERCENTS),
	840:  uint16(anon_sym_PERCENTN),
	841:  uint16(anon_sym_PERCENT_PERCENT),
	842:  uint16(anon_sym_DOLLARkernel),
	843:  uint16(anon_sym_DOLLARnumber),
	844:  uint16(anon_sym_DOLLARdevpath),
	845:  uint16(anon_sym_DOLLARid),
	846:  uint16(anon_sym_DOLLARdriver),
	847:  uint16(anon_sym_DOLLARattr),
	848:  uint16(anon_sym_DOLLARenv),
	849:  uint16(anon_sym_DOLLARmajor),
	850:  uint16(anon_sym_DOLLARminor),
	851:  uint16(anon_sym_DOLLARresult),
	852:  uint16(anon_sym_DOLLARparent),
	853:  uint16(anon_sym_DOLLARname),
	854:  uint16(anon_sym_DOLLARlinks),
	855:  uint16(anon_sym_DOLLARroot),
	856:  uint16(anon_sym_DOLLARsys),
	857:  uint16(anon_sym_DOLLARdevnode),
	858:  uint16(anon_sym_DOLLAR_DOLLAR),
	859:  uint16(3),
	860:  uint16(67),
	861:  uint16(1),
	862:  uint16(sym_linebreak),
	863:  uint16(190),
	864:  uint16(1),
	865:  uint16(anon_sym_DQUOTE2),
	866:  uint16(192),
	867:  uint16(41),
	868:  uint16(aux_sym_content_token1),
	869:  uint16(anon_sym_STAR),
	870:  uint16(anon_sym_QMARK),
	871:  uint16(anon_sym_PIPE),
	872:  uint16(aux_sym_pattern_token1),
	873:  uint16(aux_sym_c_escape_token1),
	874:  uint16(aux_sym_c_escape_token2),
	875:  uint16(aux_sym_c_escape_token3),
	876:  uint16(aux_sym_c_escape_token4),
	877:  uint16(aux_sym_c_escape_token5),
	878:  uint16(anon_sym_PERCENTk),
	879:  uint16(anon_sym_PERCENTn),
	880:  uint16(anon_sym_PERCENTp),
	881:  uint16(anon_sym_PERCENTb),
	882:  uint16(anon_sym_PERCENTs),
	883:  uint16(anon_sym_PERCENTE),
	884:  uint16(anon_sym_PERCENTM),
	885:  uint16(anon_sym_PERCENTm),
	886:  uint16(anon_sym_PERCENTc),
	887:  uint16(anon_sym_PERCENTP),
	888:  uint16(anon_sym_PERCENTr),
	889:  uint16(anon_sym_PERCENTS),
	890:  uint16(anon_sym_PERCENTN),
	891:  uint16(anon_sym_PERCENT_PERCENT),
	892:  uint16(anon_sym_DOLLARkernel),
	893:  uint16(anon_sym_DOLLARnumber),
	894:  uint16(anon_sym_DOLLARdevpath),
	895:  uint16(anon_sym_DOLLARid),
	896:  uint16(anon_sym_DOLLARdriver),
	897:  uint16(anon_sym_DOLLARattr),
	898:  uint16(anon_sym_DOLLARenv),
	899:  uint16(anon_sym_DOLLARmajor),
	900:  uint16(anon_sym_DOLLARminor),
	901:  uint16(anon_sym_DOLLARresult),
	902:  uint16(anon_sym_DOLLARparent),
	903:  uint16(anon_sym_DOLLARname),
	904:  uint16(anon_sym_DOLLARlinks),
	905:  uint16(anon_sym_DOLLARroot),
	906:  uint16(anon_sym_DOLLARsys),
	907:  uint16(anon_sym_DOLLARdevnode),
	908:  uint16(anon_sym_DOLLAR_DOLLAR),
	909:  uint16(3),
	910:  uint16(67),
	911:  uint16(1),
	912:  uint16(sym_linebreak),
	913:  uint16(194),
	914:  uint16(1),
	915:  uint16(anon_sym_DQUOTE2),
	916:  uint16(196),
	917:  uint16(41),
	918:  uint16(aux_sym_content_token1),
	919:  uint16(anon_sym_STAR),
	920:  uint16(anon_sym_QMARK),
	921:  uint16(anon_sym_PIPE),
	922:  uint16(aux_sym_pattern_token1),
	923:  uint16(aux_sym_c_escape_token1),
	924:  uint16(aux_sym_c_escape_token2),
	925:  uint16(aux_sym_c_escape_token3),
	926:  uint16(aux_sym_c_escape_token4),
	927:  uint16(aux_sym_c_escape_token5),
	928:  uint16(anon_sym_PERCENTk),
	929:  uint16(anon_sym_PERCENTn),
	930:  uint16(anon_sym_PERCENTp),
	931:  uint16(anon_sym_PERCENTb),
	932:  uint16(anon_sym_PERCENTs),
	933:  uint16(anon_sym_PERCENTE),
	934:  uint16(anon_sym_PERCENTM),
	935:  uint16(anon_sym_PERCENTm),
	936:  uint16(anon_sym_PERCENTc),
	937:  uint16(anon_sym_PERCENTP),
	938:  uint16(anon_sym_PERCENTr),
	939:  uint16(anon_sym_PERCENTS),
	940:  uint16(anon_sym_PERCENTN),
	941:  uint16(anon_sym_PERCENT_PERCENT),
	942:  uint16(anon_sym_DOLLARkernel),
	943:  uint16(anon_sym_DOLLARnumber),
	944:  uint16(anon_sym_DOLLARdevpath),
	945:  uint16(anon_sym_DOLLARid),
	946:  uint16(anon_sym_DOLLARdriver),
	947:  uint16(anon_sym_DOLLARattr),
	948:  uint16(anon_sym_DOLLARenv),
	949:  uint16(anon_sym_DOLLARmajor),
	950:  uint16(anon_sym_DOLLARminor),
	951:  uint16(anon_sym_DOLLARresult),
	952:  uint16(anon_sym_DOLLARparent),
	953:  uint16(anon_sym_DOLLARname),
	954:  uint16(anon_sym_DOLLARlinks),
	955:  uint16(anon_sym_DOLLARroot),
	956:  uint16(anon_sym_DOLLARsys),
	957:  uint16(anon_sym_DOLLARdevnode),
	958:  uint16(anon_sym_DOLLAR_DOLLAR),
	959:  uint16(3),
	960:  uint16(67),
	961:  uint16(1),
	962:  uint16(sym_linebreak),
	963:  uint16(198),
	964:  uint16(1),
	965:  uint16(anon_sym_DQUOTE2),
	966:  uint16(200),
	967:  uint16(41),
	968:  uint16(aux_sym_content_token1),
	969:  uint16(anon_sym_STAR),
	970:  uint16(anon_sym_QMARK),
	971:  uint16(anon_sym_PIPE),
	972:  uint16(aux_sym_pattern_token1),
	973:  uint16(aux_sym_c_escape_token1),
	974:  uint16(aux_sym_c_escape_token2),
	975:  uint16(aux_sym_c_escape_token3),
	976:  uint16(aux_sym_c_escape_token4),
	977:  uint16(aux_sym_c_escape_token5),
	978:  uint16(anon_sym_PERCENTk),
	979:  uint16(anon_sym_PERCENTn),
	980:  uint16(anon_sym_PERCENTp),
	981:  uint16(anon_sym_PERCENTb),
	982:  uint16(anon_sym_PERCENTs),
	983:  uint16(anon_sym_PERCENTE),
	984:  uint16(anon_sym_PERCENTM),
	985:  uint16(anon_sym_PERCENTm),
	986:  uint16(anon_sym_PERCENTc),
	987:  uint16(anon_sym_PERCENTP),
	988:  uint16(anon_sym_PERCENTr),
	989:  uint16(anon_sym_PERCENTS),
	990:  uint16(anon_sym_PERCENTN),
	991:  uint16(anon_sym_PERCENT_PERCENT),
	992:  uint16(anon_sym_DOLLARkernel),
	993:  uint16(anon_sym_DOLLARnumber),
	994:  uint16(anon_sym_DOLLARdevpath),
	995:  uint16(anon_sym_DOLLARid),
	996:  uint16(anon_sym_DOLLARdriver),
	997:  uint16(anon_sym_DOLLARattr),
	998:  uint16(anon_sym_DOLLARenv),
	999:  uint16(anon_sym_DOLLARmajor),
	1000: uint16(anon_sym_DOLLARminor),
	1001: uint16(anon_sym_DOLLARresult),
	1002: uint16(anon_sym_DOLLARparent),
	1003: uint16(anon_sym_DOLLARname),
	1004: uint16(anon_sym_DOLLARlinks),
	1005: uint16(anon_sym_DOLLARroot),
	1006: uint16(anon_sym_DOLLARsys),
	1007: uint16(anon_sym_DOLLARdevnode),
	1008: uint16(anon_sym_DOLLAR_DOLLAR),
	1009: uint16(4),
	1010: uint16(67),
	1011: uint16(1),
	1012: uint16(sym_linebreak),
	1013: uint16(114),
	1014: uint16(1),
	1015: uint16(anon_sym_DQUOTE2),
	1016: uint16(202),
	1017: uint16(1),
	1018: uint16(anon_sym_LBRACE),
	1019: uint16(116),
	1020: uint16(37),
	1021: uint16(aux_sym_content_token1),
	1022: uint16(anon_sym_BSLASH_DQUOTE),
	1023: uint16(anon_sym_STAR),
	1024: uint16(anon_sym_QMARK),
	1025: uint16(anon_sym_PIPE),
	1026: uint16(aux_sym_pattern_token1),
	1027: uint16(anon_sym_PERCENTk),
	1028: uint16(anon_sym_PERCENTn),
	1029: uint16(anon_sym_PERCENTp),
	1030: uint16(anon_sym_PERCENTb),
	1031: uint16(anon_sym_PERCENTs),
	1032: uint16(anon_sym_PERCENTE),
	1033: uint16(anon_sym_PERCENTM),
	1034: uint16(anon_sym_PERCENTm),
	1035: uint16(anon_sym_PERCENTc),
	1036: uint16(anon_sym_PERCENTP),
	1037: uint16(anon_sym_PERCENTr),
	1038: uint16(anon_sym_PERCENTS),
	1039: uint16(anon_sym_PERCENTN),
	1040: uint16(anon_sym_PERCENT_PERCENT),
	1041: uint16(anon_sym_DOLLARkernel),
	1042: uint16(anon_sym_DOLLARnumber),
	1043: uint16(anon_sym_DOLLARdevpath),
	1044: uint16(anon_sym_DOLLARid),
	1045: uint16(anon_sym_DOLLARdriver),
	1046: uint16(anon_sym_DOLLARattr),
	1047: uint16(anon_sym_DOLLARenv),
	1048: uint16(anon_sym_DOLLARmajor),
	1049: uint16(anon_sym_DOLLARminor),
	1050: uint16(anon_sym_DOLLARresult),
	1051: uint16(anon_sym_DOLLARparent),
	1052: uint16(anon_sym_DOLLARname),
	1053: uint16(anon_sym_DOLLARlinks),
	1054: uint16(anon_sym_DOLLARroot),
	1055: uint16(anon_sym_DOLLARsys),
	1056: uint16(anon_sym_DOLLARdevnode),
	1057: uint16(anon_sym_DOLLAR_DOLLAR),
	1058: uint16(4),
	1059: uint16(67),
	1060: uint16(1),
	1061: uint16(sym_linebreak),
	1062: uint16(108),
	1063: uint16(1),
	1064: uint16(anon_sym_DQUOTE2),
	1065: uint16(204),
	1066: uint16(1),
	1067: uint16(anon_sym_LBRACE),
	1068: uint16(110),
	1069: uint16(37),
	1070: uint16(aux_sym_content_token1),
	1071: uint16(anon_sym_BSLASH_DQUOTE),
	1072: uint16(anon_sym_STAR),
	1073: uint16(anon_sym_QMARK),
	1074: uint16(anon_sym_PIPE),
	1075: uint16(aux_sym_pattern_token1),
	1076: uint16(anon_sym_PERCENTk),
	1077: uint16(anon_sym_PERCENTn),
	1078: uint16(anon_sym_PERCENTp),
	1079: uint16(anon_sym_PERCENTb),
	1080: uint16(anon_sym_PERCENTs),
	1081: uint16(anon_sym_PERCENTE),
	1082: uint16(anon_sym_PERCENTM),
	1083: uint16(anon_sym_PERCENTm),
	1084: uint16(anon_sym_PERCENTc),
	1085: uint16(anon_sym_PERCENTP),
	1086: uint16(anon_sym_PERCENTr),
	1087: uint16(anon_sym_PERCENTS),
	1088: uint16(anon_sym_PERCENTN),
	1089: uint16(anon_sym_PERCENT_PERCENT),
	1090: uint16(anon_sym_DOLLARkernel),
	1091: uint16(anon_sym_DOLLARnumber),
	1092: uint16(anon_sym_DOLLARdevpath),
	1093: uint16(anon_sym_DOLLARid),
	1094: uint16(anon_sym_DOLLARdriver),
	1095: uint16(anon_sym_DOLLARattr),
	1096: uint16(anon_sym_DOLLARenv),
	1097: uint16(anon_sym_DOLLARmajor),
	1098: uint16(anon_sym_DOLLARminor),
	1099: uint16(anon_sym_DOLLARresult),
	1100: uint16(anon_sym_DOLLARparent),
	1101: uint16(anon_sym_DOLLARname),
	1102: uint16(anon_sym_DOLLARlinks),
	1103: uint16(anon_sym_DOLLARroot),
	1104: uint16(anon_sym_DOLLARsys),
	1105: uint16(anon_sym_DOLLARdevnode),
	1106: uint16(anon_sym_DOLLAR_DOLLAR),
	1107: uint16(3),
	1108: uint16(67),
	1109: uint16(1),
	1110: uint16(sym_linebreak),
	1111: uint16(206),
	1112: uint16(1),
	1113: uint16(anon_sym_DQUOTE2),
	1114: uint16(208),
	1115: uint16(37),
	1116: uint16(aux_sym_content_token1),
	1117: uint16(anon_sym_BSLASH_DQUOTE),
	1118: uint16(anon_sym_STAR),
	1119: uint16(anon_sym_QMARK),
	1120: uint16(anon_sym_PIPE),
	1121: uint16(aux_sym_pattern_token1),
	1122: uint16(anon_sym_PERCENTk),
	1123: uint16(anon_sym_PERCENTn),
	1124: uint16(anon_sym_PERCENTp),
	1125: uint16(anon_sym_PERCENTb),
	1126: uint16(anon_sym_PERCENTs),
	1127: uint16(anon_sym_PERCENTE),
	1128: uint16(anon_sym_PERCENTM),
	1129: uint16(anon_sym_PERCENTm),
	1130: uint16(anon_sym_PERCENTc),
	1131: uint16(anon_sym_PERCENTP),
	1132: uint16(anon_sym_PERCENTr),
	1133: uint16(anon_sym_PERCENTS),
	1134: uint16(anon_sym_PERCENTN),
	1135: uint16(anon_sym_PERCENT_PERCENT),
	1136: uint16(anon_sym_DOLLARkernel),
	1137: uint16(anon_sym_DOLLARnumber),
	1138: uint16(anon_sym_DOLLARdevpath),
	1139: uint16(anon_sym_DOLLARid),
	1140: uint16(anon_sym_DOLLARdriver),
	1141: uint16(anon_sym_DOLLARattr),
	1142: uint16(anon_sym_DOLLARenv),
	1143: uint16(anon_sym_DOLLARmajor),
	1144: uint16(anon_sym_DOLLARminor),
	1145: uint16(anon_sym_DOLLARresult),
	1146: uint16(anon_sym_DOLLARparent),
	1147: uint16(anon_sym_DOLLARname),
	1148: uint16(anon_sym_DOLLARlinks),
	1149: uint16(anon_sym_DOLLARroot),
	1150: uint16(anon_sym_DOLLARsys),
	1151: uint16(anon_sym_DOLLARdevnode),
	1152: uint16(anon_sym_DOLLAR_DOLLAR),
	1153: uint16(3),
	1154: uint16(67),
	1155: uint16(1),
	1156: uint16(sym_linebreak),
	1157: uint16(114),
	1158: uint16(1),
	1159: uint16(anon_sym_DQUOTE2),
	1160: uint16(116),
	1161: uint16(37),
	1162: uint16(aux_sym_content_token1),
	1163: uint16(anon_sym_BSLASH_DQUOTE),
	1164: uint16(anon_sym_STAR),
	1165: uint16(anon_sym_QMARK),
	1166: uint16(anon_sym_PIPE),
	1167: uint16(aux_sym_pattern_token1),
	1168: uint16(anon_sym_PERCENTk),
	1169: uint16(anon_sym_PERCENTn),
	1170: uint16(anon_sym_PERCENTp),
	1171: uint16(anon_sym_PERCENTb),
	1172: uint16(anon_sym_PERCENTs),
	1173: uint16(anon_sym_PERCENTE),
	1174: uint16(anon_sym_PERCENTM),
	1175: uint16(anon_sym_PERCENTm),
	1176: uint16(anon_sym_PERCENTc),
	1177: uint16(anon_sym_PERCENTP),
	1178: uint16(anon_sym_PERCENTr),
	1179: uint16(anon_sym_PERCENTS),
	1180: uint16(anon_sym_PERCENTN),
	1181: uint16(anon_sym_PERCENT_PERCENT),
	1182: uint16(anon_sym_DOLLARkernel),
	1183: uint16(anon_sym_DOLLARnumber),
	1184: uint16(anon_sym_DOLLARdevpath),
	1185: uint16(anon_sym_DOLLARid),
	1186: uint16(anon_sym_DOLLARdriver),
	1187: uint16(anon_sym_DOLLARattr),
	1188: uint16(anon_sym_DOLLARenv),
	1189: uint16(anon_sym_DOLLARmajor),
	1190: uint16(anon_sym_DOLLARminor),
	1191: uint16(anon_sym_DOLLARresult),
	1192: uint16(anon_sym_DOLLARparent),
	1193: uint16(anon_sym_DOLLARname),
	1194: uint16(anon_sym_DOLLARlinks),
	1195: uint16(anon_sym_DOLLARroot),
	1196: uint16(anon_sym_DOLLARsys),
	1197: uint16(anon_sym_DOLLARdevnode),
	1198: uint16(anon_sym_DOLLAR_DOLLAR),
	1199: uint16(3),
	1200: uint16(67),
	1201: uint16(1),
	1202: uint16(sym_linebreak),
	1203: uint16(186),
	1204: uint16(1),
	1205: uint16(anon_sym_DQUOTE2),
	1206: uint16(188),
	1207: uint16(37),
	1208: uint16(aux_sym_content_token1),
	1209: uint16(anon_sym_BSLASH_DQUOTE),
	1210: uint16(anon_sym_STAR),
	1211: uint16(anon_sym_QMARK),
	1212: uint16(anon_sym_PIPE),
	1213: uint16(aux_sym_pattern_token1),
	1214: uint16(anon_sym_PERCENTk),
	1215: uint16(anon_sym_PERCENTn),
	1216: uint16(anon_sym_PERCENTp),
	1217: uint16(anon_sym_PERCENTb),
	1218: uint16(anon_sym_PERCENTs),
	1219: uint16(anon_sym_PERCENTE),
	1220: uint16(anon_sym_PERCENTM),
	1221: uint16(anon_sym_PERCENTm),
	1222: uint16(anon_sym_PERCENTc),
	1223: uint16(anon_sym_PERCENTP),
	1224: uint16(anon_sym_PERCENTr),
	1225: uint16(anon_sym_PERCENTS),
	1226: uint16(anon_sym_PERCENTN),
	1227: uint16(anon_sym_PERCENT_PERCENT),
	1228: uint16(anon_sym_DOLLARkernel),
	1229: uint16(anon_sym_DOLLARnumber),
	1230: uint16(anon_sym_DOLLARdevpath),
	1231: uint16(anon_sym_DOLLARid),
	1232: uint16(anon_sym_DOLLARdriver),
	1233: uint16(anon_sym_DOLLARattr),
	1234: uint16(anon_sym_DOLLARenv),
	1235: uint16(anon_sym_DOLLARmajor),
	1236: uint16(anon_sym_DOLLARminor),
	1237: uint16(anon_sym_DOLLARresult),
	1238: uint16(anon_sym_DOLLARparent),
	1239: uint16(anon_sym_DOLLARname),
	1240: uint16(anon_sym_DOLLARlinks),
	1241: uint16(anon_sym_DOLLARroot),
	1242: uint16(anon_sym_DOLLARsys),
	1243: uint16(anon_sym_DOLLARdevnode),
	1244: uint16(anon_sym_DOLLAR_DOLLAR),
	1245: uint16(3),
	1246: uint16(67),
	1247: uint16(1),
	1248: uint16(sym_linebreak),
	1249: uint16(108),
	1250: uint16(1),
	1251: uint16(anon_sym_DQUOTE2),
	1252: uint16(110),
	1253: uint16(37),
	1254: uint16(aux_sym_content_token1),
	1255: uint16(anon_sym_BSLASH_DQUOTE),
	1256: uint16(anon_sym_STAR),
	1257: uint16(anon_sym_QMARK),
	1258: uint16(anon_sym_PIPE),
	1259: uint16(aux_sym_pattern_token1),
	1260: uint16(anon_sym_PERCENTk),
	1261: uint16(anon_sym_PERCENTn),
	1262: uint16(anon_sym_PERCENTp),
	1263: uint16(anon_sym_PERCENTb),
	1264: uint16(anon_sym_PERCENTs),
	1265: uint16(anon_sym_PERCENTE),
	1266: uint16(anon_sym_PERCENTM),
	1267: uint16(anon_sym_PERCENTm),
	1268: uint16(anon_sym_PERCENTc),
	1269: uint16(anon_sym_PERCENTP),
	1270: uint16(anon_sym_PERCENTr),
	1271: uint16(anon_sym_PERCENTS),
	1272: uint16(anon_sym_PERCENTN),
	1273: uint16(anon_sym_PERCENT_PERCENT),
	1274: uint16(anon_sym_DOLLARkernel),
	1275: uint16(anon_sym_DOLLARnumber),
	1276: uint16(anon_sym_DOLLARdevpath),
	1277: uint16(anon_sym_DOLLARid),
	1278: uint16(anon_sym_DOLLARdriver),
	1279: uint16(anon_sym_DOLLARattr),
	1280: uint16(anon_sym_DOLLARenv),
	1281: uint16(anon_sym_DOLLARmajor),
	1282: uint16(anon_sym_DOLLARminor),
	1283: uint16(anon_sym_DOLLARresult),
	1284: uint16(anon_sym_DOLLARparent),
	1285: uint16(anon_sym_DOLLARname),
	1286: uint16(anon_sym_DOLLARlinks),
	1287: uint16(anon_sym_DOLLARroot),
	1288: uint16(anon_sym_DOLLARsys),
	1289: uint16(anon_sym_DOLLARdevnode),
	1290: uint16(anon_sym_DOLLAR_DOLLAR),
	1291: uint16(3),
	1292: uint16(67),
	1293: uint16(1),
	1294: uint16(sym_linebreak),
	1295: uint16(190),
	1296: uint16(1),
	1297: uint16(anon_sym_DQUOTE2),
	1298: uint16(192),
	1299: uint16(37),
	1300: uint16(aux_sym_content_token1),
	1301: uint16(anon_sym_BSLASH_DQUOTE),
	1302: uint16(anon_sym_STAR),
	1303: uint16(anon_sym_QMARK),
	1304: uint16(anon_sym_PIPE),
	1305: uint16(aux_sym_pattern_token1),
	1306: uint16(anon_sym_PERCENTk),
	1307: uint16(anon_sym_PERCENTn),
	1308: uint16(anon_sym_PERCENTp),
	1309: uint16(anon_sym_PERCENTb),
	1310: uint16(anon_sym_PERCENTs),
	1311: uint16(anon_sym_PERCENTE),
	1312: uint16(anon_sym_PERCENTM),
	1313: uint16(anon_sym_PERCENTm),
	1314: uint16(anon_sym_PERCENTc),
	1315: uint16(anon_sym_PERCENTP),
	1316: uint16(anon_sym_PERCENTr),
	1317: uint16(anon_sym_PERCENTS),
	1318: uint16(anon_sym_PERCENTN),
	1319: uint16(anon_sym_PERCENT_PERCENT),
	1320: uint16(anon_sym_DOLLARkernel),
	1321: uint16(anon_sym_DOLLARnumber),
	1322: uint16(anon_sym_DOLLARdevpath),
	1323: uint16(anon_sym_DOLLARid),
	1324: uint16(anon_sym_DOLLARdriver),
	1325: uint16(anon_sym_DOLLARattr),
	1326: uint16(anon_sym_DOLLARenv),
	1327: uint16(anon_sym_DOLLARmajor),
	1328: uint16(anon_sym_DOLLARminor),
	1329: uint16(anon_sym_DOLLARresult),
	1330: uint16(anon_sym_DOLLARparent),
	1331: uint16(anon_sym_DOLLARname),
	1332: uint16(anon_sym_DOLLARlinks),
	1333: uint16(anon_sym_DOLLARroot),
	1334: uint16(anon_sym_DOLLARsys),
	1335: uint16(anon_sym_DOLLARdevnode),
	1336: uint16(anon_sym_DOLLAR_DOLLAR),
	1337: uint16(3),
	1338: uint16(67),
	1339: uint16(1),
	1340: uint16(sym_linebreak),
	1341: uint16(198),
	1342: uint16(1),
	1343: uint16(anon_sym_DQUOTE2),
	1344: uint16(200),
	1345: uint16(37),
	1346: uint16(aux_sym_content_token1),
	1347: uint16(anon_sym_BSLASH_DQUOTE),
	1348: uint16(anon_sym_STAR),
	1349: uint16(anon_sym_QMARK),
	1350: uint16(anon_sym_PIPE),
	1351: uint16(aux_sym_pattern_token1),
	1352: uint16(anon_sym_PERCENTk),
	1353: uint16(anon_sym_PERCENTn),
	1354: uint16(anon_sym_PERCENTp),
	1355: uint16(anon_sym_PERCENTb),
	1356: uint16(anon_sym_PERCENTs),
	1357: uint16(anon_sym_PERCENTE),
	1358: uint16(anon_sym_PERCENTM),
	1359: uint16(anon_sym_PERCENTm),
	1360: uint16(anon_sym_PERCENTc),
	1361: uint16(anon_sym_PERCENTP),
	1362: uint16(anon_sym_PERCENTr),
	1363: uint16(anon_sym_PERCENTS),
	1364: uint16(anon_sym_PERCENTN),
	1365: uint16(anon_sym_PERCENT_PERCENT),
	1366: uint16(anon_sym_DOLLARkernel),
	1367: uint16(anon_sym_DOLLARnumber),
	1368: uint16(anon_sym_DOLLARdevpath),
	1369: uint16(anon_sym_DOLLARid),
	1370: uint16(anon_sym_DOLLARdriver),
	1371: uint16(anon_sym_DOLLARattr),
	1372: uint16(anon_sym_DOLLARenv),
	1373: uint16(anon_sym_DOLLARmajor),
	1374: uint16(anon_sym_DOLLARminor),
	1375: uint16(anon_sym_DOLLARresult),
	1376: uint16(anon_sym_DOLLARparent),
	1377: uint16(anon_sym_DOLLARname),
	1378: uint16(anon_sym_DOLLARlinks),
	1379: uint16(anon_sym_DOLLARroot),
	1380: uint16(anon_sym_DOLLARsys),
	1381: uint16(anon_sym_DOLLARdevnode),
	1382: uint16(anon_sym_DOLLAR_DOLLAR),
	1383: uint16(3),
	1384: uint16(67),
	1385: uint16(1),
	1386: uint16(sym_linebreak),
	1387: uint16(194),
	1388: uint16(1),
	1389: uint16(anon_sym_DQUOTE2),
	1390: uint16(196),
	1391: uint16(37),
	1392: uint16(aux_sym_content_token1),
	1393: uint16(anon_sym_BSLASH_DQUOTE),
	1394: uint16(anon_sym_STAR),
	1395: uint16(anon_sym_QMARK),
	1396: uint16(anon_sym_PIPE),
	1397: uint16(aux_sym_pattern_token1),
	1398: uint16(anon_sym_PERCENTk),
	1399: uint16(anon_sym_PERCENTn),
	1400: uint16(anon_sym_PERCENTp),
	1401: uint16(anon_sym_PERCENTb),
	1402: uint16(anon_sym_PERCENTs),
	1403: uint16(anon_sym_PERCENTE),
	1404: uint16(anon_sym_PERCENTM),
	1405: uint16(anon_sym_PERCENTm),
	1406: uint16(anon_sym_PERCENTc),
	1407: uint16(anon_sym_PERCENTP),
	1408: uint16(anon_sym_PERCENTr),
	1409: uint16(anon_sym_PERCENTS),
	1410: uint16(anon_sym_PERCENTN),
	1411: uint16(anon_sym_PERCENT_PERCENT),
	1412: uint16(anon_sym_DOLLARkernel),
	1413: uint16(anon_sym_DOLLARnumber),
	1414: uint16(anon_sym_DOLLARdevpath),
	1415: uint16(anon_sym_DOLLARid),
	1416: uint16(anon_sym_DOLLARdriver),
	1417: uint16(anon_sym_DOLLARattr),
	1418: uint16(anon_sym_DOLLARenv),
	1419: uint16(anon_sym_DOLLARmajor),
	1420: uint16(anon_sym_DOLLARminor),
	1421: uint16(anon_sym_DOLLARresult),
	1422: uint16(anon_sym_DOLLARparent),
	1423: uint16(anon_sym_DOLLARname),
	1424: uint16(anon_sym_DOLLARlinks),
	1425: uint16(anon_sym_DOLLARroot),
	1426: uint16(anon_sym_DOLLARsys),
	1427: uint16(anon_sym_DOLLARdevnode),
	1428: uint16(anon_sym_DOLLAR_DOLLAR),
	1429: uint16(3),
	1430: uint16(67),
	1431: uint16(1),
	1432: uint16(sym_linebreak),
	1433: uint16(178),
	1434: uint16(1),
	1435: uint16(anon_sym_DQUOTE2),
	1436: uint16(180),
	1437: uint16(37),
	1438: uint16(aux_sym_content_token1),
	1439: uint16(anon_sym_BSLASH_DQUOTE),
	1440: uint16(anon_sym_STAR),
	1441: uint16(anon_sym_QMARK),
	1442: uint16(anon_sym_PIPE),
	1443: uint16(aux_sym_pattern_token1),
	1444: uint16(anon_sym_PERCENTk),
	1445: uint16(anon_sym_PERCENTn),
	1446: uint16(anon_sym_PERCENTp),
	1447: uint16(anon_sym_PERCENTb),
	1448: uint16(anon_sym_PERCENTs),
	1449: uint16(anon_sym_PERCENTE),
	1450: uint16(anon_sym_PERCENTM),
	1451: uint16(anon_sym_PERCENTm),
	1452: uint16(anon_sym_PERCENTc),
	1453: uint16(anon_sym_PERCENTP),
	1454: uint16(anon_sym_PERCENTr),
	1455: uint16(anon_sym_PERCENTS),
	1456: uint16(anon_sym_PERCENTN),
	1457: uint16(anon_sym_PERCENT_PERCENT),
	1458: uint16(anon_sym_DOLLARkernel),
	1459: uint16(anon_sym_DOLLARnumber),
	1460: uint16(anon_sym_DOLLARdevpath),
	1461: uint16(anon_sym_DOLLARid),
	1462: uint16(anon_sym_DOLLARdriver),
	1463: uint16(anon_sym_DOLLARattr),
	1464: uint16(anon_sym_DOLLARenv),
	1465: uint16(anon_sym_DOLLARmajor),
	1466: uint16(anon_sym_DOLLARminor),
	1467: uint16(anon_sym_DOLLARresult),
	1468: uint16(anon_sym_DOLLARparent),
	1469: uint16(anon_sym_DOLLARname),
	1470: uint16(anon_sym_DOLLARlinks),
	1471: uint16(anon_sym_DOLLARroot),
	1472: uint16(anon_sym_DOLLARsys),
	1473: uint16(anon_sym_DOLLARdevnode),
	1474: uint16(anon_sym_DOLLAR_DOLLAR),
	1475: uint16(25),
	1476: uint16(3),
	1477: uint16(1),
	1478: uint16(sym_linebreak),
	1479: uint16(210),
	1480: uint16(1),
	1482: uint16(212),
	1483: uint16(1),
	1484: uint16(aux_sym_rules_token1),
	1485: uint16(224),
	1486: uint16(1),
	1487: uint16(anon_sym_ATTR),
	1488: uint16(227),
	1489: uint16(1),
	1490: uint16(anon_sym_ATTRS),
	1491: uint16(230),
	1492: uint16(1),
	1493: uint16(anon_sym_SYSCTL),
	1494: uint16(233),
	1495: uint16(1),
	1496: uint16(anon_sym_ENV),
	1497: uint16(236),
	1498: uint16(1),
	1499: uint16(anon_sym_CONST),
	1500: uint16(239),
	1501: uint16(1),
	1502: uint16(anon_sym_TAG),
	1503: uint16(242),
	1504: uint16(1),
	1505: uint16(anon_sym_TEST),
	1506: uint16(245),
	1507: uint16(1),
	1508: uint16(anon_sym_PROGRAM),
	1509: uint16(251),
	1510: uint16(1),
	1511: uint16(anon_sym_SECLABEL),
	1512: uint16(254),
	1513: uint16(1),
	1514: uint16(anon_sym_RUN),
	1515: uint16(260),
	1516: uint16(1),
	1517: uint16(anon_sym_IMPORT),
	1518: uint16(263),
	1519: uint16(1),
	1520: uint16(sym_comment),
	1521: uint16(29),
	1522: uint16(1),
	1523: uint16(aux_sym_rules_repeat1),
	1524: uint16(31),
	1525: uint16(1),
	1526: uint16(aux_sym_rule_repeat1),
	1527: uint16(70),
	1528: uint16(1),
	1529: uint16(sym_assignment),
	1530: uint16(143),
	1531: uint16(1),
	1532: uint16(sym_rule),
	1533: uint16(151),
	1534: uint16(1),
	1535: uint16(sym_match),
	1536: uint16(221),
	1537: uint16(2),
	1538: uint16(anon_sym_NAME),
	1539: uint16(anon_sym_SYMLINK),
	1540: uint16(218),
	1541: uint16(3),
	1542: uint16(anon_sym_KERNEL),
	1543: uint16(anon_sym_SUBSYSTEM),
	1544: uint16(anon_sym_DRIVER),
	1545: uint16(248),
	1546: uint16(3),
	1547: uint16(anon_sym_OWNER),
	1548: uint16(anon_sym_GROUP),
	1549: uint16(anon_sym_MODE),
	1550: uint16(257),
	1551: uint16(3),
	1552: uint16(anon_sym_LABEL),
	1553: uint16(anon_sym_GOTO),
	1554: uint16(anon_sym_OPTIONS),
	1555: uint16(215),
	1556: uint16(7),
	1557: uint16(anon_sym_ACTION),
	1558: uint16(anon_sym_DEVPATH),
	1559: uint16(anon_sym_KERNELS),
	1560: uint16(anon_sym_SUBSYSTEMS),
	1561: uint16(anon_sym_DRIVERS),
	1562: uint16(anon_sym_TAGS),
	1563: uint16(anon_sym_RESULT),
	1564: uint16(25),
	1565: uint16(3),
	1566: uint16(1),
	1567: uint16(sym_linebreak),
	1568: uint16(15),
	1569: uint16(1),
	1570: uint16(anon_sym_ATTR),
	1571: uint16(17),
	1572: uint16(1),
	1573: uint16(anon_sym_ATTRS),
	1574: uint16(19),
	1575: uint16(1),
	1576: uint16(anon_sym_SYSCTL),
	1577: uint16(21),
	1578: uint16(1),
	1579: uint16(anon_sym_ENV),
	1580: uint16(23),
	1581: uint16(1),
	1582: uint16(anon_sym_CONST),
	1583: uint16(25),
	1584: uint16(1),
	1585: uint16(anon_sym_TAG),
	1586: uint16(27),
	1587: uint16(1),
	1588: uint16(anon_sym_TEST),
	1589: uint16(29),
	1590: uint16(1),
	1591: uint16(anon_sym_PROGRAM),
	1592: uint16(33),
	1593: uint16(1),
	1594: uint16(anon_sym_SECLABEL),
	1595: uint16(35),
	1596: uint16(1),
	1597: uint16(anon_sym_RUN),
	1598: uint16(39),
	1599: uint16(1),
	1600: uint16(anon_sym_IMPORT),
	1601: uint16(41),
	1602: uint16(1),
	1603: uint16(sym_comment),
	1604: uint16(266),
	1605: uint16(1),
	1607: uint16(268),
	1608: uint16(1),
	1609: uint16(aux_sym_rules_token1),
	1610: uint16(29),
	1611: uint16(1),
	1612: uint16(aux_sym_rules_repeat1),
	1613: uint16(31),
	1614: uint16(1),
	1615: uint16(aux_sym_rule_repeat1),
	1616: uint16(70),
	1617: uint16(1),
	1618: uint16(sym_assignment),
	1619: uint16(143),
	1620: uint16(1),
	1621: uint16(sym_rule),
	1622: uint16(151),
	1623: uint16(1),
	1624: uint16(sym_match),
	1625: uint16(13),
	1626: uint16(2),
	1627: uint16(anon_sym_NAME),
	1628: uint16(anon_sym_SYMLINK),
	1629: uint16(11),
	1630: uint16(3),
	1631: uint16(anon_sym_KERNEL),
	1632: uint16(anon_sym_SUBSYSTEM),
	1633: uint16(anon_sym_DRIVER),
	1634: uint16(31),
	1635: uint16(3),
	1636: uint16(anon_sym_OWNER),
	1637: uint16(anon_sym_GROUP),
	1638: uint16(anon_sym_MODE),
	1639: uint16(37),
	1640: uint16(3),
	1641: uint16(anon_sym_LABEL),
	1642: uint16(anon_sym_GOTO),
	1643: uint16(anon_sym_OPTIONS),
	1644: uint16(9),
	1645: uint16(7),
	1646: uint16(anon_sym_ACTION),
	1647: uint16(anon_sym_DEVPATH),
	1648: uint16(anon_sym_KERNELS),
	1649: uint16(anon_sym_SUBSYSTEMS),
	1650: uint16(anon_sym_DRIVERS),
	1651: uint16(anon_sym_TAGS),
	1652: uint16(anon_sym_RESULT),
	1653: uint16(20),
	1654: uint16(3),
	1655: uint16(1),
	1656: uint16(sym_linebreak),
	1657: uint16(15),
	1658: uint16(1),
	1659: uint16(anon_sym_ATTR),
	1660: uint16(17),
	1661: uint16(1),
	1662: uint16(anon_sym_ATTRS),
	1663: uint16(19),
	1664: uint16(1),
	1665: uint16(anon_sym_SYSCTL),
	1666: uint16(21),
	1667: uint16(1),
	1668: uint16(anon_sym_ENV),
	1669: uint16(23),
	1670: uint16(1),
	1671: uint16(anon_sym_CONST),
	1672: uint16(25),
	1673: uint16(1),
	1674: uint16(anon_sym_TAG),
	1675: uint16(27),
	1676: uint16(1),
	1677: uint16(anon_sym_TEST),
	1678: uint16(29),
	1679: uint16(1),
	1680: uint16(anon_sym_PROGRAM),
	1681: uint16(33),
	1682: uint16(1),
	1683: uint16(anon_sym_SECLABEL),
	1684: uint16(35),
	1685: uint16(1),
	1686: uint16(anon_sym_RUN),
	1687: uint16(39),
	1688: uint16(1),
	1689: uint16(anon_sym_IMPORT),
	1690: uint16(34),
	1691: uint16(1),
	1692: uint16(aux_sym_rule_repeat1),
	1693: uint16(60),
	1694: uint16(1),
	1695: uint16(sym_assignment),
	1696: uint16(151),
	1697: uint16(1),
	1698: uint16(sym_match),
	1699: uint16(13),
	1700: uint16(2),
	1701: uint16(anon_sym_NAME),
	1702: uint16(anon_sym_SYMLINK),
	1703: uint16(11),
	1704: uint16(3),
	1705: uint16(anon_sym_KERNEL),
	1706: uint16(anon_sym_SUBSYSTEM),
	1707: uint16(anon_sym_DRIVER),
	1708: uint16(31),
	1709: uint16(3),
	1710: uint16(anon_sym_OWNER),
	1711: uint16(anon_sym_GROUP),
	1712: uint16(anon_sym_MODE),
	1713: uint16(37),
	1714: uint16(3),
	1715: uint16(anon_sym_LABEL),
	1716: uint16(anon_sym_GOTO),
	1717: uint16(anon_sym_OPTIONS),
	1718: uint16(9),
	1719: uint16(7),
	1720: uint16(anon_sym_ACTION),
	1721: uint16(anon_sym_DEVPATH),
	1722: uint16(anon_sym_KERNELS),
	1723: uint16(anon_sym_SUBSYSTEMS),
	1724: uint16(anon_sym_DRIVERS),
	1725: uint16(anon_sym_TAGS),
	1726: uint16(anon_sym_RESULT),
	1727: uint16(3),
	1728: uint16(3),
	1729: uint16(1),
	1730: uint16(sym_linebreak),
	1731: uint16(270),
	1732: uint16(5),
	1733: uint16(anon_sym_KERNEL),
	1734: uint16(anon_sym_SUBSYSTEM),
	1735: uint16(anon_sym_DRIVER),
	1736: uint16(anon_sym_ATTR),
	1737: uint16(anon_sym_TAG),
	1738: uint16(210),
	1739: uint16(27),
	1741: uint16(aux_sym_rules_token1),
	1742: uint16(anon_sym_ACTION),
	1743: uint16(anon_sym_DEVPATH),
	1744: uint16(anon_sym_KERNELS),
	1745: uint16(anon_sym_NAME),
	1746: uint16(anon_sym_SYMLINK),
	1747: uint16(anon_sym_SUBSYSTEMS),
	1748: uint16(anon_sym_DRIVERS),
	1749: uint16(anon_sym_ATTRS),
	1750: uint16(anon_sym_SYSCTL),
	1751: uint16(anon_sym_ENV),
	1752: uint16(anon_sym_CONST),
	1753: uint16(anon_sym_TAGS),
	1754: uint16(anon_sym_TEST),
	1755: uint16(anon_sym_PROGRAM),
	1756: uint16(anon_sym_RESULT),
	1757: uint16(anon_sym_OWNER),
	1758: uint16(anon_sym_GROUP),
	1759: uint16(anon_sym_MODE),
	1760: uint16(anon_sym_SECLABEL),
	1761: uint16(anon_sym_RUN),
	1762: uint16(anon_sym_LABEL),
	1763: uint16(anon_sym_GOTO),
	1764: uint16(anon_sym_IMPORT),
	1765: uint16(anon_sym_OPTIONS),
	1766: uint16(sym_comment),
	1767: uint16(18),
	1768: uint16(3),
	1769: uint16(1),
	1770: uint16(sym_linebreak),
	1771: uint16(15),
	1772: uint16(1),
	1773: uint16(anon_sym_ATTR),
	1774: uint16(17),
	1775: uint16(1),
	1776: uint16(anon_sym_ATTRS),
	1777: uint16(19),
	1778: uint16(1),
	1779: uint16(anon_sym_SYSCTL),
	1780: uint16(21),
	1781: uint16(1),
	1782: uint16(anon_sym_ENV),
	1783: uint16(23),
	1784: uint16(1),
	1785: uint16(anon_sym_CONST),
	1786: uint16(25),
	1787: uint16(1),
	1788: uint16(anon_sym_TAG),
	1789: uint16(27),
	1790: uint16(1),
	1791: uint16(anon_sym_TEST),
	1792: uint16(29),
	1793: uint16(1),
	1794: uint16(anon_sym_PROGRAM),
	1795: uint16(33),
	1796: uint16(1),
	1797: uint16(anon_sym_SECLABEL),
	1798: uint16(35),
	1799: uint16(1),
	1800: uint16(anon_sym_RUN),
	1801: uint16(39),
	1802: uint16(1),
	1803: uint16(anon_sym_IMPORT),
	1804: uint16(13),
	1805: uint16(2),
	1806: uint16(anon_sym_NAME),
	1807: uint16(anon_sym_SYMLINK),
	1808: uint16(94),
	1809: uint16(2),
	1810: uint16(sym_match),
	1811: uint16(sym_assignment),
	1812: uint16(11),
	1813: uint16(3),
	1814: uint16(anon_sym_KERNEL),
	1815: uint16(anon_sym_SUBSYSTEM),
	1816: uint16(anon_sym_DRIVER),
	1817: uint16(31),
	1818: uint16(3),
	1819: uint16(anon_sym_OWNER),
	1820: uint16(anon_sym_GROUP),
	1821: uint16(anon_sym_MODE),
	1822: uint16(37),
	1823: uint16(3),
	1824: uint16(anon_sym_LABEL),
	1825: uint16(anon_sym_GOTO),
	1826: uint16(anon_sym_OPTIONS),
	1827: uint16(9),
	1828: uint16(7),
	1829: uint16(anon_sym_ACTION),
	1830: uint16(anon_sym_DEVPATH),
	1831: uint16(anon_sym_KERNELS),
	1832: uint16(anon_sym_SUBSYSTEMS),
	1833: uint16(anon_sym_DRIVERS),
	1834: uint16(anon_sym_TAGS),
	1835: uint16(anon_sym_RESULT),
	1836: uint16(14),
	1837: uint16(3),
	1838: uint16(1),
	1839: uint16(sym_linebreak),
	1840: uint16(281),
	1841: uint16(1),
	1842: uint16(anon_sym_ATTR),
	1843: uint16(284),
	1844: uint16(1),
	1845: uint16(anon_sym_ATTRS),
	1846: uint16(287),
	1847: uint16(1),
	1848: uint16(anon_sym_SYSCTL),
	1849: uint16(290),
	1850: uint16(1),
	1851: uint16(anon_sym_ENV),
	1852: uint16(293),
	1853: uint16(1),
	1854: uint16(anon_sym_CONST),
	1855: uint16(296),
	1856: uint16(1),
	1857: uint16(anon_sym_TEST),
	1858: uint16(299),
	1859: uint16(1),
	1860: uint16(anon_sym_PROGRAM),
	1861: uint16(34),
	1862: uint16(1),
	1863: uint16(aux_sym_rule_repeat1),
	1864: uint16(151),
	1865: uint16(1),
	1866: uint16(sym_match),
	1867: uint16(278),
	1868: uint16(2),
	1869: uint16(anon_sym_NAME),
	1870: uint16(anon_sym_SYMLINK),
	1871: uint16(275),
	1872: uint16(4),
	1873: uint16(anon_sym_KERNEL),
	1874: uint16(anon_sym_SUBSYSTEM),
	1875: uint16(anon_sym_DRIVER),
	1876: uint16(anon_sym_TAG),
	1877: uint16(272),
	1878: uint16(7),
	1879: uint16(anon_sym_ACTION),
	1880: uint16(anon_sym_DEVPATH),
	1881: uint16(anon_sym_KERNELS),
	1882: uint16(anon_sym_SUBSYSTEMS),
	1883: uint16(anon_sym_DRIVERS),
	1884: uint16(anon_sym_TAGS),
	1885: uint16(anon_sym_RESULT),
	1886: uint16(302),
	1887: uint16(9),
	1888: uint16(anon_sym_OWNER),
	1889: uint16(anon_sym_GROUP),
	1890: uint16(anon_sym_MODE),
	1891: uint16(anon_sym_SECLABEL),
	1892: uint16(anon_sym_RUN),
	1893: uint16(anon_sym_LABEL),
	1894: uint16(anon_sym_GOTO),
	1895: uint16(anon_sym_IMPORT),
	1896: uint16(anon_sym_OPTIONS),
	1897: uint16(3),
	1898: uint16(3),
	1899: uint16(1),
	1900: uint16(sym_linebreak),
	1901: uint16(304),
	1902: uint16(5),
	1903: uint16(anon_sym_KERNEL),
	1904: uint16(anon_sym_SUBSYSTEM),
	1905: uint16(anon_sym_DRIVER),
	1906: uint16(anon_sym_ATTR),
	1907: uint16(anon_sym_TAG),
	1908: uint16(302),
	1909: uint16(24),
	1910: uint16(anon_sym_ACTION),
	1911: uint16(anon_sym_DEVPATH),
	1912: uint16(anon_sym_KERNELS),
	1913: uint16(anon_sym_NAME),
	1914: uint16(anon_sym_SYMLINK),
	1915: uint16(anon_sym_SUBSYSTEMS),
	1916: uint16(anon_sym_DRIVERS),
	1917: uint16(anon_sym_ATTRS),
	1918: uint16(anon_sym_SYSCTL),
	1919: uint16(anon_sym_ENV),
	1920: uint16(anon_sym_CONST),
	1921: uint16(anon_sym_TAGS),
	1922: uint16(anon_sym_TEST),
	1923: uint16(anon_sym_PROGRAM),
	1924: uint16(anon_sym_RESULT),
	1925: uint16(anon_sym_OWNER),
	1926: uint16(anon_sym_GROUP),
	1927: uint16(anon_sym_MODE),
	1928: uint16(anon_sym_SECLABEL),
	1929: uint16(anon_sym_RUN),
	1930: uint16(anon_sym_LABEL),
	1931: uint16(anon_sym_GOTO),
	1932: uint16(anon_sym_IMPORT),
	1933: uint16(anon_sym_OPTIONS),
	1934: uint16(9),
	1935: uint16(3),
	1936: uint16(1),
	1937: uint16(sym_linebreak),
	1938: uint16(306),
	1939: uint16(1),
	1940: uint16(aux_sym_attribute_token1),
	1941: uint16(312),
	1942: uint16(1),
	1943: uint16(anon_sym_PERCENTs),
	1944: uint16(314),
	1945: uint16(1),
	1946: uint16(anon_sym_PERCENTE),
	1947: uint16(316),
	1948: uint16(1),
	1949: uint16(anon_sym_PERCENTc),
	1950: uint16(152),
	1951: uint16(1),
	1952: uint16(sym_attribute),
	1953: uint16(40),
	1954: uint16(3),
	1955: uint16(sym_pattern),
	1956: uint16(sym_fmt_sub),
	1957: uint16(aux_sym_attribute_repeat1),
	1958: uint16(308),
	1959: uint16(4),
	1960: uint16(anon_sym_STAR),
	1961: uint16(anon_sym_QMARK),
	1962: uint16(anon_sym_PIPE),
	1963: uint16(aux_sym_pattern_token1),
	1964: uint16(310),
	1965: uint16(11),
	1966: uint16(anon_sym_PERCENTk),
	1967: uint16(anon_sym_PERCENTn),
	1968: uint16(anon_sym_PERCENTp),
	1969: uint16(anon_sym_PERCENTb),
	1970: uint16(anon_sym_PERCENTM),
	1971: uint16(anon_sym_PERCENTm),
	1972: uint16(anon_sym_PERCENTP),
	1973: uint16(anon_sym_PERCENTr),
	1974: uint16(anon_sym_PERCENTS),
	1975: uint16(anon_sym_PERCENTN),
	1976: uint16(anon_sym_PERCENT_PERCENT),
	1977: uint16(9),
	1978: uint16(3),
	1979: uint16(1),
	1980: uint16(sym_linebreak),
	1981: uint16(318),
	1982: uint16(1),
	1983: uint16(anon_sym_RBRACE),
	1984: uint16(320),
	1985: uint16(1),
	1986: uint16(aux_sym_attribute_token1),
	1987: uint16(329),
	1988: uint16(1),
	1989: uint16(anon_sym_PERCENTs),
	1990: uint16(332),
	1991: uint16(1),
	1992: uint16(anon_sym_PERCENTE),
	1993: uint16(335),
	1994: uint16(1),
	1995: uint16(anon_sym_PERCENTc),
	1996: uint16(37),
	1997: uint16(3),
	1998: uint16(sym_pattern),
	1999: uint16(sym_fmt_sub),
	2000: uint16(aux_sym_attribute_repeat1),
	2001: uint16(323),
	2002: uint16(4),
	2003: uint16(anon_sym_STAR),
	2004: uint16(anon_sym_QMARK),
	2005: uint16(anon_sym_PIPE),
	2006: uint16(aux_sym_pattern_token1),
	2007: uint16(326),
	2008: uint16(11),
	2009: uint16(anon_sym_PERCENTk),
	2010: uint16(anon_sym_PERCENTn),
	2011: uint16(anon_sym_PERCENTp),
	2012: uint16(anon_sym_PERCENTb),
	2013: uint16(anon_sym_PERCENTM),
	2014: uint16(anon_sym_PERCENTm),
	2015: uint16(anon_sym_PERCENTP),
	2016: uint16(anon_sym_PERCENTr),
	2017: uint16(anon_sym_PERCENTS),
	2018: uint16(anon_sym_PERCENTN),
	2019: uint16(anon_sym_PERCENT_PERCENT),
	2020: uint16(9),
	2021: uint16(3),
	2022: uint16(1),
	2023: uint16(sym_linebreak),
	2024: uint16(306),
	2025: uint16(1),
	2026: uint16(aux_sym_attribute_token1),
	2027: uint16(312),
	2028: uint16(1),
	2029: uint16(anon_sym_PERCENTs),
	2030: uint16(314),
	2031: uint16(1),
	2032: uint16(anon_sym_PERCENTE),
	2033: uint16(316),
	2034: uint16(1),
	2035: uint16(anon_sym_PERCENTc),
	2036: uint16(112),
	2037: uint16(1),
	2038: uint16(sym_attribute),
	2039: uint16(40),
	2040: uint16(3),
	2041: uint16(sym_pattern),
	2042: uint16(sym_fmt_sub),
	2043: uint16(aux_sym_attribute_repeat1),
	2044: uint16(308),
	2045: uint16(4),
	2046: uint16(anon_sym_STAR),
	2047: uint16(anon_sym_QMARK),
	2048: uint16(anon_sym_PIPE),
	2049: uint16(aux_sym_pattern_token1),
	2050: uint16(310),
	2051: uint16(11),
	2052: uint16(anon_sym_PERCENTk),
	2053: uint16(anon_sym_PERCENTn),
	2054: uint16(anon_sym_PERCENTp),
	2055: uint16(anon_sym_PERCENTb),
	2056: uint16(anon_sym_PERCENTM),
	2057: uint16(anon_sym_PERCENTm),
	2058: uint16(anon_sym_PERCENTP),
	2059: uint16(anon_sym_PERCENTr),
	2060: uint16(anon_sym_PERCENTS),
	2061: uint16(anon_sym_PERCENTN),
	2062: uint16(anon_sym_PERCENT_PERCENT),
	2063: uint16(9),
	2064: uint16(3),
	2065: uint16(1),
	2066: uint16(sym_linebreak),
	2067: uint16(306),
	2068: uint16(1),
	2069: uint16(aux_sym_attribute_token1),
	2070: uint16(312),
	2071: uint16(1),
	2072: uint16(anon_sym_PERCENTs),
	2073: uint16(314),
	2074: uint16(1),
	2075: uint16(anon_sym_PERCENTE),
	2076: uint16(316),
	2077: uint16(1),
	2078: uint16(anon_sym_PERCENTc),
	2079: uint16(127),
	2080: uint16(1),
	2081: uint16(sym_attribute),
	2082: uint16(40),
	2083: uint16(3),
	2084: uint16(sym_pattern),
	2085: uint16(sym_fmt_sub),
	2086: uint16(aux_sym_attribute_repeat1),
	2087: uint16(308),
	2088: uint16(4),
	2089: uint16(anon_sym_STAR),
	2090: uint16(anon_sym_QMARK),
	2091: uint16(anon_sym_PIPE),
	2092: uint16(aux_sym_pattern_token1),
	2093: uint16(310),
	2094: uint16(11),
	2095: uint16(anon_sym_PERCENTk),
	2096: uint16(anon_sym_PERCENTn),
	2097: uint16(anon_sym_PERCENTp),
	2098: uint16(anon_sym_PERCENTb),
	2099: uint16(anon_sym_PERCENTM),
	2100: uint16(anon_sym_PERCENTm),
	2101: uint16(anon_sym_PERCENTP),
	2102: uint16(anon_sym_PERCENTr),
	2103: uint16(anon_sym_PERCENTS),
	2104: uint16(anon_sym_PERCENTN),
	2105: uint16(anon_sym_PERCENT_PERCENT),
	2106: uint16(9),
	2107: uint16(3),
	2108: uint16(1),
	2109: uint16(sym_linebreak),
	2110: uint16(312),
	2111: uint16(1),
	2112: uint16(anon_sym_PERCENTs),
	2113: uint16(314),
	2114: uint16(1),
	2115: uint16(anon_sym_PERCENTE),
	2116: uint16(316),
	2117: uint16(1),
	2118: uint16(anon_sym_PERCENTc),
	2119: uint16(338),
	2120: uint16(1),
	2121: uint16(anon_sym_RBRACE),
	2122: uint16(340),
	2123: uint16(1),
	2124: uint16(aux_sym_attribute_token1),
	2125: uint16(37),
	2126: uint16(3),
	2127: uint16(sym_pattern),
	2128: uint16(sym_fmt_sub),
	2129: uint16(aux_sym_attribute_repeat1),
	2130: uint16(308),
	2131: uint16(4),
	2132: uint16(anon_sym_STAR),
	2133: uint16(anon_sym_QMARK),
	2134: uint16(anon_sym_PIPE),
	2135: uint16(aux_sym_pattern_token1),
	2136: uint16(310),
	2137: uint16(11),
	2138: uint16(anon_sym_PERCENTk),
	2139: uint16(anon_sym_PERCENTn),
	2140: uint16(anon_sym_PERCENTp),
	2141: uint16(anon_sym_PERCENTb),
	2142: uint16(anon_sym_PERCENTM),
	2143: uint16(anon_sym_PERCENTm),
	2144: uint16(anon_sym_PERCENTP),
	2145: uint16(anon_sym_PERCENTr),
	2146: uint16(anon_sym_PERCENTS),
	2147: uint16(anon_sym_PERCENTN),
	2148: uint16(anon_sym_PERCENT_PERCENT),
	2149: uint16(9),
	2150: uint16(3),
	2151: uint16(1),
	2152: uint16(sym_linebreak),
	2153: uint16(306),
	2154: uint16(1),
	2155: uint16(aux_sym_attribute_token1),
	2156: uint16(312),
	2157: uint16(1),
	2158: uint16(anon_sym_PERCENTs),
	2159: uint16(314),
	2160: uint16(1),
	2161: uint16(anon_sym_PERCENTE),
	2162: uint16(316),
	2163: uint16(1),
	2164: uint16(anon_sym_PERCENTc),
	2165: uint16(125),
	2166: uint16(1),
	2167: uint16(sym_attribute),
	2168: uint16(40),
	2169: uint16(3),
	2170: uint16(sym_pattern),
	2171: uint16(sym_fmt_sub),
	2172: uint16(aux_sym_attribute_repeat1),
	2173: uint16(308),
	2174: uint16(4),
	2175: uint16(anon_sym_STAR),
	2176: uint16(anon_sym_QMARK),
	2177: uint16(anon_sym_PIPE),
	2178: uint16(aux_sym_pattern_token1),
	2179: uint16(310),
	2180: uint16(11),
	2181: uint16(anon_sym_PERCENTk),
	2182: uint16(anon_sym_PERCENTn),
	2183: uint16(anon_sym_PERCENTp),
	2184: uint16(anon_sym_PERCENTb),
	2185: uint16(anon_sym_PERCENTM),
	2186: uint16(anon_sym_PERCENTm),
	2187: uint16(anon_sym_PERCENTP),
	2188: uint16(anon_sym_PERCENTr),
	2189: uint16(anon_sym_PERCENTS),
	2190: uint16(anon_sym_PERCENTN),
	2191: uint16(anon_sym_PERCENT_PERCENT),
	2192: uint16(9),
	2193: uint16(3),
	2194: uint16(1),
	2195: uint16(sym_linebreak),
	2196: uint16(306),
	2197: uint16(1),
	2198: uint16(aux_sym_attribute_token1),
	2199: uint16(312),
	2200: uint16(1),
	2201: uint16(anon_sym_PERCENTs),
	2202: uint16(314),
	2203: uint16(1),
	2204: uint16(anon_sym_PERCENTE),
	2205: uint16(316),
	2206: uint16(1),
	2207: uint16(anon_sym_PERCENTc),
	2208: uint16(142),
	2209: uint16(1),
	2210: uint16(sym_attribute),
	2211: uint16(40),
	2212: uint16(3),
	2213: uint16(sym_pattern),
	2214: uint16(sym_fmt_sub),
	2215: uint16(aux_sym_attribute_repeat1),
	2216: uint16(308),
	2217: uint16(4),
	2218: uint16(anon_sym_STAR),
	2219: uint16(anon_sym_QMARK),
	2220: uint16(anon_sym_PIPE),
	2221: uint16(aux_sym_pattern_token1),
	2222: uint16(310),
	2223: uint16(11),
	2224: uint16(anon_sym_PERCENTk),
	2225: uint16(anon_sym_PERCENTn),
	2226: uint16(anon_sym_PERCENTp),
	2227: uint16(anon_sym_PERCENTb),
	2228: uint16(anon_sym_PERCENTM),
	2229: uint16(anon_sym_PERCENTm),
	2230: uint16(anon_sym_PERCENTP),
	2231: uint16(anon_sym_PERCENTr),
	2232: uint16(anon_sym_PERCENTS),
	2233: uint16(anon_sym_PERCENTN),
	2234: uint16(anon_sym_PERCENT_PERCENT),
	2235: uint16(9),
	2236: uint16(3),
	2237: uint16(1),
	2238: uint16(sym_linebreak),
	2239: uint16(306),
	2240: uint16(1),
	2241: uint16(aux_sym_attribute_token1),
	2242: uint16(312),
	2243: uint16(1),
	2244: uint16(anon_sym_PERCENTs),
	2245: uint16(314),
	2246: uint16(1),
	2247: uint16(anon_sym_PERCENTE),
	2248: uint16(316),
	2249: uint16(1),
	2250: uint16(anon_sym_PERCENTc),
	2251: uint16(144),
	2252: uint16(1),
	2253: uint16(sym_attribute),
	2254: uint16(40),
	2255: uint16(3),
	2256: uint16(sym_pattern),
	2257: uint16(sym_fmt_sub),
	2258: uint16(aux_sym_attribute_repeat1),
	2259: uint16(308),
	2260: uint16(4),
	2261: uint16(anon_sym_STAR),
	2262: uint16(anon_sym_QMARK),
	2263: uint16(anon_sym_PIPE),
	2264: uint16(aux_sym_pattern_token1),
	2265: uint16(310),
	2266: uint16(11),
	2267: uint16(anon_sym_PERCENTk),
	2268: uint16(anon_sym_PERCENTn),
	2269: uint16(anon_sym_PERCENTp),
	2270: uint16(anon_sym_PERCENTb),
	2271: uint16(anon_sym_PERCENTM),
	2272: uint16(anon_sym_PERCENTm),
	2273: uint16(anon_sym_PERCENTP),
	2274: uint16(anon_sym_PERCENTr),
	2275: uint16(anon_sym_PERCENTS),
	2276: uint16(anon_sym_PERCENTN),
	2277: uint16(anon_sym_PERCENT_PERCENT),
	2278: uint16(9),
	2279: uint16(3),
	2280: uint16(1),
	2281: uint16(sym_linebreak),
	2282: uint16(306),
	2283: uint16(1),
	2284: uint16(aux_sym_attribute_token1),
	2285: uint16(312),
	2286: uint16(1),
	2287: uint16(anon_sym_PERCENTs),
	2288: uint16(314),
	2289: uint16(1),
	2290: uint16(anon_sym_PERCENTE),
	2291: uint16(316),
	2292: uint16(1),
	2293: uint16(anon_sym_PERCENTc),
	2294: uint16(148),
	2295: uint16(1),
	2296: uint16(sym_attribute),
	2297: uint16(40),
	2298: uint16(3),
	2299: uint16(sym_pattern),
	2300: uint16(sym_fmt_sub),
	2301: uint16(aux_sym_attribute_repeat1),
	2302: uint16(308),
	2303: uint16(4),
	2304: uint16(anon_sym_STAR),
	2305: uint16(anon_sym_QMARK),
	2306: uint16(anon_sym_PIPE),
	2307: uint16(aux_sym_pattern_token1),
	2308: uint16(310),
	2309: uint16(11),
	2310: uint16(anon_sym_PERCENTk),
	2311: uint16(anon_sym_PERCENTn),
	2312: uint16(anon_sym_PERCENTp),
	2313: uint16(anon_sym_PERCENTb),
	2314: uint16(anon_sym_PERCENTM),
	2315: uint16(anon_sym_PERCENTm),
	2316: uint16(anon_sym_PERCENTP),
	2317: uint16(anon_sym_PERCENTr),
	2318: uint16(anon_sym_PERCENTS),
	2319: uint16(anon_sym_PERCENTN),
	2320: uint16(anon_sym_PERCENT_PERCENT),
	2321: uint16(3),
	2322: uint16(3),
	2323: uint16(1),
	2324: uint16(sym_linebreak),
	2325: uint16(342),
	2326: uint16(1),
	2327: uint16(anon_sym_LBRACE),
	2328: uint16(108),
	2329: uint16(20),
	2330: uint16(anon_sym_RBRACE),
	2331: uint16(aux_sym_attribute_token1),
	2332: uint16(anon_sym_STAR),
	2333: uint16(anon_sym_QMARK),
	2334: uint16(anon_sym_PIPE),
	2335: uint16(aux_sym_pattern_token1),
	2336: uint16(anon_sym_PERCENTk),
	2337: uint16(anon_sym_PERCENTn),
	2338: uint16(anon_sym_PERCENTp),
	2339: uint16(anon_sym_PERCENTb),
	2340: uint16(anon_sym_PERCENTs),
	2341: uint16(anon_sym_PERCENTE),
	2342: uint16(anon_sym_PERCENTM),
	2343: uint16(anon_sym_PERCENTm),
	2344: uint16(anon_sym_PERCENTc),
	2345: uint16(anon_sym_PERCENTP),
	2346: uint16(anon_sym_PERCENTr),
	2347: uint16(anon_sym_PERCENTS),
	2348: uint16(anon_sym_PERCENTN),
	2349: uint16(anon_sym_PERCENT_PERCENT),
	2350: uint16(2),
	2351: uint16(3),
	2352: uint16(1),
	2353: uint16(sym_linebreak),
	2354: uint16(194),
	2355: uint16(20),
	2356: uint16(anon_sym_RBRACE),
	2357: uint16(aux_sym_attribute_token1),
	2358: uint16(anon_sym_STAR),
	2359: uint16(anon_sym_QMARK),
	2360: uint16(anon_sym_PIPE),
	2361: uint16(aux_sym_pattern_token1),
	2362: uint16(anon_sym_PERCENTk),
	2363: uint16(anon_sym_PERCENTn),
	2364: uint16(anon_sym_PERCENTp),
	2365: uint16(anon_sym_PERCENTb),
	2366: uint16(anon_sym_PERCENTs),
	2367: uint16(anon_sym_PERCENTE),
	2368: uint16(anon_sym_PERCENTM),
	2369: uint16(anon_sym_PERCENTm),
	2370: uint16(anon_sym_PERCENTc),
	2371: uint16(anon_sym_PERCENTP),
	2372: uint16(anon_sym_PERCENTr),
	2373: uint16(anon_sym_PERCENTS),
	2374: uint16(anon_sym_PERCENTN),
	2375: uint16(anon_sym_PERCENT_PERCENT),
	2376: uint16(2),
	2377: uint16(3),
	2378: uint16(1),
	2379: uint16(sym_linebreak),
	2380: uint16(190),
	2381: uint16(20),
	2382: uint16(anon_sym_RBRACE),
	2383: uint16(aux_sym_attribute_token1),
	2384: uint16(anon_sym_STAR),
	2385: uint16(anon_sym_QMARK),
	2386: uint16(anon_sym_PIPE),
	2387: uint16(aux_sym_pattern_token1),
	2388: uint16(anon_sym_PERCENTk),
	2389: uint16(anon_sym_PERCENTn),
	2390: uint16(anon_sym_PERCENTp),
	2391: uint16(anon_sym_PERCENTb),
	2392: uint16(anon_sym_PERCENTs),
	2393: uint16(anon_sym_PERCENTE),
	2394: uint16(anon_sym_PERCENTM),
	2395: uint16(anon_sym_PERCENTm),
	2396: uint16(anon_sym_PERCENTc),
	2397: uint16(anon_sym_PERCENTP),
	2398: uint16(anon_sym_PERCENTr),
	2399: uint16(anon_sym_PERCENTS),
	2400: uint16(anon_sym_PERCENTN),
	2401: uint16(anon_sym_PERCENT_PERCENT),
	2402: uint16(2),
	2403: uint16(3),
	2404: uint16(1),
	2405: uint16(sym_linebreak),
	2406: uint16(198),
	2407: uint16(20),
	2408: uint16(anon_sym_RBRACE),
	2409: uint16(aux_sym_attribute_token1),
	2410: uint16(anon_sym_STAR),
	2411: uint16(anon_sym_QMARK),
	2412: uint16(anon_sym_PIPE),
	2413: uint16(aux_sym_pattern_token1),
	2414: uint16(anon_sym_PERCENTk),
	2415: uint16(anon_sym_PERCENTn),
	2416: uint16(anon_sym_PERCENTp),
	2417: uint16(anon_sym_PERCENTb),
	2418: uint16(anon_sym_PERCENTs),
	2419: uint16(anon_sym_PERCENTE),
	2420: uint16(anon_sym_PERCENTM),
	2421: uint16(anon_sym_PERCENTm),
	2422: uint16(anon_sym_PERCENTc),
	2423: uint16(anon_sym_PERCENTP),
	2424: uint16(anon_sym_PERCENTr),
	2425: uint16(anon_sym_PERCENTS),
	2426: uint16(anon_sym_PERCENTN),
	2427: uint16(anon_sym_PERCENT_PERCENT),
	2428: uint16(2),
	2429: uint16(3),
	2430: uint16(1),
	2431: uint16(sym_linebreak),
	2432: uint16(108),
	2433: uint16(20),
	2434: uint16(anon_sym_RBRACE),
	2435: uint16(aux_sym_attribute_token1),
	2436: uint16(anon_sym_STAR),
	2437: uint16(anon_sym_QMARK),
	2438: uint16(anon_sym_PIPE),
	2439: uint16(aux_sym_pattern_token1),
	2440: uint16(anon_sym_PERCENTk),
	2441: uint16(anon_sym_PERCENTn),
	2442: uint16(anon_sym_PERCENTp),
	2443: uint16(anon_sym_PERCENTb),
	2444: uint16(anon_sym_PERCENTs),
	2445: uint16(anon_sym_PERCENTE),
	2446: uint16(anon_sym_PERCENTM),
	2447: uint16(anon_sym_PERCENTm),
	2448: uint16(anon_sym_PERCENTc),
	2449: uint16(anon_sym_PERCENTP),
	2450: uint16(anon_sym_PERCENTr),
	2451: uint16(anon_sym_PERCENTS),
	2452: uint16(anon_sym_PERCENTN),
	2453: uint16(anon_sym_PERCENT_PERCENT),
	2454: uint16(7),
	2455: uint16(67),
	2456: uint16(1),
	2457: uint16(sym_linebreak),
	2458: uint16(344),
	2459: uint16(1),
	2460: uint16(anon_sym_DQUOTE2),
	2461: uint16(346),
	2462: uint16(1),
	2463: uint16(aux_sym_content_token1),
	2464: uint16(51),
	2465: uint16(1),
	2466: uint16(aux_sym__c_content),
	2467: uint16(54),
	2468: uint16(2),
	2469: uint16(sym_pattern),
	2470: uint16(sym_c_escape),
	2471: uint16(348),
	2472: uint16(4),
	2473: uint16(anon_sym_STAR),
	2474: uint16(anon_sym_QMARK),
	2475: uint16(anon_sym_PIPE),
	2476: uint16(aux_sym_pattern_token1),
	2477: uint16(350),
	2478: uint16(5),
	2479: uint16(aux_sym_c_escape_token1),
	2480: uint16(aux_sym_c_escape_token2),
	2481: uint16(aux_sym_c_escape_token3),
	2482: uint16(aux_sym_c_escape_token4),
	2483: uint16(aux_sym_c_escape_token5),
	2484: uint16(7),
	2485: uint16(67),
	2486: uint16(1),
	2487: uint16(sym_linebreak),
	2488: uint16(352),
	2489: uint16(1),
	2490: uint16(anon_sym_DQUOTE2),
	2491: uint16(354),
	2492: uint16(1),
	2493: uint16(aux_sym_content_token1),
	2494: uint16(51),
	2495: uint16(1),
	2496: uint16(aux_sym__c_content),
	2497: uint16(54),
	2498: uint16(2),
	2499: uint16(sym_pattern),
	2500: uint16(sym_c_escape),
	2501: uint16(357),
	2502: uint16(4),
	2503: uint16(anon_sym_STAR),
	2504: uint16(anon_sym_QMARK),
	2505: uint16(anon_sym_PIPE),
	2506: uint16(aux_sym_pattern_token1),
	2507: uint16(360),
	2508: uint16(5),
	2509: uint16(aux_sym_c_escape_token1),
	2510: uint16(aux_sym_c_escape_token2),
	2511: uint16(aux_sym_c_escape_token3),
	2512: uint16(aux_sym_c_escape_token4),
	2513: uint16(aux_sym_c_escape_token5),
	2514: uint16(7),
	2515: uint16(67),
	2516: uint16(1),
	2517: uint16(sym_linebreak),
	2518: uint16(346),
	2519: uint16(1),
	2520: uint16(aux_sym_content_token1),
	2521: uint16(363),
	2522: uint16(1),
	2523: uint16(anon_sym_DQUOTE2),
	2524: uint16(50),
	2525: uint16(1),
	2526: uint16(aux_sym__c_content),
	2527: uint16(54),
	2528: uint16(2),
	2529: uint16(sym_pattern),
	2530: uint16(sym_c_escape),
	2531: uint16(348),
	2532: uint16(4),
	2533: uint16(anon_sym_STAR),
	2534: uint16(anon_sym_QMARK),
	2535: uint16(anon_sym_PIPE),
	2536: uint16(aux_sym_pattern_token1),
	2537: uint16(350),
	2538: uint16(5),
	2539: uint16(aux_sym_c_escape_token1),
	2540: uint16(aux_sym_c_escape_token2),
	2541: uint16(aux_sym_c_escape_token3),
	2542: uint16(aux_sym_c_escape_token4),
	2543: uint16(aux_sym_c_escape_token5),
	2544: uint16(3),
	2545: uint16(67),
	2546: uint16(1),
	2547: uint16(sym_linebreak),
	2548: uint16(182),
	2549: uint16(1),
	2550: uint16(anon_sym_DQUOTE2),
	2551: uint16(184),
	2552: uint16(10),
	2553: uint16(aux_sym_content_token1),
	2554: uint16(anon_sym_STAR),
	2555: uint16(anon_sym_QMARK),
	2556: uint16(anon_sym_PIPE),
	2557: uint16(aux_sym_pattern_token1),
	2558: uint16(aux_sym_c_escape_token1),
	2559: uint16(aux_sym_c_escape_token2),
	2560: uint16(aux_sym_c_escape_token3),
	2561: uint16(aux_sym_c_escape_token4),
	2562: uint16(aux_sym_c_escape_token5),
	2563: uint16(3),
	2564: uint16(67),
	2565: uint16(1),
	2566: uint16(sym_linebreak),
	2567: uint16(365),
	2568: uint16(1),
	2569: uint16(anon_sym_DQUOTE2),
	2570: uint16(367),
	2571: uint16(10),
	2572: uint16(aux_sym_content_token1),
	2573: uint16(anon_sym_STAR),
	2574: uint16(anon_sym_QMARK),
	2575: uint16(anon_sym_PIPE),
	2576: uint16(aux_sym_pattern_token1),
	2577: uint16(aux_sym_c_escape_token1),
	2578: uint16(aux_sym_c_escape_token2),
	2579: uint16(aux_sym_c_escape_token3),
	2580: uint16(aux_sym_c_escape_token4),
	2581: uint16(aux_sym_c_escape_token5),
	2582: uint16(3),
	2583: uint16(67),
	2584: uint16(1),
	2585: uint16(sym_linebreak),
	2586: uint16(198),
	2587: uint16(1),
	2588: uint16(anon_sym_DQUOTE2),
	2589: uint16(200),
	2590: uint16(10),
	2591: uint16(aux_sym_content_token1),
	2592: uint16(anon_sym_STAR),
	2593: uint16(anon_sym_QMARK),
	2594: uint16(anon_sym_PIPE),
	2595: uint16(aux_sym_pattern_token1),
	2596: uint16(aux_sym_c_escape_token1),
	2597: uint16(aux_sym_c_escape_token2),
	2598: uint16(aux_sym_c_escape_token3),
	2599: uint16(aux_sym_c_escape_token4),
	2600: uint16(aux_sym_c_escape_token5),
	2601: uint16(6),
	2602: uint16(67),
	2603: uint16(1),
	2604: uint16(sym_linebreak),
	2605: uint16(369),
	2606: uint16(1),
	2607: uint16(anon_sym_DQUOTE2),
	2608: uint16(140),
	2609: uint16(1),
	2610: uint16(sym_content),
	2611: uint16(371),
	2612: uint16(2),
	2613: uint16(aux_sym_content_token1),
	2614: uint16(anon_sym_BSLASH_DQUOTE),
	2615: uint16(58),
	2616: uint16(2),
	2617: uint16(sym_pattern),
	2618: uint16(aux_sym_content_repeat1),
	2619: uint16(373),
	2620: uint16(4),
	2621: uint16(anon_sym_STAR),
	2622: uint16(anon_sym_QMARK),
	2623: uint16(anon_sym_PIPE),
	2624: uint16(aux_sym_pattern_token1),
	2625: uint16(5),
	2626: uint16(67),
	2627: uint16(1),
	2628: uint16(sym_linebreak),
	2629: uint16(375),
	2630: uint16(1),
	2631: uint16(anon_sym_DQUOTE2),
	2632: uint16(377),
	2633: uint16(2),
	2634: uint16(aux_sym_content_token1),
	2635: uint16(anon_sym_BSLASH_DQUOTE),
	2636: uint16(57),
	2637: uint16(2),
	2638: uint16(sym_pattern),
	2639: uint16(aux_sym_content_repeat1),
	2640: uint16(380),
	2641: uint16(4),
	2642: uint16(anon_sym_STAR),
	2643: uint16(anon_sym_QMARK),
	2644: uint16(anon_sym_PIPE),
	2645: uint16(aux_sym_pattern_token1),
	2646: uint16(5),
	2647: uint16(67),
	2648: uint16(1),
	2649: uint16(sym_linebreak),
	2650: uint16(383),
	2651: uint16(1),
	2652: uint16(anon_sym_DQUOTE2),
	2653: uint16(385),
	2654: uint16(2),
	2655: uint16(aux_sym_content_token1),
	2656: uint16(anon_sym_BSLASH_DQUOTE),
	2657: uint16(57),
	2658: uint16(2),
	2659: uint16(sym_pattern),
	2660: uint16(aux_sym_content_repeat1),
	2661: uint16(373),
	2662: uint16(4),
	2663: uint16(anon_sym_STAR),
	2664: uint16(anon_sym_QMARK),
	2665: uint16(anon_sym_PIPE),
	2666: uint16(aux_sym_pattern_token1),
	2667: uint16(3),
	2668: uint16(67),
	2669: uint16(1),
	2670: uint16(sym_linebreak),
	2671: uint16(198),
	2672: uint16(1),
	2673: uint16(anon_sym_DQUOTE2),
	2674: uint16(200),
	2675: uint16(6),
	2676: uint16(aux_sym_content_token1),
	2677: uint16(anon_sym_BSLASH_DQUOTE),
	2678: uint16(anon_sym_STAR),
	2679: uint16(anon_sym_QMARK),
	2680: uint16(anon_sym_PIPE),
	2681: uint16(aux_sym_pattern_token1),
	2682: uint16(4),
	2683: uint16(3),
	2684: uint16(1),
	2685: uint16(sym_linebreak),
	2686: uint16(387),
	2687: uint16(1),
	2688: uint16(aux_sym_rules_token1),
	2689: uint16(389),
	2690: uint16(1),
	2691: uint16(anon_sym_COMMA),
	2692: uint16(62),
	2693: uint16(1),
	2694: uint16(aux_sym_rule_repeat2),
	2695: uint16(4),
	2696: uint16(3),
	2697: uint16(1),
	2698: uint16(sym_linebreak),
	2699: uint16(391),
	2700: uint16(1),
	2701: uint16(aux_sym_rules_token1),
	2702: uint16(393),
	2703: uint16(1),
	2704: uint16(anon_sym_COMMA),
	2705: uint16(61),
	2706: uint16(1),
	2707: uint16(aux_sym_rule_repeat2),
	2708: uint16(4),
	2709: uint16(3),
	2710: uint16(1),
	2711: uint16(sym_linebreak),
	2712: uint16(389),
	2713: uint16(1),
	2714: uint16(anon_sym_COMMA),
	2715: uint16(396),
	2716: uint16(1),
	2717: uint16(aux_sym_rules_token1),
	2718: uint16(61),
	2719: uint16(1),
	2720: uint16(aux_sym_rule_repeat2),
	2721: uint16(4),
	2722: uint16(3),
	2723: uint16(1),
	2724: uint16(sym_linebreak),
	2725: uint16(398),
	2726: uint16(1),
	2727: uint16(anon_sym_DQUOTE),
	2728: uint16(400),
	2729: uint16(1),
	2730: uint16(anon_sym_e),
	2731: uint16(93),
	2732: uint16(1),
	2733: uint16(sym__sub_value),
	2734: uint16(4),
	2735: uint16(3),
	2736: uint16(1),
	2737: uint16(sym_linebreak),
	2738: uint16(402),
	2739: uint16(1),
	2740: uint16(anon_sym_DQUOTE),
	2741: uint16(404),
	2742: uint16(1),
	2743: uint16(anon_sym_e),
	2744: uint16(92),
	2745: uint16(1),
	2746: uint16(sym_value),
	2747: uint16(4),
	2748: uint16(3),
	2749: uint16(1),
	2750: uint16(sym_linebreak),
	2751: uint16(398),
	2752: uint16(1),
	2753: uint16(anon_sym_DQUOTE),
	2754: uint16(400),
	2755: uint16(1),
	2756: uint16(anon_sym_e),
	2757: uint16(86),
	2758: uint16(1),
	2759: uint16(sym__sub_value),
	2760: uint16(4),
	2761: uint16(3),
	2762: uint16(1),
	2763: uint16(sym_linebreak),
	2764: uint16(398),
	2765: uint16(1),
	2766: uint16(anon_sym_DQUOTE),
	2767: uint16(400),
	2768: uint16(1),
	2769: uint16(anon_sym_e),
	2770: uint16(106),
	2771: uint16(1),
	2772: uint16(sym__sub_value),
	2773: uint16(4),
	2774: uint16(3),
	2775: uint16(1),
	2776: uint16(sym_linebreak),
	2777: uint16(387),
	2778: uint16(1),
	2779: uint16(aux_sym_rules_token1),
	2780: uint16(389),
	2781: uint16(1),
	2782: uint16(anon_sym_COMMA),
	2783: uint16(61),
	2784: uint16(1),
	2785: uint16(aux_sym_rule_repeat2),
	2786: uint16(4),
	2787: uint16(3),
	2788: uint16(1),
	2789: uint16(sym_linebreak),
	2790: uint16(402),
	2791: uint16(1),
	2792: uint16(anon_sym_DQUOTE),
	2793: uint16(404),
	2794: uint16(1),
	2795: uint16(anon_sym_e),
	2796: uint16(86),
	2797: uint16(1),
	2798: uint16(sym_value),
	2799: uint16(4),
	2800: uint16(3),
	2801: uint16(1),
	2802: uint16(sym_linebreak),
	2803: uint16(402),
	2804: uint16(1),
	2805: uint16(anon_sym_DQUOTE),
	2806: uint16(404),
	2807: uint16(1),
	2808: uint16(anon_sym_e),
	2809: uint16(93),
	2810: uint16(1),
	2811: uint16(sym_value),
	2812: uint16(4),
	2813: uint16(3),
	2814: uint16(1),
	2815: uint16(sym_linebreak),
	2816: uint16(389),
	2817: uint16(1),
	2818: uint16(anon_sym_COMMA),
	2819: uint16(406),
	2820: uint16(1),
	2821: uint16(aux_sym_rules_token1),
	2822: uint16(67),
	2823: uint16(1),
	2824: uint16(aux_sym_rule_repeat2),
	2825: uint16(4),
	2826: uint16(3),
	2827: uint16(1),
	2828: uint16(sym_linebreak),
	2829: uint16(398),
	2830: uint16(1),
	2831: uint16(anon_sym_DQUOTE),
	2832: uint16(400),
	2833: uint16(1),
	2834: uint16(anon_sym_e),
	2835: uint16(92),
	2836: uint16(1),
	2837: uint16(sym__sub_value),
	2838: uint16(4),
	2839: uint16(3),
	2840: uint16(1),
	2841: uint16(sym_linebreak),
	2842: uint16(402),
	2843: uint16(1),
	2844: uint16(anon_sym_DQUOTE),
	2845: uint16(404),
	2846: uint16(1),
	2847: uint16(anon_sym_e),
	2848: uint16(106),
	2849: uint16(1),
	2850: uint16(sym_value),
	2851: uint16(3),
	2852: uint16(3),
	2853: uint16(1),
	2854: uint16(sym_linebreak),
	2855: uint16(408),
	2856: uint16(1),
	2857: uint16(anon_sym_LBRACE),
	2858: uint16(410),
	2859: uint16(1),
	2860: uint16(sym_assignment_op),
	2861: uint16(3),
	2862: uint16(3),
	2863: uint16(1),
	2864: uint16(sym_linebreak),
	2865: uint16(412),
	2866: uint16(1),
	2867: uint16(sym_match_op),
	2868: uint16(414),
	2869: uint16(1),
	2870: uint16(sym_assignment_op),
	2871: uint16(3),
	2872: uint16(3),
	2873: uint16(1),
	2874: uint16(sym_linebreak),
	2875: uint16(416),
	2876: uint16(1),
	2877: uint16(aux_sym_env_var_token1),
	2878: uint16(152),
	2879: uint16(1),
	2880: uint16(sym_env_var),
	2881: uint16(3),
	2882: uint16(3),
	2883: uint16(1),
	2884: uint16(sym_linebreak),
	2885: uint16(418),
	2886: uint16(1),
	2887: uint16(sym_match_op),
	2888: uint16(420),
	2889: uint16(1),
	2890: uint16(sym_assignment_op),
	2891: uint16(3),
	2892: uint16(3),
	2893: uint16(1),
	2894: uint16(sym_linebreak),
	2895: uint16(422),
	2896: uint16(1),
	2897: uint16(sym_match_op),
	2898: uint16(424),
	2899: uint16(1),
	2900: uint16(sym_assignment_op),
	2901: uint16(3),
	2902: uint16(3),
	2903: uint16(1),
	2904: uint16(sym_linebreak),
	2905: uint16(420),
	2906: uint16(1),
	2907: uint16(sym_assignment_op),
	2908: uint16(426),
	2909: uint16(1),
	2910: uint16(sym_match_op),
	2911: uint16(3),
	2912: uint16(3),
	2913: uint16(1),
	2914: uint16(sym_linebreak),
	2915: uint16(428),
	2916: uint16(1),
	2917: uint16(aux_sym_env_var_token1),
	2918: uint16(127),
	2919: uint16(1),
	2920: uint16(sym_kernel_param),
	2921: uint16(3),
	2922: uint16(3),
	2923: uint16(1),
	2924: uint16(sym_linebreak),
	2925: uint16(416),
	2926: uint16(1),
	2927: uint16(aux_sym_env_var_token1),
	2928: uint16(109),
	2929: uint16(1),
	2930: uint16(sym_env_var),
	2931: uint16(2),
	2932: uint16(3),
	2933: uint16(1),
	2934: uint16(sym_linebreak),
	2935: uint16(430),
	2936: uint16(2),
	2937: uint16(aux_sym_rules_token1),
	2938: uint16(anon_sym_COMMA),
	2939: uint16(3),
	2940: uint16(3),
	2941: uint16(1),
	2942: uint16(sym_linebreak),
	2943: uint16(416),
	2944: uint16(1),
	2945: uint16(aux_sym_env_var_token1),
	2946: uint16(112),
	2947: uint16(1),
	2948: uint16(sym_env_var),
	2949: uint16(2),
	2950: uint16(3),
	2951: uint16(1),
	2952: uint16(sym_linebreak),
	2953: uint16(432),
	2954: uint16(2),
	2955: uint16(aux_sym_rules_token1),
	2956: uint16(anon_sym_COMMA),
	2957: uint16(2),
	2958: uint16(3),
	2959: uint16(1),
	2960: uint16(sym_linebreak),
	2961: uint16(434),
	2962: uint16(2),
	2963: uint16(aux_sym_rules_token1),
	2964: uint16(anon_sym_COMMA),
	2965: uint16(3),
	2966: uint16(3),
	2967: uint16(1),
	2968: uint16(sym_linebreak),
	2969: uint16(436),
	2970: uint16(1),
	2971: uint16(anon_sym_PLUS),
	2972: uint16(438),
	2973: uint16(1),
	2974: uint16(anon_sym_RBRACE2),
	2975: uint16(2),
	2976: uint16(3),
	2977: uint16(1),
	2978: uint16(sym_linebreak),
	2979: uint16(440),
	2980: uint16(2),
	2981: uint16(aux_sym_rules_token1),
	2982: uint16(anon_sym_COMMA),
	2983: uint16(3),
	2984: uint16(3),
	2985: uint16(1),
	2986: uint16(sym_linebreak),
	2987: uint16(412),
	2988: uint16(1),
	2989: uint16(sym_match_op),
	2990: uint16(442),
	2991: uint16(1),
	2992: uint16(sym_assignment_op),
	2993: uint16(2),
	2994: uint16(3),
	2995: uint16(1),
	2996: uint16(sym_linebreak),
	2997: uint16(444),
	2998: uint16(2),
	2999: uint16(aux_sym_rules_token1),
	3000: uint16(anon_sym_COMMA),
	3001: uint16(3),
	3002: uint16(3),
	3003: uint16(1),
	3004: uint16(sym_linebreak),
	3005: uint16(416),
	3006: uint16(1),
	3007: uint16(aux_sym_env_var_token1),
	3008: uint16(148),
	3009: uint16(1),
	3010: uint16(sym_env_var),
	3011: uint16(3),
	3012: uint16(3),
	3013: uint16(1),
	3014: uint16(sym_linebreak),
	3015: uint16(446),
	3016: uint16(1),
	3017: uint16(anon_sym_RBRACE),
	3018: uint16(448),
	3019: uint16(1),
	3020: uint16(anon_sym_PLUS),
	3021: uint16(2),
	3022: uint16(3),
	3023: uint16(1),
	3024: uint16(sym_linebreak),
	3025: uint16(450),
	3026: uint16(2),
	3027: uint16(aux_sym_rules_token1),
	3028: uint16(anon_sym_COMMA),
	3029: uint16(2),
	3030: uint16(3),
	3031: uint16(1),
	3032: uint16(sym_linebreak),
	3033: uint16(452),
	3034: uint16(2),
	3035: uint16(aux_sym_rules_token1),
	3036: uint16(anon_sym_COMMA),
	3037: uint16(2),
	3038: uint16(3),
	3039: uint16(1),
	3040: uint16(sym_linebreak),
	3041: uint16(454),
	3042: uint16(2),
	3043: uint16(aux_sym_rules_token1),
	3044: uint16(anon_sym_COMMA),
	3045: uint16(2),
	3046: uint16(3),
	3047: uint16(1),
	3048: uint16(sym_linebreak),
	3049: uint16(391),
	3050: uint16(2),
	3051: uint16(aux_sym_rules_token1),
	3052: uint16(anon_sym_COMMA),
	3053: uint16(3),
	3054: uint16(3),
	3055: uint16(1),
	3056: uint16(sym_linebreak),
	3057: uint16(456),
	3058: uint16(1),
	3059: uint16(sym_match_op),
	3060: uint16(458),
	3061: uint16(1),
	3062: uint16(sym_assignment_op),
	3063: uint16(2),
	3064: uint16(3),
	3065: uint16(1),
	3066: uint16(sym_linebreak),
	3067: uint16(460),
	3068: uint16(2),
	3069: uint16(aux_sym_rules_token1),
	3070: uint16(anon_sym_COMMA),
	3071: uint16(3),
	3072: uint16(3),
	3073: uint16(1),
	3074: uint16(sym_linebreak),
	3075: uint16(456),
	3076: uint16(1),
	3077: uint16(sym_match_op),
	3078: uint16(462),
	3079: uint16(1),
	3080: uint16(anon_sym_LBRACE),
	3081: uint16(3),
	3082: uint16(3),
	3083: uint16(1),
	3084: uint16(sym_linebreak),
	3085: uint16(464),
	3086: uint16(1),
	3087: uint16(anon_sym_PLUS),
	3088: uint16(466),
	3089: uint16(1),
	3090: uint16(anon_sym_RBRACE2),
	3091: uint16(3),
	3092: uint16(3),
	3093: uint16(1),
	3094: uint16(sym_linebreak),
	3095: uint16(468),
	3096: uint16(1),
	3097: uint16(anon_sym_RBRACE),
	3098: uint16(470),
	3099: uint16(1),
	3100: uint16(anon_sym_PLUS),
	3101: uint16(3),
	3102: uint16(3),
	3103: uint16(1),
	3104: uint16(sym_linebreak),
	3105: uint16(472),
	3106: uint16(1),
	3107: uint16(anon_sym_PLUS),
	3108: uint16(474),
	3109: uint16(1),
	3110: uint16(anon_sym_RBRACE2),
	3111: uint16(3),
	3112: uint16(3),
	3113: uint16(1),
	3114: uint16(sym_linebreak),
	3115: uint16(428),
	3116: uint16(1),
	3117: uint16(aux_sym_env_var_token1),
	3118: uint16(125),
	3119: uint16(1),
	3120: uint16(sym_kernel_param),
	3121: uint16(3),
	3122: uint16(3),
	3123: uint16(1),
	3124: uint16(sym_linebreak),
	3125: uint16(416),
	3126: uint16(1),
	3127: uint16(aux_sym_env_var_token1),
	3128: uint16(142),
	3129: uint16(1),
	3130: uint16(sym_env_var),
	3131: uint16(2),
	3132: uint16(3),
	3133: uint16(1),
	3134: uint16(sym_linebreak),
	3135: uint16(476),
	3136: uint16(2),
	3137: uint16(aux_sym_rules_token1),
	3138: uint16(anon_sym_COMMA),
	3139: uint16(3),
	3140: uint16(3),
	3141: uint16(1),
	3142: uint16(sym_linebreak),
	3143: uint16(416),
	3144: uint16(1),
	3145: uint16(aux_sym_env_var_token1),
	3146: uint16(144),
	3147: uint16(1),
	3148: uint16(sym_env_var),
	3149: uint16(3),
	3150: uint16(3),
	3151: uint16(1),
	3152: uint16(sym_linebreak),
	3153: uint16(416),
	3154: uint16(1),
	3155: uint16(aux_sym_env_var_token1),
	3156: uint16(130),
	3157: uint16(1),
	3158: uint16(sym_env_var),
	3159: uint16(2),
	3160: uint16(3),
	3161: uint16(1),
	3162: uint16(sym_linebreak),
	3163: uint16(478),
	3164: uint16(2),
	3165: uint16(aux_sym_rules_token1),
	3166: uint16(anon_sym_COMMA),
	3167: uint16(2),
	3168: uint16(3),
	3169: uint16(1),
	3170: uint16(sym_linebreak),
	3171: uint16(480),
	3172: uint16(1),
	3173: uint16(anon_sym_LBRACE),
	3174: uint16(2),
	3175: uint16(3),
	3176: uint16(1),
	3177: uint16(sym_linebreak),
	3178: uint16(482),
	3179: uint16(1),
	3180: uint16(anon_sym_DQUOTE2),
	3181: uint16(2),
	3182: uint16(3),
	3183: uint16(1),
	3184: uint16(sym_linebreak),
	3185: uint16(484),
	3186: uint16(1),
	3187: uint16(anon_sym_RBRACE),
	3188: uint16(2),
	3189: uint16(3),
	3190: uint16(1),
	3191: uint16(sym_linebreak),
	3192: uint16(486),
	3193: uint16(1),
	3194: uint16(sym_number),
	3195: uint16(2),
	3196: uint16(3),
	3197: uint16(1),
	3198: uint16(sym_linebreak),
	3199: uint16(488),
	3200: uint16(1),
	3201: uint16(sym_octal),
	3202: uint16(2),
	3203: uint16(3),
	3204: uint16(1),
	3205: uint16(sym_linebreak),
	3206: uint16(446),
	3207: uint16(1),
	3208: uint16(anon_sym_RBRACE),
	3209: uint16(2),
	3210: uint16(3),
	3211: uint16(1),
	3212: uint16(sym_linebreak),
	3213: uint16(410),
	3214: uint16(1),
	3215: uint16(sym_assignment_op),
	3216: uint16(2),
	3217: uint16(3),
	3218: uint16(1),
	3219: uint16(sym_linebreak),
	3220: uint16(490),
	3221: uint16(1),
	3222: uint16(anon_sym_LBRACE),
	3223: uint16(2),
	3224: uint16(3),
	3225: uint16(1),
	3226: uint16(sym_linebreak),
	3227: uint16(492),
	3228: uint16(1),
	3229: uint16(anon_sym_LBRACE),
	3230: uint16(2),
	3231: uint16(3),
	3232: uint16(1),
	3233: uint16(sym_linebreak),
	3234: uint16(494),
	3235: uint16(1),
	3236: uint16(anon_sym_LBRACE),
	3237: uint16(2),
	3238: uint16(3),
	3239: uint16(1),
	3240: uint16(sym_linebreak),
	3241: uint16(496),
	3242: uint16(1),
	3243: uint16(anon_sym_RBRACE2),
	3244: uint16(2),
	3245: uint16(3),
	3246: uint16(1),
	3247: uint16(sym_linebreak),
	3248: uint16(456),
	3249: uint16(1),
	3250: uint16(sym_match_op),
	3251: uint16(2),
	3252: uint16(3),
	3253: uint16(1),
	3254: uint16(sym_linebreak),
	3255: uint16(498),
	3256: uint16(1),
	3257: uint16(sym_assignment_op),
	3258: uint16(2),
	3259: uint16(3),
	3260: uint16(1),
	3261: uint16(sym_linebreak),
	3262: uint16(422),
	3263: uint16(1),
	3264: uint16(sym_match_op),
	3265: uint16(2),
	3266: uint16(3),
	3267: uint16(1),
	3268: uint16(sym_linebreak),
	3269: uint16(488),
	3270: uint16(1),
	3271: uint16(sym_system_const),
	3272: uint16(2),
	3273: uint16(3),
	3274: uint16(1),
	3275: uint16(sym_linebreak),
	3276: uint16(500),
	3277: uint16(1),
	3278: uint16(anon_sym_RBRACE),
	3279: uint16(2),
	3280: uint16(3),
	3281: uint16(1),
	3282: uint16(sym_linebreak),
	3283: uint16(502),
	3284: uint16(1),
	3285: uint16(anon_sym_LBRACE),
	3286: uint16(2),
	3287: uint16(3),
	3288: uint16(1),
	3289: uint16(sym_linebreak),
	3290: uint16(504),
	3291: uint16(1),
	3292: uint16(anon_sym_LBRACE),
	3293: uint16(2),
	3294: uint16(3),
	3295: uint16(1),
	3296: uint16(sym_linebreak),
	3297: uint16(506),
	3298: uint16(1),
	3299: uint16(anon_sym_RBRACE),
	3300: uint16(2),
	3301: uint16(3),
	3302: uint16(1),
	3303: uint16(sym_linebreak),
	3304: uint16(508),
	3305: uint16(1),
	3306: uint16(anon_sym_LBRACE),
	3307: uint16(2),
	3308: uint16(3),
	3309: uint16(1),
	3310: uint16(sym_linebreak),
	3311: uint16(510),
	3312: uint16(1),
	3313: uint16(anon_sym_RBRACE),
	3314: uint16(2),
	3315: uint16(3),
	3316: uint16(1),
	3317: uint16(sym_linebreak),
	3318: uint16(512),
	3319: uint16(1),
	3320: uint16(anon_sym_RBRACE),
	3321: uint16(2),
	3322: uint16(3),
	3323: uint16(1),
	3324: uint16(sym_linebreak),
	3325: uint16(514),
	3326: uint16(1),
	3327: uint16(anon_sym_RBRACE),
	3328: uint16(2),
	3329: uint16(3),
	3330: uint16(1),
	3331: uint16(sym_linebreak),
	3332: uint16(516),
	3333: uint16(1),
	3334: uint16(anon_sym_RBRACE),
	3335: uint16(2),
	3336: uint16(3),
	3337: uint16(1),
	3338: uint16(sym_linebreak),
	3339: uint16(518),
	3340: uint16(1),
	3341: uint16(anon_sym_RBRACE),
	3342: uint16(2),
	3343: uint16(3),
	3344: uint16(1),
	3345: uint16(sym_linebreak),
	3346: uint16(520),
	3347: uint16(1),
	3348: uint16(anon_sym_RBRACE),
	3349: uint16(2),
	3350: uint16(3),
	3351: uint16(1),
	3352: uint16(sym_linebreak),
	3353: uint16(522),
	3354: uint16(1),
	3355: uint16(anon_sym_LBRACE),
	3356: uint16(2),
	3357: uint16(3),
	3358: uint16(1),
	3359: uint16(sym_linebreak),
	3360: uint16(524),
	3361: uint16(1),
	3362: uint16(anon_sym_LBRACE),
	3363: uint16(2),
	3364: uint16(3),
	3365: uint16(1),
	3366: uint16(sym_linebreak),
	3367: uint16(526),
	3368: uint16(1),
	3369: uint16(anon_sym_LBRACE),
	3370: uint16(2),
	3371: uint16(3),
	3372: uint16(1),
	3373: uint16(sym_linebreak),
	3374: uint16(412),
	3375: uint16(1),
	3376: uint16(sym_match_op),
	3377: uint16(2),
	3378: uint16(3),
	3379: uint16(1),
	3380: uint16(sym_linebreak),
	3381: uint16(528),
	3382: uint16(1),
	3383: uint16(anon_sym_LBRACE),
	3384: uint16(2),
	3385: uint16(3),
	3386: uint16(1),
	3387: uint16(sym_linebreak),
	3388: uint16(530),
	3389: uint16(1),
	3390: uint16(anon_sym_LBRACE),
	3391: uint16(2),
	3392: uint16(3),
	3393: uint16(1),
	3394: uint16(sym_linebreak),
	3395: uint16(532),
	3396: uint16(1),
	3397: uint16(sym_number),
	3398: uint16(2),
	3399: uint16(3),
	3400: uint16(1),
	3401: uint16(sym_linebreak),
	3402: uint16(363),
	3403: uint16(1),
	3404: uint16(anon_sym_DQUOTE2),
	3405: uint16(2),
	3406: uint16(3),
	3407: uint16(1),
	3408: uint16(sym_linebreak),
	3409: uint16(418),
	3410: uint16(1),
	3411: uint16(sym_match_op),
	3412: uint16(2),
	3413: uint16(3),
	3414: uint16(1),
	3415: uint16(sym_linebreak),
	3416: uint16(466),
	3417: uint16(1),
	3418: uint16(anon_sym_RBRACE),
	3419: uint16(2),
	3420: uint16(3),
	3421: uint16(1),
	3422: uint16(sym_linebreak),
	3423: uint16(534),
	3424: uint16(1),
	3425: uint16(aux_sym_rules_token1),
	3426: uint16(2),
	3427: uint16(3),
	3428: uint16(1),
	3429: uint16(sym_linebreak),
	3430: uint16(468),
	3431: uint16(1),
	3432: uint16(anon_sym_RBRACE),
	3433: uint16(2),
	3434: uint16(3),
	3435: uint16(1),
	3436: uint16(sym_linebreak),
	3437: uint16(536),
	3438: uint16(1),
	3439: uint16(sym_seclabel),
	3440: uint16(2),
	3441: uint16(3),
	3442: uint16(1),
	3443: uint16(sym_linebreak),
	3444: uint16(538),
	3445: uint16(1),
	3446: uint16(anon_sym_RBRACE2),
	3447: uint16(2),
	3448: uint16(3),
	3449: uint16(1),
	3450: uint16(sym_linebreak),
	3451: uint16(540),
	3452: uint16(1),
	3453: uint16(anon_sym_RBRACE),
	3454: uint16(2),
	3455: uint16(3),
	3456: uint16(1),
	3457: uint16(sym_linebreak),
	3458: uint16(474),
	3459: uint16(1),
	3460: uint16(anon_sym_RBRACE),
	3461: uint16(2),
	3462: uint16(3),
	3463: uint16(1),
	3464: uint16(sym_linebreak),
	3465: uint16(542),
	3466: uint16(1),
	3468: uint16(2),
	3469: uint16(3),
	3470: uint16(1),
	3471: uint16(sym_linebreak),
	3472: uint16(544),
	3473: uint16(1),
	3474: uint16(anon_sym_RBRACE2),
	3475: uint16(2),
	3476: uint16(3),
	3477: uint16(1),
	3478: uint16(sym_linebreak),
	3479: uint16(546),
	3480: uint16(1),
	3481: uint16(anon_sym_COMMA),
	3482: uint16(2),
	3483: uint16(3),
	3484: uint16(1),
	3485: uint16(sym_linebreak),
	3486: uint16(438),
	3487: uint16(1),
	3488: uint16(anon_sym_RBRACE),
	3489: uint16(2),
	3490: uint16(3),
	3491: uint16(1),
	3492: uint16(sym_linebreak),
	3493: uint16(548),
	3494: uint16(1),
	3495: uint16(sym_number),
	3496: uint16(2),
	3497: uint16(3),
	3498: uint16(1),
	3499: uint16(sym_linebreak),
	3500: uint16(550),
	3501: uint16(1),
	3502: uint16(sym_assignment_op),
	3503: uint16(2),
	3504: uint16(3),
	3505: uint16(1),
	3506: uint16(sym_linebreak),
	3507: uint16(552),
	3508: uint16(1),
	3509: uint16(anon_sym_DQUOTE2),
	3510: uint16(2),
	3511: uint16(3),
	3512: uint16(1),
	3513: uint16(sym_linebreak),
	3514: uint16(554),
	3515: uint16(1),
	3516: uint16(sym_number),
	3517: uint16(2),
	3518: uint16(3),
	3519: uint16(1),
	3520: uint16(sym_linebreak),
	3521: uint16(536),
	3522: uint16(1),
	3523: uint16(sym_run_type),
	3524: uint16(2),
	3525: uint16(3),
	3526: uint16(1),
	3527: uint16(sym_linebreak),
	3528: uint16(556),
	3529: uint16(1),
	3530: uint16(anon_sym_LBRACE),
	3531: uint16(2),
	3532: uint16(3),
	3533: uint16(1),
	3534: uint16(sym_linebreak),
	3535: uint16(558),
	3536: uint16(1),
	3537: uint16(sym_number),
	3538: uint16(2),
	3539: uint16(3),
	3540: uint16(1),
	3541: uint16(sym_linebreak),
	3542: uint16(560),
	3543: uint16(1),
	3544: uint16(anon_sym_LBRACE),
	3545: uint16(2),
	3546: uint16(3),
	3547: uint16(1),
	3548: uint16(sym_linebreak),
	3549: uint16(562),
	3550: uint16(1),
	3551: uint16(anon_sym_LBRACE),
	3552: uint16(2),
	3553: uint16(3),
	3554: uint16(1),
	3555: uint16(sym_linebreak),
	3556: uint16(564),
	3557: uint16(1),
	3558: uint16(anon_sym_LBRACE),
	3559: uint16(2),
	3560: uint16(3),
	3561: uint16(1),
	3562: uint16(sym_linebreak),
	3563: uint16(566),
	3564: uint16(1),
	3565: uint16(anon_sym_LBRACE),
	3566: uint16(2),
	3567: uint16(3),
	3568: uint16(1),
	3569: uint16(sym_linebreak),
	3570: uint16(568),
	3571: uint16(1),
	3572: uint16(anon_sym_LBRACE),
	3573: uint16(2),
	3574: uint16(3),
	3575: uint16(1),
	3576: uint16(sym_linebreak),
	3577: uint16(570),
	3578: uint16(1),
	3579: uint16(anon_sym_LBRACE),
	3580: uint16(2),
	3581: uint16(3),
	3582: uint16(1),
	3583: uint16(sym_linebreak),
	3584: uint16(572),
	3585: uint16(1),
	3586: uint16(sym_import_type),
}

var ts_small_parse_table_map = [165]uint32_t{
	1:   uint32(79),
	2:   uint32(158),
	3:   uint32(237),
	4:   uint32(290),
	5:   uint32(343),
	6:   uint32(415),
	7:   uint32(465),
	8:   uint32(537),
	9:   uint32(609),
	10:  uint32(659),
	11:  uint32(709),
	12:  uint32(759),
	13:  uint32(809),
	14:  uint32(859),
	15:  uint32(909),
	16:  uint32(959),
	17:  uint32(1009),
	18:  uint32(1058),
	19:  uint32(1107),
	20:  uint32(1153),
	21:  uint32(1199),
	22:  uint32(1245),
	23:  uint32(1291),
	24:  uint32(1337),
	25:  uint32(1383),
	26:  uint32(1429),
	27:  uint32(1475),
	28:  uint32(1564),
	29:  uint32(1653),
	30:  uint32(1727),
	31:  uint32(1767),
	32:  uint32(1836),
	33:  uint32(1897),
	34:  uint32(1934),
	35:  uint32(1977),
	36:  uint32(2020),
	37:  uint32(2063),
	38:  uint32(2106),
	39:  uint32(2149),
	40:  uint32(2192),
	41:  uint32(2235),
	42:  uint32(2278),
	43:  uint32(2321),
	44:  uint32(2350),
	45:  uint32(2376),
	46:  uint32(2402),
	47:  uint32(2428),
	48:  uint32(2454),
	49:  uint32(2484),
	50:  uint32(2514),
	51:  uint32(2544),
	52:  uint32(2563),
	53:  uint32(2582),
	54:  uint32(2601),
	55:  uint32(2625),
	56:  uint32(2646),
	57:  uint32(2667),
	58:  uint32(2682),
	59:  uint32(2695),
	60:  uint32(2708),
	61:  uint32(2721),
	62:  uint32(2734),
	63:  uint32(2747),
	64:  uint32(2760),
	65:  uint32(2773),
	66:  uint32(2786),
	67:  uint32(2799),
	68:  uint32(2812),
	69:  uint32(2825),
	70:  uint32(2838),
	71:  uint32(2851),
	72:  uint32(2861),
	73:  uint32(2871),
	74:  uint32(2881),
	75:  uint32(2891),
	76:  uint32(2901),
	77:  uint32(2911),
	78:  uint32(2921),
	79:  uint32(2931),
	80:  uint32(2939),
	81:  uint32(2949),
	82:  uint32(2957),
	83:  uint32(2965),
	84:  uint32(2975),
	85:  uint32(2983),
	86:  uint32(2993),
	87:  uint32(3001),
	88:  uint32(3011),
	89:  uint32(3021),
	90:  uint32(3029),
	91:  uint32(3037),
	92:  uint32(3045),
	93:  uint32(3053),
	94:  uint32(3063),
	95:  uint32(3071),
	96:  uint32(3081),
	97:  uint32(3091),
	98:  uint32(3101),
	99:  uint32(3111),
	100: uint32(3121),
	101: uint32(3131),
	102: uint32(3139),
	103: uint32(3149),
	104: uint32(3159),
	105: uint32(3167),
	106: uint32(3174),
	107: uint32(3181),
	108: uint32(3188),
	109: uint32(3195),
	110: uint32(3202),
	111: uint32(3209),
	112: uint32(3216),
	113: uint32(3223),
	114: uint32(3230),
	115: uint32(3237),
	116: uint32(3244),
	117: uint32(3251),
	118: uint32(3258),
	119: uint32(3265),
	120: uint32(3272),
	121: uint32(3279),
	122: uint32(3286),
	123: uint32(3293),
	124: uint32(3300),
	125: uint32(3307),
	126: uint32(3314),
	127: uint32(3321),
	128: uint32(3328),
	129: uint32(3335),
	130: uint32(3342),
	131: uint32(3349),
	132: uint32(3356),
	133: uint32(3363),
	134: uint32(3370),
	135: uint32(3377),
	136: uint32(3384),
	137: uint32(3391),
	138: uint32(3398),
	139: uint32(3405),
	140: uint32(3412),
	141: uint32(3419),
	142: uint32(3426),
	143: uint32(3433),
	144: uint32(3440),
	145: uint32(3447),
	146: uint32(3454),
	147: uint32(3461),
	148: uint32(3468),
	149: uint32(3475),
	150: uint32(3482),
	151: uint32(3489),
	152: uint32(3496),
	153: uint32(3503),
	154: uint32(3510),
	155: uint32(3517),
	156: uint32(3524),
	157: uint32(3531),
	158: uint32(3538),
	159: uint32(3545),
	160: uint32(3552),
	161: uint32(3559),
	162: uint32(3566),
	163: uint32(3573),
	164: uint32(3580),
}

var ts_parse_actions = [574]TSParseActionEntry{
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
		Fsymbol:     uint16(sym_rules),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(30)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(118)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(118)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(87)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(115)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(116)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(124)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(126)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(134)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(95)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(97)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(74)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(113)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(114)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(73)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(119)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(133)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(143)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(84)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(8)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(18)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(13)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(15)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(164)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(165)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(5)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(12)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(162)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(163)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(6)),
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
		Fcount: uint8(1),
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
		Fextra: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(91)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__sub_c_content),
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
		Fsymbol:      uint16(aux_sym__sub_c_content),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(8)),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym__sub_c_content),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(18)),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym__sub_c_content),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(13)),
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
		Fsymbol:      uint16(aux_sym__sub_c_content),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(15)),
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
		Fsymbol:      uint16(aux_sym__sub_c_content),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(164)),
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
		Fsymbol:      uint16(aux_sym__sub_c_content),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(165)),
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
		Fsymbol:      uint16(aux_sym__sub_c_content),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(5)),
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
		Fsymbol:      uint16(aux_sym__sub_c_content),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(12)),
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
		Fsymbol:      uint16(aux_sym__sub_c_content),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(162)),
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
		Fsymbol:      uint16(aux_sym__sub_c_content),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(163)),
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
		Fsymbol:      uint16(aux_sym__sub_c_content),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(6)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(159)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_fmt_sub),
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
		Fcount: uint8(1),
	}})),
	111: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_fmt_sub),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(156)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_var_sub),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_var_sub),
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
		Fsymbol:      uint16(aux_sym__sub_content),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym__sub_content),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(21)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	124: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__sub_content),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(26)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	127: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__sub_content),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(24)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym__sub_content),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(160)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym__sub_content),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(161)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	136: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__sub_content),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(20)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym__sub_content),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(22)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	142: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__sub_content),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(135)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym__sub_content),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(158)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	148: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__sub_content),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(19)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym__sub_c_content),
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
		Fcount: uint8(1),
	}})),
	153: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym__sub_c_content),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(83)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(21)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(26)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(24)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(160)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(161)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(20)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(22)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(135)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(158)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(19)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(103)),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_var_sub),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_var_sub),
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
		Fsymbol:      uint16(sym_c_escape),
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
		Fsymbol:      uint16(sym_c_escape),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_var_sub),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_var_sub),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_fmt_sub),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_fmt_sub),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_fmt_sub),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_fmt_sub),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_pattern),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_pattern),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(139)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(153)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym__sub_content),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym__sub_content),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_rules_repeat1),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_rules_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(29)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	216: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_rules_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(118)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	219: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_rules_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(118)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_rules_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(87)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	225: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_rules_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(115)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_rules_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(116)),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_rules_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(124)),
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
		Fsymbol:      uint16(aux_sym_rules_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(126)),
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
		Fsymbol:      uint16(aux_sym_rules_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(134)),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_rules_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(95)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	243: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_rules_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(97)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_rules_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(74)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	249: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_rules_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(113)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	252: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_rules_repeat1),
	})))),
	253: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_rules_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(73)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_rules_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(119)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_rules_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(133)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	264: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_rules_repeat1),
	})))),
	265: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(143)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_rules),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fsymbol:      uint16(aux_sym_rules_repeat1),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_rule_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(118)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	276: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_rule_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(118)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	279: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_rule_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(136)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	282: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_rule_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(116)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	285: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_rule_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(116)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	288: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_rule_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(137)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_rule_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(138)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_rule_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(134)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	297: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_rule_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(97)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	300: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_rule_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(74)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_rule_repeat1),
	})))),
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
		Fcount: uint8(1),
	}})),
	305: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_rule_repeat1),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(40)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(48)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(49)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(107)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(123)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(45)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	321: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	322: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(37)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_attribute_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(48)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	327: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	328: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(49)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_attribute_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(107)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	333: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(123)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_attribute_repeat1),
	})))),
	337: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	339: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(110)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(88)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(54)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(55)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(53)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	353: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__c_content),
	})))),
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
		Fsymbol:      uint16(aux_sym__c_content),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(54)),
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
		Fsymbol:      uint16(aux_sym__c_content),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(55)),
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
		Fsymbol:      uint16(aux_sym__c_content),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(53)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	364: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	366: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym__c_content),
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
		Fcount: uint8(1),
	}})),
	368: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym__c_content),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(96)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(58)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(59)),
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
		Fsymbol:      uint16(aux_sym_content_repeat1),
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
		Fcount: uint8(2),
	}})),
	378: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_content_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(57)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	381: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_content_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(59)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_content),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_rule),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_rule_repeat2),
	})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	394: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_rule_repeat2),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(33)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_rule),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(10)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(108)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(56)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(155)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_rule),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	409: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(157)),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(65)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(65)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(129)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(64)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(69)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(71)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(63)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(69)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(128)),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_value),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	433: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym__sub_value),
		Fproduction_id: uint16(2),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__sub_value),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(117)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	439: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_match),
		Fproduction_id: uint16(1),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_value),
		Fproduction_id: uint16(3),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(122)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	451: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym__sub_value),
		Fproduction_id: uint16(3),
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
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_match),
		Fproduction_id: uint16(1),
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
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_assignment),
		Fproduction_id: uint16(1),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	457: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_value),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	465: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(146)),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	467: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(147)),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	473: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(150)),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	475: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__sub_value),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	479: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_assignment),
		Fproduction_id: uint16(1),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	485: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(120)),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	489: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(127)),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(145)),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	493: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	497: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(72)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(28)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(75)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(101)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(76)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(105)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(141)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_kernel_param),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_env_var),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(77)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(154)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(166)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(121)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(38)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(79)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(80)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(90)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(32)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(131)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	539: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	543: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f uint8_t
		_ [7]byte
	}{f: uint8(TSParseActionTypeAccept)})),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	547: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(98)),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(63)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	553: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	555: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(99)),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(100)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(42)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(102)),
	}})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(43)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(132)),
	}})))),
}

func tree_sitter_udev(tls *libc.TLS) (r uintptr) {
	return uintptr(unsafe.Pointer(&language))
}

var language = TSLanguage{
	Fabi_version:               uint32(LANGUAGE_VERSION),
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
	Fprimary_state_ids:         uintptr(unsafe.Pointer(&ts_primary_state_ids)),
	Fname:                      __ccgo_ts + 848,
	Fmetadata: TSLanguageMetadata{
		Fminor_version: uint8(3),
	},
}

func init() {
	p := unsafe.Pointer(&language)
	*(*uintptr)(unsafe.Add(p, 92)) = __ccgo_fp(ts_lex)
}

func __ccgo_fp(f interface{}) uintptr {
	type iface [2]uintptr
	return (*iface)(unsafe.Pointer(&f))[1]
}

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = "end\x00rules_token1\x00,\x00ACTION\x00DEVPATH\x00KERNEL\x00KERNELS\x00NAME\x00SYMLINK\x00SUBSYSTEM\x00SUBSYSTEMS\x00DRIVER\x00DRIVERS\x00ATTR\x00{\x00}\x00ATTRS\x00SYSCTL\x00ENV\x00CONST\x00TAG\x00TAGS\x00TEST\x00PROGRAM\x00RESULT\x00OWNER\x00GROUP\x00MODE\x00SECLABEL\x00RUN\x00LABEL\x00GOTO\x00IMPORT\x00OPTIONS\x00system_const\x00run_type\x00import_type\x00attribute_token1\x00env_var_token1\x00seclabel\x00octal\x00number\x00match_op\x00assignment_op\x00\"\x00e\x00content_token1\x00\\\"\x00*\x00?\x00|\x00pattern_token1\x00c_escape_token1\x00c_escape_token2\x00c_escape_token3\x00c_escape_token4\x00c_escape_token5\x00%k\x00%n\x00%p\x00%b\x00%s\x00%E\x00%M\x00%m\x00%c\x00+\x00%P\x00%r\x00%S\x00%N\x00%%\x00$kernel\x00$number\x00$devpath\x00$id\x00$driver\x00$attr\x00$env\x00$major\x00$minor\x00$result\x00$parent\x00$name\x00$links\x00$root\x00$sys\x00$devnode\x00$$\x00linebreak\x00comment\x00rules\x00rule\x00match\x00assignment\x00attribute\x00env_var\x00kernel_param\x00value\x00content\x00_sub_content\x00_c_content\x00_sub_c_content\x00pattern\x00c_escape\x00fmt_sub\x00var_sub\x00rules_repeat1\x00rule_repeat1\x00rule_repeat2\x00attribute_repeat1\x00content_repeat1\x00key\x00udev\x00"
