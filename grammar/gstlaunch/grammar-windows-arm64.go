// Code generated for windows/arm64 by 'ccgo -ignore-unsupported-alignment -ignore-unsupported-atomic-sizes -ignore-vector-functions --winapi-no-errno -ignore-link-errors -std=gnu11 -O0 -D__extension__= -D__forceinline=static inline -D__attribute__(...)= -D__declspec(x)= -D__cdecl= -D__stdcall= -D__fastcall= -D__thiscall= -D_cdecl= -D__restrict= -D__restrict__= -D__MINGW_EXTENSION= -D_X86INTRIN_H_INCLUDED -D_X86GPRINTRIN_H_INCLUDED -D_IMMINTRIN_H_INCLUDED -D_MMINTRIN_H_INCLUDED -D_XMMINTRIN_H_INCLUDED -D_EMMINTRIN_H_INCLUDED -D_PMMINTRIN_H_INCLUDED -D_MM3DNOW_H_INCLUDED -Dfunc=func_token -Dinterface=interface_token -Dselect=select_token -Dchan=chan_token -Dgo=go_token -Dmap=map_token -Dpackage=package_token -Dtype=type_token -Dvar=var_token -Dimport=import_token -Ddefer=defer_token -Dfallthrough=fallthrough_token -Drange=range_token -I C:\a\ccgo-tree-sitter\ccgo-tree-sitter\third-party\tree-sitter-gstlaunch\src -I C:\a\ccgo-tree-sitter\ccgo-tree-sitter\third-party\tree-sitter-gstlaunch -I C:\a\ccgo-tree-sitter\ccgo-tree-sitter\third-party -I C:\Users\runneradmin\.cache\workspaced\sources\github\7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2\lib\include -I C:\Users\runneradmin\.cache\workspaced\sources\github\7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2\lib\src C:\a\ccgo-tree-sitter\ccgo-tree-sitter\third-party\tree-sitter-gstlaunch\src\parser.c -o grammar.go', DO NOT EDIT.

//go:build windows && arm64

package grammar_gstlaunch

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
const FIELD_COUNT = 5
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
const LARGE_STATE_COUNT = 2
const MAX_ALIAS_SEQUENCE_LENGTH = 6
const MB_LEN_MAX = 1
const MINGW_HAS_DDK_H = 1
const MINGW_HAS_SECURE_API = 1
const PRODUCTION_ID_COUNT = 8
const PTRDIFF_MAX = "__PTRDIFF_MAX__"
const RAND_MAX = 0x7fff
const SIZE_MAX = "__SIZE_MAX__"
const STATE_COUNT = 54
const SYMBOL_COUNT = 32
const TOKEN_COUNT = 14
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

type ts_symbol_identifiers = int32

const anon_sym_BANG = 1
const anon_sym_DQUOTE = 2
const aux_sym_string_literal_token1 = 3
const anon_sym_SQUOTE = 4
const anon_sym_DOT = 5
const anon_sym_LPAREN = 6
const anon_sym_RPAREN = 7
const anon_sym_COMMA = 8
const anon_sym_EQ = 9
const sym_identifier = 10
const aux_sym_value_token1 = 11
const anon_sym_SEMI = 12
const anon_sym_SLASH = 13
const sym_pipeline = 14
const sym_fragment = 15
const sym_element = 16
const sym_string_literal = 17
const sym_bin = 18
const sym_simple_element = 19
const sym_reference = 20
const sym_property = 21
const sym_value = 22
const sym_caps = 23
const sym_cap = 24
const aux_sym_pipeline_repeat1 = 25
const aux_sym_fragment_repeat1 = 26
const aux_sym_string_literal_repeat1 = 27
const aux_sym_bin_repeat1 = 28
const aux_sym_reference_repeat1 = 29
const aux_sym_caps_repeat1 = 30
const aux_sym_cap_repeat1 = 31

var ts_symbol_names = [32]uintptr{
	0:  __ccgo_ts,
	1:  __ccgo_ts + 4,
	2:  __ccgo_ts + 6,
	3:  __ccgo_ts + 8,
	4:  __ccgo_ts + 23,
	5:  __ccgo_ts + 25,
	6:  __ccgo_ts + 27,
	7:  __ccgo_ts + 29,
	8:  __ccgo_ts + 31,
	9:  __ccgo_ts + 33,
	10: __ccgo_ts + 35,
	11: __ccgo_ts + 46,
	12: __ccgo_ts + 59,
	13: __ccgo_ts + 61,
	14: __ccgo_ts + 63,
	15: __ccgo_ts + 72,
	16: __ccgo_ts + 81,
	17: __ccgo_ts + 89,
	18: __ccgo_ts + 104,
	19: __ccgo_ts + 108,
	20: __ccgo_ts + 123,
	21: __ccgo_ts + 133,
	22: __ccgo_ts + 142,
	23: __ccgo_ts + 148,
	24: __ccgo_ts + 153,
	25: __ccgo_ts + 157,
	26: __ccgo_ts + 174,
	27: __ccgo_ts + 191,
	28: __ccgo_ts + 214,
	29: __ccgo_ts + 226,
	30: __ccgo_ts + 244,
	31: __ccgo_ts + 257,
}

var ts_symbol_map = [32]TSSymbol{
	1:  uint16(anon_sym_BANG),
	2:  uint16(anon_sym_DQUOTE),
	3:  uint16(aux_sym_string_literal_token1),
	4:  uint16(anon_sym_SQUOTE),
	5:  uint16(anon_sym_DOT),
	6:  uint16(anon_sym_LPAREN),
	7:  uint16(anon_sym_RPAREN),
	8:  uint16(anon_sym_COMMA),
	9:  uint16(anon_sym_EQ),
	10: uint16(sym_identifier),
	11: uint16(aux_sym_value_token1),
	12: uint16(anon_sym_SEMI),
	13: uint16(anon_sym_SLASH),
	14: uint16(sym_pipeline),
	15: uint16(sym_fragment),
	16: uint16(sym_element),
	17: uint16(sym_string_literal),
	18: uint16(sym_bin),
	19: uint16(sym_simple_element),
	20: uint16(sym_reference),
	21: uint16(sym_property),
	22: uint16(sym_value),
	23: uint16(sym_caps),
	24: uint16(sym_cap),
	25: uint16(aux_sym_pipeline_repeat1),
	26: uint16(aux_sym_fragment_repeat1),
	27: uint16(aux_sym_string_literal_repeat1),
	28: uint16(aux_sym_bin_repeat1),
	29: uint16(aux_sym_reference_repeat1),
	30: uint16(aux_sym_caps_repeat1),
	31: uint16(aux_sym_cap_repeat1),
}

var ts_symbol_metadata = [32]TSSymbolMetadata{
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
	11: {},
	12: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	13: {
		Fvisible: libc.BoolUint8(true1 != 0),
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
	25: {},
	26: {},
	27: {},
	28: {},
	29: {},
	30: {},
	31: {},
}

type ts_field_identifiers = int32

const field_element = 1
const field_key = 2
const field_pad = 3
const field_type = 4
const field_value = 5

var ts_field_names = [6]uintptr{
	0: libc.UintptrFromInt32(0),
	1: __ccgo_ts + 81,
	2: __ccgo_ts + 269,
	3: __ccgo_ts + 273,
	4: __ccgo_ts + 277,
	5: __ccgo_ts + 142,
}

var ts_field_map_slices = [8]TSFieldMapSlice{
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
		Flength: uint16(2),
	},
	5: {
		Findex:  uint16(6),
		Flength: uint16(2),
	},
	6: {
		Findex:  uint16(8),
		Flength: uint16(3),
	},
	7: {
		Findex:  uint16(11),
		Flength: uint16(2),
	},
}

