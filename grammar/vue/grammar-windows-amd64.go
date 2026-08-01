// Code generated for windows/amd64 by 'ccgo -ignore-unsupported-alignment -ignore-unsupported-atomic-sizes -ignore-vector-functions --winapi-no-errno -ignore-link-errors -std=gnu11 -O0 -D__extension__= -D__forceinline=static inline -D__attribute__(...)= -D__declspec(x)= -D__cdecl= -D__stdcall= -D__fastcall= -D__thiscall= -D_cdecl= -D__restrict= -D__restrict__= -D__MINGW_EXTENSION= -D_X86INTRIN_H_INCLUDED -D_X86GPRINTRIN_H_INCLUDED -D_IMMINTRIN_H_INCLUDED -D_MMINTRIN_H_INCLUDED -D_XMMINTRIN_H_INCLUDED -D_EMMINTRIN_H_INCLUDED -D_PMMINTRIN_H_INCLUDED -D_MM3DNOW_H_INCLUDED -Dfunc=func_token -Dinterface=interface_token -Dselect=select_token -Dchan=chan_token -Dgo=go_token -Dmap=map_token -Dpackage=package_token -Dtype=type_token -Dvar=var_token -Dimport=import_token -Ddefer=defer_token -Dfallthrough=fallthrough_token -Drange=range_token -I D:\a\ccgo-tree-sitter\ccgo-tree-sitter\third-party\tree-sitter-vue\src -I D:\a\ccgo-tree-sitter\ccgo-tree-sitter\third-party\tree-sitter-vue -I D:\a\ccgo-tree-sitter\ccgo-tree-sitter\third-party -I C:\Users\runneradmin\.cache\workspaced\sources\github\7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2\lib\include -I C:\Users\runneradmin\.cache\workspaced\sources\github\7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2\lib\src combined.c -o grammar.go', DO NOT EDIT.

//go:build windows && amd64

package grammar_vue

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
const EXTERNAL_TOKEN_COUNT = 12
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
const LANGUAGE_VERSION = 15
const LARGE_STATE_COUNT = 2
const L_tmpnam = 260
const L_tmpnam_s = "L_tmpnam"
const MAX_ALIAS_SEQUENCE_LENGTH = 6
const MAX_RESERVED_WORD_SET_SIZE = 0
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
const SIZE_MAX = "_UI64_MAX"
const SSIZE_MAX = "_I64_MAX"
const STATE_COUNT = 170
const STDERR_FILENO = 2
const STDIN_FILENO = 0
const STDOUT_FILENO = 1
const STRUNCATE = 80
const SUPERTYPE_COUNT = 0
const SYMBOL_COUNT = 67
const SYS_OPEN = "_SYS_OPEN"
const TMP_MAX = 2147483647
const TMP_MAX_S = "TMP_MAX"
const TOKEN_COUNT = 41
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

type Array = struct {
	Fcontents uintptr
	Fsize     uint32_t
	Fcapacity uint32_t
}

type TagType = int32

const AREA = 0
const BASE = 1
const BASEFONT = 2
const BGSOUND = 3
const BR = 4
const COL = 5
const COMMAND = 6
const EMBED = 7
const FRAME = 8
const HR = 9
const IMAGE = 10
const IMG = 11
const INPUT = 12
const ISINDEX = 13
const KEYGEN = 14
const LINK = 15
const MENUITEM = 16
const META = 17
const NEXTID = 18
const PARAM = 19
const SOURCE = 20
const TRACK = 21
const WBR = 22
const END_OF_VOID_TAGS = 23
const A = 24
const ABBR = 25
const ADDRESS = 26
const ARTICLE = 27
const ASIDE = 28
const AUDIO = 29
const B = 30
const BDI = 31
const BDO = 32
const BLOCKQUOTE = 33
const BODY = 34
const BUTTON = 35
const CANVAS = 36
const CAPTION = 37
const CITE = 38
const CODE = 39
const COLGROUP = 40
const DATA = 41
const DATALIST = 42
const DD = 43
const DEL = 44
const DETAILS = 45
const DFN = 46
const DIALOG = 47
const DIV = 48
const DL = 49
const DT = 50
const EM = 51
const FIELDSET = 52
const FIGCAPTION = 53
const FIGURE = 54
const FOOTER = 55
const FORM = 56
const H1 = 57
const H2 = 58
const H3 = 59
const H4 = 60
const H5 = 61
const H6 = 62
const HEAD = 63
const HEADER = 64
const HGROUP = 65
const HTML = 66
const I = 67
const IFRAME = 68
const INS = 69
const KBD = 70
const LABEL = 71
const LEGEND = 72
const LI = 73
const MAIN = 74
const MAP = 75
const MARK = 76
const MATH = 77
const MENU = 78
const METER = 79
const NAV = 80
const NOSCRIPT = 81
const OBJECT = 82
const OL = 83
const OPTGROUP = 84
const OPTION = 85
const OUTPUT = 86
const P = 87
const PICTURE = 88
const PRE = 89
const PROGRESS = 90
const Q = 91
const RB = 92
const RP = 93
const RT = 94
const RTC = 95
const RUBY = 96
const S = 97
const SAMP = 98
const SCRIPT = 99
const SECTION = 100
const SELECT = 101
const SLOT = 102
const SMALL = 103
const SPAN = 104
const STRONG = 105
const STYLE = 106
const SUB = 107
const SUMMARY = 108
const SUP = 109
const SVG = 110
const TABLE = 111
const TBODY = 112
const TD = 113
const TEMPLATE = 114
const TEXTAREA = 115
const TFOOT = 116
const TH = 117
const THEAD = 118
const TIME = 119
const TITLE = 120
const TR = 121
const U = 122
const UL = 123
const VAR = 124
const VIDEO = 125
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
		Ftag_name: [16]int8{'A', 'R', 'E', 'A'},
	},
	1: {
		Ftag_name: [16]int8{'B', 'A', 'S', 'E'},
		Ftag_type: int32(BASE),
	},
	2: {
		Ftag_name: [16]int8{'B', 'A', 'S', 'E', 'F', 'O', 'N', 'T'},
		Ftag_type: int32(BASEFONT),
	},
	3: {
		Ftag_name: [16]int8{'B', 'G', 'S', 'O', 'U', 'N', 'D'},
		Ftag_type: int32(BGSOUND),
	},
	4: {
		Ftag_name: [16]int8{'B', 'R'},
		Ftag_type: int32(BR),
	},
	5: {
		Ftag_name: [16]int8{'C', 'O', 'L'},
		Ftag_type: int32(COL),
	},
	6: {
		Ftag_name: [16]int8{'C', 'O', 'M', 'M', 'A', 'N', 'D'},
		Ftag_type: int32(COMMAND),
	},
	7: {
		Ftag_name: [16]int8{'E', 'M', 'B', 'E', 'D'},
		Ftag_type: int32(EMBED),
	},
	8: {
		Ftag_name: [16]int8{'F', 'R', 'A', 'M', 'E'},
		Ftag_type: int32(FRAME),
	},
	9: {
		Ftag_name: [16]int8{'H', 'R'},
		Ftag_type: int32(HR),
	},
	10: {
		Ftag_name: [16]int8{'I', 'M', 'A', 'G', 'E'},
		Ftag_type: int32(IMAGE),
	},
	11: {
		Ftag_name: [16]int8{'I', 'M', 'G'},
		Ftag_type: int32(IMG),
	},
	12: {
		Ftag_name: [16]int8{'I', 'N', 'P', 'U', 'T'},
		Ftag_type: int32(INPUT),
	},
	13: {
		Ftag_name: [16]int8{'I', 'S', 'I', 'N', 'D', 'E', 'X'},
		Ftag_type: int32(ISINDEX),
	},
	14: {
		Ftag_name: [16]int8{'K', 'E', 'Y', 'G', 'E', 'N'},
		Ftag_type: int32(KEYGEN),
	},
	15: {
		Ftag_name: [16]int8{'L', 'I', 'N', 'K'},
		Ftag_type: int32(LINK),
	},
	16: {
		Ftag_name: [16]int8{'M', 'E', 'N', 'U', 'I', 'T', 'E', 'M'},
		Ftag_type: int32(MENUITEM),
	},
	17: {
		Ftag_name: [16]int8{'M', 'E', 'T', 'A'},
		Ftag_type: int32(META),
	},
	18: {
		Ftag_name: [16]int8{'N', 'E', 'X', 'T', 'I', 'D'},
		Ftag_type: int32(NEXTID),
	},
	19: {
		Ftag_name: [16]int8{'P', 'A', 'R', 'A', 'M'},
		Ftag_type: int32(PARAM),
	},
	20: {
		Ftag_name: [16]int8{'S', 'O', 'U', 'R', 'C', 'E'},
		Ftag_type: int32(SOURCE),
	},
	21: {
		Ftag_name: [16]int8{'T', 'R', 'A', 'C', 'K'},
		Ftag_type: int32(TRACK),
	},
	22: {
		Ftag_name: [16]int8{'W', 'B', 'R'},
		Ftag_type: int32(WBR),
	},
	23: {
		Ftag_name: [16]int8{'A'},
		Ftag_type: int32(A),
	},
	24: {
		Ftag_name: [16]int8{'A', 'B', 'B', 'R'},
		Ftag_type: int32(ABBR),
	},
	25: {
		Ftag_name: [16]int8{'A', 'D', 'D', 'R', 'E', 'S', 'S'},
		Ftag_type: int32(ADDRESS),
	},
	26: {
		Ftag_name: [16]int8{'A', 'R', 'T', 'I', 'C', 'L', 'E'},
		Ftag_type: int32(ARTICLE),
	},
	27: {
		Ftag_name: [16]int8{'A', 'S', 'I', 'D', 'E'},
		Ftag_type: int32(ASIDE),
	},
	28: {
		Ftag_name: [16]int8{'A', 'U', 'D', 'I', 'O'},
		Ftag_type: int32(AUDIO),
	},
	29: {
		Ftag_name: [16]int8{'B'},
		Ftag_type: int32(B),
	},
	30: {
		Ftag_name: [16]int8{'B', 'D', 'I'},
		Ftag_type: int32(BDI),
	},
	31: {
		Ftag_name: [16]int8{'B', 'D', 'O'},
		Ftag_type: int32(BDO),
	},
	32: {
		Ftag_name: [16]int8{'B', 'L', 'O', 'C', 'K', 'Q', 'U', 'O', 'T', 'E'},
		Ftag_type: int32(BLOCKQUOTE),
	},
	33: {
		Ftag_name: [16]int8{'B', 'O', 'D', 'Y'},
		Ftag_type: int32(BODY),
	},
	34: {
		Ftag_name: [16]int8{'B', 'U', 'T', 'T', 'O', 'N'},
		Ftag_type: int32(BUTTON),
	},
	35: {
		Ftag_name: [16]int8{'C', 'A', 'N', 'V', 'A', 'S'},
		Ftag_type: int32(CANVAS),
	},
	36: {
		Ftag_name: [16]int8{'C', 'A', 'P', 'T', 'I', 'O', 'N'},
		Ftag_type: int32(CAPTION),
	},
	37: {
		Ftag_name: [16]int8{'C', 'I', 'T', 'E'},
		Ftag_type: int32(CITE),
	},
	38: {
		Ftag_name: [16]int8{'C', 'O', 'D', 'E'},
		Ftag_type: int32(CODE),
	},
	39: {
		Ftag_name: [16]int8{'C', 'O', 'L', 'G', 'R', 'O', 'U', 'P'},
		Ftag_type: int32(COLGROUP),
	},
	40: {
		Ftag_name: [16]int8{'D', 'A', 'T', 'A'},
		Ftag_type: int32(DATA),
	},
	41: {
		Ftag_name: [16]int8{'D', 'A', 'T', 'A', 'L', 'I', 'S', 'T'},
		Ftag_type: int32(DATALIST),
	},
	42: {
		Ftag_name: [16]int8{'D', 'D'},
		Ftag_type: int32(DD),
	},
	43: {
		Ftag_name: [16]int8{'D', 'E', 'L'},
		Ftag_type: int32(DEL),
	},
	44: {
		Ftag_name: [16]int8{'D', 'E', 'T', 'A', 'I', 'L', 'S'},
		Ftag_type: int32(DETAILS),
	},
	45: {
		Ftag_name: [16]int8{'D', 'F', 'N'},
		Ftag_type: int32(DFN),
	},
	46: {
		Ftag_name: [16]int8{'D', 'I', 'A', 'L', 'O', 'G'},
		Ftag_type: int32(DIALOG),
	},
	47: {
		Ftag_name: [16]int8{'D', 'I', 'V'},
		Ftag_type: int32(DIV),
	},
	48: {
		Ftag_name: [16]int8{'D', 'L'},
		Ftag_type: int32(DL),
	},
	49: {
		Ftag_name: [16]int8{'D', 'T'},
		Ftag_type: int32(DT),
	},
	50: {
		Ftag_name: [16]int8{'E', 'M'},
		Ftag_type: int32(EM),
	},
	51: {
		Ftag_name: [16]int8{'F', 'I', 'E', 'L', 'D', 'S', 'E', 'T'},
		Ftag_type: int32(FIELDSET),
	},
	52: {
		Ftag_name: [16]int8{'F', 'I', 'G', 'C', 'A', 'P', 'T', 'I', 'O', 'N'},
		Ftag_type: int32(FIGCAPTION),
	},
	53: {
		Ftag_name: [16]int8{'F', 'I', 'G', 'U', 'R', 'E'},
		Ftag_type: int32(FIGURE),
	},
	54: {
		Ftag_name: [16]int8{'F', 'O', 'O', 'T', 'E', 'R'},
		Ftag_type: int32(FOOTER),
	},
	55: {
		Ftag_name: [16]int8{'F', 'O', 'R', 'M'},
		Ftag_type: int32(FORM),
	},
	56: {
		Ftag_name: [16]int8{'H', '1'},
		Ftag_type: int32(H1),
	},
	57: {
		Ftag_name: [16]int8{'H', '2'},
		Ftag_type: int32(H2),
	},
	58: {
		Ftag_name: [16]int8{'H', '3'},
		Ftag_type: int32(H3),
	},
	59: {
		Ftag_name: [16]int8{'H', '4'},
		Ftag_type: int32(H4),
	},
	60: {
		Ftag_name: [16]int8{'H', '5'},
		Ftag_type: int32(H5),
	},
	61: {
		Ftag_name: [16]int8{'H', '6'},
		Ftag_type: int32(H6),
	},
	62: {
		Ftag_name: [16]int8{'H', 'E', 'A', 'D'},
		Ftag_type: int32(HEAD),
	},
	63: {
		Ftag_name: [16]int8{'H', 'E', 'A', 'D', 'E', 'R'},
		Ftag_type: int32(HEADER),
	},
	64: {
		Ftag_name: [16]int8{'H', 'G', 'R', 'O', 'U', 'P'},
		Ftag_type: int32(HGROUP),
	},
	65: {
		Ftag_name: [16]int8{'H', 'T', 'M', 'L'},
		Ftag_type: int32(HTML),
	},
	66: {
		Ftag_name: [16]int8{'I'},
		Ftag_type: int32(I),
	},
	67: {
		Ftag_name: [16]int8{'I', 'F', 'R', 'A', 'M', 'E'},
		Ftag_type: int32(IFRAME),
	},
	68: {
		Ftag_name: [16]int8{'I', 'N', 'S'},
		Ftag_type: int32(INS),
	},
	69: {
		Ftag_name: [16]int8{'K', 'B', 'D'},
		Ftag_type: int32(KBD),
	},
	70: {
		Ftag_name: [16]int8{'L', 'A', 'B', 'E', 'L'},
		Ftag_type: int32(LABEL),
	},
	71: {
		Ftag_name: [16]int8{'L', 'E', 'G', 'E', 'N', 'D'},
		Ftag_type: int32(LEGEND),
	},
	72: {
		Ftag_name: [16]int8{'L', 'I'},
		Ftag_type: int32(LI),
	},
	73: {
		Ftag_name: [16]int8{'M', 'A', 'I', 'N'},
		Ftag_type: int32(MAIN),
	},
	74: {
		Ftag_name: [16]int8{'M', 'A', 'P'},
		Ftag_type: int32(MAP),
	},
	75: {
		Ftag_name: [16]int8{'M', 'A', 'R', 'K'},
		Ftag_type: int32(MARK),
	},
	76: {
		Ftag_name: [16]int8{'M', 'A', 'T', 'H'},
		Ftag_type: int32(MATH),
	},
	77: {
		Ftag_name: [16]int8{'M', 'E', 'N', 'U'},
		Ftag_type: int32(MENU),
	},
	78: {
		Ftag_name: [16]int8{'M', 'E', 'T', 'E', 'R'},
		Ftag_type: int32(METER),
	},
	79: {
		Ftag_name: [16]int8{'N', 'A', 'V'},
		Ftag_type: int32(NAV),
	},
	80: {
		Ftag_name: [16]int8{'N', 'O', 'S', 'C', 'R', 'I', 'P', 'T'},
		Ftag_type: int32(NOSCRIPT),
	},
	81: {
		Ftag_name: [16]int8{'O', 'B', 'J', 'E', 'C', 'T'},
		Ftag_type: int32(OBJECT),
	},
	82: {
		Ftag_name: [16]int8{'O', 'L'},
		Ftag_type: int32(OL),
	},
	83: {
		Ftag_name: [16]int8{'O', 'P', 'T', 'G', 'R', 'O', 'U', 'P'},
		Ftag_type: int32(OPTGROUP),
	},
	84: {
		Ftag_name: [16]int8{'O', 'P', 'T', 'I', 'O', 'N'},
		Ftag_type: int32(OPTION),
	},
	85: {
		Ftag_name: [16]int8{'O', 'U', 'T', 'P', 'U', 'T'},
		Ftag_type: int32(OUTPUT),
	},
	86: {
		Ftag_name: [16]int8{'P'},
		Ftag_type: int32(P),
	},
	87: {
		Ftag_name: [16]int8{'P', 'I', 'C', 'T', 'U', 'R', 'E'},
		Ftag_type: int32(PICTURE),
	},
	88: {
		Ftag_name: [16]int8{'P', 'R', 'E'},
		Ftag_type: int32(PRE),
	},
	89: {
		Ftag_name: [16]int8{'P', 'R', 'O', 'G', 'R', 'E', 'S', 'S'},
		Ftag_type: int32(PROGRESS),
	},
	90: {
		Ftag_name: [16]int8{'Q'},
		Ftag_type: int32(Q),
	},
	91: {
		Ftag_name: [16]int8{'R', 'B'},
		Ftag_type: int32(RB),
	},
	92: {
		Ftag_name: [16]int8{'R', 'P'},
		Ftag_type: int32(RP),
	},
	93: {
		Ftag_name: [16]int8{'R', 'T'},
		Ftag_type: int32(RT),
	},
	94: {
		Ftag_name: [16]int8{'R', 'T', 'C'},
		Ftag_type: int32(RTC),
	},
	95: {
		Ftag_name: [16]int8{'R', 'U', 'B', 'Y'},
		Ftag_type: int32(RUBY),
	},
	96: {
		Ftag_name: [16]int8{'S'},
		Ftag_type: int32(S),
	},
	97: {
		Ftag_name: [16]int8{'S', 'A', 'M', 'P'},
		Ftag_type: int32(SAMP),
	},
	98: {
		Ftag_name: [16]int8{'S', 'C', 'R', 'I', 'P', 'T'},
		Ftag_type: int32(SCRIPT),
	},
	99: {
		Ftag_name: [16]int8{'S', 'E', 'C', 'T', 'I', 'O', 'N'},
		Ftag_type: int32(SECTION),
	},
	100: {
		Ftag_name: [16]int8{'S', 'E', 'L', 'E', 'C', 'T'},
		Ftag_type: int32(SELECT),
	},
	101: {
		Ftag_name: [16]int8{'S', 'L', 'O', 'T'},
		Ftag_type: int32(SLOT),
	},
	102: {
		Ftag_name: [16]int8{'S', 'M', 'A', 'L', 'L'},
		Ftag_type: int32(SMALL),
	},
	103: {
		Ftag_name: [16]int8{'S', 'P', 'A', 'N'},
		Ftag_type: int32(SPAN),
	},
	104: {
		Ftag_name: [16]int8{'S', 'T', 'R', 'O', 'N', 'G'},
		Ftag_type: int32(STRONG),
	},
	105: {
		Ftag_name: [16]int8{'S', 'T', 'Y', 'L', 'E'},
		Ftag_type: int32(STYLE),
	},
	106: {
		Ftag_name: [16]int8{'S', 'U', 'B'},
		Ftag_type: int32(SUB),
	},
	107: {
		Ftag_name: [16]int8{'S', 'U', 'M', 'M', 'A', 'R', 'Y'},
		Ftag_type: int32(SUMMARY),
	},
	108: {
		Ftag_name: [16]int8{'S', 'U', 'P'},
		Ftag_type: int32(SUP),
	},
	109: {
		Ftag_name: [16]int8{'S', 'V', 'G'},
		Ftag_type: int32(SVG),
	},
	110: {
		Ftag_name: [16]int8{'T', 'A', 'B', 'L', 'E'},
		Ftag_type: int32(TABLE),
	},
	111: {
		Ftag_name: [16]int8{'T', 'B', 'O', 'D', 'Y'},
		Ftag_type: int32(TBODY),
	},
	112: {
		Ftag_name: [16]int8{'T', 'D'},
		Ftag_type: int32(TD),
	},
	113: {
		Ftag_name: [16]int8{'T', 'E', 'M', 'P', 'L', 'A', 'T', 'E'},
		Ftag_type: int32(TEMPLATE),
	},
	114: {
		Ftag_name: [16]int8{'T', 'E', 'X', 'T', 'A', 'R', 'E', 'A'},
		Ftag_type: int32(TEXTAREA),
	},
	115: {
		Ftag_name: [16]int8{'T', 'F', 'O', 'O', 'T'},
		Ftag_type: int32(TFOOT),
	},
	116: {
		Ftag_name: [16]int8{'T', 'H'},
		Ftag_type: int32(TH),
	},
	117: {
		Ftag_name: [16]int8{'T', 'H', 'E', 'A', 'D'},
		Ftag_type: int32(THEAD),
	},
	118: {
		Ftag_name: [16]int8{'T', 'I', 'M', 'E'},
		Ftag_type: int32(TIME),
	},
	119: {
		Ftag_name: [16]int8{'T', 'I', 'T', 'L', 'E'},
		Ftag_type: int32(TITLE),
	},
	120: {
		Ftag_name: [16]int8{'T', 'R'},
		Ftag_type: int32(TR),
	},
	121: {
		Ftag_name: [16]int8{'U'},
		Ftag_type: int32(U),
	},
	122: {
		Ftag_name: [16]int8{'U', 'L'},
		Ftag_type: int32(UL),
	},
	123: {
		Ftag_name: [16]int8{'V', 'A', 'R'},
		Ftag_type: int32(VAR),
	},
	124: {
		Ftag_name: [16]int8{'V', 'I', 'D', 'E', 'O'},
		Ftag_type: int32(VIDEO),
	},
	125: {
		Ftag_name: [16]int8{'C', 'U', 'S', 'T', 'O', 'M'},
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
	if !(i < int32(26)) {
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

const START_TAG_NAME = 0
const SCRIPT_START_TAG_NAME = 1
const STYLE_START_TAG_NAME = 2
const END_TAG_NAME = 3
const ERRONEOUS_END_TAG_NAME = 4
const SELF_CLOSING_TAG_DELIMITER = 5
const IMPLICIT_END_TAG = 6
const RAW_TEXT = 7
const COMMENT = 8
const TEMPLATE_START_TAG_NAME = 9
const TEXT_FRAGMENT = 10
const INTERPOLATION_TEXT = 11

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

// #define advance(lexer)                                                                                                 //     do {                                                                                                               //         printf("advance %c, line: %d\n", lexer->lookahead, __LINE__);                                                  //         (lexer->advance)(lexer, false);                                                                                //     } while (0)
//
// #define skip(lexer) //     do { //         printf("skip %c, line: %d\n", lexer->lookahead, __LINE__); //         (lexer->advance)(lexer, true); //     } while (0)

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
	for libc.Xiswctype(tls, uint16((*TSLexer)(unsafe.Pointer(lexer)).Flookahead), uint16(libc.Int32FromInt32(0x0100)|libc.Int32FromInt32(_UPPER)|libc.Int32FromInt32(_LOWER)|libc.Int32FromInt32(_DIGIT))) != 0 || (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('-') || (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32(':') {
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
		*(*int8)(unsafe.Pointer((*String)(unsafe.Pointer(bp)).Fcontents + uintptr(v5))) = int8(libc.Xtowupper(tls, uint16((*TSLexer)(unsafe.Pointer(lexer)).Flookahead)))
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

func scan_raw_text(tls *libc.TLS, scanner uintptr, lexer uintptr) (r uint8) {
	var delimiter_index uint32
	var end_delimiter, v1 uintptr
	var v2 bool
	_, _, _, _ = delimiter_index, end_delimiter, v1, v2
	if (*Scanner)(unsafe.Pointer(scanner)).Ftags.Fsize == uint32(0) {
		return libc.BoolUint8(false1 != 0)
	}
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
		libc.X_assert(tls, __ccgo_ts, __ccgo_ts+64, uint32(165))
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
		v1 = __ccgo_ts + 75
	} else {
		v1 = __ccgo_ts + 84
	}
	end_delimiter = v1
	delimiter_index = uint32(0)
	for (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != 0 {
		if int32(libc.Xtowupper(tls, uint16((*TSLexer)(unsafe.Pointer(lexer)).Flookahead))) == int32(*(*int8)(unsafe.Pointer(end_delimiter + uintptr(delimiter_index)))) {
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
			libc.X_assert(tls, __ccgo_ts, __ccgo_ts+64, uint32(192))
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
				libc.X_assert(tls, __ccgo_ts, __ccgo_ts+64, uint32(216))
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
		// the case of malformed HTML)
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
	var new_capacity1, new_size, v11, v12 uint32_t
	var tag, tag1, tag2, v3, v6 Tag
	var v1, v5, v8, v9 uintptr
	var v10 size_t
	var v2 String
	var _ /* tag_name at bp+16 */ String
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = new_capacity1, new_size, tag, tag1, tag2, v1, v10, v11, v12, v2, v3, v5, v6, v8, v9
	*(*String)(unsafe.Pointer(bp + 16)) = scan_tag_name(tls, lexer)
	if (*(*String)(unsafe.Pointer(bp + 16))).Fsize == uint32(0) {
		v1 = bp + 16
		if (*Array)(unsafe.Pointer(v1)).Fcontents != 0 {
			libc.Xfree(tls, (*Array)(unsafe.Pointer(v1)).Fcontents)
			(*Array)(unsafe.Pointer(v1)).Fcontents = libc.UintptrFromInt32(0)
			(*Array)(unsafe.Pointer(v1)).Fsize = uint32(0)
			(*Array)(unsafe.Pointer(v1)).Fcapacity = uint32(0)
		}
		return libc.BoolUint8(false1 != 0)
	}
	v2 = *(*String)(unsafe.Pointer(bp + 16))
	*(*String)(unsafe.Pointer(bp)) = v2
	tag.Ftype_token = int32(END_)
	tag.Fcustom_tag_name = String{}
	v3 = tag
	goto _4
_4:
	tag1 = v3
	tag1.Ftype_token = tag_type_for_name(tls, bp)
	if tag1.Ftype_token == int32(CUSTOM) {
		tag1.Fcustom_tag_name = v2
	} else {
		v1 = bp
		if (*Array)(unsafe.Pointer(v1)).Fcontents != 0 {
			libc.Xfree(tls, (*Array)(unsafe.Pointer(v1)).Fcontents)
			(*Array)(unsafe.Pointer(v1)).Fcontents = libc.UintptrFromInt32(0)
			(*Array)(unsafe.Pointer(v1)).Fsize = uint32(0)
			(*Array)(unsafe.Pointer(v1)).Fcapacity = uint32(0)
		}
	}
	v6 = tag1
	goto _7
_7:
	tag2 = v6
	v5 = scanner
	new_size = (*Array)(unsafe.Pointer(v5)).Fsize + uint32(1)
	if new_size > (*Array)(unsafe.Pointer(v5)).Fcapacity {
		new_capacity1 = (*Array)(unsafe.Pointer(v5)).Fcapacity * uint32(2)
		if new_capacity1 < libc.Uint32FromInt32(8) {
			new_capacity1 = uint32(8)
		}
		if new_capacity1 < new_size {
			new_capacity1 = new_size
		}
		v8 = v5
		v10 = libc.Uint64FromInt64(24)
		v11 = new_capacity1
		if v11 > (*Array)(unsafe.Pointer(v8)).Fcapacity {
			if (*Array)(unsafe.Pointer(v8)).Fcontents != 0 {
				(*Array)(unsafe.Pointer(v8)).Fcontents = libc.Xrealloc(tls, (*Array)(unsafe.Pointer(v8)).Fcontents, uint64(v11)*v10)
			} else {
				(*Array)(unsafe.Pointer(v8)).Fcontents = libc.Xmalloc(tls, uint64(v11)*v10)
			}
			(*Array)(unsafe.Pointer(v8)).Fcapacity = v11
		}
	}
	v9 = scanner + 8
	v12 = *(*uint32_t)(unsafe.Pointer(v9))
	*(*uint32_t)(unsafe.Pointer(v9)) = *(*uint32_t)(unsafe.Pointer(v9)) + 1
	*(*Tag)(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(scanner)).Fcontents + uintptr(v12)*24)) = tag2
	switch tag2.Ftype_token {
	case int32(TEMPLATE):
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(TEMPLATE_START_TAG_NAME)
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
	var tag, tag1, v3, v6 Tag
	var v1, v5, v9 uintptr
	var v11 uint8
	var v13, v8 bool
	var v2 String
	var _ /* tag at bp+32 */ Tag
	var _ /* tag_name at bp+16 */ String
	_, _, _, _, _, _, _, _, _, _, _ = tag, tag1, v1, v11, v13, v2, v3, v5, v6, v8, v9
	*(*String)(unsafe.Pointer(bp + 16)) = scan_tag_name(tls, lexer)
	if (*(*String)(unsafe.Pointer(bp + 16))).Fsize == uint32(0) {
		v1 = bp + 16
		if (*Array)(unsafe.Pointer(v1)).Fcontents != 0 {
			libc.Xfree(tls, (*Array)(unsafe.Pointer(v1)).Fcontents)
			(*Array)(unsafe.Pointer(v1)).Fcontents = libc.UintptrFromInt32(0)
			(*Array)(unsafe.Pointer(v1)).Fsize = uint32(0)
			(*Array)(unsafe.Pointer(v1)).Fcapacity = uint32(0)
		}
		return libc.BoolUint8(false1 != 0)
	}
	v2 = *(*String)(unsafe.Pointer(bp + 16))
	*(*String)(unsafe.Pointer(bp)) = v2
	tag.Ftype_token = int32(END_)
	tag.Fcustom_tag_name = String{}
	v3 = tag
	goto _4
_4:
	tag1 = v3
	tag1.Ftype_token = tag_type_for_name(tls, bp)
	if tag1.Ftype_token == int32(CUSTOM) {
		tag1.Fcustom_tag_name = v2
	} else {
		v1 = bp
		if (*Array)(unsafe.Pointer(v1)).Fcontents != 0 {
			libc.Xfree(tls, (*Array)(unsafe.Pointer(v1)).Fcontents)
			(*Array)(unsafe.Pointer(v1)).Fcontents = libc.UintptrFromInt32(0)
			(*Array)(unsafe.Pointer(v1)).Fsize = uint32(0)
			(*Array)(unsafe.Pointer(v1)).Fcapacity = uint32(0)
		}
	}
	v6 = tag1
	goto _7
_7:
	*(*Tag)(unsafe.Pointer(bp + 32)) = v6
	if v13 = (*Scanner)(unsafe.Pointer(scanner)).Ftags.Fsize > uint32(0); v13 {
		if v8 = !!((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(scanner)).Fsize-libc.Uint32FromInt32(1) < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(scanner)).Fsize); !v8 {
			libc.X_assert(tls, __ccgo_ts, __ccgo_ts+64, uint32(281))
		}
		_ = v8 || libc.Bool(libc.Int32FromInt32(0) != 0)
		v5 = (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(scanner)).Fcontents + uintptr((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(scanner)).Fsize-uint32(1))*24
		v9 = bp + 32
		if (*Tag)(unsafe.Pointer(v5)).Ftype_token != (*Tag)(unsafe.Pointer(v9)).Ftype_token {
			v11 = libc.BoolUint8(false1 != 0)
			goto _12
		}
		if (*Tag)(unsafe.Pointer(v5)).Ftype_token == int32(CUSTOM) {
			if (*Tag)(unsafe.Pointer(v5)).Fcustom_tag_name.Fsize != (*Tag)(unsafe.Pointer(v9)).Fcustom_tag_name.Fsize {
				v11 = libc.BoolUint8(false1 != 0)
				goto _12
			}
			if libc.Xmemcmp(tls, (*Tag)(unsafe.Pointer(v5)).Fcustom_tag_name.Fcontents, (*Tag)(unsafe.Pointer(v9)).Fcustom_tag_name.Fcontents, uint64((*Tag)(unsafe.Pointer(v5)).Fcustom_tag_name.Fsize)) != libc.Int32FromInt32(0) {
				v11 = libc.BoolUint8(false1 != 0)
				goto _12
			}
		}
		v11 = libc.BoolUint8(true1 != 0)
		goto _12
	_12:
	}
	if v13 && v11 != 0 {
		pop_tag(tls, scanner)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(END_TAG_NAME)
	} else {
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(ERRONEOUS_END_TAG_NAME)
	}
	v1 = bp + 32
	if (*Tag)(unsafe.Pointer(v1)).Ftype_token == int32(CUSTOM) {
		v5 = v1 + 8
		if (*Array)(unsafe.Pointer(v5)).Fcontents != 0 {
			libc.Xfree(tls, (*Array)(unsafe.Pointer(v5)).Fcontents)
			(*Array)(unsafe.Pointer(v5)).Fcontents = libc.UintptrFromInt32(0)
			(*Array)(unsafe.Pointer(v5)).Fsize = uint32(0)
			(*Array)(unsafe.Pointer(v5)).Fcapacity = uint32(0)
		}
	}
	return libc.BoolUint8(true1 != 0)
}

func scan_self_closing_tag_delimiter(tls *libc.TLS, scanner uintptr, lexer uintptr) (r uint8) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var v1 uint32_t
	var v2, v3, v4 uintptr
	var _ /* last_tag at bp+0 */ Tag
	_, _, _, _ = v1, v2, v3, v4
	advance(tls, lexer)
	if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('>') {
		advance(tls, lexer)
		if (*Scanner)(unsafe.Pointer(scanner)).Ftags.Fsize > uint32(0) {
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
			(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(SELF_CLOSING_TAG_DELIMITER)
		}
		return libc.BoolUint8(true1 != 0)
	}
	return libc.BoolUint8(false1 != 0)
}

func scan(tls *libc.TLS, scanner uintptr, lexer uintptr, valid_symbols uintptr) (r uint8) {
	var advanced_once, is_error_recovery uint8
	var v1 int32
	_, _, _ = advanced_once, is_error_recovery, v1
	is_error_recovery = libc.BoolUint8(*(*uint8)(unsafe.Pointer(valid_symbols + uintptr(START_TAG_NAME))) != 0 && *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(RAW_TEXT))) != 0)
	if !(is_error_recovery != 0) && (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32('<') && (*(*uint8)(unsafe.Pointer(valid_symbols + uintptr(TEXT_FRAGMENT))) != 0 || *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(INTERPOLATION_TEXT))) != 0) {
		advanced_once = libc.BoolUint8(false1 != 0)
		if !(*(*uint8)(unsafe.Pointer(valid_symbols + uintptr(COMMENT))) != 0) {
			for libc.Xiswctype(tls, uint16((*TSLexer)(unsafe.Pointer(lexer)).Flookahead), uint16(_SPACE)) != 0 {
				skip(tls, lexer)
			}
		}
		for !((*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Feof})))(tls, lexer) != 0) {
			switch (*TSLexer)(unsafe.Pointer(lexer)).Flookahead {
			case int32('<'):
				(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
				advance(tls, lexer)
				if libc.Xiswctype(tls, uint16((*TSLexer)(unsafe.Pointer(lexer)).Flookahead), uint16(libc.Int32FromInt32(0x0100)|libc.Int32FromInt32(_UPPER)|libc.Int32FromInt32(_LOWER))) != 0 || (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('!') || (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('?') || (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('/') {
					goto loop_exit
				}
				advanced_once = libc.BoolUint8(true1 != 0)
			case int32('{'):
				(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
				advance(tls, lexer)
				if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('{') {
					goto loop_exit
				}
				advanced_once = libc.BoolUint8(true1 != 0)
			case int32('}'):
				if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(INTERPOLATION_TEXT))) != 0 {
					(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
					advance(tls, lexer)
					if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('}') {
						(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(INTERPOLATION_TEXT)
						return advanced_once
					}
					break
				}
				advance(tls, lexer)
				advanced_once = libc.BoolUint8(true1 != 0)
			case int32('\r'):
				advance(tls, lexer)
				// Mirror `default` case behavior if this is _not_ a CRLF, otherwise fallthrough and handle newline
				if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32('\n') {
					advanced_once = libc.BoolUint8(true1 != 0)
					advance(tls, lexer)
					break
				}
				fallthrough
			case int32('\n'):
				if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(TEXT_FRAGMENT))) != 0 {
					(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
					for libc.Xiswctype(tls, uint16((*TSLexer)(unsafe.Pointer(lexer)).Flookahead), uint16(_SPACE)) != 0 {
						if advanced_once != 0 {
							advance(tls, lexer)
						} else {
							skip(tls, lexer)
						}
					}
					if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('<') || (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('>') {
						goto loop_exit
					}
				} else {
					advance(tls, lexer)
				}
			default:
				advanced_once = uint8(libc.BoolInt32(int32(advanced_once)|libc.BoolInt32((*TSLexer)(unsafe.Pointer(lexer)).Flookahead != libc.Int32FromUint8('\n')) != 0))
				advance(tls, lexer)
				break
			}
		}
		if (*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Feof})))(tls, lexer) != 0 {
			return libc.BoolUint8(false1 != 0)
		}
		goto loop_exit
	loop_exit:
		;
		if advanced_once != 0 {
			(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(TEXT_FRAGMENT)
			return libc.BoolUint8(true1 != 0)
		}
	}
	if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(RAW_TEXT))) != 0 && !(*(*uint8)(unsafe.Pointer(valid_symbols + uintptr(START_TAG_NAME))) != 0) && !(*(*uint8)(unsafe.Pointer(valid_symbols + uintptr(END_TAG_NAME))) != 0) {
		return scan_raw_text(tls, scanner, lexer)
	}
	for libc.Xiswctype(tls, uint16((*TSLexer)(unsafe.Pointer(lexer)).Flookahead), uint16(_SPACE)) != 0 {
		skip(tls, lexer)
	}
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
	case int32('\000'):
		if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(IMPLICIT_END_TAG))) != 0 {
			return scan_implicit_end_tag(tls, scanner, lexer)
		}
	case int32('/'):
		if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(SELF_CLOSING_TAG_DELIMITER))) != 0 {
			return scan_self_closing_tag_delimiter(tls, scanner, lexer)
		}
	default:
		if (*(*uint8)(unsafe.Pointer(valid_symbols + uintptr(START_TAG_NAME))) != 0 || *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(END_TAG_NAME))) != 0) && !(*(*uint8)(unsafe.Pointer(valid_symbols + uintptr(RAW_TEXT))) != 0) {
			if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(START_TAG_NAME))) != 0 {
				v1 = int32(scan_start_tag_name(tls, scanner, lexer))
			} else {
				v1 = int32(scan_end_tag_name(tls, scanner, lexer))
			}
			return uint8(libc.BoolInt32(v1 != 0))
		}
	}
	return libc.BoolUint8(false1 != 0)
}

