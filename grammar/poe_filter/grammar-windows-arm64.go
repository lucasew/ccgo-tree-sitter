// Code generated for windows/arm64 by 'ccgo -ignore-unsupported-alignment -ignore-unsupported-atomic-sizes -ignore-vector-functions --winapi-no-errno -ignore-link-errors -std=gnu11 -O0 -D__extension__= -D__forceinline=static inline -D__attribute__(...)= -D__declspec(x)= -D__cdecl= -D__stdcall= -D__fastcall= -D__thiscall= -D_cdecl= -D__restrict= -D__restrict__= -D__MINGW_EXTENSION= -D_X86INTRIN_H_INCLUDED -D_X86GPRINTRIN_H_INCLUDED -D_IMMINTRIN_H_INCLUDED -D_MMINTRIN_H_INCLUDED -D_XMMINTRIN_H_INCLUDED -D_EMMINTRIN_H_INCLUDED -D_PMMINTRIN_H_INCLUDED -D_MM3DNOW_H_INCLUDED -Dfunc=func_token -Dinterface=interface_token -Dselect=select_token -Dchan=chan_token -Dgo=go_token -Dmap=map_token -Dpackage=package_token -Dtype=type_token -Dvar=var_token -Dimport=import_token -Ddefer=defer_token -Dfallthrough=fallthrough_token -Drange=range_token -I C:\a\ccgo-tree-sitter\ccgo-tree-sitter\third-party\tree-sitter-poe-filter\src -I C:\a\ccgo-tree-sitter\ccgo-tree-sitter\third-party\tree-sitter-poe-filter -I C:\a\ccgo-tree-sitter\ccgo-tree-sitter\third-party -I C:\Users\runneradmin\.cache\workspaced\sources\github\7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2\lib\include -I C:\Users\runneradmin\.cache\workspaced\sources\github\7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2\lib\src C:\a\ccgo-tree-sitter\ccgo-tree-sitter\third-party\tree-sitter-poe-filter\src\parser.c -o grammar.go', DO NOT EDIT.

//go:build windows && arm64

package grammar_poe_filter

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
const FIELD_COUNT = 13
const INTMAX_MAX = "__INTMAX_MAX__"
const INTPTR_MAX = "__INTPTR_MAX__"
const INT_FAST16_MAX = "__INT_LEAST16_MAX"
const INT_FAST16_MIN = "__INT_LEAST16_MIN"
const INT_FAST32_MAX = "__INT_LEAST32_MAX"
const INT_FAST32_MIN = "__INT_LEAST32_MIN"
const INT_FAST64_MAX = "__INT_LEAST64_MAX"
const INT_FAST64_MIN = "__INT_LEAST64_MIN"
const INT_FAST8_MAX = "__INT_LEAST8_MAX"
const INT_FAST8_MIN = "__INT_LEAST8_MIN"
const INT_LEAST16_MAX = "__INT_LEAST16_MAX"
const INT_LEAST16_MIN = "__INT_LEAST16_MIN"
const INT_LEAST32_MAX = "__INT_LEAST32_MAX"
const INT_LEAST32_MIN = "__INT_LEAST32_MIN"
const INT_LEAST64_MAX = "__INT_LEAST64_MAX"
const INT_LEAST64_MIN = "__INT_LEAST64_MIN"
const INT_LEAST8_MAX = "__INT_LEAST8_MAX"
const INT_LEAST8_MIN = "__INT_LEAST8_MIN"
const LANGUAGE_VERSION = 14
const LARGE_STATE_COUNT = 33
const MAX_ALIAS_SEQUENCE_LENGTH = 10
const MAX_RESERVED_WORD_SET_SIZE = 0
const MB_LEN_MAX = 1
const MINGW_HAS_DDK_H = 1
const MINGW_HAS_SECURE_API = 1
const PRODUCTION_ID_COUNT = 27
const PTRDIFF_MAX = "__PTRDIFF_MAX__"
const RAND_MAX = 0x7fff
const SIZE_MAX = "__SIZE_MAX__"
const STATE_COUNT = 195
const SUPERTYPE_COUNT = 0
const SYMBOL_COUNT = 167
const TOKEN_COUNT = 135
const TREE_SITTER_SERIALIZATION_BUFFER_SIZE = 1024
const UINTMAX_MAX = "__UINTMAX_MAX__"
const UINTPTR_MAX = "__UINTPTR_MAX__"
const UINT_FAST16_MAX = "__UINT_LEAST16_MAX"
const UINT_FAST32_MAX = "__UINT_LEAST32_MAX"
const UINT_FAST64_MAX = "__UINT_LEAST64_MAX"
const UINT_FAST8_MAX = "__UINT_LEAST8_MAX"
const UINT_LEAST16_MAX = "__UINT_LEAST16_MAX"
const UINT_LEAST32_MAX = "__UINT_LEAST32_MAX"
const UINT_LEAST64_MAX = "__UINT_LEAST64_MAX"
const UINT_LEAST8_MAX = "__UINT_LEAST8_MAX"
const USE___UUIDOF = 0
const WCHAR_MAX = "__WCHAR_MAX__"
const WIN32 = 1
const WIN64 = 1
const WINNT = 1
const _ALLOCA_S_HEAP_MARKER = 56797
const _ALLOCA_S_MARKER_SIZE = 16
const _ALLOCA_S_STACK_MARKER = 0xCCCC
const _ALLOCA_S_THRESHOLD = 1024
const _ANONYMOUS_STRUCT = "__MINGW_EXTENSION"
const _ANONYMOUS_UNION = "__MINGW_EXTENSION"
const _ARGMAX = 100
const _ARM64_ = 1
const _CALL_REPORTFAULT = 0x2
const _CRTIMP2 = "_CRTIMP"
const _CRTIMP_ALTERNATIVE = "_CRTIMP"
const _CRTIMP_NOIA64 = "_CRTIMP"
const _CRTIMP_PURE = "_CRTIMP"
const _EMMINTRIN_H_INCLUDED = 1
const _FREEENTRY = 0
const _HEAP_MAXREQ = 0xFFFFFFFFFFFFFFE0
const _IMMINTRIN_H_INCLUDED = 1
const _MAX_DIR = 256
const _MAX_DRIVE = 3
const _MAX_ENV = 32767
const _MAX_EXT = 256
const _MAX_FNAME = 256
const _MAX_PATH = 260
const _MAX_WAIT_MALLOC_CRT = 60000
const _MCRTIMP = "_CRTIMP"
const _MM3DNOW_H_INCLUDED = 1
const _MMINTRIN_H_INCLUDED = 1
const _MRTIMP2 = "_CRTIMP"
const _M_ARM64 = 1
const _OUT_TO_DEFAULT = 0
const _OUT_TO_MSGBOX = 2
const _OUT_TO_STDERR = 1
const _PMMINTRIN_H_INCLUDED = 1
const _REPORT_ERRMODE = 3
const _SECURECRT_FILL_BUFFER_PATTERN = 0xFD
const _USEDENTRY = 1
const _WIN32 = 1
const _WIN32_WINNT = 0x601
const _WIN64 = 1
const _WRITE_ABORT_MSG = 0x1
const _X86GPRINTRIN_H_INCLUDED = 1
const _X86INTRIN_H_INCLUDED = 1
const _XMMINTRIN_H_INCLUDED = 1
const __AARCH64EL__ = 1
const __AARCH64_CMODEL_SMALL__ = 1
const __ARM_64BIT_STATE = 1
const __ARM_ACLE = 202420
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
const __ARM_NEON_SVE_BRIDGE = 1
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
const __BITINT_MAXWIDTH__ = 128
const __BOOL_WIDTH__ = 1
const __BYTE_ORDER__ = "__ORDER_LITTLE_ENDIAN__"
const __C89_NAMELESS = "__MINGW_EXTENSION"
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
const __CRTDECL = "__cdecl"
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
const __DBL_NORM_MAX__ = 1.7976931348623157e+308
const __DECIMAL_DIG__ = "__LDBL_DECIMAL_DIG__"
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
const __FLT16_NORM_MAX__ = 6.5504e+4
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
const __FLT_NORM_MAX__ = 3.40282347e+38
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
const __FUNCTION_MULTI_VERSIONING_SUPPORT_LEVEL = 202430
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
const __GCC_CONSTRUCTIVE_SIZE = 64
const __GCC_DESTRUCTIVE_SIZE = 256
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_1 = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_16 = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_2 = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_4 = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_8 = 1
const __GNUC_MINOR__ = 2
const __GNUC_PATCHLEVEL__ = 1
const __GNUC_STDC_INLINE__ = 1
const __GNUC__ = 4
const __GNU_EXTENSION = "__MINGW_EXTENSION"
const __GOT_SECURE_LIB__ = "__STDC_SECURE_LIB__"
const __GXX_ABI_VERSION = 1002
const __GXX_TYPEINFO_EQUALITY_INLINE = 0
const __HAVE_FUNCTION_MULTI_VERSIONING = 1
const __INT16_FMTd__ = "hd"
const __INT16_FMTi__ = "hi"
const __INT16_MAX__ = 32767
const __INT16_TYPE__ = "short"
const __INT32_FMTd__ = "d"
const __INT32_FMTi__ = "i"
const __INT32_MAX__ = 2147483647
const __INT32_TYPE__ = "int"
const __INT64_C_SUFFIX__ = "LL"
const __INT64_FMTd__ = "lld"
const __INT64_FMTi__ = "lli"
const __INT64_MAX__ = 9223372036854775807
const __INT8_FMTd__ = "hhd"
const __INT8_FMTi__ = "hhi"
const __INT8_MAX__ = 127
const __INTMAX_C_SUFFIX__ = "LL"
const __INTMAX_FMTd__ = "lld"
const __INTMAX_FMTi__ = "lli"
const __INTMAX_MAX__ = 9223372036854775807
const __INTMAX_WIDTH__ = 64
const __INTPTR_FMTd__ = "lld"
const __INTPTR_FMTi__ = "lli"
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
const __INT_FAST64_FMTd__ = "lld"
const __INT_FAST64_FMTi__ = "lli"
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
const __INT_LEAST64_FMTd__ = "lld"
const __INT_LEAST64_FMTi__ = "lli"
const __INT_LEAST64_MAX = "INT64_MAX"
const __INT_LEAST64_MAX__ = 9223372036854775807
const __INT_LEAST64_MIN = "INT64_MIN"
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
const __LDBL_NORM_MAX__ = 1.7976931348623157e+308
const __LITTLE_ENDIAN__ = 1
const __LLONG_WIDTH__ = 64
const __LONG32 = "long"
const __LONG_LONG_MAX__ = 9223372036854775807
const __LONG_MAX__ = 2147483647
const __LONG_WIDTH__ = 32
const __MEMORY_SCOPE_DEVICE = 1
const __MEMORY_SCOPE_SINGLE = 4
const __MEMORY_SCOPE_SYSTEM = 0
const __MEMORY_SCOPE_WRKGRP = 2
const __MEMORY_SCOPE_WVFRNT = 3
const __MINGW32_MAJOR_VERSION = 3
const __MINGW32_MINOR_VERSION = 11
const __MINGW32__ = 1
const __MINGW64_VERSION_BUGFIX = 0
const __MINGW64_VERSION_MAJOR = 14
const __MINGW64_VERSION_MINOR = 0
const __MINGW64_VERSION_RC = 0
const __MINGW64_VERSION_STATE = "alpha"
const __MINGW64__ = 1
const __MINGW_DEBUGBREAK_IMPL = 1
const __MINGW_FASTFAIL_IMPL = 1
const __MINGW_FORTIFY_LEVEL = 0
const __MINGW_FORTIFY_VA_ARG = 0
const __MINGW_HAVE_ANSI_C99_PRINTF = 1
const __MINGW_HAVE_ANSI_C99_SCANF = 1
const __MINGW_HAVE_WIDE_C99_PRINTF = 1
const __MINGW_HAVE_WIDE_C99_SCANF = 1
const __MINGW_MSVC2005_DEPREC_STR = "This POSIX function is deprecated beginning in Visual C++ 2005, use _CRT_NONSTDC_NO_DEPRECATE to disable deprecation"
const __MINGW_PREFETCH_IMPL = 1
const __MINGW_SEC_WARN_STR = "This function or variable may be unsafe, use _CRT_SECURE_NO_WARNINGS to disable deprecation"
const __MINGW_USE_UNDERSCORE_PREFIX = 0
const __MSVCRT_VERSION__ = 0xE00
const __MSVCRT__ = 1
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
const __POINTER_WIDTH__ = 64
const __PRAGMA_REDEFINE_EXTNAME = 1
const __PRETTY_FUNCTION__ = "__func__"
const __PTRDIFF_FMTd__ = "lld"
const __PTRDIFF_FMTi__ = "lli"
const __PTRDIFF_MAX__ = 9223372036854775807
const __PTRDIFF_WIDTH__ = 64
const __SCHAR_MAX__ = 127
const __SEH__ = 1
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
const __SIZEOF_LONG__ = 4
const __SIZEOF_POINTER__ = 8
const __SIZEOF_PTRDIFF_T__ = 8
const __SIZEOF_SHORT__ = 2
const __SIZEOF_SIZE_T__ = 8
const __SIZEOF_WCHAR_T__ = 2
const __SIZEOF_WINT_T__ = 2
const __SIZE_FMTX__ = "llX"
const __SIZE_FMTo__ = "llo"
const __SIZE_FMTu__ = "llu"
const __SIZE_FMTx__ = "llx"
const __SIZE_MAX__ = "18446744073709551615U"
const __SIZE_WIDTH__ = 64
const __STDC_EMBED_EMPTY__ = 2
const __STDC_EMBED_FOUND__ = 1
const __STDC_EMBED_NOT_FOUND__ = 0
const __STDC_HOSTED__ = 1
const __STDC_SECURE_LIB__ = 200411
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
const __UINT64_C_SUFFIX__ = "ULL"
const __UINT64_FMTX__ = "llX"
const __UINT64_FMTo__ = "llo"
const __UINT64_FMTu__ = "llu"
const __UINT64_FMTx__ = "llx"
const __UINT64_MAX__ = "18446744073709551615U"
const __UINT8_FMTX__ = "hhX"
const __UINT8_FMTo__ = "hho"
const __UINT8_FMTu__ = "hhu"
const __UINT8_FMTx__ = "hhx"
const __UINT8_MAX__ = 255
const __UINTMAX_C_SUFFIX__ = "ULL"
const __UINTMAX_FMTX__ = "llX"
const __UINTMAX_FMTo__ = "llo"
const __UINTMAX_FMTu__ = "llu"
const __UINTMAX_FMTx__ = "llx"
const __UINTMAX_MAX__ = "18446744073709551615U"
const __UINTMAX_WIDTH__ = 64
const __UINTPTR_FMTX__ = "llX"
const __UINTPTR_FMTo__ = "llo"
const __UINTPTR_FMTu__ = "llu"
const __UINTPTR_FMTx__ = "llx"
const __UINTPTR_MAX__ = "18446744073709551615U"
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
const __UINT_FAST64_FMTX__ = "llX"
const __UINT_FAST64_FMTo__ = "llo"
const __UINT_FAST64_FMTu__ = "llu"
const __UINT_FAST64_FMTx__ = "llx"
const __UINT_FAST64_MAX__ = "18446744073709551615U"
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
const __UINT_LEAST64_FMTX__ = "llX"
const __UINT_LEAST64_FMTo__ = "llo"
const __UINT_LEAST64_FMTu__ = "llu"
const __UINT_LEAST64_FMTx__ = "llx"
const __UINT_LEAST64_MAX = "UINT64_MAX"
const __UINT_LEAST64_MAX__ = "18446744073709551615U"
const __UINT_LEAST8_FMTX__ = "hhX"
const __UINT_LEAST8_FMTo__ = "hho"
const __UINT_LEAST8_FMTu__ = "hhu"
const __UINT_LEAST8_FMTx__ = "hhx"
const __UINT_LEAST8_MAX__ = 255
const __USE_MINGW_ANSI_STDIO = 0
const __VERSION__ = "Clang 21.1.8 (https://github.com/llvm/llvm-project.git 2078da43e25a4623cab2d0d60decddf709aaea28)"
const __WCHAR_MAX__ = 65535
const __WCHAR_UNSIGNED__ = 1
const __WCHAR_WIDTH__ = 16
const __WIN32 = 1
const __WIN32__ = 1
const __WIN64 = 1
const __WIN64__ = 1
const __WINNT = 1
const __WINNT__ = 1
const __WINT_MAX__ = 65535
const __WINT_UNSIGNED__ = 1
const __WINT_WIDTH__ = 16
const __aarch64__ = 1
const __bool_true_false_are_defined = 1
const __clang__ = 1
const __clang_literal_encoding__ = "UTF-8"
const __clang_major__ = 21
const __clang_minor__ = 1
const __clang_patchlevel__ = 8
const __clang_version__ = "21.1.8 (https://github.com/llvm/llvm-project.git 2078da43e25a4623cab2d0d60decddf709aaea28)"
const __clang_wide_literal_encoding__ = "UTF-16"
const __int16 = "short"
const __int32 = "int"
const __int8 = "char"
const __llvm__ = 1
const __mingw_bos_ovr = "__mingw_ovr"
const __pic__ = 2
const _inline = "__inline"
const bool1 = "_Bool"
const chan1 = "chan_token"
const defer1 = "defer_token"
const environ1 = "_environ"
const fallthrough1 = "fallthrough_token"
const false1 = 0
const func1 = "func_token"
const go1 = "go_token"
const import1 = "import_token"
const interface1 = "interface_token"
const map1 = "map_token"
const onexit_t = "_onexit_t"
const package1 = "package_token"
const range1 = "range_token"
const select2 = "select_token"
const sys_errlist = "_sys_errlist"
const sys_nerr = "_sys_nerr"
const true1 = 1
const ts_builtin_sym_end = 0
const type1 = "type_token"
const var1 = "var_token"

type __builtin_va_list = uintptr

type __predefined_size_t = uint64

type __predefined_wchar_t = uint16

type __predefined_ptrdiff_t = int64

type int64_t = int64

type uint64_t = uint64

type int_least64_t = int64

type uint_least64_t = uint64

type int_fast64_t = int64

type uint_fast64_t = uint64

type int32_t = int32

type uint32_t = uint32

type int_least32_t = int32

type uint_least32_t = uint32

type int_fast32_t = int32

type uint_fast32_t = uint32

type int16_t = int16

type uint16_t = uint16

type int_least16_t = int16

type uint_least16_t = uint16

type int_fast16_t = int16

type uint_fast16_t = uint16

type int8_t = int8

type uint8_t = uint8

type int_least8_t = int8

type uint_least8_t = uint8

type int_fast8_t = int8

type uint_fast8_t = uint8

type intptr_t = int64

type uintptr_t = uint64

type intmax_t = int64

type uintmax_t = uint64

type __gnuc_va_list = uintptr

type va_list = uintptr

type size_t = uint64

type ssize_t = int64

type rsize_t = uint64

type ptrdiff_t = int64

type wchar_t = uint16

type wint_t = uint16

type wctype_t = uint16

type errno_t = int32

type __time32_t = int32

type __time64_t = int64

type time_t = int64

type threadlocaleinfostruct = struct {
	F_locale_pctype      uintptr
	F_locale_mb_cur_max  int32
	F_locale_lc_codepage uint32
}

type pthreadlocinfo = uintptr

type pthreadmbcinfo = uintptr

type _locale_tstruct = struct {
	Flocinfo pthreadlocinfo
	Fmbcinfo pthreadmbcinfo
}

type localeinfo_struct = _locale_tstruct

type _locale_t = uintptr

type LC_ID = struct {
	FwLanguage uint16
	FwCountry  uint16
	FwCodePage uint16
}

type tagLC_ID = LC_ID

type LPLC_ID = uintptr

type threadlocinfo = struct {
	F_locale_pctype      uintptr
	F_locale_mb_cur_max  int32
	F_locale_lc_codepage uint32
}

type _onexit_t = uintptr

type div_t = struct {
	Fquot int32
	Frem  int32
}

type _div_t = div_t

type ldiv_t = struct {
	Fquot int32
	Frem  int32
}

type _ldiv_t = ldiv_t

type _LDOUBLE = struct {
	Fld [10]uint8
}

type _CRT_DOUBLE = struct {
	Fx float64
}

type _CRT_FLOAT = struct {
	Ff float32
}

type _LONGDOUBLE = struct {
	Fx float64
}

type _LDBL12 = struct {
	Fld12 [12]uint8
}

type _purecall_handler = uintptr

type _invalid_parameter_handler = uintptr

type lldiv_t = struct {
	Fquot int64
	Frem  int64
}

type _HEAPINFO = struct {
	F_pentry  uintptr
	F_size    size_t
	F_useflag int32
}

type _heapinfo = _HEAPINFO

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

const anon_sym_Show = 1
const anon_sym_Hide = 2
const anon_sym_Minimal = 3
const anon_sym_Import = 4
const anon_sym_Optional = 5
const anon_sym_AlternateQuality = 6
const anon_sym_AnyEnchantment = 7
const anon_sym_ArchnemesisMod = 8
const anon_sym_AreaLevel = 9
const anon_sym_BaseArmour = 10
const anon_sym_BaseDefencePercentile = 11
const anon_sym_BaseEnergyShield = 12
const anon_sym_BaseEvasion = 13
const anon_sym_BaseType = 14
const anon_sym_BaseWard = 15
const anon_sym_BlightedMap = 16
const anon_sym_Class = 17
const anon_sym_Corrupted = 18
const anon_sym_CorruptedMods = 19
const anon_sym_DropLevel = 20
const anon_sym_ElderItem = 21
const anon_sym_ElderMap = 22
const anon_sym_EnchantmentPassiveNode = 23
const anon_sym_EnchantmentPassiveNum = 24
const anon_sym_FracturedItem = 25
const anon_sym_GemLevel = 26
const anon_sym_GemQualityType = 27
const anon_sym_HasCruciblePassiveTree = 28
const anon_sym_HasEaterOfWorldsImplicit = 29
const anon_sym_HasEnchantment = 30
const anon_sym_HasExplicitMod = 31
const anon_sym_HasImplicitMod = 32
const anon_sym_HasInfluence = 33
const anon_sym_HasSearingExarchImplicit = 34
const anon_sym_Height = 35
const anon_sym_Identified = 36
const anon_sym_ItemLevel = 37
const anon_sym_LinkedSockets = 38
const anon_sym_MapTier = 39
const anon_sym_MemoryStrands = 40
const anon_sym_Mirrored = 41
const anon_sym_Quality = 42
const anon_sym_Rarity = 43
const anon_sym_Replica = 44
const anon_sym_Scourged = 45
const anon_sym_ShapedMap = 46
const anon_sym_ShaperItem = 47
const anon_sym_SocketGroup = 48
const anon_sym_Sockets = 49
const anon_sym_StackSize = 50
const anon_sym_SynthesisedItem = 51
const anon_sym_TransfiguredGem = 52
const anon_sym_UberBlightedMap = 53
const anon_sym_UnidentifiedItemTier = 54
const anon_sym_WaystoneTier = 55
const anon_sym_Width = 56
const anon_sym_ZanaMemory = 57
const anon_sym_PlayAlertSound = 58
const anon_sym_None = 59
const anon_sym_PlayAlertSoundPositional = 60
const anon_sym_CustomAlertSound = 61
const aux_sym_action_token1 = 62
const anon_sym_CustomAlertSoundOptional = 63
const anon_sym_DisableDropSound = 64
const anon_sym_EnableDropSound = 65
const anon_sym_DisableDropSoundIfAlertSound = 66
const anon_sym_EnableDropSoundIfAlertSound = 67
const anon_sym_MinimapIcon = 68
const anon_sym_DASH1 = 69
const anon_sym_PlayEffect = 70
const anon_sym_Temp = 71
const anon_sym_SetBackgroundColor = 72
const anon_sym_SetBorderColor = 73
const anon_sym_SetFontSize = 74
const anon_sym_SetTextColor = 75
const anon_sym_Continue = 76
const aux_sym__equal_operator_token1 = 77
const aux_sym__range_operator_token1 = 78
const sym_boolean = 79
const anon_sym_DQUOTE = 80
const aux_sym_quality_token1 = 81
const anon_sym_Superior = 82
const anon_sym_Divergent = 83
const anon_sym_Anomalous = 84
const anon_sym_Phantasmal = 85
const aux_sym_rarity_token1 = 86
const anon_sym_Normal = 87
const anon_sym_Magic = 88
const anon_sym_Rare = 89
const anon_sym_Unique = 90
const aux_sym_influence_token1 = 91
const anon_sym_Shaper = 92
const anon_sym_Elder = 93
const anon_sym_Crusader = 94
const anon_sym_Hunter = 95
const anon_sym_Redeemer = 96
const anon_sym_Warlord = 97
const aux_sym_sockets_token1 = 98
const aux_sym_sockets_token2 = 99
const anon_sym_Red = 100
const anon_sym_Green = 101
const anon_sym_Blue = 102
const anon_sym_Brown = 103
const anon_sym_White = 104
const anon_sym_Yellow = 105
const anon_sym_Cyan = 106
const anon_sym_Grey = 107
const anon_sym_Orange = 108
const anon_sym_Pink = 109
const anon_sym_Purple = 110
const anon_sym_Circle = 111
const anon_sym_Diamond = 112
const anon_sym_Hexagon = 113
const anon_sym_Square = 114
const anon_sym_Star = 115
const anon_sym_Triangle = 116
const anon_sym_Cross = 117
const anon_sym_Moon = 118
const anon_sym_Raindrop = 119
const anon_sym_Kite = 120
const anon_sym_Pentagon = 121
const anon_sym_UpsideDownHouse = 122
const aux_sym_string_token1 = 123
const aux_sym_string_token2 = 124
const aux_sym_file_token1 = 125
const sym_number = 126
const aux_sym__id_token1 = 127
const aux_sym__volume_token1 = 128
const aux_sym__color_token1 = 129
const aux_sym__icon_size_token1 = 130
const aux_sym__font_size_token1 = 131
const sym_comment = 132
const sym__space = 133
const sym__eol = 134
const sym_filter = 135
const sym_block = 136
const sym_import = 137
const sym_condition = 138
const sym_action = 139
const sym_continue = 140
const sym__equal_operator = 141
const sym__range_operator = 142
const sym_quality = 143
const sym_rarity = 144
const sym_influence = 145
const sym_sockets = 146
const sym_colour = 147
const sym_shape = 148
const sym_string = 149
const sym_file = 150
const sym__quantity = 151
const sym__id = 152
const sym__volume = 153
const sym__color = 154
const sym__icon_size = 155
const sym__font_size = 156
const aux_sym_filter_repeat1 = 157
const aux_sym_block_repeat1 = 158
const aux_sym_condition_repeat1 = 159
const aux_sym_condition_repeat2 = 160
const aux_sym_condition_repeat3 = 161
const aux_sym_condition_repeat4 = 162
const aux_sym_condition_repeat5 = 163
const aux_sym_condition_repeat6 = 164
const aux_sym_condition_repeat7 = 165
const aux_sym_condition_repeat8 = 166

var ts_symbol_names = [167]uintptr{
	0:   __ccgo_ts,
	1:   __ccgo_ts + 4,
	2:   __ccgo_ts + 9,
	3:   __ccgo_ts + 14,
	4:   __ccgo_ts + 22,
	5:   __ccgo_ts + 29,
	6:   __ccgo_ts + 38,
	7:   __ccgo_ts + 38,
	8:   __ccgo_ts + 38,
	9:   __ccgo_ts + 38,
	10:  __ccgo_ts + 38,
	11:  __ccgo_ts + 38,
	12:  __ccgo_ts + 38,
	13:  __ccgo_ts + 38,
	14:  __ccgo_ts + 38,
	15:  __ccgo_ts + 38,
	16:  __ccgo_ts + 38,
	17:  __ccgo_ts + 38,
	18:  __ccgo_ts + 38,
	19:  __ccgo_ts + 38,
	20:  __ccgo_ts + 38,
	21:  __ccgo_ts + 38,
	22:  __ccgo_ts + 38,
	23:  __ccgo_ts + 38,
	24:  __ccgo_ts + 38,
	25:  __ccgo_ts + 38,
	26:  __ccgo_ts + 38,
	27:  __ccgo_ts + 38,
	28:  __ccgo_ts + 38,
	29:  __ccgo_ts + 38,
	30:  __ccgo_ts + 38,
	31:  __ccgo_ts + 38,
	32:  __ccgo_ts + 38,
	33:  __ccgo_ts + 38,
	34:  __ccgo_ts + 38,
	35:  __ccgo_ts + 38,
	36:  __ccgo_ts + 38,
	37:  __ccgo_ts + 38,
	38:  __ccgo_ts + 38,
	39:  __ccgo_ts + 38,
	40:  __ccgo_ts + 38,
	41:  __ccgo_ts + 38,
	42:  __ccgo_ts + 38,
	43:  __ccgo_ts + 38,
	44:  __ccgo_ts + 38,
	45:  __ccgo_ts + 38,
	46:  __ccgo_ts + 38,
	47:  __ccgo_ts + 38,
	48:  __ccgo_ts + 38,
	49:  __ccgo_ts + 38,
	50:  __ccgo_ts + 38,
	51:  __ccgo_ts + 38,
	52:  __ccgo_ts + 38,
	53:  __ccgo_ts + 38,
	54:  __ccgo_ts + 38,
	55:  __ccgo_ts + 38,
	56:  __ccgo_ts + 38,
	57:  __ccgo_ts + 38,
	58:  __ccgo_ts + 38,
	59:  __ccgo_ts + 43,
	60:  __ccgo_ts + 38,
	61:  __ccgo_ts + 38,
	62:  __ccgo_ts + 48,
	63:  __ccgo_ts + 38,
	64:  __ccgo_ts + 38,
	65:  __ccgo_ts + 38,
	66:  __ccgo_ts + 38,
	67:  __ccgo_ts + 38,
	68:  __ccgo_ts + 38,
	69:  __ccgo_ts + 48,
	70:  __ccgo_ts + 38,
	71:  __ccgo_ts + 56,
	72:  __ccgo_ts + 38,
	73:  __ccgo_ts + 38,
	74:  __ccgo_ts + 38,
	75:  __ccgo_ts + 38,
	76:  __ccgo_ts + 61,
	77:  __ccgo_ts + 70,
	78:  __ccgo_ts + 70,
	79:  __ccgo_ts + 79,
	80:  __ccgo_ts + 87,
	81:  __ccgo_ts + 89,
	82:  __ccgo_ts + 104,
	83:  __ccgo_ts + 113,
	84:  __ccgo_ts + 123,
	85:  __ccgo_ts + 133,
	86:  __ccgo_ts + 144,
	87:  __ccgo_ts + 158,
	88:  __ccgo_ts + 165,
	89:  __ccgo_ts + 171,
	90:  __ccgo_ts + 176,
	91:  __ccgo_ts + 183,
	92:  __ccgo_ts + 200,
	93:  __ccgo_ts + 207,
	94:  __ccgo_ts + 213,
	95:  __ccgo_ts + 222,
	96:  __ccgo_ts + 229,
	97:  __ccgo_ts + 238,
	98:  __ccgo_ts + 246,
	99:  __ccgo_ts + 261,
	100: __ccgo_ts + 276,
	101: __ccgo_ts + 280,
	102: __ccgo_ts + 286,
	103: __ccgo_ts + 291,
	104: __ccgo_ts + 297,
	105: __ccgo_ts + 303,
	106: __ccgo_ts + 310,
	107: __ccgo_ts + 315,
	108: __ccgo_ts + 320,
	109: __ccgo_ts + 327,
	110: __ccgo_ts + 332,
	111: __ccgo_ts + 339,
	112: __ccgo_ts + 346,
	113: __ccgo_ts + 354,
	114: __ccgo_ts + 362,
	115: __ccgo_ts + 369,
	116: __ccgo_ts + 374,
	117: __ccgo_ts + 383,
	118: __ccgo_ts + 389,
	119: __ccgo_ts + 394,
	120: __ccgo_ts + 403,
	121: __ccgo_ts + 408,
	122: __ccgo_ts + 417,
	123: __ccgo_ts + 433,
	124: __ccgo_ts + 447,
	125: __ccgo_ts + 461,
	126: __ccgo_ts + 473,
	127: __ccgo_ts + 473,
	128: __ccgo_ts + 473,
	129: __ccgo_ts + 473,
	130: __ccgo_ts + 473,
	131: __ccgo_ts + 473,
	132: __ccgo_ts + 480,
	133: __ccgo_ts + 488,
	134: __ccgo_ts + 495,
	135: __ccgo_ts + 500,
	136: __ccgo_ts + 507,
	137: __ccgo_ts + 513,
	138: __ccgo_ts + 520,
	139: __ccgo_ts + 530,
	140: __ccgo_ts + 537,
	141: __ccgo_ts + 546,
	142: __ccgo_ts + 562,
	143: __ccgo_ts + 578,
	144: __ccgo_ts + 586,
	145: __ccgo_ts + 593,
	146: __ccgo_ts + 603,
	147: __ccgo_ts + 611,
	148: __ccgo_ts + 618,
	149: __ccgo_ts + 624,
	150: __ccgo_ts + 631,
	151: __ccgo_ts + 636,
	152: __ccgo_ts + 646,
	153: __ccgo_ts + 650,
	154: __ccgo_ts + 658,
	155: __ccgo_ts + 665,
	156: __ccgo_ts + 676,
	157: __ccgo_ts + 687,
	158: __ccgo_ts + 702,
	159: __ccgo_ts + 716,
	160: __ccgo_ts + 734,
	161: __ccgo_ts + 752,
	162: __ccgo_ts + 770,
	163: __ccgo_ts + 788,
	164: __ccgo_ts + 806,
	165: __ccgo_ts + 824,
	166: __ccgo_ts + 842,
}

var ts_symbol_map = [167]TSSymbol{
	1:   uint16(anon_sym_Show),
	2:   uint16(anon_sym_Hide),
	3:   uint16(anon_sym_Minimal),
	4:   uint16(anon_sym_Import),
	5:   uint16(anon_sym_Optional),
	6:   uint16(anon_sym_AlternateQuality),
	7:   uint16(anon_sym_AlternateQuality),
	8:   uint16(anon_sym_AlternateQuality),
	9:   uint16(anon_sym_AlternateQuality),
	10:  uint16(anon_sym_AlternateQuality),
	11:  uint16(anon_sym_AlternateQuality),
	12:  uint16(anon_sym_AlternateQuality),
	13:  uint16(anon_sym_AlternateQuality),
	14:  uint16(anon_sym_AlternateQuality),
	15:  uint16(anon_sym_AlternateQuality),
	16:  uint16(anon_sym_AlternateQuality),
	17:  uint16(anon_sym_AlternateQuality),
	18:  uint16(anon_sym_AlternateQuality),
	19:  uint16(anon_sym_AlternateQuality),
	20:  uint16(anon_sym_AlternateQuality),
	21:  uint16(anon_sym_AlternateQuality),
	22:  uint16(anon_sym_AlternateQuality),
	23:  uint16(anon_sym_AlternateQuality),
	24:  uint16(anon_sym_AlternateQuality),
	25:  uint16(anon_sym_AlternateQuality),
	26:  uint16(anon_sym_AlternateQuality),
	27:  uint16(anon_sym_AlternateQuality),
	28:  uint16(anon_sym_AlternateQuality),
	29:  uint16(anon_sym_AlternateQuality),
	30:  uint16(anon_sym_AlternateQuality),
	31:  uint16(anon_sym_AlternateQuality),
	32:  uint16(anon_sym_AlternateQuality),
	33:  uint16(anon_sym_AlternateQuality),
	34:  uint16(anon_sym_AlternateQuality),
	35:  uint16(anon_sym_AlternateQuality),
	36:  uint16(anon_sym_AlternateQuality),
	37:  uint16(anon_sym_AlternateQuality),
	38:  uint16(anon_sym_AlternateQuality),
	39:  uint16(anon_sym_AlternateQuality),
	40:  uint16(anon_sym_AlternateQuality),
	41:  uint16(anon_sym_AlternateQuality),
	42:  uint16(anon_sym_AlternateQuality),
	43:  uint16(anon_sym_AlternateQuality),
	44:  uint16(anon_sym_AlternateQuality),
	45:  uint16(anon_sym_AlternateQuality),
	46:  uint16(anon_sym_AlternateQuality),
	47:  uint16(anon_sym_AlternateQuality),
	48:  uint16(anon_sym_AlternateQuality),
	49:  uint16(anon_sym_AlternateQuality),
	50:  uint16(anon_sym_AlternateQuality),
	51:  uint16(anon_sym_AlternateQuality),
	52:  uint16(anon_sym_AlternateQuality),
	53:  uint16(anon_sym_AlternateQuality),
	54:  uint16(anon_sym_AlternateQuality),
	55:  uint16(anon_sym_AlternateQuality),
	56:  uint16(anon_sym_AlternateQuality),
	57:  uint16(anon_sym_AlternateQuality),
	58:  uint16(anon_sym_AlternateQuality),
	59:  uint16(anon_sym_None),
	60:  uint16(anon_sym_AlternateQuality),
	61:  uint16(anon_sym_AlternateQuality),
	62:  uint16(aux_sym_action_token1),
	63:  uint16(anon_sym_AlternateQuality),
	64:  uint16(anon_sym_AlternateQuality),
	65:  uint16(anon_sym_AlternateQuality),
	66:  uint16(anon_sym_AlternateQuality),
	67:  uint16(anon_sym_AlternateQuality),
	68:  uint16(anon_sym_AlternateQuality),
	69:  uint16(aux_sym_action_token1),
	70:  uint16(anon_sym_AlternateQuality),
	71:  uint16(anon_sym_Temp),
	72:  uint16(anon_sym_AlternateQuality),
	73:  uint16(anon_sym_AlternateQuality),
	74:  uint16(anon_sym_AlternateQuality),
	75:  uint16(anon_sym_AlternateQuality),
	76:  uint16(anon_sym_Continue),
	77:  uint16(aux_sym__equal_operator_token1),
	78:  uint16(aux_sym__equal_operator_token1),
	79:  uint16(sym_boolean),
	80:  uint16(anon_sym_DQUOTE),
	81:  uint16(aux_sym_quality_token1),
	82:  uint16(anon_sym_Superior),
	83:  uint16(anon_sym_Divergent),
	84:  uint16(anon_sym_Anomalous),
	85:  uint16(anon_sym_Phantasmal),
	86:  uint16(aux_sym_rarity_token1),
	87:  uint16(anon_sym_Normal),
	88:  uint16(anon_sym_Magic),
	89:  uint16(anon_sym_Rare),
	90:  uint16(anon_sym_Unique),
	91:  uint16(aux_sym_influence_token1),
	92:  uint16(anon_sym_Shaper),
	93:  uint16(anon_sym_Elder),
	94:  uint16(anon_sym_Crusader),
	95:  uint16(anon_sym_Hunter),
	96:  uint16(anon_sym_Redeemer),
	97:  uint16(anon_sym_Warlord),
	98:  uint16(aux_sym_sockets_token1),
	99:  uint16(aux_sym_sockets_token2),
	100: uint16(anon_sym_Red),
	101: uint16(anon_sym_Green),
	102: uint16(anon_sym_Blue),
	103: uint16(anon_sym_Brown),
	104: uint16(anon_sym_White),
	105: uint16(anon_sym_Yellow),
	106: uint16(anon_sym_Cyan),
	107: uint16(anon_sym_Grey),
	108: uint16(anon_sym_Orange),
	109: uint16(anon_sym_Pink),
	110: uint16(anon_sym_Purple),
	111: uint16(anon_sym_Circle),
	112: uint16(anon_sym_Diamond),
	113: uint16(anon_sym_Hexagon),
	114: uint16(anon_sym_Square),
	115: uint16(anon_sym_Star),
	116: uint16(anon_sym_Triangle),
	117: uint16(anon_sym_Cross),
	118: uint16(anon_sym_Moon),
	119: uint16(anon_sym_Raindrop),
	120: uint16(anon_sym_Kite),
	121: uint16(anon_sym_Pentagon),
	122: uint16(anon_sym_UpsideDownHouse),
	123: uint16(aux_sym_string_token1),
	124: uint16(aux_sym_string_token2),
	125: uint16(aux_sym_file_token1),
	126: uint16(sym_number),
	127: uint16(sym_number),
	128: uint16(sym_number),
	129: uint16(sym_number),
	130: uint16(sym_number),
	131: uint16(sym_number),
	132: uint16(sym_comment),
	133: uint16(sym__space),
	134: uint16(sym__eol),
	135: uint16(sym_filter),
	136: uint16(sym_block),
	137: uint16(sym_import),
	138: uint16(sym_condition),
	139: uint16(sym_action),
	140: uint16(sym_continue),
	141: uint16(sym__equal_operator),
	142: uint16(sym__range_operator),
	143: uint16(sym_quality),
	144: uint16(sym_rarity),
	145: uint16(sym_influence),
	146: uint16(sym_sockets),
	147: uint16(sym_colour),
	148: uint16(sym_shape),
	149: uint16(sym_string),
	150: uint16(sym_file),
	151: uint16(sym__quantity),
	152: uint16(sym__id),
	153: uint16(sym__volume),
	154: uint16(sym__color),
	155: uint16(sym__icon_size),
	156: uint16(sym__font_size),
	157: uint16(aux_sym_filter_repeat1),
	158: uint16(aux_sym_block_repeat1),
	159: uint16(aux_sym_condition_repeat1),
	160: uint16(aux_sym_condition_repeat2),
	161: uint16(aux_sym_condition_repeat3),
	162: uint16(aux_sym_condition_repeat4),
	163: uint16(aux_sym_condition_repeat5),
	164: uint16(aux_sym_condition_repeat6),
	165: uint16(aux_sym_condition_repeat7),
	166: uint16(aux_sym_condition_repeat8),
}

var ts_symbol_metadata = [167]TSSymbolMetadata{
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
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	7: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	12: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	13: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	14: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
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
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	18: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	36: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	37: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	45: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	58: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	59: {
		Fvisible: libc.BoolUint8(true1 != 0),
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
	71: {
		Fvisible: libc.BoolUint8(true1 != 0),
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
	},
	81: {},
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
	86: {},
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
	91: {},
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
	98: {},
	99: {},
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
	123: {},
	124: {},
	125: {},
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
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	133: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	134: {
		Fnamed: libc.BoolUint8(true1 != 0),
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
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	141: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	142: {
		Fnamed: libc.BoolUint8(true1 != 0),
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
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	153: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	154: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	155: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	156: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	157: {},
	158: {},
	159: {},
	160: {},
	161: {},
	162: {},
	163: {},
	164: {},
	165: {},
	166: {},
}

type ts_field_identifiers = int32

const field_alpha = 1
const field_blue = 2
const field_class = 3
const field_enchantment = 4
const field_gem = 5
const field_green = 6
const field_id = 7
const field_modifier = 8
const field_red = 9
const field_size = 10
const field_sockets = 11
const field_type = 12
const field_volume = 13

var ts_field_names = [14]uintptr{
	0:  libc.UintptrFromInt32(0),
	1:  __ccgo_ts + 860,
	2:  __ccgo_ts + 866,
	3:  __ccgo_ts + 871,
	4:  __ccgo_ts + 877,
	5:  __ccgo_ts + 889,
	6:  __ccgo_ts + 893,
	7:  __ccgo_ts + 899,
	8:  __ccgo_ts + 902,
	9:  __ccgo_ts + 911,
	10: __ccgo_ts + 915,
	11: __ccgo_ts + 603,
	12: __ccgo_ts + 920,
	13: __ccgo_ts + 925,
}

var ts_field_map_slices = [27]TSMapSlice{
	1: {
		Flength: uint16(1),
	},
	2: {
		Findex:  uint16(1),
		Flength: uint16(1),
	},
	3: {
		Findex:  uint16(2),
		Flength: uint16(2),
	},
	4: {
		Findex:  uint16(4),
		Flength: uint16(1),
	},
	5: {
		Findex:  uint16(5),
		Flength: uint16(1),
	},
	6: {
		Findex:  uint16(6),
		Flength: uint16(2),
	},
	7: {
		Findex:  uint16(8),
		Flength: uint16(1),
	},
	8: {
		Findex:  uint16(9),
		Flength: uint16(1),
	},
	9: {
		Findex:  uint16(10),
		Flength: uint16(2),
	},
	10: {
		Findex:  uint16(12),
		Flength: uint16(1),
	},
	11: {
		Findex:  uint16(13),
		Flength: uint16(1),
	},
	12: {
		Findex:  uint16(14),
		Flength: uint16(2),
	},
	13: {
		Findex:  uint16(16),
		Flength: uint16(1),
	},
	14: {
		Findex:  uint16(17),
		Flength: uint16(1),
	},
	16: {
		Findex:  uint16(18),
		Flength: uint16(1),
	},
	17: {
		Findex:  uint16(19),
		Flength: uint16(1),
	},
	18: {
		Findex:  uint16(20),
		Flength: uint16(1),
	},
	19: {
		Findex:  uint16(21),
		Flength: uint16(1),
	},
	20: {
		Findex:  uint16(22),
		Flength: uint16(1),
	},
	21: {
		Findex:  uint16(23),
		Flength: uint16(1),
	},
	22: {
		Findex:  uint16(24),
		Flength: uint16(1),
	},
	23: {
		Findex:  uint16(25),
		Flength: uint16(2),
	},
	24: {
		Findex:  uint16(24),
		Flength: uint16(1),
	},
	25: {
		Findex:  uint16(27),
		Flength: uint16(3),
	},
	26: {
		Findex:  uint16(30),
		Flength: uint16(4),
	},
}

var ts_field_map_entries = [34]TSFieldMapEntry{
	0: {
		Ffield_id:    uint16(field_modifier),
		Fchild_index: uint8(1),
	},
	1: {
		Ffield_id:    uint16(field_modifier),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	2: {
		Ffield_id:  uint16(field_modifier),
		Finherited: libc.BoolUint8(true1 != 0),
	},
	3: {
		Ffield_id:    uint16(field_modifier),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	4: {
		Ffield_id:    uint16(field_type),
		Fchild_index: uint8(1),
	},
	5: {
		Ffield_id:    uint16(field_type),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	6: {
		Ffield_id:  uint16(field_type),
		Finherited: libc.BoolUint8(true1 != 0),
	},
	7: {
		Ffield_id:    uint16(field_type),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	8: {
		Ffield_id:    uint16(field_class),
		Fchild_index: uint8(1),
	},
	9: {
		Ffield_id:    uint16(field_class),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	10: {
		Ffield_id:  uint16(field_class),
		Finherited: libc.BoolUint8(true1 != 0),
	},
	11: {
		Ffield_id:    uint16(field_class),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	12: {
		Ffield_id:    uint16(field_enchantment),
		Fchild_index: uint8(1),
	},
	13: {
		Ffield_id:    uint16(field_enchantment),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	14: {
		Ffield_id:  uint16(field_enchantment),
		Finherited: libc.BoolUint8(true1 != 0),
	},
	15: {
		Ffield_id:    uint16(field_enchantment),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	16: {
		Ffield_id: uint16(field_sockets),
	},
	17: {
		Ffield_id:    uint16(field_gem),
		Fchild_index: uint8(2),
	},
	18: {
		Ffield_id:    uint16(field_id),
		Fchild_index: uint8(2),
	},
	19: {
		Ffield_id:    uint16(field_size),
		Fchild_index: uint8(2),
	},
	20: {
		Ffield_id:    uint16(field_modifier),
		Fchild_index: uint8(3),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	21: {
		Ffield_id:    uint16(field_type),
		Fchild_index: uint8(3),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	22: {
		Ffield_id:    uint16(field_class),
		Fchild_index: uint8(3),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	23: {
		Ffield_id:    uint16(field_enchantment),
		Fchild_index: uint8(3),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	24: {
		Ffield_id:    uint16(field_volume),
		Fchild_index: uint8(4),
	},
	25: {
		Ffield_id:    uint16(field_id),
		Fchild_index: uint8(2),
	},
	26: {
		Ffield_id:    uint16(field_volume),
		Fchild_index: uint8(4),
	},
	27: {
		Ffield_id:    uint16(field_blue),
		Fchild_index: uint8(6),
	},
	28: {
		Ffield_id:    uint16(field_green),
		Fchild_index: uint8(4),
	},
	29: {
		Ffield_id:    uint16(field_red),
		Fchild_index: uint8(2),
	},
	30: {
		Ffield_id:    uint16(field_alpha),
		Fchild_index: uint8(8),
	},
	31: {
		Ffield_id:    uint16(field_blue),
		Fchild_index: uint8(6),
	},
	32: {
		Ffield_id:    uint16(field_green),
		Fchild_index: uint8(4),
	},
	33: {
		Ffield_id:    uint16(field_red),
		Fchild_index: uint8(2),
	},
}

var ts_alias_sequences = [27][10]TSSymbol{
	0: {},
	15: {
		2: uint16(aux_sym_action_token1),
	},
	22: {
		2: uint16(aux_sym_action_token1),
	},
}

var ts_non_terminal_alias_map = [1]uint16_t{}

var ts_primary_state_ids = [195]TSStateId{
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
	171: uint16(171),
	172: uint16(172),
	173: uint16(173),
	174: uint16(174),
	175: uint16(175),
	176: uint16(176),
	177: uint16(177),
	178: uint16(178),
	179: uint16(179),
	180: uint16(180),
	181: uint16(181),
	182: uint16(182),
	183: uint16(183),
	184: uint16(184),
	185: uint16(185),
	186: uint16(186),
	187: uint16(187),
	188: uint16(113),
	189: uint16(176),
	190: uint16(116),
	191: uint16(160),
	192: uint16(174),
	193: uint16(146),
	194: uint16(194),
}

var aux_sym_string_token1_character_set_1 = [681]TSCharacterRange{
	0: {
		Fstart: int32(' '),
		Fend:   int32(' '),
	},
	1: {
		Fstart: int32('\''),
		Fend:   int32('\''),
	},
	2: {
		Fstart: int32(','),
		Fend:   int32('-'),
	},
	3: {
		Fstart: int32(':'),
		Fend:   int32(':'),
	},
	4: {
		Fstart: int32('A'),
		Fend:   int32('Z'),
	},
	5: {
		Fstart: int32('a'),
		Fend:   int32('z'),
	},
	6: {
		Fstart: int32(0xaa),
		Fend:   int32(0xaa),
	},
	7: {
		Fstart: int32(0xb5),
		Fend:   int32(0xb5),
	},
	8: {
		Fstart: int32(0xba),
		Fend:   int32(0xba),
	},
	9: {
		Fstart: int32(0xc0),
		Fend:   int32(0xd6),
	},
	10: {
		Fstart: int32(0xd8),
		Fend:   int32(0xf6),
	},
	11: {
		Fstart: int32(0xf8),
		Fend:   int32(0x2c1),
	},
	12: {
		Fstart: int32(0x2c6),
		Fend:   int32(0x2d1),
	},
	13: {
		Fstart: int32(0x2e0),
		Fend:   int32(0x2e4),
	},
	14: {
		Fstart: int32(0x2ec),
		Fend:   int32(0x2ec),
	},
	15: {
		Fstart: int32(0x2ee),
		Fend:   int32(0x2ee),
	},
	16: {
		Fstart: int32(0x370),
		Fend:   int32(0x374),
	},
	17: {
		Fstart: int32(0x376),
		Fend:   int32(0x377),
	},
	18: {
		Fstart: int32(0x37a),
		Fend:   int32(0x37d),
	},
	19: {
		Fstart: int32(0x37f),
		Fend:   int32(0x37f),
	},
	20: {
		Fstart: int32(0x386),
		Fend:   int32(0x386),
	},
	21: {
		Fstart: int32(0x388),
		Fend:   int32(0x38a),
	},
	22: {
		Fstart: int32(0x38c),
		Fend:   int32(0x38c),
	},
	23: {
		Fstart: int32(0x38e),
		Fend:   int32(0x3a1),
	},
	24: {
		Fstart: int32(0x3a3),
		Fend:   int32(0x3f5),
	},
	25: {
		Fstart: int32(0x3f7),
		Fend:   int32(0x481),
	},
	26: {
		Fstart: int32(0x48a),
		Fend:   int32(0x52f),
	},
	27: {
		Fstart: int32(0x531),
		Fend:   int32(0x556),
	},
	28: {
		Fstart: int32(0x559),
		Fend:   int32(0x559),
	},
	29: {
		Fstart: int32(0x560),
		Fend:   int32(0x588),
	},
	30: {
		Fstart: int32(0x5d0),
		Fend:   int32(0x5ea),
	},
	31: {
		Fstart: int32(0x5ef),
		Fend:   int32(0x5f2),
	},
	32: {
		Fstart: int32(0x620),
		Fend:   int32(0x64a),
	},
	33: {
		Fstart: int32(0x66e),
		Fend:   int32(0x66f),
	},
	34: {
		Fstart: int32(0x671),
		Fend:   int32(0x6d3),
	},
	35: {
		Fstart: int32(0x6d5),
		Fend:   int32(0x6d5),
	},
	36: {
		Fstart: int32(0x6e5),
		Fend:   int32(0x6e6),
	},
	37: {
		Fstart: int32(0x6ee),
		Fend:   int32(0x6ef),
	},
	38: {
		Fstart: int32(0x6fa),
		Fend:   int32(0x6fc),
	},
	39: {
		Fstart: int32(0x6ff),
		Fend:   int32(0x6ff),
	},
	40: {
		Fstart: int32(0x710),
		Fend:   int32(0x710),
	},
	41: {
		Fstart: int32(0x712),
		Fend:   int32(0x72f),
	},
	42: {
		Fstart: int32(0x74d),
		Fend:   int32(0x7a5),
	},
	43: {
		Fstart: int32(0x7b1),
		Fend:   int32(0x7b1),
	},
	44: {
		Fstart: int32(0x7ca),
		Fend:   int32(0x7ea),
	},
	45: {
		Fstart: int32(0x7f4),
		Fend:   int32(0x7f5),
	},
	46: {
		Fstart: int32(0x7fa),
		Fend:   int32(0x7fa),
	},
	47: {
		Fstart: int32(0x800),
		Fend:   int32(0x815),
	},
	48: {
		Fstart: int32(0x81a),
		Fend:   int32(0x81a),
	},
	49: {
		Fstart: int32(0x824),
		Fend:   int32(0x824),
	},
	50: {
		Fstart: int32(0x828),
		Fend:   int32(0x828),
	},
	51: {
		Fstart: int32(0x840),
		Fend:   int32(0x858),
	},
	52: {
		Fstart: int32(0x860),
		Fend:   int32(0x86a),
	},
	53: {
		Fstart: int32(0x870),
		Fend:   int32(0x887),
	},
	54: {
		Fstart: int32(0x889),
		Fend:   int32(0x88e),
	},
	55: {
		Fstart: int32(0x8a0),
		Fend:   int32(0x8c9),
	},
	56: {
		Fstart: int32(0x904),
		Fend:   int32(0x939),
	},
	57: {
		Fstart: int32(0x93d),
		Fend:   int32(0x93d),
	},
	58: {
		Fstart: int32(0x950),
		Fend:   int32(0x950),
	},
	59: {
		Fstart: int32(0x958),
		Fend:   int32(0x961),
	},
	60: {
		Fstart: int32(0x971),
		Fend:   int32(0x980),
	},
	61: {
		Fstart: int32(0x985),
		Fend:   int32(0x98c),
	},
	62: {
		Fstart: int32(0x98f),
		Fend:   int32(0x990),
	},
	63: {
		Fstart: int32(0x993),
		Fend:   int32(0x9a8),
	},
	64: {
		Fstart: int32(0x9aa),
		Fend:   int32(0x9b0),
	},
	65: {
		Fstart: int32(0x9b2),
		Fend:   int32(0x9b2),
	},
	66: {
		Fstart: int32(0x9b6),
		Fend:   int32(0x9b9),
	},
	67: {
		Fstart: int32(0x9bd),
		Fend:   int32(0x9bd),
	},
	68: {
		Fstart: int32(0x9ce),
		Fend:   int32(0x9ce),
	},
	69: {
		Fstart: int32(0x9dc),
		Fend:   int32(0x9dd),
	},
	70: {
		Fstart: int32(0x9df),
		Fend:   int32(0x9e1),
	},
	71: {
		Fstart: int32(0x9f0),
		Fend:   int32(0x9f1),
	},
	72: {
		Fstart: int32(0x9fc),
		Fend:   int32(0x9fc),
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
		Fstart: int32(0xa59),
		Fend:   int32(0xa5c),
	},
	81: {
		Fstart: int32(0xa5e),
		Fend:   int32(0xa5e),
	},
	82: {
		Fstart: int32(0xa72),
		Fend:   int32(0xa74),
	},
	83: {
		Fstart: int32(0xa85),
		Fend:   int32(0xa8d),
	},
	84: {
		Fstart: int32(0xa8f),
		Fend:   int32(0xa91),
	},
	85: {
		Fstart: int32(0xa93),
		Fend:   int32(0xaa8),
	},
	86: {
		Fstart: int32(0xaaa),
		Fend:   int32(0xab0),
	},
	87: {
		Fstart: int32(0xab2),
		Fend:   int32(0xab3),
	},
	88: {
		Fstart: int32(0xab5),
		Fend:   int32(0xab9),
	},
	89: {
		Fstart: int32(0xabd),
		Fend:   int32(0xabd),
	},
	90: {
		Fstart: int32(0xad0),
		Fend:   int32(0xad0),
	},
	91: {
		Fstart: int32(0xae0),
		Fend:   int32(0xae1),
	},
	92: {
		Fstart: int32(0xaf9),
		Fend:   int32(0xaf9),
	},
	93: {
		Fstart: int32(0xb05),
		Fend:   int32(0xb0c),
	},
	94: {
		Fstart: int32(0xb0f),
		Fend:   int32(0xb10),
	},
	95: {
		Fstart: int32(0xb13),
		Fend:   int32(0xb28),
	},
	96: {
		Fstart: int32(0xb2a),
		Fend:   int32(0xb30),
	},
	97: {
		Fstart: int32(0xb32),
		Fend:   int32(0xb33),
	},
	98: {
		Fstart: int32(0xb35),
		Fend:   int32(0xb39),
	},
	99: {
		Fstart: int32(0xb3d),
		Fend:   int32(0xb3d),
	},
	100: {
		Fstart: int32(0xb5c),
		Fend:   int32(0xb5d),
	},
	101: {
		Fstart: int32(0xb5f),
		Fend:   int32(0xb61),
	},
	102: {
		Fstart: int32(0xb71),
		Fend:   int32(0xb71),
	},
	103: {
		Fstart: int32(0xb83),
		Fend:   int32(0xb83),
	},
	104: {
		Fstart: int32(0xb85),
		Fend:   int32(0xb8a),
	},
	105: {
		Fstart: int32(0xb8e),
		Fend:   int32(0xb90),
	},
	106: {
		Fstart: int32(0xb92),
		Fend:   int32(0xb95),
	},
	107: {
		Fstart: int32(0xb99),
		Fend:   int32(0xb9a),
	},
	108: {
		Fstart: int32(0xb9c),
		Fend:   int32(0xb9c),
	},
	109: {
		Fstart: int32(0xb9e),
		Fend:   int32(0xb9f),
	},
	110: {
		Fstart: int32(0xba3),
		Fend:   int32(0xba4),
	},
	111: {
		Fstart: int32(0xba8),
		Fend:   int32(0xbaa),
	},
	112: {
		Fstart: int32(0xbae),
		Fend:   int32(0xbb9),
	},
	113: {
		Fstart: int32(0xbd0),
		Fend:   int32(0xbd0),
	},
	114: {
		Fstart: int32(0xc05),
		Fend:   int32(0xc0c),
	},
	115: {
		Fstart: int32(0xc0e),
		Fend:   int32(0xc10),
	},
	116: {
		Fstart: int32(0xc12),
		Fend:   int32(0xc28),
	},
	117: {
		Fstart: int32(0xc2a),
		Fend:   int32(0xc39),
	},
	118: {
		Fstart: int32(0xc3d),
		Fend:   int32(0xc3d),
	},
	119: {
		Fstart: int32(0xc58),
		Fend:   int32(0xc5a),
	},
	120: {
		Fstart: int32(0xc5d),
		Fend:   int32(0xc5d),
	},
	121: {
		Fstart: int32(0xc60),
		Fend:   int32(0xc61),
	},
	122: {
		Fstart: int32(0xc80),
		Fend:   int32(0xc80),
	},
	123: {
		Fstart: int32(0xc85),
		Fend:   int32(0xc8c),
	},
	124: {
		Fstart: int32(0xc8e),
		Fend:   int32(0xc90),
	},
	125: {
		Fstart: int32(0xc92),
		Fend:   int32(0xca8),
	},
	126: {
		Fstart: int32(0xcaa),
		Fend:   int32(0xcb3),
	},
	127: {
		Fstart: int32(0xcb5),
		Fend:   int32(0xcb9),
	},
	128: {
		Fstart: int32(0xcbd),
		Fend:   int32(0xcbd),
	},
	129: {
		Fstart: int32(0xcdd),
		Fend:   int32(0xcde),
	},
	130: {
		Fstart: int32(0xce0),
		Fend:   int32(0xce1),
	},
	131: {
		Fstart: int32(0xcf1),
		Fend:   int32(0xcf2),
	},
	132: {
		Fstart: int32(0xd04),
		Fend:   int32(0xd0c),
	},
	133: {
		Fstart: int32(0xd0e),
		Fend:   int32(0xd10),
	},
	134: {
		Fstart: int32(0xd12),
		Fend:   int32(0xd3a),
	},
	135: {
		Fstart: int32(0xd3d),
		Fend:   int32(0xd3d),
	},
	136: {
		Fstart: int32(0xd4e),
		Fend:   int32(0xd4e),
	},
	137: {
		Fstart: int32(0xd54),
		Fend:   int32(0xd56),
	},
	138: {
		Fstart: int32(0xd5f),
		Fend:   int32(0xd61),
	},
	139: {
		Fstart: int32(0xd7a),
		Fend:   int32(0xd7f),
	},
	140: {
		Fstart: int32(0xd85),
		Fend:   int32(0xd96),
	},
	141: {
		Fstart: int32(0xd9a),
		Fend:   int32(0xdb1),
	},
	142: {
		Fstart: int32(0xdb3),
		Fend:   int32(0xdbb),
	},
	143: {
		Fstart: int32(0xdbd),
		Fend:   int32(0xdbd),
	},
	144: {
		Fstart: int32(0xdc0),
		Fend:   int32(0xdc6),
	},
	145: {
		Fstart: int32(0xe01),
		Fend:   int32(0xe30),
	},
	146: {
		Fstart: int32(0xe32),
		Fend:   int32(0xe33),
	},
	147: {
		Fstart: int32(0xe40),
		Fend:   int32(0xe46),
	},
	148: {
		Fstart: int32(0xe81),
		Fend:   int32(0xe82),
	},
	149: {
		Fstart: int32(0xe84),
		Fend:   int32(0xe84),
	},
	150: {
		Fstart: int32(0xe86),
		Fend:   int32(0xe8a),
	},
	151: {
		Fstart: int32(0xe8c),
		Fend:   int32(0xea3),
	},
	152: {
		Fstart: int32(0xea5),
		Fend:   int32(0xea5),
	},
	153: {
		Fstart: int32(0xea7),
		Fend:   int32(0xeb0),
	},
	154: {
		Fstart: int32(0xeb2),
		Fend:   int32(0xeb3),
	},
	155: {
		Fstart: int32(0xebd),
		Fend:   int32(0xebd),
	},
	156: {
		Fstart: int32(0xec0),
		Fend:   int32(0xec4),
	},
	157: {
		Fstart: int32(0xec6),
		Fend:   int32(0xec6),
	},
	158: {
		Fstart: int32(0xedc),
		Fend:   int32(0xedf),
	},
	159: {
		Fstart: int32(0xf00),
		Fend:   int32(0xf00),
	},
	160: {
		Fstart: int32(0xf40),
		Fend:   int32(0xf47),
	},
	161: {
		Fstart: int32(0xf49),
		Fend:   int32(0xf6c),
	},
	162: {
		Fstart: int32(0xf88),
		Fend:   int32(0xf8c),
	},
	163: {
		Fstart: int32(0x1000),
		Fend:   int32(0x102a),
	},
	164: {
		Fstart: int32(0x103f),
		Fend:   int32(0x103f),
	},
	165: {
		Fstart: int32(0x1050),
		Fend:   int32(0x1055),
	},
	166: {
		Fstart: int32(0x105a),
		Fend:   int32(0x105d),
	},
	167: {
		Fstart: int32(0x1061),
		Fend:   int32(0x1061),
	},
	168: {
		Fstart: int32(0x1065),
		Fend:   int32(0x1066),
	},
	169: {
		Fstart: int32(0x106e),
		Fend:   int32(0x1070),
	},
	170: {
		Fstart: int32(0x1075),
		Fend:   int32(0x1081),
	},
	171: {
		Fstart: int32(0x108e),
		Fend:   int32(0x108e),
	},
	172: {
		Fstart: int32(0x10a0),
		Fend:   int32(0x10c5),
	},
	173: {
		Fstart: int32(0x10c7),
		Fend:   int32(0x10c7),
	},
	174: {
		Fstart: int32(0x10cd),
		Fend:   int32(0x10cd),
	},
	175: {
		Fstart: int32(0x10d0),
		Fend:   int32(0x10fa),
	},
	176: {
		Fstart: int32(0x10fc),
		Fend:   int32(0x1248),
	},
	177: {
		Fstart: int32(0x124a),
		Fend:   int32(0x124d),
	},
	178: {
		Fstart: int32(0x1250),
		Fend:   int32(0x1256),
	},
	179: {
		Fstart: int32(0x1258),
		Fend:   int32(0x1258),
	},
	180: {
		Fstart: int32(0x125a),
		Fend:   int32(0x125d),
	},
	181: {
		Fstart: int32(0x1260),
		Fend:   int32(0x1288),
	},
	182: {
		Fstart: int32(0x128a),
		Fend:   int32(0x128d),
	},
	183: {
		Fstart: int32(0x1290),
		Fend:   int32(0x12b0),
	},
	184: {
		Fstart: int32(0x12b2),
		Fend:   int32(0x12b5),
	},
	185: {
		Fstart: int32(0x12b8),
		Fend:   int32(0x12be),
	},
	186: {
		Fstart: int32(0x12c0),
		Fend:   int32(0x12c0),
	},
	187: {
		Fstart: int32(0x12c2),
		Fend:   int32(0x12c5),
	},
	188: {
		Fstart: int32(0x12c8),
		Fend:   int32(0x12d6),
	},
	189: {
		Fstart: int32(0x12d8),
		Fend:   int32(0x1310),
	},
	190: {
		Fstart: int32(0x1312),
		Fend:   int32(0x1315),
	},
	191: {
		Fstart: int32(0x1318),
		Fend:   int32(0x135a),
	},
	192: {
		Fstart: int32(0x1380),
		Fend:   int32(0x138f),
	},
	193: {
		Fstart: int32(0x13a0),
		Fend:   int32(0x13f5),
	},
	194: {
		Fstart: int32(0x13f8),
		Fend:   int32(0x13fd),
	},
	195: {
		Fstart: int32(0x1401),
		Fend:   int32(0x166c),
	},
	196: {
		Fstart: int32(0x166f),
		Fend:   int32(0x167f),
	},
	197: {
		Fstart: int32(0x1681),
		Fend:   int32(0x169a),
	},
	198: {
		Fstart: int32(0x16a0),
		Fend:   int32(0x16ea),
	},
	199: {
		Fstart: int32(0x16f1),
		Fend:   int32(0x16f8),
	},
	200: {
		Fstart: int32(0x1700),
		Fend:   int32(0x1711),
	},
	201: {
		Fstart: int32(0x171f),
		Fend:   int32(0x1731),
	},
	202: {
		Fstart: int32(0x1740),
		Fend:   int32(0x1751),
	},
	203: {
		Fstart: int32(0x1760),
		Fend:   int32(0x176c),
	},
	204: {
		Fstart: int32(0x176e),
		Fend:   int32(0x1770),
	},
	205: {
		Fstart: int32(0x1780),
		Fend:   int32(0x17b3),
	},
	206: {
		Fstart: int32(0x17d7),
		Fend:   int32(0x17d7),
	},
	207: {
		Fstart: int32(0x17dc),
		Fend:   int32(0x17dc),
	},
	208: {
		Fstart: int32(0x1820),
		Fend:   int32(0x1878),
	},
	209: {
		Fstart: int32(0x1880),
		Fend:   int32(0x1884),
	},
	210: {
		Fstart: int32(0x1887),
		Fend:   int32(0x18a8),
	},
	211: {
		Fstart: int32(0x18aa),
		Fend:   int32(0x18aa),
	},
	212: {
		Fstart: int32(0x18b0),
		Fend:   int32(0x18f5),
	},
	213: {
		Fstart: int32(0x1900),
		Fend:   int32(0x191e),
	},
	214: {
		Fstart: int32(0x1950),
		Fend:   int32(0x196d),
	},
	215: {
		Fstart: int32(0x1970),
		Fend:   int32(0x1974),
	},
	216: {
		Fstart: int32(0x1980),
		Fend:   int32(0x19ab),
	},
	217: {
		Fstart: int32(0x19b0),
		Fend:   int32(0x19c9),
	},
	218: {
		Fstart: int32(0x1a00),
		Fend:   int32(0x1a16),
	},
	219: {
		Fstart: int32(0x1a20),
		Fend:   int32(0x1a54),
	},
	220: {
		Fstart: int32(0x1aa7),
		Fend:   int32(0x1aa7),
	},
	221: {
		Fstart: int32(0x1b05),
		Fend:   int32(0x1b33),
	},
	222: {
		Fstart: int32(0x1b45),
		Fend:   int32(0x1b4c),
	},
	223: {
		Fstart: int32(0x1b83),
		Fend:   int32(0x1ba0),
	},
	224: {
		Fstart: int32(0x1bae),
		Fend:   int32(0x1baf),
	},
	225: {
		Fstart: int32(0x1bba),
		Fend:   int32(0x1be5),
	},
	226: {
		Fstart: int32(0x1c00),
		Fend:   int32(0x1c23),
	},
	227: {
		Fstart: int32(0x1c4d),
		Fend:   int32(0x1c4f),
	},
	228: {
		Fstart: int32(0x1c5a),
		Fend:   int32(0x1c7d),
	},
	229: {
		Fstart: int32(0x1c80),
		Fend:   int32(0x1c8a),
	},
	230: {
		Fstart: int32(0x1c90),
		Fend:   int32(0x1cba),
	},
	231: {
		Fstart: int32(0x1cbd),
		Fend:   int32(0x1cbf),
	},
	232: {
		Fstart: int32(0x1ce9),
		Fend:   int32(0x1cec),
	},
	233: {
		Fstart: int32(0x1cee),
		Fend:   int32(0x1cf3),
	},
	234: {
		Fstart: int32(0x1cf5),
		Fend:   int32(0x1cf6),
	},
	235: {
		Fstart: int32(0x1cfa),
		Fend:   int32(0x1cfa),
	},
	236: {
		Fstart: int32(0x1d00),
		Fend:   int32(0x1dbf),
	},
	237: {
		Fstart: int32(0x1e00),
		Fend:   int32(0x1f15),
	},
	238: {
		Fstart: int32(0x1f18),
		Fend:   int32(0x1f1d),
	},
	239: {
		Fstart: int32(0x1f20),
		Fend:   int32(0x1f45),
	},
	240: {
		Fstart: int32(0x1f48),
		Fend:   int32(0x1f4d),
	},
	241: {
		Fstart: int32(0x1f50),
		Fend:   int32(0x1f57),
	},
	242: {
		Fstart: int32(0x1f59),
		Fend:   int32(0x1f59),
	},
	243: {
		Fstart: int32(0x1f5b),
		Fend:   int32(0x1f5b),
	},
	244: {
		Fstart: int32(0x1f5d),
		Fend:   int32(0x1f5d),
	},
	245: {
		Fstart: int32(0x1f5f),
		Fend:   int32(0x1f7d),
	},
	246: {
		Fstart: int32(0x1f80),
		Fend:   int32(0x1fb4),
	},
	247: {
		Fstart: int32(0x1fb6),
		Fend:   int32(0x1fbc),
	},
	248: {
		Fstart: int32(0x1fbe),
		Fend:   int32(0x1fbe),
	},
	249: {
		Fstart: int32(0x1fc2),
		Fend:   int32(0x1fc4),
	},
	250: {
		Fstart: int32(0x1fc6),
		Fend:   int32(0x1fcc),
	},
	251: {
		Fstart: int32(0x1fd0),
		Fend:   int32(0x1fd3),
	},
	252: {
		Fstart: int32(0x1fd6),
		Fend:   int32(0x1fdb),
	},
	253: {
		Fstart: int32(0x1fe0),
		Fend:   int32(0x1fec),
	},
	254: {
		Fstart: int32(0x1ff2),
		Fend:   int32(0x1ff4),
	},
	255: {
		Fstart: int32(0x1ff6),
		Fend:   int32(0x1ffc),
	},
	256: {
		Fstart: int32(0x2071),
		Fend:   int32(0x2071),
	},
	257: {
		Fstart: int32(0x207f),
		Fend:   int32(0x207f),
	},
	258: {
		Fstart: int32(0x2090),
		Fend:   int32(0x209c),
	},
	259: {
		Fstart: int32(0x2102),
		Fend:   int32(0x2102),
	},
	260: {
		Fstart: int32(0x2107),
		Fend:   int32(0x2107),
	},
	261: {
		Fstart: int32(0x210a),
		Fend:   int32(0x2113),
	},
	262: {
		Fstart: int32(0x2115),
		Fend:   int32(0x2115),
	},
	263: {
		Fstart: int32(0x2119),
		Fend:   int32(0x211d),
	},
	264: {
		Fstart: int32(0x2124),
		Fend:   int32(0x2124),
	},
	265: {
		Fstart: int32(0x2126),
		Fend:   int32(0x2126),
	},
	266: {
		Fstart: int32(0x2128),
		Fend:   int32(0x2128),
	},
	267: {
		Fstart: int32(0x212a),
		Fend:   int32(0x212d),
	},
	268: {
		Fstart: int32(0x212f),
		Fend:   int32(0x2139),
	},
	269: {
		Fstart: int32(0x213c),
		Fend:   int32(0x213f),
	},
	270: {
		Fstart: int32(0x2145),
		Fend:   int32(0x2149),
	},
	271: {
		Fstart: int32(0x214e),
		Fend:   int32(0x214e),
	},
	272: {
		Fstart: int32(0x2183),
		Fend:   int32(0x2184),
	},
	273: {
		Fstart: int32(0x2c00),
		Fend:   int32(0x2ce4),
	},
	274: {
		Fstart: int32(0x2ceb),
		Fend:   int32(0x2cee),
	},
	275: {
		Fstart: int32(0x2cf2),
		Fend:   int32(0x2cf3),
	},
	276: {
		Fstart: int32(0x2d00),
		Fend:   int32(0x2d25),
	},
	277: {
		Fstart: int32(0x2d27),
		Fend:   int32(0x2d27),
	},
	278: {
		Fstart: int32(0x2d2d),
		Fend:   int32(0x2d2d),
	},
	279: {
		Fstart: int32(0x2d30),
		Fend:   int32(0x2d67),
	},
	280: {
		Fstart: int32(0x2d6f),
		Fend:   int32(0x2d6f),
	},
	281: {
		Fstart: int32(0x2d80),
		Fend:   int32(0x2d96),
	},
	282: {
		Fstart: int32(0x2da0),
		Fend:   int32(0x2da6),
	},
	283: {
		Fstart: int32(0x2da8),
		Fend:   int32(0x2dae),
	},
	284: {
		Fstart: int32(0x2db0),
		Fend:   int32(0x2db6),
	},
	285: {
		Fstart: int32(0x2db8),
		Fend:   int32(0x2dbe),
	},
	286: {
		Fstart: int32(0x2dc0),
		Fend:   int32(0x2dc6),
	},
	287: {
		Fstart: int32(0x2dc8),
		Fend:   int32(0x2dce),
	},
	288: {
		Fstart: int32(0x2dd0),
		Fend:   int32(0x2dd6),
	},
	289: {
		Fstart: int32(0x2dd8),
		Fend:   int32(0x2dde),
	},
	290: {
		Fstart: int32(0x2e2f),
		Fend:   int32(0x2e2f),
	},
	291: {
		Fstart: int32(0x3005),
		Fend:   int32(0x3006),
	},
	292: {
		Fstart: int32(0x3031),
		Fend:   int32(0x3035),
	},
	293: {
		Fstart: int32(0x303b),
		Fend:   int32(0x303c),
	},
	294: {
		Fstart: int32(0x3041),
		Fend:   int32(0x3096),
	},
	295: {
		Fstart: int32(0x309d),
		Fend:   int32(0x309f),
	},
	296: {
		Fstart: int32(0x30a1),
		Fend:   int32(0x30fa),
	},
	297: {
		Fstart: int32(0x30fc),
		Fend:   int32(0x30ff),
	},
	298: {
		Fstart: int32(0x3105),
		Fend:   int32(0x312f),
	},
	299: {
		Fstart: int32(0x3131),
		Fend:   int32(0x318e),
	},
	300: {
		Fstart: int32(0x31a0),
		Fend:   int32(0x31bf),
	},
	301: {
		Fstart: int32(0x31f0),
		Fend:   int32(0x31ff),
	},
	302: {
		Fstart: int32(0x3400),
		Fend:   int32(0x4dbf),
	},
	303: {
		Fstart: int32(0x4e00),
		Fend:   int32(0xa48c),
	},
	304: {
		Fstart: int32(0xa4d0),
		Fend:   int32(0xa4fd),
	},
	305: {
		Fstart: int32(0xa500),
		Fend:   int32(0xa60c),
	},
	306: {
		Fstart: int32(0xa610),
		Fend:   int32(0xa61f),
	},
	307: {
		Fstart: int32(0xa62a),
		Fend:   int32(0xa62b),
	},
	308: {
		Fstart: int32(0xa640),
		Fend:   int32(0xa66e),
	},
	309: {
		Fstart: int32(0xa67f),
		Fend:   int32(0xa69d),
	},
	310: {
		Fstart: int32(0xa6a0),
		Fend:   int32(0xa6e5),
	},
	311: {
		Fstart: int32(0xa717),
		Fend:   int32(0xa71f),
	},
	312: {
		Fstart: int32(0xa722),
		Fend:   int32(0xa788),
	},
	313: {
		Fstart: int32(0xa78b),
		Fend:   int32(0xa7cd),
	},
	314: {
		Fstart: int32(0xa7d0),
		Fend:   int32(0xa7d1),
	},
	315: {
		Fstart: int32(0xa7d3),
		Fend:   int32(0xa7d3),
	},
	316: {
		Fstart: int32(0xa7d5),
		Fend:   int32(0xa7dc),
	},
	317: {
		Fstart: int32(0xa7f2),
		Fend:   int32(0xa801),
	},
	318: {
		Fstart: int32(0xa803),
		Fend:   int32(0xa805),
	},
	319: {
		Fstart: int32(0xa807),
		Fend:   int32(0xa80a),
	},
	320: {
		Fstart: int32(0xa80c),
		Fend:   int32(0xa822),
	},
	321: {
		Fstart: int32(0xa840),
		Fend:   int32(0xa873),
	},
	322: {
		Fstart: int32(0xa882),
		Fend:   int32(0xa8b3),
	},
	323: {
		Fstart: int32(0xa8f2),
		Fend:   int32(0xa8f7),
	},
	324: {
		Fstart: int32(0xa8fb),
		Fend:   int32(0xa8fb),
	},
	325: {
		Fstart: int32(0xa8fd),
		Fend:   int32(0xa8fe),
	},
	326: {
		Fstart: int32(0xa90a),
		Fend:   int32(0xa925),
	},
	327: {
		Fstart: int32(0xa930),
		Fend:   int32(0xa946),
	},
	328: {
		Fstart: int32(0xa960),
		Fend:   int32(0xa97c),
	},
	329: {
		Fstart: int32(0xa984),
		Fend:   int32(0xa9b2),
	},
	330: {
		Fstart: int32(0xa9cf),
		Fend:   int32(0xa9cf),
	},
	331: {
		Fstart: int32(0xa9e0),
		Fend:   int32(0xa9e4),
	},
	332: {
		Fstart: int32(0xa9e6),
		Fend:   int32(0xa9ef),
	},
	333: {
		Fstart: int32(0xa9fa),
		Fend:   int32(0xa9fe),
	},
	334: {
		Fstart: int32(0xaa00),
		Fend:   int32(0xaa28),
	},
	335: {
		Fstart: int32(0xaa40),
		Fend:   int32(0xaa42),
	},
	336: {
		Fstart: int32(0xaa44),
		Fend:   int32(0xaa4b),
	},
	337: {
		Fstart: int32(0xaa60),
		Fend:   int32(0xaa76),
	},
	338: {
		Fstart: int32(0xaa7a),
		Fend:   int32(0xaa7a),
	},
	339: {
		Fstart: int32(0xaa7e),
		Fend:   int32(0xaaaf),
	},
	340: {
		Fstart: int32(0xaab1),
		Fend:   int32(0xaab1),
	},
	341: {
		Fstart: int32(0xaab5),
		Fend:   int32(0xaab6),
	},
	342: {
		Fstart: int32(0xaab9),
		Fend:   int32(0xaabd),
	},
	343: {
		Fstart: int32(0xaac0),
		Fend:   int32(0xaac0),
	},
	344: {
		Fstart: int32(0xaac2),
		Fend:   int32(0xaac2),
	},
	345: {
		Fstart: int32(0xaadb),
		Fend:   int32(0xaadd),
	},
	346: {
		Fstart: int32(0xaae0),
		Fend:   int32(0xaaea),
	},
	347: {
		Fstart: int32(0xaaf2),
		Fend:   int32(0xaaf4),
	},
	348: {
		Fstart: int32(0xab01),
		Fend:   int32(0xab06),
	},
	349: {
		Fstart: int32(0xab09),
		Fend:   int32(0xab0e),
	},
	350: {
		Fstart: int32(0xab11),
		Fend:   int32(0xab16),
	},
	351: {
		Fstart: int32(0xab20),
		Fend:   int32(0xab26),
	},
	352: {
		Fstart: int32(0xab28),
		Fend:   int32(0xab2e),
	},
	353: {
		Fstart: int32(0xab30),
		Fend:   int32(0xab5a),
	},
	354: {
		Fstart: int32(0xab5c),
		Fend:   int32(0xab69),
	},
	355: {
		Fstart: int32(0xab70),
		Fend:   int32(0xabe2),
	},
	356: {
		Fstart: int32(0xac00),
		Fend:   int32(0xd7a3),
	},
	357: {
		Fstart: int32(0xd7b0),
		Fend:   int32(0xd7c6),
	},
	358: {
		Fstart: int32(0xd7cb),
		Fend:   int32(0xd7fb),
	},
	359: {
		Fstart: int32(0xf900),
		Fend:   int32(0xfa6d),
	},
	360: {
		Fstart: int32(0xfa70),
		Fend:   int32(0xfad9),
	},
	361: {
		Fstart: int32(0xfb00),
		Fend:   int32(0xfb06),
	},
	362: {
		Fstart: int32(0xfb13),
		Fend:   int32(0xfb17),
	},
	363: {
		Fstart: int32(0xfb1d),
		Fend:   int32(0xfb1d),
	},
	364: {
		Fstart: int32(0xfb1f),
		Fend:   int32(0xfb28),
	},
	365: {
		Fstart: int32(0xfb2a),
		Fend:   int32(0xfb36),
	},
	366: {
		Fstart: int32(0xfb38),
		Fend:   int32(0xfb3c),
	},
	367: {
		Fstart: int32(0xfb3e),
		Fend:   int32(0xfb3e),
	},
	368: {
		Fstart: int32(0xfb40),
		Fend:   int32(0xfb41),
	},
	369: {
		Fstart: int32(0xfb43),
		Fend:   int32(0xfb44),
	},
	370: {
		Fstart: int32(0xfb46),
		Fend:   int32(0xfbb1),
	},
	371: {
		Fstart: int32(0xfbd3),
		Fend:   int32(0xfd3d),
	},
	372: {
		Fstart: int32(0xfd50),
		Fend:   int32(0xfd8f),
	},
	373: {
		Fstart: int32(0xfd92),
		Fend:   int32(0xfdc7),
	},
	374: {
		Fstart: int32(0xfdf0),
		Fend:   int32(0xfdfb),
	},
	375: {
		Fstart: int32(0xfe70),
		Fend:   int32(0xfe74),
	},
	376: {
		Fstart: int32(0xfe76),
		Fend:   int32(0xfefc),
	},
	377: {
		Fstart: int32(0xff21),
		Fend:   int32(0xff3a),
	},
	378: {
		Fstart: int32(0xff41),
		Fend:   int32(0xff5a),
	},
	379: {
		Fstart: int32(0xff66),
		Fend:   int32(0xffbe),
	},
	380: {
		Fstart: int32(0xffc2),
		Fend:   int32(0xffc7),
	},
	381: {
		Fstart: int32(0xffca),
		Fend:   int32(0xffcf),
	},
	382: {
		Fstart: int32(0xffd2),
		Fend:   int32(0xffd7),
	},
	383: {
		Fstart: int32(0xffda),
		Fend:   int32(0xffdc),
	},
	384: {
		Fstart: int32(0x10000),
		Fend:   int32(0x1000b),
	},
	385: {
		Fstart: int32(0x1000d),
		Fend:   int32(0x10026),
	},
	386: {
		Fstart: int32(0x10028),
		Fend:   int32(0x1003a),
	},
	387: {
		Fstart: int32(0x1003c),
		Fend:   int32(0x1003d),
	},
	388: {
		Fstart: int32(0x1003f),
		Fend:   int32(0x1004d),
	},
	389: {
		Fstart: int32(0x10050),
		Fend:   int32(0x1005d),
	},
	390: {
		Fstart: int32(0x10080),
		Fend:   int32(0x100fa),
	},
	391: {
		Fstart: int32(0x10280),
		Fend:   int32(0x1029c),
	},
	392: {
		Fstart: int32(0x102a0),
		Fend:   int32(0x102d0),
	},
	393: {
		Fstart: int32(0x10300),
		Fend:   int32(0x1031f),
	},
	394: {
		Fstart: int32(0x1032d),
		Fend:   int32(0x10340),
	},
	395: {
		Fstart: int32(0x10342),
		Fend:   int32(0x10349),
	},
	396: {
		Fstart: int32(0x10350),
		Fend:   int32(0x10375),
	},
	397: {
		Fstart: int32(0x10380),
		Fend:   int32(0x1039d),
	},
	398: {
		Fstart: int32(0x103a0),
		Fend:   int32(0x103c3),
	},
	399: {
		Fstart: int32(0x103c8),
		Fend:   int32(0x103cf),
	},
	400: {
		Fstart: int32(0x10400),
		Fend:   int32(0x1049d),
	},
	401: {
		Fstart: int32(0x104b0),
		Fend:   int32(0x104d3),
	},
	402: {
		Fstart: int32(0x104d8),
		Fend:   int32(0x104fb),
	},
	403: {
		Fstart: int32(0x10500),
		Fend:   int32(0x10527),
	},
	404: {
		Fstart: int32(0x10530),
		Fend:   int32(0x10563),
	},
	405: {
		Fstart: int32(0x10570),
		Fend:   int32(0x1057a),
	},
	406: {
		Fstart: int32(0x1057c),
		Fend:   int32(0x1058a),
	},
	407: {
		Fstart: int32(0x1058c),
		Fend:   int32(0x10592),
	},
	408: {
		Fstart: int32(0x10594),
		Fend:   int32(0x10595),
	},
	409: {
		Fstart: int32(0x10597),
		Fend:   int32(0x105a1),
	},
	410: {
		Fstart: int32(0x105a3),
		Fend:   int32(0x105b1),
	},
	411: {
		Fstart: int32(0x105b3),
		Fend:   int32(0x105b9),
	},
	412: {
		Fstart: int32(0x105bb),
		Fend:   int32(0x105bc),
	},
	413: {
		Fstart: int32(0x105c0),
		Fend:   int32(0x105f3),
	},
	414: {
		Fstart: int32(0x10600),
		Fend:   int32(0x10736),
	},
	415: {
		Fstart: int32(0x10740),
		Fend:   int32(0x10755),
	},
	416: {
		Fstart: int32(0x10760),
		Fend:   int32(0x10767),
	},
	417: {
		Fstart: int32(0x10780),
		Fend:   int32(0x10785),
	},
	418: {
		Fstart: int32(0x10787),
		Fend:   int32(0x107b0),
	},
	419: {
		Fstart: int32(0x107b2),
		Fend:   int32(0x107ba),
	},
	420: {
		Fstart: int32(0x10800),
		Fend:   int32(0x10805),
	},
	421: {
		Fstart: int32(0x10808),
		Fend:   int32(0x10808),
	},
	422: {
		Fstart: int32(0x1080a),
		Fend:   int32(0x10835),
	},
	423: {
		Fstart: int32(0x10837),
		Fend:   int32(0x10838),
	},
	424: {
		Fstart: int32(0x1083c),
		Fend:   int32(0x1083c),
	},
	425: {
		Fstart: int32(0x1083f),
		Fend:   int32(0x10855),
	},
	426: {
		Fstart: int32(0x10860),
		Fend:   int32(0x10876),
	},
	427: {
		Fstart: int32(0x10880),
		Fend:   int32(0x1089e),
	},
	428: {
		Fstart: int32(0x108e0),
		Fend:   int32(0x108f2),
	},
	429: {
		Fstart: int32(0x108f4),
		Fend:   int32(0x108f5),
	},
	430: {
		Fstart: int32(0x10900),
		Fend:   int32(0x10915),
	},
	431: {
		Fstart: int32(0x10920),
		Fend:   int32(0x10939),
	},
	432: {
		Fstart: int32(0x10980),
		Fend:   int32(0x109b7),
	},
	433: {
		Fstart: int32(0x109be),
		Fend:   int32(0x109bf),
	},
	434: {
		Fstart: int32(0x10a00),
		Fend:   int32(0x10a00),
	},
	435: {
		Fstart: int32(0x10a10),
		Fend:   int32(0x10a13),
	},
	436: {
		Fstart: int32(0x10a15),
		Fend:   int32(0x10a17),
	},
	437: {
		Fstart: int32(0x10a19),
		Fend:   int32(0x10a35),
	},
	438: {
		Fstart: int32(0x10a60),
		Fend:   int32(0x10a7c),
	},
	439: {
		Fstart: int32(0x10a80),
		Fend:   int32(0x10a9c),
	},
	440: {
		Fstart: int32(0x10ac0),
		Fend:   int32(0x10ac7),
	},
	441: {
		Fstart: int32(0x10ac9),
		Fend:   int32(0x10ae4),
	},
	442: {
		Fstart: int32(0x10b00),
		Fend:   int32(0x10b35),
	},
	443: {
		Fstart: int32(0x10b40),
		Fend:   int32(0x10b55),
	},
	444: {
		Fstart: int32(0x10b60),
		Fend:   int32(0x10b72),
	},
	445: {
		Fstart: int32(0x10b80),
		Fend:   int32(0x10b91),
	},
	446: {
		Fstart: int32(0x10c00),
		Fend:   int32(0x10c48),
	},
	447: {
		Fstart: int32(0x10c80),
		Fend:   int32(0x10cb2),
	},
	448: {
		Fstart: int32(0x10cc0),
		Fend:   int32(0x10cf2),
	},
	449: {
		Fstart: int32(0x10d00),
		Fend:   int32(0x10d23),
	},
	450: {
		Fstart: int32(0x10d4a),
		Fend:   int32(0x10d65),
	},
	451: {
		Fstart: int32(0x10d6f),
		Fend:   int32(0x10d85),
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
		Fstart: int32(0x10ec2),
		Fend:   int32(0x10ec4),
	},
	455: {
		Fstart: int32(0x10f00),
		Fend:   int32(0x10f1c),
	},
	456: {
		Fstart: int32(0x10f27),
		Fend:   int32(0x10f27),
	},
	457: {
		Fstart: int32(0x10f30),
		Fend:   int32(0x10f45),
	},
	458: {
		Fstart: int32(0x10f70),
		Fend:   int32(0x10f81),
	},
	459: {
		Fstart: int32(0x10fb0),
		Fend:   int32(0x10fc4),
	},
	460: {
		Fstart: int32(0x10fe0),
		Fend:   int32(0x10ff6),
	},
	461: {
		Fstart: int32(0x11003),
		Fend:   int32(0x11037),
	},
	462: {
		Fstart: int32(0x11071),
		Fend:   int32(0x11072),
	},
	463: {
		Fstart: int32(0x11075),
		Fend:   int32(0x11075),
	},
	464: {
		Fstart: int32(0x11083),
		Fend:   int32(0x110af),
	},
	465: {
		Fstart: int32(0x110d0),
		Fend:   int32(0x110e8),
	},
	466: {
		Fstart: int32(0x11103),
		Fend:   int32(0x11126),
	},
	467: {
		Fstart: int32(0x11144),
		Fend:   int32(0x11144),
	},
	468: {
		Fstart: int32(0x11147),
		Fend:   int32(0x11147),
	},
	469: {
		Fstart: int32(0x11150),
		Fend:   int32(0x11172),
	},
	470: {
		Fstart: int32(0x11176),
		Fend:   int32(0x11176),
	},
	471: {
		Fstart: int32(0x11183),
		Fend:   int32(0x111b2),
	},
	472: {
		Fstart: int32(0x111c1),
		Fend:   int32(0x111c4),
	},
	473: {
		Fstart: int32(0x111da),
		Fend:   int32(0x111da),
	},
	474: {
		Fstart: int32(0x111dc),
		Fend:   int32(0x111dc),
	},
	475: {
		Fstart: int32(0x11200),
		Fend:   int32(0x11211),
	},
	476: {
		Fstart: int32(0x11213),
		Fend:   int32(0x1122b),
	},
	477: {
		Fstart: int32(0x1123f),
		Fend:   int32(0x11240),
	},
	478: {
		Fstart: int32(0x11280),
		Fend:   int32(0x11286),
	},
	479: {
		Fstart: int32(0x11288),
		Fend:   int32(0x11288),
	},
	480: {
		Fstart: int32(0x1128a),
		Fend:   int32(0x1128d),
	},
	481: {
		Fstart: int32(0x1128f),
		Fend:   int32(0x1129d),
	},
	482: {
		Fstart: int32(0x1129f),
		Fend:   int32(0x112a8),
	},
	483: {
		Fstart: int32(0x112b0),
		Fend:   int32(0x112de),
	},
	484: {
		Fstart: int32(0x11305),
		Fend:   int32(0x1130c),
	},
	485: {
		Fstart: int32(0x1130f),
		Fend:   int32(0x11310),
	},
	486: {
		Fstart: int32(0x11313),
		Fend:   int32(0x11328),
	},
	487: {
		Fstart: int32(0x1132a),
		Fend:   int32(0x11330),
	},
	488: {
		Fstart: int32(0x11332),
		Fend:   int32(0x11333),
	},
	489: {
		Fstart: int32(0x11335),
		Fend:   int32(0x11339),
	},
	490: {
		Fstart: int32(0x1133d),
		Fend:   int32(0x1133d),
	},
	491: {
		Fstart: int32(0x11350),
		Fend:   int32(0x11350),
	},
	492: {
		Fstart: int32(0x1135d),
		Fend:   int32(0x11361),
	},
	493: {
		Fstart: int32(0x11380),
		Fend:   int32(0x11389),
	},
	494: {
		Fstart: int32(0x1138b),
		Fend:   int32(0x1138b),
	},
	495: {
		Fstart: int32(0x1138e),
		Fend:   int32(0x1138e),
	},
	496: {
		Fstart: int32(0x11390),
		Fend:   int32(0x113b5),
	},
	497: {
		Fstart: int32(0x113b7),
		Fend:   int32(0x113b7),
	},
	498: {
		Fstart: int32(0x113d1),
		Fend:   int32(0x113d1),
	},
	499: {
		Fstart: int32(0x113d3),
		Fend:   int32(0x113d3),
	},
	500: {
		Fstart: int32(0x11400),
		Fend:   int32(0x11434),
	},
	501: {
		Fstart: int32(0x11447),
		Fend:   int32(0x1144a),
	},
	502: {
		Fstart: int32(0x1145f),
		Fend:   int32(0x11461),
	},
	503: {
		Fstart: int32(0x11480),
		Fend:   int32(0x114af),
	},
	504: {
		Fstart: int32(0x114c4),
		Fend:   int32(0x114c5),
	},
	505: {
		Fstart: int32(0x114c7),
		Fend:   int32(0x114c7),
	},
	506: {
		Fstart: int32(0x11580),
		Fend:   int32(0x115ae),
	},
	507: {
		Fstart: int32(0x115d8),
		Fend:   int32(0x115db),
	},
	508: {
		Fstart: int32(0x11600),
		Fend:   int32(0x1162f),
	},
	509: {
		Fstart: int32(0x11644),
		Fend:   int32(0x11644),
	},
	510: {
		Fstart: int32(0x11680),
		Fend:   int32(0x116aa),
	},
	511: {
		Fstart: int32(0x116b8),
		Fend:   int32(0x116b8),
	},
	512: {
		Fstart: int32(0x11700),
		Fend:   int32(0x1171a),
	},
	513: {
		Fstart: int32(0x11740),
		Fend:   int32(0x11746),
	},
	514: {
		Fstart: int32(0x11800),
		Fend:   int32(0x1182b),
	},
	515: {
		Fstart: int32(0x118a0),
		Fend:   int32(0x118df),
	},
	516: {
		Fstart: int32(0x118ff),
		Fend:   int32(0x11906),
	},
	517: {
		Fstart: int32(0x11909),
		Fend:   int32(0x11909),
	},
	518: {
		Fstart: int32(0x1190c),
		Fend:   int32(0x11913),
	},
	519: {
		Fstart: int32(0x11915),
		Fend:   int32(0x11916),
	},
	520: {
		Fstart: int32(0x11918),
		Fend:   int32(0x1192f),
	},
	521: {
		Fstart: int32(0x1193f),
		Fend:   int32(0x1193f),
	},
	522: {
		Fstart: int32(0x11941),
		Fend:   int32(0x11941),
	},
	523: {
		Fstart: int32(0x119a0),
		Fend:   int32(0x119a7),
	},
	524: {
		Fstart: int32(0x119aa),
		Fend:   int32(0x119d0),
	},
	525: {
		Fstart: int32(0x119e1),
		Fend:   int32(0x119e1),
	},
	526: {
		Fstart: int32(0x119e3),
		Fend:   int32(0x119e3),
	},
	527: {
		Fstart: int32(0x11a00),
		Fend:   int32(0x11a00),
	},
	528: {
		Fstart: int32(0x11a0b),
		Fend:   int32(0x11a32),
	},
	529: {
		Fstart: int32(0x11a3a),
		Fend:   int32(0x11a3a),
	},
	530: {
		Fstart: int32(0x11a50),
		Fend:   int32(0x11a50),
	},
	531: {
		Fstart: int32(0x11a5c),
		Fend:   int32(0x11a89),
	},
	532: {
		Fstart: int32(0x11a9d),
		Fend:   int32(0x11a9d),
	},
	533: {
		Fstart: int32(0x11ab0),
		Fend:   int32(0x11af8),
	},
	534: {
		Fstart: int32(0x11bc0),
		Fend:   int32(0x11be0),
	},
	535: {
		Fstart: int32(0x11c00),
		Fend:   int32(0x11c08),
	},
	536: {
		Fstart: int32(0x11c0a),
		Fend:   int32(0x11c2e),
	},
	537: {
		Fstart: int32(0x11c40),
		Fend:   int32(0x11c40),
	},
	538: {
		Fstart: int32(0x11c72),
		Fend:   int32(0x11c8f),
	},
	539: {
		Fstart: int32(0x11d00),
		Fend:   int32(0x11d06),
	},
	540: {
		Fstart: int32(0x11d08),
		Fend:   int32(0x11d09),
	},
	541: {
		Fstart: int32(0x11d0b),
		Fend:   int32(0x11d30),
	},
	542: {
		Fstart: int32(0x11d46),
		Fend:   int32(0x11d46),
	},
	543: {
		Fstart: int32(0x11d60),
		Fend:   int32(0x11d65),
	},
	544: {
		Fstart: int32(0x11d67),
		Fend:   int32(0x11d68),
	},
	545: {
		Fstart: int32(0x11d6a),
		Fend:   int32(0x11d89),
	},
	546: {
		Fstart: int32(0x11d98),
		Fend:   int32(0x11d98),
	},
	547: {
		Fstart: int32(0x11ee0),
		Fend:   int32(0x11ef2),
	},
	548: {
		Fstart: int32(0x11f02),
		Fend:   int32(0x11f02),
	},
	549: {
		Fstart: int32(0x11f04),
		Fend:   int32(0x11f10),
	},
	550: {
		Fstart: int32(0x11f12),
		Fend:   int32(0x11f33),
	},
	551: {
		Fstart: int32(0x11fb0),
		Fend:   int32(0x11fb0),
	},
	552: {
		Fstart: int32(0x12000),
		Fend:   int32(0x12399),
	},
	553: {
		Fstart: int32(0x12480),
		Fend:   int32(0x12543),
	},
	554: {
		Fstart: int32(0x12f90),
		Fend:   int32(0x12ff0),
	},
	555: {
		Fstart: int32(0x13000),
		Fend:   int32(0x1342f),
	},
	556: {
		Fstart: int32(0x13441),
		Fend:   int32(0x13446),
	},
	557: {
		Fstart: int32(0x13460),
		Fend:   int32(0x143fa),
	},
	558: {
		Fstart: int32(0x14400),
		Fend:   int32(0x14646),
	},
	559: {
		Fstart: int32(0x16100),
		Fend:   int32(0x1611d),
	},
	560: {
		Fstart: int32(0x16800),
		Fend:   int32(0x16a38),
	},
	561: {
		Fstart: int32(0x16a40),
		Fend:   int32(0x16a5e),
	},
	562: {
		Fstart: int32(0x16a70),
		Fend:   int32(0x16abe),
	},
	563: {
		Fstart: int32(0x16ad0),
		Fend:   int32(0x16aed),
	},
	564: {
		Fstart: int32(0x16b00),
		Fend:   int32(0x16b2f),
	},
	565: {
		Fstart: int32(0x16b40),
		Fend:   int32(0x16b43),
	},
	566: {
		Fstart: int32(0x16b63),
		Fend:   int32(0x16b77),
	},
	567: {
		Fstart: int32(0x16b7d),
		Fend:   int32(0x16b8f),
	},
	568: {
		Fstart: int32(0x16d40),
		Fend:   int32(0x16d6c),
	},
	569: {
		Fstart: int32(0x16e40),
		Fend:   int32(0x16e7f),
	},
	570: {
		Fstart: int32(0x16f00),
		Fend:   int32(0x16f4a),
	},
	571: {
		Fstart: int32(0x16f50),
		Fend:   int32(0x16f50),
	},
	572: {
		Fstart: int32(0x16f93),
		Fend:   int32(0x16f9f),
	},
	573: {
		Fstart: int32(0x16fe0),
		Fend:   int32(0x16fe1),
	},
	574: {
		Fstart: int32(0x16fe3),
		Fend:   int32(0x16fe3),
	},
	575: {
		Fstart: int32(0x17000),
		Fend:   int32(0x187f7),
	},
	576: {
		Fstart: int32(0x18800),
		Fend:   int32(0x18cd5),
	},
	577: {
		Fstart: int32(0x18cff),
		Fend:   int32(0x18d08),
	},
	578: {
		Fstart: int32(0x1aff0),
		Fend:   int32(0x1aff3),
	},
	579: {
		Fstart: int32(0x1aff5),
		Fend:   int32(0x1affb),
	},
	580: {
		Fstart: int32(0x1affd),
		Fend:   int32(0x1affe),
	},
	581: {
		Fstart: int32(0x1b000),
		Fend:   int32(0x1b122),
	},
	582: {
		Fstart: int32(0x1b132),
		Fend:   int32(0x1b132),
	},
	583: {
		Fstart: int32(0x1b150),
		Fend:   int32(0x1b152),
	},
	584: {
		Fstart: int32(0x1b155),
		Fend:   int32(0x1b155),
	},
	585: {
		Fstart: int32(0x1b164),
		Fend:   int32(0x1b167),
	},
	586: {
		Fstart: int32(0x1b170),
		Fend:   int32(0x1b2fb),
	},
	587: {
		Fstart: int32(0x1bc00),
		Fend:   int32(0x1bc6a),
	},
	588: {
		Fstart: int32(0x1bc70),
		Fend:   int32(0x1bc7c),
	},
	589: {
		Fstart: int32(0x1bc80),
		Fend:   int32(0x1bc88),
	},
	590: {
		Fstart: int32(0x1bc90),
		Fend:   int32(0x1bc99),
	},
	591: {
		Fstart: int32(0x1d400),
		Fend:   int32(0x1d454),
	},
	592: {
		Fstart: int32(0x1d456),
		Fend:   int32(0x1d49c),
	},
	593: {
		Fstart: int32(0x1d49e),
		Fend:   int32(0x1d49f),
	},
	594: {
		Fstart: int32(0x1d4a2),
		Fend:   int32(0x1d4a2),
	},
	595: {
		Fstart: int32(0x1d4a5),
		Fend:   int32(0x1d4a6),
	},
	596: {
		Fstart: int32(0x1d4a9),
		Fend:   int32(0x1d4ac),
	},
	597: {
		Fstart: int32(0x1d4ae),
		Fend:   int32(0x1d4b9),
	},
	598: {
		Fstart: int32(0x1d4bb),
		Fend:   int32(0x1d4bb),
	},
	599: {
		Fstart: int32(0x1d4bd),
		Fend:   int32(0x1d4c3),
	},
	600: {
		Fstart: int32(0x1d4c5),
		Fend:   int32(0x1d505),
	},
	601: {
		Fstart: int32(0x1d507),
		Fend:   int32(0x1d50a),
	},
	602: {
		Fstart: int32(0x1d50d),
		Fend:   int32(0x1d514),
	},
	603: {
		Fstart: int32(0x1d516),
		Fend:   int32(0x1d51c),
	},
	604: {
		Fstart: int32(0x1d51e),
		Fend:   int32(0x1d539),
	},
	605: {
		Fstart: int32(0x1d53b),
		Fend:   int32(0x1d53e),
	},
	606: {
		Fstart: int32(0x1d540),
		Fend:   int32(0x1d544),
	},
	607: {
		Fstart: int32(0x1d546),
		Fend:   int32(0x1d546),
	},
	608: {
		Fstart: int32(0x1d54a),
		Fend:   int32(0x1d550),
	},
	609: {
		Fstart: int32(0x1d552),
		Fend:   int32(0x1d6a5),
	},
	610: {
		Fstart: int32(0x1d6a8),
		Fend:   int32(0x1d6c0),
	},
	611: {
		Fstart: int32(0x1d6c2),
		Fend:   int32(0x1d6da),
	},
	612: {
		Fstart: int32(0x1d6dc),
		Fend:   int32(0x1d6fa),
	},
	613: {
		Fstart: int32(0x1d6fc),
		Fend:   int32(0x1d714),
	},
	614: {
		Fstart: int32(0x1d716),
		Fend:   int32(0x1d734),
	},
	615: {
		Fstart: int32(0x1d736),
		Fend:   int32(0x1d74e),
	},
	616: {
		Fstart: int32(0x1d750),
		Fend:   int32(0x1d76e),
	},
	617: {
		Fstart: int32(0x1d770),
		Fend:   int32(0x1d788),
	},
	618: {
		Fstart: int32(0x1d78a),
		Fend:   int32(0x1d7a8),
	},
	619: {
		Fstart: int32(0x1d7aa),
		Fend:   int32(0x1d7c2),
	},
	620: {
		Fstart: int32(0x1d7c4),
		Fend:   int32(0x1d7cb),
	},
	621: {
		Fstart: int32(0x1df00),
		Fend:   int32(0x1df1e),
	},
	622: {
		Fstart: int32(0x1df25),
		Fend:   int32(0x1df2a),
	},
	623: {
		Fstart: int32(0x1e030),
		Fend:   int32(0x1e06d),
	},
	624: {
		Fstart: int32(0x1e100),
		Fend:   int32(0x1e12c),
	},
	625: {
		Fstart: int32(0x1e137),
		Fend:   int32(0x1e13d),
	},
	626: {
		Fstart: int32(0x1e14e),
		Fend:   int32(0x1e14e),
	},
	627: {
		Fstart: int32(0x1e290),
		Fend:   int32(0x1e2ad),
	},
	628: {
		Fstart: int32(0x1e2c0),
		Fend:   int32(0x1e2eb),
	},
	629: {
		Fstart: int32(0x1e4d0),
		Fend:   int32(0x1e4eb),
	},
	630: {
		Fstart: int32(0x1e5d0),
		Fend:   int32(0x1e5ed),
	},
	631: {
		Fstart: int32(0x1e5f0),
		Fend:   int32(0x1e5f0),
	},
	632: {
		Fstart: int32(0x1e7e0),
		Fend:   int32(0x1e7e6),
	},
	633: {
		Fstart: int32(0x1e7e8),
		Fend:   int32(0x1e7eb),
	},
	634: {
		Fstart: int32(0x1e7ed),
		Fend:   int32(0x1e7ee),
	},
	635: {
		Fstart: int32(0x1e7f0),
		Fend:   int32(0x1e7fe),
	},
	636: {
		Fstart: int32(0x1e800),
		Fend:   int32(0x1e8c4),
	},
	637: {
		Fstart: int32(0x1e900),
		Fend:   int32(0x1e943),
	},
	638: {
		Fstart: int32(0x1e94b),
		Fend:   int32(0x1e94b),
	},
	639: {
		Fstart: int32(0x1ee00),
		Fend:   int32(0x1ee03),
	},
	640: {
		Fstart: int32(0x1ee05),
		Fend:   int32(0x1ee1f),
	},
	641: {
		Fstart: int32(0x1ee21),
		Fend:   int32(0x1ee22),
	},
	642: {
		Fstart: int32(0x1ee24),
		Fend:   int32(0x1ee24),
	},
	643: {
		Fstart: int32(0x1ee27),
		Fend:   int32(0x1ee27),
	},
	644: {
		Fstart: int32(0x1ee29),
		Fend:   int32(0x1ee32),
	},
	645: {
		Fstart: int32(0x1ee34),
		Fend:   int32(0x1ee37),
	},
	646: {
		Fstart: int32(0x1ee39),
		Fend:   int32(0x1ee39),
	},
	647: {
		Fstart: int32(0x1ee3b),
		Fend:   int32(0x1ee3b),
	},
	648: {
		Fstart: int32(0x1ee42),
		Fend:   int32(0x1ee42),
	},
	649: {
		Fstart: int32(0x1ee47),
		Fend:   int32(0x1ee47),
	},
	650: {
		Fstart: int32(0x1ee49),
		Fend:   int32(0x1ee49),
	},
	651: {
		Fstart: int32(0x1ee4b),
		Fend:   int32(0x1ee4b),
	},
	652: {
		Fstart: int32(0x1ee4d),
		Fend:   int32(0x1ee4f),
	},
	653: {
		Fstart: int32(0x1ee51),
		Fend:   int32(0x1ee52),
	},
	654: {
		Fstart: int32(0x1ee54),
		Fend:   int32(0x1ee54),
	},
	655: {
		Fstart: int32(0x1ee57),
		Fend:   int32(0x1ee57),
	},
	656: {
		Fstart: int32(0x1ee59),
		Fend:   int32(0x1ee59),
	},
	657: {
		Fstart: int32(0x1ee5b),
		Fend:   int32(0x1ee5b),
	},
	658: {
		Fstart: int32(0x1ee5d),
		Fend:   int32(0x1ee5d),
	},
	659: {
		Fstart: int32(0x1ee5f),
		Fend:   int32(0x1ee5f),
	},
	660: {
		Fstart: int32(0x1ee61),
		Fend:   int32(0x1ee62),
	},
	661: {
		Fstart: int32(0x1ee64),
		Fend:   int32(0x1ee64),
	},
	662: {
		Fstart: int32(0x1ee67),
		Fend:   int32(0x1ee6a),
	},
	663: {
		Fstart: int32(0x1ee6c),
		Fend:   int32(0x1ee72),
	},
	664: {
		Fstart: int32(0x1ee74),
		Fend:   int32(0x1ee77),
	},
	665: {
		Fstart: int32(0x1ee79),
		Fend:   int32(0x1ee7c),
	},
	666: {
		Fstart: int32(0x1ee7e),
		Fend:   int32(0x1ee7e),
	},
	667: {
		Fstart: int32(0x1ee80),
		Fend:   int32(0x1ee89),
	},
	668: {
		Fstart: int32(0x1ee8b),
		Fend:   int32(0x1ee9b),
	},
	669: {
		Fstart: int32(0x1eea1),
		Fend:   int32(0x1eea3),
	},
	670: {
		Fstart: int32(0x1eea5),
		Fend:   int32(0x1eea9),
	},
	671: {
		Fstart: int32(0x1eeab),
		Fend:   int32(0x1eebb),
	},
	672: {
		Fstart: int32(0x20000),
		Fend:   int32(0x2a6df),
	},
	673: {
		Fstart: int32(0x2a700),
		Fend:   int32(0x2b739),
	},
	674: {
		Fstart: int32(0x2b740),
		Fend:   int32(0x2b81d),
	},
	675: {
		Fstart: int32(0x2b820),
		Fend:   int32(0x2cea1),
	},
	676: {
		Fstart: int32(0x2ceb0),
		Fend:   int32(0x2ebe0),
	},
	677: {
		Fstart: int32(0x2ebf0),
		Fend:   int32(0x2ee5d),
	},
	678: {
		Fstart: int32(0x2f800),
		Fend:   int32(0x2fa1d),
	},
	679: {
		Fstart: int32(0x30000),
		Fend:   int32(0x3134a),
	},
	680: {
		Fstart: int32(0x31350),
		Fend:   int32(0x323af),
	},
}

var aux_sym_string_token2_character_set_1 = [680]TSCharacterRange{
	0: {
		Fstart: int32('\''),
		Fend:   int32('\''),
	},
	1: {
		Fstart: int32(','),
		Fend:   int32('-'),
	},
	2: {
		Fstart: int32(':'),
		Fend:   int32(':'),
	},
	3: {
		Fstart: int32('A'),
		Fend:   int32('Z'),
	},
	4: {
		Fstart: int32('a'),
		Fend:   int32('z'),
	},
	5: {
		Fstart: int32(0xaa),
		Fend:   int32(0xaa),
	},
	6: {
		Fstart: int32(0xb5),
		Fend:   int32(0xb5),
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
		Fstart: int32(0x370),
		Fend:   int32(0x374),
	},
	16: {
		Fstart: int32(0x376),
		Fend:   int32(0x377),
	},
	17: {
		Fstart: int32(0x37a),
		Fend:   int32(0x37d),
	},
	18: {
		Fstart: int32(0x37f),
		Fend:   int32(0x37f),
	},
	19: {
		Fstart: int32(0x386),
		Fend:   int32(0x386),
	},
	20: {
		Fstart: int32(0x388),
		Fend:   int32(0x38a),
	},
	21: {
		Fstart: int32(0x38c),
		Fend:   int32(0x38c),
	},
	22: {
		Fstart: int32(0x38e),
		Fend:   int32(0x3a1),
	},
	23: {
		Fstart: int32(0x3a3),
		Fend:   int32(0x3f5),
	},
	24: {
		Fstart: int32(0x3f7),
		Fend:   int32(0x481),
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
		Fstart: int32(0x5d0),
		Fend:   int32(0x5ea),
	},
	30: {
		Fstart: int32(0x5ef),
		Fend:   int32(0x5f2),
	},
	31: {
		Fstart: int32(0x620),
		Fend:   int32(0x64a),
	},
	32: {
		Fstart: int32(0x66e),
		Fend:   int32(0x66f),
	},
	33: {
		Fstart: int32(0x671),
		Fend:   int32(0x6d3),
	},
	34: {
		Fstart: int32(0x6d5),
		Fend:   int32(0x6d5),
	},
	35: {
		Fstart: int32(0x6e5),
		Fend:   int32(0x6e6),
	},
	36: {
		Fstart: int32(0x6ee),
		Fend:   int32(0x6ef),
	},
	37: {
		Fstart: int32(0x6fa),
		Fend:   int32(0x6fc),
	},
	38: {
		Fstart: int32(0x6ff),
		Fend:   int32(0x6ff),
	},
	39: {
		Fstart: int32(0x710),
		Fend:   int32(0x710),
	},
	40: {
		Fstart: int32(0x712),
		Fend:   int32(0x72f),
	},
	41: {
		Fstart: int32(0x74d),
		Fend:   int32(0x7a5),
	},
	42: {
		Fstart: int32(0x7b1),
		Fend:   int32(0x7b1),
	},
	43: {
		Fstart: int32(0x7ca),
		Fend:   int32(0x7ea),
	},
	44: {
		Fstart: int32(0x7f4),
		Fend:   int32(0x7f5),
	},
	45: {
		Fstart: int32(0x7fa),
		Fend:   int32(0x7fa),
	},
	46: {
		Fstart: int32(0x800),
		Fend:   int32(0x815),
	},
	47: {
		Fstart: int32(0x81a),
		Fend:   int32(0x81a),
	},
	48: {
		Fstart: int32(0x824),
		Fend:   int32(0x824),
	},
	49: {
		Fstart: int32(0x828),
		Fend:   int32(0x828),
	},
	50: {
		Fstart: int32(0x840),
		Fend:   int32(0x858),
	},
	51: {
		Fstart: int32(0x860),
		Fend:   int32(0x86a),
	},
	52: {
		Fstart: int32(0x870),
		Fend:   int32(0x887),
	},
	53: {
		Fstart: int32(0x889),
		Fend:   int32(0x88e),
	},
	54: {
		Fstart: int32(0x8a0),
		Fend:   int32(0x8c9),
	},
	55: {
		Fstart: int32(0x904),
		Fend:   int32(0x939),
	},
	56: {
		Fstart: int32(0x93d),
		Fend:   int32(0x93d),
	},
	57: {
		Fstart: int32(0x950),
		Fend:   int32(0x950),
	},
	58: {
		Fstart: int32(0x958),
		Fend:   int32(0x961),
	},
	59: {
		Fstart: int32(0x971),
		Fend:   int32(0x980),
	},
	60: {
		Fstart: int32(0x985),
		Fend:   int32(0x98c),
	},
	61: {
		Fstart: int32(0x98f),
		Fend:   int32(0x990),
	},
	62: {
		Fstart: int32(0x993),
		Fend:   int32(0x9a8),
	},
	63: {
		Fstart: int32(0x9aa),
		Fend:   int32(0x9b0),
	},
	64: {
		Fstart: int32(0x9b2),
		Fend:   int32(0x9b2),
	},
	65: {
		Fstart: int32(0x9b6),
		Fend:   int32(0x9b9),
	},
	66: {
		Fstart: int32(0x9bd),
		Fend:   int32(0x9bd),
	},
	67: {
		Fstart: int32(0x9ce),
		Fend:   int32(0x9ce),
	},
	68: {
		Fstart: int32(0x9dc),
		Fend:   int32(0x9dd),
	},
	69: {
		Fstart: int32(0x9df),
		Fend:   int32(0x9e1),
	},
	70: {
		Fstart: int32(0x9f0),
		Fend:   int32(0x9f1),
	},
	71: {
		Fstart: int32(0x9fc),
		Fend:   int32(0x9fc),
	},
	72: {
		Fstart: int32(0xa05),
		Fend:   int32(0xa0a),
	},
	73: {
		Fstart: int32(0xa0f),
		Fend:   int32(0xa10),
	},
	74: {
		Fstart: int32(0xa13),
		Fend:   int32(0xa28),
	},
	75: {
		Fstart: int32(0xa2a),
		Fend:   int32(0xa30),
	},
	76: {
		Fstart: int32(0xa32),
		Fend:   int32(0xa33),
	},
	77: {
		Fstart: int32(0xa35),
		Fend:   int32(0xa36),
	},
	78: {
		Fstart: int32(0xa38),
		Fend:   int32(0xa39),
	},
	79: {
		Fstart: int32(0xa59),
		Fend:   int32(0xa5c),
	},
	80: {
		Fstart: int32(0xa5e),
		Fend:   int32(0xa5e),
	},
	81: {
		Fstart: int32(0xa72),
		Fend:   int32(0xa74),
	},
	82: {
		Fstart: int32(0xa85),
		Fend:   int32(0xa8d),
	},
	83: {
		Fstart: int32(0xa8f),
		Fend:   int32(0xa91),
	},
	84: {
		Fstart: int32(0xa93),
		Fend:   int32(0xaa8),
	},
	85: {
		Fstart: int32(0xaaa),
		Fend:   int32(0xab0),
	},
	86: {
		Fstart: int32(0xab2),
		Fend:   int32(0xab3),
	},
	87: {
		Fstart: int32(0xab5),
		Fend:   int32(0xab9),
	},
	88: {
		Fstart: int32(0xabd),
		Fend:   int32(0xabd),
	},
	89: {
		Fstart: int32(0xad0),
		Fend:   int32(0xad0),
	},
	90: {
		Fstart: int32(0xae0),
		Fend:   int32(0xae1),
	},
	91: {
		Fstart: int32(0xaf9),
		Fend:   int32(0xaf9),
	},
	92: {
		Fstart: int32(0xb05),
		Fend:   int32(0xb0c),
	},
	93: {
		Fstart: int32(0xb0f),
		Fend:   int32(0xb10),
	},
	94: {
		Fstart: int32(0xb13),
		Fend:   int32(0xb28),
	},
	95: {
		Fstart: int32(0xb2a),
		Fend:   int32(0xb30),
	},
	96: {
		Fstart: int32(0xb32),
		Fend:   int32(0xb33),
	},
	97: {
		Fstart: int32(0xb35),
		Fend:   int32(0xb39),
	},
	98: {
		Fstart: int32(0xb3d),
		Fend:   int32(0xb3d),
	},
	99: {
		Fstart: int32(0xb5c),
		Fend:   int32(0xb5d),
	},
	100: {
		Fstart: int32(0xb5f),
		Fend:   int32(0xb61),
	},
	101: {
		Fstart: int32(0xb71),
		Fend:   int32(0xb71),
	},
	102: {
		Fstart: int32(0xb83),
		Fend:   int32(0xb83),
	},
	103: {
		Fstart: int32(0xb85),
		Fend:   int32(0xb8a),
	},
	104: {
		Fstart: int32(0xb8e),
		Fend:   int32(0xb90),
	},
	105: {
		Fstart: int32(0xb92),
		Fend:   int32(0xb95),
	},
	106: {
		Fstart: int32(0xb99),
		Fend:   int32(0xb9a),
	},
	107: {
		Fstart: int32(0xb9c),
		Fend:   int32(0xb9c),
	},
	108: {
		Fstart: int32(0xb9e),
		Fend:   int32(0xb9f),
	},
	109: {
		Fstart: int32(0xba3),
		Fend:   int32(0xba4),
	},
	110: {
		Fstart: int32(0xba8),
		Fend:   int32(0xbaa),
	},
	111: {
		Fstart: int32(0xbae),
		Fend:   int32(0xbb9),
	},
	112: {
		Fstart: int32(0xbd0),
		Fend:   int32(0xbd0),
	},
	113: {
		Fstart: int32(0xc05),
		Fend:   int32(0xc0c),
	},
	114: {
		Fstart: int32(0xc0e),
		Fend:   int32(0xc10),
	},
	115: {
		Fstart: int32(0xc12),
		Fend:   int32(0xc28),
	},
	116: {
		Fstart: int32(0xc2a),
		Fend:   int32(0xc39),
	},
	117: {
		Fstart: int32(0xc3d),
		Fend:   int32(0xc3d),
	},
	118: {
		Fstart: int32(0xc58),
		Fend:   int32(0xc5a),
	},
	119: {
		Fstart: int32(0xc5d),
		Fend:   int32(0xc5d),
	},
	120: {
		Fstart: int32(0xc60),
		Fend:   int32(0xc61),
	},
	121: {
		Fstart: int32(0xc80),
		Fend:   int32(0xc80),
	},
	122: {
		Fstart: int32(0xc85),
		Fend:   int32(0xc8c),
	},
	123: {
		Fstart: int32(0xc8e),
		Fend:   int32(0xc90),
	},
	124: {
		Fstart: int32(0xc92),
		Fend:   int32(0xca8),
	},
	125: {
		Fstart: int32(0xcaa),
		Fend:   int32(0xcb3),
	},
	126: {
		Fstart: int32(0xcb5),
		Fend:   int32(0xcb9),
	},
	127: {
		Fstart: int32(0xcbd),
		Fend:   int32(0xcbd),
	},
	128: {
		Fstart: int32(0xcdd),
		Fend:   int32(0xcde),
	},
	129: {
		Fstart: int32(0xce0),
		Fend:   int32(0xce1),
	},
	130: {
		Fstart: int32(0xcf1),
		Fend:   int32(0xcf2),
	},
	131: {
		Fstart: int32(0xd04),
		Fend:   int32(0xd0c),
	},
	132: {
		Fstart: int32(0xd0e),
		Fend:   int32(0xd10),
	},
	133: {
		Fstart: int32(0xd12),
		Fend:   int32(0xd3a),
	},
	134: {
		Fstart: int32(0xd3d),
		Fend:   int32(0xd3d),
	},
	135: {
		Fstart: int32(0xd4e),
		Fend:   int32(0xd4e),
	},
	136: {
		Fstart: int32(0xd54),
		Fend:   int32(0xd56),
	},
	137: {
		Fstart: int32(0xd5f),
		Fend:   int32(0xd61),
	},
	138: {
		Fstart: int32(0xd7a),
		Fend:   int32(0xd7f),
	},
	139: {
		Fstart: int32(0xd85),
		Fend:   int32(0xd96),
	},
	140: {
		Fstart: int32(0xd9a),
		Fend:   int32(0xdb1),
	},
	141: {
		Fstart: int32(0xdb3),
		Fend:   int32(0xdbb),
	},
	142: {
		Fstart: int32(0xdbd),
		Fend:   int32(0xdbd),
	},
	143: {
		Fstart: int32(0xdc0),
		Fend:   int32(0xdc6),
	},
	144: {
		Fstart: int32(0xe01),
		Fend:   int32(0xe30),
	},
	145: {
		Fstart: int32(0xe32),
		Fend:   int32(0xe33),
	},
	146: {
		Fstart: int32(0xe40),
		Fend:   int32(0xe46),
	},
	147: {
		Fstart: int32(0xe81),
		Fend:   int32(0xe82),
	},
	148: {
		Fstart: int32(0xe84),
		Fend:   int32(0xe84),
	},
	149: {
		Fstart: int32(0xe86),
		Fend:   int32(0xe8a),
	},
	150: {
		Fstart: int32(0xe8c),
		Fend:   int32(0xea3),
	},
	151: {
		Fstart: int32(0xea5),
		Fend:   int32(0xea5),
	},
	152: {
		Fstart: int32(0xea7),
		Fend:   int32(0xeb0),
	},
	153: {
		Fstart: int32(0xeb2),
		Fend:   int32(0xeb3),
	},
	154: {
		Fstart: int32(0xebd),
		Fend:   int32(0xebd),
	},
	155: {
		Fstart: int32(0xec0),
		Fend:   int32(0xec4),
	},
	156: {
		Fstart: int32(0xec6),
		Fend:   int32(0xec6),
	},
	157: {
		Fstart: int32(0xedc),
		Fend:   int32(0xedf),
	},
	158: {
		Fstart: int32(0xf00),
		Fend:   int32(0xf00),
	},
	159: {
		Fstart: int32(0xf40),
		Fend:   int32(0xf47),
	},
	160: {
		Fstart: int32(0xf49),
		Fend:   int32(0xf6c),
	},
	161: {
		Fstart: int32(0xf88),
		Fend:   int32(0xf8c),
	},
	162: {
		Fstart: int32(0x1000),
		Fend:   int32(0x102a),
	},
	163: {
		Fstart: int32(0x103f),
		Fend:   int32(0x103f),
	},
	164: {
		Fstart: int32(0x1050),
		Fend:   int32(0x1055),
	},
	165: {
		Fstart: int32(0x105a),
		Fend:   int32(0x105d),
	},
	166: {
		Fstart: int32(0x1061),
		Fend:   int32(0x1061),
	},
	167: {
		Fstart: int32(0x1065),
		Fend:   int32(0x1066),
	},
	168: {
		Fstart: int32(0x106e),
		Fend:   int32(0x1070),
	},
	169: {
		Fstart: int32(0x1075),
		Fend:   int32(0x1081),
	},
	170: {
		Fstart: int32(0x108e),
		Fend:   int32(0x108e),
	},
	171: {
		Fstart: int32(0x10a0),
		Fend:   int32(0x10c5),
	},
	172: {
		Fstart: int32(0x10c7),
		Fend:   int32(0x10c7),
	},
	173: {
		Fstart: int32(0x10cd),
		Fend:   int32(0x10cd),
	},
	174: {
		Fstart: int32(0x10d0),
		Fend:   int32(0x10fa),
	},
	175: {
		Fstart: int32(0x10fc),
		Fend:   int32(0x1248),
	},
	176: {
		Fstart: int32(0x124a),
		Fend:   int32(0x124d),
	},
	177: {
		Fstart: int32(0x1250),
		Fend:   int32(0x1256),
	},
	178: {
		Fstart: int32(0x1258),
		Fend:   int32(0x1258),
	},
	179: {
		Fstart: int32(0x125a),
		Fend:   int32(0x125d),
	},
	180: {
		Fstart: int32(0x1260),
		Fend:   int32(0x1288),
	},
	181: {
		Fstart: int32(0x128a),
		Fend:   int32(0x128d),
	},
	182: {
		Fstart: int32(0x1290),
		Fend:   int32(0x12b0),
	},
	183: {
		Fstart: int32(0x12b2),
		Fend:   int32(0x12b5),
	},
	184: {
		Fstart: int32(0x12b8),
		Fend:   int32(0x12be),
	},
	185: {
		Fstart: int32(0x12c0),
		Fend:   int32(0x12c0),
	},
	186: {
		Fstart: int32(0x12c2),
		Fend:   int32(0x12c5),
	},
	187: {
		Fstart: int32(0x12c8),
		Fend:   int32(0x12d6),
	},
	188: {
		Fstart: int32(0x12d8),
		Fend:   int32(0x1310),
	},
	189: {
		Fstart: int32(0x1312),
		Fend:   int32(0x1315),
	},
	190: {
		Fstart: int32(0x1318),
		Fend:   int32(0x135a),
	},
	191: {
		Fstart: int32(0x1380),
		Fend:   int32(0x138f),
	},
	192: {
		Fstart: int32(0x13a0),
		Fend:   int32(0x13f5),
	},
	193: {
		Fstart: int32(0x13f8),
		Fend:   int32(0x13fd),
	},
	194: {
		Fstart: int32(0x1401),
		Fend:   int32(0x166c),
	},
	195: {
		Fstart: int32(0x166f),
		Fend:   int32(0x167f),
	},
	196: {
		Fstart: int32(0x1681),
		Fend:   int32(0x169a),
	},
	197: {
		Fstart: int32(0x16a0),
		Fend:   int32(0x16ea),
	},
	198: {
		Fstart: int32(0x16f1),
		Fend:   int32(0x16f8),
	},
	199: {
		Fstart: int32(0x1700),
		Fend:   int32(0x1711),
	},
	200: {
		Fstart: int32(0x171f),
		Fend:   int32(0x1731),
	},
	201: {
		Fstart: int32(0x1740),
		Fend:   int32(0x1751),
	},
	202: {
		Fstart: int32(0x1760),
		Fend:   int32(0x176c),
	},
	203: {
		Fstart: int32(0x176e),
		Fend:   int32(0x1770),
	},
	204: {
		Fstart: int32(0x1780),
		Fend:   int32(0x17b3),
	},
	205: {
		Fstart: int32(0x17d7),
		Fend:   int32(0x17d7),
	},
	206: {
		Fstart: int32(0x17dc),
		Fend:   int32(0x17dc),
	},
	207: {
		Fstart: int32(0x1820),
		Fend:   int32(0x1878),
	},
	208: {
		Fstart: int32(0x1880),
		Fend:   int32(0x1884),
	},
	209: {
		Fstart: int32(0x1887),
		Fend:   int32(0x18a8),
	},
	210: {
		Fstart: int32(0x18aa),
		Fend:   int32(0x18aa),
	},
	211: {
		Fstart: int32(0x18b0),
		Fend:   int32(0x18f5),
	},
	212: {
		Fstart: int32(0x1900),
		Fend:   int32(0x191e),
	},
	213: {
		Fstart: int32(0x1950),
		Fend:   int32(0x196d),
	},
	214: {
		Fstart: int32(0x1970),
		Fend:   int32(0x1974),
	},
	215: {
		Fstart: int32(0x1980),
		Fend:   int32(0x19ab),
	},
	216: {
		Fstart: int32(0x19b0),
		Fend:   int32(0x19c9),
	},
	217: {
		Fstart: int32(0x1a00),
		Fend:   int32(0x1a16),
	},
	218: {
		Fstart: int32(0x1a20),
		Fend:   int32(0x1a54),
	},
	219: {
		Fstart: int32(0x1aa7),
		Fend:   int32(0x1aa7),
	},
	220: {
		Fstart: int32(0x1b05),
		Fend:   int32(0x1b33),
	},
	221: {
		Fstart: int32(0x1b45),
		Fend:   int32(0x1b4c),
	},
	222: {
		Fstart: int32(0x1b83),
		Fend:   int32(0x1ba0),
	},
	223: {
		Fstart: int32(0x1bae),
		Fend:   int32(0x1baf),
	},
	224: {
		Fstart: int32(0x1bba),
		Fend:   int32(0x1be5),
	},
	225: {
		Fstart: int32(0x1c00),
		Fend:   int32(0x1c23),
	},
	226: {
		Fstart: int32(0x1c4d),
		Fend:   int32(0x1c4f),
	},
	227: {
		Fstart: int32(0x1c5a),
		Fend:   int32(0x1c7d),
	},
	228: {
		Fstart: int32(0x1c80),
		Fend:   int32(0x1c8a),
	},
	229: {
		Fstart: int32(0x1c90),
		Fend:   int32(0x1cba),
	},
	230: {
		Fstart: int32(0x1cbd),
		Fend:   int32(0x1cbf),
	},
	231: {
		Fstart: int32(0x1ce9),
		Fend:   int32(0x1cec),
	},
	232: {
		Fstart: int32(0x1cee),
		Fend:   int32(0x1cf3),
	},
	233: {
		Fstart: int32(0x1cf5),
		Fend:   int32(0x1cf6),
	},
	234: {
		Fstart: int32(0x1cfa),
		Fend:   int32(0x1cfa),
	},
	235: {
		Fstart: int32(0x1d00),
		Fend:   int32(0x1dbf),
	},
	236: {
		Fstart: int32(0x1e00),
		Fend:   int32(0x1f15),
	},
	237: {
		Fstart: int32(0x1f18),
		Fend:   int32(0x1f1d),
	},
	238: {
		Fstart: int32(0x1f20),
		Fend:   int32(0x1f45),
	},
	239: {
		Fstart: int32(0x1f48),
		Fend:   int32(0x1f4d),
	},
	240: {
		Fstart: int32(0x1f50),
		Fend:   int32(0x1f57),
	},
	241: {
		Fstart: int32(0x1f59),
		Fend:   int32(0x1f59),
	},
	242: {
		Fstart: int32(0x1f5b),
		Fend:   int32(0x1f5b),
	},
	243: {
		Fstart: int32(0x1f5d),
		Fend:   int32(0x1f5d),
	},
	244: {
		Fstart: int32(0x1f5f),
		Fend:   int32(0x1f7d),
	},
	245: {
		Fstart: int32(0x1f80),
		Fend:   int32(0x1fb4),
	},
	246: {
		Fstart: int32(0x1fb6),
		Fend:   int32(0x1fbc),
	},
	247: {
		Fstart: int32(0x1fbe),
		Fend:   int32(0x1fbe),
	},
	248: {
		Fstart: int32(0x1fc2),
		Fend:   int32(0x1fc4),
	},
	249: {
		Fstart: int32(0x1fc6),
		Fend:   int32(0x1fcc),
	},
	250: {
		Fstart: int32(0x1fd0),
		Fend:   int32(0x1fd3),
	},
	251: {
		Fstart: int32(0x1fd6),
		Fend:   int32(0x1fdb),
	},
	252: {
		Fstart: int32(0x1fe0),
		Fend:   int32(0x1fec),
	},
	253: {
		Fstart: int32(0x1ff2),
		Fend:   int32(0x1ff4),
	},
	254: {
		Fstart: int32(0x1ff6),
		Fend:   int32(0x1ffc),
	},
	255: {
		Fstart: int32(0x2071),
		Fend:   int32(0x2071),
	},
	256: {
		Fstart: int32(0x207f),
		Fend:   int32(0x207f),
	},
	257: {
		Fstart: int32(0x2090),
		Fend:   int32(0x209c),
	},
	258: {
		Fstart: int32(0x2102),
		Fend:   int32(0x2102),
	},
	259: {
		Fstart: int32(0x2107),
		Fend:   int32(0x2107),
	},
	260: {
		Fstart: int32(0x210a),
		Fend:   int32(0x2113),
	},
	261: {
		Fstart: int32(0x2115),
		Fend:   int32(0x2115),
	},
	262: {
		Fstart: int32(0x2119),
		Fend:   int32(0x211d),
	},
	263: {
		Fstart: int32(0x2124),
		Fend:   int32(0x2124),
	},
	264: {
		Fstart: int32(0x2126),
		Fend:   int32(0x2126),
	},
	265: {
		Fstart: int32(0x2128),
		Fend:   int32(0x2128),
	},
	266: {
		Fstart: int32(0x212a),
		Fend:   int32(0x212d),
	},
	267: {
		Fstart: int32(0x212f),
		Fend:   int32(0x2139),
	},
	268: {
		Fstart: int32(0x213c),
		Fend:   int32(0x213f),
	},
	269: {
		Fstart: int32(0x2145),
		Fend:   int32(0x2149),
	},
	270: {
		Fstart: int32(0x214e),
		Fend:   int32(0x214e),
	},
	271: {
		Fstart: int32(0x2183),
		Fend:   int32(0x2184),
	},
	272: {
		Fstart: int32(0x2c00),
		Fend:   int32(0x2ce4),
	},
	273: {
		Fstart: int32(0x2ceb),
		Fend:   int32(0x2cee),
	},
	274: {
		Fstart: int32(0x2cf2),
		Fend:   int32(0x2cf3),
	},
	275: {
		Fstart: int32(0x2d00),
		Fend:   int32(0x2d25),
	},
	276: {
		Fstart: int32(0x2d27),
		Fend:   int32(0x2d27),
	},
	277: {
		Fstart: int32(0x2d2d),
		Fend:   int32(0x2d2d),
	},
	278: {
		Fstart: int32(0x2d30),
		Fend:   int32(0x2d67),
	},
	279: {
		Fstart: int32(0x2d6f),
		Fend:   int32(0x2d6f),
	},
	280: {
		Fstart: int32(0x2d80),
		Fend:   int32(0x2d96),
	},
	281: {
		Fstart: int32(0x2da0),
		Fend:   int32(0x2da6),
	},
	282: {
		Fstart: int32(0x2da8),
		Fend:   int32(0x2dae),
	},
	283: {
		Fstart: int32(0x2db0),
		Fend:   int32(0x2db6),
	},
	284: {
		Fstart: int32(0x2db8),
		Fend:   int32(0x2dbe),
	},
	285: {
		Fstart: int32(0x2dc0),
		Fend:   int32(0x2dc6),
	},
	286: {
		Fstart: int32(0x2dc8),
		Fend:   int32(0x2dce),
	},
	287: {
		Fstart: int32(0x2dd0),
		Fend:   int32(0x2dd6),
	},
	288: {
		Fstart: int32(0x2dd8),
		Fend:   int32(0x2dde),
	},
	289: {
		Fstart: int32(0x2e2f),
		Fend:   int32(0x2e2f),
	},
	290: {
		Fstart: int32(0x3005),
		Fend:   int32(0x3006),
	},
	291: {
		Fstart: int32(0x3031),
		Fend:   int32(0x3035),
	},
	292: {
		Fstart: int32(0x303b),
		Fend:   int32(0x303c),
	},
	293: {
		Fstart: int32(0x3041),
		Fend:   int32(0x3096),
	},
	294: {
		Fstart: int32(0x309d),
		Fend:   int32(0x309f),
	},
	295: {
		Fstart: int32(0x30a1),
		Fend:   int32(0x30fa),
	},
	296: {
		Fstart: int32(0x30fc),
		Fend:   int32(0x30ff),
	},
	297: {
		Fstart: int32(0x3105),
		Fend:   int32(0x312f),
	},
	298: {
		Fstart: int32(0x3131),
		Fend:   int32(0x318e),
	},
	299: {
		Fstart: int32(0x31a0),
		Fend:   int32(0x31bf),
	},
	300: {
		Fstart: int32(0x31f0),
		Fend:   int32(0x31ff),
	},
	301: {
		Fstart: int32(0x3400),
		Fend:   int32(0x4dbf),
	},
	302: {
		Fstart: int32(0x4e00),
		Fend:   int32(0xa48c),
	},
	303: {
		Fstart: int32(0xa4d0),
		Fend:   int32(0xa4fd),
	},
	304: {
		Fstart: int32(0xa500),
		Fend:   int32(0xa60c),
	},
	305: {
		Fstart: int32(0xa610),
		Fend:   int32(0xa61f),
	},
	306: {
		Fstart: int32(0xa62a),
		Fend:   int32(0xa62b),
	},
	307: {
		Fstart: int32(0xa640),
		Fend:   int32(0xa66e),
	},
	308: {
		Fstart: int32(0xa67f),
		Fend:   int32(0xa69d),
	},
	309: {
		Fstart: int32(0xa6a0),
		Fend:   int32(0xa6e5),
	},
	310: {
		Fstart: int32(0xa717),
		Fend:   int32(0xa71f),
	},
	311: {
		Fstart: int32(0xa722),
		Fend:   int32(0xa788),
	},
	312: {
		Fstart: int32(0xa78b),
		Fend:   int32(0xa7cd),
	},
	313: {
		Fstart: int32(0xa7d0),
		Fend:   int32(0xa7d1),
	},
	314: {
		Fstart: int32(0xa7d3),
		Fend:   int32(0xa7d3),
	},
	315: {
		Fstart: int32(0xa7d5),
		Fend:   int32(0xa7dc),
	},
	316: {
		Fstart: int32(0xa7f2),
		Fend:   int32(0xa801),
	},
	317: {
		Fstart: int32(0xa803),
		Fend:   int32(0xa805),
	},
	318: {
		Fstart: int32(0xa807),
		Fend:   int32(0xa80a),
	},
	319: {
		Fstart: int32(0xa80c),
		Fend:   int32(0xa822),
	},
	320: {
		Fstart: int32(0xa840),
		Fend:   int32(0xa873),
	},
	321: {
		Fstart: int32(0xa882),
		Fend:   int32(0xa8b3),
	},
	322: {
		Fstart: int32(0xa8f2),
		Fend:   int32(0xa8f7),
	},
	323: {
		Fstart: int32(0xa8fb),
		Fend:   int32(0xa8fb),
	},
	324: {
		Fstart: int32(0xa8fd),
		Fend:   int32(0xa8fe),
	},
	325: {
		Fstart: int32(0xa90a),
		Fend:   int32(0xa925),
	},
	326: {
		Fstart: int32(0xa930),
		Fend:   int32(0xa946),
	},
	327: {
		Fstart: int32(0xa960),
		Fend:   int32(0xa97c),
	},
	328: {
		Fstart: int32(0xa984),
		Fend:   int32(0xa9b2),
	},
	329: {
		Fstart: int32(0xa9cf),
		Fend:   int32(0xa9cf),
	},
	330: {
		Fstart: int32(0xa9e0),
		Fend:   int32(0xa9e4),
	},
	331: {
		Fstart: int32(0xa9e6),
		Fend:   int32(0xa9ef),
	},
	332: {
		Fstart: int32(0xa9fa),
		Fend:   int32(0xa9fe),
	},
	333: {
		Fstart: int32(0xaa00),
		Fend:   int32(0xaa28),
	},
	334: {
		Fstart: int32(0xaa40),
		Fend:   int32(0xaa42),
	},
	335: {
		Fstart: int32(0xaa44),
		Fend:   int32(0xaa4b),
	},
	336: {
		Fstart: int32(0xaa60),
		Fend:   int32(0xaa76),
	},
	337: {
		Fstart: int32(0xaa7a),
		Fend:   int32(0xaa7a),
	},
	338: {
		Fstart: int32(0xaa7e),
		Fend:   int32(0xaaaf),
	},
	339: {
		Fstart: int32(0xaab1),
		Fend:   int32(0xaab1),
	},
	340: {
		Fstart: int32(0xaab5),
		Fend:   int32(0xaab6),
	},
	341: {
		Fstart: int32(0xaab9),
		Fend:   int32(0xaabd),
	},
	342: {
		Fstart: int32(0xaac0),
		Fend:   int32(0xaac0),
	},
	343: {
		Fstart: int32(0xaac2),
		Fend:   int32(0xaac2),
	},
	344: {
		Fstart: int32(0xaadb),
		Fend:   int32(0xaadd),
	},
	345: {
		Fstart: int32(0xaae0),
		Fend:   int32(0xaaea),
	},
	346: {
		Fstart: int32(0xaaf2),
		Fend:   int32(0xaaf4),
	},
	347: {
		Fstart: int32(0xab01),
		Fend:   int32(0xab06),
	},
	348: {
		Fstart: int32(0xab09),
		Fend:   int32(0xab0e),
	},
	349: {
		Fstart: int32(0xab11),
		Fend:   int32(0xab16),
	},
	350: {
		Fstart: int32(0xab20),
		Fend:   int32(0xab26),
	},
	351: {
		Fstart: int32(0xab28),
		Fend:   int32(0xab2e),
	},
	352: {
		Fstart: int32(0xab30),
		Fend:   int32(0xab5a),
	},
	353: {
		Fstart: int32(0xab5c),
		Fend:   int32(0xab69),
	},
	354: {
		Fstart: int32(0xab70),
		Fend:   int32(0xabe2),
	},
	355: {
		Fstart: int32(0xac00),
		Fend:   int32(0xd7a3),
	},
	356: {
		Fstart: int32(0xd7b0),
		Fend:   int32(0xd7c6),
	},
	357: {
		Fstart: int32(0xd7cb),
		Fend:   int32(0xd7fb),
	},
	358: {
		Fstart: int32(0xf900),
		Fend:   int32(0xfa6d),
	},
	359: {
		Fstart: int32(0xfa70),
		Fend:   int32(0xfad9),
	},
	360: {
		Fstart: int32(0xfb00),
		Fend:   int32(0xfb06),
	},
	361: {
		Fstart: int32(0xfb13),
		Fend:   int32(0xfb17),
	},
	362: {
		Fstart: int32(0xfb1d),
		Fend:   int32(0xfb1d),
	},
	363: {
		Fstart: int32(0xfb1f),
		Fend:   int32(0xfb28),
	},
	364: {
		Fstart: int32(0xfb2a),
		Fend:   int32(0xfb36),
	},
	365: {
		Fstart: int32(0xfb38),
		Fend:   int32(0xfb3c),
	},
	366: {
		Fstart: int32(0xfb3e),
		Fend:   int32(0xfb3e),
	},
	367: {
		Fstart: int32(0xfb40),
		Fend:   int32(0xfb41),
	},
	368: {
		Fstart: int32(0xfb43),
		Fend:   int32(0xfb44),
	},
	369: {
		Fstart: int32(0xfb46),
		Fend:   int32(0xfbb1),
	},
	370: {
		Fstart: int32(0xfbd3),
		Fend:   int32(0xfd3d),
	},
	371: {
		Fstart: int32(0xfd50),
		Fend:   int32(0xfd8f),
	},
	372: {
		Fstart: int32(0xfd92),
		Fend:   int32(0xfdc7),
	},
	373: {
		Fstart: int32(0xfdf0),
		Fend:   int32(0xfdfb),
	},
	374: {
		Fstart: int32(0xfe70),
		Fend:   int32(0xfe74),
	},
	375: {
		Fstart: int32(0xfe76),
		Fend:   int32(0xfefc),
	},
	376: {
		Fstart: int32(0xff21),
		Fend:   int32(0xff3a),
	},
	377: {
		Fstart: int32(0xff41),
		Fend:   int32(0xff5a),
	},
	378: {
		Fstart: int32(0xff66),
		Fend:   int32(0xffbe),
	},
	379: {
		Fstart: int32(0xffc2),
		Fend:   int32(0xffc7),
	},
	380: {
		Fstart: int32(0xffca),
		Fend:   int32(0xffcf),
	},
	381: {
		Fstart: int32(0xffd2),
		Fend:   int32(0xffd7),
	},
	382: {
		Fstart: int32(0xffda),
		Fend:   int32(0xffdc),
	},
	383: {
		Fstart: int32(0x10000),
		Fend:   int32(0x1000b),
	},
	384: {
		Fstart: int32(0x1000d),
		Fend:   int32(0x10026),
	},
	385: {
		Fstart: int32(0x10028),
		Fend:   int32(0x1003a),
	},
	386: {
		Fstart: int32(0x1003c),
		Fend:   int32(0x1003d),
	},
	387: {
		Fstart: int32(0x1003f),
		Fend:   int32(0x1004d),
	},
	388: {
		Fstart: int32(0x10050),
		Fend:   int32(0x1005d),
	},
	389: {
		Fstart: int32(0x10080),
		Fend:   int32(0x100fa),
	},
	390: {
		Fstart: int32(0x10280),
		Fend:   int32(0x1029c),
	},
	391: {
		Fstart: int32(0x102a0),
		Fend:   int32(0x102d0),
	},
	392: {
		Fstart: int32(0x10300),
		Fend:   int32(0x1031f),
	},
	393: {
		Fstart: int32(0x1032d),
		Fend:   int32(0x10340),
	},
	394: {
		Fstart: int32(0x10342),
		Fend:   int32(0x10349),
	},
	395: {
		Fstart: int32(0x10350),
		Fend:   int32(0x10375),
	},
	396: {
		Fstart: int32(0x10380),
		Fend:   int32(0x1039d),
	},
	397: {
		Fstart: int32(0x103a0),
		Fend:   int32(0x103c3),
	},
	398: {
		Fstart: int32(0x103c8),
		Fend:   int32(0x103cf),
	},
	399: {
		Fstart: int32(0x10400),
		Fend:   int32(0x1049d),
	},
	400: {
		Fstart: int32(0x104b0),
		Fend:   int32(0x104d3),
	},
	401: {
		Fstart: int32(0x104d8),
		Fend:   int32(0x104fb),
	},
	402: {
		Fstart: int32(0x10500),
		Fend:   int32(0x10527),
	},
	403: {
		Fstart: int32(0x10530),
		Fend:   int32(0x10563),
	},
	404: {
		Fstart: int32(0x10570),
		Fend:   int32(0x1057a),
	},
	405: {
		Fstart: int32(0x1057c),
		Fend:   int32(0x1058a),
	},
	406: {
		Fstart: int32(0x1058c),
		Fend:   int32(0x10592),
	},
	407: {
		Fstart: int32(0x10594),
		Fend:   int32(0x10595),
	},
	408: {
		Fstart: int32(0x10597),
		Fend:   int32(0x105a1),
	},
	409: {
		Fstart: int32(0x105a3),
		Fend:   int32(0x105b1),
	},
	410: {
		Fstart: int32(0x105b3),
		Fend:   int32(0x105b9),
	},
	411: {
		Fstart: int32(0x105bb),
		Fend:   int32(0x105bc),
	},
	412: {
		Fstart: int32(0x105c0),
		Fend:   int32(0x105f3),
	},
	413: {
		Fstart: int32(0x10600),
		Fend:   int32(0x10736),
	},
	414: {
		Fstart: int32(0x10740),
		Fend:   int32(0x10755),
	},
	415: {
		Fstart: int32(0x10760),
		Fend:   int32(0x10767),
	},
	416: {
		Fstart: int32(0x10780),
		Fend:   int32(0x10785),
	},
	417: {
		Fstart: int32(0x10787),
		Fend:   int32(0x107b0),
	},
	418: {
		Fstart: int32(0x107b2),
		Fend:   int32(0x107ba),
	},
	419: {
		Fstart: int32(0x10800),
		Fend:   int32(0x10805),
	},
	420: {
		Fstart: int32(0x10808),
		Fend:   int32(0x10808),
	},
	421: {
		Fstart: int32(0x1080a),
		Fend:   int32(0x10835),
	},
	422: {
		Fstart: int32(0x10837),
		Fend:   int32(0x10838),
	},
	423: {
		Fstart: int32(0x1083c),
		Fend:   int32(0x1083c),
	},
	424: {
		Fstart: int32(0x1083f),
		Fend:   int32(0x10855),
	},
	425: {
		Fstart: int32(0x10860),
		Fend:   int32(0x10876),
	},
	426: {
		Fstart: int32(0x10880),
		Fend:   int32(0x1089e),
	},
	427: {
		Fstart: int32(0x108e0),
		Fend:   int32(0x108f2),
	},
	428: {
		Fstart: int32(0x108f4),
		Fend:   int32(0x108f5),
	},
	429: {
		Fstart: int32(0x10900),
		Fend:   int32(0x10915),
	},
	430: {
		Fstart: int32(0x10920),
		Fend:   int32(0x10939),
	},
	431: {
		Fstart: int32(0x10980),
		Fend:   int32(0x109b7),
	},
	432: {
		Fstart: int32(0x109be),
		Fend:   int32(0x109bf),
	},
	433: {
		Fstart: int32(0x10a00),
		Fend:   int32(0x10a00),
	},
	434: {
		Fstart: int32(0x10a10),
		Fend:   int32(0x10a13),
	},
	435: {
		Fstart: int32(0x10a15),
		Fend:   int32(0x10a17),
	},
	436: {
		Fstart: int32(0x10a19),
		Fend:   int32(0x10a35),
	},
	437: {
		Fstart: int32(0x10a60),
		Fend:   int32(0x10a7c),
	},
	438: {
		Fstart: int32(0x10a80),
		Fend:   int32(0x10a9c),
	},
	439: {
		Fstart: int32(0x10ac0),
		Fend:   int32(0x10ac7),
	},
	440: {
		Fstart: int32(0x10ac9),
		Fend:   int32(0x10ae4),
	},
	441: {
		Fstart: int32(0x10b00),
		Fend:   int32(0x10b35),
	},
	442: {
		Fstart: int32(0x10b40),
		Fend:   int32(0x10b55),
	},
	443: {
		Fstart: int32(0x10b60),
		Fend:   int32(0x10b72),
	},
	444: {
		Fstart: int32(0x10b80),
		Fend:   int32(0x10b91),
	},
	445: {
		Fstart: int32(0x10c00),
		Fend:   int32(0x10c48),
	},
	446: {
		Fstart: int32(0x10c80),
		Fend:   int32(0x10cb2),
	},
	447: {
		Fstart: int32(0x10cc0),
		Fend:   int32(0x10cf2),
	},
	448: {
		Fstart: int32(0x10d00),
		Fend:   int32(0x10d23),
	},
	449: {
		Fstart: int32(0x10d4a),
		Fend:   int32(0x10d65),
	},
	450: {
		Fstart: int32(0x10d6f),
		Fend:   int32(0x10d85),
	},
	451: {
		Fstart: int32(0x10e80),
		Fend:   int32(0x10ea9),
	},
	452: {
		Fstart: int32(0x10eb0),
		Fend:   int32(0x10eb1),
	},
	453: {
		Fstart: int32(0x10ec2),
		Fend:   int32(0x10ec4),
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
		Fstart: int32(0x1123f),
		Fend:   int32(0x11240),
	},
	477: {
		Fstart: int32(0x11280),
		Fend:   int32(0x11286),
	},
	478: {
		Fstart: int32(0x11288),
		Fend:   int32(0x11288),
	},
	479: {
		Fstart: int32(0x1128a),
		Fend:   int32(0x1128d),
	},
	480: {
		Fstart: int32(0x1128f),
		Fend:   int32(0x1129d),
	},
	481: {
		Fstart: int32(0x1129f),
		Fend:   int32(0x112a8),
	},
	482: {
		Fstart: int32(0x112b0),
		Fend:   int32(0x112de),
	},
	483: {
		Fstart: int32(0x11305),
		Fend:   int32(0x1130c),
	},
	484: {
		Fstart: int32(0x1130f),
		Fend:   int32(0x11310),
	},
	485: {
		Fstart: int32(0x11313),
		Fend:   int32(0x11328),
	},
	486: {
		Fstart: int32(0x1132a),
		Fend:   int32(0x11330),
	},
	487: {
		Fstart: int32(0x11332),
		Fend:   int32(0x11333),
	},
	488: {
		Fstart: int32(0x11335),
		Fend:   int32(0x11339),
	},
	489: {
		Fstart: int32(0x1133d),
		Fend:   int32(0x1133d),
	},
	490: {
		Fstart: int32(0x11350),
		Fend:   int32(0x11350),
	},
	491: {
		Fstart: int32(0x1135d),
		Fend:   int32(0x11361),
	},
	492: {
		Fstart: int32(0x11380),
		Fend:   int32(0x11389),
	},
	493: {
		Fstart: int32(0x1138b),
		Fend:   int32(0x1138b),
	},
	494: {
		Fstart: int32(0x1138e),
		Fend:   int32(0x1138e),
	},
	495: {
		Fstart: int32(0x11390),
		Fend:   int32(0x113b5),
	},
	496: {
		Fstart: int32(0x113b7),
		Fend:   int32(0x113b7),
	},
	497: {
		Fstart: int32(0x113d1),
		Fend:   int32(0x113d1),
	},
	498: {
		Fstart: int32(0x113d3),
		Fend:   int32(0x113d3),
	},
	499: {
		Fstart: int32(0x11400),
		Fend:   int32(0x11434),
	},
	500: {
		Fstart: int32(0x11447),
		Fend:   int32(0x1144a),
	},
	501: {
		Fstart: int32(0x1145f),
		Fend:   int32(0x11461),
	},
	502: {
		Fstart: int32(0x11480),
		Fend:   int32(0x114af),
	},
	503: {
		Fstart: int32(0x114c4),
		Fend:   int32(0x114c5),
	},
	504: {
		Fstart: int32(0x114c7),
		Fend:   int32(0x114c7),
	},
	505: {
		Fstart: int32(0x11580),
		Fend:   int32(0x115ae),
	},
	506: {
		Fstart: int32(0x115d8),
		Fend:   int32(0x115db),
	},
	507: {
		Fstart: int32(0x11600),
		Fend:   int32(0x1162f),
	},
	508: {
		Fstart: int32(0x11644),
		Fend:   int32(0x11644),
	},
	509: {
		Fstart: int32(0x11680),
		Fend:   int32(0x116aa),
	},
	510: {
		Fstart: int32(0x116b8),
		Fend:   int32(0x116b8),
	},
	511: {
		Fstart: int32(0x11700),
		Fend:   int32(0x1171a),
	},
	512: {
		Fstart: int32(0x11740),
		Fend:   int32(0x11746),
	},
	513: {
		Fstart: int32(0x11800),
		Fend:   int32(0x1182b),
	},
	514: {
		Fstart: int32(0x118a0),
		Fend:   int32(0x118df),
	},
	515: {
		Fstart: int32(0x118ff),
		Fend:   int32(0x11906),
	},
	516: {
		Fstart: int32(0x11909),
		Fend:   int32(0x11909),
	},
	517: {
		Fstart: int32(0x1190c),
		Fend:   int32(0x11913),
	},
	518: {
		Fstart: int32(0x11915),
		Fend:   int32(0x11916),
	},
	519: {
		Fstart: int32(0x11918),
		Fend:   int32(0x1192f),
	},
	520: {
		Fstart: int32(0x1193f),
		Fend:   int32(0x1193f),
	},
	521: {
		Fstart: int32(0x11941),
		Fend:   int32(0x11941),
	},
	522: {
		Fstart: int32(0x119a0),
		Fend:   int32(0x119a7),
	},
	523: {
		Fstart: int32(0x119aa),
		Fend:   int32(0x119d0),
	},
	524: {
		Fstart: int32(0x119e1),
		Fend:   int32(0x119e1),
	},
	525: {
		Fstart: int32(0x119e3),
		Fend:   int32(0x119e3),
	},
	526: {
		Fstart: int32(0x11a00),
		Fend:   int32(0x11a00),
	},
	527: {
		Fstart: int32(0x11a0b),
		Fend:   int32(0x11a32),
	},
	528: {
		Fstart: int32(0x11a3a),
		Fend:   int32(0x11a3a),
	},
	529: {
		Fstart: int32(0x11a50),
		Fend:   int32(0x11a50),
	},
	530: {
		Fstart: int32(0x11a5c),
		Fend:   int32(0x11a89),
	},
	531: {
		Fstart: int32(0x11a9d),
		Fend:   int32(0x11a9d),
	},
	532: {
		Fstart: int32(0x11ab0),
		Fend:   int32(0x11af8),
	},
	533: {
		Fstart: int32(0x11bc0),
		Fend:   int32(0x11be0),
	},
	534: {
		Fstart: int32(0x11c00),
		Fend:   int32(0x11c08),
	},
	535: {
		Fstart: int32(0x11c0a),
		Fend:   int32(0x11c2e),
	},
	536: {
		Fstart: int32(0x11c40),
		Fend:   int32(0x11c40),
	},
	537: {
		Fstart: int32(0x11c72),
		Fend:   int32(0x11c8f),
	},
	538: {
		Fstart: int32(0x11d00),
		Fend:   int32(0x11d06),
	},
	539: {
		Fstart: int32(0x11d08),
		Fend:   int32(0x11d09),
	},
	540: {
		Fstart: int32(0x11d0b),
		Fend:   int32(0x11d30),
	},
	541: {
		Fstart: int32(0x11d46),
		Fend:   int32(0x11d46),
	},
	542: {
		Fstart: int32(0x11d60),
		Fend:   int32(0x11d65),
	},
	543: {
		Fstart: int32(0x11d67),
		Fend:   int32(0x11d68),
	},
	544: {
		Fstart: int32(0x11d6a),
		Fend:   int32(0x11d89),
	},
	545: {
		Fstart: int32(0x11d98),
		Fend:   int32(0x11d98),
	},
	546: {
		Fstart: int32(0x11ee0),
		Fend:   int32(0x11ef2),
	},
	547: {
		Fstart: int32(0x11f02),
		Fend:   int32(0x11f02),
	},
	548: {
		Fstart: int32(0x11f04),
		Fend:   int32(0x11f10),
	},
	549: {
		Fstart: int32(0x11f12),
		Fend:   int32(0x11f33),
	},
	550: {
		Fstart: int32(0x11fb0),
		Fend:   int32(0x11fb0),
	},
	551: {
		Fstart: int32(0x12000),
		Fend:   int32(0x12399),
	},
	552: {
		Fstart: int32(0x12480),
		Fend:   int32(0x12543),
	},
	553: {
		Fstart: int32(0x12f90),
		Fend:   int32(0x12ff0),
	},
	554: {
		Fstart: int32(0x13000),
		Fend:   int32(0x1342f),
	},
	555: {
		Fstart: int32(0x13441),
		Fend:   int32(0x13446),
	},
	556: {
		Fstart: int32(0x13460),
		Fend:   int32(0x143fa),
	},
	557: {
		Fstart: int32(0x14400),
		Fend:   int32(0x14646),
	},
	558: {
		Fstart: int32(0x16100),
		Fend:   int32(0x1611d),
	},
	559: {
		Fstart: int32(0x16800),
		Fend:   int32(0x16a38),
	},
	560: {
		Fstart: int32(0x16a40),
		Fend:   int32(0x16a5e),
	},
	561: {
		Fstart: int32(0x16a70),
		Fend:   int32(0x16abe),
	},
	562: {
		Fstart: int32(0x16ad0),
		Fend:   int32(0x16aed),
	},
	563: {
		Fstart: int32(0x16b00),
		Fend:   int32(0x16b2f),
	},
	564: {
		Fstart: int32(0x16b40),
		Fend:   int32(0x16b43),
	},
	565: {
		Fstart: int32(0x16b63),
		Fend:   int32(0x16b77),
	},
	566: {
		Fstart: int32(0x16b7d),
		Fend:   int32(0x16b8f),
	},
	567: {
		Fstart: int32(0x16d40),
		Fend:   int32(0x16d6c),
	},
	568: {
		Fstart: int32(0x16e40),
		Fend:   int32(0x16e7f),
	},
	569: {
		Fstart: int32(0x16f00),
		Fend:   int32(0x16f4a),
	},
	570: {
		Fstart: int32(0x16f50),
		Fend:   int32(0x16f50),
	},
	571: {
		Fstart: int32(0x16f93),
		Fend:   int32(0x16f9f),
	},
	572: {
		Fstart: int32(0x16fe0),
		Fend:   int32(0x16fe1),
	},
	573: {
		Fstart: int32(0x16fe3),
		Fend:   int32(0x16fe3),
	},
	574: {
		Fstart: int32(0x17000),
		Fend:   int32(0x187f7),
	},
	575: {
		Fstart: int32(0x18800),
		Fend:   int32(0x18cd5),
	},
	576: {
		Fstart: int32(0x18cff),
		Fend:   int32(0x18d08),
	},
	577: {
		Fstart: int32(0x1aff0),
		Fend:   int32(0x1aff3),
	},
	578: {
		Fstart: int32(0x1aff5),
		Fend:   int32(0x1affb),
	},
	579: {
		Fstart: int32(0x1affd),
		Fend:   int32(0x1affe),
	},
	580: {
		Fstart: int32(0x1b000),
		Fend:   int32(0x1b122),
	},
	581: {
		Fstart: int32(0x1b132),
		Fend:   int32(0x1b132),
	},
	582: {
		Fstart: int32(0x1b150),
		Fend:   int32(0x1b152),
	},
	583: {
		Fstart: int32(0x1b155),
		Fend:   int32(0x1b155),
	},
	584: {
		Fstart: int32(0x1b164),
		Fend:   int32(0x1b167),
	},
	585: {
		Fstart: int32(0x1b170),
		Fend:   int32(0x1b2fb),
	},
	586: {
		Fstart: int32(0x1bc00),
		Fend:   int32(0x1bc6a),
	},
	587: {
		Fstart: int32(0x1bc70),
		Fend:   int32(0x1bc7c),
	},
	588: {
		Fstart: int32(0x1bc80),
		Fend:   int32(0x1bc88),
	},
	589: {
		Fstart: int32(0x1bc90),
		Fend:   int32(0x1bc99),
	},
	590: {
		Fstart: int32(0x1d400),
		Fend:   int32(0x1d454),
	},
	591: {
		Fstart: int32(0x1d456),
		Fend:   int32(0x1d49c),
	},
	592: {
		Fstart: int32(0x1d49e),
		Fend:   int32(0x1d49f),
	},
	593: {
		Fstart: int32(0x1d4a2),
		Fend:   int32(0x1d4a2),
	},
	594: {
		Fstart: int32(0x1d4a5),
		Fend:   int32(0x1d4a6),
	},
	595: {
		Fstart: int32(0x1d4a9),
		Fend:   int32(0x1d4ac),
	},
	596: {
		Fstart: int32(0x1d4ae),
		Fend:   int32(0x1d4b9),
	},
	597: {
		Fstart: int32(0x1d4bb),
		Fend:   int32(0x1d4bb),
	},
	598: {
		Fstart: int32(0x1d4bd),
		Fend:   int32(0x1d4c3),
	},
	599: {
		Fstart: int32(0x1d4c5),
		Fend:   int32(0x1d505),
	},
	600: {
		Fstart: int32(0x1d507),
		Fend:   int32(0x1d50a),
	},
	601: {
		Fstart: int32(0x1d50d),
		Fend:   int32(0x1d514),
	},
	602: {
		Fstart: int32(0x1d516),
		Fend:   int32(0x1d51c),
	},
	603: {
		Fstart: int32(0x1d51e),
		Fend:   int32(0x1d539),
	},
	604: {
		Fstart: int32(0x1d53b),
		Fend:   int32(0x1d53e),
	},
	605: {
		Fstart: int32(0x1d540),
		Fend:   int32(0x1d544),
	},
	606: {
		Fstart: int32(0x1d546),
		Fend:   int32(0x1d546),
	},
	607: {
		Fstart: int32(0x1d54a),
		Fend:   int32(0x1d550),
	},
	608: {
		Fstart: int32(0x1d552),
		Fend:   int32(0x1d6a5),
	},
	609: {
		Fstart: int32(0x1d6a8),
		Fend:   int32(0x1d6c0),
	},
	610: {
		Fstart: int32(0x1d6c2),
		Fend:   int32(0x1d6da),
	},
	611: {
		Fstart: int32(0x1d6dc),
		Fend:   int32(0x1d6fa),
	},
	612: {
		Fstart: int32(0x1d6fc),
		Fend:   int32(0x1d714),
	},
	613: {
		Fstart: int32(0x1d716),
		Fend:   int32(0x1d734),
	},
	614: {
		Fstart: int32(0x1d736),
		Fend:   int32(0x1d74e),
	},
	615: {
		Fstart: int32(0x1d750),
		Fend:   int32(0x1d76e),
	},
	616: {
		Fstart: int32(0x1d770),
		Fend:   int32(0x1d788),
	},
	617: {
		Fstart: int32(0x1d78a),
		Fend:   int32(0x1d7a8),
	},
	618: {
		Fstart: int32(0x1d7aa),
		Fend:   int32(0x1d7c2),
	},
	619: {
		Fstart: int32(0x1d7c4),
		Fend:   int32(0x1d7cb),
	},
	620: {
		Fstart: int32(0x1df00),
		Fend:   int32(0x1df1e),
	},
	621: {
		Fstart: int32(0x1df25),
		Fend:   int32(0x1df2a),
	},
	622: {
		Fstart: int32(0x1e030),
		Fend:   int32(0x1e06d),
	},
	623: {
		Fstart: int32(0x1e100),
		Fend:   int32(0x1e12c),
	},
	624: {
		Fstart: int32(0x1e137),
		Fend:   int32(0x1e13d),
	},
	625: {
		Fstart: int32(0x1e14e),
		Fend:   int32(0x1e14e),
	},
	626: {
		Fstart: int32(0x1e290),
		Fend:   int32(0x1e2ad),
	},
	627: {
		Fstart: int32(0x1e2c0),
		Fend:   int32(0x1e2eb),
	},
	628: {
		Fstart: int32(0x1e4d0),
		Fend:   int32(0x1e4eb),
	},
	629: {
		Fstart: int32(0x1e5d0),
		Fend:   int32(0x1e5ed),
	},
	630: {
		Fstart: int32(0x1e5f0),
		Fend:   int32(0x1e5f0),
	},
	631: {
		Fstart: int32(0x1e7e0),
		Fend:   int32(0x1e7e6),
	},
	632: {
		Fstart: int32(0x1e7e8),
		Fend:   int32(0x1e7eb),
	},
	633: {
		Fstart: int32(0x1e7ed),
		Fend:   int32(0x1e7ee),
	},
	634: {
		Fstart: int32(0x1e7f0),
		Fend:   int32(0x1e7fe),
	},
	635: {
		Fstart: int32(0x1e800),
		Fend:   int32(0x1e8c4),
	},
	636: {
		Fstart: int32(0x1e900),
		Fend:   int32(0x1e943),
	},
	637: {
		Fstart: int32(0x1e94b),
		Fend:   int32(0x1e94b),
	},
	638: {
		Fstart: int32(0x1ee00),
		Fend:   int32(0x1ee03),
	},
	639: {
		Fstart: int32(0x1ee05),
		Fend:   int32(0x1ee1f),
	},
	640: {
		Fstart: int32(0x1ee21),
		Fend:   int32(0x1ee22),
	},
	641: {
		Fstart: int32(0x1ee24),
		Fend:   int32(0x1ee24),
	},
	642: {
		Fstart: int32(0x1ee27),
		Fend:   int32(0x1ee27),
	},
	643: {
		Fstart: int32(0x1ee29),
		Fend:   int32(0x1ee32),
	},
	644: {
		Fstart: int32(0x1ee34),
		Fend:   int32(0x1ee37),
	},
	645: {
		Fstart: int32(0x1ee39),
		Fend:   int32(0x1ee39),
	},
	646: {
		Fstart: int32(0x1ee3b),
		Fend:   int32(0x1ee3b),
	},
	647: {
		Fstart: int32(0x1ee42),
		Fend:   int32(0x1ee42),
	},
	648: {
		Fstart: int32(0x1ee47),
		Fend:   int32(0x1ee47),
	},
	649: {
		Fstart: int32(0x1ee49),
		Fend:   int32(0x1ee49),
	},
	650: {
		Fstart: int32(0x1ee4b),
		Fend:   int32(0x1ee4b),
	},
	651: {
		Fstart: int32(0x1ee4d),
		Fend:   int32(0x1ee4f),
	},
	652: {
		Fstart: int32(0x1ee51),
		Fend:   int32(0x1ee52),
	},
	653: {
		Fstart: int32(0x1ee54),
		Fend:   int32(0x1ee54),
	},
	654: {
		Fstart: int32(0x1ee57),
		Fend:   int32(0x1ee57),
	},
	655: {
		Fstart: int32(0x1ee59),
		Fend:   int32(0x1ee59),
	},
	656: {
		Fstart: int32(0x1ee5b),
		Fend:   int32(0x1ee5b),
	},
	657: {
		Fstart: int32(0x1ee5d),
		Fend:   int32(0x1ee5d),
	},
	658: {
		Fstart: int32(0x1ee5f),
		Fend:   int32(0x1ee5f),
	},
	659: {
		Fstart: int32(0x1ee61),
		Fend:   int32(0x1ee62),
	},
	660: {
		Fstart: int32(0x1ee64),
		Fend:   int32(0x1ee64),
	},
	661: {
		Fstart: int32(0x1ee67),
		Fend:   int32(0x1ee6a),
	},
	662: {
		Fstart: int32(0x1ee6c),
		Fend:   int32(0x1ee72),
	},
	663: {
		Fstart: int32(0x1ee74),
		Fend:   int32(0x1ee77),
	},
	664: {
		Fstart: int32(0x1ee79),
		Fend:   int32(0x1ee7c),
	},
	665: {
		Fstart: int32(0x1ee7e),
		Fend:   int32(0x1ee7e),
	},
	666: {
		Fstart: int32(0x1ee80),
		Fend:   int32(0x1ee89),
	},
	667: {
		Fstart: int32(0x1ee8b),
		Fend:   int32(0x1ee9b),
	},
	668: {
		Fstart: int32(0x1eea1),
		Fend:   int32(0x1eea3),
	},
	669: {
		Fstart: int32(0x1eea5),
		Fend:   int32(0x1eea9),
	},
	670: {
		Fstart: int32(0x1eeab),
		Fend:   int32(0x1eebb),
	},
	671: {
		Fstart: int32(0x20000),
		Fend:   int32(0x2a6df),
	},
	672: {
		Fstart: int32(0x2a700),
		Fend:   int32(0x2b739),
	},
	673: {
		Fstart: int32(0x2b740),
		Fend:   int32(0x2b81d),
	},
	674: {
		Fstart: int32(0x2b820),
		Fend:   int32(0x2cea1),
	},
	675: {
		Fstart: int32(0x2ceb0),
		Fend:   int32(0x2ebe0),
	},
	676: {
		Fstart: int32(0x2ebf0),
		Fend:   int32(0x2ee5d),
	},
	677: {
		Fstart: int32(0x2f800),
		Fend:   int32(0x2fa1d),
	},
	678: {
		Fstart: int32(0x30000),
		Fend:   int32(0x3134a),
	},
	679: {
		Fstart: int32(0x31350),
		Fend:   int32(0x323af),
	},
}

func ts_lex(tls *libc.TLS, lexer uintptr, state TSStateId) (r uint8) {
	var eof, result, skip, v4 uint8
	var half_size, i, i1, i10, i11, i12, i13, i14, i15, i2, i3, i4, i5, i6, i7, i8, i9, index, mid_index, size uint32_t
	var lookahead1, v3 int32_t
	var range_token, range_token1, v2 uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = eof, half_size, i, i1, i10, i11, i12, i13, i14, i15, i2, i3, i4, i5, i6, i7, i8, i9, index, lookahead1, mid_index, range_token, range_token1, result, size, skip, v2, v3, v4
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
	switch int32(state) {
	case 0:
		if eof != 0 {
			state = uint16(918)
			goto next_state
		}
		i = uint32(0)
		for {
			if !(uint64(i) < libc.Uint64FromInt64(148)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token[i]) == lookahead1 {
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
			state = uint16(914)
			goto next_state
		}
		if lookahead1 == int32('5') || lookahead1 == int32('6') {
			state = uint16(1033)
			goto next_state
		}
		if int32('7') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(1085)
			goto next_state
		}
		return result
	case int32(1):
		if lookahead1 == int32('\t') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(17)
			goto next_state
		}
		if lookahead1 == int32(' ') {
			state = uint16(1070)
			goto next_state
		}
		if lookahead1 == int32('#') {
			state = uint16(1102)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&aux_sym_string_token2_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(680) - index
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
			state = uint16(1071)
			goto next_state
		}
		return result
	case int32(2):
		if lookahead1 == int32('\t') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(17)
			goto next_state
		}
		if lookahead1 == int32(' ') {
			state = uint16(1079)
			goto next_state
		}
		if lookahead1 == int32('#') {
			state = uint16(1080)
			goto next_state
		}
		if lookahead1 > int32(' ') && lookahead1 != int32('"') && lookahead1 != int32('#') && lookahead1 != int32('*') && lookahead1 != int32('<') && lookahead1 != int32('>') && lookahead1 != int32('?') && lookahead1 != int32('|') && (lookahead1 < int32(0x7f) || int32(0x9f) < lookahead1) {
			state = uint16(1081)
			goto next_state
		}
		return result
	case int32(3):
		if lookahead1 == int32('\n') {
			state = uint16(1104)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead1 == int32('!') {
			state = uint16(996)
			goto next_state
		}
		if lookahead1 == int32('"') {
			state = uint16(1001)
			goto next_state
		}
		if lookahead1 == int32('#') {
			state = uint16(1102)
			goto next_state
		}
		if lookahead1 == int32('0') {
			state = uint16(1082)
			goto next_state
		}
		if lookahead1 == int32('<') {
			state = uint16(998)
			goto next_state
		}
		if lookahead1 == int32('=') {
			state = uint16(996)
			goto next_state
		}
		if lookahead1 == int32('>') {
			state = uint16(998)
			goto next_state
		}
		if lookahead1 == int32('\t') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(6)
			goto next_state
		}
		if int32('1') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(1085)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&aux_sym_string_token2_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(680) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _9
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _9
	_9:
		if v4 != 0 {
			state = uint16(1078)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead1 == int32('!') {
			state = uint16(996)
			goto next_state
		}
		if lookahead1 == int32('"') {
			state = uint16(1001)
			goto next_state
		}
		if lookahead1 == int32('#') {
			state = uint16(1102)
			goto next_state
		}
		if lookahead1 == int32('<') {
			state = uint16(998)
			goto next_state
		}
		if lookahead1 == int32('=') {
			state = uint16(996)
			goto next_state
		}
		if lookahead1 == int32('>') {
			state = uint16(998)
			goto next_state
		}
		if lookahead1 == int32('\t') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(7)
			goto next_state
		}
		if lookahead1 == int32('A') || lookahead1 == int32('B') || lookahead1 == int32('D') || lookahead1 == int32('G') || lookahead1 == int32('R') || lookahead1 == int32('W') {
			state = uint16(1044)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('6') {
			state = uint16(1044)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead1 == int32('!') {
			state = uint16(996)
			goto next_state
		}
		if lookahead1 == int32('#') {
			state = uint16(1102)
			goto next_state
		}
		if lookahead1 == int32('0') {
			state = uint16(1082)
			goto next_state
		}
		if lookahead1 == int32('<') {
			state = uint16(998)
			goto next_state
		}
		if lookahead1 == int32('=') {
			state = uint16(996)
			goto next_state
		}
		if lookahead1 == int32('>') {
			state = uint16(998)
			goto next_state
		}
		if lookahead1 == int32('\t') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(6)
			goto next_state
		}
		if int32('1') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(1085)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&aux_sym_string_token2_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(680) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _13
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _13
	_13:
		if v4 != 0 {
			state = uint16(1078)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead1 == int32('!') {
			state = uint16(996)
			goto next_state
		}
		if lookahead1 == int32('#') {
			state = uint16(1102)
			goto next_state
		}
		if lookahead1 == int32('<') {
			state = uint16(998)
			goto next_state
		}
		if lookahead1 == int32('=') {
			state = uint16(996)
			goto next_state
		}
		if lookahead1 == int32('>') {
			state = uint16(998)
			goto next_state
		}
		if lookahead1 == int32('\t') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(7)
			goto next_state
		}
		if lookahead1 == int32('A') || lookahead1 == int32('B') || lookahead1 == int32('D') || lookahead1 == int32('G') || lookahead1 == int32('R') || lookahead1 == int32('W') {
			state = uint16(1044)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('6') {
			state = uint16(1044)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead1 == int32('"') {
			state = uint16(980)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead1 == int32('"') {
			state = uint16(1001)
			goto next_state
		}
		if lookahead1 == int32('#') {
			state = uint16(1102)
			goto next_state
		}
		if lookahead1 == int32('0') {
			state = uint16(1093)
			goto next_state
		}
		if lookahead1 == int32('1') {
			state = uint16(1097)
			goto next_state
		}
		if lookahead1 == int32('2') {
			state = uint16(1094)
			goto next_state
		}
		if lookahead1 == int32('F') {
			state = uint16(1072)
			goto next_state
		}
		if lookahead1 == int32('T') {
			state = uint16(1075)
			goto next_state
		}
		if lookahead1 == int32('\t') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(13)
			goto next_state
		}
		if int32('3') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(1096)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&aux_sym_string_token2_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(680) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _17
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _17
	_17:
		if v4 != 0 {
			state = uint16(1078)
			goto next_state
		}
		return result
	case int32(10):
		i1 = uint32(0)
		for {
			if !(uint64(i1) < libc.Uint64FromInt64(60)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token1[i1]) == lookahead1 {
				state = map_token1[i1+uint32(1)]
				goto next_state
			}
			goto _18
		_18:
			;
			i1 = i1 + uint32(2)
		}
		if lookahead1 == int32('\t') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(15)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(1086)
			goto next_state
		}
		return result
	case int32(11):
		i2 = uint32(0)
		for {
			if !(uint64(i2) < libc.Uint64FromInt64(36)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token2[i2]) == lookahead1 {
				state = map_token2[i2+uint32(1)]
				goto next_state
			}
			goto _19
		_19:
			;
			i2 = i2 + uint32(2)
		}
		if lookahead1 == int32('\t') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(12)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('2') {
			state = uint16(1098)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead1 == int32('#') {
			state = uint16(1102)
			goto next_state
		}
		if lookahead1 == int32('-') {
			state = uint16(19)
			goto next_state
		}
		if lookahead1 == int32('\t') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(12)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('2') {
			state = uint16(1098)
			goto next_state
		}
		return result
	case int32(13):
		if lookahead1 == int32('#') {
			state = uint16(1102)
			goto next_state
		}
		if lookahead1 == int32('0') {
			state = uint16(1093)
			goto next_state
		}
		if lookahead1 == int32('1') {
			state = uint16(1097)
			goto next_state
		}
		if lookahead1 == int32('2') {
			state = uint16(1094)
			goto next_state
		}
		if lookahead1 == int32('F') {
			state = uint16(1072)
			goto next_state
		}
		if lookahead1 == int32('T') {
			state = uint16(1075)
			goto next_state
		}
		if lookahead1 == int32('\t') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(13)
			goto next_state
		}
		if int32('3') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(1096)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&aux_sym_string_token2_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(680) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _23
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _23
	_23:
		if v4 != 0 {
			state = uint16(1078)
			goto next_state
		}
		return result
	case int32(14):
		if lookahead1 == int32('#') {
			state = uint16(1102)
			goto next_state
		}
		if lookahead1 == int32('0') {
			state = uint16(1088)
			goto next_state
		}
		if lookahead1 == int32('3') {
			state = uint16(1090)
			goto next_state
		}
		if lookahead1 == int32('\t') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(14)
			goto next_state
		}
		if lookahead1 == int32('1') || lookahead1 == int32('2') {
			state = uint16(1092)
			goto next_state
		}
		if int32('4') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(1091)
			goto next_state
		}
		return result
	case int32(15):
		i3 = uint32(0)
		for {
			if !(uint64(i3) < libc.Uint64FromInt64(36)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token3[i3]) == lookahead1 {
				state = map_token3[i3+uint32(1)]
				goto next_state
			}
			goto _24
		_24:
			;
			i3 = i3 + uint32(2)
		}
		if lookahead1 == int32('\t') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(15)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(1086)
			goto next_state
		}
		return result
	case int32(16):
		if lookahead1 == int32('#') {
			state = uint16(1102)
			goto next_state
		}
		if lookahead1 == int32('4') {
			state = uint16(1100)
			goto next_state
		}
		if lookahead1 == int32('\t') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(16)
			goto next_state
		}
		if int32('1') <= lookahead1 && lookahead1 <= int32('3') {
			state = uint16(1101)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(1099)
			goto next_state
		}
		return result
	case int32(17):
		if lookahead1 == int32('#') {
			state = uint16(1102)
			goto next_state
		}
		if lookahead1 == int32('\t') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(17)
			goto next_state
		}
		return result
	case int32(18):
		if lookahead1 == int32('#') {
			state = uint16(1102)
			goto next_state
		}
		if lookahead1 == int32('\t') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(17)
			goto next_state
		}
		if lookahead1 == int32('A') || lookahead1 == int32('B') || lookahead1 == int32('D') || lookahead1 == int32('G') || lookahead1 == int32('R') || lookahead1 == int32('W') {
			state = uint16(1032)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('6') {
			state = uint16(1032)
			goto next_state
		}
		return result
	case int32(19):
		if lookahead1 == int32('1') {
			state = uint16(987)
			goto next_state
		}
		return result
	case int32(20):
		if lookahead1 == int32('A') {
			state = uint16(498)
			goto next_state
		}
		if lookahead1 == int32('E') {
			state = uint16(354)
			goto next_state
		}
		return result
	case int32(21):
		if lookahead1 == int32('A') {
			state = uint16(750)
			goto next_state
		}
		if lookahead1 == int32('D') {
			state = uint16(273)
			goto next_state
		}
		if lookahead1 == int32('E') {
			state = uint16(600)
			goto next_state
		}
		if lookahead1 == int32('T') {
			state = uint16(910)
			goto next_state
		}
		if lookahead1 == int32('W') {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(22):
		if lookahead1 == int32('A') {
			state = uint16(505)
			goto next_state
		}
		return result
	case int32(23):
		if lookahead1 == int32('A') {
			state = uint16(506)
			goto next_state
		}
		return result
	case int32(24):
		if lookahead1 == int32('A') {
			state = uint16(507)
			goto next_state
		}
		return result
	case int32(25):
		if lookahead1 == int32('B') {
			state = uint16(127)
			goto next_state
		}
		if lookahead1 == int32('F') {
			state = uint16(660)
			goto next_state
		}
		if lookahead1 == int32('T') {
			state = uint16(241)
			goto next_state
		}
		return result
	case int32(26):
		if lookahead1 == int32('B') {
			state = uint16(477)
			goto next_state
		}
		return result
	case int32(27):
		if lookahead1 == int32('C') {
			state = uint16(726)
			goto next_state
		}
		if lookahead1 == int32('E') {
			state = uint16(145)
			goto next_state
		}
		if lookahead1 == int32('I') {
			state = uint16(542)
			goto next_state
		}
		if lookahead1 == int32('S') {
			state = uint16(325)
			goto next_state
		}
		return result
	case int32(28):
		if lookahead1 == int32('C') {
			state = uint16(667)
			goto next_state
		}
		return result
	case int32(29):
		if lookahead1 == int32('C') {
			state = uint16(668)
			goto next_state
		}
		return result
	case int32(30):
		if lookahead1 == int32('C') {
			state = uint16(669)
			goto next_state
		}
		return result
	case int32(31):
		if lookahead1 == int32('D') {
			state = uint16(644)
			goto next_state
		}
		return result
	case int32(32):
		if lookahead1 == int32('D') {
			state = uint16(736)
			goto next_state
		}
		return result
	case int32(33):
		if lookahead1 == int32('D') {
			state = uint16(762)
			goto next_state
		}
		return result
	case int32(34):
		if lookahead1 == int32('E') {
			state = uint16(900)
			goto next_state
		}
		return result
	case int32(35):
		if lookahead1 == int32('E') {
			state = uint16(559)
			goto next_state
		}
		return result
	case int32(36):
		if lookahead1 == int32('G') {
			state = uint16(738)
			goto next_state
		}
		if lookahead1 == int32('s') {
			state = uint16(967)
			goto next_state
		}
		return result
	case int32(37):
		if lookahead1 == int32('G') {
			state = uint16(297)
			goto next_state
		}
		return result
	case int32(38):
		if lookahead1 == int32('H') {
			state = uint16(619)
			goto next_state
		}
		return result
	case int32(39):
		if lookahead1 == int32('I') {
			state = uint16(177)
			goto next_state
		}
		return result
	case int32(40):
		if lookahead1 == int32('I') {
			state = uint16(538)
			goto next_state
		}
		return result
	case int32(41):
		if lookahead1 == int32('I') {
			state = uint16(836)
			goto next_state
		}
		if lookahead1 == int32('M') {
			state = uint16(103)
			goto next_state
		}
		return result
	case int32(42):
		if lookahead1 == int32('I') {
			state = uint16(837)
			goto next_state
		}
		return result
	case int32(43):
		if lookahead1 == int32('I') {
			state = uint16(838)
			goto next_state
		}
		return result
	case int32(44):
		if lookahead1 == int32('I') {
			state = uint16(840)
			goto next_state
		}
		return result
	case int32(45):
		if lookahead1 == int32('I') {
			state = uint16(841)
			goto next_state
		}
		return result
	case int32(46):
		if lookahead1 == int32('I') {
			state = uint16(541)
			goto next_state
		}
		return result
	case int32(47):
		if lookahead1 == int32('L') {
			state = uint16(240)
			goto next_state
		}
		if lookahead1 == int32('Q') {
			state = uint16(884)
			goto next_state
		}
		return result
	case int32(48):
		if lookahead1 == int32('L') {
			state = uint16(333)
			goto next_state
		}
		return result
	case int32(49):
		if lookahead1 == int32('L') {
			state = uint16(341)
			goto next_state
		}
		return result
	case int32(50):
		if lookahead1 == int32('L') {
			state = uint16(343)
			goto next_state
		}
		return result
	case int32(51):
		if lookahead1 == int32('M') {
			state = uint16(628)
			goto next_state
		}
		return result
	case int32(52):
		if lookahead1 == int32('M') {
			state = uint16(107)
			goto next_state
		}
		return result
	case int32(53):
		if lookahead1 == int32('M') {
			state = uint16(112)
			goto next_state
		}
		return result
	case int32(54):
		if lookahead1 == int32('M') {
			state = uint16(114)
			goto next_state
		}
		return result
	case int32(55):
		if lookahead1 == int32('M') {
			state = uint16(334)
			goto next_state
		}
		return result
	case int32(56):
		if lookahead1 == int32('M') {
			state = uint16(629)
			goto next_state
		}
		return result
	case int32(57):
		if lookahead1 == int32('M') {
			state = uint16(631)
			goto next_state
		}
		return result
	case int32(58):
		if lookahead1 == int32('N') {
			state = uint16(616)
			goto next_state
		}
		return result
	case int32(59):
		if lookahead1 == int32('N') {
			state = uint16(656)
			goto next_state
		}
		return result
	case int32(60):
		if lookahead1 == int32('O') {
			state = uint16(353)
			goto next_state
		}
		return result
	case int32(61):
		if lookahead1 == int32('P') {
			state = uint16(318)
			goto next_state
		}
		return result
	case int32(62):
		if lookahead1 == int32('P') {
			state = uint16(115)
			goto next_state
		}
		return result
	case int32(63):
		if lookahead1 == int32('P') {
			state = uint16(143)
			goto next_state
		}
		return result
	case int32(64):
		if lookahead1 == int32('Q') {
			state = uint16(885)
			goto next_state
		}
		return result
	case int32(65):
		if lookahead1 == int32('S') {
			state = uint16(382)
			goto next_state
		}
		return result
	case int32(66):
		if lookahead1 == int32('S') {
			state = uint16(400)
			goto next_state
		}
		return result
	case int32(67):
		if lookahead1 == int32('S') {
			state = uint16(665)
			goto next_state
		}
		return result
	case int32(68):
		if lookahead1 == int32('S') {
			state = uint16(827)
			goto next_state
		}
		return result
	case int32(69):
		if lookahead1 == int32('S') {
			state = uint16(424)
			goto next_state
		}
		return result
	case int32(70):
		if lookahead1 == int32('S') {
			state = uint16(649)
			goto next_state
		}
		return result
	case int32(71):
		if lookahead1 == int32('S') {
			state = uint16(655)
			goto next_state
		}
		return result
	case int32(72):
		if lookahead1 == int32('S') {
			state = uint16(659)
			goto next_state
		}
		return result
	case int32(73):
		if lookahead1 == int32('S') {
			state = uint16(662)
			goto next_state
		}
		return result
	case int32(74):
		if lookahead1 == int32('S') {
			state = uint16(663)
			goto next_state
		}
		return result
	case int32(75):
		if lookahead1 == int32('S') {
			state = uint16(664)
			goto next_state
		}
		return result
	case int32(76):
		if lookahead1 == int32('T') {
			state = uint16(438)
			goto next_state
		}
		return result
	case int32(77):
		if lookahead1 == int32('T') {
			state = uint16(748)
			goto next_state
		}
		return result
	case int32(78):
		if lookahead1 == int32('T') {
			state = uint16(911)
			goto next_state
		}
		return result
	case int32(79):
		if lookahead1 == int32('T') {
			state = uint16(443)
			goto next_state
		}
		return result
	case int32(80):
		if lookahead1 == int32('T') {
			state = uint16(444)
			goto next_state
		}
		return result
	case int32(81):
		if lookahead1 == int32('W') {
			state = uint16(639)
			goto next_state
		}
		return result
	case int32(82):
		if lookahead1 == int32('a') {
			state = uint16(780)
			goto next_state
		}
		if lookahead1 == int32('l') {
			state = uint16(398)
			goto next_state
		}
		if lookahead1 == int32('r') {
			state = uint16(608)
			goto next_state
		}
		return result
	case int32(83):
		if lookahead1 == int32('a') {
			state = uint16(487)
			goto next_state
		}
		if lookahead1 == int32('r') {
			state = uint16(97)
			goto next_state
		}
		return result
	case int32(84):
		if lookahead1 == int32('a') {
			state = uint16(365)
			goto next_state
		}
		if lookahead1 == int32('e') {
			state = uint16(528)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(557)
			goto next_state
		}
		if lookahead1 == int32('o') {
			state = uint16(620)
			goto next_state
		}
		return result
	case int32(85):
		if lookahead1 == int32('a') {
			state = uint16(408)
			goto next_state
		}
		if lookahead1 == int32('e') {
			state = uint16(201)
			goto next_state
		}
		return result
	case int32(86):
		if lookahead1 == int32('a') {
			state = uint16(154)
			goto next_state
		}
		if lookahead1 == int32('c') {
			state = uint16(388)
			goto next_state
		}
		return result
	case int32(87):
		if lookahead1 == int32('a') {
			state = uint16(901)
			goto next_state
		}
		return result
	case int32(88):
		if lookahead1 == int32('a') {
			state = uint16(160)
			goto next_state
		}
		return result
	case int32(89):
		if lookahead1 == int32('a') {
			state = uint16(48)
			goto next_state
		}
		return result
	case int32(90):
		if lookahead1 == int32('a') {
			state = uint16(55)
			goto next_state
		}
		return result
	case int32(91):
		if lookahead1 == int32('a') {
			state = uint16(463)
			goto next_state
		}
		return result
	case int32(92):
		if lookahead1 == int32('a') {
			state = uint16(962)
			goto next_state
		}
		return result
	case int32(93):
		if lookahead1 == int32('a') {
			state = uint16(374)
			goto next_state
		}
		return result
	case int32(94):
		if lookahead1 == int32('a') {
			state = uint16(768)
			goto next_state
		}
		if lookahead1 == int32('e') {
			state = uint16(423)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(208)
			goto next_state
		}
		return result
	case int32(95):
		if lookahead1 == int32('a') {
			state = uint16(768)
			goto next_state
		}
		if lookahead1 == int32('e') {
			state = uint16(423)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(208)
			goto next_state
		}
		if lookahead1 == int32('u') {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(96):
		if lookahead1 == int32('a') {
			state = uint16(486)
			goto next_state
		}
		return result
	case int32(97):
		if lookahead1 == int32('a') {
			state = uint16(166)
			goto next_state
		}
		return result
	case int32(98):
		if lookahead1 == int32('a') {
			state = uint16(782)
			goto next_state
		}
		return result
	case int32(99):
		if lookahead1 == int32('a') {
			state = uint16(529)
			goto next_state
		}
		if lookahead1 == int32('s') {
			state = uint16(152)
			goto next_state
		}
		if lookahead1 == int32('v') {
			state = uint16(277)
			goto next_state
		}
		return result
	case int32(100):
		if lookahead1 == int32('a') {
			state = uint16(570)
			goto next_state
		}
		return result
	case int32(101):
		if lookahead1 == int32('a') {
			state = uint16(461)
			goto next_state
		}
		return result
	case int32(102):
		if lookahead1 == int32('a') {
			state = uint16(544)
			goto next_state
		}
		return result
	case int32(103):
		if lookahead1 == int32('a') {
			state = uint16(673)
			goto next_state
		}
		return result
	case int32(104):
		if lookahead1 == int32('a') {
			state = uint16(909)
			goto next_state
		}
		if lookahead1 == int32('h') {
			state = uint16(427)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(209)
			goto next_state
		}
		return result
	case int32(105):
		if lookahead1 == int32('a') {
			state = uint16(792)
			goto next_state
		}
		return result
	case int32(106):
		if lookahead1 == int32('a') {
			state = uint16(555)
			goto next_state
		}
		return result
	case int32(107):
		if lookahead1 == int32('a') {
			state = uint16(675)
			goto next_state
		}
		return result
	case int32(108):
		if lookahead1 == int32('a') {
			state = uint16(371)
			goto next_state
		}
		return result
	case int32(109):
		if lookahead1 == int32('a') {
			state = uint16(465)
			goto next_state
		}
		return result
	case int32(110):
		if lookahead1 == int32('a') {
			state = uint16(563)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(117)
			goto next_state
		}
		return result
	case int32(111):
		if lookahead1 == int32('a') {
			state = uint16(563)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(117)
			goto next_state
		}
		if lookahead1 == int32('u') {
			state = uint16(238)
			goto next_state
		}
		return result
	case int32(112):
		if lookahead1 == int32('a') {
			state = uint16(676)
			goto next_state
		}
		return result
	case int32(113):
		if lookahead1 == int32('a') {
			state = uint16(743)
			goto next_state
		}
		return result
	case int32(114):
		if lookahead1 == int32('a') {
			state = uint16(678)
			goto next_state
		}
		return result
	case int32(115):
		if lookahead1 == int32('a') {
			state = uint16(793)
			goto next_state
		}
		return result
	case int32(116):
		if lookahead1 == int32('a') {
			state = uint16(469)
			goto next_state
		}
		return result
	case int32(117):
		if lookahead1 == int32('a') {
			state = uint16(560)
			goto next_state
		}
		return result
	case int32(118):
		if lookahead1 == int32('a') {
			state = uint16(470)
			goto next_state
		}
		return result
	case int32(119):
		if lookahead1 == int32('a') {
			state = uint16(730)
			goto next_state
		}
		return result
	case int32(120):
		if lookahead1 == int32('a') {
			state = uint16(471)
			goto next_state
		}
		return result
	case int32(121):
		if lookahead1 == int32('a') {
			state = uint16(755)
			goto next_state
		}
		return result
	case int32(122):
		if lookahead1 == int32('a') {
			state = uint16(472)
			goto next_state
		}
		return result
	case int32(123):
		if lookahead1 == int32('a') {
			state = uint16(473)
			goto next_state
		}
		return result
	case int32(124):
		if lookahead1 == int32('a') {
			state = uint16(462)
			goto next_state
		}
		return result
	case int32(125):
		if lookahead1 == int32('a') {
			state = uint16(572)
			goto next_state
		}
		return result
	case int32(126):
		if lookahead1 == int32('a') {
			state = uint16(727)
			goto next_state
		}
		return result
	case int32(127):
		if lookahead1 == int32('a') {
			state = uint16(161)
			goto next_state
		}
		if lookahead1 == int32('o') {
			state = uint16(759)
			goto next_state
		}
		return result
	case int32(128):
		if lookahead1 == int32('a') {
			state = uint16(681)
			goto next_state
		}
		if lookahead1 == int32('o') {
			state = uint16(895)
			goto next_state
		}
		return result
	case int32(129):
		if lookahead1 == int32('a') {
			state = uint16(737)
			goto next_state
		}
		return result
	case int32(130):
		if lookahead1 == int32('a') {
			state = uint16(568)
			goto next_state
		}
		return result
	case int32(131):
		if lookahead1 == int32('a') {
			state = uint16(797)
			goto next_state
		}
		return result
	case int32(132):
		if lookahead1 == int32('a') {
			state = uint16(476)
			goto next_state
		}
		return result
	case int32(133):
		if lookahead1 == int32('a') {
			state = uint16(574)
			goto next_state
		}
		return result
	case int32(134):
		if lookahead1 == int32('a') {
			state = uint16(490)
			goto next_state
		}
		return result
	case int32(135):
		if lookahead1 == int32('a') {
			state = uint16(684)
			goto next_state
		}
		if lookahead1 == int32('o') {
			state = uint16(895)
			goto next_state
		}
		return result
	case int32(136):
		if lookahead1 == int32('a') {
			state = uint16(749)
			goto next_state
		}
		if lookahead1 == int32('e') {
			state = uint16(215)
			goto next_state
		}
		return result
	case int32(137):
		if lookahead1 == int32('a') {
			state = uint16(830)
			goto next_state
		}
		return result
	case int32(138):
		if lookahead1 == int32('a') {
			state = uint16(220)
			goto next_state
		}
		return result
	case int32(139):
		if lookahead1 == int32('a') {
			state = uint16(688)
			goto next_state
		}
		return result
	case int32(140):
		if lookahead1 == int32('a') {
			state = uint16(375)
			goto next_state
		}
		return result
	case int32(141):
		if lookahead1 == int32('a') {
			state = uint16(491)
			goto next_state
		}
		return result
	case int32(142):
		if lookahead1 == int32('a') {
			state = uint16(796)
			goto next_state
		}
		return result
	case int32(143):
		if lookahead1 == int32('a') {
			state = uint16(795)
			goto next_state
		}
		return result
	case int32(144):
		if lookahead1 == int32('a') {
			state = uint16(760)
			goto next_state
		}
		return result
	case int32(145):
		if lookahead1 == int32('a') {
			state = uint16(842)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(181)
			goto next_state
		}
		if lookahead1 == int32('x') {
			state = uint16(690)
			goto next_state
		}
		return result
	case int32(146):
		if lookahead1 == int32('a') {
			state = uint16(689)
			goto next_state
		}
		return result
	case int32(147):
		if lookahead1 == int32('a') {
			state = uint16(494)
			goto next_state
		}
		return result
	case int32(148):
		if lookahead1 == int32('a') {
			state = uint16(224)
			goto next_state
		}
		return result
	case int32(149):
		if lookahead1 == int32('a') {
			state = uint16(605)
			goto next_state
		}
		return result
	case int32(150):
		if lookahead1 == int32('a') {
			state = uint16(606)
			goto next_state
		}
		return result
	case int32(151):
		if lookahead1 == int32('a') {
			state = uint16(607)
			goto next_state
		}
		return result
	case int32(152):
		if lookahead1 == int32('a') {
			state = uint16(156)
			goto next_state
		}
		return result
	case int32(153):
		if lookahead1 == int32('b') {
			state = uint16(274)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(399)
			goto next_state
		}
		if lookahead1 == int32('p') {
			state = uint16(781)
			goto next_state
		}
		return result
	case int32(154):
		if lookahead1 == int32('b') {
			state = uint16(479)
			goto next_state
		}
		return result
	case int32(155):
		if lookahead1 == int32('b') {
			state = uint16(484)
			goto next_state
		}
		return result
	case int32(156):
		if lookahead1 == int32('b') {
			state = uint16(508)
			goto next_state
		}
		return result
	case int32(157):
		if lookahead1 == int32('c') {
			state = uint16(381)
			goto next_state
		}
		if lookahead1 == int32('e') {
			state = uint16(89)
			goto next_state
		}
		return result
	case int32(158):
		if lookahead1 == int32('c') {
			state = uint16(1010)
			goto next_state
		}
		return result
	case int32(159):
		if lookahead1 == int32('c') {
			state = uint16(1008)
			goto next_state
		}
		return result
	case int32(160):
		if lookahead1 == int32('c') {
			state = uint16(454)
			goto next_state
		}
		if lookahead1 == int32('r') {
			state = uint16(1062)
			goto next_state
		}
		return result
	case int32(161):
		if lookahead1 == int32('c') {
			state = uint16(455)
			goto next_state
		}
		return result
	case int32(162):
		i4 = uint32(0)
		for {
			if !(uint64(i4) < libc.Uint64FromInt64(32)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token4[i4]) == lookahead1 {
				state = map_token4[i4+uint32(1)]
				goto next_state
			}
			goto _25
		_25:
			;
			i4 = i4 + uint32(2)
		}
		return result
	case int32(163):
		i5 = uint32(0)
		for {
			if !(uint64(i5) < libc.Uint64FromInt64(32)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token5[i5]) == lookahead1 {
				state = map_token5[i5+uint32(1)]
				goto next_state
			}
			goto _26
		_26:
			;
			i5 = i5 + uint32(2)
		}
		return result
	case int32(164):
		if lookahead1 == int32('c') {
			state = uint16(478)
			goto next_state
		}
		return result
	case int32(165):
		if lookahead1 == int32('c') {
			state = uint16(394)
			goto next_state
		}
		return result
	case int32(166):
		if lookahead1 == int32('c') {
			state = uint16(832)
			goto next_state
		}
		return result
	case int32(167):
		if lookahead1 == int32('c') {
			state = uint16(445)
			goto next_state
		}
		return result
	case int32(168):
		if lookahead1 == int32('c') {
			state = uint16(92)
			goto next_state
		}
		return result
	case int32(169):
		if lookahead1 == int32('c') {
			state = uint16(806)
			goto next_state
		}
		return result
	case int32(170):
		if lookahead1 == int32('c') {
			state = uint16(255)
			goto next_state
		}
		return result
	case int32(171):
		if lookahead1 == int32('c') {
			state = uint16(257)
			goto next_state
		}
		return result
	case int32(172):
		if lookahead1 == int32('c') {
			state = uint16(457)
			goto next_state
		}
		return result
	case int32(173):
		if lookahead1 == int32('c') {
			state = uint16(458)
			goto next_state
		}
		return result
	case int32(174):
		if lookahead1 == int32('c') {
			state = uint16(415)
			goto next_state
		}
		return result
	case int32(175):
		if lookahead1 == int32('c') {
			state = uint16(417)
			goto next_state
		}
		return result
	case int32(176):
		if lookahead1 == int32('c') {
			state = uint16(418)
			goto next_state
		}
		return result
	case int32(177):
		if lookahead1 == int32('c') {
			state = uint16(636)
			goto next_state
		}
		return result
	case int32(178):
		if lookahead1 == int32('c') {
			state = uint16(337)
			goto next_state
		}
		return result
	case int32(179):
		if lookahead1 == int32('c') {
			state = uint16(449)
			goto next_state
		}
		return result
	case int32(180):
		if lookahead1 == int32('c') {
			state = uint16(392)
			goto next_state
		}
		return result
	case int32(181):
		if lookahead1 == int32('c') {
			state = uint16(393)
			goto next_state
		}
		return result
	case int32(182):
		if lookahead1 == int32('d') {
			state = uint16(1047)
			goto next_state
		}
		if lookahead1 == int32('p') {
			state = uint16(488)
			goto next_state
		}
		return result
	case int32(183):
		if lookahead1 == int32('d') {
			state = uint16(1059)
			goto next_state
		}
		return result
	case int32(184):
		if lookahead1 == int32('d') {
			state = uint16(1021)
			goto next_state
		}
		return result
	case int32(185):
		if lookahead1 == int32('d') {
			state = uint16(933)
			goto next_state
		}
		return result
	case int32(186):
		if lookahead1 == int32('d') {
			state = uint16(959)
			goto next_state
		}
		return result
	case int32(187):
		if lookahead1 == int32('d') {
			state = uint16(963)
			goto next_state
		}
		return result
	case int32(188):
		if lookahead1 == int32('d') {
			state = uint16(936)
			goto next_state
		}
		return result
	case int32(189):
		if lookahead1 == int32('d') {
			state = uint16(954)
			goto next_state
		}
		return result
	case int32(190):
		if lookahead1 == int32('d') {
			state = uint16(37)
			goto next_state
		}
		return result
	case int32(191):
		if lookahead1 == int32('d') {
			state = uint16(926)
			goto next_state
		}
		return result
	case int32(192):
		if lookahead1 == int32('d') {
			state = uint16(949)
			goto next_state
		}
		return result
	case int32(193):
		if lookahead1 == int32('d') {
			state = uint16(950)
			goto next_state
		}
		return result
	case int32(194):
		if lookahead1 == int32('d') {
			state = uint16(976)
			goto next_state
		}
		return result
	case int32(195):
		if lookahead1 == int32('d') {
			state = uint16(983)
			goto next_state
		}
		return result
	case int32(196):
		if lookahead1 == int32('d') {
			state = uint16(930)
			goto next_state
		}
		return result
	case int32(197):
		if lookahead1 == int32('d') {
			state = uint16(979)
			goto next_state
		}
		return result
	case int32(198):
		if lookahead1 == int32('d') {
			state = uint16(982)
			goto next_state
		}
		return result
	case int32(199):
		if lookahead1 == int32('d') {
			state = uint16(985)
			goto next_state
		}
		return result
	case int32(200):
		if lookahead1 == int32('d') {
			state = uint16(984)
			goto next_state
		}
		return result
	case int32(201):
		if lookahead1 == int32('d') {
			state = uint16(1046)
			goto next_state
		}
		if lookahead1 == int32('p') {
			state = uint16(488)
			goto next_state
		}
		return result
	case int32(202):
		if lookahead1 == int32('d') {
			state = uint16(1013)
			goto next_state
		}
		return result
	case int32(203):
		if lookahead1 == int32('d') {
			state = uint16(52)
			goto next_state
		}
		if lookahead1 == int32('r') {
			state = uint16(1015)
			goto next_state
		}
		return result
	case int32(204):
		if lookahead1 == int32('d') {
			state = uint16(52)
			goto next_state
		}
		if lookahead1 == int32('r') {
			state = uint16(42)
			goto next_state
		}
		return result
	case int32(205):
		if lookahead1 == int32('d') {
			state = uint16(67)
			goto next_state
		}
		return result
	case int32(206):
		if lookahead1 == int32('d') {
			state = uint16(43)
			goto next_state
		}
		return result
	case int32(207):
		if lookahead1 == int32('d') {
			state = uint16(281)
			goto next_state
		}
		if lookahead1 == int32('m') {
			state = uint16(680)
			goto next_state
		}
		if lookahead1 == int32('t') {
			state = uint16(276)
			goto next_state
		}
		return result
	case int32(208):
		if lookahead1 == int32('d') {
			state = uint16(234)
			goto next_state
		}
		return result
	case int32(209):
		if lookahead1 == int32('d') {
			state = uint16(817)
			goto next_state
		}
		return result
	case int32(210):
		if lookahead1 == int32('d') {
			state = uint16(772)
			goto next_state
		}
		return result
	case int32(211):
		if lookahead1 == int32('d') {
			state = uint16(774)
			goto next_state
		}
		return result
	case int32(212):
		if lookahead1 == int32('d') {
			state = uint16(778)
			goto next_state
		}
		return result
	case int32(213):
		if lookahead1 == int32('d') {
			state = uint16(268)
			goto next_state
		}
		return result
	case int32(214):
		if lookahead1 == int32('d') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(215):
		if lookahead1 == int32('d') {
			state = uint16(320)
			goto next_state
		}
		return result
	case int32(216):
		if lookahead1 == int32('d') {
			state = uint16(53)
			goto next_state
		}
		return result
	case int32(217):
		if lookahead1 == int32('d') {
			state = uint16(734)
			goto next_state
		}
		return result
	case int32(218):
		if lookahead1 == int32('d') {
			state = uint16(54)
			goto next_state
		}
		return result
	case int32(219):
		if lookahead1 == int32('d') {
			state = uint16(278)
			goto next_state
		}
		return result
	case int32(220):
		if lookahead1 == int32('d') {
			state = uint16(306)
			goto next_state
		}
		return result
	case int32(221):
		if lookahead1 == int32('d') {
			state = uint16(310)
			goto next_state
		}
		return result
	case int32(222):
		if lookahead1 == int32('d') {
			state = uint16(313)
			goto next_state
		}
		return result
	case int32(223):
		if lookahead1 == int32('d') {
			state = uint16(314)
			goto next_state
		}
		return result
	case int32(224):
		if lookahead1 == int32('d') {
			state = uint16(316)
			goto next_state
		}
		return result
	case int32(225):
		if lookahead1 == int32('d') {
			state = uint16(44)
			goto next_state
		}
		return result
	case int32(226):
		if lookahead1 == int32('d') {
			state = uint16(45)
			goto next_state
		}
		return result
	case int32(227):
		if lookahead1 == int32('d') {
			state = uint16(340)
			goto next_state
		}
		return result
	case int32(228):
		if lookahead1 == int32('d') {
			state = uint16(347)
			goto next_state
		}
		if lookahead1 == int32('q') {
			state = uint16(869)
			goto next_state
		}
		return result
	case int32(229):
		if lookahead1 == int32('d') {
			state = uint16(30)
			goto next_state
		}
		return result
	case int32(230):
		if lookahead1 == int32('e') {
			state = uint16(511)
			goto next_state
		}
		if lookahead1 == int32('r') {
			state = uint16(231)
			goto next_state
		}
		return result
	case int32(231):
		if lookahead1 == int32('e') {
			state = uint16(302)
			goto next_state
		}
		return result
	case int32(232):
		if lookahead1 == int32('e') {
			state = uint16(21)
			goto next_state
		}
		return result
	case int32(233):
		if lookahead1 == int32('e') {
			state = uint16(1049)
			goto next_state
		}
		return result
	case int32(234):
		if lookahead1 == int32('e') {
			state = uint16(920)
			goto next_state
		}
		return result
	case int32(235):
		if lookahead1 == int32('e') {
			state = uint16(1067)
			goto next_state
		}
		return result
	case int32(236):
		if lookahead1 == int32('e') {
			state = uint16(977)
			goto next_state
		}
		return result
	case int32(237):
		if lookahead1 == int32('e') {
			state = uint16(1011)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(813)
			goto next_state
		}
		return result
	case int32(238):
		if lookahead1 == int32('e') {
			state = uint16(999)
			goto next_state
		}
		return result
	case int32(239):
		if lookahead1 == int32('e') {
			state = uint16(8)
			goto next_state
		}
		return result
	case int32(240):
		if lookahead1 == int32('e') {
			state = uint16(890)
			goto next_state
		}
		return result
	case int32(241):
		if lookahead1 == int32('e') {
			state = uint16(899)
			goto next_state
		}
		return result
	case int32(242):
		if lookahead1 == int32('e') {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(243):
		if lookahead1 == int32('e') {
			state = uint16(1051)
			goto next_state
		}
		return result
	case int32(244):
		if lookahead1 == int32('e') {
			state = uint16(1058)
			goto next_state
		}
		return result
	case int32(245):
		if lookahead1 == int32('e') {
			state = uint16(32)
			goto next_state
		}
		return result
	case int32(246):
		if lookahead1 == int32('e') {
			state = uint16(1055)
			goto next_state
		}
		return result
	case int32(247):
		if lookahead1 == int32('e') {
			state = uint16(1057)
			goto next_state
		}
		return result
	case int32(248):
		if lookahead1 == int32('e') {
			state = uint16(1061)
			goto next_state
		}
		return result
	case int32(249):
		if lookahead1 == int32('e') {
			state = uint16(1012)
			goto next_state
		}
		return result
	case int32(250):
		if lookahead1 == int32('e') {
			state = uint16(932)
			goto next_state
		}
		return result
	case int32(251):
		if lookahead1 == int32('e') {
			state = uint16(994)
			goto next_state
		}
		return result
	case int32(252):
		if lookahead1 == int32('e') {
			state = uint16(1063)
			goto next_state
		}
		return result
	case int32(253):
		if lookahead1 == int32('e') {
			state = uint16(64)
			goto next_state
		}
		return result
	case int32(254):
		if lookahead1 == int32('e') {
			state = uint16(968)
			goto next_state
		}
		return result
	case int32(255):
		if lookahead1 == int32('e') {
			state = uint16(61)
			goto next_state
		}
		return result
	case int32(256):
		if lookahead1 == int32('e') {
			state = uint16(992)
			goto next_state
		}
		return result
	case int32(257):
		if lookahead1 == int32('e') {
			state = uint16(951)
			goto next_state
		}
		return result
	case int32(258):
		if lookahead1 == int32('e') {
			state = uint16(945)
			goto next_state
		}
		return result
	case int32(259):
		if lookahead1 == int32('e') {
			state = uint16(1069)
			goto next_state
		}
		return result
	case int32(260):
		if lookahead1 == int32('e') {
			state = uint16(929)
			goto next_state
		}
		return result
	case int32(261):
		if lookahead1 == int32('e') {
			state = uint16(941)
			goto next_state
		}
		return result
	case int32(262):
		if lookahead1 == int32('e') {
			state = uint16(946)
			goto next_state
		}
		return result
	case int32(263):
		if lookahead1 == int32('e') {
			state = uint16(1008)
			goto next_state
		}
		return result
	case int32(264):
		if lookahead1 == int32('e') {
			state = uint16(1013)
			goto next_state
		}
		return result
	case int32(265):
		if lookahead1 == int32('e') {
			state = uint16(489)
			goto next_state
		}
		return result
	case int32(266):
		if lookahead1 == int32('e') {
			state = uint16(59)
			goto next_state
		}
		return result
	case int32(267):
		if lookahead1 == int32('e') {
			state = uint16(204)
			goto next_state
		}
		return result
	case int32(268):
		if lookahead1 == int32('e') {
			state = uint16(31)
			goto next_state
		}
		return result
	case int32(269):
		if lookahead1 == int32('e') {
			state = uint16(519)
			goto next_state
		}
		if lookahead1 == int32('r') {
			state = uint16(111)
			goto next_state
		}
		return result
	case int32(270):
		if lookahead1 == int32('e') {
			state = uint16(561)
			goto next_state
		}
		if lookahead1 == int32('h') {
			state = uint16(130)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(553)
			goto next_state
		}
		if lookahead1 == int32('l') {
			state = uint16(87)
			goto next_state
		}
		if lookahead1 == int32('u') {
			state = uint16(724)
			goto next_state
		}
		return result
	case int32(271):
		if lookahead1 == int32('e') {
			state = uint16(77)
			goto next_state
		}
		return result
	case int32(272):
		if lookahead1 == int32('e') {
			state = uint16(205)
			goto next_state
		}
		return result
	case int32(273):
		if lookahead1 == int32('e') {
			state = uint16(360)
			goto next_state
		}
		return result
	case int32(274):
		if lookahead1 == int32('e') {
			state = uint16(700)
			goto next_state
		}
		return result
	case int32(275):
		if lookahead1 == int32('e') {
			state = uint16(752)
			goto next_state
		}
		return result
	case int32(276):
		if lookahead1 == int32('e') {
			state = uint16(540)
			goto next_state
		}
		return result
	case int32(277):
		if lookahead1 == int32('e') {
			state = uint16(725)
			goto next_state
		}
		return result
	case int32(278):
		if lookahead1 == int32('e') {
			state = uint16(701)
			goto next_state
		}
		return result
	case int32(279):
		if lookahead1 == int32('e') {
			state = uint16(216)
			goto next_state
		}
		return result
	case int32(280):
		if lookahead1 == int32('e') {
			state = uint16(531)
			goto next_state
		}
		return result
	case int32(281):
		if lookahead1 == int32('e') {
			state = uint16(595)
			goto next_state
		}
		return result
	case int32(282):
		if lookahead1 == int32('e') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(283):
		if lookahead1 == int32('e') {
			state = uint16(534)
			goto next_state
		}
		return result
	case int32(284):
		if lookahead1 == int32('e') {
			state = uint16(187)
			goto next_state
		}
		return result
	case int32(285):
		if lookahead1 == int32('e') {
			state = uint16(188)
			goto next_state
		}
		return result
	case int32(286):
		if lookahead1 == int32('e') {
			state = uint16(464)
			goto next_state
		}
		return result
	case int32(287):
		if lookahead1 == int32('e') {
			state = uint16(513)
			goto next_state
		}
		return result
	case int32(288):
		if lookahead1 == int32('e') {
			state = uint16(206)
			goto next_state
		}
		return result
	case int32(289):
		if lookahead1 == int32('e') {
			state = uint16(514)
			goto next_state
		}
		return result
	case int32(290):
		if lookahead1 == int32('e') {
			state = uint16(189)
			goto next_state
		}
		return result
	case int32(291):
		if lookahead1 == int32('e') {
			state = uint16(466)
			goto next_state
		}
		return result
	case int32(292):
		if lookahead1 == int32('e') {
			state = uint16(515)
			goto next_state
		}
		return result
	case int32(293):
		if lookahead1 == int32('e') {
			state = uint16(733)
			goto next_state
		}
		return result
	case int32(294):
		if lookahead1 == int32('e') {
			state = uint16(467)
			goto next_state
		}
		return result
	case int32(295):
		if lookahead1 == int32('e') {
			state = uint16(516)
			goto next_state
		}
		return result
	case int32(296):
		if lookahead1 == int32('e') {
			state = uint16(468)
			goto next_state
		}
		return result
	case int32(297):
		if lookahead1 == int32('e') {
			state = uint16(517)
			goto next_state
		}
		return result
	case int32(298):
		if lookahead1 == int32('e') {
			state = uint16(803)
			goto next_state
		}
		return result
	case int32(299):
		if lookahead1 == int32('e') {
			state = uint16(702)
			goto next_state
		}
		return result
	case int32(300):
		if lookahead1 == int32('e') {
			state = uint16(190)
			goto next_state
		}
		return result
	case int32(301):
		if lookahead1 == int32('e') {
			state = uint16(539)
			goto next_state
		}
		return result
	case int32(302):
		if lookahead1 == int32('e') {
			state = uint16(547)
			goto next_state
		}
		if lookahead1 == int32('y') {
			state = uint16(1054)
			goto next_state
		}
		return result
	case int32(303):
		if lookahead1 == int32('e') {
			state = uint16(703)
			goto next_state
		}
		return result
	case int32(304):
		if lookahead1 == int32('e') {
			state = uint16(732)
			goto next_state
		}
		return result
	case int32(305):
		if lookahead1 == int32('e') {
			state = uint16(569)
			goto next_state
		}
		return result
	case int32(306):
		if lookahead1 == int32('e') {
			state = uint16(704)
			goto next_state
		}
		return result
	case int32(307):
		if lookahead1 == int32('e') {
			state = uint16(705)
			goto next_state
		}
		return result
	case int32(308):
		if lookahead1 == int32('e') {
			state = uint16(706)
			goto next_state
		}
		return result
	case int32(309):
		if lookahead1 == int32('e') {
			state = uint16(824)
			goto next_state
		}
		return result
	case int32(310):
		if lookahead1 == int32('e') {
			state = uint16(766)
			goto next_state
		}
		return result
	case int32(311):
		if lookahead1 == int32('e') {
			state = uint16(710)
			goto next_state
		}
		return result
	case int32(312):
		if lookahead1 == int32('e') {
			state = uint16(713)
			goto next_state
		}
		return result
	case int32(313):
		if lookahead1 == int32('e') {
			state = uint16(715)
			goto next_state
		}
		return result
	case int32(314):
		if lookahead1 == int32('e') {
			state = uint16(716)
			goto next_state
		}
		return result
	case int32(315):
		if lookahead1 == int32('e') {
			state = uint16(717)
			goto next_state
		}
		return result
	case int32(316):
		if lookahead1 == int32('e') {
			state = uint16(719)
			goto next_state
		}
		return result
	case int32(317):
		if lookahead1 == int32('e') {
			state = uint16(262)
			goto next_state
		}
		return result
	case int32(318):
		if lookahead1 == int32('e') {
			state = uint16(735)
			goto next_state
		}
		return result
	case int32(319):
		if lookahead1 == int32('e') {
			state = uint16(169)
			goto next_state
		}
		return result
	case int32(320):
		if lookahead1 == int32('e') {
			state = uint16(280)
			goto next_state
		}
		return result
	case int32(321):
		if lookahead1 == int32('e') {
			state = uint16(787)
			goto next_state
		}
		return result
	case int32(322):
		if lookahead1 == int32('e') {
			state = uint16(742)
			goto next_state
		}
		return result
	case int32(323):
		if lookahead1 == int32('e') {
			state = uint16(482)
			goto next_state
		}
		return result
	case int32(324):
		if lookahead1 == int32('e') {
			state = uint16(597)
			goto next_state
		}
		return result
	case int32(325):
		if lookahead1 == int32('e') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(326):
		if lookahead1 == int32('e') {
			state = uint16(788)
			goto next_state
		}
		return result
	case int32(327):
		if lookahead1 == int32('e') {
			state = uint16(215)
			goto next_state
		}
		return result
	case int32(328):
		if lookahead1 == int32('e') {
			state = uint16(580)
			goto next_state
		}
		return result
	case int32(329):
		if lookahead1 == int32('e') {
			state = uint16(584)
			goto next_state
		}
		return result
	case int32(330):
		if lookahead1 == int32('e') {
			state = uint16(589)
			goto next_state
		}
		return result
	case int32(331):
		if lookahead1 == int32('e') {
			state = uint16(591)
			goto next_state
		}
		return result
	case int32(332):
		if lookahead1 == int32('e') {
			state = uint16(592)
			goto next_state
		}
		return result
	case int32(333):
		if lookahead1 == int32('e') {
			state = uint16(891)
			goto next_state
		}
		return result
	case int32(334):
		if lookahead1 == int32('e') {
			state = uint16(530)
			goto next_state
		}
		return result
	case int32(335):
		if lookahead1 == int32('e') {
			state = uint16(218)
			goto next_state
		}
		return result
	case int32(336):
		if lookahead1 == int32('e') {
			state = uint16(532)
			goto next_state
		}
		return result
	case int32(337):
		if lookahead1 == int32('e') {
			state = uint16(598)
			goto next_state
		}
		return result
	case int32(338):
		if lookahead1 == int32('e') {
			state = uint16(225)
			goto next_state
		}
		return result
	case int32(339):
		if lookahead1 == int32('e') {
			state = uint16(758)
			goto next_state
		}
		return result
	case int32(340):
		if lookahead1 == int32('e') {
			state = uint16(336)
			goto next_state
		}
		return result
	case int32(341):
		if lookahead1 == int32('e') {
			state = uint16(892)
			goto next_state
		}
		return result
	case int32(342):
		if lookahead1 == int32('e') {
			state = uint16(226)
			goto next_state
		}
		return result
	case int32(343):
		if lookahead1 == int32('e') {
			state = uint16(893)
			goto next_state
		}
		return result
	case int32(344):
		if lookahead1 == int32('e') {
			state = uint16(79)
			goto next_state
		}
		return result
	case int32(345):
		if lookahead1 == int32('e') {
			state = uint16(63)
			goto next_state
		}
		return result
	case int32(346):
		if lookahead1 == int32('e') {
			state = uint16(761)
			goto next_state
		}
		return result
	case int32(347):
		if lookahead1 == int32('e') {
			state = uint16(603)
			goto next_state
		}
		return result
	case int32(348):
		if lookahead1 == int32('e') {
			state = uint16(763)
			goto next_state
		}
		return result
	case int32(349):
		if lookahead1 == int32('e') {
			state = uint16(227)
			goto next_state
		}
		return result
	case int32(350):
		if lookahead1 == int32('e') {
			state = uint16(764)
			goto next_state
		}
		return result
	case int32(351):
		if lookahead1 == int32('e') {
			state = uint16(765)
			goto next_state
		}
		return result
	case int32(352):
		if lookahead1 == int32('e') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(353):
		if lookahead1 == int32('f') {
			state = uint16(81)
			goto next_state
		}
		return result
	case int32(354):
		if lookahead1 == int32('f') {
			state = uint16(358)
			goto next_state
		}
		return result
	case int32(355):
		if lookahead1 == int32('f') {
			state = uint16(474)
			goto next_state
		}
		return result
	case int32(356):
		if lookahead1 == int32('f') {
			state = uint16(406)
			goto next_state
		}
		return result
	case int32(357):
		if lookahead1 == int32('f') {
			state = uint16(420)
			goto next_state
		}
		return result
	case int32(358):
		if lookahead1 == int32('f') {
			state = uint16(319)
			goto next_state
		}
		return result
	case int32(359):
		if lookahead1 == int32('f') {
			state = uint16(448)
			goto next_state
		}
		return result
	case int32(360):
		if lookahead1 == int32('f') {
			state = uint16(305)
			goto next_state
		}
		return result
	case int32(361):
		if lookahead1 == int32('f') {
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(362):
		if lookahead1 == int32('f') {
			state = uint16(24)
			goto next_state
		}
		return result
	case int32(363):
		if lookahead1 == int32('g') {
			state = uint16(34)
			goto next_state
		}
		return result
	case int32(364):
		if lookahead1 == int32('g') {
			state = uint16(387)
			goto next_state
		}
		return result
	case int32(365):
		if lookahead1 == int32('g') {
			state = uint16(404)
			goto next_state
		}
		if lookahead1 == int32('p') {
			state = uint16(76)
			goto next_state
		}
		return result
	case int32(366):
		if lookahead1 == int32('g') {
			state = uint16(380)
			goto next_state
		}
		return result
	case int32(367):
		if lookahead1 == int32('g') {
			state = uint16(246)
			goto next_state
		}
		return result
	case int32(368):
		if lookahead1 == int32('g') {
			state = uint16(284)
			goto next_state
		}
		return result
	case int32(369):
		if lookahead1 == int32('g') {
			state = uint16(328)
			goto next_state
		}
		return result
	case int32(370):
		if lookahead1 == int32('g') {
			state = uint16(908)
			goto next_state
		}
		return result
	case int32(371):
		if lookahead1 == int32('g') {
			state = uint16(410)
			goto next_state
		}
		return result
	case int32(372):
		if lookahead1 == int32('g') {
			state = uint16(483)
			goto next_state
		}
		return result
	case int32(373):
		if lookahead1 == int32('g') {
			state = uint16(740)
			goto next_state
		}
		return result
	case int32(374):
		if lookahead1 == int32('g') {
			state = uint16(632)
			goto next_state
		}
		return result
	case int32(375):
		if lookahead1 == int32('g') {
			state = uint16(633)
			goto next_state
		}
		return result
	case int32(376):
		if lookahead1 == int32('g') {
			state = uint16(883)
			goto next_state
		}
		return result
	case int32(377):
		if lookahead1 == int32('g') {
			state = uint16(332)
			goto next_state
		}
		return result
	case int32(378):
		if lookahead1 == int32('g') {
			state = uint16(389)
			goto next_state
		}
		return result
	case int32(379):
		if lookahead1 == int32('h') {
			state = uint16(974)
			goto next_state
		}
		return result
	case int32(380):
		if lookahead1 == int32('h') {
			state = uint16(801)
			goto next_state
		}
		return result
	case int32(381):
		if lookahead1 == int32('h') {
			state = uint16(599)
			goto next_state
		}
		return result
	case int32(382):
		if lookahead1 == int32('h') {
			state = uint16(435)
			goto next_state
		}
		return result
	case int32(383):
		if lookahead1 == int32('h') {
			state = uint16(610)
			goto next_state
		}
		return result
	case int32(384):
		if lookahead1 == int32('h') {
			state = uint16(321)
			goto next_state
		}
		return result
	case int32(385):
		if lookahead1 == int32('h') {
			state = uint16(139)
			goto next_state
		}
		return result
	case int32(386):
		if lookahead1 == int32('h') {
			state = uint16(139)
			goto next_state
		}
		if lookahead1 == int32('u') {
			state = uint16(692)
			goto next_state
		}
		return result
	case int32(387):
		if lookahead1 == int32('h') {
			state = uint16(828)
			goto next_state
		}
		return result
	case int32(388):
		if lookahead1 == int32('h') {
			state = uint16(133)
			goto next_state
		}
		return result
	case int32(389):
		if lookahead1 == int32('h') {
			state = uint16(849)
			goto next_state
		}
		return result
	case int32(390):
		if lookahead1 == int32('h') {
			state = uint16(146)
			goto next_state
		}
		return result
	case int32(391):
		if lookahead1 == int32('h') {
			state = uint16(149)
			goto next_state
		}
		return result
	case int32(392):
		if lookahead1 == int32('h') {
			state = uint16(150)
			goto next_state
		}
		return result
	case int32(393):
		if lookahead1 == int32('h') {
			state = uint16(151)
			goto next_state
		}
		return result
	case int32(394):
		if lookahead1 == int32('h') {
			state = uint16(46)
			goto next_state
		}
		return result
	case int32(395):
		if lookahead1 == int32('i') {
			state = uint16(720)
			goto next_state
		}
		if lookahead1 == int32('l') {
			state = uint16(98)
			goto next_state
		}
		if lookahead1 == int32('o') {
			state = uint16(556)
			goto next_state
		}
		if lookahead1 == int32('r') {
			state = uint16(647)
			goto next_state
		}
		if lookahead1 == int32('u') {
			state = uint16(779)
			goto next_state
		}
		if lookahead1 == int32('y') {
			state = uint16(102)
			goto next_state
		}
		return result
	case int32(396):
		if lookahead1 == int32('i') {
			state = uint16(720)
			goto next_state
		}
		if lookahead1 == int32('l') {
			state = uint16(98)
			goto next_state
		}
		if lookahead1 == int32('o') {
			state = uint16(556)
			goto next_state
		}
		if lookahead1 == int32('r') {
			state = uint16(646)
			goto next_state
		}
		if lookahead1 == int32('u') {
			state = uint16(779)
			goto next_state
		}
		if lookahead1 == int32('y') {
			state = uint16(102)
			goto next_state
		}
		return result
	case int32(397):
		if lookahead1 == int32('i') {
			state = uint16(99)
			goto next_state
		}
		if lookahead1 == int32('r') {
			state = uint16(612)
			goto next_state
		}
		return result
	case int32(398):
		if lookahead1 == int32('i') {
			state = uint16(364)
			goto next_state
		}
		if lookahead1 == int32('u') {
			state = uint16(233)
			goto next_state
		}
		return result
	case int32(399):
		if lookahead1 == int32('i') {
			state = uint16(228)
			goto next_state
		}
		return result
	case int32(400):
		if lookahead1 == int32('i') {
			state = uint16(912)
			goto next_state
		}
		return result
	case int32(401):
		if lookahead1 == int32('i') {
			state = uint16(699)
			goto next_state
		}
		return result
	case int32(402):
		if lookahead1 == int32('i') {
			state = uint16(543)
			goto next_state
		}
		return result
	case int32(403):
		if lookahead1 == int32('i') {
			state = uint16(888)
			goto next_state
		}
		return result
	case int32(404):
		if lookahead1 == int32('i') {
			state = uint16(158)
			goto next_state
		}
		return result
	case int32(405):
		if lookahead1 == int32('i') {
			state = uint16(168)
			goto next_state
		}
		return result
	case int32(406):
		if lookahead1 == int32('i') {
			state = uint16(376)
			goto next_state
		}
		return result
	case int32(407):
		if lookahead1 == int32('i') {
			state = uint16(174)
			goto next_state
		}
		return result
	case int32(408):
		if lookahead1 == int32('i') {
			state = uint16(554)
			goto next_state
		}
		if lookahead1 == int32('r') {
			state = uint16(237)
			goto next_state
		}
		return result
	case int32(409):
		if lookahead1 == int32('i') {
			state = uint16(776)
			goto next_state
		}
		return result
	case int32(410):
		if lookahead1 == int32('i') {
			state = uint16(159)
			goto next_state
		}
		return result
	case int32(411):
		if lookahead1 == int32('i') {
			state = uint16(815)
			goto next_state
		}
		return result
	case int32(412):
		if lookahead1 == int32('i') {
			state = uint16(596)
			goto next_state
		}
		return result
	case int32(413):
		if lookahead1 == int32('i') {
			state = uint16(818)
			goto next_state
		}
		return result
	case int32(414):
		if lookahead1 == int32('i') {
			state = uint16(564)
			goto next_state
		}
		return result
	case int32(415):
		if lookahead1 == int32('i') {
			state = uint16(846)
			goto next_state
		}
		return result
	case int32(416):
		if lookahead1 == int32('i') {
			state = uint16(820)
			goto next_state
		}
		return result
	case int32(417):
		if lookahead1 == int32('i') {
			state = uint16(809)
			goto next_state
		}
		return result
	case int32(418):
		if lookahead1 == int32('i') {
			state = uint16(810)
			goto next_state
		}
		return result
	case int32(419):
		if lookahead1 == int32('i') {
			state = uint16(602)
			goto next_state
		}
		return result
	case int32(420):
		if lookahead1 == int32('i') {
			state = uint16(290)
			goto next_state
		}
		return result
	case int32(421):
		if lookahead1 == int32('i') {
			state = uint16(822)
			goto next_state
		}
		return result
	case int32(422):
		if lookahead1 == int32('i') {
			state = uint16(208)
			goto next_state
		}
		return result
	case int32(423):
		if lookahead1 == int32('i') {
			state = uint16(366)
			goto next_state
		}
		if lookahead1 == int32('x') {
			state = uint16(93)
			goto next_state
		}
		return result
	case int32(424):
		if lookahead1 == int32('i') {
			state = uint16(913)
			goto next_state
		}
		return result
	case int32(425):
		if lookahead1 == int32('i') {
			state = uint16(889)
			goto next_state
		}
		return result
	case int32(426):
		if lookahead1 == int32('i') {
			state = uint16(357)
			goto next_state
		}
		return result
	case int32(427):
		if lookahead1 == int32('i') {
			state = uint16(825)
			goto next_state
		}
		return result
	case int32(428):
		if lookahead1 == int32('i') {
			state = uint16(798)
			goto next_state
		}
		return result
	case int32(429):
		if lookahead1 == int32('i') {
			state = uint16(213)
			goto next_state
		}
		return result
	case int32(430):
		if lookahead1 == int32('i') {
			state = uint16(894)
			goto next_state
		}
		return result
	case int32(431):
		if lookahead1 == int32('i') {
			state = uint16(359)
			goto next_state
		}
		return result
	case int32(432):
		if lookahead1 == int32('i') {
			state = uint16(522)
			goto next_state
		}
		return result
	case int32(433):
		if lookahead1 == int32('i') {
			state = uint16(652)
			goto next_state
		}
		return result
	case int32(434):
		if lookahead1 == int32('i') {
			state = uint16(635)
			goto next_state
		}
		return result
	case int32(435):
		if lookahead1 == int32('i') {
			state = uint16(323)
			goto next_state
		}
		return result
	case int32(436):
		if lookahead1 == int32('i') {
			state = uint16(485)
			goto next_state
		}
		return result
	case int32(437):
		if lookahead1 == int32('i') {
			state = uint16(526)
			goto next_state
		}
		return result
	case int32(438):
		if lookahead1 == int32('i') {
			state = uint16(303)
			goto next_state
		}
		return result
	case int32(439):
		if lookahead1 == int32('i') {
			state = uint16(634)
			goto next_state
		}
		return result
	case int32(440):
		if lookahead1 == int32('i') {
			state = uint16(657)
			goto next_state
		}
		return result
	case int32(441):
		if lookahead1 == int32('i') {
			state = uint16(661)
			goto next_state
		}
		return result
	case int32(442):
		if lookahead1 == int32('i') {
			state = uint16(642)
			goto next_state
		}
		return result
	case int32(443):
		if lookahead1 == int32('i') {
			state = uint16(311)
			goto next_state
		}
		return result
	case int32(444):
		if lookahead1 == int32('i') {
			state = uint16(312)
			goto next_state
		}
		return result
	case int32(445):
		if lookahead1 == int32('i') {
			state = uint16(155)
			goto next_state
		}
		return result
	case int32(446):
		if lookahead1 == int32('i') {
			state = uint16(175)
			goto next_state
		}
		return result
	case int32(447):
		if lookahead1 == int32('i') {
			state = uint16(176)
			goto next_state
		}
		return result
	case int32(448):
		if lookahead1 == int32('i') {
			state = uint16(342)
			goto next_state
		}
		return result
	case int32(449):
		if lookahead1 == int32('i') {
			state = uint16(851)
			goto next_state
		}
		return result
	case int32(450):
		if lookahead1 == int32('i') {
			state = uint16(854)
			goto next_state
		}
		return result
	case int32(451):
		if lookahead1 == int32('i') {
			state = uint16(378)
			goto next_state
		}
		return result
	case int32(452):
		if lookahead1 == int32('i') {
			state = uint16(179)
			goto next_state
		}
		return result
	case int32(453):
		if lookahead1 == int32('k') {
			state = uint16(1056)
			goto next_state
		}
		return result
	case int32(454):
		if lookahead1 == int32('k') {
			state = uint16(66)
			goto next_state
		}
		return result
	case int32(455):
		if lookahead1 == int32('k') {
			state = uint16(373)
			goto next_state
		}
		return result
	case int32(456):
		if lookahead1 == int32('k') {
			state = uint16(272)
			goto next_state
		}
		return result
	case int32(457):
		if lookahead1 == int32('k') {
			state = uint16(298)
			goto next_state
		}
		return result
	case int32(458):
		if lookahead1 == int32('k') {
			state = uint16(309)
			goto next_state
		}
		return result
	case int32(459):
		if lookahead1 == int32('l') {
			state = uint16(835)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(614)
			goto next_state
		}
		if lookahead1 == int32('r') {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(460):
		if lookahead1 == int32('l') {
			state = uint16(219)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(86)
			goto next_state
		}
		return result
	case int32(461):
		if lookahead1 == int32('l') {
			state = uint16(1009)
			goto next_state
		}
		return result
	case int32(462):
		if lookahead1 == int32('l') {
			state = uint16(921)
			goto next_state
		}
		return result
	case int32(463):
		if lookahead1 == int32('l') {
			state = uint16(921)
			goto next_state
		}
		if lookahead1 == int32('p') {
			state = uint16(39)
			goto next_state
		}
		return result
	case int32(464):
		if lookahead1 == int32('l') {
			state = uint16(944)
			goto next_state
		}
		return result
	case int32(465):
		if lookahead1 == int32('l') {
			state = uint16(923)
			goto next_state
		}
		return result
	case int32(466):
		if lookahead1 == int32('l') {
			state = uint16(927)
			goto next_state
		}
		return result
	case int32(467):
		if lookahead1 == int32('l') {
			state = uint16(938)
			goto next_state
		}
		return result
	case int32(468):
		if lookahead1 == int32('l') {
			state = uint16(955)
			goto next_state
		}
		return result
	case int32(469):
		if lookahead1 == int32('l') {
			state = uint16(1007)
			goto next_state
		}
		return result
	case int32(470):
		if lookahead1 == int32('l') {
			state = uint16(981)
			goto next_state
		}
		return result
	case int32(471):
		if lookahead1 == int32('l') {
			state = uint16(978)
			goto next_state
		}
		return result
	case int32(472):
		if lookahead1 == int32('l') {
			state = uint16(1008)
			goto next_state
		}
		return result
	case int32(473):
		if lookahead1 == int32('l') {
			state = uint16(1003)
			goto next_state
		}
		return result
	case int32(474):
		if lookahead1 == int32('l') {
			state = uint16(882)
			goto next_state
		}
		return result
	case int32(475):
		if lookahead1 == int32('l') {
			state = uint16(613)
			goto next_state
		}
		return result
	case int32(476):
		if lookahead1 == int32('l') {
			state = uint16(615)
			goto next_state
		}
		return result
	case int32(477):
		if lookahead1 == int32('l') {
			state = uint16(451)
			goto next_state
		}
		return result
	case int32(478):
		if lookahead1 == int32('l') {
			state = uint16(244)
			goto next_state
		}
		return result
	case int32(479):
		if lookahead1 == int32('l') {
			state = uint16(245)
			goto next_state
		}
		return result
	case int32(480):
		if lookahead1 == int32('l') {
			state = uint16(212)
			goto next_state
		}
		return result
	case int32(481):
		if lookahead1 == int32('l') {
			state = uint16(247)
			goto next_state
		}
		return result
	case int32(482):
		if lookahead1 == int32('l') {
			state = uint16(196)
			goto next_state
		}
		return result
	case int32(483):
		if lookahead1 == int32('l') {
			state = uint16(252)
			goto next_state
		}
		return result
	case int32(484):
		if lookahead1 == int32('l') {
			state = uint16(345)
			goto next_state
		}
		return result
	case int32(485):
		if lookahead1 == int32('l') {
			state = uint16(260)
			goto next_state
		}
		return result
	case int32(486):
		if lookahead1 == int32('l') {
			state = uint16(411)
			goto next_state
		}
		return result
	case int32(487):
		if lookahead1 == int32('l') {
			state = uint16(784)
			goto next_state
		}
		return result
	case int32(488):
		if lookahead1 == int32('l') {
			state = uint16(405)
			goto next_state
		}
		return result
	case int32(489):
		if lookahead1 == int32('l') {
			state = uint16(475)
			goto next_state
		}
		return result
	case int32(490):
		if lookahead1 == int32('l') {
			state = uint16(653)
			goto next_state
		}
		return result
	case int32(491):
		if lookahead1 == int32('l') {
			state = uint16(413)
			goto next_state
		}
		return result
	case int32(492):
		if lookahead1 == int32('l') {
			state = uint16(407)
			goto next_state
		}
		return result
	case int32(493):
		if lookahead1 == int32('l') {
			state = uint16(627)
			goto next_state
		}
		return result
	case int32(494):
		if lookahead1 == int32('l') {
			state = uint16(416)
			goto next_state
		}
		return result
	case int32(495):
		if lookahead1 == int32('l') {
			state = uint16(638)
			goto next_state
		}
		return result
	case int32(496):
		if lookahead1 == int32('l') {
			state = uint16(640)
			goto next_state
		}
		return result
	case int32(497):
		if lookahead1 == int32('l') {
			state = uint16(641)
			goto next_state
		}
		return result
	case int32(498):
		if lookahead1 == int32('l') {
			state = uint16(322)
			goto next_state
		}
		return result
	case int32(499):
		if lookahead1 == int32('l') {
			state = uint16(643)
			goto next_state
		}
		return result
	case int32(500):
		if lookahead1 == int32('l') {
			state = uint16(446)
			goto next_state
		}
		return result
	case int32(501):
		if lookahead1 == int32('l') {
			state = uint16(447)
			goto next_state
		}
		return result
	case int32(502):
		if lookahead1 == int32('l') {
			state = uint16(222)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(86)
			goto next_state
		}
		return result
	case int32(503):
		if lookahead1 == int32('l') {
			state = uint16(223)
			goto next_state
		}
		return result
	case int32(504):
		if lookahead1 == int32('l') {
			state = uint16(224)
			goto next_state
		}
		return result
	case int32(505):
		if lookahead1 == int32('l') {
			state = uint16(348)
			goto next_state
		}
		return result
	case int32(506):
		if lookahead1 == int32('l') {
			state = uint16(350)
			goto next_state
		}
		return result
	case int32(507):
		if lookahead1 == int32('l') {
			state = uint16(351)
			goto next_state
		}
		return result
	case int32(508):
		if lookahead1 == int32('l') {
			state = uint16(352)
			goto next_state
		}
		return result
	case int32(509):
		if lookahead1 == int32('l') {
			state = uint16(452)
			goto next_state
		}
		return result
	case int32(510):
		if lookahead1 == int32('m') {
			state = uint16(680)
			goto next_state
		}
		return result
	case int32(511):
		if lookahead1 == int32('m') {
			state = uint16(47)
			goto next_state
		}
		return result
	case int32(512):
		if lookahead1 == int32('m') {
			state = uint16(22)
			goto next_state
		}
		return result
	case int32(513):
		if lookahead1 == int32('m') {
			state = uint16(939)
			goto next_state
		}
		return result
	case int32(514):
		if lookahead1 == int32('m') {
			state = uint16(965)
			goto next_state
		}
		return result
	case int32(515):
		if lookahead1 == int32('m') {
			state = uint16(943)
			goto next_state
		}
		return result
	case int32(516):
		if lookahead1 == int32('m') {
			state = uint16(969)
			goto next_state
		}
		return result
	case int32(517):
		if lookahead1 == int32('m') {
			state = uint16(970)
			goto next_state
		}
		return result
	case int32(518):
		if lookahead1 == int32('m') {
			state = uint16(942)
			goto next_state
		}
		return result
	case int32(519):
		if lookahead1 == int32('m') {
			state = uint16(672)
			goto next_state
		}
		return result
	case int32(520):
		if lookahead1 == int32('m') {
			state = uint16(132)
			goto next_state
		}
		return result
	case int32(521):
		if lookahead1 == int32('m') {
			state = uint16(101)
			goto next_state
		}
		return result
	case int32(522):
		if lookahead1 == int32('m') {
			state = uint16(91)
			goto next_state
		}
		return result
	case int32(523):
		if lookahead1 == int32('m') {
			state = uint16(116)
			goto next_state
		}
		return result
	case int32(524):
		if lookahead1 == int32('m') {
			state = uint16(122)
			goto next_state
		}
		return result
	case int32(525):
		if lookahead1 == int32('m') {
			state = uint16(123)
			goto next_state
		}
		return result
	case int32(526):
		if lookahead1 == int32('m') {
			state = uint16(124)
			goto next_state
		}
		return result
	case int32(527):
		if lookahead1 == int32('m') {
			state = uint16(651)
			goto next_state
		}
		return result
	case int32(528):
		if lookahead1 == int32('m') {
			state = uint16(623)
			goto next_state
		}
		return result
	case int32(529):
		if lookahead1 == int32('m') {
			state = uint16(630)
			goto next_state
		}
		return result
	case int32(530):
		if lookahead1 == int32('m') {
			state = uint16(637)
			goto next_state
		}
		return result
	case int32(531):
		if lookahead1 == int32('m') {
			state = uint16(308)
			goto next_state
		}
		return result
	case int32(532):
		if lookahead1 == int32('m') {
			state = uint16(316)
			goto next_state
		}
		return result
	case int32(533):
		if lookahead1 == int32('m') {
			state = uint16(134)
			goto next_state
		}
		return result
	case int32(534):
		if lookahead1 == int32('m') {
			state = uint16(326)
			goto next_state
		}
		return result
	case int32(535):
		if lookahead1 == int32('m') {
			state = uint16(329)
			goto next_state
		}
		return result
	case int32(536):
		if lookahead1 == int32('m') {
			state = uint16(330)
			goto next_state
		}
		return result
	case int32(537):
		if lookahead1 == int32('m') {
			state = uint16(331)
			goto next_state
		}
		return result
	case int32(538):
		if lookahead1 == int32('m') {
			state = uint16(693)
			goto next_state
		}
		return result
	case int32(539):
		if lookahead1 == int32('m') {
			state = uint16(80)
			goto next_state
		}
		return result
	case int32(540):
		if lookahead1 == int32('m') {
			state = uint16(50)
			goto next_state
		}
		return result
	case int32(541):
		if lookahead1 == int32('m') {
			state = uint16(694)
			goto next_state
		}
		return result
	case int32(542):
		if lookahead1 == int32('m') {
			state = uint16(698)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(355)
			goto next_state
		}
		return result
	case int32(543):
		if lookahead1 == int32('n') {
			state = uint16(456)
			goto next_state
		}
		return result
	case int32(544):
		if lookahead1 == int32('n') {
			state = uint16(1053)
			goto next_state
		}
		return result
	case int32(545):
		if lookahead1 == int32('n') {
			state = uint16(1065)
			goto next_state
		}
		return result
	case int32(546):
		if lookahead1 == int32('n') {
			state = uint16(1050)
			goto next_state
		}
		return result
	case int32(547):
		if lookahead1 == int32('n') {
			state = uint16(1048)
			goto next_state
		}
		return result
	case int32(548):
		if lookahead1 == int32('n') {
			state = uint16(1060)
			goto next_state
		}
		return result
	case int32(549):
		if lookahead1 == int32('n') {
			state = uint16(1068)
			goto next_state
		}
		return result
	case int32(550):
		if lookahead1 == int32('n') {
			state = uint16(38)
			goto next_state
		}
		return result
	case int32(551):
		if lookahead1 == int32('n') {
			state = uint16(931)
			goto next_state
		}
		return result
	case int32(552):
		if lookahead1 == int32('n') {
			state = uint16(986)
			goto next_state
		}
		return result
	case int32(553):
		if lookahead1 == int32('n') {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(554):
		if lookahead1 == int32('n') {
			state = uint16(217)
			goto next_state
		}
		return result
	case int32(555):
		if lookahead1 == int32('n') {
			state = uint16(367)
			goto next_state
		}
		return result
	case int32(556):
		if lookahead1 == int32('n') {
			state = uint16(831)
			goto next_state
		}
		if lookahead1 == int32('r') {
			state = uint16(721)
			goto next_state
		}
		return result
	case int32(557):
		if lookahead1 == int32('n') {
			state = uint16(432)
			goto next_state
		}
		if lookahead1 == int32('r') {
			state = uint16(756)
			goto next_state
		}
		return result
	case int32(558):
		if lookahead1 == int32('n') {
			state = uint16(183)
			goto next_state
		}
		return result
	case int32(559):
		if lookahead1 == int32('n') {
			state = uint16(180)
			goto next_state
		}
		return result
	case int32(560):
		if lookahead1 == int32('n') {
			state = uint16(372)
			goto next_state
		}
		return result
	case int32(561):
		if lookahead1 == int32('n') {
			state = uint16(848)
			goto next_state
		}
		return result
	case int32(562):
		if lookahead1 == int32('n') {
			state = uint16(814)
			goto next_state
		}
		return result
	case int32(563):
		if lookahead1 == int32('n') {
			state = uint16(777)
			goto next_state
		}
		return result
	case int32(564):
		if lookahead1 == int32('n') {
			state = uint16(363)
			goto next_state
		}
		return result
	case int32(565):
		if lookahead1 == int32('n') {
			state = uint16(236)
			goto next_state
		}
		return result
	case int32(566):
		if lookahead1 == int32('n') {
			state = uint16(236)
			goto next_state
		}
		if lookahead1 == int32('r') {
			state = uint16(521)
			goto next_state
		}
		return result
	case int32(567):
		if lookahead1 == int32('n') {
			state = uint16(236)
			goto next_state
		}
		if lookahead1 == int32('r') {
			state = uint16(524)
			goto next_state
		}
		return result
	case int32(568):
		if lookahead1 == int32('n') {
			state = uint16(833)
			goto next_state
		}
		return result
	case int32(569):
		if lookahead1 == int32('n') {
			state = uint16(170)
			goto next_state
		}
		return result
	case int32(570):
		if lookahead1 == int32('n') {
			state = uint16(90)
			goto next_state
		}
		return result
	case int32(571):
		if lookahead1 == int32('n') {
			state = uint16(239)
			goto next_state
		}
		return result
	case int32(572):
		if lookahead1 == int32('n') {
			state = uint16(211)
			goto next_state
		}
		return result
	case int32(573):
		if lookahead1 == int32('n') {
			state = uint16(137)
			goto next_state
		}
		return result
	case int32(574):
		if lookahead1 == int32('n') {
			state = uint16(823)
			goto next_state
		}
		return result
	case int32(575):
		if lookahead1 == int32('n') {
			state = uint16(229)
			goto next_state
		}
		return result
	case int32(576):
		if lookahead1 == int32('n') {
			state = uint16(109)
			goto next_state
		}
		return result
	case int32(577):
		if lookahead1 == int32('n') {
			state = uint16(844)
			goto next_state
		}
		return result
	case int32(578):
		if lookahead1 == int32('n') {
			state = uint16(401)
			goto next_state
		}
		return result
	case int32(579):
		if lookahead1 == int32('n') {
			state = uint16(670)
			goto next_state
		}
		return result
	case int32(580):
		if lookahead1 == int32('n') {
			state = uint16(805)
			goto next_state
		}
		return result
	case int32(581):
		if lookahead1 == int32('n') {
			state = uint16(194)
			goto next_state
		}
		return result
	case int32(582):
		if lookahead1 == int32('n') {
			state = uint16(195)
			goto next_state
		}
		return result
	case int32(583):
		if lookahead1 == int32('n') {
			state = uint16(118)
			goto next_state
		}
		return result
	case int32(584):
		if lookahead1 == int32('n') {
			state = uint16(812)
			goto next_state
		}
		return result
	case int32(585):
		if lookahead1 == int32('n') {
			state = uint16(197)
			goto next_state
		}
		return result
	case int32(586):
		if lookahead1 == int32('n') {
			state = uint16(120)
			goto next_state
		}
		return result
	case int32(587):
		if lookahead1 == int32('n') {
			state = uint16(198)
			goto next_state
		}
		return result
	case int32(588):
		if lookahead1 == int32('n') {
			state = uint16(199)
			goto next_state
		}
		return result
	case int32(589):
		if lookahead1 == int32('n') {
			state = uint16(807)
			goto next_state
		}
		return result
	case int32(590):
		if lookahead1 == int32('n') {
			state = uint16(200)
			goto next_state
		}
		return result
	case int32(591):
		if lookahead1 == int32('n') {
			state = uint16(808)
			goto next_state
		}
		return result
	case int32(592):
		if lookahead1 == int32('n') {
			state = uint16(811)
			goto next_state
		}
		return result
	case int32(593):
		if lookahead1 == int32('n') {
			state = uint16(344)
			goto next_state
		}
		return result
	case int32(594):
		if lookahead1 == int32('n') {
			state = uint16(264)
			goto next_state
		}
		return result
	case int32(595):
		if lookahead1 == int32('n') {
			state = uint16(821)
			goto next_state
		}
		return result
	case int32(596):
		if lookahead1 == int32('n') {
			state = uint16(877)
			goto next_state
		}
		return result
	case int32(597):
		if lookahead1 == int32('n') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(598):
		if lookahead1 == int32('n') {
			state = uint16(826)
			goto next_state
		}
		return result
	case int32(599):
		if lookahead1 == int32('n') {
			state = uint16(283)
			goto next_state
		}
		return result
	case int32(600):
		if lookahead1 == int32('n') {
			state = uint16(304)
			goto next_state
		}
		if lookahead1 == int32('v') {
			state = uint16(131)
			goto next_state
		}
		return result
	case int32(601):
		if lookahead1 == int32('n') {
			state = uint16(839)
			goto next_state
		}
		return result
	case int32(602):
		if lookahead1 == int32('n') {
			state = uint16(437)
			goto next_state
		}
		return result
	case int32(603):
		if lookahead1 == int32('n') {
			state = uint16(847)
			goto next_state
		}
		return result
	case int32(604):
		if lookahead1 == int32('n') {
			state = uint16(843)
			goto next_state
		}
		return result
	case int32(605):
		if lookahead1 == int32('n') {
			state = uint16(850)
			goto next_state
		}
		return result
	case int32(606):
		if lookahead1 == int32('n') {
			state = uint16(857)
			goto next_state
		}
		return result
	case int32(607):
		if lookahead1 == int32('n') {
			state = uint16(858)
			goto next_state
		}
		return result
	case int32(608):
		if lookahead1 == int32('o') {
			state = uint16(897)
			goto next_state
		}
		return result
	case int32(609):
		if lookahead1 == int32('o') {
			state = uint16(866)
			goto next_state
		}
		return result
	case int32(610):
		if lookahead1 == int32('o') {
			state = uint16(895)
			goto next_state
		}
		return result
	case int32(611):
		if lookahead1 == int32('o') {
			state = uint16(566)
			goto next_state
		}
		return result
	case int32(612):
		if lookahead1 == int32('o') {
			state = uint16(696)
			goto next_state
		}
		return result
	case int32(613):
		if lookahead1 == int32('o') {
			state = uint16(896)
			goto next_state
		}
		return result
	case int32(614):
		if lookahead1 == int32('o') {
			state = uint16(520)
			goto next_state
		}
		if lookahead1 == int32('y') {
			state = uint16(35)
			goto next_state
		}
		return result
	case int32(615):
		if lookahead1 == int32('o') {
			state = uint16(865)
			goto next_state
		}
		return result
	case int32(616):
		if lookahead1 == int32('o') {
			state = uint16(571)
			goto next_state
		}
		return result
	case int32(617):
		if lookahead1 == int32('o') {
			state = uint16(512)
			goto next_state
		}
		return result
	case int32(618):
		if lookahead1 == int32('o') {
			state = uint16(870)
			goto next_state
		}
		return result
	case int32(619):
		if lookahead1 == int32('o') {
			state = uint16(881)
			goto next_state
		}
		return result
	case int32(620):
		if lookahead1 == int32('o') {
			state = uint16(545)
			goto next_state
		}
		return result
	case int32(621):
		if lookahead1 == int32('o') {
			state = uint16(674)
			goto next_state
		}
		return result
	case int32(622):
		if lookahead1 == int32('o') {
			state = uint16(739)
			goto next_state
		}
		return result
	case int32(623):
		if lookahead1 == int32('o') {
			state = uint16(723)
			goto next_state
		}
		return result
	case int32(624):
		if lookahead1 == int32('o') {
			state = uint16(691)
			goto next_state
		}
		return result
	case int32(625):
		if lookahead1 == int32('o') {
			state = uint16(210)
			goto next_state
		}
		return result
	case int32(626):
		if lookahead1 == int32('o') {
			state = uint16(794)
			goto next_state
		}
		return result
	case int32(627):
		if lookahead1 == int32('o') {
			state = uint16(729)
			goto next_state
		}
		return result
	case int32(628):
		if lookahead1 == int32('o') {
			state = uint16(191)
			goto next_state
		}
		return result
	case int32(629):
		if lookahead1 == int32('o') {
			state = uint16(192)
			goto next_state
		}
		return result
	case int32(630):
		if lookahead1 == int32('o') {
			state = uint16(558)
			goto next_state
		}
		return result
	case int32(631):
		if lookahead1 == int32('o') {
			state = uint16(193)
			goto next_state
		}
		return result
	case int32(632):
		if lookahead1 == int32('o') {
			state = uint16(548)
			goto next_state
		}
		return result
	case int32(633):
		if lookahead1 == int32('o') {
			state = uint16(549)
			goto next_state
		}
		return result
	case int32(634):
		if lookahead1 == int32('o') {
			state = uint16(551)
			goto next_state
		}
		return result
	case int32(635):
		if lookahead1 == int32('o') {
			state = uint16(707)
			goto next_state
		}
		return result
	case int32(636):
		if lookahead1 == int32('o') {
			state = uint16(552)
			goto next_state
		}
		return result
	case int32(637):
		if lookahead1 == int32('o') {
			state = uint16(731)
			goto next_state
		}
		return result
	case int32(638):
		if lookahead1 == int32('o') {
			state = uint16(709)
			goto next_state
		}
		return result
	case int32(639):
		if lookahead1 == int32('o') {
			state = uint16(741)
			goto next_state
		}
		return result
	case int32(640):
		if lookahead1 == int32('o') {
			state = uint16(711)
			goto next_state
		}
		return result
	case int32(641):
		if lookahead1 == int32('o') {
			state = uint16(712)
			goto next_state
		}
		return result
	case int32(642):
		if lookahead1 == int32('o') {
			state = uint16(718)
			goto next_state
		}
		return result
	case int32(643):
		if lookahead1 == int32('o') {
			state = uint16(744)
			goto next_state
		}
		return result
	case int32(644):
		if lookahead1 == int32('o') {
			state = uint16(898)
			goto next_state
		}
		return result
	case int32(645):
		if lookahead1 == int32('o') {
			state = uint16(867)
			goto next_state
		}
		return result
	case int32(646):
		if lookahead1 == int32('o') {
			state = uint16(783)
			goto next_state
		}
		return result
	case int32(647):
		if lookahead1 == int32('o') {
			state = uint16(783)
			goto next_state
		}
		if lookahead1 == int32('u') {
			state = uint16(786)
			goto next_state
		}
		return result
	case int32(648):
		if lookahead1 == int32('o') {
			state = uint16(593)
			goto next_state
		}
		return result
	case int32(649):
		if lookahead1 == int32('o') {
			state = uint16(871)
			goto next_state
		}
		return result
	case int32(650):
		if lookahead1 == int32('o') {
			state = uint16(745)
			goto next_state
		}
		return result
	case int32(651):
		if lookahead1 == int32('o') {
			state = uint16(872)
			goto next_state
		}
		return result
	case int32(652):
		if lookahead1 == int32('o') {
			state = uint16(576)
			goto next_state
		}
		return result
	case int32(653):
		if lookahead1 == int32('o') {
			state = uint16(868)
			goto next_state
		}
		return result
	case int32(654):
		if lookahead1 == int32('o') {
			state = uint16(565)
			goto next_state
		}
		return result
	case int32(655):
		if lookahead1 == int32('o') {
			state = uint16(873)
			goto next_state
		}
		return result
	case int32(656):
		if lookahead1 == int32('o') {
			state = uint16(214)
			goto next_state
		}
		if lookahead1 == int32('u') {
			state = uint16(518)
			goto next_state
		}
		return result
	case int32(657):
		if lookahead1 == int32('o') {
			state = uint16(583)
			goto next_state
		}
		return result
	case int32(658):
		if lookahead1 == int32('o') {
			state = uint16(594)
			goto next_state
		}
		return result
	case int32(659):
		if lookahead1 == int32('o') {
			state = uint16(874)
			goto next_state
		}
		return result
	case int32(660):
		if lookahead1 == int32('o') {
			state = uint16(577)
			goto next_state
		}
		return result
	case int32(661):
		if lookahead1 == int32('o') {
			state = uint16(586)
			goto next_state
		}
		return result
	case int32(662):
		if lookahead1 == int32('o') {
			state = uint16(875)
			goto next_state
		}
		return result
	case int32(663):
		if lookahead1 == int32('o') {
			state = uint16(876)
			goto next_state
		}
		return result
	case int32(664):
		if lookahead1 == int32('o') {
			state = uint16(878)
			goto next_state
		}
		return result
	case int32(665):
		if lookahead1 == int32('o') {
			state = uint16(173)
			goto next_state
		}
		return result
	case int32(666):
		if lookahead1 == int32('o') {
			state = uint16(567)
			goto next_state
		}
		return result
	case int32(667):
		if lookahead1 == int32('o') {
			state = uint16(495)
			goto next_state
		}
		return result
	case int32(668):
		if lookahead1 == int32('o') {
			state = uint16(496)
			goto next_state
		}
		return result
	case int32(669):
		if lookahead1 == int32('o') {
			state = uint16(497)
			goto next_state
		}
		return result
	case int32(670):
		if lookahead1 == int32('o') {
			state = uint16(533)
			goto next_state
		}
		return result
	case int32(671):
		if lookahead1 == int32('o') {
			state = uint16(695)
			goto next_state
		}
		return result
	case int32(672):
		if lookahead1 == int32('p') {
			state = uint16(989)
			goto next_state
		}
		return result
	case int32(673):
		if lookahead1 == int32('p') {
			state = uint16(940)
			goto next_state
		}
		return result
	case int32(674):
		if lookahead1 == int32('p') {
			state = uint16(1066)
			goto next_state
		}
		return result
	case int32(675):
		if lookahead1 == int32('p') {
			state = uint16(964)
			goto next_state
		}
		return result
	case int32(676):
		if lookahead1 == int32('p') {
			state = uint16(934)
			goto next_state
		}
		return result
	case int32(677):
		if lookahead1 == int32('p') {
			state = uint16(966)
			goto next_state
		}
		return result
	case int32(678):
		if lookahead1 == int32('p') {
			state = uint16(971)
			goto next_state
		}
		return result
	case int32(679):
		if lookahead1 == int32('p') {
			state = uint16(816)
			goto next_state
		}
		if lookahead1 == int32('r') {
			state = uint16(106)
			goto next_state
		}
		return result
	case int32(680):
		if lookahead1 == int32('p') {
			state = uint16(622)
			goto next_state
		}
		return result
	case int32(681):
		if lookahead1 == int32('p') {
			state = uint16(242)
			goto next_state
		}
		return result
	case int32(682):
		if lookahead1 == int32('p') {
			state = uint16(250)
			goto next_state
		}
		return result
	case int32(683):
		if lookahead1 == int32('p') {
			state = uint16(258)
			goto next_state
		}
		return result
	case int32(684):
		if lookahead1 == int32('p') {
			state = uint16(267)
			goto next_state
		}
		return result
	case int32(685):
		if lookahead1 == int32('p') {
			state = uint16(481)
			goto next_state
		}
		return result
	case int32(686):
		if lookahead1 == int32('p') {
			state = uint16(829)
			goto next_state
		}
		return result
	case int32(687):
		if lookahead1 == int32('p') {
			state = uint16(293)
			goto next_state
		}
		return result
	case int32(688):
		if lookahead1 == int32('p') {
			state = uint16(315)
			goto next_state
		}
		return result
	case int32(689):
		if lookahead1 == int32('p') {
			state = uint16(316)
			goto next_state
		}
		return result
	case int32(690):
		if lookahead1 == int32('p') {
			state = uint16(492)
			goto next_state
		}
		return result
	case int32(691):
		if lookahead1 == int32('p') {
			state = uint16(71)
			goto next_state
		}
		return result
	case int32(692):
		if lookahead1 == int32('p') {
			state = uint16(339)
			goto next_state
		}
		return result
	case int32(693):
		if lookahead1 == int32('p') {
			state = uint16(500)
			goto next_state
		}
		return result
	case int32(694):
		if lookahead1 == int32('p') {
			state = uint16(501)
			goto next_state
		}
		return result
	case int32(695):
		if lookahead1 == int32('p') {
			state = uint16(73)
			goto next_state
		}
		return result
	case int32(696):
		if lookahead1 == int32('p') {
			state = uint16(49)
			goto next_state
		}
		return result
	case int32(697):
		if lookahead1 == int32('p') {
			state = uint16(852)
			goto next_state
		}
		return result
	case int32(698):
		if lookahead1 == int32('p') {
			state = uint16(509)
			goto next_state
		}
		return result
	case int32(699):
		if lookahead1 == int32('q') {
			state = uint16(879)
			goto next_state
		}
		return result
	case int32(700):
		if lookahead1 == int32('r') {
			state = uint16(26)
			goto next_state
		}
		return result
	case int32(701):
		if lookahead1 == int32('r') {
			state = uint16(1017)
			goto next_state
		}
		return result
	case int32(702):
		if lookahead1 == int32('r') {
			state = uint16(1019)
			goto next_state
		}
		return result
	case int32(703):
		if lookahead1 == int32('r') {
			state = uint16(957)
			goto next_state
		}
		return result
	case int32(704):
		if lookahead1 == int32('r') {
			state = uint16(1018)
			goto next_state
		}
		return result
	case int32(705):
		if lookahead1 == int32('r') {
			state = uint16(60)
			goto next_state
		}
		return result
	case int32(706):
		if lookahead1 == int32('r') {
			state = uint16(1020)
			goto next_state
		}
		return result
	case int32(707):
		if lookahead1 == int32('r') {
			state = uint16(1004)
			goto next_state
		}
		return result
	case int32(708):
		if lookahead1 == int32('r') {
			state = uint16(928)
			goto next_state
		}
		return result
	case int32(709):
		if lookahead1 == int32('r') {
			state = uint16(993)
			goto next_state
		}
		return result
	case int32(710):
		if lookahead1 == int32('r') {
			state = uint16(973)
			goto next_state
		}
		return result
	case int32(711):
		if lookahead1 == int32('r') {
			state = uint16(991)
			goto next_state
		}
		return result
	case int32(712):
		if lookahead1 == int32('r') {
			state = uint16(990)
			goto next_state
		}
		return result
	case int32(713):
		if lookahead1 == int32('r') {
			state = uint16(972)
			goto next_state
		}
		return result
	case int32(714):
		if lookahead1 == int32('r') {
			state = uint16(110)
			goto next_state
		}
		return result
	case int32(715):
		if lookahead1 == int32('r') {
			state = uint16(41)
			goto next_state
		}
		return result
	case int32(716):
		if lookahead1 == int32('r') {
			state = uint16(1016)
			goto next_state
		}
		return result
	case int32(717):
		if lookahead1 == int32('r') {
			state = uint16(1014)
			goto next_state
		}
		return result
	case int32(718):
		if lookahead1 == int32('r') {
			state = uint16(1003)
			goto next_state
		}
		return result
	case int32(719):
		if lookahead1 == int32('r') {
			state = uint16(1013)
			goto next_state
		}
		return result
	case int32(720):
		if lookahead1 == int32('r') {
			state = uint16(164)
			goto next_state
		}
		return result
	case int32(721):
		if lookahead1 == int32('r') {
			state = uint16(862)
			goto next_state
		}
		return result
	case int32(722):
		if lookahead1 == int32('r') {
			state = uint16(97)
			goto next_state
		}
		return result
	case int32(723):
		if lookahead1 == int32('r') {
			state = uint16(907)
			goto next_state
		}
		return result
	case int32(724):
		if lookahead1 == int32('r') {
			state = uint16(685)
			goto next_state
		}
		return result
	case int32(725):
		if lookahead1 == int32('r') {
			state = uint16(369)
			goto next_state
		}
		return result
	case int32(726):
		if lookahead1 == int32('r') {
			state = uint16(863)
			goto next_state
		}
		return result
	case int32(727):
		if lookahead1 == int32('r') {
			state = uint16(493)
			goto next_state
		}
		return result
	case int32(728):
		if lookahead1 == int32('r') {
			state = uint16(493)
			goto next_state
		}
		if lookahead1 == int32('y') {
			state = uint16(785)
			goto next_state
		}
		return result
	case int32(729):
		if lookahead1 == int32('r') {
			state = uint16(184)
			goto next_state
		}
		return result
	case int32(730):
		if lookahead1 == int32('r') {
			state = uint16(185)
			goto next_state
		}
		return result
	case int32(731):
		if lookahead1 == int32('r') {
			state = uint16(904)
			goto next_state
		}
		return result
	case int32(732):
		if lookahead1 == int32('r') {
			state = uint16(370)
			goto next_state
		}
		return result
	case int32(733):
		if lookahead1 == int32('r') {
			state = uint16(434)
			goto next_state
		}
		return result
	case int32(734):
		if lookahead1 == int32('r') {
			state = uint16(621)
			goto next_state
		}
		return result
	case int32(735):
		if lookahead1 == int32('r') {
			state = uint16(178)
			goto next_state
		}
		return result
	case int32(736):
		if lookahead1 == int32('r') {
			state = uint16(624)
			goto next_state
		}
		return result
	case int32(737):
		if lookahead1 == int32('r') {
			state = uint16(165)
			goto next_state
		}
		return result
	case int32(738):
		if lookahead1 == int32('r') {
			state = uint16(645)
			goto next_state
		}
		return result
	case int32(739):
		if lookahead1 == int32('r') {
			state = uint16(802)
			goto next_state
		}
		return result
	case int32(740):
		if lookahead1 == int32('r') {
			state = uint16(618)
			goto next_state
		}
		return result
	case int32(741):
		if lookahead1 == int32('r') {
			state = uint16(480)
			goto next_state
		}
		return result
	case int32(742):
		if lookahead1 == int32('r') {
			state = uint16(834)
			goto next_state
		}
		return result
	case int32(743):
		if lookahead1 == int32('r') {
			state = uint16(248)
			goto next_state
		}
		return result
	case int32(744):
		if lookahead1 == int32('r') {
			state = uint16(202)
			goto next_state
		}
		return result
	case int32(745):
		if lookahead1 == int32('r') {
			state = uint16(282)
			goto next_state
		}
		return result
	case int32(746):
		if lookahead1 == int32('r') {
			state = uint16(288)
			goto next_state
		}
		return result
	case int32(747):
		if lookahead1 == int32('r') {
			state = uint16(300)
			goto next_state
		}
		return result
	case int32(748):
		if lookahead1 == int32('r') {
			state = uint16(317)
			goto next_state
		}
		return result
	case int32(749):
		if lookahead1 == int32('r') {
			state = uint16(263)
			goto next_state
		}
		return result
	case int32(750):
		if lookahead1 == int32('r') {
			state = uint16(527)
			goto next_state
		}
		return result
	case int32(751):
		if lookahead1 == int32('r') {
			state = uint16(368)
			goto next_state
		}
		return result
	case int32(752):
		if lookahead1 == int32('r') {
			state = uint16(573)
			goto next_state
		}
		return result
	case int32(753):
		if lookahead1 == int32('r') {
			state = uint16(861)
			goto next_state
		}
		return result
	case int32(754):
		if lookahead1 == int32('r') {
			state = uint16(106)
			goto next_state
		}
		return result
	case int32(755):
		if lookahead1 == int32('r') {
			state = uint16(414)
			goto next_state
		}
		return result
	case int32(756):
		if lookahead1 == int32('r') {
			state = uint16(650)
			goto next_state
		}
		return result
	case int32(757):
		if lookahead1 == int32('r') {
			state = uint16(125)
			goto next_state
		}
		return result
	case int32(758):
		if lookahead1 == int32('r') {
			state = uint16(442)
			goto next_state
		}
		return result
	case int32(759):
		if lookahead1 == int32('r') {
			state = uint16(221)
			goto next_state
		}
		return result
	case int32(760):
		if lookahead1 == int32('r') {
			state = uint16(499)
			goto next_state
		}
		return result
	case int32(761):
		if lookahead1 == int32('r') {
			state = uint16(377)
			goto next_state
		}
		return result
	case int32(762):
		if lookahead1 == int32('r') {
			state = uint16(671)
			goto next_state
		}
		return result
	case int32(763):
		if lookahead1 == int32('r') {
			state = uint16(853)
			goto next_state
		}
		return result
	case int32(764):
		if lookahead1 == int32('r') {
			state = uint16(855)
			goto next_state
		}
		return result
	case int32(765):
		if lookahead1 == int32('r') {
			state = uint16(856)
			goto next_state
		}
		return result
	case int32(766):
		if lookahead1 == int32('r') {
			state = uint16(29)
			goto next_state
		}
		return result
	case int32(767):
		if lookahead1 == int32('r') {
			state = uint16(887)
			goto next_state
		}
		return result
	case int32(768):
		if lookahead1 == int32('s') {
			state = uint16(27)
			goto next_state
		}
		return result
	case int32(769):
		if lookahead1 == int32('s') {
			state = uint16(935)
			goto next_state
		}
		return result
	case int32(770):
		if lookahead1 == int32('s') {
			state = uint16(1064)
			goto next_state
		}
		return result
	case int32(771):
		if lookahead1 == int32('s') {
			state = uint16(1006)
			goto next_state
		}
		return result
	case int32(772):
		if lookahead1 == int32('s') {
			state = uint16(937)
			goto next_state
		}
		return result
	case int32(773):
		if lookahead1 == int32('s') {
			state = uint16(956)
			goto next_state
		}
		return result
	case int32(774):
		if lookahead1 == int32('s') {
			state = uint16(958)
			goto next_state
		}
		return result
	case int32(775):
		if lookahead1 == int32('s') {
			state = uint16(1003)
			goto next_state
		}
		return result
	case int32(776):
		if lookahead1 == int32('s') {
			state = uint16(51)
			goto next_state
		}
		return result
	case int32(777):
		if lookahead1 == int32('s') {
			state = uint16(356)
			goto next_state
		}
		return result
	case int32(778):
		if lookahead1 == int32('s') {
			state = uint16(40)
			goto next_state
		}
		return result
	case int32(779):
		if lookahead1 == int32('s') {
			state = uint16(819)
			goto next_state
		}
		return result
	case int32(780):
		if lookahead1 == int32('s') {
			state = uint16(232)
			goto next_state
		}
		return result
	case int32(781):
		if lookahead1 == int32('s') {
			state = uint16(429)
			goto next_state
		}
		return result
	case int32(782):
		if lookahead1 == int32('s') {
			state = uint16(769)
			goto next_state
		}
		return result
	case int32(783):
		if lookahead1 == int32('s') {
			state = uint16(770)
			goto next_state
		}
		return result
	case int32(784):
		if lookahead1 == int32('s') {
			state = uint16(238)
			goto next_state
		}
		return result
	case int32(785):
		if lookahead1 == int32('s') {
			state = uint16(845)
			goto next_state
		}
		return result
	case int32(786):
		if lookahead1 == int32('s') {
			state = uint16(138)
			goto next_state
		}
		return result
	case int32(787):
		if lookahead1 == int32('s') {
			state = uint16(428)
			goto next_state
		}
		return result
	case int32(788):
		if lookahead1 == int32('s') {
			state = uint16(409)
			goto next_state
		}
		return result
	case int32(789):
		if lookahead1 == int32('s') {
			state = uint16(403)
			goto next_state
		}
		return result
	case int32(790):
		if lookahead1 == int32('s') {
			state = uint16(259)
			goto next_state
		}
		return result
	case int32(791):
		if lookahead1 == int32('s') {
			state = uint16(425)
			goto next_state
		}
		return result
	case int32(792):
		if lookahead1 == int32('s') {
			state = uint16(523)
			goto next_state
		}
		return result
	case int32(793):
		if lookahead1 == int32('s') {
			state = uint16(789)
			goto next_state
		}
		return result
	case int32(794):
		if lookahead1 == int32('s') {
			state = uint16(450)
			goto next_state
		}
		return result
	case int32(795):
		if lookahead1 == int32('s') {
			state = uint16(791)
			goto next_state
		}
		return result
	case int32(796):
		if lookahead1 == int32('s') {
			state = uint16(525)
			goto next_state
		}
		return result
	case int32(797):
		if lookahead1 == int32('s') {
			state = uint16(439)
			goto next_state
		}
		return result
	case int32(798):
		if lookahead1 == int32('s') {
			state = uint16(338)
			goto next_state
		}
		return result
	case int32(799):
		if lookahead1 == int32('s') {
			state = uint16(148)
			goto next_state
		}
		return result
	case int32(800):
		if lookahead1 == int32('t') {
			state = uint16(25)
			goto next_state
		}
		return result
	case int32(801):
		if lookahead1 == int32('t') {
			state = uint16(953)
			goto next_state
		}
		return result
	case int32(802):
		if lookahead1 == int32('t') {
			state = uint16(922)
			goto next_state
		}
		return result
	case int32(803):
		if lookahead1 == int32('t') {
			state = uint16(36)
			goto next_state
		}
		return result
	case int32(804):
		if lookahead1 == int32('t') {
			state = uint16(28)
			goto next_state
		}
		return result
	case int32(805):
		if lookahead1 == int32('t') {
			state = uint16(1005)
			goto next_state
		}
		return result
	case int32(806):
		if lookahead1 == int32('t') {
			state = uint16(988)
			goto next_state
		}
		return result
	case int32(807):
		if lookahead1 == int32('t') {
			state = uint16(925)
			goto next_state
		}
		return result
	case int32(808):
		if lookahead1 == int32('t') {
			state = uint16(948)
			goto next_state
		}
		return result
	case int32(809):
		if lookahead1 == int32('t') {
			state = uint16(947)
			goto next_state
		}
		return result
	case int32(810):
		if lookahead1 == int32('t') {
			state = uint16(952)
			goto next_state
		}
		return result
	case int32(811):
		if lookahead1 == int32('t') {
			state = uint16(1003)
			goto next_state
		}
		return result
	case int32(812):
		if lookahead1 == int32('t') {
			state = uint16(62)
			goto next_state
		}
		return result
	case int32(813):
		if lookahead1 == int32('t') {
			state = uint16(902)
			goto next_state
		}
		return result
	case int32(814):
		if lookahead1 == int32('t') {
			state = uint16(384)
			goto next_state
		}
		return result
	case int32(815):
		if lookahead1 == int32('t') {
			state = uint16(903)
			goto next_state
		}
		return result
	case int32(816):
		if lookahead1 == int32('t') {
			state = uint16(433)
			goto next_state
		}
		return result
	case int32(817):
		if lookahead1 == int32('t') {
			state = uint16(379)
			goto next_state
		}
		return result
	case int32(818):
		if lookahead1 == int32('t') {
			state = uint16(906)
			goto next_state
		}
		return result
	case int32(819):
		if lookahead1 == int32('t') {
			state = uint16(617)
			goto next_state
		}
		return result
	case int32(820):
		if lookahead1 == int32('t') {
			state = uint16(905)
			goto next_state
		}
		return result
	case int32(821):
		if lookahead1 == int32('t') {
			state = uint16(426)
			goto next_state
		}
		return result
	case int32(822):
		if lookahead1 == int32('t') {
			state = uint16(235)
			goto next_state
		}
		return result
	case int32(823):
		if lookahead1 == int32('t') {
			state = uint16(535)
			goto next_state
		}
		return result
	case int32(824):
		if lookahead1 == int32('t') {
			state = uint16(773)
			goto next_state
		}
		return result
	case int32(825):
		if lookahead1 == int32('t') {
			state = uint16(243)
			goto next_state
		}
		return result
	case int32(826):
		if lookahead1 == int32('t') {
			state = uint16(436)
			goto next_state
		}
		return result
	case int32(827):
		if lookahead1 == int32('t') {
			state = uint16(757)
			goto next_state
		}
		return result
	case int32(828):
		if lookahead1 == int32('t') {
			state = uint16(279)
			goto next_state
		}
		return result
	case int32(829):
		if lookahead1 == int32('t') {
			state = uint16(285)
			goto next_state
		}
		return result
	case int32(830):
		if lookahead1 == int32('t') {
			state = uint16(253)
			goto next_state
		}
		return result
	case int32(831):
		if lookahead1 == int32('t') {
			state = uint16(412)
			goto next_state
		}
		return result
	case int32(832):
		if lookahead1 == int32('t') {
			state = uint16(880)
			goto next_state
		}
		return result
	case int32(833):
		if lookahead1 == int32('t') {
			state = uint16(105)
			goto next_state
		}
		return result
	case int32(834):
		if lookahead1 == int32('t') {
			state = uint16(70)
			goto next_state
		}
		return result
	case int32(835):
		if lookahead1 == int32('t') {
			state = uint16(275)
			goto next_state
		}
		return result
	case int32(836):
		if lookahead1 == int32('t') {
			state = uint16(287)
			goto next_state
		}
		return result
	case int32(837):
		if lookahead1 == int32('t') {
			state = uint16(289)
			goto next_state
		}
		return result
	case int32(838):
		if lookahead1 == int32('t') {
			state = uint16(292)
			goto next_state
		}
		return result
	case int32(839):
		if lookahead1 == int32('t') {
			state = uint16(299)
			goto next_state
		}
		return result
	case int32(840):
		if lookahead1 == int32('t') {
			state = uint16(295)
			goto next_state
		}
		return result
	case int32(841):
		if lookahead1 == int32('t') {
			state = uint16(301)
			goto next_state
		}
		return result
	case int32(842):
		if lookahead1 == int32('t') {
			state = uint16(307)
			goto next_state
		}
		return result
	case int32(843):
		if lookahead1 == int32('t') {
			state = uint16(316)
			goto next_state
		}
		return result
	case int32(844):
		if lookahead1 == int32('t') {
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(845):
		if lookahead1 == int32('t') {
			state = uint16(648)
			goto next_state
		}
		return result
	case int32(846):
		if lookahead1 == int32('t') {
			state = uint16(56)
			goto next_state
		}
		return result
	case int32(847):
		if lookahead1 == int32('t') {
			state = uint16(431)
			goto next_state
		}
		return result
	case int32(848):
		if lookahead1 == int32('t') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(849):
		if lookahead1 == int32('t') {
			state = uint16(335)
			goto next_state
		}
		return result
	case int32(850):
		if lookahead1 == int32('t') {
			state = uint16(142)
			goto next_state
		}
		return result
	case int32(851):
		if lookahead1 == int32('t') {
			state = uint16(57)
			goto next_state
		}
		return result
	case int32(852):
		if lookahead1 == int32('t') {
			state = uint16(440)
			goto next_state
		}
		return result
	case int32(853):
		if lookahead1 == int32('t') {
			state = uint16(72)
			goto next_state
		}
		return result
	case int32(854):
		if lookahead1 == int32('t') {
			state = uint16(441)
			goto next_state
		}
		return result
	case int32(855):
		if lookahead1 == int32('t') {
			state = uint16(74)
			goto next_state
		}
		return result
	case int32(856):
		if lookahead1 == int32('t') {
			state = uint16(75)
			goto next_state
		}
		return result
	case int32(857):
		if lookahead1 == int32('t') {
			state = uint16(536)
			goto next_state
		}
		return result
	case int32(858):
		if lookahead1 == int32('t') {
			state = uint16(537)
			goto next_state
		}
		return result
	case int32(859):
		if lookahead1 == int32('u') {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(860):
		if lookahead1 == int32('u') {
			state = uint16(96)
			goto next_state
		}
		return result
	case int32(861):
		if lookahead1 == int32('u') {
			state = uint16(786)
			goto next_state
		}
		return result
	case int32(862):
		if lookahead1 == int32('u') {
			state = uint16(686)
			goto next_state
		}
		return result
	case int32(863):
		if lookahead1 == int32('u') {
			state = uint16(167)
			goto next_state
		}
		return result
	case int32(864):
		if lookahead1 == int32('u') {
			state = uint16(113)
			goto next_state
		}
		return result
	case int32(865):
		if lookahead1 == int32('u') {
			state = uint16(771)
			goto next_state
		}
		return result
	case int32(866):
		if lookahead1 == int32('u') {
			state = uint16(751)
			goto next_state
		}
		return result
	case int32(867):
		if lookahead1 == int32('u') {
			state = uint16(677)
			goto next_state
		}
		return result
	case int32(868):
		if lookahead1 == int32('u') {
			state = uint16(775)
			goto next_state
		}
		return result
	case int32(869):
		if lookahead1 == int32('u') {
			state = uint16(249)
			goto next_state
		}
		return result
	case int32(870):
		if lookahead1 == int32('u') {
			state = uint16(575)
			goto next_state
		}
		return result
	case int32(871):
		if lookahead1 == int32('u') {
			state = uint16(581)
			goto next_state
		}
		return result
	case int32(872):
		if lookahead1 == int32('u') {
			state = uint16(708)
			goto next_state
		}
		return result
	case int32(873):
		if lookahead1 == int32('u') {
			state = uint16(582)
			goto next_state
		}
		return result
	case int32(874):
		if lookahead1 == int32('u') {
			state = uint16(585)
			goto next_state
		}
		return result
	case int32(875):
		if lookahead1 == int32('u') {
			state = uint16(587)
			goto next_state
		}
		return result
	case int32(876):
		if lookahead1 == int32('u') {
			state = uint16(588)
			goto next_state
		}
		return result
	case int32(877):
		if lookahead1 == int32('u') {
			state = uint16(251)
			goto next_state
		}
		return result
	case int32(878):
		if lookahead1 == int32('u') {
			state = uint16(590)
			goto next_state
		}
		return result
	case int32(879):
		if lookahead1 == int32('u') {
			state = uint16(263)
			goto next_state
		}
		return result
	case int32(880):
		if lookahead1 == int32('u') {
			state = uint16(746)
			goto next_state
		}
		return result
	case int32(881):
		if lookahead1 == int32('u') {
			state = uint16(790)
			goto next_state
		}
		return result
	case int32(882):
		if lookahead1 == int32('u') {
			state = uint16(324)
			goto next_state
		}
		return result
	case int32(883):
		if lookahead1 == int32('u') {
			state = uint16(747)
			goto next_state
		}
		return result
	case int32(884):
		if lookahead1 == int32('u') {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(885):
		if lookahead1 == int32('u') {
			state = uint16(147)
			goto next_state
		}
		return result
	case int32(886):
		if lookahead1 == int32('u') {
			state = uint16(604)
			goto next_state
		}
		return result
	case int32(887):
		if lookahead1 == int32('u') {
			state = uint16(799)
			goto next_state
		}
		return result
	case int32(888):
		if lookahead1 == int32('v') {
			state = uint16(266)
			goto next_state
		}
		return result
	case int32(889):
		if lookahead1 == int32('v') {
			state = uint16(271)
			goto next_state
		}
		return result
	case int32(890):
		if lookahead1 == int32('v') {
			state = uint16(286)
			goto next_state
		}
		return result
	case int32(891):
		if lookahead1 == int32('v') {
			state = uint16(291)
			goto next_state
		}
		return result
	case int32(892):
		if lookahead1 == int32('v') {
			state = uint16(294)
			goto next_state
		}
		return result
	case int32(893):
		if lookahead1 == int32('v') {
			state = uint16(296)
			goto next_state
		}
		return result
	case int32(894):
		if lookahead1 == int32('v') {
			state = uint16(346)
			goto next_state
		}
		return result
	case int32(895):
		if lookahead1 == int32('w') {
			state = uint16(919)
			goto next_state
		}
		return result
	case int32(896):
		if lookahead1 == int32('w') {
			state = uint16(1052)
			goto next_state
		}
		return result
	case int32(897):
		if lookahead1 == int32('w') {
			state = uint16(546)
			goto next_state
		}
		return result
	case int32(898):
		if lookahead1 == int32('w') {
			state = uint16(550)
			goto next_state
		}
		return result
	case int32(899):
		if lookahead1 == int32('x') {
			state = uint16(804)
			goto next_state
		}
		return result
	case int32(900):
		if lookahead1 == int32('x') {
			state = uint16(129)
			goto next_state
		}
		return result
	case int32(901):
		if lookahead1 == int32('y') {
			state = uint16(20)
			goto next_state
		}
		return result
	case int32(902):
		if lookahead1 == int32('y') {
			state = uint16(961)
			goto next_state
		}
		return result
	case int32(903):
		if lookahead1 == int32('y') {
			state = uint16(960)
			goto next_state
		}
		return result
	case int32(904):
		if lookahead1 == int32('y') {
			state = uint16(975)
			goto next_state
		}
		return result
	case int32(905):
		if lookahead1 == int32('y') {
			state = uint16(924)
			goto next_state
		}
		return result
	case int32(906):
		if lookahead1 == int32('y') {
			state = uint16(78)
			goto next_state
		}
		return result
	case int32(907):
		if lookahead1 == int32('y') {
			state = uint16(68)
			goto next_state
		}
		return result
	case int32(908):
		if lookahead1 == int32('y') {
			state = uint16(65)
			goto next_state
		}
		return result
	case int32(909):
		if lookahead1 == int32('y') {
			state = uint16(785)
			goto next_state
		}
		return result
	case int32(910):
		if lookahead1 == int32('y') {
			state = uint16(682)
			goto next_state
		}
		return result
	case int32(911):
		if lookahead1 == int32('y') {
			state = uint16(683)
			goto next_state
		}
		return result
	case int32(912):
		if lookahead1 == int32('z') {
			state = uint16(254)
			goto next_state
		}
		return result
	case int32(913):
		if lookahead1 == int32('z') {
			state = uint16(256)
			goto next_state
		}
		return result
	case int32(914):
		if eof != 0 {
			state = uint16(918)
			goto next_state
		}
		i6 = uint32(0)
		for {
			if !(uint64(i6) < libc.Uint64FromInt64(148)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token6[i6]) == lookahead1 {
				state = map_token6[i6+uint32(1)]
				goto next_state
			}
			goto _27
		_27:
			;
			i6 = i6 + uint32(2)
		}
		if lookahead1 == int32('\t') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(914)
			goto next_state
		}
		if lookahead1 == int32('5') || lookahead1 == int32('6') {
			state = uint16(1045)
			goto next_state
		}
		if int32('7') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(1085)
			goto next_state
		}
		return result
	case int32(915):
		if eof != 0 {
			state = uint16(918)
			goto next_state
		}
		i7 = uint32(0)
		for {
			if !(uint64(i7) < libc.Uint64FromInt64(128)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token7[i7]) == lookahead1 {
				state = map_token7[i7+uint32(1)]
				goto next_state
			}
			goto _28
		_28:
			;
			i7 = i7 + uint32(2)
		}
		if lookahead1 == int32('\t') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(916)
			goto next_state
		}
		if int32('1') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(1085)
			goto next_state
		}
		return result
	case int32(916):
		if eof != 0 {
			state = uint16(918)
			goto next_state
		}
		i8 = uint32(0)
		for {
			if !(uint64(i8) < libc.Uint64FromInt64(124)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token8[i8]) == lookahead1 {
				state = map_token8[i8+uint32(1)]
				goto next_state
			}
			goto _29
		_29:
			;
			i8 = i8 + uint32(2)
		}
		if lookahead1 == int32('\t') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(916)
			goto next_state
		}
		if int32('1') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(1085)
			goto next_state
		}
		return result
	case int32(917):
		if eof != 0 {
			state = uint16(918)
			goto next_state
		}
		i9 = uint32(0)
		for {
			if !(uint64(i9) < libc.Uint64FromInt64(36)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token9[i9]) == lookahead1 {
				state = map_token9[i9+uint32(1)]
				goto next_state
			}
			goto _30
		_30:
			;
			i9 = i9 + uint32(2)
		}
		return result
	case int32(918):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(ts_builtin_sym_end)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(919):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Show)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(920):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Hide)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(921):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Minimal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(922):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Import)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(923):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Optional)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(924):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AlternateQuality)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(925):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AnyEnchantment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(926):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ArchnemesisMod)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(927):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AreaLevel)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(928):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BaseArmour)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(929):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BaseDefencePercentile)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(930):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BaseEnergyShield)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(931):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BaseEvasion)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(932):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BaseType)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(933):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BaseWard)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(934):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BlightedMap)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(935):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Class)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(936):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Corrupted)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('M') {
			state = uint16(625)
			goto next_state
		}
		return result
	case int32(937):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_CorruptedMods)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(938):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DropLevel)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(939):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ElderItem)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(940):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ElderMap)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(941):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EnchantmentPassiveNode)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(942):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EnchantmentPassiveNum)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(943):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_FracturedItem)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(944):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_GemLevel)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(945):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_GemQualityType)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(946):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_HasCruciblePassiveTree)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(947):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_HasEaterOfWorldsImplicit)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(948):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_HasEnchantment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(949):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_HasExplicitMod)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(950):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_HasImplicitMod)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(951):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_HasInfluence)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(952):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_HasSearingExarchImplicit)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(953):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Height)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(954):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Identified)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(955):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ItemLevel)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(956):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LinkedSockets)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(957):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_MapTier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(958):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_MemoryStrands)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(959):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Mirrored)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(960):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Quality)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(961):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Rarity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(962):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Replica)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(963):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Scourged)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(964):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ShapedMap)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(965):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ShaperItem)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(966):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SocketGroup)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(967):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Sockets)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(968):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_StackSize)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(969):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SynthesisedItem)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(970):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_TransfiguredGem)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(971):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_UberBlightedMap)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(972):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_UnidentifiedItemTier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(973):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_WaystoneTier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(974):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Width)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(975):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ZanaMemory)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(976):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PlayAlertSound)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('P') {
			state = uint16(626)
			goto next_state
		}
		return result
	case int32(977):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_None)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(978):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PlayAlertSoundPositional)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(979):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_CustomAlertSound)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('O') {
			state = uint16(697)
			goto next_state
		}
		return result
	case int32(980):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_action_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(981):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_CustomAlertSoundOptional)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(982):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DisableDropSound)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('I') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(983):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EnableDropSound)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('I') {
			state = uint16(361)
			goto next_state
		}
		return result
	case int32(984):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DisableDropSoundIfAlertSound)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(985):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EnableDropSoundIfAlertSound)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(986):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_MinimapIcon)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(987):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DASH1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(988):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PlayEffect)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(989):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Temp)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(990):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SetBackgroundColor)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(991):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SetBorderColor)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(992):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SetFontSize)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(993):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SetTextColor)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(994):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Continue)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(995):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__equal_operator_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(996):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__equal_operator_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('=') {
			state = uint16(995)
			goto next_state
		}
		return result
	case int32(997):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__range_operator_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(998):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__range_operator_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('=') {
			state = uint16(997)
			goto next_state
		}
		return result
	case int32(999):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_boolean)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(1000):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_boolean)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		v2 = uintptr(unsafe.Pointer(&aux_sym_string_token2_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(680) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _34
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _34
	_34:
		if v4 != 0 {
			state = uint16(1078)
			goto next_state
		}
		return result
	case int32(1001):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(1002):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('N') {
			state = uint16(616)
			goto next_state
		}
		return result
	case int32(1003):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_quality_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(1004):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Superior)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(1005):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Divergent)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(1006):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Anomalous)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(1007):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Phantasmal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(1008):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_rarity_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(1009):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Normal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(1010):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Magic)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(1011):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Rare)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(1012):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Unique)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(1013):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_influence_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(1014):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Shaper)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(1015):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Shaper)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('I') {
			state = uint16(837)
			goto next_state
		}
		return result
	case int32(1016):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Elder)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(1017):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Elder)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('I') {
			state = uint16(836)
			goto next_state
		}
		if lookahead1 == int32('M') {
			state = uint16(103)
			goto next_state
		}
		return result
	case int32(1018):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Crusader)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(1019):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Hunter)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(1020):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Redeemer)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(1021):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Warlord)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(1022):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_sockets_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('0') {
			state = uint16(1083)
			goto next_state
		}
		if lookahead1 == int32('A') || lookahead1 == int32('B') || lookahead1 == int32('D') || lookahead1 == int32('G') || lookahead1 == int32('R') || lookahead1 == int32('W') {
			state = uint16(1032)
			goto next_state
		}
		if int32('1') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(1085)
			goto next_state
		}
		return result
	case int32(1023):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_sockets_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('5') {
			state = uint16(1084)
			goto next_state
		}
		if int32('6') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(1085)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('4') {
			state = uint16(1085)
			goto next_state
		}
		if lookahead1 == int32('A') || lookahead1 == int32('B') || lookahead1 == int32('D') || lookahead1 == int32('G') || lookahead1 == int32('R') || lookahead1 == int32('W') {
			state = uint16(1032)
			goto next_state
		}
		return result
	case int32(1024):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_sockets_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i10 = uint32(0)
		for {
			if !(uint64(i10) < libc.Uint64FromInt64(36)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token10[i10]) == lookahead1 {
				state = map_token10[i10+uint32(1)]
				goto next_state
			}
			goto _35
		_35:
			;
			i10 = i10 + uint32(2)
		}
		return result
	case int32(1025):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_sockets_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('a') {
			state = uint16(408)
			goto next_state
		}
		if lookahead1 == int32('e') {
			state = uint16(182)
			goto next_state
		}
		if lookahead1 == int32('A') || lookahead1 == int32('B') || lookahead1 == int32('D') || lookahead1 == int32('G') || lookahead1 == int32('R') || lookahead1 == int32('W') {
			state = uint16(1032)
			goto next_state
		}
		return result
	case int32(1026):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_sockets_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i11 = uint32(0)
		for {
			if !(uint64(i11) < libc.Uint64FromInt64(36)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token11[i11]) == lookahead1 {
				state = map_token11[i11+uint32(1)]
				goto next_state
			}
			goto _36
		_36:
			;
			i11 = i11 + uint32(2)
		}
		return result
	case int32(1027):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_sockets_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('e') {
			state = uint16(511)
			goto next_state
		}
		if lookahead1 == int32('r') {
			state = uint16(231)
			goto next_state
		}
		if lookahead1 == int32('A') || lookahead1 == int32('B') || lookahead1 == int32('D') || lookahead1 == int32('G') || lookahead1 == int32('R') || lookahead1 == int32('W') {
			state = uint16(1032)
			goto next_state
		}
		return result
	case int32(1028):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_sockets_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('i') {
			state = uint16(99)
			goto next_state
		}
		if lookahead1 == int32('r') {
			state = uint16(612)
			goto next_state
		}
		if lookahead1 == int32('A') || lookahead1 == int32('B') || lookahead1 == int32('D') || lookahead1 == int32('G') || lookahead1 == int32('R') || lookahead1 == int32('W') {
			state = uint16(1032)
			goto next_state
		}
		return result
	case int32(1029):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_sockets_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i12 = uint32(0)
		for {
			if !(uint64(i12) < libc.Uint64FromInt64(36)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token12[i12]) == lookahead1 {
				state = map_token12[i12+uint32(1)]
				goto next_state
			}
			goto _37
		_37:
			;
			i12 = i12 + uint32(2)
		}
		return result
	case int32(1030):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_sockets_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('7') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(1085)
			goto next_state
		}
		if lookahead1 == int32('A') || lookahead1 == int32('B') || lookahead1 == int32('D') || lookahead1 == int32('G') || lookahead1 == int32('R') || lookahead1 == int32('W') {
			state = uint16(1032)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('6') {
			state = uint16(1085)
			goto next_state
		}
		return result
	case int32(1031):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_sockets_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('6') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(1085)
			goto next_state
		}
		if lookahead1 == int32('A') || lookahead1 == int32('B') || lookahead1 == int32('D') || lookahead1 == int32('G') || lookahead1 == int32('R') || lookahead1 == int32('W') {
			state = uint16(1032)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('5') {
			state = uint16(1085)
			goto next_state
		}
		return result
	case int32(1032):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_sockets_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('A') || lookahead1 == int32('B') || lookahead1 == int32('D') || lookahead1 == int32('G') || lookahead1 == int32('R') || lookahead1 == int32('W') {
			state = uint16(1032)
			goto next_state
		}
		return result
	case int32(1033):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_sockets_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('A') || lookahead1 == int32('B') || lookahead1 == int32('D') || lookahead1 == int32('G') || lookahead1 == int32('R') || lookahead1 == int32('W') {
			state = uint16(1032)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(1085)
			goto next_state
		}
		return result
	case int32(1034):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_sockets_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('0') {
			state = uint16(1083)
			goto next_state
		}
		if lookahead1 == int32('A') || lookahead1 == int32('B') || lookahead1 == int32('D') || lookahead1 == int32('G') || lookahead1 == int32('R') || lookahead1 == int32('W') {
			state = uint16(1044)
			goto next_state
		}
		if int32('1') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(1085)
			goto next_state
		}
		return result
	case int32(1035):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_sockets_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('5') {
			state = uint16(1084)
			goto next_state
		}
		if int32('6') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(1085)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('4') {
			state = uint16(1085)
			goto next_state
		}
		if lookahead1 == int32('A') || lookahead1 == int32('B') || lookahead1 == int32('D') || lookahead1 == int32('G') || lookahead1 == int32('R') || lookahead1 == int32('W') {
			state = uint16(1044)
			goto next_state
		}
		return result
	case int32(1036):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_sockets_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i13 = uint32(0)
		for {
			if !(uint64(i13) < libc.Uint64FromInt64(36)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token13[i13]) == lookahead1 {
				state = map_token13[i13+uint32(1)]
				goto next_state
			}
			goto _38
		_38:
			;
			i13 = i13 + uint32(2)
		}
		return result
	case int32(1037):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_sockets_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('a') {
			state = uint16(408)
			goto next_state
		}
		if lookahead1 == int32('e') {
			state = uint16(182)
			goto next_state
		}
		if lookahead1 == int32('A') || lookahead1 == int32('B') || lookahead1 == int32('D') || lookahead1 == int32('G') || lookahead1 == int32('R') || lookahead1 == int32('W') {
			state = uint16(1044)
			goto next_state
		}
		return result
	case int32(1038):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_sockets_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i14 = uint32(0)
		for {
			if !(uint64(i14) < libc.Uint64FromInt64(36)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token14[i14]) == lookahead1 {
				state = map_token14[i14+uint32(1)]
				goto next_state
			}
			goto _39
		_39:
			;
			i14 = i14 + uint32(2)
		}
		return result
	case int32(1039):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_sockets_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('e') {
			state = uint16(511)
			goto next_state
		}
		if lookahead1 == int32('r') {
			state = uint16(231)
			goto next_state
		}
		if lookahead1 == int32('A') || lookahead1 == int32('B') || lookahead1 == int32('D') || lookahead1 == int32('G') || lookahead1 == int32('R') || lookahead1 == int32('W') {
			state = uint16(1044)
			goto next_state
		}
		return result
	case int32(1040):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_sockets_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('i') {
			state = uint16(99)
			goto next_state
		}
		if lookahead1 == int32('r') {
			state = uint16(612)
			goto next_state
		}
		if lookahead1 == int32('A') || lookahead1 == int32('B') || lookahead1 == int32('D') || lookahead1 == int32('G') || lookahead1 == int32('R') || lookahead1 == int32('W') {
			state = uint16(1044)
			goto next_state
		}
		return result
	case int32(1041):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_sockets_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i15 = uint32(0)
		for {
			if !(uint64(i15) < libc.Uint64FromInt64(36)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token15[i15]) == lookahead1 {
				state = map_token15[i15+uint32(1)]
				goto next_state
			}
			goto _40
		_40:
			;
			i15 = i15 + uint32(2)
		}
		return result
	case int32(1042):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_sockets_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('7') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(1085)
			goto next_state
		}
		if lookahead1 == int32('A') || lookahead1 == int32('B') || lookahead1 == int32('D') || lookahead1 == int32('G') || lookahead1 == int32('R') || lookahead1 == int32('W') {
			state = uint16(1044)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('6') {
			state = uint16(1085)
			goto next_state
		}
		return result
	case int32(1043):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_sockets_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('6') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(1085)
			goto next_state
		}
		if lookahead1 == int32('A') || lookahead1 == int32('B') || lookahead1 == int32('D') || lookahead1 == int32('G') || lookahead1 == int32('R') || lookahead1 == int32('W') {
			state = uint16(1044)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('5') {
			state = uint16(1085)
			goto next_state
		}
		return result
	case int32(1044):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_sockets_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('A') || lookahead1 == int32('B') || lookahead1 == int32('D') || lookahead1 == int32('G') || lookahead1 == int32('R') || lookahead1 == int32('W') {
			state = uint16(1044)
			goto next_state
		}
		return result
	case int32(1045):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_sockets_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('A') || lookahead1 == int32('B') || lookahead1 == int32('D') || lookahead1 == int32('G') || lookahead1 == int32('R') || lookahead1 == int32('W') {
			state = uint16(1044)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(1085)
			goto next_state
		}
		return result
	case int32(1046):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Red)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(1047):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Red)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('e') {
			state = uint16(280)
			goto next_state
		}
		return result
	case int32(1048):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Green)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(1049):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Blue)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(1050):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Brown)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(1051):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_White)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(1052):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Yellow)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(1053):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Cyan)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(1054):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Grey)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(1055):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Orange)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(1056):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Pink)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(1057):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Purple)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(1058):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Circle)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(1059):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Diamond)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(1060):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Hexagon)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(1061):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Square)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(1062):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Star)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(1063):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Triangle)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(1064):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Cross)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(1065):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Moon)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(1066):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Raindrop)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(1067):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Kite)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(1068):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Pentagon)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(1069):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_UpsideDownHouse)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(1070):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(' ') {
			state = uint16(1070)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&aux_sym_string_token2_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(680) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _44
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _44
	_44:
		if v4 != 0 {
			state = uint16(1071)
			goto next_state
		}
		return result
	case int32(1071):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		v2 = uintptr(unsafe.Pointer(&aux_sym_string_token1_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(681) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _48
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _48
	_48:
		if v4 != 0 {
			state = uint16(1071)
			goto next_state
		}
		return result
	case int32(1072):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('a') {
			state = uint16(1074)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&aux_sym_string_token2_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(680) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _52
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _52
	_52:
		if v4 != 0 {
			state = uint16(1078)
			goto next_state
		}
		return result
	case int32(1073):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('e') {
			state = uint16(1000)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&aux_sym_string_token2_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(680) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _56
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _56
	_56:
		if v4 != 0 {
			state = uint16(1078)
			goto next_state
		}
		return result
	case int32(1074):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('l') {
			state = uint16(1076)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&aux_sym_string_token2_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(680) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _60
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _60
	_60:
		if v4 != 0 {
			state = uint16(1078)
			goto next_state
		}
		return result
	case int32(1075):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('r') {
			state = uint16(1077)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&aux_sym_string_token2_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(680) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _64
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _64
	_64:
		if v4 != 0 {
			state = uint16(1078)
			goto next_state
		}
		return result
	case int32(1076):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('s') {
			state = uint16(1073)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&aux_sym_string_token2_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(680) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _68
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _68
	_68:
		if v4 != 0 {
			state = uint16(1078)
			goto next_state
		}
		return result
	case int32(1077):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('u') {
			state = uint16(1073)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&aux_sym_string_token2_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(680) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _72
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _72
	_72:
		if v4 != 0 {
			state = uint16(1078)
			goto next_state
		}
		return result
	case int32(1078):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		v2 = uintptr(unsafe.Pointer(&aux_sym_string_token2_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(680) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _76
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _76
	_76:
		if v4 != 0 {
			state = uint16(1078)
			goto next_state
		}
		return result
	case int32(1079):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_file_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(' ') {
			state = uint16(1079)
			goto next_state
		}
		if lookahead1 == int32('#') {
			state = uint16(1080)
			goto next_state
		}
		if lookahead1 > int32(' ') && lookahead1 != int32('"') && lookahead1 != int32('#') && lookahead1 != int32('*') && lookahead1 != int32('<') && lookahead1 != int32('>') && lookahead1 != int32('?') && lookahead1 != int32('|') && (lookahead1 < int32(0x7f) || int32(0x9f) < lookahead1) {
			state = uint16(1081)
			goto next_state
		}
		return result
	case int32(1080):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_file_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 <= int32('\t') || int32(0x0b) <= lookahead1 && lookahead1 <= int32(0x1f) || lookahead1 == int32('"') || lookahead1 == int32('*') || lookahead1 == int32('<') || lookahead1 == int32('>') || lookahead1 == int32('?') || lookahead1 == int32('|') || int32(0x7f) <= lookahead1 && lookahead1 <= int32(0x9f) {
			state = uint16(1102)
			goto next_state
		}
		if lookahead1 > int32(0x1f) {
			state = uint16(1080)
			goto next_state
		}
		return result
	case int32(1081):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_file_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 > int32(0x1f) && lookahead1 != int32('"') && lookahead1 != int32('*') && lookahead1 != int32('<') && lookahead1 != int32('>') && lookahead1 != int32('?') && lookahead1 != int32('|') && (lookahead1 < int32(0x7f) || int32(0x9f) < lookahead1) {
			state = uint16(1081)
			goto next_state
		}
		return result
	case int32(1082):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(1083):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('0') {
			state = uint16(1085)
			goto next_state
		}
		if int32('1') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(1085)
			goto next_state
		}
		return result
	case int32(1084):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('6') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(1085)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('5') {
			state = uint16(1085)
			goto next_state
		}
		return result
	case int32(1085):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(1085)
			goto next_state
		}
		return result
	case int32(1086):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__id_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(1087):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__id_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('6') {
			state = uint16(1086)
			goto next_state
		}
		return result
	case int32(1088):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__volume_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(1089):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__volume_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('0') {
			state = uint16(1088)
			goto next_state
		}
		return result
	case int32(1090):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__volume_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('0') {
			state = uint16(1089)
			goto next_state
		}
		if int32('1') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(1088)
			goto next_state
		}
		return result
	case int32(1091):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__volume_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(1088)
			goto next_state
		}
		return result
	case int32(1092):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__volume_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(1091)
			goto next_state
		}
		return result
	case int32(1093):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__color_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(1094):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__color_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('5') {
			state = uint16(1095)
			goto next_state
		}
		if int32('6') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(1093)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('4') {
			state = uint16(1096)
			goto next_state
		}
		return result
	case int32(1095):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__color_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('5') {
			state = uint16(1093)
			goto next_state
		}
		return result
	case int32(1096):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__color_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(1093)
			goto next_state
		}
		return result
	case int32(1097):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__color_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(1096)
			goto next_state
		}
		return result
	case int32(1098):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__icon_size_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(1099):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__font_size_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(1100):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__font_size_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('5') {
			state = uint16(1099)
			goto next_state
		}
		return result
	case int32(1101):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__font_size_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(1099)
			goto next_state
		}
		return result
	case int32(1102):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 != 0 && lookahead1 != int32('\n') {
			state = uint16(1102)
			goto next_state
		}
		return result
	case int32(1103):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__space)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\t') || lookahead1 == int32(' ') {
			state = uint16(1103)
			goto next_state
		}
		return result
	case int32(1104):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__eol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var map_token = [74]uint16_t{
	0:  uint16('\n'),
	1:  uint16(1104),
	2:  uint16('\r'),
	3:  uint16(3),
	4:  uint16('!'),
	5:  uint16(996),
	6:  uint16('"'),
	7:  uint16(1002),
	8:  uint16('#'),
	9:  uint16(1102),
	10: uint16('-'),
	11: uint16(19),
	12: uint16('0'),
	13: uint16(1032),
	14: uint16('1'),
	15: uint16(1030),
	16: uint16('2'),
	17: uint16(1023),
	18: uint16('3'),
	19: uint16(1022),
	20: uint16('4'),
	21: uint16(1031),
	22: uint16('<'),
	23: uint16(998),
	24: uint16('='),
	25: uint16(996),
	26: uint16('>'),
	27: uint16(998),
	28: uint16('A'),
	29: uint16(1029),
	30: uint16('B'),
	31: uint16(1024),
	32: uint16('C'),
	33: uint16(395),
	34: uint16('D'),
	35: uint16(1028),
	36: uint16('E'),
	37: uint16(460),
	38: uint16('F'),
	39: uint16(83),
	40: uint16('G'),
	41: uint16(1027),
	42: uint16('H'),
	43: uint16(95),
	44: uint16('I'),
	45: uint16(207),
	46: uint16('K'),
	47: uint16(421),
	48: uint16('L'),
	49: uint16(402),
	50: uint16('M'),
	51: uint16(84),
	52: uint16('N'),
	53: uint16(611),
	54: uint16('O'),
	55: uint16(679),
	56: uint16('P'),
	57: uint16(270),
	58: uint16('Q'),
	59: uint16(860),
	60: uint16('R'),
	61: uint16(1025),
	62: uint16('S'),
	63: uint16(162),
	64: uint16('T'),
	65: uint16(269),
	66: uint16('U'),
	67: uint16(153),
	68: uint16('W'),
	69: uint16(1026),
	70: uint16('Y'),
	71: uint16(265),
	72: uint16('Z'),
	73: uint16(100),
}

var map_token1 = [30]uint16_t{
	0:  uint16('"'),
	1:  uint16(1001),
	2:  uint16('#'),
	3:  uint16(1102),
	4:  uint16('1'),
	5:  uint16(1087),
	6:  uint16('A'),
	7:  uint16(579),
	8:  uint16('C'),
	9:  uint16(753),
	10: uint16('D'),
	11: uint16(430),
	12: uint16('E'),
	13: uint16(503),
	14: uint16('H'),
	15: uint16(859),
	16: uint16('M'),
	17: uint16(108),
	18: uint16('N'),
	19: uint16(666),
	20: uint16('P'),
	21: uint16(391),
	22: uint16('R'),
	23: uint16(136),
	24: uint16('S'),
	25: uint16(386),
	26: uint16('U'),
	27: uint16(578),
	28: uint16('W'),
	29: uint16(126),
}

var map_token2 = [18]uint16_t{
	0:  uint16('#'),
	1:  uint16(1102),
	2:  uint16('-'),
	3:  uint16(19),
	4:  uint16('C'),
	5:  uint16(767),
	6:  uint16('E'),
	7:  uint16(504),
	8:  uint16('H'),
	9:  uint16(886),
	10: uint16('N'),
	11: uint16(658),
	12: uint16('R'),
	13: uint16(349),
	14: uint16('S'),
	15: uint16(390),
	16: uint16('W'),
	17: uint16(144),
}

var map_token3 = [18]uint16_t{
	0:  uint16('#'),
	1:  uint16(1102),
	2:  uint16('1'),
	3:  uint16(1087),
	4:  uint16('C'),
	5:  uint16(753),
	6:  uint16('E'),
	7:  uint16(503),
	8:  uint16('H'),
	9:  uint16(859),
	10: uint16('N'),
	11: uint16(654),
	12: uint16('R'),
	13: uint16(327),
	14: uint16('S'),
	15: uint16(385),
	16: uint16('W'),
	17: uint16(126),
}

var map_token4 = [16]uint16_t{
	0:  uint16('c'),
	1:  uint16(609),
	2:  uint16('e'),
	3:  uint16(800),
	4:  uint16('h'),
	5:  uint16(128),
	6:  uint16('o'),
	7:  uint16(172),
	8:  uint16('q'),
	9:  uint16(864),
	10: uint16('t'),
	11: uint16(88),
	12: uint16('u'),
	13: uint16(687),
	14: uint16('y'),
	15: uint16(562),
}

var map_token5 = [16]uint16_t{
	0:  uint16('c'),
	1:  uint16(609),
	2:  uint16('e'),
	3:  uint16(800),
	4:  uint16('h'),
	5:  uint16(135),
	6:  uint16('o'),
	7:  uint16(172),
	8:  uint16('q'),
	9:  uint16(864),
	10: uint16('t'),
	11: uint16(88),
	12: uint16('u'),
	13: uint16(687),
	14: uint16('y'),
	15: uint16(562),
}

var map_token6 = [74]uint16_t{
	0:  uint16('\n'),
	1:  uint16(1104),
	2:  uint16('\r'),
	3:  uint16(3),
	4:  uint16('!'),
	5:  uint16(996),
	6:  uint16('"'),
	7:  uint16(58),
	8:  uint16('#'),
	9:  uint16(1102),
	10: uint16('-'),
	11: uint16(19),
	12: uint16('0'),
	13: uint16(1044),
	14: uint16('1'),
	15: uint16(1042),
	16: uint16('2'),
	17: uint16(1035),
	18: uint16('3'),
	19: uint16(1034),
	20: uint16('4'),
	21: uint16(1043),
	22: uint16('<'),
	23: uint16(998),
	24: uint16('='),
	25: uint16(996),
	26: uint16('>'),
	27: uint16(998),
	28: uint16('A'),
	29: uint16(1041),
	30: uint16('B'),
	31: uint16(1036),
	32: uint16('C'),
	33: uint16(395),
	34: uint16('D'),
	35: uint16(1040),
	36: uint16('E'),
	37: uint16(460),
	38: uint16('F'),
	39: uint16(83),
	40: uint16('G'),
	41: uint16(1039),
	42: uint16('H'),
	43: uint16(95),
	44: uint16('I'),
	45: uint16(207),
	46: uint16('K'),
	47: uint16(421),
	48: uint16('L'),
	49: uint16(402),
	50: uint16('M'),
	51: uint16(84),
	52: uint16('N'),
	53: uint16(611),
	54: uint16('O'),
	55: uint16(679),
	56: uint16('P'),
	57: uint16(270),
	58: uint16('Q'),
	59: uint16(860),
	60: uint16('R'),
	61: uint16(1037),
	62: uint16('S'),
	63: uint16(162),
	64: uint16('T'),
	65: uint16(269),
	66: uint16('U'),
	67: uint16(153),
	68: uint16('W'),
	69: uint16(1038),
	70: uint16('Y'),
	71: uint16(265),
	72: uint16('Z'),
	73: uint16(100),
}

var map_token7 = [64]uint16_t{
	0:  uint16('\n'),
	1:  uint16(1104),
	2:  uint16('\r'),
	3:  uint16(3),
	4:  uint16('!'),
	5:  uint16(996),
	6:  uint16('"'),
	7:  uint16(1001),
	8:  uint16('#'),
	9:  uint16(1102),
	10: uint16('0'),
	11: uint16(1082),
	12: uint16('<'),
	13: uint16(998),
	14: uint16('='),
	15: uint16(996),
	16: uint16('>'),
	17: uint16(998),
	18: uint16('A'),
	19: uint16(459),
	20: uint16('B'),
	21: uint16(82),
	22: uint16('C'),
	23: uint16(396),
	24: uint16('D'),
	25: uint16(397),
	26: uint16('E'),
	27: uint16(502),
	28: uint16('F'),
	29: uint16(722),
	30: uint16('G'),
	31: uint16(230),
	32: uint16('H'),
	33: uint16(94),
	34: uint16('I'),
	35: uint16(207),
	36: uint16('K'),
	37: uint16(421),
	38: uint16('L'),
	39: uint16(402),
	40: uint16('M'),
	41: uint16(84),
	42: uint16('N'),
	43: uint16(611),
	44: uint16('O'),
	45: uint16(754),
	46: uint16('P'),
	47: uint16(270),
	48: uint16('Q'),
	49: uint16(860),
	50: uint16('R'),
	51: uint16(85),
	52: uint16('S'),
	53: uint16(163),
	54: uint16('T'),
	55: uint16(714),
	56: uint16('U'),
	57: uint16(153),
	58: uint16('W'),
	59: uint16(104),
	60: uint16('Y'),
	61: uint16(265),
	62: uint16('Z'),
	63: uint16(100),
}

var map_token8 = [62]uint16_t{
	0:  uint16('\n'),
	1:  uint16(1104),
	2:  uint16('\r'),
	3:  uint16(3),
	4:  uint16('!'),
	5:  uint16(996),
	6:  uint16('#'),
	7:  uint16(1102),
	8:  uint16('0'),
	9:  uint16(1082),
	10: uint16('<'),
	11: uint16(998),
	12: uint16('='),
	13: uint16(996),
	14: uint16('>'),
	15: uint16(998),
	16: uint16('A'),
	17: uint16(459),
	18: uint16('B'),
	19: uint16(82),
	20: uint16('C'),
	21: uint16(396),
	22: uint16('D'),
	23: uint16(397),
	24: uint16('E'),
	25: uint16(502),
	26: uint16('F'),
	27: uint16(722),
	28: uint16('G'),
	29: uint16(230),
	30: uint16('H'),
	31: uint16(94),
	32: uint16('I'),
	33: uint16(207),
	34: uint16('K'),
	35: uint16(421),
	36: uint16('L'),
	37: uint16(402),
	38: uint16('M'),
	39: uint16(84),
	40: uint16('N'),
	41: uint16(611),
	42: uint16('O'),
	43: uint16(754),
	44: uint16('P'),
	45: uint16(270),
	46: uint16('Q'),
	47: uint16(860),
	48: uint16('R'),
	49: uint16(85),
	50: uint16('S'),
	51: uint16(163),
	52: uint16('T'),
	53: uint16(714),
	54: uint16('U'),
	55: uint16(153),
	56: uint16('W'),
	57: uint16(104),
	58: uint16('Y'),
	59: uint16(265),
	60: uint16('Z'),
	61: uint16(100),
}

var map_token9 = [18]uint16_t{
	0:  uint16('\n'),
	1:  uint16(1104),
	2:  uint16('\r'),
	3:  uint16(3),
	4:  uint16('#'),
	5:  uint16(1102),
	6:  uint16('H'),
	7:  uint16(422),
	8:  uint16('I'),
	9:  uint16(510),
	10: uint16('M'),
	11: uint16(419),
	12: uint16('S'),
	13: uint16(383),
	14: uint16('\t'),
	15: uint16(1103),
	16: uint16(' '),
	17: uint16(1103),
}

var map_token10 = [18]uint16_t{
	0:  uint16('a'),
	1:  uint16(780),
	2:  uint16('l'),
	3:  uint16(398),
	4:  uint16('r'),
	5:  uint16(608),
	6:  uint16('A'),
	7:  uint16(1032),
	8:  uint16('B'),
	9:  uint16(1032),
	10: uint16('D'),
	11: uint16(1032),
	12: uint16('G'),
	13: uint16(1032),
	14: uint16('R'),
	15: uint16(1032),
	16: uint16('W'),
	17: uint16(1032),
}

var map_token11 = [18]uint16_t{
	0:  uint16('a'),
	1:  uint16(728),
	2:  uint16('h'),
	3:  uint16(427),
	4:  uint16('i'),
	5:  uint16(209),
	6:  uint16('A'),
	7:  uint16(1032),
	8:  uint16('B'),
	9:  uint16(1032),
	10: uint16('D'),
	11: uint16(1032),
	12: uint16('G'),
	13: uint16(1032),
	14: uint16('R'),
	15: uint16(1032),
	16: uint16('W'),
	17: uint16(1032),
}

var map_token12 = [18]uint16_t{
	0:  uint16('l'),
	1:  uint16(835),
	2:  uint16('n'),
	3:  uint16(614),
	4:  uint16('r'),
	5:  uint16(157),
	6:  uint16('A'),
	7:  uint16(1032),
	8:  uint16('B'),
	9:  uint16(1032),
	10: uint16('D'),
	11: uint16(1032),
	12: uint16('G'),
	13: uint16(1032),
	14: uint16('R'),
	15: uint16(1032),
	16: uint16('W'),
	17: uint16(1032),
}

var map_token13 = [18]uint16_t{
	0:  uint16('a'),
	1:  uint16(780),
	2:  uint16('l'),
	3:  uint16(398),
	4:  uint16('r'),
	5:  uint16(608),
	6:  uint16('A'),
	7:  uint16(1044),
	8:  uint16('B'),
	9:  uint16(1044),
	10: uint16('D'),
	11: uint16(1044),
	12: uint16('G'),
	13: uint16(1044),
	14: uint16('R'),
	15: uint16(1044),
	16: uint16('W'),
	17: uint16(1044),
}

var map_token14 = [18]uint16_t{
	0:  uint16('a'),
	1:  uint16(728),
	2:  uint16('h'),
	3:  uint16(427),
	4:  uint16('i'),
	5:  uint16(209),
	6:  uint16('A'),
	7:  uint16(1044),
	8:  uint16('B'),
	9:  uint16(1044),
	10: uint16('D'),
	11: uint16(1044),
	12: uint16('G'),
	13: uint16(1044),
	14: uint16('R'),
	15: uint16(1044),
	16: uint16('W'),
	17: uint16(1044),
}

var map_token15 = [18]uint16_t{
	0:  uint16('l'),
	1:  uint16(835),
	2:  uint16('n'),
	3:  uint16(614),
	4:  uint16('r'),
	5:  uint16(157),
	6:  uint16('A'),
	7:  uint16(1044),
	8:  uint16('B'),
	9:  uint16(1044),
	10: uint16('D'),
	11: uint16(1044),
	12: uint16('G'),
	13: uint16(1044),
	14: uint16('R'),
	15: uint16(1044),
	16: uint16('W'),
	17: uint16(1044),
}

var ts_lex_modes = [195]TSLexMode{
	0: {},
	1: {},
	2: {
		Flex_state: uint16(915),
	},
	3: {
		Flex_state: uint16(915),
	},
	4: {
		Flex_state: uint16(915),
	},
	5: {
		Flex_state: uint16(915),
	},
	6: {
		Flex_state: uint16(915),
	},
	7: {
		Flex_state: uint16(915),
	},
	8: {
		Flex_state: uint16(915),
	},
	9: {
		Flex_state: uint16(915),
	},
	10: {
		Flex_state: uint16(915),
	},
	11: {
		Flex_state: uint16(915),
	},
	12: {
		Flex_state: uint16(915),
	},
	13: {
		Flex_state: uint16(915),
	},
	14: {
		Flex_state: uint16(915),
	},
	15: {
		Flex_state: uint16(915),
	},
	16: {
		Flex_state: uint16(915),
	},
	17: {
		Flex_state: uint16(915),
	},
	18: {
		Flex_state: uint16(915),
	},
	19: {
		Flex_state: uint16(915),
	},
	20: {
		Flex_state: uint16(915),
	},
	21: {
		Flex_state: uint16(915),
	},
	22: {
		Flex_state: uint16(915),
	},
	23: {
		Flex_state: uint16(915),
	},
	24: {
		Flex_state: uint16(915),
	},
	25: {
		Flex_state: uint16(915),
	},
	26: {
		Flex_state: uint16(915),
	},
	27: {
		Flex_state: uint16(915),
	},
	28: {
		Flex_state: uint16(915),
	},
	29: {
		Flex_state: uint16(915),
	},
	30: {
		Flex_state: uint16(915),
	},
	31: {
		Flex_state: uint16(915),
	},
	32: {
		Flex_state: uint16(915),
	},
	33: {
		Flex_state: uint16(915),
	},
	34: {
		Flex_state: uint16(915),
	},
	35: {
		Flex_state: uint16(915),
	},
	36: {
		Flex_state: uint16(915),
	},
	37: {
		Flex_state: uint16(915),
	},
	38: {
		Flex_state: uint16(915),
	},
	39: {
		Flex_state: uint16(10),
	},
	40: {},
	41: {
		Flex_state: uint16(4),
	},
	42: {},
	43: {
		Flex_state: uint16(5),
	},
	44: {
		Flex_state: uint16(917),
	},
	45: {
		Flex_state: uint16(917),
	},
	46: {
		Flex_state: uint16(915),
	},
	47: {
		Flex_state: uint16(915),
	},
	48: {},
	49: {
		Flex_state: uint16(915),
	},
	50: {
		Flex_state: uint16(4),
	},
	51: {
		Flex_state: uint16(4),
	},
	52: {
		Flex_state: uint16(4),
	},
	53: {
		Flex_state: uint16(4),
	},
	54: {
		Flex_state: uint16(9),
	},
	55: {
		Flex_state: uint16(917),
	},
	56: {
		Flex_state: uint16(917),
	},
	57: {
		Flex_state: uint16(917),
	},
	58: {
		Flex_state: uint16(917),
	},
	59: {
		Flex_state: uint16(917),
	},
	60: {
		Flex_state: uint16(917),
	},
	61: {
		Flex_state: uint16(917),
	},
	62: {
		Flex_state: uint16(917),
	},
	63: {
		Flex_state: uint16(917),
	},
	64: {
		Flex_state: uint16(917),
	},
	65: {
		Flex_state: uint16(5),
	},
	66: {
		Flex_state: uint16(10),
	},
	67: {},
	68: {
		Flex_state: uint16(917),
	},
	69: {
		Flex_state: uint16(11),
	},
	70: {
		Flex_state: uint16(917),
	},
	71: {
		Flex_state: uint16(917),
	},
	72: {
		Flex_state: uint16(917),
	},
	73: {
		Flex_state: uint16(4),
	},
	74: {
		Flex_state: uint16(917),
	},
	75: {
		Flex_state: uint16(917),
	},
	76: {
		Flex_state: uint16(4),
	},
	77: {
		Flex_state: uint16(917),
	},
	78: {
		Flex_state: uint16(917),
	},
	79: {
		Flex_state: uint16(4),
	},
	80: {
		Flex_state: uint16(917),
	},
	81: {
		Flex_state: uint16(917),
	},
	82: {
		Flex_state: uint16(4),
	},
	83: {
		Flex_state: uint16(917),
	},
	84: {
		Flex_state: uint16(917),
	},
	85: {
		Flex_state: uint16(917),
	},
	86: {
		Flex_state: uint16(917),
	},
	87: {
		Flex_state: uint16(917),
	},
	88: {
		Flex_state: uint16(917),
	},
	89: {
		Flex_state: uint16(917),
	},
	90: {
		Flex_state: uint16(917),
	},
	91: {
		Flex_state: uint16(917),
	},
	92: {
		Flex_state: uint16(917),
	},
	93: {
		Flex_state: uint16(917),
	},
	94: {
		Flex_state: uint16(917),
	},
	95: {
		Flex_state: uint16(917),
	},
	96: {
		Flex_state: uint16(917),
	},
	97: {
		Flex_state: uint16(917),
	},
	98: {
		Flex_state: uint16(917),
	},
	99: {
		Flex_state: uint16(917),
	},
	100: {
		Flex_state: uint16(917),
	},
	101: {
		Flex_state: uint16(917),
	},
	102: {
		Flex_state: uint16(917),
	},
	103: {
		Flex_state: uint16(917),
	},
	104: {
		Flex_state: uint16(917),
	},
	105: {
		Flex_state: uint16(915),
	},
	106: {
		Flex_state: uint16(915),
	},
	107: {
		Flex_state: uint16(9),
	},
	108: {
		Flex_state: uint16(16),
	},
	109: {
		Flex_state: uint16(14),
	},
	110: {
		Flex_state: uint16(917),
	},
	111: {
		Flex_state: uint16(14),
	},
	112: {
		Flex_state: uint16(917),
	},
	113: {
		Flex_state: uint16(917),
	},
	114: {
		Flex_state: uint16(9),
	},
	115: {
		Flex_state: uint16(917),
	},
	116: {
		Flex_state: uint16(917),
	},
	117: {
		Flex_state: uint16(917),
	},
	118: {
		Flex_state: uint16(917),
	},
	119: {
		Flex_state: uint16(917),
	},
	120: {
		Flex_state: uint16(917),
	},
	121: {
		Flex_state: uint16(917),
	},
	122: {
		Flex_state: uint16(917),
	},
	123: {
		Flex_state: uint16(917),
	},
	124: {
		Flex_state: uint16(917),
	},
	125: {
		Flex_state: uint16(917),
	},
	126: {
		Flex_state: uint16(917),
	},
	127: {
		Flex_state: uint16(917),
	},
	128: {
		Flex_state: uint16(917),
	},
	129: {
		Flex_state: uint16(9),
	},
	130: {
		Flex_state: uint16(917),
	},
	131: {
		Flex_state: uint16(917),
	},
	132: {
		Flex_state: uint16(917),
	},
	133: {
		Flex_state: uint16(9),
	},
	134: {
		Flex_state: uint16(917),
	},
	135: {
		Flex_state: uint16(917),
	},
	136: {
		Flex_state: uint16(14),
	},
	137: {},
	138: {},
	139: {},
	140: {
		Flex_state: uint16(917),
	},
	141: {
		Flex_state: uint16(915),
	},
	142: {
		Flex_state: uint16(917),
	},
	143: {},
	144: {},
	145: {
		Flex_state: uint16(10),
	},
	146: {
		Flex_state: uint16(1),
	},
	147: {
		Flex_state: uint16(917),
	},
	148: {},
	149: {},
	150: {
		Flex_state: uint16(917),
	},
	151: {},
	152: {
		Flex_state: uint16(917),
	},
	153: {
		Flex_state: uint16(917),
	},
	154: {
		Flex_state: uint16(917),
	},
	155: {
		Flex_state: uint16(917),
	},
	156: {
		Flex_state: uint16(917),
	},
	157: {
		Flex_state: uint16(917),
	},
	158: {
		Flex_state: uint16(917),
	},
	159: {
		Flex_state: uint16(917),
	},
	160: {
		Flex_state: uint16(915),
	},
	161: {},
	162: {},
	163: {},
	164: {},
	165: {
		Flex_state: uint16(915),
	},
	166: {
		Flex_state: uint16(917),
	},
	167: {},
	168: {},
	169: {
		Flex_state: uint16(917),
	},
	170: {},
	171: {
		Flex_state: uint16(18),
	},
	172: {},
	173: {},
	174: {
		Flex_state: uint16(915),
	},
	175: {
		Flex_state: uint16(917),
	},
	176: {},
	177: {
		Flex_state: uint16(2),
	},
	178: {},
	179: {},
	180: {},
	181: {
		Flex_state: uint16(11),
	},
	182: {
		Flex_state: uint16(10),
	},
	183: {
		Flex_state: uint16(915),
	},
	184: {
		Flex_state: uint16(915),
	},
	185: {
		Flex_state: uint16(915),
	},
	186: {},
	187: {
		Flex_state: uint16(917),
	},
	188: {},
	189: {
		Flex_state: uint16(917),
	},
	190: {},
	191: {
		Flex_state: uint16(915),
	},
	192: {
		Flex_state: uint16(915),
	},
	193: {
		Flex_state: uint16(1),
	},
	194: {
		Flex_state: uint16(917),
	},
}

var ts_parse_table = [33][167]uint16_t{
	0: {
		0:   uint16(1),
		1:   uint16(1),
		2:   uint16(1),
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
		126: uint16(1),
		127: uint16(1),
		128: uint16(1),
		129: uint16(1),
		130: uint16(1),
		131: uint16(1),
		132: uint16(3),
		134: uint16(1),
	},
	1: {
		0:   uint16(5),
		1:   uint16(7),
		2:   uint16(7),
		3:   uint16(7),
		4:   uint16(9),
		132: uint16(3),
		134: uint16(11),
		135: uint16(139),
		136: uint16(40),
		137: uint16(40),
		157: uint16(40),
	},
	2: {
		0:   uint16(13),
		1:   uint16(13),
		2:   uint16(13),
		3:   uint16(13),
		4:   uint16(13),
		6:   uint16(15),
		7:   uint16(15),
		8:   uint16(17),
		9:   uint16(19),
		10:  uint16(19),
		11:  uint16(19),
		12:  uint16(19),
		13:  uint16(19),
		14:  uint16(21),
		15:  uint16(19),
		16:  uint16(15),
		17:  uint16(23),
		18:  uint16(25),
		19:  uint16(19),
		20:  uint16(19),
		21:  uint16(15),
		22:  uint16(15),
		23:  uint16(27),
		24:  uint16(19),
		25:  uint16(15),
		26:  uint16(19),
		27:  uint16(29),
		28:  uint16(15),
		29:  uint16(19),
		30:  uint16(27),
		31:  uint16(31),
		32:  uint16(15),
		33:  uint16(33),
		34:  uint16(19),
		35:  uint16(19),
		36:  uint16(15),
		37:  uint16(19),
		38:  uint16(19),
		39:  uint16(19),
		40:  uint16(19),
		41:  uint16(15),
		42:  uint16(19),
		43:  uint16(35),
		44:  uint16(15),
		45:  uint16(15),
		46:  uint16(15),
		47:  uint16(15),
		48:  uint16(37),
		49:  uint16(37),
		50:  uint16(19),
		51:  uint16(15),
		52:  uint16(39),
		53:  uint16(15),
		54:  uint16(19),
		55:  uint16(19),
		56:  uint16(19),
		57:  uint16(15),
		58:  uint16(41),
		60:  uint16(43),
		61:  uint16(45),
		63:  uint16(47),
		64:  uint16(49),
		65:  uint16(49),
		66:  uint16(51),
		67:  uint16(51),
		68:  uint16(53),
		70:  uint16(55),
		72:  uint16(57),
		73:  uint16(57),
		74:  uint16(59),
		75:  uint16(57),
		76:  uint16(61),
		132: uint16(3),
		134: uint16(63),
		138: uint16(3),
		139: uint16(3),
		140: uint16(3),
		158: uint16(3),
	},
	3: {
		0:   uint16(65),
		1:   uint16(65),
		2:   uint16(65),
		3:   uint16(65),
		4:   uint16(65),
		6:   uint16(67),
		7:   uint16(67),
		8:   uint16(70),
		9:   uint16(73),
		10:  uint16(73),
		11:  uint16(73),
		12:  uint16(73),
		13:  uint16(73),
		14:  uint16(76),
		15:  uint16(73),
		16:  uint16(67),
		17:  uint16(79),
		18:  uint16(82),
		19:  uint16(73),
		20:  uint16(73),
		21:  uint16(67),
		22:  uint16(67),
		23:  uint16(85),
		24:  uint16(73),
		25:  uint16(67),
		26:  uint16(73),
		27:  uint16(88),
		28:  uint16(67),
		29:  uint16(73),
		30:  uint16(85),
		31:  uint16(91),
		32:  uint16(67),
		33:  uint16(94),
		34:  uint16(73),
		35:  uint16(73),
		36:  uint16(67),
		37:  uint16(73),
		38:  uint16(73),
		39:  uint16(73),
		40:  uint16(73),
		41:  uint16(67),
		42:  uint16(73),
		43:  uint16(97),
		44:  uint16(67),
		45:  uint16(67),
		46:  uint16(67),
		47:  uint16(67),
		48:  uint16(100),
		49:  uint16(100),
		50:  uint16(73),
		51:  uint16(67),
		52:  uint16(103),
		53:  uint16(67),
		54:  uint16(73),
		55:  uint16(73),
		56:  uint16(73),
		57:  uint16(67),
		58:  uint16(106),
		60:  uint16(109),
		61:  uint16(112),
		63:  uint16(115),
		64:  uint16(118),
		65:  uint16(118),
		66:  uint16(121),
		67:  uint16(121),
		68:  uint16(124),
		70:  uint16(127),
		72:  uint16(130),
		73:  uint16(130),
		74:  uint16(133),
		75:  uint16(130),
		76:  uint16(136),
		132: uint16(3),
		134: uint16(139),
		138: uint16(3),
		139: uint16(3),
		140: uint16(3),
		158: uint16(3),
	},
	4: {
		0:   uint16(142),
		1:   uint16(142),
		2:   uint16(142),
		3:   uint16(142),
		4:   uint16(142),
		6:   uint16(142),
		7:   uint16(142),
		8:   uint16(142),
		9:   uint16(142),
		10:  uint16(142),
		11:  uint16(142),
		12:  uint16(142),
		13:  uint16(142),
		14:  uint16(142),
		15:  uint16(142),
		16:  uint16(142),
		17:  uint16(142),
		18:  uint16(144),
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
		40:  uint16(142),
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
		58:  uint16(144),
		60:  uint16(142),
		61:  uint16(144),
		63:  uint16(142),
		64:  uint16(144),
		65:  uint16(144),
		66:  uint16(142),
		67:  uint16(142),
		68:  uint16(142),
		70:  uint16(142),
		72:  uint16(142),
		73:  uint16(142),
		74:  uint16(142),
		75:  uint16(142),
		76:  uint16(142),
		132: uint16(3),
		134: uint16(142),
	},
	5: {
		0:   uint16(146),
		1:   uint16(146),
		2:   uint16(146),
		3:   uint16(146),
		4:   uint16(146),
		6:   uint16(146),
		7:   uint16(146),
		8:   uint16(146),
		9:   uint16(146),
		10:  uint16(146),
		11:  uint16(146),
		12:  uint16(146),
		13:  uint16(146),
		14:  uint16(146),
		15:  uint16(146),
		16:  uint16(146),
		17:  uint16(146),
		18:  uint16(148),
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
		40:  uint16(146),
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
		58:  uint16(148),
		60:  uint16(146),
		61:  uint16(148),
		63:  uint16(146),
		64:  uint16(148),
		65:  uint16(148),
		66:  uint16(146),
		67:  uint16(146),
		68:  uint16(146),
		70:  uint16(146),
		72:  uint16(146),
		73:  uint16(146),
		74:  uint16(146),
		75:  uint16(146),
		76:  uint16(146),
		132: uint16(3),
		134: uint16(146),
	},
	6: {
		0:   uint16(150),
		1:   uint16(150),
		2:   uint16(150),
		3:   uint16(150),
		4:   uint16(150),
		6:   uint16(150),
		7:   uint16(150),
		8:   uint16(150),
		9:   uint16(150),
		10:  uint16(150),
		11:  uint16(150),
		12:  uint16(150),
		13:  uint16(150),
		14:  uint16(150),
		15:  uint16(150),
		16:  uint16(150),
		17:  uint16(150),
		18:  uint16(152),
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
		40:  uint16(150),
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
		58:  uint16(152),
		60:  uint16(150),
		61:  uint16(152),
		63:  uint16(150),
		64:  uint16(152),
		65:  uint16(152),
		66:  uint16(150),
		67:  uint16(150),
		68:  uint16(150),
		70:  uint16(150),
		72:  uint16(150),
		73:  uint16(150),
		74:  uint16(150),
		75:  uint16(150),
		76:  uint16(150),
		132: uint16(3),
		134: uint16(150),
	},
	7: {
		0:   uint16(154),
		1:   uint16(154),
		2:   uint16(154),
		3:   uint16(154),
		4:   uint16(154),
		6:   uint16(154),
		7:   uint16(154),
		8:   uint16(154),
		9:   uint16(154),
		10:  uint16(154),
		11:  uint16(154),
		12:  uint16(154),
		13:  uint16(154),
		14:  uint16(154),
		15:  uint16(154),
		16:  uint16(154),
		17:  uint16(154),
		18:  uint16(156),
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
		40:  uint16(154),
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
		58:  uint16(156),
		60:  uint16(154),
		61:  uint16(156),
		63:  uint16(154),
		64:  uint16(156),
		65:  uint16(156),
		66:  uint16(154),
		67:  uint16(154),
		68:  uint16(154),
		70:  uint16(154),
		72:  uint16(154),
		73:  uint16(154),
		74:  uint16(154),
		75:  uint16(154),
		76:  uint16(154),
		132: uint16(3),
		134: uint16(154),
	},
	8: {
		0:   uint16(158),
		1:   uint16(158),
		2:   uint16(158),
		3:   uint16(158),
		4:   uint16(158),
		6:   uint16(158),
		7:   uint16(158),
		8:   uint16(158),
		9:   uint16(158),
		10:  uint16(158),
		11:  uint16(158),
		12:  uint16(158),
		13:  uint16(158),
		14:  uint16(158),
		15:  uint16(158),
		16:  uint16(158),
		17:  uint16(158),
		18:  uint16(160),
		19:  uint16(158),
		20:  uint16(158),
		21:  uint16(158),
		22:  uint16(158),
		23:  uint16(158),
		24:  uint16(158),
		25:  uint16(158),
		26:  uint16(158),
		27:  uint16(158),
		28:  uint16(158),
		29:  uint16(158),
		30:  uint16(158),
		31:  uint16(158),
		32:  uint16(158),
		33:  uint16(158),
		34:  uint16(158),
		35:  uint16(158),
		36:  uint16(158),
		37:  uint16(158),
		38:  uint16(158),
		39:  uint16(158),
		40:  uint16(158),
		41:  uint16(158),
		42:  uint16(158),
		43:  uint16(158),
		44:  uint16(158),
		45:  uint16(158),
		46:  uint16(158),
		47:  uint16(158),
		48:  uint16(158),
		49:  uint16(158),
		50:  uint16(158),
		51:  uint16(158),
		52:  uint16(158),
		53:  uint16(158),
		54:  uint16(158),
		55:  uint16(158),
		56:  uint16(158),
		57:  uint16(158),
		58:  uint16(160),
		60:  uint16(158),
		61:  uint16(160),
		63:  uint16(158),
		64:  uint16(160),
		65:  uint16(160),
		66:  uint16(158),
		67:  uint16(158),
		68:  uint16(158),
		70:  uint16(158),
		72:  uint16(158),
		73:  uint16(158),
		74:  uint16(158),
		75:  uint16(158),
		76:  uint16(158),
		132: uint16(3),
		134: uint16(158),
	},
	9: {
		0:   uint16(162),
		1:   uint16(162),
		2:   uint16(162),
		3:   uint16(162),
		4:   uint16(162),
		6:   uint16(162),
		7:   uint16(162),
		8:   uint16(162),
		9:   uint16(162),
		10:  uint16(162),
		11:  uint16(162),
		12:  uint16(162),
		13:  uint16(162),
		14:  uint16(162),
		15:  uint16(162),
		16:  uint16(162),
		17:  uint16(162),
		18:  uint16(164),
		19:  uint16(162),
		20:  uint16(162),
		21:  uint16(162),
		22:  uint16(162),
		23:  uint16(162),
		24:  uint16(162),
		25:  uint16(162),
		26:  uint16(162),
		27:  uint16(162),
		28:  uint16(162),
		29:  uint16(162),
		30:  uint16(162),
		31:  uint16(162),
		32:  uint16(162),
		33:  uint16(162),
		34:  uint16(162),
		35:  uint16(162),
		36:  uint16(162),
		37:  uint16(162),
		38:  uint16(162),
		39:  uint16(162),
		40:  uint16(162),
		41:  uint16(162),
		42:  uint16(162),
		43:  uint16(162),
		44:  uint16(162),
		45:  uint16(162),
		46:  uint16(162),
		47:  uint16(162),
		48:  uint16(162),
		49:  uint16(162),
		50:  uint16(162),
		51:  uint16(162),
		52:  uint16(162),
		53:  uint16(162),
		54:  uint16(162),
		55:  uint16(162),
		56:  uint16(162),
		57:  uint16(162),
		58:  uint16(164),
		60:  uint16(162),
		61:  uint16(164),
		63:  uint16(162),
		64:  uint16(164),
		65:  uint16(164),
		66:  uint16(162),
		67:  uint16(162),
		68:  uint16(162),
		70:  uint16(162),
		72:  uint16(162),
		73:  uint16(162),
		74:  uint16(162),
		75:  uint16(162),
		76:  uint16(162),
		132: uint16(3),
		134: uint16(162),
	},
	10: {
		0:   uint16(166),
		1:   uint16(166),
		2:   uint16(166),
		3:   uint16(166),
		4:   uint16(166),
		6:   uint16(166),
		7:   uint16(166),
		8:   uint16(166),
		9:   uint16(166),
		10:  uint16(166),
		11:  uint16(166),
		12:  uint16(166),
		13:  uint16(166),
		14:  uint16(166),
		15:  uint16(166),
		16:  uint16(166),
		17:  uint16(166),
		18:  uint16(168),
		19:  uint16(166),
		20:  uint16(166),
		21:  uint16(166),
		22:  uint16(166),
		23:  uint16(166),
		24:  uint16(166),
		25:  uint16(166),
		26:  uint16(166),
		27:  uint16(166),
		28:  uint16(166),
		29:  uint16(166),
		30:  uint16(166),
		31:  uint16(166),
		32:  uint16(166),
		33:  uint16(166),
		34:  uint16(166),
		35:  uint16(166),
		36:  uint16(166),
		37:  uint16(166),
		38:  uint16(166),
		39:  uint16(166),
		40:  uint16(166),
		41:  uint16(166),
		42:  uint16(166),
		43:  uint16(166),
		44:  uint16(166),
		45:  uint16(166),
		46:  uint16(166),
		47:  uint16(166),
		48:  uint16(166),
		49:  uint16(166),
		50:  uint16(166),
		51:  uint16(166),
		52:  uint16(166),
		53:  uint16(166),
		54:  uint16(166),
		55:  uint16(166),
		56:  uint16(166),
		57:  uint16(166),
		58:  uint16(168),
		60:  uint16(166),
		61:  uint16(168),
		63:  uint16(166),
		64:  uint16(168),
		65:  uint16(168),
		66:  uint16(166),
		67:  uint16(166),
		68:  uint16(166),
		70:  uint16(166),
		72:  uint16(166),
		73:  uint16(166),
		74:  uint16(166),
		75:  uint16(166),
		76:  uint16(166),
		132: uint16(3),
		134: uint16(166),
	},
	11: {
		0:   uint16(170),
		1:   uint16(170),
		2:   uint16(170),
		3:   uint16(170),
		4:   uint16(170),
		6:   uint16(170),
		7:   uint16(170),
		8:   uint16(170),
		9:   uint16(170),
		10:  uint16(170),
		11:  uint16(170),
		12:  uint16(170),
		13:  uint16(170),
		14:  uint16(170),
		15:  uint16(170),
		16:  uint16(170),
		17:  uint16(170),
		18:  uint16(172),
		19:  uint16(170),
		20:  uint16(170),
		21:  uint16(170),
		22:  uint16(170),
		23:  uint16(170),
		24:  uint16(170),
		25:  uint16(170),
		26:  uint16(170),
		27:  uint16(170),
		28:  uint16(170),
		29:  uint16(170),
		30:  uint16(170),
		31:  uint16(170),
		32:  uint16(170),
		33:  uint16(170),
		34:  uint16(170),
		35:  uint16(170),
		36:  uint16(170),
		37:  uint16(170),
		38:  uint16(170),
		39:  uint16(170),
		40:  uint16(170),
		41:  uint16(170),
		42:  uint16(170),
		43:  uint16(170),
		44:  uint16(170),
		45:  uint16(170),
		46:  uint16(170),
		47:  uint16(170),
		48:  uint16(170),
		49:  uint16(170),
		50:  uint16(170),
		51:  uint16(170),
		52:  uint16(170),
		53:  uint16(170),
		54:  uint16(170),
		55:  uint16(170),
		56:  uint16(170),
		57:  uint16(170),
		58:  uint16(172),
		60:  uint16(170),
		61:  uint16(172),
		63:  uint16(170),
		64:  uint16(172),
		65:  uint16(172),
		66:  uint16(170),
		67:  uint16(170),
		68:  uint16(170),
		70:  uint16(170),
		72:  uint16(170),
		73:  uint16(170),
		74:  uint16(170),
		75:  uint16(170),
		76:  uint16(170),
		132: uint16(3),
		134: uint16(170),
	},
	12: {
		0:   uint16(174),
		1:   uint16(174),
		2:   uint16(174),
		3:   uint16(174),
		4:   uint16(174),
		6:   uint16(174),
		7:   uint16(174),
		8:   uint16(174),
		9:   uint16(174),
		10:  uint16(174),
		11:  uint16(174),
		12:  uint16(174),
		13:  uint16(174),
		14:  uint16(174),
		15:  uint16(174),
		16:  uint16(174),
		17:  uint16(174),
		18:  uint16(176),
		19:  uint16(174),
		20:  uint16(174),
		21:  uint16(174),
		22:  uint16(174),
		23:  uint16(174),
		24:  uint16(174),
		25:  uint16(174),
		26:  uint16(174),
		27:  uint16(174),
		28:  uint16(174),
		29:  uint16(174),
		30:  uint16(174),
		31:  uint16(174),
		32:  uint16(174),
		33:  uint16(174),
		34:  uint16(174),
		35:  uint16(174),
		36:  uint16(174),
		37:  uint16(174),
		38:  uint16(174),
		39:  uint16(174),
		40:  uint16(174),
		41:  uint16(174),
		42:  uint16(174),
		43:  uint16(174),
		44:  uint16(174),
		45:  uint16(174),
		46:  uint16(174),
		47:  uint16(174),
		48:  uint16(174),
		49:  uint16(174),
		50:  uint16(174),
		51:  uint16(174),
		52:  uint16(174),
		53:  uint16(174),
		54:  uint16(174),
		55:  uint16(174),
		56:  uint16(174),
		57:  uint16(174),
		58:  uint16(176),
		60:  uint16(174),
		61:  uint16(176),
		63:  uint16(174),
		64:  uint16(176),
		65:  uint16(176),
		66:  uint16(174),
		67:  uint16(174),
		68:  uint16(174),
		70:  uint16(174),
		72:  uint16(174),
		73:  uint16(174),
		74:  uint16(174),
		75:  uint16(174),
		76:  uint16(174),
		132: uint16(3),
		134: uint16(174),
	},
	13: {
		0:   uint16(158),
		1:   uint16(158),
		2:   uint16(158),
		3:   uint16(158),
		4:   uint16(158),
		6:   uint16(158),
		7:   uint16(158),
		8:   uint16(158),
		9:   uint16(158),
		10:  uint16(158),
		11:  uint16(158),
		12:  uint16(158),
		13:  uint16(158),
		14:  uint16(158),
		15:  uint16(158),
		16:  uint16(158),
		17:  uint16(158),
		18:  uint16(160),
		19:  uint16(158),
		20:  uint16(158),
		21:  uint16(158),
		22:  uint16(158),
		23:  uint16(158),
		24:  uint16(158),
		25:  uint16(158),
		26:  uint16(158),
		27:  uint16(158),
		28:  uint16(158),
		29:  uint16(158),
		30:  uint16(158),
		31:  uint16(158),
		32:  uint16(158),
		33:  uint16(158),
		34:  uint16(158),
		35:  uint16(158),
		36:  uint16(158),
		37:  uint16(158),
		38:  uint16(158),
		39:  uint16(158),
		40:  uint16(158),
		41:  uint16(158),
		42:  uint16(158),
		43:  uint16(158),
		44:  uint16(158),
		45:  uint16(158),
		46:  uint16(158),
		47:  uint16(158),
		48:  uint16(158),
		49:  uint16(158),
		50:  uint16(158),
		51:  uint16(158),
		52:  uint16(158),
		53:  uint16(158),
		54:  uint16(158),
		55:  uint16(158),
		56:  uint16(158),
		57:  uint16(158),
		58:  uint16(160),
		60:  uint16(158),
		61:  uint16(160),
		63:  uint16(158),
		64:  uint16(160),
		65:  uint16(160),
		66:  uint16(158),
		67:  uint16(158),
		68:  uint16(158),
		70:  uint16(158),
		72:  uint16(158),
		73:  uint16(158),
		74:  uint16(158),
		75:  uint16(158),
		76:  uint16(158),
		132: uint16(3),
		134: uint16(158),
	},
	14: {
		0:   uint16(178),
		1:   uint16(178),
		2:   uint16(178),
		3:   uint16(178),
		4:   uint16(178),
		6:   uint16(178),
		7:   uint16(178),
		8:   uint16(178),
		9:   uint16(178),
		10:  uint16(178),
		11:  uint16(178),
		12:  uint16(178),
		13:  uint16(178),
		14:  uint16(178),
		15:  uint16(178),
		16:  uint16(178),
		17:  uint16(178),
		18:  uint16(180),
		19:  uint16(178),
		20:  uint16(178),
		21:  uint16(178),
		22:  uint16(178),
		23:  uint16(178),
		24:  uint16(178),
		25:  uint16(178),
		26:  uint16(178),
		27:  uint16(178),
		28:  uint16(178),
		29:  uint16(178),
		30:  uint16(178),
		31:  uint16(178),
		32:  uint16(178),
		33:  uint16(178),
		34:  uint16(178),
		35:  uint16(178),
		36:  uint16(178),
		37:  uint16(178),
		38:  uint16(178),
		39:  uint16(178),
		40:  uint16(178),
		41:  uint16(178),
		42:  uint16(178),
		43:  uint16(178),
		44:  uint16(178),
		45:  uint16(178),
		46:  uint16(178),
		47:  uint16(178),
		48:  uint16(178),
		49:  uint16(178),
		50:  uint16(178),
		51:  uint16(178),
		52:  uint16(178),
		53:  uint16(178),
		54:  uint16(178),
		55:  uint16(178),
		56:  uint16(178),
		57:  uint16(178),
		58:  uint16(180),
		60:  uint16(178),
		61:  uint16(180),
		63:  uint16(178),
		64:  uint16(180),
		65:  uint16(180),
		66:  uint16(178),
		67:  uint16(178),
		68:  uint16(178),
		70:  uint16(178),
		72:  uint16(178),
		73:  uint16(178),
		74:  uint16(178),
		75:  uint16(178),
		76:  uint16(178),
		132: uint16(3),
		134: uint16(178),
	},
	15: {
		0:   uint16(182),
		1:   uint16(182),
		2:   uint16(182),
		3:   uint16(182),
		4:   uint16(182),
		6:   uint16(182),
		7:   uint16(182),
		8:   uint16(182),
		9:   uint16(182),
		10:  uint16(182),
		11:  uint16(182),
		12:  uint16(182),
		13:  uint16(182),
		14:  uint16(182),
		15:  uint16(182),
		16:  uint16(182),
		17:  uint16(182),
		18:  uint16(184),
		19:  uint16(182),
		20:  uint16(182),
		21:  uint16(182),
		22:  uint16(182),
		23:  uint16(182),
		24:  uint16(182),
		25:  uint16(182),
		26:  uint16(182),
		27:  uint16(182),
		28:  uint16(182),
		29:  uint16(182),
		30:  uint16(182),
		31:  uint16(182),
		32:  uint16(182),
		33:  uint16(182),
		34:  uint16(182),
		35:  uint16(182),
		36:  uint16(182),
		37:  uint16(182),
		38:  uint16(182),
		39:  uint16(182),
		40:  uint16(182),
		41:  uint16(182),
		42:  uint16(182),
		43:  uint16(182),
		44:  uint16(182),
		45:  uint16(182),
		46:  uint16(182),
		47:  uint16(182),
		48:  uint16(182),
		49:  uint16(182),
		50:  uint16(182),
		51:  uint16(182),
		52:  uint16(182),
		53:  uint16(182),
		54:  uint16(182),
		55:  uint16(182),
		56:  uint16(182),
		57:  uint16(182),
		58:  uint16(184),
		60:  uint16(182),
		61:  uint16(184),
		63:  uint16(182),
		64:  uint16(184),
		65:  uint16(184),
		66:  uint16(182),
		67:  uint16(182),
		68:  uint16(182),
		70:  uint16(182),
		72:  uint16(182),
		73:  uint16(182),
		74:  uint16(182),
		75:  uint16(182),
		76:  uint16(182),
		132: uint16(3),
		134: uint16(182),
	},
	16: {
		0:   uint16(186),
		1:   uint16(186),
		2:   uint16(186),
		3:   uint16(186),
		4:   uint16(186),
		6:   uint16(186),
		7:   uint16(186),
		8:   uint16(186),
		9:   uint16(186),
		10:  uint16(186),
		11:  uint16(186),
		12:  uint16(186),
		13:  uint16(186),
		14:  uint16(186),
		15:  uint16(186),
		16:  uint16(186),
		17:  uint16(186),
		18:  uint16(188),
		19:  uint16(186),
		20:  uint16(186),
		21:  uint16(186),
		22:  uint16(186),
		23:  uint16(186),
		24:  uint16(186),
		25:  uint16(186),
		26:  uint16(186),
		27:  uint16(186),
		28:  uint16(186),
		29:  uint16(186),
		30:  uint16(186),
		31:  uint16(186),
		32:  uint16(186),
		33:  uint16(186),
		34:  uint16(186),
		35:  uint16(186),
		36:  uint16(186),
		37:  uint16(186),
		38:  uint16(186),
		39:  uint16(186),
		40:  uint16(186),
		41:  uint16(186),
		42:  uint16(186),
		43:  uint16(186),
		44:  uint16(186),
		45:  uint16(186),
		46:  uint16(186),
		47:  uint16(186),
		48:  uint16(186),
		49:  uint16(186),
		50:  uint16(186),
		51:  uint16(186),
		52:  uint16(186),
		53:  uint16(186),
		54:  uint16(186),
		55:  uint16(186),
		56:  uint16(186),
		57:  uint16(186),
		58:  uint16(188),
		60:  uint16(186),
		61:  uint16(188),
		63:  uint16(186),
		64:  uint16(188),
		65:  uint16(188),
		66:  uint16(186),
		67:  uint16(186),
		68:  uint16(186),
		70:  uint16(186),
		72:  uint16(186),
		73:  uint16(186),
		74:  uint16(186),
		75:  uint16(186),
		76:  uint16(186),
		132: uint16(3),
		134: uint16(186),
	},
	17: {
		0:   uint16(190),
		1:   uint16(190),
		2:   uint16(190),
		3:   uint16(190),
		4:   uint16(190),
		6:   uint16(190),
		7:   uint16(190),
		8:   uint16(190),
		9:   uint16(190),
		10:  uint16(190),
		11:  uint16(190),
		12:  uint16(190),
		13:  uint16(190),
		14:  uint16(190),
		15:  uint16(190),
		16:  uint16(190),
		17:  uint16(190),
		18:  uint16(192),
		19:  uint16(190),
		20:  uint16(190),
		21:  uint16(190),
		22:  uint16(190),
		23:  uint16(190),
		24:  uint16(190),
		25:  uint16(190),
		26:  uint16(190),
		27:  uint16(190),
		28:  uint16(190),
		29:  uint16(190),
		30:  uint16(190),
		31:  uint16(190),
		32:  uint16(190),
		33:  uint16(190),
		34:  uint16(190),
		35:  uint16(190),
		36:  uint16(190),
		37:  uint16(190),
		38:  uint16(190),
		39:  uint16(190),
		40:  uint16(190),
		41:  uint16(190),
		42:  uint16(190),
		43:  uint16(190),
		44:  uint16(190),
		45:  uint16(190),
		46:  uint16(190),
		47:  uint16(190),
		48:  uint16(190),
		49:  uint16(190),
		50:  uint16(190),
		51:  uint16(190),
		52:  uint16(190),
		53:  uint16(190),
		54:  uint16(190),
		55:  uint16(190),
		56:  uint16(190),
		57:  uint16(190),
		58:  uint16(192),
		60:  uint16(190),
		61:  uint16(192),
		63:  uint16(190),
		64:  uint16(192),
		65:  uint16(192),
		66:  uint16(190),
		67:  uint16(190),
		68:  uint16(190),
		70:  uint16(190),
		72:  uint16(190),
		73:  uint16(190),
		74:  uint16(190),
		75:  uint16(190),
		76:  uint16(190),
		132: uint16(3),
		134: uint16(190),
	},
	18: {
		0:   uint16(194),
		1:   uint16(194),
		2:   uint16(194),
		3:   uint16(194),
		4:   uint16(194),
		6:   uint16(194),
		7:   uint16(194),
		8:   uint16(194),
		9:   uint16(194),
		10:  uint16(194),
		11:  uint16(194),
		12:  uint16(194),
		13:  uint16(194),
		14:  uint16(194),
		15:  uint16(194),
		16:  uint16(194),
		17:  uint16(194),
		18:  uint16(196),
		19:  uint16(194),
		20:  uint16(194),
		21:  uint16(194),
		22:  uint16(194),
		23:  uint16(194),
		24:  uint16(194),
		25:  uint16(194),
		26:  uint16(194),
		27:  uint16(194),
		28:  uint16(194),
		29:  uint16(194),
		30:  uint16(194),
		31:  uint16(194),
		32:  uint16(194),
		33:  uint16(194),
		34:  uint16(194),
		35:  uint16(194),
		36:  uint16(194),
		37:  uint16(194),
		38:  uint16(194),
		39:  uint16(194),
		40:  uint16(194),
		41:  uint16(194),
		42:  uint16(194),
		43:  uint16(194),
		44:  uint16(194),
		45:  uint16(194),
		46:  uint16(194),
		47:  uint16(194),
		48:  uint16(194),
		49:  uint16(194),
		50:  uint16(194),
		51:  uint16(194),
		52:  uint16(194),
		53:  uint16(194),
		54:  uint16(194),
		55:  uint16(194),
		56:  uint16(194),
		57:  uint16(194),
		58:  uint16(196),
		60:  uint16(194),
		61:  uint16(196),
		63:  uint16(194),
		64:  uint16(196),
		65:  uint16(196),
		66:  uint16(194),
		67:  uint16(194),
		68:  uint16(194),
		70:  uint16(194),
		72:  uint16(194),
		73:  uint16(194),
		74:  uint16(194),
		75:  uint16(194),
		76:  uint16(194),
		132: uint16(3),
		134: uint16(194),
	},
	19: {
		0:   uint16(198),
		1:   uint16(198),
		2:   uint16(198),
		3:   uint16(198),
		4:   uint16(198),
		6:   uint16(198),
		7:   uint16(198),
		8:   uint16(198),
		9:   uint16(198),
		10:  uint16(198),
		11:  uint16(198),
		12:  uint16(198),
		13:  uint16(198),
		14:  uint16(198),
		15:  uint16(198),
		16:  uint16(198),
		17:  uint16(198),
		18:  uint16(200),
		19:  uint16(198),
		20:  uint16(198),
		21:  uint16(198),
		22:  uint16(198),
		23:  uint16(198),
		24:  uint16(198),
		25:  uint16(198),
		26:  uint16(198),
		27:  uint16(198),
		28:  uint16(198),
		29:  uint16(198),
		30:  uint16(198),
		31:  uint16(198),
		32:  uint16(198),
		33:  uint16(198),
		34:  uint16(198),
		35:  uint16(198),
		36:  uint16(198),
		37:  uint16(198),
		38:  uint16(198),
		39:  uint16(198),
		40:  uint16(198),
		41:  uint16(198),
		42:  uint16(198),
		43:  uint16(198),
		44:  uint16(198),
		45:  uint16(198),
		46:  uint16(198),
		47:  uint16(198),
		48:  uint16(198),
		49:  uint16(198),
		50:  uint16(198),
		51:  uint16(198),
		52:  uint16(198),
		53:  uint16(198),
		54:  uint16(198),
		55:  uint16(198),
		56:  uint16(198),
		57:  uint16(198),
		58:  uint16(200),
		60:  uint16(198),
		61:  uint16(200),
		63:  uint16(198),
		64:  uint16(200),
		65:  uint16(200),
		66:  uint16(198),
		67:  uint16(198),
		68:  uint16(198),
		70:  uint16(198),
		72:  uint16(198),
		73:  uint16(198),
		74:  uint16(198),
		75:  uint16(198),
		76:  uint16(198),
		132: uint16(3),
		134: uint16(198),
	},
	20: {
		0:   uint16(202),
		1:   uint16(202),
		2:   uint16(202),
		3:   uint16(202),
		4:   uint16(202),
		6:   uint16(202),
		7:   uint16(202),
		8:   uint16(202),
		9:   uint16(202),
		10:  uint16(202),
		11:  uint16(202),
		12:  uint16(202),
		13:  uint16(202),
		14:  uint16(202),
		15:  uint16(202),
		16:  uint16(202),
		17:  uint16(202),
		18:  uint16(204),
		19:  uint16(202),
		20:  uint16(202),
		21:  uint16(202),
		22:  uint16(202),
		23:  uint16(202),
		24:  uint16(202),
		25:  uint16(202),
		26:  uint16(202),
		27:  uint16(202),
		28:  uint16(202),
		29:  uint16(202),
		30:  uint16(202),
		31:  uint16(202),
		32:  uint16(202),
		33:  uint16(202),
		34:  uint16(202),
		35:  uint16(202),
		36:  uint16(202),
		37:  uint16(202),
		38:  uint16(202),
		39:  uint16(202),
		40:  uint16(202),
		41:  uint16(202),
		42:  uint16(202),
		43:  uint16(202),
		44:  uint16(202),
		45:  uint16(202),
		46:  uint16(202),
		47:  uint16(202),
		48:  uint16(202),
		49:  uint16(202),
		50:  uint16(202),
		51:  uint16(202),
		52:  uint16(202),
		53:  uint16(202),
		54:  uint16(202),
		55:  uint16(202),
		56:  uint16(202),
		57:  uint16(202),
		58:  uint16(204),
		60:  uint16(202),
		61:  uint16(204),
		63:  uint16(202),
		64:  uint16(204),
		65:  uint16(204),
		66:  uint16(202),
		67:  uint16(202),
		68:  uint16(202),
		70:  uint16(202),
		72:  uint16(202),
		73:  uint16(202),
		74:  uint16(202),
		75:  uint16(202),
		76:  uint16(202),
		132: uint16(3),
		134: uint16(202),
	},
	21: {
		0:   uint16(206),
		1:   uint16(206),
		2:   uint16(206),
		3:   uint16(206),
		4:   uint16(206),
		6:   uint16(206),
		7:   uint16(206),
		8:   uint16(206),
		9:   uint16(206),
		10:  uint16(206),
		11:  uint16(206),
		12:  uint16(206),
		13:  uint16(206),
		14:  uint16(206),
		15:  uint16(206),
		16:  uint16(206),
		17:  uint16(206),
		18:  uint16(208),
		19:  uint16(206),
		20:  uint16(206),
		21:  uint16(206),
		22:  uint16(206),
		23:  uint16(206),
		24:  uint16(206),
		25:  uint16(206),
		26:  uint16(206),
		27:  uint16(206),
		28:  uint16(206),
		29:  uint16(206),
		30:  uint16(206),
		31:  uint16(206),
		32:  uint16(206),
		33:  uint16(206),
		34:  uint16(206),
		35:  uint16(206),
		36:  uint16(206),
		37:  uint16(206),
		38:  uint16(206),
		39:  uint16(206),
		40:  uint16(206),
		41:  uint16(206),
		42:  uint16(206),
		43:  uint16(206),
		44:  uint16(206),
		45:  uint16(206),
		46:  uint16(206),
		47:  uint16(206),
		48:  uint16(206),
		49:  uint16(206),
		50:  uint16(206),
		51:  uint16(206),
		52:  uint16(206),
		53:  uint16(206),
		54:  uint16(206),
		55:  uint16(206),
		56:  uint16(206),
		57:  uint16(206),
		58:  uint16(208),
		60:  uint16(206),
		61:  uint16(208),
		63:  uint16(206),
		64:  uint16(208),
		65:  uint16(208),
		66:  uint16(206),
		67:  uint16(206),
		68:  uint16(206),
		70:  uint16(206),
		72:  uint16(206),
		73:  uint16(206),
		74:  uint16(206),
		75:  uint16(206),
		76:  uint16(206),
		132: uint16(3),
		134: uint16(206),
	},
	22: {
		0:   uint16(210),
		1:   uint16(210),
		2:   uint16(210),
		3:   uint16(210),
		4:   uint16(210),
		6:   uint16(210),
		7:   uint16(210),
		8:   uint16(210),
		9:   uint16(210),
		10:  uint16(210),
		11:  uint16(210),
		12:  uint16(210),
		13:  uint16(210),
		14:  uint16(210),
		15:  uint16(210),
		16:  uint16(210),
		17:  uint16(210),
		18:  uint16(212),
		19:  uint16(210),
		20:  uint16(210),
		21:  uint16(210),
		22:  uint16(210),
		23:  uint16(210),
		24:  uint16(210),
		25:  uint16(210),
		26:  uint16(210),
		27:  uint16(210),
		28:  uint16(210),
		29:  uint16(210),
		30:  uint16(210),
		31:  uint16(210),
		32:  uint16(210),
		33:  uint16(210),
		34:  uint16(210),
		35:  uint16(210),
		36:  uint16(210),
		37:  uint16(210),
		38:  uint16(210),
		39:  uint16(210),
		40:  uint16(210),
		41:  uint16(210),
		42:  uint16(210),
		43:  uint16(210),
		44:  uint16(210),
		45:  uint16(210),
		46:  uint16(210),
		47:  uint16(210),
		48:  uint16(210),
		49:  uint16(210),
		50:  uint16(210),
		51:  uint16(210),
		52:  uint16(210),
		53:  uint16(210),
		54:  uint16(210),
		55:  uint16(210),
		56:  uint16(210),
		57:  uint16(210),
		58:  uint16(212),
		60:  uint16(210),
		61:  uint16(212),
		63:  uint16(210),
		64:  uint16(212),
		65:  uint16(212),
		66:  uint16(210),
		67:  uint16(210),
		68:  uint16(210),
		70:  uint16(210),
		72:  uint16(210),
		73:  uint16(210),
		74:  uint16(210),
		75:  uint16(210),
		76:  uint16(210),
		132: uint16(3),
		134: uint16(210),
	},
	23: {
		0:   uint16(214),
		1:   uint16(214),
		2:   uint16(214),
		3:   uint16(214),
		4:   uint16(214),
		6:   uint16(214),
		7:   uint16(214),
		8:   uint16(214),
		9:   uint16(214),
		10:  uint16(214),
		11:  uint16(214),
		12:  uint16(214),
		13:  uint16(214),
		14:  uint16(214),
		15:  uint16(214),
		16:  uint16(214),
		17:  uint16(214),
		18:  uint16(216),
		19:  uint16(214),
		20:  uint16(214),
		21:  uint16(214),
		22:  uint16(214),
		23:  uint16(214),
		24:  uint16(214),
		25:  uint16(214),
		26:  uint16(214),
		27:  uint16(214),
		28:  uint16(214),
		29:  uint16(214),
		30:  uint16(214),
		31:  uint16(214),
		32:  uint16(214),
		33:  uint16(214),
		34:  uint16(214),
		35:  uint16(214),
		36:  uint16(214),
		37:  uint16(214),
		38:  uint16(214),
		39:  uint16(214),
		40:  uint16(214),
		41:  uint16(214),
		42:  uint16(214),
		43:  uint16(214),
		44:  uint16(214),
		45:  uint16(214),
		46:  uint16(214),
		47:  uint16(214),
		48:  uint16(214),
		49:  uint16(214),
		50:  uint16(214),
		51:  uint16(214),
		52:  uint16(214),
		53:  uint16(214),
		54:  uint16(214),
		55:  uint16(214),
		56:  uint16(214),
		57:  uint16(214),
		58:  uint16(216),
		60:  uint16(214),
		61:  uint16(216),
		63:  uint16(214),
		64:  uint16(216),
		65:  uint16(216),
		66:  uint16(214),
		67:  uint16(214),
		68:  uint16(214),
		70:  uint16(214),
		72:  uint16(214),
		73:  uint16(214),
		74:  uint16(214),
		75:  uint16(214),
		76:  uint16(214),
		132: uint16(3),
		134: uint16(214),
	},
	24: {
		0:   uint16(218),
		1:   uint16(218),
		2:   uint16(218),
		3:   uint16(218),
		4:   uint16(218),
		6:   uint16(218),
		7:   uint16(218),
		8:   uint16(218),
		9:   uint16(218),
		10:  uint16(218),
		11:  uint16(218),
		12:  uint16(218),
		13:  uint16(218),
		14:  uint16(218),
		15:  uint16(218),
		16:  uint16(218),
		17:  uint16(218),
		18:  uint16(220),
		19:  uint16(218),
		20:  uint16(218),
		21:  uint16(218),
		22:  uint16(218),
		23:  uint16(218),
		24:  uint16(218),
		25:  uint16(218),
		26:  uint16(218),
		27:  uint16(218),
		28:  uint16(218),
		29:  uint16(218),
		30:  uint16(218),
		31:  uint16(218),
		32:  uint16(218),
		33:  uint16(218),
		34:  uint16(218),
		35:  uint16(218),
		36:  uint16(218),
		37:  uint16(218),
		38:  uint16(218),
		39:  uint16(218),
		40:  uint16(218),
		41:  uint16(218),
		42:  uint16(218),
		43:  uint16(218),
		44:  uint16(218),
		45:  uint16(218),
		46:  uint16(218),
		47:  uint16(218),
		48:  uint16(218),
		49:  uint16(218),
		50:  uint16(218),
		51:  uint16(218),
		52:  uint16(218),
		53:  uint16(218),
		54:  uint16(218),
		55:  uint16(218),
		56:  uint16(218),
		57:  uint16(218),
		58:  uint16(220),
		60:  uint16(218),
		61:  uint16(220),
		63:  uint16(218),
		64:  uint16(220),
		65:  uint16(220),
		66:  uint16(218),
		67:  uint16(218),
		68:  uint16(218),
		70:  uint16(218),
		72:  uint16(218),
		73:  uint16(218),
		74:  uint16(218),
		75:  uint16(218),
		76:  uint16(218),
		132: uint16(3),
		134: uint16(218),
	},
	25: {
		0:   uint16(222),
		1:   uint16(222),
		2:   uint16(222),
		3:   uint16(222),
		4:   uint16(222),
		6:   uint16(222),
		7:   uint16(222),
		8:   uint16(222),
		9:   uint16(222),
		10:  uint16(222),
		11:  uint16(222),
		12:  uint16(222),
		13:  uint16(222),
		14:  uint16(222),
		15:  uint16(222),
		16:  uint16(222),
		17:  uint16(222),
		18:  uint16(224),
		19:  uint16(222),
		20:  uint16(222),
		21:  uint16(222),
		22:  uint16(222),
		23:  uint16(222),
		24:  uint16(222),
		25:  uint16(222),
		26:  uint16(222),
		27:  uint16(222),
		28:  uint16(222),
		29:  uint16(222),
		30:  uint16(222),
		31:  uint16(222),
		32:  uint16(222),
		33:  uint16(222),
		34:  uint16(222),
		35:  uint16(222),
		36:  uint16(222),
		37:  uint16(222),
		38:  uint16(222),
		39:  uint16(222),
		40:  uint16(222),
		41:  uint16(222),
		42:  uint16(222),
		43:  uint16(222),
		44:  uint16(222),
		45:  uint16(222),
		46:  uint16(222),
		47:  uint16(222),
		48:  uint16(222),
		49:  uint16(222),
		50:  uint16(222),
		51:  uint16(222),
		52:  uint16(222),
		53:  uint16(222),
		54:  uint16(222),
		55:  uint16(222),
		56:  uint16(222),
		57:  uint16(222),
		58:  uint16(224),
		60:  uint16(222),
		61:  uint16(224),
		63:  uint16(222),
		64:  uint16(224),
		65:  uint16(224),
		66:  uint16(222),
		67:  uint16(222),
		68:  uint16(222),
		70:  uint16(222),
		72:  uint16(222),
		73:  uint16(222),
		74:  uint16(222),
		75:  uint16(222),
		76:  uint16(222),
		132: uint16(3),
		134: uint16(222),
	},
	26: {
		0:   uint16(226),
		1:   uint16(226),
		2:   uint16(226),
		3:   uint16(226),
		4:   uint16(226),
		6:   uint16(226),
		7:   uint16(226),
		8:   uint16(226),
		9:   uint16(226),
		10:  uint16(226),
		11:  uint16(226),
		12:  uint16(226),
		13:  uint16(226),
		14:  uint16(226),
		15:  uint16(226),
		16:  uint16(226),
		17:  uint16(226),
		18:  uint16(228),
		19:  uint16(226),
		20:  uint16(226),
		21:  uint16(226),
		22:  uint16(226),
		23:  uint16(226),
		24:  uint16(226),
		25:  uint16(226),
		26:  uint16(226),
		27:  uint16(226),
		28:  uint16(226),
		29:  uint16(226),
		30:  uint16(226),
		31:  uint16(226),
		32:  uint16(226),
		33:  uint16(226),
		34:  uint16(226),
		35:  uint16(226),
		36:  uint16(226),
		37:  uint16(226),
		38:  uint16(226),
		39:  uint16(226),
		40:  uint16(226),
		41:  uint16(226),
		42:  uint16(226),
		43:  uint16(226),
		44:  uint16(226),
		45:  uint16(226),
		46:  uint16(226),
		47:  uint16(226),
		48:  uint16(226),
		49:  uint16(226),
		50:  uint16(226),
		51:  uint16(226),
		52:  uint16(226),
		53:  uint16(226),
		54:  uint16(226),
		55:  uint16(226),
		56:  uint16(226),
		57:  uint16(226),
		58:  uint16(228),
		60:  uint16(226),
		61:  uint16(228),
		63:  uint16(226),
		64:  uint16(228),
		65:  uint16(228),
		66:  uint16(226),
		67:  uint16(226),
		68:  uint16(226),
		70:  uint16(226),
		72:  uint16(226),
		73:  uint16(226),
		74:  uint16(226),
		75:  uint16(226),
		76:  uint16(226),
		132: uint16(3),
		134: uint16(226),
	},
	27: {
		0:   uint16(230),
		1:   uint16(230),
		2:   uint16(230),
		3:   uint16(230),
		4:   uint16(230),
		6:   uint16(230),
		7:   uint16(230),
		8:   uint16(230),
		9:   uint16(230),
		10:  uint16(230),
		11:  uint16(230),
		12:  uint16(230),
		13:  uint16(230),
		14:  uint16(230),
		15:  uint16(230),
		16:  uint16(230),
		17:  uint16(230),
		18:  uint16(232),
		19:  uint16(230),
		20:  uint16(230),
		21:  uint16(230),
		22:  uint16(230),
		23:  uint16(230),
		24:  uint16(230),
		25:  uint16(230),
		26:  uint16(230),
		27:  uint16(230),
		28:  uint16(230),
		29:  uint16(230),
		30:  uint16(230),
		31:  uint16(230),
		32:  uint16(230),
		33:  uint16(230),
		34:  uint16(230),
		35:  uint16(230),
		36:  uint16(230),
		37:  uint16(230),
		38:  uint16(230),
		39:  uint16(230),
		40:  uint16(230),
		41:  uint16(230),
		42:  uint16(230),
		43:  uint16(230),
		44:  uint16(230),
		45:  uint16(230),
		46:  uint16(230),
		47:  uint16(230),
		48:  uint16(230),
		49:  uint16(230),
		50:  uint16(230),
		51:  uint16(230),
		52:  uint16(230),
		53:  uint16(230),
		54:  uint16(230),
		55:  uint16(230),
		56:  uint16(230),
		57:  uint16(230),
		58:  uint16(232),
		60:  uint16(230),
		61:  uint16(232),
		63:  uint16(230),
		64:  uint16(232),
		65:  uint16(232),
		66:  uint16(230),
		67:  uint16(230),
		68:  uint16(230),
		70:  uint16(230),
		72:  uint16(230),
		73:  uint16(230),
		74:  uint16(230),
		75:  uint16(230),
		76:  uint16(230),
		132: uint16(3),
		134: uint16(230),
	},
	28: {
		0:   uint16(234),
		1:   uint16(234),
		2:   uint16(234),
		3:   uint16(234),
		4:   uint16(234),
		6:   uint16(234),
		7:   uint16(234),
		8:   uint16(234),
		9:   uint16(234),
		10:  uint16(234),
		11:  uint16(234),
		12:  uint16(234),
		13:  uint16(234),
		14:  uint16(234),
		15:  uint16(234),
		16:  uint16(234),
		17:  uint16(234),
		18:  uint16(236),
		19:  uint16(234),
		20:  uint16(234),
		21:  uint16(234),
		22:  uint16(234),
		23:  uint16(234),
		24:  uint16(234),
		25:  uint16(234),
		26:  uint16(234),
		27:  uint16(234),
		28:  uint16(234),
		29:  uint16(234),
		30:  uint16(234),
		31:  uint16(234),
		32:  uint16(234),
		33:  uint16(234),
		34:  uint16(234),
		35:  uint16(234),
		36:  uint16(234),
		37:  uint16(234),
		38:  uint16(234),
		39:  uint16(234),
		40:  uint16(234),
		41:  uint16(234),
		42:  uint16(234),
		43:  uint16(234),
		44:  uint16(234),
		45:  uint16(234),
		46:  uint16(234),
		47:  uint16(234),
		48:  uint16(234),
		49:  uint16(234),
		50:  uint16(234),
		51:  uint16(234),
		52:  uint16(234),
		53:  uint16(234),
		54:  uint16(234),
		55:  uint16(234),
		56:  uint16(234),
		57:  uint16(234),
		58:  uint16(236),
		60:  uint16(234),
		61:  uint16(236),
		63:  uint16(234),
		64:  uint16(236),
		65:  uint16(236),
		66:  uint16(234),
		67:  uint16(234),
		68:  uint16(234),
		70:  uint16(234),
		72:  uint16(234),
		73:  uint16(234),
		74:  uint16(234),
		75:  uint16(234),
		76:  uint16(234),
		132: uint16(3),
		134: uint16(234),
	},
	29: {
		0:   uint16(238),
		1:   uint16(238),
		2:   uint16(238),
		3:   uint16(238),
		4:   uint16(238),
		6:   uint16(238),
		7:   uint16(238),
		8:   uint16(238),
		9:   uint16(238),
		10:  uint16(238),
		11:  uint16(238),
		12:  uint16(238),
		13:  uint16(238),
		14:  uint16(238),
		15:  uint16(238),
		16:  uint16(238),
		17:  uint16(238),
		18:  uint16(240),
		19:  uint16(238),
		20:  uint16(238),
		21:  uint16(238),
		22:  uint16(238),
		23:  uint16(238),
		24:  uint16(238),
		25:  uint16(238),
		26:  uint16(238),
		27:  uint16(238),
		28:  uint16(238),
		29:  uint16(238),
		30:  uint16(238),
		31:  uint16(238),
		32:  uint16(238),
		33:  uint16(238),
		34:  uint16(238),
		35:  uint16(238),
		36:  uint16(238),
		37:  uint16(238),
		38:  uint16(238),
		39:  uint16(238),
		40:  uint16(238),
		41:  uint16(238),
		42:  uint16(238),
		43:  uint16(238),
		44:  uint16(238),
		45:  uint16(238),
		46:  uint16(238),
		47:  uint16(238),
		48:  uint16(238),
		49:  uint16(238),
		50:  uint16(238),
		51:  uint16(238),
		52:  uint16(238),
		53:  uint16(238),
		54:  uint16(238),
		55:  uint16(238),
		56:  uint16(238),
		57:  uint16(238),
		58:  uint16(240),
		60:  uint16(238),
		61:  uint16(240),
		63:  uint16(238),
		64:  uint16(240),
		65:  uint16(240),
		66:  uint16(238),
		67:  uint16(238),
		68:  uint16(238),
		70:  uint16(238),
		72:  uint16(238),
		73:  uint16(238),
		74:  uint16(238),
		75:  uint16(238),
		76:  uint16(238),
		132: uint16(3),
		134: uint16(238),
	},
	30: {
		0:   uint16(242),
		1:   uint16(242),
		2:   uint16(242),
		3:   uint16(242),
		4:   uint16(242),
		6:   uint16(242),
		7:   uint16(242),
		8:   uint16(242),
		9:   uint16(242),
		10:  uint16(242),
		11:  uint16(242),
		12:  uint16(242),
		13:  uint16(242),
		14:  uint16(242),
		15:  uint16(242),
		16:  uint16(242),
		17:  uint16(242),
		18:  uint16(244),
		19:  uint16(242),
		20:  uint16(242),
		21:  uint16(242),
		22:  uint16(242),
		23:  uint16(242),
		24:  uint16(242),
		25:  uint16(242),
		26:  uint16(242),
		27:  uint16(242),
		28:  uint16(242),
		29:  uint16(242),
		30:  uint16(242),
		31:  uint16(242),
		32:  uint16(242),
		33:  uint16(242),
		34:  uint16(242),
		35:  uint16(242),
		36:  uint16(242),
		37:  uint16(242),
		38:  uint16(242),
		39:  uint16(242),
		40:  uint16(242),
		41:  uint16(242),
		42:  uint16(242),
		43:  uint16(242),
		44:  uint16(242),
		45:  uint16(242),
		46:  uint16(242),
		47:  uint16(242),
		48:  uint16(242),
		49:  uint16(242),
		50:  uint16(242),
		51:  uint16(242),
		52:  uint16(242),
		53:  uint16(242),
		54:  uint16(242),
		55:  uint16(242),
		56:  uint16(242),
		57:  uint16(242),
		58:  uint16(244),
		60:  uint16(242),
		61:  uint16(244),
		63:  uint16(242),
		64:  uint16(244),
		65:  uint16(244),
		66:  uint16(242),
		67:  uint16(242),
		68:  uint16(242),
		70:  uint16(242),
		72:  uint16(242),
		73:  uint16(242),
		74:  uint16(242),
		75:  uint16(242),
		76:  uint16(242),
		132: uint16(3),
		134: uint16(242),
	},
	31: {
		0:   uint16(246),
		1:   uint16(246),
		2:   uint16(246),
		3:   uint16(246),
		4:   uint16(246),
		6:   uint16(246),
		7:   uint16(246),
		8:   uint16(246),
		9:   uint16(246),
		10:  uint16(246),
		11:  uint16(246),
		12:  uint16(246),
		13:  uint16(246),
		14:  uint16(246),
		15:  uint16(246),
		16:  uint16(246),
		17:  uint16(246),
		18:  uint16(248),
		19:  uint16(246),
		20:  uint16(246),
		21:  uint16(246),
		22:  uint16(246),
		23:  uint16(246),
		24:  uint16(246),
		25:  uint16(246),
		26:  uint16(246),
		27:  uint16(246),
		28:  uint16(246),
		29:  uint16(246),
		30:  uint16(246),
		31:  uint16(246),
		32:  uint16(246),
		33:  uint16(246),
		34:  uint16(246),
		35:  uint16(246),
		36:  uint16(246),
		37:  uint16(246),
		38:  uint16(246),
		39:  uint16(246),
		40:  uint16(246),
		41:  uint16(246),
		42:  uint16(246),
		43:  uint16(246),
		44:  uint16(246),
		45:  uint16(246),
		46:  uint16(246),
		47:  uint16(246),
		48:  uint16(246),
		49:  uint16(246),
		50:  uint16(246),
		51:  uint16(246),
		52:  uint16(246),
		53:  uint16(246),
		54:  uint16(246),
		55:  uint16(246),
		56:  uint16(246),
		57:  uint16(246),
		58:  uint16(248),
		60:  uint16(246),
		61:  uint16(248),
		63:  uint16(246),
		64:  uint16(248),
		65:  uint16(248),
		66:  uint16(246),
		67:  uint16(246),
		68:  uint16(246),
		70:  uint16(246),
		72:  uint16(246),
		73:  uint16(246),
		74:  uint16(246),
		75:  uint16(246),
		76:  uint16(246),
		132: uint16(3),
		134: uint16(246),
	},
	32: {
		6:   uint16(15),
		7:   uint16(15),
		8:   uint16(17),
		9:   uint16(19),
		10:  uint16(19),
		11:  uint16(19),
		12:  uint16(19),
		13:  uint16(19),
		14:  uint16(21),
		15:  uint16(19),
		16:  uint16(15),
		17:  uint16(23),
		18:  uint16(25),
		19:  uint16(19),
		20:  uint16(19),
		21:  uint16(15),
		22:  uint16(15),
		23:  uint16(27),
		24:  uint16(19),
		25:  uint16(15),
		26:  uint16(19),
		27:  uint16(29),
		28:  uint16(15),
		29:  uint16(19),
		30:  uint16(27),
		31:  uint16(31),
		32:  uint16(15),
		33:  uint16(33),
		34:  uint16(19),
		35:  uint16(19),
		36:  uint16(15),
		37:  uint16(19),
		38:  uint16(19),
		39:  uint16(19),
		40:  uint16(19),
		41:  uint16(15),
		42:  uint16(19),
		43:  uint16(35),
		44:  uint16(15),
		45:  uint16(15),
		46:  uint16(15),
		47:  uint16(15),
		48:  uint16(37),
		49:  uint16(37),
		50:  uint16(19),
		51:  uint16(15),
		52:  uint16(39),
		53:  uint16(15),
		54:  uint16(19),
		55:  uint16(19),
		56:  uint16(19),
		57:  uint16(15),
		58:  uint16(41),
		60:  uint16(43),
		61:  uint16(45),
		63:  uint16(47),
		64:  uint16(49),
		65:  uint16(49),
		66:  uint16(51),
		67:  uint16(51),
		68:  uint16(53),
		70:  uint16(55),
		72:  uint16(57),
		73:  uint16(57),
		74:  uint16(59),
		75:  uint16(57),
		76:  uint16(61),
		132: uint16(3),
		134: uint16(250),
		138: uint16(2),
		139: uint16(2),
		140: uint16(2),
		158: uint16(2),
	},
}

var ts_small_parse_table = [1747]uint16_t{
	0:    uint16(3),
	1:    uint16(3),
	2:    uint16(1),
	3:    uint16(sym_comment),
	4:    uint16(179),
	5:    uint16(1),
	6:    uint16(sym_shape),
	7:    uint16(252),
	8:    uint16(12),
	9:    uint16(anon_sym_Circle),
	10:   uint16(anon_sym_Diamond),
	11:   uint16(anon_sym_Hexagon),
	12:   uint16(anon_sym_Square),
	13:   uint16(anon_sym_Star),
	14:   uint16(anon_sym_Triangle),
	15:   uint16(anon_sym_Cross),
	16:   uint16(anon_sym_Moon),
	17:   uint16(anon_sym_Raindrop),
	18:   uint16(anon_sym_Kite),
	19:   uint16(anon_sym_Pentagon),
	20:   uint16(anon_sym_UpsideDownHouse),
	21:   uint16(3),
	22:   uint16(3),
	23:   uint16(1),
	24:   uint16(sym_comment),
	25:   uint16(180),
	26:   uint16(1),
	27:   uint16(sym_shape),
	28:   uint16(252),
	29:   uint16(12),
	30:   uint16(anon_sym_Circle),
	31:   uint16(anon_sym_Diamond),
	32:   uint16(anon_sym_Hexagon),
	33:   uint16(anon_sym_Square),
	34:   uint16(anon_sym_Star),
	35:   uint16(anon_sym_Triangle),
	36:   uint16(anon_sym_Cross),
	37:   uint16(anon_sym_Moon),
	38:   uint16(anon_sym_Raindrop),
	39:   uint16(anon_sym_Kite),
	40:   uint16(anon_sym_Pentagon),
	41:   uint16(anon_sym_UpsideDownHouse),
	42:   uint16(4),
	43:   uint16(3),
	44:   uint16(1),
	45:   uint16(sym_comment),
	46:   uint16(254),
	47:   uint16(1),
	48:   uint16(anon_sym_None),
	49:   uint16(104),
	50:   uint16(1),
	51:   uint16(sym_colour),
	52:   uint16(256),
	53:   uint16(11),
	54:   uint16(anon_sym_Red),
	55:   uint16(anon_sym_Green),
	56:   uint16(anon_sym_Blue),
	57:   uint16(anon_sym_Brown),
	58:   uint16(anon_sym_White),
	59:   uint16(anon_sym_Yellow),
	60:   uint16(anon_sym_Cyan),
	61:   uint16(anon_sym_Grey),
	62:   uint16(anon_sym_Orange),
	63:   uint16(anon_sym_Pink),
	64:   uint16(anon_sym_Purple),
	65:   uint16(3),
	66:   uint16(3),
	67:   uint16(1),
	68:   uint16(sym_comment),
	69:   uint16(166),
	70:   uint16(1),
	71:   uint16(sym_colour),
	72:   uint16(256),
	73:   uint16(11),
	74:   uint16(anon_sym_Red),
	75:   uint16(anon_sym_Green),
	76:   uint16(anon_sym_Blue),
	77:   uint16(anon_sym_Brown),
	78:   uint16(anon_sym_White),
	79:   uint16(anon_sym_Yellow),
	80:   uint16(anon_sym_Cyan),
	81:   uint16(anon_sym_Grey),
	82:   uint16(anon_sym_Orange),
	83:   uint16(anon_sym_Pink),
	84:   uint16(anon_sym_Purple),
	85:   uint16(3),
	86:   uint16(3),
	87:   uint16(1),
	88:   uint16(sym_comment),
	89:   uint16(194),
	90:   uint16(1),
	91:   uint16(sym_colour),
	92:   uint16(256),
	93:   uint16(11),
	94:   uint16(anon_sym_Red),
	95:   uint16(anon_sym_Green),
	96:   uint16(anon_sym_Blue),
	97:   uint16(anon_sym_Brown),
	98:   uint16(anon_sym_White),
	99:   uint16(anon_sym_Yellow),
	100:  uint16(anon_sym_Cyan),
	101:  uint16(anon_sym_Grey),
	102:  uint16(anon_sym_Orange),
	103:  uint16(anon_sym_Pink),
	104:  uint16(anon_sym_Purple),
	105:  uint16(6),
	106:  uint16(3),
	107:  uint16(1),
	108:  uint16(sym_comment),
	109:  uint16(260),
	110:  uint16(1),
	111:  uint16(anon_sym_DQUOTE),
	112:  uint16(95),
	113:  uint16(1),
	114:  uint16(sym_rarity),
	115:  uint16(258),
	116:  uint16(2),
	117:  uint16(aux_sym__equal_operator_token1),
	118:  uint16(aux_sym__range_operator_token1),
	119:  uint16(94),
	120:  uint16(2),
	121:  uint16(sym__equal_operator),
	122:  uint16(sym__range_operator),
	123:  uint16(262),
	124:  uint16(4),
	125:  uint16(anon_sym_Normal),
	126:  uint16(anon_sym_Magic),
	127:  uint16(anon_sym_Rare),
	128:  uint16(anon_sym_Unique),
	129:  uint16(4),
	130:  uint16(3),
	131:  uint16(1),
	132:  uint16(sym_comment),
	133:  uint16(266),
	134:  uint16(1),
	135:  uint16(anon_sym_DQUOTE),
	136:  uint16(91),
	137:  uint16(1),
	138:  uint16(sym_influence),
	139:  uint16(264),
	140:  uint16(7),
	141:  uint16(anon_sym_None),
	142:  uint16(anon_sym_Shaper),
	143:  uint16(anon_sym_Elder),
	144:  uint16(anon_sym_Crusader),
	145:  uint16(anon_sym_Hunter),
	146:  uint16(anon_sym_Redeemer),
	147:  uint16(anon_sym_Warlord),
	148:  uint16(6),
	149:  uint16(3),
	150:  uint16(1),
	151:  uint16(sym_comment),
	152:  uint16(9),
	153:  uint16(1),
	154:  uint16(anon_sym_Import),
	155:  uint16(268),
	156:  uint16(1),
	158:  uint16(270),
	159:  uint16(1),
	160:  uint16(sym__eol),
	161:  uint16(7),
	162:  uint16(3),
	163:  uint16(anon_sym_Show),
	164:  uint16(anon_sym_Hide),
	165:  uint16(anon_sym_Minimal),
	166:  uint16(42),
	167:  uint16(3),
	168:  uint16(sym_block),
	169:  uint16(sym_import),
	170:  uint16(aux_sym_filter_repeat1),
	171:  uint16(8),
	172:  uint16(3),
	173:  uint16(1),
	174:  uint16(sym_comment),
	175:  uint16(274),
	176:  uint16(1),
	177:  uint16(anon_sym_DQUOTE),
	178:  uint16(276),
	179:  uint16(1),
	180:  uint16(aux_sym_string_token2),
	181:  uint16(278),
	182:  uint16(1),
	183:  uint16(sym_number),
	184:  uint16(115),
	185:  uint16(1),
	186:  uint16(sym__quantity),
	187:  uint16(117),
	188:  uint16(1),
	189:  uint16(sym_string),
	190:  uint16(272),
	191:  uint16(2),
	192:  uint16(aux_sym__equal_operator_token1),
	193:  uint16(aux_sym__range_operator_token1),
	194:  uint16(191),
	195:  uint16(2),
	196:  uint16(sym__equal_operator),
	197:  uint16(sym__range_operator),
	198:  uint16(6),
	199:  uint16(3),
	200:  uint16(1),
	201:  uint16(sym_comment),
	202:  uint16(280),
	203:  uint16(1),
	205:  uint16(285),
	206:  uint16(1),
	207:  uint16(anon_sym_Import),
	208:  uint16(288),
	209:  uint16(1),
	210:  uint16(sym__eol),
	211:  uint16(282),
	212:  uint16(3),
	213:  uint16(anon_sym_Show),
	214:  uint16(anon_sym_Hide),
	215:  uint16(anon_sym_Minimal),
	216:  uint16(42),
	217:  uint16(3),
	218:  uint16(sym_block),
	219:  uint16(sym_import),
	220:  uint16(aux_sym_filter_repeat1),
	221:  uint16(6),
	222:  uint16(3),
	223:  uint16(1),
	224:  uint16(sym_comment),
	225:  uint16(293),
	226:  uint16(1),
	227:  uint16(anon_sym_DQUOTE),
	228:  uint16(295),
	229:  uint16(1),
	230:  uint16(aux_sym_sockets_token2),
	231:  uint16(98),
	232:  uint16(1),
	233:  uint16(sym_sockets),
	234:  uint16(291),
	235:  uint16(2),
	236:  uint16(aux_sym__equal_operator_token1),
	237:  uint16(aux_sym__range_operator_token1),
	238:  uint16(97),
	239:  uint16(2),
	240:  uint16(sym__equal_operator),
	241:  uint16(sym__range_operator),
	242:  uint16(4),
	243:  uint16(297),
	244:  uint16(1),
	246:  uint16(301),
	247:  uint16(1),
	248:  uint16(sym_comment),
	249:  uint16(303),
	250:  uint16(1),
	251:  uint16(sym__space),
	252:  uint16(299),
	253:  uint16(5),
	254:  uint16(anon_sym_Show),
	255:  uint16(anon_sym_Hide),
	256:  uint16(anon_sym_Minimal),
	257:  uint16(anon_sym_Import),
	258:  uint16(sym__eol),
	259:  uint16(3),
	260:  uint16(301),
	261:  uint16(1),
	262:  uint16(sym_comment),
	263:  uint16(305),
	264:  uint16(2),
	266:  uint16(sym__space),
	267:  uint16(307),
	268:  uint16(5),
	269:  uint16(anon_sym_Show),
	270:  uint16(anon_sym_Hide),
	271:  uint16(anon_sym_Minimal),
	272:  uint16(anon_sym_Import),
	273:  uint16(sym__eol),
	274:  uint16(5),
	275:  uint16(3),
	276:  uint16(1),
	277:  uint16(sym_comment),
	278:  uint16(311),
	279:  uint16(1),
	280:  uint16(sym_number),
	281:  uint16(137),
	282:  uint16(1),
	283:  uint16(sym__quantity),
	284:  uint16(309),
	285:  uint16(2),
	286:  uint16(aux_sym__equal_operator_token1),
	287:  uint16(aux_sym__range_operator_token1),
	288:  uint16(160),
	289:  uint16(2),
	290:  uint16(sym__equal_operator),
	291:  uint16(sym__range_operator),
	292:  uint16(4),
	293:  uint16(3),
	294:  uint16(1),
	295:  uint16(sym_comment),
	296:  uint16(260),
	297:  uint16(1),
	298:  uint16(anon_sym_DQUOTE),
	299:  uint16(95),
	300:  uint16(1),
	301:  uint16(sym_rarity),
	302:  uint16(262),
	303:  uint16(4),
	304:  uint16(anon_sym_Normal),
	305:  uint16(anon_sym_Magic),
	306:  uint16(anon_sym_Rare),
	307:  uint16(anon_sym_Unique),
	308:  uint16(2),
	309:  uint16(3),
	310:  uint16(1),
	311:  uint16(sym_comment),
	312:  uint16(313),
	313:  uint16(6),
	315:  uint16(anon_sym_Show),
	316:  uint16(anon_sym_Hide),
	317:  uint16(anon_sym_Minimal),
	318:  uint16(anon_sym_Import),
	319:  uint16(sym__eol),
	320:  uint16(4),
	321:  uint16(3),
	322:  uint16(1),
	323:  uint16(sym_comment),
	324:  uint16(315),
	325:  uint16(1),
	326:  uint16(anon_sym_DQUOTE),
	327:  uint16(89),
	328:  uint16(1),
	329:  uint16(sym_quality),
	330:  uint16(317),
	331:  uint16(4),
	332:  uint16(anon_sym_Superior),
	333:  uint16(anon_sym_Divergent),
	334:  uint16(anon_sym_Anomalous),
	335:  uint16(anon_sym_Phantasmal),
	336:  uint16(6),
	337:  uint16(3),
	338:  uint16(1),
	339:  uint16(sym_comment),
	340:  uint16(274),
	341:  uint16(1),
	342:  uint16(anon_sym_DQUOTE),
	343:  uint16(276),
	344:  uint16(1),
	345:  uint16(aux_sym_string_token2),
	346:  uint16(278),
	347:  uint16(1),
	348:  uint16(aux_sym__equal_operator_token1),
	349:  uint16(115),
	350:  uint16(1),
	351:  uint16(sym__equal_operator),
	352:  uint16(117),
	353:  uint16(1),
	354:  uint16(sym_string),
	355:  uint16(6),
	356:  uint16(3),
	357:  uint16(1),
	358:  uint16(sym_comment),
	359:  uint16(274),
	360:  uint16(1),
	361:  uint16(anon_sym_DQUOTE),
	362:  uint16(276),
	363:  uint16(1),
	364:  uint16(aux_sym_string_token2),
	365:  uint16(319),
	366:  uint16(1),
	367:  uint16(aux_sym__equal_operator_token1),
	368:  uint16(123),
	369:  uint16(1),
	370:  uint16(sym__equal_operator),
	371:  uint16(124),
	372:  uint16(1),
	373:  uint16(sym_string),
	374:  uint16(6),
	375:  uint16(3),
	376:  uint16(1),
	377:  uint16(sym_comment),
	378:  uint16(274),
	379:  uint16(1),
	380:  uint16(anon_sym_DQUOTE),
	381:  uint16(276),
	382:  uint16(1),
	383:  uint16(aux_sym_string_token2),
	384:  uint16(321),
	385:  uint16(1),
	386:  uint16(aux_sym__equal_operator_token1),
	387:  uint16(127),
	388:  uint16(1),
	389:  uint16(sym__equal_operator),
	390:  uint16(128),
	391:  uint16(1),
	392:  uint16(sym_string),
	393:  uint16(6),
	394:  uint16(3),
	395:  uint16(1),
	396:  uint16(sym_comment),
	397:  uint16(274),
	398:  uint16(1),
	399:  uint16(anon_sym_DQUOTE),
	400:  uint16(276),
	401:  uint16(1),
	402:  uint16(aux_sym_string_token2),
	403:  uint16(323),
	404:  uint16(1),
	405:  uint16(aux_sym__equal_operator_token1),
	406:  uint16(112),
	407:  uint16(1),
	408:  uint16(sym_string),
	409:  uint16(135),
	410:  uint16(1),
	411:  uint16(sym__equal_operator),
	412:  uint16(5),
	413:  uint16(3),
	414:  uint16(1),
	415:  uint16(sym_comment),
	416:  uint16(325),
	417:  uint16(1),
	418:  uint16(sym_boolean),
	419:  uint16(327),
	420:  uint16(1),
	421:  uint16(anon_sym_DQUOTE),
	422:  uint16(329),
	423:  uint16(1),
	424:  uint16(aux_sym_string_token2),
	425:  uint16(144),
	426:  uint16(1),
	427:  uint16(sym_string),
	428:  uint16(4),
	429:  uint16(301),
	430:  uint16(1),
	431:  uint16(sym_comment),
	432:  uint16(331),
	433:  uint16(1),
	434:  uint16(sym__space),
	435:  uint16(333),
	436:  uint16(1),
	437:  uint16(sym__eol),
	438:  uint16(78),
	439:  uint16(1),
	440:  uint16(aux_sym_condition_repeat2),
	441:  uint16(4),
	442:  uint16(301),
	443:  uint16(1),
	444:  uint16(sym_comment),
	445:  uint16(335),
	446:  uint16(1),
	447:  uint16(sym__space),
	448:  uint16(337),
	449:  uint16(1),
	450:  uint16(sym__eol),
	451:  uint16(71),
	452:  uint16(1),
	453:  uint16(aux_sym_condition_repeat6),
	454:  uint16(4),
	455:  uint16(301),
	456:  uint16(1),
	457:  uint16(sym_comment),
	458:  uint16(339),
	459:  uint16(1),
	460:  uint16(sym__space),
	461:  uint16(341),
	462:  uint16(1),
	463:  uint16(sym__eol),
	464:  uint16(75),
	465:  uint16(1),
	466:  uint16(aux_sym_condition_repeat1),
	467:  uint16(4),
	468:  uint16(301),
	469:  uint16(1),
	470:  uint16(sym_comment),
	471:  uint16(331),
	472:  uint16(1),
	473:  uint16(sym__space),
	474:  uint16(343),
	475:  uint16(1),
	476:  uint16(sym__eol),
	477:  uint16(78),
	478:  uint16(1),
	479:  uint16(aux_sym_condition_repeat2),
	480:  uint16(4),
	481:  uint16(301),
	482:  uint16(1),
	483:  uint16(sym_comment),
	484:  uint16(345),
	485:  uint16(1),
	486:  uint16(sym__space),
	487:  uint16(347),
	488:  uint16(1),
	489:  uint16(sym__eol),
	490:  uint16(81),
	491:  uint16(1),
	492:  uint16(aux_sym_condition_repeat3),
	493:  uint16(4),
	494:  uint16(301),
	495:  uint16(1),
	496:  uint16(sym_comment),
	497:  uint16(349),
	498:  uint16(1),
	499:  uint16(sym__space),
	500:  uint16(351),
	501:  uint16(1),
	502:  uint16(sym__eol),
	503:  uint16(84),
	504:  uint16(1),
	505:  uint16(aux_sym_condition_repeat4),
	506:  uint16(4),
	507:  uint16(301),
	508:  uint16(1),
	509:  uint16(sym_comment),
	510:  uint16(353),
	511:  uint16(1),
	512:  uint16(sym__space),
	513:  uint16(355),
	514:  uint16(1),
	515:  uint16(sym__eol),
	516:  uint16(74),
	517:  uint16(1),
	518:  uint16(aux_sym_condition_repeat7),
	519:  uint16(4),
	520:  uint16(301),
	521:  uint16(1),
	522:  uint16(sym_comment),
	523:  uint16(355),
	524:  uint16(1),
	525:  uint16(sym__eol),
	526:  uint16(357),
	527:  uint16(1),
	528:  uint16(sym__space),
	529:  uint16(68),
	530:  uint16(1),
	531:  uint16(aux_sym_condition_repeat8),
	532:  uint16(4),
	533:  uint16(301),
	534:  uint16(1),
	535:  uint16(sym_comment),
	536:  uint16(359),
	537:  uint16(1),
	538:  uint16(sym__space),
	539:  uint16(362),
	540:  uint16(1),
	541:  uint16(sym__eol),
	542:  uint16(63),
	543:  uint16(1),
	544:  uint16(aux_sym_condition_repeat5),
	545:  uint16(4),
	546:  uint16(301),
	547:  uint16(1),
	548:  uint16(sym_comment),
	549:  uint16(337),
	550:  uint16(1),
	551:  uint16(sym__eol),
	552:  uint16(357),
	553:  uint16(1),
	554:  uint16(sym__space),
	555:  uint16(68),
	556:  uint16(1),
	557:  uint16(aux_sym_condition_repeat8),
	558:  uint16(4),
	559:  uint16(3),
	560:  uint16(1),
	561:  uint16(sym_comment),
	562:  uint16(293),
	563:  uint16(1),
	564:  uint16(anon_sym_DQUOTE),
	565:  uint16(295),
	566:  uint16(1),
	567:  uint16(aux_sym_sockets_token2),
	568:  uint16(98),
	569:  uint16(1),
	570:  uint16(sym_sockets),
	571:  uint16(4),
	572:  uint16(3),
	573:  uint16(1),
	574:  uint16(sym_comment),
	575:  uint16(364),
	576:  uint16(1),
	577:  uint16(anon_sym_None),
	578:  uint16(366),
	579:  uint16(1),
	580:  uint16(aux_sym__id_token1),
	581:  uint16(100),
	582:  uint16(1),
	583:  uint16(sym__id),
	584:  uint16(4),
	585:  uint16(3),
	586:  uint16(1),
	587:  uint16(sym_comment),
	588:  uint16(364),
	589:  uint16(1),
	590:  uint16(aux_sym_action_token1),
	591:  uint16(368),
	592:  uint16(1),
	593:  uint16(anon_sym_DQUOTE),
	594:  uint16(101),
	595:  uint16(1),
	596:  uint16(sym_file),
	597:  uint16(4),
	598:  uint16(301),
	599:  uint16(1),
	600:  uint16(sym_comment),
	601:  uint16(370),
	602:  uint16(1),
	603:  uint16(sym__space),
	604:  uint16(373),
	605:  uint16(1),
	606:  uint16(sym__eol),
	607:  uint16(68),
	608:  uint16(1),
	609:  uint16(aux_sym_condition_repeat8),
	610:  uint16(4),
	611:  uint16(3),
	612:  uint16(1),
	613:  uint16(sym_comment),
	614:  uint16(375),
	615:  uint16(1),
	616:  uint16(anon_sym_DASH1),
	617:  uint16(377),
	618:  uint16(1),
	619:  uint16(aux_sym__icon_size_token1),
	620:  uint16(153),
	621:  uint16(1),
	622:  uint16(sym__icon_size),
	623:  uint16(4),
	624:  uint16(301),
	625:  uint16(1),
	626:  uint16(sym_comment),
	627:  uint16(345),
	628:  uint16(1),
	629:  uint16(sym__space),
	630:  uint16(379),
	631:  uint16(1),
	632:  uint16(sym__eol),
	633:  uint16(81),
	634:  uint16(1),
	635:  uint16(aux_sym_condition_repeat3),
	636:  uint16(4),
	637:  uint16(301),
	638:  uint16(1),
	639:  uint16(sym_comment),
	640:  uint16(381),
	641:  uint16(1),
	642:  uint16(sym__space),
	643:  uint16(384),
	644:  uint16(1),
	645:  uint16(sym__eol),
	646:  uint16(71),
	647:  uint16(1),
	648:  uint16(aux_sym_condition_repeat6),
	649:  uint16(4),
	650:  uint16(301),
	651:  uint16(1),
	652:  uint16(sym_comment),
	653:  uint16(349),
	654:  uint16(1),
	655:  uint16(sym__space),
	656:  uint16(386),
	657:  uint16(1),
	658:  uint16(sym__eol),
	659:  uint16(84),
	660:  uint16(1),
	661:  uint16(aux_sym_condition_repeat4),
	662:  uint16(4),
	663:  uint16(3),
	664:  uint16(1),
	665:  uint16(sym_comment),
	666:  uint16(274),
	667:  uint16(1),
	668:  uint16(anon_sym_DQUOTE),
	669:  uint16(276),
	670:  uint16(1),
	671:  uint16(aux_sym_string_token2),
	672:  uint16(117),
	673:  uint16(1),
	674:  uint16(sym_string),
	675:  uint16(4),
	676:  uint16(301),
	677:  uint16(1),
	678:  uint16(sym_comment),
	679:  uint16(388),
	680:  uint16(1),
	681:  uint16(sym__space),
	682:  uint16(391),
	683:  uint16(1),
	684:  uint16(sym__eol),
	685:  uint16(74),
	686:  uint16(1),
	687:  uint16(aux_sym_condition_repeat7),
	688:  uint16(4),
	689:  uint16(301),
	690:  uint16(1),
	691:  uint16(sym_comment),
	692:  uint16(393),
	693:  uint16(1),
	694:  uint16(sym__space),
	695:  uint16(396),
	696:  uint16(1),
	697:  uint16(sym__eol),
	698:  uint16(75),
	699:  uint16(1),
	700:  uint16(aux_sym_condition_repeat1),
	701:  uint16(4),
	702:  uint16(3),
	703:  uint16(1),
	704:  uint16(sym_comment),
	705:  uint16(274),
	706:  uint16(1),
	707:  uint16(anon_sym_DQUOTE),
	708:  uint16(276),
	709:  uint16(1),
	710:  uint16(aux_sym_string_token2),
	711:  uint16(124),
	712:  uint16(1),
	713:  uint16(sym_string),
	714:  uint16(4),
	715:  uint16(301),
	716:  uint16(1),
	717:  uint16(sym_comment),
	718:  uint16(337),
	719:  uint16(1),
	720:  uint16(sym__eol),
	721:  uint16(398),
	722:  uint16(1),
	723:  uint16(sym__space),
	724:  uint16(63),
	725:  uint16(1),
	726:  uint16(aux_sym_condition_repeat5),
	727:  uint16(4),
	728:  uint16(301),
	729:  uint16(1),
	730:  uint16(sym_comment),
	731:  uint16(400),
	732:  uint16(1),
	733:  uint16(sym__space),
	734:  uint16(403),
	735:  uint16(1),
	736:  uint16(sym__eol),
	737:  uint16(78),
	738:  uint16(1),
	739:  uint16(aux_sym_condition_repeat2),
	740:  uint16(4),
	741:  uint16(3),
	742:  uint16(1),
	743:  uint16(sym_comment),
	744:  uint16(274),
	745:  uint16(1),
	746:  uint16(anon_sym_DQUOTE),
	747:  uint16(276),
	748:  uint16(1),
	749:  uint16(aux_sym_string_token2),
	750:  uint16(128),
	751:  uint16(1),
	752:  uint16(sym_string),
	753:  uint16(4),
	754:  uint16(301),
	755:  uint16(1),
	756:  uint16(sym_comment),
	757:  uint16(339),
	758:  uint16(1),
	759:  uint16(sym__space),
	760:  uint16(405),
	761:  uint16(1),
	762:  uint16(sym__eol),
	763:  uint16(75),
	764:  uint16(1),
	765:  uint16(aux_sym_condition_repeat1),
	766:  uint16(4),
	767:  uint16(301),
	768:  uint16(1),
	769:  uint16(sym_comment),
	770:  uint16(407),
	771:  uint16(1),
	772:  uint16(sym__space),
	773:  uint16(410),
	774:  uint16(1),
	775:  uint16(sym__eol),
	776:  uint16(81),
	777:  uint16(1),
	778:  uint16(aux_sym_condition_repeat3),
	779:  uint16(4),
	780:  uint16(3),
	781:  uint16(1),
	782:  uint16(sym_comment),
	783:  uint16(274),
	784:  uint16(1),
	785:  uint16(anon_sym_DQUOTE),
	786:  uint16(276),
	787:  uint16(1),
	788:  uint16(aux_sym_string_token2),
	789:  uint16(112),
	790:  uint16(1),
	791:  uint16(sym_string),
	792:  uint16(4),
	793:  uint16(301),
	794:  uint16(1),
	795:  uint16(sym_comment),
	796:  uint16(339),
	797:  uint16(1),
	798:  uint16(sym__space),
	799:  uint16(412),
	800:  uint16(1),
	801:  uint16(sym__eol),
	802:  uint16(75),
	803:  uint16(1),
	804:  uint16(aux_sym_condition_repeat1),
	805:  uint16(4),
	806:  uint16(301),
	807:  uint16(1),
	808:  uint16(sym_comment),
	809:  uint16(414),
	810:  uint16(1),
	811:  uint16(sym__space),
	812:  uint16(417),
	813:  uint16(1),
	814:  uint16(sym__eol),
	815:  uint16(84),
	816:  uint16(1),
	817:  uint16(aux_sym_condition_repeat4),
	818:  uint16(4),
	819:  uint16(301),
	820:  uint16(1),
	821:  uint16(sym_comment),
	822:  uint16(337),
	823:  uint16(1),
	824:  uint16(sym__eol),
	825:  uint16(353),
	826:  uint16(1),
	827:  uint16(sym__space),
	828:  uint16(74),
	829:  uint16(1),
	830:  uint16(aux_sym_condition_repeat7),
	831:  uint16(3),
	832:  uint16(301),
	833:  uint16(1),
	834:  uint16(sym_comment),
	835:  uint16(419),
	836:  uint16(1),
	837:  uint16(sym__space),
	838:  uint16(55),
	839:  uint16(1),
	840:  uint16(aux_sym_condition_repeat2),
	841:  uint16(3),
	842:  uint16(301),
	843:  uint16(1),
	844:  uint16(sym_comment),
	845:  uint16(421),
	846:  uint16(1),
	847:  uint16(sym__space),
	848:  uint16(64),
	849:  uint16(1),
	850:  uint16(aux_sym_condition_repeat8),
	851:  uint16(3),
	852:  uint16(301),
	853:  uint16(1),
	854:  uint16(sym_comment),
	855:  uint16(423),
	856:  uint16(1),
	857:  uint16(sym__space),
	858:  uint16(425),
	859:  uint16(1),
	860:  uint16(sym__eol),
	861:  uint16(3),
	862:  uint16(301),
	863:  uint16(1),
	864:  uint16(sym_comment),
	865:  uint16(362),
	866:  uint16(1),
	867:  uint16(sym__eol),
	868:  uint16(427),
	869:  uint16(1),
	870:  uint16(sym__space),
	871:  uint16(3),
	872:  uint16(301),
	873:  uint16(1),
	874:  uint16(sym_comment),
	875:  uint16(429),
	876:  uint16(1),
	877:  uint16(sym__space),
	878:  uint16(431),
	879:  uint16(1),
	880:  uint16(sym__eol),
	881:  uint16(3),
	882:  uint16(301),
	883:  uint16(1),
	884:  uint16(sym_comment),
	885:  uint16(384),
	886:  uint16(1),
	887:  uint16(sym__eol),
	888:  uint16(433),
	889:  uint16(1),
	890:  uint16(sym__space),
	891:  uint16(3),
	892:  uint16(301),
	893:  uint16(1),
	894:  uint16(sym_comment),
	895:  uint16(435),
	896:  uint16(1),
	897:  uint16(sym__space),
	898:  uint16(437),
	899:  uint16(1),
	900:  uint16(sym__eol),
	901:  uint16(3),
	902:  uint16(301),
	903:  uint16(1),
	904:  uint16(sym_comment),
	905:  uint16(439),
	906:  uint16(1),
	907:  uint16(sym__space),
	908:  uint16(441),
	909:  uint16(1),
	910:  uint16(sym__eol),
	911:  uint16(3),
	912:  uint16(301),
	913:  uint16(1),
	914:  uint16(sym_comment),
	915:  uint16(353),
	916:  uint16(1),
	917:  uint16(sym__space),
	918:  uint16(61),
	919:  uint16(1),
	920:  uint16(aux_sym_condition_repeat7),
	921:  uint16(3),
	922:  uint16(301),
	923:  uint16(1),
	924:  uint16(sym_comment),
	925:  uint16(391),
	926:  uint16(1),
	927:  uint16(sym__eol),
	928:  uint16(443),
	929:  uint16(1),
	930:  uint16(sym__space),
	931:  uint16(3),
	932:  uint16(301),
	933:  uint16(1),
	934:  uint16(sym_comment),
	935:  uint16(445),
	936:  uint16(1),
	937:  uint16(sym__space),
	938:  uint16(447),
	939:  uint16(1),
	940:  uint16(sym__eol),
	941:  uint16(3),
	942:  uint16(301),
	943:  uint16(1),
	944:  uint16(sym_comment),
	945:  uint16(357),
	946:  uint16(1),
	947:  uint16(sym__space),
	948:  uint16(62),
	949:  uint16(1),
	950:  uint16(aux_sym_condition_repeat8),
	951:  uint16(3),
	952:  uint16(301),
	953:  uint16(1),
	954:  uint16(sym_comment),
	955:  uint16(373),
	956:  uint16(1),
	957:  uint16(sym__eol),
	958:  uint16(449),
	959:  uint16(1),
	960:  uint16(sym__space),
	961:  uint16(3),
	962:  uint16(301),
	963:  uint16(1),
	964:  uint16(sym_comment),
	965:  uint16(451),
	966:  uint16(1),
	967:  uint16(sym__space),
	968:  uint16(453),
	969:  uint16(1),
	970:  uint16(sym__eol),
	971:  uint16(3),
	972:  uint16(301),
	973:  uint16(1),
	974:  uint16(sym_comment),
	975:  uint16(455),
	976:  uint16(1),
	977:  uint16(sym__space),
	978:  uint16(457),
	979:  uint16(1),
	980:  uint16(sym__eol),
	981:  uint16(3),
	982:  uint16(301),
	983:  uint16(1),
	984:  uint16(sym_comment),
	985:  uint16(459),
	986:  uint16(1),
	987:  uint16(sym__space),
	988:  uint16(461),
	989:  uint16(1),
	990:  uint16(sym__eol),
	991:  uint16(3),
	992:  uint16(301),
	993:  uint16(1),
	994:  uint16(sym_comment),
	995:  uint16(453),
	996:  uint16(1),
	997:  uint16(sym__eol),
	998:  uint16(463),
	999:  uint16(1),
	1000: uint16(sym__space),
	1001: uint16(3),
	1002: uint16(301),
	1003: uint16(1),
	1004: uint16(sym_comment),
	1005: uint16(465),
	1006: uint16(1),
	1007: uint16(sym__space),
	1008: uint16(467),
	1009: uint16(1),
	1010: uint16(sym__eol),
	1011: uint16(3),
	1012: uint16(301),
	1013: uint16(1),
	1014: uint16(sym_comment),
	1015: uint16(461),
	1016: uint16(1),
	1017: uint16(sym__eol),
	1018: uint16(469),
	1019: uint16(1),
	1020: uint16(sym__space),
	1021: uint16(3),
	1022: uint16(3),
	1023: uint16(1),
	1024: uint16(sym_comment),
	1025: uint16(471),
	1026: uint16(1),
	1027: uint16(anon_sym_DQUOTE),
	1028: uint16(101),
	1029: uint16(1),
	1030: uint16(sym_file),
	1031: uint16(3),
	1032: uint16(3),
	1033: uint16(1),
	1034: uint16(sym_comment),
	1035: uint16(471),
	1036: uint16(1),
	1037: uint16(anon_sym_DQUOTE),
	1038: uint16(44),
	1039: uint16(1),
	1040: uint16(sym_file),
	1041: uint16(3),
	1042: uint16(3),
	1043: uint16(1),
	1044: uint16(sym_comment),
	1045: uint16(473),
	1046: uint16(1),
	1047: uint16(aux_sym__color_token1),
	1048: uint16(159),
	1049: uint16(1),
	1050: uint16(sym__color),
	1051: uint16(3),
	1052: uint16(3),
	1053: uint16(1),
	1054: uint16(sym_comment),
	1055: uint16(475),
	1056: uint16(1),
	1057: uint16(aux_sym__font_size_token1),
	1058: uint16(161),
	1059: uint16(1),
	1060: uint16(sym__font_size),
	1061: uint16(3),
	1062: uint16(3),
	1063: uint16(1),
	1064: uint16(sym_comment),
	1065: uint16(477),
	1066: uint16(1),
	1067: uint16(aux_sym__volume_token1),
	1068: uint16(162),
	1069: uint16(1),
	1070: uint16(sym__volume),
	1071: uint16(3),
	1072: uint16(301),
	1073: uint16(1),
	1074: uint16(sym_comment),
	1075: uint16(479),
	1076: uint16(1),
	1077: uint16(sym__space),
	1078: uint16(85),
	1079: uint16(1),
	1080: uint16(aux_sym_condition_repeat7),
	1081: uint16(3),
	1082: uint16(3),
	1083: uint16(1),
	1084: uint16(sym_comment),
	1085: uint16(481),
	1086: uint16(1),
	1087: uint16(aux_sym__volume_token1),
	1088: uint16(164),
	1089: uint16(1),
	1090: uint16(sym__volume),
	1091: uint16(3),
	1092: uint16(301),
	1093: uint16(1),
	1094: uint16(sym_comment),
	1095: uint16(483),
	1096: uint16(1),
	1097: uint16(sym__space),
	1098: uint16(485),
	1099: uint16(1),
	1100: uint16(sym__eol),
	1101: uint16(3),
	1102: uint16(301),
	1103: uint16(1),
	1104: uint16(sym_comment),
	1105: uint16(487),
	1106: uint16(1),
	1107: uint16(sym__space),
	1108: uint16(489),
	1109: uint16(1),
	1110: uint16(sym__eol),
	1111: uint16(3),
	1112: uint16(3),
	1113: uint16(1),
	1114: uint16(sym_comment),
	1115: uint16(491),
	1116: uint16(1),
	1117: uint16(aux_sym__color_token1),
	1118: uint16(169),
	1119: uint16(1),
	1120: uint16(sym__color),
	1121: uint16(3),
	1122: uint16(301),
	1123: uint16(1),
	1124: uint16(sym_comment),
	1125: uint16(339),
	1126: uint16(1),
	1127: uint16(sym__space),
	1128: uint16(57),
	1129: uint16(1),
	1130: uint16(aux_sym_condition_repeat1),
	1131: uint16(3),
	1132: uint16(301),
	1133: uint16(1),
	1134: uint16(sym_comment),
	1135: uint16(493),
	1136: uint16(1),
	1137: uint16(sym__space),
	1138: uint16(495),
	1139: uint16(1),
	1140: uint16(sym__eol),
	1141: uint16(3),
	1142: uint16(301),
	1143: uint16(1),
	1144: uint16(sym_comment),
	1145: uint16(497),
	1146: uint16(1),
	1147: uint16(sym__space),
	1148: uint16(499),
	1149: uint16(1),
	1150: uint16(sym__eol),
	1151: uint16(3),
	1152: uint16(301),
	1153: uint16(1),
	1154: uint16(sym_comment),
	1155: uint16(501),
	1156: uint16(1),
	1157: uint16(sym__space),
	1158: uint16(80),
	1159: uint16(1),
	1160: uint16(aux_sym_condition_repeat1),
	1161: uint16(3),
	1162: uint16(301),
	1163: uint16(1),
	1164: uint16(sym_comment),
	1165: uint16(503),
	1166: uint16(1),
	1167: uint16(sym__space),
	1168: uint16(505),
	1169: uint16(1),
	1170: uint16(sym__eol),
	1171: uint16(3),
	1172: uint16(301),
	1173: uint16(1),
	1174: uint16(sym_comment),
	1175: uint16(507),
	1176: uint16(1),
	1177: uint16(sym__space),
	1178: uint16(509),
	1179: uint16(1),
	1180: uint16(sym__eol),
	1181: uint16(3),
	1182: uint16(301),
	1183: uint16(1),
	1184: uint16(sym_comment),
	1185: uint16(511),
	1186: uint16(1),
	1187: uint16(sym__space),
	1188: uint16(513),
	1189: uint16(1),
	1190: uint16(sym__eol),
	1191: uint16(3),
	1192: uint16(301),
	1193: uint16(1),
	1194: uint16(sym_comment),
	1195: uint16(515),
	1196: uint16(1),
	1197: uint16(sym__space),
	1198: uint16(517),
	1199: uint16(1),
	1200: uint16(sym__eol),
	1201: uint16(3),
	1202: uint16(301),
	1203: uint16(1),
	1204: uint16(sym_comment),
	1205: uint16(331),
	1206: uint16(1),
	1207: uint16(sym__space),
	1208: uint16(58),
	1209: uint16(1),
	1210: uint16(aux_sym_condition_repeat2),
	1211: uint16(3),
	1212: uint16(301),
	1213: uint16(1),
	1214: uint16(sym_comment),
	1215: uint16(519),
	1216: uint16(1),
	1217: uint16(sym__space),
	1218: uint16(521),
	1219: uint16(1),
	1220: uint16(sym__eol),
	1221: uint16(3),
	1222: uint16(301),
	1223: uint16(1),
	1224: uint16(sym_comment),
	1225: uint16(523),
	1226: uint16(1),
	1227: uint16(sym__space),
	1228: uint16(70),
	1229: uint16(1),
	1230: uint16(aux_sym_condition_repeat3),
	1231: uint16(3),
	1232: uint16(301),
	1233: uint16(1),
	1234: uint16(sym_comment),
	1235: uint16(525),
	1236: uint16(1),
	1237: uint16(sym__space),
	1238: uint16(72),
	1239: uint16(1),
	1240: uint16(aux_sym_condition_repeat4),
	1241: uint16(3),
	1242: uint16(301),
	1243: uint16(1),
	1244: uint16(sym_comment),
	1245: uint16(345),
	1246: uint16(1),
	1247: uint16(sym__space),
	1248: uint16(59),
	1249: uint16(1),
	1250: uint16(aux_sym_condition_repeat3),
	1251: uint16(3),
	1252: uint16(301),
	1253: uint16(1),
	1254: uint16(sym_comment),
	1255: uint16(527),
	1256: uint16(1),
	1257: uint16(sym__space),
	1258: uint16(529),
	1259: uint16(1),
	1260: uint16(sym__eol),
	1261: uint16(3),
	1262: uint16(3),
	1263: uint16(1),
	1264: uint16(sym_comment),
	1265: uint16(531),
	1266: uint16(1),
	1267: uint16(aux_sym__color_token1),
	1268: uint16(130),
	1269: uint16(1),
	1270: uint16(sym__color),
	1271: uint16(3),
	1272: uint16(301),
	1273: uint16(1),
	1274: uint16(sym_comment),
	1275: uint16(533),
	1276: uint16(1),
	1277: uint16(sym__space),
	1278: uint16(535),
	1279: uint16(1),
	1280: uint16(sym__eol),
	1281: uint16(3),
	1282: uint16(301),
	1283: uint16(1),
	1284: uint16(sym_comment),
	1285: uint16(398),
	1286: uint16(1),
	1287: uint16(sym__space),
	1288: uint16(77),
	1289: uint16(1),
	1290: uint16(aux_sym_condition_repeat5),
	1291: uint16(3),
	1292: uint16(301),
	1293: uint16(1),
	1294: uint16(sym_comment),
	1295: uint16(537),
	1296: uint16(1),
	1297: uint16(sym__space),
	1298: uint16(83),
	1299: uint16(1),
	1300: uint16(aux_sym_condition_repeat1),
	1301: uint16(3),
	1302: uint16(3),
	1303: uint16(1),
	1304: uint16(sym_comment),
	1305: uint16(539),
	1306: uint16(1),
	1307: uint16(aux_sym__color_token1),
	1308: uint16(186),
	1309: uint16(1),
	1310: uint16(sym__color),
	1311: uint16(3),
	1312: uint16(301),
	1313: uint16(1),
	1314: uint16(sym_comment),
	1315: uint16(335),
	1316: uint16(1),
	1317: uint16(sym__space),
	1318: uint16(56),
	1319: uint16(1),
	1320: uint16(aux_sym_condition_repeat6),
	1321: uint16(3),
	1322: uint16(301),
	1323: uint16(1),
	1324: uint16(sym_comment),
	1325: uint16(349),
	1326: uint16(1),
	1327: uint16(sym__space),
	1328: uint16(60),
	1329: uint16(1),
	1330: uint16(aux_sym_condition_repeat4),
	1331: uint16(3),
	1332: uint16(3),
	1333: uint16(1),
	1334: uint16(sym_comment),
	1335: uint16(541),
	1336: uint16(1),
	1337: uint16(aux_sym__volume_token1),
	1338: uint16(163),
	1339: uint16(1),
	1340: uint16(sym__volume),
	1341: uint16(2),
	1342: uint16(3),
	1343: uint16(1),
	1344: uint16(sym_comment),
	1345: uint16(543),
	1346: uint16(1),
	1347: uint16(sym__eol),
	1348: uint16(2),
	1349: uint16(3),
	1350: uint16(1),
	1351: uint16(sym_comment),
	1352: uint16(311),
	1353: uint16(1),
	1354: uint16(sym_boolean),
	1355: uint16(2),
	1356: uint16(3),
	1357: uint16(1),
	1358: uint16(sym_comment),
	1359: uint16(545),
	1360: uint16(1),
	1362: uint16(2),
	1363: uint16(301),
	1364: uint16(1),
	1365: uint16(sym_comment),
	1366: uint16(547),
	1367: uint16(1),
	1368: uint16(sym__space),
	1369: uint16(2),
	1370: uint16(3),
	1371: uint16(1),
	1372: uint16(sym_comment),
	1373: uint16(549),
	1374: uint16(1),
	1375: uint16(anon_sym_DQUOTE),
	1376: uint16(2),
	1377: uint16(301),
	1378: uint16(1),
	1379: uint16(sym_comment),
	1380: uint16(551),
	1381: uint16(1),
	1382: uint16(sym__space),
	1383: uint16(2),
	1384: uint16(3),
	1385: uint16(1),
	1386: uint16(sym_comment),
	1387: uint16(553),
	1388: uint16(1),
	1389: uint16(anon_sym_Optional),
	1390: uint16(2),
	1391: uint16(3),
	1392: uint16(1),
	1393: uint16(sym_comment),
	1394: uint16(555),
	1395: uint16(1),
	1396: uint16(sym__eol),
	1397: uint16(2),
	1398: uint16(3),
	1399: uint16(1),
	1400: uint16(sym_comment),
	1401: uint16(557),
	1402: uint16(1),
	1403: uint16(aux_sym_rarity_token1),
	1404: uint16(2),
	1405: uint16(301),
	1406: uint16(1),
	1407: uint16(sym_comment),
	1408: uint16(559),
	1409: uint16(1),
	1410: uint16(aux_sym_string_token1),
	1411: uint16(2),
	1412: uint16(301),
	1413: uint16(1),
	1414: uint16(sym_comment),
	1415: uint16(561),
	1416: uint16(1),
	1417: uint16(sym__space),
	1418: uint16(2),
	1419: uint16(3),
	1420: uint16(1),
	1421: uint16(sym_comment),
	1422: uint16(563),
	1423: uint16(1),
	1424: uint16(anon_sym_Temp),
	1425: uint16(2),
	1426: uint16(3),
	1427: uint16(1),
	1428: uint16(sym_comment),
	1429: uint16(565),
	1430: uint16(1),
	1431: uint16(anon_sym_Temp),
	1432: uint16(2),
	1433: uint16(301),
	1434: uint16(1),
	1435: uint16(sym_comment),
	1436: uint16(567),
	1437: uint16(1),
	1438: uint16(sym__space),
	1439: uint16(2),
	1440: uint16(3),
	1441: uint16(1),
	1442: uint16(sym_comment),
	1443: uint16(569),
	1444: uint16(1),
	1445: uint16(sym__eol),
	1446: uint16(2),
	1447: uint16(301),
	1448: uint16(1),
	1449: uint16(sym_comment),
	1450: uint16(571),
	1451: uint16(1),
	1452: uint16(sym__space),
	1453: uint16(2),
	1454: uint16(301),
	1455: uint16(1),
	1456: uint16(sym_comment),
	1457: uint16(573),
	1458: uint16(1),
	1459: uint16(sym__space),
	1460: uint16(2),
	1461: uint16(301),
	1462: uint16(1),
	1463: uint16(sym_comment),
	1464: uint16(575),
	1465: uint16(1),
	1466: uint16(sym__space),
	1467: uint16(2),
	1468: uint16(301),
	1469: uint16(1),
	1470: uint16(sym_comment),
	1471: uint16(577),
	1472: uint16(1),
	1473: uint16(sym__space),
	1474: uint16(2),
	1475: uint16(301),
	1476: uint16(1),
	1477: uint16(sym_comment),
	1478: uint16(579),
	1479: uint16(1),
	1480: uint16(sym__space),
	1481: uint16(2),
	1482: uint16(301),
	1483: uint16(1),
	1484: uint16(sym_comment),
	1485: uint16(581),
	1486: uint16(1),
	1487: uint16(sym__space),
	1488: uint16(2),
	1489: uint16(301),
	1490: uint16(1),
	1491: uint16(sym_comment),
	1492: uint16(583),
	1493: uint16(1),
	1494: uint16(sym__space),
	1495: uint16(2),
	1496: uint16(301),
	1497: uint16(1),
	1498: uint16(sym_comment),
	1499: uint16(585),
	1500: uint16(1),
	1501: uint16(sym__space),
	1502: uint16(2),
	1503: uint16(3),
	1504: uint16(1),
	1505: uint16(sym_comment),
	1506: uint16(587),
	1507: uint16(1),
	1508: uint16(sym_number),
	1509: uint16(2),
	1510: uint16(3),
	1511: uint16(1),
	1512: uint16(sym_comment),
	1513: uint16(589),
	1514: uint16(1),
	1515: uint16(sym__eol),
	1516: uint16(2),
	1517: uint16(3),
	1518: uint16(1),
	1519: uint16(sym_comment),
	1520: uint16(591),
	1521: uint16(1),
	1522: uint16(sym__eol),
	1523: uint16(2),
	1524: uint16(3),
	1525: uint16(1),
	1526: uint16(sym_comment),
	1527: uint16(593),
	1528: uint16(1),
	1529: uint16(sym__eol),
	1530: uint16(2),
	1531: uint16(3),
	1532: uint16(1),
	1533: uint16(sym_comment),
	1534: uint16(595),
	1535: uint16(1),
	1536: uint16(sym__eol),
	1537: uint16(2),
	1538: uint16(3),
	1539: uint16(1),
	1540: uint16(sym_comment),
	1541: uint16(597),
	1542: uint16(1),
	1543: uint16(anon_sym_DQUOTE),
	1544: uint16(2),
	1545: uint16(301),
	1546: uint16(1),
	1547: uint16(sym_comment),
	1548: uint16(599),
	1549: uint16(1),
	1550: uint16(sym__space),
	1551: uint16(2),
	1552: uint16(3),
	1553: uint16(1),
	1554: uint16(sym_comment),
	1555: uint16(601),
	1556: uint16(1),
	1557: uint16(sym__eol),
	1558: uint16(2),
	1559: uint16(3),
	1560: uint16(1),
	1561: uint16(sym_comment),
	1562: uint16(603),
	1563: uint16(1),
	1564: uint16(sym__eol),
	1565: uint16(2),
	1566: uint16(301),
	1567: uint16(1),
	1568: uint16(sym_comment),
	1569: uint16(605),
	1570: uint16(1),
	1571: uint16(sym__space),
	1572: uint16(2),
	1573: uint16(3),
	1574: uint16(1),
	1575: uint16(sym_comment),
	1576: uint16(607),
	1577: uint16(1),
	1578: uint16(sym__eol),
	1579: uint16(2),
	1580: uint16(3),
	1581: uint16(1),
	1582: uint16(sym_comment),
	1583: uint16(609),
	1584: uint16(1),
	1585: uint16(aux_sym_sockets_token1),
	1586: uint16(2),
	1587: uint16(3),
	1588: uint16(1),
	1589: uint16(sym_comment),
	1590: uint16(611),
	1591: uint16(1),
	1592: uint16(sym__eol),
	1593: uint16(2),
	1594: uint16(3),
	1595: uint16(1),
	1596: uint16(sym_comment),
	1597: uint16(613),
	1598: uint16(1),
	1599: uint16(sym_boolean),
	1600: uint16(2),
	1601: uint16(3),
	1602: uint16(1),
	1603: uint16(sym_comment),
	1604: uint16(615),
	1605: uint16(1),
	1606: uint16(anon_sym_DQUOTE),
	1607: uint16(2),
	1608: uint16(301),
	1609: uint16(1),
	1610: uint16(sym_comment),
	1611: uint16(617),
	1612: uint16(1),
	1613: uint16(sym__space),
	1614: uint16(2),
	1615: uint16(3),
	1616: uint16(1),
	1617: uint16(sym_comment),
	1618: uint16(619),
	1619: uint16(1),
	1620: uint16(sym__eol),
	1621: uint16(2),
	1622: uint16(301),
	1623: uint16(1),
	1624: uint16(sym_comment),
	1625: uint16(621),
	1626: uint16(1),
	1627: uint16(aux_sym_file_token1),
	1628: uint16(2),
	1629: uint16(3),
	1630: uint16(1),
	1631: uint16(sym_comment),
	1632: uint16(623),
	1633: uint16(1),
	1634: uint16(sym__eol),
	1635: uint16(2),
	1636: uint16(3),
	1637: uint16(1),
	1638: uint16(sym_comment),
	1639: uint16(625),
	1640: uint16(1),
	1641: uint16(sym__eol),
	1642: uint16(2),
	1643: uint16(3),
	1644: uint16(1),
	1645: uint16(sym_comment),
	1646: uint16(627),
	1647: uint16(1),
	1648: uint16(sym__eol),
	1649: uint16(2),
	1650: uint16(3),
	1651: uint16(1),
	1652: uint16(sym_comment),
	1653: uint16(629),
	1654: uint16(1),
	1655: uint16(aux_sym_influence_token1),
	1656: uint16(2),
	1657: uint16(3),
	1658: uint16(1),
	1659: uint16(sym_comment),
	1660: uint16(631),
	1661: uint16(1),
	1662: uint16(aux_sym_quality_token1),
	1663: uint16(2),
	1664: uint16(3),
	1665: uint16(1),
	1666: uint16(sym_comment),
	1667: uint16(633),
	1668: uint16(1),
	1669: uint16(anon_sym_DQUOTE),
	1670: uint16(2),
	1671: uint16(3),
	1672: uint16(1),
	1673: uint16(sym_comment),
	1674: uint16(635),
	1675: uint16(1),
	1676: uint16(anon_sym_DQUOTE),
	1677: uint16(2),
	1678: uint16(3),
	1679: uint16(1),
	1680: uint16(sym_comment),
	1681: uint16(637),
	1682: uint16(1),
	1683: uint16(anon_sym_DQUOTE),
	1684: uint16(2),
	1685: uint16(3),
	1686: uint16(1),
	1687: uint16(sym_comment),
	1688: uint16(639),
	1689: uint16(1),
	1690: uint16(sym__eol),
	1691: uint16(2),
	1692: uint16(301),
	1693: uint16(1),
	1694: uint16(sym_comment),
	1695: uint16(641),
	1696: uint16(1),
	1697: uint16(sym__space),
	1698: uint16(2),
	1699: uint16(3),
	1700: uint16(1),
	1701: uint16(sym_comment),
	1702: uint16(487),
	1703: uint16(1),
	1704: uint16(sym__eol),
	1705: uint16(2),
	1706: uint16(301),
	1707: uint16(1),
	1708: uint16(sym_comment),
	1709: uint16(619),
	1710: uint16(1),
	1711: uint16(sym__space),
	1712: uint16(2),
	1713: uint16(3),
	1714: uint16(1),
	1715: uint16(sym_comment),
	1716: uint16(493),
	1717: uint16(1),
	1718: uint16(sym__eol),
	1719: uint16(2),
	1720: uint16(3),
	1721: uint16(1),
	1722: uint16(sym_comment),
	1723: uint16(643),
	1724: uint16(1),
	1725: uint16(sym_number),
	1726: uint16(2),
	1727: uint16(3),
	1728: uint16(1),
	1729: uint16(sym_comment),
	1730: uint16(645),
	1731: uint16(1),
	1732: uint16(anon_sym_DQUOTE),
	1733: uint16(2),
	1734: uint16(301),
	1735: uint16(1),
	1736: uint16(sym_comment),
	1737: uint16(647),
	1738: uint16(1),
	1739: uint16(aux_sym_string_token1),
	1740: uint16(2),
	1741: uint16(301),
	1742: uint16(1),
	1743: uint16(sym_comment),
	1744: uint16(649),
	1745: uint16(1),
	1746: uint16(sym__space),
}

var ts_small_parse_table_map = [162]uint32_t{
	1:   uint32(21),
	2:   uint32(42),
	3:   uint32(65),
	4:   uint32(85),
	5:   uint32(105),
	6:   uint32(129),
	7:   uint32(148),
	8:   uint32(171),
	9:   uint32(198),
	10:  uint32(221),
	11:  uint32(242),
	12:  uint32(259),
	13:  uint32(274),
	14:  uint32(292),
	15:  uint32(308),
	16:  uint32(320),
	17:  uint32(336),
	18:  uint32(355),
	19:  uint32(374),
	20:  uint32(393),
	21:  uint32(412),
	22:  uint32(428),
	23:  uint32(441),
	24:  uint32(454),
	25:  uint32(467),
	26:  uint32(480),
	27:  uint32(493),
	28:  uint32(506),
	29:  uint32(519),
	30:  uint32(532),
	31:  uint32(545),
	32:  uint32(558),
	33:  uint32(571),
	34:  uint32(584),
	35:  uint32(597),
	36:  uint32(610),
	37:  uint32(623),
	38:  uint32(636),
	39:  uint32(649),
	40:  uint32(662),
	41:  uint32(675),
	42:  uint32(688),
	43:  uint32(701),
	44:  uint32(714),
	45:  uint32(727),
	46:  uint32(740),
	47:  uint32(753),
	48:  uint32(766),
	49:  uint32(779),
	50:  uint32(792),
	51:  uint32(805),
	52:  uint32(818),
	53:  uint32(831),
	54:  uint32(841),
	55:  uint32(851),
	56:  uint32(861),
	57:  uint32(871),
	58:  uint32(881),
	59:  uint32(891),
	60:  uint32(901),
	61:  uint32(911),
	62:  uint32(921),
	63:  uint32(931),
	64:  uint32(941),
	65:  uint32(951),
	66:  uint32(961),
	67:  uint32(971),
	68:  uint32(981),
	69:  uint32(991),
	70:  uint32(1001),
	71:  uint32(1011),
	72:  uint32(1021),
	73:  uint32(1031),
	74:  uint32(1041),
	75:  uint32(1051),
	76:  uint32(1061),
	77:  uint32(1071),
	78:  uint32(1081),
	79:  uint32(1091),
	80:  uint32(1101),
	81:  uint32(1111),
	82:  uint32(1121),
	83:  uint32(1131),
	84:  uint32(1141),
	85:  uint32(1151),
	86:  uint32(1161),
	87:  uint32(1171),
	88:  uint32(1181),
	89:  uint32(1191),
	90:  uint32(1201),
	91:  uint32(1211),
	92:  uint32(1221),
	93:  uint32(1231),
	94:  uint32(1241),
	95:  uint32(1251),
	96:  uint32(1261),
	97:  uint32(1271),
	98:  uint32(1281),
	99:  uint32(1291),
	100: uint32(1301),
	101: uint32(1311),
	102: uint32(1321),
	103: uint32(1331),
	104: uint32(1341),
	105: uint32(1348),
	106: uint32(1355),
	107: uint32(1362),
	108: uint32(1369),
	109: uint32(1376),
	110: uint32(1383),
	111: uint32(1390),
	112: uint32(1397),
	113: uint32(1404),
	114: uint32(1411),
	115: uint32(1418),
	116: uint32(1425),
	117: uint32(1432),
	118: uint32(1439),
	119: uint32(1446),
	120: uint32(1453),
	121: uint32(1460),
	122: uint32(1467),
	123: uint32(1474),
	124: uint32(1481),
	125: uint32(1488),
	126: uint32(1495),
	127: uint32(1502),
	128: uint32(1509),
	129: uint32(1516),
	130: uint32(1523),
	131: uint32(1530),
	132: uint32(1537),
	133: uint32(1544),
	134: uint32(1551),
	135: uint32(1558),
	136: uint32(1565),
	137: uint32(1572),
	138: uint32(1579),
	139: uint32(1586),
	140: uint32(1593),
	141: uint32(1600),
	142: uint32(1607),
	143: uint32(1614),
	144: uint32(1621),
	145: uint32(1628),
	146: uint32(1635),
	147: uint32(1642),
	148: uint32(1649),
	149: uint32(1656),
	150: uint32(1663),
	151: uint32(1670),
	152: uint32(1677),
	153: uint32(1684),
	154: uint32(1691),
	155: uint32(1698),
	156: uint32(1705),
	157: uint32(1712),
	158: uint32(1719),
	159: uint32(1726),
	160: uint32(1733),
	161: uint32(1740),
}

var ts_parse_actions = [651]TSParseActionEntry{
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
		Fsymbol:     uint16(sym_filter),
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
		Fstate: uint16(libc.Int32FromInt32(172)),
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
		Fstate: uint16(libc.Int32FromInt32(175)),
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
		Fstate: uint16(libc.Int32FromInt32(40)),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_block),
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
		Fstate: uint16(libc.Int32FromInt32(154)),
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
		Fstate: uint16(libc.Int32FromInt32(118)),
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
		Fstate: uint16(libc.Int32FromInt32(156)),
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
		Fstate: uint16(libc.Int32FromInt32(86)),
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
		Fstate: uint16(libc.Int32FromInt32(125)),
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
		Fstate: uint16(libc.Int32FromInt32(154)),
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
		Fstate: uint16(libc.Int32FromInt32(126)),
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
		Fstate: uint16(libc.Int32FromInt32(131)),
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
		Fstate: uint16(libc.Int32FromInt32(132)),
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
		Fstate: uint16(libc.Int32FromInt32(134)),
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
		Fstate: uint16(libc.Int32FromInt32(110)),
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
		Fstate: uint16(libc.Int32FromInt32(87)),
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
		Fstate: uint16(libc.Int32FromInt32(155)),
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
		Fstate: uint16(libc.Int32FromInt32(187)),
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
		Fstate: uint16(libc.Int32FromInt32(187)),
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
		Fstate: uint16(libc.Int32FromInt32(147)),
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
		Fstate: uint16(libc.Int32FromInt32(150)),
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
		Fstate: uint16(libc.Int32FromInt32(92)),
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
		Fstate: uint16(libc.Int32FromInt32(92)),
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
		Fstate: uint16(libc.Int32FromInt32(157)),
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
		Fstate: uint16(libc.Int32FromInt32(158)),
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
		Fstate: uint16(libc.Int32FromInt32(140)),
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
		Fstate: uint16(libc.Int32FromInt32(142)),
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
		Fstate: uint16(libc.Int32FromInt32(170)),
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
		Fstate: uint16(libc.Int32FromInt32(3)),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_block_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(154)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate:      uint16(libc.Int32FromInt32(118)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate:      uint16(libc.Int32FromInt32(156)),
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
		Fstate:      uint16(libc.Int32FromInt32(86)),
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
		Fstate:      uint16(libc.Int32FromInt32(125)),
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
		Fstate:      uint16(libc.Int32FromInt32(154)),
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
		Fcount:    uint8(2),
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
		Fstate:      uint16(libc.Int32FromInt32(126)),
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
		Fcount:    uint8(2),
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
		Fstate:      uint16(libc.Int32FromInt32(131)),
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
		Fcount:    uint8(2),
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
		Fstate:      uint16(libc.Int32FromInt32(132)),
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
		Fcount:    uint8(2),
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
		Fstate:      uint16(libc.Int32FromInt32(134)),
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
		Fcount:    uint8(2),
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
		Fstate:      uint16(libc.Int32FromInt32(110)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate:      uint16(libc.Int32FromInt32(87)),
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
		Fcount:    uint8(2),
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
		Fstate:      uint16(libc.Int32FromInt32(155)),
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
		Fcount: uint8(2),
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
		Fstate:      uint16(libc.Int32FromInt32(187)),
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
		Fcount:    uint8(2),
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
		Fstate:      uint16(libc.Int32FromInt32(187)),
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
		Fcount: uint8(2),
	}})),
	113: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      uint16(libc.Int32FromInt32(147)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_block_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(150)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_block_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(92)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_block_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(92)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_block_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(157)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_block_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(158)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	131: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      uint16(libc.Int32FromInt32(140)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_block_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(142)),
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
		Fsymbol:      uint16(aux_sym_block_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(170)),
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
		Fsymbol:      uint16(aux_sym_block_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(3)),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_condition),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_condition),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(10),
		Fsymbol:        uint16(sym_action),
		Fproduction_id: uint16(26),
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
		Fcount: uint8(1),
	}})),
	149: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(10),
		Fsymbol:        uint16(sym_action),
		Fproduction_id: uint16(26),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_action),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_action),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_continue),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_continue),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_condition),
		Fproduction_id: uint16(2),
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
		Fcount: uint8(1),
	}})),
	161: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_condition),
		Fproduction_id: uint16(2),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_condition),
		Fproduction_id: uint16(5),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_condition),
		Fproduction_id: uint16(5),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_condition),
		Fproduction_id: uint16(8),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_condition),
		Fproduction_id: uint16(8),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_condition),
		Fproduction_id: uint16(11),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_condition),
		Fproduction_id: uint16(11),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_condition),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_condition),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_condition),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_condition),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_condition),
		Fproduction_id: uint16(14),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_condition),
		Fproduction_id: uint16(14),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_action),
		Fproduction_id: uint16(15),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_action),
		Fproduction_id: uint16(15),
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
		Fsymbol:      uint16(sym_action),
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
		Fsymbol:      uint16(sym_action),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_action),
		Fproduction_id: uint16(17),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_action),
		Fproduction_id: uint16(17),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_condition),
		Fproduction_id: uint16(18),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_condition),
		Fproduction_id: uint16(18),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_condition),
		Fproduction_id: uint16(19),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_condition),
		Fproduction_id: uint16(19),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_condition),
		Fproduction_id: uint16(20),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_condition),
		Fproduction_id: uint16(20),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_condition),
		Fproduction_id: uint16(21),
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
		Fcount: uint8(1),
	}})),
	213: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_condition),
		Fproduction_id: uint16(21),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_action),
		Fproduction_id: uint16(22),
	})))),
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
		Fcount: uint8(1),
	}})),
	217: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_action),
		Fproduction_id: uint16(22),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_action),
		Fproduction_id: uint16(23),
	})))),
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
		Fcount: uint8(1),
	}})),
	221: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_action),
		Fproduction_id: uint16(23),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_action),
		Fproduction_id: uint16(24),
	})))),
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
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_action),
		Fproduction_id: uint16(24),
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
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_action),
		Fproduction_id: uint16(15),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_action),
		Fproduction_id: uint16(15),
	})))),
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
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_action),
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
		Fcount: uint8(1),
	}})),
	233: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_action),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	235: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_action),
		Fproduction_id: uint16(15),
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
		Fcount: uint8(1),
	}})),
	237: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_action),
		Fproduction_id: uint16(15),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	239: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_action),
		Fproduction_id: uint16(17),
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
		Fcount: uint8(1),
	}})),
	241: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_action),
		Fproduction_id: uint16(17),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	243: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_action),
		Fproduction_id: uint16(25),
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
		Fcount: uint8(1),
	}})),
	245: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_action),
		Fproduction_id: uint16(25),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_action),
		Fproduction_id: uint16(16),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_action),
		Fproduction_id: uint16(16),
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
		Fstate: uint16(libc.Int32FromInt32(2)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(178)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	255: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(102)),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(103)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(94)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(145)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(93)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(90)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(181)),
	}})))),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_filter),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(42)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(191)),
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
		Fstate: uint16(libc.Int32FromInt32(146)),
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
		Fstate: uint16(libc.Int32FromInt32(113)),
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
		Fstate: uint16(libc.Int32FromInt32(115)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_filter_repeat1),
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
		Fsymbol:      uint16(aux_sym_filter_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(172)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_filter_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(175)),
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
		Fsymbol:      uint16(aux_sym_filter_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(42)),
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
		Fstate: uint16(libc.Int32FromInt32(97)),
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
		Fstate: uint16(libc.Int32FromInt32(171)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(96)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_import),
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
		Fsymbol:      uint16(sym_import),
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
		Fextra: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(143)),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_file),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_file),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(160)),
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
		Fstate: uint16(libc.Int32FromInt32(137)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_import),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(182)),
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
		Fstate: uint16(libc.Int32FromInt32(88)),
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
		Fstate: uint16(libc.Int32FromInt32(123)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(127)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(135)),
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
		Fstate: uint16(libc.Int32FromInt32(137)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(193)),
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
		Fstate: uint16(libc.Int32FromInt32(188)),
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
		Fstate: uint16(libc.Int32FromInt32(76)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(9)),
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
		Fstate: uint16(libc.Int32FromInt32(39)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(12)),
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
		Fstate: uint16(libc.Int32FromInt32(73)),
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
		Fstate: uint16(libc.Int32FromInt32(19)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(20)),
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
		Fstate: uint16(libc.Int32FromInt32(79)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(21)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(82)),
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
		Fcount: uint8(1),
	}})),
	352: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(22)),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(47)),
	}})))),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(4)),
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
		Fstate: uint16(libc.Int32FromInt32(65)),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_condition_repeat5),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(49)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(1),
	}})),
	363: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_condition_repeat5),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(99)),
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
		Fstate: uint16(libc.Int32FromInt32(100)),
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
		Fstate: uint16(libc.Int32FromInt32(177)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	371: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_condition_repeat8),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(65)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_condition_repeat8),
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
		Fstate: uint16(libc.Int32FromInt32(152)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(153)),
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
		Fstate: uint16(libc.Int32FromInt32(10)),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_condition_repeat6),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(39)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_condition_repeat6),
	})))),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(11)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	389: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_condition_repeat7),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(47)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_condition_repeat7),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_condition_repeat1),
		Fproduction_id: uint16(3),
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
		Fstate:      uint16(libc.Int32FromInt32(73)),
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
		Fcount: uint8(1),
	}})),
	397: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_condition_repeat1),
		Fproduction_id: uint16(3),
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
		Fstate: uint16(libc.Int32FromInt32(49)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	401: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_condition_repeat2),
		Fproduction_id: uint16(6),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(76)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_condition_repeat2),
		Fproduction_id: uint16(6),
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
		Fstate: uint16(libc.Int32FromInt32(8)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	408: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_condition_repeat3),
		Fproduction_id: uint16(9),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(79)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(1),
	}})),
	411: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_condition_repeat3),
		Fproduction_id: uint16(9),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(13)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	415: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_condition_repeat4),
		Fproduction_id: uint16(12),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(82)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(1),
	}})),
	418: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_condition_repeat4),
		Fproduction_id: uint16(12),
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
		Fstate: uint16(libc.Int32FromInt32(51)),
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
		Fstate: uint16(libc.Int32FromInt32(43)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	424: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_quality),
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
		Fcount: uint8(1),
	}})),
	426: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_quality),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	428: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_condition_repeat5),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_influence),
	})))),
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
		Fcount: uint8(1),
	}})),
	432: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_influence),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_condition_repeat6),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	436: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(173)),
	}})))),
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
		Fcount: uint8(1),
	}})),
	438: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(6)),
	}})))),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_rarity),
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
		Fcount: uint8(1),
	}})),
	442: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_rarity),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_condition_repeat7),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_sockets),
		Fproduction_id: uint16(13),
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
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_sockets),
		Fproduction_id: uint16(13),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_condition_repeat8),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(109)),
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
		Fcount: uint8(1),
	}})),
	454: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(16)),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	456: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(136)),
	}})))),
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
		Fcount: uint8(1),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(31)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	460: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(111)),
	}})))),
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
		Fstate: uint16(libc.Int32FromInt32(17)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(148)),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_colour),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_colour),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(149)),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(177)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(159)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(161)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(162)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(38)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(164)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	484: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_condition_repeat4),
		Fproduction_id: uint16(10),
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
		Fcount: uint8(1),
	}})),
	486: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_condition_repeat4),
		Fproduction_id: uint16(10),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	488: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fcount: uint8(1),
	}})),
	490: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(169)),
	}})))),
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
		Fsymbol:      uint16(sym_string),
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
		Fcount: uint8(1),
	}})),
	496: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	498: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_condition_repeat1),
		Fproduction_id: uint16(1),
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
		Fcount: uint8(1),
	}})),
	500: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_condition_repeat1),
		Fproduction_id: uint16(1),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	502: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(50)),
	}})))),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_quality),
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
		Fcount: uint8(1),
	}})),
	506: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_quality),
	})))),
	507: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	508: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_influence),
	})))),
	509: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	510: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_influence),
	})))),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_rarity),
	})))),
	513: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	514: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_rarity),
	})))),
	515: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	516: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_sockets),
	})))),
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
		Fcount: uint8(1),
	}})),
	518: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_sockets),
	})))),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_condition_repeat2),
		Fproduction_id: uint16(4),
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
		Fcount: uint8(1),
	}})),
	522: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_condition_repeat2),
		Fproduction_id: uint16(4),
	})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(52)),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(53)),
	}})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_condition_repeat3),
		Fproduction_id: uint16(7),
	})))),
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
		Fcount: uint8(1),
	}})),
	530: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_condition_repeat3),
		Fproduction_id: uint16(7),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(130)),
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
		Fstate: uint16(libc.Int32FromInt32(133)),
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
		Fcount: uint8(1),
	}})),
	536: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(30)),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(41)),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(186)),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(163)),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(14)),
	}})))),
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
	546: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f uint8_t
		_ [7]byte
	}{f: uint8(TSParseActionTypeAccept)})),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(107)),
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
		Fstate: uint16(libc.Int32FromInt32(45)),
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
		Fstate: uint16(libc.Int32FromInt32(108)),
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
		Fstate: uint16(libc.Int32FromInt32(48)),
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
		Fstate: uint16(libc.Int32FromInt32(15)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(185)),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	560: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(174)),
	}})))),
	561: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(67)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	564: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(167)),
	}})))),
	565: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	566: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(168)),
	}})))),
	567: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(105)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	570: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(17)),
	}})))),
	571: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	572: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(37)),
	}})))),
	573: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	574: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(36)),
	}})))),
	575: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	576: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(138)),
	}})))),
	577: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	578: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(54)),
	}})))),
	579: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(46)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	582: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(69)),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	584: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(35)),
	}})))),
	585: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	586: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(114)),
	}})))),
	587: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(176)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	590: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(18)),
	}})))),
	591: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	592: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(23)),
	}})))),
	593: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(24)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	596: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(25)),
	}})))),
	597: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	598: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(122)),
	}})))),
	599: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(34)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	602: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(26)),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(27)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(129)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(7)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(165)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(32)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(151)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	616: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(116)),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(106)),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	620: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__quantity),
	})))),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(141)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_shape),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(28)),
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
		Fstate: uint16(libc.Int32FromInt32(29)),
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
		Fstate: uint16(libc.Int32FromInt32(184)),
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
		Fstate: uint16(libc.Int32FromInt32(183)),
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
		Fstate: uint16(libc.Int32FromInt32(119)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(120)),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(121)),
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
		Fstate: uint16(libc.Int32FromInt32(5)),
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
		Fstate: uint16(libc.Int32FromInt32(66)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(189)),
	}})))),
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
		Fstate: uint16(libc.Int32FromInt32(190)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(192)),
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
		Fstate: uint16(libc.Int32FromInt32(33)),
	}})))),
}

func tree_sitter_poe_filter(tls *libc.TLS) (r uintptr) {
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
}

func init() {
	p := unsafe.Pointer(&language)
	*(*uintptr)(unsafe.Add(p, 144)) = __ccgo_fp(ts_lex)
}

func __ccgo_fp(f interface{}) uintptr {
	type iface [2]uintptr
	return (*iface)(unsafe.Pointer(&f))[1]
}

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = "end\x00Show\x00Hide\x00Minimal\x00Import\x00Optional\x00name\x00None\x00disable\x00Temp\x00Continue\x00operator\x00boolean\x00\"\x00quality_token1\x00Superior\x00Divergent\x00Anomalous\x00Phantasmal\x00rarity_token1\x00Normal\x00Magic\x00Rare\x00Unique\x00influence_token1\x00Shaper\x00Elder\x00Crusader\x00Hunter\x00Redeemer\x00Warlord\x00sockets_token1\x00sockets_token2\x00Red\x00Green\x00Blue\x00Brown\x00White\x00Yellow\x00Cyan\x00Grey\x00Orange\x00Pink\x00Purple\x00Circle\x00Diamond\x00Hexagon\x00Square\x00Star\x00Triangle\x00Cross\x00Moon\x00Raindrop\x00Kite\x00Pentagon\x00UpsideDownHouse\x00string_token1\x00string_token2\x00file_token1\x00number\x00comment\x00_space\x00_eol\x00filter\x00block\x00import\x00condition\x00action\x00continue\x00_equal_operator\x00_range_operator\x00quality\x00rarity\x00influence\x00sockets\x00colour\x00shape\x00string\x00file\x00_quantity\x00_id\x00_volume\x00_color\x00_icon_size\x00_font_size\x00filter_repeat1\x00block_repeat1\x00condition_repeat1\x00condition_repeat2\x00condition_repeat3\x00condition_repeat4\x00condition_repeat5\x00condition_repeat6\x00condition_repeat7\x00condition_repeat8\x00alpha\x00blue\x00class\x00enchantment\x00gem\x00green\x00id\x00modifier\x00red\x00size\x00type\x00volume\x00"