var ts_field_map_entries = [13]TSFieldMapEntry{
	0: {
		Ffield_id: uint16(field_type),
	},
	1: {
		Ffield_id:    uint16(field_pad),
		Fchild_index: uint8(1),
	},
	2: {
		Ffield_id:    uint16(field_pad),
		Fchild_index: uint8(1),
	},
	3: {
		Ffield_id:    uint16(field_pad),
		Fchild_index: uint8(2),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	4: {
		Ffield_id: uint16(field_element),
	},
	5: {
		Ffield_id:    uint16(field_pad),
		Fchild_index: uint8(2),
	},
	6: {
		Ffield_id:  uint16(field_pad),
		Finherited: libc.BoolUint8(true1 != 0),
	},
	7: {
		Ffield_id:    uint16(field_pad),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	8: {
		Ffield_id: uint16(field_element),
	},
	9: {
		Ffield_id:    uint16(field_pad),
		Fchild_index: uint8(2),
	},
	10: {
		Ffield_id:    uint16(field_pad),
		Fchild_index: uint8(3),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	11: {
		Ffield_id: uint16(field_key),
	},
	12: {
		Ffield_id:    uint16(field_value),
		Fchild_index: uint8(2),
	},
}

var ts_alias_sequences = [8][6]TSSymbol{}

var ts_non_terminal_alias_map = [1]uint16_t{}

var ts_primary_state_ids = [54]TSStateId{
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
}

func sym_identifier_character_set_1(tls *libc.TLS, c int32_t) (r uint8) {
	var v1, v10, v100, v101, v105, v106, v107, v108, v109, v110, v112, v113, v115, v116, v117, v118, v119, v12, v122, v123, v124, v125, v126, v127, v129, v13, v130, v131, v132, v133, v139, v140, v141, v142, v143, v144, v145, v146, v148, v149, v15, v151, v152, v153, v154, v155, v158, v159, v16, v160, v161, v162, v163, v165, v166, v167, v168, v169, v17, v173, v174, v175, v176, v177, v178, v18, v180, v181, v183, v184, v185, v186, v187, v19, v190, v191, v192, v193, v194, v195, v197, v198, v199, v2, v200, v201, v206, v207, v208, v209, v210, v211, v212, v214, v215, v217, v218, v219, v22, v220, v221, v224, v225, v226, v227, v228, v229, v23, v231, v232, v233, v234, v235, v239, v24, v240, v241, v242, v243, v244, v246, v247, v249, v25, v250, v251, v252, v253, v256, v257, v258, v259, v26, v260, v261, v263, v264, v265, v266, v267, v274, v275, v276, v277, v278, v279, v28, v280, v281, v282, v284, v285, v287, v288, v289, v29, v290, v291, v294, v295, v296, v297, v298, v299, v3, v301, v302, v303, v304, v305, v309, v31, v310, v311, v312, v313, v314, v316, v317, v319, v32, v320, v321, v322, v323, v326, v327, v328, v329, v33, v330, v331, v333, v334, v335, v336, v337, v34, v342, v343, v344, v345, v346, v347, v348, v35, v350, v351, v353, v354, v355, v356, v357, v360, v361, v362, v363, v364, v365, v367, v368, v369, v370, v371, v375, v376, v377, v378, v379, v380, v382, v383, v385, v386, v387, v388, v389, v39, v392, v393, v394, v395, v396, v397, v399, v4, v40, v400, v401, v402, v403, v409, v41, v410, v411, v412, v413, v414, v415, v416, v418, v419, v42, v421, v422, v423, v424, v425, v428, v429, v43, v430, v431, v432, v433, v435, v436, v437, v438, v439, v44, v443, v444, v445, v446, v447, v448, v450, v451, v453, v454, v455, v456, v457, v46, v460, v461, v462, v463, v464, v465, v467, v468, v469, v47, v470, v471, v476, v477, v478, v479, v480, v481, v482, v484, v485, v487, v488, v489, v49, v490, v491, v494, v495, v496, v497, v498, v499, v5, v50, v501, v502, v503, v504, v505, v509, v51, v510, v511, v512, v513, v514, v516, v517, v519, v52, v520, v521, v522, v523, v526, v527, v528, v529, v53, v530, v531, v533, v534, v535, v536, v537, v56, v57, v58, v59, v6, v60, v61, v63, v64, v65, v66, v67, v7, v72, v73, v74, v75, v76, v77, v78, v8, v80, v81, v83, v84, v85, v86, v87, v9, v90, v91, v92, v93, v94, v95, v97, v98, v99 int32
	var v102, v103, v104, v11, v111, v114, v120, v121, v128, v134, v135, v136, v137, v138, v14, v147, v150, v156, v157, v164, v170, v171, v172, v179, v182, v188, v189, v196, v20, v202, v203, v204, v205, v21, v213, v216, v222, v223, v230, v236, v237, v238, v245, v248, v254, v255, v262, v268, v269, v27, v270, v271, v272, v273, v283, v286, v292, v293, v30, v300, v306, v307, v308, v315, v318, v324, v325, v332, v338, v339, v340, v341, v349, v352, v358, v359, v36, v366, v37, v372, v373, v374, v38, v381, v384, v390, v391, v398, v404, v405, v406, v407, v408, v417, v420, v426, v427, v434, v440, v441, v442, v449, v45, v452, v458, v459, v466, v472, v473, v474, v475, v48, v483, v486, v492, v493, v500, v506, v507, v508, v515, v518, v524, v525, v532, v538, v539, v54, v540, v541, v542, v543, v544, v55, v62, v68, v69, v70, v71, v79, v82, v88, v89, v96 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = v1, v10, v100, v101, v102, v103, v104, v105, v106, v107, v108, v109, v11, v110, v111, v112, v113, v114, v115, v116, v117, v118, v119, v12, v120, v121, v122, v123, v124, v125, v126, v127, v128, v129, v13, v130, v131, v132, v133, v134, v135, v136, v137, v138, v139, v14, v140, v141, v142, v143, v144, v145, v146, v147, v148, v149, v15, v150, v151, v152, v153, v154, v155, v156, v157, v158, v159, v16, v160, v161, v162, v163, v164, v165, v166, v167, v168, v169, v17, v170, v171, v172, v173, v174, v175, v176, v177, v178, v179, v18, v180, v181, v182, v183, v184, v185, v186, v187, v188, v189, v19, v190, v191, v192, v193, v194, v195, v196, v197, v198, v199, v2, v20, v200, v201, v202, v203, v204, v205, v206, v207, v208, v209, v21, v210, v211, v212, v213, v214, v215, v216, v217, v218, v219, v22, v220, v221, v222, v223, v224, v225, v226, v227, v228, v229, v23, v230, v231, v232, v233, v234, v235, v236, v237, v238, v239, v24, v240, v241, v242, v243, v244, v245, v246, v247, v248, v249, v25, v250, v251, v252, v253, v254, v255, v256, v257, v258, v259, v26, v260, v261, v262, v263, v264, v265, v266, v267, v268, v269, v27, v270, v271, v272, v273, v274, v275, v276, v277, v278, v279, v28, v280, v281, v282, v283, v284, v285, v286, v287, v288, v289, v29, v290, v291, v292, v293, v294, v295, v296, v297, v298, v299, v3, v30, v300, v301, v302, v303, v304, v305, v306, v307, v308, v309, v31, v310, v311, v312, v313, v314, v315, v316, v317, v318, v319, v32, v320, v321, v322, v323, v324, v325, v326, v327, v328, v329, v33, v330, v331, v332, v333, v334, v335, v336, v337, v338, v339, v34, v340, v341, v342, v343, v344, v345, v346, v347, v348, v349, v35, v350, v351, v352, v353, v354, v355, v356, v357, v358, v359, v36, v360, v361, v362, v363, v364, v365, v366, v367, v368, v369, v37, v370, v371, v372, v373, v374, v375, v376, v377, v378, v379, v38, v380, v381, v382, v383, v384, v385, v386, v387, v388, v389, v39, v390, v391, v392, v393, v394, v395, v396, v397, v398, v399, v4, v40, v400, v401, v402, v403, v404, v405, v406, v407, v408, v409, v41, v410, v411, v412, v413, v414, v415, v416, v417, v418, v419, v42, v420, v421, v422, v423, v424, v425, v426, v427, v428, v429, v43, v430, v431, v432, v433, v434, v435, v436, v437, v438, v439, v44, v440, v441, v442, v443, v444, v445, v446, v447, v448, v449, v45, v450, v451, v452, v453, v454, v455, v456, v457, v458, v459, v46, v460, v461, v462, v463, v464, v465, v466, v467, v468, v469, v47, v470, v471, v472, v473, v474, v475, v476, v477, v478, v479, v48, v480, v481, v482, v483, v484, v485, v486, v487, v488, v489, v49, v490, v491, v492, v493, v494, v495, v496, v497, v498, v499, v5, v50, v500, v501, v502, v503, v504, v505, v506, v507, v508, v509, v51, v510, v511, v512, v513, v514, v515, v516, v517, v518, v519, v52, v520, v521, v522, v523, v524, v525, v526, v527, v528, v529, v53, v530, v531, v532, v533, v534, v535, v536, v537, v538, v539, v54, v540, v541, v542, v543, v544, v55, v56, v57, v58, v59, v6, v60, v61, v62, v63, v64, v65, v66, v67, v68, v69, v7, v70, v71, v72, v73, v74, v75, v76, v77, v78, v79, v8, v80, v81, v82, v83, v84, v85, v86, v87, v88, v89, v9, v90, v91, v92, v93, v94, v95, v96, v97, v98, v99
	if c < int32(43514) {
		if c < int32(4193) {
			if c < int32(2707) {
				if c < int32(1994) {
					if c < int32(910) {
						if c < int32(736) {
							if c < int32(186) {
								if c < int32('a') {
									if c < int32('_') {
										v9 = libc.BoolInt32(c >= int32('A') && c <= int32('Z'))
									} else {
										v9 = libc.BoolInt32(c <= int32('_'))
									}
									v8 = v9
								} else {
									if v11 = c <= int32('z'); !v11 {
										if c < int32(181) {
											v10 = libc.BoolInt32(c == int32(170))
										} else {
											v10 = libc.BoolInt32(c <= int32(181))
										}
									}
									v8 = libc.BoolInt32(v11 || v10 != 0)
								}
								v7 = v8
							} else {
								if v14 = c <= int32(186); !v14 {
									if c < int32(248) {
										if c < int32(216) {
											v13 = libc.BoolInt32(c >= int32(192) && c <= int32(214))
										} else {
											v13 = libc.BoolInt32(c <= int32(246))
										}
										v12 = v13
									} else {
										v12 = libc.BoolInt32(c <= int32(705) || c >= int32(710) && c <= int32(721))
									}
								}
								v7 = libc.BoolInt32(v14 || v12 != 0)
							}
							v6 = v7
						} else {
							if v21 = c <= int32(740); !v21 {
								if c < int32(891) {
									if c < int32(880) {
										if c < int32(750) {
											v17 = libc.BoolInt32(c == int32(748))
										} else {
											v17 = libc.BoolInt32(c <= int32(750))
										}
										v16 = v17
									} else {
										v16 = libc.BoolInt32(c <= int32(884) || c >= int32(886) && c <= int32(887))
									}
									v15 = v16
								} else {
									if v20 = c <= int32(893); !v20 {
										if c < int32(904) {
											if c < int32(902) {
												v19 = libc.BoolInt32(c == int32(895))
											} else {
												v19 = libc.BoolInt32(c <= int32(902))
											}
											v18 = v19
										} else {
											v18 = libc.BoolInt32(c <= int32(906) || c == int32(908))
										}
									}
									v15 = libc.BoolInt32(v20 || v18 != 0)
								}
							}
							v6 = libc.BoolInt32(v21 || v15 != 0)
						}
						v5 = v6
					} else {
						if v38 = c <= int32(929); !v38 {
							if c < int32(1649) {
								if c < int32(1376) {
									if c < int32(1162) {
										if c < int32(1015) {
											v25 = libc.BoolInt32(c >= int32(931) && c <= int32(1013))
										} else {
											v25 = libc.BoolInt32(c <= int32(1153))
										}
										v24 = v25
									} else {
										if v27 = c <= int32(1327); !v27 {
											if c < int32(1369) {
												v26 = libc.BoolInt32(c >= int32(1329) && c <= int32(1366))
											} else {
												v26 = libc.BoolInt32(c <= int32(1369))
											}
										}
										v24 = libc.BoolInt32(v27 || v26 != 0)
									}
									v23 = v24
								} else {
									if v30 = c <= int32(1416); !v30 {
										if c < int32(1568) {
											if c < int32(1519) {
												v29 = libc.BoolInt32(c >= int32(1488) && c <= int32(1514))
											} else {
												v29 = libc.BoolInt32(c <= int32(1522))
											}
											v28 = v29
										} else {
											v28 = libc.BoolInt32(c <= int32(1610) || c >= int32(1646) && c <= int32(1647))
										}
									}
									v23 = libc.BoolInt32(v30 || v28 != 0)
								}
								v22 = v23
							} else {
								if v37 = c <= int32(1747); !v37 {
									if c < int32(1791) {
										if c < int32(1774) {
											if c < int32(1765) {
												v33 = libc.BoolInt32(c == int32(1749))
											} else {
												v33 = libc.BoolInt32(c <= int32(1766))
											}
											v32 = v33
										} else {
											v32 = libc.BoolInt32(c <= int32(1775) || c >= int32(1786) && c <= int32(1788))
										}
										v31 = v32
									} else {
										if v36 = c <= int32(1791); !v36 {
											if c < int32(1869) {
												if c < int32(1810) {
													v35 = libc.BoolInt32(c == int32(1808))
												} else {
													v35 = libc.BoolInt32(c <= int32(1839))
												}
												v34 = v35
											} else {
												v34 = libc.BoolInt32(c <= int32(1957) || c == int32(1969))
											}
										}
										v31 = libc.BoolInt32(v36 || v34 != 0)
									}
								}
								v22 = libc.BoolInt32(v37 || v31 != 0)
							}
						}
						v5 = libc.BoolInt32(v38 || v22 != 0)
					}
					v4 = v5
				} else {
					if v71 = c <= int32(2026); !v71 {
						if c < int32(2482) {
							if c < int32(2208) {
								if c < int32(2088) {
									if c < int32(2048) {
										if c < int32(2042) {
											v43 = libc.BoolInt32(c >= int32(2036) && c <= int32(2037))
										} else {
											v43 = libc.BoolInt32(c <= int32(2042))
										}
										v42 = v43
									} else {
										if v45 = c <= int32(2069); !v45 {
											if c < int32(2084) {
												v44 = libc.BoolInt32(c == int32(2074))
											} else {
												v44 = libc.BoolInt32(c <= int32(2084))
											}
										}
										v42 = libc.BoolInt32(v45 || v44 != 0)
									}
									v41 = v42
								} else {
									if v48 = c <= int32(2088); !v48 {
										if c < int32(2160) {
											if c < int32(2144) {
												v47 = libc.BoolInt32(c >= int32(2112) && c <= int32(2136))
											} else {
												v47 = libc.BoolInt32(c <= int32(2154))
											}
											v46 = v47
										} else {
											v46 = libc.BoolInt32(c <= int32(2183) || c >= int32(2185) && c <= int32(2190))
										}
									}
									v41 = libc.BoolInt32(v48 || v46 != 0)
								}
								v40 = v41
							} else {
								if v55 = c <= int32(2249); !v55 {
									if c < int32(2417) {
										if c < int32(2384) {
											if c < int32(2365) {
												v51 = libc.BoolInt32(c >= int32(2308) && c <= int32(2361))
											} else {
												v51 = libc.BoolInt32(c <= int32(2365))
											}
											v50 = v51
										} else {
											v50 = libc.BoolInt32(c <= int32(2384) || c >= int32(2392) && c <= int32(2401))
										}
										v49 = v50
									} else {
										if v54 = c <= int32(2432); !v54 {
											if c < int32(2451) {
												if c < int32(2447) {
													v53 = libc.BoolInt32(c >= int32(2437) && c <= int32(2444))
												} else {
													v53 = libc.BoolInt32(c <= int32(2448))
												}
												v52 = v53
											} else {
												v52 = libc.BoolInt32(c <= int32(2472) || c >= int32(2474) && c <= int32(2480))
											}
										}
										v49 = libc.BoolInt32(v54 || v52 != 0)
									}
								}
								v40 = libc.BoolInt32(v55 || v49 != 0)
							}
							v39 = v40
						} else {
							if v70 = c <= int32(2482); !v70 {
								if c < int32(2579) {
									if c < int32(2527) {
										if c < int32(2510) {
											if c < int32(2493) {
												v59 = libc.BoolInt32(c >= int32(2486) && c <= int32(2489))
											} else {
												v59 = libc.BoolInt32(c <= int32(2493))
											}
											v58 = v59
										} else {
											v58 = libc.BoolInt32(c <= int32(2510) || c >= int32(2524) && c <= int32(2525))
										}
										v57 = v58
									} else {
										if v62 = c <= int32(2529); !v62 {
											if c < int32(2565) {
												if c < int32(2556) {
													v61 = libc.BoolInt32(c >= int32(2544) && c <= int32(2545))
												} else {
													v61 = libc.BoolInt32(c <= int32(2556))
												}
												v60 = v61
											} else {
												v60 = libc.BoolInt32(c <= int32(2570) || c >= int32(2575) && c <= int32(2576))
											}
										}
										v57 = libc.BoolInt32(v62 || v60 != 0)
									}
									v56 = v57
								} else {
									if v69 = c <= int32(2600); !v69 {
										if c < int32(2649) {
											if c < int32(2613) {
												if c < int32(2610) {
													v65 = libc.BoolInt32(c >= int32(2602) && c <= int32(2608))
												} else {
													v65 = libc.BoolInt32(c <= int32(2611))
												}
												v64 = v65
											} else {
												v64 = libc.BoolInt32(c <= int32(2614) || c >= int32(2616) && c <= int32(2617))
											}
											v63 = v64
										} else {
											if v68 = c <= int32(2652); !v68 {
												if c < int32(2693) {
													if c < int32(2674) {
														v67 = libc.BoolInt32(c == int32(2654))
													} else {
														v67 = libc.BoolInt32(c <= int32(2676))
													}
													v66 = v67
												} else {
													v66 = libc.BoolInt32(c <= int32(2701) || c >= int32(2703) && c <= int32(2705))
												}
											}
											v63 = libc.BoolInt32(v68 || v66 != 0)
										}
									}
									v56 = libc.BoolInt32(v69 || v63 != 0)
								}
							}
							v39 = libc.BoolInt32(v70 || v56 != 0)
						}
					}
					v4 = libc.BoolInt32(v71 || v39 != 0)
				}
				v3 = v4
			} else {
				if v138 = c <= int32(2728); !v138 {
					if c < int32(3242) {
						if c < int32(2962) {
							if c < int32(2858) {
								if c < int32(2784) {
									if c < int32(2741) {
										if c < int32(2738) {
											v77 = libc.BoolInt32(c >= int32(2730) && c <= int32(2736))
										} else {
											v77 = libc.BoolInt32(c <= int32(2739))
										}
										v76 = v77
									} else {
										if v79 = c <= int32(2745); !v79 {
											if c < int32(2768) {
												v78 = libc.BoolInt32(c == int32(2749))
											} else {
												v78 = libc.BoolInt32(c <= int32(2768))
											}
										}
										v76 = libc.BoolInt32(v79 || v78 != 0)
									}
									v75 = v76
								} else {
									if v82 = c <= int32(2785); !v82 {
										if c < int32(2831) {
											if c < int32(2821) {
												v81 = libc.BoolInt32(c == int32(2809))
											} else {
												v81 = libc.BoolInt32(c <= int32(2828))
											}
											v80 = v81
										} else {
											v80 = libc.BoolInt32(c <= int32(2832) || c >= int32(2835) && c <= int32(2856))
										}
									}
									v75 = libc.BoolInt32(v82 || v80 != 0)
								}
								v74 = v75
							} else {
								if v89 = c <= int32(2864); !v89 {
									if c < int32(2911) {
										if c < int32(2877) {
											if c < int32(2869) {
												v85 = libc.BoolInt32(c >= int32(2866) && c <= int32(2867))
											} else {
												v85 = libc.BoolInt32(c <= int32(2873))
											}
											v84 = v85
										} else {
											v84 = libc.BoolInt32(c <= int32(2877) || c >= int32(2908) && c <= int32(2909))
										}
										v83 = v84
									} else {
										if v88 = c <= int32(2913); !v88 {
											if c < int32(2949) {
												if c < int32(2947) {
													v87 = libc.BoolInt32(c == int32(2929))
												} else {
													v87 = libc.BoolInt32(c <= int32(2947))
												}
												v86 = v87
											} else {
												v86 = libc.BoolInt32(c <= int32(2954) || c >= int32(2958) && c <= int32(2960))
											}
										}
										v83 = libc.BoolInt32(v88 || v86 != 0)
									}
								}
								v74 = libc.BoolInt32(v89 || v83 != 0)
							}
							v73 = v74
						} else {
							if v104 = c <= int32(2965); !v104 {
								if c < int32(3090) {
									if c < int32(2984) {
										if c < int32(2974) {
											if c < int32(2972) {
												v93 = libc.BoolInt32(c >= int32(2969) && c <= int32(2970))
											} else {
												v93 = libc.BoolInt32(c <= int32(2972))
											}
											v92 = v93
										} else {
											v92 = libc.BoolInt32(c <= int32(2975) || c >= int32(2979) && c <= int32(2980))
										}
										v91 = v92
									} else {
										if v96 = c <= int32(2986); !v96 {
											if c < int32(3077) {
												if c < int32(3024) {
													v95 = libc.BoolInt32(c >= int32(2990) && c <= int32(3001))
												} else {
													v95 = libc.BoolInt32(c <= int32(3024))
												}
												v94 = v95
											} else {
												v94 = libc.BoolInt32(c <= int32(3084) || c >= int32(3086) && c <= int32(3088))
											}
										}
										v91 = libc.BoolInt32(v96 || v94 != 0)
									}
									v90 = v91
								} else {
									if v103 = c <= int32(3112); !v103 {
										if c < int32(3168) {
											if c < int32(3160) {
												if c < int32(3133) {
													v99 = libc.BoolInt32(c >= int32(3114) && c <= int32(3129))
												} else {
													v99 = libc.BoolInt32(c <= int32(3133))
												}
												v98 = v99
											} else {
												v98 = libc.BoolInt32(c <= int32(3162) || c == int32(3165))
											}
											v97 = v98
										} else {
											if v102 = c <= int32(3169); !v102 {
												if c < int32(3214) {
													if c < int32(3205) {
														v101 = libc.BoolInt32(c == int32(3200))
													} else {
														v101 = libc.BoolInt32(c <= int32(3212))
													}
													v100 = v101
												} else {
													v100 = libc.BoolInt32(c <= int32(3216) || c >= int32(3218) && c <= int32(3240))
												}
											}
											v97 = libc.BoolInt32(v102 || v100 != 0)
										}
									}
									v90 = libc.BoolInt32(v103 || v97 != 0)
								}
							}
							v73 = libc.BoolInt32(v104 || v90 != 0)
						}
						v72 = v73
					} else {
						if v137 = c <= int32(3251); !v137 {
							if c < int32(3648) {
								if c < int32(3412) {
									if c < int32(3332) {
										if c < int32(3293) {
											if c < int32(3261) {
												v109 = libc.BoolInt32(c >= int32(3253) && c <= int32(3257))
											} else {
												v109 = libc.BoolInt32(c <= int32(3261))
											}
											v108 = v109
										} else {
											if v111 = c <= int32(3294); !v111 {
												if c < int32(3313) {
													v110 = libc.BoolInt32(c >= int32(3296) && c <= int32(3297))
												} else {
													v110 = libc.BoolInt32(c <= int32(3314))
												}
											}
											v108 = libc.BoolInt32(v111 || v110 != 0)
										}
										v107 = v108
									} else {
										if v114 = c <= int32(3340); !v114 {
											if c < int32(3389) {
												if c < int32(3346) {
													v113 = libc.BoolInt32(c >= int32(3342) && c <= int32(3344))
												} else {
													v113 = libc.BoolInt32(c <= int32(3386))
												}
												v112 = v113
											} else {
												v112 = libc.BoolInt32(c <= int32(3389) || c == int32(3406))
											}
										}
										v107 = libc.BoolInt32(v114 || v112 != 0)
									}
									v106 = v107
								} else {
									if v121 = c <= int32(3414); !v121 {
										if c < int32(3507) {
											if c < int32(3461) {
												if c < int32(3450) {
													v117 = libc.BoolInt32(c >= int32(3423) && c <= int32(3425))
												} else {
													v117 = libc.BoolInt32(c <= int32(3455))
												}
												v116 = v117
											} else {
												v116 = libc.BoolInt32(c <= int32(3478) || c >= int32(3482) && c <= int32(3505))
											}
											v115 = v116
										} else {
											if v120 = c <= int32(3515); !v120 {
												if c < int32(3585) {
													if c < int32(3520) {
														v119 = libc.BoolInt32(c == int32(3517))
													} else {
														v119 = libc.BoolInt32(c <= int32(3526))
													}
													v118 = v119
												} else {
													v118 = libc.BoolInt32(c <= int32(3632) || c == int32(3634))
												}
											}
											v115 = libc.BoolInt32(v120 || v118 != 0)
										}
									}
									v106 = libc.BoolInt32(v121 || v115 != 0)
								}
								v105 = v106
							} else {
								if v136 = c <= int32(3654); !v136 {
									if c < int32(3782) {
										if c < int32(3749) {
											if c < int32(3718) {
												if c < int32(3716) {
													v125 = libc.BoolInt32(c >= int32(3713) && c <= int32(3714))
												} else {
													v125 = libc.BoolInt32(c <= int32(3716))
												}
												v124 = v125
											} else {
												v124 = libc.BoolInt32(c <= int32(3722) || c >= int32(3724) && c <= int32(3747))
											}
											v123 = v124
										} else {
											if v128 = c <= int32(3749); !v128 {
												if c < int32(3773) {
													if c < int32(3762) {
														v127 = libc.BoolInt32(c >= int32(3751) && c <= int32(3760))
													} else {
														v127 = libc.BoolInt32(c <= int32(3762))
													}
													v126 = v127
												} else {
													v126 = libc.BoolInt32(c <= int32(3773) || c >= int32(3776) && c <= int32(3780))
												}
											}
											v123 = libc.BoolInt32(v128 || v126 != 0)
										}
										v122 = v123
									} else {
										if v135 = c <= int32(3782); !v135 {
											if c < int32(3976) {
												if c < int32(3904) {
													if c < int32(3840) {
														v131 = libc.BoolInt32(c >= int32(3804) && c <= int32(3807))
													} else {
														v131 = libc.BoolInt32(c <= int32(3840))
													}
													v130 = v131
												} else {
													v130 = libc.BoolInt32(c <= int32(3911) || c >= int32(3913) && c <= int32(3948))
												}
												v129 = v130
											} else {
												if v134 = c <= int32(3980); !v134 {
													if c < int32(4176) {
														if c < int32(4159) {
															v133 = libc.BoolInt32(c >= int32(4096) && c <= int32(4138))
														} else {
															v133 = libc.BoolInt32(c <= int32(4159))
														}
														v132 = v133
													} else {
														v132 = libc.BoolInt32(c <= int32(4181) || c >= int32(4186) && c <= int32(4189))
													}
												}
												v129 = libc.BoolInt32(v134 || v132 != 0)
											}
										}
										v122 = libc.BoolInt32(v135 || v129 != 0)
									}
								}
								v105 = libc.BoolInt32(v136 || v122 != 0)
							}
						}
						v72 = libc.BoolInt32(v137 || v105 != 0)
					}
				}
				v3 = libc.BoolInt32(v138 || v72 != 0)
			}
			v2 = v3
		} else {
			if v273 = c <= int32(4193); !v273 {
				if c < int32(8134) {
					if c < int32(6176) {
						if c < int32(4808) {
							if c < int32(4688) {
								if c < int32(4295) {
									if c < int32(4213) {
										if c < int32(4206) {
											v145 = libc.BoolInt32(c >= int32(4197) && c <= int32(4198))
										} else {
											v145 = libc.BoolInt32(c <= int32(4208))
										}
										v144 = v145
									} else {
										if v147 = c <= int32(4225); !v147 {
											if c < int32(4256) {
												v146 = libc.BoolInt32(c == int32(4238))
											} else {
												v146 = libc.BoolInt32(c <= int32(4293))
											}
										}
										v144 = libc.BoolInt32(v147 || v146 != 0)
									}
									v143 = v144
								} else {
									if v150 = c <= int32(4295); !v150 {
										if c < int32(4348) {
											if c < int32(4304) {
												v149 = libc.BoolInt32(c == int32(4301))
											} else {
												v149 = libc.BoolInt32(c <= int32(4346))
											}
											v148 = v149
										} else {
											v148 = libc.BoolInt32(c <= int32(4680) || c >= int32(4682) && c <= int32(4685))
										}
									}
									v143 = libc.BoolInt32(v150 || v148 != 0)
								}
								v142 = v143
							} else {
								if v157 = c <= int32(4694); !v157 {
									if c < int32(4752) {
										if c < int32(4704) {
											if c < int32(4698) {
												v153 = libc.BoolInt32(c == int32(4696))
											} else {
												v153 = libc.BoolInt32(c <= int32(4701))
											}
											v152 = v153
										} else {
											v152 = libc.BoolInt32(c <= int32(4744) || c >= int32(4746) && c <= int32(4749))
										}
										v151 = v152
									} else {
										if v156 = c <= int32(4784); !v156 {
											if c < int32(4800) {
												if c < int32(4792) {
													v155 = libc.BoolInt32(c >= int32(4786) && c <= int32(4789))
												} else {
													v155 = libc.BoolInt32(c <= int32(4798))
												}
												v154 = v155
											} else {
												v154 = libc.BoolInt32(c <= int32(4800) || c >= int32(4802) && c <= int32(4805))
											}
										}
										v151 = libc.BoolInt32(v156 || v154 != 0)
									}
								}
								v142 = libc.BoolInt32(v157 || v151 != 0)
							}
							v141 = v142
						} else {
							if v172 = c <= int32(4822); !v172 {
								if c < int32(5792) {
									if c < int32(5024) {
										if c < int32(4888) {
											if c < int32(4882) {
												v161 = libc.BoolInt32(c >= int32(4824) && c <= int32(4880))
											} else {
												v161 = libc.BoolInt32(c <= int32(4885))
											}
											v160 = v161
										} else {
											v160 = libc.BoolInt32(c <= int32(4954) || c >= int32(4992) && c <= int32(5007))
										}
										v159 = v160
									} else {
										if v164 = c <= int32(5109); !v164 {
											if c < int32(5743) {
												if c < int32(5121) {
													v163 = libc.BoolInt32(c >= int32(5112) && c <= int32(5117))
												} else {
													v163 = libc.BoolInt32(c <= int32(5740))
												}
												v162 = v163
											} else {
												v162 = libc.BoolInt32(c <= int32(5759) || c >= int32(5761) && c <= int32(5786))
											}
										}
										v159 = libc.BoolInt32(v164 || v162 != 0)
									}
									v158 = v159
								} else {
									if v171 = c <= int32(5866); !v171 {
										if c < int32(5984) {
											if c < int32(5919) {
												if c < int32(5888) {
													v167 = libc.BoolInt32(c >= int32(5870) && c <= int32(5880))
												} else {
													v167 = libc.BoolInt32(c <= int32(5905))
												}
												v166 = v167
											} else {
												v166 = libc.BoolInt32(c <= int32(5937) || c >= int32(5952) && c <= int32(5969))
											}
											v165 = v166
										} else {
											if v170 = c <= int32(5996); !v170 {
												if c < int32(6103) {
													if c < int32(6016) {
														v169 = libc.BoolInt32(c >= int32(5998) && c <= int32(6000))
													} else {
														v169 = libc.BoolInt32(c <= int32(6067))
													}
													v168 = v169
												} else {
													v168 = libc.BoolInt32(c <= int32(6103) || c == int32(6108))
												}
											}
											v165 = libc.BoolInt32(v170 || v168 != 0)
										}
									}
									v158 = libc.BoolInt32(v171 || v165 != 0)
								}
							}
							v141 = libc.BoolInt32(v172 || v158 != 0)
						}
						v140 = v141
					} else {
						if v205 = c <= int32(6264); !v205 {
							if c < int32(7312) {
								if c < int32(6823) {
									if c < int32(6512) {
										if c < int32(6320) {
											if c < int32(6314) {
												v177 = libc.BoolInt32(c >= int32(6272) && c <= int32(6312))
											} else {
												v177 = libc.BoolInt32(c <= int32(6314))
											}
											v176 = v177
										} else {
											if v179 = c <= int32(6389); !v179 {
												if c < int32(6480) {
													v178 = libc.BoolInt32(c >= int32(6400) && c <= int32(6430))
												} else {
													v178 = libc.BoolInt32(c <= int32(6509))
												}
											}
											v176 = libc.BoolInt32(v179 || v178 != 0)
										}
										v175 = v176
									} else {
										if v182 = c <= int32(6516); !v182 {
											if c < int32(6656) {
												if c < int32(6576) {
													v181 = libc.BoolInt32(c >= int32(6528) && c <= int32(6571))
												} else {
													v181 = libc.BoolInt32(c <= int32(6601))
												}
												v180 = v181
											} else {
												v180 = libc.BoolInt32(c <= int32(6678) || c >= int32(6688) && c <= int32(6740))
											}
										}
										v175 = libc.BoolInt32(v182 || v180 != 0)
									}
									v174 = v175
								} else {
									if v189 = c <= int32(6823); !v189 {
										if c < int32(7098) {
											if c < int32(7043) {
												if c < int32(6981) {
													v185 = libc.BoolInt32(c >= int32(6917) && c <= int32(6963))
												} else {
													v185 = libc.BoolInt32(c <= int32(6988))
												}
												v184 = v185
											} else {
												v184 = libc.BoolInt32(c <= int32(7072) || c >= int32(7086) && c <= int32(7087))
											}
											v183 = v184
										} else {
											if v188 = c <= int32(7141); !v188 {
												if c < int32(7258) {
													if c < int32(7245) {
														v187 = libc.BoolInt32(c >= int32(7168) && c <= int32(7203))
													} else {
														v187 = libc.BoolInt32(c <= int32(7247))
													}
													v186 = v187
												} else {
													v186 = libc.BoolInt32(c <= int32(7293) || c >= int32(7296) && c <= int32(7304))
												}
											}
											v183 = libc.BoolInt32(v188 || v186 != 0)
										}
									}
									v174 = libc.BoolInt32(v189 || v183 != 0)
								}
								v173 = v174
							} else {
								if v204 = c <= int32(7354); !v204 {
									if c < int32(8008) {
										if c < int32(7418) {
											if c < int32(7406) {
												if c < int32(7401) {
													v193 = libc.BoolInt32(c >= int32(7357) && c <= int32(7359))
												} else {
													v193 = libc.BoolInt32(c <= int32(7404))
												}
												v192 = v193
											} else {
												v192 = libc.BoolInt32(c <= int32(7411) || c >= int32(7413) && c <= int32(7414))
											}
											v191 = v192
										} else {
											if v196 = c <= int32(7418); !v196 {
												if c < int32(7960) {
													if c < int32(7680) {
														v195 = libc.BoolInt32(c >= int32(7424) && c <= int32(7615))
													} else {
														v195 = libc.BoolInt32(c <= int32(7957))
													}
													v194 = v195
												} else {
													v194 = libc.BoolInt32(c <= int32(7965) || c >= int32(7968) && c <= int32(8005))
												}
											}
											v191 = libc.BoolInt32(v196 || v194 != 0)
										}
										v190 = v191
									} else {
										if v203 = c <= int32(8013); !v203 {
											if c < int32(8031) {
												if c < int32(8027) {
													if c < int32(8025) {
														v199 = libc.BoolInt32(c >= int32(8016) && c <= int32(8023))
													} else {
														v199 = libc.BoolInt32(c <= int32(8025))
													}
													v198 = v199
												} else {
													v198 = libc.BoolInt32(c <= int32(8027) || c == int32(8029))
												}
												v197 = v198
											} else {
												if v202 = c <= int32(8061); !v202 {
													if c < int32(8126) {
														if c < int32(8118) {
															v201 = libc.BoolInt32(c >= int32(8064) && c <= int32(8116))
														} else {
															v201 = libc.BoolInt32(c <= int32(8124))
														}
														v200 = v201
													} else {
														v200 = libc.BoolInt32(c <= int32(8126) || c >= int32(8130) && c <= int32(8132))
													}
												}
												v197 = libc.BoolInt32(v202 || v200 != 0)
											}
										}
										v190 = libc.BoolInt32(v203 || v197 != 0)
									}
								}
								v173 = libc.BoolInt32(v204 || v190 != 0)
							}
						}
						v140 = libc.BoolInt32(v205 || v173 != 0)
					}
					v139 = v140
				} else {
					if v272 = c <= int32(8140); !v272 {
						if c < int32(12337) {
							if c < int32(8544) {
								if c < int32(8458) {
									if c < int32(8305) {
										if c < int32(8160) {
											if c < int32(8150) {
												v211 = libc.BoolInt32(c >= int32(8144) && c <= int32(8147))
											} else {
												v211 = libc.BoolInt32(c <= int32(8155))
											}
											v210 = v211
										} else {
											if v213 = c <= int32(8172); !v213 {
												if c < int32(8182) {
													v212 = libc.BoolInt32(c >= int32(8178) && c <= int32(8180))
												} else {
													v212 = libc.BoolInt32(c <= int32(8188))
												}
											}
											v210 = libc.BoolInt32(v213 || v212 != 0)
										}
										v209 = v210
									} else {
										if v216 = c <= int32(8305); !v216 {
											if c < int32(8450) {
												if c < int32(8336) {
													v215 = libc.BoolInt32(c == int32(8319))
												} else {
													v215 = libc.BoolInt32(c <= int32(8348))
												}
												v214 = v215
											} else {
												v214 = libc.BoolInt32(c <= int32(8450) || c == int32(8455))
											}
										}
										v209 = libc.BoolInt32(v216 || v214 != 0)
									}
									v208 = v209
								} else {
									if v223 = c <= int32(8467); !v223 {
										if c < int32(8488) {
											if c < int32(8484) {
												if c < int32(8472) {
													v219 = libc.BoolInt32(c == int32(8469))
												} else {
													v219 = libc.BoolInt32(c <= int32(8477))
												}
												v218 = v219
											} else {
												v218 = libc.BoolInt32(c <= int32(8484) || c == int32(8486))
											}
											v217 = v218
										} else {
											if v222 = c <= int32(8488); !v222 {
												if c < int32(8517) {
													if c < int32(8508) {
														v221 = libc.BoolInt32(c >= int32(8490) && c <= int32(8505))
													} else {
														v221 = libc.BoolInt32(c <= int32(8511))
													}
													v220 = v221
												} else {
													v220 = libc.BoolInt32(c <= int32(8521) || c == int32(8526))
												}
											}
											v217 = libc.BoolInt32(v222 || v220 != 0)
										}
									}
									v208 = libc.BoolInt32(v223 || v217 != 0)
								}
								v207 = v208
							} else {
								if v238 = c <= int32(8584); !v238 {
									if c < int32(11680) {
										if c < int32(11559) {
											if c < int32(11506) {
												if c < int32(11499) {
													v227 = libc.BoolInt32(c >= int32(11264) && c <= int32(11492))
												} else {
													v227 = libc.BoolInt32(c <= int32(11502))
												}
												v226 = v227
											} else {
												v226 = libc.BoolInt32(c <= int32(11507) || c >= int32(11520) && c <= int32(11557))
											}
											v225 = v226
										} else {
											if v230 = c <= int32(11559); !v230 {
												if c < int32(11631) {
													if c < int32(11568) {
														v229 = libc.BoolInt32(c == int32(11565))
													} else {
														v229 = libc.BoolInt32(c <= int32(11623))
													}
													v228 = v229
												} else {
													v228 = libc.BoolInt32(c <= int32(11631) || c >= int32(11648) && c <= int32(11670))
												}
											}
											v225 = libc.BoolInt32(v230 || v228 != 0)
										}
										v224 = v225
									} else {
										if v237 = c <= int32(11686); !v237 {
											if c < int32(11720) {
												if c < int32(11704) {
													if c < int32(11696) {
														v233 = libc.BoolInt32(c >= int32(11688) && c <= int32(11694))
													} else {
														v233 = libc.BoolInt32(c <= int32(11702))
													}
													v232 = v233
												} else {
													v232 = libc.BoolInt32(c <= int32(11710) || c >= int32(11712) && c <= int32(11718))
												}
												v231 = v232
											} else {
												if v236 = c <= int32(11726); !v236 {
													if c < int32(12293) {
														if c < int32(11736) {
															v235 = libc.BoolInt32(c >= int32(11728) && c <= int32(11734))
														} else {
															v235 = libc.BoolInt32(c <= int32(11742))
														}
														v234 = v235
													} else {
														v234 = libc.BoolInt32(c <= int32(12295) || c >= int32(12321) && c <= int32(12329))
													}
												}
												v231 = libc.BoolInt32(v236 || v234 != 0)
											}
										}
										v224 = libc.BoolInt32(v237 || v231 != 0)
									}
								}
								v207 = libc.BoolInt32(v238 || v224 != 0)
							}
							v206 = v207
						} else {
							if v271 = c <= int32(12341); !v271 {
								if c < int32(42891) {
									if c < int32(19968) {
										if c < int32(12549) {
											if c < int32(12445) {
												if c < int32(12353) {
													v243 = libc.BoolInt32(c >= int32(12344) && c <= int32(12348))
												} else {
													v243 = libc.BoolInt32(c <= int32(12438))
												}
												v242 = v243
											} else {
												if v245 = c <= int32(12447); !v245 {
													if c < int32(12540) {
														v244 = libc.BoolInt32(c >= int32(12449) && c <= int32(12538))
													} else {
														v244 = libc.BoolInt32(c <= int32(12543))
													}
												}
												v242 = libc.BoolInt32(v245 || v244 != 0)
											}
											v241 = v242
										} else {
											if v248 = c <= int32(12591); !v248 {
												if c < int32(12784) {
													if c < int32(12704) {
														v247 = libc.BoolInt32(c >= int32(12593) && c <= int32(12686))
													} else {
														v247 = libc.BoolInt32(c <= int32(12735))
													}
													v246 = v247
												} else {
													v246 = libc.BoolInt32(c <= int32(12799) || c >= int32(13312) && c <= int32(19903))
												}
											}
											v241 = libc.BoolInt32(v248 || v246 != 0)
										}
										v240 = v241
									} else {
										if v255 = c <= int32(42124); !v255 {
											if c < int32(42560) {
												if c < int32(42512) {
													if c < int32(42240) {
														v251 = libc.BoolInt32(c >= int32(42192) && c <= int32(42237))
													} else {
														v251 = libc.BoolInt32(c <= int32(42508))
													}
													v250 = v251
												} else {
													v250 = libc.BoolInt32(c <= int32(42527) || c >= int32(42538) && c <= int32(42539))
												}
												v249 = v250
											} else {
												if v254 = c <= int32(42606); !v254 {
													if c < int32(42775) {
														if c < int32(42656) {
															v253 = libc.BoolInt32(c >= int32(42623) && c <= int32(42653))
														} else {
															v253 = libc.BoolInt32(c <= int32(42735))
														}
														v252 = v253
													} else {
														v252 = libc.BoolInt32(c <= int32(42783) || c >= int32(42786) && c <= int32(42888))
													}
												}
												v249 = libc.BoolInt32(v254 || v252 != 0)
											}
										}
										v240 = libc.BoolInt32(v255 || v249 != 0)
									}
									v239 = v240
								} else {
									if v270 = c <= int32(42954); !v270 {
										if c < int32(43250) {
											if c < int32(43011) {
												if c < int32(42965) {
													if c < int32(42963) {
														v259 = libc.BoolInt32(c >= int32(42960) && c <= int32(42961))
													} else {
														v259 = libc.BoolInt32(c <= int32(42963))
													}
													v258 = v259
												} else {
													v258 = libc.BoolInt32(c <= int32(42969) || c >= int32(42994) && c <= int32(43009))
												}
												v257 = v258
											} else {
												if v262 = c <= int32(43013); !v262 {
													if c < int32(43072) {
														if c < int32(43020) {
															v261 = libc.BoolInt32(c >= int32(43015) && c <= int32(43018))
														} else {
															v261 = libc.BoolInt32(c <= int32(43042))
														}
														v260 = v261
													} else {
														v260 = libc.BoolInt32(c <= int32(43123) || c >= int32(43138) && c <= int32(43187))
													}
												}
												v257 = libc.BoolInt32(v262 || v260 != 0)
											}
											v256 = v257
										} else {
											if v269 = c <= int32(43255); !v269 {
												if c < int32(43360) {
													if c < int32(43274) {
														if c < int32(43261) {
															v265 = libc.BoolInt32(c == int32(43259))
														} else {
															v265 = libc.BoolInt32(c <= int32(43262))
														}
														v264 = v265
													} else {
														v264 = libc.BoolInt32(c <= int32(43301) || c >= int32(43312) && c <= int32(43334))
													}
													v263 = v264
												} else {
													if v268 = c <= int32(43388); !v268 {
														if c < int32(43488) {
															if c < int32(43471) {
																v267 = libc.BoolInt32(c >= int32(43396) && c <= int32(43442))
															} else {
																v267 = libc.BoolInt32(c <= int32(43471))
															}
															v266 = v267
														} else {
															v266 = libc.BoolInt32(c <= int32(43492) || c >= int32(43494) && c <= int32(43503))
														}
													}
													v263 = libc.BoolInt32(v268 || v266 != 0)
												}
											}
											v256 = libc.BoolInt32(v269 || v263 != 0)
										}
									}
									v239 = libc.BoolInt32(v270 || v256 != 0)
								}
							}
							v206 = libc.BoolInt32(v271 || v239 != 0)
						}
					}
					v139 = libc.BoolInt32(v272 || v206 != 0)
				}
			}
			v2 = libc.BoolInt32(v273 || v139 != 0)
		}
		v1 = v2
	} else {
		if v544 = c <= int32(43518); !v544 {
			if c < int32(70727) {
				if c < int32(66956) {
					if c < int32(64914) {
						if c < int32(43868) {
							if c < int32(43714) {
								if c < int32(43646) {
									if c < int32(43588) {
										if c < int32(43584) {
											v281 = libc.BoolInt32(c >= int32(43520) && c <= int32(43560))
										} else {
											v281 = libc.BoolInt32(c <= int32(43586))
										}
										v280 = v281
									} else {
										if v283 = c <= int32(43595); !v283 {
											if c < int32(43642) {
												v282 = libc.BoolInt32(c >= int32(43616) && c <= int32(43638))
											} else {
												v282 = libc.BoolInt32(c <= int32(43642))
											}
										}
										v280 = libc.BoolInt32(v283 || v282 != 0)
									}
									v279 = v280
								} else {
									if v286 = c <= int32(43695); !v286 {
										if c < int32(43705) {
											if c < int32(43701) {
												v285 = libc.BoolInt32(c == int32(43697))
											} else {
												v285 = libc.BoolInt32(c <= int32(43702))
											}
											v284 = v285
										} else {
											v284 = libc.BoolInt32(c <= int32(43709) || c == int32(43712))
										}
									}
									v279 = libc.BoolInt32(v286 || v284 != 0)
								}
								v278 = v279
							} else {
								if v293 = c <= int32(43714); !v293 {
									if c < int32(43785) {
										if c < int32(43762) {
											if c < int32(43744) {
												v289 = libc.BoolInt32(c >= int32(43739) && c <= int32(43741))
											} else {
												v289 = libc.BoolInt32(c <= int32(43754))
											}
											v288 = v289
										} else {
											v288 = libc.BoolInt32(c <= int32(43764) || c >= int32(43777) && c <= int32(43782))
										}
										v287 = v288
									} else {
										if v292 = c <= int32(43790); !v292 {
											if c < int32(43816) {
												if c < int32(43808) {
													v291 = libc.BoolInt32(c >= int32(43793) && c <= int32(43798))
												} else {
													v291 = libc.BoolInt32(c <= int32(43814))
												}
												v290 = v291
											} else {
												v290 = libc.BoolInt32(c <= int32(43822) || c >= int32(43824) && c <= int32(43866))
											}
										}
										v287 = libc.BoolInt32(v292 || v290 != 0)
									}
								}
								v278 = libc.BoolInt32(v293 || v287 != 0)
							}
							v277 = v278
						} else {
							if v308 = c <= int32(43881); !v308 {
								if c < int32(64287) {
									if c < int32(63744) {
										if c < int32(55216) {
											if c < int32(44032) {
												v297 = libc.BoolInt32(c >= int32(43888) && c <= int32(44002))
											} else {
												v297 = libc.BoolInt32(c <= int32(55203))
											}
											v296 = v297
										} else {
											v296 = libc.BoolInt32(c <= int32(55238) || c >= int32(55243) && c <= int32(55291))
										}
										v295 = v296
									} else {
										if v300 = c <= int32(64109); !v300 {
											if c < int32(64275) {
												if c < int32(64256) {
													v299 = libc.BoolInt32(c >= int32(64112) && c <= int32(64217))
												} else {
													v299 = libc.BoolInt32(c <= int32(64262))
												}
												v298 = v299
											} else {
												v298 = libc.BoolInt32(c <= int32(64279) || c == int32(64285))
											}
										}
										v295 = libc.BoolInt32(v300 || v298 != 0)
									}
									v294 = v295
								} else {
									if v307 = c <= int32(64296); !v307 {
										if c < int32(64323) {
											if c < int32(64318) {
												if c < int32(64312) {
													v303 = libc.BoolInt32(c >= int32(64298) && c <= int32(64310))
												} else {
													v303 = libc.BoolInt32(c <= int32(64316))
												}
												v302 = v303
											} else {
												v302 = libc.BoolInt32(c <= int32(64318) || c >= int32(64320) && c <= int32(64321))
											}
											v301 = v302
										} else {
											if v306 = c <= int32(64324); !v306 {
												if c < int32(64612) {
													if c < int32(64467) {
														v305 = libc.BoolInt32(c >= int32(64326) && c <= int32(64433))
													} else {
														v305 = libc.BoolInt32(c <= int32(64605))
													}
													v304 = v305
												} else {
													v304 = libc.BoolInt32(c <= int32(64829) || c >= int32(64848) && c <= int32(64911))
												}
											}
											v301 = libc.BoolInt32(v306 || v304 != 0)
										}
									}
									v294 = libc.BoolInt32(v307 || v301 != 0)
								}
							}
							v277 = libc.BoolInt32(v308 || v294 != 0)
						}
						v276 = v277
					} else {
						if v341 = c <= int32(64967); !v341 {
							if c < int32(65599) {
								if c < int32(65382) {
									if c < int32(65147) {
										if c < int32(65139) {
											if c < int32(65137) {
												v313 = libc.BoolInt32(c >= int32(65008) && c <= int32(65017))
											} else {
												v313 = libc.BoolInt32(c <= int32(65137))
											}
											v312 = v313
										} else {
											if v315 = c <= int32(65139); !v315 {
												if c < int32(65145) {
													v314 = libc.BoolInt32(c == int32(65143))
												} else {
													v314 = libc.BoolInt32(c <= int32(65145))
												}
											}
											v312 = libc.BoolInt32(v315 || v314 != 0)
										}
										v311 = v312
									} else {
										if v318 = c <= int32(65147); !v318 {
											if c < int32(65313) {
												if c < int32(65151) {
													v317 = libc.BoolInt32(c == int32(65149))
												} else {
													v317 = libc.BoolInt32(c <= int32(65276))
												}
												v316 = v317
											} else {
												v316 = libc.BoolInt32(c <= int32(65338) || c >= int32(65345) && c <= int32(65370))
											}
										}
										v311 = libc.BoolInt32(v318 || v316 != 0)
									}
									v310 = v311
								} else {
									if v325 = c <= int32(65437); !v325 {
										if c < int32(65498) {
											if c < int32(65482) {
												if c < int32(65474) {
													v321 = libc.BoolInt32(c >= int32(65440) && c <= int32(65470))
												} else {
													v321 = libc.BoolInt32(c <= int32(65479))
												}
												v320 = v321
											} else {
												v320 = libc.BoolInt32(c <= int32(65487) || c >= int32(65490) && c <= int32(65495))
											}
											v319 = v320
										} else {
											if v324 = c <= int32(65500); !v324 {
												if c < int32(65576) {
													if c < int32(65549) {
														v323 = libc.BoolInt32(c >= int32(65536) && c <= int32(65547))
													} else {
														v323 = libc.BoolInt32(c <= int32(65574))
													}
													v322 = v323
												} else {
													v322 = libc.BoolInt32(c <= int32(65594) || c >= int32(65596) && c <= int32(65597))
												}
											}
											v319 = libc.BoolInt32(v324 || v322 != 0)
										}
									}
									v310 = libc.BoolInt32(v325 || v319 != 0)
								}
								v309 = v310
							} else {
								if v340 = c <= int32(65613); !v340 {
									if c < int32(66464) {
										if c < int32(66208) {
											if c < int32(65856) {
												if c < int32(65664) {
													v329 = libc.BoolInt32(c >= int32(65616) && c <= int32(65629))
												} else {
													v329 = libc.BoolInt32(c <= int32(65786))
												}
												v328 = v329
											} else {
												v328 = libc.BoolInt32(c <= int32(65908) || c >= int32(66176) && c <= int32(66204))
											}
											v327 = v328
										} else {
											if v332 = c <= int32(66256); !v332 {
												if c < int32(66384) {
													if c < int32(66349) {
														v331 = libc.BoolInt32(c >= int32(66304) && c <= int32(66335))
													} else {
														v331 = libc.BoolInt32(c <= int32(66378))
													}
													v330 = v331
												} else {
													v330 = libc.BoolInt32(c <= int32(66421) || c >= int32(66432) && c <= int32(66461))
												}
											}
											v327 = libc.BoolInt32(v332 || v330 != 0)
										}
										v326 = v327
									} else {
										if v339 = c <= int32(66499); !v339 {
											if c < int32(66776) {
												if c < int32(66560) {
													if c < int32(66513) {
														v335 = libc.BoolInt32(c >= int32(66504) && c <= int32(66511))
													} else {
														v335 = libc.BoolInt32(c <= int32(66517))
													}
													v334 = v335
												} else {
													v334 = libc.BoolInt32(c <= int32(66717) || c >= int32(66736) && c <= int32(66771))
												}
												v333 = v334
											} else {
												if v338 = c <= int32(66811); !v338 {
													if c < int32(66928) {
														if c < int32(66864) {
															v337 = libc.BoolInt32(c >= int32(66816) && c <= int32(66855))
														} else {
															v337 = libc.BoolInt32(c <= int32(66915))
														}
														v336 = v337
													} else {
														v336 = libc.BoolInt32(c <= int32(66938) || c >= int32(66940) && c <= int32(66954))
													}
												}
												v333 = libc.BoolInt32(v338 || v336 != 0)
											}
										}
										v326 = libc.BoolInt32(v339 || v333 != 0)
									}
								}
								v309 = libc.BoolInt32(v340 || v326 != 0)
							}
						}
						v276 = libc.BoolInt32(v341 || v309 != 0)
					}
					v275 = v276
				} else {
					if v408 = c <= int32(66962); !v408 {
						if c < int32(68864) {
							if c < int32(67828) {
								if c < int32(67506) {
									if c < int32(67072) {
										if c < int32(66979) {
											if c < int32(66967) {
												v347 = libc.BoolInt32(c >= int32(66964) && c <= int32(66965))
											} else {
												v347 = libc.BoolInt32(c <= int32(66977))
											}
											v346 = v347
										} else {
											if v349 = c <= int32(66993); !v349 {
												if c < int32(67003) {
													v348 = libc.BoolInt32(c >= int32(66995) && c <= int32(67001))
												} else {
													v348 = libc.BoolInt32(c <= int32(67004))
												}
											}
											v346 = libc.BoolInt32(v349 || v348 != 0)
										}
										v345 = v346
									} else {
										if v352 = c <= int32(67382); !v352 {
											if c < int32(67456) {
												if c < int32(67424) {
													v351 = libc.BoolInt32(c >= int32(67392) && c <= int32(67413))
												} else {
													v351 = libc.BoolInt32(c <= int32(67431))
												}
												v350 = v351
											} else {
												v350 = libc.BoolInt32(c <= int32(67461) || c >= int32(67463) && c <= int32(67504))
											}
										}
										v345 = libc.BoolInt32(v352 || v350 != 0)
									}
									v344 = v345
								} else {
									if v359 = c <= int32(67514); !v359 {
										if c < int32(67644) {
											if c < int32(67594) {
												if c < int32(67592) {
													v355 = libc.BoolInt32(c >= int32(67584) && c <= int32(67589))
												} else {
													v355 = libc.BoolInt32(c <= int32(67592))
												}
												v354 = v355
											} else {
												v354 = libc.BoolInt32(c <= int32(67637) || c >= int32(67639) && c <= int32(67640))
											}
											v353 = v354
										} else {
											if v358 = c <= int32(67644); !v358 {
												if c < int32(67712) {
													if c < int32(67680) {
														v357 = libc.BoolInt32(c >= int32(67647) && c <= int32(67669))
													} else {
														v357 = libc.BoolInt32(c <= int32(67702))
													}
													v356 = v357
												} else {
													v356 = libc.BoolInt32(c <= int32(67742) || c >= int32(67808) && c <= int32(67826))
												}
											}
											v353 = libc.BoolInt32(v358 || v356 != 0)
										}
									}
									v344 = libc.BoolInt32(v359 || v353 != 0)
								}
								v343 = v344
							} else {
								if v374 = c <= int32(67829); !v374 {
									if c < int32(68224) {
										if c < int32(68096) {
											if c < int32(67968) {
												if c < int32(67872) {
													v363 = libc.BoolInt32(c >= int32(67840) && c <= int32(67861))
												} else {
													v363 = libc.BoolInt32(c <= int32(67897))
												}
												v362 = v363
											} else {
												v362 = libc.BoolInt32(c <= int32(68023) || c >= int32(68030) && c <= int32(68031))
											}
											v361 = v362
										} else {
											if v366 = c <= int32(68096); !v366 {
												if c < int32(68121) {
													if c < int32(68117) {
														v365 = libc.BoolInt32(c >= int32(68112) && c <= int32(68115))
													} else {
														v365 = libc.BoolInt32(c <= int32(68119))
													}
													v364 = v365
												} else {
													v364 = libc.BoolInt32(c <= int32(68149) || c >= int32(68192) && c <= int32(68220))
												}
											}
											v361 = libc.BoolInt32(v366 || v364 != 0)
										}
										v360 = v361
									} else {
										if v373 = c <= int32(68252); !v373 {
											if c < int32(68448) {
												if c < int32(68352) {
													if c < int32(68297) {
														v369 = libc.BoolInt32(c >= int32(68288) && c <= int32(68295))
													} else {
														v369 = libc.BoolInt32(c <= int32(68324))
													}
													v368 = v369
												} else {
													v368 = libc.BoolInt32(c <= int32(68405) || c >= int32(68416) && c <= int32(68437))
												}
												v367 = v368
											} else {
												if v372 = c <= int32(68466); !v372 {
													if c < int32(68736) {
														if c < int32(68608) {
															v371 = libc.BoolInt32(c >= int32(68480) && c <= int32(68497))
														} else {
															v371 = libc.BoolInt32(c <= int32(68680))
														}
														v370 = v371
													} else {
														v370 = libc.BoolInt32(c <= int32(68786) || c >= int32(68800) && c <= int32(68850))
													}
												}
												v367 = libc.BoolInt32(v372 || v370 != 0)
											}
										}
										v360 = libc.BoolInt32(v373 || v367 != 0)
									}
								}
								v343 = libc.BoolInt32(v374 || v360 != 0)
							}
							v342 = v343
						} else {
							if v407 = c <= int32(68899); !v407 {
								if c < int32(70106) {
									if c < int32(69749) {
										if c < int32(69488) {
											if c < int32(69376) {
												if c < int32(69296) {
													v379 = libc.BoolInt32(c >= int32(69248) && c <= int32(69289))
												} else {
													v379 = libc.BoolInt32(c <= int32(69297))
												}
												v378 = v379
											} else {
												if v381 = c <= int32(69404); !v381 {
													if c < int32(69424) {
														v380 = libc.BoolInt32(c == int32(69415))
													} else {
														v380 = libc.BoolInt32(c <= int32(69445))
													}
												}
												v378 = libc.BoolInt32(v381 || v380 != 0)
											}
											v377 = v378
										} else {
											if v384 = c <= int32(69505); !v384 {
												if c < int32(69635) {
													if c < int32(69600) {
														v383 = libc.BoolInt32(c >= int32(69552) && c <= int32(69572))
													} else {
														v383 = libc.BoolInt32(c <= int32(69622))
													}
													v382 = v383
												} else {
													v382 = libc.BoolInt32(c <= int32(69687) || c >= int32(69745) && c <= int32(69746))
												}
											}
											v377 = libc.BoolInt32(v384 || v382 != 0)
										}
										v376 = v377
									} else {
										if v391 = c <= int32(69749); !v391 {
											if c < int32(69959) {
												if c < int32(69891) {
													if c < int32(69840) {
														v387 = libc.BoolInt32(c >= int32(69763) && c <= int32(69807))
													} else {
														v387 = libc.BoolInt32(c <= int32(69864))
													}
													v386 = v387
												} else {
													v386 = libc.BoolInt32(c <= int32(69926) || c == int32(69956))
												}
												v385 = v386
											} else {
												if v390 = c <= int32(69959); !v390 {
													if c < int32(70019) {
														if c < int32(70006) {
															v389 = libc.BoolInt32(c >= int32(69968) && c <= int32(70002))
														} else {
															v389 = libc.BoolInt32(c <= int32(70006))
														}
														v388 = v389
													} else {
														v388 = libc.BoolInt32(c <= int32(70066) || c >= int32(70081) && c <= int32(70084))
													}
												}
												v385 = libc.BoolInt32(v390 || v388 != 0)
											}
										}
										v376 = libc.BoolInt32(v391 || v385 != 0)
									}
									v375 = v376
								} else {
									if v406 = c <= int32(70106); !v406 {
										if c < int32(70405) {
											if c < int32(70280) {
												if c < int32(70163) {
													if c < int32(70144) {
														v395 = libc.BoolInt32(c == int32(70108))
													} else {
														v395 = libc.BoolInt32(c <= int32(70161))
													}
													v394 = v395
												} else {
													v394 = libc.BoolInt32(c <= int32(70187) || c >= int32(70272) && c <= int32(70278))
												}
												v393 = v394
											} else {
												if v398 = c <= int32(70280); !v398 {
													if c < int32(70303) {
														if c < int32(70287) {
															v397 = libc.BoolInt32(c >= int32(70282) && c <= int32(70285))
														} else {
															v397 = libc.BoolInt32(c <= int32(70301))
														}
														v396 = v397
													} else {
														v396 = libc.BoolInt32(c <= int32(70312) || c >= int32(70320) && c <= int32(70366))
													}
												}
												v393 = libc.BoolInt32(v398 || v396 != 0)
											}
											v392 = v393
										} else {
											if v405 = c <= int32(70412); !v405 {
												if c < int32(70453) {
													if c < int32(70442) {
														if c < int32(70419) {
															v401 = libc.BoolInt32(c >= int32(70415) && c <= int32(70416))
														} else {
															v401 = libc.BoolInt32(c <= int32(70440))
														}
														v400 = v401
													} else {
														v400 = libc.BoolInt32(c <= int32(70448) || c >= int32(70450) && c <= int32(70451))
													}
													v399 = v400
												} else {
													if v404 = c <= int32(70457); !v404 {
														if c < int32(70493) {
															if c < int32(70480) {
																v403 = libc.BoolInt32(c == int32(70461))
															} else {
																v403 = libc.BoolInt32(c <= int32(70480))
															}
															v402 = v403
														} else {
															v402 = libc.BoolInt32(c <= int32(70497) || c >= int32(70656) && c <= int32(70708))
														}
													}
													v399 = libc.BoolInt32(v404 || v402 != 0)
												}
											}
											v392 = libc.BoolInt32(v405 || v399 != 0)
										}
									}
									v375 = libc.BoolInt32(v406 || v392 != 0)
								}
							}
							v342 = libc.BoolInt32(v407 || v375 != 0)
						}
					}
					v275 = libc.BoolInt32(v408 || v342 != 0)
				}
				v274 = v275
			} else {
				if v543 = c <= int32(70730); !v543 {
					if c < int32(119894) {
						if c < int32(73056) {
							if c < int32(72001) {
								if c < int32(71424) {
									if c < int32(71128) {
										if c < int32(70852) {
											if c < int32(70784) {
												v415 = libc.BoolInt32(c >= int32(70751) && c <= int32(70753))
											} else {
												v415 = libc.BoolInt32(c <= int32(70831))
											}
											v414 = v415
										} else {
											if v417 = c <= int32(70853); !v417 {
												if c < int32(71040) {
													v416 = libc.BoolInt32(c == int32(70855))
												} else {
													v416 = libc.BoolInt32(c <= int32(71086))
												}
											}
											v414 = libc.BoolInt32(v417 || v416 != 0)
										}
										v413 = v414
									} else {
										if v420 = c <= int32(71131); !v420 {
											if c < int32(71296) {
												if c < int32(71236) {
													v419 = libc.BoolInt32(c >= int32(71168) && c <= int32(71215))
												} else {
													v419 = libc.BoolInt32(c <= int32(71236))
												}
												v418 = v419
											} else {
												v418 = libc.BoolInt32(c <= int32(71338) || c == int32(71352))
											}
										}
										v413 = libc.BoolInt32(v420 || v418 != 0)
									}
									v412 = v413
								} else {
									if v427 = c <= int32(71450); !v427 {
										if c < int32(71945) {
											if c < int32(71840) {
												if c < int32(71680) {
													v423 = libc.BoolInt32(c >= int32(71488) && c <= int32(71494))
												} else {
													v423 = libc.BoolInt32(c <= int32(71723))
												}
												v422 = v423
											} else {
												v422 = libc.BoolInt32(c <= int32(71903) || c >= int32(71935) && c <= int32(71942))
											}
											v421 = v422
										} else {
											if v426 = c <= int32(71945); !v426 {
												if c < int32(71960) {
													if c < int32(71957) {
														v425 = libc.BoolInt32(c >= int32(71948) && c <= int32(71955))
													} else {
														v425 = libc.BoolInt32(c <= int32(71958))
													}
													v424 = v425
												} else {
													v424 = libc.BoolInt32(c <= int32(71983) || c == int32(71999))
												}
											}
											v421 = libc.BoolInt32(v426 || v424 != 0)
										}
									}
									v412 = libc.BoolInt32(v427 || v421 != 0)
								}
								v411 = v412
							} else {
								if v442 = c <= int32(72001); !v442 {
									if c < int32(72349) {
										if c < int32(72192) {
											if c < int32(72161) {
												if c < int32(72106) {
													v431 = libc.BoolInt32(c >= int32(72096) && c <= int32(72103))
												} else {
													v431 = libc.BoolInt32(c <= int32(72144))
												}
												v430 = v431
											} else {
												v430 = libc.BoolInt32(c <= int32(72161) || c == int32(72163))
											}
											v429 = v430
										} else {
											if v434 = c <= int32(72192); !v434 {
												if c < int32(72272) {
													if c < int32(72250) {
														v433 = libc.BoolInt32(c >= int32(72203) && c <= int32(72242))
													} else {
														v433 = libc.BoolInt32(c <= int32(72250))
													}
													v432 = v433
												} else {
													v432 = libc.BoolInt32(c <= int32(72272) || c >= int32(72284) && c <= int32(72329))
												}
											}
											v429 = libc.BoolInt32(v434 || v432 != 0)
										}
										v428 = v429
									} else {
										if v441 = c <= int32(72349); !v441 {
											if c < int32(72818) {
												if c < int32(72714) {
													if c < int32(72704) {
														v437 = libc.BoolInt32(c >= int32(72368) && c <= int32(72440))
													} else {
														v437 = libc.BoolInt32(c <= int32(72712))
													}
													v436 = v437
												} else {
													v436 = libc.BoolInt32(c <= int32(72750) || c == int32(72768))
												}
												v435 = v436
											} else {
												if v440 = c <= int32(72847); !v440 {
													if c < int32(72971) {
														if c < int32(72968) {
															v439 = libc.BoolInt32(c >= int32(72960) && c <= int32(72966))
														} else {
															v439 = libc.BoolInt32(c <= int32(72969))
														}
														v438 = v439
													} else {
														v438 = libc.BoolInt32(c <= int32(73008) || c == int32(73030))
													}
												}
												v435 = libc.BoolInt32(v440 || v438 != 0)
											}
										}
										v428 = libc.BoolInt32(v441 || v435 != 0)
									}
								}
								v411 = libc.BoolInt32(v442 || v428 != 0)
							}
							v410 = v411
						} else {
							if v475 = c <= int32(73061); !v475 {
								if c < int32(93952) {
									if c < int32(82944) {
										if c < int32(73728) {
											if c < int32(73112) {
												if c < int32(73066) {
													v447 = libc.BoolInt32(c >= int32(73063) && c <= int32(73064))
												} else {
													v447 = libc.BoolInt32(c <= int32(73097))
												}
												v446 = v447
											} else {
												if v449 = c <= int32(73112); !v449 {
													if c < int32(73648) {
														v448 = libc.BoolInt32(c >= int32(73440) && c <= int32(73458))
													} else {
														v448 = libc.BoolInt32(c <= int32(73648))
													}
												}
												v446 = libc.BoolInt32(v449 || v448 != 0)
											}
											v445 = v446
										} else {
											if v452 = c <= int32(74649); !v452 {
												if c < int32(77712) {
													if c < int32(74880) {
														v451 = libc.BoolInt32(c >= int32(74752) && c <= int32(74862))
													} else {
														v451 = libc.BoolInt32(c <= int32(75075))
													}
													v450 = v451
												} else {
													v450 = libc.BoolInt32(c <= int32(77808) || c >= int32(77824) && c <= int32(78894))
												}
											}
											v445 = libc.BoolInt32(v452 || v450 != 0)
										}
										v444 = v445
									} else {
										if v459 = c <= int32(83526); !v459 {
											if c < int32(92928) {
												if c < int32(92784) {
													if c < int32(92736) {
														v455 = libc.BoolInt32(c >= int32(92160) && c <= int32(92728))
													} else {
														v455 = libc.BoolInt32(c <= int32(92766))
													}
													v454 = v455
												} else {
													v454 = libc.BoolInt32(c <= int32(92862) || c >= int32(92880) && c <= int32(92909))
												}
												v453 = v454
											} else {
												if v458 = c <= int32(92975); !v458 {
													if c < int32(93053) {
														if c < int32(93027) {
															v457 = libc.BoolInt32(c >= int32(92992) && c <= int32(92995))
														} else {
															v457 = libc.BoolInt32(c <= int32(93047))
														}
														v456 = v457
													} else {
														v456 = libc.BoolInt32(c <= int32(93071) || c >= int32(93760) && c <= int32(93823))
													}
												}
												v453 = libc.BoolInt32(v458 || v456 != 0)
											}
										}
										v444 = libc.BoolInt32(v459 || v453 != 0)
									}
									v443 = v444
								} else {
									if v474 = c <= int32(94026); !v474 {
										if c < int32(110589) {
											if c < int32(94208) {
												if c < int32(94176) {
													if c < int32(94099) {
														v463 = libc.BoolInt32(c == int32(94032))
													} else {
														v463 = libc.BoolInt32(c <= int32(94111))
													}
													v462 = v463
												} else {
													v462 = libc.BoolInt32(c <= int32(94177) || c == int32(94179))
												}
												v461 = v462
											} else {
												if v466 = c <= int32(100343); !v466 {
													if c < int32(110576) {
														if c < int32(101632) {
															v465 = libc.BoolInt32(c >= int32(100352) && c <= int32(101589))
														} else {
															v465 = libc.BoolInt32(c <= int32(101640))
														}
														v464 = v465
													} else {
														v464 = libc.BoolInt32(c <= int32(110579) || c >= int32(110581) && c <= int32(110587))
													}
												}
												v461 = libc.BoolInt32(v466 || v464 != 0)
											}
											v460 = v461
										} else {
											if v473 = c <= int32(110590); !v473 {
												if c < int32(113664) {
													if c < int32(110948) {
														if c < int32(110928) {
															v469 = libc.BoolInt32(c >= int32(110592) && c <= int32(110882))
														} else {
															v469 = libc.BoolInt32(c <= int32(110930))
														}
														v468 = v469
													} else {
														v468 = libc.BoolInt32(c <= int32(110951) || c >= int32(110960) && c <= int32(111355))
													}
													v467 = v468
												} else {
													if v472 = c <= int32(113770); !v472 {
														if c < int32(113808) {
															if c < int32(113792) {
																v471 = libc.BoolInt32(c >= int32(113776) && c <= int32(113788))
															} else {
																v471 = libc.BoolInt32(c <= int32(113800))
															}
															v470 = v471
														} else {
															v470 = libc.BoolInt32(c <= int32(113817) || c >= int32(119808) && c <= int32(119892))
														}
													}
													v467 = libc.BoolInt32(v472 || v470 != 0)
												}
											}
											v460 = libc.BoolInt32(v473 || v467 != 0)
										}
									}
									v443 = libc.BoolInt32(v474 || v460 != 0)
								}
							}
							v410 = libc.BoolInt32(v475 || v443 != 0)
						}
						v409 = v410
					} else {
						if v542 = c <= int32(119964); !v542 {
							if c < int32(125259) {
								if c < int32(120572) {
									if c < int32(120086) {
										if c < int32(119995) {
											if c < int32(119973) {
												if c < int32(119970) {
													v481 = libc.BoolInt32(c >= int32(119966) && c <= int32(119967))
												} else {
													v481 = libc.BoolInt32(c <= int32(119970))
												}
												v480 = v481
											} else {
												if v483 = c <= int32(119974); !v483 {
													if c < int32(119982) {
														v482 = libc.BoolInt32(c >= int32(119977) && c <= int32(119980))
													} else {
														v482 = libc.BoolInt32(c <= int32(119993))
													}
												}
												v480 = libc.BoolInt32(v483 || v482 != 0)
											}
											v479 = v480
										} else {
											if v486 = c <= int32(119995); !v486 {
												if c < int32(120071) {
													if c < int32(120005) {
														v485 = libc.BoolInt32(c >= int32(119997) && c <= int32(120003))
													} else {
														v485 = libc.BoolInt32(c <= int32(120069))
													}
													v484 = v485
												} else {
													v484 = libc.BoolInt32(c <= int32(120074) || c >= int32(120077) && c <= int32(120084))
												}
											}
											v479 = libc.BoolInt32(v486 || v484 != 0)
										}
										v478 = v479
									} else {
										if v493 = c <= int32(120092); !v493 {
											if c < int32(120138) {
												if c < int32(120128) {
													if c < int32(120123) {
														v489 = libc.BoolInt32(c >= int32(120094) && c <= int32(120121))
													} else {
														v489 = libc.BoolInt32(c <= int32(120126))
													}
													v488 = v489
												} else {
													v488 = libc.BoolInt32(c <= int32(120132) || c == int32(120134))
												}
												v487 = v488
											} else {
												if v492 = c <= int32(120144); !v492 {
													if c < int32(120514) {
														if c < int32(120488) {
															v491 = libc.BoolInt32(c >= int32(120146) && c <= int32(120485))
														} else {
															v491 = libc.BoolInt32(c <= int32(120512))
														}
														v490 = v491
													} else {
														v490 = libc.BoolInt32(c <= int32(120538) || c >= int32(120540) && c <= int32(120570))
													}
												}
												v487 = libc.BoolInt32(v492 || v490 != 0)
											}
										}
										v478 = libc.BoolInt32(v493 || v487 != 0)
									}
									v477 = v478
								} else {
									if v508 = c <= int32(120596); !v508 {
										if c < int32(123191) {
											if c < int32(120714) {
												if c < int32(120656) {
													if c < int32(120630) {
														v497 = libc.BoolInt32(c >= int32(120598) && c <= int32(120628))
													} else {
														v497 = libc.BoolInt32(c <= int32(120654))
													}
													v496 = v497
												} else {
													v496 = libc.BoolInt32(c <= int32(120686) || c >= int32(120688) && c <= int32(120712))
												}
												v495 = v496
											} else {
												if v500 = c <= int32(120744); !v500 {
													if c < int32(122624) {
														if c < int32(120772) {
															v499 = libc.BoolInt32(c >= int32(120746) && c <= int32(120770))
														} else {
															v499 = libc.BoolInt32(c <= int32(120779))
														}
														v498 = v499
													} else {
														v498 = libc.BoolInt32(c <= int32(122654) || c >= int32(123136) && c <= int32(123180))
													}
												}
												v495 = libc.BoolInt32(v500 || v498 != 0)
											}
											v494 = v495
										} else {
											if v507 = c <= int32(123197); !v507 {
												if c < int32(124904) {
													if c < int32(123584) {
														if c < int32(123536) {
															v503 = libc.BoolInt32(c == int32(123214))
														} else {
															v503 = libc.BoolInt32(c <= int32(123565))
														}
														v502 = v503
													} else {
														v502 = libc.BoolInt32(c <= int32(123627) || c >= int32(124896) && c <= int32(124902))
													}
													v501 = v502
												} else {
													if v506 = c <= int32(124907); !v506 {
														if c < int32(124928) {
															if c < int32(124912) {
																v505 = libc.BoolInt32(c >= int32(124909) && c <= int32(124910))
															} else {
																v505 = libc.BoolInt32(c <= int32(124926))
															}
															v504 = v505
														} else {
															v504 = libc.BoolInt32(c <= int32(125124) || c >= int32(125184) && c <= int32(125251))
														}
													}
													v501 = libc.BoolInt32(v506 || v504 != 0)
												}
											}
											v494 = libc.BoolInt32(v507 || v501 != 0)
										}
									}
									v477 = libc.BoolInt32(v508 || v494 != 0)
								}
								v476 = v477
							} else {
								if v541 = c <= int32(125259); !v541 {
									if c < int32(126559) {
										if c < int32(126535) {
											if c < int32(126505) {
												if c < int32(126497) {
													if c < int32(126469) {
														v513 = libc.BoolInt32(c >= int32(126464) && c <= int32(126467))
													} else {
														v513 = libc.BoolInt32(c <= int32(126495))
													}
													v512 = v513
												} else {
													if v515 = c <= int32(126498); !v515 {
														if c < int32(126503) {
															v514 = libc.BoolInt32(c == int32(126500))
														} else {
															v514 = libc.BoolInt32(c <= int32(126503))
														}
													}
													v512 = libc.BoolInt32(v515 || v514 != 0)
												}
												v511 = v512
											} else {
												if v518 = c <= int32(126514); !v518 {
													if c < int32(126523) {
														if c < int32(126521) {
															v517 = libc.BoolInt32(c >= int32(126516) && c <= int32(126519))
														} else {
															v517 = libc.BoolInt32(c <= int32(126521))
														}
														v516 = v517
													} else {
														v516 = libc.BoolInt32(c <= int32(126523) || c == int32(126530))
													}
												}
												v511 = libc.BoolInt32(v518 || v516 != 0)
											}
											v510 = v511
										} else {
											if v525 = c <= int32(126535); !v525 {
												if c < int32(126548) {
													if c < int32(126541) {
														if c < int32(126539) {
															v521 = libc.BoolInt32(c == int32(126537))
														} else {
															v521 = libc.BoolInt32(c <= int32(126539))
														}
														v520 = v521
													} else {
														v520 = libc.BoolInt32(c <= int32(126543) || c >= int32(126545) && c <= int32(126546))
													}
													v519 = v520
												} else {
													if v524 = c <= int32(126548); !v524 {
														if c < int32(126555) {
															if c < int32(126553) {
																v523 = libc.BoolInt32(c == int32(126551))
															} else {
																v523 = libc.BoolInt32(c <= int32(126553))
															}
															v522 = v523
														} else {
															v522 = libc.BoolInt32(c <= int32(126555) || c == int32(126557))
														}
													}
													v519 = libc.BoolInt32(v524 || v522 != 0)
												}
											}
											v510 = libc.BoolInt32(v525 || v519 != 0)
										}
										v509 = v510
									} else {
										if v540 = c <= int32(126559); !v540 {
											if c < int32(126625) {
												if c < int32(126580) {
													if c < int32(126567) {
														if c < int32(126564) {
															v529 = libc.BoolInt32(c >= int32(126561) && c <= int32(126562))
														} else {
															v529 = libc.BoolInt32(c <= int32(126564))
														}
														v528 = v529
													} else {
														v528 = libc.BoolInt32(c <= int32(126570) || c >= int32(126572) && c <= int32(126578))
													}
													v527 = v528
												} else {
													if v532 = c <= int32(126583); !v532 {
														if c < int32(126592) {
															if c < int32(126590) {
																v531 = libc.BoolInt32(c >= int32(126585) && c <= int32(126588))
															} else {
																v531 = libc.BoolInt32(c <= int32(126590))
															}
															v530 = v531
														} else {
															v530 = libc.BoolInt32(c <= int32(126601) || c >= int32(126603) && c <= int32(126619))
														}
													}
													v527 = libc.BoolInt32(v532 || v530 != 0)
												}
												v526 = v527
											} else {
												if v539 = c <= int32(126627); !v539 {
													if c < int32(177984) {
														if c < int32(131072) {
															if c < int32(126635) {
																v535 = libc.BoolInt32(c >= int32(126629) && c <= int32(126633))
															} else {
																v535 = libc.BoolInt32(c <= int32(126651))
															}
															v534 = v535
														} else {
															v534 = libc.BoolInt32(c <= int32(173791) || c >= int32(173824) && c <= int32(177976))
														}
														v533 = v534
													} else {
														if v538 = c <= int32(178205); !v538 {
															if c < int32(194560) {
																if c < int32(183984) {
																	v537 = libc.BoolInt32(c >= int32(178208) && c <= int32(183969))
																} else {
																	v537 = libc.BoolInt32(c <= int32(191456))
																}
																v536 = v537
															} else {
																v536 = libc.BoolInt32(c <= int32(195101) || c >= int32(196608) && c <= int32(201546))
															}
														}
														v533 = libc.BoolInt32(v538 || v536 != 0)
													}
												}
												v526 = libc.BoolInt32(v539 || v533 != 0)
											}
										}
										v509 = libc.BoolInt32(v540 || v526 != 0)
									}
								}
								v476 = libc.BoolInt32(v541 || v509 != 0)
							}
						}
						v409 = libc.BoolInt32(v542 || v476 != 0)
					}
				}
				v274 = libc.BoolInt32(v543 || v409 != 0)
			}
		}
		v1 = libc.BoolInt32(v544 || v274 != 0)
	}
	return uint8(libc.BoolInt32(v1 != 0))
}

func sym_identifier_character_set_2(tls *libc.TLS, c int32_t) (r uint8) {
	var v1, v10, v100, v101, v102, v103, v104, v106, v107, v108, v111, v112, v113, v114, v116, v117, v118, v12, v122, v123, v124, v125, v126, v128, v129, v13, v130, v133, v134, v135, v136, v138, v139, v14, v140, v145, v146, v147, v148, v149, v150, v152, v153, v154, v157, v158, v159, v160, v162, v163, v164, v168, v169, v17, v170, v171, v172, v174, v175, v176, v179, v18, v180, v181, v182, v184, v185, v186, v19, v193, v194, v195, v196, v197, v198, v199, v2, v20, v200, v202, v203, v204, v207, v208, v209, v210, v212, v213, v214, v218, v219, v22, v220, v221, v222, v224, v225, v226, v229, v23, v230, v231, v232, v234, v235, v236, v24, v241, v242, v243, v244, v245, v246, v248, v249, v250, v253, v254, v255, v256, v258, v259, v260, v264, v265, v266, v267, v268, v270, v271, v272, v275, v276, v277, v278, v28, v280, v281, v282, v288, v289, v29, v290, v291, v292, v293, v294, v296, v297, v298, v3, v30, v301, v302, v303, v304, v306, v307, v308, v31, v312, v313, v314, v315, v316, v318, v319, v32, v320, v323, v324, v325, v326, v328, v329, v330, v335, v336, v337, v338, v339, v34, v340, v342, v343, v344, v347, v348, v349, v35, v350, v352, v353, v354, v358, v359, v36, v360, v361, v362, v364, v365, v366, v369, v370, v371, v372, v374, v375, v382, v383, v384, v385, v386, v387, v388, v389, v39, v390, v392, v393, v394, v397, v398, v399, v4, v40, v400, v402, v403, v404, v408, v409, v41, v410, v411, v412, v414, v415, v416, v419, v42, v420, v421, v422, v424, v425, v426, v431, v432, v433, v434, v435, v436, v438, v439, v44, v440, v443, v444, v445, v446, v448, v449, v45, v450, v454, v455, v456, v457, v458, v46, v460, v461, v462, v465, v466, v467, v468, v470, v471, v472, v478, v479, v480, v481, v482, v483, v484, v486, v487, v488, v491, v492, v493, v494, v496, v497, v498, v5, v502, v503, v504, v505, v506, v508, v509, v51, v510, v513, v514, v515, v516, v518, v519, v52, v520, v525, v526, v527, v528, v529, v53, v530, v532, v533, v534, v537, v538, v539, v54, v540, v542, v543, v544, v548, v549, v55, v550, v551, v552, v554, v555, v556, v559, v56, v560, v561, v562, v564, v565, v571, v572, v573, v574, v575, v576, v577, v578, v58, v580, v581, v582, v585, v586, v587, v588, v59, v590, v591, v592, v596, v597, v598, v599, v6, v60, v600, v602, v603, v604, v607, v608, v609, v610, v612, v613, v614, v619, v620, v621, v622, v623, v624, v626, v627, v628, v63, v631, v632, v633, v634, v636, v637, v638, v64, v642, v643, v644, v645, v646, v648, v649, v65, v650, v653, v654, v655, v656, v658, v659, v66, v660, v666, v667, v668, v669, v670, v671, v672, v674, v675, v676, v679, v68, v680, v681, v682, v684, v685, v686, v69, v690, v691, v692, v693, v694, v696, v697, v698, v7, v70, v701, v702, v703, v704, v706, v707, v708, v713, v714, v715, v716, v717, v718, v720, v721, v722, v725, v726, v727, v728, v730, v731, v732, v736, v737, v738, v739, v74, v740, v742, v743, v744, v747, v748, v749, v75, v750, v752, v753, v76, v77, v78, v8, v80, v81, v82, v85, v86, v87, v88, v9, v90, v91, v92, v98, v99 int32
	var v105, v109, v11, v110, v115, v119, v120, v121, v127, v131, v132, v137, v141, v142, v143, v144, v15, v151, v155, v156, v16, v161, v165, v166, v167, v173, v177, v178, v183, v187, v188, v189, v190, v191, v192, v201, v205, v206, v21, v211, v215, v216, v217, v223, v227, v228, v233, v237, v238, v239, v240, v247, v25, v251, v252, v257, v26, v261, v262, v263, v269, v27, v273, v274, v279, v283, v284, v285, v286, v287, v295, v299, v300, v305, v309, v310, v311, v317, v321, v322, v327, v33, v331, v332, v333, v334, v341, v345, v346, v351, v355, v356, v357, v363, v367, v368, v37, v373, v376, v377, v378, v379, v38, v380, v381, v391, v395, v396, v401, v405, v406, v407, v413, v417, v418, v423, v427, v428, v429, v43, v430, v437, v441, v442, v447, v451, v452, v453, v459, v463, v464, v469, v47, v473, v474, v475, v476, v477, v48, v485, v489, v49, v490, v495, v499, v50, v500, v501, v507, v511, v512, v517, v521, v522, v523, v524, v531, v535, v536, v541, v545, v546, v547, v553, v557, v558, v563, v566, v567, v568, v569, v57, v570, v579, v583, v584, v589, v593, v594, v595, v601, v605, v606, v61, v611, v615, v616, v617, v618, v62, v625, v629, v630, v635, v639, v640, v641, v647, v651, v652, v657, v661, v662, v663, v664, v665, v67, v673, v677, v678, v683, v687, v688, v689, v695, v699, v700, v705, v709, v71, v710, v711, v712, v719, v72, v723, v724, v729, v73, v733, v734, v735, v741, v745, v746, v751, v754, v755, v756, v757, v758, v759, v760, v79, v83, v84, v89, v93, v94, v95, v96, v97 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = v1, v10, v100, v101, v102, v103, v104, v105, v106, v107, v108, v109, v11, v110, v111, v112, v113, v114, v115, v116, v117, v118, v119, v12, v120, v121, v122, v123, v124, v125, v126, v127, v128, v129, v13, v130, v131, v132, v133, v134, v135, v136, v137, v138, v139, v14, v140, v141, v142, v143, v144, v145, v146, v147, v148, v149, v15, v150, v151, v152, v153, v154, v155, v156, v157, v158, v159, v16, v160, v161, v162, v163, v164, v165, v166, v167, v168, v169, v17, v170, v171, v172, v173, v174, v175, v176, v177, v178, v179, v18, v180, v181, v182, v183, v184, v185, v186, v187, v188, v189, v19, v190, v191, v192, v193, v194, v195, v196, v197, v198, v199, v2, v20, v200, v201, v202, v203, v204, v205, v206, v207, v208, v209, v21, v210, v211, v212, v213, v214, v215, v216, v217, v218, v219, v22, v220, v221, v222, v223, v224, v225, v226, v227, v228, v229, v23, v230, v231, v232, v233, v234, v235, v236, v237, v238, v239, v24, v240, v241, v242, v243, v244, v245, v246, v247, v248, v249, v25, v250, v251, v252, v253, v254, v255, v256, v257, v258, v259, v26, v260, v261, v262, v263, v264, v265, v266, v267, v268, v269, v27, v270, v271, v272, v273, v274, v275, v276, v277, v278, v279, v28, v280, v281, v282, v283, v284, v285, v286, v287, v288, v289, v29, v290, v291, v292, v293, v294, v295, v296, v297, v298, v299, v3, v30, v300, v301, v302, v303, v304, v305, v306, v307, v308, v309, v31, v310, v311, v312, v313, v314, v315, v316, v317, v318, v319, v32, v320, v321, v322, v323, v324, v325, v326, v327, v328, v329, v33, v330, v331, v332, v333, v334, v335, v336, v337, v338, v339, v34, v340, v341, v342, v343, v344, v345, v346, v347, v348, v349, v35, v350, v351, v352, v353, v354, v355, v356, v357, v358, v359, v36, v360, v361, v362, v363, v364, v365, v366, v367, v368, v369, v37, v370, v371, v372, v373, v374, v375, v376, v377, v378, v379, v38, v380, v381, v382, v383, v384, v385, v386, v387, v388, v389, v39, v390, v391, v392, v393, v394, v395, v396, v397, v398, v399, v4, v40, v400, v401, v402, v403, v404, v405, v406, v407, v408, v409, v41, v410, v411, v412, v413, v414, v415, v416, v417, v418, v419, v42, v420, v421, v422, v423, v424, v425, v426, v427, v428, v429, v43, v430, v431, v432, v433, v434, v435, v436, v437, v438, v439, v44, v440, v441, v442, v443, v444, v445, v446, v447, v448, v449, v45, v450, v451, v452, v453, v454, v455, v456, v457, v458, v459, v46, v460, v461, v462, v463, v464, v465, v466, v467, v468, v469, v47, v470, v471, v472, v473, v474, v475, v476, v477, v478, v479, v48, v480, v481, v482, v483, v484, v485, v486, v487, v488, v489, v49, v490, v491, v492, v493, v494, v495, v496, v497, v498, v499, v5, v50, v500, v501, v502, v503, v504, v505, v506, v507, v508, v509, v51, v510, v511, v512, v513, v514, v515, v516, v517, v518, v519, v52, v520, v521, v522, v523, v524, v525, v526, v527, v528, v529, v53, v530, v531, v532, v533, v534, v535, v536, v537, v538, v539, v54, v540, v541, v542, v543, v544, v545, v546, v547, v548, v549, v55, v550, v551, v552, v553, v554, v555, v556, v557, v558, v559, v56, v560, v561, v562, v563, v564, v565, v566, v567, v568, v569, v57, v570, v571, v572, v573, v574, v575, v576, v577, v578, v579, v58, v580, v581, v582, v583, v584, v585, v586, v587, v588, v589, v59, v590, v591, v592, v593, v594, v595, v596, v597, v598, v599, v6, v60, v600, v601, v602, v603, v604, v605, v606, v607, v608, v609, v61, v610, v611, v612, v613, v614, v615, v616, v617, v618, v619, v62, v620, v621, v622, v623, v624, v625, v626, v627, v628, v629, v63, v630, v631, v632, v633, v634, v635, v636, v637, v638, v639, v64, v640, v641, v642, v643, v644, v645, v646, v647, v648, v649, v65, v650, v651, v652, v653, v654, v655, v656, v657, v658, v659, v66, v660, v661, v662, v663, v664, v665, v666, v667, v668, v669, v67, v670, v671, v672, v673, v674, v675, v676, v677, v678, v679, v68, v680, v681, v682, v683, v684, v685, v686, v687, v688, v689, v69, v690, v691, v692, v693, v694, v695, v696, v697, v698, v699, v7, v70, v700, v701, v702, v703, v704, v705, v706, v707, v708, v709, v71, v710, v711, v712, v713, v714, v715, v716, v717, v718, v719, v72, v720, v721, v722, v723, v724, v725, v726, v727, v728, v729, v73, v730, v731, v732, v733, v734, v735, v736, v737, v738, v739, v74, v740, v741, v742, v743, v744, v745, v746, v747, v748, v749, v75, v750, v751, v752, v753, v754, v755, v756, v757, v758, v759, v76, v760, v77, v78, v79, v8, v80, v81, v82, v83, v84, v85, v86, v87, v88, v89, v9, v90, v91, v92, v93, v94, v95, v96, v97, v98, v99
	if c < int32(43616) {
		if c < int32(3782) {
			if c < int32(2741) {
				if c < int32(2042) {
					if c < int32(931) {
						if c < int32(248) {
							if c < int32(170) {
								if c < int32('A') {
									if c < int32('0') {
										v9 = libc.BoolInt32(c == int32('-'))
									} else {
										v9 = libc.BoolInt32(c <= int32('9'))
									}
									v8 = v9
								} else {
									if v11 = c <= int32('Z'); !v11 {
										if c < int32('a') {
											v10 = libc.BoolInt32(c == int32('_'))
										} else {
											v10 = libc.BoolInt32(c <= int32('z'))
										}
									}
									v8 = libc.BoolInt32(v11 || v10 != 0)
								}
								v7 = v8
							} else {
								if v16 = c <= int32(170); !v16 {
									if c < int32(186) {
										if c < int32(183) {
											v13 = libc.BoolInt32(c == int32(181))
										} else {
											v13 = libc.BoolInt32(c <= int32(183))
										}
										v12 = v13
									} else {
										if v15 = c <= int32(186); !v15 {
											if c < int32(216) {
												v14 = libc.BoolInt32(c >= int32(192) && c <= int32(214))
											} else {
												v14 = libc.BoolInt32(c <= int32(246))
											}
										}
										v12 = libc.BoolInt32(v15 || v14 != 0)
									}
								}
								v7 = libc.BoolInt32(v16 || v12 != 0)
							}
							v6 = v7
						} else {
							if v27 = c <= int32(705); !v27 {
								if c < int32(886) {
									if c < int32(748) {
										if c < int32(736) {
											v19 = libc.BoolInt32(c >= int32(710) && c <= int32(721))
										} else {
											v19 = libc.BoolInt32(c <= int32(740))
										}
										v18 = v19
									} else {
										if v21 = c <= int32(748); !v21 {
											if c < int32(768) {
												v20 = libc.BoolInt32(c == int32(750))
											} else {
												v20 = libc.BoolInt32(c <= int32(884))
											}
										}
										v18 = libc.BoolInt32(v21 || v20 != 0)
									}
									v17 = v18
								} else {
									if v26 = c <= int32(887); !v26 {
										if c < int32(902) {
											if c < int32(895) {
												v23 = libc.BoolInt32(c >= int32(891) && c <= int32(893))
											} else {
												v23 = libc.BoolInt32(c <= int32(895))
											}
											v22 = v23
										} else {
											if v25 = c <= int32(906); !v25 {
												if c < int32(910) {
													v24 = libc.BoolInt32(c == int32(908))
												} else {
													v24 = libc.BoolInt32(c <= int32(929))
												}
											}
											v22 = libc.BoolInt32(v25 || v24 != 0)
										}
									}
									v17 = libc.BoolInt32(v26 || v22 != 0)
								}
							}
							v6 = libc.BoolInt32(v27 || v17 != 0)
						}
						v5 = v6
					} else {
						if v50 = c <= int32(1013); !v50 {
							if c < int32(1488) {
								if c < int32(1376) {
									if c < int32(1162) {
										if c < int32(1155) {
											v31 = libc.BoolInt32(c >= int32(1015) && c <= int32(1153))
										} else {
											v31 = libc.BoolInt32(c <= int32(1159))
										}
										v30 = v31
									} else {
										if v33 = c <= int32(1327); !v33 {
											if c < int32(1369) {
												v32 = libc.BoolInt32(c >= int32(1329) && c <= int32(1366))
											} else {
												v32 = libc.BoolInt32(c <= int32(1369))
											}
										}
										v30 = libc.BoolInt32(v33 || v32 != 0)
									}
									v29 = v30
								} else {
									if v38 = c <= int32(1416); !v38 {
										if c < int32(1473) {
											if c < int32(1471) {
												v35 = libc.BoolInt32(c >= int32(1425) && c <= int32(1469))
											} else {
												v35 = libc.BoolInt32(c <= int32(1471))
											}
											v34 = v35
										} else {
											if v37 = c <= int32(1474); !v37 {
												if c < int32(1479) {
													v36 = libc.BoolInt32(c >= int32(1476) && c <= int32(1477))
												} else {
													v36 = libc.BoolInt32(c <= int32(1479))
												}
											}
											v34 = libc.BoolInt32(v37 || v36 != 0)
										}
									}
									v29 = libc.BoolInt32(v38 || v34 != 0)
								}
								v28 = v29
							} else {
								if v49 = c <= int32(1514); !v49 {
									if c < int32(1759) {
										if c < int32(1568) {
											if c < int32(1552) {
												v41 = libc.BoolInt32(c >= int32(1519) && c <= int32(1522))
											} else {
												v41 = libc.BoolInt32(c <= int32(1562))
											}
											v40 = v41
										} else {
											if v43 = c <= int32(1641); !v43 {
												if c < int32(1749) {
													v42 = libc.BoolInt32(c >= int32(1646) && c <= int32(1747))
												} else {
													v42 = libc.BoolInt32(c <= int32(1756))
												}
											}
											v40 = libc.BoolInt32(v43 || v42 != 0)
										}
										v39 = v40
									} else {
										if v48 = c <= int32(1768); !v48 {
											if c < int32(1808) {
												if c < int32(1791) {
													v45 = libc.BoolInt32(c >= int32(1770) && c <= int32(1788))
												} else {
													v45 = libc.BoolInt32(c <= int32(1791))
												}
												v44 = v45
											} else {
												if v47 = c <= int32(1866); !v47 {
													if c < int32(1984) {
														v46 = libc.BoolInt32(c >= int32(1869) && c <= int32(1969))
													} else {
														v46 = libc.BoolInt32(c <= int32(2037))
													}
												}
												v44 = libc.BoolInt32(v47 || v46 != 0)
											}
										}
										v39 = libc.BoolInt32(v48 || v44 != 0)
									}
								}
								v28 = libc.BoolInt32(v49 || v39 != 0)
							}
						}
						v5 = libc.BoolInt32(v50 || v28 != 0)
					}
					v4 = v5
				} else {
					if v97 = c <= int32(2042); !v97 {
						if c < int32(2556) {
							if c < int32(2447) {
								if c < int32(2185) {
									if c < int32(2112) {
										if c < int32(2048) {
											v55 = libc.BoolInt32(c == int32(2045))
										} else {
											v55 = libc.BoolInt32(c <= int32(2093))
										}
										v54 = v55
									} else {
										if v57 = c <= int32(2139); !v57 {
											if c < int32(2160) {
												v56 = libc.BoolInt32(c >= int32(2144) && c <= int32(2154))
											} else {
												v56 = libc.BoolInt32(c <= int32(2183))
											}
										}
										v54 = libc.BoolInt32(v57 || v56 != 0)
									}
									v53 = v54
								} else {
									if v62 = c <= int32(2190); !v62 {
										if c < int32(2406) {
											if c < int32(2275) {
												v59 = libc.BoolInt32(c >= int32(2200) && c <= int32(2273))
											} else {
												v59 = libc.BoolInt32(c <= int32(2403))
											}
											v58 = v59
										} else {
											if v61 = c <= int32(2415); !v61 {
												if c < int32(2437) {
													v60 = libc.BoolInt32(c >= int32(2417) && c <= int32(2435))
												} else {
													v60 = libc.BoolInt32(c <= int32(2444))
												}
											}
											v58 = libc.BoolInt32(v61 || v60 != 0)
										}
									}
									v53 = libc.BoolInt32(v62 || v58 != 0)
								}
								v52 = v53
							} else {
								if v73 = c <= int32(2448); !v73 {
									if c < int32(2503) {
										if c < int32(2482) {
											if c < int32(2474) {
												v65 = libc.BoolInt32(c >= int32(2451) && c <= int32(2472))
											} else {
												v65 = libc.BoolInt32(c <= int32(2480))
											}
											v64 = v65
										} else {
											if v67 = c <= int32(2482); !v67 {
												if c < int32(2492) {
													v66 = libc.BoolInt32(c >= int32(2486) && c <= int32(2489))
												} else {
													v66 = libc.BoolInt32(c <= int32(2500))
												}
											}
											v64 = libc.BoolInt32(v67 || v66 != 0)
										}
										v63 = v64
									} else {
										if v72 = c <= int32(2504); !v72 {
											if c < int32(2524) {
												if c < int32(2519) {
													v69 = libc.BoolInt32(c >= int32(2507) && c <= int32(2510))
												} else {
													v69 = libc.BoolInt32(c <= int32(2519))
												}
												v68 = v69
											} else {
												if v71 = c <= int32(2525); !v71 {
													if c < int32(2534) {
														v70 = libc.BoolInt32(c >= int32(2527) && c <= int32(2531))
													} else {
														v70 = libc.BoolInt32(c <= int32(2545))
													}
												}
												v68 = libc.BoolInt32(v71 || v70 != 0)
											}
										}
										v63 = libc.BoolInt32(v72 || v68 != 0)
									}
								}
								v52 = libc.BoolInt32(v73 || v63 != 0)
							}
							v51 = v52
						} else {
							if v96 = c <= int32(2556); !v96 {
								if c < int32(2631) {
									if c < int32(2602) {
										if c < int32(2565) {
											if c < int32(2561) {
												v77 = libc.BoolInt32(c == int32(2558))
											} else {
												v77 = libc.BoolInt32(c <= int32(2563))
											}
											v76 = v77
										} else {
											if v79 = c <= int32(2570); !v79 {
												if c < int32(2579) {
													v78 = libc.BoolInt32(c >= int32(2575) && c <= int32(2576))
												} else {
													v78 = libc.BoolInt32(c <= int32(2600))
												}
											}
											v76 = libc.BoolInt32(v79 || v78 != 0)
										}
										v75 = v76
									} else {
										if v84 = c <= int32(2608); !v84 {
											if c < int32(2616) {
												if c < int32(2613) {
													v81 = libc.BoolInt32(c >= int32(2610) && c <= int32(2611))
												} else {
													v81 = libc.BoolInt32(c <= int32(2614))
												}
												v80 = v81
											} else {
												if v83 = c <= int32(2617); !v83 {
													if c < int32(2622) {
														v82 = libc.BoolInt32(c == int32(2620))
													} else {
														v82 = libc.BoolInt32(c <= int32(2626))
													}
												}
												v80 = libc.BoolInt32(v83 || v82 != 0)
											}
										}
										v75 = libc.BoolInt32(v84 || v80 != 0)
									}
									v74 = v75
								} else {
									if v95 = c <= int32(2632); !v95 {
										if c < int32(2689) {
											if c < int32(2649) {
												if c < int32(2641) {
													v87 = libc.BoolInt32(c >= int32(2635) && c <= int32(2637))
												} else {
													v87 = libc.BoolInt32(c <= int32(2641))
												}
												v86 = v87
											} else {
												if v89 = c <= int32(2652); !v89 {
													if c < int32(2662) {
														v88 = libc.BoolInt32(c == int32(2654))
													} else {
														v88 = libc.BoolInt32(c <= int32(2677))
													}
												}
												v86 = libc.BoolInt32(v89 || v88 != 0)
											}
											v85 = v86
										} else {
											if v94 = c <= int32(2691); !v94 {
												if c < int32(2707) {
													if c < int32(2703) {
														v91 = libc.BoolInt32(c >= int32(2693) && c <= int32(2701))
													} else {
														v91 = libc.BoolInt32(c <= int32(2705))
													}
													v90 = v91
												} else {
													if v93 = c <= int32(2728); !v93 {
														if c < int32(2738) {
															v92 = libc.BoolInt32(c >= int32(2730) && c <= int32(2736))
														} else {
															v92 = libc.BoolInt32(c <= int32(2739))
														}
													}
													v90 = libc.BoolInt32(v93 || v92 != 0)
												}
											}
											v85 = libc.BoolInt32(v94 || v90 != 0)
										}
									}
									v74 = libc.BoolInt32(v95 || v85 != 0)
								}
							}
							v51 = libc.BoolInt32(v96 || v74 != 0)
						}
					}
					v4 = libc.BoolInt32(v97 || v51 != 0)
				}
				v3 = v4
			} else {
				if v192 = c <= int32(2745); !v192 {
					if c < int32(3165) {
						if c < int32(2949) {
							if c < int32(2858) {
								if c < int32(2790) {
									if c < int32(2763) {
										if c < int32(2759) {
											v103 = libc.BoolInt32(c >= int32(2748) && c <= int32(2757))
										} else {
											v103 = libc.BoolInt32(c <= int32(2761))
										}
										v102 = v103
									} else {
										if v105 = c <= int32(2765); !v105 {
											if c < int32(2784) {
												v104 = libc.BoolInt32(c == int32(2768))
											} else {
												v104 = libc.BoolInt32(c <= int32(2787))
											}
										}
										v102 = libc.BoolInt32(v105 || v104 != 0)
									}
									v101 = v102
								} else {
									if v110 = c <= int32(2799); !v110 {
										if c < int32(2821) {
											if c < int32(2817) {
												v107 = libc.BoolInt32(c >= int32(2809) && c <= int32(2815))
											} else {
												v107 = libc.BoolInt32(c <= int32(2819))
											}
											v106 = v107
										} else {
											if v109 = c <= int32(2828); !v109 {
												if c < int32(2835) {
													v108 = libc.BoolInt32(c >= int32(2831) && c <= int32(2832))
												} else {
													v108 = libc.BoolInt32(c <= int32(2856))
												}
											}
											v106 = libc.BoolInt32(v109 || v108 != 0)
										}
									}
									v101 = libc.BoolInt32(v110 || v106 != 0)
								}
								v100 = v101
							} else {
								if v121 = c <= int32(2864); !v121 {
									if c < int32(2901) {
										if c < int32(2876) {
											if c < int32(2869) {
												v113 = libc.BoolInt32(c >= int32(2866) && c <= int32(2867))
											} else {
												v113 = libc.BoolInt32(c <= int32(2873))
											}
											v112 = v113
										} else {
											if v115 = c <= int32(2884); !v115 {
												if c < int32(2891) {
													v114 = libc.BoolInt32(c >= int32(2887) && c <= int32(2888))
												} else {
													v114 = libc.BoolInt32(c <= int32(2893))
												}
											}
											v112 = libc.BoolInt32(v115 || v114 != 0)
										}
										v111 = v112
									} else {
										if v120 = c <= int32(2903); !v120 {
											if c < int32(2918) {
												if c < int32(2911) {
													v117 = libc.BoolInt32(c >= int32(2908) && c <= int32(2909))
												} else {
													v117 = libc.BoolInt32(c <= int32(2915))
												}
												v116 = v117
											} else {
												if v119 = c <= int32(2927); !v119 {
													if c < int32(2946) {
														v118 = libc.BoolInt32(c == int32(2929))
													} else {
														v118 = libc.BoolInt32(c <= int32(2947))
													}
												}
												v116 = libc.BoolInt32(v119 || v118 != 0)
											}
										}
										v111 = libc.BoolInt32(v120 || v116 != 0)
									}
								}
								v100 = libc.BoolInt32(v121 || v111 != 0)
							}
							v99 = v100
						} else {
							if v144 = c <= int32(2954); !v144 {
								if c < int32(3024) {
									if c < int32(2979) {
										if c < int32(2969) {
											if c < int32(2962) {
												v125 = libc.BoolInt32(c >= int32(2958) && c <= int32(2960))
											} else {
												v125 = libc.BoolInt32(c <= int32(2965))
											}
											v124 = v125
										} else {
											if v127 = c <= int32(2970); !v127 {
												if c < int32(2974) {
													v126 = libc.BoolInt32(c == int32(2972))
												} else {
													v126 = libc.BoolInt32(c <= int32(2975))
												}
											}
											v124 = libc.BoolInt32(v127 || v126 != 0)
										}
										v123 = v124
									} else {
										if v132 = c <= int32(2980); !v132 {
											if c < int32(3006) {
												if c < int32(2990) {
													v129 = libc.BoolInt32(c >= int32(2984) && c <= int32(2986))
												} else {
													v129 = libc.BoolInt32(c <= int32(3001))
												}
												v128 = v129
											} else {
												if v131 = c <= int32(3010); !v131 {
													if c < int32(3018) {
														v130 = libc.BoolInt32(c >= int32(3014) && c <= int32(3016))
													} else {
														v130 = libc.BoolInt32(c <= int32(3021))
													}
												}
												v128 = libc.BoolInt32(v131 || v130 != 0)
											}
										}
										v123 = libc.BoolInt32(v132 || v128 != 0)
									}
									v122 = v123
								} else {
									if v143 = c <= int32(3024); !v143 {
										if c < int32(3114) {
											if c < int32(3072) {
												if c < int32(3046) {
													v135 = libc.BoolInt32(c == int32(3031))
												} else {
													v135 = libc.BoolInt32(c <= int32(3055))
												}
												v134 = v135
											} else {
												if v137 = c <= int32(3084); !v137 {
													if c < int32(3090) {
														v136 = libc.BoolInt32(c >= int32(3086) && c <= int32(3088))
													} else {
														v136 = libc.BoolInt32(c <= int32(3112))
													}
												}
												v134 = libc.BoolInt32(v137 || v136 != 0)
											}
											v133 = v134
										} else {
											if v142 = c <= int32(3129); !v142 {
												if c < int32(3146) {
													if c < int32(3142) {
														v139 = libc.BoolInt32(c >= int32(3132) && c <= int32(3140))
													} else {
														v139 = libc.BoolInt32(c <= int32(3144))
													}
													v138 = v139
												} else {
													if v141 = c <= int32(3149); !v141 {
														if c < int32(3160) {
															v140 = libc.BoolInt32(c >= int32(3157) && c <= int32(3158))
														} else {
															v140 = libc.BoolInt32(c <= int32(3162))
														}
													}
													v138 = libc.BoolInt32(v141 || v140 != 0)
												}
											}
											v133 = libc.BoolInt32(v142 || v138 != 0)
										}
									}
									v122 = libc.BoolInt32(v143 || v133 != 0)
								}
							}
							v99 = libc.BoolInt32(v144 || v122 != 0)
						}
						v98 = v99
					} else {
						if v191 = c <= int32(3165); !v191 {
							if c < int32(3430) {
								if c < int32(3285) {
									if c < int32(3218) {
										if c < int32(3200) {
											if c < int32(3174) {
												v149 = libc.BoolInt32(c >= int32(3168) && c <= int32(3171))
											} else {
												v149 = libc.BoolInt32(c <= int32(3183))
											}
											v148 = v149
										} else {
											if v151 = c <= int32(3203); !v151 {
												if c < int32(3214) {
													v150 = libc.BoolInt32(c >= int32(3205) && c <= int32(3212))
												} else {
													v150 = libc.BoolInt32(c <= int32(3216))
												}
											}
											v148 = libc.BoolInt32(v151 || v150 != 0)
										}
										v147 = v148
									} else {
										if v156 = c <= int32(3240); !v156 {
											if c < int32(3260) {
												if c < int32(3253) {
													v153 = libc.BoolInt32(c >= int32(3242) && c <= int32(3251))
												} else {
													v153 = libc.BoolInt32(c <= int32(3257))
												}
												v152 = v153
											} else {
												if v155 = c <= int32(3268); !v155 {
													if c < int32(3274) {
														v154 = libc.BoolInt32(c >= int32(3270) && c <= int32(3272))
													} else {
														v154 = libc.BoolInt32(c <= int32(3277))
													}
												}
												v152 = libc.BoolInt32(v155 || v154 != 0)
											}
										}
										v147 = libc.BoolInt32(v156 || v152 != 0)
									}
									v146 = v147
								} else {
									if v167 = c <= int32(3286); !v167 {
										if c < int32(3342) {
											if c < int32(3302) {
												if c < int32(3296) {
													v159 = libc.BoolInt32(c >= int32(3293) && c <= int32(3294))
												} else {
													v159 = libc.BoolInt32(c <= int32(3299))
												}
												v158 = v159
											} else {
												if v161 = c <= int32(3311); !v161 {
													if c < int32(3328) {
														v160 = libc.BoolInt32(c >= int32(3313) && c <= int32(3314))
													} else {
														v160 = libc.BoolInt32(c <= int32(3340))
													}
												}
												v158 = libc.BoolInt32(v161 || v160 != 0)
											}
											v157 = v158
										} else {
											if v166 = c <= int32(3344); !v166 {
												if c < int32(3402) {
													if c < int32(3398) {
														v163 = libc.BoolInt32(c >= int32(3346) && c <= int32(3396))
													} else {
														v163 = libc.BoolInt32(c <= int32(3400))
													}
													v162 = v163
												} else {
													if v165 = c <= int32(3406); !v165 {
														if c < int32(3423) {
															v164 = libc.BoolInt32(c >= int32(3412) && c <= int32(3415))
														} else {
															v164 = libc.BoolInt32(c <= int32(3427))
														}
													}
													v162 = libc.BoolInt32(v165 || v164 != 0)
												}
											}
											v157 = libc.BoolInt32(v166 || v162 != 0)
										}
									}
									v146 = libc.BoolInt32(v167 || v157 != 0)
								}
								v145 = v146
							} else {
								if v190 = c <= int32(3439); !v190 {
									if c < int32(3558) {
										if c < int32(3517) {
											if c < int32(3461) {
												if c < int32(3457) {
													v171 = libc.BoolInt32(c >= int32(3450) && c <= int32(3455))
												} else {
													v171 = libc.BoolInt32(c <= int32(3459))
												}
												v170 = v171
											} else {
												if v173 = c <= int32(3478); !v173 {
													if c < int32(3507) {
														v172 = libc.BoolInt32(c >= int32(3482) && c <= int32(3505))
													} else {
														v172 = libc.BoolInt32(c <= int32(3515))
													}
												}
												v170 = libc.BoolInt32(v173 || v172 != 0)
											}
											v169 = v170
										} else {
											if v178 = c <= int32(3517); !v178 {
												if c < int32(3535) {
													if c < int32(3530) {
														v175 = libc.BoolInt32(c >= int32(3520) && c <= int32(3526))
													} else {
														v175 = libc.BoolInt32(c <= int32(3530))
													}
													v174 = v175
												} else {
													if v177 = c <= int32(3540); !v177 {
														if c < int32(3544) {
															v176 = libc.BoolInt32(c == int32(3542))
														} else {
															v176 = libc.BoolInt32(c <= int32(3551))
														}
													}
													v174 = libc.BoolInt32(v177 || v176 != 0)
												}
											}
											v169 = libc.BoolInt32(v178 || v174 != 0)
										}
										v168 = v169
									} else {
										if v189 = c <= int32(3567); !v189 {
											if c < int32(3716) {
												if c < int32(3648) {
													if c < int32(3585) {
														v181 = libc.BoolInt32(c >= int32(3570) && c <= int32(3571))
													} else {
														v181 = libc.BoolInt32(c <= int32(3642))
													}
													v180 = v181
												} else {
													if v183 = c <= int32(3662); !v183 {
														if c < int32(3713) {
															v182 = libc.BoolInt32(c >= int32(3664) && c <= int32(3673))
														} else {
															v182 = libc.BoolInt32(c <= int32(3714))
														}
													}
													v180 = libc.BoolInt32(v183 || v182 != 0)
												}
												v179 = v180
											} else {
												if v188 = c <= int32(3716); !v188 {
													if c < int32(3749) {
														if c < int32(3724) {
															v185 = libc.BoolInt32(c >= int32(3718) && c <= int32(3722))
														} else {
															v185 = libc.BoolInt32(c <= int32(3747))
														}
														v184 = v185
													} else {
														if v187 = c <= int32(3749); !v187 {
															if c < int32(3776) {
																v186 = libc.BoolInt32(c >= int32(3751) && c <= int32(3773))
															} else {
																v186 = libc.BoolInt32(c <= int32(3780))
															}
														}
														v184 = libc.BoolInt32(v187 || v186 != 0)
													}
												}
												v179 = libc.BoolInt32(v188 || v184 != 0)
											}
										}
										v168 = libc.BoolInt32(v189 || v179 != 0)
									}
								}
								v145 = libc.BoolInt32(v190 || v168 != 0)
							}
						}
						v98 = libc.BoolInt32(v191 || v145 != 0)
					}
				}
				v3 = libc.BoolInt32(v192 || v98 != 0)
			}
			v2 = v3
		} else {
			if v381 = c <= int32(3782); !v381 {
				if c < int32(8025) {
					if c < int32(5888) {
						if c < int32(4688) {
							if c < int32(3953) {
								if c < int32(3872) {
									if c < int32(3804) {
										if c < int32(3792) {
											v199 = libc.BoolInt32(c >= int32(3784) && c <= int32(3789))
										} else {
											v199 = libc.BoolInt32(c <= int32(3801))
										}
										v198 = v199
									} else {
										if v201 = c <= int32(3807); !v201 {
											if c < int32(3864) {
												v200 = libc.BoolInt32(c == int32(3840))
											} else {
												v200 = libc.BoolInt32(c <= int32(3865))
											}
										}
										v198 = libc.BoolInt32(v201 || v200 != 0)
									}
									v197 = v198
								} else {
									if v206 = c <= int32(3881); !v206 {
										if c < int32(3897) {
											if c < int32(3895) {
												v203 = libc.BoolInt32(c == int32(3893))
											} else {
												v203 = libc.BoolInt32(c <= int32(3895))
											}
											v202 = v203
										} else {
											if v205 = c <= int32(3897); !v205 {
												if c < int32(3913) {
													v204 = libc.BoolInt32(c >= int32(3902) && c <= int32(3911))
												} else {
													v204 = libc.BoolInt32(c <= int32(3948))
												}
											}
											v202 = libc.BoolInt32(v205 || v204 != 0)
										}
									}
									v197 = libc.BoolInt32(v206 || v202 != 0)
								}
								v196 = v197
							} else {
								if v217 = c <= int32(3972); !v217 {
									if c < int32(4256) {
										if c < int32(4038) {
											if c < int32(3993) {
												v209 = libc.BoolInt32(c >= int32(3974) && c <= int32(3991))
											} else {
												v209 = libc.BoolInt32(c <= int32(4028))
											}
											v208 = v209
										} else {
											if v211 = c <= int32(4038); !v211 {
												if c < int32(4176) {
													v210 = libc.BoolInt32(c >= int32(4096) && c <= int32(4169))
												} else {
													v210 = libc.BoolInt32(c <= int32(4253))
												}
											}
											v208 = libc.BoolInt32(v211 || v210 != 0)
										}
										v207 = v208
									} else {
										if v216 = c <= int32(4293); !v216 {
											if c < int32(4304) {
												if c < int32(4301) {
													v213 = libc.BoolInt32(c == int32(4295))
												} else {
													v213 = libc.BoolInt32(c <= int32(4301))
												}
												v212 = v213
											} else {
												if v215 = c <= int32(4346); !v215 {
													if c < int32(4682) {
														v214 = libc.BoolInt32(c >= int32(4348) && c <= int32(4680))
													} else {
														v214 = libc.BoolInt32(c <= int32(4685))
													}
												}
												v212 = libc.BoolInt32(v215 || v214 != 0)
											}
										}
										v207 = libc.BoolInt32(v216 || v212 != 0)
									}
								}
								v196 = libc.BoolInt32(v217 || v207 != 0)
							}
							v195 = v196
						} else {
							if v240 = c <= int32(4694); !v240 {
								if c < int32(4882) {
									if c < int32(4786) {
										if c < int32(4704) {
											if c < int32(4698) {
												v221 = libc.BoolInt32(c == int32(4696))
											} else {
												v221 = libc.BoolInt32(c <= int32(4701))
											}
											v220 = v221
										} else {
											if v223 = c <= int32(4744); !v223 {
												if c < int32(4752) {
													v222 = libc.BoolInt32(c >= int32(4746) && c <= int32(4749))
												} else {
													v222 = libc.BoolInt32(c <= int32(4784))
												}
											}
											v220 = libc.BoolInt32(v223 || v222 != 0)
										}
										v219 = v220
									} else {
										if v228 = c <= int32(4789); !v228 {
											if c < int32(4802) {
												if c < int32(4800) {
													v225 = libc.BoolInt32(c >= int32(4792) && c <= int32(4798))
												} else {
													v225 = libc.BoolInt32(c <= int32(4800))
												}
												v224 = v225
											} else {
												if v227 = c <= int32(4805); !v227 {
													if c < int32(4824) {
														v226 = libc.BoolInt32(c >= int32(4808) && c <= int32(4822))
													} else {
														v226 = libc.BoolInt32(c <= int32(4880))
													}
												}
												v224 = libc.BoolInt32(v227 || v226 != 0)
											}
										}
										v219 = libc.BoolInt32(v228 || v224 != 0)
									}
									v218 = v219
								} else {
									if v239 = c <= int32(4885); !v239 {
										if c < int32(5112) {
											if c < int32(4969) {
												if c < int32(4957) {
													v231 = libc.BoolInt32(c >= int32(4888) && c <= int32(4954))
												} else {
													v231 = libc.BoolInt32(c <= int32(4959))
												}
												v230 = v231
											} else {
												if v233 = c <= int32(4977); !v233 {
													if c < int32(5024) {
														v232 = libc.BoolInt32(c >= int32(4992) && c <= int32(5007))
													} else {
														v232 = libc.BoolInt32(c <= int32(5109))
													}
												}
												v230 = libc.BoolInt32(v233 || v232 != 0)
											}
											v229 = v230
										} else {
											if v238 = c <= int32(5117); !v238 {
												if c < int32(5761) {
													if c < int32(5743) {
														v235 = libc.BoolInt32(c >= int32(5121) && c <= int32(5740))
													} else {
														v235 = libc.BoolInt32(c <= int32(5759))
													}
													v234 = v235
												} else {
													if v237 = c <= int32(5786); !v237 {
														if c < int32(5870) {
															v236 = libc.BoolInt32(c >= int32(5792) && c <= int32(5866))
														} else {
															v236 = libc.BoolInt32(c <= int32(5880))
														}
													}
													v234 = libc.BoolInt32(v237 || v236 != 0)
												}
											}
											v229 = libc.BoolInt32(v238 || v234 != 0)
										}
									}
									v218 = libc.BoolInt32(v239 || v229 != 0)
								}
							}
							v195 = libc.BoolInt32(v240 || v218 != 0)
						}
						v194 = v195
					} else {
						if v287 = c <= int32(5909); !v287 {
							if c < int32(6688) {
								if c < int32(6176) {
									if c < int32(6016) {
										if c < int32(5984) {
											if c < int32(5952) {
												v245 = libc.BoolInt32(c >= int32(5919) && c <= int32(5940))
											} else {
												v245 = libc.BoolInt32(c <= int32(5971))
											}
											v244 = v245
										} else {
											if v247 = c <= int32(5996); !v247 {
												if c < int32(6002) {
													v246 = libc.BoolInt32(c >= int32(5998) && c <= int32(6000))
												} else {
													v246 = libc.BoolInt32(c <= int32(6003))
												}
											}
											v244 = libc.BoolInt32(v247 || v246 != 0)
										}
										v243 = v244
									} else {
										if v252 = c <= int32(6099); !v252 {
											if c < int32(6112) {
												if c < int32(6108) {
													v249 = libc.BoolInt32(c == int32(6103))
												} else {
													v249 = libc.BoolInt32(c <= int32(6109))
												}
												v248 = v249
											} else {
												if v251 = c <= int32(6121); !v251 {
													if c < int32(6159) {
														v250 = libc.BoolInt32(c >= int32(6155) && c <= int32(6157))
													} else {
														v250 = libc.BoolInt32(c <= int32(6169))
													}
												}
												v248 = libc.BoolInt32(v251 || v250 != 0)
											}
										}
										v243 = libc.BoolInt32(v252 || v248 != 0)
									}
									v242 = v243
								} else {
									if v263 = c <= int32(6264); !v263 {
										if c < int32(6470) {
											if c < int32(6400) {
												if c < int32(6320) {
													v255 = libc.BoolInt32(c >= int32(6272) && c <= int32(6314))
												} else {
													v255 = libc.BoolInt32(c <= int32(6389))
												}
												v254 = v255
											} else {
												if v257 = c <= int32(6430); !v257 {
													if c < int32(6448) {
														v256 = libc.BoolInt32(c >= int32(6432) && c <= int32(6443))
													} else {
														v256 = libc.BoolInt32(c <= int32(6459))
													}
												}
												v254 = libc.BoolInt32(v257 || v256 != 0)
											}
											v253 = v254
										} else {
											if v262 = c <= int32(6509); !v262 {
												if c < int32(6576) {
													if c < int32(6528) {
														v259 = libc.BoolInt32(c >= int32(6512) && c <= int32(6516))
													} else {
														v259 = libc.BoolInt32(c <= int32(6571))
													}
													v258 = v259
												} else {
													if v261 = c <= int32(6601); !v261 {
														if c < int32(6656) {
															v260 = libc.BoolInt32(c >= int32(6608) && c <= int32(6618))
														} else {
															v260 = libc.BoolInt32(c <= int32(6683))
														}
													}
													v258 = libc.BoolInt32(v261 || v260 != 0)
												}
											}
											v253 = libc.BoolInt32(v262 || v258 != 0)
										}
									}
									v242 = libc.BoolInt32(v263 || v253 != 0)
								}
								v241 = v242
							} else {
								if v286 = c <= int32(6750); !v286 {
									if c < int32(7232) {
										if c < int32(6847) {
											if c < int32(6800) {
												if c < int32(6783) {
													v267 = libc.BoolInt32(c >= int32(6752) && c <= int32(6780))
												} else {
													v267 = libc.BoolInt32(c <= int32(6793))
												}
												v266 = v267
											} else {
												if v269 = c <= int32(6809); !v269 {
													if c < int32(6832) {
														v268 = libc.BoolInt32(c == int32(6823))
													} else {
														v268 = libc.BoolInt32(c <= int32(6845))
													}
												}
												v266 = libc.BoolInt32(v269 || v268 != 0)
											}
											v265 = v266
										} else {
											if v274 = c <= int32(6862); !v274 {
												if c < int32(7019) {
													if c < int32(6992) {
														v271 = libc.BoolInt32(c >= int32(6912) && c <= int32(6988))
													} else {
														v271 = libc.BoolInt32(c <= int32(7001))
													}
													v270 = v271
												} else {
													if v273 = c <= int32(7027); !v273 {
														if c < int32(7168) {
															v272 = libc.BoolInt32(c >= int32(7040) && c <= int32(7155))
														} else {
															v272 = libc.BoolInt32(c <= int32(7223))
														}
													}
													v270 = libc.BoolInt32(v273 || v272 != 0)
												}
											}
											v265 = libc.BoolInt32(v274 || v270 != 0)
										}
										v264 = v265
									} else {
										if v285 = c <= int32(7241); !v285 {
											if c < int32(7380) {
												if c < int32(7312) {
													if c < int32(7296) {
														v277 = libc.BoolInt32(c >= int32(7245) && c <= int32(7293))
													} else {
														v277 = libc.BoolInt32(c <= int32(7304))
													}
													v276 = v277
												} else {
													if v279 = c <= int32(7354); !v279 {
														if c < int32(7376) {
															v278 = libc.BoolInt32(c >= int32(7357) && c <= int32(7359))
														} else {
															v278 = libc.BoolInt32(c <= int32(7378))
														}
													}
													v276 = libc.BoolInt32(v279 || v278 != 0)
												}
												v275 = v276
											} else {
												if v284 = c <= int32(7418); !v284 {
													if c < int32(7968) {
														if c < int32(7960) {
															v281 = libc.BoolInt32(c >= int32(7424) && c <= int32(7957))
														} else {
															v281 = libc.BoolInt32(c <= int32(7965))
														}
														v280 = v281
													} else {
														if v283 = c <= int32(8005); !v283 {
															if c < int32(8016) {
																v282 = libc.BoolInt32(c >= int32(8008) && c <= int32(8013))
															} else {
																v282 = libc.BoolInt32(c <= int32(8023))
															}
														}
														v280 = libc.BoolInt32(v283 || v282 != 0)
													}
												}
												v275 = libc.BoolInt32(v284 || v280 != 0)
											}
										}
										v264 = libc.BoolInt32(v285 || v275 != 0)
									}
								}
								v241 = libc.BoolInt32(v286 || v264 != 0)
							}
						}
						v194 = libc.BoolInt32(v287 || v241 != 0)
					}
					v193 = v194
				} else {
					if v380 = c <= int32(8025); !v380 {
						if c < int32(11720) {
							if c < int32(8458) {
								if c < int32(8178) {
									if c < int32(8126) {
										if c < int32(8031) {
											if c < int32(8029) {
												v293 = libc.BoolInt32(c == int32(8027))
											} else {
												v293 = libc.BoolInt32(c <= int32(8029))
											}
											v292 = v293
										} else {
											if v295 = c <= int32(8061); !v295 {
												if c < int32(8118) {
													v294 = libc.BoolInt32(c >= int32(8064) && c <= int32(8116))
												} else {
													v294 = libc.BoolInt32(c <= int32(8124))
												}
											}
											v292 = libc.BoolInt32(v295 || v294 != 0)
										}
										v291 = v292
									} else {
										if v300 = c <= int32(8126); !v300 {
											if c < int32(8144) {
												if c < int32(8134) {
													v297 = libc.BoolInt32(c >= int32(8130) && c <= int32(8132))
												} else {
													v297 = libc.BoolInt32(c <= int32(8140))
												}
												v296 = v297
											} else {
												if v299 = c <= int32(8147); !v299 {
													if c < int32(8160) {
														v298 = libc.BoolInt32(c >= int32(8150) && c <= int32(8155))
													} else {
														v298 = libc.BoolInt32(c <= int32(8172))
													}
												}
												v296 = libc.BoolInt32(v299 || v298 != 0)
											}
										}
										v291 = libc.BoolInt32(v300 || v296 != 0)
									}
									v290 = v291
								} else {
									if v311 = c <= int32(8180); !v311 {
										if c < int32(8336) {
											if c < int32(8276) {
												if c < int32(8255) {
													v303 = libc.BoolInt32(c >= int32(8182) && c <= int32(8188))
												} else {
													v303 = libc.BoolInt32(c <= int32(8256))
												}
												v302 = v303
											} else {
												if v305 = c <= int32(8276); !v305 {
													if c < int32(8319) {
														v304 = libc.BoolInt32(c == int32(8305))
													} else {
														v304 = libc.BoolInt32(c <= int32(8319))
													}
												}
												v302 = libc.BoolInt32(v305 || v304 != 0)
											}
											v301 = v302
										} else {
											if v310 = c <= int32(8348); !v310 {
												if c < int32(8421) {
													if c < int32(8417) {
														v307 = libc.BoolInt32(c >= int32(8400) && c <= int32(8412))
													} else {
														v307 = libc.BoolInt32(c <= int32(8417))
													}
													v306 = v307
												} else {
													if v309 = c <= int32(8432); !v309 {
														if c < int32(8455) {
															v308 = libc.BoolInt32(c == int32(8450))
														} else {
															v308 = libc.BoolInt32(c <= int32(8455))
														}
													}
													v306 = libc.BoolInt32(v309 || v308 != 0)
												}
											}
											v301 = libc.BoolInt32(v310 || v306 != 0)
										}
									}
									v290 = libc.BoolInt32(v311 || v301 != 0)
								}
								v289 = v290
							} else {
								if v334 = c <= int32(8467); !v334 {
									if c < int32(11499) {
										if c < int32(8490) {
											if c < int32(8484) {
												if c < int32(8472) {
													v315 = libc.BoolInt32(c == int32(8469))
												} else {
													v315 = libc.BoolInt32(c <= int32(8477))
												}
												v314 = v315
											} else {
												if v317 = c <= int32(8484); !v317 {
													if c < int32(8488) {
														v316 = libc.BoolInt32(c == int32(8486))
													} else {
														v316 = libc.BoolInt32(c <= int32(8488))
													}
												}
												v314 = libc.BoolInt32(v317 || v316 != 0)
											}
											v313 = v314
										} else {
											if v322 = c <= int32(8505); !v322 {
												if c < int32(8526) {
													if c < int32(8517) {
														v319 = libc.BoolInt32(c >= int32(8508) && c <= int32(8511))
													} else {
														v319 = libc.BoolInt32(c <= int32(8521))
													}
													v318 = v319
												} else {
													if v321 = c <= int32(8526); !v321 {
														if c < int32(11264) {
															v320 = libc.BoolInt32(c >= int32(8544) && c <= int32(8584))
														} else {
															v320 = libc.BoolInt32(c <= int32(11492))
														}
													}
													v318 = libc.BoolInt32(v321 || v320 != 0)
												}
											}
											v313 = libc.BoolInt32(v322 || v318 != 0)
										}
										v312 = v313
									} else {
										if v333 = c <= int32(11507); !v333 {
											if c < int32(11647) {
												if c < int32(11565) {
													if c < int32(11559) {
														v325 = libc.BoolInt32(c >= int32(11520) && c <= int32(11557))
													} else {
														v325 = libc.BoolInt32(c <= int32(11559))
													}
													v324 = v325
												} else {
													if v327 = c <= int32(11565); !v327 {
														if c < int32(11631) {
															v326 = libc.BoolInt32(c >= int32(11568) && c <= int32(11623))
														} else {
															v326 = libc.BoolInt32(c <= int32(11631))
														}
													}
													v324 = libc.BoolInt32(v327 || v326 != 0)
												}
												v323 = v324
											} else {
												if v332 = c <= int32(11670); !v332 {
													if c < int32(11696) {
														if c < int32(11688) {
															v329 = libc.BoolInt32(c >= int32(11680) && c <= int32(11686))
														} else {
															v329 = libc.BoolInt32(c <= int32(11694))
														}
														v328 = v329
													} else {
														if v331 = c <= int32(11702); !v331 {
															if c < int32(11712) {
																v330 = libc.BoolInt32(c >= int32(11704) && c <= int32(11710))
															} else {
																v330 = libc.BoolInt32(c <= int32(11718))
															}
														}
														v328 = libc.BoolInt32(v331 || v330 != 0)
													}
												}
												v323 = libc.BoolInt32(v332 || v328 != 0)
											}
										}
										v312 = libc.BoolInt32(v333 || v323 != 0)
									}
								}
								v289 = libc.BoolInt32(v334 || v312 != 0)
							}
							v288 = v289
						} else {
							if v379 = c <= int32(11726); !v379 {
								if c < int32(42623) {
									if c < int32(12540) {
										if c < int32(12337) {
											if c < int32(11744) {
												if c < int32(11736) {
													v339 = libc.BoolInt32(c >= int32(11728) && c <= int32(11734))
												} else {
													v339 = libc.BoolInt32(c <= int32(11742))
												}
												v338 = v339
											} else {
												if v341 = c <= int32(11775); !v341 {
													if c < int32(12321) {
														v340 = libc.BoolInt32(c >= int32(12293) && c <= int32(12295))
													} else {
														v340 = libc.BoolInt32(c <= int32(12335))
													}
												}
												v338 = libc.BoolInt32(v341 || v340 != 0)
											}
											v337 = v338
										} else {
											if v346 = c <= int32(12341); !v346 {
												if c < int32(12441) {
													if c < int32(12353) {
														v343 = libc.BoolInt32(c >= int32(12344) && c <= int32(12348))
													} else {
														v343 = libc.BoolInt32(c <= int32(12438))
													}
													v342 = v343
												} else {
													if v345 = c <= int32(12442); !v345 {
														if c < int32(12449) {
															v344 = libc.BoolInt32(c >= int32(12445) && c <= int32(12447))
														} else {
															v344 = libc.BoolInt32(c <= int32(12538))
														}
													}
													v342 = libc.BoolInt32(v345 || v344 != 0)
												}
											}
											v337 = libc.BoolInt32(v346 || v342 != 0)
										}
										v336 = v337
									} else {
										if v357 = c <= int32(12543); !v357 {
											if c < int32(19968) {
												if c < int32(12704) {
													if c < int32(12593) {
														v349 = libc.BoolInt32(c >= int32(12549) && c <= int32(12591))
													} else {
														v349 = libc.BoolInt32(c <= int32(12686))
													}
													v348 = v349
												} else {
													if v351 = c <= int32(12735); !v351 {
														if c < int32(13312) {
															v350 = libc.BoolInt32(c >= int32(12784) && c <= int32(12799))
														} else {
															v350 = libc.BoolInt32(c <= int32(19903))
														}
													}
													v348 = libc.BoolInt32(v351 || v350 != 0)
												}
												v347 = v348
											} else {
												if v356 = c <= int32(42124); !v356 {
													if c < int32(42512) {
														if c < int32(42240) {
															v353 = libc.BoolInt32(c >= int32(42192) && c <= int32(42237))
														} else {
															v353 = libc.BoolInt32(c <= int32(42508))
														}
														v352 = v353
													} else {
														if v355 = c <= int32(42539); !v355 {
															if c < int32(42612) {
																v354 = libc.BoolInt32(c >= int32(42560) && c <= int32(42607))
															} else {
																v354 = libc.BoolInt32(c <= int32(42621))
															}
														}
														v352 = libc.BoolInt32(v355 || v354 != 0)
													}
												}
												v347 = libc.BoolInt32(v356 || v352 != 0)
											}
										}
										v336 = libc.BoolInt32(v357 || v347 != 0)
									}
									v335 = v336
								} else {
									if v378 = c <= int32(42737); !v378 {
										if c < int32(43232) {
											if c < int32(42965) {
												if c < int32(42891) {
													if c < int32(42786) {
														v361 = libc.BoolInt32(c >= int32(42775) && c <= int32(42783))
													} else {
														v361 = libc.BoolInt32(c <= int32(42888))
													}
													v360 = v361
												} else {
													if v363 = c <= int32(42954); !v363 {
														if c < int32(42963) {
															v362 = libc.BoolInt32(c >= int32(42960) && c <= int32(42961))
														} else {
															v362 = libc.BoolInt32(c <= int32(42963))
														}
													}
													v360 = libc.BoolInt32(v363 || v362 != 0)
												}
												v359 = v360
											} else {
												if v368 = c <= int32(42969); !v368 {
													if c < int32(43072) {
														if c < int32(43052) {
															v365 = libc.BoolInt32(c >= int32(42994) && c <= int32(43047))
														} else {
															v365 = libc.BoolInt32(c <= int32(43052))
														}
														v364 = v365
													} else {
														if v367 = c <= int32(43123); !v367 {
															if c < int32(43216) {
																v366 = libc.BoolInt32(c >= int32(43136) && c <= int32(43205))
															} else {
																v366 = libc.BoolInt32(c <= int32(43225))
															}
														}
														v364 = libc.BoolInt32(v367 || v366 != 0)
													}
												}
												v359 = libc.BoolInt32(v368 || v364 != 0)
											}
											v358 = v359
										} else {
											if v377 = c <= int32(43255); !v377 {
												if c < int32(43471) {
													if c < int32(43312) {
														if c < int32(43261) {
															v371 = libc.BoolInt32(c == int32(43259))
														} else {
															v371 = libc.BoolInt32(c <= int32(43309))
														}
														v370 = v371
													} else {
														if v373 = c <= int32(43347); !v373 {
															if c < int32(43392) {
																v372 = libc.BoolInt32(c >= int32(43360) && c <= int32(43388))
															} else {
																v372 = libc.BoolInt32(c <= int32(43456))
															}
														}
														v370 = libc.BoolInt32(v373 || v372 != 0)
													}
													v369 = v370
												} else {
													if v376 = c <= int32(43481); !v376 {
														if c < int32(43584) {
															if c < int32(43520) {
																v375 = libc.BoolInt32(c >= int32(43488) && c <= int32(43518))
															} else {
																v375 = libc.BoolInt32(c <= int32(43574))
															}
															v374 = v375
														} else {
															v374 = libc.BoolInt32(c <= int32(43597) || c >= int32(43600) && c <= int32(43609))
														}
													}
													v369 = libc.BoolInt32(v376 || v374 != 0)
												}
											}
											v358 = libc.BoolInt32(v377 || v369 != 0)
										}
									}
									v335 = libc.BoolInt32(v378 || v358 != 0)
								}
							}
							v288 = libc.BoolInt32(v379 || v335 != 0)
						}
					}
					v193 = libc.BoolInt32(v380 || v288 != 0)
				}
			}
			v2 = libc.BoolInt32(v381 || v193 != 0)
		}
		v1 = v2
	} else {
		if v760 = c <= int32(43638); !v760 {
			if c < int32(71453) {
				if c < int32(67639) {
					if c < int32(65345) {
						if c < int32(64312) {
							if c < int32(43888) {
								if c < int32(43785) {
									if c < int32(43744) {
										if c < int32(43739) {
											v389 = libc.BoolInt32(c >= int32(43642) && c <= int32(43714))
										} else {
											v389 = libc.BoolInt32(c <= int32(43741))
										}
										v388 = v389
									} else {
										if v391 = c <= int32(43759); !v391 {
											if c < int32(43777) {
												v390 = libc.BoolInt32(c >= int32(43762) && c <= int32(43766))
											} else {
												v390 = libc.BoolInt32(c <= int32(43782))
											}
										}
										v388 = libc.BoolInt32(v391 || v390 != 0)
									}
									v387 = v388
								} else {
									if v396 = c <= int32(43790); !v396 {
										if c < int32(43816) {
											if c < int32(43808) {
												v393 = libc.BoolInt32(c >= int32(43793) && c <= int32(43798))
											} else {
												v393 = libc.BoolInt32(c <= int32(43814))
											}
											v392 = v393
										} else {
											if v395 = c <= int32(43822); !v395 {
												if c < int32(43868) {
													v394 = libc.BoolInt32(c >= int32(43824) && c <= int32(43866))
												} else {
													v394 = libc.BoolInt32(c <= int32(43881))
												}
											}
											v392 = libc.BoolInt32(v395 || v394 != 0)
										}
									}
									v387 = libc.BoolInt32(v396 || v392 != 0)
								}
								v386 = v387
							} else {
								if v407 = c <= int32(44010); !v407 {
									if c < int32(63744) {
										if c < int32(44032) {
											if c < int32(44016) {
												v399 = libc.BoolInt32(c >= int32(44012) && c <= int32(44013))
											} else {
												v399 = libc.BoolInt32(c <= int32(44025))
											}
											v398 = v399
										} else {
											if v401 = c <= int32(55203); !v401 {
												if c < int32(55243) {
													v400 = libc.BoolInt32(c >= int32(55216) && c <= int32(55238))
												} else {
													v400 = libc.BoolInt32(c <= int32(55291))
												}
											}
											v398 = libc.BoolInt32(v401 || v400 != 0)
										}
										v397 = v398
									} else {
										if v406 = c <= int32(64109); !v406 {
											if c < int32(64275) {
												if c < int32(64256) {
													v403 = libc.BoolInt32(c >= int32(64112) && c <= int32(64217))
												} else {
													v403 = libc.BoolInt32(c <= int32(64262))
												}
												v402 = v403
											} else {
												if v405 = c <= int32(64279); !v405 {
													if c < int32(64298) {
														v404 = libc.BoolInt32(c >= int32(64285) && c <= int32(64296))
													} else {
														v404 = libc.BoolInt32(c <= int32(64310))
													}
												}
												v402 = libc.BoolInt32(v405 || v404 != 0)
											}
										}
										v397 = libc.BoolInt32(v406 || v402 != 0)
									}
								}
								v386 = libc.BoolInt32(v407 || v397 != 0)
							}
							v385 = v386
						} else {
							if v430 = c <= int32(64316); !v430 {
								if c < int32(65075) {
									if c < int32(64612) {
										if c < int32(64323) {
											if c < int32(64320) {
												v411 = libc.BoolInt32(c == int32(64318))
											} else {
												v411 = libc.BoolInt32(c <= int32(64321))
											}
											v410 = v411
										} else {
											if v413 = c <= int32(64324); !v413 {
												if c < int32(64467) {
													v412 = libc.BoolInt32(c >= int32(64326) && c <= int32(64433))
												} else {
													v412 = libc.BoolInt32(c <= int32(64605))
												}
											}
											v410 = libc.BoolInt32(v413 || v412 != 0)
										}
										v409 = v410
									} else {
										if v418 = c <= int32(64829); !v418 {
											if c < int32(65008) {
												if c < int32(64914) {
													v415 = libc.BoolInt32(c >= int32(64848) && c <= int32(64911))
												} else {
													v415 = libc.BoolInt32(c <= int32(64967))
												}
												v414 = v415
											} else {
												if v417 = c <= int32(65017); !v417 {
													if c < int32(65056) {
														v416 = libc.BoolInt32(c >= int32(65024) && c <= int32(65039))
													} else {
														v416 = libc.BoolInt32(c <= int32(65071))
													}
												}
												v414 = libc.BoolInt32(v417 || v416 != 0)
											}
										}
										v409 = libc.BoolInt32(v418 || v414 != 0)
									}
									v408 = v409
								} else {
									if v429 = c <= int32(65076); !v429 {
										if c < int32(65147) {
											if c < int32(65139) {
												if c < int32(65137) {
													v421 = libc.BoolInt32(c >= int32(65101) && c <= int32(65103))
												} else {
													v421 = libc.BoolInt32(c <= int32(65137))
												}
												v420 = v421
											} else {
												if v423 = c <= int32(65139); !v423 {
													if c < int32(65145) {
														v422 = libc.BoolInt32(c == int32(65143))
													} else {
														v422 = libc.BoolInt32(c <= int32(65145))
													}
												}
												v420 = libc.BoolInt32(v423 || v422 != 0)
											}
											v419 = v420
										} else {
											if v428 = c <= int32(65147); !v428 {
												if c < int32(65296) {
													if c < int32(65151) {
														v425 = libc.BoolInt32(c == int32(65149))
													} else {
														v425 = libc.BoolInt32(c <= int32(65276))
													}
													v424 = v425
												} else {
													if v427 = c <= int32(65305); !v427 {
														if c < int32(65343) {
															v426 = libc.BoolInt32(c >= int32(65313) && c <= int32(65338))
														} else {
															v426 = libc.BoolInt32(c <= int32(65343))
														}
													}
													v424 = libc.BoolInt32(v427 || v426 != 0)
												}
											}
											v419 = libc.BoolInt32(v428 || v424 != 0)
										}
									}
									v408 = libc.BoolInt32(v429 || v419 != 0)
								}
							}
							v385 = libc.BoolInt32(v430 || v408 != 0)
						}
						v384 = v385
					} else {
						if v477 = c <= int32(65370); !v477 {
							if c < int32(66513) {
								if c < int32(65664) {
									if c < int32(65536) {
										if c < int32(65482) {
											if c < int32(65474) {
												v435 = libc.BoolInt32(c >= int32(65382) && c <= int32(65470))
											} else {
												v435 = libc.BoolInt32(c <= int32(65479))
											}
											v434 = v435
										} else {
											if v437 = c <= int32(65487); !v437 {
												if c < int32(65498) {
													v436 = libc.BoolInt32(c >= int32(65490) && c <= int32(65495))
												} else {
													v436 = libc.BoolInt32(c <= int32(65500))
												}
											}
											v434 = libc.BoolInt32(v437 || v436 != 0)
										}
										v433 = v434
									} else {
										if v442 = c <= int32(65547); !v442 {
											if c < int32(65596) {
												if c < int32(65576) {
													v439 = libc.BoolInt32(c >= int32(65549) && c <= int32(65574))
												} else {
													v439 = libc.BoolInt32(c <= int32(65594))
												}
												v438 = v439
											} else {
												if v441 = c <= int32(65597); !v441 {
													if c < int32(65616) {
														v440 = libc.BoolInt32(c >= int32(65599) && c <= int32(65613))
													} else {
														v440 = libc.BoolInt32(c <= int32(65629))
													}
												}
												v438 = libc.BoolInt32(v441 || v440 != 0)
											}
										}
										v433 = libc.BoolInt32(v442 || v438 != 0)
									}
									v432 = v433
								} else {
									if v453 = c <= int32(65786); !v453 {
										if c < int32(66304) {
											if c < int32(66176) {
												if c < int32(66045) {
													v445 = libc.BoolInt32(c >= int32(65856) && c <= int32(65908))
												} else {
													v445 = libc.BoolInt32(c <= int32(66045))
												}
												v444 = v445
											} else {
												if v447 = c <= int32(66204); !v447 {
													if c < int32(66272) {
														v446 = libc.BoolInt32(c >= int32(66208) && c <= int32(66256))
													} else {
														v446 = libc.BoolInt32(c <= int32(66272))
													}
												}
												v444 = libc.BoolInt32(v447 || v446 != 0)
											}
											v443 = v444
										} else {
											if v452 = c <= int32(66335); !v452 {
												if c < int32(66432) {
													if c < int32(66384) {
														v449 = libc.BoolInt32(c >= int32(66349) && c <= int32(66378))
													} else {
														v449 = libc.BoolInt32(c <= int32(66426))
													}
													v448 = v449
												} else {
													if v451 = c <= int32(66461); !v451 {
														if c < int32(66504) {
															v450 = libc.BoolInt32(c >= int32(66464) && c <= int32(66499))
														} else {
															v450 = libc.BoolInt32(c <= int32(66511))
														}
													}
													v448 = libc.BoolInt32(v451 || v450 != 0)
												}
											}
											v443 = libc.BoolInt32(v452 || v448 != 0)
										}
									}
									v432 = libc.BoolInt32(v453 || v443 != 0)
								}
								v431 = v432
							} else {
								if v476 = c <= int32(66517); !v476 {
									if c < int32(66979) {
										if c < int32(66864) {
											if c < int32(66736) {
												if c < int32(66720) {
													v457 = libc.BoolInt32(c >= int32(66560) && c <= int32(66717))
												} else {
													v457 = libc.BoolInt32(c <= int32(66729))
												}
												v456 = v457
											} else {
												if v459 = c <= int32(66771); !v459 {
													if c < int32(66816) {
														v458 = libc.BoolInt32(c >= int32(66776) && c <= int32(66811))
													} else {
														v458 = libc.BoolInt32(c <= int32(66855))
													}
												}
												v456 = libc.BoolInt32(v459 || v458 != 0)
											}
											v455 = v456
										} else {
											if v464 = c <= int32(66915); !v464 {
												if c < int32(66956) {
													if c < int32(66940) {
														v461 = libc.BoolInt32(c >= int32(66928) && c <= int32(66938))
													} else {
														v461 = libc.BoolInt32(c <= int32(66954))
													}
													v460 = v461
												} else {
													if v463 = c <= int32(66962); !v463 {
														if c < int32(66967) {
															v462 = libc.BoolInt32(c >= int32(66964) && c <= int32(66965))
														} else {
															v462 = libc.BoolInt32(c <= int32(66977))
														}
													}
													v460 = libc.BoolInt32(v463 || v462 != 0)
												}
											}
											v455 = libc.BoolInt32(v464 || v460 != 0)
										}
										v454 = v455
									} else {
										if v475 = c <= int32(66993); !v475 {
											if c < int32(67456) {
												if c < int32(67072) {
													if c < int32(67003) {
														v467 = libc.BoolInt32(c >= int32(66995) && c <= int32(67001))
													} else {
														v467 = libc.BoolInt32(c <= int32(67004))
													}
													v466 = v467
												} else {
													if v469 = c <= int32(67382); !v469 {
														if c < int32(67424) {
															v468 = libc.BoolInt32(c >= int32(67392) && c <= int32(67413))
														} else {
															v468 = libc.BoolInt32(c <= int32(67431))
														}
													}
													v466 = libc.BoolInt32(v469 || v468 != 0)
												}
												v465 = v466
											} else {
												if v474 = c <= int32(67461); !v474 {
													if c < int32(67584) {
														if c < int32(67506) {
															v471 = libc.BoolInt32(c >= int32(67463) && c <= int32(67504))
														} else {
															v471 = libc.BoolInt32(c <= int32(67514))
														}
														v470 = v471
													} else {
														if v473 = c <= int32(67589); !v473 {
															if c < int32(67594) {
																v472 = libc.BoolInt32(c == int32(67592))
															} else {
																v472 = libc.BoolInt32(c <= int32(67637))
															}
														}
														v470 = libc.BoolInt32(v473 || v472 != 0)
													}
												}
												v465 = libc.BoolInt32(v474 || v470 != 0)
											}
										}
										v454 = libc.BoolInt32(v475 || v465 != 0)
									}
								}
								v431 = libc.BoolInt32(v476 || v454 != 0)
							}
						}
						v384 = libc.BoolInt32(v477 || v431 != 0)
					}
					v383 = v384
				} else {
					if v570 = c <= int32(67640); !v570 {
						if c < int32(69956) {
							if c < int32(68448) {
								if c < int32(68101) {
									if c < int32(67828) {
										if c < int32(67680) {
											if c < int32(67647) {
												v483 = libc.BoolInt32(c == int32(67644))
											} else {
												v483 = libc.BoolInt32(c <= int32(67669))
											}
											v482 = v483
										} else {
											if v485 = c <= int32(67702); !v485 {
												if c < int32(67808) {
													v484 = libc.BoolInt32(c >= int32(67712) && c <= int32(67742))
												} else {
													v484 = libc.BoolInt32(c <= int32(67826))
												}
											}
											v482 = libc.BoolInt32(v485 || v484 != 0)
										}
										v481 = v482
									} else {
										if v490 = c <= int32(67829); !v490 {
											if c < int32(67968) {
												if c < int32(67872) {
													v487 = libc.BoolInt32(c >= int32(67840) && c <= int32(67861))
												} else {
													v487 = libc.BoolInt32(c <= int32(67897))
												}
												v486 = v487
											} else {
												if v489 = c <= int32(68023); !v489 {
													if c < int32(68096) {
														v488 = libc.BoolInt32(c >= int32(68030) && c <= int32(68031))
													} else {
														v488 = libc.BoolInt32(c <= int32(68099))
													}
												}
												v486 = libc.BoolInt32(v489 || v488 != 0)
											}
										}
										v481 = libc.BoolInt32(v490 || v486 != 0)
									}
									v480 = v481
								} else {
									if v501 = c <= int32(68102); !v501 {
										if c < int32(68192) {
											if c < int32(68121) {
												if c < int32(68117) {
													v493 = libc.BoolInt32(c >= int32(68108) && c <= int32(68115))
												} else {
													v493 = libc.BoolInt32(c <= int32(68119))
												}
												v492 = v493
											} else {
												if v495 = c <= int32(68149); !v495 {
													if c < int32(68159) {
														v494 = libc.BoolInt32(c >= int32(68152) && c <= int32(68154))
													} else {
														v494 = libc.BoolInt32(c <= int32(68159))
													}
												}
												v492 = libc.BoolInt32(v495 || v494 != 0)
											}
											v491 = v492
										} else {
											if v500 = c <= int32(68220); !v500 {
												if c < int32(68297) {
													if c < int32(68288) {
														v497 = libc.BoolInt32(c >= int32(68224) && c <= int32(68252))
													} else {
														v497 = libc.BoolInt32(c <= int32(68295))
													}
													v496 = v497
												} else {
													if v499 = c <= int32(68326); !v499 {
														if c < int32(68416) {
															v498 = libc.BoolInt32(c >= int32(68352) && c <= int32(68405))
														} else {
															v498 = libc.BoolInt32(c <= int32(68437))
														}
													}
													v496 = libc.BoolInt32(v499 || v498 != 0)
												}
											}
											v491 = libc.BoolInt32(v500 || v496 != 0)
										}
									}
									v480 = libc.BoolInt32(v501 || v491 != 0)
								}
								v479 = v480
							} else {
								if v524 = c <= int32(68466); !v524 {
									if c < int32(69424) {
										if c < int32(68912) {
											if c < int32(68736) {
												if c < int32(68608) {
													v505 = libc.BoolInt32(c >= int32(68480) && c <= int32(68497))
												} else {
													v505 = libc.BoolInt32(c <= int32(68680))
												}
												v504 = v505
											} else {
												if v507 = c <= int32(68786); !v507 {
													if c < int32(68864) {
														v506 = libc.BoolInt32(c >= int32(68800) && c <= int32(68850))
													} else {
														v506 = libc.BoolInt32(c <= int32(68903))
													}
												}
												v504 = libc.BoolInt32(v507 || v506 != 0)
											}
											v503 = v504
										} else {
											if v512 = c <= int32(68921); !v512 {
												if c < int32(69296) {
													if c < int32(69291) {
														v509 = libc.BoolInt32(c >= int32(69248) && c <= int32(69289))
													} else {
														v509 = libc.BoolInt32(c <= int32(69292))
													}
													v508 = v509
												} else {
													if v511 = c <= int32(69297); !v511 {
														if c < int32(69415) {
															v510 = libc.BoolInt32(c >= int32(69376) && c <= int32(69404))
														} else {
															v510 = libc.BoolInt32(c <= int32(69415))
														}
													}
													v508 = libc.BoolInt32(v511 || v510 != 0)
												}
											}
											v503 = libc.BoolInt32(v512 || v508 != 0)
										}
										v502 = v503
									} else {
										if v523 = c <= int32(69456); !v523 {
											if c < int32(69759) {
												if c < int32(69600) {
													if c < int32(69552) {
														v515 = libc.BoolInt32(c >= int32(69488) && c <= int32(69509))
													} else {
														v515 = libc.BoolInt32(c <= int32(69572))
													}
													v514 = v515
												} else {
													if v517 = c <= int32(69622); !v517 {
														if c < int32(69734) {
															v516 = libc.BoolInt32(c >= int32(69632) && c <= int32(69702))
														} else {
															v516 = libc.BoolInt32(c <= int32(69749))
														}
													}
													v514 = libc.BoolInt32(v517 || v516 != 0)
												}
												v513 = v514
											} else {
												if v522 = c <= int32(69818); !v522 {
													if c < int32(69872) {
														if c < int32(69840) {
															v519 = libc.BoolInt32(c == int32(69826))
														} else {
															v519 = libc.BoolInt32(c <= int32(69864))
														}
														v518 = v519
													} else {
														if v521 = c <= int32(69881); !v521 {
															if c < int32(69942) {
																v520 = libc.BoolInt32(c >= int32(69888) && c <= int32(69940))
															} else {
																v520 = libc.BoolInt32(c <= int32(69951))
															}
														}
														v518 = libc.BoolInt32(v521 || v520 != 0)
													}
												}
												v513 = libc.BoolInt32(v522 || v518 != 0)
											}
										}
										v502 = libc.BoolInt32(v523 || v513 != 0)
									}
								}
								v479 = libc.BoolInt32(v524 || v502 != 0)
							}
							v478 = v479
						} else {
							if v569 = c <= int32(69959); !v569 {
								if c < int32(70459) {
									if c < int32(70282) {
										if c < int32(70108) {
											if c < int32(70016) {
												if c < int32(70006) {
													v529 = libc.BoolInt32(c >= int32(69968) && c <= int32(70003))
												} else {
													v529 = libc.BoolInt32(c <= int32(70006))
												}
												v528 = v529
											} else {
												if v531 = c <= int32(70084); !v531 {
													if c < int32(70094) {
														v530 = libc.BoolInt32(c >= int32(70089) && c <= int32(70092))
													} else {
														v530 = libc.BoolInt32(c <= int32(70106))
													}
												}
												v528 = libc.BoolInt32(v531 || v530 != 0)
											}
											v527 = v528
										} else {
											if v536 = c <= int32(70108); !v536 {
												if c < int32(70206) {
													if c < int32(70163) {
														v533 = libc.BoolInt32(c >= int32(70144) && c <= int32(70161))
													} else {
														v533 = libc.BoolInt32(c <= int32(70199))
													}
													v532 = v533
												} else {
													if v535 = c <= int32(70206); !v535 {
														if c < int32(70280) {
															v534 = libc.BoolInt32(c >= int32(70272) && c <= int32(70278))
														} else {
															v534 = libc.BoolInt32(c <= int32(70280))
														}
													}
													v532 = libc.BoolInt32(v535 || v534 != 0)
												}
											}
											v527 = libc.BoolInt32(v536 || v532 != 0)
										}
										v526 = v527
									} else {
										if v547 = c <= int32(70285); !v547 {
											if c < int32(70405) {
												if c < int32(70320) {
													if c < int32(70303) {
														v539 = libc.BoolInt32(c >= int32(70287) && c <= int32(70301))
													} else {
														v539 = libc.BoolInt32(c <= int32(70312))
													}
													v538 = v539
												} else {
													if v541 = c <= int32(70378); !v541 {
														if c < int32(70400) {
															v540 = libc.BoolInt32(c >= int32(70384) && c <= int32(70393))
														} else {
															v540 = libc.BoolInt32(c <= int32(70403))
														}
													}
													v538 = libc.BoolInt32(v541 || v540 != 0)
												}
												v537 = v538
											} else {
												if v546 = c <= int32(70412); !v546 {
													if c < int32(70442) {
														if c < int32(70419) {
															v543 = libc.BoolInt32(c >= int32(70415) && c <= int32(70416))
														} else {
															v543 = libc.BoolInt32(c <= int32(70440))
														}
														v542 = v543
													} else {
														if v545 = c <= int32(70448); !v545 {
															if c < int32(70453) {
																v544 = libc.BoolInt32(c >= int32(70450) && c <= int32(70451))
															} else {
																v544 = libc.BoolInt32(c <= int32(70457))
															}
														}
														v542 = libc.BoolInt32(v545 || v544 != 0)
													}
												}
												v537 = libc.BoolInt32(v546 || v542 != 0)
											}
										}
										v526 = libc.BoolInt32(v547 || v537 != 0)
									}
									v525 = v526
								} else {
									if v568 = c <= int32(70468); !v568 {
										if c < int32(70855) {
											if c < int32(70502) {
												if c < int32(70480) {
													if c < int32(70475) {
														v551 = libc.BoolInt32(c >= int32(70471) && c <= int32(70472))
													} else {
														v551 = libc.BoolInt32(c <= int32(70477))
													}
													v550 = v551
												} else {
													if v553 = c <= int32(70480); !v553 {
														if c < int32(70493) {
															v552 = libc.BoolInt32(c == int32(70487))
														} else {
															v552 = libc.BoolInt32(c <= int32(70499))
														}
													}
													v550 = libc.BoolInt32(v553 || v552 != 0)
												}
												v549 = v550
											} else {
												if v558 = c <= int32(70508); !v558 {
													if c < int32(70736) {
														if c < int32(70656) {
															v555 = libc.BoolInt32(c >= int32(70512) && c <= int32(70516))
														} else {
															v555 = libc.BoolInt32(c <= int32(70730))
														}
														v554 = v555
													} else {
														if v557 = c <= int32(70745); !v557 {
															if c < int32(70784) {
																v556 = libc.BoolInt32(c >= int32(70750) && c <= int32(70753))
															} else {
																v556 = libc.BoolInt32(c <= int32(70853))
															}
														}
														v554 = libc.BoolInt32(v557 || v556 != 0)
													}
												}
												v549 = libc.BoolInt32(v558 || v554 != 0)
											}
											v548 = v549
										} else {
											if v567 = c <= int32(70855); !v567 {
												if c < int32(71236) {
													if c < int32(71096) {
														if c < int32(71040) {
															v561 = libc.BoolInt32(c >= int32(70864) && c <= int32(70873))
														} else {
															v561 = libc.BoolInt32(c <= int32(71093))
														}
														v560 = v561
													} else {
														if v563 = c <= int32(71104); !v563 {
															if c < int32(71168) {
																v562 = libc.BoolInt32(c >= int32(71128) && c <= int32(71133))
															} else {
																v562 = libc.BoolInt32(c <= int32(71232))
															}
														}
														v560 = libc.BoolInt32(v563 || v562 != 0)
													}
													v559 = v560
												} else {
													if v566 = c <= int32(71236); !v566 {
														if c < int32(71360) {
															if c < int32(71296) {
																v565 = libc.BoolInt32(c >= int32(71248) && c <= int32(71257))
															} else {
																v565 = libc.BoolInt32(c <= int32(71352))
															}
															v564 = v565
														} else {
															v564 = libc.BoolInt32(c <= int32(71369) || c >= int32(71424) && c <= int32(71450))
														}
													}
													v559 = libc.BoolInt32(v566 || v564 != 0)
												}
											}
											v548 = libc.BoolInt32(v567 || v559 != 0)
										}
									}
									v525 = libc.BoolInt32(v568 || v548 != 0)
								}
							}
							v478 = libc.BoolInt32(v569 || v525 != 0)
						}
					}
					v383 = libc.BoolInt32(v570 || v478 != 0)
				}
				v382 = v383
			} else {
				if v759 = c <= int32(71467); !v759 {
					if c < int32(119973) {
						if c < int32(77824) {
							if c < int32(72760) {
								if c < int32(72016) {
									if c < int32(71945) {
										if c < int32(71680) {
											if c < int32(71488) {
												v577 = libc.BoolInt32(c >= int32(71472) && c <= int32(71481))
											} else {
												v577 = libc.BoolInt32(c <= int32(71494))
											}
											v576 = v577
										} else {
											if v579 = c <= int32(71738); !v579 {
												if c < int32(71935) {
													v578 = libc.BoolInt32(c >= int32(71840) && c <= int32(71913))
												} else {
													v578 = libc.BoolInt32(c <= int32(71942))
												}
											}
											v576 = libc.BoolInt32(v579 || v578 != 0)
										}
										v575 = v576
									} else {
										if v584 = c <= int32(71945); !v584 {
											if c < int32(71960) {
												if c < int32(71957) {
													v581 = libc.BoolInt32(c >= int32(71948) && c <= int32(71955))
												} else {
													v581 = libc.BoolInt32(c <= int32(71958))
												}
												v580 = v581
											} else {
												if v583 = c <= int32(71989); !v583 {
													if c < int32(71995) {
														v582 = libc.BoolInt32(c >= int32(71991) && c <= int32(71992))
													} else {
														v582 = libc.BoolInt32(c <= int32(72003))
													}
												}
												v580 = libc.BoolInt32(v583 || v582 != 0)
											}
										}
										v575 = libc.BoolInt32(v584 || v580 != 0)
									}
									v574 = v575
								} else {
									if v595 = c <= int32(72025); !v595 {
										if c < int32(72263) {
											if c < int32(72154) {
												if c < int32(72106) {
													v587 = libc.BoolInt32(c >= int32(72096) && c <= int32(72103))
												} else {
													v587 = libc.BoolInt32(c <= int32(72151))
												}
												v586 = v587
											} else {
												if v589 = c <= int32(72161); !v589 {
													if c < int32(72192) {
														v588 = libc.BoolInt32(c >= int32(72163) && c <= int32(72164))
													} else {
														v588 = libc.BoolInt32(c <= int32(72254))
													}
												}
												v586 = libc.BoolInt32(v589 || v588 != 0)
											}
											v585 = v586
										} else {
											if v594 = c <= int32(72263); !v594 {
												if c < int32(72368) {
													if c < int32(72349) {
														v591 = libc.BoolInt32(c >= int32(72272) && c <= int32(72345))
													} else {
														v591 = libc.BoolInt32(c <= int32(72349))
													}
													v590 = v591
												} else {
													if v593 = c <= int32(72440); !v593 {
														if c < int32(72714) {
															v592 = libc.BoolInt32(c >= int32(72704) && c <= int32(72712))
														} else {
															v592 = libc.BoolInt32(c <= int32(72758))
														}
													}
													v590 = libc.BoolInt32(v593 || v592 != 0)
												}
											}
											v585 = libc.BoolInt32(v594 || v590 != 0)
										}
									}
									v574 = libc.BoolInt32(v595 || v585 != 0)
								}
								v573 = v574
							} else {
								if v618 = c <= int32(72768); !v618 {
									if c < int32(73056) {
										if c < int32(72968) {
											if c < int32(72850) {
												if c < int32(72818) {
													v599 = libc.BoolInt32(c >= int32(72784) && c <= int32(72793))
												} else {
													v599 = libc.BoolInt32(c <= int32(72847))
												}
												v598 = v599
											} else {
												if v601 = c <= int32(72871); !v601 {
													if c < int32(72960) {
														v600 = libc.BoolInt32(c >= int32(72873) && c <= int32(72886))
													} else {
														v600 = libc.BoolInt32(c <= int32(72966))
													}
												}
												v598 = libc.BoolInt32(v601 || v600 != 0)
											}
											v597 = v598
										} else {
											if v606 = c <= int32(72969); !v606 {
												if c < int32(73020) {
													if c < int32(73018) {
														v603 = libc.BoolInt32(c >= int32(72971) && c <= int32(73014))
													} else {
														v603 = libc.BoolInt32(c <= int32(73018))
													}
													v602 = v603
												} else {
													if v605 = c <= int32(73021); !v605 {
														if c < int32(73040) {
															v604 = libc.BoolInt32(c >= int32(73023) && c <= int32(73031))
														} else {
															v604 = libc.BoolInt32(c <= int32(73049))
														}
													}
													v602 = libc.BoolInt32(v605 || v604 != 0)
												}
											}
											v597 = libc.BoolInt32(v606 || v602 != 0)
										}
										v596 = v597
									} else {
										if v617 = c <= int32(73061); !v617 {
											if c < int32(73440) {
												if c < int32(73104) {
													if c < int32(73066) {
														v609 = libc.BoolInt32(c >= int32(73063) && c <= int32(73064))
													} else {
														v609 = libc.BoolInt32(c <= int32(73102))
													}
													v608 = v609
												} else {
													if v611 = c <= int32(73105); !v611 {
														if c < int32(73120) {
															v610 = libc.BoolInt32(c >= int32(73107) && c <= int32(73112))
														} else {
															v610 = libc.BoolInt32(c <= int32(73129))
														}
													}
													v608 = libc.BoolInt32(v611 || v610 != 0)
												}
												v607 = v608
											} else {
												if v616 = c <= int32(73462); !v616 {
													if c < int32(74752) {
														if c < int32(73728) {
															v613 = libc.BoolInt32(c == int32(73648))
														} else {
															v613 = libc.BoolInt32(c <= int32(74649))
														}
														v612 = v613
													} else {
														if v615 = c <= int32(74862); !v615 {
															if c < int32(77712) {
																v614 = libc.BoolInt32(c >= int32(74880) && c <= int32(75075))
															} else {
																v614 = libc.BoolInt32(c <= int32(77808))
															}
														}
														v612 = libc.BoolInt32(v615 || v614 != 0)
													}
												}
												v607 = libc.BoolInt32(v616 || v612 != 0)
											}
										}
										v596 = libc.BoolInt32(v617 || v607 != 0)
									}
								}
								v573 = libc.BoolInt32(v618 || v596 != 0)
							}
							v572 = v573
						} else {
							if v665 = c <= int32(78894); !v665 {
								if c < int32(110576) {
									if c < int32(93027) {
										if c < int32(92864) {
											if c < int32(92736) {
												if c < int32(92160) {
													v623 = libc.BoolInt32(c >= int32(82944) && c <= int32(83526))
												} else {
													v623 = libc.BoolInt32(c <= int32(92728))
												}
												v622 = v623
											} else {
												if v625 = c <= int32(92766); !v625 {
													if c < int32(92784) {
														v624 = libc.BoolInt32(c >= int32(92768) && c <= int32(92777))
													} else {
														v624 = libc.BoolInt32(c <= int32(92862))
													}
												}
												v622 = libc.BoolInt32(v625 || v624 != 0)
											}
											v621 = v622
										} else {
											if v630 = c <= int32(92873); !v630 {
												if c < int32(92928) {
													if c < int32(92912) {
														v627 = libc.BoolInt32(c >= int32(92880) && c <= int32(92909))
													} else {
														v627 = libc.BoolInt32(c <= int32(92916))
													}
													v626 = v627
												} else {
													if v629 = c <= int32(92982); !v629 {
														if c < int32(93008) {
															v628 = libc.BoolInt32(c >= int32(92992) && c <= int32(92995))
														} else {
															v628 = libc.BoolInt32(c <= int32(93017))
														}
													}
													v626 = libc.BoolInt32(v629 || v628 != 0)
												}
											}
											v621 = libc.BoolInt32(v630 || v626 != 0)
										}
										v620 = v621
									} else {
										if v641 = c <= int32(93047); !v641 {
											if c < int32(94176) {
												if c < int32(93952) {
													if c < int32(93760) {
														v633 = libc.BoolInt32(c >= int32(93053) && c <= int32(93071))
													} else {
														v633 = libc.BoolInt32(c <= int32(93823))
													}
													v632 = v633
												} else {
													if v635 = c <= int32(94026); !v635 {
														if c < int32(94095) {
															v634 = libc.BoolInt32(c >= int32(94031) && c <= int32(94087))
														} else {
															v634 = libc.BoolInt32(c <= int32(94111))
														}
													}
													v632 = libc.BoolInt32(v635 || v634 != 0)
												}
												v631 = v632
											} else {
												if v640 = c <= int32(94177); !v640 {
													if c < int32(94208) {
														if c < int32(94192) {
															v637 = libc.BoolInt32(c >= int32(94179) && c <= int32(94180))
														} else {
															v637 = libc.BoolInt32(c <= int32(94193))
														}
														v636 = v637
													} else {
														if v639 = c <= int32(100343); !v639 {
															if c < int32(101632) {
																v638 = libc.BoolInt32(c >= int32(100352) && c <= int32(101589))
															} else {
																v638 = libc.BoolInt32(c <= int32(101640))
															}
														}
														v636 = libc.BoolInt32(v639 || v638 != 0)
													}
												}
												v631 = libc.BoolInt32(v640 || v636 != 0)
											}
										}
										v620 = libc.BoolInt32(v641 || v631 != 0)
									}
									v619 = v620
								} else {
									if v664 = c <= int32(110579); !v664 {
										if c < int32(118528) {
											if c < int32(110960) {
												if c < int32(110592) {
													if c < int32(110589) {
														v645 = libc.BoolInt32(c >= int32(110581) && c <= int32(110587))
													} else {
														v645 = libc.BoolInt32(c <= int32(110590))
													}
													v644 = v645
												} else {
													if v647 = c <= int32(110882); !v647 {
														if c < int32(110948) {
															v646 = libc.BoolInt32(c >= int32(110928) && c <= int32(110930))
														} else {
															v646 = libc.BoolInt32(c <= int32(110951))
														}
													}
													v644 = libc.BoolInt32(v647 || v646 != 0)
												}
												v643 = v644
											} else {
												if v652 = c <= int32(111355); !v652 {
													if c < int32(113792) {
														if c < int32(113776) {
															v649 = libc.BoolInt32(c >= int32(113664) && c <= int32(113770))
														} else {
															v649 = libc.BoolInt32(c <= int32(113788))
														}
														v648 = v649
													} else {
														if v651 = c <= int32(113800); !v651 {
															if c < int32(113821) {
																v650 = libc.BoolInt32(c >= int32(113808) && c <= int32(113817))
															} else {
																v650 = libc.BoolInt32(c <= int32(113822))
															}
														}
														v648 = libc.BoolInt32(v651 || v650 != 0)
													}
												}
												v643 = libc.BoolInt32(v652 || v648 != 0)
											}
											v642 = v643
										} else {
											if v663 = c <= int32(118573); !v663 {
												if c < int32(119210) {
													if c < int32(119149) {
														if c < int32(119141) {
															v655 = libc.BoolInt32(c >= int32(118576) && c <= int32(118598))
														} else {
															v655 = libc.BoolInt32(c <= int32(119145))
														}
														v654 = v655
													} else {
														if v657 = c <= int32(119154); !v657 {
															if c < int32(119173) {
																v656 = libc.BoolInt32(c >= int32(119163) && c <= int32(119170))
															} else {
																v656 = libc.BoolInt32(c <= int32(119179))
															}
														}
														v654 = libc.BoolInt32(v657 || v656 != 0)
													}
													v653 = v654
												} else {
													if v662 = c <= int32(119213); !v662 {
														if c < int32(119894) {
															if c < int32(119808) {
																v659 = libc.BoolInt32(c >= int32(119362) && c <= int32(119364))
															} else {
																v659 = libc.BoolInt32(c <= int32(119892))
															}
															v658 = v659
														} else {
															if v661 = c <= int32(119964); !v661 {
																if c < int32(119970) {
																	v660 = libc.BoolInt32(c >= int32(119966) && c <= int32(119967))
																} else {
																	v660 = libc.BoolInt32(c <= int32(119970))
																}
															}
															v658 = libc.BoolInt32(v661 || v660 != 0)
														}
													}
													v653 = libc.BoolInt32(v662 || v658 != 0)
												}
											}
											v642 = libc.BoolInt32(v663 || v653 != 0)
										}
									}
									v619 = libc.BoolInt32(v664 || v642 != 0)
								}
							}
							v572 = libc.BoolInt32(v665 || v619 != 0)
						}
						v571 = v572
					} else {
						if v758 = c <= int32(119974); !v758 {
							if c < int32(124912) {
								if c < int32(120746) {
									if c < int32(120134) {
										if c < int32(120071) {
											if c < int32(119995) {
												if c < int32(119982) {
													v671 = libc.BoolInt32(c >= int32(119977) && c <= int32(119980))
												} else {
													v671 = libc.BoolInt32(c <= int32(119993))
												}
												v670 = v671
											} else {
												if v673 = c <= int32(119995); !v673 {
													if c < int32(120005) {
														v672 = libc.BoolInt32(c >= int32(119997) && c <= int32(120003))
													} else {
														v672 = libc.BoolInt32(c <= int32(120069))
													}
												}
												v670 = libc.BoolInt32(v673 || v672 != 0)
											}
											v669 = v670
										} else {
											if v678 = c <= int32(120074); !v678 {
												if c < int32(120094) {
													if c < int32(120086) {
														v675 = libc.BoolInt32(c >= int32(120077) && c <= int32(120084))
													} else {
														v675 = libc.BoolInt32(c <= int32(120092))
													}
													v674 = v675
												} else {
													if v677 = c <= int32(120121); !v677 {
														if c < int32(120128) {
															v676 = libc.BoolInt32(c >= int32(120123) && c <= int32(120126))
														} else {
															v676 = libc.BoolInt32(c <= int32(120132))
														}
													}
													v674 = libc.BoolInt32(v677 || v676 != 0)
												}
											}
											v669 = libc.BoolInt32(v678 || v674 != 0)
										}
										v668 = v669
									} else {
										if v689 = c <= int32(120134); !v689 {
											if c < int32(120572) {
												if c < int32(120488) {
													if c < int32(120146) {
														v681 = libc.BoolInt32(c >= int32(120138) && c <= int32(120144))
													} else {
														v681 = libc.BoolInt32(c <= int32(120485))
													}
													v680 = v681
												} else {
													if v683 = c <= int32(120512); !v683 {
														if c < int32(120540) {
															v682 = libc.BoolInt32(c >= int32(120514) && c <= int32(120538))
														} else {
															v682 = libc.BoolInt32(c <= int32(120570))
														}
													}
													v680 = libc.BoolInt32(v683 || v682 != 0)
												}
												v679 = v680
											} else {
												if v688 = c <= int32(120596); !v688 {
													if c < int32(120656) {
														if c < int32(120630) {
															v685 = libc.BoolInt32(c >= int32(120598) && c <= int32(120628))
														} else {
															v685 = libc.BoolInt32(c <= int32(120654))
														}
														v684 = v685
													} else {
														if v687 = c <= int32(120686); !v687 {
															if c < int32(120714) {
																v686 = libc.BoolInt32(c >= int32(120688) && c <= int32(120712))
															} else {
																v686 = libc.BoolInt32(c <= int32(120744))
															}
														}
														v684 = libc.BoolInt32(v687 || v686 != 0)
													}
												}
												v679 = libc.BoolInt32(v688 || v684 != 0)
											}
										}
										v668 = libc.BoolInt32(v689 || v679 != 0)
									}
									v667 = v668
								} else {
									if v712 = c <= int32(120770); !v712 {
										if c < int32(122907) {
											if c < int32(121476) {
												if c < int32(121344) {
													if c < int32(120782) {
														v693 = libc.BoolInt32(c >= int32(120772) && c <= int32(120779))
													} else {
														v693 = libc.BoolInt32(c <= int32(120831))
													}
													v692 = v693
												} else {
													if v695 = c <= int32(121398); !v695 {
														if c < int32(121461) {
															v694 = libc.BoolInt32(c >= int32(121403) && c <= int32(121452))
														} else {
															v694 = libc.BoolInt32(c <= int32(121461))
														}
													}
													v692 = libc.BoolInt32(v695 || v694 != 0)
												}
												v691 = v692
											} else {
												if v700 = c <= int32(121476); !v700 {
													if c < int32(122624) {
														if c < int32(121505) {
															v697 = libc.BoolInt32(c >= int32(121499) && c <= int32(121503))
														} else {
															v697 = libc.BoolInt32(c <= int32(121519))
														}
														v696 = v697
													} else {
														if v699 = c <= int32(122654); !v699 {
															if c < int32(122888) {
																v698 = libc.BoolInt32(c >= int32(122880) && c <= int32(122886))
															} else {
																v698 = libc.BoolInt32(c <= int32(122904))
															}
														}
														v696 = libc.BoolInt32(v699 || v698 != 0)
													}
												}
												v691 = libc.BoolInt32(v700 || v696 != 0)
											}
											v690 = v691
										} else {
											if v711 = c <= int32(122913); !v711 {
												if c < int32(123214) {
													if c < int32(123136) {
														if c < int32(122918) {
															v703 = libc.BoolInt32(c >= int32(122915) && c <= int32(122916))
														} else {
															v703 = libc.BoolInt32(c <= int32(122922))
														}
														v702 = v703
													} else {
														if v705 = c <= int32(123180); !v705 {
															if c < int32(123200) {
																v704 = libc.BoolInt32(c >= int32(123184) && c <= int32(123197))
															} else {
																v704 = libc.BoolInt32(c <= int32(123209))
															}
														}
														v702 = libc.BoolInt32(v705 || v704 != 0)
													}
													v701 = v702
												} else {
													if v710 = c <= int32(123214); !v710 {
														if c < int32(124896) {
															if c < int32(123584) {
																v707 = libc.BoolInt32(c >= int32(123536) && c <= int32(123566))
															} else {
																v707 = libc.BoolInt32(c <= int32(123641))
															}
															v706 = v707
														} else {
															if v709 = c <= int32(124902); !v709 {
																if c < int32(124909) {
																	v708 = libc.BoolInt32(c >= int32(124904) && c <= int32(124907))
																} else {
																	v708 = libc.BoolInt32(c <= int32(124910))
																}
															}
															v706 = libc.BoolInt32(v709 || v708 != 0)
														}
													}
													v701 = libc.BoolInt32(v710 || v706 != 0)
												}
											}
											v690 = libc.BoolInt32(v711 || v701 != 0)
										}
									}
									v667 = libc.BoolInt32(v712 || v690 != 0)
								}
								v666 = v667
							} else {
								if v757 = c <= int32(124926); !v757 {
									if c < int32(126557) {
										if c < int32(126521) {
											if c < int32(126469) {
												if c < int32(125184) {
													if c < int32(125136) {
														v717 = libc.BoolInt32(c >= int32(124928) && c <= int32(125124))
													} else {
														v717 = libc.BoolInt32(c <= int32(125142))
													}
													v716 = v717
												} else {
													if v719 = c <= int32(125259); !v719 {
														if c < int32(126464) {
															v718 = libc.BoolInt32(c >= int32(125264) && c <= int32(125273))
														} else {
															v718 = libc.BoolInt32(c <= int32(126467))
														}
													}
													v716 = libc.BoolInt32(v719 || v718 != 0)
												}
												v715 = v716
											} else {
												if v724 = c <= int32(126495); !v724 {
													if c < int32(126503) {
														if c < int32(126500) {
															v721 = libc.BoolInt32(c >= int32(126497) && c <= int32(126498))
														} else {
															v721 = libc.BoolInt32(c <= int32(126500))
														}
														v720 = v721
													} else {
														if v723 = c <= int32(126503); !v723 {
															if c < int32(126516) {
																v722 = libc.BoolInt32(c >= int32(126505) && c <= int32(126514))
															} else {
																v722 = libc.BoolInt32(c <= int32(126519))
															}
														}
														v720 = libc.BoolInt32(v723 || v722 != 0)
													}
												}
												v715 = libc.BoolInt32(v724 || v720 != 0)
											}
											v714 = v715
										} else {
											if v735 = c <= int32(126521); !v735 {
												if c < int32(126541) {
													if c < int32(126535) {
														if c < int32(126530) {
															v727 = libc.BoolInt32(c == int32(126523))
														} else {
															v727 = libc.BoolInt32(c <= int32(126530))
														}
														v726 = v727
													} else {
														if v729 = c <= int32(126535); !v729 {
															if c < int32(126539) {
																v728 = libc.BoolInt32(c == int32(126537))
															} else {
																v728 = libc.BoolInt32(c <= int32(126539))
															}
														}
														v726 = libc.BoolInt32(v729 || v728 != 0)
													}
													v725 = v726
												} else {
													if v734 = c <= int32(126543); !v734 {
														if c < int32(126551) {
															if c < int32(126548) {
																v731 = libc.BoolInt32(c >= int32(126545) && c <= int32(126546))
															} else {
																v731 = libc.BoolInt32(c <= int32(126548))
															}
															v730 = v731
														} else {
															if v733 = c <= int32(126551); !v733 {
																if c < int32(126555) {
																	v732 = libc.BoolInt32(c == int32(126553))
																} else {
																	v732 = libc.BoolInt32(c <= int32(126555))
																}
															}
															v730 = libc.BoolInt32(v733 || v732 != 0)
														}
													}
													v725 = libc.BoolInt32(v734 || v730 != 0)
												}
											}
											v714 = libc.BoolInt32(v735 || v725 != 0)
										}
										v713 = v714
									} else {
										if v756 = c <= int32(126557); !v756 {
											if c < int32(126629) {
												if c < int32(126580) {
													if c < int32(126564) {
														if c < int32(126561) {
															v739 = libc.BoolInt32(c == int32(126559))
														} else {
															v739 = libc.BoolInt32(c <= int32(126562))
														}
														v738 = v739
													} else {
														if v741 = c <= int32(126564); !v741 {
															if c < int32(126572) {
																v740 = libc.BoolInt32(c >= int32(126567) && c <= int32(126570))
															} else {
																v740 = libc.BoolInt32(c <= int32(126578))
															}
														}
														v738 = libc.BoolInt32(v741 || v740 != 0)
													}
													v737 = v738
												} else {
													if v746 = c <= int32(126583); !v746 {
														if c < int32(126592) {
															if c < int32(126590) {
																v743 = libc.BoolInt32(c >= int32(126585) && c <= int32(126588))
															} else {
																v743 = libc.BoolInt32(c <= int32(126590))
															}
															v742 = v743
														} else {
															if v745 = c <= int32(126601); !v745 {
																if c < int32(126625) {
																	v744 = libc.BoolInt32(c >= int32(126603) && c <= int32(126619))
																} else {
																	v744 = libc.BoolInt32(c <= int32(126627))
																}
															}
															v742 = libc.BoolInt32(v745 || v744 != 0)
														}
													}
													v737 = libc.BoolInt32(v746 || v742 != 0)
												}
												v736 = v737
											} else {
												if v755 = c <= int32(126633); !v755 {
													if c < int32(178208) {
														if c < int32(131072) {
															if c < int32(130032) {
																v749 = libc.BoolInt32(c >= int32(126635) && c <= int32(126651))
															} else {
																v749 = libc.BoolInt32(c <= int32(130041))
															}
															v748 = v749
														} else {
															if v751 = c <= int32(173791); !v751 {
																if c < int32(177984) {
																	v750 = libc.BoolInt32(c >= int32(173824) && c <= int32(177976))
																} else {
																	v750 = libc.BoolInt32(c <= int32(178205))
																}
															}
															v748 = libc.BoolInt32(v751 || v750 != 0)
														}
														v747 = v748
													} else {
														if v754 = c <= int32(183969); !v754 {
															if c < int32(196608) {
																if c < int32(194560) {
																	v753 = libc.BoolInt32(c >= int32(183984) && c <= int32(191456))
																} else {
																	v753 = libc.BoolInt32(c <= int32(195101))
																}
																v752 = v753
															} else {
																v752 = libc.BoolInt32(c <= int32(201546) || c >= int32(917760) && c <= int32(917999))
															}
														}
														v747 = libc.BoolInt32(v754 || v752 != 0)
													}
												}
												v736 = libc.BoolInt32(v755 || v747 != 0)
											}
										}
										v713 = libc.BoolInt32(v756 || v736 != 0)
									}
								}
								v666 = libc.BoolInt32(v757 || v713 != 0)
							}
						}
						v571 = libc.BoolInt32(v758 || v666 != 0)
					}
				}
				v382 = libc.BoolInt32(v759 || v571 != 0)
			}
		}
		v1 = libc.BoolInt32(v760 || v382 != 0)
	}
	return uint8(libc.BoolInt32(v1 != 0))
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
	switch int32(state) {
	case 0:
		if eof != 0 {
			state = uint16(21)
			goto next_state
		}
		if lookahead == int32('!') {
			state = uint16(22)
			goto next_state
		}
		if lookahead == int32('"') {
			state = uint16(23)
			goto next_state
		}
		if lookahead == int32('\'') {
			state = uint16(26)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(29)
			goto next_state
		}
		if lookahead == int32(')') {
			state = uint16(30)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(31)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(28)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(36)
			goto next_state
		}
		if lookahead == int32(';') {
			state = uint16(35)
			goto next_state
		}
		if lookahead == int32('=') {
			state = uint16(32)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(2)
			goto next_state
		}
		if sym_identifier_character_set_1(tls, lookahead) != 0 {
			state = uint16(33)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(0)
			goto next_state
		}
		return result
	case int32(1):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(8)
			goto next_state
		}
		return result
	case int32(2):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(8)
			goto next_state
		}
		if lookahead == int32('\r') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(1)
			goto next_state
		}
		if lookahead == int32('U') {
			state = uint16(20)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(16)
			goto next_state
		}
		return result
	case int32(3):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(9)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(9)
			goto next_state
		}
		if lookahead == int32('\r') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(3)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(11)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(11)
			goto next_state
		}
		if lookahead == int32('\r') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(5)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(11)
			goto next_state
		}
		if lookahead == int32('"') {
			state = uint16(23)
			goto next_state
		}
		if lookahead == int32('\\') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(6)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(24)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(25)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead == int32('!') {
			state = uint16(22)
			goto next_state
		}
		if lookahead == int32('"') {
			state = uint16(23)
			goto next_state
		}
		if lookahead == int32('\'') {
			state = uint16(26)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(29)
			goto next_state
		}
		if lookahead == int32(')') {
			state = uint16(30)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(31)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(28)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(36)
			goto next_state
		}
		if lookahead == int32(';') {
			state = uint16(35)
			goto next_state
		}
		if lookahead == int32('=') {
			state = uint16(32)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(2)
			goto next_state
		}
		if sym_identifier_character_set_1(tls, lookahead) != 0 {
			state = uint16(33)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(8)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead == int32('"') {
			state = uint16(23)
			goto next_state
		}
		if lookahead == int32('\'') {
			state = uint16(26)
			goto next_state
		}
		if lookahead == int32('\\') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(4)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(9)
			goto next_state
		}
		return result
	case int32(10):
		if lookahead == int32('"') {
			state = uint16(23)
			goto next_state
		}
		if lookahead == int32('\'') {
			state = uint16(27)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(34)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(9)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('!') && lookahead != int32(',') && lookahead != int32(';') {
			state = uint16(34)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead == int32('"') {
			state = uint16(23)
			goto next_state
		}
		if lookahead == int32('\\') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(6)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(11)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead == int32('U') {
			state = uint16(20)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(16)
			goto next_state
		}
		return result
	case int32(13):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(14):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(13)
			goto next_state
		}
		return result
	case int32(15):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(14)
			goto next_state
		}
		return result
	case int32(16):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(15)
			goto next_state
		}
		return result
	case int32(17):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(16)
			goto next_state
		}
		return result
	case int32(18):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(17)
			goto next_state
		}
		return result
	case int32(19):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(18)
			goto next_state
		}
		return result
	case int32(20):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(19)
			goto next_state
		}
		return result
	case int32(21):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(ts_builtin_sym_end)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(22):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BANG)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(23):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(24):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_literal_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\t') || int32(11) <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(24)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('"') && lookahead != int32('\\') {
			state = uint16(25)
			goto next_state
		}
		return result
	case int32(25):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_literal_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('"') && lookahead != int32('\\') {
			state = uint16(25)
			goto next_state
		}
		return result
	case int32(26):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(27):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && (lookahead < int32(' ') || int32('"') < lookahead) && lookahead != int32(',') && lookahead != int32(';') {
			state = uint16(34)
			goto next_state
		}
		return result
	case int32(28):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(29):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(30):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(31):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COMMA)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(32):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(33):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if sym_identifier_character_set_2(tls, lookahead) != 0 {
			state = uint16(33)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(12)
			goto next_state
		}
		return result
	case int32(34):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_value_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && (lookahead < int32(' ') || int32('"') < lookahead) && lookahead != int32(',') && lookahead != int32(';') {
			state = uint16(34)
			goto next_state
		}
		return result
	case int32(35):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SEMI)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(36):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SLASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var ts_lex_modes = [54]TSLexMode{
	0:  {},
	1:  {},
	2:  {},
	3:  {},
	4:  {},
	5:  {},
	6:  {},
	7:  {},
	8:  {},
	9:  {},
	10: {},
	11: {},
	12: {},
	13: {},
	14: {},
	15: {},
	16: {},
	17: {},
	18: {},
	19: {},
	20: {},
	21: {},
	22: {},
	23: {},
	24: {},
	25: {},
	26: {},
	27: {},
	28: {},
	29: {},
	30: {},
	31: {},
	32: {},
	33: {},
	34: {},
	35: {},
	36: {},
	37: {},
	38: {},
	39: {},
	40: {
		Flex_state: uint16(10),
	},
	41: {
		Flex_state: uint16(7),
	},
	42: {
		Flex_state: uint16(7),
	},
	43: {
		Flex_state: uint16(7),
	},
	44: {},
	45: {},
	46: {},
	47: {},
	48: {},
	49: {},
	50: {},
	51: {},
	52: {},
	53: {},
}

