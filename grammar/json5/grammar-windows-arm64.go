// Code generated for windows/arm64 by 'ccgo -ignore-unsupported-alignment -ignore-unsupported-atomic-sizes -ignore-vector-functions --winapi-no-errno -ignore-link-errors -std=gnu11 -O0 -D__extension__= -D__forceinline=static inline -D__attribute__(...)= -D__declspec(x)= -D__cdecl= -D__stdcall= -D__fastcall= -D__thiscall= -D_cdecl= -D__restrict= -D__restrict__= -D__MINGW_EXTENSION= -D_X86INTRIN_H_INCLUDED -D_X86GPRINTRIN_H_INCLUDED -D_IMMINTRIN_H_INCLUDED -D_MMINTRIN_H_INCLUDED -D_XMMINTRIN_H_INCLUDED -D_EMMINTRIN_H_INCLUDED -D_PMMINTRIN_H_INCLUDED -D_MM3DNOW_H_INCLUDED -Dfunc=func_token -Dinterface=interface_token -Dselect=select_token -Dchan=chan_token -Dgo=go_token -Dmap=map_token -Dpackage=package_token -Dtype=type_token -Dvar=var_token -Dimport=import_token -Ddefer=defer_token -Dfallthrough=fallthrough_token -Drange=range_token -I C:\a\ccgo-tree-sitter\ccgo-tree-sitter\third-party\tree-sitter-json5\src -I C:\a\ccgo-tree-sitter\ccgo-tree-sitter\third-party\tree-sitter-json5 -I C:\a\ccgo-tree-sitter\ccgo-tree-sitter\third-party -I C:\Users\runneradmin\.cache\workspaced\sources\github\7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2\lib\include -I C:\Users\runneradmin\.cache\workspaced\sources\github\7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2\lib\src C:\a\ccgo-tree-sitter\ccgo-tree-sitter\third-party\tree-sitter-json5\src\parser.c -o grammar.go', DO NOT EDIT.

//go:build windows && arm64

package grammar_json5

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
const FIELD_COUNT = 2
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
const LANGUAGE_VERSION = 15
const LARGE_STATE_COUNT = 7
const MAX_ALIAS_SEQUENCE_LENGTH = 5
const MAX_RESERVED_WORD_SET_SIZE = 0
const MB_LEN_MAX = 1
const MINGW_HAS_DDK_H = 1
const MINGW_HAS_SECURE_API = 1
const PRODUCTION_ID_COUNT = 2
const PTRDIFF_MAX = "__PTRDIFF_MAX__"
const RAND_MAX = 0x7fff
const SIZE_MAX = "__SIZE_MAX__"
const STATE_COUNT = 31
const SUPERTYPE_COUNT = 0
const SYMBOL_COUNT = 21
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

const sym_comment = 1
const anon_sym_LBRACE = 2
const anon_sym_COMMA = 3
const anon_sym_RBRACE = 4
const anon_sym_COLON = 5
const sym_identifier = 6
const anon_sym_LBRACK = 7
const anon_sym_RBRACK = 8
const sym_string = 9
const sym_number = 10
const sym_null = 11
const sym_true = 12
const sym_false = 13
const sym_file = 14
const sym_object = 15
const sym_member = 16
const sym_array = 17
const sym__value = 18
const aux_sym_object_repeat1 = 19
const aux_sym_array_repeat1 = 20

var ts_symbol_names = [21]uintptr{
	0:  __ccgo_ts,
	1:  __ccgo_ts + 4,
	2:  __ccgo_ts + 12,
	3:  __ccgo_ts + 14,
	4:  __ccgo_ts + 16,
	5:  __ccgo_ts + 18,
	6:  __ccgo_ts + 20,
	7:  __ccgo_ts + 31,
	8:  __ccgo_ts + 33,
	9:  __ccgo_ts + 35,
	10: __ccgo_ts + 42,
	11: __ccgo_ts + 49,
	12: __ccgo_ts + 54,
	13: __ccgo_ts + 59,
	14: __ccgo_ts + 65,
	15: __ccgo_ts + 70,
	16: __ccgo_ts + 77,
	17: __ccgo_ts + 84,
	18: __ccgo_ts + 90,
	19: __ccgo_ts + 97,
	20: __ccgo_ts + 112,
}

var ts_symbol_map = [21]TSSymbol{
	1:  uint16(sym_comment),
	2:  uint16(anon_sym_LBRACE),
	3:  uint16(anon_sym_COMMA),
	4:  uint16(anon_sym_RBRACE),
	5:  uint16(anon_sym_COLON),
	6:  uint16(sym_identifier),
	7:  uint16(anon_sym_LBRACK),
	8:  uint16(anon_sym_RBRACK),
	9:  uint16(sym_string),
	10: uint16(sym_number),
	11: uint16(sym_null),
	12: uint16(sym_true),
	13: uint16(sym_false),
	14: uint16(sym_file),
	15: uint16(sym_object),
	16: uint16(sym_member),
	17: uint16(sym_array),
	18: uint16(sym__value),
	19: uint16(aux_sym_object_repeat1),
	20: uint16(aux_sym_array_repeat1),
}

var ts_symbol_metadata = [21]TSSymbolMetadata{
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
		Fnamed:   libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	19: {},
	20: {},
}

type ts_field_identifiers = int32

const field_name = 1
const field_value = 2

var ts_field_names = [3]uintptr{
	0: libc.UintptrFromInt32(0),
	1: __ccgo_ts + 126,
	2: __ccgo_ts + 131,
}

var ts_field_map_slices = [2]TSMapSlice{
	1: {
		Flength: uint16(2),
	},
}

var ts_field_map_entries = [2]TSFieldMapEntry{
	0: {
		Ffield_id: uint16(field_name),
	},
	1: {
		Ffield_id:    uint16(field_value),
		Fchild_index: uint8(2),
	},
}

var ts_alias_sequences = [2][5]TSSymbol{}

var ts_non_terminal_alias_map = [1]uint16_t{}

var ts_primary_state_ids = [31]TSStateId{
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
}

