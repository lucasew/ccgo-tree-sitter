// Code generated for windows/arm64 by 'ccgo -ignore-unsupported-alignment -ignore-unsupported-atomic-sizes -ignore-vector-functions --winapi-no-errno -ignore-link-errors -std=gnu11 -O0 -D__extension__= -D__forceinline=static inline -D__attribute__(...)= -D__declspec(x)= -D__cdecl= -D__stdcall= -D__fastcall= -D__thiscall= -D_cdecl= -D__restrict= -D__restrict__= -D__MINGW_EXTENSION= -D_X86INTRIN_H_INCLUDED -D_X86GPRINTRIN_H_INCLUDED -D_IMMINTRIN_H_INCLUDED -D_MMINTRIN_H_INCLUDED -D_XMMINTRIN_H_INCLUDED -D_EMMINTRIN_H_INCLUDED -D_PMMINTRIN_H_INCLUDED -D_MM3DNOW_H_INCLUDED -Dfunc=func_token -Dinterface=interface_token -Dselect=select_token -Dchan=chan_token -Dgo=go_token -Dmap=map_token -Dpackage=package_token -Dtype=type_token -Dvar=var_token -Dimport=import_token -Ddefer=defer_token -Dfallthrough=fallthrough_token -Drange=range_token -I C:\a\ccgo-tree-sitter\ccgo-tree-sitter\third-party\tree-sitter-xml\dtd\src -I C:\a\ccgo-tree-sitter\ccgo-tree-sitter\third-party\tree-sitter-xml\dtd -I C:\a\ccgo-tree-sitter\ccgo-tree-sitter\third-party\tree-sitter-xml -I C:\Users\runneradmin\.cache\workspaced\sources\github\7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2\lib\include -I C:\Users\runneradmin\.cache\workspaced\sources\github\7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2\lib\src combined.c -o grammar.go', DO NOT EDIT.

//go:build windows && arm64

package grammar_dtd

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
const EXTERNAL_TOKEN_COUNT = 3
const FIELD_COUNT = 1
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
const MAX_ALIAS_SEQUENCE_LENGTH = 10
const MAX_RESERVED_WORD_SET_SIZE = 0
const MB_LEN_MAX = 1
const MINGW_HAS_DDK_H = 1
const MINGW_HAS_SECURE_API = 1
const PRODUCTION_ID_COUNT = 2
const PTRDIFF_MAX = "__PTRDIFF_MAX__"
const RAND_MAX = 0x7fff
const SIZE_MAX = "__SIZE_MAX__"
const STATE_COUNT = 334
const SUPERTYPE_COUNT = 0
const SYMBOL_COUNT = 111
const TOKEN_COUNT = 61
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
const _BLANK = 0x40
const _CALL_REPORTFAULT = 0x2
const _CONTROL = 0x20
const _CRTIMP2 = "_CRTIMP"
const _CRTIMP_ALTERNATIVE = "_CRTIMP"
const _CRTIMP_NOIA64 = "_CRTIMP"
const _CRTIMP_PURE = "_CRTIMP"
const _DIGIT = 4
const _EMMINTRIN_H_INCLUDED = 1
const _FREEENTRY = 0
const _HEAP_MAXREQ = 0xFFFFFFFFFFFFFFE0
const _HEX = 0x80
const _IMMINTRIN_H_INCLUDED = 1
const _LEADBYTE = 0x8000
const _LOWER = 2
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
const _PUNCT = 0x10
const _REPORT_ERRMODE = 3
const _SECURECRT_FILL_BUFFER_PATTERN = 0xFD
const _SPACE = 0x8
const _UPPER = 1
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

type wctrans_t = uint16

type TokenType = int32

const PI_TARGET = 0
const PI_CONTENT = 1
const COMMENT = 2

func scan_pi_target(tls *libc.TLS, lexer1 uintptr, valid_symbols uintptr) (r uint8) {
	var advanced_once, found_x_first, v2 uint8
	var v1 wchar_t
	var v4 uintptr
	_, _, _, _, _ = advanced_once, found_x_first, v1, v2, v4
	advanced_once = libc.BoolUint8(false1 != 0)
	found_x_first = libc.BoolUint8(false1 != 0)
	_ = valid_symbols
	v1 = uint16((*TSLexer)(unsafe.Pointer(lexer1)).Flookahead)
	v2 = libc.BoolUint8(libc.Xiswctype(tls, v1, uint16(libc.Int32FromInt32(0x0100)|libc.Int32FromInt32(_UPPER)|libc.Int32FromInt32(_LOWER))) != 0 || int32(v1) == int32('_') || int32(v1) == int32(':'))
	goto _3
_3:
	if v2 != 0 {
		if (*TSLexer)(unsafe.Pointer(lexer1)).Flookahead == int32('x') || (*TSLexer)(unsafe.Pointer(lexer1)).Flookahead == int32('X') {
			found_x_first = libc.BoolUint8(true1 != 0)
			(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer1)).Fmark_end})))(tls, lexer1)
		}
		advanced_once = libc.BoolUint8(true1 != 0)
		v4 = lexer1
		(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(v4)).Fadvance})))(tls, v4, libc.BoolUint8(false1 != 0))
	}
	if advanced_once != 0 {
		for {
			v1 = uint16((*TSLexer)(unsafe.Pointer(lexer1)).Flookahead)
			v2 = libc.BoolUint8(libc.Xiswctype(tls, v1, uint16(libc.Int32FromInt32(0x0100)|libc.Int32FromInt32(_UPPER)|libc.Int32FromInt32(_LOWER)|libc.Int32FromInt32(_DIGIT))) != 0 || int32(v1) == int32('_') || int32(v1) == int32(':') || int32(v1) == int32('.') || int32(v1) == int32('-') || int32(v1) == int32(0xB7))
			goto _7
		_7:
			if !(v2 != 0) {
				break
			}
			if found_x_first != 0 && ((*TSLexer)(unsafe.Pointer(lexer1)).Flookahead == int32('m') || (*TSLexer)(unsafe.Pointer(lexer1)).Flookahead == int32('M')) {
				v4 = lexer1
				(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(v4)).Fadvance})))(tls, v4, libc.BoolUint8(false1 != 0))
				if (*TSLexer)(unsafe.Pointer(lexer1)).Flookahead == int32('l') || (*TSLexer)(unsafe.Pointer(lexer1)).Flookahead == int32('L') {
					v4 = lexer1
					(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(v4)).Fadvance})))(tls, v4, libc.BoolUint8(false1 != 0))
					v1 = uint16((*TSLexer)(unsafe.Pointer(lexer1)).Flookahead)
					v2 = libc.BoolUint8(libc.Xiswctype(tls, v1, uint16(libc.Int32FromInt32(0x0100)|libc.Int32FromInt32(_UPPER)|libc.Int32FromInt32(_LOWER)|libc.Int32FromInt32(_DIGIT))) != 0 || int32(v1) == int32('_') || int32(v1) == int32(':') || int32(v1) == int32('.') || int32(v1) == int32('-') || int32(v1) == int32(0xB7))
					goto _12
				_12:
					if v2 != 0 {
					} else {
						return libc.BoolUint8(false1 != 0)
					}
				}
			}
			found_x_first = libc.BoolUint8(false1 != 0)
			v4 = lexer1
			(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(v4)).Fadvance})))(tls, v4, libc.BoolUint8(false1 != 0))
		}
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer1)).Fmark_end})))(tls, lexer1)
		(*TSLexer)(unsafe.Pointer(lexer1)).Fresult_symbol = uint16(PI_TARGET)
		return libc.BoolUint8(true1 != 0)
	}
	return libc.BoolUint8(false1 != 0)
}

func scan_pi_content(tls *libc.TLS, lexer1 uintptr) (r uint8) {
	var v1 uintptr
	_ = v1
	for !((*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer1)).Feof})))(tls, lexer1) != 0) && (*TSLexer)(unsafe.Pointer(lexer1)).Flookahead != int32('\n') && (*TSLexer)(unsafe.Pointer(lexer1)).Flookahead != int32('?') {
		v1 = lexer1
		(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(v1)).Fadvance})))(tls, v1, libc.BoolUint8(false1 != 0))
	}
	if (*TSLexer)(unsafe.Pointer(lexer1)).Flookahead != int32('?') {
		return libc.BoolUint8(false1 != 0)
	}
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer1)).Fmark_end})))(tls, lexer1)
	v1 = lexer1
	(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(v1)).Fadvance})))(tls, v1, libc.BoolUint8(false1 != 0))
	if (*TSLexer)(unsafe.Pointer(lexer1)).Flookahead == int32('>') {
		v1 = lexer1
		(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(v1)).Fadvance})))(tls, v1, libc.BoolUint8(false1 != 0))
		for (*TSLexer)(unsafe.Pointer(lexer1)).Flookahead == int32(' ') {
			v1 = lexer1
			(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(v1)).Fadvance})))(tls, v1, libc.BoolUint8(false1 != 0))
		}
		if !((*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer1)).Feof})))(tls, lexer1) != 0) && (*TSLexer)(unsafe.Pointer(lexer1)).Flookahead == int32('\n') {
			v1 = lexer1
			(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(v1)).Fadvance})))(tls, v1, libc.BoolUint8(false1 != 0))
		} else {
			return libc.BoolUint8(false1 != 0)
		}
		(*TSLexer)(unsafe.Pointer(lexer1)).Fresult_symbol = uint16(PI_CONTENT)
		return libc.BoolUint8(true1 != 0)
	}
	return libc.BoolUint8(false1 != 0)
}

func scan_comment(tls *libc.TLS, lexer1 uintptr) (r uint8) {
	var v1 uintptr
	_ = v1
	if !((*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer1)).Feof})))(tls, lexer1) != 0) && (*TSLexer)(unsafe.Pointer(lexer1)).Flookahead == int32('-') {
		v1 = lexer1
		(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(v1)).Fadvance})))(tls, v1, libc.BoolUint8(false1 != 0))
	} else {
		return libc.BoolUint8(false1 != 0)
	}
	if !((*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer1)).Feof})))(tls, lexer1) != 0) && (*TSLexer)(unsafe.Pointer(lexer1)).Flookahead == int32('-') {
		v1 = lexer1
		(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(v1)).Fadvance})))(tls, v1, libc.BoolUint8(false1 != 0))
	} else {
		return libc.BoolUint8(false1 != 0)
	}
	for !((*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer1)).Feof})))(tls, lexer1) != 0) {
		if (*TSLexer)(unsafe.Pointer(lexer1)).Flookahead == int32('-') {
			v1 = lexer1
			(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(v1)).Fadvance})))(tls, v1, libc.BoolUint8(false1 != 0))
			if (*TSLexer)(unsafe.Pointer(lexer1)).Flookahead == int32('-') {
				v1 = lexer1
				(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(v1)).Fadvance})))(tls, v1, libc.BoolUint8(false1 != 0))
				break
			}
		} else {
			v1 = lexer1
			(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(v1)).Fadvance})))(tls, v1, libc.BoolUint8(false1 != 0))
		}
	}
	if (*TSLexer)(unsafe.Pointer(lexer1)).Flookahead == int32('>') {
		v1 = lexer1
		(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(v1)).Fadvance})))(tls, v1, libc.BoolUint8(false1 != 0))
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer1)).Fmark_end})))(tls, lexer1)
		(*TSLexer)(unsafe.Pointer(lexer1)).Fresult_symbol = uint16(COMMENT)
		return libc.BoolUint8(true1 != 0)
	}
	return libc.BoolUint8(false1 != 0)
}

// C documentation
//
//	/// Check if the lexer is in error recovery mode
func in_error_recovery(tls *libc.TLS, valid_symbols uintptr) (r uint8) {
	return libc.BoolUint8(*(*uint8)(unsafe.Pointer(valid_symbols + uintptr(PI_TARGET))) != 0 && *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(PI_CONTENT))) != 0 && *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(COMMENT))) != 0)
}

func tree_sitter_dtd_external_scanner_scan(tls *libc.TLS, payload uintptr, lexer1 uintptr, valid_symbols uintptr) (r uint8) {
	var v1 uintptr
	_ = v1
	if in_error_recovery(tls, valid_symbols) != 0 {
		return libc.BoolUint8(false1 != 0)
	}
	if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(PI_TARGET))) != 0 {
		return scan_pi_target(tls, lexer1, valid_symbols)
	}
	if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(PI_CONTENT))) != 0 {
		return scan_pi_content(tls, lexer1)
	}
	if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(COMMENT))) != 0 {
		if !((*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer1)).Feof})))(tls, lexer1) != 0) && (*TSLexer)(unsafe.Pointer(lexer1)).Flookahead == int32('<') {
			v1 = lexer1
			(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(v1)).Fadvance})))(tls, v1, libc.BoolUint8(false1 != 0))
		} else {
			return libc.BoolUint8(false1 != 0)
		}
		if !((*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer1)).Feof})))(tls, lexer1) != 0) && (*TSLexer)(unsafe.Pointer(lexer1)).Flookahead == int32('!') {
			v1 = lexer1
			(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(v1)).Fadvance})))(tls, v1, libc.BoolUint8(false1 != 0))
		} else {
			return libc.BoolUint8(false1 != 0)
		}
		return scan_comment(tls, lexer1)
	}
	return libc.BoolUint8(false1 != 0)
}

func tree_sitter_dtd_external_scanner_create(tls *libc.TLS) (r uintptr) {
	return libc.UintptrFromInt32(0)
}

func tree_sitter_dtd_external_scanner_destroy(tls *libc.TLS, payload uintptr) {
}

func tree_sitter_dtd_external_scanner_serialize(tls *libc.TLS, payload uintptr, buffer uintptr) (r uint32) {
	return uint32(0)
}

func tree_sitter_dtd_external_scanner_deserialize(tls *libc.TLS, payload uintptr, buffer uintptr, length uint32) {
}

/* Automatically @generated by tree-sitter v0.25.4 (683a5044a8f496d8e8b57c06ac4f1a722c62205f) */

type ts_symbol_identifiers = int32

const sym_Name = 1
const anon_sym_LT_QMARK = 2
const anon_sym_xml = 3
const anon_sym_QMARK_GT = 4
const anon_sym_LT_BANG_LBRACK = 5
const anon_sym_IGNORE = 6
const anon_sym_INCLUDE = 7
const anon_sym_LBRACK = 8
const anon_sym_RBRACK_RBRACK_GT = 9
const anon_sym_LT_BANG = 10
const anon_sym_ELEMENT = 11
const anon_sym_GT = 12
const anon_sym_EMPTY = 13
const anon_sym_ANY = 14
const anon_sym_LPAREN = 15
const anon_sym_POUNDPCDATA = 16
const anon_sym_PIPE = 17
const anon_sym_RPAREN = 18
const anon_sym_STAR = 19
const anon_sym_QMARK = 20
const anon_sym_PLUS = 21
const anon_sym_COMMA = 22
const anon_sym_ATTLIST = 23
const sym_StringType = 24
const sym_TokenizedType = 25
const anon_sym_NOTATION = 26
const anon_sym_POUNDREQUIRED = 27
const anon_sym_POUNDIMPLIED = 28
const anon_sym_POUNDFIXED = 29
const anon_sym_ENTITY = 30
const anon_sym_PERCENT = 31
const anon_sym_DQUOTE = 32
const aux_sym_EntityValue_token1 = 33
const anon_sym_SQUOTE = 34
const aux_sym_EntityValue_token2 = 35
const anon_sym_NDATA = 36
const anon_sym_SEMI = 37
const sym__S = 38
const sym_Nmtoken = 39
const anon_sym_AMP = 40
const anon_sym_AMP_POUND = 41
const aux_sym_CharRef_token1 = 42
const anon_sym_AMP_POUNDx = 43
const aux_sym_CharRef_token2 = 44
const aux_sym_AttValue_token1 = 45
const aux_sym_AttValue_token2 = 46
const anon_sym_SYSTEM = 47
const anon_sym_PUBLIC = 48
const aux_sym_SystemLiteral_token1 = 49
const aux_sym_SystemLiteral_token2 = 50
const aux_sym_PubidLiteral_token1 = 51
const aux_sym_PubidLiteral_token2 = 52
const anon_sym_version = 53
const sym_VersionNum = 54
const anon_sym_encoding = 55
const sym_EncName = 56
const anon_sym_EQ = 57
const sym_PITarget = 58
const sym__pi_content = 59
const sym_Comment = 60
const sym_extSubset = 61
const sym_TextDecl = 62
const sym__extSubsetDecl = 63
const sym_conditionalSect = 64
const sym__markupdecl = 65
const sym__DeclSep = 66
const sym_elementdecl = 67
const sym_contentspec = 68
const sym_Mixed = 69
const sym_children = 70
const sym__cp = 71
const sym__choice = 72
const sym_AttlistDecl = 73
const sym_AttDef = 74
const sym__AttType = 75
const sym__EnumeratedType = 76
const sym_NotationType = 77
const sym_Enumeration = 78
const sym_DefaultDecl = 79
const sym__EntityDecl = 80
const sym_GEDecl = 81
const sym_PEDecl = 82
const sym_EntityValue = 83
const sym_NDataDecl = 84
const sym_NotationDecl = 85
const sym_PEReference = 86
const sym__Reference = 87
const sym_EntityRef = 88
const sym_CharRef = 89
const sym_AttValue = 90
const sym_ExternalID = 91
const sym_PublicID = 92
const sym_SystemLiteral = 93
const sym_PubidLiteral = 94
const sym__VersionInfo = 95
const sym__EncodingDecl = 96
const sym_PI = 97
const sym__Eq = 98
const aux_sym_extSubset_repeat1 = 99
const aux_sym_Mixed_repeat1 = 100
const aux_sym_Mixed_repeat2 = 101
const aux_sym__choice_repeat1 = 102
const aux_sym__choice_repeat2 = 103
const aux_sym_AttlistDecl_repeat1 = 104
const aux_sym_NotationType_repeat1 = 105
const aux_sym_Enumeration_repeat1 = 106
const aux_sym_EntityValue_repeat1 = 107
const aux_sym_EntityValue_repeat2 = 108
const aux_sym_AttValue_repeat1 = 109
const aux_sym_AttValue_repeat2 = 110

var ts_symbol_names = [111]uintptr{
	0:   __ccgo_ts,
	1:   __ccgo_ts + 4,
	2:   __ccgo_ts + 9,
	3:   __ccgo_ts + 12,
	4:   __ccgo_ts + 16,
	5:   __ccgo_ts + 19,
	6:   __ccgo_ts + 23,
	7:   __ccgo_ts + 30,
	8:   __ccgo_ts + 38,
	9:   __ccgo_ts + 40,
	10:  __ccgo_ts + 44,
	11:  __ccgo_ts + 47,
	12:  __ccgo_ts + 55,
	13:  __ccgo_ts + 57,
	14:  __ccgo_ts + 63,
	15:  __ccgo_ts + 67,
	16:  __ccgo_ts + 69,
	17:  __ccgo_ts + 77,
	18:  __ccgo_ts + 79,
	19:  __ccgo_ts + 81,
	20:  __ccgo_ts + 83,
	21:  __ccgo_ts + 85,
	22:  __ccgo_ts + 87,
	23:  __ccgo_ts + 89,
	24:  __ccgo_ts + 97,
	25:  __ccgo_ts + 108,
	26:  __ccgo_ts + 122,
	27:  __ccgo_ts + 131,
	28:  __ccgo_ts + 141,
	29:  __ccgo_ts + 150,
	30:  __ccgo_ts + 157,
	31:  __ccgo_ts + 164,
	32:  __ccgo_ts + 166,
	33:  __ccgo_ts + 168,
	34:  __ccgo_ts + 187,
	35:  __ccgo_ts + 189,
	36:  __ccgo_ts + 208,
	37:  __ccgo_ts + 214,
	38:  __ccgo_ts + 216,
	39:  __ccgo_ts + 219,
	40:  __ccgo_ts + 227,
	41:  __ccgo_ts + 229,
	42:  __ccgo_ts + 232,
	43:  __ccgo_ts + 247,
	44:  __ccgo_ts + 251,
	45:  __ccgo_ts + 266,
	46:  __ccgo_ts + 282,
	47:  __ccgo_ts + 298,
	48:  __ccgo_ts + 305,
	49:  __ccgo_ts + 312,
	50:  __ccgo_ts + 312,
	51:  __ccgo_ts + 316,
	52:  __ccgo_ts + 336,
	53:  __ccgo_ts + 356,
	54:  __ccgo_ts + 364,
	55:  __ccgo_ts + 375,
	56:  __ccgo_ts + 384,
	57:  __ccgo_ts + 392,
	58:  __ccgo_ts + 394,
	59:  __ccgo_ts + 403,
	60:  __ccgo_ts + 415,
	61:  __ccgo_ts + 423,
	62:  __ccgo_ts + 433,
	63:  __ccgo_ts + 442,
	64:  __ccgo_ts + 457,
	65:  __ccgo_ts + 473,
	66:  __ccgo_ts + 485,
	67:  __ccgo_ts + 494,
	68:  __ccgo_ts + 506,
	69:  __ccgo_ts + 518,
	70:  __ccgo_ts + 524,
	71:  __ccgo_ts + 533,
	72:  __ccgo_ts + 537,
	73:  __ccgo_ts + 545,
	74:  __ccgo_ts + 557,
	75:  __ccgo_ts + 564,
	76:  __ccgo_ts + 573,
	77:  __ccgo_ts + 589,
	78:  __ccgo_ts + 602,
	79:  __ccgo_ts + 614,
	80:  __ccgo_ts + 626,
	81:  __ccgo_ts + 638,
	82:  __ccgo_ts + 645,
	83:  __ccgo_ts + 652,
	84:  __ccgo_ts + 664,
	85:  __ccgo_ts + 674,
	86:  __ccgo_ts + 687,
	87:  __ccgo_ts + 699,
	88:  __ccgo_ts + 710,
	89:  __ccgo_ts + 720,
	90:  __ccgo_ts + 728,
	91:  __ccgo_ts + 737,
	92:  __ccgo_ts + 748,
	93:  __ccgo_ts + 757,
	94:  __ccgo_ts + 771,
	95:  __ccgo_ts + 784,
	96:  __ccgo_ts + 797,
	97:  __ccgo_ts + 811,
	98:  __ccgo_ts + 814,
	99:  __ccgo_ts + 818,
	100: __ccgo_ts + 836,
	101: __ccgo_ts + 850,
	102: __ccgo_ts + 864,
	103: __ccgo_ts + 880,
	104: __ccgo_ts + 896,
	105: __ccgo_ts + 916,
	106: __ccgo_ts + 937,
	107: __ccgo_ts + 957,
	108: __ccgo_ts + 977,
	109: __ccgo_ts + 997,
	110: __ccgo_ts + 1014,
}

var ts_symbol_map = [111]TSSymbol{
	1:   uint16(sym_Name),
	2:   uint16(anon_sym_LT_QMARK),
	3:   uint16(anon_sym_xml),
	4:   uint16(anon_sym_QMARK_GT),
	5:   uint16(anon_sym_LT_BANG_LBRACK),
	6:   uint16(anon_sym_IGNORE),
	7:   uint16(anon_sym_INCLUDE),
	8:   uint16(anon_sym_LBRACK),
	9:   uint16(anon_sym_RBRACK_RBRACK_GT),
	10:  uint16(anon_sym_LT_BANG),
	11:  uint16(anon_sym_ELEMENT),
	12:  uint16(anon_sym_GT),
	13:  uint16(anon_sym_EMPTY),
	14:  uint16(anon_sym_ANY),
	15:  uint16(anon_sym_LPAREN),
	16:  uint16(anon_sym_POUNDPCDATA),
	17:  uint16(anon_sym_PIPE),
	18:  uint16(anon_sym_RPAREN),
	19:  uint16(anon_sym_STAR),
	20:  uint16(anon_sym_QMARK),
	21:  uint16(anon_sym_PLUS),
	22:  uint16(anon_sym_COMMA),
	23:  uint16(anon_sym_ATTLIST),
	24:  uint16(sym_StringType),
	25:  uint16(sym_TokenizedType),
	26:  uint16(anon_sym_NOTATION),
	27:  uint16(anon_sym_POUNDREQUIRED),
	28:  uint16(anon_sym_POUNDIMPLIED),
	29:  uint16(anon_sym_POUNDFIXED),
	30:  uint16(anon_sym_ENTITY),
	31:  uint16(anon_sym_PERCENT),
	32:  uint16(anon_sym_DQUOTE),
	33:  uint16(aux_sym_EntityValue_token1),
	34:  uint16(anon_sym_SQUOTE),
	35:  uint16(aux_sym_EntityValue_token2),
	36:  uint16(anon_sym_NDATA),
	37:  uint16(anon_sym_SEMI),
	38:  uint16(sym__S),
	39:  uint16(sym_Nmtoken),
	40:  uint16(anon_sym_AMP),
	41:  uint16(anon_sym_AMP_POUND),
	42:  uint16(aux_sym_CharRef_token1),
	43:  uint16(anon_sym_AMP_POUNDx),
	44:  uint16(aux_sym_CharRef_token2),
	45:  uint16(aux_sym_AttValue_token1),
	46:  uint16(aux_sym_AttValue_token2),
	47:  uint16(anon_sym_SYSTEM),
	48:  uint16(anon_sym_PUBLIC),
	49:  uint16(aux_sym_SystemLiteral_token1),
	50:  uint16(aux_sym_SystemLiteral_token1),
	51:  uint16(aux_sym_PubidLiteral_token1),
	52:  uint16(aux_sym_PubidLiteral_token2),
	53:  uint16(anon_sym_version),
	54:  uint16(sym_VersionNum),
	55:  uint16(anon_sym_encoding),
	56:  uint16(sym_EncName),
	57:  uint16(anon_sym_EQ),
	58:  uint16(sym_PITarget),
	59:  uint16(sym__pi_content),
	60:  uint16(sym_Comment),
	61:  uint16(sym_extSubset),
	62:  uint16(sym_TextDecl),
	63:  uint16(sym__extSubsetDecl),
	64:  uint16(sym_conditionalSect),
	65:  uint16(sym__markupdecl),
	66:  uint16(sym__DeclSep),
	67:  uint16(sym_elementdecl),
	68:  uint16(sym_contentspec),
	69:  uint16(sym_Mixed),
	70:  uint16(sym_children),
	71:  uint16(sym__cp),
	72:  uint16(sym__choice),
	73:  uint16(sym_AttlistDecl),
	74:  uint16(sym_AttDef),
	75:  uint16(sym__AttType),
	76:  uint16(sym__EnumeratedType),
	77:  uint16(sym_NotationType),
	78:  uint16(sym_Enumeration),
	79:  uint16(sym_DefaultDecl),
	80:  uint16(sym__EntityDecl),
	81:  uint16(sym_GEDecl),
	82:  uint16(sym_PEDecl),
	83:  uint16(sym_EntityValue),
	84:  uint16(sym_NDataDecl),
	85:  uint16(sym_NotationDecl),
	86:  uint16(sym_PEReference),
	87:  uint16(sym__Reference),
	88:  uint16(sym_EntityRef),
	89:  uint16(sym_CharRef),
	90:  uint16(sym_AttValue),
	91:  uint16(sym_ExternalID),
	92:  uint16(sym_PublicID),
	93:  uint16(sym_SystemLiteral),
	94:  uint16(sym_PubidLiteral),
	95:  uint16(sym__VersionInfo),
	96:  uint16(sym__EncodingDecl),
	97:  uint16(sym_PI),
	98:  uint16(sym__Eq),
	99:  uint16(aux_sym_extSubset_repeat1),
	100: uint16(aux_sym_Mixed_repeat1),
	101: uint16(aux_sym_Mixed_repeat2),
	102: uint16(aux_sym__choice_repeat1),
	103: uint16(aux_sym__choice_repeat2),
	104: uint16(aux_sym_AttlistDecl_repeat1),
	105: uint16(aux_sym_NotationType_repeat1),
	106: uint16(aux_sym_Enumeration_repeat1),
	107: uint16(aux_sym_EntityValue_repeat1),
	108: uint16(aux_sym_EntityValue_repeat2),
	109: uint16(aux_sym_AttValue_repeat1),
	110: uint16(aux_sym_AttValue_repeat2),
}