var ts_parse_table = [2][32]uint16_t{
	0: {
		0:  uint16(1),
		1:  uint16(1),
		2:  uint16(1),
		4:  uint16(1),
		5:  uint16(1),
		6:  uint16(1),
		7:  uint16(1),
		8:  uint16(1),
		9:  uint16(1),
		10: uint16(1),
		12: uint16(1),
		13: uint16(1),
	},
	1: {
		0:  uint16(3),
		5:  uint16(5),
		10: uint16(7),
		14: uint16(53),
		15: uint16(6),
		16: uint16(31),
		18: uint16(37),
		19: uint16(37),
		20: uint16(37),
		23: uint16(37),
		24: uint16(25),
		25: uint16(6),
	},
}

var ts_small_parse_table = [664]uint16_t{
	0:   uint16(8),
	1:   uint16(5),
	2:   uint16(1),
	3:   uint16(anon_sym_DOT),
	4:   uint16(9),
	5:   uint16(1),
	6:   uint16(anon_sym_RPAREN),
	7:   uint16(11),
	8:   uint16(1),
	9:   uint16(sym_identifier),
	10:  uint16(25),
	11:  uint16(1),
	12:  uint16(sym_cap),
	13:  uint16(31),
	14:  uint16(1),
	15:  uint16(sym_element),
	16:  uint16(5),
	17:  uint16(2),
	18:  uint16(sym_fragment),
	19:  uint16(aux_sym_pipeline_repeat1),
	20:  uint16(23),
	21:  uint16(2),
	22:  uint16(sym_property),
	23:  uint16(aux_sym_bin_repeat1),
	24:  uint16(37),
	25:  uint16(4),
	26:  uint16(sym_bin),
	27:  uint16(sym_simple_element),
	28:  uint16(sym_reference),
	29:  uint16(sym_caps),
	30:  uint16(8),
	31:  uint16(5),
	32:  uint16(1),
	33:  uint16(anon_sym_DOT),
	34:  uint16(11),
	35:  uint16(1),
	36:  uint16(sym_identifier),
	37:  uint16(13),
	38:  uint16(1),
	39:  uint16(anon_sym_RPAREN),
	40:  uint16(25),
	41:  uint16(1),
	42:  uint16(sym_cap),
	43:  uint16(31),
	44:  uint16(1),
	45:  uint16(sym_element),
	46:  uint16(2),
	47:  uint16(2),
	48:  uint16(sym_property),
	49:  uint16(aux_sym_bin_repeat1),
	50:  uint16(7),
	51:  uint16(2),
	52:  uint16(sym_fragment),
	53:  uint16(aux_sym_pipeline_repeat1),
	54:  uint16(37),
	55:  uint16(4),
	56:  uint16(sym_bin),
	57:  uint16(sym_simple_element),
	58:  uint16(sym_reference),
	59:  uint16(sym_caps),
	60:  uint16(7),
	61:  uint16(17),
	62:  uint16(1),
	63:  uint16(anon_sym_DOT),
	64:  uint16(20),
	65:  uint16(1),
	66:  uint16(sym_identifier),
	67:  uint16(25),
	68:  uint16(1),
	69:  uint16(sym_cap),
	70:  uint16(31),
	71:  uint16(1),
	72:  uint16(sym_element),
	73:  uint16(15),
	74:  uint16(2),
	76:  uint16(anon_sym_RPAREN),
	77:  uint16(4),
	78:  uint16(2),
	79:  uint16(sym_fragment),
	80:  uint16(aux_sym_pipeline_repeat1),
	81:  uint16(37),
	82:  uint16(4),
	83:  uint16(sym_bin),
	84:  uint16(sym_simple_element),
	85:  uint16(sym_reference),
	86:  uint16(sym_caps),
	87:  uint16(7),
	88:  uint16(5),
	89:  uint16(1),
	90:  uint16(anon_sym_DOT),
	91:  uint16(7),
	92:  uint16(1),
	93:  uint16(sym_identifier),
	94:  uint16(23),
	95:  uint16(1),
	96:  uint16(anon_sym_RPAREN),
	97:  uint16(25),
	98:  uint16(1),
	99:  uint16(sym_cap),
	100: uint16(31),
	101: uint16(1),
	102: uint16(sym_element),
	103: uint16(4),
	104: uint16(2),
	105: uint16(sym_fragment),
	106: uint16(aux_sym_pipeline_repeat1),
	107: uint16(37),
	108: uint16(4),
	109: uint16(sym_bin),
	110: uint16(sym_simple_element),
	111: uint16(sym_reference),
	112: uint16(sym_caps),
	113: uint16(7),
	114: uint16(5),
	115: uint16(1),
	116: uint16(anon_sym_DOT),
	117: uint16(7),
	118: uint16(1),
	119: uint16(sym_identifier),
	120: uint16(25),
	121: uint16(1),
	123: uint16(25),
	124: uint16(1),
	125: uint16(sym_cap),
	126: uint16(31),
	127: uint16(1),
	128: uint16(sym_element),
	129: uint16(4),
	130: uint16(2),
	131: uint16(sym_fragment),
	132: uint16(aux_sym_pipeline_repeat1),
	133: uint16(37),
	134: uint16(4),
	135: uint16(sym_bin),
	136: uint16(sym_simple_element),
	137: uint16(sym_reference),
	138: uint16(sym_caps),
	139: uint16(7),
	140: uint16(5),
	141: uint16(1),
	142: uint16(anon_sym_DOT),
	143: uint16(7),
	144: uint16(1),
	145: uint16(sym_identifier),
	146: uint16(9),
	147: uint16(1),
	148: uint16(anon_sym_RPAREN),
	149: uint16(25),
	150: uint16(1),
	151: uint16(sym_cap),
	152: uint16(31),
	153: uint16(1),
	154: uint16(sym_element),
	155: uint16(4),
	156: uint16(2),
	157: uint16(sym_fragment),
	158: uint16(aux_sym_pipeline_repeat1),
	159: uint16(37),
	160: uint16(4),
	161: uint16(sym_bin),
	162: uint16(sym_simple_element),
	163: uint16(sym_reference),
	164: uint16(sym_caps),
	165: uint16(5),
	166: uint16(5),
	167: uint16(1),
	168: uint16(anon_sym_DOT),
	169: uint16(7),
	170: uint16(1),
	171: uint16(sym_identifier),
	172: uint16(25),
	173: uint16(1),
	174: uint16(sym_cap),
	175: uint16(39),
	176: uint16(1),
	177: uint16(sym_element),
	178: uint16(37),
	179: uint16(4),
	180: uint16(sym_bin),
	181: uint16(sym_simple_element),
	182: uint16(sym_reference),
	183: uint16(sym_caps),
	184: uint16(3),
	185: uint16(29),
	186: uint16(1),
	187: uint16(anon_sym_COMMA),
	188: uint16(9),
	189: uint16(1),
	190: uint16(aux_sym_cap_repeat1),
	191: uint16(27),
	192: uint16(6),
	194: uint16(anon_sym_BANG),
	195: uint16(anon_sym_DOT),
	196: uint16(anon_sym_RPAREN),
	197: uint16(sym_identifier),
	198: uint16(anon_sym_SEMI),
	199: uint16(5),
	200: uint16(34),
	201: uint16(1),
	202: uint16(anon_sym_DOT),
	203: uint16(37),
	204: uint16(1),
	205: uint16(sym_identifier),
	206: uint16(40),
	207: uint16(1),
	208: uint16(anon_sym_SLASH),
	209: uint16(16),
	210: uint16(2),
	211: uint16(sym_property),
	212: uint16(aux_sym_bin_repeat1),
	213: uint16(32),
	214: uint16(3),
	216: uint16(anon_sym_BANG),
	217: uint16(anon_sym_RPAREN),
	218: uint16(3),
	219: uint16(44),
	220: uint16(1),
	221: uint16(anon_sym_COMMA),
	222: uint16(9),
	223: uint16(1),
	224: uint16(aux_sym_cap_repeat1),
	225: uint16(42),
	226: uint16(6),
	228: uint16(anon_sym_BANG),
	229: uint16(anon_sym_DOT),
	230: uint16(anon_sym_RPAREN),
	231: uint16(sym_identifier),
	232: uint16(anon_sym_SEMI),
	233: uint16(3),
	234: uint16(44),
	235: uint16(1),
	236: uint16(anon_sym_COMMA),
	237: uint16(11),
	238: uint16(1),
	239: uint16(aux_sym_cap_repeat1),
	240: uint16(46),
	241: uint16(6),
	243: uint16(anon_sym_BANG),
	244: uint16(anon_sym_DOT),
	245: uint16(anon_sym_RPAREN),
	246: uint16(sym_identifier),
	247: uint16(anon_sym_SEMI),
	248: uint16(6),
	249: uint16(34),
	250: uint16(1),
	251: uint16(anon_sym_DOT),
	252: uint16(37),
	253: uint16(1),
	254: uint16(sym_identifier),
	255: uint16(40),
	256: uint16(1),
	257: uint16(anon_sym_SLASH),
	258: uint16(48),
	259: uint16(1),
	260: uint16(anon_sym_EQ),
	261: uint16(32),
	262: uint16(2),
	263: uint16(anon_sym_BANG),
	264: uint16(anon_sym_RPAREN),
	265: uint16(16),
	266: uint16(2),
	267: uint16(sym_property),
	268: uint16(aux_sym_bin_repeat1),
	269: uint16(1),
	270: uint16(50),
	271: uint16(7),
	273: uint16(anon_sym_BANG),
	274: uint16(anon_sym_DOT),
	275: uint16(anon_sym_RPAREN),
	276: uint16(anon_sym_COMMA),
	277: uint16(sym_identifier),
	278: uint16(anon_sym_SEMI),
	279: uint16(3),
	280: uint16(54),
	281: uint16(1),
	282: uint16(anon_sym_COMMA),
	283: uint16(27),
	284: uint16(1),
	285: uint16(aux_sym_reference_repeat1),
	286: uint16(52),
	287: uint16(5),
	289: uint16(anon_sym_BANG),
	290: uint16(anon_sym_DOT),
	291: uint16(anon_sym_RPAREN),
	292: uint16(sym_identifier),
	293: uint16(3),
	294: uint16(58),
	295: uint16(1),
	296: uint16(sym_identifier),
	297: uint16(23),
	298: uint16(2),
	299: uint16(sym_property),
	300: uint16(aux_sym_bin_repeat1),
	301: uint16(56),
	302: uint16(4),
	304: uint16(anon_sym_BANG),
	305: uint16(anon_sym_DOT),
	306: uint16(anon_sym_RPAREN),
	307: uint16(1),
	308: uint16(27),
	309: uint16(7),
	311: uint16(anon_sym_BANG),
	312: uint16(anon_sym_DOT),
	313: uint16(anon_sym_RPAREN),
	314: uint16(anon_sym_COMMA),
	315: uint16(sym_identifier),
	316: uint16(anon_sym_SEMI),
	317: uint16(3),
	318: uint16(63),
	319: uint16(1),
	320: uint16(anon_sym_SEMI),
	321: uint16(26),
	322: uint16(1),
	323: uint16(aux_sym_caps_repeat1),
	324: uint16(61),
	325: uint16(5),
	327: uint16(anon_sym_BANG),
	328: uint16(anon_sym_DOT),
	329: uint16(anon_sym_RPAREN),
	330: uint16(sym_identifier),
	331: uint16(3),
	332: uint16(54),
	333: uint16(1),
	334: uint16(anon_sym_COMMA),
	335: uint16(27),
	336: uint16(1),
	337: uint16(aux_sym_reference_repeat1),
	338: uint16(65),
	339: uint16(5),
	341: uint16(anon_sym_BANG),
	342: uint16(anon_sym_DOT),
	343: uint16(anon_sym_RPAREN),
	344: uint16(sym_identifier),
	345: uint16(3),
	346: uint16(54),
	347: uint16(1),
	348: uint16(anon_sym_COMMA),
	349: uint16(15),
	350: uint16(1),
	351: uint16(aux_sym_reference_repeat1),
	352: uint16(67),
	353: uint16(5),
	355: uint16(anon_sym_BANG),
	356: uint16(anon_sym_DOT),
	357: uint16(anon_sym_RPAREN),
	358: uint16(sym_identifier),
	359: uint16(1),
	360: uint16(69),
	361: uint16(7),
	363: uint16(anon_sym_BANG),
	364: uint16(anon_sym_DOT),
	365: uint16(anon_sym_RPAREN),
	366: uint16(anon_sym_COMMA),
	367: uint16(sym_identifier),
	368: uint16(anon_sym_SEMI),
	369: uint16(2),
	370: uint16(73),
	371: uint16(1),
	372: uint16(anon_sym_COMMA),
	373: uint16(71),
	374: uint16(6),
	376: uint16(anon_sym_BANG),
	377: uint16(anon_sym_DOT),
	378: uint16(anon_sym_RPAREN),
	379: uint16(sym_identifier),
	380: uint16(anon_sym_SEMI),
	381: uint16(3),
	382: uint16(77),
	383: uint16(1),
	384: uint16(sym_identifier),
	385: uint16(23),
	386: uint16(2),
	387: uint16(sym_property),
	388: uint16(aux_sym_bin_repeat1),
	389: uint16(75),
	390: uint16(4),
	392: uint16(anon_sym_BANG),
	393: uint16(anon_sym_DOT),
	394: uint16(anon_sym_RPAREN),
	395: uint16(1),
	396: uint16(80),
	397: uint16(7),
	399: uint16(anon_sym_BANG),
	400: uint16(anon_sym_DOT),
	401: uint16(anon_sym_RPAREN),
	402: uint16(anon_sym_COMMA),
	403: uint16(sym_identifier),
	404: uint16(anon_sym_SEMI),
	405: uint16(3),
	406: uint16(63),
	407: uint16(1),
	408: uint16(anon_sym_SEMI),
	409: uint16(18),
	410: uint16(1),
	411: uint16(aux_sym_caps_repeat1),
	412: uint16(82),
	413: uint16(5),
	415: uint16(anon_sym_BANG),
	416: uint16(anon_sym_DOT),
	417: uint16(anon_sym_RPAREN),
	418: uint16(sym_identifier),
	419: uint16(3),
	420: uint16(86),
	421: uint16(1),
	422: uint16(anon_sym_SEMI),
	423: uint16(26),
	424: uint16(1),
	425: uint16(aux_sym_caps_repeat1),
	426: uint16(84),
	427: uint16(5),
	429: uint16(anon_sym_BANG),
	430: uint16(anon_sym_DOT),
	431: uint16(anon_sym_RPAREN),
	432: uint16(sym_identifier),
	433: uint16(3),
	434: uint16(91),
	435: uint16(1),
	436: uint16(anon_sym_COMMA),
	437: uint16(27),
	438: uint16(1),
	439: uint16(aux_sym_reference_repeat1),
	440: uint16(89),
	441: uint16(5),
	443: uint16(anon_sym_BANG),
	444: uint16(anon_sym_DOT),
	445: uint16(anon_sym_RPAREN),
	446: uint16(sym_identifier),
	447: uint16(3),
	448: uint16(54),
	449: uint16(1),
	450: uint16(anon_sym_COMMA),
	451: uint16(19),
	452: uint16(1),
	453: uint16(aux_sym_reference_repeat1),
	454: uint16(94),
	455: uint16(5),
	457: uint16(anon_sym_BANG),
	458: uint16(anon_sym_DOT),
	459: uint16(anon_sym_RPAREN),
	460: uint16(sym_identifier),
	461: uint16(1),
	462: uint16(96),
	463: uint16(7),
	465: uint16(anon_sym_BANG),
	466: uint16(anon_sym_DOT),
	467: uint16(anon_sym_RPAREN),
	468: uint16(anon_sym_COMMA),
	469: uint16(sym_identifier),
	470: uint16(anon_sym_SEMI),
	471: uint16(3),
	472: uint16(100),
	473: uint16(1),
	474: uint16(anon_sym_BANG),
	475: uint16(32),
	476: uint16(1),
	477: uint16(aux_sym_fragment_repeat1),
	478: uint16(98),
	479: uint16(4),
	481: uint16(anon_sym_DOT),
	482: uint16(anon_sym_RPAREN),
	483: uint16(sym_identifier),
	484: uint16(3),
	485: uint16(100),
	486: uint16(1),
	487: uint16(anon_sym_BANG),
	488: uint16(30),
	489: uint16(1),
	490: uint16(aux_sym_fragment_repeat1),
	491: uint16(102),
	492: uint16(4),
	494: uint16(anon_sym_DOT),
	495: uint16(anon_sym_RPAREN),
	496: uint16(sym_identifier),
	497: uint16(3),
	498: uint16(106),
	499: uint16(1),
	500: uint16(anon_sym_BANG),
	501: uint16(32),
	502: uint16(1),
	503: uint16(aux_sym_fragment_repeat1),
	504: uint16(104),
	505: uint16(4),
	507: uint16(anon_sym_DOT),
	508: uint16(anon_sym_RPAREN),
	509: uint16(sym_identifier),
	510: uint16(1),
	511: uint16(84),
	512: uint16(6),
	514: uint16(anon_sym_BANG),
	515: uint16(anon_sym_DOT),
	516: uint16(anon_sym_RPAREN),
	517: uint16(sym_identifier),
	518: uint16(anon_sym_SEMI),
	519: uint16(1),
	520: uint16(109),
	521: uint16(6),
	523: uint16(anon_sym_BANG),
	524: uint16(anon_sym_DOT),
	525: uint16(anon_sym_RPAREN),
	526: uint16(anon_sym_COMMA),
	527: uint16(sym_identifier),
	528: uint16(1),
	529: uint16(111),
	530: uint16(5),
	532: uint16(anon_sym_BANG),
	533: uint16(anon_sym_DOT),
	534: uint16(anon_sym_RPAREN),
	535: uint16(sym_identifier),
	536: uint16(1),
	537: uint16(113),
	538: uint16(5),
	540: uint16(anon_sym_BANG),
	541: uint16(anon_sym_DOT),
	542: uint16(anon_sym_RPAREN),
	543: uint16(sym_identifier),
	544: uint16(1),
	545: uint16(115),
	546: uint16(5),
	548: uint16(anon_sym_BANG),
	549: uint16(anon_sym_DOT),
	550: uint16(anon_sym_RPAREN),
	551: uint16(sym_identifier),
	552: uint16(1),
	553: uint16(117),
	554: uint16(5),
	556: uint16(anon_sym_BANG),
	557: uint16(anon_sym_DOT),
	558: uint16(anon_sym_RPAREN),
	559: uint16(sym_identifier),
	560: uint16(1),
	561: uint16(104),
	562: uint16(5),
	564: uint16(anon_sym_BANG),
	565: uint16(anon_sym_DOT),
	566: uint16(anon_sym_RPAREN),
	567: uint16(sym_identifier),
	568: uint16(4),
	569: uint16(121),
	570: uint16(1),
	571: uint16(aux_sym_value_token1),
	572: uint16(14),
	573: uint16(1),
	574: uint16(sym_value),
	575: uint16(29),
	576: uint16(1),
	577: uint16(sym_string_literal),
	578: uint16(119),
	579: uint16(2),
	580: uint16(anon_sym_DQUOTE),
	581: uint16(anon_sym_SQUOTE),
	582: uint16(3),
	583: uint16(123),
	584: uint16(1),
	585: uint16(anon_sym_DQUOTE),
	586: uint16(125),
	587: uint16(1),
	588: uint16(aux_sym_string_literal_token1),
	589: uint16(42),
	590: uint16(1),
	591: uint16(aux_sym_string_literal_repeat1),
	592: uint16(3),
	593: uint16(127),
	594: uint16(1),
	595: uint16(anon_sym_DQUOTE),
	596: uint16(129),
	597: uint16(1),
	598: uint16(aux_sym_string_literal_token1),
	599: uint16(43),
	600: uint16(1),
	601: uint16(aux_sym_string_literal_repeat1),
	602: uint16(3),
	603: uint16(131),
	604: uint16(1),
	605: uint16(anon_sym_DQUOTE),
	606: uint16(133),
	607: uint16(1),
	608: uint16(aux_sym_string_literal_token1),
	609: uint16(43),
	610: uint16(1),
	611: uint16(aux_sym_string_literal_repeat1),
	612: uint16(2),
	613: uint16(136),
	614: uint16(1),
	615: uint16(anon_sym_LPAREN),
	616: uint16(138),
	617: uint16(1),
	618: uint16(sym_identifier),
	619: uint16(2),
	620: uint16(140),
	621: uint16(1),
	622: uint16(sym_identifier),
	623: uint16(33),
	624: uint16(1),
	625: uint16(sym_cap),
	626: uint16(2),
	627: uint16(142),
	628: uint16(1),
	629: uint16(sym_identifier),
	630: uint16(12),
	631: uint16(1),
	632: uint16(sym_property),
	633: uint16(2),
	634: uint16(142),
	635: uint16(1),
	636: uint16(sym_identifier),
	637: uint16(17),
	638: uint16(1),
	639: uint16(sym_property),
	640: uint16(1),
	641: uint16(48),
	642: uint16(1),
	643: uint16(anon_sym_EQ),
	644: uint16(1),
	645: uint16(144),
	646: uint16(1),
	647: uint16(sym_identifier),
	648: uint16(1),
	649: uint16(146),
	650: uint16(1),
	651: uint16(sym_identifier),
	652: uint16(1),
	653: uint16(40),
	654: uint16(1),
	655: uint16(anon_sym_SLASH),
	656: uint16(1),
	657: uint16(148),
	658: uint16(1),
	659: uint16(sym_identifier),
	660: uint16(1),
	661: uint16(150),
	662: uint16(1),
}