func tree_sitter_vue_external_scanner_create(tls *libc.TLS) (r uintptr) {
	var scanner uintptr
	_ = scanner
	scanner = libc.Xcalloc(tls, uint64(1), uint64(16))
	return scanner
}

func tree_sitter_vue_external_scanner_scan(tls *libc.TLS, payload uintptr, lexer uintptr, valid_symbols uintptr) (r uint8) {
	var scanner uintptr
	_ = scanner
	scanner = payload
	return scan(tls, scanner, lexer, valid_symbols)
}

func tree_sitter_vue_external_scanner_serialize(tls *libc.TLS, payload uintptr, buffer uintptr) (r uint32) {
	var scanner uintptr
	_ = scanner
	scanner = payload
	return serialize(tls, scanner, buffer)
}

func tree_sitter_vue_external_scanner_deserialize(tls *libc.TLS, payload uintptr, buffer uintptr, length uint32) {
	var scanner uintptr
	_ = scanner
	scanner = payload
	deserialize(tls, scanner, buffer, length)
}

func tree_sitter_vue_external_scanner_destroy(tls *libc.TLS, payload uintptr) {
	var i uint32
	var scanner, v3, v4 uintptr
	var v2 bool
	_, _, _, _, _ = i, scanner, v2, v3, v4
	scanner = payload
	i = uint32(0)
	for {
		if !(i < (*Scanner)(unsafe.Pointer(scanner)).Ftags.Fsize) {
			break
		}
		if v2 = !!(i < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(scanner)).Fsize); !v2 {
			libc.X_assert(tls, __ccgo_ts+92, __ccgo_ts+64, uint32(465))
		}
		_ = v2 || libc.Bool(libc.Int32FromInt32(0) != 0)
		v3 = (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(scanner)).Fcontents + uintptr(i)*24
		if (*Tag)(unsafe.Pointer(v3)).Ftype_token == int32(CUSTOM) {
			v4 = v3 + 8
			if (*Array)(unsafe.Pointer(v4)).Fcontents != 0 {
				libc.Xfree(tls, (*Array)(unsafe.Pointer(v4)).Fcontents)
				(*Array)(unsafe.Pointer(v4)).Fcontents = libc.UintptrFromInt32(0)
				(*Array)(unsafe.Pointer(v4)).Fsize = uint32(0)
				(*Array)(unsafe.Pointer(v4)).Fcapacity = uint32(0)
			}
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	v3 = scanner
	if (*Array)(unsafe.Pointer(v3)).Fcontents != 0 {
		libc.Xfree(tls, (*Array)(unsafe.Pointer(v3)).Fcontents)
		(*Array)(unsafe.Pointer(v3)).Fcontents = libc.UintptrFromInt32(0)
		(*Array)(unsafe.Pointer(v3)).Fsize = uint32(0)
		(*Array)(unsafe.Pointer(v3)).Fcapacity = uint32(0)
	}
	libc.Xfree(tls, scanner)
}

/* Automatically @generated by tree-sitter */

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
const anon_sym_LBRACE_LBRACE = 16
const anon_sym_RBRACE_RBRACE = 17
const anon_sym_COLON = 18
const anon_sym_COLON2 = 19
const anon_sym_DOT = 20
const anon_sym_AT = 21
const anon_sym_POUND = 22
const sym_directive_name = 23
const sym_directive_value = 24
const anon_sym_LBRACK = 25
const anon_sym_RBRACK = 26
const sym_dynamic_directive_inner_value = 27
const anon_sym_DOT2 = 28
const sym_directive_modifier = 29
const sym__start_tag_name = 30
const sym__script_start_tag_name = 31
const sym__style_start_tag_name = 32
const sym__end_tag_name = 33
const sym_erroneous_end_tag_name = 34
const sym__implicit_end_tag = 35
const sym_raw_text = 36
const sym_comment = 37
const sym__template_start_tag_name = 38
const sym__text_fragment = 39
const sym__interpolation_text = 40
const sym_document = 41
const sym_doctype = 42
const sym__node = 43
const sym_element = 44
const sym_script_element = 45
const sym_style_element = 46
const sym_start_tag = 47
const sym_script_start_tag = 48
const sym_style_start_tag = 49
const sym_self_closing_tag = 50
const sym_end_tag = 51
const sym_erroneous_end_tag = 52
const sym_attribute = 53
const sym_quoted_attribute_value = 54
const sym_text = 55
const sym_template_element = 56
const sym_template_start_tag = 57
const sym_interpolation = 58
const sym__attribute = 59
const sym_directive_attribute = 60
const sym_dynamic_directive_value = 61
const sym_directive_modifiers = 62
const aux_sym_document_repeat1 = 63
const aux_sym_start_tag_repeat1 = 64
const aux_sym_directive_attribute_repeat1 = 65
const aux_sym_directive_modifiers_repeat1 = 66

var ts_symbol_names = [67]uintptr{
	0:  __ccgo_ts + 131,
	1:  __ccgo_ts + 135,
	2:  __ccgo_ts + 138,
	3:  __ccgo_ts + 153,
	4:  __ccgo_ts + 155,
	5:  __ccgo_ts + 163,
	6:  __ccgo_ts + 165,
	7:  __ccgo_ts + 168,
	8:  __ccgo_ts + 171,
	9:  __ccgo_ts + 173,
	10: __ccgo_ts + 188,
	11: __ccgo_ts + 204,
	12: __ccgo_ts + 211,
	13: __ccgo_ts + 188,
	14: __ccgo_ts + 213,
	15: __ccgo_ts + 188,
	16: __ccgo_ts + 215,
	17: __ccgo_ts + 218,
	18: __ccgo_ts + 221,
	19: __ccgo_ts + 221,
	20: __ccgo_ts + 223,
	21: __ccgo_ts + 225,
	22: __ccgo_ts + 227,
	23: __ccgo_ts + 229,
	24: __ccgo_ts + 244,
	25: __ccgo_ts + 260,
	26: __ccgo_ts + 262,
	27: __ccgo_ts + 264,
	28: __ccgo_ts + 223,
	29: __ccgo_ts + 294,
	30: __ccgo_ts + 313,
	31: __ccgo_ts + 313,
	32: __ccgo_ts + 313,
	33: __ccgo_ts + 313,
	34: __ccgo_ts + 322,
	35: __ccgo_ts + 345,
	36: __ccgo_ts + 363,
	37: __ccgo_ts + 372,
	38: __ccgo_ts + 313,
	39: __ccgo_ts + 380,
	40: __ccgo_ts + 363,
	41: __ccgo_ts + 395,
	42: __ccgo_ts + 155,
	43: __ccgo_ts + 404,
	44: __ccgo_ts + 410,
	45: __ccgo_ts + 418,
	46: __ccgo_ts + 433,
	47: __ccgo_ts + 447,
	48: __ccgo_ts + 447,
	49: __ccgo_ts + 447,
	50: __ccgo_ts + 457,
	51: __ccgo_ts + 474,
	52: __ccgo_ts + 482,
	53: __ccgo_ts + 500,
	54: __ccgo_ts + 510,
	55: __ccgo_ts + 533,
	56: __ccgo_ts + 538,
	57: __ccgo_ts + 447,
	58: __ccgo_ts + 555,
	59: __ccgo_ts + 569,
	60: __ccgo_ts + 580,
	61: __ccgo_ts + 600,
	62: __ccgo_ts + 624,
	63: __ccgo_ts + 644,
	64: __ccgo_ts + 661,
	65: __ccgo_ts + 679,
	66: __ccgo_ts + 707,
}

var ts_symbol_map = [67]TSSymbol{
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
	16: uint16(anon_sym_LBRACE_LBRACE),
	17: uint16(anon_sym_RBRACE_RBRACE),
	18: uint16(anon_sym_COLON),
	19: uint16(anon_sym_COLON),
	20: uint16(anon_sym_DOT),
	21: uint16(anon_sym_AT),
	22: uint16(anon_sym_POUND),
	23: uint16(sym_directive_name),
	24: uint16(sym_directive_value),
	25: uint16(anon_sym_LBRACK),
	26: uint16(anon_sym_RBRACK),
	27: uint16(sym_dynamic_directive_inner_value),
	28: uint16(anon_sym_DOT),
	29: uint16(sym_directive_modifier),
	30: uint16(sym__start_tag_name),
	31: uint16(sym__start_tag_name),
	32: uint16(sym__start_tag_name),
	33: uint16(sym__start_tag_name),
	34: uint16(sym_erroneous_end_tag_name),
	35: uint16(sym__implicit_end_tag),
	36: uint16(sym_raw_text),
	37: uint16(sym_comment),
	38: uint16(sym__start_tag_name),
	39: uint16(sym__text_fragment),
	40: uint16(sym_raw_text),
	41: uint16(sym_document),
	42: uint16(sym_doctype),
	43: uint16(sym__node),
	44: uint16(sym_element),
	45: uint16(sym_script_element),
	46: uint16(sym_style_element),
	47: uint16(sym_start_tag),
	48: uint16(sym_start_tag),
	49: uint16(sym_start_tag),
	50: uint16(sym_self_closing_tag),
	51: uint16(sym_end_tag),
	52: uint16(sym_erroneous_end_tag),
	53: uint16(sym_attribute),
	54: uint16(sym_quoted_attribute_value),
	55: uint16(sym_text),
	56: uint16(sym_template_element),
	57: uint16(sym_start_tag),
	58: uint16(sym_interpolation),
	59: uint16(sym__attribute),
	60: uint16(sym_directive_attribute),
	61: uint16(sym_dynamic_directive_value),
	62: uint16(sym_directive_modifiers),
	63: uint16(aux_sym_document_repeat1),
	64: uint16(aux_sym_start_tag_repeat1),
	65: uint16(aux_sym_directive_attribute_repeat1),
	66: uint16(aux_sym_directive_modifiers_repeat1),
}

var ts_symbol_metadata = [67]TSSymbolMetadata{
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
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	24: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	25: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	26: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	27: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	28: {
		Fvisible: libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
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
	63: {},
	64: {},
	65: {},
	66: {},
}

var ts_alias_sequences = [1][6]TSSymbol{}

var ts_non_terminal_alias_map = [1]uint16_t{}

var ts_primary_state_ids = [170]TSStateId{
	1:   uint16(1),
	2:   uint16(2),
	3:   uint16(3),
	4:   uint16(2),
	5:   uint16(3),
	6:   uint16(6),
	7:   uint16(7),
	8:   uint16(8),
	9:   uint16(9),
	10:  uint16(9),
	11:  uint16(6),
	12:  uint16(8),
	13:  uint16(13),
	14:  uint16(14),
	15:  uint16(15),
	16:  uint16(16),
	17:  uint16(17),
	18:  uint16(18),
	19:  uint16(16),
	20:  uint16(15),
	21:  uint16(21),
	22:  uint16(14),
	23:  uint16(18),
	24:  uint16(24),
	25:  uint16(25),
	26:  uint16(26),
	27:  uint16(27),
	28:  uint16(21),
	29:  uint16(29),
	30:  uint16(13),
	31:  uint16(17),
	32:  uint16(32),
	33:  uint16(33),
	34:  uint16(34),
	35:  uint16(29),
	36:  uint16(33),
	37:  uint16(34),
	38:  uint16(38),
	39:  uint16(39),
	40:  uint16(32),
	41:  uint16(41),
	42:  uint16(42),
	43:  uint16(43),
	44:  uint16(42),
	45:  uint16(38),
	46:  uint16(41),
	47:  uint16(39),
	48:  uint16(48),
	49:  uint16(49),
	50:  uint16(50),
	51:  uint16(51),
	52:  uint16(49),
	53:  uint16(53),
	54:  uint16(43),
	55:  uint16(55),
	56:  uint16(56),
	57:  uint16(57),
	58:  uint16(48),
	59:  uint16(59),
	60:  uint16(60),
	61:  uint16(61),
	62:  uint16(62),
	63:  uint16(63),
	64:  uint16(64),
	65:  uint16(65),
	66:  uint16(66),
	67:  uint16(60),
	68:  uint16(68),
	69:  uint16(69),
	70:  uint16(70),
	71:  uint16(71),
	72:  uint16(68),
	73:  uint16(73),
	74:  uint16(63),
	75:  uint16(65),
	76:  uint16(66),
	77:  uint16(70),
	78:  uint16(78),
	79:  uint16(79),
	80:  uint16(80),
	81:  uint16(81),
	82:  uint16(82),
	83:  uint16(83),
	84:  uint16(84),
	85:  uint16(62),
	86:  uint16(71),
	87:  uint16(64),
	88:  uint16(51),
	89:  uint16(78),
	90:  uint16(57),
	91:  uint16(79),
	92:  uint16(80),
	93:  uint16(56),
	94:  uint16(81),
	95:  uint16(82),
	96:  uint16(59),
	97:  uint16(55),
	98:  uint16(53),
	99:  uint16(50),
	100: uint16(83),
	101: uint16(84),
	102: uint16(73),
	103: uint16(103),
	104: uint16(104),
	105: uint16(105),
	106: uint16(106),
	107: uint16(107),
	108: uint16(108),
	109: uint16(109),
	110: uint16(110),
	111: uint16(110),
	112: uint16(107),
	113: uint16(109),
	114: uint16(105),
	115: uint16(106),
	116: uint16(108),
	117: uint16(117),
	118: uint16(118),
	119: uint16(119),
	120: uint16(117),
	121: uint16(118),
	122: uint16(122),
	123: uint16(119),
	124: uint16(122),
	125: uint16(125),
	126: uint16(126),
	127: uint16(127),
	128: uint16(128),
	129: uint16(129),
	130: uint16(130),
	131: uint16(130),
	132: uint16(132),
	133: uint16(126),
	134: uint16(134),
	135: uint16(135),
	136: uint16(136),
	137: uint16(125),
	138: uint16(136),
	139: uint16(134),
	140: uint16(129),
	141: uint16(141),
	142: uint16(126),
	143: uint16(126),
	144: uint16(135),
	145: uint16(145),
	146: uint16(146),
	147: uint16(147),
	148: uint16(148),
	149: uint16(149),
	150: uint16(150),
	151: uint16(151),
	152: uint16(152),
	153: uint16(153),
	154: uint16(148),
	155: uint16(155),
	156: uint16(146),
	157: uint16(147),
	158: uint16(149),
	159: uint16(155),
	160: uint16(160),
	161: uint16(160),
	162: uint16(162),
	163: uint16(163),
	164: uint16(163),
	165: uint16(153),
	166: uint16(152),
	167: uint16(150),
	168: uint16(162),
	169: uint16(145),
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
			state = uint16(27)
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
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(26)
			goto next_state
		}
		return result
	case int32(1):
		if lookahead == int32('"') {
			state = uint16(84)
			goto next_state
		}
		if lookahead == int32('\'') {
			state = uint16(81)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(1)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('<') || int32('>') < lookahead) {
			state = uint16(40)
			goto next_state
		}
		return result
	case int32(2):
		if lookahead == int32('"') {
			state = uint16(84)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(85)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(86)
			goto next_state
		}
		return result
	case int32(3):
		i1 = uint32(0)
		for {
			if !(uint64(i1) < libc.Uint64FromInt64(32)/libc.Uint64FromInt64(2)) {
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
			state = uint16(5)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('"') && lookahead != int32('#') && lookahead != int32('\'') && (lookahead < int32('<') || int32('>') < lookahead) {
			state = uint16(39)
			goto next_state
		}
		return result
	case int32(4):
		i2 = uint32(0)
		for {
			if !(uint64(i2) < libc.Uint64FromInt64(32)/libc.Uint64FromInt64(2)) {
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
			state = uint16(5)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('"') && lookahead != int32('#') && lookahead != int32('\'') && (lookahead < int32('<') || int32('>') < lookahead) {
			state = uint16(39)
			goto next_state
		}
		return result
	case int32(5):
		i3 = uint32(0)
		for {
			if !(uint64(i3) < libc.Uint64FromInt64(32)/libc.Uint64FromInt64(2)) {
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
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(5)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('"') && lookahead != int32('#') && lookahead != int32('\'') && (lookahead < int32('<') || int32('>') < lookahead) {
			state = uint16(39)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead == int32('#') {
			state = uint16(19)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(80)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead == int32('\'') {
			state = uint16(81)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(82)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(83)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead == int32('-') {
			state = uint16(24)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead == int32('>') {
			state = uint16(34)
			goto next_state
		}
		return result
	case int32(10):
		if lookahead == int32('[') {
			state = uint16(96)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(21)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('"') && lookahead != int32('\'') && lookahead != int32('.') && lookahead != int32('/') && lookahead != int32(':') && (lookahead < int32('<') || int32('>') < lookahead) {
			state = uint16(95)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead == int32(']') {
			state = uint16(97)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32('"') && lookahead != int32('\'') && lookahead != int32('/') && (lookahead < int32('<') || int32('>') < lookahead) {
			state = uint16(98)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead == int32('{') {
			state = uint16(87)
			goto next_state
		}
		return result
	case int32(13):
		if lookahead == int32('}') {
			state = uint16(88)
			goto next_state
		}
		return result
	case int32(14):
		if lookahead == int32('C') || lookahead == int32('c') {
			state = uint16(18)
			goto next_state
		}
		return result
	case int32(15):
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(32)
			goto next_state
		}
		return result
	case int32(16):
		if lookahead == int32('O') || lookahead == int32('o') {
			state = uint16(14)
			goto next_state
		}
		return result
	case int32(17):
		if lookahead == int32('P') || lookahead == int32('p') {
			state = uint16(15)
			goto next_state
		}
		return result
	case int32(18):
		if lookahead == int32('T') || lookahead == int32('t') {
			state = uint16(20)
			goto next_state
		}
		return result
	case int32(19):
		if lookahead == int32('X') || lookahead == int32('x') {
			state = uint16(23)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(46)
			goto next_state
		}
		return result
	case int32(20):
		if lookahead == int32('Y') || lookahead == int32('y') {
			state = uint16(17)
			goto next_state
		}
		return result
	case int32(21):
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(21)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('"') && lookahead != int32('\'') && lookahead != int32('.') && lookahead != int32('/') && lookahead != int32(':') && (lookahead < int32('<') || int32('>') < lookahead) {
			state = uint16(95)
			goto next_state
		}
		return result
	case int32(22):
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(29)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('>') {
			state = uint16(30)
			goto next_state
		}
		return result
	case int32(23):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(51)
			goto next_state
		}
		return result
	case int32(24):
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32('"') && lookahead != int32('\'') && lookahead != int32('.') && lookahead != int32('/') && lookahead != int32(':') && (lookahead < int32('<') || int32('>') < lookahead) {
			state = uint16(94)
			goto next_state
		}
		return result
	case int32(25):
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32('"') && lookahead != int32('\'') && lookahead != int32('.') && lookahead != int32('/') && (lookahead < int32('<') || int32('>') < lookahead) {
			state = uint16(100)
			goto next_state
		}
		return result
	case int32(26):
		if eof != 0 {
			state = uint16(27)
			goto next_state
		}
		i4 = uint32(0)
		for {
			if !(uint64(i4) < libc.Uint64FromInt64(64)/libc.Uint64FromInt64(2)) {
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
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(26)
			goto next_state
		}
		return result
	case int32(27):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(ts_builtin_sym_end)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(28):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT_BANG)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(29):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_doctype_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(29)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('>') {
			state = uint16(30)
			goto next_state
		}
		return result
	case int32(30):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_doctype_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('>') {
			state = uint16(30)
			goto next_state
		}
		return result
	case int32(31):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_GT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(32):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__doctype)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(33):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('!') {
			state = uint16(28)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(35)
			goto next_state
		}
		return result
	case int32(34):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SLASH_GT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(35):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT_SLASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(36):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(37):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_attribute_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') {
			state = uint16(38)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32('"') && lookahead != int32('\'') && lookahead != int32('/') && (lookahead < int32('<') || int32('>') < lookahead) {
			state = uint16(39)
			goto next_state
		}
		return result
	case int32(38):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_attribute_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') || lookahead == int32(':') {
			state = uint16(39)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32('"') && lookahead != int32('\'') && lookahead != int32('.') && lookahead != int32('/') && (lookahead < int32('<') || int32('>') < lookahead) {
			state = uint16(94)
			goto next_state
		}
		return result
	case int32(39):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_attribute_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32('"') && lookahead != int32('\'') && lookahead != int32('/') && (lookahead < int32('<') || int32('>') < lookahead) {
			state = uint16(39)
			goto next_state
		}
		return result
	case int32(40):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_attribute_value)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32('"') && lookahead != int32('\'') && (lookahead < int32('<') || int32('>') < lookahead) {
			state = uint16(40)
			goto next_state
		}
		return result
	case int32(41):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(42):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(41)
			goto next_state
		}
		return result
	case int32(43):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(41)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(42)
			goto next_state
		}
		return result
	case int32(44):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(41)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(43)
			goto next_state
		}
		return result
	case int32(45):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(41)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(44)
			goto next_state
		}
		return result
	case int32(46):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(41)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(45)
			goto next_state
		}
		return result
	case int32(47):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(41)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(42)
			goto next_state
		}
		return result
	case int32(48):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(41)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(47)
			goto next_state
		}
		return result
	case int32(49):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(41)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(48)
			goto next_state
		}
		return result
	case int32(50):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(41)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(49)
			goto next_state
		}
		return result
	case int32(51):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(41)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(50)
			goto next_state
		}
		return result
	case int32(52):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(41)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(42)
			goto next_state
		}
		return result
	case int32(53):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(41)
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
			state = uint16(41)
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
			state = uint16(41)
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
			state = uint16(41)
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
			state = uint16(41)
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
			state = uint16(41)
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
			state = uint16(41)
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
			state = uint16(41)
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
			state = uint16(41)
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
			state = uint16(41)
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
			state = uint16(41)
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
			state = uint16(41)
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
			state = uint16(41)
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
			state = uint16(41)
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
			state = uint16(41)
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
			state = uint16(41)
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
			state = uint16(41)
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
			state = uint16(41)
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
			state = uint16(41)
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
			state = uint16(41)
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
			state = uint16(41)
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
			state = uint16(41)
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
			state = uint16(41)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(74)
			goto next_state
		}
		return result
	case int32(76):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(41)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(75)
			goto next_state
		}
		return result
	case int32(77):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(41)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(76)
			goto next_state
		}
		return result
	case int32(78):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(41)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(77)
			goto next_state
		}
		return result
	case int32(79):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(41)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(78)
			goto next_state
		}
		return result
	case int32(80):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(41)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(79)
			goto next_state
		}
		return result
	case int32(81):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(82):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_quoted_attribute_value_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(82)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\'') {
			state = uint16(83)
			goto next_state
		}
		return result
	case int32(83):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_quoted_attribute_value_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\'') {
			state = uint16(83)
			goto next_state
		}
		return result
	case int32(84):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(85):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_quoted_attribute_value_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(85)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('"') {
			state = uint16(86)
			goto next_state
		}
		return result
	case int32(86):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_quoted_attribute_value_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('"') {
			state = uint16(86)
			goto next_state
		}
		return result
	case int32(87):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACE_LBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(88):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACE_RBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(89):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COLON)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(90):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COLON2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(91):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(92):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(93):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUND)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(94):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_directive_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32('"') && lookahead != int32('\'') && lookahead != int32('.') && lookahead != int32('/') && lookahead != int32(':') && (lookahead < int32('<') || int32('>') < lookahead) {
			state = uint16(94)
			goto next_state
		}
		return result
	case int32(95):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_directive_value)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32('"') && lookahead != int32('\'') && lookahead != int32('.') && lookahead != int32('/') && lookahead != int32(':') && (lookahead < int32('<') || int32('>') < lookahead) {
			state = uint16(95)
			goto next_state
		}
		return result
	case int32(96):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(97):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(98):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_dynamic_directive_inner_value)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32('"') && lookahead != int32('\'') && lookahead != int32('/') && (lookahead < int32('<') || int32('>') < lookahead) && lookahead != int32(']') {
			state = uint16(98)
			goto next_state
		}
		return result
	case int32(99):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOT2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(100):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_directive_modifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32('"') && lookahead != int32('\'') && lookahead != int32('.') && lookahead != int32('/') && (lookahead < int32('<') || int32('>') < lookahead) {
			state = uint16(100)
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
	1:  uint16(84),
	2:  uint16('#'),
	3:  uint16(93),
	4:  uint16('&'),
	5:  uint16(6),
	6:  uint16('\''),
	7:  uint16(81),
	8:  uint16('.'),
	9:  uint16(99),
	10: uint16('/'),
	11: uint16(9),
	12: uint16(':'),
	13: uint16(89),
	14: uint16('<'),
	15: uint16(33),
	16: uint16('='),
	17: uint16(36),
	18: uint16('>'),
	19: uint16(31),
	20: uint16('@'),
	21: uint16(92),
	22: uint16('['),
	23: uint16(96),
	24: uint16(']'),
	25: uint16(97),
	26: uint16('v'),
	27: uint16(8),
	28: uint16('{'),
	29: uint16(12),
	30: uint16('}'),
	31: uint16(13),
	32: uint16('D'),
	33: uint16(16),
	34: uint16('d'),
	35: uint16(16),
}