var ts_symbol_metadata = [111]TSSymbolMetadata{
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
	},
	22: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	23: {
		Fvisible: libc.BoolUint8(true1 != 0),
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
	33: {},
	34: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	35: {},
	36: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	37: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	38: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	39: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	40: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	41: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	42: {},
	43: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	44: {},
	45: {},
	46: {},
	47: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	48: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	49: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	50: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	51: {},
	52: {},
	53: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	54: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	55: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	56: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	57: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	58: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	59: {
		Fnamed: libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	64: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	65: {
		Fnamed:     libc.BoolUint8(true1 != 0),
		Fsupertype: libc.BoolUint8(true1 != 0),
	},
	66: {
		Fnamed: libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	72: {
		Fnamed: libc.BoolUint8(true1 != 0),
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
		Fnamed:     libc.BoolUint8(true1 != 0),
		Fsupertype: libc.BoolUint8(true1 != 0),
	},
	76: {
		Fnamed:     libc.BoolUint8(true1 != 0),
		Fsupertype: libc.BoolUint8(true1 != 0),
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
		Fnamed:     libc.BoolUint8(true1 != 0),
		Fsupertype: libc.BoolUint8(true1 != 0),
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
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	87: {
		Fnamed:     libc.BoolUint8(true1 != 0),
		Fsupertype: libc.BoolUint8(true1 != 0),
	},
	88: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	89: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	90: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	96: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	97: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	98: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	99:  {},
	100: {},
	101: {},
	102: {},
	103: {},
	104: {},
	105: {},
	106: {},
	107: {},
	108: {},
	109: {},
	110: {},
}

type ts_field_identifiers = int32

const field_content = 1

var ts_field_names = [2]uintptr{
	0: libc.UintptrFromInt32(0),
	1: __ccgo_ts + 1031,
}

var ts_field_map_slices = [2]TSMapSlice{
	1: {
		Flength: uint16(1),
	},
}

var ts_field_map_entries = [1]TSFieldMapEntry{
	0: {
		Ffield_id:    uint16(field_content),
		Fchild_index: uint8(1),
	},
}

var ts_alias_sequences = [2][10]TSSymbol{}

var ts_non_terminal_alias_map = [1]uint16_t{}

var ts_primary_state_ids = [334]TSStateId{
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
	42:  uint16(12),
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
	71:  uint16(12),
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
	88:  uint16(75),
	89:  uint16(76),
	90:  uint16(12),
	91:  uint16(70),
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
	110: uint16(70),
	111: uint16(75),
	112: uint16(76),
	113: uint16(113),
	114: uint16(70),
	115: uint16(75),
	116: uint16(76),
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
	188: uint16(188),
	189: uint16(189),
	190: uint16(190),
	191: uint16(191),
	192: uint16(192),
	193: uint16(193),
	194: uint16(194),
	195: uint16(195),
	196: uint16(196),
	197: uint16(197),
	198: uint16(198),
	199: uint16(199),
	200: uint16(200),
	201: uint16(201),
	202: uint16(202),
	203: uint16(203),
	204: uint16(204),
	205: uint16(205),
	206: uint16(206),
	207: uint16(207),
	208: uint16(208),
	209: uint16(209),
	210: uint16(210),
	211: uint16(211),
	212: uint16(212),
	213: uint16(213),
	214: uint16(214),
	215: uint16(215),
	216: uint16(216),
	217: uint16(217),
	218: uint16(218),
	219: uint16(219),
	220: uint16(220),
	221: uint16(221),
	222: uint16(222),
	223: uint16(223),
	224: uint16(224),
	225: uint16(225),
	226: uint16(226),
	227: uint16(227),
	228: uint16(228),
	229: uint16(229),
	230: uint16(230),
	231: uint16(231),
	232: uint16(232),
	233: uint16(233),
	234: uint16(234),
	235: uint16(235),
	236: uint16(236),
	237: uint16(237),
	238: uint16(238),
	239: uint16(239),
	240: uint16(240),
	241: uint16(241),
	242: uint16(242),
	243: uint16(243),
	244: uint16(244),
	245: uint16(245),
	246: uint16(246),
	247: uint16(247),
	248: uint16(248),
	249: uint16(249),
	250: uint16(250),
	251: uint16(251),
	252: uint16(252),
	253: uint16(253),
	254: uint16(254),
	255: uint16(255),
	256: uint16(256),
	257: uint16(257),
	258: uint16(258),
	259: uint16(259),
	260: uint16(260),
	261: uint16(261),
	262: uint16(262),
	263: uint16(263),
	264: uint16(264),
	265: uint16(265),
	266: uint16(266),
	267: uint16(267),
	268: uint16(268),
	269: uint16(269),
	270: uint16(270),
	271: uint16(271),
	272: uint16(272),
	273: uint16(273),
	274: uint16(274),
	275: uint16(275),
	276: uint16(276),
	277: uint16(277),
	278: uint16(278),
	279: uint16(279),
	280: uint16(280),
	281: uint16(281),
	282: uint16(282),
	283: uint16(283),
	284: uint16(284),
	285: uint16(285),
	286: uint16(286),
	287: uint16(287),
	288: uint16(244),
	289: uint16(289),
	290: uint16(290),
	291: uint16(291),
	292: uint16(292),
	293: uint16(293),
	294: uint16(294),
	295: uint16(295),
	296: uint16(296),
	297: uint16(297),
	298: uint16(298),
	299: uint16(299),
	300: uint16(300),
	301: uint16(301),
	302: uint16(302),
	303: uint16(303),
	304: uint16(304),
	305: uint16(305),
	306: uint16(306),
	307: uint16(307),
	308: uint16(308),
	309: uint16(309),
	310: uint16(310),
	311: uint16(311),
	312: uint16(280),
	313: uint16(287),
	314: uint16(244),
	315: uint16(280),
	316: uint16(287),
	317: uint16(244),
	318: uint16(280),
	319: uint16(287),
	320: uint16(320),
	321: uint16(279),
	322: uint16(306),
	323: uint16(311),
	324: uint16(246),
	325: uint16(279),
	326: uint16(306),
	327: uint16(311),
	328: uint16(246),
	329: uint16(279),
	330: uint16(306),
	331: uint16(311),
	332: uint16(246),
	333: uint16(333),
}

var aux_sym_PubidLiteral_token1_character_set_1 = [9]TSCharacterRange{
	0: {
		Fstart: int32('\n'),
		Fend:   int32('\n'),
	},
	1: {
		Fstart: int32('\r'),
		Fend:   int32('\r'),
	},
	2: {
		Fstart: int32(' '),
		Fend:   int32('!'),
	},
	3: {
		Fstart: int32('#'),
		Fend:   int32('%'),
	},
	4: {
		Fstart: int32('\''),
		Fend:   int32(';'),
	},
	5: {
		Fstart: int32('='),
		Fend:   int32('='),
	},
	6: {
		Fstart: int32('?'),
		Fend:   int32('Z'),
	},
	7: {
		Fstart: int32('_'),
		Fend:   int32('_'),
	},
	8: {
		Fstart: int32('a'),
		Fend:   int32('z'),
	},
}

var aux_sym_PubidLiteral_token2_character_set_1 = [9]TSCharacterRange{
	0: {
		Fstart: int32('\n'),
		Fend:   int32('\n'),
	},
	1: {
		Fstart: int32('\r'),
		Fend:   int32('\r'),
	},
	2: {
		Fstart: int32(' '),
		Fend:   int32('!'),
	},
	3: {
		Fstart: int32('#'),
		Fend:   int32('%'),
	},
	4: {
		Fstart: int32('('),
		Fend:   int32(';'),
	},
	5: {
		Fstart: int32('='),
		Fend:   int32('='),
	},
	6: {
		Fstart: int32('?'),
		Fend:   int32('Z'),
	},
	7: {
		Fstart: int32('_'),
		Fend:   int32('_'),
	},
	8: {
		Fstart: int32('a'),
		Fend:   int32('z'),
	},
}

func ts_lex(tls *libc.TLS, lexer uintptr, state TSStateId) (r uint8) {
	var eof, result, skip, v6 uint8
	var half_size, i, i1, i2, index, mid_index, size uint32_t
	var lookahead1, v5 int32_t
	var range_token, range_token1, v4 uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = eof, half_size, i, i1, i2, index, lookahead1, mid_index, range_token, range_token1, result, size, skip, v4, v5, v6
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
			state = uint16(40)
			goto next_state
		}
		i = uint32(0)
		for {
			if !(uint64(i) < libc.Uint64FromInt64(124)/libc.Uint64FromInt64(2)) {
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
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(77)
			goto next_state
		}
		if int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(74)
			goto next_state
		}
		if int32('G') <= lookahead1 && lookahead1 <= int32('Z') || int32('g') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(75)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(67)
			goto next_state
		}
		return result
	case int32(1):
		if lookahead1 == int32('!') {
			state = uint16(46)
			goto next_state
		}
		if lookahead1 == int32('?') {
			state = uint16(41)
			goto next_state
		}
		return result
	case int32(2):
		if lookahead1 == int32('"') {
			state = uint16(66)
			goto next_state
		}
		if lookahead1 == int32('%') {
			state = uint16(65)
			goto next_state
		}
		if lookahead1 == int32('&') {
			state = uint16(120)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32('<') {
			state = uint16(67)
			goto next_state
		}
		return result
	case int32(3):
		if lookahead1 == int32('"') {
			state = uint16(66)
			goto next_state
		}
		if lookahead1 == int32('&') {
			state = uint16(120)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32('<') {
			state = uint16(125)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead1 == int32('%') {
			state = uint16(65)
			goto next_state
		}
		if lookahead1 == int32('&') {
			state = uint16(120)
			goto next_state
		}
		if lookahead1 == int32('\'') {
			state = uint16(80)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32('<') {
			state = uint16(81)
			goto next_state
		}
		return result
	case int32(5):
		i1 = uint32(0)
		for {
			if !(uint64(i1) < libc.Uint64FromInt64(40)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token1[i1]) == lookahead1 {
				state = map_token1[i1+uint32(1)]
				goto next_state
			}
			goto _2
		_2:
			;
			i1 = i1 + uint32(2)
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(122)
			goto next_state
		}
		if int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(115)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead1 == int32('&') {
			state = uint16(120)
			goto next_state
		}
		if lookahead1 == int32('\'') {
			state = uint16(80)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32('<') {
			state = uint16(126)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead1 == int32('.') {
			state = uint16(36)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead1 == int32('>') {
			state = uint16(45)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead1 == int32('>') {
			state = uint16(42)
			goto next_state
		}
		return result
	case int32(10):
		if lookahead1 == int32('?') {
			state = uint16(9)
			goto next_state
		}
		if int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(115)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead1 == int32('A') {
			state = uint16(31)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead1 == int32('A') {
			state = uint16(49)
			goto next_state
		}
		return result
	case int32(13):
		if lookahead1 == int32('C') {
			state = uint16(14)
			goto next_state
		}
		return result
	case int32(14):
		if lookahead1 == int32('D') {
			state = uint16(11)
			goto next_state
		}
		return result
	case int32(15):
		if lookahead1 == int32('D') {
			state = uint16(64)
			goto next_state
		}
		return result
	case int32(16):
		if lookahead1 == int32('D') {
			state = uint16(63)
			goto next_state
		}
		return result
	case int32(17):
		if lookahead1 == int32('D') {
			state = uint16(62)
			goto next_state
		}
		return result
	case int32(18):
		if lookahead1 == int32('E') {
			state = uint16(29)
			goto next_state
		}
		return result
	case int32(19):
		if lookahead1 == int32('E') {
			state = uint16(15)
			goto next_state
		}
		return result
	case int32(20):
		if lookahead1 == int32('E') {
			state = uint16(16)
			goto next_state
		}
		return result
	case int32(21):
		if lookahead1 == int32('E') {
			state = uint16(17)
			goto next_state
		}
		return result
	case int32(22):
		if lookahead1 == int32('F') {
			state = uint16(23)
			goto next_state
		}
		if lookahead1 == int32('I') {
			state = uint16(27)
			goto next_state
		}
		if lookahead1 == int32('P') {
			state = uint16(13)
			goto next_state
		}
		if lookahead1 == int32('R') {
			state = uint16(18)
			goto next_state
		}
		return result
	case int32(23):
		if lookahead1 == int32('I') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(24):
		if lookahead1 == int32('I') {
			state = uint16(30)
			goto next_state
		}
		return result
	case int32(25):
		if lookahead1 == int32('I') {
			state = uint16(20)
			goto next_state
		}
		return result
	case int32(26):
		if lookahead1 == int32('L') {
			state = uint16(25)
			goto next_state
		}
		return result
	case int32(27):
		if lookahead1 == int32('M') {
			state = uint16(28)
			goto next_state
		}
		return result
	case int32(28):
		if lookahead1 == int32('P') {
			state = uint16(26)
			goto next_state
		}
		return result
	case int32(29):
		if lookahead1 == int32('Q') {
			state = uint16(32)
			goto next_state
		}
		return result
	case int32(30):
		if lookahead1 == int32('R') {
			state = uint16(21)
			goto next_state
		}
		return result
	case int32(31):
		if lookahead1 == int32('T') {
			state = uint16(12)
			goto next_state
		}
		return result
	case int32(32):
		if lookahead1 == int32('U') {
			state = uint16(24)
			goto next_state
		}
		return result
	case int32(33):
		if lookahead1 == int32('X') {
			state = uint16(19)
			goto next_state
		}
		return result
	case int32(34):
		if lookahead1 == int32(']') {
			state = uint16(8)
			goto next_state
		}
		return result
	case int32(35):
		if lookahead1 == int32('\t') || lookahead1 == int32('\n') || lookahead1 == int32('\r') || lookahead1 == int32(' ') {
			state = uint16(83)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32(':') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') || lookahead1 == int32(0xb7) {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(36):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(131)
			goto next_state
		}
		return result
	case int32(37):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(124)
			goto next_state
		}
		return result
	case int32(38):
		if int32('A') <= lookahead1 && lookahead1 <= int32('Z') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(132)
			goto next_state
		}
		return result
	case int32(39):
		if eof != 0 {
			state = uint16(40)
			goto next_state
		}
		i2 = uint32(0)
		for {
			if !(uint64(i2) < libc.Uint64FromInt64(88)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token2[i2]) == lookahead1 {
				state = map_token2[i2+uint32(1)]
				goto next_state
			}
			goto _3
		_3:
			;
			i2 = i2 + uint32(2)
		}
		if int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(115)
			goto next_state
		}
		return result
	case int32(40):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(ts_builtin_sym_end)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(41):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT_QMARK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(42):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_QMARK_GT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(43):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT_BANG_LBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(44):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(45):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACK_RBRACK_GT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(46):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT_BANG)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('[') {
			state = uint16(43)
			goto next_state
		}
		return result
	case int32(47):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_GT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(48):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(49):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUNDPCDATA)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(50):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PIPE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(51):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(52):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STAR)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(53):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_QMARK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(54):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PLUS)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(55):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COMMA)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(56):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_TokenizedType)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('R') {
			state = uint16(85)
			goto next_state
		}
		if lookahead1 == int32(':') || lookahead1 == int32(0xb7) {
			state = uint16(115)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(114)
			goto next_state
		}
		return result
	case int32(57):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_TokenizedType)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('R') {
			state = uint16(88)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32(':') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') || lookahead1 == int32(0xb7) {
			state = uint16(115)
			goto next_state
		}
		return result
	case int32(58):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_TokenizedType)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('S') {
			state = uint16(60)
			goto next_state
		}
		if lookahead1 == int32(':') || lookahead1 == int32(0xb7) {
			state = uint16(115)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(114)
			goto next_state
		}
		return result
	case int32(59):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_TokenizedType)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('S') {
			state = uint16(61)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32(':') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') || lookahead1 == int32(0xb7) {
			state = uint16(115)
			goto next_state
		}
		return result
	case int32(60):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_TokenizedType)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(':') || lookahead1 == int32(0xb7) {
			state = uint16(115)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(114)
			goto next_state
		}
		return result
	case int32(61):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_TokenizedType)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32(':') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') || lookahead1 == int32(0xb7) {
			state = uint16(115)
			goto next_state
		}
		return result
	case int32(62):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUNDREQUIRED)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(63):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUNDIMPLIED)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(64):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUNDFIXED)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(65):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PERCENT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(66):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(67):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_EntityValue_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(68):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_EntityValue_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('.') {
			state = uint16(116)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(117)
			goto next_state
		}
		if int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(118)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32(':') || int32('G') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('g') <= lookahead1 && lookahead1 <= int32('z') || lookahead1 == int32(0xb7) {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(69):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_EntityValue_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('D') {
			state = uint16(56)
			goto next_state
		}
		if lookahead1 == int32(':') || lookahead1 == int32(0xb7) {
			state = uint16(115)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(114)
			goto next_state
		}
		return result
	case int32(70):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_EntityValue_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('F') {
			state = uint16(23)
			goto next_state
		}
		if lookahead1 == int32('I') {
			state = uint16(27)
			goto next_state
		}
		if lookahead1 == int32('P') {
			state = uint16(13)
			goto next_state
		}
		if lookahead1 == int32('R') {
			state = uint16(18)
			goto next_state
		}
		return result
	case int32(71):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_EntityValue_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('M') {
			state = uint16(108)
			goto next_state
		}
		if lookahead1 == int32(':') || lookahead1 == int32(0xb7) {
			state = uint16(115)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(114)
			goto next_state
		}
		return result
	case int32(72):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_EntityValue_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('N') {
			state = uint16(107)
			goto next_state
		}
		if lookahead1 == int32(':') || lookahead1 == int32(0xb7) {
			state = uint16(115)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(113)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('G') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('g') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(114)
			goto next_state
		}
		return result
	case int32(73):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_EntityValue_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(']') {
			state = uint16(8)
			goto next_state
		}
		return result
	case int32(74):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_EntityValue_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(':') || lookahead1 == int32(0xb7) {
			state = uint16(115)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(113)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('G') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('g') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(114)
			goto next_state
		}
		return result
	case int32(75):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_EntityValue_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(':') || lookahead1 == int32(0xb7) {
			state = uint16(115)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(114)
			goto next_state
		}
		return result
	case int32(76):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_EntityValue_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\t') || lookahead1 == int32('\n') || lookahead1 == int32('\r') || lookahead1 == int32(' ') {
			state = uint16(83)
			goto next_state
		}
		return result
	case int32(77):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_EntityValue_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(117)
			goto next_state
		}
		if int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(118)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || lookahead1 == int32(':') || int32('G') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('g') <= lookahead1 && lookahead1 <= int32('z') || lookahead1 == int32(0xb7) {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(78):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_EntityValue_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32(':') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') || lookahead1 == int32(0xb7) {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(79):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_EntityValue_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32(':') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') || lookahead1 == int32(0xb7) {
			state = uint16(115)
			goto next_state
		}
		return result
	case int32(80):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(81):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_EntityValue_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(82):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SEMI)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(83):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__S)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\t') || lookahead1 == int32('\n') || lookahead1 == int32('\r') || lookahead1 == int32(' ') {
			state = uint16(83)
			goto next_state
		}
		return result
	case int32(84):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('D') {
			state = uint16(57)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32(':') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') || lookahead1 == int32(0xb7) {
			state = uint16(115)
			goto next_state
		}
		return result
	case int32(85):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('E') {
			state = uint16(91)
			goto next_state
		}
		if lookahead1 == int32(':') || lookahead1 == int32(0xb7) {
			state = uint16(115)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(114)
			goto next_state
		}
		return result
	case int32(86):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('E') {
			state = uint16(100)
			goto next_state
		}
		if lookahead1 == int32(':') || lookahead1 == int32(0xb7) {
			state = uint16(115)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(114)
			goto next_state
		}
		return result
	case int32(87):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('E') {
			state = uint16(105)
			goto next_state
		}
		if lookahead1 == int32(':') || lookahead1 == int32(0xb7) {
			state = uint16(115)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(114)
			goto next_state
		}
		return result
	case int32(88):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('E') {
			state = uint16(92)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32(':') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') || lookahead1 == int32(0xb7) {
			state = uint16(115)
			goto next_state
		}
		return result
	case int32(89):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('E') {
			state = uint16(106)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32(':') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') || lookahead1 == int32(0xb7) {
			state = uint16(115)
			goto next_state
		}
		return result
	case int32(90):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('E') {
			state = uint16(102)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32(':') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') || lookahead1 == int32(0xb7) {
			state = uint16(115)
			goto next_state
		}
		return result
	case int32(91):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('F') {
			state = uint16(58)
			goto next_state
		}
		if lookahead1 == int32(':') || lookahead1 == int32(0xb7) {
			state = uint16(115)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(114)
			goto next_state
		}
		return result
	case int32(92):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('F') {
			state = uint16(59)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32(':') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') || lookahead1 == int32(0xb7) {
			state = uint16(115)
			goto next_state
		}
		return result
	case int32(93):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('I') {
			state = uint16(109)
			goto next_state
		}
		if lookahead1 == int32(':') || lookahead1 == int32(0xb7) {
			state = uint16(115)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(114)
			goto next_state
		}
		return result
	case int32(94):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('I') {
			state = uint16(87)
			goto next_state
		}
		if lookahead1 == int32('Y') {
			state = uint16(60)
			goto next_state
		}
		if lookahead1 == int32(':') || lookahead1 == int32(0xb7) {
			state = uint16(115)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(114)
			goto next_state
		}
		return result
	case int32(95):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('I') {
			state = uint16(112)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32(':') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') || lookahead1 == int32(0xb7) {
			state = uint16(115)
			goto next_state
		}
		return result
	case int32(96):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('I') {
			state = uint16(89)
			goto next_state
		}
		if lookahead1 == int32('Y') {
			state = uint16(61)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32(':') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') || lookahead1 == int32(0xb7) {
			state = uint16(115)
			goto next_state
		}
		return result
	case int32(97):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('K') {
			state = uint16(86)
			goto next_state
		}
		if lookahead1 == int32(':') || lookahead1 == int32(0xb7) {
			state = uint16(115)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(114)
			goto next_state
		}
		return result
	case int32(98):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('K') {
			state = uint16(90)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32(':') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') || lookahead1 == int32(0xb7) {
			state = uint16(115)
			goto next_state
		}
		return result
	case int32(99):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('M') {
			state = uint16(111)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32(':') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') || lookahead1 == int32(0xb7) {
			state = uint16(115)
			goto next_state
		}
		return result
	case int32(100):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('N') {
			state = uint16(58)
			goto next_state
		}
		if lookahead1 == int32(':') || lookahead1 == int32(0xb7) {
			state = uint16(115)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(114)
			goto next_state
		}
		return result
	case int32(101):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('N') {
			state = uint16(110)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32(':') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') || lookahead1 == int32(0xb7) {
			state = uint16(115)
			goto next_state
		}
		return result
	case int32(102):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('N') {
			state = uint16(59)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32(':') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') || lookahead1 == int32(0xb7) {
			state = uint16(115)
			goto next_state
		}
		return result
	case int32(103):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('O') {
			state = uint16(97)
			goto next_state
		}
		if lookahead1 == int32(':') || lookahead1 == int32(0xb7) {
			state = uint16(115)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(114)
			goto next_state
		}
		return result
	case int32(104):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('O') {
			state = uint16(98)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32(':') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') || lookahead1 == int32(0xb7) {
			state = uint16(115)
			goto next_state
		}
		return result
	case int32(105):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('S') {
			state = uint16(60)
			goto next_state
		}
		if lookahead1 == int32(':') || lookahead1 == int32(0xb7) {
			state = uint16(115)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(114)
			goto next_state
		}
		return result
	case int32(106):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('S') {
			state = uint16(61)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32(':') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') || lookahead1 == int32(0xb7) {
			state = uint16(115)
			goto next_state
		}
		return result
	case int32(107):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('T') {
			state = uint16(93)
			goto next_state
		}
		if lookahead1 == int32(':') || lookahead1 == int32(0xb7) {
			state = uint16(115)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(114)
			goto next_state
		}
		return result
	case int32(108):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('T') {
			state = uint16(103)
			goto next_state
		}
		if lookahead1 == int32(':') || lookahead1 == int32(0xb7) {
			state = uint16(115)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(114)
			goto next_state
		}
		return result
	case int32(109):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('T') {
			state = uint16(94)
			goto next_state
		}
		if lookahead1 == int32(':') || lookahead1 == int32(0xb7) {
			state = uint16(115)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(114)
			goto next_state
		}
		return result
	case int32(110):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('T') {
			state = uint16(95)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32(':') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') || lookahead1 == int32(0xb7) {
			state = uint16(115)
			goto next_state
		}
		return result
	case int32(111):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('T') {
			state = uint16(104)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32(':') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') || lookahead1 == int32(0xb7) {
			state = uint16(115)
			goto next_state
		}
		return result
	case int32(112):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('T') {
			state = uint16(96)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32(':') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') || lookahead1 == int32(0xb7) {
			state = uint16(115)
			goto next_state
		}
		return result
	case int32(113):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(':') || lookahead1 == int32(0xb7) {
			state = uint16(115)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(113)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('G') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('g') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(114)
			goto next_state
		}
		return result
	case int32(114):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(':') || lookahead1 == int32(0xb7) {
			state = uint16(115)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(114)
			goto next_state
		}
		return result
	case int32(115):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32(':') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') || lookahead1 == int32(0xb7) {
			state = uint16(115)
			goto next_state
		}
		return result
	case int32(116):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Nmtoken)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(116)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || lookahead1 == int32(':') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') || lookahead1 == int32(0xb7) {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(117):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Nmtoken)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(117)
			goto next_state
		}
		if int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(118)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || lookahead1 == int32(':') || int32('G') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('g') <= lookahead1 && lookahead1 <= int32('z') || lookahead1 == int32(0xb7) {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(118):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Nmtoken)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(118)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || lookahead1 == int32(':') || int32('G') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('g') <= lookahead1 && lookahead1 <= int32('z') || lookahead1 == int32(0xb7) {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(119):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Nmtoken)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32(':') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') || lookahead1 == int32(0xb7) {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(120):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AMP)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(121):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AMP_POUND)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('x') {
			state = uint16(123)
			goto next_state
		}
		return result
	case int32(122):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_CharRef_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(122)
			goto next_state
		}
		return result
	case int32(123):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AMP_POUNDx)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(124):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_CharRef_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(124)
			goto next_state
		}
		return result
	case int32(125):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_AttValue_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(126):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_AttValue_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(127):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_SystemLiteral_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 != 0 && lookahead1 != int32('"') {
			state = uint16(127)
			goto next_state
		}
		return result
	case int32(128):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_SystemLiteral_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 != 0 && lookahead1 != int32('\'') {
			state = uint16(128)
			goto next_state
		}
		return result
	case int32(129):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_PubidLiteral_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		v4 = uintptr(unsafe.Pointer(&aux_sym_PubidLiteral_token1_character_set_1))
		v5 = lookahead1
		index = uint32(0)
		size = uint32(9) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v4 + uintptr(mid_index)*8
			if v5 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v5 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v6 = libc.BoolUint8(true1 != 0)
				goto _7
			} else {
				if v5 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v4 + uintptr(index)*8
		v6 = libc.BoolUint8(v5 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v5 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _7
	_7:
		if v6 != 0 {
			state = uint16(129)
			goto next_state
		}
		return result
	case int32(130):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_PubidLiteral_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		v4 = uintptr(unsafe.Pointer(&aux_sym_PubidLiteral_token2_character_set_1))
		v5 = lookahead1
		index = uint32(0)
		size = uint32(9) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v4 + uintptr(mid_index)*8
			if v5 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v5 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v6 = libc.BoolUint8(true1 != 0)
				goto _11
			} else {
				if v5 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v4 + uintptr(index)*8
		v6 = libc.BoolUint8(v5 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v5 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _11
	_11:
		if v6 != 0 {
			state = uint16(130)
			goto next_state
		}
		return result
	case int32(131):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_VersionNum)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(131)
			goto next_state
		}
		return result
	case int32(132):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_EncName)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(132)
			goto next_state
		}
		return result
	case int32(133):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var map_token = [62]uint16_t{
	0:  uint16('"'),
	1:  uint16(66),
	2:  uint16('#'),
	3:  uint16(70),
	4:  uint16('%'),
	5:  uint16(65),
	6:  uint16('&'),
	7:  uint16(120),
	8:  uint16('\''),
	9:  uint16(80),
	10: uint16('('),
	11: uint16(48),
	12: uint16(')'),
	13: uint16(51),
	14: uint16('*'),
	15: uint16(52),
	16: uint16('+'),
	17: uint16(54),
	18: uint16(','),
	19: uint16(55),
	20: uint16('1'),
	21: uint16(68),
	22: uint16(';'),
	23: uint16(82),
	24: uint16('<'),
	25: uint16(1),
	26: uint16('='),
	27: uint16(133),
	28: uint16('>'),
	29: uint16(47),
	30: uint16('?'),
	31: uint16(53),
	32: uint16('E'),
	33: uint16(72),
	34: uint16('I'),
	35: uint16(69),
	36: uint16('N'),
	37: uint16(71),
	38: uint16('['),
	39: uint16(44),
	40: uint16(']'),
	41: uint16(73),
	42: uint16('_'),
	43: uint16(79),
	44: uint16('|'),
	45: uint16(50),
	46: uint16('\t'),
	47: uint16(76),
	48: uint16('\n'),
	49: uint16(76),
	50: uint16('\r'),
	51: uint16(76),
	52: uint16(' '),
	53: uint16(76),
	54: uint16('-'),
	55: uint16(78),
	56: uint16('.'),
	57: uint16(78),
	58: uint16(':'),
	59: uint16(78),
	60: uint16(0xb7),
	61: uint16(78),
}

var map_token1 = [20]uint16_t{
	0:  uint16('%'),
	1:  uint16(65),
	2:  uint16('('),
	3:  uint16(48),
	4:  uint16('?'),
	5:  uint16(9),
	6:  uint16('E'),
	7:  uint16(101),
	8:  uint16('I'),
	9:  uint16(84),
	10: uint16('N'),
	11: uint16(99),
	12: uint16('\t'),
	13: uint16(83),
	14: uint16('\n'),
	15: uint16(83),
	16: uint16('\r'),
	17: uint16(83),
	18: uint16(' '),
	19: uint16(83),
}

