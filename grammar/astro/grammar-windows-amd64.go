// Code generated for windows/amd64 by 'ccgo -ignore-unsupported-alignment -ignore-unsupported-atomic-sizes -ignore-vector-functions --winapi-no-errno -ignore-link-errors -std=gnu11 -O0 -D__extension__= -D__forceinline=static inline -D__attribute__(...)= -D__declspec(x)= -D__cdecl= -D__stdcall= -D__fastcall= -D__thiscall= -D_cdecl= -D__restrict= -D__restrict__= -D__MINGW_EXTENSION= -D_X86INTRIN_H_INCLUDED -D_X86GPRINTRIN_H_INCLUDED -D_IMMINTRIN_H_INCLUDED -D_MMINTRIN_H_INCLUDED -D_XMMINTRIN_H_INCLUDED -D_EMMINTRIN_H_INCLUDED -D_PMMINTRIN_H_INCLUDED -D_MM3DNOW_H_INCLUDED -Dfunc=func_token -Dinterface=interface_token -Dselect=select_token -Dchan=chan_token -Dgo=go_token -Dmap=map_token -Dpackage=package_token -Dtype=type_token -Dvar=var_token -Dimport=import_token -Ddefer=defer_token -Dfallthrough=fallthrough_token -Drange=range_token -I D:\a\ccgo-tree-sitter\ccgo-tree-sitter\third-party\tree-sitter-astro\src -I D:\a\ccgo-tree-sitter\ccgo-tree-sitter\third-party\tree-sitter-astro -I D:\a\ccgo-tree-sitter\ccgo-tree-sitter\third-party -I C:\Users\runneradmin\.cache\workspaced\sources\github\7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2\lib\include -I C:\Users\runneradmin\.cache\workspaced\sources\github\7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2\lib\src combined.c -o grammar.go', DO NOT EDIT.

//go:build windows && amd64

package grammar_astro

import (
	"reflect"
	"unsafe"

	"modernc.org/libc"
)

var _ reflect.Type
var _ unsafe.Pointer

const ALIAS_COUNT = 0
const BUFSIZ = 512
const E2BIG = 7
const EACCES = 13
const EADDRINUSE = 100
const EADDRNOTAVAIL = 101
const EAFNOSUPPORT = 102
const EAGAIN = 11
const EALREADY = 103
const EBADF = 9
const EBADMSG = 104
const EBUSY = 16
const ECANCELED = 105
const ECHILD = 10
const ECONNABORTED = 106
const ECONNREFUSED = 107
const ECONNRESET = 108
const EDEADLK = 36
const EDEADLOCK = "EDEADLK"
const EDESTADDRREQ = 109
const EDOM = 33
const EEXIST = 17
const EFAULT = 14
const EFBIG = 27
const EHOSTUNREACH = 110
const EIDRM = 111
const EILSEQ = 42
const EINPROGRESS = 112
const EINTR = 4
const EINVAL = 22
const EIO = 5
const EISCONN = 113
const EISDIR = 21
const ELOOP = 114
const EMFILE = 24
const EMLINK = 31
const EMSGSIZE = 115
const ENAMETOOLONG = 38
const ENETDOWN = 116
const ENETRESET = 117
const ENETUNREACH = 118
const ENFILE = 23
const ENOBUFS = 119
const ENODATA = 120
const ENODEV = 19
const ENOENT = 2
const ENOEXEC = 8
const ENOFILE = "ENOENT"
const ENOLCK = 39
const ENOLINK = 121
const ENOMEM = 12
const ENOMSG = 122
const ENOPROTOOPT = 123
const ENOSPC = 28
const ENOSR = 124
const ENOSTR = 125
const ENOSYS = 40
const ENOTCONN = 126
const ENOTDIR = 20
const ENOTEMPTY = 41
const ENOTRECOVERABLE = 127
const ENOTSOCK = 128
const ENOTSUP = 129
const ENOTTY = 25
const ENXIO = 6
const EOPNOTSUPP = 130
const EOVERFLOW = 132
const EOWNERDEAD = 133
const EPERM = 1
const EPIPE = 32
const EPROTO = 134
const EPROTONOSUPPORT = 135
const EPROTOTYPE = 136
const ERANGE = 34
const EROFS = 30
const ESPIPE = 29
const ESRCH = 3
const ETIME = 137
const ETIMEDOUT = 138
const ETXTBSY = 139
const EWOULDBLOCK = 140
const EXDEV = 18
const EXIT_FAILURE = 1
const EXIT_SUCCESS = 0
const EXTERNAL_TOKEN_COUNT = 16
const FIELD_COUNT = 0
const FILENAME_MAX = 260
const FOPEN_MAX = 20
const INT16_MAX = 32767
const INT32_MAX = 2147483647
const INT64_MAX = 9223372036854775807
const INT8_MAX = 127
const INTMAX_MAX = "INT64_MAX"
const INTMAX_MIN = "INT64_MIN"
const INTPTR_MAX = "INT64_MAX"
const INTPTR_MIN = "INT64_MIN"
const INT_FAST16_MAX = "INT16_MAX"
const INT_FAST16_MIN = "INT16_MIN"
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
const LANGUAGE_VERSION = 14
const LARGE_STATE_COUNT = 2
const L_tmpnam = 260
const L_tmpnam_s = "L_tmpnam"
const MAX_ALIAS_SEQUENCE_LENGTH = 4
const MB_LEN_MAX = 5
const MINGW_HAS_DDK_H = 1
const MINGW_HAS_SECURE_API = 1
const PATH_MAX = 260
const PRODUCTION_ID_COUNT = 1
const PTRDIFF_MAX = "INT64_MAX"
const PTRDIFF_MIN = "INT64_MIN"
const P_tmpdir = "_P_tmpdir"
const RAND_MAX = 0x7fff
const SEEK_CUR = 1
const SEEK_END = 2
const SEEK_SET = 0
const SIG_ATOMIC_MAX = "INT32_MAX"
const SIG_ATOMIC_MIN = "INT32_MIN"
const SIZE_MAX = "UINT64_MAX"
const SSIZE_MAX = "_I64_MAX"
const STATE_COUNT = 163
const STDERR_FILENO = 2
const STDIN_FILENO = 0
const STDOUT_FILENO = 1
const STRUNCATE = 80
const SYMBOL_COUNT = 59
const SYS_OPEN = "_SYS_OPEN"
const TMP_MAX = 2147483647
const TMP_MAX_S = "TMP_MAX"
const TOKEN_COUNT = 36
const TREE_SITTER_SERIALIZATION_BUFFER_SIZE = 1024
const UINT16_MAX = 65535
const UINT32_MAX = 0xffffffff
const UINT64_MAX = "0xffffffffffffffffU"
const UINT8_MAX = 255
const UINTMAX_MAX = "UINT64_MAX"
const UINTPTR_MAX = "UINT64_MAX"
const UINT_FAST16_MAX = "UINT16_MAX"
const UINT_FAST32_MAX = "UINT32_MAX"
const UINT_FAST64_MAX = "UINT64_MAX"
const UINT_FAST8_MAX = "UINT8_MAX"
const UINT_LEAST16_MAX = "UINT16_MAX"
const UINT_LEAST32_MAX = "UINT32_MAX"
const UINT_LEAST64_MAX = "UINT64_MAX"
const UINT_LEAST8_MAX = "UINT8_MAX"
const UNALIGNED = "__unaligned"
const USE___UUIDOF = 0
const WCHAR_MAX = 0xffff
const WCHAR_MIN = 0
const WIN32 = 1
const WIN64 = 1
const WINNT = 1
const WINT_MAX = 0xffff
const WINT_MIN = 0
const _ALLOCA_S_HEAP_MARKER = 56797
const _ALLOCA_S_MARKER_SIZE = 16
const _ALLOCA_S_STACK_MARKER = 0xCCCC
const _ALLOCA_S_THRESHOLD = 1024
const _ANONYMOUS_STRUCT = "__MINGW_EXTENSION"
const _ANONYMOUS_UNION = "__MINGW_EXTENSION"
const _ARGMAX = 100
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
const _DIGIT = 4
const _EMMINTRIN_H_INCLUDED = 1
const _FREEENTRY = 0
const _HEAP_MAXREQ = 0xFFFFFFFFFFFFFFE0
const _HEX = 0x80
const _I16_MAX = 32767
const _I32_MAX = 2147483647
const _I64_MAX = "9223372036854775807ll"
const _I8_MAX = 127
const _IMMINTRIN_H_INCLUDED = 1
const _INTEGRAL_MAX_BITS = 64
const _IOB_ENTRIES = 20
const _IOFBF = 0x0000
const _IOLBF = 0x0040
const _IONBF = 0x0004
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
const _M_AMD64 = 100
const _M_X64 = 100
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
const _REENTRANT = 1
const _REPORT_ERRMODE = 3
const _SECURECRT_FILL_BUFFER_PATTERN = 0xFD
const _SPACE = 8
const _SYS_OPEN = 20
const _UI16_MAX = "0xffffu"
const _UI32_MAX = "0xffffffffu"
const _UI64_MAX = "0xffffffffffffffffull"
const _UI8_MAX = "0xffu"
const _UPPER = 1
const _USEDENTRY = 1
const _WAIT_CHILD = 0
const _WAIT_GRANDCHILD = 1
const _WConst_return = "_CONST_RETURN"
const _WIN32 = 1
const _WIN32_WINNT = 0xa00
const _WIN64 = 1
const _WRITE_ABORT_MSG = 0x1
const _X86GPRINTRIN_H_INCLUDED = 1
const _X86INTRIN_H_INCLUDED = 1
const _XMMINTRIN_H_INCLUDED = 1
const __ATOMIC_ACQUIRE = 2
const __ATOMIC_ACQ_REL = 4
const __ATOMIC_CONSUME = 1
const __ATOMIC_HLE_ACQUIRE = 65536
const __ATOMIC_HLE_RELEASE = 131072
const __ATOMIC_RELAXED = 0
const __ATOMIC_RELEASE = 3
const __ATOMIC_SEQ_CST = 5
const __BFLT16_DECIMAL_DIG__ = 4
const __BFLT16_DENORM_MIN__ = "9.18354961579912115600575419704879436e-41B"
const __BFLT16_DIG__ = 2
const __BFLT16_EPSILON__ = "7.81250000000000000000000000000000000e-3B"
const __BFLT16_HAS_DENORM__ = 1
const __BFLT16_HAS_INFINITY__ = 1
const __BFLT16_HAS_QUIET_NAN__ = 1
const __BFLT16_IS_IEC_60559__ = 0
const __BFLT16_MANT_DIG__ = 8
const __BFLT16_MAX_10_EXP__ = 38
const __BFLT16_MAX_EXP__ = 128
const __BFLT16_MAX__ = "3.38953138925153547590470800371487867e+38B"
const __BFLT16_MIN__ = "1.17549435082228750796873653722224568e-38B"
const __BFLT16_NORM_MAX__ = "3.38953138925153547590470800371487867e+38B"
const __BIGGEST_ALIGNMENT__ = 16
const __BITINT_MAXWIDTH__ = 65535
const __BYTE_ORDER__ = "__ORDER_LITTLE_ENDIAN__"
const __C89_NAMELESS = "__MINGW_EXTENSION"
const __CCGO__ = 1
const __CHAR_BIT__ = 8
const __CRTDECL = "__cdecl"
const __DBL_DECIMAL_DIG__ = 17
const __DBL_DIG__ = 15
const __DBL_HAS_DENORM__ = 1
const __DBL_HAS_INFINITY__ = 1
const __DBL_HAS_QUIET_NAN__ = 1
const __DBL_IS_IEC_60559__ = 1
const __DBL_MANT_DIG__ = 53
const __DBL_MAX_10_EXP__ = 308
const __DBL_MAX_EXP__ = 1024
const __DEC128_EPSILON__ = 1e-33
const __DEC128_MANT_DIG__ = 34
const __DEC128_MAX_EXP__ = 6145
const __DEC128_MAX__ = "9.999999999999999999999999999999999E6144"
const __DEC128_MIN__ = 1e-6143
const __DEC128_SUBNORMAL_MIN__ = 0.000000000000000000000000000000001e-6143
const __DEC32_EPSILON__ = 1e-6
const __DEC32_MANT_DIG__ = 7
const __DEC32_MAX_EXP__ = 97
const __DEC32_MAX__ = 9.999999e96
const __DEC32_MIN__ = 1e-95
const __DEC32_SUBNORMAL_MIN__ = 0.000001e-95
const __DEC64X_EPSILON__ = "1E-33D64x"
const __DEC64X_MANT_DIG__ = 34
const __DEC64X_MAX_EXP__ = 6145
const __DEC64X_MAX__ = "9.999999999999999999999999999999999E6144D64x"
const __DEC64X_MIN__ = "1E-6143D64x"
const __DEC64X_SUBNORMAL_MIN__ = "0.000000000000000000000000000000001E-6143D64x"
const __DEC64_EPSILON__ = 1e-15
const __DEC64_MANT_DIG__ = 16
const __DEC64_MAX_EXP__ = 385
const __DEC64_MAX__ = "9.999999999999999E384"
const __DEC64_MIN__ = 1e-383
const __DEC64_SUBNORMAL_MIN__ = 0.000000000000001e-383
const __DECIMAL_BID_FORMAT__ = 1
const __DECIMAL_DIG__ = 17
const __DEC_EVAL_METHOD__ = 2
const __FINITE_MATH_ONLY__ = 0
const __FLOAT_WORD_ORDER__ = "__ORDER_LITTLE_ENDIAN__"
const __FLT128_DECIMAL_DIG__ = 36
const __FLT128_DENORM_MIN__ = 6.47517511943802511092443895822764655e-4966
const __FLT128_DIG__ = 33
const __FLT128_EPSILON__ = 1.92592994438723585305597794258492732e-34
const __FLT128_HAS_DENORM__ = 1
const __FLT128_HAS_INFINITY__ = 1
const __FLT128_HAS_QUIET_NAN__ = 1
const __FLT128_IS_IEC_60559__ = 1
const __FLT128_MANT_DIG__ = 113
const __FLT128_MAX_10_EXP__ = 4932
const __FLT128_MAX_EXP__ = 16384
const __FLT128_MAX__ = "1.18973149535723176508575932662800702e+4932"
const __FLT128_MIN__ = 3.36210314311209350626267781732175260e-4932
const __FLT128_NORM_MAX__ = "1.18973149535723176508575932662800702e+4932"
const __FLT16_DECIMAL_DIG__ = 5
const __FLT16_DENORM_MIN__ = 5.96046447753906250000000000000000000e-8
const __FLT16_DIG__ = 3
const __FLT16_EPSILON__ = 9.76562500000000000000000000000000000e-4
const __FLT16_HAS_DENORM__ = 1
const __FLT16_HAS_INFINITY__ = 1
const __FLT16_HAS_QUIET_NAN__ = 1
const __FLT16_IS_IEC_60559__ = 1
const __FLT16_MANT_DIG__ = 11
const __FLT16_MAX_10_EXP__ = 4
const __FLT16_MAX_EXP__ = 16
const __FLT16_MAX__ = 6.55040000000000000000000000000000000e+4
const __FLT16_MIN__ = 6.10351562500000000000000000000000000e-5
const __FLT16_NORM_MAX__ = 6.55040000000000000000000000000000000e+4
const __FLT32X_DECIMAL_DIG__ = 17
const __FLT32X_DENORM_MIN__ = 4.94065645841246544176568792868221372e-324
const __FLT32X_DIG__ = 15
const __FLT32X_EPSILON__ = 2.22044604925031308084726333618164062e-16
const __FLT32X_HAS_DENORM__ = 1
const __FLT32X_HAS_INFINITY__ = 1
const __FLT32X_HAS_QUIET_NAN__ = 1
const __FLT32X_IS_IEC_60559__ = 1
const __FLT32X_MANT_DIG__ = 53
const __FLT32X_MAX_10_EXP__ = 308
const __FLT32X_MAX_EXP__ = 1024
const __FLT32X_MAX__ = 1.79769313486231570814527423731704357e+308
const __FLT32X_MIN__ = 2.22507385850720138309023271733240406e-308
const __FLT32X_NORM_MAX__ = 1.79769313486231570814527423731704357e+308
const __FLT32_DECIMAL_DIG__ = 9
const __FLT32_DENORM_MIN__ = 1.40129846432481707092372958328991613e-45
const __FLT32_DIG__ = 6
const __FLT32_EPSILON__ = 1.19209289550781250000000000000000000e-7
const __FLT32_HAS_DENORM__ = 1
const __FLT32_HAS_INFINITY__ = 1
const __FLT32_HAS_QUIET_NAN__ = 1
const __FLT32_IS_IEC_60559__ = 1
const __FLT32_MANT_DIG__ = 24
const __FLT32_MAX_10_EXP__ = 38
const __FLT32_MAX_EXP__ = 128
const __FLT32_MAX__ = 3.40282346638528859811704183484516925e+38
const __FLT32_MIN__ = 1.17549435082228750796873653722224568e-38
const __FLT32_NORM_MAX__ = 3.40282346638528859811704183484516925e+38
const __FLT64X_DECIMAL_DIG__ = 36
const __FLT64X_DENORM_MIN__ = 6.47517511943802511092443895822764655e-4966
const __FLT64X_DIG__ = 33
const __FLT64X_EPSILON__ = 1.92592994438723585305597794258492732e-34
const __FLT64X_HAS_DENORM__ = 1
const __FLT64X_HAS_INFINITY__ = 1
const __FLT64X_HAS_QUIET_NAN__ = 1
const __FLT64X_IS_IEC_60559__ = 1
const __FLT64X_MANT_DIG__ = 113
const __FLT64X_MAX_10_EXP__ = 4932
const __FLT64X_MAX_EXP__ = 16384
const __FLT64X_MAX__ = "1.18973149535723176508575932662800702e+4932"
const __FLT64X_MIN__ = 3.36210314311209350626267781732175260e-4932
const __FLT64X_NORM_MAX__ = "1.18973149535723176508575932662800702e+4932"
const __FLT64_DECIMAL_DIG__ = 17
const __FLT64_DENORM_MIN__ = 4.94065645841246544176568792868221372e-324
const __FLT64_DIG__ = 15
const __FLT64_EPSILON__ = 2.22044604925031308084726333618164062e-16
const __FLT64_HAS_DENORM__ = 1
const __FLT64_HAS_INFINITY__ = 1
const __FLT64_HAS_QUIET_NAN__ = 1
const __FLT64_IS_IEC_60559__ = 1
const __FLT64_MANT_DIG__ = 53
const __FLT64_MAX_10_EXP__ = 308
const __FLT64_MAX_EXP__ = 1024
const __FLT64_MAX__ = 1.79769313486231570814527423731704357e+308
const __FLT64_MIN__ = 2.22507385850720138309023271733240406e-308
const __FLT64_NORM_MAX__ = 1.79769313486231570814527423731704357e+308
const __FLT_DECIMAL_DIG__ = 9
const __FLT_DENORM_MIN__ = 1.40129846432481707092372958328991613e-45
const __FLT_DIG__ = 6
const __FLT_EPSILON__ = 1.19209289550781250000000000000000000e-7
const __FLT_EVAL_METHOD_TS_18661_3__ = 2
const __FLT_EVAL_METHOD__ = 2
const __FLT_HAS_DENORM__ = 1
const __FLT_HAS_INFINITY__ = 1
const __FLT_HAS_QUIET_NAN__ = 1
const __FLT_IS_IEC_60559__ = 1
const __FLT_MANT_DIG__ = 24
const __FLT_MAX_10_EXP__ = 38
const __FLT_MAX_EXP__ = 128
const __FLT_MAX__ = 3.40282346638528859811704183484516925e+38
const __FLT_MIN__ = 1.17549435082228750796873653722224568e-38
const __FLT_NORM_MAX__ = 3.40282346638528859811704183484516925e+38
const __FLT_RADIX__ = 2
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
const __GCC_CONSTRUCTIVE_SIZE = 64
const __GCC_DESTRUCTIVE_SIZE = 64
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_1 = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_16 = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_2 = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_4 = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_8 = 1
const __GCC_IEC_559 = 2
const __GCC_IEC_559_COMPLEX = 2
const __GNUC_EXECUTION_CHARSET_NAME = "UTF-8"
const __GNUC_MINOR__ = 2
const __GNUC_PATCHLEVEL__ = 0
const __GNUC_STDC_INLINE__ = 1
const __GNUC_WIDE_EXECUTION_CHARSET_NAME = "UTF-16LE"
const __GNUC__ = 15
const __GNU_EXTENSION = "__MINGW_EXTENSION"
const __GOT_SECURE_LIB__ = "__STDC_SECURE_LIB__"
const __GXX_ABI_VERSION = 1020
const __GXX_MERGED_TYPEINFO_NAMES = 0
const __GXX_TYPEINFO_EQUALITY_INLINE = 0
const __HAVE_SPECULATION_SAFE_VALUE = 1
const __INT16_MAX__ = 0x7fff
const __INT32_MAX__ = 0x7fffffff
const __INT32_TYPE__ = "int"
const __INT64_MAX__ = 0x7fffffffffffffff
const __INT8_MAX__ = 0x7f
const __INTMAX_MAX__ = 0x7fffffffffffffff
const __INTMAX_WIDTH__ = 64
const __INTPTR_MAX__ = 0x7fffffffffffffff
const __INTPTR_WIDTH__ = 64
const __INT_FAST16_MAX__ = 0x7fff
const __INT_FAST16_WIDTH__ = 16
const __INT_FAST32_MAX__ = 0x7fffffff
const __INT_FAST32_TYPE__ = "int"
const __INT_FAST32_WIDTH__ = 32
const __INT_FAST64_MAX__ = 0x7fffffffffffffff
const __INT_FAST64_WIDTH__ = 64
const __INT_FAST8_MAX__ = 0x7f
const __INT_FAST8_WIDTH__ = 8
const __INT_LEAST16_MAX__ = 0x7fff
const __INT_LEAST16_WIDTH__ = 16
const __INT_LEAST32_MAX__ = 0x7fffffff
const __INT_LEAST32_TYPE__ = "int"
const __INT_LEAST32_WIDTH__ = 32
const __INT_LEAST64_MAX__ = 0x7fffffffffffffff
const __INT_LEAST64_WIDTH__ = 64
const __INT_LEAST8_MAX__ = 0x7f
const __INT_LEAST8_WIDTH__ = 8
const __INT_MAX__ = 0x7fffffff
const __INT_WIDTH__ = 32
const __LDBL_DECIMAL_DIG__ = 17
const __LDBL_DENORM_MIN__ = 4.94065645841246544176568792868221372e-324
const __LDBL_DIG__ = 15
const __LDBL_EPSILON__ = 2.22044604925031308084726333618164062e-16
const __LDBL_HAS_DENORM__ = 1
const __LDBL_HAS_INFINITY__ = 1
const __LDBL_HAS_QUIET_NAN__ = 1
const __LDBL_IS_IEC_60559__ = 1
const __LDBL_MANT_DIG__ = 53
const __LDBL_MAX_10_EXP__ = 308
const __LDBL_MAX_EXP__ = 1024
const __LDBL_MAX__ = 1.79769313486231570814527423731704357e+308
const __LDBL_MIN__ = 2.22507385850720138309023271733240406e-308
const __LDBL_NORM_MAX__ = 1.79769313486231570814527423731704357e+308
const __LONG32 = "long"
const __LONG_DOUBLE_64__ = 1
const __LONG_LONG_MAX__ = 0x7fffffffffffffff
const __LONG_LONG_WIDTH__ = 64
const __LONG_MAX__ = 0x7fffffff
const __LONG_WIDTH__ = 32
const __MINGW32_MAJOR_VERSION = 3
const __MINGW32_MINOR_VERSION = 11
const __MINGW32__ = 1
const __MINGW64_VERSION_BUGFIX = 0
const __MINGW64_VERSION_MAJOR = 13
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
const __ORDER_BIG_ENDIAN__ = 4321
const __ORDER_LITTLE_ENDIAN__ = 1234
const __ORDER_PDP_ENDIAN__ = 3412
const __PIC__ = 1
const __PRAGMA_REDEFINE_EXTNAME = 1
const __PRETTY_FUNCTION__ = "__func__"
const __PTRDIFF_MAX__ = 0x7fffffffffffffff
const __PTRDIFF_WIDTH__ = 64
const __SCHAR_MAX__ = 0x7f
const __SCHAR_WIDTH__ = 8
const __SEG_FS = 1
const __SEG_GS = 1
const __SEH__ = 1
const __SHRT_MAX__ = 0x7fff
const __SHRT_WIDTH__ = 16
const __SIG_ATOMIC_MAX__ = 0x7fffffff
const __SIG_ATOMIC_TYPE__ = "int"
const __SIG_ATOMIC_WIDTH__ = 32
const __SIZEOF_DOUBLE__ = 8
const __SIZEOF_FLOAT128__ = 16
const __SIZEOF_FLOAT80__ = 16
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
const __SIZE_MAX__ = "0xffffffffffffffffU"
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
const __UINT16_MAX__ = 0xffff
const __UINT32_MAX__ = 0xffffffff
const __UINT64_MAX__ = "0xffffffffffffffffU"
const __UINT8_MAX__ = 0xff
const __UINTMAX_MAX__ = "0xffffffffffffffffU"
const __UINTPTR_MAX__ = "0xffffffffffffffffU"
const __UINT_FAST16_MAX__ = 0xffff
const __UINT_FAST32_MAX__ = 0xffffffff
const __UINT_FAST64_MAX__ = "0xffffffffffffffffU"
const __UINT_FAST8_MAX__ = 0xff
const __UINT_LEAST16_MAX__ = 0xffff
const __UINT_LEAST32_MAX__ = 0xffffffff
const __UINT_LEAST64_MAX__ = "0xffffffffffffffffU"
const __UINT_LEAST8_MAX__ = 0xff
const __USE_MINGW_ANSI_STDIO = 0
const __USING_POSIXTHREAD__ = 1
const __VERSION__ = "15.2.0"
const __WCHAR_MAX__ = 0xffff
const __WCHAR_MIN__ = 0
const __WCHAR_WIDTH__ = 16
const __WIN32 = 1
const __WIN32__ = 1
const __WIN64 = 1
const __WIN64__ = 1
const __WINNT = 1
const __WINNT__ = 1
const __WINT_MAX__ = 0xffff
const __WINT_MIN__ = 0
const __WINT_WIDTH__ = 16
const __amd64 = 1
const __amd64__ = 1
const __bool_true_false_are_defined = 1
const __code_model_medium__ = 1
const __int16 = "short"
const __int32 = "int"
const __int8 = "char"
const __mingw_bos_ovr = "__mingw_ovr"
const __nocona = 1
const __nocona__ = 1
const __pic__ = 1
const __tune_core2__ = 1
const __x86_64 = 1
const __x86_64__ = 1
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
const strcasecmp1 = "_stricmp"
const strncasecmp = "_strnicmp"
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
const wcswcs = "wcsstr"
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

type max_align_t = struct {
	F__max_align_ll int64
	F__max_align_ld float64
}

type int8_t = int8

type uint8_t = uint8

type int16_t = int16

type uint16_t = uint16

type int32_t = int32

type uint32_t = uint32

type int64_t = int64

type uint64_t = uint64

type int_least8_t = int8

type uint_least8_t = uint8

type int_least16_t = int16

type uint_least16_t = uint16

type int_least32_t = int32

type uint_least32_t = uint32

type int_least64_t = int64

type uint_least64_t = uint64

type int_fast8_t = int8

type uint_fast8_t = uint8

type int_fast16_t = int16

type uint_fast16_t = uint16

type int_fast32_t = int32

type uint_fast32_t = uint32

type int_fast64_t = int64

type uint_fast64_t = uint64

type intmax_t = int64

type uintmax_t = uint64

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

type Array = struct {
	Fcontents uintptr
	Fsize     uint32_t
	Fcapacity uint32_t
}

type TagType = int32

const AREA = 0
const BASE = 1
const BR = 2
const COL = 3
const COMMAND = 4
const EMBED = 5
const FRAME = 6
const HR = 7
const IMAGE = 8
const IMG = 9
const INPUT = 10
const ISINDEX = 11
const KEYGEN = 12
const LINK = 13
const MENUITEM = 14
const META = 15
const NEXTID = 16
const PARAM = 17
const SOURCE = 18
const TRACK = 19
const WBR = 20
const END_OF_VOID_TAGS = 21
const A = 22
const ABBR = 23
const ADDRESS = 24
const ARTICLE = 25
const ASIDE = 26
const AUDIO = 27
const B = 28
const BDI = 29
const BDO = 30
const BLOCKQUOTE = 31
const BODY = 32
const BUTTON = 33
const CANVAS = 34
const CAPTION = 35
const CITE = 36
const CODE = 37
const COLGROUP = 38
const DATA = 39
const DATALIST = 40
const DD = 41
const DEL = 42
const DETAILS = 43
const DFN = 44
const DIALOG = 45
const DIV = 46
const DL = 47
const DT = 48
const EM = 49
const FIELDSET = 50
const FIGCAPTION = 51
const FIGURE = 52
const FOOTER = 53
const FORM = 54
const H1 = 55
const H2 = 56
const H3 = 57
const H4 = 58
const H5 = 59
const H6 = 60
const HEAD = 61
const HEADER = 62
const HGROUP = 63
const HTML = 64
const I = 65
const IFRAME = 66
const INS = 67
const KBD = 68
const LABEL = 69
const LEGEND = 70
const LI = 71
const MAIN = 72
const MAP = 73
const MARK = 74
const MATH = 75
const MENU = 76
const METER = 77
const NAV = 78
const NOSCRIPT = 79
const OBJECT = 80
const OL = 81
const OPTGROUP = 82
const OPTION = 83
const OUTPUT = 84
const P = 85
const PICTURE = 86
const PRE = 87
const PROGRESS = 88
const Q = 89
const RB = 90
const RP = 91
const RT = 92
const RTC = 93
const RUBY = 94
const S = 95
const SAMP = 96
const SCRIPT = 97
const SECTION = 98
const SELECT = 99
const SLOT = 100
const SMALL = 101
const SPAN = 102
const STRONG = 103
const STYLE = 104
const SUB = 105
const SUMMARY = 106
const SUP = 107
const SVG = 108
const TABLE = 109
const TBODY = 110
const TD = 111
const TEMPLATE = 112
const TEXTAREA = 113
const TFOOT = 114
const TH = 115
const THEAD = 116
const TIME = 117
const TITLE = 118
const TR = 119
const U = 120
const UL = 121
const VAR = 122
const VIDEO = 123
const INTERPOLATION = 124
const FRAGMENT = 125
const CUSTOM = 126
const END_ = 127

type String = struct {
	Fcontents uintptr
	Fsize     uint32_t
	Fcapacity uint32_t
}

type TagMapEntry = struct {
	Ftag_name [16]int8
	Ftag_type TagType
}

type Tag = struct {
	Ftype_token      TagType
	Fcustom_tag_name String
}