var sym_identifier_character_set_1 = [679]TSCharacterRange{
	0: {
		Fstart: int32('$'),
		Fend:   int32('$'),
	},
	1: {
		Fstart: int32('A'),
		Fend:   int32('Z'),
	},
	2: {
		Fstart: int32('_'),
		Fend:   int32('_'),
	},
	3: {
		Fstart: int32('a'),
		Fend:   int32('z'),
	},
	4: {
		Fstart: int32(0xaa),
		Fend:   int32(0xaa),
	},
	5: {
		Fstart: int32(0xb5),
		Fend:   int32(0xb5),
	},
	6: {
		Fstart: int32(0xba),
		Fend:   int32(0xba),
	},
	7: {
		Fstart: int32(0xc0),
		Fend:   int32(0xd6),
	},
	8: {
		Fstart: int32(0xd8),
		Fend:   int32(0xf6),
	},
	9: {
		Fstart: int32(0xf8),
		Fend:   int32(0x2c1),
	},
	10: {
		Fstart: int32(0x2c6),
		Fend:   int32(0x2d1),
	},
	11: {
		Fstart: int32(0x2e0),
		Fend:   int32(0x2e4),
	},
	12: {
		Fstart: int32(0x2ec),
		Fend:   int32(0x2ec),
	},
	13: {
		Fstart: int32(0x2ee),
		Fend:   int32(0x2ee),
	},
	14: {
		Fstart: int32(0x370),
		Fend:   int32(0x374),
	},
	15: {
		Fstart: int32(0x376),
		Fend:   int32(0x377),
	},
	16: {
		Fstart: int32(0x37a),
		Fend:   int32(0x37d),
	},
	17: {
		Fstart: int32(0x37f),
		Fend:   int32(0x37f),
	},
	18: {
		Fstart: int32(0x386),
		Fend:   int32(0x386),
	},
	19: {
		Fstart: int32(0x388),
		Fend:   int32(0x38a),
	},
	20: {
		Fstart: int32(0x38c),
		Fend:   int32(0x38c),
	},
	21: {
		Fstart: int32(0x38e),
		Fend:   int32(0x3a1),
	},
	22: {
		Fstart: int32(0x3a3),
		Fend:   int32(0x3f5),
	},
	23: {
		Fstart: int32(0x3f7),
		Fend:   int32(0x481),
	},
	24: {
		Fstart: int32(0x48a),
		Fend:   int32(0x52f),
	},
	25: {
		Fstart: int32(0x531),
		Fend:   int32(0x556),
	},
	26: {
		Fstart: int32(0x559),
		Fend:   int32(0x559),
	},
	27: {
		Fstart: int32(0x560),
		Fend:   int32(0x588),
	},
	28: {
		Fstart: int32(0x5d0),
		Fend:   int32(0x5ea),
	},
	29: {
		Fstart: int32(0x5ef),
		Fend:   int32(0x5f2),
	},
	30: {
		Fstart: int32(0x620),
		Fend:   int32(0x64a),
	},
	31: {
		Fstart: int32(0x66e),
		Fend:   int32(0x66f),
	},
	32: {
		Fstart: int32(0x671),
		Fend:   int32(0x6d3),
	},
	33: {
		Fstart: int32(0x6d5),
		Fend:   int32(0x6d5),
	},
	34: {
		Fstart: int32(0x6e5),
		Fend:   int32(0x6e6),
	},
	35: {
		Fstart: int32(0x6ee),
		Fend:   int32(0x6ef),
	},
	36: {
		Fstart: int32(0x6fa),
		Fend:   int32(0x6fc),
	},
	37: {
		Fstart: int32(0x6ff),
		Fend:   int32(0x6ff),
	},
	38: {
		Fstart: int32(0x710),
		Fend:   int32(0x710),
	},
	39: {
		Fstart: int32(0x712),
		Fend:   int32(0x72f),
	},
	40: {
		Fstart: int32(0x74d),
		Fend:   int32(0x7a5),
	},
	41: {
		Fstart: int32(0x7b1),
		Fend:   int32(0x7b1),
	},
	42: {
		Fstart: int32(0x7ca),
		Fend:   int32(0x7ea),
	},
	43: {
		Fstart: int32(0x7f4),
		Fend:   int32(0x7f5),
	},
	44: {
		Fstart: int32(0x7fa),
		Fend:   int32(0x7fa),
	},
	45: {
		Fstart: int32(0x800),
		Fend:   int32(0x815),
	},
	46: {
		Fstart: int32(0x81a),
		Fend:   int32(0x81a),
	},
	47: {
		Fstart: int32(0x824),
		Fend:   int32(0x824),
	},
	48: {
		Fstart: int32(0x828),
		Fend:   int32(0x828),
	},
	49: {
		Fstart: int32(0x840),
		Fend:   int32(0x858),
	},
	50: {
		Fstart: int32(0x860),
		Fend:   int32(0x86a),
	},
	51: {
		Fstart: int32(0x870),
		Fend:   int32(0x887),
	},
	52: {
		Fstart: int32(0x889),
		Fend:   int32(0x88e),
	},
	53: {
		Fstart: int32(0x8a0),
		Fend:   int32(0x8c9),
	},
	54: {
		Fstart: int32(0x904),
		Fend:   int32(0x939),
	},
	55: {
		Fstart: int32(0x93d),
		Fend:   int32(0x93d),
	},
	56: {
		Fstart: int32(0x950),
		Fend:   int32(0x950),
	},
	57: {
		Fstart: int32(0x958),
		Fend:   int32(0x961),
	},
	58: {
		Fstart: int32(0x971),
		Fend:   int32(0x980),
	},
	59: {
		Fstart: int32(0x985),
		Fend:   int32(0x98c),
	},
	60: {
		Fstart: int32(0x98f),
		Fend:   int32(0x990),
	},
	61: {
		Fstart: int32(0x993),
		Fend:   int32(0x9a8),
	},
	62: {
		Fstart: int32(0x9aa),
		Fend:   int32(0x9b0),
	},
	63: {
		Fstart: int32(0x9b2),
		Fend:   int32(0x9b2),
	},
	64: {
		Fstart: int32(0x9b6),
		Fend:   int32(0x9b9),
	},
	65: {
		Fstart: int32(0x9bd),
		Fend:   int32(0x9bd),
	},
	66: {
		Fstart: int32(0x9ce),
		Fend:   int32(0x9ce),
	},
	67: {
		Fstart: int32(0x9dc),
		Fend:   int32(0x9dd),
	},
	68: {
		Fstart: int32(0x9df),
		Fend:   int32(0x9e1),
	},
	69: {
		Fstart: int32(0x9f0),
		Fend:   int32(0x9f1),
	},
	70: {
		Fstart: int32(0x9fc),
		Fend:   int32(0x9fc),
	},
	71: {
		Fstart: int32(0xa05),
		Fend:   int32(0xa0a),
	},
	72: {
		Fstart: int32(0xa0f),
		Fend:   int32(0xa10),
	},
	73: {
		Fstart: int32(0xa13),
		Fend:   int32(0xa28),
	},
	74: {
		Fstart: int32(0xa2a),
		Fend:   int32(0xa30),
	},
	75: {
		Fstart: int32(0xa32),
		Fend:   int32(0xa33),
	},
	76: {
		Fstart: int32(0xa35),
		Fend:   int32(0xa36),
	},
	77: {
		Fstart: int32(0xa38),
		Fend:   int32(0xa39),
	},
	78: {
		Fstart: int32(0xa59),
		Fend:   int32(0xa5c),
	},
	79: {
		Fstart: int32(0xa5e),
		Fend:   int32(0xa5e),
	},
	80: {
		Fstart: int32(0xa72),
		Fend:   int32(0xa74),
	},
	81: {
		Fstart: int32(0xa85),
		Fend:   int32(0xa8d),
	},
	82: {
		Fstart: int32(0xa8f),
		Fend:   int32(0xa91),
	},
	83: {
		Fstart: int32(0xa93),
		Fend:   int32(0xaa8),
	},
	84: {
		Fstart: int32(0xaaa),
		Fend:   int32(0xab0),
	},
	85: {
		Fstart: int32(0xab2),
		Fend:   int32(0xab3),
	},
	86: {
		Fstart: int32(0xab5),
		Fend:   int32(0xab9),
	},
	87: {
		Fstart: int32(0xabd),
		Fend:   int32(0xabd),
	},
	88: {
		Fstart: int32(0xad0),
		Fend:   int32(0xad0),
	},
	89: {
		Fstart: int32(0xae0),
		Fend:   int32(0xae1),
	},
	90: {
		Fstart: int32(0xaf9),
		Fend:   int32(0xaf9),
	},
	91: {
		Fstart: int32(0xb05),
		Fend:   int32(0xb0c),
	},
	92: {
		Fstart: int32(0xb0f),
		Fend:   int32(0xb10),
	},
	93: {
		Fstart: int32(0xb13),
		Fend:   int32(0xb28),
	},
	94: {
		Fstart: int32(0xb2a),
		Fend:   int32(0xb30),
	},
	95: {
		Fstart: int32(0xb32),
		Fend:   int32(0xb33),
	},
	96: {
		Fstart: int32(0xb35),
		Fend:   int32(0xb39),
	},
	97: {
		Fstart: int32(0xb3d),
		Fend:   int32(0xb3d),
	},
	98: {
		Fstart: int32(0xb5c),
		Fend:   int32(0xb5d),
	},
	99: {
		Fstart: int32(0xb5f),
		Fend:   int32(0xb61),
	},
	100: {
		Fstart: int32(0xb71),
		Fend:   int32(0xb71),
	},
	101: {
		Fstart: int32(0xb83),
		Fend:   int32(0xb83),
	},
	102: {
		Fstart: int32(0xb85),
		Fend:   int32(0xb8a),
	},
	103: {
		Fstart: int32(0xb8e),
		Fend:   int32(0xb90),
	},
	104: {
		Fstart: int32(0xb92),
		Fend:   int32(0xb95),
	},
	105: {
		Fstart: int32(0xb99),
		Fend:   int32(0xb9a),
	},
	106: {
		Fstart: int32(0xb9c),
		Fend:   int32(0xb9c),
	},
	107: {
		Fstart: int32(0xb9e),
		Fend:   int32(0xb9f),
	},
	108: {
		Fstart: int32(0xba3),
		Fend:   int32(0xba4),
	},
	109: {
		Fstart: int32(0xba8),
		Fend:   int32(0xbaa),
	},
	110: {
		Fstart: int32(0xbae),
		Fend:   int32(0xbb9),
	},
	111: {
		Fstart: int32(0xbd0),
		Fend:   int32(0xbd0),
	},
	112: {
		Fstart: int32(0xc05),
		Fend:   int32(0xc0c),
	},
	113: {
		Fstart: int32(0xc0e),
		Fend:   int32(0xc10),
	},
	114: {
		Fstart: int32(0xc12),
		Fend:   int32(0xc28),
	},
	115: {
		Fstart: int32(0xc2a),
		Fend:   int32(0xc39),
	},
	116: {
		Fstart: int32(0xc3d),
		Fend:   int32(0xc3d),
	},
	117: {
		Fstart: int32(0xc58),
		Fend:   int32(0xc5a),
	},
	118: {
		Fstart: int32(0xc5d),
		Fend:   int32(0xc5d),
	},
	119: {
		Fstart: int32(0xc60),
		Fend:   int32(0xc61),
	},
	120: {
		Fstart: int32(0xc80),
		Fend:   int32(0xc80),
	},
	121: {
		Fstart: int32(0xc85),
		Fend:   int32(0xc8c),
	},
	122: {
		Fstart: int32(0xc8e),
		Fend:   int32(0xc90),
	},
	123: {
		Fstart: int32(0xc92),
		Fend:   int32(0xca8),
	},
	124: {
		Fstart: int32(0xcaa),
		Fend:   int32(0xcb3),
	},
	125: {
		Fstart: int32(0xcb5),
		Fend:   int32(0xcb9),
	},
	126: {
		Fstart: int32(0xcbd),
		Fend:   int32(0xcbd),
	},
	127: {
		Fstart: int32(0xcdd),
		Fend:   int32(0xcde),
	},
	128: {
		Fstart: int32(0xce0),
		Fend:   int32(0xce1),
	},
	129: {
		Fstart: int32(0xcf1),
		Fend:   int32(0xcf2),
	},
	130: {
		Fstart: int32(0xd04),
		Fend:   int32(0xd0c),
	},
	131: {
		Fstart: int32(0xd0e),
		Fend:   int32(0xd10),
	},
	132: {
		Fstart: int32(0xd12),
		Fend:   int32(0xd3a),
	},
	133: {
		Fstart: int32(0xd3d),
		Fend:   int32(0xd3d),
	},
	134: {
		Fstart: int32(0xd4e),
		Fend:   int32(0xd4e),
	},
	135: {
		Fstart: int32(0xd54),
		Fend:   int32(0xd56),
	},
	136: {
		Fstart: int32(0xd5f),
		Fend:   int32(0xd61),
	},
	137: {
		Fstart: int32(0xd7a),
		Fend:   int32(0xd7f),
	},
	138: {
		Fstart: int32(0xd85),
		Fend:   int32(0xd96),
	},
	139: {
		Fstart: int32(0xd9a),
		Fend:   int32(0xdb1),
	},
	140: {
		Fstart: int32(0xdb3),
		Fend:   int32(0xdbb),
	},
	141: {
		Fstart: int32(0xdbd),
		Fend:   int32(0xdbd),
	},
	142: {
		Fstart: int32(0xdc0),
		Fend:   int32(0xdc6),
	},
	143: {
		Fstart: int32(0xe01),
		Fend:   int32(0xe30),
	},
	144: {
		Fstart: int32(0xe32),
		Fend:   int32(0xe33),
	},
	145: {
		Fstart: int32(0xe40),
		Fend:   int32(0xe46),
	},
	146: {
		Fstart: int32(0xe81),
		Fend:   int32(0xe82),
	},
	147: {
		Fstart: int32(0xe84),
		Fend:   int32(0xe84),
	},
	148: {
		Fstart: int32(0xe86),
		Fend:   int32(0xe8a),
	},
	149: {
		Fstart: int32(0xe8c),
		Fend:   int32(0xea3),
	},
	150: {
		Fstart: int32(0xea5),
		Fend:   int32(0xea5),
	},
	151: {
		Fstart: int32(0xea7),
		Fend:   int32(0xeb0),
	},
	152: {
		Fstart: int32(0xeb2),
		Fend:   int32(0xeb3),
	},
	153: {
		Fstart: int32(0xebd),
		Fend:   int32(0xebd),
	},
	154: {
		Fstart: int32(0xec0),
		Fend:   int32(0xec4),
	},
	155: {
		Fstart: int32(0xec6),
		Fend:   int32(0xec6),
	},
	156: {
		Fstart: int32(0xedc),
		Fend:   int32(0xedf),
	},
	157: {
		Fstart: int32(0xf00),
		Fend:   int32(0xf00),
	},
	158: {
		Fstart: int32(0xf40),
		Fend:   int32(0xf47),
	},
	159: {
		Fstart: int32(0xf49),
		Fend:   int32(0xf6c),
	},
	160: {
		Fstart: int32(0xf88),
		Fend:   int32(0xf8c),
	},
	161: {
		Fstart: int32(0x1000),
		Fend:   int32(0x102a),
	},
	162: {
		Fstart: int32(0x103f),
		Fend:   int32(0x103f),
	},
	163: {
		Fstart: int32(0x1050),
		Fend:   int32(0x1055),
	},
	164: {
		Fstart: int32(0x105a),
		Fend:   int32(0x105d),
	},
	165: {
		Fstart: int32(0x1061),
		Fend:   int32(0x1061),
	},
	166: {
		Fstart: int32(0x1065),
		Fend:   int32(0x1066),
	},
	167: {
		Fstart: int32(0x106e),
		Fend:   int32(0x1070),
	},
	168: {
		Fstart: int32(0x1075),
		Fend:   int32(0x1081),
	},
	169: {
		Fstart: int32(0x108e),
		Fend:   int32(0x108e),
	},
	170: {
		Fstart: int32(0x10a0),
		Fend:   int32(0x10c5),
	},
	171: {
		Fstart: int32(0x10c7),
		Fend:   int32(0x10c7),
	},
	172: {
		Fstart: int32(0x10cd),
		Fend:   int32(0x10cd),
	},
	173: {
		Fstart: int32(0x10d0),
		Fend:   int32(0x10fa),
	},
	174: {
		Fstart: int32(0x10fc),
		Fend:   int32(0x1248),
	},
	175: {
		Fstart: int32(0x124a),
		Fend:   int32(0x124d),
	},
	176: {
		Fstart: int32(0x1250),
		Fend:   int32(0x1256),
	},
	177: {
		Fstart: int32(0x1258),
		Fend:   int32(0x1258),
	},
	178: {
		Fstart: int32(0x125a),
		Fend:   int32(0x125d),
	},
	179: {
		Fstart: int32(0x1260),
		Fend:   int32(0x1288),
	},
	180: {
		Fstart: int32(0x128a),
		Fend:   int32(0x128d),
	},
	181: {
		Fstart: int32(0x1290),
		Fend:   int32(0x12b0),
	},
	182: {
		Fstart: int32(0x12b2),
		Fend:   int32(0x12b5),
	},
	183: {
		Fstart: int32(0x12b8),
		Fend:   int32(0x12be),
	},
	184: {
		Fstart: int32(0x12c0),
		Fend:   int32(0x12c0),
	},
	185: {
		Fstart: int32(0x12c2),
		Fend:   int32(0x12c5),
	},
	186: {
		Fstart: int32(0x12c8),
		Fend:   int32(0x12d6),
	},
	187: {
		Fstart: int32(0x12d8),
		Fend:   int32(0x1310),
	},
	188: {
		Fstart: int32(0x1312),
		Fend:   int32(0x1315),
	},
	189: {
		Fstart: int32(0x1318),
		Fend:   int32(0x135a),
	},
	190: {
		Fstart: int32(0x1380),
		Fend:   int32(0x138f),
	},
	191: {
		Fstart: int32(0x13a0),
		Fend:   int32(0x13f5),
	},
	192: {
		Fstart: int32(0x13f8),
		Fend:   int32(0x13fd),
	},
	193: {
		Fstart: int32(0x1401),
		Fend:   int32(0x166c),
	},
	194: {
		Fstart: int32(0x166f),
		Fend:   int32(0x167f),
	},
	195: {
		Fstart: int32(0x1681),
		Fend:   int32(0x169a),
	},
	196: {
		Fstart: int32(0x16a0),
		Fend:   int32(0x16ea),
	},
	197: {
		Fstart: int32(0x16f1),
		Fend:   int32(0x16f8),
	},
	198: {
		Fstart: int32(0x1700),
		Fend:   int32(0x1711),
	},
	199: {
		Fstart: int32(0x171f),
		Fend:   int32(0x1731),
	},
	200: {
		Fstart: int32(0x1740),
		Fend:   int32(0x1751),
	},
	201: {
		Fstart: int32(0x1760),
		Fend:   int32(0x176c),
	},
	202: {
		Fstart: int32(0x176e),
		Fend:   int32(0x1770),
	},
	203: {
		Fstart: int32(0x1780),
		Fend:   int32(0x17b3),
	},
	204: {
		Fstart: int32(0x17d7),
		Fend:   int32(0x17d7),
	},
	205: {
		Fstart: int32(0x17dc),
		Fend:   int32(0x17dc),
	},
	206: {
		Fstart: int32(0x1820),
		Fend:   int32(0x1878),
	},
	207: {
		Fstart: int32(0x1880),
		Fend:   int32(0x1884),
	},
	208: {
		Fstart: int32(0x1887),
		Fend:   int32(0x18a8),
	},
	209: {
		Fstart: int32(0x18aa),
		Fend:   int32(0x18aa),
	},
	210: {
		Fstart: int32(0x18b0),
		Fend:   int32(0x18f5),
	},
	211: {
		Fstart: int32(0x1900),
		Fend:   int32(0x191e),
	},
	212: {
		Fstart: int32(0x1950),
		Fend:   int32(0x196d),
	},
	213: {
		Fstart: int32(0x1970),
		Fend:   int32(0x1974),
	},
	214: {
		Fstart: int32(0x1980),
		Fend:   int32(0x19ab),
	},
	215: {
		Fstart: int32(0x19b0),
		Fend:   int32(0x19c9),
	},
	216: {
		Fstart: int32(0x1a00),
		Fend:   int32(0x1a16),
	},
	217: {
		Fstart: int32(0x1a20),
		Fend:   int32(0x1a54),
	},
	218: {
		Fstart: int32(0x1aa7),
		Fend:   int32(0x1aa7),
	},
	219: {
		Fstart: int32(0x1b05),
		Fend:   int32(0x1b33),
	},
	220: {
		Fstart: int32(0x1b45),
		Fend:   int32(0x1b4c),
	},
	221: {
		Fstart: int32(0x1b83),
		Fend:   int32(0x1ba0),
	},
	222: {
		Fstart: int32(0x1bae),
		Fend:   int32(0x1baf),
	},
	223: {
		Fstart: int32(0x1bba),
		Fend:   int32(0x1be5),
	},
	224: {
		Fstart: int32(0x1c00),
		Fend:   int32(0x1c23),
	},
	225: {
		Fstart: int32(0x1c4d),
		Fend:   int32(0x1c4f),
	},
	226: {
		Fstart: int32(0x1c5a),
		Fend:   int32(0x1c7d),
	},
	227: {
		Fstart: int32(0x1c80),
		Fend:   int32(0x1c8a),
	},
	228: {
		Fstart: int32(0x1c90),
		Fend:   int32(0x1cba),
	},
	229: {
		Fstart: int32(0x1cbd),
		Fend:   int32(0x1cbf),
	},
	230: {
		Fstart: int32(0x1ce9),
		Fend:   int32(0x1cec),
	},
	231: {
		Fstart: int32(0x1cee),
		Fend:   int32(0x1cf3),
	},
	232: {
		Fstart: int32(0x1cf5),
		Fend:   int32(0x1cf6),
	},
	233: {
		Fstart: int32(0x1cfa),
		Fend:   int32(0x1cfa),
	},
	234: {
		Fstart: int32(0x1d00),
		Fend:   int32(0x1dbf),
	},
	235: {
		Fstart: int32(0x1e00),
		Fend:   int32(0x1f15),
	},
	236: {
		Fstart: int32(0x1f18),
		Fend:   int32(0x1f1d),
	},
	237: {
		Fstart: int32(0x1f20),
		Fend:   int32(0x1f45),
	},
	238: {
		Fstart: int32(0x1f48),
		Fend:   int32(0x1f4d),
	},
	239: {
		Fstart: int32(0x1f50),
		Fend:   int32(0x1f57),
	},
	240: {
		Fstart: int32(0x1f59),
		Fend:   int32(0x1f59),
	},
	241: {
		Fstart: int32(0x1f5b),
		Fend:   int32(0x1f5b),
	},
	242: {
		Fstart: int32(0x1f5d),
		Fend:   int32(0x1f5d),
	},
	243: {
		Fstart: int32(0x1f5f),
		Fend:   int32(0x1f7d),
	},
	244: {
		Fstart: int32(0x1f80),
		Fend:   int32(0x1fb4),
	},
	245: {
		Fstart: int32(0x1fb6),
		Fend:   int32(0x1fbc),
	},
	246: {
		Fstart: int32(0x1fbe),
		Fend:   int32(0x1fbe),
	},
	247: {
		Fstart: int32(0x1fc2),
		Fend:   int32(0x1fc4),
	},
	248: {
		Fstart: int32(0x1fc6),
		Fend:   int32(0x1fcc),
	},
	249: {
		Fstart: int32(0x1fd0),
		Fend:   int32(0x1fd3),
	},
	250: {
		Fstart: int32(0x1fd6),
		Fend:   int32(0x1fdb),
	},
	251: {
		Fstart: int32(0x1fe0),
		Fend:   int32(0x1fec),
	},
	252: {
		Fstart: int32(0x1ff2),
		Fend:   int32(0x1ff4),
	},
	253: {
		Fstart: int32(0x1ff6),
		Fend:   int32(0x1ffc),
	},
	254: {
		Fstart: int32(0x2071),
		Fend:   int32(0x2071),
	},
	255: {
		Fstart: int32(0x207f),
		Fend:   int32(0x207f),
	},
	256: {
		Fstart: int32(0x2090),
		Fend:   int32(0x209c),
	},
	257: {
		Fstart: int32(0x2102),
		Fend:   int32(0x2102),
	},
	258: {
		Fstart: int32(0x2107),
		Fend:   int32(0x2107),
	},
	259: {
		Fstart: int32(0x210a),
		Fend:   int32(0x2113),
	},
	260: {
		Fstart: int32(0x2115),
		Fend:   int32(0x2115),
	},
	261: {
		Fstart: int32(0x2119),
		Fend:   int32(0x211d),
	},
	262: {
		Fstart: int32(0x2124),
		Fend:   int32(0x2124),
	},
	263: {
		Fstart: int32(0x2126),
		Fend:   int32(0x2126),
	},
	264: {
		Fstart: int32(0x2128),
		Fend:   int32(0x2128),
	},
	265: {
		Fstart: int32(0x212a),
		Fend:   int32(0x212d),
	},
	266: {
		Fstart: int32(0x212f),
		Fend:   int32(0x2139),
	},
	267: {
		Fstart: int32(0x213c),
		Fend:   int32(0x213f),
	},
	268: {
		Fstart: int32(0x2145),
		Fend:   int32(0x2149),
	},
	269: {
		Fstart: int32(0x214e),
		Fend:   int32(0x214e),
	},
	270: {
		Fstart: int32(0x2183),
		Fend:   int32(0x2184),
	},
	271: {
		Fstart: int32(0x2c00),
		Fend:   int32(0x2ce4),
	},
	272: {
		Fstart: int32(0x2ceb),
		Fend:   int32(0x2cee),
	},
	273: {
		Fstart: int32(0x2cf2),
		Fend:   int32(0x2cf3),
	},
	274: {
		Fstart: int32(0x2d00),
		Fend:   int32(0x2d25),
	},
	275: {
		Fstart: int32(0x2d27),
		Fend:   int32(0x2d27),
	},
	276: {
		Fstart: int32(0x2d2d),
		Fend:   int32(0x2d2d),
	},
	277: {
		Fstart: int32(0x2d30),
		Fend:   int32(0x2d67),
	},
	278: {
		Fstart: int32(0x2d6f),
		Fend:   int32(0x2d6f),
	},
	279: {
		Fstart: int32(0x2d80),
		Fend:   int32(0x2d96),
	},
	280: {
		Fstart: int32(0x2da0),
		Fend:   int32(0x2da6),
	},
	281: {
		Fstart: int32(0x2da8),
		Fend:   int32(0x2dae),
	},
	282: {
		Fstart: int32(0x2db0),
		Fend:   int32(0x2db6),
	},
	283: {
		Fstart: int32(0x2db8),
		Fend:   int32(0x2dbe),
	},
	284: {
		Fstart: int32(0x2dc0),
		Fend:   int32(0x2dc6),
	},
	285: {
		Fstart: int32(0x2dc8),
		Fend:   int32(0x2dce),
	},
	286: {
		Fstart: int32(0x2dd0),
		Fend:   int32(0x2dd6),
	},
	287: {
		Fstart: int32(0x2dd8),
		Fend:   int32(0x2dde),
	},
	288: {
		Fstart: int32(0x2e2f),
		Fend:   int32(0x2e2f),
	},
	289: {
		Fstart: int32(0x3005),
		Fend:   int32(0x3006),
	},
	290: {
		Fstart: int32(0x3031),
		Fend:   int32(0x3035),
	},
	291: {
		Fstart: int32(0x303b),
		Fend:   int32(0x303c),
	},
	292: {
		Fstart: int32(0x3041),
		Fend:   int32(0x3096),
	},
	293: {
		Fstart: int32(0x309d),
		Fend:   int32(0x309f),
	},
	294: {
		Fstart: int32(0x30a1),
		Fend:   int32(0x30fa),
	},
	295: {
		Fstart: int32(0x30fc),
		Fend:   int32(0x30ff),
	},
	296: {
		Fstart: int32(0x3105),
		Fend:   int32(0x312f),
	},
	297: {
		Fstart: int32(0x3131),
		Fend:   int32(0x318e),
	},
	298: {
		Fstart: int32(0x31a0),
		Fend:   int32(0x31bf),
	},
	299: {
		Fstart: int32(0x31f0),
		Fend:   int32(0x31ff),
	},
	300: {
		Fstart: int32(0x3400),
		Fend:   int32(0x4dbf),
	},
	301: {
		Fstart: int32(0x4e00),
		Fend:   int32(0xa48c),
	},
	302: {
		Fstart: int32(0xa4d0),
		Fend:   int32(0xa4fd),
	},
	303: {
		Fstart: int32(0xa500),
		Fend:   int32(0xa60c),
	},
	304: {
		Fstart: int32(0xa610),
		Fend:   int32(0xa61f),
	},
	305: {
		Fstart: int32(0xa62a),
		Fend:   int32(0xa62b),
	},
	306: {
		Fstart: int32(0xa640),
		Fend:   int32(0xa66e),
	},
	307: {
		Fstart: int32(0xa67f),
		Fend:   int32(0xa69d),
	},
	308: {
		Fstart: int32(0xa6a0),
		Fend:   int32(0xa6e5),
	},
	309: {
		Fstart: int32(0xa717),
		Fend:   int32(0xa71f),
	},
	310: {
		Fstart: int32(0xa722),
		Fend:   int32(0xa788),
	},
	311: {
		Fstart: int32(0xa78b),
		Fend:   int32(0xa7cd),
	},
	312: {
		Fstart: int32(0xa7d0),
		Fend:   int32(0xa7d1),
	},
	313: {
		Fstart: int32(0xa7d3),
		Fend:   int32(0xa7d3),
	},
	314: {
		Fstart: int32(0xa7d5),
		Fend:   int32(0xa7dc),
	},
	315: {
		Fstart: int32(0xa7f2),
		Fend:   int32(0xa801),
	},
	316: {
		Fstart: int32(0xa803),
		Fend:   int32(0xa805),
	},
	317: {
		Fstart: int32(0xa807),
		Fend:   int32(0xa80a),
	},
	318: {
		Fstart: int32(0xa80c),
		Fend:   int32(0xa822),
	},
	319: {
		Fstart: int32(0xa840),
		Fend:   int32(0xa873),
	},
	320: {
		Fstart: int32(0xa882),
		Fend:   int32(0xa8b3),
	},
	321: {
		Fstart: int32(0xa8f2),
		Fend:   int32(0xa8f7),
	},
	322: {
		Fstart: int32(0xa8fb),
		Fend:   int32(0xa8fb),
	},
	323: {
		Fstart: int32(0xa8fd),
		Fend:   int32(0xa8fe),
	},
	324: {
		Fstart: int32(0xa90a),
		Fend:   int32(0xa925),
	},
	325: {
		Fstart: int32(0xa930),
		Fend:   int32(0xa946),
	},
	326: {
		Fstart: int32(0xa960),
		Fend:   int32(0xa97c),
	},
	327: {
		Fstart: int32(0xa984),
		Fend:   int32(0xa9b2),
	},
	328: {
		Fstart: int32(0xa9cf),
		Fend:   int32(0xa9cf),
	},
	329: {
		Fstart: int32(0xa9e0),
		Fend:   int32(0xa9e4),
	},
	330: {
		Fstart: int32(0xa9e6),
		Fend:   int32(0xa9ef),
	},
	331: {
		Fstart: int32(0xa9fa),
		Fend:   int32(0xa9fe),
	},
	332: {
		Fstart: int32(0xaa00),
		Fend:   int32(0xaa28),
	},
	333: {
		Fstart: int32(0xaa40),
		Fend:   int32(0xaa42),
	},
	334: {
		Fstart: int32(0xaa44),
		Fend:   int32(0xaa4b),
	},
	335: {
		Fstart: int32(0xaa60),
		Fend:   int32(0xaa76),
	},
	336: {
		Fstart: int32(0xaa7a),
		Fend:   int32(0xaa7a),
	},
	337: {
		Fstart: int32(0xaa7e),
		Fend:   int32(0xaaaf),
	},
	338: {
		Fstart: int32(0xaab1),
		Fend:   int32(0xaab1),
	},
	339: {
		Fstart: int32(0xaab5),
		Fend:   int32(0xaab6),
	},
	340: {
		Fstart: int32(0xaab9),
		Fend:   int32(0xaabd),
	},
	341: {
		Fstart: int32(0xaac0),
		Fend:   int32(0xaac0),
	},
	342: {
		Fstart: int32(0xaac2),
		Fend:   int32(0xaac2),
	},
	343: {
		Fstart: int32(0xaadb),
		Fend:   int32(0xaadd),
	},
	344: {
		Fstart: int32(0xaae0),
		Fend:   int32(0xaaea),
	},
	345: {
		Fstart: int32(0xaaf2),
		Fend:   int32(0xaaf4),
	},
	346: {
		Fstart: int32(0xab01),
		Fend:   int32(0xab06),
	},
	347: {
		Fstart: int32(0xab09),
		Fend:   int32(0xab0e),
	},
	348: {
		Fstart: int32(0xab11),
		Fend:   int32(0xab16),
	},
	349: {
		Fstart: int32(0xab20),
		Fend:   int32(0xab26),
	},
	350: {
		Fstart: int32(0xab28),
		Fend:   int32(0xab2e),
	},
	351: {
		Fstart: int32(0xab30),
		Fend:   int32(0xab5a),
	},
	352: {
		Fstart: int32(0xab5c),
		Fend:   int32(0xab69),
	},
	353: {
		Fstart: int32(0xab70),
		Fend:   int32(0xabe2),
	},
	354: {
		Fstart: int32(0xac00),
		Fend:   int32(0xd7a3),
	},
	355: {
		Fstart: int32(0xd7b0),
		Fend:   int32(0xd7c6),
	},
	356: {
		Fstart: int32(0xd7cb),
		Fend:   int32(0xd7fb),
	},
	357: {
		Fstart: int32(0xf900),
		Fend:   int32(0xfa6d),
	},
	358: {
		Fstart: int32(0xfa70),
		Fend:   int32(0xfad9),
	},
	359: {
		Fstart: int32(0xfb00),
		Fend:   int32(0xfb06),
	},
	360: {
		Fstart: int32(0xfb13),
		Fend:   int32(0xfb17),
	},
	361: {
		Fstart: int32(0xfb1d),
		Fend:   int32(0xfb1d),
	},
	362: {
		Fstart: int32(0xfb1f),
		Fend:   int32(0xfb28),
	},
	363: {
		Fstart: int32(0xfb2a),
		Fend:   int32(0xfb36),
	},
	364: {
		Fstart: int32(0xfb38),
		Fend:   int32(0xfb3c),
	},
	365: {
		Fstart: int32(0xfb3e),
		Fend:   int32(0xfb3e),
	},
	366: {
		Fstart: int32(0xfb40),
		Fend:   int32(0xfb41),
	},
	367: {
		Fstart: int32(0xfb43),
		Fend:   int32(0xfb44),
	},
	368: {
		Fstart: int32(0xfb46),
		Fend:   int32(0xfbb1),
	},
	369: {
		Fstart: int32(0xfbd3),
		Fend:   int32(0xfd3d),
	},
	370: {
		Fstart: int32(0xfd50),
		Fend:   int32(0xfd8f),
	},
	371: {
		Fstart: int32(0xfd92),
		Fend:   int32(0xfdc7),
	},
	372: {
		Fstart: int32(0xfdf0),
		Fend:   int32(0xfdfb),
	},
	373: {
		Fstart: int32(0xfe70),
		Fend:   int32(0xfe74),
	},
	374: {
		Fstart: int32(0xfe76),
		Fend:   int32(0xfefc),
	},
	375: {
		Fstart: int32(0xff21),
		Fend:   int32(0xff3a),
	},
	376: {
		Fstart: int32(0xff41),
		Fend:   int32(0xff5a),
	},
	377: {
		Fstart: int32(0xff66),
		Fend:   int32(0xffbe),
	},
	378: {
		Fstart: int32(0xffc2),
		Fend:   int32(0xffc7),
	},
	379: {
		Fstart: int32(0xffca),
		Fend:   int32(0xffcf),
	},
	380: {
		Fstart: int32(0xffd2),
		Fend:   int32(0xffd7),
	},
	381: {
		Fstart: int32(0xffda),
		Fend:   int32(0xffdc),
	},
	382: {
		Fstart: int32(0x10000),
		Fend:   int32(0x1000b),
	},
	383: {
		Fstart: int32(0x1000d),
		Fend:   int32(0x10026),
	},
	384: {
		Fstart: int32(0x10028),
		Fend:   int32(0x1003a),
	},
	385: {
		Fstart: int32(0x1003c),
		Fend:   int32(0x1003d),
	},
	386: {
		Fstart: int32(0x1003f),
		Fend:   int32(0x1004d),
	},
	387: {
		Fstart: int32(0x10050),
		Fend:   int32(0x1005d),
	},
	388: {
		Fstart: int32(0x10080),
		Fend:   int32(0x100fa),
	},
	389: {
		Fstart: int32(0x10280),
		Fend:   int32(0x1029c),
	},
	390: {
		Fstart: int32(0x102a0),
		Fend:   int32(0x102d0),
	},
	391: {
		Fstart: int32(0x10300),
		Fend:   int32(0x1031f),
	},
	392: {
		Fstart: int32(0x1032d),
		Fend:   int32(0x10340),
	},
	393: {
		Fstart: int32(0x10342),
		Fend:   int32(0x10349),
	},
	394: {
		Fstart: int32(0x10350),
		Fend:   int32(0x10375),
	},
	395: {
		Fstart: int32(0x10380),
		Fend:   int32(0x1039d),
	},
	396: {
		Fstart: int32(0x103a0),
		Fend:   int32(0x103c3),
	},
	397: {
		Fstart: int32(0x103c8),
		Fend:   int32(0x103cf),
	},
	398: {
		Fstart: int32(0x10400),
		Fend:   int32(0x1049d),
	},
	399: {
		Fstart: int32(0x104b0),
		Fend:   int32(0x104d3),
	},
	400: {
		Fstart: int32(0x104d8),
		Fend:   int32(0x104fb),
	},
	401: {
		Fstart: int32(0x10500),
		Fend:   int32(0x10527),
	},
	402: {
		Fstart: int32(0x10530),
		Fend:   int32(0x10563),
	},
	403: {
		Fstart: int32(0x10570),
		Fend:   int32(0x1057a),
	},
	404: {
		Fstart: int32(0x1057c),
		Fend:   int32(0x1058a),
	},
	405: {
		Fstart: int32(0x1058c),
		Fend:   int32(0x10592),
	},
	406: {
		Fstart: int32(0x10594),
		Fend:   int32(0x10595),
	},
	407: {
		Fstart: int32(0x10597),
		Fend:   int32(0x105a1),
	},
	408: {
		Fstart: int32(0x105a3),
		Fend:   int32(0x105b1),
	},
	409: {
		Fstart: int32(0x105b3),
		Fend:   int32(0x105b9),
	},
	410: {
		Fstart: int32(0x105bb),
		Fend:   int32(0x105bc),
	},
	411: {
		Fstart: int32(0x105c0),
		Fend:   int32(0x105f3),
	},
	412: {
		Fstart: int32(0x10600),
		Fend:   int32(0x10736),
	},
	413: {
		Fstart: int32(0x10740),
		Fend:   int32(0x10755),
	},
	414: {
		Fstart: int32(0x10760),
		Fend:   int32(0x10767),
	},
	415: {
		Fstart: int32(0x10780),
		Fend:   int32(0x10785),
	},
	416: {
		Fstart: int32(0x10787),
		Fend:   int32(0x107b0),
	},
	417: {
		Fstart: int32(0x107b2),
		Fend:   int32(0x107ba),
	},
	418: {
		Fstart: int32(0x10800),
		Fend:   int32(0x10805),
	},
	419: {
		Fstart: int32(0x10808),
		Fend:   int32(0x10808),
	},
	420: {
		Fstart: int32(0x1080a),
		Fend:   int32(0x10835),
	},
	421: {
		Fstart: int32(0x10837),
		Fend:   int32(0x10838),
	},
	422: {
		Fstart: int32(0x1083c),
		Fend:   int32(0x1083c),
	},
	423: {
		Fstart: int32(0x1083f),
		Fend:   int32(0x10855),
	},
	424: {
		Fstart: int32(0x10860),
		Fend:   int32(0x10876),
	},
	425: {
		Fstart: int32(0x10880),
		Fend:   int32(0x1089e),
	},
	426: {
		Fstart: int32(0x108e0),
		Fend:   int32(0x108f2),
	},
	427: {
		Fstart: int32(0x108f4),
		Fend:   int32(0x108f5),
	},
	428: {
		Fstart: int32(0x10900),
		Fend:   int32(0x10915),
	},
	429: {
		Fstart: int32(0x10920),
		Fend:   int32(0x10939),
	},
	430: {
		Fstart: int32(0x10980),
		Fend:   int32(0x109b7),
	},
	431: {
		Fstart: int32(0x109be),
		Fend:   int32(0x109bf),
	},
	432: {
		Fstart: int32(0x10a00),
		Fend:   int32(0x10a00),
	},
	433: {
		Fstart: int32(0x10a10),
		Fend:   int32(0x10a13),
	},
	434: {
		Fstart: int32(0x10a15),
		Fend:   int32(0x10a17),
	},
	435: {
		Fstart: int32(0x10a19),
		Fend:   int32(0x10a35),
	},
	436: {
		Fstart: int32(0x10a60),
		Fend:   int32(0x10a7c),
	},
	437: {
		Fstart: int32(0x10a80),
		Fend:   int32(0x10a9c),
	},
	438: {
		Fstart: int32(0x10ac0),
		Fend:   int32(0x10ac7),
	},
	439: {
		Fstart: int32(0x10ac9),
		Fend:   int32(0x10ae4),
	},
	440: {
		Fstart: int32(0x10b00),
		Fend:   int32(0x10b35),
	},
	441: {
		Fstart: int32(0x10b40),
		Fend:   int32(0x10b55),
	},
	442: {
		Fstart: int32(0x10b60),
		Fend:   int32(0x10b72),
	},
	443: {
		Fstart: int32(0x10b80),
		Fend:   int32(0x10b91),
	},
	444: {
		Fstart: int32(0x10c00),
		Fend:   int32(0x10c48),
	},
	445: {
		Fstart: int32(0x10c80),
		Fend:   int32(0x10cb2),
	},
	446: {
		Fstart: int32(0x10cc0),
		Fend:   int32(0x10cf2),
	},
	447: {
		Fstart: int32(0x10d00),
		Fend:   int32(0x10d23),
	},
	448: {
		Fstart: int32(0x10d4a),
		Fend:   int32(0x10d65),
	},
	449: {
		Fstart: int32(0x10d6f),
		Fend:   int32(0x10d85),
	},
	450: {
		Fstart: int32(0x10e80),
		Fend:   int32(0x10ea9),
	},
	451: {
		Fstart: int32(0x10eb0),
		Fend:   int32(0x10eb1),
	},
	452: {
		Fstart: int32(0x10ec2),
		Fend:   int32(0x10ec4),
	},
	453: {
		Fstart: int32(0x10f00),
		Fend:   int32(0x10f1c),
	},
	454: {
		Fstart: int32(0x10f27),
		Fend:   int32(0x10f27),
	},
	455: {
		Fstart: int32(0x10f30),
		Fend:   int32(0x10f45),
	},
	456: {
		Fstart: int32(0x10f70),
		Fend:   int32(0x10f81),
	},
	457: {
		Fstart: int32(0x10fb0),
		Fend:   int32(0x10fc4),
	},
	458: {
		Fstart: int32(0x10fe0),
		Fend:   int32(0x10ff6),
	},
	459: {
		Fstart: int32(0x11003),
		Fend:   int32(0x11037),
	},
	460: {
		Fstart: int32(0x11071),
		Fend:   int32(0x11072),
	},
	461: {
		Fstart: int32(0x11075),
		Fend:   int32(0x11075),
	},
	462: {
		Fstart: int32(0x11083),
		Fend:   int32(0x110af),
	},
	463: {
		Fstart: int32(0x110d0),
		Fend:   int32(0x110e8),
	},
	464: {
		Fstart: int32(0x11103),
		Fend:   int32(0x11126),
	},
	465: {
		Fstart: int32(0x11144),
		Fend:   int32(0x11144),
	},
	466: {
		Fstart: int32(0x11147),
		Fend:   int32(0x11147),
	},
	467: {
		Fstart: int32(0x11150),
		Fend:   int32(0x11172),
	},
	468: {
		Fstart: int32(0x11176),
		Fend:   int32(0x11176),
	},
	469: {
		Fstart: int32(0x11183),
		Fend:   int32(0x111b2),
	},
	470: {
		Fstart: int32(0x111c1),
		Fend:   int32(0x111c4),
	},
	471: {
		Fstart: int32(0x111da),
		Fend:   int32(0x111da),
	},
	472: {
		Fstart: int32(0x111dc),
		Fend:   int32(0x111dc),
	},
	473: {
		Fstart: int32(0x11200),
		Fend:   int32(0x11211),
	},
	474: {
		Fstart: int32(0x11213),
		Fend:   int32(0x1122b),
	},
	475: {
		Fstart: int32(0x1123f),
		Fend:   int32(0x11240),
	},
	476: {
		Fstart: int32(0x11280),
		Fend:   int32(0x11286),
	},
	477: {
		Fstart: int32(0x11288),
		Fend:   int32(0x11288),
	},
	478: {
		Fstart: int32(0x1128a),
		Fend:   int32(0x1128d),
	},
	479: {
		Fstart: int32(0x1128f),
		Fend:   int32(0x1129d),
	},
	480: {
		Fstart: int32(0x1129f),
		Fend:   int32(0x112a8),
	},
	481: {
		Fstart: int32(0x112b0),
		Fend:   int32(0x112de),
	},
	482: {
		Fstart: int32(0x11305),
		Fend:   int32(0x1130c),
	},
	483: {
		Fstart: int32(0x1130f),
		Fend:   int32(0x11310),
	},
	484: {
		Fstart: int32(0x11313),
		Fend:   int32(0x11328),
	},
	485: {
		Fstart: int32(0x1132a),
		Fend:   int32(0x11330),
	},
	486: {
		Fstart: int32(0x11332),
		Fend:   int32(0x11333),
	},
	487: {
		Fstart: int32(0x11335),
		Fend:   int32(0x11339),
	},
	488: {
		Fstart: int32(0x1133d),
		Fend:   int32(0x1133d),
	},
	489: {
		Fstart: int32(0x11350),
		Fend:   int32(0x11350),
	},
	490: {
		Fstart: int32(0x1135d),
		Fend:   int32(0x11361),
	},
	491: {
		Fstart: int32(0x11380),
		Fend:   int32(0x11389),
	},
	492: {
		Fstart: int32(0x1138b),
		Fend:   int32(0x1138b),
	},
	493: {
		Fstart: int32(0x1138e),
		Fend:   int32(0x1138e),
	},
	494: {
		Fstart: int32(0x11390),
		Fend:   int32(0x113b5),
	},
	495: {
		Fstart: int32(0x113b7),
		Fend:   int32(0x113b7),
	},
	496: {
		Fstart: int32(0x113d1),
		Fend:   int32(0x113d1),
	},
	497: {
		Fstart: int32(0x113d3),
		Fend:   int32(0x113d3),
	},
	498: {
		Fstart: int32(0x11400),
		Fend:   int32(0x11434),
	},
	499: {
		Fstart: int32(0x11447),
		Fend:   int32(0x1144a),
	},
	500: {
		Fstart: int32(0x1145f),
		Fend:   int32(0x11461),
	},
	501: {
		Fstart: int32(0x11480),
		Fend:   int32(0x114af),
	},
	502: {
		Fstart: int32(0x114c4),
		Fend:   int32(0x114c5),
	},
	503: {
		Fstart: int32(0x114c7),
		Fend:   int32(0x114c7),
	},
	504: {
		Fstart: int32(0x11580),
		Fend:   int32(0x115ae),
	},
	505: {
		Fstart: int32(0x115d8),
		Fend:   int32(0x115db),
	},
	506: {
		Fstart: int32(0x11600),
		Fend:   int32(0x1162f),
	},
	507: {
		Fstart: int32(0x11644),
		Fend:   int32(0x11644),
	},
	508: {
		Fstart: int32(0x11680),
		Fend:   int32(0x116aa),
	},
	509: {
		Fstart: int32(0x116b8),
		Fend:   int32(0x116b8),
	},
	510: {
		Fstart: int32(0x11700),
		Fend:   int32(0x1171a),
	},
	511: {
		Fstart: int32(0x11740),
		Fend:   int32(0x11746),
	},
	512: {
		Fstart: int32(0x11800),
		Fend:   int32(0x1182b),
	},
	513: {
		Fstart: int32(0x118a0),
		Fend:   int32(0x118df),
	},
	514: {
		Fstart: int32(0x118ff),
		Fend:   int32(0x11906),
	},
	515: {
		Fstart: int32(0x11909),
		Fend:   int32(0x11909),
	},
	516: {
		Fstart: int32(0x1190c),
		Fend:   int32(0x11913),
	},
	517: {
		Fstart: int32(0x11915),
		Fend:   int32(0x11916),
	},
	518: {
		Fstart: int32(0x11918),
		Fend:   int32(0x1192f),
	},
	519: {
		Fstart: int32(0x1193f),
		Fend:   int32(0x1193f),
	},
	520: {
		Fstart: int32(0x11941),
		Fend:   int32(0x11941),
	},
	521: {
		Fstart: int32(0x119a0),
		Fend:   int32(0x119a7),
	},
	522: {
		Fstart: int32(0x119aa),
		Fend:   int32(0x119d0),
	},
	523: {
		Fstart: int32(0x119e1),
		Fend:   int32(0x119e1),
	},
	524: {
		Fstart: int32(0x119e3),
		Fend:   int32(0x119e3),
	},
	525: {
		Fstart: int32(0x11a00),
		Fend:   int32(0x11a00),
	},
	526: {
		Fstart: int32(0x11a0b),
		Fend:   int32(0x11a32),
	},
	527: {
		Fstart: int32(0x11a3a),
		Fend:   int32(0x11a3a),
	},
	528: {
		Fstart: int32(0x11a50),
		Fend:   int32(0x11a50),
	},
	529: {
		Fstart: int32(0x11a5c),
		Fend:   int32(0x11a89),
	},
	530: {
		Fstart: int32(0x11a9d),
		Fend:   int32(0x11a9d),
	},
	531: {
		Fstart: int32(0x11ab0),
		Fend:   int32(0x11af8),
	},
	532: {
		Fstart: int32(0x11bc0),
		Fend:   int32(0x11be0),
	},
	533: {
		Fstart: int32(0x11c00),
		Fend:   int32(0x11c08),
	},
	534: {
		Fstart: int32(0x11c0a),
		Fend:   int32(0x11c2e),
	},
	535: {
		Fstart: int32(0x11c40),
		Fend:   int32(0x11c40),
	},
	536: {
		Fstart: int32(0x11c72),
		Fend:   int32(0x11c8f),
	},
	537: {
		Fstart: int32(0x11d00),
		Fend:   int32(0x11d06),
	},
	538: {
		Fstart: int32(0x11d08),
		Fend:   int32(0x11d09),
	},
	539: {
		Fstart: int32(0x11d0b),
		Fend:   int32(0x11d30),
	},
	540: {
		Fstart: int32(0x11d46),
		Fend:   int32(0x11d46),
	},
	541: {
		Fstart: int32(0x11d60),
		Fend:   int32(0x11d65),
	},
	542: {
		Fstart: int32(0x11d67),
		Fend:   int32(0x11d68),
	},
	543: {
		Fstart: int32(0x11d6a),
		Fend:   int32(0x11d89),
	},
	544: {
		Fstart: int32(0x11d98),
		Fend:   int32(0x11d98),
	},
	545: {
		Fstart: int32(0x11ee0),
		Fend:   int32(0x11ef2),
	},
	546: {
		Fstart: int32(0x11f02),
		Fend:   int32(0x11f02),
	},
	547: {
		Fstart: int32(0x11f04),
		Fend:   int32(0x11f10),
	},
	548: {
		Fstart: int32(0x11f12),
		Fend:   int32(0x11f33),
	},
	549: {
		Fstart: int32(0x11fb0),
		Fend:   int32(0x11fb0),
	},
	550: {
		Fstart: int32(0x12000),
		Fend:   int32(0x12399),
	},
	551: {
		Fstart: int32(0x12480),
		Fend:   int32(0x12543),
	},
	552: {
		Fstart: int32(0x12f90),
		Fend:   int32(0x12ff0),
	},
	553: {
		Fstart: int32(0x13000),
		Fend:   int32(0x1342f),
	},
	554: {
		Fstart: int32(0x13441),
		Fend:   int32(0x13446),
	},
	555: {
		Fstart: int32(0x13460),
		Fend:   int32(0x143fa),
	},
	556: {
		Fstart: int32(0x14400),
		Fend:   int32(0x14646),
	},
	557: {
		Fstart: int32(0x16100),
		Fend:   int32(0x1611d),
	},
	558: {
		Fstart: int32(0x16800),
		Fend:   int32(0x16a38),
	},
	559: {
		Fstart: int32(0x16a40),
		Fend:   int32(0x16a5e),
	},
	560: {
		Fstart: int32(0x16a70),
		Fend:   int32(0x16abe),
	},
	561: {
		Fstart: int32(0x16ad0),
		Fend:   int32(0x16aed),
	},
	562: {
		Fstart: int32(0x16b00),
		Fend:   int32(0x16b2f),
	},
	563: {
		Fstart: int32(0x16b40),
		Fend:   int32(0x16b43),
	},
	564: {
		Fstart: int32(0x16b63),
		Fend:   int32(0x16b77),
	},
	565: {
		Fstart: int32(0x16b7d),
		Fend:   int32(0x16b8f),
	},
	566: {
		Fstart: int32(0x16d40),
		Fend:   int32(0x16d6c),
	},
	567: {
		Fstart: int32(0x16e40),
		Fend:   int32(0x16e7f),
	},
	568: {
		Fstart: int32(0x16f00),
		Fend:   int32(0x16f4a),
	},
	569: {
		Fstart: int32(0x16f50),
		Fend:   int32(0x16f50),
	},
	570: {
		Fstart: int32(0x16f93),
		Fend:   int32(0x16f9f),
	},
	571: {
		Fstart: int32(0x16fe0),
		Fend:   int32(0x16fe1),
	},
	572: {
		Fstart: int32(0x16fe3),
		Fend:   int32(0x16fe3),
	},
	573: {
		Fstart: int32(0x17000),
		Fend:   int32(0x187f7),
	},
	574: {
		Fstart: int32(0x18800),
		Fend:   int32(0x18cd5),
	},
	575: {
		Fstart: int32(0x18cff),
		Fend:   int32(0x18d08),
	},
	576: {
		Fstart: int32(0x1aff0),
		Fend:   int32(0x1aff3),
	},
	577: {
		Fstart: int32(0x1aff5),
		Fend:   int32(0x1affb),
	},
	578: {
		Fstart: int32(0x1affd),
		Fend:   int32(0x1affe),
	},
	579: {
		Fstart: int32(0x1b000),
		Fend:   int32(0x1b122),
	},
	580: {
		Fstart: int32(0x1b132),
		Fend:   int32(0x1b132),
	},
	581: {
		Fstart: int32(0x1b150),
		Fend:   int32(0x1b152),
	},
	582: {
		Fstart: int32(0x1b155),
		Fend:   int32(0x1b155),
	},
	583: {
		Fstart: int32(0x1b164),
		Fend:   int32(0x1b167),
	},
	584: {
		Fstart: int32(0x1b170),
		Fend:   int32(0x1b2fb),
	},
	585: {
		Fstart: int32(0x1bc00),
		Fend:   int32(0x1bc6a),
	},
	586: {
		Fstart: int32(0x1bc70),
		Fend:   int32(0x1bc7c),
	},
	587: {
		Fstart: int32(0x1bc80),
		Fend:   int32(0x1bc88),
	},
	588: {
		Fstart: int32(0x1bc90),
		Fend:   int32(0x1bc99),
	},
	589: {
		Fstart: int32(0x1d400),
		Fend:   int32(0x1d454),
	},
	590: {
		Fstart: int32(0x1d456),
		Fend:   int32(0x1d49c),
	},
	591: {
		Fstart: int32(0x1d49e),
		Fend:   int32(0x1d49f),
	},
	592: {
		Fstart: int32(0x1d4a2),
		Fend:   int32(0x1d4a2),
	},
	593: {
		Fstart: int32(0x1d4a5),
		Fend:   int32(0x1d4a6),
	},
	594: {
		Fstart: int32(0x1d4a9),
		Fend:   int32(0x1d4ac),
	},
	595: {
		Fstart: int32(0x1d4ae),
		Fend:   int32(0x1d4b9),
	},
	596: {
		Fstart: int32(0x1d4bb),
		Fend:   int32(0x1d4bb),
	},
	597: {
		Fstart: int32(0x1d4bd),
		Fend:   int32(0x1d4c3),
	},
	598: {
		Fstart: int32(0x1d4c5),
		Fend:   int32(0x1d505),
	},
	599: {
		Fstart: int32(0x1d507),
		Fend:   int32(0x1d50a),
	},
	600: {
		Fstart: int32(0x1d50d),
		Fend:   int32(0x1d514),
	},
	601: {
		Fstart: int32(0x1d516),
		Fend:   int32(0x1d51c),
	},
	602: {
		Fstart: int32(0x1d51e),
		Fend:   int32(0x1d539),
	},
	603: {
		Fstart: int32(0x1d53b),
		Fend:   int32(0x1d53e),
	},
	604: {
		Fstart: int32(0x1d540),
		Fend:   int32(0x1d544),
	},
	605: {
		Fstart: int32(0x1d546),
		Fend:   int32(0x1d546),
	},
	606: {
		Fstart: int32(0x1d54a),
		Fend:   int32(0x1d550),
	},
	607: {
		Fstart: int32(0x1d552),
		Fend:   int32(0x1d6a5),
	},
	608: {
		Fstart: int32(0x1d6a8),
		Fend:   int32(0x1d6c0),
	},
	609: {
		Fstart: int32(0x1d6c2),
		Fend:   int32(0x1d6da),
	},
	610: {
		Fstart: int32(0x1d6dc),
		Fend:   int32(0x1d6fa),
	},
	611: {
		Fstart: int32(0x1d6fc),
		Fend:   int32(0x1d714),
	},
	612: {
		Fstart: int32(0x1d716),
		Fend:   int32(0x1d734),
	},
	613: {
		Fstart: int32(0x1d736),
		Fend:   int32(0x1d74e),
	},
	614: {
		Fstart: int32(0x1d750),
		Fend:   int32(0x1d76e),
	},
	615: {
		Fstart: int32(0x1d770),
		Fend:   int32(0x1d788),
	},
	616: {
		Fstart: int32(0x1d78a),
		Fend:   int32(0x1d7a8),
	},
	617: {
		Fstart: int32(0x1d7aa),
		Fend:   int32(0x1d7c2),
	},
	618: {
		Fstart: int32(0x1d7c4),
		Fend:   int32(0x1d7cb),
	},
	619: {
		Fstart: int32(0x1df00),
		Fend:   int32(0x1df1e),
	},
	620: {
		Fstart: int32(0x1df25),
		Fend:   int32(0x1df2a),
	},
	621: {
		Fstart: int32(0x1e030),
		Fend:   int32(0x1e06d),
	},
	622: {
		Fstart: int32(0x1e100),
		Fend:   int32(0x1e12c),
	},
	623: {
		Fstart: int32(0x1e137),
		Fend:   int32(0x1e13d),
	},
	624: {
		Fstart: int32(0x1e14e),
		Fend:   int32(0x1e14e),
	},
	625: {
		Fstart: int32(0x1e290),
		Fend:   int32(0x1e2ad),
	},
	626: {
		Fstart: int32(0x1e2c0),
		Fend:   int32(0x1e2eb),
	},
	627: {
		Fstart: int32(0x1e4d0),
		Fend:   int32(0x1e4eb),
	},
	628: {
		Fstart: int32(0x1e5d0),
		Fend:   int32(0x1e5ed),
	},
	629: {
		Fstart: int32(0x1e5f0),
		Fend:   int32(0x1e5f0),
	},
	630: {
		Fstart: int32(0x1e7e0),
		Fend:   int32(0x1e7e6),
	},
	631: {
		Fstart: int32(0x1e7e8),
		Fend:   int32(0x1e7eb),
	},
	632: {
		Fstart: int32(0x1e7ed),
		Fend:   int32(0x1e7ee),
	},
	633: {
		Fstart: int32(0x1e7f0),
		Fend:   int32(0x1e7fe),
	},
	634: {
		Fstart: int32(0x1e800),
		Fend:   int32(0x1e8c4),
	},
	635: {
		Fstart: int32(0x1e900),
		Fend:   int32(0x1e943),
	},
	636: {
		Fstart: int32(0x1e94b),
		Fend:   int32(0x1e94b),
	},
	637: {
		Fstart: int32(0x1ee00),
		Fend:   int32(0x1ee03),
	},
	638: {
		Fstart: int32(0x1ee05),
		Fend:   int32(0x1ee1f),
	},
	639: {
		Fstart: int32(0x1ee21),
		Fend:   int32(0x1ee22),
	},
	640: {
		Fstart: int32(0x1ee24),
		Fend:   int32(0x1ee24),
	},
	641: {
		Fstart: int32(0x1ee27),
		Fend:   int32(0x1ee27),
	},
	642: {
		Fstart: int32(0x1ee29),
		Fend:   int32(0x1ee32),
	},
	643: {
		Fstart: int32(0x1ee34),
		Fend:   int32(0x1ee37),
	},
	644: {
		Fstart: int32(0x1ee39),
		Fend:   int32(0x1ee39),
	},
	645: {
		Fstart: int32(0x1ee3b),
		Fend:   int32(0x1ee3b),
	},
	646: {
		Fstart: int32(0x1ee42),
		Fend:   int32(0x1ee42),
	},
	647: {
		Fstart: int32(0x1ee47),
		Fend:   int32(0x1ee47),
	},
	648: {
		Fstart: int32(0x1ee49),
		Fend:   int32(0x1ee49),
	},
	649: {
		Fstart: int32(0x1ee4b),
		Fend:   int32(0x1ee4b),
	},
	650: {
		Fstart: int32(0x1ee4d),
		Fend:   int32(0x1ee4f),
	},
	651: {
		Fstart: int32(0x1ee51),
		Fend:   int32(0x1ee52),
	},
	652: {
		Fstart: int32(0x1ee54),
		Fend:   int32(0x1ee54),
	},
	653: {
		Fstart: int32(0x1ee57),
		Fend:   int32(0x1ee57),
	},
	654: {
		Fstart: int32(0x1ee59),
		Fend:   int32(0x1ee59),
	},
	655: {
		Fstart: int32(0x1ee5b),
		Fend:   int32(0x1ee5b),
	},
	656: {
		Fstart: int32(0x1ee5d),
		Fend:   int32(0x1ee5d),
	},
	657: {
		Fstart: int32(0x1ee5f),
		Fend:   int32(0x1ee5f),
	},
	658: {
		Fstart: int32(0x1ee61),
		Fend:   int32(0x1ee62),
	},
	659: {
		Fstart: int32(0x1ee64),
		Fend:   int32(0x1ee64),
	},
	660: {
		Fstart: int32(0x1ee67),
		Fend:   int32(0x1ee6a),
	},
	661: {
		Fstart: int32(0x1ee6c),
		Fend:   int32(0x1ee72),
	},
	662: {
		Fstart: int32(0x1ee74),
		Fend:   int32(0x1ee77),
	},
	663: {
		Fstart: int32(0x1ee79),
		Fend:   int32(0x1ee7c),
	},
	664: {
		Fstart: int32(0x1ee7e),
		Fend:   int32(0x1ee7e),
	},
	665: {
		Fstart: int32(0x1ee80),
		Fend:   int32(0x1ee89),
	},
	666: {
		Fstart: int32(0x1ee8b),
		Fend:   int32(0x1ee9b),
	},
	667: {
		Fstart: int32(0x1eea1),
		Fend:   int32(0x1eea3),
	},
	668: {
		Fstart: int32(0x1eea5),
		Fend:   int32(0x1eea9),
	},
	669: {
		Fstart: int32(0x1eeab),
		Fend:   int32(0x1eebb),
	},
	670: {
		Fstart: int32(0x20000),
		Fend:   int32(0x2a6df),
	},
	671: {
		Fstart: int32(0x2a700),
		Fend:   int32(0x2b739),
	},
	672: {
		Fstart: int32(0x2b740),
		Fend:   int32(0x2b81d),
	},
	673: {
		Fstart: int32(0x2b820),
		Fend:   int32(0x2cea1),
	},
	674: {
		Fstart: int32(0x2ceb0),
		Fend:   int32(0x2ebe0),
	},
	675: {
		Fstart: int32(0x2ebf0),
		Fend:   int32(0x2ee5d),
	},
	676: {
		Fstart: int32(0x2f800),
		Fend:   int32(0x2fa1d),
	},
	677: {
		Fstart: int32(0x30000),
		Fend:   int32(0x3134a),
	},
	678: {
		Fstart: int32(0x31350),
		Fend:   int32(0x323af),
	},
}