var map_token2 = [44]uint16_t{
	0:  uint16('"'),
	1:  uint16(66),
	2:  uint16('#'),
	3:  uint16(22),
	4:  uint16('%'),
	5:  uint16(65),
	6:  uint16('\''),
	7:  uint16(80),
	8:  uint16('('),
	9:  uint16(48),
	10: uint16(')'),
	11: uint16(51),
	12: uint16('*'),
	13: uint16(52),
	14: uint16('+'),
	15: uint16(54),
	16: uint16(','),
	17: uint16(55),
	18: uint16('1'),
	19: uint16(7),
	20: uint16(';'),
	21: uint16(82),
	22: uint16('<'),
	23: uint16(1),
	24: uint16('='),
	25: uint16(133),
	26: uint16('>'),
	27: uint16(47),
	28: uint16('?'),
	29: uint16(53),
	30: uint16('['),
	31: uint16(44),
	32: uint16(']'),
	33: uint16(34),
	34: uint16('|'),
	35: uint16(50),
	36: uint16('\t'),
	37: uint16(83),
	38: uint16('\n'),
	39: uint16(83),
	40: uint16('\r'),
	41: uint16(83),
	42: uint16(' '),
	43: uint16(83),
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
	switch int32(state) {
	case 0:
		i = uint32(0)
		for {
			if !(uint64(i) < libc.Uint64FromInt64(40)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token3[i]) == lookahead {
				state = map_token3[i+uint32(1)]
				goto next_state
			}
			goto _1
		_1:
			;
			i = i + uint32(2)
		}
		return result
	case int32(1):
		if lookahead == int32('N') {
			state = uint16(11)
			goto next_state
		}
		if lookahead == int32('T') {
			state = uint16(12)
			goto next_state
		}
		return result
	case int32(2):
		if lookahead == int32('D') {
			state = uint16(13)
			goto next_state
		}
		return result
	case int32(3):
		if lookahead == int32('L') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('M') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32('N') {
			state = uint16(16)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead == int32('G') {
			state = uint16(17)
			goto next_state
		}
		if lookahead == int32('N') {
			state = uint16(18)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead == int32('D') {
			state = uint16(19)
			goto next_state
		}
		if lookahead == int32('O') {
			state = uint16(20)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead == int32('U') {
			state = uint16(21)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead == int32('Y') {
			state = uint16(22)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead == int32('n') {
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead == int32('e') {
			state = uint16(24)
			goto next_state
		}
		return result
	case int32(10):
		if lookahead == int32('m') {
			state = uint16(25)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead == int32('Y') {
			state = uint16(26)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead == int32('T') {
			state = uint16(27)
			goto next_state
		}
		return result
	case int32(13):
		if lookahead == int32('A') {
			state = uint16(28)
			goto next_state
		}
		return result
	case int32(14):
		if lookahead == int32('E') {
			state = uint16(29)
			goto next_state
		}
		return result
	case int32(15):
		if lookahead == int32('P') {
			state = uint16(30)
			goto next_state
		}
		return result
	case int32(16):
		if lookahead == int32('T') {
			state = uint16(31)
			goto next_state
		}
		return result
	case int32(17):
		if lookahead == int32('N') {
			state = uint16(32)
			goto next_state
		}
		return result
	case int32(18):
		if lookahead == int32('C') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(19):
		if lookahead == int32('A') {
			state = uint16(34)
			goto next_state
		}
		return result
	case int32(20):
		if lookahead == int32('T') {
			state = uint16(35)
			goto next_state
		}
		return result
	case int32(21):
		if lookahead == int32('B') {
			state = uint16(36)
			goto next_state
		}
		return result
	case int32(22):
		if lookahead == int32('S') {
			state = uint16(37)
			goto next_state
		}
		return result
	case int32(23):
		if lookahead == int32('c') {
			state = uint16(38)
			goto next_state
		}
		return result
	case int32(24):
		if lookahead == int32('r') {
			state = uint16(39)
			goto next_state
		}
		return result
	case int32(25):
		if lookahead == int32('l') {
			state = uint16(40)
			goto next_state
		}
		return result
	case int32(26):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ANY)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(27):
		if lookahead == int32('L') {
			state = uint16(41)
			goto next_state
		}
		return result
	case int32(28):
		if lookahead == int32('T') {
			state = uint16(42)
			goto next_state
		}
		return result
	case int32(29):
		if lookahead == int32('M') {
			state = uint16(43)
			goto next_state
		}
		return result
	case int32(30):
		if lookahead == int32('T') {
			state = uint16(44)
			goto next_state
		}
		return result
	case int32(31):
		if lookahead == int32('I') {
			state = uint16(45)
			goto next_state
		}
		return result
	case int32(32):
		if lookahead == int32('O') {
			state = uint16(46)
			goto next_state
		}
		return result
	case int32(33):
		if lookahead == int32('L') {
			state = uint16(47)
			goto next_state
		}
		return result
	case int32(34):
		if lookahead == int32('T') {
			state = uint16(48)
			goto next_state
		}
		return result
	case int32(35):
		if lookahead == int32('A') {
			state = uint16(49)
			goto next_state
		}
		return result
	case int32(36):
		if lookahead == int32('L') {
			state = uint16(50)
			goto next_state
		}
		return result
	case int32(37):
		if lookahead == int32('T') {
			state = uint16(51)
			goto next_state
		}
		return result
	case int32(38):
		if lookahead == int32('o') {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(39):
		if lookahead == int32('s') {
			state = uint16(53)
			goto next_state
		}
		return result
	case int32(40):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_xml)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(41):
		if lookahead == int32('I') {
			state = uint16(54)
			goto next_state
		}
		return result
	case int32(42):
		if lookahead == int32('A') {
			state = uint16(55)
			goto next_state
		}
		return result
	case int32(43):
		if lookahead == int32('E') {
			state = uint16(56)
			goto next_state
		}
		return result
	case int32(44):
		if lookahead == int32('Y') {
			state = uint16(57)
			goto next_state
		}
		return result
	case int32(45):
		if lookahead == int32('T') {
			state = uint16(58)
			goto next_state
		}
		return result
	case int32(46):
		if lookahead == int32('R') {
			state = uint16(59)
			goto next_state
		}
		return result
	case int32(47):
		if lookahead == int32('U') {
			state = uint16(60)
			goto next_state
		}
		return result
	case int32(48):
		if lookahead == int32('A') {
			state = uint16(61)
			goto next_state
		}
		return result
	case int32(49):
		if lookahead == int32('T') {
			state = uint16(62)
			goto next_state
		}
		return result
	case int32(50):
		if lookahead == int32('I') {
			state = uint16(63)
			goto next_state
		}
		return result
	case int32(51):
		if lookahead == int32('E') {
			state = uint16(64)
			goto next_state
		}
		return result
	case int32(52):
		if lookahead == int32('d') {
			state = uint16(65)
			goto next_state
		}
		return result
	case int32(53):
		if lookahead == int32('i') {
			state = uint16(66)
			goto next_state
		}
		return result
	case int32(54):
		if lookahead == int32('S') {
			state = uint16(67)
			goto next_state
		}
		return result
	case int32(55):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_StringType)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(56):
		if lookahead == int32('N') {
			state = uint16(68)
			goto next_state
		}
		return result
	case int32(57):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EMPTY)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(58):
		if lookahead == int32('Y') {
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(59):
		if lookahead == int32('E') {
			state = uint16(70)
			goto next_state
		}
		return result
	case int32(60):
		if lookahead == int32('D') {
			state = uint16(71)
			goto next_state
		}
		return result
	case int32(61):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_NDATA)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(62):
		if lookahead == int32('I') {
			state = uint16(72)
			goto next_state
		}
		return result
	case int32(63):
		if lookahead == int32('C') {
			state = uint16(73)
			goto next_state
		}
		return result
	case int32(64):
		if lookahead == int32('M') {
			state = uint16(74)
			goto next_state
		}
		return result
	case int32(65):
		if lookahead == int32('i') {
			state = uint16(75)
			goto next_state
		}
		return result
	case int32(66):
		if lookahead == int32('o') {
			state = uint16(76)
			goto next_state
		}
		return result
	case int32(67):
		if lookahead == int32('T') {
			state = uint16(77)
			goto next_state
		}
		return result
	case int32(68):
		if lookahead == int32('T') {
			state = uint16(78)
			goto next_state
		}
		return result
	case int32(69):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ENTITY)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(70):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_IGNORE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(71):
		if lookahead == int32('E') {
			state = uint16(79)
			goto next_state
		}
		return result
	case int32(72):
		if lookahead == int32('O') {
			state = uint16(80)
			goto next_state
		}
		return result
	case int32(73):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PUBLIC)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(74):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SYSTEM)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(75):
		if lookahead == int32('n') {
			state = uint16(81)
			goto next_state
		}
		return result
	case int32(76):
		if lookahead == int32('n') {
			state = uint16(82)
			goto next_state
		}
		return result
	case int32(77):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ATTLIST)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(78):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ELEMENT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(79):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_INCLUDE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(80):
		if lookahead == int32('N') {
			state = uint16(83)
			goto next_state
		}
		return result
	case int32(81):
		if lookahead == int32('g') {
			state = uint16(84)
			goto next_state
		}
		return result
	case int32(82):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_version)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(83):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_NOTATION)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(84):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_encoding)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var map_token3 = [20]uint16_t{
	0:  uint16('A'),
	1:  uint16(1),
	2:  uint16('C'),
	3:  uint16(2),
	4:  uint16('E'),
	5:  uint16(3),
	6:  uint16('I'),
	7:  uint16(4),
	8:  uint16('N'),
	9:  uint16(5),
	10: uint16('P'),
	11: uint16(6),
	12: uint16('S'),
	13: uint16(7),
	14: uint16('e'),
	15: uint16(8),
	16: uint16('v'),
	17: uint16(9),
	18: uint16('x'),
	19: uint16(10),
}

var ts_lex_modes = [334]TSLexMode{
	0: {
		Fexternal_lex_state: uint16(1),
	},
	1: {
		Flex_state:          uint16(39),
		Fexternal_lex_state: uint16(2),
	},
	2: {
		Flex_state:          uint16(39),
		Fexternal_lex_state: uint16(2),
	},
	3: {
		Flex_state:          uint16(39),
		Fexternal_lex_state: uint16(2),
	},
	4: {
		Flex_state:          uint16(39),
		Fexternal_lex_state: uint16(2),
	},
	5: {
		Flex_state:          uint16(39),
		Fexternal_lex_state: uint16(2),
	},
	6: {
		Flex_state:          uint16(39),
		Fexternal_lex_state: uint16(2),
	},
	7: {
		Flex_state:          uint16(39),
		Fexternal_lex_state: uint16(2),
	},
	8: {
		Flex_state:          uint16(39),
		Fexternal_lex_state: uint16(2),
	},
	9: {
		Flex_state:          uint16(39),
		Fexternal_lex_state: uint16(2),
	},
	10: {
		Flex_state:          uint16(39),
		Fexternal_lex_state: uint16(2),
	},
	11: {
		Flex_state:          uint16(39),
		Fexternal_lex_state: uint16(2),
	},
	12: {
		Flex_state: uint16(39),
	},
	13: {
		Flex_state: uint16(2),
	},
	14: {
		Flex_state: uint16(4),
	},
	15: {
		Flex_state: uint16(39),
	},
	16: {
		Flex_state: uint16(4),
	},
	17: {
		Flex_state: uint16(2),
	},
	18: {
		Flex_state: uint16(4),
	},
	19: {
		Flex_state: uint16(39),
	},
	20: {
		Flex_state: uint16(2),
	},
	21: {
		Flex_state: uint16(5),
	},
	22: {
		Flex_state: uint16(39),
	},
	23: {
		Flex_state: uint16(6),
	},
	24: {
		Flex_state: uint16(39),
	},
	25: {
		Flex_state: uint16(39),
	},
	26: {
		Flex_state: uint16(39),
	},
	27: {
		Flex_state: uint16(39),
	},
	28: {
		Flex_state: uint16(3),
	},
	29: {
		Flex_state: uint16(6),
	},
	30: {
		Flex_state: uint16(39),
	},
	31: {
		Flex_state: uint16(3),
	},
	32: {
		Flex_state: uint16(6),
	},
	33: {
		Flex_state: uint16(39),
	},
	34: {
		Flex_state: uint16(3),
	},
	35: {
		Flex_state:          uint16(39),
		Fexternal_lex_state: uint16(2),
	},
	36: {
		Flex_state:          uint16(39),
		Fexternal_lex_state: uint16(2),
	},
	37: {
		Flex_state: uint16(39),
	},
	38: {
		Flex_state: uint16(39),
	},
	39: {
		Flex_state:          uint16(39),
		Fexternal_lex_state: uint16(2),
	},
	40: {
		Flex_state:          uint16(39),
		Fexternal_lex_state: uint16(2),
	},
	41: {
		Flex_state:          uint16(39),
		Fexternal_lex_state: uint16(2),
	},
	42: {
		Flex_state:          uint16(39),
		Fexternal_lex_state: uint16(2),
	},
	43: {
		Flex_state:          uint16(39),
		Fexternal_lex_state: uint16(2),
	},
	44: {
		Flex_state: uint16(39),
	},
	45: {
		Flex_state:          uint16(39),
		Fexternal_lex_state: uint16(2),
	},
	46: {
		Flex_state:          uint16(39),
		Fexternal_lex_state: uint16(2),
	},
	47: {
		Flex_state:          uint16(39),
		Fexternal_lex_state: uint16(2),
	},
	48: {
		Flex_state:          uint16(39),
		Fexternal_lex_state: uint16(2),
	},
	49: {
		Flex_state:          uint16(39),
		Fexternal_lex_state: uint16(2),
	},
	50: {
		Flex_state:          uint16(39),
		Fexternal_lex_state: uint16(2),
	},
	51: {
		Flex_state: uint16(39),
	},
	52: {
		Flex_state:          uint16(39),
		Fexternal_lex_state: uint16(2),
	},
	53: {
		Flex_state:          uint16(39),
		Fexternal_lex_state: uint16(2),
	},
	54: {
		Flex_state:          uint16(39),
		Fexternal_lex_state: uint16(2),
	},
	55: {
		Flex_state:          uint16(39),
		Fexternal_lex_state: uint16(2),
	},
	56: {
		Flex_state:          uint16(39),
		Fexternal_lex_state: uint16(2),
	},
	57: {
		Flex_state:          uint16(39),
		Fexternal_lex_state: uint16(2),
	},
	58: {
		Flex_state:          uint16(39),
		Fexternal_lex_state: uint16(2),
	},
	59: {
		Flex_state: uint16(39),
	},
	60: {
		Flex_state:          uint16(39),
		Fexternal_lex_state: uint16(2),
	},
	61: {
		Flex_state: uint16(39),
	},
	62: {
		Flex_state: uint16(39),
	},
	63: {
		Flex_state: uint16(39),
	},
	64: {
		Flex_state: uint16(39),
	},
	65: {
		Flex_state: uint16(39),
	},
	66: {
		Flex_state: uint16(39),
	},
	67: {
		Flex_state: uint16(39),
	},
	68: {
		Flex_state: uint16(39),
	},
	69: {
		Flex_state: uint16(39),
	},
	70: {
		Flex_state: uint16(4),
	},
	71: {
		Flex_state: uint16(4),
	},
	72: {
		Flex_state: uint16(39),
	},
	73: {
		Flex_state: uint16(39),
	},
	74: {
		Flex_state: uint16(39),
	},
	75: {
		Flex_state: uint16(2),
	},
	76: {
		Flex_state: uint16(2),
	},
	77: {
		Flex_state:          uint16(39),
		Fexternal_lex_state: uint16(2),
	},
	78: {
		Flex_state: uint16(39),
	},
	79: {
		Flex_state: uint16(39),
	},
	80: {
		Flex_state: uint16(39),
	},
	81: {
		Flex_state: uint16(39),
	},
	82: {
		Flex_state: uint16(39),
	},
	83: {
		Flex_state:          uint16(39),
		Fexternal_lex_state: uint16(2),
	},
	84: {
		Flex_state: uint16(39),
	},
	85: {
		Flex_state: uint16(39),
	},
	86: {
		Flex_state: uint16(39),
	},
	87: {
		Flex_state:          uint16(39),
		Fexternal_lex_state: uint16(2),
	},
	88: {
		Flex_state: uint16(4),
	},
	89: {
		Flex_state: uint16(4),
	},
	90: {
		Flex_state: uint16(2),
	},
	91: {
		Flex_state: uint16(2),
	},
	92: {
		Flex_state: uint16(39),
	},
	93: {
		Flex_state: uint16(39),
	},
	94: {
		Flex_state: uint16(39),
	},
	95: {
		Flex_state: uint16(39),
	},
	96: {
		Flex_state: uint16(39),
	},
	97: {
		Flex_state: uint16(39),
	},
	98: {
		Flex_state: uint16(39),
	},
	99: {
		Flex_state: uint16(39),
	},
	100: {
		Flex_state: uint16(39),
	},
	101: {
		Flex_state: uint16(39),
	},
	102: {
		Flex_state: uint16(39),
	},
	103: {
		Flex_state: uint16(39),
	},
	104: {
		Flex_state: uint16(39),
	},
	105: {
		Flex_state: uint16(39),
	},
	106: {
		Flex_state: uint16(39),
	},
	107: {
		Flex_state: uint16(39),
	},
	108: {
		Flex_state: uint16(39),
	},
	109: {
		Flex_state: uint16(39),
	},
	110: {
		Flex_state: uint16(3),
	},
	111: {
		Flex_state: uint16(3),
	},
	112: {
		Flex_state: uint16(3),
	},
	113: {
		Flex_state: uint16(39),
	},
	114: {
		Flex_state: uint16(6),
	},
	115: {
		Flex_state: uint16(6),
	},
	116: {
		Flex_state: uint16(6),
	},
	117: {
		Flex_state: uint16(39),
	},
	118: {
		Flex_state: uint16(39),
	},
	119: {
		Flex_state: uint16(39),
	},
	120: {
		Flex_state: uint16(39),
	},
	121: {
		Flex_state: uint16(39),
	},
	122: {
		Flex_state: uint16(39),
	},
	123: {
		Flex_state: uint16(39),
	},
	124: {
		Flex_state: uint16(39),
	},
	125: {
		Flex_state: uint16(39),
	},
	126: {
		Flex_state: uint16(39),
	},
	127: {
		Flex_state: uint16(39),
	},
	128: {
		Flex_state: uint16(39),
	},
	129: {
		Flex_state: uint16(39),
	},
	130: {
		Flex_state: uint16(39),
	},
	131: {
		Flex_state: uint16(39),
	},
	132: {
		Flex_state: uint16(39),
	},
	133: {
		Flex_state: uint16(39),
	},
	134: {
		Flex_state: uint16(39),
	},
	135: {
		Flex_state: uint16(39),
	},
	136: {
		Flex_state: uint16(39),
	},
	137: {
		Flex_state: uint16(39),
	},
	138: {
		Flex_state: uint16(39),
	},
	139: {
		Flex_state: uint16(39),
	},
	140: {
		Flex_state: uint16(39),
	},
	141: {
		Flex_state: uint16(39),
	},
	142: {
		Flex_state: uint16(39),
	},
	143: {
		Flex_state: uint16(39),
	},
	144: {
		Flex_state: uint16(39),
	},
	145: {
		Flex_state: uint16(39),
	},
	146: {
		Flex_state: uint16(39),
	},
	147: {
		Flex_state: uint16(39),
	},
	148: {
		Flex_state: uint16(39),
	},
	149: {
		Flex_state: uint16(39),
	},
	150: {
		Flex_state: uint16(39),
	},
	151: {
		Flex_state: uint16(39),
	},
	152: {
		Flex_state: uint16(39),
	},
	153: {
		Flex_state: uint16(39),
	},
	154: {
		Flex_state: uint16(39),
	},
	155: {
		Flex_state: uint16(39),
	},
	156: {
		Flex_state: uint16(39),
	},
	157: {
		Flex_state: uint16(39),
	},
	158: {
		Flex_state: uint16(5),
	},
	159: {
		Flex_state: uint16(39),
	},
	160: {
		Flex_state: uint16(39),
	},
	161: {
		Flex_state: uint16(39),
	},
	162: {
		Flex_state: uint16(39),
	},
	163: {
		Flex_state: uint16(39),
	},
	164: {
		Flex_state: uint16(39),
	},
	165: {
		Flex_state: uint16(39),
	},
	166: {
		Flex_state: uint16(39),
	},
	167: {
		Flex_state: uint16(39),
	},
	168: {
		Flex_state: uint16(39),
	},
	169: {
		Flex_state: uint16(39),
	},
	170: {
		Flex_state: uint16(39),
	},
	171: {
		Flex_state: uint16(39),
	},
	172: {
		Flex_state: uint16(39),
	},
	173: {
		Flex_state: uint16(39),
	},
	174: {
		Flex_state: uint16(39),
	},
	175: {
		Flex_state: uint16(39),
	},
	176: {
		Flex_state: uint16(39),
	},
	177: {
		Flex_state: uint16(39),
	},
	178: {
		Flex_state: uint16(39),
	},
	179: {
		Flex_state: uint16(39),
	},
	180: {
		Flex_state: uint16(39),
	},
	181: {
		Flex_state: uint16(39),
	},
	182: {
		Flex_state: uint16(39),
	},
	183: {
		Flex_state: uint16(39),
	},
	184: {
		Flex_state:          uint16(39),
		Fexternal_lex_state: uint16(3),
	},
	185: {
		Flex_state: uint16(39),
	},
	186: {
		Flex_state: uint16(39),
	},
	187: {
		Flex_state: uint16(10),
	},
	188: {
		Flex_state: uint16(39),
	},
	189: {
		Flex_state: uint16(39),
	},
	190: {
		Flex_state: uint16(39),
	},
	191: {
		Flex_state: uint16(39),
	},
	192: {
		Flex_state: uint16(39),
	},
	193: {
		Flex_state: uint16(5),
	},
	194: {
		Flex_state: uint16(5),
	},
	195: {
		Flex_state: uint16(39),
	},
	196: {
		Flex_state: uint16(39),
	},
	197: {
		Flex_state: uint16(35),
	},
	198: {
		Flex_state: uint16(39),
	},
	199: {
		Flex_state: uint16(39),
	},
	200: {
		Flex_state: uint16(39),
	},
	201: {
		Flex_state: uint16(5),
	},
	202: {
		Flex_state: uint16(39),
	},
	203: {
		Flex_state: uint16(39),
	},
	204: {
		Flex_state: uint16(39),
	},
	205: {
		Flex_state: uint16(39),
	},
	206: {
		Flex_state: uint16(39),
	},
	207: {
		Flex_state: uint16(39),
	},
	208: {
		Flex_state: uint16(39),
	},
	209: {
		Flex_state: uint16(39),
	},
	210: {
		Flex_state: uint16(39),
	},
	211: {
		Flex_state: uint16(39),
	},
	212: {
		Flex_state: uint16(39),
	},
	213: {
		Flex_state: uint16(39),
	},
	214: {
		Flex_state: uint16(35),
	},
	215: {
		Flex_state: uint16(39),
	},
	216: {
		Flex_state: uint16(39),
	},
	217: {
		Flex_state: uint16(39),
	},
	218: {
		Flex_state: uint16(39),
	},
	219: {
		Flex_state: uint16(39),
	},
	220: {
		Flex_state: uint16(39),
	},
	221: {
		Flex_state: uint16(39),
	},
	222: {
		Flex_state: uint16(39),
	},
	223: {
		Flex_state: uint16(39),
	},
	224: {
		Flex_state: uint16(39),
	},
	225: {
		Flex_state: uint16(39),
	},
	226: {
		Flex_state: uint16(39),
	},
	227: {
		Flex_state: uint16(39),
	},
	228: {
		Flex_state: uint16(39),
	},
	229: {
		Flex_state: uint16(39),
	},
	230: {
		Flex_state: uint16(39),
	},
	231: {
		Flex_state: uint16(39),
	},
	232: {
		Flex_state: uint16(39),
	},
	233: {
		Flex_state: uint16(39),
	},
	234: {
		Flex_state: uint16(39),
	},
	235: {
		Flex_state: uint16(39),
	},
	236: {
		Flex_state: uint16(39),
	},
	237: {
		Flex_state: uint16(39),
	},
	238: {
		Flex_state: uint16(39),
	},
	239: {
		Flex_state: uint16(39),
	},
	240: {
		Flex_state: uint16(5),
	},
	241: {
		Flex_state: uint16(39),
	},
	242: {
		Flex_state: uint16(39),
	},
	243: {
		Flex_state: uint16(35),
	},
	244: {
		Flex_state: uint16(39),
	},
	245: {
		Flex_state: uint16(39),
	},
	246: {
		Flex_state: uint16(37),
	},
	247: {
		Flex_state: uint16(39),
	},
	248: {
		Flex_state: uint16(39),
	},
	249: {
		Fexternal_lex_state: uint16(4),
	},
	250: {
		Flex_state: uint16(39),
	},
	251: {
		Flex_state: uint16(39),
	},
	252: {
		Flex_state: uint16(39),
	},
	253: {
		Flex_state: uint16(39),
	},
	254: {
		Flex_state: uint16(39),
	},
	255: {
		Flex_state: uint16(39),
	},
	256: {
		Flex_state: uint16(35),
	},
	257: {
		Flex_state: uint16(5),
	},
	258: {
		Flex_state: uint16(39),
	},
	259: {
		Flex_state: uint16(39),
	},
	260: {
		Flex_state: uint16(39),
	},
	261: {
		Flex_state: uint16(39),
	},
	262: {
		Flex_state: uint16(39),
	},
	263: {
		Flex_state: uint16(39),
	},
	264: {
		Flex_state: uint16(39),
	},
	265: {
		Flex_state: uint16(39),
	},
	266: {
		Flex_state: uint16(39),
	},
	267: {
		Flex_state: uint16(39),
	},
	268: {
		Flex_state: uint16(127),
	},
	269: {
		Flex_state: uint16(39),
	},
	270: {
		Flex_state: uint16(39),
	},
	271: {
		Flex_state: uint16(128),
	},
	272: {
		Flex_state: uint16(35),
	},
	273: {
		Flex_state: uint16(39),
	},
	274: {
		Flex_state: uint16(39),
	},
	275: {
		Flex_state: uint16(129),
	},
	276: {
		Flex_state: uint16(39),
	},
	277: {
		Flex_state: uint16(5),
	},
	278: {
		Flex_state: uint16(130),
	},
	279: {
		Flex_state: uint16(39),
	},
	280: {
		Flex_state: uint16(39),
	},
	281: {
		Flex_state: uint16(35),
	},
	282: {},
	283: {
		Flex_state: uint16(39),
	},
	284: {
		Flex_state: uint16(39),
	},
	285: {
		Flex_state: uint16(39),
	},
	286: {
		Flex_state: uint16(39),
	},
	287: {
		Flex_state: uint16(39),
	},
	288: {
		Flex_state: uint16(39),
	},
	289: {
		Flex_state: uint16(39),
	},
	290: {
		Flex_state: uint16(39),
	},
	291: {
		Flex_state: uint16(39),
	},
	292: {
		Flex_state: uint16(38),
	},
	293: {
		Flex_state: uint16(39),
	},
	294: {
		Flex_state: uint16(39),
	},
	295: {
		Flex_state: uint16(39),
	},
	296: {
		Flex_state: uint16(39),
	},
	297: {
		Flex_state: uint16(39),
	},
	298: {
		Flex_state: uint16(39),
	},
	299: {
		Fexternal_lex_state: uint16(3),
	},
	300: {
		Flex_state: uint16(39),
	},
	301: {
		Flex_state: uint16(39),
	},
	302: {
		Flex_state: uint16(39),
	},
	303: {
		Flex_state: uint16(39),
	},
	304: {
		Flex_state: uint16(39),
	},
	305: {
		Flex_state: uint16(39),
	},
	306: {
		Flex_state: uint16(39),
	},
	307: {
		Flex_state: uint16(39),
	},
	308: {
		Flex_state: uint16(39),
	},
	309: {
		Flex_state: uint16(39),
	},
	310: {
		Flex_state: uint16(38),
	},
	311: {
		Flex_state: uint16(5),
	},
	312: {
		Flex_state: uint16(39),
	},
	313: {
		Flex_state: uint16(39),
	},
	314: {
		Flex_state: uint16(39),
	},
	315: {
		Flex_state: uint16(39),
	},
	316: {
		Flex_state: uint16(39),
	},
	317: {
		Flex_state: uint16(39),
	},
	318: {
		Flex_state: uint16(39),
	},
	319: {
		Flex_state: uint16(39),
	},
	320: {
		Flex_state: uint16(39),
	},
	321: {
		Flex_state: uint16(39),
	},
	322: {
		Flex_state: uint16(39),
	},
	323: {
		Flex_state: uint16(5),
	},
	324: {
		Flex_state: uint16(37),
	},
	325: {
		Flex_state: uint16(39),
	},
	326: {
		Flex_state: uint16(39),
	},
	327: {
		Flex_state: uint16(5),
	},
	328: {
		Flex_state: uint16(37),
	},
	329: {
		Flex_state: uint16(39),
	},
	330: {
		Flex_state: uint16(39),
	},
	331: {
		Flex_state: uint16(5),
	},
	332: {
		Flex_state: uint16(37),
	},
	333: {
		Flex_state: uint16(39),
	},
}

var ts_parse_table = [2][111]uint16_t{
	0: {
		0:  uint16(1),
		1:  uint16(1),
		2:  uint16(1),
		3:  uint16(1),
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
		42: uint16(1),
		43: uint16(1),
		44: uint16(1),
		45: uint16(1),
		46: uint16(1),
		47: uint16(1),
		48: uint16(1),
		53: uint16(1),
		54: uint16(1),
		55: uint16(1),
		56: uint16(1),
		57: uint16(1),
		58: uint16(1),
		59: uint16(1),
		60: uint16(1),
	},
	1: {
		2:  uint16(3),
		5:  uint16(5),
		10: uint16(7),
		31: uint16(9),
		38: uint16(11),
		60: uint16(13),
		61: uint16(282),
		62: uint16(11),
		63: uint16(6),
		64: uint16(6),
		65: uint16(6),
		66: uint16(6),
		67: uint16(36),
		73: uint16(36),
		80: uint16(36),
		81: uint16(45),
		82: uint16(45),
		85: uint16(36),
		86: uint16(6),
		97: uint16(36),
		99: uint16(6),
	},
}

