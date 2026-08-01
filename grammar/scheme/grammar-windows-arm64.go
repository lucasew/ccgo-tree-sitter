// Code generated for windows/arm64 by 'ccgo -ignore-unsupported-alignment -ignore-unsupported-atomic-sizes -ignore-vector-functions --winapi-no-errno -ignore-link-errors -std=gnu11 -O0 -D__extension__= -D__forceinline=static inline -D__attribute__(...)= -D__declspec(x)= -D__cdecl= -D__stdcall= -D__fastcall= -D__thiscall= -D_cdecl= -D__restrict= -D__restrict__= -D__MINGW_EXTENSION= -D_X86INTRIN_H_INCLUDED -D_X86GPRINTRIN_H_INCLUDED -D_IMMINTRIN_H_INCLUDED -D_MMINTRIN_H_INCLUDED -D_XMMINTRIN_H_INCLUDED -D_EMMINTRIN_H_INCLUDED -D_PMMINTRIN_H_INCLUDED -D_MM3DNOW_H_INCLUDED -Dfunc=func_token -Dinterface=interface_token -Dselect=select_token -Dchan=chan_token -Dgo=go_token -Dmap=map_token -Dpackage=package_token -Dtype=type_token -Dvar=var_token -Dimport=import_token -Ddefer=defer_token -Dfallthrough=fallthrough_token -Drange=range_token -I C:\a\ccgo-tree-sitter\ccgo-tree-sitter\third-party\tree-sitter-scheme\src -I C:\a\ccgo-tree-sitter\ccgo-tree-sitter\third-party\tree-sitter-scheme -I C:\a\ccgo-tree-sitter\ccgo-tree-sitter\third-party -I C:\Users\runneradmin\.cache\workspaced\sources\github\7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2\lib\include -I C:\Users\runneradmin\.cache\workspaced\sources\github\7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2\lib\src C:\a\ccgo-tree-sitter\ccgo-tree-sitter\third-party\tree-sitter-scheme\src\parser.c -o grammar.go', DO NOT EDIT.

//go:build windows && arm64

package grammar_scheme

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
const FIELD_COUNT = 0
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
const LARGE_STATE_COUNT = 61
const MAX_ALIAS_SEQUENCE_LENGTH = 3
const MB_LEN_MAX = 1
const MINGW_HAS_DDK_H = 1
const MINGW_HAS_SECURE_API = 1
const PRODUCTION_ID_COUNT = 1
const PTRDIFF_MAX = "__PTRDIFF_MAX__"
const RAND_MAX = 0x7fff
const SIZE_MAX = "__SIZE_MAX__"
const STATE_COUNT = 143
const SYMBOL_COUNT = 56
const TOKEN_COUNT = 33
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

/*
 *  Lexer Macros
 */

/*
 *  Parse Table Macros
 */

type ts_symbol_identifiers = int32

const aux_sym__intertoken_token1 = 1
const aux_sym_comment_token1 = 2
const anon_sym_POUND_SEMI = 3
const anon_sym_POUND_BANG = 4
const aux_sym_directive_token1 = 5
const anon_sym_POUND_PIPE = 6
const aux_sym_block_comment_token1 = 7
const anon_sym_PIPE_POUND = 8
const sym_boolean = 9
const sym_number = 10
const sym_character = 11
const anon_sym_DQUOTE = 12
const aux_sym_string_token1 = 13
const sym_escape_sequence = 14
const sym_symbol = 15
const sym_keyword = 16
const anon_sym_LPAREN = 17
const anon_sym_RPAREN = 18
const anon_sym_LBRACK = 19
const anon_sym_RBRACK = 20
const anon_sym_LBRACE = 21
const anon_sym_RBRACE = 22
const anon_sym_SQUOTE = 23
const anon_sym_BQUOTE = 24
const anon_sym_POUND_SQUOTE = 25
const anon_sym_POUND_BQUOTE = 26
const anon_sym_COMMA = 27
const anon_sym_COMMA_AT = 28
const anon_sym_POUND_COMMA = 29
const anon_sym_POUND_COMMA_AT = 30
const anon_sym_POUND_LPAREN = 31
const anon_sym_POUNDvu8_LPAREN = 32
const sym_program = 33
const sym__token = 34
const sym__intertoken = 35
const sym_comment = 36
const sym_directive = 37
const sym_block_comment = 38
const sym__datum = 39
const sym_string = 40
const sym_list = 41
const sym_quote = 42
const sym_quasiquote = 43
const sym_syntax = 44
const sym_quasisyntax = 45
const sym_unquote = 46
const sym_unquote_splicing = 47
const sym_unsyntax = 48
const sym_unsyntax_splicing = 49
const sym_vector = 50
const sym_byte_vector = 51
const aux_sym_program_repeat1 = 52
const aux_sym_comment_repeat1 = 53
const aux_sym_block_comment_repeat1 = 54
const aux_sym_string_repeat1 = 55

var ts_symbol_names = [56]uintptr{
	0:  __ccgo_ts,
	1:  __ccgo_ts + 4,
	2:  __ccgo_ts + 23,
	3:  __ccgo_ts + 38,
	4:  __ccgo_ts + 41,
	5:  __ccgo_ts + 44,
	6:  __ccgo_ts + 61,
	7:  __ccgo_ts + 64,
	8:  __ccgo_ts + 85,
	9:  __ccgo_ts + 88,
	10: __ccgo_ts + 96,
	11: __ccgo_ts + 103,
	12: __ccgo_ts + 113,
	13: __ccgo_ts + 115,
	14: __ccgo_ts + 129,
	15: __ccgo_ts + 145,
	16: __ccgo_ts + 152,
	17: __ccgo_ts + 160,
	18: __ccgo_ts + 162,
	19: __ccgo_ts + 164,
	20: __ccgo_ts + 166,
	21: __ccgo_ts + 168,
	22: __ccgo_ts + 170,
	23: __ccgo_ts + 172,
	24: __ccgo_ts + 174,
	25: __ccgo_ts + 176,
	26: __ccgo_ts + 179,
	27: __ccgo_ts + 182,
	28: __ccgo_ts + 184,
	29: __ccgo_ts + 187,
	30: __ccgo_ts + 190,
	31: __ccgo_ts + 194,
	32: __ccgo_ts + 197,
	33: __ccgo_ts + 203,
	34: __ccgo_ts + 211,
	35: __ccgo_ts + 218,
	36: __ccgo_ts + 230,
	37: __ccgo_ts + 238,
	38: __ccgo_ts + 248,
	39: __ccgo_ts + 262,
	40: __ccgo_ts + 269,
	41: __ccgo_ts + 276,
	42: __ccgo_ts + 281,
	43: __ccgo_ts + 287,
	44: __ccgo_ts + 298,
	45: __ccgo_ts + 305,
	46: __ccgo_ts + 317,
	47: __ccgo_ts + 325,
	48: __ccgo_ts + 342,
	49: __ccgo_ts + 351,
	50: __ccgo_ts + 369,
	51: __ccgo_ts + 376,
	52: __ccgo_ts + 388,
	53: __ccgo_ts + 404,
	54: __ccgo_ts + 420,
	55: __ccgo_ts + 442,
}

var ts_symbol_map = [56]TSSymbol{
	1:  uint16(aux_sym__intertoken_token1),
	2:  uint16(aux_sym_comment_token1),
	3:  uint16(anon_sym_POUND_SEMI),
	4:  uint16(anon_sym_POUND_BANG),
	5:  uint16(aux_sym_directive_token1),
	6:  uint16(anon_sym_POUND_PIPE),
	7:  uint16(aux_sym_block_comment_token1),
	8:  uint16(anon_sym_PIPE_POUND),
	9:  uint16(sym_boolean),
	10: uint16(sym_number),
	11: uint16(sym_character),
	12: uint16(anon_sym_DQUOTE),
	13: uint16(aux_sym_string_token1),
	14: uint16(sym_escape_sequence),
	15: uint16(sym_symbol),
	16: uint16(sym_keyword),
	17: uint16(anon_sym_LPAREN),
	18: uint16(anon_sym_RPAREN),
	19: uint16(anon_sym_LBRACK),
	20: uint16(anon_sym_RBRACK),
	21: uint16(anon_sym_LBRACE),
	22: uint16(anon_sym_RBRACE),
	23: uint16(anon_sym_SQUOTE),
	24: uint16(anon_sym_BQUOTE),
	25: uint16(anon_sym_POUND_SQUOTE),
	26: uint16(anon_sym_POUND_BQUOTE),
	27: uint16(anon_sym_COMMA),
	28: uint16(anon_sym_COMMA_AT),
	29: uint16(anon_sym_POUND_COMMA),
	30: uint16(anon_sym_POUND_COMMA_AT),
	31: uint16(anon_sym_POUND_LPAREN),
	32: uint16(anon_sym_POUNDvu8_LPAREN),
	33: uint16(sym_program),
	34: uint16(sym__token),
	35: uint16(sym__intertoken),
	36: uint16(sym_comment),
	37: uint16(sym_directive),
	38: uint16(sym_block_comment),
	39: uint16(sym__datum),
	40: uint16(sym_string),
	41: uint16(sym_list),
	42: uint16(sym_quote),
	43: uint16(sym_quasiquote),
	44: uint16(sym_syntax),
	45: uint16(sym_quasisyntax),
	46: uint16(sym_unquote),
	47: uint16(sym_unquote_splicing),
	48: uint16(sym_unsyntax),
	49: uint16(sym_unsyntax_splicing),
	50: uint16(sym_vector),
	51: uint16(sym_byte_vector),
	52: uint16(aux_sym_program_repeat1),
	53: uint16(aux_sym_comment_repeat1),
	54: uint16(aux_sym_block_comment_repeat1),
	55: uint16(aux_sym_string_repeat1),
}