var ts_small_parse_table_map = [52]uint32_t{
	1:  uint32(30),
	2:  uint32(60),
	3:  uint32(87),
	4:  uint32(113),
	5:  uint32(139),
	6:  uint32(165),
	7:  uint32(184),
	8:  uint32(199),
	9:  uint32(218),
	10: uint32(233),
	11: uint32(248),
	12: uint32(269),
	13: uint32(279),
	14: uint32(293),
	15: uint32(307),
	16: uint32(317),
	17: uint32(331),
	18: uint32(345),
	19: uint32(359),
	20: uint32(369),
	21: uint32(381),
	22: uint32(395),
	23: uint32(405),
	24: uint32(419),
	25: uint32(433),
	26: uint32(447),
	27: uint32(461),
	28: uint32(471),
	29: uint32(484),
	30: uint32(497),
	31: uint32(510),
	32: uint32(519),
	33: uint32(528),
	34: uint32(536),
	35: uint32(544),
	36: uint32(552),
	37: uint32(560),
	38: uint32(568),
	39: uint32(582),
	40: uint32(592),
	41: uint32(602),
	42: uint32(612),
	43: uint32(619),
	44: uint32(626),
	45: uint32(633),
	46: uint32(640),
	47: uint32(644),
	48: uint32(648),
	49: uint32(652),
	50: uint32(656),
	51: uint32(660),
}