var ts_small_parse_table = [3470]uint16_t{
	0:    uint16(10),
	1:    uint16(17),
	2:    uint16(1),
	3:    uint16(anon_sym_LT_QMARK),
	4:    uint16(20),
	5:    uint16(1),
	6:    uint16(anon_sym_LT_BANG_LBRACK),
	7:    uint16(23),
	8:    uint16(1),
	9:    uint16(anon_sym_LT_BANG),
	10:   uint16(26),
	11:   uint16(1),
	12:   uint16(anon_sym_PERCENT),
	13:   uint16(29),
	14:   uint16(1),
	15:   uint16(sym__S),
	16:   uint16(32),
	17:   uint16(1),
	18:   uint16(sym_Comment),
	19:   uint16(15),
	20:   uint16(2),
	22:   uint16(anon_sym_RBRACK_RBRACK_GT),
	23:   uint16(45),
	24:   uint16(2),
	25:   uint16(sym_GEDecl),
	26:   uint16(sym_PEDecl),
	27:   uint16(36),
	28:   uint16(5),
	29:   uint16(sym_elementdecl),
	30:   uint16(sym_AttlistDecl),
	31:   uint16(sym__EntityDecl),
	32:   uint16(sym_NotationDecl),
	33:   uint16(sym_PI),
	34:   uint16(2),
	35:   uint16(6),
	36:   uint16(sym__extSubsetDecl),
	37:   uint16(sym_conditionalSect),
	38:   uint16(sym__markupdecl),
	39:   uint16(sym__DeclSep),
	40:   uint16(sym_PEReference),
	41:   uint16(aux_sym_extSubset_repeat1),
	42:   uint16(10),
	43:   uint16(5),
	44:   uint16(1),
	45:   uint16(anon_sym_LT_BANG_LBRACK),
	46:   uint16(7),
	47:   uint16(1),
	48:   uint16(anon_sym_LT_BANG),
	49:   uint16(9),
	50:   uint16(1),
	51:   uint16(anon_sym_PERCENT),
	52:   uint16(13),
	53:   uint16(1),
	54:   uint16(sym_Comment),
	55:   uint16(35),
	56:   uint16(1),
	57:   uint16(anon_sym_LT_QMARK),
	58:   uint16(37),
	59:   uint16(1),
	60:   uint16(anon_sym_RBRACK_RBRACK_GT),
	61:   uint16(39),
	62:   uint16(1),
	63:   uint16(sym__S),
	64:   uint16(45),
	65:   uint16(2),
	66:   uint16(sym_GEDecl),
	67:   uint16(sym_PEDecl),
	68:   uint16(36),
	69:   uint16(5),
	70:   uint16(sym_elementdecl),
	71:   uint16(sym_AttlistDecl),
	72:   uint16(sym__EntityDecl),
	73:   uint16(sym_NotationDecl),
	74:   uint16(sym_PI),
	75:   uint16(5),
	76:   uint16(6),
	77:   uint16(sym__extSubsetDecl),
	78:   uint16(sym_conditionalSect),
	79:   uint16(sym__markupdecl),
	80:   uint16(sym__DeclSep),
	81:   uint16(sym_PEReference),
	82:   uint16(aux_sym_extSubset_repeat1),
	83:   uint16(10),
	84:   uint16(5),
	85:   uint16(1),
	86:   uint16(anon_sym_LT_BANG_LBRACK),
	87:   uint16(7),
	88:   uint16(1),
	89:   uint16(anon_sym_LT_BANG),
	90:   uint16(9),
	91:   uint16(1),
	92:   uint16(anon_sym_PERCENT),
	93:   uint16(13),
	94:   uint16(1),
	95:   uint16(sym_Comment),
	96:   uint16(35),
	97:   uint16(1),
	98:   uint16(anon_sym_LT_QMARK),
	99:   uint16(41),
	100:  uint16(1),
	102:  uint16(43),
	103:  uint16(1),
	104:  uint16(sym__S),
	105:  uint16(45),
	106:  uint16(2),
	107:  uint16(sym_GEDecl),
	108:  uint16(sym_PEDecl),
	109:  uint16(36),
	110:  uint16(5),
	111:  uint16(sym_elementdecl),
	112:  uint16(sym_AttlistDecl),
	113:  uint16(sym__EntityDecl),
	114:  uint16(sym_NotationDecl),
	115:  uint16(sym_PI),
	116:  uint16(2),
	117:  uint16(6),
	118:  uint16(sym__extSubsetDecl),
	119:  uint16(sym_conditionalSect),
	120:  uint16(sym__markupdecl),
	121:  uint16(sym__DeclSep),
	122:  uint16(sym_PEReference),
	123:  uint16(aux_sym_extSubset_repeat1),
	124:  uint16(10),
	125:  uint16(5),
	126:  uint16(1),
	127:  uint16(anon_sym_LT_BANG_LBRACK),
	128:  uint16(7),
	129:  uint16(1),
	130:  uint16(anon_sym_LT_BANG),
	131:  uint16(9),
	132:  uint16(1),
	133:  uint16(anon_sym_PERCENT),
	134:  uint16(13),
	135:  uint16(1),
	136:  uint16(sym_Comment),
	137:  uint16(35),
	138:  uint16(1),
	139:  uint16(anon_sym_LT_QMARK),
	140:  uint16(43),
	141:  uint16(1),
	142:  uint16(sym__S),
	143:  uint16(45),
	144:  uint16(1),
	145:  uint16(anon_sym_RBRACK_RBRACK_GT),
	146:  uint16(45),
	147:  uint16(2),
	148:  uint16(sym_GEDecl),
	149:  uint16(sym_PEDecl),
	150:  uint16(36),
	151:  uint16(5),
	152:  uint16(sym_elementdecl),
	153:  uint16(sym_AttlistDecl),
	154:  uint16(sym__EntityDecl),
	155:  uint16(sym_NotationDecl),
	156:  uint16(sym_PI),
	157:  uint16(2),
	158:  uint16(6),
	159:  uint16(sym__extSubsetDecl),
	160:  uint16(sym_conditionalSect),
	161:  uint16(sym__markupdecl),
	162:  uint16(sym__DeclSep),
	163:  uint16(sym_PEReference),
	164:  uint16(aux_sym_extSubset_repeat1),
	165:  uint16(10),
	166:  uint16(5),
	167:  uint16(1),
	168:  uint16(anon_sym_LT_BANG_LBRACK),
	169:  uint16(7),
	170:  uint16(1),
	171:  uint16(anon_sym_LT_BANG),
	172:  uint16(9),
	173:  uint16(1),
	174:  uint16(anon_sym_PERCENT),
	175:  uint16(13),
	176:  uint16(1),
	177:  uint16(sym_Comment),
	178:  uint16(35),
	179:  uint16(1),
	180:  uint16(anon_sym_LT_QMARK),
	181:  uint16(43),
	182:  uint16(1),
	183:  uint16(sym__S),
	184:  uint16(47),
	185:  uint16(1),
	187:  uint16(45),
	188:  uint16(2),
	189:  uint16(sym_GEDecl),
	190:  uint16(sym_PEDecl),
	191:  uint16(36),
	192:  uint16(5),
	193:  uint16(sym_elementdecl),
	194:  uint16(sym_AttlistDecl),
	195:  uint16(sym__EntityDecl),
	196:  uint16(sym_NotationDecl),
	197:  uint16(sym_PI),
	198:  uint16(2),
	199:  uint16(6),
	200:  uint16(sym__extSubsetDecl),
	201:  uint16(sym_conditionalSect),
	202:  uint16(sym__markupdecl),
	203:  uint16(sym__DeclSep),
	204:  uint16(sym_PEReference),
	205:  uint16(aux_sym_extSubset_repeat1),
	206:  uint16(10),
	207:  uint16(5),
	208:  uint16(1),
	209:  uint16(anon_sym_LT_BANG_LBRACK),
	210:  uint16(7),
	211:  uint16(1),
	212:  uint16(anon_sym_LT_BANG),
	213:  uint16(9),
	214:  uint16(1),
	215:  uint16(anon_sym_PERCENT),
	216:  uint16(13),
	217:  uint16(1),
	218:  uint16(sym_Comment),
	219:  uint16(35),
	220:  uint16(1),
	221:  uint16(anon_sym_LT_QMARK),
	222:  uint16(49),
	223:  uint16(1),
	224:  uint16(anon_sym_RBRACK_RBRACK_GT),
	225:  uint16(51),
	226:  uint16(1),
	227:  uint16(sym__S),
	228:  uint16(45),
	229:  uint16(2),
	230:  uint16(sym_GEDecl),
	231:  uint16(sym_PEDecl),
	232:  uint16(36),
	233:  uint16(5),
	234:  uint16(sym_elementdecl),
	235:  uint16(sym_AttlistDecl),
	236:  uint16(sym__EntityDecl),
	237:  uint16(sym_NotationDecl),
	238:  uint16(sym_PI),
	239:  uint16(8),
	240:  uint16(6),
	241:  uint16(sym__extSubsetDecl),
	242:  uint16(sym_conditionalSect),
	243:  uint16(sym__markupdecl),
	244:  uint16(sym__DeclSep),
	245:  uint16(sym_PEReference),
	246:  uint16(aux_sym_extSubset_repeat1),
	247:  uint16(10),
	248:  uint16(5),
	249:  uint16(1),
	250:  uint16(anon_sym_LT_BANG_LBRACK),
	251:  uint16(7),
	252:  uint16(1),
	253:  uint16(anon_sym_LT_BANG),
	254:  uint16(9),
	255:  uint16(1),
	256:  uint16(anon_sym_PERCENT),
	257:  uint16(13),
	258:  uint16(1),
	259:  uint16(sym_Comment),
	260:  uint16(35),
	261:  uint16(1),
	262:  uint16(anon_sym_LT_QMARK),
	263:  uint16(37),
	264:  uint16(1),
	265:  uint16(anon_sym_RBRACK_RBRACK_GT),
	266:  uint16(43),
	267:  uint16(1),
	268:  uint16(sym__S),
	269:  uint16(45),
	270:  uint16(2),
	271:  uint16(sym_GEDecl),
	272:  uint16(sym_PEDecl),
	273:  uint16(36),
	274:  uint16(5),
	275:  uint16(sym_elementdecl),
	276:  uint16(sym_AttlistDecl),
	277:  uint16(sym__EntityDecl),
	278:  uint16(sym_NotationDecl),
	279:  uint16(sym_PI),
	280:  uint16(2),
	281:  uint16(6),
	282:  uint16(sym__extSubsetDecl),
	283:  uint16(sym_conditionalSect),
	284:  uint16(sym__markupdecl),
	285:  uint16(sym__DeclSep),
	286:  uint16(sym_PEReference),
	287:  uint16(aux_sym_extSubset_repeat1),
	288:  uint16(10),
	289:  uint16(5),
	290:  uint16(1),
	291:  uint16(anon_sym_LT_BANG_LBRACK),
	292:  uint16(7),
	293:  uint16(1),
	294:  uint16(anon_sym_LT_BANG),
	295:  uint16(9),
	296:  uint16(1),
	297:  uint16(anon_sym_PERCENT),
	298:  uint16(13),
	299:  uint16(1),
	300:  uint16(sym_Comment),
	301:  uint16(35),
	302:  uint16(1),
	303:  uint16(anon_sym_LT_QMARK),
	304:  uint16(45),
	305:  uint16(1),
	306:  uint16(anon_sym_RBRACK_RBRACK_GT),
	307:  uint16(53),
	308:  uint16(1),
	309:  uint16(sym__S),
	310:  uint16(45),
	311:  uint16(2),
	312:  uint16(sym_GEDecl),
	313:  uint16(sym_PEDecl),
	314:  uint16(36),
	315:  uint16(5),
	316:  uint16(sym_elementdecl),
	317:  uint16(sym_AttlistDecl),
	318:  uint16(sym__EntityDecl),
	319:  uint16(sym_NotationDecl),
	320:  uint16(sym_PI),
	321:  uint16(10),
	322:  uint16(6),
	323:  uint16(sym__extSubsetDecl),
	324:  uint16(sym_conditionalSect),
	325:  uint16(sym__markupdecl),
	326:  uint16(sym__DeclSep),
	327:  uint16(sym_PEReference),
	328:  uint16(aux_sym_extSubset_repeat1),
	329:  uint16(10),
	330:  uint16(5),
	331:  uint16(1),
	332:  uint16(anon_sym_LT_BANG_LBRACK),
	333:  uint16(7),
	334:  uint16(1),
	335:  uint16(anon_sym_LT_BANG),
	336:  uint16(9),
	337:  uint16(1),
	338:  uint16(anon_sym_PERCENT),
	339:  uint16(13),
	340:  uint16(1),
	341:  uint16(sym_Comment),
	342:  uint16(35),
	343:  uint16(1),
	344:  uint16(anon_sym_LT_QMARK),
	345:  uint16(43),
	346:  uint16(1),
	347:  uint16(sym__S),
	348:  uint16(55),
	349:  uint16(1),
	350:  uint16(anon_sym_RBRACK_RBRACK_GT),
	351:  uint16(45),
	352:  uint16(2),
	353:  uint16(sym_GEDecl),
	354:  uint16(sym_PEDecl),
	355:  uint16(36),
	356:  uint16(5),
	357:  uint16(sym_elementdecl),
	358:  uint16(sym_AttlistDecl),
	359:  uint16(sym__EntityDecl),
	360:  uint16(sym_NotationDecl),
	361:  uint16(sym_PI),
	362:  uint16(2),
	363:  uint16(6),
	364:  uint16(sym__extSubsetDecl),
	365:  uint16(sym_conditionalSect),
	366:  uint16(sym__markupdecl),
	367:  uint16(sym__DeclSep),
	368:  uint16(sym_PEReference),
	369:  uint16(aux_sym_extSubset_repeat1),
	370:  uint16(9),
	371:  uint16(5),
	372:  uint16(1),
	373:  uint16(anon_sym_LT_BANG_LBRACK),
	374:  uint16(7),
	375:  uint16(1),
	376:  uint16(anon_sym_LT_BANG),
	377:  uint16(9),
	378:  uint16(1),
	379:  uint16(anon_sym_PERCENT),
	380:  uint16(13),
	381:  uint16(1),
	382:  uint16(sym_Comment),
	383:  uint16(35),
	384:  uint16(1),
	385:  uint16(anon_sym_LT_QMARK),
	386:  uint16(57),
	387:  uint16(1),
	388:  uint16(sym__S),
	389:  uint16(45),
	390:  uint16(2),
	391:  uint16(sym_GEDecl),
	392:  uint16(sym_PEDecl),
	393:  uint16(36),
	394:  uint16(5),
	395:  uint16(sym_elementdecl),
	396:  uint16(sym_AttlistDecl),
	397:  uint16(sym__EntityDecl),
	398:  uint16(sym_NotationDecl),
	399:  uint16(sym_PI),
	400:  uint16(4),
	401:  uint16(6),
	402:  uint16(sym__extSubsetDecl),
	403:  uint16(sym_conditionalSect),
	404:  uint16(sym__markupdecl),
	405:  uint16(sym__DeclSep),
	406:  uint16(sym_PEReference),
	407:  uint16(aux_sym_extSubset_repeat1),
	408:  uint16(1),
	409:  uint16(59),
	410:  uint16(11),
	411:  uint16(anon_sym_LBRACK),
	412:  uint16(anon_sym_GT),
	413:  uint16(anon_sym_PIPE),
	414:  uint16(anon_sym_RPAREN),
	415:  uint16(anon_sym_STAR),
	416:  uint16(anon_sym_QMARK),
	417:  uint16(anon_sym_PLUS),
	418:  uint16(anon_sym_COMMA),
	419:  uint16(anon_sym_PERCENT),
	420:  uint16(sym__S),
	421:  uint16(sym_Name),
	422:  uint16(8),
	423:  uint16(61),
	424:  uint16(1),
	425:  uint16(anon_sym_PERCENT),
	426:  uint16(63),
	427:  uint16(1),
	428:  uint16(anon_sym_DQUOTE),
	429:  uint16(65),
	430:  uint16(1),
	431:  uint16(aux_sym_EntityValue_token1),
	432:  uint16(67),
	433:  uint16(1),
	434:  uint16(anon_sym_AMP),
	435:  uint16(69),
	436:  uint16(1),
	437:  uint16(anon_sym_AMP_POUND),
	438:  uint16(71),
	439:  uint16(1),
	440:  uint16(anon_sym_AMP_POUNDx),
	441:  uint16(91),
	442:  uint16(2),
	443:  uint16(sym_EntityRef),
	444:  uint16(sym_CharRef),
	445:  uint16(17),
	446:  uint16(3),
	447:  uint16(sym_PEReference),
	448:  uint16(sym__Reference),
	449:  uint16(aux_sym_EntityValue_repeat1),
	450:  uint16(8),
	451:  uint16(63),
	452:  uint16(1),
	453:  uint16(anon_sym_SQUOTE),
	454:  uint16(73),
	455:  uint16(1),
	456:  uint16(anon_sym_PERCENT),
	457:  uint16(75),
	458:  uint16(1),
	459:  uint16(aux_sym_EntityValue_token2),
	460:  uint16(77),
	461:  uint16(1),
	462:  uint16(anon_sym_AMP),
	463:  uint16(79),
	464:  uint16(1),
	465:  uint16(anon_sym_AMP_POUND),
	466:  uint16(81),
	467:  uint16(1),
	468:  uint16(anon_sym_AMP_POUNDx),
	469:  uint16(70),
	470:  uint16(2),
	471:  uint16(sym_EntityRef),
	472:  uint16(sym_CharRef),
	473:  uint16(18),
	474:  uint16(3),
	475:  uint16(sym_PEReference),
	476:  uint16(sym__Reference),
	477:  uint16(aux_sym_EntityValue_repeat2),
	478:  uint16(5),
	479:  uint16(65),
	480:  uint16(1),
	481:  uint16(aux_sym_Mixed_repeat1),
	482:  uint16(143),
	483:  uint16(1),
	484:  uint16(aux_sym_Mixed_repeat2),
	485:  uint16(165),
	486:  uint16(1),
	487:  uint16(sym_PEReference),
	488:  uint16(85),
	489:  uint16(3),
	490:  uint16(anon_sym_STAR),
	491:  uint16(anon_sym_QMARK),
	492:  uint16(anon_sym_PLUS),
	493:  uint16(83),
	494:  uint16(5),
	495:  uint16(anon_sym_PIPE),
	496:  uint16(anon_sym_RPAREN),
	497:  uint16(anon_sym_COMMA),
	498:  uint16(anon_sym_PERCENT),
	499:  uint16(sym__S),
	500:  uint16(8),
	501:  uint16(87),
	502:  uint16(1),
	503:  uint16(anon_sym_PERCENT),
	504:  uint16(90),
	505:  uint16(1),
	506:  uint16(anon_sym_SQUOTE),
	507:  uint16(92),
	508:  uint16(1),
	509:  uint16(aux_sym_EntityValue_token2),
	510:  uint16(95),
	511:  uint16(1),
	512:  uint16(anon_sym_AMP),
	513:  uint16(98),
	514:  uint16(1),
	515:  uint16(anon_sym_AMP_POUND),
	516:  uint16(101),
	517:  uint16(1),
	518:  uint16(anon_sym_AMP_POUNDx),
	519:  uint16(70),
	520:  uint16(2),
	521:  uint16(sym_EntityRef),
	522:  uint16(sym_CharRef),
	523:  uint16(16),
	524:  uint16(3),
	525:  uint16(sym_PEReference),
	526:  uint16(sym__Reference),
	527:  uint16(aux_sym_EntityValue_repeat2),
	528:  uint16(8),
	529:  uint16(61),
	530:  uint16(1),
	531:  uint16(anon_sym_PERCENT),
	532:  uint16(67),
	533:  uint16(1),
	534:  uint16(anon_sym_AMP),
	535:  uint16(69),
	536:  uint16(1),
	537:  uint16(anon_sym_AMP_POUND),
	538:  uint16(71),
	539:  uint16(1),
	540:  uint16(anon_sym_AMP_POUNDx),
	541:  uint16(104),
	542:  uint16(1),
	543:  uint16(anon_sym_DQUOTE),
	544:  uint16(106),
	545:  uint16(1),
	546:  uint16(aux_sym_EntityValue_token1),
	547:  uint16(91),
	548:  uint16(2),
	549:  uint16(sym_EntityRef),
	550:  uint16(sym_CharRef),
	551:  uint16(20),
	552:  uint16(3),
	553:  uint16(sym_PEReference),
	554:  uint16(sym__Reference),
	555:  uint16(aux_sym_EntityValue_repeat1),
	556:  uint16(8),
	557:  uint16(73),
	558:  uint16(1),
	559:  uint16(anon_sym_PERCENT),
	560:  uint16(77),
	561:  uint16(1),
	562:  uint16(anon_sym_AMP),
	563:  uint16(79),
	564:  uint16(1),
	565:  uint16(anon_sym_AMP_POUND),
	566:  uint16(81),
	567:  uint16(1),
	568:  uint16(anon_sym_AMP_POUNDx),
	569:  uint16(104),
	570:  uint16(1),
	571:  uint16(anon_sym_SQUOTE),
	572:  uint16(108),
	573:  uint16(1),
	574:  uint16(aux_sym_EntityValue_token2),
	575:  uint16(70),
	576:  uint16(2),
	577:  uint16(sym_EntityRef),
	578:  uint16(sym_CharRef),
	579:  uint16(16),
	580:  uint16(3),
	581:  uint16(sym_PEReference),
	582:  uint16(sym__Reference),
	583:  uint16(aux_sym_EntityValue_repeat2),
	584:  uint16(5),
	585:  uint16(64),
	586:  uint16(1),
	587:  uint16(aux_sym_Mixed_repeat1),
	588:  uint16(132),
	589:  uint16(1),
	590:  uint16(aux_sym_Mixed_repeat2),
	591:  uint16(165),
	592:  uint16(1),
	593:  uint16(sym_PEReference),
	594:  uint16(85),
	595:  uint16(3),
	596:  uint16(anon_sym_STAR),
	597:  uint16(anon_sym_QMARK),
	598:  uint16(anon_sym_PLUS),
	599:  uint16(83),
	600:  uint16(5),
	601:  uint16(anon_sym_PIPE),
	602:  uint16(anon_sym_RPAREN),
	603:  uint16(anon_sym_COMMA),
	604:  uint16(anon_sym_PERCENT),
	605:  uint16(sym__S),
	606:  uint16(8),
	607:  uint16(110),
	608:  uint16(1),
	609:  uint16(anon_sym_PERCENT),
	610:  uint16(113),
	611:  uint16(1),
	612:  uint16(anon_sym_DQUOTE),
	613:  uint16(115),
	614:  uint16(1),
	615:  uint16(aux_sym_EntityValue_token1),
	616:  uint16(118),
	617:  uint16(1),
	618:  uint16(anon_sym_AMP),
	619:  uint16(121),
	620:  uint16(1),
	621:  uint16(anon_sym_AMP_POUND),
	622:  uint16(124),
	623:  uint16(1),
	624:  uint16(anon_sym_AMP_POUNDx),
	625:  uint16(91),
	626:  uint16(2),
	627:  uint16(sym_EntityRef),
	628:  uint16(sym_CharRef),
	629:  uint16(20),
	630:  uint16(3),
	631:  uint16(sym_PEReference),
	632:  uint16(sym__Reference),
	633:  uint16(aux_sym_EntityValue_repeat1),
	634:  uint16(7),
	635:  uint16(127),
	636:  uint16(1),
	637:  uint16(anon_sym_LPAREN),
	638:  uint16(131),
	639:  uint16(1),
	640:  uint16(anon_sym_NOTATION),
	641:  uint16(133),
	642:  uint16(1),
	643:  uint16(anon_sym_PERCENT),
	644:  uint16(218),
	645:  uint16(1),
	646:  uint16(sym__AttType),
	647:  uint16(129),
	648:  uint16(2),
	649:  uint16(sym_StringType),
	650:  uint16(sym_TokenizedType),
	651:  uint16(216),
	652:  uint16(2),
	653:  uint16(sym__EnumeratedType),
	654:  uint16(sym_PEReference),
	655:  uint16(219),
	656:  uint16(2),
	657:  uint16(sym_NotationType),
	658:  uint16(sym_Enumeration),
	659:  uint16(1),
	660:  uint16(135),
	661:  uint16(9),
	662:  uint16(anon_sym_GT),
	663:  uint16(anon_sym_PIPE),
	664:  uint16(anon_sym_RPAREN),
	665:  uint16(anon_sym_STAR),
	666:  uint16(anon_sym_QMARK),
	667:  uint16(anon_sym_PLUS),
	668:  uint16(anon_sym_COMMA),
	669:  uint16(anon_sym_PERCENT),
	670:  uint16(sym__S),
	671:  uint16(7),
	672:  uint16(137),
	673:  uint16(1),
	674:  uint16(anon_sym_SQUOTE),
	675:  uint16(139),
	676:  uint16(1),
	677:  uint16(anon_sym_AMP),
	678:  uint16(142),
	679:  uint16(1),
	680:  uint16(anon_sym_AMP_POUND),
	681:  uint16(145),
	682:  uint16(1),
	683:  uint16(anon_sym_AMP_POUNDx),
	684:  uint16(148),
	685:  uint16(1),
	686:  uint16(aux_sym_AttValue_token2),
	687:  uint16(23),
	688:  uint16(2),
	689:  uint16(sym__Reference),
	690:  uint16(aux_sym_AttValue_repeat2),
	691:  uint16(114),
	692:  uint16(2),
	693:  uint16(sym_EntityRef),
	694:  uint16(sym_CharRef),
	695:  uint16(6),
	696:  uint16(133),
	697:  uint16(1),
	698:  uint16(anon_sym_PERCENT),
	699:  uint16(153),
	700:  uint16(1),
	701:  uint16(anon_sym_LPAREN),
	702:  uint16(100),
	703:  uint16(1),
	704:  uint16(sym__choice),
	705:  uint16(203),
	706:  uint16(1),
	707:  uint16(sym_contentspec),
	708:  uint16(151),
	709:  uint16(2),
	710:  uint16(anon_sym_EMPTY),
	711:  uint16(anon_sym_ANY),
	712:  uint16(200),
	713:  uint16(3),
	714:  uint16(sym_Mixed),
	715:  uint16(sym_children),
	716:  uint16(sym_PEReference),
	717:  uint16(1),
	718:  uint16(155),
	719:  uint16(9),
	720:  uint16(anon_sym_GT),
	721:  uint16(anon_sym_PIPE),
	722:  uint16(anon_sym_RPAREN),
	723:  uint16(anon_sym_STAR),
	724:  uint16(anon_sym_QMARK),
	725:  uint16(anon_sym_PLUS),
	726:  uint16(anon_sym_COMMA),
	727:  uint16(anon_sym_PERCENT),
	728:  uint16(sym__S),
	729:  uint16(7),
	730:  uint16(133),
	731:  uint16(1),
	732:  uint16(anon_sym_PERCENT),
	733:  uint16(159),
	734:  uint16(1),
	735:  uint16(anon_sym_POUNDFIXED),
	736:  uint16(161),
	737:  uint16(1),
	738:  uint16(anon_sym_DQUOTE),
	739:  uint16(163),
	740:  uint16(1),
	741:  uint16(anon_sym_SQUOTE),
	742:  uint16(204),
	743:  uint16(1),
	744:  uint16(sym_DefaultDecl),
	745:  uint16(157),
	746:  uint16(2),
	747:  uint16(anon_sym_POUNDREQUIRED),
	748:  uint16(anon_sym_POUNDIMPLIED),
	749:  uint16(202),
	750:  uint16(2),
	751:  uint16(sym_PEReference),
	752:  uint16(sym_AttValue),
	753:  uint16(1),
	754:  uint16(165),
	755:  uint16(9),
	756:  uint16(anon_sym_GT),
	757:  uint16(anon_sym_PIPE),
	758:  uint16(anon_sym_RPAREN),
	759:  uint16(anon_sym_STAR),
	760:  uint16(anon_sym_QMARK),
	761:  uint16(anon_sym_PLUS),
	762:  uint16(anon_sym_COMMA),
	763:  uint16(anon_sym_PERCENT),
	764:  uint16(sym__S),
	765:  uint16(7),
	766:  uint16(167),
	767:  uint16(1),
	768:  uint16(anon_sym_DQUOTE),
	769:  uint16(169),
	770:  uint16(1),
	771:  uint16(anon_sym_AMP),
	772:  uint16(171),
	773:  uint16(1),
	774:  uint16(anon_sym_AMP_POUND),
	775:  uint16(173),
	776:  uint16(1),
	777:  uint16(anon_sym_AMP_POUNDx),
	778:  uint16(175),
	779:  uint16(1),
	780:  uint16(aux_sym_AttValue_token1),
	781:  uint16(31),
	782:  uint16(2),
	783:  uint16(sym__Reference),
	784:  uint16(aux_sym_AttValue_repeat1),
	785:  uint16(110),
	786:  uint16(2),
	787:  uint16(sym_EntityRef),
	788:  uint16(sym_CharRef),
	789:  uint16(7),
	790:  uint16(167),
	791:  uint16(1),
	792:  uint16(anon_sym_SQUOTE),
	793:  uint16(177),
	794:  uint16(1),
	795:  uint16(anon_sym_AMP),
	796:  uint16(179),
	797:  uint16(1),
	798:  uint16(anon_sym_AMP_POUND),
	799:  uint16(181),
	800:  uint16(1),
	801:  uint16(anon_sym_AMP_POUNDx),
	802:  uint16(183),
	803:  uint16(1),
	804:  uint16(aux_sym_AttValue_token2),
	805:  uint16(32),
	806:  uint16(2),
	807:  uint16(sym__Reference),
	808:  uint16(aux_sym_AttValue_repeat2),
	809:  uint16(114),
	810:  uint16(2),
	811:  uint16(sym_EntityRef),
	812:  uint16(sym_CharRef),
	813:  uint16(1),
	814:  uint16(185),
	815:  uint16(9),
	816:  uint16(anon_sym_GT),
	817:  uint16(anon_sym_PIPE),
	818:  uint16(anon_sym_RPAREN),
	819:  uint16(anon_sym_STAR),
	820:  uint16(anon_sym_QMARK),
	821:  uint16(anon_sym_PLUS),
	822:  uint16(anon_sym_COMMA),
	823:  uint16(anon_sym_PERCENT),
	824:  uint16(sym__S),
	825:  uint16(7),
	826:  uint16(169),
	827:  uint16(1),
	828:  uint16(anon_sym_AMP),
	829:  uint16(171),
	830:  uint16(1),
	831:  uint16(anon_sym_AMP_POUND),
	832:  uint16(173),
	833:  uint16(1),
	834:  uint16(anon_sym_AMP_POUNDx),
	835:  uint16(187),
	836:  uint16(1),
	837:  uint16(anon_sym_DQUOTE),
	838:  uint16(189),
	839:  uint16(1),
	840:  uint16(aux_sym_AttValue_token1),
	841:  uint16(34),
	842:  uint16(2),
	843:  uint16(sym__Reference),
	844:  uint16(aux_sym_AttValue_repeat1),
	845:  uint16(110),
	846:  uint16(2),
	847:  uint16(sym_EntityRef),
	848:  uint16(sym_CharRef),
	849:  uint16(7),
	850:  uint16(177),
	851:  uint16(1),
	852:  uint16(anon_sym_AMP),
	853:  uint16(179),
	854:  uint16(1),
	855:  uint16(anon_sym_AMP_POUND),
	856:  uint16(181),
	857:  uint16(1),
	858:  uint16(anon_sym_AMP_POUNDx),
	859:  uint16(187),
	860:  uint16(1),
	861:  uint16(anon_sym_SQUOTE),
	862:  uint16(191),
	863:  uint16(1),
	864:  uint16(aux_sym_AttValue_token2),
	865:  uint16(23),
	866:  uint16(2),
	867:  uint16(sym__Reference),
	868:  uint16(aux_sym_AttValue_repeat2),
	869:  uint16(114),
	870:  uint16(2),
	871:  uint16(sym_EntityRef),
	872:  uint16(sym_CharRef),
	873:  uint16(1),
	874:  uint16(193),
	875:  uint16(9),
	876:  uint16(anon_sym_GT),
	877:  uint16(anon_sym_PIPE),
	878:  uint16(anon_sym_RPAREN),
	879:  uint16(anon_sym_STAR),
	880:  uint16(anon_sym_QMARK),
	881:  uint16(anon_sym_PLUS),
	882:  uint16(anon_sym_COMMA),
	883:  uint16(anon_sym_PERCENT),
	884:  uint16(sym__S),
	885:  uint16(7),
	886:  uint16(195),
	887:  uint16(1),
	888:  uint16(anon_sym_DQUOTE),
	889:  uint16(197),
	890:  uint16(1),
	891:  uint16(anon_sym_AMP),
	892:  uint16(200),
	893:  uint16(1),
	894:  uint16(anon_sym_AMP_POUND),
	895:  uint16(203),
	896:  uint16(1),
	897:  uint16(anon_sym_AMP_POUNDx),
	898:  uint16(206),
	899:  uint16(1),
	900:  uint16(aux_sym_AttValue_token1),
	901:  uint16(34),
	902:  uint16(2),
	903:  uint16(sym__Reference),
	904:  uint16(aux_sym_AttValue_repeat1),
	905:  uint16(110),
	906:  uint16(2),
	907:  uint16(sym_EntityRef),
	908:  uint16(sym_CharRef),
	909:  uint16(2),
	910:  uint16(211),
	911:  uint16(1),
	912:  uint16(anon_sym_LT_BANG),
	913:  uint16(209),
	914:  uint16(7),
	915:  uint16(sym_Comment),
	917:  uint16(anon_sym_LT_QMARK),
	918:  uint16(anon_sym_LT_BANG_LBRACK),
	919:  uint16(anon_sym_RBRACK_RBRACK_GT),
	920:  uint16(anon_sym_PERCENT),
	921:  uint16(sym__S),
	922:  uint16(2),
	923:  uint16(215),
	924:  uint16(1),
	925:  uint16(anon_sym_LT_BANG),
	926:  uint16(213),
	927:  uint16(7),
	928:  uint16(sym_Comment),
	930:  uint16(anon_sym_LT_QMARK),
	931:  uint16(anon_sym_LT_BANG_LBRACK),
	932:  uint16(anon_sym_RBRACK_RBRACK_GT),
	933:  uint16(anon_sym_PERCENT),
	934:  uint16(sym__S),
	935:  uint16(2),
	936:  uint16(85),
	937:  uint16(3),
	938:  uint16(anon_sym_STAR),
	939:  uint16(anon_sym_QMARK),
	940:  uint16(anon_sym_PLUS),
	941:  uint16(83),
	942:  uint16(5),
	943:  uint16(anon_sym_PIPE),
	944:  uint16(anon_sym_RPAREN),
	945:  uint16(anon_sym_COMMA),
	946:  uint16(anon_sym_PERCENT),
	947:  uint16(sym__S),
	948:  uint16(6),
	949:  uint16(133),
	950:  uint16(1),
	951:  uint16(anon_sym_PERCENT),
	952:  uint16(219),
	953:  uint16(1),
	954:  uint16(anon_sym_RPAREN),
	955:  uint16(221),
	956:  uint16(1),
	957:  uint16(sym__S),
	958:  uint16(61),
	959:  uint16(1),
	960:  uint16(aux_sym__choice_repeat1),
	961:  uint16(217),
	962:  uint16(2),
	963:  uint16(anon_sym_PIPE),
	964:  uint16(anon_sym_COMMA),
	965:  uint16(113),
	966:  uint16(2),
	967:  uint16(sym_PEReference),
	968:  uint16(aux_sym__choice_repeat2),
	969:  uint16(2),
	970:  uint16(225),
	971:  uint16(1),
	972:  uint16(anon_sym_LT_BANG),
	973:  uint16(223),
	974:  uint16(7),
	975:  uint16(sym_Comment),
	977:  uint16(anon_sym_LT_QMARK),
	978:  uint16(anon_sym_LT_BANG_LBRACK),
	979:  uint16(anon_sym_RBRACK_RBRACK_GT),
	980:  uint16(anon_sym_PERCENT),
	981:  uint16(sym__S),
	982:  uint16(2),
	983:  uint16(229),
	984:  uint16(1),
	985:  uint16(anon_sym_LT_BANG),
	986:  uint16(227),
	987:  uint16(7),
	988:  uint16(sym_Comment),
	990:  uint16(anon_sym_LT_QMARK),
	991:  uint16(anon_sym_LT_BANG_LBRACK),
	992:  uint16(anon_sym_RBRACK_RBRACK_GT),
	993:  uint16(anon_sym_PERCENT),
	994:  uint16(sym__S),
	995:  uint16(2),
	996:  uint16(233),
	997:  uint16(1),
	998:  uint16(anon_sym_LT_BANG),
	999:  uint16(231),
	1000: uint16(7),
	1001: uint16(sym_Comment),
	1003: uint16(anon_sym_LT_QMARK),
	1004: uint16(anon_sym_LT_BANG_LBRACK),
	1005: uint16(anon_sym_RBRACK_RBRACK_GT),
	1006: uint16(anon_sym_PERCENT),
	1007: uint16(sym__S),
	1008: uint16(2),
	1009: uint16(235),
	1010: uint16(1),
	1011: uint16(anon_sym_LT_BANG),
	1012: uint16(59),
	1013: uint16(7),
	1014: uint16(sym_Comment),
	1016: uint16(anon_sym_LT_QMARK),
	1017: uint16(anon_sym_LT_BANG_LBRACK),
	1018: uint16(anon_sym_RBRACK_RBRACK_GT),
	1019: uint16(anon_sym_PERCENT),
	1020: uint16(sym__S),
	1021: uint16(2),
	1022: uint16(239),
	1023: uint16(1),
	1024: uint16(anon_sym_LT_BANG),
	1025: uint16(237),
	1026: uint16(7),
	1027: uint16(sym_Comment),
	1029: uint16(anon_sym_LT_QMARK),
	1030: uint16(anon_sym_LT_BANG_LBRACK),
	1031: uint16(anon_sym_RBRACK_RBRACK_GT),
	1032: uint16(anon_sym_PERCENT),
	1033: uint16(sym__S),
	1034: uint16(6),
	1035: uint16(133),
	1036: uint16(1),
	1037: uint16(anon_sym_PERCENT),
	1038: uint16(241),
	1039: uint16(1),
	1040: uint16(anon_sym_RPAREN),
	1041: uint16(243),
	1042: uint16(1),
	1043: uint16(sym__S),
	1044: uint16(51),
	1045: uint16(1),
	1046: uint16(aux_sym__choice_repeat1),
	1047: uint16(217),
	1048: uint16(2),
	1049: uint16(anon_sym_PIPE),
	1050: uint16(anon_sym_COMMA),
	1051: uint16(92),
	1052: uint16(2),
	1053: uint16(sym_PEReference),
	1054: uint16(aux_sym__choice_repeat2),
	1055: uint16(2),
	1056: uint16(247),
	1057: uint16(1),
	1058: uint16(anon_sym_LT_BANG),
	1059: uint16(245),
	1060: uint16(7),
	1061: uint16(sym_Comment),
	1063: uint16(anon_sym_LT_QMARK),
	1064: uint16(anon_sym_LT_BANG_LBRACK),
	1065: uint16(anon_sym_RBRACK_RBRACK_GT),
	1066: uint16(anon_sym_PERCENT),
	1067: uint16(sym__S),
	1068: uint16(2),
	1069: uint16(251),
	1070: uint16(1),
	1071: uint16(anon_sym_LT_BANG),
	1072: uint16(249),
	1073: uint16(7),
	1074: uint16(sym_Comment),
	1076: uint16(anon_sym_LT_QMARK),
	1077: uint16(anon_sym_LT_BANG_LBRACK),
	1078: uint16(anon_sym_RBRACK_RBRACK_GT),
	1079: uint16(anon_sym_PERCENT),
	1080: uint16(sym__S),
	1081: uint16(2),
	1082: uint16(255),
	1083: uint16(1),
	1084: uint16(anon_sym_LT_BANG),
	1085: uint16(253),
	1086: uint16(7),
	1087: uint16(sym_Comment),
	1089: uint16(anon_sym_LT_QMARK),
	1090: uint16(anon_sym_LT_BANG_LBRACK),
	1091: uint16(anon_sym_RBRACK_RBRACK_GT),
	1092: uint16(anon_sym_PERCENT),
	1093: uint16(sym__S),
	1094: uint16(2),
	1095: uint16(259),
	1096: uint16(1),
	1097: uint16(anon_sym_LT_BANG),
	1098: uint16(257),
	1099: uint16(7),
	1100: uint16(sym_Comment),
	1102: uint16(anon_sym_LT_QMARK),
	1103: uint16(anon_sym_LT_BANG_LBRACK),
	1104: uint16(anon_sym_RBRACK_RBRACK_GT),
	1105: uint16(anon_sym_PERCENT),
	1106: uint16(sym__S),
	1107: uint16(2),
	1108: uint16(263),
	1109: uint16(1),
	1110: uint16(anon_sym_LT_BANG),
	1111: uint16(261),
	1112: uint16(7),
	1113: uint16(sym_Comment),
	1115: uint16(anon_sym_LT_QMARK),
	1116: uint16(anon_sym_LT_BANG_LBRACK),
	1117: uint16(anon_sym_RBRACK_RBRACK_GT),
	1118: uint16(anon_sym_PERCENT),
	1119: uint16(sym__S),
	1120: uint16(2),
	1121: uint16(267),
	1122: uint16(1),
	1123: uint16(anon_sym_LT_BANG),
	1124: uint16(265),
	1125: uint16(7),
	1126: uint16(sym_Comment),
	1128: uint16(anon_sym_LT_QMARK),
	1129: uint16(anon_sym_LT_BANG_LBRACK),
	1130: uint16(anon_sym_RBRACK_RBRACK_GT),
	1131: uint16(anon_sym_PERCENT),
	1132: uint16(sym__S),
	1133: uint16(6),
	1134: uint16(133),
	1135: uint16(1),
	1136: uint16(anon_sym_PERCENT),
	1137: uint16(269),
	1138: uint16(1),
	1139: uint16(anon_sym_RPAREN),
	1140: uint16(271),
	1141: uint16(1),
	1142: uint16(sym__S),
	1143: uint16(73),
	1144: uint16(1),
	1145: uint16(aux_sym__choice_repeat1),
	1146: uint16(217),
	1147: uint16(2),
	1148: uint16(anon_sym_PIPE),
	1149: uint16(anon_sym_COMMA),
	1150: uint16(99),
	1151: uint16(2),
	1152: uint16(sym_PEReference),
	1153: uint16(aux_sym__choice_repeat2),
	1154: uint16(2),
	1155: uint16(275),
	1156: uint16(1),
	1157: uint16(anon_sym_LT_BANG),
	1158: uint16(273),
	1159: uint16(7),
	1160: uint16(sym_Comment),
	1162: uint16(anon_sym_LT_QMARK),
	1163: uint16(anon_sym_LT_BANG_LBRACK),
	1164: uint16(anon_sym_RBRACK_RBRACK_GT),
	1165: uint16(anon_sym_PERCENT),
	1166: uint16(sym__S),
	1167: uint16(2),
	1168: uint16(279),
	1169: uint16(1),
	1170: uint16(anon_sym_LT_BANG),
	1171: uint16(277),
	1172: uint16(7),
	1173: uint16(sym_Comment),
	1175: uint16(anon_sym_LT_QMARK),
	1176: uint16(anon_sym_LT_BANG_LBRACK),
	1177: uint16(anon_sym_RBRACK_RBRACK_GT),
	1178: uint16(anon_sym_PERCENT),
	1179: uint16(sym__S),
	1180: uint16(2),
	1181: uint16(283),
	1182: uint16(1),
	1183: uint16(anon_sym_LT_BANG),
	1184: uint16(281),
	1185: uint16(7),
	1186: uint16(sym_Comment),
	1188: uint16(anon_sym_LT_QMARK),
	1189: uint16(anon_sym_LT_BANG_LBRACK),
	1190: uint16(anon_sym_RBRACK_RBRACK_GT),
	1191: uint16(anon_sym_PERCENT),
	1192: uint16(sym__S),
	1193: uint16(2),
	1194: uint16(287),
	1195: uint16(1),
	1196: uint16(anon_sym_LT_BANG),
	1197: uint16(285),
	1198: uint16(7),
	1199: uint16(sym_Comment),
	1201: uint16(anon_sym_LT_QMARK),
	1202: uint16(anon_sym_LT_BANG_LBRACK),
	1203: uint16(anon_sym_RBRACK_RBRACK_GT),
	1204: uint16(anon_sym_PERCENT),
	1205: uint16(sym__S),
	1206: uint16(2),
	1207: uint16(291),
	1208: uint16(1),
	1209: uint16(anon_sym_LT_BANG),
	1210: uint16(289),
	1211: uint16(7),
	1212: uint16(sym_Comment),
	1214: uint16(anon_sym_LT_QMARK),
	1215: uint16(anon_sym_LT_BANG_LBRACK),
	1216: uint16(anon_sym_RBRACK_RBRACK_GT),
	1217: uint16(anon_sym_PERCENT),
	1218: uint16(sym__S),
	1219: uint16(2),
	1220: uint16(295),
	1221: uint16(1),
	1222: uint16(anon_sym_LT_BANG),
	1223: uint16(293),
	1224: uint16(7),
	1225: uint16(sym_Comment),
	1227: uint16(anon_sym_LT_QMARK),
	1228: uint16(anon_sym_LT_BANG_LBRACK),
	1229: uint16(anon_sym_RBRACK_RBRACK_GT),
	1230: uint16(anon_sym_PERCENT),
	1231: uint16(sym__S),
	1232: uint16(2),
	1233: uint16(299),
	1234: uint16(1),
	1235: uint16(anon_sym_LT_BANG),
	1236: uint16(297),
	1237: uint16(7),
	1238: uint16(sym_Comment),
	1240: uint16(anon_sym_LT_QMARK),
	1241: uint16(anon_sym_LT_BANG_LBRACK),
	1242: uint16(anon_sym_RBRACK_RBRACK_GT),
	1243: uint16(anon_sym_PERCENT),
	1244: uint16(sym__S),
	1245: uint16(8),
	1246: uint16(133),
	1247: uint16(1),
	1248: uint16(anon_sym_PERCENT),
	1249: uint16(301),
	1250: uint16(1),
	1251: uint16(sym_Name),
	1252: uint16(303),
	1253: uint16(1),
	1254: uint16(anon_sym_LPAREN),
	1255: uint16(305),
	1256: uint16(1),
	1257: uint16(anon_sym_POUNDPCDATA),
	1258: uint16(307),
	1259: uint16(1),
	1260: uint16(sym__S),
	1261: uint16(15),
	1262: uint16(1),
	1263: uint16(sym_PEReference),
	1264: uint16(37),
	1265: uint16(1),
	1266: uint16(sym__choice),
	1267: uint16(38),
	1268: uint16(1),
	1269: uint16(sym__cp),
	1270: uint16(2),
	1271: uint16(311),
	1272: uint16(1),
	1273: uint16(anon_sym_LT_BANG),
	1274: uint16(309),
	1275: uint16(7),
	1276: uint16(sym_Comment),
	1278: uint16(anon_sym_LT_QMARK),
	1279: uint16(anon_sym_LT_BANG_LBRACK),
	1280: uint16(anon_sym_RBRACK_RBRACK_GT),
	1281: uint16(anon_sym_PERCENT),
	1282: uint16(sym__S),
	1283: uint16(6),
	1284: uint16(133),
	1285: uint16(1),
	1286: uint16(anon_sym_PERCENT),
	1287: uint16(241),
	1288: uint16(1),
	1289: uint16(anon_sym_RPAREN),
	1290: uint16(243),
	1291: uint16(1),
	1292: uint16(sym__S),
	1293: uint16(73),
	1294: uint16(1),
	1295: uint16(aux_sym__choice_repeat1),
	1296: uint16(217),
	1297: uint16(2),
	1298: uint16(anon_sym_PIPE),
	1299: uint16(anon_sym_COMMA),
	1300: uint16(92),
	1301: uint16(2),
	1302: uint16(sym_PEReference),
	1303: uint16(aux_sym__choice_repeat2),
	1304: uint16(6),
	1305: uint16(133),
	1306: uint16(1),
	1307: uint16(anon_sym_PERCENT),
	1308: uint16(301),
	1309: uint16(1),
	1310: uint16(sym_Name),
	1311: uint16(303),
	1312: uint16(1),
	1313: uint16(anon_sym_LPAREN),
	1314: uint16(313),
	1315: uint16(1),
	1316: uint16(sym__S),
	1317: uint16(38),
	1318: uint16(1),
	1319: uint16(sym__cp),
	1320: uint16(37),
	1321: uint16(2),
	1322: uint16(sym__choice),
	1323: uint16(sym_PEReference),
	1324: uint16(7),
	1325: uint16(133),
	1326: uint16(1),
	1327: uint16(anon_sym_PERCENT),
	1328: uint16(315),
	1329: uint16(1),
	1330: uint16(anon_sym_PIPE),
	1331: uint16(317),
	1332: uint16(1),
	1333: uint16(anon_sym_RPAREN),
	1334: uint16(319),
	1335: uint16(1),
	1336: uint16(sym__S),
	1337: uint16(65),
	1338: uint16(1),
	1339: uint16(aux_sym_Mixed_repeat1),
	1340: uint16(143),
	1341: uint16(1),
	1342: uint16(aux_sym_Mixed_repeat2),
	1343: uint16(165),
	1344: uint16(1),
	1345: uint16(sym_PEReference),
	1346: uint16(7),
	1347: uint16(133),
	1348: uint16(1),
	1349: uint16(anon_sym_PERCENT),
	1350: uint16(315),
	1351: uint16(1),
	1352: uint16(anon_sym_PIPE),
	1353: uint16(321),
	1354: uint16(1),
	1355: uint16(anon_sym_RPAREN),
	1356: uint16(323),
	1357: uint16(1),
	1358: uint16(sym__S),
	1359: uint16(96),
	1360: uint16(1),
	1361: uint16(aux_sym_Mixed_repeat1),
	1362: uint16(120),
	1363: uint16(1),
	1364: uint16(aux_sym_Mixed_repeat2),
	1365: uint16(165),
	1366: uint16(1),
	1367: uint16(sym_PEReference),
	1368: uint16(7),
	1369: uint16(133),
	1370: uint16(1),
	1371: uint16(anon_sym_PERCENT),
	1372: uint16(315),
	1373: uint16(1),
	1374: uint16(anon_sym_PIPE),
	1375: uint16(325),
	1376: uint16(1),
	1377: uint16(anon_sym_RPAREN),
	1378: uint16(327),
	1379: uint16(1),
	1380: uint16(sym__S),
	1381: uint16(96),
	1382: uint16(1),
	1383: uint16(aux_sym_Mixed_repeat1),
	1384: uint16(134),
	1385: uint16(1),
	1386: uint16(aux_sym_Mixed_repeat2),
	1387: uint16(165),
	1388: uint16(1),
	1389: uint16(sym_PEReference),
	1390: uint16(6),
	1391: uint16(133),
	1392: uint16(1),
	1393: uint16(anon_sym_PERCENT),
	1394: uint16(301),
	1395: uint16(1),
	1396: uint16(sym_Name),
	1397: uint16(303),
	1398: uint16(1),
	1399: uint16(anon_sym_LPAREN),
	1400: uint16(329),
	1401: uint16(1),
	1402: uint16(sym__S),
	1403: uint16(103),
	1404: uint16(1),
	1405: uint16(sym__cp),
	1406: uint16(37),
	1407: uint16(2),
	1408: uint16(sym__choice),
	1409: uint16(sym_PEReference),
	1410: uint16(7),
	1411: uint16(133),
	1412: uint16(1),
	1413: uint16(anon_sym_PERCENT),
	1414: uint16(315),
	1415: uint16(1),
	1416: uint16(anon_sym_PIPE),
	1417: uint16(331),
	1418: uint16(1),
	1419: uint16(anon_sym_RPAREN),
	1420: uint16(333),
	1421: uint16(1),
	1422: uint16(sym__S),
	1423: uint16(64),
	1424: uint16(1),
	1425: uint16(aux_sym_Mixed_repeat1),
	1426: uint16(132),
	1427: uint16(1),
	1428: uint16(aux_sym_Mixed_repeat2),
	1429: uint16(165),
	1430: uint16(1),
	1431: uint16(sym_PEReference),
	1432: uint16(7),
	1433: uint16(133),
	1434: uint16(1),
	1435: uint16(anon_sym_PERCENT),
	1436: uint16(301),
	1437: uint16(1),
	1438: uint16(sym_Name),
	1439: uint16(303),
	1440: uint16(1),
	1441: uint16(anon_sym_LPAREN),
	1442: uint16(335),
	1443: uint16(1),
	1444: uint16(anon_sym_POUNDPCDATA),
	1445: uint16(19),
	1446: uint16(1),
	1447: uint16(sym_PEReference),
	1448: uint16(37),
	1449: uint16(1),
	1450: uint16(sym__choice),
	1451: uint16(44),
	1452: uint16(1),
	1453: uint16(sym__cp),
	1454: uint16(6),
	1455: uint16(133),
	1456: uint16(1),
	1457: uint16(anon_sym_PERCENT),
	1458: uint16(301),
	1459: uint16(1),
	1460: uint16(sym_Name),
	1461: uint16(303),
	1462: uint16(1),
	1463: uint16(anon_sym_LPAREN),
	1464: uint16(337),
	1465: uint16(1),
	1466: uint16(sym__S),
	1467: uint16(107),
	1468: uint16(1),
	1469: uint16(sym__cp),
	1470: uint16(37),
	1471: uint16(2),
	1472: uint16(sym__choice),
	1473: uint16(sym_PEReference),
	1474: uint16(2),
	1475: uint16(341),
	1476: uint16(2),
	1477: uint16(anon_sym_AMP),
	1478: uint16(anon_sym_AMP_POUND),
	1479: uint16(339),
	1480: uint16(4),
	1481: uint16(anon_sym_PERCENT),
	1482: uint16(anon_sym_SQUOTE),
	1483: uint16(aux_sym_EntityValue_token2),
	1484: uint16(anon_sym_AMP_POUNDx),
	1485: uint16(2),
	1486: uint16(235),
	1487: uint16(2),
	1488: uint16(anon_sym_AMP),
	1489: uint16(anon_sym_AMP_POUND),
	1490: uint16(59),
	1491: uint16(4),
	1492: uint16(anon_sym_PERCENT),
	1493: uint16(anon_sym_SQUOTE),
	1494: uint16(aux_sym_EntityValue_token2),
	1495: uint16(anon_sym_AMP_POUNDx),
	1496: uint16(5),
	1497: uint16(133),
	1498: uint16(1),
	1499: uint16(anon_sym_PERCENT),
	1500: uint16(301),
	1501: uint16(1),
	1502: uint16(sym_Name),
	1503: uint16(303),
	1504: uint16(1),
	1505: uint16(anon_sym_LPAREN),
	1506: uint16(103),
	1507: uint16(1),
	1508: uint16(sym__cp),
	1509: uint16(37),
	1510: uint16(2),
	1511: uint16(sym__choice),
	1512: uint16(sym_PEReference),
	1513: uint16(4),
	1514: uint16(348),
	1515: uint16(1),
	1516: uint16(sym__S),
	1517: uint16(73),
	1518: uint16(1),
	1519: uint16(aux_sym__choice_repeat1),
	1520: uint16(343),
	1521: uint16(2),
	1522: uint16(anon_sym_PIPE),
	1523: uint16(anon_sym_COMMA),
	1524: uint16(346),
	1525: uint16(2),
	1526: uint16(anon_sym_RPAREN),
	1527: uint16(anon_sym_PERCENT),
	1528: uint16(5),
	1529: uint16(133),
	1530: uint16(1),
	1531: uint16(anon_sym_PERCENT),
	1532: uint16(301),
	1533: uint16(1),
	1534: uint16(sym_Name),
	1535: uint16(303),
	1536: uint16(1),
	1537: uint16(anon_sym_LPAREN),
	1538: uint16(44),
	1539: uint16(1),
	1540: uint16(sym__cp),
	1541: uint16(37),
	1542: uint16(2),
	1543: uint16(sym__choice),
	1544: uint16(sym_PEReference),
	1545: uint16(2),
	1546: uint16(353),
	1547: uint16(2),
	1548: uint16(anon_sym_AMP),
	1549: uint16(anon_sym_AMP_POUND),
	1550: uint16(351),
	1551: uint16(4),
	1552: uint16(anon_sym_PERCENT),
	1553: uint16(anon_sym_DQUOTE),
	1554: uint16(aux_sym_EntityValue_token1),
	1555: uint16(anon_sym_AMP_POUNDx),
	1556: uint16(2),
	1557: uint16(357),
	1558: uint16(2),
	1559: uint16(anon_sym_AMP),
	1560: uint16(anon_sym_AMP_POUND),
	1561: uint16(355),
	1562: uint16(4),
	1563: uint16(anon_sym_PERCENT),
	1564: uint16(anon_sym_DQUOTE),
	1565: uint16(aux_sym_EntityValue_token1),
	1566: uint16(anon_sym_AMP_POUNDx),
	1567: uint16(2),
	1568: uint16(361),
	1569: uint16(1),
	1570: uint16(anon_sym_LT_BANG),
	1571: uint16(359),
	1572: uint16(5),
	1573: uint16(sym_Comment),
	1574: uint16(anon_sym_LT_QMARK),
	1575: uint16(anon_sym_LT_BANG_LBRACK),
	1576: uint16(anon_sym_PERCENT),
	1577: uint16(sym__S),
	1578: uint16(5),
	1579: uint16(133),
	1580: uint16(1),
	1581: uint16(anon_sym_PERCENT),
	1582: uint16(301),
	1583: uint16(1),
	1584: uint16(sym_Name),
	1585: uint16(303),
	1586: uint16(1),
	1587: uint16(anon_sym_LPAREN),
	1588: uint16(93),
	1589: uint16(1),
	1590: uint16(sym__cp),
	1591: uint16(37),
	1592: uint16(2),
	1593: uint16(sym__choice),
	1594: uint16(sym_PEReference),
	1595: uint16(5),
	1596: uint16(363),
	1597: uint16(1),
	1598: uint16(anon_sym_DQUOTE),
	1599: uint16(365),
	1600: uint16(1),
	1601: uint16(anon_sym_SQUOTE),
	1602: uint16(367),
	1603: uint16(1),
	1604: uint16(anon_sym_SYSTEM),
	1605: uint16(369),
	1606: uint16(1),
	1607: uint16(anon_sym_PUBLIC),
	1608: uint16(232),
	1609: uint16(2),
	1610: uint16(sym_EntityValue),
	1611: uint16(sym_ExternalID),
	1612: uint16(5),
	1613: uint16(133),
	1614: uint16(1),
	1615: uint16(anon_sym_PERCENT),
	1616: uint16(367),
	1617: uint16(1),
	1618: uint16(anon_sym_SYSTEM),
	1619: uint16(371),
	1620: uint16(1),
	1621: uint16(anon_sym_PUBLIC),
	1622: uint16(255),
	1623: uint16(1),
	1624: uint16(sym_PEReference),
	1625: uint16(213),
	1626: uint16(2),
	1627: uint16(sym_ExternalID),
	1628: uint16(sym_PublicID),
	1629: uint16(6),
	1630: uint16(363),
	1631: uint16(1),
	1632: uint16(anon_sym_DQUOTE),
	1633: uint16(365),
	1634: uint16(1),
	1635: uint16(anon_sym_SQUOTE),
	1636: uint16(367),
	1637: uint16(1),
	1638: uint16(anon_sym_SYSTEM),
	1639: uint16(369),
	1640: uint16(1),
	1641: uint16(anon_sym_PUBLIC),
	1642: uint16(147),
	1643: uint16(1),
	1644: uint16(sym_ExternalID),
	1645: uint16(239),
	1646: uint16(1),
	1647: uint16(sym_EntityValue),
	1648: uint16(6),
	1649: uint16(133),
	1650: uint16(1),
	1651: uint16(anon_sym_PERCENT),
	1652: uint16(373),
	1653: uint16(1),
	1654: uint16(sym_Name),
	1655: uint16(375),
	1656: uint16(1),
	1657: uint16(anon_sym_PIPE),
	1658: uint16(377),
	1659: uint16(1),
	1660: uint16(sym__S),
	1661: uint16(85),
	1662: uint16(1),
	1663: uint16(aux_sym_NotationType_repeat1),
	1664: uint16(224),
	1665: uint16(1),
	1666: uint16(sym_PEReference),
	1667: uint16(2),
	1668: uint16(381),
	1669: uint16(1),
	1670: uint16(anon_sym_LT_BANG),
	1671: uint16(379),
	1672: uint16(5),
	1673: uint16(sym_Comment),
	1674: uint16(anon_sym_LT_QMARK),
	1675: uint16(anon_sym_LT_BANG_LBRACK),
	1676: uint16(anon_sym_PERCENT),
	1677: uint16(sym__S),
	1678: uint16(6),
	1679: uint16(133),
	1680: uint16(1),
	1681: uint16(anon_sym_PERCENT),
	1682: uint16(375),
	1683: uint16(1),
	1684: uint16(anon_sym_PIPE),
	1685: uint16(377),
	1686: uint16(1),
	1687: uint16(sym__S),
	1688: uint16(383),
	1689: uint16(1),
	1690: uint16(sym_Name),
	1691: uint16(86),
	1692: uint16(1),
	1693: uint16(aux_sym_NotationType_repeat1),
	1694: uint16(179),
	1695: uint16(1),
	1696: uint16(sym_PEReference),
	1697: uint16(6),
	1698: uint16(133),
	1699: uint16(1),
	1700: uint16(anon_sym_PERCENT),
	1701: uint16(375),
	1702: uint16(1),
	1703: uint16(anon_sym_PIPE),
	1704: uint16(377),
	1705: uint16(1),
	1706: uint16(sym__S),
	1707: uint16(383),
	1708: uint16(1),
	1709: uint16(sym_Name),
	1710: uint16(105),
	1711: uint16(1),
	1712: uint16(aux_sym_NotationType_repeat1),
	1713: uint16(179),
	1714: uint16(1),
	1715: uint16(sym_PEReference),
	1716: uint16(6),
	1717: uint16(133),
	1718: uint16(1),
	1719: uint16(anon_sym_PERCENT),
	1720: uint16(375),
	1721: uint16(1),
	1722: uint16(anon_sym_PIPE),
	1723: uint16(377),
	1724: uint16(1),
	1725: uint16(sym__S),
	1726: uint16(385),
	1727: uint16(1),
	1728: uint16(sym_Name),
	1729: uint16(105),
	1730: uint16(1),
	1731: uint16(aux_sym_NotationType_repeat1),
	1732: uint16(236),
	1733: uint16(1),
	1734: uint16(sym_PEReference),
	1735: uint16(2),
	1736: uint16(389),
	1737: uint16(1),
	1738: uint16(anon_sym_LT_BANG),
	1739: uint16(387),
	1740: uint16(5),
	1741: uint16(sym_Comment),
	1742: uint16(anon_sym_LT_QMARK),
	1743: uint16(anon_sym_LT_BANG_LBRACK),
	1744: uint16(anon_sym_PERCENT),
	1745: uint16(sym__S),
	1746: uint16(2),
	1747: uint16(353),
	1748: uint16(2),
	1749: uint16(anon_sym_AMP),
	1750: uint16(anon_sym_AMP_POUND),
	1751: uint16(351),
	1752: uint16(4),
	1753: uint16(anon_sym_PERCENT),
	1754: uint16(anon_sym_SQUOTE),
	1755: uint16(aux_sym_EntityValue_token2),
	1756: uint16(anon_sym_AMP_POUNDx),
	1757: uint16(2),
	1758: uint16(357),
	1759: uint16(2),
	1760: uint16(anon_sym_AMP),
	1761: uint16(anon_sym_AMP_POUND),
	1762: uint16(355),
	1763: uint16(4),
	1764: uint16(anon_sym_PERCENT),
	1765: uint16(anon_sym_SQUOTE),
	1766: uint16(aux_sym_EntityValue_token2),
	1767: uint16(anon_sym_AMP_POUNDx),
	1768: uint16(2),
	1769: uint16(235),
	1770: uint16(2),
	1771: uint16(anon_sym_AMP),
	1772: uint16(anon_sym_AMP_POUND),
	1773: uint16(59),
	1774: uint16(4),
	1775: uint16(anon_sym_PERCENT),
	1776: uint16(anon_sym_DQUOTE),
	1777: uint16(aux_sym_EntityValue_token1),
	1778: uint16(anon_sym_AMP_POUNDx),
	1779: uint16(2),
	1780: uint16(341),
	1781: uint16(2),
	1782: uint16(anon_sym_AMP),
	1783: uint16(anon_sym_AMP_POUND),
	1784: uint16(339),
	1785: uint16(4),
	1786: uint16(anon_sym_PERCENT),
	1787: uint16(anon_sym_DQUOTE),
	1788: uint16(aux_sym_EntityValue_token1),
	1789: uint16(anon_sym_AMP_POUNDx),
	1790: uint16(4),
	1791: uint16(133),
	1792: uint16(1),
	1793: uint16(anon_sym_PERCENT),
	1794: uint16(269),
	1795: uint16(1),
	1796: uint16(anon_sym_RPAREN),
	1797: uint16(391),
	1798: uint16(1),
	1799: uint16(sym__S),
	1800: uint16(108),
	1801: uint16(2),
	1802: uint16(sym_PEReference),
	1803: uint16(aux_sym__choice_repeat2),
	1804: uint16(1),
	1805: uint16(393),
	1806: uint16(5),
	1807: uint16(anon_sym_PIPE),
	1808: uint16(anon_sym_RPAREN),
	1809: uint16(anon_sym_COMMA),
	1810: uint16(anon_sym_PERCENT),
	1811: uint16(sym__S),
	1812: uint16(5),
	1813: uint16(133),
	1814: uint16(1),
	1815: uint16(anon_sym_PERCENT),
	1816: uint16(395),
	1817: uint16(1),
	1818: uint16(anon_sym_PIPE),
	1819: uint16(397),
	1820: uint16(1),
	1821: uint16(anon_sym_RPAREN),
	1822: uint16(128),
	1823: uint16(1),
	1824: uint16(aux_sym_Mixed_repeat2),
	1825: uint16(165),
	1826: uint16(1),
	1827: uint16(sym_PEReference),
	1828: uint16(5),
	1829: uint16(133),
	1830: uint16(1),
	1831: uint16(anon_sym_PERCENT),
	1832: uint16(331),
	1833: uint16(1),
	1834: uint16(anon_sym_RPAREN),
	1835: uint16(395),
	1836: uint16(1),
	1837: uint16(anon_sym_PIPE),
	1838: uint16(132),
	1839: uint16(1),
	1840: uint16(aux_sym_Mixed_repeat2),
	1841: uint16(165),
	1842: uint16(1),
	1843: uint16(sym_PEReference),
	1844: uint16(4),
	1845: uint16(399),
	1846: uint16(1),
	1847: uint16(anon_sym_PIPE),
	1848: uint16(404),
	1849: uint16(1),
	1850: uint16(sym__S),
	1851: uint16(96),
	1852: uint16(1),
	1853: uint16(aux_sym_Mixed_repeat1),
	1854: uint16(402),
	1855: uint16(2),
	1856: uint16(anon_sym_RPAREN),
	1857: uint16(anon_sym_PERCENT),
	1858: uint16(5),
	1859: uint16(133),
	1860: uint16(1),
	1861: uint16(anon_sym_PERCENT),
	1862: uint16(395),
	1863: uint16(1),
	1864: uint16(anon_sym_PIPE),
	1865: uint16(407),
	1866: uint16(1),
	1867: uint16(anon_sym_RPAREN),
	1868: uint16(121),
	1869: uint16(1),
	1870: uint16(aux_sym_Mixed_repeat2),
	1871: uint16(165),
	1872: uint16(1),
	1873: uint16(sym_PEReference),
	1874: uint16(4),
	1875: uint16(133),
	1876: uint16(1),
	1877: uint16(anon_sym_PERCENT),
	1878: uint16(411),
	1879: uint16(1),
	1880: uint16(anon_sym_RPAREN),
	1881: uint16(170),
	1882: uint16(1),
	1883: uint16(sym_PEReference),
	1884: uint16(409),
	1885: uint16(2),
	1886: uint16(anon_sym_PIPE),
	1887: uint16(anon_sym_COMMA),
	1888: uint16(4),
	1889: uint16(133),
	1890: uint16(1),
	1891: uint16(anon_sym_PERCENT),
	1892: uint16(411),
	1893: uint16(1),
	1894: uint16(anon_sym_RPAREN),
	1895: uint16(413),
	1896: uint16(1),
	1897: uint16(sym__S),
	1898: uint16(108),
	1899: uint16(2),
	1900: uint16(sym_PEReference),
	1901: uint16(aux_sym__choice_repeat2),
	1902: uint16(2),
	1903: uint16(415),
	1904: uint16(2),
	1905: uint16(anon_sym_GT),
	1906: uint16(sym__S),
	1907: uint16(417),
	1908: uint16(3),
	1909: uint16(anon_sym_STAR),
	1910: uint16(anon_sym_QMARK),
	1911: uint16(anon_sym_PLUS),
	1912: uint16(4),
	1913: uint16(133),
	1914: uint16(1),
	1915: uint16(anon_sym_PERCENT),
	1916: uint16(241),
	1917: uint16(1),
	1918: uint16(anon_sym_RPAREN),
	1919: uint16(170),
	1920: uint16(1),
	1921: uint16(sym_PEReference),
	1922: uint16(409),
	1923: uint16(2),
	1924: uint16(anon_sym_PIPE),
	1925: uint16(anon_sym_COMMA),
	1926: uint16(4),
	1927: uint16(133),
	1928: uint16(1),
	1929: uint16(anon_sym_PERCENT),
	1930: uint16(421),
	1931: uint16(1),
	1932: uint16(sym__S),
	1933: uint16(182),
	1934: uint16(1),
	1935: uint16(sym_PEReference),
	1936: uint16(419),
	1937: uint16(2),
	1938: uint16(anon_sym_IGNORE),
	1939: uint16(anon_sym_INCLUDE),
	1940: uint16(1),
	1941: uint16(423),
	1942: uint16(5),
	1943: uint16(anon_sym_PIPE),
	1944: uint16(anon_sym_RPAREN),
	1945: uint16(anon_sym_COMMA),
	1946: uint16(anon_sym_PERCENT),
	1947: uint16(sym__S),
	1948: uint16(4),
	1949: uint16(133),
	1950: uint16(1),
	1951: uint16(anon_sym_PERCENT),
	1952: uint16(269),
	1953: uint16(1),
	1954: uint16(anon_sym_RPAREN),
	1955: uint16(170),
	1956: uint16(1),
	1957: uint16(sym_PEReference),
	1958: uint16(409),
	1959: uint16(2),
	1960: uint16(anon_sym_PIPE),
	1961: uint16(anon_sym_COMMA),
	1962: uint16(4),
	1963: uint16(427),
	1964: uint16(1),
	1965: uint16(anon_sym_PIPE),
	1966: uint16(430),
	1967: uint16(1),
	1968: uint16(sym__S),
	1969: uint16(105),
	1970: uint16(1),
	1971: uint16(aux_sym_NotationType_repeat1),
	1972: uint16(425),
	1973: uint16(2),
	1974: uint16(anon_sym_PERCENT),
	1975: uint16(sym_Name),
	1976: uint16(1),
	1977: uint16(433),
	1978: uint16(5),
	1979: uint16(anon_sym_PIPE),
	1980: uint16(anon_sym_RPAREN),
	1981: uint16(anon_sym_COMMA),
	1982: uint16(anon_sym_PERCENT),
	1983: uint16(sym__S),
	1984: uint16(1),
	1985: uint16(346),
	1986: uint16(5),
	1987: uint16(anon_sym_PIPE),
	1988: uint16(anon_sym_RPAREN),
	1989: uint16(anon_sym_COMMA),
	1990: uint16(anon_sym_PERCENT),
	1991: uint16(sym__S),
	1992: uint16(4),
	1993: uint16(435),
	1994: uint16(1),
	1995: uint16(anon_sym_RPAREN),
	1996: uint16(437),
	1997: uint16(1),
	1998: uint16(anon_sym_PERCENT),
	1999: uint16(440),
	2000: uint16(1),
	2001: uint16(sym__S),
	2002: uint16(108),
	2003: uint16(2),
	2004: uint16(sym_PEReference),
	2005: uint16(aux_sym__choice_repeat2),
	2006: uint16(5),
	2007: uint16(133),
	2008: uint16(1),
	2009: uint16(anon_sym_PERCENT),
	2010: uint16(321),
	2011: uint16(1),
	2012: uint16(anon_sym_RPAREN),
	2013: uint16(395),
	2014: uint16(1),
	2015: uint16(anon_sym_PIPE),
	2016: uint16(120),
	2017: uint16(1),
	2018: uint16(aux_sym_Mixed_repeat2),
	2019: uint16(165),
	2020: uint16(1),
	2021: uint16(sym_PEReference),
	2022: uint16(2),
	2023: uint16(341),
	2024: uint16(2),
	2025: uint16(anon_sym_AMP),
	2026: uint16(anon_sym_AMP_POUND),
	2027: uint16(339),
	2028: uint16(3),
	2029: uint16(anon_sym_DQUOTE),
	2030: uint16(anon_sym_AMP_POUNDx),
	2031: uint16(aux_sym_AttValue_token1),
	2032: uint16(2),
	2033: uint16(353),
	2034: uint16(2),
	2035: uint16(anon_sym_AMP),
	2036: uint16(anon_sym_AMP_POUND),
	2037: uint16(351),
	2038: uint16(3),
	2039: uint16(anon_sym_DQUOTE),
	2040: uint16(anon_sym_AMP_POUNDx),
	2041: uint16(aux_sym_AttValue_token1),
	2042: uint16(2),
	2043: uint16(357),
	2044: uint16(2),
	2045: uint16(anon_sym_AMP),
	2046: uint16(anon_sym_AMP_POUND),
	2047: uint16(355),
	2048: uint16(3),
	2049: uint16(anon_sym_DQUOTE),
	2050: uint16(anon_sym_AMP_POUNDx),
	2051: uint16(aux_sym_AttValue_token1),
	2052: uint16(4),
	2053: uint16(133),
	2054: uint16(1),
	2055: uint16(anon_sym_PERCENT),
	2056: uint16(241),
	2057: uint16(1),
	2058: uint16(anon_sym_RPAREN),
	2059: uint16(443),
	2060: uint16(1),
	2061: uint16(sym__S),
	2062: uint16(108),
	2063: uint16(2),
	2064: uint16(sym_PEReference),
	2065: uint16(aux_sym__choice_repeat2),
	2066: uint16(2),
	2067: uint16(341),
	2068: uint16(2),
	2069: uint16(anon_sym_AMP),
	2070: uint16(anon_sym_AMP_POUND),
	2071: uint16(339),
	2072: uint16(3),
	2073: uint16(anon_sym_SQUOTE),
	2074: uint16(anon_sym_AMP_POUNDx),
	2075: uint16(aux_sym_AttValue_token2),
	2076: uint16(2),
	2077: uint16(353),
	2078: uint16(2),
	2079: uint16(anon_sym_AMP),
	2080: uint16(anon_sym_AMP_POUND),
	2081: uint16(351),
	2082: uint16(3),
	2083: uint16(anon_sym_SQUOTE),
	2084: uint16(anon_sym_AMP_POUNDx),
	2085: uint16(aux_sym_AttValue_token2),
	2086: uint16(2),
	2087: uint16(357),
	2088: uint16(2),
	2089: uint16(anon_sym_AMP),
	2090: uint16(anon_sym_AMP_POUND),
	2091: uint16(355),
	2092: uint16(3),
	2093: uint16(anon_sym_SQUOTE),
	2094: uint16(anon_sym_AMP_POUNDx),
	2095: uint16(aux_sym_AttValue_token2),
	2096: uint16(2),
	2097: uint16(445),
	2098: uint16(1),
	2099: uint16(sym__S),
	2100: uint16(425),
	2101: uint16(3),
	2102: uint16(anon_sym_PIPE),
	2103: uint16(anon_sym_PERCENT),
	2104: uint16(sym_Name),
	2105: uint16(4),
	2106: uint16(133),
	2107: uint16(1),
	2108: uint16(anon_sym_PERCENT),
	2109: uint16(448),
	2110: uint16(1),
	2111: uint16(sym_Name),
	2112: uint16(450),
	2113: uint16(1),
	2114: uint16(anon_sym_GT),
	2115: uint16(210),
	2116: uint16(1),
	2117: uint16(sym_PEReference),
	2118: uint16(1),
	2119: uint16(452),
	2120: uint16(4),
	2121: uint16(anon_sym_PIPE),
	2122: uint16(anon_sym_RPAREN),
	2123: uint16(anon_sym_PERCENT),
	2124: uint16(sym__S),
	2125: uint16(4),
	2126: uint16(133),
	2127: uint16(1),
	2128: uint16(anon_sym_PERCENT),
	2129: uint16(397),
	2130: uint16(1),
	2131: uint16(anon_sym_RPAREN),
	2132: uint16(135),
	2133: uint16(1),
	2134: uint16(aux_sym_Mixed_repeat2),
	2135: uint16(165),
	2136: uint16(1),
	2137: uint16(sym_PEReference),
	2138: uint16(4),
	2139: uint16(133),
	2140: uint16(1),
	2141: uint16(anon_sym_PERCENT),
	2142: uint16(454),
	2143: uint16(1),
	2144: uint16(anon_sym_RPAREN),
	2145: uint16(135),
	2146: uint16(1),
	2147: uint16(aux_sym_Mixed_repeat2),
	2148: uint16(165),
	2149: uint16(1),
	2150: uint16(sym_PEReference),
	2151: uint16(4),
	2152: uint16(456),
	2153: uint16(1),
	2154: uint16(anon_sym_ELEMENT),
	2155: uint16(458),
	2156: uint16(1),
	2157: uint16(anon_sym_ATTLIST),
	2158: uint16(460),
	2159: uint16(1),
	2160: uint16(anon_sym_NOTATION),
	2161: uint16(462),
	2162: uint16(1),
	2163: uint16(anon_sym_ENTITY),
	2164: uint16(4),
	2165: uint16(464),
	2166: uint16(1),
	2167: uint16(anon_sym_PIPE),
	2168: uint16(466),
	2169: uint16(1),
	2170: uint16(anon_sym_RPAREN),
	2171: uint16(468),
	2172: uint16(1),
	2173: uint16(sym__S),
	2174: uint16(130),
	2175: uint16(1),
	2176: uint16(aux_sym_Enumeration_repeat1),
	2177: uint16(4),
	2178: uint16(464),
	2179: uint16(1),
	2180: uint16(anon_sym_PIPE),
	2181: uint16(466),
	2182: uint16(1),
	2183: uint16(anon_sym_RPAREN),
	2184: uint16(468),
	2185: uint16(1),
	2186: uint16(sym__S),
	2187: uint16(131),
	2188: uint16(1),
	2189: uint16(aux_sym_Enumeration_repeat1),
	2190: uint16(4),
	2191: uint16(133),
	2192: uint16(1),
	2193: uint16(anon_sym_PERCENT),
	2194: uint16(470),
	2195: uint16(1),
	2196: uint16(sym_Name),
	2197: uint16(472),
	2198: uint16(1),
	2199: uint16(sym__S),
	2200: uint16(82),
	2201: uint16(1),
	2202: uint16(sym_PEReference),
	2203: uint16(1),
	2204: uint16(402),
	2205: uint16(4),
	2206: uint16(anon_sym_PIPE),
	2207: uint16(anon_sym_RPAREN),
	2208: uint16(anon_sym_PERCENT),
	2209: uint16(sym__S),
	2210: uint16(1),
	2211: uint16(474),
	2212: uint16(4),
	2213: uint16(anon_sym_PIPE),
	2214: uint16(anon_sym_RPAREN),
	2215: uint16(anon_sym_PERCENT),
	2216: uint16(sym__S),
	2217: uint16(4),
	2218: uint16(133),
	2219: uint16(1),
	2220: uint16(anon_sym_PERCENT),
	2221: uint16(476),
	2222: uint16(1),
	2223: uint16(anon_sym_RPAREN),
	2224: uint16(135),
	2225: uint16(1),
	2226: uint16(aux_sym_Mixed_repeat2),
	2227: uint16(165),
	2228: uint16(1),
	2229: uint16(sym_PEReference),
	2230: uint16(4),
	2231: uint16(133),
	2232: uint16(1),
	2233: uint16(anon_sym_PERCENT),
	2234: uint16(478),
	2235: uint16(1),
	2236: uint16(sym_Name),
	2237: uint16(480),
	2238: uint16(1),
	2239: uint16(sym__S),
	2240: uint16(119),
	2241: uint16(1),
	2242: uint16(sym_PEReference),
	2243: uint16(4),
	2244: uint16(464),
	2245: uint16(1),
	2246: uint16(anon_sym_PIPE),
	2247: uint16(482),
	2248: uint16(1),
	2249: uint16(anon_sym_RPAREN),
	2250: uint16(484),
	2251: uint16(1),
	2252: uint16(sym__S),
	2253: uint16(131),
	2254: uint16(1),
	2255: uint16(aux_sym_Enumeration_repeat1),
	2256: uint16(4),
	2257: uint16(486),
	2258: uint16(1),
	2259: uint16(anon_sym_PIPE),
	2260: uint16(489),
	2261: uint16(1),
	2262: uint16(anon_sym_RPAREN),
	2263: uint16(491),
	2264: uint16(1),
	2265: uint16(sym__S),
	2266: uint16(131),
	2267: uint16(1),
	2268: uint16(aux_sym_Enumeration_repeat1),
	2269: uint16(4),
	2270: uint16(133),
	2271: uint16(1),
	2272: uint16(anon_sym_PERCENT),
	2273: uint16(407),
	2274: uint16(1),
	2275: uint16(anon_sym_RPAREN),
	2276: uint16(135),
	2277: uint16(1),
	2278: uint16(aux_sym_Mixed_repeat2),
	2279: uint16(165),
	2280: uint16(1),
	2281: uint16(sym_PEReference),
	2282: uint16(4),
	2283: uint16(133),
	2284: uint16(1),
	2285: uint16(anon_sym_PERCENT),
	2286: uint16(448),
	2287: uint16(1),
	2288: uint16(sym_Name),
	2289: uint16(494),
	2290: uint16(1),
	2291: uint16(anon_sym_GT),
	2292: uint16(210),
	2293: uint16(1),
	2294: uint16(sym_PEReference),
	2295: uint16(4),
	2296: uint16(133),
	2297: uint16(1),
	2298: uint16(anon_sym_PERCENT),
	2299: uint16(321),
	2300: uint16(1),
	2301: uint16(anon_sym_RPAREN),
	2302: uint16(135),
	2303: uint16(1),
	2304: uint16(aux_sym_Mixed_repeat2),
	2305: uint16(165),
	2306: uint16(1),
	2307: uint16(sym_PEReference),
	2308: uint16(4),
	2309: uint16(496),
	2310: uint16(1),
	2311: uint16(anon_sym_RPAREN),
	2312: uint16(498),
	2313: uint16(1),
	2314: uint16(anon_sym_PERCENT),
	2315: uint16(135),
	2316: uint16(1),
	2317: uint16(aux_sym_Mixed_repeat2),
	2318: uint16(165),
	2319: uint16(1),
	2320: uint16(sym_PEReference),
	2321: uint16(2),
	2322: uint16(503),
	2323: uint16(1),
	2324: uint16(sym__S),
	2325: uint16(501),
	2326: uint16(3),
	2327: uint16(anon_sym_PIPE),
	2328: uint16(anon_sym_PERCENT),
	2329: uint16(sym_Name),
	2330: uint16(3),
	2331: uint16(494),
	2332: uint16(1),
	2333: uint16(anon_sym_GT),
	2334: uint16(506),
	2335: uint16(1),
	2336: uint16(sym__S),
	2337: uint16(145),
	2338: uint16(2),
	2339: uint16(sym_AttDef),
	2340: uint16(aux_sym_AttlistDecl_repeat1),
	2341: uint16(3),
	2342: uint16(133),
	2343: uint16(1),
	2344: uint16(anon_sym_PERCENT),
	2345: uint16(183),
	2346: uint16(1),
	2347: uint16(sym_PEReference),
	2348: uint16(508),
	2349: uint16(2),
	2350: uint16(anon_sym_IGNORE),
	2351: uint16(anon_sym_INCLUDE),
	2352: uint16(4),
	2353: uint16(133),
	2354: uint16(1),
	2355: uint16(anon_sym_PERCENT),
	2356: uint16(510),
	2357: uint16(1),
	2358: uint16(sym_Name),
	2359: uint16(512),
	2360: uint16(1),
	2361: uint16(sym__S),
	2362: uint16(126),
	2363: uint16(1),
	2364: uint16(sym_PEReference),
	2365: uint16(1),
	2366: uint16(425),
	2367: uint16(4),
	2368: uint16(anon_sym_PIPE),
	2369: uint16(anon_sym_PERCENT),
	2370: uint16(sym__S),
	2371: uint16(sym_Name),
	2372: uint16(1),
	2373: uint16(514),
	2374: uint16(4),
	2375: uint16(anon_sym_PIPE),
	2376: uint16(anon_sym_PERCENT),
	2377: uint16(sym__S),
	2378: uint16(sym_Name),
	2379: uint16(3),
	2380: uint16(516),
	2381: uint16(1),
	2382: uint16(anon_sym_GT),
	2383: uint16(518),
	2384: uint16(1),
	2385: uint16(sym__S),
	2386: uint16(137),
	2387: uint16(2),
	2388: uint16(sym_AttDef),
	2389: uint16(aux_sym_AttlistDecl_repeat1),
	2390: uint16(4),
	2391: uint16(133),
	2392: uint16(1),
	2393: uint16(anon_sym_PERCENT),
	2394: uint16(331),
	2395: uint16(1),
	2396: uint16(anon_sym_RPAREN),
	2397: uint16(135),
	2398: uint16(1),
	2399: uint16(aux_sym_Mixed_repeat2),
	2400: uint16(165),
	2401: uint16(1),
	2402: uint16(sym_PEReference),
	2403: uint16(4),
	2404: uint16(464),
	2405: uint16(1),
	2406: uint16(anon_sym_PIPE),
	2407: uint16(520),
	2408: uint16(1),
	2409: uint16(anon_sym_RPAREN),
	2410: uint16(522),
	2411: uint16(1),
	2412: uint16(sym__S),
	2413: uint16(124),
	2414: uint16(1),
	2415: uint16(aux_sym_Enumeration_repeat1),
	2416: uint16(3),
	2417: uint16(524),
	2418: uint16(1),
	2419: uint16(anon_sym_GT),
	2420: uint16(526),
	2421: uint16(1),
	2422: uint16(sym__S),
	2423: uint16(145),
	2424: uint16(2),
	2425: uint16(sym_AttDef),
	2426: uint16(aux_sym_AttlistDecl_repeat1),
	2427: uint16(2),
	2428: uint16(531),
	2429: uint16(1),
	2430: uint16(anon_sym_STAR),
	2431: uint16(529),
	2432: uint16(2),
	2433: uint16(anon_sym_GT),
	2434: uint16(sym__S),
	2435: uint16(3),
	2436: uint16(533),
	2437: uint16(1),
	2438: uint16(anon_sym_GT),
	2439: uint16(535),
	2440: uint16(1),
	2441: uint16(sym__S),
	2442: uint16(192),
	2443: uint16(1),
	2444: uint16(sym_NDataDecl),
	2445: uint16(3),
	2446: uint16(133),
	2447: uint16(1),
	2448: uint16(anon_sym_PERCENT),
	2449: uint16(537),
	2450: uint16(1),
	2451: uint16(sym_Name),
	2452: uint16(84),
	2453: uint16(1),
	2454: uint16(sym_PEReference),
	2455: uint16(3),
	2456: uint16(133),
	2457: uint16(1),
	2458: uint16(anon_sym_PERCENT),
	2459: uint16(448),
	2460: uint16(1),
	2461: uint16(sym_Name),
	2462: uint16(210),
	2463: uint16(1),
	2464: uint16(sym_PEReference),
	2465: uint16(3),
	2466: uint16(161),
	2467: uint16(1),
	2468: uint16(anon_sym_DQUOTE),
	2469: uint16(163),
	2470: uint16(1),
	2471: uint16(anon_sym_SQUOTE),
	2472: uint16(225),
	2473: uint16(1),
	2474: uint16(sym_AttValue),
	2475: uint16(3),
	2476: uint16(133),
	2477: uint16(1),
	2478: uint16(anon_sym_PERCENT),
	2479: uint16(539),
	2480: uint16(1),
	2481: uint16(sym_Name),
	2482: uint16(127),
	2483: uint16(1),
	2484: uint16(sym_PEReference),
	2485: uint16(3),
	2486: uint16(541),
	2487: uint16(1),
	2488: uint16(anon_sym_DQUOTE),
	2489: uint16(543),
	2490: uint16(1),
	2491: uint16(anon_sym_SQUOTE),
	2492: uint16(300),
	2493: uint16(1),
	2494: uint16(sym_PubidLiteral),
	2495: uint16(3),
	2496: uint16(545),
	2497: uint16(1),
	2498: uint16(anon_sym_DQUOTE),
	2499: uint16(547),
	2500: uint16(1),
	2501: uint16(anon_sym_SQUOTE),
	2502: uint16(223),
	2503: uint16(1),
	2504: uint16(sym_SystemLiteral),
	2505: uint16(1),
	2506: uint16(549),
	2507: uint16(3),
	2508: uint16(anon_sym_PIPE),
	2509: uint16(anon_sym_RPAREN),
	2510: uint16(sym__S),
	2511: uint16(3),
	2512: uint16(551),
	2513: uint16(1),
	2514: uint16(sym_Name),
	2515: uint16(553),
	2516: uint16(1),
	2517: uint16(anon_sym_PERCENT),
	2518: uint16(333),
	2519: uint16(1),
	2520: uint16(sym_PEReference),
	2521: uint16(3),
	2522: uint16(133),
	2523: uint16(1),
	2524: uint16(anon_sym_PERCENT),
	2525: uint16(478),
	2526: uint16(1),
	2527: uint16(sym_Name),
	2528: uint16(119),
	2529: uint16(1),
	2530: uint16(sym_PEReference),
	2531: uint16(3),
	2532: uint16(541),
	2533: uint16(1),
	2534: uint16(anon_sym_DQUOTE),
	2535: uint16(543),
	2536: uint16(1),
	2537: uint16(anon_sym_SQUOTE),
	2538: uint16(227),
	2539: uint16(1),
	2540: uint16(sym_PubidLiteral),
	2541: uint16(3),
	2542: uint16(555),
	2543: uint16(1),
	2544: uint16(anon_sym_QMARK_GT),
	2545: uint16(557),
	2546: uint16(1),
	2547: uint16(sym__S),
	2548: uint16(194),
	2549: uint16(1),
	2550: uint16(sym__EncodingDecl),
	2551: uint16(3),
	2552: uint16(541),
	2553: uint16(1),
	2554: uint16(anon_sym_DQUOTE),
	2555: uint16(543),
	2556: uint16(1),
	2557: uint16(anon_sym_SQUOTE),
	2558: uint16(228),
	2559: uint16(1),
	2560: uint16(sym_PubidLiteral),
	2561: uint16(1),
	2562: uint16(559),
	2563: uint16(3),
	2564: uint16(anon_sym_PIPE),
	2565: uint16(anon_sym_RPAREN),
	2566: uint16(sym__S),
	2567: uint16(2),
	2568: uint16(563),
	2569: uint16(1),
	2570: uint16(anon_sym_STAR),
	2571: uint16(561),
	2572: uint16(2),
	2573: uint16(anon_sym_GT),
	2574: uint16(sym__S),
	2575: uint16(3),
	2576: uint16(133),
	2577: uint16(1),
	2578: uint16(anon_sym_PERCENT),
	2579: uint16(565),
	2580: uint16(1),
	2581: uint16(sym_Name),
	2582: uint16(208),
	2583: uint16(1),
	2584: uint16(sym_PEReference),
	2585: uint16(2),
	2586: uint16(569),
	2587: uint16(1),
	2588: uint16(anon_sym_STAR),
	2589: uint16(567),
	2590: uint16(2),
	2591: uint16(anon_sym_GT),
	2592: uint16(sym__S),
	2593: uint16(3),
	2594: uint16(571),
	2595: uint16(1),
	2596: uint16(sym__S),
	2597: uint16(573),
	2598: uint16(1),
	2599: uint16(anon_sym_EQ),
	2600: uint16(190),
	2601: uint16(1),
	2602: uint16(sym__Eq),
	2603: uint16(2),
	2604: uint16(577),
	2605: uint16(1),
	2606: uint16(sym__S),
	2607: uint16(575),
	2608: uint16(2),
	2609: uint16(anon_sym_RPAREN),
	2610: uint16(anon_sym_PERCENT),
	2611: uint16(3),
	2612: uint16(133),
	2613: uint16(1),
	2614: uint16(anon_sym_PERCENT),
	2615: uint16(411),
	2616: uint16(1),
	2617: uint16(anon_sym_RPAREN),
	2618: uint16(170),
	2619: uint16(1),
	2620: uint16(sym_PEReference),
	2621: uint16(3),
	2622: uint16(133),
	2623: uint16(1),
	2624: uint16(anon_sym_PERCENT),
	2625: uint16(579),
	2626: uint16(1),
	2627: uint16(anon_sym_RPAREN),
	2628: uint16(170),
	2629: uint16(1),
	2630: uint16(sym_PEReference),
	2631: uint16(2),
	2632: uint16(583),
	2633: uint16(1),
	2634: uint16(sym__S),
	2635: uint16(581),
	2636: uint16(2),
	2637: uint16(anon_sym_DQUOTE),
	2638: uint16(anon_sym_SQUOTE),
	2639: uint16(2),
	2640: uint16(587),
	2641: uint16(1),
	2642: uint16(anon_sym_STAR),
	2643: uint16(585),
	2644: uint16(2),
	2645: uint16(anon_sym_GT),
	2646: uint16(sym__S),
	2647: uint16(1),
	2648: uint16(435),
	2649: uint16(3),
	2650: uint16(anon_sym_RPAREN),
	2651: uint16(anon_sym_PERCENT),
	2652: uint16(sym__S),
	2653: uint16(2),
	2654: uint16(591),
	2655: uint16(1),
	2656: uint16(sym__S),
	2657: uint16(589),
	2658: uint16(2),
	2659: uint16(anon_sym_DQUOTE),
	2660: uint16(anon_sym_SQUOTE),
	2661: uint16(3),
	2662: uint16(133),
	2663: uint16(1),
	2664: uint16(anon_sym_PERCENT),
	2665: uint16(269),
	2666: uint16(1),
	2667: uint16(anon_sym_RPAREN),
	2668: uint16(170),
	2669: uint16(1),
	2670: uint16(sym_PEReference),
	2671: uint16(3),
	2672: uint16(133),
	2673: uint16(1),
	2674: uint16(anon_sym_PERCENT),
	2675: uint16(593),
	2676: uint16(1),
	2677: uint16(sym_Name),
	2678: uint16(142),
	2679: uint16(1),
	2680: uint16(sym_PEReference),
	2681: uint16(3),
	2682: uint16(133),
	2683: uint16(1),
	2684: uint16(anon_sym_PERCENT),
	2685: uint16(595),
	2686: uint16(1),
	2687: uint16(sym_Name),
	2688: uint16(307),
	2689: uint16(1),
	2690: uint16(sym_PEReference),
	2691: uint16(3),
	2692: uint16(571),
	2693: uint16(1),
	2694: uint16(sym__S),
	2695: uint16(573),
	2696: uint16(1),
	2697: uint16(anon_sym_EQ),
	2698: uint16(181),
	2699: uint16(1),
	2700: uint16(sym__Eq),
	2701: uint16(3),
	2702: uint16(545),
	2703: uint16(1),
	2704: uint16(anon_sym_DQUOTE),
	2705: uint16(547),
	2706: uint16(1),
	2707: uint16(anon_sym_SQUOTE),
	2708: uint16(207),
	2709: uint16(1),
	2710: uint16(sym_SystemLiteral),
	2711: uint16(1),
	2712: uint16(489),
	2713: uint16(3),
	2714: uint16(anon_sym_PIPE),
	2715: uint16(anon_sym_RPAREN),
	2716: uint16(sym__S),
	2717: uint16(3),
	2718: uint16(133),
	2719: uint16(1),
	2720: uint16(anon_sym_PERCENT),
	2721: uint16(597),
	2722: uint16(1),
	2723: uint16(sym_Name),
	2724: uint16(250),
	2725: uint16(1),
	2726: uint16(sym_PEReference),
	2727: uint16(2),
	2728: uint16(599),
	2729: uint16(1),
	2730: uint16(anon_sym_RPAREN),
	2731: uint16(601),
	2732: uint16(1),
	2733: uint16(sym__S),
	2734: uint16(2),
	2735: uint16(603),
	2736: uint16(1),
	2737: uint16(sym__S),
	2738: uint16(158),
	2739: uint16(1),
	2740: uint16(sym__VersionInfo),
	2741: uint16(2),
	2742: uint16(605),
	2743: uint16(1),
	2744: uint16(anon_sym_DQUOTE),
	2745: uint16(607),
	2746: uint16(1),
	2747: uint16(anon_sym_SQUOTE),
	2748: uint16(2),
	2749: uint16(609),
	2750: uint16(1),
	2751: uint16(anon_sym_LBRACK),
	2752: uint16(611),
	2753: uint16(1),
	2754: uint16(sym__S),
	2755: uint16(2),
	2756: uint16(613),
	2757: uint16(1),
	2758: uint16(anon_sym_LBRACK),
	2759: uint16(615),
	2760: uint16(1),
	2761: uint16(sym__S),
	2762: uint16(2),
	2763: uint16(617),
	2764: uint16(1),
	2765: uint16(anon_sym_xml),
	2766: uint16(619),
	2767: uint16(1),
	2768: uint16(sym_PITarget),
	2769: uint16(1),
	2770: uint16(621),
	2771: uint16(2),
	2772: uint16(anon_sym_GT),
	2773: uint16(sym__S),
	2774: uint16(1),
	2775: uint16(623),
	2776: uint16(2),
	2777: uint16(anon_sym_GT),
	2778: uint16(sym__S),
	2779: uint16(2),
	2780: uint16(625),
	2781: uint16(1),
	2782: uint16(anon_sym_QMARK_GT),
	2783: uint16(627),
	2784: uint16(1),
	2785: uint16(anon_sym_encoding),
	2786: uint16(1),
	2787: uint16(589),
	2788: uint16(2),
	2789: uint16(anon_sym_DQUOTE),
	2790: uint16(anon_sym_SQUOTE),
	2791: uint16(1),
	2792: uint16(567),
	2793: uint16(2),
	2794: uint16(anon_sym_GT),
	2795: uint16(sym__S),
	2796: uint16(2),
	2797: uint16(629),
	2798: uint16(1),
	2799: uint16(anon_sym_DQUOTE),
	2800: uint16(631),
	2801: uint16(1),
	2802: uint16(anon_sym_SQUOTE),
	2803: uint16(2),
	2804: uint16(633),
	2805: uint16(1),
	2806: uint16(anon_sym_GT),
	2807: uint16(635),
	2808: uint16(1),
	2809: uint16(anon_sym_NDATA),
	2810: uint16(2),
	2811: uint16(633),
	2812: uint16(1),
	2813: uint16(anon_sym_GT),
	2814: uint16(637),
	2815: uint16(1),
	2816: uint16(sym__S),
	2817: uint16(1),
	2818: uint16(639),
	2819: uint16(2),
	2820: uint16(anon_sym_QMARK_GT),
	2821: uint16(sym__S),
	2822: uint16(2),
	2823: uint16(625),
	2824: uint16(1),
	2825: uint16(anon_sym_QMARK_GT),
	2826: uint16(641),
	2827: uint16(1),
	2828: uint16(sym__S),
	2829: uint16(1),
	2830: uint16(409),
	2831: uint16(2),
	2832: uint16(anon_sym_PIPE),
	2833: uint16(anon_sym_COMMA),
	2834: uint16(2),
	2835: uint16(133),
	2836: uint16(1),
	2837: uint16(anon_sym_PERCENT),
	2838: uint16(170),
	2839: uint16(1),
	2840: uint16(sym_PEReference),
	2841: uint16(2),
	2842: uint16(643),
	2843: uint16(1),
	2844: uint16(sym__S),
	2845: uint16(645),
	2846: uint16(1),
	2847: uint16(sym_Nmtoken),
	2848: uint16(1),
	2849: uint16(647),
	2850: uint16(2),
	2851: uint16(anon_sym_GT),
	2852: uint16(sym__S),
	2853: uint16(2),
	2854: uint16(466),
	2855: uint16(1),
	2856: uint16(anon_sym_RPAREN),
	2857: uint16(649),
	2858: uint16(1),
	2859: uint16(anon_sym_PIPE),
	2860: uint16(1),
	2861: uint16(651),
	2862: uint16(2),
	2863: uint16(anon_sym_GT),
	2864: uint16(sym__S),
	2865: uint16(2),
	2866: uint16(653),
	2867: uint16(1),
	2868: uint16(anon_sym_QMARK_GT),
	2869: uint16(655),
	2870: uint16(1),
	2871: uint16(sym__S),
	2872: uint16(1),
	2873: uint16(657),
	2874: uint16(2),
	2875: uint16(anon_sym_GT),
	2876: uint16(sym__S),
	2877: uint16(2),
	2878: uint16(659),
	2879: uint16(1),
	2880: uint16(anon_sym_GT),
	2881: uint16(661),
	2882: uint16(1),
	2883: uint16(sym__S),
	2884: uint16(1),
	2885: uint16(663),
	2886: uint16(2),
	2887: uint16(anon_sym_GT),
	2888: uint16(sym__S),
	2889: uint16(1),
	2890: uint16(665),
	2891: uint16(2),
	2892: uint16(anon_sym_GT),
	2893: uint16(sym__S),
	2894: uint16(1),
	2895: uint16(667),
	2896: uint16(2),
	2897: uint16(anon_sym_GT),
	2898: uint16(sym__S),
	2899: uint16(1),
	2900: uint16(669),
	2901: uint16(2),
	2902: uint16(anon_sym_GT),
	2903: uint16(sym__S),
	2904: uint16(1),
	2905: uint16(671),
	2906: uint16(2),
	2907: uint16(anon_sym_GT),
	2908: uint16(sym__S),
	2909: uint16(1),
	2910: uint16(561),
	2911: uint16(2),
	2912: uint16(anon_sym_GT),
	2913: uint16(sym__S),
	2914: uint16(2),
	2915: uint16(524),
	2916: uint16(1),
	2917: uint16(anon_sym_GT),
	2918: uint16(673),
	2919: uint16(1),
	2920: uint16(sym__S),
	2921: uint16(1),
	2922: uint16(676),
	2923: uint16(2),
	2924: uint16(anon_sym_GT),
	2925: uint16(sym__S),
	2926: uint16(2),
	2927: uint16(482),
	2928: uint16(1),
	2929: uint16(anon_sym_RPAREN),
	2930: uint16(649),
	2931: uint16(1),
	2932: uint16(anon_sym_PIPE),
	2933: uint16(2),
	2934: uint16(678),
	2935: uint16(1),
	2936: uint16(anon_sym_GT),
	2937: uint16(680),
	2938: uint16(1),
	2939: uint16(sym__S),
	2940: uint16(2),
	2941: uint16(682),
	2942: uint16(1),
	2943: uint16(sym__S),
	2944: uint16(684),
	2945: uint16(1),
	2946: uint16(sym_Nmtoken),
	2947: uint16(1),
	2948: uint16(496),
	2949: uint16(2),
	2950: uint16(anon_sym_RPAREN),
	2951: uint16(anon_sym_PERCENT),
	2952: uint16(1),
	2953: uint16(686),
	2954: uint16(2),
	2955: uint16(anon_sym_GT),
	2956: uint16(sym__S),
	2957: uint16(1),
	2958: uint16(688),
	2959: uint16(2),
	2960: uint16(anon_sym_GT),
	2961: uint16(sym__S),
	2962: uint16(2),
	2963: uint16(690),
	2964: uint16(1),
	2965: uint16(anon_sym_GT),
	2966: uint16(692),
	2967: uint16(1),
	2968: uint16(sym__S),
	2969: uint16(1),
	2970: uint16(694),
	2971: uint16(2),
	2972: uint16(anon_sym_GT),
	2973: uint16(sym__S),
	2974: uint16(1),
	2975: uint16(696),
	2976: uint16(2),
	2977: uint16(anon_sym_GT),
	2978: uint16(sym__S),
	2979: uint16(1),
	2980: uint16(698),
	2981: uint16(2),
	2982: uint16(anon_sym_GT),
	2983: uint16(sym__S),
	2984: uint16(2),
	2985: uint16(649),
	2986: uint16(1),
	2987: uint16(anon_sym_PIPE),
	2988: uint16(700),
	2989: uint16(1),
	2990: uint16(anon_sym_RPAREN),
	2991: uint16(1),
	2992: uint16(702),
	2993: uint16(2),
	2994: uint16(anon_sym_GT),
	2995: uint16(sym__S),
	2996: uint16(2),
	2997: uint16(704),
	2998: uint16(1),
	2999: uint16(anon_sym_RPAREN),
	3000: uint16(706),
	3001: uint16(1),
	3002: uint16(sym__S),
	3003: uint16(1),
	3004: uint16(708),
	3005: uint16(2),
	3006: uint16(anon_sym_GT),
	3007: uint16(sym__S),
	3008: uint16(1),
	3009: uint16(710),
	3010: uint16(2),
	3011: uint16(anon_sym_GT),
	3012: uint16(sym__S),
	3013: uint16(2),
	3014: uint16(712),
	3015: uint16(1),
	3016: uint16(anon_sym_GT),
	3017: uint16(714),
	3018: uint16(1),
	3019: uint16(sym__S),
	3020: uint16(1),
	3021: uint16(712),
	3022: uint16(2),
	3023: uint16(anon_sym_GT),
	3024: uint16(sym__S),
	3025: uint16(1),
	3026: uint16(716),
	3027: uint16(2),
	3028: uint16(anon_sym_GT),
	3029: uint16(sym__S),
	3030: uint16(1),
	3031: uint16(718),
	3032: uint16(2),
	3033: uint16(anon_sym_GT),
	3034: uint16(sym__S),
	3035: uint16(2),
	3036: uint16(720),
	3037: uint16(1),
	3038: uint16(sym_Name),
	3039: uint16(722),
	3040: uint16(1),
	3041: uint16(sym__S),
	3042: uint16(2),
	3043: uint16(724),
	3044: uint16(1),
	3045: uint16(anon_sym_GT),
	3046: uint16(726),
	3047: uint16(1),
	3048: uint16(sym__S),
	3049: uint16(1),
	3050: uint16(728),
	3051: uint16(2),
	3052: uint16(anon_sym_GT),
	3053: uint16(sym__S),
	3054: uint16(1),
	3055: uint16(730),
	3056: uint16(2),
	3057: uint16(anon_sym_GT),
	3058: uint16(sym__S),
	3059: uint16(1),
	3060: uint16(732),
	3061: uint16(2),
	3062: uint16(anon_sym_GT),
	3063: uint16(sym__S),
	3064: uint16(2),
	3065: uint16(734),
	3066: uint16(1),
	3067: uint16(anon_sym_RPAREN),
	3068: uint16(736),
	3069: uint16(1),
	3070: uint16(sym__S),
	3071: uint16(1),
	3072: uint16(738),
	3073: uint16(2),
	3074: uint16(anon_sym_GT),
	3075: uint16(sym__S),
	3076: uint16(1),
	3077: uint16(740),
	3078: uint16(2),
	3079: uint16(anon_sym_GT),
	3080: uint16(sym__S),
	3081: uint16(2),
	3082: uint16(533),
	3083: uint16(1),
	3084: uint16(anon_sym_GT),
	3085: uint16(742),
	3086: uint16(1),
	3087: uint16(sym__S),
	3088: uint16(1),
	3089: uint16(744),
	3090: uint16(2),
	3091: uint16(anon_sym_QMARK_GT),
	3092: uint16(sym__S),
	3093: uint16(1),
	3094: uint16(746),
	3095: uint16(2),
	3096: uint16(anon_sym_DQUOTE),
	3097: uint16(anon_sym_SQUOTE),
	3098: uint16(1),
	3099: uint16(585),
	3100: uint16(2),
	3101: uint16(anon_sym_GT),
	3102: uint16(sym__S),
	3103: uint16(2),
	3104: uint16(748),
	3105: uint16(1),
	3106: uint16(sym__S),
	3107: uint16(750),
	3108: uint16(1),
	3109: uint16(sym_Nmtoken),
	3110: uint16(1),
	3111: uint16(752),
	3112: uint16(1),
	3113: uint16(anon_sym_SEMI),
	3114: uint16(1),
	3115: uint16(754),
	3116: uint16(1),
	3117: uint16(sym__S),
	3118: uint16(1),
	3119: uint16(756),
	3120: uint16(1),
	3121: uint16(aux_sym_CharRef_token2),
	3122: uint16(1),
	3123: uint16(563),
	3124: uint16(1),
	3125: uint16(anon_sym_STAR),
	3126: uint16(1),
	3127: uint16(613),
	3128: uint16(1),
	3129: uint16(anon_sym_LBRACK),
	3130: uint16(1),
	3131: uint16(758),
	3132: uint16(1),
	3133: uint16(sym__pi_content),
	3134: uint16(1),
	3135: uint16(760),
	3136: uint16(1),
	3137: uint16(sym__S),
	3138: uint16(1),
	3139: uint16(762),
	3140: uint16(1),
	3141: uint16(sym__S),
	3142: uint16(1),
	3143: uint16(764),
	3144: uint16(1),
	3145: uint16(sym__S),
	3146: uint16(1),
	3147: uint16(766),
	3148: uint16(1),
	3149: uint16(anon_sym_GT),
	3150: uint16(1),
	3151: uint16(768),
	3152: uint16(1),
	3153: uint16(sym_VersionNum),
	3154: uint16(1),
	3155: uint16(770),
	3156: uint16(1),
	3157: uint16(sym__S),
	3158: uint16(1),
	3159: uint16(684),
	3160: uint16(1),
	3161: uint16(sym_Nmtoken),
	3162: uint16(1),
	3163: uint16(772),
	3164: uint16(1),
	3165: uint16(anon_sym_QMARK_GT),
	3166: uint16(1),
	3167: uint16(774),
	3168: uint16(1),
	3169: uint16(sym_VersionNum),
	3170: uint16(1),
	3171: uint16(776),
	3172: uint16(1),
	3173: uint16(sym__S),
	3174: uint16(1),
	3175: uint16(569),
	3176: uint16(1),
	3177: uint16(anon_sym_STAR),
	3178: uint16(1),
	3179: uint16(778),
	3180: uint16(1),
	3181: uint16(sym__S),
	3182: uint16(1),
	3183: uint16(780),
	3184: uint16(1),
	3185: uint16(sym__S),
	3186: uint16(1),
	3187: uint16(395),
	3188: uint16(1),
	3189: uint16(anon_sym_PIPE),
	3190: uint16(1),
	3191: uint16(633),
	3192: uint16(1),
	3193: uint16(anon_sym_GT),
	3194: uint16(1),
	3195: uint16(782),
	3196: uint16(1),
	3197: uint16(sym__S),
	3198: uint16(1),
	3199: uint16(784),
	3200: uint16(1),
	3201: uint16(sym__S),
	3202: uint16(1),
	3203: uint16(786),
	3204: uint16(1),
	3205: uint16(anon_sym_STAR),
	3206: uint16(1),
	3207: uint16(788),
	3208: uint16(1),
	3209: uint16(aux_sym_SystemLiteral_token1),
	3210: uint16(1),
	3211: uint16(790),
	3212: uint16(1),
	3213: uint16(sym__S),
	3214: uint16(1),
	3215: uint16(792),
	3216: uint16(1),
	3217: uint16(anon_sym_DQUOTE),
	3218: uint16(1),
	3219: uint16(794),
	3220: uint16(1),
	3221: uint16(aux_sym_SystemLiteral_token2),
	3222: uint16(1),
	3223: uint16(796),
	3224: uint16(1),
	3225: uint16(sym_Nmtoken),
	3226: uint16(1),
	3227: uint16(649),
	3228: uint16(1),
	3229: uint16(anon_sym_PIPE),
	3230: uint16(1),
	3231: uint16(792),
	3232: uint16(1),
	3233: uint16(anon_sym_SQUOTE),
	3234: uint16(1),
	3235: uint16(798),
	3236: uint16(1),
	3237: uint16(aux_sym_PubidLiteral_token1),
	3238: uint16(1),
	3239: uint16(800),
	3240: uint16(1),
	3241: uint16(anon_sym_PIPE),
	3242: uint16(1),
	3243: uint16(802),
	3244: uint16(1),
	3245: uint16(anon_sym_QMARK_GT),
	3246: uint16(1),
	3247: uint16(804),
	3248: uint16(1),
	3249: uint16(aux_sym_PubidLiteral_token2),
	3250: uint16(1),
	3251: uint16(806),
	3252: uint16(1),
	3253: uint16(sym_Name),
	3254: uint16(1),
	3255: uint16(808),
	3256: uint16(1),
	3257: uint16(anon_sym_SEMI),
	3258: uint16(1),
	3259: uint16(810),
	3260: uint16(1),
	3261: uint16(sym_Nmtoken),
	3262: uint16(1),
	3263: uint16(812),
	3264: uint16(1),
	3266: uint16(1),
	3267: uint16(814),
	3268: uint16(1),
	3269: uint16(sym_Name),
	3270: uint16(1),
	3271: uint16(816),
	3272: uint16(1),
	3273: uint16(anon_sym_LPAREN),
	3274: uint16(1),
	3275: uint16(818),
	3276: uint16(1),
	3277: uint16(anon_sym_GT),
	3278: uint16(1),
	3279: uint16(820),
	3280: uint16(1),
	3281: uint16(anon_sym_DQUOTE),
	3282: uint16(1),
	3283: uint16(822),
	3284: uint16(1),
	3285: uint16(anon_sym_SEMI),
	3286: uint16(1),
	3287: uint16(824),
	3288: uint16(1),
	3289: uint16(anon_sym_SEMI),
	3290: uint16(1),
	3291: uint16(820),
	3292: uint16(1),
	3293: uint16(anon_sym_SQUOTE),
	3294: uint16(1),
	3295: uint16(826),
	3296: uint16(1),
	3297: uint16(anon_sym_DQUOTE),
	3298: uint16(1),
	3299: uint16(599),
	3300: uint16(1),
	3301: uint16(anon_sym_RPAREN),
	3302: uint16(1),
	3303: uint16(828),
	3304: uint16(1),
	3305: uint16(sym_EncName),
	3306: uint16(1),
	3307: uint16(830),
	3308: uint16(1),
	3309: uint16(sym__S),
	3310: uint16(1),
	3311: uint16(734),
	3312: uint16(1),
	3313: uint16(anon_sym_RPAREN),
	3314: uint16(1),
	3315: uint16(826),
	3316: uint16(1),
	3317: uint16(anon_sym_SQUOTE),
	3318: uint16(1),
	3319: uint16(832),
	3320: uint16(1),
	3321: uint16(sym__S),
	3322: uint16(1),
	3323: uint16(834),
	3324: uint16(1),
	3325: uint16(anon_sym_LBRACK),
	3326: uint16(1),
	3327: uint16(836),
	3328: uint16(1),
	3329: uint16(anon_sym_RPAREN),
	3330: uint16(1),
	3331: uint16(619),
	3332: uint16(1),
	3333: uint16(sym_PITarget),
	3334: uint16(1),
	3335: uint16(714),
	3336: uint16(1),
	3337: uint16(sym__S),
	3338: uint16(1),
	3339: uint16(838),
	3340: uint16(1),
	3341: uint16(anon_sym_EQ),
	3342: uint16(1),
	3343: uint16(840),
	3344: uint16(1),
	3345: uint16(sym__S),
	3346: uint16(1),
	3347: uint16(842),
	3348: uint16(1),
	3349: uint16(anon_sym_GT),
	3350: uint16(1),
	3351: uint16(844),
	3352: uint16(1),
	3353: uint16(anon_sym_GT),
	3354: uint16(1),
	3355: uint16(846),
	3356: uint16(1),
	3357: uint16(anon_sym_version),
	3358: uint16(1),
	3359: uint16(848),
	3360: uint16(1),
	3361: uint16(sym_Name),
	3362: uint16(1),
	3363: uint16(850),
	3364: uint16(1),
	3365: uint16(sym__S),
	3366: uint16(1),
	3367: uint16(852),
	3368: uint16(1),
	3369: uint16(anon_sym_DQUOTE),
	3370: uint16(1),
	3371: uint16(852),
	3372: uint16(1),
	3373: uint16(anon_sym_SQUOTE),
	3374: uint16(1),
	3375: uint16(854),
	3376: uint16(1),
	3377: uint16(sym_EncName),
	3378: uint16(1),
	3379: uint16(756),
	3380: uint16(1),
	3381: uint16(aux_sym_CharRef_token1),
	3382: uint16(1),
	3383: uint16(856),
	3384: uint16(1),
	3385: uint16(anon_sym_SEMI),
	3386: uint16(1),
	3387: uint16(858),
	3388: uint16(1),
	3389: uint16(anon_sym_SEMI),
	3390: uint16(1),
	3391: uint16(860),
	3392: uint16(1),
	3393: uint16(anon_sym_SEMI),
	3394: uint16(1),
	3395: uint16(862),
	3396: uint16(1),
	3397: uint16(anon_sym_SEMI),
	3398: uint16(1),
	3399: uint16(864),
	3400: uint16(1),
	3401: uint16(anon_sym_SEMI),
	3402: uint16(1),
	3403: uint16(866),
	3404: uint16(1),
	3405: uint16(anon_sym_SEMI),
	3406: uint16(1),
	3407: uint16(868),
	3408: uint16(1),
	3409: uint16(anon_sym_SEMI),
	3410: uint16(1),
	3411: uint16(870),
	3412: uint16(1),
	3413: uint16(anon_sym_SEMI),
	3414: uint16(1),
	3415: uint16(587),
	3416: uint16(1),
	3417: uint16(anon_sym_STAR),
	3418: uint16(1),
	3419: uint16(720),
	3420: uint16(1),
	3421: uint16(sym_Name),
	3422: uint16(1),
	3423: uint16(872),
	3424: uint16(1),
	3425: uint16(sym_Name),
	3426: uint16(1),
	3427: uint16(874),
	3428: uint16(1),
	3429: uint16(aux_sym_CharRef_token1),
	3430: uint16(1),
	3431: uint16(874),
	3432: uint16(1),
	3433: uint16(aux_sym_CharRef_token2),
	3434: uint16(1),
	3435: uint16(876),
	3436: uint16(1),
	3437: uint16(sym_Name),
	3438: uint16(1),
	3439: uint16(878),
	3440: uint16(1),
	3441: uint16(sym_Name),
	3442: uint16(1),
	3443: uint16(880),
	3444: uint16(1),
	3445: uint16(aux_sym_CharRef_token1),
	3446: uint16(1),
	3447: uint16(880),
	3448: uint16(1),
	3449: uint16(aux_sym_CharRef_token2),
	3450: uint16(1),
	3451: uint16(882),
	3452: uint16(1),
	3453: uint16(sym_Name),
	3454: uint16(1),
	3455: uint16(884),
	3456: uint16(1),
	3457: uint16(sym_Name),
	3458: uint16(1),
	3459: uint16(886),
	3460: uint16(1),
	3461: uint16(aux_sym_CharRef_token1),
	3462: uint16(1),
	3463: uint16(886),
	3464: uint16(1),
	3465: uint16(aux_sym_CharRef_token2),
	3466: uint16(1),
	3467: uint16(888),
	3468: uint16(1),
	3469: uint16(sym__S),
}