var ts_symbol_metadata = [56]TSSymbolMetadata{
	0: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	1: {},
	2: {},
	3: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	4: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	5: {},
	6: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	7: {},
	8: {
		Fvisible: libc.BoolUint8(true1 != 0),
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
	},
	13: {},
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
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	34: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	35: {
		Fnamed: libc.BoolUint8(true1 != 0),
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
	52: {},
	53: {},
	54: {},
	55: {},
}

var ts_alias_sequences = [1][3]TSSymbol{}

var ts_non_terminal_alias_map = [1]uint16_t{}

var ts_primary_state_ids = [143]TSStateId{
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
	12:  uint16(4),
	13:  uint16(13),
	14:  uint16(11),
	15:  uint16(13),
	16:  uint16(16),
	17:  uint16(5),
	18:  uint16(6),
	19:  uint16(8),
	20:  uint16(9),
	21:  uint16(10),
	22:  uint16(3),
	23:  uint16(16),
	24:  uint16(24),
	25:  uint16(25),
	26:  uint16(26),
	27:  uint16(24),
	28:  uint16(28),
	29:  uint16(29),
	30:  uint16(30),
	31:  uint16(31),
	32:  uint16(28),
	33:  uint16(29),
	34:  uint16(34),
	35:  uint16(35),
	36:  uint16(36),
	37:  uint16(37),
	38:  uint16(31),
	39:  uint16(39),
	40:  uint16(40),
	41:  uint16(41),
	42:  uint16(42),
	43:  uint16(43),
	44:  uint16(44),
	45:  uint16(39),
	46:  uint16(25),
	47:  uint16(26),
	48:  uint16(30),
	49:  uint16(34),
	50:  uint16(50),
	51:  uint16(35),
	52:  uint16(36),
	53:  uint16(37),
	54:  uint16(40),
	55:  uint16(41),
	56:  uint16(42),
	57:  uint16(43),
	58:  uint16(44),
	59:  uint16(50),
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
	94:  uint16(60),
	95:  uint16(93),
	96:  uint16(92),
	97:  uint16(74),
	98:  uint16(73),
	99:  uint16(80),
	100: uint16(75),
	101: uint16(81),
	102: uint16(76),
	103: uint16(83),
	104: uint16(84),
	105: uint16(86),
	106: uint16(88),
	107: uint16(89),
	108: uint16(90),
	109: uint16(65),
	110: uint16(62),
	111: uint16(63),
	112: uint16(64),
	113: uint16(91),
	114: uint16(66),
	115: uint16(61),
	116: uint16(67),
	117: uint16(68),
	118: uint16(69),
	119: uint16(70),
	120: uint16(77),
	121: uint16(79),
	122: uint16(82),
	123: uint16(85),
	124: uint16(71),
	125: uint16(78),
	126: uint16(72),
	127: uint16(87),
	128: uint16(128),
	129: uint16(129),
	130: uint16(130),
	131: uint16(130),
	132: uint16(128),
	133: uint16(128),
	134: uint16(130),
	135: uint16(135),
	136: uint16(136),
	137: uint16(137),
	138: uint16(136),
	139: uint16(135),
	140: uint16(89),
	141: uint16(79),
	142: uint16(142),
}

var aux_sym__intertoken_token1_character_set_1 = [9]TSCharacterRange{
	0: {
		Fstart: int32('\t'),
		Fend:   int32('\r'),
	},
	1: {
		Fstart: int32(' '),
		Fend:   int32(' '),
	},
	2: {
		Fstart: int32(0xa0),
		Fend:   int32(0xa0),
	},
	3: {
		Fstart: int32(0x1680),
		Fend:   int32(0x1680),
	},
	4: {
		Fstart: int32(0x2000),
		Fend:   int32(0x200a),
	},
	5: {
		Fstart: int32(0x2028),
		Fend:   int32(0x2029),
	},
	6: {
		Fstart: int32(0x202f),
		Fend:   int32(0x202f),
	},
	7: {
		Fstart: int32(0x205f),
		Fend:   int32(0x205f),
	},
	8: {
		Fstart: int32(0x3000),
		Fend:   int32(0x3000),
	},
}

var aux_sym_directive_token1_character_set_2 = [17]TSCharacterRange{
	0: {
		Fend: int32(0x08),
	},
	1: {
		Fstart: int32(0x0e),
		Fend:   int32(0x1f),
	},
	2: {
		Fstart: int32('!'),
		Fend:   int32('!'),
	},
	3: {
		Fstart: int32('$'),
		Fend:   int32('&'),
	},
	4: {
		Fstart: int32('*'),
		Fend:   int32('+'),
	},
	5: {
		Fstart: int32('-'),
		Fend:   int32(':'),
	},
	6: {
		Fstart: int32('<'),
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
		Fstart: int32('~'),
		Fend:   int32(0x9f),
	},
	10: {
		Fstart: int32(0xa1),
		Fend:   int32(0x167f),
	},
	11: {
		Fstart: int32(0x1681),
		Fend:   int32(0x1fff),
	},
	12: {
		Fstart: int32(0x200b),
		Fend:   int32(0x2027),
	},
	13: {
		Fstart: int32(0x202a),
		Fend:   int32(0x202e),
	},
	14: {
		Fstart: int32(0x2030),
		Fend:   int32(0x205e),
	},
	15: {
		Fstart: int32(0x2060),
		Fend:   int32(0x2fff),
	},
	16: {
		Fstart: int32(0x3001),
		Fend:   int32(0x10ffff),
	},
}

var sym_escape_sequence_character_set_1 = [9]TSCharacterRange{
	0: {
		Fstart: int32('\t'),
		Fend:   int32('\n'),
	},
	1: {
		Fstart: int32(' '),
		Fend:   int32(' '),
	},
	2: {
		Fstart: int32(0x85),
		Fend:   int32(0x85),
	},
	3: {
		Fstart: int32(0xa0),
		Fend:   int32(0xa0),
	},
	4: {
		Fstart: int32(0x1680),
		Fend:   int32(0x1680),
	},
	5: {
		Fstart: int32(0x2000),
		Fend:   int32(0x200a),
	},
	6: {
		Fstart: int32(0x202f),
		Fend:   int32(0x202f),
	},
	7: {
		Fstart: int32(0x205f),
		Fend:   int32(0x205f),
	},
	8: {
		Fstart: int32(0x3000),
		Fend:   int32(0x3000),
	},
}

var sym_escape_sequence_character_set_2 = [11]TSCharacterRange{
	0: {
		Fstart: int32('\t'),
		Fend:   int32('\n'),
	},
	1: {
		Fstart: int32('\r'),
		Fend:   int32('\r'),
	},
	2: {
		Fstart: int32(' '),
		Fend:   int32(' '),
	},
	3: {
		Fstart: int32(0x85),
		Fend:   int32(0x85),
	},
	4: {
		Fstart: int32(0xa0),
		Fend:   int32(0xa0),
	},
	5: {
		Fstart: int32(0x1680),
		Fend:   int32(0x1680),
	},
	6: {
		Fstart: int32(0x2000),
		Fend:   int32(0x200a),
	},
	7: {
		Fstart: int32(0x2028),
		Fend:   int32(0x2028),
	},
	8: {
		Fstart: int32(0x202f),
		Fend:   int32(0x202f),
	},
	9: {
		Fstart: int32(0x205f),
		Fend:   int32(0x205f),
	},
	10: {
		Fstart: int32(0x3000),
		Fend:   int32(0x3000),
	},
}

func ts_lex(tls *libc.TLS, lexer uintptr, state TSStateId) (r uint8) {
	var eof, result, skip, v4 uint8
	var half_size, i, i1, i10, i11, i12, i13, i14, i15, i16, i17, i18, i19, i2, i20, i21, i22, i23, i24, i25, i26, i27, i28, i29, i3, i30, i31, i4, i5, i6, i7, i8, i9, index, mid_index, size uint32_t
	var lookahead1, v3 int32_t
	var range_token, range_token1, v2 uintptr
	var v28 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = eof, half_size, i, i1, i10, i11, i12, i13, i14, i15, i16, i17, i18, i19, i2, i20, i21, i22, i23, i24, i25, i26, i27, i28, i29, i3, i30, i31, i4, i5, i6, i7, i8, i9, index, lookahead1, mid_index, range_token, range_token1, result, size, skip, v2, v28, v3, v4
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
			state = uint16(295)
			goto next_state
		}
		i = uint32(0)
		for {
			if !(uint64(i) < libc.Uint64FromInt64(52)/libc.Uint64FromInt64(2)) {
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
		v2 = uintptr(unsafe.Pointer(&aux_sym__intertoken_token1_character_set_1))
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
			state = uint16(296)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(303)
			goto next_state
		}
		return result
	case int32(1):
		if lookahead1 == int32('\n') {
			state = uint16(506)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(505)
			goto next_state
		}
		if lookahead1 == int32('X') {
			state = uint16(507)
			goto next_state
		}
		if lookahead1 == int32('x') {
			state = uint16(507)
			goto next_state
		}
		if lookahead1 == int32(0x85) || lookahead1 == int32(0x2028) {
			state = uint16(506)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_escape_sequence_character_set_2))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(11) - index
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
			state = uint16(504)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(503)
			goto next_state
		}
		return result
	case int32(2):
		if lookahead1 == int32('\r') {
			state = uint16(505)
			goto next_state
		}
		if lookahead1 == int32('\n') || lookahead1 == int32(0x85) || lookahead1 == int32(0x2028) {
			state = uint16(506)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_escape_sequence_character_set_2))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(11) - index
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
			state = uint16(2)
			goto next_state
		}
		return result
	case int32(3):
		i1 = uint32(0)
		for {
			if !(uint64(i1) < libc.Uint64FromInt64(104)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token1[i1]) == lookahead1 {
				state = map_token1[i1+uint32(1)]
				goto next_state
			}
			goto _14
		_14:
			;
			i1 = i1 + uint32(2)
		}
		return result
	case int32(4):
		if lookahead1 == int32('!') {
			state = uint16(299)
			goto next_state
		}
		if lookahead1 == int32(';') {
			state = uint16(298)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(302)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead1 == int32('"') {
			state = uint16(501)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(1)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(502)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead1 == int32('#') {
			state = uint16(305)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(304)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(303)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead1 == int32('#') {
			state = uint16(193)
			goto next_state
		}
		if lookahead1 == int32('.') {
			state = uint16(246)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(37)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(313)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead1 == int32('#') {
			state = uint16(9)
			goto next_state
		}
		if lookahead1 == int32('.') {
			state = uint16(15)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(260)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(177)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(8)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead1 == int32('#') {
			state = uint16(9)
			goto next_state
		}
		if lookahead1 == int32('.') {
			state = uint16(14)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(260)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(177)
			goto next_state
		}
		return result
	case int32(10):
		i2 = uint32(0)
		for {
			if !(uint64(i2) < libc.Uint64FromInt64(60)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token2[i2]) == lookahead1 {
				state = map_token2[i2+uint32(1)]
				goto next_state
			}
			goto _15
		_15:
			;
			i2 = i2 + uint32(2)
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(10)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead1 == int32('#') {
			state = uint16(9)
			goto next_state
		}
		if lookahead1 == int32('.') {
			state = uint16(13)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(260)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(255)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(184)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(11)
			goto next_state
		}
		return result
	case int32(12):
		i3 = uint32(0)
		for {
			if !(uint64(i3) < libc.Uint64FromInt64(52)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token3[i3]) == lookahead1 {
				state = map_token3[i3+uint32(1)]
				goto next_state
			}
			goto _16
		_16:
			;
			i3 = i3 + uint32(2)
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(12)
			goto next_state
		}
		return result
	case int32(13):
		if lookahead1 == int32('#') {
			state = uint16(14)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(255)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(184)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(13)
			goto next_state
		}
		return result
	case int32(14):
		if lookahead1 == int32('#') {
			state = uint16(14)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(177)
			goto next_state
		}
		return result
	case int32(15):
		if lookahead1 == int32('#') {
			state = uint16(14)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(177)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(15)
			goto next_state
		}
		return result
	case int32(16):
		if lookahead1 == int32('#') {
			state = uint16(16)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(228)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		return result
	case int32(17):
		if lookahead1 == int32('#') {
			state = uint16(16)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(228)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(17)
			goto next_state
		}
		return result
	case int32(18):
		if lookahead1 == int32('#') {
			state = uint16(18)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		return result
	case int32(19):
		if lookahead1 == int32('#') {
			state = uint16(18)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(19)
			goto next_state
		}
		return result
	case int32(20):
		if lookahead1 == int32('#') {
			state = uint16(18)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(20)
			goto next_state
		}
		return result
	case int32(21):
		if lookahead1 == int32('#') {
			state = uint16(18)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(21)
			goto next_state
		}
		return result
	case int32(22):
		if lookahead1 == int32('#') {
			state = uint16(18)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(22)
			goto next_state
		}
		return result
	case int32(23):
		if lookahead1 == int32('#') {
			state = uint16(4)
			goto next_state
		}
		if lookahead1 == int32(';') {
			state = uint16(297)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(118)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&aux_sym__intertoken_token1_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(9) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _20
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _20
	_20:
		if v4 != 0 {
			state = uint16(296)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32('"') && lookahead1 != int32('#') && (lookahead1 < int32('\'') || int32(')') < lookahead1) && lookahead1 != int32(',') && (lookahead1 < int32('[') || int32(']') < lookahead1) && lookahead1 != int32('`') && (lookahead1 < int32('{') || int32('}') < lookahead1) {
			state = uint16(301)
			goto next_state
		}
		return result
	case int32(24):
		if lookahead1 == int32('#') {
			state = uint16(236)
			goto next_state
		}
		if lookahead1 == int32('.') {
			state = uint16(246)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(37)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(313)
			goto next_state
		}
		return result
	case int32(25):
		if lookahead1 == int32('#') {
			state = uint16(25)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(240)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		return result
	case int32(26):
		if lookahead1 == int32('#') {
			state = uint16(25)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(240)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(26)
			goto next_state
		}
		return result
	case int32(27):
		if lookahead1 == int32('#') {
			state = uint16(27)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(287)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		return result
	case int32(28):
		if lookahead1 == int32('#') {
			state = uint16(27)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(287)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(28)
			goto next_state
		}
		return result
	case int32(29):
		if lookahead1 == int32('(') {
			state = uint16(620)
			goto next_state
		}
		return result
	case int32(30):
		if lookahead1 == int32('.') {
			state = uint16(140)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(255)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(255)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(184)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(30)
			goto next_state
		}
		return result
	case int32(31):
		if lookahead1 == int32('.') {
			state = uint16(67)
			goto next_state
		}
		return result
	case int32(32):
		if lookahead1 == int32('.') {
			state = uint16(142)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(255)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if lookahead1 == int32('E') || lookahead1 == int32('e') {
			state = uint16(177)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(32)
			goto next_state
		}
		return result
	case int32(33):
		i4 = uint32(0)
		for {
			if !(uint64(i4) < libc.Uint64FromInt64(56)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token4[i4]) == lookahead1 {
				state = map_token4[i4+uint32(1)]
				goto next_state
			}
			goto _21
		_21:
			;
			i4 = i4 + uint32(2)
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(34):
		if lookahead1 == int32('.') {
			state = uint16(246)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(37)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(313)
			goto next_state
		}
		return result
	case int32(35):
		if lookahead1 == int32('.') {
			state = uint16(66)
			goto next_state
		}
		return result
	case int32(36):
		if lookahead1 == int32('.') {
			state = uint16(68)
			goto next_state
		}
		return result
	case int32(37):
		if lookahead1 == int32('.') {
			state = uint16(250)
			goto next_state
		}
		if lookahead1 == int32('I') {
			state = uint16(216)
			goto next_state
		}
		if lookahead1 == int32('N') {
			state = uint16(190)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(440)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(85)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(320)
			goto next_state
		}
		return result
	case int32(38):
		if lookahead1 == int32('.') {
			state = uint16(253)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(39)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(331)
			goto next_state
		}
		return result
	case int32(39):
		if lookahead1 == int32('.') {
			state = uint16(253)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(331)
			goto next_state
		}
		return result
	case int32(40):
		if lookahead1 == int32('.') {
			state = uint16(254)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(8)
			goto next_state
		}
		return result
	case int32(41):
		if lookahead1 == int32('.') {
			state = uint16(257)
			goto next_state
		}
		if lookahead1 == int32('I') {
			state = uint16(211)
			goto next_state
		}
		if lookahead1 == int32('N') {
			state = uint16(188)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(111)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(83)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(333)
			goto next_state
		}
		return result
	case int32(42):
		if lookahead1 == int32('.') {
			state = uint16(257)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(41)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(333)
			goto next_state
		}
		return result
	case int32(43):
		if lookahead1 == int32('.') {
			state = uint16(258)
			goto next_state
		}
		if lookahead1 == int32('I') {
			state = uint16(212)
			goto next_state
		}
		if lookahead1 == int32('N') {
			state = uint16(189)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(439)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(84)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(10)
			goto next_state
		}
		return result
	case int32(44):
		if lookahead1 == int32('.') {
			state = uint16(261)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(162)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(122)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(388)
			goto next_state
		}
		return result
	case int32(45):
		if lookahead1 == int32('.') {
			state = uint16(261)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(44)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(388)
			goto next_state
		}
		return result
	case int32(46):
		if lookahead1 == int32('.') {
			state = uint16(262)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(447)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(121)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(30)
			goto next_state
		}
		return result
	case int32(47):
		if lookahead1 == int32('.') {
			state = uint16(265)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(162)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(122)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(335)
			goto next_state
		}
		return result
	case int32(48):
		if lookahead1 == int32('.') {
			state = uint16(265)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(47)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(335)
			goto next_state
		}
		return result
	case int32(49):
		if lookahead1 == int32('.') {
			state = uint16(266)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(447)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(121)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(11)
			goto next_state
		}
		return result
	case int32(50):
		if lookahead1 == int32('.') {
			state = uint16(267)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(51)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(391)
			goto next_state
		}
		return result
	case int32(51):
		if lookahead1 == int32('.') {
			state = uint16(267)
			goto next_state
		}
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(211)
			goto next_state
		}
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(188)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(391)
			goto next_state
		}
		return result
	case int32(52):
		if lookahead1 == int32('.') {
			state = uint16(268)
			goto next_state
		}
		if lookahead1 == int32('I') {
			state = uint16(212)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(474)
			goto next_state
		}
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(189)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(32)
			goto next_state
		}
		return result
	case int32(53):
		if lookahead1 == int32('.') {
			state = uint16(269)
			goto next_state
		}
		if lookahead1 == int32('I') {
			state = uint16(211)
			goto next_state
		}
		if lookahead1 == int32('N') {
			state = uint16(188)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(111)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(83)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(392)
			goto next_state
		}
		return result
	case int32(54):
		if lookahead1 == int32('.') {
			state = uint16(269)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(53)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(392)
			goto next_state
		}
		return result
	case int32(55):
		if lookahead1 == int32('.') {
			state = uint16(270)
			goto next_state
		}
		if lookahead1 == int32('I') {
			state = uint16(212)
			goto next_state
		}
		if lookahead1 == int32('N') {
			state = uint16(189)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(439)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(84)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(56):
		if lookahead1 == int32('.') {
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(57):
		if lookahead1 == int32('.') {
			state = uint16(70)
			goto next_state
		}
		return result
	case int32(58):
		if lookahead1 == int32('.') {
			state = uint16(71)
			goto next_state
		}
		return result
	case int32(59):
		if lookahead1 == int32('.') {
			state = uint16(72)
			goto next_state
		}
		return result
	case int32(60):
		if lookahead1 == int32('.') {
			state = uint16(73)
			goto next_state
		}
		return result
	case int32(61):
		if lookahead1 == int32('.') {
			state = uint16(74)
			goto next_state
		}
		return result
	case int32(62):
		if lookahead1 == int32('.') {
			state = uint16(75)
			goto next_state
		}
		return result
	case int32(63):
		if lookahead1 == int32('/') {
			state = uint16(229)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(63)
			goto next_state
		}
		return result
	case int32(64):
		if lookahead1 == int32('/') {
			state = uint16(241)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(64)
			goto next_state
		}
		return result
	case int32(65):
		if lookahead1 == int32('/') {
			state = uint16(289)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(65)
			goto next_state
		}
		return result
	case int32(66):
		if lookahead1 == int32('0') {
			state = uint16(310)
			goto next_state
		}
		return result
	case int32(67):
		if lookahead1 == int32('0') {
			state = uint16(414)
			goto next_state
		}
		return result
	case int32(68):
		if lookahead1 == int32('0') {
			state = uint16(138)
			goto next_state
		}
		return result
	case int32(69):
		if lookahead1 == int32('0') {
			state = uint16(418)
			goto next_state
		}
		return result
	case int32(70):
		if lookahead1 == int32('0') {
			state = uint16(428)
			goto next_state
		}
		return result
	case int32(71):
		if lookahead1 == int32('0') {
			state = uint16(429)
			goto next_state
		}
		return result
	case int32(72):
		if lookahead1 == int32('0') {
			state = uint16(415)
			goto next_state
		}
		return result
	case int32(73):
		if lookahead1 == int32('0') {
			state = uint16(419)
			goto next_state
		}
		return result
	case int32(74):
		if lookahead1 == int32('0') {
			state = uint16(416)
			goto next_state
		}
		return result
	case int32(75):
		if lookahead1 == int32('0') {
			state = uint16(420)
			goto next_state
		}
		return result
	case int32(76):
		if lookahead1 == int32('8') {
			state = uint16(29)
			goto next_state
		}
		return result
	case int32(77):
		if lookahead1 == int32(';') {
			state = uint16(116)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(77)
			goto next_state
		}
		return result
	case int32(78):
		if lookahead1 == int32(';') {
			state = uint16(503)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(78)
			goto next_state
		}
		return result
	case int32(79):
		if lookahead1 == int32(';') {
			state = uint16(117)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(79)
			goto next_state
		}
		return result
	case int32(80):
		if lookahead1 == int32(';') {
			state = uint16(118)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(80)
			goto next_state
		}
		return result
	case int32(81):
		if lookahead1 == int32('A') {
			state = uint16(194)
			goto next_state
		}
		if lookahead1 == int32('a') {
			state = uint16(88)
			goto next_state
		}
		return result
	case int32(82):
		if lookahead1 == int32('A') {
			state = uint16(207)
			goto next_state
		}
		if lookahead1 == int32('a') {
			state = uint16(107)
			goto next_state
		}
		return result
	case int32(83):
		if lookahead1 == int32('A') {
			state = uint16(209)
			goto next_state
		}
		if lookahead1 == int32('a') {
			state = uint16(109)
			goto next_state
		}
		return result
	case int32(84):
		if lookahead1 == int32('A') {
			state = uint16(210)
			goto next_state
		}
		if lookahead1 == int32('a') {
			state = uint16(110)
			goto next_state
		}
		return result
	case int32(85):
		if lookahead1 == int32('A') {
			state = uint16(213)
			goto next_state
		}
		if lookahead1 == int32('a') {
			state = uint16(112)
			goto next_state
		}
		return result
	case int32(86):
		if lookahead1 == int32('A') {
			state = uint16(214)
			goto next_state
		}
		if lookahead1 == int32('a') {
			state = uint16(113)
			goto next_state
		}
		return result
	case int32(87):
		if lookahead1 == int32('A') {
			state = uint16(215)
			goto next_state
		}
		if lookahead1 == int32('a') {
			state = uint16(114)
			goto next_state
		}
		return result
	case int32(88):
		if lookahead1 == int32('C') {
			state = uint16(195)
			goto next_state
		}
		if lookahead1 == int32('c') {
			state = uint16(195)
			goto next_state
		}
		return result
	case int32(89):
		if lookahead1 == int32('F') {
			state = uint16(31)
			goto next_state
		}
		if lookahead1 == int32('f') {
			state = uint16(56)
			goto next_state
		}
		return result
	case int32(90):
		if lookahead1 == int32('F') {
			state = uint16(35)
			goto next_state
		}
		if lookahead1 == int32('f') {
			state = uint16(35)
			goto next_state
		}
		return result
	case int32(91):
		if lookahead1 == int32('F') {
			state = uint16(36)
			goto next_state
		}
		if lookahead1 == int32('f') {
			state = uint16(36)
			goto next_state
		}
		return result
	case int32(92):
		if lookahead1 == int32('F') {
			state = uint16(57)
			goto next_state
		}
		if lookahead1 == int32('f') {
			state = uint16(58)
			goto next_state
		}
		return result
	case int32(93):
		if lookahead1 == int32('F') {
			state = uint16(59)
			goto next_state
		}
		if lookahead1 == int32('f') {
			state = uint16(60)
			goto next_state
		}
		return result
	case int32(94):
		if lookahead1 == int32('F') {
			state = uint16(61)
			goto next_state
		}
		if lookahead1 == int32('f') {
			state = uint16(62)
			goto next_state
		}
		return result
	case int32(95):
		if lookahead1 == int32('I') {
			state = uint16(208)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(108)
			goto next_state
		}
		return result
	case int32(96):
		if lookahead1 == int32('I') {
			state = uint16(212)
			goto next_state
		}
		if lookahead1 == int32('N') {
			state = uint16(189)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(439)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(84)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(255)
			goto next_state
		}
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(17)
			goto next_state
		}
		return result
	case int32(97):
		if lookahead1 == int32('I') {
			state = uint16(212)
			goto next_state
		}
		if lookahead1 == int32('N') {
			state = uint16(189)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(439)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(84)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(255)
			goto next_state
		}
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(63)
			goto next_state
		}
		return result
	case int32(98):
		if lookahead1 == int32('I') {
			state = uint16(212)
			goto next_state
		}
		if lookahead1 == int32('N') {
			state = uint16(189)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(439)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(84)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(255)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(26)
			goto next_state
		}
		return result
	case int32(99):
		if lookahead1 == int32('I') {
			state = uint16(212)
			goto next_state
		}
		if lookahead1 == int32('N') {
			state = uint16(189)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(439)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(84)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(255)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(64)
			goto next_state
		}
		return result
	case int32(100):
		if lookahead1 == int32('I') {
			state = uint16(212)
			goto next_state
		}
		if lookahead1 == int32('N') {
			state = uint16(189)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(439)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(84)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(255)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(28)
			goto next_state
		}
		return result
	case int32(101):
		if lookahead1 == int32('I') {
			state = uint16(212)
			goto next_state
		}
		if lookahead1 == int32('N') {
			state = uint16(189)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(439)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(84)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(255)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(65)
			goto next_state
		}
		return result
	case int32(102):
		if lookahead1 == int32('I') {
			state = uint16(212)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(474)
			goto next_state
		}
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(189)
			goto next_state
		}
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(63)
			goto next_state
		}
		return result
	case int32(103):
		if lookahead1 == int32('I') {
			state = uint16(212)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(474)
			goto next_state
		}
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(189)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(64)
			goto next_state
		}
		return result
	case int32(104):
		if lookahead1 == int32('I') {
			state = uint16(212)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(474)
			goto next_state
		}
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(189)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(65)
			goto next_state
		}
		return result
	case int32(105):
		if lookahead1 == int32('L') {
			state = uint16(203)
			goto next_state
		}
		if lookahead1 == int32('l') {
			state = uint16(95)
			goto next_state
		}
		return result
	case int32(106):
		i5 = uint32(0)
		for {
			if !(uint64(i5) < libc.Uint64FromInt64(64)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token5[i5]) == lookahead1 {
				state = map_token5[i5+uint32(1)]
				goto next_state
			}
			goto _22
		_22:
			;
			i5 = i5 + uint32(2)
		}
		if lookahead1 != 0 {
			state = uint16(484)
			goto next_state
		}
		return result
	case int32(107):
		if lookahead1 == int32('N') {
			state = uint16(31)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(56)
			goto next_state
		}
		return result
	case int32(108):
		if lookahead1 == int32('N') {
			state = uint16(195)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(195)
			goto next_state
		}
		return result
	case int32(109):
		if lookahead1 == int32('N') {
			state = uint16(35)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(35)
			goto next_state
		}
		return result
	case int32(110):
		if lookahead1 == int32('N') {
			state = uint16(36)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(36)
			goto next_state
		}
		return result
	case int32(111):
		if lookahead1 == int32('N') {
			state = uint16(198)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(90)
			goto next_state
		}
		return result
	case int32(112):
		if lookahead1 == int32('N') {
			state = uint16(57)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(58)
			goto next_state
		}
		return result
	case int32(113):
		if lookahead1 == int32('N') {
			state = uint16(59)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(60)
			goto next_state
		}
		return result
	case int32(114):
		if lookahead1 == int32('N') {
			state = uint16(61)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(62)
			goto next_state
		}
		return result
	case int32(115):
		if lookahead1 == int32('W') {
			state = uint16(205)
			goto next_state
		}
		if lookahead1 == int32('l') {
			state = uint16(484)
			goto next_state
		}
		if lookahead1 == int32('w') {
			state = uint16(105)
			goto next_state
		}
		return result
	case int32(116):
		if lookahead1 == int32('\\') {
			state = uint16(222)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(508)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(116)
			goto next_state
		}
		return result
	case int32(117):
		if lookahead1 == int32('\\') {
			state = uint16(223)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(602)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(117)
			goto next_state
		}
		return result
	case int32(118):
		if lookahead1 == int32('\\') {
			state = uint16(224)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(300)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(118)
			goto next_state
		}
		return result
	case int32(119):
		if lookahead1 == int32('a') {
			state = uint16(166)
			goto next_state
		}
		return result
	case int32(120):
		if lookahead1 == int32('a') {
			state = uint16(127)
			goto next_state
		}
		return result
	case int32(121):
		if lookahead1 == int32('a') {
			state = uint16(161)
			goto next_state
		}
		return result
	case int32(122):
		if lookahead1 == int32('a') {
			state = uint16(160)
			goto next_state
		}
		return result
	case int32(123):
		if lookahead1 == int32('b') {
			state = uint16(484)
			goto next_state
		}
		return result
	case int32(124):
		if lookahead1 == int32('b') {
			state = uint16(163)
			goto next_state
		}
		return result
	case int32(125):
		if lookahead1 == int32('c') {
			state = uint16(153)
			goto next_state
		}
		return result
	case int32(126):
		if lookahead1 == int32('c') {
			state = uint16(489)
			goto next_state
		}
		return result
	case int32(127):
		if lookahead1 == int32('c') {
			state = uint16(129)
			goto next_state
		}
		return result
	case int32(128):
		if lookahead1 == int32('d') {
			state = uint16(484)
			goto next_state
		}
		return result
	case int32(129):
		if lookahead1 == int32('e') {
			state = uint16(484)
			goto next_state
		}
		return result
	case int32(130):
		if lookahead1 == int32('e') {
			state = uint16(134)
			goto next_state
		}
		return result
	case int32(131):
		if lookahead1 == int32('e') {
			state = uint16(128)
			goto next_state
		}
		return result
	case int32(132):
		if lookahead1 == int32('e') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(133):
		if lookahead1 == int32('e') {
			state = uint16(131)
			goto next_state
		}
		return result
	case int32(134):
		if lookahead1 == int32('f') {
			state = uint16(133)
			goto next_state
		}
		return result
	case int32(135):
		if lookahead1 == int32('f') {
			state = uint16(35)
			goto next_state
		}
		return result
	case int32(136):
		if lookahead1 == int32('f') {
			state = uint16(36)
			goto next_state
		}
		return result
	case int32(137):
		if lookahead1 == int32('g') {
			state = uint16(129)
			goto next_state
		}
		return result
	case int32(138):
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		return result
	case int32(139):
		i6 = uint32(0)
		for {
			if !(uint64(i6) < libc.Uint64FromInt64(48)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token6[i6]) == lookahead1 {
				state = map_token6[i6+uint32(1)]
				goto next_state
			}
			goto _23
		_23:
			;
			i6 = i6 + uint32(2)
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(139)
			goto next_state
		}
		return result
	case int32(140):
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(255)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(184)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(141):
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(255)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(142):
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if lookahead1 == int32('E') || lookahead1 == int32('e') {
			state = uint16(177)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(142)
			goto next_state
		}
		return result
	case int32(143):
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(17)
			goto next_state
		}
		return result
	case int32(144):
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(144)
			goto next_state
		}
		return result
	case int32(145):
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(26)
			goto next_state
		}
		return result
	case int32(146):
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(146)
			goto next_state
		}
		return result
	case int32(147):
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(147)
			goto next_state
		}
		return result
	case int32(148):
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(28)
			goto next_state
		}
		return result
	case int32(149):
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(149)
			goto next_state
		}
		return result
	case int32(150):
		if lookahead1 == int32('i') {
			state = uint16(447)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(121)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(255)
			goto next_state
		}
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(63)
			goto next_state
		}
		return result
	case int32(151):
		if lookahead1 == int32('i') {
			state = uint16(447)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(121)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(255)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(64)
			goto next_state
		}
		return result
	case int32(152):
		if lookahead1 == int32('i') {
			state = uint16(447)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(121)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(255)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(65)
			goto next_state
		}
		return result
	case int32(153):
		if lookahead1 == int32('k') {
			state = uint16(168)
			goto next_state
		}
		return result
	case int32(154):
		if lookahead1 == int32('l') {
			state = uint16(484)
			goto next_state
		}
		return result
	case int32(155):
		if lookahead1 == int32('l') {
			state = uint16(132)
			goto next_state
		}
		return result
	case int32(156):
		if lookahead1 == int32('l') {
			state = uint16(494)
			goto next_state
		}
		return result
	case int32(157):
		if lookahead1 == int32('m') {
			state = uint16(484)
			goto next_state
		}
		return result
	case int32(158):
		if lookahead1 == int32('n') {
			state = uint16(484)
			goto next_state
		}
		return result
	case int32(159):
		if lookahead1 == int32('n') {
			state = uint16(130)
			goto next_state
		}
		return result
	case int32(160):
		if lookahead1 == int32('n') {
			state = uint16(35)
			goto next_state
		}
		return result
	case int32(161):
		if lookahead1 == int32('n') {
			state = uint16(36)
			goto next_state
		}
		return result
	case int32(162):
		if lookahead1 == int32('n') {
			state = uint16(135)
			goto next_state
		}
		return result
	case int32(163):
		if lookahead1 == int32('o') {
			state = uint16(174)
			goto next_state
		}
		return result
	case int32(164):
		if lookahead1 == int32('p') {
			state = uint16(120)
			goto next_state
		}
		return result
	case int32(165):
		if lookahead1 == int32('p') {
			state = uint16(129)
			goto next_state
		}
		return result
	case int32(166):
		if lookahead1 == int32('r') {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(167):
		if lookahead1 == int32('r') {
			state = uint16(158)
			goto next_state
		}
		return result
	case int32(168):
		if lookahead1 == int32('s') {
			state = uint16(164)
			goto next_state
		}
		return result
	case int32(169):
		if lookahead1 == int32('t') {
			state = uint16(484)
			goto next_state
		}
		return result
	case int32(170):
		if lookahead1 == int32('t') {
			state = uint16(173)
			goto next_state
		}
		return result
	case int32(171):
		if lookahead1 == int32('t') {
			state = uint16(129)
			goto next_state
		}
		return result
	case int32(172):
		if lookahead1 == int32('u') {
			state = uint16(76)
			goto next_state
		}
		return result
	case int32(173):
		if lookahead1 == int32('u') {
			state = uint16(167)
			goto next_state
		}
		return result
	case int32(174):
		if lookahead1 == int32('u') {
			state = uint16(169)
			goto next_state
		}
		return result
	case int32(175):
		if lookahead1 == int32('|') {
			state = uint16(117)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _27
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _27
		_27:
		}
		if v28 && v4 != 0 {
			state = uint16(603)
			goto next_state
		}
		return result
	case int32(176):
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(249)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(479)
			goto next_state
		}
		return result
	case int32(177):
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(255)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(147)
			goto next_state
		}
		return result
	case int32(178):
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(251)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(422)
			goto next_state
		}
		return result
	case int32(179):
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(274)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(427)
			goto next_state
		}
		return result
	case int32(180):
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(282)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(404)
			goto next_state
		}
		return result
	case int32(181):
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(256)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(421)
			goto next_state
		}
		return result
	case int32(182):
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(276)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(426)
			goto next_state
		}
		return result
	case int32(183):
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(263)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(461)
			goto next_state
		}
		return result
	case int32(184):
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(264)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(185):
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(281)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(405)
			goto next_state
		}
		return result
	case int32(186):
		if lookahead1 == int32('A') || lookahead1 == int32('a') {
			state = uint16(194)
			goto next_state
		}
		return result
	case int32(187):
		if lookahead1 == int32('A') || lookahead1 == int32('a') {
			state = uint16(207)
			goto next_state
		}
		return result
	case int32(188):
		if lookahead1 == int32('A') || lookahead1 == int32('a') {
			state = uint16(209)
			goto next_state
		}
		return result
	case int32(189):
		if lookahead1 == int32('A') || lookahead1 == int32('a') {
			state = uint16(210)
			goto next_state
		}
		return result
	case int32(190):
		if lookahead1 == int32('A') || lookahead1 == int32('a') {
			state = uint16(213)
			goto next_state
		}
		return result
	case int32(191):
		if lookahead1 == int32('A') || lookahead1 == int32('a') {
			state = uint16(214)
			goto next_state
		}
		return result
	case int32(192):
		if lookahead1 == int32('A') || lookahead1 == int32('a') {
			state = uint16(215)
			goto next_state
		}
		return result
	case int32(193):
		i7 = uint32(0)
		for {
			if !(uint64(i7) < libc.Uint64FromInt64(32)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token7[i7]) == lookahead1 {
				state = map_token7[i7+uint32(1)]
				goto next_state
			}
			goto _29
		_29:
			;
			i7 = i7 + uint32(2)
		}
		return result
	case int32(194):
		if lookahead1 == int32('C') || lookahead1 == int32('c') {
			state = uint16(195)
			goto next_state
		}
		return result
	case int32(195):
		if lookahead1 == int32('E') || lookahead1 == int32('e') {
			state = uint16(484)
			goto next_state
		}
		return result
	case int32(196):
		if lookahead1 == int32('E') || lookahead1 == int32('e') {
			state = uint16(307)
			goto next_state
		}
		return result
	case int32(197):
		if lookahead1 == int32('F') || lookahead1 == int32('f') {
			state = uint16(31)
			goto next_state
		}
		return result
	case int32(198):
		if lookahead1 == int32('F') || lookahead1 == int32('f') {
			state = uint16(35)
			goto next_state
		}
		return result
	case int32(199):
		if lookahead1 == int32('F') || lookahead1 == int32('f') {
			state = uint16(36)
			goto next_state
		}
		return result
	case int32(200):
		if lookahead1 == int32('F') || lookahead1 == int32('f') {
			state = uint16(57)
			goto next_state
		}
		return result
	case int32(201):
		if lookahead1 == int32('F') || lookahead1 == int32('f') {
			state = uint16(59)
			goto next_state
		}
		return result
	case int32(202):
		if lookahead1 == int32('F') || lookahead1 == int32('f') {
			state = uint16(61)
			goto next_state
		}
		return result
	case int32(203):
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(208)
			goto next_state
		}
		return result
	case int32(204):
		if lookahead1 == int32('L') || lookahead1 == int32('l') {
			state = uint16(219)
			goto next_state
		}
		return result
	case int32(205):
		if lookahead1 == int32('L') || lookahead1 == int32('l') {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(206):
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(197)
			goto next_state
		}
		return result
	case int32(207):
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(31)
			goto next_state
		}
		return result
	case int32(208):
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(195)
			goto next_state
		}
		return result
	case int32(209):
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(35)
			goto next_state
		}
		return result
	case int32(210):
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(36)
			goto next_state
		}
		return result
	case int32(211):
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(198)
			goto next_state
		}
		return result
	case int32(212):
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(199)
			goto next_state
		}
		return result
	case int32(213):
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(57)
			goto next_state
		}
		return result
	case int32(214):
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(59)
			goto next_state
		}
		return result
	case int32(215):
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(61)
			goto next_state
		}
		return result
	case int32(216):
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(200)
			goto next_state
		}
		return result
	case int32(217):
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(201)
			goto next_state
		}
		return result
	case int32(218):
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(202)
			goto next_state
		}
		return result
	case int32(219):
		if lookahead1 == int32('S') || lookahead1 == int32('s') {
			state = uint16(196)
			goto next_state
		}
		return result
	case int32(220):
		if lookahead1 == int32('U') || lookahead1 == int32('u') {
			state = uint16(196)
			goto next_state
		}
		return result
	case int32(221):
		if lookahead1 == int32('W') || lookahead1 == int32('w') {
			state = uint16(205)
			goto next_state
		}
		return result
	case int32(222):
		if lookahead1 == int32('X') || lookahead1 == int32('x') {
			state = uint16(283)
			goto next_state
		}
		if lookahead1 == int32('a') || lookahead1 == int32('b') || lookahead1 == int32('n') || lookahead1 == int32('r') || lookahead1 == int32('t') || lookahead1 == int32('|') {
			state = uint16(116)
			goto next_state
		}
		return result
	case int32(223):
		if lookahead1 == int32('X') || lookahead1 == int32('x') {
			state = uint16(288)
			goto next_state
		}
		if lookahead1 == int32('a') || lookahead1 == int32('b') || lookahead1 == int32('n') || lookahead1 == int32('r') || lookahead1 == int32('t') || lookahead1 == int32('|') {
			state = uint16(117)
			goto next_state
		}
		return result
	case int32(224):
		if lookahead1 == int32('X') || lookahead1 == int32('x') {
			state = uint16(293)
			goto next_state
		}
		if lookahead1 == int32('a') || lookahead1 == int32('b') || lookahead1 == int32('n') || lookahead1 == int32('r') || lookahead1 == int32('t') || lookahead1 == int32('|') {
			state = uint16(118)
			goto next_state
		}
		return result
	case int32(225):
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(358)
			goto next_state
		}
		return result
	case int32(226):
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(351)
			goto next_state
		}
		return result
	case int32(227):
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(365)
			goto next_state
		}
		return result
	case int32(228):
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(19)
			goto next_state
		}
		return result
	case int32(229):
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(144)
			goto next_state
		}
		return result
	case int32(230):
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(360)
			goto next_state
		}
		return result
	case int32(231):
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(367)
			goto next_state
		}
		return result
	case int32(232):
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(476)
			goto next_state
		}
		return result
	case int32(233):
		if lookahead1 == int32('E') || lookahead1 == int32('I') || lookahead1 == int32('e') || lookahead1 == int32('i') {
			state = uint16(399)
			goto next_state
		}
		return result
	case int32(234):
		if lookahead1 == int32('E') || lookahead1 == int32('I') || lookahead1 == int32('e') || lookahead1 == int32('i') {
			state = uint16(401)
			goto next_state
		}
		return result
	case int32(235):
		if lookahead1 == int32('E') || lookahead1 == int32('I') || lookahead1 == int32('e') || lookahead1 == int32('i') {
			state = uint16(403)
			goto next_state
		}
		return result
	case int32(236):
		if lookahead1 == int32('E') || lookahead1 == int32('I') || lookahead1 == int32('e') || lookahead1 == int32('i') {
			state = uint16(34)
			goto next_state
		}
		return result
	case int32(237):
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(371)
			goto next_state
		}
		return result
	case int32(238):
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(352)
			goto next_state
		}
		return result
	case int32(239):
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(378)
			goto next_state
		}
		return result
	case int32(240):
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(20)
			goto next_state
		}
		return result
	case int32(241):
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(146)
			goto next_state
		}
		return result
	case int32(242):
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(373)
			goto next_state
		}
		return result
	case int32(243):
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(380)
			goto next_state
		}
		return result
	case int32(244):
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(478)
			goto next_state
		}
		return result
	case int32(245):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(425)
			goto next_state
		}
		return result
	case int32(246):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(317)
			goto next_state
		}
		return result
	case int32(247):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(424)
			goto next_state
		}
		return result
	case int32(248):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(328)
			goto next_state
		}
		return result
	case int32(249):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(479)
			goto next_state
		}
		return result
	case int32(250):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(325)
			goto next_state
		}
		return result
	case int32(251):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(422)
			goto next_state
		}
		return result
	case int32(252):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(342)
			goto next_state
		}
		return result
	case int32(253):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(349)
			goto next_state
		}
		return result
	case int32(254):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(15)
			goto next_state
		}
		return result
	case int32(255):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(147)
			goto next_state
		}
		return result
	case int32(256):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(421)
			goto next_state
		}
		return result
	case int32(257):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(345)
			goto next_state
		}
		return result
	case int32(258):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(12)
			goto next_state
		}
		return result
	case int32(259):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(354)
			goto next_state
		}
		return result
	case int32(260):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(21)
			goto next_state
		}
		return result
	case int32(261):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(459)
			goto next_state
		}
		return result
	case int32(262):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(263):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(461)
			goto next_state
		}
		return result
	case int32(264):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(265):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(347)
			goto next_state
		}
		return result
	case int32(266):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(13)
			goto next_state
		}
		return result
	case int32(267):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(469)
			goto next_state
		}
		return result
	case int32(268):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(142)
			goto next_state
		}
		return result
	case int32(269):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(458)
			goto next_state
		}
		return result
	case int32(270):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(139)
			goto next_state
		}
		return result
	case int32(271):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(409)
			goto next_state
		}
		return result
	case int32(272):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(408)
			goto next_state
		}
		return result
	case int32(273):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(329)
			goto next_state
		}
		return result
	case int32(274):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(427)
			goto next_state
		}
		return result
	case int32(275):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(343)
			goto next_state
		}
		return result
	case int32(276):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(426)
			goto next_state
		}
		return result
	case int32(277):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(411)
			goto next_state
		}
		return result
	case int32(278):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(410)
			goto next_state
		}
		return result
	case int32(279):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(413)
			goto next_state
		}
		return result
	case int32(280):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(412)
			goto next_state
		}
		return result
	case int32(281):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(405)
			goto next_state
		}
		return result
	case int32(282):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(404)
			goto next_state
		}
		return result
	case int32(283):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(77)
			goto next_state
		}
		return result
	case int32(284):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(382)
			goto next_state
		}
		return result
	case int32(285):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(355)
			goto next_state
		}
		return result
	case int32(286):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(385)
			goto next_state
		}
		return result
	case int32(287):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(22)
			goto next_state
		}
		return result
	case int32(288):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(79)
			goto next_state
		}
		return result
	case int32(289):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(149)
			goto next_state
		}
		return result
	case int32(290):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(384)
			goto next_state
		}
		return result
	case int32(291):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(387)
			goto next_state
		}
		return result
	case int32(292):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(482)
			goto next_state
		}
		return result
	case int32(293):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(80)
			goto next_state
		}
		return result
	case int32(294):
		if eof != 0 {
			state = uint16(295)
			goto next_state
		}
		i8 = uint32(0)
		for {
			if !(uint64(i8) < libc.Uint64FromInt64(64)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token8[i8]) == lookahead1 {
				state = map_token8[i8+uint32(1)]
				goto next_state
			}
			goto _30
		_30:
			;
			i8 = i8 + uint32(2)
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(311)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&aux_sym__intertoken_token1_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(9) - index
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
			state = uint16(296)
			goto next_state
		}
		if lookahead1 != 0 && (lookahead1 < int32('[') || int32(']') < lookahead1) {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(295):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(ts_builtin_sym_end)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(296):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__intertoken_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		v2 = uintptr(unsafe.Pointer(&aux_sym__intertoken_token1_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(9) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _38
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _38
	_38:
		if v4 != 0 {
			state = uint16(296)
			goto next_state
		}
		return result
	case int32(297):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_comment_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 != 0 && lookahead1 != int32('\n') {
			state = uint16(297)
			goto next_state
		}
		return result
	case int32(298):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUND_SEMI)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(299):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUND_BANG)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(300):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_directive_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(301):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_directive_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _42
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _42
		_42:
		}
		if v28 && v4 != 0 {
			state = uint16(301)
			goto next_state
		}
		return result
	case int32(302):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUND_PIPE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(303):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_block_comment_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(304):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_block_comment_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(306)
			goto next_state
		}
		return result
	case int32(305):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_block_comment_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('|') {
			state = uint16(302)
			goto next_state
		}
		return result
	case int32(306):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PIPE_POUND)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(307):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_boolean)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(308):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_boolean)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('A') || lookahead1 == int32('a') {
			state = uint16(204)
			goto next_state
		}
		return result
	case int32(309):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_boolean)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('R') || lookahead1 == int32('r') {
			state = uint16(220)
			goto next_state
		}
		return result
	case int32(310):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(311):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i9 = uint32(0)
		for {
			if !(uint64(i9) < libc.Uint64FromInt64(68)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token9[i9]) == lookahead1 {
				state = map_token9[i9+uint32(1)]
				goto next_state
			}
			goto _44
		_44:
			;
			i9 = i9 + uint32(2)
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(311)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
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
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(312):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(312)
			goto next_state
		}
		if lookahead1 == int32('.') {
			state = uint16(316)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(248)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(38)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(40)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(185)
			goto next_state
		}
		return result
	case int32(313):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i10 = uint32(0)
		for {
			if !(uint64(i10) < libc.Uint64FromInt64(68)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token10[i10]) == lookahead1 {
				state = map_token10[i10+uint32(1)]
				goto next_state
			}
			goto _50
		_50:
			;
			i10 = i10 + uint32(2)
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(313)
			goto next_state
		}
		return result
	case int32(314):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(233)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(448)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(271)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(398)
			goto next_state
		}
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(321)
			goto next_state
		}
		return result
	case int32(315):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i11 = uint32(0)
		for {
			if !(uint64(i11) < libc.Uint64FromInt64(60)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token11[i11]) == lookahead1 {
				state = map_token11[i11+uint32(1)]
				goto next_state
			}
			goto _51
		_51:
			;
			i11 = i11 + uint32(2)
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(315)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _55
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _55
		_55:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(316):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(316)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(38)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(40)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(185)
			goto next_state
		}
		return result
	case int32(317):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i12 = uint32(0)
		for {
			if !(uint64(i12) < libc.Uint64FromInt64(60)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token12[i12]) == lookahead1 {
				state = map_token12[i12+uint32(1)]
				goto next_state
			}
			goto _57
		_57:
			;
			i12 = i12 + uint32(2)
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(317)
			goto next_state
		}
		return result
	case int32(318):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i13 = uint32(0)
		for {
			if !(uint64(i13) < libc.Uint64FromInt64(72)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token13[i13]) == lookahead1 {
				state = map_token13[i13+uint32(1)]
				goto next_state
			}
			goto _58
		_58:
			;
			i13 = i13 + uint32(2)
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(318)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _62
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _62
		_62:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(319):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(319)
			goto next_state
		}
		if lookahead1 == int32('.') {
			state = uint16(324)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(252)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(38)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(40)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(180)
			goto next_state
		}
		return result
	case int32(320):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i14 = uint32(0)
		for {
			if !(uint64(i14) < libc.Uint64FromInt64(72)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token14[i14]) == lookahead1 {
				state = map_token14[i14+uint32(1)]
				goto next_state
			}
			goto _64
		_64:
			;
			i14 = i14 + uint32(2)
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(320)
			goto next_state
		}
		return result
	case int32(321):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(322)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(225)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(448)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(96)
			goto next_state
		}
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(321)
			goto next_state
		}
		return result
	case int32(322):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(322)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(230)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(462)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(143)
			goto next_state
		}
		return result
	case int32(323):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i15 = uint32(0)
		for {
			if !(uint64(i15) < libc.Uint64FromInt64(64)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token15[i15]) == lookahead1 {
				state = map_token15[i15+uint32(1)]
				goto next_state
			}
			goto _65
		_65:
			;
			i15 = i15 + uint32(2)
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(323)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
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
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(324):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(324)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(38)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(40)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(180)
			goto next_state
		}
		return result
	case int32(325):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i16 = uint32(0)
		for {
			if !(uint64(i16) < libc.Uint64FromInt64(64)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token16[i16]) == lookahead1 {
				state = map_token16[i16+uint32(1)]
				goto next_state
			}
			goto _71
		_71:
			;
			i16 = i16 + uint32(2)
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(325)
			goto next_state
		}
		return result
	case int32(326):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(327)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(521)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(522)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(326)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _75
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _75
		_75:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(327):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(327)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(38)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(40)
			goto next_state
		}
		return result
	case int32(328):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(327)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(38)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(40)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(328)
			goto next_state
		}
		return result
	case int32(329):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(327)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(42)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(43)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(329)
			goto next_state
		}
		return result
	case int32(330):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i17 = uint32(0)
		for {
			if !(uint64(i17) < libc.Uint64FromInt64(56)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token17[i17]) == lookahead1 {
				state = map_token17[i17+uint32(1)]
				goto next_state
			}
			goto _77
		_77:
			;
			i17 = i17 + uint32(2)
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(330)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _81
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _81
		_81:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(331):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(332)
			goto next_state
		}
		if lookahead1 == int32('.') {
			state = uint16(349)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(259)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(176)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(331)
			goto next_state
		}
		return result
	case int32(332):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(332)
			goto next_state
		}
		if lookahead1 == int32('.') {
			state = uint16(348)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(259)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(176)
			goto next_state
		}
		return result
	case int32(333):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i18 = uint32(0)
		for {
			if !(uint64(i18) < libc.Uint64FromInt64(56)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token18[i18]) == lookahead1 {
				state = map_token18[i18+uint32(1)]
				goto next_state
			}
			goto _83
		_83:
			;
			i18 = i18 + uint32(2)
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(333)
			goto next_state
		}
		return result
	case int32(334):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(332)
			goto next_state
		}
		if lookahead1 == int32('.') {
			state = uint16(346)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(587)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(249)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(563)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(334)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _87
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _87
		_87:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(335):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(332)
			goto next_state
		}
		if lookahead1 == int32('.') {
			state = uint16(347)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(259)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(249)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(183)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(335)
			goto next_state
		}
		return result
	case int32(336):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(336)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(226)
			goto next_state
		}
		return result
	case int32(337):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(336)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(226)
			goto next_state
		}
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(337)
			goto next_state
		}
		return result
	case int32(338):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(339)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(227)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(448)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(96)
			goto next_state
		}
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(338)
			goto next_state
		}
		return result
	case int32(339):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(339)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(231)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(462)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(143)
			goto next_state
		}
		return result
	case int32(340):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(341)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(521)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(483)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(522)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(340)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _92
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _92
		_92:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(341):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(341)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(38)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(40)
			goto next_state
		}
		return result
	case int32(342):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(341)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(38)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(40)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(342)
			goto next_state
		}
		return result
	case int32(343):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(341)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(42)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(43)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(343)
			goto next_state
		}
		return result
	case int32(344):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i19 = uint32(0)
		for {
			if !(uint64(i19) < libc.Uint64FromInt64(48)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token19[i19]) == lookahead1 {
				state = map_token19[i19+uint32(1)]
				goto next_state
			}
			goto _94
		_94:
			;
			i19 = i19 + uint32(2)
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(344)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _98
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _98
		_98:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(345):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i20 = uint32(0)
		for {
			if !(uint64(i20) < libc.Uint64FromInt64(48)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token20[i20]) == lookahead1 {
				state = map_token20[i20+uint32(1)]
				goto next_state
			}
			goto _100
		_100:
			;
			i20 = i20 + uint32(2)
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(345)
			goto next_state
		}
		return result
	case int32(346):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(348)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(249)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(563)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(346)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _104
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _104
		_104:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(347):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(348)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(249)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(183)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(347)
			goto next_state
		}
		return result
	case int32(348):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(348)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(176)
			goto next_state
		}
		return result
	case int32(349):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(348)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(176)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(349)
			goto next_state
		}
		return result
	case int32(350):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(350)
			goto next_state
		}
		return result
	case int32(351):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(350)
			goto next_state
		}
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(351)
			goto next_state
		}
		return result
	case int32(352):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(350)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(352)
			goto next_state
		}
		return result
	case int32(353):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(350)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(353)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _109
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _109
		_109:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(354):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(350)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(354)
			goto next_state
		}
		return result
	case int32(355):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(350)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(355)
			goto next_state
		}
		return result
	case int32(356):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(357)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(237)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(449)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(98)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(356)
			goto next_state
		}
		return result
	case int32(357):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(357)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(242)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(463)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(145)
			goto next_state
		}
		return result
	case int32(358):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(359)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(448)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(96)
			goto next_state
		}
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(358)
			goto next_state
		}
		return result
	case int32(359):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(359)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(462)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(143)
			goto next_state
		}
		return result
	case int32(360):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(359)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(462)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(143)
			goto next_state
		}
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(360)
			goto next_state
		}
		return result
	case int32(361):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(361)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(238)
			goto next_state
		}
		return result
	case int32(362):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(361)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(238)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(363):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(364)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(239)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(449)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(98)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(363)
			goto next_state
		}
		return result
	case int32(364):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(364)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(243)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(463)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(145)
			goto next_state
		}
		return result
	case int32(365):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(366)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(448)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(96)
			goto next_state
		}
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(365)
			goto next_state
		}
		return result
	case int32(366):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(366)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(462)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(143)
			goto next_state
		}
		return result
	case int32(367):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(366)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(462)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(143)
			goto next_state
		}
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(367)
			goto next_state
		}
		return result
	case int32(368):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(234)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(449)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(277)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(400)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(356)
			goto next_state
		}
		return result
	case int32(369):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(370)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(284)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(450)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(100)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(369)
			goto next_state
		}
		return result
	case int32(370):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(370)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(290)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(464)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(148)
			goto next_state
		}
		return result
	case int32(371):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(372)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(449)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(98)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(371)
			goto next_state
		}
		return result
	case int32(372):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(372)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(463)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(145)
			goto next_state
		}
		return result
	case int32(373):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(372)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(463)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(145)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(373)
			goto next_state
		}
		return result
	case int32(374):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(374)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(285)
			goto next_state
		}
		return result
	case int32(375):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(374)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(285)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(375)
			goto next_state
		}
		return result
	case int32(376):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(377)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(286)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(450)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(100)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(376)
			goto next_state
		}
		return result
	case int32(377):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(377)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(291)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(464)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(148)
			goto next_state
		}
		return result
	case int32(378):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(379)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(449)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(98)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(378)
			goto next_state
		}
		return result
	case int32(379):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(379)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(463)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(145)
			goto next_state
		}
		return result
	case int32(380):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(379)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(463)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(145)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(380)
			goto next_state
		}
		return result
	case int32(381):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(235)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(450)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(279)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(402)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(369)
			goto next_state
		}
		return result
	case int32(382):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(383)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(450)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(100)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(382)
			goto next_state
		}
		return result
	case int32(383):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(383)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(464)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(148)
			goto next_state
		}
		return result
	case int32(384):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(383)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(464)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(148)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(384)
			goto next_state
		}
		return result
	case int32(385):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(386)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(450)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(100)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(385)
			goto next_state
		}
		return result
	case int32(386):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(386)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(464)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(148)
			goto next_state
		}
		return result
	case int32(387):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(386)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(464)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(148)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(387)
			goto next_state
		}
		return result
	case int32(388):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('.') {
			state = uint16(459)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(249)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(249)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(183)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(388)
			goto next_state
		}
		return result
	case int32(389):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('.') {
			state = uint16(468)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(597)
			goto next_state
		}
		if lookahead1 == int32('E') || lookahead1 == int32('e') {
			state = uint16(565)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(389)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _114
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _114
		_114:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(390):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i21 = uint32(0)
		for {
			if !(uint64(i21) < libc.Uint64FromInt64(52)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token21[i21]) == lookahead1 {
				state = map_token21[i21+uint32(1)]
				goto next_state
			}
			goto _116
		_116:
			;
			i21 = i21 + uint32(2)
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(390)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _120
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _120
		_120:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(391):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('.') {
			state = uint16(469)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(249)
			goto next_state
		}
		if lookahead1 == int32('E') || lookahead1 == int32('e') {
			state = uint16(176)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(391)
			goto next_state
		}
		return result
	case int32(392):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i22 = uint32(0)
		for {
			if !(uint64(i22) < libc.Uint64FromInt64(52)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token22[i22]) == lookahead1 {
				state = map_token22[i22+uint32(1)]
				goto next_state
			}
			goto _122
		_122:
			;
			i22 = i22 + uint32(2)
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(392)
			goto next_state
		}
		return result
	case int32(393):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('/') {
			state = uint16(232)
			goto next_state
		}
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(393)
			goto next_state
		}
		return result
	case int32(394):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('/') {
			state = uint16(244)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(394)
			goto next_state
		}
		return result
	case int32(395):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('/') {
			state = uint16(292)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(395)
			goto next_state
		}
		return result
	case int32(396):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('@') {
			state = uint16(521)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(483)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(247)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(522)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(396)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _126
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _126
		_126:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(397):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('@') {
			state = uint16(521)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(245)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(522)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(397)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _131
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _131
		_131:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(398):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i23 = uint32(0)
		for {
			if !(uint64(i23) < libc.Uint64FromInt64(40)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token23[i23]) == lookahead1 {
				state = map_token23[i23+uint32(1)]
				goto next_state
			}
			goto _133
		_133:
			;
			i23 = i23 + uint32(2)
		}
		return result
	case int32(399):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('@') {
			state = uint16(448)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(271)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(398)
			goto next_state
		}
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(321)
			goto next_state
		}
		return result
	case int32(400):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i24 = uint32(0)
		for {
			if !(uint64(i24) < libc.Uint64FromInt64(32)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token24[i24]) == lookahead1 {
				state = map_token24[i24+uint32(1)]
				goto next_state
			}
			goto _134
		_134:
			;
			i24 = i24 + uint32(2)
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(363)
			goto next_state
		}
		return result
	case int32(401):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('@') {
			state = uint16(449)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(277)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(400)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(356)
			goto next_state
		}
		return result
	case int32(402):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i25 = uint32(0)
		for {
			if !(uint64(i25) < libc.Uint64FromInt64(32)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token25[i25]) == lookahead1 {
				state = map_token25[i25+uint32(1)]
				goto next_state
			}
			goto _135
		_135:
			;
			i25 = i25 + uint32(2)
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(376)
			goto next_state
		}
		return result
	case int32(403):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('@') {
			state = uint16(450)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(279)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(402)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(369)
			goto next_state
		}
		return result
	case int32(404):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('@') {
			state = uint16(38)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(40)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(404)
			goto next_state
		}
		return result
	case int32(405):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('@') {
			state = uint16(38)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(40)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(405)
			goto next_state
		}
		return result
	case int32(406):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('@') {
			state = uint16(524)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(483)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(247)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(525)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(406)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _139
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _139
		_139:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(407):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('@') {
			state = uint16(524)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(245)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(525)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(407)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _144
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _144
		_144:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(408):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('@') {
			state = uint16(451)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(150)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(408)
			goto next_state
		}
		return result
	case int32(409):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('@') {
			state = uint16(451)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(150)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(409)
			goto next_state
		}
		return result
	case int32(410):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('@') {
			state = uint16(452)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(151)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(410)
			goto next_state
		}
		return result
	case int32(411):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('@') {
			state = uint16(452)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(151)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(411)
			goto next_state
		}
		return result
	case int32(412):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('@') {
			state = uint16(453)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(152)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(412)
			goto next_state
		}
		return result
	case int32(413):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('@') {
			state = uint16(453)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(152)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(413)
			goto next_state
		}
		return result
	case int32(414):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('@') {
			state = uint16(465)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(102)
			goto next_state
		}
		return result
	case int32(415):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('@') {
			state = uint16(466)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(103)
			goto next_state
		}
		return result
	case int32(416):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('@') {
			state = uint16(467)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(417):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('@') {
			state = uint16(526)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(483)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(528)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _149
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _149
		_149:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(418):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('@') {
			state = uint16(454)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(97)
			goto next_state
		}
		return result
	case int32(419):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('@') {
			state = uint16(455)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(99)
			goto next_state
		}
		return result
	case int32(420):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('@') {
			state = uint16(456)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(101)
			goto next_state
		}
		return result
	case int32(421):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('@') {
			state = uint16(42)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(247)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(43)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(421)
			goto next_state
		}
		return result
	case int32(422):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('@') {
			state = uint16(42)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(245)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(43)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(422)
			goto next_state
		}
		return result
	case int32(423):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('@') {
			state = uint16(530)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(483)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(531)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _154
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _154
		_154:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(424):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('@') {
			state = uint16(45)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(46)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(424)
			goto next_state
		}
		return result
	case int32(425):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('@') {
			state = uint16(45)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(46)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(425)
			goto next_state
		}
		return result
	case int32(426):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('@') {
			state = uint16(48)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(247)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(49)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(426)
			goto next_state
		}
		return result
	case int32(427):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('@') {
			state = uint16(48)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(245)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(49)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(427)
			goto next_state
		}
		return result
	case int32(428):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('@') {
			state = uint16(50)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(429):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('@') {
			state = uint16(54)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(310)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(55)
			goto next_state
		}
		return result
	case int32(430):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('I') {
			state = uint16(211)
			goto next_state
		}
		if lookahead1 == int32('N') {
			state = uint16(188)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(111)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(83)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(249)
			goto next_state
		}
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(337)
			goto next_state
		}
		return result
	case int32(431):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('I') {
			state = uint16(211)
			goto next_state
		}
		if lookahead1 == int32('N') {
			state = uint16(188)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(111)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(83)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(249)
			goto next_state
		}
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(393)
			goto next_state
		}
		return result
	case int32(432):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('I') {
			state = uint16(211)
			goto next_state
		}
		if lookahead1 == int32('N') {
			state = uint16(188)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(111)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(83)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(249)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(433):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('I') {
			state = uint16(211)
			goto next_state
		}
		if lookahead1 == int32('N') {
			state = uint16(188)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(111)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(83)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(249)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(394)
			goto next_state
		}
		return result
	case int32(434):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('I') {
			state = uint16(211)
			goto next_state
		}
		if lookahead1 == int32('N') {
			state = uint16(188)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(111)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(83)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(249)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(375)
			goto next_state
		}
		return result
	case int32(435):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('I') {
			state = uint16(211)
			goto next_state
		}
		if lookahead1 == int32('N') {
			state = uint16(188)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(111)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(83)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(249)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(395)
			goto next_state
		}
		return result
	case int32(436):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('N') {
			state = uint16(570)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(540)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _159
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _159
		_159:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(437):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('N') {
			state = uint16(197)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(89)
			goto next_state
		}
		return result
	case int32(438):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('N') {
			state = uint16(571)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(541)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _164
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _164
		_164:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(439):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('N') {
			state = uint16(199)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(91)
			goto next_state
		}
		return result
	case int32(440):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('N') {
			state = uint16(200)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(92)
			goto next_state
		}
		return result
	case int32(441):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('N') {
			state = uint16(201)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(93)
			goto next_state
		}
		return result
	case int32(442):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('N') {
			state = uint16(202)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(94)
			goto next_state
		}
		return result
	case int32(443):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('i') {
			state = uint16(162)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(122)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(249)
			goto next_state
		}
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(393)
			goto next_state
		}
		return result
	case int32(444):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('i') {
			state = uint16(162)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(122)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(249)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(394)
			goto next_state
		}
		return result
	case int32(445):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('i') {
			state = uint16(162)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(122)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(249)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(395)
			goto next_state
		}
		return result
	case int32(446):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('n') {
			state = uint16(549)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _169
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _169
		_169:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(447):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('n') {
			state = uint16(136)
			goto next_state
		}
		return result
	case int32(448):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('|') {
			state = uint16(249)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(430)
			goto next_state
		}
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(337)
			goto next_state
		}
		return result
	case int32(449):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('|') {
			state = uint16(249)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(432)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(450):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('|') {
			state = uint16(249)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(434)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(375)
			goto next_state
		}
		return result
	case int32(451):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('|') {
			state = uint16(249)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(443)
			goto next_state
		}
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(393)
			goto next_state
		}
		return result
	case int32(452):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('|') {
			state = uint16(249)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(444)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(394)
			goto next_state
		}
		return result
	case int32(453):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('|') {
			state = uint16(249)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(445)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(395)
			goto next_state
		}
		return result
	case int32(454):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('|') {
			state = uint16(249)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(431)
			goto next_state
		}
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(393)
			goto next_state
		}
		return result
	case int32(455):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('|') {
			state = uint16(249)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(433)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(394)
			goto next_state
		}
		return result
	case int32(456):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('|') {
			state = uint16(249)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(435)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(395)
			goto next_state
		}
		return result
	case int32(457):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i26 = uint32(0)
		for {
			if !(uint64(i26) < libc.Uint64FromInt64(44)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token26[i26]) == lookahead1 {
				state = map_token26[i26+uint32(1)]
				goto next_state
			}
			goto _171
		_171:
			;
			i26 = i26 + uint32(2)
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(457)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _175
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _175
		_175:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(458):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i27 = uint32(0)
		for {
			if !(uint64(i27) < libc.Uint64FromInt64(44)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token27[i27]) == lookahead1 {
				state = map_token27[i27+uint32(1)]
				goto next_state
			}
			goto _177
		_177:
			;
			i27 = i27 + uint32(2)
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(458)
			goto next_state
		}
		return result
	case int32(459):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('|') {
			state = uint16(249)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(183)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(459)
			goto next_state
		}
		return result
	case int32(460):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('|') {
			state = uint16(249)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(460)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _181
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _181
		_181:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(461):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('|') {
			state = uint16(249)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(461)
			goto next_state
		}
		return result
	case int32(462):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(475)
			goto next_state
		}
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(337)
			goto next_state
		}
		return result
	case int32(463):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(477)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(464):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(481)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(375)
			goto next_state
		}
		return result
	case int32(465):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(470)
			goto next_state
		}
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(393)
			goto next_state
		}
		return result
	case int32(466):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(471)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(394)
			goto next_state
		}
		return result
	case int32(467):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(472)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(395)
			goto next_state
		}
		return result
	case int32(468):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('E') || lookahead1 == int32('e') {
			state = uint16(565)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(468)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _186
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _186
		_186:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(469):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('E') || lookahead1 == int32('e') {
			state = uint16(176)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(469)
			goto next_state
		}
		return result
	case int32(470):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(211)
			goto next_state
		}
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(188)
			goto next_state
		}
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(393)
			goto next_state
		}
		return result
	case int32(471):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(211)
			goto next_state
		}
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(188)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(394)
			goto next_state
		}
		return result
	case int32(472):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(211)
			goto next_state
		}
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(188)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(395)
			goto next_state
		}
		return result
	case int32(473):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(571)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _191
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _191
		_191:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(474):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(199)
			goto next_state
		}
		return result
	case int32(475):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(337)
			goto next_state
		}
		return result
	case int32(476):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(476)
			goto next_state
		}
		return result
	case int32(477):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(478):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(478)
			goto next_state
		}
		return result
	case int32(479):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(479)
			goto next_state
		}
		return result
	case int32(480):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(480)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _196
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _196
		_196:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(481):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(375)
			goto next_state
		}
		return result
	case int32(482):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(482)
			goto next_state
		}
		return result
	case int32(483):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _201
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _201
		_201:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(484):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_character)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(485):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_character)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('E') {
			state = uint16(221)
			goto next_state
		}
		if lookahead1 == int32('e') {
			state = uint16(115)
			goto next_state
		}
		if lookahead1 == int32('u') {
			state = uint16(156)
			goto next_state
		}
		return result
	case int32(486):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_character)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('P') {
			state = uint16(186)
			goto next_state
		}
		if lookahead1 == int32('p') {
			state = uint16(81)
			goto next_state
		}
		return result
	case int32(487):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_character)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('a') {
			state = uint16(125)
			goto next_state
		}
		if lookahead1 == int32('e') {
			state = uint16(154)
			goto next_state
		}
		return result
	case int32(488):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_character)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('a') {
			state = uint16(137)
			goto next_state
		}
		return result
	case int32(489):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_character)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('a') {
			state = uint16(165)
			goto next_state
		}
		return result
	case int32(490):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_character)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('a') {
			state = uint16(123)
			goto next_state
		}
		return result
	case int32(491):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_character)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('e') {
			state = uint16(170)
			goto next_state
		}
		if lookahead1 == int32('u') {
			state = uint16(124)
			goto next_state
		}
		return result
	case int32(492):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_character)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('e') {
			state = uint16(155)
			goto next_state
		}
		return result
	case int32(493):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_character)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('i') {
			state = uint16(159)
			goto next_state
		}
		if lookahead1 == int32('s') {
			state = uint16(484)
			goto next_state
		}
		return result
	case int32(494):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_character)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('l') {
			state = uint16(484)
			goto next_state
		}
		return result
	case int32(495):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_character)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('l') {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(496):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_character)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('s') {
			state = uint16(126)
			goto next_state
		}
		return result
	case int32(497):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_character)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('t') {
			state = uint16(490)
			goto next_state
		}
		return result
	case int32(498):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_character)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('E') || lookahead1 == int32('e') {
			state = uint16(221)
			goto next_state
		}
		return result
	case int32(499):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_character)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('P') || lookahead1 == int32('p') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(500):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_character)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(500)
			goto next_state
		}
		return result
	case int32(501):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(502):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 != 0 && lookahead1 != int32('"') && lookahead1 != int32('\\') {
			state = uint16(502)
			goto next_state
		}
		return result
	case int32(503):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_escape_sequence)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(504):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_escape_sequence)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\r') {
			state = uint16(505)
			goto next_state
		}
		if lookahead1 == int32('\n') || lookahead1 == int32(0x85) || lookahead1 == int32(0x2028) {
			state = uint16(506)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_escape_sequence_character_set_2))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(11) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _206
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _206
	_206:
		if v4 != 0 {
			state = uint16(2)
			goto next_state
		}
		return result
	case int32(505):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_escape_sequence)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\n') || lookahead1 == int32(0x85) {
			state = uint16(506)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_escape_sequence_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(9) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _210
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _210
	_210:
		if v4 != 0 {
			state = uint16(506)
			goto next_state
		}
		return result
	case int32(506):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_escape_sequence)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		v2 = uintptr(unsafe.Pointer(&sym_escape_sequence_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(9) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _214
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _214
	_214:
		;
		if v4 != 0 && lookahead1 != int32('\n') && lookahead1 != int32(0x85) {
			state = uint16(506)
			goto next_state
		}
		return result
	case int32(507):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_escape_sequence)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(78)
			goto next_state
		}
		return result
	case int32(508):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(509):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i28 = uint32(0)
		for {
			if !(uint64(i28) < libc.Uint64FromInt64(60)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token28[i28]) == lookahead1 {
				state = map_token28[i28+uint32(1)]
				goto next_state
			}
			goto _215
		_215:
			;
			i28 = i28 + uint32(2)
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(509)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _219
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _219
		_219:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(510):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(9)
			goto next_state
		}
		if lookahead1 == int32('.') {
			state = uint16(512)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(589)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(483)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(255)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(564)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(510)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _224
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _224
		_224:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(511):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i29 = uint32(0)
		for {
			if !(uint64(i29) < libc.Uint64FromInt64(52)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token29[i29]) == lookahead1 {
				state = map_token29[i29+uint32(1)]
				goto next_state
			}
			goto _226
		_226:
			;
			i29 = i29 + uint32(2)
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(511)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _230
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _230
		_230:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(512):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(14)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(483)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(255)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(564)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(512)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _235
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _235
		_235:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(513):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(18)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(483)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(513)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _240
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _240
		_240:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(514):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('.') {
			state = uint16(534)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _245
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _245
		_245:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(515):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('.') {
			state = uint16(553)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(598)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(483)
			goto next_state
		}
		if lookahead1 == int32('E') || lookahead1 == int32('e') {
			state = uint16(566)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(515)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _250
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _250
		_250:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(516):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i30 = uint32(0)
		for {
			if !(uint64(i30) < libc.Uint64FromInt64(56)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token30[i30]) == lookahead1 {
				state = map_token30[i30+uint32(1)]
				goto next_state
			}
			goto _252
		_252:
			;
			i30 = i30 + uint32(2)
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(516)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _256
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _256
		_256:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(517):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('.') {
			state = uint16(580)
			goto next_state
		}
		if lookahead1 == int32('I') {
			state = uint16(573)
			goto next_state
		}
		if lookahead1 == int32('N') {
			state = uint16(567)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(436)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(537)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(318)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _261
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _261
		_261:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(518):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('.') {
			state = uint16(535)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _266
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _266
		_266:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(519):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('.') {
			state = uint16(533)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _271
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _271
		_271:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(520):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('.') {
			state = uint16(585)
			goto next_state
		}
		if lookahead1 == int32('I') {
			state = uint16(578)
			goto next_state
		}
		if lookahead1 == int32('N') {
			state = uint16(569)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(546)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(539)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(330)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _276
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _276
		_276:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(521):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('.') {
			state = uint16(585)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(520)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(330)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _281
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _281
		_281:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(522):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('.') {
			state = uint16(586)
			goto next_state
		}
		if lookahead1 == int32('I') {
			state = uint16(577)
			goto next_state
		}
		if lookahead1 == int32('N') {
			state = uint16(568)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(438)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(538)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(509)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _286
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _286
		_286:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(523):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('.') {
			state = uint16(591)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(558)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(548)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(334)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _291
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _291
		_291:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(524):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('.') {
			state = uint16(591)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(523)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(334)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _296
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _296
		_296:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(525):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('.') {
			state = uint16(592)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(446)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(547)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(510)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _301
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _301
		_301:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(526):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('.') {
			state = uint16(593)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(527)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(389)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _306
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _306
		_306:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(527):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('.') {
			state = uint16(593)
			goto next_state
		}
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(578)
			goto next_state
		}
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(569)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(389)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _311
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _311
		_311:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(528):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('.') {
			state = uint16(594)
			goto next_state
		}
		if lookahead1 == int32('I') {
			state = uint16(577)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(473)
			goto next_state
		}
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(568)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(515)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _316
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _316
		_316:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(529):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('.') {
			state = uint16(595)
			goto next_state
		}
		if lookahead1 == int32('I') {
			state = uint16(578)
			goto next_state
		}
		if lookahead1 == int32('N') {
			state = uint16(569)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(546)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(539)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(390)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _321
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _321
		_321:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(530):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('.') {
			state = uint16(595)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(529)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(390)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _326
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _326
		_326:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(531):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('.') {
			state = uint16(596)
			goto next_state
		}
		if lookahead1 == int32('I') {
			state = uint16(577)
			goto next_state
		}
		if lookahead1 == int32('N') {
			state = uint16(568)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(438)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(538)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(516)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _331
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _331
		_331:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(532):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('.') {
			state = uint16(536)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _336
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _336
		_336:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(533):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('0') {
			state = uint16(483)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _341
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _341
		_341:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(534):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('0') {
			state = uint16(417)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _346
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _346
		_346:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(535):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('0') {
			state = uint16(555)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _351
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _351
		_351:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(536):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('0') {
			state = uint16(423)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _356
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _356
		_356:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(537):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('A') {
			state = uint16(574)
			goto next_state
		}
		if lookahead1 == int32('a') {
			state = uint16(543)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _361
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _361
		_361:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(538):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('A') {
			state = uint16(575)
			goto next_state
		}
		if lookahead1 == int32('a') {
			state = uint16(544)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _366
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _366
		_366:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(539):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('A') {
			state = uint16(576)
			goto next_state
		}
		if lookahead1 == int32('a') {
			state = uint16(545)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _371
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _371
		_371:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(540):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('F') {
			state = uint16(514)
			goto next_state
		}
		if lookahead1 == int32('f') {
			state = uint16(532)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _376
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _376
		_376:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(541):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('F') {
			state = uint16(518)
			goto next_state
		}
		if lookahead1 == int32('f') {
			state = uint16(518)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _381
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _381
		_381:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(542):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('F') {
			state = uint16(519)
			goto next_state
		}
		if lookahead1 == int32('f') {
			state = uint16(519)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _386
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _386
		_386:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(543):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('N') {
			state = uint16(514)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(532)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _391
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _391
		_391:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(544):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('N') {
			state = uint16(518)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(518)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _396
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _396
		_396:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(545):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('N') {
			state = uint16(519)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(519)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _401
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _401
		_401:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(546):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('N') {
			state = uint16(572)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(542)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _406
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _406
		_406:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(547):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('a') {
			state = uint16(556)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _411
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _411
		_411:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(548):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('a') {
			state = uint16(557)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _416
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _416
		_416:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(549):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('f') {
			state = uint16(518)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _421
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _421
		_421:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(550):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('f') {
			state = uint16(519)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _426
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _426
		_426:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(551):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i31 = uint32(0)
		for {
			if !(uint64(i31) < libc.Uint64FromInt64(48)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token31[i31]) == lookahead1 {
				state = map_token31[i31+uint32(1)]
				goto next_state
			}
			goto _428
		_428:
			;
			i31 = i31 + uint32(2)
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(551)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _432
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _432
		_432:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(552):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('i') {
			state = uint16(483)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(255)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(552)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _437
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _437
		_437:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(553):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('i') {
			state = uint16(483)
			goto next_state
		}
		if lookahead1 == int32('E') || lookahead1 == int32('e') {
			state = uint16(566)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(553)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _442
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _442
		_442:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(554):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('i') {
			state = uint16(483)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(554)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _447
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _447
		_447:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(555):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('i') {
			state = uint16(483)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _452
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _452
		_452:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(556):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('n') {
			state = uint16(518)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _457
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _457
		_457:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(557):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('n') {
			state = uint16(519)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _462
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _462
		_462:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(558):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('n') {
			state = uint16(550)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _467
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _467
		_467:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(559):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(582)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(397)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _472
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _472
		_472:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(560):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(599)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(407)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _477
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _477
		_477:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(561):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(584)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(396)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _482
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _482
		_482:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(562):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(600)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(406)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _487
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _487
		_487:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(563):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(588)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(460)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _492
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _492
		_492:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(564):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(590)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(552)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _497
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _497
		_497:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(565):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(597)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(480)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _502
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _502
		_502:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(566):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(598)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(554)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _507
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _507
		_507:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(567):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('A') || lookahead1 == int32('a') {
			state = uint16(574)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _512
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _512
		_512:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(568):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('A') || lookahead1 == int32('a') {
			state = uint16(575)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _517
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _517
		_517:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(569):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('A') || lookahead1 == int32('a') {
			state = uint16(576)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _522
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _522
		_522:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(570):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('F') || lookahead1 == int32('f') {
			state = uint16(514)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _527
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _527
		_527:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(571):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('F') || lookahead1 == int32('f') {
			state = uint16(518)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _532
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _532
		_532:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(572):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('F') || lookahead1 == int32('f') {
			state = uint16(519)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _537
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _537
		_537:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(573):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(570)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _542
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _542
		_542:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(574):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(514)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _547
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _547
		_547:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(575):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(518)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _552
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _552
		_552:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(576):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(519)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _557
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _557
		_557:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(577):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(571)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _562
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _562
		_562:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(578):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(572)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _567
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _567
		_567:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(579):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(315)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _572
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _572
		_572:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(580):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(323)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _577
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _577
		_577:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(581):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(326)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _582
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _582
		_582:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(582):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(397)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _587
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _587
		_587:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(583):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(340)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _592
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _592
		_592:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(584):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(396)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _597
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _597
		_597:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(585):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(344)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _602
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _602
		_602:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(586):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(511)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _607
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _607
		_607:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(587):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(353)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _612
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _612
		_612:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(588):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(460)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _617
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _617
		_617:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(589):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(513)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _622
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _622
		_622:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(590):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(552)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _627
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _627
		_627:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(591):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(346)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _632
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _632
		_632:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(592):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(512)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _637
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _637
		_637:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(593):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(468)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _642
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _642
		_642:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(594):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(553)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _647
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _647
		_647:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(595):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(457)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _652
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _652
		_652:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(596):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(551)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _657
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _657
		_657:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(597):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(480)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _662
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _662
		_662:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(598):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(554)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _667
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _667
		_667:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(599):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(407)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _672
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _672
		_672:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(600):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(406)
			goto next_state
		}
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _677
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _677
		_677:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(601):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _682
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _682
		_682:
		}
		if v28 && v4 != 0 {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(602):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_keyword)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(603):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_keyword)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if v28 = !(eof != 0); v28 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_directive_token1_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(17) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _687
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _687
		_687:
		}
		if v28 && v4 != 0 {
			state = uint16(603)
			goto next_state
		}
		return result
	case int32(604):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(605):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(606):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(607):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(608):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(609):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(610):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(611):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(612):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUND_SQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(613):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUND_BQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(614):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COMMA)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(615):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COMMA)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('@') {
			state = uint16(616)
			goto next_state
		}
		return result
	case int32(616):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COMMA_AT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(617):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUND_COMMA)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('@') {
			state = uint16(618)
			goto next_state
		}
		return result
	case int32(618):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUND_COMMA_AT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(619):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUND_LPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(620):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUNDvu8_LPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var map_token = [26]uint16_t{
	0:  uint16('"'),
	1:  uint16(501),
	2:  uint16('#'),
	3:  uint16(305),
	4:  uint16('\''),
	5:  uint16(610),
	6:  uint16('('),
	7:  uint16(604),
	8:  uint16(')'),
	9:  uint16(605),
	10: uint16(','),
	11: uint16(614),
	12: uint16(';'),
	13: uint16(297),
	14: uint16('['),
	15: uint16(606),
	16: uint16(']'),
	17: uint16(607),
	18: uint16('`'),
	19: uint16(611),
	20: uint16('{'),
	21: uint16(608),
	22: uint16('|'),
	23: uint16(304),
	24: uint16('}'),
	25: uint16(609),
}

var map_token1 = [52]uint16_t{
	0:  uint16('!'),
	1:  uint16(299),
	2:  uint16('\''),
	3:  uint16(612),
	4:  uint16('('),
	5:  uint16(619),
	6:  uint16(','),
	7:  uint16(617),
	8:  uint16(':'),
	9:  uint16(175),
	10: uint16(';'),
	11: uint16(298),
	12: uint16('\\'),
	13: uint16(106),
	14: uint16('`'),
	15: uint16(613),
	16: uint16('v'),
	17: uint16(172),
	18: uint16('|'),
	19: uint16(302),
	20: uint16('B'),
	21: uint16(314),
	22: uint16('b'),
	23: uint16(314),
	24: uint16('D'),
	25: uint16(24),
	26: uint16('d'),
	27: uint16(24),
	28: uint16('F'),
	29: uint16(308),
	30: uint16('f'),
	31: uint16(308),
	32: uint16('O'),
	33: uint16(368),
	34: uint16('o'),
	35: uint16(368),
	36: uint16('T'),
	37: uint16(309),
	38: uint16('t'),
	39: uint16(309),
	40: uint16('X'),
	41: uint16(381),
	42: uint16('x'),
	43: uint16(381),
	44: uint16('E'),
	45: uint16(7),
	46: uint16('I'),
	47: uint16(7),
	48: uint16('e'),
	49: uint16(7),
	50: uint16('i'),
	51: uint16(7),
}

var map_token2 = [30]uint16_t{
	0:  uint16('#'),
	1:  uint16(9),
	2:  uint16('.'),
	3:  uint16(12),
	4:  uint16('/'),
	5:  uint16(260),
	6:  uint16('i'),
	7:  uint16(310),
	8:  uint16('|'),
	9:  uint16(255),
	10: uint16('E'),
	11: uint16(184),
	12: uint16('e'),
	13: uint16(184),
	14: uint16('D'),
	15: uint16(184),
	16: uint16('F'),
	17: uint16(184),
	18: uint16('L'),
	19: uint16(184),
	20: uint16('S'),
	21: uint16(184),
	22: uint16('d'),
	23: uint16(184),
	24: uint16('f'),
	25: uint16(184),
	26: uint16('l'),
	27: uint16(184),
	28: uint16('s'),
	29: uint16(184),
}

var map_token3 = [26]uint16_t{
	0:  uint16('#'),
	1:  uint16(14),
	2:  uint16('i'),
	3:  uint16(310),
	4:  uint16('|'),
	5:  uint16(255),
	6:  uint16('E'),
	7:  uint16(184),
	8:  uint16('e'),
	9:  uint16(184),
	10: uint16('D'),
	11: uint16(184),
	12: uint16('F'),
	13: uint16(184),
	14: uint16('L'),
	15: uint16(184),
	16: uint16('S'),
	17: uint16(184),
	18: uint16('d'),
	19: uint16(184),
	20: uint16('f'),
	21: uint16(184),
	22: uint16('l'),
	23: uint16(184),
	24: uint16('s'),
	25: uint16(184),
}

var map_token4 = [28]uint16_t{
	0:  uint16('.'),
	1:  uint16(139),
	2:  uint16('/'),
	3:  uint16(255),
	4:  uint16('i'),
	5:  uint16(310),
	6:  uint16('|'),
	7:  uint16(255),
	8:  uint16('E'),
	9:  uint16(184),
	10: uint16('e'),
	11: uint16(184),
	12: uint16('D'),
	13: uint16(184),
	14: uint16('F'),
	15: uint16(184),
	16: uint16('L'),
	17: uint16(184),
	18: uint16('S'),
	19: uint16(184),
	20: uint16('d'),
	21: uint16(184),
	22: uint16('f'),
	23: uint16(184),
	24: uint16('l'),
	25: uint16(184),
	26: uint16('s'),
	27: uint16(184),
}

var map_token5 = [32]uint16_t{
	0:  uint16('N'),
	1:  uint16(498),
	2:  uint16('S'),
	3:  uint16(499),
	4:  uint16('X'),
	5:  uint16(500),
	6:  uint16('a'),
	7:  uint16(495),
	8:  uint16('b'),
	9:  uint16(487),
	10: uint16('d'),
	11: uint16(492),
	12: uint16('e'),
	13: uint16(496),
	14: uint16('l'),
	15: uint16(493),
	16: uint16('n'),
	17: uint16(485),
	18: uint16('p'),
	19: uint16(488),
	20: uint16('r'),
	21: uint16(491),
	22: uint16('s'),
	23: uint16(486),
	24: uint16('t'),
	25: uint16(490),
	26: uint16('u'),
	27: uint16(500),
	28: uint16('v'),
	29: uint16(497),
	30: uint16('x'),
	31: uint16(500),
}

var map_token6 = [24]uint16_t{
	0:  uint16('i'),
	1:  uint16(310),
	2:  uint16('|'),
	3:  uint16(255),
	4:  uint16('E'),
	5:  uint16(184),
	6:  uint16('e'),
	7:  uint16(184),
	8:  uint16('D'),
	9:  uint16(184),
	10: uint16('F'),
	11: uint16(184),
	12: uint16('L'),
	13: uint16(184),
	14: uint16('S'),
	15: uint16(184),
	16: uint16('d'),
	17: uint16(184),
	18: uint16('f'),
	19: uint16(184),
	20: uint16('l'),
	21: uint16(184),
	22: uint16('s'),
	23: uint16(184),
}

var map_token7 = [16]uint16_t{
	0:  uint16('B'),
	1:  uint16(399),
	2:  uint16('b'),
	3:  uint16(399),
	4:  uint16('D'),
	5:  uint16(34),
	6:  uint16('d'),
	7:  uint16(34),
	8:  uint16('O'),
	9:  uint16(401),
	10: uint16('o'),
	11: uint16(401),
	12: uint16('X'),
	13: uint16(403),
	14: uint16('x'),
	15: uint16(403),
}

var map_token8 = [32]uint16_t{
	0:  uint16('"'),
	1:  uint16(501),
	2:  uint16('#'),
	3:  uint16(3),
	4:  uint16('\''),
	5:  uint16(610),
	6:  uint16('('),
	7:  uint16(604),
	8:  uint16(')'),
	9:  uint16(605),
	10: uint16(','),
	11: uint16(615),
	12: uint16('.'),
	13: uint16(579),
	14: uint16(';'),
	15: uint16(297),
	16: uint16('['),
	17: uint16(606),
	18: uint16(']'),
	19: uint16(607),
	20: uint16('`'),
	21: uint16(611),
	22: uint16('{'),
	23: uint16(608),
	24: uint16('|'),
	25: uint16(116),
	26: uint16('}'),
	27: uint16(609),
	28: uint16('+'),
	29: uint16(517),
	30: uint16('-'),
	31: uint16(517),
}

var map_token9 = [34]uint16_t{
	0:  uint16('#'),
	1:  uint16(312),
	2:  uint16('.'),
	3:  uint16(315),
	4:  uint16('/'),
	5:  uint16(581),
	6:  uint16('@'),
	7:  uint16(521),
	8:  uint16('|'),
	9:  uint16(245),
	10: uint16('+'),
	11: uint16(522),
	12: uint16('-'),
	13: uint16(522),
	14: uint16('E'),
	15: uint16(559),
	16: uint16('e'),
	17: uint16(559),
	18: uint16('D'),
	19: uint16(560),
	20: uint16('F'),
	21: uint16(560),
	22: uint16('L'),
	23: uint16(560),
	24: uint16('S'),
	25: uint16(560),
	26: uint16('d'),
	27: uint16(560),
	28: uint16('f'),
	29: uint16(560),
	30: uint16('l'),
	31: uint16(560),
	32: uint16('s'),
	33: uint16(560),
}

var map_token10 = [34]uint16_t{
	0:  uint16('#'),
	1:  uint16(312),
	2:  uint16('.'),
	3:  uint16(317),
	4:  uint16('/'),
	5:  uint16(273),
	6:  uint16('@'),
	7:  uint16(42),
	8:  uint16('|'),
	9:  uint16(245),
	10: uint16('+'),
	11: uint16(43),
	12: uint16('-'),
	13: uint16(43),
	14: uint16('E'),
	15: uint16(178),
	16: uint16('e'),
	17: uint16(178),
	18: uint16('D'),
	19: uint16(179),
	20: uint16('F'),
	21: uint16(179),
	22: uint16('L'),
	23: uint16(179),
	24: uint16('S'),
	25: uint16(179),
	26: uint16('d'),
	27: uint16(179),
	28: uint16('f'),
	29: uint16(179),
	30: uint16('l'),
	31: uint16(179),
	32: uint16('s'),
	33: uint16(179),
}

var map_token11 = [30]uint16_t{
	0:  uint16('#'),
	1:  uint16(316),
	2:  uint16('@'),
	3:  uint16(521),
	4:  uint16('|'),
	5:  uint16(245),
	6:  uint16('+'),
	7:  uint16(522),
	8:  uint16('-'),
	9:  uint16(522),
	10: uint16('E'),
	11: uint16(559),
	12: uint16('e'),
	13: uint16(559),
	14: uint16('D'),
	15: uint16(560),
	16: uint16('F'),
	17: uint16(560),
	18: uint16('L'),
	19: uint16(560),
	20: uint16('S'),
	21: uint16(560),
	22: uint16('d'),
	23: uint16(560),
	24: uint16('f'),
	25: uint16(560),
	26: uint16('l'),
	27: uint16(560),
	28: uint16('s'),
	29: uint16(560),
}

var map_token12 = [30]uint16_t{
	0:  uint16('#'),
	1:  uint16(316),
	2:  uint16('@'),
	3:  uint16(42),
	4:  uint16('|'),
	5:  uint16(245),
	6:  uint16('+'),
	7:  uint16(43),
	8:  uint16('-'),
	9:  uint16(43),
	10: uint16('E'),
	11: uint16(178),
	12: uint16('e'),
	13: uint16(178),
	14: uint16('D'),
	15: uint16(179),
	16: uint16('F'),
	17: uint16(179),
	18: uint16('L'),
	19: uint16(179),
	20: uint16('S'),
	21: uint16(179),
	22: uint16('d'),
	23: uint16(179),
	24: uint16('f'),
	25: uint16(179),
	26: uint16('l'),
	27: uint16(179),
	28: uint16('s'),
	29: uint16(179),
}

var map_token13 = [36]uint16_t{
	0:  uint16('#'),
	1:  uint16(319),
	2:  uint16('.'),
	3:  uint16(323),
	4:  uint16('/'),
	5:  uint16(583),
	6:  uint16('@'),
	7:  uint16(521),
	8:  uint16('i'),
	9:  uint16(483),
	10: uint16('|'),
	11: uint16(247),
	12: uint16('+'),
	13: uint16(522),
	14: uint16('-'),
	15: uint16(522),
	16: uint16('E'),
	17: uint16(561),
	18: uint16('e'),
	19: uint16(561),
	20: uint16('D'),
	21: uint16(562),
	22: uint16('F'),
	23: uint16(562),
	24: uint16('L'),
	25: uint16(562),
	26: uint16('S'),
	27: uint16(562),
	28: uint16('d'),
	29: uint16(562),
	30: uint16('f'),
	31: uint16(562),
	32: uint16('l'),
	33: uint16(562),
	34: uint16('s'),
	35: uint16(562),
}

var map_token14 = [36]uint16_t{
	0:  uint16('#'),
	1:  uint16(319),
	2:  uint16('.'),
	3:  uint16(325),
	4:  uint16('/'),
	5:  uint16(275),
	6:  uint16('@'),
	7:  uint16(42),
	8:  uint16('i'),
	9:  uint16(310),
	10: uint16('|'),
	11: uint16(247),
	12: uint16('+'),
	13: uint16(43),
	14: uint16('-'),
	15: uint16(43),
	16: uint16('E'),
	17: uint16(181),
	18: uint16('e'),
	19: uint16(181),
	20: uint16('D'),
	21: uint16(182),
	22: uint16('F'),
	23: uint16(182),
	24: uint16('L'),
	25: uint16(182),
	26: uint16('S'),
	27: uint16(182),
	28: uint16('d'),
	29: uint16(182),
	30: uint16('f'),
	31: uint16(182),
	32: uint16('l'),
	33: uint16(182),
	34: uint16('s'),
	35: uint16(182),
}

var map_token15 = [32]uint16_t{
	0:  uint16('#'),
	1:  uint16(324),
	2:  uint16('@'),
	3:  uint16(521),
	4:  uint16('i'),
	5:  uint16(483),
	6:  uint16('|'),
	7:  uint16(247),
	8:  uint16('+'),
	9:  uint16(522),
	10: uint16('-'),
	11: uint16(522),
	12: uint16('E'),
	13: uint16(561),
	14: uint16('e'),
	15: uint16(561),
	16: uint16('D'),
	17: uint16(562),
	18: uint16('F'),
	19: uint16(562),
	20: uint16('L'),
	21: uint16(562),
	22: uint16('S'),
	23: uint16(562),
	24: uint16('d'),
	25: uint16(562),
	26: uint16('f'),
	27: uint16(562),
	28: uint16('l'),
	29: uint16(562),
	30: uint16('s'),
	31: uint16(562),
}

var map_token16 = [32]uint16_t{
	0:  uint16('#'),
	1:  uint16(324),
	2:  uint16('@'),
	3:  uint16(42),
	4:  uint16('i'),
	5:  uint16(310),
	6:  uint16('|'),
	7:  uint16(247),
	8:  uint16('+'),
	9:  uint16(43),
	10: uint16('-'),
	11: uint16(43),
	12: uint16('E'),
	13: uint16(181),
	14: uint16('e'),
	15: uint16(181),
	16: uint16('D'),
	17: uint16(182),
	18: uint16('F'),
	19: uint16(182),
	20: uint16('L'),
	21: uint16(182),
	22: uint16('S'),
	23: uint16(182),
	24: uint16('d'),
	25: uint16(182),
	26: uint16('f'),
	27: uint16(182),
	28: uint16('l'),
	29: uint16(182),
	30: uint16('s'),
	31: uint16(182),
}

var map_token17 = [28]uint16_t{
	0:  uint16('#'),
	1:  uint16(332),
	2:  uint16('.'),
	3:  uint16(344),
	4:  uint16('/'),
	5:  uint16(587),
	6:  uint16('|'),
	7:  uint16(249),
	8:  uint16('E'),
	9:  uint16(563),
	10: uint16('e'),
	11: uint16(563),
	12: uint16('D'),
	13: uint16(563),
	14: uint16('F'),
	15: uint16(563),
	16: uint16('L'),
	17: uint16(563),
	18: uint16('S'),
	19: uint16(563),
	20: uint16('d'),
	21: uint16(563),
	22: uint16('f'),
	23: uint16(563),
	24: uint16('l'),
	25: uint16(563),
	26: uint16('s'),
	27: uint16(563),
}

var map_token18 = [28]uint16_t{
	0:  uint16('#'),
	1:  uint16(332),
	2:  uint16('.'),
	3:  uint16(345),
	4:  uint16('/'),
	5:  uint16(259),
	6:  uint16('|'),
	7:  uint16(249),
	8:  uint16('E'),
	9:  uint16(183),
	10: uint16('e'),
	11: uint16(183),
	12: uint16('D'),
	13: uint16(183),
	14: uint16('F'),
	15: uint16(183),
	16: uint16('L'),
	17: uint16(183),
	18: uint16('S'),
	19: uint16(183),
	20: uint16('d'),
	21: uint16(183),
	22: uint16('f'),
	23: uint16(183),
	24: uint16('l'),
	25: uint16(183),
	26: uint16('s'),
	27: uint16(183),
}

var map_token19 = [24]uint16_t{
	0:  uint16('#'),
	1:  uint16(348),
	2:  uint16('|'),
	3:  uint16(249),
	4:  uint16('E'),
	5:  uint16(563),
	6:  uint16('e'),
	7:  uint16(563),
	8:  uint16('D'),
	9:  uint16(563),
	10: uint16('F'),
	11: uint16(563),
	12: uint16('L'),
	13: uint16(563),
	14: uint16('S'),
	15: uint16(563),
	16: uint16('d'),
	17: uint16(563),
	18: uint16('f'),
	19: uint16(563),
	20: uint16('l'),
	21: uint16(563),
	22: uint16('s'),
	23: uint16(563),
}

var map_token20 = [24]uint16_t{
	0:  uint16('#'),
	1:  uint16(348),
	2:  uint16('|'),
	3:  uint16(249),
	4:  uint16('E'),
	5:  uint16(183),
	6:  uint16('e'),
	7:  uint16(183),
	8:  uint16('D'),
	9:  uint16(183),
	10: uint16('F'),
	11: uint16(183),
	12: uint16('L'),
	13: uint16(183),
	14: uint16('S'),
	15: uint16(183),
	16: uint16('d'),
	17: uint16(183),
	18: uint16('f'),
	19: uint16(183),
	20: uint16('l'),
	21: uint16(183),
	22: uint16('s'),
	23: uint16(183),
}

var map_token21 = [26]uint16_t{
	0:  uint16('.'),
	1:  uint16(457),
	2:  uint16('/'),
	3:  uint16(597),
	4:  uint16('|'),
	5:  uint16(249),
	6:  uint16('E'),
	7:  uint16(563),
	8:  uint16('e'),
	9:  uint16(563),
	10: uint16('D'),
	11: uint16(563),
	12: uint16('F'),
	13: uint16(563),
	14: uint16('L'),
	15: uint16(563),
	16: uint16('S'),
	17: uint16(563),
	18: uint16('d'),
	19: uint16(563),
	20: uint16('f'),
	21: uint16(563),
	22: uint16('l'),
	23: uint16(563),
	24: uint16('s'),
	25: uint16(563),
}

var map_token22 = [26]uint16_t{
	0:  uint16('.'),
	1:  uint16(458),
	2:  uint16('/'),
	3:  uint16(249),
	4:  uint16('|'),
	5:  uint16(249),
	6:  uint16('E'),
	7:  uint16(183),
	8:  uint16('e'),
	9:  uint16(183),
	10: uint16('D'),
	11: uint16(183),
	12: uint16('F'),
	13: uint16(183),
	14: uint16('L'),
	15: uint16(183),
	16: uint16('S'),
	17: uint16(183),
	18: uint16('d'),
	19: uint16(183),
	20: uint16('f'),
	21: uint16(183),
	22: uint16('l'),
	23: uint16(183),
	24: uint16('s'),
	25: uint16(183),
}

var map_token23 = [20]uint16_t{
	0:  uint16('@'),
	1:  uint16(448),
	2:  uint16('I'),
	3:  uint16(206),
	4:  uint16('N'),
	5:  uint16(187),
	6:  uint16('i'),
	7:  uint16(437),
	8:  uint16('n'),
	9:  uint16(82),
	10: uint16('|'),
	11: uint16(272),
	12: uint16('+'),
	13: uint16(96),
	14: uint16('-'),
	15: uint16(96),
	16: uint16('0'),
	17: uint16(338),
	18: uint16('1'),
	19: uint16(338),
}

var map_token24 = [16]uint16_t{
	0:  uint16('@'),
	1:  uint16(449),
	2:  uint16('I'),
	3:  uint16(217),
	4:  uint16('N'),
	5:  uint16(191),
	6:  uint16('i'),
	7:  uint16(441),
	8:  uint16('n'),
	9:  uint16(86),
	10: uint16('|'),
	11: uint16(278),
	12: uint16('+'),
	13: uint16(98),
	14: uint16('-'),
	15: uint16(98),
}

var map_token25 = [16]uint16_t{
	0:  uint16('@'),
	1:  uint16(450),
	2:  uint16('I'),
	3:  uint16(218),
	4:  uint16('N'),
	5:  uint16(192),
	6:  uint16('i'),
	7:  uint16(442),
	8:  uint16('n'),
	9:  uint16(87),
	10: uint16('|'),
	11: uint16(280),
	12: uint16('+'),
	13: uint16(100),
	14: uint16('-'),
	15: uint16(100),
}

var map_token26 = [22]uint16_t{
	0:  uint16('|'),
	1:  uint16(249),
	2:  uint16('E'),
	3:  uint16(563),
	4:  uint16('e'),
	5:  uint16(563),
	6:  uint16('D'),
	7:  uint16(563),
	8:  uint16('F'),
	9:  uint16(563),
	10: uint16('L'),
	11: uint16(563),
	12: uint16('S'),
	13: uint16(563),
	14: uint16('d'),
	15: uint16(563),
	16: uint16('f'),
	17: uint16(563),
	18: uint16('l'),
	19: uint16(563),
	20: uint16('s'),
	21: uint16(563),
}

var map_token27 = [22]uint16_t{
	0:  uint16('|'),
	1:  uint16(249),
	2:  uint16('E'),
	3:  uint16(183),
	4:  uint16('e'),
	5:  uint16(183),
	6:  uint16('D'),
	7:  uint16(183),
	8:  uint16('F'),
	9:  uint16(183),
	10: uint16('L'),
	11: uint16(183),
	12: uint16('S'),
	13: uint16(183),
	14: uint16('d'),
	15: uint16(183),
	16: uint16('f'),
	17: uint16(183),
	18: uint16('l'),
	19: uint16(183),
	20: uint16('s'),
	21: uint16(183),
}

var map_token28 = [30]uint16_t{
	0:  uint16('#'),
	1:  uint16(9),
	2:  uint16('.'),
	3:  uint16(511),
	4:  uint16('/'),
	5:  uint16(589),
	6:  uint16('i'),
	7:  uint16(483),
	8:  uint16('|'),
	9:  uint16(255),
	10: uint16('E'),
	11: uint16(564),
	12: uint16('e'),
	13: uint16(564),
	14: uint16('D'),
	15: uint16(564),
	16: uint16('F'),
	17: uint16(564),
	18: uint16('L'),
	19: uint16(564),
	20: uint16('S'),
	21: uint16(564),
	22: uint16('d'),
	23: uint16(564),
	24: uint16('f'),
	25: uint16(564),
	26: uint16('l'),
	27: uint16(564),
	28: uint16('s'),
	29: uint16(564),
}

var map_token29 = [26]uint16_t{
	0:  uint16('#'),
	1:  uint16(14),
	2:  uint16('i'),
	3:  uint16(483),
	4:  uint16('|'),
	5:  uint16(255),
	6:  uint16('E'),
	7:  uint16(564),
	8:  uint16('e'),
	9:  uint16(564),
	10: uint16('D'),
	11: uint16(564),
	12: uint16('F'),
	13: uint16(564),
	14: uint16('L'),
	15: uint16(564),
	16: uint16('S'),
	17: uint16(564),
	18: uint16('d'),
	19: uint16(564),
	20: uint16('f'),
	21: uint16(564),
	22: uint16('l'),
	23: uint16(564),
	24: uint16('s'),
	25: uint16(564),
}

var map_token30 = [28]uint16_t{
	0:  uint16('.'),
	1:  uint16(551),
	2:  uint16('/'),
	3:  uint16(598),
	4:  uint16('i'),
	5:  uint16(483),
	6:  uint16('|'),
	7:  uint16(255),
	8:  uint16('E'),
	9:  uint16(564),
	10: uint16('e'),
	11: uint16(564),
	12: uint16('D'),
	13: uint16(564),
	14: uint16('F'),
	15: uint16(564),
	16: uint16('L'),
	17: uint16(564),
	18: uint16('S'),
	19: uint16(564),
	20: uint16('d'),
	21: uint16(564),
	22: uint16('f'),
	23: uint16(564),
	24: uint16('l'),
	25: uint16(564),
	26: uint16('s'),
	27: uint16(564),
}

var map_token31 = [24]uint16_t{
	0:  uint16('i'),
	1:  uint16(483),
	2:  uint16('|'),
	3:  uint16(255),
	4:  uint16('E'),
	5:  uint16(564),
	6:  uint16('e'),
	7:  uint16(564),
	8:  uint16('D'),
	9:  uint16(564),
	10: uint16('F'),
	11: uint16(564),
	12: uint16('L'),
	13: uint16(564),
	14: uint16('S'),
	15: uint16(564),
	16: uint16('d'),
	17: uint16(564),
	18: uint16('f'),
	19: uint16(564),
	20: uint16('l'),
	21: uint16(564),
	22: uint16('s'),
	23: uint16(564),
}

var ts_lex_modes = [143]TSLexMode{
	0: {},
	1: {
		Flex_state: uint16(294),
	},
	2: {
		Flex_state: uint16(294),
	},
	3: {
		Flex_state: uint16(294),
	},
	4: {
		Flex_state: uint16(294),
	},
	5: {
		Flex_state: uint16(294),
	},
	6: {
		Flex_state: uint16(294),
	},
	7: {
		Flex_state: uint16(294),
	},
	8: {
		Flex_state: uint16(294),
	},
	9: {
		Flex_state: uint16(294),
	},
	10: {
		Flex_state: uint16(294),
	},
	11: {
		Flex_state: uint16(294),
	},
	12: {
		Flex_state: uint16(294),
	},
	13: {
		Flex_state: uint16(294),
	},
	14: {
		Flex_state: uint16(294),
	},
	15: {
		Flex_state: uint16(294),
	},
	16: {
		Flex_state: uint16(294),
	},
	17: {
		Flex_state: uint16(294),
	},
	18: {
		Flex_state: uint16(294),
	},
	19: {
		Flex_state: uint16(294),
	},
	20: {
		Flex_state: uint16(294),
	},
	21: {
		Flex_state: uint16(294),
	},
	22: {
		Flex_state: uint16(294),
	},
	23: {
		Flex_state: uint16(294),
	},
	24: {
		Flex_state: uint16(294),
	},
	25: {
		Flex_state: uint16(294),
	},
	26: {
		Flex_state: uint16(294),
	},
	27: {
		Flex_state: uint16(294),
	},
	28: {
		Flex_state: uint16(294),
	},
	29: {
		Flex_state: uint16(294),
	},
	30: {
		Flex_state: uint16(294),
	},
	31: {
		Flex_state: uint16(294),
	},
	32: {
		Flex_state: uint16(294),
	},
	33: {
		Flex_state: uint16(294),
	},
	34: {
		Flex_state: uint16(294),
	},
	35: {
		Flex_state: uint16(294),
	},
	36: {
		Flex_state: uint16(294),
	},
	37: {
		Flex_state: uint16(294),
	},
	38: {
		Flex_state: uint16(294),
	},
	39: {
		Flex_state: uint16(294),
	},
	40: {
		Flex_state: uint16(294),
	},
	41: {
		Flex_state: uint16(294),
	},
	42: {
		Flex_state: uint16(294),
	},
	43: {
		Flex_state: uint16(294),
	},
	44: {
		Flex_state: uint16(294),
	},
	45: {
		Flex_state: uint16(294),
	},
	46: {
		Flex_state: uint16(294),
	},
	47: {
		Flex_state: uint16(294),
	},
	48: {
		Flex_state: uint16(294),
	},
	49: {
		Flex_state: uint16(294),
	},
	50: {
		Flex_state: uint16(294),
	},
	51: {
		Flex_state: uint16(294),
	},
	52: {
		Flex_state: uint16(294),
	},
	53: {
		Flex_state: uint16(294),
	},
	54: {
		Flex_state: uint16(294),
	},
	55: {
		Flex_state: uint16(294),
	},
	56: {
		Flex_state: uint16(294),
	},
	57: {
		Flex_state: uint16(294),
	},
	58: {
		Flex_state: uint16(294),
	},
	59: {
		Flex_state: uint16(294),
	},
	60: {
		Flex_state: uint16(294),
	},
	61: {
		Flex_state: uint16(294),
	},
	62: {
		Flex_state: uint16(294),
	},
	63: {
		Flex_state: uint16(294),
	},
	64: {
		Flex_state: uint16(294),
	},
	65: {
		Flex_state: uint16(294),
	},
	66: {
		Flex_state: uint16(294),
	},
	67: {
		Flex_state: uint16(294),
	},
	68: {
		Flex_state: uint16(294),
	},
	69: {
		Flex_state: uint16(294),
	},
	70: {
		Flex_state: uint16(294),
	},
	71: {
		Flex_state: uint16(294),
	},
	72: {
		Flex_state: uint16(294),
	},
	73: {
		Flex_state: uint16(294),
	},
	74: {
		Flex_state: uint16(294),
	},
	75: {
		Flex_state: uint16(294),
	},
	76: {
		Flex_state: uint16(294),
	},
	77: {
		Flex_state: uint16(294),
	},
	78: {
		Flex_state: uint16(294),
	},
	79: {
		Flex_state: uint16(294),
	},
	80: {
		Flex_state: uint16(294),
	},
	81: {
		Flex_state: uint16(294),
	},
	82: {
		Flex_state: uint16(294),
	},
	83: {
		Flex_state: uint16(294),
	},
	84: {
		Flex_state: uint16(294),
	},
	85: {
		Flex_state: uint16(294),
	},
	86: {
		Flex_state: uint16(294),
	},
	87: {
		Flex_state: uint16(294),
	},
	88: {
		Flex_state: uint16(294),
	},
	89: {
		Flex_state: uint16(294),
	},
	90: {
		Flex_state: uint16(294),
	},
	91: {
		Flex_state: uint16(294),
	},
	92: {
		Flex_state: uint16(23),
	},
	93: {
		Flex_state: uint16(23),
	},
	94: {
		Flex_state: uint16(23),
	},
	95: {
		Flex_state: uint16(23),
	},
	96: {
		Flex_state: uint16(23),
	},
	97: {
		Flex_state: uint16(23),
	},
	98: {
		Flex_state: uint16(23),
	},
	99: {
		Flex_state: uint16(23),
	},
	100: {
		Flex_state: uint16(23),
	},
	101: {
		Flex_state: uint16(23),
	},
	102: {
		Flex_state: uint16(23),
	},
	103: {
		Flex_state: uint16(23),
	},
	104: {
		Flex_state: uint16(23),
	},
	105: {
		Flex_state: uint16(23),
	},
	106: {
		Flex_state: uint16(23),
	},
	107: {
		Flex_state: uint16(23),
	},
	108: {
		Flex_state: uint16(23),
	},
	109: {
		Flex_state: uint16(23),
	},
	110: {
		Flex_state: uint16(23),
	},
	111: {
		Flex_state: uint16(23),
	},
	112: {
		Flex_state: uint16(23),
	},
	113: {
		Flex_state: uint16(23),
	},
	114: {
		Flex_state: uint16(23),
	},
	115: {
		Flex_state: uint16(23),
	},
	116: {
		Flex_state: uint16(23),
	},
	117: {
		Flex_state: uint16(23),
	},
	118: {
		Flex_state: uint16(23),
	},
	119: {
		Flex_state: uint16(23),
	},
	120: {
		Flex_state: uint16(23),
	},
	121: {
		Flex_state: uint16(23),
	},
	122: {
		Flex_state: uint16(23),
	},
	123: {
		Flex_state: uint16(23),
	},
	124: {
		Flex_state: uint16(23),
	},
	125: {
		Flex_state: uint16(23),
	},
	126: {
		Flex_state: uint16(23),
	},
	127: {
		Flex_state: uint16(23),
	},
	128: {
		Flex_state: uint16(6),
	},
	129: {
		Flex_state: uint16(6),
	},
	130: {
		Flex_state: uint16(6),
	},
	131: {
		Flex_state: uint16(6),
	},
	132: {
		Flex_state: uint16(6),
	},
	133: {
		Flex_state: uint16(6),
	},
	134: {
		Flex_state: uint16(6),
	},
	135: {
		Flex_state: uint16(5),
	},
	136: {
		Flex_state: uint16(5),
	},
	137: {
		Flex_state: uint16(5),
	},
	138: {
		Flex_state: uint16(5),
	},
	139: {
		Flex_state: uint16(5),
	},
	140: {
		Flex_state: uint16(6),
	},
	141: {
		Flex_state: uint16(6),
	},
	142: {},
}

var ts_parse_table = [61][56]uint16_t{
	0: {
		0:  uint16(1),
		1:  uint16(1),
		2:  uint16(1),
		6:  uint16(1),
		7:  uint16(1),
		8:  uint16(1),
		12: uint16(1),
		17: uint16(1),
		18: uint16(1),
		19: uint16(1),
		20: uint16(1),
		21: uint16(1),
		22: uint16(1),
		23: uint16(1),
		24: uint16(1),
		27: uint16(1),
	},
	1: {
		0:  uint16(3),
		1:  uint16(5),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		6:  uint16(13),
		9:  uint16(5),
		10: uint16(15),
		11: uint16(5),
		12: uint16(17),
		15: uint16(15),
		16: uint16(5),
		17: uint16(19),
		19: uint16(21),
		21: uint16(23),
		23: uint16(25),
		24: uint16(27),
		25: uint16(29),
		26: uint16(31),
		27: uint16(33),
		28: uint16(35),
		29: uint16(37),
		30: uint16(39),
		31: uint16(41),
		32: uint16(43),
		33: uint16(142),
		34: uint16(7),
		35: uint16(7),
		36: uint16(7),
		37: uint16(7),
		38: uint16(7),
		39: uint16(7),
		40: uint16(7),
		41: uint16(7),
		42: uint16(7),
		43: uint16(7),
		44: uint16(7),
		45: uint16(7),
		46: uint16(7),
		47: uint16(7),
		48: uint16(7),
		49: uint16(7),
		50: uint16(7),
		51: uint16(7),
		52: uint16(7),
	},
	2: {
		0:  uint16(45),
		1:  uint16(47),
		2:  uint16(50),
		3:  uint16(53),
		4:  uint16(56),
		6:  uint16(59),
		9:  uint16(47),
		10: uint16(62),
		11: uint16(47),
		12: uint16(65),
		15: uint16(62),
		16: uint16(47),
		17: uint16(68),
		18: uint16(45),
		19: uint16(71),
		20: uint16(45),
		21: uint16(74),
		22: uint16(45),
		23: uint16(77),
		24: uint16(80),
		25: uint16(83),
		26: uint16(86),
		27: uint16(89),
		28: uint16(92),
		29: uint16(95),
		30: uint16(98),
		31: uint16(101),
		32: uint16(104),
		34: uint16(2),
		35: uint16(2),
		36: uint16(2),
		37: uint16(2),
		38: uint16(2),
		39: uint16(2),
		40: uint16(2),
		41: uint16(2),
		42: uint16(2),
		43: uint16(2),
		44: uint16(2),
		45: uint16(2),
		46: uint16(2),
		47: uint16(2),
		48: uint16(2),
		49: uint16(2),
		50: uint16(2),
		51: uint16(2),
		52: uint16(2),
	},
	3: {
		1:  uint16(107),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		6:  uint16(13),
		9:  uint16(107),
		10: uint16(109),
		11: uint16(107),
		12: uint16(17),
		15: uint16(109),
		16: uint16(107),
		17: uint16(19),
		18: uint16(111),
		19: uint16(21),
		21: uint16(23),
		23: uint16(25),
		24: uint16(27),
		25: uint16(29),
		26: uint16(31),
		27: uint16(33),
		28: uint16(35),
		29: uint16(37),
		30: uint16(39),
		31: uint16(41),
		32: uint16(43),
		34: uint16(2),
		35: uint16(2),
		36: uint16(2),
		37: uint16(2),
		38: uint16(2),
		39: uint16(2),
		40: uint16(2),
		41: uint16(2),
		42: uint16(2),
		43: uint16(2),
		44: uint16(2),
		45: uint16(2),
		46: uint16(2),
		47: uint16(2),
		48: uint16(2),
		49: uint16(2),
		50: uint16(2),
		51: uint16(2),
		52: uint16(2),
	},
	4: {
		1:  uint16(107),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		6:  uint16(13),
		9:  uint16(107),
		10: uint16(109),
		11: uint16(107),
		12: uint16(17),
		15: uint16(109),
		16: uint16(107),
		17: uint16(19),
		18: uint16(113),
		19: uint16(21),
		21: uint16(23),
		23: uint16(25),
		24: uint16(27),
		25: uint16(29),
		26: uint16(31),
		27: uint16(33),
		28: uint16(35),
		29: uint16(37),
		30: uint16(39),
		31: uint16(41),
		32: uint16(43),
		34: uint16(2),
		35: uint16(2),
		36: uint16(2),
		37: uint16(2),
		38: uint16(2),
		39: uint16(2),
		40: uint16(2),
		41: uint16(2),
		42: uint16(2),
		43: uint16(2),
		44: uint16(2),
		45: uint16(2),
		46: uint16(2),
		47: uint16(2),
		48: uint16(2),
		49: uint16(2),
		50: uint16(2),
		51: uint16(2),
		52: uint16(2),
	},
	5: {
		1:  uint16(115),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		6:  uint16(13),
		9:  uint16(115),
		10: uint16(117),
		11: uint16(115),
		12: uint16(17),
		15: uint16(117),
		16: uint16(115),
		17: uint16(19),
		18: uint16(119),
		19: uint16(21),
		21: uint16(23),
		23: uint16(25),
		24: uint16(27),
		25: uint16(29),
		26: uint16(31),
		27: uint16(33),
		28: uint16(35),
		29: uint16(37),
		30: uint16(39),
		31: uint16(41),
		32: uint16(43),
		34: uint16(3),
		35: uint16(3),
		36: uint16(3),
		37: uint16(3),
		38: uint16(3),
		39: uint16(3),
		40: uint16(3),
		41: uint16(3),
		42: uint16(3),
		43: uint16(3),
		44: uint16(3),
		45: uint16(3),
		46: uint16(3),
		47: uint16(3),
		48: uint16(3),
		49: uint16(3),
		50: uint16(3),
		51: uint16(3),
		52: uint16(3),
	},
	6: {
		1:  uint16(121),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		6:  uint16(13),
		9:  uint16(121),
		10: uint16(123),
		11: uint16(121),
		12: uint16(17),
		15: uint16(123),
		16: uint16(121),
		17: uint16(19),
		18: uint16(125),
		19: uint16(21),
		21: uint16(23),
		23: uint16(25),
		24: uint16(27),
		25: uint16(29),
		26: uint16(31),
		27: uint16(33),
		28: uint16(35),
		29: uint16(37),
		30: uint16(39),
		31: uint16(41),
		32: uint16(43),
		34: uint16(12),
		35: uint16(12),
		36: uint16(12),
		37: uint16(12),
		38: uint16(12),
		39: uint16(12),
		40: uint16(12),
		41: uint16(12),
		42: uint16(12),
		43: uint16(12),
		44: uint16(12),
		45: uint16(12),
		46: uint16(12),
		47: uint16(12),
		48: uint16(12),
		49: uint16(12),
		50: uint16(12),
		51: uint16(12),
		52: uint16(12),
	},
	7: {
		0:  uint16(127),
		1:  uint16(107),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		6:  uint16(13),
		9:  uint16(107),
		10: uint16(109),
		11: uint16(107),
		12: uint16(17),
		15: uint16(109),
		16: uint16(107),
		17: uint16(19),
		19: uint16(21),
		21: uint16(23),
		23: uint16(25),
		24: uint16(27),
		25: uint16(29),
		26: uint16(31),
		27: uint16(33),
		28: uint16(35),
		29: uint16(37),
		30: uint16(39),
		31: uint16(41),
		32: uint16(43),
		34: uint16(2),
		35: uint16(2),
		36: uint16(2),
		37: uint16(2),
		38: uint16(2),
		39: uint16(2),
		40: uint16(2),
		41: uint16(2),
		42: uint16(2),
		43: uint16(2),
		44: uint16(2),
		45: uint16(2),
		46: uint16(2),
		47: uint16(2),
		48: uint16(2),
		49: uint16(2),
		50: uint16(2),
		51: uint16(2),
		52: uint16(2),
	},
	8: {
		1:  uint16(107),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		6:  uint16(13),
		9:  uint16(107),
		10: uint16(109),
		11: uint16(107),
		12: uint16(17),
		15: uint16(109),
		16: uint16(107),
		17: uint16(19),
		18: uint16(129),
		19: uint16(21),
		21: uint16(23),
		23: uint16(25),
		24: uint16(27),
		25: uint16(29),
		26: uint16(31),
		27: uint16(33),
		28: uint16(35),
		29: uint16(37),
		30: uint16(39),
		31: uint16(41),
		32: uint16(43),
		34: uint16(2),
		35: uint16(2),
		36: uint16(2),
		37: uint16(2),
		38: uint16(2),
		39: uint16(2),
		40: uint16(2),
		41: uint16(2),
		42: uint16(2),
		43: uint16(2),
		44: uint16(2),
		45: uint16(2),
		46: uint16(2),
		47: uint16(2),
		48: uint16(2),
		49: uint16(2),
		50: uint16(2),
		51: uint16(2),
		52: uint16(2),
	},
	9: {
		1:  uint16(107),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		6:  uint16(13),
		9:  uint16(107),
		10: uint16(109),
		11: uint16(107),
		12: uint16(17),
		15: uint16(109),
		16: uint16(107),
		17: uint16(19),
		19: uint16(21),
		20: uint16(129),
		21: uint16(23),
		23: uint16(25),
		24: uint16(27),
		25: uint16(29),
		26: uint16(31),
		27: uint16(33),
		28: uint16(35),
		29: uint16(37),
		30: uint16(39),
		31: uint16(41),
		32: uint16(43),
		34: uint16(2),
		35: uint16(2),
		36: uint16(2),
		37: uint16(2),
		38: uint16(2),
		39: uint16(2),
		40: uint16(2),
		41: uint16(2),
		42: uint16(2),
		43: uint16(2),
		44: uint16(2),
		45: uint16(2),
		46: uint16(2),
		47: uint16(2),
		48: uint16(2),
		49: uint16(2),
		50: uint16(2),
		51: uint16(2),
		52: uint16(2),
	},
	10: {
		1:  uint16(107),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		6:  uint16(13),
		9:  uint16(107),
		10: uint16(109),
		11: uint16(107),
		12: uint16(17),
		15: uint16(109),
		16: uint16(107),
		17: uint16(19),
		19: uint16(21),
		21: uint16(23),
		22: uint16(129),
		23: uint16(25),
		24: uint16(27),
		25: uint16(29),
		26: uint16(31),
		27: uint16(33),
		28: uint16(35),
		29: uint16(37),
		30: uint16(39),
		31: uint16(41),
		32: uint16(43),
		34: uint16(2),
		35: uint16(2),
		36: uint16(2),
		37: uint16(2),
		38: uint16(2),
		39: uint16(2),
		40: uint16(2),
		41: uint16(2),
		42: uint16(2),
		43: uint16(2),
		44: uint16(2),
		45: uint16(2),
		46: uint16(2),
		47: uint16(2),
		48: uint16(2),
		49: uint16(2),
		50: uint16(2),
		51: uint16(2),
		52: uint16(2),
	},
	11: {
		1:  uint16(131),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		6:  uint16(13),
		9:  uint16(131),
		10: uint16(133),
		11: uint16(131),
		12: uint16(17),
		15: uint16(133),
		16: uint16(131),
		17: uint16(19),
		18: uint16(135),
		19: uint16(21),
		21: uint16(23),
		23: uint16(25),
		24: uint16(27),
		25: uint16(29),
		26: uint16(31),
		27: uint16(33),
		28: uint16(35),
		29: uint16(37),
		30: uint16(39),
		31: uint16(41),
		32: uint16(43),
		34: uint16(8),
		35: uint16(8),
		36: uint16(8),
		37: uint16(8),
		38: uint16(8),
		39: uint16(8),
		40: uint16(8),
		41: uint16(8),
		42: uint16(8),
		43: uint16(8),
		44: uint16(8),
		45: uint16(8),
		46: uint16(8),
		47: uint16(8),
		48: uint16(8),
		49: uint16(8),
		50: uint16(8),
		51: uint16(8),
		52: uint16(8),
	},
	12: {
		1:  uint16(107),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		6:  uint16(13),
		9:  uint16(107),
		10: uint16(109),
		11: uint16(107),
		12: uint16(17),
		15: uint16(109),
		16: uint16(107),
		17: uint16(19),
		18: uint16(137),
		19: uint16(21),
		21: uint16(23),
		23: uint16(25),
		24: uint16(27),
		25: uint16(29),
		26: uint16(31),
		27: uint16(33),
		28: uint16(35),
		29: uint16(37),
		30: uint16(39),
		31: uint16(41),
		32: uint16(43),
		34: uint16(2),
		35: uint16(2),
		36: uint16(2),
		37: uint16(2),
		38: uint16(2),
		39: uint16(2),
		40: uint16(2),
		41: uint16(2),
		42: uint16(2),
		43: uint16(2),
		44: uint16(2),
		45: uint16(2),
		46: uint16(2),
		47: uint16(2),
		48: uint16(2),
		49: uint16(2),
		50: uint16(2),
		51: uint16(2),
		52: uint16(2),
	},
	13: {
		1:  uint16(139),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		6:  uint16(13),
		9:  uint16(139),
		10: uint16(141),
		11: uint16(139),
		12: uint16(17),
		15: uint16(141),
		16: uint16(139),
		17: uint16(19),
		19: uint16(21),
		20: uint16(135),
		21: uint16(23),
		23: uint16(25),
		24: uint16(27),
		25: uint16(29),
		26: uint16(31),
		27: uint16(33),
		28: uint16(35),
		29: uint16(37),
		30: uint16(39),
		31: uint16(41),
		32: uint16(43),
		34: uint16(9),
		35: uint16(9),
		36: uint16(9),
		37: uint16(9),
		38: uint16(9),
		39: uint16(9),
		40: uint16(9),
		41: uint16(9),
		42: uint16(9),
		43: uint16(9),
		44: uint16(9),
		45: uint16(9),
		46: uint16(9),
		47: uint16(9),
		48: uint16(9),
		49: uint16(9),
		50: uint16(9),
		51: uint16(9),
		52: uint16(9),
	},
	14: {
		1:  uint16(143),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		6:  uint16(13),
		9:  uint16(143),
		10: uint16(145),
		11: uint16(143),
		12: uint16(17),
		15: uint16(145),
		16: uint16(143),
		17: uint16(19),
		18: uint16(147),
		19: uint16(21),
		21: uint16(23),
		23: uint16(25),
		24: uint16(27),
		25: uint16(29),
		26: uint16(31),
		27: uint16(33),
		28: uint16(35),
		29: uint16(37),
		30: uint16(39),
		31: uint16(41),
		32: uint16(43),
		34: uint16(19),
		35: uint16(19),
		36: uint16(19),
		37: uint16(19),
		38: uint16(19),
		39: uint16(19),
		40: uint16(19),
		41: uint16(19),
		42: uint16(19),
		43: uint16(19),
		44: uint16(19),
		45: uint16(19),
		46: uint16(19),
		47: uint16(19),
		48: uint16(19),
		49: uint16(19),
		50: uint16(19),
		51: uint16(19),
		52: uint16(19),
	},
	15: {
		1:  uint16(149),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		6:  uint16(13),
		9:  uint16(149),
		10: uint16(151),
		11: uint16(149),
		12: uint16(17),
		15: uint16(151),
		16: uint16(149),
		17: uint16(19),
		19: uint16(21),
		20: uint16(147),
		21: uint16(23),
		23: uint16(25),
		24: uint16(27),
		25: uint16(29),
		26: uint16(31),
		27: uint16(33),
		28: uint16(35),
		29: uint16(37),
		30: uint16(39),
		31: uint16(41),
		32: uint16(43),
		34: uint16(20),
		35: uint16(20),
		36: uint16(20),
		37: uint16(20),
		38: uint16(20),
		39: uint16(20),
		40: uint16(20),
		41: uint16(20),
		42: uint16(20),
		43: uint16(20),
		44: uint16(20),
		45: uint16(20),
		46: uint16(20),
		47: uint16(20),
		48: uint16(20),
		49: uint16(20),
		50: uint16(20),
		51: uint16(20),
		52: uint16(20),
	},
	16: {
		1:  uint16(153),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		6:  uint16(13),
		9:  uint16(153),
		10: uint16(155),
		11: uint16(153),
		12: uint16(17),
		15: uint16(155),
		16: uint16(153),
		17: uint16(19),
		19: uint16(21),
		21: uint16(23),
		22: uint16(147),
		23: uint16(25),
		24: uint16(27),
		25: uint16(29),
		26: uint16(31),
		27: uint16(33),
		28: uint16(35),
		29: uint16(37),
		30: uint16(39),
		31: uint16(41),
		32: uint16(43),
		34: uint16(21),
		35: uint16(21),
		36: uint16(21),
		37: uint16(21),
		38: uint16(21),
		39: uint16(21),
		40: uint16(21),
		41: uint16(21),
		42: uint16(21),
		43: uint16(21),
		44: uint16(21),
		45: uint16(21),
		46: uint16(21),
		47: uint16(21),
		48: uint16(21),
		49: uint16(21),
		50: uint16(21),
		51: uint16(21),
		52: uint16(21),
	},
	17: {
		1:  uint16(157),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		6:  uint16(13),
		9:  uint16(157),
		10: uint16(159),
		11: uint16(157),
		12: uint16(17),
		15: uint16(159),
		16: uint16(157),
		17: uint16(19),
		18: uint16(161),
		19: uint16(21),
		21: uint16(23),
		23: uint16(25),
		24: uint16(27),
		25: uint16(29),
		26: uint16(31),
		27: uint16(33),
		28: uint16(35),
		29: uint16(37),
		30: uint16(39),
		31: uint16(41),
		32: uint16(43),
		34: uint16(22),
		35: uint16(22),
		36: uint16(22),
		37: uint16(22),
		38: uint16(22),
		39: uint16(22),
		40: uint16(22),
		41: uint16(22),
		42: uint16(22),
		43: uint16(22),
		44: uint16(22),
		45: uint16(22),
		46: uint16(22),
		47: uint16(22),
		48: uint16(22),
		49: uint16(22),
		50: uint16(22),
		51: uint16(22),
		52: uint16(22),
	},
	18: {
		1:  uint16(163),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		6:  uint16(13),
		9:  uint16(163),
		10: uint16(165),
		11: uint16(163),
		12: uint16(17),
		15: uint16(165),
		16: uint16(163),
		17: uint16(19),
		18: uint16(167),
		19: uint16(21),
		21: uint16(23),
		23: uint16(25),
		24: uint16(27),
		25: uint16(29),
		26: uint16(31),
		27: uint16(33),
		28: uint16(35),
		29: uint16(37),
		30: uint16(39),
		31: uint16(41),
		32: uint16(43),
		34: uint16(4),
		35: uint16(4),
		36: uint16(4),
		37: uint16(4),
		38: uint16(4),
		39: uint16(4),
		40: uint16(4),
		41: uint16(4),
		42: uint16(4),
		43: uint16(4),
		44: uint16(4),
		45: uint16(4),
		46: uint16(4),
		47: uint16(4),
		48: uint16(4),
		49: uint16(4),
		50: uint16(4),
		51: uint16(4),
		52: uint16(4),
	},
	19: {
		1:  uint16(107),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		6:  uint16(13),
		9:  uint16(107),
		10: uint16(109),
		11: uint16(107),
		12: uint16(17),
		15: uint16(109),
		16: uint16(107),
		17: uint16(19),
		18: uint16(169),
		19: uint16(21),
		21: uint16(23),
		23: uint16(25),
		24: uint16(27),
		25: uint16(29),
		26: uint16(31),
		27: uint16(33),
		28: uint16(35),
		29: uint16(37),
		30: uint16(39),
		31: uint16(41),
		32: uint16(43),
		34: uint16(2),
		35: uint16(2),
		36: uint16(2),
		37: uint16(2),
		38: uint16(2),
		39: uint16(2),
		40: uint16(2),
		41: uint16(2),
		42: uint16(2),
		43: uint16(2),
		44: uint16(2),
		45: uint16(2),
		46: uint16(2),
		47: uint16(2),
		48: uint16(2),
		49: uint16(2),
		50: uint16(2),
		51: uint16(2),
		52: uint16(2),
	},
	20: {
		1:  uint16(107),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		6:  uint16(13),
		9:  uint16(107),
		10: uint16(109),
		11: uint16(107),
		12: uint16(17),
		15: uint16(109),
		16: uint16(107),
		17: uint16(19),
		19: uint16(21),
		20: uint16(169),
		21: uint16(23),
		23: uint16(25),
		24: uint16(27),
		25: uint16(29),
		26: uint16(31),
		27: uint16(33),
		28: uint16(35),
		29: uint16(37),
		30: uint16(39),
		31: uint16(41),
		32: uint16(43),
		34: uint16(2),
		35: uint16(2),
		36: uint16(2),
		37: uint16(2),
		38: uint16(2),
		39: uint16(2),
		40: uint16(2),
		41: uint16(2),
		42: uint16(2),
		43: uint16(2),
		44: uint16(2),
		45: uint16(2),
		46: uint16(2),
		47: uint16(2),
		48: uint16(2),
		49: uint16(2),
		50: uint16(2),
		51: uint16(2),
		52: uint16(2),
	},
	21: {
		1:  uint16(107),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		6:  uint16(13),
		9:  uint16(107),
		10: uint16(109),
		11: uint16(107),
		12: uint16(17),
		15: uint16(109),
		16: uint16(107),
		17: uint16(19),
		19: uint16(21),
		21: uint16(23),
		22: uint16(169),
		23: uint16(25),
		24: uint16(27),
		25: uint16(29),
		26: uint16(31),
		27: uint16(33),
		28: uint16(35),
		29: uint16(37),
		30: uint16(39),
		31: uint16(41),
		32: uint16(43),
		34: uint16(2),
		35: uint16(2),
		36: uint16(2),
		37: uint16(2),
		38: uint16(2),
		39: uint16(2),
		40: uint16(2),
		41: uint16(2),
		42: uint16(2),
		43: uint16(2),
		44: uint16(2),
		45: uint16(2),
		46: uint16(2),
		47: uint16(2),
		48: uint16(2),
		49: uint16(2),
		50: uint16(2),
		51: uint16(2),
		52: uint16(2),
	},
	22: {
		1:  uint16(107),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		6:  uint16(13),
		9:  uint16(107),
		10: uint16(109),
		11: uint16(107),
		12: uint16(17),
		15: uint16(109),
		16: uint16(107),
		17: uint16(19),
		18: uint16(171),
		19: uint16(21),
		21: uint16(23),
		23: uint16(25),
		24: uint16(27),
		25: uint16(29),
		26: uint16(31),
		27: uint16(33),
		28: uint16(35),
		29: uint16(37),
		30: uint16(39),
		31: uint16(41),
		32: uint16(43),
		34: uint16(2),
		35: uint16(2),
		36: uint16(2),
		37: uint16(2),
		38: uint16(2),
		39: uint16(2),
		40: uint16(2),
		41: uint16(2),
		42: uint16(2),
		43: uint16(2),
		44: uint16(2),
		45: uint16(2),
		46: uint16(2),
		47: uint16(2),
		48: uint16(2),
		49: uint16(2),
		50: uint16(2),
		51: uint16(2),
		52: uint16(2),
	},
	23: {
		1:  uint16(173),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		6:  uint16(13),
		9:  uint16(173),
		10: uint16(175),
		11: uint16(173),
		12: uint16(17),
		15: uint16(175),
		16: uint16(173),
		17: uint16(19),
		19: uint16(21),
		21: uint16(23),
		22: uint16(135),
		23: uint16(25),
		24: uint16(27),
		25: uint16(29),
		26: uint16(31),
		27: uint16(33),
		28: uint16(35),
		29: uint16(37),
		30: uint16(39),
		31: uint16(41),
		32: uint16(43),
		34: uint16(10),
		35: uint16(10),
		36: uint16(10),
		37: uint16(10),
		38: uint16(10),
		39: uint16(10),
		40: uint16(10),
		41: uint16(10),
		42: uint16(10),
		43: uint16(10),
		44: uint16(10),
		45: uint16(10),
		46: uint16(10),
		47: uint16(10),
		48: uint16(10),
		49: uint16(10),
		50: uint16(10),
		51: uint16(10),
		52: uint16(10),
	},
	24: {
		1:  uint16(177),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		6:  uint16(13),
		9:  uint16(179),
		10: uint16(181),
		11: uint16(179),
		12: uint16(183),
		15: uint16(181),
		16: uint16(179),
		17: uint16(185),
		19: uint16(187),
		21: uint16(189),
		23: uint16(191),
		24: uint16(193),
		25: uint16(195),
		26: uint16(197),
		27: uint16(199),
		28: uint16(201),
		29: uint16(203),
		30: uint16(205),
		31: uint16(207),
		32: uint16(209),
		35: uint16(47),
		36: uint16(47),
		37: uint16(47),
		38: uint16(47),
		39: uint16(101),
		40: uint16(101),
		41: uint16(101),
		42: uint16(101),
		43: uint16(101),
		44: uint16(101),
		45: uint16(101),
		46: uint16(101),
		47: uint16(101),
		48: uint16(101),
		49: uint16(101),
		50: uint16(101),
		51: uint16(101),
		53: uint16(47),
	},
	25: {
		1:  uint16(211),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		6:  uint16(13),
		9:  uint16(213),
		10: uint16(215),
		11: uint16(213),
		12: uint16(17),
		15: uint16(215),
		16: uint16(213),
		17: uint16(19),
		19: uint16(21),
		21: uint16(23),
		23: uint16(25),
		24: uint16(27),
		25: uint16(29),
		26: uint16(31),
		27: uint16(33),
		28: uint16(35),
		29: uint16(37),
		30: uint16(39),
		31: uint16(41),
		32: uint16(43),
		35: uint16(60),
		36: uint16(60),
		37: uint16(60),
		38: uint16(60),
		39: uint16(67),
		40: uint16(67),
		41: uint16(67),
		42: uint16(67),
		43: uint16(67),
		44: uint16(67),
		45: uint16(67),
		46: uint16(67),
		47: uint16(67),
		48: uint16(67),
		49: uint16(67),
		50: uint16(67),
		51: uint16(67),
		53: uint16(60),
	},
	26: {
		1:  uint16(211),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		6:  uint16(13),
		9:  uint16(217),
		10: uint16(219),
		11: uint16(217),
		12: uint16(17),
		15: uint16(219),
		16: uint16(217),
		17: uint16(19),
		19: uint16(21),
		21: uint16(23),
		23: uint16(25),
		24: uint16(27),
		25: uint16(29),
		26: uint16(31),
		27: uint16(33),
		28: uint16(35),
		29: uint16(37),
		30: uint16(39),
		31: uint16(41),
		32: uint16(43),
		35: uint16(60),
		36: uint16(60),
		37: uint16(60),
		38: uint16(60),
		39: uint16(68),
		40: uint16(68),
		41: uint16(68),
		42: uint16(68),
		43: uint16(68),
		44: uint16(68),
		45: uint16(68),
		46: uint16(68),
		47: uint16(68),
		48: uint16(68),
		49: uint16(68),
		50: uint16(68),
		51: uint16(68),
		53: uint16(60),
	},
	27: {
		1:  uint16(221),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		6:  uint16(13),
		9:  uint16(223),
		10: uint16(225),
		11: uint16(223),
		12: uint16(17),
		15: uint16(225),
		16: uint16(223),
		17: uint16(19),
		19: uint16(21),
		21: uint16(23),
		23: uint16(25),
		24: uint16(27),
		25: uint16(29),
		26: uint16(31),
		27: uint16(33),
		28: uint16(35),
		29: uint16(37),
		30: uint16(39),
		31: uint16(41),
		32: uint16(43),
		35: uint16(26),
		36: uint16(26),
		37: uint16(26),
		38: uint16(26),
		39: uint16(81),
		40: uint16(81),
		41: uint16(81),
		42: uint16(81),
		43: uint16(81),
		44: uint16(81),
		45: uint16(81),
		46: uint16(81),
		47: uint16(81),
		48: uint16(81),
		49: uint16(81),
		50: uint16(81),
		51: uint16(81),
		53: uint16(26),
	},
	28: {
		1:  uint16(227),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		6:  uint16(13),
		9:  uint16(229),
		10: uint16(231),
		11: uint16(229),
		12: uint16(17),
		15: uint16(231),
		16: uint16(229),
		17: uint16(19),
		19: uint16(21),
		21: uint16(23),
		23: uint16(25),
		24: uint16(27),
		25: uint16(29),
		26: uint16(31),
		27: uint16(33),
		28: uint16(35),
		29: uint16(37),
		30: uint16(39),
		31: uint16(41),
		32: uint16(43),
		35: uint16(55),
		36: uint16(55),
		37: uint16(55),
		38: uint16(55),
		39: uint16(72),
		40: uint16(72),
		41: uint16(72),
		42: uint16(72),
		43: uint16(72),
		44: uint16(72),
		45: uint16(72),
		46: uint16(72),
		47: uint16(72),
		48: uint16(72),
		49: uint16(72),
		50: uint16(72),
		51: uint16(72),
		53: uint16(55),
	},
	29: {
		1:  uint16(233),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		6:  uint16(13),
		9:  uint16(235),
		10: uint16(237),
		11: uint16(235),
		12: uint16(17),
		15: uint16(237),
		16: uint16(235),
		17: uint16(19),
		19: uint16(21),
		21: uint16(23),
		23: uint16(25),
		24: uint16(27),
		25: uint16(29),
		26: uint16(31),
		27: uint16(33),
		28: uint16(35),
		29: uint16(37),
		30: uint16(39),
		31: uint16(41),
		32: uint16(43),
		35: uint16(56),
		36: uint16(56),
		37: uint16(56),
		38: uint16(56),
		39: uint16(74),
		40: uint16(74),
		41: uint16(74),
		42: uint16(74),
		43: uint16(74),
		44: uint16(74),
		45: uint16(74),
		46: uint16(74),
		47: uint16(74),
		48: uint16(74),
		49: uint16(74),
		50: uint16(74),
		51: uint16(74),
		53: uint16(56),
	},
	30: {
		1:  uint16(239),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		6:  uint16(13),
		9:  uint16(241),
		10: uint16(243),
		11: uint16(241),
		12: uint16(17),
		15: uint16(243),
		16: uint16(241),
		17: uint16(19),
		19: uint16(21),
		21: uint16(23),
		23: uint16(25),
		24: uint16(27),
		25: uint16(29),
		26: uint16(31),
		27: uint16(33),
		28: uint16(35),
		29: uint16(37),
		30: uint16(39),
		31: uint16(41),
		32: uint16(43),
		35: uint16(50),
		36: uint16(50),
		37: uint16(50),
		38: uint16(50),
		39: uint16(73),
		40: uint16(73),
		41: uint16(73),
		42: uint16(73),
		43: uint16(73),
		44: uint16(73),
		45: uint16(73),
		46: uint16(73),
		47: uint16(73),
		48: uint16(73),
		49: uint16(73),
		50: uint16(73),
		51: uint16(73),
		53: uint16(50),
	},
	31: {
		1:  uint16(245),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		6:  uint16(13),
		9:  uint16(247),
		10: uint16(249),
		11: uint16(247),
		12: uint16(183),
		15: uint16(249),
		16: uint16(247),
		17: uint16(185),
		19: uint16(187),
		21: uint16(189),
		23: uint16(191),
		24: uint16(193),
		25: uint16(195),
		26: uint16(197),
		27: uint16(199),
		28: uint16(201),
		29: uint16(203),
		30: uint16(205),
		31: uint16(207),
		32: uint16(209),
		35: uint16(40),
		36: uint16(40),
		37: uint16(40),
		38: uint16(40),
		39: uint16(124),
		40: uint16(124),
		41: uint16(124),
		42: uint16(124),
		43: uint16(124),
		44: uint16(124),
		45: uint16(124),
		46: uint16(124),
		47: uint16(124),
		48: uint16(124),
		49: uint16(124),
		50: uint16(124),
		51: uint16(124),
		53: uint16(40),
	},
	32: {
		1:  uint16(251),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		6:  uint16(13),
		9:  uint16(253),
		10: uint16(255),
		11: uint16(253),
		12: uint16(183),
		15: uint16(255),
		16: uint16(253),
		17: uint16(185),
		19: uint16(187),
		21: uint16(189),
		23: uint16(191),
		24: uint16(193),
		25: uint16(195),
		26: uint16(197),
		27: uint16(199),
		28: uint16(201),
		29: uint16(203),
		30: uint16(205),
		31: uint16(207),
		32: uint16(209),
		35: uint16(41),
		36: uint16(41),
		37: uint16(41),
		38: uint16(41),
		39: uint16(126),
		40: uint16(126),
		41: uint16(126),
		42: uint16(126),
		43: uint16(126),
		44: uint16(126),
		45: uint16(126),
		46: uint16(126),
		47: uint16(126),
		48: uint16(126),
		49: uint16(126),
		50: uint16(126),
		51: uint16(126),
		53: uint16(41),
	},
	33: {
		1:  uint16(257),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		6:  uint16(13),
		9:  uint16(259),
		10: uint16(261),
		11: uint16(259),
		12: uint16(183),
		15: uint16(261),
		16: uint16(259),
		17: uint16(185),
		19: uint16(187),
		21: uint16(189),
		23: uint16(191),
		24: uint16(193),
		25: uint16(195),
		26: uint16(197),
		27: uint16(199),
		28: uint16(201),
		29: uint16(203),
		30: uint16(205),
		31: uint16(207),
		32: uint16(209),
		35: uint16(42),
		36: uint16(42),
		37: uint16(42),
		38: uint16(42),
		39: uint16(97),
		40: uint16(97),
		41: uint16(97),
		42: uint16(97),
		43: uint16(97),
		44: uint16(97),
		45: uint16(97),
		46: uint16(97),
		47: uint16(97),
		48: uint16(97),
		49: uint16(97),
		50: uint16(97),
		51: uint16(97),
		53: uint16(42),
	},
	34: {
		1:  uint16(263),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		6:  uint16(13),
		9:  uint16(265),
		10: uint16(267),
		11: uint16(265),
		12: uint16(183),
		15: uint16(267),
		16: uint16(265),
		17: uint16(185),
		19: uint16(187),
		21: uint16(189),
		23: uint16(191),
		24: uint16(193),
		25: uint16(195),
		26: uint16(197),
		27: uint16(199),
		28: uint16(201),
		29: uint16(203),
		30: uint16(205),
		31: uint16(207),
		32: uint16(209),
		35: uint16(43),
		36: uint16(43),
		37: uint16(43),
		38: uint16(43),
		39: uint16(100),
		40: uint16(100),
		41: uint16(100),
		42: uint16(100),
		43: uint16(100),
		44: uint16(100),
		45: uint16(100),
		46: uint16(100),
		47: uint16(100),
		48: uint16(100),
		49: uint16(100),
		50: uint16(100),
		51: uint16(100),
		53: uint16(43),
	},
	35: {
		1:  uint16(269),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		6:  uint16(13),
		9:  uint16(271),
		10: uint16(273),
		11: uint16(271),
		12: uint16(183),
		15: uint16(273),
		16: uint16(271),
		17: uint16(185),
		19: uint16(187),
		21: uint16(189),
		23: uint16(191),
		24: uint16(193),
		25: uint16(195),
		26: uint16(197),
		27: uint16(199),
		28: uint16(201),
		29: uint16(203),
		30: uint16(205),
		31: uint16(207),
		32: uint16(209),
		35: uint16(44),
		36: uint16(44),
		37: uint16(44),
		38: uint16(44),
		39: uint16(120),
		40: uint16(120),
		41: uint16(120),
		42: uint16(120),
		43: uint16(120),
		44: uint16(120),
		45: uint16(120),
		46: uint16(120),
		47: uint16(120),
		48: uint16(120),
		49: uint16(120),
		50: uint16(120),
		51: uint16(120),
		53: uint16(44),
	},
	36: {
		1:  uint16(275),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		6:  uint16(13),
		9:  uint16(277),
		10: uint16(279),
		11: uint16(277),
		12: uint16(183),
		15: uint16(279),
		16: uint16(277),
		17: uint16(185),
		19: uint16(187),
		21: uint16(189),
		23: uint16(191),
		24: uint16(193),
		25: uint16(195),
		26: uint16(197),
		27: uint16(199),
		28: uint16(201),
		29: uint16(203),
		30: uint16(205),
		31: uint16(207),
		32: uint16(209),
		35: uint16(45),
		36: uint16(45),
		37: uint16(45),
		38: uint16(45),
		39: uint16(125),
		40: uint16(125),
		41: uint16(125),
		42: uint16(125),
		43: uint16(125),
		44: uint16(125),
		45: uint16(125),
		46: uint16(125),
		47: uint16(125),
		48: uint16(125),
		49: uint16(125),
		50: uint16(125),
		51: uint16(125),
		53: uint16(45),
	},
	37: {
		1:  uint16(281),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		6:  uint16(13),
		9:  uint16(283),
		10: uint16(285),
		11: uint16(283),
		12: uint16(183),
		15: uint16(285),
		16: uint16(283),
		17: uint16(185),
		19: uint16(187),
		21: uint16(189),
		23: uint16(191),
		24: uint16(193),
		25: uint16(195),
		26: uint16(197),
		27: uint16(199),
		28: uint16(201),
		29: uint16(203),
		30: uint16(205),
		31: uint16(207),
		32: uint16(209),
		35: uint16(46),
		36: uint16(46),
		37: uint16(46),
		38: uint16(46),
		39: uint16(99),
		40: uint16(99),
		41: uint16(99),
		42: uint16(99),
		43: uint16(99),
		44: uint16(99),
		45: uint16(99),
		46: uint16(99),
		47: uint16(99),
		48: uint16(99),
		49: uint16(99),
		50: uint16(99),
		51: uint16(99),
		53: uint16(46),
	},
	38: {
		1:  uint16(287),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		6:  uint16(13),
		9:  uint16(289),
		10: uint16(291),
		11: uint16(289),
		12: uint16(17),
		15: uint16(291),
		16: uint16(289),
		17: uint16(19),
		19: uint16(21),
		21: uint16(23),
		23: uint16(25),
		24: uint16(27),
		25: uint16(29),
		26: uint16(31),
		27: uint16(33),
		28: uint16(35),
		29: uint16(37),
		30: uint16(39),
		31: uint16(41),
		32: uint16(43),
		35: uint16(54),
		36: uint16(54),
		37: uint16(54),
		38: uint16(54),
		39: uint16(71),
		40: uint16(71),
		41: uint16(71),
		42: uint16(71),
		43: uint16(71),
		44: uint16(71),
		45: uint16(71),
		46: uint16(71),
		47: uint16(71),
		48: uint16(71),
		49: uint16(71),
		50: uint16(71),
		51: uint16(71),
		53: uint16(54),
	},
	39: {
		1:  uint16(211),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		6:  uint16(13),
		9:  uint16(293),
		10: uint16(295),
		11: uint16(293),
		12: uint16(17),
		15: uint16(295),
		16: uint16(293),
		17: uint16(19),
		19: uint16(21),
		21: uint16(23),
		23: uint16(25),
		24: uint16(27),
		25: uint16(29),
		26: uint16(31),
		27: uint16(33),
		28: uint16(35),
		29: uint16(37),
		30: uint16(39),
		31: uint16(41),
		32: uint16(43),
		35: uint16(60),
		36: uint16(60),
		37: uint16(60),
		38: uint16(60),
		39: uint16(61),
		40: uint16(61),
		41: uint16(61),
		42: uint16(61),
		43: uint16(61),
		44: uint16(61),
		45: uint16(61),
		46: uint16(61),
		47: uint16(61),
		48: uint16(61),
		49: uint16(61),
		50: uint16(61),
		51: uint16(61),
		53: uint16(60),
	},
	40: {
		1:  uint16(211),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		6:  uint16(13),
		9:  uint16(297),
		10: uint16(299),
		11: uint16(297),
		12: uint16(183),
		15: uint16(299),
		16: uint16(297),
		17: uint16(185),
		19: uint16(187),
		21: uint16(189),
		23: uint16(191),
		24: uint16(193),
		25: uint16(195),
		26: uint16(197),
		27: uint16(199),
		28: uint16(201),
		29: uint16(203),
		30: uint16(205),
		31: uint16(207),
		32: uint16(209),
		35: uint16(60),
		36: uint16(60),
		37: uint16(60),
		38: uint16(60),
		39: uint16(110),
		40: uint16(110),
		41: uint16(110),
		42: uint16(110),
		43: uint16(110),
		44: uint16(110),
		45: uint16(110),
		46: uint16(110),
		47: uint16(110),
		48: uint16(110),
		49: uint16(110),
		50: uint16(110),
		51: uint16(110),
		53: uint16(60),
	},
	41: {
		1:  uint16(211),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		6:  uint16(13),
		9:  uint16(301),
		10: uint16(303),
		11: uint16(301),
		12: uint16(183),
		15: uint16(303),
		16: uint16(301),
		17: uint16(185),
		19: uint16(187),
		21: uint16(189),
		23: uint16(191),
		24: uint16(193),
		25: uint16(195),
		26: uint16(197),
		27: uint16(199),
		28: uint16(201),
		29: uint16(203),
		30: uint16(205),
		31: uint16(207),
		32: uint16(209),
		35: uint16(60),
		36: uint16(60),
		37: uint16(60),
		38: uint16(60),
		39: uint16(111),
		40: uint16(111),
		41: uint16(111),
		42: uint16(111),
		43: uint16(111),
		44: uint16(111),
		45: uint16(111),
		46: uint16(111),
		47: uint16(111),
		48: uint16(111),
		49: uint16(111),
		50: uint16(111),
		51: uint16(111),
		53: uint16(60),
	},
	42: {
		1:  uint16(211),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		6:  uint16(13),
		9:  uint16(305),
		10: uint16(307),
		11: uint16(305),
		12: uint16(183),
		15: uint16(307),
		16: uint16(305),
		17: uint16(185),
		19: uint16(187),
		21: uint16(189),
		23: uint16(191),
		24: uint16(193),
		25: uint16(195),
		26: uint16(197),
		27: uint16(199),
		28: uint16(201),
		29: uint16(203),
		30: uint16(205),
		31: uint16(207),
		32: uint16(209),
		35: uint16(60),
		36: uint16(60),
		37: uint16(60),
		38: uint16(60),
		39: uint16(112),
		40: uint16(112),
		41: uint16(112),
		42: uint16(112),
		43: uint16(112),
		44: uint16(112),
		45: uint16(112),
		46: uint16(112),
		47: uint16(112),
		48: uint16(112),
		49: uint16(112),
		50: uint16(112),
		51: uint16(112),
		53: uint16(60),
	},
	43: {
		1:  uint16(211),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		6:  uint16(13),
		9:  uint16(309),
		10: uint16(311),
		11: uint16(309),
		12: uint16(183),
		15: uint16(311),
		16: uint16(309),
		17: uint16(185),
		19: uint16(187),
		21: uint16(189),
		23: uint16(191),
		24: uint16(193),
		25: uint16(195),
		26: uint16(197),
		27: uint16(199),
		28: uint16(201),
		29: uint16(203),
		30: uint16(205),
		31: uint16(207),
		32: uint16(209),
		35: uint16(60),
		36: uint16(60),
		37: uint16(60),
		38: uint16(60),
		39: uint16(113),
		40: uint16(113),
		41: uint16(113),
		42: uint16(113),
		43: uint16(113),
		44: uint16(113),
		45: uint16(113),
		46: uint16(113),
		47: uint16(113),
		48: uint16(113),
		49: uint16(113),
		50: uint16(113),
		51: uint16(113),
		53: uint16(60),
	},
	44: {
		1:  uint16(211),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		6:  uint16(13),
		9:  uint16(313),
		10: uint16(315),
		11: uint16(313),
		12: uint16(183),
		15: uint16(315),
		16: uint16(313),
		17: uint16(185),
		19: uint16(187),
		21: uint16(189),
		23: uint16(191),
		24: uint16(193),
		25: uint16(195),
		26: uint16(197),
		27: uint16(199),
		28: uint16(201),
		29: uint16(203),
		30: uint16(205),
		31: uint16(207),
		32: uint16(209),
		35: uint16(60),
		36: uint16(60),
		37: uint16(60),
		38: uint16(60),
		39: uint16(114),
		40: uint16(114),
		41: uint16(114),
		42: uint16(114),
		43: uint16(114),
		44: uint16(114),
		45: uint16(114),
		46: uint16(114),
		47: uint16(114),
		48: uint16(114),
		49: uint16(114),
		50: uint16(114),
		51: uint16(114),
		53: uint16(60),
	},
	45: {
		1:  uint16(211),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		6:  uint16(13),
		9:  uint16(317),
		10: uint16(319),
		11: uint16(317),
		12: uint16(183),
		15: uint16(319),
		16: uint16(317),
		17: uint16(185),
		19: uint16(187),
		21: uint16(189),
		23: uint16(191),
		24: uint16(193),
		25: uint16(195),
		26: uint16(197),
		27: uint16(199),
		28: uint16(201),
		29: uint16(203),
		30: uint16(205),
		31: uint16(207),
		32: uint16(209),
		35: uint16(60),
		36: uint16(60),
		37: uint16(60),
		38: uint16(60),
		39: uint16(115),
		40: uint16(115),
		41: uint16(115),
		42: uint16(115),
		43: uint16(115),
		44: uint16(115),
		45: uint16(115),
		46: uint16(115),
		47: uint16(115),
		48: uint16(115),
		49: uint16(115),
		50: uint16(115),
		51: uint16(115),
		53: uint16(60),
	},
	46: {
		1:  uint16(211),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		6:  uint16(13),
		9:  uint16(321),
		10: uint16(323),
		11: uint16(321),
		12: uint16(183),
		15: uint16(323),
		16: uint16(321),
		17: uint16(185),
		19: uint16(187),
		21: uint16(189),
		23: uint16(191),
		24: uint16(193),
		25: uint16(195),
		26: uint16(197),
		27: uint16(199),
		28: uint16(201),
		29: uint16(203),
		30: uint16(205),
		31: uint16(207),
		32: uint16(209),
		35: uint16(60),
		36: uint16(60),
		37: uint16(60),
		38: uint16(60),
		39: uint16(116),
		40: uint16(116),
		41: uint16(116),
		42: uint16(116),
		43: uint16(116),
		44: uint16(116),
		45: uint16(116),
		46: uint16(116),
		47: uint16(116),
		48: uint16(116),
		49: uint16(116),
		50: uint16(116),
		51: uint16(116),
		53: uint16(60),
	},
	47: {
		1:  uint16(211),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		6:  uint16(13),
		9:  uint16(325),
		10: uint16(327),
		11: uint16(325),
		12: uint16(183),
		15: uint16(327),
		16: uint16(325),
		17: uint16(185),
		19: uint16(187),
		21: uint16(189),
		23: uint16(191),
		24: uint16(193),
		25: uint16(195),
		26: uint16(197),
		27: uint16(199),
		28: uint16(201),
		29: uint16(203),
		30: uint16(205),
		31: uint16(207),
		32: uint16(209),
		35: uint16(60),
		36: uint16(60),
		37: uint16(60),
		38: uint16(60),
		39: uint16(117),
		40: uint16(117),
		41: uint16(117),
		42: uint16(117),
		43: uint16(117),
		44: uint16(117),
		45: uint16(117),
		46: uint16(117),
		47: uint16(117),
		48: uint16(117),
		49: uint16(117),
		50: uint16(117),
		51: uint16(117),
		53: uint16(60),
	},
	48: {
		1:  uint16(329),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		6:  uint16(13),
		9:  uint16(331),
		10: uint16(333),
		11: uint16(331),
		12: uint16(183),
		15: uint16(333),
		16: uint16(331),
		17: uint16(185),
		19: uint16(187),
		21: uint16(189),
		23: uint16(191),
		24: uint16(193),
		25: uint16(195),
		26: uint16(197),
		27: uint16(199),
		28: uint16(201),
		29: uint16(203),
		30: uint16(205),
		31: uint16(207),
		32: uint16(209),
		35: uint16(59),
		36: uint16(59),
		37: uint16(59),
		38: uint16(59),
		39: uint16(98),
		40: uint16(98),
		41: uint16(98),
		42: uint16(98),
		43: uint16(98),
		44: uint16(98),
		45: uint16(98),
		46: uint16(98),
		47: uint16(98),
		48: uint16(98),
		49: uint16(98),
		50: uint16(98),
		51: uint16(98),
		53: uint16(59),
	},
	49: {
		1:  uint16(335),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		6:  uint16(13),
		9:  uint16(337),
		10: uint16(339),
		11: uint16(337),
		12: uint16(17),
		15: uint16(339),
		16: uint16(337),
		17: uint16(19),
		19: uint16(21),
		21: uint16(23),
		23: uint16(25),
		24: uint16(27),
		25: uint16(29),
		26: uint16(31),
		27: uint16(33),
		28: uint16(35),
		29: uint16(37),
		30: uint16(39),
		31: uint16(41),
		32: uint16(43),
		35: uint16(57),
		36: uint16(57),
		37: uint16(57),
		38: uint16(57),
		39: uint16(75),
		40: uint16(75),
		41: uint16(75),
		42: uint16(75),
		43: uint16(75),
		44: uint16(75),
		45: uint16(75),
		46: uint16(75),
		47: uint16(75),
		48: uint16(75),
		49: uint16(75),
		50: uint16(75),
		51: uint16(75),
		53: uint16(57),
	},
	50: {
		1:  uint16(211),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		6:  uint16(13),
		9:  uint16(341),
		10: uint16(343),
		11: uint16(341),
		12: uint16(17),
		15: uint16(343),
		16: uint16(341),
		17: uint16(19),
		19: uint16(21),
		21: uint16(23),
		23: uint16(25),
		24: uint16(27),
		25: uint16(29),
		26: uint16(31),
		27: uint16(33),
		28: uint16(35),
		29: uint16(37),
		30: uint16(39),
		31: uint16(41),
		32: uint16(43),
		35: uint16(60),
		36: uint16(60),
		37: uint16(60),
		38: uint16(60),
		39: uint16(86),
		40: uint16(86),
		41: uint16(86),
		42: uint16(86),
		43: uint16(86),
		44: uint16(86),
		45: uint16(86),
		46: uint16(86),
		47: uint16(86),
		48: uint16(86),
		49: uint16(86),
		50: uint16(86),
		51: uint16(86),
		53: uint16(60),
	},
	51: {
		1:  uint16(345),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		6:  uint16(13),
		9:  uint16(347),
		10: uint16(349),
		11: uint16(347),
		12: uint16(17),
		15: uint16(349),
		16: uint16(347),
		17: uint16(19),
		19: uint16(21),
		21: uint16(23),
		23: uint16(25),
		24: uint16(27),
		25: uint16(29),
		26: uint16(31),
		27: uint16(33),
		28: uint16(35),
		29: uint16(37),
		30: uint16(39),
		31: uint16(41),
		32: uint16(43),
		35: uint16(58),
		36: uint16(58),
		37: uint16(58),
		38: uint16(58),
		39: uint16(77),
		40: uint16(77),
		41: uint16(77),
		42: uint16(77),
		43: uint16(77),
		44: uint16(77),
		45: uint16(77),
		46: uint16(77),
		47: uint16(77),
		48: uint16(77),
		49: uint16(77),
		50: uint16(77),
		51: uint16(77),
		53: uint16(58),
	},
	52: {
		1:  uint16(351),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		6:  uint16(13),
		9:  uint16(353),
		10: uint16(355),
		11: uint16(353),
		12: uint16(17),
		15: uint16(355),
		16: uint16(353),
		17: uint16(19),
		19: uint16(21),
		21: uint16(23),
		23: uint16(25),
		24: uint16(27),
		25: uint16(29),
		26: uint16(31),
		27: uint16(33),
		28: uint16(35),
		29: uint16(37),
		30: uint16(39),
		31: uint16(41),
		32: uint16(43),
		35: uint16(39),
		36: uint16(39),
		37: uint16(39),
		38: uint16(39),
		39: uint16(78),
		40: uint16(78),
		41: uint16(78),
		42: uint16(78),
		43: uint16(78),
		44: uint16(78),
		45: uint16(78),
		46: uint16(78),
		47: uint16(78),
		48: uint16(78),
		49: uint16(78),
		50: uint16(78),
		51: uint16(78),
		53: uint16(39),
	},
	53: {
		1:  uint16(357),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		6:  uint16(13),
		9:  uint16(359),
		10: uint16(361),
		11: uint16(359),
		12: uint16(17),
		15: uint16(361),
		16: uint16(359),
		17: uint16(19),
		19: uint16(21),
		21: uint16(23),
		23: uint16(25),
		24: uint16(27),
		25: uint16(29),
		26: uint16(31),
		27: uint16(33),
		28: uint16(35),
		29: uint16(37),
		30: uint16(39),
		31: uint16(41),
		32: uint16(43),
		35: uint16(25),
		36: uint16(25),
		37: uint16(25),
		38: uint16(25),
		39: uint16(80),
		40: uint16(80),
		41: uint16(80),
		42: uint16(80),
		43: uint16(80),
		44: uint16(80),
		45: uint16(80),
		46: uint16(80),
		47: uint16(80),
		48: uint16(80),
		49: uint16(80),
		50: uint16(80),
		51: uint16(80),
		53: uint16(25),
	},
	54: {
		1:  uint16(211),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		6:  uint16(13),
		9:  uint16(363),
		10: uint16(365),
		11: uint16(363),
		12: uint16(17),
		15: uint16(365),
		16: uint16(363),
		17: uint16(19),
		19: uint16(21),
		21: uint16(23),
		23: uint16(25),
		24: uint16(27),
		25: uint16(29),
		26: uint16(31),
		27: uint16(33),
		28: uint16(35),
		29: uint16(37),
		30: uint16(39),
		31: uint16(41),
		32: uint16(43),
		35: uint16(60),
		36: uint16(60),
		37: uint16(60),
		38: uint16(60),
		39: uint16(62),
		40: uint16(62),
		41: uint16(62),
		42: uint16(62),
		43: uint16(62),
		44: uint16(62),
		45: uint16(62),
		46: uint16(62),
		47: uint16(62),
		48: uint16(62),
		49: uint16(62),
		50: uint16(62),
		51: uint16(62),
		53: uint16(60),
	},
	55: {
		1:  uint16(211),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		6:  uint16(13),
		9:  uint16(367),
		10: uint16(369),
		11: uint16(367),
		12: uint16(17),
		15: uint16(369),
		16: uint16(367),
		17: uint16(19),
		19: uint16(21),
		21: uint16(23),
		23: uint16(25),
		24: uint16(27),
		25: uint16(29),
		26: uint16(31),
		27: uint16(33),
		28: uint16(35),
		29: uint16(37),
		30: uint16(39),
		31: uint16(41),
		32: uint16(43),
		35: uint16(60),
		36: uint16(60),
		37: uint16(60),
		38: uint16(60),
		39: uint16(63),
		40: uint16(63),
		41: uint16(63),
		42: uint16(63),
		43: uint16(63),
		44: uint16(63),
		45: uint16(63),
		46: uint16(63),
		47: uint16(63),
		48: uint16(63),
		49: uint16(63),
		50: uint16(63),
		51: uint16(63),
		53: uint16(60),
	},
	56: {
		1:  uint16(211),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		6:  uint16(13),
		9:  uint16(371),
		10: uint16(373),
		11: uint16(371),
		12: uint16(17),
		15: uint16(373),
		16: uint16(371),
		17: uint16(19),
		19: uint16(21),
		21: uint16(23),
		23: uint16(25),
		24: uint16(27),
		25: uint16(29),
		26: uint16(31),
		27: uint16(33),
		28: uint16(35),
		29: uint16(37),
		30: uint16(39),
		31: uint16(41),
		32: uint16(43),
		35: uint16(60),
		36: uint16(60),
		37: uint16(60),
		38: uint16(60),
		39: uint16(64),
		40: uint16(64),
		41: uint16(64),
		42: uint16(64),
		43: uint16(64),
		44: uint16(64),
		45: uint16(64),
		46: uint16(64),
		47: uint16(64),
		48: uint16(64),
		49: uint16(64),
		50: uint16(64),
		51: uint16(64),
		53: uint16(60),
	},
	57: {
		1:  uint16(211),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		6:  uint16(13),
		9:  uint16(375),
		10: uint16(377),
		11: uint16(375),
		12: uint16(17),
		15: uint16(377),
		16: uint16(375),
		17: uint16(19),
		19: uint16(21),
		21: uint16(23),
		23: uint16(25),
		24: uint16(27),
		25: uint16(29),
		26: uint16(31),
		27: uint16(33),
		28: uint16(35),
		29: uint16(37),
		30: uint16(39),
		31: uint16(41),
		32: uint16(43),
		35: uint16(60),
		36: uint16(60),
		37: uint16(60),
		38: uint16(60),
		39: uint16(91),
		40: uint16(91),
		41: uint16(91),
		42: uint16(91),
		43: uint16(91),
		44: uint16(91),
		45: uint16(91),
		46: uint16(91),
		47: uint16(91),
		48: uint16(91),
		49: uint16(91),
		50: uint16(91),
		51: uint16(91),
		53: uint16(60),
	},
	58: {
		1:  uint16(211),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		6:  uint16(13),
		9:  uint16(379),
		10: uint16(381),
		11: uint16(379),
		12: uint16(17),
		15: uint16(381),
		16: uint16(379),
		17: uint16(19),
		19: uint16(21),
		21: uint16(23),
		23: uint16(25),
		24: uint16(27),
		25: uint16(29),
		26: uint16(31),
		27: uint16(33),
		28: uint16(35),
		29: uint16(37),
		30: uint16(39),
		31: uint16(41),
		32: uint16(43),
		35: uint16(60),
		36: uint16(60),
		37: uint16(60),
		38: uint16(60),
		39: uint16(66),
		40: uint16(66),
		41: uint16(66),
		42: uint16(66),
		43: uint16(66),
		44: uint16(66),
		45: uint16(66),
		46: uint16(66),
		47: uint16(66),
		48: uint16(66),
		49: uint16(66),
		50: uint16(66),
		51: uint16(66),
		53: uint16(60),
	},
	59: {
		1:  uint16(211),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		6:  uint16(13),
		9:  uint16(383),
		10: uint16(385),
		11: uint16(383),
		12: uint16(183),
		15: uint16(385),
		16: uint16(383),
		17: uint16(185),
		19: uint16(187),
		21: uint16(189),
		23: uint16(191),
		24: uint16(193),
		25: uint16(195),
		26: uint16(197),
		27: uint16(199),
		28: uint16(201),
		29: uint16(203),
		30: uint16(205),
		31: uint16(207),
		32: uint16(209),
		35: uint16(60),
		36: uint16(60),
		37: uint16(60),
		38: uint16(60),
		39: uint16(105),
		40: uint16(105),
		41: uint16(105),
		42: uint16(105),
		43: uint16(105),
		44: uint16(105),
		45: uint16(105),
		46: uint16(105),
		47: uint16(105),
		48: uint16(105),
		49: uint16(105),
		50: uint16(105),
		51: uint16(105),
		53: uint16(60),
	},
	60: {
		1:  uint16(387),
		2:  uint16(390),
		3:  uint16(393),
		4:  uint16(396),
		6:  uint16(399),
		9:  uint16(402),
		10: uint16(404),
		11: uint16(402),
		12: uint16(402),
		15: uint16(404),
		16: uint16(402),
		17: uint16(402),
		19: uint16(402),
		21: uint16(402),
		23: uint16(402),
		24: uint16(402),
		25: uint16(402),
		26: uint16(402),
		27: uint16(404),
		28: uint16(402),
		29: uint16(404),
		30: uint16(402),
		31: uint16(402),
		32: uint16(402),
		35: uint16(60),
		36: uint16(60),
		37: uint16(60),
		38: uint16(60),
		53: uint16(60),
	},
}

var ts_small_parse_table = [1605]uint16_t{
	0:    uint16(2),
	1:    uint16(408),
	2:    uint16(4),
	3:    uint16(sym_number),
	4:    uint16(sym_symbol),
	5:    uint16(anon_sym_COMMA),
	6:    uint16(anon_sym_POUND_COMMA),
	7:    uint16(406),
	8:    uint16(24),
	10:   uint16(aux_sym__intertoken_token1),
	11:   uint16(aux_sym_comment_token1),
	12:   uint16(anon_sym_POUND_SEMI),
	13:   uint16(anon_sym_POUND_BANG),
	14:   uint16(anon_sym_POUND_PIPE),
	15:   uint16(sym_boolean),
	16:   uint16(sym_character),
	17:   uint16(anon_sym_DQUOTE),
	18:   uint16(sym_keyword),
	19:   uint16(anon_sym_LPAREN),
	20:   uint16(anon_sym_RPAREN),
	21:   uint16(anon_sym_LBRACK),
	22:   uint16(anon_sym_RBRACK),
	23:   uint16(anon_sym_LBRACE),
	24:   uint16(anon_sym_RBRACE),
	25:   uint16(anon_sym_SQUOTE),
	26:   uint16(anon_sym_BQUOTE),
	27:   uint16(anon_sym_POUND_SQUOTE),
	28:   uint16(anon_sym_POUND_BQUOTE),
	29:   uint16(anon_sym_COMMA_AT),
	30:   uint16(anon_sym_POUND_COMMA_AT),
	31:   uint16(anon_sym_POUND_LPAREN),
	32:   uint16(anon_sym_POUNDvu8_LPAREN),
	33:   uint16(2),
	34:   uint16(412),
	35:   uint16(4),
	36:   uint16(sym_number),
	37:   uint16(sym_symbol),
	38:   uint16(anon_sym_COMMA),
	39:   uint16(anon_sym_POUND_COMMA),
	40:   uint16(410),
	41:   uint16(24),
	43:   uint16(aux_sym__intertoken_token1),
	44:   uint16(aux_sym_comment_token1),
	45:   uint16(anon_sym_POUND_SEMI),
	46:   uint16(anon_sym_POUND_BANG),
	47:   uint16(anon_sym_POUND_PIPE),
	48:   uint16(sym_boolean),
	49:   uint16(sym_character),
	50:   uint16(anon_sym_DQUOTE),
	51:   uint16(sym_keyword),
	52:   uint16(anon_sym_LPAREN),
	53:   uint16(anon_sym_RPAREN),
	54:   uint16(anon_sym_LBRACK),
	55:   uint16(anon_sym_RBRACK),
	56:   uint16(anon_sym_LBRACE),
	57:   uint16(anon_sym_RBRACE),
	58:   uint16(anon_sym_SQUOTE),
	59:   uint16(anon_sym_BQUOTE),
	60:   uint16(anon_sym_POUND_SQUOTE),
	61:   uint16(anon_sym_POUND_BQUOTE),
	62:   uint16(anon_sym_COMMA_AT),
	63:   uint16(anon_sym_POUND_COMMA_AT),
	64:   uint16(anon_sym_POUND_LPAREN),
	65:   uint16(anon_sym_POUNDvu8_LPAREN),
	66:   uint16(2),
	67:   uint16(416),
	68:   uint16(4),
	69:   uint16(sym_number),
	70:   uint16(sym_symbol),
	71:   uint16(anon_sym_COMMA),
	72:   uint16(anon_sym_POUND_COMMA),
	73:   uint16(414),
	74:   uint16(24),
	76:   uint16(aux_sym__intertoken_token1),
	77:   uint16(aux_sym_comment_token1),
	78:   uint16(anon_sym_POUND_SEMI),
	79:   uint16(anon_sym_POUND_BANG),
	80:   uint16(anon_sym_POUND_PIPE),
	81:   uint16(sym_boolean),
	82:   uint16(sym_character),
	83:   uint16(anon_sym_DQUOTE),
	84:   uint16(sym_keyword),
	85:   uint16(anon_sym_LPAREN),
	86:   uint16(anon_sym_RPAREN),
	87:   uint16(anon_sym_LBRACK),
	88:   uint16(anon_sym_RBRACK),
	89:   uint16(anon_sym_LBRACE),
	90:   uint16(anon_sym_RBRACE),
	91:   uint16(anon_sym_SQUOTE),
	92:   uint16(anon_sym_BQUOTE),
	93:   uint16(anon_sym_POUND_SQUOTE),
	94:   uint16(anon_sym_POUND_BQUOTE),
	95:   uint16(anon_sym_COMMA_AT),
	96:   uint16(anon_sym_POUND_COMMA_AT),
	97:   uint16(anon_sym_POUND_LPAREN),
	98:   uint16(anon_sym_POUNDvu8_LPAREN),
	99:   uint16(2),
	100:  uint16(420),
	101:  uint16(4),
	102:  uint16(sym_number),
	103:  uint16(sym_symbol),
	104:  uint16(anon_sym_COMMA),
	105:  uint16(anon_sym_POUND_COMMA),
	106:  uint16(418),
	107:  uint16(24),
	109:  uint16(aux_sym__intertoken_token1),
	110:  uint16(aux_sym_comment_token1),
	111:  uint16(anon_sym_POUND_SEMI),
	112:  uint16(anon_sym_POUND_BANG),
	113:  uint16(anon_sym_POUND_PIPE),
	114:  uint16(sym_boolean),
	115:  uint16(sym_character),
	116:  uint16(anon_sym_DQUOTE),
	117:  uint16(sym_keyword),
	118:  uint16(anon_sym_LPAREN),
	119:  uint16(anon_sym_RPAREN),
	120:  uint16(anon_sym_LBRACK),
	121:  uint16(anon_sym_RBRACK),
	122:  uint16(anon_sym_LBRACE),
	123:  uint16(anon_sym_RBRACE),
	124:  uint16(anon_sym_SQUOTE),
	125:  uint16(anon_sym_BQUOTE),
	126:  uint16(anon_sym_POUND_SQUOTE),
	127:  uint16(anon_sym_POUND_BQUOTE),
	128:  uint16(anon_sym_COMMA_AT),
	129:  uint16(anon_sym_POUND_COMMA_AT),
	130:  uint16(anon_sym_POUND_LPAREN),
	131:  uint16(anon_sym_POUNDvu8_LPAREN),
	132:  uint16(2),
	133:  uint16(424),
	134:  uint16(4),
	135:  uint16(sym_number),
	136:  uint16(sym_symbol),
	137:  uint16(anon_sym_COMMA),
	138:  uint16(anon_sym_POUND_COMMA),
	139:  uint16(422),
	140:  uint16(24),
	142:  uint16(aux_sym__intertoken_token1),
	143:  uint16(aux_sym_comment_token1),
	144:  uint16(anon_sym_POUND_SEMI),
	145:  uint16(anon_sym_POUND_BANG),
	146:  uint16(anon_sym_POUND_PIPE),
	147:  uint16(sym_boolean),
	148:  uint16(sym_character),
	149:  uint16(anon_sym_DQUOTE),
	150:  uint16(sym_keyword),
	151:  uint16(anon_sym_LPAREN),
	152:  uint16(anon_sym_RPAREN),
	153:  uint16(anon_sym_LBRACK),
	154:  uint16(anon_sym_RBRACK),
	155:  uint16(anon_sym_LBRACE),
	156:  uint16(anon_sym_RBRACE),
	157:  uint16(anon_sym_SQUOTE),
	158:  uint16(anon_sym_BQUOTE),
	159:  uint16(anon_sym_POUND_SQUOTE),
	160:  uint16(anon_sym_POUND_BQUOTE),
	161:  uint16(anon_sym_COMMA_AT),
	162:  uint16(anon_sym_POUND_COMMA_AT),
	163:  uint16(anon_sym_POUND_LPAREN),
	164:  uint16(anon_sym_POUNDvu8_LPAREN),
	165:  uint16(2),
	166:  uint16(428),
	167:  uint16(4),
	168:  uint16(sym_number),
	169:  uint16(sym_symbol),
	170:  uint16(anon_sym_COMMA),
	171:  uint16(anon_sym_POUND_COMMA),
	172:  uint16(426),
	173:  uint16(24),
	175:  uint16(aux_sym__intertoken_token1),
	176:  uint16(aux_sym_comment_token1),
	177:  uint16(anon_sym_POUND_SEMI),
	178:  uint16(anon_sym_POUND_BANG),
	179:  uint16(anon_sym_POUND_PIPE),
	180:  uint16(sym_boolean),
	181:  uint16(sym_character),
	182:  uint16(anon_sym_DQUOTE),
	183:  uint16(sym_keyword),
	184:  uint16(anon_sym_LPAREN),
	185:  uint16(anon_sym_RPAREN),
	186:  uint16(anon_sym_LBRACK),
	187:  uint16(anon_sym_RBRACK),
	188:  uint16(anon_sym_LBRACE),
	189:  uint16(anon_sym_RBRACE),
	190:  uint16(anon_sym_SQUOTE),
	191:  uint16(anon_sym_BQUOTE),
	192:  uint16(anon_sym_POUND_SQUOTE),
	193:  uint16(anon_sym_POUND_BQUOTE),
	194:  uint16(anon_sym_COMMA_AT),
	195:  uint16(anon_sym_POUND_COMMA_AT),
	196:  uint16(anon_sym_POUND_LPAREN),
	197:  uint16(anon_sym_POUNDvu8_LPAREN),
	198:  uint16(2),
	199:  uint16(432),
	200:  uint16(4),
	201:  uint16(sym_number),
	202:  uint16(sym_symbol),
	203:  uint16(anon_sym_COMMA),
	204:  uint16(anon_sym_POUND_COMMA),
	205:  uint16(430),
	206:  uint16(24),
	208:  uint16(aux_sym__intertoken_token1),
	209:  uint16(aux_sym_comment_token1),
	210:  uint16(anon_sym_POUND_SEMI),
	211:  uint16(anon_sym_POUND_BANG),
	212:  uint16(anon_sym_POUND_PIPE),
	213:  uint16(sym_boolean),
	214:  uint16(sym_character),
	215:  uint16(anon_sym_DQUOTE),
	216:  uint16(sym_keyword),
	217:  uint16(anon_sym_LPAREN),
	218:  uint16(anon_sym_RPAREN),
	219:  uint16(anon_sym_LBRACK),
	220:  uint16(anon_sym_RBRACK),
	221:  uint16(anon_sym_LBRACE),
	222:  uint16(anon_sym_RBRACE),
	223:  uint16(anon_sym_SQUOTE),
	224:  uint16(anon_sym_BQUOTE),
	225:  uint16(anon_sym_POUND_SQUOTE),
	226:  uint16(anon_sym_POUND_BQUOTE),
	227:  uint16(anon_sym_COMMA_AT),
	228:  uint16(anon_sym_POUND_COMMA_AT),
	229:  uint16(anon_sym_POUND_LPAREN),
	230:  uint16(anon_sym_POUNDvu8_LPAREN),
	231:  uint16(2),
	232:  uint16(436),
	233:  uint16(4),
	234:  uint16(sym_number),
	235:  uint16(sym_symbol),
	236:  uint16(anon_sym_COMMA),
	237:  uint16(anon_sym_POUND_COMMA),
	238:  uint16(434),
	239:  uint16(24),
	241:  uint16(aux_sym__intertoken_token1),
	242:  uint16(aux_sym_comment_token1),
	243:  uint16(anon_sym_POUND_SEMI),
	244:  uint16(anon_sym_POUND_BANG),
	245:  uint16(anon_sym_POUND_PIPE),
	246:  uint16(sym_boolean),
	247:  uint16(sym_character),
	248:  uint16(anon_sym_DQUOTE),
	249:  uint16(sym_keyword),
	250:  uint16(anon_sym_LPAREN),
	251:  uint16(anon_sym_RPAREN),
	252:  uint16(anon_sym_LBRACK),
	253:  uint16(anon_sym_RBRACK),
	254:  uint16(anon_sym_LBRACE),
	255:  uint16(anon_sym_RBRACE),
	256:  uint16(anon_sym_SQUOTE),
	257:  uint16(anon_sym_BQUOTE),
	258:  uint16(anon_sym_POUND_SQUOTE),
	259:  uint16(anon_sym_POUND_BQUOTE),
	260:  uint16(anon_sym_COMMA_AT),
	261:  uint16(anon_sym_POUND_COMMA_AT),
	262:  uint16(anon_sym_POUND_LPAREN),
	263:  uint16(anon_sym_POUNDvu8_LPAREN),
	264:  uint16(2),
	265:  uint16(440),
	266:  uint16(4),
	267:  uint16(sym_number),
	268:  uint16(sym_symbol),
	269:  uint16(anon_sym_COMMA),
	270:  uint16(anon_sym_POUND_COMMA),
	271:  uint16(438),
	272:  uint16(24),
	274:  uint16(aux_sym__intertoken_token1),
	275:  uint16(aux_sym_comment_token1),
	276:  uint16(anon_sym_POUND_SEMI),
	277:  uint16(anon_sym_POUND_BANG),
	278:  uint16(anon_sym_POUND_PIPE),
	279:  uint16(sym_boolean),
	280:  uint16(sym_character),
	281:  uint16(anon_sym_DQUOTE),
	282:  uint16(sym_keyword),
	283:  uint16(anon_sym_LPAREN),
	284:  uint16(anon_sym_RPAREN),
	285:  uint16(anon_sym_LBRACK),
	286:  uint16(anon_sym_RBRACK),
	287:  uint16(anon_sym_LBRACE),
	288:  uint16(anon_sym_RBRACE),
	289:  uint16(anon_sym_SQUOTE),
	290:  uint16(anon_sym_BQUOTE),
	291:  uint16(anon_sym_POUND_SQUOTE),
	292:  uint16(anon_sym_POUND_BQUOTE),
	293:  uint16(anon_sym_COMMA_AT),
	294:  uint16(anon_sym_POUND_COMMA_AT),
	295:  uint16(anon_sym_POUND_LPAREN),
	296:  uint16(anon_sym_POUNDvu8_LPAREN),
	297:  uint16(2),
	298:  uint16(444),
	299:  uint16(4),
	300:  uint16(sym_number),
	301:  uint16(sym_symbol),
	302:  uint16(anon_sym_COMMA),
	303:  uint16(anon_sym_POUND_COMMA),
	304:  uint16(442),
	305:  uint16(24),
	307:  uint16(aux_sym__intertoken_token1),
	308:  uint16(aux_sym_comment_token1),
	309:  uint16(anon_sym_POUND_SEMI),
	310:  uint16(anon_sym_POUND_BANG),
	311:  uint16(anon_sym_POUND_PIPE),
	312:  uint16(sym_boolean),
	313:  uint16(sym_character),
	314:  uint16(anon_sym_DQUOTE),
	315:  uint16(sym_keyword),
	316:  uint16(anon_sym_LPAREN),
	317:  uint16(anon_sym_RPAREN),
	318:  uint16(anon_sym_LBRACK),
	319:  uint16(anon_sym_RBRACK),
	320:  uint16(anon_sym_LBRACE),
	321:  uint16(anon_sym_RBRACE),
	322:  uint16(anon_sym_SQUOTE),
	323:  uint16(anon_sym_BQUOTE),
	324:  uint16(anon_sym_POUND_SQUOTE),
	325:  uint16(anon_sym_POUND_BQUOTE),
	326:  uint16(anon_sym_COMMA_AT),
	327:  uint16(anon_sym_POUND_COMMA_AT),
	328:  uint16(anon_sym_POUND_LPAREN),
	329:  uint16(anon_sym_POUNDvu8_LPAREN),
	330:  uint16(2),
	331:  uint16(448),
	332:  uint16(4),
	333:  uint16(sym_number),
	334:  uint16(sym_symbol),
	335:  uint16(anon_sym_COMMA),
	336:  uint16(anon_sym_POUND_COMMA),
	337:  uint16(446),
	338:  uint16(24),
	340:  uint16(aux_sym__intertoken_token1),
	341:  uint16(aux_sym_comment_token1),
	342:  uint16(anon_sym_POUND_SEMI),
	343:  uint16(anon_sym_POUND_BANG),
	344:  uint16(anon_sym_POUND_PIPE),
	345:  uint16(sym_boolean),
	346:  uint16(sym_character),
	347:  uint16(anon_sym_DQUOTE),
	348:  uint16(sym_keyword),
	349:  uint16(anon_sym_LPAREN),
	350:  uint16(anon_sym_RPAREN),
	351:  uint16(anon_sym_LBRACK),
	352:  uint16(anon_sym_RBRACK),
	353:  uint16(anon_sym_LBRACE),
	354:  uint16(anon_sym_RBRACE),
	355:  uint16(anon_sym_SQUOTE),
	356:  uint16(anon_sym_BQUOTE),
	357:  uint16(anon_sym_POUND_SQUOTE),
	358:  uint16(anon_sym_POUND_BQUOTE),
	359:  uint16(anon_sym_COMMA_AT),
	360:  uint16(anon_sym_POUND_COMMA_AT),
	361:  uint16(anon_sym_POUND_LPAREN),
	362:  uint16(anon_sym_POUNDvu8_LPAREN),
	363:  uint16(2),
	364:  uint16(452),
	365:  uint16(4),
	366:  uint16(sym_number),
	367:  uint16(sym_symbol),
	368:  uint16(anon_sym_COMMA),
	369:  uint16(anon_sym_POUND_COMMA),
	370:  uint16(450),
	371:  uint16(24),
	373:  uint16(aux_sym__intertoken_token1),
	374:  uint16(aux_sym_comment_token1),
	375:  uint16(anon_sym_POUND_SEMI),
	376:  uint16(anon_sym_POUND_BANG),
	377:  uint16(anon_sym_POUND_PIPE),
	378:  uint16(sym_boolean),
	379:  uint16(sym_character),
	380:  uint16(anon_sym_DQUOTE),
	381:  uint16(sym_keyword),
	382:  uint16(anon_sym_LPAREN),
	383:  uint16(anon_sym_RPAREN),
	384:  uint16(anon_sym_LBRACK),
	385:  uint16(anon_sym_RBRACK),
	386:  uint16(anon_sym_LBRACE),
	387:  uint16(anon_sym_RBRACE),
	388:  uint16(anon_sym_SQUOTE),
	389:  uint16(anon_sym_BQUOTE),
	390:  uint16(anon_sym_POUND_SQUOTE),
	391:  uint16(anon_sym_POUND_BQUOTE),
	392:  uint16(anon_sym_COMMA_AT),
	393:  uint16(anon_sym_POUND_COMMA_AT),
	394:  uint16(anon_sym_POUND_LPAREN),
	395:  uint16(anon_sym_POUNDvu8_LPAREN),
	396:  uint16(2),
	397:  uint16(456),
	398:  uint16(4),
	399:  uint16(sym_number),
	400:  uint16(sym_symbol),
	401:  uint16(anon_sym_COMMA),
	402:  uint16(anon_sym_POUND_COMMA),
	403:  uint16(454),
	404:  uint16(24),
	406:  uint16(aux_sym__intertoken_token1),
	407:  uint16(aux_sym_comment_token1),
	408:  uint16(anon_sym_POUND_SEMI),
	409:  uint16(anon_sym_POUND_BANG),
	410:  uint16(anon_sym_POUND_PIPE),
	411:  uint16(sym_boolean),
	412:  uint16(sym_character),
	413:  uint16(anon_sym_DQUOTE),
	414:  uint16(sym_keyword),
	415:  uint16(anon_sym_LPAREN),
	416:  uint16(anon_sym_RPAREN),
	417:  uint16(anon_sym_LBRACK),
	418:  uint16(anon_sym_RBRACK),
	419:  uint16(anon_sym_LBRACE),
	420:  uint16(anon_sym_RBRACE),
	421:  uint16(anon_sym_SQUOTE),
	422:  uint16(anon_sym_BQUOTE),
	423:  uint16(anon_sym_POUND_SQUOTE),
	424:  uint16(anon_sym_POUND_BQUOTE),
	425:  uint16(anon_sym_COMMA_AT),
	426:  uint16(anon_sym_POUND_COMMA_AT),
	427:  uint16(anon_sym_POUND_LPAREN),
	428:  uint16(anon_sym_POUNDvu8_LPAREN),
	429:  uint16(2),
	430:  uint16(460),
	431:  uint16(4),
	432:  uint16(sym_number),
	433:  uint16(sym_symbol),
	434:  uint16(anon_sym_COMMA),
	435:  uint16(anon_sym_POUND_COMMA),
	436:  uint16(458),
	437:  uint16(24),
	439:  uint16(aux_sym__intertoken_token1),
	440:  uint16(aux_sym_comment_token1),
	441:  uint16(anon_sym_POUND_SEMI),
	442:  uint16(anon_sym_POUND_BANG),
	443:  uint16(anon_sym_POUND_PIPE),
	444:  uint16(sym_boolean),
	445:  uint16(sym_character),
	446:  uint16(anon_sym_DQUOTE),
	447:  uint16(sym_keyword),
	448:  uint16(anon_sym_LPAREN),
	449:  uint16(anon_sym_RPAREN),
	450:  uint16(anon_sym_LBRACK),
	451:  uint16(anon_sym_RBRACK),
	452:  uint16(anon_sym_LBRACE),
	453:  uint16(anon_sym_RBRACE),
	454:  uint16(anon_sym_SQUOTE),
	455:  uint16(anon_sym_BQUOTE),
	456:  uint16(anon_sym_POUND_SQUOTE),
	457:  uint16(anon_sym_POUND_BQUOTE),
	458:  uint16(anon_sym_COMMA_AT),
	459:  uint16(anon_sym_POUND_COMMA_AT),
	460:  uint16(anon_sym_POUND_LPAREN),
	461:  uint16(anon_sym_POUNDvu8_LPAREN),
	462:  uint16(2),
	463:  uint16(464),
	464:  uint16(4),
	465:  uint16(sym_number),
	466:  uint16(sym_symbol),
	467:  uint16(anon_sym_COMMA),
	468:  uint16(anon_sym_POUND_COMMA),
	469:  uint16(462),
	470:  uint16(24),
	472:  uint16(aux_sym__intertoken_token1),
	473:  uint16(aux_sym_comment_token1),
	474:  uint16(anon_sym_POUND_SEMI),
	475:  uint16(anon_sym_POUND_BANG),
	476:  uint16(anon_sym_POUND_PIPE),
	477:  uint16(sym_boolean),
	478:  uint16(sym_character),
	479:  uint16(anon_sym_DQUOTE),
	480:  uint16(sym_keyword),
	481:  uint16(anon_sym_LPAREN),
	482:  uint16(anon_sym_RPAREN),
	483:  uint16(anon_sym_LBRACK),
	484:  uint16(anon_sym_RBRACK),
	485:  uint16(anon_sym_LBRACE),
	486:  uint16(anon_sym_RBRACE),
	487:  uint16(anon_sym_SQUOTE),
	488:  uint16(anon_sym_BQUOTE),
	489:  uint16(anon_sym_POUND_SQUOTE),
	490:  uint16(anon_sym_POUND_BQUOTE),
	491:  uint16(anon_sym_COMMA_AT),
	492:  uint16(anon_sym_POUND_COMMA_AT),
	493:  uint16(anon_sym_POUND_LPAREN),
	494:  uint16(anon_sym_POUNDvu8_LPAREN),
	495:  uint16(2),
	496:  uint16(468),
	497:  uint16(4),
	498:  uint16(sym_number),
	499:  uint16(sym_symbol),
	500:  uint16(anon_sym_COMMA),
	501:  uint16(anon_sym_POUND_COMMA),
	502:  uint16(466),
	503:  uint16(24),
	505:  uint16(aux_sym__intertoken_token1),
	506:  uint16(aux_sym_comment_token1),
	507:  uint16(anon_sym_POUND_SEMI),
	508:  uint16(anon_sym_POUND_BANG),
	509:  uint16(anon_sym_POUND_PIPE),
	510:  uint16(sym_boolean),
	511:  uint16(sym_character),
	512:  uint16(anon_sym_DQUOTE),
	513:  uint16(sym_keyword),
	514:  uint16(anon_sym_LPAREN),
	515:  uint16(anon_sym_RPAREN),
	516:  uint16(anon_sym_LBRACK),
	517:  uint16(anon_sym_RBRACK),
	518:  uint16(anon_sym_LBRACE),
	519:  uint16(anon_sym_RBRACE),
	520:  uint16(anon_sym_SQUOTE),
	521:  uint16(anon_sym_BQUOTE),
	522:  uint16(anon_sym_POUND_SQUOTE),
	523:  uint16(anon_sym_POUND_BQUOTE),
	524:  uint16(anon_sym_COMMA_AT),
	525:  uint16(anon_sym_POUND_COMMA_AT),
	526:  uint16(anon_sym_POUND_LPAREN),
	527:  uint16(anon_sym_POUNDvu8_LPAREN),
	528:  uint16(2),
	529:  uint16(472),
	530:  uint16(4),
	531:  uint16(sym_number),
	532:  uint16(sym_symbol),
	533:  uint16(anon_sym_COMMA),
	534:  uint16(anon_sym_POUND_COMMA),
	535:  uint16(470),
	536:  uint16(24),
	538:  uint16(aux_sym__intertoken_token1),
	539:  uint16(aux_sym_comment_token1),
	540:  uint16(anon_sym_POUND_SEMI),
	541:  uint16(anon_sym_POUND_BANG),
	542:  uint16(anon_sym_POUND_PIPE),
	543:  uint16(sym_boolean),
	544:  uint16(sym_character),
	545:  uint16(anon_sym_DQUOTE),
	546:  uint16(sym_keyword),
	547:  uint16(anon_sym_LPAREN),
	548:  uint16(anon_sym_RPAREN),
	549:  uint16(anon_sym_LBRACK),
	550:  uint16(anon_sym_RBRACK),
	551:  uint16(anon_sym_LBRACE),
	552:  uint16(anon_sym_RBRACE),
	553:  uint16(anon_sym_SQUOTE),
	554:  uint16(anon_sym_BQUOTE),
	555:  uint16(anon_sym_POUND_SQUOTE),
	556:  uint16(anon_sym_POUND_BQUOTE),
	557:  uint16(anon_sym_COMMA_AT),
	558:  uint16(anon_sym_POUND_COMMA_AT),
	559:  uint16(anon_sym_POUND_LPAREN),
	560:  uint16(anon_sym_POUNDvu8_LPAREN),
	561:  uint16(2),
	562:  uint16(476),
	563:  uint16(4),
	564:  uint16(sym_number),
	565:  uint16(sym_symbol),
	566:  uint16(anon_sym_COMMA),
	567:  uint16(anon_sym_POUND_COMMA),
	568:  uint16(474),
	569:  uint16(24),
	571:  uint16(aux_sym__intertoken_token1),
	572:  uint16(aux_sym_comment_token1),
	573:  uint16(anon_sym_POUND_SEMI),
	574:  uint16(anon_sym_POUND_BANG),
	575:  uint16(anon_sym_POUND_PIPE),
	576:  uint16(sym_boolean),
	577:  uint16(sym_character),
	578:  uint16(anon_sym_DQUOTE),
	579:  uint16(sym_keyword),
	580:  uint16(anon_sym_LPAREN),
	581:  uint16(anon_sym_RPAREN),
	582:  uint16(anon_sym_LBRACK),
	583:  uint16(anon_sym_RBRACK),
	584:  uint16(anon_sym_LBRACE),
	585:  uint16(anon_sym_RBRACE),
	586:  uint16(anon_sym_SQUOTE),
	587:  uint16(anon_sym_BQUOTE),
	588:  uint16(anon_sym_POUND_SQUOTE),
	589:  uint16(anon_sym_POUND_BQUOTE),
	590:  uint16(anon_sym_COMMA_AT),
	591:  uint16(anon_sym_POUND_COMMA_AT),
	592:  uint16(anon_sym_POUND_LPAREN),
	593:  uint16(anon_sym_POUNDvu8_LPAREN),
	594:  uint16(2),
	595:  uint16(480),
	596:  uint16(4),
	597:  uint16(sym_number),
	598:  uint16(sym_symbol),
	599:  uint16(anon_sym_COMMA),
	600:  uint16(anon_sym_POUND_COMMA),
	601:  uint16(478),
	602:  uint16(24),
	604:  uint16(aux_sym__intertoken_token1),
	605:  uint16(aux_sym_comment_token1),
	606:  uint16(anon_sym_POUND_SEMI),
	607:  uint16(anon_sym_POUND_BANG),
	608:  uint16(anon_sym_POUND_PIPE),
	609:  uint16(sym_boolean),
	610:  uint16(sym_character),
	611:  uint16(anon_sym_DQUOTE),
	612:  uint16(sym_keyword),
	613:  uint16(anon_sym_LPAREN),
	614:  uint16(anon_sym_RPAREN),
	615:  uint16(anon_sym_LBRACK),
	616:  uint16(anon_sym_RBRACK),
	617:  uint16(anon_sym_LBRACE),
	618:  uint16(anon_sym_RBRACE),
	619:  uint16(anon_sym_SQUOTE),
	620:  uint16(anon_sym_BQUOTE),
	621:  uint16(anon_sym_POUND_SQUOTE),
	622:  uint16(anon_sym_POUND_BQUOTE),
	623:  uint16(anon_sym_COMMA_AT),
	624:  uint16(anon_sym_POUND_COMMA_AT),
	625:  uint16(anon_sym_POUND_LPAREN),
	626:  uint16(anon_sym_POUNDvu8_LPAREN),
	627:  uint16(2),
	628:  uint16(484),
	629:  uint16(4),
	630:  uint16(sym_number),
	631:  uint16(sym_symbol),
	632:  uint16(anon_sym_COMMA),
	633:  uint16(anon_sym_POUND_COMMA),
	634:  uint16(482),
	635:  uint16(24),
	637:  uint16(aux_sym__intertoken_token1),
	638:  uint16(aux_sym_comment_token1),
	639:  uint16(anon_sym_POUND_SEMI),
	640:  uint16(anon_sym_POUND_BANG),
	641:  uint16(anon_sym_POUND_PIPE),
	642:  uint16(sym_boolean),
	643:  uint16(sym_character),
	644:  uint16(anon_sym_DQUOTE),
	645:  uint16(sym_keyword),
	646:  uint16(anon_sym_LPAREN),
	647:  uint16(anon_sym_RPAREN),
	648:  uint16(anon_sym_LBRACK),
	649:  uint16(anon_sym_RBRACK),
	650:  uint16(anon_sym_LBRACE),
	651:  uint16(anon_sym_RBRACE),
	652:  uint16(anon_sym_SQUOTE),
	653:  uint16(anon_sym_BQUOTE),
	654:  uint16(anon_sym_POUND_SQUOTE),
	655:  uint16(anon_sym_POUND_BQUOTE),
	656:  uint16(anon_sym_COMMA_AT),
	657:  uint16(anon_sym_POUND_COMMA_AT),
	658:  uint16(anon_sym_POUND_LPAREN),
	659:  uint16(anon_sym_POUNDvu8_LPAREN),
	660:  uint16(2),
	661:  uint16(488),
	662:  uint16(4),
	663:  uint16(sym_number),
	664:  uint16(sym_symbol),
	665:  uint16(anon_sym_COMMA),
	666:  uint16(anon_sym_POUND_COMMA),
	667:  uint16(486),
	668:  uint16(24),
	670:  uint16(aux_sym__intertoken_token1),
	671:  uint16(aux_sym_comment_token1),
	672:  uint16(anon_sym_POUND_SEMI),
	673:  uint16(anon_sym_POUND_BANG),
	674:  uint16(anon_sym_POUND_PIPE),
	675:  uint16(sym_boolean),
	676:  uint16(sym_character),
	677:  uint16(anon_sym_DQUOTE),
	678:  uint16(sym_keyword),
	679:  uint16(anon_sym_LPAREN),
	680:  uint16(anon_sym_RPAREN),
	681:  uint16(anon_sym_LBRACK),
	682:  uint16(anon_sym_RBRACK),
	683:  uint16(anon_sym_LBRACE),
	684:  uint16(anon_sym_RBRACE),
	685:  uint16(anon_sym_SQUOTE),
	686:  uint16(anon_sym_BQUOTE),
	687:  uint16(anon_sym_POUND_SQUOTE),
	688:  uint16(anon_sym_POUND_BQUOTE),
	689:  uint16(anon_sym_COMMA_AT),
	690:  uint16(anon_sym_POUND_COMMA_AT),
	691:  uint16(anon_sym_POUND_LPAREN),
	692:  uint16(anon_sym_POUNDvu8_LPAREN),
	693:  uint16(2),
	694:  uint16(492),
	695:  uint16(4),
	696:  uint16(sym_number),
	697:  uint16(sym_symbol),
	698:  uint16(anon_sym_COMMA),
	699:  uint16(anon_sym_POUND_COMMA),
	700:  uint16(490),
	701:  uint16(24),
	703:  uint16(aux_sym__intertoken_token1),
	704:  uint16(aux_sym_comment_token1),
	705:  uint16(anon_sym_POUND_SEMI),
	706:  uint16(anon_sym_POUND_BANG),
	707:  uint16(anon_sym_POUND_PIPE),
	708:  uint16(sym_boolean),
	709:  uint16(sym_character),
	710:  uint16(anon_sym_DQUOTE),
	711:  uint16(sym_keyword),
	712:  uint16(anon_sym_LPAREN),
	713:  uint16(anon_sym_RPAREN),
	714:  uint16(anon_sym_LBRACK),
	715:  uint16(anon_sym_RBRACK),
	716:  uint16(anon_sym_LBRACE),
	717:  uint16(anon_sym_RBRACE),
	718:  uint16(anon_sym_SQUOTE),
	719:  uint16(anon_sym_BQUOTE),
	720:  uint16(anon_sym_POUND_SQUOTE),
	721:  uint16(anon_sym_POUND_BQUOTE),
	722:  uint16(anon_sym_COMMA_AT),
	723:  uint16(anon_sym_POUND_COMMA_AT),
	724:  uint16(anon_sym_POUND_LPAREN),
	725:  uint16(anon_sym_POUNDvu8_LPAREN),
	726:  uint16(2),
	727:  uint16(496),
	728:  uint16(4),
	729:  uint16(sym_number),
	730:  uint16(sym_symbol),
	731:  uint16(anon_sym_COMMA),
	732:  uint16(anon_sym_POUND_COMMA),
	733:  uint16(494),
	734:  uint16(24),
	736:  uint16(aux_sym__intertoken_token1),
	737:  uint16(aux_sym_comment_token1),
	738:  uint16(anon_sym_POUND_SEMI),
	739:  uint16(anon_sym_POUND_BANG),
	740:  uint16(anon_sym_POUND_PIPE),
	741:  uint16(sym_boolean),
	742:  uint16(sym_character),
	743:  uint16(anon_sym_DQUOTE),
	744:  uint16(sym_keyword),
	745:  uint16(anon_sym_LPAREN),
	746:  uint16(anon_sym_RPAREN),
	747:  uint16(anon_sym_LBRACK),
	748:  uint16(anon_sym_RBRACK),
	749:  uint16(anon_sym_LBRACE),
	750:  uint16(anon_sym_RBRACE),
	751:  uint16(anon_sym_SQUOTE),
	752:  uint16(anon_sym_BQUOTE),
	753:  uint16(anon_sym_POUND_SQUOTE),
	754:  uint16(anon_sym_POUND_BQUOTE),
	755:  uint16(anon_sym_COMMA_AT),
	756:  uint16(anon_sym_POUND_COMMA_AT),
	757:  uint16(anon_sym_POUND_LPAREN),
	758:  uint16(anon_sym_POUNDvu8_LPAREN),
	759:  uint16(2),
	760:  uint16(500),
	761:  uint16(4),
	762:  uint16(sym_number),
	763:  uint16(sym_symbol),
	764:  uint16(anon_sym_COMMA),
	765:  uint16(anon_sym_POUND_COMMA),
	766:  uint16(498),
	767:  uint16(24),
	769:  uint16(aux_sym__intertoken_token1),
	770:  uint16(aux_sym_comment_token1),
	771:  uint16(anon_sym_POUND_SEMI),
	772:  uint16(anon_sym_POUND_BANG),
	773:  uint16(anon_sym_POUND_PIPE),
	774:  uint16(sym_boolean),
	775:  uint16(sym_character),
	776:  uint16(anon_sym_DQUOTE),
	777:  uint16(sym_keyword),
	778:  uint16(anon_sym_LPAREN),
	779:  uint16(anon_sym_RPAREN),
	780:  uint16(anon_sym_LBRACK),
	781:  uint16(anon_sym_RBRACK),
	782:  uint16(anon_sym_LBRACE),
	783:  uint16(anon_sym_RBRACE),
	784:  uint16(anon_sym_SQUOTE),
	785:  uint16(anon_sym_BQUOTE),
	786:  uint16(anon_sym_POUND_SQUOTE),
	787:  uint16(anon_sym_POUND_BQUOTE),
	788:  uint16(anon_sym_COMMA_AT),
	789:  uint16(anon_sym_POUND_COMMA_AT),
	790:  uint16(anon_sym_POUND_LPAREN),
	791:  uint16(anon_sym_POUNDvu8_LPAREN),
	792:  uint16(2),
	793:  uint16(504),
	794:  uint16(4),
	795:  uint16(sym_number),
	796:  uint16(sym_symbol),
	797:  uint16(anon_sym_COMMA),
	798:  uint16(anon_sym_POUND_COMMA),
	799:  uint16(502),
	800:  uint16(24),
	802:  uint16(aux_sym__intertoken_token1),
	803:  uint16(aux_sym_comment_token1),
	804:  uint16(anon_sym_POUND_SEMI),
	805:  uint16(anon_sym_POUND_BANG),
	806:  uint16(anon_sym_POUND_PIPE),
	807:  uint16(sym_boolean),
	808:  uint16(sym_character),
	809:  uint16(anon_sym_DQUOTE),
	810:  uint16(sym_keyword),
	811:  uint16(anon_sym_LPAREN),
	812:  uint16(anon_sym_RPAREN),
	813:  uint16(anon_sym_LBRACK),
	814:  uint16(anon_sym_RBRACK),
	815:  uint16(anon_sym_LBRACE),
	816:  uint16(anon_sym_RBRACE),
	817:  uint16(anon_sym_SQUOTE),
	818:  uint16(anon_sym_BQUOTE),
	819:  uint16(anon_sym_POUND_SQUOTE),
	820:  uint16(anon_sym_POUND_BQUOTE),
	821:  uint16(anon_sym_COMMA_AT),
	822:  uint16(anon_sym_POUND_COMMA_AT),
	823:  uint16(anon_sym_POUND_LPAREN),
	824:  uint16(anon_sym_POUNDvu8_LPAREN),
	825:  uint16(2),
	826:  uint16(508),
	827:  uint16(4),
	828:  uint16(sym_number),
	829:  uint16(sym_symbol),
	830:  uint16(anon_sym_COMMA),
	831:  uint16(anon_sym_POUND_COMMA),
	832:  uint16(506),
	833:  uint16(24),
	835:  uint16(aux_sym__intertoken_token1),
	836:  uint16(aux_sym_comment_token1),
	837:  uint16(anon_sym_POUND_SEMI),
	838:  uint16(anon_sym_POUND_BANG),
	839:  uint16(anon_sym_POUND_PIPE),
	840:  uint16(sym_boolean),
	841:  uint16(sym_character),
	842:  uint16(anon_sym_DQUOTE),
	843:  uint16(sym_keyword),
	844:  uint16(anon_sym_LPAREN),
	845:  uint16(anon_sym_RPAREN),
	846:  uint16(anon_sym_LBRACK),
	847:  uint16(anon_sym_RBRACK),
	848:  uint16(anon_sym_LBRACE),
	849:  uint16(anon_sym_RBRACE),
	850:  uint16(anon_sym_SQUOTE),
	851:  uint16(anon_sym_BQUOTE),
	852:  uint16(anon_sym_POUND_SQUOTE),
	853:  uint16(anon_sym_POUND_BQUOTE),
	854:  uint16(anon_sym_COMMA_AT),
	855:  uint16(anon_sym_POUND_COMMA_AT),
	856:  uint16(anon_sym_POUND_LPAREN),
	857:  uint16(anon_sym_POUNDvu8_LPAREN),
	858:  uint16(2),
	859:  uint16(512),
	860:  uint16(4),
	861:  uint16(sym_number),
	862:  uint16(sym_symbol),
	863:  uint16(anon_sym_COMMA),
	864:  uint16(anon_sym_POUND_COMMA),
	865:  uint16(510),
	866:  uint16(24),
	868:  uint16(aux_sym__intertoken_token1),
	869:  uint16(aux_sym_comment_token1),
	870:  uint16(anon_sym_POUND_SEMI),
	871:  uint16(anon_sym_POUND_BANG),
	872:  uint16(anon_sym_POUND_PIPE),
	873:  uint16(sym_boolean),
	874:  uint16(sym_character),
	875:  uint16(anon_sym_DQUOTE),
	876:  uint16(sym_keyword),
	877:  uint16(anon_sym_LPAREN),
	878:  uint16(anon_sym_RPAREN),
	879:  uint16(anon_sym_LBRACK),
	880:  uint16(anon_sym_RBRACK),
	881:  uint16(anon_sym_LBRACE),
	882:  uint16(anon_sym_RBRACE),
	883:  uint16(anon_sym_SQUOTE),
	884:  uint16(anon_sym_BQUOTE),
	885:  uint16(anon_sym_POUND_SQUOTE),
	886:  uint16(anon_sym_POUND_BQUOTE),
	887:  uint16(anon_sym_COMMA_AT),
	888:  uint16(anon_sym_POUND_COMMA_AT),
	889:  uint16(anon_sym_POUND_LPAREN),
	890:  uint16(anon_sym_POUNDvu8_LPAREN),
	891:  uint16(2),
	892:  uint16(516),
	893:  uint16(4),
	894:  uint16(sym_number),
	895:  uint16(sym_symbol),
	896:  uint16(anon_sym_COMMA),
	897:  uint16(anon_sym_POUND_COMMA),
	898:  uint16(514),
	899:  uint16(24),
	901:  uint16(aux_sym__intertoken_token1),
	902:  uint16(aux_sym_comment_token1),
	903:  uint16(anon_sym_POUND_SEMI),
	904:  uint16(anon_sym_POUND_BANG),
	905:  uint16(anon_sym_POUND_PIPE),
	906:  uint16(sym_boolean),
	907:  uint16(sym_character),
	908:  uint16(anon_sym_DQUOTE),
	909:  uint16(sym_keyword),
	910:  uint16(anon_sym_LPAREN),
	911:  uint16(anon_sym_RPAREN),
	912:  uint16(anon_sym_LBRACK),
	913:  uint16(anon_sym_RBRACK),
	914:  uint16(anon_sym_LBRACE),
	915:  uint16(anon_sym_RBRACE),
	916:  uint16(anon_sym_SQUOTE),
	917:  uint16(anon_sym_BQUOTE),
	918:  uint16(anon_sym_POUND_SQUOTE),
	919:  uint16(anon_sym_POUND_BQUOTE),
	920:  uint16(anon_sym_COMMA_AT),
	921:  uint16(anon_sym_POUND_COMMA_AT),
	922:  uint16(anon_sym_POUND_LPAREN),
	923:  uint16(anon_sym_POUNDvu8_LPAREN),
	924:  uint16(2),
	925:  uint16(520),
	926:  uint16(4),
	927:  uint16(sym_number),
	928:  uint16(sym_symbol),
	929:  uint16(anon_sym_COMMA),
	930:  uint16(anon_sym_POUND_COMMA),
	931:  uint16(518),
	932:  uint16(24),
	934:  uint16(aux_sym__intertoken_token1),
	935:  uint16(aux_sym_comment_token1),
	936:  uint16(anon_sym_POUND_SEMI),
	937:  uint16(anon_sym_POUND_BANG),
	938:  uint16(anon_sym_POUND_PIPE),
	939:  uint16(sym_boolean),
	940:  uint16(sym_character),
	941:  uint16(anon_sym_DQUOTE),
	942:  uint16(sym_keyword),
	943:  uint16(anon_sym_LPAREN),
	944:  uint16(anon_sym_RPAREN),
	945:  uint16(anon_sym_LBRACK),
	946:  uint16(anon_sym_RBRACK),
	947:  uint16(anon_sym_LBRACE),
	948:  uint16(anon_sym_RBRACE),
	949:  uint16(anon_sym_SQUOTE),
	950:  uint16(anon_sym_BQUOTE),
	951:  uint16(anon_sym_POUND_SQUOTE),
	952:  uint16(anon_sym_POUND_BQUOTE),
	953:  uint16(anon_sym_COMMA_AT),
	954:  uint16(anon_sym_POUND_COMMA_AT),
	955:  uint16(anon_sym_POUND_LPAREN),
	956:  uint16(anon_sym_POUNDvu8_LPAREN),
	957:  uint16(2),
	958:  uint16(524),
	959:  uint16(4),
	960:  uint16(sym_number),
	961:  uint16(sym_symbol),
	962:  uint16(anon_sym_COMMA),
	963:  uint16(anon_sym_POUND_COMMA),
	964:  uint16(522),
	965:  uint16(24),
	967:  uint16(aux_sym__intertoken_token1),
	968:  uint16(aux_sym_comment_token1),
	969:  uint16(anon_sym_POUND_SEMI),
	970:  uint16(anon_sym_POUND_BANG),
	971:  uint16(anon_sym_POUND_PIPE),
	972:  uint16(sym_boolean),
	973:  uint16(sym_character),
	974:  uint16(anon_sym_DQUOTE),
	975:  uint16(sym_keyword),
	976:  uint16(anon_sym_LPAREN),
	977:  uint16(anon_sym_RPAREN),
	978:  uint16(anon_sym_LBRACK),
	979:  uint16(anon_sym_RBRACK),
	980:  uint16(anon_sym_LBRACE),
	981:  uint16(anon_sym_RBRACE),
	982:  uint16(anon_sym_SQUOTE),
	983:  uint16(anon_sym_BQUOTE),
	984:  uint16(anon_sym_POUND_SQUOTE),
	985:  uint16(anon_sym_POUND_BQUOTE),
	986:  uint16(anon_sym_COMMA_AT),
	987:  uint16(anon_sym_POUND_COMMA_AT),
	988:  uint16(anon_sym_POUND_LPAREN),
	989:  uint16(anon_sym_POUNDvu8_LPAREN),
	990:  uint16(2),
	991:  uint16(528),
	992:  uint16(4),
	993:  uint16(sym_number),
	994:  uint16(sym_symbol),
	995:  uint16(anon_sym_COMMA),
	996:  uint16(anon_sym_POUND_COMMA),
	997:  uint16(526),
	998:  uint16(24),
	1000: uint16(aux_sym__intertoken_token1),
	1001: uint16(aux_sym_comment_token1),
	1002: uint16(anon_sym_POUND_SEMI),
	1003: uint16(anon_sym_POUND_BANG),
	1004: uint16(anon_sym_POUND_PIPE),
	1005: uint16(sym_boolean),
	1006: uint16(sym_character),
	1007: uint16(anon_sym_DQUOTE),
	1008: uint16(sym_keyword),
	1009: uint16(anon_sym_LPAREN),
	1010: uint16(anon_sym_RPAREN),
	1011: uint16(anon_sym_LBRACK),
	1012: uint16(anon_sym_RBRACK),
	1013: uint16(anon_sym_LBRACE),
	1014: uint16(anon_sym_RBRACE),
	1015: uint16(anon_sym_SQUOTE),
	1016: uint16(anon_sym_BQUOTE),
	1017: uint16(anon_sym_POUND_SQUOTE),
	1018: uint16(anon_sym_POUND_BQUOTE),
	1019: uint16(anon_sym_COMMA_AT),
	1020: uint16(anon_sym_POUND_COMMA_AT),
	1021: uint16(anon_sym_POUND_LPAREN),
	1022: uint16(anon_sym_POUNDvu8_LPAREN),
	1023: uint16(7),
	1024: uint16(530),
	1025: uint16(1),
	1026: uint16(aux_sym__intertoken_token1),
	1027: uint16(532),
	1028: uint16(1),
	1029: uint16(aux_sym_comment_token1),
	1030: uint16(534),
	1031: uint16(1),
	1032: uint16(anon_sym_POUND_SEMI),
	1033: uint16(536),
	1034: uint16(1),
	1035: uint16(anon_sym_POUND_BANG),
	1036: uint16(538),
	1037: uint16(1),
	1038: uint16(aux_sym_directive_token1),
	1039: uint16(540),
	1040: uint16(1),
	1041: uint16(anon_sym_POUND_PIPE),
	1042: uint16(94),
	1043: uint16(5),
	1044: uint16(sym__intertoken),
	1045: uint16(sym_comment),
	1046: uint16(sym_directive),
	1047: uint16(sym_block_comment),
	1048: uint16(aux_sym_comment_repeat1),
	1049: uint16(7),
	1050: uint16(532),
	1051: uint16(1),
	1052: uint16(aux_sym_comment_token1),
	1053: uint16(534),
	1054: uint16(1),
	1055: uint16(anon_sym_POUND_SEMI),
	1056: uint16(536),
	1057: uint16(1),
	1058: uint16(anon_sym_POUND_BANG),
	1059: uint16(540),
	1060: uint16(1),
	1061: uint16(anon_sym_POUND_PIPE),
	1062: uint16(542),
	1063: uint16(1),
	1064: uint16(aux_sym__intertoken_token1),
	1065: uint16(544),
	1066: uint16(1),
	1067: uint16(aux_sym_directive_token1),
	1068: uint16(92),
	1069: uint16(5),
	1070: uint16(sym__intertoken),
	1071: uint16(sym_comment),
	1072: uint16(sym_directive),
	1073: uint16(sym_block_comment),
	1074: uint16(aux_sym_comment_repeat1),
	1075: uint16(7),
	1076: uint16(402),
	1077: uint16(1),
	1078: uint16(aux_sym_directive_token1),
	1079: uint16(546),
	1080: uint16(1),
	1081: uint16(aux_sym__intertoken_token1),
	1082: uint16(549),
	1083: uint16(1),
	1084: uint16(aux_sym_comment_token1),
	1085: uint16(552),
	1086: uint16(1),
	1087: uint16(anon_sym_POUND_SEMI),
	1088: uint16(555),
	1089: uint16(1),
	1090: uint16(anon_sym_POUND_BANG),
	1091: uint16(558),
	1092: uint16(1),
	1093: uint16(anon_sym_POUND_PIPE),
	1094: uint16(94),
	1095: uint16(5),
	1096: uint16(sym__intertoken),
	1097: uint16(sym_comment),
	1098: uint16(sym_directive),
	1099: uint16(sym_block_comment),
	1100: uint16(aux_sym_comment_repeat1),
	1101: uint16(7),
	1102: uint16(532),
	1103: uint16(1),
	1104: uint16(aux_sym_comment_token1),
	1105: uint16(534),
	1106: uint16(1),
	1107: uint16(anon_sym_POUND_SEMI),
	1108: uint16(536),
	1109: uint16(1),
	1110: uint16(anon_sym_POUND_BANG),
	1111: uint16(540),
	1112: uint16(1),
	1113: uint16(anon_sym_POUND_PIPE),
	1114: uint16(561),
	1115: uint16(1),
	1116: uint16(aux_sym__intertoken_token1),
	1117: uint16(563),
	1118: uint16(1),
	1119: uint16(aux_sym_directive_token1),
	1120: uint16(96),
	1121: uint16(5),
	1122: uint16(sym__intertoken),
	1123: uint16(sym_comment),
	1124: uint16(sym_directive),
	1125: uint16(sym_block_comment),
	1126: uint16(aux_sym_comment_repeat1),
	1127: uint16(7),
	1128: uint16(530),
	1129: uint16(1),
	1130: uint16(aux_sym__intertoken_token1),
	1131: uint16(532),
	1132: uint16(1),
	1133: uint16(aux_sym_comment_token1),
	1134: uint16(534),
	1135: uint16(1),
	1136: uint16(anon_sym_POUND_SEMI),
	1137: uint16(536),
	1138: uint16(1),
	1139: uint16(anon_sym_POUND_BANG),
	1140: uint16(540),
	1141: uint16(1),
	1142: uint16(anon_sym_POUND_PIPE),
	1143: uint16(565),
	1144: uint16(1),
	1145: uint16(aux_sym_directive_token1),
	1146: uint16(94),
	1147: uint16(5),
	1148: uint16(sym__intertoken),
	1149: uint16(sym_comment),
	1150: uint16(sym_directive),
	1151: uint16(sym_block_comment),
	1152: uint16(aux_sym_comment_repeat1),
	1153: uint16(1),
	1154: uint16(458),
	1155: uint16(6),
	1156: uint16(aux_sym__intertoken_token1),
	1157: uint16(aux_sym_comment_token1),
	1158: uint16(anon_sym_POUND_SEMI),
	1159: uint16(anon_sym_POUND_BANG),
	1160: uint16(aux_sym_directive_token1),
	1161: uint16(anon_sym_POUND_PIPE),
	1162: uint16(1),
	1163: uint16(454),
	1164: uint16(6),
	1165: uint16(aux_sym__intertoken_token1),
	1166: uint16(aux_sym_comment_token1),
	1167: uint16(anon_sym_POUND_SEMI),
	1168: uint16(anon_sym_POUND_BANG),
	1169: uint16(aux_sym_directive_token1),
	1170: uint16(anon_sym_POUND_PIPE),
	1171: uint16(1),
	1172: uint16(482),
	1173: uint16(6),
	1174: uint16(aux_sym__intertoken_token1),
	1175: uint16(aux_sym_comment_token1),
	1176: uint16(anon_sym_POUND_SEMI),
	1177: uint16(anon_sym_POUND_BANG),
	1178: uint16(aux_sym_directive_token1),
	1179: uint16(anon_sym_POUND_PIPE),
	1180: uint16(1),
	1181: uint16(462),
	1182: uint16(6),
	1183: uint16(aux_sym__intertoken_token1),
	1184: uint16(aux_sym_comment_token1),
	1185: uint16(anon_sym_POUND_SEMI),
	1186: uint16(anon_sym_POUND_BANG),
	1187: uint16(aux_sym_directive_token1),
	1188: uint16(anon_sym_POUND_PIPE),
	1189: uint16(1),
	1190: uint16(486),
	1191: uint16(6),
	1192: uint16(aux_sym__intertoken_token1),
	1193: uint16(aux_sym_comment_token1),
	1194: uint16(anon_sym_POUND_SEMI),
	1195: uint16(anon_sym_POUND_BANG),
	1196: uint16(aux_sym_directive_token1),
	1197: uint16(anon_sym_POUND_PIPE),
	1198: uint16(1),
	1199: uint16(466),
	1200: uint16(6),
	1201: uint16(aux_sym__intertoken_token1),
	1202: uint16(aux_sym_comment_token1),
	1203: uint16(anon_sym_POUND_SEMI),
	1204: uint16(anon_sym_POUND_BANG),
	1205: uint16(aux_sym_directive_token1),
	1206: uint16(anon_sym_POUND_PIPE),
	1207: uint16(1),
	1208: uint16(494),
	1209: uint16(6),
	1210: uint16(aux_sym__intertoken_token1),
	1211: uint16(aux_sym_comment_token1),
	1212: uint16(anon_sym_POUND_SEMI),
	1213: uint16(anon_sym_POUND_BANG),
	1214: uint16(aux_sym_directive_token1),
	1215: uint16(anon_sym_POUND_PIPE),
	1216: uint16(1),
	1217: uint16(498),
	1218: uint16(6),
	1219: uint16(aux_sym__intertoken_token1),
	1220: uint16(aux_sym_comment_token1),
	1221: uint16(anon_sym_POUND_SEMI),
	1222: uint16(anon_sym_POUND_BANG),
	1223: uint16(aux_sym_directive_token1),
	1224: uint16(anon_sym_POUND_PIPE),
	1225: uint16(1),
	1226: uint16(506),
	1227: uint16(6),
	1228: uint16(aux_sym__intertoken_token1),
	1229: uint16(aux_sym_comment_token1),
	1230: uint16(anon_sym_POUND_SEMI),
	1231: uint16(anon_sym_POUND_BANG),
	1232: uint16(aux_sym_directive_token1),
	1233: uint16(anon_sym_POUND_PIPE),
	1234: uint16(1),
	1235: uint16(514),
	1236: uint16(6),
	1237: uint16(aux_sym__intertoken_token1),
	1238: uint16(aux_sym_comment_token1),
	1239: uint16(anon_sym_POUND_SEMI),
	1240: uint16(anon_sym_POUND_BANG),
	1241: uint16(aux_sym_directive_token1),
	1242: uint16(anon_sym_POUND_PIPE),
	1243: uint16(1),
	1244: uint16(518),
	1245: uint16(6),
	1246: uint16(aux_sym__intertoken_token1),
	1247: uint16(aux_sym_comment_token1),
	1248: uint16(anon_sym_POUND_SEMI),
	1249: uint16(anon_sym_POUND_BANG),
	1250: uint16(aux_sym_directive_token1),
	1251: uint16(anon_sym_POUND_PIPE),
	1252: uint16(1),
	1253: uint16(522),
	1254: uint16(6),
	1255: uint16(aux_sym__intertoken_token1),
	1256: uint16(aux_sym_comment_token1),
	1257: uint16(anon_sym_POUND_SEMI),
	1258: uint16(anon_sym_POUND_BANG),
	1259: uint16(aux_sym_directive_token1),
	1260: uint16(anon_sym_POUND_PIPE),
	1261: uint16(1),
	1262: uint16(422),
	1263: uint16(6),
	1264: uint16(aux_sym__intertoken_token1),
	1265: uint16(aux_sym_comment_token1),
	1266: uint16(anon_sym_POUND_SEMI),
	1267: uint16(anon_sym_POUND_BANG),
	1268: uint16(aux_sym_directive_token1),
	1269: uint16(anon_sym_POUND_PIPE),
	1270: uint16(1),
	1271: uint16(410),
	1272: uint16(6),
	1273: uint16(aux_sym__intertoken_token1),
	1274: uint16(aux_sym_comment_token1),
	1275: uint16(anon_sym_POUND_SEMI),
	1276: uint16(anon_sym_POUND_BANG),
	1277: uint16(aux_sym_directive_token1),
	1278: uint16(anon_sym_POUND_PIPE),
	1279: uint16(1),
	1280: uint16(414),
	1281: uint16(6),
	1282: uint16(aux_sym__intertoken_token1),
	1283: uint16(aux_sym_comment_token1),
	1284: uint16(anon_sym_POUND_SEMI),
	1285: uint16(anon_sym_POUND_BANG),
	1286: uint16(aux_sym_directive_token1),
	1287: uint16(anon_sym_POUND_PIPE),
	1288: uint16(1),
	1289: uint16(418),
	1290: uint16(6),
	1291: uint16(aux_sym__intertoken_token1),
	1292: uint16(aux_sym_comment_token1),
	1293: uint16(anon_sym_POUND_SEMI),
	1294: uint16(anon_sym_POUND_BANG),
	1295: uint16(aux_sym_directive_token1),
	1296: uint16(anon_sym_POUND_PIPE),
	1297: uint16(1),
	1298: uint16(526),
	1299: uint16(6),
	1300: uint16(aux_sym__intertoken_token1),
	1301: uint16(aux_sym_comment_token1),
	1302: uint16(anon_sym_POUND_SEMI),
	1303: uint16(anon_sym_POUND_BANG),
	1304: uint16(aux_sym_directive_token1),
	1305: uint16(anon_sym_POUND_PIPE),
	1306: uint16(1),
	1307: uint16(426),
	1308: uint16(6),
	1309: uint16(aux_sym__intertoken_token1),
	1310: uint16(aux_sym_comment_token1),
	1311: uint16(anon_sym_POUND_SEMI),
	1312: uint16(anon_sym_POUND_BANG),
	1313: uint16(aux_sym_directive_token1),
	1314: uint16(anon_sym_POUND_PIPE),
	1315: uint16(1),
	1316: uint16(406),
	1317: uint16(6),
	1318: uint16(aux_sym__intertoken_token1),
	1319: uint16(aux_sym_comment_token1),
	1320: uint16(anon_sym_POUND_SEMI),
	1321: uint16(anon_sym_POUND_BANG),
	1322: uint16(aux_sym_directive_token1),
	1323: uint16(anon_sym_POUND_PIPE),
	1324: uint16(1),
	1325: uint16(430),
	1326: uint16(6),
	1327: uint16(aux_sym__intertoken_token1),
	1328: uint16(aux_sym_comment_token1),
	1329: uint16(anon_sym_POUND_SEMI),
	1330: uint16(anon_sym_POUND_BANG),
	1331: uint16(aux_sym_directive_token1),
	1332: uint16(anon_sym_POUND_PIPE),
	1333: uint16(1),
	1334: uint16(434),
	1335: uint16(6),
	1336: uint16(aux_sym__intertoken_token1),
	1337: uint16(aux_sym_comment_token1),
	1338: uint16(anon_sym_POUND_SEMI),
	1339: uint16(anon_sym_POUND_BANG),
	1340: uint16(aux_sym_directive_token1),
	1341: uint16(anon_sym_POUND_PIPE),
	1342: uint16(1),
	1343: uint16(438),
	1344: uint16(6),
	1345: uint16(aux_sym__intertoken_token1),
	1346: uint16(aux_sym_comment_token1),
	1347: uint16(anon_sym_POUND_SEMI),
	1348: uint16(anon_sym_POUND_BANG),
	1349: uint16(aux_sym_directive_token1),
	1350: uint16(anon_sym_POUND_PIPE),
	1351: uint16(1),
	1352: uint16(442),
	1353: uint16(6),
	1354: uint16(aux_sym__intertoken_token1),
	1355: uint16(aux_sym_comment_token1),
	1356: uint16(anon_sym_POUND_SEMI),
	1357: uint16(anon_sym_POUND_BANG),
	1358: uint16(aux_sym_directive_token1),
	1359: uint16(anon_sym_POUND_PIPE),
	1360: uint16(1),
	1361: uint16(470),
	1362: uint16(6),
	1363: uint16(aux_sym__intertoken_token1),
	1364: uint16(aux_sym_comment_token1),
	1365: uint16(anon_sym_POUND_SEMI),
	1366: uint16(anon_sym_POUND_BANG),
	1367: uint16(aux_sym_directive_token1),
	1368: uint16(anon_sym_POUND_PIPE),
	1369: uint16(1),
	1370: uint16(478),
	1371: uint16(6),
	1372: uint16(aux_sym__intertoken_token1),
	1373: uint16(aux_sym_comment_token1),
	1374: uint16(anon_sym_POUND_SEMI),
	1375: uint16(anon_sym_POUND_BANG),
	1376: uint16(aux_sym_directive_token1),
	1377: uint16(anon_sym_POUND_PIPE),
	1378: uint16(1),
	1379: uint16(490),
	1380: uint16(6),
	1381: uint16(aux_sym__intertoken_token1),
	1382: uint16(aux_sym_comment_token1),
	1383: uint16(anon_sym_POUND_SEMI),
	1384: uint16(anon_sym_POUND_BANG),
	1385: uint16(aux_sym_directive_token1),
	1386: uint16(anon_sym_POUND_PIPE),
	1387: uint16(1),
	1388: uint16(502),
	1389: uint16(6),
	1390: uint16(aux_sym__intertoken_token1),
	1391: uint16(aux_sym_comment_token1),
	1392: uint16(anon_sym_POUND_SEMI),
	1393: uint16(anon_sym_POUND_BANG),
	1394: uint16(aux_sym_directive_token1),
	1395: uint16(anon_sym_POUND_PIPE),
	1396: uint16(1),
	1397: uint16(446),
	1398: uint16(6),
	1399: uint16(aux_sym__intertoken_token1),
	1400: uint16(aux_sym_comment_token1),
	1401: uint16(anon_sym_POUND_SEMI),
	1402: uint16(anon_sym_POUND_BANG),
	1403: uint16(aux_sym_directive_token1),
	1404: uint16(anon_sym_POUND_PIPE),
	1405: uint16(1),
	1406: uint16(474),
	1407: uint16(6),
	1408: uint16(aux_sym__intertoken_token1),
	1409: uint16(aux_sym_comment_token1),
	1410: uint16(anon_sym_POUND_SEMI),
	1411: uint16(anon_sym_POUND_BANG),
	1412: uint16(aux_sym_directive_token1),
	1413: uint16(anon_sym_POUND_PIPE),
	1414: uint16(1),
	1415: uint16(450),
	1416: uint16(6),
	1417: uint16(aux_sym__intertoken_token1),
	1418: uint16(aux_sym_comment_token1),
	1419: uint16(anon_sym_POUND_SEMI),
	1420: uint16(anon_sym_POUND_BANG),
	1421: uint16(aux_sym_directive_token1),
	1422: uint16(anon_sym_POUND_PIPE),
	1423: uint16(1),
	1424: uint16(510),
	1425: uint16(6),
	1426: uint16(aux_sym__intertoken_token1),
	1427: uint16(aux_sym_comment_token1),
	1428: uint16(anon_sym_POUND_SEMI),
	1429: uint16(anon_sym_POUND_BANG),
	1430: uint16(aux_sym_directive_token1),
	1431: uint16(anon_sym_POUND_PIPE),
	1432: uint16(4),
	1433: uint16(567),
	1434: uint16(1),
	1435: uint16(anon_sym_POUND_PIPE),
	1436: uint16(569),
	1437: uint16(1),
	1438: uint16(aux_sym_block_comment_token1),
	1439: uint16(571),
	1440: uint16(1),
	1441: uint16(anon_sym_PIPE_POUND),
	1442: uint16(130),
	1443: uint16(2),
	1444: uint16(sym_block_comment),
	1445: uint16(aux_sym_block_comment_repeat1),
	1446: uint16(4),
	1447: uint16(573),
	1448: uint16(1),
	1449: uint16(anon_sym_POUND_PIPE),
	1450: uint16(576),
	1451: uint16(1),
	1452: uint16(aux_sym_block_comment_token1),
	1453: uint16(579),
	1454: uint16(1),
	1455: uint16(anon_sym_PIPE_POUND),
	1456: uint16(129),
	1457: uint16(2),
	1458: uint16(sym_block_comment),
	1459: uint16(aux_sym_block_comment_repeat1),
	1460: uint16(4),
	1461: uint16(567),
	1462: uint16(1),
	1463: uint16(anon_sym_POUND_PIPE),
	1464: uint16(581),
	1465: uint16(1),
	1466: uint16(aux_sym_block_comment_token1),
	1467: uint16(583),
	1468: uint16(1),
	1469: uint16(anon_sym_PIPE_POUND),
	1470: uint16(129),
	1471: uint16(2),
	1472: uint16(sym_block_comment),
	1473: uint16(aux_sym_block_comment_repeat1),
	1474: uint16(4),
	1475: uint16(567),
	1476: uint16(1),
	1477: uint16(anon_sym_POUND_PIPE),
	1478: uint16(581),
	1479: uint16(1),
	1480: uint16(aux_sym_block_comment_token1),
	1481: uint16(585),
	1482: uint16(1),
	1483: uint16(anon_sym_PIPE_POUND),
	1484: uint16(129),
	1485: uint16(2),
	1486: uint16(sym_block_comment),
	1487: uint16(aux_sym_block_comment_repeat1),
	1488: uint16(4),
	1489: uint16(567),
	1490: uint16(1),
	1491: uint16(anon_sym_POUND_PIPE),
	1492: uint16(587),
	1493: uint16(1),
	1494: uint16(aux_sym_block_comment_token1),
	1495: uint16(589),
	1496: uint16(1),
	1497: uint16(anon_sym_PIPE_POUND),
	1498: uint16(131),
	1499: uint16(2),
	1500: uint16(sym_block_comment),
	1501: uint16(aux_sym_block_comment_repeat1),
	1502: uint16(4),
	1503: uint16(567),
	1504: uint16(1),
	1505: uint16(anon_sym_POUND_PIPE),
	1506: uint16(591),
	1507: uint16(1),
	1508: uint16(aux_sym_block_comment_token1),
	1509: uint16(593),
	1510: uint16(1),
	1511: uint16(anon_sym_PIPE_POUND),
	1512: uint16(134),
	1513: uint16(2),
	1514: uint16(sym_block_comment),
	1515: uint16(aux_sym_block_comment_repeat1),
	1516: uint16(4),
	1517: uint16(567),
	1518: uint16(1),
	1519: uint16(anon_sym_POUND_PIPE),
	1520: uint16(581),
	1521: uint16(1),
	1522: uint16(aux_sym_block_comment_token1),
	1523: uint16(595),
	1524: uint16(1),
	1525: uint16(anon_sym_PIPE_POUND),
	1526: uint16(129),
	1527: uint16(2),
	1528: uint16(sym_block_comment),
	1529: uint16(aux_sym_block_comment_repeat1),
	1530: uint16(3),
	1531: uint16(597),
	1532: uint16(1),
	1533: uint16(anon_sym_DQUOTE),
	1534: uint16(136),
	1535: uint16(1),
	1536: uint16(aux_sym_string_repeat1),
	1537: uint16(599),
	1538: uint16(2),
	1539: uint16(aux_sym_string_token1),
	1540: uint16(sym_escape_sequence),
	1541: uint16(3),
	1542: uint16(601),
	1543: uint16(1),
	1544: uint16(anon_sym_DQUOTE),
	1545: uint16(137),
	1546: uint16(1),
	1547: uint16(aux_sym_string_repeat1),
	1548: uint16(603),
	1549: uint16(2),
	1550: uint16(aux_sym_string_token1),
	1551: uint16(sym_escape_sequence),
	1552: uint16(3),
	1553: uint16(605),
	1554: uint16(1),
	1555: uint16(anon_sym_DQUOTE),
	1556: uint16(137),
	1557: uint16(1),
	1558: uint16(aux_sym_string_repeat1),
	1559: uint16(607),
	1560: uint16(2),
	1561: uint16(aux_sym_string_token1),
	1562: uint16(sym_escape_sequence),
	1563: uint16(3),
	1564: uint16(610),
	1565: uint16(1),
	1566: uint16(anon_sym_DQUOTE),
	1567: uint16(137),
	1568: uint16(1),
	1569: uint16(aux_sym_string_repeat1),
	1570: uint16(603),
	1571: uint16(2),
	1572: uint16(aux_sym_string_token1),
	1573: uint16(sym_escape_sequence),
	1574: uint16(3),
	1575: uint16(612),
	1576: uint16(1),
	1577: uint16(anon_sym_DQUOTE),
	1578: uint16(138),
	1579: uint16(1),
	1580: uint16(aux_sym_string_repeat1),
	1581: uint16(614),
	1582: uint16(2),
	1583: uint16(aux_sym_string_token1),
	1584: uint16(sym_escape_sequence),
	1585: uint16(2),
	1586: uint16(520),
	1587: uint16(1),
	1588: uint16(aux_sym_block_comment_token1),
	1589: uint16(518),
	1590: uint16(2),
	1591: uint16(anon_sym_POUND_PIPE),
	1592: uint16(anon_sym_PIPE_POUND),
	1593: uint16(2),
	1594: uint16(480),
	1595: uint16(1),
	1596: uint16(aux_sym_block_comment_token1),
	1597: uint16(478),
	1598: uint16(2),
	1599: uint16(anon_sym_POUND_PIPE),
	1600: uint16(anon_sym_PIPE_POUND),
	1601: uint16(1),
	1602: uint16(616),
	1603: uint16(1),
}

var ts_small_parse_table_map = [82]uint32_t{
	1:  uint32(33),
	2:  uint32(66),
	3:  uint32(99),
	4:  uint32(132),
	5:  uint32(165),
	6:  uint32(198),
	7:  uint32(231),
	8:  uint32(264),
	9:  uint32(297),
	10: uint32(330),
	11: uint32(363),
	12: uint32(396),
	13: uint32(429),
	14: uint32(462),
	15: uint32(495),
	16: uint32(528),
	17: uint32(561),
	18: uint32(594),
	19: uint32(627),
	20: uint32(660),
	21: uint32(693),
	22: uint32(726),
	23: uint32(759),
	24: uint32(792),
	25: uint32(825),
	26: uint32(858),
	27: uint32(891),
	28: uint32(924),
	29: uint32(957),
	30: uint32(990),
	31: uint32(1023),
	32: uint32(1049),
	33: uint32(1075),
	34: uint32(1101),
	35: uint32(1127),
	36: uint32(1153),
	37: uint32(1162),
	38: uint32(1171),
	39: uint32(1180),
	40: uint32(1189),
	41: uint32(1198),
	42: uint32(1207),
	43: uint32(1216),
	44: uint32(1225),
	45: uint32(1234),
	46: uint32(1243),
	47: uint32(1252),
	48: uint32(1261),
	49: uint32(1270),
	50: uint32(1279),
	51: uint32(1288),
	52: uint32(1297),
	53: uint32(1306),
	54: uint32(1315),
	55: uint32(1324),
	56: uint32(1333),
	57: uint32(1342),
	58: uint32(1351),
	59: uint32(1360),
	60: uint32(1369),
	61: uint32(1378),
	62: uint32(1387),
	63: uint32(1396),
	64: uint32(1405),
	65: uint32(1414),
	66: uint32(1423),
	67: uint32(1432),
	68: uint32(1446),
	69: uint32(1460),
	70: uint32(1474),
	71: uint32(1488),
	72: uint32(1502),
	73: uint32(1516),
	74: uint32(1530),
	75: uint32(1541),
	76: uint32(1552),
	77: uint32(1563),
	78: uint32(1574),
	79: uint32(1585),
	80: uint32(1593),
	81: uint32(1601),
}

var ts_parse_actions = [618]TSParseActionEntry{
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token: uint8(TSParseActionTypeReduce),
		Fsymbol:     uint16(sym_program),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(7)),
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
		Fstate: uint16(libc.Int32FromInt32(87)),
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
		Fstate: uint16(libc.Int32FromInt32(30)),
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
		Fstate: uint16(libc.Int32FromInt32(93)),
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
		Fstate: uint16(libc.Int32FromInt32(128)),
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
		Fstate: uint16(libc.Int32FromInt32(7)),
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
		Fstate: uint16(libc.Int32FromInt32(135)),
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
		Fstate: uint16(libc.Int32FromInt32(11)),
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
		Fstate: uint16(libc.Int32FromInt32(13)),
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
		Fstate: uint16(libc.Int32FromInt32(23)),
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
		Fstate: uint16(libc.Int32FromInt32(38)),
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
		Fstate: uint16(libc.Int32FromInt32(28)),
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
		Fstate: uint16(libc.Int32FromInt32(29)),
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
		Fstate: uint16(libc.Int32FromInt32(49)),
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
		Fstate: uint16(libc.Int32FromInt32(51)),
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
		Fstate: uint16(libc.Int32FromInt32(52)),
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
		Fstate: uint16(libc.Int32FromInt32(53)),
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
		Fstate: uint16(libc.Int32FromInt32(27)),
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
		Fstate: uint16(libc.Int32FromInt32(5)),
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
		Fstate: uint16(libc.Int32FromInt32(6)),
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
		Fsymbol:      uint16(aux_sym_program_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(2)),
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
		Fsymbol:      uint16(aux_sym_program_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(87)),
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
		Fsymbol:      uint16(aux_sym_program_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(30)),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_program_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(93)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_program_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(128)),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_program_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(2)),
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
		Fsymbol:      uint16(aux_sym_program_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(135)),
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
		Fsymbol:      uint16(aux_sym_program_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(11)),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_program_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(13)),
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
		Fsymbol:      uint16(aux_sym_program_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(23)),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_program_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(38)),
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
		Fsymbol:      uint16(aux_sym_program_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(28)),
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
		Fsymbol:      uint16(aux_sym_program_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(29)),
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
		Fsymbol:      uint16(aux_sym_program_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(49)),
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
		Fsymbol:      uint16(aux_sym_program_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(51)),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_program_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(52)),
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
		Fsymbol:      uint16(aux_sym_program_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(53)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_program_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(27)),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_program_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(5)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_program_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(6)),
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
		Fstate: uint16(libc.Int32FromInt32(2)),
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
		Fstate: uint16(libc.Int32FromInt32(2)),
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
		Fstate: uint16(libc.Int32FromInt32(69)),
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
		Fstate: uint16(libc.Int32FromInt32(119)),
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
		Fstate: uint16(libc.Int32FromInt32(3)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(3)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(83)),
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
		Fstate: uint16(libc.Int32FromInt32(12)),
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
		Fstate: uint16(libc.Int32FromInt32(12)),
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
		Fstate: uint16(libc.Int32FromInt32(84)),
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
		Fstate: uint16(libc.Int32FromInt32(65)),
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
		Fstate: uint16(libc.Int32FromInt32(8)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(8)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(85)),
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
		Fstate: uint16(libc.Int32FromInt32(70)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(9)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(9)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(19)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(123)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(20)),
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
		Fstate: uint16(libc.Int32FromInt32(20)),
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
		Fstate: uint16(libc.Int32FromInt32(21)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(21)),
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
		Fstate: uint16(libc.Int32FromInt32(22)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(22)),
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
		Fstate: uint16(libc.Int32FromInt32(103)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
	166: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(104)),
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
		Fstate: uint16(libc.Int32FromInt32(109)),
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
		Fstate: uint16(libc.Int32FromInt32(118)),
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
		Fstate: uint16(libc.Int32FromInt32(10)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(47)),
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
		Fstate: uint16(libc.Int32FromInt32(101)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(101)),
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
		Fstate: uint16(libc.Int32FromInt32(139)),
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
		Fstate: uint16(libc.Int32FromInt32(14)),
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
		Fstate: uint16(libc.Int32FromInt32(15)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(16)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	192: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	194: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(36)),
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
		Fstate: uint16(libc.Int32FromInt32(37)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(24)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(17)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(60)),
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
		Fstate: uint16(libc.Int32FromInt32(67)),
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
		Fstate: uint16(libc.Int32FromInt32(67)),
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
		Fstate: uint16(libc.Int32FromInt32(68)),
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
		Fstate: uint16(libc.Int32FromInt32(68)),
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
		Fstate: uint16(libc.Int32FromInt32(26)),
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
		Fstate: uint16(libc.Int32FromInt32(81)),
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
		Fstate: uint16(libc.Int32FromInt32(81)),
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
		Fstate: uint16(libc.Int32FromInt32(55)),
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
		Fstate: uint16(libc.Int32FromInt32(72)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(72)),
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
		Fstate: uint16(libc.Int32FromInt32(56)),
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
	236: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(74)),
	}})))),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(74)),
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
		Fstate: uint16(libc.Int32FromInt32(50)),
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
		Fstate: uint16(libc.Int32FromInt32(73)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(73)),
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
		Fstate: uint16(libc.Int32FromInt32(40)),
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
		Fstate: uint16(libc.Int32FromInt32(124)),
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
		Fstate: uint16(libc.Int32FromInt32(124)),
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
		Fstate: uint16(libc.Int32FromInt32(41)),
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
		Fstate: uint16(libc.Int32FromInt32(126)),
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
		Fstate: uint16(libc.Int32FromInt32(126)),
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
		Fstate: uint16(libc.Int32FromInt32(42)),
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
		Fstate: uint16(libc.Int32FromInt32(97)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(97)),
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
		Fstate: uint16(libc.Int32FromInt32(43)),
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
		Fstate: uint16(libc.Int32FromInt32(100)),
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
		Fstate: uint16(libc.Int32FromInt32(100)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(44)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(120)),
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
		Fstate: uint16(libc.Int32FromInt32(120)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(45)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(125)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(125)),
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
		Fstate: uint16(libc.Int32FromInt32(46)),
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
		Fstate: uint16(libc.Int32FromInt32(99)),
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
		Fstate: uint16(libc.Int32FromInt32(99)),
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
		Fstate: uint16(libc.Int32FromInt32(54)),
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
		Fstate: uint16(libc.Int32FromInt32(71)),
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
		Fstate: uint16(libc.Int32FromInt32(71)),
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
		Fstate: uint16(libc.Int32FromInt32(61)),
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
		Fstate: uint16(libc.Int32FromInt32(61)),
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
		Fstate: uint16(libc.Int32FromInt32(110)),
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
		Fstate: uint16(libc.Int32FromInt32(110)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(111)),
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
		Fstate: uint16(libc.Int32FromInt32(111)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(112)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(112)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(114)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(114)),
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
		Fstate: uint16(libc.Int32FromInt32(115)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(115)),
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
		Fstate: uint16(libc.Int32FromInt32(116)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(116)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(117)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(117)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(59)),
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
		Fstate: uint16(libc.Int32FromInt32(98)),
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
		Fstate: uint16(libc.Int32FromInt32(98)),
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
		Fstate: uint16(libc.Int32FromInt32(57)),
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
		Fstate: uint16(libc.Int32FromInt32(75)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(75)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(86)),
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
		Fstate: uint16(libc.Int32FromInt32(86)),
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
		Fstate: uint16(libc.Int32FromInt32(58)),
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
		Fstate: uint16(libc.Int32FromInt32(77)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(77)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(39)),
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
		Fstate: uint16(libc.Int32FromInt32(78)),
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
		Fstate: uint16(libc.Int32FromInt32(78)),
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
		Fstate: uint16(libc.Int32FromInt32(25)),
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
		Fstate: uint16(libc.Int32FromInt32(80)),
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
		Fcount: uint8(1),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(80)),
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
		Fstate: uint16(libc.Int32FromInt32(62)),
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
		Fcount: uint8(1),
	}})),
	366: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(62)),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(63)),
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
		Fstate: uint16(libc.Int32FromInt32(63)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(64)),
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
		Fstate: uint16(libc.Int32FromInt32(64)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(91)),
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
		Fstate: uint16(libc.Int32FromInt32(91)),
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
		Fstate: uint16(libc.Int32FromInt32(66)),
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
		Fstate: uint16(libc.Int32FromInt32(66)),
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
		Fstate: uint16(libc.Int32FromInt32(105)),
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
		Fstate: uint16(libc.Int32FromInt32(105)),
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
		Fsymbol:      uint16(aux_sym_comment_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(60)),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_comment_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(87)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_comment_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(30)),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_comment_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(93)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_comment_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(128)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_comment_repeat1),
	})))),
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
		Fcount: uint8(1),
	}})),
	405: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_comment_repeat1),
	})))),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_unquote_splicing),
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
		Fcount: uint8(1),
	}})),
	409: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_unquote_splicing),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_quote),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_quote),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	415: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_quasiquote),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_quasiquote),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_syntax),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_syntax),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_unquote),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_unquote),
	})))),
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
		Fsymbol:      uint16(sym_unsyntax),
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
		Fcount: uint8(1),
	}})),
	433: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_unsyntax),
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
		Fsymbol:      uint16(sym_unsyntax_splicing),
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
		Fcount: uint8(1),
	}})),
	437: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_unsyntax_splicing),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_vector),
	})))),
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
		Fcount: uint8(1),
	}})),
	441: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_vector),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	443: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_byte_vector),
	})))),
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
		Fcount: uint8(1),
	}})),
	445: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_byte_vector),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_quote),
	})))),
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
		Fcount: uint8(1),
	}})),
	449: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_quote),
	})))),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_quasiquote),
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
		Fcount: uint8(1),
	}})),
	453: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_quasiquote),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_comment),
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
		Fcount: uint8(1),
	}})),
	457: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_comment),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	459: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_syntax),
	})))),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_syntax),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_quasisyntax),
	})))),
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
		Fcount: uint8(1),
	}})),
	465: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_quasisyntax),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_directive),
	})))),
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
		Fcount: uint8(1),
	}})),
	469: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_directive),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_unquote),
	})))),
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
		Fcount: uint8(1),
	}})),
	473: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_unquote),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_unquote_splicing),
	})))),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_unquote_splicing),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_block_comment),
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
		Fcount: uint8(1),
	}})),
	481: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_block_comment),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_unsyntax),
	})))),
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
		Fcount: uint8(1),
	}})),
	485: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_unsyntax),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_unsyntax_splicing),
	})))),
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
		Fcount: uint8(1),
	}})),
	489: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_unsyntax_splicing),
	})))),
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
		Fcount: uint8(1),
	}})),
	493: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_vector),
	})))),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_vector),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_byte_vector),
	})))),
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
		Fcount: uint8(1),
	}})),
	501: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_byte_vector),
	})))),
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
		Fcount: uint8(1),
	}})),
	505: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_comment),
	})))),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_comment),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_comment),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_comment),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_directive),
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
		Fcount: uint8(1),
	}})),
	517: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_directive),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_block_comment),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_block_comment),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_quasisyntax),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_quasisyntax),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(94)),
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
		Fstate: uint16(libc.Int32FromInt32(127)),
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
		Fstate: uint16(libc.Int32FromInt32(48)),
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
		Fstate: uint16(libc.Int32FromInt32(95)),
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
		Fstate: uint16(libc.Int32FromInt32(88)),
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
		Fstate: uint16(libc.Int32FromInt32(132)),
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
	543: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(76)),
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
		Fsymbol:      uint16(aux_sym_comment_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(94)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	550: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_comment_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(127)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	553: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_comment_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(48)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	556: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_comment_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(95)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	559: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_comment_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(132)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(96)),
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
		Fstate: uint16(libc.Int32FromInt32(102)),
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
		Fstate: uint16(libc.Int32FromInt32(106)),
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
		Fstate: uint16(libc.Int32FromInt32(133)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(130)),
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
		Fstate: uint16(libc.Int32FromInt32(79)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	574: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_block_comment_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(133)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	577: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_block_comment_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(129)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_block_comment_repeat1),
	})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(89)),
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
		Fstate: uint16(libc.Int32FromInt32(107)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(131)),
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
		Fstate: uint16(libc.Int32FromInt32(121)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(134)),
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
		Fstate: uint16(libc.Int32FromInt32(141)),
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
		Fstate: uint16(libc.Int32FromInt32(140)),
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
		Fstate: uint16(libc.Int32FromInt32(82)),
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
		Fstate: uint16(libc.Int32FromInt32(136)),
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
		Fstate: uint16(libc.Int32FromInt32(90)),
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
		Fstate: uint16(libc.Int32FromInt32(137)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	608: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	609: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(libc.Int32FromInt32(137)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	610: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(108)),
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
		Fstate: uint16(libc.Int32FromInt32(122)),
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
		Fstate: uint16(libc.Int32FromInt32(138)),
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
	617: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f uint8_t
		_ [7]byte
	}{f: uint8(TSParseActionTypeAccept)})),
}

func tree_sitter_scheme(tls *libc.TLS) (r uintptr) {
	return uintptr(unsafe.Pointer(&language))
}

var language = TSLanguage{
	Fversion:                   uint32(LANGUAGE_VERSION),
	Fsymbol_count:              uint32(SYMBOL_COUNT),
	Ftoken_count:               uint32(TOKEN_COUNT),
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

var __ccgo_ts1 = "end\x00_intertoken_token1\x00comment_token1\x00#;\x00#!\x00directive_token1\x00#|\x00block_comment_token1\x00|#\x00boolean\x00number\x00character\x00\"\x00string_token1\x00escape_sequence\x00symbol\x00keyword\x00(\x00)\x00[\x00]\x00{\x00}\x00'\x00`\x00#'\x00#`\x00,\x00,@\x00#,\x00#,@\x00#(\x00#vu8(\x00program\x00_token\x00_intertoken\x00comment\x00directive\x00block_comment\x00_datum\x00string\x00list\x00quote\x00quasiquote\x00syntax\x00quasisyntax\x00unquote\x00unquote_splicing\x00unsyntax\x00unsyntax_splicing\x00vector\x00byte_vector\x00program_repeat1\x00comment_repeat1\x00block_comment_repeat1\x00string_repeat1\x00"