var map_token1 = [16]uint16_t{
	0:  uint16('#'),
	1:  uint16(93),
	2:  uint16('.'),
	3:  uint16(99),
	4:  uint16('/'),
	5:  uint16(9),
	6:  uint16(':'),
	7:  uint16(89),
	8:  uint16('='),
	9:  uint16(36),
	10: uint16('>'),
	11: uint16(31),
	12: uint16('@'),
	13: uint16(92),
	14: uint16('v'),
	15: uint16(37),
}

var map_token2 = [16]uint16_t{
	0:  uint16('#'),
	1:  uint16(93),
	2:  uint16('.'),
	3:  uint16(99),
	4:  uint16('/'),
	5:  uint16(9),
	6:  uint16(':'),
	7:  uint16(90),
	8:  uint16('='),
	9:  uint16(36),
	10: uint16('>'),
	11: uint16(31),
	12: uint16('@'),
	13: uint16(92),
	14: uint16('v'),
	15: uint16(37),
}

var map_token3 = [16]uint16_t{
	0:  uint16('#'),
	1:  uint16(93),
	2:  uint16('.'),
	3:  uint16(91),
	4:  uint16('/'),
	5:  uint16(9),
	6:  uint16(':'),
	7:  uint16(90),
	8:  uint16('='),
	9:  uint16(36),
	10: uint16('>'),
	11: uint16(31),
	12: uint16('@'),
	13: uint16(92),
	14: uint16('v'),
	15: uint16(37),
}

var map_token4 = [32]uint16_t{
	0:  uint16('"'),
	1:  uint16(84),
	2:  uint16('#'),
	3:  uint16(93),
	4:  uint16('&'),
	5:  uint16(6),
	6:  uint16('\''),
	7:  uint16(81),
	8:  uint16('.'),
	9:  uint16(91),
	10: uint16('/'),
	11: uint16(9),
	12: uint16(':'),
	13: uint16(90),
	14: uint16('<'),
	15: uint16(33),
	16: uint16('='),
	17: uint16(36),
	18: uint16('>'),
	19: uint16(31),
	20: uint16('@'),
	21: uint16(92),
	22: uint16('v'),
	23: uint16(8),
	24: uint16('{'),
	25: uint16(12),
	26: uint16('}'),
	27: uint16(13),
	28: uint16('D'),
	29: uint16(16),
	30: uint16('d'),
	31: uint16(16),
}

var ts_lex_modes = [170]TSLexerMode{
	0: {
		Fexternal_lex_state: uint16(1),
	},
	1: {
		Fexternal_lex_state: uint16(2),
	},
	2: {
		Fexternal_lex_state: uint16(3),
	},
	3: {
		Fexternal_lex_state: uint16(3),
	},
	4: {
		Fexternal_lex_state: uint16(3),
	},
	5: {
		Fexternal_lex_state: uint16(3),
	},
	6: {
		Fexternal_lex_state: uint16(2),
	},
	7: {
		Fexternal_lex_state: uint16(2),
	},
	8: {
		Fexternal_lex_state: uint16(2),
	},
	9: {
		Fexternal_lex_state: uint16(2),
	},
	10: {
		Fexternal_lex_state: uint16(3),
	},
	11: {
		Fexternal_lex_state: uint16(2),
	},
	12: {
		Fexternal_lex_state: uint16(2),
	},
	13: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(4),
	},
	14: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(4),
	},
	15: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(4),
	},
	16: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(4),
	},
	17: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(4),
	},
	18: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(4),
	},
	19: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(4),
	},
	20: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(4),
	},
	21: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(4),
	},
	22: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(4),
	},
	23: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(4),
	},
	24: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(5),
	},
	25: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(5),
	},
	26: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(5),
	},
	27: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(5),
	},
	28: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(5),
	},
	29: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(4),
	},
	30: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(5),
	},
	31: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(5),
	},
	32: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(4),
	},
	33: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(4),
	},
	34: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(4),
	},
	35: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(5),
	},
	36: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(5),
	},
	37: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(5),
	},
	38: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(4),
	},
	39: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(4),
	},
	40: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(5),
	},
	41: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(4),
	},
	42: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(4),
	},
	43: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(4),
	},
	44: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(5),
	},
	45: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(5),
	},
	46: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(5),
	},
	47: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(5),
	},
	48: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(4),
	},
	49: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(4),
	},
	50: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(4),
	},
	51: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(4),
	},
	52: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(5),
	},
	53: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(4),
	},
	54: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(5),
	},
	55: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(4),
	},
	56: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(4),
	},
	57: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(4),
	},
	58: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(5),
	},
	59: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(4),
	},
	60: {
		Fexternal_lex_state: uint16(3),
	},
	61: {
		Fexternal_lex_state: uint16(3),
	},
	62: {
		Fexternal_lex_state: uint16(2),
	},
	63: {
		Fexternal_lex_state: uint16(2),
	},
	64: {
		Fexternal_lex_state: uint16(2),
	},
	65: {
		Fexternal_lex_state: uint16(2),
	},
	66: {
		Fexternal_lex_state: uint16(2),
	},
	67: {
		Fexternal_lex_state: uint16(2),
	},
	68: {
		Fexternal_lex_state: uint16(2),
	},
	69: {
		Fexternal_lex_state: uint16(3),
	},
	70: {
		Fexternal_lex_state: uint16(2),
	},
	71: {
		Fexternal_lex_state: uint16(3),
	},
	72: {
		Fexternal_lex_state: uint16(3),
	},
	73: {
		Fexternal_lex_state: uint16(3),
	},
	74: {
		Fexternal_lex_state: uint16(3),
	},
	75: {
		Fexternal_lex_state: uint16(3),
	},
	76: {
		Fexternal_lex_state: uint16(3),
	},
	77: {
		Fexternal_lex_state: uint16(3),
	},
	78: {
		Fexternal_lex_state: uint16(3),
	},
	79: {
		Fexternal_lex_state: uint16(3),
	},
	80: {
		Fexternal_lex_state: uint16(3),
	},
	81: {
		Fexternal_lex_state: uint16(3),
	},
	82: {
		Fexternal_lex_state: uint16(3),
	},
	83: {
		Fexternal_lex_state: uint16(3),
	},
	84: {
		Fexternal_lex_state: uint16(3),
	},
	85: {
		Fexternal_lex_state: uint16(3),
	},
	86: {
		Fexternal_lex_state: uint16(2),
	},
	87: {
		Fexternal_lex_state: uint16(3),
	},
	88: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(5),
	},
	89: {
		Fexternal_lex_state: uint16(2),
	},
	90: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(5),
	},
	91: {
		Fexternal_lex_state: uint16(2),
	},
	92: {
		Fexternal_lex_state: uint16(2),
	},
	93: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(5),
	},
	94: {
		Fexternal_lex_state: uint16(2),
	},
	95: {
		Fexternal_lex_state: uint16(2),
	},
	96: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(5),
	},
	97: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(5),
	},
	98: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(5),
	},
	99: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(5),
	},
	100: {
		Fexternal_lex_state: uint16(2),
	},
	101: {
		Fexternal_lex_state: uint16(2),
	},
	102: {
		Fexternal_lex_state: uint16(2),
	},
	103: {
		Fexternal_lex_state: uint16(2),
	},
	104: {
		Fexternal_lex_state: uint16(2),
	},
	105: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(5),
	},
	106: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(5),
	},
	107: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(5),
	},
	108: {
		Fexternal_lex_state: uint16(6),
	},
	109: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(5),
	},
	110: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(5),
	},
	111: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(5),
	},
	112: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(5),
	},
	113: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(5),
	},
	114: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(5),
	},
	115: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(5),
	},
	116: {
		Fexternal_lex_state: uint16(6),
	},
	117: {
		Fexternal_lex_state: uint16(7),
	},
	118: {
		Fexternal_lex_state: uint16(7),
	},
	119: {
		Flex_state:          uint16(10),
		Fexternal_lex_state: uint16(5),
	},
	120: {
		Fexternal_lex_state: uint16(7),
	},
	121: {
		Fexternal_lex_state: uint16(7),
	},
	122: {
		Flex_state:          uint16(10),
		Fexternal_lex_state: uint16(5),
	},
	123: {
		Flex_state:          uint16(10),
		Fexternal_lex_state: uint16(5),
	},
	124: {
		Flex_state:          uint16(10),
		Fexternal_lex_state: uint16(5),
	},
	125: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(5),
	},
	126: {
		Fexternal_lex_state: uint16(8),
	},
	127: {
		Fexternal_lex_state: uint16(7),
	},
	128: {
		Fexternal_lex_state: uint16(7),
	},
	129: {
		Flex_state:          uint16(7),
		Fexternal_lex_state: uint16(5),
	},
	130: {
		Fexternal_lex_state: uint16(9),
	},
	131: {
		Fexternal_lex_state: uint16(9),
	},
	132: {
		Fexternal_lex_state: uint16(7),
	},
	133: {
		Fexternal_lex_state: uint16(8),
	},
	134: {
		Flex_state:          uint16(11),
		Fexternal_lex_state: uint16(5),
	},
	135: {
		Fexternal_lex_state: uint16(5),
	},
	136: {
		Fexternal_lex_state: uint16(5),
	},
	137: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(5),
	},
	138: {
		Fexternal_lex_state: uint16(5),
	},
	139: {
		Flex_state:          uint16(11),
		Fexternal_lex_state: uint16(5),
	},
	140: {
		Flex_state:          uint16(7),
		Fexternal_lex_state: uint16(5),
	},
	141: {
		Fexternal_lex_state: uint16(7),
	},
	142: {
		Fexternal_lex_state: uint16(8),
	},
	143: {
		Fexternal_lex_state: uint16(8),
	},
	144: {
		Fexternal_lex_state: uint16(5),
	},
	145: {
		Fexternal_lex_state: uint16(5),
	},
	146: {
		Fexternal_lex_state: uint16(5),
	},
	147: {
		Fexternal_lex_state: uint16(5),
	},
	148: {
		Flex_state:          uint16(25),
		Fexternal_lex_state: uint16(5),
	},
	149: {
		Fexternal_lex_state: uint16(5),
	},
	150: {
		Fexternal_lex_state: uint16(5),
	},
	151: {
		Fexternal_lex_state: uint16(5),
	},
	152: {
		Flex_state:          uint16(22),
		Fexternal_lex_state: uint16(5),
	},
	153: {
		Fexternal_lex_state: uint16(10),
	},
	154: {
		Flex_state:          uint16(25),
		Fexternal_lex_state: uint16(5),
	},
	155: {
		Fexternal_lex_state: uint16(5),
	},
	156: {
		Fexternal_lex_state: uint16(5),
	},
	157: {
		Fexternal_lex_state: uint16(5),
	},
	158: {
		Fexternal_lex_state: uint16(5),
	},
	159: {
		Fexternal_lex_state: uint16(5),
	},
	160: {
		Fexternal_lex_state: uint16(5),
	},
	161: {
		Fexternal_lex_state: uint16(5),
	},
	162: {
		Fexternal_lex_state: uint16(11),
	},
	163: {
		Fexternal_lex_state: uint16(5),
	},
	164: {
		Fexternal_lex_state: uint16(5),
	},
	165: {
		Fexternal_lex_state: uint16(10),
	},
	166: {
		Flex_state:          uint16(22),
		Fexternal_lex_state: uint16(5),
	},
	167: {
		Fexternal_lex_state: uint16(5),
	},
	168: {
		Fexternal_lex_state: uint16(11),
	},
	169: {
		Fexternal_lex_state: uint16(5),
	},
}