var sym_identifier_character_set_2 = [680]TSCharacterRange{
	0: {
		Fstart: int32('$'),
		Fend:   int32('$'),
	},
	1: {
		Fstart: int32('0'),
		Fend:   int32('9'),
	},
	2: {
		Fstart: int32('A'),
		Fend:   int32('Z'),
	},
	3: {
		Fstart: int32('_'),
		Fend:   int32('_'),
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
	var half_size, i, i1, i2, i3, index, mid_index, size uint32_t
	var lookahead1, v3 int32_t
	var range_token, range_token1, v2 uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = eof, half_size, i, i1, i2, i3, index, lookahead1, mid_index, range_token, range_token1, result, size, skip, v2, v3, v4
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
			state = uint16(43)
			goto next_state
		}
		i = uint32(0)
		for {
			if !(uint64(i) < libc.Uint64FromInt64(72)/libc.Uint64FromInt64(2)) {
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
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(0)
			goto next_state
		}
		if int32('1') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(75)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_identifier_character_set_2))
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
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(1):
		if lookahead1 == int32('\n') {
			state = uint16(7)
			goto next_state
		}
		return result
	case int32(2):
		if lookahead1 == int32('\n') {
			state = uint16(8)
			goto next_state
		}
		return result
	case int32(3):
		i1 = uint32(0)
		for {
			if !(uint64(i1) < libc.Uint64FromInt64(56)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token1[i1]) == lookahead1 {
				state = map_token1[i1+uint32(1)]
				goto next_state
			}
			goto _6
		_6:
			;
			i1 = i1 + uint32(2)
		}
		return result
	case int32(4):
		i2 = uint32(0)
		for {
			if !(uint64(i2) < libc.Uint64FromInt64(56)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token2[i2]) == lookahead1 {
				state = map_token2[i2+uint32(1)]
				goto next_state
			}
			goto _7
		_7:
			;
			i2 = i2 + uint32(2)
		}
		return result
	case int32(5):
		i3 = uint32(0)
		for {
			if !(uint64(i3) < libc.Uint64FromInt64(60)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token3[i3]) == lookahead1 {
				state = map_token3[i3+uint32(1)]
				goto next_state
			}
			goto _8
		_8:
			;
			i3 = i3 + uint32(2)
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(5)
			goto next_state
		}
		if int32('1') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(75)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead1 == int32('"') {
			state = uint16(7)
			goto next_state
		}
		if lookahead1 == int32('\'') {
			state = uint16(8)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(9)
			goto next_state
		}
		if lookahead1 == int32('}') {
			state = uint16(48)
			goto next_state
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(6)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_identifier_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(679) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _12
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _12
	_12:
		if v4 != 0 {
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead1 == int32('"') {
			state = uint16(72)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(3)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(7)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead1 == int32('\'') {
			state = uint16(72)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(4)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(8)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead1 == int32('*') {
			state = uint16(11)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(45)
			goto next_state
		}
		return result
	case int32(10):
		if lookahead1 == int32('*') {
			state = uint16(10)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(44)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(11)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead1 == int32('*') {
			state = uint16(10)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(11)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead1 == int32('.') {
			state = uint16(76)
			goto next_state
		}
		if lookahead1 == int32('0') {
			state = uint16(74)
			goto next_state
		}
		if lookahead1 == int32('I') {
			state = uint16(24)
			goto next_state
		}
		if lookahead1 == int32('N') {
			state = uint16(14)
			goto next_state
		}
		if int32('1') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(75)
			goto next_state
		}
		return result
	case int32(13):
		if lookahead1 == int32('N') {
			state = uint16(73)
			goto next_state
		}
		return result
	case int32(14):
		if lookahead1 == int32('a') {
			state = uint16(13)
			goto next_state
		}
		return result
	case int32(15):
		if lookahead1 == int32('a') {
			state = uint16(21)
			goto next_state
		}
		return result
	case int32(16):
		if lookahead1 == int32('e') {
			state = uint16(81)
			goto next_state
		}
		return result
	case int32(17):
		if lookahead1 == int32('e') {
			state = uint16(83)
			goto next_state
		}
		return result
	case int32(18):
		if lookahead1 == int32('f') {
			state = uint16(20)
			goto next_state
		}
		return result
	case int32(19):
		if lookahead1 == int32('i') {
			state = uint16(28)
			goto next_state
		}
		return result
	case int32(20):
		if lookahead1 == int32('i') {
			state = uint16(25)
			goto next_state
		}
		return result
	case int32(21):
		if lookahead1 == int32('l') {
			state = uint16(27)
			goto next_state
		}
		return result
	case int32(22):
		if lookahead1 == int32('l') {
			state = uint16(79)
			goto next_state
		}
		return result
	case int32(23):
		if lookahead1 == int32('l') {
			state = uint16(22)
			goto next_state
		}
		return result
	case int32(24):
		if lookahead1 == int32('n') {
			state = uint16(18)
			goto next_state
		}
		return result
	case int32(25):
		if lookahead1 == int32('n') {
			state = uint16(19)
			goto next_state
		}
		return result
	case int32(26):
		if lookahead1 == int32('r') {
			state = uint16(29)
			goto next_state
		}
		return result
	case int32(27):
		if lookahead1 == int32('s') {
			state = uint16(17)
			goto next_state
		}
		return result
	case int32(28):
		if lookahead1 == int32('t') {
			state = uint16(31)
			goto next_state
		}
		return result
	case int32(29):
		if lookahead1 == int32('u') {
			state = uint16(16)
			goto next_state
		}
		return result
	case int32(30):
		if lookahead1 == int32('u') {
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(31):
		if lookahead1 == int32('y') {
			state = uint16(73)
			goto next_state
		}
		return result
	case int32(32):
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(33)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(77)
			goto next_state
		}
		return result
	case int32(33):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(77)
			goto next_state
		}
		return result
	case int32(34):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(7)
			goto next_state
		}
		return result
	case int32(35):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(8)
			goto next_state
		}
		return result
	case int32(36):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(78)
			goto next_state
		}
		return result
	case int32(37):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(34)
			goto next_state
		}
		return result
	case int32(38):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(35)
			goto next_state
		}
		return result
	case int32(39):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(37)
			goto next_state
		}
		return result
	case int32(40):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(38)
			goto next_state
		}
		return result
	case int32(41):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(39)
			goto next_state
		}
		return result
	case int32(42):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(40)
			goto next_state
		}
		return result
	case int32(43):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(ts_builtin_sym_end)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(44):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(45):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 != 0 && lookahead1 != int32('\n') {
			state = uint16(45)
			goto next_state
		}
		return result
	case int32(46):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(47):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COMMA)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(48):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(49):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COLON)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(50):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('N') {
			state = uint16(69)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_identifier_character_set_2))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(680) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _16
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _16
	_16:
		if v4 != 0 {
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(51):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('a') {
			state = uint16(50)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_identifier_character_set_2))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(680) - index
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
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(52):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('a') {
			state = uint16(58)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_identifier_character_set_2))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(680) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _24
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _24
	_24:
		if v4 != 0 {
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(53):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('e') {
			state = uint16(82)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_identifier_character_set_2))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(680) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _28
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _28
	_28:
		if v4 != 0 {
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(54):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('e') {
			state = uint16(84)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_identifier_character_set_2))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(680) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _32
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _32
	_32:
		if v4 != 0 {
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(55):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('f') {
			state = uint16(57)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_identifier_character_set_2))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(680) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _36
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _36
	_36:
		if v4 != 0 {
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(56):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('i') {
			state = uint16(65)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_identifier_character_set_2))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(680) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _40
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _40
	_40:
		if v4 != 0 {
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(57):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('i') {
			state = uint16(62)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_identifier_character_set_2))
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
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(58):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('l') {
			state = uint16(64)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_identifier_character_set_2))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(680) - index
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
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(59):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('l') {
			state = uint16(80)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_identifier_character_set_2))
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
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(60):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('l') {
			state = uint16(59)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_identifier_character_set_2))
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
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(61):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('n') {
			state = uint16(55)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_identifier_character_set_2))
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
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(62):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('n') {
			state = uint16(56)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_identifier_character_set_2))
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
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(63):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('r') {
			state = uint16(66)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_identifier_character_set_2))
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
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(64):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('s') {
			state = uint16(54)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_identifier_character_set_2))
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
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(65):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('t') {
			state = uint16(68)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_identifier_character_set_2))
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
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(66):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('u') {
			state = uint16(53)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_identifier_character_set_2))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(680) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _80
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _80
	_80:
		if v4 != 0 {
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(67):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('u') {
			state = uint16(60)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_identifier_character_set_2))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(680) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _84
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _84
	_84:
		if v4 != 0 {
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(68):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('y') {
			state = uint16(69)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_identifier_character_set_2))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(680) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _88
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _88
	_88:
		if v4 != 0 {
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(69):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		v2 = uintptr(unsafe.Pointer(&sym_identifier_character_set_2))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(680) - index
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
		if v4 != 0 {
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(70):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(71):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(72):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(73):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(74):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('.') {
			state = uint16(76)
			goto next_state
		}
		if lookahead1 == int32('E') || lookahead1 == int32('e') {
			state = uint16(32)
			goto next_state
		}
		if lookahead1 == int32('X') || lookahead1 == int32('x') {
			state = uint16(36)
			goto next_state
		}
		return result
	case int32(75):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('.') {
			state = uint16(76)
			goto next_state
		}
		if lookahead1 == int32('E') || lookahead1 == int32('e') {
			state = uint16(32)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(75)
			goto next_state
		}
		return result
	case int32(76):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('E') || lookahead1 == int32('e') {
			state = uint16(32)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(76)
			goto next_state
		}
		return result
	case int32(77):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(77)
			goto next_state
		}
		return result
	case int32(78):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(78)
			goto next_state
		}
		return result
	case int32(79):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_null)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(80):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_null)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		v2 = uintptr(unsafe.Pointer(&sym_identifier_character_set_2))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(680) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _96
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _96
	_96:
		if v4 != 0 {
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(81):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_true)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(82):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_true)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		v2 = uintptr(unsafe.Pointer(&sym_identifier_character_set_2))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(680) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _100
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _100
	_100:
		if v4 != 0 {
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(83):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_false)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(84):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_false)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		v2 = uintptr(unsafe.Pointer(&sym_identifier_character_set_2))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(680) - index
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
		if v4 != 0 {
			state = uint16(69)
			goto next_state
		}
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var map_token = [36]uint16_t{
	0:  uint16('"'),
	1:  uint16(7),
	2:  uint16('\''),
	3:  uint16(8),
	4:  uint16(','),
	5:  uint16(47),
	6:  uint16('.'),
	7:  uint16(76),
	8:  uint16('/'),
	9:  uint16(9),
	10: uint16('0'),
	11: uint16(74),
	12: uint16(':'),
	13: uint16(49),
	14: uint16('I'),
	15: uint16(61),
	16: uint16('N'),
	17: uint16(51),
	18: uint16('['),
	19: uint16(70),
	20: uint16(']'),
	21: uint16(71),
	22: uint16('f'),
	23: uint16(52),
	24: uint16('n'),
	25: uint16(67),
	26: uint16('t'),
	27: uint16(63),
	28: uint16('{'),
	29: uint16(46),
	30: uint16('}'),
	31: uint16(48),
	32: uint16('+'),
	33: uint16(12),
	34: uint16('-'),
	35: uint16(12),
}