var TAG_TYPES_BY_TAG_NAME = [126]TagMapEntry{
	0: {
		Ftag_name: [16]int8{'a', 'r', 'e', 'a'},
	},
	1: {
		Ftag_name: [16]int8{'b', 'a', 's', 'e'},
		Ftag_type: int32(BASE),
	},
	2: {
		Ftag_name: [16]int8{'b', 'r'},
		Ftag_type: int32(BR),
	},
	3: {
		Ftag_name: [16]int8{'c', 'o', 'l'},
		Ftag_type: int32(COL),
	},
	4: {
		Ftag_name: [16]int8{'c', 'o', 'm', 'm', 'a', 'n', 'd'},
		Ftag_type: int32(COMMAND),
	},
	5: {
		Ftag_name: [16]int8{'e', 'm', 'b', 'e', 'd'},
		Ftag_type: int32(EMBED),
	},
	6: {
		Ftag_name: [16]int8{'f', 'r', 'a', 'm', 'e'},
		Ftag_type: int32(FRAME),
	},
	7: {
		Ftag_name: [16]int8{'h', 'r'},
		Ftag_type: int32(HR),
	},
	8: {
		Ftag_name: [16]int8{'i', 'm', 'a', 'g', 'e'},
		Ftag_type: int32(IMAGE),
	},
	9: {
		Ftag_name: [16]int8{'i', 'm', 'g'},
		Ftag_type: int32(IMG),
	},
	10: {
		Ftag_name: [16]int8{'i', 'n', 'p', 'u', 't'},
		Ftag_type: int32(INPUT),
	},
	11: {
		Ftag_name: [16]int8{'i', 's', 'i', 'n', 'd', 'e', 'x'},
		Ftag_type: int32(ISINDEX),
	},
	12: {
		Ftag_name: [16]int8{'k', 'e', 'y', 'g', 'e', 'n'},
		Ftag_type: int32(KEYGEN),
	},
	13: {
		Ftag_name: [16]int8{'l', 'i', 'n', 'k'},
		Ftag_type: int32(LINK),
	},
	14: {
		Ftag_name: [16]int8{'m', 'e', 'n', 'u', 'i', 't', 'e', 'm'},
		Ftag_type: int32(MENUITEM),
	},
	15: {
		Ftag_name: [16]int8{'m', 'e', 't', 'a'},
		Ftag_type: int32(META),
	},
	16: {
		Ftag_name: [16]int8{'n', 'e', 'x', 't', 'i', 'd'},
		Ftag_type: int32(NEXTID),
	},
	17: {
		Ftag_name: [16]int8{'p', 'a', 'r', 'a', 'm'},
		Ftag_type: int32(PARAM),
	},
	18: {
		Ftag_name: [16]int8{'s', 'o', 'u', 'r', 'c', 'e'},
		Ftag_type: int32(SOURCE),
	},
	19: {
		Ftag_name: [16]int8{'t', 'r', 'a', 'c', 'k'},
		Ftag_type: int32(TRACK),
	},
	20: {
		Ftag_name: [16]int8{'w', 'b', 'r'},
		Ftag_type: int32(WBR),
	},
	21: {
		Ftag_name: [16]int8{'a'},
		Ftag_type: int32(A),
	},
	22: {
		Ftag_name: [16]int8{'a', 'b', 'b', 'r'},
		Ftag_type: int32(ABBR),
	},
	23: {
		Ftag_name: [16]int8{'a', 'd', 'd', 'r', 'e', 's', 's'},
		Ftag_type: int32(ADDRESS),
	},
	24: {
		Ftag_name: [16]int8{'a', 'r', 't', 'i', 'c', 'l', 'e'},
		Ftag_type: int32(ARTICLE),
	},
	25: {
		Ftag_name: [16]int8{'a', 's', 'i', 'd', 'e'},
		Ftag_type: int32(ASIDE),
	},
	26: {
		Ftag_name: [16]int8{'a', 'u', 'd', 'i', 'o'},
		Ftag_type: int32(AUDIO),
	},
	27: {
		Ftag_name: [16]int8{'b'},
		Ftag_type: int32(B),
	},
	28: {
		Ftag_name: [16]int8{'b', 'd', 'i'},
		Ftag_type: int32(BDI),
	},
	29: {
		Ftag_name: [16]int8{'b', 'd', 'o'},
		Ftag_type: int32(BDO),
	},
	30: {
		Ftag_name: [16]int8{'b', 'l', 'o', 'c', 'k', 'q', 'u', 'o', 't', 'e'},
		Ftag_type: int32(BLOCKQUOTE),
	},
	31: {
		Ftag_name: [16]int8{'b', 'o', 'd', 'y'},
		Ftag_type: int32(BODY),
	},
	32: {
		Ftag_name: [16]int8{'b', 'u', 't', 't', 'o', 'n'},
		Ftag_type: int32(BUTTON),
	},
	33: {
		Ftag_name: [16]int8{'c', 'a', 'n', 'v', 'a', 's'},
		Ftag_type: int32(CANVAS),
	},
	34: {
		Ftag_name: [16]int8{'c', 'a', 'p', 't', 'i', 'o', 'n'},
		Ftag_type: int32(CAPTION),
	},
	35: {
		Ftag_name: [16]int8{'c', 'i', 't', 'e'},
		Ftag_type: int32(CITE),
	},
	36: {
		Ftag_name: [16]int8{'c', 'o', 'd', 'e'},
		Ftag_type: int32(CODE),
	},
	37: {
		Ftag_name: [16]int8{'c', 'o', 'l', 'g', 'r', 'o', 'u', 'p'},
		Ftag_type: int32(COLGROUP),
	},
	38: {
		Ftag_name: [16]int8{'d', 'a', 't', 'a'},
		Ftag_type: int32(DATA),
	},
	39: {
		Ftag_name: [16]int8{'d', 'a', 't', 'a', 'l', 'i', 's', 't'},
		Ftag_type: int32(DATALIST),
	},
	40: {
		Ftag_name: [16]int8{'d', 'd'},
		Ftag_type: int32(DD),
	},
	41: {
		Ftag_name: [16]int8{'d', 'e', 'l'},
		Ftag_type: int32(DEL),
	},
	42: {
		Ftag_name: [16]int8{'d', 'e', 't', 'a', 'i', 'l', 's'},
		Ftag_type: int32(DETAILS),
	},
	43: {
		Ftag_name: [16]int8{'d', 'f', 'n'},
		Ftag_type: int32(DFN),
	},
	44: {
		Ftag_name: [16]int8{'d', 'i', 'a', 'l', 'o', 'g'},
		Ftag_type: int32(DIALOG),
	},
	45: {
		Ftag_name: [16]int8{'d', 'i', 'v'},
		Ftag_type: int32(DIV),
	},
	46: {
		Ftag_name: [16]int8{'d', 'l'},
		Ftag_type: int32(DL),
	},
	47: {
		Ftag_name: [16]int8{'d', 't'},
		Ftag_type: int32(DT),
	},
	48: {
		Ftag_name: [16]int8{'e', 'm'},
		Ftag_type: int32(EM),
	},
	49: {
		Ftag_name: [16]int8{'f', 'i', 'e', 'l', 'd', 's', 'e', 't'},
		Ftag_type: int32(FIELDSET),
	},
	50: {
		Ftag_name: [16]int8{'f', 'i', 'g', 'c', 'a', 'p', 't', 'i', 'o', 'n'},
		Ftag_type: int32(FIGCAPTION),
	},
	51: {
		Ftag_name: [16]int8{'f', 'i', 'g', 'u', 'r', 'e'},
		Ftag_type: int32(FIGURE),
	},
	52: {
		Ftag_name: [16]int8{'f', 'o', 'o', 't', 'e', 'r'},
		Ftag_type: int32(FOOTER),
	},
	53: {
		Ftag_name: [16]int8{'f', 'o', 'r', 'm'},
		Ftag_type: int32(FORM),
	},
	54: {
		Ftag_name: [16]int8{'h', '1'},
		Ftag_type: int32(H1),
	},
	55: {
		Ftag_name: [16]int8{'h', '2'},
		Ftag_type: int32(H2),
	},
	56: {
		Ftag_name: [16]int8{'h', '3'},
		Ftag_type: int32(H3),
	},
	57: {
		Ftag_name: [16]int8{'h', '4'},
		Ftag_type: int32(H4),
	},
	58: {
		Ftag_name: [16]int8{'h', '5'},
		Ftag_type: int32(H5),
	},
	59: {
		Ftag_name: [16]int8{'h', '6'},
		Ftag_type: int32(H6),
	},
	60: {
		Ftag_name: [16]int8{'h', 'e', 'a', 'd'},
		Ftag_type: int32(HEAD),
	},
	61: {
		Ftag_name: [16]int8{'h', 'e', 'a', 'd', 'e', 'r'},
		Ftag_type: int32(HEADER),
	},
	62: {
		Ftag_name: [16]int8{'h', 'g', 'r', 'o', 'u', 'p'},
		Ftag_type: int32(HGROUP),
	},
	63: {
		Ftag_name: [16]int8{'h', 't', 'm', 'l'},
		Ftag_type: int32(HTML),
	},
	64: {
		Ftag_name: [16]int8{'i'},
		Ftag_type: int32(I),
	},
	65: {
		Ftag_name: [16]int8{'i', 'f', 'r', 'a', 'm', 'e'},
		Ftag_type: int32(IFRAME),
	},
	66: {
		Ftag_name: [16]int8{'i', 'n', 's'},
		Ftag_type: int32(INS),
	},
	67: {
		Ftag_name: [16]int8{'k', 'b', 'd'},
		Ftag_type: int32(KBD),
	},
	68: {
		Ftag_name: [16]int8{'l', 'a', 'b', 'e', 'l'},
		Ftag_type: int32(LABEL),
	},
	69: {
		Ftag_name: [16]int8{'l', 'e', 'g', 'e', 'n', 'd'},
		Ftag_type: int32(LEGEND),
	},
	70: {
		Ftag_name: [16]int8{'l', 'i'},
		Ftag_type: int32(LI),
	},
	71: {
		Ftag_name: [16]int8{'m', 'a', 'i', 'n'},
		Ftag_type: int32(MAIN),
	},
	72: {
		Ftag_name: [16]int8{'m', 'a', 'p'},
		Ftag_type: int32(MAP),
	},
	73: {
		Ftag_name: [16]int8{'m', 'a', 'r', 'k'},
		Ftag_type: int32(MARK),
	},
	74: {
		Ftag_name: [16]int8{'m', 'a', 't', 'h'},
		Ftag_type: int32(MATH),
	},
	75: {
		Ftag_name: [16]int8{'m', 'e', 'n', 'u'},
		Ftag_type: int32(MENU),
	},
	76: {
		Ftag_name: [16]int8{'m', 'e', 't', 'e', 'r'},
		Ftag_type: int32(METER),
	},
	77: {
		Ftag_name: [16]int8{'n', 'a', 'v'},
		Ftag_type: int32(NAV),
	},
	78: {
		Ftag_name: [16]int8{'n', 'o', 's', 'c', 'r', 'i', 'p', 't'},
		Ftag_type: int32(NOSCRIPT),
	},
	79: {
		Ftag_name: [16]int8{'o', 'b', 'j', 'e', 'c', 't'},
		Ftag_type: int32(OBJECT),
	},
	80: {
		Ftag_name: [16]int8{'o', 'l'},
		Ftag_type: int32(OL),
	},
	81: {
		Ftag_name: [16]int8{'o', 'p', 't', 'g', 'r', 'o', 'u', 'p'},
		Ftag_type: int32(OPTGROUP),
	},
	82: {
		Ftag_name: [16]int8{'o', 'p', 't', 'i', 'o', 'n'},
		Ftag_type: int32(OPTION),
	},
	83: {
		Ftag_name: [16]int8{'o', 'u', 't', 'p', 'u', 't'},
		Ftag_type: int32(OUTPUT),
	},
	84: {
		Ftag_name: [16]int8{'p'},
		Ftag_type: int32(P),
	},
	85: {
		Ftag_name: [16]int8{'p', 'i', 'c', 't', 'u', 'r', 'e'},
		Ftag_type: int32(PICTURE),
	},
	86: {
		Ftag_name: [16]int8{'p', 'r', 'e'},
		Ftag_type: int32(PRE),
	},
	87: {
		Ftag_name: [16]int8{'p', 'r', 'o', 'g', 'r', 'e', 's', 's'},
		Ftag_type: int32(PROGRESS),
	},
	88: {
		Ftag_name: [16]int8{'q'},
		Ftag_type: int32(Q),
	},
	89: {
		Ftag_name: [16]int8{'r', 'b'},
		Ftag_type: int32(RB),
	},
	90: {
		Ftag_name: [16]int8{'r', 'p'},
		Ftag_type: int32(RP),
	},
	91: {
		Ftag_name: [16]int8{'r', 't'},
		Ftag_type: int32(RT),
	},
	92: {
		Ftag_name: [16]int8{'r', 't', 'c'},
		Ftag_type: int32(RTC),
	},
	93: {
		Ftag_name: [16]int8{'r', 'u', 'b', 'y'},
		Ftag_type: int32(RUBY),
	},
	94: {
		Ftag_name: [16]int8{'s'},
		Ftag_type: int32(S),
	},
	95: {
		Ftag_name: [16]int8{'s', 'a', 'm', 'p'},
		Ftag_type: int32(SAMP),
	},
	96: {
		Ftag_name: [16]int8{'s', 'c', 'r', 'i', 'p', 't'},
		Ftag_type: int32(SCRIPT),
	},
	97: {
		Ftag_name: [16]int8{'s', 'e', 'c', 't', 'i', 'o', 'n'},
		Ftag_type: int32(SECTION),
	},
	98: {
		Ftag_name: [16]int8{'s', 'e', 'l', 'e', 'c', 't'},
		Ftag_type: int32(SELECT),
	},
	99: {
		Ftag_name: [16]int8{'s', 'l', 'o', 't'},
		Ftag_type: int32(SLOT),
	},
	100: {
		Ftag_name: [16]int8{'s', 'm', 'a', 'l', 'l'},
		Ftag_type: int32(SMALL),
	},
	101: {
		Ftag_name: [16]int8{'s', 'p', 'a', 'n'},
		Ftag_type: int32(SPAN),
	},
	102: {
		Ftag_name: [16]int8{'s', 't', 'r', 'o', 'n', 'g'},
		Ftag_type: int32(STRONG),
	},
	103: {
		Ftag_name: [16]int8{'s', 't', 'y', 'l', 'e'},
		Ftag_type: int32(STYLE),
	},
	104: {
		Ftag_name: [16]int8{'s', 'u', 'b'},
		Ftag_type: int32(SUB),
	},
	105: {
		Ftag_name: [16]int8{'s', 'u', 'm', 'm', 'a', 'r', 'y'},
		Ftag_type: int32(SUMMARY),
	},
	106: {
		Ftag_name: [16]int8{'s', 'u', 'p'},
		Ftag_type: int32(SUP),
	},
	107: {
		Ftag_name: [16]int8{'s', 'v', 'g'},
		Ftag_type: int32(SVG),
	},
	108: {
		Ftag_name: [16]int8{'t', 'a', 'b', 'l', 'e'},
		Ftag_type: int32(TABLE),
	},
	109: {
		Ftag_name: [16]int8{'t', 'b', 'o', 'd', 'y'},
		Ftag_type: int32(TBODY),
	},
	110: {
		Ftag_name: [16]int8{'t', 'd'},
		Ftag_type: int32(TD),
	},
	111: {
		Ftag_name: [16]int8{'t', 'e', 'm', 'p', 'l', 'a', 't', 'e'},
		Ftag_type: int32(TEMPLATE),
	},
	112: {
		Ftag_name: [16]int8{'t', 'e', 'x', 't', 'a', 'r', 'e', 'a'},
		Ftag_type: int32(TEXTAREA),
	},
	113: {
		Ftag_name: [16]int8{'t', 'f', 'o', 'o', 't'},
		Ftag_type: int32(TFOOT),
	},
	114: {
		Ftag_name: [16]int8{'t', 'h'},
		Ftag_type: int32(TH),
	},
	115: {
		Ftag_name: [16]int8{'t', 'h', 'e', 'a', 'd'},
		Ftag_type: int32(THEAD),
	},
	116: {
		Ftag_name: [16]int8{'t', 'i', 'm', 'e'},
		Ftag_type: int32(TIME),
	},
	117: {
		Ftag_name: [16]int8{'t', 'i', 't', 'l', 'e'},
		Ftag_type: int32(TITLE),
	},
	118: {
		Ftag_name: [16]int8{'t', 'r'},
		Ftag_type: int32(TR),
	},
	119: {
		Ftag_name: [16]int8{'u'},
		Ftag_type: int32(U),
	},
	120: {
		Ftag_name: [16]int8{'u', 'l'},
		Ftag_type: int32(UL),
	},
	121: {
		Ftag_name: [16]int8{'v', 'a', 'r'},
		Ftag_type: int32(VAR),
	},
	122: {
		Ftag_name: [16]int8{'v', 'i', 'd', 'e', 'o'},
		Ftag_type: int32(VIDEO),
	},
	123: {
		Ftag_name: [16]int8{'c', 'u', 's', 't', 'o', 'm'},
		Ftag_type: int32(CUSTOM),
	},
}
var TAG_TYPES_NOT_ALLOWED_IN_PARAGRAPHS = [26]TagType{
	0:  int32(ADDRESS),
	1:  int32(ARTICLE),
	2:  int32(ASIDE),
	3:  int32(BLOCKQUOTE),
	4:  int32(DETAILS),
	5:  int32(DIV),
	6:  int32(DL),
	7:  int32(FIELDSET),
	8:  int32(FIGCAPTION),
	9:  int32(FIGURE),
	10: int32(FOOTER),
	11: int32(FORM),
	12: int32(H1),
	13: int32(H2),
	14: int32(H3),
	15: int32(H4),
	16: int32(H5),
	17: int32(H6),
	18: int32(HEADER),
	19: int32(HR),
	20: int32(MAIN),
	21: int32(NAV),
	22: int32(OL),
	23: int32(P),
	24: int32(PRE),
	25: int32(SECTION),
}