var ts_parse_actions = [152]TSParseActionEntry{
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
		Fsymbol:     uint16(sym_pipeline),
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
		Fstate: uint16(libc.Int32FromInt32(49)),
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
		Fstate: uint16(libc.Int32FromInt32(10)),
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
		Fstate: uint16(libc.Int32FromInt32(36)),
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
		Fstate: uint16(libc.Int32FromInt32(13)),
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
		Fstate: uint16(libc.Int32FromInt32(35)),
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
		Fsymbol:      uint16(aux_sym_pipeline_repeat1),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_pipeline_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(49)),
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
		Fsymbol:      uint16(aux_sym_pipeline_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(10)),
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
		Fstate: uint16(libc.Int32FromInt32(38)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_pipeline),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_cap_repeat1),
	})))),
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
		Fsymbol:      uint16(aux_sym_cap_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(47)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	33: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_simple_element),
		Fproduction_id: uint16(1),
	})))),
	34: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	35: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_simple_element),
		Fproduction_id: uint16(1),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(44)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	38: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_simple_element),
		Fproduction_id: uint16(1),
	})))),
	39: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	40: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	41: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	42: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	43: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_cap),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	45: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	46: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	47: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_cap),
	})))),
	48: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(40)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	51: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_property),
		Fproduction_id: uint16(7),
	})))),
	52: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	53: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_reference),
		Fproduction_id: uint16(6),
	})))),
	54: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(52)),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_simple_element),
		Fproduction_id: uint16(1),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	59: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_simple_element),
		Fproduction_id: uint16(1),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(48)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_caps),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(45)),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_reference),
		Fproduction_id: uint16(3),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_reference),
		Fproduction_id: uint16(4),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_string_literal),
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
		Fsymbol:      uint16(sym_cap),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_bin_repeat1),
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
		Fsymbol:      uint16(aux_sym_bin_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(48)),
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
		Fcount:    uint8(1),
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
		Fsymbol:      uint16(sym_string_literal),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_caps),
	})))),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_caps_repeat1),
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
		Fsymbol:      uint16(aux_sym_caps_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(45)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	90: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_reference_repeat1),
		Fproduction_id: uint16(5),
	})))),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_reference_repeat1),
		Fproduction_id: uint16(5),
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
		Fstate:      uint16(libc.Int32FromInt32(52)),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_reference),
		Fproduction_id: uint16(2),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(sym_value),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_fragment),
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
		Fstate: uint16(libc.Int32FromInt32(8)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_fragment),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_fragment_repeat1),
	})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_fragment_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(8)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	110: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_reference_repeat1),
		Fproduction_id: uint16(2),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_bin),
		Fproduction_id: uint16(1),
	})))),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_bin),
		Fproduction_id: uint16(1),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_element),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_bin),
		Fproduction_id: uint16(1),
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
		Fstate: uint16(libc.Int32FromInt32(41)),
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
		Fstate: uint16(libc.Int32FromInt32(29)),
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
		Fstate: uint16(libc.Int32FromInt32(24)),
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
		Fstate: uint16(libc.Int32FromInt32(42)),
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
		Fstate: uint16(libc.Int32FromInt32(21)),
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
		Fstate: uint16(libc.Int32FromInt32(43)),
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
		Fcount: uint8(1),
	}})),
	132: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_string_literal_repeat1),
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
		Fsymbol:      uint16(aux_sym_string_literal_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(43)),
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
		Fstate: uint16(libc.Int32FromInt32(3)),
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
		Fstate: uint16(libc.Int32FromInt32(20)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(34)),
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
	151: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f uint8_t
		_ [7]byte
	}{f: uint8(TSParseActionTypeAccept)})),
}

func tree_sitter_gstlaunch(tls *libc.TLS) (r uintptr) {
	return uintptr(unsafe.Pointer(&language))
}

var language = TSLanguage{
	Fversion:                   uint32(LANGUAGE_VERSION),
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

var __ccgo_ts1 = "end\x00!\x00\"\x00string_content\x00'\x00.\x00(\x00)\x00,\x00=\x00identifier\x00value_token1\x00;\x00/\x00pipeline\x00fragment\x00element\x00string_literal\x00bin\x00simple_element\x00reference\x00property\x00value\x00caps\x00cap\x00pipeline_repeat1\x00fragment_repeat1\x00string_literal_repeat1\x00bin_repeat1\x00reference_repeat1\x00caps_repeat1\x00cap_repeat1\x00key\x00pad\x00type\x00"