var map_token1 = [28]uint16_t{
	0:  uint16('\r'),
	1:  uint16(1),
	2:  uint16('u'),
	3:  uint16(41),
	4:  uint16('x'),
	5:  uint16(37),
	6:  uint16('\n'),
	7:  uint16(7),
	8:  uint16('"'),
	9:  uint16(7),
	10: uint16('\''),
	11: uint16(7),
	12: uint16('/'),
	13: uint16(7),
	14: uint16('\\'),
	15: uint16(7),
	16: uint16('b'),
	17: uint16(7),
	18: uint16('f'),
	19: uint16(7),
	20: uint16('n'),
	21: uint16(7),
	22: uint16('r'),
	23: uint16(7),
	24: uint16('t'),
	25: uint16(7),
	26: uint16('v'),
	27: uint16(7),
}

var map_token2 = [28]uint16_t{
	0:  uint16('\r'),
	1:  uint16(2),
	2:  uint16('u'),
	3:  uint16(42),
	4:  uint16('x'),
	5:  uint16(38),
	6:  uint16('\n'),
	7:  uint16(8),
	8:  uint16('"'),
	9:  uint16(8),
	10: uint16('\''),
	11: uint16(8),
	12: uint16('/'),
	13: uint16(8),
	14: uint16('\\'),
	15: uint16(8),
	16: uint16('b'),
	17: uint16(8),
	18: uint16('f'),
	19: uint16(8),
	20: uint16('n'),
	21: uint16(8),
	22: uint16('r'),
	23: uint16(8),
	24: uint16('t'),
	25: uint16(8),
	26: uint16('v'),
	27: uint16(8),
}