func tag_type_for_name(tls *libc.TLS, tag_name uintptr) (r TagType) {
	var entry uintptr
	var i int32
	_, _ = entry, i
	i = 0
	for {
		if !(i < int32(126)) {
			break
		}
		entry = uintptr(unsafe.Pointer(&TAG_TYPES_BY_TAG_NAME)) + uintptr(i)*20
		if libc.Xstrlen(tls, entry) == uint64((*String)(unsafe.Pointer(tag_name)).Fsize) && libc.Xmemcmp(tls, (*String)(unsafe.Pointer(tag_name)).Fcontents, entry, uint64((*String)(unsafe.Pointer(tag_name)).Fsize)) == 0 {
			return (*TagMapEntry)(unsafe.Pointer(entry)).Ftag_type
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	return int32(CUSTOM)
}

func tag_can_contain(tls *libc.TLS, self uintptr, other uintptr) (r uint8) {
	var child TagType
	var i int32
	_, _ = child, i
	child = (*Tag)(unsafe.Pointer(other)).Ftype_token
	if child == int32(INTERPOLATION) {
		return libc.BoolUint8(true1 != 0)
	}
	switch (*Tag)(unsafe.Pointer(self)).Ftype_token {
	case int32(LI):
		goto _1
	case int32(DD):
		goto _2
	case int32(DT):
		goto _3
	case int32(P):
		goto _4
	case int32(COLGROUP):
		goto _5
	case int32(RP):
		goto _6
	case int32(RT):
		goto _7
	case int32(RB):
		goto _8
	case int32(OPTGROUP):
		goto _9
	case int32(TR):
		goto _10
	case int32(TH):
		goto _11
	case int32(TD):
		goto _12
	default:
		goto _13
	}
	goto _14
_1:
	;
	return libc.BoolUint8(child != int32(LI))
_3:
	;
_2:
	;
	return libc.BoolUint8(child != int32(DT) && child != int32(DD))
_4:
	;
	i = 0
_17:
	;
	if !(uint64(i) < libc.Uint64FromInt64(104)/libc.Uint64FromInt64(4)) {
		goto _15
	}
	if child == TAG_TYPES_NOT_ALLOWED_IN_PARAGRAPHS[i] {
		return libc.BoolUint8(false1 != 0)
	}
	goto _16
_16:
	;
	i = i + 1
	goto _17
_15:
	;
	return libc.BoolUint8(true1 != 0)
_5:
	;
	return libc.BoolUint8(child == int32(COL))
_8:
	;
_7:
	;
_6:
	;
	return libc.BoolUint8(child != int32(RB) && child != int32(RT) && child != int32(RP))
_9:
	;
	return libc.BoolUint8(child != int32(OPTGROUP))
_10:
	;
	return libc.BoolUint8(child != int32(TR))
_12:
	;
_11:
	;
	return libc.BoolUint8(child != int32(TD) && child != int32(TH) && child != int32(TR))
_13:
	;
	return libc.BoolUint8(true1 != 0)
_14:
	;
	return r
}

type wctrans_t = uint16

type TokenType = int32

const START_TAG_NAME = 0
const SCRIPT_START_TAG_NAME = 1
const STYLE_START_TAG_NAME = 2
const END_TAG_NAME = 3
const ERRONEOUS_END_TAG_NAME = 4
const SELF_CLOSING_TAG_DELIMITER = 5
const IMPLICIT_END_TAG = 6
const RAW_TEXT = 7
const COMMENT = 8
const HTML_INTERPOLATION_START = 9
const HTML_INTERPOLATION_END = 10
const FRONTMATTER_JS_BLOCK = 11
const ATTRIBUTE_JS_EXPR = 12
const ATTRIBUTE_BACKTICK_STRING = 13
const PERMISSIBLE_TEXT = 14
const FRAGMENT_TAG_DELIM = 15

type Scanner = struct {
	Ftags struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	}
}

func advance(tls *libc.TLS, lexer uintptr) {
	(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(false1 != 0))
}

func skip(tls *libc.TLS, lexer uintptr) {
	(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(true1 != 0))
}

func serialize(tls *libc.TLS, scanner uintptr, buffer uintptr) (r uint32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var name_length, size, v1 uint32
	var tag Tag
	var _ /* serialized_tag_count at bp+2 */ uint16_t
	var _ /* tag_count at bp+0 */ uint16_t
	_, _, _, _ = name_length, size, tag, v1
	if (*Scanner)(unsafe.Pointer(scanner)).Ftags.Fsize > uint32(UINT16_MAX) {
		v1 = uint32(UINT16_MAX)
	} else {
		v1 = (*Scanner)(unsafe.Pointer(scanner)).Ftags.Fsize
	}
	*(*uint16_t)(unsafe.Pointer(bp)) = uint16(v1)
	*(*uint16_t)(unsafe.Pointer(bp + 2)) = uint16(0)
	size = uint32(2)
	libc.Xmemcpy(tls, buffer+uintptr(size), bp, uint64(2))
	size = uint32(uint64(size) + libc.Uint64FromInt64(2))
	for {
		if !(int32(*(*uint16_t)(unsafe.Pointer(bp + 2))) < int32(*(*uint16_t)(unsafe.Pointer(bp)))) {
			break
		}
		tag = *(*Tag)(unsafe.Pointer((*Scanner)(unsafe.Pointer(scanner)).Ftags.Fcontents + uintptr(*(*uint16_t)(unsafe.Pointer(bp + 2)))*24))
		if tag.Ftype_token == int32(CUSTOM) {
			name_length = tag.Fcustom_tag_name.Fsize
			if name_length > uint32(UINT8_MAX) {
				name_length = uint32(UINT8_MAX)
			}
			if size+uint32(2)+name_length >= uint32(TREE_SITTER_SERIALIZATION_BUFFER_SIZE) {
				break
			}
			v1 = size
			size = size + 1
			*(*int8)(unsafe.Pointer(buffer + uintptr(v1))) = int8(tag.Ftype_token)
			v1 = size
			size = size + 1
			*(*int8)(unsafe.Pointer(buffer + uintptr(v1))) = int8(name_length)
			libc.Xstrncpy(tls, buffer+uintptr(size), tag.Fcustom_tag_name.Fcontents, uint64(name_length))
			size = size + name_length
		} else {
			if size+uint32(1) >= uint32(TREE_SITTER_SERIALIZATION_BUFFER_SIZE) {
				break
			}
			v1 = size
			size = size + 1
			*(*int8)(unsafe.Pointer(buffer + uintptr(v1))) = int8(tag.Ftype_token)
		}
		goto _2
	_2:
		;
		*(*uint16_t)(unsafe.Pointer(bp + 2)) = *(*uint16_t)(unsafe.Pointer(bp + 2)) + 1
	}
	libc.Xmemcpy(tls, buffer, bp+2, uint64(2))
	return size
}

func deserialize(tls *libc.TLS, scanner uintptr, buffer uintptr, length uint32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var i, iter, size, v10 uint32
	var name_length uint16_t
	var new_capacity1, new_size, v14, v6 uint32_t
	var tag, v8 Tag
	var v2, v3, v4 uintptr
	var v5 size_t
	var _ /* serialized_tag_count at bp+2 */ uint16_t
	var _ /* tag at bp+8 */ Tag
	var _ /* tag_count at bp+0 */ uint16_t
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = i, iter, name_length, new_capacity1, new_size, size, tag, v10, v14, v2, v3, v4, v5, v6, v8
	i = uint32(0)
	for {
		if !(i < (*Scanner)(unsafe.Pointer(scanner)).Ftags.Fsize) {
			break
		}
		v2 = (*Scanner)(unsafe.Pointer(scanner)).Ftags.Fcontents + uintptr(i)*24
		if (*Tag)(unsafe.Pointer(v2)).Ftype_token == int32(CUSTOM) {
			v3 = v2 + 8
			if (*Array)(unsafe.Pointer(v3)).Fcontents != 0 {
				libc.Xfree(tls, (*Array)(unsafe.Pointer(v3)).Fcontents)
				(*Array)(unsafe.Pointer(v3)).Fcontents = libc.UintptrFromInt32(0)
				(*Array)(unsafe.Pointer(v3)).Fsize = uint32(0)
				(*Array)(unsafe.Pointer(v3)).Fcapacity = uint32(0)
			}
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(scanner)).Fsize = uint32(0)
	if length > uint32(0) {
		size = uint32(0)
		*(*uint16_t)(unsafe.Pointer(bp)) = uint16(0)
		*(*uint16_t)(unsafe.Pointer(bp + 2)) = uint16(0)
		libc.Xmemcpy(tls, bp+2, buffer+uintptr(size), uint64(2))
		size = uint32(uint64(size) + libc.Uint64FromInt64(2))
		libc.Xmemcpy(tls, bp, buffer+uintptr(size), uint64(2))
		size = uint32(uint64(size) + libc.Uint64FromInt64(2))
		v2 = scanner
		v5 = libc.Uint64FromInt64(24)
		v6 = uint32(*(*uint16_t)(unsafe.Pointer(bp)))
		if v6 > (*Array)(unsafe.Pointer(v2)).Fcapacity {
			if (*Array)(unsafe.Pointer(v2)).Fcontents != 0 {
				(*Array)(unsafe.Pointer(v2)).Fcontents = libc.Xrealloc(tls, (*Array)(unsafe.Pointer(v2)).Fcontents, uint64(v6)*v5)
			} else {
				(*Array)(unsafe.Pointer(v2)).Fcontents = libc.Xmalloc(tls, uint64(v6)*v5)
			}
			(*Array)(unsafe.Pointer(v2)).Fcapacity = v6
		}
		if int32(*(*uint16_t)(unsafe.Pointer(bp))) > 0 {
			iter = uint32(0)
			iter = uint32(0)
			for {
				if !(iter < uint32(*(*uint16_t)(unsafe.Pointer(bp + 2)))) {
					break
				}
				tag.Ftype_token = int32(END_)
				tag.Fcustom_tag_name = String{}
				v8 = tag
				goto _9
			_9:
				*(*Tag)(unsafe.Pointer(bp + 8)) = v8
				v10 = size
				size = size + 1
				(*(*Tag)(unsafe.Pointer(bp + 8))).Ftype_token = int32(*(*int8)(unsafe.Pointer(buffer + uintptr(v10))))
				if (*(*Tag)(unsafe.Pointer(bp + 8))).Ftype_token == int32(CUSTOM) {
					v10 = size
					size = size + 1
					name_length = uint16(uint8(*(*int8)(unsafe.Pointer(buffer + uintptr(v10)))))
					v2 = bp + 8 + 8
					v5 = libc.Uint64FromInt64(1)
					v6 = uint32(name_length)
					if v6 > (*Array)(unsafe.Pointer(v2)).Fcapacity {
						if (*Array)(unsafe.Pointer(v2)).Fcontents != 0 {
							(*Array)(unsafe.Pointer(v2)).Fcontents = libc.Xrealloc(tls, (*Array)(unsafe.Pointer(v2)).Fcontents, uint64(v6)*v5)
						} else {
							(*Array)(unsafe.Pointer(v2)).Fcontents = libc.Xmalloc(tls, uint64(v6)*v5)
						}
						(*Array)(unsafe.Pointer(v2)).Fcapacity = v6
					}
					(*(*Tag)(unsafe.Pointer(bp + 8))).Fcustom_tag_name.Fsize = uint32(name_length)
					libc.Xmemcpy(tls, (*(*Tag)(unsafe.Pointer(bp + 8))).Fcustom_tag_name.Fcontents, buffer+uintptr(size), uint64(name_length))
					size = size + uint32(name_length)
				}
				v2 = scanner
				new_size = (*Array)(unsafe.Pointer(v2)).Fsize + uint32(1)
				if new_size > (*Array)(unsafe.Pointer(v2)).Fcapacity {
					new_capacity1 = (*Array)(unsafe.Pointer(v2)).Fcapacity * uint32(2)
					if new_capacity1 < libc.Uint32FromInt32(8) {
						new_capacity1 = uint32(8)
					}
					if new_capacity1 < new_size {
						new_capacity1 = new_size
					}
					v3 = v2
					v5 = libc.Uint64FromInt64(24)
					v6 = new_capacity1
					if v6 > (*Array)(unsafe.Pointer(v3)).Fcapacity {
						if (*Array)(unsafe.Pointer(v3)).Fcontents != 0 {
							(*Array)(unsafe.Pointer(v3)).Fcontents = libc.Xrealloc(tls, (*Array)(unsafe.Pointer(v3)).Fcontents, uint64(v6)*v5)
						} else {
							(*Array)(unsafe.Pointer(v3)).Fcontents = libc.Xmalloc(tls, uint64(v6)*v5)
						}
						(*Array)(unsafe.Pointer(v3)).Fcapacity = v6
					}
				}
				v4 = scanner + 8
				v14 = *(*uint32_t)(unsafe.Pointer(v4))
				*(*uint32_t)(unsafe.Pointer(v4)) = *(*uint32_t)(unsafe.Pointer(v4)) + 1
				*(*Tag)(unsafe.Pointer((*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(scanner)).Fcontents + uintptr(v14)*24)) = *(*Tag)(unsafe.Pointer(bp + 8))
				goto _7
			_7:
				;
				iter = iter + 1
			}
			// add zero tags if we didn't read enough, this is because the
			// buffer had no more room but we held more tags.
			for {
				if !(iter < uint32(*(*uint16_t)(unsafe.Pointer(bp)))) {
					break
				}
				v2 = scanner
				new_size = (*Array)(unsafe.Pointer(v2)).Fsize + uint32(1)
				if new_size > (*Array)(unsafe.Pointer(v2)).Fcapacity {
					new_capacity1 = (*Array)(unsafe.Pointer(v2)).Fcapacity * uint32(2)
					if new_capacity1 < libc.Uint32FromInt32(8) {
						new_capacity1 = uint32(8)
					}
					if new_capacity1 < new_size {
						new_capacity1 = new_size
					}
					v3 = v2
					v5 = libc.Uint64FromInt64(24)
					v6 = new_capacity1
					if v6 > (*Array)(unsafe.Pointer(v3)).Fcapacity {
						if (*Array)(unsafe.Pointer(v3)).Fcontents != 0 {
							(*Array)(unsafe.Pointer(v3)).Fcontents = libc.Xrealloc(tls, (*Array)(unsafe.Pointer(v3)).Fcontents, uint64(v6)*v5)
						} else {
							(*Array)(unsafe.Pointer(v3)).Fcontents = libc.Xmalloc(tls, uint64(v6)*v5)
						}
						(*Array)(unsafe.Pointer(v3)).Fcapacity = v6
					}
				}
				v4 = scanner + 8
				v14 = *(*uint32_t)(unsafe.Pointer(v4))
				*(*uint32_t)(unsafe.Pointer(v4)) = *(*uint32_t)(unsafe.Pointer(v4)) + 1
				tag.Ftype_token = int32(END_)
				tag.Fcustom_tag_name = String{}
				v8 = tag
				goto _29
			_29:
				*(*Tag)(unsafe.Pointer((*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(scanner)).Fcontents + uintptr(v14)*24)) = v8
				goto _21
			_21:
				;
				iter = iter + 1
			}
		}
	}
}

func scan_tag_name(tls *libc.TLS, lexer uintptr) (r String) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var new_capacity1, new_size, v4, v5 uint32_t
	var v1, v2, v6 uintptr
	var v3 size_t
	var _ /* tag_name at bp+0 */ String
	_, _, _, _, _, _, _, _ = new_capacity1, new_size, v1, v2, v3, v4, v5, v6
	*(*String)(unsafe.Pointer(bp)) = String{}
	for libc.Xiswctype(tls, uint16((*TSLexer)(unsafe.Pointer(lexer)).Flookahead), uint16(libc.Int32FromInt32(0x0100)|libc.Int32FromInt32(_UPPER)|libc.Int32FromInt32(_LOWER)|libc.Int32FromInt32(_DIGIT))) != 0 || (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('-') || (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32(':') || (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('.') {
		v1 = bp
		new_size = (*Array)(unsafe.Pointer(v1)).Fsize + uint32(1)
		if new_size > (*Array)(unsafe.Pointer(v1)).Fcapacity {
			new_capacity1 = (*Array)(unsafe.Pointer(v1)).Fcapacity * uint32(2)
			if new_capacity1 < libc.Uint32FromInt32(8) {
				new_capacity1 = uint32(8)
			}
			if new_capacity1 < new_size {
				new_capacity1 = new_size
			}
			v2 = v1
			v3 = libc.Uint64FromInt64(1)
			v4 = new_capacity1
			if v4 > (*Array)(unsafe.Pointer(v2)).Fcapacity {
				if (*Array)(unsafe.Pointer(v2)).Fcontents != 0 {
					(*Array)(unsafe.Pointer(v2)).Fcontents = libc.Xrealloc(tls, (*Array)(unsafe.Pointer(v2)).Fcontents, uint64(v4)*v3)
				} else {
					(*Array)(unsafe.Pointer(v2)).Fcontents = libc.Xmalloc(tls, uint64(v4)*v3)
				}
				(*Array)(unsafe.Pointer(v2)).Fcapacity = v4
			}
		}
		v6 = bp + 8
		v5 = *(*uint32_t)(unsafe.Pointer(v6))
		*(*uint32_t)(unsafe.Pointer(v6)) = *(*uint32_t)(unsafe.Pointer(v6)) + 1
		*(*int8)(unsafe.Pointer((*String)(unsafe.Pointer(bp)).Fcontents + uintptr(v5))) = int8((*TSLexer)(unsafe.Pointer(lexer)).Flookahead)
		advance(tls, lexer)
	}
	return *(*String)(unsafe.Pointer(bp))
}

func scan_comment(tls *libc.TLS, lexer uintptr) (r uint8) {
	var dashes uint32
	_ = dashes
	if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32('-') {
		return libc.BoolUint8(false1 != 0)
	}
	advance(tls, lexer)
	if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32('-') {
		return libc.BoolUint8(false1 != 0)
	}
	advance(tls, lexer)
	dashes = uint32(0)
	for (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != 0 {
		switch (*TSLexer)(unsafe.Pointer(lexer)).Flookahead {
		case int32('-'):
			dashes = dashes + 1
		case int32('>'):
			if dashes >= uint32(2) {
				(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(COMMENT)
				advance(tls, lexer)
				(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
				return libc.BoolUint8(true1 != 0)
			}
			fallthrough
		default:
			dashes = uint32(0)
		}
		advance(tls, lexer)
	}
	return libc.BoolUint8(false1 != 0)
}

type RawTextEndType = int32

const
// corresponds to the ending delimiter "\n---"
EndFrontmatter = 0
const
// corresponds to the ending delimiter "}",
// used for JS backtick strings and attribute interpolations.
// we have to balance brackets with this one.
EndCurly = 1

func scan_js_backtick_string(tls *libc.TLS, lexer uintptr) {
	// Advance past backtick
	(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(false1 != 0))
	for (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32('\000') {
		if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('$') {
			(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(false1 != 0))
			if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('{') {
				// String interpolation
				(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(false1 != 0))
				scan_js_expr_with_delimiter(tls, lexer, int32(EndCurly))
				// Advance past the final curly
			} else {
				// Reprocess this character
				continue
			}
		} else {
			if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('`') {
				// End of string
				(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(false1 != 0))
				break
			}
		}
		(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(false1 != 0))
	}
}

func scan_js_string(tls *libc.TLS, lexer uintptr) {
	var str_end_char int8
	_ = str_end_char
	// Assumes the lookahead char is actually valid
	if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('`') {
		scan_js_backtick_string(tls, lexer)
	} else {
		// Start and end chars are the same
		str_end_char = int8((*TSLexer)(unsafe.Pointer(lexer)).Flookahead)
		(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(false1 != 0))
		for (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32('\000') {
			// Note that this doesn't take into account newlines in basic
			// strings, for the same reason why tree-sitter-javascript
			// doesn't.
			// https://github.com/tree-sitter/tree-sitter-javascript/blob/fdeb68ac8d2bd5a78b943528bb68ceda3aade2eb/grammar.js#L860
			if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('\\') {
				// Accept any next char
				(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(false1 != 0))
			} else {
				if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32(str_end_char) {
					// End of string
					(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(false1 != 0))
					return
				}
			}
			(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(false1 != 0))
		}
	}
}

func scan_js_expr_with_delimiter(tls *libc.TLS, lexer uintptr, end_type RawTextEndType) {
	var END uintptr
	var curly_count, delimiter_index uint32
	var in_comment, v1 int32
	_, _, _, _, _ = END, curly_count, delimiter_index, in_comment, v1
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
	// `delimiter_index` is only used when `end_type == EndFrontmatter`.
	// We assign "1" here because we only enter this function when "---\n" is parsed by tree-sitter,
	// but tree-sitter leaves out the extra newline when giving the lexer to us.
	// "1" reflects the true delimiter status.
	// (This ensures parsing empty delimiters correctly.)
	delimiter_index = uint32(1)
	curly_count = uint32(0)
	in_comment = 0
	for (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32('\000') {
		if in_comment == 0 {
			// delimiter check
			if end_type == int32(EndFrontmatter) {
				// We have to `mark_end` at index 0, always,
				// so just pre-emptively do it here.
				if delimiter_index == uint32(0) {
					(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
				}
				END = __ccgo_ts
				if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32(*(*int8)(unsafe.Pointer(END + uintptr(delimiter_index)))) {
					delimiter_index = delimiter_index + 1
					if uint64(delimiter_index) == libc.Xstrlen(tls, END) {
						break
					}
				} else {
					// In case we stumble into a newline. This could happen if e.g.
					// ---
					// -
					// ---
					// was the frontmatter.
					(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
					if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('\n') {
						v1 = int32(1)
					} else {
						v1 = 0
					}
					delimiter_index = uint32(v1)
				}
			} else {
				if end_type == int32(EndCurly) {
					(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
					// balance braces
					if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('{') {
						curly_count = curly_count + 1
					} else {
						if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('}') {
							if curly_count == uint32(0) {
								// delimiter, break
								(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
								break
							}
							curly_count = curly_count - 1
						}
					}
				}
			}
			if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('"') || (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('\'') || (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('`') {
				scan_js_string(tls, lexer)
				continue
			}
			if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('/') {
				// comment?
				(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(false1 != 0))
				if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('/') {
					in_comment = 1
				} else {
					if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('*') {
						in_comment = 2
					}
				}
				continue
			}
		} else {
			if in_comment == 1 {
				if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('\n') {
					// End of comment
					in_comment = 0
					// Frontmatter delimiters start with 1 newline.
					if end_type == int32(EndFrontmatter) {
						v1 = int32(1)
					} else {
						v1 = 0
					}
					delimiter_index = uint32(v1)
					// `mark_end` here ensures that we don't accidentally skip a frontmatter delimiter start point.
					// If this was left out, delimiters of the form
					// ---
					// // comment
					// ---
					// would never `mark_end` before the delimiter, causing the returned token to be truncated in size.
					(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
				}
			} else {
				if in_comment == 2 {
					if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('*') {
						(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(false1 != 0))
						if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('/') {
							// End of comment
							in_comment = 0
							delimiter_index = uint32(0)
						} else {
							continue
						}
					}
				}
			}
		}
		(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(false1 != 0))
	}
}

func scan_raw_text(tls *libc.TLS, scanner uintptr, lexer uintptr) (r uint8) {
	var delimiter_index uint32
	var end_delimiter, v1 uintptr
	var v2 bool
	_, _, _, _ = delimiter_index, end_delimiter, v1, v2
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
	if v2 = !!((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(scanner)).Fsize-libc.Uint32FromInt32(1) < (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(scanner)).Fsize); !v2 {
		libc.X_assert(tls, __ccgo_ts+5, __ccgo_ts+69, uint32(315))
	}
	_ = v2 || libc.Bool(libc.Int32FromInt32(0) != 0)
	if (*Tag)(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(scanner)).Fcontents+uintptr((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(scanner)).Fsize-uint32(1))*24)).Ftype_token == int32(SCRIPT) {
		v1 = __ccgo_ts + 80
	} else {
		v1 = __ccgo_ts + 89
	}
	end_delimiter = v1
	delimiter_index = uint32(0)
	for (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != 0 {
		// TODO add test for uppercase SCRIPT not conflicting with lowercase script
		if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32(*(*int8)(unsafe.Pointer(end_delimiter + uintptr(delimiter_index)))) {
			delimiter_index = delimiter_index + 1
			if uint64(delimiter_index) == libc.Xstrlen(tls, end_delimiter) {
				break
			}
			advance(tls, lexer)
		} else {
			delimiter_index = uint32(0)
			advance(tls, lexer)
			(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		}
	}
	(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(RAW_TEXT)
	return libc.BoolUint8(true1 != 0)
}

func pop_tag(tls *libc.TLS, scanner uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var v1 uint32_t
	var v2, v3, v4 uintptr
	var _ /* popped_tag at bp+0 */ Tag
	_, _, _, _ = v1, v2, v3, v4
	v2 = scanner + 8
	*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) - 1
	v1 = *(*uint32_t)(unsafe.Pointer(v2))
	*(*Tag)(unsafe.Pointer(bp)) = *(*Tag)(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(scanner)).Fcontents + uintptr(v1)*24))
	v3 = bp
	if (*Tag)(unsafe.Pointer(v3)).Ftype_token == int32(CUSTOM) {
		v4 = v3 + 8
		if (*Array)(unsafe.Pointer(v4)).Fcontents != 0 {
			libc.Xfree(tls, (*Array)(unsafe.Pointer(v4)).Fcontents)
			(*Array)(unsafe.Pointer(v4)).Fcontents = libc.UintptrFromInt32(0)
			(*Array)(unsafe.Pointer(v4)).Fsize = uint32(0)
			(*Array)(unsafe.Pointer(v4)).Fcapacity = uint32(0)
		}
	}
}

func scan_implicit_end_tag(tls *libc.TLS, scanner uintptr, lexer uintptr) (r uint8) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var i uint32
	var is_closing_tag, v3 uint8
	var parent, v1, v10, v14, v15, v6 uintptr
	var tag, tag1, v11, v8 Tag
	var v2, v5 bool
	var v7 String
	var _ /* next_tag at bp+32 */ Tag
	var _ /* tag_name at bp+16 */ String
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = i, is_closing_tag, parent, tag, tag1, v1, v10, v11, v14, v15, v2, v3, v5, v6, v7, v8
	if (*Scanner)(unsafe.Pointer(scanner)).Ftags.Fsize == uint32(0) {
		v1 = libc.UintptrFromInt32(0)
	} else {
		if v2 = !!((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(scanner)).Fsize-libc.Uint32FromInt32(1) < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(scanner)).Fsize); !v2 {
			libc.X_assert(tls, __ccgo_ts+5, __ccgo_ts+69, uint32(343))
		}
		_ = v2 || libc.Bool(libc.Int32FromInt32(0) != 0)
		v1 = (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(scanner)).Fcontents + uintptr((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(scanner)).Fsize-uint32(1))*24
	}
	parent = v1
	is_closing_tag = libc.BoolUint8(false1 != 0)
	if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('/') {
		is_closing_tag = libc.BoolUint8(true1 != 0)
		advance(tls, lexer)
	} else {
		if v2 = parent != 0; v2 {
			v3 = libc.BoolUint8((*Tag)(unsafe.Pointer(parent)).Ftype_token < int32(END_OF_VOID_TAGS))
			goto _4
		_4:
		}
		if v2 && v3 != 0 {
			pop_tag(tls, scanner)
			(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(IMPLICIT_END_TAG)
			return libc.BoolUint8(true1 != 0)
		}
	}
	*(*String)(unsafe.Pointer(bp + 16)) = scan_tag_name(tls, lexer)
	if (*(*String)(unsafe.Pointer(bp + 16))).Fsize == uint32(0) && !((*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Feof})))(tls, lexer) != 0) {
		v1 = bp + 16
		if (*Array)(unsafe.Pointer(v1)).Fcontents != 0 {
			libc.Xfree(tls, (*Array)(unsafe.Pointer(v1)).Fcontents)
			(*Array)(unsafe.Pointer(v1)).Fcontents = libc.UintptrFromInt32(0)
			(*Array)(unsafe.Pointer(v1)).Fsize = uint32(0)
			(*Array)(unsafe.Pointer(v1)).Fcapacity = uint32(0)
		}
		return libc.BoolUint8(false1 != 0)
	}
	v7 = *(*String)(unsafe.Pointer(bp + 16))
	*(*String)(unsafe.Pointer(bp)) = v7
	tag.Ftype_token = int32(END_)
	tag.Fcustom_tag_name = String{}
	v8 = tag
	goto _9
_9:
	tag1 = v8
	tag1.Ftype_token = tag_type_for_name(tls, bp)
	if tag1.Ftype_token == int32(CUSTOM) {
		tag1.Fcustom_tag_name = v7
	} else {
		v1 = bp
		if (*Array)(unsafe.Pointer(v1)).Fcontents != 0 {
			libc.Xfree(tls, (*Array)(unsafe.Pointer(v1)).Fcontents)
			(*Array)(unsafe.Pointer(v1)).Fcontents = libc.UintptrFromInt32(0)
			(*Array)(unsafe.Pointer(v1)).Fsize = uint32(0)
			(*Array)(unsafe.Pointer(v1)).Fcapacity = uint32(0)
		}
	}
	v11 = tag1
	goto _12
_12:
	*(*Tag)(unsafe.Pointer(bp + 32)) = v11
	if is_closing_tag != 0 {
		// The tag correctly closes the topmost element on the stack
		if v5 = (*Scanner)(unsafe.Pointer(scanner)).Ftags.Fsize > uint32(0); v5 {
			if v2 = !!((*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(scanner)).Fsize-libc.Uint32FromInt32(1) < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(scanner)).Fsize); !v2 {
				libc.X_assert(tls, __ccgo_ts+5, __ccgo_ts+69, uint32(367))
			}
			_ = v2 || libc.Bool(libc.Int32FromInt32(0) != 0)
			v6 = (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(scanner)).Fcontents + uintptr((*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(scanner)).Fsize-uint32(1))*24
			v10 = bp + 32
			if (*Tag)(unsafe.Pointer(v6)).Ftype_token != (*Tag)(unsafe.Pointer(v10)).Ftype_token {
				v3 = libc.BoolUint8(false1 != 0)
				goto _17
			}
			if (*Tag)(unsafe.Pointer(v6)).Ftype_token == int32(CUSTOM) {
				if (*Tag)(unsafe.Pointer(v6)).Fcustom_tag_name.Fsize != (*Tag)(unsafe.Pointer(v10)).Fcustom_tag_name.Fsize {
					v3 = libc.BoolUint8(false1 != 0)
					goto _17
				}
				if libc.Xmemcmp(tls, (*Tag)(unsafe.Pointer(v6)).Fcustom_tag_name.Fcontents, (*Tag)(unsafe.Pointer(v10)).Fcustom_tag_name.Fcontents, uint64((*Tag)(unsafe.Pointer(v6)).Fcustom_tag_name.Fsize)) != libc.Int32FromInt32(0) {
					v3 = libc.BoolUint8(false1 != 0)
					goto _17
				}
			}
			v3 = libc.BoolUint8(true1 != 0)
			goto _17
		_17:
		}
		if v5 && v3 != 0 {
			v14 = bp + 32
			if (*Tag)(unsafe.Pointer(v14)).Ftype_token == int32(CUSTOM) {
				v15 = v14 + 8
				if (*Array)(unsafe.Pointer(v15)).Fcontents != 0 {
					libc.Xfree(tls, (*Array)(unsafe.Pointer(v15)).Fcontents)
					(*Array)(unsafe.Pointer(v15)).Fcontents = libc.UintptrFromInt32(0)
					(*Array)(unsafe.Pointer(v15)).Fsize = uint32(0)
					(*Array)(unsafe.Pointer(v15)).Fcapacity = uint32(0)
				}
			}
			return libc.BoolUint8(false1 != 0)
		}
		// Otherwise, dig deeper and queue implicit end tags (to be nice in
		// the case of malformed Astro)
		i = (*Scanner)(unsafe.Pointer(scanner)).Ftags.Fsize
		for {
			if !(i > uint32(0)) {
				break
			}
			if (*(*Tag)(unsafe.Pointer((*Scanner)(unsafe.Pointer(scanner)).Ftags.Fcontents + uintptr(i-uint32(1))*24))).Ftype_token == (*(*Tag)(unsafe.Pointer(bp + 32))).Ftype_token {
				pop_tag(tls, scanner)
				(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(IMPLICIT_END_TAG)
				v1 = bp + 32
				if (*Tag)(unsafe.Pointer(v1)).Ftype_token == int32(CUSTOM) {
					v6 = v1 + 8
					if (*Array)(unsafe.Pointer(v6)).Fcontents != 0 {
						libc.Xfree(tls, (*Array)(unsafe.Pointer(v6)).Fcontents)
						(*Array)(unsafe.Pointer(v6)).Fcontents = libc.UintptrFromInt32(0)
						(*Array)(unsafe.Pointer(v6)).Fsize = uint32(0)
						(*Array)(unsafe.Pointer(v6)).Fcapacity = uint32(0)
					}
				}
				return libc.BoolUint8(true1 != 0)
			}
			goto _21
		_21:
			;
			i = i - 1
		}
	} else {
		if parent != 0 && (!(tag_can_contain(tls, parent, bp+32) != 0) || ((*Tag)(unsafe.Pointer(parent)).Ftype_token == int32(HTML) || (*Tag)(unsafe.Pointer(parent)).Ftype_token == int32(HEAD) || (*Tag)(unsafe.Pointer(parent)).Ftype_token == int32(BODY)) && (*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Feof})))(tls, lexer) != 0) {
			pop_tag(tls, scanner)
			(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(IMPLICIT_END_TAG)
			v1 = bp + 32
			if (*Tag)(unsafe.Pointer(v1)).Ftype_token == int32(CUSTOM) {
				v6 = v1 + 8
				if (*Array)(unsafe.Pointer(v6)).Fcontents != 0 {
					libc.Xfree(tls, (*Array)(unsafe.Pointer(v6)).Fcontents)
					(*Array)(unsafe.Pointer(v6)).Fcontents = libc.UintptrFromInt32(0)
					(*Array)(unsafe.Pointer(v6)).Fsize = uint32(0)
					(*Array)(unsafe.Pointer(v6)).Fcapacity = uint32(0)
				}
			}
			return libc.BoolUint8(true1 != 0)
		}
	}
	v1 = bp + 32
	if (*Tag)(unsafe.Pointer(v1)).Ftype_token == int32(CUSTOM) {
		v6 = v1 + 8
		if (*Array)(unsafe.Pointer(v6)).Fcontents != 0 {
			libc.Xfree(tls, (*Array)(unsafe.Pointer(v6)).Fcontents)
			(*Array)(unsafe.Pointer(v6)).Fcontents = libc.UintptrFromInt32(0)
			(*Array)(unsafe.Pointer(v6)).Fsize = uint32(0)
			(*Array)(unsafe.Pointer(v6)).Fcapacity = uint32(0)
		}
	}
	return libc.BoolUint8(false1 != 0)
}

func scan_start_tag_name(tls *libc.TLS, scanner uintptr, lexer uintptr) (r uint8) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var new_capacity1, new_size, v5, v6 uint32_t
	var tag, tag1, tag2, tag3, v12, v9 Tag
	var v1, v2, v3, v7 uintptr
	var v4 size_t
	var v8 String
	var _ /* tag_name at bp+16 */ String
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = new_capacity1, new_size, tag, tag1, tag2, tag3, v1, v12, v2, v3, v4, v5, v6, v7, v8, v9
	*(*String)(unsafe.Pointer(bp + 16)) = scan_tag_name(tls, lexer)
	if (*(*String)(unsafe.Pointer(bp + 16))).Fsize == uint32(0) {
		v1 = bp + 16
		if (*Array)(unsafe.Pointer(v1)).Fcontents != 0 {
			libc.Xfree(tls, (*Array)(unsafe.Pointer(v1)).Fcontents)
			(*Array)(unsafe.Pointer(v1)).Fcontents = libc.UintptrFromInt32(0)
			(*Array)(unsafe.Pointer(v1)).Fsize = uint32(0)
			(*Array)(unsafe.Pointer(v1)).Fcapacity = uint32(0)
		}
		// Fragment tags don't contain spaces.
		if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('>') {
			advance(tls, lexer)
			tag2 = Tag{
				Ftype_token: int32(FRAGMENT),
			}
			v1 = scanner
			new_size = (*Array)(unsafe.Pointer(v1)).Fsize + uint32(1)
			if new_size > (*Array)(unsafe.Pointer(v1)).Fcapacity {
				new_capacity1 = (*Array)(unsafe.Pointer(v1)).Fcapacity * uint32(2)
				if new_capacity1 < libc.Uint32FromInt32(8) {
					new_capacity1 = uint32(8)
				}
				if new_capacity1 < new_size {
					new_capacity1 = new_size
				}
				v2 = v1
				v4 = libc.Uint64FromInt64(24)
				v5 = new_capacity1
				if v5 > (*Array)(unsafe.Pointer(v2)).Fcapacity {
					if (*Array)(unsafe.Pointer(v2)).Fcontents != 0 {
						(*Array)(unsafe.Pointer(v2)).Fcontents = libc.Xrealloc(tls, (*Array)(unsafe.Pointer(v2)).Fcontents, uint64(v5)*v4)
					} else {
						(*Array)(unsafe.Pointer(v2)).Fcontents = libc.Xmalloc(tls, uint64(v5)*v4)
					}
					(*Array)(unsafe.Pointer(v2)).Fcapacity = v5
				}
			}
			v3 = scanner + 8
			v6 = *(*uint32_t)(unsafe.Pointer(v3))
			*(*uint32_t)(unsafe.Pointer(v3)) = *(*uint32_t)(unsafe.Pointer(v3)) + 1
			*(*Tag)(unsafe.Pointer((*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(scanner)).Fcontents + uintptr(v6)*24)) = tag2
			(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(FRAGMENT_TAG_DELIM)
			return libc.BoolUint8(true1 != 0)
		} else {
			return libc.BoolUint8(false1 != 0)
		}
	}
	v8 = *(*String)(unsafe.Pointer(bp + 16))
	*(*String)(unsafe.Pointer(bp)) = v8
	tag.Ftype_token = int32(END_)
	tag.Fcustom_tag_name = String{}
	v9 = tag
	goto _10
_10:
	tag1 = v9
	tag1.Ftype_token = tag_type_for_name(tls, bp)
	if tag1.Ftype_token == int32(CUSTOM) {
		tag1.Fcustom_tag_name = v8
	} else {
		v1 = bp
		if (*Array)(unsafe.Pointer(v1)).Fcontents != 0 {
			libc.Xfree(tls, (*Array)(unsafe.Pointer(v1)).Fcontents)
			(*Array)(unsafe.Pointer(v1)).Fcontents = libc.UintptrFromInt32(0)
			(*Array)(unsafe.Pointer(v1)).Fsize = uint32(0)
			(*Array)(unsafe.Pointer(v1)).Fcapacity = uint32(0)
		}
	}
	v12 = tag1
	goto _13
_13:
	tag3 = v12
	v2 = scanner
	new_size = (*Array)(unsafe.Pointer(v2)).Fsize + uint32(1)
	if new_size > (*Array)(unsafe.Pointer(v2)).Fcapacity {
		new_capacity1 = (*Array)(unsafe.Pointer(v2)).Fcapacity * uint32(2)
		if new_capacity1 < libc.Uint32FromInt32(8) {
			new_capacity1 = uint32(8)
		}
		if new_capacity1 < new_size {
			new_capacity1 = new_size
		}
		v3 = v2
		v4 = libc.Uint64FromInt64(24)
		v5 = new_capacity1
		if v5 > (*Array)(unsafe.Pointer(v3)).Fcapacity {
			if (*Array)(unsafe.Pointer(v3)).Fcontents != 0 {
				(*Array)(unsafe.Pointer(v3)).Fcontents = libc.Xrealloc(tls, (*Array)(unsafe.Pointer(v3)).Fcontents, uint64(v5)*v4)
			} else {
				(*Array)(unsafe.Pointer(v3)).Fcontents = libc.Xmalloc(tls, uint64(v5)*v4)
			}
			(*Array)(unsafe.Pointer(v3)).Fcapacity = v5
		}
	}
	v7 = scanner + 8
	v6 = *(*uint32_t)(unsafe.Pointer(v7))
	*(*uint32_t)(unsafe.Pointer(v7)) = *(*uint32_t)(unsafe.Pointer(v7)) + 1
	*(*Tag)(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(scanner)).Fcontents + uintptr(v6)*24)) = tag3
	switch tag3.Ftype_token {
	case int32(SCRIPT):
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(SCRIPT_START_TAG_NAME)
	case int32(STYLE):
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(STYLE_START_TAG_NAME)
	default:
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(START_TAG_NAME)
		break
	}
	return libc.BoolUint8(true1 != 0)
}

func scan_end_tag_name(tls *libc.TLS, scanner uintptr, lexer uintptr) (r uint8) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var tag, tag1, v5, v8 Tag
	var v1, v11, v7 uintptr
	var v2, v3 bool
	var v13 uint8
	var v4 String
	var _ /* tag at bp+32 */ Tag
	var _ /* tag_name at bp+16 */ String
	_, _, _, _, _, _, _, _, _, _, _ = tag, tag1, v1, v11, v13, v2, v3, v4, v5, v7, v8
	*(*String)(unsafe.Pointer(bp + 16)) = scan_tag_name(tls, lexer)
	if (*(*String)(unsafe.Pointer(bp + 16))).Fsize == uint32(0) {
		v1 = bp + 16
		if (*Array)(unsafe.Pointer(v1)).Fcontents != 0 {
			libc.Xfree(tls, (*Array)(unsafe.Pointer(v1)).Fcontents)
			(*Array)(unsafe.Pointer(v1)).Fcontents = libc.UintptrFromInt32(0)
			(*Array)(unsafe.Pointer(v1)).Fsize = uint32(0)
			(*Array)(unsafe.Pointer(v1)).Fcapacity = uint32(0)
		}
		if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('>') {
			advance(tls, lexer)
			if v3 = (*Scanner)(unsafe.Pointer(scanner)).Ftags.Fsize > uint32(0); v3 {
				if v2 = !!((*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(scanner)).Fsize-libc.Uint32FromInt32(1) < (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(scanner)).Fsize); !v2 {
					libc.X_assert(tls, __ccgo_ts+5, __ccgo_ts+69, uint32(437))
				}
				_ = v2 || libc.Bool(libc.Int32FromInt32(0) != 0)
			}
			if v3 && (*Tag)(unsafe.Pointer((*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(scanner)).Fcontents+uintptr((*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(scanner)).Fsize-uint32(1))*24)).Ftype_token == int32(FRAGMENT) {
				pop_tag(tls, scanner)
				(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(FRAGMENT_TAG_DELIM)
				return libc.BoolUint8(true1 != 0)
			} else {
				(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(ERRONEOUS_END_TAG_NAME)
				return libc.BoolUint8(true1 != 0)
			}
		} else {
			// Closing fragment tags don't have spaces.
			return libc.BoolUint8(false1 != 0)
		}
	}
	v4 = *(*String)(unsafe.Pointer(bp + 16))
	*(*String)(unsafe.Pointer(bp)) = v4
	tag.Ftype_token = int32(END_)
	tag.Fcustom_tag_name = String{}
	v5 = tag
	goto _6
_6:
	tag1 = v5
	tag1.Ftype_token = tag_type_for_name(tls, bp)
	if tag1.Ftype_token == int32(CUSTOM) {
		tag1.Fcustom_tag_name = v4
	} else {
		v1 = bp
		if (*Array)(unsafe.Pointer(v1)).Fcontents != 0 {
			libc.Xfree(tls, (*Array)(unsafe.Pointer(v1)).Fcontents)
			(*Array)(unsafe.Pointer(v1)).Fcontents = libc.UintptrFromInt32(0)
			(*Array)(unsafe.Pointer(v1)).Fsize = uint32(0)
			(*Array)(unsafe.Pointer(v1)).Fcapacity = uint32(0)
		}
	}
	v8 = tag1
	goto _9
_9:
	*(*Tag)(unsafe.Pointer(bp + 32)) = v8
	if v3 = (*Scanner)(unsafe.Pointer(scanner)).Ftags.Fsize > uint32(0); v3 {
		if v2 = !!((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(scanner)).Fsize-libc.Uint32FromInt32(1) < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(scanner)).Fsize); !v2 {
			libc.X_assert(tls, __ccgo_ts+5, __ccgo_ts+69, uint32(452))
		}
		_ = v2 || libc.Bool(libc.Int32FromInt32(0) != 0)
		v7 = (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(scanner)).Fcontents + uintptr((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(scanner)).Fsize-uint32(1))*24
		v11 = bp + 32
		if (*Tag)(unsafe.Pointer(v7)).Ftype_token != (*Tag)(unsafe.Pointer(v11)).Ftype_token {
			v13 = libc.BoolUint8(false1 != 0)
			goto _14
		}
		if (*Tag)(unsafe.Pointer(v7)).Ftype_token == int32(CUSTOM) {
			if (*Tag)(unsafe.Pointer(v7)).Fcustom_tag_name.Fsize != (*Tag)(unsafe.Pointer(v11)).Fcustom_tag_name.Fsize {
				v13 = libc.BoolUint8(false1 != 0)
				goto _14
			}
			if libc.Xmemcmp(tls, (*Tag)(unsafe.Pointer(v7)).Fcustom_tag_name.Fcontents, (*Tag)(unsafe.Pointer(v11)).Fcustom_tag_name.Fcontents, uint64((*Tag)(unsafe.Pointer(v7)).Fcustom_tag_name.Fsize)) != libc.Int32FromInt32(0) {
				v13 = libc.BoolUint8(false1 != 0)
				goto _14
			}
		}
		v13 = libc.BoolUint8(true1 != 0)
		goto _14
	_14:
	}
	if v3 && v13 != 0 {
		pop_tag(tls, scanner)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(END_TAG_NAME)
	} else {
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(ERRONEOUS_END_TAG_NAME)
	}
	v1 = bp + 32
	if (*Tag)(unsafe.Pointer(v1)).Ftype_token == int32(CUSTOM) {
		v7 = v1 + 8
		if (*Array)(unsafe.Pointer(v7)).Fcontents != 0 {
			libc.Xfree(tls, (*Array)(unsafe.Pointer(v7)).Fcontents)
			(*Array)(unsafe.Pointer(v7)).Fcontents = libc.UintptrFromInt32(0)
			(*Array)(unsafe.Pointer(v7)).Fsize = uint32(0)
			(*Array)(unsafe.Pointer(v7)).Fcapacity = uint32(0)
		}
	}
	return libc.BoolUint8(true1 != 0)
}

func scan_self_closing_tag_delimiter(tls *libc.TLS, scanner uintptr, lexer uintptr) (r uint8) {
	advance(tls, lexer)
	if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('>') {
		advance(tls, lexer)
		if (*Scanner)(unsafe.Pointer(scanner)).Ftags.Fsize > uint32(0) {
			pop_tag(tls, scanner)
			(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(SELF_CLOSING_TAG_DELIMITER)
		}
		return libc.BoolUint8(true1 != 0)
	}
	return libc.BoolUint8(false1 != 0)
}

func scan_permissible_text(tls *libc.TLS, lexer uintptr) (r uint8) {
	var there_is_text uint8
	_ = there_is_text
	there_is_text = libc.BoolUint8(false1 != 0)
	for (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32('\000') {
		if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('{') || (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('}') {
			// Start of interpolation / end of interpolation, break
			break
		}
		if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('\'') || (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('"') || (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('`') {
			// skip string
			scan_js_string(tls, lexer)
			there_is_text = libc.BoolUint8(true1 != 0)
			goto text_found
		}
		if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('/') {
			// check for a comment
			// (c.f. https://github.com/withastro/compiler/blob/e8b6cdfc89f940a411304787632efd8140535feb/internal/token.go#L1017)
			advance(tls, lexer)
			there_is_text = libc.BoolUint8(true1 != 0)
			if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('/') {
				// single-line comment
				for (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32('\r') && (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32('\n') && (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32('\000') {
					advance(tls, lexer)
				}
			}
			if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('*') {
				// multiline comment
				for (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32('\000') {
					advance(tls, lexer)
					if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('*') {
						advance(tls, lexer)
						if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('/') {
							// end of multi-line comment
							advance(tls, lexer)
							break
						}
					}
				}
			}
			goto text_found
		}
		if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('<') {
			advance(tls, lexer)
			// This is the same logic the Astro compiler uses
			// to find when to terminate a text node.
			// https://github.com/withastro/compiler/blob/e8b6cdfc89f940a411304787632efd8140535feb/internal/token.go#L1737
			if int32('a') <= (*TSLexer)(unsafe.Pointer(lexer)).Flookahead && (*TSLexer)(unsafe.Pointer(lexer)).Flookahead <= int32('z') || int32('A') <= (*TSLexer)(unsafe.Pointer(lexer)).Flookahead && (*TSLexer)(unsafe.Pointer(lexer)).Flookahead <= int32('Z') {
				// Potential <start> tag
				break
			}
			if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('/') {
				// Potential </end> tag
				break
			}
			if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('/') || (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('?') {
				// Potential <!-- comment --> tag
				// or <!DOCTYPE ...> tag
				// or <?xml processing instructions?> tag
				break
			}
			if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('>') {
				// Fragment tag
				// (e.g. `<> <p> hi </p> </>`)
				break
			}
			// If none of the above conditions pass,
			// there's definitely text here.
			goto text_found
		}
		advance(tls, lexer)
		goto text_found
	text_found:
		;
		there_is_text = libc.BoolUint8(true1 != 0)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
	}
	// If there isn't any text,
	// then this can't be permissible text.
	if there_is_text != 0 {
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(PERMISSIBLE_TEXT)
		return libc.BoolUint8(true1 != 0)
	} else {
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

func scan(tls *libc.TLS, scanner uintptr, lexer uintptr, valid_symbols uintptr) (r uint8) {
	var definitely_not_permissible_text, invalid uint8
	var new_capacity1, new_size, v4, v5 uint32_t
	var tag Tag
	var v1, v2, v6 uintptr
	var v11 int32
	var v3 size_t
	var v7, v8 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _ = definitely_not_permissible_text, invalid, new_capacity1, new_size, tag, v1, v11, v2, v3, v4, v5, v6, v7, v8
	if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(FRONTMATTER_JS_BLOCK))) != 0 && (*Scanner)(unsafe.Pointer(scanner)).Ftags.Fsize == uint32(0) {
		scan_js_expr_with_delimiter(tls, lexer, int32(EndFrontmatter))
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(FRONTMATTER_JS_BLOCK)
		return libc.BoolUint8(true1 != 0)
	}
	if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(RAW_TEXT))) != 0 && !(*(*uint8)(unsafe.Pointer(valid_symbols + uintptr(START_TAG_NAME))) != 0) && !(*(*uint8)(unsafe.Pointer(valid_symbols + uintptr(END_TAG_NAME))) != 0) {
		return scan_raw_text(tls, scanner, lexer)
	}
	if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(ATTRIBUTE_JS_EXPR))) != 0 {
		scan_js_expr_with_delimiter(tls, lexer, int32(EndCurly))
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(ATTRIBUTE_JS_EXPR)
		return libc.BoolUint8(true1 != 0)
	}
	if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(PERMISSIBLE_TEXT))) != 0 {
		if libc.Xiswctype(tls, uint16((*TSLexer)(unsafe.Pointer(lexer)).Flookahead), uint16(_SPACE)) != 0 {
			// Can't be anything else.
			return scan_permissible_text(tls, lexer)
		}
	} else {
		for libc.Xiswctype(tls, uint16((*TSLexer)(unsafe.Pointer(lexer)).Flookahead), uint16(_SPACE)) != 0 {
			skip(tls, lexer)
		}
	}
	definitely_not_permissible_text = libc.BoolUint8(false1 != 0)
	switch (*TSLexer)(unsafe.Pointer(lexer)).Flookahead {
	case int32('<'):
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		advance(tls, lexer)
		if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('!') {
			advance(tls, lexer)
			return scan_comment(tls, lexer)
		}
		if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(IMPLICIT_END_TAG))) != 0 {
			return scan_implicit_end_tag(tls, scanner, lexer)
		}
		if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(PERMISSIBLE_TEXT))) != 0 {
			invalid = libc.BoolUint8(int32('a') <= (*TSLexer)(unsafe.Pointer(lexer)).Flookahead && (*TSLexer)(unsafe.Pointer(lexer)).Flookahead <= int32('z') || int32('A') <= (*TSLexer)(unsafe.Pointer(lexer)).Flookahead && (*TSLexer)(unsafe.Pointer(lexer)).Flookahead <= int32('Z') || (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('/') || (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('?') || (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('>'))
			if invalid != 0 {
				// This looks like an element,
				// so it can't be permissible text.
				definitely_not_permissible_text = libc.BoolUint8(true1 != 0)
			}
		}
	case int32('\000'):
		definitely_not_permissible_text = libc.BoolUint8(true1 != 0)
		if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(IMPLICIT_END_TAG))) != 0 {
			return scan_implicit_end_tag(tls, scanner, lexer)
		}
	case int32('/'):
		if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(SELF_CLOSING_TAG_DELIMITER))) != 0 {
			return scan_self_closing_tag_delimiter(tls, scanner, lexer)
		}
	case int32('{'):
		if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(HTML_INTERPOLATION_START))) != 0 {
			(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(false1 != 0))
			tag = Tag{
				Ftype_token: int32(INTERPOLATION),
			}
			v1 = scanner
			new_size = (*Array)(unsafe.Pointer(v1)).Fsize + uint32(1)
			if new_size > (*Array)(unsafe.Pointer(v1)).Fcapacity {
				new_capacity1 = (*Array)(unsafe.Pointer(v1)).Fcapacity * uint32(2)
				if new_capacity1 < libc.Uint32FromInt32(8) {
					new_capacity1 = uint32(8)
				}
				if new_capacity1 < new_size {
					new_capacity1 = new_size
				}
				v2 = v1
				v3 = libc.Uint64FromInt64(24)
				v4 = new_capacity1
				if v4 > (*Array)(unsafe.Pointer(v2)).Fcapacity {
					if (*Array)(unsafe.Pointer(v2)).Fcontents != 0 {
						(*Array)(unsafe.Pointer(v2)).Fcontents = libc.Xrealloc(tls, (*Array)(unsafe.Pointer(v2)).Fcontents, uint64(v4)*v3)
					} else {
						(*Array)(unsafe.Pointer(v2)).Fcontents = libc.Xmalloc(tls, uint64(v4)*v3)
					}
					(*Array)(unsafe.Pointer(v2)).Fcapacity = v4
				}
			}
			v6 = scanner + 8
			v5 = *(*uint32_t)(unsafe.Pointer(v6))
			*(*uint32_t)(unsafe.Pointer(v6)) = *(*uint32_t)(unsafe.Pointer(v6)) + 1
			*(*Tag)(unsafe.Pointer((*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(scanner)).Fcontents + uintptr(v5)*24)) = tag
			(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(HTML_INTERPOLATION_START)
			return libc.BoolUint8(true1 != 0)
		}
	case int32('}'):
		// Close any void tags before exiting the interpolation node.
		if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(IMPLICIT_END_TAG))) != 0 {
			return scan_implicit_end_tag(tls, scanner, lexer)
		}
		if v8 = *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(HTML_INTERPOLATION_END))) != 0 && (*Scanner)(unsafe.Pointer(scanner)).Ftags.Fsize > uint32(0); v8 {
			if v7 = !!((*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(scanner)).Fsize-libc.Uint32FromInt32(1) < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(scanner)).Fsize); !v7 {
				libc.X_assert(tls, __ccgo_ts+5, __ccgo_ts+69, uint32(650))
			}
			_ = v7 || libc.Bool(libc.Int32FromInt32(0) != 0)
		}
		if v8 && (*Tag)(unsafe.Pointer((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(scanner)).Fcontents+uintptr((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(scanner)).Fsize-uint32(1))*24)).Ftype_token == int32(INTERPOLATION) {
			(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(false1 != 0))
			v1 = scanner + 8
			*(*uint32_t)(unsafe.Pointer(v1)) = *(*uint32_t)(unsafe.Pointer(v1)) - 1
			v4 = *(*uint32_t)(unsafe.Pointer(v1))
			_ = *(*Tag)(unsafe.Pointer((*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(scanner)).Fcontents + uintptr(v4)*24))
			(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(HTML_INTERPOLATION_END)
			return libc.BoolUint8(true1 != 0)
		}
	case int32('`'):
		if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(ATTRIBUTE_BACKTICK_STRING))) != 0 {
			scan_js_backtick_string(tls, lexer)
			(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
			(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(ATTRIBUTE_BACKTICK_STRING)
			return libc.BoolUint8(true1 != 0)
		}
	default:
		if (*(*uint8)(unsafe.Pointer(valid_symbols + uintptr(START_TAG_NAME))) != 0 || *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(END_TAG_NAME))) != 0) && !(*(*uint8)(unsafe.Pointer(valid_symbols + uintptr(RAW_TEXT))) != 0) {
			if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(START_TAG_NAME))) != 0 {
				v11 = int32(scan_start_tag_name(tls, scanner, lexer))
			} else {
				v11 = int32(scan_end_tag_name(tls, scanner, lexer))
			}
			return uint8(libc.BoolInt32(v11 != 0))
		}
	}
	if !(definitely_not_permissible_text != 0) && *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(PERMISSIBLE_TEXT))) != 0 {
		// There are no other choices, it's this or nothing.
		return scan_permissible_text(tls, lexer)
	}
	return libc.BoolUint8(false1 != 0)
}