var ts_small_parse_table_map = [332]uint32_t{
	1:   uint32(42),
	2:   uint32(83),
	3:   uint32(124),
	4:   uint32(165),
	5:   uint32(206),
	6:   uint32(247),
	7:   uint32(288),
	8:   uint32(329),
	9:   uint32(370),
	10:  uint32(408),
	11:  uint32(422),
	12:  uint32(450),
	13:  uint32(478),
	14:  uint32(500),
	15:  uint32(528),
	16:  uint32(556),
	17:  uint32(584),
	18:  uint32(606),
	19:  uint32(634),
	20:  uint32(659),
	21:  uint32(671),
	22:  uint32(695),
	23:  uint32(717),
	24:  uint32(729),
	25:  uint32(753),
	26:  uint32(765),
	27:  uint32(789),
	28:  uint32(813),
	29:  uint32(825),
	30:  uint32(849),
	31:  uint32(873),
	32:  uint32(885),
	33:  uint32(909),
	34:  uint32(922),
	35:  uint32(935),
	36:  uint32(948),
	37:  uint32(969),
	38:  uint32(982),
	39:  uint32(995),
	40:  uint32(1008),
	41:  uint32(1021),
	42:  uint32(1034),
	43:  uint32(1055),
	44:  uint32(1068),
	45:  uint32(1081),
	46:  uint32(1094),
	47:  uint32(1107),
	48:  uint32(1120),
	49:  uint32(1133),
	50:  uint32(1154),
	51:  uint32(1167),
	52:  uint32(1180),
	53:  uint32(1193),
	54:  uint32(1206),
	55:  uint32(1219),
	56:  uint32(1232),
	57:  uint32(1245),
	58:  uint32(1270),
	59:  uint32(1283),
	60:  uint32(1304),
	61:  uint32(1324),
	62:  uint32(1346),
	63:  uint32(1368),
	64:  uint32(1390),
	65:  uint32(1410),
	66:  uint32(1432),
	67:  uint32(1454),
	68:  uint32(1474),
	69:  uint32(1485),
	70:  uint32(1496),
	71:  uint32(1513),
	72:  uint32(1528),
	73:  uint32(1545),
	74:  uint32(1556),
	75:  uint32(1567),
	76:  uint32(1578),
	77:  uint32(1595),
	78:  uint32(1612),
	79:  uint32(1629),
	80:  uint32(1648),
	81:  uint32(1667),
	82:  uint32(1678),
	83:  uint32(1697),
	84:  uint32(1716),
	85:  uint32(1735),
	86:  uint32(1746),
	87:  uint32(1757),
	88:  uint32(1768),
	89:  uint32(1779),
	90:  uint32(1790),
	91:  uint32(1804),
	92:  uint32(1812),
	93:  uint32(1828),
	94:  uint32(1844),
	95:  uint32(1858),
	96:  uint32(1874),
	97:  uint32(1888),
	98:  uint32(1902),
	99:  uint32(1912),
	100: uint32(1926),
	101: uint32(1940),
	102: uint32(1948),
	103: uint32(1962),
	104: uint32(1976),
	105: uint32(1984),
	106: uint32(1992),
	107: uint32(2006),
	108: uint32(2022),
	109: uint32(2032),
	110: uint32(2042),
	111: uint32(2052),
	112: uint32(2066),
	113: uint32(2076),
	114: uint32(2086),
	115: uint32(2096),
	116: uint32(2105),
	117: uint32(2118),
	118: uint32(2125),
	119: uint32(2138),
	120: uint32(2151),
	121: uint32(2164),
	122: uint32(2177),
	123: uint32(2190),
	124: uint32(2203),
	125: uint32(2210),
	126: uint32(2217),
	127: uint32(2230),
	128: uint32(2243),
	129: uint32(2256),
	130: uint32(2269),
	131: uint32(2282),
	132: uint32(2295),
	133: uint32(2308),
	134: uint32(2321),
	135: uint32(2330),
	136: uint32(2341),
	137: uint32(2352),
	138: uint32(2365),
	139: uint32(2372),
	140: uint32(2379),
	141: uint32(2390),
	142: uint32(2403),
	143: uint32(2416),
	144: uint32(2427),
	145: uint32(2435),
	146: uint32(2445),
	147: uint32(2455),
	148: uint32(2465),
	149: uint32(2475),
	150: uint32(2485),
	151: uint32(2495),
	152: uint32(2505),
	153: uint32(2511),
	154: uint32(2521),
	155: uint32(2531),
	156: uint32(2541),
	157: uint32(2551),
	158: uint32(2561),
	159: uint32(2567),
	160: uint32(2575),
	161: uint32(2585),
	162: uint32(2593),
	163: uint32(2603),
	164: uint32(2611),
	165: uint32(2621),
	166: uint32(2631),
	167: uint32(2639),
	168: uint32(2647),
	169: uint32(2653),
	170: uint32(2661),
	171: uint32(2671),
	172: uint32(2681),
	173: uint32(2691),
	174: uint32(2701),
	175: uint32(2711),
	176: uint32(2717),
	177: uint32(2727),
	178: uint32(2734),
	179: uint32(2741),
	180: uint32(2748),
	181: uint32(2755),
	182: uint32(2762),
	183: uint32(2769),
	184: uint32(2774),
	185: uint32(2779),
	186: uint32(2786),
	187: uint32(2791),
	188: uint32(2796),
	189: uint32(2803),
	190: uint32(2810),
	191: uint32(2817),
	192: uint32(2822),
	193: uint32(2829),
	194: uint32(2834),
	195: uint32(2841),
	196: uint32(2848),
	197: uint32(2853),
	198: uint32(2860),
	199: uint32(2865),
	200: uint32(2872),
	201: uint32(2877),
	202: uint32(2884),
	203: uint32(2889),
	204: uint32(2894),
	205: uint32(2899),
	206: uint32(2904),
	207: uint32(2909),
	208: uint32(2914),
	209: uint32(2921),
	210: uint32(2926),
	211: uint32(2933),
	212: uint32(2940),
	213: uint32(2947),
	214: uint32(2952),
	215: uint32(2957),
	216: uint32(2962),
	217: uint32(2969),
	218: uint32(2974),
	219: uint32(2979),
	220: uint32(2984),
	221: uint32(2991),
	222: uint32(2996),
	223: uint32(3003),
	224: uint32(3008),
	225: uint32(3013),
	226: uint32(3020),
	227: uint32(3025),
	228: uint32(3030),
	229: uint32(3035),
	230: uint32(3042),
	231: uint32(3049),
	232: uint32(3054),
	233: uint32(3059),
	234: uint32(3064),
	235: uint32(3071),
	236: uint32(3076),
	237: uint32(3081),
	238: uint32(3088),
	239: uint32(3093),
	240: uint32(3098),
	241: uint32(3103),
	242: uint32(3110),
	243: uint32(3114),
	244: uint32(3118),
	245: uint32(3122),
	246: uint32(3126),
	247: uint32(3130),
	248: uint32(3134),
	249: uint32(3138),
	250: uint32(3142),
	251: uint32(3146),
	252: uint32(3150),
	253: uint32(3154),
	254: uint32(3158),
	255: uint32(3162),
	256: uint32(3166),
	257: uint32(3170),
	258: uint32(3174),
	259: uint32(3178),
	260: uint32(3182),
	261: uint32(3186),
	262: uint32(3190),
	263: uint32(3194),
	264: uint32(3198),
	265: uint32(3202),
	266: uint32(3206),
	267: uint32(3210),
	268: uint32(3214),
	269: uint32(3218),
	270: uint32(3222),
	271: uint32(3226),
	272: uint32(3230),
	273: uint32(3234),
	274: uint32(3238),
	275: uint32(3242),
	276: uint32(3246),
	277: uint32(3250),
	278: uint32(3254),
	279: uint32(3258),
	280: uint32(3262),
	281: uint32(3266),
	282: uint32(3270),
	283: uint32(3274),
	284: uint32(3278),
	285: uint32(3282),
	286: uint32(3286),
	287: uint32(3290),
	288: uint32(3294),
	289: uint32(3298),
	290: uint32(3302),
	291: uint32(3306),
	292: uint32(3310),
	293: uint32(3314),
	294: uint32(3318),
	295: uint32(3322),
	296: uint32(3326),
	297: uint32(3330),
	298: uint32(3334),
	299: uint32(3338),
	300: uint32(3342),
	301: uint32(3346),
	302: uint32(3350),
	303: uint32(3354),
	304: uint32(3358),
	305: uint32(3362),
	306: uint32(3366),
	307: uint32(3370),
	308: uint32(3374),
	309: uint32(3378),
	310: uint32(3382),
	311: uint32(3386),
	312: uint32(3390),
	313: uint32(3394),
	314: uint32(3398),
	315: uint32(3402),
	316: uint32(3406),
	317: uint32(3410),
	318: uint32(3414),
	319: uint32(3418),
	320: uint32(3422),
	321: uint32(3426),
	322: uint32(3430),
	323: uint32(3434),
	324: uint32(3438),
	325: uint32(3442),
	326: uint32(3446),
	327: uint32(3450),
	328: uint32(3454),
	329: uint32(3458),
	330: uint32(3462),
	331: uint32(3466),
}