var map_token3 = [30]uint16_t{
	0:  uint16('"'),
	1:  uint16(7),
	2:  uint16('\''),
	3:  uint16(8),
	4:  uint16('.'),
	5:  uint16(76),
	6:  uint16('/'),
	7:  uint16(9),
	8:  uint16('0'),
	9:  uint16(74),
	10: uint16('I'),
	11: uint16(24),
	12: uint16('N'),
	13: uint16(14),
	14: uint16('['),
	15: uint16(70),
	16: uint16(']'),
	17: uint16(71),
	18: uint16('f'),
	19: uint16(15),
	20: uint16('n'),
	21: uint16(30),
	22: uint16('t'),
	23: uint16(26),
	24: uint16('{'),
	25: uint16(46),
	26: uint16('+'),
	27: uint16(12),
	28: uint16('-'),
	29: uint16(12),
}

var ts_lex_modes = [31]TSLexerMode{
	0: {},
	1: {
		Flex_state: uint16(5),
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
		Flex_state: uint16(5),
	},
	6: {
		Flex_state: uint16(5),
	},
	7: {
		Flex_state: uint16(6),
	},
	8: {},
	9: {
		Flex_state: uint16(6),
	},
	10: {},
	11: {},
	12: {},
	13: {
		Flex_state: uint16(6),
	},
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
	24: {
		Flex_state: uint16(6),
	},
	25: {},
	26: {},
	27: {},
	28: {},
	29: {},
	30: {},
}

