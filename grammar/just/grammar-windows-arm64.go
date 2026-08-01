// Code generated for windows/arm64 by 'ccgo -ignore-unsupported-alignment -ignore-unsupported-atomic-sizes -ignore-vector-functions --winapi-no-errno -ignore-link-errors -std=gnu11 -O0 -D__extension__= -D__forceinline=static inline -D__attribute__(...)= -D__declspec(x)= -D__cdecl= -D__stdcall= -D__fastcall= -D__thiscall= -D_cdecl= -D__restrict= -D__restrict__= -D__MINGW_EXTENSION= -D_X86INTRIN_H_INCLUDED -D_X86GPRINTRIN_H_INCLUDED -D_IMMINTRIN_H_INCLUDED -D_MMINTRIN_H_INCLUDED -D_XMMINTRIN_H_INCLUDED -D_EMMINTRIN_H_INCLUDED -D_PMMINTRIN_H_INCLUDED -D_MM3DNOW_H_INCLUDED -Dfunc=func_token -Dinterface=interface_token -Dselect=select_token -Dchan=chan_token -Dgo=go_token -Dmap=map_token -Dpackage=package_token -Dtype=type_token -Dvar=var_token -Dimport=import_token -Ddefer=defer_token -Dfallthrough=fallthrough_token -Drange=range_token -I C:\a\ccgo-tree-sitter\ccgo-tree-sitter\third-party\tree-sitter-just\src -I C:\a\ccgo-tree-sitter\ccgo-tree-sitter\third-party\tree-sitter-just -I C:\a\ccgo-tree-sitter\ccgo-tree-sitter\third-party -I C:\Users\runneradmin\.cache\workspaced\sources\github\7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2\lib\include -I C:\Users\runneradmin\.cache\workspaced\sources\github\7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2\lib\src combined.c -o grammar.go', DO NOT EDIT.

//go:build windows && arm64

package grammar_just

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
const EXTERNAL_TOKEN_COUNT = 5
const FIELD_COUNT = 16
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
const MAX_ALIAS_SEQUENCE_LENGTH = 9
const MAX_RESERVED_WORD_SET_SIZE = 0
const MB_LEN_MAX = 1
const MINGW_HAS_DDK_H = 1
const MINGW_HAS_SECURE_API = 1
const PRODUCTION_ID_COUNT = 49
const PTRDIFF_MAX = "__PTRDIFF_MAX__"
const P_tmpdir = "_P_tmpdir"
const RAND_MAX = 0x7fff
const SEEK_CUR = 1
const SEEK_END = 2
const SEEK_SET = 0
const SIZE_MAX = "__SIZE_MAX__"
const STATE_COUNT = 405
const STDERR_FILENO = 2
const STDIN_FILENO = 0
const STDOUT_FILENO = 1
const SUPERTYPE_COUNT = 0
const SYMBOL_COUNT = 118
const SYS_OPEN = "_SYS_OPEN"
const TMP_MAX = 2147483647
const TMP_MAX_S = "TMP_MAX"
const TOKEN_COUNT = 60
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
const _NLSCMPERROR = 2147483647
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
const _WConst_return = "_CONST_RETURN"
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
const __FILE_NAME__ = "__FILE__"
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
const fprintf_s = "fprintf"
const func1 = "func_token"
const go1 = "go_token"
const import1 = "import_token"
const interface1 = "interface_token"
const isascii1 = "__isascii"
const iscsym = "__iscsym"
const iscsymf = "__iscsymf"
const map1 = "map_token"
const onexit_t = "_onexit_t"
const package1 = "package_token"
const pclose1 = "_pclose"
const popen1 = "_popen"
const range1 = "range_token"
const select2 = "select_token"
const strcasecmp1 = "_stricmp"
const strncasecmp = "_strnicmp"
const sys_errlist = "_sys_errlist"
const sys_nerr = "_sys_nerr"
const toascii = "__toascii"
const true1 = 1
const ts_builtin_sym_end = 0
const type1 = "type_token"
const var1 = "var_token"
const wcswcs = "wcsstr"
const wpopen = "_wpopen"

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

type wctrans_t = uint16

// Enable this for debugging
// #define DEBUG_PRINT

type TokenType = int32

const INDENT = 0
const DEDENT = 1
const NEWLINE = 2
const TEXT = 3
const ERROR_RECOVERY = 4
const TOKEN_TYPE_END = 5

func assert_valid_token(tls *libc.TLS, sym TSSymbol) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	if libc.BoolInt32(!(int32(sym) >= int32(INDENT) && int32(sym) < int32(TOKEN_TYPE_END))) != 0 {
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts, libc.VaList(bp+8, __ccgo_ts+17, int32(76)))
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts+28, libc.VaList(bp+8, int32(sym)))
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts+46, 0)
		libc.Xexit(tls, int32(1))
	}
}

type Scanner = struct {
	Fprev_indent         uint32_t
	Fadvance_brace_count uint16_t
	Fhas_seen_eof        uint8
}

// C documentation
//
//	// This function should create your scanner object. It will only be called once
//	// anytime your language is set on a parser. Often, you will want to allocate
//	// memory on the heap and return a pointer to it. If your external scanner
//	// doesn’t need to maintain any state, it’s ok to return NULL.
func tree_sitter_just_external_scanner_create(tls *libc.TLS) (r uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var ptr uintptr
	_ = ptr
	ptr = libc.Xcalloc(tls, uint64(8), uint64(1))
	if libc.BoolInt32(!(ptr != 0)) != 0 {
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts, libc.VaList(bp+8, __ccgo_ts+17, int32(91)))
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts+48, 0)
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts+46, 0)
		libc.Xexit(tls, int32(1))
	}
	return ptr
}

// C documentation
//
//	// This function should free any memory used by your scanner. It is called once
//	// when a parser is deleted or assigned a different language. It receives as an
//	// argument the same pointer that was returned from the create function. If your
//	// create function didn’t allocate any memory, this function can be a noop.
func tree_sitter_just_external_scanner_destroy(tls *libc.TLS, payload uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	if libc.BoolInt32(!(payload != 0)) != 0 {
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts, libc.VaList(bp+8, __ccgo_ts+17, int32(100)))
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts+104, 0)
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts+46, 0)
		libc.Xexit(tls, int32(1))
	}
	libc.Xfree(tls, payload)
}

// C documentation
//
//	// Serialize the state of the scanner. This is called when the parser is
//	// serialized. It receives as an argument the same pointer that was returned
//	// from the create function.
func tree_sitter_just_external_scanner_serialize(tls *libc.TLS, payload uintptr, buffer uintptr) (r uint32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	if libc.BoolInt32(!(libc.Uint64FromInt64(8) < libc.Uint64FromInt32(TREE_SITTER_SERIALIZATION_BUFFER_SIZE))) != 0 {
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts, libc.VaList(bp+8, __ccgo_ts+17, int32(109)))
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts+132, 0)
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts+46, 0)
		libc.Xexit(tls, int32(1))
	}
	libc.Xmemcpy(tls, buffer, payload, uint64(8))
	return uint32(8)
}

// C documentation
//
//	// Reconstruct a scanner from the serialized state. This is called when the
//	// parser is deserialized.
func tree_sitter_just_external_scanner_deserialize(tls *libc.TLS, payload uintptr, buffer uintptr, length uint32) {
	var ptr uintptr
	_ = ptr
	ptr = payload
	if length == uint32(0) {
		(*Scanner)(unsafe.Pointer(ptr)).Fprev_indent = uint32(0)
		(*Scanner)(unsafe.Pointer(ptr)).Fhas_seen_eof = libc.BoolUint8(false1 != 0)
		return
	}
	libc.Xmemcpy(tls, ptr, buffer, uint64(8))
}

// C documentation
//
//	// Continue and include the preceding character in the token
func advance(tls *libc.TLS, lexer uintptr) {
	(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(false1 != 0))
}

// C documentation
//
//	// Continue and discard the preceding character
func skip(tls *libc.TLS, lexer uintptr) {
	(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(true1 != 0))
}

// C documentation
//
//	// An EOF works as a dedent
func handle_eof(tls *libc.TLS, lexer uintptr, state uintptr, valid_symbols uintptr) (r uint8) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	if libc.BoolInt32(!((*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Feof})))(tls, lexer) != 0)) != 0 {
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts, libc.VaList(bp+8, __ccgo_ts+17, int32(138)))
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts+153, 0)
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts+46, 0)
		libc.Xexit(tls, int32(1))
	}
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
	if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(DEDENT))) != 0 {
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(DEDENT)
		return libc.BoolUint8(true1 != 0)
	}
	if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(NEWLINE))) != 0 {
		if (*Scanner)(unsafe.Pointer(state)).Fhas_seen_eof != 0 {
			// allow EOF to count for a single symbol. Don't return true more than
			// once, otherwise it will keep calling us thinking there are more tokens.
			return libc.BoolUint8(false1 != 0)
		}
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(NEWLINE)
		(*Scanner)(unsafe.Pointer(state)).Fhas_seen_eof = libc.BoolUint8(true1 != 0)
		return libc.BoolUint8(true1 != 0)
	}
	return libc.BoolUint8(false1 != 0)
}

// C documentation
//
//	// This function is responsible for recognizing external tokens. It should
//	// return true if a token was recognized, and false otherwise.
func tree_sitter_just_external_scanner_scan(tls *libc.TLS, payload uintptr, lexer uintptr, valid_symbols uintptr) (r uint8) {
	var advanced_once, eol_found, escape uint8
	var indent uint32_t
	var scanner uintptr
	_, _, _, _, _ = advanced_once, eol_found, escape, indent, scanner
	scanner = payload
	if (*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Feof})))(tls, lexer) != 0 {
		return handle_eof(tls, lexer, scanner, valid_symbols)
	}
	// Handle backslash escaping for newlines
	if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(NEWLINE))) != 0 {
		escape = libc.BoolUint8(false1 != 0)
		if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('\\') {
			escape = libc.BoolUint8(true1 != 0)
			skip(tls, lexer)
		}
		eol_found = libc.BoolUint8(false1 != 0)
		for libc.Xiswctype(tls, uint16((*TSLexer)(unsafe.Pointer(lexer)).Flookahead), uint16(_SPACE)) != 0 {
			if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('\n') {
				skip(tls, lexer)
				eol_found = libc.BoolUint8(true1 != 0)
				break
			}
			skip(tls, lexer)
		}
		if eol_found != 0 && !(escape != 0) {
			(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(NEWLINE)
			return libc.BoolUint8(true1 != 0)
		}
	}
	if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(INDENT))) != 0 || *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(DEDENT))) != 0 {
		for !((*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Feof})))(tls, lexer) != 0) && libc.Xisspace(tls, (*TSLexer)(unsafe.Pointer(lexer)).Flookahead) != 0 {
			switch (*TSLexer)(unsafe.Pointer(lexer)).Flookahead {
			case int32('\n'):
				if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(INDENT))) != 0 {
					return libc.BoolUint8(false1 != 0)
				}
				fallthrough
			case int32('\t'):
				fallthrough
			case int32(' '):
				skip(tls, lexer)
			default:
				return libc.BoolUint8(false1 != 0)
			}
		}
		if (*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Feof})))(tls, lexer) != 0 {
			return handle_eof(tls, lexer, scanner, valid_symbols)
		}
		indent = (*(*func(*libc.TLS, uintptr) uint32_t)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fget_column})))(tls, lexer)
		if indent > (*Scanner)(unsafe.Pointer(scanner)).Fprev_indent && *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(INDENT))) != 0 && (*Scanner)(unsafe.Pointer(scanner)).Fprev_indent == uint32(0) {
			(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(INDENT)
			(*Scanner)(unsafe.Pointer(scanner)).Fprev_indent = indent
			return libc.BoolUint8(true1 != 0)
		}
		if indent < (*Scanner)(unsafe.Pointer(scanner)).Fprev_indent && *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(DEDENT))) != 0 && indent == uint32(0) {
			(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(DEDENT)
			(*Scanner)(unsafe.Pointer(scanner)).Fprev_indent = indent
			return libc.BoolUint8(true1 != 0)
		}
	}
	if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(TEXT))) != 0 {
		if (*(*func(*libc.TLS, uintptr) uint32_t)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fget_column})))(tls, lexer) == (*Scanner)(unsafe.Pointer(scanner)).Fprev_indent && ((*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('\n') || (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('@') || (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('-')) {
			return libc.BoolUint8(false1 != 0)
		}
		advanced_once = libc.BoolUint8(false1 != 0)
		for (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('{') && int32((*Scanner)(unsafe.Pointer(scanner)).Fadvance_brace_count) > 0 && !((*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Feof})))(tls, lexer) != 0) {
			(*Scanner)(unsafe.Pointer(scanner)).Fadvance_brace_count = (*Scanner)(unsafe.Pointer(scanner)).Fadvance_brace_count - 1
			advance(tls, lexer)
			advanced_once = libc.BoolUint8(true1 != 0)
		}
		for int32(1) != 0 {
			if (*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Feof})))(tls, lexer) != 0 {
				return handle_eof(tls, lexer, scanner, valid_symbols)
			}
			for !((*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Feof})))(tls, lexer) != 0) && (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32('\n') && (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32('{') {
				// Can't start with #!
				if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('#') && !(advanced_once != 0) {
					advance(tls, lexer)
					if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('!') {
						return libc.BoolUint8(false1 != 0)
					}
				}
				advance(tls, lexer)
				advanced_once = libc.BoolUint8(true1 != 0)
			}
			if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('\n') || (*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Feof})))(tls, lexer) != 0 {
				(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
				(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(TEXT)
				if advanced_once != 0 {
					return libc.BoolUint8(true1 != 0)
				}
				if (*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Feof})))(tls, lexer) != 0 {
					return handle_eof(tls, lexer, scanner, valid_symbols)
				}
				advance(tls, lexer)
			} else {
				if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('{') {
					(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
					advance(tls, lexer)
					if (*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Feof})))(tls, lexer) != 0 || (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('\n') { // EOF without anything after {
						(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
						(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(TEXT)
						return advanced_once
					}
					if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('{') {
						advance(tls, lexer)
						for (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('{') { // more braces!
							(*Scanner)(unsafe.Pointer(scanner)).Fadvance_brace_count = (*Scanner)(unsafe.Pointer(scanner)).Fadvance_brace_count + 1
							advance(tls, lexer)
						}
						// scan till a balanced pair of }} are found, then assume it's a valid
						// interpolation
						for !((*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Feof})))(tls, lexer) != 0) && (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32('\n') {
							advance(tls, lexer)
							if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('}') {
								advance(tls, lexer)
								if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('}') {
									(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(TEXT)
									return advanced_once
								}
							}
						}
						if !(advanced_once != 0) {
							return libc.BoolUint8(false1 != 0)
						}
					}
				}
			}
		}
	}
	return libc.BoolUint8(false1 != 0)
}

/* Automatically @generated by tree-sitter v0.25.5
 * (460118b4c82318b083b4d527c9c750426730f9c0) */

type ts_symbol_identifiers = int32

const sym_identifier = 1
const anon_sym_alias = 2
const anon_sym_COLON_EQ = 3
const anon_sym_export = 4
const anon_sym_import = 5
const anon_sym_QMARK = 6
const anon_sym_mod = 7
const anon_sym_set = 8
const anon_sym_LBRACK = 9
const anon_sym_COMMA = 10
const anon_sym_RBRACK = 11
const anon_sym_shell = 12
const anon_sym_true = 13
const anon_sym_false = 14
const anon_sym_SLASH = 15
const anon_sym_PLUS = 16
const anon_sym_if = 17
const anon_sym_else = 18
const anon_sym_LBRACE = 19
const anon_sym_RBRACE = 20
const anon_sym_EQ_EQ = 21
const anon_sym_BANG_EQ = 22
const anon_sym_EQ_TILDE = 23
const anon_sym_LPAREN = 24
const anon_sym_RPAREN = 25
const anon_sym_EQ = 26
const anon_sym_COLON = 27
const anon_sym_AT = 28
const anon_sym_DOLLAR = 29
const anon_sym_STAR = 30
const anon_sym_AMP_AMP = 31
const anon_sym_AT_DASH = 32
const anon_sym_DASH_AT = 33
const anon_sym_DASH = 34
const aux_sym_shebang_token1 = 35
const aux_sym__shebang_with_lang_token1 = 36
const anon_sym_env = 37
const aux_sym__shebang_with_lang_token2 = 38
const aux_sym__shebang_with_lang_token3 = 39
const sym__opaque_shebang = 40
const aux_sym_string_token1 = 41
const anon_sym_SQUOTE_SQUOTE_SQUOTE = 42
const aux_sym__raw_string_indented_token1 = 43
const anon_sym_DQUOTE = 44
const aux_sym__string_token1 = 45
const anon_sym_DQUOTE_DQUOTE_DQUOTE = 46
const aux_sym__string_indented_token1 = 47
const sym_escape_sequence = 48
const anon_sym_BQUOTE = 49
const anon_sym_BQUOTE_BQUOTE_BQUOTE = 50
const anon_sym_LBRACE_LBRACE = 51
const anon_sym_RBRACE_RBRACE = 52
const sym_numeric_error = 53
const sym_comment = 54
const sym__indent = 55
const sym__dedent = 56
const sym__newline = 57
const sym_text = 58
const sym_error_recovery = 59
const sym_source_file = 60
const sym__item = 61
const sym_alias = 62
const sym_assignment = 63
const sym_export = 64
const sym_import = 65
const sym_module = 66
const sym_setting = 67
const sym_boolean = 68
const sym_expression = 69
const sym__expression_inner = 70
const sym_if_expression = 71
const sym_else_if_clause = 72
const sym_else_clause = 73
const sym__braced_expr = 74
const sym_condition = 75
const sym_regex_literal = 76
const sym_value = 77
const sym_function_call = 78
const sym_external_command = 79
const sym_sequence = 80
const sym_attribute_kv_argument = 81
const sym_attribute = 82
const sym_recipe = 83
const sym_recipe_header = 84
const sym_parameters = 85
const sym_parameter = 86
const sym_variadic_parameter = 87
const sym_dependencies = 88
const sym_dependency = 89
const sym_dependency_expression = 90
const sym_recipe_body = 91
const sym_recipe_line = 92
const sym_recipe_line_prefix = 93
const sym_shebang = 94
const sym__shebang_with_lang = 95
const sym_string = 96
const sym__backticked = 97
const sym__indented_backticked = 98
const sym_command_body = 99
const sym_interpolation = 100
const aux_sym_source_file_repeat1 = 101
const aux_sym_alias_repeat1 = 102
const aux_sym_setting_repeat1 = 103
const aux_sym_if_expression_repeat1 = 104
const aux_sym_sequence_repeat1 = 105
const aux_sym_attribute_repeat1 = 106
const aux_sym_attribute_repeat2 = 107
const aux_sym_parameters_repeat1 = 108
const aux_sym_dependencies_repeat1 = 109
const aux_sym_dependency_expression_repeat1 = 110
const aux_sym_recipe_body_repeat1 = 111
const aux_sym_recipe_line_repeat1 = 112
const aux_sym__shebang_with_lang_repeat1 = 113
const aux_sym__raw_string_indented_repeat1 = 114
const aux_sym__string_repeat1 = 115
const aux_sym__string_indented_repeat1 = 116
const aux_sym_command_body_repeat1 = 117
const anon_alias_sym_expression = 118
const alias_sym_language = 119

var ts_symbol_names = [120]uintptr{
	0:   __ccgo_ts + 166,
	1:   __ccgo_ts + 170,
	2:   __ccgo_ts + 181,
	3:   __ccgo_ts + 187,
	4:   __ccgo_ts + 190,
	5:   __ccgo_ts + 197,
	6:   __ccgo_ts + 204,
	7:   __ccgo_ts + 206,
	8:   __ccgo_ts + 210,
	9:   __ccgo_ts + 214,
	10:  __ccgo_ts + 216,
	11:  __ccgo_ts + 218,
	12:  __ccgo_ts + 220,
	13:  __ccgo_ts + 226,
	14:  __ccgo_ts + 231,
	15:  __ccgo_ts + 237,
	16:  __ccgo_ts + 239,
	17:  __ccgo_ts + 241,
	18:  __ccgo_ts + 244,
	19:  __ccgo_ts + 249,
	20:  __ccgo_ts + 251,
	21:  __ccgo_ts + 253,
	22:  __ccgo_ts + 256,
	23:  __ccgo_ts + 259,
	24:  __ccgo_ts + 262,
	25:  __ccgo_ts + 264,
	26:  __ccgo_ts + 266,
	27:  __ccgo_ts + 268,
	28:  __ccgo_ts + 270,
	29:  __ccgo_ts + 272,
	30:  __ccgo_ts + 274,
	31:  __ccgo_ts + 276,
	32:  __ccgo_ts + 279,
	33:  __ccgo_ts + 282,
	34:  __ccgo_ts + 285,
	35:  __ccgo_ts + 287,
	36:  __ccgo_ts + 302,
	37:  __ccgo_ts + 328,
	38:  __ccgo_ts + 332,
	39:  __ccgo_ts + 358,
	40:  __ccgo_ts + 384,
	41:  __ccgo_ts + 400,
	42:  __ccgo_ts + 414,
	43:  __ccgo_ts + 418,
	44:  __ccgo_ts + 446,
	45:  __ccgo_ts + 448,
	46:  __ccgo_ts + 463,
	47:  __ccgo_ts + 467,
	48:  __ccgo_ts + 491,
	49:  __ccgo_ts + 507,
	50:  __ccgo_ts + 509,
	51:  __ccgo_ts + 513,
	52:  __ccgo_ts + 516,
	53:  __ccgo_ts + 519,
	54:  __ccgo_ts + 533,
	55:  __ccgo_ts + 541,
	56:  __ccgo_ts + 549,
	57:  __ccgo_ts + 557,
	58:  __ccgo_ts + 566,
	59:  __ccgo_ts + 571,
	60:  __ccgo_ts + 586,
	61:  __ccgo_ts + 598,
	62:  __ccgo_ts + 181,
	63:  __ccgo_ts + 604,
	64:  __ccgo_ts + 190,
	65:  __ccgo_ts + 197,
	66:  __ccgo_ts + 615,
	67:  __ccgo_ts + 622,
	68:  __ccgo_ts + 630,
	69:  __ccgo_ts + 638,
	70:  __ccgo_ts + 649,
	71:  __ccgo_ts + 667,
	72:  __ccgo_ts + 681,
	73:  __ccgo_ts + 696,
	74:  __ccgo_ts + 708,
	75:  __ccgo_ts + 721,
	76:  __ccgo_ts + 731,
	77:  __ccgo_ts + 745,
	78:  __ccgo_ts + 751,
	79:  __ccgo_ts + 765,
	80:  __ccgo_ts + 782,
	81:  __ccgo_ts + 791,
	82:  __ccgo_ts + 813,
	83:  __ccgo_ts + 823,
	84:  __ccgo_ts + 830,
	85:  __ccgo_ts + 844,
	86:  __ccgo_ts + 855,
	87:  __ccgo_ts + 865,
	88:  __ccgo_ts + 884,
	89:  __ccgo_ts + 897,
	90:  __ccgo_ts + 908,
	91:  __ccgo_ts + 930,
	92:  __ccgo_ts + 942,
	93:  __ccgo_ts + 954,
	94:  __ccgo_ts + 973,
	95:  __ccgo_ts + 981,
	96:  __ccgo_ts + 1000,
	97:  __ccgo_ts + 1007,
	98:  __ccgo_ts + 1019,
	99:  __ccgo_ts + 1040,
	100: __ccgo_ts + 1053,
	101: __ccgo_ts + 1067,
	102: __ccgo_ts + 1087,
	103: __ccgo_ts + 1101,
	104: __ccgo_ts + 1117,
	105: __ccgo_ts + 1139,
	106: __ccgo_ts + 1156,
	107: __ccgo_ts + 1174,
	108: __ccgo_ts + 1192,
	109: __ccgo_ts + 1211,
	110: __ccgo_ts + 1232,
	111: __ccgo_ts + 1262,
	112: __ccgo_ts + 1282,
	113: __ccgo_ts + 1302,
	114: __ccgo_ts + 1329,
	115: __ccgo_ts + 1358,
	116: __ccgo_ts + 1374,
	117: __ccgo_ts + 1399,
	118: __ccgo_ts + 638,
	119: __ccgo_ts + 1420,
}

var ts_symbol_map = [120]TSSymbol{
	1:   uint16(sym_identifier),
	2:   uint16(anon_sym_alias),
	3:   uint16(anon_sym_COLON_EQ),
	4:   uint16(anon_sym_export),
	5:   uint16(anon_sym_import),
	6:   uint16(anon_sym_QMARK),
	7:   uint16(anon_sym_mod),
	8:   uint16(anon_sym_set),
	9:   uint16(anon_sym_LBRACK),
	10:  uint16(anon_sym_COMMA),
	11:  uint16(anon_sym_RBRACK),
	12:  uint16(anon_sym_shell),
	13:  uint16(anon_sym_true),
	14:  uint16(anon_sym_false),
	15:  uint16(anon_sym_SLASH),
	16:  uint16(anon_sym_PLUS),
	17:  uint16(anon_sym_if),
	18:  uint16(anon_sym_else),
	19:  uint16(anon_sym_LBRACE),
	20:  uint16(anon_sym_RBRACE),
	21:  uint16(anon_sym_EQ_EQ),
	22:  uint16(anon_sym_BANG_EQ),
	23:  uint16(anon_sym_EQ_TILDE),
	24:  uint16(anon_sym_LPAREN),
	25:  uint16(anon_sym_RPAREN),
	26:  uint16(anon_sym_EQ),
	27:  uint16(anon_sym_COLON),
	28:  uint16(anon_sym_AT),
	29:  uint16(anon_sym_DOLLAR),
	30:  uint16(anon_sym_STAR),
	31:  uint16(anon_sym_AMP_AMP),
	32:  uint16(anon_sym_AT_DASH),
	33:  uint16(anon_sym_DASH_AT),
	34:  uint16(anon_sym_DASH),
	35:  uint16(aux_sym_shebang_token1),
	36:  uint16(aux_sym__shebang_with_lang_token1),
	37:  uint16(anon_sym_env),
	38:  uint16(aux_sym__shebang_with_lang_token2),
	39:  uint16(aux_sym__shebang_with_lang_token3),
	40:  uint16(sym__opaque_shebang),
	41:  uint16(aux_sym_string_token1),
	42:  uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	43:  uint16(aux_sym__raw_string_indented_token1),
	44:  uint16(anon_sym_DQUOTE),
	45:  uint16(aux_sym__string_token1),
	46:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	47:  uint16(aux_sym__string_indented_token1),
	48:  uint16(sym_escape_sequence),
	49:  uint16(anon_sym_BQUOTE),
	50:  uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	51:  uint16(anon_sym_LBRACE_LBRACE),
	52:  uint16(anon_sym_RBRACE_RBRACE),
	53:  uint16(sym_numeric_error),
	54:  uint16(sym_comment),
	55:  uint16(sym__indent),
	56:  uint16(sym__dedent),
	57:  uint16(sym__newline),
	58:  uint16(sym_text),
	59:  uint16(sym_error_recovery),
	60:  uint16(sym_source_file),
	61:  uint16(sym__item),
	62:  uint16(sym_alias),
	63:  uint16(sym_assignment),
	64:  uint16(sym_export),
	65:  uint16(sym_import),
	66:  uint16(sym_module),
	67:  uint16(sym_setting),
	68:  uint16(sym_boolean),
	69:  uint16(sym_expression),
	70:  uint16(sym__expression_inner),
	71:  uint16(sym_if_expression),
	72:  uint16(sym_else_if_clause),
	73:  uint16(sym_else_clause),
	74:  uint16(sym__braced_expr),
	75:  uint16(sym_condition),
	76:  uint16(sym_regex_literal),
	77:  uint16(sym_value),
	78:  uint16(sym_function_call),
	79:  uint16(sym_external_command),
	80:  uint16(sym_sequence),
	81:  uint16(sym_attribute_kv_argument),
	82:  uint16(sym_attribute),
	83:  uint16(sym_recipe),
	84:  uint16(sym_recipe_header),
	85:  uint16(sym_parameters),
	86:  uint16(sym_parameter),
	87:  uint16(sym_variadic_parameter),
	88:  uint16(sym_dependencies),
	89:  uint16(sym_dependency),
	90:  uint16(sym_dependency_expression),
	91:  uint16(sym_recipe_body),
	92:  uint16(sym_recipe_line),
	93:  uint16(sym_recipe_line_prefix),
	94:  uint16(sym_shebang),
	95:  uint16(sym__shebang_with_lang),
	96:  uint16(sym_string),
	97:  uint16(sym__backticked),
	98:  uint16(sym__indented_backticked),
	99:  uint16(sym_command_body),
	100: uint16(sym_interpolation),
	101: uint16(aux_sym_source_file_repeat1),
	102: uint16(aux_sym_alias_repeat1),
	103: uint16(aux_sym_setting_repeat1),
	104: uint16(aux_sym_if_expression_repeat1),
	105: uint16(aux_sym_sequence_repeat1),
	106: uint16(aux_sym_attribute_repeat1),
	107: uint16(aux_sym_attribute_repeat2),
	108: uint16(aux_sym_parameters_repeat1),
	109: uint16(aux_sym_dependencies_repeat1),
	110: uint16(aux_sym_dependency_expression_repeat1),
	111: uint16(aux_sym_recipe_body_repeat1),
	112: uint16(aux_sym_recipe_line_repeat1),
	113: uint16(aux_sym__shebang_with_lang_repeat1),
	114: uint16(aux_sym__raw_string_indented_repeat1),
	115: uint16(aux_sym__string_repeat1),
	116: uint16(aux_sym__string_indented_repeat1),
	117: uint16(aux_sym_command_body_repeat1),
	118: uint16(anon_alias_sym_expression),
	119: uint16(alias_sym_language),
}

var ts_symbol_metadata = [120]TSSymbolMetadata{
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
	35: {},
	36: {},
	37: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	38: {},
	39: {},
	40: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	41: {},
	42: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	43: {},
	44: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	45: {},
	46: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	47: {},
	48: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
	52: {
		Fvisible: libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	56: {
		Fnamed: libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	71: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	75: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	79: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	80: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	97: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	98: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	99: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	100: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
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
	111: {},
	112: {},
	113: {},
	114: {},
	115: {},
	116: {},
	117: {},
	118: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	119: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
}

type ts_field_identifiers = int32

const field_alternative = 1
const field_argument = 2
const field_arguments = 3
const field_array = 4
const field_body = 5
const field_consequence = 6
const field_content = 7
const field_default = 8
const field_element = 9
const field_key = 10
const field_kleene = 11
const field_left = 12
const field_name = 13
const field_right = 14
const field_shebang = 15
const field_value = 16

var ts_field_names = [17]uintptr{
	0:  libc.UintptrFromInt32(0),
	1:  __ccgo_ts + 1429,
	2:  __ccgo_ts + 1441,
	3:  __ccgo_ts + 1450,
	4:  __ccgo_ts + 1460,
	5:  __ccgo_ts + 1466,
	6:  __ccgo_ts + 1471,
	7:  __ccgo_ts + 1483,
	8:  __ccgo_ts + 1491,
	9:  __ccgo_ts + 1499,
	10: __ccgo_ts + 1507,
	11: __ccgo_ts + 1511,
	12: __ccgo_ts + 1518,
	13: __ccgo_ts + 1523,
	14: __ccgo_ts + 1528,
	15: __ccgo_ts + 973,
	16: __ccgo_ts + 745,
}

var ts_field_map_slices = [49]TSMapSlice{
	1: {
		Flength: uint16(1),
	},
	2: {
		Flength: uint16(1),
	},
	3: {
		Findex:  uint16(1),
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
		Findex:  uint16(1),
		Flength: uint16(1),
	},
	8: {
		Findex:  uint16(5),
		Flength: uint16(2),
	},
	9: {
		Findex:  uint16(7),
		Flength: uint16(2),
	},
	10: {
		Findex:  uint16(9),
		Flength: uint16(2),
	},
	12: {
		Findex:  uint16(11),
		Flength: uint16(2),
	},
	13: {
		Findex:  uint16(13),
		Flength: uint16(2),
	},
	14: {
		Findex:  uint16(15),
		Flength: uint16(3),
	},
	15: {
		Findex:  uint16(18),
		Flength: uint16(1),
	},
	17: {
		Findex:  uint16(19),
		Flength: uint16(2),
	},
	19: {
		Findex:  uint16(21),
		Flength: uint16(2),
	},
	20: {
		Findex:  uint16(23),
		Flength: uint16(2),
	},
	21: {
		Findex:  uint16(25),
		Flength: uint16(4),
	},
	22: {
		Findex:  uint16(29),
		Flength: uint16(1),
	},
	23: {
		Findex:  uint16(30),
		Flength: uint16(2),
	},
	25: {
		Findex:  uint16(32),
		Flength: uint16(1),
	},
	26: {
		Findex:  uint16(33),
		Flength: uint16(3),
	},
	27: {
		Findex:  uint16(36),
		Flength: uint16(3),
	},
	28: {
		Findex:  uint16(39),
		Flength: uint16(2),
	},
	29: {
		Findex:  uint16(41),
		Flength: uint16(1),
	},
	30: {
		Findex:  uint16(42),
		Flength: uint16(1),
	},
	31: {
		Findex:  uint16(43),
		Flength: uint16(3),
	},
	32: {
		Findex:  uint16(46),
		Flength: uint16(2),
	},
	33: {
		Findex:  uint16(48),
		Flength: uint16(5),
	},
	34: {
		Findex:  uint16(53),
		Flength: uint16(2),
	},
	35: {
		Findex:  uint16(55),
		Flength: uint16(1),
	},
	36: {
		Findex:  uint16(56),
		Flength: uint16(1),
	},
	37: {
		Findex:  uint16(57),
		Flength: uint16(4),
	},
	38: {
		Findex:  uint16(61),
		Flength: uint16(2),
	},
	39: {
		Findex:  uint16(63),
		Flength: uint16(4),
	},
	40: {
		Findex:  uint16(67),
		Flength: uint16(5),
	},
	41: {
		Findex:  uint16(72),
		Flength: uint16(6),
	},
	42: {
		Findex:  uint16(78),
		Flength: uint16(7),
	},
	43: {
		Findex:  uint16(85),
		Flength: uint16(2),
	},
	44: {
		Findex:  uint16(87),
		Flength: uint16(2),
	},
	45: {
		Findex:  uint16(89),
		Flength: uint16(6),
	},
	46: {
		Findex:  uint16(95),
		Flength: uint16(8),
	},
	47: {
		Findex:  uint16(103),
		Flength: uint16(3),
	},
	48: {
		Findex:  uint16(106),
		Flength: uint16(1),
	},
}

var ts_field_map_entries = [107]TSFieldMapEntry{
	0: {
		Ffield_id: uint16(field_name),
	},
	1: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
	},
	2: {
		Ffield_id: uint16(field_kleene),
	},
	3: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(2),
	},
	4: {
		Ffield_id:    uint16(field_left),
		Fchild_index: uint8(1),
	},
	5: {
		Ffield_id:    uint16(field_left),
		Fchild_index: uint8(1),
	},
	6: {
		Ffield_id:    uint16(field_right),
		Fchild_index: uint8(3),
	},
	7: {
		Ffield_id:    uint16(field_default),
		Fchild_index: uint8(2),
	},
	8: {
		Ffield_id: uint16(field_name),
	},
	9: {
		Ffield_id:  uint16(field_argument),
		Finherited: libc.BoolUint8(true1 != 0),
	},
	10: {
		Ffield_id:    uint16(field_argument),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	11: {
		Ffield_id: uint16(field_left),
	},
	12: {
		Ffield_id:    uint16(field_right),
		Fchild_index: uint8(2),
	},
	13: {
		Ffield_id:    uint16(field_default),
		Fchild_index: uint8(3),
	},
	14: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
	},
	15: {
		Ffield_id:    uint16(field_left),
		Fchild_index: uint8(1),
	},
	16: {
		Ffield_id:    uint16(field_right),
		Fchild_index: uint8(2),
	},
	17: {
		Ffield_id:    uint16(field_right),
		Fchild_index: uint8(3),
	},
	18: {
		Ffield_id:    uint16(field_argument),
		Fchild_index: uint8(2),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	19: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(2),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	20: {
		Ffield_id:    uint16(field_consequence),
		Fchild_index: uint8(2),
	},
	21: {
		Ffield_id:    uint16(field_left),
		Fchild_index: uint8(2),
	},
	22: {
		Ffield_id:    uint16(field_right),
		Fchild_index: uint8(4),
	},
	23: {
		Ffield_id:    uint16(field_right),
		Fchild_index: uint8(3),
	},
	24: {
		Ffield_id:    uint16(field_right),
		Fchild_index: uint8(4),
	},
	25: {
		Ffield_id:    uint16(field_array),
		Fchild_index: uint8(3),
	},
	26: {
		Ffield_id:    uint16(field_array),
		Fchild_index: uint8(4),
	},
	27: {
		Ffield_id:    uint16(field_left),
		Fchild_index: uint8(1),
	},
	28: {
		Ffield_id:    uint16(field_right),
		Fchild_index: uint8(2),
	},
	29: {
		Ffield_id:    uint16(field_argument),
		Fchild_index: uint8(3),
	},
	30: {
		Ffield_id: uint16(field_key),
	},
	31: {
		Ffield_id:    uint16(field_value),
		Fchild_index: uint8(2),
	},
	32: {
		Ffield_id: uint16(field_alternative),
	},
	33: {
		Ffield_id:    uint16(field_alternative),
		Fchild_index: uint8(3),
	},
	34: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(2),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	35: {
		Ffield_id:    uint16(field_consequence),
		Fchild_index: uint8(2),
	},
	36: {
		Ffield_id:    uint16(field_alternative),
		Fchild_index: uint8(3),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	37: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(2),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	38: {
		Ffield_id:    uint16(field_consequence),
		Fchild_index: uint8(2),
	},
	39: {
		Ffield_id:    uint16(field_arguments),
		Fchild_index: uint8(2),
	},
	40: {
		Ffield_id: uint16(field_name),
	},
	41: {
		Ffield_id:    uint16(field_shebang),
		Fchild_index: uint8(1),
	},
	42: {
		Ffield_id:    uint16(field_element),
		Fchild_index: uint8(1),
	},
	43: {
		Ffield_id:    uint16(field_element),
		Fchild_index: uint8(4),
	},
	44: {
		Ffield_id:    uint16(field_right),
		Fchild_index: uint8(3),
	},
	45: {
		Ffield_id:    uint16(field_right),
		Fchild_index: uint8(5),
	},
	46: {
		Ffield_id:  uint16(field_element),
		Finherited: libc.BoolUint8(true1 != 0),
	},
	47: {
		Ffield_id:    uint16(field_element),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	48: {
		Ffield_id:    uint16(field_array),
		Fchild_index: uint8(3),
	},
	49: {
		Ffield_id:    uint16(field_array),
		Fchild_index: uint8(5),
	},
	50: {
		Ffield_id:    uint16(field_element),
		Fchild_index: uint8(4),
	},
	51: {
		Ffield_id:    uint16(field_left),
		Fchild_index: uint8(1),
	},
	52: {
		Ffield_id:    uint16(field_right),
		Fchild_index: uint8(2),
	},
	53: {
		Ffield_id:    uint16(field_argument),
		Fchild_index: uint8(3),
	},
	54: {
		Ffield_id:    uint16(field_argument),
		Fchild_index: uint8(4),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	55: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(1),
	},
	56: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	57: {
		Ffield_id:    uint16(field_alternative),
		Fchild_index: uint8(3),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	58: {
		Ffield_id:    uint16(field_alternative),
		Fchild_index: uint8(4),
	},
	59: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(2),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	60: {
		Ffield_id:    uint16(field_consequence),
		Fchild_index: uint8(2),
	},
	61: {
		Ffield_id:  uint16(field_alternative),
		Finherited: libc.BoolUint8(true1 != 0),
	},
	62: {
		Ffield_id:    uint16(field_alternative),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	63: {
		Ffield_id:    uint16(field_element),
		Fchild_index: uint8(4),
	},
	64: {
		Ffield_id:    uint16(field_element),
		Fchild_index: uint8(5),
	},
	65: {
		Ffield_id:    uint16(field_right),
		Fchild_index: uint8(3),
	},
	66: {
		Ffield_id:    uint16(field_right),
		Fchild_index: uint8(6),
	},
	67: {
		Ffield_id:    uint16(field_content),
		Fchild_index: uint8(5),
	},
	68: {
		Ffield_id:    uint16(field_element),
		Fchild_index: uint8(4),
	},
	69: {
		Ffield_id:    uint16(field_element),
		Fchild_index: uint8(5),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	70: {
		Ffield_id:    uint16(field_right),
		Fchild_index: uint8(3),
	},
	71: {
		Ffield_id:    uint16(field_right),
		Fchild_index: uint8(6),
	},
	72: {
		Ffield_id:    uint16(field_array),
		Fchild_index: uint8(3),
	},
	73: {
		Ffield_id:    uint16(field_array),
		Fchild_index: uint8(6),
	},
	74: {
		Ffield_id:    uint16(field_element),
		Fchild_index: uint8(4),
	},
	75: {
		Ffield_id:    uint16(field_element),
		Fchild_index: uint8(5),
	},
	76: {
		Ffield_id:    uint16(field_left),
		Fchild_index: uint8(1),
	},
	77: {
		Ffield_id:    uint16(field_right),
		Fchild_index: uint8(2),
	},
	78: {
		Ffield_id:    uint16(field_array),
		Fchild_index: uint8(3),
	},
	79: {
		Ffield_id:    uint16(field_array),
		Fchild_index: uint8(6),
	},
	80: {
		Ffield_id:    uint16(field_content),
		Fchild_index: uint8(5),
	},
	81: {
		Ffield_id:    uint16(field_element),
		Fchild_index: uint8(4),
	},
	82: {
		Ffield_id:    uint16(field_element),
		Fchild_index: uint8(5),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	83: {
		Ffield_id:    uint16(field_left),
		Fchild_index: uint8(1),
	},
	84: {
		Ffield_id:    uint16(field_right),
		Fchild_index: uint8(2),
	},
	85: {
		Ffield_id:    uint16(field_argument),
		Fchild_index: uint8(3),
	},
	86: {
		Ffield_id:    uint16(field_argument),
		Fchild_index: uint8(4),
	},
	87: {
		Ffield_id:    uint16(field_argument),
		Fchild_index: uint8(3),
	},
	88: {
		Ffield_id:    uint16(field_argument),
		Fchild_index: uint8(5),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	89: {
		Ffield_id:    uint16(field_content),
		Fchild_index: uint8(5),
	},
	90: {
		Ffield_id:    uint16(field_element),
		Fchild_index: uint8(4),
	},
	91: {
		Ffield_id:    uint16(field_element),
		Fchild_index: uint8(5),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	92: {
		Ffield_id:    uint16(field_element),
		Fchild_index: uint8(6),
	},
	93: {
		Ffield_id:    uint16(field_right),
		Fchild_index: uint8(3),
	},
	94: {
		Ffield_id:    uint16(field_right),
		Fchild_index: uint8(7),
	},
	95: {
		Ffield_id:    uint16(field_array),
		Fchild_index: uint8(3),
	},
	96: {
		Ffield_id:    uint16(field_array),
		Fchild_index: uint8(7),
	},
	97: {
		Ffield_id:    uint16(field_content),
		Fchild_index: uint8(5),
	},
	98: {
		Ffield_id:    uint16(field_element),
		Fchild_index: uint8(4),
	},
	99: {
		Ffield_id:    uint16(field_element),
		Fchild_index: uint8(5),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	100: {
		Ffield_id:    uint16(field_element),
		Fchild_index: uint8(6),
	},
	101: {
		Ffield_id:    uint16(field_left),
		Fchild_index: uint8(1),
	},
	102: {
		Ffield_id:    uint16(field_right),
		Fchild_index: uint8(2),
	},
	103: {
		Ffield_id:    uint16(field_argument),
		Fchild_index: uint8(3),
	},
	104: {
		Ffield_id:    uint16(field_argument),
		Fchild_index: uint8(4),
	},
	105: {
		Ffield_id:    uint16(field_argument),
		Fchild_index: uint8(6),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	106: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(3),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
}

var ts_alias_sequences = [49][9]TSSymbol{
	0: {},
	1: {
		0: uint16(sym_identifier),
	},
	7: {
		1: uint16(sym_identifier),
	},
	11: {
		1: uint16(alias_sym_language),
	},
	16: {
		2: uint16(alias_sym_language),
	},
	18: {
		0: uint16(anon_alias_sym_expression),
		2: uint16(anon_alias_sym_expression),
	},
	24: {
		3: uint16(alias_sym_language),
	},
}

var ts_non_terminal_alias_map = [5]uint16_t{
	0: uint16(sym__expression_inner),
	1: uint16(2),
	2: uint16(sym__expression_inner),
	3: uint16(anon_alias_sym_expression),
}

var ts_primary_state_ids = [405]TSStateId{
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
	21:  uint16(9),
	22:  uint16(9),
	23:  uint16(23),
	24:  uint16(24),
	25:  uint16(25),
	26:  uint16(26),
	27:  uint16(27),
	28:  uint16(28),
	29:  uint16(29),
	30:  uint16(27),
	31:  uint16(24),
	32:  uint16(27),
	33:  uint16(24),
	34:  uint16(34),
	35:  uint16(35),
	36:  uint16(36),
	37:  uint16(37),
	38:  uint16(38),
	39:  uint16(39),
	40:  uint16(35),
	41:  uint16(41),
	42:  uint16(42),
	43:  uint16(38),
	44:  uint16(36),
	45:  uint16(35),
	46:  uint16(38),
	47:  uint16(36),
	48:  uint16(48),
	49:  uint16(49),
	50:  uint16(50),
	51:  uint16(51),
	52:  uint16(52),
	53:  uint16(53),
	54:  uint16(54),
	55:  uint16(55),
	56:  uint16(54),
	57:  uint16(55),
	58:  uint16(55),
	59:  uint16(59),
	60:  uint16(59),
	61:  uint16(55),
	62:  uint16(59),
	63:  uint16(59),
	64:  uint16(54),
	65:  uint16(54),
	66:  uint16(3),
	67:  uint16(2),
	68:  uint16(4),
	69:  uint16(69),
	70:  uint16(70),
	71:  uint16(71),
	72:  uint16(70),
	73:  uint16(71),
	74:  uint16(74),
	75:  uint16(75),
	76:  uint16(76),
	77:  uint16(77),
	78:  uint16(78),
	79:  uint16(79),
	80:  uint16(80),
	81:  uint16(77),
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
	94:  uint16(83),
	95:  uint16(95),
	96:  uint16(78),
	97:  uint16(97),
	98:  uint16(98),
	99:  uint16(99),
	100: uint16(100),
	101: uint16(101),
	102: uint16(82),
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
	138: uint16(6),
	139: uint16(6),
	140: uint16(140),
	141: uint16(141),
	142: uint16(5),
	143: uint16(143),
	144: uint16(144),
	145: uint16(145),
	146: uint16(5),
	147: uint16(147),
	148: uint16(148),
	149: uint16(149),
	150: uint16(20),
	151: uint16(151),
	152: uint16(152),
	153: uint16(153),
	154: uint16(154),
	155: uint16(155),
	156: uint16(20),
	157: uint16(157),
	158: uint16(151),
	159: uint16(152),
	160: uint16(152),
	161: uint16(161),
	162: uint16(151),
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
	184: uint16(170),
	185: uint16(185),
	186: uint16(182),
	187: uint16(187),
	188: uint16(188),
	189: uint16(189),
	190: uint16(190),
	191: uint16(191),
	192: uint16(39),
	193: uint16(77),
	194: uint16(41),
	195: uint16(48),
	196: uint16(196),
	197: uint16(197),
	198: uint16(198),
	199: uint16(199),
	200: uint16(200),
	201: uint16(201),
	202: uint16(39),
	203: uint16(41),
	204: uint16(204),
	205: uint16(48),
	206: uint16(77),
	207: uint16(207),
	208: uint16(208),
	209: uint16(189),
	210: uint16(198),
	211: uint16(211),
	212: uint16(212),
	213: uint16(207),
	214: uint16(208),
	215: uint16(189),
	216: uint16(198),
	217: uint16(207),
	218: uint16(218),
	219: uint16(208),
	220: uint16(220),
	221: uint16(221),
	222: uint16(222),
	223: uint16(223),
	224: uint16(224),
	225: uint16(51),
	226: uint16(226),
	227: uint16(221),
	228: uint16(228),
	229: uint16(19),
	230: uint16(82),
	231: uint16(231),
	232: uint16(232),
	233: uint16(52),
	234: uint16(11),
	235: uint16(50),
	236: uint16(49),
	237: uint16(237),
	238: uint16(83),
	239: uint16(53),
	240: uint16(7),
	241: uint16(241),
	242: uint16(2),
	243: uint16(3),
	244: uint16(244),
	245: uint16(245),
	246: uint16(4),
	247: uint16(247),
	248: uint16(8),
	249: uint16(249),
	250: uint16(250),
	251: uint16(251),
	252: uint16(252),
	253: uint16(16),
	254: uint16(11),
	255: uint16(83),
	256: uint16(53),
	257: uint16(7),
	258: uint16(78),
	259: uint16(13),
	260: uint16(14),
	261: uint16(78),
	262: uint16(8),
	263: uint16(15),
	264: uint16(16),
	265: uint16(17),
	266: uint16(52),
	267: uint16(17),
	268: uint16(51),
	269: uint16(13),
	270: uint16(221),
	271: uint16(19),
	272: uint16(50),
	273: uint16(49),
	274: uint16(274),
	275: uint16(275),
	276: uint16(82),
	277: uint16(277),
	278: uint16(278),
	279: uint16(14),
	280: uint16(224),
	281: uint16(277),
	282: uint16(277),
	283: uint16(283),
	284: uint16(284),
	285: uint16(278),
	286: uint16(286),
	287: uint16(224),
	288: uint16(288),
	289: uint16(15),
	290: uint16(278),
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
	304: uint16(161),
	305: uint16(157),
	306: uint16(299),
	307: uint16(303),
	308: uint16(308),
	309: uint16(309),
	310: uint16(310),
	311: uint16(311),
	312: uint16(312),
	313: uint16(299),
	314: uint16(303),
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
	361: uint16(333),
	362: uint16(346),
	363: uint16(363),
	364: uint16(364),
	365: uint16(347),
	366: uint16(366),
	367: uint16(323),
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
	382: uint16(333),
	383: uint16(346),
	384: uint16(358),
	385: uint16(385),
	386: uint16(347),
	387: uint16(366),
	388: uint16(323),
	389: uint16(389),
	390: uint16(390),
	391: uint16(326),
	392: uint16(392),
	393: uint16(393),
	394: uint16(366),
	395: uint16(395),
	396: uint16(396),
	397: uint16(344),
	398: uint16(398),
	399: uint16(399),
	400: uint16(400),
	401: uint16(401),
	402: uint16(402),
	403: uint16(344),
	404: uint16(358),
}

func ts_lex(tls *libc.TLS, lexer uintptr, state TSStateId) (r uint8) {
	var eof, result, skip uint8
	var i, i1, i2, i3 uint32_t
	var lookahead int32_t
	_, _, _, _, _, _, _, _ = eof, i, i1, i2, i3, lookahead, result, skip
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
			state = uint16(38)
			goto next_state
		}
		i = uint32(0)
		for {
			if !(uint64(i) < libc.Uint64FromInt64(72)/libc.Uint64FromInt64(2)) {
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
		if lookahead == int32('\\') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(33)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(43)
			goto next_state
		}
		if lookahead == int32('`') {
			state = uint16(101)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(47)
			goto next_state
		}
		if lookahead == int32('}') {
			state = uint16(48)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(0)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(106)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('_') <= lookahead && lookahead <= int32('z') {
			state = uint16(105)
			goto next_state
		}
		return result
	case int32(1):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(15)
			goto next_state
		}
		return result
	case int32(2):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32('\r') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(1)
			goto next_state
		}
		return result
	case int32(3):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(20)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(20)
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
			state = uint16(5)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(81)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(81)
			goto next_state
		}
		if lookahead == int32('`') {
			state = uint16(100)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(87)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(83)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(81)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead == int32('\n') {
			state = uint16(98)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead == int32('\n') {
			state = uint16(98)
			goto next_state
		}
		if lookahead == int32('\r') {
			state = uint16(6)
			goto next_state
		}
		if lookahead == int32('"') || lookahead == int32('\\') || lookahead == int32('n') || lookahead == int32('r') || lookahead == int32('t') {
			state = uint16(97)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(8)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(76)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(67)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(73)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(74)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(75)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(9)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(81)
			goto next_state
		}
		if lookahead == int32('\'') {
			state = uint16(85)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(81)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(82)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(81)
			goto next_state
		}
		return result
	case int32(10):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(21)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(21)
			goto next_state
		}
		if lookahead == int32('\r') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(10)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(12)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(81)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(81)
			goto next_state
		}
		if lookahead == int32('`') {
			state = uint16(86)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(87)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(84)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(81)
			goto next_state
		}
		return result
	case int32(13):
		if lookahead == int32('\n') {
			state = uint16(99)
			goto next_state
		}
		return result
	case int32(14):
		if lookahead == int32('\n') {
			state = uint16(99)
			goto next_state
		}
		if lookahead == int32('\r') {
			state = uint16(13)
			goto next_state
		}
		if lookahead == int32('"') || lookahead == int32('\\') || lookahead == int32('n') || lookahead == int32('r') || lookahead == int32('t') {
			state = uint16(97)
			goto next_state
		}
		return result
	case int32(15):
		i1 = uint32(0)
		for {
			if !(uint64(i1) < libc.Uint64FromInt64(56)/libc.Uint64FromInt64(2)) {
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
		if lookahead == int32('\\') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(2)
			goto next_state
		}
		if lookahead == int32('`') {
			state = uint16(101)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(30)
			goto next_state
		}
		if lookahead == int32('}') {
			state = uint16(48)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(15)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(106)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('_') <= lookahead && lookahead <= int32('z') {
			state = uint16(105)
			goto next_state
		}
		return result
	case int32(16):
		if lookahead == int32('"') {
			state = uint16(17)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\\') {
			state = uint16(96)
			goto next_state
		}
		return result
	case int32(17):
		if lookahead == int32('"') {
			state = uint16(93)
			goto next_state
		}
		return result
	case int32(18):
		if lookahead == int32('"') {
			state = uint16(16)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(94)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(7)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(95)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(96)
			goto next_state
		}
		return result
	case int32(19):
		if lookahead == int32('"') {
			state = uint16(88)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(90)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(14)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(91)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(92)
			goto next_state
		}
		return result
	case int32(20):
		if lookahead == int32('#') {
			state = uint16(108)
			goto next_state
		}
		if lookahead == int32('-') {
			state = uint16(65)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(59)
			goto next_state
		}
		if lookahead == int32('\\') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(4)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(30)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(20)
			goto next_state
		}
		return result
	case int32(21):
		if lookahead == int32('#') {
			state = uint16(110)
			goto next_state
		}
		if lookahead == int32('\\') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(11)
			goto next_state
		}
		if lookahead == int32('`') {
			state = uint16(100)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(21)
			goto next_state
		}
		return result
	case int32(22):
		if lookahead == int32('&') {
			state = uint16(62)
			goto next_state
		}
		return result
	case int32(23):
		if lookahead == int32('\'') {
			state = uint16(79)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(25)
			goto next_state
		}
		return result
	case int32(24):
		if lookahead == int32('\'') {
			state = uint16(80)
			goto next_state
		}
		return result
	case int32(25):
		if lookahead == int32('\'') {
			state = uint16(78)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(25)
			goto next_state
		}
		return result
	case int32(26):
		if lookahead == int32('/') {
			state = uint16(67)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') {
			state = uint16(26)
			goto next_state
		}
		return result
	case int32(27):
		if lookahead == int32('=') {
			state = uint16(50)
			goto next_state
		}
		return result
	case int32(28):
		if lookahead == int32('=') {
			state = uint16(49)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(51)
			goto next_state
		}
		return result
	case int32(29):
		if lookahead == int32('`') {
			state = uint16(102)
			goto next_state
		}
		return result
	case int32(30):
		if lookahead == int32('{') {
			state = uint16(103)
			goto next_state
		}
		return result
	case int32(31):
		if lookahead == int32('}') {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(32):
		if eof != 0 {
			state = uint16(38)
			goto next_state
		}
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(36)
			goto next_state
		}
		return result
	case int32(33):
		if eof != 0 {
			state = uint16(38)
			goto next_state
		}
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(36)
			goto next_state
		}
		if lookahead == int32('\r') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(32)
			goto next_state
		}
		return result
	case int32(34):
		if eof != 0 {
			state = uint16(38)
			goto next_state
		}
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(37)
			goto next_state
		}
		return result
	case int32(35):
		if eof != 0 {
			state = uint16(38)
			goto next_state
		}
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(37)
			goto next_state
		}
		if lookahead == int32('\r') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(34)
			goto next_state
		}
		return result
	case int32(36):
		if eof != 0 {
			state = uint16(38)
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
		if lookahead == int32('\\') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(33)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(43)
			goto next_state
		}
		if lookahead == int32('`') {
			state = uint16(101)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(47)
			goto next_state
		}
		if lookahead == int32('}') {
			state = uint16(48)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(36)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(106)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('_') <= lookahead && lookahead <= int32('z') {
			state = uint16(105)
			goto next_state
		}
		return result
	case int32(37):
		if eof != 0 {
			state = uint16(38)
			goto next_state
		}
		i3 = uint32(0)
		for {
			if !(uint64(i3) < libc.Uint64FromInt64(72)/libc.Uint64FromInt64(2)) {
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
		if lookahead == int32('\\') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(35)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(43)
			goto next_state
		}
		if lookahead == int32('`') {
			state = uint16(101)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(46)
			goto next_state
		}
		if lookahead == int32('}') {
			state = uint16(31)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(37)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(106)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('_') <= lookahead && lookahead <= int32('z') {
			state = uint16(105)
			goto next_state
		}
		return result
	case int32(38):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(ts_builtin_sym_end)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(39):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COLON_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(40):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_QMARK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(41):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(42):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COMMA)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(43):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(44):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SLASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(45):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PLUS)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(46):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(47):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('{') {
			state = uint16(103)
			goto next_state
		}
		return result
	case int32(48):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(49):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EQ_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(50):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BANG_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(51):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EQ_TILDE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(52):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(53):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(54):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(55):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('=') {
			state = uint16(49)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(51)
			goto next_state
		}
		return result
	case int32(56):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COLON)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(57):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COLON)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('=') {
			state = uint16(39)
			goto next_state
		}
		return result
	case int32(58):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(59):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') {
			state = uint16(63)
			goto next_state
		}
		return result
	case int32(60):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOLLAR)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(61):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STAR)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(62):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AMP_AMP)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(63):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AT_DASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(64):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DASH_AT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(65):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('@') {
			state = uint16(64)
			goto next_state
		}
		return result
	case int32(66):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_shebang_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\t') || lookahead == int32(' ') {
			state = uint16(66)
			goto next_state
		}
		return result
	case int32(67):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__shebang_with_lang_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('/') {
			state = uint16(67)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') {
			state = uint16(26)
			goto next_state
		}
		return result
	case int32(68):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__shebang_with_lang_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('/') {
			state = uint16(68)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') {
			state = uint16(109)
			goto next_state
		}
		return result
	case int32(69):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__shebang_with_lang_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') {
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(70):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__shebang_with_lang_token3)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\r') {
			state = uint16(72)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(72)
			goto next_state
		}
		return result
	case int32(71):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__shebang_with_lang_token3)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(72)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(70)
			goto next_state
		}
		if lookahead == int32('\t') || int32(0x0b) <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(71)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) {
			state = uint16(72)
			goto next_state
		}
		return result
	case int32(72):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__shebang_with_lang_token3)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(72)
			goto next_state
		}
		return result
	case int32(73):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__opaque_shebang)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\r') {
			state = uint16(77)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(0x0b) || lookahead == int32('\f') || lookahead == int32(' ') {
			state = uint16(77)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32('/') {
			state = uint16(75)
			goto next_state
		}
		return result
	case int32(74):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__opaque_shebang)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(76)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(73)
			goto next_state
		}
		if lookahead == int32('\t') || int32(0x0b) <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(74)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32('/') {
			state = uint16(75)
			goto next_state
		}
		return result
	case int32(75):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__opaque_shebang)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('/') {
			state = uint16(67)
			goto next_state
		}
		if lookahead == int32('\t') || int32(0x0b) <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(77)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) {
			state = uint16(75)
			goto next_state
		}
		return result
	case int32(76):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__opaque_shebang)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('/') {
			state = uint16(68)
			goto next_state
		}
		if lookahead == int32('\t') || int32(0x0b) <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(77)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) {
			state = uint16(76)
			goto next_state
		}
		return result
	case int32(77):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__opaque_shebang)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('/') {
			state = uint16(77)
			goto next_state
		}
		return result
	case int32(78):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(79):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\'') {
			state = uint16(80)
			goto next_state
		}
		return result
	case int32(80):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(81):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__raw_string_indented_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(82):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__raw_string_indented_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(81)
			goto next_state
		}
		if lookahead == int32('\'') {
			state = uint16(85)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(81)
			goto next_state
		}
		if lookahead == int32('\t') || int32(0x0b) <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(82)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) {
			state = uint16(81)
			goto next_state
		}
		return result
	case int32(83):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__raw_string_indented_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(81)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(81)
			goto next_state
		}
		if lookahead == int32('`') {
			state = uint16(100)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(87)
			goto next_state
		}
		if lookahead == int32('\t') || int32(0x0b) <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(83)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) {
			state = uint16(81)
			goto next_state
		}
		return result
	case int32(84):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__raw_string_indented_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(81)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(81)
			goto next_state
		}
		if lookahead == int32('`') {
			state = uint16(86)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(87)
			goto next_state
		}
		if lookahead == int32('\t') || int32(0x0b) <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(84)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) {
			state = uint16(81)
			goto next_state
		}
		return result
	case int32(85):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__raw_string_indented_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\'') {
			state = uint16(24)
			goto next_state
		}
		return result
	case int32(86):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__raw_string_indented_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('`') {
			state = uint16(29)
			goto next_state
		}
		return result
	case int32(87):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__raw_string_indented_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('{') {
			state = uint16(103)
			goto next_state
		}
		return result
	case int32(88):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(89):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('"') {
			state = uint16(17)
			goto next_state
		}
		return result
	case int32(90):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__string_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(92)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('"') && lookahead != int32('\\') {
			state = uint16(90)
			goto next_state
		}
		return result
	case int32(91):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__string_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(90)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(14)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(91)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('"') && lookahead != int32('#') {
			state = uint16(92)
			goto next_state
		}
		return result
	case int32(92):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__string_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('"') && lookahead != int32('\\') {
			state = uint16(92)
			goto next_state
		}
		return result
	case int32(93):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(94):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__string_indented_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(96)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('"') && lookahead != int32('\\') {
			state = uint16(94)
			goto next_state
		}
		return result
	case int32(95):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__string_indented_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('"') {
			state = uint16(16)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(94)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(7)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(95)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(96)
			goto next_state
		}
		return result
	case int32(96):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__string_indented_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('"') && lookahead != int32('\\') {
			state = uint16(96)
			goto next_state
		}
		return result
	case int32(97):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_escape_sequence)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(98):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_escape_sequence)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\\') {
			state = uint16(7)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(95)
			goto next_state
		}
		return result
	case int32(99):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_escape_sequence)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\\') {
			state = uint16(14)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(91)
			goto next_state
		}
		return result
	case int32(100):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(101):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('`') {
			state = uint16(29)
			goto next_state
		}
		return result
	case int32(102):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(103):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACE_LBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(104):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACE_RBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(105):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(105)
			goto next_state
		}
		return result
	case int32(106):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_numeric_error)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(107)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(106)
			goto next_state
		}
		return result
	case int32(107):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_numeric_error)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(107)
			goto next_state
		}
		return result
	case int32(108):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('!') {
			state = uint16(66)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(110)
			goto next_state
		}
		return result
	case int32(109):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('/') {
			state = uint16(68)
			goto next_state
		}
		if lookahead == int32('\t') || int32(0x0b) <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(110)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) {
			state = uint16(109)
			goto next_state
		}
		return result
	case int32(110):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(110)
			goto next_state
		}
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var map_token = [36]uint16_t{
	0:  uint16('!'),
	1:  uint16(27),
	2:  uint16('"'),
	3:  uint16(89),
	4:  uint16('#'),
	5:  uint16(108),
	6:  uint16('$'),
	7:  uint16(60),
	8:  uint16('&'),
	9:  uint16(22),
	10: uint16('\''),
	11: uint16(23),
	12: uint16('('),
	13: uint16(52),
	14: uint16(')'),
	15: uint16(53),
	16: uint16('*'),
	17: uint16(61),
	18: uint16('+'),
	19: uint16(45),
	20: uint16(','),
	21: uint16(42),
	22: uint16('-'),
	23: uint16(65),
	24: uint16('/'),
	25: uint16(44),
	26: uint16(':'),
	27: uint16(57),
	28: uint16('='),
	29: uint16(55),
	30: uint16('?'),
	31: uint16(40),
	32: uint16('@'),
	33: uint16(59),
	34: uint16('['),
	35: uint16(41),
}

var map_token1 = [28]uint16_t{
	0:  uint16('"'),
	1:  uint16(89),
	2:  uint16('#'),
	3:  uint16(110),
	4:  uint16('$'),
	5:  uint16(60),
	6:  uint16('\''),
	7:  uint16(23),
	8:  uint16('('),
	9:  uint16(52),
	10: uint16(')'),
	11: uint16(53),
	12: uint16('*'),
	13: uint16(61),
	14: uint16('+'),
	15: uint16(45),
	16: uint16(','),
	17: uint16(42),
	18: uint16('-'),
	19: uint16(65),
	20: uint16('/'),
	21: uint16(44),
	22: uint16(':'),
	23: uint16(56),
	24: uint16('='),
	25: uint16(54),
	26: uint16('@'),
	27: uint16(59),
}

var map_token2 = [36]uint16_t{
	0:  uint16('!'),
	1:  uint16(27),
	2:  uint16('"'),
	3:  uint16(89),
	4:  uint16('#'),
	5:  uint16(108),
	6:  uint16('$'),
	7:  uint16(60),
	8:  uint16('&'),
	9:  uint16(22),
	10: uint16('\''),
	11: uint16(23),
	12: uint16('('),
	13: uint16(52),
	14: uint16(')'),
	15: uint16(53),
	16: uint16('*'),
	17: uint16(61),
	18: uint16('+'),
	19: uint16(45),
	20: uint16(','),
	21: uint16(42),
	22: uint16('-'),
	23: uint16(65),
	24: uint16('/'),
	25: uint16(44),
	26: uint16(':'),
	27: uint16(57),
	28: uint16('='),
	29: uint16(55),
	30: uint16('?'),
	31: uint16(40),
	32: uint16('@'),
	33: uint16(59),
	34: uint16('['),
	35: uint16(41),
}

var map_token3 = [36]uint16_t{
	0:  uint16('!'),
	1:  uint16(27),
	2:  uint16('"'),
	3:  uint16(89),
	4:  uint16('#'),
	5:  uint16(110),
	6:  uint16('$'),
	7:  uint16(60),
	8:  uint16('&'),
	9:  uint16(22),
	10: uint16('\''),
	11: uint16(23),
	12: uint16('('),
	13: uint16(52),
	14: uint16(')'),
	15: uint16(53),
	16: uint16('*'),
	17: uint16(61),
	18: uint16('+'),
	19: uint16(45),
	20: uint16(','),
	21: uint16(42),
	22: uint16('-'),
	23: uint16(69),
	24: uint16('/'),
	25: uint16(44),
	26: uint16(':'),
	27: uint16(57),
	28: uint16('='),
	29: uint16(28),
	30: uint16('?'),
	31: uint16(40),
	32: uint16('@'),
	33: uint16(58),
	34: uint16('['),
	35: uint16(41),
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
	switch int32(state) {
	case 0:
		if lookahead == int32('\\') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(1)
			goto next_state
		}
		if lookahead == int32('a') {
			state = uint16(2)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(3)
			goto next_state
		}
		if lookahead == int32('f') {
			state = uint16(4)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(5)
			goto next_state
		}
		if lookahead == int32('m') {
			state = uint16(6)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(7)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(8)
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
			state = uint16(9)
			goto next_state
		}
		if lookahead == int32('\r') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(10)
			goto next_state
		}
		return result
	case int32(2):
		if lookahead == int32('l') {
			state = uint16(11)
			goto next_state
		}
		return result
	case int32(3):
		if lookahead == int32('l') {
			state = uint16(12)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(13)
			goto next_state
		}
		if lookahead == int32('x') {
			state = uint16(14)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead == int32('a') {
			state = uint16(15)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead == int32('f') {
			state = uint16(16)
			goto next_state
		}
		if lookahead == int32('m') {
			state = uint16(17)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead == int32('o') {
			state = uint16(18)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead == int32('e') {
			state = uint16(19)
			goto next_state
		}
		if lookahead == int32('h') {
			state = uint16(20)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead == int32('r') {
			state = uint16(21)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead == int32('\\') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(1)
			goto next_state
		}
		if lookahead == int32('a') {
			state = uint16(2)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(3)
			goto next_state
		}
		if lookahead == int32('f') {
			state = uint16(4)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(5)
			goto next_state
		}
		if lookahead == int32('m') {
			state = uint16(6)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(7)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(8)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(9)
			goto next_state
		}
		return result
	case int32(10):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(9)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead == int32('i') {
			state = uint16(22)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead == int32('s') {
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(13):
		if lookahead == int32('v') {
			state = uint16(24)
			goto next_state
		}
		return result
	case int32(14):
		if lookahead == int32('p') {
			state = uint16(25)
			goto next_state
		}
		return result
	case int32(15):
		if lookahead == int32('l') {
			state = uint16(26)
			goto next_state
		}
		return result
	case int32(16):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_if)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(17):
		if lookahead == int32('p') {
			state = uint16(27)
			goto next_state
		}
		return result
	case int32(18):
		if lookahead == int32('d') {
			state = uint16(28)
			goto next_state
		}
		return result
	case int32(19):
		if lookahead == int32('t') {
			state = uint16(29)
			goto next_state
		}
		return result
	case int32(20):
		if lookahead == int32('e') {
			state = uint16(30)
			goto next_state
		}
		return result
	case int32(21):
		if lookahead == int32('u') {
			state = uint16(31)
			goto next_state
		}
		return result
	case int32(22):
		if lookahead == int32('a') {
			state = uint16(32)
			goto next_state
		}
		return result
	case int32(23):
		if lookahead == int32('e') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(24):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_env)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(25):
		if lookahead == int32('o') {
			state = uint16(34)
			goto next_state
		}
		return result
	case int32(26):
		if lookahead == int32('s') {
			state = uint16(35)
			goto next_state
		}
		return result
	case int32(27):
		if lookahead == int32('o') {
			state = uint16(36)
			goto next_state
		}
		return result
	case int32(28):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_mod)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(29):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_set)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(30):
		if lookahead == int32('l') {
			state = uint16(37)
			goto next_state
		}
		return result
	case int32(31):
		if lookahead == int32('e') {
			state = uint16(38)
			goto next_state
		}
		return result
	case int32(32):
		if lookahead == int32('s') {
			state = uint16(39)
			goto next_state
		}
		return result
	case int32(33):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_else)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(34):
		if lookahead == int32('r') {
			state = uint16(40)
			goto next_state
		}
		return result
	case int32(35):
		if lookahead == int32('e') {
			state = uint16(41)
			goto next_state
		}
		return result
	case int32(36):
		if lookahead == int32('r') {
			state = uint16(42)
			goto next_state
		}
		return result
	case int32(37):
		if lookahead == int32('l') {
			state = uint16(43)
			goto next_state
		}
		return result
	case int32(38):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_true)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(39):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_alias)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(40):
		if lookahead == int32('t') {
			state = uint16(44)
			goto next_state
		}
		return result
	case int32(41):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_false)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(42):
		if lookahead == int32('t') {
			state = uint16(45)
			goto next_state
		}
		return result
	case int32(43):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_shell)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(44):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_export)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(45):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_import)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var ts_lex_modes = [405]TSLexerMode{
	0: {
		Fexternal_lex_state: uint16(1),
	},
	1: {},
	2: {
		Flex_state: uint16(37),
	},
	3: {
		Flex_state: uint16(37),
	},
	4: {
		Flex_state: uint16(37),
	},
	5: {
		Flex_state: uint16(37),
	},
	6: {
		Flex_state: uint16(37),
	},
	7: {
		Flex_state: uint16(37),
	},
	8: {
		Flex_state: uint16(37),
	},
	9: {
		Flex_state: uint16(37),
	},
	10: {
		Flex_state: uint16(37),
	},
	11: {
		Flex_state: uint16(37),
	},
	12: {
		Flex_state: uint16(37),
	},
	13: {
		Flex_state: uint16(37),
	},
	14: {
		Flex_state: uint16(37),
	},
	15: {
		Flex_state: uint16(37),
	},
	16: {
		Flex_state: uint16(37),
	},
	17: {
		Flex_state: uint16(37),
	},
	18: {
		Flex_state: uint16(37),
	},
	19: {
		Flex_state: uint16(37),
	},
	20: {
		Flex_state: uint16(37),
	},
	21: {
		Flex_state: uint16(37),
	},
	22: {
		Flex_state: uint16(37),
	},
	23: {
		Flex_state: uint16(37),
	},
	24: {
		Flex_state: uint16(37),
	},
	25: {
		Flex_state: uint16(37),
	},
	26: {
		Flex_state: uint16(37),
	},
	27: {
		Flex_state: uint16(37),
	},
	28: {
		Flex_state: uint16(37),
	},
	29: {
		Flex_state: uint16(37),
	},
	30: {
		Flex_state: uint16(37),
	},
	31: {
		Flex_state: uint16(37),
	},
	32: {
		Flex_state: uint16(37),
	},
	33: {
		Flex_state: uint16(37),
	},
	34: {
		Flex_state: uint16(37),
	},
	35: {
		Flex_state: uint16(37),
	},
	36: {
		Flex_state: uint16(37),
	},
	37: {
		Flex_state: uint16(37),
	},
	38: {
		Flex_state: uint16(37),
	},
	39: {
		Flex_state: uint16(37),
	},
	40: {
		Flex_state: uint16(37),
	},
	41: {
		Flex_state: uint16(37),
	},
	42: {
		Flex_state: uint16(37),
	},
	43: {
		Flex_state: uint16(37),
	},
	44: {
		Flex_state: uint16(37),
	},
	45: {
		Flex_state: uint16(37),
	},
	46: {
		Flex_state: uint16(37),
	},
	47: {
		Flex_state: uint16(37),
	},
	48: {
		Flex_state: uint16(37),
	},
	49: {
		Flex_state: uint16(37),
	},
	50: {
		Flex_state: uint16(37),
	},
	51: {
		Flex_state: uint16(37),
	},
	52: {
		Flex_state: uint16(37),
	},
	53: {
		Flex_state: uint16(37),
	},
	54: {
		Flex_state: uint16(37),
	},
	55: {
		Flex_state: uint16(37),
	},
	56: {
		Flex_state: uint16(37),
	},
	57: {
		Flex_state: uint16(37),
	},
	58: {
		Flex_state: uint16(37),
	},
	59: {
		Flex_state: uint16(37),
	},
	60: {
		Flex_state: uint16(37),
	},
	61: {
		Flex_state: uint16(37),
	},
	62: {
		Flex_state: uint16(37),
	},
	63: {
		Flex_state: uint16(37),
	},
	64: {
		Flex_state: uint16(37),
	},
	65: {
		Flex_state: uint16(37),
	},
	66: {
		Flex_state: uint16(15),
	},
	67: {
		Flex_state: uint16(15),
	},
	68: {
		Flex_state: uint16(15),
	},
	69: {
		Flex_state:          uint16(20),
		Fexternal_lex_state: uint16(2),
	},
	70: {
		Flex_state: uint16(37),
	},
	71: {
		Flex_state: uint16(37),
	},
	72: {
		Flex_state: uint16(37),
	},
	73: {
		Flex_state: uint16(37),
	},
	74: {
		Flex_state: uint16(37),
	},
	75: {
		Flex_state: uint16(37),
	},
	76: {
		Flex_state: uint16(37),
	},
	77: {
		Flex_state: uint16(37),
	},
	78: {
		Flex_state: uint16(37),
	},
	79: {
		Flex_state:          uint16(15),
		Fexternal_lex_state: uint16(2),
	},
	80: {
		Flex_state:          uint16(15),
		Fexternal_lex_state: uint16(2),
	},
	81: {
		Flex_state: uint16(37),
	},
	82: {
		Flex_state: uint16(37),
	},
	83: {
		Flex_state: uint16(37),
	},
	84: {
		Flex_state:          uint16(15),
		Fexternal_lex_state: uint16(2),
	},
	85: {
		Flex_state:          uint16(15),
		Fexternal_lex_state: uint16(2),
	},
	86: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(3),
	},
	87: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(3),
	},
	88: {
		Flex_state: uint16(37),
	},
	89: {
		Flex_state: uint16(37),
	},
	90: {
		Flex_state: uint16(37),
	},
	91: {
		Flex_state: uint16(37),
	},
	92: {
		Flex_state: uint16(37),
	},
	93: {
		Flex_state: uint16(37),
	},
	94: {
		Flex_state: uint16(37),
	},
	95: {
		Flex_state: uint16(37),
	},
	96: {
		Flex_state: uint16(37),
	},
	97: {
		Flex_state: uint16(37),
	},
	98: {
		Flex_state: uint16(37),
	},
	99: {
		Flex_state: uint16(37),
	},
	100: {
		Flex_state: uint16(37),
	},
	101: {
		Flex_state: uint16(37),
	},
	102: {
		Flex_state: uint16(37),
	},
	103: {
		Flex_state: uint16(37),
	},
	104: {
		Flex_state: uint16(37),
	},
	105: {
		Flex_state: uint16(37),
	},
	106: {
		Flex_state: uint16(37),
	},
	107: {
		Flex_state: uint16(37),
	},
	108: {
		Flex_state: uint16(37),
	},
	109: {
		Flex_state: uint16(37),
	},
	110: {
		Flex_state: uint16(37),
	},
	111: {
		Flex_state: uint16(37),
	},
	112: {
		Flex_state: uint16(37),
	},
	113: {
		Flex_state: uint16(37),
	},
	114: {
		Flex_state: uint16(37),
	},
	115: {
		Flex_state: uint16(37),
	},
	116: {
		Flex_state: uint16(37),
	},
	117: {
		Flex_state: uint16(37),
	},
	118: {
		Flex_state: uint16(37),
	},
	119: {
		Flex_state: uint16(37),
	},
	120: {
		Flex_state: uint16(37),
	},
	121: {
		Flex_state: uint16(37),
	},
	122: {
		Flex_state: uint16(37),
	},
	123: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	124: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	125: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	126: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	127: {
		Flex_state: uint16(37),
	},
	128: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	129: {
		Flex_state: uint16(37),
	},
	130: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	131: {
		Flex_state: uint16(37),
	},
	132: {
		Flex_state: uint16(37),
	},
	133: {
		Flex_state: uint16(37),
	},
	134: {
		Flex_state:          uint16(15),
		Fexternal_lex_state: uint16(2),
	},
	135: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	136: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	137: {
		Flex_state: uint16(37),
	},
	138: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	139: {
		Flex_state: uint16(15),
	},
	140: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	141: {
		Flex_state: uint16(37),
	},
	142: {
		Flex_state: uint16(15),
	},
	143: {
		Flex_state: uint16(37),
	},
	144: {
		Flex_state: uint16(37),
	},
	145: {
		Flex_state: uint16(37),
	},
	146: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	147: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	148: {
		Flex_state: uint16(37),
	},
	149: {
		Flex_state: uint16(37),
	},
	150: {
		Flex_state: uint16(15),
	},
	151: {
		Flex_state: uint16(5),
	},
	152: {
		Flex_state: uint16(12),
	},
	153: {
		Flex_state: uint16(37),
	},
	154: {
		Flex_state: uint16(37),
	},
	155: {
		Flex_state: uint16(37),
	},
	156: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	157: {
		Flex_state: uint16(15),
	},
	158: {
		Flex_state: uint16(5),
	},
	159: {
		Flex_state: uint16(12),
	},
	160: {
		Flex_state: uint16(12),
	},
	161: {
		Flex_state: uint16(15),
	},
	162: {
		Flex_state: uint16(5),
	},
	163: {
		Flex_state: uint16(37),
	},
	164: {
		Flex_state: uint16(37),
	},
	165: {
		Flex_state: uint16(37),
	},
	166: {
		Flex_state: uint16(37),
	},
	167: {
		Flex_state: uint16(37),
	},
	168: {
		Flex_state: uint16(37),
	},
	169: {
		Flex_state: uint16(37),
	},
	170: {
		Flex_state: uint16(5),
	},
	171: {
		Flex_state:          uint16(15),
		Fexternal_lex_state: uint16(5),
	},
	172: {
		Flex_state:          uint16(15),
		Fexternal_lex_state: uint16(5),
	},
	173: {
		Flex_state: uint16(37),
	},
	174: {
		Flex_state: uint16(37),
	},
	175: {
		Flex_state: uint16(37),
	},
	176: {
		Flex_state: uint16(37),
	},
	177: {
		Flex_state: uint16(37),
	},
	178: {
		Flex_state: uint16(37),
	},
	179: {
		Flex_state: uint16(37),
	},
	180: {
		Flex_state: uint16(37),
	},
	181: {
		Flex_state: uint16(37),
	},
	182: {
		Flex_state: uint16(12),
	},
	183: {
		Flex_state: uint16(37),
	},
	184: {
		Flex_state: uint16(12),
	},
	185: {
		Flex_state: uint16(37),
	},
	186: {
		Flex_state: uint16(5),
	},
	187: {
		Flex_state:          uint16(15),
		Fexternal_lex_state: uint16(5),
	},
	188: {
		Flex_state: uint16(18),
	},
	189: {
		Flex_state: uint16(19),
	},
	190: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	191: {
		Flex_state: uint16(15),
	},
	192: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	193: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	194: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	195: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	196: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	197: {
		Flex_state: uint16(19),
	},
	198: {
		Flex_state: uint16(18),
	},
	199: {
		Flex_state: uint16(37),
	},
	200: {
		Flex_state: uint16(15),
	},
	201: {
		Flex_state: uint16(37),
	},
	202: {
		Flex_state: uint16(15),
	},
	203: {
		Flex_state: uint16(15),
	},
	204: {
		Flex_state: uint16(37),
	},
	205: {
		Flex_state: uint16(15),
	},
	206: {
		Flex_state: uint16(15),
	},
	207: {
		Flex_state: uint16(19),
	},
	208: {
		Flex_state: uint16(18),
	},
	209: {
		Flex_state: uint16(19),
	},
	210: {
		Flex_state: uint16(18),
	},
	211: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	212: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	213: {
		Flex_state: uint16(19),
	},
	214: {
		Flex_state: uint16(18),
	},
	215: {
		Flex_state: uint16(19),
	},
	216: {
		Flex_state: uint16(18),
	},
	217: {
		Flex_state: uint16(19),
	},
	218: {
		Flex_state:          uint16(15),
		Fexternal_lex_state: uint16(6),
	},
	219: {
		Flex_state: uint16(18),
	},
	220: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	221: {
		Flex_state: uint16(12),
	},
	222: {
		Flex_state: uint16(37),
	},
	223: {
		Flex_state: uint16(37),
	},
	224: {
		Flex_state: uint16(37),
	},
	225: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	226: {
		Flex_state: uint16(37),
	},
	227: {
		Flex_state: uint16(5),
	},
	228: {
		Flex_state: uint16(37),
	},
	229: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	230: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	231: {
		Flex_state: uint16(37),
	},
	232: {
		Flex_state: uint16(37),
	},
	233: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	234: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	235: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	236: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	237: {
		Flex_state: uint16(37),
	},
	238: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	239: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	240: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	241: {
		Flex_state: uint16(37),
	},
	242: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	243: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	244: {
		Flex_state: uint16(8),
	},
	245: {
		Flex_state: uint16(37),
	},
	246: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	247: {
		Flex_state: uint16(9),
	},
	248: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	249: {
		Flex_state: uint16(37),
	},
	250: {
		Flex_state: uint16(37),
	},
	251: {
		Flex_state: uint16(37),
	},
	252: {
		Flex_state: uint16(37),
	},
	253: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	254: {
		Flex_state: uint16(15),
	},
	255: {
		Flex_state: uint16(15),
	},
	256: {
		Flex_state: uint16(15),
	},
	257: {
		Flex_state: uint16(15),
	},
	258: {
		Flex_state: uint16(15),
	},
	259: {
		Flex_state: uint16(15),
	},
	260: {
		Flex_state: uint16(15),
	},
	261: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	262: {
		Flex_state: uint16(15),
	},
	263: {
		Flex_state: uint16(15),
	},
	264: {
		Flex_state: uint16(15),
	},
	265: {
		Flex_state: uint16(15),
	},
	266: {
		Flex_state: uint16(15),
	},
	267: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	268: {
		Flex_state: uint16(15),
	},
	269: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	270: {
		Flex_state:          uint16(15),
		Fexternal_lex_state: uint16(5),
	},
	271: {
		Flex_state: uint16(15),
	},
	272: {
		Flex_state: uint16(15),
	},
	273: {
		Flex_state: uint16(15),
	},
	274: {
		Flex_state: uint16(37),
	},
	275: {
		Flex_state: uint16(37),
	},
	276: {
		Flex_state: uint16(15),
	},
	277: {
		Flex_state: uint16(9),
	},
	278: {
		Flex_state: uint16(9),
	},
	279: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	280: {
		Flex_state: uint16(37),
	},
	281: {
		Flex_state: uint16(9),
	},
	282: {
		Flex_state: uint16(9),
	},
	283: {
		Flex_state: uint16(37),
	},
	284: {
		Flex_state: uint16(37),
	},
	285: {
		Flex_state: uint16(9),
	},
	286: {
		Flex_state: uint16(15),
	},
	287: {
		Flex_state: uint16(37),
	},
	288: {
		Flex_state: uint16(37),
	},
	289: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	290: {
		Flex_state: uint16(9),
	},
	291: {
		Flex_state: uint16(37),
	},
	292: {
		Flex_state: uint16(37),
	},
	293: {
		Flex_state: uint16(37),
	},
	294: {
		Flex_state: uint16(37),
	},
	295: {
		Flex_state: uint16(37),
	},
	296: {
		Flex_state: uint16(37),
	},
	297: {
		Flex_state: uint16(37),
	},
	298: {
		Flex_state: uint16(37),
	},
	299: {
		Flex_state: uint16(37),
	},
	300: {
		Flex_state: uint16(37),
	},
	301: {
		Flex_state: uint16(37),
	},
	302: {
		Flex_state: uint16(37),
	},
	303: {
		Flex_state: uint16(37),
	},
	304: {
		Flex_state: uint16(15),
	},
	305: {
		Flex_state: uint16(15),
	},
	306: {
		Flex_state: uint16(37),
	},
	307: {
		Flex_state: uint16(37),
	},
	308: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	309: {
		Flex_state: uint16(37),
	},
	310: {
		Flex_state: uint16(37),
	},
	311: {
		Flex_state: uint16(37),
	},
	312: {
		Flex_state:          uint16(15),
		Fexternal_lex_state: uint16(6),
	},
	313: {
		Flex_state: uint16(37),
	},
	314: {
		Flex_state: uint16(37),
	},
	315: {
		Flex_state: uint16(37),
	},
	316: {
		Flex_state: uint16(71),
	},
	317: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	318: {
		Flex_state: uint16(37),
	},
	319: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	320: {
		Flex_state: uint16(37),
	},
	321: {
		Flex_state: uint16(71),
	},
	322: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	323: {
		Flex_state: uint16(15),
	},
	324: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	325: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	326: {
		Flex_state: uint16(37),
	},
	327: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	328: {
		Flex_state: uint16(37),
	},
	329: {
		Flex_state: uint16(37),
	},
	330: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	331: {
		Flex_state: uint16(37),
	},
	332: {
		Flex_state: uint16(37),
	},
	333: {
		Flex_state: uint16(37),
	},
	334: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	335: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	336: {
		Flex_state: uint16(37),
	},
	337: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	338: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	339: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	340: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	341: {
		Flex_state: uint16(71),
	},
	342: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	343: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	344: {
		Flex_state: uint16(37),
	},
	345: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	346: {
		Flex_state: uint16(21),
	},
	347: {
		Flex_state: uint16(37),
	},
	348: {
		Flex_state: uint16(37),
	},
	349: {
		Flex_state: uint16(37),
	},
	350: {
		Flex_state: uint16(37),
	},
	351: {
		Flex_state: uint16(37),
	},
	352: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	353: {
		Flex_state: uint16(37),
	},
	354: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	355: {
		Flex_state: uint16(37),
	},
	356: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	357: {
		Flex_state: uint16(37),
	},
	358: {
		Flex_state: uint16(37),
	},
	359: {
		Flex_state: uint16(37),
	},
	360: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	361: {
		Flex_state: uint16(37),
	},
	362: {
		Flex_state: uint16(21),
	},
	363: {
		Flex_state: uint16(37),
	},
	364: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	365: {
		Flex_state: uint16(37),
	},
	366: {
		Flex_state: uint16(37),
	},
	367: {
		Flex_state: uint16(15),
	},
	368: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	369: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	370: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	371: {
		Flex_state: uint16(37),
	},
	372: {
		Flex_state: uint16(37),
	},
	373: {
		Flex_state: uint16(37),
	},
	374: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	375: {
		Flex_state: uint16(37),
	},
	376: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	377: {
		Flex_state: uint16(37),
	},
	378: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	379: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	380: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	381: {
		Flex_state: uint16(37),
	},
	382: {
		Flex_state: uint16(37),
	},
	383: {
		Flex_state: uint16(21),
	},
	384: {
		Flex_state: uint16(37),
	},
	385: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	386: {
		Flex_state: uint16(37),
	},
	387: {
		Flex_state: uint16(37),
	},
	388: {
		Flex_state: uint16(15),
	},
	389: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	390: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	391: {
		Flex_state: uint16(37),
	},
	392: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	393: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	394: {
		Flex_state: uint16(37),
	},
	395: {
		Flex_state: uint16(37),
	},
	396: {
		Flex_state: uint16(37),
	},
	397: {
		Flex_state: uint16(37),
	},
	398: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	399: {
		Flex_state: uint16(37),
	},
	400: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	401: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	402: {
		Flex_state:          uint16(37),
		Fexternal_lex_state: uint16(4),
	},
	403: {
		Flex_state: uint16(37),
	},
	404: {
		Flex_state: uint16(37),
	},
}

var ts_parse_table = [2][118]uint16_t{
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
		37: uint16(1),
		41: uint16(1),
		42: uint16(1),
		44: uint16(1),
		46: uint16(1),
		49: uint16(1),
		50: uint16(1),
		51: uint16(1),
		53: uint16(1),
		54: uint16(3),
		55: uint16(1),
		56: uint16(1),
		57: uint16(1),
		58: uint16(1),
		59: uint16(1),
	},
	1: {
		0:   uint16(5),
		1:   uint16(7),
		2:   uint16(9),
		4:   uint16(11),
		5:   uint16(13),
		7:   uint16(15),
		8:   uint16(17),
		9:   uint16(19),
		28:  uint16(21),
		35:  uint16(23),
		54:  uint16(3),
		60:  uint16(357),
		61:  uint16(26),
		62:  uint16(26),
		63:  uint16(26),
		64:  uint16(26),
		65:  uint16(26),
		66:  uint16(26),
		67:  uint16(26),
		82:  uint16(127),
		83:  uint16(26),
		84:  uint16(327),
		94:  uint16(330),
		101: uint16(26),
		102: uint16(127),
	},
}

var ts_small_parse_table = [7647]uint16_t{
	0:    uint16(3),
	1:    uint16(29),
	2:    uint16(1),
	3:    uint16(sym_comment),
	4:    uint16(27),
	5:    uint16(8),
	6:    uint16(anon_sym_alias),
	7:    uint16(anon_sym_export),
	8:    uint16(anon_sym_import),
	9:    uint16(anon_sym_mod),
	10:   uint16(anon_sym_set),
	11:   uint16(aux_sym_string_token1),
	12:   uint16(anon_sym_DQUOTE),
	13:   uint16(sym_identifier),
	14:   uint16(25),
	15:   uint16(16),
	17:   uint16(anon_sym_LBRACK),
	18:   uint16(anon_sym_COMMA),
	19:   uint16(anon_sym_RBRACK),
	20:   uint16(anon_sym_SLASH),
	21:   uint16(anon_sym_PLUS),
	22:   uint16(anon_sym_LBRACE),
	23:   uint16(anon_sym_EQ_EQ),
	24:   uint16(anon_sym_BANG_EQ),
	25:   uint16(anon_sym_EQ_TILDE),
	26:   uint16(anon_sym_RPAREN),
	27:   uint16(anon_sym_COLON),
	28:   uint16(anon_sym_AT),
	29:   uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	30:   uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	31:   uint16(anon_sym_RBRACE_RBRACE),
	32:   uint16(3),
	33:   uint16(29),
	34:   uint16(1),
	35:   uint16(sym_comment),
	36:   uint16(33),
	37:   uint16(8),
	38:   uint16(anon_sym_alias),
	39:   uint16(anon_sym_export),
	40:   uint16(anon_sym_import),
	41:   uint16(anon_sym_mod),
	42:   uint16(anon_sym_set),
	43:   uint16(aux_sym_string_token1),
	44:   uint16(anon_sym_DQUOTE),
	45:   uint16(sym_identifier),
	46:   uint16(31),
	47:   uint16(16),
	49:   uint16(anon_sym_LBRACK),
	50:   uint16(anon_sym_COMMA),
	51:   uint16(anon_sym_RBRACK),
	52:   uint16(anon_sym_SLASH),
	53:   uint16(anon_sym_PLUS),
	54:   uint16(anon_sym_LBRACE),
	55:   uint16(anon_sym_EQ_EQ),
	56:   uint16(anon_sym_BANG_EQ),
	57:   uint16(anon_sym_EQ_TILDE),
	58:   uint16(anon_sym_RPAREN),
	59:   uint16(anon_sym_COLON),
	60:   uint16(anon_sym_AT),
	61:   uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	62:   uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	63:   uint16(anon_sym_RBRACE_RBRACE),
	64:   uint16(3),
	65:   uint16(29),
	66:   uint16(1),
	67:   uint16(sym_comment),
	68:   uint16(37),
	69:   uint16(8),
	70:   uint16(anon_sym_alias),
	71:   uint16(anon_sym_export),
	72:   uint16(anon_sym_import),
	73:   uint16(anon_sym_mod),
	74:   uint16(anon_sym_set),
	75:   uint16(aux_sym_string_token1),
	76:   uint16(anon_sym_DQUOTE),
	77:   uint16(sym_identifier),
	78:   uint16(35),
	79:   uint16(16),
	81:   uint16(anon_sym_LBRACK),
	82:   uint16(anon_sym_COMMA),
	83:   uint16(anon_sym_RBRACK),
	84:   uint16(anon_sym_SLASH),
	85:   uint16(anon_sym_PLUS),
	86:   uint16(anon_sym_LBRACE),
	87:   uint16(anon_sym_EQ_EQ),
	88:   uint16(anon_sym_BANG_EQ),
	89:   uint16(anon_sym_EQ_TILDE),
	90:   uint16(anon_sym_RPAREN),
	91:   uint16(anon_sym_COLON),
	92:   uint16(anon_sym_AT),
	93:   uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	94:   uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	95:   uint16(anon_sym_RBRACE_RBRACE),
	96:   uint16(7),
	97:   uint16(29),
	98:   uint16(1),
	99:   uint16(sym_comment),
	100:  uint16(43),
	101:  uint16(1),
	102:  uint16(anon_sym_else),
	103:  uint16(6),
	104:  uint16(1),
	105:  uint16(aux_sym_if_expression_repeat1),
	106:  uint16(39),
	107:  uint16(1),
	108:  uint16(sym_else_if_clause),
	109:  uint16(51),
	110:  uint16(1),
	111:  uint16(sym_else_clause),
	112:  uint16(39),
	113:  uint16(5),
	114:  uint16(anon_sym_if),
	115:  uint16(aux_sym_string_token1),
	116:  uint16(anon_sym_DQUOTE),
	117:  uint16(anon_sym_BQUOTE),
	118:  uint16(sym_identifier),
	119:  uint16(41),
	120:  uint16(14),
	121:  uint16(anon_sym_COMMA),
	122:  uint16(anon_sym_SLASH),
	123:  uint16(anon_sym_PLUS),
	124:  uint16(anon_sym_LBRACE),
	125:  uint16(anon_sym_EQ_EQ),
	126:  uint16(anon_sym_BANG_EQ),
	127:  uint16(anon_sym_EQ_TILDE),
	128:  uint16(anon_sym_LPAREN),
	129:  uint16(anon_sym_RPAREN),
	130:  uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	131:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	132:  uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	133:  uint16(anon_sym_RBRACE_RBRACE),
	134:  uint16(sym_numeric_error),
	135:  uint16(7),
	136:  uint16(29),
	137:  uint16(1),
	138:  uint16(sym_comment),
	139:  uint16(43),
	140:  uint16(1),
	141:  uint16(anon_sym_else),
	142:  uint16(20),
	143:  uint16(1),
	144:  uint16(aux_sym_if_expression_repeat1),
	145:  uint16(39),
	146:  uint16(1),
	147:  uint16(sym_else_if_clause),
	148:  uint16(49),
	149:  uint16(1),
	150:  uint16(sym_else_clause),
	151:  uint16(45),
	152:  uint16(5),
	153:  uint16(anon_sym_if),
	154:  uint16(aux_sym_string_token1),
	155:  uint16(anon_sym_DQUOTE),
	156:  uint16(anon_sym_BQUOTE),
	157:  uint16(sym_identifier),
	158:  uint16(47),
	159:  uint16(14),
	160:  uint16(anon_sym_COMMA),
	161:  uint16(anon_sym_SLASH),
	162:  uint16(anon_sym_PLUS),
	163:  uint16(anon_sym_LBRACE),
	164:  uint16(anon_sym_EQ_EQ),
	165:  uint16(anon_sym_BANG_EQ),
	166:  uint16(anon_sym_EQ_TILDE),
	167:  uint16(anon_sym_LPAREN),
	168:  uint16(anon_sym_RPAREN),
	169:  uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	170:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	171:  uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	172:  uint16(anon_sym_RBRACE_RBRACE),
	173:  uint16(sym_numeric_error),
	174:  uint16(3),
	175:  uint16(29),
	176:  uint16(1),
	177:  uint16(sym_comment),
	178:  uint16(49),
	179:  uint16(5),
	180:  uint16(anon_sym_if),
	181:  uint16(aux_sym_string_token1),
	182:  uint16(anon_sym_DQUOTE),
	183:  uint16(anon_sym_BQUOTE),
	184:  uint16(sym_identifier),
	185:  uint16(51),
	186:  uint16(17),
	187:  uint16(anon_sym_COMMA),
	188:  uint16(anon_sym_SLASH),
	189:  uint16(anon_sym_PLUS),
	190:  uint16(anon_sym_LBRACE),
	191:  uint16(anon_sym_EQ_EQ),
	192:  uint16(anon_sym_BANG_EQ),
	193:  uint16(anon_sym_EQ_TILDE),
	194:  uint16(anon_sym_LPAREN),
	195:  uint16(anon_sym_RPAREN),
	196:  uint16(anon_sym_COLON),
	197:  uint16(anon_sym_DOLLAR),
	198:  uint16(anon_sym_STAR),
	199:  uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	200:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	201:  uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	202:  uint16(anon_sym_RBRACE_RBRACE),
	203:  uint16(sym_numeric_error),
	204:  uint16(3),
	205:  uint16(29),
	206:  uint16(1),
	207:  uint16(sym_comment),
	208:  uint16(53),
	209:  uint16(5),
	210:  uint16(anon_sym_if),
	211:  uint16(aux_sym_string_token1),
	212:  uint16(anon_sym_DQUOTE),
	213:  uint16(anon_sym_BQUOTE),
	214:  uint16(sym_identifier),
	215:  uint16(55),
	216:  uint16(17),
	217:  uint16(anon_sym_COMMA),
	218:  uint16(anon_sym_SLASH),
	219:  uint16(anon_sym_PLUS),
	220:  uint16(anon_sym_LBRACE),
	221:  uint16(anon_sym_EQ_EQ),
	222:  uint16(anon_sym_BANG_EQ),
	223:  uint16(anon_sym_EQ_TILDE),
	224:  uint16(anon_sym_LPAREN),
	225:  uint16(anon_sym_RPAREN),
	226:  uint16(anon_sym_COLON),
	227:  uint16(anon_sym_DOLLAR),
	228:  uint16(anon_sym_STAR),
	229:  uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	230:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	231:  uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	232:  uint16(anon_sym_RBRACE_RBRACE),
	233:  uint16(sym_numeric_error),
	234:  uint16(19),
	235:  uint16(29),
	236:  uint16(1),
	237:  uint16(sym_comment),
	238:  uint16(57),
	239:  uint16(1),
	240:  uint16(sym_identifier),
	241:  uint16(59),
	242:  uint16(1),
	243:  uint16(anon_sym_SLASH),
	244:  uint16(61),
	245:  uint16(1),
	246:  uint16(anon_sym_if),
	247:  uint16(63),
	248:  uint16(1),
	249:  uint16(anon_sym_LPAREN),
	250:  uint16(65),
	251:  uint16(1),
	252:  uint16(anon_sym_RPAREN),
	253:  uint16(67),
	254:  uint16(1),
	255:  uint16(aux_sym_string_token1),
	256:  uint16(69),
	257:  uint16(1),
	258:  uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	259:  uint16(71),
	260:  uint16(1),
	261:  uint16(anon_sym_DQUOTE),
	262:  uint16(73),
	263:  uint16(1),
	264:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	265:  uint16(75),
	266:  uint16(1),
	267:  uint16(anon_sym_BQUOTE),
	268:  uint16(77),
	269:  uint16(1),
	270:  uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	271:  uint16(79),
	272:  uint16(1),
	273:  uint16(sym_numeric_error),
	274:  uint16(94),
	275:  uint16(1),
	276:  uint16(sym__expression_inner),
	277:  uint16(283),
	278:  uint16(1),
	279:  uint16(sym_expression),
	280:  uint16(394),
	281:  uint16(1),
	282:  uint16(sym_sequence),
	283:  uint16(7),
	284:  uint16(2),
	285:  uint16(sym__backticked),
	286:  uint16(sym__indented_backticked),
	287:  uint16(53),
	288:  uint16(2),
	289:  uint16(sym_if_expression),
	290:  uint16(sym_value),
	291:  uint16(11),
	292:  uint16(3),
	293:  uint16(sym_function_call),
	294:  uint16(sym_external_command),
	295:  uint16(sym_string),
	296:  uint16(18),
	297:  uint16(29),
	298:  uint16(1),
	299:  uint16(sym_comment),
	300:  uint16(61),
	301:  uint16(1),
	302:  uint16(anon_sym_if),
	303:  uint16(63),
	304:  uint16(1),
	305:  uint16(anon_sym_LPAREN),
	306:  uint16(75),
	307:  uint16(1),
	308:  uint16(anon_sym_BQUOTE),
	309:  uint16(77),
	310:  uint16(1),
	311:  uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	312:  uint16(79),
	313:  uint16(1),
	314:  uint16(sym_numeric_error),
	315:  uint16(81),
	316:  uint16(1),
	317:  uint16(sym_identifier),
	318:  uint16(83),
	319:  uint16(1),
	320:  uint16(anon_sym_SLASH),
	321:  uint16(85),
	322:  uint16(1),
	323:  uint16(anon_sym_RPAREN),
	324:  uint16(87),
	325:  uint16(1),
	326:  uint16(aux_sym_string_token1),
	327:  uint16(89),
	328:  uint16(1),
	329:  uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	330:  uint16(91),
	331:  uint16(1),
	332:  uint16(anon_sym_DQUOTE),
	333:  uint16(93),
	334:  uint16(1),
	335:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	336:  uint16(83),
	337:  uint16(1),
	338:  uint16(sym__expression_inner),
	339:  uint16(7),
	340:  uint16(2),
	341:  uint16(sym__backticked),
	342:  uint16(sym__indented_backticked),
	343:  uint16(12),
	344:  uint16(2),
	345:  uint16(sym_expression),
	346:  uint16(aux_sym_dependency_expression_repeat1),
	347:  uint16(53),
	348:  uint16(2),
	349:  uint16(sym_if_expression),
	350:  uint16(sym_value),
	351:  uint16(11),
	352:  uint16(3),
	353:  uint16(sym_function_call),
	354:  uint16(sym_external_command),
	355:  uint16(sym_string),
	356:  uint16(3),
	357:  uint16(29),
	358:  uint16(1),
	359:  uint16(sym_comment),
	360:  uint16(95),
	361:  uint16(5),
	362:  uint16(anon_sym_if),
	363:  uint16(aux_sym_string_token1),
	364:  uint16(anon_sym_DQUOTE),
	365:  uint16(anon_sym_BQUOTE),
	366:  uint16(sym_identifier),
	367:  uint16(97),
	368:  uint16(17),
	369:  uint16(anon_sym_COMMA),
	370:  uint16(anon_sym_SLASH),
	371:  uint16(anon_sym_PLUS),
	372:  uint16(anon_sym_LBRACE),
	373:  uint16(anon_sym_EQ_EQ),
	374:  uint16(anon_sym_BANG_EQ),
	375:  uint16(anon_sym_EQ_TILDE),
	376:  uint16(anon_sym_LPAREN),
	377:  uint16(anon_sym_RPAREN),
	378:  uint16(anon_sym_COLON),
	379:  uint16(anon_sym_DOLLAR),
	380:  uint16(anon_sym_STAR),
	381:  uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	382:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	383:  uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	384:  uint16(anon_sym_RBRACE_RBRACE),
	385:  uint16(sym_numeric_error),
	386:  uint16(18),
	387:  uint16(29),
	388:  uint16(1),
	389:  uint16(sym_comment),
	390:  uint16(99),
	391:  uint16(1),
	392:  uint16(sym_identifier),
	393:  uint16(102),
	394:  uint16(1),
	395:  uint16(anon_sym_SLASH),
	396:  uint16(105),
	397:  uint16(1),
	398:  uint16(anon_sym_if),
	399:  uint16(108),
	400:  uint16(1),
	401:  uint16(anon_sym_LPAREN),
	402:  uint16(111),
	403:  uint16(1),
	404:  uint16(anon_sym_RPAREN),
	405:  uint16(113),
	406:  uint16(1),
	407:  uint16(aux_sym_string_token1),
	408:  uint16(116),
	409:  uint16(1),
	410:  uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	411:  uint16(119),
	412:  uint16(1),
	413:  uint16(anon_sym_DQUOTE),
	414:  uint16(122),
	415:  uint16(1),
	416:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	417:  uint16(125),
	418:  uint16(1),
	419:  uint16(anon_sym_BQUOTE),
	420:  uint16(128),
	421:  uint16(1),
	422:  uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	423:  uint16(131),
	424:  uint16(1),
	425:  uint16(sym_numeric_error),
	426:  uint16(83),
	427:  uint16(1),
	428:  uint16(sym__expression_inner),
	429:  uint16(7),
	430:  uint16(2),
	431:  uint16(sym__backticked),
	432:  uint16(sym__indented_backticked),
	433:  uint16(12),
	434:  uint16(2),
	435:  uint16(sym_expression),
	436:  uint16(aux_sym_dependency_expression_repeat1),
	437:  uint16(53),
	438:  uint16(2),
	439:  uint16(sym_if_expression),
	440:  uint16(sym_value),
	441:  uint16(11),
	442:  uint16(3),
	443:  uint16(sym_function_call),
	444:  uint16(sym_external_command),
	445:  uint16(sym_string),
	446:  uint16(3),
	447:  uint16(29),
	448:  uint16(1),
	449:  uint16(sym_comment),
	450:  uint16(134),
	451:  uint16(5),
	452:  uint16(anon_sym_if),
	453:  uint16(aux_sym_string_token1),
	454:  uint16(anon_sym_DQUOTE),
	455:  uint16(anon_sym_BQUOTE),
	456:  uint16(sym_identifier),
	457:  uint16(136),
	458:  uint16(17),
	459:  uint16(anon_sym_COMMA),
	460:  uint16(anon_sym_SLASH),
	461:  uint16(anon_sym_PLUS),
	462:  uint16(anon_sym_LBRACE),
	463:  uint16(anon_sym_EQ_EQ),
	464:  uint16(anon_sym_BANG_EQ),
	465:  uint16(anon_sym_EQ_TILDE),
	466:  uint16(anon_sym_LPAREN),
	467:  uint16(anon_sym_RPAREN),
	468:  uint16(anon_sym_COLON),
	469:  uint16(anon_sym_DOLLAR),
	470:  uint16(anon_sym_STAR),
	471:  uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	472:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	473:  uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	474:  uint16(anon_sym_RBRACE_RBRACE),
	475:  uint16(sym_numeric_error),
	476:  uint16(3),
	477:  uint16(29),
	478:  uint16(1),
	479:  uint16(sym_comment),
	480:  uint16(138),
	481:  uint16(5),
	482:  uint16(anon_sym_if),
	483:  uint16(aux_sym_string_token1),
	484:  uint16(anon_sym_DQUOTE),
	485:  uint16(anon_sym_BQUOTE),
	486:  uint16(sym_identifier),
	487:  uint16(140),
	488:  uint16(17),
	489:  uint16(anon_sym_COMMA),
	490:  uint16(anon_sym_SLASH),
	491:  uint16(anon_sym_PLUS),
	492:  uint16(anon_sym_LBRACE),
	493:  uint16(anon_sym_EQ_EQ),
	494:  uint16(anon_sym_BANG_EQ),
	495:  uint16(anon_sym_EQ_TILDE),
	496:  uint16(anon_sym_LPAREN),
	497:  uint16(anon_sym_RPAREN),
	498:  uint16(anon_sym_COLON),
	499:  uint16(anon_sym_DOLLAR),
	500:  uint16(anon_sym_STAR),
	501:  uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	502:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	503:  uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	504:  uint16(anon_sym_RBRACE_RBRACE),
	505:  uint16(sym_numeric_error),
	506:  uint16(3),
	507:  uint16(29),
	508:  uint16(1),
	509:  uint16(sym_comment),
	510:  uint16(142),
	511:  uint16(5),
	512:  uint16(anon_sym_if),
	513:  uint16(aux_sym_string_token1),
	514:  uint16(anon_sym_DQUOTE),
	515:  uint16(anon_sym_BQUOTE),
	516:  uint16(sym_identifier),
	517:  uint16(144),
	518:  uint16(17),
	519:  uint16(anon_sym_COMMA),
	520:  uint16(anon_sym_SLASH),
	521:  uint16(anon_sym_PLUS),
	522:  uint16(anon_sym_LBRACE),
	523:  uint16(anon_sym_EQ_EQ),
	524:  uint16(anon_sym_BANG_EQ),
	525:  uint16(anon_sym_EQ_TILDE),
	526:  uint16(anon_sym_LPAREN),
	527:  uint16(anon_sym_RPAREN),
	528:  uint16(anon_sym_COLON),
	529:  uint16(anon_sym_DOLLAR),
	530:  uint16(anon_sym_STAR),
	531:  uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	532:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	533:  uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	534:  uint16(anon_sym_RBRACE_RBRACE),
	535:  uint16(sym_numeric_error),
	536:  uint16(3),
	537:  uint16(29),
	538:  uint16(1),
	539:  uint16(sym_comment),
	540:  uint16(146),
	541:  uint16(5),
	542:  uint16(anon_sym_if),
	543:  uint16(aux_sym_string_token1),
	544:  uint16(anon_sym_DQUOTE),
	545:  uint16(anon_sym_BQUOTE),
	546:  uint16(sym_identifier),
	547:  uint16(148),
	548:  uint16(17),
	549:  uint16(anon_sym_COMMA),
	550:  uint16(anon_sym_SLASH),
	551:  uint16(anon_sym_PLUS),
	552:  uint16(anon_sym_LBRACE),
	553:  uint16(anon_sym_EQ_EQ),
	554:  uint16(anon_sym_BANG_EQ),
	555:  uint16(anon_sym_EQ_TILDE),
	556:  uint16(anon_sym_LPAREN),
	557:  uint16(anon_sym_RPAREN),
	558:  uint16(anon_sym_COLON),
	559:  uint16(anon_sym_DOLLAR),
	560:  uint16(anon_sym_STAR),
	561:  uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	562:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	563:  uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	564:  uint16(anon_sym_RBRACE_RBRACE),
	565:  uint16(sym_numeric_error),
	566:  uint16(3),
	567:  uint16(29),
	568:  uint16(1),
	569:  uint16(sym_comment),
	570:  uint16(150),
	571:  uint16(5),
	572:  uint16(anon_sym_if),
	573:  uint16(aux_sym_string_token1),
	574:  uint16(anon_sym_DQUOTE),
	575:  uint16(anon_sym_BQUOTE),
	576:  uint16(sym_identifier),
	577:  uint16(152),
	578:  uint16(17),
	579:  uint16(anon_sym_COMMA),
	580:  uint16(anon_sym_SLASH),
	581:  uint16(anon_sym_PLUS),
	582:  uint16(anon_sym_LBRACE),
	583:  uint16(anon_sym_EQ_EQ),
	584:  uint16(anon_sym_BANG_EQ),
	585:  uint16(anon_sym_EQ_TILDE),
	586:  uint16(anon_sym_LPAREN),
	587:  uint16(anon_sym_RPAREN),
	588:  uint16(anon_sym_COLON),
	589:  uint16(anon_sym_DOLLAR),
	590:  uint16(anon_sym_STAR),
	591:  uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	592:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	593:  uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	594:  uint16(anon_sym_RBRACE_RBRACE),
	595:  uint16(sym_numeric_error),
	596:  uint16(18),
	597:  uint16(29),
	598:  uint16(1),
	599:  uint16(sym_comment),
	600:  uint16(61),
	601:  uint16(1),
	602:  uint16(anon_sym_if),
	603:  uint16(63),
	604:  uint16(1),
	605:  uint16(anon_sym_LPAREN),
	606:  uint16(75),
	607:  uint16(1),
	608:  uint16(anon_sym_BQUOTE),
	609:  uint16(77),
	610:  uint16(1),
	611:  uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	612:  uint16(79),
	613:  uint16(1),
	614:  uint16(sym_numeric_error),
	615:  uint16(81),
	616:  uint16(1),
	617:  uint16(sym_identifier),
	618:  uint16(83),
	619:  uint16(1),
	620:  uint16(anon_sym_SLASH),
	621:  uint16(87),
	622:  uint16(1),
	623:  uint16(aux_sym_string_token1),
	624:  uint16(89),
	625:  uint16(1),
	626:  uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	627:  uint16(91),
	628:  uint16(1),
	629:  uint16(anon_sym_DQUOTE),
	630:  uint16(93),
	631:  uint16(1),
	632:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	633:  uint16(154),
	634:  uint16(1),
	635:  uint16(anon_sym_RPAREN),
	636:  uint16(83),
	637:  uint16(1),
	638:  uint16(sym__expression_inner),
	639:  uint16(7),
	640:  uint16(2),
	641:  uint16(sym__backticked),
	642:  uint16(sym__indented_backticked),
	643:  uint16(10),
	644:  uint16(2),
	645:  uint16(sym_expression),
	646:  uint16(aux_sym_dependency_expression_repeat1),
	647:  uint16(53),
	648:  uint16(2),
	649:  uint16(sym_if_expression),
	650:  uint16(sym_value),
	651:  uint16(11),
	652:  uint16(3),
	653:  uint16(sym_function_call),
	654:  uint16(sym_external_command),
	655:  uint16(sym_string),
	656:  uint16(3),
	657:  uint16(29),
	658:  uint16(1),
	659:  uint16(sym_comment),
	660:  uint16(156),
	661:  uint16(5),
	662:  uint16(anon_sym_if),
	663:  uint16(aux_sym_string_token1),
	664:  uint16(anon_sym_DQUOTE),
	665:  uint16(anon_sym_BQUOTE),
	666:  uint16(sym_identifier),
	667:  uint16(158),
	668:  uint16(17),
	669:  uint16(anon_sym_COMMA),
	670:  uint16(anon_sym_SLASH),
	671:  uint16(anon_sym_PLUS),
	672:  uint16(anon_sym_LBRACE),
	673:  uint16(anon_sym_EQ_EQ),
	674:  uint16(anon_sym_BANG_EQ),
	675:  uint16(anon_sym_EQ_TILDE),
	676:  uint16(anon_sym_LPAREN),
	677:  uint16(anon_sym_RPAREN),
	678:  uint16(anon_sym_COLON),
	679:  uint16(anon_sym_DOLLAR),
	680:  uint16(anon_sym_STAR),
	681:  uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	682:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	683:  uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	684:  uint16(anon_sym_RBRACE_RBRACE),
	685:  uint16(sym_numeric_error),
	686:  uint16(6),
	687:  uint16(29),
	688:  uint16(1),
	689:  uint16(sym_comment),
	690:  uint16(164),
	691:  uint16(1),
	692:  uint16(anon_sym_else),
	693:  uint16(20),
	694:  uint16(1),
	695:  uint16(aux_sym_if_expression_repeat1),
	696:  uint16(39),
	697:  uint16(1),
	698:  uint16(sym_else_if_clause),
	699:  uint16(160),
	700:  uint16(5),
	701:  uint16(anon_sym_if),
	702:  uint16(aux_sym_string_token1),
	703:  uint16(anon_sym_DQUOTE),
	704:  uint16(anon_sym_BQUOTE),
	705:  uint16(sym_identifier),
	706:  uint16(162),
	707:  uint16(14),
	708:  uint16(anon_sym_COMMA),
	709:  uint16(anon_sym_SLASH),
	710:  uint16(anon_sym_PLUS),
	711:  uint16(anon_sym_LBRACE),
	712:  uint16(anon_sym_EQ_EQ),
	713:  uint16(anon_sym_BANG_EQ),
	714:  uint16(anon_sym_EQ_TILDE),
	715:  uint16(anon_sym_LPAREN),
	716:  uint16(anon_sym_RPAREN),
	717:  uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	718:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	719:  uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	720:  uint16(anon_sym_RBRACE_RBRACE),
	721:  uint16(sym_numeric_error),
	722:  uint16(19),
	723:  uint16(29),
	724:  uint16(1),
	725:  uint16(sym_comment),
	726:  uint16(57),
	727:  uint16(1),
	728:  uint16(sym_identifier),
	729:  uint16(59),
	730:  uint16(1),
	731:  uint16(anon_sym_SLASH),
	732:  uint16(61),
	733:  uint16(1),
	734:  uint16(anon_sym_if),
	735:  uint16(63),
	736:  uint16(1),
	737:  uint16(anon_sym_LPAREN),
	738:  uint16(67),
	739:  uint16(1),
	740:  uint16(aux_sym_string_token1),
	741:  uint16(69),
	742:  uint16(1),
	743:  uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	744:  uint16(71),
	745:  uint16(1),
	746:  uint16(anon_sym_DQUOTE),
	747:  uint16(73),
	748:  uint16(1),
	749:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	750:  uint16(75),
	751:  uint16(1),
	752:  uint16(anon_sym_BQUOTE),
	753:  uint16(77),
	754:  uint16(1),
	755:  uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	756:  uint16(79),
	757:  uint16(1),
	758:  uint16(sym_numeric_error),
	759:  uint16(167),
	760:  uint16(1),
	761:  uint16(anon_sym_RPAREN),
	762:  uint16(94),
	763:  uint16(1),
	764:  uint16(sym__expression_inner),
	765:  uint16(283),
	766:  uint16(1),
	767:  uint16(sym_expression),
	768:  uint16(366),
	769:  uint16(1),
	770:  uint16(sym_sequence),
	771:  uint16(7),
	772:  uint16(2),
	773:  uint16(sym__backticked),
	774:  uint16(sym__indented_backticked),
	775:  uint16(53),
	776:  uint16(2),
	777:  uint16(sym_if_expression),
	778:  uint16(sym_value),
	779:  uint16(11),
	780:  uint16(3),
	781:  uint16(sym_function_call),
	782:  uint16(sym_external_command),
	783:  uint16(sym_string),
	784:  uint16(19),
	785:  uint16(29),
	786:  uint16(1),
	787:  uint16(sym_comment),
	788:  uint16(57),
	789:  uint16(1),
	790:  uint16(sym_identifier),
	791:  uint16(59),
	792:  uint16(1),
	793:  uint16(anon_sym_SLASH),
	794:  uint16(61),
	795:  uint16(1),
	796:  uint16(anon_sym_if),
	797:  uint16(63),
	798:  uint16(1),
	799:  uint16(anon_sym_LPAREN),
	800:  uint16(67),
	801:  uint16(1),
	802:  uint16(aux_sym_string_token1),
	803:  uint16(69),
	804:  uint16(1),
	805:  uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	806:  uint16(71),
	807:  uint16(1),
	808:  uint16(anon_sym_DQUOTE),
	809:  uint16(73),
	810:  uint16(1),
	811:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	812:  uint16(75),
	813:  uint16(1),
	814:  uint16(anon_sym_BQUOTE),
	815:  uint16(77),
	816:  uint16(1),
	817:  uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	818:  uint16(79),
	819:  uint16(1),
	820:  uint16(sym_numeric_error),
	821:  uint16(169),
	822:  uint16(1),
	823:  uint16(anon_sym_RPAREN),
	824:  uint16(94),
	825:  uint16(1),
	826:  uint16(sym__expression_inner),
	827:  uint16(283),
	828:  uint16(1),
	829:  uint16(sym_expression),
	830:  uint16(387),
	831:  uint16(1),
	832:  uint16(sym_sequence),
	833:  uint16(7),
	834:  uint16(2),
	835:  uint16(sym__backticked),
	836:  uint16(sym__indented_backticked),
	837:  uint16(53),
	838:  uint16(2),
	839:  uint16(sym_if_expression),
	840:  uint16(sym_value),
	841:  uint16(11),
	842:  uint16(3),
	843:  uint16(sym_function_call),
	844:  uint16(sym_external_command),
	845:  uint16(sym_string),
	846:  uint16(13),
	847:  uint16(7),
	848:  uint16(1),
	849:  uint16(sym_identifier),
	850:  uint16(9),
	851:  uint16(1),
	852:  uint16(anon_sym_alias),
	853:  uint16(11),
	854:  uint16(1),
	855:  uint16(anon_sym_export),
	856:  uint16(13),
	857:  uint16(1),
	858:  uint16(anon_sym_import),
	859:  uint16(15),
	860:  uint16(1),
	861:  uint16(anon_sym_mod),
	862:  uint16(17),
	863:  uint16(1),
	864:  uint16(anon_sym_set),
	865:  uint16(19),
	866:  uint16(1),
	867:  uint16(anon_sym_LBRACK),
	868:  uint16(21),
	869:  uint16(1),
	870:  uint16(anon_sym_AT),
	871:  uint16(29),
	872:  uint16(1),
	873:  uint16(sym_comment),
	874:  uint16(171),
	875:  uint16(1),
	877:  uint16(327),
	878:  uint16(1),
	879:  uint16(sym_recipe_header),
	880:  uint16(127),
	881:  uint16(2),
	882:  uint16(sym_attribute),
	883:  uint16(aux_sym_alias_repeat1),
	884:  uint16(29),
	885:  uint16(9),
	886:  uint16(sym__item),
	887:  uint16(sym_alias),
	888:  uint16(sym_assignment),
	889:  uint16(sym_export),
	890:  uint16(sym_import),
	891:  uint16(sym_module),
	892:  uint16(sym_setting),
	893:  uint16(sym_recipe),
	894:  uint16(aux_sym_source_file_repeat1),
	895:  uint16(18),
	896:  uint16(29),
	897:  uint16(1),
	898:  uint16(sym_comment),
	899:  uint16(57),
	900:  uint16(1),
	901:  uint16(sym_identifier),
	902:  uint16(59),
	903:  uint16(1),
	904:  uint16(anon_sym_SLASH),
	905:  uint16(61),
	906:  uint16(1),
	907:  uint16(anon_sym_if),
	908:  uint16(63),
	909:  uint16(1),
	910:  uint16(anon_sym_LPAREN),
	911:  uint16(67),
	912:  uint16(1),
	913:  uint16(aux_sym_string_token1),
	914:  uint16(69),
	915:  uint16(1),
	916:  uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	917:  uint16(71),
	918:  uint16(1),
	919:  uint16(anon_sym_DQUOTE),
	920:  uint16(73),
	921:  uint16(1),
	922:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	923:  uint16(75),
	924:  uint16(1),
	925:  uint16(anon_sym_BQUOTE),
	926:  uint16(77),
	927:  uint16(1),
	928:  uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	929:  uint16(79),
	930:  uint16(1),
	931:  uint16(sym_numeric_error),
	932:  uint16(94),
	933:  uint16(1),
	934:  uint16(sym__expression_inner),
	935:  uint16(201),
	936:  uint16(1),
	937:  uint16(sym_expression),
	938:  uint16(303),
	939:  uint16(1),
	940:  uint16(sym_condition),
	941:  uint16(7),
	942:  uint16(2),
	943:  uint16(sym__backticked),
	944:  uint16(sym__indented_backticked),
	945:  uint16(53),
	946:  uint16(2),
	947:  uint16(sym_if_expression),
	948:  uint16(sym_value),
	949:  uint16(11),
	950:  uint16(3),
	951:  uint16(sym_function_call),
	952:  uint16(sym_external_command),
	953:  uint16(sym_string),
	954:  uint16(18),
	955:  uint16(29),
	956:  uint16(1),
	957:  uint16(sym_comment),
	958:  uint16(57),
	959:  uint16(1),
	960:  uint16(sym_identifier),
	961:  uint16(59),
	962:  uint16(1),
	963:  uint16(anon_sym_SLASH),
	964:  uint16(61),
	965:  uint16(1),
	966:  uint16(anon_sym_if),
	967:  uint16(63),
	968:  uint16(1),
	969:  uint16(anon_sym_LPAREN),
	970:  uint16(67),
	971:  uint16(1),
	972:  uint16(aux_sym_string_token1),
	973:  uint16(69),
	974:  uint16(1),
	975:  uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	976:  uint16(71),
	977:  uint16(1),
	978:  uint16(anon_sym_DQUOTE),
	979:  uint16(73),
	980:  uint16(1),
	981:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	982:  uint16(75),
	983:  uint16(1),
	984:  uint16(anon_sym_BQUOTE),
	985:  uint16(77),
	986:  uint16(1),
	987:  uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	988:  uint16(79),
	989:  uint16(1),
	990:  uint16(sym_numeric_error),
	991:  uint16(94),
	992:  uint16(1),
	993:  uint16(sym__expression_inner),
	994:  uint16(223),
	995:  uint16(1),
	996:  uint16(sym_string),
	997:  uint16(7),
	998:  uint16(2),
	999:  uint16(sym__backticked),
	1000: uint16(sym__indented_backticked),
	1001: uint16(11),
	1002: uint16(2),
	1003: uint16(sym_function_call),
	1004: uint16(sym_external_command),
	1005: uint16(53),
	1006: uint16(2),
	1007: uint16(sym_if_expression),
	1008: uint16(sym_value),
	1009: uint16(318),
	1010: uint16(2),
	1011: uint16(sym_expression),
	1012: uint16(sym_regex_literal),
	1013: uint16(13),
	1014: uint16(7),
	1015: uint16(1),
	1016: uint16(sym_identifier),
	1017: uint16(9),
	1018: uint16(1),
	1019: uint16(anon_sym_alias),
	1020: uint16(11),
	1021: uint16(1),
	1022: uint16(anon_sym_export),
	1023: uint16(13),
	1024: uint16(1),
	1025: uint16(anon_sym_import),
	1026: uint16(15),
	1027: uint16(1),
	1028: uint16(anon_sym_mod),
	1029: uint16(17),
	1030: uint16(1),
	1031: uint16(anon_sym_set),
	1032: uint16(19),
	1033: uint16(1),
	1034: uint16(anon_sym_LBRACK),
	1035: uint16(21),
	1036: uint16(1),
	1037: uint16(anon_sym_AT),
	1038: uint16(29),
	1039: uint16(1),
	1040: uint16(sym_comment),
	1041: uint16(173),
	1042: uint16(1),
	1044: uint16(327),
	1045: uint16(1),
	1046: uint16(sym_recipe_header),
	1047: uint16(127),
	1048: uint16(2),
	1049: uint16(sym_attribute),
	1050: uint16(aux_sym_alias_repeat1),
	1051: uint16(28),
	1052: uint16(9),
	1053: uint16(sym__item),
	1054: uint16(sym_alias),
	1055: uint16(sym_assignment),
	1056: uint16(sym_export),
	1057: uint16(sym_import),
	1058: uint16(sym_module),
	1059: uint16(sym_setting),
	1060: uint16(sym_recipe),
	1061: uint16(aux_sym_source_file_repeat1),
	1062: uint16(18),
	1063: uint16(29),
	1064: uint16(1),
	1065: uint16(sym_comment),
	1066: uint16(57),
	1067: uint16(1),
	1068: uint16(sym_identifier),
	1069: uint16(59),
	1070: uint16(1),
	1071: uint16(anon_sym_SLASH),
	1072: uint16(61),
	1073: uint16(1),
	1074: uint16(anon_sym_if),
	1075: uint16(63),
	1076: uint16(1),
	1077: uint16(anon_sym_LPAREN),
	1078: uint16(67),
	1079: uint16(1),
	1080: uint16(aux_sym_string_token1),
	1081: uint16(69),
	1082: uint16(1),
	1083: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	1084: uint16(71),
	1085: uint16(1),
	1086: uint16(anon_sym_DQUOTE),
	1087: uint16(73),
	1088: uint16(1),
	1089: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1090: uint16(75),
	1091: uint16(1),
	1092: uint16(anon_sym_BQUOTE),
	1093: uint16(77),
	1094: uint16(1),
	1095: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	1096: uint16(79),
	1097: uint16(1),
	1098: uint16(sym_numeric_error),
	1099: uint16(94),
	1100: uint16(1),
	1101: uint16(sym__expression_inner),
	1102: uint16(201),
	1103: uint16(1),
	1104: uint16(sym_expression),
	1105: uint16(299),
	1106: uint16(1),
	1107: uint16(sym_condition),
	1108: uint16(7),
	1109: uint16(2),
	1110: uint16(sym__backticked),
	1111: uint16(sym__indented_backticked),
	1112: uint16(53),
	1113: uint16(2),
	1114: uint16(sym_if_expression),
	1115: uint16(sym_value),
	1116: uint16(11),
	1117: uint16(3),
	1118: uint16(sym_function_call),
	1119: uint16(sym_external_command),
	1120: uint16(sym_string),
	1121: uint16(13),
	1122: uint16(29),
	1123: uint16(1),
	1124: uint16(sym_comment),
	1125: uint16(175),
	1126: uint16(1),
	1128: uint16(177),
	1129: uint16(1),
	1130: uint16(sym_identifier),
	1131: uint16(180),
	1132: uint16(1),
	1133: uint16(anon_sym_alias),
	1134: uint16(183),
	1135: uint16(1),
	1136: uint16(anon_sym_export),
	1137: uint16(186),
	1138: uint16(1),
	1139: uint16(anon_sym_import),
	1140: uint16(189),
	1141: uint16(1),
	1142: uint16(anon_sym_mod),
	1143: uint16(192),
	1144: uint16(1),
	1145: uint16(anon_sym_set),
	1146: uint16(195),
	1147: uint16(1),
	1148: uint16(anon_sym_LBRACK),
	1149: uint16(198),
	1150: uint16(1),
	1151: uint16(anon_sym_AT),
	1152: uint16(327),
	1153: uint16(1),
	1154: uint16(sym_recipe_header),
	1155: uint16(127),
	1156: uint16(2),
	1157: uint16(sym_attribute),
	1158: uint16(aux_sym_alias_repeat1),
	1159: uint16(28),
	1160: uint16(9),
	1161: uint16(sym__item),
	1162: uint16(sym_alias),
	1163: uint16(sym_assignment),
	1164: uint16(sym_export),
	1165: uint16(sym_import),
	1166: uint16(sym_module),
	1167: uint16(sym_setting),
	1168: uint16(sym_recipe),
	1169: uint16(aux_sym_source_file_repeat1),
	1170: uint16(13),
	1171: uint16(7),
	1172: uint16(1),
	1173: uint16(sym_identifier),
	1174: uint16(9),
	1175: uint16(1),
	1176: uint16(anon_sym_alias),
	1177: uint16(11),
	1178: uint16(1),
	1179: uint16(anon_sym_export),
	1180: uint16(13),
	1181: uint16(1),
	1182: uint16(anon_sym_import),
	1183: uint16(15),
	1184: uint16(1),
	1185: uint16(anon_sym_mod),
	1186: uint16(17),
	1187: uint16(1),
	1188: uint16(anon_sym_set),
	1189: uint16(19),
	1190: uint16(1),
	1191: uint16(anon_sym_LBRACK),
	1192: uint16(21),
	1193: uint16(1),
	1194: uint16(anon_sym_AT),
	1195: uint16(29),
	1196: uint16(1),
	1197: uint16(sym_comment),
	1198: uint16(201),
	1199: uint16(1),
	1201: uint16(327),
	1202: uint16(1),
	1203: uint16(sym_recipe_header),
	1204: uint16(127),
	1205: uint16(2),
	1206: uint16(sym_attribute),
	1207: uint16(aux_sym_alias_repeat1),
	1208: uint16(28),
	1209: uint16(9),
	1210: uint16(sym__item),
	1211: uint16(sym_alias),
	1212: uint16(sym_assignment),
	1213: uint16(sym_export),
	1214: uint16(sym_import),
	1215: uint16(sym_module),
	1216: uint16(sym_setting),
	1217: uint16(sym_recipe),
	1218: uint16(aux_sym_source_file_repeat1),
	1219: uint16(18),
	1220: uint16(29),
	1221: uint16(1),
	1222: uint16(sym_comment),
	1223: uint16(57),
	1224: uint16(1),
	1225: uint16(sym_identifier),
	1226: uint16(59),
	1227: uint16(1),
	1228: uint16(anon_sym_SLASH),
	1229: uint16(61),
	1230: uint16(1),
	1231: uint16(anon_sym_if),
	1232: uint16(63),
	1233: uint16(1),
	1234: uint16(anon_sym_LPAREN),
	1235: uint16(67),
	1236: uint16(1),
	1237: uint16(aux_sym_string_token1),
	1238: uint16(69),
	1239: uint16(1),
	1240: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	1241: uint16(71),
	1242: uint16(1),
	1243: uint16(anon_sym_DQUOTE),
	1244: uint16(73),
	1245: uint16(1),
	1246: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1247: uint16(75),
	1248: uint16(1),
	1249: uint16(anon_sym_BQUOTE),
	1250: uint16(77),
	1251: uint16(1),
	1252: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	1253: uint16(79),
	1254: uint16(1),
	1255: uint16(sym_numeric_error),
	1256: uint16(94),
	1257: uint16(1),
	1258: uint16(sym__expression_inner),
	1259: uint16(201),
	1260: uint16(1),
	1261: uint16(sym_expression),
	1262: uint16(306),
	1263: uint16(1),
	1264: uint16(sym_condition),
	1265: uint16(7),
	1266: uint16(2),
	1267: uint16(sym__backticked),
	1268: uint16(sym__indented_backticked),
	1269: uint16(53),
	1270: uint16(2),
	1271: uint16(sym_if_expression),
	1272: uint16(sym_value),
	1273: uint16(11),
	1274: uint16(3),
	1275: uint16(sym_function_call),
	1276: uint16(sym_external_command),
	1277: uint16(sym_string),
	1278: uint16(18),
	1279: uint16(29),
	1280: uint16(1),
	1281: uint16(sym_comment),
	1282: uint16(57),
	1283: uint16(1),
	1284: uint16(sym_identifier),
	1285: uint16(59),
	1286: uint16(1),
	1287: uint16(anon_sym_SLASH),
	1288: uint16(61),
	1289: uint16(1),
	1290: uint16(anon_sym_if),
	1291: uint16(63),
	1292: uint16(1),
	1293: uint16(anon_sym_LPAREN),
	1294: uint16(67),
	1295: uint16(1),
	1296: uint16(aux_sym_string_token1),
	1297: uint16(69),
	1298: uint16(1),
	1299: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	1300: uint16(71),
	1301: uint16(1),
	1302: uint16(anon_sym_DQUOTE),
	1303: uint16(73),
	1304: uint16(1),
	1305: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1306: uint16(75),
	1307: uint16(1),
	1308: uint16(anon_sym_BQUOTE),
	1309: uint16(77),
	1310: uint16(1),
	1311: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	1312: uint16(79),
	1313: uint16(1),
	1314: uint16(sym_numeric_error),
	1315: uint16(94),
	1316: uint16(1),
	1317: uint16(sym__expression_inner),
	1318: uint16(201),
	1319: uint16(1),
	1320: uint16(sym_expression),
	1321: uint16(307),
	1322: uint16(1),
	1323: uint16(sym_condition),
	1324: uint16(7),
	1325: uint16(2),
	1326: uint16(sym__backticked),
	1327: uint16(sym__indented_backticked),
	1328: uint16(53),
	1329: uint16(2),
	1330: uint16(sym_if_expression),
	1331: uint16(sym_value),
	1332: uint16(11),
	1333: uint16(3),
	1334: uint16(sym_function_call),
	1335: uint16(sym_external_command),
	1336: uint16(sym_string),
	1337: uint16(18),
	1338: uint16(29),
	1339: uint16(1),
	1340: uint16(sym_comment),
	1341: uint16(57),
	1342: uint16(1),
	1343: uint16(sym_identifier),
	1344: uint16(59),
	1345: uint16(1),
	1346: uint16(anon_sym_SLASH),
	1347: uint16(61),
	1348: uint16(1),
	1349: uint16(anon_sym_if),
	1350: uint16(63),
	1351: uint16(1),
	1352: uint16(anon_sym_LPAREN),
	1353: uint16(67),
	1354: uint16(1),
	1355: uint16(aux_sym_string_token1),
	1356: uint16(69),
	1357: uint16(1),
	1358: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	1359: uint16(71),
	1360: uint16(1),
	1361: uint16(anon_sym_DQUOTE),
	1362: uint16(73),
	1363: uint16(1),
	1364: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1365: uint16(75),
	1366: uint16(1),
	1367: uint16(anon_sym_BQUOTE),
	1368: uint16(77),
	1369: uint16(1),
	1370: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	1371: uint16(79),
	1372: uint16(1),
	1373: uint16(sym_numeric_error),
	1374: uint16(94),
	1375: uint16(1),
	1376: uint16(sym__expression_inner),
	1377: uint16(201),
	1378: uint16(1),
	1379: uint16(sym_expression),
	1380: uint16(313),
	1381: uint16(1),
	1382: uint16(sym_condition),
	1383: uint16(7),
	1384: uint16(2),
	1385: uint16(sym__backticked),
	1386: uint16(sym__indented_backticked),
	1387: uint16(53),
	1388: uint16(2),
	1389: uint16(sym_if_expression),
	1390: uint16(sym_value),
	1391: uint16(11),
	1392: uint16(3),
	1393: uint16(sym_function_call),
	1394: uint16(sym_external_command),
	1395: uint16(sym_string),
	1396: uint16(18),
	1397: uint16(29),
	1398: uint16(1),
	1399: uint16(sym_comment),
	1400: uint16(57),
	1401: uint16(1),
	1402: uint16(sym_identifier),
	1403: uint16(59),
	1404: uint16(1),
	1405: uint16(anon_sym_SLASH),
	1406: uint16(61),
	1407: uint16(1),
	1408: uint16(anon_sym_if),
	1409: uint16(63),
	1410: uint16(1),
	1411: uint16(anon_sym_LPAREN),
	1412: uint16(67),
	1413: uint16(1),
	1414: uint16(aux_sym_string_token1),
	1415: uint16(69),
	1416: uint16(1),
	1417: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	1418: uint16(71),
	1419: uint16(1),
	1420: uint16(anon_sym_DQUOTE),
	1421: uint16(73),
	1422: uint16(1),
	1423: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1424: uint16(75),
	1425: uint16(1),
	1426: uint16(anon_sym_BQUOTE),
	1427: uint16(77),
	1428: uint16(1),
	1429: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	1430: uint16(79),
	1431: uint16(1),
	1432: uint16(sym_numeric_error),
	1433: uint16(94),
	1434: uint16(1),
	1435: uint16(sym__expression_inner),
	1436: uint16(201),
	1437: uint16(1),
	1438: uint16(sym_expression),
	1439: uint16(314),
	1440: uint16(1),
	1441: uint16(sym_condition),
	1442: uint16(7),
	1443: uint16(2),
	1444: uint16(sym__backticked),
	1445: uint16(sym__indented_backticked),
	1446: uint16(53),
	1447: uint16(2),
	1448: uint16(sym_if_expression),
	1449: uint16(sym_value),
	1450: uint16(11),
	1451: uint16(3),
	1452: uint16(sym_function_call),
	1453: uint16(sym_external_command),
	1454: uint16(sym_string),
	1455: uint16(17),
	1456: uint16(29),
	1457: uint16(1),
	1458: uint16(sym_comment),
	1459: uint16(203),
	1460: uint16(1),
	1461: uint16(sym_identifier),
	1462: uint16(205),
	1463: uint16(1),
	1464: uint16(anon_sym_SLASH),
	1465: uint16(207),
	1466: uint16(1),
	1467: uint16(anon_sym_if),
	1468: uint16(209),
	1469: uint16(1),
	1470: uint16(anon_sym_LPAREN),
	1471: uint16(211),
	1472: uint16(1),
	1473: uint16(aux_sym_string_token1),
	1474: uint16(213),
	1475: uint16(1),
	1476: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	1477: uint16(215),
	1478: uint16(1),
	1479: uint16(anon_sym_DQUOTE),
	1480: uint16(217),
	1481: uint16(1),
	1482: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1483: uint16(219),
	1484: uint16(1),
	1485: uint16(anon_sym_BQUOTE),
	1486: uint16(221),
	1487: uint16(1),
	1488: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	1489: uint16(223),
	1490: uint16(1),
	1491: uint16(sym_numeric_error),
	1492: uint16(238),
	1493: uint16(1),
	1494: uint16(sym__expression_inner),
	1495: uint16(325),
	1496: uint16(1),
	1497: uint16(sym_expression),
	1498: uint16(239),
	1499: uint16(2),
	1500: uint16(sym_if_expression),
	1501: uint16(sym_value),
	1502: uint16(240),
	1503: uint16(2),
	1504: uint16(sym__backticked),
	1505: uint16(sym__indented_backticked),
	1506: uint16(234),
	1507: uint16(3),
	1508: uint16(sym_function_call),
	1509: uint16(sym_external_command),
	1510: uint16(sym_string),
	1511: uint16(17),
	1512: uint16(29),
	1513: uint16(1),
	1514: uint16(sym_comment),
	1515: uint16(87),
	1516: uint16(1),
	1517: uint16(aux_sym_string_token1),
	1518: uint16(89),
	1519: uint16(1),
	1520: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	1521: uint16(91),
	1522: uint16(1),
	1523: uint16(anon_sym_DQUOTE),
	1524: uint16(93),
	1525: uint16(1),
	1526: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1527: uint16(225),
	1528: uint16(1),
	1529: uint16(sym_identifier),
	1530: uint16(227),
	1531: uint16(1),
	1532: uint16(anon_sym_SLASH),
	1533: uint16(229),
	1534: uint16(1),
	1535: uint16(anon_sym_if),
	1536: uint16(231),
	1537: uint16(1),
	1538: uint16(anon_sym_LPAREN),
	1539: uint16(233),
	1540: uint16(1),
	1541: uint16(anon_sym_BQUOTE),
	1542: uint16(235),
	1543: uint16(1),
	1544: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	1545: uint16(237),
	1546: uint16(1),
	1547: uint16(sym_numeric_error),
	1548: uint16(255),
	1549: uint16(1),
	1550: uint16(sym__expression_inner),
	1551: uint16(388),
	1552: uint16(1),
	1553: uint16(sym_expression),
	1554: uint16(256),
	1555: uint16(2),
	1556: uint16(sym_if_expression),
	1557: uint16(sym_value),
	1558: uint16(257),
	1559: uint16(2),
	1560: uint16(sym__backticked),
	1561: uint16(sym__indented_backticked),
	1562: uint16(254),
	1563: uint16(3),
	1564: uint16(sym_function_call),
	1565: uint16(sym_external_command),
	1566: uint16(sym_string),
	1567: uint16(17),
	1568: uint16(29),
	1569: uint16(1),
	1570: uint16(sym_comment),
	1571: uint16(57),
	1572: uint16(1),
	1573: uint16(sym_identifier),
	1574: uint16(59),
	1575: uint16(1),
	1576: uint16(anon_sym_SLASH),
	1577: uint16(61),
	1578: uint16(1),
	1579: uint16(anon_sym_if),
	1580: uint16(63),
	1581: uint16(1),
	1582: uint16(anon_sym_LPAREN),
	1583: uint16(67),
	1584: uint16(1),
	1585: uint16(aux_sym_string_token1),
	1586: uint16(69),
	1587: uint16(1),
	1588: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	1589: uint16(71),
	1590: uint16(1),
	1591: uint16(anon_sym_DQUOTE),
	1592: uint16(73),
	1593: uint16(1),
	1594: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1595: uint16(75),
	1596: uint16(1),
	1597: uint16(anon_sym_BQUOTE),
	1598: uint16(77),
	1599: uint16(1),
	1600: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	1601: uint16(79),
	1602: uint16(1),
	1603: uint16(sym_numeric_error),
	1604: uint16(94),
	1605: uint16(1),
	1606: uint16(sym__expression_inner),
	1607: uint16(347),
	1608: uint16(1),
	1609: uint16(sym_expression),
	1610: uint16(7),
	1611: uint16(2),
	1612: uint16(sym__backticked),
	1613: uint16(sym__indented_backticked),
	1614: uint16(53),
	1615: uint16(2),
	1616: uint16(sym_if_expression),
	1617: uint16(sym_value),
	1618: uint16(11),
	1619: uint16(3),
	1620: uint16(sym_function_call),
	1621: uint16(sym_external_command),
	1622: uint16(sym_string),
	1623: uint16(17),
	1624: uint16(29),
	1625: uint16(1),
	1626: uint16(sym_comment),
	1627: uint16(57),
	1628: uint16(1),
	1629: uint16(sym_identifier),
	1630: uint16(59),
	1631: uint16(1),
	1632: uint16(anon_sym_SLASH),
	1633: uint16(61),
	1634: uint16(1),
	1635: uint16(anon_sym_if),
	1636: uint16(63),
	1637: uint16(1),
	1638: uint16(anon_sym_LPAREN),
	1639: uint16(67),
	1640: uint16(1),
	1641: uint16(aux_sym_string_token1),
	1642: uint16(69),
	1643: uint16(1),
	1644: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	1645: uint16(71),
	1646: uint16(1),
	1647: uint16(anon_sym_DQUOTE),
	1648: uint16(73),
	1649: uint16(1),
	1650: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1651: uint16(75),
	1652: uint16(1),
	1653: uint16(anon_sym_BQUOTE),
	1654: uint16(77),
	1655: uint16(1),
	1656: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	1657: uint16(79),
	1658: uint16(1),
	1659: uint16(sym_numeric_error),
	1660: uint16(94),
	1661: uint16(1),
	1662: uint16(sym__expression_inner),
	1663: uint16(318),
	1664: uint16(1),
	1665: uint16(sym_expression),
	1666: uint16(7),
	1667: uint16(2),
	1668: uint16(sym__backticked),
	1669: uint16(sym__indented_backticked),
	1670: uint16(53),
	1671: uint16(2),
	1672: uint16(sym_if_expression),
	1673: uint16(sym_value),
	1674: uint16(11),
	1675: uint16(3),
	1676: uint16(sym_function_call),
	1677: uint16(sym_external_command),
	1678: uint16(sym_string),
	1679: uint16(17),
	1680: uint16(29),
	1681: uint16(1),
	1682: uint16(sym_comment),
	1683: uint16(57),
	1684: uint16(1),
	1685: uint16(sym_identifier),
	1686: uint16(59),
	1687: uint16(1),
	1688: uint16(anon_sym_SLASH),
	1689: uint16(61),
	1690: uint16(1),
	1691: uint16(anon_sym_if),
	1692: uint16(63),
	1693: uint16(1),
	1694: uint16(anon_sym_LPAREN),
	1695: uint16(67),
	1696: uint16(1),
	1697: uint16(aux_sym_string_token1),
	1698: uint16(69),
	1699: uint16(1),
	1700: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	1701: uint16(71),
	1702: uint16(1),
	1703: uint16(anon_sym_DQUOTE),
	1704: uint16(73),
	1705: uint16(1),
	1706: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1707: uint16(75),
	1708: uint16(1),
	1709: uint16(anon_sym_BQUOTE),
	1710: uint16(77),
	1711: uint16(1),
	1712: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	1713: uint16(79),
	1714: uint16(1),
	1715: uint16(sym_numeric_error),
	1716: uint16(94),
	1717: uint16(1),
	1718: uint16(sym__expression_inner),
	1719: uint16(333),
	1720: uint16(1),
	1721: uint16(sym_expression),
	1722: uint16(7),
	1723: uint16(2),
	1724: uint16(sym__backticked),
	1725: uint16(sym__indented_backticked),
	1726: uint16(53),
	1727: uint16(2),
	1728: uint16(sym_if_expression),
	1729: uint16(sym_value),
	1730: uint16(11),
	1731: uint16(3),
	1732: uint16(sym_function_call),
	1733: uint16(sym_external_command),
	1734: uint16(sym_string),
	1735: uint16(3),
	1736: uint16(29),
	1737: uint16(1),
	1738: uint16(sym_comment),
	1739: uint16(239),
	1740: uint16(6),
	1741: uint16(anon_sym_if),
	1742: uint16(anon_sym_else),
	1743: uint16(aux_sym_string_token1),
	1744: uint16(anon_sym_DQUOTE),
	1745: uint16(anon_sym_BQUOTE),
	1746: uint16(sym_identifier),
	1747: uint16(241),
	1748: uint16(14),
	1749: uint16(anon_sym_COMMA),
	1750: uint16(anon_sym_SLASH),
	1751: uint16(anon_sym_PLUS),
	1752: uint16(anon_sym_LBRACE),
	1753: uint16(anon_sym_EQ_EQ),
	1754: uint16(anon_sym_BANG_EQ),
	1755: uint16(anon_sym_EQ_TILDE),
	1756: uint16(anon_sym_LPAREN),
	1757: uint16(anon_sym_RPAREN),
	1758: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	1759: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1760: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	1761: uint16(anon_sym_RBRACE_RBRACE),
	1762: uint16(sym_numeric_error),
	1763: uint16(17),
	1764: uint16(29),
	1765: uint16(1),
	1766: uint16(sym_comment),
	1767: uint16(87),
	1768: uint16(1),
	1769: uint16(aux_sym_string_token1),
	1770: uint16(89),
	1771: uint16(1),
	1772: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	1773: uint16(91),
	1774: uint16(1),
	1775: uint16(anon_sym_DQUOTE),
	1776: uint16(93),
	1777: uint16(1),
	1778: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1779: uint16(225),
	1780: uint16(1),
	1781: uint16(sym_identifier),
	1782: uint16(227),
	1783: uint16(1),
	1784: uint16(anon_sym_SLASH),
	1785: uint16(229),
	1786: uint16(1),
	1787: uint16(anon_sym_if),
	1788: uint16(231),
	1789: uint16(1),
	1790: uint16(anon_sym_LPAREN),
	1791: uint16(233),
	1792: uint16(1),
	1793: uint16(anon_sym_BQUOTE),
	1794: uint16(235),
	1795: uint16(1),
	1796: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	1797: uint16(237),
	1798: uint16(1),
	1799: uint16(sym_numeric_error),
	1800: uint16(255),
	1801: uint16(1),
	1802: uint16(sym__expression_inner),
	1803: uint16(323),
	1804: uint16(1),
	1805: uint16(sym_expression),
	1806: uint16(256),
	1807: uint16(2),
	1808: uint16(sym_if_expression),
	1809: uint16(sym_value),
	1810: uint16(257),
	1811: uint16(2),
	1812: uint16(sym__backticked),
	1813: uint16(sym__indented_backticked),
	1814: uint16(254),
	1815: uint16(3),
	1816: uint16(sym_function_call),
	1817: uint16(sym_external_command),
	1818: uint16(sym_string),
	1819: uint16(3),
	1820: uint16(29),
	1821: uint16(1),
	1822: uint16(sym_comment),
	1823: uint16(243),
	1824: uint16(6),
	1825: uint16(anon_sym_if),
	1826: uint16(anon_sym_else),
	1827: uint16(aux_sym_string_token1),
	1828: uint16(anon_sym_DQUOTE),
	1829: uint16(anon_sym_BQUOTE),
	1830: uint16(sym_identifier),
	1831: uint16(245),
	1832: uint16(14),
	1833: uint16(anon_sym_COMMA),
	1834: uint16(anon_sym_SLASH),
	1835: uint16(anon_sym_PLUS),
	1836: uint16(anon_sym_LBRACE),
	1837: uint16(anon_sym_EQ_EQ),
	1838: uint16(anon_sym_BANG_EQ),
	1839: uint16(anon_sym_EQ_TILDE),
	1840: uint16(anon_sym_LPAREN),
	1841: uint16(anon_sym_RPAREN),
	1842: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	1843: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1844: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	1845: uint16(anon_sym_RBRACE_RBRACE),
	1846: uint16(sym_numeric_error),
	1847: uint16(17),
	1848: uint16(29),
	1849: uint16(1),
	1850: uint16(sym_comment),
	1851: uint16(57),
	1852: uint16(1),
	1853: uint16(sym_identifier),
	1854: uint16(59),
	1855: uint16(1),
	1856: uint16(anon_sym_SLASH),
	1857: uint16(61),
	1858: uint16(1),
	1859: uint16(anon_sym_if),
	1860: uint16(63),
	1861: uint16(1),
	1862: uint16(anon_sym_LPAREN),
	1863: uint16(67),
	1864: uint16(1),
	1865: uint16(aux_sym_string_token1),
	1866: uint16(69),
	1867: uint16(1),
	1868: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	1869: uint16(71),
	1870: uint16(1),
	1871: uint16(anon_sym_DQUOTE),
	1872: uint16(73),
	1873: uint16(1),
	1874: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1875: uint16(75),
	1876: uint16(1),
	1877: uint16(anon_sym_BQUOTE),
	1878: uint16(77),
	1879: uint16(1),
	1880: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	1881: uint16(79),
	1882: uint16(1),
	1883: uint16(sym_numeric_error),
	1884: uint16(94),
	1885: uint16(1),
	1886: uint16(sym__expression_inner),
	1887: uint16(298),
	1888: uint16(1),
	1889: uint16(sym_expression),
	1890: uint16(7),
	1891: uint16(2),
	1892: uint16(sym__backticked),
	1893: uint16(sym__indented_backticked),
	1894: uint16(53),
	1895: uint16(2),
	1896: uint16(sym_if_expression),
	1897: uint16(sym_value),
	1898: uint16(11),
	1899: uint16(3),
	1900: uint16(sym_function_call),
	1901: uint16(sym_external_command),
	1902: uint16(sym_string),
	1903: uint16(17),
	1904: uint16(29),
	1905: uint16(1),
	1906: uint16(sym_comment),
	1907: uint16(57),
	1908: uint16(1),
	1909: uint16(sym_identifier),
	1910: uint16(59),
	1911: uint16(1),
	1912: uint16(anon_sym_SLASH),
	1913: uint16(61),
	1914: uint16(1),
	1915: uint16(anon_sym_if),
	1916: uint16(63),
	1917: uint16(1),
	1918: uint16(anon_sym_LPAREN),
	1919: uint16(67),
	1920: uint16(1),
	1921: uint16(aux_sym_string_token1),
	1922: uint16(69),
	1923: uint16(1),
	1924: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	1925: uint16(71),
	1926: uint16(1),
	1927: uint16(anon_sym_DQUOTE),
	1928: uint16(73),
	1929: uint16(1),
	1930: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1931: uint16(75),
	1932: uint16(1),
	1933: uint16(anon_sym_BQUOTE),
	1934: uint16(77),
	1935: uint16(1),
	1936: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	1937: uint16(79),
	1938: uint16(1),
	1939: uint16(sym_numeric_error),
	1940: uint16(94),
	1941: uint16(1),
	1942: uint16(sym__expression_inner),
	1943: uint16(361),
	1944: uint16(1),
	1945: uint16(sym_expression),
	1946: uint16(7),
	1947: uint16(2),
	1948: uint16(sym__backticked),
	1949: uint16(sym__indented_backticked),
	1950: uint16(53),
	1951: uint16(2),
	1952: uint16(sym_if_expression),
	1953: uint16(sym_value),
	1954: uint16(11),
	1955: uint16(3),
	1956: uint16(sym_function_call),
	1957: uint16(sym_external_command),
	1958: uint16(sym_string),
	1959: uint16(17),
	1960: uint16(29),
	1961: uint16(1),
	1962: uint16(sym_comment),
	1963: uint16(57),
	1964: uint16(1),
	1965: uint16(sym_identifier),
	1966: uint16(59),
	1967: uint16(1),
	1968: uint16(anon_sym_SLASH),
	1969: uint16(61),
	1970: uint16(1),
	1971: uint16(anon_sym_if),
	1972: uint16(63),
	1973: uint16(1),
	1974: uint16(anon_sym_LPAREN),
	1975: uint16(67),
	1976: uint16(1),
	1977: uint16(aux_sym_string_token1),
	1978: uint16(69),
	1979: uint16(1),
	1980: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	1981: uint16(71),
	1982: uint16(1),
	1983: uint16(anon_sym_DQUOTE),
	1984: uint16(73),
	1985: uint16(1),
	1986: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1987: uint16(75),
	1988: uint16(1),
	1989: uint16(anon_sym_BQUOTE),
	1990: uint16(77),
	1991: uint16(1),
	1992: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	1993: uint16(79),
	1994: uint16(1),
	1995: uint16(sym_numeric_error),
	1996: uint16(94),
	1997: uint16(1),
	1998: uint16(sym__expression_inner),
	1999: uint16(365),
	2000: uint16(1),
	2001: uint16(sym_expression),
	2002: uint16(7),
	2003: uint16(2),
	2004: uint16(sym__backticked),
	2005: uint16(sym__indented_backticked),
	2006: uint16(53),
	2007: uint16(2),
	2008: uint16(sym_if_expression),
	2009: uint16(sym_value),
	2010: uint16(11),
	2011: uint16(3),
	2012: uint16(sym_function_call),
	2013: uint16(sym_external_command),
	2014: uint16(sym_string),
	2015: uint16(17),
	2016: uint16(29),
	2017: uint16(1),
	2018: uint16(sym_comment),
	2019: uint16(87),
	2020: uint16(1),
	2021: uint16(aux_sym_string_token1),
	2022: uint16(89),
	2023: uint16(1),
	2024: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	2025: uint16(91),
	2026: uint16(1),
	2027: uint16(anon_sym_DQUOTE),
	2028: uint16(93),
	2029: uint16(1),
	2030: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2031: uint16(225),
	2032: uint16(1),
	2033: uint16(sym_identifier),
	2034: uint16(227),
	2035: uint16(1),
	2036: uint16(anon_sym_SLASH),
	2037: uint16(229),
	2038: uint16(1),
	2039: uint16(anon_sym_if),
	2040: uint16(231),
	2041: uint16(1),
	2042: uint16(anon_sym_LPAREN),
	2043: uint16(233),
	2044: uint16(1),
	2045: uint16(anon_sym_BQUOTE),
	2046: uint16(235),
	2047: uint16(1),
	2048: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	2049: uint16(237),
	2050: uint16(1),
	2051: uint16(sym_numeric_error),
	2052: uint16(255),
	2053: uint16(1),
	2054: uint16(sym__expression_inner),
	2055: uint16(367),
	2056: uint16(1),
	2057: uint16(sym_expression),
	2058: uint16(256),
	2059: uint16(2),
	2060: uint16(sym_if_expression),
	2061: uint16(sym_value),
	2062: uint16(257),
	2063: uint16(2),
	2064: uint16(sym__backticked),
	2065: uint16(sym__indented_backticked),
	2066: uint16(254),
	2067: uint16(3),
	2068: uint16(sym_function_call),
	2069: uint16(sym_external_command),
	2070: uint16(sym_string),
	2071: uint16(17),
	2072: uint16(29),
	2073: uint16(1),
	2074: uint16(sym_comment),
	2075: uint16(57),
	2076: uint16(1),
	2077: uint16(sym_identifier),
	2078: uint16(59),
	2079: uint16(1),
	2080: uint16(anon_sym_SLASH),
	2081: uint16(61),
	2082: uint16(1),
	2083: uint16(anon_sym_if),
	2084: uint16(63),
	2085: uint16(1),
	2086: uint16(anon_sym_LPAREN),
	2087: uint16(67),
	2088: uint16(1),
	2089: uint16(aux_sym_string_token1),
	2090: uint16(69),
	2091: uint16(1),
	2092: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	2093: uint16(71),
	2094: uint16(1),
	2095: uint16(anon_sym_DQUOTE),
	2096: uint16(73),
	2097: uint16(1),
	2098: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2099: uint16(75),
	2100: uint16(1),
	2101: uint16(anon_sym_BQUOTE),
	2102: uint16(77),
	2103: uint16(1),
	2104: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	2105: uint16(79),
	2106: uint16(1),
	2107: uint16(sym_numeric_error),
	2108: uint16(94),
	2109: uint16(1),
	2110: uint16(sym__expression_inner),
	2111: uint16(382),
	2112: uint16(1),
	2113: uint16(sym_expression),
	2114: uint16(7),
	2115: uint16(2),
	2116: uint16(sym__backticked),
	2117: uint16(sym__indented_backticked),
	2118: uint16(53),
	2119: uint16(2),
	2120: uint16(sym_if_expression),
	2121: uint16(sym_value),
	2122: uint16(11),
	2123: uint16(3),
	2124: uint16(sym_function_call),
	2125: uint16(sym_external_command),
	2126: uint16(sym_string),
	2127: uint16(17),
	2128: uint16(29),
	2129: uint16(1),
	2130: uint16(sym_comment),
	2131: uint16(57),
	2132: uint16(1),
	2133: uint16(sym_identifier),
	2134: uint16(59),
	2135: uint16(1),
	2136: uint16(anon_sym_SLASH),
	2137: uint16(61),
	2138: uint16(1),
	2139: uint16(anon_sym_if),
	2140: uint16(63),
	2141: uint16(1),
	2142: uint16(anon_sym_LPAREN),
	2143: uint16(67),
	2144: uint16(1),
	2145: uint16(aux_sym_string_token1),
	2146: uint16(69),
	2147: uint16(1),
	2148: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	2149: uint16(71),
	2150: uint16(1),
	2151: uint16(anon_sym_DQUOTE),
	2152: uint16(73),
	2153: uint16(1),
	2154: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2155: uint16(75),
	2156: uint16(1),
	2157: uint16(anon_sym_BQUOTE),
	2158: uint16(77),
	2159: uint16(1),
	2160: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	2161: uint16(79),
	2162: uint16(1),
	2163: uint16(sym_numeric_error),
	2164: uint16(94),
	2165: uint16(1),
	2166: uint16(sym__expression_inner),
	2167: uint16(386),
	2168: uint16(1),
	2169: uint16(sym_expression),
	2170: uint16(7),
	2171: uint16(2),
	2172: uint16(sym__backticked),
	2173: uint16(sym__indented_backticked),
	2174: uint16(53),
	2175: uint16(2),
	2176: uint16(sym_if_expression),
	2177: uint16(sym_value),
	2178: uint16(11),
	2179: uint16(3),
	2180: uint16(sym_function_call),
	2181: uint16(sym_external_command),
	2182: uint16(sym_string),
	2183: uint16(3),
	2184: uint16(29),
	2185: uint16(1),
	2186: uint16(sym_comment),
	2187: uint16(247),
	2188: uint16(6),
	2189: uint16(anon_sym_if),
	2190: uint16(anon_sym_else),
	2191: uint16(aux_sym_string_token1),
	2192: uint16(anon_sym_DQUOTE),
	2193: uint16(anon_sym_BQUOTE),
	2194: uint16(sym_identifier),
	2195: uint16(249),
	2196: uint16(14),
	2197: uint16(anon_sym_COMMA),
	2198: uint16(anon_sym_SLASH),
	2199: uint16(anon_sym_PLUS),
	2200: uint16(anon_sym_LBRACE),
	2201: uint16(anon_sym_EQ_EQ),
	2202: uint16(anon_sym_BANG_EQ),
	2203: uint16(anon_sym_EQ_TILDE),
	2204: uint16(anon_sym_LPAREN),
	2205: uint16(anon_sym_RPAREN),
	2206: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	2207: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2208: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	2209: uint16(anon_sym_RBRACE_RBRACE),
	2210: uint16(sym_numeric_error),
	2211: uint16(3),
	2212: uint16(29),
	2213: uint16(1),
	2214: uint16(sym_comment),
	2215: uint16(251),
	2216: uint16(5),
	2217: uint16(anon_sym_if),
	2218: uint16(aux_sym_string_token1),
	2219: uint16(anon_sym_DQUOTE),
	2220: uint16(anon_sym_BQUOTE),
	2221: uint16(sym_identifier),
	2222: uint16(253),
	2223: uint16(14),
	2224: uint16(anon_sym_COMMA),
	2225: uint16(anon_sym_SLASH),
	2226: uint16(anon_sym_PLUS),
	2227: uint16(anon_sym_LBRACE),
	2228: uint16(anon_sym_EQ_EQ),
	2229: uint16(anon_sym_BANG_EQ),
	2230: uint16(anon_sym_EQ_TILDE),
	2231: uint16(anon_sym_LPAREN),
	2232: uint16(anon_sym_RPAREN),
	2233: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	2234: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2235: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	2236: uint16(anon_sym_RBRACE_RBRACE),
	2237: uint16(sym_numeric_error),
	2238: uint16(3),
	2239: uint16(29),
	2240: uint16(1),
	2241: uint16(sym_comment),
	2242: uint16(255),
	2243: uint16(5),
	2244: uint16(anon_sym_if),
	2245: uint16(aux_sym_string_token1),
	2246: uint16(anon_sym_DQUOTE),
	2247: uint16(anon_sym_BQUOTE),
	2248: uint16(sym_identifier),
	2249: uint16(257),
	2250: uint16(14),
	2251: uint16(anon_sym_COMMA),
	2252: uint16(anon_sym_SLASH),
	2253: uint16(anon_sym_PLUS),
	2254: uint16(anon_sym_LBRACE),
	2255: uint16(anon_sym_EQ_EQ),
	2256: uint16(anon_sym_BANG_EQ),
	2257: uint16(anon_sym_EQ_TILDE),
	2258: uint16(anon_sym_LPAREN),
	2259: uint16(anon_sym_RPAREN),
	2260: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	2261: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2262: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	2263: uint16(anon_sym_RBRACE_RBRACE),
	2264: uint16(sym_numeric_error),
	2265: uint16(3),
	2266: uint16(29),
	2267: uint16(1),
	2268: uint16(sym_comment),
	2269: uint16(259),
	2270: uint16(5),
	2271: uint16(anon_sym_if),
	2272: uint16(aux_sym_string_token1),
	2273: uint16(anon_sym_DQUOTE),
	2274: uint16(anon_sym_BQUOTE),
	2275: uint16(sym_identifier),
	2276: uint16(261),
	2277: uint16(14),
	2278: uint16(anon_sym_COMMA),
	2279: uint16(anon_sym_SLASH),
	2280: uint16(anon_sym_PLUS),
	2281: uint16(anon_sym_LBRACE),
	2282: uint16(anon_sym_EQ_EQ),
	2283: uint16(anon_sym_BANG_EQ),
	2284: uint16(anon_sym_EQ_TILDE),
	2285: uint16(anon_sym_LPAREN),
	2286: uint16(anon_sym_RPAREN),
	2287: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	2288: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2289: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	2290: uint16(anon_sym_RBRACE_RBRACE),
	2291: uint16(sym_numeric_error),
	2292: uint16(3),
	2293: uint16(29),
	2294: uint16(1),
	2295: uint16(sym_comment),
	2296: uint16(263),
	2297: uint16(5),
	2298: uint16(anon_sym_if),
	2299: uint16(aux_sym_string_token1),
	2300: uint16(anon_sym_DQUOTE),
	2301: uint16(anon_sym_BQUOTE),
	2302: uint16(sym_identifier),
	2303: uint16(265),
	2304: uint16(14),
	2305: uint16(anon_sym_COMMA),
	2306: uint16(anon_sym_SLASH),
	2307: uint16(anon_sym_PLUS),
	2308: uint16(anon_sym_LBRACE),
	2309: uint16(anon_sym_EQ_EQ),
	2310: uint16(anon_sym_BANG_EQ),
	2311: uint16(anon_sym_EQ_TILDE),
	2312: uint16(anon_sym_LPAREN),
	2313: uint16(anon_sym_RPAREN),
	2314: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	2315: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2316: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	2317: uint16(anon_sym_RBRACE_RBRACE),
	2318: uint16(sym_numeric_error),
	2319: uint16(3),
	2320: uint16(29),
	2321: uint16(1),
	2322: uint16(sym_comment),
	2323: uint16(267),
	2324: uint16(5),
	2325: uint16(anon_sym_if),
	2326: uint16(aux_sym_string_token1),
	2327: uint16(anon_sym_DQUOTE),
	2328: uint16(anon_sym_BQUOTE),
	2329: uint16(sym_identifier),
	2330: uint16(269),
	2331: uint16(14),
	2332: uint16(anon_sym_COMMA),
	2333: uint16(anon_sym_SLASH),
	2334: uint16(anon_sym_PLUS),
	2335: uint16(anon_sym_LBRACE),
	2336: uint16(anon_sym_EQ_EQ),
	2337: uint16(anon_sym_BANG_EQ),
	2338: uint16(anon_sym_EQ_TILDE),
	2339: uint16(anon_sym_LPAREN),
	2340: uint16(anon_sym_RPAREN),
	2341: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	2342: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2343: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	2344: uint16(anon_sym_RBRACE_RBRACE),
	2345: uint16(sym_numeric_error),
	2346: uint16(15),
	2347: uint16(29),
	2348: uint16(1),
	2349: uint16(sym_comment),
	2350: uint16(61),
	2351: uint16(1),
	2352: uint16(anon_sym_if),
	2353: uint16(63),
	2354: uint16(1),
	2355: uint16(anon_sym_LPAREN),
	2356: uint16(75),
	2357: uint16(1),
	2358: uint16(anon_sym_BQUOTE),
	2359: uint16(77),
	2360: uint16(1),
	2361: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	2362: uint16(79),
	2363: uint16(1),
	2364: uint16(sym_numeric_error),
	2365: uint16(81),
	2366: uint16(1),
	2367: uint16(sym_identifier),
	2368: uint16(87),
	2369: uint16(1),
	2370: uint16(aux_sym_string_token1),
	2371: uint16(89),
	2372: uint16(1),
	2373: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	2374: uint16(91),
	2375: uint16(1),
	2376: uint16(anon_sym_DQUOTE),
	2377: uint16(93),
	2378: uint16(1),
	2379: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2380: uint16(82),
	2381: uint16(1),
	2382: uint16(sym__expression_inner),
	2383: uint16(7),
	2384: uint16(2),
	2385: uint16(sym__backticked),
	2386: uint16(sym__indented_backticked),
	2387: uint16(53),
	2388: uint16(2),
	2389: uint16(sym_if_expression),
	2390: uint16(sym_value),
	2391: uint16(11),
	2392: uint16(3),
	2393: uint16(sym_function_call),
	2394: uint16(sym_external_command),
	2395: uint16(sym_string),
	2396: uint16(15),
	2397: uint16(29),
	2398: uint16(1),
	2399: uint16(sym_comment),
	2400: uint16(203),
	2401: uint16(1),
	2402: uint16(sym_identifier),
	2403: uint16(207),
	2404: uint16(1),
	2405: uint16(anon_sym_if),
	2406: uint16(209),
	2407: uint16(1),
	2408: uint16(anon_sym_LPAREN),
	2409: uint16(211),
	2410: uint16(1),
	2411: uint16(aux_sym_string_token1),
	2412: uint16(213),
	2413: uint16(1),
	2414: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	2415: uint16(215),
	2416: uint16(1),
	2417: uint16(anon_sym_DQUOTE),
	2418: uint16(217),
	2419: uint16(1),
	2420: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2421: uint16(219),
	2422: uint16(1),
	2423: uint16(anon_sym_BQUOTE),
	2424: uint16(221),
	2425: uint16(1),
	2426: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	2427: uint16(223),
	2428: uint16(1),
	2429: uint16(sym_numeric_error),
	2430: uint16(233),
	2431: uint16(1),
	2432: uint16(sym__expression_inner),
	2433: uint16(239),
	2434: uint16(2),
	2435: uint16(sym_if_expression),
	2436: uint16(sym_value),
	2437: uint16(240),
	2438: uint16(2),
	2439: uint16(sym__backticked),
	2440: uint16(sym__indented_backticked),
	2441: uint16(234),
	2442: uint16(3),
	2443: uint16(sym_function_call),
	2444: uint16(sym_external_command),
	2445: uint16(sym_string),
	2446: uint16(15),
	2447: uint16(29),
	2448: uint16(1),
	2449: uint16(sym_comment),
	2450: uint16(87),
	2451: uint16(1),
	2452: uint16(aux_sym_string_token1),
	2453: uint16(89),
	2454: uint16(1),
	2455: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	2456: uint16(91),
	2457: uint16(1),
	2458: uint16(anon_sym_DQUOTE),
	2459: uint16(93),
	2460: uint16(1),
	2461: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2462: uint16(225),
	2463: uint16(1),
	2464: uint16(sym_identifier),
	2465: uint16(229),
	2466: uint16(1),
	2467: uint16(anon_sym_if),
	2468: uint16(231),
	2469: uint16(1),
	2470: uint16(anon_sym_LPAREN),
	2471: uint16(233),
	2472: uint16(1),
	2473: uint16(anon_sym_BQUOTE),
	2474: uint16(235),
	2475: uint16(1),
	2476: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	2477: uint16(237),
	2478: uint16(1),
	2479: uint16(sym_numeric_error),
	2480: uint16(276),
	2481: uint16(1),
	2482: uint16(sym__expression_inner),
	2483: uint16(256),
	2484: uint16(2),
	2485: uint16(sym_if_expression),
	2486: uint16(sym_value),
	2487: uint16(257),
	2488: uint16(2),
	2489: uint16(sym__backticked),
	2490: uint16(sym__indented_backticked),
	2491: uint16(254),
	2492: uint16(3),
	2493: uint16(sym_function_call),
	2494: uint16(sym_external_command),
	2495: uint16(sym_string),
	2496: uint16(15),
	2497: uint16(29),
	2498: uint16(1),
	2499: uint16(sym_comment),
	2500: uint16(61),
	2501: uint16(1),
	2502: uint16(anon_sym_if),
	2503: uint16(63),
	2504: uint16(1),
	2505: uint16(anon_sym_LPAREN),
	2506: uint16(75),
	2507: uint16(1),
	2508: uint16(anon_sym_BQUOTE),
	2509: uint16(77),
	2510: uint16(1),
	2511: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	2512: uint16(79),
	2513: uint16(1),
	2514: uint16(sym_numeric_error),
	2515: uint16(81),
	2516: uint16(1),
	2517: uint16(sym_identifier),
	2518: uint16(87),
	2519: uint16(1),
	2520: uint16(aux_sym_string_token1),
	2521: uint16(89),
	2522: uint16(1),
	2523: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	2524: uint16(91),
	2525: uint16(1),
	2526: uint16(anon_sym_DQUOTE),
	2527: uint16(93),
	2528: uint16(1),
	2529: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2530: uint16(52),
	2531: uint16(1),
	2532: uint16(sym__expression_inner),
	2533: uint16(7),
	2534: uint16(2),
	2535: uint16(sym__backticked),
	2536: uint16(sym__indented_backticked),
	2537: uint16(53),
	2538: uint16(2),
	2539: uint16(sym_if_expression),
	2540: uint16(sym_value),
	2541: uint16(11),
	2542: uint16(3),
	2543: uint16(sym_function_call),
	2544: uint16(sym_external_command),
	2545: uint16(sym_string),
	2546: uint16(15),
	2547: uint16(29),
	2548: uint16(1),
	2549: uint16(sym_comment),
	2550: uint16(87),
	2551: uint16(1),
	2552: uint16(aux_sym_string_token1),
	2553: uint16(89),
	2554: uint16(1),
	2555: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	2556: uint16(91),
	2557: uint16(1),
	2558: uint16(anon_sym_DQUOTE),
	2559: uint16(93),
	2560: uint16(1),
	2561: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2562: uint16(225),
	2563: uint16(1),
	2564: uint16(sym_identifier),
	2565: uint16(229),
	2566: uint16(1),
	2567: uint16(anon_sym_if),
	2568: uint16(231),
	2569: uint16(1),
	2570: uint16(anon_sym_LPAREN),
	2571: uint16(233),
	2572: uint16(1),
	2573: uint16(anon_sym_BQUOTE),
	2574: uint16(235),
	2575: uint16(1),
	2576: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	2577: uint16(237),
	2578: uint16(1),
	2579: uint16(sym_numeric_error),
	2580: uint16(266),
	2581: uint16(1),
	2582: uint16(sym__expression_inner),
	2583: uint16(256),
	2584: uint16(2),
	2585: uint16(sym_if_expression),
	2586: uint16(sym_value),
	2587: uint16(257),
	2588: uint16(2),
	2589: uint16(sym__backticked),
	2590: uint16(sym__indented_backticked),
	2591: uint16(254),
	2592: uint16(3),
	2593: uint16(sym_function_call),
	2594: uint16(sym_external_command),
	2595: uint16(sym_string),
	2596: uint16(15),
	2597: uint16(29),
	2598: uint16(1),
	2599: uint16(sym_comment),
	2600: uint16(87),
	2601: uint16(1),
	2602: uint16(aux_sym_string_token1),
	2603: uint16(89),
	2604: uint16(1),
	2605: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	2606: uint16(91),
	2607: uint16(1),
	2608: uint16(anon_sym_DQUOTE),
	2609: uint16(93),
	2610: uint16(1),
	2611: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2612: uint16(225),
	2613: uint16(1),
	2614: uint16(sym_identifier),
	2615: uint16(229),
	2616: uint16(1),
	2617: uint16(anon_sym_if),
	2618: uint16(231),
	2619: uint16(1),
	2620: uint16(anon_sym_LPAREN),
	2621: uint16(233),
	2622: uint16(1),
	2623: uint16(anon_sym_BQUOTE),
	2624: uint16(235),
	2625: uint16(1),
	2626: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	2627: uint16(237),
	2628: uint16(1),
	2629: uint16(sym_numeric_error),
	2630: uint16(258),
	2631: uint16(1),
	2632: uint16(sym__expression_inner),
	2633: uint16(256),
	2634: uint16(2),
	2635: uint16(sym_if_expression),
	2636: uint16(sym_value),
	2637: uint16(257),
	2638: uint16(2),
	2639: uint16(sym__backticked),
	2640: uint16(sym__indented_backticked),
	2641: uint16(254),
	2642: uint16(3),
	2643: uint16(sym_function_call),
	2644: uint16(sym_external_command),
	2645: uint16(sym_string),
	2646: uint16(15),
	2647: uint16(29),
	2648: uint16(1),
	2649: uint16(sym_comment),
	2650: uint16(203),
	2651: uint16(1),
	2652: uint16(sym_identifier),
	2653: uint16(207),
	2654: uint16(1),
	2655: uint16(anon_sym_if),
	2656: uint16(209),
	2657: uint16(1),
	2658: uint16(anon_sym_LPAREN),
	2659: uint16(211),
	2660: uint16(1),
	2661: uint16(aux_sym_string_token1),
	2662: uint16(213),
	2663: uint16(1),
	2664: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	2665: uint16(215),
	2666: uint16(1),
	2667: uint16(anon_sym_DQUOTE),
	2668: uint16(217),
	2669: uint16(1),
	2670: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2671: uint16(219),
	2672: uint16(1),
	2673: uint16(anon_sym_BQUOTE),
	2674: uint16(221),
	2675: uint16(1),
	2676: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	2677: uint16(223),
	2678: uint16(1),
	2679: uint16(sym_numeric_error),
	2680: uint16(261),
	2681: uint16(1),
	2682: uint16(sym__expression_inner),
	2683: uint16(239),
	2684: uint16(2),
	2685: uint16(sym_if_expression),
	2686: uint16(sym_value),
	2687: uint16(240),
	2688: uint16(2),
	2689: uint16(sym__backticked),
	2690: uint16(sym__indented_backticked),
	2691: uint16(234),
	2692: uint16(3),
	2693: uint16(sym_function_call),
	2694: uint16(sym_external_command),
	2695: uint16(sym_string),
	2696: uint16(15),
	2697: uint16(29),
	2698: uint16(1),
	2699: uint16(sym_comment),
	2700: uint16(57),
	2701: uint16(1),
	2702: uint16(sym_identifier),
	2703: uint16(61),
	2704: uint16(1),
	2705: uint16(anon_sym_if),
	2706: uint16(63),
	2707: uint16(1),
	2708: uint16(anon_sym_LPAREN),
	2709: uint16(67),
	2710: uint16(1),
	2711: uint16(aux_sym_string_token1),
	2712: uint16(69),
	2713: uint16(1),
	2714: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	2715: uint16(71),
	2716: uint16(1),
	2717: uint16(anon_sym_DQUOTE),
	2718: uint16(73),
	2719: uint16(1),
	2720: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2721: uint16(75),
	2722: uint16(1),
	2723: uint16(anon_sym_BQUOTE),
	2724: uint16(77),
	2725: uint16(1),
	2726: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	2727: uint16(79),
	2728: uint16(1),
	2729: uint16(sym_numeric_error),
	2730: uint16(52),
	2731: uint16(1),
	2732: uint16(sym__expression_inner),
	2733: uint16(7),
	2734: uint16(2),
	2735: uint16(sym__backticked),
	2736: uint16(sym__indented_backticked),
	2737: uint16(53),
	2738: uint16(2),
	2739: uint16(sym_if_expression),
	2740: uint16(sym_value),
	2741: uint16(11),
	2742: uint16(3),
	2743: uint16(sym_function_call),
	2744: uint16(sym_external_command),
	2745: uint16(sym_string),
	2746: uint16(15),
	2747: uint16(29),
	2748: uint16(1),
	2749: uint16(sym_comment),
	2750: uint16(57),
	2751: uint16(1),
	2752: uint16(sym_identifier),
	2753: uint16(61),
	2754: uint16(1),
	2755: uint16(anon_sym_if),
	2756: uint16(63),
	2757: uint16(1),
	2758: uint16(anon_sym_LPAREN),
	2759: uint16(67),
	2760: uint16(1),
	2761: uint16(aux_sym_string_token1),
	2762: uint16(69),
	2763: uint16(1),
	2764: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	2765: uint16(71),
	2766: uint16(1),
	2767: uint16(anon_sym_DQUOTE),
	2768: uint16(73),
	2769: uint16(1),
	2770: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2771: uint16(75),
	2772: uint16(1),
	2773: uint16(anon_sym_BQUOTE),
	2774: uint16(77),
	2775: uint16(1),
	2776: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	2777: uint16(79),
	2778: uint16(1),
	2779: uint16(sym_numeric_error),
	2780: uint16(96),
	2781: uint16(1),
	2782: uint16(sym__expression_inner),
	2783: uint16(7),
	2784: uint16(2),
	2785: uint16(sym__backticked),
	2786: uint16(sym__indented_backticked),
	2787: uint16(53),
	2788: uint16(2),
	2789: uint16(sym_if_expression),
	2790: uint16(sym_value),
	2791: uint16(11),
	2792: uint16(3),
	2793: uint16(sym_function_call),
	2794: uint16(sym_external_command),
	2795: uint16(sym_string),
	2796: uint16(15),
	2797: uint16(29),
	2798: uint16(1),
	2799: uint16(sym_comment),
	2800: uint16(61),
	2801: uint16(1),
	2802: uint16(anon_sym_if),
	2803: uint16(63),
	2804: uint16(1),
	2805: uint16(anon_sym_LPAREN),
	2806: uint16(75),
	2807: uint16(1),
	2808: uint16(anon_sym_BQUOTE),
	2809: uint16(77),
	2810: uint16(1),
	2811: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	2812: uint16(79),
	2813: uint16(1),
	2814: uint16(sym_numeric_error),
	2815: uint16(81),
	2816: uint16(1),
	2817: uint16(sym_identifier),
	2818: uint16(87),
	2819: uint16(1),
	2820: uint16(aux_sym_string_token1),
	2821: uint16(89),
	2822: uint16(1),
	2823: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	2824: uint16(91),
	2825: uint16(1),
	2826: uint16(anon_sym_DQUOTE),
	2827: uint16(93),
	2828: uint16(1),
	2829: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2830: uint16(78),
	2831: uint16(1),
	2832: uint16(sym__expression_inner),
	2833: uint16(7),
	2834: uint16(2),
	2835: uint16(sym__backticked),
	2836: uint16(sym__indented_backticked),
	2837: uint16(53),
	2838: uint16(2),
	2839: uint16(sym_if_expression),
	2840: uint16(sym_value),
	2841: uint16(11),
	2842: uint16(3),
	2843: uint16(sym_function_call),
	2844: uint16(sym_external_command),
	2845: uint16(sym_string),
	2846: uint16(15),
	2847: uint16(29),
	2848: uint16(1),
	2849: uint16(sym_comment),
	2850: uint16(57),
	2851: uint16(1),
	2852: uint16(sym_identifier),
	2853: uint16(61),
	2854: uint16(1),
	2855: uint16(anon_sym_if),
	2856: uint16(63),
	2857: uint16(1),
	2858: uint16(anon_sym_LPAREN),
	2859: uint16(67),
	2860: uint16(1),
	2861: uint16(aux_sym_string_token1),
	2862: uint16(69),
	2863: uint16(1),
	2864: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	2865: uint16(71),
	2866: uint16(1),
	2867: uint16(anon_sym_DQUOTE),
	2868: uint16(73),
	2869: uint16(1),
	2870: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2871: uint16(75),
	2872: uint16(1),
	2873: uint16(anon_sym_BQUOTE),
	2874: uint16(77),
	2875: uint16(1),
	2876: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	2877: uint16(79),
	2878: uint16(1),
	2879: uint16(sym_numeric_error),
	2880: uint16(102),
	2881: uint16(1),
	2882: uint16(sym__expression_inner),
	2883: uint16(7),
	2884: uint16(2),
	2885: uint16(sym__backticked),
	2886: uint16(sym__indented_backticked),
	2887: uint16(53),
	2888: uint16(2),
	2889: uint16(sym_if_expression),
	2890: uint16(sym_value),
	2891: uint16(11),
	2892: uint16(3),
	2893: uint16(sym_function_call),
	2894: uint16(sym_external_command),
	2895: uint16(sym_string),
	2896: uint16(15),
	2897: uint16(29),
	2898: uint16(1),
	2899: uint16(sym_comment),
	2900: uint16(203),
	2901: uint16(1),
	2902: uint16(sym_identifier),
	2903: uint16(207),
	2904: uint16(1),
	2905: uint16(anon_sym_if),
	2906: uint16(209),
	2907: uint16(1),
	2908: uint16(anon_sym_LPAREN),
	2909: uint16(211),
	2910: uint16(1),
	2911: uint16(aux_sym_string_token1),
	2912: uint16(213),
	2913: uint16(1),
	2914: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	2915: uint16(215),
	2916: uint16(1),
	2917: uint16(anon_sym_DQUOTE),
	2918: uint16(217),
	2919: uint16(1),
	2920: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2921: uint16(219),
	2922: uint16(1),
	2923: uint16(anon_sym_BQUOTE),
	2924: uint16(221),
	2925: uint16(1),
	2926: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	2927: uint16(223),
	2928: uint16(1),
	2929: uint16(sym_numeric_error),
	2930: uint16(230),
	2931: uint16(1),
	2932: uint16(sym__expression_inner),
	2933: uint16(239),
	2934: uint16(2),
	2935: uint16(sym_if_expression),
	2936: uint16(sym_value),
	2937: uint16(240),
	2938: uint16(2),
	2939: uint16(sym__backticked),
	2940: uint16(sym__indented_backticked),
	2941: uint16(234),
	2942: uint16(3),
	2943: uint16(sym_function_call),
	2944: uint16(sym_external_command),
	2945: uint16(sym_string),
	2946: uint16(3),
	2947: uint16(29),
	2948: uint16(1),
	2949: uint16(sym_comment),
	2950: uint16(33),
	2951: uint16(5),
	2952: uint16(anon_sym_if),
	2953: uint16(aux_sym_string_token1),
	2954: uint16(anon_sym_DQUOTE),
	2955: uint16(anon_sym_BQUOTE),
	2956: uint16(sym_identifier),
	2957: uint16(31),
	2958: uint16(12),
	2959: uint16(anon_sym_SLASH),
	2960: uint16(anon_sym_PLUS),
	2961: uint16(anon_sym_RBRACE),
	2962: uint16(anon_sym_LPAREN),
	2963: uint16(anon_sym_RPAREN),
	2964: uint16(anon_sym_COLON),
	2965: uint16(anon_sym_DOLLAR),
	2966: uint16(anon_sym_STAR),
	2967: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	2968: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2969: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	2970: uint16(sym_numeric_error),
	2971: uint16(3),
	2972: uint16(29),
	2973: uint16(1),
	2974: uint16(sym_comment),
	2975: uint16(27),
	2976: uint16(5),
	2977: uint16(anon_sym_if),
	2978: uint16(aux_sym_string_token1),
	2979: uint16(anon_sym_DQUOTE),
	2980: uint16(anon_sym_BQUOTE),
	2981: uint16(sym_identifier),
	2982: uint16(25),
	2983: uint16(12),
	2984: uint16(anon_sym_SLASH),
	2985: uint16(anon_sym_PLUS),
	2986: uint16(anon_sym_RBRACE),
	2987: uint16(anon_sym_LPAREN),
	2988: uint16(anon_sym_RPAREN),
	2989: uint16(anon_sym_COLON),
	2990: uint16(anon_sym_DOLLAR),
	2991: uint16(anon_sym_STAR),
	2992: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	2993: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2994: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	2995: uint16(sym_numeric_error),
	2996: uint16(3),
	2997: uint16(29),
	2998: uint16(1),
	2999: uint16(sym_comment),
	3000: uint16(37),
	3001: uint16(5),
	3002: uint16(anon_sym_if),
	3003: uint16(aux_sym_string_token1),
	3004: uint16(anon_sym_DQUOTE),
	3005: uint16(anon_sym_BQUOTE),
	3006: uint16(sym_identifier),
	3007: uint16(35),
	3008: uint16(12),
	3009: uint16(anon_sym_SLASH),
	3010: uint16(anon_sym_PLUS),
	3011: uint16(anon_sym_RBRACE),
	3012: uint16(anon_sym_LPAREN),
	3013: uint16(anon_sym_RPAREN),
	3014: uint16(anon_sym_COLON),
	3015: uint16(anon_sym_DOLLAR),
	3016: uint16(anon_sym_STAR),
	3017: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	3018: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	3019: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	3020: uint16(sym_numeric_error),
	3021: uint16(13),
	3022: uint16(3),
	3023: uint16(1),
	3024: uint16(sym_comment),
	3025: uint16(23),
	3026: uint16(1),
	3027: uint16(aux_sym_shebang_token1),
	3028: uint16(275),
	3029: uint16(1),
	3030: uint16(anon_sym_LBRACE_LBRACE),
	3031: uint16(277),
	3032: uint16(1),
	3033: uint16(sym__dedent),
	3034: uint16(279),
	3035: uint16(1),
	3036: uint16(sym__newline),
	3037: uint16(281),
	3038: uint16(1),
	3039: uint16(sym_text),
	3040: uint16(84),
	3041: uint16(1),
	3042: uint16(aux_sym_recipe_body_repeat1),
	3043: uint16(218),
	3044: uint16(1),
	3045: uint16(sym_recipe_line_prefix),
	3046: uint16(385),
	3047: uint16(1),
	3048: uint16(sym_recipe_line),
	3049: uint16(393),
	3050: uint16(1),
	3051: uint16(sym_shebang),
	3052: uint16(271),
	3053: uint16(2),
	3054: uint16(anon_sym_AT),
	3055: uint16(anon_sym_DASH),
	3056: uint16(273),
	3057: uint16(2),
	3058: uint16(anon_sym_AT_DASH),
	3059: uint16(anon_sym_DASH_AT),
	3060: uint16(187),
	3061: uint16(2),
	3062: uint16(sym_interpolation),
	3063: uint16(aux_sym_recipe_line_repeat1),
	3064: uint16(13),
	3065: uint16(29),
	3066: uint16(1),
	3067: uint16(sym_comment),
	3068: uint16(63),
	3069: uint16(1),
	3070: uint16(anon_sym_LPAREN),
	3071: uint16(67),
	3072: uint16(1),
	3073: uint16(aux_sym_string_token1),
	3074: uint16(69),
	3075: uint16(1),
	3076: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	3077: uint16(71),
	3078: uint16(1),
	3079: uint16(anon_sym_DQUOTE),
	3080: uint16(73),
	3081: uint16(1),
	3082: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	3083: uint16(75),
	3084: uint16(1),
	3085: uint16(anon_sym_BQUOTE),
	3086: uint16(77),
	3087: uint16(1),
	3088: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	3089: uint16(79),
	3090: uint16(1),
	3091: uint16(sym_numeric_error),
	3092: uint16(283),
	3093: uint16(1),
	3094: uint16(sym_identifier),
	3095: uint16(164),
	3096: uint16(1),
	3097: uint16(sym_value),
	3098: uint16(7),
	3099: uint16(2),
	3100: uint16(sym__backticked),
	3101: uint16(sym__indented_backticked),
	3102: uint16(11),
	3103: uint16(3),
	3104: uint16(sym_function_call),
	3105: uint16(sym_external_command),
	3106: uint16(sym_string),
	3107: uint16(13),
	3108: uint16(29),
	3109: uint16(1),
	3110: uint16(sym_comment),
	3111: uint16(63),
	3112: uint16(1),
	3113: uint16(anon_sym_LPAREN),
	3114: uint16(75),
	3115: uint16(1),
	3116: uint16(anon_sym_BQUOTE),
	3117: uint16(77),
	3118: uint16(1),
	3119: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	3120: uint16(79),
	3121: uint16(1),
	3122: uint16(sym_numeric_error),
	3123: uint16(87),
	3124: uint16(1),
	3125: uint16(aux_sym_string_token1),
	3126: uint16(89),
	3127: uint16(1),
	3128: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	3129: uint16(91),
	3130: uint16(1),
	3131: uint16(anon_sym_DQUOTE),
	3132: uint16(93),
	3133: uint16(1),
	3134: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	3135: uint16(283),
	3136: uint16(1),
	3137: uint16(sym_identifier),
	3138: uint16(183),
	3139: uint16(1),
	3140: uint16(sym_value),
	3141: uint16(7),
	3142: uint16(2),
	3143: uint16(sym__backticked),
	3144: uint16(sym__indented_backticked),
	3145: uint16(11),
	3146: uint16(3),
	3147: uint16(sym_function_call),
	3148: uint16(sym_external_command),
	3149: uint16(sym_string),
	3150: uint16(13),
	3151: uint16(29),
	3152: uint16(1),
	3153: uint16(sym_comment),
	3154: uint16(63),
	3155: uint16(1),
	3156: uint16(anon_sym_LPAREN),
	3157: uint16(75),
	3158: uint16(1),
	3159: uint16(anon_sym_BQUOTE),
	3160: uint16(77),
	3161: uint16(1),
	3162: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	3163: uint16(79),
	3164: uint16(1),
	3165: uint16(sym_numeric_error),
	3166: uint16(87),
	3167: uint16(1),
	3168: uint16(aux_sym_string_token1),
	3169: uint16(89),
	3170: uint16(1),
	3171: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	3172: uint16(91),
	3173: uint16(1),
	3174: uint16(anon_sym_DQUOTE),
	3175: uint16(93),
	3176: uint16(1),
	3177: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	3178: uint16(283),
	3179: uint16(1),
	3180: uint16(sym_identifier),
	3181: uint16(164),
	3182: uint16(1),
	3183: uint16(sym_value),
	3184: uint16(7),
	3185: uint16(2),
	3186: uint16(sym__backticked),
	3187: uint16(sym__indented_backticked),
	3188: uint16(11),
	3189: uint16(3),
	3190: uint16(sym_function_call),
	3191: uint16(sym_external_command),
	3192: uint16(sym_string),
	3193: uint16(13),
	3194: uint16(29),
	3195: uint16(1),
	3196: uint16(sym_comment),
	3197: uint16(63),
	3198: uint16(1),
	3199: uint16(anon_sym_LPAREN),
	3200: uint16(67),
	3201: uint16(1),
	3202: uint16(aux_sym_string_token1),
	3203: uint16(69),
	3204: uint16(1),
	3205: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	3206: uint16(71),
	3207: uint16(1),
	3208: uint16(anon_sym_DQUOTE),
	3209: uint16(73),
	3210: uint16(1),
	3211: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	3212: uint16(75),
	3213: uint16(1),
	3214: uint16(anon_sym_BQUOTE),
	3215: uint16(77),
	3216: uint16(1),
	3217: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	3218: uint16(79),
	3219: uint16(1),
	3220: uint16(sym_numeric_error),
	3221: uint16(283),
	3222: uint16(1),
	3223: uint16(sym_identifier),
	3224: uint16(183),
	3225: uint16(1),
	3226: uint16(sym_value),
	3227: uint16(7),
	3228: uint16(2),
	3229: uint16(sym__backticked),
	3230: uint16(sym__indented_backticked),
	3231: uint16(11),
	3232: uint16(3),
	3233: uint16(sym_function_call),
	3234: uint16(sym_external_command),
	3235: uint16(sym_string),
	3236: uint16(15),
	3237: uint16(29),
	3238: uint16(1),
	3239: uint16(sym_comment),
	3240: uint16(67),
	3241: uint16(1),
	3242: uint16(aux_sym_string_token1),
	3243: uint16(69),
	3244: uint16(1),
	3245: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	3246: uint16(71),
	3247: uint16(1),
	3248: uint16(anon_sym_DQUOTE),
	3249: uint16(73),
	3250: uint16(1),
	3251: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	3252: uint16(285),
	3253: uint16(1),
	3254: uint16(sym_identifier),
	3255: uint16(287),
	3256: uint16(1),
	3257: uint16(anon_sym_QMARK),
	3258: uint16(291),
	3259: uint16(1),
	3260: uint16(anon_sym_COLON),
	3261: uint16(293),
	3262: uint16(1),
	3263: uint16(anon_sym_DOLLAR),
	3264: uint16(117),
	3265: uint16(1),
	3266: uint16(sym_string),
	3267: uint16(141),
	3268: uint16(1),
	3269: uint16(aux_sym_parameters_repeat1),
	3270: uint16(169),
	3271: uint16(1),
	3272: uint16(sym_parameter),
	3273: uint16(329),
	3274: uint16(1),
	3275: uint16(sym_parameters),
	3276: uint16(349),
	3277: uint16(1),
	3278: uint16(sym_variadic_parameter),
	3279: uint16(289),
	3280: uint16(2),
	3281: uint16(anon_sym_PLUS),
	3282: uint16(anon_sym_STAR),
	3283: uint16(8),
	3284: uint16(29),
	3285: uint16(1),
	3286: uint16(sym_comment),
	3287: uint16(67),
	3288: uint16(1),
	3289: uint16(aux_sym_string_token1),
	3290: uint16(69),
	3291: uint16(1),
	3292: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	3293: uint16(71),
	3294: uint16(1),
	3295: uint16(anon_sym_DQUOTE),
	3296: uint16(73),
	3297: uint16(1),
	3298: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	3299: uint16(122),
	3300: uint16(1),
	3301: uint16(sym_string),
	3302: uint16(295),
	3303: uint16(3),
	3305: uint16(anon_sym_LBRACK),
	3306: uint16(anon_sym_AT),
	3307: uint16(297),
	3308: uint16(6),
	3309: uint16(anon_sym_alias),
	3310: uint16(anon_sym_export),
	3311: uint16(anon_sym_import),
	3312: uint16(anon_sym_mod),
	3313: uint16(anon_sym_set),
	3314: uint16(sym_identifier),
	3315: uint16(8),
	3316: uint16(29),
	3317: uint16(1),
	3318: uint16(sym_comment),
	3319: uint16(67),
	3320: uint16(1),
	3321: uint16(aux_sym_string_token1),
	3322: uint16(69),
	3323: uint16(1),
	3324: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	3325: uint16(71),
	3326: uint16(1),
	3327: uint16(anon_sym_DQUOTE),
	3328: uint16(73),
	3329: uint16(1),
	3330: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	3331: uint16(109),
	3332: uint16(1),
	3333: uint16(sym_string),
	3334: uint16(299),
	3335: uint16(3),
	3337: uint16(anon_sym_LBRACK),
	3338: uint16(anon_sym_AT),
	3339: uint16(301),
	3340: uint16(6),
	3341: uint16(anon_sym_alias),
	3342: uint16(anon_sym_export),
	3343: uint16(anon_sym_import),
	3344: uint16(anon_sym_mod),
	3345: uint16(anon_sym_set),
	3346: uint16(sym_identifier),
	3347: uint16(3),
	3348: uint16(29),
	3349: uint16(1),
	3350: uint16(sym_comment),
	3351: uint16(303),
	3352: uint16(1),
	3353: uint16(anon_sym_LPAREN),
	3354: uint16(97),
	3355: uint16(13),
	3356: uint16(anon_sym_COMMA),
	3357: uint16(anon_sym_SLASH),
	3358: uint16(anon_sym_PLUS),
	3359: uint16(anon_sym_LBRACE),
	3360: uint16(anon_sym_EQ_EQ),
	3361: uint16(anon_sym_BANG_EQ),
	3362: uint16(anon_sym_EQ_TILDE),
	3363: uint16(anon_sym_RPAREN),
	3364: uint16(anon_sym_COLON),
	3365: uint16(anon_sym_DOLLAR),
	3366: uint16(anon_sym_STAR),
	3367: uint16(anon_sym_RBRACE_RBRACE),
	3368: uint16(sym_identifier),
	3369: uint16(5),
	3370: uint16(29),
	3371: uint16(1),
	3372: uint16(sym_comment),
	3373: uint16(307),
	3374: uint16(1),
	3375: uint16(anon_sym_SLASH),
	3376: uint16(309),
	3377: uint16(1),
	3378: uint16(anon_sym_PLUS),
	3379: uint16(305),
	3380: uint16(5),
	3381: uint16(anon_sym_if),
	3382: uint16(aux_sym_string_token1),
	3383: uint16(anon_sym_DQUOTE),
	3384: uint16(anon_sym_BQUOTE),
	3385: uint16(sym_identifier),
	3386: uint16(311),
	3387: uint16(6),
	3388: uint16(anon_sym_LPAREN),
	3389: uint16(anon_sym_RPAREN),
	3390: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	3391: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	3392: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	3393: uint16(sym_numeric_error),
	3394: uint16(11),
	3395: uint16(29),
	3396: uint16(1),
	3397: uint16(sym_comment),
	3398: uint16(319),
	3399: uint16(1),
	3400: uint16(anon_sym_LBRACE_LBRACE),
	3401: uint16(322),
	3402: uint16(1),
	3403: uint16(sym__dedent),
	3404: uint16(324),
	3405: uint16(1),
	3406: uint16(sym__newline),
	3407: uint16(327),
	3408: uint16(1),
	3409: uint16(sym_text),
	3410: uint16(79),
	3411: uint16(1),
	3412: uint16(aux_sym_recipe_body_repeat1),
	3413: uint16(218),
	3414: uint16(1),
	3415: uint16(sym_recipe_line_prefix),
	3416: uint16(385),
	3417: uint16(1),
	3418: uint16(sym_recipe_line),
	3419: uint16(313),
	3420: uint16(2),
	3421: uint16(anon_sym_AT),
	3422: uint16(anon_sym_DASH),
	3423: uint16(316),
	3424: uint16(2),
	3425: uint16(anon_sym_AT_DASH),
	3426: uint16(anon_sym_DASH_AT),
	3427: uint16(187),
	3428: uint16(2),
	3429: uint16(sym_interpolation),
	3430: uint16(aux_sym_recipe_line_repeat1),
	3431: uint16(11),
	3432: uint16(29),
	3433: uint16(1),
	3434: uint16(sym_comment),
	3435: uint16(275),
	3436: uint16(1),
	3437: uint16(anon_sym_LBRACE_LBRACE),
	3438: uint16(281),
	3439: uint16(1),
	3440: uint16(sym_text),
	3441: uint16(330),
	3442: uint16(1),
	3443: uint16(sym__dedent),
	3444: uint16(332),
	3445: uint16(1),
	3446: uint16(sym__newline),
	3447: uint16(79),
	3448: uint16(1),
	3449: uint16(aux_sym_recipe_body_repeat1),
	3450: uint16(218),
	3451: uint16(1),
	3452: uint16(sym_recipe_line_prefix),
	3453: uint16(385),
	3454: uint16(1),
	3455: uint16(sym_recipe_line),
	3456: uint16(271),
	3457: uint16(2),
	3458: uint16(anon_sym_AT),
	3459: uint16(anon_sym_DASH),
	3460: uint16(273),
	3461: uint16(2),
	3462: uint16(anon_sym_AT_DASH),
	3463: uint16(anon_sym_DASH_AT),
	3464: uint16(187),
	3465: uint16(2),
	3466: uint16(sym_interpolation),
	3467: uint16(aux_sym_recipe_line_repeat1),
	3468: uint16(3),
	3469: uint16(29),
	3470: uint16(1),
	3471: uint16(sym_comment),
	3472: uint16(95),
	3473: uint16(5),
	3474: uint16(anon_sym_if),
	3475: uint16(aux_sym_string_token1),
	3476: uint16(anon_sym_DQUOTE),
	3477: uint16(anon_sym_BQUOTE),
	3478: uint16(sym_identifier),
	3479: uint16(97),
	3480: uint16(8),
	3481: uint16(anon_sym_SLASH),
	3482: uint16(anon_sym_PLUS),
	3483: uint16(anon_sym_LPAREN),
	3484: uint16(anon_sym_RPAREN),
	3485: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	3486: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	3487: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	3488: uint16(sym_numeric_error),
	3489: uint16(4),
	3490: uint16(29),
	3491: uint16(1),
	3492: uint16(sym_comment),
	3493: uint16(309),
	3494: uint16(1),
	3495: uint16(anon_sym_PLUS),
	3496: uint16(263),
	3497: uint16(5),
	3498: uint16(anon_sym_if),
	3499: uint16(aux_sym_string_token1),
	3500: uint16(anon_sym_DQUOTE),
	3501: uint16(anon_sym_BQUOTE),
	3502: uint16(sym_identifier),
	3503: uint16(265),
	3504: uint16(7),
	3505: uint16(anon_sym_SLASH),
	3506: uint16(anon_sym_LPAREN),
	3507: uint16(anon_sym_RPAREN),
	3508: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	3509: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	3510: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	3511: uint16(sym_numeric_error),
	3512: uint16(5),
	3513: uint16(29),
	3514: uint16(1),
	3515: uint16(sym_comment),
	3516: uint16(307),
	3517: uint16(1),
	3518: uint16(anon_sym_SLASH),
	3519: uint16(309),
	3520: uint16(1),
	3521: uint16(anon_sym_PLUS),
	3522: uint16(334),
	3523: uint16(5),
	3524: uint16(anon_sym_if),
	3525: uint16(aux_sym_string_token1),
	3526: uint16(anon_sym_DQUOTE),
	3527: uint16(anon_sym_BQUOTE),
	3528: uint16(sym_identifier),
	3529: uint16(336),
	3530: uint16(6),
	3531: uint16(anon_sym_LPAREN),
	3532: uint16(anon_sym_RPAREN),
	3533: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	3534: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	3535: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	3536: uint16(sym_numeric_error),
	3537: uint16(11),
	3538: uint16(29),
	3539: uint16(1),
	3540: uint16(sym_comment),
	3541: uint16(275),
	3542: uint16(1),
	3543: uint16(anon_sym_LBRACE_LBRACE),
	3544: uint16(281),
	3545: uint16(1),
	3546: uint16(sym_text),
	3547: uint16(332),
	3548: uint16(1),
	3549: uint16(sym__newline),
	3550: uint16(338),
	3551: uint16(1),
	3552: uint16(sym__dedent),
	3553: uint16(79),
	3554: uint16(1),
	3555: uint16(aux_sym_recipe_body_repeat1),
	3556: uint16(218),
	3557: uint16(1),
	3558: uint16(sym_recipe_line_prefix),
	3559: uint16(385),
	3560: uint16(1),
	3561: uint16(sym_recipe_line),
	3562: uint16(271),
	3563: uint16(2),
	3564: uint16(anon_sym_AT),
	3565: uint16(anon_sym_DASH),
	3566: uint16(273),
	3567: uint16(2),
	3568: uint16(anon_sym_AT_DASH),
	3569: uint16(anon_sym_DASH_AT),
	3570: uint16(187),
	3571: uint16(2),
	3572: uint16(sym_interpolation),
	3573: uint16(aux_sym_recipe_line_repeat1),
	3574: uint16(11),
	3575: uint16(29),
	3576: uint16(1),
	3577: uint16(sym_comment),
	3578: uint16(275),
	3579: uint16(1),
	3580: uint16(anon_sym_LBRACE_LBRACE),
	3581: uint16(281),
	3582: uint16(1),
	3583: uint16(sym_text),
	3584: uint16(340),
	3585: uint16(1),
	3586: uint16(sym__dedent),
	3587: uint16(342),
	3588: uint16(1),
	3589: uint16(sym__newline),
	3590: uint16(80),
	3591: uint16(1),
	3592: uint16(aux_sym_recipe_body_repeat1),
	3593: uint16(218),
	3594: uint16(1),
	3595: uint16(sym_recipe_line_prefix),
	3596: uint16(385),
	3597: uint16(1),
	3598: uint16(sym_recipe_line),
	3599: uint16(271),
	3600: uint16(2),
	3601: uint16(anon_sym_AT),
	3602: uint16(anon_sym_DASH),
	3603: uint16(273),
	3604: uint16(2),
	3605: uint16(anon_sym_AT_DASH),
	3606: uint16(anon_sym_DASH_AT),
	3607: uint16(187),
	3608: uint16(2),
	3609: uint16(sym_interpolation),
	3610: uint16(aux_sym_recipe_line_repeat1),
	3611: uint16(5),
	3612: uint16(29),
	3613: uint16(1),
	3614: uint16(sym_comment),
	3615: uint16(348),
	3616: uint16(1),
	3617: uint16(sym__indent),
	3618: uint16(99),
	3619: uint16(1),
	3620: uint16(sym_recipe_body),
	3621: uint16(344),
	3622: uint16(3),
	3624: uint16(anon_sym_LBRACK),
	3625: uint16(anon_sym_AT),
	3626: uint16(346),
	3627: uint16(6),
	3628: uint16(anon_sym_alias),
	3629: uint16(anon_sym_export),
	3630: uint16(anon_sym_import),
	3631: uint16(anon_sym_mod),
	3632: uint16(anon_sym_set),
	3633: uint16(sym_identifier),
	3634: uint16(5),
	3635: uint16(29),
	3636: uint16(1),
	3637: uint16(sym_comment),
	3638: uint16(348),
	3639: uint16(1),
	3640: uint16(sym__indent),
	3641: uint16(114),
	3642: uint16(1),
	3643: uint16(sym_recipe_body),
	3644: uint16(350),
	3645: uint16(3),
	3647: uint16(anon_sym_LBRACK),
	3648: uint16(anon_sym_AT),
	3649: uint16(352),
	3650: uint16(6),
	3651: uint16(anon_sym_alias),
	3652: uint16(anon_sym_export),
	3653: uint16(anon_sym_import),
	3654: uint16(anon_sym_mod),
	3655: uint16(anon_sym_set),
	3656: uint16(sym_identifier),
	3657: uint16(10),
	3658: uint16(29),
	3659: uint16(1),
	3660: uint16(sym_comment),
	3661: uint16(285),
	3662: uint16(1),
	3663: uint16(sym_identifier),
	3664: uint16(293),
	3665: uint16(1),
	3666: uint16(anon_sym_DOLLAR),
	3667: uint16(354),
	3668: uint16(1),
	3669: uint16(anon_sym_COLON_EQ),
	3670: uint16(356),
	3671: uint16(1),
	3672: uint16(anon_sym_COLON),
	3673: uint16(141),
	3674: uint16(1),
	3675: uint16(aux_sym_parameters_repeat1),
	3676: uint16(169),
	3677: uint16(1),
	3678: uint16(sym_parameter),
	3679: uint16(315),
	3680: uint16(1),
	3681: uint16(sym_parameters),
	3682: uint16(349),
	3683: uint16(1),
	3684: uint16(sym_variadic_parameter),
	3685: uint16(289),
	3686: uint16(2),
	3687: uint16(anon_sym_PLUS),
	3688: uint16(anon_sym_STAR),
	3689: uint16(3),
	3690: uint16(29),
	3691: uint16(1),
	3692: uint16(sym_comment),
	3693: uint16(358),
	3694: uint16(3),
	3696: uint16(anon_sym_LBRACK),
	3697: uint16(anon_sym_AT),
	3698: uint16(360),
	3699: uint16(6),
	3700: uint16(anon_sym_alias),
	3701: uint16(anon_sym_export),
	3702: uint16(anon_sym_import),
	3703: uint16(anon_sym_mod),
	3704: uint16(anon_sym_set),
	3705: uint16(sym_identifier),
	3706: uint16(3),
	3707: uint16(29),
	3708: uint16(1),
	3709: uint16(sym_comment),
	3710: uint16(362),
	3711: uint16(3),
	3713: uint16(anon_sym_LBRACK),
	3714: uint16(anon_sym_AT),
	3715: uint16(364),
	3716: uint16(6),
	3717: uint16(anon_sym_alias),
	3718: uint16(anon_sym_export),
	3719: uint16(anon_sym_import),
	3720: uint16(anon_sym_mod),
	3721: uint16(anon_sym_set),
	3722: uint16(sym_identifier),
	3723: uint16(3),
	3724: uint16(29),
	3725: uint16(1),
	3726: uint16(sym_comment),
	3727: uint16(366),
	3728: uint16(3),
	3730: uint16(anon_sym_LBRACK),
	3731: uint16(anon_sym_AT),
	3732: uint16(368),
	3733: uint16(6),
	3734: uint16(anon_sym_alias),
	3735: uint16(anon_sym_export),
	3736: uint16(anon_sym_import),
	3737: uint16(anon_sym_mod),
	3738: uint16(anon_sym_set),
	3739: uint16(sym_identifier),
	3740: uint16(3),
	3741: uint16(29),
	3742: uint16(1),
	3743: uint16(sym_comment),
	3744: uint16(370),
	3745: uint16(3),
	3747: uint16(anon_sym_LBRACK),
	3748: uint16(anon_sym_AT),
	3749: uint16(372),
	3750: uint16(6),
	3751: uint16(anon_sym_alias),
	3752: uint16(anon_sym_export),
	3753: uint16(anon_sym_import),
	3754: uint16(anon_sym_mod),
	3755: uint16(anon_sym_set),
	3756: uint16(sym_identifier),
	3757: uint16(8),
	3758: uint16(29),
	3759: uint16(1),
	3760: uint16(sym_comment),
	3761: uint16(211),
	3762: uint16(1),
	3763: uint16(aux_sym_string_token1),
	3764: uint16(213),
	3765: uint16(1),
	3766: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	3767: uint16(215),
	3768: uint16(1),
	3769: uint16(anon_sym_DQUOTE),
	3770: uint16(217),
	3771: uint16(1),
	3772: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	3773: uint16(374),
	3774: uint16(1),
	3775: uint16(anon_sym_LBRACK),
	3776: uint16(376),
	3777: uint16(2),
	3778: uint16(anon_sym_true),
	3779: uint16(anon_sym_false),
	3780: uint16(343),
	3781: uint16(2),
	3782: uint16(sym_boolean),
	3783: uint16(sym_string),
	3784: uint16(4),
	3785: uint16(29),
	3786: uint16(1),
	3787: uint16(sym_comment),
	3788: uint16(378),
	3789: uint16(1),
	3790: uint16(anon_sym_SLASH),
	3791: uint16(380),
	3792: uint16(1),
	3793: uint16(anon_sym_PLUS),
	3794: uint16(336),
	3795: uint16(7),
	3796: uint16(anon_sym_COMMA),
	3797: uint16(anon_sym_LBRACE),
	3798: uint16(anon_sym_EQ_EQ),
	3799: uint16(anon_sym_BANG_EQ),
	3800: uint16(anon_sym_EQ_TILDE),
	3801: uint16(anon_sym_RPAREN),
	3802: uint16(anon_sym_RBRACE_RBRACE),
	3803: uint16(9),
	3804: uint16(29),
	3805: uint16(1),
	3806: uint16(sym_comment),
	3807: uint16(285),
	3808: uint16(1),
	3809: uint16(sym_identifier),
	3810: uint16(293),
	3811: uint16(1),
	3812: uint16(anon_sym_DOLLAR),
	3813: uint16(382),
	3814: uint16(1),
	3815: uint16(anon_sym_COLON),
	3816: uint16(141),
	3817: uint16(1),
	3818: uint16(aux_sym_parameters_repeat1),
	3819: uint16(169),
	3820: uint16(1),
	3821: uint16(sym_parameter),
	3822: uint16(349),
	3823: uint16(1),
	3824: uint16(sym_variadic_parameter),
	3825: uint16(355),
	3826: uint16(1),
	3827: uint16(sym_parameters),
	3828: uint16(289),
	3829: uint16(2),
	3830: uint16(anon_sym_PLUS),
	3831: uint16(anon_sym_STAR),
	3832: uint16(4),
	3833: uint16(29),
	3834: uint16(1),
	3835: uint16(sym_comment),
	3836: uint16(378),
	3837: uint16(1),
	3838: uint16(anon_sym_SLASH),
	3839: uint16(380),
	3840: uint16(1),
	3841: uint16(anon_sym_PLUS),
	3842: uint16(311),
	3843: uint16(7),
	3844: uint16(anon_sym_COMMA),
	3845: uint16(anon_sym_LBRACE),
	3846: uint16(anon_sym_EQ_EQ),
	3847: uint16(anon_sym_BANG_EQ),
	3848: uint16(anon_sym_EQ_TILDE),
	3849: uint16(anon_sym_RPAREN),
	3850: uint16(anon_sym_RBRACE_RBRACE),
	3851: uint16(9),
	3852: uint16(29),
	3853: uint16(1),
	3854: uint16(sym_comment),
	3855: uint16(285),
	3856: uint16(1),
	3857: uint16(sym_identifier),
	3858: uint16(293),
	3859: uint16(1),
	3860: uint16(anon_sym_DOLLAR),
	3861: uint16(384),
	3862: uint16(1),
	3863: uint16(anon_sym_COLON),
	3864: uint16(141),
	3865: uint16(1),
	3866: uint16(aux_sym_parameters_repeat1),
	3867: uint16(169),
	3868: uint16(1),
	3869: uint16(sym_parameter),
	3870: uint16(320),
	3871: uint16(1),
	3872: uint16(sym_parameters),
	3873: uint16(349),
	3874: uint16(1),
	3875: uint16(sym_variadic_parameter),
	3876: uint16(289),
	3877: uint16(2),
	3878: uint16(anon_sym_PLUS),
	3879: uint16(anon_sym_STAR),
	3880: uint16(3),
	3881: uint16(29),
	3882: uint16(1),
	3883: uint16(sym_comment),
	3884: uint16(386),
	3885: uint16(3),
	3887: uint16(anon_sym_LBRACK),
	3888: uint16(anon_sym_AT),
	3889: uint16(388),
	3890: uint16(6),
	3891: uint16(anon_sym_alias),
	3892: uint16(anon_sym_export),
	3893: uint16(anon_sym_import),
	3894: uint16(anon_sym_mod),
	3895: uint16(anon_sym_set),
	3896: uint16(sym_identifier),
	3897: uint16(3),
	3898: uint16(29),
	3899: uint16(1),
	3900: uint16(sym_comment),
	3901: uint16(350),
	3902: uint16(3),
	3904: uint16(anon_sym_LBRACK),
	3905: uint16(anon_sym_AT),
	3906: uint16(352),
	3907: uint16(6),
	3908: uint16(anon_sym_alias),
	3909: uint16(anon_sym_export),
	3910: uint16(anon_sym_import),
	3911: uint16(anon_sym_mod),
	3912: uint16(anon_sym_set),
	3913: uint16(sym_identifier),
	3914: uint16(3),
	3915: uint16(29),
	3916: uint16(1),
	3917: uint16(sym_comment),
	3918: uint16(390),
	3919: uint16(3),
	3921: uint16(anon_sym_LBRACK),
	3922: uint16(anon_sym_AT),
	3923: uint16(392),
	3924: uint16(6),
	3925: uint16(anon_sym_alias),
	3926: uint16(anon_sym_export),
	3927: uint16(anon_sym_import),
	3928: uint16(anon_sym_mod),
	3929: uint16(anon_sym_set),
	3930: uint16(sym_identifier),
	3931: uint16(3),
	3932: uint16(29),
	3933: uint16(1),
	3934: uint16(sym_comment),
	3935: uint16(394),
	3936: uint16(3),
	3938: uint16(anon_sym_LBRACK),
	3939: uint16(anon_sym_AT),
	3940: uint16(396),
	3941: uint16(6),
	3942: uint16(anon_sym_alias),
	3943: uint16(anon_sym_export),
	3944: uint16(anon_sym_import),
	3945: uint16(anon_sym_mod),
	3946: uint16(anon_sym_set),
	3947: uint16(sym_identifier),
	3948: uint16(3),
	3949: uint16(29),
	3950: uint16(1),
	3951: uint16(sym_comment),
	3952: uint16(380),
	3953: uint16(1),
	3954: uint16(anon_sym_PLUS),
	3955: uint16(265),
	3956: uint16(8),
	3957: uint16(anon_sym_COMMA),
	3958: uint16(anon_sym_SLASH),
	3959: uint16(anon_sym_LBRACE),
	3960: uint16(anon_sym_EQ_EQ),
	3961: uint16(anon_sym_BANG_EQ),
	3962: uint16(anon_sym_EQ_TILDE),
	3963: uint16(anon_sym_RPAREN),
	3964: uint16(anon_sym_RBRACE_RBRACE),
	3965: uint16(3),
	3966: uint16(29),
	3967: uint16(1),
	3968: uint16(sym_comment),
	3969: uint16(398),
	3970: uint16(3),
	3972: uint16(anon_sym_LBRACK),
	3973: uint16(anon_sym_AT),
	3974: uint16(400),
	3975: uint16(6),
	3976: uint16(anon_sym_alias),
	3977: uint16(anon_sym_export),
	3978: uint16(anon_sym_import),
	3979: uint16(anon_sym_mod),
	3980: uint16(anon_sym_set),
	3981: uint16(sym_identifier),
	3982: uint16(3),
	3983: uint16(29),
	3984: uint16(1),
	3985: uint16(sym_comment),
	3986: uint16(402),
	3987: uint16(3),
	3989: uint16(anon_sym_LBRACK),
	3990: uint16(anon_sym_AT),
	3991: uint16(404),
	3992: uint16(6),
	3993: uint16(anon_sym_alias),
	3994: uint16(anon_sym_export),
	3995: uint16(anon_sym_import),
	3996: uint16(anon_sym_mod),
	3997: uint16(anon_sym_set),
	3998: uint16(sym_identifier),
	3999: uint16(3),
	4000: uint16(29),
	4001: uint16(1),
	4002: uint16(sym_comment),
	4003: uint16(406),
	4004: uint16(3),
	4006: uint16(anon_sym_LBRACK),
	4007: uint16(anon_sym_AT),
	4008: uint16(408),
	4009: uint16(6),
	4010: uint16(anon_sym_alias),
	4011: uint16(anon_sym_export),
	4012: uint16(anon_sym_import),
	4013: uint16(anon_sym_mod),
	4014: uint16(anon_sym_set),
	4015: uint16(sym_identifier),
	4016: uint16(9),
	4017: uint16(29),
	4018: uint16(1),
	4019: uint16(sym_comment),
	4020: uint16(285),
	4021: uint16(1),
	4022: uint16(sym_identifier),
	4023: uint16(291),
	4024: uint16(1),
	4025: uint16(anon_sym_COLON),
	4026: uint16(293),
	4027: uint16(1),
	4028: uint16(anon_sym_DOLLAR),
	4029: uint16(141),
	4030: uint16(1),
	4031: uint16(aux_sym_parameters_repeat1),
	4032: uint16(169),
	4033: uint16(1),
	4034: uint16(sym_parameter),
	4035: uint16(329),
	4036: uint16(1),
	4037: uint16(sym_parameters),
	4038: uint16(349),
	4039: uint16(1),
	4040: uint16(sym_variadic_parameter),
	4041: uint16(289),
	4042: uint16(2),
	4043: uint16(anon_sym_PLUS),
	4044: uint16(anon_sym_STAR),
	4045: uint16(9),
	4046: uint16(29),
	4047: uint16(1),
	4048: uint16(sym_comment),
	4049: uint16(285),
	4050: uint16(1),
	4051: uint16(sym_identifier),
	4052: uint16(293),
	4053: uint16(1),
	4054: uint16(anon_sym_DOLLAR),
	4055: uint16(410),
	4056: uint16(1),
	4057: uint16(anon_sym_COLON),
	4058: uint16(141),
	4059: uint16(1),
	4060: uint16(aux_sym_parameters_repeat1),
	4061: uint16(169),
	4062: uint16(1),
	4063: uint16(sym_parameter),
	4064: uint16(315),
	4065: uint16(1),
	4066: uint16(sym_parameters),
	4067: uint16(349),
	4068: uint16(1),
	4069: uint16(sym_variadic_parameter),
	4070: uint16(289),
	4071: uint16(2),
	4072: uint16(anon_sym_PLUS),
	4073: uint16(anon_sym_STAR),
	4074: uint16(3),
	4075: uint16(29),
	4076: uint16(1),
	4077: uint16(sym_comment),
	4078: uint16(412),
	4079: uint16(3),
	4081: uint16(anon_sym_LBRACK),
	4082: uint16(anon_sym_AT),
	4083: uint16(414),
	4084: uint16(6),
	4085: uint16(anon_sym_alias),
	4086: uint16(anon_sym_export),
	4087: uint16(anon_sym_import),
	4088: uint16(anon_sym_mod),
	4089: uint16(anon_sym_set),
	4090: uint16(sym_identifier),
	4091: uint16(3),
	4092: uint16(29),
	4093: uint16(1),
	4094: uint16(sym_comment),
	4095: uint16(416),
	4096: uint16(3),
	4098: uint16(anon_sym_LBRACK),
	4099: uint16(anon_sym_AT),
	4100: uint16(418),
	4101: uint16(6),
	4102: uint16(anon_sym_alias),
	4103: uint16(anon_sym_export),
	4104: uint16(anon_sym_import),
	4105: uint16(anon_sym_mod),
	4106: uint16(anon_sym_set),
	4107: uint16(sym_identifier),
	4108: uint16(3),
	4109: uint16(29),
	4110: uint16(1),
	4111: uint16(sym_comment),
	4112: uint16(420),
	4113: uint16(3),
	4115: uint16(anon_sym_LBRACK),
	4116: uint16(anon_sym_AT),
	4117: uint16(422),
	4118: uint16(6),
	4119: uint16(anon_sym_alias),
	4120: uint16(anon_sym_export),
	4121: uint16(anon_sym_import),
	4122: uint16(anon_sym_mod),
	4123: uint16(anon_sym_set),
	4124: uint16(sym_identifier),
	4125: uint16(3),
	4126: uint16(29),
	4127: uint16(1),
	4128: uint16(sym_comment),
	4129: uint16(424),
	4130: uint16(3),
	4132: uint16(anon_sym_LBRACK),
	4133: uint16(anon_sym_AT),
	4134: uint16(426),
	4135: uint16(6),
	4136: uint16(anon_sym_alias),
	4137: uint16(anon_sym_export),
	4138: uint16(anon_sym_import),
	4139: uint16(anon_sym_mod),
	4140: uint16(anon_sym_set),
	4141: uint16(sym_identifier),
	4142: uint16(3),
	4143: uint16(29),
	4144: uint16(1),
	4145: uint16(sym_comment),
	4146: uint16(428),
	4147: uint16(3),
	4149: uint16(anon_sym_LBRACK),
	4150: uint16(anon_sym_AT),
	4151: uint16(430),
	4152: uint16(6),
	4153: uint16(anon_sym_alias),
	4154: uint16(anon_sym_export),
	4155: uint16(anon_sym_import),
	4156: uint16(anon_sym_mod),
	4157: uint16(anon_sym_set),
	4158: uint16(sym_identifier),
	4159: uint16(3),
	4160: uint16(29),
	4161: uint16(1),
	4162: uint16(sym_comment),
	4163: uint16(432),
	4164: uint16(3),
	4166: uint16(anon_sym_LBRACK),
	4167: uint16(anon_sym_AT),
	4168: uint16(434),
	4169: uint16(6),
	4170: uint16(anon_sym_alias),
	4171: uint16(anon_sym_export),
	4172: uint16(anon_sym_import),
	4173: uint16(anon_sym_mod),
	4174: uint16(anon_sym_set),
	4175: uint16(sym_identifier),
	4176: uint16(3),
	4177: uint16(29),
	4178: uint16(1),
	4179: uint16(sym_comment),
	4180: uint16(436),
	4181: uint16(3),
	4183: uint16(anon_sym_LBRACK),
	4184: uint16(anon_sym_AT),
	4185: uint16(438),
	4186: uint16(6),
	4187: uint16(anon_sym_alias),
	4188: uint16(anon_sym_export),
	4189: uint16(anon_sym_import),
	4190: uint16(anon_sym_mod),
	4191: uint16(anon_sym_set),
	4192: uint16(sym_identifier),
	4193: uint16(3),
	4194: uint16(29),
	4195: uint16(1),
	4196: uint16(sym_comment),
	4197: uint16(440),
	4198: uint16(3),
	4200: uint16(anon_sym_LBRACK),
	4201: uint16(anon_sym_AT),
	4202: uint16(442),
	4203: uint16(6),
	4204: uint16(anon_sym_alias),
	4205: uint16(anon_sym_export),
	4206: uint16(anon_sym_import),
	4207: uint16(anon_sym_mod),
	4208: uint16(anon_sym_set),
	4209: uint16(sym_identifier),
	4210: uint16(3),
	4211: uint16(29),
	4212: uint16(1),
	4213: uint16(sym_comment),
	4214: uint16(444),
	4215: uint16(3),
	4217: uint16(anon_sym_LBRACK),
	4218: uint16(anon_sym_AT),
	4219: uint16(446),
	4220: uint16(6),
	4221: uint16(anon_sym_alias),
	4222: uint16(anon_sym_export),
	4223: uint16(anon_sym_import),
	4224: uint16(anon_sym_mod),
	4225: uint16(anon_sym_set),
	4226: uint16(sym_identifier),
	4227: uint16(3),
	4228: uint16(29),
	4229: uint16(1),
	4230: uint16(sym_comment),
	4231: uint16(448),
	4232: uint16(3),
	4234: uint16(anon_sym_LBRACK),
	4235: uint16(anon_sym_AT),
	4236: uint16(450),
	4237: uint16(6),
	4238: uint16(anon_sym_alias),
	4239: uint16(anon_sym_export),
	4240: uint16(anon_sym_import),
	4241: uint16(anon_sym_mod),
	4242: uint16(anon_sym_set),
	4243: uint16(sym_identifier),
	4244: uint16(3),
	4245: uint16(29),
	4246: uint16(1),
	4247: uint16(sym_comment),
	4248: uint16(452),
	4249: uint16(3),
	4251: uint16(anon_sym_LBRACK),
	4252: uint16(anon_sym_AT),
	4253: uint16(454),
	4254: uint16(6),
	4255: uint16(anon_sym_alias),
	4256: uint16(anon_sym_export),
	4257: uint16(anon_sym_import),
	4258: uint16(anon_sym_mod),
	4259: uint16(anon_sym_set),
	4260: uint16(sym_identifier),
	4261: uint16(3),
	4262: uint16(29),
	4263: uint16(1),
	4264: uint16(sym_comment),
	4265: uint16(456),
	4266: uint16(3),
	4268: uint16(anon_sym_LBRACK),
	4269: uint16(anon_sym_AT),
	4270: uint16(458),
	4271: uint16(6),
	4272: uint16(anon_sym_alias),
	4273: uint16(anon_sym_export),
	4274: uint16(anon_sym_import),
	4275: uint16(anon_sym_mod),
	4276: uint16(anon_sym_set),
	4277: uint16(sym_identifier),
	4278: uint16(3),
	4279: uint16(29),
	4280: uint16(1),
	4281: uint16(sym_comment),
	4282: uint16(460),
	4283: uint16(3),
	4285: uint16(anon_sym_LBRACK),
	4286: uint16(anon_sym_AT),
	4287: uint16(462),
	4288: uint16(6),
	4289: uint16(anon_sym_alias),
	4290: uint16(anon_sym_export),
	4291: uint16(anon_sym_import),
	4292: uint16(anon_sym_mod),
	4293: uint16(anon_sym_set),
	4294: uint16(sym_identifier),
	4295: uint16(3),
	4296: uint16(29),
	4297: uint16(1),
	4298: uint16(sym_comment),
	4299: uint16(464),
	4300: uint16(3),
	4302: uint16(anon_sym_LBRACK),
	4303: uint16(anon_sym_AT),
	4304: uint16(466),
	4305: uint16(6),
	4306: uint16(anon_sym_alias),
	4307: uint16(anon_sym_export),
	4308: uint16(anon_sym_import),
	4309: uint16(anon_sym_mod),
	4310: uint16(anon_sym_set),
	4311: uint16(sym_identifier),
	4312: uint16(3),
	4313: uint16(29),
	4314: uint16(1),
	4315: uint16(sym_comment),
	4316: uint16(468),
	4317: uint16(3),
	4319: uint16(anon_sym_LBRACK),
	4320: uint16(anon_sym_AT),
	4321: uint16(470),
	4322: uint16(6),
	4323: uint16(anon_sym_alias),
	4324: uint16(anon_sym_export),
	4325: uint16(anon_sym_import),
	4326: uint16(anon_sym_mod),
	4327: uint16(anon_sym_set),
	4328: uint16(sym_identifier),
	4329: uint16(8),
	4330: uint16(29),
	4331: uint16(1),
	4332: uint16(sym_comment),
	4333: uint16(472),
	4334: uint16(1),
	4335: uint16(sym_identifier),
	4336: uint16(474),
	4337: uint16(1),
	4338: uint16(anon_sym_LPAREN),
	4339: uint16(476),
	4340: uint16(1),
	4341: uint16(anon_sym_AMP_AMP),
	4342: uint16(478),
	4343: uint16(1),
	4344: uint16(sym__newline),
	4345: uint16(212),
	4346: uint16(1),
	4347: uint16(sym_dependency_expression),
	4348: uint16(364),
	4349: uint16(1),
	4350: uint16(sym_dependencies),
	4351: uint16(140),
	4352: uint16(2),
	4353: uint16(sym_dependency),
	4354: uint16(aux_sym_dependencies_repeat1),
	4355: uint16(8),
	4356: uint16(29),
	4357: uint16(1),
	4358: uint16(sym_comment),
	4359: uint16(472),
	4360: uint16(1),
	4361: uint16(sym_identifier),
	4362: uint16(474),
	4363: uint16(1),
	4364: uint16(anon_sym_LPAREN),
	4365: uint16(476),
	4366: uint16(1),
	4367: uint16(anon_sym_AMP_AMP),
	4368: uint16(480),
	4369: uint16(1),
	4370: uint16(sym__newline),
	4371: uint16(212),
	4372: uint16(1),
	4373: uint16(sym_dependency_expression),
	4374: uint16(390),
	4375: uint16(1),
	4376: uint16(sym_dependencies),
	4377: uint16(140),
	4378: uint16(2),
	4379: uint16(sym_dependency),
	4380: uint16(aux_sym_dependencies_repeat1),
	4381: uint16(8),
	4382: uint16(29),
	4383: uint16(1),
	4384: uint16(sym_comment),
	4385: uint16(472),
	4386: uint16(1),
	4387: uint16(sym_identifier),
	4388: uint16(474),
	4389: uint16(1),
	4390: uint16(anon_sym_LPAREN),
	4391: uint16(476),
	4392: uint16(1),
	4393: uint16(anon_sym_AMP_AMP),
	4394: uint16(482),
	4395: uint16(1),
	4396: uint16(sym__newline),
	4397: uint16(212),
	4398: uint16(1),
	4399: uint16(sym_dependency_expression),
	4400: uint16(374),
	4401: uint16(1),
	4402: uint16(sym_dependencies),
	4403: uint16(140),
	4404: uint16(2),
	4405: uint16(sym_dependency),
	4406: uint16(aux_sym_dependencies_repeat1),
	4407: uint16(8),
	4408: uint16(29),
	4409: uint16(1),
	4410: uint16(sym_comment),
	4411: uint16(472),
	4412: uint16(1),
	4413: uint16(sym_identifier),
	4414: uint16(474),
	4415: uint16(1),
	4416: uint16(anon_sym_LPAREN),
	4417: uint16(476),
	4418: uint16(1),
	4419: uint16(anon_sym_AMP_AMP),
	4420: uint16(484),
	4421: uint16(1),
	4422: uint16(sym__newline),
	4423: uint16(212),
	4424: uint16(1),
	4425: uint16(sym_dependency_expression),
	4426: uint16(319),
	4427: uint16(1),
	4428: uint16(sym_dependencies),
	4429: uint16(140),
	4430: uint16(2),
	4431: uint16(sym_dependency),
	4432: uint16(aux_sym_dependencies_repeat1),
	4433: uint16(8),
	4434: uint16(19),
	4435: uint16(1),
	4436: uint16(anon_sym_LBRACK),
	4437: uint16(21),
	4438: uint16(1),
	4439: uint16(anon_sym_AT),
	4440: uint16(29),
	4441: uint16(1),
	4442: uint16(sym_comment),
	4443: uint16(486),
	4444: uint16(1),
	4445: uint16(sym_identifier),
	4446: uint16(488),
	4447: uint16(1),
	4448: uint16(anon_sym_alias),
	4449: uint16(490),
	4450: uint16(1),
	4451: uint16(anon_sym_import),
	4452: uint16(400),
	4453: uint16(1),
	4454: uint16(sym_recipe_header),
	4455: uint16(143),
	4456: uint16(2),
	4457: uint16(sym_attribute),
	4458: uint16(aux_sym_alias_repeat1),
	4459: uint16(8),
	4460: uint16(29),
	4461: uint16(1),
	4462: uint16(sym_comment),
	4463: uint16(472),
	4464: uint16(1),
	4465: uint16(sym_identifier),
	4466: uint16(474),
	4467: uint16(1),
	4468: uint16(anon_sym_LPAREN),
	4469: uint16(476),
	4470: uint16(1),
	4471: uint16(anon_sym_AMP_AMP),
	4472: uint16(492),
	4473: uint16(1),
	4474: uint16(sym__newline),
	4475: uint16(212),
	4476: uint16(1),
	4477: uint16(sym_dependency_expression),
	4478: uint16(378),
	4479: uint16(1),
	4480: uint16(sym_dependencies),
	4481: uint16(140),
	4482: uint16(2),
	4483: uint16(sym_dependency),
	4484: uint16(aux_sym_dependencies_repeat1),
	4485: uint16(9),
	4486: uint16(29),
	4487: uint16(1),
	4488: uint16(sym_comment),
	4489: uint16(67),
	4490: uint16(1),
	4491: uint16(aux_sym_string_token1),
	4492: uint16(69),
	4493: uint16(1),
	4494: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	4495: uint16(71),
	4496: uint16(1),
	4497: uint16(anon_sym_DQUOTE),
	4498: uint16(73),
	4499: uint16(1),
	4500: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4501: uint16(494),
	4502: uint16(1),
	4503: uint16(anon_sym_COMMA),
	4504: uint16(496),
	4505: uint16(1),
	4506: uint16(anon_sym_RBRACK),
	4507: uint16(133),
	4508: uint16(1),
	4509: uint16(aux_sym_setting_repeat1),
	4510: uint16(350),
	4511: uint16(1),
	4512: uint16(sym_string),
	4513: uint16(8),
	4514: uint16(29),
	4515: uint16(1),
	4516: uint16(sym_comment),
	4517: uint16(472),
	4518: uint16(1),
	4519: uint16(sym_identifier),
	4520: uint16(474),
	4521: uint16(1),
	4522: uint16(anon_sym_LPAREN),
	4523: uint16(476),
	4524: uint16(1),
	4525: uint16(anon_sym_AMP_AMP),
	4526: uint16(498),
	4527: uint16(1),
	4528: uint16(sym__newline),
	4529: uint16(212),
	4530: uint16(1),
	4531: uint16(sym_dependency_expression),
	4532: uint16(335),
	4533: uint16(1),
	4534: uint16(sym_dependencies),
	4535: uint16(140),
	4536: uint16(2),
	4537: uint16(sym_dependency),
	4538: uint16(aux_sym_dependencies_repeat1),
	4539: uint16(9),
	4540: uint16(29),
	4541: uint16(1),
	4542: uint16(sym_comment),
	4543: uint16(67),
	4544: uint16(1),
	4545: uint16(aux_sym_string_token1),
	4546: uint16(69),
	4547: uint16(1),
	4548: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	4549: uint16(71),
	4550: uint16(1),
	4551: uint16(anon_sym_DQUOTE),
	4552: uint16(73),
	4553: uint16(1),
	4554: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4555: uint16(494),
	4556: uint16(1),
	4557: uint16(anon_sym_COMMA),
	4558: uint16(500),
	4559: uint16(1),
	4560: uint16(anon_sym_RBRACK),
	4561: uint16(132),
	4562: uint16(1),
	4563: uint16(aux_sym_setting_repeat1),
	4564: uint16(359),
	4565: uint16(1),
	4566: uint16(sym_string),
	4567: uint16(9),
	4568: uint16(29),
	4569: uint16(1),
	4570: uint16(sym_comment),
	4571: uint16(67),
	4572: uint16(1),
	4573: uint16(aux_sym_string_token1),
	4574: uint16(69),
	4575: uint16(1),
	4576: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	4577: uint16(71),
	4578: uint16(1),
	4579: uint16(anon_sym_DQUOTE),
	4580: uint16(73),
	4581: uint16(1),
	4582: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4583: uint16(494),
	4584: uint16(1),
	4585: uint16(anon_sym_COMMA),
	4586: uint16(502),
	4587: uint16(1),
	4588: uint16(anon_sym_RBRACK),
	4589: uint16(144),
	4590: uint16(1),
	4591: uint16(aux_sym_setting_repeat1),
	4592: uint16(353),
	4593: uint16(1),
	4594: uint16(sym_string),
	4595: uint16(9),
	4596: uint16(29),
	4597: uint16(1),
	4598: uint16(sym_comment),
	4599: uint16(67),
	4600: uint16(1),
	4601: uint16(aux_sym_string_token1),
	4602: uint16(69),
	4603: uint16(1),
	4604: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	4605: uint16(71),
	4606: uint16(1),
	4607: uint16(anon_sym_DQUOTE),
	4608: uint16(73),
	4609: uint16(1),
	4610: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4611: uint16(494),
	4612: uint16(1),
	4613: uint16(anon_sym_COMMA),
	4614: uint16(504),
	4615: uint16(1),
	4616: uint16(anon_sym_RBRACK),
	4617: uint16(144),
	4618: uint16(1),
	4619: uint16(aux_sym_setting_repeat1),
	4620: uint16(371),
	4621: uint16(1),
	4622: uint16(sym_string),
	4623: uint16(3),
	4624: uint16(29),
	4625: uint16(1),
	4626: uint16(sym_comment),
	4627: uint16(506),
	4628: uint16(2),
	4629: uint16(anon_sym_AT),
	4630: uint16(anon_sym_DASH),
	4631: uint16(322),
	4632: uint16(6),
	4633: uint16(sym__dedent),
	4634: uint16(sym__newline),
	4635: uint16(sym_text),
	4636: uint16(anon_sym_AT_DASH),
	4637: uint16(anon_sym_DASH_AT),
	4638: uint16(anon_sym_LBRACE_LBRACE),
	4639: uint16(8),
	4640: uint16(29),
	4641: uint16(1),
	4642: uint16(sym_comment),
	4643: uint16(472),
	4644: uint16(1),
	4645: uint16(sym_identifier),
	4646: uint16(474),
	4647: uint16(1),
	4648: uint16(anon_sym_LPAREN),
	4649: uint16(476),
	4650: uint16(1),
	4651: uint16(anon_sym_AMP_AMP),
	4652: uint16(508),
	4653: uint16(1),
	4654: uint16(sym__newline),
	4655: uint16(212),
	4656: uint16(1),
	4657: uint16(sym_dependency_expression),
	4658: uint16(380),
	4659: uint16(1),
	4660: uint16(sym_dependencies),
	4661: uint16(140),
	4662: uint16(2),
	4663: uint16(sym_dependency),
	4664: uint16(aux_sym_dependencies_repeat1),
	4665: uint16(8),
	4666: uint16(29),
	4667: uint16(1),
	4668: uint16(sym_comment),
	4669: uint16(472),
	4670: uint16(1),
	4671: uint16(sym_identifier),
	4672: uint16(474),
	4673: uint16(1),
	4674: uint16(anon_sym_LPAREN),
	4675: uint16(476),
	4676: uint16(1),
	4677: uint16(anon_sym_AMP_AMP),
	4678: uint16(510),
	4679: uint16(1),
	4680: uint16(sym__newline),
	4681: uint16(212),
	4682: uint16(1),
	4683: uint16(sym_dependency_expression),
	4684: uint16(360),
	4685: uint16(1),
	4686: uint16(sym_dependencies),
	4687: uint16(140),
	4688: uint16(2),
	4689: uint16(sym_dependency),
	4690: uint16(aux_sym_dependencies_repeat1),
	4691: uint16(7),
	4692: uint16(29),
	4693: uint16(1),
	4694: uint16(sym_comment),
	4695: uint16(67),
	4696: uint16(1),
	4697: uint16(aux_sym_string_token1),
	4698: uint16(69),
	4699: uint16(1),
	4700: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	4701: uint16(71),
	4702: uint16(1),
	4703: uint16(anon_sym_DQUOTE),
	4704: uint16(73),
	4705: uint16(1),
	4706: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4707: uint16(512),
	4708: uint16(1),
	4709: uint16(sym_identifier),
	4710: uint16(284),
	4711: uint16(2),
	4712: uint16(sym_attribute_kv_argument),
	4713: uint16(sym_string),
	4714: uint16(6),
	4715: uint16(29),
	4716: uint16(1),
	4717: uint16(sym_comment),
	4718: uint16(514),
	4719: uint16(1),
	4720: uint16(anon_sym_else),
	4721: uint16(156),
	4722: uint16(1),
	4723: uint16(aux_sym_if_expression_repeat1),
	4724: uint16(192),
	4725: uint16(1),
	4726: uint16(sym_else_if_clause),
	4727: uint16(236),
	4728: uint16(1),
	4729: uint16(sym_else_clause),
	4730: uint16(47),
	4731: uint16(3),
	4732: uint16(sym__newline),
	4733: uint16(anon_sym_SLASH),
	4734: uint16(anon_sym_PLUS),
	4735: uint16(6),
	4736: uint16(29),
	4737: uint16(1),
	4738: uint16(sym_comment),
	4739: uint16(516),
	4740: uint16(1),
	4741: uint16(anon_sym_else),
	4742: uint16(150),
	4743: uint16(1),
	4744: uint16(aux_sym_if_expression_repeat1),
	4745: uint16(202),
	4746: uint16(1),
	4747: uint16(sym_else_if_clause),
	4748: uint16(273),
	4749: uint16(1),
	4750: uint16(sym_else_clause),
	4751: uint16(47),
	4752: uint16(3),
	4753: uint16(anon_sym_SLASH),
	4754: uint16(anon_sym_PLUS),
	4755: uint16(anon_sym_RBRACE),
	4756: uint16(7),
	4757: uint16(29),
	4758: uint16(1),
	4759: uint16(sym_comment),
	4760: uint16(472),
	4761: uint16(1),
	4762: uint16(sym_identifier),
	4763: uint16(474),
	4764: uint16(1),
	4765: uint16(anon_sym_LPAREN),
	4766: uint16(476),
	4767: uint16(1),
	4768: uint16(anon_sym_AMP_AMP),
	4769: uint16(518),
	4770: uint16(1),
	4771: uint16(sym__newline),
	4772: uint16(212),
	4773: uint16(1),
	4774: uint16(sym_dependency_expression),
	4775: uint16(147),
	4776: uint16(2),
	4777: uint16(sym_dependency),
	4778: uint16(aux_sym_dependencies_repeat1),
	4779: uint16(7),
	4780: uint16(29),
	4781: uint16(1),
	4782: uint16(sym_comment),
	4783: uint16(285),
	4784: uint16(1),
	4785: uint16(sym_identifier),
	4786: uint16(293),
	4787: uint16(1),
	4788: uint16(anon_sym_DOLLAR),
	4789: uint16(155),
	4790: uint16(1),
	4791: uint16(aux_sym_parameters_repeat1),
	4792: uint16(165),
	4793: uint16(1),
	4794: uint16(sym_parameter),
	4795: uint16(375),
	4796: uint16(1),
	4797: uint16(sym_variadic_parameter),
	4798: uint16(289),
	4799: uint16(2),
	4800: uint16(anon_sym_PLUS),
	4801: uint16(anon_sym_STAR),
	4802: uint16(6),
	4803: uint16(29),
	4804: uint16(1),
	4805: uint16(sym_comment),
	4806: uint16(516),
	4807: uint16(1),
	4808: uint16(anon_sym_else),
	4809: uint16(139),
	4810: uint16(1),
	4811: uint16(aux_sym_if_expression_repeat1),
	4812: uint16(202),
	4813: uint16(1),
	4814: uint16(sym_else_if_clause),
	4815: uint16(268),
	4816: uint16(1),
	4817: uint16(sym_else_clause),
	4818: uint16(41),
	4819: uint16(3),
	4820: uint16(anon_sym_SLASH),
	4821: uint16(anon_sym_PLUS),
	4822: uint16(anon_sym_RBRACE),
	4823: uint16(5),
	4824: uint16(29),
	4825: uint16(1),
	4826: uint16(sym_comment),
	4827: uint16(522),
	4828: uint16(1),
	4829: uint16(anon_sym_LBRACK),
	4830: uint16(525),
	4831: uint16(1),
	4832: uint16(anon_sym_AT),
	4833: uint16(143),
	4834: uint16(2),
	4835: uint16(sym_attribute),
	4836: uint16(aux_sym_alias_repeat1),
	4837: uint16(520),
	4838: uint16(3),
	4839: uint16(anon_sym_alias),
	4840: uint16(anon_sym_import),
	4841: uint16(sym_identifier),
	4842: uint16(5),
	4843: uint16(29),
	4844: uint16(1),
	4845: uint16(sym_comment),
	4846: uint16(527),
	4847: uint16(1),
	4848: uint16(anon_sym_COMMA),
	4849: uint16(144),
	4850: uint16(1),
	4851: uint16(aux_sym_setting_repeat1),
	4852: uint16(532),
	4853: uint16(2),
	4854: uint16(aux_sym_string_token1),
	4855: uint16(anon_sym_DQUOTE),
	4856: uint16(530),
	4857: uint16(3),
	4858: uint16(anon_sym_RBRACK),
	4859: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	4860: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4861: uint16(7),
	4862: uint16(29),
	4863: uint16(1),
	4864: uint16(sym_comment),
	4865: uint16(67),
	4866: uint16(1),
	4867: uint16(aux_sym_string_token1),
	4868: uint16(69),
	4869: uint16(1),
	4870: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	4871: uint16(71),
	4872: uint16(1),
	4873: uint16(anon_sym_DQUOTE),
	4874: uint16(73),
	4875: uint16(1),
	4876: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4877: uint16(534),
	4878: uint16(1),
	4879: uint16(sym_identifier),
	4880: uint16(311),
	4881: uint16(2),
	4882: uint16(sym_attribute_kv_argument),
	4883: uint16(sym_string),
	4884: uint16(6),
	4885: uint16(29),
	4886: uint16(1),
	4887: uint16(sym_comment),
	4888: uint16(514),
	4889: uint16(1),
	4890: uint16(anon_sym_else),
	4891: uint16(138),
	4892: uint16(1),
	4893: uint16(aux_sym_if_expression_repeat1),
	4894: uint16(192),
	4895: uint16(1),
	4896: uint16(sym_else_if_clause),
	4897: uint16(225),
	4898: uint16(1),
	4899: uint16(sym_else_clause),
	4900: uint16(41),
	4901: uint16(3),
	4902: uint16(sym__newline),
	4903: uint16(anon_sym_SLASH),
	4904: uint16(anon_sym_PLUS),
	4905: uint16(7),
	4906: uint16(29),
	4907: uint16(1),
	4908: uint16(sym_comment),
	4909: uint16(536),
	4910: uint16(1),
	4911: uint16(sym_identifier),
	4912: uint16(539),
	4913: uint16(1),
	4914: uint16(anon_sym_LPAREN),
	4915: uint16(542),
	4916: uint16(1),
	4917: uint16(anon_sym_AMP_AMP),
	4918: uint16(545),
	4919: uint16(1),
	4920: uint16(sym__newline),
	4921: uint16(212),
	4922: uint16(1),
	4923: uint16(sym_dependency_expression),
	4924: uint16(147),
	4925: uint16(2),
	4926: uint16(sym_dependency),
	4927: uint16(aux_sym_dependencies_repeat1),
	4928: uint16(7),
	4929: uint16(29),
	4930: uint16(1),
	4931: uint16(sym_comment),
	4932: uint16(67),
	4933: uint16(1),
	4934: uint16(aux_sym_string_token1),
	4935: uint16(69),
	4936: uint16(1),
	4937: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	4938: uint16(71),
	4939: uint16(1),
	4940: uint16(anon_sym_DQUOTE),
	4941: uint16(73),
	4942: uint16(1),
	4943: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4944: uint16(547),
	4945: uint16(1),
	4946: uint16(sym_identifier),
	4947: uint16(250),
	4948: uint16(2),
	4949: uint16(sym_attribute_kv_argument),
	4950: uint16(sym_string),
	4951: uint16(7),
	4952: uint16(29),
	4953: uint16(1),
	4954: uint16(sym_comment),
	4955: uint16(67),
	4956: uint16(1),
	4957: uint16(aux_sym_string_token1),
	4958: uint16(69),
	4959: uint16(1),
	4960: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	4961: uint16(71),
	4962: uint16(1),
	4963: uint16(anon_sym_DQUOTE),
	4964: uint16(73),
	4965: uint16(1),
	4966: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4967: uint16(549),
	4968: uint16(1),
	4969: uint16(anon_sym_RBRACK),
	4970: uint16(131),
	4971: uint16(1),
	4972: uint16(sym_string),
	4973: uint16(5),
	4974: uint16(29),
	4975: uint16(1),
	4976: uint16(sym_comment),
	4977: uint16(551),
	4978: uint16(1),
	4979: uint16(anon_sym_else),
	4980: uint16(150),
	4981: uint16(1),
	4982: uint16(aux_sym_if_expression_repeat1),
	4983: uint16(202),
	4984: uint16(1),
	4985: uint16(sym_else_if_clause),
	4986: uint16(162),
	4987: uint16(3),
	4988: uint16(anon_sym_SLASH),
	4989: uint16(anon_sym_PLUS),
	4990: uint16(anon_sym_RBRACE),
	4991: uint16(6),
	4992: uint16(3),
	4993: uint16(1),
	4994: uint16(sym_comment),
	4995: uint16(554),
	4996: uint16(1),
	4997: uint16(aux_sym__raw_string_indented_token1),
	4998: uint16(556),
	4999: uint16(1),
	5000: uint16(anon_sym_BQUOTE),
	5001: uint16(558),
	5002: uint16(1),
	5003: uint16(anon_sym_LBRACE_LBRACE),
	5004: uint16(362),
	5005: uint16(1),
	5006: uint16(sym_command_body),
	5007: uint16(186),
	5008: uint16(2),
	5009: uint16(sym_interpolation),
	5010: uint16(aux_sym_command_body_repeat1),
	5011: uint16(6),
	5012: uint16(3),
	5013: uint16(1),
	5014: uint16(sym_comment),
	5015: uint16(560),
	5016: uint16(1),
	5017: uint16(aux_sym__raw_string_indented_token1),
	5018: uint16(562),
	5019: uint16(1),
	5020: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	5021: uint16(564),
	5022: uint16(1),
	5023: uint16(anon_sym_LBRACE_LBRACE),
	5024: uint16(404),
	5025: uint16(1),
	5026: uint16(sym_command_body),
	5027: uint16(182),
	5028: uint16(2),
	5029: uint16(sym_interpolation),
	5030: uint16(aux_sym_command_body_repeat1),
	5031: uint16(7),
	5032: uint16(29),
	5033: uint16(1),
	5034: uint16(sym_comment),
	5035: uint16(67),
	5036: uint16(1),
	5037: uint16(aux_sym_string_token1),
	5038: uint16(69),
	5039: uint16(1),
	5040: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	5041: uint16(71),
	5042: uint16(1),
	5043: uint16(anon_sym_DQUOTE),
	5044: uint16(73),
	5045: uint16(1),
	5046: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5047: uint16(566),
	5048: uint16(1),
	5049: uint16(anon_sym_RBRACK),
	5050: uint16(129),
	5051: uint16(1),
	5052: uint16(sym_string),
	5053: uint16(3),
	5054: uint16(29),
	5055: uint16(1),
	5056: uint16(sym_comment),
	5057: uint16(570),
	5058: uint16(2),
	5059: uint16(aux_sym_string_token1),
	5060: uint16(anon_sym_DQUOTE),
	5061: uint16(568),
	5062: uint16(4),
	5063: uint16(anon_sym_COMMA),
	5064: uint16(anon_sym_RBRACK),
	5065: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	5066: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5067: uint16(5),
	5068: uint16(29),
	5069: uint16(1),
	5070: uint16(sym_comment),
	5071: uint16(572),
	5072: uint16(1),
	5073: uint16(sym_identifier),
	5074: uint16(577),
	5075: uint16(1),
	5076: uint16(anon_sym_DOLLAR),
	5077: uint16(575),
	5078: uint16(2),
	5079: uint16(anon_sym_PLUS),
	5080: uint16(anon_sym_STAR),
	5081: uint16(155),
	5082: uint16(2),
	5083: uint16(sym_parameter),
	5084: uint16(aux_sym_parameters_repeat1),
	5085: uint16(5),
	5086: uint16(29),
	5087: uint16(1),
	5088: uint16(sym_comment),
	5089: uint16(580),
	5090: uint16(1),
	5091: uint16(anon_sym_else),
	5092: uint16(156),
	5093: uint16(1),
	5094: uint16(aux_sym_if_expression_repeat1),
	5095: uint16(192),
	5096: uint16(1),
	5097: uint16(sym_else_if_clause),
	5098: uint16(162),
	5099: uint16(3),
	5100: uint16(sym__newline),
	5101: uint16(anon_sym_SLASH),
	5102: uint16(anon_sym_PLUS),
	5103: uint16(3),
	5104: uint16(29),
	5105: uint16(1),
	5106: uint16(sym_comment),
	5107: uint16(585),
	5108: uint16(1),
	5109: uint16(anon_sym_EQ),
	5110: uint16(583),
	5111: uint16(5),
	5112: uint16(anon_sym_PLUS),
	5113: uint16(anon_sym_COLON),
	5114: uint16(anon_sym_DOLLAR),
	5115: uint16(anon_sym_STAR),
	5116: uint16(sym_identifier),
	5117: uint16(6),
	5118: uint16(3),
	5119: uint16(1),
	5120: uint16(sym_comment),
	5121: uint16(554),
	5122: uint16(1),
	5123: uint16(aux_sym__raw_string_indented_token1),
	5124: uint16(558),
	5125: uint16(1),
	5126: uint16(anon_sym_LBRACE_LBRACE),
	5127: uint16(587),
	5128: uint16(1),
	5129: uint16(anon_sym_BQUOTE),
	5130: uint16(346),
	5131: uint16(1),
	5132: uint16(sym_command_body),
	5133: uint16(186),
	5134: uint16(2),
	5135: uint16(sym_interpolation),
	5136: uint16(aux_sym_command_body_repeat1),
	5137: uint16(6),
	5138: uint16(3),
	5139: uint16(1),
	5140: uint16(sym_comment),
	5141: uint16(560),
	5142: uint16(1),
	5143: uint16(aux_sym__raw_string_indented_token1),
	5144: uint16(564),
	5145: uint16(1),
	5146: uint16(anon_sym_LBRACE_LBRACE),
	5147: uint16(589),
	5148: uint16(1),
	5149: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	5150: uint16(384),
	5151: uint16(1),
	5152: uint16(sym_command_body),
	5153: uint16(182),
	5154: uint16(2),
	5155: uint16(sym_interpolation),
	5156: uint16(aux_sym_command_body_repeat1),
	5157: uint16(6),
	5158: uint16(3),
	5159: uint16(1),
	5160: uint16(sym_comment),
	5161: uint16(560),
	5162: uint16(1),
	5163: uint16(aux_sym__raw_string_indented_token1),
	5164: uint16(564),
	5165: uint16(1),
	5166: uint16(anon_sym_LBRACE_LBRACE),
	5167: uint16(591),
	5168: uint16(1),
	5169: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	5170: uint16(358),
	5171: uint16(1),
	5172: uint16(sym_command_body),
	5173: uint16(182),
	5174: uint16(2),
	5175: uint16(sym_interpolation),
	5176: uint16(aux_sym_command_body_repeat1),
	5177: uint16(3),
	5178: uint16(29),
	5179: uint16(1),
	5180: uint16(sym_comment),
	5181: uint16(595),
	5182: uint16(1),
	5183: uint16(anon_sym_EQ),
	5184: uint16(593),
	5185: uint16(5),
	5186: uint16(anon_sym_PLUS),
	5187: uint16(anon_sym_COLON),
	5188: uint16(anon_sym_DOLLAR),
	5189: uint16(anon_sym_STAR),
	5190: uint16(sym_identifier),
	5191: uint16(6),
	5192: uint16(3),
	5193: uint16(1),
	5194: uint16(sym_comment),
	5195: uint16(554),
	5196: uint16(1),
	5197: uint16(aux_sym__raw_string_indented_token1),
	5198: uint16(558),
	5199: uint16(1),
	5200: uint16(anon_sym_LBRACE_LBRACE),
	5201: uint16(597),
	5202: uint16(1),
	5203: uint16(anon_sym_BQUOTE),
	5204: uint16(383),
	5205: uint16(1),
	5206: uint16(sym_command_body),
	5207: uint16(186),
	5208: uint16(2),
	5209: uint16(sym_interpolation),
	5210: uint16(aux_sym_command_body_repeat1),
	5211: uint16(3),
	5212: uint16(29),
	5213: uint16(1),
	5214: uint16(sym_comment),
	5215: uint16(601),
	5216: uint16(2),
	5217: uint16(anon_sym_LBRACK),
	5218: uint16(anon_sym_AT),
	5219: uint16(599),
	5220: uint16(3),
	5221: uint16(anon_sym_alias),
	5222: uint16(anon_sym_import),
	5223: uint16(sym_identifier),
	5224: uint16(2),
	5225: uint16(29),
	5226: uint16(1),
	5227: uint16(sym_comment),
	5228: uint16(603),
	5229: uint16(5),
	5230: uint16(anon_sym_PLUS),
	5231: uint16(anon_sym_COLON),
	5232: uint16(anon_sym_DOLLAR),
	5233: uint16(anon_sym_STAR),
	5234: uint16(sym_identifier),
	5235: uint16(3),
	5236: uint16(29),
	5237: uint16(1),
	5238: uint16(sym_comment),
	5239: uint16(607),
	5240: uint16(1),
	5241: uint16(anon_sym_COLON),
	5242: uint16(605),
	5243: uint16(4),
	5244: uint16(anon_sym_PLUS),
	5245: uint16(anon_sym_DOLLAR),
	5246: uint16(anon_sym_STAR),
	5247: uint16(sym_identifier),
	5248: uint16(6),
	5249: uint16(29),
	5250: uint16(1),
	5251: uint16(sym_comment),
	5252: uint16(67),
	5253: uint16(1),
	5254: uint16(aux_sym_string_token1),
	5255: uint16(69),
	5256: uint16(1),
	5257: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	5258: uint16(71),
	5259: uint16(1),
	5260: uint16(anon_sym_DQUOTE),
	5261: uint16(73),
	5262: uint16(1),
	5263: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5264: uint16(310),
	5265: uint16(1),
	5266: uint16(sym_string),
	5267: uint16(6),
	5268: uint16(29),
	5269: uint16(1),
	5270: uint16(sym_comment),
	5271: uint16(67),
	5272: uint16(1),
	5273: uint16(aux_sym_string_token1),
	5274: uint16(69),
	5275: uint16(1),
	5276: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	5277: uint16(71),
	5278: uint16(1),
	5279: uint16(anon_sym_DQUOTE),
	5280: uint16(73),
	5281: uint16(1),
	5282: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5283: uint16(294),
	5284: uint16(1),
	5285: uint16(sym_string),
	5286: uint16(3),
	5287: uint16(29),
	5288: uint16(1),
	5289: uint16(sym_comment),
	5290: uint16(611),
	5291: uint16(2),
	5292: uint16(anon_sym_LBRACK),
	5293: uint16(anon_sym_AT),
	5294: uint16(609),
	5295: uint16(3),
	5296: uint16(anon_sym_alias),
	5297: uint16(anon_sym_import),
	5298: uint16(sym_identifier),
	5299: uint16(3),
	5300: uint16(29),
	5301: uint16(1),
	5302: uint16(sym_comment),
	5303: uint16(613),
	5304: uint16(1),
	5305: uint16(anon_sym_COLON),
	5306: uint16(605),
	5307: uint16(4),
	5308: uint16(anon_sym_PLUS),
	5309: uint16(anon_sym_DOLLAR),
	5310: uint16(anon_sym_STAR),
	5311: uint16(sym_identifier),
	5312: uint16(5),
	5313: uint16(3),
	5314: uint16(1),
	5315: uint16(sym_comment),
	5316: uint16(615),
	5317: uint16(1),
	5318: uint16(aux_sym__raw_string_indented_token1),
	5319: uint16(618),
	5320: uint16(1),
	5321: uint16(anon_sym_BQUOTE),
	5322: uint16(620),
	5323: uint16(1),
	5324: uint16(anon_sym_LBRACE_LBRACE),
	5325: uint16(170),
	5326: uint16(2),
	5327: uint16(sym_interpolation),
	5328: uint16(aux_sym_command_body_repeat1),
	5329: uint16(5),
	5330: uint16(29),
	5331: uint16(1),
	5332: uint16(sym_comment),
	5333: uint16(275),
	5334: uint16(1),
	5335: uint16(anon_sym_LBRACE_LBRACE),
	5336: uint16(623),
	5337: uint16(1),
	5338: uint16(sym__newline),
	5339: uint16(625),
	5340: uint16(1),
	5341: uint16(sym_text),
	5342: uint16(172),
	5343: uint16(2),
	5344: uint16(sym_interpolation),
	5345: uint16(aux_sym_recipe_line_repeat1),
	5346: uint16(5),
	5347: uint16(29),
	5348: uint16(1),
	5349: uint16(sym_comment),
	5350: uint16(627),
	5351: uint16(1),
	5352: uint16(anon_sym_LBRACE_LBRACE),
	5353: uint16(630),
	5354: uint16(1),
	5355: uint16(sym__newline),
	5356: uint16(632),
	5357: uint16(1),
	5358: uint16(sym_text),
	5359: uint16(172),
	5360: uint16(2),
	5361: uint16(sym_interpolation),
	5362: uint16(aux_sym_recipe_line_repeat1),
	5363: uint16(6),
	5364: uint16(29),
	5365: uint16(1),
	5366: uint16(sym_comment),
	5367: uint16(67),
	5368: uint16(1),
	5369: uint16(aux_sym_string_token1),
	5370: uint16(69),
	5371: uint16(1),
	5372: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	5373: uint16(71),
	5374: uint16(1),
	5375: uint16(anon_sym_DQUOTE),
	5376: uint16(73),
	5377: uint16(1),
	5378: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5379: uint16(251),
	5380: uint16(1),
	5381: uint16(sym_string),
	5382: uint16(6),
	5383: uint16(29),
	5384: uint16(1),
	5385: uint16(sym_comment),
	5386: uint16(67),
	5387: uint16(1),
	5388: uint16(aux_sym_string_token1),
	5389: uint16(69),
	5390: uint16(1),
	5391: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	5392: uint16(71),
	5393: uint16(1),
	5394: uint16(anon_sym_DQUOTE),
	5395: uint16(73),
	5396: uint16(1),
	5397: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5398: uint16(154),
	5399: uint16(1),
	5400: uint16(sym_string),
	5401: uint16(6),
	5402: uint16(29),
	5403: uint16(1),
	5404: uint16(sym_comment),
	5405: uint16(635),
	5406: uint16(1),
	5407: uint16(anon_sym_COMMA),
	5408: uint16(637),
	5409: uint16(1),
	5410: uint16(anon_sym_RBRACK),
	5411: uint16(639),
	5412: uint16(1),
	5413: uint16(anon_sym_LPAREN),
	5414: uint16(641),
	5415: uint16(1),
	5416: uint16(anon_sym_COLON),
	5417: uint16(274),
	5418: uint16(1),
	5419: uint16(aux_sym_attribute_repeat2),
	5420: uint16(3),
	5421: uint16(29),
	5422: uint16(1),
	5423: uint16(sym_comment),
	5424: uint16(645),
	5425: uint16(2),
	5426: uint16(anon_sym_LBRACK),
	5427: uint16(anon_sym_AT),
	5428: uint16(643),
	5429: uint16(3),
	5430: uint16(anon_sym_alias),
	5431: uint16(anon_sym_import),
	5432: uint16(sym_identifier),
	5433: uint16(3),
	5434: uint16(29),
	5435: uint16(1),
	5436: uint16(sym_comment),
	5437: uint16(649),
	5438: uint16(2),
	5439: uint16(anon_sym_LBRACK),
	5440: uint16(anon_sym_AT),
	5441: uint16(647),
	5442: uint16(3),
	5443: uint16(anon_sym_alias),
	5444: uint16(anon_sym_import),
	5445: uint16(sym_identifier),
	5446: uint16(3),
	5447: uint16(29),
	5448: uint16(1),
	5449: uint16(sym_comment),
	5450: uint16(653),
	5451: uint16(2),
	5452: uint16(anon_sym_LBRACK),
	5453: uint16(anon_sym_AT),
	5454: uint16(651),
	5455: uint16(3),
	5456: uint16(anon_sym_alias),
	5457: uint16(anon_sym_import),
	5458: uint16(sym_identifier),
	5459: uint16(3),
	5460: uint16(29),
	5461: uint16(1),
	5462: uint16(sym_comment),
	5463: uint16(657),
	5464: uint16(2),
	5465: uint16(anon_sym_LBRACK),
	5466: uint16(anon_sym_AT),
	5467: uint16(655),
	5468: uint16(3),
	5469: uint16(anon_sym_alias),
	5470: uint16(anon_sym_import),
	5471: uint16(sym_identifier),
	5472: uint16(3),
	5473: uint16(29),
	5474: uint16(1),
	5475: uint16(sym_comment),
	5476: uint16(661),
	5477: uint16(2),
	5478: uint16(anon_sym_LBRACK),
	5479: uint16(anon_sym_AT),
	5480: uint16(659),
	5481: uint16(3),
	5482: uint16(anon_sym_alias),
	5483: uint16(anon_sym_import),
	5484: uint16(sym_identifier),
	5485: uint16(6),
	5486: uint16(29),
	5487: uint16(1),
	5488: uint16(sym_comment),
	5489: uint16(67),
	5490: uint16(1),
	5491: uint16(aux_sym_string_token1),
	5492: uint16(69),
	5493: uint16(1),
	5494: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	5495: uint16(71),
	5496: uint16(1),
	5497: uint16(anon_sym_DQUOTE),
	5498: uint16(73),
	5499: uint16(1),
	5500: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5501: uint16(110),
	5502: uint16(1),
	5503: uint16(sym_string),
	5504: uint16(5),
	5505: uint16(3),
	5506: uint16(1),
	5507: uint16(sym_comment),
	5508: uint16(564),
	5509: uint16(1),
	5510: uint16(anon_sym_LBRACE_LBRACE),
	5511: uint16(663),
	5512: uint16(1),
	5513: uint16(aux_sym__raw_string_indented_token1),
	5514: uint16(665),
	5515: uint16(1),
	5516: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	5517: uint16(184),
	5518: uint16(2),
	5519: uint16(sym_interpolation),
	5520: uint16(aux_sym_command_body_repeat1),
	5521: uint16(2),
	5522: uint16(29),
	5523: uint16(1),
	5524: uint16(sym_comment),
	5525: uint16(667),
	5526: uint16(5),
	5527: uint16(anon_sym_PLUS),
	5528: uint16(anon_sym_COLON),
	5529: uint16(anon_sym_DOLLAR),
	5530: uint16(anon_sym_STAR),
	5531: uint16(sym_identifier),
	5532: uint16(5),
	5533: uint16(3),
	5534: uint16(1),
	5535: uint16(sym_comment),
	5536: uint16(618),
	5537: uint16(1),
	5538: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	5539: uint16(669),
	5540: uint16(1),
	5541: uint16(aux_sym__raw_string_indented_token1),
	5542: uint16(672),
	5543: uint16(1),
	5544: uint16(anon_sym_LBRACE_LBRACE),
	5545: uint16(184),
	5546: uint16(2),
	5547: uint16(sym_interpolation),
	5548: uint16(aux_sym_command_body_repeat1),
	5549: uint16(3),
	5550: uint16(29),
	5551: uint16(1),
	5552: uint16(sym_comment),
	5553: uint16(677),
	5554: uint16(2),
	5555: uint16(anon_sym_LBRACK),
	5556: uint16(anon_sym_AT),
	5557: uint16(675),
	5558: uint16(3),
	5559: uint16(anon_sym_alias),
	5560: uint16(anon_sym_import),
	5561: uint16(sym_identifier),
	5562: uint16(5),
	5563: uint16(3),
	5564: uint16(1),
	5565: uint16(sym_comment),
	5566: uint16(558),
	5567: uint16(1),
	5568: uint16(anon_sym_LBRACE_LBRACE),
	5569: uint16(665),
	5570: uint16(1),
	5571: uint16(anon_sym_BQUOTE),
	5572: uint16(679),
	5573: uint16(1),
	5574: uint16(aux_sym__raw_string_indented_token1),
	5575: uint16(170),
	5576: uint16(2),
	5577: uint16(sym_interpolation),
	5578: uint16(aux_sym_command_body_repeat1),
	5579: uint16(5),
	5580: uint16(29),
	5581: uint16(1),
	5582: uint16(sym_comment),
	5583: uint16(275),
	5584: uint16(1),
	5585: uint16(anon_sym_LBRACE_LBRACE),
	5586: uint16(625),
	5587: uint16(1),
	5588: uint16(sym_text),
	5589: uint16(681),
	5590: uint16(1),
	5591: uint16(sym__newline),
	5592: uint16(172),
	5593: uint16(2),
	5594: uint16(sym_interpolation),
	5595: uint16(aux_sym_recipe_line_repeat1),
	5596: uint16(4),
	5597: uint16(3),
	5598: uint16(1),
	5599: uint16(sym_comment),
	5600: uint16(683),
	5601: uint16(1),
	5602: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5603: uint16(188),
	5604: uint16(1),
	5605: uint16(aux_sym__string_indented_repeat1),
	5606: uint16(685),
	5607: uint16(2),
	5608: uint16(aux_sym__string_indented_token1),
	5609: uint16(sym_escape_sequence),
	5610: uint16(4),
	5611: uint16(3),
	5612: uint16(1),
	5613: uint16(sym_comment),
	5614: uint16(688),
	5615: uint16(1),
	5616: uint16(anon_sym_DQUOTE),
	5617: uint16(197),
	5618: uint16(1),
	5619: uint16(aux_sym__string_repeat1),
	5620: uint16(690),
	5621: uint16(2),
	5622: uint16(aux_sym__string_token1),
	5623: uint16(sym_escape_sequence),
	5624: uint16(2),
	5625: uint16(29),
	5626: uint16(1),
	5627: uint16(sym_comment),
	5628: uint16(692),
	5629: uint16(4),
	5630: uint16(sym__newline),
	5631: uint16(anon_sym_LPAREN),
	5632: uint16(anon_sym_AMP_AMP),
	5633: uint16(sym_identifier),
	5634: uint16(5),
	5635: uint16(29),
	5636: uint16(1),
	5637: uint16(sym_comment),
	5638: uint16(694),
	5639: uint16(1),
	5640: uint16(anon_sym_COMMA),
	5641: uint16(696),
	5642: uint16(1),
	5643: uint16(anon_sym_RPAREN),
	5644: uint16(698),
	5645: uint16(1),
	5646: uint16(anon_sym_EQ),
	5647: uint16(231),
	5648: uint16(1),
	5649: uint16(aux_sym_attribute_repeat1),
	5650: uint16(2),
	5651: uint16(29),
	5652: uint16(1),
	5653: uint16(sym_comment),
	5654: uint16(241),
	5655: uint16(4),
	5656: uint16(sym__newline),
	5657: uint16(anon_sym_SLASH),
	5658: uint16(anon_sym_PLUS),
	5659: uint16(anon_sym_else),
	5660: uint16(3),
	5661: uint16(29),
	5662: uint16(1),
	5663: uint16(sym_comment),
	5664: uint16(700),
	5665: uint16(1),
	5666: uint16(anon_sym_LPAREN),
	5667: uint16(97),
	5668: uint16(3),
	5669: uint16(sym__newline),
	5670: uint16(anon_sym_SLASH),
	5671: uint16(anon_sym_PLUS),
	5672: uint16(2),
	5673: uint16(29),
	5674: uint16(1),
	5675: uint16(sym_comment),
	5676: uint16(245),
	5677: uint16(4),
	5678: uint16(sym__newline),
	5679: uint16(anon_sym_SLASH),
	5680: uint16(anon_sym_PLUS),
	5681: uint16(anon_sym_else),
	5682: uint16(2),
	5683: uint16(29),
	5684: uint16(1),
	5685: uint16(sym_comment),
	5686: uint16(249),
	5687: uint16(4),
	5688: uint16(sym__newline),
	5689: uint16(anon_sym_SLASH),
	5690: uint16(anon_sym_PLUS),
	5691: uint16(anon_sym_else),
	5692: uint16(2),
	5693: uint16(29),
	5694: uint16(1),
	5695: uint16(sym_comment),
	5696: uint16(545),
	5697: uint16(4),
	5698: uint16(sym__newline),
	5699: uint16(anon_sym_LPAREN),
	5700: uint16(anon_sym_AMP_AMP),
	5701: uint16(sym_identifier),
	5702: uint16(4),
	5703: uint16(3),
	5704: uint16(1),
	5705: uint16(sym_comment),
	5706: uint16(702),
	5707: uint16(1),
	5708: uint16(anon_sym_DQUOTE),
	5709: uint16(197),
	5710: uint16(1),
	5711: uint16(aux_sym__string_repeat1),
	5712: uint16(704),
	5713: uint16(2),
	5714: uint16(aux_sym__string_token1),
	5715: uint16(sym_escape_sequence),
	5716: uint16(4),
	5717: uint16(3),
	5718: uint16(1),
	5719: uint16(sym_comment),
	5720: uint16(688),
	5721: uint16(1),
	5722: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5723: uint16(188),
	5724: uint16(1),
	5725: uint16(aux_sym__string_indented_repeat1),
	5726: uint16(707),
	5727: uint16(2),
	5728: uint16(aux_sym__string_indented_token1),
	5729: uint16(sym_escape_sequence),
	5730: uint16(4),
	5731: uint16(29),
	5732: uint16(1),
	5733: uint16(sym_comment),
	5734: uint16(711),
	5735: uint16(1),
	5736: uint16(anon_sym_LPAREN),
	5737: uint16(713),
	5738: uint16(1),
	5739: uint16(anon_sym_COLON),
	5740: uint16(709),
	5741: uint16(2),
	5742: uint16(anon_sym_COMMA),
	5743: uint16(anon_sym_RBRACK),
	5744: uint16(5),
	5745: uint16(29),
	5746: uint16(1),
	5747: uint16(sym_comment),
	5748: uint16(694),
	5749: uint16(1),
	5750: uint16(anon_sym_COMMA),
	5751: uint16(698),
	5752: uint16(1),
	5753: uint16(anon_sym_EQ),
	5754: uint16(715),
	5755: uint16(1),
	5756: uint16(anon_sym_RPAREN),
	5757: uint16(249),
	5758: uint16(1),
	5759: uint16(aux_sym_attribute_repeat1),
	5760: uint16(4),
	5761: uint16(29),
	5762: uint16(1),
	5763: uint16(sym_comment),
	5764: uint16(717),
	5765: uint16(1),
	5766: uint16(anon_sym_LBRACE),
	5767: uint16(721),
	5768: uint16(1),
	5769: uint16(anon_sym_EQ_TILDE),
	5770: uint16(719),
	5771: uint16(2),
	5772: uint16(anon_sym_EQ_EQ),
	5773: uint16(anon_sym_BANG_EQ),
	5774: uint16(2),
	5775: uint16(29),
	5776: uint16(1),
	5777: uint16(sym_comment),
	5778: uint16(241),
	5779: uint16(4),
	5780: uint16(anon_sym_SLASH),
	5781: uint16(anon_sym_PLUS),
	5782: uint16(anon_sym_else),
	5783: uint16(anon_sym_RBRACE),
	5784: uint16(2),
	5785: uint16(29),
	5786: uint16(1),
	5787: uint16(sym_comment),
	5788: uint16(245),
	5789: uint16(4),
	5790: uint16(anon_sym_SLASH),
	5791: uint16(anon_sym_PLUS),
	5792: uint16(anon_sym_else),
	5793: uint16(anon_sym_RBRACE),
	5794: uint16(5),
	5795: uint16(29),
	5796: uint16(1),
	5797: uint16(sym_comment),
	5798: uint16(472),
	5799: uint16(1),
	5800: uint16(sym_identifier),
	5801: uint16(474),
	5802: uint16(1),
	5803: uint16(anon_sym_LPAREN),
	5804: uint16(196),
	5805: uint16(1),
	5806: uint16(sym_dependency),
	5807: uint16(212),
	5808: uint16(1),
	5809: uint16(sym_dependency_expression),
	5810: uint16(2),
	5811: uint16(29),
	5812: uint16(1),
	5813: uint16(sym_comment),
	5814: uint16(249),
	5815: uint16(4),
	5816: uint16(anon_sym_SLASH),
	5817: uint16(anon_sym_PLUS),
	5818: uint16(anon_sym_else),
	5819: uint16(anon_sym_RBRACE),
	5820: uint16(3),
	5821: uint16(29),
	5822: uint16(1),
	5823: uint16(sym_comment),
	5824: uint16(723),
	5825: uint16(1),
	5826: uint16(anon_sym_LPAREN),
	5827: uint16(97),
	5828: uint16(3),
	5829: uint16(anon_sym_SLASH),
	5830: uint16(anon_sym_PLUS),
	5831: uint16(anon_sym_RBRACE),
	5832: uint16(4),
	5833: uint16(3),
	5834: uint16(1),
	5835: uint16(sym_comment),
	5836: uint16(725),
	5837: uint16(1),
	5838: uint16(anon_sym_DQUOTE),
	5839: uint16(209),
	5840: uint16(1),
	5841: uint16(aux_sym__string_repeat1),
	5842: uint16(727),
	5843: uint16(2),
	5844: uint16(aux_sym__string_token1),
	5845: uint16(sym_escape_sequence),
	5846: uint16(4),
	5847: uint16(3),
	5848: uint16(1),
	5849: uint16(sym_comment),
	5850: uint16(725),
	5851: uint16(1),
	5852: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5853: uint16(210),
	5854: uint16(1),
	5855: uint16(aux_sym__string_indented_repeat1),
	5856: uint16(729),
	5857: uint16(2),
	5858: uint16(aux_sym__string_indented_token1),
	5859: uint16(sym_escape_sequence),
	5860: uint16(4),
	5861: uint16(3),
	5862: uint16(1),
	5863: uint16(sym_comment),
	5864: uint16(731),
	5865: uint16(1),
	5866: uint16(anon_sym_DQUOTE),
	5867: uint16(197),
	5868: uint16(1),
	5869: uint16(aux_sym__string_repeat1),
	5870: uint16(690),
	5871: uint16(2),
	5872: uint16(aux_sym__string_token1),
	5873: uint16(sym_escape_sequence),
	5874: uint16(4),
	5875: uint16(3),
	5876: uint16(1),
	5877: uint16(sym_comment),
	5878: uint16(731),
	5879: uint16(1),
	5880: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5881: uint16(188),
	5882: uint16(1),
	5883: uint16(aux_sym__string_indented_repeat1),
	5884: uint16(707),
	5885: uint16(2),
	5886: uint16(aux_sym__string_indented_token1),
	5887: uint16(sym_escape_sequence),
	5888: uint16(2),
	5889: uint16(29),
	5890: uint16(1),
	5891: uint16(sym_comment),
	5892: uint16(733),
	5893: uint16(4),
	5894: uint16(sym__newline),
	5895: uint16(anon_sym_LPAREN),
	5896: uint16(anon_sym_AMP_AMP),
	5897: uint16(sym_identifier),
	5898: uint16(2),
	5899: uint16(29),
	5900: uint16(1),
	5901: uint16(sym_comment),
	5902: uint16(735),
	5903: uint16(4),
	5904: uint16(sym__newline),
	5905: uint16(anon_sym_LPAREN),
	5906: uint16(anon_sym_AMP_AMP),
	5907: uint16(sym_identifier),
	5908: uint16(4),
	5909: uint16(3),
	5910: uint16(1),
	5911: uint16(sym_comment),
	5912: uint16(737),
	5913: uint16(1),
	5914: uint16(anon_sym_DQUOTE),
	5915: uint16(215),
	5916: uint16(1),
	5917: uint16(aux_sym__string_repeat1),
	5918: uint16(739),
	5919: uint16(2),
	5920: uint16(aux_sym__string_token1),
	5921: uint16(sym_escape_sequence),
	5922: uint16(4),
	5923: uint16(3),
	5924: uint16(1),
	5925: uint16(sym_comment),
	5926: uint16(737),
	5927: uint16(1),
	5928: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5929: uint16(216),
	5930: uint16(1),
	5931: uint16(aux_sym__string_indented_repeat1),
	5932: uint16(741),
	5933: uint16(2),
	5934: uint16(aux_sym__string_indented_token1),
	5935: uint16(sym_escape_sequence),
	5936: uint16(4),
	5937: uint16(3),
	5938: uint16(1),
	5939: uint16(sym_comment),
	5940: uint16(743),
	5941: uint16(1),
	5942: uint16(anon_sym_DQUOTE),
	5943: uint16(197),
	5944: uint16(1),
	5945: uint16(aux_sym__string_repeat1),
	5946: uint16(690),
	5947: uint16(2),
	5948: uint16(aux_sym__string_token1),
	5949: uint16(sym_escape_sequence),
	5950: uint16(4),
	5951: uint16(3),
	5952: uint16(1),
	5953: uint16(sym_comment),
	5954: uint16(743),
	5955: uint16(1),
	5956: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5957: uint16(188),
	5958: uint16(1),
	5959: uint16(aux_sym__string_indented_repeat1),
	5960: uint16(707),
	5961: uint16(2),
	5962: uint16(aux_sym__string_indented_token1),
	5963: uint16(sym_escape_sequence),
	5964: uint16(4),
	5965: uint16(3),
	5966: uint16(1),
	5967: uint16(sym_comment),
	5968: uint16(745),
	5969: uint16(1),
	5970: uint16(anon_sym_DQUOTE),
	5971: uint16(189),
	5972: uint16(1),
	5973: uint16(aux_sym__string_repeat1),
	5974: uint16(747),
	5975: uint16(2),
	5976: uint16(aux_sym__string_token1),
	5977: uint16(sym_escape_sequence),
	5978: uint16(4),
	5979: uint16(29),
	5980: uint16(1),
	5981: uint16(sym_comment),
	5982: uint16(275),
	5983: uint16(1),
	5984: uint16(anon_sym_LBRACE_LBRACE),
	5985: uint16(749),
	5986: uint16(1),
	5987: uint16(sym_text),
	5988: uint16(171),
	5989: uint16(2),
	5990: uint16(sym_interpolation),
	5991: uint16(aux_sym_recipe_line_repeat1),
	5992: uint16(4),
	5993: uint16(3),
	5994: uint16(1),
	5995: uint16(sym_comment),
	5996: uint16(745),
	5997: uint16(1),
	5998: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5999: uint16(198),
	6000: uint16(1),
	6001: uint16(aux_sym__string_indented_repeat1),
	6002: uint16(751),
	6003: uint16(2),
	6004: uint16(aux_sym__string_indented_token1),
	6005: uint16(sym_escape_sequence),
	6006: uint16(2),
	6007: uint16(29),
	6008: uint16(1),
	6009: uint16(sym_comment),
	6010: uint16(753),
	6011: uint16(4),
	6012: uint16(sym__newline),
	6013: uint16(anon_sym_LPAREN),
	6014: uint16(anon_sym_AMP_AMP),
	6015: uint16(sym_identifier),
	6016: uint16(2),
	6017: uint16(3),
	6018: uint16(1),
	6019: uint16(sym_comment),
	6020: uint16(755),
	6021: uint16(3),
	6022: uint16(aux_sym__raw_string_indented_token1),
	6023: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	6024: uint16(anon_sym_LBRACE_LBRACE),
	6025: uint16(4),
	6026: uint16(29),
	6027: uint16(1),
	6028: uint16(sym_comment),
	6029: uint16(757),
	6030: uint16(1),
	6031: uint16(anon_sym_COMMA),
	6032: uint16(760),
	6033: uint16(1),
	6034: uint16(anon_sym_RPAREN),
	6035: uint16(222),
	6036: uint16(1),
	6037: uint16(aux_sym_attribute_repeat1),
	6038: uint16(3),
	6039: uint16(29),
	6040: uint16(1),
	6041: uint16(sym_comment),
	6042: uint16(762),
	6043: uint16(1),
	6044: uint16(anon_sym_LBRACE),
	6045: uint16(97),
	6046: uint16(2),
	6047: uint16(anon_sym_SLASH),
	6048: uint16(anon_sym_PLUS),
	6049: uint16(4),
	6050: uint16(29),
	6051: uint16(1),
	6052: uint16(sym_comment),
	6053: uint16(764),
	6054: uint16(1),
	6055: uint16(anon_sym_if),
	6056: uint16(766),
	6057: uint16(1),
	6058: uint16(anon_sym_LBRACE),
	6059: uint16(235),
	6060: uint16(1),
	6061: uint16(sym__braced_expr),
	6062: uint16(2),
	6063: uint16(29),
	6064: uint16(1),
	6065: uint16(sym_comment),
	6066: uint16(261),
	6067: uint16(3),
	6068: uint16(sym__newline),
	6069: uint16(anon_sym_SLASH),
	6070: uint16(anon_sym_PLUS),
	6071: uint16(4),
	6072: uint16(29),
	6073: uint16(1),
	6074: uint16(sym_comment),
	6075: uint16(768),
	6076: uint16(1),
	6077: uint16(sym_identifier),
	6078: uint16(770),
	6079: uint16(1),
	6080: uint16(aux_sym__shebang_with_lang_token2),
	6081: uint16(293),
	6082: uint16(1),
	6083: uint16(aux_sym__shebang_with_lang_repeat1),
	6084: uint16(2),
	6085: uint16(3),
	6086: uint16(1),
	6087: uint16(sym_comment),
	6088: uint16(755),
	6089: uint16(3),
	6090: uint16(aux_sym__raw_string_indented_token1),
	6091: uint16(anon_sym_BQUOTE),
	6092: uint16(anon_sym_LBRACE_LBRACE),
	6093: uint16(4),
	6094: uint16(29),
	6095: uint16(1),
	6096: uint16(sym_comment),
	6097: uint16(772),
	6098: uint16(1),
	6099: uint16(anon_sym_COMMA),
	6100: uint16(774),
	6101: uint16(1),
	6102: uint16(anon_sym_RPAREN),
	6103: uint16(237),
	6104: uint16(1),
	6105: uint16(aux_sym_sequence_repeat1),
	6106: uint16(2),
	6107: uint16(29),
	6108: uint16(1),
	6109: uint16(sym_comment),
	6110: uint16(158),
	6111: uint16(3),
	6112: uint16(sym__newline),
	6113: uint16(anon_sym_SLASH),
	6114: uint16(anon_sym_PLUS),
	6115: uint16(3),
	6116: uint16(29),
	6117: uint16(1),
	6118: uint16(sym_comment),
	6119: uint16(776),
	6120: uint16(1),
	6121: uint16(anon_sym_PLUS),
	6122: uint16(265),
	6123: uint16(2),
	6124: uint16(sym__newline),
	6125: uint16(anon_sym_SLASH),
	6126: uint16(4),
	6127: uint16(29),
	6128: uint16(1),
	6129: uint16(sym_comment),
	6130: uint16(694),
	6131: uint16(1),
	6132: uint16(anon_sym_COMMA),
	6133: uint16(778),
	6134: uint16(1),
	6135: uint16(anon_sym_RPAREN),
	6136: uint16(222),
	6137: uint16(1),
	6138: uint16(aux_sym_attribute_repeat1),
	6139: uint16(4),
	6140: uint16(29),
	6141: uint16(1),
	6142: uint16(sym_comment),
	6143: uint16(635),
	6144: uint16(1),
	6145: uint16(anon_sym_COMMA),
	6146: uint16(780),
	6147: uint16(1),
	6148: uint16(anon_sym_RBRACK),
	6149: uint16(252),
	6150: uint16(1),
	6151: uint16(aux_sym_attribute_repeat2),
	6152: uint16(2),
	6153: uint16(29),
	6154: uint16(1),
	6155: uint16(sym_comment),
	6156: uint16(265),
	6157: uint16(3),
	6158: uint16(sym__newline),
	6159: uint16(anon_sym_SLASH),
	6160: uint16(anon_sym_PLUS),
	6161: uint16(2),
	6162: uint16(29),
	6163: uint16(1),
	6164: uint16(sym_comment),
	6165: uint16(97),
	6166: uint16(3),
	6167: uint16(sym__newline),
	6168: uint16(anon_sym_SLASH),
	6169: uint16(anon_sym_PLUS),
	6170: uint16(2),
	6171: uint16(29),
	6172: uint16(1),
	6173: uint16(sym_comment),
	6174: uint16(257),
	6175: uint16(3),
	6176: uint16(sym__newline),
	6177: uint16(anon_sym_SLASH),
	6178: uint16(anon_sym_PLUS),
	6179: uint16(2),
	6180: uint16(29),
	6181: uint16(1),
	6182: uint16(sym_comment),
	6183: uint16(253),
	6184: uint16(3),
	6185: uint16(sym__newline),
	6186: uint16(anon_sym_SLASH),
	6187: uint16(anon_sym_PLUS),
	6188: uint16(4),
	6189: uint16(29),
	6190: uint16(1),
	6191: uint16(sym_comment),
	6192: uint16(782),
	6193: uint16(1),
	6194: uint16(anon_sym_COMMA),
	6195: uint16(785),
	6196: uint16(1),
	6197: uint16(anon_sym_RPAREN),
	6198: uint16(237),
	6199: uint16(1),
	6200: uint16(aux_sym_sequence_repeat1),
	6201: uint16(4),
	6202: uint16(29),
	6203: uint16(1),
	6204: uint16(sym_comment),
	6205: uint16(336),
	6206: uint16(1),
	6207: uint16(sym__newline),
	6208: uint16(776),
	6209: uint16(1),
	6210: uint16(anon_sym_PLUS),
	6211: uint16(787),
	6212: uint16(1),
	6213: uint16(anon_sym_SLASH),
	6214: uint16(2),
	6215: uint16(29),
	6216: uint16(1),
	6217: uint16(sym_comment),
	6218: uint16(269),
	6219: uint16(3),
	6220: uint16(sym__newline),
	6221: uint16(anon_sym_SLASH),
	6222: uint16(anon_sym_PLUS),
	6223: uint16(2),
	6224: uint16(29),
	6225: uint16(1),
	6226: uint16(sym_comment),
	6227: uint16(51),
	6228: uint16(3),
	6229: uint16(sym__newline),
	6230: uint16(anon_sym_SLASH),
	6231: uint16(anon_sym_PLUS),
	6232: uint16(4),
	6233: uint16(29),
	6234: uint16(1),
	6235: uint16(sym_comment),
	6236: uint16(789),
	6237: uint16(1),
	6238: uint16(sym_identifier),
	6239: uint16(791),
	6240: uint16(1),
	6241: uint16(aux_sym__shebang_with_lang_token2),
	6242: uint16(241),
	6243: uint16(1),
	6244: uint16(aux_sym__shebang_with_lang_repeat1),
	6245: uint16(2),
	6246: uint16(29),
	6247: uint16(1),
	6248: uint16(sym_comment),
	6249: uint16(25),
	6250: uint16(3),
	6251: uint16(sym__newline),
	6252: uint16(anon_sym_SLASH),
	6253: uint16(anon_sym_PLUS),
	6254: uint16(2),
	6255: uint16(29),
	6256: uint16(1),
	6257: uint16(sym_comment),
	6258: uint16(31),
	6259: uint16(3),
	6260: uint16(sym__newline),
	6261: uint16(anon_sym_SLASH),
	6262: uint16(anon_sym_PLUS),
	6263: uint16(4),
	6264: uint16(3),
	6265: uint16(1),
	6266: uint16(sym_comment),
	6267: uint16(794),
	6268: uint16(1),
	6269: uint16(aux_sym__shebang_with_lang_token1),
	6270: uint16(796),
	6271: uint16(1),
	6272: uint16(sym__opaque_shebang),
	6273: uint16(389),
	6274: uint16(1),
	6275: uint16(sym__shebang_with_lang),
	6276: uint16(4),
	6277: uint16(29),
	6278: uint16(1),
	6279: uint16(sym_comment),
	6280: uint16(798),
	6281: uint16(1),
	6282: uint16(sym_identifier),
	6283: uint16(800),
	6284: uint16(1),
	6285: uint16(anon_sym_DOLLAR),
	6286: uint16(381),
	6287: uint16(1),
	6288: uint16(sym_parameter),
	6289: uint16(2),
	6290: uint16(29),
	6291: uint16(1),
	6292: uint16(sym_comment),
	6293: uint16(35),
	6294: uint16(3),
	6295: uint16(sym__newline),
	6296: uint16(anon_sym_SLASH),
	6297: uint16(anon_sym_PLUS),
	6298: uint16(4),
	6299: uint16(3),
	6300: uint16(1),
	6301: uint16(sym_comment),
	6302: uint16(802),
	6303: uint16(1),
	6304: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	6305: uint16(804),
	6306: uint16(1),
	6307: uint16(aux_sym__raw_string_indented_token1),
	6308: uint16(247),
	6309: uint16(1),
	6310: uint16(aux_sym__raw_string_indented_repeat1),
	6311: uint16(2),
	6312: uint16(29),
	6313: uint16(1),
	6314: uint16(sym_comment),
	6315: uint16(55),
	6316: uint16(3),
	6317: uint16(sym__newline),
	6318: uint16(anon_sym_SLASH),
	6319: uint16(anon_sym_PLUS),
	6320: uint16(4),
	6321: uint16(29),
	6322: uint16(1),
	6323: uint16(sym_comment),
	6324: uint16(694),
	6325: uint16(1),
	6326: uint16(anon_sym_COMMA),
	6327: uint16(807),
	6328: uint16(1),
	6329: uint16(anon_sym_RPAREN),
	6330: uint16(222),
	6331: uint16(1),
	6332: uint16(aux_sym_attribute_repeat1),
	6333: uint16(4),
	6334: uint16(29),
	6335: uint16(1),
	6336: uint16(sym_comment),
	6337: uint16(694),
	6338: uint16(1),
	6339: uint16(anon_sym_COMMA),
	6340: uint16(715),
	6341: uint16(1),
	6342: uint16(anon_sym_RPAREN),
	6343: uint16(249),
	6344: uint16(1),
	6345: uint16(aux_sym_attribute_repeat1),
	6346: uint16(4),
	6347: uint16(29),
	6348: uint16(1),
	6349: uint16(sym_comment),
	6350: uint16(635),
	6351: uint16(1),
	6352: uint16(anon_sym_COMMA),
	6353: uint16(809),
	6354: uint16(1),
	6355: uint16(anon_sym_RBRACK),
	6356: uint16(275),
	6357: uint16(1),
	6358: uint16(aux_sym_attribute_repeat2),
	6359: uint16(4),
	6360: uint16(29),
	6361: uint16(1),
	6362: uint16(sym_comment),
	6363: uint16(811),
	6364: uint16(1),
	6365: uint16(anon_sym_COMMA),
	6366: uint16(814),
	6367: uint16(1),
	6368: uint16(anon_sym_RBRACK),
	6369: uint16(252),
	6370: uint16(1),
	6371: uint16(aux_sym_attribute_repeat2),
	6372: uint16(2),
	6373: uint16(29),
	6374: uint16(1),
	6375: uint16(sym_comment),
	6376: uint16(148),
	6377: uint16(3),
	6378: uint16(sym__newline),
	6379: uint16(anon_sym_SLASH),
	6380: uint16(anon_sym_PLUS),
	6381: uint16(2),
	6382: uint16(29),
	6383: uint16(1),
	6384: uint16(sym_comment),
	6385: uint16(97),
	6386: uint16(3),
	6387: uint16(anon_sym_SLASH),
	6388: uint16(anon_sym_PLUS),
	6389: uint16(anon_sym_RBRACE),
	6390: uint16(4),
	6391: uint16(29),
	6392: uint16(1),
	6393: uint16(sym_comment),
	6394: uint16(336),
	6395: uint16(1),
	6396: uint16(anon_sym_RBRACE),
	6397: uint16(816),
	6398: uint16(1),
	6399: uint16(anon_sym_SLASH),
	6400: uint16(818),
	6401: uint16(1),
	6402: uint16(anon_sym_PLUS),
	6403: uint16(2),
	6404: uint16(29),
	6405: uint16(1),
	6406: uint16(sym_comment),
	6407: uint16(269),
	6408: uint16(3),
	6409: uint16(anon_sym_SLASH),
	6410: uint16(anon_sym_PLUS),
	6411: uint16(anon_sym_RBRACE),
	6412: uint16(2),
	6413: uint16(29),
	6414: uint16(1),
	6415: uint16(sym_comment),
	6416: uint16(51),
	6417: uint16(3),
	6418: uint16(anon_sym_SLASH),
	6419: uint16(anon_sym_PLUS),
	6420: uint16(anon_sym_RBRACE),
	6421: uint16(4),
	6422: uint16(29),
	6423: uint16(1),
	6424: uint16(sym_comment),
	6425: uint16(311),
	6426: uint16(1),
	6427: uint16(anon_sym_RBRACE),
	6428: uint16(816),
	6429: uint16(1),
	6430: uint16(anon_sym_SLASH),
	6431: uint16(818),
	6432: uint16(1),
	6433: uint16(anon_sym_PLUS),
	6434: uint16(2),
	6435: uint16(29),
	6436: uint16(1),
	6437: uint16(sym_comment),
	6438: uint16(136),
	6439: uint16(3),
	6440: uint16(anon_sym_SLASH),
	6441: uint16(anon_sym_PLUS),
	6442: uint16(anon_sym_RBRACE),
	6443: uint16(2),
	6444: uint16(29),
	6445: uint16(1),
	6446: uint16(sym_comment),
	6447: uint16(140),
	6448: uint16(3),
	6449: uint16(anon_sym_SLASH),
	6450: uint16(anon_sym_PLUS),
	6451: uint16(anon_sym_RBRACE),
	6452: uint16(4),
	6453: uint16(29),
	6454: uint16(1),
	6455: uint16(sym_comment),
	6456: uint16(311),
	6457: uint16(1),
	6458: uint16(sym__newline),
	6459: uint16(776),
	6460: uint16(1),
	6461: uint16(anon_sym_PLUS),
	6462: uint16(787),
	6463: uint16(1),
	6464: uint16(anon_sym_SLASH),
	6465: uint16(2),
	6466: uint16(29),
	6467: uint16(1),
	6468: uint16(sym_comment),
	6469: uint16(55),
	6470: uint16(3),
	6471: uint16(anon_sym_SLASH),
	6472: uint16(anon_sym_PLUS),
	6473: uint16(anon_sym_RBRACE),
	6474: uint16(2),
	6475: uint16(29),
	6476: uint16(1),
	6477: uint16(sym_comment),
	6478: uint16(144),
	6479: uint16(3),
	6480: uint16(anon_sym_SLASH),
	6481: uint16(anon_sym_PLUS),
	6482: uint16(anon_sym_RBRACE),
	6483: uint16(2),
	6484: uint16(29),
	6485: uint16(1),
	6486: uint16(sym_comment),
	6487: uint16(148),
	6488: uint16(3),
	6489: uint16(anon_sym_SLASH),
	6490: uint16(anon_sym_PLUS),
	6491: uint16(anon_sym_RBRACE),
	6492: uint16(2),
	6493: uint16(29),
	6494: uint16(1),
	6495: uint16(sym_comment),
	6496: uint16(152),
	6497: uint16(3),
	6498: uint16(anon_sym_SLASH),
	6499: uint16(anon_sym_PLUS),
	6500: uint16(anon_sym_RBRACE),
	6501: uint16(2),
	6502: uint16(29),
	6503: uint16(1),
	6504: uint16(sym_comment),
	6505: uint16(265),
	6506: uint16(3),
	6507: uint16(anon_sym_SLASH),
	6508: uint16(anon_sym_PLUS),
	6509: uint16(anon_sym_RBRACE),
	6510: uint16(2),
	6511: uint16(29),
	6512: uint16(1),
	6513: uint16(sym_comment),
	6514: uint16(152),
	6515: uint16(3),
	6516: uint16(sym__newline),
	6517: uint16(anon_sym_SLASH),
	6518: uint16(anon_sym_PLUS),
	6519: uint16(2),
	6520: uint16(29),
	6521: uint16(1),
	6522: uint16(sym_comment),
	6523: uint16(261),
	6524: uint16(3),
	6525: uint16(anon_sym_SLASH),
	6526: uint16(anon_sym_PLUS),
	6527: uint16(anon_sym_RBRACE),
	6528: uint16(2),
	6529: uint16(29),
	6530: uint16(1),
	6531: uint16(sym_comment),
	6532: uint16(136),
	6533: uint16(3),
	6534: uint16(sym__newline),
	6535: uint16(anon_sym_SLASH),
	6536: uint16(anon_sym_PLUS),
	6537: uint16(2),
	6538: uint16(29),
	6539: uint16(1),
	6540: uint16(sym_comment),
	6541: uint16(820),
	6542: uint16(3),
	6543: uint16(sym__newline),
	6544: uint16(sym_text),
	6545: uint16(anon_sym_LBRACE_LBRACE),
	6546: uint16(2),
	6547: uint16(29),
	6548: uint16(1),
	6549: uint16(sym_comment),
	6550: uint16(158),
	6551: uint16(3),
	6552: uint16(anon_sym_SLASH),
	6553: uint16(anon_sym_PLUS),
	6554: uint16(anon_sym_RBRACE),
	6555: uint16(2),
	6556: uint16(29),
	6557: uint16(1),
	6558: uint16(sym_comment),
	6559: uint16(257),
	6560: uint16(3),
	6561: uint16(anon_sym_SLASH),
	6562: uint16(anon_sym_PLUS),
	6563: uint16(anon_sym_RBRACE),
	6564: uint16(2),
	6565: uint16(29),
	6566: uint16(1),
	6567: uint16(sym_comment),
	6568: uint16(253),
	6569: uint16(3),
	6570: uint16(anon_sym_SLASH),
	6571: uint16(anon_sym_PLUS),
	6572: uint16(anon_sym_RBRACE),
	6573: uint16(4),
	6574: uint16(29),
	6575: uint16(1),
	6576: uint16(sym_comment),
	6577: uint16(635),
	6578: uint16(1),
	6579: uint16(anon_sym_COMMA),
	6580: uint16(822),
	6581: uint16(1),
	6582: uint16(anon_sym_RBRACK),
	6583: uint16(252),
	6584: uint16(1),
	6585: uint16(aux_sym_attribute_repeat2),
	6586: uint16(4),
	6587: uint16(29),
	6588: uint16(1),
	6589: uint16(sym_comment),
	6590: uint16(635),
	6591: uint16(1),
	6592: uint16(anon_sym_COMMA),
	6593: uint16(824),
	6594: uint16(1),
	6595: uint16(anon_sym_RBRACK),
	6596: uint16(252),
	6597: uint16(1),
	6598: uint16(aux_sym_attribute_repeat2),
	6599: uint16(3),
	6600: uint16(29),
	6601: uint16(1),
	6602: uint16(sym_comment),
	6603: uint16(818),
	6604: uint16(1),
	6605: uint16(anon_sym_PLUS),
	6606: uint16(265),
	6607: uint16(2),
	6608: uint16(anon_sym_SLASH),
	6609: uint16(anon_sym_RBRACE),
	6610: uint16(4),
	6611: uint16(3),
	6612: uint16(1),
	6613: uint16(sym_comment),
	6614: uint16(725),
	6615: uint16(1),
	6616: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	6617: uint16(826),
	6618: uint16(1),
	6619: uint16(aux_sym__raw_string_indented_token1),
	6620: uint16(278),
	6621: uint16(1),
	6622: uint16(aux_sym__raw_string_indented_repeat1),
	6623: uint16(4),
	6624: uint16(3),
	6625: uint16(1),
	6626: uint16(sym_comment),
	6627: uint16(731),
	6628: uint16(1),
	6629: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	6630: uint16(828),
	6631: uint16(1),
	6632: uint16(aux_sym__raw_string_indented_token1),
	6633: uint16(247),
	6634: uint16(1),
	6635: uint16(aux_sym__raw_string_indented_repeat1),
	6636: uint16(2),
	6637: uint16(29),
	6638: uint16(1),
	6639: uint16(sym_comment),
	6640: uint16(140),
	6641: uint16(3),
	6642: uint16(sym__newline),
	6643: uint16(anon_sym_SLASH),
	6644: uint16(anon_sym_PLUS),
	6645: uint16(4),
	6646: uint16(29),
	6647: uint16(1),
	6648: uint16(sym_comment),
	6649: uint16(830),
	6650: uint16(1),
	6651: uint16(anon_sym_if),
	6652: uint16(832),
	6653: uint16(1),
	6654: uint16(anon_sym_LBRACE),
	6655: uint16(50),
	6656: uint16(1),
	6657: uint16(sym__braced_expr),
	6658: uint16(4),
	6659: uint16(3),
	6660: uint16(1),
	6661: uint16(sym_comment),
	6662: uint16(745),
	6663: uint16(1),
	6664: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	6665: uint16(834),
	6666: uint16(1),
	6667: uint16(aux_sym__raw_string_indented_token1),
	6668: uint16(290),
	6669: uint16(1),
	6670: uint16(aux_sym__raw_string_indented_repeat1),
	6671: uint16(4),
	6672: uint16(3),
	6673: uint16(1),
	6674: uint16(sym_comment),
	6675: uint16(737),
	6676: uint16(1),
	6677: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	6678: uint16(836),
	6679: uint16(1),
	6680: uint16(aux_sym__raw_string_indented_token1),
	6681: uint16(285),
	6682: uint16(1),
	6683: uint16(aux_sym__raw_string_indented_repeat1),
	6684: uint16(4),
	6685: uint16(29),
	6686: uint16(1),
	6687: uint16(sym_comment),
	6688: uint16(772),
	6689: uint16(1),
	6690: uint16(anon_sym_COMMA),
	6691: uint16(838),
	6692: uint16(1),
	6693: uint16(anon_sym_RPAREN),
	6694: uint16(228),
	6695: uint16(1),
	6696: uint16(aux_sym_sequence_repeat1),
	6697: uint16(4),
	6698: uint16(29),
	6699: uint16(1),
	6700: uint16(sym_comment),
	6701: uint16(694),
	6702: uint16(1),
	6703: uint16(anon_sym_COMMA),
	6704: uint16(696),
	6705: uint16(1),
	6706: uint16(anon_sym_RPAREN),
	6707: uint16(231),
	6708: uint16(1),
	6709: uint16(aux_sym_attribute_repeat1),
	6710: uint16(4),
	6711: uint16(3),
	6712: uint16(1),
	6713: uint16(sym_comment),
	6714: uint16(743),
	6715: uint16(1),
	6716: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	6717: uint16(828),
	6718: uint16(1),
	6719: uint16(aux_sym__raw_string_indented_token1),
	6720: uint16(247),
	6721: uint16(1),
	6722: uint16(aux_sym__raw_string_indented_repeat1),
	6723: uint16(3),
	6724: uint16(29),
	6725: uint16(1),
	6726: uint16(sym_comment),
	6727: uint16(698),
	6728: uint16(1),
	6729: uint16(anon_sym_EQ),
	6730: uint16(760),
	6731: uint16(2),
	6732: uint16(anon_sym_COMMA),
	6733: uint16(anon_sym_RPAREN),
	6734: uint16(4),
	6735: uint16(29),
	6736: uint16(1),
	6737: uint16(sym_comment),
	6738: uint16(840),
	6739: uint16(1),
	6740: uint16(anon_sym_if),
	6741: uint16(842),
	6742: uint16(1),
	6743: uint16(anon_sym_LBRACE),
	6744: uint16(272),
	6745: uint16(1),
	6746: uint16(sym__braced_expr),
	6747: uint16(4),
	6748: uint16(29),
	6749: uint16(1),
	6750: uint16(sym_comment),
	6751: uint16(635),
	6752: uint16(1),
	6753: uint16(anon_sym_COMMA),
	6754: uint16(844),
	6755: uint16(1),
	6756: uint16(anon_sym_RBRACK),
	6757: uint16(252),
	6758: uint16(1),
	6759: uint16(aux_sym_attribute_repeat2),
	6760: uint16(2),
	6761: uint16(29),
	6762: uint16(1),
	6763: uint16(sym_comment),
	6764: uint16(144),
	6765: uint16(3),
	6766: uint16(sym__newline),
	6767: uint16(anon_sym_SLASH),
	6768: uint16(anon_sym_PLUS),
	6769: uint16(4),
	6770: uint16(3),
	6771: uint16(1),
	6772: uint16(sym_comment),
	6773: uint16(688),
	6774: uint16(1),
	6775: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	6776: uint16(828),
	6777: uint16(1),
	6778: uint16(aux_sym__raw_string_indented_token1),
	6779: uint16(247),
	6780: uint16(1),
	6781: uint16(aux_sym__raw_string_indented_repeat1),
	6782: uint16(4),
	6783: uint16(29),
	6784: uint16(1),
	6785: uint16(sym_comment),
	6786: uint16(635),
	6787: uint16(1),
	6788: uint16(anon_sym_COMMA),
	6789: uint16(846),
	6790: uint16(1),
	6791: uint16(anon_sym_RBRACK),
	6792: uint16(232),
	6793: uint16(1),
	6794: uint16(aux_sym_attribute_repeat2),
	6795: uint16(4),
	6796: uint16(29),
	6797: uint16(1),
	6798: uint16(sym_comment),
	6799: uint16(635),
	6800: uint16(1),
	6801: uint16(anon_sym_COMMA),
	6802: uint16(848),
	6803: uint16(1),
	6804: uint16(anon_sym_RBRACK),
	6805: uint16(288),
	6806: uint16(1),
	6807: uint16(aux_sym_attribute_repeat2),
	6808: uint16(4),
	6809: uint16(29),
	6810: uint16(1),
	6811: uint16(sym_comment),
	6812: uint16(850),
	6813: uint16(1),
	6814: uint16(sym_identifier),
	6815: uint16(852),
	6816: uint16(1),
	6817: uint16(aux_sym__shebang_with_lang_token2),
	6818: uint16(241),
	6819: uint16(1),
	6820: uint16(aux_sym__shebang_with_lang_repeat1),
	6821: uint16(2),
	6822: uint16(29),
	6823: uint16(1),
	6824: uint16(sym_comment),
	6825: uint16(854),
	6826: uint16(2),
	6827: uint16(anon_sym_COMMA),
	6828: uint16(anon_sym_RPAREN),
	6829: uint16(3),
	6830: uint16(29),
	6831: uint16(1),
	6832: uint16(sym_comment),
	6833: uint16(856),
	6834: uint16(1),
	6835: uint16(sym_identifier),
	6836: uint16(89),
	6837: uint16(1),
	6838: uint16(sym_assignment),
	6839: uint16(3),
	6840: uint16(29),
	6841: uint16(1),
	6842: uint16(sym_comment),
	6843: uint16(858),
	6844: uint16(1),
	6845: uint16(sym_identifier),
	6846: uint16(860),
	6847: uint16(1),
	6848: uint16(anon_sym_import),
	6849: uint16(2),
	6850: uint16(29),
	6851: uint16(1),
	6852: uint16(sym_comment),
	6853: uint16(862),
	6854: uint16(2),
	6855: uint16(anon_sym_COMMA),
	6856: uint16(anon_sym_RBRACK),
	6857: uint16(2),
	6858: uint16(29),
	6859: uint16(1),
	6860: uint16(sym_comment),
	6861: uint16(785),
	6862: uint16(2),
	6863: uint16(anon_sym_COMMA),
	6864: uint16(anon_sym_RPAREN),
	6865: uint16(3),
	6866: uint16(29),
	6867: uint16(1),
	6868: uint16(sym_comment),
	6869: uint16(766),
	6870: uint16(1),
	6871: uint16(anon_sym_LBRACE),
	6872: uint16(146),
	6873: uint16(1),
	6874: uint16(sym__braced_expr),
	6875: uint16(3),
	6876: uint16(29),
	6877: uint16(1),
	6878: uint16(sym_comment),
	6879: uint16(864),
	6880: uint16(1),
	6881: uint16(sym_identifier),
	6882: uint16(866),
	6883: uint16(1),
	6884: uint16(anon_sym_env),
	6885: uint16(3),
	6886: uint16(29),
	6887: uint16(1),
	6888: uint16(sym_comment),
	6889: uint16(868),
	6890: uint16(1),
	6891: uint16(sym_identifier),
	6892: uint16(870),
	6893: uint16(1),
	6894: uint16(anon_sym_QMARK),
	6895: uint16(2),
	6896: uint16(29),
	6897: uint16(1),
	6898: uint16(sym_comment),
	6899: uint16(872),
	6900: uint16(2),
	6901: uint16(anon_sym_COMMA),
	6902: uint16(anon_sym_RBRACK),
	6903: uint16(3),
	6904: uint16(29),
	6905: uint16(1),
	6906: uint16(sym_comment),
	6907: uint16(766),
	6908: uint16(1),
	6909: uint16(anon_sym_LBRACE),
	6910: uint16(195),
	6911: uint16(1),
	6912: uint16(sym__braced_expr),
	6913: uint16(3),
	6914: uint16(29),
	6915: uint16(1),
	6916: uint16(sym_comment),
	6917: uint16(593),
	6918: uint16(1),
	6919: uint16(anon_sym_COLON),
	6920: uint16(874),
	6921: uint16(1),
	6922: uint16(anon_sym_EQ),
	6923: uint16(3),
	6924: uint16(29),
	6925: uint16(1),
	6926: uint16(sym_comment),
	6927: uint16(583),
	6928: uint16(1),
	6929: uint16(anon_sym_COLON),
	6930: uint16(876),
	6931: uint16(1),
	6932: uint16(anon_sym_EQ),
	6933: uint16(3),
	6934: uint16(29),
	6935: uint16(1),
	6936: uint16(sym_comment),
	6937: uint16(832),
	6938: uint16(1),
	6939: uint16(anon_sym_LBRACE),
	6940: uint16(5),
	6941: uint16(1),
	6942: uint16(sym__braced_expr),
	6943: uint16(3),
	6944: uint16(29),
	6945: uint16(1),
	6946: uint16(sym_comment),
	6947: uint16(832),
	6948: uint16(1),
	6949: uint16(anon_sym_LBRACE),
	6950: uint16(48),
	6951: uint16(1),
	6952: uint16(sym__braced_expr),
	6953: uint16(3),
	6954: uint16(29),
	6955: uint16(1),
	6956: uint16(sym_comment),
	6957: uint16(878),
	6958: uint16(1),
	6959: uint16(anon_sym_COLON_EQ),
	6960: uint16(880),
	6961: uint16(1),
	6962: uint16(sym__newline),
	6963: uint16(3),
	6964: uint16(29),
	6965: uint16(1),
	6966: uint16(sym_comment),
	6967: uint16(882),
	6968: uint16(1),
	6969: uint16(sym_identifier),
	6970: uint16(884),
	6971: uint16(1),
	6972: uint16(anon_sym_shell),
	6973: uint16(2),
	6974: uint16(29),
	6975: uint16(1),
	6976: uint16(sym_comment),
	6977: uint16(886),
	6978: uint16(2),
	6979: uint16(anon_sym_COMMA),
	6980: uint16(anon_sym_RBRACK),
	6981: uint16(2),
	6982: uint16(29),
	6983: uint16(1),
	6984: uint16(sym_comment),
	6985: uint16(760),
	6986: uint16(2),
	6987: uint16(anon_sym_COMMA),
	6988: uint16(anon_sym_RPAREN),
	6989: uint16(2),
	6990: uint16(29),
	6991: uint16(1),
	6992: uint16(sym_comment),
	6993: uint16(888),
	6994: uint16(2),
	6995: uint16(sym_text),
	6996: uint16(anon_sym_LBRACE_LBRACE),
	6997: uint16(3),
	6998: uint16(29),
	6999: uint16(1),
	7000: uint16(sym_comment),
	7001: uint16(842),
	7002: uint16(1),
	7003: uint16(anon_sym_LBRACE),
	7004: uint16(142),
	7005: uint16(1),
	7006: uint16(sym__braced_expr),
	7007: uint16(3),
	7008: uint16(29),
	7009: uint16(1),
	7010: uint16(sym_comment),
	7011: uint16(842),
	7012: uint16(1),
	7013: uint16(anon_sym_LBRACE),
	7014: uint16(205),
	7015: uint16(1),
	7016: uint16(sym__braced_expr),
	7017: uint16(2),
	7018: uint16(29),
	7019: uint16(1),
	7020: uint16(sym_comment),
	7021: uint16(890),
	7022: uint16(1),
	7023: uint16(anon_sym_COLON),
	7024: uint16(2),
	7025: uint16(3),
	7026: uint16(1),
	7027: uint16(sym_comment),
	7028: uint16(892),
	7029: uint16(1),
	7030: uint16(aux_sym__shebang_with_lang_token3),
	7031: uint16(2),
	7032: uint16(29),
	7033: uint16(1),
	7034: uint16(sym_comment),
	7035: uint16(894),
	7036: uint16(1),
	7037: uint16(sym__newline),
	7038: uint16(2),
	7039: uint16(29),
	7040: uint16(1),
	7041: uint16(sym_comment),
	7042: uint16(896),
	7043: uint16(1),
	7044: uint16(anon_sym_LBRACE),
	7045: uint16(2),
	7046: uint16(29),
	7047: uint16(1),
	7048: uint16(sym_comment),
	7049: uint16(898),
	7050: uint16(1),
	7051: uint16(sym__newline),
	7052: uint16(2),
	7053: uint16(29),
	7054: uint16(1),
	7055: uint16(sym_comment),
	7056: uint16(900),
	7057: uint16(1),
	7058: uint16(anon_sym_COLON),
	7059: uint16(2),
	7060: uint16(3),
	7061: uint16(1),
	7062: uint16(sym_comment),
	7063: uint16(902),
	7064: uint16(1),
	7065: uint16(aux_sym__shebang_with_lang_token3),
	7066: uint16(2),
	7067: uint16(29),
	7068: uint16(1),
	7069: uint16(sym_comment),
	7070: uint16(904),
	7071: uint16(1),
	7072: uint16(sym__newline),
	7073: uint16(2),
	7074: uint16(29),
	7075: uint16(1),
	7076: uint16(sym_comment),
	7077: uint16(906),
	7078: uint16(1),
	7079: uint16(anon_sym_RBRACE),
	7080: uint16(2),
	7081: uint16(29),
	7082: uint16(1),
	7083: uint16(sym_comment),
	7084: uint16(908),
	7085: uint16(1),
	7086: uint16(sym__newline),
	7087: uint16(2),
	7088: uint16(29),
	7089: uint16(1),
	7090: uint16(sym_comment),
	7091: uint16(910),
	7092: uint16(1),
	7093: uint16(sym__newline),
	7094: uint16(2),
	7095: uint16(29),
	7096: uint16(1),
	7097: uint16(sym_comment),
	7098: uint16(912),
	7099: uint16(1),
	7100: uint16(sym_identifier),
	7101: uint16(2),
	7102: uint16(29),
	7103: uint16(1),
	7104: uint16(sym_comment),
	7105: uint16(914),
	7106: uint16(1),
	7107: uint16(sym__newline),
	7108: uint16(2),
	7109: uint16(29),
	7110: uint16(1),
	7111: uint16(sym_comment),
	7112: uint16(916),
	7113: uint16(1),
	7114: uint16(sym_identifier),
	7115: uint16(2),
	7116: uint16(29),
	7117: uint16(1),
	7118: uint16(sym_comment),
	7119: uint16(918),
	7120: uint16(1),
	7121: uint16(anon_sym_COLON),
	7122: uint16(2),
	7123: uint16(29),
	7124: uint16(1),
	7125: uint16(sym_comment),
	7126: uint16(920),
	7127: uint16(1),
	7128: uint16(sym__newline),
	7129: uint16(2),
	7130: uint16(29),
	7131: uint16(1),
	7132: uint16(sym_comment),
	7133: uint16(922),
	7134: uint16(1),
	7135: uint16(sym_identifier),
	7136: uint16(2),
	7137: uint16(29),
	7138: uint16(1),
	7139: uint16(sym_comment),
	7140: uint16(924),
	7141: uint16(1),
	7142: uint16(anon_sym_LBRACK),
	7143: uint16(2),
	7144: uint16(29),
	7145: uint16(1),
	7146: uint16(sym_comment),
	7147: uint16(926),
	7148: uint16(1),
	7149: uint16(anon_sym_RPAREN),
	7150: uint16(2),
	7151: uint16(29),
	7152: uint16(1),
	7153: uint16(sym_comment),
	7154: uint16(928),
	7155: uint16(1),
	7156: uint16(sym__newline),
	7157: uint16(2),
	7158: uint16(29),
	7159: uint16(1),
	7160: uint16(sym_comment),
	7161: uint16(492),
	7162: uint16(1),
	7163: uint16(sym__newline),
	7164: uint16(2),
	7165: uint16(29),
	7166: uint16(1),
	7167: uint16(sym_comment),
	7168: uint16(930),
	7169: uint16(1),
	7170: uint16(sym_identifier),
	7171: uint16(2),
	7172: uint16(29),
	7173: uint16(1),
	7174: uint16(sym_comment),
	7175: uint16(932),
	7176: uint16(1),
	7177: uint16(sym__newline),
	7178: uint16(2),
	7179: uint16(29),
	7180: uint16(1),
	7181: uint16(sym_comment),
	7182: uint16(934),
	7183: uint16(1),
	7184: uint16(sym__newline),
	7185: uint16(2),
	7186: uint16(29),
	7187: uint16(1),
	7188: uint16(sym_comment),
	7189: uint16(936),
	7190: uint16(1),
	7191: uint16(sym__newline),
	7192: uint16(2),
	7193: uint16(29),
	7194: uint16(1),
	7195: uint16(sym_comment),
	7196: uint16(938),
	7197: uint16(1),
	7198: uint16(sym__newline),
	7199: uint16(2),
	7200: uint16(3),
	7201: uint16(1),
	7202: uint16(sym_comment),
	7203: uint16(940),
	7204: uint16(1),
	7205: uint16(aux_sym__shebang_with_lang_token3),
	7206: uint16(2),
	7207: uint16(29),
	7208: uint16(1),
	7209: uint16(sym_comment),
	7210: uint16(942),
	7211: uint16(1),
	7212: uint16(sym__newline),
	7213: uint16(2),
	7214: uint16(29),
	7215: uint16(1),
	7216: uint16(sym_comment),
	7217: uint16(944),
	7218: uint16(1),
	7219: uint16(sym__newline),
	7220: uint16(2),
	7221: uint16(29),
	7222: uint16(1),
	7223: uint16(sym_comment),
	7224: uint16(764),
	7225: uint16(1),
	7226: uint16(anon_sym_if),
	7227: uint16(2),
	7228: uint16(29),
	7229: uint16(1),
	7230: uint16(sym_comment),
	7231: uint16(946),
	7232: uint16(1),
	7233: uint16(sym__newline),
	7234: uint16(2),
	7235: uint16(29),
	7236: uint16(1),
	7237: uint16(sym_comment),
	7238: uint16(948),
	7239: uint16(1),
	7240: uint16(anon_sym_BQUOTE),
	7241: uint16(2),
	7242: uint16(29),
	7243: uint16(1),
	7244: uint16(sym_comment),
	7245: uint16(950),
	7246: uint16(1),
	7247: uint16(anon_sym_RBRACE_RBRACE),
	7248: uint16(2),
	7249: uint16(29),
	7250: uint16(1),
	7251: uint16(sym_comment),
	7252: uint16(952),
	7253: uint16(1),
	7254: uint16(anon_sym_COLON_EQ),
	7255: uint16(2),
	7256: uint16(29),
	7257: uint16(1),
	7258: uint16(sym_comment),
	7259: uint16(613),
	7260: uint16(1),
	7261: uint16(anon_sym_COLON),
	7262: uint16(2),
	7263: uint16(29),
	7264: uint16(1),
	7265: uint16(sym_comment),
	7266: uint16(954),
	7267: uint16(1),
	7268: uint16(anon_sym_RBRACK),
	7269: uint16(2),
	7270: uint16(29),
	7271: uint16(1),
	7272: uint16(sym_comment),
	7273: uint16(956),
	7274: uint16(1),
	7275: uint16(anon_sym_COLON_EQ),
	7276: uint16(2),
	7277: uint16(29),
	7278: uint16(1),
	7279: uint16(sym_comment),
	7280: uint16(958),
	7281: uint16(1),
	7282: uint16(sym__newline),
	7283: uint16(2),
	7284: uint16(29),
	7285: uint16(1),
	7286: uint16(sym_comment),
	7287: uint16(960),
	7288: uint16(1),
	7289: uint16(anon_sym_RBRACK),
	7290: uint16(2),
	7291: uint16(29),
	7292: uint16(1),
	7293: uint16(sym_comment),
	7294: uint16(962),
	7295: uint16(1),
	7296: uint16(sym__newline),
	7297: uint16(2),
	7298: uint16(29),
	7299: uint16(1),
	7300: uint16(sym_comment),
	7301: uint16(964),
	7302: uint16(1),
	7303: uint16(anon_sym_COLON),
	7304: uint16(2),
	7305: uint16(29),
	7306: uint16(1),
	7307: uint16(sym_comment),
	7308: uint16(966),
	7309: uint16(1),
	7310: uint16(sym__newline),
	7311: uint16(2),
	7312: uint16(29),
	7313: uint16(1),
	7314: uint16(sym_comment),
	7315: uint16(968),
	7316: uint16(1),
	7318: uint16(2),
	7319: uint16(29),
	7320: uint16(1),
	7321: uint16(sym_comment),
	7322: uint16(970),
	7323: uint16(1),
	7324: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	7325: uint16(2),
	7326: uint16(29),
	7327: uint16(1),
	7328: uint16(sym_comment),
	7329: uint16(972),
	7330: uint16(1),
	7331: uint16(anon_sym_RBRACK),
	7332: uint16(2),
	7333: uint16(29),
	7334: uint16(1),
	7335: uint16(sym_comment),
	7336: uint16(974),
	7337: uint16(1),
	7338: uint16(sym__newline),
	7339: uint16(2),
	7340: uint16(29),
	7341: uint16(1),
	7342: uint16(sym_comment),
	7343: uint16(976),
	7344: uint16(1),
	7345: uint16(anon_sym_RPAREN),
	7346: uint16(2),
	7347: uint16(29),
	7348: uint16(1),
	7349: uint16(sym_comment),
	7350: uint16(978),
	7351: uint16(1),
	7352: uint16(anon_sym_BQUOTE),
	7353: uint16(2),
	7354: uint16(29),
	7355: uint16(1),
	7356: uint16(sym_comment),
	7357: uint16(980),
	7358: uint16(1),
	7359: uint16(sym_identifier),
	7360: uint16(2),
	7361: uint16(29),
	7362: uint16(1),
	7363: uint16(sym_comment),
	7364: uint16(510),
	7365: uint16(1),
	7366: uint16(sym__newline),
	7367: uint16(2),
	7368: uint16(29),
	7369: uint16(1),
	7370: uint16(sym_comment),
	7371: uint16(982),
	7372: uint16(1),
	7373: uint16(anon_sym_RBRACE_RBRACE),
	7374: uint16(2),
	7375: uint16(29),
	7376: uint16(1),
	7377: uint16(sym_comment),
	7378: uint16(984),
	7379: uint16(1),
	7380: uint16(anon_sym_RPAREN),
	7381: uint16(2),
	7382: uint16(29),
	7383: uint16(1),
	7384: uint16(sym_comment),
	7385: uint16(986),
	7386: uint16(1),
	7387: uint16(anon_sym_RBRACE),
	7388: uint16(2),
	7389: uint16(29),
	7390: uint16(1),
	7391: uint16(sym_comment),
	7392: uint16(988),
	7393: uint16(1),
	7394: uint16(sym__newline),
	7395: uint16(2),
	7396: uint16(29),
	7397: uint16(1),
	7398: uint16(sym_comment),
	7399: uint16(990),
	7400: uint16(1),
	7401: uint16(sym__newline),
	7402: uint16(2),
	7403: uint16(29),
	7404: uint16(1),
	7405: uint16(sym_comment),
	7406: uint16(992),
	7407: uint16(1),
	7408: uint16(sym__newline),
	7409: uint16(2),
	7410: uint16(29),
	7411: uint16(1),
	7412: uint16(sym_comment),
	7413: uint16(994),
	7414: uint16(1),
	7415: uint16(anon_sym_RBRACK),
	7416: uint16(2),
	7417: uint16(29),
	7418: uint16(1),
	7419: uint16(sym_comment),
	7420: uint16(354),
	7421: uint16(1),
	7422: uint16(anon_sym_COLON_EQ),
	7423: uint16(2),
	7424: uint16(29),
	7425: uint16(1),
	7426: uint16(sym_comment),
	7427: uint16(996),
	7428: uint16(1),
	7429: uint16(anon_sym_COLON_EQ),
	7430: uint16(2),
	7431: uint16(29),
	7432: uint16(1),
	7433: uint16(sym_comment),
	7434: uint16(998),
	7435: uint16(1),
	7436: uint16(sym__newline),
	7437: uint16(2),
	7438: uint16(29),
	7439: uint16(1),
	7440: uint16(sym_comment),
	7441: uint16(607),
	7442: uint16(1),
	7443: uint16(anon_sym_COLON),
	7444: uint16(2),
	7445: uint16(29),
	7446: uint16(1),
	7447: uint16(sym_comment),
	7448: uint16(1000),
	7449: uint16(1),
	7450: uint16(sym__newline),
	7451: uint16(2),
	7452: uint16(29),
	7453: uint16(1),
	7454: uint16(sym_comment),
	7455: uint16(1002),
	7456: uint16(1),
	7457: uint16(sym_identifier),
	7458: uint16(2),
	7459: uint16(29),
	7460: uint16(1),
	7461: uint16(sym_comment),
	7462: uint16(1004),
	7463: uint16(1),
	7464: uint16(sym__newline),
	7465: uint16(2),
	7466: uint16(29),
	7467: uint16(1),
	7468: uint16(sym_comment),
	7469: uint16(1006),
	7470: uint16(1),
	7471: uint16(sym__newline),
	7472: uint16(2),
	7473: uint16(29),
	7474: uint16(1),
	7475: uint16(sym_comment),
	7476: uint16(482),
	7477: uint16(1),
	7478: uint16(sym__newline),
	7479: uint16(2),
	7480: uint16(29),
	7481: uint16(1),
	7482: uint16(sym_comment),
	7483: uint16(1008),
	7484: uint16(1),
	7485: uint16(anon_sym_COLON),
	7486: uint16(2),
	7487: uint16(29),
	7488: uint16(1),
	7489: uint16(sym_comment),
	7490: uint16(1010),
	7491: uint16(1),
	7492: uint16(anon_sym_RPAREN),
	7493: uint16(2),
	7494: uint16(29),
	7495: uint16(1),
	7496: uint16(sym_comment),
	7497: uint16(1012),
	7498: uint16(1),
	7499: uint16(anon_sym_BQUOTE),
	7500: uint16(2),
	7501: uint16(29),
	7502: uint16(1),
	7503: uint16(sym_comment),
	7504: uint16(1014),
	7505: uint16(1),
	7506: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	7507: uint16(2),
	7508: uint16(29),
	7509: uint16(1),
	7510: uint16(sym_comment),
	7511: uint16(1016),
	7512: uint16(1),
	7513: uint16(sym__newline),
	7514: uint16(2),
	7515: uint16(29),
	7516: uint16(1),
	7517: uint16(sym_comment),
	7518: uint16(1018),
	7519: uint16(1),
	7520: uint16(anon_sym_RBRACE_RBRACE),
	7521: uint16(2),
	7522: uint16(29),
	7523: uint16(1),
	7524: uint16(sym_comment),
	7525: uint16(1020),
	7526: uint16(1),
	7527: uint16(anon_sym_RPAREN),
	7528: uint16(2),
	7529: uint16(29),
	7530: uint16(1),
	7531: uint16(sym_comment),
	7532: uint16(1022),
	7533: uint16(1),
	7534: uint16(anon_sym_RBRACE),
	7535: uint16(2),
	7536: uint16(29),
	7537: uint16(1),
	7538: uint16(sym_comment),
	7539: uint16(1024),
	7540: uint16(1),
	7541: uint16(sym__newline),
	7542: uint16(2),
	7543: uint16(29),
	7544: uint16(1),
	7545: uint16(sym_comment),
	7546: uint16(484),
	7547: uint16(1),
	7548: uint16(sym__newline),
	7549: uint16(2),
	7550: uint16(29),
	7551: uint16(1),
	7552: uint16(sym_comment),
	7553: uint16(1026),
	7554: uint16(1),
	7555: uint16(sym_identifier),
	7556: uint16(2),
	7557: uint16(29),
	7558: uint16(1),
	7559: uint16(sym_comment),
	7560: uint16(1028),
	7561: uint16(1),
	7562: uint16(sym__newline),
	7563: uint16(2),
	7564: uint16(29),
	7565: uint16(1),
	7566: uint16(sym_comment),
	7567: uint16(1030),
	7568: uint16(1),
	7569: uint16(sym__newline),
	7570: uint16(2),
	7571: uint16(29),
	7572: uint16(1),
	7573: uint16(sym_comment),
	7574: uint16(1032),
	7575: uint16(1),
	7576: uint16(anon_sym_RPAREN),
	7577: uint16(2),
	7578: uint16(29),
	7579: uint16(1),
	7580: uint16(sym_comment),
	7581: uint16(1034),
	7582: uint16(1),
	7583: uint16(sym_identifier),
	7584: uint16(2),
	7585: uint16(29),
	7586: uint16(1),
	7587: uint16(sym_comment),
	7588: uint16(1036),
	7589: uint16(1),
	7590: uint16(sym_identifier),
	7591: uint16(2),
	7592: uint16(29),
	7593: uint16(1),
	7594: uint16(sym_comment),
	7595: uint16(830),
	7596: uint16(1),
	7597: uint16(anon_sym_if),
	7598: uint16(2),
	7599: uint16(29),
	7600: uint16(1),
	7601: uint16(sym_comment),
	7602: uint16(1038),
	7603: uint16(1),
	7604: uint16(sym__newline),
	7605: uint16(2),
	7606: uint16(29),
	7607: uint16(1),
	7608: uint16(sym_comment),
	7609: uint16(1040),
	7610: uint16(1),
	7611: uint16(sym_identifier),
	7612: uint16(2),
	7613: uint16(29),
	7614: uint16(1),
	7615: uint16(sym_comment),
	7616: uint16(1042),
	7617: uint16(1),
	7618: uint16(sym__newline),
	7619: uint16(2),
	7620: uint16(29),
	7621: uint16(1),
	7622: uint16(sym_comment),
	7623: uint16(1044),
	7624: uint16(1),
	7625: uint16(sym__newline),
	7626: uint16(2),
	7627: uint16(29),
	7628: uint16(1),
	7629: uint16(sym_comment),
	7630: uint16(1046),
	7631: uint16(1),
	7632: uint16(sym__newline),
	7633: uint16(2),
	7634: uint16(29),
	7635: uint16(1),
	7636: uint16(sym_comment),
	7637: uint16(840),
	7638: uint16(1),
	7639: uint16(anon_sym_if),
	7640: uint16(2),
	7641: uint16(29),
	7642: uint16(1),
	7643: uint16(sym_comment),
	7644: uint16(1048),
	7645: uint16(1),
	7646: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
}

var ts_small_parse_table_map = [403]uint32_t{
	1:   uint32(32),
	2:   uint32(64),
	3:   uint32(96),
	4:   uint32(135),
	5:   uint32(174),
	6:   uint32(204),
	7:   uint32(234),
	8:   uint32(296),
	9:   uint32(356),
	10:  uint32(386),
	11:  uint32(446),
	12:  uint32(476),
	13:  uint32(506),
	14:  uint32(536),
	15:  uint32(566),
	16:  uint32(596),
	17:  uint32(656),
	18:  uint32(686),
	19:  uint32(722),
	20:  uint32(784),
	21:  uint32(846),
	22:  uint32(895),
	23:  uint32(954),
	24:  uint32(1013),
	25:  uint32(1062),
	26:  uint32(1121),
	27:  uint32(1170),
	28:  uint32(1219),
	29:  uint32(1278),
	30:  uint32(1337),
	31:  uint32(1396),
	32:  uint32(1455),
	33:  uint32(1511),
	34:  uint32(1567),
	35:  uint32(1623),
	36:  uint32(1679),
	37:  uint32(1735),
	38:  uint32(1763),
	39:  uint32(1819),
	40:  uint32(1847),
	41:  uint32(1903),
	42:  uint32(1959),
	43:  uint32(2015),
	44:  uint32(2071),
	45:  uint32(2127),
	46:  uint32(2183),
	47:  uint32(2211),
	48:  uint32(2238),
	49:  uint32(2265),
	50:  uint32(2292),
	51:  uint32(2319),
	52:  uint32(2346),
	53:  uint32(2396),
	54:  uint32(2446),
	55:  uint32(2496),
	56:  uint32(2546),
	57:  uint32(2596),
	58:  uint32(2646),
	59:  uint32(2696),
	60:  uint32(2746),
	61:  uint32(2796),
	62:  uint32(2846),
	63:  uint32(2896),
	64:  uint32(2946),
	65:  uint32(2971),
	66:  uint32(2996),
	67:  uint32(3021),
	68:  uint32(3064),
	69:  uint32(3107),
	70:  uint32(3150),
	71:  uint32(3193),
	72:  uint32(3236),
	73:  uint32(3283),
	74:  uint32(3315),
	75:  uint32(3347),
	76:  uint32(3369),
	77:  uint32(3394),
	78:  uint32(3431),
	79:  uint32(3468),
	80:  uint32(3489),
	81:  uint32(3512),
	82:  uint32(3537),
	83:  uint32(3574),
	84:  uint32(3611),
	85:  uint32(3634),
	86:  uint32(3657),
	87:  uint32(3689),
	88:  uint32(3706),
	89:  uint32(3723),
	90:  uint32(3740),
	91:  uint32(3757),
	92:  uint32(3784),
	93:  uint32(3803),
	94:  uint32(3832),
	95:  uint32(3851),
	96:  uint32(3880),
	97:  uint32(3897),
	98:  uint32(3914),
	99:  uint32(3931),
	100: uint32(3948),
	101: uint32(3965),
	102: uint32(3982),
	103: uint32(3999),
	104: uint32(4016),
	105: uint32(4045),
	106: uint32(4074),
	107: uint32(4091),
	108: uint32(4108),
	109: uint32(4125),
	110: uint32(4142),
	111: uint32(4159),
	112: uint32(4176),
	113: uint32(4193),
	114: uint32(4210),
	115: uint32(4227),
	116: uint32(4244),
	117: uint32(4261),
	118: uint32(4278),
	119: uint32(4295),
	120: uint32(4312),
	121: uint32(4329),
	122: uint32(4355),
	123: uint32(4381),
	124: uint32(4407),
	125: uint32(4433),
	126: uint32(4459),
	127: uint32(4485),
	128: uint32(4513),
	129: uint32(4539),
	130: uint32(4567),
	131: uint32(4595),
	132: uint32(4623),
	133: uint32(4639),
	134: uint32(4665),
	135: uint32(4691),
	136: uint32(4714),
	137: uint32(4735),
	138: uint32(4756),
	139: uint32(4779),
	140: uint32(4802),
	141: uint32(4823),
	142: uint32(4842),
	143: uint32(4861),
	144: uint32(4884),
	145: uint32(4905),
	146: uint32(4928),
	147: uint32(4951),
	148: uint32(4973),
	149: uint32(4991),
	150: uint32(5011),
	151: uint32(5031),
	152: uint32(5053),
	153: uint32(5067),
	154: uint32(5085),
	155: uint32(5103),
	156: uint32(5117),
	157: uint32(5137),
	158: uint32(5157),
	159: uint32(5177),
	160: uint32(5191),
	161: uint32(5211),
	162: uint32(5224),
	163: uint32(5235),
	164: uint32(5248),
	165: uint32(5267),
	166: uint32(5286),
	167: uint32(5299),
	168: uint32(5312),
	169: uint32(5329),
	170: uint32(5346),
	171: uint32(5363),
	172: uint32(5382),
	173: uint32(5401),
	174: uint32(5420),
	175: uint32(5433),
	176: uint32(5446),
	177: uint32(5459),
	178: uint32(5472),
	179: uint32(5485),
	180: uint32(5504),
	181: uint32(5521),
	182: uint32(5532),
	183: uint32(5549),
	184: uint32(5562),
	185: uint32(5579),
	186: uint32(5596),
	187: uint32(5610),
	188: uint32(5624),
	189: uint32(5634),
	190: uint32(5650),
	191: uint32(5660),
	192: uint32(5672),
	193: uint32(5682),
	194: uint32(5692),
	195: uint32(5702),
	196: uint32(5716),
	197: uint32(5730),
	198: uint32(5744),
	199: uint32(5760),
	200: uint32(5774),
	201: uint32(5784),
	202: uint32(5794),
	203: uint32(5810),
	204: uint32(5820),
	205: uint32(5832),
	206: uint32(5846),
	207: uint32(5860),
	208: uint32(5874),
	209: uint32(5888),
	210: uint32(5898),
	211: uint32(5908),
	212: uint32(5922),
	213: uint32(5936),
	214: uint32(5950),
	215: uint32(5964),
	216: uint32(5978),
	217: uint32(5992),
	218: uint32(6006),
	219: uint32(6016),
	220: uint32(6025),
	221: uint32(6038),
	222: uint32(6049),
	223: uint32(6062),
	224: uint32(6071),
	225: uint32(6084),
	226: uint32(6093),
	227: uint32(6106),
	228: uint32(6115),
	229: uint32(6126),
	230: uint32(6139),
	231: uint32(6152),
	232: uint32(6161),
	233: uint32(6170),
	234: uint32(6179),
	235: uint32(6188),
	236: uint32(6201),
	237: uint32(6214),
	238: uint32(6223),
	239: uint32(6232),
	240: uint32(6245),
	241: uint32(6254),
	242: uint32(6263),
	243: uint32(6276),
	244: uint32(6289),
	245: uint32(6298),
	246: uint32(6311),
	247: uint32(6320),
	248: uint32(6333),
	249: uint32(6346),
	250: uint32(6359),
	251: uint32(6372),
	252: uint32(6381),
	253: uint32(6390),
	254: uint32(6403),
	255: uint32(6412),
	256: uint32(6421),
	257: uint32(6434),
	258: uint32(6443),
	259: uint32(6452),
	260: uint32(6465),
	261: uint32(6474),
	262: uint32(6483),
	263: uint32(6492),
	264: uint32(6501),
	265: uint32(6510),
	266: uint32(6519),
	267: uint32(6528),
	268: uint32(6537),
	269: uint32(6546),
	270: uint32(6555),
	271: uint32(6564),
	272: uint32(6573),
	273: uint32(6586),
	274: uint32(6599),
	275: uint32(6610),
	276: uint32(6623),
	277: uint32(6636),
	278: uint32(6645),
	279: uint32(6658),
	280: uint32(6671),
	281: uint32(6684),
	282: uint32(6697),
	283: uint32(6710),
	284: uint32(6723),
	285: uint32(6734),
	286: uint32(6747),
	287: uint32(6760),
	288: uint32(6769),
	289: uint32(6782),
	290: uint32(6795),
	291: uint32(6808),
	292: uint32(6821),
	293: uint32(6829),
	294: uint32(6839),
	295: uint32(6849),
	296: uint32(6857),
	297: uint32(6865),
	298: uint32(6875),
	299: uint32(6885),
	300: uint32(6895),
	301: uint32(6903),
	302: uint32(6913),
	303: uint32(6923),
	304: uint32(6933),
	305: uint32(6943),
	306: uint32(6953),
	307: uint32(6963),
	308: uint32(6973),
	309: uint32(6981),
	310: uint32(6989),
	311: uint32(6997),
	312: uint32(7007),
	313: uint32(7017),
	314: uint32(7024),
	315: uint32(7031),
	316: uint32(7038),
	317: uint32(7045),
	318: uint32(7052),
	319: uint32(7059),
	320: uint32(7066),
	321: uint32(7073),
	322: uint32(7080),
	323: uint32(7087),
	324: uint32(7094),
	325: uint32(7101),
	326: uint32(7108),
	327: uint32(7115),
	328: uint32(7122),
	329: uint32(7129),
	330: uint32(7136),
	331: uint32(7143),
	332: uint32(7150),
	333: uint32(7157),
	334: uint32(7164),
	335: uint32(7171),
	336: uint32(7178),
	337: uint32(7185),
	338: uint32(7192),
	339: uint32(7199),
	340: uint32(7206),
	341: uint32(7213),
	342: uint32(7220),
	343: uint32(7227),
	344: uint32(7234),
	345: uint32(7241),
	346: uint32(7248),
	347: uint32(7255),
	348: uint32(7262),
	349: uint32(7269),
	350: uint32(7276),
	351: uint32(7283),
	352: uint32(7290),
	353: uint32(7297),
	354: uint32(7304),
	355: uint32(7311),
	356: uint32(7318),
	357: uint32(7325),
	358: uint32(7332),
	359: uint32(7339),
	360: uint32(7346),
	361: uint32(7353),
	362: uint32(7360),
	363: uint32(7367),
	364: uint32(7374),
	365: uint32(7381),
	366: uint32(7388),
	367: uint32(7395),
	368: uint32(7402),
	369: uint32(7409),
	370: uint32(7416),
	371: uint32(7423),
	372: uint32(7430),
	373: uint32(7437),
	374: uint32(7444),
	375: uint32(7451),
	376: uint32(7458),
	377: uint32(7465),
	378: uint32(7472),
	379: uint32(7479),
	380: uint32(7486),
	381: uint32(7493),
	382: uint32(7500),
	383: uint32(7507),
	384: uint32(7514),
	385: uint32(7521),
	386: uint32(7528),
	387: uint32(7535),
	388: uint32(7542),
	389: uint32(7549),
	390: uint32(7556),
	391: uint32(7563),
	392: uint32(7570),
	393: uint32(7577),
	394: uint32(7584),
	395: uint32(7591),
	396: uint32(7598),
	397: uint32(7605),
	398: uint32(7612),
	399: uint32(7619),
	400: uint32(7626),
	401: uint32(7633),
	402: uint32(7640),
}

var ts_parse_actions = [1050]TSParseActionEntry{
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
		Fstate: uint16(libc.Int32FromInt32(331)),
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
		Fstate: uint16(libc.Int32FromInt32(295)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(74)),
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
		Fstate: uint16(libc.Int32FromInt32(301)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(309)),
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
		Fstate: uint16(libc.Int32FromInt32(395)),
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
		Fstate: uint16(libc.Int32FromInt32(296)),
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
		Fstate: uint16(libc.Int32FromInt32(244)),
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
		Fsymbol:      uint16(sym_string),
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
		Fcount: uint8(1),
	}})),
	28: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fextra: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(sym_string),
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
		Fcount: uint8(1),
	}})),
	34: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fsymbol:      uint16(sym_string),
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
		Fcount: uint8(1),
	}})),
	38: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_if_expression),
		Fproduction_id: uint16(17),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_if_expression),
		Fproduction_id: uint16(17),
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
		Fstate: uint16(libc.Int32FromInt32(280)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_if_expression),
		Fproduction_id: uint16(27),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_if_expression),
		Fproduction_id: uint16(27),
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
		Fcount: uint8(1),
	}})),
	50: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_external_command),
	})))),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_external_command),
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
		Fcount: uint8(1),
	}})),
	54: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: uint16(libc.Int32FromInt32(77)),
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
		Fstate: uint16(libc.Int32FromInt32(62)),
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
		Fstate: uint16(libc.Int32FromInt32(30)),
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
		Fstate: uint16(libc.Int32FromInt32(43)),
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
		Fstate: uint16(libc.Int32FromInt32(267)),
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
		Fstate: uint16(libc.Int32FromInt32(2)),
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
		Fstate: uint16(libc.Int32FromInt32(281)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(217)),
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
		Fstate: uint16(libc.Int32FromInt32(219)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(151)),
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
		Fstate: uint16(libc.Int32FromInt32(152)),
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
		Fstate: uint16(libc.Int32FromInt32(11)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(81)),
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
		Fstate: uint16(libc.Int32FromInt32(63)),
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
		Fstate: uint16(libc.Int32FromInt32(190)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(67)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(213)),
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
		Fstate: uint16(libc.Int32FromInt32(214)),
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
		Fcount: uint8(1),
	}})),
	96: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fcount: uint8(2),
	}})),
	100: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_dependency_expression_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(81)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_dependency_expression_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(63)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	106: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_dependency_expression_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(30)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_dependency_expression_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(43)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_dependency_expression_repeat1),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_dependency_expression_repeat1),
	})))),
	115: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(libc.Int32FromInt32(67)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_dependency_expression_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(282)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_dependency_expression_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(213)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_dependency_expression_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(214)),
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
		Fsymbol:      uint16(aux_sym_dependency_expression_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(151)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_dependency_expression_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(152)),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_dependency_expression_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(11)),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym__backticked),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__backticked),
	})))),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym__indented_backticked),
	})))),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__indented_backticked),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__backticked),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__backticked),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__indented_backticked),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__indented_backticked),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_function_call),
		Fproduction_id: uint16(2),
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
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_function_call),
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
		Fstate: uint16(libc.Int32FromInt32(220)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_function_call),
		Fproduction_id: uint16(28),
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
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_function_call),
		Fproduction_id: uint16(28),
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
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_if_expression_repeat1),
		Fproduction_id: uint16(38),
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
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_if_expression_repeat1),
		Fproduction_id: uint16(38),
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
		Fcount: uint8(2),
	}})),
	165: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_if_expression_repeat1),
		Fproduction_id: uint16(38),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(397)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(17)),
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
		Fstate: uint16(libc.Int32FromInt32(265)),
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
		Fstate:      uint16(libc.Int32FromInt32(88)),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_source_file_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(331)),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_source_file_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(295)),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_source_file_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(74)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_source_file_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(301)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	193: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      uint16(libc.Int32FromInt32(309)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_source_file_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(395)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_source_file_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(296)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_source_file),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(193)),
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
		Fstate: uint16(libc.Int32FromInt32(60)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(27)),
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
		Fstate: uint16(libc.Int32FromInt32(38)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(242)),
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
		Fstate: uint16(libc.Int32FromInt32(277)),
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
		Fstate: uint16(libc.Int32FromInt32(207)),
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
		Fstate: uint16(libc.Int32FromInt32(208)),
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
		Fstate: uint16(libc.Int32FromInt32(158)),
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
		Fstate: uint16(libc.Int32FromInt32(160)),
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
		Fstate: uint16(libc.Int32FromInt32(234)),
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
		Fstate: uint16(libc.Int32FromInt32(206)),
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
		Fstate: uint16(libc.Int32FromInt32(59)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(32)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(46)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(162)),
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
		Fstate: uint16(libc.Int32FromInt32(159)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(254)),
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
		Fcount: uint8(1),
	}})),
	240: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(aux_sym_if_expression_repeat1),
		Fproduction_id: uint16(25),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(aux_sym_if_expression_repeat1),
		Fproduction_id: uint16(25),
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
		Fcount: uint8(1),
	}})),
	244: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym__braced_expr),
		Fproduction_id: uint16(35),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym__braced_expr),
		Fproduction_id: uint16(35),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_else_if_clause),
		Fproduction_id: uint16(48),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_else_if_clause),
		Fproduction_id: uint16(48),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_if_expression),
		Fproduction_id: uint16(37),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_if_expression),
		Fproduction_id: uint16(37),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_else_clause),
		Fproduction_id: uint16(36),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_else_clause),
		Fproduction_id: uint16(36),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_if_expression),
		Fproduction_id: uint16(26),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_if_expression),
		Fproduction_id: uint16(26),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym__expression_inner),
		Fproduction_id: uint16(18),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym__expression_inner),
		Fproduction_id: uint16(18),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__expression_inner),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__expression_inner),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(312)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(47)),
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
		Fstate: uint16(libc.Int32FromInt32(111)),
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
		Fstate: uint16(libc.Int32FromInt32(84)),
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
		Fstate: uint16(libc.Int32FromInt32(187)),
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
		Fstate: uint16(libc.Int32FromInt32(77)),
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
		Fstate: uint16(libc.Int32FromInt32(161)),
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
		Fstate: uint16(libc.Int32FromInt32(181)),
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
		Fstate: uint16(libc.Int32FromInt32(245)),
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
		Fstate: uint16(libc.Int32FromInt32(123)),
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
		Fstate: uint16(libc.Int32FromInt32(326)),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_module),
		Fproduction_id: uint16(5),
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
		Fcount: uint8(1),
	}})),
	298: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_module),
		Fproduction_id: uint16(5),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_module),
		Fproduction_id: uint16(3),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_module),
		Fproduction_id: uint16(3),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_expression),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(57)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_expression),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_recipe_body_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(312)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	317: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_recipe_body_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(312)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_recipe_body_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(47)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_recipe_body_repeat1),
	})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	325: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_recipe_body_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(79)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_recipe_body_repeat1),
	})))),
	329: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(112)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	333: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	334: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	335: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_expression),
	})))),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_expression),
	})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(100)),
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
		Fstate: uint16(libc.Int32FromInt32(80)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_recipe),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_recipe),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(69)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	351: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_recipe),
	})))),
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
		Fcount: uint8(1),
	}})),
	353: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_recipe),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(34)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(130)),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_export),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_export),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	363: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(9),
		Fsymbol:        uint16(sym_setting),
		Fproduction_id: uint16(45),
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
		Fcount: uint8(1),
	}})),
	365: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(9),
		Fsymbol:        uint16(sym_setting),
		Fproduction_id: uint16(45),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(9),
		Fsymbol:        uint16(sym_setting),
		Fproduction_id: uint16(46),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(9),
		Fsymbol:        uint16(sym_setting),
		Fproduction_id: uint16(46),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_setting),
		Fproduction_id: uint16(6),
	})))),
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
		Fcount: uint8(1),
	}})),
	373: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_setting),
		Fproduction_id: uint16(6),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	375: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	376: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	377: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	378: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(64)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	381: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	382: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(135)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	387: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_setting),
		Fproduction_id: uint16(20),
	})))),
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
		Fcount: uint8(1),
	}})),
	389: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_setting),
		Fproduction_id: uint16(20),
	})))),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_recipe_body),
		Fproduction_id: uint16(29),
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
		Fcount: uint8(1),
	}})),
	393: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_recipe_body),
		Fproduction_id: uint16(29),
	})))),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_alias),
		Fproduction_id: uint16(8),
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
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_alias),
		Fproduction_id: uint16(8),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_setting),
		Fproduction_id: uint16(31),
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
		Fcount: uint8(1),
	}})),
	401: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_setting),
		Fproduction_id: uint16(31),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_setting),
		Fproduction_id: uint16(33),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_setting),
		Fproduction_id: uint16(33),
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
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_setting),
		Fproduction_id: uint16(21),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_setting),
		Fproduction_id: uint16(21),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fsymbol:        uint16(sym_assignment),
		Fproduction_id: uint16(12),
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
		Fcount: uint8(1),
	}})),
	415: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_assignment),
		Fproduction_id: uint16(12),
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
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_module),
		Fproduction_id: uint16(3),
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
		Fcount: uint8(1),
	}})),
	419: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_module),
		Fproduction_id: uint16(3),
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
		Fsymbol:      uint16(sym_import),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_import),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_recipe_body),
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
		Fcount: uint8(1),
	}})),
	427: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_recipe_body),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_recipe_body),
		Fproduction_id: uint16(29),
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
		Fcount: uint8(1),
	}})),
	431: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_recipe_body),
		Fproduction_id: uint16(29),
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
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_setting),
		Fproduction_id: uint16(39),
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
		Fcount: uint8(1),
	}})),
	435: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_setting),
		Fproduction_id: uint16(39),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_recipe),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_recipe),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_setting),
		Fproduction_id: uint16(40),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_setting),
		Fproduction_id: uint16(40),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_setting),
		Fproduction_id: uint16(14),
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
		Fcount: uint8(1),
	}})),
	447: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_setting),
		Fproduction_id: uint16(14),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_import),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_import),
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
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_setting),
		Fproduction_id: uint16(41),
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
		Fcount: uint8(1),
	}})),
	455: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_setting),
		Fproduction_id: uint16(41),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_recipe_body),
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
		Fcount: uint8(1),
	}})),
	459: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_recipe_body),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_setting),
		Fproduction_id: uint16(42),
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
		Fcount: uint8(1),
	}})),
	463: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_setting),
		Fproduction_id: uint16(42),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_alias),
		Fproduction_id: uint16(19),
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
		Fcount: uint8(1),
	}})),
	467: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_alias),
		Fproduction_id: uint16(19),
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
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_module),
		Fproduction_id: uint16(5),
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
		Fcount: uint8(1),
	}})),
	471: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_module),
		Fproduction_id: uint16(5),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(399)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_recipe_header),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_recipe_header),
		Fproduction_id: uint16(3),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_recipe_header),
		Fproduction_id: uint16(7),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	485: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_recipe_header),
		Fproduction_id: uint16(3),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(107)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(377)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(106)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_recipe_header),
		Fproduction_id: uint16(2),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(345)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_recipe_header),
		Fproduction_id: uint16(2),
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
		Fstate: uint16(libc.Int32FromInt32(356)),
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
		Fstate: uint16(libc.Int32FromInt32(401)),
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
		Fstate: uint16(libc.Int32FromInt32(368)),
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
		Fcount: uint8(1),
	}})),
	507: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_recipe_body_repeat1),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	509: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_recipe_header),
		Fproduction_id: uint16(7),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_recipe_header),
		Fproduction_id: uint16(1),
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
		Fstate: uint16(libc.Int32FromInt32(191)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(287)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_dependencies),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_alias_repeat1),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_alias_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(395)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_alias_repeat1),
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
		Fcount:    uint8(2),
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
		Fsymbol:        uint16(aux_sym_setting_repeat1),
		Fproduction_id: uint16(32),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(174)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_setting_repeat1),
		Fproduction_id: uint16(32),
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
		Fcount: uint8(1),
	}})),
	533: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_setting_repeat1),
		Fproduction_id: uint16(32),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(286)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	537: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_dependencies_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(211)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	540: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_dependencies_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(399)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	543: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_dependencies_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(204)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_dependencies_repeat1),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(370)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	552: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_if_expression_repeat1),
		Fproduction_id: uint16(38),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(403)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(186)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(13)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(36)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(182)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(14)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(44)),
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
		Fstate: uint16(libc.Int32FromInt32(402)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_setting_repeat1),
		Fproduction_id: uint16(30),
	})))),
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
		Fcount: uint8(1),
	}})),
	571: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_setting_repeat1),
		Fproduction_id: uint16(30),
	})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	573: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_parameters_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(161)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_parameters_repeat1),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	578: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_parameters_repeat1),
	})))),
	579: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	580: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	581: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_if_expression_repeat1),
		Fproduction_id: uint16(38),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(344)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_parameter),
		Fproduction_id: uint16(3),
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
		Fstate: uint16(libc.Int32FromInt32(72)),
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
		Fstate: uint16(libc.Int32FromInt32(269)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(260)),
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
		Fstate: uint16(libc.Int32FromInt32(279)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_parameter),
		Fproduction_id: uint16(2),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(71)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(259)),
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
		Fcount: uint8(1),
	}})),
	600: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_attribute),
		Fproduction_id: uint16(22),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_attribute),
		Fproduction_id: uint16(22),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_parameter),
		Fproduction_id: uint16(13),
	})))),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_parameters_repeat1),
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
		Fcount:    uint8(1),
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
		Fsymbol:      uint16(sym_parameters),
	})))),
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
		Fcount: uint8(1),
	}})),
	610: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_attribute),
		Fproduction_id: uint16(15),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_attribute),
		Fproduction_id: uint16(15),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_parameters),
	})))),
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
		Fcount: uint8(2),
	}})),
	616: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_command_body_repeat1),
	})))),
	617: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
	619: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_command_body_repeat1),
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
		Fcount: uint8(2),
	}})),
	621: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_command_body_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(36)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_recipe_line),
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
		Fstate: uint16(libc.Int32FromInt32(172)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	628: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_recipe_line_repeat1),
	})))),
	629: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_recipe_line_repeat1),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_recipe_line_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(172)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(336)),
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
		Fstate: uint16(libc.Int32FromInt32(352)),
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
		Fstate: uint16(libc.Int32FromInt32(148)),
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
		Fstate: uint16(libc.Int32FromInt32(173)),
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
		Fcount: uint8(1),
	}})),
	644: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_attribute),
		Fproduction_id: uint16(22),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_attribute),
		Fproduction_id: uint16(22),
	})))),
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
		Fcount: uint8(1),
	}})),
	648: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_attribute),
		Fproduction_id: uint16(34),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_attribute),
		Fproduction_id: uint16(34),
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
		Fcount: uint8(1),
	}})),
	652: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_attribute),
		Fproduction_id: uint16(44),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_attribute),
		Fproduction_id: uint16(44),
	})))),
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
		Fcount: uint8(1),
	}})),
	656: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_attribute),
		Fproduction_id: uint16(43),
	})))),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_attribute),
		Fproduction_id: uint16(43),
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
		Fcount: uint8(1),
	}})),
	660: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(9),
		Fsymbol:        uint16(sym_attribute),
		Fproduction_id: uint16(47),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(9),
		Fsymbol:        uint16(sym_attribute),
		Fproduction_id: uint16(47),
	})))),
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
		Fcount: uint8(1),
	}})),
	664: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
	666: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_command_body),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_parameter),
		Fproduction_id: uint16(9),
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
		Fcount: uint8(2),
	}})),
	670: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_command_body_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(184)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	673: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_command_body_repeat1),
	})))),
	674: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(libc.Int32FromInt32(44)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	675: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	676: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_attribute),
	})))),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_attribute),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(170)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_recipe_line),
	})))),
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
		Fcount: uint8(1),
	}})),
	684: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__string_indented_repeat1),
	})))),
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
		Fcount: uint8(2),
	}})),
	686: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__string_indented_repeat1),
	})))),
	687: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(libc.Int32FromInt32(188)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(4)),
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
		Fcount: uint8(1),
	}})),
	691: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_dependency_expression),
		Fproduction_id: uint16(3),
	})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(9)),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(aux_sym__string_repeat1),
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
		Fcount: uint8(2),
	}})),
	705: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__string_repeat1),
	})))),
	706: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(188)),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_attribute_repeat2),
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
		Fstate: uint16(libc.Int32FromInt32(137)),
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
		Fstate: uint16(libc.Int32FromInt32(166)),
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
		Fstate: uint16(libc.Int32FromInt32(292)),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_condition),
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
		Fstate: uint16(libc.Int32FromInt32(37)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(22)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(243)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(209)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(210)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(246)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_dependency),
		Fproduction_id: uint16(2),
	})))),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_dependency),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(66)),
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
		Fcount: uint8(1),
	}})),
	740: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(216)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(68)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(3)),
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
		Fcount: uint8(1),
	}})),
	748: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(171)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(198)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_dependency_expression),
		Fproduction_id: uint16(3),
	})))),
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
		Fcount: uint8(1),
	}})),
	756: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_interpolation),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	758: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      uint16(libc.Int32FromInt32(145)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_regex_literal),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(24)),
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
		Fstate: uint16(libc.Int32FromInt32(40)),
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
		Fstate: uint16(libc.Int32FromInt32(316)),
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
		Fstate: uint16(libc.Int32FromInt32(293)),
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
		Fstate: uint16(libc.Int32FromInt32(42)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_sequence),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(55)),
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
		Fstate: uint16(libc.Int32FromInt32(302)),
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
		Fstate: uint16(libc.Int32FromInt32(342)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	783: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_sequence_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(42)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_sequence_repeat1),
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
		Fstate: uint16(libc.Int32FromInt32(65)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__shebang_with_lang_repeat1),
	})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	792: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__shebang_with_lang_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(241)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(300)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(389)),
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
		Fstate: uint16(libc.Int32FromInt32(304)),
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
		Fstate: uint16(libc.Int32FromInt32(391)),
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
		Fcount: uint8(1),
	}})),
	803: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__raw_string_indented_repeat1),
	})))),
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
		Fcount: uint8(2),
	}})),
	805: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__raw_string_indented_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(247)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(291)),
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
		Fstate: uint16(libc.Int32FromInt32(340)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	812: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_attribute_repeat2),
		Fproduction_id: uint16(10),
	})))),
	813: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_attribute_repeat2),
		Fproduction_id: uint16(10),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(56)),
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
		Fstate: uint16(libc.Int32FromInt32(58)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_interpolation),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(376)),
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
		Fstate: uint16(libc.Int32FromInt32(324)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(278)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(247)),
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
		Fstate: uint16(libc.Int32FromInt32(31)),
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
		Fstate: uint16(libc.Int32FromInt32(45)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(290)),
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
		Fstate: uint16(libc.Int32FromInt32(285)),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_sequence),
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
		Fstate: uint16(libc.Int32FromInt32(33)),
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
		Fstate: uint16(libc.Int32FromInt32(35)),
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
		Fstate: uint16(libc.Int32FromInt32(379)),
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
		Fstate: uint16(libc.Int32FromInt32(317)),
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
		Fstate: uint16(libc.Int32FromInt32(392)),
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
		Fstate: uint16(libc.Int32FromInt32(341)),
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
		Fstate: uint16(libc.Int32FromInt32(241)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_attribute_kv_argument),
		Fproduction_id: uint16(23),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(372)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(97)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(95)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(aux_sym_attribute_repeat2),
		Fproduction_id: uint16(22),
	})))),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(321)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(226)),
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
		Fstate: uint16(libc.Int32FromInt32(76)),
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
		Fstate: uint16(libc.Int32FromInt32(363)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(aux_sym_attribute_repeat2),
		Fproduction_id: uint16(43),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(73)),
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
		Fstate: uint16(libc.Int32FromInt32(70)),
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
		Fstate: uint16(libc.Int32FromInt32(93)),
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
		Fstate: uint16(libc.Int32FromInt32(92)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(308)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(351)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(aux_sym_attribute_repeat2),
		Fproduction_id: uint16(22),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_recipe_line_prefix),
	})))),
	890: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	891: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	892: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	893: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	894: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(179)),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_condition),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_recipe_header),
		Fproduction_id: uint16(3),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym__shebang_with_lang),
		Fproduction_id: uint16(11),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(194)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(86)),
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
		Fstate: uint16(libc.Int32FromInt32(101)),
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
		Fstate: uint16(libc.Int32FromInt32(136)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(348)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym__shebang_with_lang),
		Fproduction_id: uint16(16),
	})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(369)),
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
		Fstate: uint16(libc.Int32FromInt32(180)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	951: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(libc.Int32FromInt32(227)),
	}})))),
	952: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	953: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	954: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	955: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	956: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	957: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	958: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	959: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	960: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	961: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	962: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(113)),
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
		Fstate: uint16(libc.Int32FromInt32(125)),
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
		Fstate: uint16(libc.Int32FromInt32(104)),
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
	969: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f uint8_t
		_ [7]byte
	}{f: uint8(TSParseActionTypeAccept)})),
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
		Fstate: uint16(libc.Int32FromInt32(253)),
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
		Fstate: uint16(libc.Int32FromInt32(398)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_recipe_header),
		Fproduction_id: uint16(1),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(8)),
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
		Fstate: uint16(libc.Int32FromInt32(15)),
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
		Fstate: uint16(libc.Int32FromInt32(75)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(19)),
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
		Fstate: uint16(libc.Int32FromInt32(41)),
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
		Fstate: uint16(libc.Int32FromInt32(115)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym__shebang_with_lang),
		Fproduction_id: uint16(24),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(105)),
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
		Fstate: uint16(libc.Int32FromInt32(334)),
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
		Fstate: uint16(libc.Int32FromInt32(396)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_recipe_header),
		Fproduction_id: uint16(7),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(168)),
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
		Fstate: uint16(libc.Int32FromInt32(373)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_recipe_header),
		Fproduction_id: uint16(2),
	})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_variadic_parameter),
		Fproduction_id: uint16(4),
	})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(263)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(134)),
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
		Fstate: uint16(libc.Int32FromInt32(270)),
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
		Fstate: uint16(libc.Int32FromInt32(271)),
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
		Fstate: uint16(libc.Int32FromInt32(203)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_shebang),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(305)),
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
		Fstate: uint16(libc.Int32FromInt32(176)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(229)),
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
		Fstate: uint16(libc.Int32FromInt32(175)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1041: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(87)),
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
		Fstate: uint16(libc.Int32FromInt32(120)),
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
		Fstate: uint16(libc.Int32FromInt32(98)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
}

type ts_external_scanner_symbol_identifiers = int32

const ts_external_token__indent = 0
const ts_external_token__dedent = 1
const ts_external_token__newline = 2
const ts_external_token_text = 3
const ts_external_token_error_recovery = 4

var ts_external_scanner_symbol_map = [5]TSSymbol{
	0: uint16(sym__indent),
	1: uint16(sym__dedent),
	2: uint16(sym__newline),
	3: uint16(sym_text),
	4: uint16(sym_error_recovery),
}

var ts_external_scanner_states = [7][5]uint8{
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
		3: libc.BoolUint8(true1 != 0),
	},
	3: {
		0: libc.BoolUint8(true1 != 0),
	},
	4: {
		2: libc.BoolUint8(true1 != 0),
	},
	5: {
		2: libc.BoolUint8(true1 != 0),
		3: libc.BoolUint8(true1 != 0),
	},
	6: {
		3: libc.BoolUint8(true1 != 0),
	},
}

func tree_sitter_just(tls *libc.TLS) (r uintptr) {
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
	Fname:              __ccgo_ts + 1534,
	Fmetadata: TSLanguageMetadata{
		Fminor_version: uint8(2),
	},
}

func init() {
	p := unsafe.Pointer(&language)
	*(*uintptr)(unsafe.Add(p, 144)) = __ccgo_fp(ts_lex)
	*(*uintptr)(unsafe.Add(p, 152)) = __ccgo_fp(ts_lex_keywords)
	*(*uintptr)(unsafe.Add(p, 184)) = __ccgo_fp(tree_sitter_just_external_scanner_create)
	*(*uintptr)(unsafe.Add(p, 192)) = __ccgo_fp(tree_sitter_just_external_scanner_destroy)
	*(*uintptr)(unsafe.Add(p, 200)) = __ccgo_fp(tree_sitter_just_external_scanner_scan)
	*(*uintptr)(unsafe.Add(p, 208)) = __ccgo_fp(tree_sitter_just_external_scanner_serialize)
	*(*uintptr)(unsafe.Add(p, 216)) = __ccgo_fp(tree_sitter_just_external_scanner_deserialize)
}

func __ccgo_fp(f interface{}) uintptr {
	type iface [2]uintptr
	return (*iface)(unsafe.Pointer(&f))[1]
}

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = "panic at %s:%d: \x00combined.c\x00invalid symbol %d\x00\n\x00tree_sitter_just_external_scanner_create: out of memory\x00got null payload at destroy\x00invalid scanner size\x00expected EOF\x00end\x00identifier\x00alias\x00:=\x00export\x00import\x00?\x00mod\x00set\x00[\x00,\x00]\x00shell\x00true\x00false\x00/\x00+\x00if\x00else\x00{\x00}\x00==\x00!=\x00=~\x00(\x00)\x00=\x00:\x00@\x00$\x00*\x00&&\x00@-\x00-@\x00-\x00shebang_token1\x00_shebang_with_lang_token1\x00env\x00_shebang_with_lang_token2\x00_shebang_with_lang_token3\x00_opaque_shebang\x00string_token1\x00'''\x00_raw_string_indented_token1\x00\"\x00_string_token1\x00\"\"\"\x00_string_indented_token1\x00escape_sequence\x00`\x00```\x00{{\x00}}\x00numeric_error\x00comment\x00_indent\x00_dedent\x00_newline\x00text\x00error_recovery\x00source_file\x00_item\x00assignment\x00module\x00setting\x00boolean\x00expression\x00_expression_inner\x00if_expression\x00else_if_clause\x00else_clause\x00_braced_expr\x00condition\x00regex_literal\x00value\x00function_call\x00external_command\x00sequence\x00attribute_kv_argument\x00attribute\x00recipe\x00recipe_header\x00parameters\x00parameter\x00variadic_parameter\x00dependencies\x00dependency\x00dependency_expression\x00recipe_body\x00recipe_line\x00recipe_line_prefix\x00shebang\x00_shebang_with_lang\x00string\x00_backticked\x00_indented_backticked\x00command_body\x00interpolation\x00source_file_repeat1\x00alias_repeat1\x00setting_repeat1\x00if_expression_repeat1\x00sequence_repeat1\x00attribute_repeat1\x00attribute_repeat2\x00parameters_repeat1\x00dependencies_repeat1\x00dependency_expression_repeat1\x00recipe_body_repeat1\x00recipe_line_repeat1\x00_shebang_with_lang_repeat1\x00_raw_string_indented_repeat1\x00_string_repeat1\x00_string_indented_repeat1\x00command_body_repeat1\x00language\x00alternative\x00argument\x00arguments\x00array\x00body\x00consequence\x00content\x00default\x00element\x00key\x00kleene\x00left\x00name\x00right\x00just\x00"