func tree_sitter_astro_external_scanner_create(tls *libc.TLS) (r uintptr) {
	var scanner uintptr
	_ = scanner
	scanner = libc.Xcalloc(tls, uint64(1), uint64(16))
	return scanner
}

func tree_sitter_astro_external_scanner_scan(tls *libc.TLS, payload uintptr, lexer uintptr, valid_symbols uintptr) (r uint8) {
	var scanner uintptr
	_ = scanner
	scanner = payload
	return scan(tls, scanner, lexer, valid_symbols)
}

func tree_sitter_astro_external_scanner_serialize(tls *libc.TLS, payload uintptr, buffer uintptr) (r uint32) {
	var scanner uintptr
	_ = scanner
	scanner = payload
	return serialize(tls, scanner, buffer)
}

func tree_sitter_astro_external_scanner_deserialize(tls *libc.TLS, payload uintptr, buffer uintptr, length uint32) {
	var scanner uintptr
	_ = scanner
	scanner = payload
	deserialize(tls, scanner, buffer, length)
}

func tree_sitter_astro_external_scanner_destroy(tls *libc.TLS, payload uintptr) {
	var i uint32
	var scanner, v2, v3 uintptr
	_, _, _, _ = i, scanner, v2, v3
	scanner = payload
	i = uint32(0)
	for {
		if !(i < (*Scanner)(unsafe.Pointer(scanner)).Ftags.Fsize) {
			break
		}
		v2 = (*Scanner)(unsafe.Pointer(scanner)).Ftags.Fcontents + uintptr(i)*24
		if (*Tag)(unsafe.Pointer(v2)).Ftype_token == int32(CUSTOM) {
			v3 = v2 + 8
			if (*Array)(unsafe.Pointer(v3)).Fcontents != 0 {
				libc.Xfree(tls, (*Array)(unsafe.Pointer(v3)).Fcontents)
				(*Array)(unsafe.Pointer(v3)).Fcontents = libc.UintptrFromInt32(0)
				(*Array)(unsafe.Pointer(v3)).Fsize = uint32(0)
				(*Array)(unsafe.Pointer(v3)).Fcapacity = uint32(0)
			}
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	v2 = scanner
	if (*Array)(unsafe.Pointer(v2)).Fcontents != 0 {
		libc.Xfree(tls, (*Array)(unsafe.Pointer(v2)).Fcontents)
		(*Array)(unsafe.Pointer(v2)).Fcontents = libc.UintptrFromInt32(0)
		(*Array)(unsafe.Pointer(v2)).Fsize = uint32(0)
		(*Array)(unsafe.Pointer(v2)).Fcapacity = uint32(0)
	}
	libc.Xfree(tls, scanner)
}

type ts_symbol_identifiers = int32

const anon_sym_LT_BANG = 1
const aux_sym_doctype_token1 = 2
const anon_sym_GT = 3
const sym__doctype = 4
const anon_sym_LT = 5
const anon_sym_SLASH_GT = 6
const anon_sym_LT_SLASH = 7
const anon_sym_EQ = 8
const sym_attribute_name = 9
const sym_attribute_value = 10
const sym_entity = 11
const anon_sym_SQUOTE = 12
const aux_sym_quoted_attribute_value_token1 = 13
const anon_sym_DQUOTE = 14
const aux_sym_quoted_attribute_value_token2 = 15
const sym_text = 16
const anon_sym_DASH_DASH_DASH = 17
const anon_sym_DASH_DASH_DASH2 = 18
const anon_sym_LBRACE = 19
const anon_sym_RBRACE = 20
const sym__start_tag_name = 21
const sym__script_start_tag_name = 22
const sym__style_start_tag_name = 23
const sym__end_tag_name = 24
const sym_erroneous_end_tag_name = 25
const sym__implicit_end_tag = 26
const sym_raw_text = 27
const sym_comment = 28
const sym__html_interpolation_start = 29
const sym__html_interpolation_end = 30
const sym_frontmatter_js_block = 31
const sym_attribute_js_expr = 32
const sym_attribute_backtick_string = 33
const sym_permissible_text = 34
const sym__fragment_tag_delim = 35
const sym_document = 36
const sym_doctype = 37
const sym__node = 38
const sym_element = 39
const sym_script_element = 40
const sym_style_element = 41
const sym_start_tag = 42
const sym_script_start_tag = 43
const sym_style_start_tag = 44
const sym_self_closing_tag = 45
const sym_end_tag = 46
const sym_erroneous_end_tag = 47
const sym_attribute = 48
const sym_quoted_attribute_value = 49
const sym__node_with_permissible_text = 50
const sym_frontmatter = 51
const sym_attribute_interpolation = 52
const sym_html_interpolation = 53
const sym_self_closing_script_tag = 54
const sym_self_closing_style_tag = 55
const aux_sym_document_repeat1 = 56
const aux_sym_start_tag_repeat1 = 57
const aux_sym_html_interpolation_repeat1 = 58

var ts_symbol_names = [59]uintptr{
	0:  __ccgo_ts + 97,
	1:  __ccgo_ts + 101,
	2:  __ccgo_ts + 104,
	3:  __ccgo_ts + 119,
	4:  __ccgo_ts + 121,
	5:  __ccgo_ts + 129,
	6:  __ccgo_ts + 131,
	7:  __ccgo_ts + 134,
	8:  __ccgo_ts + 137,
	9:  __ccgo_ts + 139,
	10: __ccgo_ts + 154,
	11: __ccgo_ts + 170,
	12: __ccgo_ts + 177,
	13: __ccgo_ts + 154,
	14: __ccgo_ts + 179,
	15: __ccgo_ts + 154,
	16: __ccgo_ts + 181,
	17: __ccgo_ts + 186,
	18: __ccgo_ts + 186,
	19: __ccgo_ts + 190,
	20: __ccgo_ts + 192,
	21: __ccgo_ts + 194,
	22: __ccgo_ts + 194,
	23: __ccgo_ts + 194,
	24: __ccgo_ts + 194,
	25: __ccgo_ts + 203,
	26: __ccgo_ts + 226,
	27: __ccgo_ts + 244,
	28: __ccgo_ts + 253,
	29: __ccgo_ts + 190,
	30: __ccgo_ts + 192,
	31: __ccgo_ts + 261,
	32: __ccgo_ts + 282,
	33: __ccgo_ts + 300,
	34: __ccgo_ts + 326,
	35: __ccgo_ts + 119,
	36: __ccgo_ts + 343,
	37: __ccgo_ts + 121,
	38: __ccgo_ts + 352,
	39: __ccgo_ts + 358,
	40: __ccgo_ts + 366,
	41: __ccgo_ts + 381,
	42: __ccgo_ts + 395,
	43: __ccgo_ts + 395,
	44: __ccgo_ts + 395,
	45: __ccgo_ts + 405,
	46: __ccgo_ts + 422,
	47: __ccgo_ts + 430,
	48: __ccgo_ts + 448,
	49: __ccgo_ts + 458,
	50: __ccgo_ts + 481,
	51: __ccgo_ts + 509,
	52: __ccgo_ts + 521,
	53: __ccgo_ts + 545,
	54: __ccgo_ts + 405,
	55: __ccgo_ts + 405,
	56: __ccgo_ts + 564,
	57: __ccgo_ts + 581,
	58: __ccgo_ts + 599,
}

var ts_symbol_map = [59]TSSymbol{
	1:  uint16(anon_sym_LT_BANG),
	2:  uint16(aux_sym_doctype_token1),
	3:  uint16(anon_sym_GT),
	4:  uint16(sym__doctype),
	5:  uint16(anon_sym_LT),
	6:  uint16(anon_sym_SLASH_GT),
	7:  uint16(anon_sym_LT_SLASH),
	8:  uint16(anon_sym_EQ),
	9:  uint16(sym_attribute_name),
	10: uint16(sym_attribute_value),
	11: uint16(sym_entity),
	12: uint16(anon_sym_SQUOTE),
	13: uint16(sym_attribute_value),
	14: uint16(anon_sym_DQUOTE),
	15: uint16(sym_attribute_value),
	16: uint16(sym_text),
	17: uint16(anon_sym_DASH_DASH_DASH),
	18: uint16(anon_sym_DASH_DASH_DASH),
	19: uint16(anon_sym_LBRACE),
	20: uint16(anon_sym_RBRACE),
	21: uint16(sym__start_tag_name),
	22: uint16(sym__start_tag_name),
	23: uint16(sym__start_tag_name),
	24: uint16(sym__start_tag_name),
	25: uint16(sym_erroneous_end_tag_name),
	26: uint16(sym__implicit_end_tag),
	27: uint16(sym_raw_text),
	28: uint16(sym_comment),
	29: uint16(anon_sym_LBRACE),
	30: uint16(anon_sym_RBRACE),
	31: uint16(sym_frontmatter_js_block),
	32: uint16(sym_attribute_js_expr),
	33: uint16(sym_attribute_backtick_string),
	34: uint16(sym_permissible_text),
	35: uint16(anon_sym_GT),
	36: uint16(sym_document),
	37: uint16(sym_doctype),
	38: uint16(sym__node),
	39: uint16(sym_element),
	40: uint16(sym_script_element),
	41: uint16(sym_style_element),
	42: uint16(sym_start_tag),
	43: uint16(sym_start_tag),
	44: uint16(sym_start_tag),
	45: uint16(sym_self_closing_tag),
	46: uint16(sym_end_tag),
	47: uint16(sym_erroneous_end_tag),
	48: uint16(sym_attribute),
	49: uint16(sym_quoted_attribute_value),
	50: uint16(sym__node_with_permissible_text),
	51: uint16(sym_frontmatter),
	52: uint16(sym_attribute_interpolation),
	53: uint16(sym_html_interpolation),
	54: uint16(sym_self_closing_tag),
	55: uint16(sym_self_closing_tag),
	56: uint16(aux_sym_document_repeat1),
	57: uint16(aux_sym_start_tag_repeat1),
	58: uint16(aux_sym_html_interpolation_repeat1),
}

var ts_symbol_metadata = [59]TSSymbolMetadata{
	0: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	1: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	2: {},
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
	13: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	14: {
		Fvisible: libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
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
	},
	30: {
		Fvisible: libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
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
	56: {},
	57: {},
	58: {},
}

var ts_alias_sequences = [1][4]TSSymbol{}

var ts_non_terminal_alias_map = [1]uint16_t{}

var ts_primary_state_ids = [163]TSStateId{
	1:   uint16(1),
	2:   uint16(2),
	3:   uint16(2),
	4:   uint16(4),
	5:   uint16(4),
	6:   uint16(4),
	7:   uint16(2),
	8:   uint16(8),
	9:   uint16(9),
	10:  uint16(8),
	11:  uint16(11),
	12:  uint16(12),
	13:  uint16(13),
	14:  uint16(14),
	15:  uint16(13),
	16:  uint16(16),
	17:  uint16(13),
	18:  uint16(16),
	19:  uint16(16),
	20:  uint16(20),
	21:  uint16(21),
	22:  uint16(22),
	23:  uint16(23),
	24:  uint16(21),
	25:  uint16(23),
	26:  uint16(20),
	27:  uint16(27),
	28:  uint16(20),
	29:  uint16(29),
	30:  uint16(27),
	31:  uint16(23),
	32:  uint16(27),
	33:  uint16(29),
	34:  uint16(21),
	35:  uint16(22),
	36:  uint16(36),
	37:  uint16(29),
	38:  uint16(22),
	39:  uint16(39),
	40:  uint16(40),
	41:  uint16(41),
	42:  uint16(42),
	43:  uint16(43),
	44:  uint16(44),
	45:  uint16(45),
	46:  uint16(46),
	47:  uint16(41),
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
	61:  uint16(60),
	62:  uint16(62),
	63:  uint16(55),
	64:  uint16(64),
	65:  uint16(51),
	66:  uint16(66),
	67:  uint16(49),
	68:  uint16(48),
	69:  uint16(59),
	70:  uint16(58),
	71:  uint16(57),
	72:  uint16(56),
	73:  uint16(52),
	74:  uint16(74),
	75:  uint16(75),
	76:  uint16(50),
	77:  uint16(40),
	78:  uint16(43),
	79:  uint16(42),
	80:  uint16(46),
	81:  uint16(75),
	82:  uint16(74),
	83:  uint16(66),
	84:  uint16(84),
	85:  uint16(64),
	86:  uint16(53),
	87:  uint16(74),
	88:  uint16(66),
	89:  uint16(40),
	90:  uint16(56),
	91:  uint16(57),
	92:  uint16(58),
	93:  uint16(59),
	94:  uint16(60),
	95:  uint16(55),
	96:  uint16(64),
	97:  uint16(49),
	98:  uint16(48),
	99:  uint16(43),
	100: uint16(42),
	101: uint16(41),
	102: uint16(46),
	103: uint16(75),
	104: uint16(104),
	105: uint16(52),
	106: uint16(50),
	107: uint16(51),
	108: uint16(108),
	109: uint16(109),
	110: uint16(110),
	111: uint16(111),
	112: uint16(112),
	113: uint16(113),
	114: uint16(112),
	115: uint16(112),
	116: uint16(116),
	117: uint16(117),
	118: uint16(117),
	119: uint16(116),
	120: uint16(120),
	121: uint16(116),
	122: uint16(117),
	123: uint16(120),
	124: uint16(120),
	125: uint16(125),
	126: uint16(126),
	127: uint16(127),
	128: uint16(128),
	129: uint16(129),
	130: uint16(128),
	131: uint16(131),
	132: uint16(132),
	133: uint16(133),
	134: uint16(132),
	135: uint16(131),
	136: uint16(131),
	137: uint16(128),
	138: uint16(138),
	139: uint16(139),
	140: uint16(132),
	141: uint16(141),
	142: uint16(142),
	143: uint16(143),
	144: uint16(144),
	145: uint16(145),
	146: uint16(146),
	147: uint16(147),
	148: uint16(148),
	149: uint16(146),
	150: uint16(150),
	151: uint16(141),
	152: uint16(142),
	153: uint16(153),
	154: uint16(141),
	155: uint16(146),
	156: uint16(156),
	157: uint16(153),
	158: uint16(150),
	159: uint16(159),
	160: uint16(150),
	161: uint16(148),
	162: uint16(148),
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
			state = uint16(24)
			goto next_state
		}
		if lookahead == int32('"') {
			state = uint16(79)
			goto next_state
		}
		if lookahead == int32('&') {
			state = uint16(3)
			goto next_state
		}
		if lookahead == int32('\'') {
			state = uint16(76)
			goto next_state
		}
		if lookahead == int32('-') {
			state = uint16(7)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(11)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(30)
			goto next_state
		}
		if lookahead == int32('=') {
			state = uint16(33)
			goto next_state
		}
		if lookahead == int32('>') {
			state = uint16(28)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(87)
			goto next_state
		}
		if lookahead == int32('}') {
			state = uint16(88)
			goto next_state
		}
		if lookahead == int32('D') || lookahead == int32('d') {
			state = uint16(14)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(0)
			goto next_state
		}
		return result
	case int32(1):
		if lookahead == int32('"') {
			state = uint16(79)
			goto next_state
		}
		if lookahead == int32('\'') {
			state = uint16(76)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(87)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(1)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('<') || int32('>') < lookahead) {
			state = uint16(35)
			goto next_state
		}
		return result
	case int32(2):
		if lookahead == int32('"') {
			state = uint16(79)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(80)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(81)
			goto next_state
		}
		return result
	case int32(3):
		if lookahead == int32('#') {
			state = uint16(17)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(75)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead == int32('\'') {
			state = uint16(76)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(77)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(78)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead == int32('-') {
			state = uint16(85)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead == int32('-') {
			state = uint16(86)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead == int32('-') {
			state = uint16(5)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead == int32('-') {
			state = uint16(6)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead == int32('-') {
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
		if lookahead == int32('/') {
			state = uint16(11)
			goto next_state
		}
		if lookahead == int32('=') {
			state = uint16(33)
			goto next_state
		}
		if lookahead == int32('>') {
			state = uint16(28)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(87)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(10)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('"') && lookahead != int32('\'') && lookahead != int32('<') {
			state = uint16(34)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead == int32('>') {
			state = uint16(31)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead == int32('C') || lookahead == int32('c') {
			state = uint16(16)
			goto next_state
		}
		return result
	case int32(13):
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(29)
			goto next_state
		}
		return result
	case int32(14):
		if lookahead == int32('O') || lookahead == int32('o') {
			state = uint16(12)
			goto next_state
		}
		return result
	case int32(15):
		if lookahead == int32('P') || lookahead == int32('p') {
			state = uint16(13)
			goto next_state
		}
		return result
	case int32(16):
		if lookahead == int32('T') || lookahead == int32('t') {
			state = uint16(18)
			goto next_state
		}
		return result
	case int32(17):
		if lookahead == int32('X') || lookahead == int32('x') {
			state = uint16(21)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(41)
			goto next_state
		}
		return result
	case int32(18):
		if lookahead == int32('Y') || lookahead == int32('y') {
			state = uint16(15)
			goto next_state
		}
		return result
	case int32(19):
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(19)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('<') && lookahead != int32('>') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(84)
			goto next_state
		}
		return result
	case int32(20):
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(26)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('>') {
			state = uint16(27)
			goto next_state
		}
		return result
	case int32(21):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(46)
			goto next_state
		}
		return result
	case int32(22):
		if eof != 0 {
			state = uint16(24)
			goto next_state
		}
		if lookahead == int32('-') {
			state = uint16(83)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(30)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(22)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('>') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(84)
			goto next_state
		}
		return result
	case int32(23):
		if eof != 0 {
			state = uint16(24)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(30)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(23)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('>') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(84)
			goto next_state
		}
		return result
	case int32(24):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(ts_builtin_sym_end)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(25):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT_BANG)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(26):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_doctype_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(26)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('>') {
			state = uint16(27)
			goto next_state
		}
		return result
	case int32(27):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_doctype_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('>') {
			state = uint16(27)
			goto next_state
		}
		return result
	case int32(28):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_GT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(29):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__doctype)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(30):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('!') {
			state = uint16(25)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(32)
			goto next_state
		}
		return result
	case int32(31):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SLASH_GT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(32):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT_SLASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(33):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(34):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_attribute_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32('"') && lookahead != int32('\'') && lookahead != int32('/') && (lookahead < int32('<') || int32('>') < lookahead) {
			state = uint16(34)
			goto next_state
		}
		return result
	case int32(35):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_attribute_value)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32('"') && lookahead != int32('\'') && (lookahead < int32('<') || int32('>') < lookahead) {
			state = uint16(35)
			goto next_state
		}
		return result
	case int32(36):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(37):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(36)
			goto next_state
		}
		return result
	case int32(38):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(36)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(37)
			goto next_state
		}
		return result
	case int32(39):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(36)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(38)
			goto next_state
		}
		return result
	case int32(40):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(36)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(39)
			goto next_state
		}
		return result
	case int32(41):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(36)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(40)
			goto next_state
		}
		return result
	case int32(42):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(36)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(37)
			goto next_state
		}
		return result
	case int32(43):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(36)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(42)
			goto next_state
		}
		return result
	case int32(44):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(36)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(43)
			goto next_state
		}
		return result
	case int32(45):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(36)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(44)
			goto next_state
		}
		return result
	case int32(46):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(36)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(45)
			goto next_state
		}
		return result
	case int32(47):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(36)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(37)
			goto next_state
		}
		return result
	case int32(48):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(36)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(47)
			goto next_state
		}
		return result
	case int32(49):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(36)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(48)
			goto next_state
		}
		return result
	case int32(50):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(36)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(49)
			goto next_state
		}
		return result
	case int32(51):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(36)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(50)
			goto next_state
		}
		return result
	case int32(52):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(36)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(51)
			goto next_state
		}
		return result
	case int32(53):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(36)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(54):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(36)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(53)
			goto next_state
		}
		return result
	case int32(55):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(36)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(54)
			goto next_state
		}
		return result
	case int32(56):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(36)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(55)
			goto next_state
		}
		return result
	case int32(57):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(36)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(56)
			goto next_state
		}
		return result
	case int32(58):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(36)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(57)
			goto next_state
		}
		return result
	case int32(59):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(36)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(58)
			goto next_state
		}
		return result
	case int32(60):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(36)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(59)
			goto next_state
		}
		return result
	case int32(61):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(36)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(60)
			goto next_state
		}
		return result
	case int32(62):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(36)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(61)
			goto next_state
		}
		return result
	case int32(63):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(36)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(62)
			goto next_state
		}
		return result
	case int32(64):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(36)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(63)
			goto next_state
		}
		return result
	case int32(65):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(36)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(64)
			goto next_state
		}
		return result
	case int32(66):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(36)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(65)
			goto next_state
		}
		return result
	case int32(67):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(36)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(66)
			goto next_state
		}
		return result
	case int32(68):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(36)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(67)
			goto next_state
		}
		return result
	case int32(69):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(36)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(68)
			goto next_state
		}
		return result
	case int32(70):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(36)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(71):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(36)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(70)
			goto next_state
		}
		return result
	case int32(72):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(36)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(71)
			goto next_state
		}
		return result
	case int32(73):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(36)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(72)
			goto next_state
		}
		return result
	case int32(74):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(36)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(73)
			goto next_state
		}
		return result
	case int32(75):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(36)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(74)
			goto next_state
		}
		return result
	case int32(76):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(77):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_quoted_attribute_value_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(77)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\'') {
			state = uint16(78)
			goto next_state
		}
		return result
	case int32(78):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_quoted_attribute_value_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\'') {
			state = uint16(78)
			goto next_state
		}
		return result
	case int32(79):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(80):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_quoted_attribute_value_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(80)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('"') {
			state = uint16(81)
			goto next_state
		}
		return result
	case int32(81):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_quoted_attribute_value_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('"') {
			state = uint16(81)
			goto next_state
		}
		return result
	case int32(82):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') {
			state = uint16(85)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(19)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('<') && lookahead != int32('>') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(84)
			goto next_state
		}
		return result
	case int32(83):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') {
			state = uint16(82)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(19)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('<') && lookahead != int32('>') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(84)
			goto next_state
		}
		return result
	case int32(84):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(19)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('<') && lookahead != int32('>') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(84)
			goto next_state
		}
		return result
	case int32(85):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DASH_DASH_DASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(86):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DASH_DASH_DASH2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(87):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(88):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var ts_lex_modes = [163]TSLexMode{
	0: {
		Fexternal_lex_state: uint16(1),
	},
	1: {
		Flex_state:          uint16(22),
		Fexternal_lex_state: uint16(2),
	},
	2: {
		Flex_state:          uint16(23),
		Fexternal_lex_state: uint16(3),
	},
	3: {
		Flex_state:          uint16(23),
		Fexternal_lex_state: uint16(3),
	},
	4: {
		Flex_state:          uint16(23),
		Fexternal_lex_state: uint16(3),
	},
	5: {
		Flex_state:          uint16(23),
		Fexternal_lex_state: uint16(3),
	},
	6: {
		Flex_state:          uint16(23),
		Fexternal_lex_state: uint16(3),
	},
	7: {
		Flex_state:          uint16(23),
		Fexternal_lex_state: uint16(3),
	},
	8: {
		Flex_state:          uint16(23),
		Fexternal_lex_state: uint16(2),
	},
	9: {
		Flex_state:          uint16(23),
		Fexternal_lex_state: uint16(2),
	},
	10: {
		Flex_state:          uint16(23),
		Fexternal_lex_state: uint16(3),
	},
	11: {
		Flex_state:          uint16(23),
		Fexternal_lex_state: uint16(2),
	},
	12: {
		Flex_state:          uint16(23),
		Fexternal_lex_state: uint16(2),
	},
	13: {
		Fexternal_lex_state: uint16(4),
	},
	14: {
		Fexternal_lex_state: uint16(4),
	},
	15: {
		Fexternal_lex_state: uint16(4),
	},
	16: {
		Fexternal_lex_state: uint16(4),
	},
	17: {
		Fexternal_lex_state: uint16(4),
	},
	18: {
		Fexternal_lex_state: uint16(4),
	},
	19: {
		Fexternal_lex_state: uint16(4),
	},
	20: {
		Flex_state:          uint16(10),
		Fexternal_lex_state: uint16(5),
	},
	21: {
		Flex_state:          uint16(10),
		Fexternal_lex_state: uint16(5),
	},
	22: {
		Flex_state:          uint16(10),
		Fexternal_lex_state: uint16(5),
	},
	23: {
		Flex_state:          uint16(10),
		Fexternal_lex_state: uint16(5),
	},
	24: {
		Flex_state:          uint16(10),
		Fexternal_lex_state: uint16(5),
	},
	25: {
		Flex_state:          uint16(10),
		Fexternal_lex_state: uint16(5),
	},
	26: {
		Flex_state:          uint16(10),
		Fexternal_lex_state: uint16(5),
	},
	27: {
		Flex_state:          uint16(10),
		Fexternal_lex_state: uint16(5),
	},
	28: {
		Flex_state:          uint16(10),
		Fexternal_lex_state: uint16(5),
	},
	29: {
		Flex_state:          uint16(10),
		Fexternal_lex_state: uint16(5),
	},
	30: {
		Flex_state:          uint16(10),
		Fexternal_lex_state: uint16(5),
	},
	31: {
		Flex_state:          uint16(10),
		Fexternal_lex_state: uint16(5),
	},
	32: {
		Flex_state:          uint16(10),
		Fexternal_lex_state: uint16(5),
	},
	33: {
		Flex_state:          uint16(10),
		Fexternal_lex_state: uint16(5),
	},
	34: {
		Flex_state:          uint16(10),
		Fexternal_lex_state: uint16(5),
	},
	35: {
		Flex_state:          uint16(10),
		Fexternal_lex_state: uint16(5),
	},
	36: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(6),
	},
	37: {
		Flex_state:          uint16(10),
		Fexternal_lex_state: uint16(5),
	},
	38: {
		Flex_state:          uint16(10),
		Fexternal_lex_state: uint16(5),
	},
	39: {
		Flex_state:          uint16(10),
		Fexternal_lex_state: uint16(5),
	},
	40: {
		Flex_state:          uint16(23),
		Fexternal_lex_state: uint16(2),
	},
	41: {
		Flex_state:          uint16(23),
		Fexternal_lex_state: uint16(3),
	},
	42: {
		Flex_state:          uint16(23),
		Fexternal_lex_state: uint16(2),
	},
	43: {
		Flex_state:          uint16(23),
		Fexternal_lex_state: uint16(2),
	},
	44: {
		Flex_state:          uint16(23),
		Fexternal_lex_state: uint16(2),
	},
	45: {
		Flex_state:          uint16(23),
		Fexternal_lex_state: uint16(3),
	},
	46: {
		Flex_state:          uint16(23),
		Fexternal_lex_state: uint16(2),
	},
	47: {
		Flex_state:          uint16(23),
		Fexternal_lex_state: uint16(2),
	},
	48: {
		Flex_state:          uint16(23),
		Fexternal_lex_state: uint16(3),
	},
	49: {
		Flex_state:          uint16(23),
		Fexternal_lex_state: uint16(3),
	},
	50: {
		Flex_state:          uint16(23),
		Fexternal_lex_state: uint16(2),
	},
	51: {
		Flex_state:          uint16(23),
		Fexternal_lex_state: uint16(3),
	},
	52: {
		Flex_state:          uint16(23),
		Fexternal_lex_state: uint16(2),
	},
	53: {
		Flex_state:          uint16(23),
		Fexternal_lex_state: uint16(2),
	},
	54: {
		Flex_state:          uint16(23),
		Fexternal_lex_state: uint16(2),
	},
	55: {
		Flex_state:          uint16(23),
		Fexternal_lex_state: uint16(3),
	},
	56: {
		Flex_state:          uint16(23),
		Fexternal_lex_state: uint16(2),
	},
	57: {
		Flex_state:          uint16(23),
		Fexternal_lex_state: uint16(2),
	},
	58: {
		Flex_state:          uint16(23),
		Fexternal_lex_state: uint16(2),
	},
	59: {
		Flex_state:          uint16(23),
		Fexternal_lex_state: uint16(2),
	},
	60: {
		Flex_state:          uint16(23),
		Fexternal_lex_state: uint16(2),
	},
	61: {
		Flex_state:          uint16(23),
		Fexternal_lex_state: uint16(3),
	},
	62: {
		Flex_state:          uint16(23),
		Fexternal_lex_state: uint16(3),
	},
	63: {
		Flex_state:          uint16(23),
		Fexternal_lex_state: uint16(2),
	},
	64: {
		Flex_state:          uint16(23),
		Fexternal_lex_state: uint16(2),
	},
	65: {
		Flex_state:          uint16(23),
		Fexternal_lex_state: uint16(2),
	},
	66: {
		Flex_state:          uint16(23),
		Fexternal_lex_state: uint16(2),
	},
	67: {
		Flex_state:          uint16(23),
		Fexternal_lex_state: uint16(2),
	},
	68: {
		Flex_state:          uint16(23),
		Fexternal_lex_state: uint16(2),
	},
	69: {
		Flex_state:          uint16(23),
		Fexternal_lex_state: uint16(3),
	},
	70: {
		Flex_state:          uint16(23),
		Fexternal_lex_state: uint16(3),
	},
	71: {
		Flex_state:          uint16(23),
		Fexternal_lex_state: uint16(3),
	},
	72: {
		Flex_state:          uint16(23),
		Fexternal_lex_state: uint16(3),
	},
	73: {
		Flex_state:          uint16(23),
		Fexternal_lex_state: uint16(3),
	},
	74: {
		Flex_state:          uint16(23),
		Fexternal_lex_state: uint16(2),
	},
	75: {
		Flex_state:          uint16(23),
		Fexternal_lex_state: uint16(2),
	},
	76: {
		Flex_state:          uint16(23),
		Fexternal_lex_state: uint16(3),
	},
	77: {
		Flex_state:          uint16(23),
		Fexternal_lex_state: uint16(3),
	},
	78: {
		Flex_state:          uint16(23),
		Fexternal_lex_state: uint16(3),
	},
	79: {
		Flex_state:          uint16(23),
		Fexternal_lex_state: uint16(3),
	},
	80: {
		Flex_state:          uint16(23),
		Fexternal_lex_state: uint16(3),
	},
	81: {
		Flex_state:          uint16(23),
		Fexternal_lex_state: uint16(3),
	},
	82: {
		Flex_state:          uint16(23),
		Fexternal_lex_state: uint16(3),
	},
	83: {
		Flex_state:          uint16(23),
		Fexternal_lex_state: uint16(3),
	},
	84: {
		Flex_state:          uint16(23),
		Fexternal_lex_state: uint16(3),
	},
	85: {
		Flex_state:          uint16(23),
		Fexternal_lex_state: uint16(3),
	},
	86: {
		Flex_state:          uint16(23),
		Fexternal_lex_state: uint16(3),
	},
	87: {
		Fexternal_lex_state: uint16(4),
	},
	88: {
		Fexternal_lex_state: uint16(4),
	},
	89: {
		Fexternal_lex_state: uint16(4),
	},
	90: {
		Fexternal_lex_state: uint16(4),
	},
	91: {
		Fexternal_lex_state: uint16(4),
	},
	92: {
		Fexternal_lex_state: uint16(4),
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
		Fexternal_lex_state: uint16(4),
	},
	97: {
		Fexternal_lex_state: uint16(4),
	},
	98: {
		Fexternal_lex_state: uint16(4),
	},
	99: {
		Fexternal_lex_state: uint16(4),
	},
	100: {
		Fexternal_lex_state: uint16(4),
	},
	101: {
		Fexternal_lex_state: uint16(4),
	},
	102: {
		Fexternal_lex_state: uint16(4),
	},
	103: {
		Fexternal_lex_state: uint16(4),
	},
	104: {
		Flex_state:          uint16(10),
		Fexternal_lex_state: uint16(5),
	},
	105: {
		Fexternal_lex_state: uint16(4),
	},
	106: {
		Fexternal_lex_state: uint16(4),
	},
	107: {
		Fexternal_lex_state: uint16(4),
	},
	108: {
		Flex_state:          uint16(10),
		Fexternal_lex_state: uint16(5),
	},
	109: {
		Flex_state:          uint16(10),
		Fexternal_lex_state: uint16(5),
	},
	110: {
		Flex_state:          uint16(10),
		Fexternal_lex_state: uint16(5),
	},
	111: {
		Flex_state:          uint16(10),
		Fexternal_lex_state: uint16(5),
	},
	112: {
		Fexternal_lex_state: uint16(7),
	},
	113: {
		Flex_state:          uint16(10),
		Fexternal_lex_state: uint16(5),
	},
	114: {
		Fexternal_lex_state: uint16(7),
	},
	115: {
		Fexternal_lex_state: uint16(7),
	},
	116: {
		Fexternal_lex_state: uint16(8),
	},
	117: {
		Fexternal_lex_state: uint16(8),
	},
	118: {
		Fexternal_lex_state: uint16(8),
	},
	119: {
		Fexternal_lex_state: uint16(8),
	},
	120: {
		Fexternal_lex_state: uint16(9),
	},
	121: {
		Fexternal_lex_state: uint16(8),
	},
	122: {
		Fexternal_lex_state: uint16(8),
	},
	123: {
		Fexternal_lex_state: uint16(9),
	},
	124: {
		Fexternal_lex_state: uint16(9),
	},
	125: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(10),
	},
	126: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(10),
	},
	127: {
		Fexternal_lex_state: uint16(8),
	},
	128: {
		Fexternal_lex_state: uint16(10),
	},
	129: {
		Fexternal_lex_state: uint16(8),
	},
	130: {
		Fexternal_lex_state: uint16(10),
	},
	131: {
		Fexternal_lex_state: uint16(11),
	},
	132: {
		Fexternal_lex_state: uint16(10),
	},
	133: {
		Flex_state:          uint16(9),
		Fexternal_lex_state: uint16(12),
	},
	134: {
		Fexternal_lex_state: uint16(10),
	},
	135: {
		Fexternal_lex_state: uint16(11),
	},
	136: {
		Fexternal_lex_state: uint16(11),
	},
	137: {
		Fexternal_lex_state: uint16(10),
	},
	138: {
		Fexternal_lex_state: uint16(8),
	},
	139: {
		Fexternal_lex_state: uint16(8),
	},
	140: {
		Fexternal_lex_state: uint16(10),
	},
	141: {
		Fexternal_lex_state: uint16(10),
	},
	142: {
		Fexternal_lex_state: uint16(10),
	},
	143: {
		Fexternal_lex_state: uint16(10),
	},
	144: {
		Fexternal_lex_state: uint16(10),
	},
	145: {
		Fexternal_lex_state: uint16(10),
	},
	146: {
		Fexternal_lex_state: uint16(10),
	},
	147: {
		Fexternal_lex_state: uint16(10),
	},
	148: {
		Fexternal_lex_state: uint16(10),
	},
	149: {
		Fexternal_lex_state: uint16(10),
	},
	150: {
		Flex_state:          uint16(20),
		Fexternal_lex_state: uint16(10),
	},
	151: {
		Fexternal_lex_state: uint16(10),
	},
	152: {
		Fexternal_lex_state: uint16(10),
	},
	153: {
		Fexternal_lex_state: uint16(13),
	},
	154: {
		Fexternal_lex_state: uint16(10),
	},
	155: {
		Fexternal_lex_state: uint16(10),
	},
	156: {
		Flex_state:          uint16(9),
		Fexternal_lex_state: uint16(10),
	},
	157: {
		Fexternal_lex_state: uint16(13),
	},
	158: {
		Flex_state:          uint16(20),
		Fexternal_lex_state: uint16(10),
	},
	159: {
		Fexternal_lex_state: uint16(14),
	},
	160: {
		Flex_state:          uint16(20),
		Fexternal_lex_state: uint16(10),
	},
	161: {
		Fexternal_lex_state: uint16(10),
	},
	162: {
		Fexternal_lex_state: uint16(10),
	},
}