var ts_parse_table = [7][21]uint16_t{
	0: {
		0:  uint16(1),
		1:  uint16(3),
		2:  uint16(1),
		3:  uint16(1),
		4:  uint16(1),
		5:  uint16(1),
		6:  uint16(1),
		7:  uint16(1),
		8:  uint16(1),
		9:  uint16(1),
		10: uint16(1),
		11: uint16(1),
		12: uint16(1),
		13: uint16(1),
	},
	1: {
		1:  uint16(3),
		2:  uint16(5),
		7:  uint16(7),
		9:  uint16(9),
		10: uint16(9),
		11: uint16(9),
		12: uint16(9),
		13: uint16(9),
		14: uint16(29),
		15: uint16(30),
		17: uint16(30),
		18: uint16(30),
	},
	2: {
		1:  uint16(3),
		2:  uint16(5),
		7:  uint16(7),
		8:  uint16(11),
		9:  uint16(13),
		10: uint16(13),
		11: uint16(13),
		12: uint16(13),
		13: uint16(13),
		15: uint16(19),
		17: uint16(19),
		18: uint16(19),
	},
	3: {
		1:  uint16(3),
		2:  uint16(5),
		7:  uint16(7),
		8:  uint16(15),
		9:  uint16(17),
		10: uint16(17),
		11: uint16(17),
		12: uint16(17),
		13: uint16(17),
		15: uint16(25),
		17: uint16(25),
		18: uint16(25),
	},
	4: {
		1:  uint16(3),
		2:  uint16(5),
		7:  uint16(7),
		8:  uint16(19),
		9:  uint16(17),
		10: uint16(17),
		11: uint16(17),
		12: uint16(17),
		13: uint16(17),
		15: uint16(25),
		17: uint16(25),
		18: uint16(25),
	},
	5: {
		1:  uint16(3),
		2:  uint16(5),
		7:  uint16(7),
		9:  uint16(21),
		10: uint16(21),
		11: uint16(21),
		12: uint16(21),
		13: uint16(21),
		15: uint16(26),
		17: uint16(26),
		18: uint16(26),
	},
	6: {
		1:  uint16(3),
		2:  uint16(5),
		7:  uint16(7),
		9:  uint16(17),
		10: uint16(17),
		11: uint16(17),
		12: uint16(17),
		13: uint16(17),
		15: uint16(25),
		17: uint16(25),
		18: uint16(25),
	},
}