var ts_parse_table = [2][67]uint16_t{
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
		16: uint16(1),
		17: uint16(1),
		18: uint16(1),
		19: uint16(1),
		20: uint16(1),
		21: uint16(1),
		22: uint16(1),
		23: uint16(1),
		25: uint16(1),
		26: uint16(1),
		28: uint16(1),
		30: uint16(1),
		31: uint16(1),
		32: uint16(1),
		33: uint16(1),
		34: uint16(1),
		35: uint16(1),
		36: uint16(1),
		37: uint16(3),
		38: uint16(1),
		39: uint16(1),
		40: uint16(1),
	},
	1: {
		0:  uint16(5),
		1:  uint16(7),
		5:  uint16(9),
		7:  uint16(11),
		11: uint16(13),
		16: uint16(15),
		37: uint16(3),
		39: uint16(17),
		41: uint16(151),
		42: uint16(7),
		43: uint16(7),
		44: uint16(7),
		45: uint16(7),
		46: uint16(7),
		47: uint16(2),
		48: uint16(117),
		49: uint16(118),
		50: uint16(68),
		52: uint16(7),
		55: uint16(7),
		56: uint16(7),
		57: uint16(6),
		58: uint16(7),
		63: uint16(7),
	},
}

var ts_small_parse_table = [3024]uint16_t{
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
	15:   uint16(sym_entity),
	16:   uint16(27),
	17:   uint16(1),
	18:   uint16(anon_sym_LBRACE_LBRACE),
	19:   uint16(29),
	20:   uint16(1),
	21:   uint16(sym__implicit_end_tag),
	22:   uint16(31),
	23:   uint16(1),
	24:   uint16(sym__text_fragment),
	25:   uint16(4),
	26:   uint16(1),
	27:   uint16(sym_start_tag),
	28:   uint16(11),
	29:   uint16(1),
	30:   uint16(sym_template_start_tag),
	31:   uint16(63),
	32:   uint16(1),
	33:   uint16(sym_end_tag),
	34:   uint16(72),
	35:   uint16(1),
	36:   uint16(sym_self_closing_tag),
	37:   uint16(120),
	38:   uint16(1),
	39:   uint16(sym_script_start_tag),
	40:   uint16(121),
	41:   uint16(1),
	42:   uint16(sym_style_start_tag),
	43:   uint16(3),
	44:   uint16(10),
	45:   uint16(sym_doctype),
	46:   uint16(sym__node),
	47:   uint16(sym_element),
	48:   uint16(sym_script_element),
	49:   uint16(sym_style_element),
	50:   uint16(sym_erroneous_end_tag),
	51:   uint16(sym_text),
	52:   uint16(sym_template_element),
	53:   uint16(sym_interpolation),
	54:   uint16(aux_sym_document_repeat1),
	55:   uint16(15),
	56:   uint16(3),
	57:   uint16(1),
	58:   uint16(sym_comment),
	59:   uint16(19),
	60:   uint16(1),
	61:   uint16(anon_sym_LT_BANG),
	62:   uint16(21),
	63:   uint16(1),
	64:   uint16(anon_sym_LT),
	65:   uint16(23),
	66:   uint16(1),
	67:   uint16(anon_sym_LT_SLASH),
	68:   uint16(27),
	69:   uint16(1),
	70:   uint16(anon_sym_LBRACE_LBRACE),
	71:   uint16(31),
	72:   uint16(1),
	73:   uint16(sym__text_fragment),
	74:   uint16(33),
	75:   uint16(1),
	76:   uint16(sym_entity),
	77:   uint16(35),
	78:   uint16(1),
	79:   uint16(sym__implicit_end_tag),
	80:   uint16(4),
	81:   uint16(1),
	82:   uint16(sym_start_tag),
	83:   uint16(11),
	84:   uint16(1),
	85:   uint16(sym_template_start_tag),
	86:   uint16(72),
	87:   uint16(1),
	88:   uint16(sym_self_closing_tag),
	89:   uint16(92),
	90:   uint16(1),
	91:   uint16(sym_end_tag),
	92:   uint16(120),
	93:   uint16(1),
	94:   uint16(sym_script_start_tag),
	95:   uint16(121),
	96:   uint16(1),
	97:   uint16(sym_style_start_tag),
	98:   uint16(10),
	99:   uint16(10),
	100:  uint16(sym_doctype),
	101:  uint16(sym__node),
	102:  uint16(sym_element),
	103:  uint16(sym_script_element),
	104:  uint16(sym_style_element),
	105:  uint16(sym_erroneous_end_tag),
	106:  uint16(sym_text),
	107:  uint16(sym_template_element),
	108:  uint16(sym_interpolation),
	109:  uint16(aux_sym_document_repeat1),
	110:  uint16(15),
	111:  uint16(3),
	112:  uint16(1),
	113:  uint16(sym_comment),
	114:  uint16(19),
	115:  uint16(1),
	116:  uint16(anon_sym_LT_BANG),
	117:  uint16(21),
	118:  uint16(1),
	119:  uint16(anon_sym_LT),
	120:  uint16(27),
	121:  uint16(1),
	122:  uint16(anon_sym_LBRACE_LBRACE),
	123:  uint16(31),
	124:  uint16(1),
	125:  uint16(sym__text_fragment),
	126:  uint16(37),
	127:  uint16(1),
	128:  uint16(anon_sym_LT_SLASH),
	129:  uint16(39),
	130:  uint16(1),
	131:  uint16(sym_entity),
	132:  uint16(41),
	133:  uint16(1),
	134:  uint16(sym__implicit_end_tag),
	135:  uint16(4),
	136:  uint16(1),
	137:  uint16(sym_start_tag),
	138:  uint16(11),
	139:  uint16(1),
	140:  uint16(sym_template_start_tag),
	141:  uint16(72),
	142:  uint16(1),
	143:  uint16(sym_self_closing_tag),
	144:  uint16(74),
	145:  uint16(1),
	146:  uint16(sym_end_tag),
	147:  uint16(120),
	148:  uint16(1),
	149:  uint16(sym_script_start_tag),
	150:  uint16(121),
	151:  uint16(1),
	152:  uint16(sym_style_start_tag),
	153:  uint16(5),
	154:  uint16(10),
	155:  uint16(sym_doctype),
	156:  uint16(sym__node),
	157:  uint16(sym_element),
	158:  uint16(sym_script_element),
	159:  uint16(sym_style_element),
	160:  uint16(sym_erroneous_end_tag),
	161:  uint16(sym_text),
	162:  uint16(sym_template_element),
	163:  uint16(sym_interpolation),
	164:  uint16(aux_sym_document_repeat1),
	165:  uint16(15),
	166:  uint16(3),
	167:  uint16(1),
	168:  uint16(sym_comment),
	169:  uint16(19),
	170:  uint16(1),
	171:  uint16(anon_sym_LT_BANG),
	172:  uint16(21),
	173:  uint16(1),
	174:  uint16(anon_sym_LT),
	175:  uint16(27),
	176:  uint16(1),
	177:  uint16(anon_sym_LBRACE_LBRACE),
	178:  uint16(31),
	179:  uint16(1),
	180:  uint16(sym__text_fragment),
	181:  uint16(33),
	182:  uint16(1),
	183:  uint16(sym_entity),
	184:  uint16(37),
	185:  uint16(1),
	186:  uint16(anon_sym_LT_SLASH),
	187:  uint16(43),
	188:  uint16(1),
	189:  uint16(sym__implicit_end_tag),
	190:  uint16(4),
	191:  uint16(1),
	192:  uint16(sym_start_tag),
	193:  uint16(11),
	194:  uint16(1),
	195:  uint16(sym_template_start_tag),
	196:  uint16(72),
	197:  uint16(1),
	198:  uint16(sym_self_closing_tag),
	199:  uint16(80),
	200:  uint16(1),
	201:  uint16(sym_end_tag),
	202:  uint16(120),
	203:  uint16(1),
	204:  uint16(sym_script_start_tag),
	205:  uint16(121),
	206:  uint16(1),
	207:  uint16(sym_style_start_tag),
	208:  uint16(10),
	209:  uint16(10),
	210:  uint16(sym_doctype),
	211:  uint16(sym__node),
	212:  uint16(sym_element),
	213:  uint16(sym_script_element),
	214:  uint16(sym_style_element),
	215:  uint16(sym_erroneous_end_tag),
	216:  uint16(sym_text),
	217:  uint16(sym_template_element),
	218:  uint16(sym_interpolation),
	219:  uint16(aux_sym_document_repeat1),
	220:  uint16(14),
	221:  uint16(3),
	222:  uint16(1),
	223:  uint16(sym_comment),
	224:  uint16(7),
	225:  uint16(1),
	226:  uint16(anon_sym_LT_BANG),
	227:  uint16(9),
	228:  uint16(1),
	229:  uint16(anon_sym_LT),
	230:  uint16(15),
	231:  uint16(1),
	232:  uint16(anon_sym_LBRACE_LBRACE),
	233:  uint16(17),
	234:  uint16(1),
	235:  uint16(sym__text_fragment),
	236:  uint16(45),
	237:  uint16(1),
	238:  uint16(anon_sym_LT_SLASH),
	239:  uint16(47),
	240:  uint16(1),
	241:  uint16(sym_entity),
	242:  uint16(2),
	243:  uint16(1),
	244:  uint16(sym_start_tag),
	245:  uint16(6),
	246:  uint16(1),
	247:  uint16(sym_template_start_tag),
	248:  uint16(67),
	249:  uint16(1),
	250:  uint16(sym_end_tag),
	251:  uint16(68),
	252:  uint16(1),
	253:  uint16(sym_self_closing_tag),
	254:  uint16(117),
	255:  uint16(1),
	256:  uint16(sym_script_start_tag),
	257:  uint16(118),
	258:  uint16(1),
	259:  uint16(sym_style_start_tag),
	260:  uint16(8),
	261:  uint16(10),
	262:  uint16(sym_doctype),
	263:  uint16(sym__node),
	264:  uint16(sym_element),
	265:  uint16(sym_script_element),
	266:  uint16(sym_style_element),
	267:  uint16(sym_erroneous_end_tag),
	268:  uint16(sym_text),
	269:  uint16(sym_template_element),
	270:  uint16(sym_interpolation),
	271:  uint16(aux_sym_document_repeat1),
	272:  uint16(14),
	273:  uint16(3),
	274:  uint16(1),
	275:  uint16(sym_comment),
	276:  uint16(7),
	277:  uint16(1),
	278:  uint16(anon_sym_LT_BANG),
	279:  uint16(9),
	280:  uint16(1),
	281:  uint16(anon_sym_LT),
	282:  uint16(11),
	283:  uint16(1),
	284:  uint16(anon_sym_LT_SLASH),
	285:  uint16(15),
	286:  uint16(1),
	287:  uint16(anon_sym_LBRACE_LBRACE),
	288:  uint16(17),
	289:  uint16(1),
	290:  uint16(sym__text_fragment),
	291:  uint16(49),
	292:  uint16(1),
	294:  uint16(51),
	295:  uint16(1),
	296:  uint16(sym_entity),
	297:  uint16(2),
	298:  uint16(1),
	299:  uint16(sym_start_tag),
	300:  uint16(6),
	301:  uint16(1),
	302:  uint16(sym_template_start_tag),
	303:  uint16(68),
	304:  uint16(1),
	305:  uint16(sym_self_closing_tag),
	306:  uint16(117),
	307:  uint16(1),
	308:  uint16(sym_script_start_tag),
	309:  uint16(118),
	310:  uint16(1),
	311:  uint16(sym_style_start_tag),
	312:  uint16(9),
	313:  uint16(10),
	314:  uint16(sym_doctype),
	315:  uint16(sym__node),
	316:  uint16(sym_element),
	317:  uint16(sym_script_element),
	318:  uint16(sym_style_element),
	319:  uint16(sym_erroneous_end_tag),
	320:  uint16(sym_text),
	321:  uint16(sym_template_element),
	322:  uint16(sym_interpolation),
	323:  uint16(aux_sym_document_repeat1),
	324:  uint16(14),
	325:  uint16(3),
	326:  uint16(1),
	327:  uint16(sym_comment),
	328:  uint16(7),
	329:  uint16(1),
	330:  uint16(anon_sym_LT_BANG),
	331:  uint16(9),
	332:  uint16(1),
	333:  uint16(anon_sym_LT),
	334:  uint16(15),
	335:  uint16(1),
	336:  uint16(anon_sym_LBRACE_LBRACE),
	337:  uint16(17),
	338:  uint16(1),
	339:  uint16(sym__text_fragment),
	340:  uint16(45),
	341:  uint16(1),
	342:  uint16(anon_sym_LT_SLASH),
	343:  uint16(51),
	344:  uint16(1),
	345:  uint16(sym_entity),
	346:  uint16(2),
	347:  uint16(1),
	348:  uint16(sym_start_tag),
	349:  uint16(6),
	350:  uint16(1),
	351:  uint16(sym_template_start_tag),
	352:  uint16(68),
	353:  uint16(1),
	354:  uint16(sym_self_closing_tag),
	355:  uint16(100),
	356:  uint16(1),
	357:  uint16(sym_end_tag),
	358:  uint16(117),
	359:  uint16(1),
	360:  uint16(sym_script_start_tag),
	361:  uint16(118),
	362:  uint16(1),
	363:  uint16(sym_style_start_tag),
	364:  uint16(9),
	365:  uint16(10),
	366:  uint16(sym_doctype),
	367:  uint16(sym__node),
	368:  uint16(sym_element),
	369:  uint16(sym_script_element),
	370:  uint16(sym_style_element),
	371:  uint16(sym_erroneous_end_tag),
	372:  uint16(sym_text),
	373:  uint16(sym_template_element),
	374:  uint16(sym_interpolation),
	375:  uint16(aux_sym_document_repeat1),
	376:  uint16(14),
	377:  uint16(3),
	378:  uint16(1),
	379:  uint16(sym_comment),
	380:  uint16(53),
	381:  uint16(1),
	383:  uint16(55),
	384:  uint16(1),
	385:  uint16(anon_sym_LT_BANG),
	386:  uint16(58),
	387:  uint16(1),
	388:  uint16(anon_sym_LT),
	389:  uint16(61),
	390:  uint16(1),
	391:  uint16(anon_sym_LT_SLASH),
	392:  uint16(64),
	393:  uint16(1),
	394:  uint16(sym_entity),
	395:  uint16(67),
	396:  uint16(1),
	397:  uint16(anon_sym_LBRACE_LBRACE),
	398:  uint16(70),
	399:  uint16(1),
	400:  uint16(sym__text_fragment),
	401:  uint16(2),
	402:  uint16(1),
	403:  uint16(sym_start_tag),
	404:  uint16(6),
	405:  uint16(1),
	406:  uint16(sym_template_start_tag),
	407:  uint16(68),
	408:  uint16(1),
	409:  uint16(sym_self_closing_tag),
	410:  uint16(117),
	411:  uint16(1),
	412:  uint16(sym_script_start_tag),
	413:  uint16(118),
	414:  uint16(1),
	415:  uint16(sym_style_start_tag),
	416:  uint16(9),
	417:  uint16(10),
	418:  uint16(sym_doctype),
	419:  uint16(sym__node),
	420:  uint16(sym_element),
	421:  uint16(sym_script_element),
	422:  uint16(sym_style_element),
	423:  uint16(sym_erroneous_end_tag),
	424:  uint16(sym_text),
	425:  uint16(sym_template_element),
	426:  uint16(sym_interpolation),
	427:  uint16(aux_sym_document_repeat1),
	428:  uint16(14),
	429:  uint16(3),
	430:  uint16(1),
	431:  uint16(sym_comment),
	432:  uint16(53),
	433:  uint16(1),
	434:  uint16(sym__implicit_end_tag),
	435:  uint16(73),
	436:  uint16(1),
	437:  uint16(anon_sym_LT_BANG),
	438:  uint16(76),
	439:  uint16(1),
	440:  uint16(anon_sym_LT),
	441:  uint16(79),
	442:  uint16(1),
	443:  uint16(anon_sym_LT_SLASH),
	444:  uint16(82),
	445:  uint16(1),
	446:  uint16(sym_entity),
	447:  uint16(85),
	448:  uint16(1),
	449:  uint16(anon_sym_LBRACE_LBRACE),
	450:  uint16(88),
	451:  uint16(1),
	452:  uint16(sym__text_fragment),
	453:  uint16(4),
	454:  uint16(1),
	455:  uint16(sym_start_tag),
	456:  uint16(11),
	457:  uint16(1),
	458:  uint16(sym_template_start_tag),
	459:  uint16(72),
	460:  uint16(1),
	461:  uint16(sym_self_closing_tag),
	462:  uint16(120),
	463:  uint16(1),
	464:  uint16(sym_script_start_tag),
	465:  uint16(121),
	466:  uint16(1),
	467:  uint16(sym_style_start_tag),
	468:  uint16(10),
	469:  uint16(10),
	470:  uint16(sym_doctype),
	471:  uint16(sym__node),
	472:  uint16(sym_element),
	473:  uint16(sym_script_element),
	474:  uint16(sym_style_element),
	475:  uint16(sym_erroneous_end_tag),
	476:  uint16(sym_text),
	477:  uint16(sym_template_element),
	478:  uint16(sym_interpolation),
	479:  uint16(aux_sym_document_repeat1),
	480:  uint16(14),
	481:  uint16(3),
	482:  uint16(1),
	483:  uint16(sym_comment),
	484:  uint16(7),
	485:  uint16(1),
	486:  uint16(anon_sym_LT_BANG),
	487:  uint16(9),
	488:  uint16(1),
	489:  uint16(anon_sym_LT),
	490:  uint16(15),
	491:  uint16(1),
	492:  uint16(anon_sym_LBRACE_LBRACE),
	493:  uint16(17),
	494:  uint16(1),
	495:  uint16(sym__text_fragment),
	496:  uint16(91),
	497:  uint16(1),
	498:  uint16(anon_sym_LT_SLASH),
	499:  uint16(93),
	500:  uint16(1),
	501:  uint16(sym_entity),
	502:  uint16(2),
	503:  uint16(1),
	504:  uint16(sym_start_tag),
	505:  uint16(6),
	506:  uint16(1),
	507:  uint16(sym_template_start_tag),
	508:  uint16(60),
	509:  uint16(1),
	510:  uint16(sym_end_tag),
	511:  uint16(68),
	512:  uint16(1),
	513:  uint16(sym_self_closing_tag),
	514:  uint16(117),
	515:  uint16(1),
	516:  uint16(sym_script_start_tag),
	517:  uint16(118),
	518:  uint16(1),
	519:  uint16(sym_style_start_tag),
	520:  uint16(12),
	521:  uint16(10),
	522:  uint16(sym_doctype),
	523:  uint16(sym__node),
	524:  uint16(sym_element),
	525:  uint16(sym_script_element),
	526:  uint16(sym_style_element),
	527:  uint16(sym_erroneous_end_tag),
	528:  uint16(sym_text),
	529:  uint16(sym_template_element),
	530:  uint16(sym_interpolation),
	531:  uint16(aux_sym_document_repeat1),
	532:  uint16(14),
	533:  uint16(3),
	534:  uint16(1),
	535:  uint16(sym_comment),
	536:  uint16(7),
	537:  uint16(1),
	538:  uint16(anon_sym_LT_BANG),
	539:  uint16(9),
	540:  uint16(1),
	541:  uint16(anon_sym_LT),
	542:  uint16(15),
	543:  uint16(1),
	544:  uint16(anon_sym_LBRACE_LBRACE),
	545:  uint16(17),
	546:  uint16(1),
	547:  uint16(sym__text_fragment),
	548:  uint16(51),
	549:  uint16(1),
	550:  uint16(sym_entity),
	551:  uint16(91),
	552:  uint16(1),
	553:  uint16(anon_sym_LT_SLASH),
	554:  uint16(2),
	555:  uint16(1),
	556:  uint16(sym_start_tag),
	557:  uint16(6),
	558:  uint16(1),
	559:  uint16(sym_template_start_tag),
	560:  uint16(68),
	561:  uint16(1),
	562:  uint16(sym_self_closing_tag),
	563:  uint16(83),
	564:  uint16(1),
	565:  uint16(sym_end_tag),
	566:  uint16(117),
	567:  uint16(1),
	568:  uint16(sym_script_start_tag),
	569:  uint16(118),
	570:  uint16(1),
	571:  uint16(sym_style_start_tag),
	572:  uint16(9),
	573:  uint16(10),
	574:  uint16(sym_doctype),
	575:  uint16(sym__node),
	576:  uint16(sym_element),
	577:  uint16(sym_script_element),
	578:  uint16(sym_style_element),
	579:  uint16(sym_erroneous_end_tag),
	580:  uint16(sym_text),
	581:  uint16(sym_template_element),
	582:  uint16(sym_interpolation),
	583:  uint16(aux_sym_document_repeat1),
	584:  uint16(10),
	585:  uint16(3),
	586:  uint16(1),
	587:  uint16(sym_comment),
	588:  uint16(97),
	589:  uint16(1),
	590:  uint16(anon_sym_EQ),
	591:  uint16(99),
	592:  uint16(1),
	593:  uint16(sym_attribute_name),
	594:  uint16(103),
	595:  uint16(1),
	596:  uint16(anon_sym_DOT),
	597:  uint16(105),
	598:  uint16(1),
	599:  uint16(anon_sym_DOT2),
	600:  uint16(33),
	601:  uint16(1),
	602:  uint16(aux_sym_directive_attribute_repeat1),
	603:  uint16(34),
	604:  uint16(1),
	605:  uint16(aux_sym_directive_modifiers_repeat1),
	606:  uint16(49),
	607:  uint16(1),
	608:  uint16(sym_directive_modifiers),
	609:  uint16(95),
	610:  uint16(3),
	611:  uint16(anon_sym_GT),
	612:  uint16(anon_sym_SLASH_GT),
	613:  uint16(sym_directive_name),
	614:  uint16(101),
	615:  uint16(3),
	616:  uint16(anon_sym_COLON2),
	617:  uint16(anon_sym_AT),
	618:  uint16(anon_sym_POUND),
	619:  uint16(8),
	620:  uint16(3),
	621:  uint16(1),
	622:  uint16(sym_comment),
	623:  uint16(107),
	624:  uint16(1),
	625:  uint16(anon_sym_GT),
	626:  uint16(109),
	627:  uint16(1),
	628:  uint16(anon_sym_SLASH_GT),
	629:  uint16(111),
	630:  uint16(1),
	631:  uint16(sym_attribute_name),
	632:  uint16(113),
	633:  uint16(1),
	634:  uint16(sym_directive_name),
	635:  uint16(13),
	636:  uint16(1),
	637:  uint16(aux_sym_directive_attribute_repeat1),
	638:  uint16(101),
	639:  uint16(4),
	640:  uint16(anon_sym_COLON2),
	641:  uint16(anon_sym_DOT),
	642:  uint16(anon_sym_AT),
	643:  uint16(anon_sym_POUND),
	644:  uint16(21),
	645:  uint16(4),
	646:  uint16(sym_attribute),
	647:  uint16(sym__attribute),
	648:  uint16(sym_directive_attribute),
	649:  uint16(aux_sym_start_tag_repeat1),
	650:  uint16(8),
	651:  uint16(3),
	652:  uint16(1),
	653:  uint16(sym_comment),
	654:  uint16(111),
	655:  uint16(1),
	656:  uint16(sym_attribute_name),
	657:  uint16(113),
	658:  uint16(1),
	659:  uint16(sym_directive_name),
	660:  uint16(115),
	661:  uint16(1),
	662:  uint16(anon_sym_GT),
	663:  uint16(117),
	664:  uint16(1),
	665:  uint16(anon_sym_SLASH_GT),
	666:  uint16(13),
	667:  uint16(1),
	668:  uint16(aux_sym_directive_attribute_repeat1),
	669:  uint16(101),
	670:  uint16(4),
	671:  uint16(anon_sym_COLON2),
	672:  uint16(anon_sym_DOT),
	673:  uint16(anon_sym_AT),
	674:  uint16(anon_sym_POUND),
	675:  uint16(18),
	676:  uint16(4),
	677:  uint16(sym_attribute),
	678:  uint16(sym__attribute),
	679:  uint16(sym_directive_attribute),
	680:  uint16(aux_sym_start_tag_repeat1),
	681:  uint16(8),
	682:  uint16(3),
	683:  uint16(1),
	684:  uint16(sym_comment),
	685:  uint16(111),
	686:  uint16(1),
	687:  uint16(sym_attribute_name),
	688:  uint16(113),
	689:  uint16(1),
	690:  uint16(sym_directive_name),
	691:  uint16(117),
	692:  uint16(1),
	693:  uint16(anon_sym_SLASH_GT),
	694:  uint16(119),
	695:  uint16(1),
	696:  uint16(anon_sym_GT),
	697:  uint16(13),
	698:  uint16(1),
	699:  uint16(aux_sym_directive_attribute_repeat1),
	700:  uint16(101),
	701:  uint16(4),
	702:  uint16(anon_sym_COLON2),
	703:  uint16(anon_sym_DOT),
	704:  uint16(anon_sym_AT),
	705:  uint16(anon_sym_POUND),
	706:  uint16(14),
	707:  uint16(4),
	708:  uint16(sym_attribute),
	709:  uint16(sym__attribute),
	710:  uint16(sym_directive_attribute),
	711:  uint16(aux_sym_start_tag_repeat1),
	712:  uint16(8),
	713:  uint16(3),
	714:  uint16(1),
	715:  uint16(sym_comment),
	716:  uint16(97),
	717:  uint16(1),
	718:  uint16(anon_sym_EQ),
	719:  uint16(105),
	720:  uint16(1),
	721:  uint16(anon_sym_DOT2),
	722:  uint16(121),
	723:  uint16(1),
	724:  uint16(anon_sym_COLON),
	725:  uint16(34),
	726:  uint16(1),
	727:  uint16(aux_sym_directive_modifiers_repeat1),
	728:  uint16(49),
	729:  uint16(1),
	730:  uint16(sym_directive_modifiers),
	731:  uint16(99),
	732:  uint16(3),
	733:  uint16(sym_attribute_name),
	734:  uint16(anon_sym_COLON2),
	735:  uint16(anon_sym_DOT),
	736:  uint16(95),
	737:  uint16(5),
	738:  uint16(anon_sym_GT),
	739:  uint16(anon_sym_SLASH_GT),
	740:  uint16(anon_sym_AT),
	741:  uint16(anon_sym_POUND),
	742:  uint16(sym_directive_name),
	743:  uint16(8),
	744:  uint16(3),
	745:  uint16(1),
	746:  uint16(sym_comment),
	747:  uint16(109),
	748:  uint16(1),
	749:  uint16(anon_sym_SLASH_GT),
	750:  uint16(111),
	751:  uint16(1),
	752:  uint16(sym_attribute_name),
	753:  uint16(113),
	754:  uint16(1),
	755:  uint16(sym_directive_name),
	756:  uint16(123),
	757:  uint16(1),
	758:  uint16(anon_sym_GT),
	759:  uint16(13),
	760:  uint16(1),
	761:  uint16(aux_sym_directive_attribute_repeat1),
	762:  uint16(101),
	763:  uint16(4),
	764:  uint16(anon_sym_COLON2),
	765:  uint16(anon_sym_DOT),
	766:  uint16(anon_sym_AT),
	767:  uint16(anon_sym_POUND),
	768:  uint16(21),
	769:  uint16(4),
	770:  uint16(sym_attribute),
	771:  uint16(sym__attribute),
	772:  uint16(sym_directive_attribute),
	773:  uint16(aux_sym_start_tag_repeat1),
	774:  uint16(8),
	775:  uint16(3),
	776:  uint16(1),
	777:  uint16(sym_comment),
	778:  uint16(111),
	779:  uint16(1),
	780:  uint16(sym_attribute_name),
	781:  uint16(113),
	782:  uint16(1),
	783:  uint16(sym_directive_name),
	784:  uint16(119),
	785:  uint16(1),
	786:  uint16(anon_sym_GT),
	787:  uint16(125),
	788:  uint16(1),
	789:  uint16(anon_sym_SLASH_GT),
	790:  uint16(13),
	791:  uint16(1),
	792:  uint16(aux_sym_directive_attribute_repeat1),
	793:  uint16(101),
	794:  uint16(4),
	795:  uint16(anon_sym_COLON2),
	796:  uint16(anon_sym_DOT),
	797:  uint16(anon_sym_AT),
	798:  uint16(anon_sym_POUND),
	799:  uint16(22),
	800:  uint16(4),
	801:  uint16(sym_attribute),
	802:  uint16(sym__attribute),
	803:  uint16(sym_directive_attribute),
	804:  uint16(aux_sym_start_tag_repeat1),
	805:  uint16(8),
	806:  uint16(3),
	807:  uint16(1),
	808:  uint16(sym_comment),
	809:  uint16(111),
	810:  uint16(1),
	811:  uint16(sym_attribute_name),
	812:  uint16(113),
	813:  uint16(1),
	814:  uint16(sym_directive_name),
	815:  uint16(115),
	816:  uint16(1),
	817:  uint16(anon_sym_GT),
	818:  uint16(125),
	819:  uint16(1),
	820:  uint16(anon_sym_SLASH_GT),
	821:  uint16(13),
	822:  uint16(1),
	823:  uint16(aux_sym_directive_attribute_repeat1),
	824:  uint16(101),
	825:  uint16(4),
	826:  uint16(anon_sym_COLON2),
	827:  uint16(anon_sym_DOT),
	828:  uint16(anon_sym_AT),
	829:  uint16(anon_sym_POUND),
	830:  uint16(23),
	831:  uint16(4),
	832:  uint16(sym_attribute),
	833:  uint16(sym__attribute),
	834:  uint16(sym_directive_attribute),
	835:  uint16(aux_sym_start_tag_repeat1),
	836:  uint16(7),
	837:  uint16(3),
	838:  uint16(1),
	839:  uint16(sym_comment),
	840:  uint16(129),
	841:  uint16(1),
	842:  uint16(sym_attribute_name),
	843:  uint16(135),
	844:  uint16(1),
	845:  uint16(sym_directive_name),
	846:  uint16(13),
	847:  uint16(1),
	848:  uint16(aux_sym_directive_attribute_repeat1),
	849:  uint16(127),
	850:  uint16(2),
	851:  uint16(anon_sym_GT),
	852:  uint16(anon_sym_SLASH_GT),
	853:  uint16(132),
	854:  uint16(4),
	855:  uint16(anon_sym_COLON2),
	856:  uint16(anon_sym_DOT),
	857:  uint16(anon_sym_AT),
	858:  uint16(anon_sym_POUND),
	859:  uint16(21),
	860:  uint16(4),
	861:  uint16(sym_attribute),
	862:  uint16(sym__attribute),
	863:  uint16(sym_directive_attribute),
	864:  uint16(aux_sym_start_tag_repeat1),
	865:  uint16(8),
	866:  uint16(3),
	867:  uint16(1),
	868:  uint16(sym_comment),
	869:  uint16(107),
	870:  uint16(1),
	871:  uint16(anon_sym_GT),
	872:  uint16(111),
	873:  uint16(1),
	874:  uint16(sym_attribute_name),
	875:  uint16(113),
	876:  uint16(1),
	877:  uint16(sym_directive_name),
	878:  uint16(138),
	879:  uint16(1),
	880:  uint16(anon_sym_SLASH_GT),
	881:  uint16(13),
	882:  uint16(1),
	883:  uint16(aux_sym_directive_attribute_repeat1),
	884:  uint16(101),
	885:  uint16(4),
	886:  uint16(anon_sym_COLON2),
	887:  uint16(anon_sym_DOT),
	888:  uint16(anon_sym_AT),
	889:  uint16(anon_sym_POUND),
	890:  uint16(21),
	891:  uint16(4),
	892:  uint16(sym_attribute),
	893:  uint16(sym__attribute),
	894:  uint16(sym_directive_attribute),
	895:  uint16(aux_sym_start_tag_repeat1),
	896:  uint16(8),
	897:  uint16(3),
	898:  uint16(1),
	899:  uint16(sym_comment),
	900:  uint16(111),
	901:  uint16(1),
	902:  uint16(sym_attribute_name),
	903:  uint16(113),
	904:  uint16(1),
	905:  uint16(sym_directive_name),
	906:  uint16(123),
	907:  uint16(1),
	908:  uint16(anon_sym_GT),
	909:  uint16(138),
	910:  uint16(1),
	911:  uint16(anon_sym_SLASH_GT),
	912:  uint16(13),
	913:  uint16(1),
	914:  uint16(aux_sym_directive_attribute_repeat1),
	915:  uint16(101),
	916:  uint16(4),
	917:  uint16(anon_sym_COLON2),
	918:  uint16(anon_sym_DOT),
	919:  uint16(anon_sym_AT),
	920:  uint16(anon_sym_POUND),
	921:  uint16(21),
	922:  uint16(4),
	923:  uint16(sym_attribute),
	924:  uint16(sym__attribute),
	925:  uint16(sym_directive_attribute),
	926:  uint16(aux_sym_start_tag_repeat1),
	927:  uint16(7),
	928:  uint16(3),
	929:  uint16(1),
	930:  uint16(sym_comment),
	931:  uint16(140),
	932:  uint16(1),
	933:  uint16(anon_sym_GT),
	934:  uint16(142),
	935:  uint16(1),
	936:  uint16(sym_attribute_name),
	937:  uint16(146),
	938:  uint16(1),
	939:  uint16(sym_directive_name),
	940:  uint16(30),
	941:  uint16(1),
	942:  uint16(aux_sym_directive_attribute_repeat1),
	943:  uint16(144),
	944:  uint16(4),
	945:  uint16(anon_sym_COLON2),
	946:  uint16(anon_sym_DOT),
	947:  uint16(anon_sym_AT),
	948:  uint16(anon_sym_POUND),
	949:  uint16(26),
	950:  uint16(4),
	951:  uint16(sym_attribute),
	952:  uint16(sym__attribute),
	953:  uint16(sym_directive_attribute),
	954:  uint16(aux_sym_start_tag_repeat1),
	955:  uint16(7),
	956:  uint16(3),
	957:  uint16(1),
	958:  uint16(sym_comment),
	959:  uint16(142),
	960:  uint16(1),
	961:  uint16(sym_attribute_name),
	962:  uint16(146),
	963:  uint16(1),
	964:  uint16(sym_directive_name),
	965:  uint16(148),
	966:  uint16(1),
	967:  uint16(anon_sym_GT),
	968:  uint16(30),
	969:  uint16(1),
	970:  uint16(aux_sym_directive_attribute_repeat1),
	971:  uint16(144),
	972:  uint16(4),
	973:  uint16(anon_sym_COLON2),
	974:  uint16(anon_sym_DOT),
	975:  uint16(anon_sym_AT),
	976:  uint16(anon_sym_POUND),
	977:  uint16(27),
	978:  uint16(4),
	979:  uint16(sym_attribute),
	980:  uint16(sym__attribute),
	981:  uint16(sym_directive_attribute),
	982:  uint16(aux_sym_start_tag_repeat1),
	983:  uint16(7),
	984:  uint16(3),
	985:  uint16(1),
	986:  uint16(sym_comment),
	987:  uint16(142),
	988:  uint16(1),
	989:  uint16(sym_attribute_name),
	990:  uint16(146),
	991:  uint16(1),
	992:  uint16(sym_directive_name),
	993:  uint16(150),
	994:  uint16(1),
	995:  uint16(anon_sym_GT),
	996:  uint16(30),
	997:  uint16(1),
	998:  uint16(aux_sym_directive_attribute_repeat1),
	999:  uint16(144),
	1000: uint16(4),
	1001: uint16(anon_sym_COLON2),
	1002: uint16(anon_sym_DOT),
	1003: uint16(anon_sym_AT),
	1004: uint16(anon_sym_POUND),
	1005: uint16(28),
	1006: uint16(4),
	1007: uint16(sym_attribute),
	1008: uint16(sym__attribute),
	1009: uint16(sym_directive_attribute),
	1010: uint16(aux_sym_start_tag_repeat1),
	1011: uint16(7),
	1012: uint16(3),
	1013: uint16(1),
	1014: uint16(sym_comment),
	1015: uint16(142),
	1016: uint16(1),
	1017: uint16(sym_attribute_name),
	1018: uint16(146),
	1019: uint16(1),
	1020: uint16(sym_directive_name),
	1021: uint16(152),
	1022: uint16(1),
	1023: uint16(anon_sym_GT),
	1024: uint16(30),
	1025: uint16(1),
	1026: uint16(aux_sym_directive_attribute_repeat1),
	1027: uint16(144),
	1028: uint16(4),
	1029: uint16(anon_sym_COLON2),
	1030: uint16(anon_sym_DOT),
	1031: uint16(anon_sym_AT),
	1032: uint16(anon_sym_POUND),
	1033: uint16(28),
	1034: uint16(4),
	1035: uint16(sym_attribute),
	1036: uint16(sym__attribute),
	1037: uint16(sym_directive_attribute),
	1038: uint16(aux_sym_start_tag_repeat1),
	1039: uint16(7),
	1040: uint16(3),
	1041: uint16(1),
	1042: uint16(sym_comment),
	1043: uint16(127),
	1044: uint16(1),
	1045: uint16(anon_sym_GT),
	1046: uint16(154),
	1047: uint16(1),
	1048: uint16(sym_attribute_name),
	1049: uint16(160),
	1050: uint16(1),
	1051: uint16(sym_directive_name),
	1052: uint16(30),
	1053: uint16(1),
	1054: uint16(aux_sym_directive_attribute_repeat1),
	1055: uint16(157),
	1056: uint16(4),
	1057: uint16(anon_sym_COLON2),
	1058: uint16(anon_sym_DOT),
	1059: uint16(anon_sym_AT),
	1060: uint16(anon_sym_POUND),
	1061: uint16(28),
	1062: uint16(4),
	1063: uint16(sym_attribute),
	1064: uint16(sym__attribute),
	1065: uint16(sym_directive_attribute),
	1066: uint16(aux_sym_start_tag_repeat1),
	1067: uint16(7),
	1068: uint16(3),
	1069: uint16(1),
	1070: uint16(sym_comment),
	1071: uint16(105),
	1072: uint16(1),
	1073: uint16(anon_sym_DOT2),
	1074: uint16(165),
	1075: uint16(1),
	1076: uint16(anon_sym_EQ),
	1077: uint16(34),
	1078: uint16(1),
	1079: uint16(aux_sym_directive_modifiers_repeat1),
	1080: uint16(48),
	1081: uint16(1),
	1082: uint16(sym_directive_modifiers),
	1083: uint16(167),
	1084: uint16(2),
	1085: uint16(sym_attribute_name),
	1086: uint16(anon_sym_DOT),
	1087: uint16(163),
	1088: uint16(6),
	1089: uint16(anon_sym_GT),
	1090: uint16(anon_sym_SLASH_GT),
	1091: uint16(anon_sym_COLON2),
	1092: uint16(anon_sym_AT),
	1093: uint16(anon_sym_POUND),
	1094: uint16(sym_directive_name),
	1095: uint16(10),
	1096: uint16(3),
	1097: uint16(1),
	1098: uint16(sym_comment),
	1099: uint16(99),
	1100: uint16(1),
	1101: uint16(sym_attribute_name),
	1102: uint16(169),
	1103: uint16(1),
	1104: uint16(anon_sym_EQ),
	1105: uint16(171),
	1106: uint16(1),
	1107: uint16(anon_sym_DOT),
	1108: uint16(173),
	1109: uint16(1),
	1110: uint16(anon_sym_DOT2),
	1111: uint16(36),
	1112: uint16(1),
	1113: uint16(aux_sym_directive_attribute_repeat1),
	1114: uint16(37),
	1115: uint16(1),
	1116: uint16(aux_sym_directive_modifiers_repeat1),
	1117: uint16(52),
	1118: uint16(1),
	1119: uint16(sym_directive_modifiers),
	1120: uint16(95),
	1121: uint16(2),
	1122: uint16(anon_sym_GT),
	1123: uint16(sym_directive_name),
	1124: uint16(144),
	1125: uint16(3),
	1126: uint16(anon_sym_COLON2),
	1127: uint16(anon_sym_AT),
	1128: uint16(anon_sym_POUND),
	1129: uint16(8),
	1130: uint16(3),
	1131: uint16(1),
	1132: uint16(sym_comment),
	1133: uint16(169),
	1134: uint16(1),
	1135: uint16(anon_sym_EQ),
	1136: uint16(173),
	1137: uint16(1),
	1138: uint16(anon_sym_DOT2),
	1139: uint16(175),
	1140: uint16(1),
	1141: uint16(anon_sym_COLON),
	1142: uint16(37),
	1143: uint16(1),
	1144: uint16(aux_sym_directive_modifiers_repeat1),
	1145: uint16(52),
	1146: uint16(1),
	1147: uint16(sym_directive_modifiers),
	1148: uint16(99),
	1149: uint16(3),
	1150: uint16(sym_attribute_name),
	1151: uint16(anon_sym_COLON2),
	1152: uint16(anon_sym_DOT),
	1153: uint16(95),
	1154: uint16(4),
	1155: uint16(anon_sym_GT),
	1156: uint16(anon_sym_AT),
	1157: uint16(anon_sym_POUND),
	1158: uint16(sym_directive_name),
	1159: uint16(5),
	1160: uint16(3),
	1161: uint16(1),
	1162: uint16(sym_comment),
	1163: uint16(181),
	1164: uint16(1),
	1165: uint16(anon_sym_DOT2),
	1166: uint16(32),
	1167: uint16(1),
	1168: uint16(aux_sym_directive_modifiers_repeat1),
	1169: uint16(179),
	1170: uint16(2),
	1171: uint16(sym_attribute_name),
	1172: uint16(anon_sym_DOT),
	1173: uint16(177),
	1174: uint16(7),
	1175: uint16(anon_sym_GT),
	1176: uint16(anon_sym_SLASH_GT),
	1177: uint16(anon_sym_EQ),
	1178: uint16(anon_sym_COLON2),
	1179: uint16(anon_sym_AT),
	1180: uint16(anon_sym_POUND),
	1181: uint16(sym_directive_name),
	1182: uint16(6),
	1183: uint16(3),
	1184: uint16(1),
	1185: uint16(sym_comment),
	1186: uint16(186),
	1187: uint16(1),
	1188: uint16(sym_attribute_name),
	1189: uint16(191),
	1190: uint16(1),
	1191: uint16(anon_sym_DOT),
	1192: uint16(33),
	1193: uint16(1),
	1194: uint16(aux_sym_directive_attribute_repeat1),
	1195: uint16(188),
	1196: uint16(3),
	1197: uint16(anon_sym_COLON2),
	1198: uint16(anon_sym_AT),
	1199: uint16(anon_sym_POUND),
	1200: uint16(184),
	1201: uint16(5),
	1202: uint16(anon_sym_GT),
	1203: uint16(anon_sym_SLASH_GT),
	1204: uint16(anon_sym_EQ),
	1205: uint16(sym_directive_name),
	1206: uint16(anon_sym_DOT2),
	1207: uint16(5),
	1208: uint16(3),
	1209: uint16(1),
	1210: uint16(sym_comment),
	1211: uint16(105),
	1212: uint16(1),
	1213: uint16(anon_sym_DOT2),
	1214: uint16(32),
	1215: uint16(1),
	1216: uint16(aux_sym_directive_modifiers_repeat1),
	1217: uint16(196),
	1218: uint16(2),
	1219: uint16(sym_attribute_name),
	1220: uint16(anon_sym_DOT),
	1221: uint16(194),
	1222: uint16(7),
	1223: uint16(anon_sym_GT),
	1224: uint16(anon_sym_SLASH_GT),
	1225: uint16(anon_sym_EQ),
	1226: uint16(anon_sym_COLON2),
	1227: uint16(anon_sym_AT),
	1228: uint16(anon_sym_POUND),
	1229: uint16(sym_directive_name),
	1230: uint16(7),
	1231: uint16(3),
	1232: uint16(1),
	1233: uint16(sym_comment),
	1234: uint16(173),
	1235: uint16(1),
	1236: uint16(anon_sym_DOT2),
	1237: uint16(198),
	1238: uint16(1),
	1239: uint16(anon_sym_EQ),
	1240: uint16(37),
	1241: uint16(1),
	1242: uint16(aux_sym_directive_modifiers_repeat1),
	1243: uint16(58),
	1244: uint16(1),
	1245: uint16(sym_directive_modifiers),
	1246: uint16(167),
	1247: uint16(2),
	1248: uint16(sym_attribute_name),
	1249: uint16(anon_sym_DOT),
	1250: uint16(163),
	1251: uint16(5),
	1252: uint16(anon_sym_GT),
	1253: uint16(anon_sym_COLON2),
	1254: uint16(anon_sym_AT),
	1255: uint16(anon_sym_POUND),
	1256: uint16(sym_directive_name),
	1257: uint16(6),
	1258: uint16(3),
	1259: uint16(1),
	1260: uint16(sym_comment),
	1261: uint16(186),
	1262: uint16(1),
	1263: uint16(sym_attribute_name),
	1264: uint16(203),
	1265: uint16(1),
	1266: uint16(anon_sym_DOT),
	1267: uint16(36),
	1268: uint16(1),
	1269: uint16(aux_sym_directive_attribute_repeat1),
	1270: uint16(200),
	1271: uint16(3),
	1272: uint16(anon_sym_COLON2),
	1273: uint16(anon_sym_AT),
	1274: uint16(anon_sym_POUND),
	1275: uint16(184),
	1276: uint16(4),
	1277: uint16(anon_sym_GT),
	1278: uint16(anon_sym_EQ),
	1279: uint16(sym_directive_name),
	1280: uint16(anon_sym_DOT2),
	1281: uint16(5),
	1282: uint16(3),
	1283: uint16(1),
	1284: uint16(sym_comment),
	1285: uint16(173),
	1286: uint16(1),
	1287: uint16(anon_sym_DOT2),
	1288: uint16(40),
	1289: uint16(1),
	1290: uint16(aux_sym_directive_modifiers_repeat1),
	1291: uint16(196),
	1292: uint16(2),
	1293: uint16(sym_attribute_name),
	1294: uint16(anon_sym_DOT),
	1295: uint16(194),
	1296: uint16(6),
	1297: uint16(anon_sym_GT),
	1298: uint16(anon_sym_EQ),
	1299: uint16(anon_sym_COLON2),
	1300: uint16(anon_sym_AT),
	1301: uint16(anon_sym_POUND),
	1302: uint16(sym_directive_name),
	1303: uint16(3),
	1304: uint16(3),
	1305: uint16(1),
	1306: uint16(sym_comment),
	1307: uint16(208),
	1308: uint16(2),
	1309: uint16(sym_attribute_name),
	1310: uint16(anon_sym_DOT),
	1311: uint16(206),
	1312: uint16(8),
	1313: uint16(anon_sym_GT),
	1314: uint16(anon_sym_SLASH_GT),
	1315: uint16(anon_sym_EQ),
	1316: uint16(anon_sym_COLON2),
	1317: uint16(anon_sym_AT),
	1318: uint16(anon_sym_POUND),
	1319: uint16(sym_directive_name),
	1320: uint16(anon_sym_DOT2),
	1321: uint16(3),
	1322: uint16(3),
	1323: uint16(1),
	1324: uint16(sym_comment),
	1325: uint16(212),
	1326: uint16(2),
	1327: uint16(sym_attribute_name),
	1328: uint16(anon_sym_DOT),
	1329: uint16(210),
	1330: uint16(8),
	1331: uint16(anon_sym_GT),
	1332: uint16(anon_sym_SLASH_GT),
	1333: uint16(anon_sym_EQ),
	1334: uint16(anon_sym_COLON2),
	1335: uint16(anon_sym_AT),
	1336: uint16(anon_sym_POUND),
	1337: uint16(sym_directive_name),
	1338: uint16(anon_sym_DOT2),
	1339: uint16(5),
	1340: uint16(3),
	1341: uint16(1),
	1342: uint16(sym_comment),
	1343: uint16(214),
	1344: uint16(1),
	1345: uint16(anon_sym_DOT2),
	1346: uint16(40),
	1347: uint16(1),
	1348: uint16(aux_sym_directive_modifiers_repeat1),
	1349: uint16(179),
	1350: uint16(2),
	1351: uint16(sym_attribute_name),
	1352: uint16(anon_sym_DOT),
	1353: uint16(177),
	1354: uint16(6),
	1355: uint16(anon_sym_GT),
	1356: uint16(anon_sym_EQ),
	1357: uint16(anon_sym_COLON2),
	1358: uint16(anon_sym_AT),
	1359: uint16(anon_sym_POUND),
	1360: uint16(sym_directive_name),
	1361: uint16(3),
	1362: uint16(3),
	1363: uint16(1),
	1364: uint16(sym_comment),
	1365: uint16(179),
	1366: uint16(2),
	1367: uint16(sym_attribute_name),
	1368: uint16(anon_sym_DOT),
	1369: uint16(177),
	1370: uint16(8),
	1371: uint16(anon_sym_GT),
	1372: uint16(anon_sym_SLASH_GT),
	1373: uint16(anon_sym_EQ),
	1374: uint16(anon_sym_COLON2),
	1375: uint16(anon_sym_AT),
	1376: uint16(anon_sym_POUND),
	1377: uint16(sym_directive_name),
	1378: uint16(anon_sym_DOT2),
	1379: uint16(3),
	1380: uint16(3),
	1381: uint16(1),
	1382: uint16(sym_comment),
	1383: uint16(186),
	1384: uint16(2),
	1385: uint16(sym_attribute_name),
	1386: uint16(anon_sym_DOT),
	1387: uint16(184),
	1388: uint16(8),
	1389: uint16(anon_sym_GT),
	1390: uint16(anon_sym_SLASH_GT),
	1391: uint16(anon_sym_EQ),
	1392: uint16(anon_sym_COLON2),
	1393: uint16(anon_sym_AT),
	1394: uint16(anon_sym_POUND),
	1395: uint16(sym_directive_name),
	1396: uint16(anon_sym_DOT2),
	1397: uint16(4),
	1398: uint16(3),
	1399: uint16(1),
	1400: uint16(sym_comment),
	1401: uint16(219),
	1402: uint16(1),
	1403: uint16(anon_sym_EQ),
	1404: uint16(221),
	1405: uint16(1),
	1406: uint16(sym_attribute_name),
	1407: uint16(217),
	1408: uint16(7),
	1409: uint16(anon_sym_GT),
	1410: uint16(anon_sym_SLASH_GT),
	1411: uint16(anon_sym_COLON2),
	1412: uint16(anon_sym_DOT),
	1413: uint16(anon_sym_AT),
	1414: uint16(anon_sym_POUND),
	1415: uint16(sym_directive_name),
	1416: uint16(3),
	1417: uint16(3),
	1418: uint16(1),
	1419: uint16(sym_comment),
	1420: uint16(186),
	1421: uint16(2),
	1422: uint16(sym_attribute_name),
	1423: uint16(anon_sym_DOT),
	1424: uint16(184),
	1425: uint16(7),
	1426: uint16(anon_sym_GT),
	1427: uint16(anon_sym_EQ),
	1428: uint16(anon_sym_COLON2),
	1429: uint16(anon_sym_AT),
	1430: uint16(anon_sym_POUND),
	1431: uint16(sym_directive_name),
	1432: uint16(anon_sym_DOT2),
	1433: uint16(3),
	1434: uint16(3),
	1435: uint16(1),
	1436: uint16(sym_comment),
	1437: uint16(208),
	1438: uint16(2),
	1439: uint16(sym_attribute_name),
	1440: uint16(anon_sym_DOT),
	1441: uint16(206),
	1442: uint16(7),
	1443: uint16(anon_sym_GT),
	1444: uint16(anon_sym_EQ),
	1445: uint16(anon_sym_COLON2),
	1446: uint16(anon_sym_AT),
	1447: uint16(anon_sym_POUND),
	1448: uint16(sym_directive_name),
	1449: uint16(anon_sym_DOT2),
	1450: uint16(3),
	1451: uint16(3),
	1452: uint16(1),
	1453: uint16(sym_comment),
	1454: uint16(179),
	1455: uint16(2),
	1456: uint16(sym_attribute_name),
	1457: uint16(anon_sym_DOT),
	1458: uint16(177),
	1459: uint16(7),
	1460: uint16(anon_sym_GT),
	1461: uint16(anon_sym_EQ),
	1462: uint16(anon_sym_COLON2),
	1463: uint16(anon_sym_AT),
	1464: uint16(anon_sym_POUND),
	1465: uint16(sym_directive_name),
	1466: uint16(anon_sym_DOT2),
	1467: uint16(3),
	1468: uint16(3),
	1469: uint16(1),
	1470: uint16(sym_comment),
	1471: uint16(212),
	1472: uint16(2),
	1473: uint16(sym_attribute_name),
	1474: uint16(anon_sym_DOT),
	1475: uint16(210),
	1476: uint16(7),
	1477: uint16(anon_sym_GT),
	1478: uint16(anon_sym_EQ),
	1479: uint16(anon_sym_COLON2),
	1480: uint16(anon_sym_AT),
	1481: uint16(anon_sym_POUND),
	1482: uint16(sym_directive_name),
	1483: uint16(anon_sym_DOT2),
	1484: uint16(4),
	1485: uint16(3),
	1486: uint16(1),
	1487: uint16(sym_comment),
	1488: uint16(225),
	1489: uint16(1),
	1490: uint16(anon_sym_EQ),
	1491: uint16(227),
	1492: uint16(1),
	1493: uint16(sym_attribute_name),
	1494: uint16(223),
	1495: uint16(7),
	1496: uint16(anon_sym_GT),
	1497: uint16(anon_sym_SLASH_GT),
	1498: uint16(anon_sym_COLON2),
	1499: uint16(anon_sym_DOT),
	1500: uint16(anon_sym_AT),
	1501: uint16(anon_sym_POUND),
	1502: uint16(sym_directive_name),
	1503: uint16(4),
	1504: uint16(3),
	1505: uint16(1),
	1506: uint16(sym_comment),
	1507: uint16(231),
	1508: uint16(1),
	1509: uint16(anon_sym_EQ),
	1510: uint16(233),
	1511: uint16(1),
	1512: uint16(sym_attribute_name),
	1513: uint16(229),
	1514: uint16(7),
	1515: uint16(anon_sym_GT),
	1516: uint16(anon_sym_SLASH_GT),
	1517: uint16(anon_sym_COLON2),
	1518: uint16(anon_sym_DOT),
	1519: uint16(anon_sym_AT),
	1520: uint16(anon_sym_POUND),
	1521: uint16(sym_directive_name),
	1522: uint16(3),
	1523: uint16(3),
	1524: uint16(1),
	1525: uint16(sym_comment),
	1526: uint16(237),
	1527: uint16(1),
	1528: uint16(sym_attribute_name),
	1529: uint16(235),
	1530: uint16(7),
	1531: uint16(anon_sym_GT),
	1532: uint16(anon_sym_SLASH_GT),
	1533: uint16(anon_sym_COLON2),
	1534: uint16(anon_sym_DOT),
	1535: uint16(anon_sym_AT),
	1536: uint16(anon_sym_POUND),
	1537: uint16(sym_directive_name),
	1538: uint16(3),
	1539: uint16(3),
	1540: uint16(1),
	1541: uint16(sym_comment),
	1542: uint16(241),
	1543: uint16(1),
	1544: uint16(sym_attribute_name),
	1545: uint16(239),
	1546: uint16(7),
	1547: uint16(anon_sym_GT),
	1548: uint16(anon_sym_SLASH_GT),
	1549: uint16(anon_sym_COLON2),
	1550: uint16(anon_sym_DOT),
	1551: uint16(anon_sym_AT),
	1552: uint16(anon_sym_POUND),
	1553: uint16(sym_directive_name),
	1554: uint16(4),
	1555: uint16(3),
	1556: uint16(1),
	1557: uint16(sym_comment),
	1558: uint16(233),
	1559: uint16(1),
	1560: uint16(sym_attribute_name),
	1561: uint16(243),
	1562: uint16(1),
	1563: uint16(anon_sym_EQ),
	1564: uint16(229),
	1565: uint16(6),
	1566: uint16(anon_sym_GT),
	1567: uint16(anon_sym_COLON2),
	1568: uint16(anon_sym_DOT),
	1569: uint16(anon_sym_AT),
	1570: uint16(anon_sym_POUND),
	1571: uint16(sym_directive_name),
	1572: uint16(3),
	1573: uint16(3),
	1574: uint16(1),
	1575: uint16(sym_comment),
	1576: uint16(247),
	1577: uint16(1),
	1578: uint16(sym_attribute_name),
	1579: uint16(245),
	1580: uint16(7),
	1581: uint16(anon_sym_GT),
	1582: uint16(anon_sym_SLASH_GT),
	1583: uint16(anon_sym_COLON2),
	1584: uint16(anon_sym_DOT),
	1585: uint16(anon_sym_AT),
	1586: uint16(anon_sym_POUND),
	1587: uint16(sym_directive_name),
	1588: uint16(4),
	1589: uint16(3),
	1590: uint16(1),
	1591: uint16(sym_comment),
	1592: uint16(221),
	1593: uint16(1),
	1594: uint16(sym_attribute_name),
	1595: uint16(249),
	1596: uint16(1),
	1597: uint16(anon_sym_EQ),
	1598: uint16(217),
	1599: uint16(6),
	1600: uint16(anon_sym_GT),
	1601: uint16(anon_sym_COLON2),
	1602: uint16(anon_sym_DOT),
	1603: uint16(anon_sym_AT),
	1604: uint16(anon_sym_POUND),
	1605: uint16(sym_directive_name),
	1606: uint16(3),
	1607: uint16(3),
	1608: uint16(1),
	1609: uint16(sym_comment),
	1610: uint16(253),
	1611: uint16(1),
	1612: uint16(sym_attribute_name),
	1613: uint16(251),
	1614: uint16(7),
	1615: uint16(anon_sym_GT),
	1616: uint16(anon_sym_SLASH_GT),
	1617: uint16(anon_sym_COLON2),
	1618: uint16(anon_sym_DOT),
	1619: uint16(anon_sym_AT),
	1620: uint16(anon_sym_POUND),
	1621: uint16(sym_directive_name),
	1622: uint16(3),
	1623: uint16(3),
	1624: uint16(1),
	1625: uint16(sym_comment),
	1626: uint16(257),
	1627: uint16(1),
	1628: uint16(sym_attribute_name),
	1629: uint16(255),
	1630: uint16(7),
	1631: uint16(anon_sym_GT),
	1632: uint16(anon_sym_SLASH_GT),
	1633: uint16(anon_sym_COLON2),
	1634: uint16(anon_sym_DOT),
	1635: uint16(anon_sym_AT),
	1636: uint16(anon_sym_POUND),
	1637: uint16(sym_directive_name),
	1638: uint16(3),
	1639: uint16(3),
	1640: uint16(1),
	1641: uint16(sym_comment),
	1642: uint16(167),
	1643: uint16(1),
	1644: uint16(sym_attribute_name),
	1645: uint16(163),
	1646: uint16(7),
	1647: uint16(anon_sym_GT),
	1648: uint16(anon_sym_SLASH_GT),
	1649: uint16(anon_sym_COLON2),
	1650: uint16(anon_sym_DOT),
	1651: uint16(anon_sym_AT),
	1652: uint16(anon_sym_POUND),
	1653: uint16(sym_directive_name),
	1654: uint16(4),
	1655: uint16(3),
	1656: uint16(1),
	1657: uint16(sym_comment),
	1658: uint16(227),
	1659: uint16(1),
	1660: uint16(sym_attribute_name),
	1661: uint16(259),
	1662: uint16(1),
	1663: uint16(anon_sym_EQ),
	1664: uint16(223),
	1665: uint16(6),
	1666: uint16(anon_sym_GT),
	1667: uint16(anon_sym_COLON2),
	1668: uint16(anon_sym_DOT),
	1669: uint16(anon_sym_AT),
	1670: uint16(anon_sym_POUND),
	1671: uint16(sym_directive_name),
	1672: uint16(3),
	1673: uint16(3),
	1674: uint16(1),
	1675: uint16(sym_comment),
	1676: uint16(227),
	1677: uint16(1),
	1678: uint16(sym_attribute_name),
	1679: uint16(223),
	1680: uint16(7),
	1681: uint16(anon_sym_GT),
	1682: uint16(anon_sym_SLASH_GT),
	1683: uint16(anon_sym_COLON2),
	1684: uint16(anon_sym_DOT),
	1685: uint16(anon_sym_AT),
	1686: uint16(anon_sym_POUND),
	1687: uint16(sym_directive_name),
	1688: uint16(3),
	1689: uint16(3),
	1690: uint16(1),
	1691: uint16(sym_comment),
	1692: uint16(263),
	1693: uint16(1),
	1694: uint16(anon_sym_LT),
	1695: uint16(261),
	1696: uint16(6),
	1697: uint16(sym__implicit_end_tag),
	1698: uint16(sym__text_fragment),
	1699: uint16(anon_sym_LT_BANG),
	1700: uint16(anon_sym_LT_SLASH),
	1701: uint16(sym_entity),
	1702: uint16(anon_sym_LBRACE_LBRACE),
	1703: uint16(3),
	1704: uint16(3),
	1705: uint16(1),
	1706: uint16(sym_comment),
	1707: uint16(267),
	1708: uint16(1),
	1709: uint16(anon_sym_LT),
	1710: uint16(265),
	1711: uint16(6),
	1712: uint16(sym__implicit_end_tag),
	1713: uint16(sym__text_fragment),
	1714: uint16(anon_sym_LT_BANG),
	1715: uint16(anon_sym_LT_SLASH),
	1716: uint16(sym_entity),
	1717: uint16(anon_sym_LBRACE_LBRACE),
	1718: uint16(3),
	1719: uint16(3),
	1720: uint16(1),
	1721: uint16(sym_comment),
	1722: uint16(271),
	1723: uint16(1),
	1724: uint16(anon_sym_LT),
	1725: uint16(269),
	1726: uint16(6),
	1727: uint16(sym__text_fragment),
	1729: uint16(anon_sym_LT_BANG),
	1730: uint16(anon_sym_LT_SLASH),
	1731: uint16(sym_entity),
	1732: uint16(anon_sym_LBRACE_LBRACE),
	1733: uint16(3),
	1734: uint16(3),
	1735: uint16(1),
	1736: uint16(sym_comment),
	1737: uint16(275),
	1738: uint16(1),
	1739: uint16(anon_sym_LT),
	1740: uint16(273),
	1741: uint16(6),
	1742: uint16(sym__text_fragment),
	1744: uint16(anon_sym_LT_BANG),
	1745: uint16(anon_sym_LT_SLASH),
	1746: uint16(sym_entity),
	1747: uint16(anon_sym_LBRACE_LBRACE),
	1748: uint16(3),
	1749: uint16(3),
	1750: uint16(1),
	1751: uint16(sym_comment),
	1752: uint16(279),
	1753: uint16(1),
	1754: uint16(anon_sym_LT),
	1755: uint16(277),
	1756: uint16(6),
	1757: uint16(sym__text_fragment),
	1759: uint16(anon_sym_LT_BANG),
	1760: uint16(anon_sym_LT_SLASH),
	1761: uint16(sym_entity),
	1762: uint16(anon_sym_LBRACE_LBRACE),
	1763: uint16(3),
	1764: uint16(3),
	1765: uint16(1),
	1766: uint16(sym_comment),
	1767: uint16(283),
	1768: uint16(1),
	1769: uint16(anon_sym_LT),
	1770: uint16(281),
	1771: uint16(6),
	1772: uint16(sym__text_fragment),
	1774: uint16(anon_sym_LT_BANG),
	1775: uint16(anon_sym_LT_SLASH),
	1776: uint16(sym_entity),
	1777: uint16(anon_sym_LBRACE_LBRACE),
	1778: uint16(3),
	1779: uint16(3),
	1780: uint16(1),
	1781: uint16(sym_comment),
	1782: uint16(287),
	1783: uint16(1),
	1784: uint16(anon_sym_LT),
	1785: uint16(285),
	1786: uint16(6),
	1787: uint16(sym__text_fragment),
	1789: uint16(anon_sym_LT_BANG),
	1790: uint16(anon_sym_LT_SLASH),
	1791: uint16(sym_entity),
	1792: uint16(anon_sym_LBRACE_LBRACE),
	1793: uint16(3),
	1794: uint16(3),
	1795: uint16(1),
	1796: uint16(sym_comment),
	1797: uint16(263),
	1798: uint16(1),
	1799: uint16(anon_sym_LT),
	1800: uint16(261),
	1801: uint16(6),
	1802: uint16(sym__text_fragment),
	1804: uint16(anon_sym_LT_BANG),
	1805: uint16(anon_sym_LT_SLASH),
	1806: uint16(sym_entity),
	1807: uint16(anon_sym_LBRACE_LBRACE),
	1808: uint16(3),
	1809: uint16(3),
	1810: uint16(1),
	1811: uint16(sym_comment),
	1812: uint16(291),
	1813: uint16(1),
	1814: uint16(anon_sym_LT),
	1815: uint16(289),
	1816: uint16(6),
	1817: uint16(sym__text_fragment),
	1819: uint16(anon_sym_LT_BANG),
	1820: uint16(anon_sym_LT_SLASH),
	1821: uint16(sym_entity),
	1822: uint16(anon_sym_LBRACE_LBRACE),
	1823: uint16(3),
	1824: uint16(3),
	1825: uint16(1),
	1826: uint16(sym_comment),
	1827: uint16(295),
	1828: uint16(1),
	1829: uint16(anon_sym_LT),
	1830: uint16(293),
	1831: uint16(6),
	1832: uint16(sym__implicit_end_tag),
	1833: uint16(sym__text_fragment),
	1834: uint16(anon_sym_LT_BANG),
	1835: uint16(anon_sym_LT_SLASH),
	1836: uint16(sym_entity),
	1837: uint16(anon_sym_LBRACE_LBRACE),
	1838: uint16(3),
	1839: uint16(3),
	1840: uint16(1),
	1841: uint16(sym_comment),
	1842: uint16(299),
	1843: uint16(1),
	1844: uint16(anon_sym_LT),
	1845: uint16(297),
	1846: uint16(6),
	1847: uint16(sym__text_fragment),
	1849: uint16(anon_sym_LT_BANG),
	1850: uint16(anon_sym_LT_SLASH),
	1851: uint16(sym_entity),
	1852: uint16(anon_sym_LBRACE_LBRACE),
	1853: uint16(3),
	1854: uint16(3),
	1855: uint16(1),
	1856: uint16(sym_comment),
	1857: uint16(303),
	1858: uint16(1),
	1859: uint16(anon_sym_LT),
	1860: uint16(301),
	1861: uint16(6),
	1862: uint16(sym__implicit_end_tag),
	1863: uint16(sym__text_fragment),
	1864: uint16(anon_sym_LT_BANG),
	1865: uint16(anon_sym_LT_SLASH),
	1866: uint16(sym_entity),
	1867: uint16(anon_sym_LBRACE_LBRACE),
	1868: uint16(3),
	1869: uint16(3),
	1870: uint16(1),
	1871: uint16(sym_comment),
	1872: uint16(291),
	1873: uint16(1),
	1874: uint16(anon_sym_LT),
	1875: uint16(289),
	1876: uint16(6),
	1877: uint16(sym__implicit_end_tag),
	1878: uint16(sym__text_fragment),
	1879: uint16(anon_sym_LT_BANG),
	1880: uint16(anon_sym_LT_SLASH),
	1881: uint16(sym_entity),
	1882: uint16(anon_sym_LBRACE_LBRACE),
	1883: uint16(3),
	1884: uint16(3),
	1885: uint16(1),
	1886: uint16(sym_comment),
	1887: uint16(307),
	1888: uint16(1),
	1889: uint16(anon_sym_LT),
	1890: uint16(305),
	1891: uint16(6),
	1892: uint16(sym__implicit_end_tag),
	1893: uint16(sym__text_fragment),
	1894: uint16(anon_sym_LT_BANG),
	1895: uint16(anon_sym_LT_SLASH),
	1896: uint16(sym_entity),
	1897: uint16(anon_sym_LBRACE_LBRACE),
	1898: uint16(3),
	1899: uint16(3),
	1900: uint16(1),
	1901: uint16(sym_comment),
	1902: uint16(275),
	1903: uint16(1),
	1904: uint16(anon_sym_LT),
	1905: uint16(273),
	1906: uint16(6),
	1907: uint16(sym__implicit_end_tag),
	1908: uint16(sym__text_fragment),
	1909: uint16(anon_sym_LT_BANG),
	1910: uint16(anon_sym_LT_SLASH),
	1911: uint16(sym_entity),
	1912: uint16(anon_sym_LBRACE_LBRACE),
	1913: uint16(3),
	1914: uint16(3),
	1915: uint16(1),
	1916: uint16(sym_comment),
	1917: uint16(283),
	1918: uint16(1),
	1919: uint16(anon_sym_LT),
	1920: uint16(281),
	1921: uint16(6),
	1922: uint16(sym__implicit_end_tag),
	1923: uint16(sym__text_fragment),
	1924: uint16(anon_sym_LT_BANG),
	1925: uint16(anon_sym_LT_SLASH),
	1926: uint16(sym_entity),
	1927: uint16(anon_sym_LBRACE_LBRACE),
	1928: uint16(3),
	1929: uint16(3),
	1930: uint16(1),
	1931: uint16(sym_comment),
	1932: uint16(287),
	1933: uint16(1),
	1934: uint16(anon_sym_LT),
	1935: uint16(285),
	1936: uint16(6),
	1937: uint16(sym__implicit_end_tag),
	1938: uint16(sym__text_fragment),
	1939: uint16(anon_sym_LT_BANG),
	1940: uint16(anon_sym_LT_SLASH),
	1941: uint16(sym_entity),
	1942: uint16(anon_sym_LBRACE_LBRACE),
	1943: uint16(3),
	1944: uint16(3),
	1945: uint16(1),
	1946: uint16(sym_comment),
	1947: uint16(299),
	1948: uint16(1),
	1949: uint16(anon_sym_LT),
	1950: uint16(297),
	1951: uint16(6),
	1952: uint16(sym__implicit_end_tag),
	1953: uint16(sym__text_fragment),
	1954: uint16(anon_sym_LT_BANG),
	1955: uint16(anon_sym_LT_SLASH),
	1956: uint16(sym_entity),
	1957: uint16(anon_sym_LBRACE_LBRACE),
	1958: uint16(3),
	1959: uint16(3),
	1960: uint16(1),
	1961: uint16(sym_comment),
	1962: uint16(311),
	1963: uint16(1),
	1964: uint16(anon_sym_LT),
	1965: uint16(309),
	1966: uint16(6),
	1967: uint16(sym__implicit_end_tag),
	1968: uint16(sym__text_fragment),
	1969: uint16(anon_sym_LT_BANG),
	1970: uint16(anon_sym_LT_SLASH),
	1971: uint16(sym_entity),
	1972: uint16(anon_sym_LBRACE_LBRACE),
	1973: uint16(3),
	1974: uint16(3),
	1975: uint16(1),
	1976: uint16(sym_comment),
	1977: uint16(315),
	1978: uint16(1),
	1979: uint16(anon_sym_LT),
	1980: uint16(313),
	1981: uint16(6),
	1982: uint16(sym__implicit_end_tag),
	1983: uint16(sym__text_fragment),
	1984: uint16(anon_sym_LT_BANG),
	1985: uint16(anon_sym_LT_SLASH),
	1986: uint16(sym_entity),
	1987: uint16(anon_sym_LBRACE_LBRACE),
	1988: uint16(3),
	1989: uint16(3),
	1990: uint16(1),
	1991: uint16(sym_comment),
	1992: uint16(319),
	1993: uint16(1),
	1994: uint16(anon_sym_LT),
	1995: uint16(317),
	1996: uint16(6),
	1997: uint16(sym__implicit_end_tag),
	1998: uint16(sym__text_fragment),
	1999: uint16(anon_sym_LT_BANG),
	2000: uint16(anon_sym_LT_SLASH),
	2001: uint16(sym_entity),
	2002: uint16(anon_sym_LBRACE_LBRACE),
	2003: uint16(3),
	2004: uint16(3),
	2005: uint16(1),
	2006: uint16(sym_comment),
	2007: uint16(323),
	2008: uint16(1),
	2009: uint16(anon_sym_LT),
	2010: uint16(321),
	2011: uint16(6),
	2012: uint16(sym__implicit_end_tag),
	2013: uint16(sym__text_fragment),
	2014: uint16(anon_sym_LT_BANG),
	2015: uint16(anon_sym_LT_SLASH),
	2016: uint16(sym_entity),
	2017: uint16(anon_sym_LBRACE_LBRACE),
	2018: uint16(3),
	2019: uint16(3),
	2020: uint16(1),
	2021: uint16(sym_comment),
	2022: uint16(327),
	2023: uint16(1),
	2024: uint16(anon_sym_LT),
	2025: uint16(325),
	2026: uint16(6),
	2027: uint16(sym__implicit_end_tag),
	2028: uint16(sym__text_fragment),
	2029: uint16(anon_sym_LT_BANG),
	2030: uint16(anon_sym_LT_SLASH),
	2031: uint16(sym_entity),
	2032: uint16(anon_sym_LBRACE_LBRACE),
	2033: uint16(3),
	2034: uint16(3),
	2035: uint16(1),
	2036: uint16(sym_comment),
	2037: uint16(331),
	2038: uint16(1),
	2039: uint16(anon_sym_LT),
	2040: uint16(329),
	2041: uint16(6),
	2042: uint16(sym__implicit_end_tag),
	2043: uint16(sym__text_fragment),
	2044: uint16(anon_sym_LT_BANG),
	2045: uint16(anon_sym_LT_SLASH),
	2046: uint16(sym_entity),
	2047: uint16(anon_sym_LBRACE_LBRACE),
	2048: uint16(3),
	2049: uint16(3),
	2050: uint16(1),
	2051: uint16(sym_comment),
	2052: uint16(335),
	2053: uint16(1),
	2054: uint16(anon_sym_LT),
	2055: uint16(333),
	2056: uint16(6),
	2057: uint16(sym__implicit_end_tag),
	2058: uint16(sym__text_fragment),
	2059: uint16(anon_sym_LT_BANG),
	2060: uint16(anon_sym_LT_SLASH),
	2061: uint16(sym_entity),
	2062: uint16(anon_sym_LBRACE_LBRACE),
	2063: uint16(3),
	2064: uint16(3),
	2065: uint16(1),
	2066: uint16(sym_comment),
	2067: uint16(271),
	2068: uint16(1),
	2069: uint16(anon_sym_LT),
	2070: uint16(269),
	2071: uint16(6),
	2072: uint16(sym__implicit_end_tag),
	2073: uint16(sym__text_fragment),
	2074: uint16(anon_sym_LT_BANG),
	2075: uint16(anon_sym_LT_SLASH),
	2076: uint16(sym_entity),
	2077: uint16(anon_sym_LBRACE_LBRACE),
	2078: uint16(3),
	2079: uint16(3),
	2080: uint16(1),
	2081: uint16(sym_comment),
	2082: uint16(303),
	2083: uint16(1),
	2084: uint16(anon_sym_LT),
	2085: uint16(301),
	2086: uint16(6),
	2087: uint16(sym__text_fragment),
	2089: uint16(anon_sym_LT_BANG),
	2090: uint16(anon_sym_LT_SLASH),
	2091: uint16(sym_entity),
	2092: uint16(anon_sym_LBRACE_LBRACE),
	2093: uint16(3),
	2094: uint16(3),
	2095: uint16(1),
	2096: uint16(sym_comment),
	2097: uint16(279),
	2098: uint16(1),
	2099: uint16(anon_sym_LT),
	2100: uint16(277),
	2101: uint16(6),
	2102: uint16(sym__implicit_end_tag),
	2103: uint16(sym__text_fragment),
	2104: uint16(anon_sym_LT_BANG),
	2105: uint16(anon_sym_LT_SLASH),
	2106: uint16(sym_entity),
	2107: uint16(anon_sym_LBRACE_LBRACE),
	2108: uint16(3),
	2109: uint16(3),
	2110: uint16(1),
	2111: uint16(sym_comment),
	2112: uint16(241),
	2113: uint16(1),
	2114: uint16(sym_attribute_name),
	2115: uint16(239),
	2116: uint16(6),
	2117: uint16(anon_sym_GT),
	2118: uint16(anon_sym_COLON2),
	2119: uint16(anon_sym_DOT),
	2120: uint16(anon_sym_AT),
	2121: uint16(anon_sym_POUND),
	2122: uint16(sym_directive_name),
	2123: uint16(3),
	2124: uint16(3),
	2125: uint16(1),
	2126: uint16(sym_comment),
	2127: uint16(311),
	2128: uint16(1),
	2129: uint16(anon_sym_LT),
	2130: uint16(309),
	2131: uint16(6),
	2132: uint16(sym__text_fragment),
	2134: uint16(anon_sym_LT_BANG),
	2135: uint16(anon_sym_LT_SLASH),
	2136: uint16(sym_entity),
	2137: uint16(anon_sym_LBRACE_LBRACE),
	2138: uint16(3),
	2139: uint16(3),
	2140: uint16(1),
	2141: uint16(sym_comment),
	2142: uint16(167),
	2143: uint16(1),
	2144: uint16(sym_attribute_name),
	2145: uint16(163),
	2146: uint16(6),
	2147: uint16(anon_sym_GT),
	2148: uint16(anon_sym_COLON2),
	2149: uint16(anon_sym_DOT),
	2150: uint16(anon_sym_AT),
	2151: uint16(anon_sym_POUND),
	2152: uint16(sym_directive_name),
	2153: uint16(3),
	2154: uint16(3),
	2155: uint16(1),
	2156: uint16(sym_comment),
	2157: uint16(315),
	2158: uint16(1),
	2159: uint16(anon_sym_LT),
	2160: uint16(313),
	2161: uint16(6),
	2162: uint16(sym__text_fragment),
	2164: uint16(anon_sym_LT_BANG),
	2165: uint16(anon_sym_LT_SLASH),
	2166: uint16(sym_entity),
	2167: uint16(anon_sym_LBRACE_LBRACE),
	2168: uint16(3),
	2169: uint16(3),
	2170: uint16(1),
	2171: uint16(sym_comment),
	2172: uint16(319),
	2173: uint16(1),
	2174: uint16(anon_sym_LT),
	2175: uint16(317),
	2176: uint16(6),
	2177: uint16(sym__text_fragment),
	2179: uint16(anon_sym_LT_BANG),
	2180: uint16(anon_sym_LT_SLASH),
	2181: uint16(sym_entity),
	2182: uint16(anon_sym_LBRACE_LBRACE),
	2183: uint16(3),
	2184: uint16(3),
	2185: uint16(1),
	2186: uint16(sym_comment),
	2187: uint16(257),
	2188: uint16(1),
	2189: uint16(sym_attribute_name),
	2190: uint16(255),
	2191: uint16(6),
	2192: uint16(anon_sym_GT),
	2193: uint16(anon_sym_COLON2),
	2194: uint16(anon_sym_DOT),
	2195: uint16(anon_sym_AT),
	2196: uint16(anon_sym_POUND),
	2197: uint16(sym_directive_name),
	2198: uint16(3),
	2199: uint16(3),
	2200: uint16(1),
	2201: uint16(sym_comment),
	2202: uint16(323),
	2203: uint16(1),
	2204: uint16(anon_sym_LT),
	2205: uint16(321),
	2206: uint16(6),
	2207: uint16(sym__text_fragment),
	2209: uint16(anon_sym_LT_BANG),
	2210: uint16(anon_sym_LT_SLASH),
	2211: uint16(sym_entity),
	2212: uint16(anon_sym_LBRACE_LBRACE),
	2213: uint16(3),
	2214: uint16(3),
	2215: uint16(1),
	2216: uint16(sym_comment),
	2217: uint16(327),
	2218: uint16(1),
	2219: uint16(anon_sym_LT),
	2220: uint16(325),
	2221: uint16(6),
	2222: uint16(sym__text_fragment),
	2224: uint16(anon_sym_LT_BANG),
	2225: uint16(anon_sym_LT_SLASH),
	2226: uint16(sym_entity),
	2227: uint16(anon_sym_LBRACE_LBRACE),
	2228: uint16(3),
	2229: uint16(3),
	2230: uint16(1),
	2231: uint16(sym_comment),
	2232: uint16(227),
	2233: uint16(1),
	2234: uint16(sym_attribute_name),
	2235: uint16(223),
	2236: uint16(6),
	2237: uint16(anon_sym_GT),
	2238: uint16(anon_sym_COLON2),
	2239: uint16(anon_sym_DOT),
	2240: uint16(anon_sym_AT),
	2241: uint16(anon_sym_POUND),
	2242: uint16(sym_directive_name),
	2243: uint16(3),
	2244: uint16(3),
	2245: uint16(1),
	2246: uint16(sym_comment),
	2247: uint16(253),
	2248: uint16(1),
	2249: uint16(sym_attribute_name),
	2250: uint16(251),
	2251: uint16(6),
	2252: uint16(anon_sym_GT),
	2253: uint16(anon_sym_COLON2),
	2254: uint16(anon_sym_DOT),
	2255: uint16(anon_sym_AT),
	2256: uint16(anon_sym_POUND),
	2257: uint16(sym_directive_name),
	2258: uint16(3),
	2259: uint16(3),
	2260: uint16(1),
	2261: uint16(sym_comment),
	2262: uint16(247),
	2263: uint16(1),
	2264: uint16(sym_attribute_name),
	2265: uint16(245),
	2266: uint16(6),
	2267: uint16(anon_sym_GT),
	2268: uint16(anon_sym_COLON2),
	2269: uint16(anon_sym_DOT),
	2270: uint16(anon_sym_AT),
	2271: uint16(anon_sym_POUND),
	2272: uint16(sym_directive_name),
	2273: uint16(3),
	2274: uint16(3),
	2275: uint16(1),
	2276: uint16(sym_comment),
	2277: uint16(237),
	2278: uint16(1),
	2279: uint16(sym_attribute_name),
	2280: uint16(235),
	2281: uint16(6),
	2282: uint16(anon_sym_GT),
	2283: uint16(anon_sym_COLON2),
	2284: uint16(anon_sym_DOT),
	2285: uint16(anon_sym_AT),
	2286: uint16(anon_sym_POUND),
	2287: uint16(sym_directive_name),
	2288: uint16(3),
	2289: uint16(3),
	2290: uint16(1),
	2291: uint16(sym_comment),
	2292: uint16(331),
	2293: uint16(1),
	2294: uint16(anon_sym_LT),
	2295: uint16(329),
	2296: uint16(6),
	2297: uint16(sym__text_fragment),
	2299: uint16(anon_sym_LT_BANG),
	2300: uint16(anon_sym_LT_SLASH),
	2301: uint16(sym_entity),
	2302: uint16(anon_sym_LBRACE_LBRACE),
	2303: uint16(3),
	2304: uint16(3),
	2305: uint16(1),
	2306: uint16(sym_comment),
	2307: uint16(335),
	2308: uint16(1),
	2309: uint16(anon_sym_LT),
	2310: uint16(333),
	2311: uint16(6),
	2312: uint16(sym__text_fragment),
	2314: uint16(anon_sym_LT_BANG),
	2315: uint16(anon_sym_LT_SLASH),
	2316: uint16(sym_entity),
	2317: uint16(anon_sym_LBRACE_LBRACE),
	2318: uint16(3),
	2319: uint16(3),
	2320: uint16(1),
	2321: uint16(sym_comment),
	2322: uint16(307),
	2323: uint16(1),
	2324: uint16(anon_sym_LT),
	2325: uint16(305),
	2326: uint16(6),
	2327: uint16(sym__text_fragment),
	2329: uint16(anon_sym_LT_BANG),
	2330: uint16(anon_sym_LT_SLASH),
	2331: uint16(sym_entity),
	2332: uint16(anon_sym_LBRACE_LBRACE),
	2333: uint16(3),
	2334: uint16(3),
	2335: uint16(1),
	2336: uint16(sym_comment),
	2337: uint16(339),
	2338: uint16(1),
	2339: uint16(anon_sym_LT),
	2340: uint16(337),
	2341: uint16(5),
	2342: uint16(sym__text_fragment),
	2343: uint16(anon_sym_LT_BANG),
	2344: uint16(anon_sym_LT_SLASH),
	2345: uint16(sym_entity),
	2346: uint16(anon_sym_LBRACE_LBRACE),
	2347: uint16(3),
	2348: uint16(3),
	2349: uint16(1),
	2350: uint16(sym_comment),
	2351: uint16(343),
	2352: uint16(1),
	2353: uint16(anon_sym_LT),
	2354: uint16(341),
	2355: uint16(5),
	2356: uint16(sym__text_fragment),
	2357: uint16(anon_sym_LT_BANG),
	2358: uint16(anon_sym_LT_SLASH),
	2359: uint16(sym_entity),
	2360: uint16(anon_sym_LBRACE_LBRACE),
	2361: uint16(5),
	2362: uint16(3),
	2363: uint16(1),
	2364: uint16(sym_comment),
	2365: uint16(345),
	2366: uint16(1),
	2367: uint16(sym_attribute_value),
	2368: uint16(347),
	2369: uint16(1),
	2370: uint16(anon_sym_SQUOTE),
	2371: uint16(349),
	2372: uint16(1),
	2373: uint16(anon_sym_DQUOTE),
	2374: uint16(59),
	2375: uint16(1),
	2376: uint16(sym_quoted_attribute_value),
	2377: uint16(5),
	2378: uint16(3),
	2379: uint16(1),
	2380: uint16(sym_comment),
	2381: uint16(347),
	2382: uint16(1),
	2383: uint16(anon_sym_SQUOTE),
	2384: uint16(349),
	2385: uint16(1),
	2386: uint16(anon_sym_DQUOTE),
	2387: uint16(351),
	2388: uint16(1),
	2389: uint16(sym_attribute_value),
	2390: uint16(50),
	2391: uint16(1),
	2392: uint16(sym_quoted_attribute_value),
	2393: uint16(5),
	2394: uint16(3),
	2395: uint16(1),
	2396: uint16(sym_comment),
	2397: uint16(347),
	2398: uint16(1),
	2399: uint16(anon_sym_SQUOTE),
	2400: uint16(349),
	2401: uint16(1),
	2402: uint16(anon_sym_DQUOTE),
	2403: uint16(353),
	2404: uint16(1),
	2405: uint16(sym_attribute_value),
	2406: uint16(53),
	2407: uint16(1),
	2408: uint16(sym_quoted_attribute_value),
	2409: uint16(5),
	2410: uint16(3),
	2411: uint16(1),
	2412: uint16(sym_comment),
	2413: uint16(355),
	2414: uint16(1),
	2415: uint16(sym__start_tag_name),
	2416: uint16(357),
	2417: uint16(1),
	2418: uint16(sym__script_start_tag_name),
	2419: uint16(359),
	2420: uint16(1),
	2421: uint16(sym__style_start_tag_name),
	2422: uint16(361),
	2423: uint16(1),
	2424: uint16(sym__template_start_tag_name),
	2425: uint16(5),
	2426: uint16(3),
	2427: uint16(1),
	2428: uint16(sym_comment),
	2429: uint16(347),
	2430: uint16(1),
	2431: uint16(anon_sym_SQUOTE),
	2432: uint16(349),
	2433: uint16(1),
	2434: uint16(anon_sym_DQUOTE),
	2435: uint16(363),
	2436: uint16(1),
	2437: uint16(sym_attribute_value),
	2438: uint16(51),
	2439: uint16(1),
	2440: uint16(sym_quoted_attribute_value),
	2441: uint16(5),
	2442: uint16(3),
	2443: uint16(1),
	2444: uint16(sym_comment),
	2445: uint16(347),
	2446: uint16(1),
	2447: uint16(anon_sym_SQUOTE),
	2448: uint16(349),
	2449: uint16(1),
	2450: uint16(anon_sym_DQUOTE),
	2451: uint16(365),
	2452: uint16(1),
	2453: uint16(sym_attribute_value),
	2454: uint16(57),
	2455: uint16(1),
	2456: uint16(sym_quoted_attribute_value),
	2457: uint16(5),
	2458: uint16(3),
	2459: uint16(1),
	2460: uint16(sym_comment),
	2461: uint16(367),
	2462: uint16(1),
	2463: uint16(sym_attribute_value),
	2464: uint16(369),
	2465: uint16(1),
	2466: uint16(anon_sym_SQUOTE),
	2467: uint16(371),
	2468: uint16(1),
	2469: uint16(anon_sym_DQUOTE),
	2470: uint16(90),
	2471: uint16(1),
	2472: uint16(sym_quoted_attribute_value),
	2473: uint16(5),
	2474: uint16(3),
	2475: uint16(1),
	2476: uint16(sym_comment),
	2477: uint16(369),
	2478: uint16(1),
	2479: uint16(anon_sym_SQUOTE),
	2480: uint16(371),
	2481: uint16(1),
	2482: uint16(anon_sym_DQUOTE),
	2483: uint16(373),
	2484: uint16(1),
	2485: uint16(sym_attribute_value),
	2486: uint16(98),
	2487: uint16(1),
	2488: uint16(sym_quoted_attribute_value),
	2489: uint16(5),
	2490: uint16(3),
	2491: uint16(1),
	2492: uint16(sym_comment),
	2493: uint16(369),
	2494: uint16(1),
	2495: uint16(anon_sym_SQUOTE),
	2496: uint16(371),
	2497: uint16(1),
	2498: uint16(anon_sym_DQUOTE),
	2499: uint16(375),
	2500: uint16(1),
	2501: uint16(sym_attribute_value),
	2502: uint16(88),
	2503: uint16(1),
	2504: uint16(sym_quoted_attribute_value),
	2505: uint16(5),
	2506: uint16(3),
	2507: uint16(1),
	2508: uint16(sym_comment),
	2509: uint16(369),
	2510: uint16(1),
	2511: uint16(anon_sym_SQUOTE),
	2512: uint16(371),
	2513: uint16(1),
	2514: uint16(anon_sym_DQUOTE),
	2515: uint16(377),
	2516: uint16(1),
	2517: uint16(sym_attribute_value),
	2518: uint16(96),
	2519: uint16(1),
	2520: uint16(sym_quoted_attribute_value),
	2521: uint16(5),
	2522: uint16(3),
	2523: uint16(1),
	2524: uint16(sym_comment),
	2525: uint16(369),
	2526: uint16(1),
	2527: uint16(anon_sym_SQUOTE),
	2528: uint16(371),
	2529: uint16(1),
	2530: uint16(anon_sym_DQUOTE),
	2531: uint16(379),
	2532: uint16(1),
	2533: uint16(sym_attribute_value),
	2534: uint16(99),
	2535: uint16(1),
	2536: uint16(sym_quoted_attribute_value),
	2537: uint16(5),
	2538: uint16(3),
	2539: uint16(1),
	2540: uint16(sym_comment),
	2541: uint16(357),
	2542: uint16(1),
	2543: uint16(sym__script_start_tag_name),
	2544: uint16(359),
	2545: uint16(1),
	2546: uint16(sym__style_start_tag_name),
	2547: uint16(381),
	2548: uint16(1),
	2549: uint16(sym__start_tag_name),
	2550: uint16(383),
	2551: uint16(1),
	2552: uint16(sym__template_start_tag_name),
	2553: uint16(4),
	2554: uint16(3),
	2555: uint16(1),
	2556: uint16(sym_comment),
	2557: uint16(385),
	2558: uint16(1),
	2559: uint16(anon_sym_LT_SLASH),
	2560: uint16(387),
	2561: uint16(1),
	2562: uint16(sym_raw_text),
	2563: uint16(65),
	2564: uint16(1),
	2565: uint16(sym_end_tag),
	2566: uint16(4),
	2567: uint16(3),
	2568: uint16(1),
	2569: uint16(sym_comment),
	2570: uint16(385),
	2571: uint16(1),
	2572: uint16(anon_sym_LT_SLASH),
	2573: uint16(389),
	2574: uint16(1),
	2575: uint16(sym_raw_text),
	2576: uint16(66),
	2577: uint16(1),
	2578: uint16(sym_end_tag),
	2579: uint16(4),
	2580: uint16(3),
	2581: uint16(1),
	2582: uint16(sym_comment),
	2583: uint16(391),
	2584: uint16(1),
	2585: uint16(sym_directive_value),
	2586: uint16(393),
	2587: uint16(1),
	2588: uint16(anon_sym_LBRACK),
	2589: uint16(42),
	2590: uint16(1),
	2591: uint16(sym_dynamic_directive_value),
	2592: uint16(4),
	2593: uint16(3),
	2594: uint16(1),
	2595: uint16(sym_comment),
	2596: uint16(395),
	2597: uint16(1),
	2598: uint16(anon_sym_LT_SLASH),
	2599: uint16(397),
	2600: uint16(1),
	2601: uint16(sym_raw_text),
	2602: uint16(75),
	2603: uint16(1),
	2604: uint16(sym_end_tag),
	2605: uint16(4),
	2606: uint16(3),
	2607: uint16(1),
	2608: uint16(sym_comment),
	2609: uint16(395),
	2610: uint16(1),
	2611: uint16(anon_sym_LT_SLASH),
	2612: uint16(399),
	2613: uint16(1),
	2614: uint16(sym_raw_text),
	2615: uint16(76),
	2616: uint16(1),
	2617: uint16(sym_end_tag),
	2618: uint16(4),
	2619: uint16(3),
	2620: uint16(1),
	2621: uint16(sym_comment),
	2622: uint16(401),
	2623: uint16(1),
	2624: uint16(sym_directive_value),
	2625: uint16(403),
	2626: uint16(1),
	2627: uint16(anon_sym_LBRACK),
	2628: uint16(35),
	2629: uint16(1),
	2630: uint16(sym_dynamic_directive_value),
	2631: uint16(4),
	2632: uint16(3),
	2633: uint16(1),
	2634: uint16(sym_comment),
	2635: uint16(403),
	2636: uint16(1),
	2637: uint16(anon_sym_LBRACK),
	2638: uint16(405),
	2639: uint16(1),
	2640: uint16(sym_directive_value),
	2641: uint16(44),
	2642: uint16(1),
	2643: uint16(sym_dynamic_directive_value),
	2644: uint16(4),
	2645: uint16(3),
	2646: uint16(1),
	2647: uint16(sym_comment),
	2648: uint16(393),
	2649: uint16(1),
	2650: uint16(anon_sym_LBRACK),
	2651: uint16(407),
	2652: uint16(1),
	2653: uint16(sym_directive_value),
	2654: uint16(29),
	2655: uint16(1),
	2656: uint16(sym_dynamic_directive_value),
	2657: uint16(3),
	2658: uint16(3),
	2659: uint16(1),
	2660: uint16(sym_comment),
	2661: uint16(409),
	2662: uint16(1),
	2663: uint16(anon_sym_DQUOTE),
	2664: uint16(411),
	2665: uint16(1),
	2666: uint16(aux_sym_quoted_attribute_value_token2),
	2667: uint16(3),
	2668: uint16(3),
	2669: uint16(1),
	2670: uint16(sym_comment),
	2671: uint16(413),
	2672: uint16(1),
	2673: uint16(sym__end_tag_name),
	2674: uint16(415),
	2675: uint16(1),
	2676: uint16(sym_erroneous_end_tag_name),
	2677: uint16(2),
	2678: uint16(3),
	2679: uint16(1),
	2680: uint16(sym_comment),
	2681: uint16(417),
	2682: uint16(2),
	2683: uint16(sym_raw_text),
	2684: uint16(anon_sym_LT_SLASH),
	2685: uint16(2),
	2686: uint16(3),
	2687: uint16(1),
	2688: uint16(sym_comment),
	2689: uint16(419),
	2690: uint16(2),
	2691: uint16(sym_raw_text),
	2692: uint16(anon_sym_LT_SLASH),
	2693: uint16(3),
	2694: uint16(3),
	2695: uint16(1),
	2696: uint16(sym_comment),
	2697: uint16(421),
	2698: uint16(1),
	2699: uint16(anon_sym_SQUOTE),
	2700: uint16(423),
	2701: uint16(1),
	2702: uint16(aux_sym_quoted_attribute_value_token1),
	2703: uint16(3),
	2704: uint16(3),
	2705: uint16(1),
	2706: uint16(sym_comment),
	2707: uint16(425),
	2708: uint16(1),
	2709: uint16(anon_sym_RBRACE_RBRACE),
	2710: uint16(427),
	2711: uint16(1),
	2712: uint16(sym__interpolation_text),
	2713: uint16(3),
	2714: uint16(3),
	2715: uint16(1),
	2716: uint16(sym_comment),
	2717: uint16(429),
	2718: uint16(1),
	2719: uint16(anon_sym_RBRACE_RBRACE),
	2720: uint16(431),
	2721: uint16(1),
	2722: uint16(sym__interpolation_text),
	2723: uint16(2),
	2724: uint16(3),
	2725: uint16(1),
	2726: uint16(sym_comment),
	2727: uint16(433),
	2728: uint16(2),
	2729: uint16(sym_raw_text),
	2730: uint16(anon_sym_LT_SLASH),
	2731: uint16(3),
	2732: uint16(3),
	2733: uint16(1),
	2734: uint16(sym_comment),
	2735: uint16(435),
	2736: uint16(1),
	2737: uint16(sym__end_tag_name),
	2738: uint16(437),
	2739: uint16(1),
	2740: uint16(sym_erroneous_end_tag_name),
	2741: uint16(3),
	2742: uint16(3),
	2743: uint16(1),
	2744: uint16(sym_comment),
	2745: uint16(439),
	2746: uint16(1),
	2747: uint16(anon_sym_RBRACK),
	2748: uint16(441),
	2749: uint16(1),
	2750: uint16(sym_dynamic_directive_inner_value),
	2751: uint16(3),
	2752: uint16(3),
	2753: uint16(1),
	2754: uint16(sym_comment),
	2755: uint16(395),
	2756: uint16(1),
	2757: uint16(anon_sym_LT_SLASH),
	2758: uint16(81),
	2759: uint16(1),
	2760: uint16(sym_end_tag),
	2761: uint16(3),
	2762: uint16(3),
	2763: uint16(1),
	2764: uint16(sym_comment),
	2765: uint16(395),
	2766: uint16(1),
	2767: uint16(anon_sym_LT_SLASH),
	2768: uint16(82),
	2769: uint16(1),
	2770: uint16(sym_end_tag),
	2771: uint16(3),
	2772: uint16(3),
	2773: uint16(1),
	2774: uint16(sym_comment),
	2775: uint16(421),
	2776: uint16(1),
	2777: uint16(anon_sym_DQUOTE),
	2778: uint16(443),
	2779: uint16(1),
	2780: uint16(aux_sym_quoted_attribute_value_token2),
	2781: uint16(3),
	2782: uint16(3),
	2783: uint16(1),
	2784: uint16(sym_comment),
	2785: uint16(385),
	2786: uint16(1),
	2787: uint16(anon_sym_LT_SLASH),
	2788: uint16(95),
	2789: uint16(1),
	2790: uint16(sym_end_tag),
	2791: uint16(3),
	2792: uint16(3),
	2793: uint16(1),
	2794: uint16(sym_comment),
	2795: uint16(445),
	2796: uint16(1),
	2797: uint16(anon_sym_RBRACK),
	2798: uint16(447),
	2799: uint16(1),
	2800: uint16(sym_dynamic_directive_inner_value),
	2801: uint16(3),
	2802: uint16(3),
	2803: uint16(1),
	2804: uint16(sym_comment),
	2805: uint16(409),
	2806: uint16(1),
	2807: uint16(anon_sym_SQUOTE),
	2808: uint16(449),
	2809: uint16(1),
	2810: uint16(aux_sym_quoted_attribute_value_token1),
	2811: uint16(2),
	2812: uint16(3),
	2813: uint16(1),
	2814: uint16(sym_comment),
	2815: uint16(451),
	2816: uint16(2),
	2817: uint16(sym_raw_text),
	2818: uint16(anon_sym_LT_SLASH),
	2819: uint16(3),
	2820: uint16(3),
	2821: uint16(1),
	2822: uint16(sym_comment),
	2823: uint16(413),
	2824: uint16(1),
	2825: uint16(sym__end_tag_name),
	2826: uint16(437),
	2827: uint16(1),
	2828: uint16(sym_erroneous_end_tag_name),
	2829: uint16(3),
	2830: uint16(3),
	2831: uint16(1),
	2832: uint16(sym_comment),
	2833: uint16(415),
	2834: uint16(1),
	2835: uint16(sym_erroneous_end_tag_name),
	2836: uint16(435),
	2837: uint16(1),
	2838: uint16(sym__end_tag_name),
	2839: uint16(3),
	2840: uint16(3),
	2841: uint16(1),
	2842: uint16(sym_comment),
	2843: uint16(385),
	2844: uint16(1),
	2845: uint16(anon_sym_LT_SLASH),
	2846: uint16(94),
	2847: uint16(1),
	2848: uint16(sym_end_tag),
	2849: uint16(2),
	2850: uint16(3),
	2851: uint16(1),
	2852: uint16(sym_comment),
	2853: uint16(453),
	2854: uint16(1),
	2855: uint16(sym__doctype),
	2856: uint16(2),
	2857: uint16(3),
	2858: uint16(1),
	2859: uint16(sym_comment),
	2860: uint16(455),
	2861: uint16(1),
	2862: uint16(anon_sym_RBRACE_RBRACE),
	2863: uint16(2),
	2864: uint16(3),
	2865: uint16(1),
	2866: uint16(sym_comment),
	2867: uint16(457),
	2868: uint16(1),
	2869: uint16(anon_sym_RBRACK),
	2870: uint16(2),
	2871: uint16(3),
	2872: uint16(1),
	2873: uint16(sym_comment),
	2874: uint16(459),
	2875: uint16(1),
	2876: uint16(sym_directive_modifier),
	2877: uint16(2),
	2878: uint16(3),
	2879: uint16(1),
	2880: uint16(sym_comment),
	2881: uint16(461),
	2882: uint16(1),
	2883: uint16(anon_sym_GT),
	2884: uint16(2),
	2885: uint16(3),
	2886: uint16(1),
	2887: uint16(sym_comment),
	2888: uint16(463),
	2889: uint16(1),
	2890: uint16(anon_sym_GT),
	2891: uint16(2),
	2892: uint16(3),
	2893: uint16(1),
	2894: uint16(sym_comment),
	2895: uint16(465),
	2896: uint16(1),
	2898: uint16(2),
	2899: uint16(3),
	2900: uint16(1),
	2901: uint16(sym_comment),
	2902: uint16(467),
	2903: uint16(1),
	2904: uint16(aux_sym_doctype_token1),
	2905: uint16(2),
	2906: uint16(3),
	2907: uint16(1),
	2908: uint16(sym_comment),
	2909: uint16(437),
	2910: uint16(1),
	2911: uint16(sym_erroneous_end_tag_name),
	2912: uint16(2),
	2913: uint16(3),
	2914: uint16(1),
	2915: uint16(sym_comment),
	2916: uint16(469),
	2917: uint16(1),
	2918: uint16(sym_directive_modifier),
	2919: uint16(2),
	2920: uint16(3),
	2921: uint16(1),
	2922: uint16(sym_comment),
	2923: uint16(471),
	2924: uint16(1),
	2925: uint16(anon_sym_SQUOTE),
	2926: uint16(2),
	2927: uint16(3),
	2928: uint16(1),
	2929: uint16(sym_comment),
	2930: uint16(473),
	2931: uint16(1),
	2932: uint16(anon_sym_RBRACE_RBRACE),
	2933: uint16(2),
	2934: uint16(3),
	2935: uint16(1),
	2936: uint16(sym_comment),
	2937: uint16(475),
	2938: uint16(1),
	2939: uint16(anon_sym_RBRACK),
	2940: uint16(2),
	2941: uint16(3),
	2942: uint16(1),
	2943: uint16(sym_comment),
	2944: uint16(477),
	2945: uint16(1),
	2946: uint16(anon_sym_GT),
	2947: uint16(2),
	2948: uint16(3),
	2949: uint16(1),
	2950: uint16(sym_comment),
	2951: uint16(479),
	2952: uint16(1),
	2953: uint16(anon_sym_SQUOTE),
	2954: uint16(2),
	2955: uint16(3),
	2956: uint16(1),
	2957: uint16(sym_comment),
	2958: uint16(479),
	2959: uint16(1),
	2960: uint16(anon_sym_DQUOTE),
	2961: uint16(2),
	2962: uint16(3),
	2963: uint16(1),
	2964: uint16(sym_comment),
	2965: uint16(471),
	2966: uint16(1),
	2967: uint16(anon_sym_DQUOTE),
	2968: uint16(2),
	2969: uint16(3),
	2970: uint16(1),
	2971: uint16(sym_comment),
	2972: uint16(435),
	2973: uint16(1),
	2974: uint16(sym__end_tag_name),
	2975: uint16(2),
	2976: uint16(3),
	2977: uint16(1),
	2978: uint16(sym_comment),
	2979: uint16(481),
	2980: uint16(1),
	2981: uint16(anon_sym_GT),
	2982: uint16(2),
	2983: uint16(3),
	2984: uint16(1),
	2985: uint16(sym_comment),
	2986: uint16(483),
	2987: uint16(1),
	2988: uint16(anon_sym_GT),
	2989: uint16(2),
	2990: uint16(3),
	2991: uint16(1),
	2992: uint16(sym_comment),
	2993: uint16(415),
	2994: uint16(1),
	2995: uint16(sym_erroneous_end_tag_name),
	2996: uint16(2),
	2997: uint16(3),
	2998: uint16(1),
	2999: uint16(sym_comment),
	3000: uint16(485),
	3001: uint16(1),
	3002: uint16(aux_sym_doctype_token1),
	3003: uint16(2),
	3004: uint16(3),
	3005: uint16(1),
	3006: uint16(sym_comment),
	3007: uint16(487),
	3008: uint16(1),
	3009: uint16(anon_sym_GT),
	3010: uint16(2),
	3011: uint16(3),
	3012: uint16(1),
	3013: uint16(sym_comment),
	3014: uint16(413),
	3015: uint16(1),
	3016: uint16(sym__end_tag_name),
	3017: uint16(2),
	3018: uint16(3),
	3019: uint16(1),
	3020: uint16(sym_comment),
	3021: uint16(489),
	3022: uint16(1),
	3023: uint16(sym__doctype),
}