var ts_parse_table = [2][59]uint16_t{
	0: {
		0:  uint16(1),
		1:  uint16(1),
		3:  uint16(1),
		4:  uint16(1),
		5:  uint16(1),
		6:  uint16(1),
		7:  uint16(1),
		8:  uint16(1),
		11: uint16(1),
		12: uint16(1),
		14: uint16(1),
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
		28: uint16(3),
		29: uint16(1),
		30: uint16(1),
		31: uint16(1),
		32: uint16(1),
		33: uint16(1),
		34: uint16(1),
		35: uint16(1),
	},
	1: {
		0:  uint16(5),
		1:  uint16(7),
		5:  uint16(9),
		7:  uint16(11),
		16: uint16(13),
		17: uint16(15),
		28: uint16(3),
		29: uint16(17),
		36: uint16(147),
		37: uint16(12),
		38: uint16(12),
		39: uint16(12),
		40: uint16(12),
		41: uint16(12),
		42: uint16(4),
		43: uint16(119),
		44: uint16(118),
		45: uint16(64),
		47: uint16(12),
		51: uint16(11),
		53: uint16(12),
		54: uint16(66),
		55: uint16(74),
		56: uint16(12),
	},
}

var ts_small_parse_table = [2812]uint16_t{
	0:    uint16(15),
	1:    uint16(3),
	2:    uint16(1),
	3:    uint16(sym_comment),
	4:    uint16(19),
	5:    uint16(1),
	6:    uint16(anon_sym_LT_BANG),
	7:    uint16(21),
	8:    uint16(1),
	9:    uint16(anon_sym_LT),
	10:   uint16(23),
	11:   uint16(1),
	12:   uint16(anon_sym_LT_SLASH),
	13:   uint16(25),
	14:   uint16(1),
	15:   uint16(sym_text),
	16:   uint16(27),
	17:   uint16(1),
	18:   uint16(sym__implicit_end_tag),
	19:   uint16(29),
	20:   uint16(1),
	21:   uint16(sym__html_interpolation_start),
	22:   uint16(5),
	23:   uint16(1),
	24:   uint16(sym_start_tag),
	25:   uint16(82),
	26:   uint16(1),
	27:   uint16(sym_self_closing_style_tag),
	28:   uint16(83),
	29:   uint16(1),
	30:   uint16(sym_self_closing_script_tag),
	31:   uint16(85),
	32:   uint16(1),
	33:   uint16(sym_self_closing_tag),
	34:   uint16(91),
	35:   uint16(1),
	36:   uint16(sym_end_tag),
	37:   uint16(116),
	38:   uint16(1),
	39:   uint16(sym_script_start_tag),
	40:   uint16(117),
	41:   uint16(1),
	42:   uint16(sym_style_start_tag),
	43:   uint16(10),
	44:   uint16(8),
	45:   uint16(sym_doctype),
	46:   uint16(sym__node),
	47:   uint16(sym_element),
	48:   uint16(sym_script_element),
	49:   uint16(sym_style_element),
	50:   uint16(sym_erroneous_end_tag),
	51:   uint16(sym_html_interpolation),
	52:   uint16(aux_sym_document_repeat1),
	53:   uint16(15),
	54:   uint16(3),
	55:   uint16(1),
	56:   uint16(sym_comment),
	57:   uint16(19),
	58:   uint16(1),
	59:   uint16(anon_sym_LT_BANG),
	60:   uint16(21),
	61:   uint16(1),
	62:   uint16(anon_sym_LT),
	63:   uint16(25),
	64:   uint16(1),
	65:   uint16(sym_text),
	66:   uint16(29),
	67:   uint16(1),
	68:   uint16(sym__html_interpolation_start),
	69:   uint16(31),
	70:   uint16(1),
	71:   uint16(anon_sym_LT_SLASH),
	72:   uint16(33),
	73:   uint16(1),
	74:   uint16(sym__implicit_end_tag),
	75:   uint16(5),
	76:   uint16(1),
	77:   uint16(sym_start_tag),
	78:   uint16(71),
	79:   uint16(1),
	80:   uint16(sym_end_tag),
	81:   uint16(82),
	82:   uint16(1),
	83:   uint16(sym_self_closing_style_tag),
	84:   uint16(83),
	85:   uint16(1),
	86:   uint16(sym_self_closing_script_tag),
	87:   uint16(85),
	88:   uint16(1),
	89:   uint16(sym_self_closing_tag),
	90:   uint16(116),
	91:   uint16(1),
	92:   uint16(sym_script_start_tag),
	93:   uint16(117),
	94:   uint16(1),
	95:   uint16(sym_style_start_tag),
	96:   uint16(10),
	97:   uint16(8),
	98:   uint16(sym_doctype),
	99:   uint16(sym__node),
	100:  uint16(sym_element),
	101:  uint16(sym_script_element),
	102:  uint16(sym_style_element),
	103:  uint16(sym_erroneous_end_tag),
	104:  uint16(sym_html_interpolation),
	105:  uint16(aux_sym_document_repeat1),
	106:  uint16(15),
	107:  uint16(3),
	108:  uint16(1),
	109:  uint16(sym_comment),
	110:  uint16(19),
	111:  uint16(1),
	112:  uint16(anon_sym_LT_BANG),
	113:  uint16(21),
	114:  uint16(1),
	115:  uint16(anon_sym_LT),
	116:  uint16(29),
	117:  uint16(1),
	118:  uint16(sym__html_interpolation_start),
	119:  uint16(35),
	120:  uint16(1),
	121:  uint16(anon_sym_LT_SLASH),
	122:  uint16(37),
	123:  uint16(1),
	124:  uint16(sym_text),
	125:  uint16(39),
	126:  uint16(1),
	127:  uint16(sym__implicit_end_tag),
	128:  uint16(5),
	129:  uint16(1),
	130:  uint16(sym_start_tag),
	131:  uint16(46),
	132:  uint16(1),
	133:  uint16(sym_end_tag),
	134:  uint16(82),
	135:  uint16(1),
	136:  uint16(sym_self_closing_style_tag),
	137:  uint16(83),
	138:  uint16(1),
	139:  uint16(sym_self_closing_script_tag),
	140:  uint16(85),
	141:  uint16(1),
	142:  uint16(sym_self_closing_tag),
	143:  uint16(116),
	144:  uint16(1),
	145:  uint16(sym_script_start_tag),
	146:  uint16(117),
	147:  uint16(1),
	148:  uint16(sym_style_start_tag),
	149:  uint16(7),
	150:  uint16(8),
	151:  uint16(sym_doctype),
	152:  uint16(sym__node),
	153:  uint16(sym_element),
	154:  uint16(sym_script_element),
	155:  uint16(sym_style_element),
	156:  uint16(sym_erroneous_end_tag),
	157:  uint16(sym_html_interpolation),
	158:  uint16(aux_sym_document_repeat1),
	159:  uint16(15),
	160:  uint16(3),
	161:  uint16(1),
	162:  uint16(sym_comment),
	163:  uint16(19),
	164:  uint16(1),
	165:  uint16(anon_sym_LT_BANG),
	166:  uint16(21),
	167:  uint16(1),
	168:  uint16(anon_sym_LT),
	169:  uint16(29),
	170:  uint16(1),
	171:  uint16(sym__html_interpolation_start),
	172:  uint16(31),
	173:  uint16(1),
	174:  uint16(anon_sym_LT_SLASH),
	175:  uint16(41),
	176:  uint16(1),
	177:  uint16(sym_text),
	178:  uint16(43),
	179:  uint16(1),
	180:  uint16(sym__implicit_end_tag),
	181:  uint16(5),
	182:  uint16(1),
	183:  uint16(sym_start_tag),
	184:  uint16(80),
	185:  uint16(1),
	186:  uint16(sym_end_tag),
	187:  uint16(82),
	188:  uint16(1),
	189:  uint16(sym_self_closing_style_tag),
	190:  uint16(83),
	191:  uint16(1),
	192:  uint16(sym_self_closing_script_tag),
	193:  uint16(85),
	194:  uint16(1),
	195:  uint16(sym_self_closing_tag),
	196:  uint16(116),
	197:  uint16(1),
	198:  uint16(sym_script_start_tag),
	199:  uint16(117),
	200:  uint16(1),
	201:  uint16(sym_style_start_tag),
	202:  uint16(3),
	203:  uint16(8),
	204:  uint16(sym_doctype),
	205:  uint16(sym__node),
	206:  uint16(sym_element),
	207:  uint16(sym_script_element),
	208:  uint16(sym_style_element),
	209:  uint16(sym_erroneous_end_tag),
	210:  uint16(sym_html_interpolation),
	211:  uint16(aux_sym_document_repeat1),
	212:  uint16(15),
	213:  uint16(3),
	214:  uint16(1),
	215:  uint16(sym_comment),
	216:  uint16(19),
	217:  uint16(1),
	218:  uint16(anon_sym_LT_BANG),
	219:  uint16(21),
	220:  uint16(1),
	221:  uint16(anon_sym_LT),
	222:  uint16(23),
	223:  uint16(1),
	224:  uint16(anon_sym_LT_SLASH),
	225:  uint16(29),
	226:  uint16(1),
	227:  uint16(sym__html_interpolation_start),
	228:  uint16(45),
	229:  uint16(1),
	230:  uint16(sym_text),
	231:  uint16(47),
	232:  uint16(1),
	233:  uint16(sym__implicit_end_tag),
	234:  uint16(5),
	235:  uint16(1),
	236:  uint16(sym_start_tag),
	237:  uint16(82),
	238:  uint16(1),
	239:  uint16(sym_self_closing_style_tag),
	240:  uint16(83),
	241:  uint16(1),
	242:  uint16(sym_self_closing_script_tag),
	243:  uint16(85),
	244:  uint16(1),
	245:  uint16(sym_self_closing_tag),
	246:  uint16(102),
	247:  uint16(1),
	248:  uint16(sym_end_tag),
	249:  uint16(116),
	250:  uint16(1),
	251:  uint16(sym_script_start_tag),
	252:  uint16(117),
	253:  uint16(1),
	254:  uint16(sym_style_start_tag),
	255:  uint16(2),
	256:  uint16(8),
	257:  uint16(sym_doctype),
	258:  uint16(sym__node),
	259:  uint16(sym_element),
	260:  uint16(sym_script_element),
	261:  uint16(sym_style_element),
	262:  uint16(sym_erroneous_end_tag),
	263:  uint16(sym_html_interpolation),
	264:  uint16(aux_sym_document_repeat1),
	265:  uint16(15),
	266:  uint16(3),
	267:  uint16(1),
	268:  uint16(sym_comment),
	269:  uint16(19),
	270:  uint16(1),
	271:  uint16(anon_sym_LT_BANG),
	272:  uint16(21),
	273:  uint16(1),
	274:  uint16(anon_sym_LT),
	275:  uint16(25),
	276:  uint16(1),
	277:  uint16(sym_text),
	278:  uint16(29),
	279:  uint16(1),
	280:  uint16(sym__html_interpolation_start),
	281:  uint16(35),
	282:  uint16(1),
	283:  uint16(anon_sym_LT_SLASH),
	284:  uint16(49),
	285:  uint16(1),
	286:  uint16(sym__implicit_end_tag),
	287:  uint16(5),
	288:  uint16(1),
	289:  uint16(sym_start_tag),
	290:  uint16(57),
	291:  uint16(1),
	292:  uint16(sym_end_tag),
	293:  uint16(82),
	294:  uint16(1),
	295:  uint16(sym_self_closing_style_tag),
	296:  uint16(83),
	297:  uint16(1),
	298:  uint16(sym_self_closing_script_tag),
	299:  uint16(85),
	300:  uint16(1),
	301:  uint16(sym_self_closing_tag),
	302:  uint16(116),
	303:  uint16(1),
	304:  uint16(sym_script_start_tag),
	305:  uint16(117),
	306:  uint16(1),
	307:  uint16(sym_style_start_tag),
	308:  uint16(10),
	309:  uint16(8),
	310:  uint16(sym_doctype),
	311:  uint16(sym__node),
	312:  uint16(sym_element),
	313:  uint16(sym_script_element),
	314:  uint16(sym_style_element),
	315:  uint16(sym_erroneous_end_tag),
	316:  uint16(sym_html_interpolation),
	317:  uint16(aux_sym_document_repeat1),
	318:  uint16(14),
	319:  uint16(3),
	320:  uint16(1),
	321:  uint16(sym_comment),
	322:  uint16(51),
	323:  uint16(1),
	325:  uint16(53),
	326:  uint16(1),
	327:  uint16(anon_sym_LT_BANG),
	328:  uint16(56),
	329:  uint16(1),
	330:  uint16(anon_sym_LT),
	331:  uint16(59),
	332:  uint16(1),
	333:  uint16(anon_sym_LT_SLASH),
	334:  uint16(62),
	335:  uint16(1),
	336:  uint16(sym_text),
	337:  uint16(65),
	338:  uint16(1),
	339:  uint16(sym__html_interpolation_start),
	340:  uint16(4),
	341:  uint16(1),
	342:  uint16(sym_start_tag),
	343:  uint16(64),
	344:  uint16(1),
	345:  uint16(sym_self_closing_tag),
	346:  uint16(66),
	347:  uint16(1),
	348:  uint16(sym_self_closing_script_tag),
	349:  uint16(74),
	350:  uint16(1),
	351:  uint16(sym_self_closing_style_tag),
	352:  uint16(118),
	353:  uint16(1),
	354:  uint16(sym_style_start_tag),
	355:  uint16(119),
	356:  uint16(1),
	357:  uint16(sym_script_start_tag),
	358:  uint16(8),
	359:  uint16(8),
	360:  uint16(sym_doctype),
	361:  uint16(sym__node),
	362:  uint16(sym_element),
	363:  uint16(sym_script_element),
	364:  uint16(sym_style_element),
	365:  uint16(sym_erroneous_end_tag),
	366:  uint16(sym_html_interpolation),
	367:  uint16(aux_sym_document_repeat1),
	368:  uint16(14),
	369:  uint16(3),
	370:  uint16(1),
	371:  uint16(sym_comment),
	372:  uint16(7),
	373:  uint16(1),
	374:  uint16(anon_sym_LT_BANG),
	375:  uint16(9),
	376:  uint16(1),
	377:  uint16(anon_sym_LT),
	378:  uint16(11),
	379:  uint16(1),
	380:  uint16(anon_sym_LT_SLASH),
	381:  uint16(17),
	382:  uint16(1),
	383:  uint16(sym__html_interpolation_start),
	384:  uint16(68),
	385:  uint16(1),
	387:  uint16(70),
	388:  uint16(1),
	389:  uint16(sym_text),
	390:  uint16(4),
	391:  uint16(1),
	392:  uint16(sym_start_tag),
	393:  uint16(64),
	394:  uint16(1),
	395:  uint16(sym_self_closing_tag),
	396:  uint16(66),
	397:  uint16(1),
	398:  uint16(sym_self_closing_script_tag),
	399:  uint16(74),
	400:  uint16(1),
	401:  uint16(sym_self_closing_style_tag),
	402:  uint16(118),
	403:  uint16(1),
	404:  uint16(sym_style_start_tag),
	405:  uint16(119),
	406:  uint16(1),
	407:  uint16(sym_script_start_tag),
	408:  uint16(8),
	409:  uint16(8),
	410:  uint16(sym_doctype),
	411:  uint16(sym__node),
	412:  uint16(sym_element),
	413:  uint16(sym_script_element),
	414:  uint16(sym_style_element),
	415:  uint16(sym_erroneous_end_tag),
	416:  uint16(sym_html_interpolation),
	417:  uint16(aux_sym_document_repeat1),
	418:  uint16(14),
	419:  uint16(3),
	420:  uint16(1),
	421:  uint16(sym_comment),
	422:  uint16(51),
	423:  uint16(1),
	424:  uint16(sym__implicit_end_tag),
	425:  uint16(72),
	426:  uint16(1),
	427:  uint16(anon_sym_LT_BANG),
	428:  uint16(75),
	429:  uint16(1),
	430:  uint16(anon_sym_LT),
	431:  uint16(78),
	432:  uint16(1),
	433:  uint16(anon_sym_LT_SLASH),
	434:  uint16(81),
	435:  uint16(1),
	436:  uint16(sym_text),
	437:  uint16(84),
	438:  uint16(1),
	439:  uint16(sym__html_interpolation_start),
	440:  uint16(5),
	441:  uint16(1),
	442:  uint16(sym_start_tag),
	443:  uint16(82),
	444:  uint16(1),
	445:  uint16(sym_self_closing_style_tag),
	446:  uint16(83),
	447:  uint16(1),
	448:  uint16(sym_self_closing_script_tag),
	449:  uint16(85),
	450:  uint16(1),
	451:  uint16(sym_self_closing_tag),
	452:  uint16(116),
	453:  uint16(1),
	454:  uint16(sym_script_start_tag),
	455:  uint16(117),
	456:  uint16(1),
	457:  uint16(sym_style_start_tag),
	458:  uint16(10),
	459:  uint16(8),
	460:  uint16(sym_doctype),
	461:  uint16(sym__node),
	462:  uint16(sym_element),
	463:  uint16(sym_script_element),
	464:  uint16(sym_style_element),
	465:  uint16(sym_erroneous_end_tag),
	466:  uint16(sym_html_interpolation),
	467:  uint16(aux_sym_document_repeat1),
	468:  uint16(14),
	469:  uint16(3),
	470:  uint16(1),
	471:  uint16(sym_comment),
	472:  uint16(7),
	473:  uint16(1),
	474:  uint16(anon_sym_LT_BANG),
	475:  uint16(9),
	476:  uint16(1),
	477:  uint16(anon_sym_LT),
	478:  uint16(11),
	479:  uint16(1),
	480:  uint16(anon_sym_LT_SLASH),
	481:  uint16(17),
	482:  uint16(1),
	483:  uint16(sym__html_interpolation_start),
	484:  uint16(87),
	485:  uint16(1),
	487:  uint16(89),
	488:  uint16(1),
	489:  uint16(sym_text),
	490:  uint16(4),
	491:  uint16(1),
	492:  uint16(sym_start_tag),
	493:  uint16(64),
	494:  uint16(1),
	495:  uint16(sym_self_closing_tag),
	496:  uint16(66),
	497:  uint16(1),
	498:  uint16(sym_self_closing_script_tag),
	499:  uint16(74),
	500:  uint16(1),
	501:  uint16(sym_self_closing_style_tag),
	502:  uint16(118),
	503:  uint16(1),
	504:  uint16(sym_style_start_tag),
	505:  uint16(119),
	506:  uint16(1),
	507:  uint16(sym_script_start_tag),
	508:  uint16(9),
	509:  uint16(8),
	510:  uint16(sym_doctype),
	511:  uint16(sym__node),
	512:  uint16(sym_element),
	513:  uint16(sym_script_element),
	514:  uint16(sym_style_element),
	515:  uint16(sym_erroneous_end_tag),
	516:  uint16(sym_html_interpolation),
	517:  uint16(aux_sym_document_repeat1),
	518:  uint16(14),
	519:  uint16(3),
	520:  uint16(1),
	521:  uint16(sym_comment),
	522:  uint16(7),
	523:  uint16(1),
	524:  uint16(anon_sym_LT_BANG),
	525:  uint16(9),
	526:  uint16(1),
	527:  uint16(anon_sym_LT),
	528:  uint16(11),
	529:  uint16(1),
	530:  uint16(anon_sym_LT_SLASH),
	531:  uint16(17),
	532:  uint16(1),
	533:  uint16(sym__html_interpolation_start),
	534:  uint16(70),
	535:  uint16(1),
	536:  uint16(sym_text),
	537:  uint16(87),
	538:  uint16(1),
	540:  uint16(4),
	541:  uint16(1),
	542:  uint16(sym_start_tag),
	543:  uint16(64),
	544:  uint16(1),
	545:  uint16(sym_self_closing_tag),
	546:  uint16(66),
	547:  uint16(1),
	548:  uint16(sym_self_closing_script_tag),
	549:  uint16(74),
	550:  uint16(1),
	551:  uint16(sym_self_closing_style_tag),
	552:  uint16(118),
	553:  uint16(1),
	554:  uint16(sym_style_start_tag),
	555:  uint16(119),
	556:  uint16(1),
	557:  uint16(sym_script_start_tag),
	558:  uint16(8),
	559:  uint16(8),
	560:  uint16(sym_doctype),
	561:  uint16(sym__node),
	562:  uint16(sym_element),
	563:  uint16(sym_script_element),
	564:  uint16(sym_style_element),
	565:  uint16(sym_erroneous_end_tag),
	566:  uint16(sym_html_interpolation),
	567:  uint16(aux_sym_document_repeat1),
	568:  uint16(13),
	569:  uint16(3),
	570:  uint16(1),
	571:  uint16(sym_comment),
	572:  uint16(91),
	573:  uint16(1),
	574:  uint16(anon_sym_LT_BANG),
	575:  uint16(93),
	576:  uint16(1),
	577:  uint16(anon_sym_LT),
	578:  uint16(95),
	579:  uint16(1),
	580:  uint16(sym__html_interpolation_start),
	581:  uint16(97),
	582:  uint16(1),
	583:  uint16(sym__html_interpolation_end),
	584:  uint16(99),
	585:  uint16(1),
	586:  uint16(sym_permissible_text),
	587:  uint16(6),
	588:  uint16(1),
	589:  uint16(sym_start_tag),
	590:  uint16(87),
	591:  uint16(1),
	592:  uint16(sym_self_closing_style_tag),
	593:  uint16(88),
	594:  uint16(1),
	595:  uint16(sym_self_closing_script_tag),
	596:  uint16(96),
	597:  uint16(1),
	598:  uint16(sym_self_closing_tag),
	599:  uint16(121),
	600:  uint16(1),
	601:  uint16(sym_script_start_tag),
	602:  uint16(122),
	603:  uint16(1),
	604:  uint16(sym_style_start_tag),
	605:  uint16(18),
	606:  uint16(7),
	607:  uint16(sym_doctype),
	608:  uint16(sym_element),
	609:  uint16(sym_script_element),
	610:  uint16(sym_style_element),
	611:  uint16(sym__node_with_permissible_text),
	612:  uint16(sym_html_interpolation),
	613:  uint16(aux_sym_html_interpolation_repeat1),
	614:  uint16(13),
	615:  uint16(3),
	616:  uint16(1),
	617:  uint16(sym_comment),
	618:  uint16(101),
	619:  uint16(1),
	620:  uint16(anon_sym_LT_BANG),
	621:  uint16(104),
	622:  uint16(1),
	623:  uint16(anon_sym_LT),
	624:  uint16(107),
	625:  uint16(1),
	626:  uint16(sym__html_interpolation_start),
	627:  uint16(110),
	628:  uint16(1),
	629:  uint16(sym__html_interpolation_end),
	630:  uint16(112),
	631:  uint16(1),
	632:  uint16(sym_permissible_text),
	633:  uint16(6),
	634:  uint16(1),
	635:  uint16(sym_start_tag),
	636:  uint16(87),
	637:  uint16(1),
	638:  uint16(sym_self_closing_style_tag),
	639:  uint16(88),
	640:  uint16(1),
	641:  uint16(sym_self_closing_script_tag),
	642:  uint16(96),
	643:  uint16(1),
	644:  uint16(sym_self_closing_tag),
	645:  uint16(121),
	646:  uint16(1),
	647:  uint16(sym_script_start_tag),
	648:  uint16(122),
	649:  uint16(1),
	650:  uint16(sym_style_start_tag),
	651:  uint16(14),
	652:  uint16(7),
	653:  uint16(sym_doctype),
	654:  uint16(sym_element),
	655:  uint16(sym_script_element),
	656:  uint16(sym_style_element),
	657:  uint16(sym__node_with_permissible_text),
	658:  uint16(sym_html_interpolation),
	659:  uint16(aux_sym_html_interpolation_repeat1),
	660:  uint16(13),
	661:  uint16(3),
	662:  uint16(1),
	663:  uint16(sym_comment),
	664:  uint16(91),
	665:  uint16(1),
	666:  uint16(anon_sym_LT_BANG),
	667:  uint16(93),
	668:  uint16(1),
	669:  uint16(anon_sym_LT),
	670:  uint16(95),
	671:  uint16(1),
	672:  uint16(sym__html_interpolation_start),
	673:  uint16(115),
	674:  uint16(1),
	675:  uint16(sym__html_interpolation_end),
	676:  uint16(117),
	677:  uint16(1),
	678:  uint16(sym_permissible_text),
	679:  uint16(6),
	680:  uint16(1),
	681:  uint16(sym_start_tag),
	682:  uint16(87),
	683:  uint16(1),
	684:  uint16(sym_self_closing_style_tag),
	685:  uint16(88),
	686:  uint16(1),
	687:  uint16(sym_self_closing_script_tag),
	688:  uint16(96),
	689:  uint16(1),
	690:  uint16(sym_self_closing_tag),
	691:  uint16(121),
	692:  uint16(1),
	693:  uint16(sym_script_start_tag),
	694:  uint16(122),
	695:  uint16(1),
	696:  uint16(sym_style_start_tag),
	697:  uint16(19),
	698:  uint16(7),
	699:  uint16(sym_doctype),
	700:  uint16(sym_element),
	701:  uint16(sym_script_element),
	702:  uint16(sym_style_element),
	703:  uint16(sym__node_with_permissible_text),
	704:  uint16(sym_html_interpolation),
	705:  uint16(aux_sym_html_interpolation_repeat1),
	706:  uint16(13),
	707:  uint16(3),
	708:  uint16(1),
	709:  uint16(sym_comment),
	710:  uint16(91),
	711:  uint16(1),
	712:  uint16(anon_sym_LT_BANG),
	713:  uint16(93),
	714:  uint16(1),
	715:  uint16(anon_sym_LT),
	716:  uint16(95),
	717:  uint16(1),
	718:  uint16(sym__html_interpolation_start),
	719:  uint16(119),
	720:  uint16(1),
	721:  uint16(sym__html_interpolation_end),
	722:  uint16(121),
	723:  uint16(1),
	724:  uint16(sym_permissible_text),
	725:  uint16(6),
	726:  uint16(1),
	727:  uint16(sym_start_tag),
	728:  uint16(87),
	729:  uint16(1),
	730:  uint16(sym_self_closing_style_tag),
	731:  uint16(88),
	732:  uint16(1),
	733:  uint16(sym_self_closing_script_tag),
	734:  uint16(96),
	735:  uint16(1),
	736:  uint16(sym_self_closing_tag),
	737:  uint16(121),
	738:  uint16(1),
	739:  uint16(sym_script_start_tag),
	740:  uint16(122),
	741:  uint16(1),
	742:  uint16(sym_style_start_tag),
	743:  uint16(14),
	744:  uint16(7),
	745:  uint16(sym_doctype),
	746:  uint16(sym_element),
	747:  uint16(sym_script_element),
	748:  uint16(sym_style_element),
	749:  uint16(sym__node_with_permissible_text),
	750:  uint16(sym_html_interpolation),
	751:  uint16(aux_sym_html_interpolation_repeat1),
	752:  uint16(13),
	753:  uint16(3),
	754:  uint16(1),
	755:  uint16(sym_comment),
	756:  uint16(91),
	757:  uint16(1),
	758:  uint16(anon_sym_LT_BANG),
	759:  uint16(93),
	760:  uint16(1),
	761:  uint16(anon_sym_LT),
	762:  uint16(95),
	763:  uint16(1),
	764:  uint16(sym__html_interpolation_start),
	765:  uint16(123),
	766:  uint16(1),
	767:  uint16(sym__html_interpolation_end),
	768:  uint16(125),
	769:  uint16(1),
	770:  uint16(sym_permissible_text),
	771:  uint16(6),
	772:  uint16(1),
	773:  uint16(sym_start_tag),
	774:  uint16(87),
	775:  uint16(1),
	776:  uint16(sym_self_closing_style_tag),
	777:  uint16(88),
	778:  uint16(1),
	779:  uint16(sym_self_closing_script_tag),
	780:  uint16(96),
	781:  uint16(1),
	782:  uint16(sym_self_closing_tag),
	783:  uint16(121),
	784:  uint16(1),
	785:  uint16(sym_script_start_tag),
	786:  uint16(122),
	787:  uint16(1),
	788:  uint16(sym_style_start_tag),
	789:  uint16(16),
	790:  uint16(7),
	791:  uint16(sym_doctype),
	792:  uint16(sym_element),
	793:  uint16(sym_script_element),
	794:  uint16(sym_style_element),
	795:  uint16(sym__node_with_permissible_text),
	796:  uint16(sym_html_interpolation),
	797:  uint16(aux_sym_html_interpolation_repeat1),
	798:  uint16(13),
	799:  uint16(3),
	800:  uint16(1),
	801:  uint16(sym_comment),
	802:  uint16(91),
	803:  uint16(1),
	804:  uint16(anon_sym_LT_BANG),
	805:  uint16(93),
	806:  uint16(1),
	807:  uint16(anon_sym_LT),
	808:  uint16(95),
	809:  uint16(1),
	810:  uint16(sym__html_interpolation_start),
	811:  uint16(121),
	812:  uint16(1),
	813:  uint16(sym_permissible_text),
	814:  uint16(127),
	815:  uint16(1),
	816:  uint16(sym__html_interpolation_end),
	817:  uint16(6),
	818:  uint16(1),
	819:  uint16(sym_start_tag),
	820:  uint16(87),
	821:  uint16(1),
	822:  uint16(sym_self_closing_style_tag),
	823:  uint16(88),
	824:  uint16(1),
	825:  uint16(sym_self_closing_script_tag),
	826:  uint16(96),
	827:  uint16(1),
	828:  uint16(sym_self_closing_tag),
	829:  uint16(121),
	830:  uint16(1),
	831:  uint16(sym_script_start_tag),
	832:  uint16(122),
	833:  uint16(1),
	834:  uint16(sym_style_start_tag),
	835:  uint16(14),
	836:  uint16(7),
	837:  uint16(sym_doctype),
	838:  uint16(sym_element),
	839:  uint16(sym_script_element),
	840:  uint16(sym_style_element),
	841:  uint16(sym__node_with_permissible_text),
	842:  uint16(sym_html_interpolation),
	843:  uint16(aux_sym_html_interpolation_repeat1),
	844:  uint16(13),
	845:  uint16(3),
	846:  uint16(1),
	847:  uint16(sym_comment),
	848:  uint16(91),
	849:  uint16(1),
	850:  uint16(anon_sym_LT_BANG),
	851:  uint16(93),
	852:  uint16(1),
	853:  uint16(anon_sym_LT),
	854:  uint16(95),
	855:  uint16(1),
	856:  uint16(sym__html_interpolation_start),
	857:  uint16(121),
	858:  uint16(1),
	859:  uint16(sym_permissible_text),
	860:  uint16(129),
	861:  uint16(1),
	862:  uint16(sym__html_interpolation_end),
	863:  uint16(6),
	864:  uint16(1),
	865:  uint16(sym_start_tag),
	866:  uint16(87),
	867:  uint16(1),
	868:  uint16(sym_self_closing_style_tag),
	869:  uint16(88),
	870:  uint16(1),
	871:  uint16(sym_self_closing_script_tag),
	872:  uint16(96),
	873:  uint16(1),
	874:  uint16(sym_self_closing_tag),
	875:  uint16(121),
	876:  uint16(1),
	877:  uint16(sym_script_start_tag),
	878:  uint16(122),
	879:  uint16(1),
	880:  uint16(sym_style_start_tag),
	881:  uint16(14),
	882:  uint16(7),
	883:  uint16(sym_doctype),
	884:  uint16(sym_element),
	885:  uint16(sym_script_element),
	886:  uint16(sym_style_element),
	887:  uint16(sym__node_with_permissible_text),
	888:  uint16(sym_html_interpolation),
	889:  uint16(aux_sym_html_interpolation_repeat1),
	890:  uint16(7),
	891:  uint16(3),
	892:  uint16(1),
	893:  uint16(sym_comment),
	894:  uint16(131),
	895:  uint16(1),
	896:  uint16(anon_sym_GT),
	897:  uint16(133),
	898:  uint16(1),
	899:  uint16(anon_sym_SLASH_GT),
	900:  uint16(135),
	901:  uint16(1),
	902:  uint16(sym_attribute_name),
	903:  uint16(137),
	904:  uint16(1),
	905:  uint16(anon_sym_LBRACE),
	906:  uint16(113),
	907:  uint16(1),
	908:  uint16(sym_attribute_interpolation),
	909:  uint16(35),
	910:  uint16(2),
	911:  uint16(sym_attribute),
	912:  uint16(aux_sym_start_tag_repeat1),
	913:  uint16(7),
	914:  uint16(3),
	915:  uint16(1),
	916:  uint16(sym_comment),
	917:  uint16(135),
	918:  uint16(1),
	919:  uint16(sym_attribute_name),
	920:  uint16(137),
	921:  uint16(1),
	922:  uint16(anon_sym_LBRACE),
	923:  uint16(139),
	924:  uint16(1),
	925:  uint16(anon_sym_GT),
	926:  uint16(141),
	927:  uint16(1),
	928:  uint16(anon_sym_SLASH_GT),
	929:  uint16(113),
	930:  uint16(1),
	931:  uint16(sym_attribute_interpolation),
	932:  uint16(32),
	933:  uint16(2),
	934:  uint16(sym_attribute),
	935:  uint16(aux_sym_start_tag_repeat1),
	936:  uint16(7),
	937:  uint16(3),
	938:  uint16(1),
	939:  uint16(sym_comment),
	940:  uint16(135),
	941:  uint16(1),
	942:  uint16(sym_attribute_name),
	943:  uint16(137),
	944:  uint16(1),
	945:  uint16(anon_sym_LBRACE),
	946:  uint16(143),
	947:  uint16(1),
	948:  uint16(anon_sym_GT),
	949:  uint16(145),
	950:  uint16(1),
	951:  uint16(anon_sym_SLASH_GT),
	952:  uint16(113),
	953:  uint16(1),
	954:  uint16(sym_attribute_interpolation),
	955:  uint16(39),
	956:  uint16(2),
	957:  uint16(sym_attribute),
	958:  uint16(aux_sym_start_tag_repeat1),
	959:  uint16(7),
	960:  uint16(3),
	961:  uint16(1),
	962:  uint16(sym_comment),
	963:  uint16(135),
	964:  uint16(1),
	965:  uint16(sym_attribute_name),
	966:  uint16(137),
	967:  uint16(1),
	968:  uint16(anon_sym_LBRACE),
	969:  uint16(147),
	970:  uint16(1),
	971:  uint16(anon_sym_GT),
	972:  uint16(149),
	973:  uint16(1),
	974:  uint16(anon_sym_SLASH_GT),
	975:  uint16(113),
	976:  uint16(1),
	977:  uint16(sym_attribute_interpolation),
	978:  uint16(33),
	979:  uint16(2),
	980:  uint16(sym_attribute),
	981:  uint16(aux_sym_start_tag_repeat1),
	982:  uint16(7),
	983:  uint16(3),
	984:  uint16(1),
	985:  uint16(sym_comment),
	986:  uint16(135),
	987:  uint16(1),
	988:  uint16(sym_attribute_name),
	989:  uint16(137),
	990:  uint16(1),
	991:  uint16(anon_sym_LBRACE),
	992:  uint16(139),
	993:  uint16(1),
	994:  uint16(anon_sym_GT),
	995:  uint16(151),
	996:  uint16(1),
	997:  uint16(anon_sym_SLASH_GT),
	998:  uint16(113),
	999:  uint16(1),
	1000: uint16(sym_attribute_interpolation),
	1001: uint16(30),
	1002: uint16(2),
	1003: uint16(sym_attribute),
	1004: uint16(aux_sym_start_tag_repeat1),
	1005: uint16(7),
	1006: uint16(3),
	1007: uint16(1),
	1008: uint16(sym_comment),
	1009: uint16(135),
	1010: uint16(1),
	1011: uint16(sym_attribute_name),
	1012: uint16(137),
	1013: uint16(1),
	1014: uint16(anon_sym_LBRACE),
	1015: uint16(147),
	1016: uint16(1),
	1017: uint16(anon_sym_GT),
	1018: uint16(153),
	1019: uint16(1),
	1020: uint16(anon_sym_SLASH_GT),
	1021: uint16(113),
	1022: uint16(1),
	1023: uint16(sym_attribute_interpolation),
	1024: uint16(29),
	1025: uint16(2),
	1026: uint16(sym_attribute),
	1027: uint16(aux_sym_start_tag_repeat1),
	1028: uint16(7),
	1029: uint16(3),
	1030: uint16(1),
	1031: uint16(sym_comment),
	1032: uint16(131),
	1033: uint16(1),
	1034: uint16(anon_sym_GT),
	1035: uint16(135),
	1036: uint16(1),
	1037: uint16(sym_attribute_name),
	1038: uint16(137),
	1039: uint16(1),
	1040: uint16(anon_sym_LBRACE),
	1041: uint16(155),
	1042: uint16(1),
	1043: uint16(anon_sym_SLASH_GT),
	1044: uint16(113),
	1045: uint16(1),
	1046: uint16(sym_attribute_interpolation),
	1047: uint16(22),
	1048: uint16(2),
	1049: uint16(sym_attribute),
	1050: uint16(aux_sym_start_tag_repeat1),
	1051: uint16(7),
	1052: uint16(3),
	1053: uint16(1),
	1054: uint16(sym_comment),
	1055: uint16(135),
	1056: uint16(1),
	1057: uint16(sym_attribute_name),
	1058: uint16(137),
	1059: uint16(1),
	1060: uint16(anon_sym_LBRACE),
	1061: uint16(157),
	1062: uint16(1),
	1063: uint16(anon_sym_GT),
	1064: uint16(159),
	1065: uint16(1),
	1066: uint16(anon_sym_SLASH_GT),
	1067: uint16(113),
	1068: uint16(1),
	1069: uint16(sym_attribute_interpolation),
	1070: uint16(39),
	1071: uint16(2),
	1072: uint16(sym_attribute),
	1073: uint16(aux_sym_start_tag_repeat1),
	1074: uint16(7),
	1075: uint16(3),
	1076: uint16(1),
	1077: uint16(sym_comment),
	1078: uint16(131),
	1079: uint16(1),
	1080: uint16(anon_sym_GT),
	1081: uint16(135),
	1082: uint16(1),
	1083: uint16(sym_attribute_name),
	1084: uint16(137),
	1085: uint16(1),
	1086: uint16(anon_sym_LBRACE),
	1087: uint16(161),
	1088: uint16(1),
	1089: uint16(anon_sym_SLASH_GT),
	1090: uint16(113),
	1091: uint16(1),
	1092: uint16(sym_attribute_interpolation),
	1093: uint16(38),
	1094: uint16(2),
	1095: uint16(sym_attribute),
	1096: uint16(aux_sym_start_tag_repeat1),
	1097: uint16(7),
	1098: uint16(3),
	1099: uint16(1),
	1100: uint16(sym_comment),
	1101: uint16(135),
	1102: uint16(1),
	1103: uint16(sym_attribute_name),
	1104: uint16(137),
	1105: uint16(1),
	1106: uint16(anon_sym_LBRACE),
	1107: uint16(163),
	1108: uint16(1),
	1109: uint16(anon_sym_GT),
	1110: uint16(165),
	1111: uint16(1),
	1112: uint16(anon_sym_SLASH_GT),
	1113: uint16(113),
	1114: uint16(1),
	1115: uint16(sym_attribute_interpolation),
	1116: uint16(39),
	1117: uint16(2),
	1118: uint16(sym_attribute),
	1119: uint16(aux_sym_start_tag_repeat1),
	1120: uint16(7),
	1121: uint16(3),
	1122: uint16(1),
	1123: uint16(sym_comment),
	1124: uint16(135),
	1125: uint16(1),
	1126: uint16(sym_attribute_name),
	1127: uint16(137),
	1128: uint16(1),
	1129: uint16(anon_sym_LBRACE),
	1130: uint16(157),
	1131: uint16(1),
	1132: uint16(anon_sym_GT),
	1133: uint16(167),
	1134: uint16(1),
	1135: uint16(anon_sym_SLASH_GT),
	1136: uint16(113),
	1137: uint16(1),
	1138: uint16(sym_attribute_interpolation),
	1139: uint16(39),
	1140: uint16(2),
	1141: uint16(sym_attribute),
	1142: uint16(aux_sym_start_tag_repeat1),
	1143: uint16(7),
	1144: uint16(3),
	1145: uint16(1),
	1146: uint16(sym_comment),
	1147: uint16(135),
	1148: uint16(1),
	1149: uint16(sym_attribute_name),
	1150: uint16(137),
	1151: uint16(1),
	1152: uint16(anon_sym_LBRACE),
	1153: uint16(147),
	1154: uint16(1),
	1155: uint16(anon_sym_GT),
	1156: uint16(169),
	1157: uint16(1),
	1158: uint16(anon_sym_SLASH_GT),
	1159: uint16(113),
	1160: uint16(1),
	1161: uint16(sym_attribute_interpolation),
	1162: uint16(37),
	1163: uint16(2),
	1164: uint16(sym_attribute),
	1165: uint16(aux_sym_start_tag_repeat1),
	1166: uint16(7),
	1167: uint16(3),
	1168: uint16(1),
	1169: uint16(sym_comment),
	1170: uint16(135),
	1171: uint16(1),
	1172: uint16(sym_attribute_name),
	1173: uint16(137),
	1174: uint16(1),
	1175: uint16(anon_sym_LBRACE),
	1176: uint16(157),
	1177: uint16(1),
	1178: uint16(anon_sym_GT),
	1179: uint16(171),
	1180: uint16(1),
	1181: uint16(anon_sym_SLASH_GT),
	1182: uint16(113),
	1183: uint16(1),
	1184: uint16(sym_attribute_interpolation),
	1185: uint16(39),
	1186: uint16(2),
	1187: uint16(sym_attribute),
	1188: uint16(aux_sym_start_tag_repeat1),
	1189: uint16(7),
	1190: uint16(3),
	1191: uint16(1),
	1192: uint16(sym_comment),
	1193: uint16(135),
	1194: uint16(1),
	1195: uint16(sym_attribute_name),
	1196: uint16(137),
	1197: uint16(1),
	1198: uint16(anon_sym_LBRACE),
	1199: uint16(163),
	1200: uint16(1),
	1201: uint16(anon_sym_GT),
	1202: uint16(173),
	1203: uint16(1),
	1204: uint16(anon_sym_SLASH_GT),
	1205: uint16(113),
	1206: uint16(1),
	1207: uint16(sym_attribute_interpolation),
	1208: uint16(39),
	1209: uint16(2),
	1210: uint16(sym_attribute),
	1211: uint16(aux_sym_start_tag_repeat1),
	1212: uint16(7),
	1213: uint16(3),
	1214: uint16(1),
	1215: uint16(sym_comment),
	1216: uint16(135),
	1217: uint16(1),
	1218: uint16(sym_attribute_name),
	1219: uint16(137),
	1220: uint16(1),
	1221: uint16(anon_sym_LBRACE),
	1222: uint16(139),
	1223: uint16(1),
	1224: uint16(anon_sym_GT),
	1225: uint16(175),
	1226: uint16(1),
	1227: uint16(anon_sym_SLASH_GT),
	1228: uint16(113),
	1229: uint16(1),
	1230: uint16(sym_attribute_interpolation),
	1231: uint16(27),
	1232: uint16(2),
	1233: uint16(sym_attribute),
	1234: uint16(aux_sym_start_tag_repeat1),
	1235: uint16(7),
	1236: uint16(3),
	1237: uint16(1),
	1238: uint16(sym_comment),
	1239: uint16(135),
	1240: uint16(1),
	1241: uint16(sym_attribute_name),
	1242: uint16(137),
	1243: uint16(1),
	1244: uint16(anon_sym_LBRACE),
	1245: uint16(143),
	1246: uint16(1),
	1247: uint16(anon_sym_GT),
	1248: uint16(177),
	1249: uint16(1),
	1250: uint16(anon_sym_SLASH_GT),
	1251: uint16(113),
	1252: uint16(1),
	1253: uint16(sym_attribute_interpolation),
	1254: uint16(39),
	1255: uint16(2),
	1256: uint16(sym_attribute),
	1257: uint16(aux_sym_start_tag_repeat1),
	1258: uint16(7),
	1259: uint16(3),
	1260: uint16(1),
	1261: uint16(sym_comment),
	1262: uint16(137),
	1263: uint16(1),
	1264: uint16(anon_sym_LBRACE),
	1265: uint16(179),
	1266: uint16(1),
	1267: uint16(sym_attribute_value),
	1268: uint16(181),
	1269: uint16(1),
	1270: uint16(anon_sym_SQUOTE),
	1271: uint16(183),
	1272: uint16(1),
	1273: uint16(anon_sym_DQUOTE),
	1274: uint16(185),
	1275: uint16(1),
	1276: uint16(sym_attribute_backtick_string),
	1277: uint16(111),
	1278: uint16(2),
	1279: uint16(sym_quoted_attribute_value),
	1280: uint16(sym_attribute_interpolation),
	1281: uint16(7),
	1282: uint16(3),
	1283: uint16(1),
	1284: uint16(sym_comment),
	1285: uint16(135),
	1286: uint16(1),
	1287: uint16(sym_attribute_name),
	1288: uint16(137),
	1289: uint16(1),
	1290: uint16(anon_sym_LBRACE),
	1291: uint16(163),
	1292: uint16(1),
	1293: uint16(anon_sym_GT),
	1294: uint16(187),
	1295: uint16(1),
	1296: uint16(anon_sym_SLASH_GT),
	1297: uint16(113),
	1298: uint16(1),
	1299: uint16(sym_attribute_interpolation),
	1300: uint16(39),
	1301: uint16(2),
	1302: uint16(sym_attribute),
	1303: uint16(aux_sym_start_tag_repeat1),
	1304: uint16(7),
	1305: uint16(3),
	1306: uint16(1),
	1307: uint16(sym_comment),
	1308: uint16(135),
	1309: uint16(1),
	1310: uint16(sym_attribute_name),
	1311: uint16(137),
	1312: uint16(1),
	1313: uint16(anon_sym_LBRACE),
	1314: uint16(143),
	1315: uint16(1),
	1316: uint16(anon_sym_GT),
	1317: uint16(189),
	1318: uint16(1),
	1319: uint16(anon_sym_SLASH_GT),
	1320: uint16(113),
	1321: uint16(1),
	1322: uint16(sym_attribute_interpolation),
	1323: uint16(39),
	1324: uint16(2),
	1325: uint16(sym_attribute),
	1326: uint16(aux_sym_start_tag_repeat1),
	1327: uint16(6),
	1328: uint16(3),
	1329: uint16(1),
	1330: uint16(sym_comment),
	1331: uint16(193),
	1332: uint16(1),
	1333: uint16(sym_attribute_name),
	1334: uint16(196),
	1335: uint16(1),
	1336: uint16(anon_sym_LBRACE),
	1337: uint16(113),
	1338: uint16(1),
	1339: uint16(sym_attribute_interpolation),
	1340: uint16(191),
	1341: uint16(2),
	1342: uint16(anon_sym_GT),
	1343: uint16(anon_sym_SLASH_GT),
	1344: uint16(39),
	1345: uint16(2),
	1346: uint16(sym_attribute),
	1347: uint16(aux_sym_start_tag_repeat1),
	1348: uint16(3),
	1349: uint16(3),
	1350: uint16(1),
	1351: uint16(sym_comment),
	1352: uint16(201),
	1353: uint16(1),
	1354: uint16(anon_sym_LT),
	1355: uint16(199),
	1356: uint16(5),
	1357: uint16(sym__html_interpolation_start),
	1359: uint16(anon_sym_LT_BANG),
	1360: uint16(anon_sym_LT_SLASH),
	1361: uint16(sym_text),
	1362: uint16(3),
	1363: uint16(3),
	1364: uint16(1),
	1365: uint16(sym_comment),
	1366: uint16(205),
	1367: uint16(1),
	1368: uint16(anon_sym_LT),
	1369: uint16(203),
	1370: uint16(5),
	1371: uint16(sym__implicit_end_tag),
	1372: uint16(sym__html_interpolation_start),
	1373: uint16(anon_sym_LT_BANG),
	1374: uint16(anon_sym_LT_SLASH),
	1375: uint16(sym_text),
	1376: uint16(3),
	1377: uint16(3),
	1378: uint16(1),
	1379: uint16(sym_comment),
	1380: uint16(209),
	1381: uint16(1),
	1382: uint16(anon_sym_LT),
	1383: uint16(207),
	1384: uint16(5),
	1385: uint16(sym__html_interpolation_start),
	1387: uint16(anon_sym_LT_BANG),
	1388: uint16(anon_sym_LT_SLASH),
	1389: uint16(sym_text),
	1390: uint16(3),
	1391: uint16(3),
	1392: uint16(1),
	1393: uint16(sym_comment),
	1394: uint16(213),
	1395: uint16(1),
	1396: uint16(anon_sym_LT),
	1397: uint16(211),
	1398: uint16(5),
	1399: uint16(sym__html_interpolation_start),
	1401: uint16(anon_sym_LT_BANG),
	1402: uint16(anon_sym_LT_SLASH),
	1403: uint16(sym_text),
	1404: uint16(3),
	1405: uint16(3),
	1406: uint16(1),
	1407: uint16(sym_comment),
	1408: uint16(217),
	1409: uint16(1),
	1410: uint16(anon_sym_LT),
	1411: uint16(215),
	1412: uint16(5),
	1413: uint16(sym__html_interpolation_start),
	1415: uint16(anon_sym_LT_BANG),
	1416: uint16(anon_sym_LT_SLASH),
	1417: uint16(sym_text),
	1418: uint16(3),
	1419: uint16(3),
	1420: uint16(1),
	1421: uint16(sym_comment),
	1422: uint16(221),
	1423: uint16(1),
	1424: uint16(anon_sym_LT),
	1425: uint16(219),
	1426: uint16(5),
	1427: uint16(sym__implicit_end_tag),
	1428: uint16(sym__html_interpolation_start),
	1429: uint16(anon_sym_LT_BANG),
	1430: uint16(anon_sym_LT_SLASH),
	1431: uint16(sym_text),
	1432: uint16(3),
	1433: uint16(3),
	1434: uint16(1),
	1435: uint16(sym_comment),
	1436: uint16(225),
	1437: uint16(1),
	1438: uint16(anon_sym_LT),
	1439: uint16(223),
	1440: uint16(5),
	1441: uint16(sym__html_interpolation_start),
	1443: uint16(anon_sym_LT_BANG),
	1444: uint16(anon_sym_LT_SLASH),
	1445: uint16(sym_text),
	1446: uint16(3),
	1447: uint16(3),
	1448: uint16(1),
	1449: uint16(sym_comment),
	1450: uint16(205),
	1451: uint16(1),
	1452: uint16(anon_sym_LT),
	1453: uint16(203),
	1454: uint16(5),
	1455: uint16(sym__html_interpolation_start),
	1457: uint16(anon_sym_LT_BANG),
	1458: uint16(anon_sym_LT_SLASH),
	1459: uint16(sym_text),
	1460: uint16(3),
	1461: uint16(3),
	1462: uint16(1),
	1463: uint16(sym_comment),
	1464: uint16(229),
	1465: uint16(1),
	1466: uint16(anon_sym_LT),
	1467: uint16(227),
	1468: uint16(5),
	1469: uint16(sym__implicit_end_tag),
	1470: uint16(sym__html_interpolation_start),
	1471: uint16(anon_sym_LT_BANG),
	1472: uint16(anon_sym_LT_SLASH),
	1473: uint16(sym_text),
	1474: uint16(3),
	1475: uint16(3),
	1476: uint16(1),
	1477: uint16(sym_comment),
	1478: uint16(233),
	1479: uint16(1),
	1480: uint16(anon_sym_LT),
	1481: uint16(231),
	1482: uint16(5),
	1483: uint16(sym__implicit_end_tag),
	1484: uint16(sym__html_interpolation_start),
	1485: uint16(anon_sym_LT_BANG),
	1486: uint16(anon_sym_LT_SLASH),
	1487: uint16(sym_text),
	1488: uint16(3),
	1489: uint16(3),
	1490: uint16(1),
	1491: uint16(sym_comment),
	1492: uint16(237),
	1493: uint16(1),
	1494: uint16(anon_sym_LT),
	1495: uint16(235),
	1496: uint16(5),
	1497: uint16(sym__html_interpolation_start),
	1499: uint16(anon_sym_LT_BANG),
	1500: uint16(anon_sym_LT_SLASH),
	1501: uint16(sym_text),
	1502: uint16(3),
	1503: uint16(3),
	1504: uint16(1),
	1505: uint16(sym_comment),
	1506: uint16(241),
	1507: uint16(1),
	1508: uint16(anon_sym_LT),
	1509: uint16(239),
	1510: uint16(5),
	1511: uint16(sym__implicit_end_tag),
	1512: uint16(sym__html_interpolation_start),
	1513: uint16(anon_sym_LT_BANG),
	1514: uint16(anon_sym_LT_SLASH),
	1515: uint16(sym_text),
	1516: uint16(3),
	1517: uint16(3),
	1518: uint16(1),
	1519: uint16(sym_comment),
	1520: uint16(245),
	1521: uint16(1),
	1522: uint16(anon_sym_LT),
	1523: uint16(243),
	1524: uint16(5),
	1525: uint16(sym__html_interpolation_start),
	1527: uint16(anon_sym_LT_BANG),
	1528: uint16(anon_sym_LT_SLASH),
	1529: uint16(sym_text),
	1530: uint16(3),
	1531: uint16(3),
	1532: uint16(1),
	1533: uint16(sym_comment),
	1534: uint16(249),
	1535: uint16(1),
	1536: uint16(anon_sym_LT),
	1537: uint16(247),
	1538: uint16(5),
	1539: uint16(sym__html_interpolation_start),
	1541: uint16(anon_sym_LT_BANG),
	1542: uint16(anon_sym_LT_SLASH),
	1543: uint16(sym_text),
	1544: uint16(3),
	1545: uint16(3),
	1546: uint16(1),
	1547: uint16(sym_comment),
	1548: uint16(253),
	1549: uint16(1),
	1550: uint16(anon_sym_LT),
	1551: uint16(251),
	1552: uint16(5),
	1553: uint16(sym__html_interpolation_start),
	1555: uint16(anon_sym_LT_BANG),
	1556: uint16(anon_sym_LT_SLASH),
	1557: uint16(sym_text),
	1558: uint16(3),
	1559: uint16(3),
	1560: uint16(1),
	1561: uint16(sym_comment),
	1562: uint16(257),
	1563: uint16(1),
	1564: uint16(anon_sym_LT),
	1565: uint16(255),
	1566: uint16(5),
	1567: uint16(sym__implicit_end_tag),
	1568: uint16(sym__html_interpolation_start),
	1569: uint16(anon_sym_LT_BANG),
	1570: uint16(anon_sym_LT_SLASH),
	1571: uint16(sym_text),
	1572: uint16(3),
	1573: uint16(3),
	1574: uint16(1),
	1575: uint16(sym_comment),
	1576: uint16(261),
	1577: uint16(1),
	1578: uint16(anon_sym_LT),
	1579: uint16(259),
	1580: uint16(5),
	1581: uint16(sym__html_interpolation_start),
	1583: uint16(anon_sym_LT_BANG),
	1584: uint16(anon_sym_LT_SLASH),
	1585: uint16(sym_text),
	1586: uint16(3),
	1587: uint16(3),
	1588: uint16(1),
	1589: uint16(sym_comment),
	1590: uint16(265),
	1591: uint16(1),
	1592: uint16(anon_sym_LT),
	1593: uint16(263),
	1594: uint16(5),
	1595: uint16(sym__html_interpolation_start),
	1597: uint16(anon_sym_LT_BANG),
	1598: uint16(anon_sym_LT_SLASH),
	1599: uint16(sym_text),
	1600: uint16(3),
	1601: uint16(3),
	1602: uint16(1),
	1603: uint16(sym_comment),
	1604: uint16(269),
	1605: uint16(1),
	1606: uint16(anon_sym_LT),
	1607: uint16(267),
	1608: uint16(5),
	1609: uint16(sym__html_interpolation_start),
	1611: uint16(anon_sym_LT_BANG),
	1612: uint16(anon_sym_LT_SLASH),
	1613: uint16(sym_text),
	1614: uint16(3),
	1615: uint16(3),
	1616: uint16(1),
	1617: uint16(sym_comment),
	1618: uint16(273),
	1619: uint16(1),
	1620: uint16(anon_sym_LT),
	1621: uint16(271),
	1622: uint16(5),
	1623: uint16(sym__html_interpolation_start),
	1625: uint16(anon_sym_LT_BANG),
	1626: uint16(anon_sym_LT_SLASH),
	1627: uint16(sym_text),
	1628: uint16(3),
	1629: uint16(3),
	1630: uint16(1),
	1631: uint16(sym_comment),
	1632: uint16(277),
	1633: uint16(1),
	1634: uint16(anon_sym_LT),
	1635: uint16(275),
	1636: uint16(5),
	1637: uint16(sym__html_interpolation_start),
	1639: uint16(anon_sym_LT_BANG),
	1640: uint16(anon_sym_LT_SLASH),
	1641: uint16(sym_text),
	1642: uint16(3),
	1643: uint16(3),
	1644: uint16(1),
	1645: uint16(sym_comment),
	1646: uint16(277),
	1647: uint16(1),
	1648: uint16(anon_sym_LT),
	1649: uint16(275),
	1650: uint16(5),
	1651: uint16(sym__implicit_end_tag),
	1652: uint16(sym__html_interpolation_start),
	1653: uint16(anon_sym_LT_BANG),
	1654: uint16(anon_sym_LT_SLASH),
	1655: uint16(sym_text),
	1656: uint16(3),
	1657: uint16(3),
	1658: uint16(1),
	1659: uint16(sym_comment),
	1660: uint16(281),
	1661: uint16(1),
	1662: uint16(anon_sym_LT),
	1663: uint16(279),
	1664: uint16(5),
	1665: uint16(sym__implicit_end_tag),
	1666: uint16(sym__html_interpolation_start),
	1667: uint16(anon_sym_LT_BANG),
	1668: uint16(anon_sym_LT_SLASH),
	1669: uint16(sym_text),
	1670: uint16(3),
	1671: uint16(3),
	1672: uint16(1),
	1673: uint16(sym_comment),
	1674: uint16(257),
	1675: uint16(1),
	1676: uint16(anon_sym_LT),
	1677: uint16(255),
	1678: uint16(5),
	1679: uint16(sym__html_interpolation_start),
	1681: uint16(anon_sym_LT_BANG),
	1682: uint16(anon_sym_LT_SLASH),
	1683: uint16(sym_text),
	1684: uint16(3),
	1685: uint16(3),
	1686: uint16(1),
	1687: uint16(sym_comment),
	1688: uint16(285),
	1689: uint16(1),
	1690: uint16(anon_sym_LT),
	1691: uint16(283),
	1692: uint16(5),
	1693: uint16(sym__html_interpolation_start),
	1695: uint16(anon_sym_LT_BANG),
	1696: uint16(anon_sym_LT_SLASH),
	1697: uint16(sym_text),
	1698: uint16(3),
	1699: uint16(3),
	1700: uint16(1),
	1701: uint16(sym_comment),
	1702: uint16(241),
	1703: uint16(1),
	1704: uint16(anon_sym_LT),
	1705: uint16(239),
	1706: uint16(5),
	1707: uint16(sym__html_interpolation_start),
	1709: uint16(anon_sym_LT_BANG),
	1710: uint16(anon_sym_LT_SLASH),
	1711: uint16(sym_text),
	1712: uint16(3),
	1713: uint16(3),
	1714: uint16(1),
	1715: uint16(sym_comment),
	1716: uint16(289),
	1717: uint16(1),
	1718: uint16(anon_sym_LT),
	1719: uint16(287),
	1720: uint16(5),
	1721: uint16(sym__html_interpolation_start),
	1723: uint16(anon_sym_LT_BANG),
	1724: uint16(anon_sym_LT_SLASH),
	1725: uint16(sym_text),
	1726: uint16(3),
	1727: uint16(3),
	1728: uint16(1),
	1729: uint16(sym_comment),
	1730: uint16(233),
	1731: uint16(1),
	1732: uint16(anon_sym_LT),
	1733: uint16(231),
	1734: uint16(5),
	1735: uint16(sym__html_interpolation_start),
	1737: uint16(anon_sym_LT_BANG),
	1738: uint16(anon_sym_LT_SLASH),
	1739: uint16(sym_text),
	1740: uint16(3),
	1741: uint16(3),
	1742: uint16(1),
	1743: uint16(sym_comment),
	1744: uint16(229),
	1745: uint16(1),
	1746: uint16(anon_sym_LT),
	1747: uint16(227),
	1748: uint16(5),
	1749: uint16(sym__html_interpolation_start),
	1751: uint16(anon_sym_LT_BANG),
	1752: uint16(anon_sym_LT_SLASH),
	1753: uint16(sym_text),
	1754: uint16(3),
	1755: uint16(3),
	1756: uint16(1),
	1757: uint16(sym_comment),
	1758: uint16(273),
	1759: uint16(1),
	1760: uint16(anon_sym_LT),
	1761: uint16(271),
	1762: uint16(5),
	1763: uint16(sym__implicit_end_tag),
	1764: uint16(sym__html_interpolation_start),
	1765: uint16(anon_sym_LT_BANG),
	1766: uint16(anon_sym_LT_SLASH),
	1767: uint16(sym_text),
	1768: uint16(3),
	1769: uint16(3),
	1770: uint16(1),
	1771: uint16(sym_comment),
	1772: uint16(269),
	1773: uint16(1),
	1774: uint16(anon_sym_LT),
	1775: uint16(267),
	1776: uint16(5),
	1777: uint16(sym__implicit_end_tag),
	1778: uint16(sym__html_interpolation_start),
	1779: uint16(anon_sym_LT_BANG),
	1780: uint16(anon_sym_LT_SLASH),
	1781: uint16(sym_text),
	1782: uint16(3),
	1783: uint16(3),
	1784: uint16(1),
	1785: uint16(sym_comment),
	1786: uint16(265),
	1787: uint16(1),
	1788: uint16(anon_sym_LT),
	1789: uint16(263),
	1790: uint16(5),
	1791: uint16(sym__implicit_end_tag),
	1792: uint16(sym__html_interpolation_start),
	1793: uint16(anon_sym_LT_BANG),
	1794: uint16(anon_sym_LT_SLASH),
	1795: uint16(sym_text),
	1796: uint16(3),
	1797: uint16(3),
	1798: uint16(1),
	1799: uint16(sym_comment),
	1800: uint16(261),
	1801: uint16(1),
	1802: uint16(anon_sym_LT),
	1803: uint16(259),
	1804: uint16(5),
	1805: uint16(sym__implicit_end_tag),
	1806: uint16(sym__html_interpolation_start),
	1807: uint16(anon_sym_LT_BANG),
	1808: uint16(anon_sym_LT_SLASH),
	1809: uint16(sym_text),
	1810: uint16(3),
	1811: uint16(3),
	1812: uint16(1),
	1813: uint16(sym_comment),
	1814: uint16(245),
	1815: uint16(1),
	1816: uint16(anon_sym_LT),
	1817: uint16(243),
	1818: uint16(5),
	1819: uint16(sym__implicit_end_tag),
	1820: uint16(sym__html_interpolation_start),
	1821: uint16(anon_sym_LT_BANG),
	1822: uint16(anon_sym_LT_SLASH),
	1823: uint16(sym_text),
	1824: uint16(3),
	1825: uint16(3),
	1826: uint16(1),
	1827: uint16(sym_comment),
	1828: uint16(293),
	1829: uint16(1),
	1830: uint16(anon_sym_LT),
	1831: uint16(291),
	1832: uint16(5),
	1833: uint16(sym__html_interpolation_start),
	1835: uint16(anon_sym_LT_BANG),
	1836: uint16(anon_sym_LT_SLASH),
	1837: uint16(sym_text),
	1838: uint16(3),
	1839: uint16(3),
	1840: uint16(1),
	1841: uint16(sym_comment),
	1842: uint16(297),
	1843: uint16(1),
	1844: uint16(anon_sym_LT),
	1845: uint16(295),
	1846: uint16(5),
	1847: uint16(sym__html_interpolation_start),
	1849: uint16(anon_sym_LT_BANG),
	1850: uint16(anon_sym_LT_SLASH),
	1851: uint16(sym_text),
	1852: uint16(3),
	1853: uint16(3),
	1854: uint16(1),
	1855: uint16(sym_comment),
	1856: uint16(237),
	1857: uint16(1),
	1858: uint16(anon_sym_LT),
	1859: uint16(235),
	1860: uint16(5),
	1861: uint16(sym__implicit_end_tag),
	1862: uint16(sym__html_interpolation_start),
	1863: uint16(anon_sym_LT_BANG),
	1864: uint16(anon_sym_LT_SLASH),
	1865: uint16(sym_text),
	1866: uint16(3),
	1867: uint16(3),
	1868: uint16(1),
	1869: uint16(sym_comment),
	1870: uint16(201),
	1871: uint16(1),
	1872: uint16(anon_sym_LT),
	1873: uint16(199),
	1874: uint16(5),
	1875: uint16(sym__implicit_end_tag),
	1876: uint16(sym__html_interpolation_start),
	1877: uint16(anon_sym_LT_BANG),
	1878: uint16(anon_sym_LT_SLASH),
	1879: uint16(sym_text),
	1880: uint16(3),
	1881: uint16(3),
	1882: uint16(1),
	1883: uint16(sym_comment),
	1884: uint16(213),
	1885: uint16(1),
	1886: uint16(anon_sym_LT),
	1887: uint16(211),
	1888: uint16(5),
	1889: uint16(sym__implicit_end_tag),
	1890: uint16(sym__html_interpolation_start),
	1891: uint16(anon_sym_LT_BANG),
	1892: uint16(anon_sym_LT_SLASH),
	1893: uint16(sym_text),
	1894: uint16(3),
	1895: uint16(3),
	1896: uint16(1),
	1897: uint16(sym_comment),
	1898: uint16(209),
	1899: uint16(1),
	1900: uint16(anon_sym_LT),
	1901: uint16(207),
	1902: uint16(5),
	1903: uint16(sym__implicit_end_tag),
	1904: uint16(sym__html_interpolation_start),
	1905: uint16(anon_sym_LT_BANG),
	1906: uint16(anon_sym_LT_SLASH),
	1907: uint16(sym_text),
	1908: uint16(3),
	1909: uint16(3),
	1910: uint16(1),
	1911: uint16(sym_comment),
	1912: uint16(225),
	1913: uint16(1),
	1914: uint16(anon_sym_LT),
	1915: uint16(223),
	1916: uint16(5),
	1917: uint16(sym__implicit_end_tag),
	1918: uint16(sym__html_interpolation_start),
	1919: uint16(anon_sym_LT_BANG),
	1920: uint16(anon_sym_LT_SLASH),
	1921: uint16(sym_text),
	1922: uint16(3),
	1923: uint16(3),
	1924: uint16(1),
	1925: uint16(sym_comment),
	1926: uint16(297),
	1927: uint16(1),
	1928: uint16(anon_sym_LT),
	1929: uint16(295),
	1930: uint16(5),
	1931: uint16(sym__implicit_end_tag),
	1932: uint16(sym__html_interpolation_start),
	1933: uint16(anon_sym_LT_BANG),
	1934: uint16(anon_sym_LT_SLASH),
	1935: uint16(sym_text),
	1936: uint16(3),
	1937: uint16(3),
	1938: uint16(1),
	1939: uint16(sym_comment),
	1940: uint16(293),
	1941: uint16(1),
	1942: uint16(anon_sym_LT),
	1943: uint16(291),
	1944: uint16(5),
	1945: uint16(sym__implicit_end_tag),
	1946: uint16(sym__html_interpolation_start),
	1947: uint16(anon_sym_LT_BANG),
	1948: uint16(anon_sym_LT_SLASH),
	1949: uint16(sym_text),
	1950: uint16(3),
	1951: uint16(3),
	1952: uint16(1),
	1953: uint16(sym_comment),
	1954: uint16(289),
	1955: uint16(1),
	1956: uint16(anon_sym_LT),
	1957: uint16(287),
	1958: uint16(5),
	1959: uint16(sym__implicit_end_tag),
	1960: uint16(sym__html_interpolation_start),
	1961: uint16(anon_sym_LT_BANG),
	1962: uint16(anon_sym_LT_SLASH),
	1963: uint16(sym_text),
	1964: uint16(3),
	1965: uint16(3),
	1966: uint16(1),
	1967: uint16(sym_comment),
	1968: uint16(301),
	1969: uint16(1),
	1970: uint16(anon_sym_LT),
	1971: uint16(299),
	1972: uint16(5),
	1973: uint16(sym__implicit_end_tag),
	1974: uint16(sym__html_interpolation_start),
	1975: uint16(anon_sym_LT_BANG),
	1976: uint16(anon_sym_LT_SLASH),
	1977: uint16(sym_text),
	1978: uint16(3),
	1979: uint16(3),
	1980: uint16(1),
	1981: uint16(sym_comment),
	1982: uint16(285),
	1983: uint16(1),
	1984: uint16(anon_sym_LT),
	1985: uint16(283),
	1986: uint16(5),
	1987: uint16(sym__implicit_end_tag),
	1988: uint16(sym__html_interpolation_start),
	1989: uint16(anon_sym_LT_BANG),
	1990: uint16(anon_sym_LT_SLASH),
	1991: uint16(sym_text),
	1992: uint16(3),
	1993: uint16(3),
	1994: uint16(1),
	1995: uint16(sym_comment),
	1996: uint16(249),
	1997: uint16(1),
	1998: uint16(anon_sym_LT),
	1999: uint16(247),
	2000: uint16(5),
	2001: uint16(sym__implicit_end_tag),
	2002: uint16(sym__html_interpolation_start),
	2003: uint16(anon_sym_LT_BANG),
	2004: uint16(anon_sym_LT_SLASH),
	2005: uint16(sym_text),
	2006: uint16(3),
	2007: uint16(3),
	2008: uint16(1),
	2009: uint16(sym_comment),
	2010: uint16(293),
	2011: uint16(1),
	2012: uint16(anon_sym_LT),
	2013: uint16(291),
	2014: uint16(4),
	2015: uint16(sym__html_interpolation_start),
	2016: uint16(sym__html_interpolation_end),
	2017: uint16(sym_permissible_text),
	2018: uint16(anon_sym_LT_BANG),
	2019: uint16(3),
	2020: uint16(3),
	2021: uint16(1),
	2022: uint16(sym_comment),
	2023: uint16(289),
	2024: uint16(1),
	2025: uint16(anon_sym_LT),
	2026: uint16(287),
	2027: uint16(4),
	2028: uint16(sym__html_interpolation_start),
	2029: uint16(sym__html_interpolation_end),
	2030: uint16(sym_permissible_text),
	2031: uint16(anon_sym_LT_BANG),
	2032: uint16(3),
	2033: uint16(3),
	2034: uint16(1),
	2035: uint16(sym_comment),
	2036: uint16(201),
	2037: uint16(1),
	2038: uint16(anon_sym_LT),
	2039: uint16(199),
	2040: uint16(4),
	2041: uint16(sym__html_interpolation_start),
	2042: uint16(sym__html_interpolation_end),
	2043: uint16(sym_permissible_text),
	2044: uint16(anon_sym_LT_BANG),
	2045: uint16(3),
	2046: uint16(3),
	2047: uint16(1),
	2048: uint16(sym_comment),
	2049: uint16(261),
	2050: uint16(1),
	2051: uint16(anon_sym_LT),
	2052: uint16(259),
	2053: uint16(4),
	2054: uint16(sym__html_interpolation_start),
	2055: uint16(sym__html_interpolation_end),
	2056: uint16(sym_permissible_text),
	2057: uint16(anon_sym_LT_BANG),
	2058: uint16(3),
	2059: uint16(3),
	2060: uint16(1),
	2061: uint16(sym_comment),
	2062: uint16(265),
	2063: uint16(1),
	2064: uint16(anon_sym_LT),
	2065: uint16(263),
	2066: uint16(4),
	2067: uint16(sym__html_interpolation_start),
	2068: uint16(sym__html_interpolation_end),
	2069: uint16(sym_permissible_text),
	2070: uint16(anon_sym_LT_BANG),
	2071: uint16(3),
	2072: uint16(3),
	2073: uint16(1),
	2074: uint16(sym_comment),
	2075: uint16(269),
	2076: uint16(1),
	2077: uint16(anon_sym_LT),
	2078: uint16(267),
	2079: uint16(4),
	2080: uint16(sym__html_interpolation_start),
	2081: uint16(sym__html_interpolation_end),
	2082: uint16(sym_permissible_text),
	2083: uint16(anon_sym_LT_BANG),
	2084: uint16(3),
	2085: uint16(3),
	2086: uint16(1),
	2087: uint16(sym_comment),
	2088: uint16(273),
	2089: uint16(1),
	2090: uint16(anon_sym_LT),
	2091: uint16(271),
	2092: uint16(4),
	2093: uint16(sym__html_interpolation_start),
	2094: uint16(sym__html_interpolation_end),
	2095: uint16(sym_permissible_text),
	2096: uint16(anon_sym_LT_BANG),
	2097: uint16(3),
	2098: uint16(3),
	2099: uint16(1),
	2100: uint16(sym_comment),
	2101: uint16(277),
	2102: uint16(1),
	2103: uint16(anon_sym_LT),
	2104: uint16(275),
	2105: uint16(4),
	2106: uint16(sym__html_interpolation_start),
	2107: uint16(sym__html_interpolation_end),
	2108: uint16(sym_permissible_text),
	2109: uint16(anon_sym_LT_BANG),
	2110: uint16(3),
	2111: uint16(3),
	2112: uint16(1),
	2113: uint16(sym_comment),
	2114: uint16(257),
	2115: uint16(1),
	2116: uint16(anon_sym_LT),
	2117: uint16(255),
	2118: uint16(4),
	2119: uint16(sym__html_interpolation_start),
	2120: uint16(sym__html_interpolation_end),
	2121: uint16(sym_permissible_text),
	2122: uint16(anon_sym_LT_BANG),
	2123: uint16(3),
	2124: uint16(3),
	2125: uint16(1),
	2126: uint16(sym_comment),
	2127: uint16(285),
	2128: uint16(1),
	2129: uint16(anon_sym_LT),
	2130: uint16(283),
	2131: uint16(4),
	2132: uint16(sym__html_interpolation_start),
	2133: uint16(sym__html_interpolation_end),
	2134: uint16(sym_permissible_text),
	2135: uint16(anon_sym_LT_BANG),
	2136: uint16(3),
	2137: uint16(3),
	2138: uint16(1),
	2139: uint16(sym_comment),
	2140: uint16(233),
	2141: uint16(1),
	2142: uint16(anon_sym_LT),
	2143: uint16(231),
	2144: uint16(4),
	2145: uint16(sym__html_interpolation_start),
	2146: uint16(sym__html_interpolation_end),
	2147: uint16(sym_permissible_text),
	2148: uint16(anon_sym_LT_BANG),
	2149: uint16(3),
	2150: uint16(3),
	2151: uint16(1),
	2152: uint16(sym_comment),
	2153: uint16(229),
	2154: uint16(1),
	2155: uint16(anon_sym_LT),
	2156: uint16(227),
	2157: uint16(4),
	2158: uint16(sym__html_interpolation_start),
	2159: uint16(sym__html_interpolation_end),
	2160: uint16(sym_permissible_text),
	2161: uint16(anon_sym_LT_BANG),
	2162: uint16(3),
	2163: uint16(3),
	2164: uint16(1),
	2165: uint16(sym_comment),
	2166: uint16(213),
	2167: uint16(1),
	2168: uint16(anon_sym_LT),
	2169: uint16(211),
	2170: uint16(4),
	2171: uint16(sym__html_interpolation_start),
	2172: uint16(sym__html_interpolation_end),
	2173: uint16(sym_permissible_text),
	2174: uint16(anon_sym_LT_BANG),
	2175: uint16(3),
	2176: uint16(3),
	2177: uint16(1),
	2178: uint16(sym_comment),
	2179: uint16(209),
	2180: uint16(1),
	2181: uint16(anon_sym_LT),
	2182: uint16(207),
	2183: uint16(4),
	2184: uint16(sym__html_interpolation_start),
	2185: uint16(sym__html_interpolation_end),
	2186: uint16(sym_permissible_text),
	2187: uint16(anon_sym_LT_BANG),
	2188: uint16(3),
	2189: uint16(3),
	2190: uint16(1),
	2191: uint16(sym_comment),
	2192: uint16(205),
	2193: uint16(1),
	2194: uint16(anon_sym_LT),
	2195: uint16(203),
	2196: uint16(4),
	2197: uint16(sym__html_interpolation_start),
	2198: uint16(sym__html_interpolation_end),
	2199: uint16(sym_permissible_text),
	2200: uint16(anon_sym_LT_BANG),
	2201: uint16(3),
	2202: uint16(3),
	2203: uint16(1),
	2204: uint16(sym_comment),
	2205: uint16(225),
	2206: uint16(1),
	2207: uint16(anon_sym_LT),
	2208: uint16(223),
	2209: uint16(4),
	2210: uint16(sym__html_interpolation_start),
	2211: uint16(sym__html_interpolation_end),
	2212: uint16(sym_permissible_text),
	2213: uint16(anon_sym_LT_BANG),
	2214: uint16(3),
	2215: uint16(3),
	2216: uint16(1),
	2217: uint16(sym_comment),
	2218: uint16(297),
	2219: uint16(1),
	2220: uint16(anon_sym_LT),
	2221: uint16(295),
	2222: uint16(4),
	2223: uint16(sym__html_interpolation_start),
	2224: uint16(sym__html_interpolation_end),
	2225: uint16(sym_permissible_text),
	2226: uint16(anon_sym_LT_BANG),
	2227: uint16(4),
	2228: uint16(3),
	2229: uint16(1),
	2230: uint16(sym_comment),
	2231: uint16(305),
	2232: uint16(1),
	2233: uint16(anon_sym_EQ),
	2234: uint16(307),
	2235: uint16(1),
	2236: uint16(sym_attribute_name),
	2237: uint16(303),
	2238: uint16(3),
	2239: uint16(anon_sym_GT),
	2240: uint16(anon_sym_SLASH_GT),
	2241: uint16(anon_sym_LBRACE),
	2242: uint16(3),
	2243: uint16(3),
	2244: uint16(1),
	2245: uint16(sym_comment),
	2246: uint16(245),
	2247: uint16(1),
	2248: uint16(anon_sym_LT),
	2249: uint16(243),
	2250: uint16(4),
	2251: uint16(sym__html_interpolation_start),
	2252: uint16(sym__html_interpolation_end),
	2253: uint16(sym_permissible_text),
	2254: uint16(anon_sym_LT_BANG),
	2255: uint16(3),
	2256: uint16(3),
	2257: uint16(1),
	2258: uint16(sym_comment),
	2259: uint16(237),
	2260: uint16(1),
	2261: uint16(anon_sym_LT),
	2262: uint16(235),
	2263: uint16(4),
	2264: uint16(sym__html_interpolation_start),
	2265: uint16(sym__html_interpolation_end),
	2266: uint16(sym_permissible_text),
	2267: uint16(anon_sym_LT_BANG),
	2268: uint16(3),
	2269: uint16(3),
	2270: uint16(1),
	2271: uint16(sym_comment),
	2272: uint16(241),
	2273: uint16(1),
	2274: uint16(anon_sym_LT),
	2275: uint16(239),
	2276: uint16(4),
	2277: uint16(sym__html_interpolation_start),
	2278: uint16(sym__html_interpolation_end),
	2279: uint16(sym_permissible_text),
	2280: uint16(anon_sym_LT_BANG),
	2281: uint16(3),
	2282: uint16(3),
	2283: uint16(1),
	2284: uint16(sym_comment),
	2285: uint16(311),
	2286: uint16(1),
	2287: uint16(sym_attribute_name),
	2288: uint16(309),
	2289: uint16(3),
	2290: uint16(anon_sym_GT),
	2291: uint16(anon_sym_SLASH_GT),
	2292: uint16(anon_sym_LBRACE),
	2293: uint16(3),
	2294: uint16(3),
	2295: uint16(1),
	2296: uint16(sym_comment),
	2297: uint16(315),
	2298: uint16(1),
	2299: uint16(sym_attribute_name),
	2300: uint16(313),
	2301: uint16(3),
	2302: uint16(anon_sym_GT),
	2303: uint16(anon_sym_SLASH_GT),
	2304: uint16(anon_sym_LBRACE),
	2305: uint16(3),
	2306: uint16(3),
	2307: uint16(1),
	2308: uint16(sym_comment),
	2309: uint16(319),
	2310: uint16(1),
	2311: uint16(sym_attribute_name),
	2312: uint16(317),
	2313: uint16(3),
	2314: uint16(anon_sym_GT),
	2315: uint16(anon_sym_SLASH_GT),
	2316: uint16(anon_sym_LBRACE),
	2317: uint16(3),
	2318: uint16(3),
	2319: uint16(1),
	2320: uint16(sym_comment),
	2321: uint16(323),
	2322: uint16(1),
	2323: uint16(sym_attribute_name),
	2324: uint16(321),
	2325: uint16(3),
	2326: uint16(anon_sym_GT),
	2327: uint16(anon_sym_SLASH_GT),
	2328: uint16(anon_sym_LBRACE),
	2329: uint16(5),
	2330: uint16(3),
	2331: uint16(1),
	2332: uint16(sym_comment),
	2333: uint16(325),
	2334: uint16(1),
	2335: uint16(sym__start_tag_name),
	2336: uint16(327),
	2337: uint16(1),
	2338: uint16(sym__script_start_tag_name),
	2339: uint16(329),
	2340: uint16(1),
	2341: uint16(sym__style_start_tag_name),
	2342: uint16(331),
	2343: uint16(1),
	2344: uint16(sym__fragment_tag_delim),
	2345: uint16(3),
	2346: uint16(3),
	2347: uint16(1),
	2348: uint16(sym_comment),
	2349: uint16(307),
	2350: uint16(1),
	2351: uint16(sym_attribute_name),
	2352: uint16(303),
	2353: uint16(3),
	2354: uint16(anon_sym_GT),
	2355: uint16(anon_sym_SLASH_GT),
	2356: uint16(anon_sym_LBRACE),
	2357: uint16(5),
	2358: uint16(3),
	2359: uint16(1),
	2360: uint16(sym_comment),
	2361: uint16(331),
	2362: uint16(1),
	2363: uint16(sym__fragment_tag_delim),
	2364: uint16(333),
	2365: uint16(1),
	2366: uint16(sym__start_tag_name),
	2367: uint16(335),
	2368: uint16(1),
	2369: uint16(sym__script_start_tag_name),
	2370: uint16(337),
	2371: uint16(1),
	2372: uint16(sym__style_start_tag_name),
	2373: uint16(5),
	2374: uint16(3),
	2375: uint16(1),
	2376: uint16(sym_comment),
	2377: uint16(331),
	2378: uint16(1),
	2379: uint16(sym__fragment_tag_delim),
	2380: uint16(339),
	2381: uint16(1),
	2382: uint16(sym__start_tag_name),
	2383: uint16(341),
	2384: uint16(1),
	2385: uint16(sym__script_start_tag_name),
	2386: uint16(343),
	2387: uint16(1),
	2388: uint16(sym__style_start_tag_name),
	2389: uint16(4),
	2390: uint16(3),
	2391: uint16(1),
	2392: uint16(sym_comment),
	2393: uint16(345),
	2394: uint16(1),
	2395: uint16(anon_sym_LT_SLASH),
	2396: uint16(347),
	2397: uint16(1),
	2398: uint16(sym_raw_text),
	2399: uint16(41),
	2400: uint16(1),
	2401: uint16(sym_end_tag),
	2402: uint16(4),
	2403: uint16(3),
	2404: uint16(1),
	2405: uint16(sym_comment),
	2406: uint16(345),
	2407: uint16(1),
	2408: uint16(anon_sym_LT_SLASH),
	2409: uint16(349),
	2410: uint16(1),
	2411: uint16(sym_raw_text),
	2412: uint16(79),
	2413: uint16(1),
	2414: uint16(sym_end_tag),
	2415: uint16(4),
	2416: uint16(3),
	2417: uint16(1),
	2418: uint16(sym_comment),
	2419: uint16(351),
	2420: uint16(1),
	2421: uint16(anon_sym_LT_SLASH),
	2422: uint16(353),
	2423: uint16(1),
	2424: uint16(sym_raw_text),
	2425: uint16(42),
	2426: uint16(1),
	2427: uint16(sym_end_tag),
	2428: uint16(4),
	2429: uint16(3),
	2430: uint16(1),
	2431: uint16(sym_comment),
	2432: uint16(351),
	2433: uint16(1),
	2434: uint16(anon_sym_LT_SLASH),
	2435: uint16(355),
	2436: uint16(1),
	2437: uint16(sym_raw_text),
	2438: uint16(47),
	2439: uint16(1),
	2440: uint16(sym_end_tag),
	2441: uint16(4),
	2442: uint16(3),
	2443: uint16(1),
	2444: uint16(sym_comment),
	2445: uint16(357),
	2446: uint16(1),
	2447: uint16(sym__end_tag_name),
	2448: uint16(359),
	2449: uint16(1),
	2450: uint16(sym_erroneous_end_tag_name),
	2451: uint16(361),
	2452: uint16(1),
	2453: uint16(sym__fragment_tag_delim),
	2454: uint16(4),
	2455: uint16(3),
	2456: uint16(1),
	2457: uint16(sym_comment),
	2458: uint16(363),
	2459: uint16(1),
	2460: uint16(anon_sym_LT_SLASH),
	2461: uint16(365),
	2462: uint16(1),
	2463: uint16(sym_raw_text),
	2464: uint16(101),
	2465: uint16(1),
	2466: uint16(sym_end_tag),
	2467: uint16(4),
	2468: uint16(3),
	2469: uint16(1),
	2470: uint16(sym_comment),
	2471: uint16(363),
	2472: uint16(1),
	2473: uint16(anon_sym_LT_SLASH),
	2474: uint16(367),
	2475: uint16(1),
	2476: uint16(sym_raw_text),
	2477: uint16(100),
	2478: uint16(1),
	2479: uint16(sym_end_tag),
	2480: uint16(4),
	2481: uint16(3),
	2482: uint16(1),
	2483: uint16(sym_comment),
	2484: uint16(359),
	2485: uint16(1),
	2486: uint16(sym_erroneous_end_tag_name),
	2487: uint16(369),
	2488: uint16(1),
	2489: uint16(sym__end_tag_name),
	2490: uint16(371),
	2491: uint16(1),
	2492: uint16(sym__fragment_tag_delim),
	2493: uint16(4),
	2494: uint16(3),
	2495: uint16(1),
	2496: uint16(sym_comment),
	2497: uint16(359),
	2498: uint16(1),
	2499: uint16(sym_erroneous_end_tag_name),
	2500: uint16(373),
	2501: uint16(1),
	2502: uint16(sym__end_tag_name),
	2503: uint16(375),
	2504: uint16(1),
	2505: uint16(sym__fragment_tag_delim),
	2506: uint16(3),
	2507: uint16(3),
	2508: uint16(1),
	2509: uint16(sym_comment),
	2510: uint16(377),
	2511: uint16(1),
	2512: uint16(anon_sym_SQUOTE),
	2513: uint16(379),
	2514: uint16(1),
	2515: uint16(aux_sym_quoted_attribute_value_token1),
	2516: uint16(3),
	2517: uint16(3),
	2518: uint16(1),
	2519: uint16(sym_comment),
	2520: uint16(377),
	2521: uint16(1),
	2522: uint16(anon_sym_DQUOTE),
	2523: uint16(381),
	2524: uint16(1),
	2525: uint16(aux_sym_quoted_attribute_value_token2),
	2526: uint16(2),
	2527: uint16(3),
	2528: uint16(1),
	2529: uint16(sym_comment),
	2530: uint16(383),
	2531: uint16(2),
	2532: uint16(sym_raw_text),
	2533: uint16(anon_sym_LT_SLASH),
	2534: uint16(3),
	2535: uint16(3),
	2536: uint16(1),
	2537: uint16(sym_comment),
	2538: uint16(351),
	2539: uint16(1),
	2540: uint16(anon_sym_LT_SLASH),
	2541: uint16(58),
	2542: uint16(1),
	2543: uint16(sym_end_tag),
	2544: uint16(2),
	2545: uint16(3),
	2546: uint16(1),
	2547: uint16(sym_comment),
	2548: uint16(385),
	2549: uint16(2),
	2550: uint16(sym_raw_text),
	2551: uint16(anon_sym_LT_SLASH),
	2552: uint16(3),
	2553: uint16(3),
	2554: uint16(1),
	2555: uint16(sym_comment),
	2556: uint16(363),
	2557: uint16(1),
	2558: uint16(anon_sym_LT_SLASH),
	2559: uint16(92),
	2560: uint16(1),
	2561: uint16(sym_end_tag),
	2562: uint16(3),
	2563: uint16(3),
	2564: uint16(1),
	2565: uint16(sym_comment),
	2566: uint16(373),
	2567: uint16(1),
	2568: uint16(sym__end_tag_name),
	2569: uint16(375),
	2570: uint16(1),
	2571: uint16(sym__fragment_tag_delim),
	2572: uint16(3),
	2573: uint16(3),
	2574: uint16(1),
	2575: uint16(sym_comment),
	2576: uint16(363),
	2577: uint16(1),
	2578: uint16(anon_sym_LT_SLASH),
	2579: uint16(93),
	2580: uint16(1),
	2581: uint16(sym_end_tag),
	2582: uint16(3),
	2583: uint16(3),
	2584: uint16(1),
	2585: uint16(sym_comment),
	2586: uint16(387),
	2587: uint16(1),
	2588: uint16(anon_sym_DASH_DASH_DASH2),
	2589: uint16(389),
	2590: uint16(1),
	2591: uint16(sym_frontmatter_js_block),
	2592: uint16(3),
	2593: uint16(3),
	2594: uint16(1),
	2595: uint16(sym_comment),
	2596: uint16(345),
	2597: uint16(1),
	2598: uint16(anon_sym_LT_SLASH),
	2599: uint16(69),
	2600: uint16(1),
	2601: uint16(sym_end_tag),
	2602: uint16(3),
	2603: uint16(3),
	2604: uint16(1),
	2605: uint16(sym_comment),
	2606: uint16(369),
	2607: uint16(1),
	2608: uint16(sym__end_tag_name),
	2609: uint16(371),
	2610: uint16(1),
	2611: uint16(sym__fragment_tag_delim),
	2612: uint16(3),
	2613: uint16(3),
	2614: uint16(1),
	2615: uint16(sym_comment),
	2616: uint16(357),
	2617: uint16(1),
	2618: uint16(sym__end_tag_name),
	2619: uint16(361),
	2620: uint16(1),
	2621: uint16(sym__fragment_tag_delim),
	2622: uint16(3),
	2623: uint16(3),
	2624: uint16(1),
	2625: uint16(sym_comment),
	2626: uint16(345),
	2627: uint16(1),
	2628: uint16(anon_sym_LT_SLASH),
	2629: uint16(70),
	2630: uint16(1),
	2631: uint16(sym_end_tag),
	2632: uint16(2),
	2633: uint16(3),
	2634: uint16(1),
	2635: uint16(sym_comment),
	2636: uint16(391),
	2637: uint16(2),
	2638: uint16(sym_raw_text),
	2639: uint16(anon_sym_LT_SLASH),
	2640: uint16(2),
	2641: uint16(3),
	2642: uint16(1),
	2643: uint16(sym_comment),
	2644: uint16(393),
	2645: uint16(2),
	2646: uint16(sym_raw_text),
	2647: uint16(anon_sym_LT_SLASH),
	2648: uint16(3),
	2649: uint16(3),
	2650: uint16(1),
	2651: uint16(sym_comment),
	2652: uint16(351),
	2653: uint16(1),
	2654: uint16(anon_sym_LT_SLASH),
	2655: uint16(59),
	2656: uint16(1),
	2657: uint16(sym_end_tag),
	2658: uint16(2),
	2659: uint16(3),
	2660: uint16(1),
	2661: uint16(sym_comment),
	2662: uint16(395),
	2663: uint16(1),
	2664: uint16(anon_sym_GT),
	2665: uint16(2),
	2666: uint16(3),
	2667: uint16(1),
	2668: uint16(sym_comment),
	2669: uint16(397),
	2670: uint16(1),
	2671: uint16(anon_sym_GT),
	2672: uint16(2),
	2673: uint16(3),
	2674: uint16(1),
	2675: uint16(sym_comment),
	2676: uint16(399),
	2677: uint16(1),
	2678: uint16(anon_sym_SQUOTE),
	2679: uint16(2),
	2680: uint16(3),
	2681: uint16(1),
	2682: uint16(sym_comment),
	2683: uint16(399),
	2684: uint16(1),
	2685: uint16(anon_sym_DQUOTE),
	2686: uint16(2),
	2687: uint16(3),
	2688: uint16(1),
	2689: uint16(sym_comment),
	2690: uint16(401),
	2691: uint16(1),
	2692: uint16(anon_sym_RBRACE),
	2693: uint16(2),
	2694: uint16(3),
	2695: uint16(1),
	2696: uint16(sym_comment),
	2697: uint16(403),
	2698: uint16(1),
	2699: uint16(anon_sym_GT),
	2700: uint16(2),
	2701: uint16(3),
	2702: uint16(1),
	2703: uint16(sym_comment),
	2704: uint16(405),
	2705: uint16(1),
	2707: uint16(2),
	2708: uint16(3),
	2709: uint16(1),
	2710: uint16(sym_comment),
	2711: uint16(407),
	2712: uint16(1),
	2713: uint16(sym__doctype),
	2714: uint16(2),
	2715: uint16(3),
	2716: uint16(1),
	2717: uint16(sym_comment),
	2718: uint16(409),
	2719: uint16(1),
	2720: uint16(anon_sym_GT),
	2721: uint16(2),
	2722: uint16(3),
	2723: uint16(1),
	2724: uint16(sym_comment),
	2725: uint16(411),
	2726: uint16(1),
	2727: uint16(aux_sym_doctype_token1),
	2728: uint16(2),
	2729: uint16(3),
	2730: uint16(1),
	2731: uint16(sym_comment),
	2732: uint16(413),
	2733: uint16(1),
	2734: uint16(anon_sym_GT),
	2735: uint16(2),
	2736: uint16(3),
	2737: uint16(1),
	2738: uint16(sym_comment),
	2739: uint16(415),
	2740: uint16(1),
	2741: uint16(anon_sym_GT),
	2742: uint16(2),
	2743: uint16(3),
	2744: uint16(1),
	2745: uint16(sym_comment),
	2746: uint16(417),
	2747: uint16(1),
	2748: uint16(sym_erroneous_end_tag_name),
	2749: uint16(2),
	2750: uint16(3),
	2751: uint16(1),
	2752: uint16(sym_comment),
	2753: uint16(419),
	2754: uint16(1),
	2755: uint16(anon_sym_GT),
	2756: uint16(2),
	2757: uint16(3),
	2758: uint16(1),
	2759: uint16(sym_comment),
	2760: uint16(421),
	2761: uint16(1),
	2762: uint16(anon_sym_GT),
	2763: uint16(2),
	2764: uint16(3),
	2765: uint16(1),
	2766: uint16(sym_comment),
	2767: uint16(423),
	2768: uint16(1),
	2769: uint16(anon_sym_DASH_DASH_DASH2),
	2770: uint16(2),
	2771: uint16(3),
	2772: uint16(1),
	2773: uint16(sym_comment),
	2774: uint16(359),
	2775: uint16(1),
	2776: uint16(sym_erroneous_end_tag_name),
	2777: uint16(2),
	2778: uint16(3),
	2779: uint16(1),
	2780: uint16(sym_comment),
	2781: uint16(425),
	2782: uint16(1),
	2783: uint16(aux_sym_doctype_token1),
	2784: uint16(2),
	2785: uint16(3),
	2786: uint16(1),
	2787: uint16(sym_comment),
	2788: uint16(427),
	2789: uint16(1),
	2790: uint16(sym_attribute_js_expr),
	2791: uint16(2),
	2792: uint16(3),
	2793: uint16(1),
	2794: uint16(sym_comment),
	2795: uint16(429),
	2796: uint16(1),
	2797: uint16(aux_sym_doctype_token1),
	2798: uint16(2),
	2799: uint16(3),
	2800: uint16(1),
	2801: uint16(sym_comment),
	2802: uint16(431),
	2803: uint16(1),
	2804: uint16(sym__doctype),
	2805: uint16(2),
	2806: uint16(3),
	2807: uint16(1),
	2808: uint16(sym_comment),
	2809: uint16(433),
	2810: uint16(1),
	2811: uint16(sym__doctype),
}