var ts_small_parse_table = [256]uint16_t{
	0:   uint16(4),
	1:   uint16(3),
	2:   uint16(1),
	3:   uint16(sym_comment),
	4:   uint16(23),
	5:   uint16(1),
	6:   uint16(anon_sym_RBRACE),
	7:   uint16(18),
	8:   uint16(1),
	9:   uint16(sym_member),
	10:  uint16(25),
	11:  uint16(2),
	12:  uint16(sym_identifier),
	13:  uint16(sym_string),
	14:  uint16(2),
	15:  uint16(3),
	16:  uint16(1),
	17:  uint16(sym_comment),
	18:  uint16(27),
	19:  uint16(4),
	21:  uint16(anon_sym_COMMA),
	22:  uint16(anon_sym_RBRACE),
	23:  uint16(anon_sym_RBRACK),
	24:  uint16(4),
	25:  uint16(3),
	26:  uint16(1),
	27:  uint16(sym_comment),
	28:  uint16(29),
	29:  uint16(1),
	30:  uint16(anon_sym_RBRACE),
	31:  uint16(27),
	32:  uint16(1),
	33:  uint16(sym_member),
	34:  uint16(25),
	35:  uint16(2),
	36:  uint16(sym_identifier),
	37:  uint16(sym_string),
	38:  uint16(2),
	39:  uint16(3),
	40:  uint16(1),
	41:  uint16(sym_comment),
	42:  uint16(31),
	43:  uint16(4),
	45:  uint16(anon_sym_COMMA),
	46:  uint16(anon_sym_RBRACE),
	47:  uint16(anon_sym_RBRACK),
	48:  uint16(2),
	49:  uint16(3),
	50:  uint16(1),
	51:  uint16(sym_comment),
	52:  uint16(33),
	53:  uint16(4),
	55:  uint16(anon_sym_COMMA),
	56:  uint16(anon_sym_RBRACE),
	57:  uint16(anon_sym_RBRACK),
	58:  uint16(2),
	59:  uint16(3),
	60:  uint16(1),
	61:  uint16(sym_comment),
	62:  uint16(35),
	63:  uint16(4),
	65:  uint16(anon_sym_COMMA),
	66:  uint16(anon_sym_RBRACE),
	67:  uint16(anon_sym_RBRACK),
	68:  uint16(4),
	69:  uint16(3),
	70:  uint16(1),
	71:  uint16(sym_comment),
	72:  uint16(37),
	73:  uint16(1),
	74:  uint16(anon_sym_RBRACE),
	75:  uint16(27),
	76:  uint16(1),
	77:  uint16(sym_member),
	78:  uint16(25),
	79:  uint16(2),
	80:  uint16(sym_identifier),
	81:  uint16(sym_string),
	82:  uint16(2),
	83:  uint16(3),
	84:  uint16(1),
	85:  uint16(sym_comment),
	86:  uint16(39),
	87:  uint16(4),
	89:  uint16(anon_sym_COMMA),
	90:  uint16(anon_sym_RBRACE),
	91:  uint16(anon_sym_RBRACK),
	92:  uint16(2),
	93:  uint16(3),
	94:  uint16(1),
	95:  uint16(sym_comment),
	96:  uint16(41),
	97:  uint16(4),
	99:  uint16(anon_sym_COMMA),
	100: uint16(anon_sym_RBRACE),
	101: uint16(anon_sym_RBRACK),
	102: uint16(2),
	103: uint16(3),
	104: uint16(1),
	105: uint16(sym_comment),
	106: uint16(43),
	107: uint16(4),
	109: uint16(anon_sym_COMMA),
	110: uint16(anon_sym_RBRACE),
	111: uint16(anon_sym_RBRACK),
	112: uint16(2),
	113: uint16(3),
	114: uint16(1),
	115: uint16(sym_comment),
	116: uint16(45),
	117: uint16(4),
	119: uint16(anon_sym_COMMA),
	120: uint16(anon_sym_RBRACE),
	121: uint16(anon_sym_RBRACK),
	122: uint16(4),
	123: uint16(3),
	124: uint16(1),
	125: uint16(sym_comment),
	126: uint16(47),
	127: uint16(1),
	128: uint16(anon_sym_COMMA),
	129: uint16(49),
	130: uint16(1),
	131: uint16(anon_sym_RBRACE),
	132: uint16(20),
	133: uint16(1),
	134: uint16(aux_sym_object_repeat1),
	135: uint16(4),
	136: uint16(3),
	137: uint16(1),
	138: uint16(sym_comment),
	139: uint16(51),
	140: uint16(1),
	141: uint16(anon_sym_COMMA),
	142: uint16(53),
	143: uint16(1),
	144: uint16(anon_sym_RBRACK),
	145: uint16(21),
	146: uint16(1),
	147: uint16(aux_sym_array_repeat1),
	148: uint16(4),
	149: uint16(3),
	150: uint16(1),
	151: uint16(sym_comment),
	152: uint16(29),
	153: uint16(1),
	154: uint16(anon_sym_RBRACE),
	155: uint16(55),
	156: uint16(1),
	157: uint16(anon_sym_COMMA),
	158: uint16(22),
	159: uint16(1),
	160: uint16(aux_sym_object_repeat1),
	161: uint16(4),
	162: uint16(3),
	163: uint16(1),
	164: uint16(sym_comment),
	165: uint16(15),
	166: uint16(1),
	167: uint16(anon_sym_RBRACK),
	168: uint16(57),
	169: uint16(1),
	170: uint16(anon_sym_COMMA),
	171: uint16(23),
	172: uint16(1),
	173: uint16(aux_sym_array_repeat1),
	174: uint16(4),
	175: uint16(3),
	176: uint16(1),
	177: uint16(sym_comment),
	178: uint16(59),
	179: uint16(1),
	180: uint16(anon_sym_COMMA),
	181: uint16(62),
	182: uint16(1),
	183: uint16(anon_sym_RBRACE),
	184: uint16(22),
	185: uint16(1),
	186: uint16(aux_sym_object_repeat1),
	187: uint16(4),
	188: uint16(3),
	189: uint16(1),
	190: uint16(sym_comment),
	191: uint16(64),
	192: uint16(1),
	193: uint16(anon_sym_COMMA),
	194: uint16(67),
	195: uint16(1),
	196: uint16(anon_sym_RBRACK),
	197: uint16(23),
	198: uint16(1),
	199: uint16(aux_sym_array_repeat1),
	200: uint16(3),
	201: uint16(3),
	202: uint16(1),
	203: uint16(sym_comment),
	204: uint16(27),
	205: uint16(1),
	206: uint16(sym_member),
	207: uint16(25),
	208: uint16(2),
	209: uint16(sym_identifier),
	210: uint16(sym_string),
	211: uint16(2),
	212: uint16(3),
	213: uint16(1),
	214: uint16(sym_comment),
	215: uint16(67),
	216: uint16(2),
	217: uint16(anon_sym_COMMA),
	218: uint16(anon_sym_RBRACK),
	219: uint16(2),
	220: uint16(3),
	221: uint16(1),
	222: uint16(sym_comment),
	223: uint16(69),
	224: uint16(2),
	225: uint16(anon_sym_COMMA),
	226: uint16(anon_sym_RBRACE),
	227: uint16(2),
	228: uint16(3),
	229: uint16(1),
	230: uint16(sym_comment),
	231: uint16(62),
	232: uint16(2),
	233: uint16(anon_sym_COMMA),
	234: uint16(anon_sym_RBRACE),
	235: uint16(2),
	236: uint16(3),
	237: uint16(1),
	238: uint16(sym_comment),
	239: uint16(71),
	240: uint16(1),
	241: uint16(anon_sym_COLON),
	242: uint16(2),
	243: uint16(3),
	244: uint16(1),
	245: uint16(sym_comment),
	246: uint16(73),
	247: uint16(1),
	249: uint16(2),
	250: uint16(3),
	251: uint16(1),
	252: uint16(sym_comment),
	253: uint16(75),
	254: uint16(1),
}