var ts_parse_actions = [890]TSParseActionEntry{
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
		Fstate: uint16(libc.Int32FromInt32(184)),
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
		Fstate: uint16(libc.Int32FromInt32(102)),
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
		Fstate: uint16(libc.Int32FromInt32(122)),
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
		Fstate: uint16(libc.Int32FromInt32(279)),
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
		Fstate: uint16(libc.Int32FromInt32(6)),
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
		Fstate: uint16(libc.Int32FromInt32(36)),
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
		Fsymbol:      uint16(aux_sym_extSubset_repeat1),
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
		Fsymbol:      uint16(aux_sym_extSubset_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(299)),
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
		Fsymbol:      uint16(aux_sym_extSubset_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(102)),
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
		Fsymbol:      uint16(aux_sym_extSubset_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(122)),
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
		Fsymbol:      uint16(aux_sym_extSubset_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(279)),
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
		Fsymbol:      uint16(aux_sym_extSubset_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(2)),
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
		Fsymbol:      uint16(aux_sym_extSubset_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(36)),
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
		Fstate: uint16(libc.Int32FromInt32(299)),
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
		Fstate: uint16(libc.Int32FromInt32(55)),
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
		Fstate: uint16(libc.Int32FromInt32(5)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_extSubset),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(2)),
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
		Fstate: uint16(libc.Int32FromInt32(58)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_extSubset),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(49)),
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
		Fstate: uint16(libc.Int32FromInt32(8)),
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
		Fstate: uint16(libc.Int32FromInt32(10)),
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
		Fstate: uint16(libc.Int32FromInt32(46)),
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
		Fstate: uint16(libc.Int32FromInt32(4)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_PEReference),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(325)),
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
		Fstate: uint16(libc.Int32FromInt32(186)),
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
		Fstate: uint16(libc.Int32FromInt32(17)),
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
		Fstate: uint16(libc.Int32FromInt32(306)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(311)),
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
		Fstate: uint16(libc.Int32FromInt32(246)),
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
		Fstate: uint16(libc.Int32FromInt32(329)),
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
		Fstate: uint16(libc.Int32FromInt32(18)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(322)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(323)),
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
		Fstate: uint16(libc.Int32FromInt32(324)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__cp),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_EntityValue_repeat2),
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
		Fstate:      uint16(libc.Int32FromInt32(329)),
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
		Fcount:    uint8(1),
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
		Fsymbol:      uint16(aux_sym_EntityValue_repeat2),
	})))),
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
		Fsymbol:      uint16(aux_sym_EntityValue_repeat2),
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
		Fstate:      uint16(libc.Int32FromInt32(16)),
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
		Fsymbol:      uint16(aux_sym_EntityValue_repeat2),
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
		Fstate:      uint16(libc.Int32FromInt32(322)),
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
		Fsymbol:      uint16(aux_sym_EntityValue_repeat2),
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
		Fstate:      uint16(libc.Int32FromInt32(323)),
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
		Fsymbol:      uint16(aux_sym_EntityValue_repeat2),
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
		Fstate:      uint16(libc.Int32FromInt32(324)),
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
		Fstate: uint16(libc.Int32FromInt32(233)),
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
		Fstate: uint16(libc.Int32FromInt32(20)),
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
		Fstate: uint16(libc.Int32FromInt32(16)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_EntityValue_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(325)),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_EntityValue_repeat1),
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
		Fsymbol:      uint16(aux_sym_EntityValue_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(20)),
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
		Fsymbol:      uint16(aux_sym_EntityValue_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(306)),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_EntityValue_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(311)),
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
		Fsymbol:      uint16(aux_sym_EntityValue_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(246)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(243)),
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
		Fstate: uint16(libc.Int32FromInt32(216)),
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
		Fstate: uint16(libc.Int32FromInt32(262)),
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
		Fstate: uint16(libc.Int32FromInt32(321)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym__choice),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_AttValue_repeat2),
	})))),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_AttValue_repeat2),
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
		Fstate:      uint16(libc.Int32FromInt32(330)),
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
		Fsymbol:      uint16(aux_sym_AttValue_repeat2),
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
		Fstate:      uint16(libc.Int32FromInt32(331)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_AttValue_repeat2),
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
		Fstate:      uint16(libc.Int32FromInt32(332)),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_AttValue_repeat2),
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
		Fstate:      uint16(libc.Int32FromInt32(23)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(200)),
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
		Fstate: uint16(libc.Int32FromInt32(59)),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__choice),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(202)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(293)),
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
		Fstate: uint16(libc.Int32FromInt32(28)),
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
		Fstate: uint16(libc.Int32FromInt32(29)),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym__choice),
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
		Fstate: uint16(libc.Int32FromInt32(217)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(326)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(327)),
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
		Fstate: uint16(libc.Int32FromInt32(328)),
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
		Fstate: uint16(libc.Int32FromInt32(31)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(330)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(331)),
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
		Fstate: uint16(libc.Int32FromInt32(332)),
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
		Fstate: uint16(libc.Int32FromInt32(32)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym__choice),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(226)),
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
		Fstate: uint16(libc.Int32FromInt32(34)),
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
		Fstate: uint16(libc.Int32FromInt32(23)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym__choice),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_AttValue_repeat1),
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
		Fsymbol:      uint16(aux_sym_AttValue_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(326)),
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
		Fsymbol:      uint16(aux_sym_AttValue_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(327)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_AttValue_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(328)),
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
		Fsymbol:      uint16(aux_sym_AttValue_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(34)),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_PI),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_PI),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__markupdecl),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__markupdecl),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(25)),
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
		Fstate: uint16(libc.Int32FromInt32(101)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_elementdecl),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_elementdecl),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_AttlistDecl),
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
		Fcount: uint8(1),
	}})),
	230: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_AttlistDecl),
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
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_NotationDecl),
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
		Fcount: uint8(1),
	}})),
	234: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_NotationDecl),
	})))),
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
		Fcount: uint8(1),
	}})),
	236: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_PEReference),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_GEDecl),
	})))),
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
		Fcount: uint8(1),
	}})),
	240: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_GEDecl),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(104)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__EntityDecl),
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
		Fcount: uint8(1),
	}})),
	248: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__EntityDecl),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_conditionalSect),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_conditionalSect),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(8),
		Fsymbol:      uint16(sym_elementdecl),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(8),
		Fsymbol:      uint16(sym_elementdecl),
	})))),
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
		Fchild_count: uint8(8),
		Fsymbol:      uint16(sym_NotationDecl),
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
		Fcount: uint8(1),
	}})),
	260: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(8),
		Fsymbol:      uint16(sym_NotationDecl),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_conditionalSect),
	})))),
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
		Fcount: uint8(1),
	}})),
	264: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_conditionalSect),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(8),
		Fsymbol:      uint16(sym_GEDecl),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(8),
		Fsymbol:      uint16(sym_GEDecl),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(27)),
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
		Fstate: uint16(libc.Int32FromInt32(98)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	274: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_PI),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_PI),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(9),
		Fsymbol:      uint16(sym_PEDecl),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(9),
		Fsymbol:      uint16(sym_PEDecl),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	282: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(9),
		Fsymbol:      uint16(sym_GEDecl),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(9),
		Fsymbol:      uint16(sym_GEDecl),
	})))),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_conditionalSect),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_conditionalSect),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_AttlistDecl),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_AttlistDecl),
	})))),
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
		Fchild_count: uint8(10),
		Fsymbol:      uint16(sym_PEDecl),
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
		Fchild_count: uint8(10),
		Fsymbol:      uint16(sym_PEDecl),
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
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_conditionalSect),
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
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_conditionalSect),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(62)),
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
		Fstate: uint16(libc.Int32FromInt32(63)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(68)),
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
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_AttlistDecl),
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
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_AttlistDecl),
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
		Fstate: uint16(libc.Int32FromInt32(74)),
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
		Fstate: uint16(libc.Int32FromInt32(139)),
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
		Fstate: uint16(libc.Int32FromInt32(146)),
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
		Fstate: uint16(libc.Int32FromInt32(95)),
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
		Fstate: uint16(libc.Int32FromInt32(260)),
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
		Fstate: uint16(libc.Int32FromInt32(94)),
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
		Fstate: uint16(libc.Int32FromInt32(320)),
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
		Fstate: uint16(libc.Int32FromInt32(109)),
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
		Fstate: uint16(libc.Int32FromInt32(78)),
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
		Fstate: uint16(libc.Int32FromInt32(169)),
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
		Fstate: uint16(libc.Int32FromInt32(97)),
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
		Fstate: uint16(libc.Int32FromInt32(67)),
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
		Fstate: uint16(libc.Int32FromInt32(72)),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__Reference),
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
		Fcount: uint8(1),
	}})),
	342: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__Reference),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym__choice_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(69)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	347: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__choice_repeat1),
	})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym__choice_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(195)),
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
		Fsymbol:      uint16(sym_EntityRef),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_EntityRef),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_CharRef),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_CharRef),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_TextDecl),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_TextDecl),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(251)),
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
		Fstate: uint16(libc.Int32FromInt32(296)),
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
		Fstate: uint16(libc.Int32FromInt32(252)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(224)),
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
		Fstate: uint16(libc.Int32FromInt32(136)),
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
		Fstate: uint16(libc.Int32FromInt32(276)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_TextDecl),
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
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_TextDecl),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(179)),
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
		Fstate: uint16(libc.Int32FromInt32(236)),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_TextDecl),
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
		Fcount: uint8(1),
	}})),
	390: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_TextDecl),
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
		Fstate: uint16(libc.Int32FromInt32(166)),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(aux_sym__choice_repeat1),
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
		Fstate: uint16(libc.Int32FromInt32(129)),
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
		Fstate: uint16(libc.Int32FromInt32(247)),
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
		Fsymbol:      uint16(aux_sym_Mixed_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(139)),
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
		Fsymbol:      uint16(aux_sym_Mixed_repeat1),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_Mixed_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(263)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(163)),
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
		Fstate: uint16(libc.Int32FromInt32(66)),
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
		Fstate: uint16(libc.Int32FromInt32(30)),
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
		Fstate: uint16(libc.Int32FromInt32(167)),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_children),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(182)),
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
		Fstate: uint16(libc.Int32FromInt32(138)),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(aux_sym__choice_repeat1),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_NotationType_repeat1),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_NotationType_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(136)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_NotationType_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(276)),
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
		Fsymbol:      uint16(sym__cp),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__choice_repeat2),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym__choice_repeat2),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(321)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	441: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__choice_repeat2),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(196)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	444: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_NotationType_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(141)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(245)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(aux_sym_Mixed_repeat1),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(261)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(265)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(266)),
	}})))),
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
		Fstate: uint16(libc.Int32FromInt32(269)),
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
		Fstate: uint16(libc.Int32FromInt32(197)),
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
		Fstate: uint16(libc.Int32FromInt32(211)),
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
		Fstate: uint16(libc.Int32FromInt32(212)),
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
		Fstate: uint16(libc.Int32FromInt32(82)),
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
		Fstate: uint16(libc.Int32FromInt32(148)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(aux_sym_Mixed_repeat1),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	477: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(267)),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(151)),
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
		Fstate: uint16(libc.Int32FromInt32(221)),
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
		Fstate: uint16(libc.Int32FromInt32(222)),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_Enumeration_repeat1),
	})))),
	488: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(libc.Int32FromInt32(197)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fsymbol:      uint16(aux_sym_Enumeration_repeat1),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_Enumeration_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(273)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(60)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_Mixed_repeat2),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_Mixed_repeat2),
	})))),
	500: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(libc.Int32FromInt32(321)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_NotationType_repeat1),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_NotationType_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(140)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(118)),
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
		Fstate: uint16(libc.Int32FromInt32(183)),
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
		Fstate: uint16(libc.Int32FromInt32(126)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fsymbol:      uint16(aux_sym_NotationType_repeat1),
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
		Fstate: uint16(libc.Int32FromInt32(56)),
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
		Fstate: uint16(libc.Int32FromInt32(133)),
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
		Fstate: uint16(libc.Int32FromInt32(198)),
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
		Fstate: uint16(libc.Int32FromInt32(199)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_AttlistDecl_repeat1),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_AttlistDecl_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(149)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_Mixed),
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
		Fstate: uint16(libc.Int32FromInt32(242)),
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
		Fstate: uint16(libc.Int32FromInt32(43)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(84)),
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
		Fstate: uint16(libc.Int32FromInt32(127)),
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
		Fstate: uint16(libc.Int32FromInt32(275)),
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
		Fstate: uint16(libc.Int32FromInt32(278)),
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
	546: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(268)),
	}})))),
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
		Fstate: uint16(libc.Int32FromInt32(271)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(aux_sym_Enumeration_repeat1),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(333)),
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
		Fstate: uint16(libc.Int32FromInt32(231)),
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
		Fstate: uint16(libc.Int32FromInt32(77)),
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
		Fstate: uint16(libc.Int32FromInt32(187)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(aux_sym_Enumeration_repeat1),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_Mixed),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(220)),
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
		Fstate: uint16(libc.Int32FromInt32(208)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_Mixed),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(209)),
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
		Fstate: uint16(libc.Int32FromInt32(301)),
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
		Fstate: uint16(libc.Int32FromInt32(168)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_Mixed_repeat2),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(215)),
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
		Fstate: uint16(libc.Int32FromInt32(33)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__Eq),
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
		Fstate: uint16(libc.Int32FromInt32(188)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_Mixed),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(189)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__Eq),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(241)),
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
		Fstate: uint16(libc.Int32FromInt32(142)),
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
		Fstate: uint16(libc.Int32FromInt32(307)),
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
		Fstate: uint16(libc.Int32FromInt32(250)),
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
		Fstate: uint16(libc.Int32FromInt32(235)),
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
		Fstate: uint16(libc.Int32FromInt32(294)),
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
		Fstate: uint16(libc.Int32FromInt32(305)),
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
		Fstate: uint16(libc.Int32FromInt32(254)),
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
		Fstate: uint16(libc.Int32FromInt32(258)),
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
		Fstate: uint16(libc.Int32FromInt32(7)),
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
		Fstate: uint16(libc.Int32FromInt32(248)),
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
		Fstate: uint16(libc.Int32FromInt32(3)),
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
		Fstate: uint16(libc.Int32FromInt32(297)),
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
		Fstate: uint16(libc.Int32FromInt32(180)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(201)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	622: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_children),
	})))),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_EntityValue),
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
		Fstate: uint16(libc.Int32FromInt32(87)),
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
		Fstate: uint16(libc.Int32FromInt32(164)),
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
		Fstate: uint16(libc.Int32FromInt32(310)),
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
		Fstate: uint16(libc.Int32FromInt32(292)),
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
		Fstate: uint16(libc.Int32FromInt32(50)),
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
		Fstate: uint16(libc.Int32FromInt32(302)),
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
		Fstate: uint16(libc.Int32FromInt32(303)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym__VersionInfo),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(257)),
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
		Fstate: uint16(libc.Int32FromInt32(256)),
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
		Fstate: uint16(libc.Int32FromInt32(177)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_Enumeration),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(214)),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_contentspec),
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
		Fstate: uint16(libc.Int32FromInt32(52)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	656: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(249)),
	}})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_DefaultDecl),
	})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(253)),
	}})))),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_AttDef),
	})))),
	665: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	666: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_SystemLiteral),
	})))),
	667: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	668: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_PubidLiteral),
	})))),
	669: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	670: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_ExternalID),
	})))),
	671: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	672: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_NDataDecl),
	})))),
	673: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	674: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_AttlistDecl_repeat1),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(21)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_Enumeration),
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
		Fstate: uint16(libc.Int32FromInt32(41)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(285)),
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
		Fstate: uint16(libc.Int32FromInt32(272)),
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
		Fstate: uint16(libc.Int32FromInt32(154)),
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
		Fsymbol:      uint16(sym__AttType),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_AttValue),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	691: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_AttDef),
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
		Fstate: uint16(libc.Int32FromInt32(26)),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__EnumeratedType),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	697: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_Mixed),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	699: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_Enumeration),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(230)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	703: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_ExternalID),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(234)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(291)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	709: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_DefaultDecl),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	711: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_AttValue),
		Fproduction_id: uint16(1),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	713: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_PublicID),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(176)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	717: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(8),
		Fsymbol:      uint16(sym_Mixed),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	719: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_Enumeration),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(312)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(283)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(53)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(304)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	729: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_EntityValue),
		Fproduction_id: uint16(1),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	731: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_NotationType),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	733: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_NotationType),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(237)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	737: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(298)),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	739: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(8),
		Fsymbol:      uint16(sym_NotationType),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	741: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(9),
		Fsymbol:      uint16(sym_NotationType),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(264)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym__EncodingDecl),
	})))),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__Eq),
	})))),
	748: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	749: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(281)),
	}})))),
	750: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	751: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(144)),
	}})))),
	752: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(116)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	755: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(288)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	759: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(277)),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	761: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	762: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	763: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	764: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	765: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	766: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	767: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	768: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	769: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(308)),
	}})))),
	770: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	771: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	772: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	773: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	774: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	775: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(309)),
	}})))),
	776: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	777: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	778: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	779: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	780: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	781: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(284)),
	}})))),
	782: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	783: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	784: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	785: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	786: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	787: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(229)),
	}})))),
	788: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	789: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(286)),
	}})))),
	790: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	791: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	792: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	793: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(240)),
	}})))),
	794: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	795: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(289)),
	}})))),
	796: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	797: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	798: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	799: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(290)),
	}})))),
	800: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(117)),
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
		Fstate: uint16(libc.Int32FromInt32(35)),
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
		Fstate: uint16(libc.Int32FromInt32(295)),
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
		Fstate: uint16(libc.Int32FromInt32(280)),
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
		Fstate: uint16(libc.Int32FromInt32(42)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	811: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	812: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	813: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f uint8_t
		_ [7]byte
	}{f: uint8(TSParseActionTypeAccept)})),
	814: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(259)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	817: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	818: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	819: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	820: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(205)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	825: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	826: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	827: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(206)),
	}})))),
	828: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(274)),
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
		Fstate: uint16(libc.Int32FromInt32(150)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(152)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(9)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	837: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(238)),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(171)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(162)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	843: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	845: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	847: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	849: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(287)),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(24)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(193)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	855: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(270)),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	857: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	859: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	861: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	863: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	865: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(112)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	869: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(115)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	873: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(313)),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	875: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(314)),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(315)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	879: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(316)),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(317)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	883: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(318)),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(319)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	887: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(244)),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	889: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
}