var ts_small_parse_table_map = [161]uint32_t{
	1:   uint32(53),
	2:   uint32(106),
	3:   uint32(159),
	4:   uint32(212),
	5:   uint32(265),
	6:   uint32(318),
	7:   uint32(368),
	8:   uint32(418),
	9:   uint32(468),
	10:  uint32(518),
	11:  uint32(568),
	12:  uint32(614),
	13:  uint32(660),
	14:  uint32(706),
	15:  uint32(752),
	16:  uint32(798),
	17:  uint32(844),
	18:  uint32(890),
	19:  uint32(913),
	20:  uint32(936),
	21:  uint32(959),
	22:  uint32(982),
	23:  uint32(1005),
	24:  uint32(1028),
	25:  uint32(1051),
	26:  uint32(1074),
	27:  uint32(1097),
	28:  uint32(1120),
	29:  uint32(1143),
	30:  uint32(1166),
	31:  uint32(1189),
	32:  uint32(1212),
	33:  uint32(1235),
	34:  uint32(1258),
	35:  uint32(1281),
	36:  uint32(1304),
	37:  uint32(1327),
	38:  uint32(1348),
	39:  uint32(1362),
	40:  uint32(1376),
	41:  uint32(1390),
	42:  uint32(1404),
	43:  uint32(1418),
	44:  uint32(1432),
	45:  uint32(1446),
	46:  uint32(1460),
	47:  uint32(1474),
	48:  uint32(1488),
	49:  uint32(1502),
	50:  uint32(1516),
	51:  uint32(1530),
	52:  uint32(1544),
	53:  uint32(1558),
	54:  uint32(1572),
	55:  uint32(1586),
	56:  uint32(1600),
	57:  uint32(1614),
	58:  uint32(1628),
	59:  uint32(1642),
	60:  uint32(1656),
	61:  uint32(1670),
	62:  uint32(1684),
	63:  uint32(1698),
	64:  uint32(1712),
	65:  uint32(1726),
	66:  uint32(1740),
	67:  uint32(1754),
	68:  uint32(1768),
	69:  uint32(1782),
	70:  uint32(1796),
	71:  uint32(1810),
	72:  uint32(1824),
	73:  uint32(1838),
	74:  uint32(1852),
	75:  uint32(1866),
	76:  uint32(1880),
	77:  uint32(1894),
	78:  uint32(1908),
	79:  uint32(1922),
	80:  uint32(1936),
	81:  uint32(1950),
	82:  uint32(1964),
	83:  uint32(1978),
	84:  uint32(1992),
	85:  uint32(2006),
	86:  uint32(2019),
	87:  uint32(2032),
	88:  uint32(2045),
	89:  uint32(2058),
	90:  uint32(2071),
	91:  uint32(2084),
	92:  uint32(2097),
	93:  uint32(2110),
	94:  uint32(2123),
	95:  uint32(2136),
	96:  uint32(2149),
	97:  uint32(2162),
	98:  uint32(2175),
	99:  uint32(2188),
	100: uint32(2201),
	101: uint32(2214),
	102: uint32(2227),
	103: uint32(2242),
	104: uint32(2255),
	105: uint32(2268),
	106: uint32(2281),
	107: uint32(2293),
	108: uint32(2305),
	109: uint32(2317),
	110: uint32(2329),
	111: uint32(2345),
	112: uint32(2357),
	113: uint32(2373),
	114: uint32(2389),
	115: uint32(2402),
	116: uint32(2415),
	117: uint32(2428),
	118: uint32(2441),
	119: uint32(2454),
	120: uint32(2467),
	121: uint32(2480),
	122: uint32(2493),
	123: uint32(2506),
	124: uint32(2516),
	125: uint32(2526),
	126: uint32(2534),
	127: uint32(2544),
	128: uint32(2552),
	129: uint32(2562),
	130: uint32(2572),
	131: uint32(2582),
	132: uint32(2592),
	133: uint32(2602),
	134: uint32(2612),
	135: uint32(2622),
	136: uint32(2632),
	137: uint32(2640),
	138: uint32(2648),
	139: uint32(2658),
	140: uint32(2665),
	141: uint32(2672),
	142: uint32(2679),
	143: uint32(2686),
	144: uint32(2693),
	145: uint32(2700),
	146: uint32(2707),
	147: uint32(2714),
	148: uint32(2721),
	149: uint32(2728),
	150: uint32(2735),
	151: uint32(2742),
	152: uint32(2749),
	153: uint32(2756),
	154: uint32(2763),
	155: uint32(2770),
	156: uint32(2777),
	157: uint32(2784),
	158: uint32(2791),
	159: uint32(2798),
	160: uint32(2805),
}