var ts_small_parse_table_map = [168]uint32_t{
	1:   uint32(55),
	2:   uint32(110),
	3:   uint32(165),
	4:   uint32(220),
	5:   uint32(272),
	6:   uint32(324),
	7:   uint32(376),
	8:   uint32(428),
	9:   uint32(480),
	10:  uint32(532),
	11:  uint32(584),
	12:  uint32(619),
	13:  uint32(650),
	14:  uint32(681),
	15:  uint32(712),
	16:  uint32(743),
	17:  uint32(774),
	18:  uint32(805),
	19:  uint32(836),
	20:  uint32(865),
	21:  uint32(896),
	22:  uint32(927),
	23:  uint32(955),
	24:  uint32(983),
	25:  uint32(1011),
	26:  uint32(1039),
	27:  uint32(1067),
	28:  uint32(1095),
	29:  uint32(1129),
	30:  uint32(1159),
	31:  uint32(1182),
	32:  uint32(1207),
	33:  uint32(1230),
	34:  uint32(1257),
	35:  uint32(1281),
	36:  uint32(1303),
	37:  uint32(1321),
	38:  uint32(1339),
	39:  uint32(1361),
	40:  uint32(1379),
	41:  uint32(1397),
	42:  uint32(1416),
	43:  uint32(1433),
	44:  uint32(1450),
	45:  uint32(1467),
	46:  uint32(1484),
	47:  uint32(1503),
	48:  uint32(1522),
	49:  uint32(1538),
	50:  uint32(1554),
	51:  uint32(1572),
	52:  uint32(1588),
	53:  uint32(1606),
	54:  uint32(1622),
	55:  uint32(1638),
	56:  uint32(1654),
	57:  uint32(1672),
	58:  uint32(1688),
	59:  uint32(1703),
	60:  uint32(1718),
	61:  uint32(1733),
	62:  uint32(1748),
	63:  uint32(1763),
	64:  uint32(1778),
	65:  uint32(1793),
	66:  uint32(1808),
	67:  uint32(1823),
	68:  uint32(1838),
	69:  uint32(1853),
	70:  uint32(1868),
	71:  uint32(1883),
	72:  uint32(1898),
	73:  uint32(1913),
	74:  uint32(1928),
	75:  uint32(1943),
	76:  uint32(1958),
	77:  uint32(1973),
	78:  uint32(1988),
	79:  uint32(2003),
	80:  uint32(2018),
	81:  uint32(2033),
	82:  uint32(2048),
	83:  uint32(2063),
	84:  uint32(2078),
	85:  uint32(2093),
	86:  uint32(2108),
	87:  uint32(2123),
	88:  uint32(2138),
	89:  uint32(2153),
	90:  uint32(2168),
	91:  uint32(2183),
	92:  uint32(2198),
	93:  uint32(2213),
	94:  uint32(2228),
	95:  uint32(2243),
	96:  uint32(2258),
	97:  uint32(2273),
	98:  uint32(2288),
	99:  uint32(2303),
	100: uint32(2318),
	101: uint32(2333),
	102: uint32(2347),
	103: uint32(2361),
	104: uint32(2377),
	105: uint32(2393),
	106: uint32(2409),
	107: uint32(2425),
	108: uint32(2441),
	109: uint32(2457),
	110: uint32(2473),
	111: uint32(2489),
	112: uint32(2505),
	113: uint32(2521),
	114: uint32(2537),
	115: uint32(2553),
	116: uint32(2566),
	117: uint32(2579),
	118: uint32(2592),
	119: uint32(2605),
	120: uint32(2618),
	121: uint32(2631),
	122: uint32(2644),
	123: uint32(2657),
	124: uint32(2667),
	125: uint32(2677),
	126: uint32(2685),
	127: uint32(2693),
	128: uint32(2703),
	129: uint32(2713),
	130: uint32(2723),
	131: uint32(2731),
	132: uint32(2741),
	133: uint32(2751),
	134: uint32(2761),
	135: uint32(2771),
	136: uint32(2781),
	137: uint32(2791),
	138: uint32(2801),
	139: uint32(2811),
	140: uint32(2819),
	141: uint32(2829),
	142: uint32(2839),
	143: uint32(2849),
	144: uint32(2856),
	145: uint32(2863),
	146: uint32(2870),
	147: uint32(2877),
	148: uint32(2884),
	149: uint32(2891),
	150: uint32(2898),
	151: uint32(2905),
	152: uint32(2912),
	153: uint32(2919),
	154: uint32(2926),
	155: uint32(2933),
	156: uint32(2940),
	157: uint32(2947),
	158: uint32(2954),
	159: uint32(2961),
	160: uint32(2968),
	161: uint32(2975),
	162: uint32(2982),
	163: uint32(2989),
	164: uint32(2996),
	165: uint32(3003),
	166: uint32(3010),
	167: uint32(3017),
}