type ts_external_scanner_symbol_identifiers = int32

const ts_external_token_PITarget = 0
const ts_external_token__pi_content = 1
const ts_external_token_Comment = 2

var ts_external_scanner_symbol_map = [3]TSSymbol{
	0: uint16(sym_PITarget),
	1: uint16(sym__pi_content),
	2: uint16(sym_Comment),
}

var ts_external_scanner_states = [5][3]uint8{
	1: {
		0: libc.BoolUint8(true1 != 0),
		1: libc.BoolUint8(true1 != 0),
		2: libc.BoolUint8(true1 != 0),
	},
	2: {
		2: libc.BoolUint8(true1 != 0),
	},
	3: {
		0: libc.BoolUint8(true1 != 0),
	},
	4: {
		1: libc.BoolUint8(true1 != 0),
	},
}

func tree_sitter_dtd(tls *libc.TLS) (r uintptr) {
	return uintptr(unsafe.Pointer(&language))
}

var language = TSLanguage{
	Fabi_version:               uint32(LANGUAGE_VERSION),
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
	Fkeyword_capture_token:     uint16(sym_Name),
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
	*(*uintptr)(unsafe.Add(p, 184)) = __ccgo_fp(tree_sitter_dtd_external_scanner_create)
	*(*uintptr)(unsafe.Add(p, 192)) = __ccgo_fp(tree_sitter_dtd_external_scanner_destroy)
	*(*uintptr)(unsafe.Add(p, 200)) = __ccgo_fp(tree_sitter_dtd_external_scanner_scan)
	*(*uintptr)(unsafe.Add(p, 208)) = __ccgo_fp(tree_sitter_dtd_external_scanner_serialize)
	*(*uintptr)(unsafe.Add(p, 216)) = __ccgo_fp(tree_sitter_dtd_external_scanner_deserialize)
}