var ts_parse_actions = [435]TSParseActionEntry{
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
		Fsymbol:     uint16(sym_document),
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
		Fstate: uint16(libc.Int32FromInt32(148)),
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
		Fstate: uint16(libc.Int32FromInt32(112)),
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
		Fstate: uint16(libc.Int32FromInt32(153)),
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
		Fstate: uint16(libc.Int32FromInt32(12)),
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
		Fstate: uint16(libc.Int32FromInt32(133)),
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
		Fstate: uint16(libc.Int32FromInt32(13)),
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
		Fstate: uint16(libc.Int32FromInt32(162)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(115)),
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
		Fstate: uint16(libc.Int32FromInt32(124)),
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
		Fstate: uint16(libc.Int32FromInt32(10)),
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
		Fstate: uint16(libc.Int32FromInt32(91)),
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
		Fstate: uint16(libc.Int32FromInt32(17)),
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
		Fstate: uint16(libc.Int32FromInt32(120)),
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
		Fstate: uint16(libc.Int32FromInt32(71)),
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
		Fstate: uint16(libc.Int32FromInt32(123)),
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
		Fstate: uint16(libc.Int32FromInt32(7)),
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
		Fstate: uint16(libc.Int32FromInt32(46)),
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
		Fstate: uint16(libc.Int32FromInt32(3)),
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
		Fstate: uint16(libc.Int32FromInt32(80)),
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
		Fstate: uint16(libc.Int32FromInt32(2)),
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
		Fstate: uint16(libc.Int32FromInt32(102)),
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
		Fstate: uint16(libc.Int32FromInt32(57)),
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
		Fsymbol:      uint16(aux_sym_document_repeat1),
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
		Fsymbol:      uint16(aux_sym_document_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(148)),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_document_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(112)),
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
		Fsymbol:      uint16(aux_sym_document_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(153)),
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
		Fsymbol:      uint16(aux_sym_document_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(8)),
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
		Fsymbol:      uint16(aux_sym_document_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(13)),
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
		Fcount:    uint8(1),
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
		Fsymbol:      uint16(sym_document),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	71: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	72: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	73: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_document_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(162)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_document_repeat1),
	})))),
	77: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(libc.Int32FromInt32(115)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	78: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	79: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_document_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(157)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	82: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_document_repeat1),
	})))),
	83: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_document_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(17)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_document),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(9)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(161)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(114)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(15)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	98: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(18)),
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
		Fsymbol:      uint16(aux_sym_html_interpolation_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(161)),
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
		Fsymbol:      uint16(aux_sym_html_interpolation_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(114)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	108: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_html_interpolation_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(15)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_html_interpolation_repeat1),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_html_interpolation_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(14)),
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
		Fstate: uint16(libc.Int32FromInt32(103)),
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
		Fstate: uint16(libc.Int32FromInt32(19)),
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
		Fstate: uint16(libc.Int32FromInt32(78)),
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
		Fstate: uint16(libc.Int32FromInt32(14)),
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
		Fstate: uint16(libc.Int32FromInt32(81)),
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
		Fstate: uint16(libc.Int32FromInt32(16)),
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
		Fstate: uint16(libc.Int32FromInt32(43)),
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
		Fstate: uint16(libc.Int32FromInt32(99)),
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
		Fstate: uint16(libc.Int32FromInt32(45)),
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
		Fstate: uint16(libc.Int32FromInt32(77)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(104)),
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
		Fstate: uint16(libc.Int32FromInt32(159)),
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
		Fstate: uint16(libc.Int32FromInt32(129)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(52)),
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
		Fstate: uint16(libc.Int32FromInt32(62)),
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
		Fstate: uint16(libc.Int32FromInt32(95)),
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
		Fstate: uint16(libc.Int32FromInt32(127)),
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
		Fstate: uint16(libc.Int32FromInt32(50)),
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
		Fstate: uint16(libc.Int32FromInt32(105)),
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
		Fstate: uint16(libc.Int32FromInt32(106)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(139)),
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
		Fstate: uint16(libc.Int32FromInt32(49)),
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
		Fstate: uint16(libc.Int32FromInt32(40)),
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
		Fstate: uint16(libc.Int32FromInt32(138)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(97)),
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
		Fstate: uint16(libc.Int32FromInt32(76)),
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
		Fstate: uint16(libc.Int32FromInt32(67)),
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
		Fstate: uint16(libc.Int32FromInt32(65)),
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
		Fstate: uint16(libc.Int32FromInt32(73)),
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
		Fstate: uint16(libc.Int32FromInt32(55)),
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
		Fstate: uint16(libc.Int32FromInt32(111)),
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
		Fstate: uint16(libc.Int32FromInt32(125)),
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
		Fstate: uint16(libc.Int32FromInt32(126)),
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
		Fstate: uint16(libc.Int32FromInt32(111)),
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
		Fstate: uint16(libc.Int32FromInt32(51)),
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
		Fstate: uint16(libc.Int32FromInt32(63)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_start_tag_repeat1),
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
		Fcount: uint8(2),
	}})),
	194: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_start_tag_repeat1),
	})))),
	195: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(libc.Int32FromInt32(104)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_start_tag_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(159)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_self_closing_tag),
	})))),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_self_closing_tag),
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
		Fsymbol:      uint16(sym_script_element),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_script_element),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_style_element),
	})))),
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
		Fcount: uint8(1),
	}})),
	210: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_style_element),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_html_interpolation),
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
		Fcount: uint8(1),
	}})),
	214: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_html_interpolation),
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
		Fcount:    uint8(1),
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
		Fsymbol:      uint16(sym_frontmatter),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_frontmatter),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_start_tag),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_start_tag),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_element),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_element),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_end_tag),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_end_tag),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_self_closing_style_tag),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_self_closing_style_tag),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_self_closing_script_tag),
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
		Fsymbol:      uint16(sym_self_closing_script_tag),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_self_closing_script_tag),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_self_closing_script_tag),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_self_closing_style_tag),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_self_closing_style_tag),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_erroneous_end_tag),
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
		Fsymbol:      uint16(sym_erroneous_end_tag),
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
		Fsymbol:      uint16(sym_frontmatter),
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
		Fsymbol:      uint16(sym_frontmatter),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_self_closing_tag),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_self_closing_tag),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_end_tag),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_end_tag),
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
		Fsymbol:      uint16(sym_element),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_element),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_script_element),
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
		Fsymbol:      uint16(sym_script_element),
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
		Fsymbol:      uint16(sym_style_element),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_style_element),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_doctype),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_doctype),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_start_tag),
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
		Fcount: uint8(1),
	}})),
	282: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_start_tag),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	284: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_script_element),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_script_element),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	292: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_style_element),
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
		Fcount: uint8(1),
	}})),
	294: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_style_element),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_html_interpolation),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_html_interpolation),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_start_tag),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_start_tag),
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
		Fchild_count: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(36)),
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
		Fchild_count: uint8(1),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_quoted_attribute_value),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_quoted_attribute_value),
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
		Fsymbol:      uint16(sym_quoted_attribute_value),
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
		Fcount: uint8(1),
	}})),
	316: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_quoted_attribute_value),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_attribute_interpolation),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_attribute_interpolation),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	322: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_attribute),
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
		Fcount: uint8(1),
	}})),
	324: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_attribute),
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
		Fstate: uint16(libc.Int32FromInt32(28)),
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
		Fstate: uint16(libc.Int32FromInt32(23)),
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
		Fstate: uint16(libc.Int32FromInt32(21)),
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
		Fstate: uint16(libc.Int32FromInt32(84)),
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
		Fstate: uint16(libc.Int32FromInt32(26)),
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
		Fstate: uint16(libc.Int32FromInt32(25)),
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
		Fstate: uint16(libc.Int32FromInt32(24)),
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
		Fstate: uint16(libc.Int32FromInt32(20)),
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
		Fstate: uint16(libc.Int32FromInt32(31)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(34)),
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
		Fstate: uint16(libc.Int32FromInt32(136)),
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
		Fstate: uint16(libc.Int32FromInt32(137)),
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
		Fstate: uint16(libc.Int32FromInt32(134)),
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
		Fstate: uint16(libc.Int32FromInt32(135)),
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
		Fstate: uint16(libc.Int32FromInt32(140)),
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
		Fstate: uint16(libc.Int32FromInt32(128)),
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
		Fstate: uint16(libc.Int32FromInt32(155)),
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
		Fstate: uint16(libc.Int32FromInt32(152)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(131)),
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
		Fstate: uint16(libc.Int32FromInt32(130)),
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
		Fstate: uint16(libc.Int32FromInt32(132)),
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
		Fstate: uint16(libc.Int32FromInt32(146)),
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
		Fstate: uint16(libc.Int32FromInt32(56)),
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
		Fstate: uint16(libc.Int32FromInt32(149)),
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
		Fstate: uint16(libc.Int32FromInt32(90)),
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
		Fstate: uint16(libc.Int32FromInt32(109)),
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
		Fstate: uint16(libc.Int32FromInt32(143)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(144)),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_script_start_tag),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	386: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_style_start_tag),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(44)),
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
		Fstate: uint16(libc.Int32FromInt32(156)),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_script_start_tag),
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
		Fsymbol:      uint16(sym_style_start_tag),
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
		Fstate: uint16(libc.Int32FromInt32(60)),
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
		Fstate: uint16(libc.Int32FromInt32(53)),
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
		Fstate: uint16(libc.Int32FromInt32(108)),
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
		Fstate: uint16(libc.Int32FromInt32(110)),
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
		Fstate: uint16(libc.Int32FromInt32(68)),
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
	406: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f uint8_t
		_ [7]byte
	}{f: uint8(TSParseActionTypeAccept)})),
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
		Fstate: uint16(libc.Int32FromInt32(150)),
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
		Fstate: uint16(libc.Int32FromInt32(98)),
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
		Fstate: uint16(libc.Int32FromInt32(141)),
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
		Fstate: uint16(libc.Int32FromInt32(61)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(142)),
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
		Fstate: uint16(libc.Int32FromInt32(94)),
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
		Fstate: uint16(libc.Int32FromInt32(48)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
}