var ts_small_parse_table_map = [24]uint32_t{
	1:  uint32(14),
	2:  uint32(24),
	3:  uint32(38),
	4:  uint32(48),
	5:  uint32(58),
	6:  uint32(68),
	7:  uint32(82),
	8:  uint32(92),
	9:  uint32(102),
	10: uint32(112),
	11: uint32(122),
	12: uint32(135),
	13: uint32(148),
	14: uint32(161),
	15: uint32(174),
	16: uint32(187),
	17: uint32(200),
	18: uint32(211),
	19: uint32(219),
	20: uint32(227),
	21: uint32(235),
	22: uint32(242),
	23: uint32(249),
}

var ts_parse_actions = [77]TSParseActionEntry{
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(2)),
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
		Fstate: uint16(libc.Int32FromInt32(10)),
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
		Fstate: uint16(libc.Int32FromInt32(19)),
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
		Fstate: uint16(libc.Int32FromInt32(15)),
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
		Fstate: uint16(libc.Int32FromInt32(25)),
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
		Fstate: uint16(libc.Int32FromInt32(17)),
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
		Fstate: uint16(libc.Int32FromInt32(26)),
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
		Fstate: uint16(libc.Int32FromInt32(8)),
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
		Fstate: uint16(libc.Int32FromInt32(28)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_object),
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
		Fstate: uint16(libc.Int32FromInt32(14)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_object),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(16)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_object),
	})))),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_array),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_object),
	})))),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_array),
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
		Fstate: uint16(libc.Int32FromInt32(9)),
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
		Fstate: uint16(libc.Int32FromInt32(11)),
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
		Fstate: uint16(libc.Int32FromInt32(3)),
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
		Fstate: uint16(libc.Int32FromInt32(12)),
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
		Fstate: uint16(libc.Int32FromInt32(13)),
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
		Fsymbol:      uint16(aux_sym_object_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(24)),
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
		Fcount:    uint8(1),
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
		Fsymbol:      uint16(aux_sym_object_repeat1),
	})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	65: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      uint16(libc.Int32FromInt32(6)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_member),
		Fproduction_id: uint16(1),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	74: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f uint8_t
		_ [7]byte
	}{f: uint8(TSParseActionTypeAccept)})),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_file),
	})))),
}

func tree_sitter_json5(tls *libc.TLS) (r uintptr) {
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
	Fname:                      __ccgo_ts + 137,
	Fmetadata: TSLanguageMetadata{
		Fminor_version: uint8(1),
	},
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

var __ccgo_ts1 = "end\x00comment\x00{\x00,\x00}\x00:\x00identifier\x00[\x00]\x00string\x00number\x00null\x00true\x00false\x00file\x00object\x00member\x00array\x00_value\x00object_repeat1\x00array_repeat1\x00name\x00value\x00json5\x00"