var ts_parse_actions = [491]TSParseActionEntry{
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
		Fstate: uint16(libc.Int32FromInt32(145)),
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
		Fstate: uint16(libc.Int32FromInt32(108)),
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
		Fstate: uint16(libc.Int32FromInt32(7)),
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
		Fstate: uint16(libc.Int32FromInt32(131)),
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
		Fstate: uint16(libc.Int32FromInt32(86)),
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
		Fstate: uint16(libc.Int32FromInt32(169)),
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
		Fstate: uint16(libc.Int32FromInt32(116)),
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
		Fstate: uint16(libc.Int32FromInt32(143)),
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
		Fstate: uint16(libc.Int32FromInt32(3)),
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
		Fstate: uint16(libc.Int32FromInt32(130)),
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
		Fstate: uint16(libc.Int32FromInt32(63)),
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
		Fstate: uint16(libc.Int32FromInt32(71)),
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
		Fstate: uint16(libc.Int32FromInt32(10)),
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
		Fstate: uint16(libc.Int32FromInt32(92)),
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
		Fstate: uint16(libc.Int32FromInt32(126)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(133)),
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
		Fstate: uint16(libc.Int32FromInt32(8)),
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
		Fstate: uint16(libc.Int32FromInt32(9)),
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
		Fsymbol:      uint16(aux_sym_document_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(145)),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_document_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(108)),
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
		Fsymbol:      uint16(aux_sym_document_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(153)),
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
		Fsymbol:      uint16(aux_sym_document_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(9)),
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
		Fsymbol:      uint16(aux_sym_document_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(131)),
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
		Fsymbol:      uint16(aux_sym_document_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(86)),
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
		Fsymbol:      uint16(aux_sym_document_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(169)),
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
		Fsymbol:      uint16(aux_sym_document_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(116)),
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
		Fsymbol:      uint16(aux_sym_document_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(165)),
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
		Fsymbol:      uint16(aux_sym_document_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(10)),
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
		Fsymbol:      uint16(aux_sym_document_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(130)),
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
		Fsymbol:      uint16(aux_sym_document_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(71)),
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
		Fstate: uint16(libc.Int32FromInt32(142)),
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
		Fstate: uint16(libc.Int32FromInt32(12)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_directive_attribute),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
	100: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_directive_attribute),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(119)),
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
		Fcount: uint8(1),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(119)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(148)),
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
		Fstate: uint16(libc.Int32FromInt32(61)),
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
		Fstate: uint16(libc.Int32FromInt32(62)),
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
		Fstate: uint16(libc.Int32FromInt32(43)),
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
		Fstate: uint16(libc.Int32FromInt32(17)),
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
		Fstate: uint16(libc.Int32FromInt32(104)),
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
		Fstate: uint16(libc.Int32FromInt32(70)),
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
		Fstate: uint16(libc.Int32FromInt32(69)),
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
		Fstate: uint16(libc.Int32FromInt32(124)),
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
		Fstate: uint16(libc.Int32FromInt32(103)),
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
		Fstate: uint16(libc.Int32FromInt32(77)),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_start_tag_repeat1),
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
		Fsymbol:      uint16(aux_sym_start_tag_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(43)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_start_tag_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(119)),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_start_tag_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(17)),
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
		Fstate: uint16(libc.Int32FromInt32(85)),
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
		Fstate: uint16(libc.Int32FromInt32(141)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(54)),
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
		Fstate: uint16(libc.Int32FromInt32(123)),
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
		Fstate: uint16(libc.Int32FromInt32(31)),
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
		Fstate: uint16(libc.Int32FromInt32(132)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_start_tag_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(54)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	158: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      uint16(libc.Int32FromInt32(123)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	161: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      uint16(libc.Int32FromInt32(31)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_directive_attribute),
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
		Fcount: uint8(1),
	}})),
	168: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_directive_attribute),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(111)),
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
		Fstate: uint16(libc.Int32FromInt32(123)),
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
		Fstate: uint16(libc.Int32FromInt32(154)),
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
		Fstate: uint16(libc.Int32FromInt32(122)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_directive_modifiers_repeat1),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_directive_modifiers_repeat1),
	})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	182: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_directive_modifiers_repeat1),
	})))),
	183: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
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
		Fsymbol:      uint16(aux_sym_directive_attribute_repeat1),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(aux_sym_directive_attribute_repeat1),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_directive_attribute_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(119)),
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
		Fsymbol:      uint16(aux_sym_directive_attribute_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(119)),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_directive_modifiers),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_directive_modifiers),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_directive_attribute_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(123)),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_directive_attribute_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(123)),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_dynamic_directive_value),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_dynamic_directive_value),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_dynamic_directive_value),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_dynamic_directive_value),
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
		Fsymbol:      uint16(aux_sym_directive_modifiers_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(154)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(sym_attribute),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fsymbol:      uint16(sym_attribute),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_directive_attribute),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(106)),
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
		Fcount: uint8(1),
	}})),
	228: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_directive_attribute),
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
		Fcount:    uint8(1),
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
		Fsymbol:      uint16(sym_directive_attribute),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_directive_attribute),
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
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_directive_attribute),
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
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_directive_attribute),
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
		Fsymbol:      uint16(sym_attribute),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_attribute),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_directive_attribute),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_directive_attribute),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fsymbol:      uint16(sym_quoted_attribute_value),
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
		Fsymbol:      uint16(sym_quoted_attribute_value),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_quoted_attribute_value),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_quoted_attribute_value),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_template_element),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_template_element),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_start_tag),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_start_tag),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_self_closing_tag),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_element),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_element),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_end_tag),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_end_tag),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_script_element),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_script_element),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_style_element),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_style_element),
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
		Fsymbol:      uint16(sym_element),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_element),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_start_tag),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_start_tag),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_self_closing_tag),
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
		Fsymbol:      uint16(sym_self_closing_tag),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_text),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_text),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_interpolation),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_interpolation),
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
		Fsymbol:      uint16(sym_erroneous_end_tag),
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
		Fsymbol:      uint16(sym_erroneous_end_tag),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_interpolation),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_interpolation),
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
		Fsymbol:      uint16(sym_element),
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
		Fsymbol:      uint16(sym_element),
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
		Fsymbol:      uint16(sym_script_element),
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
		Fsymbol:      uint16(sym_script_element),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_template_element),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_template_element),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_doctype),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_doctype),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_template_start_tag),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_template_start_tag),
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
		Fsymbol:      uint16(sym_template_start_tag),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_template_start_tag),
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
		Fstate: uint16(libc.Int32FromInt32(59)),
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
		Fstate: uint16(libc.Int32FromInt32(129)),
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
		Fstate: uint16(libc.Int32FromInt32(137)),
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
		Fstate: uint16(libc.Int32FromInt32(50)),
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
		Fstate: uint16(libc.Int32FromInt32(53)),
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
		Fstate: uint16(libc.Int32FromInt32(16)),
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
		Fstate: uint16(libc.Int32FromInt32(24)),
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
		Fstate: uint16(libc.Int32FromInt32(25)),
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
		Fstate: uint16(libc.Int32FromInt32(15)),
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
		Fstate: uint16(libc.Int32FromInt32(51)),
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
		Fstate: uint16(libc.Int32FromInt32(57)),
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
		Fstate: uint16(libc.Int32FromInt32(90)),
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
		Fstate: uint16(libc.Int32FromInt32(140)),
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
		Fstate: uint16(libc.Int32FromInt32(125)),
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
		Fstate: uint16(libc.Int32FromInt32(98)),
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
		Fstate: uint16(libc.Int32FromInt32(88)),
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
		Fstate: uint16(libc.Int32FromInt32(96)),
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
		Fstate: uint16(libc.Int32FromInt32(99)),
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
		Fstate: uint16(libc.Int32FromInt32(19)),
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
		Fstate: uint16(libc.Int32FromInt32(20)),
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
		Fstate: uint16(libc.Int32FromInt32(162)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(138)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(168)),
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
		Fstate: uint16(libc.Int32FromInt32(135)),
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
		Fstate: uint16(libc.Int32FromInt32(136)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(35)),
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
		Fstate: uint16(libc.Int32FromInt32(139)),
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
		Fstate: uint16(libc.Int32FromInt32(44)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(29)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(93)),
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
		Fstate: uint16(libc.Int32FromInt32(160)),
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
		Fstate: uint16(libc.Int32FromInt32(149)),
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
		Fstate: uint16(libc.Int32FromInt32(164)),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_style_start_tag),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(155)),
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
		Fstate: uint16(libc.Int32FromInt32(73)),
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
		Fstate: uint16(libc.Int32FromInt32(146)),
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
		Fstate: uint16(libc.Int32FromInt32(102)),
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
		Fstate: uint16(libc.Int32FromInt32(156)),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_style_start_tag),
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
		Fstate: uint16(libc.Int32FromInt32(158)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(163)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(147)),
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
		Fstate: uint16(libc.Int32FromInt32(161)),
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
		Fstate: uint16(libc.Int32FromInt32(45)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fsymbol:      uint16(sym_script_start_tag),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(152)),
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
		Fstate: uint16(libc.Int32FromInt32(79)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(39)),
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
		Fstate: uint16(libc.Int32FromInt32(41)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(87)),
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
		Fstate: uint16(libc.Int32FromInt32(101)),
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
	466: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f uint8_t
		_ [7]byte
	}{f: uint8(TSParseActionTypeAccept)})),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	468: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(46)),
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
		Fstate: uint16(libc.Int32FromInt32(55)),
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
		Fstate: uint16(libc.Int32FromInt32(91)),
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
		Fstate: uint16(libc.Int32FromInt32(47)),
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
		Fstate: uint16(libc.Int32FromInt32(64)),
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
		Fstate: uint16(libc.Int32FromInt32(97)),
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
		Fstate: uint16(libc.Int32FromInt32(89)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(167)),
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
		Fstate: uint16(libc.Int32FromInt32(84)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
const ts_external_token__template_start_tag_name = 9
const ts_external_token__text_fragment = 10
const ts_external_token__interpolation_text = 11

var ts_external_scanner_symbol_map = [12]TSSymbol{
	0:  uint16(sym__start_tag_name),
	1:  uint16(sym__script_start_tag_name),
	2:  uint16(sym__style_start_tag_name),
	3:  uint16(sym__end_tag_name),
	4:  uint16(sym_erroneous_end_tag_name),
	5:  uint16(anon_sym_SLASH_GT),
	6:  uint16(sym__implicit_end_tag),
	7:  uint16(sym_raw_text),
	8:  uint16(sym_comment),
	9:  uint16(sym__template_start_tag_name),
	10: uint16(sym__text_fragment),
	11: uint16(sym__interpolation_text),
}

var ts_external_scanner_states = [12][12]uint8{
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
	},
	2: {
		8:  libc.BoolUint8(true1 != 0),
		10: libc.BoolUint8(true1 != 0),
	},
	3: {
		6:  libc.BoolUint8(true1 != 0),
		8:  libc.BoolUint8(true1 != 0),
		10: libc.BoolUint8(true1 != 0),
	},
	4: {
		5: libc.BoolUint8(true1 != 0),
		8: libc.BoolUint8(true1 != 0),
	},
	5: {
		8: libc.BoolUint8(true1 != 0),
	},
	6: {
		0: libc.BoolUint8(true1 != 0),
		1: libc.BoolUint8(true1 != 0),
		2: libc.BoolUint8(true1 != 0),
		8: libc.BoolUint8(true1 != 0),
		9: libc.BoolUint8(true1 != 0),
	},
	7: {
		7: libc.BoolUint8(true1 != 0),
		8: libc.BoolUint8(true1 != 0),
	},
	8: {
		3: libc.BoolUint8(true1 != 0),
		4: libc.BoolUint8(true1 != 0),
		8: libc.BoolUint8(true1 != 0),
	},
	9: {
		8:  libc.BoolUint8(true1 != 0),
		11: libc.BoolUint8(true1 != 0),
	},
	10: {
		4: libc.BoolUint8(true1 != 0),
		8: libc.BoolUint8(true1 != 0),
	},
	11: {
		3: libc.BoolUint8(true1 != 0),
		8: libc.BoolUint8(true1 != 0),
	},
}