type ts_external_scanner_symbol_identifiers = int32

const ts_external_token__start_tag_name = 0
const ts_external_token__script_start_tag_name = 1
const ts_external_token__style_start_tag_name = 2
const ts_external_token__end_tag_name = 3
const ts_external_token_erroneous_end_tag_name = 4
const ts_external_token_SLASH_GT = 5
const ts_external_token__implicit_end_tag = 6
const ts_external_token_raw_text = 7
const ts_external_token_comment = 8
const ts_external_token__html_interpolation_start = 9
const ts_external_token__html_interpolation_end = 10
const ts_external_token_frontmatter_js_block = 11
const ts_external_token_attribute_js_expr = 12
const ts_external_token_attribute_backtick_string = 13
const ts_external_token_permissible_text = 14
const ts_external_token__fragment_tag_delim = 15

var ts_external_scanner_symbol_map = [16]TSSymbol{
	0:  uint16(sym__start_tag_name),
	1:  uint16(sym__script_start_tag_name),
	2:  uint16(sym__style_start_tag_name),
	3:  uint16(sym__end_tag_name),
	4:  uint16(sym_erroneous_end_tag_name),
	5:  uint16(anon_sym_SLASH_GT),
	6:  uint16(sym__implicit_end_tag),
	7:  uint16(sym_raw_text),
	8:  uint16(sym_comment),
	9:  uint16(sym__html_interpolation_start),
	10: uint16(sym__html_interpolation_end),
	11: uint16(sym_frontmatter_js_block),
	12: uint16(sym_attribute_js_expr),
	13: uint16(sym_attribute_backtick_string),
	14: uint16(sym_permissible_text),
	15: uint16(sym__fragment_tag_delim),
}

var ts_external_scanner_states = [15][16]uint8{
	1: {
		0:  libc.BoolUint8(true1 != 0),
		1:  libc.BoolUint8(true1 != 0),
		2:  libc.BoolUint8(true1 != 0),
		3:  libc.BoolUint8(true1 != 0),
		4:  libc.BoolUint8(true1 != 0),
		5:  libc.BoolUint8(true1 != 0),
		6:  libc.BoolUint8(true1 != 0),
		7:  libc.BoolUint8(true1 != 0),
		8:  libc.BoolUint8(true1 != 0),
		9:  libc.BoolUint8(true1 != 0),
		10: libc.BoolUint8(true1 != 0),
		11: libc.BoolUint8(true1 != 0),
		12: libc.BoolUint8(true1 != 0),
		13: libc.BoolUint8(true1 != 0),
		14: libc.BoolUint8(true1 != 0),
		15: libc.BoolUint8(true1 != 0),
	},
	2: {
		8: libc.BoolUint8(true1 != 0),
		9: libc.BoolUint8(true1 != 0),
	},
	3: {
		6: libc.BoolUint8(true1 != 0),
		8: libc.BoolUint8(true1 != 0),
		9: libc.BoolUint8(true1 != 0),
	},
	4: {
		8:  libc.BoolUint8(true1 != 0),
		9:  libc.BoolUint8(true1 != 0),
		10: libc.BoolUint8(true1 != 0),
		14: libc.BoolUint8(true1 != 0),
	},
	5: {
		5: libc.BoolUint8(true1 != 0),
		8: libc.BoolUint8(true1 != 0),
	},
	6: {
		8:  libc.BoolUint8(true1 != 0),
		13: libc.BoolUint8(true1 != 0),
	},
	7: {
		0:  libc.BoolUint8(true1 != 0),
		1:  libc.BoolUint8(true1 != 0),
		2:  libc.BoolUint8(true1 != 0),
		8:  libc.BoolUint8(true1 != 0),
		15: libc.BoolUint8(true1 != 0),
	},
	8: {
		7: libc.BoolUint8(true1 != 0),
		8: libc.BoolUint8(true1 != 0),
	},
	9: {
		3:  libc.BoolUint8(true1 != 0),
		4:  libc.BoolUint8(true1 != 0),
		8:  libc.BoolUint8(true1 != 0),
		15: libc.BoolUint8(true1 != 0),
	},
	10: {
		8: libc.BoolUint8(true1 != 0),
	},
	11: {
		3:  libc.BoolUint8(true1 != 0),
		8:  libc.BoolUint8(true1 != 0),
		15: libc.BoolUint8(true1 != 0),
	},
	12: {
		8:  libc.BoolUint8(true1 != 0),
		11: libc.BoolUint8(true1 != 0),
	},
	13: {
		4: libc.BoolUint8(true1 != 0),
		8: libc.BoolUint8(true1 != 0),
	},
	14: {
		8:  libc.BoolUint8(true1 != 0),
		12: libc.BoolUint8(true1 != 0),
	},
}

func tree_sitter_astro(tls *libc.TLS) (r uintptr) {
	return uintptr(unsafe.Pointer(&language))
}

var language = TSLanguage{
	Fversion:                   uint32(LANGUAGE_VERSION),
	Fsymbol_count:              uint32(SYMBOL_COUNT),
	Ftoken_count:               uint32(TOKEN_COUNT),
	Fexternal_token_count:      uint32(EXTERNAL_TOKEN_COUNT),
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
	*(*uintptr)(unsafe.Add(p, 184)) = __ccgo_fp(tree_sitter_astro_external_scanner_create)
	*(*uintptr)(unsafe.Add(p, 192)) = __ccgo_fp(tree_sitter_astro_external_scanner_destroy)
	*(*uintptr)(unsafe.Add(p, 200)) = __ccgo_fp(tree_sitter_astro_external_scanner_scan)
	*(*uintptr)(unsafe.Add(p, 208)) = __ccgo_fp(tree_sitter_astro_external_scanner_serialize)
	*(*uintptr)(unsafe.Add(p, 216)) = __ccgo_fp(tree_sitter_astro_external_scanner_deserialize)
}

func __ccgo_fp(f interface{}) uintptr {
	type iface [2]uintptr
	return (*iface)(unsafe.Pointer(&f))[1]
}

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = "\n---\x00(uint32_t)((&scanner->tags)->size - 1) < (&scanner->tags)->size\x00combined.c\x00</script\x00</style\x00end\x00<!\x00doctype_token1\x00>\x00doctype\x00<\x00/>\x00</\x00=\x00attribute_name\x00attribute_value\x00entity\x00'\x00\"\x00text\x00---\x00{\x00}\x00tag_name\x00erroneous_end_tag_name\x00_implicit_end_tag\x00raw_text\x00comment\x00frontmatter_js_block\x00attribute_js_expr\x00attribute_backtick_string\x00permissible_text\x00document\x00_node\x00element\x00script_element\x00style_element\x00start_tag\x00self_closing_tag\x00end_tag\x00erroneous_end_tag\x00attribute\x00quoted_attribute_value\x00_node_with_permissible_text\x00frontmatter\x00attribute_interpolation\x00html_interpolation\x00document_repeat1\x00start_tag_repeat1\x00html_interpolation_repeat1\x00"