func __ccgo_fp(f interface{}) uintptr {
	type iface [2]uintptr
	return (*iface)(unsafe.Pointer(&f))[1]
}

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = "end\x00Name\x00<?\x00xml\x00?>\x00<![\x00IGNORE\x00INCLUDE\x00[\x00]]>\x00<!\x00ELEMENT\x00>\x00EMPTY\x00ANY\x00(\x00#PCDATA\x00|\x00)\x00*\x00?\x00+\x00,\x00ATTLIST\x00StringType\x00TokenizedType\x00NOTATION\x00#REQUIRED\x00#IMPLIED\x00#FIXED\x00ENTITY\x00%\x00\"\x00EntityValue_token1\x00'\x00EntityValue_token2\x00NDATA\x00;\x00_S\x00Nmtoken\x00&\x00&#\x00CharRef_token1\x00&#x\x00CharRef_token2\x00AttValue_token1\x00AttValue_token2\x00SYSTEM\x00PUBLIC\x00URI\x00PubidLiteral_token1\x00PubidLiteral_token2\x00version\x00VersionNum\x00encoding\x00EncName\x00=\x00PITarget\x00_pi_content\x00Comment\x00extSubset\x00TextDecl\x00_extSubsetDecl\x00conditionalSect\x00_markupdecl\x00_DeclSep\x00elementdecl\x00contentspec\x00Mixed\x00children\x00_cp\x00_choice\x00AttlistDecl\x00AttDef\x00_AttType\x00_EnumeratedType\x00NotationType\x00Enumeration\x00DefaultDecl\x00_EntityDecl\x00GEDecl\x00PEDecl\x00EntityValue\x00NDataDecl\x00NotationDecl\x00PEReference\x00_Reference\x00EntityRef\x00CharRef\x00AttValue\x00ExternalID\x00PublicID\x00SystemLiteral\x00PubidLiteral\x00_VersionInfo\x00_EncodingDecl\x00PI\x00_Eq\x00extSubset_repeat1\x00Mixed_repeat1\x00Mixed_repeat2\x00_choice_repeat1\x00_choice_repeat2\x00AttlistDecl_repeat1\x00NotationType_repeat1\x00Enumeration_repeat1\x00EntityValue_repeat1\x00EntityValue_repeat2\x00AttValue_repeat1\x00AttValue_repeat2\x00content\x00"