func tree_sitter_vue(tls *libc.TLS) (r uintptr) {
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
	Fname:              __ccgo_ts + 735,
	Fmetadata: TSLanguageMetadata{
		Fminor_version: uint8(1),
	},
}

func init() {
	p := unsafe.Pointer(&language)
	*(*uintptr)(unsafe.Add(p, 144)) = __ccgo_fp(ts_lex)
	*(*uintptr)(unsafe.Add(p, 184)) = __ccgo_fp(tree_sitter_vue_external_scanner_create)
	*(*uintptr)(unsafe.Add(p, 192)) = __ccgo_fp(tree_sitter_vue_external_scanner_destroy)
	*(*uintptr)(unsafe.Add(p, 200)) = __ccgo_fp(tree_sitter_vue_external_scanner_scan)
	*(*uintptr)(unsafe.Add(p, 208)) = __ccgo_fp(tree_sitter_vue_external_scanner_serialize)
	*(*uintptr)(unsafe.Add(p, 216)) = __ccgo_fp(tree_sitter_vue_external_scanner_deserialize)
}

func __ccgo_fp(f interface{}) uintptr {
	type iface [2]uintptr
	return (*iface)(unsafe.Pointer(&f))[1]
}

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = "(uint32_t)((&scanner->tags)->size - 1) < (&scanner->tags)->size\x00combined.c\x00</SCRIPT\x00</STYLE\x00(uint32_t)(i) < (&scanner->tags)->size\x00end\x00<!\x00doctype_token1\x00>\x00doctype\x00<\x00/>\x00</\x00=\x00attribute_name\x00attribute_value\x00entity\x00'\x00\"\x00{{\x00}}\x00:\x00.\x00@\x00#\x00directive_name\x00directive_value\x00[\x00]\x00dynamic_directive_inner_value\x00directive_modifier\x00tag_name\x00erroneous_end_tag_name\x00_implicit_end_tag\x00raw_text\x00comment\x00_text_fragment\x00document\x00_node\x00element\x00script_element\x00style_element\x00start_tag\x00self_closing_tag\x00end_tag\x00erroneous_end_tag\x00attribute\x00quoted_attribute_value\x00text\x00template_element\x00interpolation\x00_attribute\x00directive_attribute\x00dynamic_directive_value\x00directive_modifiers\x00document_repeat1\x00start_tag_repeat1\x00directive_attribute_repeat1\x00directive_modifiers_repeat1\x00vue\x00"
