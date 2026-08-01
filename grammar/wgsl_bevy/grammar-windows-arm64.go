// Code generated for windows/arm64 by 'ccgo -ignore-unsupported-alignment -ignore-unsupported-atomic-sizes -ignore-vector-functions --winapi-no-errno -ignore-link-errors -std=gnu11 -O0 -D__extension__= -D__forceinline=static inline -D__attribute__(...)= -D__declspec(x)= -D__cdecl= -D__stdcall= -D__fastcall= -D__thiscall= -D_cdecl= -D__restrict= -D__restrict__= -D__MINGW_EXTENSION= -D_X86INTRIN_H_INCLUDED -D_X86GPRINTRIN_H_INCLUDED -D_IMMINTRIN_H_INCLUDED -D_MMINTRIN_H_INCLUDED -D_XMMINTRIN_H_INCLUDED -D_EMMINTRIN_H_INCLUDED -D_PMMINTRIN_H_INCLUDED -D_MM3DNOW_H_INCLUDED -Dfunc=func_token -Dinterface=interface_token -Dselect=select_token -Dchan=chan_token -Dgo=go_token -Dmap=map_token -Dpackage=package_token -Dtype=type_token -Dvar=var_token -Dimport=import_token -Ddefer=defer_token -Dfallthrough=fallthrough_token -Drange=range_token -I C:\a\ccgo-tree-sitter\ccgo-tree-sitter\third-party\tree-sitter-wgsl-bevy\src -I C:\a\ccgo-tree-sitter\ccgo-tree-sitter\third-party\tree-sitter-wgsl-bevy -I C:\a\ccgo-tree-sitter\ccgo-tree-sitter\third-party -I C:\Users\runneradmin\.cache\workspaced\sources\github\7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2\lib\include -I C:\Users\runneradmin\.cache\workspaced\sources\github\7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2\lib\src combined.c -o grammar.go', DO NOT EDIT.

//go:build windows && arm64

package grammar_wgsl_bevy

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
const FIELD_COUNT = 15
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
const LARGE_STATE_COUNT = 32
const MAX_ALIAS_SEQUENCE_LENGTH = 9
const MAX_RESERVED_WORD_SET_SIZE = 0
const MB_LEN_MAX = 1
const MINGW_HAS_DDK_H = 1
const MINGW_HAS_SECURE_API = 1
const PRODUCTION_ID_COUNT = 26
const PTRDIFF_MAX = "__PTRDIFF_MAX__"
const RAND_MAX = 0x7fff
const SIZE_MAX = "__SIZE_MAX__"
const STATE_COUNT = 443
const SUPERTYPE_COUNT = 0
const SYMBOL_COUNT = 235
const TOKEN_COUNT = 146
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
const _DIGIT = 0x4
const _EMMINTRIN_H_INCLUDED = 1
const _FREEENTRY = 0
const _HEAP_MAXREQ = 0xFFFFFFFFFFFFFFE0
const _HEX = 0x80
const _IMMINTRIN_H_INCLUDED = 1
const _LEADBYTE = 0x8000
const _LOWER = 0x2
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
const _SPACE = 8
const _UPPER = 0x1
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

const BLOCK_COMMENT = 0

func tree_sitter_wgsl_bevy_external_scanner_create(tls *libc.TLS) (r uintptr) {
	return libc.UintptrFromInt32(0)
}

func tree_sitter_wgsl_bevy_external_scanner_destroy(tls *libc.TLS, p uintptr) {
}

func tree_sitter_wgsl_bevy_external_scanner_reset(tls *libc.TLS, p uintptr) {
}

func tree_sitter_wgsl_bevy_external_scanner_serialize(tls *libc.TLS, p uintptr, buffer uintptr) (r uint32) {
	return uint32(0)
}

func tree_sitter_wgsl_bevy_external_scanner_deserialize(tls *libc.TLS, p uintptr, b uintptr, n uint32) {
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
func tree_sitter_wgsl_bevy_external_scanner_scan(tls *libc.TLS, payload uintptr, lexer uintptr, valid_symbols uintptr) (r uint8) {
	var comment_depth uint32
	_ = comment_depth
	for libc.Xiswctype(tls, uint16((*TSLexer)(unsafe.Pointer(lexer)).Flookahead), uint16(_SPACE)) != 0 {
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

/* Automatically @generated by tree-sitter v0.25.9 */

type ts_symbol_identifiers = int32

const sym_identifier = 1
const sym_line_comment = 2
const anon_sym_SEMI = 3
const anon_sym_EQ = 4
const anon_sym_let = 5
const anon_sym_override = 6
const anon_sym_type = 7
const anon_sym_LPAREN = 8
const anon_sym_COMMA = 9
const anon_sym_RPAREN = 10
const anon_sym_virtual = 11
const anon_sym_fn = 12
const anon_sym_DASH_GT = 13
const anon_sym_struct = 14
const anon_sym_LBRACE = 15
const anon_sym_RBRACE = 16
const anon_sym_enable = 17
const anon_sym_AT = 18
const anon_sym__ = 19
const anon_sym_PLUS_EQ = 20
const anon_sym_DASH_EQ = 21
const anon_sym_STAR_EQ = 22
const anon_sym_SLASH_EQ = 23
const anon_sym_PERCENT_EQ = 24
const anon_sym_AMP_EQ = 25
const anon_sym_PIPE_EQ = 26
const anon_sym_CARET_EQ = 27
const anon_sym_if = 28
const anon_sym_else = 29
const anon_sym_switch = 30
const anon_sym_case = 31
const anon_sym_COLON = 32
const anon_sym_default = 33
const anon_sym_fallthrough = 34
const anon_sym_loop = 35
const anon_sym_for = 36
const anon_sym_while = 37
const anon_sym_break = 38
const anon_sym_continue = 39
const anon_sym_continuing = 40
const anon_sym_return = 41
const anon_sym_discard = 42
const anon_sym_var = 43
const anon_sym_LT = 44
const anon_sym_GT = 45
const anon_sym_PLUS_PLUS = 46
const anon_sym_DASH_DASH = 47
const sym_int_literal = 48
const aux_sym_float_literal_token1 = 49
const aux_sym_float_literal_token2 = 50
const anon_sym_true = 51
const anon_sym_false = 52
const anon_sym_bool = 53
const anon_sym_u32 = 54
const anon_sym_i32 = 55
const anon_sym_f32 = 56
const anon_sym_f16 = 57
const anon_sym_array = 58
const anon_sym_ptr = 59
const anon_sym_sampler = 60
const anon_sym_sampler_comparison = 61
const anon_sym_texture_depth_2d = 62
const anon_sym_texture_depth_2d_array = 63
const anon_sym_texture_depth_cube = 64
const anon_sym_texture_depth_cube_array = 65
const anon_sym_texture_depth_multisampled_2d = 66
const anon_sym_texture_1d = 67
const anon_sym_texture_2d = 68
const anon_sym_texture_2d_array = 69
const anon_sym_texture_3d = 70
const anon_sym_texture_cube = 71
const anon_sym_texture_cube_array = 72
const anon_sym_texture_multisampled_2d = 73
const anon_sym_texture_storage_1d = 74
const anon_sym_texture_storage_2d = 75
const anon_sym_texture_storage_2d_array = 76
const anon_sym_texture_storage_3d = 77
const anon_sym_vec2 = 78
const anon_sym_vec3 = 79
const anon_sym_vec4 = 80
const anon_sym_mat2x2 = 81
const anon_sym_mat2x3 = 82
const anon_sym_mat2x4 = 83
const anon_sym_mat3x2 = 84
const anon_sym_mat3x3 = 85
const anon_sym_mat3x4 = 86
const anon_sym_mat4x2 = 87
const anon_sym_mat4x3 = 88
const anon_sym_mat4x4 = 89
const anon_sym_rgba8unorm = 90
const anon_sym_rgba8snorm = 91
const anon_sym_rgba8uint = 92
const anon_sym_rgba8sint = 93
const anon_sym_rgba16uint = 94
const anon_sym_rgba16sint = 95
const anon_sym_rgba16float = 96
const anon_sym_r32uint = 97
const anon_sym_r32sint = 98
const anon_sym_r32float = 99
const anon_sym_rg32uint = 100
const anon_sym_rg32sint = 101
const anon_sym_rg32float = 102
const anon_sym_rgba32uint = 103
const anon_sym_rgba32sint = 104
const anon_sym_rgba32float = 105
const anon_sym_function = 106
const anon_sym_private = 107
const anon_sym_workgroup = 108
const anon_sym_uniform = 109
const anon_sym_storage = 110
const anon_sym_read = 111
const anon_sym_write = 112
const anon_sym_read_write = 113
const anon_sym_bitcast = 114
const anon_sym_PIPE_PIPE = 115
const anon_sym_AMP_AMP = 116
const anon_sym_PIPE = 117
const anon_sym_CARET = 118
const anon_sym_AMP = 119
const anon_sym_EQ_EQ = 120
const anon_sym_BANG_EQ = 121
const anon_sym_LT_EQ = 122
const anon_sym_GT_EQ = 123
const anon_sym_LT_LT = 124
const anon_sym_GT_GT = 125
const anon_sym_PLUS = 126
const anon_sym_DASH = 127
const anon_sym_STAR = 128
const anon_sym_SLASH = 129
const anon_sym_PERCENT = 130
const anon_sym_BANG = 131
const anon_sym_TILDE = 132
const anon_sym_LBRACK = 133
const anon_sym_RBRACK = 134
const anon_sym_DOT = 135
const aux_sym_preproc_import_token1 = 136
const anon_sym_LF = 137
const aux_sym_define_import_path_token1 = 138
const anon_sym_COLON_COLON = 139
const anon_sym_as = 140
const aux_sym_preproc_ifdef_token1 = 141
const aux_sym_preproc_ifdef_token2 = 142
const aux_sym_preproc_ifdef_token3 = 143
const aux_sym_preproc_else_token1 = 144
const sym_block_comment = 145
const sym_source_file = 146
const sym__declaration = 147
const sym_global_variable_declaration = 148
const sym_global_constant_declaration = 149
const sym_type_alias_declaration = 150
const sym_const_expression = 151
const sym_function_declaration = 152
const sym_function_return_type_declaration = 153
const sym_struct_declaration = 154
const sym_struct_member = 155
const sym_enable_directive = 156
const sym_attribute = 157
const sym__literal_or_identifier = 158
const sym_parameter_list = 159
const sym_parameter = 160
const sym__statement = 161
const sym_compound_statement = 162
const sym_assignment_statement = 163
const sym_compound_assignment_operator = 164
const sym_if_statement = 165
const sym_else_statement = 166
const sym_switch_statement = 167
const sym_switch_body = 168
const sym_case_selectors = 169
const sym_case_compound_statement = 170
const sym_fallthrough_statement = 171
const sym_loop_statement = 172
const sym_for_statement = 173
const sym_for_header = 174
const sym_while_statement = 175
const sym_break_statement = 176
const sym_break_if_statement = 177
const sym_continue_statement = 178
const sym_continuing_statement = 179
const sym_continuing_compound_statement = 180
const sym_return_statement = 181
const sym_discard_statement = 182
const sym_variable_statement = 183
const sym_variable_declaration = 184
const sym_variable_qualifier = 185
const sym_variable_identifier_declaration = 186
const sym_increment_statement = 187
const sym_decrement_statement = 188
const sym__expression = 189
const sym_const_literal = 190
const sym_float_literal = 191
const sym_bool_literal = 192
const sym_parenthesized_expression = 193
const sym_type_constructor_or_function_call_expression = 194
const sym_type_declaration = 195
const sym__vec_prefix = 196
const sym__mat_prefix = 197
const sym_texel_format = 198
const sym_address_space = 199
const sym_access_mode = 200
const sym_argument_list_expression = 201
const sym_bitcast_expression = 202
const sym_binary_expression = 203
const sym_unary_expression = 204
const sym_postfix_expression = 205
const sym_subscript_expression = 206
const sym_lhs_expression = 207
const sym_composite_value_decomposition_expression = 208
const sym__struct_declaration_content = 209
const sym_preproc_import = 210
const sym_define_import_path = 211
const sym_import_path = 212
const sym_alias = 213
const sym_preproc_ifdef = 214
const sym_preproc_else = 215
const sym_preproc_ifdef_in_statement = 216
const sym_preproc_else_in_statement = 217
const sym_preproc_ifdef_in_struct_declaration = 218
const sym_preproc_else_in_struct_declaration = 219
const aux_sym_source_file_repeat1 = 220
const aux_sym_source_file_repeat2 = 221
const aux_sym_global_variable_declaration_repeat1 = 222
const aux_sym_const_expression_repeat1 = 223
const aux_sym_attribute_repeat1 = 224
const aux_sym_parameter_list_repeat1 = 225
const aux_sym_compound_statement_repeat1 = 226
const aux_sym_switch_statement_repeat1 = 227
const aux_sym_case_selectors_repeat1 = 228
const aux_sym_argument_list_expression_repeat1 = 229
const aux_sym_lhs_expression_repeat1 = 230
const aux_sym__struct_declaration_content_repeat1 = 231
const aux_sym_preproc_import_repeat1 = 232
const aux_sym_import_path_repeat1 = 233
const aux_sym_preproc_ifdef_in_struct_declaration_repeat1 = 234

var ts_symbol_names = [235]uintptr{
	0:   __ccgo_ts,
	1:   __ccgo_ts + 4,
	2:   __ccgo_ts + 15,
	3:   __ccgo_ts + 28,
	4:   __ccgo_ts + 30,
	5:   __ccgo_ts + 32,
	6:   __ccgo_ts + 36,
	7:   __ccgo_ts + 45,
	8:   __ccgo_ts + 50,
	9:   __ccgo_ts + 52,
	10:  __ccgo_ts + 54,
	11:  __ccgo_ts + 56,
	12:  __ccgo_ts + 64,
	13:  __ccgo_ts + 67,
	14:  __ccgo_ts + 70,
	15:  __ccgo_ts + 77,
	16:  __ccgo_ts + 79,
	17:  __ccgo_ts + 81,
	18:  __ccgo_ts + 88,
	19:  __ccgo_ts + 90,
	20:  __ccgo_ts + 92,
	21:  __ccgo_ts + 95,
	22:  __ccgo_ts + 98,
	23:  __ccgo_ts + 101,
	24:  __ccgo_ts + 104,
	25:  __ccgo_ts + 107,
	26:  __ccgo_ts + 110,
	27:  __ccgo_ts + 113,
	28:  __ccgo_ts + 116,
	29:  __ccgo_ts + 119,
	30:  __ccgo_ts + 124,
	31:  __ccgo_ts + 131,
	32:  __ccgo_ts + 136,
	33:  __ccgo_ts + 138,
	34:  __ccgo_ts + 146,
	35:  __ccgo_ts + 158,
	36:  __ccgo_ts + 163,
	37:  __ccgo_ts + 167,
	38:  __ccgo_ts + 173,
	39:  __ccgo_ts + 179,
	40:  __ccgo_ts + 188,
	41:  __ccgo_ts + 199,
	42:  __ccgo_ts + 206,
	43:  __ccgo_ts + 214,
	44:  __ccgo_ts + 218,
	45:  __ccgo_ts + 220,
	46:  __ccgo_ts + 222,
	47:  __ccgo_ts + 225,
	48:  __ccgo_ts + 228,
	49:  __ccgo_ts + 240,
	50:  __ccgo_ts + 261,
	51:  __ccgo_ts + 282,
	52:  __ccgo_ts + 287,
	53:  __ccgo_ts + 293,
	54:  __ccgo_ts + 298,
	55:  __ccgo_ts + 302,
	56:  __ccgo_ts + 306,
	57:  __ccgo_ts + 310,
	58:  __ccgo_ts + 314,
	59:  __ccgo_ts + 320,
	60:  __ccgo_ts + 324,
	61:  __ccgo_ts + 332,
	62:  __ccgo_ts + 351,
	63:  __ccgo_ts + 368,
	64:  __ccgo_ts + 391,
	65:  __ccgo_ts + 410,
	66:  __ccgo_ts + 435,
	67:  __ccgo_ts + 465,
	68:  __ccgo_ts + 476,
	69:  __ccgo_ts + 487,
	70:  __ccgo_ts + 504,
	71:  __ccgo_ts + 515,
	72:  __ccgo_ts + 528,
	73:  __ccgo_ts + 547,
	74:  __ccgo_ts + 571,
	75:  __ccgo_ts + 590,
	76:  __ccgo_ts + 609,
	77:  __ccgo_ts + 634,
	78:  __ccgo_ts + 653,
	79:  __ccgo_ts + 658,
	80:  __ccgo_ts + 663,
	81:  __ccgo_ts + 668,
	82:  __ccgo_ts + 675,
	83:  __ccgo_ts + 682,
	84:  __ccgo_ts + 689,
	85:  __ccgo_ts + 696,
	86:  __ccgo_ts + 703,
	87:  __ccgo_ts + 710,
	88:  __ccgo_ts + 717,
	89:  __ccgo_ts + 724,
	90:  __ccgo_ts + 731,
	91:  __ccgo_ts + 742,
	92:  __ccgo_ts + 753,
	93:  __ccgo_ts + 763,
	94:  __ccgo_ts + 773,
	95:  __ccgo_ts + 784,
	96:  __ccgo_ts + 795,
	97:  __ccgo_ts + 807,
	98:  __ccgo_ts + 815,
	99:  __ccgo_ts + 823,
	100: __ccgo_ts + 832,
	101: __ccgo_ts + 841,
	102: __ccgo_ts + 850,
	103: __ccgo_ts + 860,
	104: __ccgo_ts + 871,
	105: __ccgo_ts + 882,
	106: __ccgo_ts + 894,
	107: __ccgo_ts + 903,
	108: __ccgo_ts + 911,
	109: __ccgo_ts + 921,
	110: __ccgo_ts + 929,
	111: __ccgo_ts + 937,
	112: __ccgo_ts + 942,
	113: __ccgo_ts + 948,
	114: __ccgo_ts + 959,
	115: __ccgo_ts + 967,
	116: __ccgo_ts + 970,
	117: __ccgo_ts + 973,
	118: __ccgo_ts + 975,
	119: __ccgo_ts + 977,
	120: __ccgo_ts + 979,
	121: __ccgo_ts + 982,
	122: __ccgo_ts + 985,
	123: __ccgo_ts + 988,
	124: __ccgo_ts + 991,
	125: __ccgo_ts + 994,
	126: __ccgo_ts + 997,
	127: __ccgo_ts + 999,
	128: __ccgo_ts + 1001,
	129: __ccgo_ts + 1003,
	130: __ccgo_ts + 1005,
	131: __ccgo_ts + 1007,
	132: __ccgo_ts + 1009,
	133: __ccgo_ts + 1011,
	134: __ccgo_ts + 1013,
	135: __ccgo_ts + 1015,
	136: __ccgo_ts + 1017,
	137: __ccgo_ts + 1025,
	138: __ccgo_ts + 1027,
	139: __ccgo_ts + 1047,
	140: __ccgo_ts + 1050,
	141: __ccgo_ts + 1053,
	142: __ccgo_ts + 1060,
	143: __ccgo_ts + 1068,
	144: __ccgo_ts + 1075,
	145: __ccgo_ts + 1081,
	146: __ccgo_ts + 1095,
	147: __ccgo_ts + 1107,
	148: __ccgo_ts + 1120,
	149: __ccgo_ts + 1148,
	150: __ccgo_ts + 1176,
	151: __ccgo_ts + 1199,
	152: __ccgo_ts + 1216,
	153: __ccgo_ts + 1237,
	154: __ccgo_ts + 1270,
	155: __ccgo_ts + 1289,
	156: __ccgo_ts + 1303,
	157: __ccgo_ts + 1320,
	158: __ccgo_ts + 1330,
	159: __ccgo_ts + 1353,
	160: __ccgo_ts + 1368,
	161: __ccgo_ts + 1378,
	162: __ccgo_ts + 1389,
	163: __ccgo_ts + 1408,
	164: __ccgo_ts + 1429,
	165: __ccgo_ts + 1458,
	166: __ccgo_ts + 1471,
	167: __ccgo_ts + 1486,
	168: __ccgo_ts + 1503,
	169: __ccgo_ts + 1515,
	170: __ccgo_ts + 1530,
	171: __ccgo_ts + 1554,
	172: __ccgo_ts + 1576,
	173: __ccgo_ts + 1591,
	174: __ccgo_ts + 1605,
	175: __ccgo_ts + 1616,
	176: __ccgo_ts + 1632,
	177: __ccgo_ts + 1648,
	178: __ccgo_ts + 1667,
	179: __ccgo_ts + 1686,
	180: __ccgo_ts + 1707,
	181: __ccgo_ts + 1737,
	182: __ccgo_ts + 1754,
	183: __ccgo_ts + 1772,
	184: __ccgo_ts + 1791,
	185: __ccgo_ts + 1812,
	186: __ccgo_ts + 1831,
	187: __ccgo_ts + 1863,
	188: __ccgo_ts + 1883,
	189: __ccgo_ts + 1903,
	190: __ccgo_ts + 1915,
	191: __ccgo_ts + 1929,
	192: __ccgo_ts + 1943,
	193: __ccgo_ts + 1956,
	194: __ccgo_ts + 1981,
	195: __ccgo_ts + 2026,
	196: __ccgo_ts + 2043,
	197: __ccgo_ts + 2055,
	198: __ccgo_ts + 2067,
	199: __ccgo_ts + 2080,
	200: __ccgo_ts + 2094,
	201: __ccgo_ts + 2106,
	202: __ccgo_ts + 2131,
	203: __ccgo_ts + 2150,
	204: __ccgo_ts + 2168,
	205: __ccgo_ts + 2185,
	206: __ccgo_ts + 2204,
	207: __ccgo_ts + 2225,
	208: __ccgo_ts + 2240,
	209: __ccgo_ts + 2281,
	210: __ccgo_ts + 2309,
	211: __ccgo_ts + 2324,
	212: __ccgo_ts + 2343,
	213: __ccgo_ts + 2355,
	214: __ccgo_ts + 2361,
	215: __ccgo_ts + 2375,
	216: __ccgo_ts + 2361,
	217: __ccgo_ts + 2375,
	218: __ccgo_ts + 2361,
	219: __ccgo_ts + 2375,
	220: __ccgo_ts + 2388,
	221: __ccgo_ts + 2408,
	222: __ccgo_ts + 2428,
	223: __ccgo_ts + 2464,
	224: __ccgo_ts + 2489,
	225: __ccgo_ts + 2507,
	226: __ccgo_ts + 2530,
	227: __ccgo_ts + 2557,
	228: __ccgo_ts + 2582,
	229: __ccgo_ts + 2605,
	230: __ccgo_ts + 2638,
	231: __ccgo_ts + 2661,
	232: __ccgo_ts + 2697,
	233: __ccgo_ts + 2720,
	234: __ccgo_ts + 2740,
}

var ts_symbol_map = [235]TSSymbol{
	1:   uint16(sym_identifier),
	2:   uint16(sym_line_comment),
	3:   uint16(anon_sym_SEMI),
	4:   uint16(anon_sym_EQ),
	5:   uint16(anon_sym_let),
	6:   uint16(anon_sym_override),
	7:   uint16(anon_sym_type),
	8:   uint16(anon_sym_LPAREN),
	9:   uint16(anon_sym_COMMA),
	10:  uint16(anon_sym_RPAREN),
	11:  uint16(anon_sym_virtual),
	12:  uint16(anon_sym_fn),
	13:  uint16(anon_sym_DASH_GT),
	14:  uint16(anon_sym_struct),
	15:  uint16(anon_sym_LBRACE),
	16:  uint16(anon_sym_RBRACE),
	17:  uint16(anon_sym_enable),
	18:  uint16(anon_sym_AT),
	19:  uint16(anon_sym__),
	20:  uint16(anon_sym_PLUS_EQ),
	21:  uint16(anon_sym_DASH_EQ),
	22:  uint16(anon_sym_STAR_EQ),
	23:  uint16(anon_sym_SLASH_EQ),
	24:  uint16(anon_sym_PERCENT_EQ),
	25:  uint16(anon_sym_AMP_EQ),
	26:  uint16(anon_sym_PIPE_EQ),
	27:  uint16(anon_sym_CARET_EQ),
	28:  uint16(anon_sym_if),
	29:  uint16(anon_sym_else),
	30:  uint16(anon_sym_switch),
	31:  uint16(anon_sym_case),
	32:  uint16(anon_sym_COLON),
	33:  uint16(anon_sym_default),
	34:  uint16(anon_sym_fallthrough),
	35:  uint16(anon_sym_loop),
	36:  uint16(anon_sym_for),
	37:  uint16(anon_sym_while),
	38:  uint16(anon_sym_break),
	39:  uint16(anon_sym_continue),
	40:  uint16(anon_sym_continuing),
	41:  uint16(anon_sym_return),
	42:  uint16(anon_sym_discard),
	43:  uint16(anon_sym_var),
	44:  uint16(anon_sym_LT),
	45:  uint16(anon_sym_GT),
	46:  uint16(anon_sym_PLUS_PLUS),
	47:  uint16(anon_sym_DASH_DASH),
	48:  uint16(sym_int_literal),
	49:  uint16(aux_sym_float_literal_token1),
	50:  uint16(aux_sym_float_literal_token2),
	51:  uint16(anon_sym_true),
	52:  uint16(anon_sym_false),
	53:  uint16(anon_sym_bool),
	54:  uint16(anon_sym_u32),
	55:  uint16(anon_sym_i32),
	56:  uint16(anon_sym_f32),
	57:  uint16(anon_sym_f16),
	58:  uint16(anon_sym_array),
	59:  uint16(anon_sym_ptr),
	60:  uint16(anon_sym_sampler),
	61:  uint16(anon_sym_sampler_comparison),
	62:  uint16(anon_sym_texture_depth_2d),
	63:  uint16(anon_sym_texture_depth_2d_array),
	64:  uint16(anon_sym_texture_depth_cube),
	65:  uint16(anon_sym_texture_depth_cube_array),
	66:  uint16(anon_sym_texture_depth_multisampled_2d),
	67:  uint16(anon_sym_texture_1d),
	68:  uint16(anon_sym_texture_2d),
	69:  uint16(anon_sym_texture_2d_array),
	70:  uint16(anon_sym_texture_3d),
	71:  uint16(anon_sym_texture_cube),
	72:  uint16(anon_sym_texture_cube_array),
	73:  uint16(anon_sym_texture_multisampled_2d),
	74:  uint16(anon_sym_texture_storage_1d),
	75:  uint16(anon_sym_texture_storage_2d),
	76:  uint16(anon_sym_texture_storage_2d_array),
	77:  uint16(anon_sym_texture_storage_3d),
	78:  uint16(anon_sym_vec2),
	79:  uint16(anon_sym_vec3),
	80:  uint16(anon_sym_vec4),
	81:  uint16(anon_sym_mat2x2),
	82:  uint16(anon_sym_mat2x3),
	83:  uint16(anon_sym_mat2x4),
	84:  uint16(anon_sym_mat3x2),
	85:  uint16(anon_sym_mat3x3),
	86:  uint16(anon_sym_mat3x4),
	87:  uint16(anon_sym_mat4x2),
	88:  uint16(anon_sym_mat4x3),
	89:  uint16(anon_sym_mat4x4),
	90:  uint16(anon_sym_rgba8unorm),
	91:  uint16(anon_sym_rgba8snorm),
	92:  uint16(anon_sym_rgba8uint),
	93:  uint16(anon_sym_rgba8sint),
	94:  uint16(anon_sym_rgba16uint),
	95:  uint16(anon_sym_rgba16sint),
	96:  uint16(anon_sym_rgba16float),
	97:  uint16(anon_sym_r32uint),
	98:  uint16(anon_sym_r32sint),
	99:  uint16(anon_sym_r32float),
	100: uint16(anon_sym_rg32uint),
	101: uint16(anon_sym_rg32sint),
	102: uint16(anon_sym_rg32float),
	103: uint16(anon_sym_rgba32uint),
	104: uint16(anon_sym_rgba32sint),
	105: uint16(anon_sym_rgba32float),
	106: uint16(anon_sym_function),
	107: uint16(anon_sym_private),
	108: uint16(anon_sym_workgroup),
	109: uint16(anon_sym_uniform),
	110: uint16(anon_sym_storage),
	111: uint16(anon_sym_read),
	112: uint16(anon_sym_write),
	113: uint16(anon_sym_read_write),
	114: uint16(anon_sym_bitcast),
	115: uint16(anon_sym_PIPE_PIPE),
	116: uint16(anon_sym_AMP_AMP),
	117: uint16(anon_sym_PIPE),
	118: uint16(anon_sym_CARET),
	119: uint16(anon_sym_AMP),
	120: uint16(anon_sym_EQ_EQ),
	121: uint16(anon_sym_BANG_EQ),
	122: uint16(anon_sym_LT_EQ),
	123: uint16(anon_sym_GT_EQ),
	124: uint16(anon_sym_LT_LT),
	125: uint16(anon_sym_GT_GT),
	126: uint16(anon_sym_PLUS),
	127: uint16(anon_sym_DASH),
	128: uint16(anon_sym_STAR),
	129: uint16(anon_sym_SLASH),
	130: uint16(anon_sym_PERCENT),
	131: uint16(anon_sym_BANG),
	132: uint16(anon_sym_TILDE),
	133: uint16(anon_sym_LBRACK),
	134: uint16(anon_sym_RBRACK),
	135: uint16(anon_sym_DOT),
	136: uint16(aux_sym_preproc_import_token1),
	137: uint16(anon_sym_LF),
	138: uint16(aux_sym_define_import_path_token1),
	139: uint16(anon_sym_COLON_COLON),
	140: uint16(anon_sym_as),
	141: uint16(aux_sym_preproc_ifdef_token1),
	142: uint16(aux_sym_preproc_ifdef_token2),
	143: uint16(aux_sym_preproc_ifdef_token3),
	144: uint16(aux_sym_preproc_else_token1),
	145: uint16(sym_block_comment),
	146: uint16(sym_source_file),
	147: uint16(sym__declaration),
	148: uint16(sym_global_variable_declaration),
	149: uint16(sym_global_constant_declaration),
	150: uint16(sym_type_alias_declaration),
	151: uint16(sym_const_expression),
	152: uint16(sym_function_declaration),
	153: uint16(sym_function_return_type_declaration),
	154: uint16(sym_struct_declaration),
	155: uint16(sym_struct_member),
	156: uint16(sym_enable_directive),
	157: uint16(sym_attribute),
	158: uint16(sym__literal_or_identifier),
	159: uint16(sym_parameter_list),
	160: uint16(sym_parameter),
	161: uint16(sym__statement),
	162: uint16(sym_compound_statement),
	163: uint16(sym_assignment_statement),
	164: uint16(sym_compound_assignment_operator),
	165: uint16(sym_if_statement),
	166: uint16(sym_else_statement),
	167: uint16(sym_switch_statement),
	168: uint16(sym_switch_body),
	169: uint16(sym_case_selectors),
	170: uint16(sym_case_compound_statement),
	171: uint16(sym_fallthrough_statement),
	172: uint16(sym_loop_statement),
	173: uint16(sym_for_statement),
	174: uint16(sym_for_header),
	175: uint16(sym_while_statement),
	176: uint16(sym_break_statement),
	177: uint16(sym_break_if_statement),
	178: uint16(sym_continue_statement),
	179: uint16(sym_continuing_statement),
	180: uint16(sym_continuing_compound_statement),
	181: uint16(sym_return_statement),
	182: uint16(sym_discard_statement),
	183: uint16(sym_variable_statement),
	184: uint16(sym_variable_declaration),
	185: uint16(sym_variable_qualifier),
	186: uint16(sym_variable_identifier_declaration),
	187: uint16(sym_increment_statement),
	188: uint16(sym_decrement_statement),
	189: uint16(sym__expression),
	190: uint16(sym_const_literal),
	191: uint16(sym_float_literal),
	192: uint16(sym_bool_literal),
	193: uint16(sym_parenthesized_expression),
	194: uint16(sym_type_constructor_or_function_call_expression),
	195: uint16(sym_type_declaration),
	196: uint16(sym__vec_prefix),
	197: uint16(sym__mat_prefix),
	198: uint16(sym_texel_format),
	199: uint16(sym_address_space),
	200: uint16(sym_access_mode),
	201: uint16(sym_argument_list_expression),
	202: uint16(sym_bitcast_expression),
	203: uint16(sym_binary_expression),
	204: uint16(sym_unary_expression),
	205: uint16(sym_postfix_expression),
	206: uint16(sym_subscript_expression),
	207: uint16(sym_lhs_expression),
	208: uint16(sym_composite_value_decomposition_expression),
	209: uint16(sym__struct_declaration_content),
	210: uint16(sym_preproc_import),
	211: uint16(sym_define_import_path),
	212: uint16(sym_import_path),
	213: uint16(sym_alias),
	214: uint16(sym_preproc_ifdef),
	215: uint16(sym_preproc_else),
	216: uint16(sym_preproc_ifdef),
	217: uint16(sym_preproc_else),
	218: uint16(sym_preproc_ifdef),
	219: uint16(sym_preproc_else),
	220: uint16(aux_sym_source_file_repeat1),
	221: uint16(aux_sym_source_file_repeat2),
	222: uint16(aux_sym_global_variable_declaration_repeat1),
	223: uint16(aux_sym_const_expression_repeat1),
	224: uint16(aux_sym_attribute_repeat1),
	225: uint16(aux_sym_parameter_list_repeat1),
	226: uint16(aux_sym_compound_statement_repeat1),
	227: uint16(aux_sym_switch_statement_repeat1),
	228: uint16(aux_sym_case_selectors_repeat1),
	229: uint16(aux_sym_argument_list_expression_repeat1),
	230: uint16(aux_sym_lhs_expression_repeat1),
	231: uint16(aux_sym__struct_declaration_content_repeat1),
	232: uint16(aux_sym_preproc_import_repeat1),
	233: uint16(aux_sym_import_path_repeat1),
	234: uint16(aux_sym_preproc_ifdef_in_struct_declaration_repeat1),
}

var ts_symbol_metadata = [235]TSSymbolMetadata{
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
	},
	48: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	49: {},
	50: {},
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
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	146: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	147: {
		Fnamed: libc.BoolUint8(true1 != 0),
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
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
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
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	187: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	188: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	189: {
		Fnamed: libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	197: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	198: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	199: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	200: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	201: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	202: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	203: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	204: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	205: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	206: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	207: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	208: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	209: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	210: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	211: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	212: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	213: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	214: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	215: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	216: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	217: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	218: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	219: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	220: {},
	221: {},
	222: {},
	223: {},
	224: {},
	225: {},
	226: {},
	227: {},
	228: {},
	229: {},
	230: {},
	231: {},
	232: {},
	233: {},
	234: {},
}

type ts_field_identifiers = int32

const field_accessor = 1
const field_alias = 2
const field_alternative = 3
const field_argument = 4
const field_body = 5
const field_condition = 6
const field_consequence = 7
const field_left = 8
const field_name = 9
const field_parameters = 10
const field_path = 11
const field_right = 12
const field_subscript = 13
const field_type = 14
const field_value = 15

var ts_field_names = [16]uintptr{
	0:  libc.UintptrFromInt32(0),
	1:  __ccgo_ts + 2784,
	2:  __ccgo_ts + 2355,
	3:  __ccgo_ts + 2793,
	4:  __ccgo_ts + 2805,
	5:  __ccgo_ts + 2814,
	6:  __ccgo_ts + 2819,
	7:  __ccgo_ts + 2829,
	8:  __ccgo_ts + 2841,
	9:  __ccgo_ts + 2846,
	10: __ccgo_ts + 2851,
	11: __ccgo_ts + 2862,
	12: __ccgo_ts + 2867,
	13: __ccgo_ts + 2873,
	14: __ccgo_ts + 45,
	15: __ccgo_ts + 2883,
}

var ts_field_map_slices = [26]TSMapSlice{
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
		Flength: uint16(1),
	},
	7: {
		Findex:  uint16(9),
		Flength: uint16(2),
	},
	8: {
		Findex:  uint16(11),
		Flength: uint16(2),
	},
	9: {
		Findex:  uint16(13),
		Flength: uint16(2),
	},
	10: {
		Findex:  uint16(15),
		Flength: uint16(2),
	},
	11: {
		Findex:  uint16(17),
		Flength: uint16(2),
	},
	12: {
		Findex:  uint16(19),
		Flength: uint16(3),
	},
	13: {
		Findex:  uint16(22),
		Flength: uint16(3),
	},
	14: {
		Findex:  uint16(25),
		Flength: uint16(3),
	},
	15: {
		Findex:  uint16(28),
		Flength: uint16(3),
	},
	16: {
		Findex:  uint16(31),
		Flength: uint16(2),
	},
	17: {
		Findex:  uint16(33),
		Flength: uint16(2),
	},
	18: {
		Findex:  uint16(35),
		Flength: uint16(4),
	},
	19: {
		Findex:  uint16(39),
		Flength: uint16(4),
	},
	20: {
		Findex:  uint16(43),
		Flength: uint16(3),
	},
	21: {
		Findex:  uint16(46),
		Flength: uint16(3),
	},
	22: {
		Findex:  uint16(49),
		Flength: uint16(2),
	},
	23: {
		Findex:  uint16(51),
		Flength: uint16(1),
	},
	24: {
		Findex:  uint16(52),
		Flength: uint16(4),
	},
	25: {
		Findex:  uint16(56),
		Flength: uint16(3),
	},
}

var ts_field_map_entries = [59]TSFieldMapEntry{
	0: {
		Ffield_id:    uint16(field_path),
		Fchild_index: uint8(1),
	},
	1: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
	},
	2: {
		Ffield_id: uint16(field_name),
	},
	3: {
		Ffield_id:    uint16(field_type),
		Fchild_index: uint8(2),
	},
	4: {
		Ffield_id:    uint16(field_alias),
		Fchild_index: uint8(2),
	},
	5: {
		Ffield_id:    uint16(field_path),
		Fchild_index: uint8(1),
	},
	6: {
		Ffield_id:    uint16(field_alternative),
		Fchild_index: uint8(2),
	},
	7: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
	},
	8: {
		Ffield_id:    uint16(field_argument),
		Fchild_index: uint8(1),
	},
	9: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(4),
	},
	10: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
	},
	11: {
		Ffield_id:    uint16(field_alternative),
		Fchild_index: uint8(3),
	},
	12: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
	},
	13: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(5),
	},
	14: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(2),
	},
	15: {
		Ffield_id: uint16(field_left),
	},
	16: {
		Ffield_id:    uint16(field_right),
		Fchild_index: uint8(2),
	},
	17: {
		Ffield_id:    uint16(field_accessor),
		Fchild_index: uint8(2),
	},
	18: {
		Ffield_id: uint16(field_value),
	},
	19: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(5),
	},
	20: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
	},
	21: {
		Ffield_id:    uint16(field_type),
		Fchild_index: uint8(4),
	},
	22: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(5),
	},
	23: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
	},
	24: {
		Ffield_id:    uint16(field_parameters),
		Fchild_index: uint8(3),
	},
	25: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(6),
	},
	26: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(2),
	},
	27: {
		Ffield_id:    uint16(field_type),
		Fchild_index: uint8(5),
	},
	28: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(6),
	},
	29: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(2),
	},
	30: {
		Ffield_id:    uint16(field_parameters),
		Fchild_index: uint8(4),
	},
	31: {
		Ffield_id:    uint16(field_subscript),
		Fchild_index: uint8(2),
	},
	32: {
		Ffield_id: uint16(field_value),
	},
	33: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(6),
	},
	34: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(3),
	},
	35: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(6),
	},
	36: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
	},
	37: {
		Ffield_id:    uint16(field_parameters),
		Fchild_index: uint8(3),
	},
	38: {
		Ffield_id:    uint16(field_type),
		Fchild_index: uint8(5),
	},
	39: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(7),
	},
	40: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(2),
	},
	41: {
		Ffield_id:    uint16(field_parameters),
		Fchild_index: uint8(4),
	},
	42: {
		Ffield_id:    uint16(field_type),
		Fchild_index: uint8(6),
	},
	43: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(7),
	},
	44: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(3),
	},
	45: {
		Ffield_id:    uint16(field_type),
		Fchild_index: uint8(6),
	},
	46: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(7),
	},
	47: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(3),
	},
	48: {
		Ffield_id:    uint16(field_parameters),
		Fchild_index: uint8(5),
	},
	49: {
		Ffield_id:    uint16(field_condition),
		Fchild_index: uint8(1),
	},
	50: {
		Ffield_id:    uint16(field_consequence),
		Fchild_index: uint8(2),
	},
	51: {
		Ffield_id:    uint16(field_condition),
		Fchild_index: uint8(1),
	},
	52: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(8),
	},
	53: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(3),
	},
	54: {
		Ffield_id:    uint16(field_parameters),
		Fchild_index: uint8(5),
	},
	55: {
		Ffield_id:    uint16(field_type),
		Fchild_index: uint8(7),
	},
	56: {
		Ffield_id:    uint16(field_alternative),
		Fchild_index: uint8(4),
	},
	57: {
		Ffield_id:    uint16(field_condition),
		Fchild_index: uint8(1),
	},
	58: {
		Ffield_id:    uint16(field_consequence),
		Fchild_index: uint8(2),
	},
}

var ts_alias_sequences = [26][9]TSSymbol{}

var ts_non_terminal_alias_map = [1]uint16_t{}

var ts_primary_state_ids = [443]TSStateId{
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
	25:  uint16(8),
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
	150: uint16(137),
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
	187: uint16(186),
	188: uint16(177),
	189: uint16(178),
	190: uint16(182),
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
	288: uint16(288),
	289: uint16(289),
	290: uint16(290),
	291: uint16(291),
	292: uint16(292),
	293: uint16(248),
	294: uint16(241),
	295: uint16(238),
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
	312: uint16(312),
	313: uint16(313),
	314: uint16(314),
	315: uint16(315),
	316: uint16(316),
	317: uint16(317),
	318: uint16(318),
	319: uint16(319),
	320: uint16(320),
	321: uint16(321),
	322: uint16(322),
	323: uint16(323),
	324: uint16(324),
	325: uint16(325),
	326: uint16(326),
	327: uint16(327),
	328: uint16(328),
	329: uint16(329),
	330: uint16(330),
	331: uint16(331),
	332: uint16(332),
	333: uint16(333),
	334: uint16(334),
	335: uint16(335),
	336: uint16(336),
	337: uint16(337),
	338: uint16(338),
	339: uint16(261),
	340: uint16(340),
	341: uint16(341),
	342: uint16(342),
	343: uint16(343),
	344: uint16(344),
	345: uint16(345),
	346: uint16(346),
	347: uint16(347),
	348: uint16(348),
	349: uint16(349),
	350: uint16(350),
	351: uint16(351),
	352: uint16(352),
	353: uint16(353),
	354: uint16(354),
	355: uint16(355),
	356: uint16(356),
	357: uint16(357),
	358: uint16(358),
	359: uint16(359),
	360: uint16(360),
	361: uint16(361),
	362: uint16(362),
	363: uint16(363),
	364: uint16(364),
	365: uint16(365),
	366: uint16(366),
	367: uint16(367),
	368: uint16(368),
	369: uint16(369),
	370: uint16(370),
	371: uint16(371),
	372: uint16(372),
	373: uint16(373),
	374: uint16(374),
	375: uint16(375),
	376: uint16(376),
	377: uint16(377),
	378: uint16(378),
	379: uint16(379),
	380: uint16(380),
	381: uint16(381),
	382: uint16(382),
	383: uint16(383),
	384: uint16(384),
	385: uint16(385),
	386: uint16(386),
	387: uint16(387),
	388: uint16(388),
	389: uint16(389),
	390: uint16(390),
	391: uint16(391),
	392: uint16(392),
	393: uint16(393),
	394: uint16(394),
	395: uint16(395),
	396: uint16(396),
	397: uint16(397),
	398: uint16(398),
	399: uint16(399),
	400: uint16(400),
	401: uint16(401),
	402: uint16(402),
	403: uint16(403),
	404: uint16(404),
	405: uint16(405),
	406: uint16(406),
	407: uint16(407),
	408: uint16(408),
	409: uint16(409),
	410: uint16(410),
	411: uint16(411),
	412: uint16(412),
	413: uint16(413),
	414: uint16(414),
	415: uint16(415),
	416: uint16(416),
	417: uint16(417),
	418: uint16(418),
	419: uint16(419),
	420: uint16(420),
	421: uint16(421),
	422: uint16(422),
	423: uint16(423),
	424: uint16(424),
	425: uint16(425),
	426: uint16(426),
	427: uint16(427),
	428: uint16(428),
	429: uint16(429),
	430: uint16(430),
	431: uint16(431),
	432: uint16(432),
	433: uint16(433),
	434: uint16(434),
	435: uint16(435),
	436: uint16(436),
	437: uint16(437),
	438: uint16(365),
	439: uint16(392),
	440: uint16(440),
	441: uint16(441),
	442: uint16(442),
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
	switch int32(state) {
	case 0:
		if eof != 0 {
			state = uint16(65)
			goto next_state
		}
		i = uint32(0)
		for {
			if !(uint64(i) < libc.Uint64FromInt64(108)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token[i]) == lookahead {
				state = map_token[i+uint32(1)]
				goto next_state
			}
			goto _1
		_1:
			;
			i = i + uint32(2)
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(97)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') || lookahead == int32(0xa0) || lookahead == int32(0x200b) || lookahead == int32(0x2060) || lookahead == int32(0xfeff) {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(0)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(77)
			goto next_state
		}
		return result
	case int32(1):
		if lookahead == int32('\n') {
			state = uint16(143)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(71)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(10)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(12)
			goto next_state
		}
		if lookahead == int32('_') {
			state = uint16(63)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') || lookahead == int32(0xa0) || lookahead == int32(0x200b) || lookahead == int32(0x2060) || lookahead == int32(0xfeff) {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(1)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(77)
			goto next_state
		}
		return result
	case int32(2):
		i1 = uint32(0)
		for {
			if !(uint64(i1) < libc.Uint64FromInt64(84)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token1[i1]) == lookahead {
				state = map_token1[i1+uint32(1)]
				goto next_state
			}
			goto _2
		_2:
			;
			i1 = i1 + uint32(2)
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') || lookahead == int32(0xa0) || lookahead == int32(0x200b) || lookahead == int32(0x2060) || lookahead == int32(0xfeff) {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(2)
			goto next_state
		}
		return result
	case int32(3):
		i2 = uint32(0)
		for {
			if !(uint64(i2) < libc.Uint64FromInt64(68)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token2[i2]) == lookahead {
				state = map_token2[i2+uint32(1)]
				goto next_state
			}
			goto _3
		_3:
			;
			i2 = i2 + uint32(2)
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(101)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') || lookahead == int32(0xa0) || lookahead == int32(0x200b) || lookahead == int32(0x2060) || lookahead == int32(0xfeff) {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(3)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(77)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead == int32('+') {
			state = uint16(93)
			goto next_state
		}
		if lookahead == int32('=') {
			state = uint16(80)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead == int32('-') {
			state = uint16(94)
			goto next_state
		}
		if lookahead == int32('0') {
			state = uint16(57)
			goto next_state
		}
		if lookahead == int32('=') {
			state = uint16(81)
			goto next_state
		}
		if lookahead == int32('>') {
			state = uint16(73)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(101)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead == int32('.') {
			state = uint16(104)
			goto next_state
		}
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(55)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(6)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead == int32('.') {
			state = uint16(104)
			goto next_state
		}
		if lookahead == int32('X') || lookahead == int32('x') {
			state = uint16(9)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(8)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead == int32('.') {
			state = uint16(104)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(8)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead == int32('.') {
			state = uint16(61)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(99)
			goto next_state
		}
		return result
	case int32(10):
		if lookahead == int32('/') {
			state = uint16(66)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead == int32('/') {
			state = uint16(66)
			goto next_state
		}
		if lookahead == int32('=') {
			state = uint16(83)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead == int32(':') {
			state = uint16(145)
			goto next_state
		}
		return result
	case int32(13):
		if lookahead == int32('=') {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(14):
		if lookahead == int32('=') {
			state = uint16(84)
			goto next_state
		}
		return result
	case int32(15):
		if lookahead == int32('=') {
			state = uint16(85)
			goto next_state
		}
		return result
	case int32(16):
		if lookahead == int32('=') {
			state = uint16(82)
			goto next_state
		}
		return result
	case int32(17):
		if lookahead == int32('=') {
			state = uint16(118)
			goto next_state
		}
		return result
	case int32(18):
		if lookahead == int32('=') {
			state = uint16(87)
			goto next_state
		}
		return result
	case int32(19):
		if lookahead == int32('=') {
			state = uint16(86)
			goto next_state
		}
		return result
	case int32(20):
		if lookahead == int32('_') {
			state = uint16(47)
			goto next_state
		}
		return result
	case int32(21):
		if lookahead == int32('_') {
			state = uint16(39)
			goto next_state
		}
		return result
	case int32(22):
		if lookahead == int32('a') {
			state = uint16(53)
			goto next_state
		}
		return result
	case int32(23):
		if lookahead == int32('d') {
			state = uint16(27)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(41)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(32)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(24):
		if lookahead == int32('d') {
			state = uint16(40)
			goto next_state
		}
		return result
	case int32(25):
		if lookahead == int32('d') {
			state = uint16(30)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(26)
			goto next_state
		}
		return result
	case int32(26):
		if lookahead == int32('d') {
			state = uint16(31)
			goto next_state
		}
		return result
	case int32(27):
		if lookahead == int32('e') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(28):
		if lookahead == int32('e') {
			state = uint16(149)
			goto next_state
		}
		return result
	case int32(29):
		if lookahead == int32('e') {
			state = uint16(21)
			goto next_state
		}
		return result
	case int32(30):
		if lookahead == int32('e') {
			state = uint16(35)
			goto next_state
		}
		return result
	case int32(31):
		if lookahead == int32('e') {
			state = uint16(36)
			goto next_state
		}
		return result
	case int32(32):
		if lookahead == int32('f') {
			state = uint16(25)
			goto next_state
		}
		if lookahead == int32('m') {
			state = uint16(46)
			goto next_state
		}
		return result
	case int32(33):
		if lookahead == int32('f') {
			state = uint16(38)
			goto next_state
		}
		return result
	case int32(34):
		if lookahead == int32('f') {
			state = uint16(148)
			goto next_state
		}
		return result
	case int32(35):
		if lookahead == int32('f') {
			state = uint16(146)
			goto next_state
		}
		return result
	case int32(36):
		if lookahead == int32('f') {
			state = uint16(147)
			goto next_state
		}
		return result
	case int32(37):
		if lookahead == int32('h') {
			state = uint16(144)
			goto next_state
		}
		return result
	case int32(38):
		if lookahead == int32('i') {
			state = uint16(43)
			goto next_state
		}
		return result
	case int32(39):
		if lookahead == int32('i') {
			state = uint16(42)
			goto next_state
		}
		return result
	case int32(40):
		if lookahead == int32('i') {
			state = uint16(34)
			goto next_state
		}
		return result
	case int32(41):
		if lookahead == int32('l') {
			state = uint16(51)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(24)
			goto next_state
		}
		return result
	case int32(42):
		if lookahead == int32('m') {
			state = uint16(48)
			goto next_state
		}
		return result
	case int32(43):
		if lookahead == int32('n') {
			state = uint16(29)
			goto next_state
		}
		return result
	case int32(44):
		if lookahead == int32('o') {
			state = uint16(49)
			goto next_state
		}
		return result
	case int32(45):
		if lookahead == int32('o') {
			state = uint16(50)
			goto next_state
		}
		return result
	case int32(46):
		if lookahead == int32('p') {
			state = uint16(44)
			goto next_state
		}
		return result
	case int32(47):
		if lookahead == int32('p') {
			state = uint16(22)
			goto next_state
		}
		return result
	case int32(48):
		if lookahead == int32('p') {
			state = uint16(45)
			goto next_state
		}
		return result
	case int32(49):
		if lookahead == int32('r') {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(50):
		if lookahead == int32('r') {
			state = uint16(54)
			goto next_state
		}
		return result
	case int32(51):
		if lookahead == int32('s') {
			state = uint16(28)
			goto next_state
		}
		return result
	case int32(52):
		if lookahead == int32('t') {
			state = uint16(142)
			goto next_state
		}
		return result
	case int32(53):
		if lookahead == int32('t') {
			state = uint16(37)
			goto next_state
		}
		return result
	case int32(54):
		if lookahead == int32('t') {
			state = uint16(20)
			goto next_state
		}
		return result
	case int32(55):
		if lookahead == int32('+') || lookahead == int32('-') {
			state = uint16(59)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(105)
			goto next_state
		}
		return result
	case int32(56):
		if lookahead == int32('+') || lookahead == int32('-') {
			state = uint16(60)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(107)
			goto next_state
		}
		return result
	case int32(57):
		if lookahead == int32('X') || lookahead == int32('x') {
			state = uint16(62)
			goto next_state
		}
		return result
	case int32(58):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(59):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(105)
			goto next_state
		}
		return result
	case int32(60):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(107)
			goto next_state
		}
		return result
	case int32(61):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(108)
			goto next_state
		}
		return result
	case int32(62):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(102)
			goto next_state
		}
		return result
	case int32(63):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(78)
			goto next_state
		}
		return result
	case int32(64):
		if eof != 0 {
			state = uint16(65)
			goto next_state
		}
		i3 = uint32(0)
		for {
			if !(uint64(i3) < libc.Uint64FromInt64(80)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token3[i3]) == lookahead {
				state = map_token3[i3+uint32(1)]
				goto next_state
			}
			goto _4
		_4:
			;
			i3 = i3 + uint32(2)
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(97)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') || lookahead == int32(0xa0) || lookahead == int32(0x200b) || lookahead == int32(0x2060) || lookahead == int32(0xfeff) {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(64)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(77)
			goto next_state
		}
		return result
	case int32(65):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(ts_builtin_sym_end)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(66):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_line_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(66)
			goto next_state
		}
		return result
	case int32(67):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SEMI)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(68):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(69):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('=') {
			state = uint16(118)
			goto next_state
		}
		return result
	case int32(70):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(71):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COMMA)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(72):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(73):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DASH_GT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(74):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(75):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(76):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(77):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(78)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(78)
			goto next_state
		}
		return result
	case int32(78):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(78)
			goto next_state
		}
		return result
	case int32(79):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym__)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(78)
			goto next_state
		}
		return result
	case int32(80):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PLUS_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(81):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DASH_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(82):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STAR_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(83):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SLASH_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(84):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PERCENT_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(85):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AMP_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(86):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PIPE_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(87):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_CARET_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(88):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COLON)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(89):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COLON)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(':') {
			state = uint16(145)
			goto next_state
		}
		return result
	case int32(90):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('<') {
			state = uint16(122)
			goto next_state
		}
		if lookahead == int32('=') {
			state = uint16(120)
			goto next_state
		}
		return result
	case int32(91):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_GT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(92):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_GT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('=') {
			state = uint16(121)
			goto next_state
		}
		if lookahead == int32('>') {
			state = uint16(123)
			goto next_state
		}
		return result
	case int32(93):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PLUS_PLUS)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(94):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DASH_DASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(95):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_int_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(96):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_int_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i4 = uint32(0)
		for {
			if !(uint64(i4) < libc.Uint64FromInt64(32)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token4[i4]) == lookahead {
				state = map_token4[i4+uint32(1)]
				goto next_state
			}
			goto _5
		_5:
			;
			i4 = i4 + uint32(2)
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(6)
			goto next_state
		}
		return result
	case int32(97):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_int_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(104)
			goto next_state
		}
		if lookahead == int32('f') {
			state = uint16(103)
			goto next_state
		}
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('i') || lookahead == int32('u') {
			state = uint16(95)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(97)
			goto next_state
		}
		return result
	case int32(98):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_int_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(104)
			goto next_state
		}
		if lookahead == int32('f') {
			state = uint16(103)
			goto next_state
		}
		if lookahead == int32('i') || lookahead == int32('u') {
			state = uint16(95)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(98)
			goto next_state
		}
		return result
	case int32(99):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_int_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(108)
			goto next_state
		}
		if lookahead == int32('P') || lookahead == int32('p') {
			state = uint16(56)
			goto next_state
		}
		if lookahead == int32('i') || lookahead == int32('u') {
			state = uint16(95)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(99)
			goto next_state
		}
		return result
	case int32(100):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_int_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('X') || lookahead == int32('x') {
			state = uint16(62)
			goto next_state
		}
		if lookahead == int32('i') || lookahead == int32('u') {
			state = uint16(95)
			goto next_state
		}
		return result
	case int32(101):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_int_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('i') || lookahead == int32('u') {
			state = uint16(95)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(101)
			goto next_state
		}
		return result
	case int32(102):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_int_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('i') || lookahead == int32('u') {
			state = uint16(95)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(102)
			goto next_state
		}
		return result
	case int32(103):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_float_literal_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(104):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_float_literal_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('f') {
			state = uint16(103)
			goto next_state
		}
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(55)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(105):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_float_literal_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('f') {
			state = uint16(103)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(105)
			goto next_state
		}
		return result
	case int32(106):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_float_literal_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(107):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_float_literal_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('f') {
			state = uint16(106)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(107)
			goto next_state
		}
		return result
	case int32(108):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_float_literal_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('P') || lookahead == int32('p') {
			state = uint16(56)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(108)
			goto next_state
		}
		return result
	case int32(109):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PIPE_PIPE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(110):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AMP_AMP)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(111):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PIPE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('=') {
			state = uint16(86)
			goto next_state
		}
		if lookahead == int32('|') {
			state = uint16(109)
			goto next_state
		}
		return result
	case int32(112):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PIPE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('|') {
			state = uint16(109)
			goto next_state
		}
		return result
	case int32(113):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_CARET)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(114):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_CARET)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('=') {
			state = uint16(87)
			goto next_state
		}
		return result
	case int32(115):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AMP)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(116):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AMP)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('&') {
			state = uint16(110)
			goto next_state
		}
		return result
	case int32(117):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AMP)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('&') {
			state = uint16(110)
			goto next_state
		}
		if lookahead == int32('=') {
			state = uint16(85)
			goto next_state
		}
		return result
	case int32(118):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EQ_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(119):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BANG_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(120):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(121):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_GT_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(122):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT_LT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(123):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_GT_GT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(124):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PLUS)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(125):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PLUS)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('+') {
			state = uint16(93)
			goto next_state
		}
		if lookahead == int32('=') {
			state = uint16(80)
			goto next_state
		}
		return result
	case int32(126):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(127):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(58)
			goto next_state
		}
		if lookahead == int32('0') {
			state = uint16(7)
			goto next_state
		}
		if lookahead == int32('=') {
			state = uint16(81)
			goto next_state
		}
		if lookahead == int32('>') {
			state = uint16(73)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(98)
			goto next_state
		}
		return result
	case int32(128):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(58)
			goto next_state
		}
		if lookahead == int32('0') {
			state = uint16(7)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(98)
			goto next_state
		}
		return result
	case int32(129):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STAR)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(130):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STAR)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('=') {
			state = uint16(82)
			goto next_state
		}
		return result
	case int32(131):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SLASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('/') {
			state = uint16(66)
			goto next_state
		}
		return result
	case int32(132):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SLASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('/') {
			state = uint16(66)
			goto next_state
		}
		if lookahead == int32('=') {
			state = uint16(83)
			goto next_state
		}
		return result
	case int32(133):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PERCENT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(134):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PERCENT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('=') {
			state = uint16(84)
			goto next_state
		}
		return result
	case int32(135):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BANG)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(136):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BANG)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('=') {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(137):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_TILDE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(138):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(139):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(140):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(141):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(142):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_preproc_import_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(143):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LF)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(143)
			goto next_state
		}
		return result
	case int32(144):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_define_import_path_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(145):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COLON_COLON)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(146):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_preproc_ifdef_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(147):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_preproc_ifdef_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(148):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_preproc_ifdef_token3)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(149):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_preproc_else_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var map_token = [54]uint16_t{
	0:  uint16('!'),
	1:  uint16(136),
	2:  uint16('#'),
	3:  uint16(23),
	4:  uint16('%'),
	5:  uint16(134),
	6:  uint16('&'),
	7:  uint16(117),
	8:  uint16('('),
	9:  uint16(70),
	10: uint16(')'),
	11: uint16(72),
	12: uint16('*'),
	13: uint16(130),
	14: uint16('+'),
	15: uint16(125),
	16: uint16(','),
	17: uint16(71),
	18: uint16('-'),
	19: uint16(127),
	20: uint16('.'),
	21: uint16(141),
	22: uint16('/'),
	23: uint16(132),
	24: uint16('0'),
	25: uint16(96),
	26: uint16(':'),
	27: uint16(89),
	28: uint16(';'),
	29: uint16(67),
	30: uint16('<'),
	31: uint16(90),
	32: uint16('='),
	33: uint16(69),
	34: uint16('>'),
	35: uint16(92),
	36: uint16('@'),
	37: uint16(76),
	38: uint16('['),
	39: uint16(138),
	40: uint16(']'),
	41: uint16(139),
	42: uint16('^'),
	43: uint16(114),
	44: uint16('_'),
	45: uint16(79),
	46: uint16('{'),
	47: uint16(74),
	48: uint16('|'),
	49: uint16(111),
	50: uint16('}'),
	51: uint16(75),
	52: uint16('~'),
	53: uint16(137),
}

var map_token1 = [42]uint16_t{
	0:  uint16('!'),
	1:  uint16(13),
	2:  uint16('%'),
	3:  uint16(133),
	4:  uint16('&'),
	5:  uint16(116),
	6:  uint16('('),
	7:  uint16(70),
	8:  uint16(')'),
	9:  uint16(72),
	10: uint16('*'),
	11: uint16(129),
	12: uint16('+'),
	13: uint16(124),
	14: uint16(','),
	15: uint16(71),
	16: uint16('-'),
	17: uint16(126),
	18: uint16('.'),
	19: uint16(140),
	20: uint16('/'),
	21: uint16(131),
	22: uint16(':'),
	23: uint16(88),
	24: uint16(';'),
	25: uint16(67),
	26: uint16('<'),
	27: uint16(90),
	28: uint16('='),
	29: uint16(17),
	30: uint16('>'),
	31: uint16(92),
	32: uint16('['),
	33: uint16(138),
	34: uint16(']'),
	35: uint16(139),
	36: uint16('^'),
	37: uint16(113),
	38: uint16('{'),
	39: uint16(74),
	40: uint16('|'),
	41: uint16(112),
}

var map_token2 = [34]uint16_t{
	0:  uint16('%'),
	1:  uint16(14),
	2:  uint16('&'),
	3:  uint16(15),
	4:  uint16('('),
	5:  uint16(70),
	6:  uint16(')'),
	7:  uint16(72),
	8:  uint16('*'),
	9:  uint16(16),
	10: uint16('+'),
	11: uint16(4),
	12: uint16('-'),
	13: uint16(5),
	14: uint16('.'),
	15: uint16(140),
	16: uint16('/'),
	17: uint16(11),
	18: uint16('0'),
	19: uint16(100),
	20: uint16(':'),
	21: uint16(12),
	22: uint16('='),
	23: uint16(68),
	24: uint16('['),
	25: uint16(138),
	26: uint16('^'),
	27: uint16(18),
	28: uint16('_'),
	29: uint16(63),
	30: uint16('{'),
	31: uint16(74),
	32: uint16('|'),
	33: uint16(19),
}

var map_token3 = [40]uint16_t{
	0:  uint16('!'),
	1:  uint16(135),
	2:  uint16('#'),
	3:  uint16(23),
	4:  uint16('&'),
	5:  uint16(115),
	6:  uint16('('),
	7:  uint16(70),
	8:  uint16(')'),
	9:  uint16(72),
	10: uint16('*'),
	11: uint16(129),
	12: uint16(','),
	13: uint16(71),
	14: uint16('-'),
	15: uint16(128),
	16: uint16('.'),
	17: uint16(58),
	18: uint16('/'),
	19: uint16(10),
	20: uint16('0'),
	21: uint16(96),
	22: uint16(':'),
	23: uint16(88),
	24: uint16(';'),
	25: uint16(67),
	26: uint16('='),
	27: uint16(68),
	28: uint16('>'),
	29: uint16(91),
	30: uint16('@'),
	31: uint16(76),
	32: uint16('_'),
	33: uint16(79),
	34: uint16('{'),
	35: uint16(74),
	36: uint16('}'),
	37: uint16(75),
	38: uint16('~'),
	39: uint16(137),
}

var map_token4 = [16]uint16_t{
	0:  uint16('.'),
	1:  uint16(104),
	2:  uint16('f'),
	3:  uint16(103),
	4:  uint16('E'),
	5:  uint16(55),
	6:  uint16('e'),
	7:  uint16(55),
	8:  uint16('X'),
	9:  uint16(9),
	10: uint16('x'),
	11: uint16(9),
	12: uint16('i'),
	13: uint16(95),
	14: uint16('u'),
	15: uint16(95),
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
			if !(uint64(i) < libc.Uint64FromInt64(68)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token5[i]) == lookahead {
				state = map_token5[i+uint32(1)]
				goto next_state
			}
			goto _1
		_1:
			;
			i = i + uint32(2)
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') || lookahead == int32(0xa0) || lookahead == int32(0x200b) || lookahead == int32(0x2060) || lookahead == int32(0xfeff) {
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
		if lookahead == int32('s') {
			state = uint16(19)
			goto next_state
		}
		return result
	case int32(2):
		if lookahead == int32('i') {
			state = uint16(20)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(21)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(22)
			goto next_state
		}
		return result
	case int32(3):
		if lookahead == int32('a') {
			state = uint16(23)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(24)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead == int32('e') {
			state = uint16(25)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(26)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead == int32('l') {
			state = uint16(27)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(28)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead == int32('1') {
			state = uint16(29)
			goto next_state
		}
		if lookahead == int32('3') {
			state = uint16(30)
			goto next_state
		}
		if lookahead == int32('a') {
			state = uint16(31)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(32)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(33)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(34)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead == int32('3') {
			state = uint16(35)
			goto next_state
		}
		if lookahead == int32('f') {
			state = uint16(36)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead == int32('e') {
			state = uint16(37)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(38)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead == int32('a') {
			state = uint16(39)
			goto next_state
		}
		return result
	case int32(10):
		if lookahead == int32('v') {
			state = uint16(40)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead == int32('r') {
			state = uint16(41)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(42)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead == int32('3') {
			state = uint16(43)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(44)
			goto next_state
		}
		if lookahead == int32('g') {
			state = uint16(45)
			goto next_state
		}
		return result
	case int32(13):
		if lookahead == int32('a') {
			state = uint16(46)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(47)
			goto next_state
		}
		if lookahead == int32('w') {
			state = uint16(48)
			goto next_state
		}
		return result
	case int32(14):
		if lookahead == int32('e') {
			state = uint16(49)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(50)
			goto next_state
		}
		if lookahead == int32('y') {
			state = uint16(51)
			goto next_state
		}
		return result
	case int32(15):
		if lookahead == int32('3') {
			state = uint16(52)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(53)
			goto next_state
		}
		return result
	case int32(16):
		if lookahead == int32('a') {
			state = uint16(54)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(56)
			goto next_state
		}
		return result
	case int32(17):
		if lookahead == int32('h') {
			state = uint16(57)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(58)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(59)
			goto next_state
		}
		return result
	case int32(18):
		if lookahead == int32('r') {
			state = uint16(60)
			goto next_state
		}
		return result
	case int32(19):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_as)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(20):
		if lookahead == int32('t') {
			state = uint16(61)
			goto next_state
		}
		return result
	case int32(21):
		if lookahead == int32('o') {
			state = uint16(62)
			goto next_state
		}
		return result
	case int32(22):
		if lookahead == int32('e') {
			state = uint16(63)
			goto next_state
		}
		return result
	case int32(23):
		if lookahead == int32('s') {
			state = uint16(64)
			goto next_state
		}
		return result
	case int32(24):
		if lookahead == int32('n') {
			state = uint16(65)
			goto next_state
		}
		return result
	case int32(25):
		if lookahead == int32('f') {
			state = uint16(66)
			goto next_state
		}
		return result
	case int32(26):
		if lookahead == int32('s') {
			state = uint16(67)
			goto next_state
		}
		return result
	case int32(27):
		if lookahead == int32('s') {
			state = uint16(68)
			goto next_state
		}
		return result
	case int32(28):
		if lookahead == int32('a') {
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(29):
		if lookahead == int32('6') {
			state = uint16(70)
			goto next_state
		}
		return result
	case int32(30):
		if lookahead == int32('2') {
			state = uint16(71)
			goto next_state
		}
		return result
	case int32(31):
		if lookahead == int32('l') {
			state = uint16(72)
			goto next_state
		}
		return result
	case int32(32):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_fn)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(33):
		if lookahead == int32('r') {
			state = uint16(73)
			goto next_state
		}
		return result
	case int32(34):
		if lookahead == int32('n') {
			state = uint16(74)
			goto next_state
		}
		return result
	case int32(35):
		if lookahead == int32('2') {
			state = uint16(75)
			goto next_state
		}
		return result
	case int32(36):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_if)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(37):
		if lookahead == int32('t') {
			state = uint16(76)
			goto next_state
		}
		return result
	case int32(38):
		if lookahead == int32('o') {
			state = uint16(77)
			goto next_state
		}
		return result
	case int32(39):
		if lookahead == int32('t') {
			state = uint16(78)
			goto next_state
		}
		return result
	case int32(40):
		if lookahead == int32('e') {
			state = uint16(79)
			goto next_state
		}
		return result
	case int32(41):
		if lookahead == int32('i') {
			state = uint16(80)
			goto next_state
		}
		return result
	case int32(42):
		if lookahead == int32('r') {
			state = uint16(81)
			goto next_state
		}
		return result
	case int32(43):
		if lookahead == int32('2') {
			state = uint16(82)
			goto next_state
		}
		return result
	case int32(44):
		if lookahead == int32('a') {
			state = uint16(83)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(84)
			goto next_state
		}
		return result
	case int32(45):
		if lookahead == int32('3') {
			state = uint16(85)
			goto next_state
		}
		if lookahead == int32('b') {
			state = uint16(86)
			goto next_state
		}
		return result
	case int32(46):
		if lookahead == int32('m') {
			state = uint16(87)
			goto next_state
		}
		return result
	case int32(47):
		if lookahead == int32('o') {
			state = uint16(88)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(89)
			goto next_state
		}
		return result
	case int32(48):
		if lookahead == int32('i') {
			state = uint16(90)
			goto next_state
		}
		return result
	case int32(49):
		if lookahead == int32('x') {
			state = uint16(91)
			goto next_state
		}
		return result
	case int32(50):
		if lookahead == int32('u') {
			state = uint16(92)
			goto next_state
		}
		return result
	case int32(51):
		if lookahead == int32('p') {
			state = uint16(93)
			goto next_state
		}
		return result
	case int32(52):
		if lookahead == int32('2') {
			state = uint16(94)
			goto next_state
		}
		return result
	case int32(53):
		if lookahead == int32('i') {
			state = uint16(95)
			goto next_state
		}
		return result
	case int32(54):
		if lookahead == int32('r') {
			state = uint16(96)
			goto next_state
		}
		return result
	case int32(55):
		if lookahead == int32('c') {
			state = uint16(97)
			goto next_state
		}
		return result
	case int32(56):
		if lookahead == int32('r') {
			state = uint16(98)
			goto next_state
		}
		return result
	case int32(57):
		if lookahead == int32('i') {
			state = uint16(99)
			goto next_state
		}
		return result
	case int32(58):
		if lookahead == int32('r') {
			state = uint16(100)
			goto next_state
		}
		return result
	case int32(59):
		if lookahead == int32('i') {
			state = uint16(101)
			goto next_state
		}
		return result
	case int32(60):
		if lookahead == int32('a') {
			state = uint16(102)
			goto next_state
		}
		return result
	case int32(61):
		if lookahead == int32('c') {
			state = uint16(103)
			goto next_state
		}
		return result
	case int32(62):
		if lookahead == int32('l') {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(63):
		if lookahead == int32('a') {
			state = uint16(105)
			goto next_state
		}
		return result
	case int32(64):
		if lookahead == int32('e') {
			state = uint16(106)
			goto next_state
		}
		return result
	case int32(65):
		if lookahead == int32('t') {
			state = uint16(107)
			goto next_state
		}
		return result
	case int32(66):
		if lookahead == int32('a') {
			state = uint16(108)
			goto next_state
		}
		return result
	case int32(67):
		if lookahead == int32('c') {
			state = uint16(109)
			goto next_state
		}
		return result
	case int32(68):
		if lookahead == int32('e') {
			state = uint16(110)
			goto next_state
		}
		return result
	case int32(69):
		if lookahead == int32('b') {
			state = uint16(111)
			goto next_state
		}
		return result
	case int32(70):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_f16)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(71):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_f32)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(72):
		if lookahead == int32('l') {
			state = uint16(112)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(113)
			goto next_state
		}
		return result
	case int32(73):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_for)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(74):
		if lookahead == int32('c') {
			state = uint16(114)
			goto next_state
		}
		return result
	case int32(75):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_i32)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(76):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_let)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(77):
		if lookahead == int32('p') {
			state = uint16(115)
			goto next_state
		}
		return result
	case int32(78):
		if lookahead == int32('2') {
			state = uint16(116)
			goto next_state
		}
		if lookahead == int32('3') {
			state = uint16(117)
			goto next_state
		}
		if lookahead == int32('4') {
			state = uint16(118)
			goto next_state
		}
		return result
	case int32(79):
		if lookahead == int32('r') {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(80):
		if lookahead == int32('v') {
			state = uint16(120)
			goto next_state
		}
		return result
	case int32(81):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ptr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(82):
		if lookahead == int32('f') {
			state = uint16(121)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(122)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(123)
			goto next_state
		}
		return result
	case int32(83):
		if lookahead == int32('d') {
			state = uint16(124)
			goto next_state
		}
		return result
	case int32(84):
		if lookahead == int32('u') {
			state = uint16(125)
			goto next_state
		}
		return result
	case int32(85):
		if lookahead == int32('2') {
			state = uint16(126)
			goto next_state
		}
		return result
	case int32(86):
		if lookahead == int32('a') {
			state = uint16(127)
			goto next_state
		}
		return result
	case int32(87):
		if lookahead == int32('p') {
			state = uint16(128)
			goto next_state
		}
		return result
	case int32(88):
		if lookahead == int32('r') {
			state = uint16(129)
			goto next_state
		}
		return result
	case int32(89):
		if lookahead == int32('u') {
			state = uint16(130)
			goto next_state
		}
		return result
	case int32(90):
		if lookahead == int32('t') {
			state = uint16(131)
			goto next_state
		}
		return result
	case int32(91):
		if lookahead == int32('t') {
			state = uint16(132)
			goto next_state
		}
		return result
	case int32(92):
		if lookahead == int32('e') {
			state = uint16(133)
			goto next_state
		}
		return result
	case int32(93):
		if lookahead == int32('e') {
			state = uint16(134)
			goto next_state
		}
		return result
	case int32(94):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_u32)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(95):
		if lookahead == int32('f') {
			state = uint16(135)
			goto next_state
		}
		return result
	case int32(96):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_var)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(97):
		if lookahead == int32('2') {
			state = uint16(136)
			goto next_state
		}
		if lookahead == int32('3') {
			state = uint16(137)
			goto next_state
		}
		if lookahead == int32('4') {
			state = uint16(138)
			goto next_state
		}
		return result
	case int32(98):
		if lookahead == int32('t') {
			state = uint16(139)
			goto next_state
		}
		return result
	case int32(99):
		if lookahead == int32('l') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(100):
		if lookahead == int32('k') {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(101):
		if lookahead == int32('t') {
			state = uint16(142)
			goto next_state
		}
		return result
	case int32(102):
		if lookahead == int32('y') {
			state = uint16(143)
			goto next_state
		}
		return result
	case int32(103):
		if lookahead == int32('a') {
			state = uint16(144)
			goto next_state
		}
		return result
	case int32(104):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_bool)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(105):
		if lookahead == int32('k') {
			state = uint16(145)
			goto next_state
		}
		return result
	case int32(106):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_case)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(107):
		if lookahead == int32('i') {
			state = uint16(146)
			goto next_state
		}
		return result
	case int32(108):
		if lookahead == int32('u') {
			state = uint16(147)
			goto next_state
		}
		return result
	case int32(109):
		if lookahead == int32('a') {
			state = uint16(148)
			goto next_state
		}
		return result
	case int32(110):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_else)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(111):
		if lookahead == int32('l') {
			state = uint16(149)
			goto next_state
		}
		return result
	case int32(112):
		if lookahead == int32('t') {
			state = uint16(150)
			goto next_state
		}
		return result
	case int32(113):
		if lookahead == int32('e') {
			state = uint16(151)
			goto next_state
		}
		return result
	case int32(114):
		if lookahead == int32('t') {
			state = uint16(152)
			goto next_state
		}
		return result
	case int32(115):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_loop)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(116):
		if lookahead == int32('x') {
			state = uint16(153)
			goto next_state
		}
		return result
	case int32(117):
		if lookahead == int32('x') {
			state = uint16(154)
			goto next_state
		}
		return result
	case int32(118):
		if lookahead == int32('x') {
			state = uint16(155)
			goto next_state
		}
		return result
	case int32(119):
		if lookahead == int32('r') {
			state = uint16(156)
			goto next_state
		}
		return result
	case int32(120):
		if lookahead == int32('a') {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(121):
		if lookahead == int32('l') {
			state = uint16(158)
			goto next_state
		}
		return result
	case int32(122):
		if lookahead == int32('i') {
			state = uint16(159)
			goto next_state
		}
		return result
	case int32(123):
		if lookahead == int32('i') {
			state = uint16(160)
			goto next_state
		}
		return result
	case int32(124):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_read)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(161)
			goto next_state
		}
		return result
	case int32(125):
		if lookahead == int32('r') {
			state = uint16(162)
			goto next_state
		}
		return result
	case int32(126):
		if lookahead == int32('f') {
			state = uint16(163)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(164)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(165)
			goto next_state
		}
		return result
	case int32(127):
		if lookahead == int32('1') {
			state = uint16(166)
			goto next_state
		}
		if lookahead == int32('3') {
			state = uint16(167)
			goto next_state
		}
		if lookahead == int32('8') {
			state = uint16(168)
			goto next_state
		}
		return result
	case int32(128):
		if lookahead == int32('l') {
			state = uint16(169)
			goto next_state
		}
		return result
	case int32(129):
		if lookahead == int32('a') {
			state = uint16(170)
			goto next_state
		}
		return result
	case int32(130):
		if lookahead == int32('c') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(131):
		if lookahead == int32('c') {
			state = uint16(172)
			goto next_state
		}
		return result
	case int32(132):
		if lookahead == int32('u') {
			state = uint16(173)
			goto next_state
		}
		return result
	case int32(133):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_true)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(134):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_type)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(135):
		if lookahead == int32('o') {
			state = uint16(174)
			goto next_state
		}
		return result
	case int32(136):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_vec2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(137):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_vec3)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(138):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_vec4)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(139):
		if lookahead == int32('u') {
			state = uint16(175)
			goto next_state
		}
		return result
	case int32(140):
		if lookahead == int32('e') {
			state = uint16(176)
			goto next_state
		}
		return result
	case int32(141):
		if lookahead == int32('g') {
			state = uint16(177)
			goto next_state
		}
		return result
	case int32(142):
		if lookahead == int32('e') {
			state = uint16(178)
			goto next_state
		}
		return result
	case int32(143):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_array)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(144):
		if lookahead == int32('s') {
			state = uint16(179)
			goto next_state
		}
		return result
	case int32(145):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_break)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(146):
		if lookahead == int32('n') {
			state = uint16(180)
			goto next_state
		}
		return result
	case int32(147):
		if lookahead == int32('l') {
			state = uint16(181)
			goto next_state
		}
		return result
	case int32(148):
		if lookahead == int32('r') {
			state = uint16(182)
			goto next_state
		}
		return result
	case int32(149):
		if lookahead == int32('e') {
			state = uint16(183)
			goto next_state
		}
		return result
	case int32(150):
		if lookahead == int32('h') {
			state = uint16(184)
			goto next_state
		}
		return result
	case int32(151):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_false)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(152):
		if lookahead == int32('i') {
			state = uint16(185)
			goto next_state
		}
		return result
	case int32(153):
		if lookahead == int32('2') {
			state = uint16(186)
			goto next_state
		}
		if lookahead == int32('3') {
			state = uint16(187)
			goto next_state
		}
		if lookahead == int32('4') {
			state = uint16(188)
			goto next_state
		}
		return result
	case int32(154):
		if lookahead == int32('2') {
			state = uint16(189)
			goto next_state
		}
		if lookahead == int32('3') {
			state = uint16(190)
			goto next_state
		}
		if lookahead == int32('4') {
			state = uint16(191)
			goto next_state
		}
		return result
	case int32(155):
		if lookahead == int32('2') {
			state = uint16(192)
			goto next_state
		}
		if lookahead == int32('3') {
			state = uint16(193)
			goto next_state
		}
		if lookahead == int32('4') {
			state = uint16(194)
			goto next_state
		}
		return result
	case int32(156):
		if lookahead == int32('i') {
			state = uint16(195)
			goto next_state
		}
		return result
	case int32(157):
		if lookahead == int32('t') {
			state = uint16(196)
			goto next_state
		}
		return result
	case int32(158):
		if lookahead == int32('o') {
			state = uint16(197)
			goto next_state
		}
		return result
	case int32(159):
		if lookahead == int32('n') {
			state = uint16(198)
			goto next_state
		}
		return result
	case int32(160):
		if lookahead == int32('n') {
			state = uint16(199)
			goto next_state
		}
		return result
	case int32(161):
		if lookahead == int32('w') {
			state = uint16(200)
			goto next_state
		}
		return result
	case int32(162):
		if lookahead == int32('n') {
			state = uint16(201)
			goto next_state
		}
		return result
	case int32(163):
		if lookahead == int32('l') {
			state = uint16(202)
			goto next_state
		}
		return result
	case int32(164):
		if lookahead == int32('i') {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(165):
		if lookahead == int32('i') {
			state = uint16(204)
			goto next_state
		}
		return result
	case int32(166):
		if lookahead == int32('6') {
			state = uint16(205)
			goto next_state
		}
		return result
	case int32(167):
		if lookahead == int32('2') {
			state = uint16(206)
			goto next_state
		}
		return result
	case int32(168):
		if lookahead == int32('s') {
			state = uint16(207)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(208)
			goto next_state
		}
		return result
	case int32(169):
		if lookahead == int32('e') {
			state = uint16(209)
			goto next_state
		}
		return result
	case int32(170):
		if lookahead == int32('g') {
			state = uint16(210)
			goto next_state
		}
		return result
	case int32(171):
		if lookahead == int32('t') {
			state = uint16(211)
			goto next_state
		}
		return result
	case int32(172):
		if lookahead == int32('h') {
			state = uint16(212)
			goto next_state
		}
		return result
	case int32(173):
		if lookahead == int32('r') {
			state = uint16(213)
			goto next_state
		}
		return result
	case int32(174):
		if lookahead == int32('r') {
			state = uint16(214)
			goto next_state
		}
		return result
	case int32(175):
		if lookahead == int32('a') {
			state = uint16(215)
			goto next_state
		}
		return result
	case int32(176):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_while)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(177):
		if lookahead == int32('r') {
			state = uint16(216)
			goto next_state
		}
		return result
	case int32(178):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_write)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(179):
		if lookahead == int32('t') {
			state = uint16(217)
			goto next_state
		}
		return result
	case int32(180):
		if lookahead == int32('u') {
			state = uint16(218)
			goto next_state
		}
		return result
	case int32(181):
		if lookahead == int32('t') {
			state = uint16(219)
			goto next_state
		}
		return result
	case int32(182):
		if lookahead == int32('d') {
			state = uint16(220)
			goto next_state
		}
		return result
	case int32(183):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_enable)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(184):
		if lookahead == int32('r') {
			state = uint16(221)
			goto next_state
		}
		return result
	case int32(185):
		if lookahead == int32('o') {
			state = uint16(222)
			goto next_state
		}
		return result
	case int32(186):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_mat2x2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(187):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_mat2x3)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(188):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_mat2x4)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(189):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_mat3x2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(190):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_mat3x3)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(191):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_mat3x4)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(192):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_mat4x2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(193):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_mat4x3)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(194):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_mat4x4)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(195):
		if lookahead == int32('d') {
			state = uint16(223)
			goto next_state
		}
		return result
	case int32(196):
		if lookahead == int32('e') {
			state = uint16(224)
			goto next_state
		}
		return result
	case int32(197):
		if lookahead == int32('a') {
			state = uint16(225)
			goto next_state
		}
		return result
	case int32(198):
		if lookahead == int32('t') {
			state = uint16(226)
			goto next_state
		}
		return result
	case int32(199):
		if lookahead == int32('t') {
			state = uint16(227)
			goto next_state
		}
		return result
	case int32(200):
		if lookahead == int32('r') {
			state = uint16(228)
			goto next_state
		}
		return result
	case int32(201):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_return)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(202):
		if lookahead == int32('o') {
			state = uint16(229)
			goto next_state
		}
		return result
	case int32(203):
		if lookahead == int32('n') {
			state = uint16(230)
			goto next_state
		}
		return result
	case int32(204):
		if lookahead == int32('n') {
			state = uint16(231)
			goto next_state
		}
		return result
	case int32(205):
		if lookahead == int32('f') {
			state = uint16(232)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(233)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(234)
			goto next_state
		}
		return result
	case int32(206):
		if lookahead == int32('f') {
			state = uint16(235)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(236)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(237)
			goto next_state
		}
		return result
	case int32(207):
		if lookahead == int32('i') {
			state = uint16(238)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(239)
			goto next_state
		}
		return result
	case int32(208):
		if lookahead == int32('i') {
			state = uint16(240)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(241)
			goto next_state
		}
		return result
	case int32(209):
		if lookahead == int32('r') {
			state = uint16(242)
			goto next_state
		}
		return result
	case int32(210):
		if lookahead == int32('e') {
			state = uint16(243)
			goto next_state
		}
		return result
	case int32(211):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_struct)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(212):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_switch)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(213):
		if lookahead == int32('e') {
			state = uint16(244)
			goto next_state
		}
		return result
	case int32(214):
		if lookahead == int32('m') {
			state = uint16(245)
			goto next_state
		}
		return result
	case int32(215):
		if lookahead == int32('l') {
			state = uint16(246)
			goto next_state
		}
		return result
	case int32(216):
		if lookahead == int32('o') {
			state = uint16(247)
			goto next_state
		}
		return result
	case int32(217):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_bitcast)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(218):
		if lookahead == int32('e') {
			state = uint16(248)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(249)
			goto next_state
		}
		return result
	case int32(219):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_default)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(220):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_discard)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(221):
		if lookahead == int32('o') {
			state = uint16(250)
			goto next_state
		}
		return result
	case int32(222):
		if lookahead == int32('n') {
			state = uint16(251)
			goto next_state
		}
		return result
	case int32(223):
		if lookahead == int32('e') {
			state = uint16(252)
			goto next_state
		}
		return result
	case int32(224):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_private)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(225):
		if lookahead == int32('t') {
			state = uint16(253)
			goto next_state
		}
		return result
	case int32(226):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_r32sint)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(227):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_r32uint)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(228):
		if lookahead == int32('i') {
			state = uint16(254)
			goto next_state
		}
		return result
	case int32(229):
		if lookahead == int32('a') {
			state = uint16(255)
			goto next_state
		}
		return result
	case int32(230):
		if lookahead == int32('t') {
			state = uint16(256)
			goto next_state
		}
		return result
	case int32(231):
		if lookahead == int32('t') {
			state = uint16(257)
			goto next_state
		}
		return result
	case int32(232):
		if lookahead == int32('l') {
			state = uint16(258)
			goto next_state
		}
		return result
	case int32(233):
		if lookahead == int32('i') {
			state = uint16(259)
			goto next_state
		}
		return result
	case int32(234):
		if lookahead == int32('i') {
			state = uint16(260)
			goto next_state
		}
		return result
	case int32(235):
		if lookahead == int32('l') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(236):
		if lookahead == int32('i') {
			state = uint16(262)
			goto next_state
		}
		return result
	case int32(237):
		if lookahead == int32('i') {
			state = uint16(263)
			goto next_state
		}
		return result
	case int32(238):
		if lookahead == int32('n') {
			state = uint16(264)
			goto next_state
		}
		return result
	case int32(239):
		if lookahead == int32('o') {
			state = uint16(265)
			goto next_state
		}
		return result
	case int32(240):
		if lookahead == int32('n') {
			state = uint16(266)
			goto next_state
		}
		return result
	case int32(241):
		if lookahead == int32('o') {
			state = uint16(267)
			goto next_state
		}
		return result
	case int32(242):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_sampler)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(268)
			goto next_state
		}
		return result
	case int32(243):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_storage)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(244):
		if lookahead == int32('_') {
			state = uint16(269)
			goto next_state
		}
		return result
	case int32(245):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_uniform)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(246):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_virtual)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(247):
		if lookahead == int32('u') {
			state = uint16(270)
			goto next_state
		}
		return result
	case int32(248):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_continue)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(249):
		if lookahead == int32('n') {
			state = uint16(271)
			goto next_state
		}
		return result
	case int32(250):
		if lookahead == int32('u') {
			state = uint16(272)
			goto next_state
		}
		return result
	case int32(251):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_function)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(252):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_override)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(253):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_r32float)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(254):
		if lookahead == int32('t') {
			state = uint16(273)
			goto next_state
		}
		return result
	case int32(255):
		if lookahead == int32('t') {
			state = uint16(274)
			goto next_state
		}
		return result
	case int32(256):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_rg32sint)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(257):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_rg32uint)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(258):
		if lookahead == int32('o') {
			state = uint16(275)
			goto next_state
		}
		return result
	case int32(259):
		if lookahead == int32('n') {
			state = uint16(276)
			goto next_state
		}
		return result
	case int32(260):
		if lookahead == int32('n') {
			state = uint16(277)
			goto next_state
		}
		return result
	case int32(261):
		if lookahead == int32('o') {
			state = uint16(278)
			goto next_state
		}
		return result
	case int32(262):
		if lookahead == int32('n') {
			state = uint16(279)
			goto next_state
		}
		return result
	case int32(263):
		if lookahead == int32('n') {
			state = uint16(280)
			goto next_state
		}
		return result
	case int32(264):
		if lookahead == int32('t') {
			state = uint16(281)
			goto next_state
		}
		return result
	case int32(265):
		if lookahead == int32('r') {
			state = uint16(282)
			goto next_state
		}
		return result
	case int32(266):
		if lookahead == int32('t') {
			state = uint16(283)
			goto next_state
		}
		return result
	case int32(267):
		if lookahead == int32('r') {
			state = uint16(284)
			goto next_state
		}
		return result
	case int32(268):
		if lookahead == int32('c') {
			state = uint16(285)
			goto next_state
		}
		return result
	case int32(269):
		if lookahead == int32('1') {
			state = uint16(286)
			goto next_state
		}
		if lookahead == int32('2') {
			state = uint16(287)
			goto next_state
		}
		if lookahead == int32('3') {
			state = uint16(288)
			goto next_state
		}
		if lookahead == int32('c') {
			state = uint16(289)
			goto next_state
		}
		if lookahead == int32('d') {
			state = uint16(290)
			goto next_state
		}
		if lookahead == int32('m') {
			state = uint16(291)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(292)
			goto next_state
		}
		return result
	case int32(270):
		if lookahead == int32('p') {
			state = uint16(293)
			goto next_state
		}
		return result
	case int32(271):
		if lookahead == int32('g') {
			state = uint16(294)
			goto next_state
		}
		return result
	case int32(272):
		if lookahead == int32('g') {
			state = uint16(295)
			goto next_state
		}
		return result
	case int32(273):
		if lookahead == int32('e') {
			state = uint16(296)
			goto next_state
		}
		return result
	case int32(274):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_rg32float)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(275):
		if lookahead == int32('a') {
			state = uint16(297)
			goto next_state
		}
		return result
	case int32(276):
		if lookahead == int32('t') {
			state = uint16(298)
			goto next_state
		}
		return result
	case int32(277):
		if lookahead == int32('t') {
			state = uint16(299)
			goto next_state
		}
		return result
	case int32(278):
		if lookahead == int32('a') {
			state = uint16(300)
			goto next_state
		}
		return result
	case int32(279):
		if lookahead == int32('t') {
			state = uint16(301)
			goto next_state
		}
		return result
	case int32(280):
		if lookahead == int32('t') {
			state = uint16(302)
			goto next_state
		}
		return result
	case int32(281):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_rgba8sint)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(282):
		if lookahead == int32('m') {
			state = uint16(303)
			goto next_state
		}
		return result
	case int32(283):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_rgba8uint)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(284):
		if lookahead == int32('m') {
			state = uint16(304)
			goto next_state
		}
		return result
	case int32(285):
		if lookahead == int32('o') {
			state = uint16(305)
			goto next_state
		}
		return result
	case int32(286):
		if lookahead == int32('d') {
			state = uint16(306)
			goto next_state
		}
		return result
	case int32(287):
		if lookahead == int32('d') {
			state = uint16(307)
			goto next_state
		}
		return result
	case int32(288):
		if lookahead == int32('d') {
			state = uint16(308)
			goto next_state
		}
		return result
	case int32(289):
		if lookahead == int32('u') {
			state = uint16(309)
			goto next_state
		}
		return result
	case int32(290):
		if lookahead == int32('e') {
			state = uint16(310)
			goto next_state
		}
		return result
	case int32(291):
		if lookahead == int32('u') {
			state = uint16(311)
			goto next_state
		}
		return result
	case int32(292):
		if lookahead == int32('t') {
			state = uint16(312)
			goto next_state
		}
		return result
	case int32(293):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_workgroup)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(294):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_continuing)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(295):
		if lookahead == int32('h') {
			state = uint16(313)
			goto next_state
		}
		return result
	case int32(296):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_read_write)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(297):
		if lookahead == int32('t') {
			state = uint16(314)
			goto next_state
		}
		return result
	case int32(298):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_rgba16sint)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(299):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_rgba16uint)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(300):
		if lookahead == int32('t') {
			state = uint16(315)
			goto next_state
		}
		return result
	case int32(301):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_rgba32sint)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(302):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_rgba32uint)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(303):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_rgba8snorm)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(304):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_rgba8unorm)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(305):
		if lookahead == int32('m') {
			state = uint16(316)
			goto next_state
		}
		return result
	case int32(306):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_texture_1d)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(307):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_texture_2d)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(317)
			goto next_state
		}
		return result
	case int32(308):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_texture_3d)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(309):
		if lookahead == int32('b') {
			state = uint16(318)
			goto next_state
		}
		return result
	case int32(310):
		if lookahead == int32('p') {
			state = uint16(319)
			goto next_state
		}
		return result
	case int32(311):
		if lookahead == int32('l') {
			state = uint16(320)
			goto next_state
		}
		return result
	case int32(312):
		if lookahead == int32('o') {
			state = uint16(321)
			goto next_state
		}
		return result
	case int32(313):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_fallthrough)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(314):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_rgba16float)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(315):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_rgba32float)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(316):
		if lookahead == int32('p') {
			state = uint16(322)
			goto next_state
		}
		return result
	case int32(317):
		if lookahead == int32('a') {
			state = uint16(323)
			goto next_state
		}
		return result
	case int32(318):
		if lookahead == int32('e') {
			state = uint16(324)
			goto next_state
		}
		return result
	case int32(319):
		if lookahead == int32('t') {
			state = uint16(325)
			goto next_state
		}
		return result
	case int32(320):
		if lookahead == int32('t') {
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
		if lookahead == int32('a') {
			state = uint16(328)
			goto next_state
		}
		return result
	case int32(323):
		if lookahead == int32('r') {
			state = uint16(329)
			goto next_state
		}
		return result
	case int32(324):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_texture_cube)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(330)
			goto next_state
		}
		return result
	case int32(325):
		if lookahead == int32('h') {
			state = uint16(331)
			goto next_state
		}
		return result
	case int32(326):
		if lookahead == int32('i') {
			state = uint16(332)
			goto next_state
		}
		return result
	case int32(327):
		if lookahead == int32('a') {
			state = uint16(333)
			goto next_state
		}
		return result
	case int32(328):
		if lookahead == int32('r') {
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
		if lookahead == int32('a') {
			state = uint16(336)
			goto next_state
		}
		return result
	case int32(331):
		if lookahead == int32('_') {
			state = uint16(337)
			goto next_state
		}
		return result
	case int32(332):
		if lookahead == int32('s') {
			state = uint16(338)
			goto next_state
		}
		return result
	case int32(333):
		if lookahead == int32('g') {
			state = uint16(339)
			goto next_state
		}
		return result
	case int32(334):
		if lookahead == int32('i') {
			state = uint16(340)
			goto next_state
		}
		return result
	case int32(335):
		if lookahead == int32('a') {
			state = uint16(341)
			goto next_state
		}
		return result
	case int32(336):
		if lookahead == int32('r') {
			state = uint16(342)
			goto next_state
		}
		return result
	case int32(337):
		if lookahead == int32('2') {
			state = uint16(343)
			goto next_state
		}
		if lookahead == int32('c') {
			state = uint16(344)
			goto next_state
		}
		if lookahead == int32('m') {
			state = uint16(345)
			goto next_state
		}
		return result
	case int32(338):
		if lookahead == int32('a') {
			state = uint16(346)
			goto next_state
		}
		return result
	case int32(339):
		if lookahead == int32('e') {
			state = uint16(347)
			goto next_state
		}
		return result
	case int32(340):
		if lookahead == int32('s') {
			state = uint16(348)
			goto next_state
		}
		return result
	case int32(341):
		if lookahead == int32('y') {
			state = uint16(349)
			goto next_state
		}
		return result
	case int32(342):
		if lookahead == int32('r') {
			state = uint16(350)
			goto next_state
		}
		return result
	case int32(343):
		if lookahead == int32('d') {
			state = uint16(351)
			goto next_state
		}
		return result
	case int32(344):
		if lookahead == int32('u') {
			state = uint16(352)
			goto next_state
		}
		return result
	case int32(345):
		if lookahead == int32('u') {
			state = uint16(353)
			goto next_state
		}
		return result
	case int32(346):
		if lookahead == int32('m') {
			state = uint16(354)
			goto next_state
		}
		return result
	case int32(347):
		if lookahead == int32('_') {
			state = uint16(355)
			goto next_state
		}
		return result
	case int32(348):
		if lookahead == int32('o') {
			state = uint16(356)
			goto next_state
		}
		return result
	case int32(349):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_texture_2d_array)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(350):
		if lookahead == int32('a') {
			state = uint16(357)
			goto next_state
		}
		return result
	case int32(351):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_texture_depth_2d)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(358)
			goto next_state
		}
		return result
	case int32(352):
		if lookahead == int32('b') {
			state = uint16(359)
			goto next_state
		}
		return result
	case int32(353):
		if lookahead == int32('l') {
			state = uint16(360)
			goto next_state
		}
		return result
	case int32(354):
		if lookahead == int32('p') {
			state = uint16(361)
			goto next_state
		}
		return result
	case int32(355):
		if lookahead == int32('1') {
			state = uint16(362)
			goto next_state
		}
		if lookahead == int32('2') {
			state = uint16(363)
			goto next_state
		}
		if lookahead == int32('3') {
			state = uint16(364)
			goto next_state
		}
		return result
	case int32(356):
		if lookahead == int32('n') {
			state = uint16(365)
			goto next_state
		}
		return result
	case int32(357):
		if lookahead == int32('y') {
			state = uint16(366)
			goto next_state
		}
		return result
	case int32(358):
		if lookahead == int32('a') {
			state = uint16(367)
			goto next_state
		}
		return result
	case int32(359):
		if lookahead == int32('e') {
			state = uint16(368)
			goto next_state
		}
		return result
	case int32(360):
		if lookahead == int32('t') {
			state = uint16(369)
			goto next_state
		}
		return result
	case int32(361):
		if lookahead == int32('l') {
			state = uint16(370)
			goto next_state
		}
		return result
	case int32(362):
		if lookahead == int32('d') {
			state = uint16(371)
			goto next_state
		}
		return result
	case int32(363):
		if lookahead == int32('d') {
			state = uint16(372)
			goto next_state
		}
		return result
	case int32(364):
		if lookahead == int32('d') {
			state = uint16(373)
			goto next_state
		}
		return result
	case int32(365):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_sampler_comparison)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(366):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_texture_cube_array)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(367):
		if lookahead == int32('r') {
			state = uint16(374)
			goto next_state
		}
		return result
	case int32(368):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_texture_depth_cube)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(375)
			goto next_state
		}
		return result
	case int32(369):
		if lookahead == int32('i') {
			state = uint16(376)
			goto next_state
		}
		return result
	case int32(370):
		if lookahead == int32('e') {
			state = uint16(377)
			goto next_state
		}
		return result
	case int32(371):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_texture_storage_1d)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(372):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_texture_storage_2d)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(378)
			goto next_state
		}
		return result
	case int32(373):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_texture_storage_3d)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(374):
		if lookahead == int32('r') {
			state = uint16(379)
			goto next_state
		}
		return result
	case int32(375):
		if lookahead == int32('a') {
			state = uint16(380)
			goto next_state
		}
		return result
	case int32(376):
		if lookahead == int32('s') {
			state = uint16(381)
			goto next_state
		}
		return result
	case int32(377):
		if lookahead == int32('d') {
			state = uint16(382)
			goto next_state
		}
		return result
	case int32(378):
		if lookahead == int32('a') {
			state = uint16(383)
			goto next_state
		}
		return result
	case int32(379):
		if lookahead == int32('a') {
			state = uint16(384)
			goto next_state
		}
		return result
	case int32(380):
		if lookahead == int32('r') {
			state = uint16(385)
			goto next_state
		}
		return result
	case int32(381):
		if lookahead == int32('a') {
			state = uint16(386)
			goto next_state
		}
		return result
	case int32(382):
		if lookahead == int32('_') {
			state = uint16(387)
			goto next_state
		}
		return result
	case int32(383):
		if lookahead == int32('r') {
			state = uint16(388)
			goto next_state
		}
		return result
	case int32(384):
		if lookahead == int32('y') {
			state = uint16(389)
			goto next_state
		}
		return result
	case int32(385):
		if lookahead == int32('r') {
			state = uint16(390)
			goto next_state
		}
		return result
	case int32(386):
		if lookahead == int32('m') {
			state = uint16(391)
			goto next_state
		}
		return result
	case int32(387):
		if lookahead == int32('2') {
			state = uint16(392)
			goto next_state
		}
		return result
	case int32(388):
		if lookahead == int32('r') {
			state = uint16(393)
			goto next_state
		}
		return result
	case int32(389):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_texture_depth_2d_array)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(390):
		if lookahead == int32('a') {
			state = uint16(394)
			goto next_state
		}
		return result
	case int32(391):
		if lookahead == int32('p') {
			state = uint16(395)
			goto next_state
		}
		return result
	case int32(392):
		if lookahead == int32('d') {
			state = uint16(396)
			goto next_state
		}
		return result
	case int32(393):
		if lookahead == int32('a') {
			state = uint16(397)
			goto next_state
		}
		return result
	case int32(394):
		if lookahead == int32('y') {
			state = uint16(398)
			goto next_state
		}
		return result
	case int32(395):
		if lookahead == int32('l') {
			state = uint16(399)
			goto next_state
		}
		return result
	case int32(396):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_texture_multisampled_2d)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(397):
		if lookahead == int32('y') {
			state = uint16(400)
			goto next_state
		}
		return result
	case int32(398):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_texture_depth_cube_array)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(399):
		if lookahead == int32('e') {
			state = uint16(401)
			goto next_state
		}
		return result
	case int32(400):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_texture_storage_2d_array)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(401):
		if lookahead == int32('d') {
			state = uint16(402)
			goto next_state
		}
		return result
	case int32(402):
		if lookahead == int32('_') {
			state = uint16(403)
			goto next_state
		}
		return result
	case int32(403):
		if lookahead == int32('2') {
			state = uint16(404)
			goto next_state
		}
		return result
	case int32(404):
		if lookahead == int32('d') {
			state = uint16(405)
			goto next_state
		}
		return result
	case int32(405):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_texture_depth_multisampled_2d)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var map_token5 = [34]uint16_t{
	0:  uint16('a'),
	1:  uint16(1),
	2:  uint16('b'),
	3:  uint16(2),
	4:  uint16('c'),
	5:  uint16(3),
	6:  uint16('d'),
	7:  uint16(4),
	8:  uint16('e'),
	9:  uint16(5),
	10: uint16('f'),
	11: uint16(6),
	12: uint16('i'),
	13: uint16(7),
	14: uint16('l'),
	15: uint16(8),
	16: uint16('m'),
	17: uint16(9),
	18: uint16('o'),
	19: uint16(10),
	20: uint16('p'),
	21: uint16(11),
	22: uint16('r'),
	23: uint16(12),
	24: uint16('s'),
	25: uint16(13),
	26: uint16('t'),
	27: uint16(14),
	28: uint16('u'),
	29: uint16(15),
	30: uint16('v'),
	31: uint16(16),
	32: uint16('w'),
	33: uint16(17),
}

var ts_lex_modes = [443]TSLexMode{
	0: {
		Fexternal_lex_state: uint16(1),
	},
	1: {
		Fexternal_lex_state: uint16(1),
	},
	2: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	3: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	4: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	5: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	6: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	7: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	8: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	9: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	10: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	11: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	12: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	13: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	14: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	15: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	16: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	17: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	18: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	19: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	20: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	21: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	22: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	23: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	24: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	25: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	26: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	27: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	28: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	29: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	30: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	31: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	32: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	33: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	34: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	35: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	36: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	37: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	38: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	39: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	40: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	41: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	42: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	43: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	44: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	45: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	46: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	47: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	48: {
		Fexternal_lex_state: uint16(1),
	},
	49: {
		Fexternal_lex_state: uint16(1),
	},
	50: {
		Fexternal_lex_state: uint16(1),
	},
	51: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	52: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	53: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	54: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	55: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	56: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	57: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	58: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	59: {
		Fexternal_lex_state: uint16(1),
	},
	60: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	61: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	62: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	63: {
		Fexternal_lex_state: uint16(1),
	},
	64: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	65: {
		Fexternal_lex_state: uint16(1),
	},
	66: {
		Fexternal_lex_state: uint16(1),
	},
	67: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	68: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	69: {
		Fexternal_lex_state: uint16(1),
	},
	70: {
		Fexternal_lex_state: uint16(1),
	},
	71: {
		Fexternal_lex_state: uint16(1),
	},
	72: {
		Fexternal_lex_state: uint16(1),
	},
	73: {
		Fexternal_lex_state: uint16(1),
	},
	74: {
		Fexternal_lex_state: uint16(1),
	},
	75: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	76: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	77: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	78: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	79: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	80: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	81: {
		Fexternal_lex_state: uint16(1),
	},
	82: {
		Fexternal_lex_state: uint16(1),
	},
	83: {
		Fexternal_lex_state: uint16(1),
	},
	84: {
		Fexternal_lex_state: uint16(1),
	},
	85: {
		Fexternal_lex_state: uint16(1),
	},
	86: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	87: {
		Fexternal_lex_state: uint16(1),
	},
	88: {
		Fexternal_lex_state: uint16(1),
	},
	89: {
		Fexternal_lex_state: uint16(1),
	},
	90: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	91: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(1),
	},
	92: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(1),
	},
	93: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(1),
	},
	94: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	95: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(1),
	},
	96: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(1),
	},
	97: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(1),
	},
	98: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(1),
	},
	99: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(1),
	},
	100: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(1),
	},
	101: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(1),
	},
	102: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	103: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	104: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	105: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	106: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(1),
	},
	107: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(1),
	},
	108: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(1),
	},
	109: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(1),
	},
	110: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(1),
	},
	111: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(1),
	},
	112: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(1),
	},
	113: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(1),
	},
	114: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	115: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(1),
	},
	116: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	117: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	118: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(1),
	},
	119: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(1),
	},
	120: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(1),
	},
	121: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	122: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	123: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	124: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	125: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	126: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	127: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	128: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	129: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	130: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(1),
	},
	131: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(1),
	},
	132: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(1),
	},
	133: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(1),
	},
	134: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(1),
	},
	135: {
		Flex_state:          uint16(2),
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
		Fexternal_lex_state: uint16(1),
	},
	152: {
		Fexternal_lex_state: uint16(1),
	},
	153: {
		Fexternal_lex_state: uint16(1),
	},
	154: {
		Fexternal_lex_state: uint16(1),
	},
	155: {
		Fexternal_lex_state: uint16(1),
	},
	156: {
		Fexternal_lex_state: uint16(1),
	},
	157: {
		Fexternal_lex_state: uint16(1),
	},
	158: {
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
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	174: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	175: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(1),
	},
	176: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	177: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(1),
	},
	178: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(1),
	},
	179: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(1),
	},
	180: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(1),
	},
	181: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(1),
	},
	182: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(1),
	},
	183: {
		Fexternal_lex_state: uint16(1),
	},
	184: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(1),
	},
	185: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	186: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(1),
	},
	187: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(1),
	},
	188: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(1),
	},
	189: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(1),
	},
	190: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(1),
	},
	191: {
		Fexternal_lex_state: uint16(1),
	},
	192: {
		Fexternal_lex_state: uint16(1),
	},
	193: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(1),
	},
	194: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(1),
	},
	195: {
		Fexternal_lex_state: uint16(1),
	},
	196: {
		Fexternal_lex_state: uint16(1),
	},
	197: {
		Fexternal_lex_state: uint16(1),
	},
	198: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	199: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(1),
	},
	200: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(1),
	},
	201: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(1),
	},
	202: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(1),
	},
	203: {
		Fexternal_lex_state: uint16(1),
	},
	204: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(1),
	},
	205: {
		Fexternal_lex_state: uint16(1),
	},
	206: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	207: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	208: {
		Fexternal_lex_state: uint16(1),
	},
	209: {
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
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	214: {
		Fexternal_lex_state: uint16(1),
	},
	215: {
		Fexternal_lex_state: uint16(1),
	},
	216: {
		Fexternal_lex_state: uint16(1),
	},
	217: {
		Fexternal_lex_state: uint16(1),
	},
	218: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	219: {
		Fexternal_lex_state: uint16(1),
	},
	220: {
		Fexternal_lex_state: uint16(1),
	},
	221: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	222: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	223: {
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
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	228: {
		Fexternal_lex_state: uint16(1),
	},
	229: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	230: {
		Fexternal_lex_state: uint16(1),
	},
	231: {
		Fexternal_lex_state: uint16(1),
	},
	232: {
		Flex_state:          uint16(64),
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
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	239: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	240: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	241: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	242: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	243: {
		Fexternal_lex_state: uint16(1),
	},
	244: {
		Fexternal_lex_state: uint16(1),
	},
	245: {
		Fexternal_lex_state: uint16(1),
	},
	246: {
		Fexternal_lex_state: uint16(1),
	},
	247: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	248: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	249: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(1),
	},
	250: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(1),
	},
	251: {
		Fexternal_lex_state: uint16(1),
	},
	252: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(1),
	},
	253: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	254: {
		Fexternal_lex_state: uint16(1),
	},
	255: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	256: {
		Flex_state:          uint16(3),
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
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	262: {
		Fexternal_lex_state: uint16(1),
	},
	263: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(1),
	},
	264: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	265: {
		Fexternal_lex_state: uint16(1),
	},
	266: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(1),
	},
	267: {
		Fexternal_lex_state: uint16(1),
	},
	268: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	269: {
		Flex_state:          uint16(64),
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
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	274: {
		Fexternal_lex_state: uint16(1),
	},
	275: {
		Fexternal_lex_state: uint16(1),
	},
	276: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	277: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	278: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	279: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	280: {
		Fexternal_lex_state: uint16(1),
	},
	281: {
		Flex_state:          uint16(64),
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
		Flex_state:          uint16(1),
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
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	290: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	291: {
		Fexternal_lex_state: uint16(1),
	},
	292: {
		Fexternal_lex_state: uint16(1),
	},
	293: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(1),
	},
	294: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(1),
	},
	295: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(1),
	},
	296: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	297: {
		Fexternal_lex_state: uint16(1),
	},
	298: {
		Fexternal_lex_state: uint16(1),
	},
	299: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	300: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	301: {
		Fexternal_lex_state: uint16(1),
	},
	302: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	303: {
		Fexternal_lex_state: uint16(1),
	},
	304: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(1),
	},
	305: {
		Fexternal_lex_state: uint16(1),
	},
	306: {
		Fexternal_lex_state: uint16(1),
	},
	307: {
		Flex_state:          uint16(64),
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
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	312: {
		Flex_state:          uint16(64),
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
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	318: {
		Fexternal_lex_state: uint16(1),
	},
	319: {
		Fexternal_lex_state: uint16(1),
	},
	320: {
		Fexternal_lex_state: uint16(1),
	},
	321: {
		Fexternal_lex_state: uint16(1),
	},
	322: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	323: {
		Fexternal_lex_state: uint16(1),
	},
	324: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	325: {
		Fexternal_lex_state: uint16(1),
	},
	326: {
		Fexternal_lex_state: uint16(1),
	},
	327: {
		Fexternal_lex_state: uint16(1),
	},
	328: {
		Fexternal_lex_state: uint16(1),
	},
	329: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	330: {
		Flex_state:          uint16(64),
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
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	335: {
		Fexternal_lex_state: uint16(1),
	},
	336: {
		Fexternal_lex_state: uint16(1),
	},
	337: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	338: {
		Fexternal_lex_state: uint16(1),
	},
	339: {
		Flex_state:          uint16(3),
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
		Flex_state:          uint16(1),
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
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	348: {
		Fexternal_lex_state: uint16(1),
	},
	349: {
		Flex_state:          uint16(64),
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
		Flex_state:          uint16(1),
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
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	359: {
		Fexternal_lex_state: uint16(1),
	},
	360: {
		Fexternal_lex_state: uint16(1),
	},
	361: {
		Fexternal_lex_state: uint16(1),
	},
	362: {
		Fexternal_lex_state: uint16(1),
	},
	363: {
		Fexternal_lex_state: uint16(1),
	},
	364: {
		Fexternal_lex_state: uint16(1),
	},
	365: {
		Fexternal_lex_state: uint16(1),
	},
	366: {
		Fexternal_lex_state: uint16(1),
	},
	367: {
		Fexternal_lex_state: uint16(1),
	},
	368: {
		Fexternal_lex_state: uint16(1),
	},
	369: {
		Fexternal_lex_state: uint16(1),
	},
	370: {
		Fexternal_lex_state: uint16(1),
	},
	371: {
		Fexternal_lex_state: uint16(1),
	},
	372: {
		Fexternal_lex_state: uint16(1),
	},
	373: {
		Fexternal_lex_state: uint16(1),
	},
	374: {
		Fexternal_lex_state: uint16(1),
	},
	375: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	376: {
		Fexternal_lex_state: uint16(1),
	},
	377: {
		Fexternal_lex_state: uint16(1),
	},
	378: {
		Fexternal_lex_state: uint16(1),
	},
	379: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	380: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	381: {
		Fexternal_lex_state: uint16(1),
	},
	382: {
		Fexternal_lex_state: uint16(1),
	},
	383: {
		Fexternal_lex_state: uint16(1),
	},
	384: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	385: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	386: {
		Fexternal_lex_state: uint16(1),
	},
	387: {
		Fexternal_lex_state: uint16(1),
	},
	388: {
		Fexternal_lex_state: uint16(1),
	},
	389: {
		Fexternal_lex_state: uint16(1),
	},
	390: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	391: {
		Fexternal_lex_state: uint16(1),
	},
	392: {
		Fexternal_lex_state: uint16(1),
	},
	393: {
		Fexternal_lex_state: uint16(1),
	},
	394: {
		Fexternal_lex_state: uint16(1),
	},
	395: {
		Fexternal_lex_state: uint16(1),
	},
	396: {
		Fexternal_lex_state: uint16(1),
	},
	397: {
		Fexternal_lex_state: uint16(1),
	},
	398: {
		Fexternal_lex_state: uint16(1),
	},
	399: {
		Fexternal_lex_state: uint16(1),
	},
	400: {
		Fexternal_lex_state: uint16(1),
	},
	401: {
		Fexternal_lex_state: uint16(1),
	},
	402: {
		Fexternal_lex_state: uint16(1),
	},
	403: {
		Fexternal_lex_state: uint16(1),
	},
	404: {
		Fexternal_lex_state: uint16(1),
	},
	405: {
		Fexternal_lex_state: uint16(1),
	},
	406: {
		Fexternal_lex_state: uint16(1),
	},
	407: {
		Fexternal_lex_state: uint16(1),
	},
	408: {
		Fexternal_lex_state: uint16(1),
	},
	409: {
		Fexternal_lex_state: uint16(1),
	},
	410: {
		Fexternal_lex_state: uint16(1),
	},
	411: {
		Fexternal_lex_state: uint16(1),
	},
	412: {
		Fexternal_lex_state: uint16(1),
	},
	413: {
		Fexternal_lex_state: uint16(1),
	},
	414: {
		Fexternal_lex_state: uint16(1),
	},
	415: {
		Fexternal_lex_state: uint16(1),
	},
	416: {
		Fexternal_lex_state: uint16(1),
	},
	417: {
		Fexternal_lex_state: uint16(1),
	},
	418: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	419: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	420: {
		Fexternal_lex_state: uint16(1),
	},
	421: {
		Fexternal_lex_state: uint16(1),
	},
	422: {
		Fexternal_lex_state: uint16(1),
	},
	423: {
		Fexternal_lex_state: uint16(1),
	},
	424: {
		Fexternal_lex_state: uint16(1),
	},
	425: {
		Fexternal_lex_state: uint16(1),
	},
	426: {
		Fexternal_lex_state: uint16(1),
	},
	427: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(1),
	},
	428: {
		Fexternal_lex_state: uint16(1),
	},
	429: {
		Fexternal_lex_state: uint16(1),
	},
	430: {
		Fexternal_lex_state: uint16(1),
	},
	431: {
		Fexternal_lex_state: uint16(1),
	},
	432: {
		Fexternal_lex_state: uint16(1),
	},
	433: {
		Fexternal_lex_state: uint16(1),
	},
	434: {
		Fexternal_lex_state: uint16(1),
	},
	435: {
		Fexternal_lex_state: uint16(1),
	},
	436: {
		Fexternal_lex_state: uint16(1),
	},
	437: {
		Fexternal_lex_state: uint16(1),
	},
	438: {
		Fexternal_lex_state: uint16(1),
	},
	439: {
		Fexternal_lex_state: uint16(1),
	},
	440: {
		Fexternal_lex_state: uint16(1),
	},
	441: {
		Fexternal_lex_state: uint16(1),
	},
	442: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
}

var ts_parse_table = [32][235]uint16_t{
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
		46:  uint16(1),
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
		138: uint16(1),
		139: uint16(1),
		140: uint16(1),
		141: uint16(1),
		142: uint16(1),
		143: uint16(1),
		144: uint16(1),
		145: uint16(3),
	},
	1: {
		0:   uint16(5),
		2:   uint16(3),
		3:   uint16(7),
		5:   uint16(9),
		6:   uint16(11),
		7:   uint16(13),
		11:  uint16(15),
		12:  uint16(17),
		14:  uint16(19),
		17:  uint16(21),
		18:  uint16(23),
		43:  uint16(25),
		136: uint16(27),
		138: uint16(29),
		141: uint16(31),
		142: uint16(31),
		145: uint16(3),
		146: uint16(402),
		147: uint16(88),
		148: uint16(417),
		149: uint16(417),
		150: uint16(417),
		152: uint16(88),
		154: uint16(88),
		156: uint16(81),
		157: uint16(226),
		184: uint16(324),
		210: uint16(88),
		211: uint16(88),
		214: uint16(88),
		220: uint16(81),
		221: uint16(88),
		222: uint16(226),
	},
	2: {
		1:   uint16(33),
		2:   uint16(3),
		8:   uint16(35),
		10:  uint16(37),
		48:  uint16(39),
		49:  uint16(41),
		50:  uint16(41),
		51:  uint16(43),
		52:  uint16(43),
		53:  uint16(45),
		54:  uint16(45),
		55:  uint16(45),
		56:  uint16(45),
		57:  uint16(45),
		58:  uint16(47),
		59:  uint16(49),
		60:  uint16(45),
		61:  uint16(45),
		62:  uint16(45),
		63:  uint16(45),
		64:  uint16(45),
		65:  uint16(45),
		66:  uint16(45),
		67:  uint16(51),
		68:  uint16(51),
		69:  uint16(51),
		70:  uint16(51),
		71:  uint16(51),
		72:  uint16(51),
		73:  uint16(51),
		74:  uint16(53),
		75:  uint16(53),
		76:  uint16(53),
		77:  uint16(53),
		78:  uint16(55),
		79:  uint16(55),
		80:  uint16(55),
		81:  uint16(55),
		82:  uint16(55),
		83:  uint16(55),
		84:  uint16(55),
		85:  uint16(55),
		86:  uint16(55),
		87:  uint16(55),
		88:  uint16(55),
		89:  uint16(55),
		114: uint16(57),
		119: uint16(59),
		127: uint16(61),
		128: uint16(59),
		131: uint16(59),
		132: uint16(59),
		145: uint16(3),
		189: uint16(132),
		190: uint16(132),
		191: uint16(93),
		192: uint16(93),
		193: uint16(132),
		194: uint16(132),
		195: uint16(323),
		196: uint16(283),
		197: uint16(283),
		202: uint16(132),
		203: uint16(132),
		204: uint16(132),
		206: uint16(132),
		208: uint16(132),
		229: uint16(3),
	},
	3: {
		1:   uint16(33),
		2:   uint16(3),
		8:   uint16(35),
		48:  uint16(39),
		49:  uint16(41),
		50:  uint16(41),
		51:  uint16(43),
		52:  uint16(43),
		53:  uint16(45),
		54:  uint16(45),
		55:  uint16(45),
		56:  uint16(45),
		57:  uint16(45),
		58:  uint16(47),
		59:  uint16(49),
		60:  uint16(45),
		61:  uint16(45),
		62:  uint16(45),
		63:  uint16(45),
		64:  uint16(45),
		65:  uint16(45),
		66:  uint16(45),
		67:  uint16(51),
		68:  uint16(51),
		69:  uint16(51),
		70:  uint16(51),
		71:  uint16(51),
		72:  uint16(51),
		73:  uint16(51),
		74:  uint16(53),
		75:  uint16(53),
		76:  uint16(53),
		77:  uint16(53),
		78:  uint16(55),
		79:  uint16(55),
		80:  uint16(55),
		81:  uint16(55),
		82:  uint16(55),
		83:  uint16(55),
		84:  uint16(55),
		85:  uint16(55),
		86:  uint16(55),
		87:  uint16(55),
		88:  uint16(55),
		89:  uint16(55),
		114: uint16(57),
		119: uint16(59),
		127: uint16(61),
		128: uint16(59),
		131: uint16(59),
		132: uint16(59),
		145: uint16(3),
		189: uint16(136),
		190: uint16(136),
		191: uint16(93),
		192: uint16(93),
		193: uint16(136),
		194: uint16(136),
		195: uint16(323),
		196: uint16(283),
		197: uint16(283),
		202: uint16(136),
		203: uint16(136),
		204: uint16(136),
		206: uint16(136),
		208: uint16(136),
		229: uint16(5),
	},
	4: {
		1:   uint16(33),
		2:   uint16(3),
		3:   uint16(63),
		8:   uint16(35),
		48:  uint16(39),
		49:  uint16(41),
		50:  uint16(41),
		51:  uint16(43),
		52:  uint16(43),
		53:  uint16(45),
		54:  uint16(45),
		55:  uint16(45),
		56:  uint16(45),
		57:  uint16(45),
		58:  uint16(47),
		59:  uint16(49),
		60:  uint16(45),
		61:  uint16(45),
		62:  uint16(45),
		63:  uint16(45),
		64:  uint16(45),
		65:  uint16(45),
		66:  uint16(45),
		67:  uint16(51),
		68:  uint16(51),
		69:  uint16(51),
		70:  uint16(51),
		71:  uint16(51),
		72:  uint16(51),
		73:  uint16(51),
		74:  uint16(53),
		75:  uint16(53),
		76:  uint16(53),
		77:  uint16(53),
		78:  uint16(55),
		79:  uint16(55),
		80:  uint16(55),
		81:  uint16(55),
		82:  uint16(55),
		83:  uint16(55),
		84:  uint16(55),
		85:  uint16(55),
		86:  uint16(55),
		87:  uint16(55),
		88:  uint16(55),
		89:  uint16(55),
		114: uint16(57),
		119: uint16(59),
		127: uint16(61),
		128: uint16(59),
		131: uint16(59),
		132: uint16(59),
		145: uint16(3),
		189: uint16(141),
		190: uint16(141),
		191: uint16(93),
		192: uint16(93),
		193: uint16(141),
		194: uint16(141),
		195: uint16(323),
		196: uint16(283),
		197: uint16(283),
		202: uint16(141),
		203: uint16(141),
		204: uint16(141),
		206: uint16(141),
		208: uint16(141),
	},
	5: {
		1:   uint16(65),
		2:   uint16(3),
		8:   uint16(68),
		48:  uint16(71),
		49:  uint16(74),
		50:  uint16(74),
		51:  uint16(77),
		52:  uint16(77),
		53:  uint16(80),
		54:  uint16(80),
		55:  uint16(80),
		56:  uint16(80),
		57:  uint16(80),
		58:  uint16(83),
		59:  uint16(86),
		60:  uint16(80),
		61:  uint16(80),
		62:  uint16(80),
		63:  uint16(80),
		64:  uint16(80),
		65:  uint16(80),
		66:  uint16(80),
		67:  uint16(89),
		68:  uint16(89),
		69:  uint16(89),
		70:  uint16(89),
		71:  uint16(89),
		72:  uint16(89),
		73:  uint16(89),
		74:  uint16(92),
		75:  uint16(92),
		76:  uint16(92),
		77:  uint16(92),
		78:  uint16(95),
		79:  uint16(95),
		80:  uint16(95),
		81:  uint16(95),
		82:  uint16(95),
		83:  uint16(95),
		84:  uint16(95),
		85:  uint16(95),
		86:  uint16(95),
		87:  uint16(95),
		88:  uint16(95),
		89:  uint16(95),
		114: uint16(98),
		119: uint16(101),
		127: uint16(104),
		128: uint16(101),
		131: uint16(101),
		132: uint16(101),
		145: uint16(3),
		189: uint16(145),
		190: uint16(145),
		191: uint16(93),
		192: uint16(93),
		193: uint16(145),
		194: uint16(145),
		195: uint16(323),
		196: uint16(283),
		197: uint16(283),
		202: uint16(145),
		203: uint16(145),
		204: uint16(145),
		206: uint16(145),
		208: uint16(145),
		229: uint16(5),
	},
	6: {
		1:   uint16(33),
		2:   uint16(3),
		3:   uint16(107),
		8:   uint16(35),
		48:  uint16(39),
		49:  uint16(41),
		50:  uint16(41),
		51:  uint16(43),
		52:  uint16(43),
		53:  uint16(45),
		54:  uint16(45),
		55:  uint16(45),
		56:  uint16(45),
		57:  uint16(45),
		58:  uint16(47),
		59:  uint16(49),
		60:  uint16(45),
		61:  uint16(45),
		62:  uint16(45),
		63:  uint16(45),
		64:  uint16(45),
		65:  uint16(45),
		66:  uint16(45),
		67:  uint16(51),
		68:  uint16(51),
		69:  uint16(51),
		70:  uint16(51),
		71:  uint16(51),
		72:  uint16(51),
		73:  uint16(51),
		74:  uint16(53),
		75:  uint16(53),
		76:  uint16(53),
		77:  uint16(53),
		78:  uint16(55),
		79:  uint16(55),
		80:  uint16(55),
		81:  uint16(55),
		82:  uint16(55),
		83:  uint16(55),
		84:  uint16(55),
		85:  uint16(55),
		86:  uint16(55),
		87:  uint16(55),
		88:  uint16(55),
		89:  uint16(55),
		114: uint16(57),
		119: uint16(59),
		127: uint16(61),
		128: uint16(59),
		131: uint16(59),
		132: uint16(59),
		145: uint16(3),
		189: uint16(142),
		190: uint16(142),
		191: uint16(93),
		192: uint16(93),
		193: uint16(142),
		194: uint16(142),
		195: uint16(323),
		196: uint16(283),
		197: uint16(283),
		202: uint16(142),
		203: uint16(142),
		204: uint16(142),
		206: uint16(142),
		208: uint16(142),
	},
	7: {
		1:   uint16(33),
		2:   uint16(3),
		3:   uint16(109),
		8:   uint16(35),
		48:  uint16(39),
		49:  uint16(41),
		50:  uint16(41),
		51:  uint16(43),
		52:  uint16(43),
		53:  uint16(45),
		54:  uint16(45),
		55:  uint16(45),
		56:  uint16(45),
		57:  uint16(45),
		58:  uint16(47),
		59:  uint16(49),
		60:  uint16(45),
		61:  uint16(45),
		62:  uint16(45),
		63:  uint16(45),
		64:  uint16(45),
		65:  uint16(45),
		66:  uint16(45),
		67:  uint16(51),
		68:  uint16(51),
		69:  uint16(51),
		70:  uint16(51),
		71:  uint16(51),
		72:  uint16(51),
		73:  uint16(51),
		74:  uint16(53),
		75:  uint16(53),
		76:  uint16(53),
		77:  uint16(53),
		78:  uint16(55),
		79:  uint16(55),
		80:  uint16(55),
		81:  uint16(55),
		82:  uint16(55),
		83:  uint16(55),
		84:  uint16(55),
		85:  uint16(55),
		86:  uint16(55),
		87:  uint16(55),
		88:  uint16(55),
		89:  uint16(55),
		114: uint16(57),
		119: uint16(59),
		127: uint16(61),
		128: uint16(59),
		131: uint16(59),
		132: uint16(59),
		145: uint16(3),
		189: uint16(147),
		190: uint16(147),
		191: uint16(93),
		192: uint16(93),
		193: uint16(147),
		194: uint16(147),
		195: uint16(323),
		196: uint16(283),
		197: uint16(283),
		202: uint16(147),
		203: uint16(147),
		204: uint16(147),
		206: uint16(147),
		208: uint16(147),
	},
	8: {
		1:   uint16(33),
		2:   uint16(3),
		8:   uint16(35),
		48:  uint16(39),
		49:  uint16(41),
		50:  uint16(41),
		51:  uint16(43),
		52:  uint16(43),
		53:  uint16(45),
		54:  uint16(45),
		55:  uint16(45),
		56:  uint16(45),
		57:  uint16(45),
		58:  uint16(47),
		59:  uint16(49),
		60:  uint16(45),
		61:  uint16(45),
		62:  uint16(45),
		63:  uint16(45),
		64:  uint16(45),
		65:  uint16(45),
		66:  uint16(45),
		67:  uint16(51),
		68:  uint16(51),
		69:  uint16(51),
		70:  uint16(51),
		71:  uint16(51),
		72:  uint16(51),
		73:  uint16(51),
		74:  uint16(53),
		75:  uint16(53),
		76:  uint16(53),
		77:  uint16(53),
		78:  uint16(55),
		79:  uint16(55),
		80:  uint16(55),
		81:  uint16(55),
		82:  uint16(55),
		83:  uint16(55),
		84:  uint16(55),
		85:  uint16(55),
		86:  uint16(55),
		87:  uint16(55),
		88:  uint16(55),
		89:  uint16(55),
		114: uint16(57),
		119: uint16(59),
		127: uint16(61),
		128: uint16(59),
		131: uint16(59),
		132: uint16(59),
		145: uint16(3),
		189: uint16(150),
		190: uint16(150),
		191: uint16(93),
		192: uint16(93),
		193: uint16(150),
		194: uint16(150),
		195: uint16(323),
		196: uint16(283),
		197: uint16(283),
		202: uint16(150),
		203: uint16(150),
		204: uint16(150),
		206: uint16(150),
		208: uint16(150),
	},
	9: {
		1:   uint16(33),
		2:   uint16(3),
		8:   uint16(35),
		48:  uint16(39),
		49:  uint16(41),
		50:  uint16(41),
		51:  uint16(43),
		52:  uint16(43),
		53:  uint16(45),
		54:  uint16(45),
		55:  uint16(45),
		56:  uint16(45),
		57:  uint16(45),
		58:  uint16(47),
		59:  uint16(49),
		60:  uint16(45),
		61:  uint16(45),
		62:  uint16(45),
		63:  uint16(45),
		64:  uint16(45),
		65:  uint16(45),
		66:  uint16(45),
		67:  uint16(51),
		68:  uint16(51),
		69:  uint16(51),
		70:  uint16(51),
		71:  uint16(51),
		72:  uint16(51),
		73:  uint16(51),
		74:  uint16(53),
		75:  uint16(53),
		76:  uint16(53),
		77:  uint16(53),
		78:  uint16(55),
		79:  uint16(55),
		80:  uint16(55),
		81:  uint16(55),
		82:  uint16(55),
		83:  uint16(55),
		84:  uint16(55),
		85:  uint16(55),
		86:  uint16(55),
		87:  uint16(55),
		88:  uint16(55),
		89:  uint16(55),
		114: uint16(57),
		119: uint16(59),
		127: uint16(61),
		128: uint16(59),
		131: uint16(59),
		132: uint16(59),
		145: uint16(3),
		189: uint16(108),
		190: uint16(108),
		191: uint16(93),
		192: uint16(93),
		193: uint16(108),
		194: uint16(108),
		195: uint16(323),
		196: uint16(283),
		197: uint16(283),
		202: uint16(108),
		203: uint16(108),
		204: uint16(108),
		206: uint16(108),
		208: uint16(108),
	},
	10: {
		1:   uint16(33),
		2:   uint16(3),
		8:   uint16(35),
		48:  uint16(39),
		49:  uint16(41),
		50:  uint16(41),
		51:  uint16(43),
		52:  uint16(43),
		53:  uint16(45),
		54:  uint16(45),
		55:  uint16(45),
		56:  uint16(45),
		57:  uint16(45),
		58:  uint16(47),
		59:  uint16(49),
		60:  uint16(45),
		61:  uint16(45),
		62:  uint16(45),
		63:  uint16(45),
		64:  uint16(45),
		65:  uint16(45),
		66:  uint16(45),
		67:  uint16(51),
		68:  uint16(51),
		69:  uint16(51),
		70:  uint16(51),
		71:  uint16(51),
		72:  uint16(51),
		73:  uint16(51),
		74:  uint16(53),
		75:  uint16(53),
		76:  uint16(53),
		77:  uint16(53),
		78:  uint16(55),
		79:  uint16(55),
		80:  uint16(55),
		81:  uint16(55),
		82:  uint16(55),
		83:  uint16(55),
		84:  uint16(55),
		85:  uint16(55),
		86:  uint16(55),
		87:  uint16(55),
		88:  uint16(55),
		89:  uint16(55),
		114: uint16(57),
		119: uint16(59),
		127: uint16(61),
		128: uint16(59),
		131: uint16(59),
		132: uint16(59),
		145: uint16(3),
		189: uint16(109),
		190: uint16(109),
		191: uint16(93),
		192: uint16(93),
		193: uint16(109),
		194: uint16(109),
		195: uint16(323),
		196: uint16(283),
		197: uint16(283),
		202: uint16(109),
		203: uint16(109),
		204: uint16(109),
		206: uint16(109),
		208: uint16(109),
	},
	11: {
		1:   uint16(33),
		2:   uint16(3),
		8:   uint16(35),
		48:  uint16(39),
		49:  uint16(41),
		50:  uint16(41),
		51:  uint16(43),
		52:  uint16(43),
		53:  uint16(45),
		54:  uint16(45),
		55:  uint16(45),
		56:  uint16(45),
		57:  uint16(45),
		58:  uint16(47),
		59:  uint16(49),
		60:  uint16(45),
		61:  uint16(45),
		62:  uint16(45),
		63:  uint16(45),
		64:  uint16(45),
		65:  uint16(45),
		66:  uint16(45),
		67:  uint16(51),
		68:  uint16(51),
		69:  uint16(51),
		70:  uint16(51),
		71:  uint16(51),
		72:  uint16(51),
		73:  uint16(51),
		74:  uint16(53),
		75:  uint16(53),
		76:  uint16(53),
		77:  uint16(53),
		78:  uint16(55),
		79:  uint16(55),
		80:  uint16(55),
		81:  uint16(55),
		82:  uint16(55),
		83:  uint16(55),
		84:  uint16(55),
		85:  uint16(55),
		86:  uint16(55),
		87:  uint16(55),
		88:  uint16(55),
		89:  uint16(55),
		114: uint16(57),
		119: uint16(59),
		127: uint16(61),
		128: uint16(59),
		131: uint16(59),
		132: uint16(59),
		145: uint16(3),
		189: uint16(113),
		190: uint16(113),
		191: uint16(93),
		192: uint16(93),
		193: uint16(113),
		194: uint16(113),
		195: uint16(323),
		196: uint16(283),
		197: uint16(283),
		202: uint16(113),
		203: uint16(113),
		204: uint16(113),
		206: uint16(113),
		208: uint16(113),
	},
	12: {
		1:   uint16(33),
		2:   uint16(3),
		8:   uint16(35),
		48:  uint16(39),
		49:  uint16(41),
		50:  uint16(41),
		51:  uint16(43),
		52:  uint16(43),
		53:  uint16(45),
		54:  uint16(45),
		55:  uint16(45),
		56:  uint16(45),
		57:  uint16(45),
		58:  uint16(47),
		59:  uint16(49),
		60:  uint16(45),
		61:  uint16(45),
		62:  uint16(45),
		63:  uint16(45),
		64:  uint16(45),
		65:  uint16(45),
		66:  uint16(45),
		67:  uint16(51),
		68:  uint16(51),
		69:  uint16(51),
		70:  uint16(51),
		71:  uint16(51),
		72:  uint16(51),
		73:  uint16(51),
		74:  uint16(53),
		75:  uint16(53),
		76:  uint16(53),
		77:  uint16(53),
		78:  uint16(55),
		79:  uint16(55),
		80:  uint16(55),
		81:  uint16(55),
		82:  uint16(55),
		83:  uint16(55),
		84:  uint16(55),
		85:  uint16(55),
		86:  uint16(55),
		87:  uint16(55),
		88:  uint16(55),
		89:  uint16(55),
		114: uint16(57),
		119: uint16(59),
		127: uint16(61),
		128: uint16(59),
		131: uint16(59),
		132: uint16(59),
		145: uint16(3),
		189: uint16(115),
		190: uint16(115),
		191: uint16(93),
		192: uint16(93),
		193: uint16(115),
		194: uint16(115),
		195: uint16(323),
		196: uint16(283),
		197: uint16(283),
		202: uint16(115),
		203: uint16(115),
		204: uint16(115),
		206: uint16(115),
		208: uint16(115),
	},
	13: {
		1:   uint16(33),
		2:   uint16(3),
		8:   uint16(35),
		48:  uint16(39),
		49:  uint16(41),
		50:  uint16(41),
		51:  uint16(43),
		52:  uint16(43),
		53:  uint16(45),
		54:  uint16(45),
		55:  uint16(45),
		56:  uint16(45),
		57:  uint16(45),
		58:  uint16(47),
		59:  uint16(49),
		60:  uint16(45),
		61:  uint16(45),
		62:  uint16(45),
		63:  uint16(45),
		64:  uint16(45),
		65:  uint16(45),
		66:  uint16(45),
		67:  uint16(51),
		68:  uint16(51),
		69:  uint16(51),
		70:  uint16(51),
		71:  uint16(51),
		72:  uint16(51),
		73:  uint16(51),
		74:  uint16(53),
		75:  uint16(53),
		76:  uint16(53),
		77:  uint16(53),
		78:  uint16(55),
		79:  uint16(55),
		80:  uint16(55),
		81:  uint16(55),
		82:  uint16(55),
		83:  uint16(55),
		84:  uint16(55),
		85:  uint16(55),
		86:  uint16(55),
		87:  uint16(55),
		88:  uint16(55),
		89:  uint16(55),
		114: uint16(57),
		119: uint16(59),
		127: uint16(61),
		128: uint16(59),
		131: uint16(59),
		132: uint16(59),
		145: uint16(3),
		189: uint16(119),
		190: uint16(119),
		191: uint16(93),
		192: uint16(93),
		193: uint16(119),
		194: uint16(119),
		195: uint16(323),
		196: uint16(283),
		197: uint16(283),
		202: uint16(119),
		203: uint16(119),
		204: uint16(119),
		206: uint16(119),
		208: uint16(119),
	},
	14: {
		1:   uint16(33),
		2:   uint16(3),
		8:   uint16(35),
		48:  uint16(39),
		49:  uint16(41),
		50:  uint16(41),
		51:  uint16(43),
		52:  uint16(43),
		53:  uint16(45),
		54:  uint16(45),
		55:  uint16(45),
		56:  uint16(45),
		57:  uint16(45),
		58:  uint16(47),
		59:  uint16(49),
		60:  uint16(45),
		61:  uint16(45),
		62:  uint16(45),
		63:  uint16(45),
		64:  uint16(45),
		65:  uint16(45),
		66:  uint16(45),
		67:  uint16(51),
		68:  uint16(51),
		69:  uint16(51),
		70:  uint16(51),
		71:  uint16(51),
		72:  uint16(51),
		73:  uint16(51),
		74:  uint16(53),
		75:  uint16(53),
		76:  uint16(53),
		77:  uint16(53),
		78:  uint16(55),
		79:  uint16(55),
		80:  uint16(55),
		81:  uint16(55),
		82:  uint16(55),
		83:  uint16(55),
		84:  uint16(55),
		85:  uint16(55),
		86:  uint16(55),
		87:  uint16(55),
		88:  uint16(55),
		89:  uint16(55),
		114: uint16(57),
		119: uint16(59),
		127: uint16(61),
		128: uint16(59),
		131: uint16(59),
		132: uint16(59),
		145: uint16(3),
		189: uint16(120),
		190: uint16(120),
		191: uint16(93),
		192: uint16(93),
		193: uint16(120),
		194: uint16(120),
		195: uint16(323),
		196: uint16(283),
		197: uint16(283),
		202: uint16(120),
		203: uint16(120),
		204: uint16(120),
		206: uint16(120),
		208: uint16(120),
	},
	15: {
		1:   uint16(33),
		2:   uint16(3),
		8:   uint16(35),
		48:  uint16(39),
		49:  uint16(41),
		50:  uint16(41),
		51:  uint16(43),
		52:  uint16(43),
		53:  uint16(45),
		54:  uint16(45),
		55:  uint16(45),
		56:  uint16(45),
		57:  uint16(45),
		58:  uint16(47),
		59:  uint16(49),
		60:  uint16(45),
		61:  uint16(45),
		62:  uint16(45),
		63:  uint16(45),
		64:  uint16(45),
		65:  uint16(45),
		66:  uint16(45),
		67:  uint16(51),
		68:  uint16(51),
		69:  uint16(51),
		70:  uint16(51),
		71:  uint16(51),
		72:  uint16(51),
		73:  uint16(51),
		74:  uint16(53),
		75:  uint16(53),
		76:  uint16(53),
		77:  uint16(53),
		78:  uint16(55),
		79:  uint16(55),
		80:  uint16(55),
		81:  uint16(55),
		82:  uint16(55),
		83:  uint16(55),
		84:  uint16(55),
		85:  uint16(55),
		86:  uint16(55),
		87:  uint16(55),
		88:  uint16(55),
		89:  uint16(55),
		114: uint16(57),
		119: uint16(59),
		127: uint16(61),
		128: uint16(59),
		131: uint16(59),
		132: uint16(59),
		145: uint16(3),
		189: uint16(130),
		190: uint16(130),
		191: uint16(93),
		192: uint16(93),
		193: uint16(130),
		194: uint16(130),
		195: uint16(323),
		196: uint16(283),
		197: uint16(283),
		202: uint16(130),
		203: uint16(130),
		204: uint16(130),
		206: uint16(130),
		208: uint16(130),
	},
	16: {
		1:   uint16(33),
		2:   uint16(3),
		8:   uint16(35),
		48:  uint16(39),
		49:  uint16(41),
		50:  uint16(41),
		51:  uint16(43),
		52:  uint16(43),
		53:  uint16(45),
		54:  uint16(45),
		55:  uint16(45),
		56:  uint16(45),
		57:  uint16(45),
		58:  uint16(47),
		59:  uint16(49),
		60:  uint16(45),
		61:  uint16(45),
		62:  uint16(45),
		63:  uint16(45),
		64:  uint16(45),
		65:  uint16(45),
		66:  uint16(45),
		67:  uint16(51),
		68:  uint16(51),
		69:  uint16(51),
		70:  uint16(51),
		71:  uint16(51),
		72:  uint16(51),
		73:  uint16(51),
		74:  uint16(53),
		75:  uint16(53),
		76:  uint16(53),
		77:  uint16(53),
		78:  uint16(55),
		79:  uint16(55),
		80:  uint16(55),
		81:  uint16(55),
		82:  uint16(55),
		83:  uint16(55),
		84:  uint16(55),
		85:  uint16(55),
		86:  uint16(55),
		87:  uint16(55),
		88:  uint16(55),
		89:  uint16(55),
		114: uint16(57),
		119: uint16(59),
		127: uint16(61),
		128: uint16(59),
		131: uint16(59),
		132: uint16(59),
		145: uint16(3),
		189: uint16(112),
		190: uint16(112),
		191: uint16(93),
		192: uint16(93),
		193: uint16(112),
		194: uint16(112),
		195: uint16(323),
		196: uint16(283),
		197: uint16(283),
		202: uint16(112),
		203: uint16(112),
		204: uint16(112),
		206: uint16(112),
		208: uint16(112),
	},
	17: {
		1:   uint16(33),
		2:   uint16(3),
		8:   uint16(35),
		48:  uint16(39),
		49:  uint16(41),
		50:  uint16(41),
		51:  uint16(43),
		52:  uint16(43),
		53:  uint16(45),
		54:  uint16(45),
		55:  uint16(45),
		56:  uint16(45),
		57:  uint16(45),
		58:  uint16(47),
		59:  uint16(49),
		60:  uint16(45),
		61:  uint16(45),
		62:  uint16(45),
		63:  uint16(45),
		64:  uint16(45),
		65:  uint16(45),
		66:  uint16(45),
		67:  uint16(51),
		68:  uint16(51),
		69:  uint16(51),
		70:  uint16(51),
		71:  uint16(51),
		72:  uint16(51),
		73:  uint16(51),
		74:  uint16(53),
		75:  uint16(53),
		76:  uint16(53),
		77:  uint16(53),
		78:  uint16(55),
		79:  uint16(55),
		80:  uint16(55),
		81:  uint16(55),
		82:  uint16(55),
		83:  uint16(55),
		84:  uint16(55),
		85:  uint16(55),
		86:  uint16(55),
		87:  uint16(55),
		88:  uint16(55),
		89:  uint16(55),
		114: uint16(57),
		119: uint16(59),
		127: uint16(61),
		128: uint16(59),
		131: uint16(59),
		132: uint16(59),
		145: uint16(3),
		189: uint16(97),
		190: uint16(97),
		191: uint16(93),
		192: uint16(93),
		193: uint16(97),
		194: uint16(97),
		195: uint16(323),
		196: uint16(283),
		197: uint16(283),
		202: uint16(97),
		203: uint16(97),
		204: uint16(97),
		206: uint16(97),
		208: uint16(97),
	},
	18: {
		1:   uint16(33),
		2:   uint16(3),
		8:   uint16(35),
		48:  uint16(39),
		49:  uint16(41),
		50:  uint16(41),
		51:  uint16(43),
		52:  uint16(43),
		53:  uint16(45),
		54:  uint16(45),
		55:  uint16(45),
		56:  uint16(45),
		57:  uint16(45),
		58:  uint16(47),
		59:  uint16(49),
		60:  uint16(45),
		61:  uint16(45),
		62:  uint16(45),
		63:  uint16(45),
		64:  uint16(45),
		65:  uint16(45),
		66:  uint16(45),
		67:  uint16(51),
		68:  uint16(51),
		69:  uint16(51),
		70:  uint16(51),
		71:  uint16(51),
		72:  uint16(51),
		73:  uint16(51),
		74:  uint16(53),
		75:  uint16(53),
		76:  uint16(53),
		77:  uint16(53),
		78:  uint16(55),
		79:  uint16(55),
		80:  uint16(55),
		81:  uint16(55),
		82:  uint16(55),
		83:  uint16(55),
		84:  uint16(55),
		85:  uint16(55),
		86:  uint16(55),
		87:  uint16(55),
		88:  uint16(55),
		89:  uint16(55),
		114: uint16(57),
		119: uint16(59),
		127: uint16(61),
		128: uint16(59),
		131: uint16(59),
		132: uint16(59),
		145: uint16(3),
		189: uint16(144),
		190: uint16(144),
		191: uint16(93),
		192: uint16(93),
		193: uint16(144),
		194: uint16(144),
		195: uint16(323),
		196: uint16(283),
		197: uint16(283),
		202: uint16(144),
		203: uint16(144),
		204: uint16(144),
		206: uint16(144),
		208: uint16(144),
	},
	19: {
		1:   uint16(33),
		2:   uint16(3),
		8:   uint16(35),
		48:  uint16(39),
		49:  uint16(41),
		50:  uint16(41),
		51:  uint16(43),
		52:  uint16(43),
		53:  uint16(45),
		54:  uint16(45),
		55:  uint16(45),
		56:  uint16(45),
		57:  uint16(45),
		58:  uint16(47),
		59:  uint16(49),
		60:  uint16(45),
		61:  uint16(45),
		62:  uint16(45),
		63:  uint16(45),
		64:  uint16(45),
		65:  uint16(45),
		66:  uint16(45),
		67:  uint16(51),
		68:  uint16(51),
		69:  uint16(51),
		70:  uint16(51),
		71:  uint16(51),
		72:  uint16(51),
		73:  uint16(51),
		74:  uint16(53),
		75:  uint16(53),
		76:  uint16(53),
		77:  uint16(53),
		78:  uint16(55),
		79:  uint16(55),
		80:  uint16(55),
		81:  uint16(55),
		82:  uint16(55),
		83:  uint16(55),
		84:  uint16(55),
		85:  uint16(55),
		86:  uint16(55),
		87:  uint16(55),
		88:  uint16(55),
		89:  uint16(55),
		114: uint16(57),
		119: uint16(59),
		127: uint16(61),
		128: uint16(59),
		131: uint16(59),
		132: uint16(59),
		145: uint16(3),
		189: uint16(148),
		190: uint16(148),
		191: uint16(93),
		192: uint16(93),
		193: uint16(148),
		194: uint16(148),
		195: uint16(323),
		196: uint16(283),
		197: uint16(283),
		202: uint16(148),
		203: uint16(148),
		204: uint16(148),
		206: uint16(148),
		208: uint16(148),
	},
	20: {
		1:   uint16(33),
		2:   uint16(3),
		8:   uint16(35),
		48:  uint16(39),
		49:  uint16(41),
		50:  uint16(41),
		51:  uint16(43),
		52:  uint16(43),
		53:  uint16(45),
		54:  uint16(45),
		55:  uint16(45),
		56:  uint16(45),
		57:  uint16(45),
		58:  uint16(47),
		59:  uint16(49),
		60:  uint16(45),
		61:  uint16(45),
		62:  uint16(45),
		63:  uint16(45),
		64:  uint16(45),
		65:  uint16(45),
		66:  uint16(45),
		67:  uint16(51),
		68:  uint16(51),
		69:  uint16(51),
		70:  uint16(51),
		71:  uint16(51),
		72:  uint16(51),
		73:  uint16(51),
		74:  uint16(53),
		75:  uint16(53),
		76:  uint16(53),
		77:  uint16(53),
		78:  uint16(55),
		79:  uint16(55),
		80:  uint16(55),
		81:  uint16(55),
		82:  uint16(55),
		83:  uint16(55),
		84:  uint16(55),
		85:  uint16(55),
		86:  uint16(55),
		87:  uint16(55),
		88:  uint16(55),
		89:  uint16(55),
		114: uint16(57),
		119: uint16(59),
		127: uint16(61),
		128: uint16(59),
		131: uint16(59),
		132: uint16(59),
		145: uint16(3),
		189: uint16(135),
		190: uint16(135),
		191: uint16(93),
		192: uint16(93),
		193: uint16(135),
		194: uint16(135),
		195: uint16(323),
		196: uint16(283),
		197: uint16(283),
		202: uint16(135),
		203: uint16(135),
		204: uint16(135),
		206: uint16(135),
		208: uint16(135),
	},
	21: {
		1:   uint16(33),
		2:   uint16(3),
		8:   uint16(35),
		48:  uint16(39),
		49:  uint16(41),
		50:  uint16(41),
		51:  uint16(43),
		52:  uint16(43),
		53:  uint16(45),
		54:  uint16(45),
		55:  uint16(45),
		56:  uint16(45),
		57:  uint16(45),
		58:  uint16(47),
		59:  uint16(49),
		60:  uint16(45),
		61:  uint16(45),
		62:  uint16(45),
		63:  uint16(45),
		64:  uint16(45),
		65:  uint16(45),
		66:  uint16(45),
		67:  uint16(51),
		68:  uint16(51),
		69:  uint16(51),
		70:  uint16(51),
		71:  uint16(51),
		72:  uint16(51),
		73:  uint16(51),
		74:  uint16(53),
		75:  uint16(53),
		76:  uint16(53),
		77:  uint16(53),
		78:  uint16(55),
		79:  uint16(55),
		80:  uint16(55),
		81:  uint16(55),
		82:  uint16(55),
		83:  uint16(55),
		84:  uint16(55),
		85:  uint16(55),
		86:  uint16(55),
		87:  uint16(55),
		88:  uint16(55),
		89:  uint16(55),
		114: uint16(57),
		119: uint16(59),
		127: uint16(61),
		128: uint16(59),
		131: uint16(59),
		132: uint16(59),
		145: uint16(3),
		189: uint16(143),
		190: uint16(143),
		191: uint16(93),
		192: uint16(93),
		193: uint16(143),
		194: uint16(143),
		195: uint16(323),
		196: uint16(283),
		197: uint16(283),
		202: uint16(143),
		203: uint16(143),
		204: uint16(143),
		206: uint16(143),
		208: uint16(143),
	},
	22: {
		1:   uint16(33),
		2:   uint16(3),
		8:   uint16(35),
		48:  uint16(39),
		49:  uint16(41),
		50:  uint16(41),
		51:  uint16(43),
		52:  uint16(43),
		53:  uint16(45),
		54:  uint16(45),
		55:  uint16(45),
		56:  uint16(45),
		57:  uint16(45),
		58:  uint16(47),
		59:  uint16(49),
		60:  uint16(45),
		61:  uint16(45),
		62:  uint16(45),
		63:  uint16(45),
		64:  uint16(45),
		65:  uint16(45),
		66:  uint16(45),
		67:  uint16(51),
		68:  uint16(51),
		69:  uint16(51),
		70:  uint16(51),
		71:  uint16(51),
		72:  uint16(51),
		73:  uint16(51),
		74:  uint16(53),
		75:  uint16(53),
		76:  uint16(53),
		77:  uint16(53),
		78:  uint16(55),
		79:  uint16(55),
		80:  uint16(55),
		81:  uint16(55),
		82:  uint16(55),
		83:  uint16(55),
		84:  uint16(55),
		85:  uint16(55),
		86:  uint16(55),
		87:  uint16(55),
		88:  uint16(55),
		89:  uint16(55),
		114: uint16(57),
		119: uint16(59),
		127: uint16(61),
		128: uint16(59),
		131: uint16(59),
		132: uint16(59),
		145: uint16(3),
		189: uint16(134),
		190: uint16(134),
		191: uint16(93),
		192: uint16(93),
		193: uint16(134),
		194: uint16(134),
		195: uint16(323),
		196: uint16(283),
		197: uint16(283),
		202: uint16(134),
		203: uint16(134),
		204: uint16(134),
		206: uint16(134),
		208: uint16(134),
	},
	23: {
		1:   uint16(33),
		2:   uint16(3),
		8:   uint16(35),
		48:  uint16(39),
		49:  uint16(41),
		50:  uint16(41),
		51:  uint16(43),
		52:  uint16(43),
		53:  uint16(45),
		54:  uint16(45),
		55:  uint16(45),
		56:  uint16(45),
		57:  uint16(45),
		58:  uint16(47),
		59:  uint16(49),
		60:  uint16(45),
		61:  uint16(45),
		62:  uint16(45),
		63:  uint16(45),
		64:  uint16(45),
		65:  uint16(45),
		66:  uint16(45),
		67:  uint16(51),
		68:  uint16(51),
		69:  uint16(51),
		70:  uint16(51),
		71:  uint16(51),
		72:  uint16(51),
		73:  uint16(51),
		74:  uint16(53),
		75:  uint16(53),
		76:  uint16(53),
		77:  uint16(53),
		78:  uint16(55),
		79:  uint16(55),
		80:  uint16(55),
		81:  uint16(55),
		82:  uint16(55),
		83:  uint16(55),
		84:  uint16(55),
		85:  uint16(55),
		86:  uint16(55),
		87:  uint16(55),
		88:  uint16(55),
		89:  uint16(55),
		114: uint16(57),
		119: uint16(59),
		127: uint16(61),
		128: uint16(59),
		131: uint16(59),
		132: uint16(59),
		145: uint16(3),
		189: uint16(140),
		190: uint16(140),
		191: uint16(93),
		192: uint16(93),
		193: uint16(140),
		194: uint16(140),
		195: uint16(323),
		196: uint16(283),
		197: uint16(283),
		202: uint16(140),
		203: uint16(140),
		204: uint16(140),
		206: uint16(140),
		208: uint16(140),
	},
	24: {
		1:   uint16(33),
		2:   uint16(3),
		8:   uint16(35),
		48:  uint16(39),
		49:  uint16(41),
		50:  uint16(41),
		51:  uint16(43),
		52:  uint16(43),
		53:  uint16(45),
		54:  uint16(45),
		55:  uint16(45),
		56:  uint16(45),
		57:  uint16(45),
		58:  uint16(47),
		59:  uint16(49),
		60:  uint16(45),
		61:  uint16(45),
		62:  uint16(45),
		63:  uint16(45),
		64:  uint16(45),
		65:  uint16(45),
		66:  uint16(45),
		67:  uint16(51),
		68:  uint16(51),
		69:  uint16(51),
		70:  uint16(51),
		71:  uint16(51),
		72:  uint16(51),
		73:  uint16(51),
		74:  uint16(53),
		75:  uint16(53),
		76:  uint16(53),
		77:  uint16(53),
		78:  uint16(55),
		79:  uint16(55),
		80:  uint16(55),
		81:  uint16(55),
		82:  uint16(55),
		83:  uint16(55),
		84:  uint16(55),
		85:  uint16(55),
		86:  uint16(55),
		87:  uint16(55),
		88:  uint16(55),
		89:  uint16(55),
		114: uint16(57),
		119: uint16(59),
		127: uint16(61),
		128: uint16(59),
		131: uint16(59),
		132: uint16(59),
		145: uint16(3),
		189: uint16(110),
		190: uint16(110),
		191: uint16(93),
		192: uint16(93),
		193: uint16(110),
		194: uint16(110),
		195: uint16(323),
		196: uint16(283),
		197: uint16(283),
		202: uint16(110),
		203: uint16(110),
		204: uint16(110),
		206: uint16(110),
		208: uint16(110),
	},
	25: {
		1:   uint16(33),
		2:   uint16(3),
		8:   uint16(35),
		48:  uint16(39),
		49:  uint16(41),
		50:  uint16(41),
		51:  uint16(43),
		52:  uint16(43),
		53:  uint16(45),
		54:  uint16(45),
		55:  uint16(45),
		56:  uint16(45),
		57:  uint16(45),
		58:  uint16(47),
		59:  uint16(49),
		60:  uint16(45),
		61:  uint16(45),
		62:  uint16(45),
		63:  uint16(45),
		64:  uint16(45),
		65:  uint16(45),
		66:  uint16(45),
		67:  uint16(51),
		68:  uint16(51),
		69:  uint16(51),
		70:  uint16(51),
		71:  uint16(51),
		72:  uint16(51),
		73:  uint16(51),
		74:  uint16(53),
		75:  uint16(53),
		76:  uint16(53),
		77:  uint16(53),
		78:  uint16(55),
		79:  uint16(55),
		80:  uint16(55),
		81:  uint16(55),
		82:  uint16(55),
		83:  uint16(55),
		84:  uint16(55),
		85:  uint16(55),
		86:  uint16(55),
		87:  uint16(55),
		88:  uint16(55),
		89:  uint16(55),
		114: uint16(57),
		119: uint16(59),
		127: uint16(61),
		128: uint16(59),
		131: uint16(59),
		132: uint16(59),
		145: uint16(3),
		189: uint16(137),
		190: uint16(137),
		191: uint16(93),
		192: uint16(93),
		193: uint16(137),
		194: uint16(137),
		195: uint16(323),
		196: uint16(283),
		197: uint16(283),
		202: uint16(137),
		203: uint16(137),
		204: uint16(137),
		206: uint16(137),
		208: uint16(137),
	},
	26: {
		1:   uint16(33),
		2:   uint16(3),
		8:   uint16(35),
		48:  uint16(39),
		49:  uint16(41),
		50:  uint16(41),
		51:  uint16(43),
		52:  uint16(43),
		53:  uint16(45),
		54:  uint16(45),
		55:  uint16(45),
		56:  uint16(45),
		57:  uint16(45),
		58:  uint16(47),
		59:  uint16(49),
		60:  uint16(45),
		61:  uint16(45),
		62:  uint16(45),
		63:  uint16(45),
		64:  uint16(45),
		65:  uint16(45),
		66:  uint16(45),
		67:  uint16(51),
		68:  uint16(51),
		69:  uint16(51),
		70:  uint16(51),
		71:  uint16(51),
		72:  uint16(51),
		73:  uint16(51),
		74:  uint16(53),
		75:  uint16(53),
		76:  uint16(53),
		77:  uint16(53),
		78:  uint16(55),
		79:  uint16(55),
		80:  uint16(55),
		81:  uint16(55),
		82:  uint16(55),
		83:  uint16(55),
		84:  uint16(55),
		85:  uint16(55),
		86:  uint16(55),
		87:  uint16(55),
		88:  uint16(55),
		89:  uint16(55),
		114: uint16(57),
		119: uint16(59),
		127: uint16(61),
		128: uint16(59),
		131: uint16(59),
		132: uint16(59),
		145: uint16(3),
		189: uint16(133),
		190: uint16(133),
		191: uint16(93),
		192: uint16(93),
		193: uint16(133),
		194: uint16(133),
		195: uint16(323),
		196: uint16(283),
		197: uint16(283),
		202: uint16(133),
		203: uint16(133),
		204: uint16(133),
		206: uint16(133),
		208: uint16(133),
	},
	27: {
		1:   uint16(33),
		2:   uint16(3),
		8:   uint16(35),
		48:  uint16(39),
		49:  uint16(41),
		50:  uint16(41),
		51:  uint16(43),
		52:  uint16(43),
		53:  uint16(45),
		54:  uint16(45),
		55:  uint16(45),
		56:  uint16(45),
		57:  uint16(45),
		58:  uint16(47),
		59:  uint16(49),
		60:  uint16(45),
		61:  uint16(45),
		62:  uint16(45),
		63:  uint16(45),
		64:  uint16(45),
		65:  uint16(45),
		66:  uint16(45),
		67:  uint16(51),
		68:  uint16(51),
		69:  uint16(51),
		70:  uint16(51),
		71:  uint16(51),
		72:  uint16(51),
		73:  uint16(51),
		74:  uint16(53),
		75:  uint16(53),
		76:  uint16(53),
		77:  uint16(53),
		78:  uint16(55),
		79:  uint16(55),
		80:  uint16(55),
		81:  uint16(55),
		82:  uint16(55),
		83:  uint16(55),
		84:  uint16(55),
		85:  uint16(55),
		86:  uint16(55),
		87:  uint16(55),
		88:  uint16(55),
		89:  uint16(55),
		114: uint16(57),
		119: uint16(59),
		127: uint16(61),
		128: uint16(59),
		131: uint16(59),
		132: uint16(59),
		145: uint16(3),
		189: uint16(107),
		190: uint16(107),
		191: uint16(93),
		192: uint16(93),
		193: uint16(107),
		194: uint16(107),
		195: uint16(323),
		196: uint16(283),
		197: uint16(283),
		202: uint16(107),
		203: uint16(107),
		204: uint16(107),
		206: uint16(107),
		208: uint16(107),
	},
	28: {
		1:   uint16(33),
		2:   uint16(3),
		8:   uint16(35),
		48:  uint16(39),
		49:  uint16(41),
		50:  uint16(41),
		51:  uint16(43),
		52:  uint16(43),
		53:  uint16(45),
		54:  uint16(45),
		55:  uint16(45),
		56:  uint16(45),
		57:  uint16(45),
		58:  uint16(47),
		59:  uint16(49),
		60:  uint16(45),
		61:  uint16(45),
		62:  uint16(45),
		63:  uint16(45),
		64:  uint16(45),
		65:  uint16(45),
		66:  uint16(45),
		67:  uint16(51),
		68:  uint16(51),
		69:  uint16(51),
		70:  uint16(51),
		71:  uint16(51),
		72:  uint16(51),
		73:  uint16(51),
		74:  uint16(53),
		75:  uint16(53),
		76:  uint16(53),
		77:  uint16(53),
		78:  uint16(55),
		79:  uint16(55),
		80:  uint16(55),
		81:  uint16(55),
		82:  uint16(55),
		83:  uint16(55),
		84:  uint16(55),
		85:  uint16(55),
		86:  uint16(55),
		87:  uint16(55),
		88:  uint16(55),
		89:  uint16(55),
		114: uint16(57),
		119: uint16(59),
		127: uint16(61),
		128: uint16(59),
		131: uint16(59),
		132: uint16(59),
		145: uint16(3),
		189: uint16(139),
		190: uint16(139),
		191: uint16(93),
		192: uint16(93),
		193: uint16(139),
		194: uint16(139),
		195: uint16(323),
		196: uint16(283),
		197: uint16(283),
		202: uint16(139),
		203: uint16(139),
		204: uint16(139),
		206: uint16(139),
		208: uint16(139),
	},
	29: {
		1:   uint16(33),
		2:   uint16(3),
		8:   uint16(35),
		48:  uint16(39),
		49:  uint16(41),
		50:  uint16(41),
		51:  uint16(43),
		52:  uint16(43),
		53:  uint16(45),
		54:  uint16(45),
		55:  uint16(45),
		56:  uint16(45),
		57:  uint16(45),
		58:  uint16(47),
		59:  uint16(49),
		60:  uint16(45),
		61:  uint16(45),
		62:  uint16(45),
		63:  uint16(45),
		64:  uint16(45),
		65:  uint16(45),
		66:  uint16(45),
		67:  uint16(51),
		68:  uint16(51),
		69:  uint16(51),
		70:  uint16(51),
		71:  uint16(51),
		72:  uint16(51),
		73:  uint16(51),
		74:  uint16(53),
		75:  uint16(53),
		76:  uint16(53),
		77:  uint16(53),
		78:  uint16(55),
		79:  uint16(55),
		80:  uint16(55),
		81:  uint16(55),
		82:  uint16(55),
		83:  uint16(55),
		84:  uint16(55),
		85:  uint16(55),
		86:  uint16(55),
		87:  uint16(55),
		88:  uint16(55),
		89:  uint16(55),
		114: uint16(57),
		119: uint16(59),
		127: uint16(61),
		128: uint16(59),
		131: uint16(59),
		132: uint16(59),
		145: uint16(3),
		189: uint16(138),
		190: uint16(138),
		191: uint16(93),
		192: uint16(93),
		193: uint16(138),
		194: uint16(138),
		195: uint16(323),
		196: uint16(283),
		197: uint16(283),
		202: uint16(138),
		203: uint16(138),
		204: uint16(138),
		206: uint16(138),
		208: uint16(138),
	},
	30: {
		1:   uint16(33),
		2:   uint16(3),
		8:   uint16(35),
		48:  uint16(39),
		49:  uint16(41),
		50:  uint16(41),
		51:  uint16(43),
		52:  uint16(43),
		53:  uint16(45),
		54:  uint16(45),
		55:  uint16(45),
		56:  uint16(45),
		57:  uint16(45),
		58:  uint16(47),
		59:  uint16(49),
		60:  uint16(45),
		61:  uint16(45),
		62:  uint16(45),
		63:  uint16(45),
		64:  uint16(45),
		65:  uint16(45),
		66:  uint16(45),
		67:  uint16(51),
		68:  uint16(51),
		69:  uint16(51),
		70:  uint16(51),
		71:  uint16(51),
		72:  uint16(51),
		73:  uint16(51),
		74:  uint16(53),
		75:  uint16(53),
		76:  uint16(53),
		77:  uint16(53),
		78:  uint16(55),
		79:  uint16(55),
		80:  uint16(55),
		81:  uint16(55),
		82:  uint16(55),
		83:  uint16(55),
		84:  uint16(55),
		85:  uint16(55),
		86:  uint16(55),
		87:  uint16(55),
		88:  uint16(55),
		89:  uint16(55),
		114: uint16(57),
		119: uint16(59),
		127: uint16(61),
		128: uint16(59),
		131: uint16(59),
		132: uint16(59),
		145: uint16(3),
		189: uint16(146),
		190: uint16(146),
		191: uint16(93),
		192: uint16(93),
		193: uint16(146),
		194: uint16(146),
		195: uint16(323),
		196: uint16(283),
		197: uint16(283),
		202: uint16(146),
		203: uint16(146),
		204: uint16(146),
		206: uint16(146),
		208: uint16(146),
	},
	31: {
		1:   uint16(33),
		2:   uint16(3),
		8:   uint16(35),
		48:  uint16(39),
		49:  uint16(41),
		50:  uint16(41),
		51:  uint16(43),
		52:  uint16(43),
		53:  uint16(45),
		54:  uint16(45),
		55:  uint16(45),
		56:  uint16(45),
		57:  uint16(45),
		58:  uint16(47),
		59:  uint16(49),
		60:  uint16(45),
		61:  uint16(45),
		62:  uint16(45),
		63:  uint16(45),
		64:  uint16(45),
		65:  uint16(45),
		66:  uint16(45),
		67:  uint16(51),
		68:  uint16(51),
		69:  uint16(51),
		70:  uint16(51),
		71:  uint16(51),
		72:  uint16(51),
		73:  uint16(51),
		74:  uint16(53),
		75:  uint16(53),
		76:  uint16(53),
		77:  uint16(53),
		78:  uint16(55),
		79:  uint16(55),
		80:  uint16(55),
		81:  uint16(55),
		82:  uint16(55),
		83:  uint16(55),
		84:  uint16(55),
		85:  uint16(55),
		86:  uint16(55),
		87:  uint16(55),
		88:  uint16(55),
		89:  uint16(55),
		114: uint16(57),
		119: uint16(59),
		127: uint16(61),
		128: uint16(59),
		131: uint16(59),
		132: uint16(59),
		145: uint16(3),
		189: uint16(149),
		190: uint16(149),
		191: uint16(93),
		192: uint16(93),
		193: uint16(149),
		194: uint16(149),
		195: uint16(323),
		196: uint16(283),
		197: uint16(283),
		202: uint16(149),
		203: uint16(149),
		204: uint16(149),
		206: uint16(149),
		208: uint16(149),
	},
}

var ts_small_parse_table = [11077]uint16_t{
	0:     uint16(21),
	1:     uint16(47),
	2:     uint16(1),
	3:     uint16(anon_sym_array),
	4:     uint16(49),
	5:     uint16(1),
	6:     uint16(anon_sym_ptr),
	7:     uint16(111),
	8:     uint16(1),
	9:     uint16(sym_identifier),
	10:    uint16(113),
	11:    uint16(1),
	12:    uint16(anon_sym_SEMI),
	13:    uint16(115),
	14:    uint16(1),
	15:    uint16(anon_sym_let),
	16:    uint16(117),
	17:    uint16(1),
	18:    uint16(anon_sym_LPAREN),
	19:    uint16(119),
	20:    uint16(1),
	21:    uint16(anon_sym__),
	22:    uint16(121),
	23:    uint16(1),
	24:    uint16(anon_sym_var),
	25:    uint16(201),
	26:    uint16(1),
	27:    uint16(sym_lhs_expression),
	28:    uint16(239),
	29:    uint16(1),
	30:    uint16(aux_sym_lhs_expression_repeat1),
	31:    uint16(323),
	32:    uint16(1),
	33:    uint16(sym_type_declaration),
	34:    uint16(330),
	35:    uint16(1),
	36:    uint16(sym_variable_declaration),
	37:    uint16(433),
	38:    uint16(1),
	39:    uint16(sym_for_header),
	40:    uint16(3),
	41:    uint16(2),
	42:    uint16(sym_block_comment),
	43:    uint16(sym_line_comment),
	44:    uint16(123),
	45:    uint16(2),
	46:    uint16(anon_sym_AMP),
	47:    uint16(anon_sym_STAR),
	48:    uint16(283),
	49:    uint16(2),
	50:    uint16(sym__vec_prefix),
	51:    uint16(sym__mat_prefix),
	52:    uint16(53),
	53:    uint16(4),
	54:    uint16(anon_sym_texture_storage_1d),
	55:    uint16(anon_sym_texture_storage_2d),
	56:    uint16(anon_sym_texture_storage_2d_array),
	57:    uint16(anon_sym_texture_storage_3d),
	58:    uint16(432),
	59:    uint16(5),
	60:    uint16(sym_assignment_statement),
	61:    uint16(sym_variable_statement),
	62:    uint16(sym_increment_statement),
	63:    uint16(sym_decrement_statement),
	64:    uint16(sym_type_constructor_or_function_call_expression),
	65:    uint16(51),
	66:    uint16(7),
	67:    uint16(anon_sym_texture_1d),
	68:    uint16(anon_sym_texture_2d),
	69:    uint16(anon_sym_texture_2d_array),
	70:    uint16(anon_sym_texture_3d),
	71:    uint16(anon_sym_texture_cube),
	72:    uint16(anon_sym_texture_cube_array),
	73:    uint16(anon_sym_texture_multisampled_2d),
	74:    uint16(45),
	75:    uint16(12),
	76:    uint16(anon_sym_bool),
	77:    uint16(anon_sym_u32),
	78:    uint16(anon_sym_i32),
	79:    uint16(anon_sym_f32),
	80:    uint16(anon_sym_f16),
	81:    uint16(anon_sym_sampler),
	82:    uint16(anon_sym_sampler_comparison),
	83:    uint16(anon_sym_texture_depth_2d),
	84:    uint16(anon_sym_texture_depth_2d_array),
	85:    uint16(anon_sym_texture_depth_cube),
	86:    uint16(anon_sym_texture_depth_cube_array),
	87:    uint16(anon_sym_texture_depth_multisampled_2d),
	88:    uint16(55),
	89:    uint16(12),
	90:    uint16(anon_sym_vec2),
	91:    uint16(anon_sym_vec3),
	92:    uint16(anon_sym_vec4),
	93:    uint16(anon_sym_mat2x2),
	94:    uint16(anon_sym_mat2x3),
	95:    uint16(anon_sym_mat2x4),
	96:    uint16(anon_sym_mat3x2),
	97:    uint16(anon_sym_mat3x3),
	98:    uint16(anon_sym_mat3x4),
	99:    uint16(anon_sym_mat4x2),
	100:   uint16(anon_sym_mat4x3),
	101:   uint16(anon_sym_mat4x4),
	102:   uint16(17),
	103:   uint16(39),
	104:   uint16(1),
	105:   uint16(sym_int_literal),
	106:   uint16(47),
	107:   uint16(1),
	108:   uint16(anon_sym_array),
	109:   uint16(49),
	110:   uint16(1),
	111:   uint16(anon_sym_ptr),
	112:   uint16(125),
	113:   uint16(1),
	114:   uint16(anon_sym_RPAREN),
	115:   uint16(38),
	116:   uint16(1),
	117:   uint16(aux_sym_const_expression_repeat1),
	118:   uint16(274),
	119:   uint16(1),
	120:   uint16(sym_const_literal),
	121:   uint16(303),
	122:   uint16(1),
	123:   uint16(sym_const_expression),
	124:   uint16(374),
	125:   uint16(1),
	126:   uint16(sym_type_declaration),
	127:   uint16(3),
	128:   uint16(2),
	129:   uint16(sym_block_comment),
	130:   uint16(sym_line_comment),
	131:   uint16(41),
	132:   uint16(2),
	133:   uint16(aux_sym_float_literal_token1),
	134:   uint16(aux_sym_float_literal_token2),
	135:   uint16(43),
	136:   uint16(2),
	137:   uint16(anon_sym_true),
	138:   uint16(anon_sym_false),
	139:   uint16(93),
	140:   uint16(2),
	141:   uint16(sym_float_literal),
	142:   uint16(sym_bool_literal),
	143:   uint16(376),
	144:   uint16(2),
	145:   uint16(sym__vec_prefix),
	146:   uint16(sym__mat_prefix),
	147:   uint16(53),
	148:   uint16(4),
	149:   uint16(anon_sym_texture_storage_1d),
	150:   uint16(anon_sym_texture_storage_2d),
	151:   uint16(anon_sym_texture_storage_2d_array),
	152:   uint16(anon_sym_texture_storage_3d),
	153:   uint16(51),
	154:   uint16(7),
	155:   uint16(anon_sym_texture_1d),
	156:   uint16(anon_sym_texture_2d),
	157:   uint16(anon_sym_texture_2d_array),
	158:   uint16(anon_sym_texture_3d),
	159:   uint16(anon_sym_texture_cube),
	160:   uint16(anon_sym_texture_cube_array),
	161:   uint16(anon_sym_texture_multisampled_2d),
	162:   uint16(127),
	163:   uint16(12),
	164:   uint16(anon_sym_vec2),
	165:   uint16(anon_sym_vec3),
	166:   uint16(anon_sym_vec4),
	167:   uint16(anon_sym_mat2x2),
	168:   uint16(anon_sym_mat2x3),
	169:   uint16(anon_sym_mat2x4),
	170:   uint16(anon_sym_mat3x2),
	171:   uint16(anon_sym_mat3x3),
	172:   uint16(anon_sym_mat3x4),
	173:   uint16(anon_sym_mat4x2),
	174:   uint16(anon_sym_mat4x3),
	175:   uint16(anon_sym_mat4x4),
	176:   uint16(45),
	177:   uint16(13),
	178:   uint16(sym_identifier),
	179:   uint16(anon_sym_bool),
	180:   uint16(anon_sym_u32),
	181:   uint16(anon_sym_i32),
	182:   uint16(anon_sym_f32),
	183:   uint16(anon_sym_f16),
	184:   uint16(anon_sym_sampler),
	185:   uint16(anon_sym_sampler_comparison),
	186:   uint16(anon_sym_texture_depth_2d),
	187:   uint16(anon_sym_texture_depth_2d_array),
	188:   uint16(anon_sym_texture_depth_cube),
	189:   uint16(anon_sym_texture_depth_cube_array),
	190:   uint16(anon_sym_texture_depth_multisampled_2d),
	191:   uint16(17),
	192:   uint16(47),
	193:   uint16(1),
	194:   uint16(anon_sym_array),
	195:   uint16(49),
	196:   uint16(1),
	197:   uint16(anon_sym_ptr),
	198:   uint16(111),
	199:   uint16(1),
	200:   uint16(sym_identifier),
	201:   uint16(117),
	202:   uint16(1),
	203:   uint16(anon_sym_LPAREN),
	204:   uint16(119),
	205:   uint16(1),
	206:   uint16(anon_sym__),
	207:   uint16(129),
	208:   uint16(1),
	209:   uint16(anon_sym_RPAREN),
	210:   uint16(201),
	211:   uint16(1),
	212:   uint16(sym_lhs_expression),
	213:   uint16(239),
	214:   uint16(1),
	215:   uint16(aux_sym_lhs_expression_repeat1),
	216:   uint16(323),
	217:   uint16(1),
	218:   uint16(sym_type_declaration),
	219:   uint16(3),
	220:   uint16(2),
	221:   uint16(sym_block_comment),
	222:   uint16(sym_line_comment),
	223:   uint16(123),
	224:   uint16(2),
	225:   uint16(anon_sym_AMP),
	226:   uint16(anon_sym_STAR),
	227:   uint16(283),
	228:   uint16(2),
	229:   uint16(sym__vec_prefix),
	230:   uint16(sym__mat_prefix),
	231:   uint16(53),
	232:   uint16(4),
	233:   uint16(anon_sym_texture_storage_1d),
	234:   uint16(anon_sym_texture_storage_2d),
	235:   uint16(anon_sym_texture_storage_2d_array),
	236:   uint16(anon_sym_texture_storage_3d),
	237:   uint16(396),
	238:   uint16(4),
	239:   uint16(sym_assignment_statement),
	240:   uint16(sym_increment_statement),
	241:   uint16(sym_decrement_statement),
	242:   uint16(sym_type_constructor_or_function_call_expression),
	243:   uint16(51),
	244:   uint16(7),
	245:   uint16(anon_sym_texture_1d),
	246:   uint16(anon_sym_texture_2d),
	247:   uint16(anon_sym_texture_2d_array),
	248:   uint16(anon_sym_texture_3d),
	249:   uint16(anon_sym_texture_cube),
	250:   uint16(anon_sym_texture_cube_array),
	251:   uint16(anon_sym_texture_multisampled_2d),
	252:   uint16(45),
	253:   uint16(12),
	254:   uint16(anon_sym_bool),
	255:   uint16(anon_sym_u32),
	256:   uint16(anon_sym_i32),
	257:   uint16(anon_sym_f32),
	258:   uint16(anon_sym_f16),
	259:   uint16(anon_sym_sampler),
	260:   uint16(anon_sym_sampler_comparison),
	261:   uint16(anon_sym_texture_depth_2d),
	262:   uint16(anon_sym_texture_depth_2d_array),
	263:   uint16(anon_sym_texture_depth_cube),
	264:   uint16(anon_sym_texture_depth_cube_array),
	265:   uint16(anon_sym_texture_depth_multisampled_2d),
	266:   uint16(55),
	267:   uint16(12),
	268:   uint16(anon_sym_vec2),
	269:   uint16(anon_sym_vec3),
	270:   uint16(anon_sym_vec4),
	271:   uint16(anon_sym_mat2x2),
	272:   uint16(anon_sym_mat2x3),
	273:   uint16(anon_sym_mat2x4),
	274:   uint16(anon_sym_mat3x2),
	275:   uint16(anon_sym_mat3x3),
	276:   uint16(anon_sym_mat3x4),
	277:   uint16(anon_sym_mat4x2),
	278:   uint16(anon_sym_mat4x3),
	279:   uint16(anon_sym_mat4x4),
	280:   uint16(17),
	281:   uint16(47),
	282:   uint16(1),
	283:   uint16(anon_sym_array),
	284:   uint16(49),
	285:   uint16(1),
	286:   uint16(anon_sym_ptr),
	287:   uint16(111),
	288:   uint16(1),
	289:   uint16(sym_identifier),
	290:   uint16(117),
	291:   uint16(1),
	292:   uint16(anon_sym_LPAREN),
	293:   uint16(119),
	294:   uint16(1),
	295:   uint16(anon_sym__),
	296:   uint16(131),
	297:   uint16(1),
	298:   uint16(anon_sym_RPAREN),
	299:   uint16(201),
	300:   uint16(1),
	301:   uint16(sym_lhs_expression),
	302:   uint16(239),
	303:   uint16(1),
	304:   uint16(aux_sym_lhs_expression_repeat1),
	305:   uint16(323),
	306:   uint16(1),
	307:   uint16(sym_type_declaration),
	308:   uint16(3),
	309:   uint16(2),
	310:   uint16(sym_block_comment),
	311:   uint16(sym_line_comment),
	312:   uint16(123),
	313:   uint16(2),
	314:   uint16(anon_sym_AMP),
	315:   uint16(anon_sym_STAR),
	316:   uint16(283),
	317:   uint16(2),
	318:   uint16(sym__vec_prefix),
	319:   uint16(sym__mat_prefix),
	320:   uint16(53),
	321:   uint16(4),
	322:   uint16(anon_sym_texture_storage_1d),
	323:   uint16(anon_sym_texture_storage_2d),
	324:   uint16(anon_sym_texture_storage_2d_array),
	325:   uint16(anon_sym_texture_storage_3d),
	326:   uint16(415),
	327:   uint16(4),
	328:   uint16(sym_assignment_statement),
	329:   uint16(sym_increment_statement),
	330:   uint16(sym_decrement_statement),
	331:   uint16(sym_type_constructor_or_function_call_expression),
	332:   uint16(51),
	333:   uint16(7),
	334:   uint16(anon_sym_texture_1d),
	335:   uint16(anon_sym_texture_2d),
	336:   uint16(anon_sym_texture_2d_array),
	337:   uint16(anon_sym_texture_3d),
	338:   uint16(anon_sym_texture_cube),
	339:   uint16(anon_sym_texture_cube_array),
	340:   uint16(anon_sym_texture_multisampled_2d),
	341:   uint16(45),
	342:   uint16(12),
	343:   uint16(anon_sym_bool),
	344:   uint16(anon_sym_u32),
	345:   uint16(anon_sym_i32),
	346:   uint16(anon_sym_f32),
	347:   uint16(anon_sym_f16),
	348:   uint16(anon_sym_sampler),
	349:   uint16(anon_sym_sampler_comparison),
	350:   uint16(anon_sym_texture_depth_2d),
	351:   uint16(anon_sym_texture_depth_2d_array),
	352:   uint16(anon_sym_texture_depth_cube),
	353:   uint16(anon_sym_texture_depth_cube_array),
	354:   uint16(anon_sym_texture_depth_multisampled_2d),
	355:   uint16(55),
	356:   uint16(12),
	357:   uint16(anon_sym_vec2),
	358:   uint16(anon_sym_vec3),
	359:   uint16(anon_sym_vec4),
	360:   uint16(anon_sym_mat2x2),
	361:   uint16(anon_sym_mat2x3),
	362:   uint16(anon_sym_mat2x4),
	363:   uint16(anon_sym_mat3x2),
	364:   uint16(anon_sym_mat3x3),
	365:   uint16(anon_sym_mat3x4),
	366:   uint16(anon_sym_mat4x2),
	367:   uint16(anon_sym_mat4x3),
	368:   uint16(anon_sym_mat4x4),
	369:   uint16(17),
	370:   uint16(47),
	371:   uint16(1),
	372:   uint16(anon_sym_array),
	373:   uint16(49),
	374:   uint16(1),
	375:   uint16(anon_sym_ptr),
	376:   uint16(111),
	377:   uint16(1),
	378:   uint16(sym_identifier),
	379:   uint16(117),
	380:   uint16(1),
	381:   uint16(anon_sym_LPAREN),
	382:   uint16(119),
	383:   uint16(1),
	384:   uint16(anon_sym__),
	385:   uint16(133),
	386:   uint16(1),
	387:   uint16(anon_sym_RPAREN),
	388:   uint16(201),
	389:   uint16(1),
	390:   uint16(sym_lhs_expression),
	391:   uint16(239),
	392:   uint16(1),
	393:   uint16(aux_sym_lhs_expression_repeat1),
	394:   uint16(323),
	395:   uint16(1),
	396:   uint16(sym_type_declaration),
	397:   uint16(3),
	398:   uint16(2),
	399:   uint16(sym_block_comment),
	400:   uint16(sym_line_comment),
	401:   uint16(123),
	402:   uint16(2),
	403:   uint16(anon_sym_AMP),
	404:   uint16(anon_sym_STAR),
	405:   uint16(283),
	406:   uint16(2),
	407:   uint16(sym__vec_prefix),
	408:   uint16(sym__mat_prefix),
	409:   uint16(53),
	410:   uint16(4),
	411:   uint16(anon_sym_texture_storage_1d),
	412:   uint16(anon_sym_texture_storage_2d),
	413:   uint16(anon_sym_texture_storage_2d_array),
	414:   uint16(anon_sym_texture_storage_3d),
	415:   uint16(426),
	416:   uint16(4),
	417:   uint16(sym_assignment_statement),
	418:   uint16(sym_increment_statement),
	419:   uint16(sym_decrement_statement),
	420:   uint16(sym_type_constructor_or_function_call_expression),
	421:   uint16(51),
	422:   uint16(7),
	423:   uint16(anon_sym_texture_1d),
	424:   uint16(anon_sym_texture_2d),
	425:   uint16(anon_sym_texture_2d_array),
	426:   uint16(anon_sym_texture_3d),
	427:   uint16(anon_sym_texture_cube),
	428:   uint16(anon_sym_texture_cube_array),
	429:   uint16(anon_sym_texture_multisampled_2d),
	430:   uint16(45),
	431:   uint16(12),
	432:   uint16(anon_sym_bool),
	433:   uint16(anon_sym_u32),
	434:   uint16(anon_sym_i32),
	435:   uint16(anon_sym_f32),
	436:   uint16(anon_sym_f16),
	437:   uint16(anon_sym_sampler),
	438:   uint16(anon_sym_sampler_comparison),
	439:   uint16(anon_sym_texture_depth_2d),
	440:   uint16(anon_sym_texture_depth_2d_array),
	441:   uint16(anon_sym_texture_depth_cube),
	442:   uint16(anon_sym_texture_depth_cube_array),
	443:   uint16(anon_sym_texture_depth_multisampled_2d),
	444:   uint16(55),
	445:   uint16(12),
	446:   uint16(anon_sym_vec2),
	447:   uint16(anon_sym_vec3),
	448:   uint16(anon_sym_vec4),
	449:   uint16(anon_sym_mat2x2),
	450:   uint16(anon_sym_mat2x3),
	451:   uint16(anon_sym_mat2x4),
	452:   uint16(anon_sym_mat3x2),
	453:   uint16(anon_sym_mat3x3),
	454:   uint16(anon_sym_mat3x4),
	455:   uint16(anon_sym_mat4x2),
	456:   uint16(anon_sym_mat4x3),
	457:   uint16(anon_sym_mat4x4),
	458:   uint16(4),
	459:   uint16(139),
	460:   uint16(1),
	461:   uint16(anon_sym_RPAREN),
	462:   uint16(3),
	463:   uint16(2),
	464:   uint16(sym_block_comment),
	465:   uint16(sym_line_comment),
	466:   uint16(137),
	467:   uint16(7),
	468:   uint16(anon_sym_LPAREN),
	469:   uint16(aux_sym_float_literal_token1),
	470:   uint16(aux_sym_float_literal_token2),
	471:   uint16(anon_sym_AMP),
	472:   uint16(anon_sym_STAR),
	473:   uint16(anon_sym_BANG),
	474:   uint16(anon_sym_TILDE),
	475:   uint16(135),
	476:   uint16(43),
	477:   uint16(sym_identifier),
	478:   uint16(sym_int_literal),
	479:   uint16(anon_sym_true),
	480:   uint16(anon_sym_false),
	481:   uint16(anon_sym_bool),
	482:   uint16(anon_sym_u32),
	483:   uint16(anon_sym_i32),
	484:   uint16(anon_sym_f32),
	485:   uint16(anon_sym_f16),
	486:   uint16(anon_sym_array),
	487:   uint16(anon_sym_ptr),
	488:   uint16(anon_sym_sampler),
	489:   uint16(anon_sym_sampler_comparison),
	490:   uint16(anon_sym_texture_depth_2d),
	491:   uint16(anon_sym_texture_depth_2d_array),
	492:   uint16(anon_sym_texture_depth_cube),
	493:   uint16(anon_sym_texture_depth_cube_array),
	494:   uint16(anon_sym_texture_depth_multisampled_2d),
	495:   uint16(anon_sym_texture_1d),
	496:   uint16(anon_sym_texture_2d),
	497:   uint16(anon_sym_texture_2d_array),
	498:   uint16(anon_sym_texture_3d),
	499:   uint16(anon_sym_texture_cube),
	500:   uint16(anon_sym_texture_cube_array),
	501:   uint16(anon_sym_texture_multisampled_2d),
	502:   uint16(anon_sym_texture_storage_1d),
	503:   uint16(anon_sym_texture_storage_2d),
	504:   uint16(anon_sym_texture_storage_2d_array),
	505:   uint16(anon_sym_texture_storage_3d),
	506:   uint16(anon_sym_vec2),
	507:   uint16(anon_sym_vec3),
	508:   uint16(anon_sym_vec4),
	509:   uint16(anon_sym_mat2x2),
	510:   uint16(anon_sym_mat2x3),
	511:   uint16(anon_sym_mat2x4),
	512:   uint16(anon_sym_mat3x2),
	513:   uint16(anon_sym_mat3x3),
	514:   uint16(anon_sym_mat3x4),
	515:   uint16(anon_sym_mat4x2),
	516:   uint16(anon_sym_mat4x3),
	517:   uint16(anon_sym_mat4x4),
	518:   uint16(anon_sym_bitcast),
	519:   uint16(anon_sym_DASH),
	520:   uint16(16),
	521:   uint16(39),
	522:   uint16(1),
	523:   uint16(sym_int_literal),
	524:   uint16(47),
	525:   uint16(1),
	526:   uint16(anon_sym_array),
	527:   uint16(49),
	528:   uint16(1),
	529:   uint16(anon_sym_ptr),
	530:   uint16(39),
	531:   uint16(1),
	532:   uint16(aux_sym_const_expression_repeat1),
	533:   uint16(274),
	534:   uint16(1),
	535:   uint16(sym_const_literal),
	536:   uint16(305),
	537:   uint16(1),
	538:   uint16(sym_const_expression),
	539:   uint16(374),
	540:   uint16(1),
	541:   uint16(sym_type_declaration),
	542:   uint16(3),
	543:   uint16(2),
	544:   uint16(sym_block_comment),
	545:   uint16(sym_line_comment),
	546:   uint16(41),
	547:   uint16(2),
	548:   uint16(aux_sym_float_literal_token1),
	549:   uint16(aux_sym_float_literal_token2),
	550:   uint16(43),
	551:   uint16(2),
	552:   uint16(anon_sym_true),
	553:   uint16(anon_sym_false),
	554:   uint16(93),
	555:   uint16(2),
	556:   uint16(sym_float_literal),
	557:   uint16(sym_bool_literal),
	558:   uint16(376),
	559:   uint16(2),
	560:   uint16(sym__vec_prefix),
	561:   uint16(sym__mat_prefix),
	562:   uint16(53),
	563:   uint16(4),
	564:   uint16(anon_sym_texture_storage_1d),
	565:   uint16(anon_sym_texture_storage_2d),
	566:   uint16(anon_sym_texture_storage_2d_array),
	567:   uint16(anon_sym_texture_storage_3d),
	568:   uint16(51),
	569:   uint16(7),
	570:   uint16(anon_sym_texture_1d),
	571:   uint16(anon_sym_texture_2d),
	572:   uint16(anon_sym_texture_2d_array),
	573:   uint16(anon_sym_texture_3d),
	574:   uint16(anon_sym_texture_cube),
	575:   uint16(anon_sym_texture_cube_array),
	576:   uint16(anon_sym_texture_multisampled_2d),
	577:   uint16(127),
	578:   uint16(12),
	579:   uint16(anon_sym_vec2),
	580:   uint16(anon_sym_vec3),
	581:   uint16(anon_sym_vec4),
	582:   uint16(anon_sym_mat2x2),
	583:   uint16(anon_sym_mat2x3),
	584:   uint16(anon_sym_mat2x4),
	585:   uint16(anon_sym_mat3x2),
	586:   uint16(anon_sym_mat3x3),
	587:   uint16(anon_sym_mat3x4),
	588:   uint16(anon_sym_mat4x2),
	589:   uint16(anon_sym_mat4x3),
	590:   uint16(anon_sym_mat4x4),
	591:   uint16(45),
	592:   uint16(13),
	593:   uint16(sym_identifier),
	594:   uint16(anon_sym_bool),
	595:   uint16(anon_sym_u32),
	596:   uint16(anon_sym_i32),
	597:   uint16(anon_sym_f32),
	598:   uint16(anon_sym_f16),
	599:   uint16(anon_sym_sampler),
	600:   uint16(anon_sym_sampler_comparison),
	601:   uint16(anon_sym_texture_depth_2d),
	602:   uint16(anon_sym_texture_depth_2d_array),
	603:   uint16(anon_sym_texture_depth_cube),
	604:   uint16(anon_sym_texture_depth_cube_array),
	605:   uint16(anon_sym_texture_depth_multisampled_2d),
	606:   uint16(16),
	607:   uint16(144),
	608:   uint16(1),
	609:   uint16(sym_int_literal),
	610:   uint16(153),
	611:   uint16(1),
	612:   uint16(anon_sym_array),
	613:   uint16(156),
	614:   uint16(1),
	615:   uint16(anon_sym_ptr),
	616:   uint16(39),
	617:   uint16(1),
	618:   uint16(aux_sym_const_expression_repeat1),
	619:   uint16(274),
	620:   uint16(1),
	621:   uint16(sym_const_literal),
	622:   uint16(374),
	623:   uint16(1),
	624:   uint16(sym_type_declaration),
	625:   uint16(389),
	626:   uint16(1),
	627:   uint16(sym_const_expression),
	628:   uint16(3),
	629:   uint16(2),
	630:   uint16(sym_block_comment),
	631:   uint16(sym_line_comment),
	632:   uint16(147),
	633:   uint16(2),
	634:   uint16(aux_sym_float_literal_token1),
	635:   uint16(aux_sym_float_literal_token2),
	636:   uint16(150),
	637:   uint16(2),
	638:   uint16(anon_sym_true),
	639:   uint16(anon_sym_false),
	640:   uint16(93),
	641:   uint16(2),
	642:   uint16(sym_float_literal),
	643:   uint16(sym_bool_literal),
	644:   uint16(376),
	645:   uint16(2),
	646:   uint16(sym__vec_prefix),
	647:   uint16(sym__mat_prefix),
	648:   uint16(162),
	649:   uint16(4),
	650:   uint16(anon_sym_texture_storage_1d),
	651:   uint16(anon_sym_texture_storage_2d),
	652:   uint16(anon_sym_texture_storage_2d_array),
	653:   uint16(anon_sym_texture_storage_3d),
	654:   uint16(159),
	655:   uint16(7),
	656:   uint16(anon_sym_texture_1d),
	657:   uint16(anon_sym_texture_2d),
	658:   uint16(anon_sym_texture_2d_array),
	659:   uint16(anon_sym_texture_3d),
	660:   uint16(anon_sym_texture_cube),
	661:   uint16(anon_sym_texture_cube_array),
	662:   uint16(anon_sym_texture_multisampled_2d),
	663:   uint16(165),
	664:   uint16(12),
	665:   uint16(anon_sym_vec2),
	666:   uint16(anon_sym_vec3),
	667:   uint16(anon_sym_vec4),
	668:   uint16(anon_sym_mat2x2),
	669:   uint16(anon_sym_mat2x3),
	670:   uint16(anon_sym_mat2x4),
	671:   uint16(anon_sym_mat3x2),
	672:   uint16(anon_sym_mat3x3),
	673:   uint16(anon_sym_mat3x4),
	674:   uint16(anon_sym_mat4x2),
	675:   uint16(anon_sym_mat4x3),
	676:   uint16(anon_sym_mat4x4),
	677:   uint16(141),
	678:   uint16(13),
	679:   uint16(sym_identifier),
	680:   uint16(anon_sym_bool),
	681:   uint16(anon_sym_u32),
	682:   uint16(anon_sym_i32),
	683:   uint16(anon_sym_f32),
	684:   uint16(anon_sym_f16),
	685:   uint16(anon_sym_sampler),
	686:   uint16(anon_sym_sampler_comparison),
	687:   uint16(anon_sym_texture_depth_2d),
	688:   uint16(anon_sym_texture_depth_2d_array),
	689:   uint16(anon_sym_texture_depth_cube),
	690:   uint16(anon_sym_texture_depth_cube_array),
	691:   uint16(anon_sym_texture_depth_multisampled_2d),
	692:   uint16(4),
	693:   uint16(168),
	694:   uint16(1),
	695:   uint16(anon_sym_RPAREN),
	696:   uint16(3),
	697:   uint16(2),
	698:   uint16(sym_block_comment),
	699:   uint16(sym_line_comment),
	700:   uint16(137),
	701:   uint16(7),
	702:   uint16(anon_sym_LPAREN),
	703:   uint16(aux_sym_float_literal_token1),
	704:   uint16(aux_sym_float_literal_token2),
	705:   uint16(anon_sym_AMP),
	706:   uint16(anon_sym_STAR),
	707:   uint16(anon_sym_BANG),
	708:   uint16(anon_sym_TILDE),
	709:   uint16(135),
	710:   uint16(43),
	711:   uint16(sym_identifier),
	712:   uint16(sym_int_literal),
	713:   uint16(anon_sym_true),
	714:   uint16(anon_sym_false),
	715:   uint16(anon_sym_bool),
	716:   uint16(anon_sym_u32),
	717:   uint16(anon_sym_i32),
	718:   uint16(anon_sym_f32),
	719:   uint16(anon_sym_f16),
	720:   uint16(anon_sym_array),
	721:   uint16(anon_sym_ptr),
	722:   uint16(anon_sym_sampler),
	723:   uint16(anon_sym_sampler_comparison),
	724:   uint16(anon_sym_texture_depth_2d),
	725:   uint16(anon_sym_texture_depth_2d_array),
	726:   uint16(anon_sym_texture_depth_cube),
	727:   uint16(anon_sym_texture_depth_cube_array),
	728:   uint16(anon_sym_texture_depth_multisampled_2d),
	729:   uint16(anon_sym_texture_1d),
	730:   uint16(anon_sym_texture_2d),
	731:   uint16(anon_sym_texture_2d_array),
	732:   uint16(anon_sym_texture_3d),
	733:   uint16(anon_sym_texture_cube),
	734:   uint16(anon_sym_texture_cube_array),
	735:   uint16(anon_sym_texture_multisampled_2d),
	736:   uint16(anon_sym_texture_storage_1d),
	737:   uint16(anon_sym_texture_storage_2d),
	738:   uint16(anon_sym_texture_storage_2d_array),
	739:   uint16(anon_sym_texture_storage_3d),
	740:   uint16(anon_sym_vec2),
	741:   uint16(anon_sym_vec3),
	742:   uint16(anon_sym_vec4),
	743:   uint16(anon_sym_mat2x2),
	744:   uint16(anon_sym_mat2x3),
	745:   uint16(anon_sym_mat2x4),
	746:   uint16(anon_sym_mat3x2),
	747:   uint16(anon_sym_mat3x3),
	748:   uint16(anon_sym_mat3x4),
	749:   uint16(anon_sym_mat4x2),
	750:   uint16(anon_sym_mat4x3),
	751:   uint16(anon_sym_mat4x4),
	752:   uint16(anon_sym_bitcast),
	753:   uint16(anon_sym_DASH),
	754:   uint16(15),
	755:   uint16(39),
	756:   uint16(1),
	757:   uint16(sym_int_literal),
	758:   uint16(47),
	759:   uint16(1),
	760:   uint16(anon_sym_array),
	761:   uint16(49),
	762:   uint16(1),
	763:   uint16(anon_sym_ptr),
	764:   uint16(274),
	765:   uint16(1),
	766:   uint16(sym_const_literal),
	767:   uint16(374),
	768:   uint16(1),
	769:   uint16(sym_type_declaration),
	770:   uint16(382),
	771:   uint16(1),
	772:   uint16(sym_const_expression),
	773:   uint16(3),
	774:   uint16(2),
	775:   uint16(sym_block_comment),
	776:   uint16(sym_line_comment),
	777:   uint16(41),
	778:   uint16(2),
	779:   uint16(aux_sym_float_literal_token1),
	780:   uint16(aux_sym_float_literal_token2),
	781:   uint16(43),
	782:   uint16(2),
	783:   uint16(anon_sym_true),
	784:   uint16(anon_sym_false),
	785:   uint16(93),
	786:   uint16(2),
	787:   uint16(sym_float_literal),
	788:   uint16(sym_bool_literal),
	789:   uint16(376),
	790:   uint16(2),
	791:   uint16(sym__vec_prefix),
	792:   uint16(sym__mat_prefix),
	793:   uint16(53),
	794:   uint16(4),
	795:   uint16(anon_sym_texture_storage_1d),
	796:   uint16(anon_sym_texture_storage_2d),
	797:   uint16(anon_sym_texture_storage_2d_array),
	798:   uint16(anon_sym_texture_storage_3d),
	799:   uint16(51),
	800:   uint16(7),
	801:   uint16(anon_sym_texture_1d),
	802:   uint16(anon_sym_texture_2d),
	803:   uint16(anon_sym_texture_2d_array),
	804:   uint16(anon_sym_texture_3d),
	805:   uint16(anon_sym_texture_cube),
	806:   uint16(anon_sym_texture_cube_array),
	807:   uint16(anon_sym_texture_multisampled_2d),
	808:   uint16(127),
	809:   uint16(12),
	810:   uint16(anon_sym_vec2),
	811:   uint16(anon_sym_vec3),
	812:   uint16(anon_sym_vec4),
	813:   uint16(anon_sym_mat2x2),
	814:   uint16(anon_sym_mat2x3),
	815:   uint16(anon_sym_mat2x4),
	816:   uint16(anon_sym_mat3x2),
	817:   uint16(anon_sym_mat3x3),
	818:   uint16(anon_sym_mat3x4),
	819:   uint16(anon_sym_mat4x2),
	820:   uint16(anon_sym_mat4x3),
	821:   uint16(anon_sym_mat4x4),
	822:   uint16(45),
	823:   uint16(13),
	824:   uint16(sym_identifier),
	825:   uint16(anon_sym_bool),
	826:   uint16(anon_sym_u32),
	827:   uint16(anon_sym_i32),
	828:   uint16(anon_sym_f32),
	829:   uint16(anon_sym_f16),
	830:   uint16(anon_sym_sampler),
	831:   uint16(anon_sym_sampler_comparison),
	832:   uint16(anon_sym_texture_depth_2d),
	833:   uint16(anon_sym_texture_depth_2d_array),
	834:   uint16(anon_sym_texture_depth_cube),
	835:   uint16(anon_sym_texture_depth_cube_array),
	836:   uint16(anon_sym_texture_depth_multisampled_2d),
	837:   uint16(15),
	838:   uint16(39),
	839:   uint16(1),
	840:   uint16(sym_int_literal),
	841:   uint16(47),
	842:   uint16(1),
	843:   uint16(anon_sym_array),
	844:   uint16(49),
	845:   uint16(1),
	846:   uint16(anon_sym_ptr),
	847:   uint16(274),
	848:   uint16(1),
	849:   uint16(sym_const_literal),
	850:   uint16(374),
	851:   uint16(1),
	852:   uint16(sym_type_declaration),
	853:   uint16(405),
	854:   uint16(1),
	855:   uint16(sym_const_expression),
	856:   uint16(3),
	857:   uint16(2),
	858:   uint16(sym_block_comment),
	859:   uint16(sym_line_comment),
	860:   uint16(41),
	861:   uint16(2),
	862:   uint16(aux_sym_float_literal_token1),
	863:   uint16(aux_sym_float_literal_token2),
	864:   uint16(43),
	865:   uint16(2),
	866:   uint16(anon_sym_true),
	867:   uint16(anon_sym_false),
	868:   uint16(93),
	869:   uint16(2),
	870:   uint16(sym_float_literal),
	871:   uint16(sym_bool_literal),
	872:   uint16(376),
	873:   uint16(2),
	874:   uint16(sym__vec_prefix),
	875:   uint16(sym__mat_prefix),
	876:   uint16(53),
	877:   uint16(4),
	878:   uint16(anon_sym_texture_storage_1d),
	879:   uint16(anon_sym_texture_storage_2d),
	880:   uint16(anon_sym_texture_storage_2d_array),
	881:   uint16(anon_sym_texture_storage_3d),
	882:   uint16(51),
	883:   uint16(7),
	884:   uint16(anon_sym_texture_1d),
	885:   uint16(anon_sym_texture_2d),
	886:   uint16(anon_sym_texture_2d_array),
	887:   uint16(anon_sym_texture_3d),
	888:   uint16(anon_sym_texture_cube),
	889:   uint16(anon_sym_texture_cube_array),
	890:   uint16(anon_sym_texture_multisampled_2d),
	891:   uint16(127),
	892:   uint16(12),
	893:   uint16(anon_sym_vec2),
	894:   uint16(anon_sym_vec3),
	895:   uint16(anon_sym_vec4),
	896:   uint16(anon_sym_mat2x2),
	897:   uint16(anon_sym_mat2x3),
	898:   uint16(anon_sym_mat2x4),
	899:   uint16(anon_sym_mat3x2),
	900:   uint16(anon_sym_mat3x3),
	901:   uint16(anon_sym_mat3x4),
	902:   uint16(anon_sym_mat4x2),
	903:   uint16(anon_sym_mat4x3),
	904:   uint16(anon_sym_mat4x4),
	905:   uint16(45),
	906:   uint16(13),
	907:   uint16(sym_identifier),
	908:   uint16(anon_sym_bool),
	909:   uint16(anon_sym_u32),
	910:   uint16(anon_sym_i32),
	911:   uint16(anon_sym_f32),
	912:   uint16(anon_sym_f16),
	913:   uint16(anon_sym_sampler),
	914:   uint16(anon_sym_sampler_comparison),
	915:   uint16(anon_sym_texture_depth_2d),
	916:   uint16(anon_sym_texture_depth_2d_array),
	917:   uint16(anon_sym_texture_depth_cube),
	918:   uint16(anon_sym_texture_depth_cube_array),
	919:   uint16(anon_sym_texture_depth_multisampled_2d),
	920:   uint16(3),
	921:   uint16(3),
	922:   uint16(2),
	923:   uint16(sym_block_comment),
	924:   uint16(sym_line_comment),
	925:   uint16(137),
	926:   uint16(7),
	927:   uint16(anon_sym_LPAREN),
	928:   uint16(aux_sym_float_literal_token1),
	929:   uint16(aux_sym_float_literal_token2),
	930:   uint16(anon_sym_AMP),
	931:   uint16(anon_sym_STAR),
	932:   uint16(anon_sym_BANG),
	933:   uint16(anon_sym_TILDE),
	934:   uint16(135),
	935:   uint16(43),
	936:   uint16(sym_identifier),
	937:   uint16(sym_int_literal),
	938:   uint16(anon_sym_true),
	939:   uint16(anon_sym_false),
	940:   uint16(anon_sym_bool),
	941:   uint16(anon_sym_u32),
	942:   uint16(anon_sym_i32),
	943:   uint16(anon_sym_f32),
	944:   uint16(anon_sym_f16),
	945:   uint16(anon_sym_array),
	946:   uint16(anon_sym_ptr),
	947:   uint16(anon_sym_sampler),
	948:   uint16(anon_sym_sampler_comparison),
	949:   uint16(anon_sym_texture_depth_2d),
	950:   uint16(anon_sym_texture_depth_2d_array),
	951:   uint16(anon_sym_texture_depth_cube),
	952:   uint16(anon_sym_texture_depth_cube_array),
	953:   uint16(anon_sym_texture_depth_multisampled_2d),
	954:   uint16(anon_sym_texture_1d),
	955:   uint16(anon_sym_texture_2d),
	956:   uint16(anon_sym_texture_2d_array),
	957:   uint16(anon_sym_texture_3d),
	958:   uint16(anon_sym_texture_cube),
	959:   uint16(anon_sym_texture_cube_array),
	960:   uint16(anon_sym_texture_multisampled_2d),
	961:   uint16(anon_sym_texture_storage_1d),
	962:   uint16(anon_sym_texture_storage_2d),
	963:   uint16(anon_sym_texture_storage_2d_array),
	964:   uint16(anon_sym_texture_storage_3d),
	965:   uint16(anon_sym_vec2),
	966:   uint16(anon_sym_vec3),
	967:   uint16(anon_sym_vec4),
	968:   uint16(anon_sym_mat2x2),
	969:   uint16(anon_sym_mat2x3),
	970:   uint16(anon_sym_mat2x4),
	971:   uint16(anon_sym_mat3x2),
	972:   uint16(anon_sym_mat3x3),
	973:   uint16(anon_sym_mat3x4),
	974:   uint16(anon_sym_mat4x2),
	975:   uint16(anon_sym_mat4x3),
	976:   uint16(anon_sym_mat4x4),
	977:   uint16(anon_sym_bitcast),
	978:   uint16(anon_sym_DASH),
	979:   uint16(3),
	980:   uint16(3),
	981:   uint16(2),
	982:   uint16(sym_block_comment),
	983:   uint16(sym_line_comment),
	984:   uint16(172),
	985:   uint16(7),
	986:   uint16(anon_sym_LPAREN),
	987:   uint16(aux_sym_float_literal_token1),
	988:   uint16(aux_sym_float_literal_token2),
	989:   uint16(anon_sym_AMP),
	990:   uint16(anon_sym_STAR),
	991:   uint16(anon_sym_BANG),
	992:   uint16(anon_sym_TILDE),
	993:   uint16(170),
	994:   uint16(43),
	995:   uint16(sym_identifier),
	996:   uint16(sym_int_literal),
	997:   uint16(anon_sym_true),
	998:   uint16(anon_sym_false),
	999:   uint16(anon_sym_bool),
	1000:  uint16(anon_sym_u32),
	1001:  uint16(anon_sym_i32),
	1002:  uint16(anon_sym_f32),
	1003:  uint16(anon_sym_f16),
	1004:  uint16(anon_sym_array),
	1005:  uint16(anon_sym_ptr),
	1006:  uint16(anon_sym_sampler),
	1007:  uint16(anon_sym_sampler_comparison),
	1008:  uint16(anon_sym_texture_depth_2d),
	1009:  uint16(anon_sym_texture_depth_2d_array),
	1010:  uint16(anon_sym_texture_depth_cube),
	1011:  uint16(anon_sym_texture_depth_cube_array),
	1012:  uint16(anon_sym_texture_depth_multisampled_2d),
	1013:  uint16(anon_sym_texture_1d),
	1014:  uint16(anon_sym_texture_2d),
	1015:  uint16(anon_sym_texture_2d_array),
	1016:  uint16(anon_sym_texture_3d),
	1017:  uint16(anon_sym_texture_cube),
	1018:  uint16(anon_sym_texture_cube_array),
	1019:  uint16(anon_sym_texture_multisampled_2d),
	1020:  uint16(anon_sym_texture_storage_1d),
	1021:  uint16(anon_sym_texture_storage_2d),
	1022:  uint16(anon_sym_texture_storage_2d_array),
	1023:  uint16(anon_sym_texture_storage_3d),
	1024:  uint16(anon_sym_vec2),
	1025:  uint16(anon_sym_vec3),
	1026:  uint16(anon_sym_vec4),
	1027:  uint16(anon_sym_mat2x2),
	1028:  uint16(anon_sym_mat2x3),
	1029:  uint16(anon_sym_mat2x4),
	1030:  uint16(anon_sym_mat3x2),
	1031:  uint16(anon_sym_mat3x3),
	1032:  uint16(anon_sym_mat3x4),
	1033:  uint16(anon_sym_mat4x2),
	1034:  uint16(anon_sym_mat4x3),
	1035:  uint16(anon_sym_mat4x4),
	1036:  uint16(anon_sym_bitcast),
	1037:  uint16(anon_sym_DASH),
	1038:  uint16(15),
	1039:  uint16(39),
	1040:  uint16(1),
	1041:  uint16(sym_int_literal),
	1042:  uint16(47),
	1043:  uint16(1),
	1044:  uint16(anon_sym_array),
	1045:  uint16(49),
	1046:  uint16(1),
	1047:  uint16(anon_sym_ptr),
	1048:  uint16(274),
	1049:  uint16(1),
	1050:  uint16(sym_const_literal),
	1051:  uint16(368),
	1052:  uint16(1),
	1053:  uint16(sym_const_expression),
	1054:  uint16(374),
	1055:  uint16(1),
	1056:  uint16(sym_type_declaration),
	1057:  uint16(3),
	1058:  uint16(2),
	1059:  uint16(sym_block_comment),
	1060:  uint16(sym_line_comment),
	1061:  uint16(41),
	1062:  uint16(2),
	1063:  uint16(aux_sym_float_literal_token1),
	1064:  uint16(aux_sym_float_literal_token2),
	1065:  uint16(43),
	1066:  uint16(2),
	1067:  uint16(anon_sym_true),
	1068:  uint16(anon_sym_false),
	1069:  uint16(93),
	1070:  uint16(2),
	1071:  uint16(sym_float_literal),
	1072:  uint16(sym_bool_literal),
	1073:  uint16(376),
	1074:  uint16(2),
	1075:  uint16(sym__vec_prefix),
	1076:  uint16(sym__mat_prefix),
	1077:  uint16(53),
	1078:  uint16(4),
	1079:  uint16(anon_sym_texture_storage_1d),
	1080:  uint16(anon_sym_texture_storage_2d),
	1081:  uint16(anon_sym_texture_storage_2d_array),
	1082:  uint16(anon_sym_texture_storage_3d),
	1083:  uint16(51),
	1084:  uint16(7),
	1085:  uint16(anon_sym_texture_1d),
	1086:  uint16(anon_sym_texture_2d),
	1087:  uint16(anon_sym_texture_2d_array),
	1088:  uint16(anon_sym_texture_3d),
	1089:  uint16(anon_sym_texture_cube),
	1090:  uint16(anon_sym_texture_cube_array),
	1091:  uint16(anon_sym_texture_multisampled_2d),
	1092:  uint16(127),
	1093:  uint16(12),
	1094:  uint16(anon_sym_vec2),
	1095:  uint16(anon_sym_vec3),
	1096:  uint16(anon_sym_vec4),
	1097:  uint16(anon_sym_mat2x2),
	1098:  uint16(anon_sym_mat2x3),
	1099:  uint16(anon_sym_mat2x4),
	1100:  uint16(anon_sym_mat3x2),
	1101:  uint16(anon_sym_mat3x3),
	1102:  uint16(anon_sym_mat3x4),
	1103:  uint16(anon_sym_mat4x2),
	1104:  uint16(anon_sym_mat4x3),
	1105:  uint16(anon_sym_mat4x4),
	1106:  uint16(45),
	1107:  uint16(13),
	1108:  uint16(sym_identifier),
	1109:  uint16(anon_sym_bool),
	1110:  uint16(anon_sym_u32),
	1111:  uint16(anon_sym_i32),
	1112:  uint16(anon_sym_f32),
	1113:  uint16(anon_sym_f16),
	1114:  uint16(anon_sym_sampler),
	1115:  uint16(anon_sym_sampler_comparison),
	1116:  uint16(anon_sym_texture_depth_2d),
	1117:  uint16(anon_sym_texture_depth_2d_array),
	1118:  uint16(anon_sym_texture_depth_cube),
	1119:  uint16(anon_sym_texture_depth_cube_array),
	1120:  uint16(anon_sym_texture_depth_multisampled_2d),
	1121:  uint16(26),
	1122:  uint16(174),
	1123:  uint16(1),
	1124:  uint16(sym_identifier),
	1125:  uint16(177),
	1126:  uint16(1),
	1127:  uint16(anon_sym_let),
	1128:  uint16(180),
	1129:  uint16(1),
	1130:  uint16(anon_sym_LPAREN),
	1131:  uint16(183),
	1132:  uint16(1),
	1133:  uint16(anon_sym_LBRACE),
	1134:  uint16(188),
	1135:  uint16(1),
	1136:  uint16(anon_sym__),
	1137:  uint16(191),
	1138:  uint16(1),
	1139:  uint16(anon_sym_if),
	1140:  uint16(194),
	1141:  uint16(1),
	1142:  uint16(anon_sym_switch),
	1143:  uint16(199),
	1144:  uint16(1),
	1145:  uint16(anon_sym_loop),
	1146:  uint16(202),
	1147:  uint16(1),
	1148:  uint16(anon_sym_for),
	1149:  uint16(205),
	1150:  uint16(1),
	1151:  uint16(anon_sym_while),
	1152:  uint16(208),
	1153:  uint16(1),
	1154:  uint16(anon_sym_break),
	1155:  uint16(211),
	1156:  uint16(1),
	1157:  uint16(anon_sym_continue),
	1158:  uint16(214),
	1159:  uint16(1),
	1160:  uint16(anon_sym_return),
	1161:  uint16(217),
	1162:  uint16(1),
	1163:  uint16(anon_sym_discard),
	1164:  uint16(220),
	1165:  uint16(1),
	1166:  uint16(anon_sym_var),
	1167:  uint16(226),
	1168:  uint16(1),
	1169:  uint16(aux_sym_preproc_import_token1),
	1170:  uint16(201),
	1171:  uint16(1),
	1172:  uint16(sym_lhs_expression),
	1173:  uint16(239),
	1174:  uint16(1),
	1175:  uint16(aux_sym_lhs_expression_repeat1),
	1176:  uint16(330),
	1177:  uint16(1),
	1178:  uint16(sym_variable_declaration),
	1179:  uint16(3),
	1180:  uint16(2),
	1181:  uint16(sym_block_comment),
	1182:  uint16(sym_line_comment),
	1183:  uint16(197),
	1184:  uint16(2),
	1185:  uint16(anon_sym_fallthrough),
	1186:  uint16(anon_sym_continuing),
	1187:  uint16(223),
	1188:  uint16(2),
	1189:  uint16(anon_sym_AMP),
	1190:  uint16(anon_sym_STAR),
	1191:  uint16(229),
	1192:  uint16(2),
	1193:  uint16(aux_sym_preproc_ifdef_token1),
	1194:  uint16(aux_sym_preproc_ifdef_token2),
	1195:  uint16(186),
	1196:  uint16(3),
	1197:  uint16(anon_sym_RBRACE),
	1198:  uint16(aux_sym_preproc_ifdef_token3),
	1199:  uint16(aux_sym_preproc_else_token1),
	1200:  uint16(398),
	1201:  uint16(3),
	1202:  uint16(sym_assignment_statement),
	1203:  uint16(sym_return_statement),
	1204:  uint16(sym_variable_statement),
	1205:  uint16(46),
	1206:  uint16(15),
	1207:  uint16(sym__statement),
	1208:  uint16(sym_compound_statement),
	1209:  uint16(sym_if_statement),
	1210:  uint16(sym_switch_statement),
	1211:  uint16(sym_loop_statement),
	1212:  uint16(sym_for_statement),
	1213:  uint16(sym_while_statement),
	1214:  uint16(sym_break_statement),
	1215:  uint16(sym_continue_statement),
	1216:  uint16(sym_discard_statement),
	1217:  uint16(sym_increment_statement),
	1218:  uint16(sym_decrement_statement),
	1219:  uint16(sym_preproc_import),
	1220:  uint16(sym_preproc_ifdef_in_statement),
	1221:  uint16(aux_sym_compound_statement_repeat1),
	1222:  uint16(27),
	1223:  uint16(27),
	1224:  uint16(1),
	1225:  uint16(aux_sym_preproc_import_token1),
	1226:  uint16(115),
	1227:  uint16(1),
	1228:  uint16(anon_sym_let),
	1229:  uint16(117),
	1230:  uint16(1),
	1231:  uint16(anon_sym_LPAREN),
	1232:  uint16(119),
	1233:  uint16(1),
	1234:  uint16(anon_sym__),
	1235:  uint16(121),
	1236:  uint16(1),
	1237:  uint16(anon_sym_var),
	1238:  uint16(232),
	1239:  uint16(1),
	1240:  uint16(sym_identifier),
	1241:  uint16(234),
	1242:  uint16(1),
	1243:  uint16(anon_sym_LBRACE),
	1244:  uint16(236),
	1245:  uint16(1),
	1246:  uint16(anon_sym_RBRACE),
	1247:  uint16(238),
	1248:  uint16(1),
	1249:  uint16(anon_sym_if),
	1250:  uint16(240),
	1251:  uint16(1),
	1252:  uint16(anon_sym_switch),
	1253:  uint16(242),
	1254:  uint16(1),
	1255:  uint16(anon_sym_fallthrough),
	1256:  uint16(244),
	1257:  uint16(1),
	1258:  uint16(anon_sym_loop),
	1259:  uint16(246),
	1260:  uint16(1),
	1261:  uint16(anon_sym_for),
	1262:  uint16(248),
	1263:  uint16(1),
	1264:  uint16(anon_sym_while),
	1265:  uint16(250),
	1266:  uint16(1),
	1267:  uint16(anon_sym_break),
	1268:  uint16(252),
	1269:  uint16(1),
	1270:  uint16(anon_sym_continue),
	1271:  uint16(254),
	1272:  uint16(1),
	1273:  uint16(anon_sym_return),
	1274:  uint16(256),
	1275:  uint16(1),
	1276:  uint16(anon_sym_discard),
	1277:  uint16(201),
	1278:  uint16(1),
	1279:  uint16(sym_lhs_expression),
	1280:  uint16(239),
	1281:  uint16(1),
	1282:  uint16(aux_sym_lhs_expression_repeat1),
	1283:  uint16(330),
	1284:  uint16(1),
	1285:  uint16(sym_variable_declaration),
	1286:  uint16(345),
	1287:  uint16(1),
	1288:  uint16(sym_fallthrough_statement),
	1289:  uint16(3),
	1290:  uint16(2),
	1291:  uint16(sym_block_comment),
	1292:  uint16(sym_line_comment),
	1293:  uint16(123),
	1294:  uint16(2),
	1295:  uint16(anon_sym_AMP),
	1296:  uint16(anon_sym_STAR),
	1297:  uint16(258),
	1298:  uint16(2),
	1299:  uint16(aux_sym_preproc_ifdef_token1),
	1300:  uint16(aux_sym_preproc_ifdef_token2),
	1301:  uint16(398),
	1302:  uint16(3),
	1303:  uint16(sym_assignment_statement),
	1304:  uint16(sym_return_statement),
	1305:  uint16(sym_variable_statement),
	1306:  uint16(46),
	1307:  uint16(15),
	1308:  uint16(sym__statement),
	1309:  uint16(sym_compound_statement),
	1310:  uint16(sym_if_statement),
	1311:  uint16(sym_switch_statement),
	1312:  uint16(sym_loop_statement),
	1313:  uint16(sym_for_statement),
	1314:  uint16(sym_while_statement),
	1315:  uint16(sym_break_statement),
	1316:  uint16(sym_continue_statement),
	1317:  uint16(sym_discard_statement),
	1318:  uint16(sym_increment_statement),
	1319:  uint16(sym_decrement_statement),
	1320:  uint16(sym_preproc_import),
	1321:  uint16(sym_preproc_ifdef_in_statement),
	1322:  uint16(aux_sym_compound_statement_repeat1),
	1323:  uint16(11),
	1324:  uint16(23),
	1325:  uint16(1),
	1326:  uint16(anon_sym_AT),
	1327:  uint16(47),
	1328:  uint16(1),
	1329:  uint16(anon_sym_array),
	1330:  uint16(49),
	1331:  uint16(1),
	1332:  uint16(anon_sym_ptr),
	1333:  uint16(406),
	1334:  uint16(1),
	1335:  uint16(sym_type_declaration),
	1336:  uint16(3),
	1337:  uint16(2),
	1338:  uint16(sym_block_comment),
	1339:  uint16(sym_line_comment),
	1340:  uint16(50),
	1341:  uint16(2),
	1342:  uint16(sym_attribute),
	1343:  uint16(aux_sym_global_variable_declaration_repeat1),
	1344:  uint16(376),
	1345:  uint16(2),
	1346:  uint16(sym__vec_prefix),
	1347:  uint16(sym__mat_prefix),
	1348:  uint16(53),
	1349:  uint16(4),
	1350:  uint16(anon_sym_texture_storage_1d),
	1351:  uint16(anon_sym_texture_storage_2d),
	1352:  uint16(anon_sym_texture_storage_2d_array),
	1353:  uint16(anon_sym_texture_storage_3d),
	1354:  uint16(51),
	1355:  uint16(7),
	1356:  uint16(anon_sym_texture_1d),
	1357:  uint16(anon_sym_texture_2d),
	1358:  uint16(anon_sym_texture_2d_array),
	1359:  uint16(anon_sym_texture_3d),
	1360:  uint16(anon_sym_texture_cube),
	1361:  uint16(anon_sym_texture_cube_array),
	1362:  uint16(anon_sym_texture_multisampled_2d),
	1363:  uint16(127),
	1364:  uint16(12),
	1365:  uint16(anon_sym_vec2),
	1366:  uint16(anon_sym_vec3),
	1367:  uint16(anon_sym_vec4),
	1368:  uint16(anon_sym_mat2x2),
	1369:  uint16(anon_sym_mat2x3),
	1370:  uint16(anon_sym_mat2x4),
	1371:  uint16(anon_sym_mat3x2),
	1372:  uint16(anon_sym_mat3x3),
	1373:  uint16(anon_sym_mat3x4),
	1374:  uint16(anon_sym_mat4x2),
	1375:  uint16(anon_sym_mat4x3),
	1376:  uint16(anon_sym_mat4x4),
	1377:  uint16(45),
	1378:  uint16(13),
	1379:  uint16(sym_identifier),
	1380:  uint16(anon_sym_bool),
	1381:  uint16(anon_sym_u32),
	1382:  uint16(anon_sym_i32),
	1383:  uint16(anon_sym_f32),
	1384:  uint16(anon_sym_f16),
	1385:  uint16(anon_sym_sampler),
	1386:  uint16(anon_sym_sampler_comparison),
	1387:  uint16(anon_sym_texture_depth_2d),
	1388:  uint16(anon_sym_texture_depth_2d_array),
	1389:  uint16(anon_sym_texture_depth_cube),
	1390:  uint16(anon_sym_texture_depth_cube_array),
	1391:  uint16(anon_sym_texture_depth_multisampled_2d),
	1392:  uint16(4),
	1393:  uint16(262),
	1394:  uint16(1),
	1395:  uint16(anon_sym_AT),
	1396:  uint16(3),
	1397:  uint16(2),
	1398:  uint16(sym_block_comment),
	1399:  uint16(sym_line_comment),
	1400:  uint16(49),
	1401:  uint16(2),
	1402:  uint16(sym_attribute),
	1403:  uint16(aux_sym_global_variable_declaration_repeat1),
	1404:  uint16(260),
	1405:  uint16(41),
	1406:  uint16(anon_sym_override),
	1407:  uint16(anon_sym_fn),
	1408:  uint16(sym_identifier),
	1409:  uint16(anon_sym_var),
	1410:  uint16(anon_sym_bool),
	1411:  uint16(anon_sym_u32),
	1412:  uint16(anon_sym_i32),
	1413:  uint16(anon_sym_f32),
	1414:  uint16(anon_sym_f16),
	1415:  uint16(anon_sym_array),
	1416:  uint16(anon_sym_ptr),
	1417:  uint16(anon_sym_sampler),
	1418:  uint16(anon_sym_sampler_comparison),
	1419:  uint16(anon_sym_texture_depth_2d),
	1420:  uint16(anon_sym_texture_depth_2d_array),
	1421:  uint16(anon_sym_texture_depth_cube),
	1422:  uint16(anon_sym_texture_depth_cube_array),
	1423:  uint16(anon_sym_texture_depth_multisampled_2d),
	1424:  uint16(anon_sym_texture_1d),
	1425:  uint16(anon_sym_texture_2d),
	1426:  uint16(anon_sym_texture_2d_array),
	1427:  uint16(anon_sym_texture_3d),
	1428:  uint16(anon_sym_texture_cube),
	1429:  uint16(anon_sym_texture_cube_array),
	1430:  uint16(anon_sym_texture_multisampled_2d),
	1431:  uint16(anon_sym_texture_storage_1d),
	1432:  uint16(anon_sym_texture_storage_2d),
	1433:  uint16(anon_sym_texture_storage_2d_array),
	1434:  uint16(anon_sym_texture_storage_3d),
	1435:  uint16(anon_sym_vec2),
	1436:  uint16(anon_sym_vec3),
	1437:  uint16(anon_sym_vec4),
	1438:  uint16(anon_sym_mat2x2),
	1439:  uint16(anon_sym_mat2x3),
	1440:  uint16(anon_sym_mat2x4),
	1441:  uint16(anon_sym_mat3x2),
	1442:  uint16(anon_sym_mat3x3),
	1443:  uint16(anon_sym_mat3x4),
	1444:  uint16(anon_sym_mat4x2),
	1445:  uint16(anon_sym_mat4x3),
	1446:  uint16(anon_sym_mat4x4),
	1447:  uint16(11),
	1448:  uint16(23),
	1449:  uint16(1),
	1450:  uint16(anon_sym_AT),
	1451:  uint16(47),
	1452:  uint16(1),
	1453:  uint16(anon_sym_array),
	1454:  uint16(49),
	1455:  uint16(1),
	1456:  uint16(anon_sym_ptr),
	1457:  uint16(373),
	1458:  uint16(1),
	1459:  uint16(sym_type_declaration),
	1460:  uint16(3),
	1461:  uint16(2),
	1462:  uint16(sym_block_comment),
	1463:  uint16(sym_line_comment),
	1464:  uint16(49),
	1465:  uint16(2),
	1466:  uint16(sym_attribute),
	1467:  uint16(aux_sym_global_variable_declaration_repeat1),
	1468:  uint16(376),
	1469:  uint16(2),
	1470:  uint16(sym__vec_prefix),
	1471:  uint16(sym__mat_prefix),
	1472:  uint16(53),
	1473:  uint16(4),
	1474:  uint16(anon_sym_texture_storage_1d),
	1475:  uint16(anon_sym_texture_storage_2d),
	1476:  uint16(anon_sym_texture_storage_2d_array),
	1477:  uint16(anon_sym_texture_storage_3d),
	1478:  uint16(51),
	1479:  uint16(7),
	1480:  uint16(anon_sym_texture_1d),
	1481:  uint16(anon_sym_texture_2d),
	1482:  uint16(anon_sym_texture_2d_array),
	1483:  uint16(anon_sym_texture_3d),
	1484:  uint16(anon_sym_texture_cube),
	1485:  uint16(anon_sym_texture_cube_array),
	1486:  uint16(anon_sym_texture_multisampled_2d),
	1487:  uint16(127),
	1488:  uint16(12),
	1489:  uint16(anon_sym_vec2),
	1490:  uint16(anon_sym_vec3),
	1491:  uint16(anon_sym_vec4),
	1492:  uint16(anon_sym_mat2x2),
	1493:  uint16(anon_sym_mat2x3),
	1494:  uint16(anon_sym_mat2x4),
	1495:  uint16(anon_sym_mat3x2),
	1496:  uint16(anon_sym_mat3x3),
	1497:  uint16(anon_sym_mat3x4),
	1498:  uint16(anon_sym_mat4x2),
	1499:  uint16(anon_sym_mat4x3),
	1500:  uint16(anon_sym_mat4x4),
	1501:  uint16(45),
	1502:  uint16(13),
	1503:  uint16(sym_identifier),
	1504:  uint16(anon_sym_bool),
	1505:  uint16(anon_sym_u32),
	1506:  uint16(anon_sym_i32),
	1507:  uint16(anon_sym_f32),
	1508:  uint16(anon_sym_f16),
	1509:  uint16(anon_sym_sampler),
	1510:  uint16(anon_sym_sampler_comparison),
	1511:  uint16(anon_sym_texture_depth_2d),
	1512:  uint16(anon_sym_texture_depth_2d_array),
	1513:  uint16(anon_sym_texture_depth_cube),
	1514:  uint16(anon_sym_texture_depth_cube_array),
	1515:  uint16(anon_sym_texture_depth_multisampled_2d),
	1516:  uint16(4),
	1517:  uint16(267),
	1518:  uint16(1),
	1519:  uint16(anon_sym_RPAREN),
	1520:  uint16(3),
	1521:  uint16(2),
	1522:  uint16(sym_block_comment),
	1523:  uint16(sym_line_comment),
	1524:  uint16(269),
	1525:  uint16(2),
	1526:  uint16(aux_sym_float_literal_token1),
	1527:  uint16(aux_sym_float_literal_token2),
	1528:  uint16(265),
	1529:  uint16(41),
	1530:  uint16(sym_identifier),
	1531:  uint16(sym_int_literal),
	1532:  uint16(anon_sym_true),
	1533:  uint16(anon_sym_false),
	1534:  uint16(anon_sym_bool),
	1535:  uint16(anon_sym_u32),
	1536:  uint16(anon_sym_i32),
	1537:  uint16(anon_sym_f32),
	1538:  uint16(anon_sym_f16),
	1539:  uint16(anon_sym_array),
	1540:  uint16(anon_sym_ptr),
	1541:  uint16(anon_sym_sampler),
	1542:  uint16(anon_sym_sampler_comparison),
	1543:  uint16(anon_sym_texture_depth_2d),
	1544:  uint16(anon_sym_texture_depth_2d_array),
	1545:  uint16(anon_sym_texture_depth_cube),
	1546:  uint16(anon_sym_texture_depth_cube_array),
	1547:  uint16(anon_sym_texture_depth_multisampled_2d),
	1548:  uint16(anon_sym_texture_1d),
	1549:  uint16(anon_sym_texture_2d),
	1550:  uint16(anon_sym_texture_2d_array),
	1551:  uint16(anon_sym_texture_3d),
	1552:  uint16(anon_sym_texture_cube),
	1553:  uint16(anon_sym_texture_cube_array),
	1554:  uint16(anon_sym_texture_multisampled_2d),
	1555:  uint16(anon_sym_texture_storage_1d),
	1556:  uint16(anon_sym_texture_storage_2d),
	1557:  uint16(anon_sym_texture_storage_2d_array),
	1558:  uint16(anon_sym_texture_storage_3d),
	1559:  uint16(anon_sym_vec2),
	1560:  uint16(anon_sym_vec3),
	1561:  uint16(anon_sym_vec4),
	1562:  uint16(anon_sym_mat2x2),
	1563:  uint16(anon_sym_mat2x3),
	1564:  uint16(anon_sym_mat2x4),
	1565:  uint16(anon_sym_mat3x2),
	1566:  uint16(anon_sym_mat3x3),
	1567:  uint16(anon_sym_mat3x4),
	1568:  uint16(anon_sym_mat4x2),
	1569:  uint16(anon_sym_mat4x3),
	1570:  uint16(anon_sym_mat4x4),
	1571:  uint16(27),
	1572:  uint16(27),
	1573:  uint16(1),
	1574:  uint16(aux_sym_preproc_import_token1),
	1575:  uint16(115),
	1576:  uint16(1),
	1577:  uint16(anon_sym_let),
	1578:  uint16(117),
	1579:  uint16(1),
	1580:  uint16(anon_sym_LPAREN),
	1581:  uint16(119),
	1582:  uint16(1),
	1583:  uint16(anon_sym__),
	1584:  uint16(121),
	1585:  uint16(1),
	1586:  uint16(anon_sym_var),
	1587:  uint16(232),
	1588:  uint16(1),
	1589:  uint16(sym_identifier),
	1590:  uint16(234),
	1591:  uint16(1),
	1592:  uint16(anon_sym_LBRACE),
	1593:  uint16(238),
	1594:  uint16(1),
	1595:  uint16(anon_sym_if),
	1596:  uint16(240),
	1597:  uint16(1),
	1598:  uint16(anon_sym_switch),
	1599:  uint16(244),
	1600:  uint16(1),
	1601:  uint16(anon_sym_loop),
	1602:  uint16(246),
	1603:  uint16(1),
	1604:  uint16(anon_sym_for),
	1605:  uint16(248),
	1606:  uint16(1),
	1607:  uint16(anon_sym_while),
	1608:  uint16(250),
	1609:  uint16(1),
	1610:  uint16(anon_sym_break),
	1611:  uint16(252),
	1612:  uint16(1),
	1613:  uint16(anon_sym_continue),
	1614:  uint16(254),
	1615:  uint16(1),
	1616:  uint16(anon_sym_return),
	1617:  uint16(256),
	1618:  uint16(1),
	1619:  uint16(anon_sym_discard),
	1620:  uint16(271),
	1621:  uint16(1),
	1622:  uint16(anon_sym_RBRACE),
	1623:  uint16(273),
	1624:  uint16(1),
	1625:  uint16(anon_sym_continuing),
	1626:  uint16(201),
	1627:  uint16(1),
	1628:  uint16(sym_lhs_expression),
	1629:  uint16(239),
	1630:  uint16(1),
	1631:  uint16(aux_sym_lhs_expression_repeat1),
	1632:  uint16(330),
	1633:  uint16(1),
	1634:  uint16(sym_variable_declaration),
	1635:  uint16(422),
	1636:  uint16(1),
	1637:  uint16(sym_continuing_statement),
	1638:  uint16(3),
	1639:  uint16(2),
	1640:  uint16(sym_block_comment),
	1641:  uint16(sym_line_comment),
	1642:  uint16(123),
	1643:  uint16(2),
	1644:  uint16(anon_sym_AMP),
	1645:  uint16(anon_sym_STAR),
	1646:  uint16(258),
	1647:  uint16(2),
	1648:  uint16(aux_sym_preproc_ifdef_token1),
	1649:  uint16(aux_sym_preproc_ifdef_token2),
	1650:  uint16(398),
	1651:  uint16(3),
	1652:  uint16(sym_assignment_statement),
	1653:  uint16(sym_return_statement),
	1654:  uint16(sym_variable_statement),
	1655:  uint16(54),
	1656:  uint16(15),
	1657:  uint16(sym__statement),
	1658:  uint16(sym_compound_statement),
	1659:  uint16(sym_if_statement),
	1660:  uint16(sym_switch_statement),
	1661:  uint16(sym_loop_statement),
	1662:  uint16(sym_for_statement),
	1663:  uint16(sym_while_statement),
	1664:  uint16(sym_break_statement),
	1665:  uint16(sym_continue_statement),
	1666:  uint16(sym_discard_statement),
	1667:  uint16(sym_increment_statement),
	1668:  uint16(sym_decrement_statement),
	1669:  uint16(sym_preproc_import),
	1670:  uint16(sym_preproc_ifdef_in_statement),
	1671:  uint16(aux_sym_compound_statement_repeat1),
	1672:  uint16(27),
	1673:  uint16(27),
	1674:  uint16(1),
	1675:  uint16(aux_sym_preproc_import_token1),
	1676:  uint16(115),
	1677:  uint16(1),
	1678:  uint16(anon_sym_let),
	1679:  uint16(117),
	1680:  uint16(1),
	1681:  uint16(anon_sym_LPAREN),
	1682:  uint16(119),
	1683:  uint16(1),
	1684:  uint16(anon_sym__),
	1685:  uint16(121),
	1686:  uint16(1),
	1687:  uint16(anon_sym_var),
	1688:  uint16(232),
	1689:  uint16(1),
	1690:  uint16(sym_identifier),
	1691:  uint16(234),
	1692:  uint16(1),
	1693:  uint16(anon_sym_LBRACE),
	1694:  uint16(238),
	1695:  uint16(1),
	1696:  uint16(anon_sym_if),
	1697:  uint16(240),
	1698:  uint16(1),
	1699:  uint16(anon_sym_switch),
	1700:  uint16(244),
	1701:  uint16(1),
	1702:  uint16(anon_sym_loop),
	1703:  uint16(246),
	1704:  uint16(1),
	1705:  uint16(anon_sym_for),
	1706:  uint16(248),
	1707:  uint16(1),
	1708:  uint16(anon_sym_while),
	1709:  uint16(250),
	1710:  uint16(1),
	1711:  uint16(anon_sym_break),
	1712:  uint16(252),
	1713:  uint16(1),
	1714:  uint16(anon_sym_continue),
	1715:  uint16(254),
	1716:  uint16(1),
	1717:  uint16(anon_sym_return),
	1718:  uint16(256),
	1719:  uint16(1),
	1720:  uint16(anon_sym_discard),
	1721:  uint16(275),
	1722:  uint16(1),
	1723:  uint16(aux_sym_preproc_ifdef_token3),
	1724:  uint16(277),
	1725:  uint16(1),
	1726:  uint16(aux_sym_preproc_else_token1),
	1727:  uint16(201),
	1728:  uint16(1),
	1729:  uint16(sym_lhs_expression),
	1730:  uint16(239),
	1731:  uint16(1),
	1732:  uint16(aux_sym_lhs_expression_repeat1),
	1733:  uint16(330),
	1734:  uint16(1),
	1735:  uint16(sym_variable_declaration),
	1736:  uint16(397),
	1737:  uint16(1),
	1738:  uint16(sym_preproc_else_in_statement),
	1739:  uint16(3),
	1740:  uint16(2),
	1741:  uint16(sym_block_comment),
	1742:  uint16(sym_line_comment),
	1743:  uint16(123),
	1744:  uint16(2),
	1745:  uint16(anon_sym_AMP),
	1746:  uint16(anon_sym_STAR),
	1747:  uint16(258),
	1748:  uint16(2),
	1749:  uint16(aux_sym_preproc_ifdef_token1),
	1750:  uint16(aux_sym_preproc_ifdef_token2),
	1751:  uint16(398),
	1752:  uint16(3),
	1753:  uint16(sym_assignment_statement),
	1754:  uint16(sym_return_statement),
	1755:  uint16(sym_variable_statement),
	1756:  uint16(56),
	1757:  uint16(15),
	1758:  uint16(sym__statement),
	1759:  uint16(sym_compound_statement),
	1760:  uint16(sym_if_statement),
	1761:  uint16(sym_switch_statement),
	1762:  uint16(sym_loop_statement),
	1763:  uint16(sym_for_statement),
	1764:  uint16(sym_while_statement),
	1765:  uint16(sym_break_statement),
	1766:  uint16(sym_continue_statement),
	1767:  uint16(sym_discard_statement),
	1768:  uint16(sym_increment_statement),
	1769:  uint16(sym_decrement_statement),
	1770:  uint16(sym_preproc_import),
	1771:  uint16(sym_preproc_ifdef_in_statement),
	1772:  uint16(aux_sym_compound_statement_repeat1),
	1773:  uint16(27),
	1774:  uint16(27),
	1775:  uint16(1),
	1776:  uint16(aux_sym_preproc_import_token1),
	1777:  uint16(115),
	1778:  uint16(1),
	1779:  uint16(anon_sym_let),
	1780:  uint16(117),
	1781:  uint16(1),
	1782:  uint16(anon_sym_LPAREN),
	1783:  uint16(119),
	1784:  uint16(1),
	1785:  uint16(anon_sym__),
	1786:  uint16(121),
	1787:  uint16(1),
	1788:  uint16(anon_sym_var),
	1789:  uint16(232),
	1790:  uint16(1),
	1791:  uint16(sym_identifier),
	1792:  uint16(234),
	1793:  uint16(1),
	1794:  uint16(anon_sym_LBRACE),
	1795:  uint16(238),
	1796:  uint16(1),
	1797:  uint16(anon_sym_if),
	1798:  uint16(240),
	1799:  uint16(1),
	1800:  uint16(anon_sym_switch),
	1801:  uint16(244),
	1802:  uint16(1),
	1803:  uint16(anon_sym_loop),
	1804:  uint16(246),
	1805:  uint16(1),
	1806:  uint16(anon_sym_for),
	1807:  uint16(248),
	1808:  uint16(1),
	1809:  uint16(anon_sym_while),
	1810:  uint16(250),
	1811:  uint16(1),
	1812:  uint16(anon_sym_break),
	1813:  uint16(252),
	1814:  uint16(1),
	1815:  uint16(anon_sym_continue),
	1816:  uint16(254),
	1817:  uint16(1),
	1818:  uint16(anon_sym_return),
	1819:  uint16(256),
	1820:  uint16(1),
	1821:  uint16(anon_sym_discard),
	1822:  uint16(273),
	1823:  uint16(1),
	1824:  uint16(anon_sym_continuing),
	1825:  uint16(279),
	1826:  uint16(1),
	1827:  uint16(anon_sym_RBRACE),
	1828:  uint16(201),
	1829:  uint16(1),
	1830:  uint16(sym_lhs_expression),
	1831:  uint16(239),
	1832:  uint16(1),
	1833:  uint16(aux_sym_lhs_expression_repeat1),
	1834:  uint16(330),
	1835:  uint16(1),
	1836:  uint16(sym_variable_declaration),
	1837:  uint16(369),
	1838:  uint16(1),
	1839:  uint16(sym_continuing_statement),
	1840:  uint16(3),
	1841:  uint16(2),
	1842:  uint16(sym_block_comment),
	1843:  uint16(sym_line_comment),
	1844:  uint16(123),
	1845:  uint16(2),
	1846:  uint16(anon_sym_AMP),
	1847:  uint16(anon_sym_STAR),
	1848:  uint16(258),
	1849:  uint16(2),
	1850:  uint16(aux_sym_preproc_ifdef_token1),
	1851:  uint16(aux_sym_preproc_ifdef_token2),
	1852:  uint16(398),
	1853:  uint16(3),
	1854:  uint16(sym_assignment_statement),
	1855:  uint16(sym_return_statement),
	1856:  uint16(sym_variable_statement),
	1857:  uint16(46),
	1858:  uint16(15),
	1859:  uint16(sym__statement),
	1860:  uint16(sym_compound_statement),
	1861:  uint16(sym_if_statement),
	1862:  uint16(sym_switch_statement),
	1863:  uint16(sym_loop_statement),
	1864:  uint16(sym_for_statement),
	1865:  uint16(sym_while_statement),
	1866:  uint16(sym_break_statement),
	1867:  uint16(sym_continue_statement),
	1868:  uint16(sym_discard_statement),
	1869:  uint16(sym_increment_statement),
	1870:  uint16(sym_decrement_statement),
	1871:  uint16(sym_preproc_import),
	1872:  uint16(sym_preproc_ifdef_in_statement),
	1873:  uint16(aux_sym_compound_statement_repeat1),
	1874:  uint16(4),
	1875:  uint16(281),
	1876:  uint16(1),
	1877:  uint16(anon_sym_RPAREN),
	1878:  uint16(3),
	1879:  uint16(2),
	1880:  uint16(sym_block_comment),
	1881:  uint16(sym_line_comment),
	1882:  uint16(269),
	1883:  uint16(2),
	1884:  uint16(aux_sym_float_literal_token1),
	1885:  uint16(aux_sym_float_literal_token2),
	1886:  uint16(265),
	1887:  uint16(41),
	1888:  uint16(sym_identifier),
	1889:  uint16(sym_int_literal),
	1890:  uint16(anon_sym_true),
	1891:  uint16(anon_sym_false),
	1892:  uint16(anon_sym_bool),
	1893:  uint16(anon_sym_u32),
	1894:  uint16(anon_sym_i32),
	1895:  uint16(anon_sym_f32),
	1896:  uint16(anon_sym_f16),
	1897:  uint16(anon_sym_array),
	1898:  uint16(anon_sym_ptr),
	1899:  uint16(anon_sym_sampler),
	1900:  uint16(anon_sym_sampler_comparison),
	1901:  uint16(anon_sym_texture_depth_2d),
	1902:  uint16(anon_sym_texture_depth_2d_array),
	1903:  uint16(anon_sym_texture_depth_cube),
	1904:  uint16(anon_sym_texture_depth_cube_array),
	1905:  uint16(anon_sym_texture_depth_multisampled_2d),
	1906:  uint16(anon_sym_texture_1d),
	1907:  uint16(anon_sym_texture_2d),
	1908:  uint16(anon_sym_texture_2d_array),
	1909:  uint16(anon_sym_texture_3d),
	1910:  uint16(anon_sym_texture_cube),
	1911:  uint16(anon_sym_texture_cube_array),
	1912:  uint16(anon_sym_texture_multisampled_2d),
	1913:  uint16(anon_sym_texture_storage_1d),
	1914:  uint16(anon_sym_texture_storage_2d),
	1915:  uint16(anon_sym_texture_storage_2d_array),
	1916:  uint16(anon_sym_texture_storage_3d),
	1917:  uint16(anon_sym_vec2),
	1918:  uint16(anon_sym_vec3),
	1919:  uint16(anon_sym_vec4),
	1920:  uint16(anon_sym_mat2x2),
	1921:  uint16(anon_sym_mat2x3),
	1922:  uint16(anon_sym_mat2x4),
	1923:  uint16(anon_sym_mat3x2),
	1924:  uint16(anon_sym_mat3x3),
	1925:  uint16(anon_sym_mat3x4),
	1926:  uint16(anon_sym_mat4x2),
	1927:  uint16(anon_sym_mat4x3),
	1928:  uint16(anon_sym_mat4x4),
	1929:  uint16(27),
	1930:  uint16(27),
	1931:  uint16(1),
	1932:  uint16(aux_sym_preproc_import_token1),
	1933:  uint16(115),
	1934:  uint16(1),
	1935:  uint16(anon_sym_let),
	1936:  uint16(117),
	1937:  uint16(1),
	1938:  uint16(anon_sym_LPAREN),
	1939:  uint16(119),
	1940:  uint16(1),
	1941:  uint16(anon_sym__),
	1942:  uint16(121),
	1943:  uint16(1),
	1944:  uint16(anon_sym_var),
	1945:  uint16(232),
	1946:  uint16(1),
	1947:  uint16(sym_identifier),
	1948:  uint16(234),
	1949:  uint16(1),
	1950:  uint16(anon_sym_LBRACE),
	1951:  uint16(238),
	1952:  uint16(1),
	1953:  uint16(anon_sym_if),
	1954:  uint16(240),
	1955:  uint16(1),
	1956:  uint16(anon_sym_switch),
	1957:  uint16(244),
	1958:  uint16(1),
	1959:  uint16(anon_sym_loop),
	1960:  uint16(246),
	1961:  uint16(1),
	1962:  uint16(anon_sym_for),
	1963:  uint16(248),
	1964:  uint16(1),
	1965:  uint16(anon_sym_while),
	1966:  uint16(250),
	1967:  uint16(1),
	1968:  uint16(anon_sym_break),
	1969:  uint16(252),
	1970:  uint16(1),
	1971:  uint16(anon_sym_continue),
	1972:  uint16(254),
	1973:  uint16(1),
	1974:  uint16(anon_sym_return),
	1975:  uint16(256),
	1976:  uint16(1),
	1977:  uint16(anon_sym_discard),
	1978:  uint16(277),
	1979:  uint16(1),
	1980:  uint16(aux_sym_preproc_else_token1),
	1981:  uint16(283),
	1982:  uint16(1),
	1983:  uint16(aux_sym_preproc_ifdef_token3),
	1984:  uint16(201),
	1985:  uint16(1),
	1986:  uint16(sym_lhs_expression),
	1987:  uint16(239),
	1988:  uint16(1),
	1989:  uint16(aux_sym_lhs_expression_repeat1),
	1990:  uint16(330),
	1991:  uint16(1),
	1992:  uint16(sym_variable_declaration),
	1993:  uint16(377),
	1994:  uint16(1),
	1995:  uint16(sym_preproc_else_in_statement),
	1996:  uint16(3),
	1997:  uint16(2),
	1998:  uint16(sym_block_comment),
	1999:  uint16(sym_line_comment),
	2000:  uint16(123),
	2001:  uint16(2),
	2002:  uint16(anon_sym_AMP),
	2003:  uint16(anon_sym_STAR),
	2004:  uint16(258),
	2005:  uint16(2),
	2006:  uint16(aux_sym_preproc_ifdef_token1),
	2007:  uint16(aux_sym_preproc_ifdef_token2),
	2008:  uint16(398),
	2009:  uint16(3),
	2010:  uint16(sym_assignment_statement),
	2011:  uint16(sym_return_statement),
	2012:  uint16(sym_variable_statement),
	2013:  uint16(46),
	2014:  uint16(15),
	2015:  uint16(sym__statement),
	2016:  uint16(sym_compound_statement),
	2017:  uint16(sym_if_statement),
	2018:  uint16(sym_switch_statement),
	2019:  uint16(sym_loop_statement),
	2020:  uint16(sym_for_statement),
	2021:  uint16(sym_while_statement),
	2022:  uint16(sym_break_statement),
	2023:  uint16(sym_continue_statement),
	2024:  uint16(sym_discard_statement),
	2025:  uint16(sym_increment_statement),
	2026:  uint16(sym_decrement_statement),
	2027:  uint16(sym_preproc_import),
	2028:  uint16(sym_preproc_ifdef_in_statement),
	2029:  uint16(aux_sym_compound_statement_repeat1),
	2030:  uint16(27),
	2031:  uint16(27),
	2032:  uint16(1),
	2033:  uint16(aux_sym_preproc_import_token1),
	2034:  uint16(115),
	2035:  uint16(1),
	2036:  uint16(anon_sym_let),
	2037:  uint16(117),
	2038:  uint16(1),
	2039:  uint16(anon_sym_LPAREN),
	2040:  uint16(119),
	2041:  uint16(1),
	2042:  uint16(anon_sym__),
	2043:  uint16(121),
	2044:  uint16(1),
	2045:  uint16(anon_sym_var),
	2046:  uint16(232),
	2047:  uint16(1),
	2048:  uint16(sym_identifier),
	2049:  uint16(234),
	2050:  uint16(1),
	2051:  uint16(anon_sym_LBRACE),
	2052:  uint16(238),
	2053:  uint16(1),
	2054:  uint16(anon_sym_if),
	2055:  uint16(240),
	2056:  uint16(1),
	2057:  uint16(anon_sym_switch),
	2058:  uint16(242),
	2059:  uint16(1),
	2060:  uint16(anon_sym_fallthrough),
	2061:  uint16(244),
	2062:  uint16(1),
	2063:  uint16(anon_sym_loop),
	2064:  uint16(246),
	2065:  uint16(1),
	2066:  uint16(anon_sym_for),
	2067:  uint16(248),
	2068:  uint16(1),
	2069:  uint16(anon_sym_while),
	2070:  uint16(250),
	2071:  uint16(1),
	2072:  uint16(anon_sym_break),
	2073:  uint16(252),
	2074:  uint16(1),
	2075:  uint16(anon_sym_continue),
	2076:  uint16(254),
	2077:  uint16(1),
	2078:  uint16(anon_sym_return),
	2079:  uint16(256),
	2080:  uint16(1),
	2081:  uint16(anon_sym_discard),
	2082:  uint16(285),
	2083:  uint16(1),
	2084:  uint16(anon_sym_RBRACE),
	2085:  uint16(201),
	2086:  uint16(1),
	2087:  uint16(sym_lhs_expression),
	2088:  uint16(239),
	2089:  uint16(1),
	2090:  uint16(aux_sym_lhs_expression_repeat1),
	2091:  uint16(330),
	2092:  uint16(1),
	2093:  uint16(sym_variable_declaration),
	2094:  uint16(410),
	2095:  uint16(1),
	2096:  uint16(sym_fallthrough_statement),
	2097:  uint16(3),
	2098:  uint16(2),
	2099:  uint16(sym_block_comment),
	2100:  uint16(sym_line_comment),
	2101:  uint16(123),
	2102:  uint16(2),
	2103:  uint16(anon_sym_AMP),
	2104:  uint16(anon_sym_STAR),
	2105:  uint16(258),
	2106:  uint16(2),
	2107:  uint16(aux_sym_preproc_ifdef_token1),
	2108:  uint16(aux_sym_preproc_ifdef_token2),
	2109:  uint16(398),
	2110:  uint16(3),
	2111:  uint16(sym_assignment_statement),
	2112:  uint16(sym_return_statement),
	2113:  uint16(sym_variable_statement),
	2114:  uint16(47),
	2115:  uint16(15),
	2116:  uint16(sym__statement),
	2117:  uint16(sym_compound_statement),
	2118:  uint16(sym_if_statement),
	2119:  uint16(sym_switch_statement),
	2120:  uint16(sym_loop_statement),
	2121:  uint16(sym_for_statement),
	2122:  uint16(sym_while_statement),
	2123:  uint16(sym_break_statement),
	2124:  uint16(sym_continue_statement),
	2125:  uint16(sym_discard_statement),
	2126:  uint16(sym_increment_statement),
	2127:  uint16(sym_decrement_statement),
	2128:  uint16(sym_preproc_import),
	2129:  uint16(sym_preproc_ifdef_in_statement),
	2130:  uint16(aux_sym_compound_statement_repeat1),
	2131:  uint16(26),
	2132:  uint16(27),
	2133:  uint16(1),
	2134:  uint16(aux_sym_preproc_import_token1),
	2135:  uint16(115),
	2136:  uint16(1),
	2137:  uint16(anon_sym_let),
	2138:  uint16(117),
	2139:  uint16(1),
	2140:  uint16(anon_sym_LPAREN),
	2141:  uint16(119),
	2142:  uint16(1),
	2143:  uint16(anon_sym__),
	2144:  uint16(121),
	2145:  uint16(1),
	2146:  uint16(anon_sym_var),
	2147:  uint16(232),
	2148:  uint16(1),
	2149:  uint16(sym_identifier),
	2150:  uint16(234),
	2151:  uint16(1),
	2152:  uint16(anon_sym_LBRACE),
	2153:  uint16(238),
	2154:  uint16(1),
	2155:  uint16(anon_sym_if),
	2156:  uint16(240),
	2157:  uint16(1),
	2158:  uint16(anon_sym_switch),
	2159:  uint16(244),
	2160:  uint16(1),
	2161:  uint16(anon_sym_loop),
	2162:  uint16(246),
	2163:  uint16(1),
	2164:  uint16(anon_sym_for),
	2165:  uint16(248),
	2166:  uint16(1),
	2167:  uint16(anon_sym_while),
	2168:  uint16(252),
	2169:  uint16(1),
	2170:  uint16(anon_sym_continue),
	2171:  uint16(254),
	2172:  uint16(1),
	2173:  uint16(anon_sym_return),
	2174:  uint16(256),
	2175:  uint16(1),
	2176:  uint16(anon_sym_discard),
	2177:  uint16(287),
	2178:  uint16(1),
	2179:  uint16(anon_sym_RBRACE),
	2180:  uint16(289),
	2181:  uint16(1),
	2182:  uint16(anon_sym_break),
	2183:  uint16(201),
	2184:  uint16(1),
	2185:  uint16(sym_lhs_expression),
	2186:  uint16(239),
	2187:  uint16(1),
	2188:  uint16(aux_sym_lhs_expression_repeat1),
	2189:  uint16(330),
	2190:  uint16(1),
	2191:  uint16(sym_variable_declaration),
	2192:  uint16(414),
	2193:  uint16(1),
	2194:  uint16(sym_break_if_statement),
	2195:  uint16(3),
	2196:  uint16(2),
	2197:  uint16(sym_block_comment),
	2198:  uint16(sym_line_comment),
	2199:  uint16(123),
	2200:  uint16(2),
	2201:  uint16(anon_sym_AMP),
	2202:  uint16(anon_sym_STAR),
	2203:  uint16(258),
	2204:  uint16(2),
	2205:  uint16(aux_sym_preproc_ifdef_token1),
	2206:  uint16(aux_sym_preproc_ifdef_token2),
	2207:  uint16(398),
	2208:  uint16(3),
	2209:  uint16(sym_assignment_statement),
	2210:  uint16(sym_return_statement),
	2211:  uint16(sym_variable_statement),
	2212:  uint16(46),
	2213:  uint16(15),
	2214:  uint16(sym__statement),
	2215:  uint16(sym_compound_statement),
	2216:  uint16(sym_if_statement),
	2217:  uint16(sym_switch_statement),
	2218:  uint16(sym_loop_statement),
	2219:  uint16(sym_for_statement),
	2220:  uint16(sym_while_statement),
	2221:  uint16(sym_break_statement),
	2222:  uint16(sym_continue_statement),
	2223:  uint16(sym_discard_statement),
	2224:  uint16(sym_increment_statement),
	2225:  uint16(sym_decrement_statement),
	2226:  uint16(sym_preproc_import),
	2227:  uint16(sym_preproc_ifdef_in_statement),
	2228:  uint16(aux_sym_compound_statement_repeat1),
	2229:  uint16(4),
	2230:  uint16(293),
	2231:  uint16(1),
	2232:  uint16(anon_sym_LPAREN),
	2233:  uint16(295),
	2234:  uint16(1),
	2235:  uint16(anon_sym_AT),
	2236:  uint16(3),
	2237:  uint16(2),
	2238:  uint16(sym_block_comment),
	2239:  uint16(sym_line_comment),
	2240:  uint16(291),
	2241:  uint16(41),
	2242:  uint16(anon_sym_override),
	2243:  uint16(anon_sym_fn),
	2244:  uint16(sym_identifier),
	2245:  uint16(anon_sym_var),
	2246:  uint16(anon_sym_bool),
	2247:  uint16(anon_sym_u32),
	2248:  uint16(anon_sym_i32),
	2249:  uint16(anon_sym_f32),
	2250:  uint16(anon_sym_f16),
	2251:  uint16(anon_sym_array),
	2252:  uint16(anon_sym_ptr),
	2253:  uint16(anon_sym_sampler),
	2254:  uint16(anon_sym_sampler_comparison),
	2255:  uint16(anon_sym_texture_depth_2d),
	2256:  uint16(anon_sym_texture_depth_2d_array),
	2257:  uint16(anon_sym_texture_depth_cube),
	2258:  uint16(anon_sym_texture_depth_cube_array),
	2259:  uint16(anon_sym_texture_depth_multisampled_2d),
	2260:  uint16(anon_sym_texture_1d),
	2261:  uint16(anon_sym_texture_2d),
	2262:  uint16(anon_sym_texture_2d_array),
	2263:  uint16(anon_sym_texture_3d),
	2264:  uint16(anon_sym_texture_cube),
	2265:  uint16(anon_sym_texture_cube_array),
	2266:  uint16(anon_sym_texture_multisampled_2d),
	2267:  uint16(anon_sym_texture_storage_1d),
	2268:  uint16(anon_sym_texture_storage_2d),
	2269:  uint16(anon_sym_texture_storage_2d_array),
	2270:  uint16(anon_sym_texture_storage_3d),
	2271:  uint16(anon_sym_vec2),
	2272:  uint16(anon_sym_vec3),
	2273:  uint16(anon_sym_vec4),
	2274:  uint16(anon_sym_mat2x2),
	2275:  uint16(anon_sym_mat2x3),
	2276:  uint16(anon_sym_mat2x4),
	2277:  uint16(anon_sym_mat3x2),
	2278:  uint16(anon_sym_mat3x3),
	2279:  uint16(anon_sym_mat3x4),
	2280:  uint16(anon_sym_mat4x2),
	2281:  uint16(anon_sym_mat4x3),
	2282:  uint16(anon_sym_mat4x4),
	2283:  uint16(3),
	2284:  uint16(3),
	2285:  uint16(2),
	2286:  uint16(sym_block_comment),
	2287:  uint16(sym_line_comment),
	2288:  uint16(269),
	2289:  uint16(2),
	2290:  uint16(aux_sym_float_literal_token1),
	2291:  uint16(aux_sym_float_literal_token2),
	2292:  uint16(265),
	2293:  uint16(41),
	2294:  uint16(sym_identifier),
	2295:  uint16(sym_int_literal),
	2296:  uint16(anon_sym_true),
	2297:  uint16(anon_sym_false),
	2298:  uint16(anon_sym_bool),
	2299:  uint16(anon_sym_u32),
	2300:  uint16(anon_sym_i32),
	2301:  uint16(anon_sym_f32),
	2302:  uint16(anon_sym_f16),
	2303:  uint16(anon_sym_array),
	2304:  uint16(anon_sym_ptr),
	2305:  uint16(anon_sym_sampler),
	2306:  uint16(anon_sym_sampler_comparison),
	2307:  uint16(anon_sym_texture_depth_2d),
	2308:  uint16(anon_sym_texture_depth_2d_array),
	2309:  uint16(anon_sym_texture_depth_cube),
	2310:  uint16(anon_sym_texture_depth_cube_array),
	2311:  uint16(anon_sym_texture_depth_multisampled_2d),
	2312:  uint16(anon_sym_texture_1d),
	2313:  uint16(anon_sym_texture_2d),
	2314:  uint16(anon_sym_texture_2d_array),
	2315:  uint16(anon_sym_texture_3d),
	2316:  uint16(anon_sym_texture_cube),
	2317:  uint16(anon_sym_texture_cube_array),
	2318:  uint16(anon_sym_texture_multisampled_2d),
	2319:  uint16(anon_sym_texture_storage_1d),
	2320:  uint16(anon_sym_texture_storage_2d),
	2321:  uint16(anon_sym_texture_storage_2d_array),
	2322:  uint16(anon_sym_texture_storage_3d),
	2323:  uint16(anon_sym_vec2),
	2324:  uint16(anon_sym_vec3),
	2325:  uint16(anon_sym_vec4),
	2326:  uint16(anon_sym_mat2x2),
	2327:  uint16(anon_sym_mat2x3),
	2328:  uint16(anon_sym_mat2x4),
	2329:  uint16(anon_sym_mat3x2),
	2330:  uint16(anon_sym_mat3x3),
	2331:  uint16(anon_sym_mat3x4),
	2332:  uint16(anon_sym_mat4x2),
	2333:  uint16(anon_sym_mat4x3),
	2334:  uint16(anon_sym_mat4x4),
	2335:  uint16(26),
	2336:  uint16(27),
	2337:  uint16(1),
	2338:  uint16(aux_sym_preproc_import_token1),
	2339:  uint16(115),
	2340:  uint16(1),
	2341:  uint16(anon_sym_let),
	2342:  uint16(117),
	2343:  uint16(1),
	2344:  uint16(anon_sym_LPAREN),
	2345:  uint16(119),
	2346:  uint16(1),
	2347:  uint16(anon_sym__),
	2348:  uint16(121),
	2349:  uint16(1),
	2350:  uint16(anon_sym_var),
	2351:  uint16(232),
	2352:  uint16(1),
	2353:  uint16(sym_identifier),
	2354:  uint16(234),
	2355:  uint16(1),
	2356:  uint16(anon_sym_LBRACE),
	2357:  uint16(238),
	2358:  uint16(1),
	2359:  uint16(anon_sym_if),
	2360:  uint16(240),
	2361:  uint16(1),
	2362:  uint16(anon_sym_switch),
	2363:  uint16(244),
	2364:  uint16(1),
	2365:  uint16(anon_sym_loop),
	2366:  uint16(246),
	2367:  uint16(1),
	2368:  uint16(anon_sym_for),
	2369:  uint16(248),
	2370:  uint16(1),
	2371:  uint16(anon_sym_while),
	2372:  uint16(252),
	2373:  uint16(1),
	2374:  uint16(anon_sym_continue),
	2375:  uint16(254),
	2376:  uint16(1),
	2377:  uint16(anon_sym_return),
	2378:  uint16(256),
	2379:  uint16(1),
	2380:  uint16(anon_sym_discard),
	2381:  uint16(289),
	2382:  uint16(1),
	2383:  uint16(anon_sym_break),
	2384:  uint16(297),
	2385:  uint16(1),
	2386:  uint16(anon_sym_RBRACE),
	2387:  uint16(201),
	2388:  uint16(1),
	2389:  uint16(sym_lhs_expression),
	2390:  uint16(239),
	2391:  uint16(1),
	2392:  uint16(aux_sym_lhs_expression_repeat1),
	2393:  uint16(330),
	2394:  uint16(1),
	2395:  uint16(sym_variable_declaration),
	2396:  uint16(393),
	2397:  uint16(1),
	2398:  uint16(sym_break_if_statement),
	2399:  uint16(3),
	2400:  uint16(2),
	2401:  uint16(sym_block_comment),
	2402:  uint16(sym_line_comment),
	2403:  uint16(123),
	2404:  uint16(2),
	2405:  uint16(anon_sym_AMP),
	2406:  uint16(anon_sym_STAR),
	2407:  uint16(258),
	2408:  uint16(2),
	2409:  uint16(aux_sym_preproc_ifdef_token1),
	2410:  uint16(aux_sym_preproc_ifdef_token2),
	2411:  uint16(398),
	2412:  uint16(3),
	2413:  uint16(sym_assignment_statement),
	2414:  uint16(sym_return_statement),
	2415:  uint16(sym_variable_statement),
	2416:  uint16(58),
	2417:  uint16(15),
	2418:  uint16(sym__statement),
	2419:  uint16(sym_compound_statement),
	2420:  uint16(sym_if_statement),
	2421:  uint16(sym_switch_statement),
	2422:  uint16(sym_loop_statement),
	2423:  uint16(sym_for_statement),
	2424:  uint16(sym_while_statement),
	2425:  uint16(sym_break_statement),
	2426:  uint16(sym_continue_statement),
	2427:  uint16(sym_discard_statement),
	2428:  uint16(sym_increment_statement),
	2429:  uint16(sym_decrement_statement),
	2430:  uint16(sym_preproc_import),
	2431:  uint16(sym_preproc_ifdef_in_statement),
	2432:  uint16(aux_sym_compound_statement_repeat1),
	2433:  uint16(25),
	2434:  uint16(27),
	2435:  uint16(1),
	2436:  uint16(aux_sym_preproc_import_token1),
	2437:  uint16(115),
	2438:  uint16(1),
	2439:  uint16(anon_sym_let),
	2440:  uint16(117),
	2441:  uint16(1),
	2442:  uint16(anon_sym_LPAREN),
	2443:  uint16(119),
	2444:  uint16(1),
	2445:  uint16(anon_sym__),
	2446:  uint16(121),
	2447:  uint16(1),
	2448:  uint16(anon_sym_var),
	2449:  uint16(232),
	2450:  uint16(1),
	2451:  uint16(sym_identifier),
	2452:  uint16(234),
	2453:  uint16(1),
	2454:  uint16(anon_sym_LBRACE),
	2455:  uint16(238),
	2456:  uint16(1),
	2457:  uint16(anon_sym_if),
	2458:  uint16(240),
	2459:  uint16(1),
	2460:  uint16(anon_sym_switch),
	2461:  uint16(244),
	2462:  uint16(1),
	2463:  uint16(anon_sym_loop),
	2464:  uint16(246),
	2465:  uint16(1),
	2466:  uint16(anon_sym_for),
	2467:  uint16(248),
	2468:  uint16(1),
	2469:  uint16(anon_sym_while),
	2470:  uint16(250),
	2471:  uint16(1),
	2472:  uint16(anon_sym_break),
	2473:  uint16(252),
	2474:  uint16(1),
	2475:  uint16(anon_sym_continue),
	2476:  uint16(254),
	2477:  uint16(1),
	2478:  uint16(anon_sym_return),
	2479:  uint16(256),
	2480:  uint16(1),
	2481:  uint16(anon_sym_discard),
	2482:  uint16(299),
	2483:  uint16(1),
	2484:  uint16(anon_sym_RBRACE),
	2485:  uint16(201),
	2486:  uint16(1),
	2487:  uint16(sym_lhs_expression),
	2488:  uint16(239),
	2489:  uint16(1),
	2490:  uint16(aux_sym_lhs_expression_repeat1),
	2491:  uint16(330),
	2492:  uint16(1),
	2493:  uint16(sym_variable_declaration),
	2494:  uint16(3),
	2495:  uint16(2),
	2496:  uint16(sym_block_comment),
	2497:  uint16(sym_line_comment),
	2498:  uint16(123),
	2499:  uint16(2),
	2500:  uint16(anon_sym_AMP),
	2501:  uint16(anon_sym_STAR),
	2502:  uint16(258),
	2503:  uint16(2),
	2504:  uint16(aux_sym_preproc_ifdef_token1),
	2505:  uint16(aux_sym_preproc_ifdef_token2),
	2506:  uint16(398),
	2507:  uint16(3),
	2508:  uint16(sym_assignment_statement),
	2509:  uint16(sym_return_statement),
	2510:  uint16(sym_variable_statement),
	2511:  uint16(64),
	2512:  uint16(15),
	2513:  uint16(sym__statement),
	2514:  uint16(sym_compound_statement),
	2515:  uint16(sym_if_statement),
	2516:  uint16(sym_switch_statement),
	2517:  uint16(sym_loop_statement),
	2518:  uint16(sym_for_statement),
	2519:  uint16(sym_while_statement),
	2520:  uint16(sym_break_statement),
	2521:  uint16(sym_continue_statement),
	2522:  uint16(sym_discard_statement),
	2523:  uint16(sym_increment_statement),
	2524:  uint16(sym_decrement_statement),
	2525:  uint16(sym_preproc_import),
	2526:  uint16(sym_preproc_ifdef_in_statement),
	2527:  uint16(aux_sym_compound_statement_repeat1),
	2528:  uint16(3),
	2529:  uint16(303),
	2530:  uint16(1),
	2531:  uint16(anon_sym_AT),
	2532:  uint16(3),
	2533:  uint16(2),
	2534:  uint16(sym_block_comment),
	2535:  uint16(sym_line_comment),
	2536:  uint16(301),
	2537:  uint16(41),
	2538:  uint16(anon_sym_override),
	2539:  uint16(anon_sym_fn),
	2540:  uint16(sym_identifier),
	2541:  uint16(anon_sym_var),
	2542:  uint16(anon_sym_bool),
	2543:  uint16(anon_sym_u32),
	2544:  uint16(anon_sym_i32),
	2545:  uint16(anon_sym_f32),
	2546:  uint16(anon_sym_f16),
	2547:  uint16(anon_sym_array),
	2548:  uint16(anon_sym_ptr),
	2549:  uint16(anon_sym_sampler),
	2550:  uint16(anon_sym_sampler_comparison),
	2551:  uint16(anon_sym_texture_depth_2d),
	2552:  uint16(anon_sym_texture_depth_2d_array),
	2553:  uint16(anon_sym_texture_depth_cube),
	2554:  uint16(anon_sym_texture_depth_cube_array),
	2555:  uint16(anon_sym_texture_depth_multisampled_2d),
	2556:  uint16(anon_sym_texture_1d),
	2557:  uint16(anon_sym_texture_2d),
	2558:  uint16(anon_sym_texture_2d_array),
	2559:  uint16(anon_sym_texture_3d),
	2560:  uint16(anon_sym_texture_cube),
	2561:  uint16(anon_sym_texture_cube_array),
	2562:  uint16(anon_sym_texture_multisampled_2d),
	2563:  uint16(anon_sym_texture_storage_1d),
	2564:  uint16(anon_sym_texture_storage_2d),
	2565:  uint16(anon_sym_texture_storage_2d_array),
	2566:  uint16(anon_sym_texture_storage_3d),
	2567:  uint16(anon_sym_vec2),
	2568:  uint16(anon_sym_vec3),
	2569:  uint16(anon_sym_vec4),
	2570:  uint16(anon_sym_mat2x2),
	2571:  uint16(anon_sym_mat2x3),
	2572:  uint16(anon_sym_mat2x4),
	2573:  uint16(anon_sym_mat3x2),
	2574:  uint16(anon_sym_mat3x3),
	2575:  uint16(anon_sym_mat3x4),
	2576:  uint16(anon_sym_mat4x2),
	2577:  uint16(anon_sym_mat4x3),
	2578:  uint16(anon_sym_mat4x4),
	2579:  uint16(25),
	2580:  uint16(27),
	2581:  uint16(1),
	2582:  uint16(aux_sym_preproc_import_token1),
	2583:  uint16(115),
	2584:  uint16(1),
	2585:  uint16(anon_sym_let),
	2586:  uint16(117),
	2587:  uint16(1),
	2588:  uint16(anon_sym_LPAREN),
	2589:  uint16(119),
	2590:  uint16(1),
	2591:  uint16(anon_sym__),
	2592:  uint16(121),
	2593:  uint16(1),
	2594:  uint16(anon_sym_var),
	2595:  uint16(232),
	2596:  uint16(1),
	2597:  uint16(sym_identifier),
	2598:  uint16(234),
	2599:  uint16(1),
	2600:  uint16(anon_sym_LBRACE),
	2601:  uint16(238),
	2602:  uint16(1),
	2603:  uint16(anon_sym_if),
	2604:  uint16(240),
	2605:  uint16(1),
	2606:  uint16(anon_sym_switch),
	2607:  uint16(244),
	2608:  uint16(1),
	2609:  uint16(anon_sym_loop),
	2610:  uint16(246),
	2611:  uint16(1),
	2612:  uint16(anon_sym_for),
	2613:  uint16(248),
	2614:  uint16(1),
	2615:  uint16(anon_sym_while),
	2616:  uint16(250),
	2617:  uint16(1),
	2618:  uint16(anon_sym_break),
	2619:  uint16(252),
	2620:  uint16(1),
	2621:  uint16(anon_sym_continue),
	2622:  uint16(254),
	2623:  uint16(1),
	2624:  uint16(anon_sym_return),
	2625:  uint16(256),
	2626:  uint16(1),
	2627:  uint16(anon_sym_discard),
	2628:  uint16(305),
	2629:  uint16(1),
	2630:  uint16(anon_sym_RBRACE),
	2631:  uint16(201),
	2632:  uint16(1),
	2633:  uint16(sym_lhs_expression),
	2634:  uint16(239),
	2635:  uint16(1),
	2636:  uint16(aux_sym_lhs_expression_repeat1),
	2637:  uint16(330),
	2638:  uint16(1),
	2639:  uint16(sym_variable_declaration),
	2640:  uint16(3),
	2641:  uint16(2),
	2642:  uint16(sym_block_comment),
	2643:  uint16(sym_line_comment),
	2644:  uint16(123),
	2645:  uint16(2),
	2646:  uint16(anon_sym_AMP),
	2647:  uint16(anon_sym_STAR),
	2648:  uint16(258),
	2649:  uint16(2),
	2650:  uint16(aux_sym_preproc_ifdef_token1),
	2651:  uint16(aux_sym_preproc_ifdef_token2),
	2652:  uint16(398),
	2653:  uint16(3),
	2654:  uint16(sym_assignment_statement),
	2655:  uint16(sym_return_statement),
	2656:  uint16(sym_variable_statement),
	2657:  uint16(46),
	2658:  uint16(15),
	2659:  uint16(sym__statement),
	2660:  uint16(sym_compound_statement),
	2661:  uint16(sym_if_statement),
	2662:  uint16(sym_switch_statement),
	2663:  uint16(sym_loop_statement),
	2664:  uint16(sym_for_statement),
	2665:  uint16(sym_while_statement),
	2666:  uint16(sym_break_statement),
	2667:  uint16(sym_continue_statement),
	2668:  uint16(sym_discard_statement),
	2669:  uint16(sym_increment_statement),
	2670:  uint16(sym_decrement_statement),
	2671:  uint16(sym_preproc_import),
	2672:  uint16(sym_preproc_ifdef_in_statement),
	2673:  uint16(aux_sym_compound_statement_repeat1),
	2674:  uint16(3),
	2675:  uint16(309),
	2676:  uint16(1),
	2677:  uint16(anon_sym_AT),
	2678:  uint16(3),
	2679:  uint16(2),
	2680:  uint16(sym_block_comment),
	2681:  uint16(sym_line_comment),
	2682:  uint16(307),
	2683:  uint16(41),
	2684:  uint16(anon_sym_override),
	2685:  uint16(anon_sym_fn),
	2686:  uint16(sym_identifier),
	2687:  uint16(anon_sym_var),
	2688:  uint16(anon_sym_bool),
	2689:  uint16(anon_sym_u32),
	2690:  uint16(anon_sym_i32),
	2691:  uint16(anon_sym_f32),
	2692:  uint16(anon_sym_f16),
	2693:  uint16(anon_sym_array),
	2694:  uint16(anon_sym_ptr),
	2695:  uint16(anon_sym_sampler),
	2696:  uint16(anon_sym_sampler_comparison),
	2697:  uint16(anon_sym_texture_depth_2d),
	2698:  uint16(anon_sym_texture_depth_2d_array),
	2699:  uint16(anon_sym_texture_depth_cube),
	2700:  uint16(anon_sym_texture_depth_cube_array),
	2701:  uint16(anon_sym_texture_depth_multisampled_2d),
	2702:  uint16(anon_sym_texture_1d),
	2703:  uint16(anon_sym_texture_2d),
	2704:  uint16(anon_sym_texture_2d_array),
	2705:  uint16(anon_sym_texture_3d),
	2706:  uint16(anon_sym_texture_cube),
	2707:  uint16(anon_sym_texture_cube_array),
	2708:  uint16(anon_sym_texture_multisampled_2d),
	2709:  uint16(anon_sym_texture_storage_1d),
	2710:  uint16(anon_sym_texture_storage_2d),
	2711:  uint16(anon_sym_texture_storage_2d_array),
	2712:  uint16(anon_sym_texture_storage_3d),
	2713:  uint16(anon_sym_vec2),
	2714:  uint16(anon_sym_vec3),
	2715:  uint16(anon_sym_vec4),
	2716:  uint16(anon_sym_mat2x2),
	2717:  uint16(anon_sym_mat2x3),
	2718:  uint16(anon_sym_mat2x4),
	2719:  uint16(anon_sym_mat3x2),
	2720:  uint16(anon_sym_mat3x3),
	2721:  uint16(anon_sym_mat3x4),
	2722:  uint16(anon_sym_mat4x2),
	2723:  uint16(anon_sym_mat4x3),
	2724:  uint16(anon_sym_mat4x4),
	2725:  uint16(3),
	2726:  uint16(313),
	2727:  uint16(1),
	2728:  uint16(anon_sym_AT),
	2729:  uint16(3),
	2730:  uint16(2),
	2731:  uint16(sym_block_comment),
	2732:  uint16(sym_line_comment),
	2733:  uint16(311),
	2734:  uint16(41),
	2735:  uint16(anon_sym_override),
	2736:  uint16(anon_sym_fn),
	2737:  uint16(sym_identifier),
	2738:  uint16(anon_sym_var),
	2739:  uint16(anon_sym_bool),
	2740:  uint16(anon_sym_u32),
	2741:  uint16(anon_sym_i32),
	2742:  uint16(anon_sym_f32),
	2743:  uint16(anon_sym_f16),
	2744:  uint16(anon_sym_array),
	2745:  uint16(anon_sym_ptr),
	2746:  uint16(anon_sym_sampler),
	2747:  uint16(anon_sym_sampler_comparison),
	2748:  uint16(anon_sym_texture_depth_2d),
	2749:  uint16(anon_sym_texture_depth_2d_array),
	2750:  uint16(anon_sym_texture_depth_cube),
	2751:  uint16(anon_sym_texture_depth_cube_array),
	2752:  uint16(anon_sym_texture_depth_multisampled_2d),
	2753:  uint16(anon_sym_texture_1d),
	2754:  uint16(anon_sym_texture_2d),
	2755:  uint16(anon_sym_texture_2d_array),
	2756:  uint16(anon_sym_texture_3d),
	2757:  uint16(anon_sym_texture_cube),
	2758:  uint16(anon_sym_texture_cube_array),
	2759:  uint16(anon_sym_texture_multisampled_2d),
	2760:  uint16(anon_sym_texture_storage_1d),
	2761:  uint16(anon_sym_texture_storage_2d),
	2762:  uint16(anon_sym_texture_storage_2d_array),
	2763:  uint16(anon_sym_texture_storage_3d),
	2764:  uint16(anon_sym_vec2),
	2765:  uint16(anon_sym_vec3),
	2766:  uint16(anon_sym_vec4),
	2767:  uint16(anon_sym_mat2x2),
	2768:  uint16(anon_sym_mat2x3),
	2769:  uint16(anon_sym_mat2x4),
	2770:  uint16(anon_sym_mat3x2),
	2771:  uint16(anon_sym_mat3x3),
	2772:  uint16(anon_sym_mat3x4),
	2773:  uint16(anon_sym_mat4x2),
	2774:  uint16(anon_sym_mat4x3),
	2775:  uint16(anon_sym_mat4x4),
	2776:  uint16(25),
	2777:  uint16(27),
	2778:  uint16(1),
	2779:  uint16(aux_sym_preproc_import_token1),
	2780:  uint16(115),
	2781:  uint16(1),
	2782:  uint16(anon_sym_let),
	2783:  uint16(117),
	2784:  uint16(1),
	2785:  uint16(anon_sym_LPAREN),
	2786:  uint16(119),
	2787:  uint16(1),
	2788:  uint16(anon_sym__),
	2789:  uint16(121),
	2790:  uint16(1),
	2791:  uint16(anon_sym_var),
	2792:  uint16(232),
	2793:  uint16(1),
	2794:  uint16(sym_identifier),
	2795:  uint16(234),
	2796:  uint16(1),
	2797:  uint16(anon_sym_LBRACE),
	2798:  uint16(238),
	2799:  uint16(1),
	2800:  uint16(anon_sym_if),
	2801:  uint16(240),
	2802:  uint16(1),
	2803:  uint16(anon_sym_switch),
	2804:  uint16(244),
	2805:  uint16(1),
	2806:  uint16(anon_sym_loop),
	2807:  uint16(246),
	2808:  uint16(1),
	2809:  uint16(anon_sym_for),
	2810:  uint16(248),
	2811:  uint16(1),
	2812:  uint16(anon_sym_while),
	2813:  uint16(250),
	2814:  uint16(1),
	2815:  uint16(anon_sym_break),
	2816:  uint16(252),
	2817:  uint16(1),
	2818:  uint16(anon_sym_continue),
	2819:  uint16(254),
	2820:  uint16(1),
	2821:  uint16(anon_sym_return),
	2822:  uint16(256),
	2823:  uint16(1),
	2824:  uint16(anon_sym_discard),
	2825:  uint16(315),
	2826:  uint16(1),
	2827:  uint16(aux_sym_preproc_ifdef_token3),
	2828:  uint16(201),
	2829:  uint16(1),
	2830:  uint16(sym_lhs_expression),
	2831:  uint16(239),
	2832:  uint16(1),
	2833:  uint16(aux_sym_lhs_expression_repeat1),
	2834:  uint16(330),
	2835:  uint16(1),
	2836:  uint16(sym_variable_declaration),
	2837:  uint16(3),
	2838:  uint16(2),
	2839:  uint16(sym_block_comment),
	2840:  uint16(sym_line_comment),
	2841:  uint16(123),
	2842:  uint16(2),
	2843:  uint16(anon_sym_AMP),
	2844:  uint16(anon_sym_STAR),
	2845:  uint16(258),
	2846:  uint16(2),
	2847:  uint16(aux_sym_preproc_ifdef_token1),
	2848:  uint16(aux_sym_preproc_ifdef_token2),
	2849:  uint16(398),
	2850:  uint16(3),
	2851:  uint16(sym_assignment_statement),
	2852:  uint16(sym_return_statement),
	2853:  uint16(sym_variable_statement),
	2854:  uint16(68),
	2855:  uint16(15),
	2856:  uint16(sym__statement),
	2857:  uint16(sym_compound_statement),
	2858:  uint16(sym_if_statement),
	2859:  uint16(sym_switch_statement),
	2860:  uint16(sym_loop_statement),
	2861:  uint16(sym_for_statement),
	2862:  uint16(sym_while_statement),
	2863:  uint16(sym_break_statement),
	2864:  uint16(sym_continue_statement),
	2865:  uint16(sym_discard_statement),
	2866:  uint16(sym_increment_statement),
	2867:  uint16(sym_decrement_statement),
	2868:  uint16(sym_preproc_import),
	2869:  uint16(sym_preproc_ifdef_in_statement),
	2870:  uint16(aux_sym_compound_statement_repeat1),
	2871:  uint16(25),
	2872:  uint16(27),
	2873:  uint16(1),
	2874:  uint16(aux_sym_preproc_import_token1),
	2875:  uint16(115),
	2876:  uint16(1),
	2877:  uint16(anon_sym_let),
	2878:  uint16(117),
	2879:  uint16(1),
	2880:  uint16(anon_sym_LPAREN),
	2881:  uint16(119),
	2882:  uint16(1),
	2883:  uint16(anon_sym__),
	2884:  uint16(121),
	2885:  uint16(1),
	2886:  uint16(anon_sym_var),
	2887:  uint16(232),
	2888:  uint16(1),
	2889:  uint16(sym_identifier),
	2890:  uint16(234),
	2891:  uint16(1),
	2892:  uint16(anon_sym_LBRACE),
	2893:  uint16(238),
	2894:  uint16(1),
	2895:  uint16(anon_sym_if),
	2896:  uint16(240),
	2897:  uint16(1),
	2898:  uint16(anon_sym_switch),
	2899:  uint16(244),
	2900:  uint16(1),
	2901:  uint16(anon_sym_loop),
	2902:  uint16(246),
	2903:  uint16(1),
	2904:  uint16(anon_sym_for),
	2905:  uint16(248),
	2906:  uint16(1),
	2907:  uint16(anon_sym_while),
	2908:  uint16(250),
	2909:  uint16(1),
	2910:  uint16(anon_sym_break),
	2911:  uint16(252),
	2912:  uint16(1),
	2913:  uint16(anon_sym_continue),
	2914:  uint16(254),
	2915:  uint16(1),
	2916:  uint16(anon_sym_return),
	2917:  uint16(256),
	2918:  uint16(1),
	2919:  uint16(anon_sym_discard),
	2920:  uint16(317),
	2921:  uint16(1),
	2922:  uint16(aux_sym_preproc_ifdef_token3),
	2923:  uint16(201),
	2924:  uint16(1),
	2925:  uint16(sym_lhs_expression),
	2926:  uint16(239),
	2927:  uint16(1),
	2928:  uint16(aux_sym_lhs_expression_repeat1),
	2929:  uint16(330),
	2930:  uint16(1),
	2931:  uint16(sym_variable_declaration),
	2932:  uint16(3),
	2933:  uint16(2),
	2934:  uint16(sym_block_comment),
	2935:  uint16(sym_line_comment),
	2936:  uint16(123),
	2937:  uint16(2),
	2938:  uint16(anon_sym_AMP),
	2939:  uint16(anon_sym_STAR),
	2940:  uint16(258),
	2941:  uint16(2),
	2942:  uint16(aux_sym_preproc_ifdef_token1),
	2943:  uint16(aux_sym_preproc_ifdef_token2),
	2944:  uint16(398),
	2945:  uint16(3),
	2946:  uint16(sym_assignment_statement),
	2947:  uint16(sym_return_statement),
	2948:  uint16(sym_variable_statement),
	2949:  uint16(46),
	2950:  uint16(15),
	2951:  uint16(sym__statement),
	2952:  uint16(sym_compound_statement),
	2953:  uint16(sym_if_statement),
	2954:  uint16(sym_switch_statement),
	2955:  uint16(sym_loop_statement),
	2956:  uint16(sym_for_statement),
	2957:  uint16(sym_while_statement),
	2958:  uint16(sym_break_statement),
	2959:  uint16(sym_continue_statement),
	2960:  uint16(sym_discard_statement),
	2961:  uint16(sym_increment_statement),
	2962:  uint16(sym_decrement_statement),
	2963:  uint16(sym_preproc_import),
	2964:  uint16(sym_preproc_ifdef_in_statement),
	2965:  uint16(aux_sym_compound_statement_repeat1),
	2966:  uint16(9),
	2967:  uint16(47),
	2968:  uint16(1),
	2969:  uint16(anon_sym_array),
	2970:  uint16(49),
	2971:  uint16(1),
	2972:  uint16(anon_sym_ptr),
	2973:  uint16(358),
	2974:  uint16(1),
	2975:  uint16(sym_type_declaration),
	2976:  uint16(3),
	2977:  uint16(2),
	2978:  uint16(sym_block_comment),
	2979:  uint16(sym_line_comment),
	2980:  uint16(376),
	2981:  uint16(2),
	2982:  uint16(sym__vec_prefix),
	2983:  uint16(sym__mat_prefix),
	2984:  uint16(53),
	2985:  uint16(4),
	2986:  uint16(anon_sym_texture_storage_1d),
	2987:  uint16(anon_sym_texture_storage_2d),
	2988:  uint16(anon_sym_texture_storage_2d_array),
	2989:  uint16(anon_sym_texture_storage_3d),
	2990:  uint16(51),
	2991:  uint16(7),
	2992:  uint16(anon_sym_texture_1d),
	2993:  uint16(anon_sym_texture_2d),
	2994:  uint16(anon_sym_texture_2d_array),
	2995:  uint16(anon_sym_texture_3d),
	2996:  uint16(anon_sym_texture_cube),
	2997:  uint16(anon_sym_texture_cube_array),
	2998:  uint16(anon_sym_texture_multisampled_2d),
	2999:  uint16(127),
	3000:  uint16(12),
	3001:  uint16(anon_sym_vec2),
	3002:  uint16(anon_sym_vec3),
	3003:  uint16(anon_sym_vec4),
	3004:  uint16(anon_sym_mat2x2),
	3005:  uint16(anon_sym_mat2x3),
	3006:  uint16(anon_sym_mat2x4),
	3007:  uint16(anon_sym_mat3x2),
	3008:  uint16(anon_sym_mat3x3),
	3009:  uint16(anon_sym_mat3x4),
	3010:  uint16(anon_sym_mat4x2),
	3011:  uint16(anon_sym_mat4x3),
	3012:  uint16(anon_sym_mat4x4),
	3013:  uint16(45),
	3014:  uint16(13),
	3015:  uint16(sym_identifier),
	3016:  uint16(anon_sym_bool),
	3017:  uint16(anon_sym_u32),
	3018:  uint16(anon_sym_i32),
	3019:  uint16(anon_sym_f32),
	3020:  uint16(anon_sym_f16),
	3021:  uint16(anon_sym_sampler),
	3022:  uint16(anon_sym_sampler_comparison),
	3023:  uint16(anon_sym_texture_depth_2d),
	3024:  uint16(anon_sym_texture_depth_2d_array),
	3025:  uint16(anon_sym_texture_depth_cube),
	3026:  uint16(anon_sym_texture_depth_cube_array),
	3027:  uint16(anon_sym_texture_depth_multisampled_2d),
	3028:  uint16(9),
	3029:  uint16(47),
	3030:  uint16(1),
	3031:  uint16(anon_sym_array),
	3032:  uint16(49),
	3033:  uint16(1),
	3034:  uint16(anon_sym_ptr),
	3035:  uint16(300),
	3036:  uint16(1),
	3037:  uint16(sym_type_declaration),
	3038:  uint16(3),
	3039:  uint16(2),
	3040:  uint16(sym_block_comment),
	3041:  uint16(sym_line_comment),
	3042:  uint16(376),
	3043:  uint16(2),
	3044:  uint16(sym__vec_prefix),
	3045:  uint16(sym__mat_prefix),
	3046:  uint16(53),
	3047:  uint16(4),
	3048:  uint16(anon_sym_texture_storage_1d),
	3049:  uint16(anon_sym_texture_storage_2d),
	3050:  uint16(anon_sym_texture_storage_2d_array),
	3051:  uint16(anon_sym_texture_storage_3d),
	3052:  uint16(51),
	3053:  uint16(7),
	3054:  uint16(anon_sym_texture_1d),
	3055:  uint16(anon_sym_texture_2d),
	3056:  uint16(anon_sym_texture_2d_array),
	3057:  uint16(anon_sym_texture_3d),
	3058:  uint16(anon_sym_texture_cube),
	3059:  uint16(anon_sym_texture_cube_array),
	3060:  uint16(anon_sym_texture_multisampled_2d),
	3061:  uint16(127),
	3062:  uint16(12),
	3063:  uint16(anon_sym_vec2),
	3064:  uint16(anon_sym_vec3),
	3065:  uint16(anon_sym_vec4),
	3066:  uint16(anon_sym_mat2x2),
	3067:  uint16(anon_sym_mat2x3),
	3068:  uint16(anon_sym_mat2x4),
	3069:  uint16(anon_sym_mat3x2),
	3070:  uint16(anon_sym_mat3x3),
	3071:  uint16(anon_sym_mat3x4),
	3072:  uint16(anon_sym_mat4x2),
	3073:  uint16(anon_sym_mat4x3),
	3074:  uint16(anon_sym_mat4x4),
	3075:  uint16(45),
	3076:  uint16(13),
	3077:  uint16(sym_identifier),
	3078:  uint16(anon_sym_bool),
	3079:  uint16(anon_sym_u32),
	3080:  uint16(anon_sym_i32),
	3081:  uint16(anon_sym_f32),
	3082:  uint16(anon_sym_f16),
	3083:  uint16(anon_sym_sampler),
	3084:  uint16(anon_sym_sampler_comparison),
	3085:  uint16(anon_sym_texture_depth_2d),
	3086:  uint16(anon_sym_texture_depth_2d_array),
	3087:  uint16(anon_sym_texture_depth_cube),
	3088:  uint16(anon_sym_texture_depth_cube_array),
	3089:  uint16(anon_sym_texture_depth_multisampled_2d),
	3090:  uint16(9),
	3091:  uint16(47),
	3092:  uint16(1),
	3093:  uint16(anon_sym_array),
	3094:  uint16(49),
	3095:  uint16(1),
	3096:  uint16(anon_sym_ptr),
	3097:  uint16(312),
	3098:  uint16(1),
	3099:  uint16(sym_type_declaration),
	3100:  uint16(3),
	3101:  uint16(2),
	3102:  uint16(sym_block_comment),
	3103:  uint16(sym_line_comment),
	3104:  uint16(376),
	3105:  uint16(2),
	3106:  uint16(sym__vec_prefix),
	3107:  uint16(sym__mat_prefix),
	3108:  uint16(53),
	3109:  uint16(4),
	3110:  uint16(anon_sym_texture_storage_1d),
	3111:  uint16(anon_sym_texture_storage_2d),
	3112:  uint16(anon_sym_texture_storage_2d_array),
	3113:  uint16(anon_sym_texture_storage_3d),
	3114:  uint16(51),
	3115:  uint16(7),
	3116:  uint16(anon_sym_texture_1d),
	3117:  uint16(anon_sym_texture_2d),
	3118:  uint16(anon_sym_texture_2d_array),
	3119:  uint16(anon_sym_texture_3d),
	3120:  uint16(anon_sym_texture_cube),
	3121:  uint16(anon_sym_texture_cube_array),
	3122:  uint16(anon_sym_texture_multisampled_2d),
	3123:  uint16(127),
	3124:  uint16(12),
	3125:  uint16(anon_sym_vec2),
	3126:  uint16(anon_sym_vec3),
	3127:  uint16(anon_sym_vec4),
	3128:  uint16(anon_sym_mat2x2),
	3129:  uint16(anon_sym_mat2x3),
	3130:  uint16(anon_sym_mat2x4),
	3131:  uint16(anon_sym_mat3x2),
	3132:  uint16(anon_sym_mat3x3),
	3133:  uint16(anon_sym_mat3x4),
	3134:  uint16(anon_sym_mat4x2),
	3135:  uint16(anon_sym_mat4x3),
	3136:  uint16(anon_sym_mat4x4),
	3137:  uint16(45),
	3138:  uint16(13),
	3139:  uint16(sym_identifier),
	3140:  uint16(anon_sym_bool),
	3141:  uint16(anon_sym_u32),
	3142:  uint16(anon_sym_i32),
	3143:  uint16(anon_sym_f32),
	3144:  uint16(anon_sym_f16),
	3145:  uint16(anon_sym_sampler),
	3146:  uint16(anon_sym_sampler_comparison),
	3147:  uint16(anon_sym_texture_depth_2d),
	3148:  uint16(anon_sym_texture_depth_2d_array),
	3149:  uint16(anon_sym_texture_depth_cube),
	3150:  uint16(anon_sym_texture_depth_cube_array),
	3151:  uint16(anon_sym_texture_depth_multisampled_2d),
	3152:  uint16(9),
	3153:  uint16(47),
	3154:  uint16(1),
	3155:  uint16(anon_sym_array),
	3156:  uint16(49),
	3157:  uint16(1),
	3158:  uint16(anon_sym_ptr),
	3159:  uint16(198),
	3160:  uint16(1),
	3161:  uint16(sym_type_declaration),
	3162:  uint16(3),
	3163:  uint16(2),
	3164:  uint16(sym_block_comment),
	3165:  uint16(sym_line_comment),
	3166:  uint16(376),
	3167:  uint16(2),
	3168:  uint16(sym__vec_prefix),
	3169:  uint16(sym__mat_prefix),
	3170:  uint16(53),
	3171:  uint16(4),
	3172:  uint16(anon_sym_texture_storage_1d),
	3173:  uint16(anon_sym_texture_storage_2d),
	3174:  uint16(anon_sym_texture_storage_2d_array),
	3175:  uint16(anon_sym_texture_storage_3d),
	3176:  uint16(51),
	3177:  uint16(7),
	3178:  uint16(anon_sym_texture_1d),
	3179:  uint16(anon_sym_texture_2d),
	3180:  uint16(anon_sym_texture_2d_array),
	3181:  uint16(anon_sym_texture_3d),
	3182:  uint16(anon_sym_texture_cube),
	3183:  uint16(anon_sym_texture_cube_array),
	3184:  uint16(anon_sym_texture_multisampled_2d),
	3185:  uint16(127),
	3186:  uint16(12),
	3187:  uint16(anon_sym_vec2),
	3188:  uint16(anon_sym_vec3),
	3189:  uint16(anon_sym_vec4),
	3190:  uint16(anon_sym_mat2x2),
	3191:  uint16(anon_sym_mat2x3),
	3192:  uint16(anon_sym_mat2x4),
	3193:  uint16(anon_sym_mat3x2),
	3194:  uint16(anon_sym_mat3x3),
	3195:  uint16(anon_sym_mat3x4),
	3196:  uint16(anon_sym_mat4x2),
	3197:  uint16(anon_sym_mat4x3),
	3198:  uint16(anon_sym_mat4x4),
	3199:  uint16(45),
	3200:  uint16(13),
	3201:  uint16(sym_identifier),
	3202:  uint16(anon_sym_bool),
	3203:  uint16(anon_sym_u32),
	3204:  uint16(anon_sym_i32),
	3205:  uint16(anon_sym_f32),
	3206:  uint16(anon_sym_f16),
	3207:  uint16(anon_sym_sampler),
	3208:  uint16(anon_sym_sampler_comparison),
	3209:  uint16(anon_sym_texture_depth_2d),
	3210:  uint16(anon_sym_texture_depth_2d_array),
	3211:  uint16(anon_sym_texture_depth_cube),
	3212:  uint16(anon_sym_texture_depth_cube_array),
	3213:  uint16(anon_sym_texture_depth_multisampled_2d),
	3214:  uint16(9),
	3215:  uint16(47),
	3216:  uint16(1),
	3217:  uint16(anon_sym_array),
	3218:  uint16(49),
	3219:  uint16(1),
	3220:  uint16(anon_sym_ptr),
	3221:  uint16(375),
	3222:  uint16(1),
	3223:  uint16(sym_type_declaration),
	3224:  uint16(3),
	3225:  uint16(2),
	3226:  uint16(sym_block_comment),
	3227:  uint16(sym_line_comment),
	3228:  uint16(376),
	3229:  uint16(2),
	3230:  uint16(sym__vec_prefix),
	3231:  uint16(sym__mat_prefix),
	3232:  uint16(53),
	3233:  uint16(4),
	3234:  uint16(anon_sym_texture_storage_1d),
	3235:  uint16(anon_sym_texture_storage_2d),
	3236:  uint16(anon_sym_texture_storage_2d_array),
	3237:  uint16(anon_sym_texture_storage_3d),
	3238:  uint16(51),
	3239:  uint16(7),
	3240:  uint16(anon_sym_texture_1d),
	3241:  uint16(anon_sym_texture_2d),
	3242:  uint16(anon_sym_texture_2d_array),
	3243:  uint16(anon_sym_texture_3d),
	3244:  uint16(anon_sym_texture_cube),
	3245:  uint16(anon_sym_texture_cube_array),
	3246:  uint16(anon_sym_texture_multisampled_2d),
	3247:  uint16(127),
	3248:  uint16(12),
	3249:  uint16(anon_sym_vec2),
	3250:  uint16(anon_sym_vec3),
	3251:  uint16(anon_sym_vec4),
	3252:  uint16(anon_sym_mat2x2),
	3253:  uint16(anon_sym_mat2x3),
	3254:  uint16(anon_sym_mat2x4),
	3255:  uint16(anon_sym_mat3x2),
	3256:  uint16(anon_sym_mat3x3),
	3257:  uint16(anon_sym_mat3x4),
	3258:  uint16(anon_sym_mat4x2),
	3259:  uint16(anon_sym_mat4x3),
	3260:  uint16(anon_sym_mat4x4),
	3261:  uint16(45),
	3262:  uint16(13),
	3263:  uint16(sym_identifier),
	3264:  uint16(anon_sym_bool),
	3265:  uint16(anon_sym_u32),
	3266:  uint16(anon_sym_i32),
	3267:  uint16(anon_sym_f32),
	3268:  uint16(anon_sym_f16),
	3269:  uint16(anon_sym_sampler),
	3270:  uint16(anon_sym_sampler_comparison),
	3271:  uint16(anon_sym_texture_depth_2d),
	3272:  uint16(anon_sym_texture_depth_2d_array),
	3273:  uint16(anon_sym_texture_depth_cube),
	3274:  uint16(anon_sym_texture_depth_cube_array),
	3275:  uint16(anon_sym_texture_depth_multisampled_2d),
	3276:  uint16(9),
	3277:  uint16(47),
	3278:  uint16(1),
	3279:  uint16(anon_sym_array),
	3280:  uint16(49),
	3281:  uint16(1),
	3282:  uint16(anon_sym_ptr),
	3283:  uint16(411),
	3284:  uint16(1),
	3285:  uint16(sym_type_declaration),
	3286:  uint16(3),
	3287:  uint16(2),
	3288:  uint16(sym_block_comment),
	3289:  uint16(sym_line_comment),
	3290:  uint16(376),
	3291:  uint16(2),
	3292:  uint16(sym__vec_prefix),
	3293:  uint16(sym__mat_prefix),
	3294:  uint16(53),
	3295:  uint16(4),
	3296:  uint16(anon_sym_texture_storage_1d),
	3297:  uint16(anon_sym_texture_storage_2d),
	3298:  uint16(anon_sym_texture_storage_2d_array),
	3299:  uint16(anon_sym_texture_storage_3d),
	3300:  uint16(51),
	3301:  uint16(7),
	3302:  uint16(anon_sym_texture_1d),
	3303:  uint16(anon_sym_texture_2d),
	3304:  uint16(anon_sym_texture_2d_array),
	3305:  uint16(anon_sym_texture_3d),
	3306:  uint16(anon_sym_texture_cube),
	3307:  uint16(anon_sym_texture_cube_array),
	3308:  uint16(anon_sym_texture_multisampled_2d),
	3309:  uint16(127),
	3310:  uint16(12),
	3311:  uint16(anon_sym_vec2),
	3312:  uint16(anon_sym_vec3),
	3313:  uint16(anon_sym_vec4),
	3314:  uint16(anon_sym_mat2x2),
	3315:  uint16(anon_sym_mat2x3),
	3316:  uint16(anon_sym_mat2x4),
	3317:  uint16(anon_sym_mat3x2),
	3318:  uint16(anon_sym_mat3x3),
	3319:  uint16(anon_sym_mat3x4),
	3320:  uint16(anon_sym_mat4x2),
	3321:  uint16(anon_sym_mat4x3),
	3322:  uint16(anon_sym_mat4x4),
	3323:  uint16(45),
	3324:  uint16(13),
	3325:  uint16(sym_identifier),
	3326:  uint16(anon_sym_bool),
	3327:  uint16(anon_sym_u32),
	3328:  uint16(anon_sym_i32),
	3329:  uint16(anon_sym_f32),
	3330:  uint16(anon_sym_f16),
	3331:  uint16(anon_sym_sampler),
	3332:  uint16(anon_sym_sampler_comparison),
	3333:  uint16(anon_sym_texture_depth_2d),
	3334:  uint16(anon_sym_texture_depth_2d_array),
	3335:  uint16(anon_sym_texture_depth_cube),
	3336:  uint16(anon_sym_texture_depth_cube_array),
	3337:  uint16(anon_sym_texture_depth_multisampled_2d),
	3338:  uint16(3),
	3339:  uint16(3),
	3340:  uint16(2),
	3341:  uint16(sym_block_comment),
	3342:  uint16(sym_line_comment),
	3343:  uint16(319),
	3344:  uint16(14),
	3346:  uint16(anon_sym_SEMI),
	3347:  uint16(anon_sym_LPAREN),
	3348:  uint16(anon_sym_LBRACE),
	3349:  uint16(anon_sym_RBRACE),
	3350:  uint16(anon_sym_AT),
	3351:  uint16(anon_sym_AMP),
	3352:  uint16(anon_sym_STAR),
	3353:  uint16(aux_sym_preproc_import_token1),
	3354:  uint16(aux_sym_define_import_path_token1),
	3355:  uint16(aux_sym_preproc_ifdef_token1),
	3356:  uint16(aux_sym_preproc_ifdef_token2),
	3357:  uint16(aux_sym_preproc_ifdef_token3),
	3358:  uint16(aux_sym_preproc_else_token1),
	3359:  uint16(321),
	3360:  uint16(21),
	3361:  uint16(anon_sym_let),
	3362:  uint16(anon_sym_override),
	3363:  uint16(anon_sym_type),
	3364:  uint16(anon_sym_virtual),
	3365:  uint16(anon_sym_fn),
	3366:  uint16(anon_sym_struct),
	3367:  uint16(sym_identifier),
	3368:  uint16(anon_sym__),
	3369:  uint16(anon_sym_if),
	3370:  uint16(anon_sym_else),
	3371:  uint16(anon_sym_switch),
	3372:  uint16(anon_sym_fallthrough),
	3373:  uint16(anon_sym_loop),
	3374:  uint16(anon_sym_for),
	3375:  uint16(anon_sym_while),
	3376:  uint16(anon_sym_break),
	3377:  uint16(anon_sym_continue),
	3378:  uint16(anon_sym_continuing),
	3379:  uint16(anon_sym_return),
	3380:  uint16(anon_sym_discard),
	3381:  uint16(anon_sym_var),
	3382:  uint16(3),
	3383:  uint16(3),
	3384:  uint16(2),
	3385:  uint16(sym_block_comment),
	3386:  uint16(sym_line_comment),
	3387:  uint16(323),
	3388:  uint16(15),
	3390:  uint16(anon_sym_SEMI),
	3391:  uint16(anon_sym_LPAREN),
	3392:  uint16(anon_sym_COMMA),
	3393:  uint16(anon_sym_LBRACE),
	3394:  uint16(anon_sym_RBRACE),
	3395:  uint16(anon_sym_AT),
	3396:  uint16(anon_sym_AMP),
	3397:  uint16(anon_sym_STAR),
	3398:  uint16(aux_sym_preproc_import_token1),
	3399:  uint16(aux_sym_define_import_path_token1),
	3400:  uint16(aux_sym_preproc_ifdef_token1),
	3401:  uint16(aux_sym_preproc_ifdef_token2),
	3402:  uint16(aux_sym_preproc_ifdef_token3),
	3403:  uint16(aux_sym_preproc_else_token1),
	3404:  uint16(325),
	3405:  uint16(20),
	3406:  uint16(anon_sym_let),
	3407:  uint16(anon_sym_override),
	3408:  uint16(anon_sym_type),
	3409:  uint16(anon_sym_virtual),
	3410:  uint16(anon_sym_fn),
	3411:  uint16(anon_sym_struct),
	3412:  uint16(sym_identifier),
	3413:  uint16(anon_sym__),
	3414:  uint16(anon_sym_if),
	3415:  uint16(anon_sym_switch),
	3416:  uint16(anon_sym_fallthrough),
	3417:  uint16(anon_sym_loop),
	3418:  uint16(anon_sym_for),
	3419:  uint16(anon_sym_while),
	3420:  uint16(anon_sym_break),
	3421:  uint16(anon_sym_continue),
	3422:  uint16(anon_sym_continuing),
	3423:  uint16(anon_sym_return),
	3424:  uint16(anon_sym_discard),
	3425:  uint16(anon_sym_var),
	3426:  uint16(3),
	3427:  uint16(3),
	3428:  uint16(2),
	3429:  uint16(sym_block_comment),
	3430:  uint16(sym_line_comment),
	3431:  uint16(327),
	3432:  uint16(15),
	3434:  uint16(anon_sym_SEMI),
	3435:  uint16(anon_sym_LPAREN),
	3436:  uint16(anon_sym_COMMA),
	3437:  uint16(anon_sym_LBRACE),
	3438:  uint16(anon_sym_RBRACE),
	3439:  uint16(anon_sym_AT),
	3440:  uint16(anon_sym_AMP),
	3441:  uint16(anon_sym_STAR),
	3442:  uint16(aux_sym_preproc_import_token1),
	3443:  uint16(aux_sym_define_import_path_token1),
	3444:  uint16(aux_sym_preproc_ifdef_token1),
	3445:  uint16(aux_sym_preproc_ifdef_token2),
	3446:  uint16(aux_sym_preproc_ifdef_token3),
	3447:  uint16(aux_sym_preproc_else_token1),
	3448:  uint16(329),
	3449:  uint16(20),
	3450:  uint16(anon_sym_let),
	3451:  uint16(anon_sym_override),
	3452:  uint16(anon_sym_type),
	3453:  uint16(anon_sym_virtual),
	3454:  uint16(anon_sym_fn),
	3455:  uint16(anon_sym_struct),
	3456:  uint16(sym_identifier),
	3457:  uint16(anon_sym__),
	3458:  uint16(anon_sym_if),
	3459:  uint16(anon_sym_switch),
	3460:  uint16(anon_sym_fallthrough),
	3461:  uint16(anon_sym_loop),
	3462:  uint16(anon_sym_for),
	3463:  uint16(anon_sym_while),
	3464:  uint16(anon_sym_break),
	3465:  uint16(anon_sym_continue),
	3466:  uint16(anon_sym_continuing),
	3467:  uint16(anon_sym_return),
	3468:  uint16(anon_sym_discard),
	3469:  uint16(anon_sym_var),
	3470:  uint16(3),
	3471:  uint16(3),
	3472:  uint16(2),
	3473:  uint16(sym_block_comment),
	3474:  uint16(sym_line_comment),
	3475:  uint16(331),
	3476:  uint16(15),
	3478:  uint16(anon_sym_SEMI),
	3479:  uint16(anon_sym_LPAREN),
	3480:  uint16(anon_sym_COMMA),
	3481:  uint16(anon_sym_LBRACE),
	3482:  uint16(anon_sym_RBRACE),
	3483:  uint16(anon_sym_AT),
	3484:  uint16(anon_sym_AMP),
	3485:  uint16(anon_sym_STAR),
	3486:  uint16(aux_sym_preproc_import_token1),
	3487:  uint16(aux_sym_define_import_path_token1),
	3488:  uint16(aux_sym_preproc_ifdef_token1),
	3489:  uint16(aux_sym_preproc_ifdef_token2),
	3490:  uint16(aux_sym_preproc_ifdef_token3),
	3491:  uint16(aux_sym_preproc_else_token1),
	3492:  uint16(333),
	3493:  uint16(20),
	3494:  uint16(anon_sym_let),
	3495:  uint16(anon_sym_override),
	3496:  uint16(anon_sym_type),
	3497:  uint16(anon_sym_virtual),
	3498:  uint16(anon_sym_fn),
	3499:  uint16(anon_sym_struct),
	3500:  uint16(sym_identifier),
	3501:  uint16(anon_sym__),
	3502:  uint16(anon_sym_if),
	3503:  uint16(anon_sym_switch),
	3504:  uint16(anon_sym_fallthrough),
	3505:  uint16(anon_sym_loop),
	3506:  uint16(anon_sym_for),
	3507:  uint16(anon_sym_while),
	3508:  uint16(anon_sym_break),
	3509:  uint16(anon_sym_continue),
	3510:  uint16(anon_sym_continuing),
	3511:  uint16(anon_sym_return),
	3512:  uint16(anon_sym_discard),
	3513:  uint16(anon_sym_var),
	3514:  uint16(3),
	3515:  uint16(3),
	3516:  uint16(2),
	3517:  uint16(sym_block_comment),
	3518:  uint16(sym_line_comment),
	3519:  uint16(335),
	3520:  uint16(15),
	3522:  uint16(anon_sym_SEMI),
	3523:  uint16(anon_sym_LPAREN),
	3524:  uint16(anon_sym_COMMA),
	3525:  uint16(anon_sym_LBRACE),
	3526:  uint16(anon_sym_RBRACE),
	3527:  uint16(anon_sym_AT),
	3528:  uint16(anon_sym_AMP),
	3529:  uint16(anon_sym_STAR),
	3530:  uint16(aux_sym_preproc_import_token1),
	3531:  uint16(aux_sym_define_import_path_token1),
	3532:  uint16(aux_sym_preproc_ifdef_token1),
	3533:  uint16(aux_sym_preproc_ifdef_token2),
	3534:  uint16(aux_sym_preproc_ifdef_token3),
	3535:  uint16(aux_sym_preproc_else_token1),
	3536:  uint16(337),
	3537:  uint16(20),
	3538:  uint16(anon_sym_let),
	3539:  uint16(anon_sym_override),
	3540:  uint16(anon_sym_type),
	3541:  uint16(anon_sym_virtual),
	3542:  uint16(anon_sym_fn),
	3543:  uint16(anon_sym_struct),
	3544:  uint16(sym_identifier),
	3545:  uint16(anon_sym__),
	3546:  uint16(anon_sym_if),
	3547:  uint16(anon_sym_switch),
	3548:  uint16(anon_sym_fallthrough),
	3549:  uint16(anon_sym_loop),
	3550:  uint16(anon_sym_for),
	3551:  uint16(anon_sym_while),
	3552:  uint16(anon_sym_break),
	3553:  uint16(anon_sym_continue),
	3554:  uint16(anon_sym_continuing),
	3555:  uint16(anon_sym_return),
	3556:  uint16(anon_sym_discard),
	3557:  uint16(anon_sym_var),
	3558:  uint16(3),
	3559:  uint16(3),
	3560:  uint16(2),
	3561:  uint16(sym_block_comment),
	3562:  uint16(sym_line_comment),
	3563:  uint16(339),
	3564:  uint16(14),
	3566:  uint16(anon_sym_SEMI),
	3567:  uint16(anon_sym_LPAREN),
	3568:  uint16(anon_sym_LBRACE),
	3569:  uint16(anon_sym_RBRACE),
	3570:  uint16(anon_sym_AT),
	3571:  uint16(anon_sym_AMP),
	3572:  uint16(anon_sym_STAR),
	3573:  uint16(aux_sym_preproc_import_token1),
	3574:  uint16(aux_sym_define_import_path_token1),
	3575:  uint16(aux_sym_preproc_ifdef_token1),
	3576:  uint16(aux_sym_preproc_ifdef_token2),
	3577:  uint16(aux_sym_preproc_ifdef_token3),
	3578:  uint16(aux_sym_preproc_else_token1),
	3579:  uint16(341),
	3580:  uint16(21),
	3581:  uint16(anon_sym_let),
	3582:  uint16(anon_sym_override),
	3583:  uint16(anon_sym_type),
	3584:  uint16(anon_sym_virtual),
	3585:  uint16(anon_sym_fn),
	3586:  uint16(anon_sym_struct),
	3587:  uint16(sym_identifier),
	3588:  uint16(anon_sym__),
	3589:  uint16(anon_sym_if),
	3590:  uint16(anon_sym_else),
	3591:  uint16(anon_sym_switch),
	3592:  uint16(anon_sym_fallthrough),
	3593:  uint16(anon_sym_loop),
	3594:  uint16(anon_sym_for),
	3595:  uint16(anon_sym_while),
	3596:  uint16(anon_sym_break),
	3597:  uint16(anon_sym_continue),
	3598:  uint16(anon_sym_continuing),
	3599:  uint16(anon_sym_return),
	3600:  uint16(anon_sym_discard),
	3601:  uint16(anon_sym_var),
	3602:  uint16(20),
	3603:  uint16(9),
	3604:  uint16(1),
	3605:  uint16(anon_sym_let),
	3606:  uint16(11),
	3607:  uint16(1),
	3608:  uint16(anon_sym_override),
	3609:  uint16(13),
	3610:  uint16(1),
	3611:  uint16(anon_sym_type),
	3612:  uint16(15),
	3613:  uint16(1),
	3614:  uint16(anon_sym_virtual),
	3615:  uint16(17),
	3616:  uint16(1),
	3617:  uint16(anon_sym_fn),
	3618:  uint16(19),
	3619:  uint16(1),
	3620:  uint16(anon_sym_struct),
	3621:  uint16(21),
	3622:  uint16(1),
	3623:  uint16(anon_sym_enable),
	3624:  uint16(23),
	3625:  uint16(1),
	3626:  uint16(anon_sym_AT),
	3627:  uint16(25),
	3628:  uint16(1),
	3629:  uint16(anon_sym_var),
	3630:  uint16(27),
	3631:  uint16(1),
	3632:  uint16(aux_sym_preproc_import_token1),
	3633:  uint16(29),
	3634:  uint16(1),
	3635:  uint16(aux_sym_define_import_path_token1),
	3636:  uint16(343),
	3637:  uint16(1),
	3639:  uint16(345),
	3640:  uint16(1),
	3641:  uint16(anon_sym_SEMI),
	3642:  uint16(324),
	3643:  uint16(1),
	3644:  uint16(sym_variable_declaration),
	3645:  uint16(3),
	3646:  uint16(2),
	3647:  uint16(sym_block_comment),
	3648:  uint16(sym_line_comment),
	3649:  uint16(31),
	3650:  uint16(2),
	3651:  uint16(aux_sym_preproc_ifdef_token1),
	3652:  uint16(aux_sym_preproc_ifdef_token2),
	3653:  uint16(151),
	3654:  uint16(2),
	3655:  uint16(sym_enable_directive),
	3656:  uint16(aux_sym_source_file_repeat1),
	3657:  uint16(226),
	3658:  uint16(2),
	3659:  uint16(sym_attribute),
	3660:  uint16(aux_sym_global_variable_declaration_repeat1),
	3661:  uint16(417),
	3662:  uint16(3),
	3663:  uint16(sym_global_variable_declaration),
	3664:  uint16(sym_global_constant_declaration),
	3665:  uint16(sym_type_alias_declaration),
	3666:  uint16(89),
	3667:  uint16(7),
	3668:  uint16(sym__declaration),
	3669:  uint16(sym_function_declaration),
	3670:  uint16(sym_struct_declaration),
	3671:  uint16(sym_preproc_import),
	3672:  uint16(sym_define_import_path),
	3673:  uint16(sym_preproc_ifdef),
	3674:  uint16(aux_sym_source_file_repeat2),
	3675:  uint16(20),
	3676:  uint16(9),
	3677:  uint16(1),
	3678:  uint16(anon_sym_let),
	3679:  uint16(11),
	3680:  uint16(1),
	3681:  uint16(anon_sym_override),
	3682:  uint16(13),
	3683:  uint16(1),
	3684:  uint16(anon_sym_type),
	3685:  uint16(15),
	3686:  uint16(1),
	3687:  uint16(anon_sym_virtual),
	3688:  uint16(17),
	3689:  uint16(1),
	3690:  uint16(anon_sym_fn),
	3691:  uint16(19),
	3692:  uint16(1),
	3693:  uint16(anon_sym_struct),
	3694:  uint16(23),
	3695:  uint16(1),
	3696:  uint16(anon_sym_AT),
	3697:  uint16(25),
	3698:  uint16(1),
	3699:  uint16(anon_sym_var),
	3700:  uint16(27),
	3701:  uint16(1),
	3702:  uint16(aux_sym_preproc_import_token1),
	3703:  uint16(29),
	3704:  uint16(1),
	3705:  uint16(aux_sym_define_import_path_token1),
	3706:  uint16(347),
	3707:  uint16(1),
	3708:  uint16(anon_sym_SEMI),
	3709:  uint16(349),
	3710:  uint16(1),
	3711:  uint16(aux_sym_preproc_ifdef_token3),
	3712:  uint16(351),
	3713:  uint16(1),
	3714:  uint16(aux_sym_preproc_else_token1),
	3715:  uint16(324),
	3716:  uint16(1),
	3717:  uint16(sym_variable_declaration),
	3718:  uint16(387),
	3719:  uint16(1),
	3720:  uint16(sym_preproc_else),
	3721:  uint16(3),
	3722:  uint16(2),
	3723:  uint16(sym_block_comment),
	3724:  uint16(sym_line_comment),
	3725:  uint16(31),
	3726:  uint16(2),
	3727:  uint16(aux_sym_preproc_ifdef_token1),
	3728:  uint16(aux_sym_preproc_ifdef_token2),
	3729:  uint16(226),
	3730:  uint16(2),
	3731:  uint16(sym_attribute),
	3732:  uint16(aux_sym_global_variable_declaration_repeat1),
	3733:  uint16(417),
	3734:  uint16(3),
	3735:  uint16(sym_global_variable_declaration),
	3736:  uint16(sym_global_constant_declaration),
	3737:  uint16(sym_type_alias_declaration),
	3738:  uint16(84),
	3739:  uint16(7),
	3740:  uint16(sym__declaration),
	3741:  uint16(sym_function_declaration),
	3742:  uint16(sym_struct_declaration),
	3743:  uint16(sym_preproc_import),
	3744:  uint16(sym_define_import_path),
	3745:  uint16(sym_preproc_ifdef),
	3746:  uint16(aux_sym_source_file_repeat2),
	3747:  uint16(20),
	3748:  uint16(9),
	3749:  uint16(1),
	3750:  uint16(anon_sym_let),
	3751:  uint16(11),
	3752:  uint16(1),
	3753:  uint16(anon_sym_override),
	3754:  uint16(13),
	3755:  uint16(1),
	3756:  uint16(anon_sym_type),
	3757:  uint16(15),
	3758:  uint16(1),
	3759:  uint16(anon_sym_virtual),
	3760:  uint16(17),
	3761:  uint16(1),
	3762:  uint16(anon_sym_fn),
	3763:  uint16(19),
	3764:  uint16(1),
	3765:  uint16(anon_sym_struct),
	3766:  uint16(23),
	3767:  uint16(1),
	3768:  uint16(anon_sym_AT),
	3769:  uint16(25),
	3770:  uint16(1),
	3771:  uint16(anon_sym_var),
	3772:  uint16(27),
	3773:  uint16(1),
	3774:  uint16(aux_sym_preproc_import_token1),
	3775:  uint16(29),
	3776:  uint16(1),
	3777:  uint16(aux_sym_define_import_path_token1),
	3778:  uint16(351),
	3779:  uint16(1),
	3780:  uint16(aux_sym_preproc_else_token1),
	3781:  uint16(353),
	3782:  uint16(1),
	3783:  uint16(anon_sym_SEMI),
	3784:  uint16(355),
	3785:  uint16(1),
	3786:  uint16(aux_sym_preproc_ifdef_token3),
	3787:  uint16(324),
	3788:  uint16(1),
	3789:  uint16(sym_variable_declaration),
	3790:  uint16(407),
	3791:  uint16(1),
	3792:  uint16(sym_preproc_else),
	3793:  uint16(3),
	3794:  uint16(2),
	3795:  uint16(sym_block_comment),
	3796:  uint16(sym_line_comment),
	3797:  uint16(31),
	3798:  uint16(2),
	3799:  uint16(aux_sym_preproc_ifdef_token1),
	3800:  uint16(aux_sym_preproc_ifdef_token2),
	3801:  uint16(226),
	3802:  uint16(2),
	3803:  uint16(sym_attribute),
	3804:  uint16(aux_sym_global_variable_declaration_repeat1),
	3805:  uint16(417),
	3806:  uint16(3),
	3807:  uint16(sym_global_variable_declaration),
	3808:  uint16(sym_global_constant_declaration),
	3809:  uint16(sym_type_alias_declaration),
	3810:  uint16(82),
	3811:  uint16(7),
	3812:  uint16(sym__declaration),
	3813:  uint16(sym_function_declaration),
	3814:  uint16(sym_struct_declaration),
	3815:  uint16(sym_preproc_import),
	3816:  uint16(sym_define_import_path),
	3817:  uint16(sym_preproc_ifdef),
	3818:  uint16(aux_sym_source_file_repeat2),
	3819:  uint16(18),
	3820:  uint16(359),
	3821:  uint16(1),
	3822:  uint16(anon_sym_SEMI),
	3823:  uint16(362),
	3824:  uint16(1),
	3825:  uint16(anon_sym_let),
	3826:  uint16(365),
	3827:  uint16(1),
	3828:  uint16(anon_sym_override),
	3829:  uint16(368),
	3830:  uint16(1),
	3831:  uint16(anon_sym_type),
	3832:  uint16(371),
	3833:  uint16(1),
	3834:  uint16(anon_sym_virtual),
	3835:  uint16(374),
	3836:  uint16(1),
	3837:  uint16(anon_sym_fn),
	3838:  uint16(377),
	3839:  uint16(1),
	3840:  uint16(anon_sym_struct),
	3841:  uint16(380),
	3842:  uint16(1),
	3843:  uint16(anon_sym_AT),
	3844:  uint16(383),
	3845:  uint16(1),
	3846:  uint16(anon_sym_var),
	3847:  uint16(386),
	3848:  uint16(1),
	3849:  uint16(aux_sym_preproc_import_token1),
	3850:  uint16(389),
	3851:  uint16(1),
	3852:  uint16(aux_sym_define_import_path_token1),
	3853:  uint16(324),
	3854:  uint16(1),
	3855:  uint16(sym_variable_declaration),
	3856:  uint16(3),
	3857:  uint16(2),
	3858:  uint16(sym_block_comment),
	3859:  uint16(sym_line_comment),
	3860:  uint16(392),
	3861:  uint16(2),
	3862:  uint16(aux_sym_preproc_ifdef_token1),
	3863:  uint16(aux_sym_preproc_ifdef_token2),
	3864:  uint16(226),
	3865:  uint16(2),
	3866:  uint16(sym_attribute),
	3867:  uint16(aux_sym_global_variable_declaration_repeat1),
	3868:  uint16(357),
	3869:  uint16(3),
	3871:  uint16(aux_sym_preproc_ifdef_token3),
	3872:  uint16(aux_sym_preproc_else_token1),
	3873:  uint16(417),
	3874:  uint16(3),
	3875:  uint16(sym_global_variable_declaration),
	3876:  uint16(sym_global_constant_declaration),
	3877:  uint16(sym_type_alias_declaration),
	3878:  uint16(84),
	3879:  uint16(7),
	3880:  uint16(sym__declaration),
	3881:  uint16(sym_function_declaration),
	3882:  uint16(sym_struct_declaration),
	3883:  uint16(sym_preproc_import),
	3884:  uint16(sym_define_import_path),
	3885:  uint16(sym_preproc_ifdef),
	3886:  uint16(aux_sym_source_file_repeat2),
	3887:  uint16(18),
	3888:  uint16(9),
	3889:  uint16(1),
	3890:  uint16(anon_sym_let),
	3891:  uint16(11),
	3892:  uint16(1),
	3893:  uint16(anon_sym_override),
	3894:  uint16(13),
	3895:  uint16(1),
	3896:  uint16(anon_sym_type),
	3897:  uint16(15),
	3898:  uint16(1),
	3899:  uint16(anon_sym_virtual),
	3900:  uint16(17),
	3901:  uint16(1),
	3902:  uint16(anon_sym_fn),
	3903:  uint16(19),
	3904:  uint16(1),
	3905:  uint16(anon_sym_struct),
	3906:  uint16(23),
	3907:  uint16(1),
	3908:  uint16(anon_sym_AT),
	3909:  uint16(25),
	3910:  uint16(1),
	3911:  uint16(anon_sym_var),
	3912:  uint16(27),
	3913:  uint16(1),
	3914:  uint16(aux_sym_preproc_import_token1),
	3915:  uint16(29),
	3916:  uint16(1),
	3917:  uint16(aux_sym_define_import_path_token1),
	3918:  uint16(395),
	3919:  uint16(1),
	3920:  uint16(anon_sym_SEMI),
	3921:  uint16(397),
	3922:  uint16(1),
	3923:  uint16(aux_sym_preproc_ifdef_token3),
	3924:  uint16(324),
	3925:  uint16(1),
	3926:  uint16(sym_variable_declaration),
	3927:  uint16(3),
	3928:  uint16(2),
	3929:  uint16(sym_block_comment),
	3930:  uint16(sym_line_comment),
	3931:  uint16(31),
	3932:  uint16(2),
	3933:  uint16(aux_sym_preproc_ifdef_token1),
	3934:  uint16(aux_sym_preproc_ifdef_token2),
	3935:  uint16(226),
	3936:  uint16(2),
	3937:  uint16(sym_attribute),
	3938:  uint16(aux_sym_global_variable_declaration_repeat1),
	3939:  uint16(417),
	3940:  uint16(3),
	3941:  uint16(sym_global_variable_declaration),
	3942:  uint16(sym_global_constant_declaration),
	3943:  uint16(sym_type_alias_declaration),
	3944:  uint16(87),
	3945:  uint16(7),
	3946:  uint16(sym__declaration),
	3947:  uint16(sym_function_declaration),
	3948:  uint16(sym_struct_declaration),
	3949:  uint16(sym_preproc_import),
	3950:  uint16(sym_define_import_path),
	3951:  uint16(sym_preproc_ifdef),
	3952:  uint16(aux_sym_source_file_repeat2),
	3953:  uint16(3),
	3954:  uint16(3),
	3955:  uint16(2),
	3956:  uint16(sym_block_comment),
	3957:  uint16(sym_line_comment),
	3958:  uint16(401),
	3959:  uint16(12),
	3960:  uint16(anon_sym_SEMI),
	3961:  uint16(anon_sym_LPAREN),
	3962:  uint16(anon_sym_RPAREN),
	3963:  uint16(anon_sym_LBRACE),
	3964:  uint16(anon_sym_RBRACE),
	3965:  uint16(anon_sym_AMP),
	3966:  uint16(anon_sym_STAR),
	3967:  uint16(aux_sym_preproc_import_token1),
	3968:  uint16(aux_sym_preproc_ifdef_token1),
	3969:  uint16(aux_sym_preproc_ifdef_token2),
	3970:  uint16(aux_sym_preproc_ifdef_token3),
	3971:  uint16(aux_sym_preproc_else_token1),
	3972:  uint16(399),
	3973:  uint16(15),
	3974:  uint16(anon_sym_let),
	3975:  uint16(sym_identifier),
	3976:  uint16(anon_sym__),
	3977:  uint16(anon_sym_if),
	3978:  uint16(anon_sym_switch),
	3979:  uint16(anon_sym_fallthrough),
	3980:  uint16(anon_sym_loop),
	3981:  uint16(anon_sym_for),
	3982:  uint16(anon_sym_while),
	3983:  uint16(anon_sym_break),
	3984:  uint16(anon_sym_continue),
	3985:  uint16(anon_sym_continuing),
	3986:  uint16(anon_sym_return),
	3987:  uint16(anon_sym_discard),
	3988:  uint16(anon_sym_var),
	3989:  uint16(18),
	3990:  uint16(9),
	3991:  uint16(1),
	3992:  uint16(anon_sym_let),
	3993:  uint16(11),
	3994:  uint16(1),
	3995:  uint16(anon_sym_override),
	3996:  uint16(13),
	3997:  uint16(1),
	3998:  uint16(anon_sym_type),
	3999:  uint16(15),
	4000:  uint16(1),
	4001:  uint16(anon_sym_virtual),
	4002:  uint16(17),
	4003:  uint16(1),
	4004:  uint16(anon_sym_fn),
	4005:  uint16(19),
	4006:  uint16(1),
	4007:  uint16(anon_sym_struct),
	4008:  uint16(23),
	4009:  uint16(1),
	4010:  uint16(anon_sym_AT),
	4011:  uint16(25),
	4012:  uint16(1),
	4013:  uint16(anon_sym_var),
	4014:  uint16(27),
	4015:  uint16(1),
	4016:  uint16(aux_sym_preproc_import_token1),
	4017:  uint16(29),
	4018:  uint16(1),
	4019:  uint16(aux_sym_define_import_path_token1),
	4020:  uint16(347),
	4021:  uint16(1),
	4022:  uint16(anon_sym_SEMI),
	4023:  uint16(403),
	4024:  uint16(1),
	4025:  uint16(aux_sym_preproc_ifdef_token3),
	4026:  uint16(324),
	4027:  uint16(1),
	4028:  uint16(sym_variable_declaration),
	4029:  uint16(3),
	4030:  uint16(2),
	4031:  uint16(sym_block_comment),
	4032:  uint16(sym_line_comment),
	4033:  uint16(31),
	4034:  uint16(2),
	4035:  uint16(aux_sym_preproc_ifdef_token1),
	4036:  uint16(aux_sym_preproc_ifdef_token2),
	4037:  uint16(226),
	4038:  uint16(2),
	4039:  uint16(sym_attribute),
	4040:  uint16(aux_sym_global_variable_declaration_repeat1),
	4041:  uint16(417),
	4042:  uint16(3),
	4043:  uint16(sym_global_variable_declaration),
	4044:  uint16(sym_global_constant_declaration),
	4045:  uint16(sym_type_alias_declaration),
	4046:  uint16(84),
	4047:  uint16(7),
	4048:  uint16(sym__declaration),
	4049:  uint16(sym_function_declaration),
	4050:  uint16(sym_struct_declaration),
	4051:  uint16(sym_preproc_import),
	4052:  uint16(sym_define_import_path),
	4053:  uint16(sym_preproc_ifdef),
	4054:  uint16(aux_sym_source_file_repeat2),
	4055:  uint16(18),
	4056:  uint16(9),
	4057:  uint16(1),
	4058:  uint16(anon_sym_let),
	4059:  uint16(11),
	4060:  uint16(1),
	4061:  uint16(anon_sym_override),
	4062:  uint16(13),
	4063:  uint16(1),
	4064:  uint16(anon_sym_type),
	4065:  uint16(15),
	4066:  uint16(1),
	4067:  uint16(anon_sym_virtual),
	4068:  uint16(17),
	4069:  uint16(1),
	4070:  uint16(anon_sym_fn),
	4071:  uint16(19),
	4072:  uint16(1),
	4073:  uint16(anon_sym_struct),
	4074:  uint16(23),
	4075:  uint16(1),
	4076:  uint16(anon_sym_AT),
	4077:  uint16(25),
	4078:  uint16(1),
	4079:  uint16(anon_sym_var),
	4080:  uint16(27),
	4081:  uint16(1),
	4082:  uint16(aux_sym_preproc_import_token1),
	4083:  uint16(29),
	4084:  uint16(1),
	4085:  uint16(aux_sym_define_import_path_token1),
	4086:  uint16(343),
	4087:  uint16(1),
	4089:  uint16(347),
	4090:  uint16(1),
	4091:  uint16(anon_sym_SEMI),
	4092:  uint16(324),
	4093:  uint16(1),
	4094:  uint16(sym_variable_declaration),
	4095:  uint16(3),
	4096:  uint16(2),
	4097:  uint16(sym_block_comment),
	4098:  uint16(sym_line_comment),
	4099:  uint16(31),
	4100:  uint16(2),
	4101:  uint16(aux_sym_preproc_ifdef_token1),
	4102:  uint16(aux_sym_preproc_ifdef_token2),
	4103:  uint16(226),
	4104:  uint16(2),
	4105:  uint16(sym_attribute),
	4106:  uint16(aux_sym_global_variable_declaration_repeat1),
	4107:  uint16(417),
	4108:  uint16(3),
	4109:  uint16(sym_global_variable_declaration),
	4110:  uint16(sym_global_constant_declaration),
	4111:  uint16(sym_type_alias_declaration),
	4112:  uint16(84),
	4113:  uint16(7),
	4114:  uint16(sym__declaration),
	4115:  uint16(sym_function_declaration),
	4116:  uint16(sym_struct_declaration),
	4117:  uint16(sym_preproc_import),
	4118:  uint16(sym_define_import_path),
	4119:  uint16(sym_preproc_ifdef),
	4120:  uint16(aux_sym_source_file_repeat2),
	4121:  uint16(18),
	4122:  uint16(9),
	4123:  uint16(1),
	4124:  uint16(anon_sym_let),
	4125:  uint16(11),
	4126:  uint16(1),
	4127:  uint16(anon_sym_override),
	4128:  uint16(13),
	4129:  uint16(1),
	4130:  uint16(anon_sym_type),
	4131:  uint16(15),
	4132:  uint16(1),
	4133:  uint16(anon_sym_virtual),
	4134:  uint16(17),
	4135:  uint16(1),
	4136:  uint16(anon_sym_fn),
	4137:  uint16(19),
	4138:  uint16(1),
	4139:  uint16(anon_sym_struct),
	4140:  uint16(23),
	4141:  uint16(1),
	4142:  uint16(anon_sym_AT),
	4143:  uint16(25),
	4144:  uint16(1),
	4145:  uint16(anon_sym_var),
	4146:  uint16(27),
	4147:  uint16(1),
	4148:  uint16(aux_sym_preproc_import_token1),
	4149:  uint16(29),
	4150:  uint16(1),
	4151:  uint16(aux_sym_define_import_path_token1),
	4152:  uint16(347),
	4153:  uint16(1),
	4154:  uint16(anon_sym_SEMI),
	4155:  uint16(405),
	4156:  uint16(1),
	4158:  uint16(324),
	4159:  uint16(1),
	4160:  uint16(sym_variable_declaration),
	4161:  uint16(3),
	4162:  uint16(2),
	4163:  uint16(sym_block_comment),
	4164:  uint16(sym_line_comment),
	4165:  uint16(31),
	4166:  uint16(2),
	4167:  uint16(aux_sym_preproc_ifdef_token1),
	4168:  uint16(aux_sym_preproc_ifdef_token2),
	4169:  uint16(226),
	4170:  uint16(2),
	4171:  uint16(sym_attribute),
	4172:  uint16(aux_sym_global_variable_declaration_repeat1),
	4173:  uint16(417),
	4174:  uint16(3),
	4175:  uint16(sym_global_variable_declaration),
	4176:  uint16(sym_global_constant_declaration),
	4177:  uint16(sym_type_alias_declaration),
	4178:  uint16(84),
	4179:  uint16(7),
	4180:  uint16(sym__declaration),
	4181:  uint16(sym_function_declaration),
	4182:  uint16(sym_struct_declaration),
	4183:  uint16(sym_preproc_import),
	4184:  uint16(sym_define_import_path),
	4185:  uint16(sym_preproc_ifdef),
	4186:  uint16(aux_sym_source_file_repeat2),
	4187:  uint16(3),
	4188:  uint16(3),
	4189:  uint16(2),
	4190:  uint16(sym_block_comment),
	4191:  uint16(sym_line_comment),
	4192:  uint16(409),
	4193:  uint16(12),
	4194:  uint16(anon_sym_SEMI),
	4195:  uint16(anon_sym_LPAREN),
	4196:  uint16(anon_sym_RPAREN),
	4197:  uint16(anon_sym_LBRACE),
	4198:  uint16(anon_sym_RBRACE),
	4199:  uint16(anon_sym_AMP),
	4200:  uint16(anon_sym_STAR),
	4201:  uint16(aux_sym_preproc_import_token1),
	4202:  uint16(aux_sym_preproc_ifdef_token1),
	4203:  uint16(aux_sym_preproc_ifdef_token2),
	4204:  uint16(aux_sym_preproc_ifdef_token3),
	4205:  uint16(aux_sym_preproc_else_token1),
	4206:  uint16(407),
	4207:  uint16(15),
	4208:  uint16(anon_sym_let),
	4209:  uint16(sym_identifier),
	4210:  uint16(anon_sym__),
	4211:  uint16(anon_sym_if),
	4212:  uint16(anon_sym_switch),
	4213:  uint16(anon_sym_fallthrough),
	4214:  uint16(anon_sym_loop),
	4215:  uint16(anon_sym_for),
	4216:  uint16(anon_sym_while),
	4217:  uint16(anon_sym_break),
	4218:  uint16(anon_sym_continue),
	4219:  uint16(anon_sym_continuing),
	4220:  uint16(anon_sym_return),
	4221:  uint16(anon_sym_discard),
	4222:  uint16(anon_sym_var),
	4223:  uint16(3),
	4224:  uint16(3),
	4225:  uint16(2),
	4226:  uint16(sym_block_comment),
	4227:  uint16(sym_line_comment),
	4228:  uint16(413),
	4229:  uint16(5),
	4230:  uint16(anon_sym_LT),
	4231:  uint16(anon_sym_GT),
	4232:  uint16(anon_sym_PIPE),
	4233:  uint16(anon_sym_AMP),
	4234:  uint16(anon_sym_SLASH),
	4235:  uint16(411),
	4236:  uint16(21),
	4237:  uint16(anon_sym_SEMI),
	4238:  uint16(anon_sym_COMMA),
	4239:  uint16(anon_sym_RPAREN),
	4240:  uint16(anon_sym_LBRACE),
	4241:  uint16(anon_sym_COLON),
	4242:  uint16(anon_sym_PIPE_PIPE),
	4243:  uint16(anon_sym_AMP_AMP),
	4244:  uint16(anon_sym_CARET),
	4245:  uint16(anon_sym_EQ_EQ),
	4246:  uint16(anon_sym_BANG_EQ),
	4247:  uint16(anon_sym_LT_EQ),
	4248:  uint16(anon_sym_GT_EQ),
	4249:  uint16(anon_sym_LT_LT),
	4250:  uint16(anon_sym_GT_GT),
	4251:  uint16(anon_sym_PLUS),
	4252:  uint16(anon_sym_DASH),
	4253:  uint16(anon_sym_STAR),
	4254:  uint16(anon_sym_PERCENT),
	4255:  uint16(anon_sym_LBRACK),
	4256:  uint16(anon_sym_RBRACK),
	4257:  uint16(anon_sym_DOT),
	4258:  uint16(3),
	4259:  uint16(3),
	4260:  uint16(2),
	4261:  uint16(sym_block_comment),
	4262:  uint16(sym_line_comment),
	4263:  uint16(417),
	4264:  uint16(5),
	4265:  uint16(anon_sym_LT),
	4266:  uint16(anon_sym_GT),
	4267:  uint16(anon_sym_PIPE),
	4268:  uint16(anon_sym_AMP),
	4269:  uint16(anon_sym_SLASH),
	4270:  uint16(415),
	4271:  uint16(21),
	4272:  uint16(anon_sym_SEMI),
	4273:  uint16(anon_sym_COMMA),
	4274:  uint16(anon_sym_RPAREN),
	4275:  uint16(anon_sym_LBRACE),
	4276:  uint16(anon_sym_COLON),
	4277:  uint16(anon_sym_PIPE_PIPE),
	4278:  uint16(anon_sym_AMP_AMP),
	4279:  uint16(anon_sym_CARET),
	4280:  uint16(anon_sym_EQ_EQ),
	4281:  uint16(anon_sym_BANG_EQ),
	4282:  uint16(anon_sym_LT_EQ),
	4283:  uint16(anon_sym_GT_EQ),
	4284:  uint16(anon_sym_LT_LT),
	4285:  uint16(anon_sym_GT_GT),
	4286:  uint16(anon_sym_PLUS),
	4287:  uint16(anon_sym_DASH),
	4288:  uint16(anon_sym_STAR),
	4289:  uint16(anon_sym_PERCENT),
	4290:  uint16(anon_sym_LBRACK),
	4291:  uint16(anon_sym_RBRACK),
	4292:  uint16(anon_sym_DOT),
	4293:  uint16(3),
	4294:  uint16(3),
	4295:  uint16(2),
	4296:  uint16(sym_block_comment),
	4297:  uint16(sym_line_comment),
	4298:  uint16(421),
	4299:  uint16(5),
	4300:  uint16(anon_sym_LT),
	4301:  uint16(anon_sym_GT),
	4302:  uint16(anon_sym_PIPE),
	4303:  uint16(anon_sym_AMP),
	4304:  uint16(anon_sym_SLASH),
	4305:  uint16(419),
	4306:  uint16(21),
	4307:  uint16(anon_sym_SEMI),
	4308:  uint16(anon_sym_COMMA),
	4309:  uint16(anon_sym_RPAREN),
	4310:  uint16(anon_sym_LBRACE),
	4311:  uint16(anon_sym_COLON),
	4312:  uint16(anon_sym_PIPE_PIPE),
	4313:  uint16(anon_sym_AMP_AMP),
	4314:  uint16(anon_sym_CARET),
	4315:  uint16(anon_sym_EQ_EQ),
	4316:  uint16(anon_sym_BANG_EQ),
	4317:  uint16(anon_sym_LT_EQ),
	4318:  uint16(anon_sym_GT_EQ),
	4319:  uint16(anon_sym_LT_LT),
	4320:  uint16(anon_sym_GT_GT),
	4321:  uint16(anon_sym_PLUS),
	4322:  uint16(anon_sym_DASH),
	4323:  uint16(anon_sym_STAR),
	4324:  uint16(anon_sym_PERCENT),
	4325:  uint16(anon_sym_LBRACK),
	4326:  uint16(anon_sym_RBRACK),
	4327:  uint16(anon_sym_DOT),
	4328:  uint16(4),
	4329:  uint16(427),
	4330:  uint16(1),
	4331:  uint16(anon_sym_else),
	4332:  uint16(3),
	4333:  uint16(2),
	4334:  uint16(sym_block_comment),
	4335:  uint16(sym_line_comment),
	4336:  uint16(425),
	4337:  uint16(10),
	4338:  uint16(anon_sym_LPAREN),
	4339:  uint16(anon_sym_LBRACE),
	4340:  uint16(anon_sym_RBRACE),
	4341:  uint16(anon_sym_AMP),
	4342:  uint16(anon_sym_STAR),
	4343:  uint16(aux_sym_preproc_import_token1),
	4344:  uint16(aux_sym_preproc_ifdef_token1),
	4345:  uint16(aux_sym_preproc_ifdef_token2),
	4346:  uint16(aux_sym_preproc_ifdef_token3),
	4347:  uint16(aux_sym_preproc_else_token1),
	4348:  uint16(423),
	4349:  uint16(15),
	4350:  uint16(anon_sym_let),
	4351:  uint16(sym_identifier),
	4352:  uint16(anon_sym__),
	4353:  uint16(anon_sym_if),
	4354:  uint16(anon_sym_switch),
	4355:  uint16(anon_sym_fallthrough),
	4356:  uint16(anon_sym_loop),
	4357:  uint16(anon_sym_for),
	4358:  uint16(anon_sym_while),
	4359:  uint16(anon_sym_break),
	4360:  uint16(anon_sym_continue),
	4361:  uint16(anon_sym_continuing),
	4362:  uint16(anon_sym_return),
	4363:  uint16(anon_sym_discard),
	4364:  uint16(anon_sym_var),
	4365:  uint16(4),
	4366:  uint16(431),
	4367:  uint16(1),
	4368:  uint16(anon_sym_LPAREN),
	4369:  uint16(3),
	4370:  uint16(2),
	4371:  uint16(sym_block_comment),
	4372:  uint16(sym_line_comment),
	4373:  uint16(433),
	4374:  uint16(5),
	4375:  uint16(anon_sym_LT),
	4376:  uint16(anon_sym_GT),
	4377:  uint16(anon_sym_PIPE),
	4378:  uint16(anon_sym_AMP),
	4379:  uint16(anon_sym_SLASH),
	4380:  uint16(429),
	4381:  uint16(20),
	4382:  uint16(anon_sym_SEMI),
	4383:  uint16(anon_sym_COMMA),
	4384:  uint16(anon_sym_RPAREN),
	4385:  uint16(anon_sym_LBRACE),
	4386:  uint16(anon_sym_PIPE_PIPE),
	4387:  uint16(anon_sym_AMP_AMP),
	4388:  uint16(anon_sym_CARET),
	4389:  uint16(anon_sym_EQ_EQ),
	4390:  uint16(anon_sym_BANG_EQ),
	4391:  uint16(anon_sym_LT_EQ),
	4392:  uint16(anon_sym_GT_EQ),
	4393:  uint16(anon_sym_LT_LT),
	4394:  uint16(anon_sym_GT_GT),
	4395:  uint16(anon_sym_PLUS),
	4396:  uint16(anon_sym_DASH),
	4397:  uint16(anon_sym_STAR),
	4398:  uint16(anon_sym_PERCENT),
	4399:  uint16(anon_sym_LBRACK),
	4400:  uint16(anon_sym_RBRACK),
	4401:  uint16(anon_sym_DOT),
	4402:  uint16(3),
	4403:  uint16(3),
	4404:  uint16(2),
	4405:  uint16(sym_block_comment),
	4406:  uint16(sym_line_comment),
	4407:  uint16(437),
	4408:  uint16(5),
	4409:  uint16(anon_sym_LT),
	4410:  uint16(anon_sym_GT),
	4411:  uint16(anon_sym_PIPE),
	4412:  uint16(anon_sym_AMP),
	4413:  uint16(anon_sym_SLASH),
	4414:  uint16(435),
	4415:  uint16(20),
	4416:  uint16(anon_sym_SEMI),
	4417:  uint16(anon_sym_COMMA),
	4418:  uint16(anon_sym_RPAREN),
	4419:  uint16(anon_sym_LBRACE),
	4420:  uint16(anon_sym_PIPE_PIPE),
	4421:  uint16(anon_sym_AMP_AMP),
	4422:  uint16(anon_sym_CARET),
	4423:  uint16(anon_sym_EQ_EQ),
	4424:  uint16(anon_sym_BANG_EQ),
	4425:  uint16(anon_sym_LT_EQ),
	4426:  uint16(anon_sym_GT_EQ),
	4427:  uint16(anon_sym_LT_LT),
	4428:  uint16(anon_sym_GT_GT),
	4429:  uint16(anon_sym_PLUS),
	4430:  uint16(anon_sym_DASH),
	4431:  uint16(anon_sym_STAR),
	4432:  uint16(anon_sym_PERCENT),
	4433:  uint16(anon_sym_LBRACK),
	4434:  uint16(anon_sym_RBRACK),
	4435:  uint16(anon_sym_DOT),
	4436:  uint16(5),
	4437:  uint16(443),
	4438:  uint16(1),
	4439:  uint16(anon_sym_LBRACK),
	4440:  uint16(445),
	4441:  uint16(1),
	4442:  uint16(anon_sym_DOT),
	4443:  uint16(3),
	4444:  uint16(2),
	4445:  uint16(sym_block_comment),
	4446:  uint16(sym_line_comment),
	4447:  uint16(441),
	4448:  uint16(5),
	4449:  uint16(anon_sym_LT),
	4450:  uint16(anon_sym_GT),
	4451:  uint16(anon_sym_PIPE),
	4452:  uint16(anon_sym_AMP),
	4453:  uint16(anon_sym_SLASH),
	4454:  uint16(439),
	4455:  uint16(18),
	4456:  uint16(anon_sym_SEMI),
	4457:  uint16(anon_sym_COMMA),
	4458:  uint16(anon_sym_RPAREN),
	4459:  uint16(anon_sym_LBRACE),
	4460:  uint16(anon_sym_PIPE_PIPE),
	4461:  uint16(anon_sym_AMP_AMP),
	4462:  uint16(anon_sym_CARET),
	4463:  uint16(anon_sym_EQ_EQ),
	4464:  uint16(anon_sym_BANG_EQ),
	4465:  uint16(anon_sym_LT_EQ),
	4466:  uint16(anon_sym_GT_EQ),
	4467:  uint16(anon_sym_LT_LT),
	4468:  uint16(anon_sym_GT_GT),
	4469:  uint16(anon_sym_PLUS),
	4470:  uint16(anon_sym_DASH),
	4471:  uint16(anon_sym_STAR),
	4472:  uint16(anon_sym_PERCENT),
	4473:  uint16(anon_sym_RBRACK),
	4474:  uint16(3),
	4475:  uint16(3),
	4476:  uint16(2),
	4477:  uint16(sym_block_comment),
	4478:  uint16(sym_line_comment),
	4479:  uint16(449),
	4480:  uint16(5),
	4481:  uint16(anon_sym_LT),
	4482:  uint16(anon_sym_GT),
	4483:  uint16(anon_sym_PIPE),
	4484:  uint16(anon_sym_AMP),
	4485:  uint16(anon_sym_SLASH),
	4486:  uint16(447),
	4487:  uint16(20),
	4488:  uint16(anon_sym_SEMI),
	4489:  uint16(anon_sym_COMMA),
	4490:  uint16(anon_sym_RPAREN),
	4491:  uint16(anon_sym_LBRACE),
	4492:  uint16(anon_sym_PIPE_PIPE),
	4493:  uint16(anon_sym_AMP_AMP),
	4494:  uint16(anon_sym_CARET),
	4495:  uint16(anon_sym_EQ_EQ),
	4496:  uint16(anon_sym_BANG_EQ),
	4497:  uint16(anon_sym_LT_EQ),
	4498:  uint16(anon_sym_GT_EQ),
	4499:  uint16(anon_sym_LT_LT),
	4500:  uint16(anon_sym_GT_GT),
	4501:  uint16(anon_sym_PLUS),
	4502:  uint16(anon_sym_DASH),
	4503:  uint16(anon_sym_STAR),
	4504:  uint16(anon_sym_PERCENT),
	4505:  uint16(anon_sym_LBRACK),
	4506:  uint16(anon_sym_RBRACK),
	4507:  uint16(anon_sym_DOT),
	4508:  uint16(3),
	4509:  uint16(3),
	4510:  uint16(2),
	4511:  uint16(sym_block_comment),
	4512:  uint16(sym_line_comment),
	4513:  uint16(453),
	4514:  uint16(5),
	4515:  uint16(anon_sym_LT),
	4516:  uint16(anon_sym_GT),
	4517:  uint16(anon_sym_PIPE),
	4518:  uint16(anon_sym_AMP),
	4519:  uint16(anon_sym_SLASH),
	4520:  uint16(451),
	4521:  uint16(20),
	4522:  uint16(anon_sym_SEMI),
	4523:  uint16(anon_sym_COMMA),
	4524:  uint16(anon_sym_RPAREN),
	4525:  uint16(anon_sym_LBRACE),
	4526:  uint16(anon_sym_PIPE_PIPE),
	4527:  uint16(anon_sym_AMP_AMP),
	4528:  uint16(anon_sym_CARET),
	4529:  uint16(anon_sym_EQ_EQ),
	4530:  uint16(anon_sym_BANG_EQ),
	4531:  uint16(anon_sym_LT_EQ),
	4532:  uint16(anon_sym_GT_EQ),
	4533:  uint16(anon_sym_LT_LT),
	4534:  uint16(anon_sym_GT_GT),
	4535:  uint16(anon_sym_PLUS),
	4536:  uint16(anon_sym_DASH),
	4537:  uint16(anon_sym_STAR),
	4538:  uint16(anon_sym_PERCENT),
	4539:  uint16(anon_sym_LBRACK),
	4540:  uint16(anon_sym_RBRACK),
	4541:  uint16(anon_sym_DOT),
	4542:  uint16(3),
	4543:  uint16(3),
	4544:  uint16(2),
	4545:  uint16(sym_block_comment),
	4546:  uint16(sym_line_comment),
	4547:  uint16(457),
	4548:  uint16(5),
	4549:  uint16(anon_sym_LT),
	4550:  uint16(anon_sym_GT),
	4551:  uint16(anon_sym_PIPE),
	4552:  uint16(anon_sym_AMP),
	4553:  uint16(anon_sym_SLASH),
	4554:  uint16(455),
	4555:  uint16(20),
	4556:  uint16(anon_sym_SEMI),
	4557:  uint16(anon_sym_COMMA),
	4558:  uint16(anon_sym_RPAREN),
	4559:  uint16(anon_sym_LBRACE),
	4560:  uint16(anon_sym_PIPE_PIPE),
	4561:  uint16(anon_sym_AMP_AMP),
	4562:  uint16(anon_sym_CARET),
	4563:  uint16(anon_sym_EQ_EQ),
	4564:  uint16(anon_sym_BANG_EQ),
	4565:  uint16(anon_sym_LT_EQ),
	4566:  uint16(anon_sym_GT_EQ),
	4567:  uint16(anon_sym_LT_LT),
	4568:  uint16(anon_sym_GT_GT),
	4569:  uint16(anon_sym_PLUS),
	4570:  uint16(anon_sym_DASH),
	4571:  uint16(anon_sym_STAR),
	4572:  uint16(anon_sym_PERCENT),
	4573:  uint16(anon_sym_LBRACK),
	4574:  uint16(anon_sym_RBRACK),
	4575:  uint16(anon_sym_DOT),
	4576:  uint16(3),
	4577:  uint16(3),
	4578:  uint16(2),
	4579:  uint16(sym_block_comment),
	4580:  uint16(sym_line_comment),
	4581:  uint16(461),
	4582:  uint16(5),
	4583:  uint16(anon_sym_LT),
	4584:  uint16(anon_sym_GT),
	4585:  uint16(anon_sym_PIPE),
	4586:  uint16(anon_sym_AMP),
	4587:  uint16(anon_sym_SLASH),
	4588:  uint16(459),
	4589:  uint16(20),
	4590:  uint16(anon_sym_SEMI),
	4591:  uint16(anon_sym_COMMA),
	4592:  uint16(anon_sym_RPAREN),
	4593:  uint16(anon_sym_LBRACE),
	4594:  uint16(anon_sym_PIPE_PIPE),
	4595:  uint16(anon_sym_AMP_AMP),
	4596:  uint16(anon_sym_CARET),
	4597:  uint16(anon_sym_EQ_EQ),
	4598:  uint16(anon_sym_BANG_EQ),
	4599:  uint16(anon_sym_LT_EQ),
	4600:  uint16(anon_sym_GT_EQ),
	4601:  uint16(anon_sym_LT_LT),
	4602:  uint16(anon_sym_GT_GT),
	4603:  uint16(anon_sym_PLUS),
	4604:  uint16(anon_sym_DASH),
	4605:  uint16(anon_sym_STAR),
	4606:  uint16(anon_sym_PERCENT),
	4607:  uint16(anon_sym_LBRACK),
	4608:  uint16(anon_sym_RBRACK),
	4609:  uint16(anon_sym_DOT),
	4610:  uint16(3),
	4611:  uint16(3),
	4612:  uint16(2),
	4613:  uint16(sym_block_comment),
	4614:  uint16(sym_line_comment),
	4615:  uint16(465),
	4616:  uint16(10),
	4617:  uint16(anon_sym_LPAREN),
	4618:  uint16(anon_sym_LBRACE),
	4619:  uint16(anon_sym_RBRACE),
	4620:  uint16(anon_sym_AMP),
	4621:  uint16(anon_sym_STAR),
	4622:  uint16(aux_sym_preproc_import_token1),
	4623:  uint16(aux_sym_preproc_ifdef_token1),
	4624:  uint16(aux_sym_preproc_ifdef_token2),
	4625:  uint16(aux_sym_preproc_ifdef_token3),
	4626:  uint16(aux_sym_preproc_else_token1),
	4627:  uint16(463),
	4628:  uint16(15),
	4629:  uint16(anon_sym_let),
	4630:  uint16(sym_identifier),
	4631:  uint16(anon_sym__),
	4632:  uint16(anon_sym_if),
	4633:  uint16(anon_sym_switch),
	4634:  uint16(anon_sym_fallthrough),
	4635:  uint16(anon_sym_loop),
	4636:  uint16(anon_sym_for),
	4637:  uint16(anon_sym_while),
	4638:  uint16(anon_sym_break),
	4639:  uint16(anon_sym_continue),
	4640:  uint16(anon_sym_continuing),
	4641:  uint16(anon_sym_return),
	4642:  uint16(anon_sym_discard),
	4643:  uint16(anon_sym_var),
	4644:  uint16(3),
	4645:  uint16(3),
	4646:  uint16(2),
	4647:  uint16(sym_block_comment),
	4648:  uint16(sym_line_comment),
	4649:  uint16(469),
	4650:  uint16(10),
	4651:  uint16(anon_sym_LPAREN),
	4652:  uint16(anon_sym_LBRACE),
	4653:  uint16(anon_sym_RBRACE),
	4654:  uint16(anon_sym_AMP),
	4655:  uint16(anon_sym_STAR),
	4656:  uint16(aux_sym_preproc_import_token1),
	4657:  uint16(aux_sym_preproc_ifdef_token1),
	4658:  uint16(aux_sym_preproc_ifdef_token2),
	4659:  uint16(aux_sym_preproc_ifdef_token3),
	4660:  uint16(aux_sym_preproc_else_token1),
	4661:  uint16(467),
	4662:  uint16(15),
	4663:  uint16(anon_sym_let),
	4664:  uint16(sym_identifier),
	4665:  uint16(anon_sym__),
	4666:  uint16(anon_sym_if),
	4667:  uint16(anon_sym_switch),
	4668:  uint16(anon_sym_fallthrough),
	4669:  uint16(anon_sym_loop),
	4670:  uint16(anon_sym_for),
	4671:  uint16(anon_sym_while),
	4672:  uint16(anon_sym_break),
	4673:  uint16(anon_sym_continue),
	4674:  uint16(anon_sym_continuing),
	4675:  uint16(anon_sym_return),
	4676:  uint16(anon_sym_discard),
	4677:  uint16(anon_sym_var),
	4678:  uint16(3),
	4679:  uint16(3),
	4680:  uint16(2),
	4681:  uint16(sym_block_comment),
	4682:  uint16(sym_line_comment),
	4683:  uint16(473),
	4684:  uint16(10),
	4685:  uint16(anon_sym_LPAREN),
	4686:  uint16(anon_sym_LBRACE),
	4687:  uint16(anon_sym_RBRACE),
	4688:  uint16(anon_sym_AMP),
	4689:  uint16(anon_sym_STAR),
	4690:  uint16(aux_sym_preproc_import_token1),
	4691:  uint16(aux_sym_preproc_ifdef_token1),
	4692:  uint16(aux_sym_preproc_ifdef_token2),
	4693:  uint16(aux_sym_preproc_ifdef_token3),
	4694:  uint16(aux_sym_preproc_else_token1),
	4695:  uint16(471),
	4696:  uint16(15),
	4697:  uint16(anon_sym_let),
	4698:  uint16(sym_identifier),
	4699:  uint16(anon_sym__),
	4700:  uint16(anon_sym_if),
	4701:  uint16(anon_sym_switch),
	4702:  uint16(anon_sym_fallthrough),
	4703:  uint16(anon_sym_loop),
	4704:  uint16(anon_sym_for),
	4705:  uint16(anon_sym_while),
	4706:  uint16(anon_sym_break),
	4707:  uint16(anon_sym_continue),
	4708:  uint16(anon_sym_continuing),
	4709:  uint16(anon_sym_return),
	4710:  uint16(anon_sym_discard),
	4711:  uint16(anon_sym_var),
	4712:  uint16(3),
	4713:  uint16(3),
	4714:  uint16(2),
	4715:  uint16(sym_block_comment),
	4716:  uint16(sym_line_comment),
	4717:  uint16(477),
	4718:  uint16(10),
	4719:  uint16(anon_sym_LPAREN),
	4720:  uint16(anon_sym_LBRACE),
	4721:  uint16(anon_sym_RBRACE),
	4722:  uint16(anon_sym_AMP),
	4723:  uint16(anon_sym_STAR),
	4724:  uint16(aux_sym_preproc_import_token1),
	4725:  uint16(aux_sym_preproc_ifdef_token1),
	4726:  uint16(aux_sym_preproc_ifdef_token2),
	4727:  uint16(aux_sym_preproc_ifdef_token3),
	4728:  uint16(aux_sym_preproc_else_token1),
	4729:  uint16(475),
	4730:  uint16(15),
	4731:  uint16(anon_sym_let),
	4732:  uint16(sym_identifier),
	4733:  uint16(anon_sym__),
	4734:  uint16(anon_sym_if),
	4735:  uint16(anon_sym_switch),
	4736:  uint16(anon_sym_fallthrough),
	4737:  uint16(anon_sym_loop),
	4738:  uint16(anon_sym_for),
	4739:  uint16(anon_sym_while),
	4740:  uint16(anon_sym_break),
	4741:  uint16(anon_sym_continue),
	4742:  uint16(anon_sym_continuing),
	4743:  uint16(anon_sym_return),
	4744:  uint16(anon_sym_discard),
	4745:  uint16(anon_sym_var),
	4746:  uint16(3),
	4747:  uint16(3),
	4748:  uint16(2),
	4749:  uint16(sym_block_comment),
	4750:  uint16(sym_line_comment),
	4751:  uint16(481),
	4752:  uint16(5),
	4753:  uint16(anon_sym_LT),
	4754:  uint16(anon_sym_GT),
	4755:  uint16(anon_sym_PIPE),
	4756:  uint16(anon_sym_AMP),
	4757:  uint16(anon_sym_SLASH),
	4758:  uint16(479),
	4759:  uint16(20),
	4760:  uint16(anon_sym_SEMI),
	4761:  uint16(anon_sym_COMMA),
	4762:  uint16(anon_sym_RPAREN),
	4763:  uint16(anon_sym_LBRACE),
	4764:  uint16(anon_sym_PIPE_PIPE),
	4765:  uint16(anon_sym_AMP_AMP),
	4766:  uint16(anon_sym_CARET),
	4767:  uint16(anon_sym_EQ_EQ),
	4768:  uint16(anon_sym_BANG_EQ),
	4769:  uint16(anon_sym_LT_EQ),
	4770:  uint16(anon_sym_GT_EQ),
	4771:  uint16(anon_sym_LT_LT),
	4772:  uint16(anon_sym_GT_GT),
	4773:  uint16(anon_sym_PLUS),
	4774:  uint16(anon_sym_DASH),
	4775:  uint16(anon_sym_STAR),
	4776:  uint16(anon_sym_PERCENT),
	4777:  uint16(anon_sym_LBRACK),
	4778:  uint16(anon_sym_RBRACK),
	4779:  uint16(anon_sym_DOT),
	4780:  uint16(9),
	4781:  uint16(443),
	4782:  uint16(1),
	4783:  uint16(anon_sym_LBRACK),
	4784:  uint16(445),
	4785:  uint16(1),
	4786:  uint16(anon_sym_DOT),
	4787:  uint16(489),
	4788:  uint16(1),
	4789:  uint16(anon_sym_SLASH),
	4790:  uint16(3),
	4791:  uint16(2),
	4792:  uint16(sym_block_comment),
	4793:  uint16(sym_line_comment),
	4794:  uint16(483),
	4795:  uint16(2),
	4796:  uint16(anon_sym_LT_LT),
	4797:  uint16(anon_sym_GT_GT),
	4798:  uint16(485),
	4799:  uint16(2),
	4800:  uint16(anon_sym_PLUS),
	4801:  uint16(anon_sym_DASH),
	4802:  uint16(487),
	4803:  uint16(2),
	4804:  uint16(anon_sym_STAR),
	4805:  uint16(anon_sym_PERCENT),
	4806:  uint16(441),
	4807:  uint16(4),
	4808:  uint16(anon_sym_LT),
	4809:  uint16(anon_sym_GT),
	4810:  uint16(anon_sym_PIPE),
	4811:  uint16(anon_sym_AMP),
	4812:  uint16(439),
	4813:  uint16(12),
	4814:  uint16(anon_sym_SEMI),
	4815:  uint16(anon_sym_COMMA),
	4816:  uint16(anon_sym_RPAREN),
	4817:  uint16(anon_sym_LBRACE),
	4818:  uint16(anon_sym_PIPE_PIPE),
	4819:  uint16(anon_sym_AMP_AMP),
	4820:  uint16(anon_sym_CARET),
	4821:  uint16(anon_sym_EQ_EQ),
	4822:  uint16(anon_sym_BANG_EQ),
	4823:  uint16(anon_sym_LT_EQ),
	4824:  uint16(anon_sym_GT_EQ),
	4825:  uint16(anon_sym_RBRACK),
	4826:  uint16(15),
	4827:  uint16(443),
	4828:  uint16(1),
	4829:  uint16(anon_sym_LBRACK),
	4830:  uint16(445),
	4831:  uint16(1),
	4832:  uint16(anon_sym_DOT),
	4833:  uint16(489),
	4834:  uint16(1),
	4835:  uint16(anon_sym_SLASH),
	4836:  uint16(493),
	4837:  uint16(1),
	4838:  uint16(anon_sym_AMP_AMP),
	4839:  uint16(495),
	4840:  uint16(1),
	4841:  uint16(anon_sym_PIPE),
	4842:  uint16(497),
	4843:  uint16(1),
	4844:  uint16(anon_sym_CARET),
	4845:  uint16(499),
	4846:  uint16(1),
	4847:  uint16(anon_sym_AMP),
	4848:  uint16(3),
	4849:  uint16(2),
	4850:  uint16(sym_block_comment),
	4851:  uint16(sym_line_comment),
	4852:  uint16(483),
	4853:  uint16(2),
	4854:  uint16(anon_sym_LT_LT),
	4855:  uint16(anon_sym_GT_GT),
	4856:  uint16(485),
	4857:  uint16(2),
	4858:  uint16(anon_sym_PLUS),
	4859:  uint16(anon_sym_DASH),
	4860:  uint16(487),
	4861:  uint16(2),
	4862:  uint16(anon_sym_STAR),
	4863:  uint16(anon_sym_PERCENT),
	4864:  uint16(491),
	4865:  uint16(2),
	4866:  uint16(anon_sym_LT),
	4867:  uint16(anon_sym_GT),
	4868:  uint16(501),
	4869:  uint16(2),
	4870:  uint16(anon_sym_EQ_EQ),
	4871:  uint16(anon_sym_BANG_EQ),
	4872:  uint16(503),
	4873:  uint16(2),
	4874:  uint16(anon_sym_LT_EQ),
	4875:  uint16(anon_sym_GT_EQ),
	4876:  uint16(439),
	4877:  uint16(6),
	4878:  uint16(anon_sym_SEMI),
	4879:  uint16(anon_sym_COMMA),
	4880:  uint16(anon_sym_RPAREN),
	4881:  uint16(anon_sym_LBRACE),
	4882:  uint16(anon_sym_PIPE_PIPE),
	4883:  uint16(anon_sym_RBRACK),
	4884:  uint16(14),
	4885:  uint16(443),
	4886:  uint16(1),
	4887:  uint16(anon_sym_LBRACK),
	4888:  uint16(445),
	4889:  uint16(1),
	4890:  uint16(anon_sym_DOT),
	4891:  uint16(489),
	4892:  uint16(1),
	4893:  uint16(anon_sym_SLASH),
	4894:  uint16(495),
	4895:  uint16(1),
	4896:  uint16(anon_sym_PIPE),
	4897:  uint16(497),
	4898:  uint16(1),
	4899:  uint16(anon_sym_CARET),
	4900:  uint16(499),
	4901:  uint16(1),
	4902:  uint16(anon_sym_AMP),
	4903:  uint16(3),
	4904:  uint16(2),
	4905:  uint16(sym_block_comment),
	4906:  uint16(sym_line_comment),
	4907:  uint16(483),
	4908:  uint16(2),
	4909:  uint16(anon_sym_LT_LT),
	4910:  uint16(anon_sym_GT_GT),
	4911:  uint16(485),
	4912:  uint16(2),
	4913:  uint16(anon_sym_PLUS),
	4914:  uint16(anon_sym_DASH),
	4915:  uint16(487),
	4916:  uint16(2),
	4917:  uint16(anon_sym_STAR),
	4918:  uint16(anon_sym_PERCENT),
	4919:  uint16(491),
	4920:  uint16(2),
	4921:  uint16(anon_sym_LT),
	4922:  uint16(anon_sym_GT),
	4923:  uint16(501),
	4924:  uint16(2),
	4925:  uint16(anon_sym_EQ_EQ),
	4926:  uint16(anon_sym_BANG_EQ),
	4927:  uint16(503),
	4928:  uint16(2),
	4929:  uint16(anon_sym_LT_EQ),
	4930:  uint16(anon_sym_GT_EQ),
	4931:  uint16(439),
	4932:  uint16(7),
	4933:  uint16(anon_sym_SEMI),
	4934:  uint16(anon_sym_COMMA),
	4935:  uint16(anon_sym_RPAREN),
	4936:  uint16(anon_sym_LBRACE),
	4937:  uint16(anon_sym_PIPE_PIPE),
	4938:  uint16(anon_sym_AMP_AMP),
	4939:  uint16(anon_sym_RBRACK),
	4940:  uint16(3),
	4941:  uint16(3),
	4942:  uint16(2),
	4943:  uint16(sym_block_comment),
	4944:  uint16(sym_line_comment),
	4945:  uint16(507),
	4946:  uint16(5),
	4947:  uint16(anon_sym_LT),
	4948:  uint16(anon_sym_GT),
	4949:  uint16(anon_sym_PIPE),
	4950:  uint16(anon_sym_AMP),
	4951:  uint16(anon_sym_SLASH),
	4952:  uint16(505),
	4953:  uint16(20),
	4954:  uint16(anon_sym_SEMI),
	4955:  uint16(anon_sym_COMMA),
	4956:  uint16(anon_sym_RPAREN),
	4957:  uint16(anon_sym_LBRACE),
	4958:  uint16(anon_sym_PIPE_PIPE),
	4959:  uint16(anon_sym_AMP_AMP),
	4960:  uint16(anon_sym_CARET),
	4961:  uint16(anon_sym_EQ_EQ),
	4962:  uint16(anon_sym_BANG_EQ),
	4963:  uint16(anon_sym_LT_EQ),
	4964:  uint16(anon_sym_GT_EQ),
	4965:  uint16(anon_sym_LT_LT),
	4966:  uint16(anon_sym_GT_GT),
	4967:  uint16(anon_sym_PLUS),
	4968:  uint16(anon_sym_DASH),
	4969:  uint16(anon_sym_STAR),
	4970:  uint16(anon_sym_PERCENT),
	4971:  uint16(anon_sym_LBRACK),
	4972:  uint16(anon_sym_RBRACK),
	4973:  uint16(anon_sym_DOT),
	4974:  uint16(3),
	4975:  uint16(3),
	4976:  uint16(2),
	4977:  uint16(sym_block_comment),
	4978:  uint16(sym_line_comment),
	4979:  uint16(511),
	4980:  uint16(5),
	4981:  uint16(anon_sym_LT),
	4982:  uint16(anon_sym_GT),
	4983:  uint16(anon_sym_PIPE),
	4984:  uint16(anon_sym_AMP),
	4985:  uint16(anon_sym_SLASH),
	4986:  uint16(509),
	4987:  uint16(20),
	4988:  uint16(anon_sym_SEMI),
	4989:  uint16(anon_sym_COMMA),
	4990:  uint16(anon_sym_RPAREN),
	4991:  uint16(anon_sym_LBRACE),
	4992:  uint16(anon_sym_PIPE_PIPE),
	4993:  uint16(anon_sym_AMP_AMP),
	4994:  uint16(anon_sym_CARET),
	4995:  uint16(anon_sym_EQ_EQ),
	4996:  uint16(anon_sym_BANG_EQ),
	4997:  uint16(anon_sym_LT_EQ),
	4998:  uint16(anon_sym_GT_EQ),
	4999:  uint16(anon_sym_LT_LT),
	5000:  uint16(anon_sym_GT_GT),
	5001:  uint16(anon_sym_PLUS),
	5002:  uint16(anon_sym_DASH),
	5003:  uint16(anon_sym_STAR),
	5004:  uint16(anon_sym_PERCENT),
	5005:  uint16(anon_sym_LBRACK),
	5006:  uint16(anon_sym_RBRACK),
	5007:  uint16(anon_sym_DOT),
	5008:  uint16(7),
	5009:  uint16(443),
	5010:  uint16(1),
	5011:  uint16(anon_sym_LBRACK),
	5012:  uint16(445),
	5013:  uint16(1),
	5014:  uint16(anon_sym_DOT),
	5015:  uint16(489),
	5016:  uint16(1),
	5017:  uint16(anon_sym_SLASH),
	5018:  uint16(3),
	5019:  uint16(2),
	5020:  uint16(sym_block_comment),
	5021:  uint16(sym_line_comment),
	5022:  uint16(487),
	5023:  uint16(2),
	5024:  uint16(anon_sym_STAR),
	5025:  uint16(anon_sym_PERCENT),
	5026:  uint16(441),
	5027:  uint16(4),
	5028:  uint16(anon_sym_LT),
	5029:  uint16(anon_sym_GT),
	5030:  uint16(anon_sym_PIPE),
	5031:  uint16(anon_sym_AMP),
	5032:  uint16(439),
	5033:  uint16(16),
	5034:  uint16(anon_sym_SEMI),
	5035:  uint16(anon_sym_COMMA),
	5036:  uint16(anon_sym_RPAREN),
	5037:  uint16(anon_sym_LBRACE),
	5038:  uint16(anon_sym_PIPE_PIPE),
	5039:  uint16(anon_sym_AMP_AMP),
	5040:  uint16(anon_sym_CARET),
	5041:  uint16(anon_sym_EQ_EQ),
	5042:  uint16(anon_sym_BANG_EQ),
	5043:  uint16(anon_sym_LT_EQ),
	5044:  uint16(anon_sym_GT_EQ),
	5045:  uint16(anon_sym_LT_LT),
	5046:  uint16(anon_sym_GT_GT),
	5047:  uint16(anon_sym_PLUS),
	5048:  uint16(anon_sym_DASH),
	5049:  uint16(anon_sym_RBRACK),
	5050:  uint16(14),
	5051:  uint16(441),
	5052:  uint16(1),
	5053:  uint16(anon_sym_PIPE),
	5054:  uint16(443),
	5055:  uint16(1),
	5056:  uint16(anon_sym_LBRACK),
	5057:  uint16(445),
	5058:  uint16(1),
	5059:  uint16(anon_sym_DOT),
	5060:  uint16(489),
	5061:  uint16(1),
	5062:  uint16(anon_sym_SLASH),
	5063:  uint16(497),
	5064:  uint16(1),
	5065:  uint16(anon_sym_CARET),
	5066:  uint16(499),
	5067:  uint16(1),
	5068:  uint16(anon_sym_AMP),
	5069:  uint16(3),
	5070:  uint16(2),
	5071:  uint16(sym_block_comment),
	5072:  uint16(sym_line_comment),
	5073:  uint16(483),
	5074:  uint16(2),
	5075:  uint16(anon_sym_LT_LT),
	5076:  uint16(anon_sym_GT_GT),
	5077:  uint16(485),
	5078:  uint16(2),
	5079:  uint16(anon_sym_PLUS),
	5080:  uint16(anon_sym_DASH),
	5081:  uint16(487),
	5082:  uint16(2),
	5083:  uint16(anon_sym_STAR),
	5084:  uint16(anon_sym_PERCENT),
	5085:  uint16(491),
	5086:  uint16(2),
	5087:  uint16(anon_sym_LT),
	5088:  uint16(anon_sym_GT),
	5089:  uint16(501),
	5090:  uint16(2),
	5091:  uint16(anon_sym_EQ_EQ),
	5092:  uint16(anon_sym_BANG_EQ),
	5093:  uint16(503),
	5094:  uint16(2),
	5095:  uint16(anon_sym_LT_EQ),
	5096:  uint16(anon_sym_GT_EQ),
	5097:  uint16(439),
	5098:  uint16(7),
	5099:  uint16(anon_sym_SEMI),
	5100:  uint16(anon_sym_COMMA),
	5101:  uint16(anon_sym_RPAREN),
	5102:  uint16(anon_sym_LBRACE),
	5103:  uint16(anon_sym_PIPE_PIPE),
	5104:  uint16(anon_sym_AMP_AMP),
	5105:  uint16(anon_sym_RBRACK),
	5106:  uint16(3),
	5107:  uint16(3),
	5108:  uint16(2),
	5109:  uint16(sym_block_comment),
	5110:  uint16(sym_line_comment),
	5111:  uint16(515),
	5112:  uint16(10),
	5113:  uint16(anon_sym_LPAREN),
	5114:  uint16(anon_sym_LBRACE),
	5115:  uint16(anon_sym_RBRACE),
	5116:  uint16(anon_sym_AMP),
	5117:  uint16(anon_sym_STAR),
	5118:  uint16(aux_sym_preproc_import_token1),
	5119:  uint16(aux_sym_preproc_ifdef_token1),
	5120:  uint16(aux_sym_preproc_ifdef_token2),
	5121:  uint16(aux_sym_preproc_ifdef_token3),
	5122:  uint16(aux_sym_preproc_else_token1),
	5123:  uint16(513),
	5124:  uint16(15),
	5125:  uint16(anon_sym_let),
	5126:  uint16(sym_identifier),
	5127:  uint16(anon_sym__),
	5128:  uint16(anon_sym_if),
	5129:  uint16(anon_sym_switch),
	5130:  uint16(anon_sym_fallthrough),
	5131:  uint16(anon_sym_loop),
	5132:  uint16(anon_sym_for),
	5133:  uint16(anon_sym_while),
	5134:  uint16(anon_sym_break),
	5135:  uint16(anon_sym_continue),
	5136:  uint16(anon_sym_continuing),
	5137:  uint16(anon_sym_return),
	5138:  uint16(anon_sym_discard),
	5139:  uint16(anon_sym_var),
	5140:  uint16(13),
	5141:  uint16(441),
	5142:  uint16(1),
	5143:  uint16(anon_sym_PIPE),
	5144:  uint16(443),
	5145:  uint16(1),
	5146:  uint16(anon_sym_LBRACK),
	5147:  uint16(445),
	5148:  uint16(1),
	5149:  uint16(anon_sym_DOT),
	5150:  uint16(489),
	5151:  uint16(1),
	5152:  uint16(anon_sym_SLASH),
	5153:  uint16(499),
	5154:  uint16(1),
	5155:  uint16(anon_sym_AMP),
	5156:  uint16(3),
	5157:  uint16(2),
	5158:  uint16(sym_block_comment),
	5159:  uint16(sym_line_comment),
	5160:  uint16(483),
	5161:  uint16(2),
	5162:  uint16(anon_sym_LT_LT),
	5163:  uint16(anon_sym_GT_GT),
	5164:  uint16(485),
	5165:  uint16(2),
	5166:  uint16(anon_sym_PLUS),
	5167:  uint16(anon_sym_DASH),
	5168:  uint16(487),
	5169:  uint16(2),
	5170:  uint16(anon_sym_STAR),
	5171:  uint16(anon_sym_PERCENT),
	5172:  uint16(491),
	5173:  uint16(2),
	5174:  uint16(anon_sym_LT),
	5175:  uint16(anon_sym_GT),
	5176:  uint16(501),
	5177:  uint16(2),
	5178:  uint16(anon_sym_EQ_EQ),
	5179:  uint16(anon_sym_BANG_EQ),
	5180:  uint16(503),
	5181:  uint16(2),
	5182:  uint16(anon_sym_LT_EQ),
	5183:  uint16(anon_sym_GT_EQ),
	5184:  uint16(439),
	5185:  uint16(8),
	5186:  uint16(anon_sym_SEMI),
	5187:  uint16(anon_sym_COMMA),
	5188:  uint16(anon_sym_RPAREN),
	5189:  uint16(anon_sym_LBRACE),
	5190:  uint16(anon_sym_PIPE_PIPE),
	5191:  uint16(anon_sym_AMP_AMP),
	5192:  uint16(anon_sym_CARET),
	5193:  uint16(anon_sym_RBRACK),
	5194:  uint16(3),
	5195:  uint16(3),
	5196:  uint16(2),
	5197:  uint16(sym_block_comment),
	5198:  uint16(sym_line_comment),
	5199:  uint16(519),
	5200:  uint16(10),
	5201:  uint16(anon_sym_LPAREN),
	5202:  uint16(anon_sym_LBRACE),
	5203:  uint16(anon_sym_RBRACE),
	5204:  uint16(anon_sym_AMP),
	5205:  uint16(anon_sym_STAR),
	5206:  uint16(aux_sym_preproc_import_token1),
	5207:  uint16(aux_sym_preproc_ifdef_token1),
	5208:  uint16(aux_sym_preproc_ifdef_token2),
	5209:  uint16(aux_sym_preproc_ifdef_token3),
	5210:  uint16(aux_sym_preproc_else_token1),
	5211:  uint16(517),
	5212:  uint16(15),
	5213:  uint16(anon_sym_let),
	5214:  uint16(sym_identifier),
	5215:  uint16(anon_sym__),
	5216:  uint16(anon_sym_if),
	5217:  uint16(anon_sym_switch),
	5218:  uint16(anon_sym_fallthrough),
	5219:  uint16(anon_sym_loop),
	5220:  uint16(anon_sym_for),
	5221:  uint16(anon_sym_while),
	5222:  uint16(anon_sym_break),
	5223:  uint16(anon_sym_continue),
	5224:  uint16(anon_sym_continuing),
	5225:  uint16(anon_sym_return),
	5226:  uint16(anon_sym_discard),
	5227:  uint16(anon_sym_var),
	5228:  uint16(3),
	5229:  uint16(3),
	5230:  uint16(2),
	5231:  uint16(sym_block_comment),
	5232:  uint16(sym_line_comment),
	5233:  uint16(523),
	5234:  uint16(10),
	5235:  uint16(anon_sym_LPAREN),
	5236:  uint16(anon_sym_LBRACE),
	5237:  uint16(anon_sym_RBRACE),
	5238:  uint16(anon_sym_AMP),
	5239:  uint16(anon_sym_STAR),
	5240:  uint16(aux_sym_preproc_import_token1),
	5241:  uint16(aux_sym_preproc_ifdef_token1),
	5242:  uint16(aux_sym_preproc_ifdef_token2),
	5243:  uint16(aux_sym_preproc_ifdef_token3),
	5244:  uint16(aux_sym_preproc_else_token1),
	5245:  uint16(521),
	5246:  uint16(15),
	5247:  uint16(anon_sym_let),
	5248:  uint16(sym_identifier),
	5249:  uint16(anon_sym__),
	5250:  uint16(anon_sym_if),
	5251:  uint16(anon_sym_switch),
	5252:  uint16(anon_sym_fallthrough),
	5253:  uint16(anon_sym_loop),
	5254:  uint16(anon_sym_for),
	5255:  uint16(anon_sym_while),
	5256:  uint16(anon_sym_break),
	5257:  uint16(anon_sym_continue),
	5258:  uint16(anon_sym_continuing),
	5259:  uint16(anon_sym_return),
	5260:  uint16(anon_sym_discard),
	5261:  uint16(anon_sym_var),
	5262:  uint16(3),
	5263:  uint16(3),
	5264:  uint16(2),
	5265:  uint16(sym_block_comment),
	5266:  uint16(sym_line_comment),
	5267:  uint16(527),
	5268:  uint16(5),
	5269:  uint16(anon_sym_LT),
	5270:  uint16(anon_sym_GT),
	5271:  uint16(anon_sym_PIPE),
	5272:  uint16(anon_sym_AMP),
	5273:  uint16(anon_sym_SLASH),
	5274:  uint16(525),
	5275:  uint16(20),
	5276:  uint16(anon_sym_SEMI),
	5277:  uint16(anon_sym_COMMA),
	5278:  uint16(anon_sym_RPAREN),
	5279:  uint16(anon_sym_LBRACE),
	5280:  uint16(anon_sym_PIPE_PIPE),
	5281:  uint16(anon_sym_AMP_AMP),
	5282:  uint16(anon_sym_CARET),
	5283:  uint16(anon_sym_EQ_EQ),
	5284:  uint16(anon_sym_BANG_EQ),
	5285:  uint16(anon_sym_LT_EQ),
	5286:  uint16(anon_sym_GT_EQ),
	5287:  uint16(anon_sym_LT_LT),
	5288:  uint16(anon_sym_GT_GT),
	5289:  uint16(anon_sym_PLUS),
	5290:  uint16(anon_sym_DASH),
	5291:  uint16(anon_sym_STAR),
	5292:  uint16(anon_sym_PERCENT),
	5293:  uint16(anon_sym_LBRACK),
	5294:  uint16(anon_sym_RBRACK),
	5295:  uint16(anon_sym_DOT),
	5296:  uint16(12),
	5297:  uint16(443),
	5298:  uint16(1),
	5299:  uint16(anon_sym_LBRACK),
	5300:  uint16(445),
	5301:  uint16(1),
	5302:  uint16(anon_sym_DOT),
	5303:  uint16(489),
	5304:  uint16(1),
	5305:  uint16(anon_sym_SLASH),
	5306:  uint16(3),
	5307:  uint16(2),
	5308:  uint16(sym_block_comment),
	5309:  uint16(sym_line_comment),
	5310:  uint16(441),
	5311:  uint16(2),
	5312:  uint16(anon_sym_PIPE),
	5313:  uint16(anon_sym_AMP),
	5314:  uint16(483),
	5315:  uint16(2),
	5316:  uint16(anon_sym_LT_LT),
	5317:  uint16(anon_sym_GT_GT),
	5318:  uint16(485),
	5319:  uint16(2),
	5320:  uint16(anon_sym_PLUS),
	5321:  uint16(anon_sym_DASH),
	5322:  uint16(487),
	5323:  uint16(2),
	5324:  uint16(anon_sym_STAR),
	5325:  uint16(anon_sym_PERCENT),
	5326:  uint16(491),
	5327:  uint16(2),
	5328:  uint16(anon_sym_LT),
	5329:  uint16(anon_sym_GT),
	5330:  uint16(501),
	5331:  uint16(2),
	5332:  uint16(anon_sym_EQ_EQ),
	5333:  uint16(anon_sym_BANG_EQ),
	5334:  uint16(503),
	5335:  uint16(2),
	5336:  uint16(anon_sym_LT_EQ),
	5337:  uint16(anon_sym_GT_EQ),
	5338:  uint16(439),
	5339:  uint16(8),
	5340:  uint16(anon_sym_SEMI),
	5341:  uint16(anon_sym_COMMA),
	5342:  uint16(anon_sym_RPAREN),
	5343:  uint16(anon_sym_LBRACE),
	5344:  uint16(anon_sym_PIPE_PIPE),
	5345:  uint16(anon_sym_AMP_AMP),
	5346:  uint16(anon_sym_CARET),
	5347:  uint16(anon_sym_RBRACK),
	5348:  uint16(11),
	5349:  uint16(443),
	5350:  uint16(1),
	5351:  uint16(anon_sym_LBRACK),
	5352:  uint16(445),
	5353:  uint16(1),
	5354:  uint16(anon_sym_DOT),
	5355:  uint16(489),
	5356:  uint16(1),
	5357:  uint16(anon_sym_SLASH),
	5358:  uint16(3),
	5359:  uint16(2),
	5360:  uint16(sym_block_comment),
	5361:  uint16(sym_line_comment),
	5362:  uint16(441),
	5363:  uint16(2),
	5364:  uint16(anon_sym_PIPE),
	5365:  uint16(anon_sym_AMP),
	5366:  uint16(483),
	5367:  uint16(2),
	5368:  uint16(anon_sym_LT_LT),
	5369:  uint16(anon_sym_GT_GT),
	5370:  uint16(485),
	5371:  uint16(2),
	5372:  uint16(anon_sym_PLUS),
	5373:  uint16(anon_sym_DASH),
	5374:  uint16(487),
	5375:  uint16(2),
	5376:  uint16(anon_sym_STAR),
	5377:  uint16(anon_sym_PERCENT),
	5378:  uint16(491),
	5379:  uint16(2),
	5380:  uint16(anon_sym_LT),
	5381:  uint16(anon_sym_GT),
	5382:  uint16(503),
	5383:  uint16(2),
	5384:  uint16(anon_sym_LT_EQ),
	5385:  uint16(anon_sym_GT_EQ),
	5386:  uint16(439),
	5387:  uint16(10),
	5388:  uint16(anon_sym_SEMI),
	5389:  uint16(anon_sym_COMMA),
	5390:  uint16(anon_sym_RPAREN),
	5391:  uint16(anon_sym_LBRACE),
	5392:  uint16(anon_sym_PIPE_PIPE),
	5393:  uint16(anon_sym_AMP_AMP),
	5394:  uint16(anon_sym_CARET),
	5395:  uint16(anon_sym_EQ_EQ),
	5396:  uint16(anon_sym_BANG_EQ),
	5397:  uint16(anon_sym_RBRACK),
	5398:  uint16(3),
	5399:  uint16(3),
	5400:  uint16(2),
	5401:  uint16(sym_block_comment),
	5402:  uint16(sym_line_comment),
	5403:  uint16(531),
	5404:  uint16(10),
	5405:  uint16(anon_sym_LPAREN),
	5406:  uint16(anon_sym_LBRACE),
	5407:  uint16(anon_sym_RBRACE),
	5408:  uint16(anon_sym_AMP),
	5409:  uint16(anon_sym_STAR),
	5410:  uint16(aux_sym_preproc_import_token1),
	5411:  uint16(aux_sym_preproc_ifdef_token1),
	5412:  uint16(aux_sym_preproc_ifdef_token2),
	5413:  uint16(aux_sym_preproc_ifdef_token3),
	5414:  uint16(aux_sym_preproc_else_token1),
	5415:  uint16(529),
	5416:  uint16(15),
	5417:  uint16(anon_sym_let),
	5418:  uint16(sym_identifier),
	5419:  uint16(anon_sym__),
	5420:  uint16(anon_sym_if),
	5421:  uint16(anon_sym_switch),
	5422:  uint16(anon_sym_fallthrough),
	5423:  uint16(anon_sym_loop),
	5424:  uint16(anon_sym_for),
	5425:  uint16(anon_sym_while),
	5426:  uint16(anon_sym_break),
	5427:  uint16(anon_sym_continue),
	5428:  uint16(anon_sym_continuing),
	5429:  uint16(anon_sym_return),
	5430:  uint16(anon_sym_discard),
	5431:  uint16(anon_sym_var),
	5432:  uint16(3),
	5433:  uint16(3),
	5434:  uint16(2),
	5435:  uint16(sym_block_comment),
	5436:  uint16(sym_line_comment),
	5437:  uint16(535),
	5438:  uint16(10),
	5439:  uint16(anon_sym_LPAREN),
	5440:  uint16(anon_sym_LBRACE),
	5441:  uint16(anon_sym_RBRACE),
	5442:  uint16(anon_sym_AMP),
	5443:  uint16(anon_sym_STAR),
	5444:  uint16(aux_sym_preproc_import_token1),
	5445:  uint16(aux_sym_preproc_ifdef_token1),
	5446:  uint16(aux_sym_preproc_ifdef_token2),
	5447:  uint16(aux_sym_preproc_ifdef_token3),
	5448:  uint16(aux_sym_preproc_else_token1),
	5449:  uint16(533),
	5450:  uint16(15),
	5451:  uint16(anon_sym_let),
	5452:  uint16(sym_identifier),
	5453:  uint16(anon_sym__),
	5454:  uint16(anon_sym_if),
	5455:  uint16(anon_sym_switch),
	5456:  uint16(anon_sym_fallthrough),
	5457:  uint16(anon_sym_loop),
	5458:  uint16(anon_sym_for),
	5459:  uint16(anon_sym_while),
	5460:  uint16(anon_sym_break),
	5461:  uint16(anon_sym_continue),
	5462:  uint16(anon_sym_continuing),
	5463:  uint16(anon_sym_return),
	5464:  uint16(anon_sym_discard),
	5465:  uint16(anon_sym_var),
	5466:  uint16(3),
	5467:  uint16(3),
	5468:  uint16(2),
	5469:  uint16(sym_block_comment),
	5470:  uint16(sym_line_comment),
	5471:  uint16(539),
	5472:  uint16(10),
	5473:  uint16(anon_sym_LPAREN),
	5474:  uint16(anon_sym_LBRACE),
	5475:  uint16(anon_sym_RBRACE),
	5476:  uint16(anon_sym_AMP),
	5477:  uint16(anon_sym_STAR),
	5478:  uint16(aux_sym_preproc_import_token1),
	5479:  uint16(aux_sym_preproc_ifdef_token1),
	5480:  uint16(aux_sym_preproc_ifdef_token2),
	5481:  uint16(aux_sym_preproc_ifdef_token3),
	5482:  uint16(aux_sym_preproc_else_token1),
	5483:  uint16(537),
	5484:  uint16(15),
	5485:  uint16(anon_sym_let),
	5486:  uint16(sym_identifier),
	5487:  uint16(anon_sym__),
	5488:  uint16(anon_sym_if),
	5489:  uint16(anon_sym_switch),
	5490:  uint16(anon_sym_fallthrough),
	5491:  uint16(anon_sym_loop),
	5492:  uint16(anon_sym_for),
	5493:  uint16(anon_sym_while),
	5494:  uint16(anon_sym_break),
	5495:  uint16(anon_sym_continue),
	5496:  uint16(anon_sym_continuing),
	5497:  uint16(anon_sym_return),
	5498:  uint16(anon_sym_discard),
	5499:  uint16(anon_sym_var),
	5500:  uint16(3),
	5501:  uint16(3),
	5502:  uint16(2),
	5503:  uint16(sym_block_comment),
	5504:  uint16(sym_line_comment),
	5505:  uint16(543),
	5506:  uint16(10),
	5507:  uint16(anon_sym_LPAREN),
	5508:  uint16(anon_sym_LBRACE),
	5509:  uint16(anon_sym_RBRACE),
	5510:  uint16(anon_sym_AMP),
	5511:  uint16(anon_sym_STAR),
	5512:  uint16(aux_sym_preproc_import_token1),
	5513:  uint16(aux_sym_preproc_ifdef_token1),
	5514:  uint16(aux_sym_preproc_ifdef_token2),
	5515:  uint16(aux_sym_preproc_ifdef_token3),
	5516:  uint16(aux_sym_preproc_else_token1),
	5517:  uint16(541),
	5518:  uint16(15),
	5519:  uint16(anon_sym_let),
	5520:  uint16(sym_identifier),
	5521:  uint16(anon_sym__),
	5522:  uint16(anon_sym_if),
	5523:  uint16(anon_sym_switch),
	5524:  uint16(anon_sym_fallthrough),
	5525:  uint16(anon_sym_loop),
	5526:  uint16(anon_sym_for),
	5527:  uint16(anon_sym_while),
	5528:  uint16(anon_sym_break),
	5529:  uint16(anon_sym_continue),
	5530:  uint16(anon_sym_continuing),
	5531:  uint16(anon_sym_return),
	5532:  uint16(anon_sym_discard),
	5533:  uint16(anon_sym_var),
	5534:  uint16(3),
	5535:  uint16(3),
	5536:  uint16(2),
	5537:  uint16(sym_block_comment),
	5538:  uint16(sym_line_comment),
	5539:  uint16(547),
	5540:  uint16(10),
	5541:  uint16(anon_sym_LPAREN),
	5542:  uint16(anon_sym_LBRACE),
	5543:  uint16(anon_sym_RBRACE),
	5544:  uint16(anon_sym_AMP),
	5545:  uint16(anon_sym_STAR),
	5546:  uint16(aux_sym_preproc_import_token1),
	5547:  uint16(aux_sym_preproc_ifdef_token1),
	5548:  uint16(aux_sym_preproc_ifdef_token2),
	5549:  uint16(aux_sym_preproc_ifdef_token3),
	5550:  uint16(aux_sym_preproc_else_token1),
	5551:  uint16(545),
	5552:  uint16(15),
	5553:  uint16(anon_sym_let),
	5554:  uint16(sym_identifier),
	5555:  uint16(anon_sym__),
	5556:  uint16(anon_sym_if),
	5557:  uint16(anon_sym_switch),
	5558:  uint16(anon_sym_fallthrough),
	5559:  uint16(anon_sym_loop),
	5560:  uint16(anon_sym_for),
	5561:  uint16(anon_sym_while),
	5562:  uint16(anon_sym_break),
	5563:  uint16(anon_sym_continue),
	5564:  uint16(anon_sym_continuing),
	5565:  uint16(anon_sym_return),
	5566:  uint16(anon_sym_discard),
	5567:  uint16(anon_sym_var),
	5568:  uint16(3),
	5569:  uint16(3),
	5570:  uint16(2),
	5571:  uint16(sym_block_comment),
	5572:  uint16(sym_line_comment),
	5573:  uint16(551),
	5574:  uint16(10),
	5575:  uint16(anon_sym_LPAREN),
	5576:  uint16(anon_sym_LBRACE),
	5577:  uint16(anon_sym_RBRACE),
	5578:  uint16(anon_sym_AMP),
	5579:  uint16(anon_sym_STAR),
	5580:  uint16(aux_sym_preproc_import_token1),
	5581:  uint16(aux_sym_preproc_ifdef_token1),
	5582:  uint16(aux_sym_preproc_ifdef_token2),
	5583:  uint16(aux_sym_preproc_ifdef_token3),
	5584:  uint16(aux_sym_preproc_else_token1),
	5585:  uint16(549),
	5586:  uint16(15),
	5587:  uint16(anon_sym_let),
	5588:  uint16(sym_identifier),
	5589:  uint16(anon_sym__),
	5590:  uint16(anon_sym_if),
	5591:  uint16(anon_sym_switch),
	5592:  uint16(anon_sym_fallthrough),
	5593:  uint16(anon_sym_loop),
	5594:  uint16(anon_sym_for),
	5595:  uint16(anon_sym_while),
	5596:  uint16(anon_sym_break),
	5597:  uint16(anon_sym_continue),
	5598:  uint16(anon_sym_continuing),
	5599:  uint16(anon_sym_return),
	5600:  uint16(anon_sym_discard),
	5601:  uint16(anon_sym_var),
	5602:  uint16(3),
	5603:  uint16(3),
	5604:  uint16(2),
	5605:  uint16(sym_block_comment),
	5606:  uint16(sym_line_comment),
	5607:  uint16(555),
	5608:  uint16(10),
	5609:  uint16(anon_sym_LPAREN),
	5610:  uint16(anon_sym_LBRACE),
	5611:  uint16(anon_sym_RBRACE),
	5612:  uint16(anon_sym_AMP),
	5613:  uint16(anon_sym_STAR),
	5614:  uint16(aux_sym_preproc_import_token1),
	5615:  uint16(aux_sym_preproc_ifdef_token1),
	5616:  uint16(aux_sym_preproc_ifdef_token2),
	5617:  uint16(aux_sym_preproc_ifdef_token3),
	5618:  uint16(aux_sym_preproc_else_token1),
	5619:  uint16(553),
	5620:  uint16(15),
	5621:  uint16(anon_sym_let),
	5622:  uint16(sym_identifier),
	5623:  uint16(anon_sym__),
	5624:  uint16(anon_sym_if),
	5625:  uint16(anon_sym_switch),
	5626:  uint16(anon_sym_fallthrough),
	5627:  uint16(anon_sym_loop),
	5628:  uint16(anon_sym_for),
	5629:  uint16(anon_sym_while),
	5630:  uint16(anon_sym_break),
	5631:  uint16(anon_sym_continue),
	5632:  uint16(anon_sym_continuing),
	5633:  uint16(anon_sym_return),
	5634:  uint16(anon_sym_discard),
	5635:  uint16(anon_sym_var),
	5636:  uint16(3),
	5637:  uint16(3),
	5638:  uint16(2),
	5639:  uint16(sym_block_comment),
	5640:  uint16(sym_line_comment),
	5641:  uint16(559),
	5642:  uint16(10),
	5643:  uint16(anon_sym_LPAREN),
	5644:  uint16(anon_sym_LBRACE),
	5645:  uint16(anon_sym_RBRACE),
	5646:  uint16(anon_sym_AMP),
	5647:  uint16(anon_sym_STAR),
	5648:  uint16(aux_sym_preproc_import_token1),
	5649:  uint16(aux_sym_preproc_ifdef_token1),
	5650:  uint16(aux_sym_preproc_ifdef_token2),
	5651:  uint16(aux_sym_preproc_ifdef_token3),
	5652:  uint16(aux_sym_preproc_else_token1),
	5653:  uint16(557),
	5654:  uint16(15),
	5655:  uint16(anon_sym_let),
	5656:  uint16(sym_identifier),
	5657:  uint16(anon_sym__),
	5658:  uint16(anon_sym_if),
	5659:  uint16(anon_sym_switch),
	5660:  uint16(anon_sym_fallthrough),
	5661:  uint16(anon_sym_loop),
	5662:  uint16(anon_sym_for),
	5663:  uint16(anon_sym_while),
	5664:  uint16(anon_sym_break),
	5665:  uint16(anon_sym_continue),
	5666:  uint16(anon_sym_continuing),
	5667:  uint16(anon_sym_return),
	5668:  uint16(anon_sym_discard),
	5669:  uint16(anon_sym_var),
	5670:  uint16(3),
	5671:  uint16(3),
	5672:  uint16(2),
	5673:  uint16(sym_block_comment),
	5674:  uint16(sym_line_comment),
	5675:  uint16(563),
	5676:  uint16(10),
	5677:  uint16(anon_sym_LPAREN),
	5678:  uint16(anon_sym_LBRACE),
	5679:  uint16(anon_sym_RBRACE),
	5680:  uint16(anon_sym_AMP),
	5681:  uint16(anon_sym_STAR),
	5682:  uint16(aux_sym_preproc_import_token1),
	5683:  uint16(aux_sym_preproc_ifdef_token1),
	5684:  uint16(aux_sym_preproc_ifdef_token2),
	5685:  uint16(aux_sym_preproc_ifdef_token3),
	5686:  uint16(aux_sym_preproc_else_token1),
	5687:  uint16(561),
	5688:  uint16(15),
	5689:  uint16(anon_sym_let),
	5690:  uint16(sym_identifier),
	5691:  uint16(anon_sym__),
	5692:  uint16(anon_sym_if),
	5693:  uint16(anon_sym_switch),
	5694:  uint16(anon_sym_fallthrough),
	5695:  uint16(anon_sym_loop),
	5696:  uint16(anon_sym_for),
	5697:  uint16(anon_sym_while),
	5698:  uint16(anon_sym_break),
	5699:  uint16(anon_sym_continue),
	5700:  uint16(anon_sym_continuing),
	5701:  uint16(anon_sym_return),
	5702:  uint16(anon_sym_discard),
	5703:  uint16(anon_sym_var),
	5704:  uint16(8),
	5705:  uint16(443),
	5706:  uint16(1),
	5707:  uint16(anon_sym_LBRACK),
	5708:  uint16(445),
	5709:  uint16(1),
	5710:  uint16(anon_sym_DOT),
	5711:  uint16(489),
	5712:  uint16(1),
	5713:  uint16(anon_sym_SLASH),
	5714:  uint16(3),
	5715:  uint16(2),
	5716:  uint16(sym_block_comment),
	5717:  uint16(sym_line_comment),
	5718:  uint16(485),
	5719:  uint16(2),
	5720:  uint16(anon_sym_PLUS),
	5721:  uint16(anon_sym_DASH),
	5722:  uint16(487),
	5723:  uint16(2),
	5724:  uint16(anon_sym_STAR),
	5725:  uint16(anon_sym_PERCENT),
	5726:  uint16(441),
	5727:  uint16(4),
	5728:  uint16(anon_sym_LT),
	5729:  uint16(anon_sym_GT),
	5730:  uint16(anon_sym_PIPE),
	5731:  uint16(anon_sym_AMP),
	5732:  uint16(439),
	5733:  uint16(14),
	5734:  uint16(anon_sym_SEMI),
	5735:  uint16(anon_sym_COMMA),
	5736:  uint16(anon_sym_RPAREN),
	5737:  uint16(anon_sym_LBRACE),
	5738:  uint16(anon_sym_PIPE_PIPE),
	5739:  uint16(anon_sym_AMP_AMP),
	5740:  uint16(anon_sym_CARET),
	5741:  uint16(anon_sym_EQ_EQ),
	5742:  uint16(anon_sym_BANG_EQ),
	5743:  uint16(anon_sym_LT_EQ),
	5744:  uint16(anon_sym_GT_EQ),
	5745:  uint16(anon_sym_LT_LT),
	5746:  uint16(anon_sym_GT_GT),
	5747:  uint16(anon_sym_RBRACK),
	5748:  uint16(3),
	5749:  uint16(3),
	5750:  uint16(2),
	5751:  uint16(sym_block_comment),
	5752:  uint16(sym_line_comment),
	5753:  uint16(567),
	5754:  uint16(5),
	5755:  uint16(anon_sym_LT),
	5756:  uint16(anon_sym_GT),
	5757:  uint16(anon_sym_PIPE),
	5758:  uint16(anon_sym_AMP),
	5759:  uint16(anon_sym_SLASH),
	5760:  uint16(565),
	5761:  uint16(20),
	5762:  uint16(anon_sym_SEMI),
	5763:  uint16(anon_sym_COMMA),
	5764:  uint16(anon_sym_RPAREN),
	5765:  uint16(anon_sym_LBRACE),
	5766:  uint16(anon_sym_PIPE_PIPE),
	5767:  uint16(anon_sym_AMP_AMP),
	5768:  uint16(anon_sym_CARET),
	5769:  uint16(anon_sym_EQ_EQ),
	5770:  uint16(anon_sym_BANG_EQ),
	5771:  uint16(anon_sym_LT_EQ),
	5772:  uint16(anon_sym_GT_EQ),
	5773:  uint16(anon_sym_LT_LT),
	5774:  uint16(anon_sym_GT_GT),
	5775:  uint16(anon_sym_PLUS),
	5776:  uint16(anon_sym_DASH),
	5777:  uint16(anon_sym_STAR),
	5778:  uint16(anon_sym_PERCENT),
	5779:  uint16(anon_sym_LBRACK),
	5780:  uint16(anon_sym_RBRACK),
	5781:  uint16(anon_sym_DOT),
	5782:  uint16(17),
	5783:  uint16(443),
	5784:  uint16(1),
	5785:  uint16(anon_sym_LBRACK),
	5786:  uint16(445),
	5787:  uint16(1),
	5788:  uint16(anon_sym_DOT),
	5789:  uint16(489),
	5790:  uint16(1),
	5791:  uint16(anon_sym_SLASH),
	5792:  uint16(493),
	5793:  uint16(1),
	5794:  uint16(anon_sym_AMP_AMP),
	5795:  uint16(495),
	5796:  uint16(1),
	5797:  uint16(anon_sym_PIPE),
	5798:  uint16(497),
	5799:  uint16(1),
	5800:  uint16(anon_sym_CARET),
	5801:  uint16(499),
	5802:  uint16(1),
	5803:  uint16(anon_sym_AMP),
	5804:  uint16(569),
	5805:  uint16(1),
	5806:  uint16(anon_sym_COMMA),
	5807:  uint16(571),
	5808:  uint16(1),
	5809:  uint16(anon_sym_RPAREN),
	5810:  uint16(573),
	5811:  uint16(1),
	5812:  uint16(anon_sym_PIPE_PIPE),
	5813:  uint16(3),
	5814:  uint16(2),
	5815:  uint16(sym_block_comment),
	5816:  uint16(sym_line_comment),
	5817:  uint16(483),
	5818:  uint16(2),
	5819:  uint16(anon_sym_LT_LT),
	5820:  uint16(anon_sym_GT_GT),
	5821:  uint16(485),
	5822:  uint16(2),
	5823:  uint16(anon_sym_PLUS),
	5824:  uint16(anon_sym_DASH),
	5825:  uint16(487),
	5826:  uint16(2),
	5827:  uint16(anon_sym_STAR),
	5828:  uint16(anon_sym_PERCENT),
	5829:  uint16(491),
	5830:  uint16(2),
	5831:  uint16(anon_sym_LT),
	5832:  uint16(anon_sym_GT),
	5833:  uint16(501),
	5834:  uint16(2),
	5835:  uint16(anon_sym_EQ_EQ),
	5836:  uint16(anon_sym_BANG_EQ),
	5837:  uint16(503),
	5838:  uint16(2),
	5839:  uint16(anon_sym_LT_EQ),
	5840:  uint16(anon_sym_GT_EQ),
	5841:  uint16(16),
	5842:  uint16(443),
	5843:  uint16(1),
	5844:  uint16(anon_sym_LBRACK),
	5845:  uint16(445),
	5846:  uint16(1),
	5847:  uint16(anon_sym_DOT),
	5848:  uint16(489),
	5849:  uint16(1),
	5850:  uint16(anon_sym_SLASH),
	5851:  uint16(493),
	5852:  uint16(1),
	5853:  uint16(anon_sym_AMP_AMP),
	5854:  uint16(495),
	5855:  uint16(1),
	5856:  uint16(anon_sym_PIPE),
	5857:  uint16(497),
	5858:  uint16(1),
	5859:  uint16(anon_sym_CARET),
	5860:  uint16(499),
	5861:  uint16(1),
	5862:  uint16(anon_sym_AMP),
	5863:  uint16(573),
	5864:  uint16(1),
	5865:  uint16(anon_sym_PIPE_PIPE),
	5866:  uint16(3),
	5867:  uint16(2),
	5868:  uint16(sym_block_comment),
	5869:  uint16(sym_line_comment),
	5870:  uint16(483),
	5871:  uint16(2),
	5872:  uint16(anon_sym_LT_LT),
	5873:  uint16(anon_sym_GT_GT),
	5874:  uint16(485),
	5875:  uint16(2),
	5876:  uint16(anon_sym_PLUS),
	5877:  uint16(anon_sym_DASH),
	5878:  uint16(487),
	5879:  uint16(2),
	5880:  uint16(anon_sym_STAR),
	5881:  uint16(anon_sym_PERCENT),
	5882:  uint16(491),
	5883:  uint16(2),
	5884:  uint16(anon_sym_LT),
	5885:  uint16(anon_sym_GT),
	5886:  uint16(501),
	5887:  uint16(2),
	5888:  uint16(anon_sym_EQ_EQ),
	5889:  uint16(anon_sym_BANG_EQ),
	5890:  uint16(503),
	5891:  uint16(2),
	5892:  uint16(anon_sym_LT_EQ),
	5893:  uint16(anon_sym_GT_EQ),
	5894:  uint16(575),
	5895:  uint16(2),
	5896:  uint16(anon_sym_SEMI),
	5897:  uint16(anon_sym_RPAREN),
	5898:  uint16(17),
	5899:  uint16(234),
	5900:  uint16(1),
	5901:  uint16(anon_sym_LBRACE),
	5902:  uint16(443),
	5903:  uint16(1),
	5904:  uint16(anon_sym_LBRACK),
	5905:  uint16(445),
	5906:  uint16(1),
	5907:  uint16(anon_sym_DOT),
	5908:  uint16(489),
	5909:  uint16(1),
	5910:  uint16(anon_sym_SLASH),
	5911:  uint16(493),
	5912:  uint16(1),
	5913:  uint16(anon_sym_AMP_AMP),
	5914:  uint16(495),
	5915:  uint16(1),
	5916:  uint16(anon_sym_PIPE),
	5917:  uint16(497),
	5918:  uint16(1),
	5919:  uint16(anon_sym_CARET),
	5920:  uint16(499),
	5921:  uint16(1),
	5922:  uint16(anon_sym_AMP),
	5923:  uint16(573),
	5924:  uint16(1),
	5925:  uint16(anon_sym_PIPE_PIPE),
	5926:  uint16(116),
	5927:  uint16(1),
	5928:  uint16(sym_compound_statement),
	5929:  uint16(3),
	5930:  uint16(2),
	5931:  uint16(sym_block_comment),
	5932:  uint16(sym_line_comment),
	5933:  uint16(483),
	5934:  uint16(2),
	5935:  uint16(anon_sym_LT_LT),
	5936:  uint16(anon_sym_GT_GT),
	5937:  uint16(485),
	5938:  uint16(2),
	5939:  uint16(anon_sym_PLUS),
	5940:  uint16(anon_sym_DASH),
	5941:  uint16(487),
	5942:  uint16(2),
	5943:  uint16(anon_sym_STAR),
	5944:  uint16(anon_sym_PERCENT),
	5945:  uint16(491),
	5946:  uint16(2),
	5947:  uint16(anon_sym_LT),
	5948:  uint16(anon_sym_GT),
	5949:  uint16(501),
	5950:  uint16(2),
	5951:  uint16(anon_sym_EQ_EQ),
	5952:  uint16(anon_sym_BANG_EQ),
	5953:  uint16(503),
	5954:  uint16(2),
	5955:  uint16(anon_sym_LT_EQ),
	5956:  uint16(anon_sym_GT_EQ),
	5957:  uint16(17),
	5958:  uint16(234),
	5959:  uint16(1),
	5960:  uint16(anon_sym_LBRACE),
	5961:  uint16(443),
	5962:  uint16(1),
	5963:  uint16(anon_sym_LBRACK),
	5964:  uint16(445),
	5965:  uint16(1),
	5966:  uint16(anon_sym_DOT),
	5967:  uint16(489),
	5968:  uint16(1),
	5969:  uint16(anon_sym_SLASH),
	5970:  uint16(493),
	5971:  uint16(1),
	5972:  uint16(anon_sym_AMP_AMP),
	5973:  uint16(495),
	5974:  uint16(1),
	5975:  uint16(anon_sym_PIPE),
	5976:  uint16(497),
	5977:  uint16(1),
	5978:  uint16(anon_sym_CARET),
	5979:  uint16(499),
	5980:  uint16(1),
	5981:  uint16(anon_sym_AMP),
	5982:  uint16(573),
	5983:  uint16(1),
	5984:  uint16(anon_sym_PIPE_PIPE),
	5985:  uint16(94),
	5986:  uint16(1),
	5987:  uint16(sym_compound_statement),
	5988:  uint16(3),
	5989:  uint16(2),
	5990:  uint16(sym_block_comment),
	5991:  uint16(sym_line_comment),
	5992:  uint16(483),
	5993:  uint16(2),
	5994:  uint16(anon_sym_LT_LT),
	5995:  uint16(anon_sym_GT_GT),
	5996:  uint16(485),
	5997:  uint16(2),
	5998:  uint16(anon_sym_PLUS),
	5999:  uint16(anon_sym_DASH),
	6000:  uint16(487),
	6001:  uint16(2),
	6002:  uint16(anon_sym_STAR),
	6003:  uint16(anon_sym_PERCENT),
	6004:  uint16(491),
	6005:  uint16(2),
	6006:  uint16(anon_sym_LT),
	6007:  uint16(anon_sym_GT),
	6008:  uint16(501),
	6009:  uint16(2),
	6010:  uint16(anon_sym_EQ_EQ),
	6011:  uint16(anon_sym_BANG_EQ),
	6012:  uint16(503),
	6013:  uint16(2),
	6014:  uint16(anon_sym_LT_EQ),
	6015:  uint16(anon_sym_GT_EQ),
	6016:  uint16(17),
	6017:  uint16(168),
	6018:  uint16(1),
	6019:  uint16(anon_sym_RPAREN),
	6020:  uint16(443),
	6021:  uint16(1),
	6022:  uint16(anon_sym_LBRACK),
	6023:  uint16(445),
	6024:  uint16(1),
	6025:  uint16(anon_sym_DOT),
	6026:  uint16(489),
	6027:  uint16(1),
	6028:  uint16(anon_sym_SLASH),
	6029:  uint16(493),
	6030:  uint16(1),
	6031:  uint16(anon_sym_AMP_AMP),
	6032:  uint16(495),
	6033:  uint16(1),
	6034:  uint16(anon_sym_PIPE),
	6035:  uint16(497),
	6036:  uint16(1),
	6037:  uint16(anon_sym_CARET),
	6038:  uint16(499),
	6039:  uint16(1),
	6040:  uint16(anon_sym_AMP),
	6041:  uint16(573),
	6042:  uint16(1),
	6043:  uint16(anon_sym_PIPE_PIPE),
	6044:  uint16(577),
	6045:  uint16(1),
	6046:  uint16(anon_sym_COMMA),
	6047:  uint16(3),
	6048:  uint16(2),
	6049:  uint16(sym_block_comment),
	6050:  uint16(sym_line_comment),
	6051:  uint16(483),
	6052:  uint16(2),
	6053:  uint16(anon_sym_LT_LT),
	6054:  uint16(anon_sym_GT_GT),
	6055:  uint16(485),
	6056:  uint16(2),
	6057:  uint16(anon_sym_PLUS),
	6058:  uint16(anon_sym_DASH),
	6059:  uint16(487),
	6060:  uint16(2),
	6061:  uint16(anon_sym_STAR),
	6062:  uint16(anon_sym_PERCENT),
	6063:  uint16(491),
	6064:  uint16(2),
	6065:  uint16(anon_sym_LT),
	6066:  uint16(anon_sym_GT),
	6067:  uint16(501),
	6068:  uint16(2),
	6069:  uint16(anon_sym_EQ_EQ),
	6070:  uint16(anon_sym_BANG_EQ),
	6071:  uint16(503),
	6072:  uint16(2),
	6073:  uint16(anon_sym_LT_EQ),
	6074:  uint16(anon_sym_GT_EQ),
	6075:  uint16(16),
	6076:  uint16(443),
	6077:  uint16(1),
	6078:  uint16(anon_sym_LBRACK),
	6079:  uint16(445),
	6080:  uint16(1),
	6081:  uint16(anon_sym_DOT),
	6082:  uint16(489),
	6083:  uint16(1),
	6084:  uint16(anon_sym_SLASH),
	6085:  uint16(493),
	6086:  uint16(1),
	6087:  uint16(anon_sym_AMP_AMP),
	6088:  uint16(495),
	6089:  uint16(1),
	6090:  uint16(anon_sym_PIPE),
	6091:  uint16(497),
	6092:  uint16(1),
	6093:  uint16(anon_sym_CARET),
	6094:  uint16(499),
	6095:  uint16(1),
	6096:  uint16(anon_sym_AMP),
	6097:  uint16(573),
	6098:  uint16(1),
	6099:  uint16(anon_sym_PIPE_PIPE),
	6100:  uint16(579),
	6101:  uint16(1),
	6102:  uint16(anon_sym_RBRACK),
	6103:  uint16(3),
	6104:  uint16(2),
	6105:  uint16(sym_block_comment),
	6106:  uint16(sym_line_comment),
	6107:  uint16(483),
	6108:  uint16(2),
	6109:  uint16(anon_sym_LT_LT),
	6110:  uint16(anon_sym_GT_GT),
	6111:  uint16(485),
	6112:  uint16(2),
	6113:  uint16(anon_sym_PLUS),
	6114:  uint16(anon_sym_DASH),
	6115:  uint16(487),
	6116:  uint16(2),
	6117:  uint16(anon_sym_STAR),
	6118:  uint16(anon_sym_PERCENT),
	6119:  uint16(491),
	6120:  uint16(2),
	6121:  uint16(anon_sym_LT),
	6122:  uint16(anon_sym_GT),
	6123:  uint16(501),
	6124:  uint16(2),
	6125:  uint16(anon_sym_EQ_EQ),
	6126:  uint16(anon_sym_BANG_EQ),
	6127:  uint16(503),
	6128:  uint16(2),
	6129:  uint16(anon_sym_LT_EQ),
	6130:  uint16(anon_sym_GT_EQ),
	6131:  uint16(16),
	6132:  uint16(443),
	6133:  uint16(1),
	6134:  uint16(anon_sym_LBRACK),
	6135:  uint16(445),
	6136:  uint16(1),
	6137:  uint16(anon_sym_DOT),
	6138:  uint16(489),
	6139:  uint16(1),
	6140:  uint16(anon_sym_SLASH),
	6141:  uint16(493),
	6142:  uint16(1),
	6143:  uint16(anon_sym_AMP_AMP),
	6144:  uint16(495),
	6145:  uint16(1),
	6146:  uint16(anon_sym_PIPE),
	6147:  uint16(497),
	6148:  uint16(1),
	6149:  uint16(anon_sym_CARET),
	6150:  uint16(499),
	6151:  uint16(1),
	6152:  uint16(anon_sym_AMP),
	6153:  uint16(573),
	6154:  uint16(1),
	6155:  uint16(anon_sym_PIPE_PIPE),
	6156:  uint16(581),
	6157:  uint16(1),
	6158:  uint16(anon_sym_SEMI),
	6159:  uint16(3),
	6160:  uint16(2),
	6161:  uint16(sym_block_comment),
	6162:  uint16(sym_line_comment),
	6163:  uint16(483),
	6164:  uint16(2),
	6165:  uint16(anon_sym_LT_LT),
	6166:  uint16(anon_sym_GT_GT),
	6167:  uint16(485),
	6168:  uint16(2),
	6169:  uint16(anon_sym_PLUS),
	6170:  uint16(anon_sym_DASH),
	6171:  uint16(487),
	6172:  uint16(2),
	6173:  uint16(anon_sym_STAR),
	6174:  uint16(anon_sym_PERCENT),
	6175:  uint16(491),
	6176:  uint16(2),
	6177:  uint16(anon_sym_LT),
	6178:  uint16(anon_sym_GT),
	6179:  uint16(501),
	6180:  uint16(2),
	6181:  uint16(anon_sym_EQ_EQ),
	6182:  uint16(anon_sym_BANG_EQ),
	6183:  uint16(503),
	6184:  uint16(2),
	6185:  uint16(anon_sym_LT_EQ),
	6186:  uint16(anon_sym_GT_EQ),
	6187:  uint16(16),
	6188:  uint16(443),
	6189:  uint16(1),
	6190:  uint16(anon_sym_LBRACK),
	6191:  uint16(445),
	6192:  uint16(1),
	6193:  uint16(anon_sym_DOT),
	6194:  uint16(489),
	6195:  uint16(1),
	6196:  uint16(anon_sym_SLASH),
	6197:  uint16(493),
	6198:  uint16(1),
	6199:  uint16(anon_sym_AMP_AMP),
	6200:  uint16(495),
	6201:  uint16(1),
	6202:  uint16(anon_sym_PIPE),
	6203:  uint16(497),
	6204:  uint16(1),
	6205:  uint16(anon_sym_CARET),
	6206:  uint16(499),
	6207:  uint16(1),
	6208:  uint16(anon_sym_AMP),
	6209:  uint16(573),
	6210:  uint16(1),
	6211:  uint16(anon_sym_PIPE_PIPE),
	6212:  uint16(583),
	6213:  uint16(1),
	6214:  uint16(anon_sym_SEMI),
	6215:  uint16(3),
	6216:  uint16(2),
	6217:  uint16(sym_block_comment),
	6218:  uint16(sym_line_comment),
	6219:  uint16(483),
	6220:  uint16(2),
	6221:  uint16(anon_sym_LT_LT),
	6222:  uint16(anon_sym_GT_GT),
	6223:  uint16(485),
	6224:  uint16(2),
	6225:  uint16(anon_sym_PLUS),
	6226:  uint16(anon_sym_DASH),
	6227:  uint16(487),
	6228:  uint16(2),
	6229:  uint16(anon_sym_STAR),
	6230:  uint16(anon_sym_PERCENT),
	6231:  uint16(491),
	6232:  uint16(2),
	6233:  uint16(anon_sym_LT),
	6234:  uint16(anon_sym_GT),
	6235:  uint16(501),
	6236:  uint16(2),
	6237:  uint16(anon_sym_EQ_EQ),
	6238:  uint16(anon_sym_BANG_EQ),
	6239:  uint16(503),
	6240:  uint16(2),
	6241:  uint16(anon_sym_LT_EQ),
	6242:  uint16(anon_sym_GT_EQ),
	6243:  uint16(16),
	6244:  uint16(443),
	6245:  uint16(1),
	6246:  uint16(anon_sym_LBRACK),
	6247:  uint16(445),
	6248:  uint16(1),
	6249:  uint16(anon_sym_DOT),
	6250:  uint16(489),
	6251:  uint16(1),
	6252:  uint16(anon_sym_SLASH),
	6253:  uint16(493),
	6254:  uint16(1),
	6255:  uint16(anon_sym_AMP_AMP),
	6256:  uint16(495),
	6257:  uint16(1),
	6258:  uint16(anon_sym_PIPE),
	6259:  uint16(497),
	6260:  uint16(1),
	6261:  uint16(anon_sym_CARET),
	6262:  uint16(499),
	6263:  uint16(1),
	6264:  uint16(anon_sym_AMP),
	6265:  uint16(573),
	6266:  uint16(1),
	6267:  uint16(anon_sym_PIPE_PIPE),
	6268:  uint16(585),
	6269:  uint16(1),
	6270:  uint16(anon_sym_RPAREN),
	6271:  uint16(3),
	6272:  uint16(2),
	6273:  uint16(sym_block_comment),
	6274:  uint16(sym_line_comment),
	6275:  uint16(483),
	6276:  uint16(2),
	6277:  uint16(anon_sym_LT_LT),
	6278:  uint16(anon_sym_GT_GT),
	6279:  uint16(485),
	6280:  uint16(2),
	6281:  uint16(anon_sym_PLUS),
	6282:  uint16(anon_sym_DASH),
	6283:  uint16(487),
	6284:  uint16(2),
	6285:  uint16(anon_sym_STAR),
	6286:  uint16(anon_sym_PERCENT),
	6287:  uint16(491),
	6288:  uint16(2),
	6289:  uint16(anon_sym_LT),
	6290:  uint16(anon_sym_GT),
	6291:  uint16(501),
	6292:  uint16(2),
	6293:  uint16(anon_sym_EQ_EQ),
	6294:  uint16(anon_sym_BANG_EQ),
	6295:  uint16(503),
	6296:  uint16(2),
	6297:  uint16(anon_sym_LT_EQ),
	6298:  uint16(anon_sym_GT_EQ),
	6299:  uint16(16),
	6300:  uint16(443),
	6301:  uint16(1),
	6302:  uint16(anon_sym_LBRACK),
	6303:  uint16(445),
	6304:  uint16(1),
	6305:  uint16(anon_sym_DOT),
	6306:  uint16(489),
	6307:  uint16(1),
	6308:  uint16(anon_sym_SLASH),
	6309:  uint16(493),
	6310:  uint16(1),
	6311:  uint16(anon_sym_AMP_AMP),
	6312:  uint16(495),
	6313:  uint16(1),
	6314:  uint16(anon_sym_PIPE),
	6315:  uint16(497),
	6316:  uint16(1),
	6317:  uint16(anon_sym_CARET),
	6318:  uint16(499),
	6319:  uint16(1),
	6320:  uint16(anon_sym_AMP),
	6321:  uint16(573),
	6322:  uint16(1),
	6323:  uint16(anon_sym_PIPE_PIPE),
	6324:  uint16(587),
	6325:  uint16(1),
	6326:  uint16(anon_sym_SEMI),
	6327:  uint16(3),
	6328:  uint16(2),
	6329:  uint16(sym_block_comment),
	6330:  uint16(sym_line_comment),
	6331:  uint16(483),
	6332:  uint16(2),
	6333:  uint16(anon_sym_LT_LT),
	6334:  uint16(anon_sym_GT_GT),
	6335:  uint16(485),
	6336:  uint16(2),
	6337:  uint16(anon_sym_PLUS),
	6338:  uint16(anon_sym_DASH),
	6339:  uint16(487),
	6340:  uint16(2),
	6341:  uint16(anon_sym_STAR),
	6342:  uint16(anon_sym_PERCENT),
	6343:  uint16(491),
	6344:  uint16(2),
	6345:  uint16(anon_sym_LT),
	6346:  uint16(anon_sym_GT),
	6347:  uint16(501),
	6348:  uint16(2),
	6349:  uint16(anon_sym_EQ_EQ),
	6350:  uint16(anon_sym_BANG_EQ),
	6351:  uint16(503),
	6352:  uint16(2),
	6353:  uint16(anon_sym_LT_EQ),
	6354:  uint16(anon_sym_GT_EQ),
	6355:  uint16(16),
	6356:  uint16(109),
	6357:  uint16(1),
	6358:  uint16(anon_sym_SEMI),
	6359:  uint16(443),
	6360:  uint16(1),
	6361:  uint16(anon_sym_LBRACK),
	6362:  uint16(445),
	6363:  uint16(1),
	6364:  uint16(anon_sym_DOT),
	6365:  uint16(489),
	6366:  uint16(1),
	6367:  uint16(anon_sym_SLASH),
	6368:  uint16(493),
	6369:  uint16(1),
	6370:  uint16(anon_sym_AMP_AMP),
	6371:  uint16(495),
	6372:  uint16(1),
	6373:  uint16(anon_sym_PIPE),
	6374:  uint16(497),
	6375:  uint16(1),
	6376:  uint16(anon_sym_CARET),
	6377:  uint16(499),
	6378:  uint16(1),
	6379:  uint16(anon_sym_AMP),
	6380:  uint16(573),
	6381:  uint16(1),
	6382:  uint16(anon_sym_PIPE_PIPE),
	6383:  uint16(3),
	6384:  uint16(2),
	6385:  uint16(sym_block_comment),
	6386:  uint16(sym_line_comment),
	6387:  uint16(483),
	6388:  uint16(2),
	6389:  uint16(anon_sym_LT_LT),
	6390:  uint16(anon_sym_GT_GT),
	6391:  uint16(485),
	6392:  uint16(2),
	6393:  uint16(anon_sym_PLUS),
	6394:  uint16(anon_sym_DASH),
	6395:  uint16(487),
	6396:  uint16(2),
	6397:  uint16(anon_sym_STAR),
	6398:  uint16(anon_sym_PERCENT),
	6399:  uint16(491),
	6400:  uint16(2),
	6401:  uint16(anon_sym_LT),
	6402:  uint16(anon_sym_GT),
	6403:  uint16(501),
	6404:  uint16(2),
	6405:  uint16(anon_sym_EQ_EQ),
	6406:  uint16(anon_sym_BANG_EQ),
	6407:  uint16(503),
	6408:  uint16(2),
	6409:  uint16(anon_sym_LT_EQ),
	6410:  uint16(anon_sym_GT_EQ),
	6411:  uint16(16),
	6412:  uint16(443),
	6413:  uint16(1),
	6414:  uint16(anon_sym_LBRACK),
	6415:  uint16(445),
	6416:  uint16(1),
	6417:  uint16(anon_sym_DOT),
	6418:  uint16(489),
	6419:  uint16(1),
	6420:  uint16(anon_sym_SLASH),
	6421:  uint16(493),
	6422:  uint16(1),
	6423:  uint16(anon_sym_AMP_AMP),
	6424:  uint16(495),
	6425:  uint16(1),
	6426:  uint16(anon_sym_PIPE),
	6427:  uint16(497),
	6428:  uint16(1),
	6429:  uint16(anon_sym_CARET),
	6430:  uint16(499),
	6431:  uint16(1),
	6432:  uint16(anon_sym_AMP),
	6433:  uint16(573),
	6434:  uint16(1),
	6435:  uint16(anon_sym_PIPE_PIPE),
	6436:  uint16(589),
	6437:  uint16(1),
	6438:  uint16(anon_sym_LBRACE),
	6439:  uint16(3),
	6440:  uint16(2),
	6441:  uint16(sym_block_comment),
	6442:  uint16(sym_line_comment),
	6443:  uint16(483),
	6444:  uint16(2),
	6445:  uint16(anon_sym_LT_LT),
	6446:  uint16(anon_sym_GT_GT),
	6447:  uint16(485),
	6448:  uint16(2),
	6449:  uint16(anon_sym_PLUS),
	6450:  uint16(anon_sym_DASH),
	6451:  uint16(487),
	6452:  uint16(2),
	6453:  uint16(anon_sym_STAR),
	6454:  uint16(anon_sym_PERCENT),
	6455:  uint16(491),
	6456:  uint16(2),
	6457:  uint16(anon_sym_LT),
	6458:  uint16(anon_sym_GT),
	6459:  uint16(501),
	6460:  uint16(2),
	6461:  uint16(anon_sym_EQ_EQ),
	6462:  uint16(anon_sym_BANG_EQ),
	6463:  uint16(503),
	6464:  uint16(2),
	6465:  uint16(anon_sym_LT_EQ),
	6466:  uint16(anon_sym_GT_EQ),
	6467:  uint16(16),
	6468:  uint16(443),
	6469:  uint16(1),
	6470:  uint16(anon_sym_LBRACK),
	6471:  uint16(445),
	6472:  uint16(1),
	6473:  uint16(anon_sym_DOT),
	6474:  uint16(489),
	6475:  uint16(1),
	6476:  uint16(anon_sym_SLASH),
	6477:  uint16(493),
	6478:  uint16(1),
	6479:  uint16(anon_sym_AMP_AMP),
	6480:  uint16(495),
	6481:  uint16(1),
	6482:  uint16(anon_sym_PIPE),
	6483:  uint16(497),
	6484:  uint16(1),
	6485:  uint16(anon_sym_CARET),
	6486:  uint16(499),
	6487:  uint16(1),
	6488:  uint16(anon_sym_AMP),
	6489:  uint16(573),
	6490:  uint16(1),
	6491:  uint16(anon_sym_PIPE_PIPE),
	6492:  uint16(591),
	6493:  uint16(1),
	6494:  uint16(anon_sym_RBRACK),
	6495:  uint16(3),
	6496:  uint16(2),
	6497:  uint16(sym_block_comment),
	6498:  uint16(sym_line_comment),
	6499:  uint16(483),
	6500:  uint16(2),
	6501:  uint16(anon_sym_LT_LT),
	6502:  uint16(anon_sym_GT_GT),
	6503:  uint16(485),
	6504:  uint16(2),
	6505:  uint16(anon_sym_PLUS),
	6506:  uint16(anon_sym_DASH),
	6507:  uint16(487),
	6508:  uint16(2),
	6509:  uint16(anon_sym_STAR),
	6510:  uint16(anon_sym_PERCENT),
	6511:  uint16(491),
	6512:  uint16(2),
	6513:  uint16(anon_sym_LT),
	6514:  uint16(anon_sym_GT),
	6515:  uint16(501),
	6516:  uint16(2),
	6517:  uint16(anon_sym_EQ_EQ),
	6518:  uint16(anon_sym_BANG_EQ),
	6519:  uint16(503),
	6520:  uint16(2),
	6521:  uint16(anon_sym_LT_EQ),
	6522:  uint16(anon_sym_GT_EQ),
	6523:  uint16(16),
	6524:  uint16(443),
	6525:  uint16(1),
	6526:  uint16(anon_sym_LBRACK),
	6527:  uint16(445),
	6528:  uint16(1),
	6529:  uint16(anon_sym_DOT),
	6530:  uint16(489),
	6531:  uint16(1),
	6532:  uint16(anon_sym_SLASH),
	6533:  uint16(493),
	6534:  uint16(1),
	6535:  uint16(anon_sym_AMP_AMP),
	6536:  uint16(495),
	6537:  uint16(1),
	6538:  uint16(anon_sym_PIPE),
	6539:  uint16(497),
	6540:  uint16(1),
	6541:  uint16(anon_sym_CARET),
	6542:  uint16(499),
	6543:  uint16(1),
	6544:  uint16(anon_sym_AMP),
	6545:  uint16(573),
	6546:  uint16(1),
	6547:  uint16(anon_sym_PIPE_PIPE),
	6548:  uint16(593),
	6549:  uint16(1),
	6550:  uint16(anon_sym_COMMA),
	6551:  uint16(3),
	6552:  uint16(2),
	6553:  uint16(sym_block_comment),
	6554:  uint16(sym_line_comment),
	6555:  uint16(483),
	6556:  uint16(2),
	6557:  uint16(anon_sym_LT_LT),
	6558:  uint16(anon_sym_GT_GT),
	6559:  uint16(485),
	6560:  uint16(2),
	6561:  uint16(anon_sym_PLUS),
	6562:  uint16(anon_sym_DASH),
	6563:  uint16(487),
	6564:  uint16(2),
	6565:  uint16(anon_sym_STAR),
	6566:  uint16(anon_sym_PERCENT),
	6567:  uint16(491),
	6568:  uint16(2),
	6569:  uint16(anon_sym_LT),
	6570:  uint16(anon_sym_GT),
	6571:  uint16(501),
	6572:  uint16(2),
	6573:  uint16(anon_sym_EQ_EQ),
	6574:  uint16(anon_sym_BANG_EQ),
	6575:  uint16(503),
	6576:  uint16(2),
	6577:  uint16(anon_sym_LT_EQ),
	6578:  uint16(anon_sym_GT_EQ),
	6579:  uint16(16),
	6580:  uint16(443),
	6581:  uint16(1),
	6582:  uint16(anon_sym_LBRACK),
	6583:  uint16(445),
	6584:  uint16(1),
	6585:  uint16(anon_sym_DOT),
	6586:  uint16(489),
	6587:  uint16(1),
	6588:  uint16(anon_sym_SLASH),
	6589:  uint16(493),
	6590:  uint16(1),
	6591:  uint16(anon_sym_AMP_AMP),
	6592:  uint16(495),
	6593:  uint16(1),
	6594:  uint16(anon_sym_PIPE),
	6595:  uint16(497),
	6596:  uint16(1),
	6597:  uint16(anon_sym_CARET),
	6598:  uint16(499),
	6599:  uint16(1),
	6600:  uint16(anon_sym_AMP),
	6601:  uint16(573),
	6602:  uint16(1),
	6603:  uint16(anon_sym_PIPE_PIPE),
	6604:  uint16(595),
	6605:  uint16(1),
	6606:  uint16(anon_sym_SEMI),
	6607:  uint16(3),
	6608:  uint16(2),
	6609:  uint16(sym_block_comment),
	6610:  uint16(sym_line_comment),
	6611:  uint16(483),
	6612:  uint16(2),
	6613:  uint16(anon_sym_LT_LT),
	6614:  uint16(anon_sym_GT_GT),
	6615:  uint16(485),
	6616:  uint16(2),
	6617:  uint16(anon_sym_PLUS),
	6618:  uint16(anon_sym_DASH),
	6619:  uint16(487),
	6620:  uint16(2),
	6621:  uint16(anon_sym_STAR),
	6622:  uint16(anon_sym_PERCENT),
	6623:  uint16(491),
	6624:  uint16(2),
	6625:  uint16(anon_sym_LT),
	6626:  uint16(anon_sym_GT),
	6627:  uint16(501),
	6628:  uint16(2),
	6629:  uint16(anon_sym_EQ_EQ),
	6630:  uint16(anon_sym_BANG_EQ),
	6631:  uint16(503),
	6632:  uint16(2),
	6633:  uint16(anon_sym_LT_EQ),
	6634:  uint16(anon_sym_GT_EQ),
	6635:  uint16(16),
	6636:  uint16(443),
	6637:  uint16(1),
	6638:  uint16(anon_sym_LBRACK),
	6639:  uint16(445),
	6640:  uint16(1),
	6641:  uint16(anon_sym_DOT),
	6642:  uint16(489),
	6643:  uint16(1),
	6644:  uint16(anon_sym_SLASH),
	6645:  uint16(493),
	6646:  uint16(1),
	6647:  uint16(anon_sym_AMP_AMP),
	6648:  uint16(495),
	6649:  uint16(1),
	6650:  uint16(anon_sym_PIPE),
	6651:  uint16(497),
	6652:  uint16(1),
	6653:  uint16(anon_sym_CARET),
	6654:  uint16(499),
	6655:  uint16(1),
	6656:  uint16(anon_sym_AMP),
	6657:  uint16(573),
	6658:  uint16(1),
	6659:  uint16(anon_sym_PIPE_PIPE),
	6660:  uint16(597),
	6661:  uint16(1),
	6662:  uint16(anon_sym_SEMI),
	6663:  uint16(3),
	6664:  uint16(2),
	6665:  uint16(sym_block_comment),
	6666:  uint16(sym_line_comment),
	6667:  uint16(483),
	6668:  uint16(2),
	6669:  uint16(anon_sym_LT_LT),
	6670:  uint16(anon_sym_GT_GT),
	6671:  uint16(485),
	6672:  uint16(2),
	6673:  uint16(anon_sym_PLUS),
	6674:  uint16(anon_sym_DASH),
	6675:  uint16(487),
	6676:  uint16(2),
	6677:  uint16(anon_sym_STAR),
	6678:  uint16(anon_sym_PERCENT),
	6679:  uint16(491),
	6680:  uint16(2),
	6681:  uint16(anon_sym_LT),
	6682:  uint16(anon_sym_GT),
	6683:  uint16(501),
	6684:  uint16(2),
	6685:  uint16(anon_sym_EQ_EQ),
	6686:  uint16(anon_sym_BANG_EQ),
	6687:  uint16(503),
	6688:  uint16(2),
	6689:  uint16(anon_sym_LT_EQ),
	6690:  uint16(anon_sym_GT_EQ),
	6691:  uint16(16),
	6692:  uint16(443),
	6693:  uint16(1),
	6694:  uint16(anon_sym_LBRACK),
	6695:  uint16(445),
	6696:  uint16(1),
	6697:  uint16(anon_sym_DOT),
	6698:  uint16(489),
	6699:  uint16(1),
	6700:  uint16(anon_sym_SLASH),
	6701:  uint16(493),
	6702:  uint16(1),
	6703:  uint16(anon_sym_AMP_AMP),
	6704:  uint16(495),
	6705:  uint16(1),
	6706:  uint16(anon_sym_PIPE),
	6707:  uint16(497),
	6708:  uint16(1),
	6709:  uint16(anon_sym_CARET),
	6710:  uint16(499),
	6711:  uint16(1),
	6712:  uint16(anon_sym_AMP),
	6713:  uint16(573),
	6714:  uint16(1),
	6715:  uint16(anon_sym_PIPE_PIPE),
	6716:  uint16(599),
	6717:  uint16(1),
	6718:  uint16(anon_sym_SEMI),
	6719:  uint16(3),
	6720:  uint16(2),
	6721:  uint16(sym_block_comment),
	6722:  uint16(sym_line_comment),
	6723:  uint16(483),
	6724:  uint16(2),
	6725:  uint16(anon_sym_LT_LT),
	6726:  uint16(anon_sym_GT_GT),
	6727:  uint16(485),
	6728:  uint16(2),
	6729:  uint16(anon_sym_PLUS),
	6730:  uint16(anon_sym_DASH),
	6731:  uint16(487),
	6732:  uint16(2),
	6733:  uint16(anon_sym_STAR),
	6734:  uint16(anon_sym_PERCENT),
	6735:  uint16(491),
	6736:  uint16(2),
	6737:  uint16(anon_sym_LT),
	6738:  uint16(anon_sym_GT),
	6739:  uint16(501),
	6740:  uint16(2),
	6741:  uint16(anon_sym_EQ_EQ),
	6742:  uint16(anon_sym_BANG_EQ),
	6743:  uint16(503),
	6744:  uint16(2),
	6745:  uint16(anon_sym_LT_EQ),
	6746:  uint16(anon_sym_GT_EQ),
	6747:  uint16(16),
	6748:  uint16(443),
	6749:  uint16(1),
	6750:  uint16(anon_sym_LBRACK),
	6751:  uint16(445),
	6752:  uint16(1),
	6753:  uint16(anon_sym_DOT),
	6754:  uint16(489),
	6755:  uint16(1),
	6756:  uint16(anon_sym_SLASH),
	6757:  uint16(493),
	6758:  uint16(1),
	6759:  uint16(anon_sym_AMP_AMP),
	6760:  uint16(495),
	6761:  uint16(1),
	6762:  uint16(anon_sym_PIPE),
	6763:  uint16(497),
	6764:  uint16(1),
	6765:  uint16(anon_sym_CARET),
	6766:  uint16(499),
	6767:  uint16(1),
	6768:  uint16(anon_sym_AMP),
	6769:  uint16(573),
	6770:  uint16(1),
	6771:  uint16(anon_sym_PIPE_PIPE),
	6772:  uint16(601),
	6773:  uint16(1),
	6774:  uint16(anon_sym_SEMI),
	6775:  uint16(3),
	6776:  uint16(2),
	6777:  uint16(sym_block_comment),
	6778:  uint16(sym_line_comment),
	6779:  uint16(483),
	6780:  uint16(2),
	6781:  uint16(anon_sym_LT_LT),
	6782:  uint16(anon_sym_GT_GT),
	6783:  uint16(485),
	6784:  uint16(2),
	6785:  uint16(anon_sym_PLUS),
	6786:  uint16(anon_sym_DASH),
	6787:  uint16(487),
	6788:  uint16(2),
	6789:  uint16(anon_sym_STAR),
	6790:  uint16(anon_sym_PERCENT),
	6791:  uint16(491),
	6792:  uint16(2),
	6793:  uint16(anon_sym_LT),
	6794:  uint16(anon_sym_GT),
	6795:  uint16(501),
	6796:  uint16(2),
	6797:  uint16(anon_sym_EQ_EQ),
	6798:  uint16(anon_sym_BANG_EQ),
	6799:  uint16(503),
	6800:  uint16(2),
	6801:  uint16(anon_sym_LT_EQ),
	6802:  uint16(anon_sym_GT_EQ),
	6803:  uint16(16),
	6804:  uint16(443),
	6805:  uint16(1),
	6806:  uint16(anon_sym_LBRACK),
	6807:  uint16(445),
	6808:  uint16(1),
	6809:  uint16(anon_sym_DOT),
	6810:  uint16(489),
	6811:  uint16(1),
	6812:  uint16(anon_sym_SLASH),
	6813:  uint16(493),
	6814:  uint16(1),
	6815:  uint16(anon_sym_AMP_AMP),
	6816:  uint16(495),
	6817:  uint16(1),
	6818:  uint16(anon_sym_PIPE),
	6819:  uint16(497),
	6820:  uint16(1),
	6821:  uint16(anon_sym_CARET),
	6822:  uint16(499),
	6823:  uint16(1),
	6824:  uint16(anon_sym_AMP),
	6825:  uint16(573),
	6826:  uint16(1),
	6827:  uint16(anon_sym_PIPE_PIPE),
	6828:  uint16(603),
	6829:  uint16(1),
	6830:  uint16(anon_sym_RBRACK),
	6831:  uint16(3),
	6832:  uint16(2),
	6833:  uint16(sym_block_comment),
	6834:  uint16(sym_line_comment),
	6835:  uint16(483),
	6836:  uint16(2),
	6837:  uint16(anon_sym_LT_LT),
	6838:  uint16(anon_sym_GT_GT),
	6839:  uint16(485),
	6840:  uint16(2),
	6841:  uint16(anon_sym_PLUS),
	6842:  uint16(anon_sym_DASH),
	6843:  uint16(487),
	6844:  uint16(2),
	6845:  uint16(anon_sym_STAR),
	6846:  uint16(anon_sym_PERCENT),
	6847:  uint16(491),
	6848:  uint16(2),
	6849:  uint16(anon_sym_LT),
	6850:  uint16(anon_sym_GT),
	6851:  uint16(501),
	6852:  uint16(2),
	6853:  uint16(anon_sym_EQ_EQ),
	6854:  uint16(anon_sym_BANG_EQ),
	6855:  uint16(503),
	6856:  uint16(2),
	6857:  uint16(anon_sym_LT_EQ),
	6858:  uint16(anon_sym_GT_EQ),
	6859:  uint16(4),
	6860:  uint16(607),
	6861:  uint16(1),
	6862:  uint16(anon_sym_enable),
	6863:  uint16(3),
	6864:  uint16(2),
	6865:  uint16(sym_block_comment),
	6866:  uint16(sym_line_comment),
	6867:  uint16(151),
	6868:  uint16(2),
	6869:  uint16(sym_enable_directive),
	6870:  uint16(aux_sym_source_file_repeat1),
	6871:  uint16(605),
	6872:  uint16(14),
	6874:  uint16(anon_sym_SEMI),
	6875:  uint16(anon_sym_let),
	6876:  uint16(anon_sym_override),
	6877:  uint16(anon_sym_type),
	6878:  uint16(anon_sym_virtual),
	6879:  uint16(anon_sym_fn),
	6880:  uint16(anon_sym_struct),
	6881:  uint16(anon_sym_AT),
	6882:  uint16(anon_sym_var),
	6883:  uint16(aux_sym_preproc_import_token1),
	6884:  uint16(aux_sym_define_import_path_token1),
	6885:  uint16(aux_sym_preproc_ifdef_token1),
	6886:  uint16(aux_sym_preproc_ifdef_token2),
	6887:  uint16(3),
	6888:  uint16(360),
	6889:  uint16(1),
	6890:  uint16(sym_texel_format),
	6891:  uint16(3),
	6892:  uint16(2),
	6893:  uint16(sym_block_comment),
	6894:  uint16(sym_line_comment),
	6895:  uint16(610),
	6896:  uint16(16),
	6897:  uint16(anon_sym_rgba8unorm),
	6898:  uint16(anon_sym_rgba8snorm),
	6899:  uint16(anon_sym_rgba8uint),
	6900:  uint16(anon_sym_rgba8sint),
	6901:  uint16(anon_sym_rgba16uint),
	6902:  uint16(anon_sym_rgba16sint),
	6903:  uint16(anon_sym_rgba16float),
	6904:  uint16(anon_sym_r32uint),
	6905:  uint16(anon_sym_r32sint),
	6906:  uint16(anon_sym_r32float),
	6907:  uint16(anon_sym_rg32uint),
	6908:  uint16(anon_sym_rg32sint),
	6909:  uint16(anon_sym_rg32float),
	6910:  uint16(anon_sym_rgba32uint),
	6911:  uint16(anon_sym_rgba32sint),
	6912:  uint16(anon_sym_rgba32float),
	6913:  uint16(2),
	6914:  uint16(3),
	6915:  uint16(2),
	6916:  uint16(sym_block_comment),
	6917:  uint16(sym_line_comment),
	6918:  uint16(612),
	6919:  uint16(16),
	6921:  uint16(anon_sym_SEMI),
	6922:  uint16(anon_sym_let),
	6923:  uint16(anon_sym_override),
	6924:  uint16(anon_sym_type),
	6925:  uint16(anon_sym_virtual),
	6926:  uint16(anon_sym_fn),
	6927:  uint16(anon_sym_struct),
	6928:  uint16(anon_sym_AT),
	6929:  uint16(anon_sym_var),
	6930:  uint16(aux_sym_preproc_import_token1),
	6931:  uint16(aux_sym_define_import_path_token1),
	6932:  uint16(aux_sym_preproc_ifdef_token1),
	6933:  uint16(aux_sym_preproc_ifdef_token2),
	6934:  uint16(aux_sym_preproc_ifdef_token3),
	6935:  uint16(aux_sym_preproc_else_token1),
	6936:  uint16(2),
	6937:  uint16(3),
	6938:  uint16(2),
	6939:  uint16(sym_block_comment),
	6940:  uint16(sym_line_comment),
	6941:  uint16(614),
	6942:  uint16(16),
	6944:  uint16(anon_sym_SEMI),
	6945:  uint16(anon_sym_let),
	6946:  uint16(anon_sym_override),
	6947:  uint16(anon_sym_type),
	6948:  uint16(anon_sym_virtual),
	6949:  uint16(anon_sym_fn),
	6950:  uint16(anon_sym_struct),
	6951:  uint16(anon_sym_AT),
	6952:  uint16(anon_sym_var),
	6953:  uint16(aux_sym_preproc_import_token1),
	6954:  uint16(aux_sym_define_import_path_token1),
	6955:  uint16(aux_sym_preproc_ifdef_token1),
	6956:  uint16(aux_sym_preproc_ifdef_token2),
	6957:  uint16(aux_sym_preproc_ifdef_token3),
	6958:  uint16(aux_sym_preproc_else_token1),
	6959:  uint16(2),
	6960:  uint16(3),
	6961:  uint16(2),
	6962:  uint16(sym_block_comment),
	6963:  uint16(sym_line_comment),
	6964:  uint16(616),
	6965:  uint16(16),
	6967:  uint16(anon_sym_SEMI),
	6968:  uint16(anon_sym_let),
	6969:  uint16(anon_sym_override),
	6970:  uint16(anon_sym_type),
	6971:  uint16(anon_sym_virtual),
	6972:  uint16(anon_sym_fn),
	6973:  uint16(anon_sym_struct),
	6974:  uint16(anon_sym_AT),
	6975:  uint16(anon_sym_var),
	6976:  uint16(aux_sym_preproc_import_token1),
	6977:  uint16(aux_sym_define_import_path_token1),
	6978:  uint16(aux_sym_preproc_ifdef_token1),
	6979:  uint16(aux_sym_preproc_ifdef_token2),
	6980:  uint16(aux_sym_preproc_ifdef_token3),
	6981:  uint16(aux_sym_preproc_else_token1),
	6982:  uint16(2),
	6983:  uint16(3),
	6984:  uint16(2),
	6985:  uint16(sym_block_comment),
	6986:  uint16(sym_line_comment),
	6987:  uint16(618),
	6988:  uint16(16),
	6990:  uint16(anon_sym_SEMI),
	6991:  uint16(anon_sym_let),
	6992:  uint16(anon_sym_override),
	6993:  uint16(anon_sym_type),
	6994:  uint16(anon_sym_virtual),
	6995:  uint16(anon_sym_fn),
	6996:  uint16(anon_sym_struct),
	6997:  uint16(anon_sym_AT),
	6998:  uint16(anon_sym_var),
	6999:  uint16(aux_sym_preproc_import_token1),
	7000:  uint16(aux_sym_define_import_path_token1),
	7001:  uint16(aux_sym_preproc_ifdef_token1),
	7002:  uint16(aux_sym_preproc_ifdef_token2),
	7003:  uint16(aux_sym_preproc_ifdef_token3),
	7004:  uint16(aux_sym_preproc_else_token1),
	7005:  uint16(2),
	7006:  uint16(3),
	7007:  uint16(2),
	7008:  uint16(sym_block_comment),
	7009:  uint16(sym_line_comment),
	7010:  uint16(620),
	7011:  uint16(16),
	7013:  uint16(anon_sym_SEMI),
	7014:  uint16(anon_sym_let),
	7015:  uint16(anon_sym_override),
	7016:  uint16(anon_sym_type),
	7017:  uint16(anon_sym_virtual),
	7018:  uint16(anon_sym_fn),
	7019:  uint16(anon_sym_struct),
	7020:  uint16(anon_sym_AT),
	7021:  uint16(anon_sym_var),
	7022:  uint16(aux_sym_preproc_import_token1),
	7023:  uint16(aux_sym_define_import_path_token1),
	7024:  uint16(aux_sym_preproc_ifdef_token1),
	7025:  uint16(aux_sym_preproc_ifdef_token2),
	7026:  uint16(aux_sym_preproc_ifdef_token3),
	7027:  uint16(aux_sym_preproc_else_token1),
	7028:  uint16(2),
	7029:  uint16(3),
	7030:  uint16(2),
	7031:  uint16(sym_block_comment),
	7032:  uint16(sym_line_comment),
	7033:  uint16(622),
	7034:  uint16(16),
	7036:  uint16(anon_sym_SEMI),
	7037:  uint16(anon_sym_let),
	7038:  uint16(anon_sym_override),
	7039:  uint16(anon_sym_type),
	7040:  uint16(anon_sym_virtual),
	7041:  uint16(anon_sym_fn),
	7042:  uint16(anon_sym_struct),
	7043:  uint16(anon_sym_AT),
	7044:  uint16(anon_sym_var),
	7045:  uint16(aux_sym_preproc_import_token1),
	7046:  uint16(aux_sym_define_import_path_token1),
	7047:  uint16(aux_sym_preproc_ifdef_token1),
	7048:  uint16(aux_sym_preproc_ifdef_token2),
	7049:  uint16(aux_sym_preproc_ifdef_token3),
	7050:  uint16(aux_sym_preproc_else_token1),
	7051:  uint16(2),
	7052:  uint16(3),
	7053:  uint16(2),
	7054:  uint16(sym_block_comment),
	7055:  uint16(sym_line_comment),
	7056:  uint16(624),
	7057:  uint16(16),
	7059:  uint16(anon_sym_SEMI),
	7060:  uint16(anon_sym_let),
	7061:  uint16(anon_sym_override),
	7062:  uint16(anon_sym_type),
	7063:  uint16(anon_sym_virtual),
	7064:  uint16(anon_sym_fn),
	7065:  uint16(anon_sym_struct),
	7066:  uint16(anon_sym_AT),
	7067:  uint16(anon_sym_var),
	7068:  uint16(aux_sym_preproc_import_token1),
	7069:  uint16(aux_sym_define_import_path_token1),
	7070:  uint16(aux_sym_preproc_ifdef_token1),
	7071:  uint16(aux_sym_preproc_ifdef_token2),
	7072:  uint16(aux_sym_preproc_ifdef_token3),
	7073:  uint16(aux_sym_preproc_else_token1),
	7074:  uint16(2),
	7075:  uint16(3),
	7076:  uint16(2),
	7077:  uint16(sym_block_comment),
	7078:  uint16(sym_line_comment),
	7079:  uint16(626),
	7080:  uint16(16),
	7082:  uint16(anon_sym_SEMI),
	7083:  uint16(anon_sym_let),
	7084:  uint16(anon_sym_override),
	7085:  uint16(anon_sym_type),
	7086:  uint16(anon_sym_virtual),
	7087:  uint16(anon_sym_fn),
	7088:  uint16(anon_sym_struct),
	7089:  uint16(anon_sym_AT),
	7090:  uint16(anon_sym_var),
	7091:  uint16(aux_sym_preproc_import_token1),
	7092:  uint16(aux_sym_define_import_path_token1),
	7093:  uint16(aux_sym_preproc_ifdef_token1),
	7094:  uint16(aux_sym_preproc_ifdef_token2),
	7095:  uint16(aux_sym_preproc_ifdef_token3),
	7096:  uint16(aux_sym_preproc_else_token1),
	7097:  uint16(2),
	7098:  uint16(3),
	7099:  uint16(2),
	7100:  uint16(sym_block_comment),
	7101:  uint16(sym_line_comment),
	7102:  uint16(628),
	7103:  uint16(16),
	7105:  uint16(anon_sym_SEMI),
	7106:  uint16(anon_sym_let),
	7107:  uint16(anon_sym_override),
	7108:  uint16(anon_sym_type),
	7109:  uint16(anon_sym_virtual),
	7110:  uint16(anon_sym_fn),
	7111:  uint16(anon_sym_struct),
	7112:  uint16(anon_sym_AT),
	7113:  uint16(anon_sym_var),
	7114:  uint16(aux_sym_preproc_import_token1),
	7115:  uint16(aux_sym_define_import_path_token1),
	7116:  uint16(aux_sym_preproc_ifdef_token1),
	7117:  uint16(aux_sym_preproc_ifdef_token2),
	7118:  uint16(aux_sym_preproc_ifdef_token3),
	7119:  uint16(aux_sym_preproc_else_token1),
	7120:  uint16(2),
	7121:  uint16(3),
	7122:  uint16(2),
	7123:  uint16(sym_block_comment),
	7124:  uint16(sym_line_comment),
	7125:  uint16(630),
	7126:  uint16(16),
	7128:  uint16(anon_sym_SEMI),
	7129:  uint16(anon_sym_let),
	7130:  uint16(anon_sym_override),
	7131:  uint16(anon_sym_type),
	7132:  uint16(anon_sym_virtual),
	7133:  uint16(anon_sym_fn),
	7134:  uint16(anon_sym_struct),
	7135:  uint16(anon_sym_AT),
	7136:  uint16(anon_sym_var),
	7137:  uint16(aux_sym_preproc_import_token1),
	7138:  uint16(aux_sym_define_import_path_token1),
	7139:  uint16(aux_sym_preproc_ifdef_token1),
	7140:  uint16(aux_sym_preproc_ifdef_token2),
	7141:  uint16(aux_sym_preproc_ifdef_token3),
	7142:  uint16(aux_sym_preproc_else_token1),
	7143:  uint16(2),
	7144:  uint16(3),
	7145:  uint16(2),
	7146:  uint16(sym_block_comment),
	7147:  uint16(sym_line_comment),
	7148:  uint16(632),
	7149:  uint16(16),
	7151:  uint16(anon_sym_SEMI),
	7152:  uint16(anon_sym_let),
	7153:  uint16(anon_sym_override),
	7154:  uint16(anon_sym_type),
	7155:  uint16(anon_sym_virtual),
	7156:  uint16(anon_sym_fn),
	7157:  uint16(anon_sym_struct),
	7158:  uint16(anon_sym_AT),
	7159:  uint16(anon_sym_var),
	7160:  uint16(aux_sym_preproc_import_token1),
	7161:  uint16(aux_sym_define_import_path_token1),
	7162:  uint16(aux_sym_preproc_ifdef_token1),
	7163:  uint16(aux_sym_preproc_ifdef_token2),
	7164:  uint16(aux_sym_preproc_ifdef_token3),
	7165:  uint16(aux_sym_preproc_else_token1),
	7166:  uint16(2),
	7167:  uint16(3),
	7168:  uint16(2),
	7169:  uint16(sym_block_comment),
	7170:  uint16(sym_line_comment),
	7171:  uint16(634),
	7172:  uint16(16),
	7174:  uint16(anon_sym_SEMI),
	7175:  uint16(anon_sym_let),
	7176:  uint16(anon_sym_override),
	7177:  uint16(anon_sym_type),
	7178:  uint16(anon_sym_virtual),
	7179:  uint16(anon_sym_fn),
	7180:  uint16(anon_sym_struct),
	7181:  uint16(anon_sym_AT),
	7182:  uint16(anon_sym_var),
	7183:  uint16(aux_sym_preproc_import_token1),
	7184:  uint16(aux_sym_define_import_path_token1),
	7185:  uint16(aux_sym_preproc_ifdef_token1),
	7186:  uint16(aux_sym_preproc_ifdef_token2),
	7187:  uint16(aux_sym_preproc_ifdef_token3),
	7188:  uint16(aux_sym_preproc_else_token1),
	7189:  uint16(2),
	7190:  uint16(3),
	7191:  uint16(2),
	7192:  uint16(sym_block_comment),
	7193:  uint16(sym_line_comment),
	7194:  uint16(636),
	7195:  uint16(16),
	7197:  uint16(anon_sym_SEMI),
	7198:  uint16(anon_sym_let),
	7199:  uint16(anon_sym_override),
	7200:  uint16(anon_sym_type),
	7201:  uint16(anon_sym_virtual),
	7202:  uint16(anon_sym_fn),
	7203:  uint16(anon_sym_struct),
	7204:  uint16(anon_sym_AT),
	7205:  uint16(anon_sym_var),
	7206:  uint16(aux_sym_preproc_import_token1),
	7207:  uint16(aux_sym_define_import_path_token1),
	7208:  uint16(aux_sym_preproc_ifdef_token1),
	7209:  uint16(aux_sym_preproc_ifdef_token2),
	7210:  uint16(aux_sym_preproc_ifdef_token3),
	7211:  uint16(aux_sym_preproc_else_token1),
	7212:  uint16(2),
	7213:  uint16(3),
	7214:  uint16(2),
	7215:  uint16(sym_block_comment),
	7216:  uint16(sym_line_comment),
	7217:  uint16(638),
	7218:  uint16(16),
	7220:  uint16(anon_sym_SEMI),
	7221:  uint16(anon_sym_let),
	7222:  uint16(anon_sym_override),
	7223:  uint16(anon_sym_type),
	7224:  uint16(anon_sym_virtual),
	7225:  uint16(anon_sym_fn),
	7226:  uint16(anon_sym_struct),
	7227:  uint16(anon_sym_AT),
	7228:  uint16(anon_sym_var),
	7229:  uint16(aux_sym_preproc_import_token1),
	7230:  uint16(aux_sym_define_import_path_token1),
	7231:  uint16(aux_sym_preproc_ifdef_token1),
	7232:  uint16(aux_sym_preproc_ifdef_token2),
	7233:  uint16(aux_sym_preproc_ifdef_token3),
	7234:  uint16(aux_sym_preproc_else_token1),
	7235:  uint16(2),
	7236:  uint16(3),
	7237:  uint16(2),
	7238:  uint16(sym_block_comment),
	7239:  uint16(sym_line_comment),
	7240:  uint16(640),
	7241:  uint16(16),
	7243:  uint16(anon_sym_SEMI),
	7244:  uint16(anon_sym_let),
	7245:  uint16(anon_sym_override),
	7246:  uint16(anon_sym_type),
	7247:  uint16(anon_sym_virtual),
	7248:  uint16(anon_sym_fn),
	7249:  uint16(anon_sym_struct),
	7250:  uint16(anon_sym_AT),
	7251:  uint16(anon_sym_var),
	7252:  uint16(aux_sym_preproc_import_token1),
	7253:  uint16(aux_sym_define_import_path_token1),
	7254:  uint16(aux_sym_preproc_ifdef_token1),
	7255:  uint16(aux_sym_preproc_ifdef_token2),
	7256:  uint16(aux_sym_preproc_ifdef_token3),
	7257:  uint16(aux_sym_preproc_else_token1),
	7258:  uint16(2),
	7259:  uint16(3),
	7260:  uint16(2),
	7261:  uint16(sym_block_comment),
	7262:  uint16(sym_line_comment),
	7263:  uint16(642),
	7264:  uint16(16),
	7266:  uint16(anon_sym_SEMI),
	7267:  uint16(anon_sym_let),
	7268:  uint16(anon_sym_override),
	7269:  uint16(anon_sym_type),
	7270:  uint16(anon_sym_virtual),
	7271:  uint16(anon_sym_fn),
	7272:  uint16(anon_sym_struct),
	7273:  uint16(anon_sym_AT),
	7274:  uint16(anon_sym_var),
	7275:  uint16(aux_sym_preproc_import_token1),
	7276:  uint16(aux_sym_define_import_path_token1),
	7277:  uint16(aux_sym_preproc_ifdef_token1),
	7278:  uint16(aux_sym_preproc_ifdef_token2),
	7279:  uint16(aux_sym_preproc_ifdef_token3),
	7280:  uint16(aux_sym_preproc_else_token1),
	7281:  uint16(2),
	7282:  uint16(3),
	7283:  uint16(2),
	7284:  uint16(sym_block_comment),
	7285:  uint16(sym_line_comment),
	7286:  uint16(644),
	7287:  uint16(16),
	7289:  uint16(anon_sym_SEMI),
	7290:  uint16(anon_sym_let),
	7291:  uint16(anon_sym_override),
	7292:  uint16(anon_sym_type),
	7293:  uint16(anon_sym_virtual),
	7294:  uint16(anon_sym_fn),
	7295:  uint16(anon_sym_struct),
	7296:  uint16(anon_sym_AT),
	7297:  uint16(anon_sym_var),
	7298:  uint16(aux_sym_preproc_import_token1),
	7299:  uint16(aux_sym_define_import_path_token1),
	7300:  uint16(aux_sym_preproc_ifdef_token1),
	7301:  uint16(aux_sym_preproc_ifdef_token2),
	7302:  uint16(aux_sym_preproc_ifdef_token3),
	7303:  uint16(aux_sym_preproc_else_token1),
	7304:  uint16(2),
	7305:  uint16(3),
	7306:  uint16(2),
	7307:  uint16(sym_block_comment),
	7308:  uint16(sym_line_comment),
	7309:  uint16(646),
	7310:  uint16(16),
	7312:  uint16(anon_sym_SEMI),
	7313:  uint16(anon_sym_let),
	7314:  uint16(anon_sym_override),
	7315:  uint16(anon_sym_type),
	7316:  uint16(anon_sym_virtual),
	7317:  uint16(anon_sym_fn),
	7318:  uint16(anon_sym_struct),
	7319:  uint16(anon_sym_AT),
	7320:  uint16(anon_sym_var),
	7321:  uint16(aux_sym_preproc_import_token1),
	7322:  uint16(aux_sym_define_import_path_token1),
	7323:  uint16(aux_sym_preproc_ifdef_token1),
	7324:  uint16(aux_sym_preproc_ifdef_token2),
	7325:  uint16(aux_sym_preproc_ifdef_token3),
	7326:  uint16(aux_sym_preproc_else_token1),
	7327:  uint16(2),
	7328:  uint16(3),
	7329:  uint16(2),
	7330:  uint16(sym_block_comment),
	7331:  uint16(sym_line_comment),
	7332:  uint16(648),
	7333:  uint16(16),
	7335:  uint16(anon_sym_SEMI),
	7336:  uint16(anon_sym_let),
	7337:  uint16(anon_sym_override),
	7338:  uint16(anon_sym_type),
	7339:  uint16(anon_sym_virtual),
	7340:  uint16(anon_sym_fn),
	7341:  uint16(anon_sym_struct),
	7342:  uint16(anon_sym_AT),
	7343:  uint16(anon_sym_var),
	7344:  uint16(aux_sym_preproc_import_token1),
	7345:  uint16(aux_sym_define_import_path_token1),
	7346:  uint16(aux_sym_preproc_ifdef_token1),
	7347:  uint16(aux_sym_preproc_ifdef_token2),
	7348:  uint16(aux_sym_preproc_ifdef_token3),
	7349:  uint16(aux_sym_preproc_else_token1),
	7350:  uint16(2),
	7351:  uint16(3),
	7352:  uint16(2),
	7353:  uint16(sym_block_comment),
	7354:  uint16(sym_line_comment),
	7355:  uint16(650),
	7356:  uint16(15),
	7358:  uint16(anon_sym_SEMI),
	7359:  uint16(anon_sym_let),
	7360:  uint16(anon_sym_override),
	7361:  uint16(anon_sym_type),
	7362:  uint16(anon_sym_virtual),
	7363:  uint16(anon_sym_fn),
	7364:  uint16(anon_sym_struct),
	7365:  uint16(anon_sym_enable),
	7366:  uint16(anon_sym_AT),
	7367:  uint16(anon_sym_var),
	7368:  uint16(aux_sym_preproc_import_token1),
	7369:  uint16(aux_sym_define_import_path_token1),
	7370:  uint16(aux_sym_preproc_ifdef_token1),
	7371:  uint16(aux_sym_preproc_ifdef_token2),
	7372:  uint16(2),
	7373:  uint16(3),
	7374:  uint16(2),
	7375:  uint16(sym_block_comment),
	7376:  uint16(sym_line_comment),
	7377:  uint16(652),
	7378:  uint16(15),
	7379:  uint16(anon_sym_SEMI),
	7380:  uint16(anon_sym_EQ),
	7381:  uint16(anon_sym_LPAREN),
	7382:  uint16(anon_sym_COMMA),
	7383:  uint16(anon_sym_RPAREN),
	7384:  uint16(anon_sym_LBRACE),
	7385:  uint16(anon_sym_RBRACE),
	7386:  uint16(anon_sym_AT),
	7387:  uint16(sym_identifier),
	7388:  uint16(anon_sym_GT),
	7389:  uint16(aux_sym_preproc_import_token1),
	7390:  uint16(aux_sym_preproc_ifdef_token1),
	7391:  uint16(aux_sym_preproc_ifdef_token2),
	7392:  uint16(aux_sym_preproc_ifdef_token3),
	7393:  uint16(aux_sym_preproc_else_token1),
	7394:  uint16(2),
	7395:  uint16(3),
	7396:  uint16(2),
	7397:  uint16(sym_block_comment),
	7398:  uint16(sym_line_comment),
	7399:  uint16(654),
	7400:  uint16(15),
	7401:  uint16(anon_sym_SEMI),
	7402:  uint16(anon_sym_EQ),
	7403:  uint16(anon_sym_LPAREN),
	7404:  uint16(anon_sym_COMMA),
	7405:  uint16(anon_sym_RPAREN),
	7406:  uint16(anon_sym_LBRACE),
	7407:  uint16(anon_sym_RBRACE),
	7408:  uint16(anon_sym_AT),
	7409:  uint16(sym_identifier),
	7410:  uint16(anon_sym_GT),
	7411:  uint16(aux_sym_preproc_import_token1),
	7412:  uint16(aux_sym_preproc_ifdef_token1),
	7413:  uint16(aux_sym_preproc_ifdef_token2),
	7414:  uint16(aux_sym_preproc_ifdef_token3),
	7415:  uint16(aux_sym_preproc_else_token1),
	7416:  uint16(5),
	7417:  uint16(658),
	7418:  uint16(1),
	7419:  uint16(anon_sym_LBRACK),
	7420:  uint16(660),
	7421:  uint16(1),
	7422:  uint16(anon_sym_DOT),
	7423:  uint16(202),
	7424:  uint16(1),
	7425:  uint16(sym_postfix_expression),
	7426:  uint16(3),
	7427:  uint16(2),
	7428:  uint16(sym_block_comment),
	7429:  uint16(sym_line_comment),
	7430:  uint16(656),
	7431:  uint16(12),
	7432:  uint16(anon_sym_EQ),
	7433:  uint16(anon_sym_RPAREN),
	7434:  uint16(anon_sym_PLUS_EQ),
	7435:  uint16(anon_sym_DASH_EQ),
	7436:  uint16(anon_sym_STAR_EQ),
	7437:  uint16(anon_sym_SLASH_EQ),
	7438:  uint16(anon_sym_PERCENT_EQ),
	7439:  uint16(anon_sym_AMP_EQ),
	7440:  uint16(anon_sym_PIPE_EQ),
	7441:  uint16(anon_sym_CARET_EQ),
	7442:  uint16(anon_sym_PLUS_PLUS),
	7443:  uint16(anon_sym_DASH_DASH),
	7444:  uint16(2),
	7445:  uint16(3),
	7446:  uint16(2),
	7447:  uint16(sym_block_comment),
	7448:  uint16(sym_line_comment),
	7449:  uint16(431),
	7450:  uint16(15),
	7451:  uint16(anon_sym_SEMI),
	7452:  uint16(anon_sym_EQ),
	7453:  uint16(anon_sym_LPAREN),
	7454:  uint16(anon_sym_COMMA),
	7455:  uint16(anon_sym_RPAREN),
	7456:  uint16(anon_sym_LBRACE),
	7457:  uint16(anon_sym_RBRACE),
	7458:  uint16(anon_sym_AT),
	7459:  uint16(sym_identifier),
	7460:  uint16(anon_sym_GT),
	7461:  uint16(aux_sym_preproc_import_token1),
	7462:  uint16(aux_sym_preproc_ifdef_token1),
	7463:  uint16(aux_sym_preproc_ifdef_token2),
	7464:  uint16(aux_sym_preproc_ifdef_token3),
	7465:  uint16(aux_sym_preproc_else_token1),
	7466:  uint16(5),
	7467:  uint16(664),
	7468:  uint16(1),
	7469:  uint16(anon_sym_LBRACK),
	7470:  uint16(666),
	7471:  uint16(1),
	7472:  uint16(anon_sym_DOT),
	7473:  uint16(182),
	7474:  uint16(1),
	7475:  uint16(sym_postfix_expression),
	7476:  uint16(3),
	7477:  uint16(2),
	7478:  uint16(sym_block_comment),
	7479:  uint16(sym_line_comment),
	7480:  uint16(662),
	7481:  uint16(12),
	7482:  uint16(anon_sym_EQ),
	7483:  uint16(anon_sym_RPAREN),
	7484:  uint16(anon_sym_PLUS_EQ),
	7485:  uint16(anon_sym_DASH_EQ),
	7486:  uint16(anon_sym_STAR_EQ),
	7487:  uint16(anon_sym_SLASH_EQ),
	7488:  uint16(anon_sym_PERCENT_EQ),
	7489:  uint16(anon_sym_AMP_EQ),
	7490:  uint16(anon_sym_PIPE_EQ),
	7491:  uint16(anon_sym_CARET_EQ),
	7492:  uint16(anon_sym_PLUS_PLUS),
	7493:  uint16(anon_sym_DASH_DASH),
	7494:  uint16(5),
	7495:  uint16(658),
	7496:  uint16(1),
	7497:  uint16(anon_sym_LBRACK),
	7498:  uint16(660),
	7499:  uint16(1),
	7500:  uint16(anon_sym_DOT),
	7501:  uint16(193),
	7502:  uint16(1),
	7503:  uint16(sym_postfix_expression),
	7504:  uint16(3),
	7505:  uint16(2),
	7506:  uint16(sym_block_comment),
	7507:  uint16(sym_line_comment),
	7508:  uint16(662),
	7509:  uint16(12),
	7510:  uint16(anon_sym_EQ),
	7511:  uint16(anon_sym_RPAREN),
	7512:  uint16(anon_sym_PLUS_EQ),
	7513:  uint16(anon_sym_DASH_EQ),
	7514:  uint16(anon_sym_STAR_EQ),
	7515:  uint16(anon_sym_SLASH_EQ),
	7516:  uint16(anon_sym_PERCENT_EQ),
	7517:  uint16(anon_sym_AMP_EQ),
	7518:  uint16(anon_sym_PIPE_EQ),
	7519:  uint16(anon_sym_CARET_EQ),
	7520:  uint16(anon_sym_PLUS_PLUS),
	7521:  uint16(anon_sym_DASH_DASH),
	7522:  uint16(5),
	7523:  uint16(658),
	7524:  uint16(1),
	7525:  uint16(anon_sym_LBRACK),
	7526:  uint16(660),
	7527:  uint16(1),
	7528:  uint16(anon_sym_DOT),
	7529:  uint16(200),
	7530:  uint16(1),
	7531:  uint16(sym_postfix_expression),
	7532:  uint16(3),
	7533:  uint16(2),
	7534:  uint16(sym_block_comment),
	7535:  uint16(sym_line_comment),
	7536:  uint16(668),
	7537:  uint16(12),
	7538:  uint16(anon_sym_EQ),
	7539:  uint16(anon_sym_RPAREN),
	7540:  uint16(anon_sym_PLUS_EQ),
	7541:  uint16(anon_sym_DASH_EQ),
	7542:  uint16(anon_sym_STAR_EQ),
	7543:  uint16(anon_sym_SLASH_EQ),
	7544:  uint16(anon_sym_PERCENT_EQ),
	7545:  uint16(anon_sym_AMP_EQ),
	7546:  uint16(anon_sym_PIPE_EQ),
	7547:  uint16(anon_sym_CARET_EQ),
	7548:  uint16(anon_sym_PLUS_PLUS),
	7549:  uint16(anon_sym_DASH_DASH),
	7550:  uint16(6),
	7551:  uint16(431),
	7552:  uint16(1),
	7553:  uint16(anon_sym_LPAREN),
	7554:  uint16(658),
	7555:  uint16(1),
	7556:  uint16(anon_sym_LBRACK),
	7557:  uint16(660),
	7558:  uint16(1),
	7559:  uint16(anon_sym_DOT),
	7560:  uint16(202),
	7561:  uint16(1),
	7562:  uint16(sym_postfix_expression),
	7563:  uint16(3),
	7564:  uint16(2),
	7565:  uint16(sym_block_comment),
	7566:  uint16(sym_line_comment),
	7567:  uint16(656),
	7568:  uint16(11),
	7569:  uint16(anon_sym_EQ),
	7570:  uint16(anon_sym_PLUS_EQ),
	7571:  uint16(anon_sym_DASH_EQ),
	7572:  uint16(anon_sym_STAR_EQ),
	7573:  uint16(anon_sym_SLASH_EQ),
	7574:  uint16(anon_sym_PERCENT_EQ),
	7575:  uint16(anon_sym_AMP_EQ),
	7576:  uint16(anon_sym_PIPE_EQ),
	7577:  uint16(anon_sym_CARET_EQ),
	7578:  uint16(anon_sym_PLUS_PLUS),
	7579:  uint16(anon_sym_DASH_DASH),
	7580:  uint16(5),
	7581:  uint16(658),
	7582:  uint16(1),
	7583:  uint16(anon_sym_LBRACK),
	7584:  uint16(660),
	7585:  uint16(1),
	7586:  uint16(anon_sym_DOT),
	7587:  uint16(204),
	7588:  uint16(1),
	7589:  uint16(sym_postfix_expression),
	7590:  uint16(3),
	7591:  uint16(2),
	7592:  uint16(sym_block_comment),
	7593:  uint16(sym_line_comment),
	7594:  uint16(670),
	7595:  uint16(12),
	7596:  uint16(anon_sym_EQ),
	7597:  uint16(anon_sym_RPAREN),
	7598:  uint16(anon_sym_PLUS_EQ),
	7599:  uint16(anon_sym_DASH_EQ),
	7600:  uint16(anon_sym_STAR_EQ),
	7601:  uint16(anon_sym_SLASH_EQ),
	7602:  uint16(anon_sym_PERCENT_EQ),
	7603:  uint16(anon_sym_AMP_EQ),
	7604:  uint16(anon_sym_PIPE_EQ),
	7605:  uint16(anon_sym_CARET_EQ),
	7606:  uint16(anon_sym_PLUS_PLUS),
	7607:  uint16(anon_sym_DASH_DASH),
	7608:  uint16(5),
	7609:  uint16(658),
	7610:  uint16(1),
	7611:  uint16(anon_sym_LBRACK),
	7612:  uint16(660),
	7613:  uint16(1),
	7614:  uint16(anon_sym_DOT),
	7615:  uint16(194),
	7616:  uint16(1),
	7617:  uint16(sym_postfix_expression),
	7618:  uint16(3),
	7619:  uint16(2),
	7620:  uint16(sym_block_comment),
	7621:  uint16(sym_line_comment),
	7622:  uint16(672),
	7623:  uint16(12),
	7624:  uint16(anon_sym_EQ),
	7625:  uint16(anon_sym_RPAREN),
	7626:  uint16(anon_sym_PLUS_EQ),
	7627:  uint16(anon_sym_DASH_EQ),
	7628:  uint16(anon_sym_STAR_EQ),
	7629:  uint16(anon_sym_SLASH_EQ),
	7630:  uint16(anon_sym_PERCENT_EQ),
	7631:  uint16(anon_sym_AMP_EQ),
	7632:  uint16(anon_sym_PIPE_EQ),
	7633:  uint16(anon_sym_CARET_EQ),
	7634:  uint16(anon_sym_PLUS_PLUS),
	7635:  uint16(anon_sym_DASH_DASH),
	7636:  uint16(12),
	7637:  uint16(23),
	7638:  uint16(1),
	7639:  uint16(anon_sym_AT),
	7640:  uint16(27),
	7641:  uint16(1),
	7642:  uint16(aux_sym_preproc_import_token1),
	7643:  uint16(674),
	7644:  uint16(1),
	7645:  uint16(sym_identifier),
	7646:  uint16(678),
	7647:  uint16(1),
	7648:  uint16(aux_sym_preproc_ifdef_token3),
	7649:  uint16(680),
	7650:  uint16(1),
	7651:  uint16(aux_sym_preproc_else_token1),
	7652:  uint16(191),
	7653:  uint16(1),
	7654:  uint16(aux_sym_preproc_ifdef_in_struct_declaration_repeat1),
	7655:  uint16(208),
	7656:  uint16(1),
	7657:  uint16(sym_variable_identifier_declaration),
	7658:  uint16(408),
	7659:  uint16(1),
	7660:  uint16(sym_preproc_else_in_struct_declaration),
	7661:  uint16(3),
	7662:  uint16(2),
	7663:  uint16(sym_block_comment),
	7664:  uint16(sym_line_comment),
	7665:  uint16(676),
	7666:  uint16(2),
	7667:  uint16(aux_sym_preproc_ifdef_token1),
	7668:  uint16(aux_sym_preproc_ifdef_token2),
	7669:  uint16(237),
	7670:  uint16(2),
	7671:  uint16(sym_attribute),
	7672:  uint16(aux_sym_global_variable_declaration_repeat1),
	7673:  uint16(219),
	7674:  uint16(3),
	7675:  uint16(sym_struct_member),
	7676:  uint16(sym_preproc_import),
	7677:  uint16(sym_preproc_ifdef_in_struct_declaration),
	7678:  uint16(5),
	7679:  uint16(658),
	7680:  uint16(1),
	7681:  uint16(anon_sym_LBRACK),
	7682:  uint16(660),
	7683:  uint16(1),
	7684:  uint16(anon_sym_DOT),
	7685:  uint16(199),
	7686:  uint16(1),
	7687:  uint16(sym_postfix_expression),
	7688:  uint16(3),
	7689:  uint16(2),
	7690:  uint16(sym_block_comment),
	7691:  uint16(sym_line_comment),
	7692:  uint16(682),
	7693:  uint16(12),
	7694:  uint16(anon_sym_EQ),
	7695:  uint16(anon_sym_RPAREN),
	7696:  uint16(anon_sym_PLUS_EQ),
	7697:  uint16(anon_sym_DASH_EQ),
	7698:  uint16(anon_sym_STAR_EQ),
	7699:  uint16(anon_sym_SLASH_EQ),
	7700:  uint16(anon_sym_PERCENT_EQ),
	7701:  uint16(anon_sym_AMP_EQ),
	7702:  uint16(anon_sym_PIPE_EQ),
	7703:  uint16(anon_sym_CARET_EQ),
	7704:  uint16(anon_sym_PLUS_PLUS),
	7705:  uint16(anon_sym_DASH_DASH),
	7706:  uint16(2),
	7707:  uint16(3),
	7708:  uint16(2),
	7709:  uint16(sym_block_comment),
	7710:  uint16(sym_line_comment),
	7711:  uint16(684),
	7712:  uint16(15),
	7713:  uint16(anon_sym_SEMI),
	7714:  uint16(anon_sym_EQ),
	7715:  uint16(anon_sym_LPAREN),
	7716:  uint16(anon_sym_COMMA),
	7717:  uint16(anon_sym_RPAREN),
	7718:  uint16(anon_sym_LBRACE),
	7719:  uint16(anon_sym_RBRACE),
	7720:  uint16(anon_sym_AT),
	7721:  uint16(sym_identifier),
	7722:  uint16(anon_sym_GT),
	7723:  uint16(aux_sym_preproc_import_token1),
	7724:  uint16(aux_sym_preproc_ifdef_token1),
	7725:  uint16(aux_sym_preproc_ifdef_token2),
	7726:  uint16(aux_sym_preproc_ifdef_token3),
	7727:  uint16(aux_sym_preproc_else_token1),
	7728:  uint16(5),
	7729:  uint16(664),
	7730:  uint16(1),
	7731:  uint16(anon_sym_LBRACK),
	7732:  uint16(666),
	7733:  uint16(1),
	7734:  uint16(anon_sym_DOT),
	7735:  uint16(178),
	7736:  uint16(1),
	7737:  uint16(sym_postfix_expression),
	7738:  uint16(3),
	7739:  uint16(2),
	7740:  uint16(sym_block_comment),
	7741:  uint16(sym_line_comment),
	7742:  uint16(686),
	7743:  uint16(12),
	7744:  uint16(anon_sym_EQ),
	7745:  uint16(anon_sym_RPAREN),
	7746:  uint16(anon_sym_PLUS_EQ),
	7747:  uint16(anon_sym_DASH_EQ),
	7748:  uint16(anon_sym_STAR_EQ),
	7749:  uint16(anon_sym_SLASH_EQ),
	7750:  uint16(anon_sym_PERCENT_EQ),
	7751:  uint16(anon_sym_AMP_EQ),
	7752:  uint16(anon_sym_PIPE_EQ),
	7753:  uint16(anon_sym_CARET_EQ),
	7754:  uint16(anon_sym_PLUS_PLUS),
	7755:  uint16(anon_sym_DASH_DASH),
	7756:  uint16(3),
	7757:  uint16(189),
	7758:  uint16(1),
	7759:  uint16(sym_postfix_expression),
	7760:  uint16(3),
	7761:  uint16(2),
	7762:  uint16(sym_block_comment),
	7763:  uint16(sym_line_comment),
	7764:  uint16(686),
	7765:  uint16(14),
	7766:  uint16(anon_sym_EQ),
	7767:  uint16(anon_sym_RPAREN),
	7768:  uint16(anon_sym_PLUS_EQ),
	7769:  uint16(anon_sym_DASH_EQ),
	7770:  uint16(anon_sym_STAR_EQ),
	7771:  uint16(anon_sym_SLASH_EQ),
	7772:  uint16(anon_sym_PERCENT_EQ),
	7773:  uint16(anon_sym_AMP_EQ),
	7774:  uint16(anon_sym_PIPE_EQ),
	7775:  uint16(anon_sym_CARET_EQ),
	7776:  uint16(anon_sym_PLUS_PLUS),
	7777:  uint16(anon_sym_DASH_DASH),
	7778:  uint16(anon_sym_LBRACK),
	7779:  uint16(anon_sym_DOT),
	7780:  uint16(3),
	7781:  uint16(190),
	7782:  uint16(1),
	7783:  uint16(sym_postfix_expression),
	7784:  uint16(3),
	7785:  uint16(2),
	7786:  uint16(sym_block_comment),
	7787:  uint16(sym_line_comment),
	7788:  uint16(662),
	7789:  uint16(14),
	7790:  uint16(anon_sym_EQ),
	7791:  uint16(anon_sym_RPAREN),
	7792:  uint16(anon_sym_PLUS_EQ),
	7793:  uint16(anon_sym_DASH_EQ),
	7794:  uint16(anon_sym_STAR_EQ),
	7795:  uint16(anon_sym_SLASH_EQ),
	7796:  uint16(anon_sym_PERCENT_EQ),
	7797:  uint16(anon_sym_AMP_EQ),
	7798:  uint16(anon_sym_PIPE_EQ),
	7799:  uint16(anon_sym_CARET_EQ),
	7800:  uint16(anon_sym_PLUS_PLUS),
	7801:  uint16(anon_sym_DASH_DASH),
	7802:  uint16(anon_sym_LBRACK),
	7803:  uint16(anon_sym_DOT),
	7804:  uint16(3),
	7805:  uint16(193),
	7806:  uint16(1),
	7807:  uint16(sym_postfix_expression),
	7808:  uint16(3),
	7809:  uint16(2),
	7810:  uint16(sym_block_comment),
	7811:  uint16(sym_line_comment),
	7812:  uint16(662),
	7813:  uint16(14),
	7814:  uint16(anon_sym_EQ),
	7815:  uint16(anon_sym_RPAREN),
	7816:  uint16(anon_sym_PLUS_EQ),
	7817:  uint16(anon_sym_DASH_EQ),
	7818:  uint16(anon_sym_STAR_EQ),
	7819:  uint16(anon_sym_SLASH_EQ),
	7820:  uint16(anon_sym_PERCENT_EQ),
	7821:  uint16(anon_sym_AMP_EQ),
	7822:  uint16(anon_sym_PIPE_EQ),
	7823:  uint16(anon_sym_CARET_EQ),
	7824:  uint16(anon_sym_PLUS_PLUS),
	7825:  uint16(anon_sym_DASH_DASH),
	7826:  uint16(anon_sym_LBRACK),
	7827:  uint16(anon_sym_DOT),
	7828:  uint16(3),
	7829:  uint16(194),
	7830:  uint16(1),
	7831:  uint16(sym_postfix_expression),
	7832:  uint16(3),
	7833:  uint16(2),
	7834:  uint16(sym_block_comment),
	7835:  uint16(sym_line_comment),
	7836:  uint16(672),
	7837:  uint16(14),
	7838:  uint16(anon_sym_EQ),
	7839:  uint16(anon_sym_RPAREN),
	7840:  uint16(anon_sym_PLUS_EQ),
	7841:  uint16(anon_sym_DASH_EQ),
	7842:  uint16(anon_sym_STAR_EQ),
	7843:  uint16(anon_sym_SLASH_EQ),
	7844:  uint16(anon_sym_PERCENT_EQ),
	7845:  uint16(anon_sym_AMP_EQ),
	7846:  uint16(anon_sym_PIPE_EQ),
	7847:  uint16(anon_sym_CARET_EQ),
	7848:  uint16(anon_sym_PLUS_PLUS),
	7849:  uint16(anon_sym_DASH_DASH),
	7850:  uint16(anon_sym_LBRACK),
	7851:  uint16(anon_sym_DOT),
	7852:  uint16(12),
	7853:  uint16(23),
	7854:  uint16(1),
	7855:  uint16(anon_sym_AT),
	7856:  uint16(27),
	7857:  uint16(1),
	7858:  uint16(aux_sym_preproc_import_token1),
	7859:  uint16(674),
	7860:  uint16(1),
	7861:  uint16(sym_identifier),
	7862:  uint16(680),
	7863:  uint16(1),
	7864:  uint16(aux_sym_preproc_else_token1),
	7865:  uint16(688),
	7866:  uint16(1),
	7867:  uint16(aux_sym_preproc_ifdef_token3),
	7868:  uint16(192),
	7869:  uint16(1),
	7870:  uint16(aux_sym_preproc_ifdef_in_struct_declaration_repeat1),
	7871:  uint16(208),
	7872:  uint16(1),
	7873:  uint16(sym_variable_identifier_declaration),
	7874:  uint16(372),
	7875:  uint16(1),
	7876:  uint16(sym_preproc_else_in_struct_declaration),
	7877:  uint16(3),
	7878:  uint16(2),
	7879:  uint16(sym_block_comment),
	7880:  uint16(sym_line_comment),
	7881:  uint16(676),
	7882:  uint16(2),
	7883:  uint16(aux_sym_preproc_ifdef_token1),
	7884:  uint16(aux_sym_preproc_ifdef_token2),
	7885:  uint16(237),
	7886:  uint16(2),
	7887:  uint16(sym_attribute),
	7888:  uint16(aux_sym_global_variable_declaration_repeat1),
	7889:  uint16(219),
	7890:  uint16(3),
	7891:  uint16(sym_struct_member),
	7892:  uint16(sym_preproc_import),
	7893:  uint16(sym_preproc_ifdef_in_struct_declaration),
	7894:  uint16(10),
	7895:  uint16(690),
	7896:  uint16(1),
	7897:  uint16(sym_identifier),
	7898:  uint16(693),
	7899:  uint16(1),
	7900:  uint16(anon_sym_AT),
	7901:  uint16(696),
	7902:  uint16(1),
	7903:  uint16(aux_sym_preproc_import_token1),
	7904:  uint16(192),
	7905:  uint16(1),
	7906:  uint16(aux_sym_preproc_ifdef_in_struct_declaration_repeat1),
	7907:  uint16(208),
	7908:  uint16(1),
	7909:  uint16(sym_variable_identifier_declaration),
	7910:  uint16(3),
	7911:  uint16(2),
	7912:  uint16(sym_block_comment),
	7913:  uint16(sym_line_comment),
	7914:  uint16(699),
	7915:  uint16(2),
	7916:  uint16(aux_sym_preproc_ifdef_token1),
	7917:  uint16(aux_sym_preproc_ifdef_token2),
	7918:  uint16(702),
	7919:  uint16(2),
	7920:  uint16(aux_sym_preproc_ifdef_token3),
	7921:  uint16(aux_sym_preproc_else_token1),
	7922:  uint16(237),
	7923:  uint16(2),
	7924:  uint16(sym_attribute),
	7925:  uint16(aux_sym_global_variable_declaration_repeat1),
	7926:  uint16(219),
	7927:  uint16(3),
	7928:  uint16(sym_struct_member),
	7929:  uint16(sym_preproc_import),
	7930:  uint16(sym_preproc_ifdef_in_struct_declaration),
	7931:  uint16(2),
	7932:  uint16(3),
	7933:  uint16(2),
	7934:  uint16(sym_block_comment),
	7935:  uint16(sym_line_comment),
	7936:  uint16(672),
	7937:  uint16(14),
	7938:  uint16(anon_sym_EQ),
	7939:  uint16(anon_sym_RPAREN),
	7940:  uint16(anon_sym_PLUS_EQ),
	7941:  uint16(anon_sym_DASH_EQ),
	7942:  uint16(anon_sym_STAR_EQ),
	7943:  uint16(anon_sym_SLASH_EQ),
	7944:  uint16(anon_sym_PERCENT_EQ),
	7945:  uint16(anon_sym_AMP_EQ),
	7946:  uint16(anon_sym_PIPE_EQ),
	7947:  uint16(anon_sym_CARET_EQ),
	7948:  uint16(anon_sym_PLUS_PLUS),
	7949:  uint16(anon_sym_DASH_DASH),
	7950:  uint16(anon_sym_LBRACK),
	7951:  uint16(anon_sym_DOT),
	7952:  uint16(2),
	7953:  uint16(3),
	7954:  uint16(2),
	7955:  uint16(sym_block_comment),
	7956:  uint16(sym_line_comment),
	7957:  uint16(704),
	7958:  uint16(14),
	7959:  uint16(anon_sym_EQ),
	7960:  uint16(anon_sym_RPAREN),
	7961:  uint16(anon_sym_PLUS_EQ),
	7962:  uint16(anon_sym_DASH_EQ),
	7963:  uint16(anon_sym_STAR_EQ),
	7964:  uint16(anon_sym_SLASH_EQ),
	7965:  uint16(anon_sym_PERCENT_EQ),
	7966:  uint16(anon_sym_AMP_EQ),
	7967:  uint16(anon_sym_PIPE_EQ),
	7968:  uint16(anon_sym_CARET_EQ),
	7969:  uint16(anon_sym_PLUS_PLUS),
	7970:  uint16(anon_sym_DASH_DASH),
	7971:  uint16(anon_sym_LBRACK),
	7972:  uint16(anon_sym_DOT),
	7973:  uint16(10),
	7974:  uint16(23),
	7975:  uint16(1),
	7976:  uint16(anon_sym_AT),
	7977:  uint16(27),
	7978:  uint16(1),
	7979:  uint16(aux_sym_preproc_import_token1),
	7980:  uint16(674),
	7981:  uint16(1),
	7982:  uint16(sym_identifier),
	7983:  uint16(706),
	7984:  uint16(1),
	7985:  uint16(aux_sym_preproc_ifdef_token3),
	7986:  uint16(196),
	7987:  uint16(1),
	7988:  uint16(aux_sym_preproc_ifdef_in_struct_declaration_repeat1),
	7989:  uint16(208),
	7990:  uint16(1),
	7991:  uint16(sym_variable_identifier_declaration),
	7992:  uint16(3),
	7993:  uint16(2),
	7994:  uint16(sym_block_comment),
	7995:  uint16(sym_line_comment),
	7996:  uint16(676),
	7997:  uint16(2),
	7998:  uint16(aux_sym_preproc_ifdef_token1),
	7999:  uint16(aux_sym_preproc_ifdef_token2),
	8000:  uint16(237),
	8001:  uint16(2),
	8002:  uint16(sym_attribute),
	8003:  uint16(aux_sym_global_variable_declaration_repeat1),
	8004:  uint16(219),
	8005:  uint16(3),
	8006:  uint16(sym_struct_member),
	8007:  uint16(sym_preproc_import),
	8008:  uint16(sym_preproc_ifdef_in_struct_declaration),
	8009:  uint16(10),
	8010:  uint16(23),
	8011:  uint16(1),
	8012:  uint16(anon_sym_AT),
	8013:  uint16(27),
	8014:  uint16(1),
	8015:  uint16(aux_sym_preproc_import_token1),
	8016:  uint16(674),
	8017:  uint16(1),
	8018:  uint16(sym_identifier),
	8019:  uint16(708),
	8020:  uint16(1),
	8021:  uint16(aux_sym_preproc_ifdef_token3),
	8022:  uint16(192),
	8023:  uint16(1),
	8024:  uint16(aux_sym_preproc_ifdef_in_struct_declaration_repeat1),
	8025:  uint16(208),
	8026:  uint16(1),
	8027:  uint16(sym_variable_identifier_declaration),
	8028:  uint16(3),
	8029:  uint16(2),
	8030:  uint16(sym_block_comment),
	8031:  uint16(sym_line_comment),
	8032:  uint16(676),
	8033:  uint16(2),
	8034:  uint16(aux_sym_preproc_ifdef_token1),
	8035:  uint16(aux_sym_preproc_ifdef_token2),
	8036:  uint16(237),
	8037:  uint16(2),
	8038:  uint16(sym_attribute),
	8039:  uint16(aux_sym_global_variable_declaration_repeat1),
	8040:  uint16(219),
	8041:  uint16(3),
	8042:  uint16(sym_struct_member),
	8043:  uint16(sym_preproc_import),
	8044:  uint16(sym_preproc_ifdef_in_struct_declaration),
	8045:  uint16(11),
	8046:  uint16(23),
	8047:  uint16(1),
	8048:  uint16(anon_sym_AT),
	8049:  uint16(27),
	8050:  uint16(1),
	8051:  uint16(aux_sym_preproc_import_token1),
	8052:  uint16(674),
	8053:  uint16(1),
	8054:  uint16(sym_identifier),
	8055:  uint16(203),
	8056:  uint16(1),
	8057:  uint16(aux_sym__struct_declaration_content_repeat1),
	8058:  uint16(208),
	8059:  uint16(1),
	8060:  uint16(sym_variable_identifier_declaration),
	8061:  uint16(308),
	8062:  uint16(1),
	8063:  uint16(sym_struct_member),
	8064:  uint16(362),
	8065:  uint16(1),
	8066:  uint16(sym__struct_declaration_content),
	8067:  uint16(3),
	8068:  uint16(2),
	8069:  uint16(sym_block_comment),
	8070:  uint16(sym_line_comment),
	8071:  uint16(676),
	8072:  uint16(2),
	8073:  uint16(aux_sym_preproc_ifdef_token1),
	8074:  uint16(aux_sym_preproc_ifdef_token2),
	8075:  uint16(225),
	8076:  uint16(2),
	8077:  uint16(sym_preproc_import),
	8078:  uint16(sym_preproc_ifdef_in_struct_declaration),
	8079:  uint16(237),
	8080:  uint16(2),
	8081:  uint16(sym_attribute),
	8082:  uint16(aux_sym_global_variable_declaration_repeat1),
	8083:  uint16(2),
	8084:  uint16(3),
	8085:  uint16(2),
	8086:  uint16(sym_block_comment),
	8087:  uint16(sym_line_comment),
	8088:  uint16(710),
	8089:  uint16(12),
	8090:  uint16(anon_sym_SEMI),
	8091:  uint16(anon_sym_EQ),
	8092:  uint16(anon_sym_COMMA),
	8093:  uint16(anon_sym_RPAREN),
	8094:  uint16(anon_sym_RBRACE),
	8095:  uint16(anon_sym_AT),
	8096:  uint16(sym_identifier),
	8097:  uint16(aux_sym_preproc_import_token1),
	8098:  uint16(aux_sym_preproc_ifdef_token1),
	8099:  uint16(aux_sym_preproc_ifdef_token2),
	8100:  uint16(aux_sym_preproc_ifdef_token3),
	8101:  uint16(aux_sym_preproc_else_token1),
	8102:  uint16(2),
	8103:  uint16(3),
	8104:  uint16(2),
	8105:  uint16(sym_block_comment),
	8106:  uint16(sym_line_comment),
	8107:  uint16(670),
	8108:  uint16(12),
	8109:  uint16(anon_sym_EQ),
	8110:  uint16(anon_sym_RPAREN),
	8111:  uint16(anon_sym_PLUS_EQ),
	8112:  uint16(anon_sym_DASH_EQ),
	8113:  uint16(anon_sym_STAR_EQ),
	8114:  uint16(anon_sym_SLASH_EQ),
	8115:  uint16(anon_sym_PERCENT_EQ),
	8116:  uint16(anon_sym_AMP_EQ),
	8117:  uint16(anon_sym_PIPE_EQ),
	8118:  uint16(anon_sym_CARET_EQ),
	8119:  uint16(anon_sym_PLUS_PLUS),
	8120:  uint16(anon_sym_DASH_DASH),
	8121:  uint16(2),
	8122:  uint16(3),
	8123:  uint16(2),
	8124:  uint16(sym_block_comment),
	8125:  uint16(sym_line_comment),
	8126:  uint16(682),
	8127:  uint16(12),
	8128:  uint16(anon_sym_EQ),
	8129:  uint16(anon_sym_RPAREN),
	8130:  uint16(anon_sym_PLUS_EQ),
	8131:  uint16(anon_sym_DASH_EQ),
	8132:  uint16(anon_sym_STAR_EQ),
	8133:  uint16(anon_sym_SLASH_EQ),
	8134:  uint16(anon_sym_PERCENT_EQ),
	8135:  uint16(anon_sym_AMP_EQ),
	8136:  uint16(anon_sym_PIPE_EQ),
	8137:  uint16(anon_sym_CARET_EQ),
	8138:  uint16(anon_sym_PLUS_PLUS),
	8139:  uint16(anon_sym_DASH_DASH),
	8140:  uint16(6),
	8141:  uint16(712),
	8142:  uint16(1),
	8143:  uint16(anon_sym_EQ),
	8144:  uint16(716),
	8145:  uint16(1),
	8146:  uint16(anon_sym_PLUS_PLUS),
	8147:  uint16(718),
	8148:  uint16(1),
	8149:  uint16(anon_sym_DASH_DASH),
	8150:  uint16(26),
	8151:  uint16(1),
	8152:  uint16(sym_compound_assignment_operator),
	8153:  uint16(3),
	8154:  uint16(2),
	8155:  uint16(sym_block_comment),
	8156:  uint16(sym_line_comment),
	8157:  uint16(714),
	8158:  uint16(8),
	8159:  uint16(anon_sym_PLUS_EQ),
	8160:  uint16(anon_sym_DASH_EQ),
	8161:  uint16(anon_sym_STAR_EQ),
	8162:  uint16(anon_sym_SLASH_EQ),
	8163:  uint16(anon_sym_PERCENT_EQ),
	8164:  uint16(anon_sym_AMP_EQ),
	8165:  uint16(anon_sym_PIPE_EQ),
	8166:  uint16(anon_sym_CARET_EQ),
	8167:  uint16(2),
	8168:  uint16(3),
	8169:  uint16(2),
	8170:  uint16(sym_block_comment),
	8171:  uint16(sym_line_comment),
	8172:  uint16(668),
	8173:  uint16(12),
	8174:  uint16(anon_sym_EQ),
	8175:  uint16(anon_sym_RPAREN),
	8176:  uint16(anon_sym_PLUS_EQ),
	8177:  uint16(anon_sym_DASH_EQ),
	8178:  uint16(anon_sym_STAR_EQ),
	8179:  uint16(anon_sym_SLASH_EQ),
	8180:  uint16(anon_sym_PERCENT_EQ),
	8181:  uint16(anon_sym_AMP_EQ),
	8182:  uint16(anon_sym_PIPE_EQ),
	8183:  uint16(anon_sym_CARET_EQ),
	8184:  uint16(anon_sym_PLUS_PLUS),
	8185:  uint16(anon_sym_DASH_DASH),
	8186:  uint16(10),
	8187:  uint16(23),
	8188:  uint16(1),
	8189:  uint16(anon_sym_AT),
	8190:  uint16(27),
	8191:  uint16(1),
	8192:  uint16(aux_sym_preproc_import_token1),
	8193:  uint16(674),
	8194:  uint16(1),
	8195:  uint16(sym_identifier),
	8196:  uint16(205),
	8197:  uint16(1),
	8198:  uint16(aux_sym__struct_declaration_content_repeat1),
	8199:  uint16(208),
	8200:  uint16(1),
	8201:  uint16(sym_variable_identifier_declaration),
	8202:  uint16(336),
	8203:  uint16(1),
	8204:  uint16(sym_struct_member),
	8205:  uint16(3),
	8206:  uint16(2),
	8207:  uint16(sym_block_comment),
	8208:  uint16(sym_line_comment),
	8209:  uint16(676),
	8210:  uint16(2),
	8211:  uint16(aux_sym_preproc_ifdef_token1),
	8212:  uint16(aux_sym_preproc_ifdef_token2),
	8213:  uint16(228),
	8214:  uint16(2),
	8215:  uint16(sym_preproc_import),
	8216:  uint16(sym_preproc_ifdef_in_struct_declaration),
	8217:  uint16(237),
	8218:  uint16(2),
	8219:  uint16(sym_attribute),
	8220:  uint16(aux_sym_global_variable_declaration_repeat1),
	8221:  uint16(2),
	8222:  uint16(3),
	8223:  uint16(2),
	8224:  uint16(sym_block_comment),
	8225:  uint16(sym_line_comment),
	8226:  uint16(720),
	8227:  uint16(12),
	8228:  uint16(anon_sym_EQ),
	8229:  uint16(anon_sym_RPAREN),
	8230:  uint16(anon_sym_PLUS_EQ),
	8231:  uint16(anon_sym_DASH_EQ),
	8232:  uint16(anon_sym_STAR_EQ),
	8233:  uint16(anon_sym_SLASH_EQ),
	8234:  uint16(anon_sym_PERCENT_EQ),
	8235:  uint16(anon_sym_AMP_EQ),
	8236:  uint16(anon_sym_PIPE_EQ),
	8237:  uint16(anon_sym_CARET_EQ),
	8238:  uint16(anon_sym_PLUS_PLUS),
	8239:  uint16(anon_sym_DASH_DASH),
	8240:  uint16(9),
	8241:  uint16(722),
	8242:  uint16(1),
	8243:  uint16(sym_identifier),
	8244:  uint16(725),
	8245:  uint16(1),
	8246:  uint16(anon_sym_AT),
	8247:  uint16(728),
	8248:  uint16(1),
	8249:  uint16(aux_sym_preproc_import_token1),
	8250:  uint16(208),
	8251:  uint16(1),
	8252:  uint16(sym_variable_identifier_declaration),
	8253:  uint16(353),
	8254:  uint16(1),
	8255:  uint16(sym_struct_member),
	8256:  uint16(3),
	8257:  uint16(2),
	8258:  uint16(sym_block_comment),
	8259:  uint16(sym_line_comment),
	8260:  uint16(731),
	8261:  uint16(2),
	8262:  uint16(aux_sym_preproc_ifdef_token1),
	8263:  uint16(aux_sym_preproc_ifdef_token2),
	8264:  uint16(237),
	8265:  uint16(2),
	8266:  uint16(sym_attribute),
	8267:  uint16(aux_sym_global_variable_declaration_repeat1),
	8268:  uint16(205),
	8269:  uint16(3),
	8270:  uint16(sym_preproc_import),
	8271:  uint16(sym_preproc_ifdef_in_struct_declaration),
	8272:  uint16(aux_sym__struct_declaration_content_repeat1),
	8273:  uint16(7),
	8274:  uint16(39),
	8275:  uint16(1),
	8276:  uint16(sym_int_literal),
	8277:  uint16(289),
	8278:  uint16(1),
	8279:  uint16(sym_const_literal),
	8280:  uint16(3),
	8281:  uint16(2),
	8282:  uint16(sym_block_comment),
	8283:  uint16(sym_line_comment),
	8284:  uint16(41),
	8285:  uint16(2),
	8286:  uint16(aux_sym_float_literal_token1),
	8287:  uint16(aux_sym_float_literal_token2),
	8288:  uint16(734),
	8289:  uint16(2),
	8290:  uint16(anon_sym_LBRACE),
	8291:  uint16(anon_sym_COLON),
	8292:  uint16(736),
	8293:  uint16(2),
	8294:  uint16(anon_sym_true),
	8295:  uint16(anon_sym_false),
	8296:  uint16(93),
	8297:  uint16(2),
	8298:  uint16(sym_float_literal),
	8299:  uint16(sym_bool_literal),
	8300:  uint16(7),
	8301:  uint16(39),
	8302:  uint16(1),
	8303:  uint16(sym_int_literal),
	8304:  uint16(289),
	8305:  uint16(1),
	8306:  uint16(sym_const_literal),
	8307:  uint16(3),
	8308:  uint16(2),
	8309:  uint16(sym_block_comment),
	8310:  uint16(sym_line_comment),
	8311:  uint16(41),
	8312:  uint16(2),
	8313:  uint16(aux_sym_float_literal_token1),
	8314:  uint16(aux_sym_float_literal_token2),
	8315:  uint16(736),
	8316:  uint16(2),
	8317:  uint16(anon_sym_true),
	8318:  uint16(anon_sym_false),
	8319:  uint16(738),
	8320:  uint16(2),
	8321:  uint16(anon_sym_LBRACE),
	8322:  uint16(anon_sym_COLON),
	8323:  uint16(93),
	8324:  uint16(2),
	8325:  uint16(sym_float_literal),
	8326:  uint16(sym_bool_literal),
	8327:  uint16(2),
	8328:  uint16(3),
	8329:  uint16(2),
	8330:  uint16(sym_block_comment),
	8331:  uint16(sym_line_comment),
	8332:  uint16(740),
	8333:  uint16(9),
	8334:  uint16(anon_sym_COMMA),
	8335:  uint16(anon_sym_RBRACE),
	8336:  uint16(anon_sym_AT),
	8337:  uint16(sym_identifier),
	8338:  uint16(aux_sym_preproc_import_token1),
	8339:  uint16(aux_sym_preproc_ifdef_token1),
	8340:  uint16(aux_sym_preproc_ifdef_token2),
	8341:  uint16(aux_sym_preproc_ifdef_token3),
	8342:  uint16(aux_sym_preproc_else_token1),
	8343:  uint16(9),
	8344:  uint16(23),
	8345:  uint16(1),
	8346:  uint16(anon_sym_AT),
	8347:  uint16(674),
	8348:  uint16(1),
	8349:  uint16(sym_identifier),
	8350:  uint16(742),
	8351:  uint16(1),
	8352:  uint16(anon_sym_RPAREN),
	8353:  uint16(224),
	8354:  uint16(1),
	8355:  uint16(aux_sym_parameter_list_repeat1),
	8356:  uint16(341),
	8357:  uint16(1),
	8358:  uint16(sym_parameter),
	8359:  uint16(342),
	8360:  uint16(1),
	8361:  uint16(sym_variable_identifier_declaration),
	8362:  uint16(412),
	8363:  uint16(1),
	8364:  uint16(sym_parameter_list),
	8365:  uint16(3),
	8366:  uint16(2),
	8367:  uint16(sym_block_comment),
	8368:  uint16(sym_line_comment),
	8369:  uint16(236),
	8370:  uint16(2),
	8371:  uint16(sym_attribute),
	8372:  uint16(aux_sym_global_variable_declaration_repeat1),
	8373:  uint16(2),
	8374:  uint16(3),
	8375:  uint16(2),
	8376:  uint16(sym_block_comment),
	8377:  uint16(sym_line_comment),
	8378:  uint16(744),
	8379:  uint16(9),
	8380:  uint16(anon_sym_COMMA),
	8381:  uint16(anon_sym_RBRACE),
	8382:  uint16(anon_sym_AT),
	8383:  uint16(sym_identifier),
	8384:  uint16(aux_sym_preproc_import_token1),
	8385:  uint16(aux_sym_preproc_ifdef_token1),
	8386:  uint16(aux_sym_preproc_ifdef_token2),
	8387:  uint16(aux_sym_preproc_ifdef_token3),
	8388:  uint16(aux_sym_preproc_else_token1),
	8389:  uint16(2),
	8390:  uint16(3),
	8391:  uint16(2),
	8392:  uint16(sym_block_comment),
	8393:  uint16(sym_line_comment),
	8394:  uint16(746),
	8395:  uint16(9),
	8396:  uint16(anon_sym_COMMA),
	8397:  uint16(anon_sym_RBRACE),
	8398:  uint16(anon_sym_AT),
	8399:  uint16(sym_identifier),
	8400:  uint16(aux_sym_preproc_import_token1),
	8401:  uint16(aux_sym_preproc_ifdef_token1),
	8402:  uint16(aux_sym_preproc_ifdef_token2),
	8403:  uint16(aux_sym_preproc_ifdef_token3),
	8404:  uint16(aux_sym_preproc_else_token1),
	8405:  uint16(2),
	8406:  uint16(3),
	8407:  uint16(2),
	8408:  uint16(sym_block_comment),
	8409:  uint16(sym_line_comment),
	8410:  uint16(748),
	8411:  uint16(9),
	8412:  uint16(anon_sym_COMMA),
	8413:  uint16(anon_sym_RBRACE),
	8414:  uint16(anon_sym_AT),
	8415:  uint16(sym_identifier),
	8416:  uint16(aux_sym_preproc_import_token1),
	8417:  uint16(aux_sym_preproc_ifdef_token1),
	8418:  uint16(aux_sym_preproc_ifdef_token2),
	8419:  uint16(aux_sym_preproc_ifdef_token3),
	8420:  uint16(aux_sym_preproc_else_token1),
	8421:  uint16(7),
	8422:  uint16(39),
	8423:  uint16(1),
	8424:  uint16(sym_int_literal),
	8425:  uint16(264),
	8426:  uint16(1),
	8427:  uint16(sym_const_literal),
	8428:  uint16(277),
	8429:  uint16(1),
	8430:  uint16(sym_case_selectors),
	8431:  uint16(3),
	8432:  uint16(2),
	8433:  uint16(sym_block_comment),
	8434:  uint16(sym_line_comment),
	8435:  uint16(41),
	8436:  uint16(2),
	8437:  uint16(aux_sym_float_literal_token1),
	8438:  uint16(aux_sym_float_literal_token2),
	8439:  uint16(736),
	8440:  uint16(2),
	8441:  uint16(anon_sym_true),
	8442:  uint16(anon_sym_false),
	8443:  uint16(93),
	8444:  uint16(2),
	8445:  uint16(sym_float_literal),
	8446:  uint16(sym_bool_literal),
	8447:  uint16(9),
	8448:  uint16(23),
	8449:  uint16(1),
	8450:  uint16(anon_sym_AT),
	8451:  uint16(674),
	8452:  uint16(1),
	8453:  uint16(sym_identifier),
	8454:  uint16(750),
	8455:  uint16(1),
	8456:  uint16(anon_sym_RPAREN),
	8457:  uint16(224),
	8458:  uint16(1),
	8459:  uint16(aux_sym_parameter_list_repeat1),
	8460:  uint16(341),
	8461:  uint16(1),
	8462:  uint16(sym_parameter),
	8463:  uint16(342),
	8464:  uint16(1),
	8465:  uint16(sym_variable_identifier_declaration),
	8466:  uint16(424),
	8467:  uint16(1),
	8468:  uint16(sym_parameter_list),
	8469:  uint16(3),
	8470:  uint16(2),
	8471:  uint16(sym_block_comment),
	8472:  uint16(sym_line_comment),
	8473:  uint16(236),
	8474:  uint16(2),
	8475:  uint16(sym_attribute),
	8476:  uint16(aux_sym_global_variable_declaration_repeat1),
	8477:  uint16(2),
	8478:  uint16(3),
	8479:  uint16(2),
	8480:  uint16(sym_block_comment),
	8481:  uint16(sym_line_comment),
	8482:  uint16(752),
	8483:  uint16(9),
	8484:  uint16(anon_sym_COMMA),
	8485:  uint16(anon_sym_RBRACE),
	8486:  uint16(anon_sym_AT),
	8487:  uint16(sym_identifier),
	8488:  uint16(aux_sym_preproc_import_token1),
	8489:  uint16(aux_sym_preproc_ifdef_token1),
	8490:  uint16(aux_sym_preproc_ifdef_token2),
	8491:  uint16(aux_sym_preproc_ifdef_token3),
	8492:  uint16(aux_sym_preproc_else_token1),
	8493:  uint16(2),
	8494:  uint16(3),
	8495:  uint16(2),
	8496:  uint16(sym_block_comment),
	8497:  uint16(sym_line_comment),
	8498:  uint16(754),
	8499:  uint16(9),
	8500:  uint16(anon_sym_COMMA),
	8501:  uint16(anon_sym_RBRACE),
	8502:  uint16(anon_sym_AT),
	8503:  uint16(sym_identifier),
	8504:  uint16(aux_sym_preproc_import_token1),
	8505:  uint16(aux_sym_preproc_ifdef_token1),
	8506:  uint16(aux_sym_preproc_ifdef_token2),
	8507:  uint16(aux_sym_preproc_ifdef_token3),
	8508:  uint16(aux_sym_preproc_else_token1),
	8509:  uint16(9),
	8510:  uint16(23),
	8511:  uint16(1),
	8512:  uint16(anon_sym_AT),
	8513:  uint16(674),
	8514:  uint16(1),
	8515:  uint16(sym_identifier),
	8516:  uint16(756),
	8517:  uint16(1),
	8518:  uint16(anon_sym_RPAREN),
	8519:  uint16(224),
	8520:  uint16(1),
	8521:  uint16(aux_sym_parameter_list_repeat1),
	8522:  uint16(341),
	8523:  uint16(1),
	8524:  uint16(sym_parameter),
	8525:  uint16(342),
	8526:  uint16(1),
	8527:  uint16(sym_variable_identifier_declaration),
	8528:  uint16(403),
	8529:  uint16(1),
	8530:  uint16(sym_parameter_list),
	8531:  uint16(3),
	8532:  uint16(2),
	8533:  uint16(sym_block_comment),
	8534:  uint16(sym_line_comment),
	8535:  uint16(236),
	8536:  uint16(2),
	8537:  uint16(sym_attribute),
	8538:  uint16(aux_sym_global_variable_declaration_repeat1),
	8539:  uint16(6),
	8540:  uint16(39),
	8541:  uint16(1),
	8542:  uint16(sym_int_literal),
	8543:  uint16(289),
	8544:  uint16(1),
	8545:  uint16(sym_const_literal),
	8546:  uint16(3),
	8547:  uint16(2),
	8548:  uint16(sym_block_comment),
	8549:  uint16(sym_line_comment),
	8550:  uint16(41),
	8551:  uint16(2),
	8552:  uint16(aux_sym_float_literal_token1),
	8553:  uint16(aux_sym_float_literal_token2),
	8554:  uint16(736),
	8555:  uint16(2),
	8556:  uint16(anon_sym_true),
	8557:  uint16(anon_sym_false),
	8558:  uint16(93),
	8559:  uint16(2),
	8560:  uint16(sym_float_literal),
	8561:  uint16(sym_bool_literal),
	8562:  uint16(3),
	8563:  uint16(760),
	8564:  uint16(1),
	8565:  uint16(anon_sym_COMMA),
	8566:  uint16(3),
	8567:  uint16(2),
	8568:  uint16(sym_block_comment),
	8569:  uint16(sym_line_comment),
	8570:  uint16(758),
	8571:  uint16(7),
	8572:  uint16(anon_sym_AT),
	8573:  uint16(sym_identifier),
	8574:  uint16(aux_sym_preproc_import_token1),
	8575:  uint16(aux_sym_preproc_ifdef_token1),
	8576:  uint16(aux_sym_preproc_ifdef_token2),
	8577:  uint16(aux_sym_preproc_ifdef_token3),
	8578:  uint16(aux_sym_preproc_else_token1),
	8579:  uint16(7),
	8580:  uint16(762),
	8581:  uint16(1),
	8582:  uint16(sym_identifier),
	8583:  uint16(765),
	8584:  uint16(1),
	8585:  uint16(anon_sym_AT),
	8586:  uint16(220),
	8587:  uint16(1),
	8588:  uint16(aux_sym_parameter_list_repeat1),
	8589:  uint16(342),
	8590:  uint16(1),
	8591:  uint16(sym_variable_identifier_declaration),
	8592:  uint16(378),
	8593:  uint16(1),
	8594:  uint16(sym_parameter),
	8595:  uint16(3),
	8596:  uint16(2),
	8597:  uint16(sym_block_comment),
	8598:  uint16(sym_line_comment),
	8599:  uint16(236),
	8600:  uint16(2),
	8601:  uint16(sym_attribute),
	8602:  uint16(aux_sym_global_variable_declaration_repeat1),
	8603:  uint16(6),
	8604:  uint16(768),
	8605:  uint16(1),
	8606:  uint16(sym_identifier),
	8607:  uint16(770),
	8608:  uint16(1),
	8609:  uint16(sym_int_literal),
	8610:  uint16(222),
	8611:  uint16(1),
	8612:  uint16(aux_sym_attribute_repeat1),
	8613:  uint16(3),
	8614:  uint16(2),
	8615:  uint16(sym_block_comment),
	8616:  uint16(sym_line_comment),
	8617:  uint16(41),
	8618:  uint16(2),
	8619:  uint16(aux_sym_float_literal_token1),
	8620:  uint16(aux_sym_float_literal_token2),
	8621:  uint16(333),
	8622:  uint16(2),
	8623:  uint16(sym__literal_or_identifier),
	8624:  uint16(sym_float_literal),
	8625:  uint16(6),
	8626:  uint16(772),
	8627:  uint16(1),
	8628:  uint16(sym_identifier),
	8629:  uint16(774),
	8630:  uint16(1),
	8631:  uint16(sym_int_literal),
	8632:  uint16(227),
	8633:  uint16(1),
	8634:  uint16(aux_sym_attribute_repeat1),
	8635:  uint16(3),
	8636:  uint16(2),
	8637:  uint16(sym_block_comment),
	8638:  uint16(sym_line_comment),
	8639:  uint16(41),
	8640:  uint16(2),
	8641:  uint16(aux_sym_float_literal_token1),
	8642:  uint16(aux_sym_float_literal_token2),
	8643:  uint16(340),
	8644:  uint16(2),
	8645:  uint16(sym__literal_or_identifier),
	8646:  uint16(sym_float_literal),
	8647:  uint16(2),
	8648:  uint16(3),
	8649:  uint16(2),
	8650:  uint16(sym_block_comment),
	8651:  uint16(sym_line_comment),
	8652:  uint16(702),
	8653:  uint16(7),
	8654:  uint16(anon_sym_AT),
	8655:  uint16(sym_identifier),
	8656:  uint16(aux_sym_preproc_import_token1),
	8657:  uint16(aux_sym_preproc_ifdef_token1),
	8658:  uint16(aux_sym_preproc_ifdef_token2),
	8659:  uint16(aux_sym_preproc_ifdef_token3),
	8660:  uint16(aux_sym_preproc_else_token1),
	8661:  uint16(7),
	8662:  uint16(23),
	8663:  uint16(1),
	8664:  uint16(anon_sym_AT),
	8665:  uint16(674),
	8666:  uint16(1),
	8667:  uint16(sym_identifier),
	8668:  uint16(220),
	8669:  uint16(1),
	8670:  uint16(aux_sym_parameter_list_repeat1),
	8671:  uint16(328),
	8672:  uint16(1),
	8673:  uint16(sym_parameter),
	8674:  uint16(342),
	8675:  uint16(1),
	8676:  uint16(sym_variable_identifier_declaration),
	8677:  uint16(3),
	8678:  uint16(2),
	8679:  uint16(sym_block_comment),
	8680:  uint16(sym_line_comment),
	8681:  uint16(236),
	8682:  uint16(2),
	8683:  uint16(sym_attribute),
	8684:  uint16(aux_sym_global_variable_declaration_repeat1),
	8685:  uint16(4),
	8686:  uint16(778),
	8687:  uint16(1),
	8688:  uint16(anon_sym_COMMA),
	8689:  uint16(780),
	8690:  uint16(1),
	8691:  uint16(anon_sym_RBRACE),
	8692:  uint16(3),
	8693:  uint16(2),
	8694:  uint16(sym_block_comment),
	8695:  uint16(sym_line_comment),
	8696:  uint16(776),
	8697:  uint16(5),
	8698:  uint16(anon_sym_AT),
	8699:  uint16(sym_identifier),
	8700:  uint16(aux_sym_preproc_import_token1),
	8701:  uint16(aux_sym_preproc_ifdef_token1),
	8702:  uint16(aux_sym_preproc_ifdef_token2),
	8703:  uint16(7),
	8704:  uint16(23),
	8705:  uint16(1),
	8706:  uint16(anon_sym_AT),
	8707:  uint16(25),
	8708:  uint16(1),
	8709:  uint16(anon_sym_var),
	8710:  uint16(782),
	8711:  uint16(1),
	8712:  uint16(anon_sym_override),
	8713:  uint16(784),
	8714:  uint16(1),
	8715:  uint16(anon_sym_fn),
	8716:  uint16(307),
	8717:  uint16(1),
	8718:  uint16(sym_variable_declaration),
	8719:  uint16(3),
	8720:  uint16(2),
	8721:  uint16(sym_block_comment),
	8722:  uint16(sym_line_comment),
	8723:  uint16(49),
	8724:  uint16(2),
	8725:  uint16(sym_attribute),
	8726:  uint16(aux_sym_global_variable_declaration_repeat1),
	8727:  uint16(6),
	8728:  uint16(786),
	8729:  uint16(1),
	8730:  uint16(sym_identifier),
	8731:  uint16(789),
	8732:  uint16(1),
	8733:  uint16(sym_int_literal),
	8734:  uint16(227),
	8735:  uint16(1),
	8736:  uint16(aux_sym_attribute_repeat1),
	8737:  uint16(3),
	8738:  uint16(2),
	8739:  uint16(sym_block_comment),
	8740:  uint16(sym_line_comment),
	8741:  uint16(792),
	8742:  uint16(2),
	8743:  uint16(aux_sym_float_literal_token1),
	8744:  uint16(aux_sym_float_literal_token2),
	8745:  uint16(363),
	8746:  uint16(2),
	8747:  uint16(sym__literal_or_identifier),
	8748:  uint16(sym_float_literal),
	8749:  uint16(4),
	8750:  uint16(795),
	8751:  uint16(1),
	8752:  uint16(anon_sym_COMMA),
	8753:  uint16(797),
	8754:  uint16(1),
	8755:  uint16(anon_sym_RBRACE),
	8756:  uint16(3),
	8757:  uint16(2),
	8758:  uint16(sym_block_comment),
	8759:  uint16(sym_line_comment),
	8760:  uint16(776),
	8761:  uint16(5),
	8762:  uint16(anon_sym_AT),
	8763:  uint16(sym_identifier),
	8764:  uint16(aux_sym_preproc_import_token1),
	8765:  uint16(aux_sym_preproc_ifdef_token1),
	8766:  uint16(aux_sym_preproc_ifdef_token2),
	8767:  uint16(6),
	8768:  uint16(117),
	8769:  uint16(1),
	8770:  uint16(anon_sym_LPAREN),
	8771:  uint16(799),
	8772:  uint16(1),
	8773:  uint16(sym_identifier),
	8774:  uint16(239),
	8775:  uint16(1),
	8776:  uint16(aux_sym_lhs_expression_repeat1),
	8777:  uint16(386),
	8778:  uint16(1),
	8779:  uint16(sym_lhs_expression),
	8780:  uint16(3),
	8781:  uint16(2),
	8782:  uint16(sym_block_comment),
	8783:  uint16(sym_line_comment),
	8784:  uint16(123),
	8785:  uint16(2),
	8786:  uint16(anon_sym_AMP),
	8787:  uint16(anon_sym_STAR),
	8788:  uint16(3),
	8789:  uint16(803),
	8790:  uint16(1),
	8791:  uint16(anon_sym_RBRACE),
	8792:  uint16(3),
	8793:  uint16(2),
	8794:  uint16(sym_block_comment),
	8795:  uint16(sym_line_comment),
	8796:  uint16(801),
	8797:  uint16(5),
	8798:  uint16(anon_sym_AT),
	8799:  uint16(sym_identifier),
	8800:  uint16(aux_sym_preproc_import_token1),
	8801:  uint16(aux_sym_preproc_ifdef_token1),
	8802:  uint16(aux_sym_preproc_ifdef_token2),
	8803:  uint16(6),
	8804:  uint16(23),
	8805:  uint16(1),
	8806:  uint16(anon_sym_AT),
	8807:  uint16(805),
	8808:  uint16(1),
	8809:  uint16(sym_identifier),
	8810:  uint16(807),
	8811:  uint16(1),
	8812:  uint16(anon_sym_fn),
	8813:  uint16(329),
	8814:  uint16(1),
	8815:  uint16(sym_variable_identifier_declaration),
	8816:  uint16(3),
	8817:  uint16(2),
	8818:  uint16(sym_block_comment),
	8819:  uint16(sym_line_comment),
	8820:  uint16(258),
	8821:  uint16(2),
	8822:  uint16(sym_attribute),
	8823:  uint16(aux_sym_global_variable_declaration_repeat1),
	8824:  uint16(6),
	8825:  uint16(117),
	8826:  uint16(1),
	8827:  uint16(anon_sym_LPAREN),
	8828:  uint16(799),
	8829:  uint16(1),
	8830:  uint16(sym_identifier),
	8831:  uint16(239),
	8832:  uint16(1),
	8833:  uint16(aux_sym_lhs_expression_repeat1),
	8834:  uint16(348),
	8835:  uint16(1),
	8836:  uint16(sym_lhs_expression),
	8837:  uint16(3),
	8838:  uint16(2),
	8839:  uint16(sym_block_comment),
	8840:  uint16(sym_line_comment),
	8841:  uint16(123),
	8842:  uint16(2),
	8843:  uint16(anon_sym_AMP),
	8844:  uint16(anon_sym_STAR),
	8845:  uint16(3),
	8846:  uint16(357),
	8847:  uint16(1),
	8848:  uint16(sym_address_space),
	8849:  uint16(3),
	8850:  uint16(2),
	8851:  uint16(sym_block_comment),
	8852:  uint16(sym_line_comment),
	8853:  uint16(809),
	8854:  uint16(5),
	8855:  uint16(anon_sym_function),
	8856:  uint16(anon_sym_private),
	8857:  uint16(anon_sym_workgroup),
	8858:  uint16(anon_sym_uniform),
	8859:  uint16(anon_sym_storage),
	8860:  uint16(3),
	8861:  uint16(797),
	8862:  uint16(1),
	8863:  uint16(anon_sym_RBRACE),
	8864:  uint16(3),
	8865:  uint16(2),
	8866:  uint16(sym_block_comment),
	8867:  uint16(sym_line_comment),
	8868:  uint16(801),
	8869:  uint16(5),
	8870:  uint16(anon_sym_AT),
	8871:  uint16(sym_identifier),
	8872:  uint16(aux_sym_preproc_import_token1),
	8873:  uint16(aux_sym_preproc_ifdef_token1),
	8874:  uint16(aux_sym_preproc_ifdef_token2),
	8875:  uint16(3),
	8876:  uint16(317),
	8877:  uint16(1),
	8878:  uint16(sym_address_space),
	8879:  uint16(3),
	8880:  uint16(2),
	8881:  uint16(sym_block_comment),
	8882:  uint16(sym_line_comment),
	8883:  uint16(809),
	8884:  uint16(5),
	8885:  uint16(anon_sym_function),
	8886:  uint16(anon_sym_private),
	8887:  uint16(anon_sym_workgroup),
	8888:  uint16(anon_sym_uniform),
	8889:  uint16(anon_sym_storage),
	8890:  uint16(5),
	8891:  uint16(23),
	8892:  uint16(1),
	8893:  uint16(anon_sym_AT),
	8894:  uint16(674),
	8895:  uint16(1),
	8896:  uint16(sym_identifier),
	8897:  uint16(327),
	8898:  uint16(1),
	8899:  uint16(sym_variable_identifier_declaration),
	8900:  uint16(3),
	8901:  uint16(2),
	8902:  uint16(sym_block_comment),
	8903:  uint16(sym_line_comment),
	8904:  uint16(49),
	8905:  uint16(2),
	8906:  uint16(sym_attribute),
	8907:  uint16(aux_sym_global_variable_declaration_repeat1),
	8908:  uint16(5),
	8909:  uint16(23),
	8910:  uint16(1),
	8911:  uint16(anon_sym_AT),
	8912:  uint16(674),
	8913:  uint16(1),
	8914:  uint16(sym_identifier),
	8915:  uint16(210),
	8916:  uint16(1),
	8917:  uint16(sym_variable_identifier_declaration),
	8918:  uint16(3),
	8919:  uint16(2),
	8920:  uint16(sym_block_comment),
	8921:  uint16(sym_line_comment),
	8922:  uint16(49),
	8923:  uint16(2),
	8924:  uint16(sym_attribute),
	8925:  uint16(aux_sym_global_variable_declaration_repeat1),
	8926:  uint16(6),
	8927:  uint16(3),
	8928:  uint16(1),
	8929:  uint16(sym_block_comment),
	8930:  uint16(813),
	8931:  uint16(1),
	8932:  uint16(sym_line_comment),
	8933:  uint16(815),
	8934:  uint16(1),
	8935:  uint16(anon_sym_LF),
	8936:  uint16(817),
	8937:  uint16(1),
	8938:  uint16(anon_sym_COLON_COLON),
	8939:  uint16(238),
	8940:  uint16(1),
	8941:  uint16(aux_sym_import_path_repeat1),
	8942:  uint16(811),
	8943:  uint16(2),
	8944:  uint16(sym_identifier),
	8945:  uint16(anon_sym_as),
	8946:  uint16(5),
	8947:  uint16(820),
	8948:  uint16(1),
	8949:  uint16(sym_identifier),
	8950:  uint16(822),
	8951:  uint16(1),
	8952:  uint16(anon_sym_LPAREN),
	8953:  uint16(242),
	8954:  uint16(1),
	8955:  uint16(aux_sym_lhs_expression_repeat1),
	8956:  uint16(3),
	8957:  uint16(2),
	8958:  uint16(sym_block_comment),
	8959:  uint16(sym_line_comment),
	8960:  uint16(824),
	8961:  uint16(2),
	8962:  uint16(anon_sym_AMP),
	8963:  uint16(anon_sym_STAR),
	8964:  uint16(4),
	8965:  uint16(828),
	8966:  uint16(1),
	8967:  uint16(anon_sym_RPAREN),
	8968:  uint16(830),
	8969:  uint16(1),
	8970:  uint16(sym_int_literal),
	8971:  uint16(3),
	8972:  uint16(2),
	8973:  uint16(sym_block_comment),
	8974:  uint16(sym_line_comment),
	8975:  uint16(826),
	8976:  uint16(3),
	8977:  uint16(sym_identifier),
	8978:  uint16(aux_sym_float_literal_token1),
	8979:  uint16(aux_sym_float_literal_token2),
	8980:  uint16(6),
	8981:  uint16(3),
	8982:  uint16(1),
	8983:  uint16(sym_block_comment),
	8984:  uint16(813),
	8985:  uint16(1),
	8986:  uint16(sym_line_comment),
	8987:  uint16(834),
	8988:  uint16(1),
	8989:  uint16(anon_sym_LF),
	8990:  uint16(836),
	8991:  uint16(1),
	8992:  uint16(anon_sym_COLON_COLON),
	8993:  uint16(238),
	8994:  uint16(1),
	8995:  uint16(aux_sym_import_path_repeat1),
	8996:  uint16(832),
	8997:  uint16(2),
	8998:  uint16(sym_identifier),
	8999:  uint16(anon_sym_as),
	9000:  uint16(4),
	9001:  uint16(242),
	9002:  uint16(1),
	9003:  uint16(aux_sym_lhs_expression_repeat1),
	9004:  uint16(3),
	9005:  uint16(2),
	9006:  uint16(sym_block_comment),
	9007:  uint16(sym_line_comment),
	9008:  uint16(838),
	9009:  uint16(2),
	9010:  uint16(anon_sym_LPAREN),
	9011:  uint16(sym_identifier),
	9012:  uint16(840),
	9013:  uint16(2),
	9014:  uint16(anon_sym_AMP),
	9015:  uint16(anon_sym_STAR),
	9016:  uint16(2),
	9017:  uint16(3),
	9018:  uint16(2),
	9019:  uint16(sym_block_comment),
	9020:  uint16(sym_line_comment),
	9021:  uint16(801),
	9022:  uint16(5),
	9023:  uint16(anon_sym_AT),
	9024:  uint16(sym_identifier),
	9025:  uint16(aux_sym_preproc_import_token1),
	9026:  uint16(aux_sym_preproc_ifdef_token1),
	9027:  uint16(aux_sym_preproc_ifdef_token2),
	9028:  uint16(5),
	9029:  uint16(234),
	9030:  uint16(1),
	9031:  uint16(anon_sym_LBRACE),
	9032:  uint16(843),
	9033:  uint16(1),
	9034:  uint16(anon_sym_if),
	9035:  uint16(125),
	9036:  uint16(1),
	9037:  uint16(sym_else_statement),
	9038:  uint16(3),
	9039:  uint16(2),
	9040:  uint16(sym_block_comment),
	9041:  uint16(sym_line_comment),
	9042:  uint16(124),
	9043:  uint16(2),
	9044:  uint16(sym_compound_statement),
	9045:  uint16(sym_if_statement),
	9046:  uint16(5),
	9047:  uint16(845),
	9048:  uint16(1),
	9049:  uint16(anon_sym_RBRACE),
	9050:  uint16(847),
	9051:  uint16(1),
	9052:  uint16(anon_sym_case),
	9053:  uint16(849),
	9054:  uint16(1),
	9055:  uint16(anon_sym_default),
	9056:  uint16(3),
	9057:  uint16(2),
	9058:  uint16(sym_block_comment),
	9059:  uint16(sym_line_comment),
	9060:  uint16(246),
	9061:  uint16(2),
	9062:  uint16(sym_switch_body),
	9063:  uint16(aux_sym_switch_statement_repeat1),
	9064:  uint16(5),
	9065:  uint16(851),
	9066:  uint16(1),
	9067:  uint16(anon_sym_RBRACE),
	9068:  uint16(853),
	9069:  uint16(1),
	9070:  uint16(anon_sym_case),
	9071:  uint16(856),
	9072:  uint16(1),
	9073:  uint16(anon_sym_default),
	9074:  uint16(3),
	9075:  uint16(2),
	9076:  uint16(sym_block_comment),
	9077:  uint16(sym_line_comment),
	9078:  uint16(246),
	9079:  uint16(2),
	9080:  uint16(sym_switch_body),
	9081:  uint16(aux_sym_switch_statement_repeat1),
	9082:  uint16(4),
	9083:  uint16(830),
	9084:  uint16(1),
	9085:  uint16(sym_int_literal),
	9086:  uint16(859),
	9087:  uint16(1),
	9088:  uint16(anon_sym_RPAREN),
	9089:  uint16(3),
	9090:  uint16(2),
	9091:  uint16(sym_block_comment),
	9092:  uint16(sym_line_comment),
	9093:  uint16(826),
	9094:  uint16(3),
	9095:  uint16(sym_identifier),
	9096:  uint16(aux_sym_float_literal_token1),
	9097:  uint16(aux_sym_float_literal_token2),
	9098:  uint16(6),
	9099:  uint16(3),
	9100:  uint16(1),
	9101:  uint16(sym_block_comment),
	9102:  uint16(813),
	9103:  uint16(1),
	9104:  uint16(sym_line_comment),
	9105:  uint16(836),
	9106:  uint16(1),
	9107:  uint16(anon_sym_COLON_COLON),
	9108:  uint16(863),
	9109:  uint16(1),
	9110:  uint16(anon_sym_LF),
	9111:  uint16(241),
	9112:  uint16(1),
	9113:  uint16(aux_sym_import_path_repeat1),
	9114:  uint16(861),
	9115:  uint16(2),
	9116:  uint16(sym_identifier),
	9117:  uint16(anon_sym_as),
	9118:  uint16(5),
	9119:  uint16(234),
	9120:  uint16(1),
	9121:  uint16(anon_sym_LBRACE),
	9122:  uint16(865),
	9123:  uint16(1),
	9124:  uint16(anon_sym_DASH_GT),
	9125:  uint16(168),
	9126:  uint16(1),
	9127:  uint16(sym_compound_statement),
	9128:  uint16(316),
	9129:  uint16(1),
	9130:  uint16(sym_function_return_type_declaration),
	9131:  uint16(3),
	9132:  uint16(2),
	9133:  uint16(sym_block_comment),
	9134:  uint16(sym_line_comment),
	9135:  uint16(5),
	9136:  uint16(234),
	9137:  uint16(1),
	9138:  uint16(anon_sym_LBRACE),
	9139:  uint16(865),
	9140:  uint16(1),
	9141:  uint16(anon_sym_DASH_GT),
	9142:  uint16(157),
	9143:  uint16(1),
	9144:  uint16(sym_compound_statement),
	9145:  uint16(325),
	9146:  uint16(1),
	9147:  uint16(sym_function_return_type_declaration),
	9148:  uint16(3),
	9149:  uint16(2),
	9150:  uint16(sym_block_comment),
	9151:  uint16(sym_line_comment),
	9152:  uint16(2),
	9153:  uint16(3),
	9154:  uint16(2),
	9155:  uint16(sym_block_comment),
	9156:  uint16(sym_line_comment),
	9157:  uint16(867),
	9158:  uint16(4),
	9159:  uint16(anon_sym_u32),
	9160:  uint16(anon_sym_i32),
	9161:  uint16(anon_sym_f32),
	9162:  uint16(anon_sym_f16),
	9163:  uint16(5),
	9164:  uint16(234),
	9165:  uint16(1),
	9166:  uint16(anon_sym_LBRACE),
	9167:  uint16(865),
	9168:  uint16(1),
	9169:  uint16(anon_sym_DASH_GT),
	9170:  uint16(159),
	9171:  uint16(1),
	9172:  uint16(sym_compound_statement),
	9173:  uint16(306),
	9174:  uint16(1),
	9175:  uint16(sym_function_return_type_declaration),
	9176:  uint16(3),
	9177:  uint16(2),
	9178:  uint16(sym_block_comment),
	9179:  uint16(sym_line_comment),
	9180:  uint16(6),
	9181:  uint16(3),
	9182:  uint16(1),
	9183:  uint16(sym_block_comment),
	9184:  uint16(813),
	9185:  uint16(1),
	9186:  uint16(sym_line_comment),
	9187:  uint16(869),
	9188:  uint16(1),
	9189:  uint16(sym_identifier),
	9190:  uint16(871),
	9191:  uint16(1),
	9192:  uint16(anon_sym_LF),
	9193:  uint16(873),
	9194:  uint16(1),
	9195:  uint16(anon_sym_as),
	9196:  uint16(390),
	9197:  uint16(1),
	9198:  uint16(sym_alias),
	9199:  uint16(4),
	9200:  uint16(875),
	9201:  uint16(1),
	9202:  uint16(anon_sym_read),
	9203:  uint16(349),
	9204:  uint16(1),
	9205:  uint16(sym_access_mode),
	9206:  uint16(3),
	9207:  uint16(2),
	9208:  uint16(sym_block_comment),
	9209:  uint16(sym_line_comment),
	9210:  uint16(877),
	9211:  uint16(2),
	9212:  uint16(anon_sym_write),
	9213:  uint16(anon_sym_read_write),
	9214:  uint16(3),
	9215:  uint16(830),
	9216:  uint16(1),
	9217:  uint16(sym_int_literal),
	9218:  uint16(3),
	9219:  uint16(2),
	9220:  uint16(sym_block_comment),
	9221:  uint16(sym_line_comment),
	9222:  uint16(826),
	9223:  uint16(3),
	9224:  uint16(sym_identifier),
	9225:  uint16(aux_sym_float_literal_token1),
	9226:  uint16(aux_sym_float_literal_token2),
	9227:  uint16(5),
	9228:  uint16(234),
	9229:  uint16(1),
	9230:  uint16(anon_sym_LBRACE),
	9231:  uint16(865),
	9232:  uint16(1),
	9233:  uint16(anon_sym_DASH_GT),
	9234:  uint16(169),
	9235:  uint16(1),
	9236:  uint16(sym_compound_statement),
	9237:  uint16(298),
	9238:  uint16(1),
	9239:  uint16(sym_function_return_type_declaration),
	9240:  uint16(3),
	9241:  uint16(2),
	9242:  uint16(sym_block_comment),
	9243:  uint16(sym_line_comment),
	9244:  uint16(4),
	9245:  uint16(875),
	9246:  uint16(1),
	9247:  uint16(anon_sym_read),
	9248:  uint16(385),
	9249:  uint16(1),
	9250:  uint16(sym_access_mode),
	9251:  uint16(3),
	9252:  uint16(2),
	9253:  uint16(sym_block_comment),
	9254:  uint16(sym_line_comment),
	9255:  uint16(877),
	9256:  uint16(2),
	9257:  uint16(anon_sym_write),
	9258:  uint16(anon_sym_read_write),
	9259:  uint16(4),
	9260:  uint16(23),
	9261:  uint16(1),
	9262:  uint16(anon_sym_AT),
	9263:  uint16(879),
	9264:  uint16(1),
	9265:  uint16(anon_sym_fn),
	9266:  uint16(3),
	9267:  uint16(2),
	9268:  uint16(sym_block_comment),
	9269:  uint16(sym_line_comment),
	9270:  uint16(49),
	9271:  uint16(2),
	9272:  uint16(sym_attribute),
	9273:  uint16(aux_sym_global_variable_declaration_repeat1),
	9274:  uint16(4),
	9275:  uint16(847),
	9276:  uint16(1),
	9277:  uint16(anon_sym_case),
	9278:  uint16(849),
	9279:  uint16(1),
	9280:  uint16(anon_sym_default),
	9281:  uint16(3),
	9282:  uint16(2),
	9283:  uint16(sym_block_comment),
	9284:  uint16(sym_line_comment),
	9285:  uint16(245),
	9286:  uint16(2),
	9287:  uint16(sym_switch_body),
	9288:  uint16(aux_sym_switch_statement_repeat1),
	9289:  uint16(4),
	9290:  uint16(23),
	9291:  uint16(1),
	9292:  uint16(anon_sym_AT),
	9293:  uint16(881),
	9294:  uint16(1),
	9295:  uint16(anon_sym_fn),
	9296:  uint16(3),
	9297:  uint16(2),
	9298:  uint16(sym_block_comment),
	9299:  uint16(sym_line_comment),
	9300:  uint16(49),
	9301:  uint16(2),
	9302:  uint16(sym_attribute),
	9303:  uint16(aux_sym_global_variable_declaration_repeat1),
	9304:  uint16(4),
	9305:  uint16(3),
	9306:  uint16(1),
	9307:  uint16(sym_block_comment),
	9308:  uint16(813),
	9309:  uint16(1),
	9310:  uint16(sym_line_comment),
	9311:  uint16(815),
	9312:  uint16(1),
	9313:  uint16(anon_sym_LF),
	9314:  uint16(811),
	9315:  uint16(3),
	9316:  uint16(sym_identifier),
	9317:  uint16(anon_sym_COLON_COLON),
	9318:  uint16(anon_sym_as),
	9319:  uint16(4),
	9320:  uint16(875),
	9321:  uint16(1),
	9322:  uint16(anon_sym_read),
	9323:  uint16(379),
	9324:  uint16(1),
	9325:  uint16(sym_access_mode),
	9326:  uint16(3),
	9327:  uint16(2),
	9328:  uint16(sym_block_comment),
	9329:  uint16(sym_line_comment),
	9330:  uint16(877),
	9331:  uint16(2),
	9332:  uint16(anon_sym_write),
	9333:  uint16(anon_sym_read_write),
	9334:  uint16(5),
	9335:  uint16(234),
	9336:  uint16(1),
	9337:  uint16(anon_sym_LBRACE),
	9338:  uint16(865),
	9339:  uint16(1),
	9340:  uint16(anon_sym_DASH_GT),
	9341:  uint16(161),
	9342:  uint16(1),
	9343:  uint16(sym_compound_statement),
	9344:  uint16(320),
	9345:  uint16(1),
	9346:  uint16(sym_function_return_type_declaration),
	9347:  uint16(3),
	9348:  uint16(2),
	9349:  uint16(sym_block_comment),
	9350:  uint16(sym_line_comment),
	9351:  uint16(4),
	9352:  uint16(883),
	9353:  uint16(1),
	9354:  uint16(anon_sym_COMMA),
	9355:  uint16(268),
	9356:  uint16(1),
	9357:  uint16(aux_sym_case_selectors_repeat1),
	9358:  uint16(3),
	9359:  uint16(2),
	9360:  uint16(sym_block_comment),
	9361:  uint16(sym_line_comment),
	9362:  uint16(885),
	9363:  uint16(2),
	9364:  uint16(anon_sym_LBRACE),
	9365:  uint16(anon_sym_COLON),
	9366:  uint16(4),
	9367:  uint16(23),
	9368:  uint16(1),
	9369:  uint16(anon_sym_AT),
	9370:  uint16(784),
	9371:  uint16(1),
	9372:  uint16(anon_sym_fn),
	9373:  uint16(3),
	9374:  uint16(2),
	9375:  uint16(sym_block_comment),
	9376:  uint16(sym_line_comment),
	9377:  uint16(260),
	9378:  uint16(2),
	9379:  uint16(sym_attribute),
	9380:  uint16(aux_sym_global_variable_declaration_repeat1),
	9381:  uint16(5),
	9382:  uint16(234),
	9383:  uint16(1),
	9384:  uint16(anon_sym_LBRACE),
	9385:  uint16(865),
	9386:  uint16(1),
	9387:  uint16(anon_sym_DASH_GT),
	9388:  uint16(165),
	9389:  uint16(1),
	9390:  uint16(sym_compound_statement),
	9391:  uint16(310),
	9392:  uint16(1),
	9393:  uint16(sym_function_return_type_declaration),
	9394:  uint16(3),
	9395:  uint16(2),
	9396:  uint16(sym_block_comment),
	9397:  uint16(sym_line_comment),
	9398:  uint16(5),
	9399:  uint16(887),
	9400:  uint16(1),
	9401:  uint16(sym_identifier),
	9402:  uint16(889),
	9403:  uint16(1),
	9404:  uint16(anon_sym_LT),
	9405:  uint16(321),
	9406:  uint16(1),
	9407:  uint16(sym_variable_qualifier),
	9408:  uint16(322),
	9409:  uint16(1),
	9410:  uint16(sym_variable_identifier_declaration),
	9411:  uint16(3),
	9412:  uint16(2),
	9413:  uint16(sym_block_comment),
	9414:  uint16(sym_line_comment),
	9415:  uint16(4),
	9416:  uint16(891),
	9417:  uint16(1),
	9418:  uint16(anon_sym_COMMA),
	9419:  uint16(269),
	9420:  uint16(1),
	9421:  uint16(aux_sym_case_selectors_repeat1),
	9422:  uint16(3),
	9423:  uint16(2),
	9424:  uint16(sym_block_comment),
	9425:  uint16(sym_line_comment),
	9426:  uint16(734),
	9427:  uint16(2),
	9428:  uint16(anon_sym_LBRACE),
	9429:  uint16(anon_sym_COLON),
	9430:  uint16(4),
	9431:  uint16(893),
	9432:  uint16(1),
	9433:  uint16(anon_sym_COMMA),
	9434:  uint16(269),
	9435:  uint16(1),
	9436:  uint16(aux_sym_case_selectors_repeat1),
	9437:  uint16(3),
	9438:  uint16(2),
	9439:  uint16(sym_block_comment),
	9440:  uint16(sym_line_comment),
	9441:  uint16(896),
	9442:  uint16(2),
	9443:  uint16(anon_sym_LBRACE),
	9444:  uint16(anon_sym_COLON),
	9445:  uint16(2),
	9446:  uint16(3),
	9447:  uint16(2),
	9448:  uint16(sym_block_comment),
	9449:  uint16(sym_line_comment),
	9450:  uint16(898),
	9451:  uint16(3),
	9452:  uint16(anon_sym_RBRACE),
	9453:  uint16(anon_sym_case),
	9454:  uint16(anon_sym_default),
	9455:  uint16(2),
	9456:  uint16(3),
	9457:  uint16(2),
	9458:  uint16(sym_block_comment),
	9459:  uint16(sym_line_comment),
	9460:  uint16(900),
	9461:  uint16(3),
	9462:  uint16(anon_sym_SEMI),
	9463:  uint16(anon_sym_COMMA),
	9464:  uint16(anon_sym_RPAREN),
	9465:  uint16(2),
	9466:  uint16(3),
	9467:  uint16(2),
	9468:  uint16(sym_block_comment),
	9469:  uint16(sym_line_comment),
	9470:  uint16(902),
	9471:  uint16(3),
	9472:  uint16(anon_sym_SEMI),
	9473:  uint16(anon_sym_COMMA),
	9474:  uint16(anon_sym_RPAREN),
	9475:  uint16(4),
	9476:  uint16(904),
	9477:  uint16(1),
	9478:  uint16(anon_sym_LBRACE),
	9479:  uint16(906),
	9480:  uint16(1),
	9481:  uint16(anon_sym_COLON),
	9482:  uint16(280),
	9483:  uint16(1),
	9484:  uint16(sym_case_compound_statement),
	9485:  uint16(3),
	9486:  uint16(2),
	9487:  uint16(sym_block_comment),
	9488:  uint16(sym_line_comment),
	9489:  uint16(2),
	9490:  uint16(3),
	9491:  uint16(2),
	9492:  uint16(sym_block_comment),
	9493:  uint16(sym_line_comment),
	9494:  uint16(908),
	9495:  uint16(3),
	9496:  uint16(anon_sym_SEMI),
	9497:  uint16(anon_sym_COMMA),
	9498:  uint16(anon_sym_RPAREN),
	9499:  uint16(2),
	9500:  uint16(3),
	9501:  uint16(2),
	9502:  uint16(sym_block_comment),
	9503:  uint16(sym_line_comment),
	9504:  uint16(910),
	9505:  uint16(3),
	9506:  uint16(anon_sym_SEMI),
	9507:  uint16(anon_sym_COMMA),
	9508:  uint16(anon_sym_RPAREN),
	9509:  uint16(4),
	9510:  uint16(912),
	9511:  uint16(1),
	9512:  uint16(anon_sym_SEMI),
	9513:  uint16(914),
	9514:  uint16(1),
	9515:  uint16(anon_sym_EQ),
	9516:  uint16(916),
	9517:  uint16(1),
	9518:  uint16(anon_sym_COLON),
	9519:  uint16(3),
	9520:  uint16(2),
	9521:  uint16(sym_block_comment),
	9522:  uint16(sym_line_comment),
	9523:  uint16(4),
	9524:  uint16(904),
	9525:  uint16(1),
	9526:  uint16(anon_sym_LBRACE),
	9527:  uint16(918),
	9528:  uint16(1),
	9529:  uint16(anon_sym_COLON),
	9530:  uint16(270),
	9531:  uint16(1),
	9532:  uint16(sym_case_compound_statement),
	9533:  uint16(3),
	9534:  uint16(2),
	9535:  uint16(sym_block_comment),
	9536:  uint16(sym_line_comment),
	9537:  uint16(4),
	9538:  uint16(916),
	9539:  uint16(1),
	9540:  uint16(anon_sym_COLON),
	9541:  uint16(920),
	9542:  uint16(1),
	9543:  uint16(anon_sym_SEMI),
	9544:  uint16(922),
	9545:  uint16(1),
	9546:  uint16(anon_sym_EQ),
	9547:  uint16(3),
	9548:  uint16(2),
	9549:  uint16(sym_block_comment),
	9550:  uint16(sym_line_comment),
	9551:  uint16(3),
	9552:  uint16(916),
	9553:  uint16(1),
	9554:  uint16(anon_sym_COLON),
	9555:  uint16(3),
	9556:  uint16(2),
	9557:  uint16(sym_block_comment),
	9558:  uint16(sym_line_comment),
	9559:  uint16(924),
	9560:  uint16(2),
	9561:  uint16(anon_sym_SEMI),
	9562:  uint16(anon_sym_EQ),
	9563:  uint16(2),
	9564:  uint16(3),
	9565:  uint16(2),
	9566:  uint16(sym_block_comment),
	9567:  uint16(sym_line_comment),
	9568:  uint16(926),
	9569:  uint16(3),
	9570:  uint16(anon_sym_RBRACE),
	9571:  uint16(anon_sym_case),
	9572:  uint16(anon_sym_default),
	9573:  uint16(3),
	9574:  uint16(916),
	9575:  uint16(1),
	9576:  uint16(anon_sym_COLON),
	9577:  uint16(3),
	9578:  uint16(2),
	9579:  uint16(sym_block_comment),
	9580:  uint16(sym_line_comment),
	9581:  uint16(928),
	9582:  uint16(2),
	9583:  uint16(anon_sym_SEMI),
	9584:  uint16(anon_sym_EQ),
	9585:  uint16(3),
	9586:  uint16(932),
	9587:  uint16(1),
	9588:  uint16(anon_sym_RPAREN),
	9589:  uint16(3),
	9590:  uint16(2),
	9591:  uint16(sym_block_comment),
	9592:  uint16(sym_line_comment),
	9593:  uint16(930),
	9594:  uint16(2),
	9595:  uint16(anon_sym_AT),
	9596:  uint16(sym_identifier),
	9597:  uint16(4),
	9598:  uint16(934),
	9599:  uint16(1),
	9600:  uint16(anon_sym_LPAREN),
	9601:  uint16(936),
	9602:  uint16(1),
	9603:  uint16(anon_sym_LT),
	9604:  uint16(96),
	9605:  uint16(1),
	9606:  uint16(sym_argument_list_expression),
	9607:  uint16(3),
	9608:  uint16(2),
	9609:  uint16(sym_block_comment),
	9610:  uint16(sym_line_comment),
	9611:  uint16(3),
	9612:  uint16(938),
	9613:  uint16(1),
	9614:  uint16(anon_sym_RPAREN),
	9615:  uint16(3),
	9616:  uint16(2),
	9617:  uint16(sym_block_comment),
	9618:  uint16(sym_line_comment),
	9619:  uint16(930),
	9620:  uint16(2),
	9621:  uint16(anon_sym_AT),
	9622:  uint16(sym_identifier),
	9623:  uint16(5),
	9624:  uint16(3),
	9625:  uint16(1),
	9626:  uint16(sym_block_comment),
	9627:  uint16(813),
	9628:  uint16(1),
	9629:  uint16(sym_line_comment),
	9630:  uint16(940),
	9631:  uint16(1),
	9632:  uint16(anon_sym_COMMA),
	9633:  uint16(942),
	9634:  uint16(1),
	9635:  uint16(anon_sym_LF),
	9636:  uint16(296),
	9637:  uint16(1),
	9638:  uint16(aux_sym_preproc_import_repeat1),
	9639:  uint16(2),
	9640:  uint16(3),
	9641:  uint16(2),
	9642:  uint16(sym_block_comment),
	9643:  uint16(sym_line_comment),
	9644:  uint16(944),
	9645:  uint16(3),
	9646:  uint16(anon_sym_SEMI),
	9647:  uint16(anon_sym_COMMA),
	9648:  uint16(anon_sym_RPAREN),
	9649:  uint16(2),
	9650:  uint16(3),
	9651:  uint16(2),
	9652:  uint16(sym_block_comment),
	9653:  uint16(sym_line_comment),
	9654:  uint16(946),
	9655:  uint16(3),
	9656:  uint16(anon_sym_RBRACE),
	9657:  uint16(anon_sym_case),
	9658:  uint16(anon_sym_default),
	9659:  uint16(2),
	9660:  uint16(3),
	9661:  uint16(2),
	9662:  uint16(sym_block_comment),
	9663:  uint16(sym_line_comment),
	9664:  uint16(948),
	9665:  uint16(3),
	9666:  uint16(anon_sym_RBRACE),
	9667:  uint16(anon_sym_case),
	9668:  uint16(anon_sym_default),
	9669:  uint16(2),
	9670:  uint16(3),
	9671:  uint16(2),
	9672:  uint16(sym_block_comment),
	9673:  uint16(sym_line_comment),
	9674:  uint16(896),
	9675:  uint16(3),
	9676:  uint16(anon_sym_COMMA),
	9677:  uint16(anon_sym_LBRACE),
	9678:  uint16(anon_sym_COLON),
	9679:  uint16(5),
	9680:  uint16(3),
	9681:  uint16(1),
	9682:  uint16(sym_block_comment),
	9683:  uint16(813),
	9684:  uint16(1),
	9685:  uint16(sym_line_comment),
	9686:  uint16(950),
	9687:  uint16(1),
	9688:  uint16(anon_sym_COMMA),
	9689:  uint16(953),
	9690:  uint16(1),
	9691:  uint16(anon_sym_LF),
	9692:  uint16(290),
	9693:  uint16(1),
	9694:  uint16(aux_sym_preproc_import_repeat1),
	9695:  uint16(2),
	9696:  uint16(3),
	9697:  uint16(2),
	9698:  uint16(sym_block_comment),
	9699:  uint16(sym_line_comment),
	9700:  uint16(955),
	9701:  uint16(3),
	9702:  uint16(anon_sym_RBRACE),
	9703:  uint16(anon_sym_case),
	9704:  uint16(anon_sym_default),
	9705:  uint16(2),
	9706:  uint16(3),
	9707:  uint16(2),
	9708:  uint16(sym_block_comment),
	9709:  uint16(sym_line_comment),
	9710:  uint16(957),
	9711:  uint16(3),
	9712:  uint16(anon_sym_RBRACE),
	9713:  uint16(anon_sym_case),
	9714:  uint16(anon_sym_default),
	9715:  uint16(4),
	9716:  uint16(863),
	9717:  uint16(1),
	9718:  uint16(anon_sym_LPAREN),
	9719:  uint16(959),
	9720:  uint16(1),
	9721:  uint16(anon_sym_COLON_COLON),
	9722:  uint16(294),
	9723:  uint16(1),
	9724:  uint16(aux_sym_import_path_repeat1),
	9725:  uint16(3),
	9726:  uint16(2),
	9727:  uint16(sym_block_comment),
	9728:  uint16(sym_line_comment),
	9729:  uint16(4),
	9730:  uint16(834),
	9731:  uint16(1),
	9732:  uint16(anon_sym_LPAREN),
	9733:  uint16(959),
	9734:  uint16(1),
	9735:  uint16(anon_sym_COLON_COLON),
	9736:  uint16(295),
	9737:  uint16(1),
	9738:  uint16(aux_sym_import_path_repeat1),
	9739:  uint16(3),
	9740:  uint16(2),
	9741:  uint16(sym_block_comment),
	9742:  uint16(sym_line_comment),
	9743:  uint16(4),
	9744:  uint16(815),
	9745:  uint16(1),
	9746:  uint16(anon_sym_LPAREN),
	9747:  uint16(961),
	9748:  uint16(1),
	9749:  uint16(anon_sym_COLON_COLON),
	9750:  uint16(295),
	9751:  uint16(1),
	9752:  uint16(aux_sym_import_path_repeat1),
	9753:  uint16(3),
	9754:  uint16(2),
	9755:  uint16(sym_block_comment),
	9756:  uint16(sym_line_comment),
	9757:  uint16(5),
	9758:  uint16(3),
	9759:  uint16(1),
	9760:  uint16(sym_block_comment),
	9761:  uint16(813),
	9762:  uint16(1),
	9763:  uint16(sym_line_comment),
	9764:  uint16(940),
	9765:  uint16(1),
	9766:  uint16(anon_sym_COMMA),
	9767:  uint16(964),
	9768:  uint16(1),
	9769:  uint16(anon_sym_LF),
	9770:  uint16(290),
	9771:  uint16(1),
	9772:  uint16(aux_sym_preproc_import_repeat1),
	9773:  uint16(3),
	9774:  uint16(966),
	9775:  uint16(1),
	9776:  uint16(sym_identifier),
	9777:  uint16(384),
	9778:  uint16(1),
	9779:  uint16(sym_variable_identifier_declaration),
	9780:  uint16(3),
	9781:  uint16(2),
	9782:  uint16(sym_block_comment),
	9783:  uint16(sym_line_comment),
	9784:  uint16(3),
	9785:  uint16(234),
	9786:  uint16(1),
	9787:  uint16(anon_sym_LBRACE),
	9788:  uint16(154),
	9789:  uint16(1),
	9790:  uint16(sym_compound_statement),
	9791:  uint16(3),
	9792:  uint16(2),
	9793:  uint16(sym_block_comment),
	9794:  uint16(sym_line_comment),
	9795:  uint16(3),
	9796:  uint16(916),
	9797:  uint16(1),
	9798:  uint16(anon_sym_COLON),
	9799:  uint16(968),
	9800:  uint16(1),
	9801:  uint16(anon_sym_EQ),
	9802:  uint16(3),
	9803:  uint16(2),
	9804:  uint16(sym_block_comment),
	9805:  uint16(sym_line_comment),
	9806:  uint16(3),
	9807:  uint16(970),
	9808:  uint16(1),
	9809:  uint16(anon_sym_COMMA),
	9810:  uint16(972),
	9811:  uint16(1),
	9812:  uint16(anon_sym_GT),
	9813:  uint16(3),
	9814:  uint16(2),
	9815:  uint16(sym_block_comment),
	9816:  uint16(sym_line_comment),
	9817:  uint16(2),
	9818:  uint16(3),
	9819:  uint16(2),
	9820:  uint16(sym_block_comment),
	9821:  uint16(sym_line_comment),
	9822:  uint16(930),
	9823:  uint16(2),
	9824:  uint16(anon_sym_AT),
	9825:  uint16(sym_identifier),
	9826:  uint16(2),
	9827:  uint16(3),
	9828:  uint16(2),
	9829:  uint16(sym_block_comment),
	9830:  uint16(sym_line_comment),
	9831:  uint16(924),
	9832:  uint16(2),
	9833:  uint16(anon_sym_SEMI),
	9834:  uint16(anon_sym_EQ),
	9835:  uint16(3),
	9836:  uint16(974),
	9837:  uint16(1),
	9838:  uint16(anon_sym_COMMA),
	9839:  uint16(976),
	9840:  uint16(1),
	9841:  uint16(anon_sym_RPAREN),
	9842:  uint16(3),
	9843:  uint16(2),
	9844:  uint16(sym_block_comment),
	9845:  uint16(sym_line_comment),
	9846:  uint16(2),
	9847:  uint16(3),
	9848:  uint16(2),
	9849:  uint16(sym_block_comment),
	9850:  uint16(sym_line_comment),
	9851:  uint16(978),
	9852:  uint16(2),
	9853:  uint16(sym_identifier),
	9854:  uint16(sym_int_literal),
	9855:  uint16(3),
	9856:  uint16(267),
	9857:  uint16(1),
	9858:  uint16(anon_sym_RPAREN),
	9859:  uint16(980),
	9860:  uint16(1),
	9861:  uint16(anon_sym_COMMA),
	9862:  uint16(3),
	9863:  uint16(2),
	9864:  uint16(sym_block_comment),
	9865:  uint16(sym_line_comment),
	9866:  uint16(3),
	9867:  uint16(234),
	9868:  uint16(1),
	9869:  uint16(anon_sym_LBRACE),
	9870:  uint16(164),
	9871:  uint16(1),
	9872:  uint16(sym_compound_statement),
	9873:  uint16(3),
	9874:  uint16(2),
	9875:  uint16(sym_block_comment),
	9876:  uint16(sym_line_comment),
	9877:  uint16(3),
	9878:  uint16(982),
	9879:  uint16(1),
	9880:  uint16(anon_sym_SEMI),
	9881:  uint16(984),
	9882:  uint16(1),
	9883:  uint16(anon_sym_EQ),
	9884:  uint16(3),
	9885:  uint16(2),
	9886:  uint16(sym_block_comment),
	9887:  uint16(sym_line_comment),
	9888:  uint16(3),
	9889:  uint16(780),
	9890:  uint16(1),
	9891:  uint16(anon_sym_RBRACE),
	9892:  uint16(986),
	9893:  uint16(1),
	9894:  uint16(anon_sym_COMMA),
	9895:  uint16(3),
	9896:  uint16(2),
	9897:  uint16(sym_block_comment),
	9898:  uint16(sym_line_comment),
	9899:  uint16(3),
	9900:  uint16(234),
	9901:  uint16(1),
	9902:  uint16(anon_sym_LBRACE),
	9903:  uint16(128),
	9904:  uint16(1),
	9905:  uint16(sym_compound_statement),
	9906:  uint16(3),
	9907:  uint16(2),
	9908:  uint16(sym_block_comment),
	9909:  uint16(sym_line_comment),
	9910:  uint16(3),
	9911:  uint16(234),
	9912:  uint16(1),
	9913:  uint16(anon_sym_LBRACE),
	9914:  uint16(156),
	9915:  uint16(1),
	9916:  uint16(sym_compound_statement),
	9917:  uint16(3),
	9918:  uint16(2),
	9919:  uint16(sym_block_comment),
	9920:  uint16(sym_line_comment),
	9921:  uint16(3),
	9922:  uint16(912),
	9923:  uint16(1),
	9924:  uint16(anon_sym_SEMI),
	9925:  uint16(914),
	9926:  uint16(1),
	9927:  uint16(anon_sym_EQ),
	9928:  uint16(3),
	9929:  uint16(2),
	9930:  uint16(sym_block_comment),
	9931:  uint16(sym_line_comment),
	9932:  uint16(3),
	9933:  uint16(988),
	9934:  uint16(1),
	9935:  uint16(anon_sym_COMMA),
	9936:  uint16(990),
	9937:  uint16(1),
	9938:  uint16(anon_sym_GT),
	9939:  uint16(3),
	9940:  uint16(2),
	9941:  uint16(sym_block_comment),
	9942:  uint16(sym_line_comment),
	9943:  uint16(3),
	9944:  uint16(992),
	9945:  uint16(1),
	9946:  uint16(sym_identifier),
	9947:  uint16(442),
	9948:  uint16(1),
	9949:  uint16(sym_import_path),
	9950:  uint16(3),
	9951:  uint16(2),
	9952:  uint16(sym_block_comment),
	9953:  uint16(sym_line_comment),
	9954:  uint16(3),
	9955:  uint16(994),
	9956:  uint16(1),
	9957:  uint16(sym_identifier),
	9958:  uint16(404),
	9959:  uint16(1),
	9960:  uint16(sym_import_path),
	9961:  uint16(3),
	9962:  uint16(2),
	9963:  uint16(sym_block_comment),
	9964:  uint16(sym_line_comment),
	9965:  uint16(3),
	9966:  uint16(904),
	9967:  uint16(1),
	9968:  uint16(anon_sym_LBRACE),
	9969:  uint16(270),
	9970:  uint16(1),
	9971:  uint16(sym_case_compound_statement),
	9972:  uint16(3),
	9973:  uint16(2),
	9974:  uint16(sym_block_comment),
	9975:  uint16(sym_line_comment),
	9976:  uint16(3),
	9977:  uint16(234),
	9978:  uint16(1),
	9979:  uint16(anon_sym_LBRACE),
	9980:  uint16(166),
	9981:  uint16(1),
	9982:  uint16(sym_compound_statement),
	9983:  uint16(3),
	9984:  uint16(2),
	9985:  uint16(sym_block_comment),
	9986:  uint16(sym_line_comment),
	9987:  uint16(3),
	9988:  uint16(996),
	9989:  uint16(1),
	9990:  uint16(anon_sym_COMMA),
	9991:  uint16(998),
	9992:  uint16(1),
	9993:  uint16(anon_sym_GT),
	9994:  uint16(3),
	9995:  uint16(2),
	9996:  uint16(sym_block_comment),
	9997:  uint16(sym_line_comment),
	9998:  uint16(3),
	9999:  uint16(994),
	10000: uint16(1),
	10001: uint16(sym_identifier),
	10002: uint16(401),
	10003: uint16(1),
	10004: uint16(sym_import_path),
	10005: uint16(3),
	10006: uint16(2),
	10007: uint16(sym_block_comment),
	10008: uint16(sym_line_comment),
	10009: uint16(3),
	10010: uint16(1000),
	10011: uint16(1),
	10012: uint16(anon_sym_SEMI),
	10013: uint16(1002),
	10014: uint16(1),
	10015: uint16(anon_sym_if),
	10016: uint16(3),
	10017: uint16(2),
	10018: uint16(sym_block_comment),
	10019: uint16(sym_line_comment),
	10020: uint16(3),
	10021: uint16(234),
	10022: uint16(1),
	10023: uint16(anon_sym_LBRACE),
	10024: uint16(163),
	10025: uint16(1),
	10026: uint16(sym_compound_statement),
	10027: uint16(3),
	10028: uint16(2),
	10029: uint16(sym_block_comment),
	10030: uint16(sym_line_comment),
	10031: uint16(3),
	10032: uint16(1004),
	10033: uint16(1),
	10034: uint16(sym_identifier),
	10035: uint16(302),
	10036: uint16(1),
	10037: uint16(sym_variable_identifier_declaration),
	10038: uint16(3),
	10039: uint16(2),
	10040: uint16(sym_block_comment),
	10041: uint16(sym_line_comment),
	10042: uint16(2),
	10043: uint16(3),
	10044: uint16(2),
	10045: uint16(sym_block_comment),
	10046: uint16(sym_line_comment),
	10047: uint16(928),
	10048: uint16(2),
	10049: uint16(anon_sym_SEMI),
	10050: uint16(anon_sym_EQ),
	10051: uint16(3),
	10052: uint16(934),
	10053: uint16(1),
	10054: uint16(anon_sym_LPAREN),
	10055: uint16(96),
	10056: uint16(1),
	10057: uint16(sym_argument_list_expression),
	10058: uint16(3),
	10059: uint16(2),
	10060: uint16(sym_block_comment),
	10061: uint16(sym_line_comment),
	10062: uint16(3),
	10063: uint16(1006),
	10064: uint16(1),
	10065: uint16(anon_sym_SEMI),
	10066: uint16(1008),
	10067: uint16(1),
	10068: uint16(anon_sym_EQ),
	10069: uint16(3),
	10070: uint16(2),
	10071: uint16(sym_block_comment),
	10072: uint16(sym_line_comment),
	10073: uint16(3),
	10074: uint16(234),
	10075: uint16(1),
	10076: uint16(anon_sym_LBRACE),
	10077: uint16(155),
	10078: uint16(1),
	10079: uint16(sym_compound_statement),
	10080: uint16(3),
	10081: uint16(2),
	10082: uint16(sym_block_comment),
	10083: uint16(sym_line_comment),
	10084: uint16(3),
	10085: uint16(904),
	10086: uint16(1),
	10087: uint16(anon_sym_LBRACE),
	10088: uint16(288),
	10089: uint16(1),
	10090: uint16(sym_case_compound_statement),
	10091: uint16(3),
	10092: uint16(2),
	10093: uint16(sym_block_comment),
	10094: uint16(sym_line_comment),
	10095: uint16(2),
	10096: uint16(3),
	10097: uint16(2),
	10098: uint16(sym_block_comment),
	10099: uint16(sym_line_comment),
	10100: uint16(1010),
	10101: uint16(2),
	10102: uint16(anon_sym_COMMA),
	10103: uint16(anon_sym_RPAREN),
	10104: uint16(3),
	10105: uint16(938),
	10106: uint16(1),
	10107: uint16(anon_sym_RPAREN),
	10108: uint16(1012),
	10109: uint16(1),
	10110: uint16(anon_sym_COMMA),
	10111: uint16(3),
	10112: uint16(2),
	10113: uint16(sym_block_comment),
	10114: uint16(sym_line_comment),
	10115: uint16(3),
	10116: uint16(920),
	10117: uint16(1),
	10118: uint16(anon_sym_SEMI),
	10119: uint16(922),
	10120: uint16(1),
	10121: uint16(anon_sym_EQ),
	10122: uint16(3),
	10123: uint16(2),
	10124: uint16(sym_block_comment),
	10125: uint16(sym_line_comment),
	10126: uint16(3),
	10127: uint16(1014),
	10128: uint16(1),
	10129: uint16(anon_sym_SEMI),
	10130: uint16(1016),
	10131: uint16(1),
	10132: uint16(anon_sym_EQ),
	10133: uint16(3),
	10134: uint16(2),
	10135: uint16(sym_block_comment),
	10136: uint16(sym_line_comment),
	10137: uint16(3),
	10138: uint16(35),
	10139: uint16(1),
	10140: uint16(anon_sym_LPAREN),
	10141: uint16(111),
	10142: uint16(1),
	10143: uint16(sym_parenthesized_expression),
	10144: uint16(3),
	10145: uint16(2),
	10146: uint16(sym_block_comment),
	10147: uint16(sym_line_comment),
	10148: uint16(3),
	10149: uint16(1018),
	10150: uint16(1),
	10151: uint16(sym_identifier),
	10152: uint16(380),
	10153: uint16(1),
	10154: uint16(sym_variable_identifier_declaration),
	10155: uint16(3),
	10156: uint16(2),
	10157: uint16(sym_block_comment),
	10158: uint16(sym_line_comment),
	10159: uint16(3),
	10160: uint16(1020),
	10161: uint16(1),
	10162: uint16(anon_sym_COMMA),
	10163: uint16(1022),
	10164: uint16(1),
	10165: uint16(anon_sym_RPAREN),
	10166: uint16(3),
	10167: uint16(2),
	10168: uint16(sym_block_comment),
	10169: uint16(sym_line_comment),
	10170: uint16(3),
	10171: uint16(916),
	10172: uint16(1),
	10173: uint16(anon_sym_COLON),
	10174: uint16(1024),
	10175: uint16(1),
	10176: uint16(anon_sym_EQ),
	10177: uint16(3),
	10178: uint16(2),
	10179: uint16(sym_block_comment),
	10180: uint16(sym_line_comment),
	10181: uint16(3),
	10182: uint16(1026),
	10183: uint16(1),
	10184: uint16(anon_sym_LBRACE),
	10185: uint16(367),
	10186: uint16(1),
	10187: uint16(sym_continuing_compound_statement),
	10188: uint16(3),
	10189: uint16(2),
	10190: uint16(sym_block_comment),
	10191: uint16(sym_line_comment),
	10192: uint16(3),
	10193: uint16(797),
	10194: uint16(1),
	10195: uint16(anon_sym_RBRACE),
	10196: uint16(1028),
	10197: uint16(1),
	10198: uint16(anon_sym_COMMA),
	10199: uint16(3),
	10200: uint16(2),
	10201: uint16(sym_block_comment),
	10202: uint16(sym_line_comment),
	10203: uint16(2),
	10204: uint16(3),
	10205: uint16(2),
	10206: uint16(sym_block_comment),
	10207: uint16(sym_line_comment),
	10208: uint16(1030),
	10209: uint16(2),
	10210: uint16(anon_sym_COMMA),
	10211: uint16(anon_sym_GT),
	10212: uint16(3),
	10213: uint16(992),
	10214: uint16(1),
	10215: uint16(sym_identifier),
	10216: uint16(253),
	10217: uint16(1),
	10218: uint16(sym_import_path),
	10219: uint16(3),
	10220: uint16(2),
	10221: uint16(sym_block_comment),
	10222: uint16(sym_line_comment),
	10223: uint16(2),
	10224: uint16(3),
	10225: uint16(2),
	10226: uint16(sym_block_comment),
	10227: uint16(sym_line_comment),
	10228: uint16(815),
	10229: uint16(2),
	10230: uint16(anon_sym_LPAREN),
	10231: uint16(anon_sym_COLON_COLON),
	10232: uint16(3),
	10233: uint16(859),
	10234: uint16(1),
	10235: uint16(anon_sym_RPAREN),
	10236: uint16(1032),
	10237: uint16(1),
	10238: uint16(anon_sym_COMMA),
	10239: uint16(3),
	10240: uint16(2),
	10241: uint16(sym_block_comment),
	10242: uint16(sym_line_comment),
	10243: uint16(3),
	10244: uint16(1034),
	10245: uint16(1),
	10246: uint16(anon_sym_COMMA),
	10247: uint16(1036),
	10248: uint16(1),
	10249: uint16(anon_sym_RPAREN),
	10250: uint16(3),
	10251: uint16(2),
	10252: uint16(sym_block_comment),
	10253: uint16(sym_line_comment),
	10254: uint16(2),
	10255: uint16(3),
	10256: uint16(2),
	10257: uint16(sym_block_comment),
	10258: uint16(sym_line_comment),
	10259: uint16(1038),
	10260: uint16(2),
	10261: uint16(anon_sym_COMMA),
	10262: uint16(anon_sym_RPAREN),
	10263: uint16(4),
	10264: uint16(3),
	10265: uint16(1),
	10266: uint16(sym_block_comment),
	10267: uint16(813),
	10268: uint16(1),
	10269: uint16(sym_line_comment),
	10270: uint16(953),
	10271: uint16(1),
	10272: uint16(anon_sym_LF),
	10273: uint16(1040),
	10274: uint16(1),
	10275: uint16(anon_sym_COMMA),
	10276: uint16(3),
	10277: uint16(1042),
	10278: uint16(1),
	10279: uint16(sym_identifier),
	10280: uint16(311),
	10281: uint16(1),
	10282: uint16(sym_variable_identifier_declaration),
	10283: uint16(3),
	10284: uint16(2),
	10285: uint16(sym_block_comment),
	10286: uint16(sym_line_comment),
	10287: uint16(2),
	10288: uint16(1044),
	10289: uint16(1),
	10290: uint16(anon_sym_RBRACE),
	10291: uint16(3),
	10292: uint16(2),
	10293: uint16(sym_block_comment),
	10294: uint16(sym_line_comment),
	10295: uint16(2),
	10296: uint16(1046),
	10297: uint16(1),
	10298: uint16(anon_sym_LT),
	10299: uint16(3),
	10300: uint16(2),
	10301: uint16(sym_block_comment),
	10302: uint16(sym_line_comment),
	10303: uint16(2),
	10304: uint16(1048),
	10305: uint16(1),
	10306: uint16(anon_sym_GT),
	10307: uint16(3),
	10308: uint16(2),
	10309: uint16(sym_block_comment),
	10310: uint16(sym_line_comment),
	10311: uint16(2),
	10312: uint16(1050),
	10313: uint16(1),
	10314: uint16(anon_sym_RPAREN),
	10315: uint16(3),
	10316: uint16(2),
	10317: uint16(sym_block_comment),
	10318: uint16(sym_line_comment),
	10319: uint16(2),
	10320: uint16(1052),
	10321: uint16(1),
	10322: uint16(anon_sym_GT),
	10323: uint16(3),
	10324: uint16(2),
	10325: uint16(sym_block_comment),
	10326: uint16(sym_line_comment),
	10327: uint16(2),
	10328: uint16(1054),
	10329: uint16(1),
	10330: uint16(sym_identifier),
	10331: uint16(3),
	10332: uint16(2),
	10333: uint16(sym_block_comment),
	10334: uint16(sym_line_comment),
	10335: uint16(2),
	10336: uint16(1056),
	10337: uint16(1),
	10338: uint16(anon_sym_LT),
	10339: uint16(3),
	10340: uint16(2),
	10341: uint16(sym_block_comment),
	10342: uint16(sym_line_comment),
	10343: uint16(2),
	10344: uint16(803),
	10345: uint16(1),
	10346: uint16(anon_sym_RBRACE),
	10347: uint16(3),
	10348: uint16(2),
	10349: uint16(sym_block_comment),
	10350: uint16(sym_line_comment),
	10351: uint16(2),
	10352: uint16(1058),
	10353: uint16(1),
	10354: uint16(anon_sym_COMMA),
	10355: uint16(3),
	10356: uint16(2),
	10357: uint16(sym_block_comment),
	10358: uint16(sym_line_comment),
	10359: uint16(3),
	10360: uint16(3),
	10361: uint16(1),
	10362: uint16(sym_block_comment),
	10363: uint16(813),
	10364: uint16(1),
	10365: uint16(sym_line_comment),
	10366: uint16(1060),
	10367: uint16(1),
	10368: uint16(anon_sym_LF),
	10369: uint16(2),
	10370: uint16(1062),
	10371: uint16(1),
	10372: uint16(anon_sym_LT),
	10373: uint16(3),
	10374: uint16(2),
	10375: uint16(sym_block_comment),
	10376: uint16(sym_line_comment),
	10377: uint16(2),
	10378: uint16(1064),
	10379: uint16(1),
	10380: uint16(sym_identifier),
	10381: uint16(3),
	10382: uint16(2),
	10383: uint16(sym_block_comment),
	10384: uint16(sym_line_comment),
	10385: uint16(2),
	10386: uint16(1066),
	10387: uint16(1),
	10388: uint16(anon_sym_COMMA),
	10389: uint16(3),
	10390: uint16(2),
	10391: uint16(sym_block_comment),
	10392: uint16(sym_line_comment),
	10393: uint16(2),
	10394: uint16(972),
	10395: uint16(1),
	10396: uint16(anon_sym_GT),
	10397: uint16(3),
	10398: uint16(2),
	10399: uint16(sym_block_comment),
	10400: uint16(sym_line_comment),
	10401: uint16(2),
	10402: uint16(1068),
	10403: uint16(1),
	10404: uint16(anon_sym_COMMA),
	10405: uint16(3),
	10406: uint16(2),
	10407: uint16(sym_block_comment),
	10408: uint16(sym_line_comment),
	10409: uint16(2),
	10410: uint16(1070),
	10411: uint16(1),
	10412: uint16(anon_sym_COMMA),
	10413: uint16(3),
	10414: uint16(2),
	10415: uint16(sym_block_comment),
	10416: uint16(sym_line_comment),
	10417: uint16(2),
	10418: uint16(1072),
	10419: uint16(1),
	10420: uint16(sym_identifier),
	10421: uint16(3),
	10422: uint16(2),
	10423: uint16(sym_block_comment),
	10424: uint16(sym_line_comment),
	10425: uint16(2),
	10426: uint16(1074),
	10427: uint16(1),
	10428: uint16(anon_sym_RBRACE),
	10429: uint16(3),
	10430: uint16(2),
	10431: uint16(sym_block_comment),
	10432: uint16(sym_line_comment),
	10433: uint16(2),
	10434: uint16(1076),
	10435: uint16(1),
	10436: uint16(anon_sym_COMMA),
	10437: uint16(3),
	10438: uint16(2),
	10439: uint16(sym_block_comment),
	10440: uint16(sym_line_comment),
	10441: uint16(2),
	10442: uint16(1078),
	10443: uint16(1),
	10444: uint16(sym_identifier),
	10445: uint16(3),
	10446: uint16(2),
	10447: uint16(sym_block_comment),
	10448: uint16(sym_line_comment),
	10449: uint16(2),
	10450: uint16(1080),
	10451: uint16(1),
	10452: uint16(sym_identifier),
	10453: uint16(3),
	10454: uint16(2),
	10455: uint16(sym_block_comment),
	10456: uint16(sym_line_comment),
	10457: uint16(2),
	10458: uint16(1082),
	10459: uint16(1),
	10460: uint16(anon_sym_LT),
	10461: uint16(3),
	10462: uint16(2),
	10463: uint16(sym_block_comment),
	10464: uint16(sym_line_comment),
	10465: uint16(2),
	10466: uint16(1084),
	10467: uint16(1),
	10468: uint16(anon_sym_RBRACE),
	10469: uint16(3),
	10470: uint16(2),
	10471: uint16(sym_block_comment),
	10472: uint16(sym_line_comment),
	10473: uint16(2),
	10474: uint16(1086),
	10475: uint16(1),
	10476: uint16(anon_sym_SEMI),
	10477: uint16(3),
	10478: uint16(2),
	10479: uint16(sym_block_comment),
	10480: uint16(sym_line_comment),
	10481: uint16(2),
	10482: uint16(1088),
	10483: uint16(1),
	10484: uint16(anon_sym_RBRACE),
	10485: uint16(3),
	10486: uint16(2),
	10487: uint16(sym_block_comment),
	10488: uint16(sym_line_comment),
	10489: uint16(2),
	10490: uint16(1090),
	10491: uint16(1),
	10492: uint16(sym_identifier),
	10493: uint16(3),
	10494: uint16(2),
	10495: uint16(sym_block_comment),
	10496: uint16(sym_line_comment),
	10497: uint16(2),
	10498: uint16(1092),
	10499: uint16(1),
	10500: uint16(sym_identifier),
	10501: uint16(3),
	10502: uint16(2),
	10503: uint16(sym_block_comment),
	10504: uint16(sym_line_comment),
	10505: uint16(2),
	10506: uint16(1094),
	10507: uint16(1),
	10508: uint16(aux_sym_preproc_ifdef_token3),
	10509: uint16(3),
	10510: uint16(2),
	10511: uint16(sym_block_comment),
	10512: uint16(sym_line_comment),
	10513: uint16(2),
	10514: uint16(1096),
	10515: uint16(1),
	10516: uint16(anon_sym_LBRACE),
	10517: uint16(3),
	10518: uint16(2),
	10519: uint16(sym_block_comment),
	10520: uint16(sym_line_comment),
	10521: uint16(2),
	10522: uint16(1098),
	10523: uint16(1),
	10524: uint16(anon_sym_LPAREN),
	10525: uint16(3),
	10526: uint16(2),
	10527: uint16(sym_block_comment),
	10528: uint16(sym_line_comment),
	10529: uint16(2),
	10530: uint16(1100),
	10531: uint16(1),
	10532: uint16(anon_sym_GT),
	10533: uint16(3),
	10534: uint16(2),
	10535: uint16(sym_block_comment),
	10536: uint16(sym_line_comment),
	10537: uint16(2),
	10538: uint16(936),
	10539: uint16(1),
	10540: uint16(anon_sym_LT),
	10541: uint16(3),
	10542: uint16(2),
	10543: uint16(sym_block_comment),
	10544: uint16(sym_line_comment),
	10545: uint16(2),
	10546: uint16(1102),
	10547: uint16(1),
	10548: uint16(aux_sym_preproc_ifdef_token3),
	10549: uint16(3),
	10550: uint16(2),
	10551: uint16(sym_block_comment),
	10552: uint16(sym_line_comment),
	10553: uint16(2),
	10554: uint16(1104),
	10555: uint16(1),
	10556: uint16(anon_sym_COMMA),
	10557: uint16(3),
	10558: uint16(2),
	10559: uint16(sym_block_comment),
	10560: uint16(sym_line_comment),
	10561: uint16(2),
	10562: uint16(1106),
	10563: uint16(1),
	10564: uint16(anon_sym_GT),
	10565: uint16(3),
	10566: uint16(2),
	10567: uint16(sym_block_comment),
	10568: uint16(sym_line_comment),
	10569: uint16(2),
	10570: uint16(1024),
	10571: uint16(1),
	10572: uint16(anon_sym_EQ),
	10573: uint16(3),
	10574: uint16(2),
	10575: uint16(sym_block_comment),
	10576: uint16(sym_line_comment),
	10577: uint16(2),
	10578: uint16(1108),
	10579: uint16(1),
	10580: uint16(anon_sym_LBRACE),
	10581: uint16(3),
	10582: uint16(2),
	10583: uint16(sym_block_comment),
	10584: uint16(sym_line_comment),
	10585: uint16(2),
	10586: uint16(599),
	10587: uint16(1),
	10588: uint16(anon_sym_SEMI),
	10589: uint16(3),
	10590: uint16(2),
	10591: uint16(sym_block_comment),
	10592: uint16(sym_line_comment),
	10593: uint16(2),
	10594: uint16(1110),
	10595: uint16(1),
	10596: uint16(anon_sym_SEMI),
	10597: uint16(3),
	10598: uint16(2),
	10599: uint16(sym_block_comment),
	10600: uint16(sym_line_comment),
	10601: uint16(2),
	10602: uint16(968),
	10603: uint16(1),
	10604: uint16(anon_sym_EQ),
	10605: uint16(3),
	10606: uint16(2),
	10607: uint16(sym_block_comment),
	10608: uint16(sym_line_comment),
	10609: uint16(2),
	10610: uint16(990),
	10611: uint16(1),
	10612: uint16(anon_sym_GT),
	10613: uint16(3),
	10614: uint16(2),
	10615: uint16(sym_block_comment),
	10616: uint16(sym_line_comment),
	10617: uint16(2),
	10618: uint16(1112),
	10619: uint16(1),
	10620: uint16(anon_sym_RPAREN),
	10621: uint16(3),
	10622: uint16(2),
	10623: uint16(sym_block_comment),
	10624: uint16(sym_line_comment),
	10625: uint16(2),
	10626: uint16(1114),
	10627: uint16(1),
	10628: uint16(aux_sym_preproc_ifdef_token3),
	10629: uint16(3),
	10630: uint16(2),
	10631: uint16(sym_block_comment),
	10632: uint16(sym_line_comment),
	10633: uint16(2),
	10634: uint16(1116),
	10635: uint16(1),
	10636: uint16(sym_identifier),
	10637: uint16(3),
	10638: uint16(2),
	10639: uint16(sym_block_comment),
	10640: uint16(sym_line_comment),
	10641: uint16(2),
	10642: uint16(1118),
	10643: uint16(1),
	10644: uint16(anon_sym_COMMA),
	10645: uint16(3),
	10646: uint16(2),
	10647: uint16(sym_block_comment),
	10648: uint16(sym_line_comment),
	10649: uint16(3),
	10650: uint16(3),
	10651: uint16(1),
	10652: uint16(sym_block_comment),
	10653: uint16(813),
	10654: uint16(1),
	10655: uint16(sym_line_comment),
	10656: uint16(1120),
	10657: uint16(1),
	10658: uint16(anon_sym_LF),
	10659: uint16(2),
	10660: uint16(1122),
	10661: uint16(1),
	10662: uint16(anon_sym_RBRACE),
	10663: uint16(3),
	10664: uint16(2),
	10665: uint16(sym_block_comment),
	10666: uint16(sym_line_comment),
	10667: uint16(2),
	10668: uint16(1124),
	10669: uint16(1),
	10670: uint16(sym_identifier),
	10671: uint16(3),
	10672: uint16(2),
	10673: uint16(sym_block_comment),
	10674: uint16(sym_line_comment),
	10675: uint16(2),
	10676: uint16(287),
	10677: uint16(1),
	10678: uint16(anon_sym_RBRACE),
	10679: uint16(3),
	10680: uint16(2),
	10681: uint16(sym_block_comment),
	10682: uint16(sym_line_comment),
	10683: uint16(2),
	10684: uint16(1126),
	10685: uint16(1),
	10686: uint16(sym_identifier),
	10687: uint16(3),
	10688: uint16(2),
	10689: uint16(sym_block_comment),
	10690: uint16(sym_line_comment),
	10691: uint16(2),
	10692: uint16(1128),
	10693: uint16(1),
	10694: uint16(anon_sym_LT),
	10695: uint16(3),
	10696: uint16(2),
	10697: uint16(sym_block_comment),
	10698: uint16(sym_line_comment),
	10699: uint16(2),
	10700: uint16(131),
	10701: uint16(1),
	10702: uint16(anon_sym_RPAREN),
	10703: uint16(3),
	10704: uint16(2),
	10705: uint16(sym_block_comment),
	10706: uint16(sym_line_comment),
	10707: uint16(2),
	10708: uint16(1130),
	10709: uint16(1),
	10710: uint16(aux_sym_preproc_ifdef_token3),
	10711: uint16(3),
	10712: uint16(2),
	10713: uint16(sym_block_comment),
	10714: uint16(sym_line_comment),
	10715: uint16(2),
	10716: uint16(1132),
	10717: uint16(1),
	10718: uint16(anon_sym_SEMI),
	10719: uint16(3),
	10720: uint16(2),
	10721: uint16(sym_block_comment),
	10722: uint16(sym_line_comment),
	10723: uint16(2),
	10724: uint16(1134),
	10725: uint16(1),
	10726: uint16(sym_identifier),
	10727: uint16(3),
	10728: uint16(2),
	10729: uint16(sym_block_comment),
	10730: uint16(sym_line_comment),
	10731: uint16(2),
	10732: uint16(1136),
	10733: uint16(1),
	10734: uint16(sym_identifier),
	10735: uint16(3),
	10736: uint16(2),
	10737: uint16(sym_block_comment),
	10738: uint16(sym_line_comment),
	10739: uint16(2),
	10740: uint16(1138),
	10741: uint16(1),
	10742: uint16(anon_sym_LPAREN),
	10743: uint16(3),
	10744: uint16(2),
	10745: uint16(sym_block_comment),
	10746: uint16(sym_line_comment),
	10747: uint16(2),
	10748: uint16(1140),
	10749: uint16(1),
	10751: uint16(3),
	10752: uint16(2),
	10753: uint16(sym_block_comment),
	10754: uint16(sym_line_comment),
	10755: uint16(2),
	10756: uint16(1142),
	10757: uint16(1),
	10758: uint16(anon_sym_RPAREN),
	10759: uint16(3),
	10760: uint16(2),
	10761: uint16(sym_block_comment),
	10762: uint16(sym_line_comment),
	10763: uint16(2),
	10764: uint16(1144),
	10765: uint16(1),
	10766: uint16(anon_sym_LPAREN),
	10767: uint16(3),
	10768: uint16(2),
	10769: uint16(sym_block_comment),
	10770: uint16(sym_line_comment),
	10771: uint16(2),
	10772: uint16(1146),
	10773: uint16(1),
	10774: uint16(anon_sym_SEMI),
	10775: uint16(3),
	10776: uint16(2),
	10777: uint16(sym_block_comment),
	10778: uint16(sym_line_comment),
	10779: uint16(2),
	10780: uint16(1148),
	10781: uint16(1),
	10782: uint16(anon_sym_LBRACE),
	10783: uint16(3),
	10784: uint16(2),
	10785: uint16(sym_block_comment),
	10786: uint16(sym_line_comment),
	10787: uint16(2),
	10788: uint16(1150),
	10789: uint16(1),
	10790: uint16(aux_sym_preproc_ifdef_token3),
	10791: uint16(3),
	10792: uint16(2),
	10793: uint16(sym_block_comment),
	10794: uint16(sym_line_comment),
	10795: uint16(2),
	10796: uint16(1152),
	10797: uint16(1),
	10798: uint16(aux_sym_preproc_ifdef_token3),
	10799: uint16(3),
	10800: uint16(2),
	10801: uint16(sym_block_comment),
	10802: uint16(sym_line_comment),
	10803: uint16(2),
	10804: uint16(1154),
	10805: uint16(1),
	10806: uint16(anon_sym_SEMI),
	10807: uint16(3),
	10808: uint16(2),
	10809: uint16(sym_block_comment),
	10810: uint16(sym_line_comment),
	10811: uint16(2),
	10812: uint16(236),
	10813: uint16(1),
	10814: uint16(anon_sym_RBRACE),
	10815: uint16(3),
	10816: uint16(2),
	10817: uint16(sym_block_comment),
	10818: uint16(sym_line_comment),
	10819: uint16(2),
	10820: uint16(1156),
	10821: uint16(1),
	10822: uint16(anon_sym_SEMI),
	10823: uint16(3),
	10824: uint16(2),
	10825: uint16(sym_block_comment),
	10826: uint16(sym_line_comment),
	10827: uint16(2),
	10828: uint16(1158),
	10829: uint16(1),
	10830: uint16(anon_sym_RPAREN),
	10831: uint16(3),
	10832: uint16(2),
	10833: uint16(sym_block_comment),
	10834: uint16(sym_line_comment),
	10835: uint16(2),
	10836: uint16(1160),
	10837: uint16(1),
	10838: uint16(anon_sym_RBRACE),
	10839: uint16(3),
	10840: uint16(2),
	10841: uint16(sym_block_comment),
	10842: uint16(sym_line_comment),
	10843: uint16(2),
	10844: uint16(1162),
	10845: uint16(1),
	10846: uint16(anon_sym_RBRACE),
	10847: uint16(3),
	10848: uint16(2),
	10849: uint16(sym_block_comment),
	10850: uint16(sym_line_comment),
	10851: uint16(2),
	10852: uint16(133),
	10853: uint16(1),
	10854: uint16(anon_sym_RPAREN),
	10855: uint16(3),
	10856: uint16(2),
	10857: uint16(sym_block_comment),
	10858: uint16(sym_line_comment),
	10859: uint16(2),
	10860: uint16(797),
	10861: uint16(1),
	10862: uint16(anon_sym_RBRACE),
	10863: uint16(3),
	10864: uint16(2),
	10865: uint16(sym_block_comment),
	10866: uint16(sym_line_comment),
	10867: uint16(2),
	10868: uint16(1164),
	10869: uint16(1),
	10870: uint16(anon_sym_SEMI),
	10871: uint16(3),
	10872: uint16(2),
	10873: uint16(sym_block_comment),
	10874: uint16(sym_line_comment),
	10875: uint16(2),
	10876: uint16(1166),
	10877: uint16(1),
	10878: uint16(anon_sym_EQ),
	10879: uint16(3),
	10880: uint16(2),
	10881: uint16(sym_block_comment),
	10882: uint16(sym_line_comment),
	10883: uint16(2),
	10884: uint16(916),
	10885: uint16(1),
	10886: uint16(anon_sym_COLON),
	10887: uint16(3),
	10888: uint16(2),
	10889: uint16(sym_block_comment),
	10890: uint16(sym_line_comment),
	10891: uint16(2),
	10892: uint16(1168),
	10893: uint16(1),
	10894: uint16(anon_sym_SEMI),
	10895: uint16(3),
	10896: uint16(2),
	10897: uint16(sym_block_comment),
	10898: uint16(sym_line_comment),
	10899: uint16(2),
	10900: uint16(1170),
	10901: uint16(1),
	10902: uint16(anon_sym_RBRACE),
	10903: uint16(3),
	10904: uint16(2),
	10905: uint16(sym_block_comment),
	10906: uint16(sym_line_comment),
	10907: uint16(2),
	10908: uint16(279),
	10909: uint16(1),
	10910: uint16(anon_sym_RBRACE),
	10911: uint16(3),
	10912: uint16(2),
	10913: uint16(sym_block_comment),
	10914: uint16(sym_line_comment),
	10915: uint16(2),
	10916: uint16(1172),
	10917: uint16(1),
	10918: uint16(anon_sym_SEMI),
	10919: uint16(3),
	10920: uint16(2),
	10921: uint16(sym_block_comment),
	10922: uint16(sym_line_comment),
	10923: uint16(2),
	10924: uint16(1174),
	10925: uint16(1),
	10926: uint16(anon_sym_RPAREN),
	10927: uint16(3),
	10928: uint16(2),
	10929: uint16(sym_block_comment),
	10930: uint16(sym_line_comment),
	10931: uint16(2),
	10932: uint16(1176),
	10933: uint16(1),
	10934: uint16(anon_sym_RBRACE),
	10935: uint16(3),
	10936: uint16(2),
	10937: uint16(sym_block_comment),
	10938: uint16(sym_line_comment),
	10939: uint16(2),
	10940: uint16(1178),
	10941: uint16(1),
	10942: uint16(anon_sym_RPAREN),
	10943: uint16(3),
	10944: uint16(2),
	10945: uint16(sym_block_comment),
	10946: uint16(sym_line_comment),
	10947: uint16(2),
	10948: uint16(712),
	10949: uint16(1),
	10950: uint16(anon_sym_EQ),
	10951: uint16(3),
	10952: uint16(2),
	10953: uint16(sym_block_comment),
	10954: uint16(sym_line_comment),
	10955: uint16(2),
	10956: uint16(1180),
	10957: uint16(1),
	10958: uint16(sym_identifier),
	10959: uint16(3),
	10960: uint16(2),
	10961: uint16(sym_block_comment),
	10962: uint16(sym_line_comment),
	10963: uint16(2),
	10964: uint16(1182),
	10965: uint16(1),
	10966: uint16(anon_sym_RBRACE),
	10967: uint16(3),
	10968: uint16(2),
	10969: uint16(sym_block_comment),
	10970: uint16(sym_line_comment),
	10971: uint16(2),
	10972: uint16(1184),
	10973: uint16(1),
	10974: uint16(sym_identifier),
	10975: uint16(3),
	10976: uint16(2),
	10977: uint16(sym_block_comment),
	10978: uint16(sym_line_comment),
	10979: uint16(2),
	10980: uint16(1186),
	10981: uint16(1),
	10982: uint16(sym_identifier),
	10983: uint16(3),
	10984: uint16(2),
	10985: uint16(sym_block_comment),
	10986: uint16(sym_line_comment),
	10987: uint16(2),
	10988: uint16(1188),
	10989: uint16(1),
	10990: uint16(anon_sym_SEMI),
	10991: uint16(3),
	10992: uint16(2),
	10993: uint16(sym_block_comment),
	10994: uint16(sym_line_comment),
	10995: uint16(2),
	10996: uint16(1190),
	10997: uint16(1),
	10998: uint16(anon_sym_RPAREN),
	10999: uint16(3),
	11000: uint16(2),
	11001: uint16(sym_block_comment),
	11002: uint16(sym_line_comment),
	11003: uint16(2),
	11004: uint16(1192),
	11005: uint16(1),
	11006: uint16(sym_identifier),
	11007: uint16(3),
	11008: uint16(2),
	11009: uint16(sym_block_comment),
	11010: uint16(sym_line_comment),
	11011: uint16(2),
	11012: uint16(1194),
	11013: uint16(1),
	11014: uint16(anon_sym_LPAREN),
	11015: uint16(3),
	11016: uint16(2),
	11017: uint16(sym_block_comment),
	11018: uint16(sym_line_comment),
	11019: uint16(2),
	11020: uint16(1196),
	11021: uint16(1),
	11022: uint16(anon_sym_LBRACE),
	11023: uint16(3),
	11024: uint16(2),
	11025: uint16(sym_block_comment),
	11026: uint16(sym_line_comment),
	11027: uint16(2),
	11028: uint16(1198),
	11029: uint16(1),
	11030: uint16(anon_sym_LPAREN),
	11031: uint16(3),
	11032: uint16(2),
	11033: uint16(sym_block_comment),
	11034: uint16(sym_line_comment),
	11035: uint16(2),
	11036: uint16(1200),
	11037: uint16(1),
	11038: uint16(sym_identifier),
	11039: uint16(3),
	11040: uint16(2),
	11041: uint16(sym_block_comment),
	11042: uint16(sym_line_comment),
	11043: uint16(2),
	11044: uint16(1202),
	11045: uint16(1),
	11046: uint16(sym_identifier),
	11047: uint16(3),
	11048: uint16(2),
	11049: uint16(sym_block_comment),
	11050: uint16(sym_line_comment),
	11051: uint16(2),
	11052: uint16(1204),
	11053: uint16(1),
	11054: uint16(sym_identifier),
	11055: uint16(3),
	11056: uint16(2),
	11057: uint16(sym_block_comment),
	11058: uint16(sym_line_comment),
	11059: uint16(2),
	11060: uint16(1000),
	11061: uint16(1),
	11062: uint16(anon_sym_SEMI),
	11063: uint16(3),
	11064: uint16(2),
	11065: uint16(sym_block_comment),
	11066: uint16(sym_line_comment),
	11067: uint16(3),
	11068: uint16(3),
	11069: uint16(1),
	11070: uint16(sym_block_comment),
	11071: uint16(813),
	11072: uint16(1),
	11073: uint16(sym_line_comment),
	11074: uint16(1206),
	11075: uint16(1),
	11076: uint16(anon_sym_LF),
}

var ts_small_parse_table_map = [411]uint32_t{
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
	11:  uint32(920),
	12:  uint32(979),
	13:  uint32(1038),
	14:  uint32(1121),
	15:  uint32(1222),
	16:  uint32(1323),
	17:  uint32(1392),
	18:  uint32(1447),
	19:  uint32(1516),
	20:  uint32(1571),
	21:  uint32(1672),
	22:  uint32(1773),
	23:  uint32(1874),
	24:  uint32(1929),
	25:  uint32(2030),
	26:  uint32(2131),
	27:  uint32(2229),
	28:  uint32(2283),
	29:  uint32(2335),
	30:  uint32(2433),
	31:  uint32(2528),
	32:  uint32(2579),
	33:  uint32(2674),
	34:  uint32(2725),
	35:  uint32(2776),
	36:  uint32(2871),
	37:  uint32(2966),
	38:  uint32(3028),
	39:  uint32(3090),
	40:  uint32(3152),
	41:  uint32(3214),
	42:  uint32(3276),
	43:  uint32(3338),
	44:  uint32(3382),
	45:  uint32(3426),
	46:  uint32(3470),
	47:  uint32(3514),
	48:  uint32(3558),
	49:  uint32(3602),
	50:  uint32(3675),
	51:  uint32(3747),
	52:  uint32(3819),
	53:  uint32(3887),
	54:  uint32(3953),
	55:  uint32(3989),
	56:  uint32(4055),
	57:  uint32(4121),
	58:  uint32(4187),
	59:  uint32(4223),
	60:  uint32(4258),
	61:  uint32(4293),
	62:  uint32(4328),
	63:  uint32(4365),
	64:  uint32(4402),
	65:  uint32(4436),
	66:  uint32(4474),
	67:  uint32(4508),
	68:  uint32(4542),
	69:  uint32(4576),
	70:  uint32(4610),
	71:  uint32(4644),
	72:  uint32(4678),
	73:  uint32(4712),
	74:  uint32(4746),
	75:  uint32(4780),
	76:  uint32(4826),
	77:  uint32(4884),
	78:  uint32(4940),
	79:  uint32(4974),
	80:  uint32(5008),
	81:  uint32(5050),
	82:  uint32(5106),
	83:  uint32(5140),
	84:  uint32(5194),
	85:  uint32(5228),
	86:  uint32(5262),
	87:  uint32(5296),
	88:  uint32(5348),
	89:  uint32(5398),
	90:  uint32(5432),
	91:  uint32(5466),
	92:  uint32(5500),
	93:  uint32(5534),
	94:  uint32(5568),
	95:  uint32(5602),
	96:  uint32(5636),
	97:  uint32(5670),
	98:  uint32(5704),
	99:  uint32(5748),
	100: uint32(5782),
	101: uint32(5841),
	102: uint32(5898),
	103: uint32(5957),
	104: uint32(6016),
	105: uint32(6075),
	106: uint32(6131),
	107: uint32(6187),
	108: uint32(6243),
	109: uint32(6299),
	110: uint32(6355),
	111: uint32(6411),
	112: uint32(6467),
	113: uint32(6523),
	114: uint32(6579),
	115: uint32(6635),
	116: uint32(6691),
	117: uint32(6747),
	118: uint32(6803),
	119: uint32(6859),
	120: uint32(6887),
	121: uint32(6913),
	122: uint32(6936),
	123: uint32(6959),
	124: uint32(6982),
	125: uint32(7005),
	126: uint32(7028),
	127: uint32(7051),
	128: uint32(7074),
	129: uint32(7097),
	130: uint32(7120),
	131: uint32(7143),
	132: uint32(7166),
	133: uint32(7189),
	134: uint32(7212),
	135: uint32(7235),
	136: uint32(7258),
	137: uint32(7281),
	138: uint32(7304),
	139: uint32(7327),
	140: uint32(7350),
	141: uint32(7372),
	142: uint32(7394),
	143: uint32(7416),
	144: uint32(7444),
	145: uint32(7466),
	146: uint32(7494),
	147: uint32(7522),
	148: uint32(7550),
	149: uint32(7580),
	150: uint32(7608),
	151: uint32(7636),
	152: uint32(7678),
	153: uint32(7706),
	154: uint32(7728),
	155: uint32(7756),
	156: uint32(7780),
	157: uint32(7804),
	158: uint32(7828),
	159: uint32(7852),
	160: uint32(7894),
	161: uint32(7931),
	162: uint32(7952),
	163: uint32(7973),
	164: uint32(8009),
	165: uint32(8045),
	166: uint32(8083),
	167: uint32(8102),
	168: uint32(8121),
	169: uint32(8140),
	170: uint32(8167),
	171: uint32(8186),
	172: uint32(8221),
	173: uint32(8240),
	174: uint32(8273),
	175: uint32(8300),
	176: uint32(8327),
	177: uint32(8343),
	178: uint32(8373),
	179: uint32(8389),
	180: uint32(8405),
	181: uint32(8421),
	182: uint32(8447),
	183: uint32(8477),
	184: uint32(8493),
	185: uint32(8509),
	186: uint32(8539),
	187: uint32(8562),
	188: uint32(8579),
	189: uint32(8603),
	190: uint32(8625),
	191: uint32(8647),
	192: uint32(8661),
	193: uint32(8685),
	194: uint32(8703),
	195: uint32(8727),
	196: uint32(8749),
	197: uint32(8767),
	198: uint32(8788),
	199: uint32(8803),
	200: uint32(8824),
	201: uint32(8845),
	202: uint32(8860),
	203: uint32(8875),
	204: uint32(8890),
	205: uint32(8908),
	206: uint32(8926),
	207: uint32(8946),
	208: uint32(8964),
	209: uint32(8980),
	210: uint32(9000),
	211: uint32(9016),
	212: uint32(9028),
	213: uint32(9046),
	214: uint32(9064),
	215: uint32(9082),
	216: uint32(9098),
	217: uint32(9118),
	218: uint32(9135),
	219: uint32(9152),
	220: uint32(9163),
	221: uint32(9180),
	222: uint32(9199),
	223: uint32(9214),
	224: uint32(9227),
	225: uint32(9244),
	226: uint32(9259),
	227: uint32(9274),
	228: uint32(9289),
	229: uint32(9304),
	230: uint32(9319),
	231: uint32(9334),
	232: uint32(9351),
	233: uint32(9366),
	234: uint32(9381),
	235: uint32(9398),
	236: uint32(9415),
	237: uint32(9430),
	238: uint32(9445),
	239: uint32(9455),
	240: uint32(9465),
	241: uint32(9475),
	242: uint32(9489),
	243: uint32(9499),
	244: uint32(9509),
	245: uint32(9523),
	246: uint32(9537),
	247: uint32(9551),
	248: uint32(9563),
	249: uint32(9573),
	250: uint32(9585),
	251: uint32(9597),
	252: uint32(9611),
	253: uint32(9623),
	254: uint32(9639),
	255: uint32(9649),
	256: uint32(9659),
	257: uint32(9669),
	258: uint32(9679),
	259: uint32(9695),
	260: uint32(9705),
	261: uint32(9715),
	262: uint32(9729),
	263: uint32(9743),
	264: uint32(9757),
	265: uint32(9773),
	266: uint32(9784),
	267: uint32(9795),
	268: uint32(9806),
	269: uint32(9817),
	270: uint32(9826),
	271: uint32(9835),
	272: uint32(9846),
	273: uint32(9855),
	274: uint32(9866),
	275: uint32(9877),
	276: uint32(9888),
	277: uint32(9899),
	278: uint32(9910),
	279: uint32(9921),
	280: uint32(9932),
	281: uint32(9943),
	282: uint32(9954),
	283: uint32(9965),
	284: uint32(9976),
	285: uint32(9987),
	286: uint32(9998),
	287: uint32(10009),
	288: uint32(10020),
	289: uint32(10031),
	290: uint32(10042),
	291: uint32(10051),
	292: uint32(10062),
	293: uint32(10073),
	294: uint32(10084),
	295: uint32(10095),
	296: uint32(10104),
	297: uint32(10115),
	298: uint32(10126),
	299: uint32(10137),
	300: uint32(10148),
	301: uint32(10159),
	302: uint32(10170),
	303: uint32(10181),
	304: uint32(10192),
	305: uint32(10203),
	306: uint32(10212),
	307: uint32(10223),
	308: uint32(10232),
	309: uint32(10243),
	310: uint32(10254),
	311: uint32(10263),
	312: uint32(10276),
	313: uint32(10287),
	314: uint32(10295),
	315: uint32(10303),
	316: uint32(10311),
	317: uint32(10319),
	318: uint32(10327),
	319: uint32(10335),
	320: uint32(10343),
	321: uint32(10351),
	322: uint32(10359),
	323: uint32(10369),
	324: uint32(10377),
	325: uint32(10385),
	326: uint32(10393),
	327: uint32(10401),
	328: uint32(10409),
	329: uint32(10417),
	330: uint32(10425),
	331: uint32(10433),
	332: uint32(10441),
	333: uint32(10449),
	334: uint32(10457),
	335: uint32(10465),
	336: uint32(10473),
	337: uint32(10481),
	338: uint32(10489),
	339: uint32(10497),
	340: uint32(10505),
	341: uint32(10513),
	342: uint32(10521),
	343: uint32(10529),
	344: uint32(10537),
	345: uint32(10545),
	346: uint32(10553),
	347: uint32(10561),
	348: uint32(10569),
	349: uint32(10577),
	350: uint32(10585),
	351: uint32(10593),
	352: uint32(10601),
	353: uint32(10609),
	354: uint32(10617),
	355: uint32(10625),
	356: uint32(10633),
	357: uint32(10641),
	358: uint32(10649),
	359: uint32(10659),
	360: uint32(10667),
	361: uint32(10675),
	362: uint32(10683),
	363: uint32(10691),
	364: uint32(10699),
	365: uint32(10707),
	366: uint32(10715),
	367: uint32(10723),
	368: uint32(10731),
	369: uint32(10739),
	370: uint32(10747),
	371: uint32(10755),
	372: uint32(10763),
	373: uint32(10771),
	374: uint32(10779),
	375: uint32(10787),
	376: uint32(10795),
	377: uint32(10803),
	378: uint32(10811),
	379: uint32(10819),
	380: uint32(10827),
	381: uint32(10835),
	382: uint32(10843),
	383: uint32(10851),
	384: uint32(10859),
	385: uint32(10867),
	386: uint32(10875),
	387: uint32(10883),
	388: uint32(10891),
	389: uint32(10899),
	390: uint32(10907),
	391: uint32(10915),
	392: uint32(10923),
	393: uint32(10931),
	394: uint32(10939),
	395: uint32(10947),
	396: uint32(10955),
	397: uint32(10963),
	398: uint32(10971),
	399: uint32(10979),
	400: uint32(10987),
	401: uint32(10995),
	402: uint32(11003),
	403: uint32(11011),
	404: uint32(11019),
	405: uint32(11027),
	406: uint32(11035),
	407: uint32(11043),
	408: uint32(11051),
	409: uint32(11059),
	410: uint32(11067),
}

var ts_parse_actions = [1208]TSParseActionEntry{
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
		Fstate: uint16(libc.Int32FromInt32(88)),
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
		Fstate: uint16(libc.Int32FromInt32(332)),
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
		Fstate: uint16(libc.Int32FromInt32(231)),
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
		Fstate: uint16(libc.Int32FromInt32(350)),
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
		Fstate: uint16(libc.Int32FromInt32(265)),
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
		Fstate: uint16(libc.Int32FromInt32(399)),
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
		Fstate: uint16(libc.Int32FromInt32(361)),
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
		Fstate: uint16(libc.Int32FromInt32(440)),
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
		Fstate: uint16(libc.Int32FromInt32(356)),
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
		Fstate: uint16(libc.Int32FromInt32(267)),
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
		Fstate: uint16(libc.Int32FromInt32(338)),
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
		Fstate: uint16(libc.Int32FromInt32(313)),
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
		Fstate: uint16(libc.Int32FromInt32(400)),
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
		Fstate: uint16(libc.Int32FromInt32(95)),
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
		Fstate: uint16(libc.Int32FromInt32(23)),
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
		Fstate: uint16(libc.Int32FromInt32(101)),
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
		Fstate: uint16(libc.Int32FromInt32(93)),
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
		Fstate: uint16(libc.Int32FromInt32(92)),
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
		Fstate: uint16(libc.Int32FromInt32(91)),
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
		Fstate: uint16(libc.Int32FromInt32(176)),
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
		Fstate: uint16(libc.Int32FromInt32(346)),
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
		Fstate: uint16(libc.Int32FromInt32(351)),
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
		Fstate: uint16(libc.Int32FromInt32(355)),
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
		Fstate: uint16(libc.Int32FromInt32(366)),
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
		Fstate: uint16(libc.Int32FromInt32(283)),
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
		Fstate: uint16(libc.Int32FromInt32(395)),
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
		Fstate: uint16(libc.Int32FromInt32(24)),
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
		Fstate: uint16(libc.Int32FromInt32(24)),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_argument_list_expression_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(95)),
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
		Fsymbol:      uint16(aux_sym_argument_list_expression_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(23)),
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
		Fsymbol:      uint16(aux_sym_argument_list_expression_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(93)),
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
		Fsymbol:      uint16(aux_sym_argument_list_expression_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(92)),
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
		Fsymbol:      uint16(aux_sym_argument_list_expression_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(91)),
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
		Fsymbol:      uint16(aux_sym_argument_list_expression_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(176)),
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
		Fsymbol:      uint16(aux_sym_argument_list_expression_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(346)),
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
		Fsymbol:      uint16(aux_sym_argument_list_expression_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(351)),
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
		Fsymbol:      uint16(aux_sym_argument_list_expression_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(355)),
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
		Fsymbol:      uint16(aux_sym_argument_list_expression_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(366)),
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
		Fsymbol:      uint16(aux_sym_argument_list_expression_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(283)),
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
		Fsymbol:      uint16(aux_sym_argument_list_expression_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(395)),
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
		Fsymbol:      uint16(aux_sym_argument_list_expression_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(24)),
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
		Fsymbol:      uint16(aux_sym_argument_list_expression_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(24)),
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
		Fstate: uint16(libc.Int32FromInt32(34)),
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
		Fstate: uint16(libc.Int32FromInt32(35)),
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
		Fstate: uint16(libc.Int32FromInt32(180)),
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
		Fstate: uint16(libc.Int32FromInt32(6)),
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
		Fstate: uint16(libc.Int32FromInt32(297)),
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
		Fstate: uint16(libc.Int32FromInt32(229)),
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
		Fstate: uint16(libc.Int32FromInt32(427)),
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
		Fstate: uint16(libc.Int32FromInt32(267)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(239)),
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
		Fstate: uint16(libc.Int32FromInt32(271)),
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
		Fstate: uint16(libc.Int32FromInt32(376)),
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
		Fsymbol:      uint16(aux_sym_argument_list_expression_repeat1),
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
		Fstate: uint16(libc.Int32FromInt32(118)),
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
		Fsymbol:      uint16(aux_sym_const_expression_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(176)),
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
		Fsymbol:      uint16(aux_sym_const_expression_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(93)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_const_expression_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(92)),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_const_expression_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(91)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	154: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_const_expression_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(346)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_const_expression_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(351)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	160: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_const_expression_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(355)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_const_expression_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(366)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	166: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_const_expression_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(376)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(131)),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_compound_assignment_operator),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_compound_statement_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(175)),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_compound_statement_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(297)),
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
		Fsymbol:      uint16(aux_sym_compound_statement_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(229)),
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
		Fsymbol:      uint16(aux_sym_compound_statement_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(62)),
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
		Fsymbol:      uint16(aux_sym_compound_statement_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(427)),
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
		Fsymbol:      uint16(aux_sym_compound_statement_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(20)),
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
		Fsymbol:      uint16(aux_sym_compound_statement_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(21)),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(aux_sym_compound_statement_repeat1),
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
		Fcount: uint8(2),
	}})),
	200: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      uint16(libc.Int32FromInt32(436)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_compound_statement_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(437)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	206: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	207: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(libc.Int32FromInt32(22)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fsymbol:      uint16(aux_sym_compound_statement_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(441)),
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
		Fsymbol:      uint16(aux_sym_compound_statement_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(423)),
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
		Fsymbol:      uint16(aux_sym_compound_statement_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(4)),
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
		Fsymbol:      uint16(aux_sym_compound_statement_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(420)),
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
		Fsymbol:      uint16(aux_sym_compound_statement_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(267)),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_compound_statement_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(239)),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_compound_statement_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(338)),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_compound_statement_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(370)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(175)),
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
		Fstate: uint16(libc.Int32FromInt32(62)),
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
		Fstate: uint16(libc.Int32FromInt32(291)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(20)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(21)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(409)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(436)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(437)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(22)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(441)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(423)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(4)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(420)),
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
		Fstate: uint16(libc.Int32FromInt32(370)),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(aux_sym_global_variable_declaration_repeat1),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_global_variable_declaration_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(356)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_const_expression_repeat1),
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
		Fstate: uint16(libc.Int32FromInt32(275)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_const_expression_repeat1),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(114)),
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
		Fstate: uint16(libc.Int32FromInt32(335)),
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
		Fstate: uint16(libc.Int32FromInt32(117)),
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
		Fstate: uint16(libc.Int32FromInt32(67)),
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
		Fstate: uint16(libc.Int32FromInt32(121)),
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
		Fstate: uint16(libc.Int32FromInt32(286)),
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
		Fstate: uint16(libc.Int32FromInt32(123)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(287)),
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
		Fstate: uint16(libc.Int32FromInt32(413)),
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
		Fstate: uint16(libc.Int32FromInt32(319)),
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
		Fstate: uint16(libc.Int32FromInt32(221)),
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
		Fstate: uint16(libc.Int32FromInt32(391)),
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
		Fstate: uint16(libc.Int32FromInt32(75)),
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
		Fstate: uint16(libc.Int32FromInt32(80)),
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
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_attribute),
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
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_attribute),
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
		Fsymbol:      uint16(sym_preproc_else_in_statement),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	318: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_preproc_else_in_statement),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_compound_statement),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_compound_statement),
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
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_preproc_import),
		Fproduction_id: uint16(4),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_preproc_import),
		Fproduction_id: uint16(4),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_preproc_import),
		Fproduction_id: uint16(1),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_preproc_import),
		Fproduction_id: uint16(1),
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
		Fsymbol:        uint16(sym_preproc_import),
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
		Fcount: uint8(1),
	}})),
	334: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_preproc_import),
		Fproduction_id: uint16(1),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_preproc_import),
		Fproduction_id: uint16(1),
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
		Fcount: uint8(1),
	}})),
	338: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_preproc_import),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_compound_statement),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_compound_statement),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_source_file),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(84)),
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
		Fstate: uint16(libc.Int32FromInt32(153)),
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
		Fstate: uint16(libc.Int32FromInt32(85)),
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
		Fstate: uint16(libc.Int32FromInt32(82)),
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
		Fstate: uint16(libc.Int32FromInt32(158)),
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
		Fsymbol:      uint16(aux_sym_source_file_repeat2),
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
		Fstate:      uint16(libc.Int32FromInt32(84)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_source_file_repeat2),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(332)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_source_file_repeat2),
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
		Fstate:      uint16(libc.Int32FromInt32(231)),
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
		Fsymbol:      uint16(aux_sym_source_file_repeat2),
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
		Fstate:      uint16(libc.Int32FromInt32(350)),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_source_file_repeat2),
	})))),
	373: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(libc.Int32FromInt32(265)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	374: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	375: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      uint16(libc.Int32FromInt32(399)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_source_file_repeat2),
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
		Fstate:      uint16(libc.Int32FromInt32(361)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_source_file_repeat2),
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
		Fstate:      uint16(libc.Int32FromInt32(356)),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_source_file_repeat2),
	})))),
	385: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(libc.Int32FromInt32(267)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	387: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      uint16(libc.Int32FromInt32(338)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_source_file_repeat2),
	})))),
	391: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(libc.Int32FromInt32(313)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fsymbol:      uint16(aux_sym_source_file_repeat2),
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
		Fstate:      uint16(libc.Int32FromInt32(400)),
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
		Fstate: uint16(libc.Int32FromInt32(87)),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_preproc_else),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(sym_preproc_else),
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
		Fsymbol:      uint16(sym_source_file),
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
		Fcount: uint8(1),
	}})),
	408: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fcount: uint8(1),
	}})),
	414: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fsymbol:      uint16(sym_float_literal),
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
		Fcount: uint8(1),
	}})),
	418: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fcount: uint8(1),
	}})),
	422: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fproduction_id: uint16(22),
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
		Fproduction_id: uint16(22),
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
		Fstate: uint16(libc.Int32FromInt32(244)),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__expression),
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
		Fsymbol:      uint16(sym_type_constructor_or_function_call_expression),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_type_constructor_or_function_call_expression),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_binary_expression),
		Fproduction_id: uint16(10),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_binary_expression),
		Fproduction_id: uint16(10),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(371)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	448: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_subscript_expression),
		Fproduction_id: uint16(16),
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
		Fcount: uint8(1),
	}})),
	450: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_subscript_expression),
		Fproduction_id: uint16(16),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_composite_value_decomposition_expression),
		Fproduction_id: uint16(11),
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
		Fcount: uint8(1),
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
		Fsymbol:        uint16(sym_composite_value_decomposition_expression),
		Fproduction_id: uint16(11),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(sym_argument_list_expression),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_break_statement),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_break_statement),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_continue_statement),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_continue_statement),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_discard_statement),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_discard_statement),
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
		Fsymbol:      uint16(sym__statement),
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
		Fsymbol:      uint16(sym__statement),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	480: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(14)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	506: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_unary_expression),
		Fproduction_id: uint16(6),
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
		Fcount: uint8(1),
	}})),
	508: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_unary_expression),
		Fproduction_id: uint16(6),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	510: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fcount: uint8(1),
	}})),
	512: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fsymbol:      uint16(sym_loop_statement),
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
		Fsymbol:      uint16(sym_loop_statement),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_while_statement),
		Fproduction_id: uint16(23),
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
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_while_statement),
		Fproduction_id: uint16(23),
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
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_preproc_ifdef_in_statement),
		Fproduction_id: uint16(2),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_preproc_ifdef_in_statement),
		Fproduction_id: uint16(2),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_argument_list_expression),
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
		Fcount: uint8(1),
	}})),
	528: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_loop_statement),
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
		Fcount: uint8(1),
	}})),
	534: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_preproc_ifdef_in_statement),
		Fproduction_id: uint16(5),
	})))),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_preproc_ifdef_in_statement),
		Fproduction_id: uint16(5),
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
		Fcount: uint8(1),
	}})),
	538: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_preproc_ifdef_in_statement),
		Fproduction_id: uint16(2),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_preproc_ifdef_in_statement),
		Fproduction_id: uint16(2),
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
		Fcount: uint8(1),
	}})),
	542: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_else_statement),
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
		Fcount: uint8(1),
	}})),
	546: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_if_statement),
		Fproduction_id: uint16(25),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_if_statement),
		Fproduction_id: uint16(25),
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
		Fcount: uint8(1),
	}})),
	550: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fcount: uint8(1),
	}})),
	554: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fcount: uint8(1),
	}})),
	558: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_for_statement),
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
		Fcount: uint8(1),
	}})),
	562: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_preproc_ifdef_in_statement),
		Fproduction_id: uint16(8),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_preproc_ifdef_in_statement),
		Fproduction_id: uint16(8),
	})))),
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
		Fcount: uint8(1),
	}})),
	568: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: uint16(libc.Int32FromInt32(40)),
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
		Fstate: uint16(libc.Int32FromInt32(99)),
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
		Fstate: uint16(libc.Int32FromInt32(9)),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_assignment_statement),
		Fproduction_id: uint16(10),
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
		Fstate: uint16(libc.Int32FromInt32(37)),
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
		Fstate: uint16(libc.Int32FromInt32(177)),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_variable_statement),
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
		Fstate: uint16(libc.Int32FromInt32(106)),
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
		Fstate: uint16(libc.Int32FromInt32(259)),
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
		Fstate: uint16(libc.Int32FromInt32(98)),
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
		Fstate: uint16(libc.Int32FromInt32(43)),
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
		Fstate: uint16(libc.Int32FromInt32(36)),
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
		Fstate: uint16(libc.Int32FromInt32(429)),
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
		Fstate: uint16(libc.Int32FromInt32(188)),
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
		Fsymbol:      uint16(aux_sym_source_file_repeat1),
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
		Fsymbol:      uint16(aux_sym_source_file_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(440)),
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
		Fstate: uint16(libc.Int32FromInt32(359)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_preproc_ifdef),
		Fproduction_id: uint16(2),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_function_declaration),
		Fproduction_id: uint16(18),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(9),
		Fsymbol:        uint16(sym_function_declaration),
		Fproduction_id: uint16(24),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	619: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_function_declaration),
		Fproduction_id: uint16(20),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_function_declaration),
		Fproduction_id: uint16(21),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	623: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_preproc_ifdef),
		Fproduction_id: uint16(2),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	625: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	626: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	627: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_struct_declaration),
		Fproduction_id: uint16(2),
	})))),
	628: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	629: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_function_declaration),
		Fproduction_id: uint16(15),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	631: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_preproc_ifdef),
		Fproduction_id: uint16(5),
	})))),
	632: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	633: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_function_declaration),
		Fproduction_id: uint16(19),
	})))),
	634: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	635: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_function_declaration),
		Fproduction_id: uint16(14),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	637: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_function_declaration),
		Fproduction_id: uint16(17),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_function_declaration),
		Fproduction_id: uint16(12),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	641: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_function_declaration),
		Fproduction_id: uint16(7),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	645: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_function_declaration),
		Fproduction_id: uint16(13),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_preproc_ifdef),
		Fproduction_id: uint16(8),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	649: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_define_import_path),
		Fproduction_id: uint16(1),
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
		Fsymbol:      uint16(sym_enable_directive),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	653: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_type_declaration),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_lhs_expression),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	659: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(392)),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	663: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(439)),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_postfix_expression),
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
		Fstate: uint16(libc.Int32FromInt32(419)),
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
		Fstate: uint16(libc.Int32FromInt32(434)),
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
		Fstate: uint16(libc.Int32FromInt32(212)),
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
		Fstate: uint16(libc.Int32FromInt32(195)),
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
		Fchild_count: uint8(8),
		Fsymbol:      uint16(sym_type_declaration),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_postfix_expression),
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
		Fstate: uint16(libc.Int32FromInt32(216)),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_preproc_ifdef_in_struct_declaration_repeat1),
	})))),
	692: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(libc.Int32FromInt32(419)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	693: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	694: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_preproc_ifdef_in_struct_declaration_repeat1),
	})))),
	695: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(libc.Int32FromInt32(356)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_preproc_ifdef_in_struct_declaration_repeat1),
	})))),
	698: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(libc.Int32FromInt32(338)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	700: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_preproc_ifdef_in_struct_declaration_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(434)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_preproc_ifdef_in_struct_declaration_repeat1),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_preproc_else_in_struct_declaration),
	})))),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_preproc_else_in_struct_declaration),
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
		Fsymbol:        uint16(sym_variable_identifier_declaration),
		Fproduction_id: uint16(3),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(44)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	723: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__struct_declaration_content_repeat1),
	})))),
	724: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(libc.Int32FromInt32(419)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__struct_declaration_content_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(356)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	729: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__struct_declaration_content_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(338)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__struct_declaration_content_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(434)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	735: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: uint16(libc.Int32FromInt32(91)),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_case_selectors),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_struct_member),
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
		Fstate: uint16(libc.Int32FromInt32(252)),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_struct_member),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_preproc_ifdef_in_struct_declaration),
		Fproduction_id: uint16(8),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_preproc_ifdef_in_struct_declaration),
		Fproduction_id: uint16(2),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(249)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_preproc_ifdef_in_struct_declaration),
		Fproduction_id: uint16(5),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_preproc_ifdef_in_struct_declaration),
		Fproduction_id: uint16(2),
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
		Fstate: uint16(libc.Int32FromInt32(266)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_preproc_ifdef_in_struct_declaration_repeat1),
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
		Fstate: uint16(libc.Int32FromInt32(223)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	763: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      uint16(libc.Int32FromInt32(419)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	766: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      uint16(libc.Int32FromInt32(356)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(333)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(333)),
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
		Fstate: uint16(libc.Int32FromInt32(340)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(340)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym__struct_declaration_content_repeat1),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(416)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__struct_declaration_content),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(344)),
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
		Fstate: uint16(libc.Int32FromInt32(430)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	787: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      uint16(libc.Int32FromInt32(363)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	790: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      uint16(libc.Int32FromInt32(363)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	793: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      uint16(libc.Int32FromInt32(92)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(352)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__struct_declaration_content),
	})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__struct_declaration_content_repeat1),
	})))),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__struct_declaration_content),
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
		Fcount: uint8(1),
	}})),
	806: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(314)),
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
		Fstate: uint16(libc.Int32FromInt32(337)),
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
		Fcount: uint8(1),
	}})),
	812: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_import_path_repeat1),
	})))),
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
		Fcount: uint8(1),
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
		Fextra: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_import_path_repeat1),
	})))),
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
		Fcount: uint8(2),
	}})),
	818: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_import_path_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(365)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(179)),
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
		Fstate: uint16(libc.Int32FromInt32(232)),
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
		Fstate: uint16(libc.Int32FromInt32(242)),
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
		Fstate: uint16(libc.Int32FromInt32(66)),
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
		Fcount: uint8(1),
	}})),
	831: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_import_path),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_import_path),
	})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(365)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	841: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      uint16(libc.Int32FromInt32(242)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(20)),
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
		Fstate: uint16(libc.Int32FromInt32(126)),
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
		Fstate: uint16(libc.Int32FromInt32(213)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(273)),
	}})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	854: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      uint16(libc.Int32FromInt32(213)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	857: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      uint16(libc.Int32FromInt32(273)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
	862: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_import_path),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_import_path),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(48)),
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
		Fstate: uint16(libc.Int32FromInt32(358)),
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
		Fcount: uint8(1),
	}})),
	870: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(77)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(388)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(347)),
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
		Fstate: uint16(libc.Int32FromInt32(347)),
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
		Fstate: uint16(libc.Int32FromInt32(318)),
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
		Fstate: uint16(libc.Int32FromInt32(431)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_case_selectors),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(235)),
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
		Fstate: uint16(libc.Int32FromInt32(207)),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_case_selectors_repeat1),
	})))),
	895: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(libc.Int32FromInt32(218)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	897: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	899: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	901: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_const_expression),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	905: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(315)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	909: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	911: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	913: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(30)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(72)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(326)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	921: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(19)),
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
		Fcount:    uint8(1),
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
		Fsymbol:      uint16(sym_switch_body),
	})))),
	928: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	929: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	930: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	931: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	932: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	933: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(2)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(69)),
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
		Fstate: uint16(libc.Int32FromInt32(428)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_case_compound_statement),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	949: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fcount: uint8(2),
	}})),
	951: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_preproc_import_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(428)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_preproc_import_repeat1),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(438)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	962: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_import_path_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(438)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(79)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(299)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(304)),
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
		Fstate: uint16(libc.Int32FromInt32(173)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(51)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	977: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	978: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	979: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(385)),
	}})))),
	980: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	981: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	982: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	983: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	985: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	986: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	987: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	988: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	989: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	990: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	991: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	992: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	993: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	994: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	995: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	996: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	997: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	998: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	999: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(394)),
	}})))),
	1000: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1001: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1002: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(31)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(279)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1007: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: uint16(libc.Int32FromInt32(45)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1011: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: uint16(libc.Int32FromInt32(282)),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_variable_statement),
	})))),
	1016: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1017: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1018: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1019: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(334)),
	}})))),
	1020: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1021: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1022: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1023: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1024: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1025: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1026: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1027: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1028: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1029: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1030: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1031: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	1032: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1033: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1034: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1035: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1036: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1037: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	1038: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1039: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	1040: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	1041: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_preproc_import_repeat1),
	})))),
	1042: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1043: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1044: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1045: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1046: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1047: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1048: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1049: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	1050: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1051: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1052: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1053: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(364)),
	}})))),
	1054: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1055: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(418)),
	}})))),
	1056: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1057: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1058: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1059: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1060: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1061: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_alias),
	})))),
	1062: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(251)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(59)),
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
		Fstate: uint16(libc.Int32FromInt32(71)),
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
		Fstate: uint16(libc.Int32FromInt32(257)),
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
		Fstate: uint16(libc.Int32FromInt32(381)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(160)),
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
		Fstate: uint16(libc.Int32FromInt32(255)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1079: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(261)),
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
		Fstate: uint16(libc.Int32FromInt32(152)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1085: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1087: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: uint16(libc.Int32FromInt32(127)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(53)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(100)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(33)),
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
		Fstate: uint16(libc.Int32FromInt32(331)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1105: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1106: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1107: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1108: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1109: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1110: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1111: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1112: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1113: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1114: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1115: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1116: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(354)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1119: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1123: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: uint16(libc.Int32FromInt32(186)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1127: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(73)),
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
		Fstate: uint16(libc.Int32FromInt32(122)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(105)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(435)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1137: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1139: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1140: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1141: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f uint8_t
		_ [7]byte
	}{f: uint8(TSParseActionTypeAccept)})),
	1142: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(250)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(209)),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_global_variable_declaration),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1149: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: uint16(libc.Int32FromInt32(162)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(215)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(421)),
	}})))),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_type_alias_declaration),
	})))),
	1158: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1159: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(263)),
	}})))),
	1160: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1161: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	1162: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1163: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(425)),
	}})))),
	1164: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1165: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1166: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1167: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1168: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1169: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1170: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1171: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	1172: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1173: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1174: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1175: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1176: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1177: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	1178: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1179: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	1180: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1181: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(343)),
	}})))),
	1182: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1183: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: uint16(libc.Int32FromInt32(404)),
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
		Fstate: uint16(libc.Int32FromInt32(401)),
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
		Fstate: uint16(libc.Int32FromInt32(7)),
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
		Fstate: uint16(libc.Int32FromInt32(309)),
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
		Fstate: uint16(libc.Int32FromInt32(183)),
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
		Fstate: uint16(libc.Int32FromInt32(214)),
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
		Fstate: uint16(libc.Int32FromInt32(52)),
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
		Fstate: uint16(libc.Int32FromInt32(32)),
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
		Fstate: uint16(libc.Int32FromInt32(339)),
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
		Fstate: uint16(libc.Int32FromInt32(187)),
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
		Fstate: uint16(libc.Int32FromInt32(383)),
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
		Fstate: uint16(libc.Int32FromInt32(171)),
	}})))),
}

type ts_external_scanner_symbol_identifiers = int32

const ts_external_token_block_comment = 0

var ts_external_scanner_symbol_map = [1]TSSymbol{
	0: uint16(sym_block_comment),
}

var ts_external_scanner_states = [2][1]uint8{
	1: {
		0: libc.BoolUint8(true1 != 0),
	},
}

func tree_sitter_wgsl_bevy(tls *libc.TLS) (r uintptr) {
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
	*(*uintptr)(unsafe.Add(p, 144)) = __ccgo_fp(ts_lex)
	*(*uintptr)(unsafe.Add(p, 152)) = __ccgo_fp(ts_lex_keywords)
	*(*uintptr)(unsafe.Add(p, 184)) = __ccgo_fp(tree_sitter_wgsl_bevy_external_scanner_create)
	*(*uintptr)(unsafe.Add(p, 192)) = __ccgo_fp(tree_sitter_wgsl_bevy_external_scanner_destroy)
	*(*uintptr)(unsafe.Add(p, 200)) = __ccgo_fp(tree_sitter_wgsl_bevy_external_scanner_scan)
	*(*uintptr)(unsafe.Add(p, 208)) = __ccgo_fp(tree_sitter_wgsl_bevy_external_scanner_serialize)
	*(*uintptr)(unsafe.Add(p, 216)) = __ccgo_fp(tree_sitter_wgsl_bevy_external_scanner_deserialize)
}

func __ccgo_fp(f interface{}) uintptr {
	type iface [2]uintptr
	return (*iface)(unsafe.Pointer(&f))[1]
}

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = "end\x00identifier\x00line_comment\x00;\x00=\x00let\x00override\x00type\x00(\x00,\x00)\x00virtual\x00fn\x00->\x00struct\x00{\x00}\x00enable\x00@\x00_\x00+=\x00-=\x00*=\x00/=\x00%=\x00&=\x00|=\x00^=\x00if\x00else\x00switch\x00case\x00:\x00default\x00fallthrough\x00loop\x00for\x00while\x00break\x00continue\x00continuing\x00return\x00discard\x00var\x00<\x00>\x00++\x00--\x00int_literal\x00float_literal_token1\x00float_literal_token2\x00true\x00false\x00bool\x00u32\x00i32\x00f32\x00f16\x00array\x00ptr\x00sampler\x00sampler_comparison\x00texture_depth_2d\x00texture_depth_2d_array\x00texture_depth_cube\x00texture_depth_cube_array\x00texture_depth_multisampled_2d\x00texture_1d\x00texture_2d\x00texture_2d_array\x00texture_3d\x00texture_cube\x00texture_cube_array\x00texture_multisampled_2d\x00texture_storage_1d\x00texture_storage_2d\x00texture_storage_2d_array\x00texture_storage_3d\x00vec2\x00vec3\x00vec4\x00mat2x2\x00mat2x3\x00mat2x4\x00mat3x2\x00mat3x3\x00mat3x4\x00mat4x2\x00mat4x3\x00mat4x4\x00rgba8unorm\x00rgba8snorm\x00rgba8uint\x00rgba8sint\x00rgba16uint\x00rgba16sint\x00rgba16float\x00r32uint\x00r32sint\x00r32float\x00rg32uint\x00rg32sint\x00rg32float\x00rgba32uint\x00rgba32sint\x00rgba32float\x00function\x00private\x00workgroup\x00uniform\x00storage\x00read\x00write\x00read_write\x00bitcast\x00||\x00&&\x00|\x00^\x00&\x00==\x00!=\x00<=\x00>=\x00<<\x00>>\x00+\x00-\x00*\x00/\x00%\x00!\x00~\x00[\x00]\x00.\x00#import\x00\n\x00#define_import_path\x00::\x00as\x00#ifdef\x00#ifndef\x00#endif\x00#else\x00block_comment\x00source_file\x00_declaration\x00global_variable_declaration\x00global_constant_declaration\x00type_alias_declaration\x00const_expression\x00function_declaration\x00function_return_type_declaration\x00struct_declaration\x00struct_member\x00enable_directive\x00attribute\x00_literal_or_identifier\x00parameter_list\x00parameter\x00_statement\x00compound_statement\x00assignment_statement\x00compound_assignment_operator\x00if_statement\x00else_statement\x00switch_statement\x00switch_body\x00case_selectors\x00case_compound_statement\x00fallthrough_statement\x00loop_statement\x00for_statement\x00for_header\x00while_statement\x00break_statement\x00break_if_statement\x00continue_statement\x00continuing_statement\x00continuing_compound_statement\x00return_statement\x00discard_statement\x00variable_statement\x00variable_declaration\x00variable_qualifier\x00variable_identifier_declaration\x00increment_statement\x00decrement_statement\x00_expression\x00const_literal\x00float_literal\x00bool_literal\x00parenthesized_expression\x00type_constructor_or_function_call_expression\x00type_declaration\x00_vec_prefix\x00_mat_prefix\x00texel_format\x00address_space\x00access_mode\x00argument_list_expression\x00bitcast_expression\x00binary_expression\x00unary_expression\x00postfix_expression\x00subscript_expression\x00lhs_expression\x00composite_value_decomposition_expression\x00_struct_declaration_content\x00preproc_import\x00define_import_path\x00import_path\x00alias\x00preproc_ifdef\x00preproc_else\x00source_file_repeat1\x00source_file_repeat2\x00global_variable_declaration_repeat1\x00const_expression_repeat1\x00attribute_repeat1\x00parameter_list_repeat1\x00compound_statement_repeat1\x00switch_statement_repeat1\x00case_selectors_repeat1\x00argument_list_expression_repeat1\x00lhs_expression_repeat1\x00_struct_declaration_content_repeat1\x00preproc_import_repeat1\x00import_path_repeat1\x00preproc_ifdef_in_struct_declaration_repeat1\x00accessor\x00alternative\x00argument\x00body\x00condition\x00consequence\x00left\x00name\x00parameters\x00path\x00right\x00subscript\x00value\x00"
