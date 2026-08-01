// Code generated for windows/arm64 by 'ccgo -ignore-unsupported-alignment -ignore-unsupported-atomic-sizes -ignore-vector-functions --winapi-no-errno -ignore-link-errors -std=gnu11 -O0 -D__extension__= -D__forceinline=static inline -D__attribute__(...)= -D__declspec(x)= -D__cdecl= -D__stdcall= -D__fastcall= -D__thiscall= -D_cdecl= -D__restrict= -D__restrict__= -D__MINGW_EXTENSION= -D_X86INTRIN_H_INCLUDED -D_X86GPRINTRIN_H_INCLUDED -D_IMMINTRIN_H_INCLUDED -D_MMINTRIN_H_INCLUDED -D_XMMINTRIN_H_INCLUDED -D_EMMINTRIN_H_INCLUDED -D_PMMINTRIN_H_INCLUDED -D_MM3DNOW_H_INCLUDED -Dfunc=func_token -Dinterface=interface_token -Dselect=select_token -Dchan=chan_token -Dgo=go_token -Dmap=map_token -Dpackage=package_token -Dtype=type_token -Dvar=var_token -Dimport=import_token -Ddefer=defer_token -Dfallthrough=fallthrough_token -Drange=range_token -I C:\a\ccgo-tree-sitter\ccgo-tree-sitter\third-party\tree-sitter-wit\src -I C:\a\ccgo-tree-sitter\ccgo-tree-sitter\third-party\tree-sitter-wit -I C:\a\ccgo-tree-sitter\ccgo-tree-sitter\third-party -I C:\Users\runneradmin\.cache\workspaced\sources\github\7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2\lib\include -I C:\Users\runneradmin\.cache\workspaced\sources\github\7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2\lib\src combined.c -o grammar.go', DO NOT EDIT.

//go:build windows && arm64

package grammar_wit

import (
	"reflect"
	"unsafe"

	"modernc.org/libc"
)

var _ reflect.Type
var _ unsafe.Pointer

const ALIAS_COUNT = 2
const BUFSIZ = 512
const EXIT_FAILURE = 1
const EXIT_SUCCESS = 0
const EXTERNAL_TOKEN_COUNT = 4
const FIELD_COUNT = 10
const FILENAME_MAX = 260
const FOPEN_MAX = 20
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
const LARGE_STATE_COUNT = 2
const L_tmpnam = 260
const L_tmpnam_s = "L_tmpnam"
const MAX_ALIAS_SEQUENCE_LENGTH = 6
const MAX_RESERVED_WORD_SET_SIZE = 0
const MB_LEN_MAX = 1
const MINGW_HAS_DDK_H = 1
const MINGW_HAS_SECURE_API = 1
const PRODUCTION_ID_COUNT = 19
const PTRDIFF_MAX = "__PTRDIFF_MAX__"
const P_tmpdir = "_P_tmpdir"
const RAND_MAX = 0x7fff
const SEEK_CUR = 1
const SEEK_END = 2
const SEEK_SET = 0
const SIZE_MAX = "__SIZE_MAX__"
const STATE_COUNT = 349
const STDERR_FILENO = 2
const STDIN_FILENO = 0
const STDOUT_FILENO = 1
const SUPERTYPE_COUNT = 2
const SYMBOL_COUNT = 168
const SYS_OPEN = "_SYS_OPEN"
const TMP_MAX = 2147483647
const TMP_MAX_S = "TMP_MAX"
const TOKEN_COUNT = 77
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
const _CRT_INTERNAL_PRINTF_LEGACY_MSVCRT_COMPATIBILITY = "0x0008U"
const _CRT_INTERNAL_PRINTF_LEGACY_THREE_DIGIT_EXPONENTS = "0x0010U"
const _CRT_INTERNAL_PRINTF_LEGACY_VSPRINTF_NULL_TERMINATION = 1
const _CRT_INTERNAL_PRINTF_LEGACY_WIDE_SPECIFIERS = "0x0004U"
const _CRT_INTERNAL_PRINTF_STANDARD_ROUNDING = "0x0020U"
const _CRT_INTERNAL_PRINTF_STANDARD_SNPRINTF_BEHAVIOR = 2
const _CRT_INTERNAL_SCANF_LEGACY_MSVCRT_COMPATIBILITY = "0x0004U"
const _CRT_INTERNAL_SCANF_LEGACY_WIDE_SPECIFIERS = "0x0002U"
const _CRT_INTERNAL_SCANF_SECURECRT = 1
const _DIGIT = 0x4
const _EMMINTRIN_H_INCLUDED = 1
const _FREEENTRY = 0
const _HEAP_MAXREQ = 0xFFFFFFFFFFFFFFE0
const _HEX = 0x80
const _IMMINTRIN_H_INCLUDED = 1
const _IOB_ENTRIES = 20
const _IOFBF = 0x0000
const _IOLBF = 0x0040
const _IONBF = 0x0004
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
const _NFILE = "_NSTREAM_"
const _NSTREAM_ = 512
const _OLD_P_OVERLAY = 2
const _OUT_TO_DEFAULT = 0
const _OUT_TO_MSGBOX = 2
const _OUT_TO_STDERR = 1
const _PMMINTRIN_H_INCLUDED = 1
const _PUNCT = 0x10
const _P_DETACH = 4
const _P_NOWAIT = 1
const _P_NOWAITO = 3
const _P_OVERLAY = 2
const _P_WAIT = 0
const _P_tmpdir = "\\\\"
const _REPORT_ERRMODE = 3
const _SECURECRT_FILL_BUFFER_PATTERN = 0xFD
const _SPACE = 8
const _SYS_OPEN = 20
const _UPPER = 0x1
const _USEDENTRY = 1
const _WAIT_CHILD = 0
const _WAIT_GRANDCHILD = 1
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
const _wP_tmpdir = "\\\\"
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
const pclose1 = "_pclose"
const popen1 = "_popen"
const range1 = "range_token"
const select2 = "select_token"
const sys_errlist = "_sys_errlist"
const sys_nerr = "_sys_nerr"
const true1 = 1
const ts_builtin_sym_end = 0
const ts_calloc = "calloc"
const ts_free = "free"
const ts_malloc = "malloc"
const ts_realloc = "realloc"
const type1 = "type_token"
const var1 = "var_token"
const wpopen = "_wpopen"

type __builtin_va_list = uintptr

type __predefined_size_t = uint64

type __predefined_wchar_t = uint16

type __predefined_ptrdiff_t = int64

type __gnuc_va_list = uintptr

type va_list = uintptr

type size_t = uint64

type ssize_t = int64

type rsize_t = uint64

type intptr_t = int64

type uintptr_t = uint64

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

type _iobuf = struct {
	F_Placeholder uintptr
}

type FILE = struct {
	F_Placeholder uintptr
}

type _off_t = int32

type off32_t = int32

type _off64_t = int64

type off64_t = int64

type off_t = int32

type fpos_t = int64

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

type intmax_t = int64

type uintmax_t = uint64

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

const BLOCK_COMMENT_CONTENT = 0
const BLOCK_DOC_MARKER = 1
const ERROR_SENTINEL = 2
const LINE_DOC_CONTENT = 3

func tree_sitter_wit_external_scanner_create(tls *libc.TLS) (r uintptr) {
	return r
}

func tree_sitter_wit_external_scanner_destroy(tls *libc.TLS, payload uintptr) {
}

func tree_sitter_wit_external_scanner_serialize(tls *libc.TLS, payload uintptr, buffer uintptr) (r uint32) {
	return uint32(1)
}

func tree_sitter_wit_external_scanner_deserialize(tls *libc.TLS, payload uintptr, buffer uintptr, length uint32) {
}

func advance(tls *libc.TLS, lexer uintptr) {
	(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(false1 != 0))
}

func skip(tls *libc.TLS, lexer uintptr) {
	(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(true1 != 0))
}

func process_line_doc_content(tls *libc.TLS, lexer uintptr) (r uint8) {
	(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(LINE_DOC_CONTENT)
	for {
		if (*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Feof})))(tls, lexer) != 0 {
			return libc.BoolUint8(true1 != 0)
		}
		if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('\n') {
			// Include the newline in the doc content node.
			// Line endings are useful for markdown injection.
			advance(tls, lexer)
			return libc.BoolUint8(true1 != 0)
		}
		advance(tls, lexer)
		goto _1
	_1:
	}
	return r
}

type BlockCommentState = int32

const LeftForwardSlash = 0
const LeftAsterisk = 1
const Continuing = 2

type BlockCommentProcessing = struct {
	Fstate        BlockCommentState
	FnestingDepth uint32
}

func process_left_forward_slash(tls *libc.TLS, processing uintptr, current int8) {
	if int32(current) == int32('*') {
		*(*uint32)(unsafe.Pointer(processing + 4)) += uint32(1)
	}
	(*BlockCommentProcessing)(unsafe.Pointer(processing)).Fstate = int32(Continuing)
}

func process_left_asterisk(tls *libc.TLS, processing uintptr, current int8, lexer uintptr) {
	if int32(current) == int32('*') {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		(*BlockCommentProcessing)(unsafe.Pointer(processing)).Fstate = int32(LeftAsterisk)
		return
	}
	if int32(current) == int32('/') {
		*(*uint32)(unsafe.Pointer(processing + 4)) -= uint32(1)
	}
	(*BlockCommentProcessing)(unsafe.Pointer(processing)).Fstate = int32(Continuing)
}

func process_continuing(tls *libc.TLS, processing uintptr, current int8) {
	switch int32(current) {
	case int32('/'):
		(*BlockCommentProcessing)(unsafe.Pointer(processing)).Fstate = int32(LeftForwardSlash)
	case int32('*'):
		(*BlockCommentProcessing)(unsafe.Pointer(processing)).Fstate = int32(LeftAsterisk)
		break
	}
}

func process_block_comment(tls *libc.TLS, lexer uintptr, valid_symbols uintptr) (r uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var first int8
	var _ /* processing at bp+0 */ BlockCommentProcessing
	_ = first
	// The first character is stored so we can safely advance inside
	// these if blocks. However, because we only store one, we can only
	// safely advance 1 time.
	first = int8((*TSLexer)(unsafe.Pointer(lexer)).Flookahead)
	if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(BLOCK_DOC_MARKER))) != 0 && int32(first) == int32('*') {
		advance(tls, lexer)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(BLOCK_DOC_MARKER)
		// If the next token is a / that means that it's an empty block comment:
		// /**/
		if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('/') {
			return libc.BoolUint8(false1 != 0)
		}
		// If the next token is a * that means that this isn't a BLOCK_DOC_MARKER
		// as BLOCK_DOC_MARKER's only have 2 * not 3 or more.
		return libc.BoolUint8(true1 != 0)
	} else {
		// Since there's a chance that an advance could
		// happen in one state, we must advance in all states to ensure that
		// the program ends up in a sane state prior to processing the block
		// comment if need be.
		advance(tls, lexer)
	}
	if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(BLOCK_COMMENT_CONTENT))) != 0 {
		*(*BlockCommentProcessing)(unsafe.Pointer(bp)) = BlockCommentProcessing{
			Fstate:        int32(Continuing),
			FnestingDepth: uint32(1),
		}
		// Manually set the current state based on the first character
		switch int32(first) {
		case int32('*'):
			(*(*BlockCommentProcessing)(unsafe.Pointer(bp))).Fstate = int32(LeftAsterisk)
			if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('/') {
				// This case can happen in an empty doc block comment
				// like /** */. The comment has no contents, so bail.
				return libc.BoolUint8(false1 != 0)
			}
		case int32('/'):
			(*(*BlockCommentProcessing)(unsafe.Pointer(bp))).Fstate = int32(LeftForwardSlash)
		default:
			(*(*BlockCommentProcessing)(unsafe.Pointer(bp))).Fstate = int32(Continuing)
			break
		}
		// For the purposes of actually parsing WIT code, this
		// is incorrect as it considers an unterminated block comment
		// to be an error. However, for the purposes of syntax highlighting
		// this should be considered successful as otherwise you are not able
		// to syntax highlight a block of code prior to closing the
		// block comment
		for !((*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Feof})))(tls, lexer) != 0) && (*(*BlockCommentProcessing)(unsafe.Pointer(bp))).FnestingDepth != uint32(0) {
			// Set first to the current lookahead as that is the second character
			// as we force an advance in the above code when we are checking if we
			// need to handle a block comment inner or outer doc comment signifier
			// node
			first = int8((*TSLexer)(unsafe.Pointer(lexer)).Flookahead)
			switch (*(*BlockCommentProcessing)(unsafe.Pointer(bp))).Fstate {
			case int32(LeftForwardSlash):
				process_left_forward_slash(tls, bp, first)
			case int32(LeftAsterisk):
				process_left_asterisk(tls, bp, first, lexer)
			case int32(Continuing):
				(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
				process_continuing(tls, bp, first)
			default:
				break
			}
			advance(tls, lexer)
			if int32(first) == int32('/') && (*(*BlockCommentProcessing)(unsafe.Pointer(bp))).FnestingDepth != uint32(0) {
				(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
			}
		}
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(BLOCK_COMMENT_CONTENT)
		return libc.BoolUint8(true1 != 0)
	}
	return libc.BoolUint8(false1 != 0)
}

func tree_sitter_wit_external_scanner_scan(tls *libc.TLS, payload uintptr, lexer uintptr, valid_symbols uintptr) (r uint8) {
	// The documentation states that if the lexical analysis fails for some reason
	// they will mark every state as valid and pass it to the external scanner
	// However, we can't do anything to help them recover in that case so we
	// should just fail.
	/*
	   link: https://tree-sitter.github.io/tree-sitter/creating-parsers/4-external-scanners.html
	   If a syntax error is encountered during regular parsing, Tree-sitter’s
	   first action during error recovery will be to call the external scanner’s
	   scan function with all tokens marked valid. The scanner should detect this
	   case and handle it appropriately. One simple method of detection is to add
	   an unused token to the end of the externals array, for example
	   externals: $ => [$.token1, $.token2, $.error_sentinel],
	   then check whether that token is marked valid to determine whether
	   Tree-sitter is in error correction mode.
	*/
	if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(ERROR_SENTINEL))) != 0 {
		return libc.BoolUint8(false1 != 0)
	}
	if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(BLOCK_COMMENT_CONTENT))) != 0 || *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(BLOCK_DOC_MARKER))) != 0 {
		return process_block_comment(tls, lexer, valid_symbols)
	}
	if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(LINE_DOC_CONTENT))) != 0 {
		return process_line_doc_content(tls, lexer)
	}
	for libc.Xiswctype(tls, uint16((*TSLexer)(unsafe.Pointer(lexer)).Flookahead), uint16(_SPACE)) != 0 {
		skip(tls, lexer)
	}
	return libc.BoolUint8(false1 != 0)
}

/* Automatically @generated by tree-sitter */

type ts_symbol_identifiers = int32

const sym_id = 1
const anon_sym_LBRACE = 2
const anon_sym_RBRACE = 3
const anon_sym_COLON = 4
const anon_sym_SLASH = 5
const anon_sym_AT = 6
const anon_sym_package = 7
const anon_sym_SEMI = 8
const anon_sym_use = 9
const anon_sym_as = 10
const sym_string_literal = 11
const sym__valid_semver = 12
const anon_sym_world = 13
const anon_sym_export = 14
const anon_sym_import = 15
const anon_sym_interface = 16
const anon_sym_include = 17
const anon_sym_with = 18
const anon_sym_COMMA = 19
const anon_sym_async = 20
const anon_sym_func = 21
const anon_sym_LPAREN = 22
const anon_sym_RPAREN = 23
const anon_sym_DASH_GT = 24
const anon_sym_DOT = 25
const anon_sym_type = 26
const anon_sym_EQ = 27
const anon_sym_record = 28
const anon_sym_flags = 29
const anon_sym_variant = 30
const anon_sym_enum = 31
const anon_sym_resource = 32
const anon_sym_static = 33
const anon_sym_constructor = 34
const anon_sym_u8 = 35
const anon_sym_u16 = 36
const anon_sym_u32 = 37
const anon_sym_u64 = 38
const anon_sym_s8 = 39
const anon_sym_s16 = 40
const anon_sym_s32 = 41
const anon_sym_s64 = 42
const anon_sym_char = 43
const anon_sym_bool = 44
const anon_sym_string = 45
const anon_sym_f32 = 46
const anon_sym_f64 = 47
const anon_sym_tuple = 48
const anon_sym_LT = 49
const anon_sym_GT = 50
const sym_uint = 51
const anon_sym_list = 52
const anon_sym_option = 53
const anon_sym_map = 54
const anon_sym_result = 55
const anon_sym__ = 56
const anon_sym_borrow = 57
const anon_sym_future = 58
const anon_sym_stream = 59
const anon_sym_SLASH_SLASH = 60
const anon_sym_SLASH_SLASH2 = 61
const aux_sym_line_comment_token1 = 62
const anon_sym_SLASH2 = 63
const aux_sym_line_comment_token2 = 64
const anon_sym_SLASH_STAR = 65
const anon_sym_STAR_SLASH = 66
const anon_sym_external_DASHid = 67
const anon_sym_unstable = 68
const anon_sym_feature = 69
const anon_sym_since = 70
const anon_sym_deprecated = 71
const anon_sym_version = 72
const sym__block_comment_content = 73
const sym__block_doc_comment_marker = 74
const sym__error_sentinel = 75
const sym__line_doc_content = 76
const sym_source_file = 77
const sym__statement = 78
const sym__package_items = 79
const sym_nested_package_definition = 80
const sym__uri_head = 81
const sym__uri_tail = 82
const sym__version = 83
const sym_decl_head = 84
const sym_package_decl = 85
const sym_toplevel_use_item = 86
const sym_use_path = 87
const sym_world_item = 88
const sym__world_body = 89
const sym__world_items = 90
const sym_export_item = 91
const sym_import_item = 92
const sym_extern_type = 93
const sym_include_item = 94
const sym__include_names_body = 95
const sym__include_names_list = 96
const sym_include_names_item = 97
const sym_alias_item = 98
const sym_interface_item = 99
const sym__interface_body = 100
const sym__interface_items = 101
const sym__typedef_item = 102
const sym_func_item = 103
const sym_func_type = 104
const sym_param_list = 105
const sym_result_list = 106
const sym__named_type_list = 107
const sym_named_type = 108
const sym_use_item = 109
const sym__use_names_body = 110
const sym__use_names_list = 111
const sym_use_names_item = 112
const sym_type_item = 113
const sym_record_item = 114
const sym__record_body = 115
const sym__record_fields = 116
const sym_record_field = 117
const sym_flags_items = 118
const sym__flags_body = 119
const sym__flags_fields = 120
const sym_variant_items = 121
const sym__variant_body = 122
const sym__variant_cases = 123
const sym_variant_case = 124
const sym_enum_items = 125
const sym__enum_body = 126
const sym__enum_cases = 127
const sym_resource_item = 128
const sym__resource_body = 129
const sym_resource_method = 130
const sym__primitive_ty = 131
const sym_ty = 132
const sym_kt = 133
const sym_tuple = 134
const sym_tuple_list = 135
const sym_list = 136
const sym_option = 137
const sym_map = 138
const sym_result = 139
const sym_handle = 140
const sym_future = 141
const sym_stream = 142
const sym_line_comment = 143
const sym_block_comment = 144
const aux_sym__gate = 145
const sym__gate_item = 146
const sym_external_id = 147
const sym_unstable_gate = 148
const sym__feature_field = 149
const sym_since_gate = 150
const sym_deprecated_gate = 151
const sym__version_field = 152
const aux_sym_source_file_repeat1 = 153
const aux_sym_nested_package_definition_repeat1 = 154
const aux_sym_decl_head_repeat1 = 155
const aux_sym_decl_head_repeat2 = 156
const aux_sym__world_body_repeat1 = 157
const aux_sym__include_names_list_repeat1 = 158
const aux_sym__interface_body_repeat1 = 159
const aux_sym__named_type_list_repeat1 = 160
const aux_sym__use_names_list_repeat1 = 161
const aux_sym__record_fields_repeat1 = 162
const aux_sym__flags_fields_repeat1 = 163
const aux_sym__variant_cases_repeat1 = 164
const aux_sym__enum_cases_repeat1 = 165
const aux_sym__resource_body_repeat1 = 166
const aux_sym_tuple_list_repeat1 = 167
const alias_sym_enum_case = 168
const alias_sym_flags_field = 169

var ts_symbol_names = [170]uintptr{
	0:   __ccgo_ts,
	1:   __ccgo_ts + 4,
	2:   __ccgo_ts + 7,
	3:   __ccgo_ts + 9,
	4:   __ccgo_ts + 11,
	5:   __ccgo_ts + 13,
	6:   __ccgo_ts + 15,
	7:   __ccgo_ts + 17,
	8:   __ccgo_ts + 25,
	9:   __ccgo_ts + 27,
	10:  __ccgo_ts + 31,
	11:  __ccgo_ts + 34,
	12:  __ccgo_ts + 49,
	13:  __ccgo_ts + 57,
	14:  __ccgo_ts + 63,
	15:  __ccgo_ts + 70,
	16:  __ccgo_ts + 77,
	17:  __ccgo_ts + 87,
	18:  __ccgo_ts + 95,
	19:  __ccgo_ts + 100,
	20:  __ccgo_ts + 102,
	21:  __ccgo_ts + 108,
	22:  __ccgo_ts + 113,
	23:  __ccgo_ts + 115,
	24:  __ccgo_ts + 117,
	25:  __ccgo_ts + 120,
	26:  __ccgo_ts + 122,
	27:  __ccgo_ts + 127,
	28:  __ccgo_ts + 129,
	29:  __ccgo_ts + 136,
	30:  __ccgo_ts + 142,
	31:  __ccgo_ts + 150,
	32:  __ccgo_ts + 155,
	33:  __ccgo_ts + 164,
	34:  __ccgo_ts + 171,
	35:  __ccgo_ts + 183,
	36:  __ccgo_ts + 186,
	37:  __ccgo_ts + 190,
	38:  __ccgo_ts + 194,
	39:  __ccgo_ts + 198,
	40:  __ccgo_ts + 201,
	41:  __ccgo_ts + 205,
	42:  __ccgo_ts + 209,
	43:  __ccgo_ts + 213,
	44:  __ccgo_ts + 218,
	45:  __ccgo_ts + 223,
	46:  __ccgo_ts + 230,
	47:  __ccgo_ts + 234,
	48:  __ccgo_ts + 238,
	49:  __ccgo_ts + 244,
	50:  __ccgo_ts + 246,
	51:  __ccgo_ts + 248,
	52:  __ccgo_ts + 253,
	53:  __ccgo_ts + 258,
	54:  __ccgo_ts + 265,
	55:  __ccgo_ts + 269,
	56:  __ccgo_ts + 276,
	57:  __ccgo_ts + 278,
	58:  __ccgo_ts + 285,
	59:  __ccgo_ts + 292,
	60:  __ccgo_ts + 299,
	61:  __ccgo_ts + 299,
	62:  __ccgo_ts + 302,
	63:  __ccgo_ts + 13,
	64:  __ccgo_ts + 322,
	65:  __ccgo_ts + 342,
	66:  __ccgo_ts + 345,
	67:  __ccgo_ts + 348,
	68:  __ccgo_ts + 360,
	69:  __ccgo_ts + 369,
	70:  __ccgo_ts + 377,
	71:  __ccgo_ts + 383,
	72:  __ccgo_ts + 49,
	73:  __ccgo_ts + 394,
	74:  __ccgo_ts + 417,
	75:  __ccgo_ts + 443,
	76:  __ccgo_ts + 459,
	77:  __ccgo_ts + 471,
	78:  __ccgo_ts + 483,
	79:  __ccgo_ts + 494,
	80:  __ccgo_ts + 509,
	81:  __ccgo_ts + 535,
	82:  __ccgo_ts + 545,
	83:  __ccgo_ts + 555,
	84:  __ccgo_ts + 564,
	85:  __ccgo_ts + 574,
	86:  __ccgo_ts + 587,
	87:  __ccgo_ts + 605,
	88:  __ccgo_ts + 614,
	89:  __ccgo_ts + 625,
	90:  __ccgo_ts + 630,
	91:  __ccgo_ts + 643,
	92:  __ccgo_ts + 655,
	93:  __ccgo_ts + 667,
	94:  __ccgo_ts + 679,
	95:  __ccgo_ts + 692,
	96:  __ccgo_ts + 704,
	97:  __ccgo_ts + 724,
	98:  __ccgo_ts + 743,
	99:  __ccgo_ts + 754,
	100: __ccgo_ts + 625,
	101: __ccgo_ts + 769,
	102: __ccgo_ts + 786,
	103: __ccgo_ts + 800,
	104: __ccgo_ts + 810,
	105: __ccgo_ts + 820,
	106: __ccgo_ts + 831,
	107: __ccgo_ts + 843,
	108: __ccgo_ts + 860,
	109: __ccgo_ts + 871,
	110: __ccgo_ts + 692,
	111: __ccgo_ts + 880,
	112: __ccgo_ts + 896,
	113: __ccgo_ts + 911,
	114: __ccgo_ts + 921,
	115: __ccgo_ts + 625,
	116: __ccgo_ts + 933,
	117: __ccgo_ts + 948,
	118: __ccgo_ts + 961,
	119: __ccgo_ts + 625,
	120: __ccgo_ts + 973,
	121: __ccgo_ts + 987,
	122: __ccgo_ts + 625,
	123: __ccgo_ts + 1001,
	124: __ccgo_ts + 1016,
	125: __ccgo_ts + 1029,
	126: __ccgo_ts + 625,
	127: __ccgo_ts + 1040,
	128: __ccgo_ts + 1052,
	129: __ccgo_ts + 625,
	130: __ccgo_ts + 1066,
	131: __ccgo_ts + 1082,
	132: __ccgo_ts + 1096,
	133: __ccgo_ts + 1099,
	134: __ccgo_ts + 238,
	135: __ccgo_ts + 1102,
	136: __ccgo_ts + 253,
	137: __ccgo_ts + 258,
	138: __ccgo_ts + 265,
	139: __ccgo_ts + 269,
	140: __ccgo_ts + 1113,
	141: __ccgo_ts + 285,
	142: __ccgo_ts + 292,
	143: __ccgo_ts + 1120,
	144: __ccgo_ts + 1133,
	145: __ccgo_ts + 1147,
	146: __ccgo_ts + 1153,
	147: __ccgo_ts + 1164,
	148: __ccgo_ts + 1176,
	149: __ccgo_ts + 1190,
	150: __ccgo_ts + 1205,
	151: __ccgo_ts + 1216,
	152: __ccgo_ts + 1232,
	153: __ccgo_ts + 1247,
	154: __ccgo_ts + 1267,
	155: __ccgo_ts + 1301,
	156: __ccgo_ts + 1319,
	157: __ccgo_ts + 1337,
	158: __ccgo_ts + 1357,
	159: __ccgo_ts + 1385,
	160: __ccgo_ts + 1409,
	161: __ccgo_ts + 1434,
	162: __ccgo_ts + 1458,
	163: __ccgo_ts + 1481,
	164: __ccgo_ts + 1503,
	165: __ccgo_ts + 1526,
	166: __ccgo_ts + 1546,
	167: __ccgo_ts + 1569,
	168: __ccgo_ts + 1588,
	169: __ccgo_ts + 1598,
}

var ts_symbol_map = [170]TSSymbol{
	1:   uint16(sym_id),
	2:   uint16(anon_sym_LBRACE),
	3:   uint16(anon_sym_RBRACE),
	4:   uint16(anon_sym_COLON),
	5:   uint16(anon_sym_SLASH),
	6:   uint16(anon_sym_AT),
	7:   uint16(anon_sym_package),
	8:   uint16(anon_sym_SEMI),
	9:   uint16(anon_sym_use),
	10:  uint16(anon_sym_as),
	11:  uint16(sym_string_literal),
	12:  uint16(sym__valid_semver),
	13:  uint16(anon_sym_world),
	14:  uint16(anon_sym_export),
	15:  uint16(anon_sym_import),
	16:  uint16(anon_sym_interface),
	17:  uint16(anon_sym_include),
	18:  uint16(anon_sym_with),
	19:  uint16(anon_sym_COMMA),
	20:  uint16(anon_sym_async),
	21:  uint16(anon_sym_func),
	22:  uint16(anon_sym_LPAREN),
	23:  uint16(anon_sym_RPAREN),
	24:  uint16(anon_sym_DASH_GT),
	25:  uint16(anon_sym_DOT),
	26:  uint16(anon_sym_type),
	27:  uint16(anon_sym_EQ),
	28:  uint16(anon_sym_record),
	29:  uint16(anon_sym_flags),
	30:  uint16(anon_sym_variant),
	31:  uint16(anon_sym_enum),
	32:  uint16(anon_sym_resource),
	33:  uint16(anon_sym_static),
	34:  uint16(anon_sym_constructor),
	35:  uint16(anon_sym_u8),
	36:  uint16(anon_sym_u16),
	37:  uint16(anon_sym_u32),
	38:  uint16(anon_sym_u64),
	39:  uint16(anon_sym_s8),
	40:  uint16(anon_sym_s16),
	41:  uint16(anon_sym_s32),
	42:  uint16(anon_sym_s64),
	43:  uint16(anon_sym_char),
	44:  uint16(anon_sym_bool),
	45:  uint16(anon_sym_string),
	46:  uint16(anon_sym_f32),
	47:  uint16(anon_sym_f64),
	48:  uint16(anon_sym_tuple),
	49:  uint16(anon_sym_LT),
	50:  uint16(anon_sym_GT),
	51:  uint16(sym_uint),
	52:  uint16(anon_sym_list),
	53:  uint16(anon_sym_option),
	54:  uint16(anon_sym_map),
	55:  uint16(anon_sym_result),
	56:  uint16(anon_sym__),
	57:  uint16(anon_sym_borrow),
	58:  uint16(anon_sym_future),
	59:  uint16(anon_sym_stream),
	60:  uint16(anon_sym_SLASH_SLASH),
	61:  uint16(anon_sym_SLASH_SLASH),
	62:  uint16(aux_sym_line_comment_token1),
	63:  uint16(anon_sym_SLASH),
	64:  uint16(aux_sym_line_comment_token2),
	65:  uint16(anon_sym_SLASH_STAR),
	66:  uint16(anon_sym_STAR_SLASH),
	67:  uint16(anon_sym_external_DASHid),
	68:  uint16(anon_sym_unstable),
	69:  uint16(anon_sym_feature),
	70:  uint16(anon_sym_since),
	71:  uint16(anon_sym_deprecated),
	72:  uint16(anon_sym_version),
	73:  uint16(sym__block_comment_content),
	74:  uint16(sym__block_doc_comment_marker),
	75:  uint16(sym__error_sentinel),
	76:  uint16(sym__line_doc_content),
	77:  uint16(sym_source_file),
	78:  uint16(sym__statement),
	79:  uint16(sym__package_items),
	80:  uint16(sym_nested_package_definition),
	81:  uint16(sym__uri_head),
	82:  uint16(sym__uri_tail),
	83:  uint16(sym__version),
	84:  uint16(sym_decl_head),
	85:  uint16(sym_package_decl),
	86:  uint16(sym_toplevel_use_item),
	87:  uint16(sym_use_path),
	88:  uint16(sym_world_item),
	89:  uint16(sym__world_body),
	90:  uint16(sym__world_items),
	91:  uint16(sym_export_item),
	92:  uint16(sym_import_item),
	93:  uint16(sym_extern_type),
	94:  uint16(sym_include_item),
	95:  uint16(sym__include_names_body),
	96:  uint16(sym__include_names_list),
	97:  uint16(sym_include_names_item),
	98:  uint16(sym_alias_item),
	99:  uint16(sym_interface_item),
	100: uint16(sym__world_body),
	101: uint16(sym__interface_items),
	102: uint16(sym__typedef_item),
	103: uint16(sym_func_item),
	104: uint16(sym_func_type),
	105: uint16(sym_param_list),
	106: uint16(sym_result_list),
	107: uint16(sym__named_type_list),
	108: uint16(sym_named_type),
	109: uint16(sym_use_item),
	110: uint16(sym__include_names_body),
	111: uint16(sym__use_names_list),
	112: uint16(sym_use_names_item),
	113: uint16(sym_type_item),
	114: uint16(sym_record_item),
	115: uint16(sym__world_body),
	116: uint16(sym__record_fields),
	117: uint16(sym_record_field),
	118: uint16(sym_flags_items),
	119: uint16(sym__world_body),
	120: uint16(sym__flags_fields),
	121: uint16(sym_variant_items),
	122: uint16(sym__world_body),
	123: uint16(sym__variant_cases),
	124: uint16(sym_variant_case),
	125: uint16(sym_enum_items),
	126: uint16(sym__world_body),
	127: uint16(sym__enum_cases),
	128: uint16(sym_resource_item),
	129: uint16(sym__world_body),
	130: uint16(sym_resource_method),
	131: uint16(sym__primitive_ty),
	132: uint16(sym_ty),
	133: uint16(sym_kt),
	134: uint16(sym_tuple),
	135: uint16(sym_tuple_list),
	136: uint16(sym_list),
	137: uint16(sym_option),
	138: uint16(sym_map),
	139: uint16(sym_result),
	140: uint16(sym_handle),
	141: uint16(sym_future),
	142: uint16(sym_stream),
	143: uint16(sym_line_comment),
	144: uint16(sym_block_comment),
	145: uint16(aux_sym__gate),
	146: uint16(sym__gate_item),
	147: uint16(sym_external_id),
	148: uint16(sym_unstable_gate),
	149: uint16(sym__feature_field),
	150: uint16(sym_since_gate),
	151: uint16(sym_deprecated_gate),
	152: uint16(sym__version_field),
	153: uint16(aux_sym_source_file_repeat1),
	154: uint16(aux_sym_nested_package_definition_repeat1),
	155: uint16(aux_sym_decl_head_repeat1),
	156: uint16(aux_sym_decl_head_repeat2),
	157: uint16(aux_sym__world_body_repeat1),
	158: uint16(aux_sym__include_names_list_repeat1),
	159: uint16(aux_sym__interface_body_repeat1),
	160: uint16(aux_sym__named_type_list_repeat1),
	161: uint16(aux_sym__use_names_list_repeat1),
	162: uint16(aux_sym__record_fields_repeat1),
	163: uint16(aux_sym__flags_fields_repeat1),
	164: uint16(aux_sym__variant_cases_repeat1),
	165: uint16(aux_sym__enum_cases_repeat1),
	166: uint16(aux_sym__resource_body_repeat1),
	167: uint16(aux_sym_tuple_list_repeat1),
	168: uint16(alias_sym_enum_case),
	169: uint16(alias_sym_flags_field),
}

var ts_symbol_metadata = [170]TSSymbolMetadata{
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
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	12: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
		Fnamed:   libc.BoolUint8(true1 != 0),
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
	62: {},
	63: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	64: {},
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
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	74: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	75: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	76: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	77: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	78: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	79: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	80: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	81: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	82: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	83: {
		Fnamed: libc.BoolUint8(true1 != 0),
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
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	102: {
		Fnamed:     libc.BoolUint8(true1 != 0),
		Fsupertype: libc.BoolUint8(true1 != 0),
	},
	103: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	104: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	105: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	106: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	107: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	108: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	109: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	110: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	111: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	112: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	113: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	114: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	115: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	116: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	117: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	118: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	119: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	120: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	121: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	122: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	123: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	124: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	125: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	126: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	127: {
		Fnamed: libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	132: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	133: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	134: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	142: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	143: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	144: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	145: {},
	146: {
		Fnamed:     libc.BoolUint8(true1 != 0),
		Fsupertype: libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	153: {},
	154: {},
	155: {},
	156: {},
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
	167: {},
	168: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	169: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
}

type ts_field_identifiers = int32

const field_alias = 1
const field_doc = 2
const field_feature = 3
const field_id = 4
const field_key = 5
const field_name = 6
const field_path = 7
const field_size = 8
const field_type = 9
const field_value = 10

var ts_field_names = [11]uintptr{
	0:  libc.UintptrFromInt32(0),
	1:  __ccgo_ts + 1610,
	2:  __ccgo_ts + 1616,
	3:  __ccgo_ts + 369,
	4:  __ccgo_ts + 4,
	5:  __ccgo_ts + 1620,
	6:  __ccgo_ts + 1624,
	7:  __ccgo_ts + 1629,
	8:  __ccgo_ts + 1634,
	9:  __ccgo_ts + 122,
	10: __ccgo_ts + 1639,
}

var ts_field_map_slices = [19]TSMapSlice{
	1: {
		Flength: uint16(1),
	},
	2: {
		Findex:  uint16(1),
		Flength: uint16(1),
	},
	3: {
		Flength: uint16(1),
	},
	4: {
		Findex:  uint16(2),
		Flength: uint16(1),
	},
	5: {
		Findex:  uint16(3),
		Flength: uint16(1),
	},
	6: {
		Findex:  uint16(4),
		Flength: uint16(1),
	},
	7: {
		Findex:  uint16(5),
		Flength: uint16(1),
	},
	9: {
		Findex:  uint16(6),
		Flength: uint16(1),
	},
	11: {
		Findex:  uint16(7),
		Flength: uint16(1),
	},
	12: {
		Findex:  uint16(8),
		Flength: uint16(2),
	},
	13: {
		Findex:  uint16(10),
		Flength: uint16(2),
	},
	16: {
		Findex:  uint16(12),
		Flength: uint16(2),
	},
	17: {
		Findex:  uint16(14),
		Flength: uint16(2),
	},
	18: {
		Findex:  uint16(16),
		Flength: uint16(2),
	},
}

var ts_field_map_entries = [18]TSFieldMapEntry{
	0: {
		Ffield_id:    uint16(field_doc),
		Fchild_index: uint8(2),
	},
	1: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
	},
	2: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(2),
	},
	3: {
		Ffield_id:    uint16(field_feature),
		Fchild_index: uint8(3),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	4: {
		Ffield_id:    uint16(field_alias),
		Fchild_index: uint8(3),
	},
	5: {
		Ffield_id:    uint16(field_feature),
		Fchild_index: uint8(2),
	},
	6: {
		Ffield_id: uint16(field_name),
	},
	7: {
		Ffield_id:    uint16(field_id),
		Fchild_index: uint8(3),
	},
	8: {
		Ffield_id:    uint16(field_alias),
		Fchild_index: uint8(1),
	},
	9: {
		Ffield_id:    uint16(field_type),
		Fchild_index: uint8(3),
	},
	10: {
		Ffield_id: uint16(field_name),
	},
	11: {
		Ffield_id:    uint16(field_type),
		Fchild_index: uint8(2),
	},
	12: {
		Ffield_id:    uint16(field_alias),
		Fchild_index: uint8(2),
	},
	13: {
		Ffield_id: uint16(field_path),
	},
	14: {
		Ffield_id:    uint16(field_size),
		Fchild_index: uint8(3),
	},
	15: {
		Ffield_id:    uint16(field_size),
		Fchild_index: uint8(4),
	},
	16: {
		Ffield_id:    uint16(field_key),
		Fchild_index: uint8(2),
	},
	17: {
		Ffield_id:    uint16(field_value),
		Fchild_index: uint8(4),
	},
}

var ts_alias_sequences = [19][6]TSSymbol{
	0: {},
	3: {
		2: uint16(sym__line_doc_content),
	},
	8: {
		0: uint16(alias_sym_flags_field),
	},
	10: {
		0: uint16(alias_sym_enum_case),
	},
	14: {
		1: uint16(alias_sym_flags_field),
	},
	15: {
		1: uint16(alias_sym_enum_case),
	},
}

var ts_non_terminal_alias_map = [1]uint16_t{}

var ts_primary_state_ids = [349]TSStateId{
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
	288: uint16(288),
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
	339: uint16(339),
	340: uint16(340),
	341: uint16(341),
	342: uint16(342),
	343: uint16(343),
	344: uint16(344),
	345: uint16(345),
	346: uint16(346),
	347: uint16(347),
	348: uint16(348),
}

var ts_supertype_symbols = [2]TSSymbol{
	0: uint16(sym__gate_item),
	1: uint16(sym__typedef_item),
}

var ts_supertype_map_slices = [147]TSMapSlice{
	102: {
		Findex:  uint16(3),
		Flength: uint16(6),
	},
	146: {
		Flength: uint16(3),
	},
}

var ts_supertype_map_entries = [9]TSSymbol{
	0: uint16(sym_deprecated_gate),
	1: uint16(sym_since_gate),
	2: uint16(sym_unstable_gate),
	3: uint16(sym_enum_items),
	4: uint16(sym_flags_items),
	5: uint16(sym_record_item),
	6: uint16(sym_resource_item),
	7: uint16(sym_type_item),
	8: uint16(sym_variant_items),
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
	switch int32(state) {
	case 0:
		if eof != 0 {
			state = uint16(20)
			goto next_state
		}
		i = uint32(0)
		for {
			if !(uint64(i) < libc.Uint64FromInt64(76)/libc.Uint64FromInt64(2)) {
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
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(18)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(41)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') {
			state = uint16(27)
			goto next_state
		}
		if int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(28)
			goto next_state
		}
		return result
	case int32(1):
		if lookahead == int32('"') {
			state = uint16(29)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(8)
			goto next_state
		}
		if lookahead > int32(0x1f) && lookahead != int32(0x7f) && lookahead != int32(0x17f) && lookahead != int32(0x212a) {
			state = uint16(1)
			goto next_state
		}
		return result
	case int32(2):
		if lookahead == int32('*') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(44)
			goto next_state
		}
		return result
	case int32(3):
		if lookahead == int32('.') {
			state = uint16(12)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(3)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead == int32('.') {
			state = uint16(11)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(4)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead == int32('/') {
			state = uint16(57)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead == int32('/') {
			state = uint16(2)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(6)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(42)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead == int32('>') {
			state = uint16(36)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead == int32('u') {
			state = uint16(9)
			goto next_state
		}
		if lookahead == int32('"') || lookahead == int32('\'') || lookahead == int32('\\') || lookahead == int32('n') || lookahead == int32('r') || lookahead == int32('t') {
			state = uint16(1)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(13)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead == int32('{') {
			state = uint16(14)
			goto next_state
		}
		return result
	case int32(10):
		if lookahead == int32('}') {
			state = uint16(1)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(10)
			goto next_state
		}
		return result
	case int32(11):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(30)
			goto next_state
		}
		return result
	case int32(12):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(4)
			goto next_state
		}
		return result
	case int32(13):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(1)
			goto next_state
		}
		return result
	case int32(14):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(10)
			goto next_state
		}
		return result
	case int32(15):
		if int32('A') <= lookahead && lookahead <= int32('Z') {
			state = uint16(27)
			goto next_state
		}
		if int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(28)
			goto next_state
		}
		return result
	case int32(16):
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(32)
			goto next_state
		}
		return result
	case int32(17):
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(31)
			goto next_state
		}
		return result
	case int32(18):
		if eof != 0 {
			state = uint16(20)
			goto next_state
		}
		i1 = uint32(0)
		for {
			if !(uint64(i1) < libc.Uint64FromInt64(76)/libc.Uint64FromInt64(2)) {
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
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(18)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(41)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') {
			state = uint16(27)
			goto next_state
		}
		if int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(28)
			goto next_state
		}
		return result
	case int32(19):
		if eof != 0 {
			state = uint16(20)
			goto next_state
		}
		i2 = uint32(0)
		for {
			if !(uint64(i2) < libc.Uint64FromInt64(72)/libc.Uint64FromInt64(2)) {
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
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(19)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(3)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') {
			state = uint16(27)
			goto next_state
		}
		if int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(28)
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
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COLON)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(24):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SLASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('*') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(44)
			goto next_state
		}
		return result
	case int32(25):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(26):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SEMI)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(27):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_id)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') {
			state = uint16(15)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') {
			state = uint16(27)
			goto next_state
		}
		return result
	case int32(28):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_id)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') {
			state = uint16(15)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(28)
			goto next_state
		}
		return result
	case int32(29):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(30):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__valid_semver)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('+') {
			state = uint16(16)
			goto next_state
		}
		if lookahead == int32('-') {
			state = uint16(17)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(30)
			goto next_state
		}
		return result
	case int32(31):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__valid_semver)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('+') {
			state = uint16(16)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(17)
			goto next_state
		}
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(31)
			goto next_state
		}
		return result
	case int32(32):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__valid_semver)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(16)
			goto next_state
		}
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(32)
			goto next_state
		}
		return result
	case int32(33):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COMMA)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(34):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(35):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(36):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DASH_GT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(37):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(38):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(39):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(40):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_GT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(41):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_uint)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(12)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(41)
			goto next_state
		}
		return result
	case int32(42):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_uint)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(42)
			goto next_state
		}
		return result
	case int32(43):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym__)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(44):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SLASH_SLASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(45):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SLASH_SLASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32(0x17f) && lookahead != int32(0x212a) {
			state = uint16(49)
			goto next_state
		}
		return result
	case int32(46):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SLASH_SLASH2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(47):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_line_comment_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('*') {
			state = uint16(56)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(45)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32(0x17f) && lookahead != int32(0x212a) {
			state = uint16(49)
			goto next_state
		}
		return result
	case int32(48):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_line_comment_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('/') {
			state = uint16(47)
			goto next_state
		}
		if lookahead == int32('\t') || int32(0x0b) <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(48)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(0x17f) && lookahead != int32(0x212a) {
			state = uint16(49)
			goto next_state
		}
		return result
	case int32(49):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_line_comment_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32(0x17f) && lookahead != int32(0x212a) {
			state = uint16(49)
			goto next_state
		}
		return result
	case int32(50):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SLASH2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('/') {
			state = uint16(46)
			goto next_state
		}
		return result
	case int32(51):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_line_comment_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('*') {
			state = uint16(54)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(54)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32(0x17f) && lookahead != int32(0x212a) {
			state = uint16(54)
			goto next_state
		}
		return result
	case int32(52):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_line_comment_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('/') {
			state = uint16(50)
			goto next_state
		}
		if lookahead == int32('\t') || int32(0x0b) <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(53)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(0x17f) && lookahead != int32(0x212a) {
			state = uint16(54)
			goto next_state
		}
		return result
	case int32(53):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_line_comment_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('/') {
			state = uint16(51)
			goto next_state
		}
		if lookahead == int32('\t') || int32(0x0b) <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(53)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(0x17f) && lookahead != int32(0x212a) {
			state = uint16(54)
			goto next_state
		}
		return result
	case int32(54):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_line_comment_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32(0x17f) && lookahead != int32(0x212a) {
			state = uint16(54)
			goto next_state
		}
		return result
	case int32(55):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SLASH_STAR)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(56):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SLASH_STAR)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32(0x17f) && lookahead != int32(0x212a) {
			state = uint16(49)
			goto next_state
		}
		return result
	case int32(57):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STAR_SLASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var map_token = [38]uint16_t{
	0:  uint16('"'),
	1:  uint16(1),
	2:  uint16('%'),
	3:  uint16(15),
	4:  uint16('('),
	5:  uint16(34),
	6:  uint16(')'),
	7:  uint16(35),
	8:  uint16('*'),
	9:  uint16(5),
	10: uint16(','),
	11: uint16(33),
	12: uint16('-'),
	13: uint16(7),
	14: uint16('.'),
	15: uint16(37),
	16: uint16('/'),
	17: uint16(50),
	18: uint16('0'),
	19: uint16(3),
	20: uint16(':'),
	21: uint16(23),
	22: uint16(';'),
	23: uint16(26),
	24: uint16('<'),
	25: uint16(39),
	26: uint16('='),
	27: uint16(38),
	28: uint16('>'),
	29: uint16(40),
	30: uint16('@'),
	31: uint16(25),
	32: uint16('_'),
	33: uint16(43),
	34: uint16('{'),
	35: uint16(21),
	36: uint16('}'),
	37: uint16(22),
}

var map_token1 = [38]uint16_t{
	0:  uint16('"'),
	1:  uint16(1),
	2:  uint16('%'),
	3:  uint16(15),
	4:  uint16('('),
	5:  uint16(34),
	6:  uint16(')'),
	7:  uint16(35),
	8:  uint16('*'),
	9:  uint16(5),
	10: uint16(','),
	11: uint16(33),
	12: uint16('-'),
	13: uint16(7),
	14: uint16('.'),
	15: uint16(37),
	16: uint16('/'),
	17: uint16(24),
	18: uint16('0'),
	19: uint16(3),
	20: uint16(':'),
	21: uint16(23),
	22: uint16(';'),
	23: uint16(26),
	24: uint16('<'),
	25: uint16(39),
	26: uint16('='),
	27: uint16(38),
	28: uint16('>'),
	29: uint16(40),
	30: uint16('@'),
	31: uint16(25),
	32: uint16('_'),
	33: uint16(43),
	34: uint16('{'),
	35: uint16(21),
	36: uint16('}'),
	37: uint16(22),
}

var map_token2 = [36]uint16_t{
	0:  uint16('"'),
	1:  uint16(1),
	2:  uint16('%'),
	3:  uint16(15),
	4:  uint16('('),
	5:  uint16(34),
	6:  uint16(')'),
	7:  uint16(35),
	8:  uint16('*'),
	9:  uint16(5),
	10: uint16(','),
	11: uint16(33),
	12: uint16('-'),
	13: uint16(7),
	14: uint16('.'),
	15: uint16(37),
	16: uint16('/'),
	17: uint16(24),
	18: uint16(':'),
	19: uint16(23),
	20: uint16(';'),
	21: uint16(26),
	22: uint16('<'),
	23: uint16(39),
	24: uint16('='),
	25: uint16(38),
	26: uint16('>'),
	27: uint16(40),
	28: uint16('@'),
	29: uint16(25),
	30: uint16('_'),
	31: uint16(43),
	32: uint16('{'),
	33: uint16(21),
	34: uint16('}'),
	35: uint16(22),
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
			if int32(map_token3[i]) == lookahead {
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
		if lookahead == int32('s') {
			state = uint16(18)
			goto next_state
		}
		return result
	case int32(2):
		if lookahead == int32('o') {
			state = uint16(19)
			goto next_state
		}
		return result
	case int32(3):
		if lookahead == int32('h') {
			state = uint16(20)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(21)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead == int32('e') {
			state = uint16(22)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead == int32('n') {
			state = uint16(23)
			goto next_state
		}
		if lookahead == int32('x') {
			state = uint16(24)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead == int32('3') {
			state = uint16(25)
			goto next_state
		}
		if lookahead == int32('6') {
			state = uint16(26)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(27)
			goto next_state
		}
		if lookahead == int32('l') {
			state = uint16(28)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(29)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead == int32('m') {
			state = uint16(30)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(31)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead == int32('i') {
			state = uint16(32)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead == int32('a') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(10):
		if lookahead == int32('p') {
			state = uint16(34)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead == int32('a') {
			state = uint16(35)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead == int32('e') {
			state = uint16(36)
			goto next_state
		}
		return result
	case int32(13):
		if lookahead == int32('1') {
			state = uint16(37)
			goto next_state
		}
		if lookahead == int32('3') {
			state = uint16(38)
			goto next_state
		}
		if lookahead == int32('6') {
			state = uint16(39)
			goto next_state
		}
		if lookahead == int32('8') {
			state = uint16(40)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(41)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(42)
			goto next_state
		}
		return result
	case int32(14):
		if lookahead == int32('u') {
			state = uint16(43)
			goto next_state
		}
		if lookahead == int32('y') {
			state = uint16(44)
			goto next_state
		}
		return result
	case int32(15):
		if lookahead == int32('1') {
			state = uint16(45)
			goto next_state
		}
		if lookahead == int32('3') {
			state = uint16(46)
			goto next_state
		}
		if lookahead == int32('6') {
			state = uint16(47)
			goto next_state
		}
		if lookahead == int32('8') {
			state = uint16(48)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(49)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(50)
			goto next_state
		}
		return result
	case int32(16):
		if lookahead == int32('a') {
			state = uint16(51)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(17):
		if lookahead == int32('i') {
			state = uint16(53)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(54)
			goto next_state
		}
		return result
	case int32(18):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_as)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('y') {
			state = uint16(55)
			goto next_state
		}
		return result
	case int32(19):
		if lookahead == int32('o') {
			state = uint16(56)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(57)
			goto next_state
		}
		return result
	case int32(20):
		if lookahead == int32('a') {
			state = uint16(58)
			goto next_state
		}
		return result
	case int32(21):
		if lookahead == int32('n') {
			state = uint16(59)
			goto next_state
		}
		return result
	case int32(22):
		if lookahead == int32('p') {
			state = uint16(60)
			goto next_state
		}
		return result
	case int32(23):
		if lookahead == int32('u') {
			state = uint16(61)
			goto next_state
		}
		return result
	case int32(24):
		if lookahead == int32('p') {
			state = uint16(62)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(63)
			goto next_state
		}
		return result
	case int32(25):
		if lookahead == int32('2') {
			state = uint16(64)
			goto next_state
		}
		return result
	case int32(26):
		if lookahead == int32('4') {
			state = uint16(65)
			goto next_state
		}
		return result
	case int32(27):
		if lookahead == int32('a') {
			state = uint16(66)
			goto next_state
		}
		return result
	case int32(28):
		if lookahead == int32('a') {
			state = uint16(67)
			goto next_state
		}
		return result
	case int32(29):
		if lookahead == int32('n') {
			state = uint16(68)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(30):
		if lookahead == int32('p') {
			state = uint16(70)
			goto next_state
		}
		return result
	case int32(31):
		if lookahead == int32('c') {
			state = uint16(71)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(72)
			goto next_state
		}
		return result
	case int32(32):
		if lookahead == int32('s') {
			state = uint16(73)
			goto next_state
		}
		return result
	case int32(33):
		if lookahead == int32('p') {
			state = uint16(74)
			goto next_state
		}
		return result
	case int32(34):
		if lookahead == int32('t') {
			state = uint16(75)
			goto next_state
		}
		return result
	case int32(35):
		if lookahead == int32('c') {
			state = uint16(76)
			goto next_state
		}
		return result
	case int32(36):
		if lookahead == int32('c') {
			state = uint16(77)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(78)
			goto next_state
		}
		return result
	case int32(37):
		if lookahead == int32('6') {
			state = uint16(79)
			goto next_state
		}
		return result
	case int32(38):
		if lookahead == int32('2') {
			state = uint16(80)
			goto next_state
		}
		return result
	case int32(39):
		if lookahead == int32('4') {
			state = uint16(81)
			goto next_state
		}
		return result
	case int32(40):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_s8)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(41):
		if lookahead == int32('n') {
			state = uint16(82)
			goto next_state
		}
		return result
	case int32(42):
		if lookahead == int32('a') {
			state = uint16(83)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(84)
			goto next_state
		}
		return result
	case int32(43):
		if lookahead == int32('p') {
			state = uint16(85)
			goto next_state
		}
		return result
	case int32(44):
		if lookahead == int32('p') {
			state = uint16(86)
			goto next_state
		}
		return result
	case int32(45):
		if lookahead == int32('6') {
			state = uint16(87)
			goto next_state
		}
		return result
	case int32(46):
		if lookahead == int32('2') {
			state = uint16(88)
			goto next_state
		}
		return result
	case int32(47):
		if lookahead == int32('4') {
			state = uint16(89)
			goto next_state
		}
		return result
	case int32(48):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_u8)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(49):
		if lookahead == int32('s') {
			state = uint16(90)
			goto next_state
		}
		return result
	case int32(50):
		if lookahead == int32('e') {
			state = uint16(91)
			goto next_state
		}
		return result
	case int32(51):
		if lookahead == int32('r') {
			state = uint16(92)
			goto next_state
		}
		return result
	case int32(52):
		if lookahead == int32('r') {
			state = uint16(93)
			goto next_state
		}
		return result
	case int32(53):
		if lookahead == int32('t') {
			state = uint16(94)
			goto next_state
		}
		return result
	case int32(54):
		if lookahead == int32('r') {
			state = uint16(95)
			goto next_state
		}
		return result
	case int32(55):
		if lookahead == int32('n') {
			state = uint16(96)
			goto next_state
		}
		return result
	case int32(56):
		if lookahead == int32('l') {
			state = uint16(97)
			goto next_state
		}
		return result
	case int32(57):
		if lookahead == int32('r') {
			state = uint16(98)
			goto next_state
		}
		return result
	case int32(58):
		if lookahead == int32('r') {
			state = uint16(99)
			goto next_state
		}
		return result
	case int32(59):
		if lookahead == int32('s') {
			state = uint16(100)
			goto next_state
		}
		return result
	case int32(60):
		if lookahead == int32('r') {
			state = uint16(101)
			goto next_state
		}
		return result
	case int32(61):
		if lookahead == int32('m') {
			state = uint16(102)
			goto next_state
		}
		return result
	case int32(62):
		if lookahead == int32('o') {
			state = uint16(103)
			goto next_state
		}
		return result
	case int32(63):
		if lookahead == int32('e') {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(64):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_f32)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(65):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_f64)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(66):
		if lookahead == int32('t') {
			state = uint16(105)
			goto next_state
		}
		return result
	case int32(67):
		if lookahead == int32('g') {
			state = uint16(106)
			goto next_state
		}
		return result
	case int32(68):
		if lookahead == int32('c') {
			state = uint16(107)
			goto next_state
		}
		return result
	case int32(69):
		if lookahead == int32('u') {
			state = uint16(108)
			goto next_state
		}
		return result
	case int32(70):
		if lookahead == int32('o') {
			state = uint16(109)
			goto next_state
		}
		return result
	case int32(71):
		if lookahead == int32('l') {
			state = uint16(110)
			goto next_state
		}
		return result
	case int32(72):
		if lookahead == int32('e') {
			state = uint16(111)
			goto next_state
		}
		return result
	case int32(73):
		if lookahead == int32('t') {
			state = uint16(112)
			goto next_state
		}
		return result
	case int32(74):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_map)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(75):
		if lookahead == int32('i') {
			state = uint16(113)
			goto next_state
		}
		return result
	case int32(76):
		if lookahead == int32('k') {
			state = uint16(114)
			goto next_state
		}
		return result
	case int32(77):
		if lookahead == int32('o') {
			state = uint16(115)
			goto next_state
		}
		return result
	case int32(78):
		if lookahead == int32('o') {
			state = uint16(116)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(117)
			goto next_state
		}
		return result
	case int32(79):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_s16)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(80):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_s32)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(81):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_s64)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(82):
		if lookahead == int32('c') {
			state = uint16(118)
			goto next_state
		}
		return result
	case int32(83):
		if lookahead == int32('t') {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(84):
		if lookahead == int32('e') {
			state = uint16(120)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(85):
		if lookahead == int32('l') {
			state = uint16(122)
			goto next_state
		}
		return result
	case int32(86):
		if lookahead == int32('e') {
			state = uint16(123)
			goto next_state
		}
		return result
	case int32(87):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_u16)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(88):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_u32)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(89):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_u64)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(90):
		if lookahead == int32('t') {
			state = uint16(124)
			goto next_state
		}
		return result
	case int32(91):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_use)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(92):
		if lookahead == int32('i') {
			state = uint16(125)
			goto next_state
		}
		return result
	case int32(93):
		if lookahead == int32('s') {
			state = uint16(126)
			goto next_state
		}
		return result
	case int32(94):
		if lookahead == int32('h') {
			state = uint16(127)
			goto next_state
		}
		return result
	case int32(95):
		if lookahead == int32('l') {
			state = uint16(128)
			goto next_state
		}
		return result
	case int32(96):
		if lookahead == int32('c') {
			state = uint16(129)
			goto next_state
		}
		return result
	case int32(97):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_bool)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(98):
		if lookahead == int32('o') {
			state = uint16(130)
			goto next_state
		}
		return result
	case int32(99):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_char)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(100):
		if lookahead == int32('t') {
			state = uint16(131)
			goto next_state
		}
		return result
	case int32(101):
		if lookahead == int32('e') {
			state = uint16(132)
			goto next_state
		}
		return result
	case int32(102):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_enum)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(103):
		if lookahead == int32('r') {
			state = uint16(133)
			goto next_state
		}
		return result
	case int32(104):
		if lookahead == int32('r') {
			state = uint16(134)
			goto next_state
		}
		return result
	case int32(105):
		if lookahead == int32('u') {
			state = uint16(135)
			goto next_state
		}
		return result
	case int32(106):
		if lookahead == int32('s') {
			state = uint16(136)
			goto next_state
		}
		return result
	case int32(107):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_func)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(108):
		if lookahead == int32('r') {
			state = uint16(137)
			goto next_state
		}
		return result
	case int32(109):
		if lookahead == int32('r') {
			state = uint16(138)
			goto next_state
		}
		return result
	case int32(110):
		if lookahead == int32('u') {
			state = uint16(139)
			goto next_state
		}
		return result
	case int32(111):
		if lookahead == int32('r') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(112):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_list)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(113):
		if lookahead == int32('o') {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(114):
		if lookahead == int32('a') {
			state = uint16(142)
			goto next_state
		}
		return result
	case int32(115):
		if lookahead == int32('r') {
			state = uint16(143)
			goto next_state
		}
		return result
	case int32(116):
		if lookahead == int32('u') {
			state = uint16(144)
			goto next_state
		}
		return result
	case int32(117):
		if lookahead == int32('l') {
			state = uint16(145)
			goto next_state
		}
		return result
	case int32(118):
		if lookahead == int32('e') {
			state = uint16(146)
			goto next_state
		}
		return result
	case int32(119):
		if lookahead == int32('i') {
			state = uint16(147)
			goto next_state
		}
		return result
	case int32(120):
		if lookahead == int32('a') {
			state = uint16(148)
			goto next_state
		}
		return result
	case int32(121):
		if lookahead == int32('n') {
			state = uint16(149)
			goto next_state
		}
		return result
	case int32(122):
		if lookahead == int32('e') {
			state = uint16(150)
			goto next_state
		}
		return result
	case int32(123):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_type)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(124):
		if lookahead == int32('a') {
			state = uint16(151)
			goto next_state
		}
		return result
	case int32(125):
		if lookahead == int32('a') {
			state = uint16(152)
			goto next_state
		}
		return result
	case int32(126):
		if lookahead == int32('i') {
			state = uint16(153)
			goto next_state
		}
		return result
	case int32(127):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_with)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(128):
		if lookahead == int32('d') {
			state = uint16(154)
			goto next_state
		}
		return result
	case int32(129):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_async)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(130):
		if lookahead == int32('w') {
			state = uint16(155)
			goto next_state
		}
		return result
	case int32(131):
		if lookahead == int32('r') {
			state = uint16(156)
			goto next_state
		}
		return result
	case int32(132):
		if lookahead == int32('c') {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(133):
		if lookahead == int32('t') {
			state = uint16(158)
			goto next_state
		}
		return result
	case int32(134):
		if lookahead == int32('n') {
			state = uint16(159)
			goto next_state
		}
		return result
	case int32(135):
		if lookahead == int32('r') {
			state = uint16(160)
			goto next_state
		}
		return result
	case int32(136):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_flags)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(137):
		if lookahead == int32('e') {
			state = uint16(161)
			goto next_state
		}
		return result
	case int32(138):
		if lookahead == int32('t') {
			state = uint16(162)
			goto next_state
		}
		return result
	case int32(139):
		if lookahead == int32('d') {
			state = uint16(163)
			goto next_state
		}
		return result
	case int32(140):
		if lookahead == int32('f') {
			state = uint16(164)
			goto next_state
		}
		return result
	case int32(141):
		if lookahead == int32('n') {
			state = uint16(165)
			goto next_state
		}
		return result
	case int32(142):
		if lookahead == int32('g') {
			state = uint16(166)
			goto next_state
		}
		return result
	case int32(143):
		if lookahead == int32('d') {
			state = uint16(167)
			goto next_state
		}
		return result
	case int32(144):
		if lookahead == int32('r') {
			state = uint16(168)
			goto next_state
		}
		return result
	case int32(145):
		if lookahead == int32('t') {
			state = uint16(169)
			goto next_state
		}
		return result
	case int32(146):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_since)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(147):
		if lookahead == int32('c') {
			state = uint16(170)
			goto next_state
		}
		return result
	case int32(148):
		if lookahead == int32('m') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(149):
		if lookahead == int32('g') {
			state = uint16(172)
			goto next_state
		}
		return result
	case int32(150):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_tuple)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(151):
		if lookahead == int32('b') {
			state = uint16(173)
			goto next_state
		}
		return result
	case int32(152):
		if lookahead == int32('n') {
			state = uint16(174)
			goto next_state
		}
		return result
	case int32(153):
		if lookahead == int32('o') {
			state = uint16(175)
			goto next_state
		}
		return result
	case int32(154):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_world)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(155):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_borrow)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(156):
		if lookahead == int32('u') {
			state = uint16(176)
			goto next_state
		}
		return result
	case int32(157):
		if lookahead == int32('a') {
			state = uint16(177)
			goto next_state
		}
		return result
	case int32(158):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_export)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(159):
		if lookahead == int32('a') {
			state = uint16(178)
			goto next_state
		}
		return result
	case int32(160):
		if lookahead == int32('e') {
			state = uint16(179)
			goto next_state
		}
		return result
	case int32(161):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_future)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(162):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_import)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(163):
		if lookahead == int32('e') {
			state = uint16(180)
			goto next_state
		}
		return result
	case int32(164):
		if lookahead == int32('a') {
			state = uint16(181)
			goto next_state
		}
		return result
	case int32(165):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_option)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(166):
		if lookahead == int32('e') {
			state = uint16(182)
			goto next_state
		}
		return result
	case int32(167):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_record)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(168):
		if lookahead == int32('c') {
			state = uint16(183)
			goto next_state
		}
		return result
	case int32(169):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_result)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(170):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_static)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(171):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_stream)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(172):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(173):
		if lookahead == int32('l') {
			state = uint16(184)
			goto next_state
		}
		return result
	case int32(174):
		if lookahead == int32('t') {
			state = uint16(185)
			goto next_state
		}
		return result
	case int32(175):
		if lookahead == int32('n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(176):
		if lookahead == int32('c') {
			state = uint16(187)
			goto next_state
		}
		return result
	case int32(177):
		if lookahead == int32('t') {
			state = uint16(188)
			goto next_state
		}
		return result
	case int32(178):
		if lookahead == int32('l') {
			state = uint16(189)
			goto next_state
		}
		return result
	case int32(179):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_feature)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(180):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_include)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(181):
		if lookahead == int32('c') {
			state = uint16(190)
			goto next_state
		}
		return result
	case int32(182):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_package)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(183):
		if lookahead == int32('e') {
			state = uint16(191)
			goto next_state
		}
		return result
	case int32(184):
		if lookahead == int32('e') {
			state = uint16(192)
			goto next_state
		}
		return result
	case int32(185):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_variant)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(186):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_version)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(187):
		if lookahead == int32('t') {
			state = uint16(193)
			goto next_state
		}
		return result
	case int32(188):
		if lookahead == int32('e') {
			state = uint16(194)
			goto next_state
		}
		return result
	case int32(189):
		if lookahead == int32('-') {
			state = uint16(195)
			goto next_state
		}
		return result
	case int32(190):
		if lookahead == int32('e') {
			state = uint16(196)
			goto next_state
		}
		return result
	case int32(191):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_resource)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(192):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_unstable)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(193):
		if lookahead == int32('o') {
			state = uint16(197)
			goto next_state
		}
		return result
	case int32(194):
		if lookahead == int32('d') {
			state = uint16(198)
			goto next_state
		}
		return result
	case int32(195):
		if lookahead == int32('i') {
			state = uint16(199)
			goto next_state
		}
		return result
	case int32(196):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_interface)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(197):
		if lookahead == int32('r') {
			state = uint16(200)
			goto next_state
		}
		return result
	case int32(198):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_deprecated)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(199):
		if lookahead == int32('d') {
			state = uint16(201)
			goto next_state
		}
		return result
	case int32(200):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_constructor)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(201):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_external_DASHid)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var map_token3 = [34]uint16_t{
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

var ts_lex_modes = [349]TSLexerMode{
	0: {
		Fexternal_lex_state: uint16(1),
	},
	1: {
		Flex_state: uint16(19),
	},
	2: {
		Flex_state: uint16(19),
	},
	3: {
		Flex_state: uint16(19),
	},
	4: {
		Flex_state: uint16(19),
	},
	5: {
		Flex_state: uint16(19),
	},
	6: {
		Flex_state: uint16(19),
	},
	7: {
		Flex_state: uint16(19),
	},
	8: {
		Flex_state: uint16(19),
	},
	9: {
		Flex_state: uint16(19),
	},
	10: {
		Flex_state: uint16(19),
	},
	11: {
		Flex_state: uint16(19),
	},
	12: {
		Flex_state: uint16(19),
	},
	13: {
		Flex_state: uint16(19),
	},
	14: {
		Flex_state: uint16(19),
	},
	15: {
		Flex_state: uint16(19),
	},
	16: {
		Flex_state: uint16(19),
	},
	17: {
		Flex_state: uint16(19),
	},
	18: {
		Flex_state: uint16(19),
	},
	19: {
		Flex_state: uint16(19),
	},
	20: {
		Flex_state: uint16(19),
	},
	21: {
		Flex_state: uint16(19),
	},
	22: {
		Flex_state: uint16(19),
	},
	23: {
		Flex_state: uint16(19),
	},
	24: {
		Flex_state: uint16(19),
	},
	25: {
		Flex_state: uint16(19),
	},
	26: {
		Flex_state: uint16(19),
	},
	27: {
		Flex_state: uint16(19),
	},
	28: {
		Flex_state: uint16(19),
	},
	29: {
		Flex_state: uint16(19),
	},
	30: {
		Flex_state: uint16(19),
	},
	31: {
		Flex_state: uint16(19),
	},
	32: {
		Flex_state: uint16(19),
	},
	33: {
		Flex_state: uint16(19),
	},
	34: {
		Flex_state: uint16(19),
	},
	35: {
		Flex_state: uint16(19),
	},
	36: {
		Flex_state: uint16(19),
	},
	37: {
		Flex_state: uint16(19),
	},
	38: {
		Flex_state: uint16(19),
	},
	39: {
		Flex_state: uint16(19),
	},
	40: {
		Flex_state: uint16(19),
	},
	41: {
		Flex_state: uint16(19),
	},
	42: {
		Flex_state: uint16(19),
	},
	43: {
		Flex_state: uint16(19),
	},
	44: {
		Flex_state: uint16(19),
	},
	45: {
		Flex_state: uint16(19),
	},
	46: {
		Flex_state: uint16(19),
	},
	47: {
		Flex_state: uint16(19),
	},
	48: {
		Flex_state: uint16(19),
	},
	49: {
		Flex_state: uint16(19),
	},
	50: {
		Flex_state: uint16(19),
	},
	51: {
		Flex_state: uint16(19),
	},
	52: {
		Flex_state: uint16(19),
	},
	53: {
		Flex_state: uint16(19),
	},
	54: {
		Flex_state: uint16(19),
	},
	55: {
		Flex_state: uint16(19),
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
		Flex_state: uint16(19),
	},
	61: {
		Flex_state: uint16(19),
	},
	62: {
		Flex_state: uint16(19),
	},
	63: {
		Flex_state: uint16(19),
	},
	64: {
		Flex_state: uint16(19),
	},
	65: {
		Flex_state: uint16(19),
	},
	66: {
		Flex_state: uint16(19),
	},
	67: {
		Flex_state: uint16(19),
	},
	68: {
		Flex_state: uint16(19),
	},
	69: {
		Flex_state: uint16(19),
	},
	70: {
		Flex_state: uint16(19),
	},
	71: {
		Flex_state: uint16(19),
	},
	72: {
		Flex_state: uint16(19),
	},
	73: {
		Flex_state: uint16(19),
	},
	74: {
		Flex_state: uint16(19),
	},
	75: {
		Flex_state: uint16(19),
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
		Flex_state: uint16(19),
	},
	89: {
		Flex_state: uint16(19),
	},
	90: {
		Flex_state: uint16(19),
	},
	91: {
		Flex_state: uint16(19),
	},
	92: {
		Flex_state: uint16(19),
	},
	93: {
		Flex_state: uint16(19),
	},
	94: {
		Flex_state: uint16(19),
	},
	95: {
		Flex_state: uint16(19),
	},
	96: {
		Flex_state: uint16(19),
	},
	97: {
		Flex_state: uint16(19),
	},
	98: {
		Flex_state: uint16(19),
	},
	99: {
		Flex_state: uint16(19),
	},
	100: {
		Flex_state: uint16(19),
	},
	101: {
		Flex_state: uint16(19),
	},
	102: {
		Flex_state: uint16(19),
	},
	103: {
		Flex_state: uint16(19),
	},
	104: {
		Flex_state: uint16(19),
	},
	105: {
		Flex_state: uint16(19),
	},
	106: {
		Flex_state: uint16(19),
	},
	107: {
		Flex_state: uint16(19),
	},
	108: {
		Flex_state: uint16(19),
	},
	109: {
		Flex_state: uint16(19),
	},
	110: {
		Flex_state: uint16(19),
	},
	111: {
		Flex_state: uint16(19),
	},
	112: {
		Flex_state: uint16(19),
	},
	113: {
		Flex_state: uint16(19),
	},
	114: {
		Flex_state: uint16(19),
	},
	115: {
		Flex_state: uint16(19),
	},
	116: {
		Flex_state: uint16(19),
	},
	117: {
		Flex_state: uint16(19),
	},
	118: {
		Flex_state: uint16(19),
	},
	119: {
		Flex_state: uint16(19),
	},
	120: {
		Flex_state: uint16(19),
	},
	121: {
		Flex_state: uint16(19),
	},
	122: {
		Flex_state: uint16(19),
	},
	123: {
		Flex_state: uint16(19),
	},
	124: {
		Flex_state: uint16(19),
	},
	125: {
		Flex_state: uint16(19),
	},
	126: {
		Flex_state: uint16(19),
	},
	127: {
		Flex_state: uint16(19),
	},
	128: {
		Flex_state: uint16(19),
	},
	129: {
		Flex_state: uint16(19),
	},
	130: {
		Flex_state: uint16(19),
	},
	131: {
		Flex_state: uint16(19),
	},
	132: {
		Flex_state: uint16(19),
	},
	133: {
		Flex_state: uint16(19),
	},
	134: {
		Flex_state: uint16(19),
	},
	135: {
		Flex_state: uint16(19),
	},
	136: {
		Flex_state: uint16(19),
	},
	137: {
		Flex_state: uint16(19),
	},
	138: {
		Flex_state: uint16(19),
	},
	139: {
		Flex_state: uint16(19),
	},
	140: {
		Flex_state: uint16(19),
	},
	141: {
		Flex_state: uint16(19),
	},
	142: {
		Flex_state: uint16(19),
	},
	143: {
		Flex_state: uint16(19),
	},
	144: {
		Flex_state: uint16(19),
	},
	145: {
		Flex_state: uint16(19),
	},
	146: {
		Flex_state: uint16(19),
	},
	147: {
		Flex_state: uint16(19),
	},
	148: {
		Flex_state: uint16(19),
	},
	149: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	150: {
		Flex_state: uint16(19),
	},
	151: {
		Flex_state: uint16(19),
	},
	152: {
		Flex_state: uint16(19),
	},
	153: {
		Flex_state: uint16(19),
	},
	154: {
		Flex_state: uint16(19),
	},
	155: {
		Flex_state: uint16(19),
	},
	156: {
		Flex_state: uint16(19),
	},
	157: {
		Flex_state: uint16(19),
	},
	158: {
		Flex_state: uint16(52),
	},
	159: {
		Flex_state: uint16(19),
	},
	160: {
		Flex_state: uint16(19),
	},
	161: {
		Flex_state: uint16(19),
	},
	162: {
		Flex_state: uint16(19),
	},
	163: {
		Flex_state: uint16(19),
	},
	164: {
		Flex_state: uint16(19),
	},
	165: {
		Flex_state: uint16(19),
	},
	166: {
		Flex_state: uint16(19),
	},
	167: {
		Flex_state: uint16(19),
	},
	168: {
		Flex_state: uint16(19),
	},
	169: {
		Flex_state: uint16(19),
	},
	170: {
		Flex_state: uint16(19),
	},
	171: {
		Flex_state: uint16(19),
	},
	172: {
		Flex_state: uint16(19),
	},
	173: {
		Flex_state: uint16(19),
	},
	174: {
		Flex_state: uint16(19),
	},
	175: {
		Flex_state: uint16(19),
	},
	176: {
		Flex_state: uint16(19),
	},
	177: {
		Flex_state: uint16(19),
	},
	178: {
		Flex_state: uint16(19),
	},
	179: {
		Flex_state: uint16(19),
	},
	180: {
		Flex_state: uint16(19),
	},
	181: {
		Flex_state: uint16(19),
	},
	182: {
		Flex_state: uint16(19),
	},
	183: {
		Flex_state: uint16(19),
	},
	184: {
		Flex_state: uint16(19),
	},
	185: {
		Flex_state: uint16(19),
	},
	186: {
		Flex_state: uint16(19),
	},
	187: {
		Flex_state: uint16(19),
	},
	188: {
		Flex_state: uint16(19),
	},
	189: {
		Flex_state: uint16(19),
	},
	190: {
		Flex_state: uint16(19),
	},
	191: {
		Flex_state: uint16(19),
	},
	192: {
		Flex_state: uint16(19),
	},
	193: {
		Flex_state: uint16(19),
	},
	194: {
		Flex_state: uint16(19),
	},
	195: {
		Flex_state: uint16(19),
	},
	196: {
		Flex_state: uint16(19),
	},
	197: {
		Flex_state: uint16(19),
	},
	198: {
		Flex_state: uint16(19),
	},
	199: {
		Flex_state: uint16(19),
	},
	200: {
		Flex_state: uint16(19),
	},
	201: {
		Flex_state: uint16(19),
	},
	202: {
		Flex_state: uint16(19),
	},
	203: {
		Flex_state: uint16(19),
	},
	204: {
		Flex_state: uint16(19),
	},
	205: {
		Flex_state: uint16(19),
	},
	206: {
		Flex_state: uint16(19),
	},
	207: {
		Flex_state: uint16(19),
	},
	208: {
		Flex_state: uint16(19),
	},
	209: {
		Flex_state: uint16(19),
	},
	210: {
		Flex_state: uint16(19),
	},
	211: {
		Flex_state: uint16(19),
	},
	212: {
		Flex_state: uint16(19),
	},
	213: {
		Flex_state: uint16(19),
	},
	214: {
		Flex_state: uint16(19),
	},
	215: {
		Flex_state: uint16(19),
	},
	216: {
		Flex_state: uint16(19),
	},
	217: {
		Flex_state: uint16(19),
	},
	218: {
		Flex_state: uint16(19),
	},
	219: {
		Flex_state: uint16(19),
	},
	220: {
		Flex_state: uint16(19),
	},
	221: {
		Flex_state: uint16(19),
	},
	222: {
		Flex_state: uint16(19),
	},
	223: {
		Flex_state: uint16(19),
	},
	224: {
		Flex_state: uint16(19),
	},
	225: {
		Flex_state: uint16(19),
	},
	226: {
		Flex_state: uint16(19),
	},
	227: {
		Flex_state: uint16(19),
	},
	228: {
		Flex_state: uint16(19),
	},
	229: {
		Flex_state: uint16(19),
	},
	230: {
		Flex_state: uint16(19),
	},
	231: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	232: {
		Flex_state: uint16(19),
	},
	233: {
		Flex_state: uint16(19),
	},
	234: {
		Flex_state: uint16(19),
	},
	235: {
		Flex_state: uint16(19),
	},
	236: {
		Flex_state: uint16(19),
	},
	237: {
		Flex_state: uint16(19),
	},
	238: {
		Flex_state: uint16(19),
	},
	239: {
		Flex_state: uint16(19),
	},
	240: {
		Flex_state: uint16(19),
	},
	241: {
		Flex_state: uint16(19),
	},
	242: {
		Flex_state: uint16(19),
	},
	243: {
		Flex_state: uint16(19),
	},
	244: {
		Flex_state: uint16(19),
	},
	245: {
		Flex_state: uint16(19),
	},
	246: {
		Flex_state: uint16(19),
	},
	247: {
		Flex_state: uint16(19),
	},
	248: {
		Flex_state: uint16(19),
	},
	249: {
		Flex_state: uint16(19),
	},
	250: {
		Flex_state: uint16(19),
	},
	251: {
		Flex_state: uint16(19),
	},
	252: {
		Flex_state: uint16(19),
	},
	253: {
		Flex_state: uint16(19),
	},
	254: {
		Flex_state: uint16(19),
	},
	255: {
		Flex_state: uint16(19),
	},
	256: {
		Flex_state: uint16(19),
	},
	257: {
		Flex_state: uint16(19),
	},
	258: {
		Flex_state: uint16(19),
	},
	259: {
		Flex_state: uint16(19),
	},
	260: {
		Flex_state: uint16(19),
	},
	261: {
		Flex_state: uint16(19),
	},
	262: {
		Flex_state: uint16(19),
	},
	263: {
		Flex_state: uint16(19),
	},
	264: {
		Flex_state: uint16(19),
	},
	265: {
		Flex_state: uint16(19),
	},
	266: {
		Flex_state: uint16(19),
	},
	267: {
		Flex_state: uint16(19),
	},
	268: {
		Flex_state: uint16(19),
	},
	269: {
		Flex_state: uint16(19),
	},
	270: {
		Flex_state: uint16(19),
	},
	271: {
		Flex_state: uint16(19),
	},
	272: {
		Flex_state: uint16(19),
	},
	273: {
		Flex_state: uint16(19),
	},
	274: {
		Flex_state: uint16(19),
	},
	275: {
		Flex_state: uint16(19),
	},
	276: {
		Flex_state: uint16(19),
	},
	277: {
		Flex_state: uint16(19),
	},
	278: {
		Flex_state: uint16(19),
	},
	279: {
		Flex_state: uint16(19),
	},
	280: {
		Flex_state: uint16(19),
	},
	281: {
		Flex_state: uint16(19),
	},
	282: {
		Flex_state: uint16(19),
	},
	283: {
		Flex_state: uint16(19),
	},
	284: {
		Flex_state: uint16(19),
	},
	285: {
		Flex_state: uint16(19),
	},
	286: {
		Flex_state: uint16(19),
	},
	287: {
		Flex_state: uint16(19),
	},
	288: {
		Flex_state: uint16(19),
	},
	289: {
		Flex_state: uint16(19),
	},
	290: {
		Flex_state: uint16(19),
	},
	291: {
		Flex_state: uint16(19),
	},
	292: {
		Flex_state: uint16(19),
	},
	293: {
		Flex_state: uint16(19),
	},
	294: {
		Flex_state: uint16(19),
	},
	295: {
		Flex_state: uint16(19),
	},
	296: {
		Flex_state: uint16(19),
	},
	297: {
		Flex_state: uint16(19),
	},
	298: {
		Flex_state: uint16(19),
	},
	299: {
		Flex_state: uint16(19),
	},
	300: {
		Flex_state: uint16(19),
	},
	301: {
		Flex_state: uint16(19),
	},
	302: {
		Flex_state: uint16(19),
	},
	303: {
		Flex_state: uint16(19),
	},
	304: {
		Flex_state: uint16(19),
	},
	305: {
		Flex_state: uint16(19),
	},
	306: {
		Flex_state: uint16(19),
	},
	307: {
		Flex_state: uint16(19),
	},
	308: {
		Flex_state: uint16(6),
	},
	309: {
		Flex_state: uint16(19),
	},
	310: {
		Flex_state: uint16(19),
	},
	311: {
		Flex_state: uint16(19),
	},
	312: {
		Flex_state: uint16(19),
	},
	313: {
		Flex_state: uint16(19),
	},
	314: {
		Flex_state: uint16(19),
	},
	315: {
		Flex_state: uint16(19),
	},
	316: {
		Flex_state: uint16(19),
	},
	317: {
		Flex_state: uint16(48),
	},
	318: {
		Flex_state: uint16(19),
	},
	319: {
		Flex_state: uint16(19),
	},
	320: {
		Flex_state: uint16(19),
	},
	321: {
		Flex_state: uint16(19),
	},
	322: {
		Flex_state: uint16(19),
	},
	323: {
		Flex_state: uint16(19),
	},
	324: {
		Flex_state: uint16(19),
	},
	325: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(4),
	},
	326: {
		Flex_state: uint16(19),
	},
	327: {
		Flex_state: uint16(19),
	},
	328: {
		Flex_state: uint16(19),
	},
	329: {
		Flex_state: uint16(19),
	},
	330: {
		Flex_state: uint16(19),
	},
	331: {
		Flex_state: uint16(19),
	},
	332: {
		Flex_state: uint16(19),
	},
	333: {
		Flex_state: uint16(19),
	},
	334: {
		Flex_state: uint16(19),
	},
	335: {
		Flex_state: uint16(19),
	},
	336: {
		Flex_state: uint16(19),
	},
	337: {
		Flex_state: uint16(19),
	},
	338: {
		Flex_state: uint16(19),
	},
	339: {
		Flex_state: uint16(19),
	},
	340: {
		Flex_state: uint16(19),
	},
	341: {
		Flex_state: uint16(19),
	},
	342: {
		Flex_state: uint16(19),
	},
	343: {
		Flex_state: uint16(-libc.Int32FromInt32(1)),
	},
	344: {
		Flex_state: uint16(-libc.Int32FromInt32(1)),
	},
	345: {
		Flex_state: uint16(-libc.Int32FromInt32(1)),
	},
	346: {
		Flex_state: uint16(-libc.Int32FromInt32(1)),
	},
	347: {
		Flex_state: uint16(-libc.Int32FromInt32(1)),
	},
	348: {
		Flex_state: uint16(-libc.Int32FromInt32(1)),
	},
}

var ts_parse_table = [2][168]uint16_t{
	0: {
		0:  uint16(1),
		1:  uint16(1),
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
		60: uint16(3),
		61: uint16(1),
		63: uint16(1),
		65: uint16(5),
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
	},
	1: {
		0:   uint16(7),
		6:   uint16(9),
		7:   uint16(11),
		9:   uint16(13),
		13:  uint16(15),
		16:  uint16(17),
		60:  uint16(19),
		65:  uint16(21),
		77:  uint16(314),
		78:  uint16(99),
		79:  uint16(100),
		80:  uint16(100),
		84:  uint16(219),
		85:  uint16(100),
		86:  uint16(97),
		88:  uint16(97),
		99:  uint16(97),
		143: uint16(1),
		144: uint16(1),
		145: uint16(82),
		146: uint16(37),
		148: uint16(38),
		150: uint16(38),
		151: uint16(38),
		153: uint16(27),
	},
}

var ts_small_parse_table = [7979]uint16_t{
	0:    uint16(17),
	1:    uint16(19),
	2:    uint16(1),
	3:    uint16(anon_sym_SLASH_SLASH),
	4:    uint16(21),
	5:    uint16(1),
	6:    uint16(anon_sym_SLASH_STAR),
	7:    uint16(23),
	8:    uint16(1),
	9:    uint16(sym_id),
	10:   uint16(29),
	11:   uint16(1),
	12:   uint16(anon_sym_tuple),
	13:   uint16(31),
	14:   uint16(1),
	15:   uint16(anon_sym_list),
	16:   uint16(33),
	17:   uint16(1),
	18:   uint16(anon_sym_option),
	19:   uint16(35),
	20:   uint16(1),
	21:   uint16(anon_sym_map),
	22:   uint16(37),
	23:   uint16(1),
	24:   uint16(anon_sym_result),
	25:   uint16(39),
	26:   uint16(1),
	27:   uint16(anon_sym__),
	28:   uint16(41),
	29:   uint16(1),
	30:   uint16(anon_sym_borrow),
	31:   uint16(43),
	32:   uint16(1),
	33:   uint16(anon_sym_future),
	34:   uint16(45),
	35:   uint16(1),
	36:   uint16(anon_sym_stream),
	37:   uint16(201),
	38:   uint16(1),
	39:   uint16(sym_ty),
	40:   uint16(27),
	41:   uint16(2),
	42:   uint16(anon_sym_f32),
	43:   uint16(anon_sym_f64),
	44:   uint16(2),
	45:   uint16(2),
	46:   uint16(sym_line_comment),
	47:   uint16(sym_block_comment),
	48:   uint16(108),
	49:   uint16(9),
	50:   uint16(sym__primitive_ty),
	51:   uint16(sym_tuple),
	52:   uint16(sym_list),
	53:   uint16(sym_option),
	54:   uint16(sym_map),
	55:   uint16(sym_result),
	56:   uint16(sym_handle),
	57:   uint16(sym_future),
	58:   uint16(sym_stream),
	59:   uint16(25),
	60:   uint16(11),
	61:   uint16(anon_sym_u8),
	62:   uint16(anon_sym_u16),
	63:   uint16(anon_sym_u32),
	64:   uint16(anon_sym_u64),
	65:   uint16(anon_sym_s8),
	66:   uint16(anon_sym_s16),
	67:   uint16(anon_sym_s32),
	68:   uint16(anon_sym_s64),
	69:   uint16(anon_sym_char),
	70:   uint16(anon_sym_bool),
	71:   uint16(anon_sym_string),
	72:   uint16(17),
	73:   uint16(19),
	74:   uint16(1),
	75:   uint16(anon_sym_SLASH_SLASH),
	76:   uint16(21),
	77:   uint16(1),
	78:   uint16(anon_sym_SLASH_STAR),
	79:   uint16(23),
	80:   uint16(1),
	81:   uint16(sym_id),
	82:   uint16(29),
	83:   uint16(1),
	84:   uint16(anon_sym_tuple),
	85:   uint16(31),
	86:   uint16(1),
	87:   uint16(anon_sym_list),
	88:   uint16(33),
	89:   uint16(1),
	90:   uint16(anon_sym_option),
	91:   uint16(35),
	92:   uint16(1),
	93:   uint16(anon_sym_map),
	94:   uint16(37),
	95:   uint16(1),
	96:   uint16(anon_sym_result),
	97:   uint16(41),
	98:   uint16(1),
	99:   uint16(anon_sym_borrow),
	100:  uint16(43),
	101:  uint16(1),
	102:  uint16(anon_sym_future),
	103:  uint16(45),
	104:  uint16(1),
	105:  uint16(anon_sym_stream),
	106:  uint16(47),
	107:  uint16(1),
	108:  uint16(anon_sym_GT),
	109:  uint16(244),
	110:  uint16(1),
	111:  uint16(sym_ty),
	112:  uint16(27),
	113:  uint16(2),
	114:  uint16(anon_sym_f32),
	115:  uint16(anon_sym_f64),
	116:  uint16(3),
	117:  uint16(2),
	118:  uint16(sym_line_comment),
	119:  uint16(sym_block_comment),
	120:  uint16(108),
	121:  uint16(9),
	122:  uint16(sym__primitive_ty),
	123:  uint16(sym_tuple),
	124:  uint16(sym_list),
	125:  uint16(sym_option),
	126:  uint16(sym_map),
	127:  uint16(sym_result),
	128:  uint16(sym_handle),
	129:  uint16(sym_future),
	130:  uint16(sym_stream),
	131:  uint16(25),
	132:  uint16(11),
	133:  uint16(anon_sym_u8),
	134:  uint16(anon_sym_u16),
	135:  uint16(anon_sym_u32),
	136:  uint16(anon_sym_u64),
	137:  uint16(anon_sym_s8),
	138:  uint16(anon_sym_s16),
	139:  uint16(anon_sym_s32),
	140:  uint16(anon_sym_s64),
	141:  uint16(anon_sym_char),
	142:  uint16(anon_sym_bool),
	143:  uint16(anon_sym_string),
	144:  uint16(17),
	145:  uint16(19),
	146:  uint16(1),
	147:  uint16(anon_sym_SLASH_SLASH),
	148:  uint16(21),
	149:  uint16(1),
	150:  uint16(anon_sym_SLASH_STAR),
	151:  uint16(23),
	152:  uint16(1),
	153:  uint16(sym_id),
	154:  uint16(29),
	155:  uint16(1),
	156:  uint16(anon_sym_tuple),
	157:  uint16(31),
	158:  uint16(1),
	159:  uint16(anon_sym_list),
	160:  uint16(33),
	161:  uint16(1),
	162:  uint16(anon_sym_option),
	163:  uint16(35),
	164:  uint16(1),
	165:  uint16(anon_sym_map),
	166:  uint16(37),
	167:  uint16(1),
	168:  uint16(anon_sym_result),
	169:  uint16(41),
	170:  uint16(1),
	171:  uint16(anon_sym_borrow),
	172:  uint16(43),
	173:  uint16(1),
	174:  uint16(anon_sym_future),
	175:  uint16(45),
	176:  uint16(1),
	177:  uint16(anon_sym_stream),
	178:  uint16(175),
	179:  uint16(1),
	180:  uint16(sym_ty),
	181:  uint16(266),
	182:  uint16(1),
	183:  uint16(sym_tuple_list),
	184:  uint16(27),
	185:  uint16(2),
	186:  uint16(anon_sym_f32),
	187:  uint16(anon_sym_f64),
	188:  uint16(4),
	189:  uint16(2),
	190:  uint16(sym_line_comment),
	191:  uint16(sym_block_comment),
	192:  uint16(108),
	193:  uint16(9),
	194:  uint16(sym__primitive_ty),
	195:  uint16(sym_tuple),
	196:  uint16(sym_list),
	197:  uint16(sym_option),
	198:  uint16(sym_map),
	199:  uint16(sym_result),
	200:  uint16(sym_handle),
	201:  uint16(sym_future),
	202:  uint16(sym_stream),
	203:  uint16(25),
	204:  uint16(11),
	205:  uint16(anon_sym_u8),
	206:  uint16(anon_sym_u16),
	207:  uint16(anon_sym_u32),
	208:  uint16(anon_sym_u64),
	209:  uint16(anon_sym_s8),
	210:  uint16(anon_sym_s16),
	211:  uint16(anon_sym_s32),
	212:  uint16(anon_sym_s64),
	213:  uint16(anon_sym_char),
	214:  uint16(anon_sym_bool),
	215:  uint16(anon_sym_string),
	216:  uint16(17),
	217:  uint16(19),
	218:  uint16(1),
	219:  uint16(anon_sym_SLASH_SLASH),
	220:  uint16(21),
	221:  uint16(1),
	222:  uint16(anon_sym_SLASH_STAR),
	223:  uint16(23),
	224:  uint16(1),
	225:  uint16(sym_id),
	226:  uint16(29),
	227:  uint16(1),
	228:  uint16(anon_sym_tuple),
	229:  uint16(31),
	230:  uint16(1),
	231:  uint16(anon_sym_list),
	232:  uint16(33),
	233:  uint16(1),
	234:  uint16(anon_sym_option),
	235:  uint16(35),
	236:  uint16(1),
	237:  uint16(anon_sym_map),
	238:  uint16(37),
	239:  uint16(1),
	240:  uint16(anon_sym_result),
	241:  uint16(41),
	242:  uint16(1),
	243:  uint16(anon_sym_borrow),
	244:  uint16(43),
	245:  uint16(1),
	246:  uint16(anon_sym_future),
	247:  uint16(45),
	248:  uint16(1),
	249:  uint16(anon_sym_stream),
	250:  uint16(49),
	251:  uint16(1),
	252:  uint16(anon_sym_LPAREN),
	253:  uint16(296),
	254:  uint16(1),
	255:  uint16(sym_ty),
	256:  uint16(27),
	257:  uint16(2),
	258:  uint16(anon_sym_f32),
	259:  uint16(anon_sym_f64),
	260:  uint16(5),
	261:  uint16(2),
	262:  uint16(sym_line_comment),
	263:  uint16(sym_block_comment),
	264:  uint16(108),
	265:  uint16(9),
	266:  uint16(sym__primitive_ty),
	267:  uint16(sym_tuple),
	268:  uint16(sym_list),
	269:  uint16(sym_option),
	270:  uint16(sym_map),
	271:  uint16(sym_result),
	272:  uint16(sym_handle),
	273:  uint16(sym_future),
	274:  uint16(sym_stream),
	275:  uint16(25),
	276:  uint16(11),
	277:  uint16(anon_sym_u8),
	278:  uint16(anon_sym_u16),
	279:  uint16(anon_sym_u32),
	280:  uint16(anon_sym_u64),
	281:  uint16(anon_sym_s8),
	282:  uint16(anon_sym_s16),
	283:  uint16(anon_sym_s32),
	284:  uint16(anon_sym_s64),
	285:  uint16(anon_sym_char),
	286:  uint16(anon_sym_bool),
	287:  uint16(anon_sym_string),
	288:  uint16(17),
	289:  uint16(19),
	290:  uint16(1),
	291:  uint16(anon_sym_SLASH_SLASH),
	292:  uint16(21),
	293:  uint16(1),
	294:  uint16(anon_sym_SLASH_STAR),
	295:  uint16(23),
	296:  uint16(1),
	297:  uint16(sym_id),
	298:  uint16(29),
	299:  uint16(1),
	300:  uint16(anon_sym_tuple),
	301:  uint16(31),
	302:  uint16(1),
	303:  uint16(anon_sym_list),
	304:  uint16(33),
	305:  uint16(1),
	306:  uint16(anon_sym_option),
	307:  uint16(35),
	308:  uint16(1),
	309:  uint16(anon_sym_map),
	310:  uint16(37),
	311:  uint16(1),
	312:  uint16(anon_sym_result),
	313:  uint16(41),
	314:  uint16(1),
	315:  uint16(anon_sym_borrow),
	316:  uint16(43),
	317:  uint16(1),
	318:  uint16(anon_sym_future),
	319:  uint16(45),
	320:  uint16(1),
	321:  uint16(anon_sym_stream),
	322:  uint16(51),
	323:  uint16(1),
	324:  uint16(anon_sym_GT),
	325:  uint16(244),
	326:  uint16(1),
	327:  uint16(sym_ty),
	328:  uint16(27),
	329:  uint16(2),
	330:  uint16(anon_sym_f32),
	331:  uint16(anon_sym_f64),
	332:  uint16(6),
	333:  uint16(2),
	334:  uint16(sym_line_comment),
	335:  uint16(sym_block_comment),
	336:  uint16(108),
	337:  uint16(9),
	338:  uint16(sym__primitive_ty),
	339:  uint16(sym_tuple),
	340:  uint16(sym_list),
	341:  uint16(sym_option),
	342:  uint16(sym_map),
	343:  uint16(sym_result),
	344:  uint16(sym_handle),
	345:  uint16(sym_future),
	346:  uint16(sym_stream),
	347:  uint16(25),
	348:  uint16(11),
	349:  uint16(anon_sym_u8),
	350:  uint16(anon_sym_u16),
	351:  uint16(anon_sym_u32),
	352:  uint16(anon_sym_u64),
	353:  uint16(anon_sym_s8),
	354:  uint16(anon_sym_s16),
	355:  uint16(anon_sym_s32),
	356:  uint16(anon_sym_s64),
	357:  uint16(anon_sym_char),
	358:  uint16(anon_sym_bool),
	359:  uint16(anon_sym_string),
	360:  uint16(16),
	361:  uint16(19),
	362:  uint16(1),
	363:  uint16(anon_sym_SLASH_SLASH),
	364:  uint16(21),
	365:  uint16(1),
	366:  uint16(anon_sym_SLASH_STAR),
	367:  uint16(23),
	368:  uint16(1),
	369:  uint16(sym_id),
	370:  uint16(29),
	371:  uint16(1),
	372:  uint16(anon_sym_tuple),
	373:  uint16(31),
	374:  uint16(1),
	375:  uint16(anon_sym_list),
	376:  uint16(33),
	377:  uint16(1),
	378:  uint16(anon_sym_option),
	379:  uint16(35),
	380:  uint16(1),
	381:  uint16(anon_sym_map),
	382:  uint16(37),
	383:  uint16(1),
	384:  uint16(anon_sym_result),
	385:  uint16(41),
	386:  uint16(1),
	387:  uint16(anon_sym_borrow),
	388:  uint16(43),
	389:  uint16(1),
	390:  uint16(anon_sym_future),
	391:  uint16(45),
	392:  uint16(1),
	393:  uint16(anon_sym_stream),
	394:  uint16(200),
	395:  uint16(1),
	396:  uint16(sym_ty),
	397:  uint16(27),
	398:  uint16(2),
	399:  uint16(anon_sym_f32),
	400:  uint16(anon_sym_f64),
	401:  uint16(7),
	402:  uint16(2),
	403:  uint16(sym_line_comment),
	404:  uint16(sym_block_comment),
	405:  uint16(108),
	406:  uint16(9),
	407:  uint16(sym__primitive_ty),
	408:  uint16(sym_tuple),
	409:  uint16(sym_list),
	410:  uint16(sym_option),
	411:  uint16(sym_map),
	412:  uint16(sym_result),
	413:  uint16(sym_handle),
	414:  uint16(sym_future),
	415:  uint16(sym_stream),
	416:  uint16(25),
	417:  uint16(11),
	418:  uint16(anon_sym_u8),
	419:  uint16(anon_sym_u16),
	420:  uint16(anon_sym_u32),
	421:  uint16(anon_sym_u64),
	422:  uint16(anon_sym_s8),
	423:  uint16(anon_sym_s16),
	424:  uint16(anon_sym_s32),
	425:  uint16(anon_sym_s64),
	426:  uint16(anon_sym_char),
	427:  uint16(anon_sym_bool),
	428:  uint16(anon_sym_string),
	429:  uint16(16),
	430:  uint16(19),
	431:  uint16(1),
	432:  uint16(anon_sym_SLASH_SLASH),
	433:  uint16(21),
	434:  uint16(1),
	435:  uint16(anon_sym_SLASH_STAR),
	436:  uint16(23),
	437:  uint16(1),
	438:  uint16(sym_id),
	439:  uint16(29),
	440:  uint16(1),
	441:  uint16(anon_sym_tuple),
	442:  uint16(31),
	443:  uint16(1),
	444:  uint16(anon_sym_list),
	445:  uint16(33),
	446:  uint16(1),
	447:  uint16(anon_sym_option),
	448:  uint16(35),
	449:  uint16(1),
	450:  uint16(anon_sym_map),
	451:  uint16(37),
	452:  uint16(1),
	453:  uint16(anon_sym_result),
	454:  uint16(41),
	455:  uint16(1),
	456:  uint16(anon_sym_borrow),
	457:  uint16(43),
	458:  uint16(1),
	459:  uint16(anon_sym_future),
	460:  uint16(45),
	461:  uint16(1),
	462:  uint16(anon_sym_stream),
	463:  uint16(268),
	464:  uint16(1),
	465:  uint16(sym_ty),
	466:  uint16(27),
	467:  uint16(2),
	468:  uint16(anon_sym_f32),
	469:  uint16(anon_sym_f64),
	470:  uint16(8),
	471:  uint16(2),
	472:  uint16(sym_line_comment),
	473:  uint16(sym_block_comment),
	474:  uint16(108),
	475:  uint16(9),
	476:  uint16(sym__primitive_ty),
	477:  uint16(sym_tuple),
	478:  uint16(sym_list),
	479:  uint16(sym_option),
	480:  uint16(sym_map),
	481:  uint16(sym_result),
	482:  uint16(sym_handle),
	483:  uint16(sym_future),
	484:  uint16(sym_stream),
	485:  uint16(25),
	486:  uint16(11),
	487:  uint16(anon_sym_u8),
	488:  uint16(anon_sym_u16),
	489:  uint16(anon_sym_u32),
	490:  uint16(anon_sym_u64),
	491:  uint16(anon_sym_s8),
	492:  uint16(anon_sym_s16),
	493:  uint16(anon_sym_s32),
	494:  uint16(anon_sym_s64),
	495:  uint16(anon_sym_char),
	496:  uint16(anon_sym_bool),
	497:  uint16(anon_sym_string),
	498:  uint16(16),
	499:  uint16(19),
	500:  uint16(1),
	501:  uint16(anon_sym_SLASH_SLASH),
	502:  uint16(21),
	503:  uint16(1),
	504:  uint16(anon_sym_SLASH_STAR),
	505:  uint16(23),
	506:  uint16(1),
	507:  uint16(sym_id),
	508:  uint16(29),
	509:  uint16(1),
	510:  uint16(anon_sym_tuple),
	511:  uint16(31),
	512:  uint16(1),
	513:  uint16(anon_sym_list),
	514:  uint16(33),
	515:  uint16(1),
	516:  uint16(anon_sym_option),
	517:  uint16(35),
	518:  uint16(1),
	519:  uint16(anon_sym_map),
	520:  uint16(37),
	521:  uint16(1),
	522:  uint16(anon_sym_result),
	523:  uint16(41),
	524:  uint16(1),
	525:  uint16(anon_sym_borrow),
	526:  uint16(43),
	527:  uint16(1),
	528:  uint16(anon_sym_future),
	529:  uint16(45),
	530:  uint16(1),
	531:  uint16(anon_sym_stream),
	532:  uint16(315),
	533:  uint16(1),
	534:  uint16(sym_ty),
	535:  uint16(27),
	536:  uint16(2),
	537:  uint16(anon_sym_f32),
	538:  uint16(anon_sym_f64),
	539:  uint16(9),
	540:  uint16(2),
	541:  uint16(sym_line_comment),
	542:  uint16(sym_block_comment),
	543:  uint16(108),
	544:  uint16(9),
	545:  uint16(sym__primitive_ty),
	546:  uint16(sym_tuple),
	547:  uint16(sym_list),
	548:  uint16(sym_option),
	549:  uint16(sym_map),
	550:  uint16(sym_result),
	551:  uint16(sym_handle),
	552:  uint16(sym_future),
	553:  uint16(sym_stream),
	554:  uint16(25),
	555:  uint16(11),
	556:  uint16(anon_sym_u8),
	557:  uint16(anon_sym_u16),
	558:  uint16(anon_sym_u32),
	559:  uint16(anon_sym_u64),
	560:  uint16(anon_sym_s8),
	561:  uint16(anon_sym_s16),
	562:  uint16(anon_sym_s32),
	563:  uint16(anon_sym_s64),
	564:  uint16(anon_sym_char),
	565:  uint16(anon_sym_bool),
	566:  uint16(anon_sym_string),
	567:  uint16(16),
	568:  uint16(19),
	569:  uint16(1),
	570:  uint16(anon_sym_SLASH_SLASH),
	571:  uint16(21),
	572:  uint16(1),
	573:  uint16(anon_sym_SLASH_STAR),
	574:  uint16(23),
	575:  uint16(1),
	576:  uint16(sym_id),
	577:  uint16(29),
	578:  uint16(1),
	579:  uint16(anon_sym_tuple),
	580:  uint16(31),
	581:  uint16(1),
	582:  uint16(anon_sym_list),
	583:  uint16(33),
	584:  uint16(1),
	585:  uint16(anon_sym_option),
	586:  uint16(35),
	587:  uint16(1),
	588:  uint16(anon_sym_map),
	589:  uint16(37),
	590:  uint16(1),
	591:  uint16(anon_sym_result),
	592:  uint16(41),
	593:  uint16(1),
	594:  uint16(anon_sym_borrow),
	595:  uint16(43),
	596:  uint16(1),
	597:  uint16(anon_sym_future),
	598:  uint16(45),
	599:  uint16(1),
	600:  uint16(anon_sym_stream),
	601:  uint16(274),
	602:  uint16(1),
	603:  uint16(sym_ty),
	604:  uint16(27),
	605:  uint16(2),
	606:  uint16(anon_sym_f32),
	607:  uint16(anon_sym_f64),
	608:  uint16(10),
	609:  uint16(2),
	610:  uint16(sym_line_comment),
	611:  uint16(sym_block_comment),
	612:  uint16(108),
	613:  uint16(9),
	614:  uint16(sym__primitive_ty),
	615:  uint16(sym_tuple),
	616:  uint16(sym_list),
	617:  uint16(sym_option),
	618:  uint16(sym_map),
	619:  uint16(sym_result),
	620:  uint16(sym_handle),
	621:  uint16(sym_future),
	622:  uint16(sym_stream),
	623:  uint16(25),
	624:  uint16(11),
	625:  uint16(anon_sym_u8),
	626:  uint16(anon_sym_u16),
	627:  uint16(anon_sym_u32),
	628:  uint16(anon_sym_u64),
	629:  uint16(anon_sym_s8),
	630:  uint16(anon_sym_s16),
	631:  uint16(anon_sym_s32),
	632:  uint16(anon_sym_s64),
	633:  uint16(anon_sym_char),
	634:  uint16(anon_sym_bool),
	635:  uint16(anon_sym_string),
	636:  uint16(16),
	637:  uint16(19),
	638:  uint16(1),
	639:  uint16(anon_sym_SLASH_SLASH),
	640:  uint16(21),
	641:  uint16(1),
	642:  uint16(anon_sym_SLASH_STAR),
	643:  uint16(23),
	644:  uint16(1),
	645:  uint16(sym_id),
	646:  uint16(29),
	647:  uint16(1),
	648:  uint16(anon_sym_tuple),
	649:  uint16(31),
	650:  uint16(1),
	651:  uint16(anon_sym_list),
	652:  uint16(33),
	653:  uint16(1),
	654:  uint16(anon_sym_option),
	655:  uint16(35),
	656:  uint16(1),
	657:  uint16(anon_sym_map),
	658:  uint16(37),
	659:  uint16(1),
	660:  uint16(anon_sym_result),
	661:  uint16(41),
	662:  uint16(1),
	663:  uint16(anon_sym_borrow),
	664:  uint16(43),
	665:  uint16(1),
	666:  uint16(anon_sym_future),
	667:  uint16(45),
	668:  uint16(1),
	669:  uint16(anon_sym_stream),
	670:  uint16(275),
	671:  uint16(1),
	672:  uint16(sym_ty),
	673:  uint16(27),
	674:  uint16(2),
	675:  uint16(anon_sym_f32),
	676:  uint16(anon_sym_f64),
	677:  uint16(11),
	678:  uint16(2),
	679:  uint16(sym_line_comment),
	680:  uint16(sym_block_comment),
	681:  uint16(108),
	682:  uint16(9),
	683:  uint16(sym__primitive_ty),
	684:  uint16(sym_tuple),
	685:  uint16(sym_list),
	686:  uint16(sym_option),
	687:  uint16(sym_map),
	688:  uint16(sym_result),
	689:  uint16(sym_handle),
	690:  uint16(sym_future),
	691:  uint16(sym_stream),
	692:  uint16(25),
	693:  uint16(11),
	694:  uint16(anon_sym_u8),
	695:  uint16(anon_sym_u16),
	696:  uint16(anon_sym_u32),
	697:  uint16(anon_sym_u64),
	698:  uint16(anon_sym_s8),
	699:  uint16(anon_sym_s16),
	700:  uint16(anon_sym_s32),
	701:  uint16(anon_sym_s64),
	702:  uint16(anon_sym_char),
	703:  uint16(anon_sym_bool),
	704:  uint16(anon_sym_string),
	705:  uint16(16),
	706:  uint16(19),
	707:  uint16(1),
	708:  uint16(anon_sym_SLASH_SLASH),
	709:  uint16(21),
	710:  uint16(1),
	711:  uint16(anon_sym_SLASH_STAR),
	712:  uint16(23),
	713:  uint16(1),
	714:  uint16(sym_id),
	715:  uint16(29),
	716:  uint16(1),
	717:  uint16(anon_sym_tuple),
	718:  uint16(31),
	719:  uint16(1),
	720:  uint16(anon_sym_list),
	721:  uint16(33),
	722:  uint16(1),
	723:  uint16(anon_sym_option),
	724:  uint16(35),
	725:  uint16(1),
	726:  uint16(anon_sym_map),
	727:  uint16(37),
	728:  uint16(1),
	729:  uint16(anon_sym_result),
	730:  uint16(41),
	731:  uint16(1),
	732:  uint16(anon_sym_borrow),
	733:  uint16(43),
	734:  uint16(1),
	735:  uint16(anon_sym_future),
	736:  uint16(45),
	737:  uint16(1),
	738:  uint16(anon_sym_stream),
	739:  uint16(244),
	740:  uint16(1),
	741:  uint16(sym_ty),
	742:  uint16(27),
	743:  uint16(2),
	744:  uint16(anon_sym_f32),
	745:  uint16(anon_sym_f64),
	746:  uint16(12),
	747:  uint16(2),
	748:  uint16(sym_line_comment),
	749:  uint16(sym_block_comment),
	750:  uint16(108),
	751:  uint16(9),
	752:  uint16(sym__primitive_ty),
	753:  uint16(sym_tuple),
	754:  uint16(sym_list),
	755:  uint16(sym_option),
	756:  uint16(sym_map),
	757:  uint16(sym_result),
	758:  uint16(sym_handle),
	759:  uint16(sym_future),
	760:  uint16(sym_stream),
	761:  uint16(25),
	762:  uint16(11),
	763:  uint16(anon_sym_u8),
	764:  uint16(anon_sym_u16),
	765:  uint16(anon_sym_u32),
	766:  uint16(anon_sym_u64),
	767:  uint16(anon_sym_s8),
	768:  uint16(anon_sym_s16),
	769:  uint16(anon_sym_s32),
	770:  uint16(anon_sym_s64),
	771:  uint16(anon_sym_char),
	772:  uint16(anon_sym_bool),
	773:  uint16(anon_sym_string),
	774:  uint16(16),
	775:  uint16(19),
	776:  uint16(1),
	777:  uint16(anon_sym_SLASH_SLASH),
	778:  uint16(21),
	779:  uint16(1),
	780:  uint16(anon_sym_SLASH_STAR),
	781:  uint16(23),
	782:  uint16(1),
	783:  uint16(sym_id),
	784:  uint16(29),
	785:  uint16(1),
	786:  uint16(anon_sym_tuple),
	787:  uint16(31),
	788:  uint16(1),
	789:  uint16(anon_sym_list),
	790:  uint16(33),
	791:  uint16(1),
	792:  uint16(anon_sym_option),
	793:  uint16(35),
	794:  uint16(1),
	795:  uint16(anon_sym_map),
	796:  uint16(37),
	797:  uint16(1),
	798:  uint16(anon_sym_result),
	799:  uint16(41),
	800:  uint16(1),
	801:  uint16(anon_sym_borrow),
	802:  uint16(43),
	803:  uint16(1),
	804:  uint16(anon_sym_future),
	805:  uint16(45),
	806:  uint16(1),
	807:  uint16(anon_sym_stream),
	808:  uint16(202),
	809:  uint16(1),
	810:  uint16(sym_ty),
	811:  uint16(27),
	812:  uint16(2),
	813:  uint16(anon_sym_f32),
	814:  uint16(anon_sym_f64),
	815:  uint16(13),
	816:  uint16(2),
	817:  uint16(sym_line_comment),
	818:  uint16(sym_block_comment),
	819:  uint16(108),
	820:  uint16(9),
	821:  uint16(sym__primitive_ty),
	822:  uint16(sym_tuple),
	823:  uint16(sym_list),
	824:  uint16(sym_option),
	825:  uint16(sym_map),
	826:  uint16(sym_result),
	827:  uint16(sym_handle),
	828:  uint16(sym_future),
	829:  uint16(sym_stream),
	830:  uint16(25),
	831:  uint16(11),
	832:  uint16(anon_sym_u8),
	833:  uint16(anon_sym_u16),
	834:  uint16(anon_sym_u32),
	835:  uint16(anon_sym_u64),
	836:  uint16(anon_sym_s8),
	837:  uint16(anon_sym_s16),
	838:  uint16(anon_sym_s32),
	839:  uint16(anon_sym_s64),
	840:  uint16(anon_sym_char),
	841:  uint16(anon_sym_bool),
	842:  uint16(anon_sym_string),
	843:  uint16(16),
	844:  uint16(19),
	845:  uint16(1),
	846:  uint16(anon_sym_SLASH_SLASH),
	847:  uint16(21),
	848:  uint16(1),
	849:  uint16(anon_sym_SLASH_STAR),
	850:  uint16(23),
	851:  uint16(1),
	852:  uint16(sym_id),
	853:  uint16(29),
	854:  uint16(1),
	855:  uint16(anon_sym_tuple),
	856:  uint16(31),
	857:  uint16(1),
	858:  uint16(anon_sym_list),
	859:  uint16(33),
	860:  uint16(1),
	861:  uint16(anon_sym_option),
	862:  uint16(35),
	863:  uint16(1),
	864:  uint16(anon_sym_map),
	865:  uint16(37),
	866:  uint16(1),
	867:  uint16(anon_sym_result),
	868:  uint16(41),
	869:  uint16(1),
	870:  uint16(anon_sym_borrow),
	871:  uint16(43),
	872:  uint16(1),
	873:  uint16(anon_sym_future),
	874:  uint16(45),
	875:  uint16(1),
	876:  uint16(anon_sym_stream),
	877:  uint16(282),
	878:  uint16(1),
	879:  uint16(sym_ty),
	880:  uint16(27),
	881:  uint16(2),
	882:  uint16(anon_sym_f32),
	883:  uint16(anon_sym_f64),
	884:  uint16(14),
	885:  uint16(2),
	886:  uint16(sym_line_comment),
	887:  uint16(sym_block_comment),
	888:  uint16(108),
	889:  uint16(9),
	890:  uint16(sym__primitive_ty),
	891:  uint16(sym_tuple),
	892:  uint16(sym_list),
	893:  uint16(sym_option),
	894:  uint16(sym_map),
	895:  uint16(sym_result),
	896:  uint16(sym_handle),
	897:  uint16(sym_future),
	898:  uint16(sym_stream),
	899:  uint16(25),
	900:  uint16(11),
	901:  uint16(anon_sym_u8),
	902:  uint16(anon_sym_u16),
	903:  uint16(anon_sym_u32),
	904:  uint16(anon_sym_u64),
	905:  uint16(anon_sym_s8),
	906:  uint16(anon_sym_s16),
	907:  uint16(anon_sym_s32),
	908:  uint16(anon_sym_s64),
	909:  uint16(anon_sym_char),
	910:  uint16(anon_sym_bool),
	911:  uint16(anon_sym_string),
	912:  uint16(16),
	913:  uint16(19),
	914:  uint16(1),
	915:  uint16(anon_sym_SLASH_SLASH),
	916:  uint16(21),
	917:  uint16(1),
	918:  uint16(anon_sym_SLASH_STAR),
	919:  uint16(23),
	920:  uint16(1),
	921:  uint16(sym_id),
	922:  uint16(29),
	923:  uint16(1),
	924:  uint16(anon_sym_tuple),
	925:  uint16(31),
	926:  uint16(1),
	927:  uint16(anon_sym_list),
	928:  uint16(33),
	929:  uint16(1),
	930:  uint16(anon_sym_option),
	931:  uint16(35),
	932:  uint16(1),
	933:  uint16(anon_sym_map),
	934:  uint16(37),
	935:  uint16(1),
	936:  uint16(anon_sym_result),
	937:  uint16(41),
	938:  uint16(1),
	939:  uint16(anon_sym_borrow),
	940:  uint16(43),
	941:  uint16(1),
	942:  uint16(anon_sym_future),
	943:  uint16(45),
	944:  uint16(1),
	945:  uint16(anon_sym_stream),
	946:  uint16(238),
	947:  uint16(1),
	948:  uint16(sym_ty),
	949:  uint16(27),
	950:  uint16(2),
	951:  uint16(anon_sym_f32),
	952:  uint16(anon_sym_f64),
	953:  uint16(15),
	954:  uint16(2),
	955:  uint16(sym_line_comment),
	956:  uint16(sym_block_comment),
	957:  uint16(108),
	958:  uint16(9),
	959:  uint16(sym__primitive_ty),
	960:  uint16(sym_tuple),
	961:  uint16(sym_list),
	962:  uint16(sym_option),
	963:  uint16(sym_map),
	964:  uint16(sym_result),
	965:  uint16(sym_handle),
	966:  uint16(sym_future),
	967:  uint16(sym_stream),
	968:  uint16(25),
	969:  uint16(11),
	970:  uint16(anon_sym_u8),
	971:  uint16(anon_sym_u16),
	972:  uint16(anon_sym_u32),
	973:  uint16(anon_sym_u64),
	974:  uint16(anon_sym_s8),
	975:  uint16(anon_sym_s16),
	976:  uint16(anon_sym_s32),
	977:  uint16(anon_sym_s64),
	978:  uint16(anon_sym_char),
	979:  uint16(anon_sym_bool),
	980:  uint16(anon_sym_string),
	981:  uint16(16),
	982:  uint16(19),
	983:  uint16(1),
	984:  uint16(anon_sym_SLASH_SLASH),
	985:  uint16(21),
	986:  uint16(1),
	987:  uint16(anon_sym_SLASH_STAR),
	988:  uint16(23),
	989:  uint16(1),
	990:  uint16(sym_id),
	991:  uint16(29),
	992:  uint16(1),
	993:  uint16(anon_sym_tuple),
	994:  uint16(31),
	995:  uint16(1),
	996:  uint16(anon_sym_list),
	997:  uint16(33),
	998:  uint16(1),
	999:  uint16(anon_sym_option),
	1000: uint16(35),
	1001: uint16(1),
	1002: uint16(anon_sym_map),
	1003: uint16(37),
	1004: uint16(1),
	1005: uint16(anon_sym_result),
	1006: uint16(41),
	1007: uint16(1),
	1008: uint16(anon_sym_borrow),
	1009: uint16(43),
	1010: uint16(1),
	1011: uint16(anon_sym_future),
	1012: uint16(45),
	1013: uint16(1),
	1014: uint16(anon_sym_stream),
	1015: uint16(333),
	1016: uint16(1),
	1017: uint16(sym_ty),
	1018: uint16(27),
	1019: uint16(2),
	1020: uint16(anon_sym_f32),
	1021: uint16(anon_sym_f64),
	1022: uint16(16),
	1023: uint16(2),
	1024: uint16(sym_line_comment),
	1025: uint16(sym_block_comment),
	1026: uint16(108),
	1027: uint16(9),
	1028: uint16(sym__primitive_ty),
	1029: uint16(sym_tuple),
	1030: uint16(sym_list),
	1031: uint16(sym_option),
	1032: uint16(sym_map),
	1033: uint16(sym_result),
	1034: uint16(sym_handle),
	1035: uint16(sym_future),
	1036: uint16(sym_stream),
	1037: uint16(25),
	1038: uint16(11),
	1039: uint16(anon_sym_u8),
	1040: uint16(anon_sym_u16),
	1041: uint16(anon_sym_u32),
	1042: uint16(anon_sym_u64),
	1043: uint16(anon_sym_s8),
	1044: uint16(anon_sym_s16),
	1045: uint16(anon_sym_s32),
	1046: uint16(anon_sym_s64),
	1047: uint16(anon_sym_char),
	1048: uint16(anon_sym_bool),
	1049: uint16(anon_sym_string),
	1050: uint16(16),
	1051: uint16(19),
	1052: uint16(1),
	1053: uint16(anon_sym_SLASH_SLASH),
	1054: uint16(21),
	1055: uint16(1),
	1056: uint16(anon_sym_SLASH_STAR),
	1057: uint16(23),
	1058: uint16(1),
	1059: uint16(sym_id),
	1060: uint16(29),
	1061: uint16(1),
	1062: uint16(anon_sym_tuple),
	1063: uint16(31),
	1064: uint16(1),
	1065: uint16(anon_sym_list),
	1066: uint16(33),
	1067: uint16(1),
	1068: uint16(anon_sym_option),
	1069: uint16(35),
	1070: uint16(1),
	1071: uint16(anon_sym_map),
	1072: uint16(37),
	1073: uint16(1),
	1074: uint16(anon_sym_result),
	1075: uint16(41),
	1076: uint16(1),
	1077: uint16(anon_sym_borrow),
	1078: uint16(43),
	1079: uint16(1),
	1080: uint16(anon_sym_future),
	1081: uint16(45),
	1082: uint16(1),
	1083: uint16(anon_sym_stream),
	1084: uint16(334),
	1085: uint16(1),
	1086: uint16(sym_ty),
	1087: uint16(27),
	1088: uint16(2),
	1089: uint16(anon_sym_f32),
	1090: uint16(anon_sym_f64),
	1091: uint16(17),
	1092: uint16(2),
	1093: uint16(sym_line_comment),
	1094: uint16(sym_block_comment),
	1095: uint16(108),
	1096: uint16(9),
	1097: uint16(sym__primitive_ty),
	1098: uint16(sym_tuple),
	1099: uint16(sym_list),
	1100: uint16(sym_option),
	1101: uint16(sym_map),
	1102: uint16(sym_result),
	1103: uint16(sym_handle),
	1104: uint16(sym_future),
	1105: uint16(sym_stream),
	1106: uint16(25),
	1107: uint16(11),
	1108: uint16(anon_sym_u8),
	1109: uint16(anon_sym_u16),
	1110: uint16(anon_sym_u32),
	1111: uint16(anon_sym_u64),
	1112: uint16(anon_sym_s8),
	1113: uint16(anon_sym_s16),
	1114: uint16(anon_sym_s32),
	1115: uint16(anon_sym_s64),
	1116: uint16(anon_sym_char),
	1117: uint16(anon_sym_bool),
	1118: uint16(anon_sym_string),
	1119: uint16(22),
	1120: uint16(19),
	1121: uint16(1),
	1122: uint16(anon_sym_SLASH_SLASH),
	1123: uint16(21),
	1124: uint16(1),
	1125: uint16(anon_sym_SLASH_STAR),
	1126: uint16(53),
	1127: uint16(1),
	1128: uint16(anon_sym_RBRACE),
	1129: uint16(55),
	1130: uint16(1),
	1131: uint16(anon_sym_AT),
	1132: uint16(58),
	1133: uint16(1),
	1134: uint16(anon_sym_use),
	1135: uint16(61),
	1136: uint16(1),
	1137: uint16(anon_sym_export),
	1138: uint16(64),
	1139: uint16(1),
	1140: uint16(anon_sym_import),
	1141: uint16(67),
	1142: uint16(1),
	1143: uint16(anon_sym_include),
	1144: uint16(70),
	1145: uint16(1),
	1146: uint16(anon_sym_type),
	1147: uint16(73),
	1148: uint16(1),
	1149: uint16(anon_sym_record),
	1150: uint16(76),
	1151: uint16(1),
	1152: uint16(anon_sym_flags),
	1153: uint16(79),
	1154: uint16(1),
	1155: uint16(anon_sym_variant),
	1156: uint16(82),
	1157: uint16(1),
	1158: uint16(anon_sym_enum),
	1159: uint16(85),
	1160: uint16(1),
	1161: uint16(anon_sym_resource),
	1162: uint16(21),
	1163: uint16(1),
	1164: uint16(aux_sym__gate),
	1165: uint16(37),
	1166: uint16(1),
	1167: uint16(sym__gate_item),
	1168: uint16(64),
	1169: uint16(1),
	1170: uint16(sym__world_items),
	1171: uint16(225),
	1172: uint16(1),
	1173: uint16(sym_external_id),
	1174: uint16(18),
	1175: uint16(3),
	1176: uint16(sym_line_comment),
	1177: uint16(sym_block_comment),
	1178: uint16(aux_sym__world_body_repeat1),
	1179: uint16(38),
	1180: uint16(3),
	1181: uint16(sym_unstable_gate),
	1182: uint16(sym_since_gate),
	1183: uint16(sym_deprecated_gate),
	1184: uint16(65),
	1185: uint16(5),
	1186: uint16(sym_export_item),
	1187: uint16(sym_import_item),
	1188: uint16(sym_include_item),
	1189: uint16(sym__typedef_item),
	1190: uint16(sym_use_item),
	1191: uint16(55),
	1192: uint16(6),
	1193: uint16(sym_type_item),
	1194: uint16(sym_record_item),
	1195: uint16(sym_flags_items),
	1196: uint16(sym_variant_items),
	1197: uint16(sym_enum_items),
	1198: uint16(sym_resource_item),
	1199: uint16(23),
	1200: uint16(19),
	1201: uint16(1),
	1202: uint16(anon_sym_SLASH_SLASH),
	1203: uint16(21),
	1204: uint16(1),
	1205: uint16(anon_sym_SLASH_STAR),
	1206: uint16(88),
	1207: uint16(1),
	1208: uint16(anon_sym_RBRACE),
	1209: uint16(90),
	1210: uint16(1),
	1211: uint16(anon_sym_AT),
	1212: uint16(92),
	1213: uint16(1),
	1214: uint16(anon_sym_use),
	1215: uint16(94),
	1216: uint16(1),
	1217: uint16(anon_sym_export),
	1218: uint16(96),
	1219: uint16(1),
	1220: uint16(anon_sym_import),
	1221: uint16(98),
	1222: uint16(1),
	1223: uint16(anon_sym_include),
	1224: uint16(100),
	1225: uint16(1),
	1226: uint16(anon_sym_type),
	1227: uint16(102),
	1228: uint16(1),
	1229: uint16(anon_sym_record),
	1230: uint16(104),
	1231: uint16(1),
	1232: uint16(anon_sym_flags),
	1233: uint16(106),
	1234: uint16(1),
	1235: uint16(anon_sym_variant),
	1236: uint16(108),
	1237: uint16(1),
	1238: uint16(anon_sym_enum),
	1239: uint16(110),
	1240: uint16(1),
	1241: uint16(anon_sym_resource),
	1242: uint16(20),
	1243: uint16(1),
	1244: uint16(aux_sym__world_body_repeat1),
	1245: uint16(21),
	1246: uint16(1),
	1247: uint16(aux_sym__gate),
	1248: uint16(37),
	1249: uint16(1),
	1250: uint16(sym__gate_item),
	1251: uint16(64),
	1252: uint16(1),
	1253: uint16(sym__world_items),
	1254: uint16(225),
	1255: uint16(1),
	1256: uint16(sym_external_id),
	1257: uint16(19),
	1258: uint16(2),
	1259: uint16(sym_line_comment),
	1260: uint16(sym_block_comment),
	1261: uint16(38),
	1262: uint16(3),
	1263: uint16(sym_unstable_gate),
	1264: uint16(sym_since_gate),
	1265: uint16(sym_deprecated_gate),
	1266: uint16(65),
	1267: uint16(5),
	1268: uint16(sym_export_item),
	1269: uint16(sym_import_item),
	1270: uint16(sym_include_item),
	1271: uint16(sym__typedef_item),
	1272: uint16(sym_use_item),
	1273: uint16(55),
	1274: uint16(6),
	1275: uint16(sym_type_item),
	1276: uint16(sym_record_item),
	1277: uint16(sym_flags_items),
	1278: uint16(sym_variant_items),
	1279: uint16(sym_enum_items),
	1280: uint16(sym_resource_item),
	1281: uint16(23),
	1282: uint16(19),
	1283: uint16(1),
	1284: uint16(anon_sym_SLASH_SLASH),
	1285: uint16(21),
	1286: uint16(1),
	1287: uint16(anon_sym_SLASH_STAR),
	1288: uint16(90),
	1289: uint16(1),
	1290: uint16(anon_sym_AT),
	1291: uint16(92),
	1292: uint16(1),
	1293: uint16(anon_sym_use),
	1294: uint16(94),
	1295: uint16(1),
	1296: uint16(anon_sym_export),
	1297: uint16(96),
	1298: uint16(1),
	1299: uint16(anon_sym_import),
	1300: uint16(98),
	1301: uint16(1),
	1302: uint16(anon_sym_include),
	1303: uint16(100),
	1304: uint16(1),
	1305: uint16(anon_sym_type),
	1306: uint16(102),
	1307: uint16(1),
	1308: uint16(anon_sym_record),
	1309: uint16(104),
	1310: uint16(1),
	1311: uint16(anon_sym_flags),
	1312: uint16(106),
	1313: uint16(1),
	1314: uint16(anon_sym_variant),
	1315: uint16(108),
	1316: uint16(1),
	1317: uint16(anon_sym_enum),
	1318: uint16(110),
	1319: uint16(1),
	1320: uint16(anon_sym_resource),
	1321: uint16(112),
	1322: uint16(1),
	1323: uint16(anon_sym_RBRACE),
	1324: uint16(18),
	1325: uint16(1),
	1326: uint16(aux_sym__world_body_repeat1),
	1327: uint16(21),
	1328: uint16(1),
	1329: uint16(aux_sym__gate),
	1330: uint16(37),
	1331: uint16(1),
	1332: uint16(sym__gate_item),
	1333: uint16(64),
	1334: uint16(1),
	1335: uint16(sym__world_items),
	1336: uint16(225),
	1337: uint16(1),
	1338: uint16(sym_external_id),
	1339: uint16(20),
	1340: uint16(2),
	1341: uint16(sym_line_comment),
	1342: uint16(sym_block_comment),
	1343: uint16(38),
	1344: uint16(3),
	1345: uint16(sym_unstable_gate),
	1346: uint16(sym_since_gate),
	1347: uint16(sym_deprecated_gate),
	1348: uint16(65),
	1349: uint16(5),
	1350: uint16(sym_export_item),
	1351: uint16(sym_import_item),
	1352: uint16(sym_include_item),
	1353: uint16(sym__typedef_item),
	1354: uint16(sym_use_item),
	1355: uint16(55),
	1356: uint16(6),
	1357: uint16(sym_type_item),
	1358: uint16(sym_record_item),
	1359: uint16(sym_flags_items),
	1360: uint16(sym_variant_items),
	1361: uint16(sym_enum_items),
	1362: uint16(sym_resource_item),
	1363: uint16(20),
	1364: uint16(19),
	1365: uint16(1),
	1366: uint16(anon_sym_SLASH_SLASH),
	1367: uint16(21),
	1368: uint16(1),
	1369: uint16(anon_sym_SLASH_STAR),
	1370: uint16(90),
	1371: uint16(1),
	1372: uint16(anon_sym_AT),
	1373: uint16(92),
	1374: uint16(1),
	1375: uint16(anon_sym_use),
	1376: uint16(94),
	1377: uint16(1),
	1378: uint16(anon_sym_export),
	1379: uint16(96),
	1380: uint16(1),
	1381: uint16(anon_sym_import),
	1382: uint16(98),
	1383: uint16(1),
	1384: uint16(anon_sym_include),
	1385: uint16(100),
	1386: uint16(1),
	1387: uint16(anon_sym_type),
	1388: uint16(102),
	1389: uint16(1),
	1390: uint16(anon_sym_record),
	1391: uint16(104),
	1392: uint16(1),
	1393: uint16(anon_sym_flags),
	1394: uint16(106),
	1395: uint16(1),
	1396: uint16(anon_sym_variant),
	1397: uint16(108),
	1398: uint16(1),
	1399: uint16(anon_sym_enum),
	1400: uint16(110),
	1401: uint16(1),
	1402: uint16(anon_sym_resource),
	1403: uint16(28),
	1404: uint16(1),
	1405: uint16(aux_sym__gate),
	1406: uint16(37),
	1407: uint16(1),
	1408: uint16(sym__gate_item),
	1409: uint16(225),
	1410: uint16(1),
	1411: uint16(sym_external_id),
	1412: uint16(21),
	1413: uint16(2),
	1414: uint16(sym_line_comment),
	1415: uint16(sym_block_comment),
	1416: uint16(38),
	1417: uint16(3),
	1418: uint16(sym_unstable_gate),
	1419: uint16(sym_since_gate),
	1420: uint16(sym_deprecated_gate),
	1421: uint16(58),
	1422: uint16(5),
	1423: uint16(sym_export_item),
	1424: uint16(sym_import_item),
	1425: uint16(sym_include_item),
	1426: uint16(sym__typedef_item),
	1427: uint16(sym_use_item),
	1428: uint16(55),
	1429: uint16(6),
	1430: uint16(sym_type_item),
	1431: uint16(sym_record_item),
	1432: uint16(sym_flags_items),
	1433: uint16(sym_variant_items),
	1434: uint16(sym_enum_items),
	1435: uint16(sym_resource_item),
	1436: uint16(20),
	1437: uint16(19),
	1438: uint16(1),
	1439: uint16(anon_sym_SLASH_SLASH),
	1440: uint16(21),
	1441: uint16(1),
	1442: uint16(anon_sym_SLASH_STAR),
	1443: uint16(114),
	1444: uint16(1),
	1445: uint16(sym_id),
	1446: uint16(117),
	1447: uint16(1),
	1448: uint16(anon_sym_RBRACE),
	1449: uint16(119),
	1450: uint16(1),
	1451: uint16(anon_sym_AT),
	1452: uint16(122),
	1453: uint16(1),
	1454: uint16(anon_sym_use),
	1455: uint16(125),
	1456: uint16(1),
	1457: uint16(anon_sym_type),
	1458: uint16(128),
	1459: uint16(1),
	1460: uint16(anon_sym_record),
	1461: uint16(131),
	1462: uint16(1),
	1463: uint16(anon_sym_flags),
	1464: uint16(134),
	1465: uint16(1),
	1466: uint16(anon_sym_variant),
	1467: uint16(137),
	1468: uint16(1),
	1469: uint16(anon_sym_enum),
	1470: uint16(140),
	1471: uint16(1),
	1472: uint16(anon_sym_resource),
	1473: uint16(25),
	1474: uint16(1),
	1475: uint16(aux_sym__gate),
	1476: uint16(35),
	1477: uint16(1),
	1478: uint16(sym_external_id),
	1479: uint16(37),
	1480: uint16(1),
	1481: uint16(sym__gate_item),
	1482: uint16(73),
	1483: uint16(1),
	1484: uint16(sym__interface_items),
	1485: uint16(22),
	1486: uint16(3),
	1487: uint16(sym_line_comment),
	1488: uint16(sym_block_comment),
	1489: uint16(aux_sym__interface_body_repeat1),
	1490: uint16(38),
	1491: uint16(3),
	1492: uint16(sym_unstable_gate),
	1493: uint16(sym_since_gate),
	1494: uint16(sym_deprecated_gate),
	1495: uint16(74),
	1496: uint16(3),
	1497: uint16(sym__typedef_item),
	1498: uint16(sym_func_item),
	1499: uint16(sym_use_item),
	1500: uint16(55),
	1501: uint16(6),
	1502: uint16(sym_type_item),
	1503: uint16(sym_record_item),
	1504: uint16(sym_flags_items),
	1505: uint16(sym_variant_items),
	1506: uint16(sym_enum_items),
	1507: uint16(sym_resource_item),
	1508: uint16(21),
	1509: uint16(19),
	1510: uint16(1),
	1511: uint16(anon_sym_SLASH_SLASH),
	1512: uint16(21),
	1513: uint16(1),
	1514: uint16(anon_sym_SLASH_STAR),
	1515: uint16(90),
	1516: uint16(1),
	1517: uint16(anon_sym_AT),
	1518: uint16(143),
	1519: uint16(1),
	1520: uint16(sym_id),
	1521: uint16(145),
	1522: uint16(1),
	1523: uint16(anon_sym_RBRACE),
	1524: uint16(147),
	1525: uint16(1),
	1526: uint16(anon_sym_use),
	1527: uint16(149),
	1528: uint16(1),
	1529: uint16(anon_sym_type),
	1530: uint16(151),
	1531: uint16(1),
	1532: uint16(anon_sym_record),
	1533: uint16(153),
	1534: uint16(1),
	1535: uint16(anon_sym_flags),
	1536: uint16(155),
	1537: uint16(1),
	1538: uint16(anon_sym_variant),
	1539: uint16(157),
	1540: uint16(1),
	1541: uint16(anon_sym_enum),
	1542: uint16(159),
	1543: uint16(1),
	1544: uint16(anon_sym_resource),
	1545: uint16(24),
	1546: uint16(1),
	1547: uint16(aux_sym__interface_body_repeat1),
	1548: uint16(25),
	1549: uint16(1),
	1550: uint16(aux_sym__gate),
	1551: uint16(35),
	1552: uint16(1),
	1553: uint16(sym_external_id),
	1554: uint16(37),
	1555: uint16(1),
	1556: uint16(sym__gate_item),
	1557: uint16(73),
	1558: uint16(1),
	1559: uint16(sym__interface_items),
	1560: uint16(23),
	1561: uint16(2),
	1562: uint16(sym_line_comment),
	1563: uint16(sym_block_comment),
	1564: uint16(38),
	1565: uint16(3),
	1566: uint16(sym_unstable_gate),
	1567: uint16(sym_since_gate),
	1568: uint16(sym_deprecated_gate),
	1569: uint16(74),
	1570: uint16(3),
	1571: uint16(sym__typedef_item),
	1572: uint16(sym_func_item),
	1573: uint16(sym_use_item),
	1574: uint16(55),
	1575: uint16(6),
	1576: uint16(sym_type_item),
	1577: uint16(sym_record_item),
	1578: uint16(sym_flags_items),
	1579: uint16(sym_variant_items),
	1580: uint16(sym_enum_items),
	1581: uint16(sym_resource_item),
	1582: uint16(21),
	1583: uint16(19),
	1584: uint16(1),
	1585: uint16(anon_sym_SLASH_SLASH),
	1586: uint16(21),
	1587: uint16(1),
	1588: uint16(anon_sym_SLASH_STAR),
	1589: uint16(90),
	1590: uint16(1),
	1591: uint16(anon_sym_AT),
	1592: uint16(143),
	1593: uint16(1),
	1594: uint16(sym_id),
	1595: uint16(147),
	1596: uint16(1),
	1597: uint16(anon_sym_use),
	1598: uint16(149),
	1599: uint16(1),
	1600: uint16(anon_sym_type),
	1601: uint16(151),
	1602: uint16(1),
	1603: uint16(anon_sym_record),
	1604: uint16(153),
	1605: uint16(1),
	1606: uint16(anon_sym_flags),
	1607: uint16(155),
	1608: uint16(1),
	1609: uint16(anon_sym_variant),
	1610: uint16(157),
	1611: uint16(1),
	1612: uint16(anon_sym_enum),
	1613: uint16(159),
	1614: uint16(1),
	1615: uint16(anon_sym_resource),
	1616: uint16(161),
	1617: uint16(1),
	1618: uint16(anon_sym_RBRACE),
	1619: uint16(22),
	1620: uint16(1),
	1621: uint16(aux_sym__interface_body_repeat1),
	1622: uint16(25),
	1623: uint16(1),
	1624: uint16(aux_sym__gate),
	1625: uint16(35),
	1626: uint16(1),
	1627: uint16(sym_external_id),
	1628: uint16(37),
	1629: uint16(1),
	1630: uint16(sym__gate_item),
	1631: uint16(73),
	1632: uint16(1),
	1633: uint16(sym__interface_items),
	1634: uint16(24),
	1635: uint16(2),
	1636: uint16(sym_line_comment),
	1637: uint16(sym_block_comment),
	1638: uint16(38),
	1639: uint16(3),
	1640: uint16(sym_unstable_gate),
	1641: uint16(sym_since_gate),
	1642: uint16(sym_deprecated_gate),
	1643: uint16(74),
	1644: uint16(3),
	1645: uint16(sym__typedef_item),
	1646: uint16(sym_func_item),
	1647: uint16(sym_use_item),
	1648: uint16(55),
	1649: uint16(6),
	1650: uint16(sym_type_item),
	1651: uint16(sym_record_item),
	1652: uint16(sym_flags_items),
	1653: uint16(sym_variant_items),
	1654: uint16(sym_enum_items),
	1655: uint16(sym_resource_item),
	1656: uint16(18),
	1657: uint16(19),
	1658: uint16(1),
	1659: uint16(anon_sym_SLASH_SLASH),
	1660: uint16(21),
	1661: uint16(1),
	1662: uint16(anon_sym_SLASH_STAR),
	1663: uint16(90),
	1664: uint16(1),
	1665: uint16(anon_sym_AT),
	1666: uint16(143),
	1667: uint16(1),
	1668: uint16(sym_id),
	1669: uint16(147),
	1670: uint16(1),
	1671: uint16(anon_sym_use),
	1672: uint16(149),
	1673: uint16(1),
	1674: uint16(anon_sym_type),
	1675: uint16(151),
	1676: uint16(1),
	1677: uint16(anon_sym_record),
	1678: uint16(153),
	1679: uint16(1),
	1680: uint16(anon_sym_flags),
	1681: uint16(155),
	1682: uint16(1),
	1683: uint16(anon_sym_variant),
	1684: uint16(157),
	1685: uint16(1),
	1686: uint16(anon_sym_enum),
	1687: uint16(159),
	1688: uint16(1),
	1689: uint16(anon_sym_resource),
	1690: uint16(28),
	1691: uint16(1),
	1692: uint16(aux_sym__gate),
	1693: uint16(32),
	1694: uint16(1),
	1695: uint16(sym_external_id),
	1696: uint16(37),
	1697: uint16(1),
	1698: uint16(sym__gate_item),
	1699: uint16(25),
	1700: uint16(2),
	1701: uint16(sym_line_comment),
	1702: uint16(sym_block_comment),
	1703: uint16(38),
	1704: uint16(3),
	1705: uint16(sym_unstable_gate),
	1706: uint16(sym_since_gate),
	1707: uint16(sym_deprecated_gate),
	1708: uint16(71),
	1709: uint16(3),
	1710: uint16(sym__typedef_item),
	1711: uint16(sym_func_item),
	1712: uint16(sym_use_item),
	1713: uint16(55),
	1714: uint16(6),
	1715: uint16(sym_type_item),
	1716: uint16(sym_record_item),
	1717: uint16(sym_flags_items),
	1718: uint16(sym_variant_items),
	1719: uint16(sym_enum_items),
	1720: uint16(sym_resource_item),
	1721: uint16(16),
	1722: uint16(19),
	1723: uint16(1),
	1724: uint16(anon_sym_SLASH_SLASH),
	1725: uint16(21),
	1726: uint16(1),
	1727: uint16(anon_sym_SLASH_STAR),
	1728: uint16(163),
	1729: uint16(1),
	1731: uint16(165),
	1732: uint16(1),
	1733: uint16(anon_sym_AT),
	1734: uint16(168),
	1735: uint16(1),
	1736: uint16(anon_sym_package),
	1737: uint16(171),
	1738: uint16(1),
	1739: uint16(anon_sym_use),
	1740: uint16(174),
	1741: uint16(1),
	1742: uint16(anon_sym_world),
	1743: uint16(177),
	1744: uint16(1),
	1745: uint16(anon_sym_interface),
	1746: uint16(37),
	1747: uint16(1),
	1748: uint16(sym__gate_item),
	1749: uint16(82),
	1750: uint16(1),
	1751: uint16(aux_sym__gate),
	1752: uint16(99),
	1753: uint16(1),
	1754: uint16(sym__statement),
	1755: uint16(219),
	1756: uint16(1),
	1757: uint16(sym_decl_head),
	1758: uint16(26),
	1759: uint16(3),
	1760: uint16(sym_line_comment),
	1761: uint16(sym_block_comment),
	1762: uint16(aux_sym_source_file_repeat1),
	1763: uint16(38),
	1764: uint16(3),
	1765: uint16(sym_unstable_gate),
	1766: uint16(sym_since_gate),
	1767: uint16(sym_deprecated_gate),
	1768: uint16(97),
	1769: uint16(3),
	1770: uint16(sym_toplevel_use_item),
	1771: uint16(sym_world_item),
	1772: uint16(sym_interface_item),
	1773: uint16(100),
	1774: uint16(3),
	1775: uint16(sym__package_items),
	1776: uint16(sym_nested_package_definition),
	1777: uint16(sym_package_decl),
	1778: uint16(17),
	1779: uint16(9),
	1780: uint16(1),
	1781: uint16(anon_sym_AT),
	1782: uint16(11),
	1783: uint16(1),
	1784: uint16(anon_sym_package),
	1785: uint16(13),
	1786: uint16(1),
	1787: uint16(anon_sym_use),
	1788: uint16(15),
	1789: uint16(1),
	1790: uint16(anon_sym_world),
	1791: uint16(17),
	1792: uint16(1),
	1793: uint16(anon_sym_interface),
	1794: uint16(19),
	1795: uint16(1),
	1796: uint16(anon_sym_SLASH_SLASH),
	1797: uint16(21),
	1798: uint16(1),
	1799: uint16(anon_sym_SLASH_STAR),
	1800: uint16(180),
	1801: uint16(1),
	1803: uint16(26),
	1804: uint16(1),
	1805: uint16(aux_sym_source_file_repeat1),
	1806: uint16(37),
	1807: uint16(1),
	1808: uint16(sym__gate_item),
	1809: uint16(82),
	1810: uint16(1),
	1811: uint16(aux_sym__gate),
	1812: uint16(99),
	1813: uint16(1),
	1814: uint16(sym__statement),
	1815: uint16(219),
	1816: uint16(1),
	1817: uint16(sym_decl_head),
	1818: uint16(27),
	1819: uint16(2),
	1820: uint16(sym_line_comment),
	1821: uint16(sym_block_comment),
	1822: uint16(38),
	1823: uint16(3),
	1824: uint16(sym_unstable_gate),
	1825: uint16(sym_since_gate),
	1826: uint16(sym_deprecated_gate),
	1827: uint16(97),
	1828: uint16(3),
	1829: uint16(sym_toplevel_use_item),
	1830: uint16(sym_world_item),
	1831: uint16(sym_interface_item),
	1832: uint16(100),
	1833: uint16(3),
	1834: uint16(sym__package_items),
	1835: uint16(sym_nested_package_definition),
	1836: uint16(sym_package_decl),
	1837: uint16(7),
	1838: uint16(19),
	1839: uint16(1),
	1840: uint16(anon_sym_SLASH_SLASH),
	1841: uint16(21),
	1842: uint16(1),
	1843: uint16(anon_sym_SLASH_STAR),
	1844: uint16(184),
	1845: uint16(1),
	1846: uint16(anon_sym_AT),
	1847: uint16(37),
	1848: uint16(1),
	1849: uint16(sym__gate_item),
	1850: uint16(28),
	1851: uint16(3),
	1852: uint16(sym_line_comment),
	1853: uint16(sym_block_comment),
	1854: uint16(aux_sym__gate),
	1855: uint16(38),
	1856: uint16(3),
	1857: uint16(sym_unstable_gate),
	1858: uint16(sym_since_gate),
	1859: uint16(sym_deprecated_gate),
	1860: uint16(182),
	1861: uint16(13),
	1862: uint16(anon_sym_use),
	1863: uint16(sym_id),
	1864: uint16(anon_sym_world),
	1865: uint16(anon_sym_export),
	1866: uint16(anon_sym_import),
	1867: uint16(anon_sym_interface),
	1868: uint16(anon_sym_include),
	1869: uint16(anon_sym_type),
	1870: uint16(anon_sym_record),
	1871: uint16(anon_sym_flags),
	1872: uint16(anon_sym_variant),
	1873: uint16(anon_sym_enum),
	1874: uint16(anon_sym_resource),
	1875: uint16(4),
	1876: uint16(19),
	1877: uint16(1),
	1878: uint16(anon_sym_SLASH_SLASH),
	1879: uint16(21),
	1880: uint16(1),
	1881: uint16(anon_sym_SLASH_STAR),
	1882: uint16(29),
	1883: uint16(2),
	1884: uint16(sym_line_comment),
	1885: uint16(sym_block_comment),
	1886: uint16(187),
	1887: uint16(16),
	1889: uint16(anon_sym_RBRACE),
	1890: uint16(anon_sym_AT),
	1891: uint16(anon_sym_package),
	1892: uint16(anon_sym_use),
	1893: uint16(anon_sym_world),
	1894: uint16(anon_sym_export),
	1895: uint16(anon_sym_import),
	1896: uint16(anon_sym_interface),
	1897: uint16(anon_sym_include),
	1898: uint16(anon_sym_type),
	1899: uint16(anon_sym_record),
	1900: uint16(anon_sym_flags),
	1901: uint16(anon_sym_variant),
	1902: uint16(anon_sym_enum),
	1903: uint16(anon_sym_resource),
	1904: uint16(4),
	1905: uint16(19),
	1906: uint16(1),
	1907: uint16(anon_sym_SLASH_SLASH),
	1908: uint16(21),
	1909: uint16(1),
	1910: uint16(anon_sym_SLASH_STAR),
	1911: uint16(30),
	1912: uint16(2),
	1913: uint16(sym_line_comment),
	1914: uint16(sym_block_comment),
	1915: uint16(189),
	1916: uint16(16),
	1918: uint16(anon_sym_RBRACE),
	1919: uint16(anon_sym_AT),
	1920: uint16(anon_sym_package),
	1921: uint16(anon_sym_use),
	1922: uint16(anon_sym_world),
	1923: uint16(anon_sym_export),
	1924: uint16(anon_sym_import),
	1925: uint16(anon_sym_interface),
	1926: uint16(anon_sym_include),
	1927: uint16(anon_sym_type),
	1928: uint16(anon_sym_record),
	1929: uint16(anon_sym_flags),
	1930: uint16(anon_sym_variant),
	1931: uint16(anon_sym_enum),
	1932: uint16(anon_sym_resource),
	1933: uint16(8),
	1934: uint16(19),
	1935: uint16(1),
	1936: uint16(anon_sym_SLASH_SLASH),
	1937: uint16(21),
	1938: uint16(1),
	1939: uint16(anon_sym_SLASH_STAR),
	1940: uint16(193),
	1941: uint16(1),
	1942: uint16(anon_sym_LBRACE),
	1943: uint16(197),
	1944: uint16(1),
	1945: uint16(anon_sym_SEMI),
	1946: uint16(54),
	1947: uint16(1),
	1948: uint16(sym__resource_body),
	1949: uint16(195),
	1950: uint16(2),
	1951: uint16(anon_sym_RBRACE),
	1952: uint16(anon_sym_AT),
	1953: uint16(31),
	1954: uint16(2),
	1955: uint16(sym_line_comment),
	1956: uint16(sym_block_comment),
	1957: uint16(191),
	1958: uint16(11),
	1959: uint16(anon_sym_use),
	1960: uint16(sym_id),
	1961: uint16(anon_sym_export),
	1962: uint16(anon_sym_import),
	1963: uint16(anon_sym_include),
	1964: uint16(anon_sym_type),
	1965: uint16(anon_sym_record),
	1966: uint16(anon_sym_flags),
	1967: uint16(anon_sym_variant),
	1968: uint16(anon_sym_enum),
	1969: uint16(anon_sym_resource),
	1970: uint16(12),
	1971: uint16(19),
	1972: uint16(1),
	1973: uint16(anon_sym_SLASH_SLASH),
	1974: uint16(21),
	1975: uint16(1),
	1976: uint16(anon_sym_SLASH_STAR),
	1977: uint16(143),
	1978: uint16(1),
	1979: uint16(sym_id),
	1980: uint16(149),
	1981: uint16(1),
	1982: uint16(anon_sym_type),
	1983: uint16(151),
	1984: uint16(1),
	1985: uint16(anon_sym_record),
	1986: uint16(153),
	1987: uint16(1),
	1988: uint16(anon_sym_flags),
	1989: uint16(155),
	1990: uint16(1),
	1991: uint16(anon_sym_variant),
	1992: uint16(157),
	1993: uint16(1),
	1994: uint16(anon_sym_enum),
	1995: uint16(159),
	1996: uint16(1),
	1997: uint16(anon_sym_resource),
	1998: uint16(32),
	1999: uint16(2),
	2000: uint16(sym_line_comment),
	2001: uint16(sym_block_comment),
	2002: uint16(75),
	2003: uint16(2),
	2004: uint16(sym__typedef_item),
	2005: uint16(sym_func_item),
	2006: uint16(55),
	2007: uint16(6),
	2008: uint16(sym_type_item),
	2009: uint16(sym_record_item),
	2010: uint16(sym_flags_items),
	2011: uint16(sym_variant_items),
	2012: uint16(sym_enum_items),
	2013: uint16(sym_resource_item),
	2014: uint16(14),
	2015: uint16(9),
	2016: uint16(1),
	2017: uint16(anon_sym_AT),
	2018: uint16(13),
	2019: uint16(1),
	2020: uint16(anon_sym_use),
	2021: uint16(15),
	2022: uint16(1),
	2023: uint16(anon_sym_world),
	2024: uint16(17),
	2025: uint16(1),
	2026: uint16(anon_sym_interface),
	2027: uint16(19),
	2028: uint16(1),
	2029: uint16(anon_sym_SLASH_SLASH),
	2030: uint16(21),
	2031: uint16(1),
	2032: uint16(anon_sym_SLASH_STAR),
	2033: uint16(199),
	2034: uint16(1),
	2035: uint16(anon_sym_RBRACE),
	2036: uint16(34),
	2037: uint16(1),
	2038: uint16(aux_sym_nested_package_definition_repeat1),
	2039: uint16(37),
	2040: uint16(1),
	2041: uint16(sym__gate_item),
	2042: uint16(82),
	2043: uint16(1),
	2044: uint16(aux_sym__gate),
	2045: uint16(114),
	2046: uint16(1),
	2047: uint16(sym__package_items),
	2048: uint16(33),
	2049: uint16(2),
	2050: uint16(sym_line_comment),
	2051: uint16(sym_block_comment),
	2052: uint16(38),
	2053: uint16(3),
	2054: uint16(sym_unstable_gate),
	2055: uint16(sym_since_gate),
	2056: uint16(sym_deprecated_gate),
	2057: uint16(97),
	2058: uint16(3),
	2059: uint16(sym_toplevel_use_item),
	2060: uint16(sym_world_item),
	2061: uint16(sym_interface_item),
	2062: uint16(14),
	2063: uint16(9),
	2064: uint16(1),
	2065: uint16(anon_sym_AT),
	2066: uint16(13),
	2067: uint16(1),
	2068: uint16(anon_sym_use),
	2069: uint16(15),
	2070: uint16(1),
	2071: uint16(anon_sym_world),
	2072: uint16(17),
	2073: uint16(1),
	2074: uint16(anon_sym_interface),
	2075: uint16(19),
	2076: uint16(1),
	2077: uint16(anon_sym_SLASH_SLASH),
	2078: uint16(21),
	2079: uint16(1),
	2080: uint16(anon_sym_SLASH_STAR),
	2081: uint16(201),
	2082: uint16(1),
	2083: uint16(anon_sym_RBRACE),
	2084: uint16(36),
	2085: uint16(1),
	2086: uint16(aux_sym_nested_package_definition_repeat1),
	2087: uint16(37),
	2088: uint16(1),
	2089: uint16(sym__gate_item),
	2090: uint16(82),
	2091: uint16(1),
	2092: uint16(aux_sym__gate),
	2093: uint16(114),
	2094: uint16(1),
	2095: uint16(sym__package_items),
	2096: uint16(34),
	2097: uint16(2),
	2098: uint16(sym_line_comment),
	2099: uint16(sym_block_comment),
	2100: uint16(38),
	2101: uint16(3),
	2102: uint16(sym_unstable_gate),
	2103: uint16(sym_since_gate),
	2104: uint16(sym_deprecated_gate),
	2105: uint16(97),
	2106: uint16(3),
	2107: uint16(sym_toplevel_use_item),
	2108: uint16(sym_world_item),
	2109: uint16(sym_interface_item),
	2110: uint16(12),
	2111: uint16(19),
	2112: uint16(1),
	2113: uint16(anon_sym_SLASH_SLASH),
	2114: uint16(21),
	2115: uint16(1),
	2116: uint16(anon_sym_SLASH_STAR),
	2117: uint16(143),
	2118: uint16(1),
	2119: uint16(sym_id),
	2120: uint16(149),
	2121: uint16(1),
	2122: uint16(anon_sym_type),
	2123: uint16(151),
	2124: uint16(1),
	2125: uint16(anon_sym_record),
	2126: uint16(153),
	2127: uint16(1),
	2128: uint16(anon_sym_flags),
	2129: uint16(155),
	2130: uint16(1),
	2131: uint16(anon_sym_variant),
	2132: uint16(157),
	2133: uint16(1),
	2134: uint16(anon_sym_enum),
	2135: uint16(159),
	2136: uint16(1),
	2137: uint16(anon_sym_resource),
	2138: uint16(35),
	2139: uint16(2),
	2140: uint16(sym_line_comment),
	2141: uint16(sym_block_comment),
	2142: uint16(71),
	2143: uint16(2),
	2144: uint16(sym__typedef_item),
	2145: uint16(sym_func_item),
	2146: uint16(55),
	2147: uint16(6),
	2148: uint16(sym_type_item),
	2149: uint16(sym_record_item),
	2150: uint16(sym_flags_items),
	2151: uint16(sym_variant_items),
	2152: uint16(sym_enum_items),
	2153: uint16(sym_resource_item),
	2154: uint16(13),
	2155: uint16(19),
	2156: uint16(1),
	2157: uint16(anon_sym_SLASH_SLASH),
	2158: uint16(21),
	2159: uint16(1),
	2160: uint16(anon_sym_SLASH_STAR),
	2161: uint16(203),
	2162: uint16(1),
	2163: uint16(anon_sym_RBRACE),
	2164: uint16(205),
	2165: uint16(1),
	2166: uint16(anon_sym_AT),
	2167: uint16(208),
	2168: uint16(1),
	2169: uint16(anon_sym_use),
	2170: uint16(211),
	2171: uint16(1),
	2172: uint16(anon_sym_world),
	2173: uint16(214),
	2174: uint16(1),
	2175: uint16(anon_sym_interface),
	2176: uint16(37),
	2177: uint16(1),
	2178: uint16(sym__gate_item),
	2179: uint16(82),
	2180: uint16(1),
	2181: uint16(aux_sym__gate),
	2182: uint16(114),
	2183: uint16(1),
	2184: uint16(sym__package_items),
	2185: uint16(36),
	2186: uint16(3),
	2187: uint16(sym_line_comment),
	2188: uint16(sym_block_comment),
	2189: uint16(aux_sym_nested_package_definition_repeat1),
	2190: uint16(38),
	2191: uint16(3),
	2192: uint16(sym_unstable_gate),
	2193: uint16(sym_since_gate),
	2194: uint16(sym_deprecated_gate),
	2195: uint16(97),
	2196: uint16(3),
	2197: uint16(sym_toplevel_use_item),
	2198: uint16(sym_world_item),
	2199: uint16(sym_interface_item),
	2200: uint16(5),
	2201: uint16(19),
	2202: uint16(1),
	2203: uint16(anon_sym_SLASH_SLASH),
	2204: uint16(21),
	2205: uint16(1),
	2206: uint16(anon_sym_SLASH_STAR),
	2207: uint16(219),
	2208: uint16(1),
	2209: uint16(anon_sym_AT),
	2210: uint16(37),
	2211: uint16(2),
	2212: uint16(sym_line_comment),
	2213: uint16(sym_block_comment),
	2214: uint16(217),
	2215: uint16(13),
	2216: uint16(anon_sym_use),
	2217: uint16(sym_id),
	2218: uint16(anon_sym_world),
	2219: uint16(anon_sym_export),
	2220: uint16(anon_sym_import),
	2221: uint16(anon_sym_interface),
	2222: uint16(anon_sym_include),
	2223: uint16(anon_sym_type),
	2224: uint16(anon_sym_record),
	2225: uint16(anon_sym_flags),
	2226: uint16(anon_sym_variant),
	2227: uint16(anon_sym_enum),
	2228: uint16(anon_sym_resource),
	2229: uint16(5),
	2230: uint16(19),
	2231: uint16(1),
	2232: uint16(anon_sym_SLASH_SLASH),
	2233: uint16(21),
	2234: uint16(1),
	2235: uint16(anon_sym_SLASH_STAR),
	2236: uint16(223),
	2237: uint16(1),
	2238: uint16(anon_sym_AT),
	2239: uint16(38),
	2240: uint16(2),
	2241: uint16(sym_line_comment),
	2242: uint16(sym_block_comment),
	2243: uint16(221),
	2244: uint16(13),
	2245: uint16(anon_sym_use),
	2246: uint16(sym_id),
	2247: uint16(anon_sym_world),
	2248: uint16(anon_sym_export),
	2249: uint16(anon_sym_import),
	2250: uint16(anon_sym_interface),
	2251: uint16(anon_sym_include),
	2252: uint16(anon_sym_type),
	2253: uint16(anon_sym_record),
	2254: uint16(anon_sym_flags),
	2255: uint16(anon_sym_variant),
	2256: uint16(anon_sym_enum),
	2257: uint16(anon_sym_resource),
	2258: uint16(5),
	2259: uint16(19),
	2260: uint16(1),
	2261: uint16(anon_sym_SLASH_SLASH),
	2262: uint16(21),
	2263: uint16(1),
	2264: uint16(anon_sym_SLASH_STAR),
	2265: uint16(227),
	2266: uint16(1),
	2267: uint16(anon_sym_AT),
	2268: uint16(39),
	2269: uint16(2),
	2270: uint16(sym_line_comment),
	2271: uint16(sym_block_comment),
	2272: uint16(225),
	2273: uint16(13),
	2274: uint16(anon_sym_use),
	2275: uint16(sym_id),
	2276: uint16(anon_sym_world),
	2277: uint16(anon_sym_export),
	2278: uint16(anon_sym_import),
	2279: uint16(anon_sym_interface),
	2280: uint16(anon_sym_include),
	2281: uint16(anon_sym_type),
	2282: uint16(anon_sym_record),
	2283: uint16(anon_sym_flags),
	2284: uint16(anon_sym_variant),
	2285: uint16(anon_sym_enum),
	2286: uint16(anon_sym_resource),
	2287: uint16(5),
	2288: uint16(19),
	2289: uint16(1),
	2290: uint16(anon_sym_SLASH_SLASH),
	2291: uint16(21),
	2292: uint16(1),
	2293: uint16(anon_sym_SLASH_STAR),
	2294: uint16(231),
	2295: uint16(1),
	2296: uint16(anon_sym_AT),
	2297: uint16(40),
	2298: uint16(2),
	2299: uint16(sym_line_comment),
	2300: uint16(sym_block_comment),
	2301: uint16(229),
	2302: uint16(13),
	2303: uint16(anon_sym_use),
	2304: uint16(sym_id),
	2305: uint16(anon_sym_world),
	2306: uint16(anon_sym_export),
	2307: uint16(anon_sym_import),
	2308: uint16(anon_sym_interface),
	2309: uint16(anon_sym_include),
	2310: uint16(anon_sym_type),
	2311: uint16(anon_sym_record),
	2312: uint16(anon_sym_flags),
	2313: uint16(anon_sym_variant),
	2314: uint16(anon_sym_enum),
	2315: uint16(anon_sym_resource),
	2316: uint16(5),
	2317: uint16(19),
	2318: uint16(1),
	2319: uint16(anon_sym_SLASH_SLASH),
	2320: uint16(21),
	2321: uint16(1),
	2322: uint16(anon_sym_SLASH_STAR),
	2323: uint16(235),
	2324: uint16(1),
	2325: uint16(anon_sym_AT),
	2326: uint16(41),
	2327: uint16(2),
	2328: uint16(sym_line_comment),
	2329: uint16(sym_block_comment),
	2330: uint16(233),
	2331: uint16(13),
	2332: uint16(anon_sym_use),
	2333: uint16(sym_id),
	2334: uint16(anon_sym_world),
	2335: uint16(anon_sym_export),
	2336: uint16(anon_sym_import),
	2337: uint16(anon_sym_interface),
	2338: uint16(anon_sym_include),
	2339: uint16(anon_sym_type),
	2340: uint16(anon_sym_record),
	2341: uint16(anon_sym_flags),
	2342: uint16(anon_sym_variant),
	2343: uint16(anon_sym_enum),
	2344: uint16(anon_sym_resource),
	2345: uint16(5),
	2346: uint16(19),
	2347: uint16(1),
	2348: uint16(anon_sym_SLASH_SLASH),
	2349: uint16(21),
	2350: uint16(1),
	2351: uint16(anon_sym_SLASH_STAR),
	2352: uint16(239),
	2353: uint16(2),
	2354: uint16(anon_sym_RBRACE),
	2355: uint16(anon_sym_AT),
	2356: uint16(42),
	2357: uint16(2),
	2358: uint16(sym_line_comment),
	2359: uint16(sym_block_comment),
	2360: uint16(237),
	2361: uint16(11),
	2362: uint16(anon_sym_use),
	2363: uint16(sym_id),
	2364: uint16(anon_sym_export),
	2365: uint16(anon_sym_import),
	2366: uint16(anon_sym_include),
	2367: uint16(anon_sym_type),
	2368: uint16(anon_sym_record),
	2369: uint16(anon_sym_flags),
	2370: uint16(anon_sym_variant),
	2371: uint16(anon_sym_enum),
	2372: uint16(anon_sym_resource),
	2373: uint16(5),
	2374: uint16(19),
	2375: uint16(1),
	2376: uint16(anon_sym_SLASH_SLASH),
	2377: uint16(21),
	2378: uint16(1),
	2379: uint16(anon_sym_SLASH_STAR),
	2380: uint16(243),
	2381: uint16(2),
	2382: uint16(anon_sym_RBRACE),
	2383: uint16(anon_sym_AT),
	2384: uint16(43),
	2385: uint16(2),
	2386: uint16(sym_line_comment),
	2387: uint16(sym_block_comment),
	2388: uint16(241),
	2389: uint16(11),
	2390: uint16(anon_sym_use),
	2391: uint16(sym_id),
	2392: uint16(anon_sym_export),
	2393: uint16(anon_sym_import),
	2394: uint16(anon_sym_include),
	2395: uint16(anon_sym_type),
	2396: uint16(anon_sym_record),
	2397: uint16(anon_sym_flags),
	2398: uint16(anon_sym_variant),
	2399: uint16(anon_sym_enum),
	2400: uint16(anon_sym_resource),
	2401: uint16(5),
	2402: uint16(19),
	2403: uint16(1),
	2404: uint16(anon_sym_SLASH_SLASH),
	2405: uint16(21),
	2406: uint16(1),
	2407: uint16(anon_sym_SLASH_STAR),
	2408: uint16(247),
	2409: uint16(2),
	2410: uint16(anon_sym_RBRACE),
	2411: uint16(anon_sym_AT),
	2412: uint16(44),
	2413: uint16(2),
	2414: uint16(sym_line_comment),
	2415: uint16(sym_block_comment),
	2416: uint16(245),
	2417: uint16(11),
	2418: uint16(anon_sym_use),
	2419: uint16(sym_id),
	2420: uint16(anon_sym_export),
	2421: uint16(anon_sym_import),
	2422: uint16(anon_sym_include),
	2423: uint16(anon_sym_type),
	2424: uint16(anon_sym_record),
	2425: uint16(anon_sym_flags),
	2426: uint16(anon_sym_variant),
	2427: uint16(anon_sym_enum),
	2428: uint16(anon_sym_resource),
	2429: uint16(5),
	2430: uint16(19),
	2431: uint16(1),
	2432: uint16(anon_sym_SLASH_SLASH),
	2433: uint16(21),
	2434: uint16(1),
	2435: uint16(anon_sym_SLASH_STAR),
	2436: uint16(251),
	2437: uint16(2),
	2438: uint16(anon_sym_RBRACE),
	2439: uint16(anon_sym_AT),
	2440: uint16(45),
	2441: uint16(2),
	2442: uint16(sym_line_comment),
	2443: uint16(sym_block_comment),
	2444: uint16(249),
	2445: uint16(11),
	2446: uint16(anon_sym_use),
	2447: uint16(sym_id),
	2448: uint16(anon_sym_export),
	2449: uint16(anon_sym_import),
	2450: uint16(anon_sym_include),
	2451: uint16(anon_sym_type),
	2452: uint16(anon_sym_record),
	2453: uint16(anon_sym_flags),
	2454: uint16(anon_sym_variant),
	2455: uint16(anon_sym_enum),
	2456: uint16(anon_sym_resource),
	2457: uint16(5),
	2458: uint16(19),
	2459: uint16(1),
	2460: uint16(anon_sym_SLASH_SLASH),
	2461: uint16(21),
	2462: uint16(1),
	2463: uint16(anon_sym_SLASH_STAR),
	2464: uint16(255),
	2465: uint16(2),
	2466: uint16(anon_sym_RBRACE),
	2467: uint16(anon_sym_AT),
	2468: uint16(46),
	2469: uint16(2),
	2470: uint16(sym_line_comment),
	2471: uint16(sym_block_comment),
	2472: uint16(253),
	2473: uint16(11),
	2474: uint16(anon_sym_use),
	2475: uint16(sym_id),
	2476: uint16(anon_sym_export),
	2477: uint16(anon_sym_import),
	2478: uint16(anon_sym_include),
	2479: uint16(anon_sym_type),
	2480: uint16(anon_sym_record),
	2481: uint16(anon_sym_flags),
	2482: uint16(anon_sym_variant),
	2483: uint16(anon_sym_enum),
	2484: uint16(anon_sym_resource),
	2485: uint16(5),
	2486: uint16(19),
	2487: uint16(1),
	2488: uint16(anon_sym_SLASH_SLASH),
	2489: uint16(21),
	2490: uint16(1),
	2491: uint16(anon_sym_SLASH_STAR),
	2492: uint16(259),
	2493: uint16(2),
	2494: uint16(anon_sym_RBRACE),
	2495: uint16(anon_sym_AT),
	2496: uint16(47),
	2497: uint16(2),
	2498: uint16(sym_line_comment),
	2499: uint16(sym_block_comment),
	2500: uint16(257),
	2501: uint16(11),
	2502: uint16(anon_sym_use),
	2503: uint16(sym_id),
	2504: uint16(anon_sym_export),
	2505: uint16(anon_sym_import),
	2506: uint16(anon_sym_include),
	2507: uint16(anon_sym_type),
	2508: uint16(anon_sym_record),
	2509: uint16(anon_sym_flags),
	2510: uint16(anon_sym_variant),
	2511: uint16(anon_sym_enum),
	2512: uint16(anon_sym_resource),
	2513: uint16(5),
	2514: uint16(19),
	2515: uint16(1),
	2516: uint16(anon_sym_SLASH_SLASH),
	2517: uint16(21),
	2518: uint16(1),
	2519: uint16(anon_sym_SLASH_STAR),
	2520: uint16(263),
	2521: uint16(2),
	2522: uint16(anon_sym_RBRACE),
	2523: uint16(anon_sym_AT),
	2524: uint16(48),
	2525: uint16(2),
	2526: uint16(sym_line_comment),
	2527: uint16(sym_block_comment),
	2528: uint16(261),
	2529: uint16(11),
	2530: uint16(anon_sym_use),
	2531: uint16(sym_id),
	2532: uint16(anon_sym_export),
	2533: uint16(anon_sym_import),
	2534: uint16(anon_sym_include),
	2535: uint16(anon_sym_type),
	2536: uint16(anon_sym_record),
	2537: uint16(anon_sym_flags),
	2538: uint16(anon_sym_variant),
	2539: uint16(anon_sym_enum),
	2540: uint16(anon_sym_resource),
	2541: uint16(5),
	2542: uint16(19),
	2543: uint16(1),
	2544: uint16(anon_sym_SLASH_SLASH),
	2545: uint16(21),
	2546: uint16(1),
	2547: uint16(anon_sym_SLASH_STAR),
	2548: uint16(267),
	2549: uint16(2),
	2550: uint16(anon_sym_RBRACE),
	2551: uint16(anon_sym_AT),
	2552: uint16(49),
	2553: uint16(2),
	2554: uint16(sym_line_comment),
	2555: uint16(sym_block_comment),
	2556: uint16(265),
	2557: uint16(11),
	2558: uint16(anon_sym_use),
	2559: uint16(sym_id),
	2560: uint16(anon_sym_export),
	2561: uint16(anon_sym_import),
	2562: uint16(anon_sym_include),
	2563: uint16(anon_sym_type),
	2564: uint16(anon_sym_record),
	2565: uint16(anon_sym_flags),
	2566: uint16(anon_sym_variant),
	2567: uint16(anon_sym_enum),
	2568: uint16(anon_sym_resource),
	2569: uint16(5),
	2570: uint16(19),
	2571: uint16(1),
	2572: uint16(anon_sym_SLASH_SLASH),
	2573: uint16(21),
	2574: uint16(1),
	2575: uint16(anon_sym_SLASH_STAR),
	2576: uint16(271),
	2577: uint16(2),
	2578: uint16(anon_sym_RBRACE),
	2579: uint16(anon_sym_AT),
	2580: uint16(50),
	2581: uint16(2),
	2582: uint16(sym_line_comment),
	2583: uint16(sym_block_comment),
	2584: uint16(269),
	2585: uint16(11),
	2586: uint16(anon_sym_use),
	2587: uint16(sym_id),
	2588: uint16(anon_sym_export),
	2589: uint16(anon_sym_import),
	2590: uint16(anon_sym_include),
	2591: uint16(anon_sym_type),
	2592: uint16(anon_sym_record),
	2593: uint16(anon_sym_flags),
	2594: uint16(anon_sym_variant),
	2595: uint16(anon_sym_enum),
	2596: uint16(anon_sym_resource),
	2597: uint16(5),
	2598: uint16(19),
	2599: uint16(1),
	2600: uint16(anon_sym_SLASH_SLASH),
	2601: uint16(21),
	2602: uint16(1),
	2603: uint16(anon_sym_SLASH_STAR),
	2604: uint16(275),
	2605: uint16(2),
	2606: uint16(anon_sym_RBRACE),
	2607: uint16(anon_sym_AT),
	2608: uint16(51),
	2609: uint16(2),
	2610: uint16(sym_line_comment),
	2611: uint16(sym_block_comment),
	2612: uint16(273),
	2613: uint16(11),
	2614: uint16(anon_sym_use),
	2615: uint16(sym_id),
	2616: uint16(anon_sym_export),
	2617: uint16(anon_sym_import),
	2618: uint16(anon_sym_include),
	2619: uint16(anon_sym_type),
	2620: uint16(anon_sym_record),
	2621: uint16(anon_sym_flags),
	2622: uint16(anon_sym_variant),
	2623: uint16(anon_sym_enum),
	2624: uint16(anon_sym_resource),
	2625: uint16(5),
	2626: uint16(19),
	2627: uint16(1),
	2628: uint16(anon_sym_SLASH_SLASH),
	2629: uint16(21),
	2630: uint16(1),
	2631: uint16(anon_sym_SLASH_STAR),
	2632: uint16(279),
	2633: uint16(2),
	2634: uint16(anon_sym_RBRACE),
	2635: uint16(anon_sym_AT),
	2636: uint16(52),
	2637: uint16(2),
	2638: uint16(sym_line_comment),
	2639: uint16(sym_block_comment),
	2640: uint16(277),
	2641: uint16(11),
	2642: uint16(anon_sym_use),
	2643: uint16(sym_id),
	2644: uint16(anon_sym_export),
	2645: uint16(anon_sym_import),
	2646: uint16(anon_sym_include),
	2647: uint16(anon_sym_type),
	2648: uint16(anon_sym_record),
	2649: uint16(anon_sym_flags),
	2650: uint16(anon_sym_variant),
	2651: uint16(anon_sym_enum),
	2652: uint16(anon_sym_resource),
	2653: uint16(6),
	2654: uint16(19),
	2655: uint16(1),
	2656: uint16(anon_sym_SLASH_SLASH),
	2657: uint16(21),
	2658: uint16(1),
	2659: uint16(anon_sym_SLASH_STAR),
	2660: uint16(269),
	2661: uint16(1),
	2662: uint16(sym__primitive_ty),
	2663: uint16(270),
	2664: uint16(1),
	2665: uint16(sym_kt),
	2666: uint16(53),
	2667: uint16(2),
	2668: uint16(sym_line_comment),
	2669: uint16(sym_block_comment),
	2670: uint16(281),
	2671: uint16(11),
	2672: uint16(anon_sym_u8),
	2673: uint16(anon_sym_u16),
	2674: uint16(anon_sym_u32),
	2675: uint16(anon_sym_u64),
	2676: uint16(anon_sym_s8),
	2677: uint16(anon_sym_s16),
	2678: uint16(anon_sym_s32),
	2679: uint16(anon_sym_s64),
	2680: uint16(anon_sym_char),
	2681: uint16(anon_sym_bool),
	2682: uint16(anon_sym_string),
	2683: uint16(5),
	2684: uint16(19),
	2685: uint16(1),
	2686: uint16(anon_sym_SLASH_SLASH),
	2687: uint16(21),
	2688: uint16(1),
	2689: uint16(anon_sym_SLASH_STAR),
	2690: uint16(285),
	2691: uint16(2),
	2692: uint16(anon_sym_RBRACE),
	2693: uint16(anon_sym_AT),
	2694: uint16(54),
	2695: uint16(2),
	2696: uint16(sym_line_comment),
	2697: uint16(sym_block_comment),
	2698: uint16(283),
	2699: uint16(11),
	2700: uint16(anon_sym_use),
	2701: uint16(sym_id),
	2702: uint16(anon_sym_export),
	2703: uint16(anon_sym_import),
	2704: uint16(anon_sym_include),
	2705: uint16(anon_sym_type),
	2706: uint16(anon_sym_record),
	2707: uint16(anon_sym_flags),
	2708: uint16(anon_sym_variant),
	2709: uint16(anon_sym_enum),
	2710: uint16(anon_sym_resource),
	2711: uint16(5),
	2712: uint16(19),
	2713: uint16(1),
	2714: uint16(anon_sym_SLASH_SLASH),
	2715: uint16(21),
	2716: uint16(1),
	2717: uint16(anon_sym_SLASH_STAR),
	2718: uint16(289),
	2719: uint16(2),
	2720: uint16(anon_sym_RBRACE),
	2721: uint16(anon_sym_AT),
	2722: uint16(55),
	2723: uint16(2),
	2724: uint16(sym_line_comment),
	2725: uint16(sym_block_comment),
	2726: uint16(287),
	2727: uint16(11),
	2728: uint16(anon_sym_use),
	2729: uint16(sym_id),
	2730: uint16(anon_sym_export),
	2731: uint16(anon_sym_import),
	2732: uint16(anon_sym_include),
	2733: uint16(anon_sym_type),
	2734: uint16(anon_sym_record),
	2735: uint16(anon_sym_flags),
	2736: uint16(anon_sym_variant),
	2737: uint16(anon_sym_enum),
	2738: uint16(anon_sym_resource),
	2739: uint16(5),
	2740: uint16(19),
	2741: uint16(1),
	2742: uint16(anon_sym_SLASH_SLASH),
	2743: uint16(21),
	2744: uint16(1),
	2745: uint16(anon_sym_SLASH_STAR),
	2746: uint16(293),
	2747: uint16(2),
	2748: uint16(anon_sym_RBRACE),
	2749: uint16(anon_sym_AT),
	2750: uint16(56),
	2751: uint16(2),
	2752: uint16(sym_line_comment),
	2753: uint16(sym_block_comment),
	2754: uint16(291),
	2755: uint16(11),
	2756: uint16(anon_sym_use),
	2757: uint16(sym_id),
	2758: uint16(anon_sym_export),
	2759: uint16(anon_sym_import),
	2760: uint16(anon_sym_include),
	2761: uint16(anon_sym_type),
	2762: uint16(anon_sym_record),
	2763: uint16(anon_sym_flags),
	2764: uint16(anon_sym_variant),
	2765: uint16(anon_sym_enum),
	2766: uint16(anon_sym_resource),
	2767: uint16(4),
	2768: uint16(19),
	2769: uint16(1),
	2770: uint16(anon_sym_SLASH_SLASH),
	2771: uint16(21),
	2772: uint16(1),
	2773: uint16(anon_sym_SLASH_STAR),
	2774: uint16(57),
	2775: uint16(2),
	2776: uint16(sym_line_comment),
	2777: uint16(sym_block_comment),
	2778: uint16(295),
	2779: uint16(12),
	2780: uint16(anon_sym_RBRACE),
	2781: uint16(anon_sym_AT),
	2782: uint16(anon_sym_use),
	2783: uint16(anon_sym_export),
	2784: uint16(anon_sym_import),
	2785: uint16(anon_sym_include),
	2786: uint16(anon_sym_type),
	2787: uint16(anon_sym_record),
	2788: uint16(anon_sym_flags),
	2789: uint16(anon_sym_variant),
	2790: uint16(anon_sym_enum),
	2791: uint16(anon_sym_resource),
	2792: uint16(4),
	2793: uint16(19),
	2794: uint16(1),
	2795: uint16(anon_sym_SLASH_SLASH),
	2796: uint16(21),
	2797: uint16(1),
	2798: uint16(anon_sym_SLASH_STAR),
	2799: uint16(58),
	2800: uint16(2),
	2801: uint16(sym_line_comment),
	2802: uint16(sym_block_comment),
	2803: uint16(297),
	2804: uint16(12),
	2805: uint16(anon_sym_RBRACE),
	2806: uint16(anon_sym_AT),
	2807: uint16(anon_sym_use),
	2808: uint16(anon_sym_export),
	2809: uint16(anon_sym_import),
	2810: uint16(anon_sym_include),
	2811: uint16(anon_sym_type),
	2812: uint16(anon_sym_record),
	2813: uint16(anon_sym_flags),
	2814: uint16(anon_sym_variant),
	2815: uint16(anon_sym_enum),
	2816: uint16(anon_sym_resource),
	2817: uint16(4),
	2818: uint16(19),
	2819: uint16(1),
	2820: uint16(anon_sym_SLASH_SLASH),
	2821: uint16(21),
	2822: uint16(1),
	2823: uint16(anon_sym_SLASH_STAR),
	2824: uint16(59),
	2825: uint16(2),
	2826: uint16(sym_line_comment),
	2827: uint16(sym_block_comment),
	2828: uint16(299),
	2829: uint16(12),
	2830: uint16(anon_sym_RBRACE),
	2831: uint16(anon_sym_AT),
	2832: uint16(anon_sym_use),
	2833: uint16(anon_sym_export),
	2834: uint16(anon_sym_import),
	2835: uint16(anon_sym_include),
	2836: uint16(anon_sym_type),
	2837: uint16(anon_sym_record),
	2838: uint16(anon_sym_flags),
	2839: uint16(anon_sym_variant),
	2840: uint16(anon_sym_enum),
	2841: uint16(anon_sym_resource),
	2842: uint16(4),
	2843: uint16(19),
	2844: uint16(1),
	2845: uint16(anon_sym_SLASH_SLASH),
	2846: uint16(21),
	2847: uint16(1),
	2848: uint16(anon_sym_SLASH_STAR),
	2849: uint16(60),
	2850: uint16(2),
	2851: uint16(sym_line_comment),
	2852: uint16(sym_block_comment),
	2853: uint16(301),
	2854: uint16(12),
	2855: uint16(anon_sym_RBRACE),
	2856: uint16(anon_sym_AT),
	2857: uint16(anon_sym_use),
	2858: uint16(anon_sym_export),
	2859: uint16(anon_sym_import),
	2860: uint16(anon_sym_include),
	2861: uint16(anon_sym_type),
	2862: uint16(anon_sym_record),
	2863: uint16(anon_sym_flags),
	2864: uint16(anon_sym_variant),
	2865: uint16(anon_sym_enum),
	2866: uint16(anon_sym_resource),
	2867: uint16(4),
	2868: uint16(19),
	2869: uint16(1),
	2870: uint16(anon_sym_SLASH_SLASH),
	2871: uint16(21),
	2872: uint16(1),
	2873: uint16(anon_sym_SLASH_STAR),
	2874: uint16(61),
	2875: uint16(2),
	2876: uint16(sym_line_comment),
	2877: uint16(sym_block_comment),
	2878: uint16(303),
	2879: uint16(12),
	2880: uint16(anon_sym_RBRACE),
	2881: uint16(anon_sym_AT),
	2882: uint16(anon_sym_use),
	2883: uint16(anon_sym_export),
	2884: uint16(anon_sym_import),
	2885: uint16(anon_sym_include),
	2886: uint16(anon_sym_type),
	2887: uint16(anon_sym_record),
	2888: uint16(anon_sym_flags),
	2889: uint16(anon_sym_variant),
	2890: uint16(anon_sym_enum),
	2891: uint16(anon_sym_resource),
	2892: uint16(4),
	2893: uint16(19),
	2894: uint16(1),
	2895: uint16(anon_sym_SLASH_SLASH),
	2896: uint16(21),
	2897: uint16(1),
	2898: uint16(anon_sym_SLASH_STAR),
	2899: uint16(62),
	2900: uint16(2),
	2901: uint16(sym_line_comment),
	2902: uint16(sym_block_comment),
	2903: uint16(305),
	2904: uint16(12),
	2905: uint16(anon_sym_RBRACE),
	2906: uint16(anon_sym_AT),
	2907: uint16(anon_sym_use),
	2908: uint16(anon_sym_export),
	2909: uint16(anon_sym_import),
	2910: uint16(anon_sym_include),
	2911: uint16(anon_sym_type),
	2912: uint16(anon_sym_record),
	2913: uint16(anon_sym_flags),
	2914: uint16(anon_sym_variant),
	2915: uint16(anon_sym_enum),
	2916: uint16(anon_sym_resource),
	2917: uint16(4),
	2918: uint16(19),
	2919: uint16(1),
	2920: uint16(anon_sym_SLASH_SLASH),
	2921: uint16(21),
	2922: uint16(1),
	2923: uint16(anon_sym_SLASH_STAR),
	2924: uint16(63),
	2925: uint16(2),
	2926: uint16(sym_line_comment),
	2927: uint16(sym_block_comment),
	2928: uint16(307),
	2929: uint16(12),
	2930: uint16(anon_sym_RBRACE),
	2931: uint16(anon_sym_AT),
	2932: uint16(anon_sym_use),
	2933: uint16(anon_sym_export),
	2934: uint16(anon_sym_import),
	2935: uint16(anon_sym_include),
	2936: uint16(anon_sym_type),
	2937: uint16(anon_sym_record),
	2938: uint16(anon_sym_flags),
	2939: uint16(anon_sym_variant),
	2940: uint16(anon_sym_enum),
	2941: uint16(anon_sym_resource),
	2942: uint16(4),
	2943: uint16(19),
	2944: uint16(1),
	2945: uint16(anon_sym_SLASH_SLASH),
	2946: uint16(21),
	2947: uint16(1),
	2948: uint16(anon_sym_SLASH_STAR),
	2949: uint16(64),
	2950: uint16(2),
	2951: uint16(sym_line_comment),
	2952: uint16(sym_block_comment),
	2953: uint16(309),
	2954: uint16(12),
	2955: uint16(anon_sym_RBRACE),
	2956: uint16(anon_sym_AT),
	2957: uint16(anon_sym_use),
	2958: uint16(anon_sym_export),
	2959: uint16(anon_sym_import),
	2960: uint16(anon_sym_include),
	2961: uint16(anon_sym_type),
	2962: uint16(anon_sym_record),
	2963: uint16(anon_sym_flags),
	2964: uint16(anon_sym_variant),
	2965: uint16(anon_sym_enum),
	2966: uint16(anon_sym_resource),
	2967: uint16(4),
	2968: uint16(19),
	2969: uint16(1),
	2970: uint16(anon_sym_SLASH_SLASH),
	2971: uint16(21),
	2972: uint16(1),
	2973: uint16(anon_sym_SLASH_STAR),
	2974: uint16(65),
	2975: uint16(2),
	2976: uint16(sym_line_comment),
	2977: uint16(sym_block_comment),
	2978: uint16(311),
	2979: uint16(12),
	2980: uint16(anon_sym_RBRACE),
	2981: uint16(anon_sym_AT),
	2982: uint16(anon_sym_use),
	2983: uint16(anon_sym_export),
	2984: uint16(anon_sym_import),
	2985: uint16(anon_sym_include),
	2986: uint16(anon_sym_type),
	2987: uint16(anon_sym_record),
	2988: uint16(anon_sym_flags),
	2989: uint16(anon_sym_variant),
	2990: uint16(anon_sym_enum),
	2991: uint16(anon_sym_resource),
	2992: uint16(4),
	2993: uint16(19),
	2994: uint16(1),
	2995: uint16(anon_sym_SLASH_SLASH),
	2996: uint16(21),
	2997: uint16(1),
	2998: uint16(anon_sym_SLASH_STAR),
	2999: uint16(66),
	3000: uint16(2),
	3001: uint16(sym_line_comment),
	3002: uint16(sym_block_comment),
	3003: uint16(313),
	3004: uint16(12),
	3005: uint16(anon_sym_RBRACE),
	3006: uint16(anon_sym_AT),
	3007: uint16(anon_sym_use),
	3008: uint16(anon_sym_export),
	3009: uint16(anon_sym_import),
	3010: uint16(anon_sym_include),
	3011: uint16(anon_sym_type),
	3012: uint16(anon_sym_record),
	3013: uint16(anon_sym_flags),
	3014: uint16(anon_sym_variant),
	3015: uint16(anon_sym_enum),
	3016: uint16(anon_sym_resource),
	3017: uint16(4),
	3018: uint16(19),
	3019: uint16(1),
	3020: uint16(anon_sym_SLASH_SLASH),
	3021: uint16(21),
	3022: uint16(1),
	3023: uint16(anon_sym_SLASH_STAR),
	3024: uint16(67),
	3025: uint16(2),
	3026: uint16(sym_line_comment),
	3027: uint16(sym_block_comment),
	3028: uint16(315),
	3029: uint16(12),
	3030: uint16(anon_sym_RBRACE),
	3031: uint16(anon_sym_AT),
	3032: uint16(anon_sym_use),
	3033: uint16(anon_sym_export),
	3034: uint16(anon_sym_import),
	3035: uint16(anon_sym_include),
	3036: uint16(anon_sym_type),
	3037: uint16(anon_sym_record),
	3038: uint16(anon_sym_flags),
	3039: uint16(anon_sym_variant),
	3040: uint16(anon_sym_enum),
	3041: uint16(anon_sym_resource),
	3042: uint16(4),
	3043: uint16(19),
	3044: uint16(1),
	3045: uint16(anon_sym_SLASH_SLASH),
	3046: uint16(21),
	3047: uint16(1),
	3048: uint16(anon_sym_SLASH_STAR),
	3049: uint16(68),
	3050: uint16(2),
	3051: uint16(sym_line_comment),
	3052: uint16(sym_block_comment),
	3053: uint16(317),
	3054: uint16(12),
	3055: uint16(anon_sym_RBRACE),
	3056: uint16(anon_sym_AT),
	3057: uint16(anon_sym_use),
	3058: uint16(anon_sym_export),
	3059: uint16(anon_sym_import),
	3060: uint16(anon_sym_include),
	3061: uint16(anon_sym_type),
	3062: uint16(anon_sym_record),
	3063: uint16(anon_sym_flags),
	3064: uint16(anon_sym_variant),
	3065: uint16(anon_sym_enum),
	3066: uint16(anon_sym_resource),
	3067: uint16(4),
	3068: uint16(19),
	3069: uint16(1),
	3070: uint16(anon_sym_SLASH_SLASH),
	3071: uint16(21),
	3072: uint16(1),
	3073: uint16(anon_sym_SLASH_STAR),
	3074: uint16(69),
	3075: uint16(2),
	3076: uint16(sym_line_comment),
	3077: uint16(sym_block_comment),
	3078: uint16(319),
	3079: uint16(12),
	3080: uint16(anon_sym_RBRACE),
	3081: uint16(anon_sym_AT),
	3082: uint16(anon_sym_use),
	3083: uint16(anon_sym_export),
	3084: uint16(anon_sym_import),
	3085: uint16(anon_sym_include),
	3086: uint16(anon_sym_type),
	3087: uint16(anon_sym_record),
	3088: uint16(anon_sym_flags),
	3089: uint16(anon_sym_variant),
	3090: uint16(anon_sym_enum),
	3091: uint16(anon_sym_resource),
	3092: uint16(5),
	3093: uint16(19),
	3094: uint16(1),
	3095: uint16(anon_sym_SLASH_SLASH),
	3096: uint16(21),
	3097: uint16(1),
	3098: uint16(anon_sym_SLASH_STAR),
	3099: uint16(323),
	3100: uint16(2),
	3101: uint16(anon_sym_RBRACE),
	3102: uint16(anon_sym_AT),
	3103: uint16(70),
	3104: uint16(2),
	3105: uint16(sym_line_comment),
	3106: uint16(sym_block_comment),
	3107: uint16(321),
	3108: uint16(9),
	3109: uint16(anon_sym_use),
	3110: uint16(sym_id),
	3111: uint16(anon_sym_type),
	3112: uint16(anon_sym_record),
	3113: uint16(anon_sym_flags),
	3114: uint16(anon_sym_variant),
	3115: uint16(anon_sym_enum),
	3116: uint16(anon_sym_resource),
	3117: uint16(anon_sym_constructor),
	3118: uint16(5),
	3119: uint16(19),
	3120: uint16(1),
	3121: uint16(anon_sym_SLASH_SLASH),
	3122: uint16(21),
	3123: uint16(1),
	3124: uint16(anon_sym_SLASH_STAR),
	3125: uint16(327),
	3126: uint16(2),
	3127: uint16(anon_sym_RBRACE),
	3128: uint16(anon_sym_AT),
	3129: uint16(71),
	3130: uint16(2),
	3131: uint16(sym_line_comment),
	3132: uint16(sym_block_comment),
	3133: uint16(325),
	3134: uint16(8),
	3135: uint16(anon_sym_use),
	3136: uint16(sym_id),
	3137: uint16(anon_sym_type),
	3138: uint16(anon_sym_record),
	3139: uint16(anon_sym_flags),
	3140: uint16(anon_sym_variant),
	3141: uint16(anon_sym_enum),
	3142: uint16(anon_sym_resource),
	3143: uint16(4),
	3144: uint16(19),
	3145: uint16(1),
	3146: uint16(anon_sym_SLASH_SLASH),
	3147: uint16(21),
	3148: uint16(1),
	3149: uint16(anon_sym_SLASH_STAR),
	3150: uint16(72),
	3151: uint16(2),
	3152: uint16(sym_line_comment),
	3153: uint16(sym_block_comment),
	3154: uint16(329),
	3155: uint16(10),
	3156: uint16(sym_id),
	3157: uint16(anon_sym_export),
	3158: uint16(anon_sym_import),
	3159: uint16(anon_sym_type),
	3160: uint16(anon_sym_record),
	3161: uint16(anon_sym_flags),
	3162: uint16(anon_sym_variant),
	3163: uint16(anon_sym_enum),
	3164: uint16(anon_sym_resource),
	3165: uint16(anon_sym_constructor),
	3166: uint16(5),
	3167: uint16(19),
	3168: uint16(1),
	3169: uint16(anon_sym_SLASH_SLASH),
	3170: uint16(21),
	3171: uint16(1),
	3172: uint16(anon_sym_SLASH_STAR),
	3173: uint16(333),
	3174: uint16(2),
	3175: uint16(anon_sym_RBRACE),
	3176: uint16(anon_sym_AT),
	3177: uint16(73),
	3178: uint16(2),
	3179: uint16(sym_line_comment),
	3180: uint16(sym_block_comment),
	3181: uint16(331),
	3182: uint16(8),
	3183: uint16(anon_sym_use),
	3184: uint16(sym_id),
	3185: uint16(anon_sym_type),
	3186: uint16(anon_sym_record),
	3187: uint16(anon_sym_flags),
	3188: uint16(anon_sym_variant),
	3189: uint16(anon_sym_enum),
	3190: uint16(anon_sym_resource),
	3191: uint16(5),
	3192: uint16(19),
	3193: uint16(1),
	3194: uint16(anon_sym_SLASH_SLASH),
	3195: uint16(21),
	3196: uint16(1),
	3197: uint16(anon_sym_SLASH_STAR),
	3198: uint16(337),
	3199: uint16(2),
	3200: uint16(anon_sym_RBRACE),
	3201: uint16(anon_sym_AT),
	3202: uint16(74),
	3203: uint16(2),
	3204: uint16(sym_line_comment),
	3205: uint16(sym_block_comment),
	3206: uint16(335),
	3207: uint16(8),
	3208: uint16(anon_sym_use),
	3209: uint16(sym_id),
	3210: uint16(anon_sym_type),
	3211: uint16(anon_sym_record),
	3212: uint16(anon_sym_flags),
	3213: uint16(anon_sym_variant),
	3214: uint16(anon_sym_enum),
	3215: uint16(anon_sym_resource),
	3216: uint16(5),
	3217: uint16(19),
	3218: uint16(1),
	3219: uint16(anon_sym_SLASH_SLASH),
	3220: uint16(21),
	3221: uint16(1),
	3222: uint16(anon_sym_SLASH_STAR),
	3223: uint16(341),
	3224: uint16(2),
	3225: uint16(anon_sym_RBRACE),
	3226: uint16(anon_sym_AT),
	3227: uint16(75),
	3228: uint16(2),
	3229: uint16(sym_line_comment),
	3230: uint16(sym_block_comment),
	3231: uint16(339),
	3232: uint16(8),
	3233: uint16(anon_sym_use),
	3234: uint16(sym_id),
	3235: uint16(anon_sym_type),
	3236: uint16(anon_sym_record),
	3237: uint16(anon_sym_flags),
	3238: uint16(anon_sym_variant),
	3239: uint16(anon_sym_enum),
	3240: uint16(anon_sym_resource),
	3241: uint16(11),
	3242: uint16(19),
	3243: uint16(1),
	3244: uint16(anon_sym_SLASH_SLASH),
	3245: uint16(21),
	3246: uint16(1),
	3247: uint16(anon_sym_SLASH_STAR),
	3248: uint16(343),
	3249: uint16(1),
	3250: uint16(sym_id),
	3251: uint16(345),
	3252: uint16(1),
	3253: uint16(anon_sym_interface),
	3254: uint16(347),
	3255: uint16(1),
	3256: uint16(anon_sym_async),
	3257: uint16(349),
	3258: uint16(1),
	3259: uint16(anon_sym_func),
	3260: uint16(67),
	3261: uint16(1),
	3262: uint16(sym_extern_type),
	3263: uint16(150),
	3264: uint16(1),
	3265: uint16(aux_sym_decl_head_repeat1),
	3266: uint16(281),
	3267: uint16(1),
	3268: uint16(sym__uri_head),
	3269: uint16(76),
	3270: uint16(2),
	3271: uint16(sym_line_comment),
	3272: uint16(sym_block_comment),
	3273: uint16(277),
	3274: uint16(2),
	3275: uint16(sym_use_path),
	3276: uint16(sym_func_type),
	3277: uint16(11),
	3278: uint16(19),
	3279: uint16(1),
	3280: uint16(anon_sym_SLASH_SLASH),
	3281: uint16(21),
	3282: uint16(1),
	3283: uint16(anon_sym_SLASH_STAR),
	3284: uint16(345),
	3285: uint16(1),
	3286: uint16(anon_sym_interface),
	3287: uint16(347),
	3288: uint16(1),
	3289: uint16(anon_sym_async),
	3290: uint16(349),
	3291: uint16(1),
	3292: uint16(anon_sym_func),
	3293: uint16(351),
	3294: uint16(1),
	3295: uint16(sym_id),
	3296: uint16(63),
	3297: uint16(1),
	3298: uint16(sym_extern_type),
	3299: uint16(150),
	3300: uint16(1),
	3301: uint16(aux_sym_decl_head_repeat1),
	3302: uint16(281),
	3303: uint16(1),
	3304: uint16(sym__uri_head),
	3305: uint16(77),
	3306: uint16(2),
	3307: uint16(sym_line_comment),
	3308: uint16(sym_block_comment),
	3309: uint16(277),
	3310: uint16(2),
	3311: uint16(sym_use_path),
	3312: uint16(sym_func_type),
	3313: uint16(11),
	3314: uint16(19),
	3315: uint16(1),
	3316: uint16(anon_sym_SLASH_SLASH),
	3317: uint16(21),
	3318: uint16(1),
	3319: uint16(anon_sym_SLASH_STAR),
	3320: uint16(343),
	3321: uint16(1),
	3322: uint16(sym_id),
	3323: uint16(345),
	3324: uint16(1),
	3325: uint16(anon_sym_interface),
	3326: uint16(347),
	3327: uint16(1),
	3328: uint16(anon_sym_async),
	3329: uint16(349),
	3330: uint16(1),
	3331: uint16(anon_sym_func),
	3332: uint16(68),
	3333: uint16(1),
	3334: uint16(sym_extern_type),
	3335: uint16(150),
	3336: uint16(1),
	3337: uint16(aux_sym_decl_head_repeat1),
	3338: uint16(281),
	3339: uint16(1),
	3340: uint16(sym__uri_head),
	3341: uint16(78),
	3342: uint16(2),
	3343: uint16(sym_line_comment),
	3344: uint16(sym_block_comment),
	3345: uint16(277),
	3346: uint16(2),
	3347: uint16(sym_use_path),
	3348: uint16(sym_func_type),
	3349: uint16(9),
	3350: uint16(19),
	3351: uint16(1),
	3352: uint16(anon_sym_SLASH_SLASH),
	3353: uint16(21),
	3354: uint16(1),
	3355: uint16(anon_sym_SLASH_STAR),
	3356: uint16(354),
	3357: uint16(1),
	3358: uint16(anon_sym_SLASH),
	3359: uint16(356),
	3360: uint16(1),
	3361: uint16(anon_sym_AT),
	3362: uint16(81),
	3363: uint16(1),
	3364: uint16(aux_sym_decl_head_repeat2),
	3365: uint16(88),
	3366: uint16(1),
	3367: uint16(sym__uri_tail),
	3368: uint16(131),
	3369: uint16(1),
	3370: uint16(sym__version),
	3371: uint16(79),
	3372: uint16(2),
	3373: uint16(sym_line_comment),
	3374: uint16(sym_block_comment),
	3375: uint16(358),
	3376: uint16(4),
	3377: uint16(anon_sym_SEMI),
	3378: uint16(anon_sym_as),
	3379: uint16(anon_sym_with),
	3380: uint16(anon_sym_DOT),
	3381: uint16(11),
	3382: uint16(19),
	3383: uint16(1),
	3384: uint16(anon_sym_SLASH_SLASH),
	3385: uint16(21),
	3386: uint16(1),
	3387: uint16(anon_sym_SLASH_STAR),
	3388: uint16(345),
	3389: uint16(1),
	3390: uint16(anon_sym_interface),
	3391: uint16(347),
	3392: uint16(1),
	3393: uint16(anon_sym_async),
	3394: uint16(349),
	3395: uint16(1),
	3396: uint16(anon_sym_func),
	3397: uint16(351),
	3398: uint16(1),
	3399: uint16(sym_id),
	3400: uint16(62),
	3401: uint16(1),
	3402: uint16(sym_extern_type),
	3403: uint16(150),
	3404: uint16(1),
	3405: uint16(aux_sym_decl_head_repeat1),
	3406: uint16(281),
	3407: uint16(1),
	3408: uint16(sym__uri_head),
	3409: uint16(80),
	3410: uint16(2),
	3411: uint16(sym_line_comment),
	3412: uint16(sym_block_comment),
	3413: uint16(277),
	3414: uint16(2),
	3415: uint16(sym_use_path),
	3416: uint16(sym_func_type),
	3417: uint16(6),
	3418: uint16(19),
	3419: uint16(1),
	3420: uint16(anon_sym_SLASH_SLASH),
	3421: uint16(21),
	3422: uint16(1),
	3423: uint16(anon_sym_SLASH_STAR),
	3424: uint16(362),
	3425: uint16(1),
	3426: uint16(anon_sym_SLASH),
	3427: uint16(88),
	3428: uint16(1),
	3429: uint16(sym__uri_tail),
	3430: uint16(81),
	3431: uint16(3),
	3432: uint16(sym_line_comment),
	3433: uint16(sym_block_comment),
	3434: uint16(aux_sym_decl_head_repeat2),
	3435: uint16(360),
	3436: uint16(6),
	3437: uint16(anon_sym_LBRACE),
	3438: uint16(anon_sym_AT),
	3439: uint16(anon_sym_SEMI),
	3440: uint16(anon_sym_as),
	3441: uint16(anon_sym_with),
	3442: uint16(anon_sym_DOT),
	3443: uint16(9),
	3444: uint16(9),
	3445: uint16(1),
	3446: uint16(anon_sym_AT),
	3447: uint16(19),
	3448: uint16(1),
	3449: uint16(anon_sym_SLASH_SLASH),
	3450: uint16(21),
	3451: uint16(1),
	3452: uint16(anon_sym_SLASH_STAR),
	3453: uint16(365),
	3454: uint16(1),
	3455: uint16(anon_sym_world),
	3456: uint16(367),
	3457: uint16(1),
	3458: uint16(anon_sym_interface),
	3459: uint16(28),
	3460: uint16(1),
	3461: uint16(aux_sym__gate),
	3462: uint16(37),
	3463: uint16(1),
	3464: uint16(sym__gate_item),
	3465: uint16(82),
	3466: uint16(2),
	3467: uint16(sym_line_comment),
	3468: uint16(sym_block_comment),
	3469: uint16(38),
	3470: uint16(3),
	3471: uint16(sym_unstable_gate),
	3472: uint16(sym_since_gate),
	3473: uint16(sym_deprecated_gate),
	3474: uint16(11),
	3475: uint16(19),
	3476: uint16(1),
	3477: uint16(anon_sym_SLASH_SLASH),
	3478: uint16(21),
	3479: uint16(1),
	3480: uint16(anon_sym_SLASH_STAR),
	3481: uint16(369),
	3482: uint16(1),
	3483: uint16(sym_id),
	3484: uint16(371),
	3485: uint16(1),
	3486: uint16(anon_sym_RBRACE),
	3487: uint16(373),
	3488: uint16(1),
	3489: uint16(anon_sym_AT),
	3490: uint16(375),
	3491: uint16(1),
	3492: uint16(anon_sym_constructor),
	3493: uint16(85),
	3494: uint16(1),
	3495: uint16(aux_sym__resource_body_repeat1),
	3496: uint16(128),
	3497: uint16(1),
	3498: uint16(sym_resource_method),
	3499: uint16(133),
	3500: uint16(1),
	3501: uint16(sym_external_id),
	3502: uint16(141),
	3503: uint16(1),
	3504: uint16(sym_func_item),
	3505: uint16(83),
	3506: uint16(2),
	3507: uint16(sym_line_comment),
	3508: uint16(sym_block_comment),
	3509: uint16(10),
	3510: uint16(19),
	3511: uint16(1),
	3512: uint16(anon_sym_SLASH_SLASH),
	3513: uint16(21),
	3514: uint16(1),
	3515: uint16(anon_sym_SLASH_STAR),
	3516: uint16(354),
	3517: uint16(1),
	3518: uint16(anon_sym_SLASH),
	3519: uint16(356),
	3520: uint16(1),
	3521: uint16(anon_sym_AT),
	3522: uint16(379),
	3523: uint16(1),
	3524: uint16(anon_sym_COLON),
	3525: uint16(87),
	3526: uint16(1),
	3527: uint16(aux_sym_decl_head_repeat2),
	3528: uint16(88),
	3529: uint16(1),
	3530: uint16(sym__uri_tail),
	3531: uint16(242),
	3532: uint16(1),
	3533: uint16(sym__version),
	3534: uint16(377),
	3535: uint16(2),
	3536: uint16(anon_sym_LBRACE),
	3537: uint16(anon_sym_SEMI),
	3538: uint16(84),
	3539: uint16(2),
	3540: uint16(sym_line_comment),
	3541: uint16(sym_block_comment),
	3542: uint16(10),
	3543: uint16(19),
	3544: uint16(1),
	3545: uint16(anon_sym_SLASH_SLASH),
	3546: uint16(21),
	3547: uint16(1),
	3548: uint16(anon_sym_SLASH_STAR),
	3549: uint16(381),
	3550: uint16(1),
	3551: uint16(sym_id),
	3552: uint16(384),
	3553: uint16(1),
	3554: uint16(anon_sym_RBRACE),
	3555: uint16(386),
	3556: uint16(1),
	3557: uint16(anon_sym_AT),
	3558: uint16(389),
	3559: uint16(1),
	3560: uint16(anon_sym_constructor),
	3561: uint16(128),
	3562: uint16(1),
	3563: uint16(sym_resource_method),
	3564: uint16(133),
	3565: uint16(1),
	3566: uint16(sym_external_id),
	3567: uint16(141),
	3568: uint16(1),
	3569: uint16(sym_func_item),
	3570: uint16(85),
	3571: uint16(3),
	3572: uint16(sym_line_comment),
	3573: uint16(sym_block_comment),
	3574: uint16(aux_sym__resource_body_repeat1),
	3575: uint16(11),
	3576: uint16(19),
	3577: uint16(1),
	3578: uint16(anon_sym_SLASH_SLASH),
	3579: uint16(21),
	3580: uint16(1),
	3581: uint16(anon_sym_SLASH_STAR),
	3582: uint16(369),
	3583: uint16(1),
	3584: uint16(sym_id),
	3585: uint16(373),
	3586: uint16(1),
	3587: uint16(anon_sym_AT),
	3588: uint16(375),
	3589: uint16(1),
	3590: uint16(anon_sym_constructor),
	3591: uint16(392),
	3592: uint16(1),
	3593: uint16(anon_sym_RBRACE),
	3594: uint16(83),
	3595: uint16(1),
	3596: uint16(aux_sym__resource_body_repeat1),
	3597: uint16(128),
	3598: uint16(1),
	3599: uint16(sym_resource_method),
	3600: uint16(133),
	3601: uint16(1),
	3602: uint16(sym_external_id),
	3603: uint16(141),
	3604: uint16(1),
	3605: uint16(sym_func_item),
	3606: uint16(86),
	3607: uint16(2),
	3608: uint16(sym_line_comment),
	3609: uint16(sym_block_comment),
	3610: uint16(9),
	3611: uint16(19),
	3612: uint16(1),
	3613: uint16(anon_sym_SLASH_SLASH),
	3614: uint16(21),
	3615: uint16(1),
	3616: uint16(anon_sym_SLASH_STAR),
	3617: uint16(354),
	3618: uint16(1),
	3619: uint16(anon_sym_SLASH),
	3620: uint16(356),
	3621: uint16(1),
	3622: uint16(anon_sym_AT),
	3623: uint16(81),
	3624: uint16(1),
	3625: uint16(aux_sym_decl_head_repeat2),
	3626: uint16(88),
	3627: uint16(1),
	3628: uint16(sym__uri_tail),
	3629: uint16(220),
	3630: uint16(1),
	3631: uint16(sym__version),
	3632: uint16(394),
	3633: uint16(2),
	3634: uint16(anon_sym_LBRACE),
	3635: uint16(anon_sym_SEMI),
	3636: uint16(87),
	3637: uint16(2),
	3638: uint16(sym_line_comment),
	3639: uint16(sym_block_comment),
	3640: uint16(5),
	3641: uint16(19),
	3642: uint16(1),
	3643: uint16(anon_sym_SLASH_SLASH),
	3644: uint16(21),
	3645: uint16(1),
	3646: uint16(anon_sym_SLASH_STAR),
	3647: uint16(398),
	3648: uint16(1),
	3649: uint16(anon_sym_SLASH),
	3650: uint16(88),
	3651: uint16(2),
	3652: uint16(sym_line_comment),
	3653: uint16(sym_block_comment),
	3654: uint16(396),
	3655: uint16(6),
	3656: uint16(anon_sym_LBRACE),
	3657: uint16(anon_sym_AT),
	3658: uint16(anon_sym_SEMI),
	3659: uint16(anon_sym_as),
	3660: uint16(anon_sym_with),
	3661: uint16(anon_sym_DOT),
	3662: uint16(4),
	3663: uint16(19),
	3664: uint16(1),
	3665: uint16(anon_sym_SLASH_SLASH),
	3666: uint16(21),
	3667: uint16(1),
	3668: uint16(anon_sym_SLASH_STAR),
	3669: uint16(89),
	3670: uint16(2),
	3671: uint16(sym_line_comment),
	3672: uint16(sym_block_comment),
	3673: uint16(400),
	3674: uint16(7),
	3676: uint16(anon_sym_RBRACE),
	3677: uint16(anon_sym_AT),
	3678: uint16(anon_sym_package),
	3679: uint16(anon_sym_use),
	3680: uint16(anon_sym_world),
	3681: uint16(anon_sym_interface),
	3682: uint16(5),
	3683: uint16(19),
	3684: uint16(1),
	3685: uint16(anon_sym_SLASH_SLASH),
	3686: uint16(21),
	3687: uint16(1),
	3688: uint16(anon_sym_SLASH_STAR),
	3689: uint16(404),
	3690: uint16(1),
	3691: uint16(anon_sym_SLASH),
	3692: uint16(90),
	3693: uint16(2),
	3694: uint16(sym_line_comment),
	3695: uint16(sym_block_comment),
	3696: uint16(402),
	3697: uint16(6),
	3698: uint16(anon_sym_LBRACE),
	3699: uint16(anon_sym_AT),
	3700: uint16(anon_sym_SEMI),
	3701: uint16(anon_sym_as),
	3702: uint16(anon_sym_with),
	3703: uint16(anon_sym_DOT),
	3704: uint16(4),
	3705: uint16(19),
	3706: uint16(1),
	3707: uint16(anon_sym_SLASH_SLASH),
	3708: uint16(21),
	3709: uint16(1),
	3710: uint16(anon_sym_SLASH_STAR),
	3711: uint16(91),
	3712: uint16(2),
	3713: uint16(sym_line_comment),
	3714: uint16(sym_block_comment),
	3715: uint16(406),
	3716: uint16(7),
	3718: uint16(anon_sym_RBRACE),
	3719: uint16(anon_sym_AT),
	3720: uint16(anon_sym_package),
	3721: uint16(anon_sym_use),
	3722: uint16(anon_sym_world),
	3723: uint16(anon_sym_interface),
	3724: uint16(4),
	3725: uint16(19),
	3726: uint16(1),
	3727: uint16(anon_sym_SLASH_SLASH),
	3728: uint16(21),
	3729: uint16(1),
	3730: uint16(anon_sym_SLASH_STAR),
	3731: uint16(92),
	3732: uint16(2),
	3733: uint16(sym_line_comment),
	3734: uint16(sym_block_comment),
	3735: uint16(408),
	3736: uint16(7),
	3738: uint16(anon_sym_RBRACE),
	3739: uint16(anon_sym_AT),
	3740: uint16(anon_sym_package),
	3741: uint16(anon_sym_use),
	3742: uint16(anon_sym_world),
	3743: uint16(anon_sym_interface),
	3744: uint16(4),
	3745: uint16(19),
	3746: uint16(1),
	3747: uint16(anon_sym_SLASH_SLASH),
	3748: uint16(21),
	3749: uint16(1),
	3750: uint16(anon_sym_SLASH_STAR),
	3751: uint16(93),
	3752: uint16(2),
	3753: uint16(sym_line_comment),
	3754: uint16(sym_block_comment),
	3755: uint16(410),
	3756: uint16(7),
	3758: uint16(anon_sym_RBRACE),
	3759: uint16(anon_sym_AT),
	3760: uint16(anon_sym_package),
	3761: uint16(anon_sym_use),
	3762: uint16(anon_sym_world),
	3763: uint16(anon_sym_interface),
	3764: uint16(4),
	3765: uint16(19),
	3766: uint16(1),
	3767: uint16(anon_sym_SLASH_SLASH),
	3768: uint16(21),
	3769: uint16(1),
	3770: uint16(anon_sym_SLASH_STAR),
	3771: uint16(94),
	3772: uint16(2),
	3773: uint16(sym_line_comment),
	3774: uint16(sym_block_comment),
	3775: uint16(412),
	3776: uint16(7),
	3778: uint16(anon_sym_RBRACE),
	3779: uint16(anon_sym_AT),
	3780: uint16(anon_sym_package),
	3781: uint16(anon_sym_use),
	3782: uint16(anon_sym_world),
	3783: uint16(anon_sym_interface),
	3784: uint16(4),
	3785: uint16(19),
	3786: uint16(1),
	3787: uint16(anon_sym_SLASH_SLASH),
	3788: uint16(21),
	3789: uint16(1),
	3790: uint16(anon_sym_SLASH_STAR),
	3791: uint16(95),
	3792: uint16(2),
	3793: uint16(sym_line_comment),
	3794: uint16(sym_block_comment),
	3795: uint16(414),
	3796: uint16(7),
	3798: uint16(anon_sym_RBRACE),
	3799: uint16(anon_sym_AT),
	3800: uint16(anon_sym_package),
	3801: uint16(anon_sym_use),
	3802: uint16(anon_sym_world),
	3803: uint16(anon_sym_interface),
	3804: uint16(4),
	3805: uint16(19),
	3806: uint16(1),
	3807: uint16(anon_sym_SLASH_SLASH),
	3808: uint16(21),
	3809: uint16(1),
	3810: uint16(anon_sym_SLASH_STAR),
	3811: uint16(96),
	3812: uint16(2),
	3813: uint16(sym_line_comment),
	3814: uint16(sym_block_comment),
	3815: uint16(416),
	3816: uint16(7),
	3818: uint16(anon_sym_RBRACE),
	3819: uint16(anon_sym_AT),
	3820: uint16(anon_sym_package),
	3821: uint16(anon_sym_use),
	3822: uint16(anon_sym_world),
	3823: uint16(anon_sym_interface),
	3824: uint16(4),
	3825: uint16(19),
	3826: uint16(1),
	3827: uint16(anon_sym_SLASH_SLASH),
	3828: uint16(21),
	3829: uint16(1),
	3830: uint16(anon_sym_SLASH_STAR),
	3831: uint16(97),
	3832: uint16(2),
	3833: uint16(sym_line_comment),
	3834: uint16(sym_block_comment),
	3835: uint16(418),
	3836: uint16(7),
	3838: uint16(anon_sym_RBRACE),
	3839: uint16(anon_sym_AT),
	3840: uint16(anon_sym_package),
	3841: uint16(anon_sym_use),
	3842: uint16(anon_sym_world),
	3843: uint16(anon_sym_interface),
	3844: uint16(4),
	3845: uint16(19),
	3846: uint16(1),
	3847: uint16(anon_sym_SLASH_SLASH),
	3848: uint16(21),
	3849: uint16(1),
	3850: uint16(anon_sym_SLASH_STAR),
	3851: uint16(98),
	3852: uint16(2),
	3853: uint16(sym_line_comment),
	3854: uint16(sym_block_comment),
	3855: uint16(420),
	3856: uint16(7),
	3858: uint16(anon_sym_RBRACE),
	3859: uint16(anon_sym_AT),
	3860: uint16(anon_sym_package),
	3861: uint16(anon_sym_use),
	3862: uint16(anon_sym_world),
	3863: uint16(anon_sym_interface),
	3864: uint16(4),
	3865: uint16(19),
	3866: uint16(1),
	3867: uint16(anon_sym_SLASH_SLASH),
	3868: uint16(21),
	3869: uint16(1),
	3870: uint16(anon_sym_SLASH_STAR),
	3871: uint16(99),
	3872: uint16(2),
	3873: uint16(sym_line_comment),
	3874: uint16(sym_block_comment),
	3875: uint16(422),
	3876: uint16(6),
	3878: uint16(anon_sym_AT),
	3879: uint16(anon_sym_package),
	3880: uint16(anon_sym_use),
	3881: uint16(anon_sym_world),
	3882: uint16(anon_sym_interface),
	3883: uint16(4),
	3884: uint16(19),
	3885: uint16(1),
	3886: uint16(anon_sym_SLASH_SLASH),
	3887: uint16(21),
	3888: uint16(1),
	3889: uint16(anon_sym_SLASH_STAR),
	3890: uint16(100),
	3891: uint16(2),
	3892: uint16(sym_line_comment),
	3893: uint16(sym_block_comment),
	3894: uint16(424),
	3895: uint16(6),
	3897: uint16(anon_sym_AT),
	3898: uint16(anon_sym_package),
	3899: uint16(anon_sym_use),
	3900: uint16(anon_sym_world),
	3901: uint16(anon_sym_interface),
	3902: uint16(5),
	3903: uint16(19),
	3904: uint16(1),
	3905: uint16(anon_sym_SLASH_SLASH),
	3906: uint16(21),
	3907: uint16(1),
	3908: uint16(anon_sym_SLASH_STAR),
	3909: uint16(428),
	3910: uint16(1),
	3911: uint16(anon_sym_LT),
	3912: uint16(101),
	3913: uint16(2),
	3914: uint16(sym_line_comment),
	3915: uint16(sym_block_comment),
	3916: uint16(426),
	3917: uint16(5),
	3918: uint16(anon_sym_RBRACE),
	3919: uint16(anon_sym_SEMI),
	3920: uint16(anon_sym_COMMA),
	3921: uint16(anon_sym_RPAREN),
	3922: uint16(anon_sym_GT),
	3923: uint16(5),
	3924: uint16(19),
	3925: uint16(1),
	3926: uint16(anon_sym_SLASH_SLASH),
	3927: uint16(21),
	3928: uint16(1),
	3929: uint16(anon_sym_SLASH_STAR),
	3930: uint16(432),
	3931: uint16(1),
	3932: uint16(anon_sym_LT),
	3933: uint16(102),
	3934: uint16(2),
	3935: uint16(sym_line_comment),
	3936: uint16(sym_block_comment),
	3937: uint16(430),
	3938: uint16(5),
	3939: uint16(anon_sym_RBRACE),
	3940: uint16(anon_sym_SEMI),
	3941: uint16(anon_sym_COMMA),
	3942: uint16(anon_sym_RPAREN),
	3943: uint16(anon_sym_GT),
	3944: uint16(5),
	3945: uint16(19),
	3946: uint16(1),
	3947: uint16(anon_sym_SLASH_SLASH),
	3948: uint16(21),
	3949: uint16(1),
	3950: uint16(anon_sym_SLASH_STAR),
	3951: uint16(436),
	3952: uint16(1),
	3953: uint16(anon_sym_LT),
	3954: uint16(103),
	3955: uint16(2),
	3956: uint16(sym_line_comment),
	3957: uint16(sym_block_comment),
	3958: uint16(434),
	3959: uint16(5),
	3960: uint16(anon_sym_RBRACE),
	3961: uint16(anon_sym_SEMI),
	3962: uint16(anon_sym_COMMA),
	3963: uint16(anon_sym_RPAREN),
	3964: uint16(anon_sym_GT),
	3965: uint16(4),
	3966: uint16(19),
	3967: uint16(1),
	3968: uint16(anon_sym_SLASH_SLASH),
	3969: uint16(21),
	3970: uint16(1),
	3971: uint16(anon_sym_SLASH_STAR),
	3972: uint16(104),
	3973: uint16(2),
	3974: uint16(sym_line_comment),
	3975: uint16(sym_block_comment),
	3976: uint16(438),
	3977: uint16(6),
	3979: uint16(anon_sym_AT),
	3980: uint16(anon_sym_package),
	3981: uint16(anon_sym_use),
	3982: uint16(anon_sym_world),
	3983: uint16(anon_sym_interface),
	3984: uint16(4),
	3985: uint16(19),
	3986: uint16(1),
	3987: uint16(anon_sym_SLASH_SLASH),
	3988: uint16(21),
	3989: uint16(1),
	3990: uint16(anon_sym_SLASH_STAR),
	3991: uint16(105),
	3992: uint16(2),
	3993: uint16(sym_line_comment),
	3994: uint16(sym_block_comment),
	3995: uint16(440),
	3996: uint16(6),
	3998: uint16(anon_sym_AT),
	3999: uint16(anon_sym_package),
	4000: uint16(anon_sym_use),
	4001: uint16(anon_sym_world),
	4002: uint16(anon_sym_interface),
	4003: uint16(4),
	4004: uint16(19),
	4005: uint16(1),
	4006: uint16(anon_sym_SLASH_SLASH),
	4007: uint16(21),
	4008: uint16(1),
	4009: uint16(anon_sym_SLASH_STAR),
	4010: uint16(106),
	4011: uint16(2),
	4012: uint16(sym_line_comment),
	4013: uint16(sym_block_comment),
	4014: uint16(442),
	4015: uint16(6),
	4017: uint16(anon_sym_AT),
	4018: uint16(anon_sym_package),
	4019: uint16(anon_sym_use),
	4020: uint16(anon_sym_world),
	4021: uint16(anon_sym_interface),
	4022: uint16(4),
	4023: uint16(19),
	4024: uint16(1),
	4025: uint16(anon_sym_SLASH_SLASH),
	4026: uint16(21),
	4027: uint16(1),
	4028: uint16(anon_sym_SLASH_STAR),
	4029: uint16(107),
	4030: uint16(2),
	4031: uint16(sym_line_comment),
	4032: uint16(sym_block_comment),
	4033: uint16(444),
	4034: uint16(5),
	4035: uint16(anon_sym_RBRACE),
	4036: uint16(anon_sym_SEMI),
	4037: uint16(anon_sym_COMMA),
	4038: uint16(anon_sym_RPAREN),
	4039: uint16(anon_sym_GT),
	4040: uint16(4),
	4041: uint16(19),
	4042: uint16(1),
	4043: uint16(anon_sym_SLASH_SLASH),
	4044: uint16(21),
	4045: uint16(1),
	4046: uint16(anon_sym_SLASH_STAR),
	4047: uint16(108),
	4048: uint16(2),
	4049: uint16(sym_line_comment),
	4050: uint16(sym_block_comment),
	4051: uint16(446),
	4052: uint16(5),
	4053: uint16(anon_sym_RBRACE),
	4054: uint16(anon_sym_SEMI),
	4055: uint16(anon_sym_COMMA),
	4056: uint16(anon_sym_RPAREN),
	4057: uint16(anon_sym_GT),
	4058: uint16(5),
	4059: uint16(19),
	4060: uint16(1),
	4061: uint16(anon_sym_SLASH_SLASH),
	4062: uint16(21),
	4063: uint16(1),
	4064: uint16(anon_sym_SLASH_STAR),
	4065: uint16(379),
	4066: uint16(1),
	4067: uint16(anon_sym_COLON),
	4068: uint16(109),
	4069: uint16(2),
	4070: uint16(sym_line_comment),
	4071: uint16(sym_block_comment),
	4072: uint16(448),
	4073: uint16(4),
	4074: uint16(anon_sym_SEMI),
	4075: uint16(anon_sym_as),
	4076: uint16(anon_sym_with),
	4077: uint16(anon_sym_DOT),
	4078: uint16(4),
	4079: uint16(19),
	4080: uint16(1),
	4081: uint16(anon_sym_SLASH_SLASH),
	4082: uint16(21),
	4083: uint16(1),
	4084: uint16(anon_sym_SLASH_STAR),
	4085: uint16(110),
	4086: uint16(2),
	4087: uint16(sym_line_comment),
	4088: uint16(sym_block_comment),
	4089: uint16(450),
	4090: uint16(5),
	4091: uint16(anon_sym_LBRACE),
	4092: uint16(anon_sym_SEMI),
	4093: uint16(anon_sym_as),
	4094: uint16(anon_sym_with),
	4095: uint16(anon_sym_DOT),
	4096: uint16(4),
	4097: uint16(19),
	4098: uint16(1),
	4099: uint16(anon_sym_SLASH_SLASH),
	4100: uint16(21),
	4101: uint16(1),
	4102: uint16(anon_sym_SLASH_STAR),
	4103: uint16(111),
	4104: uint16(2),
	4105: uint16(sym_line_comment),
	4106: uint16(sym_block_comment),
	4107: uint16(452),
	4108: uint16(5),
	4109: uint16(anon_sym_RBRACE),
	4110: uint16(anon_sym_SEMI),
	4111: uint16(anon_sym_COMMA),
	4112: uint16(anon_sym_RPAREN),
	4113: uint16(anon_sym_GT),
	4114: uint16(4),
	4115: uint16(19),
	4116: uint16(1),
	4117: uint16(anon_sym_SLASH_SLASH),
	4118: uint16(21),
	4119: uint16(1),
	4120: uint16(anon_sym_SLASH_STAR),
	4121: uint16(112),
	4122: uint16(2),
	4123: uint16(sym_line_comment),
	4124: uint16(sym_block_comment),
	4125: uint16(454),
	4126: uint16(5),
	4127: uint16(anon_sym_RBRACE),
	4128: uint16(anon_sym_SEMI),
	4129: uint16(anon_sym_COMMA),
	4130: uint16(anon_sym_RPAREN),
	4131: uint16(anon_sym_GT),
	4132: uint16(4),
	4133: uint16(19),
	4134: uint16(1),
	4135: uint16(anon_sym_SLASH_SLASH),
	4136: uint16(21),
	4137: uint16(1),
	4138: uint16(anon_sym_SLASH_STAR),
	4139: uint16(113),
	4140: uint16(2),
	4141: uint16(sym_line_comment),
	4142: uint16(sym_block_comment),
	4143: uint16(456),
	4144: uint16(5),
	4145: uint16(anon_sym_RBRACE),
	4146: uint16(anon_sym_SEMI),
	4147: uint16(anon_sym_COMMA),
	4148: uint16(anon_sym_RPAREN),
	4149: uint16(anon_sym_GT),
	4150: uint16(4),
	4151: uint16(19),
	4152: uint16(1),
	4153: uint16(anon_sym_SLASH_SLASH),
	4154: uint16(21),
	4155: uint16(1),
	4156: uint16(anon_sym_SLASH_STAR),
	4157: uint16(114),
	4158: uint16(2),
	4159: uint16(sym_line_comment),
	4160: uint16(sym_block_comment),
	4161: uint16(458),
	4162: uint16(5),
	4163: uint16(anon_sym_RBRACE),
	4164: uint16(anon_sym_AT),
	4165: uint16(anon_sym_use),
	4166: uint16(anon_sym_world),
	4167: uint16(anon_sym_interface),
	4168: uint16(4),
	4169: uint16(19),
	4170: uint16(1),
	4171: uint16(anon_sym_SLASH_SLASH),
	4172: uint16(21),
	4173: uint16(1),
	4174: uint16(anon_sym_SLASH_STAR),
	4175: uint16(115),
	4176: uint16(2),
	4177: uint16(sym_line_comment),
	4178: uint16(sym_block_comment),
	4179: uint16(460),
	4180: uint16(5),
	4181: uint16(anon_sym_RBRACE),
	4182: uint16(anon_sym_SEMI),
	4183: uint16(anon_sym_COMMA),
	4184: uint16(anon_sym_RPAREN),
	4185: uint16(anon_sym_GT),
	4186: uint16(4),
	4187: uint16(19),
	4188: uint16(1),
	4189: uint16(anon_sym_SLASH_SLASH),
	4190: uint16(21),
	4191: uint16(1),
	4192: uint16(anon_sym_SLASH_STAR),
	4193: uint16(116),
	4194: uint16(2),
	4195: uint16(sym_line_comment),
	4196: uint16(sym_block_comment),
	4197: uint16(462),
	4198: uint16(5),
	4199: uint16(anon_sym_RBRACE),
	4200: uint16(anon_sym_SEMI),
	4201: uint16(anon_sym_COMMA),
	4202: uint16(anon_sym_RPAREN),
	4203: uint16(anon_sym_GT),
	4204: uint16(4),
	4205: uint16(19),
	4206: uint16(1),
	4207: uint16(anon_sym_SLASH_SLASH),
	4208: uint16(21),
	4209: uint16(1),
	4210: uint16(anon_sym_SLASH_STAR),
	4211: uint16(117),
	4212: uint16(2),
	4213: uint16(sym_line_comment),
	4214: uint16(sym_block_comment),
	4215: uint16(464),
	4216: uint16(5),
	4217: uint16(anon_sym_RBRACE),
	4218: uint16(anon_sym_SEMI),
	4219: uint16(anon_sym_COMMA),
	4220: uint16(anon_sym_RPAREN),
	4221: uint16(anon_sym_GT),
	4222: uint16(4),
	4223: uint16(19),
	4224: uint16(1),
	4225: uint16(anon_sym_SLASH_SLASH),
	4226: uint16(21),
	4227: uint16(1),
	4228: uint16(anon_sym_SLASH_STAR),
	4229: uint16(118),
	4230: uint16(2),
	4231: uint16(sym_line_comment),
	4232: uint16(sym_block_comment),
	4233: uint16(446),
	4234: uint16(5),
	4235: uint16(anon_sym_RBRACE),
	4236: uint16(anon_sym_SEMI),
	4237: uint16(anon_sym_COMMA),
	4238: uint16(anon_sym_RPAREN),
	4239: uint16(anon_sym_GT),
	4240: uint16(4),
	4241: uint16(19),
	4242: uint16(1),
	4243: uint16(anon_sym_SLASH_SLASH),
	4244: uint16(21),
	4245: uint16(1),
	4246: uint16(anon_sym_SLASH_STAR),
	4247: uint16(119),
	4248: uint16(2),
	4249: uint16(sym_line_comment),
	4250: uint16(sym_block_comment),
	4251: uint16(466),
	4252: uint16(5),
	4253: uint16(anon_sym_RBRACE),
	4254: uint16(anon_sym_SEMI),
	4255: uint16(anon_sym_COMMA),
	4256: uint16(anon_sym_RPAREN),
	4257: uint16(anon_sym_GT),
	4258: uint16(4),
	4259: uint16(19),
	4260: uint16(1),
	4261: uint16(anon_sym_SLASH_SLASH),
	4262: uint16(21),
	4263: uint16(1),
	4264: uint16(anon_sym_SLASH_STAR),
	4265: uint16(120),
	4266: uint16(2),
	4267: uint16(sym_line_comment),
	4268: uint16(sym_block_comment),
	4269: uint16(468),
	4270: uint16(5),
	4271: uint16(anon_sym_RBRACE),
	4272: uint16(anon_sym_SEMI),
	4273: uint16(anon_sym_COMMA),
	4274: uint16(anon_sym_RPAREN),
	4275: uint16(anon_sym_GT),
	4276: uint16(4),
	4277: uint16(19),
	4278: uint16(1),
	4279: uint16(anon_sym_SLASH_SLASH),
	4280: uint16(21),
	4281: uint16(1),
	4282: uint16(anon_sym_SLASH_STAR),
	4283: uint16(121),
	4284: uint16(2),
	4285: uint16(sym_line_comment),
	4286: uint16(sym_block_comment),
	4287: uint16(470),
	4288: uint16(5),
	4289: uint16(anon_sym_RBRACE),
	4290: uint16(anon_sym_SEMI),
	4291: uint16(anon_sym_COMMA),
	4292: uint16(anon_sym_RPAREN),
	4293: uint16(anon_sym_GT),
	4294: uint16(4),
	4295: uint16(19),
	4296: uint16(1),
	4297: uint16(anon_sym_SLASH_SLASH),
	4298: uint16(21),
	4299: uint16(1),
	4300: uint16(anon_sym_SLASH_STAR),
	4301: uint16(122),
	4302: uint16(2),
	4303: uint16(sym_line_comment),
	4304: uint16(sym_block_comment),
	4305: uint16(472),
	4306: uint16(5),
	4307: uint16(anon_sym_RBRACE),
	4308: uint16(anon_sym_SEMI),
	4309: uint16(anon_sym_COMMA),
	4310: uint16(anon_sym_RPAREN),
	4311: uint16(anon_sym_GT),
	4312: uint16(5),
	4313: uint16(19),
	4314: uint16(1),
	4315: uint16(anon_sym_SLASH_SLASH),
	4316: uint16(21),
	4317: uint16(1),
	4318: uint16(anon_sym_SLASH_STAR),
	4319: uint16(474),
	4320: uint16(2),
	4321: uint16(sym_id),
	4322: uint16(anon_sym_constructor),
	4323: uint16(476),
	4324: uint16(2),
	4325: uint16(anon_sym_RBRACE),
	4326: uint16(anon_sym_AT),
	4327: uint16(123),
	4328: uint16(2),
	4329: uint16(sym_line_comment),
	4330: uint16(sym_block_comment),
	4331: uint16(7),
	4332: uint16(19),
	4333: uint16(1),
	4334: uint16(anon_sym_SLASH_SLASH),
	4335: uint16(21),
	4336: uint16(1),
	4337: uint16(anon_sym_SLASH_STAR),
	4338: uint16(478),
	4339: uint16(1),
	4340: uint16(sym_id),
	4341: uint16(150),
	4342: uint16(1),
	4343: uint16(aux_sym_decl_head_repeat1),
	4344: uint16(255),
	4345: uint16(1),
	4346: uint16(sym_use_path),
	4347: uint16(281),
	4348: uint16(1),
	4349: uint16(sym__uri_head),
	4350: uint16(124),
	4351: uint16(2),
	4352: uint16(sym_line_comment),
	4353: uint16(sym_block_comment),
	4354: uint16(7),
	4355: uint16(19),
	4356: uint16(1),
	4357: uint16(anon_sym_SLASH_SLASH),
	4358: uint16(21),
	4359: uint16(1),
	4360: uint16(anon_sym_SLASH_STAR),
	4361: uint16(480),
	4362: uint16(1),
	4363: uint16(sym_id),
	4364: uint16(482),
	4365: uint16(1),
	4366: uint16(anon_sym_RPAREN),
	4367: uint16(172),
	4368: uint16(1),
	4369: uint16(sym_named_type),
	4370: uint16(327),
	4371: uint16(1),
	4372: uint16(sym__named_type_list),
	4373: uint16(125),
	4374: uint16(2),
	4375: uint16(sym_line_comment),
	4376: uint16(sym_block_comment),
	4377: uint16(7),
	4378: uint16(19),
	4379: uint16(1),
	4380: uint16(anon_sym_SLASH_SLASH),
	4381: uint16(21),
	4382: uint16(1),
	4383: uint16(anon_sym_SLASH_STAR),
	4384: uint16(484),
	4385: uint16(1),
	4386: uint16(sym_id),
	4387: uint16(486),
	4388: uint16(1),
	4389: uint16(anon_sym_RBRACE),
	4390: uint16(197),
	4391: uint16(1),
	4392: uint16(sym_alias_item),
	4393: uint16(223),
	4394: uint16(1),
	4395: uint16(sym_use_names_item),
	4396: uint16(126),
	4397: uint16(2),
	4398: uint16(sym_line_comment),
	4399: uint16(sym_block_comment),
	4400: uint16(7),
	4401: uint16(19),
	4402: uint16(1),
	4403: uint16(anon_sym_SLASH_SLASH),
	4404: uint16(21),
	4405: uint16(1),
	4406: uint16(anon_sym_SLASH_STAR),
	4407: uint16(488),
	4408: uint16(1),
	4409: uint16(anon_sym_external_DASHid),
	4410: uint16(490),
	4411: uint16(1),
	4412: uint16(anon_sym_unstable),
	4413: uint16(492),
	4414: uint16(1),
	4415: uint16(anon_sym_since),
	4416: uint16(494),
	4417: uint16(1),
	4418: uint16(anon_sym_deprecated),
	4419: uint16(127),
	4420: uint16(2),
	4421: uint16(sym_line_comment),
	4422: uint16(sym_block_comment),
	4423: uint16(5),
	4424: uint16(19),
	4425: uint16(1),
	4426: uint16(anon_sym_SLASH_SLASH),
	4427: uint16(21),
	4428: uint16(1),
	4429: uint16(anon_sym_SLASH_STAR),
	4430: uint16(496),
	4431: uint16(2),
	4432: uint16(sym_id),
	4433: uint16(anon_sym_constructor),
	4434: uint16(498),
	4435: uint16(2),
	4436: uint16(anon_sym_RBRACE),
	4437: uint16(anon_sym_AT),
	4438: uint16(128),
	4439: uint16(2),
	4440: uint16(sym_line_comment),
	4441: uint16(sym_block_comment),
	4442: uint16(7),
	4443: uint16(19),
	4444: uint16(1),
	4445: uint16(anon_sym_SLASH_SLASH),
	4446: uint16(21),
	4447: uint16(1),
	4448: uint16(anon_sym_SLASH_STAR),
	4449: uint16(354),
	4450: uint16(1),
	4451: uint16(anon_sym_SLASH),
	4452: uint16(379),
	4453: uint16(1),
	4454: uint16(anon_sym_COLON),
	4455: uint16(79),
	4456: uint16(1),
	4457: uint16(aux_sym_decl_head_repeat2),
	4458: uint16(88),
	4459: uint16(1),
	4460: uint16(sym__uri_tail),
	4461: uint16(129),
	4462: uint16(2),
	4463: uint16(sym_line_comment),
	4464: uint16(sym_block_comment),
	4465: uint16(7),
	4466: uint16(19),
	4467: uint16(1),
	4468: uint16(anon_sym_SLASH_SLASH),
	4469: uint16(21),
	4470: uint16(1),
	4471: uint16(anon_sym_SLASH_STAR),
	4472: uint16(500),
	4473: uint16(1),
	4474: uint16(anon_sym_async),
	4475: uint16(502),
	4476: uint16(1),
	4477: uint16(anon_sym_func),
	4478: uint16(504),
	4479: uint16(1),
	4480: uint16(anon_sym_static),
	4481: uint16(250),
	4482: uint16(1),
	4483: uint16(sym_func_type),
	4484: uint16(130),
	4485: uint16(2),
	4486: uint16(sym_line_comment),
	4487: uint16(sym_block_comment),
	4488: uint16(4),
	4489: uint16(19),
	4490: uint16(1),
	4491: uint16(anon_sym_SLASH_SLASH),
	4492: uint16(21),
	4493: uint16(1),
	4494: uint16(anon_sym_SLASH_STAR),
	4495: uint16(131),
	4496: uint16(2),
	4497: uint16(sym_line_comment),
	4498: uint16(sym_block_comment),
	4499: uint16(506),
	4500: uint16(4),
	4501: uint16(anon_sym_SEMI),
	4502: uint16(anon_sym_as),
	4503: uint16(anon_sym_with),
	4504: uint16(anon_sym_DOT),
	4505: uint16(5),
	4506: uint16(19),
	4507: uint16(1),
	4508: uint16(anon_sym_SLASH_SLASH),
	4509: uint16(21),
	4510: uint16(1),
	4511: uint16(anon_sym_SLASH_STAR),
	4512: uint16(384),
	4513: uint16(2),
	4514: uint16(anon_sym_RBRACE),
	4515: uint16(anon_sym_AT),
	4516: uint16(508),
	4517: uint16(2),
	4518: uint16(sym_id),
	4519: uint16(anon_sym_constructor),
	4520: uint16(132),
	4521: uint16(2),
	4522: uint16(sym_line_comment),
	4523: uint16(sym_block_comment),
	4524: uint16(7),
	4525: uint16(19),
	4526: uint16(1),
	4527: uint16(anon_sym_SLASH_SLASH),
	4528: uint16(21),
	4529: uint16(1),
	4530: uint16(anon_sym_SLASH_STAR),
	4531: uint16(369),
	4532: uint16(1),
	4533: uint16(sym_id),
	4534: uint16(375),
	4535: uint16(1),
	4536: uint16(anon_sym_constructor),
	4537: uint16(132),
	4538: uint16(1),
	4539: uint16(sym_resource_method),
	4540: uint16(141),
	4541: uint16(1),
	4542: uint16(sym_func_item),
	4543: uint16(133),
	4544: uint16(2),
	4545: uint16(sym_line_comment),
	4546: uint16(sym_block_comment),
	4547: uint16(7),
	4548: uint16(19),
	4549: uint16(1),
	4550: uint16(anon_sym_SLASH_SLASH),
	4551: uint16(21),
	4552: uint16(1),
	4553: uint16(anon_sym_SLASH_STAR),
	4554: uint16(510),
	4555: uint16(1),
	4556: uint16(sym_id),
	4557: uint16(150),
	4558: uint16(1),
	4559: uint16(aux_sym_decl_head_repeat1),
	4560: uint16(281),
	4561: uint16(1),
	4562: uint16(sym__uri_head),
	4563: uint16(331),
	4564: uint16(1),
	4565: uint16(sym_use_path),
	4566: uint16(134),
	4567: uint16(2),
	4568: uint16(sym_line_comment),
	4569: uint16(sym_block_comment),
	4570: uint16(7),
	4571: uint16(19),
	4572: uint16(1),
	4573: uint16(anon_sym_SLASH_SLASH),
	4574: uint16(21),
	4575: uint16(1),
	4576: uint16(anon_sym_SLASH_STAR),
	4577: uint16(510),
	4578: uint16(1),
	4579: uint16(sym_id),
	4580: uint16(150),
	4581: uint16(1),
	4582: uint16(aux_sym_decl_head_repeat1),
	4583: uint16(195),
	4584: uint16(1),
	4585: uint16(sym_use_path),
	4586: uint16(281),
	4587: uint16(1),
	4588: uint16(sym__uri_head),
	4589: uint16(135),
	4590: uint16(2),
	4591: uint16(sym_line_comment),
	4592: uint16(sym_block_comment),
	4593: uint16(7),
	4594: uint16(19),
	4595: uint16(1),
	4596: uint16(anon_sym_SLASH_SLASH),
	4597: uint16(21),
	4598: uint16(1),
	4599: uint16(anon_sym_SLASH_STAR),
	4600: uint16(484),
	4601: uint16(1),
	4602: uint16(sym_id),
	4603: uint16(161),
	4604: uint16(1),
	4605: uint16(sym_use_names_item),
	4606: uint16(197),
	4607: uint16(1),
	4608: uint16(sym_alias_item),
	4609: uint16(285),
	4610: uint16(1),
	4611: uint16(sym__use_names_list),
	4612: uint16(136),
	4613: uint16(2),
	4614: uint16(sym_line_comment),
	4615: uint16(sym_block_comment),
	4616: uint16(5),
	4617: uint16(19),
	4618: uint16(1),
	4619: uint16(anon_sym_SLASH_SLASH),
	4620: uint16(21),
	4621: uint16(1),
	4622: uint16(anon_sym_SLASH_STAR),
	4623: uint16(512),
	4624: uint16(2),
	4625: uint16(sym_id),
	4626: uint16(anon_sym_constructor),
	4627: uint16(514),
	4628: uint16(2),
	4629: uint16(anon_sym_RBRACE),
	4630: uint16(anon_sym_AT),
	4631: uint16(137),
	4632: uint16(2),
	4633: uint16(sym_line_comment),
	4634: uint16(sym_block_comment),
	4635: uint16(7),
	4636: uint16(19),
	4637: uint16(1),
	4638: uint16(anon_sym_SLASH_SLASH),
	4639: uint16(21),
	4640: uint16(1),
	4641: uint16(anon_sym_SLASH_STAR),
	4642: uint16(510),
	4643: uint16(1),
	4644: uint16(sym_id),
	4645: uint16(150),
	4646: uint16(1),
	4647: uint16(aux_sym_decl_head_repeat1),
	4648: uint16(210),
	4649: uint16(1),
	4650: uint16(sym_use_path),
	4651: uint16(281),
	4652: uint16(1),
	4653: uint16(sym__uri_head),
	4654: uint16(138),
	4655: uint16(2),
	4656: uint16(sym_line_comment),
	4657: uint16(sym_block_comment),
	4658: uint16(5),
	4659: uint16(19),
	4660: uint16(1),
	4661: uint16(anon_sym_SLASH_SLASH),
	4662: uint16(21),
	4663: uint16(1),
	4664: uint16(anon_sym_SLASH_STAR),
	4665: uint16(516),
	4666: uint16(2),
	4667: uint16(sym_id),
	4668: uint16(anon_sym_constructor),
	4669: uint16(518),
	4670: uint16(2),
	4671: uint16(anon_sym_RBRACE),
	4672: uint16(anon_sym_AT),
	4673: uint16(139),
	4674: uint16(2),
	4675: uint16(sym_line_comment),
	4676: uint16(sym_block_comment),
	4677: uint16(7),
	4678: uint16(19),
	4679: uint16(1),
	4680: uint16(anon_sym_SLASH_SLASH),
	4681: uint16(21),
	4682: uint16(1),
	4683: uint16(anon_sym_SLASH_STAR),
	4684: uint16(520),
	4685: uint16(1),
	4686: uint16(sym_id),
	4687: uint16(150),
	4688: uint16(1),
	4689: uint16(aux_sym_decl_head_repeat1),
	4690: uint16(281),
	4691: uint16(1),
	4692: uint16(sym__uri_head),
	4693: uint16(339),
	4694: uint16(1),
	4695: uint16(sym_use_path),
	4696: uint16(140),
	4697: uint16(2),
	4698: uint16(sym_line_comment),
	4699: uint16(sym_block_comment),
	4700: uint16(5),
	4701: uint16(19),
	4702: uint16(1),
	4703: uint16(anon_sym_SLASH_SLASH),
	4704: uint16(21),
	4705: uint16(1),
	4706: uint16(anon_sym_SLASH_STAR),
	4707: uint16(522),
	4708: uint16(2),
	4709: uint16(sym_id),
	4710: uint16(anon_sym_constructor),
	4711: uint16(524),
	4712: uint16(2),
	4713: uint16(anon_sym_RBRACE),
	4714: uint16(anon_sym_AT),
	4715: uint16(141),
	4716: uint16(2),
	4717: uint16(sym_line_comment),
	4718: uint16(sym_block_comment),
	4719: uint16(7),
	4720: uint16(19),
	4721: uint16(1),
	4722: uint16(anon_sym_SLASH_SLASH),
	4723: uint16(21),
	4724: uint16(1),
	4725: uint16(anon_sym_SLASH_STAR),
	4726: uint16(484),
	4727: uint16(1),
	4728: uint16(sym_id),
	4729: uint16(526),
	4730: uint16(1),
	4731: uint16(anon_sym_RBRACE),
	4732: uint16(197),
	4733: uint16(1),
	4734: uint16(sym_alias_item),
	4735: uint16(223),
	4736: uint16(1),
	4737: uint16(sym_use_names_item),
	4738: uint16(142),
	4739: uint16(2),
	4740: uint16(sym_line_comment),
	4741: uint16(sym_block_comment),
	4742: uint16(7),
	4743: uint16(19),
	4744: uint16(1),
	4745: uint16(anon_sym_SLASH_SLASH),
	4746: uint16(21),
	4747: uint16(1),
	4748: uint16(anon_sym_SLASH_STAR),
	4749: uint16(480),
	4750: uint16(1),
	4751: uint16(sym_id),
	4752: uint16(528),
	4753: uint16(1),
	4754: uint16(anon_sym_RPAREN),
	4755: uint16(172),
	4756: uint16(1),
	4757: uint16(sym_named_type),
	4758: uint16(253),
	4759: uint16(1),
	4760: uint16(sym__named_type_list),
	4761: uint16(143),
	4762: uint16(2),
	4763: uint16(sym_line_comment),
	4764: uint16(sym_block_comment),
	4765: uint16(5),
	4766: uint16(19),
	4767: uint16(1),
	4768: uint16(anon_sym_SLASH_SLASH),
	4769: uint16(21),
	4770: uint16(1),
	4771: uint16(anon_sym_SLASH_STAR),
	4772: uint16(532),
	4773: uint16(1),
	4774: uint16(anon_sym_LPAREN),
	4775: uint16(530),
	4776: uint16(2),
	4777: uint16(anon_sym_RBRACE),
	4778: uint16(anon_sym_COMMA),
	4779: uint16(144),
	4780: uint16(2),
	4781: uint16(sym_line_comment),
	4782: uint16(sym_block_comment),
	4783: uint16(5),
	4784: uint16(19),
	4785: uint16(1),
	4786: uint16(anon_sym_SLASH_SLASH),
	4787: uint16(21),
	4788: uint16(1),
	4789: uint16(anon_sym_SLASH_STAR),
	4790: uint16(534),
	4791: uint16(1),
	4792: uint16(anon_sym_COMMA),
	4793: uint16(537),
	4794: uint16(1),
	4795: uint16(anon_sym_GT),
	4796: uint16(145),
	4797: uint16(3),
	4798: uint16(sym_line_comment),
	4799: uint16(sym_block_comment),
	4800: uint16(aux_sym_tuple_list_repeat1),
	4801: uint16(6),
	4802: uint16(19),
	4803: uint16(1),
	4804: uint16(anon_sym_SLASH_SLASH),
	4805: uint16(21),
	4806: uint16(1),
	4807: uint16(anon_sym_SLASH_STAR),
	4808: uint16(539),
	4809: uint16(1),
	4810: uint16(sym_id),
	4811: uint16(152),
	4812: uint16(1),
	4813: uint16(sym_record_field),
	4814: uint16(322),
	4815: uint16(1),
	4816: uint16(sym__record_fields),
	4817: uint16(146),
	4818: uint16(2),
	4819: uint16(sym_line_comment),
	4820: uint16(sym_block_comment),
	4821: uint16(6),
	4822: uint16(19),
	4823: uint16(1),
	4824: uint16(anon_sym_SLASH_SLASH),
	4825: uint16(21),
	4826: uint16(1),
	4827: uint16(anon_sym_SLASH_STAR),
	4828: uint16(541),
	4829: uint16(1),
	4830: uint16(sym_id),
	4831: uint16(154),
	4832: uint16(1),
	4833: uint16(sym_variant_case),
	4834: uint16(249),
	4835: uint16(1),
	4836: uint16(sym__variant_cases),
	4837: uint16(147),
	4838: uint16(2),
	4839: uint16(sym_line_comment),
	4840: uint16(sym_block_comment),
	4841: uint16(6),
	4842: uint16(19),
	4843: uint16(1),
	4844: uint16(anon_sym_SLASH_SLASH),
	4845: uint16(21),
	4846: uint16(1),
	4847: uint16(anon_sym_SLASH_STAR),
	4848: uint16(543),
	4849: uint16(1),
	4850: uint16(sym_id),
	4851: uint16(163),
	4852: uint16(1),
	4853: uint16(aux_sym_decl_head_repeat1),
	4854: uint16(281),
	4855: uint16(1),
	4856: uint16(sym__uri_head),
	4857: uint16(148),
	4858: uint16(2),
	4859: uint16(sym_line_comment),
	4860: uint16(sym_block_comment),
	4861: uint16(6),
	4862: uint16(19),
	4863: uint16(1),
	4864: uint16(anon_sym_SLASH_SLASH),
	4865: uint16(21),
	4866: uint16(1),
	4867: uint16(anon_sym_SLASH_STAR),
	4868: uint16(545),
	4869: uint16(1),
	4870: uint16(anon_sym_STAR_SLASH),
	4871: uint16(547),
	4872: uint16(1),
	4873: uint16(sym__block_comment_content),
	4874: uint16(549),
	4875: uint16(1),
	4876: uint16(sym__block_doc_comment_marker),
	4877: uint16(149),
	4878: uint16(2),
	4879: uint16(sym_line_comment),
	4880: uint16(sym_block_comment),
	4881: uint16(6),
	4882: uint16(19),
	4883: uint16(1),
	4884: uint16(anon_sym_SLASH_SLASH),
	4885: uint16(21),
	4886: uint16(1),
	4887: uint16(anon_sym_SLASH_STAR),
	4888: uint16(551),
	4889: uint16(1),
	4890: uint16(sym_id),
	4891: uint16(163),
	4892: uint16(1),
	4893: uint16(aux_sym_decl_head_repeat1),
	4894: uint16(281),
	4895: uint16(1),
	4896: uint16(sym__uri_head),
	4897: uint16(150),
	4898: uint16(2),
	4899: uint16(sym_line_comment),
	4900: uint16(sym_block_comment),
	4901: uint16(6),
	4902: uint16(19),
	4903: uint16(1),
	4904: uint16(anon_sym_SLASH_SLASH),
	4905: uint16(21),
	4906: uint16(1),
	4907: uint16(anon_sym_SLASH_STAR),
	4908: uint16(553),
	4909: uint16(1),
	4910: uint16(sym_id),
	4911: uint16(162),
	4912: uint16(1),
	4913: uint16(sym_include_names_item),
	4914: uint16(291),
	4915: uint16(1),
	4916: uint16(sym__include_names_list),
	4917: uint16(151),
	4918: uint16(2),
	4919: uint16(sym_line_comment),
	4920: uint16(sym_block_comment),
	4921: uint16(6),
	4922: uint16(19),
	4923: uint16(1),
	4924: uint16(anon_sym_SLASH_SLASH),
	4925: uint16(21),
	4926: uint16(1),
	4927: uint16(anon_sym_SLASH_STAR),
	4928: uint16(555),
	4929: uint16(1),
	4930: uint16(anon_sym_RBRACE),
	4931: uint16(557),
	4932: uint16(1),
	4933: uint16(anon_sym_COMMA),
	4934: uint16(165),
	4935: uint16(1),
	4936: uint16(aux_sym__record_fields_repeat1),
	4937: uint16(152),
	4938: uint16(2),
	4939: uint16(sym_line_comment),
	4940: uint16(sym_block_comment),
	4941: uint16(6),
	4942: uint16(19),
	4943: uint16(1),
	4944: uint16(anon_sym_SLASH_SLASH),
	4945: uint16(21),
	4946: uint16(1),
	4947: uint16(anon_sym_SLASH_STAR),
	4948: uint16(559),
	4949: uint16(1),
	4950: uint16(anon_sym_RBRACE),
	4951: uint16(561),
	4952: uint16(1),
	4953: uint16(anon_sym_COMMA),
	4954: uint16(166),
	4955: uint16(1),
	4956: uint16(aux_sym__flags_fields_repeat1),
	4957: uint16(153),
	4958: uint16(2),
	4959: uint16(sym_line_comment),
	4960: uint16(sym_block_comment),
	4961: uint16(6),
	4962: uint16(19),
	4963: uint16(1),
	4964: uint16(anon_sym_SLASH_SLASH),
	4965: uint16(21),
	4966: uint16(1),
	4967: uint16(anon_sym_SLASH_STAR),
	4968: uint16(563),
	4969: uint16(1),
	4970: uint16(anon_sym_RBRACE),
	4971: uint16(565),
	4972: uint16(1),
	4973: uint16(anon_sym_COMMA),
	4974: uint16(168),
	4975: uint16(1),
	4976: uint16(aux_sym__variant_cases_repeat1),
	4977: uint16(154),
	4978: uint16(2),
	4979: uint16(sym_line_comment),
	4980: uint16(sym_block_comment),
	4981: uint16(6),
	4982: uint16(19),
	4983: uint16(1),
	4984: uint16(anon_sym_SLASH_SLASH),
	4985: uint16(21),
	4986: uint16(1),
	4987: uint16(anon_sym_SLASH_STAR),
	4988: uint16(567),
	4989: uint16(1),
	4990: uint16(anon_sym_RBRACE),
	4991: uint16(569),
	4992: uint16(1),
	4993: uint16(anon_sym_COMMA),
	4994: uint16(169),
	4995: uint16(1),
	4996: uint16(aux_sym__enum_cases_repeat1),
	4997: uint16(155),
	4998: uint16(2),
	4999: uint16(sym_line_comment),
	5000: uint16(sym_block_comment),
	5001: uint16(6),
	5002: uint16(19),
	5003: uint16(1),
	5004: uint16(anon_sym_SLASH_SLASH),
	5005: uint16(21),
	5006: uint16(1),
	5007: uint16(anon_sym_SLASH_STAR),
	5008: uint16(490),
	5009: uint16(1),
	5010: uint16(anon_sym_unstable),
	5011: uint16(492),
	5012: uint16(1),
	5013: uint16(anon_sym_since),
	5014: uint16(494),
	5015: uint16(1),
	5016: uint16(anon_sym_deprecated),
	5017: uint16(156),
	5018: uint16(2),
	5019: uint16(sym_line_comment),
	5020: uint16(sym_block_comment),
	5021: uint16(6),
	5022: uint16(19),
	5023: uint16(1),
	5024: uint16(anon_sym_SLASH_SLASH),
	5025: uint16(21),
	5026: uint16(1),
	5027: uint16(anon_sym_SLASH_STAR),
	5028: uint16(571),
	5029: uint16(1),
	5030: uint16(sym_id),
	5031: uint16(148),
	5032: uint16(1),
	5033: uint16(aux_sym_decl_head_repeat1),
	5034: uint16(281),
	5035: uint16(1),
	5036: uint16(sym__uri_head),
	5037: uint16(157),
	5038: uint16(2),
	5039: uint16(sym_line_comment),
	5040: uint16(sym_block_comment),
	5041: uint16(6),
	5042: uint16(3),
	5043: uint16(1),
	5044: uint16(anon_sym_SLASH_SLASH),
	5045: uint16(5),
	5046: uint16(1),
	5047: uint16(anon_sym_SLASH_STAR),
	5048: uint16(573),
	5049: uint16(1),
	5050: uint16(anon_sym_SLASH_SLASH2),
	5051: uint16(575),
	5052: uint16(1),
	5053: uint16(anon_sym_SLASH2),
	5054: uint16(577),
	5055: uint16(1),
	5056: uint16(aux_sym_line_comment_token2),
	5057: uint16(158),
	5058: uint16(2),
	5059: uint16(sym_line_comment),
	5060: uint16(sym_block_comment),
	5061: uint16(6),
	5062: uint16(19),
	5063: uint16(1),
	5064: uint16(anon_sym_SLASH_SLASH),
	5065: uint16(21),
	5066: uint16(1),
	5067: uint16(anon_sym_SLASH_STAR),
	5068: uint16(579),
	5069: uint16(1),
	5070: uint16(anon_sym_SEMI),
	5071: uint16(581),
	5072: uint16(1),
	5073: uint16(anon_sym_DASH_GT),
	5074: uint16(256),
	5075: uint16(1),
	5076: uint16(sym_result_list),
	5077: uint16(159),
	5078: uint16(2),
	5079: uint16(sym_line_comment),
	5080: uint16(sym_block_comment),
	5081: uint16(5),
	5082: uint16(19),
	5083: uint16(1),
	5084: uint16(anon_sym_SLASH_SLASH),
	5085: uint16(21),
	5086: uint16(1),
	5087: uint16(anon_sym_SLASH_STAR),
	5088: uint16(585),
	5089: uint16(1),
	5090: uint16(anon_sym_as),
	5091: uint16(583),
	5092: uint16(2),
	5093: uint16(anon_sym_RBRACE),
	5094: uint16(anon_sym_COMMA),
	5095: uint16(160),
	5096: uint16(2),
	5097: uint16(sym_line_comment),
	5098: uint16(sym_block_comment),
	5099: uint16(6),
	5100: uint16(19),
	5101: uint16(1),
	5102: uint16(anon_sym_SLASH_SLASH),
	5103: uint16(21),
	5104: uint16(1),
	5105: uint16(anon_sym_SLASH_STAR),
	5106: uint16(587),
	5107: uint16(1),
	5108: uint16(anon_sym_RBRACE),
	5109: uint16(589),
	5110: uint16(1),
	5111: uint16(anon_sym_COMMA),
	5112: uint16(193),
	5113: uint16(1),
	5114: uint16(aux_sym__use_names_list_repeat1),
	5115: uint16(161),
	5116: uint16(2),
	5117: uint16(sym_line_comment),
	5118: uint16(sym_block_comment),
	5119: uint16(6),
	5120: uint16(19),
	5121: uint16(1),
	5122: uint16(anon_sym_SLASH_SLASH),
	5123: uint16(21),
	5124: uint16(1),
	5125: uint16(anon_sym_SLASH_STAR),
	5126: uint16(591),
	5127: uint16(1),
	5128: uint16(anon_sym_RBRACE),
	5129: uint16(593),
	5130: uint16(1),
	5131: uint16(anon_sym_COMMA),
	5132: uint16(174),
	5133: uint16(1),
	5134: uint16(aux_sym__include_names_list_repeat1),
	5135: uint16(162),
	5136: uint16(2),
	5137: uint16(sym_line_comment),
	5138: uint16(sym_block_comment),
	5139: uint16(5),
	5140: uint16(19),
	5141: uint16(1),
	5142: uint16(anon_sym_SLASH_SLASH),
	5143: uint16(21),
	5144: uint16(1),
	5145: uint16(anon_sym_SLASH_STAR),
	5146: uint16(595),
	5147: uint16(1),
	5148: uint16(sym_id),
	5149: uint16(281),
	5150: uint16(1),
	5151: uint16(sym__uri_head),
	5152: uint16(163),
	5153: uint16(3),
	5154: uint16(sym_line_comment),
	5155: uint16(sym_block_comment),
	5156: uint16(aux_sym_decl_head_repeat1),
	5157: uint16(6),
	5158: uint16(19),
	5159: uint16(1),
	5160: uint16(anon_sym_SLASH_SLASH),
	5161: uint16(21),
	5162: uint16(1),
	5163: uint16(anon_sym_SLASH_STAR),
	5164: uint16(539),
	5165: uint16(1),
	5166: uint16(sym_id),
	5167: uint16(598),
	5168: uint16(1),
	5169: uint16(anon_sym_RBRACE),
	5170: uint16(203),
	5171: uint16(1),
	5172: uint16(sym_record_field),
	5173: uint16(164),
	5174: uint16(2),
	5175: uint16(sym_line_comment),
	5176: uint16(sym_block_comment),
	5177: uint16(6),
	5178: uint16(19),
	5179: uint16(1),
	5180: uint16(anon_sym_SLASH_SLASH),
	5181: uint16(21),
	5182: uint16(1),
	5183: uint16(anon_sym_SLASH_STAR),
	5184: uint16(598),
	5185: uint16(1),
	5186: uint16(anon_sym_RBRACE),
	5187: uint16(600),
	5188: uint16(1),
	5189: uint16(anon_sym_COMMA),
	5190: uint16(177),
	5191: uint16(1),
	5192: uint16(aux_sym__record_fields_repeat1),
	5193: uint16(165),
	5194: uint16(2),
	5195: uint16(sym_line_comment),
	5196: uint16(sym_block_comment),
	5197: uint16(6),
	5198: uint16(19),
	5199: uint16(1),
	5200: uint16(anon_sym_SLASH_SLASH),
	5201: uint16(21),
	5202: uint16(1),
	5203: uint16(anon_sym_SLASH_STAR),
	5204: uint16(602),
	5205: uint16(1),
	5206: uint16(anon_sym_RBRACE),
	5207: uint16(604),
	5208: uint16(1),
	5209: uint16(anon_sym_COMMA),
	5210: uint16(178),
	5211: uint16(1),
	5212: uint16(aux_sym__flags_fields_repeat1),
	5213: uint16(166),
	5214: uint16(2),
	5215: uint16(sym_line_comment),
	5216: uint16(sym_block_comment),
	5217: uint16(6),
	5218: uint16(19),
	5219: uint16(1),
	5220: uint16(anon_sym_SLASH_SLASH),
	5221: uint16(21),
	5222: uint16(1),
	5223: uint16(anon_sym_SLASH_STAR),
	5224: uint16(541),
	5225: uint16(1),
	5226: uint16(sym_id),
	5227: uint16(606),
	5228: uint16(1),
	5229: uint16(anon_sym_RBRACE),
	5230: uint16(209),
	5231: uint16(1),
	5232: uint16(sym_variant_case),
	5233: uint16(167),
	5234: uint16(2),
	5235: uint16(sym_line_comment),
	5236: uint16(sym_block_comment),
	5237: uint16(6),
	5238: uint16(19),
	5239: uint16(1),
	5240: uint16(anon_sym_SLASH_SLASH),
	5241: uint16(21),
	5242: uint16(1),
	5243: uint16(anon_sym_SLASH_STAR),
	5244: uint16(606),
	5245: uint16(1),
	5246: uint16(anon_sym_RBRACE),
	5247: uint16(608),
	5248: uint16(1),
	5249: uint16(anon_sym_COMMA),
	5250: uint16(180),
	5251: uint16(1),
	5252: uint16(aux_sym__variant_cases_repeat1),
	5253: uint16(168),
	5254: uint16(2),
	5255: uint16(sym_line_comment),
	5256: uint16(sym_block_comment),
	5257: uint16(6),
	5258: uint16(19),
	5259: uint16(1),
	5260: uint16(anon_sym_SLASH_SLASH),
	5261: uint16(21),
	5262: uint16(1),
	5263: uint16(anon_sym_SLASH_STAR),
	5264: uint16(610),
	5265: uint16(1),
	5266: uint16(anon_sym_RBRACE),
	5267: uint16(612),
	5268: uint16(1),
	5269: uint16(anon_sym_COMMA),
	5270: uint16(181),
	5271: uint16(1),
	5272: uint16(aux_sym__enum_cases_repeat1),
	5273: uint16(169),
	5274: uint16(2),
	5275: uint16(sym_line_comment),
	5276: uint16(sym_block_comment),
	5277: uint16(6),
	5278: uint16(19),
	5279: uint16(1),
	5280: uint16(anon_sym_SLASH_SLASH),
	5281: uint16(21),
	5282: uint16(1),
	5283: uint16(anon_sym_SLASH_STAR),
	5284: uint16(581),
	5285: uint16(1),
	5286: uint16(anon_sym_DASH_GT),
	5287: uint16(614),
	5288: uint16(1),
	5289: uint16(anon_sym_SEMI),
	5290: uint16(289),
	5291: uint16(1),
	5292: uint16(sym_result_list),
	5293: uint16(170),
	5294: uint16(2),
	5295: uint16(sym_line_comment),
	5296: uint16(sym_block_comment),
	5297: uint16(6),
	5298: uint16(19),
	5299: uint16(1),
	5300: uint16(anon_sym_SLASH_SLASH),
	5301: uint16(21),
	5302: uint16(1),
	5303: uint16(anon_sym_SLASH_STAR),
	5304: uint16(581),
	5305: uint16(1),
	5306: uint16(anon_sym_DASH_GT),
	5307: uint16(616),
	5308: uint16(1),
	5309: uint16(anon_sym_SEMI),
	5310: uint16(290),
	5311: uint16(1),
	5312: uint16(sym_result_list),
	5313: uint16(171),
	5314: uint16(2),
	5315: uint16(sym_line_comment),
	5316: uint16(sym_block_comment),
	5317: uint16(6),
	5318: uint16(19),
	5319: uint16(1),
	5320: uint16(anon_sym_SLASH_SLASH),
	5321: uint16(21),
	5322: uint16(1),
	5323: uint16(anon_sym_SLASH_STAR),
	5324: uint16(618),
	5325: uint16(1),
	5326: uint16(anon_sym_COMMA),
	5327: uint16(620),
	5328: uint16(1),
	5329: uint16(anon_sym_RPAREN),
	5330: uint16(184),
	5331: uint16(1),
	5332: uint16(aux_sym__named_type_list_repeat1),
	5333: uint16(172),
	5334: uint16(2),
	5335: uint16(sym_line_comment),
	5336: uint16(sym_block_comment),
	5337: uint16(6),
	5338: uint16(19),
	5339: uint16(1),
	5340: uint16(anon_sym_SLASH_SLASH),
	5341: uint16(21),
	5342: uint16(1),
	5343: uint16(anon_sym_SLASH_STAR),
	5344: uint16(553),
	5345: uint16(1),
	5346: uint16(sym_id),
	5347: uint16(622),
	5348: uint16(1),
	5349: uint16(anon_sym_RBRACE),
	5350: uint16(227),
	5351: uint16(1),
	5352: uint16(sym_include_names_item),
	5353: uint16(173),
	5354: uint16(2),
	5355: uint16(sym_line_comment),
	5356: uint16(sym_block_comment),
	5357: uint16(6),
	5358: uint16(19),
	5359: uint16(1),
	5360: uint16(anon_sym_SLASH_SLASH),
	5361: uint16(21),
	5362: uint16(1),
	5363: uint16(anon_sym_SLASH_STAR),
	5364: uint16(622),
	5365: uint16(1),
	5366: uint16(anon_sym_RBRACE),
	5367: uint16(624),
	5368: uint16(1),
	5369: uint16(anon_sym_COMMA),
	5370: uint16(187),
	5371: uint16(1),
	5372: uint16(aux_sym__include_names_list_repeat1),
	5373: uint16(174),
	5374: uint16(2),
	5375: uint16(sym_line_comment),
	5376: uint16(sym_block_comment),
	5377: uint16(6),
	5378: uint16(19),
	5379: uint16(1),
	5380: uint16(anon_sym_SLASH_SLASH),
	5381: uint16(21),
	5382: uint16(1),
	5383: uint16(anon_sym_SLASH_STAR),
	5384: uint16(626),
	5385: uint16(1),
	5386: uint16(anon_sym_COMMA),
	5387: uint16(628),
	5388: uint16(1),
	5389: uint16(anon_sym_GT),
	5390: uint16(188),
	5391: uint16(1),
	5392: uint16(aux_sym_tuple_list_repeat1),
	5393: uint16(175),
	5394: uint16(2),
	5395: uint16(sym_line_comment),
	5396: uint16(sym_block_comment),
	5397: uint16(6),
	5398: uint16(19),
	5399: uint16(1),
	5400: uint16(anon_sym_SLASH_SLASH),
	5401: uint16(21),
	5402: uint16(1),
	5403: uint16(anon_sym_SLASH_STAR),
	5404: uint16(539),
	5405: uint16(1),
	5406: uint16(sym_id),
	5407: uint16(630),
	5408: uint16(1),
	5409: uint16(anon_sym_RBRACE),
	5410: uint16(203),
	5411: uint16(1),
	5412: uint16(sym_record_field),
	5413: uint16(176),
	5414: uint16(2),
	5415: uint16(sym_line_comment),
	5416: uint16(sym_block_comment),
	5417: uint16(5),
	5418: uint16(19),
	5419: uint16(1),
	5420: uint16(anon_sym_SLASH_SLASH),
	5421: uint16(21),
	5422: uint16(1),
	5423: uint16(anon_sym_SLASH_STAR),
	5424: uint16(632),
	5425: uint16(1),
	5426: uint16(anon_sym_RBRACE),
	5427: uint16(634),
	5428: uint16(1),
	5429: uint16(anon_sym_COMMA),
	5430: uint16(177),
	5431: uint16(3),
	5432: uint16(sym_line_comment),
	5433: uint16(sym_block_comment),
	5434: uint16(aux_sym__record_fields_repeat1),
	5435: uint16(5),
	5436: uint16(19),
	5437: uint16(1),
	5438: uint16(anon_sym_SLASH_SLASH),
	5439: uint16(21),
	5440: uint16(1),
	5441: uint16(anon_sym_SLASH_STAR),
	5442: uint16(637),
	5443: uint16(1),
	5444: uint16(anon_sym_RBRACE),
	5445: uint16(639),
	5446: uint16(1),
	5447: uint16(anon_sym_COMMA),
	5448: uint16(178),
	5449: uint16(3),
	5450: uint16(sym_line_comment),
	5451: uint16(sym_block_comment),
	5452: uint16(aux_sym__flags_fields_repeat1),
	5453: uint16(6),
	5454: uint16(19),
	5455: uint16(1),
	5456: uint16(anon_sym_SLASH_SLASH),
	5457: uint16(21),
	5458: uint16(1),
	5459: uint16(anon_sym_SLASH_STAR),
	5460: uint16(541),
	5461: uint16(1),
	5462: uint16(sym_id),
	5463: uint16(642),
	5464: uint16(1),
	5465: uint16(anon_sym_RBRACE),
	5466: uint16(209),
	5467: uint16(1),
	5468: uint16(sym_variant_case),
	5469: uint16(179),
	5470: uint16(2),
	5471: uint16(sym_line_comment),
	5472: uint16(sym_block_comment),
	5473: uint16(5),
	5474: uint16(19),
	5475: uint16(1),
	5476: uint16(anon_sym_SLASH_SLASH),
	5477: uint16(21),
	5478: uint16(1),
	5479: uint16(anon_sym_SLASH_STAR),
	5480: uint16(644),
	5481: uint16(1),
	5482: uint16(anon_sym_RBRACE),
	5483: uint16(646),
	5484: uint16(1),
	5485: uint16(anon_sym_COMMA),
	5486: uint16(180),
	5487: uint16(3),
	5488: uint16(sym_line_comment),
	5489: uint16(sym_block_comment),
	5490: uint16(aux_sym__variant_cases_repeat1),
	5491: uint16(5),
	5492: uint16(19),
	5493: uint16(1),
	5494: uint16(anon_sym_SLASH_SLASH),
	5495: uint16(21),
	5496: uint16(1),
	5497: uint16(anon_sym_SLASH_STAR),
	5498: uint16(649),
	5499: uint16(1),
	5500: uint16(anon_sym_RBRACE),
	5501: uint16(651),
	5502: uint16(1),
	5503: uint16(anon_sym_COMMA),
	5504: uint16(181),
	5505: uint16(3),
	5506: uint16(sym_line_comment),
	5507: uint16(sym_block_comment),
	5508: uint16(aux_sym__enum_cases_repeat1),
	5509: uint16(6),
	5510: uint16(19),
	5511: uint16(1),
	5512: uint16(anon_sym_SLASH_SLASH),
	5513: uint16(21),
	5514: uint16(1),
	5515: uint16(anon_sym_SLASH_STAR),
	5516: uint16(500),
	5517: uint16(1),
	5518: uint16(anon_sym_async),
	5519: uint16(502),
	5520: uint16(1),
	5521: uint16(anon_sym_func),
	5522: uint16(320),
	5523: uint16(1),
	5524: uint16(sym_func_type),
	5525: uint16(182),
	5526: uint16(2),
	5527: uint16(sym_line_comment),
	5528: uint16(sym_block_comment),
	5529: uint16(6),
	5530: uint16(19),
	5531: uint16(1),
	5532: uint16(anon_sym_SLASH_SLASH),
	5533: uint16(21),
	5534: uint16(1),
	5535: uint16(anon_sym_SLASH_STAR),
	5536: uint16(480),
	5537: uint16(1),
	5538: uint16(sym_id),
	5539: uint16(654),
	5540: uint16(1),
	5541: uint16(anon_sym_RPAREN),
	5542: uint16(239),
	5543: uint16(1),
	5544: uint16(sym_named_type),
	5545: uint16(183),
	5546: uint16(2),
	5547: uint16(sym_line_comment),
	5548: uint16(sym_block_comment),
	5549: uint16(6),
	5550: uint16(19),
	5551: uint16(1),
	5552: uint16(anon_sym_SLASH_SLASH),
	5553: uint16(21),
	5554: uint16(1),
	5555: uint16(anon_sym_SLASH_STAR),
	5556: uint16(654),
	5557: uint16(1),
	5558: uint16(anon_sym_RPAREN),
	5559: uint16(656),
	5560: uint16(1),
	5561: uint16(anon_sym_COMMA),
	5562: uint16(191),
	5563: uint16(1),
	5564: uint16(aux_sym__named_type_list_repeat1),
	5565: uint16(184),
	5566: uint16(2),
	5567: uint16(sym_line_comment),
	5568: uint16(sym_block_comment),
	5569: uint16(5),
	5570: uint16(19),
	5571: uint16(1),
	5572: uint16(anon_sym_SLASH_SLASH),
	5573: uint16(21),
	5574: uint16(1),
	5575: uint16(anon_sym_SLASH_STAR),
	5576: uint16(658),
	5577: uint16(1),
	5578: uint16(anon_sym_RBRACE),
	5579: uint16(660),
	5580: uint16(1),
	5581: uint16(anon_sym_COMMA),
	5582: uint16(185),
	5583: uint16(3),
	5584: uint16(sym_line_comment),
	5585: uint16(sym_block_comment),
	5586: uint16(aux_sym__use_names_list_repeat1),
	5587: uint16(6),
	5588: uint16(19),
	5589: uint16(1),
	5590: uint16(anon_sym_SLASH_SLASH),
	5591: uint16(21),
	5592: uint16(1),
	5593: uint16(anon_sym_SLASH_STAR),
	5594: uint16(553),
	5595: uint16(1),
	5596: uint16(sym_id),
	5597: uint16(663),
	5598: uint16(1),
	5599: uint16(anon_sym_RBRACE),
	5600: uint16(227),
	5601: uint16(1),
	5602: uint16(sym_include_names_item),
	5603: uint16(186),
	5604: uint16(2),
	5605: uint16(sym_line_comment),
	5606: uint16(sym_block_comment),
	5607: uint16(5),
	5608: uint16(19),
	5609: uint16(1),
	5610: uint16(anon_sym_SLASH_SLASH),
	5611: uint16(21),
	5612: uint16(1),
	5613: uint16(anon_sym_SLASH_STAR),
	5614: uint16(665),
	5615: uint16(1),
	5616: uint16(anon_sym_RBRACE),
	5617: uint16(667),
	5618: uint16(1),
	5619: uint16(anon_sym_COMMA),
	5620: uint16(187),
	5621: uint16(3),
	5622: uint16(sym_line_comment),
	5623: uint16(sym_block_comment),
	5624: uint16(aux_sym__include_names_list_repeat1),
	5625: uint16(6),
	5626: uint16(19),
	5627: uint16(1),
	5628: uint16(anon_sym_SLASH_SLASH),
	5629: uint16(21),
	5630: uint16(1),
	5631: uint16(anon_sym_SLASH_STAR),
	5632: uint16(51),
	5633: uint16(1),
	5634: uint16(anon_sym_GT),
	5635: uint16(670),
	5636: uint16(1),
	5637: uint16(anon_sym_COMMA),
	5638: uint16(145),
	5639: uint16(1),
	5640: uint16(aux_sym_tuple_list_repeat1),
	5641: uint16(188),
	5642: uint16(2),
	5643: uint16(sym_line_comment),
	5644: uint16(sym_block_comment),
	5645: uint16(6),
	5646: uint16(19),
	5647: uint16(1),
	5648: uint16(anon_sym_SLASH_SLASH),
	5649: uint16(21),
	5650: uint16(1),
	5651: uint16(anon_sym_SLASH_STAR),
	5652: uint16(500),
	5653: uint16(1),
	5654: uint16(anon_sym_async),
	5655: uint16(502),
	5656: uint16(1),
	5657: uint16(anon_sym_func),
	5658: uint16(250),
	5659: uint16(1),
	5660: uint16(sym_func_type),
	5661: uint16(189),
	5662: uint16(2),
	5663: uint16(sym_line_comment),
	5664: uint16(sym_block_comment),
	5665: uint16(6),
	5666: uint16(19),
	5667: uint16(1),
	5668: uint16(anon_sym_SLASH_SLASH),
	5669: uint16(21),
	5670: uint16(1),
	5671: uint16(anon_sym_SLASH_STAR),
	5672: uint16(480),
	5673: uint16(1),
	5674: uint16(sym_id),
	5675: uint16(672),
	5676: uint16(1),
	5677: uint16(anon_sym_RPAREN),
	5678: uint16(239),
	5679: uint16(1),
	5680: uint16(sym_named_type),
	5681: uint16(190),
	5682: uint16(2),
	5683: uint16(sym_line_comment),
	5684: uint16(sym_block_comment),
	5685: uint16(5),
	5686: uint16(19),
	5687: uint16(1),
	5688: uint16(anon_sym_SLASH_SLASH),
	5689: uint16(21),
	5690: uint16(1),
	5691: uint16(anon_sym_SLASH_STAR),
	5692: uint16(674),
	5693: uint16(1),
	5694: uint16(anon_sym_COMMA),
	5695: uint16(677),
	5696: uint16(1),
	5697: uint16(anon_sym_RPAREN),
	5698: uint16(191),
	5699: uint16(3),
	5700: uint16(sym_line_comment),
	5701: uint16(sym_block_comment),
	5702: uint16(aux_sym__named_type_list_repeat1),
	5703: uint16(6),
	5704: uint16(19),
	5705: uint16(1),
	5706: uint16(anon_sym_SLASH_SLASH),
	5707: uint16(21),
	5708: uint16(1),
	5709: uint16(anon_sym_SLASH_STAR),
	5710: uint16(484),
	5711: uint16(1),
	5712: uint16(sym_id),
	5713: uint16(197),
	5714: uint16(1),
	5715: uint16(sym_alias_item),
	5716: uint16(223),
	5717: uint16(1),
	5718: uint16(sym_use_names_item),
	5719: uint16(192),
	5720: uint16(2),
	5721: uint16(sym_line_comment),
	5722: uint16(sym_block_comment),
	5723: uint16(6),
	5724: uint16(19),
	5725: uint16(1),
	5726: uint16(anon_sym_SLASH_SLASH),
	5727: uint16(21),
	5728: uint16(1),
	5729: uint16(anon_sym_SLASH_STAR),
	5730: uint16(526),
	5731: uint16(1),
	5732: uint16(anon_sym_RBRACE),
	5733: uint16(679),
	5734: uint16(1),
	5735: uint16(anon_sym_COMMA),
	5736: uint16(185),
	5737: uint16(1),
	5738: uint16(aux_sym__use_names_list_repeat1),
	5739: uint16(193),
	5740: uint16(2),
	5741: uint16(sym_line_comment),
	5742: uint16(sym_block_comment),
	5743: uint16(5),
	5744: uint16(19),
	5745: uint16(1),
	5746: uint16(anon_sym_SLASH_SLASH),
	5747: uint16(21),
	5748: uint16(1),
	5749: uint16(anon_sym_SLASH_STAR),
	5750: uint16(681),
	5751: uint16(1),
	5752: uint16(anon_sym_LBRACE),
	5753: uint16(49),
	5754: uint16(1),
	5755: uint16(sym__flags_body),
	5756: uint16(194),
	5757: uint16(2),
	5758: uint16(sym_line_comment),
	5759: uint16(sym_block_comment),
	5760: uint16(5),
	5761: uint16(19),
	5762: uint16(1),
	5763: uint16(anon_sym_SLASH_SLASH),
	5764: uint16(21),
	5765: uint16(1),
	5766: uint16(anon_sym_SLASH_STAR),
	5767: uint16(683),
	5768: uint16(1),
	5769: uint16(anon_sym_SEMI),
	5770: uint16(685),
	5771: uint16(1),
	5772: uint16(anon_sym_with),
	5773: uint16(195),
	5774: uint16(2),
	5775: uint16(sym_line_comment),
	5776: uint16(sym_block_comment),
	5777: uint16(5),
	5778: uint16(19),
	5779: uint16(1),
	5780: uint16(anon_sym_SLASH_SLASH),
	5781: uint16(21),
	5782: uint16(1),
	5783: uint16(anon_sym_SLASH_STAR),
	5784: uint16(687),
	5785: uint16(1),
	5786: uint16(anon_sym_LBRACE),
	5787: uint16(265),
	5788: uint16(1),
	5789: uint16(sym__use_names_body),
	5790: uint16(196),
	5791: uint16(2),
	5792: uint16(sym_line_comment),
	5793: uint16(sym_block_comment),
	5794: uint16(4),
	5795: uint16(19),
	5796: uint16(1),
	5797: uint16(anon_sym_SLASH_SLASH),
	5798: uint16(21),
	5799: uint16(1),
	5800: uint16(anon_sym_SLASH_STAR),
	5801: uint16(583),
	5802: uint16(2),
	5803: uint16(anon_sym_RBRACE),
	5804: uint16(anon_sym_COMMA),
	5805: uint16(197),
	5806: uint16(2),
	5807: uint16(sym_line_comment),
	5808: uint16(sym_block_comment),
	5809: uint16(5),
	5810: uint16(19),
	5811: uint16(1),
	5812: uint16(anon_sym_SLASH_SLASH),
	5813: uint16(21),
	5814: uint16(1),
	5815: uint16(anon_sym_SLASH_STAR),
	5816: uint16(689),
	5817: uint16(1),
	5818: uint16(anon_sym_LBRACE),
	5819: uint16(92),
	5820: uint16(1),
	5821: uint16(sym__interface_body),
	5822: uint16(198),
	5823: uint16(2),
	5824: uint16(sym_line_comment),
	5825: uint16(sym_block_comment),
	5826: uint16(5),
	5827: uint16(19),
	5828: uint16(1),
	5829: uint16(anon_sym_SLASH_SLASH),
	5830: uint16(21),
	5831: uint16(1),
	5832: uint16(anon_sym_SLASH_STAR),
	5833: uint16(689),
	5834: uint16(1),
	5835: uint16(anon_sym_LBRACE),
	5836: uint16(66),
	5837: uint16(1),
	5838: uint16(sym__interface_body),
	5839: uint16(199),
	5840: uint16(2),
	5841: uint16(sym_line_comment),
	5842: uint16(sym_block_comment),
	5843: uint16(5),
	5844: uint16(19),
	5845: uint16(1),
	5846: uint16(anon_sym_SLASH_SLASH),
	5847: uint16(21),
	5848: uint16(1),
	5849: uint16(anon_sym_SLASH_STAR),
	5850: uint16(691),
	5851: uint16(1),
	5852: uint16(anon_sym_COMMA),
	5853: uint16(693),
	5854: uint16(1),
	5855: uint16(anon_sym_GT),
	5856: uint16(200),
	5857: uint16(2),
	5858: uint16(sym_line_comment),
	5859: uint16(sym_block_comment),
	5860: uint16(5),
	5861: uint16(19),
	5862: uint16(1),
	5863: uint16(anon_sym_SLASH_SLASH),
	5864: uint16(21),
	5865: uint16(1),
	5866: uint16(anon_sym_SLASH_STAR),
	5867: uint16(695),
	5868: uint16(1),
	5869: uint16(anon_sym_COMMA),
	5870: uint16(697),
	5871: uint16(1),
	5872: uint16(anon_sym_GT),
	5873: uint16(201),
	5874: uint16(2),
	5875: uint16(sym_line_comment),
	5876: uint16(sym_block_comment),
	5877: uint16(4),
	5878: uint16(19),
	5879: uint16(1),
	5880: uint16(anon_sym_SLASH_SLASH),
	5881: uint16(21),
	5882: uint16(1),
	5883: uint16(anon_sym_SLASH_STAR),
	5884: uint16(699),
	5885: uint16(2),
	5886: uint16(anon_sym_RBRACE),
	5887: uint16(anon_sym_COMMA),
	5888: uint16(202),
	5889: uint16(2),
	5890: uint16(sym_line_comment),
	5891: uint16(sym_block_comment),
	5892: uint16(4),
	5893: uint16(19),
	5894: uint16(1),
	5895: uint16(anon_sym_SLASH_SLASH),
	5896: uint16(21),
	5897: uint16(1),
	5898: uint16(anon_sym_SLASH_STAR),
	5899: uint16(632),
	5900: uint16(2),
	5901: uint16(anon_sym_RBRACE),
	5902: uint16(anon_sym_COMMA),
	5903: uint16(203),
	5904: uint16(2),
	5905: uint16(sym_line_comment),
	5906: uint16(sym_block_comment),
	5907: uint16(5),
	5908: uint16(19),
	5909: uint16(1),
	5910: uint16(anon_sym_SLASH_SLASH),
	5911: uint16(21),
	5912: uint16(1),
	5913: uint16(anon_sym_SLASH_STAR),
	5914: uint16(701),
	5915: uint16(1),
	5916: uint16(anon_sym_feature),
	5917: uint16(340),
	5918: uint16(1),
	5919: uint16(sym__feature_field),
	5920: uint16(204),
	5921: uint16(2),
	5922: uint16(sym_line_comment),
	5923: uint16(sym_block_comment),
	5924: uint16(5),
	5925: uint16(19),
	5926: uint16(1),
	5927: uint16(anon_sym_SLASH_SLASH),
	5928: uint16(21),
	5929: uint16(1),
	5930: uint16(anon_sym_SLASH_STAR),
	5931: uint16(703),
	5932: uint16(1),
	5933: uint16(anon_sym_version),
	5934: uint16(326),
	5935: uint16(1),
	5936: uint16(sym__version_field),
	5937: uint16(205),
	5938: uint16(2),
	5939: uint16(sym_line_comment),
	5940: uint16(sym_block_comment),
	5941: uint16(4),
	5942: uint16(19),
	5943: uint16(1),
	5944: uint16(anon_sym_SLASH_SLASH),
	5945: uint16(21),
	5946: uint16(1),
	5947: uint16(anon_sym_SLASH_STAR),
	5948: uint16(705),
	5949: uint16(2),
	5950: uint16(anon_sym_RBRACE),
	5951: uint16(anon_sym_COMMA),
	5952: uint16(206),
	5953: uint16(2),
	5954: uint16(sym_line_comment),
	5955: uint16(sym_block_comment),
	5956: uint16(5),
	5957: uint16(19),
	5958: uint16(1),
	5959: uint16(anon_sym_SLASH_SLASH),
	5960: uint16(21),
	5961: uint16(1),
	5962: uint16(anon_sym_SLASH_STAR),
	5963: uint16(707),
	5964: uint16(1),
	5965: uint16(sym_id),
	5966: uint16(709),
	5967: uint16(1),
	5968: uint16(anon_sym_RBRACE),
	5969: uint16(207),
	5970: uint16(2),
	5971: uint16(sym_line_comment),
	5972: uint16(sym_block_comment),
	5973: uint16(5),
	5974: uint16(19),
	5975: uint16(1),
	5976: uint16(anon_sym_SLASH_SLASH),
	5977: uint16(21),
	5978: uint16(1),
	5979: uint16(anon_sym_SLASH_STAR),
	5980: uint16(703),
	5981: uint16(1),
	5982: uint16(anon_sym_version),
	5983: uint16(297),
	5984: uint16(1),
	5985: uint16(sym__version_field),
	5986: uint16(208),
	5987: uint16(2),
	5988: uint16(sym_line_comment),
	5989: uint16(sym_block_comment),
	5990: uint16(4),
	5991: uint16(19),
	5992: uint16(1),
	5993: uint16(anon_sym_SLASH_SLASH),
	5994: uint16(21),
	5995: uint16(1),
	5996: uint16(anon_sym_SLASH_STAR),
	5997: uint16(644),
	5998: uint16(2),
	5999: uint16(anon_sym_RBRACE),
	6000: uint16(anon_sym_COMMA),
	6001: uint16(209),
	6002: uint16(2),
	6003: uint16(sym_line_comment),
	6004: uint16(sym_block_comment),
	6005: uint16(5),
	6006: uint16(19),
	6007: uint16(1),
	6008: uint16(anon_sym_SLASH_SLASH),
	6009: uint16(21),
	6010: uint16(1),
	6011: uint16(anon_sym_SLASH_STAR),
	6012: uint16(711),
	6013: uint16(1),
	6014: uint16(anon_sym_SEMI),
	6015: uint16(713),
	6016: uint16(1),
	6017: uint16(anon_sym_as),
	6018: uint16(210),
	6019: uint16(2),
	6020: uint16(sym_line_comment),
	6021: uint16(sym_block_comment),
	6022: uint16(5),
	6023: uint16(19),
	6024: uint16(1),
	6025: uint16(anon_sym_SLASH_SLASH),
	6026: uint16(21),
	6027: uint16(1),
	6028: uint16(anon_sym_SLASH_STAR),
	6029: uint16(715),
	6030: uint16(1),
	6031: uint16(sym_id),
	6032: uint16(254),
	6033: uint16(1),
	6034: uint16(sym__enum_cases),
	6035: uint16(211),
	6036: uint16(2),
	6037: uint16(sym_line_comment),
	6038: uint16(sym_block_comment),
	6039: uint16(4),
	6040: uint16(19),
	6041: uint16(1),
	6042: uint16(anon_sym_SLASH_SLASH),
	6043: uint16(21),
	6044: uint16(1),
	6045: uint16(anon_sym_SLASH_STAR),
	6046: uint16(717),
	6047: uint16(2),
	6048: uint16(anon_sym_RBRACE),
	6049: uint16(anon_sym_COMMA),
	6050: uint16(212),
	6051: uint16(2),
	6052: uint16(sym_line_comment),
	6053: uint16(sym_block_comment),
	6054: uint16(5),
	6055: uint16(19),
	6056: uint16(1),
	6057: uint16(anon_sym_SLASH_SLASH),
	6058: uint16(21),
	6059: uint16(1),
	6060: uint16(anon_sym_SLASH_STAR),
	6061: uint16(719),
	6062: uint16(1),
	6063: uint16(sym_id),
	6064: uint16(721),
	6065: uint16(1),
	6066: uint16(anon_sym_RBRACE),
	6067: uint16(213),
	6068: uint16(2),
	6069: uint16(sym_line_comment),
	6070: uint16(sym_block_comment),
	6071: uint16(5),
	6072: uint16(19),
	6073: uint16(1),
	6074: uint16(anon_sym_SLASH_SLASH),
	6075: uint16(21),
	6076: uint16(1),
	6077: uint16(anon_sym_SLASH_STAR),
	6078: uint16(723),
	6079: uint16(1),
	6080: uint16(anon_sym_LPAREN),
	6081: uint16(170),
	6082: uint16(1),
	6083: uint16(sym_param_list),
	6084: uint16(214),
	6085: uint16(2),
	6086: uint16(sym_line_comment),
	6087: uint16(sym_block_comment),
	6088: uint16(5),
	6089: uint16(19),
	6090: uint16(1),
	6091: uint16(anon_sym_SLASH_SLASH),
	6092: uint16(21),
	6093: uint16(1),
	6094: uint16(anon_sym_SLASH_STAR),
	6095: uint16(602),
	6096: uint16(1),
	6097: uint16(anon_sym_RBRACE),
	6098: uint16(707),
	6099: uint16(1),
	6100: uint16(sym_id),
	6101: uint16(215),
	6102: uint16(2),
	6103: uint16(sym_line_comment),
	6104: uint16(sym_block_comment),
	6105: uint16(5),
	6106: uint16(19),
	6107: uint16(1),
	6108: uint16(anon_sym_SLASH_SLASH),
	6109: uint16(21),
	6110: uint16(1),
	6111: uint16(anon_sym_SLASH_STAR),
	6112: uint16(725),
	6113: uint16(1),
	6114: uint16(anon_sym_LBRACE),
	6115: uint16(56),
	6116: uint16(1),
	6117: uint16(sym__record_body),
	6118: uint16(216),
	6119: uint16(2),
	6120: uint16(sym_line_comment),
	6121: uint16(sym_block_comment),
	6122: uint16(5),
	6123: uint16(19),
	6124: uint16(1),
	6125: uint16(anon_sym_SLASH_SLASH),
	6126: uint16(21),
	6127: uint16(1),
	6128: uint16(anon_sym_SLASH_STAR),
	6129: uint16(727),
	6130: uint16(1),
	6131: uint16(anon_sym_LBRACE),
	6132: uint16(57),
	6133: uint16(1),
	6134: uint16(sym__include_names_body),
	6135: uint16(217),
	6136: uint16(2),
	6137: uint16(sym_line_comment),
	6138: uint16(sym_block_comment),
	6139: uint16(5),
	6140: uint16(19),
	6141: uint16(1),
	6142: uint16(anon_sym_SLASH_SLASH),
	6143: uint16(21),
	6144: uint16(1),
	6145: uint16(anon_sym_SLASH_STAR),
	6146: uint16(729),
	6147: uint16(1),
	6148: uint16(sym_id),
	6149: uint16(328),
	6150: uint16(1),
	6151: uint16(sym__flags_fields),
	6152: uint16(218),
	6153: uint16(2),
	6154: uint16(sym_line_comment),
	6155: uint16(sym_block_comment),
	6156: uint16(5),
	6157: uint16(19),
	6158: uint16(1),
	6159: uint16(anon_sym_SLASH_SLASH),
	6160: uint16(21),
	6161: uint16(1),
	6162: uint16(anon_sym_SLASH_STAR),
	6163: uint16(731),
	6164: uint16(1),
	6165: uint16(anon_sym_LBRACE),
	6166: uint16(733),
	6167: uint16(1),
	6168: uint16(anon_sym_SEMI),
	6169: uint16(219),
	6170: uint16(2),
	6171: uint16(sym_line_comment),
	6172: uint16(sym_block_comment),
	6173: uint16(4),
	6174: uint16(19),
	6175: uint16(1),
	6176: uint16(anon_sym_SLASH_SLASH),
	6177: uint16(21),
	6178: uint16(1),
	6179: uint16(anon_sym_SLASH_STAR),
	6180: uint16(735),
	6181: uint16(2),
	6182: uint16(anon_sym_LBRACE),
	6183: uint16(anon_sym_SEMI),
	6184: uint16(220),
	6185: uint16(2),
	6186: uint16(sym_line_comment),
	6187: uint16(sym_block_comment),
	6188: uint16(5),
	6189: uint16(19),
	6190: uint16(1),
	6191: uint16(anon_sym_SLASH_SLASH),
	6192: uint16(21),
	6193: uint16(1),
	6194: uint16(anon_sym_SLASH_STAR),
	6195: uint16(737),
	6196: uint16(1),
	6197: uint16(anon_sym_LBRACE),
	6198: uint16(51),
	6199: uint16(1),
	6200: uint16(sym__variant_body),
	6201: uint16(221),
	6202: uint16(2),
	6203: uint16(sym_line_comment),
	6204: uint16(sym_block_comment),
	6205: uint16(4),
	6206: uint16(19),
	6207: uint16(1),
	6208: uint16(anon_sym_SLASH_SLASH),
	6209: uint16(21),
	6210: uint16(1),
	6211: uint16(anon_sym_SLASH_STAR),
	6212: uint16(739),
	6213: uint16(2),
	6214: uint16(anon_sym_RBRACE),
	6215: uint16(anon_sym_COMMA),
	6216: uint16(222),
	6217: uint16(2),
	6218: uint16(sym_line_comment),
	6219: uint16(sym_block_comment),
	6220: uint16(4),
	6221: uint16(19),
	6222: uint16(1),
	6223: uint16(anon_sym_SLASH_SLASH),
	6224: uint16(21),
	6225: uint16(1),
	6226: uint16(anon_sym_SLASH_STAR),
	6227: uint16(658),
	6228: uint16(2),
	6229: uint16(anon_sym_RBRACE),
	6230: uint16(anon_sym_COMMA),
	6231: uint16(223),
	6232: uint16(2),
	6233: uint16(sym_line_comment),
	6234: uint16(sym_block_comment),
	6235: uint16(5),
	6236: uint16(19),
	6237: uint16(1),
	6238: uint16(anon_sym_SLASH_SLASH),
	6239: uint16(21),
	6240: uint16(1),
	6241: uint16(anon_sym_SLASH_STAR),
	6242: uint16(741),
	6243: uint16(1),
	6244: uint16(anon_sym_LBRACE),
	6245: uint16(52),
	6246: uint16(1),
	6247: uint16(sym__enum_body),
	6248: uint16(224),
	6249: uint16(2),
	6250: uint16(sym_line_comment),
	6251: uint16(sym_block_comment),
	6252: uint16(5),
	6253: uint16(19),
	6254: uint16(1),
	6255: uint16(anon_sym_SLASH_SLASH),
	6256: uint16(21),
	6257: uint16(1),
	6258: uint16(anon_sym_SLASH_STAR),
	6259: uint16(743),
	6260: uint16(1),
	6261: uint16(anon_sym_export),
	6262: uint16(745),
	6263: uint16(1),
	6264: uint16(anon_sym_import),
	6265: uint16(225),
	6266: uint16(2),
	6267: uint16(sym_line_comment),
	6268: uint16(sym_block_comment),
	6269: uint16(4),
	6270: uint16(19),
	6271: uint16(1),
	6272: uint16(anon_sym_SLASH_SLASH),
	6273: uint16(21),
	6274: uint16(1),
	6275: uint16(anon_sym_SLASH_STAR),
	6276: uint16(747),
	6277: uint16(2),
	6278: uint16(anon_sym_RBRACE),
	6279: uint16(anon_sym_COMMA),
	6280: uint16(226),
	6281: uint16(2),
	6282: uint16(sym_line_comment),
	6283: uint16(sym_block_comment),
	6284: uint16(4),
	6285: uint16(19),
	6286: uint16(1),
	6287: uint16(anon_sym_SLASH_SLASH),
	6288: uint16(21),
	6289: uint16(1),
	6290: uint16(anon_sym_SLASH_STAR),
	6291: uint16(665),
	6292: uint16(2),
	6293: uint16(anon_sym_RBRACE),
	6294: uint16(anon_sym_COMMA),
	6295: uint16(227),
	6296: uint16(2),
	6297: uint16(sym_line_comment),
	6298: uint16(sym_block_comment),
	6299: uint16(5),
	6300: uint16(19),
	6301: uint16(1),
	6302: uint16(anon_sym_SLASH_SLASH),
	6303: uint16(21),
	6304: uint16(1),
	6305: uint16(anon_sym_SLASH_STAR),
	6306: uint16(723),
	6307: uint16(1),
	6308: uint16(anon_sym_LPAREN),
	6309: uint16(159),
	6310: uint16(1),
	6311: uint16(sym_param_list),
	6312: uint16(228),
	6313: uint16(2),
	6314: uint16(sym_line_comment),
	6315: uint16(sym_block_comment),
	6316: uint16(5),
	6317: uint16(19),
	6318: uint16(1),
	6319: uint16(anon_sym_SLASH_SLASH),
	6320: uint16(21),
	6321: uint16(1),
	6322: uint16(anon_sym_SLASH_STAR),
	6323: uint16(610),
	6324: uint16(1),
	6325: uint16(anon_sym_RBRACE),
	6326: uint16(719),
	6327: uint16(1),
	6328: uint16(sym_id),
	6329: uint16(229),
	6330: uint16(2),
	6331: uint16(sym_line_comment),
	6332: uint16(sym_block_comment),
	6333: uint16(5),
	6334: uint16(19),
	6335: uint16(1),
	6336: uint16(anon_sym_SLASH_SLASH),
	6337: uint16(21),
	6338: uint16(1),
	6339: uint16(anon_sym_SLASH_STAR),
	6340: uint16(723),
	6341: uint16(1),
	6342: uint16(anon_sym_LPAREN),
	6343: uint16(171),
	6344: uint16(1),
	6345: uint16(sym_param_list),
	6346: uint16(230),
	6347: uint16(2),
	6348: uint16(sym_line_comment),
	6349: uint16(sym_block_comment),
	6350: uint16(5),
	6351: uint16(19),
	6352: uint16(1),
	6353: uint16(anon_sym_SLASH_SLASH),
	6354: uint16(21),
	6355: uint16(1),
	6356: uint16(anon_sym_SLASH_STAR),
	6357: uint16(749),
	6358: uint16(1),
	6359: uint16(anon_sym_STAR_SLASH),
	6360: uint16(751),
	6361: uint16(1),
	6362: uint16(sym__block_comment_content),
	6363: uint16(231),
	6364: uint16(2),
	6365: uint16(sym_line_comment),
	6366: uint16(sym_block_comment),
	6367: uint16(5),
	6368: uint16(19),
	6369: uint16(1),
	6370: uint16(anon_sym_SLASH_SLASH),
	6371: uint16(21),
	6372: uint16(1),
	6373: uint16(anon_sym_SLASH_STAR),
	6374: uint16(753),
	6375: uint16(1),
	6376: uint16(anon_sym_LBRACE),
	6377: uint16(95),
	6378: uint16(1),
	6379: uint16(sym__world_body),
	6380: uint16(232),
	6381: uint16(2),
	6382: uint16(sym_line_comment),
	6383: uint16(sym_block_comment),
	6384: uint16(5),
	6385: uint16(19),
	6386: uint16(1),
	6387: uint16(anon_sym_SLASH_SLASH),
	6388: uint16(21),
	6389: uint16(1),
	6390: uint16(anon_sym_SLASH_STAR),
	6391: uint16(753),
	6392: uint16(1),
	6393: uint16(anon_sym_LBRACE),
	6394: uint16(94),
	6395: uint16(1),
	6396: uint16(sym__world_body),
	6397: uint16(233),
	6398: uint16(2),
	6399: uint16(sym_line_comment),
	6400: uint16(sym_block_comment),
	6401: uint16(5),
	6402: uint16(19),
	6403: uint16(1),
	6404: uint16(anon_sym_SLASH_SLASH),
	6405: uint16(21),
	6406: uint16(1),
	6407: uint16(anon_sym_SLASH_STAR),
	6408: uint16(689),
	6409: uint16(1),
	6410: uint16(anon_sym_LBRACE),
	6411: uint16(91),
	6412: uint16(1),
	6413: uint16(sym__interface_body),
	6414: uint16(234),
	6415: uint16(2),
	6416: uint16(sym_line_comment),
	6417: uint16(sym_block_comment),
	6418: uint16(5),
	6419: uint16(19),
	6420: uint16(1),
	6421: uint16(anon_sym_SLASH_SLASH),
	6422: uint16(21),
	6423: uint16(1),
	6424: uint16(anon_sym_SLASH_STAR),
	6425: uint16(539),
	6426: uint16(1),
	6427: uint16(sym_id),
	6428: uint16(203),
	6429: uint16(1),
	6430: uint16(sym_record_field),
	6431: uint16(235),
	6432: uint16(2),
	6433: uint16(sym_line_comment),
	6434: uint16(sym_block_comment),
	6435: uint16(4),
	6436: uint16(19),
	6437: uint16(1),
	6438: uint16(anon_sym_SLASH_SLASH),
	6439: uint16(21),
	6440: uint16(1),
	6441: uint16(anon_sym_SLASH_STAR),
	6442: uint16(755),
	6443: uint16(2),
	6444: uint16(anon_sym_RBRACE),
	6445: uint16(anon_sym_COMMA),
	6446: uint16(236),
	6447: uint16(2),
	6448: uint16(sym_line_comment),
	6449: uint16(sym_block_comment),
	6450: uint16(5),
	6451: uint16(19),
	6452: uint16(1),
	6453: uint16(anon_sym_SLASH_SLASH),
	6454: uint16(21),
	6455: uint16(1),
	6456: uint16(anon_sym_SLASH_STAR),
	6457: uint16(541),
	6458: uint16(1),
	6459: uint16(sym_id),
	6460: uint16(209),
	6461: uint16(1),
	6462: uint16(sym_variant_case),
	6463: uint16(237),
	6464: uint16(2),
	6465: uint16(sym_line_comment),
	6466: uint16(sym_block_comment),
	6467: uint16(4),
	6468: uint16(19),
	6469: uint16(1),
	6470: uint16(anon_sym_SLASH_SLASH),
	6471: uint16(21),
	6472: uint16(1),
	6473: uint16(anon_sym_SLASH_STAR),
	6474: uint16(757),
	6475: uint16(2),
	6476: uint16(anon_sym_COMMA),
	6477: uint16(anon_sym_RPAREN),
	6478: uint16(238),
	6479: uint16(2),
	6480: uint16(sym_line_comment),
	6481: uint16(sym_block_comment),
	6482: uint16(4),
	6483: uint16(19),
	6484: uint16(1),
	6485: uint16(anon_sym_SLASH_SLASH),
	6486: uint16(21),
	6487: uint16(1),
	6488: uint16(anon_sym_SLASH_STAR),
	6489: uint16(677),
	6490: uint16(2),
	6491: uint16(anon_sym_COMMA),
	6492: uint16(anon_sym_RPAREN),
	6493: uint16(239),
	6494: uint16(2),
	6495: uint16(sym_line_comment),
	6496: uint16(sym_block_comment),
	6497: uint16(5),
	6498: uint16(19),
	6499: uint16(1),
	6500: uint16(anon_sym_SLASH_SLASH),
	6501: uint16(21),
	6502: uint16(1),
	6503: uint16(anon_sym_SLASH_STAR),
	6504: uint16(448),
	6505: uint16(1),
	6506: uint16(anon_sym_SEMI),
	6507: uint16(759),
	6508: uint16(1),
	6509: uint16(anon_sym_COLON),
	6510: uint16(240),
	6511: uint16(2),
	6512: uint16(sym_line_comment),
	6513: uint16(sym_block_comment),
	6514: uint16(5),
	6515: uint16(19),
	6516: uint16(1),
	6517: uint16(anon_sym_SLASH_SLASH),
	6518: uint16(21),
	6519: uint16(1),
	6520: uint16(anon_sym_SLASH_STAR),
	6521: uint16(448),
	6522: uint16(1),
	6523: uint16(anon_sym_SEMI),
	6524: uint16(761),
	6525: uint16(1),
	6526: uint16(anon_sym_COLON),
	6527: uint16(241),
	6528: uint16(2),
	6529: uint16(sym_line_comment),
	6530: uint16(sym_block_comment),
	6531: uint16(4),
	6532: uint16(19),
	6533: uint16(1),
	6534: uint16(anon_sym_SLASH_SLASH),
	6535: uint16(21),
	6536: uint16(1),
	6537: uint16(anon_sym_SLASH_STAR),
	6538: uint16(394),
	6539: uint16(2),
	6540: uint16(anon_sym_LBRACE),
	6541: uint16(anon_sym_SEMI),
	6542: uint16(242),
	6543: uint16(2),
	6544: uint16(sym_line_comment),
	6545: uint16(sym_block_comment),
	6546: uint16(5),
	6547: uint16(19),
	6548: uint16(1),
	6549: uint16(anon_sym_SLASH_SLASH),
	6550: uint16(21),
	6551: uint16(1),
	6552: uint16(anon_sym_SLASH_STAR),
	6553: uint16(553),
	6554: uint16(1),
	6555: uint16(sym_id),
	6556: uint16(227),
	6557: uint16(1),
	6558: uint16(sym_include_names_item),
	6559: uint16(243),
	6560: uint16(2),
	6561: uint16(sym_line_comment),
	6562: uint16(sym_block_comment),
	6563: uint16(4),
	6564: uint16(19),
	6565: uint16(1),
	6566: uint16(anon_sym_SLASH_SLASH),
	6567: uint16(21),
	6568: uint16(1),
	6569: uint16(anon_sym_SLASH_STAR),
	6570: uint16(537),
	6571: uint16(2),
	6572: uint16(anon_sym_COMMA),
	6573: uint16(anon_sym_GT),
	6574: uint16(244),
	6575: uint16(2),
	6576: uint16(sym_line_comment),
	6577: uint16(sym_block_comment),
	6578: uint16(4),
	6579: uint16(19),
	6580: uint16(1),
	6581: uint16(anon_sym_SLASH_SLASH),
	6582: uint16(21),
	6583: uint16(1),
	6584: uint16(anon_sym_SLASH_STAR),
	6585: uint16(763),
	6586: uint16(2),
	6587: uint16(anon_sym_SEMI),
	6588: uint16(anon_sym_DASH_GT),
	6589: uint16(245),
	6590: uint16(2),
	6591: uint16(sym_line_comment),
	6592: uint16(sym_block_comment),
	6593: uint16(5),
	6594: uint16(19),
	6595: uint16(1),
	6596: uint16(anon_sym_SLASH_SLASH),
	6597: uint16(21),
	6598: uint16(1),
	6599: uint16(anon_sym_SLASH_STAR),
	6600: uint16(480),
	6601: uint16(1),
	6602: uint16(sym_id),
	6603: uint16(239),
	6604: uint16(1),
	6605: uint16(sym_named_type),
	6606: uint16(246),
	6607: uint16(2),
	6608: uint16(sym_line_comment),
	6609: uint16(sym_block_comment),
	6610: uint16(4),
	6611: uint16(19),
	6612: uint16(1),
	6613: uint16(anon_sym_SLASH_SLASH),
	6614: uint16(21),
	6615: uint16(1),
	6616: uint16(anon_sym_SLASH_STAR),
	6617: uint16(765),
	6618: uint16(2),
	6619: uint16(anon_sym_SEMI),
	6620: uint16(anon_sym_DASH_GT),
	6621: uint16(247),
	6622: uint16(2),
	6623: uint16(sym_line_comment),
	6624: uint16(sym_block_comment),
	6625: uint16(4),
	6626: uint16(19),
	6627: uint16(1),
	6628: uint16(anon_sym_SLASH_SLASH),
	6629: uint16(21),
	6630: uint16(1),
	6631: uint16(anon_sym_SLASH_STAR),
	6632: uint16(767),
	6633: uint16(1),
	6634: uint16(anon_sym_SEMI),
	6635: uint16(248),
	6636: uint16(2),
	6637: uint16(sym_line_comment),
	6638: uint16(sym_block_comment),
	6639: uint16(4),
	6640: uint16(19),
	6641: uint16(1),
	6642: uint16(anon_sym_SLASH_SLASH),
	6643: uint16(21),
	6644: uint16(1),
	6645: uint16(anon_sym_SLASH_STAR),
	6646: uint16(769),
	6647: uint16(1),
	6648: uint16(anon_sym_RBRACE),
	6649: uint16(249),
	6650: uint16(2),
	6651: uint16(sym_line_comment),
	6652: uint16(sym_block_comment),
	6653: uint16(4),
	6654: uint16(19),
	6655: uint16(1),
	6656: uint16(anon_sym_SLASH_SLASH),
	6657: uint16(21),
	6658: uint16(1),
	6659: uint16(anon_sym_SLASH_STAR),
	6660: uint16(771),
	6661: uint16(1),
	6662: uint16(anon_sym_SEMI),
	6663: uint16(250),
	6664: uint16(2),
	6665: uint16(sym_line_comment),
	6666: uint16(sym_block_comment),
	6667: uint16(4),
	6668: uint16(19),
	6669: uint16(1),
	6670: uint16(anon_sym_SLASH_SLASH),
	6671: uint16(21),
	6672: uint16(1),
	6673: uint16(anon_sym_SLASH_STAR),
	6674: uint16(773),
	6675: uint16(1),
	6676: uint16(anon_sym_COLON),
	6677: uint16(251),
	6678: uint16(2),
	6679: uint16(sym_line_comment),
	6680: uint16(sym_block_comment),
	6681: uint16(4),
	6682: uint16(19),
	6683: uint16(1),
	6684: uint16(anon_sym_SLASH_SLASH),
	6685: uint16(21),
	6686: uint16(1),
	6687: uint16(anon_sym_SLASH_STAR),
	6688: uint16(775),
	6689: uint16(1),
	6690: uint16(anon_sym_LPAREN),
	6691: uint16(252),
	6692: uint16(2),
	6693: uint16(sym_line_comment),
	6694: uint16(sym_block_comment),
	6695: uint16(4),
	6696: uint16(19),
	6697: uint16(1),
	6698: uint16(anon_sym_SLASH_SLASH),
	6699: uint16(21),
	6700: uint16(1),
	6701: uint16(anon_sym_SLASH_STAR),
	6702: uint16(777),
	6703: uint16(1),
	6704: uint16(anon_sym_RPAREN),
	6705: uint16(253),
	6706: uint16(2),
	6707: uint16(sym_line_comment),
	6708: uint16(sym_block_comment),
	6709: uint16(4),
	6710: uint16(19),
	6711: uint16(1),
	6712: uint16(anon_sym_SLASH_SLASH),
	6713: uint16(21),
	6714: uint16(1),
	6715: uint16(anon_sym_SLASH_STAR),
	6716: uint16(779),
	6717: uint16(1),
	6718: uint16(anon_sym_RBRACE),
	6719: uint16(254),
	6720: uint16(2),
	6721: uint16(sym_line_comment),
	6722: uint16(sym_block_comment),
	6723: uint16(4),
	6724: uint16(19),
	6725: uint16(1),
	6726: uint16(anon_sym_SLASH_SLASH),
	6727: uint16(21),
	6728: uint16(1),
	6729: uint16(anon_sym_SLASH_STAR),
	6730: uint16(781),
	6731: uint16(1),
	6732: uint16(anon_sym_SEMI),
	6733: uint16(255),
	6734: uint16(2),
	6735: uint16(sym_line_comment),
	6736: uint16(sym_block_comment),
	6737: uint16(4),
	6738: uint16(19),
	6739: uint16(1),
	6740: uint16(anon_sym_SLASH_SLASH),
	6741: uint16(21),
	6742: uint16(1),
	6743: uint16(anon_sym_SLASH_STAR),
	6744: uint16(616),
	6745: uint16(1),
	6746: uint16(anon_sym_SEMI),
	6747: uint16(256),
	6748: uint16(2),
	6749: uint16(sym_line_comment),
	6750: uint16(sym_block_comment),
	6751: uint16(4),
	6752: uint16(19),
	6753: uint16(1),
	6754: uint16(anon_sym_SLASH_SLASH),
	6755: uint16(21),
	6756: uint16(1),
	6757: uint16(anon_sym_SLASH_STAR),
	6758: uint16(783),
	6759: uint16(1),
	6760: uint16(sym_id),
	6761: uint16(257),
	6762: uint16(2),
	6763: uint16(sym_line_comment),
	6764: uint16(sym_block_comment),
	6765: uint16(4),
	6766: uint16(19),
	6767: uint16(1),
	6768: uint16(anon_sym_SLASH_SLASH),
	6769: uint16(21),
	6770: uint16(1),
	6771: uint16(anon_sym_SLASH_STAR),
	6772: uint16(785),
	6773: uint16(1),
	6774: uint16(anon_sym_SEMI),
	6775: uint16(258),
	6776: uint16(2),
	6777: uint16(sym_line_comment),
	6778: uint16(sym_block_comment),
	6779: uint16(4),
	6780: uint16(19),
	6781: uint16(1),
	6782: uint16(anon_sym_SLASH_SLASH),
	6783: uint16(21),
	6784: uint16(1),
	6785: uint16(anon_sym_SLASH_STAR),
	6786: uint16(787),
	6787: uint16(1),
	6788: uint16(anon_sym_RPAREN),
	6789: uint16(259),
	6790: uint16(2),
	6791: uint16(sym_line_comment),
	6792: uint16(sym_block_comment),
	6793: uint16(4),
	6794: uint16(19),
	6795: uint16(1),
	6796: uint16(anon_sym_SLASH_SLASH),
	6797: uint16(21),
	6798: uint16(1),
	6799: uint16(anon_sym_SLASH_STAR),
	6800: uint16(488),
	6801: uint16(1),
	6802: uint16(anon_sym_external_DASHid),
	6803: uint16(260),
	6804: uint16(2),
	6805: uint16(sym_line_comment),
	6806: uint16(sym_block_comment),
	6807: uint16(4),
	6808: uint16(19),
	6809: uint16(1),
	6810: uint16(anon_sym_SLASH_SLASH),
	6811: uint16(21),
	6812: uint16(1),
	6813: uint16(anon_sym_SLASH_STAR),
	6814: uint16(789),
	6815: uint16(1),
	6816: uint16(sym_id),
	6817: uint16(261),
	6818: uint16(2),
	6819: uint16(sym_line_comment),
	6820: uint16(sym_block_comment),
	6821: uint16(4),
	6822: uint16(19),
	6823: uint16(1),
	6824: uint16(anon_sym_SLASH_SLASH),
	6825: uint16(21),
	6826: uint16(1),
	6827: uint16(anon_sym_SLASH_STAR),
	6828: uint16(791),
	6829: uint16(1),
	6830: uint16(anon_sym_EQ),
	6831: uint16(262),
	6832: uint16(2),
	6833: uint16(sym_line_comment),
	6834: uint16(sym_block_comment),
	6835: uint16(4),
	6836: uint16(19),
	6837: uint16(1),
	6838: uint16(anon_sym_SLASH_SLASH),
	6839: uint16(21),
	6840: uint16(1),
	6841: uint16(anon_sym_SLASH_STAR),
	6842: uint16(793),
	6843: uint16(1),
	6844: uint16(anon_sym_COLON),
	6845: uint16(263),
	6846: uint16(2),
	6847: uint16(sym_line_comment),
	6848: uint16(sym_block_comment),
	6849: uint16(4),
	6850: uint16(19),
	6851: uint16(1),
	6852: uint16(anon_sym_SLASH_SLASH),
	6853: uint16(21),
	6854: uint16(1),
	6855: uint16(anon_sym_SLASH_STAR),
	6856: uint16(795),
	6857: uint16(1),
	6858: uint16(anon_sym_LPAREN),
	6859: uint16(264),
	6860: uint16(2),
	6861: uint16(sym_line_comment),
	6862: uint16(sym_block_comment),
	6863: uint16(4),
	6864: uint16(19),
	6865: uint16(1),
	6866: uint16(anon_sym_SLASH_SLASH),
	6867: uint16(21),
	6868: uint16(1),
	6869: uint16(anon_sym_SLASH_STAR),
	6870: uint16(797),
	6871: uint16(1),
	6872: uint16(anon_sym_SEMI),
	6873: uint16(265),
	6874: uint16(2),
	6875: uint16(sym_line_comment),
	6876: uint16(sym_block_comment),
	6877: uint16(4),
	6878: uint16(19),
	6879: uint16(1),
	6880: uint16(anon_sym_SLASH_SLASH),
	6881: uint16(21),
	6882: uint16(1),
	6883: uint16(anon_sym_SLASH_STAR),
	6884: uint16(799),
	6885: uint16(1),
	6886: uint16(anon_sym_GT),
	6887: uint16(266),
	6888: uint16(2),
	6889: uint16(sym_line_comment),
	6890: uint16(sym_block_comment),
	6891: uint16(4),
	6892: uint16(19),
	6893: uint16(1),
	6894: uint16(anon_sym_SLASH_SLASH),
	6895: uint16(21),
	6896: uint16(1),
	6897: uint16(anon_sym_SLASH_STAR),
	6898: uint16(801),
	6899: uint16(1),
	6900: uint16(sym_id),
	6901: uint16(267),
	6902: uint16(2),
	6903: uint16(sym_line_comment),
	6904: uint16(sym_block_comment),
	6905: uint16(4),
	6906: uint16(19),
	6907: uint16(1),
	6908: uint16(anon_sym_SLASH_SLASH),
	6909: uint16(21),
	6910: uint16(1),
	6911: uint16(anon_sym_SLASH_STAR),
	6912: uint16(803),
	6913: uint16(1),
	6914: uint16(anon_sym_GT),
	6915: uint16(268),
	6916: uint16(2),
	6917: uint16(sym_line_comment),
	6918: uint16(sym_block_comment),
	6919: uint16(4),
	6920: uint16(19),
	6921: uint16(1),
	6922: uint16(anon_sym_SLASH_SLASH),
	6923: uint16(21),
	6924: uint16(1),
	6925: uint16(anon_sym_SLASH_STAR),
	6926: uint16(805),
	6927: uint16(1),
	6928: uint16(anon_sym_COMMA),
	6929: uint16(269),
	6930: uint16(2),
	6931: uint16(sym_line_comment),
	6932: uint16(sym_block_comment),
	6933: uint16(4),
	6934: uint16(19),
	6935: uint16(1),
	6936: uint16(anon_sym_SLASH_SLASH),
	6937: uint16(21),
	6938: uint16(1),
	6939: uint16(anon_sym_SLASH_STAR),
	6940: uint16(807),
	6941: uint16(1),
	6942: uint16(anon_sym_COMMA),
	6943: uint16(270),
	6944: uint16(2),
	6945: uint16(sym_line_comment),
	6946: uint16(sym_block_comment),
	6947: uint16(4),
	6948: uint16(19),
	6949: uint16(1),
	6950: uint16(anon_sym_SLASH_SLASH),
	6951: uint16(21),
	6952: uint16(1),
	6953: uint16(anon_sym_SLASH_STAR),
	6954: uint16(695),
	6955: uint16(1),
	6956: uint16(anon_sym_COMMA),
	6957: uint16(271),
	6958: uint16(2),
	6959: uint16(sym_line_comment),
	6960: uint16(sym_block_comment),
	6961: uint16(4),
	6962: uint16(19),
	6963: uint16(1),
	6964: uint16(anon_sym_SLASH_SLASH),
	6965: uint16(21),
	6966: uint16(1),
	6967: uint16(anon_sym_SLASH_STAR),
	6968: uint16(809),
	6969: uint16(1),
	6970: uint16(sym__valid_semver),
	6971: uint16(272),
	6972: uint16(2),
	6973: uint16(sym_line_comment),
	6974: uint16(sym_block_comment),
	6975: uint16(4),
	6976: uint16(19),
	6977: uint16(1),
	6978: uint16(anon_sym_SLASH_SLASH),
	6979: uint16(21),
	6980: uint16(1),
	6981: uint16(anon_sym_SLASH_STAR),
	6982: uint16(811),
	6983: uint16(1),
	6984: uint16(anon_sym_GT),
	6985: uint16(273),
	6986: uint16(2),
	6987: uint16(sym_line_comment),
	6988: uint16(sym_block_comment),
	6989: uint16(4),
	6990: uint16(19),
	6991: uint16(1),
	6992: uint16(anon_sym_SLASH_SLASH),
	6993: uint16(21),
	6994: uint16(1),
	6995: uint16(anon_sym_SLASH_STAR),
	6996: uint16(813),
	6997: uint16(1),
	6998: uint16(anon_sym_GT),
	6999: uint16(274),
	7000: uint16(2),
	7001: uint16(sym_line_comment),
	7002: uint16(sym_block_comment),
	7003: uint16(4),
	7004: uint16(19),
	7005: uint16(1),
	7006: uint16(anon_sym_SLASH_SLASH),
	7007: uint16(21),
	7008: uint16(1),
	7009: uint16(anon_sym_SLASH_STAR),
	7010: uint16(815),
	7011: uint16(1),
	7012: uint16(anon_sym_GT),
	7013: uint16(275),
	7014: uint16(2),
	7015: uint16(sym_line_comment),
	7016: uint16(sym_block_comment),
	7017: uint16(4),
	7018: uint16(19),
	7019: uint16(1),
	7020: uint16(anon_sym_SLASH_SLASH),
	7021: uint16(21),
	7022: uint16(1),
	7023: uint16(anon_sym_SLASH_STAR),
	7024: uint16(817),
	7025: uint16(1),
	7026: uint16(sym_id),
	7027: uint16(276),
	7028: uint16(2),
	7029: uint16(sym_line_comment),
	7030: uint16(sym_block_comment),
	7031: uint16(4),
	7032: uint16(19),
	7033: uint16(1),
	7034: uint16(anon_sym_SLASH_SLASH),
	7035: uint16(21),
	7036: uint16(1),
	7037: uint16(anon_sym_SLASH_STAR),
	7038: uint16(819),
	7039: uint16(1),
	7040: uint16(anon_sym_SEMI),
	7041: uint16(277),
	7042: uint16(2),
	7043: uint16(sym_line_comment),
	7044: uint16(sym_block_comment),
	7045: uint16(4),
	7046: uint16(19),
	7047: uint16(1),
	7048: uint16(anon_sym_SLASH_SLASH),
	7049: uint16(21),
	7050: uint16(1),
	7051: uint16(anon_sym_SLASH_STAR),
	7052: uint16(821),
	7053: uint16(1),
	7054: uint16(anon_sym_COLON),
	7055: uint16(278),
	7056: uint16(2),
	7057: uint16(sym_line_comment),
	7058: uint16(sym_block_comment),
	7059: uint16(4),
	7060: uint16(19),
	7061: uint16(1),
	7062: uint16(anon_sym_SLASH_SLASH),
	7063: uint16(21),
	7064: uint16(1),
	7065: uint16(anon_sym_SLASH_STAR),
	7066: uint16(823),
	7067: uint16(1),
	7068: uint16(anon_sym_SEMI),
	7069: uint16(279),
	7070: uint16(2),
	7071: uint16(sym_line_comment),
	7072: uint16(sym_block_comment),
	7073: uint16(4),
	7074: uint16(19),
	7075: uint16(1),
	7076: uint16(anon_sym_SLASH_SLASH),
	7077: uint16(21),
	7078: uint16(1),
	7079: uint16(anon_sym_SLASH_STAR),
	7080: uint16(379),
	7081: uint16(1),
	7082: uint16(anon_sym_COLON),
	7083: uint16(280),
	7084: uint16(2),
	7085: uint16(sym_line_comment),
	7086: uint16(sym_block_comment),
	7087: uint16(4),
	7088: uint16(19),
	7089: uint16(1),
	7090: uint16(anon_sym_SLASH_SLASH),
	7091: uint16(21),
	7092: uint16(1),
	7093: uint16(anon_sym_SLASH_STAR),
	7094: uint16(825),
	7095: uint16(1),
	7096: uint16(sym_id),
	7097: uint16(281),
	7098: uint16(2),
	7099: uint16(sym_line_comment),
	7100: uint16(sym_block_comment),
	7101: uint16(4),
	7102: uint16(19),
	7103: uint16(1),
	7104: uint16(anon_sym_SLASH_SLASH),
	7105: uint16(21),
	7106: uint16(1),
	7107: uint16(anon_sym_SLASH_STAR),
	7108: uint16(827),
	7109: uint16(1),
	7110: uint16(anon_sym_RPAREN),
	7111: uint16(282),
	7112: uint16(2),
	7113: uint16(sym_line_comment),
	7114: uint16(sym_block_comment),
	7115: uint16(4),
	7116: uint16(19),
	7117: uint16(1),
	7118: uint16(anon_sym_SLASH_SLASH),
	7119: uint16(21),
	7120: uint16(1),
	7121: uint16(anon_sym_SLASH_STAR),
	7122: uint16(829),
	7123: uint16(1),
	7124: uint16(anon_sym_COLON),
	7125: uint16(283),
	7126: uint16(2),
	7127: uint16(sym_line_comment),
	7128: uint16(sym_block_comment),
	7129: uint16(4),
	7130: uint16(19),
	7131: uint16(1),
	7132: uint16(anon_sym_SLASH_SLASH),
	7133: uint16(21),
	7134: uint16(1),
	7135: uint16(anon_sym_SLASH_STAR),
	7136: uint16(831),
	7137: uint16(1),
	7138: uint16(sym_id),
	7139: uint16(284),
	7140: uint16(2),
	7141: uint16(sym_line_comment),
	7142: uint16(sym_block_comment),
	7143: uint16(4),
	7144: uint16(19),
	7145: uint16(1),
	7146: uint16(anon_sym_SLASH_SLASH),
	7147: uint16(21),
	7148: uint16(1),
	7149: uint16(anon_sym_SLASH_STAR),
	7150: uint16(833),
	7151: uint16(1),
	7152: uint16(anon_sym_RBRACE),
	7153: uint16(285),
	7154: uint16(2),
	7155: uint16(sym_line_comment),
	7156: uint16(sym_block_comment),
	7157: uint16(4),
	7158: uint16(19),
	7159: uint16(1),
	7160: uint16(anon_sym_SLASH_SLASH),
	7161: uint16(21),
	7162: uint16(1),
	7163: uint16(anon_sym_SLASH_STAR),
	7164: uint16(835),
	7165: uint16(1),
	7166: uint16(sym_id),
	7167: uint16(286),
	7168: uint16(2),
	7169: uint16(sym_line_comment),
	7170: uint16(sym_block_comment),
	7171: uint16(4),
	7172: uint16(19),
	7173: uint16(1),
	7174: uint16(anon_sym_SLASH_SLASH),
	7175: uint16(21),
	7176: uint16(1),
	7177: uint16(anon_sym_SLASH_STAR),
	7178: uint16(837),
	7179: uint16(1),
	7180: uint16(anon_sym_STAR_SLASH),
	7181: uint16(287),
	7182: uint16(2),
	7183: uint16(sym_line_comment),
	7184: uint16(sym_block_comment),
	7185: uint16(4),
	7186: uint16(19),
	7187: uint16(1),
	7188: uint16(anon_sym_SLASH_SLASH),
	7189: uint16(21),
	7190: uint16(1),
	7191: uint16(anon_sym_SLASH_STAR),
	7192: uint16(839),
	7193: uint16(1),
	7194: uint16(anon_sym_as),
	7195: uint16(288),
	7196: uint16(2),
	7197: uint16(sym_line_comment),
	7198: uint16(sym_block_comment),
	7199: uint16(4),
	7200: uint16(19),
	7201: uint16(1),
	7202: uint16(anon_sym_SLASH_SLASH),
	7203: uint16(21),
	7204: uint16(1),
	7205: uint16(anon_sym_SLASH_STAR),
	7206: uint16(841),
	7207: uint16(1),
	7208: uint16(anon_sym_SEMI),
	7209: uint16(289),
	7210: uint16(2),
	7211: uint16(sym_line_comment),
	7212: uint16(sym_block_comment),
	7213: uint16(4),
	7214: uint16(19),
	7215: uint16(1),
	7216: uint16(anon_sym_SLASH_SLASH),
	7217: uint16(21),
	7218: uint16(1),
	7219: uint16(anon_sym_SLASH_STAR),
	7220: uint16(843),
	7221: uint16(1),
	7222: uint16(anon_sym_SEMI),
	7223: uint16(290),
	7224: uint16(2),
	7225: uint16(sym_line_comment),
	7226: uint16(sym_block_comment),
	7227: uint16(4),
	7228: uint16(19),
	7229: uint16(1),
	7230: uint16(anon_sym_SLASH_SLASH),
	7231: uint16(21),
	7232: uint16(1),
	7233: uint16(anon_sym_SLASH_STAR),
	7234: uint16(845),
	7235: uint16(1),
	7236: uint16(anon_sym_RBRACE),
	7237: uint16(291),
	7238: uint16(2),
	7239: uint16(sym_line_comment),
	7240: uint16(sym_block_comment),
	7241: uint16(4),
	7242: uint16(19),
	7243: uint16(1),
	7244: uint16(anon_sym_SLASH_SLASH),
	7245: uint16(21),
	7246: uint16(1),
	7247: uint16(anon_sym_SLASH_STAR),
	7248: uint16(847),
	7249: uint16(1),
	7250: uint16(sym_id),
	7251: uint16(292),
	7252: uint16(2),
	7253: uint16(sym_line_comment),
	7254: uint16(sym_block_comment),
	7255: uint16(4),
	7256: uint16(19),
	7257: uint16(1),
	7258: uint16(anon_sym_SLASH_SLASH),
	7259: uint16(21),
	7260: uint16(1),
	7261: uint16(anon_sym_SLASH_STAR),
	7262: uint16(849),
	7263: uint16(1),
	7264: uint16(anon_sym_LT),
	7265: uint16(293),
	7266: uint16(2),
	7267: uint16(sym_line_comment),
	7268: uint16(sym_block_comment),
	7269: uint16(4),
	7270: uint16(19),
	7271: uint16(1),
	7272: uint16(anon_sym_SLASH_SLASH),
	7273: uint16(21),
	7274: uint16(1),
	7275: uint16(anon_sym_SLASH_STAR),
	7276: uint16(851),
	7277: uint16(1),
	7278: uint16(anon_sym_LT),
	7279: uint16(294),
	7280: uint16(2),
	7281: uint16(sym_line_comment),
	7282: uint16(sym_block_comment),
	7283: uint16(4),
	7284: uint16(19),
	7285: uint16(1),
	7286: uint16(anon_sym_SLASH_SLASH),
	7287: uint16(21),
	7288: uint16(1),
	7289: uint16(anon_sym_SLASH_STAR),
	7290: uint16(853),
	7291: uint16(1),
	7292: uint16(anon_sym_LT),
	7293: uint16(295),
	7294: uint16(2),
	7295: uint16(sym_line_comment),
	7296: uint16(sym_block_comment),
	7297: uint16(4),
	7298: uint16(19),
	7299: uint16(1),
	7300: uint16(anon_sym_SLASH_SLASH),
	7301: uint16(21),
	7302: uint16(1),
	7303: uint16(anon_sym_SLASH_STAR),
	7304: uint16(855),
	7305: uint16(1),
	7306: uint16(anon_sym_SEMI),
	7307: uint16(296),
	7308: uint16(2),
	7309: uint16(sym_line_comment),
	7310: uint16(sym_block_comment),
	7311: uint16(4),
	7312: uint16(19),
	7313: uint16(1),
	7314: uint16(anon_sym_SLASH_SLASH),
	7315: uint16(21),
	7316: uint16(1),
	7317: uint16(anon_sym_SLASH_STAR),
	7318: uint16(857),
	7319: uint16(1),
	7320: uint16(anon_sym_RPAREN),
	7321: uint16(297),
	7322: uint16(2),
	7323: uint16(sym_line_comment),
	7324: uint16(sym_block_comment),
	7325: uint16(4),
	7326: uint16(19),
	7327: uint16(1),
	7328: uint16(anon_sym_SLASH_SLASH),
	7329: uint16(21),
	7330: uint16(1),
	7331: uint16(anon_sym_SLASH_STAR),
	7332: uint16(859),
	7333: uint16(1),
	7334: uint16(sym_id),
	7335: uint16(298),
	7336: uint16(2),
	7337: uint16(sym_line_comment),
	7338: uint16(sym_block_comment),
	7339: uint16(4),
	7340: uint16(19),
	7341: uint16(1),
	7342: uint16(anon_sym_SLASH_SLASH),
	7343: uint16(21),
	7344: uint16(1),
	7345: uint16(anon_sym_SLASH_STAR),
	7346: uint16(861),
	7347: uint16(1),
	7348: uint16(sym_id),
	7349: uint16(299),
	7350: uint16(2),
	7351: uint16(sym_line_comment),
	7352: uint16(sym_block_comment),
	7353: uint16(4),
	7354: uint16(19),
	7355: uint16(1),
	7356: uint16(anon_sym_SLASH_SLASH),
	7357: uint16(21),
	7358: uint16(1),
	7359: uint16(anon_sym_SLASH_STAR),
	7360: uint16(863),
	7361: uint16(1),
	7362: uint16(sym_id),
	7363: uint16(300),
	7364: uint16(2),
	7365: uint16(sym_line_comment),
	7366: uint16(sym_block_comment),
	7367: uint16(4),
	7368: uint16(19),
	7369: uint16(1),
	7370: uint16(anon_sym_SLASH_SLASH),
	7371: uint16(21),
	7372: uint16(1),
	7373: uint16(anon_sym_SLASH_STAR),
	7374: uint16(865),
	7375: uint16(1),
	7376: uint16(sym__valid_semver),
	7377: uint16(301),
	7378: uint16(2),
	7379: uint16(sym_line_comment),
	7380: uint16(sym_block_comment),
	7381: uint16(4),
	7382: uint16(19),
	7383: uint16(1),
	7384: uint16(anon_sym_SLASH_SLASH),
	7385: uint16(21),
	7386: uint16(1),
	7387: uint16(anon_sym_SLASH_STAR),
	7388: uint16(867),
	7389: uint16(1),
	7390: uint16(anon_sym_LT),
	7391: uint16(302),
	7392: uint16(2),
	7393: uint16(sym_line_comment),
	7394: uint16(sym_block_comment),
	7395: uint16(4),
	7396: uint16(19),
	7397: uint16(1),
	7398: uint16(anon_sym_SLASH_SLASH),
	7399: uint16(21),
	7400: uint16(1),
	7401: uint16(anon_sym_SLASH_STAR),
	7402: uint16(869),
	7403: uint16(1),
	7404: uint16(sym_id),
	7405: uint16(303),
	7406: uint16(2),
	7407: uint16(sym_line_comment),
	7408: uint16(sym_block_comment),
	7409: uint16(4),
	7410: uint16(19),
	7411: uint16(1),
	7412: uint16(anon_sym_SLASH_SLASH),
	7413: uint16(21),
	7414: uint16(1),
	7415: uint16(anon_sym_SLASH_STAR),
	7416: uint16(871),
	7417: uint16(1),
	7418: uint16(sym_id),
	7419: uint16(304),
	7420: uint16(2),
	7421: uint16(sym_line_comment),
	7422: uint16(sym_block_comment),
	7423: uint16(4),
	7424: uint16(19),
	7425: uint16(1),
	7426: uint16(anon_sym_SLASH_SLASH),
	7427: uint16(21),
	7428: uint16(1),
	7429: uint16(anon_sym_SLASH_STAR),
	7430: uint16(749),
	7431: uint16(1),
	7432: uint16(anon_sym_STAR_SLASH),
	7433: uint16(305),
	7434: uint16(2),
	7435: uint16(sym_line_comment),
	7436: uint16(sym_block_comment),
	7437: uint16(4),
	7438: uint16(19),
	7439: uint16(1),
	7440: uint16(anon_sym_SLASH_SLASH),
	7441: uint16(21),
	7442: uint16(1),
	7443: uint16(anon_sym_SLASH_STAR),
	7444: uint16(873),
	7445: uint16(1),
	7446: uint16(sym_id),
	7447: uint16(306),
	7448: uint16(2),
	7449: uint16(sym_line_comment),
	7450: uint16(sym_block_comment),
	7451: uint16(4),
	7452: uint16(19),
	7453: uint16(1),
	7454: uint16(anon_sym_SLASH_SLASH),
	7455: uint16(21),
	7456: uint16(1),
	7457: uint16(anon_sym_SLASH_STAR),
	7458: uint16(875),
	7459: uint16(1),
	7460: uint16(anon_sym_LT),
	7461: uint16(307),
	7462: uint16(2),
	7463: uint16(sym_line_comment),
	7464: uint16(sym_block_comment),
	7465: uint16(4),
	7466: uint16(19),
	7467: uint16(1),
	7468: uint16(anon_sym_SLASH_SLASH),
	7469: uint16(21),
	7470: uint16(1),
	7471: uint16(anon_sym_SLASH_STAR),
	7472: uint16(877),
	7473: uint16(1),
	7474: uint16(sym_uint),
	7475: uint16(308),
	7476: uint16(2),
	7477: uint16(sym_line_comment),
	7478: uint16(sym_block_comment),
	7479: uint16(4),
	7480: uint16(19),
	7481: uint16(1),
	7482: uint16(anon_sym_SLASH_SLASH),
	7483: uint16(21),
	7484: uint16(1),
	7485: uint16(anon_sym_SLASH_STAR),
	7486: uint16(879),
	7487: uint16(1),
	7488: uint16(sym_id),
	7489: uint16(309),
	7490: uint16(2),
	7491: uint16(sym_line_comment),
	7492: uint16(sym_block_comment),
	7493: uint16(4),
	7494: uint16(19),
	7495: uint16(1),
	7496: uint16(anon_sym_SLASH_SLASH),
	7497: uint16(21),
	7498: uint16(1),
	7499: uint16(anon_sym_SLASH_STAR),
	7500: uint16(881),
	7501: uint16(1),
	7502: uint16(sym_id),
	7503: uint16(310),
	7504: uint16(2),
	7505: uint16(sym_line_comment),
	7506: uint16(sym_block_comment),
	7507: uint16(4),
	7508: uint16(19),
	7509: uint16(1),
	7510: uint16(anon_sym_SLASH_SLASH),
	7511: uint16(21),
	7512: uint16(1),
	7513: uint16(anon_sym_SLASH_STAR),
	7514: uint16(883),
	7515: uint16(1),
	7516: uint16(anon_sym_LPAREN),
	7517: uint16(311),
	7518: uint16(2),
	7519: uint16(sym_line_comment),
	7520: uint16(sym_block_comment),
	7521: uint16(4),
	7522: uint16(19),
	7523: uint16(1),
	7524: uint16(anon_sym_SLASH_SLASH),
	7525: uint16(21),
	7526: uint16(1),
	7527: uint16(anon_sym_SLASH_STAR),
	7528: uint16(885),
	7529: uint16(1),
	7530: uint16(sym_id),
	7531: uint16(312),
	7532: uint16(2),
	7533: uint16(sym_line_comment),
	7534: uint16(sym_block_comment),
	7535: uint16(4),
	7536: uint16(19),
	7537: uint16(1),
	7538: uint16(anon_sym_SLASH_SLASH),
	7539: uint16(21),
	7540: uint16(1),
	7541: uint16(anon_sym_SLASH_STAR),
	7542: uint16(887),
	7543: uint16(1),
	7544: uint16(sym_id),
	7545: uint16(313),
	7546: uint16(2),
	7547: uint16(sym_line_comment),
	7548: uint16(sym_block_comment),
	7549: uint16(4),
	7550: uint16(19),
	7551: uint16(1),
	7552: uint16(anon_sym_SLASH_SLASH),
	7553: uint16(21),
	7554: uint16(1),
	7555: uint16(anon_sym_SLASH_STAR),
	7556: uint16(889),
	7557: uint16(1),
	7559: uint16(314),
	7560: uint16(2),
	7561: uint16(sym_line_comment),
	7562: uint16(sym_block_comment),
	7563: uint16(4),
	7564: uint16(19),
	7565: uint16(1),
	7566: uint16(anon_sym_SLASH_SLASH),
	7567: uint16(21),
	7568: uint16(1),
	7569: uint16(anon_sym_SLASH_STAR),
	7570: uint16(891),
	7571: uint16(1),
	7572: uint16(anon_sym_SEMI),
	7573: uint16(315),
	7574: uint16(2),
	7575: uint16(sym_line_comment),
	7576: uint16(sym_block_comment),
	7577: uint16(4),
	7578: uint16(19),
	7579: uint16(1),
	7580: uint16(anon_sym_SLASH_SLASH),
	7581: uint16(21),
	7582: uint16(1),
	7583: uint16(anon_sym_SLASH_STAR),
	7584: uint16(707),
	7585: uint16(1),
	7586: uint16(sym_id),
	7587: uint16(316),
	7588: uint16(2),
	7589: uint16(sym_line_comment),
	7590: uint16(sym_block_comment),
	7591: uint16(4),
	7592: uint16(3),
	7593: uint16(1),
	7594: uint16(anon_sym_SLASH_SLASH),
	7595: uint16(5),
	7596: uint16(1),
	7597: uint16(anon_sym_SLASH_STAR),
	7598: uint16(893),
	7599: uint16(1),
	7600: uint16(aux_sym_line_comment_token1),
	7601: uint16(317),
	7602: uint16(2),
	7603: uint16(sym_line_comment),
	7604: uint16(sym_block_comment),
	7605: uint16(4),
	7606: uint16(19),
	7607: uint16(1),
	7608: uint16(anon_sym_SLASH_SLASH),
	7609: uint16(21),
	7610: uint16(1),
	7611: uint16(anon_sym_SLASH_STAR),
	7612: uint16(895),
	7613: uint16(1),
	7614: uint16(anon_sym_COLON),
	7615: uint16(318),
	7616: uint16(2),
	7617: uint16(sym_line_comment),
	7618: uint16(sym_block_comment),
	7619: uint16(4),
	7620: uint16(19),
	7621: uint16(1),
	7622: uint16(anon_sym_SLASH_SLASH),
	7623: uint16(21),
	7624: uint16(1),
	7625: uint16(anon_sym_SLASH_STAR),
	7626: uint16(719),
	7627: uint16(1),
	7628: uint16(sym_id),
	7629: uint16(319),
	7630: uint16(2),
	7631: uint16(sym_line_comment),
	7632: uint16(sym_block_comment),
	7633: uint16(4),
	7634: uint16(19),
	7635: uint16(1),
	7636: uint16(anon_sym_SLASH_SLASH),
	7637: uint16(21),
	7638: uint16(1),
	7639: uint16(anon_sym_SLASH_STAR),
	7640: uint16(897),
	7641: uint16(1),
	7642: uint16(anon_sym_SEMI),
	7643: uint16(320),
	7644: uint16(2),
	7645: uint16(sym_line_comment),
	7646: uint16(sym_block_comment),
	7647: uint16(4),
	7648: uint16(19),
	7649: uint16(1),
	7650: uint16(anon_sym_SLASH_SLASH),
	7651: uint16(21),
	7652: uint16(1),
	7653: uint16(anon_sym_SLASH_STAR),
	7654: uint16(899),
	7655: uint16(1),
	7656: uint16(anon_sym_EQ),
	7657: uint16(321),
	7658: uint16(2),
	7659: uint16(sym_line_comment),
	7660: uint16(sym_block_comment),
	7661: uint16(4),
	7662: uint16(19),
	7663: uint16(1),
	7664: uint16(anon_sym_SLASH_SLASH),
	7665: uint16(21),
	7666: uint16(1),
	7667: uint16(anon_sym_SLASH_STAR),
	7668: uint16(901),
	7669: uint16(1),
	7670: uint16(anon_sym_RBRACE),
	7671: uint16(322),
	7672: uint16(2),
	7673: uint16(sym_line_comment),
	7674: uint16(sym_block_comment),
	7675: uint16(4),
	7676: uint16(19),
	7677: uint16(1),
	7678: uint16(anon_sym_SLASH_SLASH),
	7679: uint16(21),
	7680: uint16(1),
	7681: uint16(anon_sym_SLASH_STAR),
	7682: uint16(903),
	7683: uint16(1),
	7684: uint16(anon_sym_COLON),
	7685: uint16(323),
	7686: uint16(2),
	7687: uint16(sym_line_comment),
	7688: uint16(sym_block_comment),
	7689: uint16(4),
	7690: uint16(19),
	7691: uint16(1),
	7692: uint16(anon_sym_SLASH_SLASH),
	7693: uint16(21),
	7694: uint16(1),
	7695: uint16(anon_sym_SLASH_STAR),
	7696: uint16(905),
	7697: uint16(1),
	7698: uint16(anon_sym_func),
	7699: uint16(324),
	7700: uint16(2),
	7701: uint16(sym_line_comment),
	7702: uint16(sym_block_comment),
	7703: uint16(4),
	7704: uint16(19),
	7705: uint16(1),
	7706: uint16(anon_sym_SLASH_SLASH),
	7707: uint16(21),
	7708: uint16(1),
	7709: uint16(anon_sym_SLASH_STAR),
	7710: uint16(907),
	7711: uint16(1),
	7712: uint16(sym__line_doc_content),
	7713: uint16(325),
	7714: uint16(2),
	7715: uint16(sym_line_comment),
	7716: uint16(sym_block_comment),
	7717: uint16(4),
	7718: uint16(19),
	7719: uint16(1),
	7720: uint16(anon_sym_SLASH_SLASH),
	7721: uint16(21),
	7722: uint16(1),
	7723: uint16(anon_sym_SLASH_STAR),
	7724: uint16(909),
	7725: uint16(1),
	7726: uint16(anon_sym_RPAREN),
	7727: uint16(326),
	7728: uint16(2),
	7729: uint16(sym_line_comment),
	7730: uint16(sym_block_comment),
	7731: uint16(4),
	7732: uint16(19),
	7733: uint16(1),
	7734: uint16(anon_sym_SLASH_SLASH),
	7735: uint16(21),
	7736: uint16(1),
	7737: uint16(anon_sym_SLASH_STAR),
	7738: uint16(911),
	7739: uint16(1),
	7740: uint16(anon_sym_RPAREN),
	7741: uint16(327),
	7742: uint16(2),
	7743: uint16(sym_line_comment),
	7744: uint16(sym_block_comment),
	7745: uint16(4),
	7746: uint16(19),
	7747: uint16(1),
	7748: uint16(anon_sym_SLASH_SLASH),
	7749: uint16(21),
	7750: uint16(1),
	7751: uint16(anon_sym_SLASH_STAR),
	7752: uint16(913),
	7753: uint16(1),
	7754: uint16(anon_sym_RBRACE),
	7755: uint16(328),
	7756: uint16(2),
	7757: uint16(sym_line_comment),
	7758: uint16(sym_block_comment),
	7759: uint16(4),
	7760: uint16(19),
	7761: uint16(1),
	7762: uint16(anon_sym_SLASH_SLASH),
	7763: uint16(21),
	7764: uint16(1),
	7765: uint16(anon_sym_SLASH_STAR),
	7766: uint16(915),
	7767: uint16(1),
	7768: uint16(anon_sym_LPAREN),
	7769: uint16(329),
	7770: uint16(2),
	7771: uint16(sym_line_comment),
	7772: uint16(sym_block_comment),
	7773: uint16(4),
	7774: uint16(19),
	7775: uint16(1),
	7776: uint16(anon_sym_SLASH_SLASH),
	7777: uint16(21),
	7778: uint16(1),
	7779: uint16(anon_sym_SLASH_STAR),
	7780: uint16(917),
	7781: uint16(1),
	7782: uint16(anon_sym_RPAREN),
	7783: uint16(330),
	7784: uint16(2),
	7785: uint16(sym_line_comment),
	7786: uint16(sym_block_comment),
	7787: uint16(4),
	7788: uint16(19),
	7789: uint16(1),
	7790: uint16(anon_sym_SLASH_SLASH),
	7791: uint16(21),
	7792: uint16(1),
	7793: uint16(anon_sym_SLASH_STAR),
	7794: uint16(919),
	7795: uint16(1),
	7796: uint16(anon_sym_DOT),
	7797: uint16(331),
	7798: uint16(2),
	7799: uint16(sym_line_comment),
	7800: uint16(sym_block_comment),
	7801: uint16(4),
	7802: uint16(19),
	7803: uint16(1),
	7804: uint16(anon_sym_SLASH_SLASH),
	7805: uint16(21),
	7806: uint16(1),
	7807: uint16(anon_sym_SLASH_STAR),
	7808: uint16(921),
	7809: uint16(1),
	7810: uint16(anon_sym_GT),
	7811: uint16(332),
	7812: uint16(2),
	7813: uint16(sym_line_comment),
	7814: uint16(sym_block_comment),
	7815: uint16(4),
	7816: uint16(19),
	7817: uint16(1),
	7818: uint16(anon_sym_SLASH_SLASH),
	7819: uint16(21),
	7820: uint16(1),
	7821: uint16(anon_sym_SLASH_STAR),
	7822: uint16(923),
	7823: uint16(1),
	7824: uint16(anon_sym_GT),
	7825: uint16(333),
	7826: uint16(2),
	7827: uint16(sym_line_comment),
	7828: uint16(sym_block_comment),
	7829: uint16(4),
	7830: uint16(19),
	7831: uint16(1),
	7832: uint16(anon_sym_SLASH_SLASH),
	7833: uint16(21),
	7834: uint16(1),
	7835: uint16(anon_sym_SLASH_STAR),
	7836: uint16(925),
	7837: uint16(1),
	7838: uint16(anon_sym_GT),
	7839: uint16(334),
	7840: uint16(2),
	7841: uint16(sym_line_comment),
	7842: uint16(sym_block_comment),
	7843: uint16(4),
	7844: uint16(19),
	7845: uint16(1),
	7846: uint16(anon_sym_SLASH_SLASH),
	7847: uint16(21),
	7848: uint16(1),
	7849: uint16(anon_sym_SLASH_STAR),
	7850: uint16(927),
	7851: uint16(1),
	7852: uint16(anon_sym_RPAREN),
	7853: uint16(335),
	7854: uint16(2),
	7855: uint16(sym_line_comment),
	7856: uint16(sym_block_comment),
	7857: uint16(4),
	7858: uint16(19),
	7859: uint16(1),
	7860: uint16(anon_sym_SLASH_SLASH),
	7861: uint16(21),
	7862: uint16(1),
	7863: uint16(anon_sym_SLASH_STAR),
	7864: uint16(929),
	7865: uint16(1),
	7866: uint16(sym_id),
	7867: uint16(336),
	7868: uint16(2),
	7869: uint16(sym_line_comment),
	7870: uint16(sym_block_comment),
	7871: uint16(4),
	7872: uint16(19),
	7873: uint16(1),
	7874: uint16(anon_sym_SLASH_SLASH),
	7875: uint16(21),
	7876: uint16(1),
	7877: uint16(anon_sym_SLASH_STAR),
	7878: uint16(931),
	7879: uint16(1),
	7880: uint16(anon_sym_SEMI),
	7881: uint16(337),
	7882: uint16(2),
	7883: uint16(sym_line_comment),
	7884: uint16(sym_block_comment),
	7885: uint16(4),
	7886: uint16(19),
	7887: uint16(1),
	7888: uint16(anon_sym_SLASH_SLASH),
	7889: uint16(21),
	7890: uint16(1),
	7891: uint16(anon_sym_SLASH_STAR),
	7892: uint16(933),
	7893: uint16(1),
	7894: uint16(sym_string_literal),
	7895: uint16(338),
	7896: uint16(2),
	7897: uint16(sym_line_comment),
	7898: uint16(sym_block_comment),
	7899: uint16(4),
	7900: uint16(19),
	7901: uint16(1),
	7902: uint16(anon_sym_SLASH_SLASH),
	7903: uint16(21),
	7904: uint16(1),
	7905: uint16(anon_sym_SLASH_STAR),
	7906: uint16(935),
	7907: uint16(1),
	7908: uint16(anon_sym_SEMI),
	7909: uint16(339),
	7910: uint16(2),
	7911: uint16(sym_line_comment),
	7912: uint16(sym_block_comment),
	7913: uint16(4),
	7914: uint16(19),
	7915: uint16(1),
	7916: uint16(anon_sym_SLASH_SLASH),
	7917: uint16(21),
	7918: uint16(1),
	7919: uint16(anon_sym_SLASH_STAR),
	7920: uint16(937),
	7921: uint16(1),
	7922: uint16(anon_sym_RPAREN),
	7923: uint16(340),
	7924: uint16(2),
	7925: uint16(sym_line_comment),
	7926: uint16(sym_block_comment),
	7927: uint16(4),
	7928: uint16(19),
	7929: uint16(1),
	7930: uint16(anon_sym_SLASH_SLASH),
	7931: uint16(21),
	7932: uint16(1),
	7933: uint16(anon_sym_SLASH_STAR),
	7934: uint16(939),
	7935: uint16(1),
	7936: uint16(anon_sym_EQ),
	7937: uint16(341),
	7938: uint16(2),
	7939: uint16(sym_line_comment),
	7940: uint16(sym_block_comment),
	7941: uint16(4),
	7942: uint16(19),
	7943: uint16(1),
	7944: uint16(anon_sym_SLASH_SLASH),
	7945: uint16(21),
	7946: uint16(1),
	7947: uint16(anon_sym_SLASH_STAR),
	7948: uint16(941),
	7949: uint16(1),
	7950: uint16(sym_id),
	7951: uint16(342),
	7952: uint16(2),
	7953: uint16(sym_line_comment),
	7954: uint16(sym_block_comment),
	7955: uint16(1),
	7956: uint16(943),
	7957: uint16(1),
	7959: uint16(1),
	7960: uint16(945),
	7961: uint16(1),
	7963: uint16(1),
	7964: uint16(947),
	7965: uint16(1),
	7967: uint16(1),
	7968: uint16(949),
	7969: uint16(1),
	7971: uint16(1),
	7972: uint16(951),
	7973: uint16(1),
	7975: uint16(1),
	7976: uint16(953),
	7977: uint16(1),
}

var ts_small_parse_table_map = [347]uint32_t{
	1:   uint32(72),
	2:   uint32(144),
	3:   uint32(216),
	4:   uint32(288),
	5:   uint32(360),
	6:   uint32(429),
	7:   uint32(498),
	8:   uint32(567),
	9:   uint32(636),
	10:  uint32(705),
	11:  uint32(774),
	12:  uint32(843),
	13:  uint32(912),
	14:  uint32(981),
	15:  uint32(1050),
	16:  uint32(1119),
	17:  uint32(1199),
	18:  uint32(1281),
	19:  uint32(1363),
	20:  uint32(1436),
	21:  uint32(1508),
	22:  uint32(1582),
	23:  uint32(1656),
	24:  uint32(1721),
	25:  uint32(1778),
	26:  uint32(1837),
	27:  uint32(1875),
	28:  uint32(1904),
	29:  uint32(1933),
	30:  uint32(1970),
	31:  uint32(2014),
	32:  uint32(2062),
	33:  uint32(2110),
	34:  uint32(2154),
	35:  uint32(2200),
	36:  uint32(2229),
	37:  uint32(2258),
	38:  uint32(2287),
	39:  uint32(2316),
	40:  uint32(2345),
	41:  uint32(2373),
	42:  uint32(2401),
	43:  uint32(2429),
	44:  uint32(2457),
	45:  uint32(2485),
	46:  uint32(2513),
	47:  uint32(2541),
	48:  uint32(2569),
	49:  uint32(2597),
	50:  uint32(2625),
	51:  uint32(2653),
	52:  uint32(2683),
	53:  uint32(2711),
	54:  uint32(2739),
	55:  uint32(2767),
	56:  uint32(2792),
	57:  uint32(2817),
	58:  uint32(2842),
	59:  uint32(2867),
	60:  uint32(2892),
	61:  uint32(2917),
	62:  uint32(2942),
	63:  uint32(2967),
	64:  uint32(2992),
	65:  uint32(3017),
	66:  uint32(3042),
	67:  uint32(3067),
	68:  uint32(3092),
	69:  uint32(3118),
	70:  uint32(3143),
	71:  uint32(3166),
	72:  uint32(3191),
	73:  uint32(3216),
	74:  uint32(3241),
	75:  uint32(3277),
	76:  uint32(3313),
	77:  uint32(3349),
	78:  uint32(3381),
	79:  uint32(3417),
	80:  uint32(3443),
	81:  uint32(3474),
	82:  uint32(3509),
	83:  uint32(3542),
	84:  uint32(3575),
	85:  uint32(3610),
	86:  uint32(3640),
	87:  uint32(3662),
	88:  uint32(3682),
	89:  uint32(3704),
	90:  uint32(3724),
	91:  uint32(3744),
	92:  uint32(3764),
	93:  uint32(3784),
	94:  uint32(3804),
	95:  uint32(3824),
	96:  uint32(3844),
	97:  uint32(3864),
	98:  uint32(3883),
	99:  uint32(3902),
	100: uint32(3923),
	101: uint32(3944),
	102: uint32(3965),
	103: uint32(3984),
	104: uint32(4003),
	105: uint32(4022),
	106: uint32(4040),
	107: uint32(4058),
	108: uint32(4078),
	109: uint32(4096),
	110: uint32(4114),
	111: uint32(4132),
	112: uint32(4150),
	113: uint32(4168),
	114: uint32(4186),
	115: uint32(4204),
	116: uint32(4222),
	117: uint32(4240),
	118: uint32(4258),
	119: uint32(4276),
	120: uint32(4294),
	121: uint32(4312),
	122: uint32(4331),
	123: uint32(4354),
	124: uint32(4377),
	125: uint32(4400),
	126: uint32(4423),
	127: uint32(4442),
	128: uint32(4465),
	129: uint32(4488),
	130: uint32(4505),
	131: uint32(4524),
	132: uint32(4547),
	133: uint32(4570),
	134: uint32(4593),
	135: uint32(4616),
	136: uint32(4635),
	137: uint32(4658),
	138: uint32(4677),
	139: uint32(4700),
	140: uint32(4719),
	141: uint32(4742),
	142: uint32(4765),
	143: uint32(4783),
	144: uint32(4801),
	145: uint32(4821),
	146: uint32(4841),
	147: uint32(4861),
	148: uint32(4881),
	149: uint32(4901),
	150: uint32(4921),
	151: uint32(4941),
	152: uint32(4961),
	153: uint32(4981),
	154: uint32(5001),
	155: uint32(5021),
	156: uint32(5041),
	157: uint32(5061),
	158: uint32(5081),
	159: uint32(5099),
	160: uint32(5119),
	161: uint32(5139),
	162: uint32(5157),
	163: uint32(5177),
	164: uint32(5197),
	165: uint32(5217),
	166: uint32(5237),
	167: uint32(5257),
	168: uint32(5277),
	169: uint32(5297),
	170: uint32(5317),
	171: uint32(5337),
	172: uint32(5357),
	173: uint32(5377),
	174: uint32(5397),
	175: uint32(5417),
	176: uint32(5435),
	177: uint32(5453),
	178: uint32(5473),
	179: uint32(5491),
	180: uint32(5509),
	181: uint32(5529),
	182: uint32(5549),
	183: uint32(5569),
	184: uint32(5587),
	185: uint32(5607),
	186: uint32(5625),
	187: uint32(5645),
	188: uint32(5665),
	189: uint32(5685),
	190: uint32(5703),
	191: uint32(5723),
	192: uint32(5743),
	193: uint32(5760),
	194: uint32(5777),
	195: uint32(5794),
	196: uint32(5809),
	197: uint32(5826),
	198: uint32(5843),
	199: uint32(5860),
	200: uint32(5877),
	201: uint32(5892),
	202: uint32(5907),
	203: uint32(5924),
	204: uint32(5941),
	205: uint32(5956),
	206: uint32(5973),
	207: uint32(5990),
	208: uint32(6005),
	209: uint32(6022),
	210: uint32(6039),
	211: uint32(6054),
	212: uint32(6071),
	213: uint32(6088),
	214: uint32(6105),
	215: uint32(6122),
	216: uint32(6139),
	217: uint32(6156),
	218: uint32(6173),
	219: uint32(6188),
	220: uint32(6205),
	221: uint32(6220),
	222: uint32(6235),
	223: uint32(6252),
	224: uint32(6269),
	225: uint32(6284),
	226: uint32(6299),
	227: uint32(6316),
	228: uint32(6333),
	229: uint32(6350),
	230: uint32(6367),
	231: uint32(6384),
	232: uint32(6401),
	233: uint32(6418),
	234: uint32(6435),
	235: uint32(6450),
	236: uint32(6467),
	237: uint32(6482),
	238: uint32(6497),
	239: uint32(6514),
	240: uint32(6531),
	241: uint32(6546),
	242: uint32(6563),
	243: uint32(6578),
	244: uint32(6593),
	245: uint32(6610),
	246: uint32(6625),
	247: uint32(6639),
	248: uint32(6653),
	249: uint32(6667),
	250: uint32(6681),
	251: uint32(6695),
	252: uint32(6709),
	253: uint32(6723),
	254: uint32(6737),
	255: uint32(6751),
	256: uint32(6765),
	257: uint32(6779),
	258: uint32(6793),
	259: uint32(6807),
	260: uint32(6821),
	261: uint32(6835),
	262: uint32(6849),
	263: uint32(6863),
	264: uint32(6877),
	265: uint32(6891),
	266: uint32(6905),
	267: uint32(6919),
	268: uint32(6933),
	269: uint32(6947),
	270: uint32(6961),
	271: uint32(6975),
	272: uint32(6989),
	273: uint32(7003),
	274: uint32(7017),
	275: uint32(7031),
	276: uint32(7045),
	277: uint32(7059),
	278: uint32(7073),
	279: uint32(7087),
	280: uint32(7101),
	281: uint32(7115),
	282: uint32(7129),
	283: uint32(7143),
	284: uint32(7157),
	285: uint32(7171),
	286: uint32(7185),
	287: uint32(7199),
	288: uint32(7213),
	289: uint32(7227),
	290: uint32(7241),
	291: uint32(7255),
	292: uint32(7269),
	293: uint32(7283),
	294: uint32(7297),
	295: uint32(7311),
	296: uint32(7325),
	297: uint32(7339),
	298: uint32(7353),
	299: uint32(7367),
	300: uint32(7381),
	301: uint32(7395),
	302: uint32(7409),
	303: uint32(7423),
	304: uint32(7437),
	305: uint32(7451),
	306: uint32(7465),
	307: uint32(7479),
	308: uint32(7493),
	309: uint32(7507),
	310: uint32(7521),
	311: uint32(7535),
	312: uint32(7549),
	313: uint32(7563),
	314: uint32(7577),
	315: uint32(7591),
	316: uint32(7605),
	317: uint32(7619),
	318: uint32(7633),
	319: uint32(7647),
	320: uint32(7661),
	321: uint32(7675),
	322: uint32(7689),
	323: uint32(7703),
	324: uint32(7717),
	325: uint32(7731),
	326: uint32(7745),
	327: uint32(7759),
	328: uint32(7773),
	329: uint32(7787),
	330: uint32(7801),
	331: uint32(7815),
	332: uint32(7829),
	333: uint32(7843),
	334: uint32(7857),
	335: uint32(7871),
	336: uint32(7885),
	337: uint32(7899),
	338: uint32(7913),
	339: uint32(7927),
	340: uint32(7941),
	341: uint32(7955),
	342: uint32(7959),
	343: uint32(7963),
	344: uint32(7967),
	345: uint32(7971),
	346: uint32(7975),
}

var ts_parse_actions = [955]TSParseActionEntry{
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(158)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(149)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token: uint8(TSParseActionTypeReduce),
		Fsymbol:     uint16(sym_source_file),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(156)),
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
		Fstate: uint16(libc.Int32FromInt32(157)),
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
		Fstate: uint16(libc.Int32FromInt32(138)),
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
		Fstate: uint16(libc.Int32FromInt32(300)),
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
		Fstate: uint16(libc.Int32FromInt32(303)),
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
		Fstate: uint16(libc.Int32FromInt32(158)),
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
		Fstate: uint16(libc.Int32FromInt32(149)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(118)),
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
		Fstate: uint16(libc.Int32FromInt32(122)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(108)),
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
		Fstate: uint16(libc.Int32FromInt32(293)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(294)),
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
		Fstate: uint16(libc.Int32FromInt32(295)),
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
		Fstate: uint16(libc.Int32FromInt32(302)),
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
		Fstate: uint16(libc.Int32FromInt32(271)),
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
		Fstate: uint16(libc.Int32FromInt32(307)),
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
		Fstate: uint16(libc.Int32FromInt32(102)),
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
		Fstate: uint16(libc.Int32FromInt32(103)),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_tuple_list),
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
		Fstate: uint16(libc.Int32FromInt32(125)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_tuple_list),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__world_body_repeat1),
	})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	56: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__world_body_repeat1),
	})))),
	57: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__world_body_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(134)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym__world_body_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(140)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym__world_body_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(124)),
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
		Fsymbol:      uint16(aux_sym__world_body_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(135)),
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
		Fsymbol:      uint16(aux_sym__world_body_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(284)),
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
		Fsymbol:      uint16(aux_sym__world_body_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(286)),
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
		Fsymbol:      uint16(aux_sym__world_body_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(292)),
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
		Fsymbol:      uint16(aux_sym__world_body_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(306)),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__world_body_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(312)),
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
		Fsymbol:      uint16(aux_sym__world_body_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(313)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(93)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(134)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(124)),
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
		Fstate: uint16(libc.Int32FromInt32(135)),
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
		Fstate: uint16(libc.Int32FromInt32(284)),
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
		Fstate: uint16(libc.Int32FromInt32(286)),
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
		Fstate: uint16(libc.Int32FromInt32(292)),
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
		Fstate: uint16(libc.Int32FromInt32(306)),
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
		Fstate: uint16(libc.Int32FromInt32(312)),
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
		Fstate: uint16(libc.Int32FromInt32(313)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(2),
	}})),
	115: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__interface_body_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(278)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__interface_body_repeat1),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	120: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__interface_body_repeat1),
	})))),
	121: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(2),
	}})),
	123: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__interface_body_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(134)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	126: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__interface_body_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(284)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	129: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__interface_body_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(286)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym__interface_body_repeat1),
	})))),
	133: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(libc.Int32FromInt32(292)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym__interface_body_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(306)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym__interface_body_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(312)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym__interface_body_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(313)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(278)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(30)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(134)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(284)),
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
		Fstate: uint16(libc.Int32FromInt32(286)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(292)),
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
		Fstate: uint16(libc.Int32FromInt32(306)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(312)),
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
		Fstate: uint16(libc.Int32FromInt32(313)),
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
		Fstate: uint16(libc.Int32FromInt32(29)),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_source_file_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(156)),
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
		Fsymbol:      uint16(aux_sym_source_file_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(157)),
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
		Fsymbol:      uint16(aux_sym_source_file_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(138)),
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
		Fsymbol:      uint16(aux_sym_source_file_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(300)),
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
		Fsymbol:      uint16(aux_sym_source_file_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(303)),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_source_file),
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
		Fcount: uint8(1),
	}})),
	183: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__gate),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	185: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__gate),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(156)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__interface_body),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__interface_body),
	})))),
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
		Fcount: uint8(1),
	}})),
	192: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_resource_item),
		Fproduction_id: uint16(2),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(86)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_resource_item),
		Fproduction_id: uint16(2),
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
		Fstate: uint16(libc.Int32FromInt32(54)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(106)),
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
		Fstate: uint16(libc.Int32FromInt32(104)),
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
		Fcount:    uint8(1),
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
		Fsymbol:      uint16(aux_sym_nested_package_definition_repeat1),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_nested_package_definition_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(156)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_nested_package_definition_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(138)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_nested_package_definition_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(300)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_nested_package_definition_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(303)),
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
		Fcount: uint8(1),
	}})),
	218: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym__gate),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym__gate),
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
		Fcount: uint8(1),
	}})),
	222: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__gate_item),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__gate_item),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_unstable_gate),
		Fproduction_id: uint16(5),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_unstable_gate),
		Fproduction_id: uint16(5),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_since_gate),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_since_gate),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_deprecated_gate),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	236: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_deprecated_gate),
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
		Fcount: uint8(1),
	}})),
	238: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__enum_body),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__enum_body),
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
		Fcount: uint8(1),
	}})),
	242: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_type_item),
		Fproduction_id: uint16(12),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_type_item),
		Fproduction_id: uint16(12),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym__resource_body),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__resource_body),
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
		Fcount: uint8(1),
	}})),
	250: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__record_body),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__record_body),
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
		Fcount: uint8(1),
	}})),
	254: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__flags_body),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	256: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__flags_body),
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
		Fcount: uint8(1),
	}})),
	258: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_use_item),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	260: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_use_item),
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
		Fcount: uint8(1),
	}})),
	262: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__variant_body),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__variant_body),
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
		Fcount: uint8(1),
	}})),
	266: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_flags_items),
		Fproduction_id: uint16(2),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_flags_items),
		Fproduction_id: uint16(2),
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
		Fcount: uint8(1),
	}})),
	270: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__resource_body),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__resource_body),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_variant_items),
		Fproduction_id: uint16(2),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	276: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_variant_items),
		Fproduction_id: uint16(2),
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
		Fcount: uint8(1),
	}})),
	278: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_enum_items),
		Fproduction_id: uint16(2),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	280: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_enum_items),
		Fproduction_id: uint16(2),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_resource_item),
		Fproduction_id: uint16(2),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_resource_item),
		Fproduction_id: uint16(2),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__typedef_item),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__typedef_item),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_record_item),
		Fproduction_id: uint16(2),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_record_item),
		Fproduction_id: uint16(2),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_include_item),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__world_items),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_export_item),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_import_item),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_include_item),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_export_item),
		Fproduction_id: uint16(2),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	308: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_import_item),
		Fproduction_id: uint16(2),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym__world_body_repeat1),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__world_items),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_extern_type),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_export_item),
		Fproduction_id: uint16(4),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_import_item),
		Fproduction_id: uint16(4),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__include_names_body),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_func_item),
		Fproduction_id: uint16(9),
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
		Fsymbol:        uint16(sym_func_item),
		Fproduction_id: uint16(9),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__interface_items),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__interface_items),
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
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_external_id),
		Fproduction_id: uint16(11),
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
		Fcount: uint8(1),
	}})),
	332: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym__interface_body_repeat1),
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
		Fsymbol:      uint16(aux_sym__interface_body_repeat1),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym__interface_items),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__interface_items),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym__interface_items),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__interface_items),
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
		Fstate: uint16(libc.Int32FromInt32(109)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(199)),
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
		Fstate: uint16(libc.Int32FromInt32(324)),
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
		Fstate: uint16(libc.Int32FromInt32(228)),
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
		Fcount: uint8(2),
	}})),
	352: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__uri_head),
	})))),
	353: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
	355: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(336)),
	}})))),
	356: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(272)),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_use_path),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_decl_head_repeat2),
	})))),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_decl_head_repeat2),
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
		Fstate:      uint16(libc.Int32FromInt32(336)),
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
		Fstate: uint16(libc.Int32FromInt32(267)),
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
		Fstate: uint16(libc.Int32FromInt32(276)),
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
		Fstate: uint16(libc.Int32FromInt32(263)),
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
		Fstate: uint16(libc.Int32FromInt32(50)),
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
		Fstate: uint16(libc.Int32FromInt32(260)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(214)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_decl_head),
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
		Fstate: uint16(libc.Int32FromInt32(342)),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym__resource_body_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(263)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym__resource_body_repeat1),
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
		Fsymbol:      uint16(aux_sym__resource_body_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(260)),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym__resource_body_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(214)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	393: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	394: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	395: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_decl_head),
	})))),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_decl_head_repeat2),
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
		Fcount: uint8(1),
	}})),
	399: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_decl_head_repeat2),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_toplevel_use_item),
	})))),
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
		Fsymbol:      uint16(sym__uri_tail),
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
		Fsymbol:      uint16(sym__uri_tail),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_interface_item),
		Fproduction_id: uint16(4),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_interface_item),
		Fproduction_id: uint16(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__world_body),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	413: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_world_item),
		Fproduction_id: uint16(4),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_world_item),
		Fproduction_id: uint16(2),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	417: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_toplevel_use_item),
		Fproduction_id: uint16(6),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__package_items),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(sym__world_body),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_source_file_repeat1),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	425: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__statement),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_result),
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
		Fstate: uint16(libc.Int32FromInt32(2)),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_future),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_stream),
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
		Fstate: uint16(libc.Int32FromInt32(11)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_nested_package_definition),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_package_decl),
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
		Fsymbol:      uint16(sym_nested_package_definition),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_future),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_ty),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	449: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_use_path),
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
		Fsymbol:      uint16(sym__version),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_tuple),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_list),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_option),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_nested_package_definition_repeat1),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_result),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_handle),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	465: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_stream),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_list),
		Fproduction_id: uint16(17),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	469: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_map),
		Fproduction_id: uint16(18),
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
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_result),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	473: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__primitive_ty),
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
		Fcount: uint8(1),
	}})),
	475: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_resource_method),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_resource_method),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(251)),
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
		Fstate: uint16(libc.Int32FromInt32(248)),
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
		Fstate: uint16(libc.Int32FromInt32(160)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__use_names_list),
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
		Fstate: uint16(libc.Int32FromInt32(329)),
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
		Fstate: uint16(libc.Int32FromInt32(311)),
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
		Fstate: uint16(libc.Int32FromInt32(252)),
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
		Fstate: uint16(libc.Int32FromInt32(264)),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym__resource_body_repeat1),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym__resource_body_repeat1),
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
		Fstate: uint16(libc.Int32FromInt32(324)),
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
		Fstate: uint16(libc.Int32FromInt32(228)),
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
		Fstate: uint16(libc.Int32FromInt32(182)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_use_path),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__resource_body_repeat1),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_resource_method),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_resource_method),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_resource_method),
		Fproduction_id: uint16(9),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_resource_method),
		Fproduction_id: uint16(9),
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
		Fstate: uint16(libc.Int32FromInt32(240)),
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
		Fcount: uint8(1),
	}})),
	523: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_resource_method),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_resource_method),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__use_names_list),
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
		Fstate: uint16(libc.Int32FromInt32(245)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_variant_case),
		Fproduction_id: uint16(9),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(14)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	535: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_tuple_list_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(12)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_tuple_list_repeat1),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(144)),
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
		Fstate: uint16(libc.Int32FromInt32(84)),
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
		Fstate: uint16(libc.Int32FromInt32(347)),
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
		Fstate: uint16(libc.Int32FromInt32(305)),
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
		Fstate: uint16(libc.Int32FromInt32(231)),
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
		Fstate: uint16(libc.Int32FromInt32(129)),
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
		Fstate: uint16(libc.Int32FromInt32(288)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__record_fields),
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
		Fstate: uint16(libc.Int32FromInt32(164)),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym__flags_fields),
		Fproduction_id: uint16(8),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__variant_cases),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym__enum_cases),
		Fproduction_id: uint16(10),
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
		Fstate: uint16(libc.Int32FromInt32(229)),
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
		Fstate: uint16(libc.Int32FromInt32(280)),
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
		Fstate: uint16(libc.Int32FromInt32(317)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(325)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(346)),
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
		Fsymbol:      uint16(sym_func_type),
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
		Fstate: uint16(libc.Int32FromInt32(5)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_use_names_item),
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
		Fstate: uint16(libc.Int32FromInt32(257)),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__use_names_list),
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
		Fstate: uint16(libc.Int32FromInt32(142)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__include_names_list),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(173)),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_decl_head_repeat1),
	})))),
	597: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(libc.Int32FromInt32(280)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	598: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	599: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__record_fields),
	})))),
	600: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	601: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	602: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	603: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym__flags_fields),
		Fproduction_id: uint16(8),
	})))),
	604: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	605: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	606: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	607: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__variant_cases),
	})))),
	608: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(179)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym__enum_cases),
		Fproduction_id: uint16(10),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(213)),
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
		Fstate: uint16(libc.Int32FromInt32(123)),
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
	617: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_func_type),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__named_type_list),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__include_names_list),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_tuple_list),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__record_fields),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__record_fields_repeat1),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	635: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__record_fields_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(235)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__flags_fields_repeat1),
	})))),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__flags_fields_repeat1),
	})))),
	641: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(libc.Int32FromInt32(316)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__variant_cases),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__variant_cases_repeat1),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	647: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__variant_cases_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(237)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__enum_cases_repeat1),
	})))),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__enum_cases_repeat1),
	})))),
	653: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(libc.Int32FromInt32(319)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__named_type_list),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__use_names_list_repeat1),
	})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	661: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__use_names_list_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(192)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__include_names_list),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__include_names_list_repeat1),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__include_names_list_repeat1),
	})))),
	669: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(libc.Int32FromInt32(243)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__named_type_list),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	675: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__named_type_list_repeat1),
	})))),
	676: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	677: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	678: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__named_type_list_repeat1),
	})))),
	679: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	680: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	681: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	682: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(218)),
	}})))),
	683: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	684: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	685: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	686: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	687: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	688: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	689: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	690: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	691: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(308)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	694: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	695: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	696: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	697: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(115)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	700: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_record_field),
		Fproduction_id: uint16(13),
	})))),
	701: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	702: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	703: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	704: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(341)),
	}})))),
	705: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	706: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym__flags_fields_repeat1),
		Fproduction_id: uint16(14),
	})))),
	707: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	708: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	709: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	710: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym__flags_fields),
		Fproduction_id: uint16(8),
	})))),
	711: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	712: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	713: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	714: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	715: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	716: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	717: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	718: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym__enum_cases_repeat1),
		Fproduction_id: uint16(15),
	})))),
	719: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	720: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	721: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	722: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym__enum_cases),
		Fproduction_id: uint16(10),
	})))),
	723: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(143)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	726: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	727: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	728: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	729: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(153)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	732: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	733: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	734: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	735: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	736: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_decl_head),
	})))),
	737: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	738: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	739: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	740: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_alias_item),
		Fproduction_id: uint16(16),
	})))),
	741: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	742: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	743: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	744: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	745: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	746: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	747: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	748: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_include_names_item),
		Fproduction_id: uint16(16),
	})))),
	749: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	750: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(345)),
	}})))),
	751: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	752: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	753: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	754: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	755: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	756: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_variant_case),
		Fproduction_id: uint16(13),
	})))),
	757: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	758: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_named_type),
		Fproduction_id: uint16(13),
	})))),
	759: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	760: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	761: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	762: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	763: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	764: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_param_list),
	})))),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_param_list),
	})))),
	767: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	768: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_result_list),
	})))),
	769: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	770: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	771: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	772: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	773: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	774: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	775: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	776: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	777: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	778: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	779: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	780: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	781: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	782: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	783: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	784: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	785: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	786: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__use_names_body),
	})))),
	787: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(72)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	790: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	791: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	792: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	793: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(130)),
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
		Fstate: uint16(libc.Int32FromInt32(208)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(111)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	806: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_kt),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(16)),
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
		Fstate: uint16(libc.Int32FromInt32(110)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	812: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(107)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	818: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	819: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	820: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	821: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	822: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	823: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	824: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	825: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	826: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_decl_head_repeat1),
	})))),
	827: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	828: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	829: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	830: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	831: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	832: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	833: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	834: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	835: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	836: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	837: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	838: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(348)),
	}})))),
	839: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	840: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	841: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(137)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_func_type),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(69)),
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
		Fstate: uint16(libc.Int32FromInt32(194)),
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
		Fstate: uint16(libc.Int32FromInt32(4)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	854: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	855: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	856: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_result_list),
	})))),
	857: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(41)),
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
		Fstate: uint16(libc.Int32FromInt32(330)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	862: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(335)),
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
		Fstate: uint16(libc.Int32FromInt32(53)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(198)),
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
		Fstate: uint16(libc.Int32FromInt32(279)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(221)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(299)),
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
		Fstate: uint16(libc.Int32FromInt32(332)),
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
		Fstate: uint16(libc.Int32FromInt32(283)),
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
		Fstate: uint16(libc.Int32FromInt32(323)),
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
		Fstate: uint16(libc.Int32FromInt32(204)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(31)),
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
	890: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f uint8_t
		_ [7]byte
	}{f: uint8(TSParseActionTypeAccept)})),
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
		Fstate: uint16(libc.Int32FromInt32(43)),
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
		Fcount: uint8(1),
	}})),
	894: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	895: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	896: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	897: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	898: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	899: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	900: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	901: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	902: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	903: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	904: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	905: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	906: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	907: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	908: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	909: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	910: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	911: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	912: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	913: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	914: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	915: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	916: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	917: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	918: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym__feature_field),
		Fproduction_id: uint16(7),
	})))),
	919: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	920: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(196)),
	}})))),
	921: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	922: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	923: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	924: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	925: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	926: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	927: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	928: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__version_field),
	})))),
	929: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	930: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	931: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	932: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_result_list),
	})))),
	933: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	934: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	935: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	936: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	937: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	938: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	939: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	940: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	941: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	942: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__uri_head),
	})))),
	943: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	944: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_line_comment),
	})))),
	945: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	946: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_line_comment),
		Fproduction_id: uint16(1),
	})))),
	947: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	948: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	949: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	950: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_line_comment),
	})))),
	951: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	952: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_block_comment),
		Fproduction_id: uint16(3),
	})))),
}

type ts_external_scanner_symbol_identifiers = int32

const ts_external_token__block_comment_content = 0
const ts_external_token__block_doc_comment_marker = 1
const ts_external_token__error_sentinel = 2
const ts_external_token__line_doc_content = 3

var ts_external_scanner_symbol_map = [4]TSSymbol{
	0: uint16(sym__block_comment_content),
	1: uint16(sym__block_doc_comment_marker),
	2: uint16(sym__error_sentinel),
	3: uint16(sym__line_doc_content),
}

var ts_external_scanner_states = [5][4]uint8{
	1: {
		0: libc.BoolUint8(true1 != 0),
		1: libc.BoolUint8(true1 != 0),
		2: libc.BoolUint8(true1 != 0),
		3: libc.BoolUint8(true1 != 0),
	},
	2: {
		0: libc.BoolUint8(true1 != 0),
		1: libc.BoolUint8(true1 != 0),
	},
	3: {
		0: libc.BoolUint8(true1 != 0),
	},
	4: {
		3: libc.BoolUint8(true1 != 0),
	},
}

func tree_sitter_wit(tls *libc.TLS) (r uintptr) {
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
	Fkeyword_capture_token:     uint16(sym_id),
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
	Fname:                  __ccgo_ts + 1645,
	Fsupertype_count:       uint32(SUPERTYPE_COUNT),
	Fsupertype_symbols:     uintptr(unsafe.Pointer(&ts_supertype_symbols)),
	Fsupertype_map_slices:  uintptr(unsafe.Pointer(&ts_supertype_map_slices)),
	Fsupertype_map_entries: uintptr(unsafe.Pointer(&ts_supertype_map_entries)),
	Fmetadata: TSLanguageMetadata{
		Fmajor_version: uint8(1),
		Fminor_version: uint8(4),
	},
}

func init() {
	p := unsafe.Pointer(&language)
	*(*uintptr)(unsafe.Add(p, 144)) = __ccgo_fp(ts_lex)
	*(*uintptr)(unsafe.Add(p, 152)) = __ccgo_fp(ts_lex_keywords)
	*(*uintptr)(unsafe.Add(p, 184)) = __ccgo_fp(tree_sitter_wit_external_scanner_create)
	*(*uintptr)(unsafe.Add(p, 192)) = __ccgo_fp(tree_sitter_wit_external_scanner_destroy)
	*(*uintptr)(unsafe.Add(p, 200)) = __ccgo_fp(tree_sitter_wit_external_scanner_scan)
	*(*uintptr)(unsafe.Add(p, 208)) = __ccgo_fp(tree_sitter_wit_external_scanner_serialize)
	*(*uintptr)(unsafe.Add(p, 216)) = __ccgo_fp(tree_sitter_wit_external_scanner_deserialize)
}

func __ccgo_fp(f interface{}) uintptr {
	type iface [2]uintptr
	return (*iface)(unsafe.Pointer(&f))[1]
}

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = "end\x00id\x00{\x00}\x00:\x00/\x00@\x00package\x00;\x00use\x00as\x00string_literal\x00version\x00world\x00export\x00import\x00interface\x00include\x00with\x00,\x00async\x00func\x00(\x00)\x00->\x00.\x00type\x00=\x00record\x00flags\x00variant\x00enum\x00resource\x00static\x00constructor\x00u8\x00u16\x00u32\x00u64\x00s8\x00s16\x00s32\x00s64\x00char\x00bool\x00string\x00f32\x00f64\x00tuple\x00<\x00>\x00uint\x00list\x00option\x00map\x00result\x00_\x00borrow\x00future\x00stream\x00//\x00line_comment_token1\x00line_comment_token2\x00/*\x00*/\x00external-id\x00unstable\x00feature\x00since\x00deprecated\x00_block_comment_content\x00_block_doc_comment_marker\x00_error_sentinel\x00doc_comment\x00source_file\x00_statement\x00_package_items\x00nested_package_definition\x00_uri_head\x00_uri_tail\x00_version\x00decl_head\x00package_decl\x00toplevel_use_item\x00use_path\x00world_item\x00body\x00_world_items\x00export_item\x00import_item\x00extern_type\x00include_item\x00definitions\x00_include_names_list\x00include_names_item\x00alias_item\x00interface_item\x00_interface_items\x00_typedef_item\x00func_item\x00func_type\x00param_list\x00result_list\x00_named_type_list\x00named_type\x00use_item\x00_use_names_list\x00use_names_item\x00type_item\x00record_item\x00_record_fields\x00record_field\x00flags_items\x00_flags_fields\x00variant_items\x00_variant_cases\x00variant_case\x00enum_items\x00_enum_cases\x00resource_item\x00resource_method\x00_primitive_ty\x00ty\x00kt\x00tuple_list\x00handle\x00line_comment\x00block_comment\x00_gate\x00_gate_item\x00external_id\x00unstable_gate\x00_feature_field\x00since_gate\x00deprecated_gate\x00_version_field\x00source_file_repeat1\x00nested_package_definition_repeat1\x00decl_head_repeat1\x00decl_head_repeat2\x00_world_body_repeat1\x00_include_names_list_repeat1\x00_interface_body_repeat1\x00_named_type_list_repeat1\x00_use_names_list_repeat1\x00_record_fields_repeat1\x00_flags_fields_repeat1\x00_variant_cases_repeat1\x00_enum_cases_repeat1\x00_resource_body_repeat1\x00tuple_list_repeat1\x00enum_case\x00flags_field\x00alias\x00doc\x00key\x00name\x00path\x00size\x00value\x00wit\x00"
