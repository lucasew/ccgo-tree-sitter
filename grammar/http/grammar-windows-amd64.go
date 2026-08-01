// Code generated for windows/amd64 by 'ccgo -ignore-unsupported-alignment -ignore-unsupported-atomic-sizes -ignore-vector-functions --winapi-no-errno -ignore-link-errors -std=gnu11 -O0 -D__extension__= -D__forceinline=static inline -D__attribute__(...)= -D__declspec(x)= -D__cdecl= -D__stdcall= -D__fastcall= -D__thiscall= -D_cdecl= -D__restrict= -D__restrict__= -D__MINGW_EXTENSION= -D_X86INTRIN_H_INCLUDED -D_X86GPRINTRIN_H_INCLUDED -D_IMMINTRIN_H_INCLUDED -D_MMINTRIN_H_INCLUDED -D_XMMINTRIN_H_INCLUDED -D_EMMINTRIN_H_INCLUDED -D_PMMINTRIN_H_INCLUDED -D_MM3DNOW_H_INCLUDED -Dfunc=func_token -Dinterface=interface_token -Dselect=select_token -Dchan=chan_token -Dgo=go_token -Dmap=map_token -Dpackage=package_token -Dtype=type_token -Dvar=var_token -Dimport=import_token -Ddefer=defer_token -Dfallthrough=fallthrough_token -Drange=range_token -I D:\a\ccgo-tree-sitter\ccgo-tree-sitter\third-party\tree-sitter-http\src -I D:\a\ccgo-tree-sitter\ccgo-tree-sitter\third-party\tree-sitter-http -I D:\a\ccgo-tree-sitter\ccgo-tree-sitter\third-party -I C:\Users\runneradmin\.cache\workspaced\sources\github\7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2\lib\include -I C:\Users\runneradmin\.cache\workspaced\sources\github\7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2\lib\src D:\a\ccgo-tree-sitter\ccgo-tree-sitter\third-party\tree-sitter-http\src\parser.c -o grammar.go', DO NOT EDIT.

//go:build windows && amd64

package grammar_http

import (
	"reflect"
	"unsafe"

	"modernc.org/libc"
)

var _ reflect.Type
var _ unsafe.Pointer

const ALIAS_COUNT = 0
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
const EXTERNAL_TOKEN_COUNT = 0
const FIELD_COUNT = 10
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
const MAX_ALIAS_SEQUENCE_LENGTH = 9
const MB_LEN_MAX = 5
const MINGW_HAS_DDK_H = 1
const MINGW_HAS_SECURE_API = 1
const PATH_MAX = 260
const PRODUCTION_ID_COUNT = 49
const PTRDIFF_MAX = "INT64_MAX"
const PTRDIFF_MIN = "INT64_MIN"
const RAND_MAX = 0x7fff
const SIG_ATOMIC_MAX = "INT32_MAX"
const SIG_ATOMIC_MIN = "INT32_MIN"
const SIZE_MAX = "UINT64_MAX"
const SSIZE_MAX = "_I64_MAX"
const STATE_COUNT = 261
const STRUNCATE = 80
const SYMBOL_COUNT = 82
const TOKEN_COUNT = 40
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
const _CALL_REPORTFAULT = 0x2
const _CRTIMP2 = "_CRTIMP"
const _CRTIMP_ALTERNATIVE = "_CRTIMP"
const _CRTIMP_NOIA64 = "_CRTIMP"
const _CRTIMP_PURE = "_CRTIMP"
const _EMMINTRIN_H_INCLUDED = 1
const _FREEENTRY = 0
const _HEAP_MAXREQ = 0xFFFFFFFFFFFFFFE0
const _I16_MAX = 32767
const _I32_MAX = 2147483647
const _I64_MAX = "9223372036854775807ll"
const _I8_MAX = 127
const _IMMINTRIN_H_INCLUDED = 1
const _INTEGRAL_MAX_BITS = 64
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
const _OUT_TO_DEFAULT = 0
const _OUT_TO_MSGBOX = 2
const _OUT_TO_STDERR = 1
const _PMMINTRIN_H_INCLUDED = 1
const _REENTRANT = 1
const _REPORT_ERRMODE = 3
const _SECURECRT_FILL_BUFFER_PATTERN = 0xFD
const _UI16_MAX = "0xffffu"
const _UI32_MAX = "0xffffffffu"
const _UI64_MAX = "0xffffffffffffffffull"
const _UI8_MAX = "0xffu"
const _USEDENTRY = 1
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

const aux_sym_WORD_CHAR_token1 = 1
const aux_sym_PUNCTUATION_token1 = 2
const aux_sym_WS_token1 = 3
const aux_sym_NL_token1 = 4
const aux_sym_LINE_TAIL_token1 = 5
const aux_sym_COMMENT_PREFIX_token1 = 6
const anon_sym_AT = 7
const anon_sym_EQ = 8
const aux_sym__var_comment_token1 = 9
const aux_sym_request_separator_token1 = 10
const sym_method = 11
const aux_sym_http_version_token1 = 12
const sym_status_code = 13
const sym_status_text = 14
const anon_sym_COLON = 15
const anon_sym_LBRACE_LBRACE = 16
const anon_sym_RBRACE_RBRACE = 17
const anon_sym_LT = 18
const aux_sym_pre_request_script_token1 = 19
const anon_sym_GT = 20
const anon_sym_LBRACE_PERCENT = 21
const anon_sym_PERCENT_RBRACE = 22
const aux_sym_res_redirect_token1 = 23
const anon_sym_AT2 = 24
const aux_sym_xml_body_token1 = 25
const aux_sym_json_body_token1 = 26
const aux_sym_graphql_data_token1 = 27
const aux_sym_graphql_json_body_token1 = 28
const anon_sym_LT2 = 29
const anon_sym_DASH_DASH = 30
const aux_sym_multipart_form_data_token1 = 31
const aux_sym_multipart_form_data_token2 = 32
const aux_sym_raw_body_token1 = 33
const aux_sym__raw_body_token1 = 34
const sym__not_comment = 35
const sym_header_entity = 36
const sym_identifier = 37
const aux_sym_path_token1 = 38
const aux_sym__blank_line_token1 = 39
const sym_document = 40
const sym_comment = 41
const sym__plain_comment = 42
const sym__var_comment = 43
const sym_request_separator = 44
const sym_section = 45
const sym__section_content = 46
const sym_http_version = 47
const aux_sym__target_url_line = 48
const sym_target_url = 49
const sym_response = 50
const sym_request = 51
const sym_header = 52
const sym_variable = 53
const sym_pre_request_script = 54
const sym_res_handler_script = 55
const sym_script = 56
const sym_res_redirect = 57
const sym_variable_declaration = 58
const sym_xml_body = 59
const sym_json_body = 60
const sym_graphql_body = 61
const sym_graphql_data = 62
const sym_graphql_json_body = 63
const sym__external_body = 64
const sym_external_body = 65
const sym_multipart_form_data = 66
const sym_raw_body = 67
const sym__raw_body = 68
const sym_path = 69
const sym_value = 70
const sym__blank_line = 71
const aux_sym_document_repeat1 = 72
const aux_sym_target_url_repeat1 = 73
const aux_sym___body_repeat1 = 74
const aux_sym___body_repeat2 = 75
const aux_sym_response_repeat1 = 76
const aux_sym_request_repeat1 = 77
const aux_sym_script_repeat1 = 78
const aux_sym_multipart_form_data_repeat1 = 79
const aux_sym_path_repeat1 = 80
const aux_sym_value_repeat1 = 81

var ts_symbol_names = [82]uintptr{
	0:  __ccgo_ts,
	1:  __ccgo_ts + 4,
	2:  __ccgo_ts + 21,
	3:  __ccgo_ts + 40,
	4:  __ccgo_ts + 50,
	5:  __ccgo_ts + 60,
	6:  __ccgo_ts + 77,
	7:  __ccgo_ts + 99,
	8:  __ccgo_ts + 101,
	9:  __ccgo_ts + 103,
	10: __ccgo_ts + 123,
	11: __ccgo_ts + 148,
	12: __ccgo_ts + 155,
	13: __ccgo_ts + 175,
	14: __ccgo_ts + 187,
	15: __ccgo_ts + 199,
	16: __ccgo_ts + 201,
	17: __ccgo_ts + 204,
	18: __ccgo_ts + 207,
	19: __ccgo_ts + 209,
	20: __ccgo_ts + 235,
	21: __ccgo_ts + 237,
	22: __ccgo_ts + 240,
	23: __ccgo_ts + 243,
	24: __ccgo_ts + 99,
	25: __ccgo_ts + 263,
	26: __ccgo_ts + 279,
	27: __ccgo_ts + 296,
	28: __ccgo_ts + 316,
	29: __ccgo_ts + 207,
	30: __ccgo_ts + 341,
	31: __ccgo_ts + 344,
	32: __ccgo_ts + 371,
	33: __ccgo_ts + 398,
	34: __ccgo_ts + 414,
	35: __ccgo_ts + 431,
	36: __ccgo_ts + 444,
	37: __ccgo_ts + 458,
	38: __ccgo_ts + 469,
	39: __ccgo_ts + 481,
	40: __ccgo_ts + 500,
	41: __ccgo_ts + 509,
	42: __ccgo_ts + 517,
	43: __ccgo_ts + 532,
	44: __ccgo_ts + 545,
	45: __ccgo_ts + 563,
	46: __ccgo_ts + 571,
	47: __ccgo_ts + 588,
	48: __ccgo_ts + 601,
	49: __ccgo_ts + 618,
	50: __ccgo_ts + 629,
	51: __ccgo_ts + 638,
	52: __ccgo_ts + 646,
	53: __ccgo_ts + 653,
	54: __ccgo_ts + 662,
	55: __ccgo_ts + 681,
	56: __ccgo_ts + 700,
	57: __ccgo_ts + 707,
	58: __ccgo_ts + 720,
	59: __ccgo_ts + 741,
	60: __ccgo_ts + 750,
	61: __ccgo_ts + 760,
	62: __ccgo_ts + 773,
	63: __ccgo_ts + 750,
	64: __ccgo_ts + 786,
	65: __ccgo_ts + 801,
	66: __ccgo_ts + 815,
	67: __ccgo_ts + 835,
	68: __ccgo_ts + 844,
	69: __ccgo_ts + 854,
	70: __ccgo_ts + 859,
	71: __ccgo_ts + 865,
	72: __ccgo_ts + 877,
	73: __ccgo_ts + 894,
	74: __ccgo_ts + 913,
	75: __ccgo_ts + 928,
	76: __ccgo_ts + 943,
	77: __ccgo_ts + 960,
	78: __ccgo_ts + 976,
	79: __ccgo_ts + 991,
	80: __ccgo_ts + 1019,
	81: __ccgo_ts + 1032,
}

var ts_symbol_map = [82]TSSymbol{
	1:  uint16(aux_sym_WORD_CHAR_token1),
	2:  uint16(aux_sym_PUNCTUATION_token1),
	3:  uint16(aux_sym_WS_token1),
	4:  uint16(aux_sym_NL_token1),
	5:  uint16(aux_sym_LINE_TAIL_token1),
	6:  uint16(aux_sym_COMMENT_PREFIX_token1),
	7:  uint16(anon_sym_AT),
	8:  uint16(anon_sym_EQ),
	9:  uint16(aux_sym__var_comment_token1),
	10: uint16(aux_sym_request_separator_token1),
	11: uint16(sym_method),
	12: uint16(aux_sym_http_version_token1),
	13: uint16(sym_status_code),
	14: uint16(sym_status_text),
	15: uint16(anon_sym_COLON),
	16: uint16(anon_sym_LBRACE_LBRACE),
	17: uint16(anon_sym_RBRACE_RBRACE),
	18: uint16(anon_sym_LT),
	19: uint16(aux_sym_pre_request_script_token1),
	20: uint16(anon_sym_GT),
	21: uint16(anon_sym_LBRACE_PERCENT),
	22: uint16(anon_sym_PERCENT_RBRACE),
	23: uint16(aux_sym_res_redirect_token1),
	24: uint16(anon_sym_AT),
	25: uint16(aux_sym_xml_body_token1),
	26: uint16(aux_sym_json_body_token1),
	27: uint16(aux_sym_graphql_data_token1),
	28: uint16(aux_sym_graphql_json_body_token1),
	29: uint16(anon_sym_LT),
	30: uint16(anon_sym_DASH_DASH),
	31: uint16(aux_sym_multipart_form_data_token1),
	32: uint16(aux_sym_multipart_form_data_token2),
	33: uint16(aux_sym_raw_body_token1),
	34: uint16(aux_sym__raw_body_token1),
	35: uint16(sym__not_comment),
	36: uint16(sym_header_entity),
	37: uint16(sym_identifier),
	38: uint16(aux_sym_path_token1),
	39: uint16(aux_sym__blank_line_token1),
	40: uint16(sym_document),
	41: uint16(sym_comment),
	42: uint16(sym__plain_comment),
	43: uint16(sym__var_comment),
	44: uint16(sym_request_separator),
	45: uint16(sym_section),
	46: uint16(sym__section_content),
	47: uint16(sym_http_version),
	48: uint16(aux_sym__target_url_line),
	49: uint16(sym_target_url),
	50: uint16(sym_response),
	51: uint16(sym_request),
	52: uint16(sym_header),
	53: uint16(sym_variable),
	54: uint16(sym_pre_request_script),
	55: uint16(sym_res_handler_script),
	56: uint16(sym_script),
	57: uint16(sym_res_redirect),
	58: uint16(sym_variable_declaration),
	59: uint16(sym_xml_body),
	60: uint16(sym_json_body),
	61: uint16(sym_graphql_body),
	62: uint16(sym_graphql_data),
	63: uint16(sym_json_body),
	64: uint16(sym__external_body),
	65: uint16(sym_external_body),
	66: uint16(sym_multipart_form_data),
	67: uint16(sym_raw_body),
	68: uint16(sym__raw_body),
	69: uint16(sym_path),
	70: uint16(sym_value),
	71: uint16(sym__blank_line),
	72: uint16(aux_sym_document_repeat1),
	73: uint16(aux_sym_target_url_repeat1),
	74: uint16(aux_sym___body_repeat1),
	75: uint16(aux_sym___body_repeat2),
	76: uint16(aux_sym_response_repeat1),
	77: uint16(aux_sym_request_repeat1),
	78: uint16(aux_sym_script_repeat1),
	79: uint16(aux_sym_multipart_form_data_repeat1),
	80: uint16(aux_sym_path_repeat1),
	81: uint16(aux_sym_value_repeat1),
}

var ts_symbol_metadata = [82]TSSymbolMetadata{
	0: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	1: {},
	2: {},
	3: {},
	4: {},
	5: {},
	6: {},
	7: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	8: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	9:  {},
	10: {},
	11: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	12: {},
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
	19: {},
	20: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	21: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	22: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	23: {},
	24: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	25: {},
	26: {},
	27: {},
	28: {},
	29: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	30: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	31: {},
	32: {},
	33: {},
	34: {},
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
	38: {},
	39: {},
	40: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	41: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	42: {
		Fnamed: libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	47: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	48: {},
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
		Fnamed:   libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
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
	72: {},
	73: {},
	74: {},
	75: {},
	76: {},
	77: {},
	78: {},
	79: {},
	80: {},
	81: {},
}

type ts_field_identifiers = int32

const field_body = 1
const field_header = 2
const field_method = 3
const field_name = 4
const field_path = 5
const field_request = 6
const field_response = 7
const field_url = 8
const field_value = 9
const field_version = 10

var ts_field_names = [11]uintptr{
	0:  libc.UintptrFromInt32(0),
	1:  __ccgo_ts + 1046,
	2:  __ccgo_ts + 646,
	3:  __ccgo_ts + 148,
	4:  __ccgo_ts + 1051,
	5:  __ccgo_ts + 854,
	6:  __ccgo_ts + 638,
	7:  __ccgo_ts + 629,
	8:  __ccgo_ts + 1056,
	9:  __ccgo_ts + 859,
	10: __ccgo_ts + 1060,
}

var ts_field_map_slices = [49]TSFieldMapSlice{
	1: {
		Flength: uint16(2),
	},
	2: {
		Findex:  uint16(2),
		Flength: uint16(2),
	},
	3: {
		Findex:  uint16(4),
		Flength: uint16(1),
	},
	4: {
		Findex:  uint16(5),
		Flength: uint16(1),
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
		Flength: uint16(1),
	},
	8: {
		Findex:  uint16(10),
		Flength: uint16(1),
	},
	9: {
		Findex:  uint16(11),
		Flength: uint16(1),
	},
	10: {
		Findex:  uint16(12),
		Flength: uint16(2),
	},
	11: {
		Findex:  uint16(14),
		Flength: uint16(1),
	},
	12: {
		Findex:  uint16(15),
		Flength: uint16(1),
	},
	13: {
		Findex:  uint16(16),
		Flength: uint16(2),
	},
	14: {
		Findex:  uint16(18),
		Flength: uint16(2),
	},
	15: {
		Flength: uint16(2),
	},
	16: {
		Findex:  uint16(20),
		Flength: uint16(1),
	},
	17: {
		Findex:  uint16(21),
		Flength: uint16(2),
	},
	18: {
		Findex:  uint16(23),
		Flength: uint16(2),
	},
	19: {
		Findex:  uint16(25),
		Flength: uint16(3),
	},
	20: {
		Findex:  uint16(28),
		Flength: uint16(2),
	},
	21: {
		Findex:  uint16(30),
		Flength: uint16(3),
	},
	22: {
		Findex:  uint16(33),
		Flength: uint16(1),
	},
	23: {
		Findex:  uint16(34),
		Flength: uint16(2),
	},
	24: {
		Findex:  uint16(36),
		Flength: uint16(3),
	},
	25: {
		Findex:  uint16(39),
		Flength: uint16(2),
	},
	26: {
		Findex:  uint16(41),
		Flength: uint16(3),
	},
	27: {
		Findex:  uint16(44),
		Flength: uint16(3),
	},
	28: {
		Findex:  uint16(47),
		Flength: uint16(2),
	},
	29: {
		Findex:  uint16(49),
		Flength: uint16(1),
	},
	30: {
		Findex:  uint16(50),
		Flength: uint16(3),
	},
	31: {
		Findex:  uint16(53),
		Flength: uint16(2),
	},
	32: {
		Findex:  uint16(55),
		Flength: uint16(1),
	},
	33: {
		Findex:  uint16(56),
		Flength: uint16(2),
	},
	34: {
		Findex:  uint16(58),
		Flength: uint16(4),
	},
	35: {
		Findex:  uint16(62),
		Flength: uint16(4),
	},
	36: {
		Findex:  uint16(66),
		Flength: uint16(2),
	},
	37: {
		Findex:  uint16(68),
		Flength: uint16(1),
	},
	38: {
		Findex:  uint16(69),
		Flength: uint16(1),
	},
	39: {
		Findex:  uint16(70),
		Flength: uint16(4),
	},
	40: {
		Findex:  uint16(74),
		Flength: uint16(2),
	},
	41: {
		Findex:  uint16(76),
		Flength: uint16(1),
	},
	42: {
		Findex:  uint16(77),
		Flength: uint16(4),
	},
	43: {
		Findex:  uint16(81),
		Flength: uint16(2),
	},
	44: {
		Findex:  uint16(83),
		Flength: uint16(1),
	},
	45: {
		Findex:  uint16(84),
		Flength: uint16(2),
	},
	46: {
		Findex:  uint16(86),
		Flength: uint16(2),
	},
	47: {
		Findex:  uint16(88),
		Flength: uint16(5),
	},
	48: {
		Findex:  uint16(93),
		Flength: uint16(2),
	},
}

var ts_field_map_entries = [95]TSFieldMapEntry{
	0: {
		Ffield_id:  uint16(field_name),
		Finherited: libc.BoolUint8(true1 != 0),
	},
	1: {
		Ffield_id:  uint16(field_value),
		Finherited: libc.BoolUint8(true1 != 0),
	},
	2: {
		Ffield_id:  uint16(field_request),
		Finherited: libc.BoolUint8(true1 != 0),
	},
	3: {
		Ffield_id:  uint16(field_response),
		Finherited: libc.BoolUint8(true1 != 0),
	},
	4: {
		Ffield_id: uint16(field_response),
	},
	5: {
		Ffield_id: uint16(field_request),
	},
	6: {
		Ffield_id:    uint16(field_request),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	7: {
		Ffield_id:    uint16(field_response),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	8: {
		Ffield_id: uint16(field_url),
	},
	9: {
		Ffield_id:    uint16(field_value),
		Fchild_index: uint8(1),
	},
	10: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
	},
	11: {
		Ffield_id: uint16(field_header),
	},
	12: {
		Ffield_id:    uint16(field_header),
		Fchild_index: uint8(2),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	13: {
		Ffield_id: uint16(field_url),
	},
	14: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(2),
	},
	15: {
		Ffield_id:    uint16(field_value),
		Fchild_index: uint8(2),
	},
	16: {
		Ffield_id: uint16(field_method),
	},
	17: {
		Ffield_id:    uint16(field_url),
		Fchild_index: uint8(2),
	},
	18: {
		Ffield_id: uint16(field_url),
	},
	19: {
		Ffield_id:    uint16(field_version),
		Fchild_index: uint8(2),
	},
	20: {
		Ffield_id: uint16(field_body),
	},
	21: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(3),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	22: {
		Ffield_id: uint16(field_url),
	},
	23: {
		Ffield_id:  uint16(field_header),
		Finherited: libc.BoolUint8(true1 != 0),
	},
	24: {
		Ffield_id:    uint16(field_header),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	25: {
		Ffield_id:    uint16(field_header),
		Fchild_index: uint8(4),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	26: {
		Ffield_id: uint16(field_method),
	},
	27: {
		Ffield_id:    uint16(field_url),
		Fchild_index: uint8(2),
	},
	28: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
	},
	29: {
		Ffield_id:    uint16(field_value),
		Fchild_index: uint8(3),
	},
	30: {
		Ffield_id:    uint16(field_header),
		Fchild_index: uint8(4),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	31: {
		Ffield_id: uint16(field_url),
	},
	32: {
		Ffield_id:    uint16(field_version),
		Fchild_index: uint8(2),
	},
	33: {
		Ffield_id: uint16(field_name),
	},
	34: {
		Ffield_id:  uint16(field_body),
		Finherited: libc.BoolUint8(true1 != 0),
	},
	35: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	36: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(4),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	37: {
		Ffield_id:    uint16(field_header),
		Fchild_index: uint8(2),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	38: {
		Ffield_id: uint16(field_url),
	},
	39: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(2),
	},
	40: {
		Ffield_id:    uint16(field_value),
		Fchild_index: uint8(4),
	},
	41: {
		Ffield_id: uint16(field_method),
	},
	42: {
		Ffield_id:    uint16(field_url),
		Fchild_index: uint8(2),
	},
	43: {
		Ffield_id:    uint16(field_version),
		Fchild_index: uint8(4),
	},
	44: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(5),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	45: {
		Ffield_id: uint16(field_method),
	},
	46: {
		Ffield_id:    uint16(field_url),
		Fchild_index: uint8(2),
	},
	47: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
	},
	48: {
		Ffield_id:    uint16(field_value),
		Fchild_index: uint8(4),
	},
	49: {
		Ffield_id:    uint16(field_header),
		Fchild_index: uint8(5),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	50: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(5),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	51: {
		Ffield_id: uint16(field_url),
	},
	52: {
		Ffield_id:    uint16(field_version),
		Fchild_index: uint8(2),
	},
	53: {
		Ffield_id: uint16(field_name),
	},
	54: {
		Ffield_id:    uint16(field_value),
		Fchild_index: uint8(2),
	},
	55: {
		Ffield_id:    uint16(field_path),
		Fchild_index: uint8(2),
	},
	56: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(2),
	},
	57: {
		Ffield_id:    uint16(field_value),
		Fchild_index: uint8(5),
	},
	58: {
		Ffield_id:    uint16(field_header),
		Fchild_index: uint8(6),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	59: {
		Ffield_id: uint16(field_method),
	},
	60: {
		Ffield_id:    uint16(field_url),
		Fchild_index: uint8(2),
	},
	61: {
		Ffield_id:    uint16(field_version),
		Fchild_index: uint8(4),
	},
	62: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(6),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	63: {
		Ffield_id:    uint16(field_header),
		Fchild_index: uint8(4),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	64: {
		Ffield_id: uint16(field_method),
	},
	65: {
		Ffield_id:    uint16(field_url),
		Fchild_index: uint8(2),
	},
	66: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
	},
	67: {
		Ffield_id:    uint16(field_value),
		Fchild_index: uint8(5),
	},
	68: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(6),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	69: {
		Ffield_id:    uint16(field_header),
		Fchild_index: uint8(6),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	70: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(6),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	71: {
		Ffield_id:    uint16(field_header),
		Fchild_index: uint8(4),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	72: {
		Ffield_id: uint16(field_url),
	},
	73: {
		Ffield_id:    uint16(field_version),
		Fchild_index: uint8(2),
	},
	74: {
		Ffield_id: uint16(field_name),
	},
	75: {
		Ffield_id:    uint16(field_value),
		Fchild_index: uint8(3),
	},
	76: {
		Ffield_id:    uint16(field_path),
		Fchild_index: uint8(3),
	},
	77: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(7),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	78: {
		Ffield_id: uint16(field_method),
	},
	79: {
		Ffield_id:    uint16(field_url),
		Fchild_index: uint8(2),
	},
	80: {
		Ffield_id:    uint16(field_version),
		Fchild_index: uint8(4),
	},
	81: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(7),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	82: {
		Ffield_id:    uint16(field_header),
		Fchild_index: uint8(5),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	83: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(7),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	84: {
		Ffield_id: uint16(field_name),
	},
	85: {
		Ffield_id:    uint16(field_value),
		Fchild_index: uint8(4),
	},
	86: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(2),
	},
	87: {
		Ffield_id:    uint16(field_path),
		Fchild_index: uint8(4),
	},
	88: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(8),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	89: {
		Ffield_id:    uint16(field_header),
		Fchild_index: uint8(6),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	90: {
		Ffield_id: uint16(field_method),
	},
	91: {
		Ffield_id:    uint16(field_url),
		Fchild_index: uint8(2),
	},
	92: {
		Ffield_id:    uint16(field_version),
		Fchild_index: uint8(4),
	},
	93: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(8),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	94: {
		Ffield_id:    uint16(field_header),
		Fchild_index: uint8(6),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
}

var ts_alias_sequences = [49][9]TSSymbol{
	0: {},
	15: {
		0: uint16(sym_comment),
	},
}

var ts_non_terminal_alias_map = [5]uint16_t{
	0: uint16(sym__var_comment),
	1: uint16(2),
	2: uint16(sym__var_comment),
	3: uint16(sym_comment),
}

var ts_primary_state_ids = [261]TSStateId{
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
	37:  uint16(35),
	38:  uint16(38),
	39:  uint16(39),
	40:  uint16(40),
	41:  uint16(41),
	42:  uint16(34),
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
	55:  uint16(46),
	56:  uint16(45),
	57:  uint16(57),
	58:  uint16(58),
	59:  uint16(59),
	60:  uint16(60),
	61:  uint16(61),
	62:  uint16(44),
	63:  uint16(63),
	64:  uint16(51),
	65:  uint16(65),
	66:  uint16(66),
	67:  uint16(50),
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
	84:  uint16(36),
	85:  uint16(85),
	86:  uint16(86),
	87:  uint16(87),
	88:  uint16(88),
	89:  uint16(89),
	90:  uint16(90),
	91:  uint16(91),
	92:  uint16(39),
	93:  uint16(44),
	94:  uint16(45),
	95:  uint16(38),
	96:  uint16(46),
	97:  uint16(97),
	98:  uint16(98),
	99:  uint16(99),
	100: uint16(100),
	101: uint16(44),
	102: uint16(46),
	103: uint16(103),
	104: uint16(104),
	105: uint16(105),
	106: uint16(57),
	107: uint16(36),
	108: uint16(108),
	109: uint16(45),
	110: uint16(110),
	111: uint16(111),
	112: uint16(38),
	113: uint16(113),
	114: uint16(114),
	115: uint16(115),
	116: uint16(39),
	117: uint16(117),
	118: uint16(118),
	119: uint16(119),
	120: uint16(120),
	121: uint16(121),
	122: uint16(122),
	123: uint16(123),
	124: uint16(124),
	125: uint16(124),
	126: uint16(126),
	127: uint16(127),
	128: uint16(124),
	129: uint16(129),
	130: uint16(130),
	131: uint16(131),
	132: uint16(132),
	133: uint16(124),
	134: uint16(134),
	135: uint16(135),
	136: uint16(136),
	137: uint16(137),
	138: uint16(136),
	139: uint16(139),
	140: uint16(136),
	141: uint16(130),
	142: uint16(136),
	143: uint16(143),
	144: uint16(144),
	145: uint16(145),
	146: uint16(146),
	147: uint16(147),
	148: uint16(148),
	149: uint16(149),
	150: uint16(129),
	151: uint16(151),
	152: uint16(152),
	153: uint16(153),
	154: uint16(154),
	155: uint16(155),
	156: uint16(156),
	157: uint16(152),
	158: uint16(153),
	159: uint16(154),
	160: uint16(151),
	161: uint16(152),
	162: uint16(153),
	163: uint16(154),
	164: uint16(151),
	165: uint16(165),
	166: uint16(166),
	167: uint16(167),
	168: uint16(168),
	169: uint16(169),
	170: uint16(170),
	171: uint16(170),
	172: uint16(170),
	173: uint16(170),
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
	184: uint16(183),
	185: uint16(185),
	186: uint16(186),
	187: uint16(187),
	188: uint16(178),
	189: uint16(187),
	190: uint16(186),
	191: uint16(191),
	192: uint16(187),
	193: uint16(178),
	194: uint16(194),
	195: uint16(183),
	196: uint16(196),
	197: uint16(197),
	198: uint16(198),
	199: uint16(199),
	200: uint16(200),
	201: uint16(201),
	202: uint16(202),
	203: uint16(186),
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
	218: uint16(206),
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
	234: uint16(232),
	235: uint16(235),
	236: uint16(236),
	237: uint16(237),
	238: uint16(238),
	239: uint16(239),
	240: uint16(240),
	241: uint16(206),
	242: uint16(233),
	243: uint16(232),
	244: uint16(236),
	245: uint16(236),
	246: uint16(232),
	247: uint16(236),
	248: uint16(248),
	249: uint16(227),
	250: uint16(237),
	251: uint16(251),
	252: uint16(252),
	253: uint16(233),
	254: uint16(227),
	255: uint16(237),
	256: uint16(256),
	257: uint16(257),
	258: uint16(227),
	259: uint16(259),
	260: uint16(235),
}

var aux_sym_WORD_CHAR_token1_character_set_1 = [475]TSCharacterRange{
	0: {
		Fstart: int32('0'),
		Fend:   int32('9'),
	},
	1: {
		Fstart: int32('A'),
		Fend:   int32('Z'),
	},
	2: {
		Fstart: int32('a'),
		Fend:   int32('z'),
	},
	3: {
		Fstart: int32(0xaa),
		Fend:   int32(0xaa),
	},
	4: {
		Fstart: int32(0xb2),
		Fend:   int32(0xb3),
	},
	5: {
		Fstart: int32(0xb5),
		Fend:   int32(0xb5),
	},
	6: {
		Fstart: int32(0xb9),
		Fend:   int32(0xba),
	},
	7: {
		Fstart: int32(0xbc),
		Fend:   int32(0xbe),
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
		Fstart: int32(0x660),
		Fend:   int32(0x669),
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
		Fstart: int32(0x7c0),
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
		Fstart: int32(0x966),
		Fend:   int32(0x96f),
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
		Fstart: int32(0x9e6),
		Fend:   int32(0x9f1),
	},
	72: {
		Fstart: int32(0x9f4),
		Fend:   int32(0x9f9),
	},
	73: {
		Fstart: int32(0x9fc),
		Fend:   int32(0x9fc),
	},
	74: {
		Fstart: int32(0xa05),
		Fend:   int32(0xa0a),
	},
	75: {
		Fstart: int32(0xa0f),
		Fend:   int32(0xa10),
	},
	76: {
		Fstart: int32(0xa13),
		Fend:   int32(0xa28),
	},
	77: {
		Fstart: int32(0xa2a),
		Fend:   int32(0xa30),
	},
	78: {
		Fstart: int32(0xa32),
		Fend:   int32(0xa33),
	},
	79: {
		Fstart: int32(0xa35),
		Fend:   int32(0xa36),
	},
	80: {
		Fstart: int32(0xa38),
		Fend:   int32(0xa39),
	},
	81: {
		Fstart: int32(0xa59),
		Fend:   int32(0xa5c),
	},
	82: {
		Fstart: int32(0xa5e),
		Fend:   int32(0xa5e),
	},
	83: {
		Fstart: int32(0xa66),
		Fend:   int32(0xa6f),
	},
	84: {
		Fstart: int32(0xa72),
		Fend:   int32(0xa74),
	},
	85: {
		Fstart: int32(0xa85),
		Fend:   int32(0xa8d),
	},
	86: {
		Fstart: int32(0xa8f),
		Fend:   int32(0xa91),
	},
	87: {
		Fstart: int32(0xa93),
		Fend:   int32(0xaa8),
	},
	88: {
		Fstart: int32(0xaaa),
		Fend:   int32(0xab0),
	},
	89: {
		Fstart: int32(0xab2),
		Fend:   int32(0xab3),
	},
	90: {
		Fstart: int32(0xab5),
		Fend:   int32(0xab9),
	},
	91: {
		Fstart: int32(0xabd),
		Fend:   int32(0xabd),
	},
	92: {
		Fstart: int32(0xad0),
		Fend:   int32(0xad0),
	},
	93: {
		Fstart: int32(0xae0),
		Fend:   int32(0xae1),
	},
	94: {
		Fstart: int32(0xae6),
		Fend:   int32(0xaef),
	},
	95: {
		Fstart: int32(0xaf9),
		Fend:   int32(0xaf9),
	},
	96: {
		Fstart: int32(0xb05),
		Fend:   int32(0xb0c),
	},
	97: {
		Fstart: int32(0xb0f),
		Fend:   int32(0xb10),
	},
	98: {
		Fstart: int32(0xb13),
		Fend:   int32(0xb28),
	},
	99: {
		Fstart: int32(0xb2a),
		Fend:   int32(0xb30),
	},
	100: {
		Fstart: int32(0xb32),
		Fend:   int32(0xb33),
	},
	101: {
		Fstart: int32(0xb35),
		Fend:   int32(0xb39),
	},
	102: {
		Fstart: int32(0xb3d),
		Fend:   int32(0xb3d),
	},
	103: {
		Fstart: int32(0xb5c),
		Fend:   int32(0xb5d),
	},
	104: {
		Fstart: int32(0xb5f),
		Fend:   int32(0xb61),
	},
	105: {
		Fstart: int32(0xb66),
		Fend:   int32(0xb6f),
	},
	106: {
		Fstart: int32(0xb71),
		Fend:   int32(0xb77),
	},
	107: {
		Fstart: int32(0xb83),
		Fend:   int32(0xb83),
	},
	108: {
		Fstart: int32(0xb85),
		Fend:   int32(0xb8a),
	},
	109: {
		Fstart: int32(0xb8e),
		Fend:   int32(0xb90),
	},
	110: {
		Fstart: int32(0xb92),
		Fend:   int32(0xb95),
	},
	111: {
		Fstart: int32(0xb99),
		Fend:   int32(0xb9a),
	},
	112: {
		Fstart: int32(0xb9c),
		Fend:   int32(0xb9c),
	},
	113: {
		Fstart: int32(0xb9e),
		Fend:   int32(0xb9f),
	},
	114: {
		Fstart: int32(0xba3),
		Fend:   int32(0xba4),
	},
	115: {
		Fstart: int32(0xba8),
		Fend:   int32(0xbaa),
	},
	116: {
		Fstart: int32(0xbae),
		Fend:   int32(0xbb9),
	},
	117: {
		Fstart: int32(0xbd0),
		Fend:   int32(0xbd0),
	},
	118: {
		Fstart: int32(0xbe6),
		Fend:   int32(0xbf2),
	},
	119: {
		Fstart: int32(0xc05),
		Fend:   int32(0xc0c),
	},
	120: {
		Fstart: int32(0xc0e),
		Fend:   int32(0xc10),
	},
	121: {
		Fstart: int32(0xc12),
		Fend:   int32(0xc28),
	},
	122: {
		Fstart: int32(0xc2a),
		Fend:   int32(0xc39),
	},
	123: {
		Fstart: int32(0xc3d),
		Fend:   int32(0xc3d),
	},
	124: {
		Fstart: int32(0xc58),
		Fend:   int32(0xc5a),
	},
	125: {
		Fstart: int32(0xc5d),
		Fend:   int32(0xc5d),
	},
	126: {
		Fstart: int32(0xc60),
		Fend:   int32(0xc61),
	},
	127: {
		Fstart: int32(0xc66),
		Fend:   int32(0xc6f),
	},
	128: {
		Fstart: int32(0xc78),
		Fend:   int32(0xc7e),
	},
	129: {
		Fstart: int32(0xc80),
		Fend:   int32(0xc80),
	},
	130: {
		Fstart: int32(0xc85),
		Fend:   int32(0xc8c),
	},
	131: {
		Fstart: int32(0xc8e),
		Fend:   int32(0xc90),
	},
	132: {
		Fstart: int32(0xc92),
		Fend:   int32(0xca8),
	},
	133: {
		Fstart: int32(0xcaa),
		Fend:   int32(0xcb3),
	},
	134: {
		Fstart: int32(0xcb5),
		Fend:   int32(0xcb9),
	},
	135: {
		Fstart: int32(0xcbd),
		Fend:   int32(0xcbd),
	},
	136: {
		Fstart: int32(0xcdd),
		Fend:   int32(0xcde),
	},
	137: {
		Fstart: int32(0xce0),
		Fend:   int32(0xce1),
	},
	138: {
		Fstart: int32(0xce6),
		Fend:   int32(0xcef),
	},
	139: {
		Fstart: int32(0xcf1),
		Fend:   int32(0xcf2),
	},
	140: {
		Fstart: int32(0xd04),
		Fend:   int32(0xd0c),
	},
	141: {
		Fstart: int32(0xd0e),
		Fend:   int32(0xd10),
	},
	142: {
		Fstart: int32(0xd12),
		Fend:   int32(0xd3a),
	},
	143: {
		Fstart: int32(0xd3d),
		Fend:   int32(0xd3d),
	},
	144: {
		Fstart: int32(0xd4e),
		Fend:   int32(0xd4e),
	},
	145: {
		Fstart: int32(0xd54),
		Fend:   int32(0xd56),
	},
	146: {
		Fstart: int32(0xd58),
		Fend:   int32(0xd61),
	},
	147: {
		Fstart: int32(0xd66),
		Fend:   int32(0xd78),
	},
	148: {
		Fstart: int32(0xd7a),
		Fend:   int32(0xd7f),
	},
	149: {
		Fstart: int32(0xd85),
		Fend:   int32(0xd96),
	},
	150: {
		Fstart: int32(0xd9a),
		Fend:   int32(0xdb1),
	},
	151: {
		Fstart: int32(0xdb3),
		Fend:   int32(0xdbb),
	},
	152: {
		Fstart: int32(0xdbd),
		Fend:   int32(0xdbd),
	},
	153: {
		Fstart: int32(0xdc0),
		Fend:   int32(0xdc6),
	},
	154: {
		Fstart: int32(0xde6),
		Fend:   int32(0xdef),
	},
	155: {
		Fstart: int32(0xe01),
		Fend:   int32(0xe30),
	},
	156: {
		Fstart: int32(0xe32),
		Fend:   int32(0xe33),
	},
	157: {
		Fstart: int32(0xe40),
		Fend:   int32(0xe46),
	},
	158: {
		Fstart: int32(0xe50),
		Fend:   int32(0xe59),
	},
	159: {
		Fstart: int32(0xe81),
		Fend:   int32(0xe82),
	},
	160: {
		Fstart: int32(0xe84),
		Fend:   int32(0xe84),
	},
	161: {
		Fstart: int32(0xe86),
		Fend:   int32(0xe8a),
	},
	162: {
		Fstart: int32(0xe8c),
		Fend:   int32(0xea3),
	},
	163: {
		Fstart: int32(0xea5),
		Fend:   int32(0xea5),
	},
	164: {
		Fstart: int32(0xea7),
		Fend:   int32(0xeb0),
	},
	165: {
		Fstart: int32(0xeb2),
		Fend:   int32(0xeb3),
	},
	166: {
		Fstart: int32(0xebd),
		Fend:   int32(0xebd),
	},
	167: {
		Fstart: int32(0xec0),
		Fend:   int32(0xec4),
	},
	168: {
		Fstart: int32(0xec6),
		Fend:   int32(0xec6),
	},
	169: {
		Fstart: int32(0xed0),
		Fend:   int32(0xed9),
	},
	170: {
		Fstart: int32(0xedc),
		Fend:   int32(0xedf),
	},
	171: {
		Fstart: int32(0xf00),
		Fend:   int32(0xf00),
	},
	172: {
		Fstart: int32(0xf20),
		Fend:   int32(0xf33),
	},
	173: {
		Fstart: int32(0xf40),
		Fend:   int32(0xf47),
	},
	174: {
		Fstart: int32(0xf49),
		Fend:   int32(0xf6c),
	},
	175: {
		Fstart: int32(0xf88),
		Fend:   int32(0xf8c),
	},
	176: {
		Fstart: int32(0x1000),
		Fend:   int32(0x102a),
	},
	177: {
		Fstart: int32(0x103f),
		Fend:   int32(0x1049),
	},
	178: {
		Fstart: int32(0x1050),
		Fend:   int32(0x1055),
	},
	179: {
		Fstart: int32(0x105a),
		Fend:   int32(0x105d),
	},
	180: {
		Fstart: int32(0x1061),
		Fend:   int32(0x1061),
	},
	181: {
		Fstart: int32(0x1065),
		Fend:   int32(0x1066),
	},
	182: {
		Fstart: int32(0x106e),
		Fend:   int32(0x1070),
	},
	183: {
		Fstart: int32(0x1075),
		Fend:   int32(0x1081),
	},
	184: {
		Fstart: int32(0x108e),
		Fend:   int32(0x108e),
	},
	185: {
		Fstart: int32(0x1090),
		Fend:   int32(0x1099),
	},
	186: {
		Fstart: int32(0x10a0),
		Fend:   int32(0x10c5),
	},
	187: {
		Fstart: int32(0x10c7),
		Fend:   int32(0x10c7),
	},
	188: {
		Fstart: int32(0x10cd),
		Fend:   int32(0x10cd),
	},
	189: {
		Fstart: int32(0x10d0),
		Fend:   int32(0x10fa),
	},
	190: {
		Fstart: int32(0x10fc),
		Fend:   int32(0x1248),
	},
	191: {
		Fstart: int32(0x124a),
		Fend:   int32(0x124d),
	},
	192: {
		Fstart: int32(0x1250),
		Fend:   int32(0x1256),
	},
	193: {
		Fstart: int32(0x1258),
		Fend:   int32(0x1258),
	},
	194: {
		Fstart: int32(0x125a),
		Fend:   int32(0x125d),
	},
	195: {
		Fstart: int32(0x1260),
		Fend:   int32(0x1288),
	},
	196: {
		Fstart: int32(0x128a),
		Fend:   int32(0x128d),
	},
	197: {
		Fstart: int32(0x1290),
		Fend:   int32(0x12b0),
	},
	198: {
		Fstart: int32(0x12b2),
		Fend:   int32(0x12b5),
	},
	199: {
		Fstart: int32(0x12b8),
		Fend:   int32(0x12be),
	},
	200: {
		Fstart: int32(0x12c0),
		Fend:   int32(0x12c0),
	},
	201: {
		Fstart: int32(0x12c2),
		Fend:   int32(0x12c5),
	},
	202: {
		Fstart: int32(0x12c8),
		Fend:   int32(0x12d6),
	},
	203: {
		Fstart: int32(0x12d8),
		Fend:   int32(0x1310),
	},
	204: {
		Fstart: int32(0x1312),
		Fend:   int32(0x1315),
	},
	205: {
		Fstart: int32(0x1318),
		Fend:   int32(0x135a),
	},
	206: {
		Fstart: int32(0x1369),
		Fend:   int32(0x137c),
	},
	207: {
		Fstart: int32(0x1380),
		Fend:   int32(0x138f),
	},
	208: {
		Fstart: int32(0x13a0),
		Fend:   int32(0x13f5),
	},
	209: {
		Fstart: int32(0x13f8),
		Fend:   int32(0x13fd),
	},
	210: {
		Fstart: int32(0x1401),
		Fend:   int32(0x166c),
	},
	211: {
		Fstart: int32(0x166f),
		Fend:   int32(0x167f),
	},
	212: {
		Fstart: int32(0x1681),
		Fend:   int32(0x169a),
	},
	213: {
		Fstart: int32(0x16a0),
		Fend:   int32(0x16ea),
	},
	214: {
		Fstart: int32(0x16ee),
		Fend:   int32(0x16f8),
	},
	215: {
		Fstart: int32(0x1700),
		Fend:   int32(0x1711),
	},
	216: {
		Fstart: int32(0x171f),
		Fend:   int32(0x1731),
	},
	217: {
		Fstart: int32(0x1740),
		Fend:   int32(0x1751),
	},
	218: {
		Fstart: int32(0x1760),
		Fend:   int32(0x176c),
	},
	219: {
		Fstart: int32(0x176e),
		Fend:   int32(0x1770),
	},
	220: {
		Fstart: int32(0x1780),
		Fend:   int32(0x17b3),
	},
	221: {
		Fstart: int32(0x17d7),
		Fend:   int32(0x17d7),
	},
	222: {
		Fstart: int32(0x17dc),
		Fend:   int32(0x17dc),
	},
	223: {
		Fstart: int32(0x17e0),
		Fend:   int32(0x17e9),
	},
	224: {
		Fstart: int32(0x17f0),
		Fend:   int32(0x17f9),
	},
	225: {
		Fstart: int32(0x1810),
		Fend:   int32(0x1819),
	},
	226: {
		Fstart: int32(0x1820),
		Fend:   int32(0x1878),
	},
	227: {
		Fstart: int32(0x1880),
		Fend:   int32(0x1884),
	},
	228: {
		Fstart: int32(0x1887),
		Fend:   int32(0x18a8),
	},
	229: {
		Fstart: int32(0x18aa),
		Fend:   int32(0x18aa),
	},
	230: {
		Fstart: int32(0x18b0),
		Fend:   int32(0x18f5),
	},
	231: {
		Fstart: int32(0x1900),
		Fend:   int32(0x191e),
	},
	232: {
		Fstart: int32(0x1946),
		Fend:   int32(0x196d),
	},
	233: {
		Fstart: int32(0x1970),
		Fend:   int32(0x1974),
	},
	234: {
		Fstart: int32(0x1980),
		Fend:   int32(0x19ab),
	},
	235: {
		Fstart: int32(0x19b0),
		Fend:   int32(0x19c9),
	},
	236: {
		Fstart: int32(0x19d0),
		Fend:   int32(0x19da),
	},
	237: {
		Fstart: int32(0x1a00),
		Fend:   int32(0x1a16),
	},
	238: {
		Fstart: int32(0x1a20),
		Fend:   int32(0x1a54),
	},
	239: {
		Fstart: int32(0x1a80),
		Fend:   int32(0x1a89),
	},
	240: {
		Fstart: int32(0x1a90),
		Fend:   int32(0x1a99),
	},
	241: {
		Fstart: int32(0x1aa7),
		Fend:   int32(0x1aa7),
	},
	242: {
		Fstart: int32(0x1b05),
		Fend:   int32(0x1b33),
	},
	243: {
		Fstart: int32(0x1b45),
		Fend:   int32(0x1b4c),
	},
	244: {
		Fstart: int32(0x1b50),
		Fend:   int32(0x1b59),
	},
	245: {
		Fstart: int32(0x1b83),
		Fend:   int32(0x1ba0),
	},
	246: {
		Fstart: int32(0x1bae),
		Fend:   int32(0x1be5),
	},
	247: {
		Fstart: int32(0x1c00),
		Fend:   int32(0x1c23),
	},
	248: {
		Fstart: int32(0x1c40),
		Fend:   int32(0x1c49),
	},
	249: {
		Fstart: int32(0x1c4d),
		Fend:   int32(0x1c7d),
	},
	250: {
		Fstart: int32(0x1c80),
		Fend:   int32(0x1c88),
	},
	251: {
		Fstart: int32(0x1c90),
		Fend:   int32(0x1cba),
	},
	252: {
		Fstart: int32(0x1cbd),
		Fend:   int32(0x1cbf),
	},
	253: {
		Fstart: int32(0x1ce9),
		Fend:   int32(0x1cec),
	},
	254: {
		Fstart: int32(0x1cee),
		Fend:   int32(0x1cf3),
	},
	255: {
		Fstart: int32(0x1cf5),
		Fend:   int32(0x1cf6),
	},
	256: {
		Fstart: int32(0x1cfa),
		Fend:   int32(0x1cfa),
	},
	257: {
		Fstart: int32(0x1d00),
		Fend:   int32(0x1dbf),
	},
	258: {
		Fstart: int32(0x1e00),
		Fend:   int32(0x1f15),
	},
	259: {
		Fstart: int32(0x1f18),
		Fend:   int32(0x1f1d),
	},
	260: {
		Fstart: int32(0x1f20),
		Fend:   int32(0x1f45),
	},
	261: {
		Fstart: int32(0x1f48),
		Fend:   int32(0x1f4d),
	},
	262: {
		Fstart: int32(0x1f50),
		Fend:   int32(0x1f57),
	},
	263: {
		Fstart: int32(0x1f59),
		Fend:   int32(0x1f59),
	},
	264: {
		Fstart: int32(0x1f5b),
		Fend:   int32(0x1f5b),
	},
	265: {
		Fstart: int32(0x1f5d),
		Fend:   int32(0x1f5d),
	},
	266: {
		Fstart: int32(0x1f5f),
		Fend:   int32(0x1f7d),
	},
	267: {
		Fstart: int32(0x1f80),
		Fend:   int32(0x1fb4),
	},
	268: {
		Fstart: int32(0x1fb6),
		Fend:   int32(0x1fbc),
	},
	269: {
		Fstart: int32(0x1fbe),
		Fend:   int32(0x1fbe),
	},
	270: {
		Fstart: int32(0x1fc2),
		Fend:   int32(0x1fc4),
	},
	271: {
		Fstart: int32(0x1fc6),
		Fend:   int32(0x1fcc),
	},
	272: {
		Fstart: int32(0x1fd0),
		Fend:   int32(0x1fd3),
	},
	273: {
		Fstart: int32(0x1fd6),
		Fend:   int32(0x1fdb),
	},
	274: {
		Fstart: int32(0x1fe0),
		Fend:   int32(0x1fec),
	},
	275: {
		Fstart: int32(0x1ff2),
		Fend:   int32(0x1ff4),
	},
	276: {
		Fstart: int32(0x1ff6),
		Fend:   int32(0x1ffc),
	},
	277: {
		Fstart: int32(0x2070),
		Fend:   int32(0x2071),
	},
	278: {
		Fstart: int32(0x2074),
		Fend:   int32(0x2079),
	},
	279: {
		Fstart: int32(0x207f),
		Fend:   int32(0x2089),
	},
	280: {
		Fstart: int32(0x2090),
		Fend:   int32(0x209c),
	},
	281: {
		Fstart: int32(0x2102),
		Fend:   int32(0x2102),
	},
	282: {
		Fstart: int32(0x2107),
		Fend:   int32(0x2107),
	},
	283: {
		Fstart: int32(0x210a),
		Fend:   int32(0x2113),
	},
	284: {
		Fstart: int32(0x2115),
		Fend:   int32(0x2115),
	},
	285: {
		Fstart: int32(0x2119),
		Fend:   int32(0x211d),
	},
	286: {
		Fstart: int32(0x2124),
		Fend:   int32(0x2124),
	},
	287: {
		Fstart: int32(0x2126),
		Fend:   int32(0x2126),
	},
	288: {
		Fstart: int32(0x2128),
		Fend:   int32(0x2128),
	},
	289: {
		Fstart: int32(0x212a),
		Fend:   int32(0x212d),
	},
	290: {
		Fstart: int32(0x212f),
		Fend:   int32(0x2139),
	},
	291: {
		Fstart: int32(0x213c),
		Fend:   int32(0x213f),
	},
	292: {
		Fstart: int32(0x2145),
		Fend:   int32(0x2149),
	},
	293: {
		Fstart: int32(0x214e),
		Fend:   int32(0x214e),
	},
	294: {
		Fstart: int32(0x2150),
		Fend:   int32(0x2189),
	},
	295: {
		Fstart: int32(0x2460),
		Fend:   int32(0x249b),
	},
	296: {
		Fstart: int32(0x24ea),
		Fend:   int32(0x24ff),
	},
	297: {
		Fstart: int32(0x2776),
		Fend:   int32(0x2793),
	},
	298: {
		Fstart: int32(0x2c00),
		Fend:   int32(0x2ce4),
	},
	299: {
		Fstart: int32(0x2ceb),
		Fend:   int32(0x2cee),
	},
	300: {
		Fstart: int32(0x2cf2),
		Fend:   int32(0x2cf3),
	},
	301: {
		Fstart: int32(0x2cfd),
		Fend:   int32(0x2cfd),
	},
	302: {
		Fstart: int32(0x2d00),
		Fend:   int32(0x2d25),
	},
	303: {
		Fstart: int32(0x2d27),
		Fend:   int32(0x2d27),
	},
	304: {
		Fstart: int32(0x2d2d),
		Fend:   int32(0x2d2d),
	},
	305: {
		Fstart: int32(0x2d30),
		Fend:   int32(0x2d67),
	},
	306: {
		Fstart: int32(0x2d6f),
		Fend:   int32(0x2d6f),
	},
	307: {
		Fstart: int32(0x2d80),
		Fend:   int32(0x2d96),
	},
	308: {
		Fstart: int32(0x2da0),
		Fend:   int32(0x2da6),
	},
	309: {
		Fstart: int32(0x2da8),
		Fend:   int32(0x2dae),
	},
	310: {
		Fstart: int32(0x2db0),
		Fend:   int32(0x2db6),
	},
	311: {
		Fstart: int32(0x2db8),
		Fend:   int32(0x2dbe),
	},
	312: {
		Fstart: int32(0x2dc0),
		Fend:   int32(0x2dc6),
	},
	313: {
		Fstart: int32(0x2dc8),
		Fend:   int32(0x2dce),
	},
	314: {
		Fstart: int32(0x2dd0),
		Fend:   int32(0x2dd6),
	},
	315: {
		Fstart: int32(0x2dd8),
		Fend:   int32(0x2dde),
	},
	316: {
		Fstart: int32(0x2e2f),
		Fend:   int32(0x2e2f),
	},
	317: {
		Fstart: int32(0x3005),
		Fend:   int32(0x3007),
	},
	318: {
		Fstart: int32(0x3021),
		Fend:   int32(0x3029),
	},
	319: {
		Fstart: int32(0x3031),
		Fend:   int32(0x3035),
	},
	320: {
		Fstart: int32(0x3038),
		Fend:   int32(0x303c),
	},
	321: {
		Fstart: int32(0x3041),
		Fend:   int32(0x3096),
	},
	322: {
		Fstart: int32(0x309d),
		Fend:   int32(0x309f),
	},
	323: {
		Fstart: int32(0x30a1),
		Fend:   int32(0x30fa),
	},
	324: {
		Fstart: int32(0x30fc),
		Fend:   int32(0x30ff),
	},
	325: {
		Fstart: int32(0x3105),
		Fend:   int32(0x312f),
	},
	326: {
		Fstart: int32(0x3131),
		Fend:   int32(0x318e),
	},
	327: {
		Fstart: int32(0x3192),
		Fend:   int32(0x3195),
	},
	328: {
		Fstart: int32(0x31a0),
		Fend:   int32(0x31bf),
	},
	329: {
		Fstart: int32(0x31f0),
		Fend:   int32(0x31ff),
	},
	330: {
		Fstart: int32(0x3220),
		Fend:   int32(0x3229),
	},
	331: {
		Fstart: int32(0x3248),
		Fend:   int32(0x324f),
	},
	332: {
		Fstart: int32(0x3251),
		Fend:   int32(0x325f),
	},
	333: {
		Fstart: int32(0x3280),
		Fend:   int32(0x3289),
	},
	334: {
		Fstart: int32(0x32b1),
		Fend:   int32(0x32bf),
	},
	335: {
		Fstart: int32(0x3400),
		Fend:   int32(0x3400),
	},
	336: {
		Fstart: int32(0x4dbf),
		Fend:   int32(0x4dbf),
	},
	337: {
		Fstart: int32(0x4e00),
		Fend:   int32(0x4e00),
	},
	338: {
		Fstart: int32(0x9fff),
		Fend:   int32(0xa48c),
	},
	339: {
		Fstart: int32(0xa4d0),
		Fend:   int32(0xa4fd),
	},
	340: {
		Fstart: int32(0xa500),
		Fend:   int32(0xa60c),
	},
	341: {
		Fstart: int32(0xa610),
		Fend:   int32(0xa62b),
	},
	342: {
		Fstart: int32(0xa640),
		Fend:   int32(0xa66e),
	},
	343: {
		Fstart: int32(0xa67f),
		Fend:   int32(0xa69d),
	},
	344: {
		Fstart: int32(0xa6a0),
		Fend:   int32(0xa6ef),
	},
	345: {
		Fstart: int32(0xa717),
		Fend:   int32(0xa71f),
	},
	346: {
		Fstart: int32(0xa722),
		Fend:   int32(0xa788),
	},
	347: {
		Fstart: int32(0xa78b),
		Fend:   int32(0xa7ca),
	},
	348: {
		Fstart: int32(0xa7d0),
		Fend:   int32(0xa7d1),
	},
	349: {
		Fstart: int32(0xa7d3),
		Fend:   int32(0xa7d3),
	},
	350: {
		Fstart: int32(0xa7d5),
		Fend:   int32(0xa7d9),
	},
	351: {
		Fstart: int32(0xa7f2),
		Fend:   int32(0xa801),
	},
	352: {
		Fstart: int32(0xa803),
		Fend:   int32(0xa805),
	},
	353: {
		Fstart: int32(0xa807),
		Fend:   int32(0xa80a),
	},
	354: {
		Fstart: int32(0xa80c),
		Fend:   int32(0xa822),
	},
	355: {
		Fstart: int32(0xa830),
		Fend:   int32(0xa835),
	},
	356: {
		Fstart: int32(0xa840),
		Fend:   int32(0xa873),
	},
	357: {
		Fstart: int32(0xa882),
		Fend:   int32(0xa8b3),
	},
	358: {
		Fstart: int32(0xa8d0),
		Fend:   int32(0xa8d9),
	},
	359: {
		Fstart: int32(0xa8f2),
		Fend:   int32(0xa8f7),
	},
	360: {
		Fstart: int32(0xa8fb),
		Fend:   int32(0xa8fb),
	},
	361: {
		Fstart: int32(0xa8fd),
		Fend:   int32(0xa8fe),
	},
	362: {
		Fstart: int32(0xa900),
		Fend:   int32(0xa925),
	},
	363: {
		Fstart: int32(0xa930),
		Fend:   int32(0xa946),
	},
	364: {
		Fstart: int32(0xa960),
		Fend:   int32(0xa97c),
	},
	365: {
		Fstart: int32(0xa984),
		Fend:   int32(0xa9b2),
	},
	366: {
		Fstart: int32(0xa9cf),
		Fend:   int32(0xa9d9),
	},
	367: {
		Fstart: int32(0xa9e0),
		Fend:   int32(0xa9e4),
	},
	368: {
		Fstart: int32(0xa9e6),
		Fend:   int32(0xa9fe),
	},
	369: {
		Fstart: int32(0xaa00),
		Fend:   int32(0xaa28),
	},
	370: {
		Fstart: int32(0xaa40),
		Fend:   int32(0xaa42),
	},
	371: {
		Fstart: int32(0xaa44),
		Fend:   int32(0xaa4b),
	},
	372: {
		Fstart: int32(0xaa50),
		Fend:   int32(0xaa59),
	},
	373: {
		Fstart: int32(0xaa60),
		Fend:   int32(0xaa76),
	},
	374: {
		Fstart: int32(0xaa7a),
		Fend:   int32(0xaa7a),
	},
	375: {
		Fstart: int32(0xaa7e),
		Fend:   int32(0xaaaf),
	},
	376: {
		Fstart: int32(0xaab1),
		Fend:   int32(0xaab1),
	},
	377: {
		Fstart: int32(0xaab5),
		Fend:   int32(0xaab6),
	},
	378: {
		Fstart: int32(0xaab9),
		Fend:   int32(0xaabd),
	},
	379: {
		Fstart: int32(0xaac0),
		Fend:   int32(0xaac0),
	},
	380: {
		Fstart: int32(0xaac2),
		Fend:   int32(0xaac2),
	},
	381: {
		Fstart: int32(0xaadb),
		Fend:   int32(0xaadd),
	},
	382: {
		Fstart: int32(0xaae0),
		Fend:   int32(0xaaea),
	},
	383: {
		Fstart: int32(0xaaf2),
		Fend:   int32(0xaaf4),
	},
	384: {
		Fstart: int32(0xab01),
		Fend:   int32(0xab06),
	},
	385: {
		Fstart: int32(0xab09),
		Fend:   int32(0xab0e),
	},
	386: {
		Fstart: int32(0xab11),
		Fend:   int32(0xab16),
	},
	387: {
		Fstart: int32(0xab20),
		Fend:   int32(0xab26),
	},
	388: {
		Fstart: int32(0xab28),
		Fend:   int32(0xab2e),
	},
	389: {
		Fstart: int32(0xab30),
		Fend:   int32(0xab5a),
	},
	390: {
		Fstart: int32(0xab5c),
		Fend:   int32(0xab69),
	},
	391: {
		Fstart: int32(0xab70),
		Fend:   int32(0xabe2),
	},
	392: {
		Fstart: int32(0xabf0),
		Fend:   int32(0xabf9),
	},
	393: {
		Fstart: int32(0xac00),
		Fend:   int32(0xac00),
	},
	394: {
		Fstart: int32(0xd7a3),
		Fend:   int32(0xd7a3),
	},
	395: {
		Fstart: int32(0xd7b0),
		Fend:   int32(0xd7c6),
	},
	396: {
		Fstart: int32(0xd7cb),
		Fend:   int32(0xd7fb),
	},
	397: {
		Fstart: int32(0xf900),
		Fend:   int32(0xfa6d),
	},
	398: {
		Fstart: int32(0xfa70),
		Fend:   int32(0xfad9),
	},
	399: {
		Fstart: int32(0xfb00),
		Fend:   int32(0xfb06),
	},
	400: {
		Fstart: int32(0xfb13),
		Fend:   int32(0xfb17),
	},
	401: {
		Fstart: int32(0xfb1d),
		Fend:   int32(0xfb1d),
	},
	402: {
		Fstart: int32(0xfb1f),
		Fend:   int32(0xfb28),
	},
	403: {
		Fstart: int32(0xfb2a),
		Fend:   int32(0xfb36),
	},
	404: {
		Fstart: int32(0xfb38),
		Fend:   int32(0xfb3c),
	},
	405: {
		Fstart: int32(0xfb3e),
		Fend:   int32(0xfb3e),
	},
	406: {
		Fstart: int32(0xfb40),
		Fend:   int32(0xfb41),
	},
	407: {
		Fstart: int32(0xfb43),
		Fend:   int32(0xfb44),
	},
	408: {
		Fstart: int32(0xfb46),
		Fend:   int32(0xfbb1),
	},
	409: {
		Fstart: int32(0xfbd3),
		Fend:   int32(0xfd3d),
	},
	410: {
		Fstart: int32(0xfd50),
		Fend:   int32(0xfd8f),
	},
	411: {
		Fstart: int32(0xfd92),
		Fend:   int32(0xfdc7),
	},
	412: {
		Fstart: int32(0xfdf0),
		Fend:   int32(0xfdfb),
	},
	413: {
		Fstart: int32(0xfe70),
		Fend:   int32(0xfe74),
	},
	414: {
		Fstart: int32(0xfe76),
		Fend:   int32(0xfefc),
	},
	415: {
		Fstart: int32(0xff10),
		Fend:   int32(0xff19),
	},
	416: {
		Fstart: int32(0xff21),
		Fend:   int32(0xff3a),
	},
	417: {
		Fstart: int32(0xff41),
		Fend:   int32(0xff5a),
	},
	418: {
		Fstart: int32(0xff66),
		Fend:   int32(0xffbe),
	},
	419: {
		Fstart: int32(0xffc2),
		Fend:   int32(0xffc7),
	},
	420: {
		Fstart: int32(0xffca),
		Fend:   int32(0xffcf),
	},
	421: {
		Fstart: int32(0xffd2),
		Fend:   int32(0xffd7),
	},
	422: {
		Fstart: int32(0xffda),
		Fend:   int32(0xffdc),
	},
	423: {
		Fstart: int32(0x10000),
		Fend:   int32(0x1000b),
	},
	424: {
		Fstart: int32(0x1000d),
		Fend:   int32(0x10026),
	},
	425: {
		Fstart: int32(0x10028),
		Fend:   int32(0x1003a),
	},
	426: {
		Fstart: int32(0x1003c),
		Fend:   int32(0x1003d),
	},
	427: {
		Fstart: int32(0x1003f),
		Fend:   int32(0x1004d),
	},
	428: {
		Fstart: int32(0x10050),
		Fend:   int32(0x1005d),
	},
	429: {
		Fstart: int32(0x10080),
		Fend:   int32(0x100fa),
	},
	430: {
		Fstart: int32(0x10107),
		Fend:   int32(0x10133),
	},
	431: {
		Fstart: int32(0x10140),
		Fend:   int32(0x10178),
	},
	432: {
		Fstart: int32(0x1018a),
		Fend:   int32(0x1018b),
	},
	433: {
		Fstart: int32(0x10280),
		Fend:   int32(0x1029c),
	},
	434: {
		Fstart: int32(0x102a0),
		Fend:   int32(0x102d0),
	},
	435: {
		Fstart: int32(0x102e1),
		Fend:   int32(0x102fb),
	},
	436: {
		Fstart: int32(0x10300),
		Fend:   int32(0x10323),
	},
	437: {
		Fstart: int32(0x1032d),
		Fend:   int32(0x1034a),
	},
	438: {
		Fstart: int32(0x10350),
		Fend:   int32(0x10375),
	},
	439: {
		Fstart: int32(0x10380),
		Fend:   int32(0x1039d),
	},
	440: {
		Fstart: int32(0x103a0),
		Fend:   int32(0x103c3),
	},
	441: {
		Fstart: int32(0x103c8),
		Fend:   int32(0x103cf),
	},
	442: {
		Fstart: int32(0x103d1),
		Fend:   int32(0x103d5),
	},
	443: {
		Fstart: int32(0x10400),
		Fend:   int32(0x1049d),
	},
	444: {
		Fstart: int32(0x104a0),
		Fend:   int32(0x104a9),
	},
	445: {
		Fstart: int32(0x104b0),
		Fend:   int32(0x104d3),
	},
	446: {
		Fstart: int32(0x104d8),
		Fend:   int32(0x104fb),
	},
	447: {
		Fstart: int32(0x10500),
		Fend:   int32(0x10527),
	},
	448: {
		Fstart: int32(0x10530),
		Fend:   int32(0x10563),
	},
	449: {
		Fstart: int32(0x10570),
		Fend:   int32(0x1057a),
	},
	450: {
		Fstart: int32(0x1057c),
		Fend:   int32(0x1058a),
	},
	451: {
		Fstart: int32(0x1058c),
		Fend:   int32(0x10592),
	},
	452: {
		Fstart: int32(0x10594),
		Fend:   int32(0x10595),
	},
	453: {
		Fstart: int32(0x10597),
		Fend:   int32(0x105a1),
	},
	454: {
		Fstart: int32(0x105a3),
		Fend:   int32(0x105b1),
	},
	455: {
		Fstart: int32(0x105b3),
		Fend:   int32(0x105b9),
	},
	456: {
		Fstart: int32(0x105bb),
		Fend:   int32(0x105bc),
	},
	457: {
		Fstart: int32(0x10600),
		Fend:   int32(0x10736),
	},
	458: {
		Fstart: int32(0x10740),
		Fend:   int32(0x10755),
	},
	459: {
		Fstart: int32(0x10760),
		Fend:   int32(0x10767),
	},
	460: {
		Fstart: int32(0x10780),
		Fend:   int32(0x10785),
	},
	461: {
		Fstart: int32(0x10787),
		Fend:   int32(0x107b0),
	},
	462: {
		Fstart: int32(0x107b2),
		Fend:   int32(0x107ba),
	},
	463: {
		Fstart: int32(0x10800),
		Fend:   int32(0x10805),
	},
	464: {
		Fstart: int32(0x10808),
		Fend:   int32(0x10808),
	},
	465: {
		Fstart: int32(0x1080a),
		Fend:   int32(0x10835),
	},
	466: {
		Fstart: int32(0x10837),
		Fend:   int32(0x10838),
	},
	467: {
		Fstart: int32(0x1083c),
		Fend:   int32(0x1083c),
	},
	468: {
		Fstart: int32(0x1083f),
		Fend:   int32(0x10855),
	},
	469: {
		Fstart: int32(0x10858),
		Fend:   int32(0x10876),
	},
	470: {
		Fstart: int32(0x10879),
		Fend:   int32(0x1089e),
	},
	471: {
		Fstart: int32(0x108a7),
		Fend:   int32(0x108af),
	},
	472: {
		Fstart: int32(0x108e0),
		Fend:   int32(0x108f2),
	},
	473: {
		Fstart: int32(0x108f4),
		Fend:   int32(0x108f5),
	},
	474: {
		Fstart: int32(0x108fb),
		Fend:   int32(0x1091b),
	},
}

var aux_sym_PUNCTUATION_token1_character_set_1 = [484]TSCharacterRange{
	0: {
		Fend: int32('\t'),
	},
	1: {
		Fstart: int32(0x0b),
		Fend:   int32('\f'),
	},
	2: {
		Fstart: int32(0x0e),
		Fend:   int32(0x1f),
	},
	3: {
		Fstart: int32('!'),
		Fend:   int32('/'),
	},
	4: {
		Fstart: int32(':'),
		Fend:   int32('@'),
	},
	5: {
		Fstart: int32('['),
		Fend:   int32('`'),
	},
	6: {
		Fstart: int32('{'),
		Fend:   int32(0x9f),
	},
	7: {
		Fstart: int32(0xa1),
		Fend:   int32(0xa9),
	},
	8: {
		Fstart: int32(0xab),
		Fend:   int32(0xb1),
	},
	9: {
		Fstart: int32(0xb4),
		Fend:   int32(0xb4),
	},
	10: {
		Fstart: int32(0xb6),
		Fend:   int32(0xb8),
	},
	11: {
		Fstart: int32(0xbb),
		Fend:   int32(0xbb),
	},
	12: {
		Fstart: int32(0xbf),
		Fend:   int32(0xbf),
	},
	13: {
		Fstart: int32(0xd7),
		Fend:   int32(0xd7),
	},
	14: {
		Fstart: int32(0xf7),
		Fend:   int32(0xf7),
	},
	15: {
		Fstart: int32(0x2c2),
		Fend:   int32(0x2c5),
	},
	16: {
		Fstart: int32(0x2d2),
		Fend:   int32(0x2df),
	},
	17: {
		Fstart: int32(0x2e5),
		Fend:   int32(0x2eb),
	},
	18: {
		Fstart: int32(0x2ed),
		Fend:   int32(0x2ed),
	},
	19: {
		Fstart: int32(0x2ef),
		Fend:   int32(0x36f),
	},
	20: {
		Fstart: int32(0x375),
		Fend:   int32(0x375),
	},
	21: {
		Fstart: int32(0x378),
		Fend:   int32(0x379),
	},
	22: {
		Fstart: int32(0x37e),
		Fend:   int32(0x37e),
	},
	23: {
		Fstart: int32(0x380),
		Fend:   int32(0x385),
	},
	24: {
		Fstart: int32(0x387),
		Fend:   int32(0x387),
	},
	25: {
		Fstart: int32(0x38b),
		Fend:   int32(0x38b),
	},
	26: {
		Fstart: int32(0x38d),
		Fend:   int32(0x38d),
	},
	27: {
		Fstart: int32(0x3a2),
		Fend:   int32(0x3a2),
	},
	28: {
		Fstart: int32(0x3f6),
		Fend:   int32(0x3f6),
	},
	29: {
		Fstart: int32(0x482),
		Fend:   int32(0x489),
	},
	30: {
		Fstart: int32(0x530),
		Fend:   int32(0x530),
	},
	31: {
		Fstart: int32(0x557),
		Fend:   int32(0x558),
	},
	32: {
		Fstart: int32(0x55a),
		Fend:   int32(0x55f),
	},
	33: {
		Fstart: int32(0x589),
		Fend:   int32(0x5cf),
	},
	34: {
		Fstart: int32(0x5eb),
		Fend:   int32(0x5ee),
	},
	35: {
		Fstart: int32(0x5f3),
		Fend:   int32(0x61f),
	},
	36: {
		Fstart: int32(0x64b),
		Fend:   int32(0x65f),
	},
	37: {
		Fstart: int32(0x66a),
		Fend:   int32(0x66d),
	},
	38: {
		Fstart: int32(0x670),
		Fend:   int32(0x670),
	},
	39: {
		Fstart: int32(0x6d4),
		Fend:   int32(0x6d4),
	},
	40: {
		Fstart: int32(0x6d6),
		Fend:   int32(0x6e4),
	},
	41: {
		Fstart: int32(0x6e7),
		Fend:   int32(0x6ed),
	},
	42: {
		Fstart: int32(0x6fd),
		Fend:   int32(0x6fe),
	},
	43: {
		Fstart: int32(0x700),
		Fend:   int32(0x70f),
	},
	44: {
		Fstart: int32(0x711),
		Fend:   int32(0x711),
	},
	45: {
		Fstart: int32(0x730),
		Fend:   int32(0x74c),
	},
	46: {
		Fstart: int32(0x7a6),
		Fend:   int32(0x7b0),
	},
	47: {
		Fstart: int32(0x7b2),
		Fend:   int32(0x7bf),
	},
	48: {
		Fstart: int32(0x7eb),
		Fend:   int32(0x7f3),
	},
	49: {
		Fstart: int32(0x7f6),
		Fend:   int32(0x7f9),
	},
	50: {
		Fstart: int32(0x7fb),
		Fend:   int32(0x7ff),
	},
	51: {
		Fstart: int32(0x816),
		Fend:   int32(0x819),
	},
	52: {
		Fstart: int32(0x81b),
		Fend:   int32(0x823),
	},
	53: {
		Fstart: int32(0x825),
		Fend:   int32(0x827),
	},
	54: {
		Fstart: int32(0x829),
		Fend:   int32(0x83f),
	},
	55: {
		Fstart: int32(0x859),
		Fend:   int32(0x85f),
	},
	56: {
		Fstart: int32(0x86b),
		Fend:   int32(0x86f),
	},
	57: {
		Fstart: int32(0x888),
		Fend:   int32(0x888),
	},
	58: {
		Fstart: int32(0x88f),
		Fend:   int32(0x89f),
	},
	59: {
		Fstart: int32(0x8ca),
		Fend:   int32(0x903),
	},
	60: {
		Fstart: int32(0x93a),
		Fend:   int32(0x93c),
	},
	61: {
		Fstart: int32(0x93e),
		Fend:   int32(0x94f),
	},
	62: {
		Fstart: int32(0x951),
		Fend:   int32(0x957),
	},
	63: {
		Fstart: int32(0x962),
		Fend:   int32(0x965),
	},
	64: {
		Fstart: int32(0x970),
		Fend:   int32(0x970),
	},
	65: {
		Fstart: int32(0x981),
		Fend:   int32(0x984),
	},
	66: {
		Fstart: int32(0x98d),
		Fend:   int32(0x98e),
	},
	67: {
		Fstart: int32(0x991),
		Fend:   int32(0x992),
	},
	68: {
		Fstart: int32(0x9a9),
		Fend:   int32(0x9a9),
	},
	69: {
		Fstart: int32(0x9b1),
		Fend:   int32(0x9b1),
	},
	70: {
		Fstart: int32(0x9b3),
		Fend:   int32(0x9b5),
	},
	71: {
		Fstart: int32(0x9ba),
		Fend:   int32(0x9bc),
	},
	72: {
		Fstart: int32(0x9be),
		Fend:   int32(0x9cd),
	},
	73: {
		Fstart: int32(0x9cf),
		Fend:   int32(0x9db),
	},
	74: {
		Fstart: int32(0x9de),
		Fend:   int32(0x9de),
	},
	75: {
		Fstart: int32(0x9e2),
		Fend:   int32(0x9e5),
	},
	76: {
		Fstart: int32(0x9f2),
		Fend:   int32(0x9f3),
	},
	77: {
		Fstart: int32(0x9fa),
		Fend:   int32(0x9fb),
	},
	78: {
		Fstart: int32(0x9fd),
		Fend:   int32(0xa04),
	},
	79: {
		Fstart: int32(0xa0b),
		Fend:   int32(0xa0e),
	},
	80: {
		Fstart: int32(0xa11),
		Fend:   int32(0xa12),
	},
	81: {
		Fstart: int32(0xa29),
		Fend:   int32(0xa29),
	},
	82: {
		Fstart: int32(0xa31),
		Fend:   int32(0xa31),
	},
	83: {
		Fstart: int32(0xa34),
		Fend:   int32(0xa34),
	},
	84: {
		Fstart: int32(0xa37),
		Fend:   int32(0xa37),
	},
	85: {
		Fstart: int32(0xa3a),
		Fend:   int32(0xa58),
	},
	86: {
		Fstart: int32(0xa5d),
		Fend:   int32(0xa5d),
	},
	87: {
		Fstart: int32(0xa5f),
		Fend:   int32(0xa65),
	},
	88: {
		Fstart: int32(0xa70),
		Fend:   int32(0xa71),
	},
	89: {
		Fstart: int32(0xa75),
		Fend:   int32(0xa84),
	},
	90: {
		Fstart: int32(0xa8e),
		Fend:   int32(0xa8e),
	},
	91: {
		Fstart: int32(0xa92),
		Fend:   int32(0xa92),
	},
	92: {
		Fstart: int32(0xaa9),
		Fend:   int32(0xaa9),
	},
	93: {
		Fstart: int32(0xab1),
		Fend:   int32(0xab1),
	},
	94: {
		Fstart: int32(0xab4),
		Fend:   int32(0xab4),
	},
	95: {
		Fstart: int32(0xaba),
		Fend:   int32(0xabc),
	},
	96: {
		Fstart: int32(0xabe),
		Fend:   int32(0xacf),
	},
	97: {
		Fstart: int32(0xad1),
		Fend:   int32(0xadf),
	},
	98: {
		Fstart: int32(0xae2),
		Fend:   int32(0xae5),
	},
	99: {
		Fstart: int32(0xaf0),
		Fend:   int32(0xaf8),
	},
	100: {
		Fstart: int32(0xafa),
		Fend:   int32(0xb04),
	},
	101: {
		Fstart: int32(0xb0d),
		Fend:   int32(0xb0e),
	},
	102: {
		Fstart: int32(0xb11),
		Fend:   int32(0xb12),
	},
	103: {
		Fstart: int32(0xb29),
		Fend:   int32(0xb29),
	},
	104: {
		Fstart: int32(0xb31),
		Fend:   int32(0xb31),
	},
	105: {
		Fstart: int32(0xb34),
		Fend:   int32(0xb34),
	},
	106: {
		Fstart: int32(0xb3a),
		Fend:   int32(0xb3c),
	},
	107: {
		Fstart: int32(0xb3e),
		Fend:   int32(0xb5b),
	},
	108: {
		Fstart: int32(0xb5e),
		Fend:   int32(0xb5e),
	},
	109: {
		Fstart: int32(0xb62),
		Fend:   int32(0xb65),
	},
	110: {
		Fstart: int32(0xb70),
		Fend:   int32(0xb70),
	},
	111: {
		Fstart: int32(0xb78),
		Fend:   int32(0xb82),
	},
	112: {
		Fstart: int32(0xb84),
		Fend:   int32(0xb84),
	},
	113: {
		Fstart: int32(0xb8b),
		Fend:   int32(0xb8d),
	},
	114: {
		Fstart: int32(0xb91),
		Fend:   int32(0xb91),
	},
	115: {
		Fstart: int32(0xb96),
		Fend:   int32(0xb98),
	},
	116: {
		Fstart: int32(0xb9b),
		Fend:   int32(0xb9b),
	},
	117: {
		Fstart: int32(0xb9d),
		Fend:   int32(0xb9d),
	},
	118: {
		Fstart: int32(0xba0),
		Fend:   int32(0xba2),
	},
	119: {
		Fstart: int32(0xba5),
		Fend:   int32(0xba7),
	},
	120: {
		Fstart: int32(0xbab),
		Fend:   int32(0xbad),
	},
	121: {
		Fstart: int32(0xbba),
		Fend:   int32(0xbcf),
	},
	122: {
		Fstart: int32(0xbd1),
		Fend:   int32(0xbe5),
	},
	123: {
		Fstart: int32(0xbf3),
		Fend:   int32(0xc04),
	},
	124: {
		Fstart: int32(0xc0d),
		Fend:   int32(0xc0d),
	},
	125: {
		Fstart: int32(0xc11),
		Fend:   int32(0xc11),
	},
	126: {
		Fstart: int32(0xc29),
		Fend:   int32(0xc29),
	},
	127: {
		Fstart: int32(0xc3a),
		Fend:   int32(0xc3c),
	},
	128: {
		Fstart: int32(0xc3e),
		Fend:   int32(0xc57),
	},
	129: {
		Fstart: int32(0xc5b),
		Fend:   int32(0xc5c),
	},
	130: {
		Fstart: int32(0xc5e),
		Fend:   int32(0xc5f),
	},
	131: {
		Fstart: int32(0xc62),
		Fend:   int32(0xc65),
	},
	132: {
		Fstart: int32(0xc70),
		Fend:   int32(0xc77),
	},
	133: {
		Fstart: int32(0xc7f),
		Fend:   int32(0xc7f),
	},
	134: {
		Fstart: int32(0xc81),
		Fend:   int32(0xc84),
	},
	135: {
		Fstart: int32(0xc8d),
		Fend:   int32(0xc8d),
	},
	136: {
		Fstart: int32(0xc91),
		Fend:   int32(0xc91),
	},
	137: {
		Fstart: int32(0xca9),
		Fend:   int32(0xca9),
	},
	138: {
		Fstart: int32(0xcb4),
		Fend:   int32(0xcb4),
	},
	139: {
		Fstart: int32(0xcba),
		Fend:   int32(0xcbc),
	},
	140: {
		Fstart: int32(0xcbe),
		Fend:   int32(0xcdc),
	},
	141: {
		Fstart: int32(0xcdf),
		Fend:   int32(0xcdf),
	},
	142: {
		Fstart: int32(0xce2),
		Fend:   int32(0xce5),
	},
	143: {
		Fstart: int32(0xcf0),
		Fend:   int32(0xcf0),
	},
	144: {
		Fstart: int32(0xcf3),
		Fend:   int32(0xd03),
	},
	145: {
		Fstart: int32(0xd0d),
		Fend:   int32(0xd0d),
	},
	146: {
		Fstart: int32(0xd11),
		Fend:   int32(0xd11),
	},
	147: {
		Fstart: int32(0xd3b),
		Fend:   int32(0xd3c),
	},
	148: {
		Fstart: int32(0xd3e),
		Fend:   int32(0xd4d),
	},
	149: {
		Fstart: int32(0xd4f),
		Fend:   int32(0xd53),
	},
	150: {
		Fstart: int32(0xd57),
		Fend:   int32(0xd57),
	},
	151: {
		Fstart: int32(0xd62),
		Fend:   int32(0xd65),
	},
	152: {
		Fstart: int32(0xd79),
		Fend:   int32(0xd79),
	},
	153: {
		Fstart: int32(0xd80),
		Fend:   int32(0xd84),
	},
	154: {
		Fstart: int32(0xd97),
		Fend:   int32(0xd99),
	},
	155: {
		Fstart: int32(0xdb2),
		Fend:   int32(0xdb2),
	},
	156: {
		Fstart: int32(0xdbc),
		Fend:   int32(0xdbc),
	},
	157: {
		Fstart: int32(0xdbe),
		Fend:   int32(0xdbf),
	},
	158: {
		Fstart: int32(0xdc7),
		Fend:   int32(0xde5),
	},
	159: {
		Fstart: int32(0xdf0),
		Fend:   int32(0xe00),
	},
	160: {
		Fstart: int32(0xe31),
		Fend:   int32(0xe31),
	},
	161: {
		Fstart: int32(0xe34),
		Fend:   int32(0xe3f),
	},
	162: {
		Fstart: int32(0xe47),
		Fend:   int32(0xe4f),
	},
	163: {
		Fstart: int32(0xe5a),
		Fend:   int32(0xe80),
	},
	164: {
		Fstart: int32(0xe83),
		Fend:   int32(0xe83),
	},
	165: {
		Fstart: int32(0xe85),
		Fend:   int32(0xe85),
	},
	166: {
		Fstart: int32(0xe8b),
		Fend:   int32(0xe8b),
	},
	167: {
		Fstart: int32(0xea4),
		Fend:   int32(0xea4),
	},
	168: {
		Fstart: int32(0xea6),
		Fend:   int32(0xea6),
	},
	169: {
		Fstart: int32(0xeb1),
		Fend:   int32(0xeb1),
	},
	170: {
		Fstart: int32(0xeb4),
		Fend:   int32(0xebc),
	},
	171: {
		Fstart: int32(0xebe),
		Fend:   int32(0xebf),
	},
	172: {
		Fstart: int32(0xec5),
		Fend:   int32(0xec5),
	},
	173: {
		Fstart: int32(0xec7),
		Fend:   int32(0xecf),
	},
	174: {
		Fstart: int32(0xeda),
		Fend:   int32(0xedb),
	},
	175: {
		Fstart: int32(0xee0),
		Fend:   int32(0xeff),
	},
	176: {
		Fstart: int32(0xf01),
		Fend:   int32(0xf1f),
	},
	177: {
		Fstart: int32(0xf34),
		Fend:   int32(0xf3f),
	},
	178: {
		Fstart: int32(0xf48),
		Fend:   int32(0xf48),
	},
	179: {
		Fstart: int32(0xf6d),
		Fend:   int32(0xf87),
	},
	180: {
		Fstart: int32(0xf8d),
		Fend:   int32(0xfff),
	},
	181: {
		Fstart: int32(0x102b),
		Fend:   int32(0x103e),
	},
	182: {
		Fstart: int32(0x104a),
		Fend:   int32(0x104f),
	},
	183: {
		Fstart: int32(0x1056),
		Fend:   int32(0x1059),
	},
	184: {
		Fstart: int32(0x105e),
		Fend:   int32(0x1060),
	},
	185: {
		Fstart: int32(0x1062),
		Fend:   int32(0x1064),
	},
	186: {
		Fstart: int32(0x1067),
		Fend:   int32(0x106d),
	},
	187: {
		Fstart: int32(0x1071),
		Fend:   int32(0x1074),
	},
	188: {
		Fstart: int32(0x1082),
		Fend:   int32(0x108d),
	},
	189: {
		Fstart: int32(0x108f),
		Fend:   int32(0x108f),
	},
	190: {
		Fstart: int32(0x109a),
		Fend:   int32(0x109f),
	},
	191: {
		Fstart: int32(0x10c6),
		Fend:   int32(0x10c6),
	},
	192: {
		Fstart: int32(0x10c8),
		Fend:   int32(0x10cc),
	},
	193: {
		Fstart: int32(0x10ce),
		Fend:   int32(0x10cf),
	},
	194: {
		Fstart: int32(0x10fb),
		Fend:   int32(0x10fb),
	},
	195: {
		Fstart: int32(0x1249),
		Fend:   int32(0x1249),
	},
	196: {
		Fstart: int32(0x124e),
		Fend:   int32(0x124f),
	},
	197: {
		Fstart: int32(0x1257),
		Fend:   int32(0x1257),
	},
	198: {
		Fstart: int32(0x1259),
		Fend:   int32(0x1259),
	},
	199: {
		Fstart: int32(0x125e),
		Fend:   int32(0x125f),
	},
	200: {
		Fstart: int32(0x1289),
		Fend:   int32(0x1289),
	},
	201: {
		Fstart: int32(0x128e),
		Fend:   int32(0x128f),
	},
	202: {
		Fstart: int32(0x12b1),
		Fend:   int32(0x12b1),
	},
	203: {
		Fstart: int32(0x12b6),
		Fend:   int32(0x12b7),
	},
	204: {
		Fstart: int32(0x12bf),
		Fend:   int32(0x12bf),
	},
	205: {
		Fstart: int32(0x12c1),
		Fend:   int32(0x12c1),
	},
	206: {
		Fstart: int32(0x12c6),
		Fend:   int32(0x12c7),
	},
	207: {
		Fstart: int32(0x12d7),
		Fend:   int32(0x12d7),
	},
	208: {
		Fstart: int32(0x1311),
		Fend:   int32(0x1311),
	},
	209: {
		Fstart: int32(0x1316),
		Fend:   int32(0x1317),
	},
	210: {
		Fstart: int32(0x135b),
		Fend:   int32(0x1368),
	},
	211: {
		Fstart: int32(0x137d),
		Fend:   int32(0x137f),
	},
	212: {
		Fstart: int32(0x1390),
		Fend:   int32(0x139f),
	},
	213: {
		Fstart: int32(0x13f6),
		Fend:   int32(0x13f7),
	},
	214: {
		Fstart: int32(0x13fe),
		Fend:   int32(0x1400),
	},
	215: {
		Fstart: int32(0x166d),
		Fend:   int32(0x166e),
	},
	216: {
		Fstart: int32(0x169b),
		Fend:   int32(0x169f),
	},
	217: {
		Fstart: int32(0x16eb),
		Fend:   int32(0x16ed),
	},
	218: {
		Fstart: int32(0x16f9),
		Fend:   int32(0x16ff),
	},
	219: {
		Fstart: int32(0x1712),
		Fend:   int32(0x171e),
	},
	220: {
		Fstart: int32(0x1732),
		Fend:   int32(0x173f),
	},
	221: {
		Fstart: int32(0x1752),
		Fend:   int32(0x175f),
	},
	222: {
		Fstart: int32(0x176d),
		Fend:   int32(0x176d),
	},
	223: {
		Fstart: int32(0x1771),
		Fend:   int32(0x177f),
	},
	224: {
		Fstart: int32(0x17b4),
		Fend:   int32(0x17d6),
	},
	225: {
		Fstart: int32(0x17d8),
		Fend:   int32(0x17db),
	},
	226: {
		Fstart: int32(0x17dd),
		Fend:   int32(0x17df),
	},
	227: {
		Fstart: int32(0x17ea),
		Fend:   int32(0x17ef),
	},
	228: {
		Fstart: int32(0x17fa),
		Fend:   int32(0x180f),
	},
	229: {
		Fstart: int32(0x181a),
		Fend:   int32(0x181f),
	},
	230: {
		Fstart: int32(0x1879),
		Fend:   int32(0x187f),
	},
	231: {
		Fstart: int32(0x1885),
		Fend:   int32(0x1886),
	},
	232: {
		Fstart: int32(0x18a9),
		Fend:   int32(0x18a9),
	},
	233: {
		Fstart: int32(0x18ab),
		Fend:   int32(0x18af),
	},
	234: {
		Fstart: int32(0x18f6),
		Fend:   int32(0x18ff),
	},
	235: {
		Fstart: int32(0x191f),
		Fend:   int32(0x1945),
	},
	236: {
		Fstart: int32(0x196e),
		Fend:   int32(0x196f),
	},
	237: {
		Fstart: int32(0x1975),
		Fend:   int32(0x197f),
	},
	238: {
		Fstart: int32(0x19ac),
		Fend:   int32(0x19af),
	},
	239: {
		Fstart: int32(0x19ca),
		Fend:   int32(0x19cf),
	},
	240: {
		Fstart: int32(0x19db),
		Fend:   int32(0x19ff),
	},
	241: {
		Fstart: int32(0x1a17),
		Fend:   int32(0x1a1f),
	},
	242: {
		Fstart: int32(0x1a55),
		Fend:   int32(0x1a7f),
	},
	243: {
		Fstart: int32(0x1a8a),
		Fend:   int32(0x1a8f),
	},
	244: {
		Fstart: int32(0x1a9a),
		Fend:   int32(0x1aa6),
	},
	245: {
		Fstart: int32(0x1aa8),
		Fend:   int32(0x1b04),
	},
	246: {
		Fstart: int32(0x1b34),
		Fend:   int32(0x1b44),
	},
	247: {
		Fstart: int32(0x1b4d),
		Fend:   int32(0x1b4f),
	},
	248: {
		Fstart: int32(0x1b5a),
		Fend:   int32(0x1b82),
	},
	249: {
		Fstart: int32(0x1ba1),
		Fend:   int32(0x1bad),
	},
	250: {
		Fstart: int32(0x1be6),
		Fend:   int32(0x1bff),
	},
	251: {
		Fstart: int32(0x1c24),
		Fend:   int32(0x1c3f),
	},
	252: {
		Fstart: int32(0x1c4a),
		Fend:   int32(0x1c4c),
	},
	253: {
		Fstart: int32(0x1c7e),
		Fend:   int32(0x1c7f),
	},
	254: {
		Fstart: int32(0x1c89),
		Fend:   int32(0x1c8f),
	},
	255: {
		Fstart: int32(0x1cbb),
		Fend:   int32(0x1cbc),
	},
	256: {
		Fstart: int32(0x1cc0),
		Fend:   int32(0x1ce8),
	},
	257: {
		Fstart: int32(0x1ced),
		Fend:   int32(0x1ced),
	},
	258: {
		Fstart: int32(0x1cf4),
		Fend:   int32(0x1cf4),
	},
	259: {
		Fstart: int32(0x1cf7),
		Fend:   int32(0x1cf9),
	},
	260: {
		Fstart: int32(0x1cfb),
		Fend:   int32(0x1cff),
	},
	261: {
		Fstart: int32(0x1dc0),
		Fend:   int32(0x1dff),
	},
	262: {
		Fstart: int32(0x1f16),
		Fend:   int32(0x1f17),
	},
	263: {
		Fstart: int32(0x1f1e),
		Fend:   int32(0x1f1f),
	},
	264: {
		Fstart: int32(0x1f46),
		Fend:   int32(0x1f47),
	},
	265: {
		Fstart: int32(0x1f4e),
		Fend:   int32(0x1f4f),
	},
	266: {
		Fstart: int32(0x1f58),
		Fend:   int32(0x1f58),
	},
	267: {
		Fstart: int32(0x1f5a),
		Fend:   int32(0x1f5a),
	},
	268: {
		Fstart: int32(0x1f5c),
		Fend:   int32(0x1f5c),
	},
	269: {
		Fstart: int32(0x1f5e),
		Fend:   int32(0x1f5e),
	},
	270: {
		Fstart: int32(0x1f7e),
		Fend:   int32(0x1f7f),
	},
	271: {
		Fstart: int32(0x1fb5),
		Fend:   int32(0x1fb5),
	},
	272: {
		Fstart: int32(0x1fbd),
		Fend:   int32(0x1fbd),
	},
	273: {
		Fstart: int32(0x1fbf),
		Fend:   int32(0x1fc1),
	},
	274: {
		Fstart: int32(0x1fc5),
		Fend:   int32(0x1fc5),
	},
	275: {
		Fstart: int32(0x1fcd),
		Fend:   int32(0x1fcf),
	},
	276: {
		Fstart: int32(0x1fd4),
		Fend:   int32(0x1fd5),
	},
	277: {
		Fstart: int32(0x1fdc),
		Fend:   int32(0x1fdf),
	},
	278: {
		Fstart: int32(0x1fed),
		Fend:   int32(0x1ff1),
	},
	279: {
		Fstart: int32(0x1ff5),
		Fend:   int32(0x1ff5),
	},
	280: {
		Fstart: int32(0x1ffd),
		Fend:   int32(0x1fff),
	},
	281: {
		Fstart: int32(0x200b),
		Fend:   int32(0x2027),
	},
	282: {
		Fstart: int32(0x202a),
		Fend:   int32(0x202e),
	},
	283: {
		Fstart: int32(0x2030),
		Fend:   int32(0x205e),
	},
	284: {
		Fstart: int32(0x2060),
		Fend:   int32(0x206f),
	},
	285: {
		Fstart: int32(0x2072),
		Fend:   int32(0x2073),
	},
	286: {
		Fstart: int32(0x207a),
		Fend:   int32(0x207e),
	},
	287: {
		Fstart: int32(0x208a),
		Fend:   int32(0x208f),
	},
	288: {
		Fstart: int32(0x209d),
		Fend:   int32(0x2101),
	},
	289: {
		Fstart: int32(0x2103),
		Fend:   int32(0x2106),
	},
	290: {
		Fstart: int32(0x2108),
		Fend:   int32(0x2109),
	},
	291: {
		Fstart: int32(0x2114),
		Fend:   int32(0x2114),
	},
	292: {
		Fstart: int32(0x2116),
		Fend:   int32(0x2118),
	},
	293: {
		Fstart: int32(0x211e),
		Fend:   int32(0x2123),
	},
	294: {
		Fstart: int32(0x2125),
		Fend:   int32(0x2125),
	},
	295: {
		Fstart: int32(0x2127),
		Fend:   int32(0x2127),
	},
	296: {
		Fstart: int32(0x2129),
		Fend:   int32(0x2129),
	},
	297: {
		Fstart: int32(0x212e),
		Fend:   int32(0x212e),
	},
	298: {
		Fstart: int32(0x213a),
		Fend:   int32(0x213b),
	},
	299: {
		Fstart: int32(0x2140),
		Fend:   int32(0x2144),
	},
	300: {
		Fstart: int32(0x214a),
		Fend:   int32(0x214d),
	},
	301: {
		Fstart: int32(0x214f),
		Fend:   int32(0x214f),
	},
	302: {
		Fstart: int32(0x218a),
		Fend:   int32(0x245f),
	},
	303: {
		Fstart: int32(0x249c),
		Fend:   int32(0x24e9),
	},
	304: {
		Fstart: int32(0x2500),
		Fend:   int32(0x2775),
	},
	305: {
		Fstart: int32(0x2794),
		Fend:   int32(0x2bff),
	},
	306: {
		Fstart: int32(0x2ce5),
		Fend:   int32(0x2cea),
	},
	307: {
		Fstart: int32(0x2cef),
		Fend:   int32(0x2cf1),
	},
	308: {
		Fstart: int32(0x2cf4),
		Fend:   int32(0x2cfc),
	},
	309: {
		Fstart: int32(0x2cfe),
		Fend:   int32(0x2cff),
	},
	310: {
		Fstart: int32(0x2d26),
		Fend:   int32(0x2d26),
	},
	311: {
		Fstart: int32(0x2d28),
		Fend:   int32(0x2d2c),
	},
	312: {
		Fstart: int32(0x2d2e),
		Fend:   int32(0x2d2f),
	},
	313: {
		Fstart: int32(0x2d68),
		Fend:   int32(0x2d6e),
	},
	314: {
		Fstart: int32(0x2d70),
		Fend:   int32(0x2d7f),
	},
	315: {
		Fstart: int32(0x2d97),
		Fend:   int32(0x2d9f),
	},
	316: {
		Fstart: int32(0x2da7),
		Fend:   int32(0x2da7),
	},
	317: {
		Fstart: int32(0x2daf),
		Fend:   int32(0x2daf),
	},
	318: {
		Fstart: int32(0x2db7),
		Fend:   int32(0x2db7),
	},
	319: {
		Fstart: int32(0x2dbf),
		Fend:   int32(0x2dbf),
	},
	320: {
		Fstart: int32(0x2dc7),
		Fend:   int32(0x2dc7),
	},
	321: {
		Fstart: int32(0x2dcf),
		Fend:   int32(0x2dcf),
	},
	322: {
		Fstart: int32(0x2dd7),
		Fend:   int32(0x2dd7),
	},
	323: {
		Fstart: int32(0x2ddf),
		Fend:   int32(0x2e2e),
	},
	324: {
		Fstart: int32(0x2e30),
		Fend:   int32(0x2fff),
	},
	325: {
		Fstart: int32(0x3001),
		Fend:   int32(0x3004),
	},
	326: {
		Fstart: int32(0x3008),
		Fend:   int32(0x3020),
	},
	327: {
		Fstart: int32(0x302a),
		Fend:   int32(0x3030),
	},
	328: {
		Fstart: int32(0x3036),
		Fend:   int32(0x3037),
	},
	329: {
		Fstart: int32(0x303d),
		Fend:   int32(0x3040),
	},
	330: {
		Fstart: int32(0x3097),
		Fend:   int32(0x309c),
	},
	331: {
		Fstart: int32(0x30a0),
		Fend:   int32(0x30a0),
	},
	332: {
		Fstart: int32(0x30fb),
		Fend:   int32(0x30fb),
	},
	333: {
		Fstart: int32(0x3100),
		Fend:   int32(0x3104),
	},
	334: {
		Fstart: int32(0x3130),
		Fend:   int32(0x3130),
	},
	335: {
		Fstart: int32(0x318f),
		Fend:   int32(0x3191),
	},
	336: {
		Fstart: int32(0x3196),
		Fend:   int32(0x319f),
	},
	337: {
		Fstart: int32(0x31c0),
		Fend:   int32(0x31ef),
	},
	338: {
		Fstart: int32(0x3200),
		Fend:   int32(0x321f),
	},
	339: {
		Fstart: int32(0x322a),
		Fend:   int32(0x3247),
	},
	340: {
		Fstart: int32(0x3250),
		Fend:   int32(0x3250),
	},
	341: {
		Fstart: int32(0x3260),
		Fend:   int32(0x327f),
	},
	342: {
		Fstart: int32(0x328a),
		Fend:   int32(0x32b0),
	},
	343: {
		Fstart: int32(0x32c0),
		Fend:   int32(0x33ff),
	},
	344: {
		Fstart: int32(0x3401),
		Fend:   int32(0x4dbe),
	},
	345: {
		Fstart: int32(0x4dc0),
		Fend:   int32(0x4dff),
	},
	346: {
		Fstart: int32(0x4e01),
		Fend:   int32(0x9ffe),
	},
	347: {
		Fstart: int32(0xa48d),
		Fend:   int32(0xa4cf),
	},
	348: {
		Fstart: int32(0xa4fe),
		Fend:   int32(0xa4ff),
	},
	349: {
		Fstart: int32(0xa60d),
		Fend:   int32(0xa60f),
	},
	350: {
		Fstart: int32(0xa62c),
		Fend:   int32(0xa63f),
	},
	351: {
		Fstart: int32(0xa66f),
		Fend:   int32(0xa67e),
	},
	352: {
		Fstart: int32(0xa69e),
		Fend:   int32(0xa69f),
	},
	353: {
		Fstart: int32(0xa6f0),
		Fend:   int32(0xa716),
	},
	354: {
		Fstart: int32(0xa720),
		Fend:   int32(0xa721),
	},
	355: {
		Fstart: int32(0xa789),
		Fend:   int32(0xa78a),
	},
	356: {
		Fstart: int32(0xa7cb),
		Fend:   int32(0xa7cf),
	},
	357: {
		Fstart: int32(0xa7d2),
		Fend:   int32(0xa7d2),
	},
	358: {
		Fstart: int32(0xa7d4),
		Fend:   int32(0xa7d4),
	},
	359: {
		Fstart: int32(0xa7da),
		Fend:   int32(0xa7f1),
	},
	360: {
		Fstart: int32(0xa802),
		Fend:   int32(0xa802),
	},
	361: {
		Fstart: int32(0xa806),
		Fend:   int32(0xa806),
	},
	362: {
		Fstart: int32(0xa80b),
		Fend:   int32(0xa80b),
	},
	363: {
		Fstart: int32(0xa823),
		Fend:   int32(0xa82f),
	},
	364: {
		Fstart: int32(0xa836),
		Fend:   int32(0xa83f),
	},
	365: {
		Fstart: int32(0xa874),
		Fend:   int32(0xa881),
	},
	366: {
		Fstart: int32(0xa8b4),
		Fend:   int32(0xa8cf),
	},
	367: {
		Fstart: int32(0xa8da),
		Fend:   int32(0xa8f1),
	},
	368: {
		Fstart: int32(0xa8f8),
		Fend:   int32(0xa8fa),
	},
	369: {
		Fstart: int32(0xa8fc),
		Fend:   int32(0xa8fc),
	},
	370: {
		Fstart: int32(0xa8ff),
		Fend:   int32(0xa8ff),
	},
	371: {
		Fstart: int32(0xa926),
		Fend:   int32(0xa92f),
	},
	372: {
		Fstart: int32(0xa947),
		Fend:   int32(0xa95f),
	},
	373: {
		Fstart: int32(0xa97d),
		Fend:   int32(0xa983),
	},
	374: {
		Fstart: int32(0xa9b3),
		Fend:   int32(0xa9ce),
	},
	375: {
		Fstart: int32(0xa9da),
		Fend:   int32(0xa9df),
	},
	376: {
		Fstart: int32(0xa9e5),
		Fend:   int32(0xa9e5),
	},
	377: {
		Fstart: int32(0xa9ff),
		Fend:   int32(0xa9ff),
	},
	378: {
		Fstart: int32(0xaa29),
		Fend:   int32(0xaa3f),
	},
	379: {
		Fstart: int32(0xaa43),
		Fend:   int32(0xaa43),
	},
	380: {
		Fstart: int32(0xaa4c),
		Fend:   int32(0xaa4f),
	},
	381: {
		Fstart: int32(0xaa5a),
		Fend:   int32(0xaa5f),
	},
	382: {
		Fstart: int32(0xaa77),
		Fend:   int32(0xaa79),
	},
	383: {
		Fstart: int32(0xaa7b),
		Fend:   int32(0xaa7d),
	},
	384: {
		Fstart: int32(0xaab0),
		Fend:   int32(0xaab0),
	},
	385: {
		Fstart: int32(0xaab2),
		Fend:   int32(0xaab4),
	},
	386: {
		Fstart: int32(0xaab7),
		Fend:   int32(0xaab8),
	},
	387: {
		Fstart: int32(0xaabe),
		Fend:   int32(0xaabf),
	},
	388: {
		Fstart: int32(0xaac1),
		Fend:   int32(0xaac1),
	},
	389: {
		Fstart: int32(0xaac3),
		Fend:   int32(0xaada),
	},
	390: {
		Fstart: int32(0xaade),
		Fend:   int32(0xaadf),
	},
	391: {
		Fstart: int32(0xaaeb),
		Fend:   int32(0xaaf1),
	},
	392: {
		Fstart: int32(0xaaf5),
		Fend:   int32(0xab00),
	},
	393: {
		Fstart: int32(0xab07),
		Fend:   int32(0xab08),
	},
	394: {
		Fstart: int32(0xab0f),
		Fend:   int32(0xab10),
	},
	395: {
		Fstart: int32(0xab17),
		Fend:   int32(0xab1f),
	},
	396: {
		Fstart: int32(0xab27),
		Fend:   int32(0xab27),
	},
	397: {
		Fstart: int32(0xab2f),
		Fend:   int32(0xab2f),
	},
	398: {
		Fstart: int32(0xab5b),
		Fend:   int32(0xab5b),
	},
	399: {
		Fstart: int32(0xab6a),
		Fend:   int32(0xab6f),
	},
	400: {
		Fstart: int32(0xabe3),
		Fend:   int32(0xabef),
	},
	401: {
		Fstart: int32(0xabfa),
		Fend:   int32(0xabff),
	},
	402: {
		Fstart: int32(0xac01),
		Fend:   int32(0xd7a2),
	},
	403: {
		Fstart: int32(0xd7a4),
		Fend:   int32(0xd7af),
	},
	404: {
		Fstart: int32(0xd7c7),
		Fend:   int32(0xd7ca),
	},
	405: {
		Fstart: int32(0xd7fc),
		Fend:   int32(0xf8ff),
	},
	406: {
		Fstart: int32(0xfa6e),
		Fend:   int32(0xfa6f),
	},
	407: {
		Fstart: int32(0xfada),
		Fend:   int32(0xfaff),
	},
	408: {
		Fstart: int32(0xfb07),
		Fend:   int32(0xfb12),
	},
	409: {
		Fstart: int32(0xfb18),
		Fend:   int32(0xfb1c),
	},
	410: {
		Fstart: int32(0xfb1e),
		Fend:   int32(0xfb1e),
	},
	411: {
		Fstart: int32(0xfb29),
		Fend:   int32(0xfb29),
	},
	412: {
		Fstart: int32(0xfb37),
		Fend:   int32(0xfb37),
	},
	413: {
		Fstart: int32(0xfb3d),
		Fend:   int32(0xfb3d),
	},
	414: {
		Fstart: int32(0xfb3f),
		Fend:   int32(0xfb3f),
	},
	415: {
		Fstart: int32(0xfb42),
		Fend:   int32(0xfb42),
	},
	416: {
		Fstart: int32(0xfb45),
		Fend:   int32(0xfb45),
	},
	417: {
		Fstart: int32(0xfbb2),
		Fend:   int32(0xfbd2),
	},
	418: {
		Fstart: int32(0xfd3e),
		Fend:   int32(0xfd4f),
	},
	419: {
		Fstart: int32(0xfd90),
		Fend:   int32(0xfd91),
	},
	420: {
		Fstart: int32(0xfdc8),
		Fend:   int32(0xfdef),
	},
	421: {
		Fstart: int32(0xfdfc),
		Fend:   int32(0xfe6f),
	},
	422: {
		Fstart: int32(0xfe75),
		Fend:   int32(0xfe75),
	},
	423: {
		Fstart: int32(0xfefd),
		Fend:   int32(0xff0f),
	},
	424: {
		Fstart: int32(0xff1a),
		Fend:   int32(0xff20),
	},
	425: {
		Fstart: int32(0xff3b),
		Fend:   int32(0xff40),
	},
	426: {
		Fstart: int32(0xff5b),
		Fend:   int32(0xff65),
	},
	427: {
		Fstart: int32(0xffbf),
		Fend:   int32(0xffc1),
	},
	428: {
		Fstart: int32(0xffc8),
		Fend:   int32(0xffc9),
	},
	429: {
		Fstart: int32(0xffd0),
		Fend:   int32(0xffd1),
	},
	430: {
		Fstart: int32(0xffd8),
		Fend:   int32(0xffd9),
	},
	431: {
		Fstart: int32(0xffdd),
		Fend:   int32(0xffff),
	},
	432: {
		Fstart: int32(0x1000c),
		Fend:   int32(0x1000c),
	},
	433: {
		Fstart: int32(0x10027),
		Fend:   int32(0x10027),
	},
	434: {
		Fstart: int32(0x1003b),
		Fend:   int32(0x1003b),
	},
	435: {
		Fstart: int32(0x1003e),
		Fend:   int32(0x1003e),
	},
	436: {
		Fstart: int32(0x1004e),
		Fend:   int32(0x1004f),
	},
	437: {
		Fstart: int32(0x1005e),
		Fend:   int32(0x1007f),
	},
	438: {
		Fstart: int32(0x100fb),
		Fend:   int32(0x10106),
	},
	439: {
		Fstart: int32(0x10134),
		Fend:   int32(0x1013f),
	},
	440: {
		Fstart: int32(0x10179),
		Fend:   int32(0x10189),
	},
	441: {
		Fstart: int32(0x1018c),
		Fend:   int32(0x1027f),
	},
	442: {
		Fstart: int32(0x1029d),
		Fend:   int32(0x1029f),
	},
	443: {
		Fstart: int32(0x102d1),
		Fend:   int32(0x102e0),
	},
	444: {
		Fstart: int32(0x102fc),
		Fend:   int32(0x102ff),
	},
	445: {
		Fstart: int32(0x10324),
		Fend:   int32(0x1032c),
	},
	446: {
		Fstart: int32(0x1034b),
		Fend:   int32(0x1034f),
	},
	447: {
		Fstart: int32(0x10376),
		Fend:   int32(0x1037f),
	},
	448: {
		Fstart: int32(0x1039e),
		Fend:   int32(0x1039f),
	},
	449: {
		Fstart: int32(0x103c4),
		Fend:   int32(0x103c7),
	},
	450: {
		Fstart: int32(0x103d0),
		Fend:   int32(0x103d0),
	},
	451: {
		Fstart: int32(0x103d6),
		Fend:   int32(0x103ff),
	},
	452: {
		Fstart: int32(0x1049e),
		Fend:   int32(0x1049f),
	},
	453: {
		Fstart: int32(0x104aa),
		Fend:   int32(0x104af),
	},
	454: {
		Fstart: int32(0x104d4),
		Fend:   int32(0x104d7),
	},
	455: {
		Fstart: int32(0x104fc),
		Fend:   int32(0x104ff),
	},
	456: {
		Fstart: int32(0x10528),
		Fend:   int32(0x1052f),
	},
	457: {
		Fstart: int32(0x10564),
		Fend:   int32(0x1056f),
	},
	458: {
		Fstart: int32(0x1057b),
		Fend:   int32(0x1057b),
	},
	459: {
		Fstart: int32(0x1058b),
		Fend:   int32(0x1058b),
	},
	460: {
		Fstart: int32(0x10593),
		Fend:   int32(0x10593),
	},
	461: {
		Fstart: int32(0x10596),
		Fend:   int32(0x10596),
	},
	462: {
		Fstart: int32(0x105a2),
		Fend:   int32(0x105a2),
	},
	463: {
		Fstart: int32(0x105b2),
		Fend:   int32(0x105b2),
	},
	464: {
		Fstart: int32(0x105ba),
		Fend:   int32(0x105ba),
	},
	465: {
		Fstart: int32(0x105bd),
		Fend:   int32(0x105ff),
	},
	466: {
		Fstart: int32(0x10737),
		Fend:   int32(0x1073f),
	},
	467: {
		Fstart: int32(0x10756),
		Fend:   int32(0x1075f),
	},
	468: {
		Fstart: int32(0x10768),
		Fend:   int32(0x1077f),
	},
	469: {
		Fstart: int32(0x10786),
		Fend:   int32(0x10786),
	},
	470: {
		Fstart: int32(0x107b1),
		Fend:   int32(0x107b1),
	},
	471: {
		Fstart: int32(0x107bb),
		Fend:   int32(0x107ff),
	},
	472: {
		Fstart: int32(0x10806),
		Fend:   int32(0x10807),
	},
	473: {
		Fstart: int32(0x10809),
		Fend:   int32(0x10809),
	},
	474: {
		Fstart: int32(0x10836),
		Fend:   int32(0x10836),
	},
	475: {
		Fstart: int32(0x10839),
		Fend:   int32(0x1083b),
	},
	476: {
		Fstart: int32(0x1083d),
		Fend:   int32(0x1083e),
	},
	477: {
		Fstart: int32(0x10856),
		Fend:   int32(0x10857),
	},
	478: {
		Fstart: int32(0x10877),
		Fend:   int32(0x10878),
	},
	479: {
		Fstart: int32(0x1089f),
		Fend:   int32(0x108a6),
	},
	480: {
		Fstart: int32(0x108b0),
		Fend:   int32(0x108df),
	},
	481: {
		Fstart: int32(0x108f3),
		Fend:   int32(0x108f3),
	},
	482: {
		Fstart: int32(0x108f6),
		Fend:   int32(0x108fa),
	},
	483: {
		Fstart: int32(0x1091c),
		Fend:   int32(0x10ffff),
	},
}

func ts_lex(tls *libc.TLS, lexer uintptr, state TSStateId) (r uint8) {
	var eof, result, skip, v4 uint8
	var half_size, i, i1, i2, i3, i4, i5, i6, i7, i8, index, mid_index, size uint32_t
	var lookahead1, v3 int32_t
	var range_token, range_token1, v2 uintptr
	var v26 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = eof, half_size, i, i1, i2, i3, i4, i5, i6, i7, i8, index, lookahead1, mid_index, range_token, range_token1, result, size, skip, v2, v26, v3, v4
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
			state = uint16(806)
			goto next_state
		}
		i = uint32(0)
		for {
			if !(uint64(i) < libc.Uint64FromInt64(96)/libc.Uint64FromInt64(2)) {
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
		if lookahead1 == int32(' ') || lookahead1 == int32(0xa0) || lookahead1 == int32(0x1680) || int32(0x2000) <= lookahead1 && lookahead1 <= int32(0x200a) || lookahead1 == int32(0x202f) || lookahead1 == int32(0x205f) || lookahead1 == int32(0x3000) {
			state = uint16(910)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(862)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&aux_sym_WORD_CHAR_token1_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(475) - index
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
			state = uint16(807)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32(0x2028) && lookahead1 != int32(0x2029) {
			state = uint16(863)
			goto next_state
		}
		return result
	case int32(1):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(863)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(897)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(899)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(891)
			goto next_state
		}
		if lookahead1 == int32('{') {
			state = uint16(888)
			goto next_state
		}
		if lookahead1 == int32(' ') || lookahead1 == int32(0xa0) || lookahead1 == int32(0x1680) || int32(0x2000) <= lookahead1 && lookahead1 <= int32(0x200a) || lookahead1 == int32(0x202f) || lookahead1 == int32(0x205f) || lookahead1 == int32(0x3000) {
			state = uint16(896)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&aux_sym_WORD_CHAR_token1_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(475) - index
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
			state = uint16(807)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32(0x2028) && lookahead1 != int32(0x2029) {
			state = uint16(863)
			goto next_state
		}
		return result
	case int32(2):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(863)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(897)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(899)
			goto next_state
		}
		if lookahead1 == int32('{') {
			state = uint16(888)
			goto next_state
		}
		if lookahead1 == int32(' ') || lookahead1 == int32(0xa0) || lookahead1 == int32(0x1680) || int32(0x2000) <= lookahead1 && lookahead1 <= int32(0x200a) || lookahead1 == int32(0x202f) || lookahead1 == int32(0x205f) || lookahead1 == int32(0x3000) {
			state = uint16(910)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&aux_sym_WORD_CHAR_token1_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(475) - index
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
			state = uint16(807)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32(0x2028) && lookahead1 != int32(0x2029) {
			state = uint16(863)
			goto next_state
		}
		return result
	case int32(3):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(863)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(897)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(899)
			goto next_state
		}
		if lookahead1 == int32('{') {
			state = uint16(888)
			goto next_state
		}
		if lookahead1 == int32(' ') || lookahead1 == int32(0xa0) || lookahead1 == int32(0x1680) || int32(0x2000) <= lookahead1 && lookahead1 <= int32(0x200a) || lookahead1 == int32(0x202f) || lookahead1 == int32(0x205f) || lookahead1 == int32(0x3000) {
			state = uint16(896)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&aux_sym_WORD_CHAR_token1_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(475) - index
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
			state = uint16(807)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32(0x2028) && lookahead1 != int32(0x2029) {
			state = uint16(863)
			goto next_state
		}
		return result
	case int32(4):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('#') {
			state = uint16(911)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(5):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(50)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(6):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('A') {
			state = uint16(30)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(7):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('A') {
			state = uint16(14)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(8):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('A') {
			state = uint16(13)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(9):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('B') {
			state = uint16(35)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(10):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('C') {
			state = uint16(36)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(11):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('C') {
			state = uint16(22)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(12):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('C') {
			state = uint16(19)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(13):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('C') {
			state = uint16(16)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(14):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('D') {
			state = uint16(914)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(15):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('E') {
			state = uint16(36)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(16):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('E') {
			state = uint16(914)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(17):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('E') {
			state = uint16(10)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(18):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('E') {
			state = uint16(40)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(19):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('H') {
			state = uint16(914)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(20):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('H') {
			state = uint16(32)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(21):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('I') {
			state = uint16(28)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(22):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('K') {
			state = uint16(15)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(23):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('L') {
			state = uint16(18)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(24):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('L') {
			state = uint16(914)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(25):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('N') {
			state = uint16(26)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(26):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('N') {
			state = uint16(17)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(27):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('N') {
			state = uint16(34)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(28):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('O') {
			state = uint16(27)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(29):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('O') {
			state = uint16(11)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(30):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('P') {
			state = uint16(20)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(31):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('P') {
			state = uint16(5)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(32):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('Q') {
			state = uint16(24)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(33):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('S') {
			state = uint16(36)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(34):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('S') {
			state = uint16(914)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(35):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('S') {
			state = uint16(29)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(36):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('T') {
			state = uint16(914)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(37):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('T') {
			state = uint16(21)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(38):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('T') {
			state = uint16(12)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(39):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('T') {
			state = uint16(31)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(40):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('T') {
			state = uint16(16)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(41):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('a') {
			state = uint16(48)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(42):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('e') {
			state = uint16(46)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(43):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(45)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(44):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(51)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(45):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('o') {
			state = uint16(44)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(46):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('r') {
			state = uint16(49)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(47):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('t') {
			state = uint16(41)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(48):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('t') {
			state = uint16(43)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(49):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('y') {
			state = uint16(51)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(50):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(918)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(51):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32(' ') || lookahead1 == int32(0xa0) || lookahead1 == int32(0x1680) || int32(0x2000) <= lookahead1 && lookahead1 <= int32(0x200a) || lookahead1 == int32(0x202f) || lookahead1 == int32(0x205f) || lookahead1 == int32(0x3000) {
			state = uint16(53)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(52):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(53):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(965)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(965)
			goto next_state
		}
		if lookahead1 == int32('{') {
			state = uint16(55)
			goto next_state
		}
		if lookahead1 == int32(' ') || lookahead1 == int32(0xa0) || lookahead1 == int32(0x1680) || int32(0x2000) <= lookahead1 && lookahead1 <= int32(0x200a) || lookahead1 == int32(0x202f) || lookahead1 == int32(0x205f) || lookahead1 == int32(0x3000) {
			state = uint16(53)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(54)
			goto next_state
		}
		return result
	case int32(54):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(965)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(965)
			goto next_state
		}
		if lookahead1 == int32('{') {
			state = uint16(55)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(54)
			goto next_state
		}
		return result
	case int32(55):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(948)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(946)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(947)
			goto next_state
		}
		if lookahead1 == int32('{') {
			state = uint16(55)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(54)
			goto next_state
		}
		return result
	case int32(56):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('#') {
			state = uint16(911)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(57):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(102)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(58):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('A') {
			state = uint16(66)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(59):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('A') {
			state = uint16(83)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(60):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('A') {
			state = uint16(65)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(61):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('B') {
			state = uint16(86)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(62):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('C') {
			state = uint16(74)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(63):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('C') {
			state = uint16(71)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(64):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('C') {
			state = uint16(88)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(65):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('C') {
			state = uint16(67)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(66):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('D') {
			state = uint16(915)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(67):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('E') {
			state = uint16(915)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(68):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('E') {
			state = uint16(88)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(69):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('E') {
			state = uint16(64)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(70):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('E') {
			state = uint16(92)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(71):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('H') {
			state = uint16(915)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(72):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('H') {
			state = uint16(84)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(73):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('I') {
			state = uint16(80)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(74):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('K') {
			state = uint16(68)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(75):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('L') {
			state = uint16(915)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(76):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('L') {
			state = uint16(70)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(77):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('N') {
			state = uint16(79)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(78):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('N') {
			state = uint16(85)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(79):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('N') {
			state = uint16(69)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(80):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('O') {
			state = uint16(78)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(81):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('O') {
			state = uint16(62)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(82):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('P') {
			state = uint16(57)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(83):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('P') {
			state = uint16(72)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(84):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('Q') {
			state = uint16(75)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(85):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('S') {
			state = uint16(915)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(86):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('S') {
			state = uint16(81)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(87):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('S') {
			state = uint16(88)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(88):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('T') {
			state = uint16(915)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(89):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('T') {
			state = uint16(73)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(90):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('T') {
			state = uint16(82)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(91):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('T') {
			state = uint16(63)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(92):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('T') {
			state = uint16(67)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(93):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('a') {
			state = uint16(100)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(94):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('e') {
			state = uint16(98)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(95):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(97)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(96):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(103)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(97):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('o') {
			state = uint16(96)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(98):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('r') {
			state = uint16(101)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(99):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('t') {
			state = uint16(93)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(100):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('t') {
			state = uint16(95)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(101):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('y') {
			state = uint16(103)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(102):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(919)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(103):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32(' ') || lookahead1 == int32(0xa0) || lookahead1 == int32(0x1680) || int32(0x2000) <= lookahead1 && lookahead1 <= int32(0x200a) || lookahead1 == int32(0x202f) || lookahead1 == int32(0x205f) || lookahead1 == int32(0x3000) {
			state = uint16(105)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(104):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(105):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(960)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(960)
			goto next_state
		}
		if lookahead1 == int32('{') {
			state = uint16(165)
			goto next_state
		}
		if lookahead1 == int32(' ') || lookahead1 == int32(0xa0) || lookahead1 == int32(0x1680) || int32(0x2000) <= lookahead1 && lookahead1 <= int32(0x200a) || lookahead1 == int32(0x202f) || lookahead1 == int32(0x205f) || lookahead1 == int32(0x3000) {
			state = uint16(105)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(106)
			goto next_state
		}
		return result
	case int32(106):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(960)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(960)
			goto next_state
		}
		if lookahead1 == int32('{') {
			state = uint16(165)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(106)
			goto next_state
		}
		return result
	case int32(107):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('#') {
			state = uint16(911)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(108):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('#') {
			state = uint16(905)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(109)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(109):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(905)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(110):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(155)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(111):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('A') {
			state = uint16(119)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(112):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('A') {
			state = uint16(136)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(113):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('A') {
			state = uint16(117)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(114):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('B') {
			state = uint16(139)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(115):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('C') {
			state = uint16(127)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(116):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('C') {
			state = uint16(124)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(117):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('C') {
			state = uint16(120)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(118):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('C') {
			state = uint16(141)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(119):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('D') {
			state = uint16(916)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(120):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('E') {
			state = uint16(916)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(121):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('E') {
			state = uint16(141)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(122):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('E') {
			state = uint16(118)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(123):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('E') {
			state = uint16(145)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(124):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('H') {
			state = uint16(916)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(125):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('H') {
			state = uint16(137)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(126):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('I') {
			state = uint16(133)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(127):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('K') {
			state = uint16(121)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(128):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('L') {
			state = uint16(916)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(129):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('L') {
			state = uint16(123)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(130):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('N') {
			state = uint16(132)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(131):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('N') {
			state = uint16(138)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(132):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('N') {
			state = uint16(122)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(133):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('O') {
			state = uint16(131)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(134):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('O') {
			state = uint16(115)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(135):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('P') {
			state = uint16(110)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(136):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('P') {
			state = uint16(125)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(137):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('Q') {
			state = uint16(128)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(138):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('S') {
			state = uint16(916)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(139):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('S') {
			state = uint16(134)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(140):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('S') {
			state = uint16(141)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(141):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('T') {
			state = uint16(916)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(142):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('T') {
			state = uint16(126)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(143):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('T') {
			state = uint16(135)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(144):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('T') {
			state = uint16(116)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(145):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('T') {
			state = uint16(120)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(146):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('a') {
			state = uint16(153)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(147):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('e') {
			state = uint16(151)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(148):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(150)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(149):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(156)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(150):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('o') {
			state = uint16(149)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(151):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('r') {
			state = uint16(154)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(152):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('t') {
			state = uint16(146)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(153):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('t') {
			state = uint16(148)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(154):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('y') {
			state = uint16(156)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(155):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(920)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(156):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32(' ') || lookahead1 == int32(0xa0) || lookahead1 == int32(0x1680) || int32(0x2000) <= lookahead1 && lookahead1 <= int32(0x200a) || lookahead1 == int32(0x202f) || lookahead1 == int32(0x205f) || lookahead1 == int32(0x3000) {
			state = uint16(158)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(157):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(158):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(969)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(969)
			goto next_state
		}
		if lookahead1 == int32('{') {
			state = uint16(159)
			goto next_state
		}
		if lookahead1 == int32(' ') || lookahead1 == int32(0xa0) || lookahead1 == int32(0x1680) || int32(0x2000) <= lookahead1 && lookahead1 <= int32(0x200a) || lookahead1 == int32(0x202f) || lookahead1 == int32(0x205f) || lookahead1 == int32(0x3000) {
			state = uint16(158)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(159)
			goto next_state
		}
		return result
	case int32(159):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(969)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(969)
			goto next_state
		}
		if lookahead1 == int32('{') {
			state = uint16(159)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(159)
			goto next_state
		}
		return result
	case int32(160):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(883)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(931)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(931)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(891)
			goto next_state
		}
		if lookahead1 == int32('{') {
			state = uint16(888)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&aux_sym_WORD_CHAR_token1_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(475) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _21
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _21
	_21:
		if v4 != 0 {
			state = uint16(807)
			goto next_state
		}
		if v26 = !(eof != 0); v26 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_PUNCTUATION_token1_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(484) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _25
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _25
		_25:
		}
		if v26 && v4 != 0 {
			state = uint16(863)
			goto next_state
		}
		return result
	case int32(161):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(901)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(900)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(901)
			goto next_state
		}
		if lookahead1 == int32('%') {
			state = uint16(163)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(164)
			goto next_state
		}
		return result
	case int32(162):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(901)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(900)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(901)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(908)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(164)
			goto next_state
		}
		return result
	case int32(163):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(901)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(900)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(901)
			goto next_state
		}
		if lookahead1 == int32('}') {
			state = uint16(934)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(164)
			goto next_state
		}
		return result
	case int32(164):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(901)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(900)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(901)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(164)
			goto next_state
		}
		return result
	case int32(165):
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(950)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(946)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(949)
			goto next_state
		}
		if lookahead1 == int32('{') {
			state = uint16(165)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(106)
			goto next_state
		}
		return result
	case int32(166):
		if lookahead1 == int32('\r') {
			state = uint16(1011)
			goto next_state
		}
		if !(eof != 0) && lookahead1 == 00 || lookahead1 == int32('\n') {
			state = uint16(1010)
			goto next_state
		}
		if lookahead1 == int32('$') || lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') || int32(0xa1) <= lookahead1 && lookahead1 <= int32(0xffff) {
			state = uint16(1008)
			goto next_state
		}
		return result
	case int32(167):
		i1 = uint32(0)
		for {
			if !(uint64(i1) < libc.Uint64FromInt64(96)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token1[i1]) == lookahead1 {
				state = map_token1[i1+uint32(1)]
				goto next_state
			}
			goto _27
		_27:
			;
			i1 = i1 + uint32(2)
		}
		if int32('1') <= lookahead1 && lookahead1 <= int32('5') {
			state = uint16(797)
			goto next_state
		}
		if lookahead1 == int32(' ') || lookahead1 == int32(0xa0) || lookahead1 == int32(0x1680) || int32(0x2000) <= lookahead1 && lookahead1 <= int32(0x200a) || lookahead1 == int32(0x202f) || lookahead1 == int32(0x205f) || lookahead1 == int32(0x3000) {
			state = uint16(896)
			goto next_state
		}
		return result
	case int32(168):
		if lookahead1 == int32('\r') {
			state = uint16(931)
			goto next_state
		}
		if !(eof != 0) && lookahead1 == 00 || lookahead1 == int32('\n') {
			state = uint16(931)
			goto next_state
		}
		if lookahead1 == int32(' ') || lookahead1 == int32(0xa0) {
			state = uint16(896)
			goto next_state
		}
		if lookahead1 == int32(0x1680) || int32(0x2000) <= lookahead1 && lookahead1 <= int32(0x200a) || lookahead1 == int32(0x202f) || lookahead1 == int32(0x205f) || lookahead1 == int32(0x3000) {
			state = uint16(895)
			goto next_state
		}
		if lookahead1 == int32('$') || lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') || int32(0xa1) <= lookahead1 && lookahead1 <= int32(0xffff) {
			state = uint16(1008)
			goto next_state
		}
		return result
	case int32(169):
		if lookahead1 == int32('\r') {
			state = uint16(970)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(908)
			goto next_state
		}
		if !(eof != 0) && lookahead1 == 00 || lookahead1 == int32('\n') {
			state = uint16(970)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(170)
			goto next_state
		}
		return result
	case int32(170):
		if lookahead1 == int32('\r') {
			state = uint16(970)
			goto next_state
		}
		if !(eof != 0) && lookahead1 == 00 || lookahead1 == int32('\n') {
			state = uint16(970)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32('@') {
			state = uint16(170)
			goto next_state
		}
		return result
	case int32(171):
		if lookahead1 == int32(' ') {
			state = uint16(314)
			goto next_state
		}
		return result
	case int32(172):
		if lookahead1 == int32(' ') {
			state = uint16(256)
			goto next_state
		}
		return result
	case int32(173):
		if lookahead1 == int32(' ') {
			state = uint16(231)
			goto next_state
		}
		return result
	case int32(174):
		if lookahead1 == int32(' ') {
			state = uint16(251)
			goto next_state
		}
		return result
	case int32(175):
		if lookahead1 == int32(' ') {
			state = uint16(316)
			goto next_state
		}
		return result
	case int32(176):
		if lookahead1 == int32(' ') {
			state = uint16(291)
			goto next_state
		}
		return result
	case int32(177):
		if lookahead1 == int32(' ') {
			state = uint16(259)
			goto next_state
		}
		return result
	case int32(178):
		if lookahead1 == int32(' ') {
			state = uint16(253)
			goto next_state
		}
		return result
	case int32(179):
		if lookahead1 == int32(' ') {
			state = uint16(270)
			goto next_state
		}
		return result
	case int32(180):
		if lookahead1 == int32(' ') {
			state = uint16(254)
			goto next_state
		}
		return result
	case int32(181):
		if lookahead1 == int32(' ') {
			state = uint16(243)
			goto next_state
		}
		return result
	case int32(182):
		if lookahead1 == int32(' ') {
			state = uint16(315)
			goto next_state
		}
		return result
	case int32(183):
		if lookahead1 == int32(' ') {
			state = uint16(292)
			goto next_state
		}
		return result
	case int32(184):
		if lookahead1 == int32(' ') {
			state = uint16(255)
			goto next_state
		}
		return result
	case int32(185):
		if lookahead1 == int32(' ') {
			state = uint16(269)
			goto next_state
		}
		return result
	case int32(186):
		if lookahead1 == int32(' ') {
			state = uint16(262)
			goto next_state
		}
		return result
	case int32(187):
		if lookahead1 == int32(' ') {
			state = uint16(280)
			goto next_state
		}
		return result
	case int32(188):
		if lookahead1 == int32(' ') {
			state = uint16(252)
			goto next_state
		}
		return result
	case int32(189):
		if lookahead1 == int32(' ') {
			state = uint16(274)
			goto next_state
		}
		return result
	case int32(190):
		if lookahead1 == int32(' ') {
			state = uint16(268)
			goto next_state
		}
		return result
	case int32(191):
		if lookahead1 == int32(' ') {
			state = uint16(285)
			goto next_state
		}
		return result
	case int32(192):
		if lookahead1 == int32(' ') {
			state = uint16(285)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(575)
			goto next_state
		}
		return result
	case int32(193):
		if lookahead1 == int32(' ') {
			state = uint16(336)
			goto next_state
		}
		return result
	case int32(194):
		if lookahead1 == int32(' ') {
			state = uint16(275)
			goto next_state
		}
		return result
	case int32(195):
		if lookahead1 == int32(' ') {
			state = uint16(267)
			goto next_state
		}
		return result
	case int32(196):
		if lookahead1 == int32(' ') {
			state = uint16(233)
			goto next_state
		}
		return result
	case int32(197):
		if lookahead1 == int32(' ') {
			state = uint16(299)
			goto next_state
		}
		return result
	case int32(198):
		if lookahead1 == int32(' ') {
			state = uint16(250)
			goto next_state
		}
		return result
	case int32(199):
		if lookahead1 == int32(' ') {
			state = uint16(284)
			goto next_state
		}
		return result
	case int32(200):
		if lookahead1 == int32(' ') {
			state = uint16(240)
			goto next_state
		}
		return result
	case int32(201):
		if lookahead1 == int32(' ') {
			state = uint16(240)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(224)
			goto next_state
		}
		if lookahead1 == int32('t') {
			state = uint16(173)
			goto next_state
		}
		return result
	case int32(202):
		if lookahead1 == int32(' ') {
			state = uint16(301)
			goto next_state
		}
		return result
	case int32(203):
		if lookahead1 == int32(' ') {
			state = uint16(249)
			goto next_state
		}
		return result
	case int32(204):
		if lookahead1 == int32(' ') {
			state = uint16(241)
			goto next_state
		}
		return result
	case int32(205):
		if lookahead1 == int32(' ') {
			state = uint16(310)
			goto next_state
		}
		return result
	case int32(206):
		if lookahead1 == int32(' ') {
			state = uint16(302)
			goto next_state
		}
		return result
	case int32(207):
		if lookahead1 == int32(' ') {
			state = uint16(311)
			goto next_state
		}
		return result
	case int32(208):
		if lookahead1 == int32(' ') {
			state = uint16(309)
			goto next_state
		}
		return result
	case int32(209):
		if lookahead1 == int32(' ') {
			state = uint16(744)
			goto next_state
		}
		return result
	case int32(210):
		if lookahead1 == int32(' ') {
			state = uint16(293)
			goto next_state
		}
		return result
	case int32(211):
		if lookahead1 == int32(' ') {
			state = uint16(244)
			goto next_state
		}
		return result
	case int32(212):
		if lookahead1 == int32(' ') {
			state = uint16(232)
			goto next_state
		}
		return result
	case int32(213):
		if lookahead1 == int32(' ') {
			state = uint16(298)
			goto next_state
		}
		return result
	case int32(214):
		if lookahead1 == int32(' ') {
			state = uint16(234)
			goto next_state
		}
		return result
	case int32(215):
		if lookahead1 == int32(' ') {
			state = uint16(290)
			goto next_state
		}
		return result
	case int32(216):
		if lookahead1 == int32(' ') {
			state = uint16(294)
			goto next_state
		}
		return result
	case int32(217):
		if lookahead1 == int32(' ') {
			state = uint16(286)
			goto next_state
		}
		return result
	case int32(218):
		if lookahead1 == int32(' ') {
			state = uint16(276)
			goto next_state
		}
		return result
	case int32(219):
		if lookahead1 == int32(' ') {
			state = uint16(313)
			goto next_state
		}
		return result
	case int32(220):
		if lookahead1 == int32(' ') {
			state = uint16(289)
			goto next_state
		}
		return result
	case int32(221):
		if lookahead1 == int32(' ') {
			state = uint16(277)
			goto next_state
		}
		return result
	case int32(222):
		if lookahead1 == int32('#') {
			state = uint16(911)
			goto next_state
		}
		return result
	case int32(223):
		if lookahead1 == int32('\'') {
			state = uint16(562)
			goto next_state
		}
		if lookahead1 == int32('M') {
			state = uint16(171)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(697)
			goto next_state
		}
		return result
	case int32(224):
		if lookahead1 == int32('-') {
			state = uint16(230)
			goto next_state
		}
		return result
	case int32(225):
		if lookahead1 == int32('-') {
			state = uint16(300)
			goto next_state
		}
		if lookahead1 == int32('p') {
			state = uint16(557)
			goto next_state
		}
		return result
	case int32(226):
		if lookahead1 == int32('/') {
			state = uint16(798)
			goto next_state
		}
		return result
	case int32(227):
		if lookahead1 == int32('A') {
			state = uint16(281)
			goto next_state
		}
		return result
	case int32(228):
		if lookahead1 == int32('A') {
			state = uint16(242)
			goto next_state
		}
		return result
	case int32(229):
		if lookahead1 == int32('A') {
			state = uint16(239)
			goto next_state
		}
		return result
	case int32(230):
		if lookahead1 == int32('A') {
			state = uint16(771)
			goto next_state
		}
		return result
	case int32(231):
		if lookahead1 == int32('A') {
			state = uint16(383)
			goto next_state
		}
		if lookahead1 == int32('E') {
			state = uint16(786)
			goto next_state
		}
		if lookahead1 == int32('F') {
			state = uint16(606)
			goto next_state
		}
		if lookahead1 == int32('I') {
			state = uint16(564)
			goto next_state
		}
		if lookahead1 == int32('M') {
			state = uint16(630)
			goto next_state
		}
		return result
	case int32(232):
		if lookahead1 == int32('A') {
			state = uint16(769)
			goto next_state
		}
		return result
	case int32(233):
		if lookahead1 == int32('A') {
			state = uint16(545)
			goto next_state
		}
		return result
	case int32(234):
		if lookahead1 == int32('A') {
			state = uint16(548)
			goto next_state
		}
		return result
	case int32(235):
		if lookahead1 == int32('B') {
			state = uint16(297)
			goto next_state
		}
		return result
	case int32(236):
		if lookahead1 == int32('C') {
			state = uint16(303)
			goto next_state
		}
		return result
	case int32(237):
		if lookahead1 == int32('C') {
			state = uint16(264)
			goto next_state
		}
		return result
	case int32(238):
		if lookahead1 == int32('C') {
			state = uint16(257)
			goto next_state
		}
		return result
	case int32(239):
		if lookahead1 == int32('C') {
			state = uint16(246)
			goto next_state
		}
		return result
	case int32(240):
		if lookahead1 == int32('C') {
			state = uint16(610)
			goto next_state
		}
		return result
	case int32(241):
		if lookahead1 == int32('C') {
			state = uint16(495)
			goto next_state
		}
		return result
	case int32(242):
		if lookahead1 == int32('D') {
			state = uint16(913)
			goto next_state
		}
		return result
	case int32(243):
		if lookahead1 == int32('D') {
			state = uint16(466)
			goto next_state
		}
		return result
	case int32(244):
		if lookahead1 == int32('D') {
			state = uint16(454)
			goto next_state
		}
		return result
	case int32(245):
		if lookahead1 == int32('E') {
			state = uint16(303)
			goto next_state
		}
		return result
	case int32(246):
		if lookahead1 == int32('E') {
			state = uint16(913)
			goto next_state
		}
		return result
	case int32(247):
		if lookahead1 == int32('E') {
			state = uint16(236)
			goto next_state
		}
		return result
	case int32(248):
		if lookahead1 == int32('E') {
			state = uint16(307)
			goto next_state
		}
		return result
	case int32(249):
		if lookahead1 == int32('E') {
			state = uint16(681)
			goto next_state
		}
		return result
	case int32(250):
		if lookahead1 == int32('E') {
			state = uint16(584)
			goto next_state
		}
		return result
	case int32(251):
		if lookahead1 == int32('E') {
			state = uint16(333)
			goto next_state
		}
		if lookahead1 == int32('M') {
			state = uint16(335)
			goto next_state
		}
		return result
	case int32(252):
		if lookahead1 == int32('F') {
			state = uint16(516)
			goto next_state
		}
		return result
	case int32(253):
		if lookahead1 == int32('F') {
			state = uint16(356)
			goto next_state
		}
		return result
	case int32(254):
		if lookahead1 == int32('F') {
			state = uint16(356)
			goto next_state
		}
		if lookahead1 == int32('R') {
			state = uint16(456)
			goto next_state
		}
		return result
	case int32(255):
		if lookahead1 == int32('F') {
			state = uint16(636)
			goto next_state
		}
		return result
	case int32(256):
		if lookahead1 == int32('G') {
			state = uint16(357)
			goto next_state
		}
		if lookahead1 == int32('R') {
			state = uint16(408)
			goto next_state
		}
		return result
	case int32(257):
		if lookahead1 == int32('H') {
			state = uint16(913)
			goto next_state
		}
		return result
	case int32(258):
		if lookahead1 == int32('H') {
			state = uint16(287)
			goto next_state
		}
		return result
	case int32(259):
		if lookahead1 == int32('H') {
			state = uint16(464)
			goto next_state
		}
		if lookahead1 == int32('T') {
			state = uint16(522)
			goto next_state
		}
		return result
	case int32(260):
		if lookahead1 == int32('I') {
			state = uint16(278)
			goto next_state
		}
		return result
	case int32(261):
		if lookahead1 == int32('I') {
			state = uint16(205)
			goto next_state
		}
		return result
	case int32(262):
		if lookahead1 == int32('I') {
			state = uint16(574)
			goto next_state
		}
		return result
	case int32(263):
		if lookahead1 == int32('K') {
			state = uint16(923)
			goto next_state
		}
		return result
	case int32(264):
		if lookahead1 == int32('K') {
			state = uint16(245)
			goto next_state
		}
		return result
	case int32(265):
		if lookahead1 == int32('L') {
			state = uint16(248)
			goto next_state
		}
		return result
	case int32(266):
		if lookahead1 == int32('L') {
			state = uint16(913)
			goto next_state
		}
		return result
	case int32(267):
		if lookahead1 == int32('L') {
			state = uint16(429)
			goto next_state
		}
		return result
	case int32(268):
		if lookahead1 == int32('L') {
			state = uint16(339)
			goto next_state
		}
		return result
	case int32(269):
		if lookahead1 == int32('L') {
			state = uint16(619)
			goto next_state
		}
		return result
	case int32(270):
		if lookahead1 == int32('M') {
			state = uint16(465)
			goto next_state
		}
		return result
	case int32(271):
		if lookahead1 == int32('N') {
			state = uint16(272)
			goto next_state
		}
		return result
	case int32(272):
		if lookahead1 == int32('N') {
			state = uint16(247)
			goto next_state
		}
		return result
	case int32(273):
		if lookahead1 == int32('N') {
			state = uint16(296)
			goto next_state
		}
		return result
	case int32(274):
		if lookahead1 == int32('N') {
			state = uint16(639)
			goto next_state
		}
		return result
	case int32(275):
		if lookahead1 == int32('N') {
			state = uint16(426)
			goto next_state
		}
		return result
	case int32(276):
		if lookahead1 == int32('N') {
			state = uint16(640)
			goto next_state
		}
		return result
	case int32(277):
		if lookahead1 == int32('N') {
			state = uint16(641)
			goto next_state
		}
		return result
	case int32(278):
		if lookahead1 == int32('O') {
			state = uint16(273)
			goto next_state
		}
		return result
	case int32(279):
		if lookahead1 == int32('O') {
			state = uint16(237)
			goto next_state
		}
		return result
	case int32(280):
		if lookahead1 == int32('O') {
			state = uint16(718)
			goto next_state
		}
		return result
	case int32(281):
		if lookahead1 == int32('P') {
			state = uint16(258)
			goto next_state
		}
		return result
	case int32(282):
		if lookahead1 == int32('P') {
			state = uint16(226)
			goto next_state
		}
		return result
	case int32(283):
		if lookahead1 == int32('P') {
			state = uint16(175)
			goto next_state
		}
		return result
	case int32(284):
		if lookahead1 == int32('P') {
			state = uint16(477)
			goto next_state
		}
		return result
	case int32(285):
		if lookahead1 == int32('P') {
			state = uint16(676)
			goto next_state
		}
		return result
	case int32(286):
		if lookahead1 == int32('P') {
			state = uint16(678)
			goto next_state
		}
		return result
	case int32(287):
		if lookahead1 == int32('Q') {
			state = uint16(266)
			goto next_state
		}
		return result
	case int32(288):
		if lookahead1 == int32('R') {
			state = uint16(261)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(322)
			goto next_state
		}
		if lookahead1 == int32('p') {
			state = uint16(485)
			goto next_state
		}
		if lookahead1 == int32('s') {
			state = uint16(428)
			goto next_state
		}
		return result
	case int32(289):
		if lookahead1 == int32('R') {
			state = uint16(408)
			goto next_state
		}
		return result
	case int32(290):
		if lookahead1 == int32('R') {
			state = uint16(476)
			goto next_state
		}
		return result
	case int32(291):
		if lookahead1 == int32('R') {
			state = uint16(456)
			goto next_state
		}
		return result
	case int32(292):
		if lookahead1 == int32('R') {
			state = uint16(444)
			goto next_state
		}
		return result
	case int32(293):
		if lookahead1 == int32('R') {
			state = uint16(472)
			goto next_state
		}
		return result
	case int32(294):
		if lookahead1 == int32('R') {
			state = uint16(475)
			goto next_state
		}
		return result
	case int32(295):
		if lookahead1 == int32('S') {
			state = uint16(303)
			goto next_state
		}
		return result
	case int32(296):
		if lookahead1 == int32('S') {
			state = uint16(913)
			goto next_state
		}
		return result
	case int32(297):
		if lookahead1 == int32('S') {
			state = uint16(279)
			goto next_state
		}
		return result
	case int32(298):
		if lookahead1 == int32('S') {
			state = uint16(753)
			goto next_state
		}
		return result
	case int32(299):
		if lookahead1 == int32('S') {
			state = uint16(439)
			goto next_state
		}
		return result
	case int32(300):
		if lookahead1 == int32('S') {
			state = uint16(727)
			goto next_state
		}
		return result
	case int32(301):
		if lookahead1 == int32('S') {
			state = uint16(352)
			goto next_state
		}
		return result
	case int32(302):
		if lookahead1 == int32('S') {
			state = uint16(774)
			goto next_state
		}
		return result
	case int32(303):
		if lookahead1 == int32('T') {
			state = uint16(913)
			goto next_state
		}
		return result
	case int32(304):
		if lookahead1 == int32('T') {
			state = uint16(260)
			goto next_state
		}
		return result
	case int32(305):
		if lookahead1 == int32('T') {
			state = uint16(238)
			goto next_state
		}
		return result
	case int32(306):
		if lookahead1 == int32('T') {
			state = uint16(282)
			goto next_state
		}
		return result
	case int32(307):
		if lookahead1 == int32('T') {
			state = uint16(246)
			goto next_state
		}
		return result
	case int32(308):
		if lookahead1 == int32('T') {
			state = uint16(283)
			goto next_state
		}
		return result
	case int32(309):
		if lookahead1 == int32('T') {
			state = uint16(791)
			goto next_state
		}
		return result
	case int32(310):
		if lookahead1 == int32('T') {
			state = uint16(628)
			goto next_state
		}
		return result
	case int32(311):
		if lookahead1 == int32('T') {
			state = uint16(522)
			goto next_state
		}
		return result
	case int32(312):
		if lookahead1 == int32('T') {
			state = uint16(308)
			goto next_state
		}
		return result
	case int32(313):
		if lookahead1 == int32('T') {
			state = uint16(635)
			goto next_state
		}
		return result
	case int32(314):
		if lookahead1 == int32('U') {
			state = uint16(701)
			goto next_state
		}
		return result
	case int32(315):
		if lookahead1 == int32('U') {
			state = uint16(581)
			goto next_state
		}
		return result
	case int32(316):
		if lookahead1 == int32('V') {
			state = uint16(432)
			goto next_state
		}
		return result
	case int32(317):
		if lookahead1 == int32('\\') {
			state = uint16(891)
			goto next_state
		}
		if lookahead1 == int32('{') {
			state = uint16(884)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&aux_sym_WORD_CHAR_token1_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(475) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _31
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _31
	_31:
		if v4 != 0 {
			state = uint16(807)
			goto next_state
		}
		if v26 = !(eof != 0); v26 {
			v2 = uintptr(unsafe.Pointer(&aux_sym_PUNCTUATION_token1_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(484) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _35
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _35
		_35:
		}
		if v26 && v4 != 0 {
			state = uint16(863)
			goto next_state
		}
		return result
	case int32(318):
		if lookahead1 == int32('a') {
			state = uint16(386)
			goto next_state
		}
		return result
	case int32(319):
		if lookahead1 == int32('a') {
			state = uint16(502)
			goto next_state
		}
		if lookahead1 == int32('o') {
			state = uint16(664)
			goto next_state
		}
		return result
	case int32(320):
		if lookahead1 == int32('a') {
			state = uint16(719)
			goto next_state
		}
		if lookahead1 == int32('o') {
			state = uint16(577)
			goto next_state
		}
		return result
	case int32(321):
		if lookahead1 == int32('a') {
			state = uint16(669)
			goto next_state
		}
		if lookahead1 == int32('e') {
			state = uint16(667)
			goto next_state
		}
		if lookahead1 == int32('r') {
			state = uint16(409)
			goto next_state
		}
		return result
	case int32(322):
		if lookahead1 == int32('a') {
			state = uint16(768)
			goto next_state
		}
		if lookahead1 == int32('p') {
			state = uint16(671)
			goto next_state
		}
		if lookahead1 == int32('s') {
			state = uint16(763)
			goto next_state
		}
		return result
	case int32(323):
		if lookahead1 == int32('a') {
			state = uint16(364)
			goto next_state
		}
		return result
	case int32(324):
		if lookahead1 == int32('a') {
			state = uint16(787)
			goto next_state
		}
		return result
	case int32(325):
		if lookahead1 == int32('a') {
			state = uint16(666)
			goto next_state
		}
		return result
	case int32(326):
		if lookahead1 == int32('a') {
			state = uint16(779)
			goto next_state
		}
		return result
	case int32(327):
		if lookahead1 == int32('a') {
			state = uint16(363)
			goto next_state
		}
		return result
	case int32(328):
		if lookahead1 == int32('a') {
			state = uint16(397)
			goto next_state
		}
		return result
	case int32(329):
		if lookahead1 == int32('a') {
			state = uint16(573)
			goto next_state
		}
		if lookahead1 == int32('e') {
			state = uint16(659)
			goto next_state
		}
		return result
	case int32(330):
		if lookahead1 == int32('a') {
			state = uint16(552)
			goto next_state
		}
		return result
	case int32(331):
		if lookahead1 == int32('a') {
			state = uint16(400)
			goto next_state
		}
		return result
	case int32(332):
		if lookahead1 == int32('a') {
			state = uint16(486)
			goto next_state
		}
		return result
	case int32(333):
		if lookahead1 == int32('a') {
			state = uint16(674)
			goto next_state
		}
		return result
	case int32(334):
		if lookahead1 == int32('a') {
			state = uint16(693)
			goto next_state
		}
		return result
	case int32(335):
		if lookahead1 == int32('a') {
			state = uint16(596)
			goto next_state
		}
		return result
	case int32(336):
		if lookahead1 == int32('a') {
			state = uint16(209)
			goto next_state
		}
		return result
	case int32(337):
		if lookahead1 == int32('a') {
			state = uint16(708)
			goto next_state
		}
		return result
	case int32(338):
		if lookahead1 == int32('a') {
			state = uint16(722)
			goto next_state
		}
		return result
	case int32(339):
		if lookahead1 == int32('a') {
			state = uint16(672)
			goto next_state
		}
		return result
	case int32(340):
		if lookahead1 == int32('a') {
			state = uint16(208)
			goto next_state
		}
		return result
	case int32(341):
		if lookahead1 == int32('a') {
			state = uint16(507)
			goto next_state
		}
		return result
	case int32(342):
		if lookahead1 == int32('a') {
			state = uint16(725)
			goto next_state
		}
		return result
	case int32(343):
		if lookahead1 == int32('a') {
			state = uint16(601)
			goto next_state
		}
		return result
	case int32(344):
		if lookahead1 == int32('a') {
			state = uint16(649)
			goto next_state
		}
		return result
	case int32(345):
		if lookahead1 == int32('a') {
			state = uint16(726)
			goto next_state
		}
		return result
	case int32(346):
		if lookahead1 == int32('a') {
			state = uint16(554)
			goto next_state
		}
		return result
	case int32(347):
		if lookahead1 == int32('a') {
			state = uint16(523)
			goto next_state
		}
		return result
	case int32(348):
		if lookahead1 == int32('a') {
			state = uint16(789)
			goto next_state
		}
		return result
	case int32(349):
		if lookahead1 == int32('a') {
			state = uint16(587)
			goto next_state
		}
		return result
	case int32(350):
		if lookahead1 == int32('a') {
			state = uint16(556)
			goto next_state
		}
		return result
	case int32(351):
		if lookahead1 == int32('a') {
			state = uint16(395)
			goto next_state
		}
		return result
	case int32(352):
		if lookahead1 == int32('a') {
			state = uint16(732)
			goto next_state
		}
		return result
	case int32(353):
		if lookahead1 == int32('a') {
			state = uint16(403)
			goto next_state
		}
		return result
	case int32(354):
		if lookahead1 == int32('a') {
			state = uint16(735)
			goto next_state
		}
		return result
	case int32(355):
		if lookahead1 == int32('a') {
			state = uint16(742)
			goto next_state
		}
		return result
	case int32(356):
		if lookahead1 == int32('a') {
			state = uint16(529)
			goto next_state
		}
		return result
	case int32(357):
		if lookahead1 == int32('a') {
			state = uint16(748)
			goto next_state
		}
		return result
	case int32(358):
		if lookahead1 == int32('a') {
			state = uint16(598)
			goto next_state
		}
		return result
	case int32(359):
		if lookahead1 == int32('a') {
			state = uint16(365)
			goto next_state
		}
		return result
	case int32(360):
		if lookahead1 == int32('a') {
			state = uint16(759)
			goto next_state
		}
		return result
	case int32(361):
		if lookahead1 == int32('a') {
			state = uint16(760)
			goto next_state
		}
		return result
	case int32(362):
		if lookahead1 == int32('b') {
			state = uint16(503)
			goto next_state
		}
		return result
	case int32(363):
		if lookahead1 == int32('b') {
			state = uint16(543)
			goto next_state
		}
		return result
	case int32(364):
		if lookahead1 == int32('b') {
			state = uint16(558)
			goto next_state
		}
		return result
	case int32(365):
		if lookahead1 == int32('b') {
			state = uint16(559)
			goto next_state
		}
		return result
	case int32(366):
		if lookahead1 == int32('c') {
			state = uint16(369)
			goto next_state
		}
		if lookahead1 == int32('l') {
			state = uint16(668)
			goto next_state
		}
		return result
	case int32(367):
		if lookahead1 == int32('c') {
			state = uint16(538)
			goto next_state
		}
		if lookahead1 == int32('o') {
			state = uint16(648)
			goto next_state
		}
		return result
	case int32(368):
		if lookahead1 == int32('c') {
			state = uint16(787)
			goto next_state
		}
		return result
	case int32(369):
		if lookahead1 == int32('c') {
			state = uint16(419)
			goto next_state
		}
		return result
	case int32(370):
		if lookahead1 == int32('c') {
			state = uint16(492)
			goto next_state
		}
		return result
	case int32(371):
		if lookahead1 == int32('c') {
			state = uint16(724)
			goto next_state
		}
		return result
	case int32(372):
		if lookahead1 == int32('c') {
			state = uint16(712)
			goto next_state
		}
		return result
	case int32(373):
		if lookahead1 == int32('c') {
			state = uint16(418)
			goto next_state
		}
		if lookahead1 == int32('x') {
			state = uint16(788)
			goto next_state
		}
		return result
	case int32(374):
		if lookahead1 == int32('c') {
			state = uint16(537)
			goto next_state
		}
		return result
	case int32(375):
		if lookahead1 == int32('c') {
			state = uint16(455)
			goto next_state
		}
		return result
	case int32(376):
		if lookahead1 == int32('c') {
			state = uint16(612)
			goto next_state
		}
		return result
	case int32(377):
		if lookahead1 == int32('c') {
			state = uint16(431)
			goto next_state
		}
		return result
	case int32(378):
		if lookahead1 == int32('c') {
			state = uint16(725)
			goto next_state
		}
		return result
	case int32(379):
		if lookahead1 == int32('c') {
			state = uint16(425)
			goto next_state
		}
		return result
	case int32(380):
		if lookahead1 == int32('c') {
			state = uint16(611)
			goto next_state
		}
		return result
	case int32(381):
		if lookahead1 == int32('c') {
			state = uint16(459)
			goto next_state
		}
		return result
	case int32(382):
		if lookahead1 == int32('c') {
			state = uint16(755)
			goto next_state
		}
		return result
	case int32(383):
		if lookahead1 == int32('c') {
			state = uint16(379)
			goto next_state
		}
		return result
	case int32(384):
		if lookahead1 == int32('c') {
			state = uint16(360)
			goto next_state
		}
		return result
	case int32(385):
		if lookahead1 == int32('d') {
			state = uint16(923)
			goto next_state
		}
		return result
	case int32(386):
		if lookahead1 == int32('d') {
			state = uint16(172)
			goto next_state
		}
		return result
	case int32(387):
		if lookahead1 == int32('d') {
			state = uint16(518)
			goto next_state
		}
		return result
	case int32(388):
		if lookahead1 == int32('d') {
			state = uint16(393)
			goto next_state
		}
		return result
	case int32(389):
		if lookahead1 == int32('d') {
			state = uint16(415)
			goto next_state
		}
		return result
	case int32(390):
		if lookahead1 == int32('d') {
			state = uint16(519)
			goto next_state
		}
		return result
	case int32(391):
		if lookahead1 == int32('d') {
			state = uint16(710)
			goto next_state
		}
		return result
	case int32(392):
		if lookahead1 == int32('d') {
			state = uint16(199)
			goto next_state
		}
		return result
	case int32(393):
		if lookahead1 == int32('d') {
			state = uint16(436)
			goto next_state
		}
		return result
	case int32(394):
		if lookahead1 == int32('d') {
			state = uint16(179)
			goto next_state
		}
		return result
	case int32(395):
		if lookahead1 == int32('d') {
			state = uint16(442)
			goto next_state
		}
		return result
	case int32(396):
		if lookahead1 == int32('d') {
			state = uint16(447)
			goto next_state
		}
		return result
	case int32(397):
		if lookahead1 == int32('d') {
			state = uint16(790)
			goto next_state
		}
		return result
	case int32(398):
		if lookahead1 == int32('d') {
			state = uint16(535)
			goto next_state
		}
		return result
	case int32(399):
		if lookahead1 == int32('d') {
			state = uint16(211)
			goto next_state
		}
		return result
	case int32(400):
		if lookahead1 == int32('d') {
			state = uint16(443)
			goto next_state
		}
		return result
	case int32(401):
		if lookahead1 == int32('d') {
			state = uint16(521)
			goto next_state
		}
		return result
	case int32(402):
		if lookahead1 == int32('d') {
			state = uint16(528)
			goto next_state
		}
		return result
	case int32(403):
		if lookahead1 == int32('d') {
			state = uint16(219)
			goto next_state
		}
		return result
	case int32(404):
		if lookahead1 == int32('d') {
			state = uint16(218)
			goto next_state
		}
		return result
	case int32(405):
		if lookahead1 == int32('d') {
			state = uint16(220)
			goto next_state
		}
		return result
	case int32(406):
		if lookahead1 == int32('e') {
			state = uint16(923)
			goto next_state
		}
		return result
	case int32(407):
		if lookahead1 == int32('e') {
			state = uint16(422)
			goto next_state
		}
		if lookahead1 == int32('w') {
			state = uint16(506)
			goto next_state
		}
		return result
	case int32(408):
		if lookahead1 == int32('e') {
			state = uint16(661)
			goto next_state
		}
		return result
	case int32(409):
		if lookahead1 == int32('e') {
			state = uint16(380)
			goto next_state
		}
		if lookahead1 == int32('o') {
			state = uint16(373)
			goto next_state
		}
		return result
	case int32(410):
		if lookahead1 == int32('e') {
			state = uint16(713)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(698)
			goto next_state
		}
		if lookahead1 == int32('o') {
			state = uint16(776)
			goto next_state
		}
		if lookahead1 == int32('u') {
			state = uint16(541)
			goto next_state
		}
		return result
	case int32(411):
		if lookahead1 == int32('e') {
			state = uint16(561)
			goto next_state
		}
		if lookahead1 == int32('o') {
			state = uint16(607)
			goto next_state
		}
		return result
	case int32(412):
		if lookahead1 == int32('e') {
			state = uint16(781)
			goto next_state
		}
		return result
	case int32(413):
		if lookahead1 == int32('e') {
			state = uint16(342)
			goto next_state
		}
		return result
	case int32(414):
		if lookahead1 == int32('e') {
			state = uint16(371)
			goto next_state
		}
		return result
	case int32(415):
		if lookahead1 == int32('e') {
			state = uint16(385)
			goto next_state
		}
		return result
	case int32(416):
		if lookahead1 == int32('e') {
			state = uint16(571)
			goto next_state
		}
		if lookahead1 == int32('o') {
			state = uint16(367)
			goto next_state
		}
		return result
	case int32(417):
		if lookahead1 == int32('e') {
			state = uint16(714)
			goto next_state
		}
		if lookahead1 == int32('o') {
			state = uint16(201)
			goto next_state
		}
		return result
	case int32(418):
		if lookahead1 == int32('e') {
			state = uint16(699)
			goto next_state
		}
		return result
	case int32(419):
		if lookahead1 == int32('e') {
			state = uint16(650)
			goto next_state
		}
		return result
	case int32(420):
		if lookahead1 == int32('e') {
			state = uint16(565)
			goto next_state
		}
		return result
	case int32(421):
		if lookahead1 == int32('e') {
			state = uint16(706)
			goto next_state
		}
		return result
	case int32(422):
		if lookahead1 == int32('e') {
			state = uint16(187)
			goto next_state
		}
		if lookahead1 == int32('r') {
			state = uint16(775)
			goto next_state
		}
		return result
	case int32(423):
		if lookahead1 == int32('e') {
			state = uint16(673)
			goto next_state
		}
		return result
	case int32(424):
		if lookahead1 == int32('e') {
			state = uint16(382)
			goto next_state
		}
		return result
	case int32(425):
		if lookahead1 == int32('e') {
			state = uint16(654)
			goto next_state
		}
		return result
	case int32(426):
		if lookahead1 == int32('e') {
			state = uint16(488)
			goto next_state
		}
		return result
	case int32(427):
		if lookahead1 == int32('e') {
			state = uint16(729)
			goto next_state
		}
		return result
	case int32(428):
		if lookahead1 == int32('e') {
			state = uint16(191)
			goto next_state
		}
		return result
	case int32(429):
		if lookahead1 == int32('e') {
			state = uint16(490)
			goto next_state
		}
		return result
	case int32(430):
		if lookahead1 == int32('e') {
			state = uint16(591)
			goto next_state
		}
		return result
	case int32(431):
		if lookahead1 == int32('e') {
			state = uint16(696)
			goto next_state
		}
		return result
	case int32(432):
		if lookahead1 == int32('e') {
			state = uint16(670)
			goto next_state
		}
		return result
	case int32(433):
		if lookahead1 == int32('e') {
			state = uint16(344)
			goto next_state
		}
		return result
	case int32(434):
		if lookahead1 == int32('e') {
			state = uint16(544)
			goto next_state
		}
		return result
	case int32(435):
		if lookahead1 == int32('e') {
			state = uint16(663)
			goto next_state
		}
		return result
	case int32(436):
		if lookahead1 == int32('e') {
			state = uint16(569)
			goto next_state
		}
		return result
	case int32(437):
		if lookahead1 == int32('e') {
			state = uint16(580)
			goto next_state
		}
		return result
	case int32(438):
		if lookahead1 == int32('e') {
			state = uint16(189)
			goto next_state
		}
		return result
	case int32(439):
		if lookahead1 == int32('e') {
			state = uint16(686)
			goto next_state
		}
		return result
	case int32(440):
		if lookahead1 == int32('e') {
			state = uint16(590)
			goto next_state
		}
		return result
	case int32(441):
		if lookahead1 == int32('e') {
			state = uint16(632)
			goto next_state
		}
		return result
	case int32(442):
		if lookahead1 == int32('e') {
			state = uint16(682)
			goto next_state
		}
		return result
	case int32(443):
		if lookahead1 == int32('e') {
			state = uint16(176)
			goto next_state
		}
		return result
	case int32(444):
		if lookahead1 == int32('e') {
			state = uint16(337)
			goto next_state
		}
		return result
	case int32(445):
		if lookahead1 == int32('e') {
			state = uint16(595)
			goto next_state
		}
		return result
	case int32(446):
		if lookahead1 == int32('e') {
			state = uint16(582)
			goto next_state
		}
		return result
	case int32(447):
		if lookahead1 == int32('e') {
			state = uint16(578)
			goto next_state
		}
		return result
	case int32(448):
		if lookahead1 == int32('e') {
			state = uint16(204)
			goto next_state
		}
		return result
	case int32(449):
		if lookahead1 == int32('e') {
			state = uint16(184)
			goto next_state
		}
		return result
	case int32(450):
		if lookahead1 == int32('e') {
			state = uint16(198)
			goto next_state
		}
		return result
	case int32(451):
		if lookahead1 == int32('e') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(452):
		if lookahead1 == int32('e') {
			state = uint16(328)
			goto next_state
		}
		return result
	case int32(453):
		if lookahead1 == int32('e') {
			state = uint16(392)
			goto next_state
		}
		return result
	case int32(454):
		if lookahead1 == int32('e') {
			state = uint16(656)
			goto next_state
		}
		return result
	case int32(455):
		if lookahead1 == int32('e') {
			state = uint16(182)
			goto next_state
		}
		return result
	case int32(456):
		if lookahead1 == int32('e') {
			state = uint16(660)
			goto next_state
		}
		return result
	case int32(457):
		if lookahead1 == int32('e') {
			state = uint16(783)
			goto next_state
		}
		return result
	case int32(458):
		if lookahead1 == int32('e') {
			state = uint16(372)
			goto next_state
		}
		return result
	case int32(459):
		if lookahead1 == int32('e') {
			state = uint16(700)
			goto next_state
		}
		return result
	case int32(460):
		if lookahead1 == int32('e') {
			state = uint16(702)
			goto next_state
		}
		return result
	case int32(461):
		if lookahead1 == int32('e') {
			state = uint16(683)
			goto next_state
		}
		return result
	case int32(462):
		if lookahead1 == int32('e') {
			state = uint16(399)
			goto next_state
		}
		return result
	case int32(463):
		if lookahead1 == int32('e') {
			state = uint16(703)
			goto next_state
		}
		return result
	case int32(464):
		if lookahead1 == int32('e') {
			state = uint16(351)
			goto next_state
		}
		return result
	case int32(465):
		if lookahead1 == int32('e') {
			state = uint16(401)
			goto next_state
		}
		return result
	case int32(466):
		if lookahead1 == int32('e') {
			state = uint16(754)
			goto next_state
		}
		return result
	case int32(467):
		if lookahead1 == int32('e') {
			state = uint16(405)
			goto next_state
		}
		return result
	case int32(468):
		if lookahead1 == int32('e') {
			state = uint16(394)
			goto next_state
		}
		return result
	case int32(469):
		if lookahead1 == int32('e') {
			state = uint16(378)
			goto next_state
		}
		return result
	case int32(470):
		if lookahead1 == int32('e') {
			state = uint16(588)
			goto next_state
		}
		return result
	case int32(471):
		if lookahead1 == int32('e') {
			state = uint16(597)
			goto next_state
		}
		return result
	case int32(472):
		if lookahead1 == int32('e') {
			state = uint16(657)
			goto next_state
		}
		return result
	case int32(473):
		if lookahead1 == int32('e') {
			state = uint16(599)
			goto next_state
		}
		return result
	case int32(474):
		if lookahead1 == int32('e') {
			state = uint16(600)
			goto next_state
		}
		return result
	case int32(475):
		if lookahead1 == int32('e') {
			state = uint16(662)
			goto next_state
		}
		return result
	case int32(476):
		if lookahead1 == int32('e') {
			state = uint16(402)
			goto next_state
		}
		return result
	case int32(477):
		if lookahead1 == int32('e') {
			state = uint16(694)
			goto next_state
		}
		return result
	case int32(478):
		if lookahead1 == int32('f') {
			state = uint16(480)
			goto next_state
		}
		return result
	case int32(479):
		if lookahead1 == int32('f') {
			state = uint16(546)
			goto next_state
		}
		if lookahead1 == int32('t') {
			state = uint16(505)
			goto next_state
		}
		return result
	case int32(480):
		if lookahead1 == int32('f') {
			state = uint16(504)
			goto next_state
		}
		return result
	case int32(481):
		if lookahead1 == int32('f') {
			state = uint16(512)
			goto next_state
		}
		return result
	case int32(482):
		if lookahead1 == int32('f') {
			state = uint16(643)
			goto next_state
		}
		return result
	case int32(483):
		if lookahead1 == int32('f') {
			state = uint16(513)
			goto next_state
		}
		return result
	case int32(484):
		if lookahead1 == int32('g') {
			state = uint16(923)
			goto next_state
		}
		return result
	case int32(485):
		if lookahead1 == int32('g') {
			state = uint16(689)
			goto next_state
		}
		return result
	case int32(486):
		if lookahead1 == int32('g') {
			state = uint16(406)
			goto next_state
		}
		return result
	case int32(487):
		if lookahead1 == int32('g') {
			state = uint16(715)
			goto next_state
		}
		return result
	case int32(488):
		if lookahead1 == int32('g') {
			state = uint16(638)
			goto next_state
		}
		return result
	case int32(489):
		if lookahead1 == int32('g') {
			state = uint16(438)
			goto next_state
		}
		return result
	case int32(490):
		if lookahead1 == int32('g') {
			state = uint16(350)
			goto next_state
		}
		return result
	case int32(491):
		if lookahead1 == int32('g') {
			state = uint16(217)
			goto next_state
		}
		return result
	case int32(492):
		if lookahead1 == int32('h') {
			state = uint16(192)
			goto next_state
		}
		return result
	case int32(493):
		if lookahead1 == int32('h') {
			state = uint16(637)
			goto next_state
		}
		return result
	case int32(494):
		if lookahead1 == int32('h') {
			state = uint16(631)
			goto next_state
		}
		return result
	case int32(495):
		if lookahead1 == int32('h') {
			state = uint16(644)
			goto next_state
		}
		return result
	case int32(496):
		if lookahead1 == int32('h') {
			state = uint16(176)
			goto next_state
		}
		return result
	case int32(497):
		if lookahead1 == int32('h') {
			state = uint16(435)
			goto next_state
		}
		return result
	case int32(498):
		if lookahead1 == int32('h') {
			state = uint16(445)
			goto next_state
		}
		return result
	case int32(499):
		if lookahead1 == int32('h') {
			state = uint16(617)
			goto next_state
		}
		return result
	case int32(500):
		if lookahead1 == int32('i') {
			state = uint16(225)
			goto next_state
		}
		return result
	case int32(501):
		if lookahead1 == int32('i') {
			state = uint16(794)
			goto next_state
		}
		return result
	case int32(502):
		if lookahead1 == int32('i') {
			state = uint16(560)
			goto next_state
		}
		return result
	case int32(503):
		if lookahead1 == int32('i') {
			state = uint16(388)
			goto next_state
		}
		return result
	case int32(504):
		if lookahead1 == int32('i') {
			state = uint16(374)
			goto next_state
		}
		return result
	case int32(505):
		if lookahead1 == int32('i') {
			state = uint16(589)
			goto next_state
		}
		return result
	case int32(506):
		if lookahead1 == int32('i') {
			state = uint16(717)
			goto next_state
		}
		return result
	case int32(507):
		if lookahead1 == int32('i') {
			state = uint16(549)
			goto next_state
		}
		return result
	case int32(508):
		if lookahead1 == int32('i') {
			state = uint16(358)
			goto next_state
		}
		return result
	case int32(509):
		if lookahead1 == int32('i') {
			state = uint16(384)
			goto next_state
		}
		return result
	case int32(510):
		if lookahead1 == int32('i') {
			state = uint16(330)
			goto next_state
		}
		return result
	case int32(511):
		if lookahead1 == int32('i') {
			state = uint16(705)
			goto next_state
		}
		return result
	case int32(512):
		if lookahead1 == int32('i') {
			state = uint16(415)
			goto next_state
		}
		return result
	case int32(513):
		if lookahead1 == int32('i') {
			state = uint16(327)
			goto next_state
		}
		return result
	case int32(514):
		if lookahead1 == int32('i') {
			state = uint16(576)
			goto next_state
		}
		return result
	case int32(515):
		if lookahead1 == int32('i') {
			state = uint16(716)
			goto next_state
		}
		return result
	case int32(516):
		if lookahead1 == int32('i') {
			state = uint16(434)
			goto next_state
		}
		return result
	case int32(517):
		if lookahead1 == int32('i') {
			state = uint16(375)
			goto next_state
		}
		return result
	case int32(518):
		if lookahead1 == int32('i') {
			state = uint16(680)
			goto next_state
		}
		return result
	case int32(519):
		if lookahead1 == int32('i') {
			state = uint16(481)
			goto next_state
		}
		return result
	case int32(520):
		if lookahead1 == int32('i') {
			state = uint16(372)
			goto next_state
		}
		return result
	case int32(521):
		if lookahead1 == int32('i') {
			state = uint16(340)
			goto next_state
		}
		return result
	case int32(522):
		if lookahead1 == int32('i') {
			state = uint16(566)
			goto next_state
		}
		return result
	case int32(523):
		if lookahead1 == int32('i') {
			state = uint16(551)
			goto next_state
		}
		return result
	case int32(524):
		if lookahead1 == int32('i') {
			state = uint16(675)
			goto next_state
		}
		return result
	case int32(525):
		if lookahead1 == int32('i') {
			state = uint16(778)
			goto next_state
		}
		return result
	case int32(526):
		if lookahead1 == int32('i') {
			state = uint16(620)
			goto next_state
		}
		return result
	case int32(527):
		if lookahead1 == int32('i') {
			state = uint16(377)
			goto next_state
		}
		return result
	case int32(528):
		if lookahead1 == int32('i') {
			state = uint16(688)
			goto next_state
		}
		return result
	case int32(529):
		if lookahead1 == int32('i') {
			state = uint16(550)
			goto next_state
		}
		return result
	case int32(530):
		if lookahead1 == int32('i') {
			state = uint16(622)
			goto next_state
		}
		return result
	case int32(531):
		if lookahead1 == int32('i') {
			state = uint16(633)
			goto next_state
		}
		return result
	case int32(532):
		if lookahead1 == int32('i') {
			state = uint16(624)
			goto next_state
		}
		return result
	case int32(533):
		if lookahead1 == int32('i') {
			state = uint16(616)
			goto next_state
		}
		return result
	case int32(534):
		if lookahead1 == int32('i') {
			state = uint16(355)
			goto next_state
		}
		return result
	case int32(535):
		if lookahead1 == int32('i') {
			state = uint16(757)
			goto next_state
		}
		return result
	case int32(536):
		if lookahead1 == int32('i') {
			state = uint16(756)
			goto next_state
		}
		return result
	case int32(537):
		if lookahead1 == int32('i') {
			state = uint16(474)
			goto next_state
		}
		return result
	case int32(538):
		if lookahead1 == int32('k') {
			state = uint16(415)
			goto next_state
		}
		return result
	case int32(539):
		if lookahead1 == int32('k') {
			state = uint16(212)
			goto next_state
		}
		return result
	case int32(540):
		if lookahead1 == int32('l') {
			state = uint16(787)
			goto next_state
		}
		return result
	case int32(541):
		if lookahead1 == int32('l') {
			state = uint16(720)
			goto next_state
		}
		return result
	case int32(542):
		if lookahead1 == int32('l') {
			state = uint16(642)
			goto next_state
		}
		if lookahead1 == int32('m') {
			state = uint16(430)
			goto next_state
		}
		return result
	case int32(543):
		if lookahead1 == int32('l') {
			state = uint16(406)
			goto next_state
		}
		return result
	case int32(544):
		if lookahead1 == int32('l') {
			state = uint16(391)
			goto next_state
		}
		return result
	case int32(545):
		if lookahead1 == int32('l') {
			state = uint16(707)
			goto next_state
		}
		return result
	case int32(546):
		if lookahead1 == int32('l') {
			state = uint16(520)
			goto next_state
		}
		return result
	case int32(547):
		if lookahead1 == int32('l') {
			state = uint16(696)
			goto next_state
		}
		return result
	case int32(548):
		if lookahead1 == int32('l') {
			state = uint16(553)
			goto next_state
		}
		return result
	case int32(549):
		if lookahead1 == int32('l') {
			state = uint16(323)
			goto next_state
		}
		return result
	case int32(550):
		if lookahead1 == int32('l') {
			state = uint16(415)
			goto next_state
		}
		return result
	case int32(551):
		if lookahead1 == int32('l') {
			state = uint16(327)
			goto next_state
		}
		return result
	case int32(552):
		if lookahead1 == int32('l') {
			state = uint16(200)
			goto next_state
		}
		return result
	case int32(553):
		if lookahead1 == int32('l') {
			state = uint16(605)
			goto next_state
		}
		return result
	case int32(554):
		if lookahead1 == int32('l') {
			state = uint16(197)
			goto next_state
		}
		return result
	case int32(555):
		if lookahead1 == int32('l') {
			state = uint16(420)
			goto next_state
		}
		return result
	case int32(556):
		if lookahead1 == int32('l') {
			state = uint16(183)
			goto next_state
		}
		return result
	case int32(557):
		if lookahead1 == int32('l') {
			state = uint16(448)
			goto next_state
		}
		return result
	case int32(558):
		if lookahead1 == int32('l') {
			state = uint16(449)
			goto next_state
		}
		return result
	case int32(559):
		if lookahead1 == int32('l') {
			state = uint16(450)
			goto next_state
		}
		return result
	case int32(560):
		if lookahead1 == int32('l') {
			state = uint16(462)
			goto next_state
		}
		return result
	case int32(561):
		if lookahead1 == int32('m') {
			state = uint16(652)
			goto next_state
		}
		return result
	case int32(562):
		if lookahead1 == int32('m') {
			state = uint16(193)
			goto next_state
		}
		return result
	case int32(563):
		if lookahead1 == int32('m') {
			state = uint16(343)
			goto next_state
		}
		return result
	case int32(564):
		if lookahead1 == int32('m') {
			state = uint16(651)
			goto next_state
		}
		return result
	case int32(565):
		if lookahead1 == int32('m') {
			state = uint16(470)
			goto next_state
		}
		return result
	case int32(566):
		if lookahead1 == int32('m') {
			state = uint16(441)
			goto next_state
		}
		return result
	case int32(567):
		if lookahead1 == int32('m') {
			state = uint16(349)
			goto next_state
		}
		return result
	case int32(568):
		if lookahead1 == int32('m') {
			state = uint16(361)
			goto next_state
		}
		return result
	case int32(569):
		if lookahead1 == int32('n') {
			state = uint16(923)
			goto next_state
		}
		return result
	case int32(570):
		if lookahead1 == int32('n') {
			state = uint16(479)
			goto next_state
		}
		return result
	case int32(571):
		if lookahead1 == int32('n') {
			state = uint16(487)
			goto next_state
		}
		return result
	case int32(572):
		if lookahead1 == int32('n') {
			state = uint16(385)
			goto next_state
		}
		return result
	case int32(573):
		if lookahead1 == int32('n') {
			state = uint16(489)
			goto next_state
		}
		return result
	case int32(574):
		if lookahead1 == int32('n') {
			state = uint16(482)
			goto next_state
		}
		return result
	case int32(575):
		if lookahead1 == int32('n') {
			state = uint16(491)
			goto next_state
		}
		return result
	case int32(576):
		if lookahead1 == int32('n') {
			state = uint16(484)
			goto next_state
		}
		return result
	case int32(577):
		if lookahead1 == int32('n') {
			state = uint16(406)
			goto next_state
		}
		return result
	case int32(578):
		if lookahead1 == int32('n') {
			state = uint16(368)
			goto next_state
		}
		return result
	case int32(579):
		if lookahead1 == int32('n') {
			state = uint16(696)
			goto next_state
		}
		return result
	case int32(580):
		if lookahead1 == int32('n') {
			state = uint16(712)
			goto next_state
		}
		return result
	case int32(581):
		if lookahead1 == int32('n') {
			state = uint16(326)
			goto next_state
		}
		return result
	case int32(582):
		if lookahead1 == int32('n') {
			state = uint16(721)
			goto next_state
		}
		return result
	case int32(583):
		if lookahead1 == int32('n') {
			state = uint16(176)
			goto next_state
		}
		return result
	case int32(584):
		if lookahead1 == int32('n') {
			state = uint16(752)
			goto next_state
		}
		return result
	case int32(585):
		if lookahead1 == int32('n') {
			state = uint16(178)
			goto next_state
		}
		return result
	case int32(586):
		if lookahead1 == int32('n') {
			state = uint16(180)
			goto next_state
		}
		return result
	case int32(587):
		if lookahead1 == int32('n') {
			state = uint16(446)
			goto next_state
		}
		return result
	case int32(588):
		if lookahead1 == int32('n') {
			state = uint16(725)
			goto next_state
		}
		return result
	case int32(589):
		if lookahead1 == int32('n') {
			state = uint16(764)
			goto next_state
		}
		return result
	case int32(590):
		if lookahead1 == int32('n') {
			state = uint16(389)
			goto next_state
		}
		return result
	case int32(591):
		if lookahead1 == int32('n') {
			state = uint16(730)
			goto next_state
		}
		return result
	case int32(592):
		if lookahead1 == int32('n') {
			state = uint16(346)
			goto next_state
		}
		return result
	case int32(593):
		if lookahead1 == int32('n') {
			state = uint16(738)
			goto next_state
		}
		return result
	case int32(594):
		if lookahead1 == int32('n') {
			state = uint16(398)
			goto next_state
		}
		return result
	case int32(595):
		if lookahead1 == int32('n') {
			state = uint16(731)
			goto next_state
		}
		return result
	case int32(596):
		if lookahead1 == int32('n') {
			state = uint16(793)
			goto next_state
		}
		return result
	case int32(597):
		if lookahead1 == int32('n') {
			state = uint16(396)
			goto next_state
		}
		return result
	case int32(598):
		if lookahead1 == int32('n') {
			state = uint16(736)
			goto next_state
		}
		return result
	case int32(599):
		if lookahead1 == int32('n') {
			state = uint16(751)
			goto next_state
		}
		return result
	case int32(600):
		if lookahead1 == int32('n') {
			state = uint16(741)
			goto next_state
		}
		return result
	case int32(601):
		if lookahead1 == int32('n') {
			state = uint16(473)
			goto next_state
		}
		return result
	case int32(602):
		if lookahead1 == int32('n') {
			state = uint16(221)
			goto next_state
		}
		return result
	case int32(603):
		if lookahead1 == int32('o') {
			state = uint16(570)
			goto next_state
		}
		if lookahead1 == int32('r') {
			state = uint16(413)
			goto next_state
		}
		return result
	case int32(604):
		if lookahead1 == int32('o') {
			state = uint16(785)
			goto next_state
		}
		return result
	case int32(605):
		if lookahead1 == int32('o') {
			state = uint16(782)
			goto next_state
		}
		return result
	case int32(606):
		if lookahead1 == int32('o') {
			state = uint16(762)
			goto next_state
		}
		return result
	case int32(607):
		if lookahead1 == int32('o') {
			state = uint16(174)
			goto next_state
		}
		return result
	case int32(608):
		if lookahead1 == int32('o') {
			state = uint16(665)
			goto next_state
		}
		return result
	case int32(609):
		if lookahead1 == int32('o') {
			state = uint16(690)
			goto next_state
		}
		return result
	case int32(610):
		if lookahead1 == int32('o') {
			state = uint16(593)
			goto next_state
		}
		return result
	case int32(611):
		if lookahead1 == int32('o') {
			state = uint16(594)
			goto next_state
		}
		return result
	case int32(612):
		if lookahead1 == int32('o') {
			state = uint16(547)
			goto next_state
		}
		return result
	case int32(613):
		if lookahead1 == int32('o') {
			state = uint16(695)
			goto next_state
		}
		return result
	case int32(614):
		if lookahead1 == int32('o') {
			state = uint16(712)
			goto next_state
		}
		return result
	case int32(615):
		if lookahead1 == int32('o') {
			state = uint16(663)
			goto next_state
		}
		return result
	case int32(616):
		if lookahead1 == int32('o') {
			state = uint16(569)
			goto next_state
		}
		return result
	case int32(617):
		if lookahead1 == int32('o') {
			state = uint16(691)
			goto next_state
		}
		return result
	case int32(618):
		if lookahead1 == int32('o') {
			state = uint16(734)
			goto next_state
		}
		return result
	case int32(619):
		if lookahead1 == int32('o') {
			state = uint16(576)
			goto next_state
		}
		return result
	case int32(620):
		if lookahead1 == int32('o') {
			state = uint16(585)
			goto next_state
		}
		return result
	case int32(621):
		if lookahead1 == int32('o') {
			state = uint16(679)
			goto next_state
		}
		return result
	case int32(622):
		if lookahead1 == int32('o') {
			state = uint16(602)
			goto next_state
		}
		return result
	case int32(623):
		if lookahead1 == int32('o') {
			state = uint16(185)
			goto next_state
		}
		return result
	case int32(624):
		if lookahead1 == int32('o') {
			state = uint16(583)
			goto next_state
		}
		return result
	case int32(625):
		if lookahead1 == int32('o') {
			state = uint16(579)
			goto next_state
		}
		return result
	case int32(626):
		if lookahead1 == int32('o') {
			state = uint16(190)
			goto next_state
		}
		return result
	case int32(627):
		if lookahead1 == int32('o') {
			state = uint16(194)
			goto next_state
		}
		return result
	case int32(628):
		if lookahead1 == int32('o') {
			state = uint16(623)
			goto next_state
		}
		return result
	case int32(629):
		if lookahead1 == int32('o') {
			state = uint16(376)
			goto next_state
		}
		return result
	case int32(630):
		if lookahead1 == int32('o') {
			state = uint16(390)
			goto next_state
		}
		return result
	case int32(631):
		if lookahead1 == int32('o') {
			state = uint16(677)
			goto next_state
		}
		return result
	case int32(632):
		if lookahead1 == int32('o') {
			state = uint16(766)
			goto next_state
		}
		return result
	case int32(633):
		if lookahead1 == int32('o') {
			state = uint16(586)
			goto next_state
		}
		return result
	case int32(634):
		if lookahead1 == int32('o') {
			state = uint16(381)
			goto next_state
		}
		return result
	case int32(635):
		if lookahead1 == int32('o') {
			state = uint16(626)
			goto next_state
		}
		return result
	case int32(636):
		if lookahead1 == int32('o') {
			state = uint16(684)
			goto next_state
		}
		return result
	case int32(637):
		if lookahead1 == int32('o') {
			state = uint16(404)
			goto next_state
		}
		return result
	case int32(638):
		if lookahead1 == int32('o') {
			state = uint16(737)
			goto next_state
		}
		return result
	case int32(639):
		if lookahead1 == int32('o') {
			state = uint16(740)
			goto next_state
		}
		return result
	case int32(640):
		if lookahead1 == int32('o') {
			state = uint16(749)
			goto next_state
		}
		return result
	case int32(641):
		if lookahead1 == int32('o') {
			state = uint16(743)
			goto next_state
		}
		return result
	case int32(642):
		if lookahead1 == int32('o') {
			state = uint16(353)
			goto next_state
		}
		return result
	case int32(643):
		if lookahead1 == int32('o') {
			state = uint16(687)
			goto next_state
		}
		return result
	case int32(644):
		if lookahead1 == int32('o') {
			state = uint16(527)
			goto next_state
		}
		return result
	case int32(645):
		if lookahead1 == int32('o') {
			state = uint16(685)
			goto next_state
		}
		return result
	case int32(646):
		if lookahead1 == int32('p') {
			state = uint16(414)
			goto next_state
		}
		return result
	case int32(647):
		if lookahead1 == int32('p') {
			state = uint16(406)
			goto next_state
		}
		return result
	case int32(648):
		if lookahead1 == int32('p') {
			state = uint16(181)
			goto next_state
		}
		return result
	case int32(649):
		if lookahead1 == int32('p') {
			state = uint16(614)
			goto next_state
		}
		return result
	case int32(650):
		if lookahead1 == int32('p') {
			state = uint16(725)
			goto next_state
		}
		return result
	case int32(651):
		if lookahead1 == int32('p') {
			state = uint16(555)
			goto next_state
		}
		return result
	case int32(652):
		if lookahead1 == int32('p') {
			state = uint16(609)
			goto next_state
		}
		return result
	case int32(653):
		if lookahead1 == int32('p') {
			state = uint16(613)
			goto next_state
		}
		return result
	case int32(654):
		if lookahead1 == int32('p') {
			state = uint16(728)
			goto next_state
		}
		return result
	case int32(655):
		if lookahead1 == int32('p') {
			state = uint16(653)
			goto next_state
		}
		return result
	case int32(656):
		if lookahead1 == int32('p') {
			state = uint16(471)
			goto next_state
		}
		return result
	case int32(657):
		if lookahead1 == int32('p') {
			state = uint16(645)
			goto next_state
		}
		return result
	case int32(658):
		if lookahead1 == int32('p') {
			state = uint16(657)
			goto next_state
		}
		return result
	case int32(659):
		if lookahead1 == int32('q') {
			state = uint16(767)
			goto next_state
		}
		if lookahead1 == int32('s') {
			state = uint16(427)
			goto next_state
		}
		return result
	case int32(660):
		if lookahead1 == int32('q') {
			state = uint16(773)
			goto next_state
		}
		return result
	case int32(661):
		if lookahead1 == int32('q') {
			state = uint16(770)
			goto next_state
		}
		return result
	case int32(662):
		if lookahead1 == int32('q') {
			state = uint16(772)
			goto next_state
		}
		return result
	case int32(663):
		if lookahead1 == int32('r') {
			state = uint16(923)
			goto next_state
		}
		return result
	case int32(664):
		if lookahead1 == int32('r') {
			state = uint16(362)
			goto next_state
		}
		if lookahead1 == int32('u') {
			state = uint16(572)
			goto next_state
		}
		return result
	case int32(665):
		if lookahead1 == int32('r') {
			state = uint16(539)
			goto next_state
		}
		return result
	case int32(666):
		if lookahead1 == int32('r') {
			state = uint16(508)
			goto next_state
		}
		return result
	case int32(667):
		if lookahead1 == int32('r') {
			state = uint16(563)
			goto next_state
		}
		return result
	case int32(668):
		if lookahead1 == int32('r') {
			state = uint16(452)
			goto next_state
		}
		return result
	case int32(669):
		if lookahead1 == int32('r') {
			state = uint16(746)
			goto next_state
		}
		if lookahead1 == int32('y') {
			state = uint16(542)
			goto next_state
		}
		return result
	case int32(670):
		if lookahead1 == int32('r') {
			state = uint16(709)
			goto next_state
		}
		return result
	case int32(671):
		if lookahead1 == int32('r') {
			state = uint16(634)
			goto next_state
		}
		return result
	case int32(672):
		if lookahead1 == int32('r') {
			state = uint16(486)
			goto next_state
		}
		return result
	case int32(673):
		if lookahead1 == int32('r') {
			state = uint16(592)
			goto next_state
		}
		return result
	case int32(674):
		if lookahead1 == int32('r') {
			state = uint16(540)
			goto next_state
		}
		return result
	case int32(675):
		if lookahead1 == int32('r') {
			state = uint16(415)
			goto next_state
		}
		return result
	case int32(676):
		if lookahead1 == int32('r') {
			state = uint16(604)
			goto next_state
		}
		return result
	case int32(677):
		if lookahead1 == int32('r') {
			state = uint16(501)
			goto next_state
		}
		return result
	case int32(678):
		if lookahead1 == int32('r') {
			state = uint16(618)
			goto next_state
		}
		return result
	case int32(679):
		if lookahead1 == int32('r') {
			state = uint16(332)
			goto next_state
		}
		return result
	case int32(680):
		if lookahead1 == int32('r') {
			state = uint16(424)
			goto next_state
		}
		return result
	case int32(681):
		if lookahead1 == int32('r') {
			state = uint16(692)
			goto next_state
		}
		return result
	case int32(682):
		if lookahead1 == int32('r') {
			state = uint16(188)
			goto next_state
		}
		return result
	case int32(683):
		if lookahead1 == int32('r') {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(684):
		if lookahead1 == int32('r') {
			state = uint16(195)
			goto next_state
		}
		return result
	case int32(685):
		if lookahead1 == int32('r') {
			state = uint16(725)
			goto next_state
		}
		return result
	case int32(686):
		if lookahead1 == int32('r') {
			state = uint16(777)
			goto next_state
		}
		return result
	case int32(687):
		if lookahead1 == int32('r') {
			state = uint16(568)
			goto next_state
		}
		return result
	case int32(688):
		if lookahead1 == int32('r') {
			state = uint16(458)
			goto next_state
		}
		return result
	case int32(689):
		if lookahead1 == int32('r') {
			state = uint16(331)
			goto next_state
		}
		return result
	case int32(690):
		if lookahead1 == int32('r') {
			state = uint16(334)
			goto next_state
		}
		return result
	case int32(691):
		if lookahead1 == int32('r') {
			state = uint16(536)
			goto next_state
		}
		return result
	case int32(692):
		if lookahead1 == int32('r') {
			state = uint16(615)
			goto next_state
		}
		return result
	case int32(693):
		if lookahead1 == int32('r') {
			state = uint16(792)
			goto next_state
		}
		return result
	case int32(694):
		if lookahead1 == int32('r') {
			state = uint16(567)
			goto next_state
		}
		return result
	case int32(695):
		if lookahead1 == int32('r') {
			state = uint16(758)
			goto next_state
		}
		return result
	case int32(696):
		if lookahead1 == int32('s') {
			state = uint16(923)
			goto next_state
		}
		return result
	case int32(697):
		if lookahead1 == int32('s') {
			state = uint16(761)
			goto next_state
		}
		if lookahead1 == int32('t') {
			state = uint16(423)
			goto next_state
		}
		return result
	case int32(698):
		if lookahead1 == int32('s') {
			state = uint16(387)
			goto next_state
		}
		return result
	case int32(699):
		if lookahead1 == int32('s') {
			state = uint16(704)
			goto next_state
		}
		return result
	case int32(700):
		if lookahead1 == int32('s') {
			state = uint16(711)
			goto next_state
		}
		return result
	case int32(701):
		if lookahead1 == int32('s') {
			state = uint16(415)
			goto next_state
		}
		return result
	case int32(702):
		if lookahead1 == int32('s') {
			state = uint16(712)
			goto next_state
		}
		return result
	case int32(703):
		if lookahead1 == int32('s') {
			state = uint16(723)
			goto next_state
		}
		return result
	case int32(704):
		if lookahead1 == int32('s') {
			state = uint16(514)
			goto next_state
		}
		return result
	case int32(705):
		if lookahead1 == int32('s') {
			state = uint16(483)
			goto next_state
		}
		return result
	case int32(706):
		if lookahead1 == int32('s') {
			state = uint16(733)
			goto next_state
		}
		return result
	case int32(707):
		if lookahead1 == int32('s') {
			state = uint16(627)
			goto next_state
		}
		return result
	case int32(708):
		if lookahead1 == int32('s') {
			state = uint16(625)
			goto next_state
		}
		return result
	case int32(709):
		if lookahead1 == int32('s') {
			state = uint16(530)
			goto next_state
		}
		return result
	case int32(710):
		if lookahead1 == int32('s') {
			state = uint16(219)
			goto next_state
		}
		return result
	case int32(711):
		if lookahead1 == int32('s') {
			state = uint16(359)
			goto next_state
		}
		return result
	case int32(712):
		if lookahead1 == int32('t') {
			state = uint16(923)
			goto next_state
		}
		return result
	case int32(713):
		if lookahead1 == int32('t') {
			state = uint16(493)
			goto next_state
		}
		return result
	case int32(714):
		if lookahead1 == int32('t') {
			state = uint16(780)
			goto next_state
		}
		return result
	case int32(715):
		if lookahead1 == int32('t') {
			state = uint16(496)
			goto next_state
		}
		return result
	case int32(716):
		if lookahead1 == int32('t') {
			state = uint16(787)
			goto next_state
		}
		return result
	case int32(717):
		if lookahead1 == int32('t') {
			state = uint16(370)
			goto next_state
		}
		return result
	case int32(718):
		if lookahead1 == int32('t') {
			state = uint16(497)
			goto next_state
		}
		return result
	case int32(719):
		if lookahead1 == int32('t') {
			state = uint16(412)
			goto next_state
		}
		return result
	case int32(720):
		if lookahead1 == int32('t') {
			state = uint16(500)
			goto next_state
		}
		return result
	case int32(721):
		if lookahead1 == int32('t') {
			state = uint16(540)
			goto next_state
		}
		return result
	case int32(722):
		if lookahead1 == int32('t') {
			state = uint16(765)
			goto next_state
		}
		return result
	case int32(723):
		if lookahead1 == int32('t') {
			state = uint16(696)
			goto next_state
		}
		return result
	case int32(724):
		if lookahead1 == int32('t') {
			state = uint16(345)
			goto next_state
		}
		return result
	case int32(725):
		if lookahead1 == int32('t') {
			state = uint16(415)
			goto next_state
		}
		return result
	case int32(726):
		if lookahead1 == int32('t') {
			state = uint16(526)
			goto next_state
		}
		return result
	case int32(727):
		if lookahead1 == int32('t') {
			state = uint16(338)
			goto next_state
		}
		return result
	case int32(728):
		if lookahead1 == int32('t') {
			state = uint16(327)
			goto next_state
		}
		return result
	case int32(729):
		if lookahead1 == int32('t') {
			state = uint16(200)
			goto next_state
		}
		return result
	case int32(730):
		if lookahead1 == int32('t') {
			state = uint16(176)
			goto next_state
		}
		return result
	case int32(731):
		if lookahead1 == int32('t') {
			state = uint16(509)
			goto next_state
		}
		return result
	case int32(732):
		if lookahead1 == int32('t') {
			state = uint16(511)
			goto next_state
		}
		return result
	case int32(733):
		if lookahead1 == int32('t') {
			state = uint16(177)
			goto next_state
		}
		return result
	case int32(734):
		if lookahead1 == int32('t') {
			state = uint16(629)
			goto next_state
		}
		return result
	case int32(735):
		if lookahead1 == int32('t') {
			state = uint16(525)
			goto next_state
		}
		return result
	case int32(736):
		if lookahead1 == int32('t') {
			state = uint16(196)
			goto next_state
		}
		return result
	case int32(737):
		if lookahead1 == int32('t') {
			state = uint16(534)
			goto next_state
		}
		return result
	case int32(738):
		if lookahead1 == int32('t') {
			state = uint16(437)
			goto next_state
		}
		return result
	case int32(739):
		if lookahead1 == int32('t') {
			state = uint16(440)
			goto next_state
		}
		return result
	case int32(740):
		if lookahead1 == int32('t') {
			state = uint16(202)
			goto next_state
		}
		return result
	case int32(741):
		if lookahead1 == int32('t') {
			state = uint16(213)
			goto next_state
		}
		return result
	case int32(742):
		if lookahead1 == int32('t') {
			state = uint16(431)
			goto next_state
		}
		return result
	case int32(743):
		if lookahead1 == int32('t') {
			state = uint16(206)
			goto next_state
		}
		return result
	case int32(744):
		if lookahead1 == int32('t') {
			state = uint16(433)
			goto next_state
		}
		return result
	case int32(745):
		if lookahead1 == int32('t') {
			state = uint16(494)
			goto next_state
		}
		return result
	case int32(746):
		if lookahead1 == int32('t') {
			state = uint16(510)
			goto next_state
		}
		return result
	case int32(747):
		if lookahead1 == int32('t') {
			state = uint16(498)
			goto next_state
		}
		return result
	case int32(748):
		if lookahead1 == int32('t') {
			state = uint16(457)
			goto next_state
		}
		return result
	case int32(749):
		if lookahead1 == int32('t') {
			state = uint16(214)
			goto next_state
		}
		return result
	case int32(750):
		if lookahead1 == int32('t') {
			state = uint16(499)
			goto next_state
		}
		return result
	case int32(751):
		if lookahead1 == int32('t') {
			state = uint16(215)
			goto next_state
		}
		return result
	case int32(752):
		if lookahead1 == int32('t') {
			state = uint16(515)
			goto next_state
		}
		return result
	case int32(753):
		if lookahead1 == int32('t') {
			state = uint16(621)
			goto next_state
		}
		return result
	case int32(754):
		if lookahead1 == int32('t') {
			state = uint16(469)
			goto next_state
		}
		return result
	case int32(755):
		if lookahead1 == int32('t') {
			state = uint16(467)
			goto next_state
		}
		return result
	case int32(756):
		if lookahead1 == int32('t') {
			state = uint16(354)
			goto next_state
		}
		return result
	case int32(757):
		if lookahead1 == int32('t') {
			state = uint16(531)
			goto next_state
		}
		return result
	case int32(758):
		if lookahead1 == int32('t') {
			state = uint16(468)
			goto next_state
		}
		return result
	case int32(759):
		if lookahead1 == int32('t') {
			state = uint16(532)
			goto next_state
		}
		return result
	case int32(760):
		if lookahead1 == int32('t') {
			state = uint16(533)
			goto next_state
		}
		return result
	case int32(761):
		if lookahead1 == int32('u') {
			state = uint16(478)
			goto next_state
		}
		return result
	case int32(762):
		if lookahead1 == int32('u') {
			state = uint16(572)
			goto next_state
		}
		return result
	case int32(763):
		if lookahead1 == int32('u') {
			state = uint16(655)
			goto next_state
		}
		return result
	case int32(764):
		if lookahead1 == int32('u') {
			state = uint16(406)
			goto next_state
		}
		return result
	case int32(765):
		if lookahead1 == int32('u') {
			state = uint16(696)
			goto next_state
		}
		return result
	case int32(766):
		if lookahead1 == int32('u') {
			state = uint16(712)
			goto next_state
		}
		return result
	case int32(767):
		if lookahead1 == int32('u') {
			state = uint16(421)
			goto next_state
		}
		return result
	case int32(768):
		if lookahead1 == int32('u') {
			state = uint16(745)
			goto next_state
		}
		if lookahead1 == int32('v') {
			state = uint16(341)
			goto next_state
		}
		return result
	case int32(769):
		if lookahead1 == int32('u') {
			state = uint16(747)
			goto next_state
		}
		return result
	case int32(770):
		if lookahead1 == int32('u') {
			state = uint16(460)
			goto next_state
		}
		return result
	case int32(771):
		if lookahead1 == int32('u') {
			state = uint16(750)
			goto next_state
		}
		return result
	case int32(772):
		if lookahead1 == int32('u') {
			state = uint16(463)
			goto next_state
		}
		return result
	case int32(773):
		if lookahead1 == int32('u') {
			state = uint16(524)
			goto next_state
		}
		return result
	case int32(774):
		if lookahead1 == int32('u') {
			state = uint16(658)
			goto next_state
		}
		return result
	case int32(775):
		if lookahead1 == int32('v') {
			state = uint16(517)
			goto next_state
		}
		return result
	case int32(776):
		if lookahead1 == int32('v') {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(777):
		if lookahead1 == int32('v') {
			state = uint16(461)
			goto next_state
		}
		return result
	case int32(778):
		if lookahead1 == int32('v') {
			state = uint16(451)
			goto next_state
		}
		return result
	case int32(779):
		if lookahead1 == int32('v') {
			state = uint16(347)
			goto next_state
		}
		return result
	case int32(780):
		if lookahead1 == int32('w') {
			state = uint16(608)
			goto next_state
		}
		return result
	case int32(781):
		if lookahead1 == int32('w') {
			state = uint16(348)
			goto next_state
		}
		return result
	case int32(782):
		if lookahead1 == int32('w') {
			state = uint16(415)
			goto next_state
		}
		return result
	case int32(783):
		if lookahead1 == int32('w') {
			state = uint16(324)
			goto next_state
		}
		return result
	case int32(784):
		if lookahead1 == int32('x') {
			state = uint16(646)
			goto next_state
		}
		return result
	case int32(785):
		if lookahead1 == int32('x') {
			state = uint16(787)
			goto next_state
		}
		return result
	case int32(786):
		if lookahead1 == int32('x') {
			state = uint16(739)
			goto next_state
		}
		return result
	case int32(787):
		if lookahead1 == int32('y') {
			state = uint16(923)
			goto next_state
		}
		return result
	case int32(788):
		if lookahead1 == int32('y') {
			state = uint16(212)
			goto next_state
		}
		return result
	case int32(789):
		if lookahead1 == int32('y') {
			state = uint16(207)
			goto next_state
		}
		return result
	case int32(790):
		if lookahead1 == int32('y') {
			state = uint16(210)
			goto next_state
		}
		return result
	case int32(791):
		if lookahead1 == int32('y') {
			state = uint16(647)
			goto next_state
		}
		return result
	case int32(792):
		if lookahead1 == int32('y') {
			state = uint16(215)
			goto next_state
		}
		return result
	case int32(793):
		if lookahead1 == int32('y') {
			state = uint16(216)
			goto next_state
		}
		return result
	case int32(794):
		if lookahead1 == int32('z') {
			state = uint16(415)
			goto next_state
		}
		return result
	case int32(795):
		if lookahead1 == int32('}') {
			state = uint16(929)
			goto next_state
		}
		return result
	case int32(796):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(922)
			goto next_state
		}
		return result
	case int32(797):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(796)
			goto next_state
		}
		return result
	case int32(798):
		if lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(921)
			goto next_state
		}
		return result
	case int32(799):
		if eof != 0 {
			state = uint16(806)
			goto next_state
		}
		i2 = uint32(0)
		for {
			if !(uint64(i2) < libc.Uint64FromInt64(96)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token2[i2]) == lookahead1 {
				state = map_token2[i2+uint32(1)]
				goto next_state
			}
			goto _37
		_37:
			;
			i2 = i2 + uint32(2)
		}
		if lookahead1 == int32(' ') || lookahead1 == int32(0xa0) || lookahead1 == int32(0x1680) || int32(0x2000) <= lookahead1 && lookahead1 <= int32(0x200a) || lookahead1 == int32(0x202f) || lookahead1 == int32(0x205f) || lookahead1 == int32(0x3000) {
			state = uint16(893)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&aux_sym_WORD_CHAR_token1_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(475) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _41
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _41
	_41:
		if v4 != 0 {
			state = uint16(831)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(875)
			goto next_state
		}
		return result
	case int32(800):
		if eof != 0 {
			state = uint16(806)
			goto next_state
		}
		i3 = uint32(0)
		for {
			if !(uint64(i3) < libc.Uint64FromInt64(76)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token3[i3]) == lookahead1 {
				state = map_token3[i3+uint32(1)]
				goto next_state
			}
			goto _42
		_42:
			;
			i3 = i3 + uint32(2)
		}
		if lookahead1 == int32(' ') || lookahead1 == int32(0xa0) || lookahead1 == int32(0x1680) || int32(0x2000) <= lookahead1 && lookahead1 <= int32(0x200a) || lookahead1 == int32(0x202f) || lookahead1 == int32(0x205f) || lookahead1 == int32(0x3000) {
			state = uint16(896)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(862)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&aux_sym_WORD_CHAR_token1_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(475) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _46
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _46
	_46:
		if v4 != 0 {
			state = uint16(807)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32(0x2028) && lookahead1 != int32(0x2029) {
			state = uint16(863)
			goto next_state
		}
		return result
	case int32(801):
		if eof != 0 {
			state = uint16(806)
			goto next_state
		}
		i4 = uint32(0)
		for {
			if !(uint64(i4) < libc.Uint64FromInt64(68)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token4[i4]) == lookahead1 {
				state = map_token4[i4+uint32(1)]
				goto next_state
			}
			goto _47
		_47:
			;
			i4 = i4 + uint32(2)
		}
		if lookahead1 == int32(' ') || lookahead1 == int32(0xa0) || lookahead1 == int32(0x1680) || int32(0x2000) <= lookahead1 && lookahead1 <= int32(0x200a) || lookahead1 == int32(0x202f) || lookahead1 == int32(0x205f) || lookahead1 == int32(0x3000) {
			state = uint16(896)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&aux_sym_WORD_CHAR_token1_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(475) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _51
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _51
	_51:
		if v4 != 0 {
			state = uint16(807)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32(0x2028) && lookahead1 != int32(0x2029) {
			state = uint16(863)
			goto next_state
		}
		return result
	case int32(802):
		if eof != 0 {
			state = uint16(806)
			goto next_state
		}
		i5 = uint32(0)
		for {
			if !(uint64(i5) < libc.Uint64FromInt64(96)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token5[i5]) == lookahead1 {
				state = map_token5[i5+uint32(1)]
				goto next_state
			}
			goto _52
		_52:
			;
			i5 = i5 + uint32(2)
		}
		if lookahead1 == int32(' ') || lookahead1 == int32(0xa0) || lookahead1 == int32(0x1680) || int32(0x2000) <= lookahead1 && lookahead1 <= int32(0x200a) || lookahead1 == int32(0x202f) || lookahead1 == int32(0x205f) || lookahead1 == int32(0x3000) {
			state = uint16(892)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&aux_sym_WORD_CHAR_token1_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(475) - index
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
			state = uint16(819)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(868)
			goto next_state
		}
		return result
	case int32(803):
		if eof != 0 {
			state = uint16(806)
			goto next_state
		}
		i6 = uint32(0)
		for {
			if !(uint64(i6) < libc.Uint64FromInt64(96)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token6[i6]) == lookahead1 {
				state = map_token6[i6+uint32(1)]
				goto next_state
			}
			goto _57
		_57:
			;
			i6 = i6 + uint32(2)
		}
		if lookahead1 == int32(' ') || lookahead1 == int32(0xa0) || lookahead1 == int32(0x1680) || int32(0x2000) <= lookahead1 && lookahead1 <= int32(0x200a) || lookahead1 == int32(0x202f) || lookahead1 == int32(0x205f) || lookahead1 == int32(0x3000) {
			state = uint16(892)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&aux_sym_WORD_CHAR_token1_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(475) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _61
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _61
	_61:
		if v4 != 0 {
			state = uint16(819)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(868)
			goto next_state
		}
		return result
	case int32(804):
		if eof != 0 {
			state = uint16(806)
			goto next_state
		}
		i7 = uint32(0)
		for {
			if !(uint64(i7) < libc.Uint64FromInt64(96)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token7[i7]) == lookahead1 {
				state = map_token7[i7+uint32(1)]
				goto next_state
			}
			goto _62
		_62:
			;
			i7 = i7 + uint32(2)
		}
		if lookahead1 == int32(' ') || lookahead1 == int32(0xa0) || lookahead1 == int32(0x1680) || int32(0x2000) <= lookahead1 && lookahead1 <= int32(0x200a) || lookahead1 == int32(0x202f) || lookahead1 == int32(0x205f) || lookahead1 == int32(0x3000) {
			state = uint16(894)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&aux_sym_WORD_CHAR_token1_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(475) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _66
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _66
	_66:
		if v4 != 0 {
			state = uint16(843)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(880)
			goto next_state
		}
		return result
	case int32(805):
		if eof != 0 {
			state = uint16(806)
			goto next_state
		}
		i8 = uint32(0)
		for {
			if !(uint64(i8) < libc.Uint64FromInt64(96)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token8[i8]) == lookahead1 {
				state = map_token8[i8+uint32(1)]
				goto next_state
			}
			goto _67
		_67:
			;
			i8 = i8 + uint32(2)
		}
		if lookahead1 == int32(' ') || lookahead1 == int32(0xa0) || lookahead1 == int32(0x1680) || int32(0x2000) <= lookahead1 && lookahead1 <= int32(0x200a) || lookahead1 == int32(0x202f) || lookahead1 == int32(0x205f) || lookahead1 == int32(0x3000) {
			state = uint16(894)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&aux_sym_WORD_CHAR_token1_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(475) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _71
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _71
	_71:
		if v4 != 0 {
			state = uint16(843)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(880)
			goto next_state
		}
		return result
	case int32(806):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(ts_builtin_sym_end)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(807):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WORD_CHAR_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(808):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WORD_CHAR_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('A') {
			state = uint16(38)
			goto next_state
		}
		if lookahead1 == int32('O') {
			state = uint16(33)
			goto next_state
		}
		if lookahead1 == int32('U') {
			state = uint16(36)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(809):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WORD_CHAR_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('E') {
			state = uint16(23)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(810):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WORD_CHAR_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('E') {
			state = uint16(36)
			goto next_state
		}
		if lookahead1 == int32('R') {
			state = uint16(6)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(811):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WORD_CHAR_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('E') {
			state = uint16(9)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(812):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WORD_CHAR_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('E') {
			state = uint16(7)
			goto next_state
		}
		if lookahead1 == int32('T') {
			state = uint16(39)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(813):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WORD_CHAR_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('I') {
			state = uint16(33)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(814):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WORD_CHAR_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('O') {
			state = uint16(25)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(815):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WORD_CHAR_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('P') {
			state = uint16(37)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(816):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WORD_CHAR_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('R') {
			state = uint16(8)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(817):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WORD_CHAR_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('u') {
			state = uint16(47)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(818):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WORD_CHAR_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('u') {
			state = uint16(42)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(819):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WORD_CHAR_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(820):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WORD_CHAR_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('A') {
			state = uint16(91)
			goto next_state
		}
		if lookahead1 == int32('O') {
			state = uint16(87)
			goto next_state
		}
		if lookahead1 == int32('U') {
			state = uint16(88)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(821):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WORD_CHAR_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('E') {
			state = uint16(61)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(822):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WORD_CHAR_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('E') {
			state = uint16(76)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(823):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WORD_CHAR_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('E') {
			state = uint16(58)
			goto next_state
		}
		if lookahead1 == int32('T') {
			state = uint16(90)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(824):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WORD_CHAR_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('E') {
			state = uint16(88)
			goto next_state
		}
		if lookahead1 == int32('R') {
			state = uint16(59)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(825):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WORD_CHAR_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('I') {
			state = uint16(87)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(826):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WORD_CHAR_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('O') {
			state = uint16(77)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(827):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WORD_CHAR_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('P') {
			state = uint16(89)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(828):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WORD_CHAR_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('R') {
			state = uint16(60)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(829):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WORD_CHAR_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('u') {
			state = uint16(94)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(830):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WORD_CHAR_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('u') {
			state = uint16(99)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(831):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WORD_CHAR_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(832):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WORD_CHAR_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('A') {
			state = uint16(144)
			goto next_state
		}
		if lookahead1 == int32('O') {
			state = uint16(140)
			goto next_state
		}
		if lookahead1 == int32('U') {
			state = uint16(141)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(833):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WORD_CHAR_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('E') {
			state = uint16(114)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(834):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WORD_CHAR_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('E') {
			state = uint16(129)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(835):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WORD_CHAR_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('E') {
			state = uint16(111)
			goto next_state
		}
		if lookahead1 == int32('T') {
			state = uint16(143)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(836):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WORD_CHAR_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('E') {
			state = uint16(141)
			goto next_state
		}
		if lookahead1 == int32('R') {
			state = uint16(112)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(837):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WORD_CHAR_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('I') {
			state = uint16(140)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(838):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WORD_CHAR_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('O') {
			state = uint16(130)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(839):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WORD_CHAR_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('P') {
			state = uint16(142)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(840):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WORD_CHAR_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('R') {
			state = uint16(113)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(841):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WORD_CHAR_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('u') {
			state = uint16(147)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(842):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WORD_CHAR_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('u') {
			state = uint16(152)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(843):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WORD_CHAR_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(844):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WORD_CHAR_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('A') {
			state = uint16(1004)
			goto next_state
		}
		if lookahead1 == int32('O') {
			state = uint16(999)
			goto next_state
		}
		if lookahead1 == int32('U') {
			state = uint16(1002)
			goto next_state
		}
		if lookahead1 == int32('-') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('B') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(1007)
			goto next_state
		}
		return result
	case int32(845):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WORD_CHAR_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('A') {
			state = uint16(305)
			goto next_state
		}
		if lookahead1 == int32('O') {
			state = uint16(295)
			goto next_state
		}
		if lookahead1 == int32('U') {
			state = uint16(303)
			goto next_state
		}
		return result
	case int32(846):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WORD_CHAR_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('E') {
			state = uint16(989)
			goto next_state
		}
		if lookahead1 == int32('-') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(1007)
			goto next_state
		}
		return result
	case int32(847):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WORD_CHAR_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('E') {
			state = uint16(1002)
			goto next_state
		}
		if lookahead1 == int32('R') {
			state = uint16(972)
			goto next_state
		}
		if lookahead1 == int32('-') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(1007)
			goto next_state
		}
		return result
	case int32(848):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WORD_CHAR_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('E') {
			state = uint16(975)
			goto next_state
		}
		if lookahead1 == int32('-') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(1007)
			goto next_state
		}
		return result
	case int32(849):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WORD_CHAR_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('E') {
			state = uint16(265)
			goto next_state
		}
		return result
	case int32(850):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WORD_CHAR_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('E') {
			state = uint16(303)
			goto next_state
		}
		if lookahead1 == int32('R') {
			state = uint16(227)
			goto next_state
		}
		return result
	case int32(851):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WORD_CHAR_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('E') {
			state = uint16(235)
			goto next_state
		}
		return result
	case int32(852):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WORD_CHAR_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('E') {
			state = uint16(973)
			goto next_state
		}
		if lookahead1 == int32('T') {
			state = uint16(1005)
			goto next_state
		}
		if lookahead1 == int32('-') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(1007)
			goto next_state
		}
		return result
	case int32(853):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WORD_CHAR_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('E') {
			state = uint16(228)
			goto next_state
		}
		if lookahead1 == int32('T') {
			state = uint16(306)
			goto next_state
		}
		return result
	case int32(854):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WORD_CHAR_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('I') {
			state = uint16(999)
			goto next_state
		}
		if lookahead1 == int32('-') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(1007)
			goto next_state
		}
		return result
	case int32(855):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WORD_CHAR_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('I') {
			state = uint16(295)
			goto next_state
		}
		return result
	case int32(856):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WORD_CHAR_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('O') {
			state = uint16(991)
			goto next_state
		}
		if lookahead1 == int32('-') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(1007)
			goto next_state
		}
		return result
	case int32(857):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WORD_CHAR_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('O') {
			state = uint16(271)
			goto next_state
		}
		return result
	case int32(858):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WORD_CHAR_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('P') {
			state = uint16(1003)
			goto next_state
		}
		if lookahead1 == int32('-') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(1007)
			goto next_state
		}
		return result
	case int32(859):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WORD_CHAR_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('P') {
			state = uint16(304)
			goto next_state
		}
		return result
	case int32(860):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WORD_CHAR_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('R') {
			state = uint16(974)
			goto next_state
		}
		if lookahead1 == int32('-') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(1007)
			goto next_state
		}
		return result
	case int32(861):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WORD_CHAR_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('R') {
			state = uint16(229)
			goto next_state
		}
		return result
	case int32(862):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WORD_CHAR_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('-') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(1007)
			goto next_state
		}
		return result
	case int32(863):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_PUNCTUATION_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(864):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_PUNCTUATION_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(952)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(951)
			goto next_state
		}
		if lookahead1 == int32('{') {
			state = uint16(926)
			goto next_state
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\f') || lookahead1 == int32(' ') {
			state = uint16(951)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(865):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_PUNCTUATION_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(952)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(951)
			goto next_state
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\f') || lookahead1 == int32(' ') {
			state = uint16(951)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(866):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_PUNCTUATION_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('-') {
			state = uint16(956)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(867):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_PUNCTUATION_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(905)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(868):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_PUNCTUATION_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(869):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_PUNCTUATION_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(945)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(944)
			goto next_state
		}
		if lookahead1 == int32('{') {
			state = uint16(926)
			goto next_state
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\f') || lookahead1 == int32(' ') {
			state = uint16(944)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(870):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_PUNCTUATION_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(945)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(944)
			goto next_state
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\f') || lookahead1 == int32(' ') {
			state = uint16(944)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(871):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_PUNCTUATION_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(945)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(944)
			goto next_state
		}
		if lookahead1 == int32('{') {
			state = uint16(927)
			goto next_state
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\f') || lookahead1 == int32(' ') {
			state = uint16(944)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(872):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_PUNCTUATION_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(945)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(944)
			goto next_state
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\f') || lookahead1 == int32(' ') {
			state = uint16(944)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(873):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_PUNCTUATION_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('-') {
			state = uint16(956)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(874):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_PUNCTUATION_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(905)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(875):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_PUNCTUATION_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(876):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_PUNCTUATION_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(952)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(951)
			goto next_state
		}
		if lookahead1 == int32('{') {
			state = uint16(928)
			goto next_state
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\f') || lookahead1 == int32(' ') {
			state = uint16(951)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(877):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_PUNCTUATION_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(952)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(951)
			goto next_state
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\f') || lookahead1 == int32(' ') {
			state = uint16(951)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(878):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_PUNCTUATION_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('-') {
			state = uint16(957)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(879):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_PUNCTUATION_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(905)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(880):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_PUNCTUATION_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(881):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_PUNCTUATION_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(968)
			goto next_state
		}
		if lookahead1 == int32('{') {
			state = uint16(928)
			goto next_state
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\f') || lookahead1 == int32(' ') {
			state = uint16(943)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(882):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_PUNCTUATION_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(968)
			goto next_state
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\f') || lookahead1 == int32(' ') {
			state = uint16(943)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(883):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_PUNCTUATION_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\r') {
			state = uint16(931)
			goto next_state
		}
		if !(eof != 0) && lookahead1 == 00 || lookahead1 == int32('\n') {
			state = uint16(931)
			goto next_state
		}
		return result
	case int32(884):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_PUNCTUATION_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('%') {
			state = uint16(933)
			goto next_state
		}
		if lookahead1 == int32('{') {
			state = uint16(925)
			goto next_state
		}
		return result
	case int32(885):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_PUNCTUATION_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('%') {
			state = uint16(933)
			goto next_state
		}
		if lookahead1 == int32('{') {
			state = uint16(925)
			goto next_state
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			state = uint16(952)
			goto next_state
		}
		return result
	case int32(886):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_PUNCTUATION_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('-') {
			state = uint16(956)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(1007)
			goto next_state
		}
		return result
	case int32(887):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_PUNCTUATION_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('/') {
			state = uint16(907)
			goto next_state
		}
		return result
	case int32(888):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_PUNCTUATION_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('{') {
			state = uint16(925)
			goto next_state
		}
		return result
	case int32(889):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_PUNCTUATION_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			state = uint16(952)
			goto next_state
		}
		return result
	case int32(890):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_PUNCTUATION_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('-') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(1007)
			goto next_state
		}
		return result
	case int32(891):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_PUNCTUATION_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 != 0 && lookahead1 != int32('\n') && lookahead1 != int32('\r') {
			state = uint16(1009)
			goto next_state
		}
		return result
	case int32(892):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WS_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32(' ') || lookahead1 == int32(0xa0) || lookahead1 == int32(0x1680) || int32(0x2000) <= lookahead1 && lookahead1 <= int32(0x200a) || lookahead1 == int32(0x202f) || lookahead1 == int32(0x205f) || lookahead1 == int32(0x3000) {
			state = uint16(892)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(893):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WS_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32(' ') || lookahead1 == int32(0xa0) || lookahead1 == int32(0x1680) || int32(0x2000) <= lookahead1 && lookahead1 <= int32(0x200a) || lookahead1 == int32(0x202f) || lookahead1 == int32(0x205f) || lookahead1 == int32(0x3000) {
			state = uint16(893)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(894):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WS_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32(' ') || lookahead1 == int32(0xa0) || lookahead1 == int32(0x1680) || int32(0x2000) <= lookahead1 && lookahead1 <= int32(0x200a) || lookahead1 == int32(0x202f) || lookahead1 == int32(0x205f) || lookahead1 == int32(0x3000) {
			state = uint16(894)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(895):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WS_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(' ') || lookahead1 == int32(0xa0) {
			state = uint16(896)
			goto next_state
		}
		if lookahead1 == int32(0x1680) || int32(0x2000) <= lookahead1 && lookahead1 <= int32(0x200a) || lookahead1 == int32(0x202f) || lookahead1 == int32(0x205f) || lookahead1 == int32(0x3000) {
			state = uint16(895)
			goto next_state
		}
		if lookahead1 == int32('$') || lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') || int32(0xa1) <= lookahead1 && lookahead1 <= int32(0xffff) {
			state = uint16(1008)
			goto next_state
		}
		return result
	case int32(896):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_WS_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(' ') || lookahead1 == int32(0xa0) || lookahead1 == int32(0x1680) || int32(0x2000) <= lookahead1 && lookahead1 <= int32(0x200a) || lookahead1 == int32(0x202f) || lookahead1 == int32(0x205f) || lookahead1 == int32(0x3000) {
			state = uint16(896)
			goto next_state
		}
		return result
	case int32(897):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_NL_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(898):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_NL_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(899):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_NL_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\n') {
			state = uint16(897)
			goto next_state
		}
		return result
	case int32(900):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_LINE_TAIL_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(901):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_LINE_TAIL_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(901)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(900)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(901)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(164)
			goto next_state
		}
		return result
	case int32(902):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_COMMENT_PREFIX_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\n') {
			state = uint16(907)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(905)
			goto next_state
		}
		if lookahead1 == int32('#') {
			state = uint16(4)
			goto next_state
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\f') || lookahead1 == int32(' ') {
			state = uint16(905)
			goto next_state
		}
		return result
	case int32(903):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_COMMENT_PREFIX_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\n') {
			state = uint16(907)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(905)
			goto next_state
		}
		if lookahead1 == int32('#') {
			state = uint16(56)
			goto next_state
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\f') || lookahead1 == int32(' ') {
			state = uint16(905)
			goto next_state
		}
		return result
	case int32(904):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_COMMENT_PREFIX_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\n') {
			state = uint16(907)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(905)
			goto next_state
		}
		if lookahead1 == int32('#') {
			state = uint16(107)
			goto next_state
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\f') || lookahead1 == int32(' ') {
			state = uint16(905)
			goto next_state
		}
		return result
	case int32(905):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_COMMENT_PREFIX_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\n') {
			state = uint16(907)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(905)
			goto next_state
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\f') || lookahead1 == int32(' ') {
			state = uint16(905)
			goto next_state
		}
		return result
	case int32(906):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_COMMENT_PREFIX_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(222)
			goto next_state
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			state = uint16(907)
			goto next_state
		}
		return result
	case int32(907):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_COMMENT_PREFIX_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			state = uint16(907)
			goto next_state
		}
		return result
	case int32(908):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(909):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(910):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__var_comment_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(' ') || lookahead1 == int32(0xa0) || lookahead1 == int32(0x1680) || int32(0x2000) <= lookahead1 && lookahead1 <= int32(0x200a) || lookahead1 == int32(0x202f) || lookahead1 == int32(0x205f) || lookahead1 == int32(0x3000) {
			state = uint16(910)
			goto next_state
		}
		return result
	case int32(911):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_request_separator_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(911)
			goto next_state
		}
		if lookahead1 == int32(' ') || lookahead1 == int32(0xa0) || lookahead1 == int32(0x1680) || int32(0x2000) <= lookahead1 && lookahead1 <= int32(0x200a) || lookahead1 == int32(0x202f) || lookahead1 == int32(0x205f) || lookahead1 == int32(0x3000) {
			state = uint16(912)
			goto next_state
		}
		return result
	case int32(912):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_request_separator_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(' ') || lookahead1 == int32(0xa0) || lookahead1 == int32(0x1680) || int32(0x2000) <= lookahead1 && lookahead1 <= int32(0x200a) || lookahead1 == int32(0x202f) || lookahead1 == int32(0x205f) || lookahead1 == int32(0x3000) {
			state = uint16(912)
			goto next_state
		}
		return result
	case int32(913):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_method)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(914):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_method)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(915):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_method)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(916):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_method)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(917):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_method)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('-') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(1007)
			goto next_state
		}
		return result
	case int32(918):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_http_version_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(918)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(919):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_http_version_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(919)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(920):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_http_version_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(920)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(921):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_http_version_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(921)
			goto next_state
		}
		return result
	case int32(922):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_status_code)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(923):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_status_text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(924):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COLON)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(925):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACE_LBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(926):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACE_LBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(927):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACE_LBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(928):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACE_LBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(929):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACE_RBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(930):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(931):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_pre_request_script_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\r') {
			state = uint16(931)
			goto next_state
		}
		if !(eof != 0) && lookahead1 == 00 || lookahead1 == int32('\n') {
			state = uint16(931)
			goto next_state
		}
		return result
	case int32(932):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_GT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('>') {
			state = uint16(936)
			goto next_state
		}
		return result
	case int32(933):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACE_PERCENT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(934):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PERCENT_RBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(935):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_res_redirect_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(936):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_res_redirect_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('!') {
			state = uint16(935)
			goto next_state
		}
		return result
	case int32(937):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AT2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(938):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AT2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(939):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AT2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(940):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AT2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(941):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_xml_body_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(942):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_xml_body_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(943):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_json_body_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(968)
			goto next_state
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\f') || lookahead1 == int32(' ') {
			state = uint16(943)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(944):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_json_body_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\n') {
			state = uint16(945)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(944)
			goto next_state
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\f') || lookahead1 == int32(' ') {
			state = uint16(944)
			goto next_state
		}
		return result
	case int32(945):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_json_body_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			state = uint16(945)
			goto next_state
		}
		return result
	case int32(946):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_graphql_data_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(947):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_graphql_data_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(965)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(946)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(965)
			goto next_state
		}
		if lookahead1 == int32('{') {
			state = uint16(55)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(54)
			goto next_state
		}
		return result
	case int32(948):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_graphql_data_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(965)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(965)
			goto next_state
		}
		if lookahead1 == int32('{') {
			state = uint16(55)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32('\n') {
			state = uint16(54)
			goto next_state
		}
		return result
	case int32(949):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_graphql_data_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(960)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(946)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(960)
			goto next_state
		}
		if lookahead1 == int32('{') {
			state = uint16(165)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(106)
			goto next_state
		}
		return result
	case int32(950):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_graphql_data_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(960)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(960)
			goto next_state
		}
		if lookahead1 == int32('{') {
			state = uint16(165)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32('\n') {
			state = uint16(106)
			goto next_state
		}
		return result
	case int32(951):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_graphql_json_body_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\n') {
			state = uint16(952)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(951)
			goto next_state
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\f') || lookahead1 == int32(' ') {
			state = uint16(951)
			goto next_state
		}
		return result
	case int32(952):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_graphql_json_body_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			state = uint16(952)
			goto next_state
		}
		return result
	case int32(953):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(941)
			goto next_state
		}
		if lookahead1 != 0 && (lookahead1 < int32('\t') || int32('\r') < lookahead1) && lookahead1 != int32(' ') && lookahead1 != int32('@') {
			state = uint16(941)
			goto next_state
		}
		return result
	case int32(954):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\f') || lookahead1 == int32(' ') || lookahead1 == int32('@') {
			state = uint16(157)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(942)
			goto next_state
		}
		return result
	case int32(955):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 != 0 && (lookahead1 < int32('\t') || int32('\r') < lookahead1) && lookahead1 != int32(' ') && lookahead1 != int32('@') {
			state = uint16(941)
			goto next_state
		}
		return result
	case int32(956):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DASH_DASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(957):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DASH_DASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(958):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_multipart_form_data_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(959):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_multipart_form_data_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(959)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(960):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_multipart_form_data_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(960)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(958)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(960)
			goto next_state
		}
		if lookahead1 == int32('{') {
			state = uint16(165)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(106)
			goto next_state
		}
		return result
	case int32(961):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_multipart_form_data_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(962):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_multipart_form_data_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\n') {
			state = uint16(961)
			goto next_state
		}
		return result
	case int32(963):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_raw_body_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(964):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_raw_body_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(964)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(965):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_raw_body_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(965)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(963)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(965)
			goto next_state
		}
		if lookahead1 == int32('{') {
			state = uint16(55)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(54)
			goto next_state
		}
		return result
	case int32(966):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__raw_body_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(967):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__raw_body_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(968):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__raw_body_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(967)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(968)
			goto next_state
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\f') || lookahead1 == int32(' ') {
			state = uint16(943)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(969):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__raw_body_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && lookahead1 == 00 {
			state = uint16(969)
			goto next_state
		}
		if lookahead1 == int32('\n') {
			state = uint16(966)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(969)
			goto next_state
		}
		if lookahead1 == int32('{') {
			state = uint16(159)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(159)
			goto next_state
		}
		return result
	case int32(970):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__not_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\r') {
			state = uint16(970)
			goto next_state
		}
		if !(eof != 0) && lookahead1 == 00 || lookahead1 == int32('\n') {
			state = uint16(970)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32('@') {
			state = uint16(170)
			goto next_state
		}
		return result
	case int32(971):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_header_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('/') {
			state = uint16(798)
			goto next_state
		}
		if lookahead1 == int32('-') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(1007)
			goto next_state
		}
		return result
	case int32(972):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_header_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('A') {
			state = uint16(996)
			goto next_state
		}
		if lookahead1 == int32('-') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('B') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(1007)
			goto next_state
		}
		return result
	case int32(973):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_header_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('A') {
			state = uint16(980)
			goto next_state
		}
		if lookahead1 == int32('-') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('B') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(1007)
			goto next_state
		}
		return result
	case int32(974):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_header_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('A') {
			state = uint16(979)
			goto next_state
		}
		if lookahead1 == int32('-') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('B') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(1007)
			goto next_state
		}
		return result
	case int32(975):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_header_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('B') {
			state = uint16(1001)
			goto next_state
		}
		if lookahead1 == int32('-') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(1007)
			goto next_state
		}
		return result
	case int32(976):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_header_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('C') {
			state = uint16(1002)
			goto next_state
		}
		if lookahead1 == int32('-') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(1007)
			goto next_state
		}
		return result
	case int32(977):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_header_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('C') {
			state = uint16(988)
			goto next_state
		}
		if lookahead1 == int32('-') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(1007)
			goto next_state
		}
		return result
	case int32(978):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_header_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('C') {
			state = uint16(985)
			goto next_state
		}
		if lookahead1 == int32('-') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(1007)
			goto next_state
		}
		return result
	case int32(979):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_header_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('C') {
			state = uint16(982)
			goto next_state
		}
		if lookahead1 == int32('-') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(1007)
			goto next_state
		}
		return result
	case int32(980):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_header_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('D') {
			state = uint16(917)
			goto next_state
		}
		if lookahead1 == int32('-') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(1007)
			goto next_state
		}
		return result
	case int32(981):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_header_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('E') {
			state = uint16(1002)
			goto next_state
		}
		if lookahead1 == int32('-') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(1007)
			goto next_state
		}
		return result
	case int32(982):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_header_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('E') {
			state = uint16(917)
			goto next_state
		}
		if lookahead1 == int32('-') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(1007)
			goto next_state
		}
		return result
	case int32(983):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_header_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('E') {
			state = uint16(976)
			goto next_state
		}
		if lookahead1 == int32('-') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(1007)
			goto next_state
		}
		return result
	case int32(984):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_header_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('E') {
			state = uint16(1006)
			goto next_state
		}
		if lookahead1 == int32('-') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(1007)
			goto next_state
		}
		return result
	case int32(985):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_header_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('H') {
			state = uint16(917)
			goto next_state
		}
		if lookahead1 == int32('-') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(1007)
			goto next_state
		}
		return result
	case int32(986):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_header_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('H') {
			state = uint16(998)
			goto next_state
		}
		if lookahead1 == int32('-') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(1007)
			goto next_state
		}
		return result
	case int32(987):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_header_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('I') {
			state = uint16(994)
			goto next_state
		}
		if lookahead1 == int32('-') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(1007)
			goto next_state
		}
		return result
	case int32(988):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_header_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('K') {
			state = uint16(981)
			goto next_state
		}
		if lookahead1 == int32('-') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(1007)
			goto next_state
		}
		return result
	case int32(989):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_header_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('L') {
			state = uint16(984)
			goto next_state
		}
		if lookahead1 == int32('-') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(1007)
			goto next_state
		}
		return result
	case int32(990):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_header_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('L') {
			state = uint16(917)
			goto next_state
		}
		if lookahead1 == int32('-') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(1007)
			goto next_state
		}
		return result
	case int32(991):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_header_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('N') {
			state = uint16(992)
			goto next_state
		}
		if lookahead1 == int32('-') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(1007)
			goto next_state
		}
		return result
	case int32(992):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_header_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('N') {
			state = uint16(983)
			goto next_state
		}
		if lookahead1 == int32('-') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(1007)
			goto next_state
		}
		return result
	case int32(993):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_header_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('N') {
			state = uint16(1000)
			goto next_state
		}
		if lookahead1 == int32('-') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(1007)
			goto next_state
		}
		return result
	case int32(994):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_header_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('O') {
			state = uint16(993)
			goto next_state
		}
		if lookahead1 == int32('-') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(1007)
			goto next_state
		}
		return result
	case int32(995):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_header_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('O') {
			state = uint16(977)
			goto next_state
		}
		if lookahead1 == int32('-') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(1007)
			goto next_state
		}
		return result
	case int32(996):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_header_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('P') {
			state = uint16(986)
			goto next_state
		}
		if lookahead1 == int32('-') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(1007)
			goto next_state
		}
		return result
	case int32(997):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_header_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('P') {
			state = uint16(971)
			goto next_state
		}
		if lookahead1 == int32('-') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(1007)
			goto next_state
		}
		return result
	case int32(998):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_header_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('Q') {
			state = uint16(990)
			goto next_state
		}
		if lookahead1 == int32('-') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(1007)
			goto next_state
		}
		return result
	case int32(999):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_header_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('S') {
			state = uint16(1002)
			goto next_state
		}
		if lookahead1 == int32('-') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(1007)
			goto next_state
		}
		return result
	case int32(1000):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_header_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('S') {
			state = uint16(917)
			goto next_state
		}
		if lookahead1 == int32('-') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(1007)
			goto next_state
		}
		return result
	case int32(1001):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_header_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('S') {
			state = uint16(995)
			goto next_state
		}
		if lookahead1 == int32('-') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(1007)
			goto next_state
		}
		return result
	case int32(1002):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_header_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('T') {
			state = uint16(917)
			goto next_state
		}
		if lookahead1 == int32('-') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(1007)
			goto next_state
		}
		return result
	case int32(1003):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_header_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('T') {
			state = uint16(987)
			goto next_state
		}
		if lookahead1 == int32('-') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(1007)
			goto next_state
		}
		return result
	case int32(1004):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_header_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('T') {
			state = uint16(978)
			goto next_state
		}
		if lookahead1 == int32('-') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(1007)
			goto next_state
		}
		return result
	case int32(1005):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_header_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('T') {
			state = uint16(997)
			goto next_state
		}
		if lookahead1 == int32('-') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(1007)
			goto next_state
		}
		return result
	case int32(1006):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_header_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('T') {
			state = uint16(982)
			goto next_state
		}
		if lookahead1 == int32('-') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(1007)
			goto next_state
		}
		return result
	case int32(1007):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_header_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('-') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(1007)
			goto next_state
		}
		return result
	case int32(1008):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('$') || lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') || int32(0xa1) <= lookahead1 && lookahead1 <= int32(0xffff) {
			state = uint16(1008)
			goto next_state
		}
		return result
	case int32(1009):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_path_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(1010):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__blank_line_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(1011):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__blank_line_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\n') {
			state = uint16(1010)
			goto next_state
		}
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var map_token = [48]uint16_t{
	1:  uint16(961),
	2:  uint16('\n'),
	3:  uint16(961),
	4:  uint16('\r'),
	5:  uint16(962),
	6:  uint16('#'),
	7:  uint16(906),
	8:  uint16('-'),
	9:  uint16(886),
	10: uint16('/'),
	11: uint16(887),
	12: uint16(':'),
	13: uint16(924),
	14: uint16('<'),
	15: uint16(955),
	16: uint16('='),
	17: uint16(909),
	18: uint16('>'),
	19: uint16(932),
	20: uint16('@'),
	21: uint16(908),
	22: uint16('C'),
	23: uint16(856),
	24: uint16('D'),
	25: uint16(846),
	26: uint16('G'),
	27: uint16(847),
	28: uint16('H'),
	29: uint16(852),
	30: uint16('L'),
	31: uint16(854),
	32: uint16('O'),
	33: uint16(858),
	34: uint16('P'),
	35: uint16(844),
	36: uint16('T'),
	37: uint16(860),
	38: uint16('W'),
	39: uint16(848),
	40: uint16('['),
	41: uint16(889),
	42: uint16('\\'),
	43: uint16(891),
	44: uint16('_'),
	45: uint16(890),
	46: uint16('{'),
	47: uint16(885),
}

var map_token1 = [48]uint16_t{
	0:  uint16('\r'),
	1:  uint16(899),
	2:  uint16(':'),
	3:  uint16(924),
	4:  uint16('='),
	5:  uint16(909),
	6:  uint16('A'),
	7:  uint16(366),
	8:  uint16('B'),
	9:  uint16(318),
	10: uint16('C'),
	11: uint16(603),
	12: uint16('E'),
	13: uint16(784),
	14: uint16('F'),
	15: uint16(319),
	16: uint16('G'),
	17: uint16(320),
	18: uint16('H'),
	19: uint16(312),
	20: uint16('I'),
	21: uint16(223),
	22: uint16('L'),
	23: uint16(416),
	24: uint16('M'),
	25: uint16(410),
	26: uint16('N'),
	27: uint16(417),
	28: uint16('O'),
	29: uint16(263),
	30: uint16('P'),
	31: uint16(321),
	32: uint16('R'),
	33: uint16(329),
	34: uint16('S'),
	35: uint16(407),
	36: uint16('T'),
	37: uint16(411),
	38: uint16('U'),
	39: uint16(288),
	40: uint16('V'),
	41: uint16(325),
	42: uint16('}'),
	43: uint16(795),
	45: uint16(897),
	46: uint16('\n'),
	47: uint16(897),
}

var map_token2 = [48]uint16_t{
	1:  uint16(961),
	2:  uint16('\n'),
	3:  uint16(961),
	4:  uint16('\r'),
	5:  uint16(962),
	6:  uint16('#'),
	7:  uint16(903),
	8:  uint16('-'),
	9:  uint16(873),
	10: uint16('/'),
	11: uint16(874),
	12: uint16('<'),
	13: uint16(953),
	14: uint16('>'),
	15: uint16(932),
	16: uint16('@'),
	17: uint16(939),
	18: uint16('C'),
	19: uint16(826),
	20: uint16('D'),
	21: uint16(822),
	22: uint16('G'),
	23: uint16(824),
	24: uint16('H'),
	25: uint16(823),
	26: uint16('L'),
	27: uint16(825),
	28: uint16('O'),
	29: uint16(827),
	30: uint16('P'),
	31: uint16(820),
	32: uint16('T'),
	33: uint16(828),
	34: uint16('W'),
	35: uint16(821),
	36: uint16('['),
	37: uint16(872),
	38: uint16('m'),
	39: uint16(830),
	40: uint16('q'),
	41: uint16(829),
	42: uint16('{'),
	43: uint16(871),
	44: uint16(0x2028),
	45: uint16(104),
	46: uint16(0x2029),
	47: uint16(104),
}

var map_token3 = [38]uint16_t{
	1:  uint16(863),
	2:  uint16('\n'),
	3:  uint16(1010),
	4:  uint16('\r'),
	5:  uint16(1011),
	6:  uint16('#'),
	7:  uint16(906),
	8:  uint16('/'),
	9:  uint16(887),
	10: uint16('<'),
	11: uint16(930),
	12: uint16('@'),
	13: uint16(937),
	14: uint16('C'),
	15: uint16(856),
	16: uint16('D'),
	17: uint16(846),
	18: uint16('G'),
	19: uint16(847),
	20: uint16('H'),
	21: uint16(852),
	22: uint16('L'),
	23: uint16(854),
	24: uint16('O'),
	25: uint16(858),
	26: uint16('P'),
	27: uint16(844),
	28: uint16('T'),
	29: uint16(860),
	30: uint16('W'),
	31: uint16(848),
	32: uint16('{'),
	33: uint16(888),
	34: uint16('-'),
	35: uint16(890),
	36: uint16('_'),
	37: uint16(890),
}

var map_token4 = [34]uint16_t{
	1:  uint16(863),
	2:  uint16('\n'),
	3:  uint16(1010),
	4:  uint16('\r'),
	5:  uint16(1011),
	6:  uint16('#'),
	7:  uint16(906),
	8:  uint16('/'),
	9:  uint16(887),
	10: uint16('<'),
	11: uint16(930),
	12: uint16('@'),
	13: uint16(937),
	14: uint16('C'),
	15: uint16(857),
	16: uint16('D'),
	17: uint16(849),
	18: uint16('G'),
	19: uint16(850),
	20: uint16('H'),
	21: uint16(853),
	22: uint16('L'),
	23: uint16(855),
	24: uint16('O'),
	25: uint16(859),
	26: uint16('P'),
	27: uint16(845),
	28: uint16('T'),
	29: uint16(861),
	30: uint16('W'),
	31: uint16(851),
	32: uint16('{'),
	33: uint16(888),
}

var map_token5 = [48]uint16_t{
	1:  uint16(868),
	2:  uint16('\n'),
	3:  uint16(897),
	4:  uint16('\r'),
	5:  uint16(898),
	6:  uint16('#'),
	7:  uint16(902),
	8:  uint16('-'),
	9:  uint16(866),
	10: uint16('/'),
	11: uint16(867),
	12: uint16('<'),
	13: uint16(953),
	14: uint16('>'),
	15: uint16(932),
	16: uint16('@'),
	17: uint16(938),
	18: uint16('C'),
	19: uint16(814),
	20: uint16('D'),
	21: uint16(809),
	22: uint16('G'),
	23: uint16(810),
	24: uint16('H'),
	25: uint16(812),
	26: uint16('L'),
	27: uint16(813),
	28: uint16('O'),
	29: uint16(815),
	30: uint16('P'),
	31: uint16(808),
	32: uint16('T'),
	33: uint16(816),
	34: uint16('W'),
	35: uint16(811),
	36: uint16('['),
	37: uint16(870),
	38: uint16('m'),
	39: uint16(817),
	40: uint16('q'),
	41: uint16(818),
	42: uint16('{'),
	43: uint16(869),
	44: uint16(0x2028),
	45: uint16(52),
	46: uint16(0x2029),
	47: uint16(52),
}

var map_token6 = [48]uint16_t{
	1:  uint16(868),
	2:  uint16('\n'),
	3:  uint16(897),
	4:  uint16('\r'),
	5:  uint16(898),
	6:  uint16('#'),
	7:  uint16(902),
	8:  uint16('-'),
	9:  uint16(866),
	10: uint16('/'),
	11: uint16(867),
	12: uint16('<'),
	13: uint16(953),
	14: uint16('>'),
	15: uint16(932),
	16: uint16('@'),
	17: uint16(938),
	18: uint16('C'),
	19: uint16(814),
	20: uint16('D'),
	21: uint16(809),
	22: uint16('G'),
	23: uint16(810),
	24: uint16('H'),
	25: uint16(812),
	26: uint16('L'),
	27: uint16(813),
	28: uint16('O'),
	29: uint16(815),
	30: uint16('P'),
	31: uint16(808),
	32: uint16('T'),
	33: uint16(816),
	34: uint16('W'),
	35: uint16(811),
	36: uint16('['),
	37: uint16(865),
	38: uint16('m'),
	39: uint16(817),
	40: uint16('q'),
	41: uint16(818),
	42: uint16('{'),
	43: uint16(864),
	44: uint16(0x2028),
	45: uint16(52),
	46: uint16(0x2029),
	47: uint16(52),
}

var map_token7 = [48]uint16_t{
	1:  uint16(967),
	2:  uint16('\n'),
	3:  uint16(966),
	4:  uint16('\r'),
	5:  uint16(967),
	6:  uint16('#'),
	7:  uint16(904),
	8:  uint16('-'),
	9:  uint16(878),
	10: uint16('/'),
	11: uint16(879),
	12: uint16('<'),
	13: uint16(954),
	14: uint16('>'),
	15: uint16(932),
	16: uint16('@'),
	17: uint16(940),
	18: uint16('C'),
	19: uint16(838),
	20: uint16('D'),
	21: uint16(834),
	22: uint16('G'),
	23: uint16(836),
	24: uint16('H'),
	25: uint16(835),
	26: uint16('L'),
	27: uint16(837),
	28: uint16('O'),
	29: uint16(839),
	30: uint16('P'),
	31: uint16(832),
	32: uint16('T'),
	33: uint16(840),
	34: uint16('W'),
	35: uint16(833),
	36: uint16('['),
	37: uint16(877),
	38: uint16('m'),
	39: uint16(842),
	40: uint16('q'),
	41: uint16(841),
	42: uint16('{'),
	43: uint16(876),
	44: uint16(0x2028),
	45: uint16(157),
	46: uint16(0x2029),
	47: uint16(157),
}

var map_token8 = [48]uint16_t{
	1:  uint16(967),
	2:  uint16('\n'),
	3:  uint16(966),
	4:  uint16('\r'),
	5:  uint16(967),
	6:  uint16('#'),
	7:  uint16(904),
	8:  uint16('-'),
	9:  uint16(878),
	10: uint16('/'),
	11: uint16(879),
	12: uint16('<'),
	13: uint16(954),
	14: uint16('>'),
	15: uint16(932),
	16: uint16('@'),
	17: uint16(940),
	18: uint16('C'),
	19: uint16(838),
	20: uint16('D'),
	21: uint16(834),
	22: uint16('G'),
	23: uint16(836),
	24: uint16('H'),
	25: uint16(835),
	26: uint16('L'),
	27: uint16(837),
	28: uint16('O'),
	29: uint16(839),
	30: uint16('P'),
	31: uint16(832),
	32: uint16('T'),
	33: uint16(840),
	34: uint16('W'),
	35: uint16(833),
	36: uint16('['),
	37: uint16(882),
	38: uint16('m'),
	39: uint16(842),
	40: uint16('q'),
	41: uint16(841),
	42: uint16('{'),
	43: uint16(881),
	44: uint16(0x2028),
	45: uint16(157),
	46: uint16(0x2029),
	47: uint16(157),
}

var ts_lex_modes = [261]TSLexMode{
	0: {},
	1: {
		Flex_state: uint16(801),
	},
	2: {
		Flex_state: uint16(802),
	},
	3: {
		Flex_state: uint16(802),
	},
	4: {
		Flex_state: uint16(802),
	},
	5: {
		Flex_state: uint16(802),
	},
	6: {
		Flex_state: uint16(802),
	},
	7: {
		Flex_state: uint16(802),
	},
	8: {
		Flex_state: uint16(802),
	},
	9: {
		Flex_state: uint16(802),
	},
	10: {
		Flex_state: uint16(802),
	},
	11: {
		Flex_state: uint16(802),
	},
	12: {
		Flex_state: uint16(802),
	},
	13: {
		Flex_state: uint16(802),
	},
	14: {
		Flex_state: uint16(802),
	},
	15: {
		Flex_state: uint16(802),
	},
	16: {
		Flex_state: uint16(802),
	},
	17: {
		Flex_state: uint16(802),
	},
	18: {
		Flex_state: uint16(802),
	},
	19: {
		Flex_state: uint16(802),
	},
	20: {
		Flex_state: uint16(802),
	},
	21: {
		Flex_state: uint16(802),
	},
	22: {
		Flex_state: uint16(802),
	},
	23: {
		Flex_state: uint16(802),
	},
	24: {
		Flex_state: uint16(802),
	},
	25: {
		Flex_state: uint16(802),
	},
	26: {
		Flex_state: uint16(802),
	},
	27: {
		Flex_state: uint16(801),
	},
	28: {
		Flex_state: uint16(799),
	},
	29: {
		Flex_state: uint16(801),
	},
	30: {
		Flex_state: uint16(799),
	},
	31: {
		Flex_state: uint16(799),
	},
	32: {
		Flex_state: uint16(801),
	},
	33: {
		Flex_state: uint16(801),
	},
	34: {
		Flex_state: uint16(804),
	},
	35: {
		Flex_state: uint16(804),
	},
	36: {
		Flex_state: uint16(799),
	},
	37: {
		Flex_state: uint16(805),
	},
	38: {
		Flex_state: uint16(799),
	},
	39: {
		Flex_state: uint16(799),
	},
	40: {
		Flex_state: uint16(799),
	},
	41: {
		Flex_state: uint16(805),
	},
	42: {
		Flex_state: uint16(805),
	},
	43: {
		Flex_state: uint16(805),
	},
	44: {
		Flex_state: uint16(799),
	},
	45: {
		Flex_state: uint16(799),
	},
	46: {
		Flex_state: uint16(799),
	},
	47: {
		Flex_state: uint16(803),
	},
	48: {
		Flex_state: uint16(802),
	},
	49: {
		Flex_state: uint16(803),
	},
	50: {
		Flex_state: uint16(803),
	},
	51: {
		Flex_state: uint16(803),
	},
	52: {
		Flex_state: uint16(802),
	},
	53: {
		Flex_state: uint16(802),
	},
	54: {
		Flex_state: uint16(802),
	},
	55: {
		Flex_state: uint16(802),
	},
	56: {
		Flex_state: uint16(802),
	},
	57: {
		Flex_state: uint16(802),
	},
	58: {
		Flex_state: uint16(802),
	},
	59: {
		Flex_state: uint16(802),
	},
	60: {
		Flex_state: uint16(802),
	},
	61: {
		Flex_state: uint16(802),
	},
	62: {
		Flex_state: uint16(802),
	},
	63: {
		Flex_state: uint16(802),
	},
	64: {
		Flex_state: uint16(802),
	},
	65: {
		Flex_state: uint16(802),
	},
	66: {
		Flex_state: uint16(802),
	},
	67: {
		Flex_state: uint16(802),
	},
	68: {
		Flex_state: uint16(802),
	},
	69: {
		Flex_state: uint16(800),
	},
	70: {
		Flex_state: uint16(800),
	},
	71: {
		Flex_state: uint16(800),
	},
	72: {
		Flex_state: uint16(800),
	},
	73: {
		Flex_state: uint16(800),
	},
	74: {
		Flex_state: uint16(800),
	},
	75: {
		Flex_state: uint16(800),
	},
	76: {
		Flex_state: uint16(800),
	},
	77: {
		Flex_state: uint16(800),
	},
	78: {
		Flex_state: uint16(800),
	},
	79: {
		Flex_state: uint16(800),
	},
	80: {
		Flex_state: uint16(800),
	},
	81: {
		Flex_state: uint16(800),
	},
	82: {
		Flex_state: uint16(800),
	},
	83: {
		Flex_state: uint16(800),
	},
	84: {
		Flex_state: uint16(800),
	},
	85: {
		Flex_state: uint16(800),
	},
	86: {
		Flex_state: uint16(800),
	},
	87: {
		Flex_state: uint16(800),
	},
	88: {
		Flex_state: uint16(800),
	},
	89: {
		Flex_state: uint16(800),
	},
	90: {
		Flex_state: uint16(800),
	},
	91: {
		Flex_state: uint16(800),
	},
	92: {
		Flex_state: uint16(800),
	},
	93: {
		Flex_state: uint16(800),
	},
	94: {
		Flex_state: uint16(800),
	},
	95: {
		Flex_state: uint16(800),
	},
	96: {
		Flex_state: uint16(800),
	},
	97: {
		Flex_state: uint16(801),
	},
	98: {
		Flex_state: uint16(801),
	},
	99: {
		Flex_state: uint16(801),
	},
	100: {
		Flex_state: uint16(801),
	},
	101: {
		Flex_state: uint16(801),
	},
	102: {
		Flex_state: uint16(801),
	},
	103: {
		Flex_state: uint16(801),
	},
	104: {
		Flex_state: uint16(801),
	},
	105: {
		Flex_state: uint16(801),
	},
	106: {
		Flex_state: uint16(801),
	},
	107: {
		Flex_state: uint16(801),
	},
	108: {
		Flex_state: uint16(801),
	},
	109: {
		Flex_state: uint16(801),
	},
	110: {
		Flex_state: uint16(801),
	},
	111: {
		Flex_state: uint16(801),
	},
	112: {
		Flex_state: uint16(801),
	},
	113: {
		Flex_state: uint16(801),
	},
	114: {
		Flex_state: uint16(801),
	},
	115: {
		Flex_state: uint16(801),
	},
	116: {
		Flex_state: uint16(801),
	},
	117: {
		Flex_state: uint16(2),
	},
	118: {
		Flex_state: uint16(2),
	},
	119: {
		Flex_state: uint16(317),
	},
	120: {
		Flex_state: uint16(317),
	},
	121: {
		Flex_state: uint16(2),
	},
	122: {
		Flex_state: uint16(3),
	},
	123: {
		Flex_state: uint16(2),
	},
	124: {
		Flex_state: uint16(2),
	},
	125: {
		Flex_state: uint16(2),
	},
	126: {
		Flex_state: uint16(3),
	},
	127: {
		Flex_state: uint16(3),
	},
	128: {
		Flex_state: uint16(2),
	},
	129: {
		Flex_state: uint16(1),
	},
	130: {
		Flex_state: uint16(1),
	},
	131: {
		Flex_state: uint16(2),
	},
	132: {
		Flex_state: uint16(3),
	},
	133: {
		Flex_state: uint16(2),
	},
	134: {
		Flex_state: uint16(3),
	},
	135: {
		Flex_state: uint16(3),
	},
	136: {
		Flex_state: uint16(3),
	},
	137: {
		Flex_state: uint16(1),
	},
	138: {
		Flex_state: uint16(3),
	},
	139: {
		Flex_state: uint16(1),
	},
	140: {
		Flex_state: uint16(3),
	},
	141: {
		Flex_state: uint16(160),
	},
	142: {
		Flex_state: uint16(3),
	},
	143: {
		Flex_state: uint16(3),
	},
	144: {
		Flex_state: uint16(3),
	},
	145: {
		Flex_state: uint16(3),
	},
	146: {
		Flex_state: uint16(3),
	},
	147: {
		Flex_state: uint16(1),
	},
	148: {
		Flex_state: uint16(1),
	},
	149: {
		Flex_state: uint16(3),
	},
	150: {
		Flex_state: uint16(160),
	},
	151: {
		Flex_state: uint16(1),
	},
	152: {
		Flex_state: uint16(1),
	},
	153: {
		Flex_state: uint16(1),
	},
	154: {
		Flex_state: uint16(1),
	},
	155: {
		Flex_state: uint16(3),
	},
	156: {
		Flex_state: uint16(3),
	},
	157: {
		Flex_state: uint16(160),
	},
	158: {
		Flex_state: uint16(160),
	},
	159: {
		Flex_state: uint16(160),
	},
	160: {
		Flex_state: uint16(160),
	},
	161: {
		Flex_state: uint16(3),
	},
	162: {
		Flex_state: uint16(3),
	},
	163: {
		Flex_state: uint16(3),
	},
	164: {
		Flex_state: uint16(3),
	},
	165: {
		Flex_state: uint16(161),
	},
	166: {
		Flex_state: uint16(161),
	},
	167: {
		Flex_state: uint16(108),
	},
	168: {
		Flex_state: uint16(108),
	},
	169: {
		Flex_state: uint16(108),
	},
	170: {
		Flex_state: uint16(167),
	},
	171: {
		Flex_state: uint16(167),
	},
	172: {
		Flex_state: uint16(167),
	},
	173: {
		Flex_state: uint16(167),
	},
	174: {
		Flex_state: uint16(167),
	},
	175: {
		Flex_state: uint16(167),
	},
	176: {
		Flex_state: uint16(108),
	},
	177: {
		Flex_state: uint16(161),
	},
	178: {
		Flex_state: uint16(162),
	},
	179: {
		Flex_state: uint16(169),
	},
	180: {
		Flex_state: uint16(168),
	},
	181: {},
	182: {
		Flex_state: uint16(167),
	},
	183: {
		Flex_state: uint16(168),
	},
	184: {
		Flex_state: uint16(168),
	},
	185: {
		Flex_state: uint16(167),
	},
	186: {
		Flex_state: uint16(167),
	},
	187: {
		Flex_state: uint16(167),
	},
	188: {
		Flex_state: uint16(162),
	},
	189: {
		Flex_state: uint16(167),
	},
	190: {
		Flex_state: uint16(167),
	},
	191: {
		Flex_state: uint16(167),
	},
	192: {
		Flex_state: uint16(167),
	},
	193: {
		Flex_state: uint16(162),
	},
	194: {
		Flex_state: uint16(167),
	},
	195: {
		Flex_state: uint16(168),
	},
	196: {
		Flex_state: uint16(167),
	},
	197: {
		Flex_state: uint16(167),
	},
	198: {
		Flex_state: uint16(167),
	},
	199: {
		Flex_state: uint16(167),
	},
	200: {
		Flex_state: uint16(167),
	},
	201: {
		Flex_state: uint16(801),
	},
	202: {},
	203: {
		Flex_state: uint16(167),
	},
	204: {
		Flex_state: uint16(167),
	},
	205: {
		Flex_state: uint16(167),
	},
	206: {
		Flex_state: uint16(167),
	},
	207: {
		Flex_state: uint16(104),
	},
	208: {
		Flex_state: uint16(801),
	},
	209: {
		Flex_state: uint16(168),
	},
	210: {
		Flex_state: uint16(168),
	},
	211: {
		Flex_state: uint16(167),
	},
	212: {
		Flex_state: uint16(801),
	},
	213: {
		Flex_state: uint16(801),
	},
	214: {
		Flex_state: uint16(801),
	},
	215: {
		Flex_state: uint16(167),
	},
	216: {},
	217: {
		Flex_state: uint16(166),
	},
	218: {
		Flex_state: uint16(167),
	},
	219: {
		Flex_state: uint16(801),
	},
	220: {
		Flex_state: uint16(167),
	},
	221: {
		Flex_state: uint16(168),
	},
	222: {
		Flex_state: uint16(167),
	},
	223: {
		Flex_state: uint16(167),
	},
	224: {
		Flex_state: uint16(167),
	},
	225: {
		Flex_state: uint16(801),
	},
	226: {
		Flex_state: uint16(167),
	},
	227: {
		Flex_state: uint16(166),
	},
	228: {
		Flex_state: uint16(167),
	},
	229: {
		Flex_state: uint16(168),
	},
	230: {
		Flex_state: uint16(168),
	},
	231: {
		Flex_state: uint16(167),
	},
	232: {
		Flex_state: uint16(167),
	},
	233: {
		Flex_state: uint16(167),
	},
	234: {
		Flex_state: uint16(167),
	},
	235: {
		Flex_state: uint16(169),
	},
	236: {
		Flex_state: uint16(167),
	},
	237: {
		Flex_state: uint16(166),
	},
	238: {
		Flex_state: uint16(167),
	},
	239: {
		Flex_state: uint16(167),
	},
	240: {
		Flex_state: uint16(167),
	},
	241: {
		Flex_state: uint16(167),
	},
	242: {
		Flex_state: uint16(167),
	},
	243: {
		Flex_state: uint16(167),
	},
	244: {
		Flex_state: uint16(167),
	},
	245: {
		Flex_state: uint16(167),
	},
	246: {
		Flex_state: uint16(167),
	},
	247: {
		Flex_state: uint16(167),
	},
	248: {
		Flex_state: uint16(167),
	},
	249: {
		Flex_state: uint16(166),
	},
	250: {
		Flex_state: uint16(166),
	},
	251: {
		Flex_state: uint16(801),
	},
	252: {
		Flex_state: uint16(167),
	},
	253: {
		Flex_state: uint16(167),
	},
	254: {
		Flex_state: uint16(166),
	},
	255: {
		Flex_state: uint16(166),
	},
	256: {
		Flex_state: uint16(166),
	},
	257: {
		Flex_state: uint16(167),
	},
	258: {
		Flex_state: uint16(166),
	},
	259: {
		Flex_state: uint16(801),
	},
	260: {
		Flex_state: uint16(169),
	},
}

var ts_parse_table = [2][82]uint16_t{
	0: {
		0:  uint16(1),
		1:  uint16(1),
		2:  uint16(1),
		3:  uint16(1),
		4:  uint16(1),
		6:  uint16(1),
		7:  uint16(1),
		8:  uint16(1),
		9:  uint16(1),
		10: uint16(1),
		11: uint16(1),
		12: uint16(1),
		15: uint16(1),
		16: uint16(1),
		18: uint16(1),
		20: uint16(1),
		21: uint16(1),
		23: uint16(1),
		24: uint16(1),
		25: uint16(1),
		26: uint16(1),
		28: uint16(1),
		29: uint16(1),
		30: uint16(1),
		32: uint16(1),
		36: uint16(1),
		38: uint16(1),
		39: uint16(1),
	},
	1: {
		0:  uint16(3),
		1:  uint16(5),
		2:  uint16(5),
		3:  uint16(7),
		6:  uint16(9),
		10: uint16(11),
		11: uint16(13),
		12: uint16(15),
		16: uint16(17),
		18: uint16(19),
		24: uint16(21),
		39: uint16(23),
		40: uint16(216),
		41: uint16(32),
		42: uint16(112),
		43: uint16(116),
		44: uint16(33),
		45: uint16(29),
		46: uint16(99),
		47: uint16(214),
		48: uint16(132),
		49: uint16(200),
		50: uint16(104),
		51: uint16(105),
		53: uint16(132),
		54: uint16(32),
		58: uint16(32),
		71: uint16(32),
		72: uint16(29),
	},
}

var ts_small_parse_table = [5696]uint16_t{
	0:    uint16(20),
	1:    uint16(29),
	2:    uint16(1),
	3:    uint16(aux_sym_WS_token1),
	4:    uint16(31),
	5:    uint16(1),
	6:    uint16(aux_sym_NL_token1),
	7:    uint16(33),
	8:    uint16(1),
	9:    uint16(aux_sym_COMMENT_PREFIX_token1),
	10:   uint16(35),
	11:   uint16(1),
	12:   uint16(anon_sym_GT),
	13:   uint16(37),
	14:   uint16(1),
	15:   uint16(aux_sym_res_redirect_token1),
	16:   uint16(39),
	17:   uint16(1),
	18:   uint16(aux_sym_xml_body_token1),
	19:   uint16(41),
	20:   uint16(1),
	21:   uint16(aux_sym_json_body_token1),
	22:   uint16(43),
	23:   uint16(1),
	24:   uint16(aux_sym_graphql_data_token1),
	25:   uint16(45),
	26:   uint16(1),
	27:   uint16(anon_sym_LT2),
	28:   uint16(47),
	29:   uint16(1),
	30:   uint16(anon_sym_DASH_DASH),
	31:   uint16(49),
	32:   uint16(1),
	33:   uint16(aux_sym_raw_body_token1),
	34:   uint16(51),
	35:   uint16(1),
	36:   uint16(aux_sym__blank_line_token1),
	37:   uint16(47),
	38:   uint16(1),
	39:   uint16(sym_graphql_data),
	40:   uint16(59),
	41:   uint16(1),
	42:   uint16(sym__var_comment),
	43:   uint16(205),
	44:   uint16(1),
	45:   uint16(sym_external_body),
	46:   uint16(25),
	47:   uint16(2),
	49:   uint16(aux_sym_request_separator_token1),
	50:   uint16(48),
	51:   uint16(2),
	52:   uint16(sym__blank_line),
	53:   uint16(aux_sym___body_repeat1),
	54:   uint16(25),
	55:   uint16(3),
	56:   uint16(sym_res_handler_script),
	57:   uint16(sym_res_redirect),
	58:   uint16(aux_sym___body_repeat2),
	59:   uint16(68),
	60:   uint16(6),
	61:   uint16(sym_xml_body),
	62:   uint16(sym_json_body),
	63:   uint16(sym_graphql_body),
	64:   uint16(sym__external_body),
	65:   uint16(sym_multipart_form_data),
	66:   uint16(sym_raw_body),
	67:   uint16(27),
	68:   uint16(7),
	69:   uint16(aux_sym_WORD_CHAR_token1),
	70:   uint16(aux_sym_PUNCTUATION_token1),
	71:   uint16(sym_method),
	72:   uint16(aux_sym_http_version_token1),
	73:   uint16(anon_sym_LBRACE_LBRACE),
	74:   uint16(anon_sym_LT),
	75:   uint16(anon_sym_AT2),
	76:   uint16(20),
	77:   uint16(29),
	78:   uint16(1),
	79:   uint16(aux_sym_WS_token1),
	80:   uint16(33),
	81:   uint16(1),
	82:   uint16(aux_sym_COMMENT_PREFIX_token1),
	83:   uint16(35),
	84:   uint16(1),
	85:   uint16(anon_sym_GT),
	86:   uint16(37),
	87:   uint16(1),
	88:   uint16(aux_sym_res_redirect_token1),
	89:   uint16(39),
	90:   uint16(1),
	91:   uint16(aux_sym_xml_body_token1),
	92:   uint16(41),
	93:   uint16(1),
	94:   uint16(aux_sym_json_body_token1),
	95:   uint16(43),
	96:   uint16(1),
	97:   uint16(aux_sym_graphql_data_token1),
	98:   uint16(45),
	99:   uint16(1),
	100:  uint16(anon_sym_LT2),
	101:  uint16(47),
	102:  uint16(1),
	103:  uint16(anon_sym_DASH_DASH),
	104:  uint16(49),
	105:  uint16(1),
	106:  uint16(aux_sym_raw_body_token1),
	107:  uint16(51),
	108:  uint16(1),
	109:  uint16(aux_sym__blank_line_token1),
	110:  uint16(57),
	111:  uint16(1),
	112:  uint16(aux_sym_NL_token1),
	113:  uint16(47),
	114:  uint16(1),
	115:  uint16(sym_graphql_data),
	116:  uint16(59),
	117:  uint16(1),
	118:  uint16(sym__var_comment),
	119:  uint16(205),
	120:  uint16(1),
	121:  uint16(sym_external_body),
	122:  uint16(53),
	123:  uint16(2),
	125:  uint16(aux_sym_request_separator_token1),
	126:  uint16(48),
	127:  uint16(2),
	128:  uint16(sym__blank_line),
	129:  uint16(aux_sym___body_repeat1),
	130:  uint16(26),
	131:  uint16(3),
	132:  uint16(sym_res_handler_script),
	133:  uint16(sym_res_redirect),
	134:  uint16(aux_sym___body_repeat2),
	135:  uint16(68),
	136:  uint16(6),
	137:  uint16(sym_xml_body),
	138:  uint16(sym_json_body),
	139:  uint16(sym_graphql_body),
	140:  uint16(sym__external_body),
	141:  uint16(sym_multipart_form_data),
	142:  uint16(sym_raw_body),
	143:  uint16(55),
	144:  uint16(7),
	145:  uint16(aux_sym_WORD_CHAR_token1),
	146:  uint16(aux_sym_PUNCTUATION_token1),
	147:  uint16(sym_method),
	148:  uint16(aux_sym_http_version_token1),
	149:  uint16(anon_sym_LBRACE_LBRACE),
	150:  uint16(anon_sym_LT),
	151:  uint16(anon_sym_AT2),
	152:  uint16(20),
	153:  uint16(29),
	154:  uint16(1),
	155:  uint16(aux_sym_WS_token1),
	156:  uint16(33),
	157:  uint16(1),
	158:  uint16(aux_sym_COMMENT_PREFIX_token1),
	159:  uint16(35),
	160:  uint16(1),
	161:  uint16(anon_sym_GT),
	162:  uint16(37),
	163:  uint16(1),
	164:  uint16(aux_sym_res_redirect_token1),
	165:  uint16(39),
	166:  uint16(1),
	167:  uint16(aux_sym_xml_body_token1),
	168:  uint16(41),
	169:  uint16(1),
	170:  uint16(aux_sym_json_body_token1),
	171:  uint16(43),
	172:  uint16(1),
	173:  uint16(aux_sym_graphql_data_token1),
	174:  uint16(45),
	175:  uint16(1),
	176:  uint16(anon_sym_LT2),
	177:  uint16(47),
	178:  uint16(1),
	179:  uint16(anon_sym_DASH_DASH),
	180:  uint16(49),
	181:  uint16(1),
	182:  uint16(aux_sym_raw_body_token1),
	183:  uint16(51),
	184:  uint16(1),
	185:  uint16(aux_sym__blank_line_token1),
	186:  uint16(63),
	187:  uint16(1),
	188:  uint16(aux_sym_NL_token1),
	189:  uint16(47),
	190:  uint16(1),
	191:  uint16(sym_graphql_data),
	192:  uint16(59),
	193:  uint16(1),
	194:  uint16(sym__var_comment),
	195:  uint16(205),
	196:  uint16(1),
	197:  uint16(sym_external_body),
	198:  uint16(59),
	199:  uint16(2),
	201:  uint16(aux_sym_request_separator_token1),
	202:  uint16(48),
	203:  uint16(2),
	204:  uint16(sym__blank_line),
	205:  uint16(aux_sym___body_repeat1),
	206:  uint16(15),
	207:  uint16(3),
	208:  uint16(sym_res_handler_script),
	209:  uint16(sym_res_redirect),
	210:  uint16(aux_sym___body_repeat2),
	211:  uint16(68),
	212:  uint16(6),
	213:  uint16(sym_xml_body),
	214:  uint16(sym_json_body),
	215:  uint16(sym_graphql_body),
	216:  uint16(sym__external_body),
	217:  uint16(sym_multipart_form_data),
	218:  uint16(sym_raw_body),
	219:  uint16(61),
	220:  uint16(7),
	221:  uint16(aux_sym_WORD_CHAR_token1),
	222:  uint16(aux_sym_PUNCTUATION_token1),
	223:  uint16(sym_method),
	224:  uint16(aux_sym_http_version_token1),
	225:  uint16(anon_sym_LBRACE_LBRACE),
	226:  uint16(anon_sym_LT),
	227:  uint16(anon_sym_AT2),
	228:  uint16(20),
	229:  uint16(29),
	230:  uint16(1),
	231:  uint16(aux_sym_WS_token1),
	232:  uint16(33),
	233:  uint16(1),
	234:  uint16(aux_sym_COMMENT_PREFIX_token1),
	235:  uint16(35),
	236:  uint16(1),
	237:  uint16(anon_sym_GT),
	238:  uint16(37),
	239:  uint16(1),
	240:  uint16(aux_sym_res_redirect_token1),
	241:  uint16(39),
	242:  uint16(1),
	243:  uint16(aux_sym_xml_body_token1),
	244:  uint16(41),
	245:  uint16(1),
	246:  uint16(aux_sym_json_body_token1),
	247:  uint16(43),
	248:  uint16(1),
	249:  uint16(aux_sym_graphql_data_token1),
	250:  uint16(45),
	251:  uint16(1),
	252:  uint16(anon_sym_LT2),
	253:  uint16(47),
	254:  uint16(1),
	255:  uint16(anon_sym_DASH_DASH),
	256:  uint16(49),
	257:  uint16(1),
	258:  uint16(aux_sym_raw_body_token1),
	259:  uint16(51),
	260:  uint16(1),
	261:  uint16(aux_sym__blank_line_token1),
	262:  uint16(69),
	263:  uint16(1),
	264:  uint16(aux_sym_NL_token1),
	265:  uint16(47),
	266:  uint16(1),
	267:  uint16(sym_graphql_data),
	268:  uint16(59),
	269:  uint16(1),
	270:  uint16(sym__var_comment),
	271:  uint16(205),
	272:  uint16(1),
	273:  uint16(sym_external_body),
	274:  uint16(65),
	275:  uint16(2),
	277:  uint16(aux_sym_request_separator_token1),
	278:  uint16(48),
	279:  uint16(2),
	280:  uint16(sym__blank_line),
	281:  uint16(aux_sym___body_repeat1),
	282:  uint16(14),
	283:  uint16(3),
	284:  uint16(sym_res_handler_script),
	285:  uint16(sym_res_redirect),
	286:  uint16(aux_sym___body_repeat2),
	287:  uint16(68),
	288:  uint16(6),
	289:  uint16(sym_xml_body),
	290:  uint16(sym_json_body),
	291:  uint16(sym_graphql_body),
	292:  uint16(sym__external_body),
	293:  uint16(sym_multipart_form_data),
	294:  uint16(sym_raw_body),
	295:  uint16(67),
	296:  uint16(7),
	297:  uint16(aux_sym_WORD_CHAR_token1),
	298:  uint16(aux_sym_PUNCTUATION_token1),
	299:  uint16(sym_method),
	300:  uint16(aux_sym_http_version_token1),
	301:  uint16(anon_sym_LBRACE_LBRACE),
	302:  uint16(anon_sym_LT),
	303:  uint16(anon_sym_AT2),
	304:  uint16(20),
	305:  uint16(29),
	306:  uint16(1),
	307:  uint16(aux_sym_WS_token1),
	308:  uint16(33),
	309:  uint16(1),
	310:  uint16(aux_sym_COMMENT_PREFIX_token1),
	311:  uint16(35),
	312:  uint16(1),
	313:  uint16(anon_sym_GT),
	314:  uint16(37),
	315:  uint16(1),
	316:  uint16(aux_sym_res_redirect_token1),
	317:  uint16(39),
	318:  uint16(1),
	319:  uint16(aux_sym_xml_body_token1),
	320:  uint16(41),
	321:  uint16(1),
	322:  uint16(aux_sym_json_body_token1),
	323:  uint16(43),
	324:  uint16(1),
	325:  uint16(aux_sym_graphql_data_token1),
	326:  uint16(45),
	327:  uint16(1),
	328:  uint16(anon_sym_LT2),
	329:  uint16(47),
	330:  uint16(1),
	331:  uint16(anon_sym_DASH_DASH),
	332:  uint16(49),
	333:  uint16(1),
	334:  uint16(aux_sym_raw_body_token1),
	335:  uint16(51),
	336:  uint16(1),
	337:  uint16(aux_sym__blank_line_token1),
	338:  uint16(75),
	339:  uint16(1),
	340:  uint16(aux_sym_NL_token1),
	341:  uint16(47),
	342:  uint16(1),
	343:  uint16(sym_graphql_data),
	344:  uint16(59),
	345:  uint16(1),
	346:  uint16(sym__var_comment),
	347:  uint16(205),
	348:  uint16(1),
	349:  uint16(sym_external_body),
	350:  uint16(71),
	351:  uint16(2),
	353:  uint16(aux_sym_request_separator_token1),
	354:  uint16(48),
	355:  uint16(2),
	356:  uint16(sym__blank_line),
	357:  uint16(aux_sym___body_repeat1),
	358:  uint16(18),
	359:  uint16(3),
	360:  uint16(sym_res_handler_script),
	361:  uint16(sym_res_redirect),
	362:  uint16(aux_sym___body_repeat2),
	363:  uint16(68),
	364:  uint16(6),
	365:  uint16(sym_xml_body),
	366:  uint16(sym_json_body),
	367:  uint16(sym_graphql_body),
	368:  uint16(sym__external_body),
	369:  uint16(sym_multipart_form_data),
	370:  uint16(sym_raw_body),
	371:  uint16(73),
	372:  uint16(7),
	373:  uint16(aux_sym_WORD_CHAR_token1),
	374:  uint16(aux_sym_PUNCTUATION_token1),
	375:  uint16(sym_method),
	376:  uint16(aux_sym_http_version_token1),
	377:  uint16(anon_sym_LBRACE_LBRACE),
	378:  uint16(anon_sym_LT),
	379:  uint16(anon_sym_AT2),
	380:  uint16(20),
	381:  uint16(29),
	382:  uint16(1),
	383:  uint16(aux_sym_WS_token1),
	384:  uint16(33),
	385:  uint16(1),
	386:  uint16(aux_sym_COMMENT_PREFIX_token1),
	387:  uint16(35),
	388:  uint16(1),
	389:  uint16(anon_sym_GT),
	390:  uint16(37),
	391:  uint16(1),
	392:  uint16(aux_sym_res_redirect_token1),
	393:  uint16(39),
	394:  uint16(1),
	395:  uint16(aux_sym_xml_body_token1),
	396:  uint16(41),
	397:  uint16(1),
	398:  uint16(aux_sym_json_body_token1),
	399:  uint16(43),
	400:  uint16(1),
	401:  uint16(aux_sym_graphql_data_token1),
	402:  uint16(45),
	403:  uint16(1),
	404:  uint16(anon_sym_LT2),
	405:  uint16(47),
	406:  uint16(1),
	407:  uint16(anon_sym_DASH_DASH),
	408:  uint16(49),
	409:  uint16(1),
	410:  uint16(aux_sym_raw_body_token1),
	411:  uint16(51),
	412:  uint16(1),
	413:  uint16(aux_sym__blank_line_token1),
	414:  uint16(81),
	415:  uint16(1),
	416:  uint16(aux_sym_NL_token1),
	417:  uint16(47),
	418:  uint16(1),
	419:  uint16(sym_graphql_data),
	420:  uint16(59),
	421:  uint16(1),
	422:  uint16(sym__var_comment),
	423:  uint16(205),
	424:  uint16(1),
	425:  uint16(sym_external_body),
	426:  uint16(77),
	427:  uint16(2),
	429:  uint16(aux_sym_request_separator_token1),
	430:  uint16(48),
	431:  uint16(2),
	432:  uint16(sym__blank_line),
	433:  uint16(aux_sym___body_repeat1),
	434:  uint16(17),
	435:  uint16(3),
	436:  uint16(sym_res_handler_script),
	437:  uint16(sym_res_redirect),
	438:  uint16(aux_sym___body_repeat2),
	439:  uint16(68),
	440:  uint16(6),
	441:  uint16(sym_xml_body),
	442:  uint16(sym_json_body),
	443:  uint16(sym_graphql_body),
	444:  uint16(sym__external_body),
	445:  uint16(sym_multipart_form_data),
	446:  uint16(sym_raw_body),
	447:  uint16(79),
	448:  uint16(7),
	449:  uint16(aux_sym_WORD_CHAR_token1),
	450:  uint16(aux_sym_PUNCTUATION_token1),
	451:  uint16(sym_method),
	452:  uint16(aux_sym_http_version_token1),
	453:  uint16(anon_sym_LBRACE_LBRACE),
	454:  uint16(anon_sym_LT),
	455:  uint16(anon_sym_AT2),
	456:  uint16(20),
	457:  uint16(29),
	458:  uint16(1),
	459:  uint16(aux_sym_WS_token1),
	460:  uint16(33),
	461:  uint16(1),
	462:  uint16(aux_sym_COMMENT_PREFIX_token1),
	463:  uint16(35),
	464:  uint16(1),
	465:  uint16(anon_sym_GT),
	466:  uint16(37),
	467:  uint16(1),
	468:  uint16(aux_sym_res_redirect_token1),
	469:  uint16(39),
	470:  uint16(1),
	471:  uint16(aux_sym_xml_body_token1),
	472:  uint16(41),
	473:  uint16(1),
	474:  uint16(aux_sym_json_body_token1),
	475:  uint16(43),
	476:  uint16(1),
	477:  uint16(aux_sym_graphql_data_token1),
	478:  uint16(45),
	479:  uint16(1),
	480:  uint16(anon_sym_LT2),
	481:  uint16(47),
	482:  uint16(1),
	483:  uint16(anon_sym_DASH_DASH),
	484:  uint16(49),
	485:  uint16(1),
	486:  uint16(aux_sym_raw_body_token1),
	487:  uint16(51),
	488:  uint16(1),
	489:  uint16(aux_sym__blank_line_token1),
	490:  uint16(87),
	491:  uint16(1),
	492:  uint16(aux_sym_NL_token1),
	493:  uint16(47),
	494:  uint16(1),
	495:  uint16(sym_graphql_data),
	496:  uint16(59),
	497:  uint16(1),
	498:  uint16(sym__var_comment),
	499:  uint16(205),
	500:  uint16(1),
	501:  uint16(sym_external_body),
	502:  uint16(83),
	503:  uint16(2),
	505:  uint16(aux_sym_request_separator_token1),
	506:  uint16(48),
	507:  uint16(2),
	508:  uint16(sym__blank_line),
	509:  uint16(aux_sym___body_repeat1),
	510:  uint16(19),
	511:  uint16(3),
	512:  uint16(sym_res_handler_script),
	513:  uint16(sym_res_redirect),
	514:  uint16(aux_sym___body_repeat2),
	515:  uint16(68),
	516:  uint16(6),
	517:  uint16(sym_xml_body),
	518:  uint16(sym_json_body),
	519:  uint16(sym_graphql_body),
	520:  uint16(sym__external_body),
	521:  uint16(sym_multipart_form_data),
	522:  uint16(sym_raw_body),
	523:  uint16(85),
	524:  uint16(7),
	525:  uint16(aux_sym_WORD_CHAR_token1),
	526:  uint16(aux_sym_PUNCTUATION_token1),
	527:  uint16(sym_method),
	528:  uint16(aux_sym_http_version_token1),
	529:  uint16(anon_sym_LBRACE_LBRACE),
	530:  uint16(anon_sym_LT),
	531:  uint16(anon_sym_AT2),
	532:  uint16(20),
	533:  uint16(29),
	534:  uint16(1),
	535:  uint16(aux_sym_WS_token1),
	536:  uint16(33),
	537:  uint16(1),
	538:  uint16(aux_sym_COMMENT_PREFIX_token1),
	539:  uint16(35),
	540:  uint16(1),
	541:  uint16(anon_sym_GT),
	542:  uint16(37),
	543:  uint16(1),
	544:  uint16(aux_sym_res_redirect_token1),
	545:  uint16(39),
	546:  uint16(1),
	547:  uint16(aux_sym_xml_body_token1),
	548:  uint16(41),
	549:  uint16(1),
	550:  uint16(aux_sym_json_body_token1),
	551:  uint16(43),
	552:  uint16(1),
	553:  uint16(aux_sym_graphql_data_token1),
	554:  uint16(45),
	555:  uint16(1),
	556:  uint16(anon_sym_LT2),
	557:  uint16(47),
	558:  uint16(1),
	559:  uint16(anon_sym_DASH_DASH),
	560:  uint16(49),
	561:  uint16(1),
	562:  uint16(aux_sym_raw_body_token1),
	563:  uint16(51),
	564:  uint16(1),
	565:  uint16(aux_sym__blank_line_token1),
	566:  uint16(93),
	567:  uint16(1),
	568:  uint16(aux_sym_NL_token1),
	569:  uint16(47),
	570:  uint16(1),
	571:  uint16(sym_graphql_data),
	572:  uint16(59),
	573:  uint16(1),
	574:  uint16(sym__var_comment),
	575:  uint16(205),
	576:  uint16(1),
	577:  uint16(sym_external_body),
	578:  uint16(89),
	579:  uint16(2),
	581:  uint16(aux_sym_request_separator_token1),
	582:  uint16(48),
	583:  uint16(2),
	584:  uint16(sym__blank_line),
	585:  uint16(aux_sym___body_repeat1),
	586:  uint16(24),
	587:  uint16(3),
	588:  uint16(sym_res_handler_script),
	589:  uint16(sym_res_redirect),
	590:  uint16(aux_sym___body_repeat2),
	591:  uint16(68),
	592:  uint16(6),
	593:  uint16(sym_xml_body),
	594:  uint16(sym_json_body),
	595:  uint16(sym_graphql_body),
	596:  uint16(sym__external_body),
	597:  uint16(sym_multipart_form_data),
	598:  uint16(sym_raw_body),
	599:  uint16(91),
	600:  uint16(7),
	601:  uint16(aux_sym_WORD_CHAR_token1),
	602:  uint16(aux_sym_PUNCTUATION_token1),
	603:  uint16(sym_method),
	604:  uint16(aux_sym_http_version_token1),
	605:  uint16(anon_sym_LBRACE_LBRACE),
	606:  uint16(anon_sym_LT),
	607:  uint16(anon_sym_AT2),
	608:  uint16(20),
	609:  uint16(29),
	610:  uint16(1),
	611:  uint16(aux_sym_WS_token1),
	612:  uint16(33),
	613:  uint16(1),
	614:  uint16(aux_sym_COMMENT_PREFIX_token1),
	615:  uint16(35),
	616:  uint16(1),
	617:  uint16(anon_sym_GT),
	618:  uint16(37),
	619:  uint16(1),
	620:  uint16(aux_sym_res_redirect_token1),
	621:  uint16(39),
	622:  uint16(1),
	623:  uint16(aux_sym_xml_body_token1),
	624:  uint16(41),
	625:  uint16(1),
	626:  uint16(aux_sym_json_body_token1),
	627:  uint16(43),
	628:  uint16(1),
	629:  uint16(aux_sym_graphql_data_token1),
	630:  uint16(45),
	631:  uint16(1),
	632:  uint16(anon_sym_LT2),
	633:  uint16(47),
	634:  uint16(1),
	635:  uint16(anon_sym_DASH_DASH),
	636:  uint16(49),
	637:  uint16(1),
	638:  uint16(aux_sym_raw_body_token1),
	639:  uint16(51),
	640:  uint16(1),
	641:  uint16(aux_sym__blank_line_token1),
	642:  uint16(99),
	643:  uint16(1),
	644:  uint16(aux_sym_NL_token1),
	645:  uint16(47),
	646:  uint16(1),
	647:  uint16(sym_graphql_data),
	648:  uint16(59),
	649:  uint16(1),
	650:  uint16(sym__var_comment),
	651:  uint16(205),
	652:  uint16(1),
	653:  uint16(sym_external_body),
	654:  uint16(95),
	655:  uint16(2),
	657:  uint16(aux_sym_request_separator_token1),
	658:  uint16(48),
	659:  uint16(2),
	660:  uint16(sym__blank_line),
	661:  uint16(aux_sym___body_repeat1),
	662:  uint16(23),
	663:  uint16(3),
	664:  uint16(sym_res_handler_script),
	665:  uint16(sym_res_redirect),
	666:  uint16(aux_sym___body_repeat2),
	667:  uint16(68),
	668:  uint16(6),
	669:  uint16(sym_xml_body),
	670:  uint16(sym_json_body),
	671:  uint16(sym_graphql_body),
	672:  uint16(sym__external_body),
	673:  uint16(sym_multipart_form_data),
	674:  uint16(sym_raw_body),
	675:  uint16(97),
	676:  uint16(7),
	677:  uint16(aux_sym_WORD_CHAR_token1),
	678:  uint16(aux_sym_PUNCTUATION_token1),
	679:  uint16(sym_method),
	680:  uint16(aux_sym_http_version_token1),
	681:  uint16(anon_sym_LBRACE_LBRACE),
	682:  uint16(anon_sym_LT),
	683:  uint16(anon_sym_AT2),
	684:  uint16(20),
	685:  uint16(29),
	686:  uint16(1),
	687:  uint16(aux_sym_WS_token1),
	688:  uint16(33),
	689:  uint16(1),
	690:  uint16(aux_sym_COMMENT_PREFIX_token1),
	691:  uint16(35),
	692:  uint16(1),
	693:  uint16(anon_sym_GT),
	694:  uint16(37),
	695:  uint16(1),
	696:  uint16(aux_sym_res_redirect_token1),
	697:  uint16(39),
	698:  uint16(1),
	699:  uint16(aux_sym_xml_body_token1),
	700:  uint16(41),
	701:  uint16(1),
	702:  uint16(aux_sym_json_body_token1),
	703:  uint16(43),
	704:  uint16(1),
	705:  uint16(aux_sym_graphql_data_token1),
	706:  uint16(45),
	707:  uint16(1),
	708:  uint16(anon_sym_LT2),
	709:  uint16(47),
	710:  uint16(1),
	711:  uint16(anon_sym_DASH_DASH),
	712:  uint16(49),
	713:  uint16(1),
	714:  uint16(aux_sym_raw_body_token1),
	715:  uint16(51),
	716:  uint16(1),
	717:  uint16(aux_sym__blank_line_token1),
	718:  uint16(105),
	719:  uint16(1),
	720:  uint16(aux_sym_NL_token1),
	721:  uint16(47),
	722:  uint16(1),
	723:  uint16(sym_graphql_data),
	724:  uint16(59),
	725:  uint16(1),
	726:  uint16(sym__var_comment),
	727:  uint16(205),
	728:  uint16(1),
	729:  uint16(sym_external_body),
	730:  uint16(101),
	731:  uint16(2),
	733:  uint16(aux_sym_request_separator_token1),
	734:  uint16(48),
	735:  uint16(2),
	736:  uint16(sym__blank_line),
	737:  uint16(aux_sym___body_repeat1),
	738:  uint16(22),
	739:  uint16(3),
	740:  uint16(sym_res_handler_script),
	741:  uint16(sym_res_redirect),
	742:  uint16(aux_sym___body_repeat2),
	743:  uint16(68),
	744:  uint16(6),
	745:  uint16(sym_xml_body),
	746:  uint16(sym_json_body),
	747:  uint16(sym_graphql_body),
	748:  uint16(sym__external_body),
	749:  uint16(sym_multipart_form_data),
	750:  uint16(sym_raw_body),
	751:  uint16(103),
	752:  uint16(7),
	753:  uint16(aux_sym_WORD_CHAR_token1),
	754:  uint16(aux_sym_PUNCTUATION_token1),
	755:  uint16(sym_method),
	756:  uint16(aux_sym_http_version_token1),
	757:  uint16(anon_sym_LBRACE_LBRACE),
	758:  uint16(anon_sym_LT),
	759:  uint16(anon_sym_AT2),
	760:  uint16(20),
	761:  uint16(29),
	762:  uint16(1),
	763:  uint16(aux_sym_WS_token1),
	764:  uint16(33),
	765:  uint16(1),
	766:  uint16(aux_sym_COMMENT_PREFIX_token1),
	767:  uint16(35),
	768:  uint16(1),
	769:  uint16(anon_sym_GT),
	770:  uint16(37),
	771:  uint16(1),
	772:  uint16(aux_sym_res_redirect_token1),
	773:  uint16(39),
	774:  uint16(1),
	775:  uint16(aux_sym_xml_body_token1),
	776:  uint16(41),
	777:  uint16(1),
	778:  uint16(aux_sym_json_body_token1),
	779:  uint16(43),
	780:  uint16(1),
	781:  uint16(aux_sym_graphql_data_token1),
	782:  uint16(45),
	783:  uint16(1),
	784:  uint16(anon_sym_LT2),
	785:  uint16(47),
	786:  uint16(1),
	787:  uint16(anon_sym_DASH_DASH),
	788:  uint16(49),
	789:  uint16(1),
	790:  uint16(aux_sym_raw_body_token1),
	791:  uint16(51),
	792:  uint16(1),
	793:  uint16(aux_sym__blank_line_token1),
	794:  uint16(111),
	795:  uint16(1),
	796:  uint16(aux_sym_NL_token1),
	797:  uint16(47),
	798:  uint16(1),
	799:  uint16(sym_graphql_data),
	800:  uint16(59),
	801:  uint16(1),
	802:  uint16(sym__var_comment),
	803:  uint16(205),
	804:  uint16(1),
	805:  uint16(sym_external_body),
	806:  uint16(107),
	807:  uint16(2),
	809:  uint16(aux_sym_request_separator_token1),
	810:  uint16(48),
	811:  uint16(2),
	812:  uint16(sym__blank_line),
	813:  uint16(aux_sym___body_repeat1),
	814:  uint16(21),
	815:  uint16(3),
	816:  uint16(sym_res_handler_script),
	817:  uint16(sym_res_redirect),
	818:  uint16(aux_sym___body_repeat2),
	819:  uint16(68),
	820:  uint16(6),
	821:  uint16(sym_xml_body),
	822:  uint16(sym_json_body),
	823:  uint16(sym_graphql_body),
	824:  uint16(sym__external_body),
	825:  uint16(sym_multipart_form_data),
	826:  uint16(sym_raw_body),
	827:  uint16(109),
	828:  uint16(7),
	829:  uint16(aux_sym_WORD_CHAR_token1),
	830:  uint16(aux_sym_PUNCTUATION_token1),
	831:  uint16(sym_method),
	832:  uint16(aux_sym_http_version_token1),
	833:  uint16(anon_sym_LBRACE_LBRACE),
	834:  uint16(anon_sym_LT),
	835:  uint16(anon_sym_AT2),
	836:  uint16(20),
	837:  uint16(29),
	838:  uint16(1),
	839:  uint16(aux_sym_WS_token1),
	840:  uint16(33),
	841:  uint16(1),
	842:  uint16(aux_sym_COMMENT_PREFIX_token1),
	843:  uint16(35),
	844:  uint16(1),
	845:  uint16(anon_sym_GT),
	846:  uint16(37),
	847:  uint16(1),
	848:  uint16(aux_sym_res_redirect_token1),
	849:  uint16(39),
	850:  uint16(1),
	851:  uint16(aux_sym_xml_body_token1),
	852:  uint16(41),
	853:  uint16(1),
	854:  uint16(aux_sym_json_body_token1),
	855:  uint16(43),
	856:  uint16(1),
	857:  uint16(aux_sym_graphql_data_token1),
	858:  uint16(45),
	859:  uint16(1),
	860:  uint16(anon_sym_LT2),
	861:  uint16(47),
	862:  uint16(1),
	863:  uint16(anon_sym_DASH_DASH),
	864:  uint16(49),
	865:  uint16(1),
	866:  uint16(aux_sym_raw_body_token1),
	867:  uint16(51),
	868:  uint16(1),
	869:  uint16(aux_sym__blank_line_token1),
	870:  uint16(117),
	871:  uint16(1),
	872:  uint16(aux_sym_NL_token1),
	873:  uint16(47),
	874:  uint16(1),
	875:  uint16(sym_graphql_data),
	876:  uint16(59),
	877:  uint16(1),
	878:  uint16(sym__var_comment),
	879:  uint16(205),
	880:  uint16(1),
	881:  uint16(sym_external_body),
	882:  uint16(113),
	883:  uint16(2),
	885:  uint16(aux_sym_request_separator_token1),
	886:  uint16(48),
	887:  uint16(2),
	888:  uint16(sym__blank_line),
	889:  uint16(aux_sym___body_repeat1),
	890:  uint16(20),
	891:  uint16(3),
	892:  uint16(sym_res_handler_script),
	893:  uint16(sym_res_redirect),
	894:  uint16(aux_sym___body_repeat2),
	895:  uint16(68),
	896:  uint16(6),
	897:  uint16(sym_xml_body),
	898:  uint16(sym_json_body),
	899:  uint16(sym_graphql_body),
	900:  uint16(sym__external_body),
	901:  uint16(sym_multipart_form_data),
	902:  uint16(sym_raw_body),
	903:  uint16(115),
	904:  uint16(7),
	905:  uint16(aux_sym_WORD_CHAR_token1),
	906:  uint16(aux_sym_PUNCTUATION_token1),
	907:  uint16(sym_method),
	908:  uint16(aux_sym_http_version_token1),
	909:  uint16(anon_sym_LBRACE_LBRACE),
	910:  uint16(anon_sym_LT),
	911:  uint16(anon_sym_AT2),
	912:  uint16(17),
	913:  uint16(33),
	914:  uint16(1),
	915:  uint16(aux_sym_COMMENT_PREFIX_token1),
	916:  uint16(35),
	917:  uint16(1),
	918:  uint16(anon_sym_GT),
	919:  uint16(37),
	920:  uint16(1),
	921:  uint16(aux_sym_res_redirect_token1),
	922:  uint16(39),
	923:  uint16(1),
	924:  uint16(aux_sym_xml_body_token1),
	925:  uint16(41),
	926:  uint16(1),
	927:  uint16(aux_sym_json_body_token1),
	928:  uint16(43),
	929:  uint16(1),
	930:  uint16(aux_sym_graphql_data_token1),
	931:  uint16(45),
	932:  uint16(1),
	933:  uint16(anon_sym_LT2),
	934:  uint16(47),
	935:  uint16(1),
	936:  uint16(anon_sym_DASH_DASH),
	937:  uint16(49),
	938:  uint16(1),
	939:  uint16(aux_sym_raw_body_token1),
	940:  uint16(123),
	941:  uint16(1),
	942:  uint16(aux_sym_NL_token1),
	943:  uint16(47),
	944:  uint16(1),
	945:  uint16(sym_graphql_data),
	946:  uint16(59),
	947:  uint16(1),
	948:  uint16(sym__var_comment),
	949:  uint16(205),
	950:  uint16(1),
	951:  uint16(sym_external_body),
	952:  uint16(119),
	953:  uint16(2),
	955:  uint16(aux_sym_request_separator_token1),
	956:  uint16(16),
	957:  uint16(3),
	958:  uint16(sym_res_handler_script),
	959:  uint16(sym_res_redirect),
	960:  uint16(aux_sym___body_repeat2),
	961:  uint16(68),
	962:  uint16(6),
	963:  uint16(sym_xml_body),
	964:  uint16(sym_json_body),
	965:  uint16(sym_graphql_body),
	966:  uint16(sym__external_body),
	967:  uint16(sym_multipart_form_data),
	968:  uint16(sym_raw_body),
	969:  uint16(121),
	970:  uint16(9),
	971:  uint16(aux_sym_WORD_CHAR_token1),
	972:  uint16(aux_sym_PUNCTUATION_token1),
	973:  uint16(aux_sym_WS_token1),
	974:  uint16(sym_method),
	975:  uint16(aux_sym_http_version_token1),
	976:  uint16(anon_sym_LBRACE_LBRACE),
	977:  uint16(anon_sym_LT),
	978:  uint16(anon_sym_AT2),
	979:  uint16(aux_sym__blank_line_token1),
	980:  uint16(17),
	981:  uint16(33),
	982:  uint16(1),
	983:  uint16(aux_sym_COMMENT_PREFIX_token1),
	984:  uint16(35),
	985:  uint16(1),
	986:  uint16(anon_sym_GT),
	987:  uint16(37),
	988:  uint16(1),
	989:  uint16(aux_sym_res_redirect_token1),
	990:  uint16(39),
	991:  uint16(1),
	992:  uint16(aux_sym_xml_body_token1),
	993:  uint16(41),
	994:  uint16(1),
	995:  uint16(aux_sym_json_body_token1),
	996:  uint16(43),
	997:  uint16(1),
	998:  uint16(aux_sym_graphql_data_token1),
	999:  uint16(45),
	1000: uint16(1),
	1001: uint16(anon_sym_LT2),
	1002: uint16(47),
	1003: uint16(1),
	1004: uint16(anon_sym_DASH_DASH),
	1005: uint16(49),
	1006: uint16(1),
	1007: uint16(aux_sym_raw_body_token1),
	1008: uint16(123),
	1009: uint16(1),
	1010: uint16(aux_sym_NL_token1),
	1011: uint16(47),
	1012: uint16(1),
	1013: uint16(sym_graphql_data),
	1014: uint16(59),
	1015: uint16(1),
	1016: uint16(sym__var_comment),
	1017: uint16(205),
	1018: uint16(1),
	1019: uint16(sym_external_body),
	1020: uint16(125),
	1021: uint16(2),
	1023: uint16(aux_sym_request_separator_token1),
	1024: uint16(16),
	1025: uint16(3),
	1026: uint16(sym_res_handler_script),
	1027: uint16(sym_res_redirect),
	1028: uint16(aux_sym___body_repeat2),
	1029: uint16(68),
	1030: uint16(6),
	1031: uint16(sym_xml_body),
	1032: uint16(sym_json_body),
	1033: uint16(sym_graphql_body),
	1034: uint16(sym__external_body),
	1035: uint16(sym_multipart_form_data),
	1036: uint16(sym_raw_body),
	1037: uint16(127),
	1038: uint16(9),
	1039: uint16(aux_sym_WORD_CHAR_token1),
	1040: uint16(aux_sym_PUNCTUATION_token1),
	1041: uint16(aux_sym_WS_token1),
	1042: uint16(sym_method),
	1043: uint16(aux_sym_http_version_token1),
	1044: uint16(anon_sym_LBRACE_LBRACE),
	1045: uint16(anon_sym_LT),
	1046: uint16(anon_sym_AT2),
	1047: uint16(aux_sym__blank_line_token1),
	1048: uint16(17),
	1049: uint16(133),
	1050: uint16(1),
	1051: uint16(aux_sym_NL_token1),
	1052: uint16(136),
	1053: uint16(1),
	1054: uint16(aux_sym_COMMENT_PREFIX_token1),
	1055: uint16(139),
	1056: uint16(1),
	1057: uint16(anon_sym_GT),
	1058: uint16(142),
	1059: uint16(1),
	1060: uint16(aux_sym_res_redirect_token1),
	1061: uint16(145),
	1062: uint16(1),
	1063: uint16(aux_sym_xml_body_token1),
	1064: uint16(148),
	1065: uint16(1),
	1066: uint16(aux_sym_json_body_token1),
	1067: uint16(151),
	1068: uint16(1),
	1069: uint16(aux_sym_graphql_data_token1),
	1070: uint16(154),
	1071: uint16(1),
	1072: uint16(anon_sym_LT2),
	1073: uint16(157),
	1074: uint16(1),
	1075: uint16(anon_sym_DASH_DASH),
	1076: uint16(160),
	1077: uint16(1),
	1078: uint16(aux_sym_raw_body_token1),
	1079: uint16(47),
	1080: uint16(1),
	1081: uint16(sym_graphql_data),
	1082: uint16(59),
	1083: uint16(1),
	1084: uint16(sym__var_comment),
	1085: uint16(205),
	1086: uint16(1),
	1087: uint16(sym_external_body),
	1088: uint16(129),
	1089: uint16(2),
	1091: uint16(aux_sym_request_separator_token1),
	1092: uint16(16),
	1093: uint16(3),
	1094: uint16(sym_res_handler_script),
	1095: uint16(sym_res_redirect),
	1096: uint16(aux_sym___body_repeat2),
	1097: uint16(68),
	1098: uint16(6),
	1099: uint16(sym_xml_body),
	1100: uint16(sym_json_body),
	1101: uint16(sym_graphql_body),
	1102: uint16(sym__external_body),
	1103: uint16(sym_multipart_form_data),
	1104: uint16(sym_raw_body),
	1105: uint16(131),
	1106: uint16(9),
	1107: uint16(aux_sym_WORD_CHAR_token1),
	1108: uint16(aux_sym_PUNCTUATION_token1),
	1109: uint16(aux_sym_WS_token1),
	1110: uint16(sym_method),
	1111: uint16(aux_sym_http_version_token1),
	1112: uint16(anon_sym_LBRACE_LBRACE),
	1113: uint16(anon_sym_LT),
	1114: uint16(anon_sym_AT2),
	1115: uint16(aux_sym__blank_line_token1),
	1116: uint16(17),
	1117: uint16(33),
	1118: uint16(1),
	1119: uint16(aux_sym_COMMENT_PREFIX_token1),
	1120: uint16(35),
	1121: uint16(1),
	1122: uint16(anon_sym_GT),
	1123: uint16(37),
	1124: uint16(1),
	1125: uint16(aux_sym_res_redirect_token1),
	1126: uint16(39),
	1127: uint16(1),
	1128: uint16(aux_sym_xml_body_token1),
	1129: uint16(41),
	1130: uint16(1),
	1131: uint16(aux_sym_json_body_token1),
	1132: uint16(43),
	1133: uint16(1),
	1134: uint16(aux_sym_graphql_data_token1),
	1135: uint16(45),
	1136: uint16(1),
	1137: uint16(anon_sym_LT2),
	1138: uint16(47),
	1139: uint16(1),
	1140: uint16(anon_sym_DASH_DASH),
	1141: uint16(49),
	1142: uint16(1),
	1143: uint16(aux_sym_raw_body_token1),
	1144: uint16(123),
	1145: uint16(1),
	1146: uint16(aux_sym_NL_token1),
	1147: uint16(47),
	1148: uint16(1),
	1149: uint16(sym_graphql_data),
	1150: uint16(59),
	1151: uint16(1),
	1152: uint16(sym__var_comment),
	1153: uint16(205),
	1154: uint16(1),
	1155: uint16(sym_external_body),
	1156: uint16(163),
	1157: uint16(2),
	1159: uint16(aux_sym_request_separator_token1),
	1160: uint16(16),
	1161: uint16(3),
	1162: uint16(sym_res_handler_script),
	1163: uint16(sym_res_redirect),
	1164: uint16(aux_sym___body_repeat2),
	1165: uint16(68),
	1166: uint16(6),
	1167: uint16(sym_xml_body),
	1168: uint16(sym_json_body),
	1169: uint16(sym_graphql_body),
	1170: uint16(sym__external_body),
	1171: uint16(sym_multipart_form_data),
	1172: uint16(sym_raw_body),
	1173: uint16(165),
	1174: uint16(9),
	1175: uint16(aux_sym_WORD_CHAR_token1),
	1176: uint16(aux_sym_PUNCTUATION_token1),
	1177: uint16(aux_sym_WS_token1),
	1178: uint16(sym_method),
	1179: uint16(aux_sym_http_version_token1),
	1180: uint16(anon_sym_LBRACE_LBRACE),
	1181: uint16(anon_sym_LT),
	1182: uint16(anon_sym_AT2),
	1183: uint16(aux_sym__blank_line_token1),
	1184: uint16(17),
	1185: uint16(33),
	1186: uint16(1),
	1187: uint16(aux_sym_COMMENT_PREFIX_token1),
	1188: uint16(35),
	1189: uint16(1),
	1190: uint16(anon_sym_GT),
	1191: uint16(37),
	1192: uint16(1),
	1193: uint16(aux_sym_res_redirect_token1),
	1194: uint16(39),
	1195: uint16(1),
	1196: uint16(aux_sym_xml_body_token1),
	1197: uint16(41),
	1198: uint16(1),
	1199: uint16(aux_sym_json_body_token1),
	1200: uint16(43),
	1201: uint16(1),
	1202: uint16(aux_sym_graphql_data_token1),
	1203: uint16(45),
	1204: uint16(1),
	1205: uint16(anon_sym_LT2),
	1206: uint16(47),
	1207: uint16(1),
	1208: uint16(anon_sym_DASH_DASH),
	1209: uint16(49),
	1210: uint16(1),
	1211: uint16(aux_sym_raw_body_token1),
	1212: uint16(123),
	1213: uint16(1),
	1214: uint16(aux_sym_NL_token1),
	1215: uint16(47),
	1216: uint16(1),
	1217: uint16(sym_graphql_data),
	1218: uint16(59),
	1219: uint16(1),
	1220: uint16(sym__var_comment),
	1221: uint16(205),
	1222: uint16(1),
	1223: uint16(sym_external_body),
	1224: uint16(167),
	1225: uint16(2),
	1227: uint16(aux_sym_request_separator_token1),
	1228: uint16(16),
	1229: uint16(3),
	1230: uint16(sym_res_handler_script),
	1231: uint16(sym_res_redirect),
	1232: uint16(aux_sym___body_repeat2),
	1233: uint16(68),
	1234: uint16(6),
	1235: uint16(sym_xml_body),
	1236: uint16(sym_json_body),
	1237: uint16(sym_graphql_body),
	1238: uint16(sym__external_body),
	1239: uint16(sym_multipart_form_data),
	1240: uint16(sym_raw_body),
	1241: uint16(169),
	1242: uint16(9),
	1243: uint16(aux_sym_WORD_CHAR_token1),
	1244: uint16(aux_sym_PUNCTUATION_token1),
	1245: uint16(aux_sym_WS_token1),
	1246: uint16(sym_method),
	1247: uint16(aux_sym_http_version_token1),
	1248: uint16(anon_sym_LBRACE_LBRACE),
	1249: uint16(anon_sym_LT),
	1250: uint16(anon_sym_AT2),
	1251: uint16(aux_sym__blank_line_token1),
	1252: uint16(17),
	1253: uint16(33),
	1254: uint16(1),
	1255: uint16(aux_sym_COMMENT_PREFIX_token1),
	1256: uint16(35),
	1257: uint16(1),
	1258: uint16(anon_sym_GT),
	1259: uint16(37),
	1260: uint16(1),
	1261: uint16(aux_sym_res_redirect_token1),
	1262: uint16(39),
	1263: uint16(1),
	1264: uint16(aux_sym_xml_body_token1),
	1265: uint16(41),
	1266: uint16(1),
	1267: uint16(aux_sym_json_body_token1),
	1268: uint16(43),
	1269: uint16(1),
	1270: uint16(aux_sym_graphql_data_token1),
	1271: uint16(45),
	1272: uint16(1),
	1273: uint16(anon_sym_LT2),
	1274: uint16(47),
	1275: uint16(1),
	1276: uint16(anon_sym_DASH_DASH),
	1277: uint16(49),
	1278: uint16(1),
	1279: uint16(aux_sym_raw_body_token1),
	1280: uint16(123),
	1281: uint16(1),
	1282: uint16(aux_sym_NL_token1),
	1283: uint16(47),
	1284: uint16(1),
	1285: uint16(sym_graphql_data),
	1286: uint16(59),
	1287: uint16(1),
	1288: uint16(sym__var_comment),
	1289: uint16(205),
	1290: uint16(1),
	1291: uint16(sym_external_body),
	1292: uint16(171),
	1293: uint16(2),
	1295: uint16(aux_sym_request_separator_token1),
	1296: uint16(16),
	1297: uint16(3),
	1298: uint16(sym_res_handler_script),
	1299: uint16(sym_res_redirect),
	1300: uint16(aux_sym___body_repeat2),
	1301: uint16(68),
	1302: uint16(6),
	1303: uint16(sym_xml_body),
	1304: uint16(sym_json_body),
	1305: uint16(sym_graphql_body),
	1306: uint16(sym__external_body),
	1307: uint16(sym_multipart_form_data),
	1308: uint16(sym_raw_body),
	1309: uint16(173),
	1310: uint16(9),
	1311: uint16(aux_sym_WORD_CHAR_token1),
	1312: uint16(aux_sym_PUNCTUATION_token1),
	1313: uint16(aux_sym_WS_token1),
	1314: uint16(sym_method),
	1315: uint16(aux_sym_http_version_token1),
	1316: uint16(anon_sym_LBRACE_LBRACE),
	1317: uint16(anon_sym_LT),
	1318: uint16(anon_sym_AT2),
	1319: uint16(aux_sym__blank_line_token1),
	1320: uint16(17),
	1321: uint16(33),
	1322: uint16(1),
	1323: uint16(aux_sym_COMMENT_PREFIX_token1),
	1324: uint16(35),
	1325: uint16(1),
	1326: uint16(anon_sym_GT),
	1327: uint16(37),
	1328: uint16(1),
	1329: uint16(aux_sym_res_redirect_token1),
	1330: uint16(39),
	1331: uint16(1),
	1332: uint16(aux_sym_xml_body_token1),
	1333: uint16(41),
	1334: uint16(1),
	1335: uint16(aux_sym_json_body_token1),
	1336: uint16(43),
	1337: uint16(1),
	1338: uint16(aux_sym_graphql_data_token1),
	1339: uint16(45),
	1340: uint16(1),
	1341: uint16(anon_sym_LT2),
	1342: uint16(47),
	1343: uint16(1),
	1344: uint16(anon_sym_DASH_DASH),
	1345: uint16(49),
	1346: uint16(1),
	1347: uint16(aux_sym_raw_body_token1),
	1348: uint16(123),
	1349: uint16(1),
	1350: uint16(aux_sym_NL_token1),
	1351: uint16(47),
	1352: uint16(1),
	1353: uint16(sym_graphql_data),
	1354: uint16(59),
	1355: uint16(1),
	1356: uint16(sym__var_comment),
	1357: uint16(205),
	1358: uint16(1),
	1359: uint16(sym_external_body),
	1360: uint16(175),
	1361: uint16(2),
	1363: uint16(aux_sym_request_separator_token1),
	1364: uint16(16),
	1365: uint16(3),
	1366: uint16(sym_res_handler_script),
	1367: uint16(sym_res_redirect),
	1368: uint16(aux_sym___body_repeat2),
	1369: uint16(68),
	1370: uint16(6),
	1371: uint16(sym_xml_body),
	1372: uint16(sym_json_body),
	1373: uint16(sym_graphql_body),
	1374: uint16(sym__external_body),
	1375: uint16(sym_multipart_form_data),
	1376: uint16(sym_raw_body),
	1377: uint16(177),
	1378: uint16(9),
	1379: uint16(aux_sym_WORD_CHAR_token1),
	1380: uint16(aux_sym_PUNCTUATION_token1),
	1381: uint16(aux_sym_WS_token1),
	1382: uint16(sym_method),
	1383: uint16(aux_sym_http_version_token1),
	1384: uint16(anon_sym_LBRACE_LBRACE),
	1385: uint16(anon_sym_LT),
	1386: uint16(anon_sym_AT2),
	1387: uint16(aux_sym__blank_line_token1),
	1388: uint16(17),
	1389: uint16(33),
	1390: uint16(1),
	1391: uint16(aux_sym_COMMENT_PREFIX_token1),
	1392: uint16(35),
	1393: uint16(1),
	1394: uint16(anon_sym_GT),
	1395: uint16(37),
	1396: uint16(1),
	1397: uint16(aux_sym_res_redirect_token1),
	1398: uint16(39),
	1399: uint16(1),
	1400: uint16(aux_sym_xml_body_token1),
	1401: uint16(41),
	1402: uint16(1),
	1403: uint16(aux_sym_json_body_token1),
	1404: uint16(43),
	1405: uint16(1),
	1406: uint16(aux_sym_graphql_data_token1),
	1407: uint16(45),
	1408: uint16(1),
	1409: uint16(anon_sym_LT2),
	1410: uint16(47),
	1411: uint16(1),
	1412: uint16(anon_sym_DASH_DASH),
	1413: uint16(49),
	1414: uint16(1),
	1415: uint16(aux_sym_raw_body_token1),
	1416: uint16(123),
	1417: uint16(1),
	1418: uint16(aux_sym_NL_token1),
	1419: uint16(47),
	1420: uint16(1),
	1421: uint16(sym_graphql_data),
	1422: uint16(59),
	1423: uint16(1),
	1424: uint16(sym__var_comment),
	1425: uint16(205),
	1426: uint16(1),
	1427: uint16(sym_external_body),
	1428: uint16(179),
	1429: uint16(2),
	1431: uint16(aux_sym_request_separator_token1),
	1432: uint16(16),
	1433: uint16(3),
	1434: uint16(sym_res_handler_script),
	1435: uint16(sym_res_redirect),
	1436: uint16(aux_sym___body_repeat2),
	1437: uint16(68),
	1438: uint16(6),
	1439: uint16(sym_xml_body),
	1440: uint16(sym_json_body),
	1441: uint16(sym_graphql_body),
	1442: uint16(sym__external_body),
	1443: uint16(sym_multipart_form_data),
	1444: uint16(sym_raw_body),
	1445: uint16(181),
	1446: uint16(9),
	1447: uint16(aux_sym_WORD_CHAR_token1),
	1448: uint16(aux_sym_PUNCTUATION_token1),
	1449: uint16(aux_sym_WS_token1),
	1450: uint16(sym_method),
	1451: uint16(aux_sym_http_version_token1),
	1452: uint16(anon_sym_LBRACE_LBRACE),
	1453: uint16(anon_sym_LT),
	1454: uint16(anon_sym_AT2),
	1455: uint16(aux_sym__blank_line_token1),
	1456: uint16(17),
	1457: uint16(33),
	1458: uint16(1),
	1459: uint16(aux_sym_COMMENT_PREFIX_token1),
	1460: uint16(35),
	1461: uint16(1),
	1462: uint16(anon_sym_GT),
	1463: uint16(37),
	1464: uint16(1),
	1465: uint16(aux_sym_res_redirect_token1),
	1466: uint16(39),
	1467: uint16(1),
	1468: uint16(aux_sym_xml_body_token1),
	1469: uint16(41),
	1470: uint16(1),
	1471: uint16(aux_sym_json_body_token1),
	1472: uint16(43),
	1473: uint16(1),
	1474: uint16(aux_sym_graphql_data_token1),
	1475: uint16(45),
	1476: uint16(1),
	1477: uint16(anon_sym_LT2),
	1478: uint16(47),
	1479: uint16(1),
	1480: uint16(anon_sym_DASH_DASH),
	1481: uint16(49),
	1482: uint16(1),
	1483: uint16(aux_sym_raw_body_token1),
	1484: uint16(123),
	1485: uint16(1),
	1486: uint16(aux_sym_NL_token1),
	1487: uint16(47),
	1488: uint16(1),
	1489: uint16(sym_graphql_data),
	1490: uint16(59),
	1491: uint16(1),
	1492: uint16(sym__var_comment),
	1493: uint16(205),
	1494: uint16(1),
	1495: uint16(sym_external_body),
	1496: uint16(183),
	1497: uint16(2),
	1499: uint16(aux_sym_request_separator_token1),
	1500: uint16(16),
	1501: uint16(3),
	1502: uint16(sym_res_handler_script),
	1503: uint16(sym_res_redirect),
	1504: uint16(aux_sym___body_repeat2),
	1505: uint16(68),
	1506: uint16(6),
	1507: uint16(sym_xml_body),
	1508: uint16(sym_json_body),
	1509: uint16(sym_graphql_body),
	1510: uint16(sym__external_body),
	1511: uint16(sym_multipart_form_data),
	1512: uint16(sym_raw_body),
	1513: uint16(185),
	1514: uint16(9),
	1515: uint16(aux_sym_WORD_CHAR_token1),
	1516: uint16(aux_sym_PUNCTUATION_token1),
	1517: uint16(aux_sym_WS_token1),
	1518: uint16(sym_method),
	1519: uint16(aux_sym_http_version_token1),
	1520: uint16(anon_sym_LBRACE_LBRACE),
	1521: uint16(anon_sym_LT),
	1522: uint16(anon_sym_AT2),
	1523: uint16(aux_sym__blank_line_token1),
	1524: uint16(17),
	1525: uint16(33),
	1526: uint16(1),
	1527: uint16(aux_sym_COMMENT_PREFIX_token1),
	1528: uint16(35),
	1529: uint16(1),
	1530: uint16(anon_sym_GT),
	1531: uint16(37),
	1532: uint16(1),
	1533: uint16(aux_sym_res_redirect_token1),
	1534: uint16(39),
	1535: uint16(1),
	1536: uint16(aux_sym_xml_body_token1),
	1537: uint16(41),
	1538: uint16(1),
	1539: uint16(aux_sym_json_body_token1),
	1540: uint16(43),
	1541: uint16(1),
	1542: uint16(aux_sym_graphql_data_token1),
	1543: uint16(45),
	1544: uint16(1),
	1545: uint16(anon_sym_LT2),
	1546: uint16(47),
	1547: uint16(1),
	1548: uint16(anon_sym_DASH_DASH),
	1549: uint16(49),
	1550: uint16(1),
	1551: uint16(aux_sym_raw_body_token1),
	1552: uint16(123),
	1553: uint16(1),
	1554: uint16(aux_sym_NL_token1),
	1555: uint16(47),
	1556: uint16(1),
	1557: uint16(sym_graphql_data),
	1558: uint16(59),
	1559: uint16(1),
	1560: uint16(sym__var_comment),
	1561: uint16(205),
	1562: uint16(1),
	1563: uint16(sym_external_body),
	1564: uint16(187),
	1565: uint16(2),
	1567: uint16(aux_sym_request_separator_token1),
	1568: uint16(16),
	1569: uint16(3),
	1570: uint16(sym_res_handler_script),
	1571: uint16(sym_res_redirect),
	1572: uint16(aux_sym___body_repeat2),
	1573: uint16(68),
	1574: uint16(6),
	1575: uint16(sym_xml_body),
	1576: uint16(sym_json_body),
	1577: uint16(sym_graphql_body),
	1578: uint16(sym__external_body),
	1579: uint16(sym_multipart_form_data),
	1580: uint16(sym_raw_body),
	1581: uint16(189),
	1582: uint16(9),
	1583: uint16(aux_sym_WORD_CHAR_token1),
	1584: uint16(aux_sym_PUNCTUATION_token1),
	1585: uint16(aux_sym_WS_token1),
	1586: uint16(sym_method),
	1587: uint16(aux_sym_http_version_token1),
	1588: uint16(anon_sym_LBRACE_LBRACE),
	1589: uint16(anon_sym_LT),
	1590: uint16(anon_sym_AT2),
	1591: uint16(aux_sym__blank_line_token1),
	1592: uint16(17),
	1593: uint16(33),
	1594: uint16(1),
	1595: uint16(aux_sym_COMMENT_PREFIX_token1),
	1596: uint16(35),
	1597: uint16(1),
	1598: uint16(anon_sym_GT),
	1599: uint16(37),
	1600: uint16(1),
	1601: uint16(aux_sym_res_redirect_token1),
	1602: uint16(39),
	1603: uint16(1),
	1604: uint16(aux_sym_xml_body_token1),
	1605: uint16(41),
	1606: uint16(1),
	1607: uint16(aux_sym_json_body_token1),
	1608: uint16(43),
	1609: uint16(1),
	1610: uint16(aux_sym_graphql_data_token1),
	1611: uint16(45),
	1612: uint16(1),
	1613: uint16(anon_sym_LT2),
	1614: uint16(47),
	1615: uint16(1),
	1616: uint16(anon_sym_DASH_DASH),
	1617: uint16(49),
	1618: uint16(1),
	1619: uint16(aux_sym_raw_body_token1),
	1620: uint16(123),
	1621: uint16(1),
	1622: uint16(aux_sym_NL_token1),
	1623: uint16(47),
	1624: uint16(1),
	1625: uint16(sym_graphql_data),
	1626: uint16(59),
	1627: uint16(1),
	1628: uint16(sym__var_comment),
	1629: uint16(205),
	1630: uint16(1),
	1631: uint16(sym_external_body),
	1632: uint16(191),
	1633: uint16(2),
	1635: uint16(aux_sym_request_separator_token1),
	1636: uint16(16),
	1637: uint16(3),
	1638: uint16(sym_res_handler_script),
	1639: uint16(sym_res_redirect),
	1640: uint16(aux_sym___body_repeat2),
	1641: uint16(68),
	1642: uint16(6),
	1643: uint16(sym_xml_body),
	1644: uint16(sym_json_body),
	1645: uint16(sym_graphql_body),
	1646: uint16(sym__external_body),
	1647: uint16(sym_multipart_form_data),
	1648: uint16(sym_raw_body),
	1649: uint16(193),
	1650: uint16(9),
	1651: uint16(aux_sym_WORD_CHAR_token1),
	1652: uint16(aux_sym_PUNCTUATION_token1),
	1653: uint16(aux_sym_WS_token1),
	1654: uint16(sym_method),
	1655: uint16(aux_sym_http_version_token1),
	1656: uint16(anon_sym_LBRACE_LBRACE),
	1657: uint16(anon_sym_LT),
	1658: uint16(anon_sym_AT2),
	1659: uint16(aux_sym__blank_line_token1),
	1660: uint16(17),
	1661: uint16(33),
	1662: uint16(1),
	1663: uint16(aux_sym_COMMENT_PREFIX_token1),
	1664: uint16(35),
	1665: uint16(1),
	1666: uint16(anon_sym_GT),
	1667: uint16(37),
	1668: uint16(1),
	1669: uint16(aux_sym_res_redirect_token1),
	1670: uint16(39),
	1671: uint16(1),
	1672: uint16(aux_sym_xml_body_token1),
	1673: uint16(41),
	1674: uint16(1),
	1675: uint16(aux_sym_json_body_token1),
	1676: uint16(43),
	1677: uint16(1),
	1678: uint16(aux_sym_graphql_data_token1),
	1679: uint16(45),
	1680: uint16(1),
	1681: uint16(anon_sym_LT2),
	1682: uint16(47),
	1683: uint16(1),
	1684: uint16(anon_sym_DASH_DASH),
	1685: uint16(49),
	1686: uint16(1),
	1687: uint16(aux_sym_raw_body_token1),
	1688: uint16(123),
	1689: uint16(1),
	1690: uint16(aux_sym_NL_token1),
	1691: uint16(47),
	1692: uint16(1),
	1693: uint16(sym_graphql_data),
	1694: uint16(59),
	1695: uint16(1),
	1696: uint16(sym__var_comment),
	1697: uint16(205),
	1698: uint16(1),
	1699: uint16(sym_external_body),
	1700: uint16(195),
	1701: uint16(2),
	1703: uint16(aux_sym_request_separator_token1),
	1704: uint16(16),
	1705: uint16(3),
	1706: uint16(sym_res_handler_script),
	1707: uint16(sym_res_redirect),
	1708: uint16(aux_sym___body_repeat2),
	1709: uint16(68),
	1710: uint16(6),
	1711: uint16(sym_xml_body),
	1712: uint16(sym_json_body),
	1713: uint16(sym_graphql_body),
	1714: uint16(sym__external_body),
	1715: uint16(sym_multipart_form_data),
	1716: uint16(sym_raw_body),
	1717: uint16(197),
	1718: uint16(9),
	1719: uint16(aux_sym_WORD_CHAR_token1),
	1720: uint16(aux_sym_PUNCTUATION_token1),
	1721: uint16(aux_sym_WS_token1),
	1722: uint16(sym_method),
	1723: uint16(aux_sym_http_version_token1),
	1724: uint16(anon_sym_LBRACE_LBRACE),
	1725: uint16(anon_sym_LT),
	1726: uint16(anon_sym_AT2),
	1727: uint16(aux_sym__blank_line_token1),
	1728: uint16(17),
	1729: uint16(33),
	1730: uint16(1),
	1731: uint16(aux_sym_COMMENT_PREFIX_token1),
	1732: uint16(35),
	1733: uint16(1),
	1734: uint16(anon_sym_GT),
	1735: uint16(37),
	1736: uint16(1),
	1737: uint16(aux_sym_res_redirect_token1),
	1738: uint16(39),
	1739: uint16(1),
	1740: uint16(aux_sym_xml_body_token1),
	1741: uint16(41),
	1742: uint16(1),
	1743: uint16(aux_sym_json_body_token1),
	1744: uint16(43),
	1745: uint16(1),
	1746: uint16(aux_sym_graphql_data_token1),
	1747: uint16(45),
	1748: uint16(1),
	1749: uint16(anon_sym_LT2),
	1750: uint16(47),
	1751: uint16(1),
	1752: uint16(anon_sym_DASH_DASH),
	1753: uint16(49),
	1754: uint16(1),
	1755: uint16(aux_sym_raw_body_token1),
	1756: uint16(123),
	1757: uint16(1),
	1758: uint16(aux_sym_NL_token1),
	1759: uint16(47),
	1760: uint16(1),
	1761: uint16(sym_graphql_data),
	1762: uint16(59),
	1763: uint16(1),
	1764: uint16(sym__var_comment),
	1765: uint16(205),
	1766: uint16(1),
	1767: uint16(sym_external_body),
	1768: uint16(199),
	1769: uint16(2),
	1771: uint16(aux_sym_request_separator_token1),
	1772: uint16(16),
	1773: uint16(3),
	1774: uint16(sym_res_handler_script),
	1775: uint16(sym_res_redirect),
	1776: uint16(aux_sym___body_repeat2),
	1777: uint16(68),
	1778: uint16(6),
	1779: uint16(sym_xml_body),
	1780: uint16(sym_json_body),
	1781: uint16(sym_graphql_body),
	1782: uint16(sym__external_body),
	1783: uint16(sym_multipart_form_data),
	1784: uint16(sym_raw_body),
	1785: uint16(201),
	1786: uint16(9),
	1787: uint16(aux_sym_WORD_CHAR_token1),
	1788: uint16(aux_sym_PUNCTUATION_token1),
	1789: uint16(aux_sym_WS_token1),
	1790: uint16(sym_method),
	1791: uint16(aux_sym_http_version_token1),
	1792: uint16(anon_sym_LBRACE_LBRACE),
	1793: uint16(anon_sym_LT),
	1794: uint16(anon_sym_AT2),
	1795: uint16(aux_sym__blank_line_token1),
	1796: uint16(22),
	1797: uint16(203),
	1798: uint16(1),
	1800: uint16(208),
	1801: uint16(1),
	1802: uint16(aux_sym_WS_token1),
	1803: uint16(211),
	1804: uint16(1),
	1805: uint16(aux_sym_COMMENT_PREFIX_token1),
	1806: uint16(214),
	1807: uint16(1),
	1808: uint16(aux_sym_request_separator_token1),
	1809: uint16(217),
	1810: uint16(1),
	1811: uint16(sym_method),
	1812: uint16(220),
	1813: uint16(1),
	1814: uint16(aux_sym_http_version_token1),
	1815: uint16(223),
	1816: uint16(1),
	1817: uint16(anon_sym_LBRACE_LBRACE),
	1818: uint16(226),
	1819: uint16(1),
	1820: uint16(anon_sym_LT),
	1821: uint16(229),
	1822: uint16(1),
	1823: uint16(anon_sym_AT2),
	1824: uint16(232),
	1825: uint16(1),
	1826: uint16(aux_sym__blank_line_token1),
	1827: uint16(33),
	1828: uint16(1),
	1829: uint16(sym_request_separator),
	1830: uint16(99),
	1831: uint16(1),
	1832: uint16(sym__section_content),
	1833: uint16(104),
	1834: uint16(1),
	1835: uint16(sym_response),
	1836: uint16(105),
	1837: uint16(1),
	1838: uint16(sym_request),
	1839: uint16(112),
	1840: uint16(1),
	1841: uint16(sym__plain_comment),
	1842: uint16(116),
	1843: uint16(1),
	1844: uint16(sym__var_comment),
	1845: uint16(200),
	1846: uint16(1),
	1847: uint16(sym_target_url),
	1848: uint16(214),
	1849: uint16(1),
	1850: uint16(sym_http_version),
	1851: uint16(205),
	1852: uint16(2),
	1853: uint16(aux_sym_WORD_CHAR_token1),
	1854: uint16(aux_sym_PUNCTUATION_token1),
	1855: uint16(27),
	1856: uint16(2),
	1857: uint16(sym_section),
	1858: uint16(aux_sym_document_repeat1),
	1859: uint16(132),
	1860: uint16(2),
	1861: uint16(aux_sym__target_url_line),
	1862: uint16(sym_variable),
	1863: uint16(32),
	1864: uint16(4),
	1865: uint16(sym_comment),
	1866: uint16(sym_pre_request_script),
	1867: uint16(sym_variable_declaration),
	1868: uint16(sym__blank_line),
	1869: uint16(10),
	1870: uint16(45),
	1871: uint16(1),
	1872: uint16(anon_sym_LT2),
	1873: uint16(239),
	1874: uint16(1),
	1875: uint16(aux_sym_COMMENT_PREFIX_token1),
	1876: uint16(243),
	1877: uint16(1),
	1878: uint16(aux_sym_multipart_form_data_token1),
	1879: uint16(38),
	1880: uint16(1),
	1881: uint16(sym__plain_comment),
	1882: uint16(39),
	1883: uint16(1),
	1884: uint16(sym__var_comment),
	1885: uint16(198),
	1886: uint16(1),
	1887: uint16(sym_external_body),
	1888: uint16(30),
	1889: uint16(2),
	1890: uint16(sym_comment),
	1891: uint16(aux_sym_multipart_form_data_repeat1),
	1892: uint16(235),
	1893: uint16(4),
	1895: uint16(aux_sym_request_separator_token1),
	1896: uint16(aux_sym_res_redirect_token1),
	1897: uint16(aux_sym_graphql_data_token1),
	1898: uint16(241),
	1899: uint16(4),
	1900: uint16(aux_sym_xml_body_token1),
	1901: uint16(aux_sym_json_body_token1),
	1902: uint16(anon_sym_DASH_DASH),
	1903: uint16(aux_sym_multipart_form_data_token2),
	1904: uint16(237),
	1905: uint16(12),
	1906: uint16(aux_sym_WORD_CHAR_token1),
	1907: uint16(aux_sym_PUNCTUATION_token1),
	1908: uint16(aux_sym_WS_token1),
	1909: uint16(aux_sym_NL_token1),
	1910: uint16(sym_method),
	1911: uint16(aux_sym_http_version_token1),
	1912: uint16(anon_sym_LBRACE_LBRACE),
	1913: uint16(anon_sym_LT),
	1914: uint16(anon_sym_GT),
	1915: uint16(anon_sym_AT2),
	1916: uint16(aux_sym_raw_body_token1),
	1917: uint16(aux_sym__blank_line_token1),
	1918: uint16(22),
	1919: uint16(7),
	1920: uint16(1),
	1921: uint16(aux_sym_WS_token1),
	1922: uint16(9),
	1923: uint16(1),
	1924: uint16(aux_sym_COMMENT_PREFIX_token1),
	1925: uint16(11),
	1926: uint16(1),
	1927: uint16(aux_sym_request_separator_token1),
	1928: uint16(13),
	1929: uint16(1),
	1930: uint16(sym_method),
	1931: uint16(15),
	1932: uint16(1),
	1933: uint16(aux_sym_http_version_token1),
	1934: uint16(17),
	1935: uint16(1),
	1936: uint16(anon_sym_LBRACE_LBRACE),
	1937: uint16(19),
	1938: uint16(1),
	1939: uint16(anon_sym_LT),
	1940: uint16(21),
	1941: uint16(1),
	1942: uint16(anon_sym_AT2),
	1943: uint16(23),
	1944: uint16(1),
	1945: uint16(aux_sym__blank_line_token1),
	1946: uint16(245),
	1947: uint16(1),
	1949: uint16(33),
	1950: uint16(1),
	1951: uint16(sym_request_separator),
	1952: uint16(99),
	1953: uint16(1),
	1954: uint16(sym__section_content),
	1955: uint16(104),
	1956: uint16(1),
	1957: uint16(sym_response),
	1958: uint16(105),
	1959: uint16(1),
	1960: uint16(sym_request),
	1961: uint16(112),
	1962: uint16(1),
	1963: uint16(sym__plain_comment),
	1964: uint16(116),
	1965: uint16(1),
	1966: uint16(sym__var_comment),
	1967: uint16(200),
	1968: uint16(1),
	1969: uint16(sym_target_url),
	1970: uint16(214),
	1971: uint16(1),
	1972: uint16(sym_http_version),
	1973: uint16(5),
	1974: uint16(2),
	1975: uint16(aux_sym_WORD_CHAR_token1),
	1976: uint16(aux_sym_PUNCTUATION_token1),
	1977: uint16(27),
	1978: uint16(2),
	1979: uint16(sym_section),
	1980: uint16(aux_sym_document_repeat1),
	1981: uint16(132),
	1982: uint16(2),
	1983: uint16(aux_sym__target_url_line),
	1984: uint16(sym_variable),
	1985: uint16(32),
	1986: uint16(4),
	1987: uint16(sym_comment),
	1988: uint16(sym_pre_request_script),
	1989: uint16(sym_variable_declaration),
	1990: uint16(sym__blank_line),
	1991: uint16(10),
	1992: uint16(45),
	1993: uint16(1),
	1994: uint16(anon_sym_LT2),
	1995: uint16(239),
	1996: uint16(1),
	1997: uint16(aux_sym_COMMENT_PREFIX_token1),
	1998: uint16(253),
	1999: uint16(1),
	2000: uint16(aux_sym_multipart_form_data_token1),
	2001: uint16(38),
	2002: uint16(1),
	2003: uint16(sym__plain_comment),
	2004: uint16(39),
	2005: uint16(1),
	2006: uint16(sym__var_comment),
	2007: uint16(198),
	2008: uint16(1),
	2009: uint16(sym_external_body),
	2010: uint16(31),
	2011: uint16(2),
	2012: uint16(sym_comment),
	2013: uint16(aux_sym_multipart_form_data_repeat1),
	2014: uint16(247),
	2015: uint16(4),
	2017: uint16(aux_sym_request_separator_token1),
	2018: uint16(aux_sym_res_redirect_token1),
	2019: uint16(aux_sym_graphql_data_token1),
	2020: uint16(251),
	2021: uint16(4),
	2022: uint16(aux_sym_xml_body_token1),
	2023: uint16(aux_sym_json_body_token1),
	2024: uint16(anon_sym_DASH_DASH),
	2025: uint16(aux_sym_multipart_form_data_token2),
	2026: uint16(249),
	2027: uint16(12),
	2028: uint16(aux_sym_WORD_CHAR_token1),
	2029: uint16(aux_sym_PUNCTUATION_token1),
	2030: uint16(aux_sym_WS_token1),
	2031: uint16(aux_sym_NL_token1),
	2032: uint16(sym_method),
	2033: uint16(aux_sym_http_version_token1),
	2034: uint16(anon_sym_LBRACE_LBRACE),
	2035: uint16(anon_sym_LT),
	2036: uint16(anon_sym_GT),
	2037: uint16(anon_sym_AT2),
	2038: uint16(aux_sym_raw_body_token1),
	2039: uint16(aux_sym__blank_line_token1),
	2040: uint16(10),
	2041: uint16(259),
	2042: uint16(1),
	2043: uint16(aux_sym_COMMENT_PREFIX_token1),
	2044: uint16(265),
	2045: uint16(1),
	2046: uint16(anon_sym_LT2),
	2047: uint16(268),
	2048: uint16(1),
	2049: uint16(aux_sym_multipart_form_data_token1),
	2050: uint16(38),
	2051: uint16(1),
	2052: uint16(sym__plain_comment),
	2053: uint16(39),
	2054: uint16(1),
	2055: uint16(sym__var_comment),
	2056: uint16(198),
	2057: uint16(1),
	2058: uint16(sym_external_body),
	2059: uint16(31),
	2060: uint16(2),
	2061: uint16(sym_comment),
	2062: uint16(aux_sym_multipart_form_data_repeat1),
	2063: uint16(255),
	2064: uint16(4),
	2066: uint16(aux_sym_request_separator_token1),
	2067: uint16(aux_sym_res_redirect_token1),
	2068: uint16(aux_sym_graphql_data_token1),
	2069: uint16(262),
	2070: uint16(4),
	2071: uint16(aux_sym_xml_body_token1),
	2072: uint16(aux_sym_json_body_token1),
	2073: uint16(anon_sym_DASH_DASH),
	2074: uint16(aux_sym_multipart_form_data_token2),
	2075: uint16(257),
	2076: uint16(12),
	2077: uint16(aux_sym_WORD_CHAR_token1),
	2078: uint16(aux_sym_PUNCTUATION_token1),
	2079: uint16(aux_sym_WS_token1),
	2080: uint16(aux_sym_NL_token1),
	2081: uint16(sym_method),
	2082: uint16(aux_sym_http_version_token1),
	2083: uint16(anon_sym_LBRACE_LBRACE),
	2084: uint16(anon_sym_LT),
	2085: uint16(anon_sym_GT),
	2086: uint16(anon_sym_AT2),
	2087: uint16(aux_sym_raw_body_token1),
	2088: uint16(aux_sym__blank_line_token1),
	2089: uint16(19),
	2090: uint16(276),
	2091: uint16(1),
	2092: uint16(aux_sym_WS_token1),
	2093: uint16(279),
	2094: uint16(1),
	2095: uint16(aux_sym_COMMENT_PREFIX_token1),
	2096: uint16(282),
	2097: uint16(1),
	2098: uint16(sym_method),
	2099: uint16(285),
	2100: uint16(1),
	2101: uint16(aux_sym_http_version_token1),
	2102: uint16(288),
	2103: uint16(1),
	2104: uint16(anon_sym_LBRACE_LBRACE),
	2105: uint16(291),
	2106: uint16(1),
	2107: uint16(anon_sym_LT),
	2108: uint16(294),
	2109: uint16(1),
	2110: uint16(anon_sym_AT2),
	2111: uint16(297),
	2112: uint16(1),
	2113: uint16(aux_sym__blank_line_token1),
	2114: uint16(104),
	2115: uint16(1),
	2116: uint16(sym_response),
	2117: uint16(105),
	2118: uint16(1),
	2119: uint16(sym_request),
	2120: uint16(112),
	2121: uint16(1),
	2122: uint16(sym__plain_comment),
	2123: uint16(114),
	2124: uint16(1),
	2125: uint16(sym__section_content),
	2126: uint16(116),
	2127: uint16(1),
	2128: uint16(sym__var_comment),
	2129: uint16(200),
	2130: uint16(1),
	2131: uint16(sym_target_url),
	2132: uint16(214),
	2133: uint16(1),
	2134: uint16(sym_http_version),
	2135: uint16(271),
	2136: uint16(2),
	2138: uint16(aux_sym_request_separator_token1),
	2139: uint16(273),
	2140: uint16(2),
	2141: uint16(aux_sym_WORD_CHAR_token1),
	2142: uint16(aux_sym_PUNCTUATION_token1),
	2143: uint16(132),
	2144: uint16(2),
	2145: uint16(aux_sym__target_url_line),
	2146: uint16(sym_variable),
	2147: uint16(32),
	2148: uint16(4),
	2149: uint16(sym_comment),
	2150: uint16(sym_pre_request_script),
	2151: uint16(sym_variable_declaration),
	2152: uint16(sym__blank_line),
	2153: uint16(19),
	2154: uint16(7),
	2155: uint16(1),
	2156: uint16(aux_sym_WS_token1),
	2157: uint16(9),
	2158: uint16(1),
	2159: uint16(aux_sym_COMMENT_PREFIX_token1),
	2160: uint16(13),
	2161: uint16(1),
	2162: uint16(sym_method),
	2163: uint16(15),
	2164: uint16(1),
	2165: uint16(aux_sym_http_version_token1),
	2166: uint16(17),
	2167: uint16(1),
	2168: uint16(anon_sym_LBRACE_LBRACE),
	2169: uint16(19),
	2170: uint16(1),
	2171: uint16(anon_sym_LT),
	2172: uint16(21),
	2173: uint16(1),
	2174: uint16(anon_sym_AT2),
	2175: uint16(23),
	2176: uint16(1),
	2177: uint16(aux_sym__blank_line_token1),
	2178: uint16(104),
	2179: uint16(1),
	2180: uint16(sym_response),
	2181: uint16(105),
	2182: uint16(1),
	2183: uint16(sym_request),
	2184: uint16(112),
	2185: uint16(1),
	2186: uint16(sym__plain_comment),
	2187: uint16(115),
	2188: uint16(1),
	2189: uint16(sym__section_content),
	2190: uint16(116),
	2191: uint16(1),
	2192: uint16(sym__var_comment),
	2193: uint16(200),
	2194: uint16(1),
	2195: uint16(sym_target_url),
	2196: uint16(214),
	2197: uint16(1),
	2198: uint16(sym_http_version),
	2199: uint16(5),
	2200: uint16(2),
	2201: uint16(aux_sym_WORD_CHAR_token1),
	2202: uint16(aux_sym_PUNCTUATION_token1),
	2203: uint16(300),
	2204: uint16(2),
	2206: uint16(aux_sym_request_separator_token1),
	2207: uint16(132),
	2208: uint16(2),
	2209: uint16(aux_sym__target_url_line),
	2210: uint16(sym_variable),
	2211: uint16(32),
	2212: uint16(4),
	2213: uint16(sym_comment),
	2214: uint16(sym_pre_request_script),
	2215: uint16(sym_variable_declaration),
	2216: uint16(sym__blank_line),
	2217: uint16(5),
	2218: uint16(306),
	2219: uint16(1),
	2220: uint16(aux_sym_COMMENT_PREFIX_token1),
	2221: uint16(309),
	2222: uint16(1),
	2223: uint16(aux_sym__raw_body_token1),
	2224: uint16(51),
	2225: uint16(1),
	2226: uint16(sym__raw_body),
	2227: uint16(302),
	2228: uint16(4),
	2230: uint16(aux_sym_request_separator_token1),
	2231: uint16(aux_sym_res_redirect_token1),
	2232: uint16(aux_sym_graphql_json_body_token1),
	2233: uint16(304),
	2234: uint16(17),
	2235: uint16(aux_sym_WORD_CHAR_token1),
	2236: uint16(aux_sym_PUNCTUATION_token1),
	2237: uint16(aux_sym_WS_token1),
	2238: uint16(aux_sym_NL_token1),
	2239: uint16(sym_method),
	2240: uint16(aux_sym_http_version_token1),
	2241: uint16(anon_sym_LBRACE_LBRACE),
	2242: uint16(anon_sym_LT),
	2243: uint16(anon_sym_GT),
	2244: uint16(anon_sym_AT2),
	2245: uint16(aux_sym_xml_body_token1),
	2246: uint16(aux_sym_json_body_token1),
	2247: uint16(aux_sym_graphql_data_token1),
	2248: uint16(anon_sym_LT2),
	2249: uint16(anon_sym_DASH_DASH),
	2250: uint16(aux_sym_raw_body_token1),
	2251: uint16(aux_sym__blank_line_token1),
	2252: uint16(5),
	2253: uint16(309),
	2254: uint16(1),
	2255: uint16(aux_sym__raw_body_token1),
	2256: uint16(315),
	2257: uint16(1),
	2258: uint16(aux_sym_COMMENT_PREFIX_token1),
	2259: uint16(50),
	2260: uint16(1),
	2261: uint16(sym__raw_body),
	2262: uint16(311),
	2263: uint16(4),
	2265: uint16(aux_sym_request_separator_token1),
	2266: uint16(aux_sym_res_redirect_token1),
	2267: uint16(aux_sym_graphql_json_body_token1),
	2268: uint16(313),
	2269: uint16(17),
	2270: uint16(aux_sym_WORD_CHAR_token1),
	2271: uint16(aux_sym_PUNCTUATION_token1),
	2272: uint16(aux_sym_WS_token1),
	2273: uint16(aux_sym_NL_token1),
	2274: uint16(sym_method),
	2275: uint16(aux_sym_http_version_token1),
	2276: uint16(anon_sym_LBRACE_LBRACE),
	2277: uint16(anon_sym_LT),
	2278: uint16(anon_sym_GT),
	2279: uint16(anon_sym_AT2),
	2280: uint16(aux_sym_xml_body_token1),
	2281: uint16(aux_sym_json_body_token1),
	2282: uint16(aux_sym_graphql_data_token1),
	2283: uint16(anon_sym_LT2),
	2284: uint16(anon_sym_DASH_DASH),
	2285: uint16(aux_sym_raw_body_token1),
	2286: uint16(aux_sym__blank_line_token1),
	2287: uint16(2),
	2288: uint16(318),
	2289: uint16(8),
	2291: uint16(aux_sym_request_separator_token1),
	2292: uint16(aux_sym_res_redirect_token1),
	2293: uint16(aux_sym_xml_body_token1),
	2294: uint16(aux_sym_json_body_token1),
	2295: uint16(aux_sym_graphql_data_token1),
	2296: uint16(anon_sym_DASH_DASH),
	2297: uint16(aux_sym_multipart_form_data_token2),
	2298: uint16(320),
	2299: uint16(15),
	2300: uint16(aux_sym_WORD_CHAR_token1),
	2301: uint16(aux_sym_PUNCTUATION_token1),
	2302: uint16(aux_sym_WS_token1),
	2303: uint16(aux_sym_NL_token1),
	2304: uint16(aux_sym_COMMENT_PREFIX_token1),
	2305: uint16(sym_method),
	2306: uint16(aux_sym_http_version_token1),
	2307: uint16(anon_sym_LBRACE_LBRACE),
	2308: uint16(anon_sym_LT),
	2309: uint16(anon_sym_GT),
	2310: uint16(anon_sym_AT2),
	2311: uint16(anon_sym_LT2),
	2312: uint16(aux_sym_multipart_form_data_token1),
	2313: uint16(aux_sym_raw_body_token1),
	2314: uint16(aux_sym__blank_line_token1),
	2315: uint16(5),
	2316: uint16(322),
	2317: uint16(1),
	2318: uint16(aux_sym_COMMENT_PREFIX_token1),
	2319: uint16(325),
	2320: uint16(1),
	2321: uint16(aux_sym__raw_body_token1),
	2322: uint16(67),
	2323: uint16(1),
	2324: uint16(sym__raw_body),
	2325: uint16(311),
	2326: uint16(3),
	2328: uint16(aux_sym_request_separator_token1),
	2329: uint16(aux_sym_res_redirect_token1),
	2330: uint16(313),
	2331: uint16(17),
	2332: uint16(aux_sym_WORD_CHAR_token1),
	2333: uint16(aux_sym_PUNCTUATION_token1),
	2334: uint16(aux_sym_WS_token1),
	2335: uint16(aux_sym_NL_token1),
	2336: uint16(sym_method),
	2337: uint16(aux_sym_http_version_token1),
	2338: uint16(anon_sym_LBRACE_LBRACE),
	2339: uint16(anon_sym_LT),
	2340: uint16(anon_sym_GT),
	2341: uint16(anon_sym_AT2),
	2342: uint16(aux_sym_xml_body_token1),
	2343: uint16(aux_sym_json_body_token1),
	2344: uint16(aux_sym_graphql_data_token1),
	2345: uint16(anon_sym_LT2),
	2346: uint16(anon_sym_DASH_DASH),
	2347: uint16(aux_sym_raw_body_token1),
	2348: uint16(aux_sym__blank_line_token1),
	2349: uint16(2),
	2350: uint16(327),
	2351: uint16(8),
	2353: uint16(aux_sym_request_separator_token1),
	2354: uint16(aux_sym_res_redirect_token1),
	2355: uint16(aux_sym_xml_body_token1),
	2356: uint16(aux_sym_json_body_token1),
	2357: uint16(aux_sym_graphql_data_token1),
	2358: uint16(anon_sym_DASH_DASH),
	2359: uint16(aux_sym_multipart_form_data_token2),
	2360: uint16(329),
	2361: uint16(15),
	2362: uint16(aux_sym_WORD_CHAR_token1),
	2363: uint16(aux_sym_PUNCTUATION_token1),
	2364: uint16(aux_sym_WS_token1),
	2365: uint16(aux_sym_NL_token1),
	2366: uint16(aux_sym_COMMENT_PREFIX_token1),
	2367: uint16(sym_method),
	2368: uint16(aux_sym_http_version_token1),
	2369: uint16(anon_sym_LBRACE_LBRACE),
	2370: uint16(anon_sym_LT),
	2371: uint16(anon_sym_GT),
	2372: uint16(anon_sym_AT2),
	2373: uint16(anon_sym_LT2),
	2374: uint16(aux_sym_multipart_form_data_token1),
	2375: uint16(aux_sym_raw_body_token1),
	2376: uint16(aux_sym__blank_line_token1),
	2377: uint16(2),
	2378: uint16(331),
	2379: uint16(8),
	2381: uint16(aux_sym_request_separator_token1),
	2382: uint16(aux_sym_res_redirect_token1),
	2383: uint16(aux_sym_xml_body_token1),
	2384: uint16(aux_sym_json_body_token1),
	2385: uint16(aux_sym_graphql_data_token1),
	2386: uint16(anon_sym_DASH_DASH),
	2387: uint16(aux_sym_multipart_form_data_token2),
	2388: uint16(333),
	2389: uint16(15),
	2390: uint16(aux_sym_WORD_CHAR_token1),
	2391: uint16(aux_sym_PUNCTUATION_token1),
	2392: uint16(aux_sym_WS_token1),
	2393: uint16(aux_sym_NL_token1),
	2394: uint16(aux_sym_COMMENT_PREFIX_token1),
	2395: uint16(sym_method),
	2396: uint16(aux_sym_http_version_token1),
	2397: uint16(anon_sym_LBRACE_LBRACE),
	2398: uint16(anon_sym_LT),
	2399: uint16(anon_sym_GT),
	2400: uint16(anon_sym_AT2),
	2401: uint16(anon_sym_LT2),
	2402: uint16(aux_sym_multipart_form_data_token1),
	2403: uint16(aux_sym_raw_body_token1),
	2404: uint16(aux_sym__blank_line_token1),
	2405: uint16(2),
	2406: uint16(255),
	2407: uint16(8),
	2409: uint16(aux_sym_request_separator_token1),
	2410: uint16(aux_sym_res_redirect_token1),
	2411: uint16(aux_sym_xml_body_token1),
	2412: uint16(aux_sym_json_body_token1),
	2413: uint16(aux_sym_graphql_data_token1),
	2414: uint16(anon_sym_DASH_DASH),
	2415: uint16(aux_sym_multipart_form_data_token2),
	2416: uint16(257),
	2417: uint16(15),
	2418: uint16(aux_sym_WORD_CHAR_token1),
	2419: uint16(aux_sym_PUNCTUATION_token1),
	2420: uint16(aux_sym_WS_token1),
	2421: uint16(aux_sym_NL_token1),
	2422: uint16(aux_sym_COMMENT_PREFIX_token1),
	2423: uint16(sym_method),
	2424: uint16(aux_sym_http_version_token1),
	2425: uint16(anon_sym_LBRACE_LBRACE),
	2426: uint16(anon_sym_LT),
	2427: uint16(anon_sym_GT),
	2428: uint16(anon_sym_AT2),
	2429: uint16(anon_sym_LT2),
	2430: uint16(aux_sym_multipart_form_data_token1),
	2431: uint16(aux_sym_raw_body_token1),
	2432: uint16(aux_sym__blank_line_token1),
	2433: uint16(5),
	2434: uint16(325),
	2435: uint16(1),
	2436: uint16(aux_sym__raw_body_token1),
	2437: uint16(339),
	2438: uint16(1),
	2439: uint16(aux_sym_COMMENT_PREFIX_token1),
	2440: uint16(61),
	2441: uint16(1),
	2442: uint16(sym__raw_body),
	2443: uint16(335),
	2444: uint16(3),
	2446: uint16(aux_sym_request_separator_token1),
	2447: uint16(aux_sym_res_redirect_token1),
	2448: uint16(337),
	2449: uint16(17),
	2450: uint16(aux_sym_WORD_CHAR_token1),
	2451: uint16(aux_sym_PUNCTUATION_token1),
	2452: uint16(aux_sym_WS_token1),
	2453: uint16(aux_sym_NL_token1),
	2454: uint16(sym_method),
	2455: uint16(aux_sym_http_version_token1),
	2456: uint16(anon_sym_LBRACE_LBRACE),
	2457: uint16(anon_sym_LT),
	2458: uint16(anon_sym_GT),
	2459: uint16(anon_sym_AT2),
	2460: uint16(aux_sym_xml_body_token1),
	2461: uint16(aux_sym_json_body_token1),
	2462: uint16(aux_sym_graphql_data_token1),
	2463: uint16(anon_sym_LT2),
	2464: uint16(anon_sym_DASH_DASH),
	2465: uint16(aux_sym_raw_body_token1),
	2466: uint16(aux_sym__blank_line_token1),
	2467: uint16(5),
	2468: uint16(325),
	2469: uint16(1),
	2470: uint16(aux_sym__raw_body_token1),
	2471: uint16(342),
	2472: uint16(1),
	2473: uint16(aux_sym_COMMENT_PREFIX_token1),
	2474: uint16(64),
	2475: uint16(1),
	2476: uint16(sym__raw_body),
	2477: uint16(302),
	2478: uint16(3),
	2480: uint16(aux_sym_request_separator_token1),
	2481: uint16(aux_sym_res_redirect_token1),
	2482: uint16(304),
	2483: uint16(17),
	2484: uint16(aux_sym_WORD_CHAR_token1),
	2485: uint16(aux_sym_PUNCTUATION_token1),
	2486: uint16(aux_sym_WS_token1),
	2487: uint16(aux_sym_NL_token1),
	2488: uint16(sym_method),
	2489: uint16(aux_sym_http_version_token1),
	2490: uint16(anon_sym_LBRACE_LBRACE),
	2491: uint16(anon_sym_LT),
	2492: uint16(anon_sym_GT),
	2493: uint16(anon_sym_AT2),
	2494: uint16(aux_sym_xml_body_token1),
	2495: uint16(aux_sym_json_body_token1),
	2496: uint16(aux_sym_graphql_data_token1),
	2497: uint16(anon_sym_LT2),
	2498: uint16(anon_sym_DASH_DASH),
	2499: uint16(aux_sym_raw_body_token1),
	2500: uint16(aux_sym__blank_line_token1),
	2501: uint16(5),
	2502: uint16(325),
	2503: uint16(1),
	2504: uint16(aux_sym__raw_body_token1),
	2505: uint16(349),
	2506: uint16(1),
	2507: uint16(aux_sym_COMMENT_PREFIX_token1),
	2508: uint16(60),
	2509: uint16(1),
	2510: uint16(sym__raw_body),
	2511: uint16(345),
	2512: uint16(3),
	2514: uint16(aux_sym_request_separator_token1),
	2515: uint16(aux_sym_res_redirect_token1),
	2516: uint16(347),
	2517: uint16(17),
	2518: uint16(aux_sym_WORD_CHAR_token1),
	2519: uint16(aux_sym_PUNCTUATION_token1),
	2520: uint16(aux_sym_WS_token1),
	2521: uint16(aux_sym_NL_token1),
	2522: uint16(sym_method),
	2523: uint16(aux_sym_http_version_token1),
	2524: uint16(anon_sym_LBRACE_LBRACE),
	2525: uint16(anon_sym_LT),
	2526: uint16(anon_sym_GT),
	2527: uint16(anon_sym_AT2),
	2528: uint16(aux_sym_xml_body_token1),
	2529: uint16(aux_sym_json_body_token1),
	2530: uint16(aux_sym_graphql_data_token1),
	2531: uint16(anon_sym_LT2),
	2532: uint16(anon_sym_DASH_DASH),
	2533: uint16(aux_sym_raw_body_token1),
	2534: uint16(aux_sym__blank_line_token1),
	2535: uint16(2),
	2536: uint16(352),
	2537: uint16(8),
	2539: uint16(aux_sym_request_separator_token1),
	2540: uint16(aux_sym_res_redirect_token1),
	2541: uint16(aux_sym_xml_body_token1),
	2542: uint16(aux_sym_json_body_token1),
	2543: uint16(aux_sym_graphql_data_token1),
	2544: uint16(anon_sym_DASH_DASH),
	2545: uint16(aux_sym_multipart_form_data_token2),
	2546: uint16(354),
	2547: uint16(15),
	2548: uint16(aux_sym_WORD_CHAR_token1),
	2549: uint16(aux_sym_PUNCTUATION_token1),
	2550: uint16(aux_sym_WS_token1),
	2551: uint16(aux_sym_NL_token1),
	2552: uint16(aux_sym_COMMENT_PREFIX_token1),
	2553: uint16(sym_method),
	2554: uint16(aux_sym_http_version_token1),
	2555: uint16(anon_sym_LBRACE_LBRACE),
	2556: uint16(anon_sym_LT),
	2557: uint16(anon_sym_GT),
	2558: uint16(anon_sym_AT2),
	2559: uint16(anon_sym_LT2),
	2560: uint16(aux_sym_multipart_form_data_token1),
	2561: uint16(aux_sym_raw_body_token1),
	2562: uint16(aux_sym__blank_line_token1),
	2563: uint16(2),
	2564: uint16(356),
	2565: uint16(8),
	2567: uint16(aux_sym_request_separator_token1),
	2568: uint16(aux_sym_res_redirect_token1),
	2569: uint16(aux_sym_xml_body_token1),
	2570: uint16(aux_sym_json_body_token1),
	2571: uint16(aux_sym_graphql_data_token1),
	2572: uint16(anon_sym_DASH_DASH),
	2573: uint16(aux_sym_multipart_form_data_token2),
	2574: uint16(358),
	2575: uint16(15),
	2576: uint16(aux_sym_WORD_CHAR_token1),
	2577: uint16(aux_sym_PUNCTUATION_token1),
	2578: uint16(aux_sym_WS_token1),
	2579: uint16(aux_sym_NL_token1),
	2580: uint16(aux_sym_COMMENT_PREFIX_token1),
	2581: uint16(sym_method),
	2582: uint16(aux_sym_http_version_token1),
	2583: uint16(anon_sym_LBRACE_LBRACE),
	2584: uint16(anon_sym_LT),
	2585: uint16(anon_sym_GT),
	2586: uint16(anon_sym_AT2),
	2587: uint16(anon_sym_LT2),
	2588: uint16(aux_sym_multipart_form_data_token1),
	2589: uint16(aux_sym_raw_body_token1),
	2590: uint16(aux_sym__blank_line_token1),
	2591: uint16(2),
	2592: uint16(360),
	2593: uint16(8),
	2595: uint16(aux_sym_request_separator_token1),
	2596: uint16(aux_sym_res_redirect_token1),
	2597: uint16(aux_sym_xml_body_token1),
	2598: uint16(aux_sym_json_body_token1),
	2599: uint16(aux_sym_graphql_data_token1),
	2600: uint16(anon_sym_DASH_DASH),
	2601: uint16(aux_sym_multipart_form_data_token2),
	2602: uint16(362),
	2603: uint16(15),
	2604: uint16(aux_sym_WORD_CHAR_token1),
	2605: uint16(aux_sym_PUNCTUATION_token1),
	2606: uint16(aux_sym_WS_token1),
	2607: uint16(aux_sym_NL_token1),
	2608: uint16(aux_sym_COMMENT_PREFIX_token1),
	2609: uint16(sym_method),
	2610: uint16(aux_sym_http_version_token1),
	2611: uint16(anon_sym_LBRACE_LBRACE),
	2612: uint16(anon_sym_LT),
	2613: uint16(anon_sym_GT),
	2614: uint16(anon_sym_AT2),
	2615: uint16(anon_sym_LT2),
	2616: uint16(aux_sym_multipart_form_data_token1),
	2617: uint16(aux_sym_raw_body_token1),
	2618: uint16(aux_sym__blank_line_token1),
	2619: uint16(4),
	2620: uint16(368),
	2621: uint16(1),
	2622: uint16(aux_sym_graphql_json_body_token1),
	2623: uint16(52),
	2624: uint16(1),
	2625: uint16(sym_graphql_json_body),
	2626: uint16(364),
	2627: uint16(6),
	2629: uint16(aux_sym_request_separator_token1),
	2630: uint16(aux_sym_res_redirect_token1),
	2631: uint16(aux_sym_xml_body_token1),
	2632: uint16(aux_sym_graphql_data_token1),
	2633: uint16(anon_sym_DASH_DASH),
	2634: uint16(366),
	2635: uint16(15),
	2636: uint16(aux_sym_WORD_CHAR_token1),
	2637: uint16(aux_sym_PUNCTUATION_token1),
	2638: uint16(aux_sym_WS_token1),
	2639: uint16(aux_sym_NL_token1),
	2640: uint16(aux_sym_COMMENT_PREFIX_token1),
	2641: uint16(sym_method),
	2642: uint16(aux_sym_http_version_token1),
	2643: uint16(anon_sym_LBRACE_LBRACE),
	2644: uint16(anon_sym_LT),
	2645: uint16(anon_sym_GT),
	2646: uint16(anon_sym_AT2),
	2647: uint16(aux_sym_json_body_token1),
	2648: uint16(anon_sym_LT2),
	2649: uint16(aux_sym_raw_body_token1),
	2650: uint16(aux_sym__blank_line_token1),
	2651: uint16(5),
	2652: uint16(374),
	2653: uint16(1),
	2654: uint16(aux_sym_WS_token1),
	2655: uint16(377),
	2656: uint16(1),
	2657: uint16(aux_sym__blank_line_token1),
	2658: uint16(48),
	2659: uint16(2),
	2660: uint16(sym__blank_line),
	2661: uint16(aux_sym___body_repeat1),
	2662: uint16(370),
	2663: uint16(7),
	2665: uint16(aux_sym_request_separator_token1),
	2666: uint16(aux_sym_res_redirect_token1),
	2667: uint16(aux_sym_xml_body_token1),
	2668: uint16(aux_sym_json_body_token1),
	2669: uint16(aux_sym_graphql_data_token1),
	2670: uint16(anon_sym_DASH_DASH),
	2671: uint16(372),
	2672: uint16(12),
	2673: uint16(aux_sym_WORD_CHAR_token1),
	2674: uint16(aux_sym_PUNCTUATION_token1),
	2675: uint16(aux_sym_NL_token1),
	2676: uint16(aux_sym_COMMENT_PREFIX_token1),
	2677: uint16(sym_method),
	2678: uint16(aux_sym_http_version_token1),
	2679: uint16(anon_sym_LBRACE_LBRACE),
	2680: uint16(anon_sym_LT),
	2681: uint16(anon_sym_GT),
	2682: uint16(anon_sym_AT2),
	2683: uint16(anon_sym_LT2),
	2684: uint16(aux_sym_raw_body_token1),
	2685: uint16(2),
	2686: uint16(380),
	2687: uint16(7),
	2689: uint16(aux_sym_request_separator_token1),
	2690: uint16(aux_sym_res_redirect_token1),
	2691: uint16(aux_sym_xml_body_token1),
	2692: uint16(aux_sym_graphql_data_token1),
	2693: uint16(aux_sym_graphql_json_body_token1),
	2694: uint16(anon_sym_DASH_DASH),
	2695: uint16(382),
	2696: uint16(15),
	2697: uint16(aux_sym_WORD_CHAR_token1),
	2698: uint16(aux_sym_PUNCTUATION_token1),
	2699: uint16(aux_sym_WS_token1),
	2700: uint16(aux_sym_NL_token1),
	2701: uint16(aux_sym_COMMENT_PREFIX_token1),
	2702: uint16(sym_method),
	2703: uint16(aux_sym_http_version_token1),
	2704: uint16(anon_sym_LBRACE_LBRACE),
	2705: uint16(anon_sym_LT),
	2706: uint16(anon_sym_GT),
	2707: uint16(anon_sym_AT2),
	2708: uint16(aux_sym_json_body_token1),
	2709: uint16(anon_sym_LT2),
	2710: uint16(aux_sym_raw_body_token1),
	2711: uint16(aux_sym__blank_line_token1),
	2712: uint16(2),
	2713: uint16(302),
	2714: uint16(7),
	2716: uint16(aux_sym_request_separator_token1),
	2717: uint16(aux_sym_res_redirect_token1),
	2718: uint16(aux_sym_xml_body_token1),
	2719: uint16(aux_sym_graphql_data_token1),
	2720: uint16(aux_sym_graphql_json_body_token1),
	2721: uint16(anon_sym_DASH_DASH),
	2722: uint16(304),
	2723: uint16(15),
	2724: uint16(aux_sym_WORD_CHAR_token1),
	2725: uint16(aux_sym_PUNCTUATION_token1),
	2726: uint16(aux_sym_WS_token1),
	2727: uint16(aux_sym_NL_token1),
	2728: uint16(aux_sym_COMMENT_PREFIX_token1),
	2729: uint16(sym_method),
	2730: uint16(aux_sym_http_version_token1),
	2731: uint16(anon_sym_LBRACE_LBRACE),
	2732: uint16(anon_sym_LT),
	2733: uint16(anon_sym_GT),
	2734: uint16(anon_sym_AT2),
	2735: uint16(aux_sym_json_body_token1),
	2736: uint16(anon_sym_LT2),
	2737: uint16(aux_sym_raw_body_token1),
	2738: uint16(aux_sym__blank_line_token1),
	2739: uint16(2),
	2740: uint16(384),
	2741: uint16(7),
	2743: uint16(aux_sym_request_separator_token1),
	2744: uint16(aux_sym_res_redirect_token1),
	2745: uint16(aux_sym_xml_body_token1),
	2746: uint16(aux_sym_graphql_data_token1),
	2747: uint16(aux_sym_graphql_json_body_token1),
	2748: uint16(anon_sym_DASH_DASH),
	2749: uint16(386),
	2750: uint16(15),
	2751: uint16(aux_sym_WORD_CHAR_token1),
	2752: uint16(aux_sym_PUNCTUATION_token1),
	2753: uint16(aux_sym_WS_token1),
	2754: uint16(aux_sym_NL_token1),
	2755: uint16(aux_sym_COMMENT_PREFIX_token1),
	2756: uint16(sym_method),
	2757: uint16(aux_sym_http_version_token1),
	2758: uint16(anon_sym_LBRACE_LBRACE),
	2759: uint16(anon_sym_LT),
	2760: uint16(anon_sym_GT),
	2761: uint16(anon_sym_AT2),
	2762: uint16(aux_sym_json_body_token1),
	2763: uint16(anon_sym_LT2),
	2764: uint16(aux_sym_raw_body_token1),
	2765: uint16(aux_sym__blank_line_token1),
	2766: uint16(2),
	2767: uint16(388),
	2768: uint16(7),
	2770: uint16(aux_sym_request_separator_token1),
	2771: uint16(aux_sym_res_redirect_token1),
	2772: uint16(aux_sym_xml_body_token1),
	2773: uint16(aux_sym_json_body_token1),
	2774: uint16(aux_sym_graphql_data_token1),
	2775: uint16(anon_sym_DASH_DASH),
	2776: uint16(390),
	2777: uint16(14),
	2778: uint16(aux_sym_WORD_CHAR_token1),
	2779: uint16(aux_sym_PUNCTUATION_token1),
	2780: uint16(aux_sym_WS_token1),
	2781: uint16(aux_sym_NL_token1),
	2782: uint16(aux_sym_COMMENT_PREFIX_token1),
	2783: uint16(sym_method),
	2784: uint16(aux_sym_http_version_token1),
	2785: uint16(anon_sym_LBRACE_LBRACE),
	2786: uint16(anon_sym_LT),
	2787: uint16(anon_sym_GT),
	2788: uint16(anon_sym_AT2),
	2789: uint16(anon_sym_LT2),
	2790: uint16(aux_sym_raw_body_token1),
	2791: uint16(aux_sym__blank_line_token1),
	2792: uint16(2),
	2793: uint16(392),
	2794: uint16(7),
	2796: uint16(aux_sym_request_separator_token1),
	2797: uint16(aux_sym_res_redirect_token1),
	2798: uint16(aux_sym_xml_body_token1),
	2799: uint16(aux_sym_json_body_token1),
	2800: uint16(aux_sym_graphql_data_token1),
	2801: uint16(anon_sym_DASH_DASH),
	2802: uint16(394),
	2803: uint16(14),
	2804: uint16(aux_sym_WORD_CHAR_token1),
	2805: uint16(aux_sym_PUNCTUATION_token1),
	2806: uint16(aux_sym_WS_token1),
	2807: uint16(aux_sym_NL_token1),
	2808: uint16(aux_sym_COMMENT_PREFIX_token1),
	2809: uint16(sym_method),
	2810: uint16(aux_sym_http_version_token1),
	2811: uint16(anon_sym_LBRACE_LBRACE),
	2812: uint16(anon_sym_LT),
	2813: uint16(anon_sym_GT),
	2814: uint16(anon_sym_AT2),
	2815: uint16(anon_sym_LT2),
	2816: uint16(aux_sym_raw_body_token1),
	2817: uint16(aux_sym__blank_line_token1),
	2818: uint16(2),
	2819: uint16(396),
	2820: uint16(7),
	2822: uint16(aux_sym_request_separator_token1),
	2823: uint16(aux_sym_res_redirect_token1),
	2824: uint16(aux_sym_xml_body_token1),
	2825: uint16(aux_sym_json_body_token1),
	2826: uint16(aux_sym_graphql_data_token1),
	2827: uint16(anon_sym_DASH_DASH),
	2828: uint16(398),
	2829: uint16(14),
	2830: uint16(aux_sym_WORD_CHAR_token1),
	2831: uint16(aux_sym_PUNCTUATION_token1),
	2832: uint16(aux_sym_WS_token1),
	2833: uint16(aux_sym_NL_token1),
	2834: uint16(aux_sym_COMMENT_PREFIX_token1),
	2835: uint16(sym_method),
	2836: uint16(aux_sym_http_version_token1),
	2837: uint16(anon_sym_LBRACE_LBRACE),
	2838: uint16(anon_sym_LT),
	2839: uint16(anon_sym_GT),
	2840: uint16(anon_sym_AT2),
	2841: uint16(anon_sym_LT2),
	2842: uint16(aux_sym_raw_body_token1),
	2843: uint16(aux_sym__blank_line_token1),
	2844: uint16(2),
	2845: uint16(360),
	2846: uint16(7),
	2848: uint16(aux_sym_request_separator_token1),
	2849: uint16(aux_sym_res_redirect_token1),
	2850: uint16(aux_sym_xml_body_token1),
	2851: uint16(aux_sym_json_body_token1),
	2852: uint16(aux_sym_graphql_data_token1),
	2853: uint16(anon_sym_DASH_DASH),
	2854: uint16(362),
	2855: uint16(14),
	2856: uint16(aux_sym_WORD_CHAR_token1),
	2857: uint16(aux_sym_PUNCTUATION_token1),
	2858: uint16(aux_sym_WS_token1),
	2859: uint16(aux_sym_NL_token1),
	2860: uint16(aux_sym_COMMENT_PREFIX_token1),
	2861: uint16(sym_method),
	2862: uint16(aux_sym_http_version_token1),
	2863: uint16(anon_sym_LBRACE_LBRACE),
	2864: uint16(anon_sym_LT),
	2865: uint16(anon_sym_GT),
	2866: uint16(anon_sym_AT2),
	2867: uint16(anon_sym_LT2),
	2868: uint16(aux_sym_raw_body_token1),
	2869: uint16(aux_sym__blank_line_token1),
	2870: uint16(2),
	2871: uint16(356),
	2872: uint16(7),
	2874: uint16(aux_sym_request_separator_token1),
	2875: uint16(aux_sym_res_redirect_token1),
	2876: uint16(aux_sym_xml_body_token1),
	2877: uint16(aux_sym_json_body_token1),
	2878: uint16(aux_sym_graphql_data_token1),
	2879: uint16(anon_sym_DASH_DASH),
	2880: uint16(358),
	2881: uint16(14),
	2882: uint16(aux_sym_WORD_CHAR_token1),
	2883: uint16(aux_sym_PUNCTUATION_token1),
	2884: uint16(aux_sym_WS_token1),
	2885: uint16(aux_sym_NL_token1),
	2886: uint16(aux_sym_COMMENT_PREFIX_token1),
	2887: uint16(sym_method),
	2888: uint16(aux_sym_http_version_token1),
	2889: uint16(anon_sym_LBRACE_LBRACE),
	2890: uint16(anon_sym_LT),
	2891: uint16(anon_sym_GT),
	2892: uint16(anon_sym_AT2),
	2893: uint16(anon_sym_LT2),
	2894: uint16(aux_sym_raw_body_token1),
	2895: uint16(aux_sym__blank_line_token1),
	2896: uint16(2),
	2897: uint16(400),
	2898: uint16(7),
	2900: uint16(aux_sym_request_separator_token1),
	2901: uint16(aux_sym_res_redirect_token1),
	2902: uint16(aux_sym_xml_body_token1),
	2903: uint16(aux_sym_json_body_token1),
	2904: uint16(aux_sym_graphql_data_token1),
	2905: uint16(anon_sym_DASH_DASH),
	2906: uint16(402),
	2907: uint16(14),
	2908: uint16(aux_sym_WORD_CHAR_token1),
	2909: uint16(aux_sym_PUNCTUATION_token1),
	2910: uint16(aux_sym_WS_token1),
	2911: uint16(aux_sym_NL_token1),
	2912: uint16(aux_sym_COMMENT_PREFIX_token1),
	2913: uint16(sym_method),
	2914: uint16(aux_sym_http_version_token1),
	2915: uint16(anon_sym_LBRACE_LBRACE),
	2916: uint16(anon_sym_LT),
	2917: uint16(anon_sym_GT),
	2918: uint16(anon_sym_AT2),
	2919: uint16(anon_sym_LT2),
	2920: uint16(aux_sym_raw_body_token1),
	2921: uint16(aux_sym__blank_line_token1),
	2922: uint16(2),
	2923: uint16(404),
	2924: uint16(7),
	2926: uint16(aux_sym_request_separator_token1),
	2927: uint16(aux_sym_res_redirect_token1),
	2928: uint16(aux_sym_xml_body_token1),
	2929: uint16(aux_sym_json_body_token1),
	2930: uint16(aux_sym_graphql_data_token1),
	2931: uint16(anon_sym_DASH_DASH),
	2932: uint16(406),
	2933: uint16(14),
	2934: uint16(aux_sym_WORD_CHAR_token1),
	2935: uint16(aux_sym_PUNCTUATION_token1),
	2936: uint16(aux_sym_WS_token1),
	2937: uint16(aux_sym_NL_token1),
	2938: uint16(aux_sym_COMMENT_PREFIX_token1),
	2939: uint16(sym_method),
	2940: uint16(aux_sym_http_version_token1),
	2941: uint16(anon_sym_LBRACE_LBRACE),
	2942: uint16(anon_sym_LT),
	2943: uint16(anon_sym_GT),
	2944: uint16(anon_sym_AT2),
	2945: uint16(anon_sym_LT2),
	2946: uint16(aux_sym_raw_body_token1),
	2947: uint16(aux_sym__blank_line_token1),
	2948: uint16(2),
	2949: uint16(408),
	2950: uint16(7),
	2952: uint16(aux_sym_request_separator_token1),
	2953: uint16(aux_sym_res_redirect_token1),
	2954: uint16(aux_sym_xml_body_token1),
	2955: uint16(aux_sym_json_body_token1),
	2956: uint16(aux_sym_graphql_data_token1),
	2957: uint16(anon_sym_DASH_DASH),
	2958: uint16(410),
	2959: uint16(14),
	2960: uint16(aux_sym_WORD_CHAR_token1),
	2961: uint16(aux_sym_PUNCTUATION_token1),
	2962: uint16(aux_sym_WS_token1),
	2963: uint16(aux_sym_NL_token1),
	2964: uint16(aux_sym_COMMENT_PREFIX_token1),
	2965: uint16(sym_method),
	2966: uint16(aux_sym_http_version_token1),
	2967: uint16(anon_sym_LBRACE_LBRACE),
	2968: uint16(anon_sym_LT),
	2969: uint16(anon_sym_GT),
	2970: uint16(anon_sym_AT2),
	2971: uint16(anon_sym_LT2),
	2972: uint16(aux_sym_raw_body_token1),
	2973: uint16(aux_sym__blank_line_token1),
	2974: uint16(2),
	2975: uint16(335),
	2976: uint16(7),
	2978: uint16(aux_sym_request_separator_token1),
	2979: uint16(aux_sym_res_redirect_token1),
	2980: uint16(aux_sym_xml_body_token1),
	2981: uint16(aux_sym_json_body_token1),
	2982: uint16(aux_sym_graphql_data_token1),
	2983: uint16(anon_sym_DASH_DASH),
	2984: uint16(337),
	2985: uint16(14),
	2986: uint16(aux_sym_WORD_CHAR_token1),
	2987: uint16(aux_sym_PUNCTUATION_token1),
	2988: uint16(aux_sym_WS_token1),
	2989: uint16(aux_sym_NL_token1),
	2990: uint16(aux_sym_COMMENT_PREFIX_token1),
	2991: uint16(sym_method),
	2992: uint16(aux_sym_http_version_token1),
	2993: uint16(anon_sym_LBRACE_LBRACE),
	2994: uint16(anon_sym_LT),
	2995: uint16(anon_sym_GT),
	2996: uint16(anon_sym_AT2),
	2997: uint16(anon_sym_LT2),
	2998: uint16(aux_sym_raw_body_token1),
	2999: uint16(aux_sym__blank_line_token1),
	3000: uint16(2),
	3001: uint16(412),
	3002: uint16(7),
	3004: uint16(aux_sym_request_separator_token1),
	3005: uint16(aux_sym_res_redirect_token1),
	3006: uint16(aux_sym_xml_body_token1),
	3007: uint16(aux_sym_json_body_token1),
	3008: uint16(aux_sym_graphql_data_token1),
	3009: uint16(anon_sym_DASH_DASH),
	3010: uint16(414),
	3011: uint16(14),
	3012: uint16(aux_sym_WORD_CHAR_token1),
	3013: uint16(aux_sym_PUNCTUATION_token1),
	3014: uint16(aux_sym_WS_token1),
	3015: uint16(aux_sym_NL_token1),
	3016: uint16(aux_sym_COMMENT_PREFIX_token1),
	3017: uint16(sym_method),
	3018: uint16(aux_sym_http_version_token1),
	3019: uint16(anon_sym_LBRACE_LBRACE),
	3020: uint16(anon_sym_LT),
	3021: uint16(anon_sym_GT),
	3022: uint16(anon_sym_AT2),
	3023: uint16(anon_sym_LT2),
	3024: uint16(aux_sym_raw_body_token1),
	3025: uint16(aux_sym__blank_line_token1),
	3026: uint16(2),
	3027: uint16(352),
	3028: uint16(7),
	3030: uint16(aux_sym_request_separator_token1),
	3031: uint16(aux_sym_res_redirect_token1),
	3032: uint16(aux_sym_xml_body_token1),
	3033: uint16(aux_sym_json_body_token1),
	3034: uint16(aux_sym_graphql_data_token1),
	3035: uint16(anon_sym_DASH_DASH),
	3036: uint16(354),
	3037: uint16(14),
	3038: uint16(aux_sym_WORD_CHAR_token1),
	3039: uint16(aux_sym_PUNCTUATION_token1),
	3040: uint16(aux_sym_WS_token1),
	3041: uint16(aux_sym_NL_token1),
	3042: uint16(aux_sym_COMMENT_PREFIX_token1),
	3043: uint16(sym_method),
	3044: uint16(aux_sym_http_version_token1),
	3045: uint16(anon_sym_LBRACE_LBRACE),
	3046: uint16(anon_sym_LT),
	3047: uint16(anon_sym_GT),
	3048: uint16(anon_sym_AT2),
	3049: uint16(anon_sym_LT2),
	3050: uint16(aux_sym_raw_body_token1),
	3051: uint16(aux_sym__blank_line_token1),
	3052: uint16(2),
	3053: uint16(416),
	3054: uint16(7),
	3056: uint16(aux_sym_request_separator_token1),
	3057: uint16(aux_sym_res_redirect_token1),
	3058: uint16(aux_sym_xml_body_token1),
	3059: uint16(aux_sym_json_body_token1),
	3060: uint16(aux_sym_graphql_data_token1),
	3061: uint16(anon_sym_DASH_DASH),
	3062: uint16(418),
	3063: uint16(14),
	3064: uint16(aux_sym_WORD_CHAR_token1),
	3065: uint16(aux_sym_PUNCTUATION_token1),
	3066: uint16(aux_sym_WS_token1),
	3067: uint16(aux_sym_NL_token1),
	3068: uint16(aux_sym_COMMENT_PREFIX_token1),
	3069: uint16(sym_method),
	3070: uint16(aux_sym_http_version_token1),
	3071: uint16(anon_sym_LBRACE_LBRACE),
	3072: uint16(anon_sym_LT),
	3073: uint16(anon_sym_GT),
	3074: uint16(anon_sym_AT2),
	3075: uint16(anon_sym_LT2),
	3076: uint16(aux_sym_raw_body_token1),
	3077: uint16(aux_sym__blank_line_token1),
	3078: uint16(2),
	3079: uint16(384),
	3080: uint16(7),
	3082: uint16(aux_sym_request_separator_token1),
	3083: uint16(aux_sym_res_redirect_token1),
	3084: uint16(aux_sym_xml_body_token1),
	3085: uint16(aux_sym_json_body_token1),
	3086: uint16(aux_sym_graphql_data_token1),
	3087: uint16(anon_sym_DASH_DASH),
	3088: uint16(386),
	3089: uint16(14),
	3090: uint16(aux_sym_WORD_CHAR_token1),
	3091: uint16(aux_sym_PUNCTUATION_token1),
	3092: uint16(aux_sym_WS_token1),
	3093: uint16(aux_sym_NL_token1),
	3094: uint16(aux_sym_COMMENT_PREFIX_token1),
	3095: uint16(sym_method),
	3096: uint16(aux_sym_http_version_token1),
	3097: uint16(anon_sym_LBRACE_LBRACE),
	3098: uint16(anon_sym_LT),
	3099: uint16(anon_sym_GT),
	3100: uint16(anon_sym_AT2),
	3101: uint16(anon_sym_LT2),
	3102: uint16(aux_sym_raw_body_token1),
	3103: uint16(aux_sym__blank_line_token1),
	3104: uint16(2),
	3105: uint16(420),
	3106: uint16(7),
	3108: uint16(aux_sym_request_separator_token1),
	3109: uint16(aux_sym_res_redirect_token1),
	3110: uint16(aux_sym_xml_body_token1),
	3111: uint16(aux_sym_json_body_token1),
	3112: uint16(aux_sym_graphql_data_token1),
	3113: uint16(anon_sym_DASH_DASH),
	3114: uint16(422),
	3115: uint16(14),
	3116: uint16(aux_sym_WORD_CHAR_token1),
	3117: uint16(aux_sym_PUNCTUATION_token1),
	3118: uint16(aux_sym_WS_token1),
	3119: uint16(aux_sym_NL_token1),
	3120: uint16(aux_sym_COMMENT_PREFIX_token1),
	3121: uint16(sym_method),
	3122: uint16(aux_sym_http_version_token1),
	3123: uint16(anon_sym_LBRACE_LBRACE),
	3124: uint16(anon_sym_LT),
	3125: uint16(anon_sym_GT),
	3126: uint16(anon_sym_AT2),
	3127: uint16(anon_sym_LT2),
	3128: uint16(aux_sym_raw_body_token1),
	3129: uint16(aux_sym__blank_line_token1),
	3130: uint16(2),
	3131: uint16(424),
	3132: uint16(7),
	3134: uint16(aux_sym_request_separator_token1),
	3135: uint16(aux_sym_res_redirect_token1),
	3136: uint16(aux_sym_xml_body_token1),
	3137: uint16(aux_sym_json_body_token1),
	3138: uint16(aux_sym_graphql_data_token1),
	3139: uint16(anon_sym_DASH_DASH),
	3140: uint16(426),
	3141: uint16(14),
	3142: uint16(aux_sym_WORD_CHAR_token1),
	3143: uint16(aux_sym_PUNCTUATION_token1),
	3144: uint16(aux_sym_WS_token1),
	3145: uint16(aux_sym_NL_token1),
	3146: uint16(aux_sym_COMMENT_PREFIX_token1),
	3147: uint16(sym_method),
	3148: uint16(aux_sym_http_version_token1),
	3149: uint16(anon_sym_LBRACE_LBRACE),
	3150: uint16(anon_sym_LT),
	3151: uint16(anon_sym_GT),
	3152: uint16(anon_sym_AT2),
	3153: uint16(anon_sym_LT2),
	3154: uint16(aux_sym_raw_body_token1),
	3155: uint16(aux_sym__blank_line_token1),
	3156: uint16(2),
	3157: uint16(302),
	3158: uint16(7),
	3160: uint16(aux_sym_request_separator_token1),
	3161: uint16(aux_sym_res_redirect_token1),
	3162: uint16(aux_sym_xml_body_token1),
	3163: uint16(aux_sym_json_body_token1),
	3164: uint16(aux_sym_graphql_data_token1),
	3165: uint16(anon_sym_DASH_DASH),
	3166: uint16(304),
	3167: uint16(14),
	3168: uint16(aux_sym_WORD_CHAR_token1),
	3169: uint16(aux_sym_PUNCTUATION_token1),
	3170: uint16(aux_sym_WS_token1),
	3171: uint16(aux_sym_NL_token1),
	3172: uint16(aux_sym_COMMENT_PREFIX_token1),
	3173: uint16(sym_method),
	3174: uint16(aux_sym_http_version_token1),
	3175: uint16(anon_sym_LBRACE_LBRACE),
	3176: uint16(anon_sym_LT),
	3177: uint16(anon_sym_GT),
	3178: uint16(anon_sym_AT2),
	3179: uint16(anon_sym_LT2),
	3180: uint16(aux_sym_raw_body_token1),
	3181: uint16(aux_sym__blank_line_token1),
	3182: uint16(2),
	3183: uint16(428),
	3184: uint16(7),
	3186: uint16(aux_sym_request_separator_token1),
	3187: uint16(aux_sym_res_redirect_token1),
	3188: uint16(aux_sym_xml_body_token1),
	3189: uint16(aux_sym_json_body_token1),
	3190: uint16(aux_sym_graphql_data_token1),
	3191: uint16(anon_sym_DASH_DASH),
	3192: uint16(430),
	3193: uint16(14),
	3194: uint16(aux_sym_WORD_CHAR_token1),
	3195: uint16(aux_sym_PUNCTUATION_token1),
	3196: uint16(aux_sym_WS_token1),
	3197: uint16(aux_sym_NL_token1),
	3198: uint16(aux_sym_COMMENT_PREFIX_token1),
	3199: uint16(sym_method),
	3200: uint16(aux_sym_http_version_token1),
	3201: uint16(anon_sym_LBRACE_LBRACE),
	3202: uint16(anon_sym_LT),
	3203: uint16(anon_sym_GT),
	3204: uint16(anon_sym_AT2),
	3205: uint16(anon_sym_LT2),
	3206: uint16(aux_sym_raw_body_token1),
	3207: uint16(aux_sym__blank_line_token1),
	3208: uint16(11),
	3209: uint16(436),
	3210: uint16(1),
	3211: uint16(aux_sym_WS_token1),
	3212: uint16(438),
	3213: uint16(1),
	3214: uint16(aux_sym_COMMENT_PREFIX_token1),
	3215: uint16(440),
	3216: uint16(1),
	3217: uint16(sym_header_entity),
	3218: uint16(442),
	3219: uint16(1),
	3220: uint16(aux_sym__blank_line_token1),
	3221: uint16(86),
	3222: uint16(1),
	3223: uint16(sym_header),
	3224: uint16(92),
	3225: uint16(1),
	3226: uint16(sym__var_comment),
	3227: uint16(95),
	3228: uint16(1),
	3229: uint16(sym__plain_comment),
	3230: uint16(12),
	3231: uint16(2),
	3232: uint16(sym__blank_line),
	3233: uint16(aux_sym___body_repeat1),
	3234: uint16(77),
	3235: uint16(2),
	3236: uint16(sym_comment),
	3237: uint16(aux_sym_request_repeat1),
	3238: uint16(434),
	3239: uint16(3),
	3240: uint16(aux_sym_WORD_CHAR_token1),
	3241: uint16(aux_sym_PUNCTUATION_token1),
	3242: uint16(sym_method),
	3243: uint16(432),
	3244: uint16(6),
	3246: uint16(aux_sym_request_separator_token1),
	3247: uint16(aux_sym_http_version_token1),
	3248: uint16(anon_sym_LBRACE_LBRACE),
	3249: uint16(anon_sym_LT),
	3250: uint16(anon_sym_AT2),
	3251: uint16(11),
	3252: uint16(436),
	3253: uint16(1),
	3254: uint16(aux_sym_WS_token1),
	3255: uint16(438),
	3256: uint16(1),
	3257: uint16(aux_sym_COMMENT_PREFIX_token1),
	3258: uint16(440),
	3259: uint16(1),
	3260: uint16(sym_header_entity),
	3261: uint16(448),
	3262: uint16(1),
	3263: uint16(aux_sym__blank_line_token1),
	3264: uint16(86),
	3265: uint16(1),
	3266: uint16(sym_header),
	3267: uint16(92),
	3268: uint16(1),
	3269: uint16(sym__var_comment),
	3270: uint16(95),
	3271: uint16(1),
	3272: uint16(sym__plain_comment),
	3273: uint16(6),
	3274: uint16(2),
	3275: uint16(sym__blank_line),
	3276: uint16(aux_sym___body_repeat1),
	3277: uint16(75),
	3278: uint16(2),
	3279: uint16(sym_comment),
	3280: uint16(aux_sym_request_repeat1),
	3281: uint16(446),
	3282: uint16(3),
	3283: uint16(aux_sym_WORD_CHAR_token1),
	3284: uint16(aux_sym_PUNCTUATION_token1),
	3285: uint16(sym_method),
	3286: uint16(444),
	3287: uint16(6),
	3289: uint16(aux_sym_request_separator_token1),
	3290: uint16(aux_sym_http_version_token1),
	3291: uint16(anon_sym_LBRACE_LBRACE),
	3292: uint16(anon_sym_LT),
	3293: uint16(anon_sym_AT2),
	3294: uint16(11),
	3295: uint16(436),
	3296: uint16(1),
	3297: uint16(aux_sym_WS_token1),
	3298: uint16(438),
	3299: uint16(1),
	3300: uint16(aux_sym_COMMENT_PREFIX_token1),
	3301: uint16(440),
	3302: uint16(1),
	3303: uint16(sym_header_entity),
	3304: uint16(454),
	3305: uint16(1),
	3306: uint16(aux_sym__blank_line_token1),
	3307: uint16(86),
	3308: uint16(1),
	3309: uint16(sym_header),
	3310: uint16(92),
	3311: uint16(1),
	3312: uint16(sym__var_comment),
	3313: uint16(95),
	3314: uint16(1),
	3315: uint16(sym__plain_comment),
	3316: uint16(4),
	3317: uint16(2),
	3318: uint16(sym__blank_line),
	3319: uint16(aux_sym___body_repeat1),
	3320: uint16(77),
	3321: uint16(2),
	3322: uint16(sym_comment),
	3323: uint16(aux_sym_request_repeat1),
	3324: uint16(452),
	3325: uint16(3),
	3326: uint16(aux_sym_WORD_CHAR_token1),
	3327: uint16(aux_sym_PUNCTUATION_token1),
	3328: uint16(sym_method),
	3329: uint16(450),
	3330: uint16(6),
	3332: uint16(aux_sym_request_separator_token1),
	3333: uint16(aux_sym_http_version_token1),
	3334: uint16(anon_sym_LBRACE_LBRACE),
	3335: uint16(anon_sym_LT),
	3336: uint16(anon_sym_AT2),
	3337: uint16(11),
	3338: uint16(436),
	3339: uint16(1),
	3340: uint16(aux_sym_WS_token1),
	3341: uint16(438),
	3342: uint16(1),
	3343: uint16(aux_sym_COMMENT_PREFIX_token1),
	3344: uint16(440),
	3345: uint16(1),
	3346: uint16(sym_header_entity),
	3347: uint16(460),
	3348: uint16(1),
	3349: uint16(aux_sym__blank_line_token1),
	3350: uint16(86),
	3351: uint16(1),
	3352: uint16(sym_header),
	3353: uint16(92),
	3354: uint16(1),
	3355: uint16(sym__var_comment),
	3356: uint16(95),
	3357: uint16(1),
	3358: uint16(sym__plain_comment),
	3359: uint16(5),
	3360: uint16(2),
	3361: uint16(sym__blank_line),
	3362: uint16(aux_sym___body_repeat1),
	3363: uint16(69),
	3364: uint16(2),
	3365: uint16(sym_comment),
	3366: uint16(aux_sym_request_repeat1),
	3367: uint16(458),
	3368: uint16(3),
	3369: uint16(aux_sym_WORD_CHAR_token1),
	3370: uint16(aux_sym_PUNCTUATION_token1),
	3371: uint16(sym_method),
	3372: uint16(456),
	3373: uint16(6),
	3375: uint16(aux_sym_request_separator_token1),
	3376: uint16(aux_sym_http_version_token1),
	3377: uint16(anon_sym_LBRACE_LBRACE),
	3378: uint16(anon_sym_LT),
	3379: uint16(anon_sym_AT2),
	3380: uint16(11),
	3381: uint16(436),
	3382: uint16(1),
	3383: uint16(aux_sym_WS_token1),
	3384: uint16(438),
	3385: uint16(1),
	3386: uint16(aux_sym_COMMENT_PREFIX_token1),
	3387: uint16(440),
	3388: uint16(1),
	3389: uint16(sym_header_entity),
	3390: uint16(466),
	3391: uint16(1),
	3392: uint16(aux_sym__blank_line_token1),
	3393: uint16(86),
	3394: uint16(1),
	3395: uint16(sym_header),
	3396: uint16(92),
	3397: uint16(1),
	3398: uint16(sym__var_comment),
	3399: uint16(95),
	3400: uint16(1),
	3401: uint16(sym__plain_comment),
	3402: uint16(7),
	3403: uint16(2),
	3404: uint16(sym__blank_line),
	3405: uint16(aux_sym___body_repeat1),
	3406: uint16(77),
	3407: uint16(2),
	3408: uint16(sym_comment),
	3409: uint16(aux_sym_request_repeat1),
	3410: uint16(464),
	3411: uint16(3),
	3412: uint16(aux_sym_WORD_CHAR_token1),
	3413: uint16(aux_sym_PUNCTUATION_token1),
	3414: uint16(sym_method),
	3415: uint16(462),
	3416: uint16(6),
	3418: uint16(aux_sym_request_separator_token1),
	3419: uint16(aux_sym_http_version_token1),
	3420: uint16(anon_sym_LBRACE_LBRACE),
	3421: uint16(anon_sym_LT),
	3422: uint16(anon_sym_AT2),
	3423: uint16(11),
	3424: uint16(436),
	3425: uint16(1),
	3426: uint16(aux_sym_WS_token1),
	3427: uint16(438),
	3428: uint16(1),
	3429: uint16(aux_sym_COMMENT_PREFIX_token1),
	3430: uint16(440),
	3431: uint16(1),
	3432: uint16(sym_header_entity),
	3433: uint16(472),
	3434: uint16(1),
	3435: uint16(aux_sym__blank_line_token1),
	3436: uint16(86),
	3437: uint16(1),
	3438: uint16(sym_header),
	3439: uint16(92),
	3440: uint16(1),
	3441: uint16(sym__var_comment),
	3442: uint16(95),
	3443: uint16(1),
	3444: uint16(sym__plain_comment),
	3445: uint16(8),
	3446: uint16(2),
	3447: uint16(sym__blank_line),
	3448: uint16(aux_sym___body_repeat1),
	3449: uint16(73),
	3450: uint16(2),
	3451: uint16(sym_comment),
	3452: uint16(aux_sym_request_repeat1),
	3453: uint16(470),
	3454: uint16(3),
	3455: uint16(aux_sym_WORD_CHAR_token1),
	3456: uint16(aux_sym_PUNCTUATION_token1),
	3457: uint16(sym_method),
	3458: uint16(468),
	3459: uint16(6),
	3461: uint16(aux_sym_request_separator_token1),
	3462: uint16(aux_sym_http_version_token1),
	3463: uint16(anon_sym_LBRACE_LBRACE),
	3464: uint16(anon_sym_LT),
	3465: uint16(anon_sym_AT2),
	3466: uint16(11),
	3467: uint16(436),
	3468: uint16(1),
	3469: uint16(aux_sym_WS_token1),
	3470: uint16(438),
	3471: uint16(1),
	3472: uint16(aux_sym_COMMENT_PREFIX_token1),
	3473: uint16(440),
	3474: uint16(1),
	3475: uint16(sym_header_entity),
	3476: uint16(478),
	3477: uint16(1),
	3478: uint16(aux_sym__blank_line_token1),
	3479: uint16(86),
	3480: uint16(1),
	3481: uint16(sym_header),
	3482: uint16(92),
	3483: uint16(1),
	3484: uint16(sym__var_comment),
	3485: uint16(95),
	3486: uint16(1),
	3487: uint16(sym__plain_comment),
	3488: uint16(2),
	3489: uint16(2),
	3490: uint16(sym__blank_line),
	3491: uint16(aux_sym___body_repeat1),
	3492: uint16(77),
	3493: uint16(2),
	3494: uint16(sym_comment),
	3495: uint16(aux_sym_request_repeat1),
	3496: uint16(476),
	3497: uint16(3),
	3498: uint16(aux_sym_WORD_CHAR_token1),
	3499: uint16(aux_sym_PUNCTUATION_token1),
	3500: uint16(sym_method),
	3501: uint16(474),
	3502: uint16(6),
	3504: uint16(aux_sym_request_separator_token1),
	3505: uint16(aux_sym_http_version_token1),
	3506: uint16(anon_sym_LBRACE_LBRACE),
	3507: uint16(anon_sym_LT),
	3508: uint16(anon_sym_AT2),
	3509: uint16(11),
	3510: uint16(436),
	3511: uint16(1),
	3512: uint16(aux_sym_WS_token1),
	3513: uint16(438),
	3514: uint16(1),
	3515: uint16(aux_sym_COMMENT_PREFIX_token1),
	3516: uint16(440),
	3517: uint16(1),
	3518: uint16(sym_header_entity),
	3519: uint16(484),
	3520: uint16(1),
	3521: uint16(aux_sym__blank_line_token1),
	3522: uint16(86),
	3523: uint16(1),
	3524: uint16(sym_header),
	3525: uint16(92),
	3526: uint16(1),
	3527: uint16(sym__var_comment),
	3528: uint16(95),
	3529: uint16(1),
	3530: uint16(sym__plain_comment),
	3531: uint16(11),
	3532: uint16(2),
	3533: uint16(sym__blank_line),
	3534: uint16(aux_sym___body_repeat1),
	3535: uint16(71),
	3536: uint16(2),
	3537: uint16(sym_comment),
	3538: uint16(aux_sym_request_repeat1),
	3539: uint16(482),
	3540: uint16(3),
	3541: uint16(aux_sym_WORD_CHAR_token1),
	3542: uint16(aux_sym_PUNCTUATION_token1),
	3543: uint16(sym_method),
	3544: uint16(480),
	3545: uint16(6),
	3547: uint16(aux_sym_request_separator_token1),
	3548: uint16(aux_sym_http_version_token1),
	3549: uint16(anon_sym_LBRACE_LBRACE),
	3550: uint16(anon_sym_LT),
	3551: uint16(anon_sym_AT2),
	3552: uint16(8),
	3553: uint16(490),
	3554: uint16(1),
	3555: uint16(aux_sym_COMMENT_PREFIX_token1),
	3556: uint16(493),
	3557: uint16(1),
	3558: uint16(sym_header_entity),
	3559: uint16(86),
	3560: uint16(1),
	3561: uint16(sym_header),
	3562: uint16(92),
	3563: uint16(1),
	3564: uint16(sym__var_comment),
	3565: uint16(95),
	3566: uint16(1),
	3567: uint16(sym__plain_comment),
	3568: uint16(77),
	3569: uint16(2),
	3570: uint16(sym_comment),
	3571: uint16(aux_sym_request_repeat1),
	3572: uint16(488),
	3573: uint16(4),
	3574: uint16(aux_sym_WORD_CHAR_token1),
	3575: uint16(aux_sym_PUNCTUATION_token1),
	3576: uint16(sym_method),
	3577: uint16(aux_sym__blank_line_token1),
	3578: uint16(486),
	3579: uint16(7),
	3581: uint16(aux_sym_WS_token1),
	3582: uint16(aux_sym_request_separator_token1),
	3583: uint16(aux_sym_http_version_token1),
	3584: uint16(anon_sym_LBRACE_LBRACE),
	3585: uint16(anon_sym_LT),
	3586: uint16(anon_sym_AT2),
	3587: uint16(8),
	3588: uint16(436),
	3589: uint16(1),
	3590: uint16(aux_sym_WS_token1),
	3591: uint16(440),
	3592: uint16(1),
	3593: uint16(sym_header_entity),
	3594: uint16(500),
	3595: uint16(1),
	3596: uint16(aux_sym__blank_line_token1),
	3597: uint16(82),
	3598: uint16(1),
	3599: uint16(aux_sym_response_repeat1),
	3600: uint16(87),
	3601: uint16(1),
	3602: uint16(sym_header),
	3603: uint16(13),
	3604: uint16(2),
	3605: uint16(sym__blank_line),
	3606: uint16(aux_sym___body_repeat1),
	3607: uint16(498),
	3608: uint16(4),
	3609: uint16(aux_sym_WORD_CHAR_token1),
	3610: uint16(aux_sym_PUNCTUATION_token1),
	3611: uint16(aux_sym_COMMENT_PREFIX_token1),
	3612: uint16(sym_method),
	3613: uint16(496),
	3614: uint16(6),
	3616: uint16(aux_sym_request_separator_token1),
	3617: uint16(aux_sym_http_version_token1),
	3618: uint16(anon_sym_LBRACE_LBRACE),
	3619: uint16(anon_sym_LT),
	3620: uint16(anon_sym_AT2),
	3621: uint16(8),
	3622: uint16(436),
	3623: uint16(1),
	3624: uint16(aux_sym_WS_token1),
	3625: uint16(440),
	3626: uint16(1),
	3627: uint16(sym_header_entity),
	3628: uint16(506),
	3629: uint16(1),
	3630: uint16(aux_sym__blank_line_token1),
	3631: uint16(80),
	3632: uint16(1),
	3633: uint16(aux_sym_response_repeat1),
	3634: uint16(87),
	3635: uint16(1),
	3636: uint16(sym_header),
	3637: uint16(3),
	3638: uint16(2),
	3639: uint16(sym__blank_line),
	3640: uint16(aux_sym___body_repeat1),
	3641: uint16(504),
	3642: uint16(4),
	3643: uint16(aux_sym_WORD_CHAR_token1),
	3644: uint16(aux_sym_PUNCTUATION_token1),
	3645: uint16(aux_sym_COMMENT_PREFIX_token1),
	3646: uint16(sym_method),
	3647: uint16(502),
	3648: uint16(6),
	3650: uint16(aux_sym_request_separator_token1),
	3651: uint16(aux_sym_http_version_token1),
	3652: uint16(anon_sym_LBRACE_LBRACE),
	3653: uint16(anon_sym_LT),
	3654: uint16(anon_sym_AT2),
	3655: uint16(8),
	3656: uint16(436),
	3657: uint16(1),
	3658: uint16(aux_sym_WS_token1),
	3659: uint16(440),
	3660: uint16(1),
	3661: uint16(sym_header_entity),
	3662: uint16(512),
	3663: uint16(1),
	3664: uint16(aux_sym__blank_line_token1),
	3665: uint16(82),
	3666: uint16(1),
	3667: uint16(aux_sym_response_repeat1),
	3668: uint16(87),
	3669: uint16(1),
	3670: uint16(sym_header),
	3671: uint16(9),
	3672: uint16(2),
	3673: uint16(sym__blank_line),
	3674: uint16(aux_sym___body_repeat1),
	3675: uint16(510),
	3676: uint16(4),
	3677: uint16(aux_sym_WORD_CHAR_token1),
	3678: uint16(aux_sym_PUNCTUATION_token1),
	3679: uint16(aux_sym_COMMENT_PREFIX_token1),
	3680: uint16(sym_method),
	3681: uint16(508),
	3682: uint16(6),
	3684: uint16(aux_sym_request_separator_token1),
	3685: uint16(aux_sym_http_version_token1),
	3686: uint16(anon_sym_LBRACE_LBRACE),
	3687: uint16(anon_sym_LT),
	3688: uint16(anon_sym_AT2),
	3689: uint16(8),
	3690: uint16(436),
	3691: uint16(1),
	3692: uint16(aux_sym_WS_token1),
	3693: uint16(440),
	3694: uint16(1),
	3695: uint16(sym_header_entity),
	3696: uint16(514),
	3697: uint16(1),
	3698: uint16(aux_sym__blank_line_token1),
	3699: uint16(78),
	3700: uint16(1),
	3701: uint16(aux_sym_response_repeat1),
	3702: uint16(87),
	3703: uint16(1),
	3704: uint16(sym_header),
	3705: uint16(10),
	3706: uint16(2),
	3707: uint16(sym__blank_line),
	3708: uint16(aux_sym___body_repeat1),
	3709: uint16(55),
	3710: uint16(4),
	3711: uint16(aux_sym_WORD_CHAR_token1),
	3712: uint16(aux_sym_PUNCTUATION_token1),
	3713: uint16(aux_sym_COMMENT_PREFIX_token1),
	3714: uint16(sym_method),
	3715: uint16(53),
	3716: uint16(6),
	3718: uint16(aux_sym_request_separator_token1),
	3719: uint16(aux_sym_http_version_token1),
	3720: uint16(anon_sym_LBRACE_LBRACE),
	3721: uint16(anon_sym_LT),
	3722: uint16(anon_sym_AT2),
	3723: uint16(5),
	3724: uint16(520),
	3725: uint16(1),
	3726: uint16(sym_header_entity),
	3727: uint16(82),
	3728: uint16(1),
	3729: uint16(aux_sym_response_repeat1),
	3730: uint16(87),
	3731: uint16(1),
	3732: uint16(sym_header),
	3733: uint16(518),
	3734: uint16(5),
	3735: uint16(aux_sym_WORD_CHAR_token1),
	3736: uint16(aux_sym_PUNCTUATION_token1),
	3737: uint16(aux_sym_COMMENT_PREFIX_token1),
	3738: uint16(sym_method),
	3739: uint16(aux_sym__blank_line_token1),
	3740: uint16(516),
	3741: uint16(7),
	3743: uint16(aux_sym_WS_token1),
	3744: uint16(aux_sym_request_separator_token1),
	3745: uint16(aux_sym_http_version_token1),
	3746: uint16(anon_sym_LBRACE_LBRACE),
	3747: uint16(anon_sym_LT),
	3748: uint16(anon_sym_AT2),
	3749: uint16(2),
	3750: uint16(525),
	3751: uint16(6),
	3752: uint16(aux_sym_WORD_CHAR_token1),
	3753: uint16(aux_sym_PUNCTUATION_token1),
	3754: uint16(aux_sym_COMMENT_PREFIX_token1),
	3755: uint16(sym_method),
	3756: uint16(sym_header_entity),
	3757: uint16(aux_sym__blank_line_token1),
	3758: uint16(523),
	3759: uint16(7),
	3761: uint16(aux_sym_WS_token1),
	3762: uint16(aux_sym_request_separator_token1),
	3763: uint16(aux_sym_http_version_token1),
	3764: uint16(anon_sym_LBRACE_LBRACE),
	3765: uint16(anon_sym_LT),
	3766: uint16(anon_sym_AT2),
	3767: uint16(2),
	3768: uint16(320),
	3769: uint16(6),
	3770: uint16(aux_sym_WORD_CHAR_token1),
	3771: uint16(aux_sym_PUNCTUATION_token1),
	3772: uint16(aux_sym_COMMENT_PREFIX_token1),
	3773: uint16(sym_method),
	3774: uint16(sym_header_entity),
	3775: uint16(aux_sym__blank_line_token1),
	3776: uint16(318),
	3777: uint16(7),
	3779: uint16(aux_sym_WS_token1),
	3780: uint16(aux_sym_request_separator_token1),
	3781: uint16(aux_sym_http_version_token1),
	3782: uint16(anon_sym_LBRACE_LBRACE),
	3783: uint16(anon_sym_LT),
	3784: uint16(anon_sym_AT2),
	3785: uint16(2),
	3786: uint16(529),
	3787: uint16(6),
	3788: uint16(aux_sym_WORD_CHAR_token1),
	3789: uint16(aux_sym_PUNCTUATION_token1),
	3790: uint16(aux_sym_COMMENT_PREFIX_token1),
	3791: uint16(sym_method),
	3792: uint16(sym_header_entity),
	3793: uint16(aux_sym__blank_line_token1),
	3794: uint16(527),
	3795: uint16(7),
	3797: uint16(aux_sym_WS_token1),
	3798: uint16(aux_sym_request_separator_token1),
	3799: uint16(aux_sym_http_version_token1),
	3800: uint16(anon_sym_LBRACE_LBRACE),
	3801: uint16(anon_sym_LT),
	3802: uint16(anon_sym_AT2),
	3803: uint16(2),
	3804: uint16(533),
	3805: uint16(6),
	3806: uint16(aux_sym_WORD_CHAR_token1),
	3807: uint16(aux_sym_PUNCTUATION_token1),
	3808: uint16(aux_sym_COMMENT_PREFIX_token1),
	3809: uint16(sym_method),
	3810: uint16(sym_header_entity),
	3811: uint16(aux_sym__blank_line_token1),
	3812: uint16(531),
	3813: uint16(7),
	3815: uint16(aux_sym_WS_token1),
	3816: uint16(aux_sym_request_separator_token1),
	3817: uint16(aux_sym_http_version_token1),
	3818: uint16(anon_sym_LBRACE_LBRACE),
	3819: uint16(anon_sym_LT),
	3820: uint16(anon_sym_AT2),
	3821: uint16(2),
	3822: uint16(537),
	3823: uint16(6),
	3824: uint16(aux_sym_WORD_CHAR_token1),
	3825: uint16(aux_sym_PUNCTUATION_token1),
	3826: uint16(aux_sym_COMMENT_PREFIX_token1),
	3827: uint16(sym_method),
	3828: uint16(sym_header_entity),
	3829: uint16(aux_sym__blank_line_token1),
	3830: uint16(535),
	3831: uint16(7),
	3833: uint16(aux_sym_WS_token1),
	3834: uint16(aux_sym_request_separator_token1),
	3835: uint16(aux_sym_http_version_token1),
	3836: uint16(anon_sym_LBRACE_LBRACE),
	3837: uint16(anon_sym_LT),
	3838: uint16(anon_sym_AT2),
	3839: uint16(2),
	3840: uint16(541),
	3841: uint16(6),
	3842: uint16(aux_sym_WORD_CHAR_token1),
	3843: uint16(aux_sym_PUNCTUATION_token1),
	3844: uint16(aux_sym_COMMENT_PREFIX_token1),
	3845: uint16(sym_method),
	3846: uint16(sym_header_entity),
	3847: uint16(aux_sym__blank_line_token1),
	3848: uint16(539),
	3849: uint16(7),
	3851: uint16(aux_sym_WS_token1),
	3852: uint16(aux_sym_request_separator_token1),
	3853: uint16(aux_sym_http_version_token1),
	3854: uint16(anon_sym_LBRACE_LBRACE),
	3855: uint16(anon_sym_LT),
	3856: uint16(anon_sym_AT2),
	3857: uint16(2),
	3858: uint16(545),
	3859: uint16(6),
	3860: uint16(aux_sym_WORD_CHAR_token1),
	3861: uint16(aux_sym_PUNCTUATION_token1),
	3862: uint16(aux_sym_COMMENT_PREFIX_token1),
	3863: uint16(sym_method),
	3864: uint16(sym_header_entity),
	3865: uint16(aux_sym__blank_line_token1),
	3866: uint16(543),
	3867: uint16(7),
	3869: uint16(aux_sym_WS_token1),
	3870: uint16(aux_sym_request_separator_token1),
	3871: uint16(aux_sym_http_version_token1),
	3872: uint16(anon_sym_LBRACE_LBRACE),
	3873: uint16(anon_sym_LT),
	3874: uint16(anon_sym_AT2),
	3875: uint16(2),
	3876: uint16(549),
	3877: uint16(6),
	3878: uint16(aux_sym_WORD_CHAR_token1),
	3879: uint16(aux_sym_PUNCTUATION_token1),
	3880: uint16(aux_sym_COMMENT_PREFIX_token1),
	3881: uint16(sym_method),
	3882: uint16(sym_header_entity),
	3883: uint16(aux_sym__blank_line_token1),
	3884: uint16(547),
	3885: uint16(7),
	3887: uint16(aux_sym_WS_token1),
	3888: uint16(aux_sym_request_separator_token1),
	3889: uint16(aux_sym_http_version_token1),
	3890: uint16(anon_sym_LBRACE_LBRACE),
	3891: uint16(anon_sym_LT),
	3892: uint16(anon_sym_AT2),
	3893: uint16(2),
	3894: uint16(553),
	3895: uint16(6),
	3896: uint16(aux_sym_WORD_CHAR_token1),
	3897: uint16(aux_sym_PUNCTUATION_token1),
	3898: uint16(aux_sym_COMMENT_PREFIX_token1),
	3899: uint16(sym_method),
	3900: uint16(sym_header_entity),
	3901: uint16(aux_sym__blank_line_token1),
	3902: uint16(551),
	3903: uint16(7),
	3905: uint16(aux_sym_WS_token1),
	3906: uint16(aux_sym_request_separator_token1),
	3907: uint16(aux_sym_http_version_token1),
	3908: uint16(anon_sym_LBRACE_LBRACE),
	3909: uint16(anon_sym_LT),
	3910: uint16(anon_sym_AT2),
	3911: uint16(2),
	3912: uint16(333),
	3913: uint16(6),
	3914: uint16(aux_sym_WORD_CHAR_token1),
	3915: uint16(aux_sym_PUNCTUATION_token1),
	3916: uint16(aux_sym_COMMENT_PREFIX_token1),
	3917: uint16(sym_method),
	3918: uint16(sym_header_entity),
	3919: uint16(aux_sym__blank_line_token1),
	3920: uint16(331),
	3921: uint16(7),
	3923: uint16(aux_sym_WS_token1),
	3924: uint16(aux_sym_request_separator_token1),
	3925: uint16(aux_sym_http_version_token1),
	3926: uint16(anon_sym_LBRACE_LBRACE),
	3927: uint16(anon_sym_LT),
	3928: uint16(anon_sym_AT2),
	3929: uint16(2),
	3930: uint16(354),
	3931: uint16(6),
	3932: uint16(aux_sym_WORD_CHAR_token1),
	3933: uint16(aux_sym_PUNCTUATION_token1),
	3934: uint16(aux_sym_COMMENT_PREFIX_token1),
	3935: uint16(sym_method),
	3936: uint16(sym_header_entity),
	3937: uint16(aux_sym__blank_line_token1),
	3938: uint16(352),
	3939: uint16(7),
	3941: uint16(aux_sym_WS_token1),
	3942: uint16(aux_sym_request_separator_token1),
	3943: uint16(aux_sym_http_version_token1),
	3944: uint16(anon_sym_LBRACE_LBRACE),
	3945: uint16(anon_sym_LT),
	3946: uint16(anon_sym_AT2),
	3947: uint16(2),
	3948: uint16(358),
	3949: uint16(6),
	3950: uint16(aux_sym_WORD_CHAR_token1),
	3951: uint16(aux_sym_PUNCTUATION_token1),
	3952: uint16(aux_sym_COMMENT_PREFIX_token1),
	3953: uint16(sym_method),
	3954: uint16(sym_header_entity),
	3955: uint16(aux_sym__blank_line_token1),
	3956: uint16(356),
	3957: uint16(7),
	3959: uint16(aux_sym_WS_token1),
	3960: uint16(aux_sym_request_separator_token1),
	3961: uint16(aux_sym_http_version_token1),
	3962: uint16(anon_sym_LBRACE_LBRACE),
	3963: uint16(anon_sym_LT),
	3964: uint16(anon_sym_AT2),
	3965: uint16(2),
	3966: uint16(329),
	3967: uint16(6),
	3968: uint16(aux_sym_WORD_CHAR_token1),
	3969: uint16(aux_sym_PUNCTUATION_token1),
	3970: uint16(aux_sym_COMMENT_PREFIX_token1),
	3971: uint16(sym_method),
	3972: uint16(sym_header_entity),
	3973: uint16(aux_sym__blank_line_token1),
	3974: uint16(327),
	3975: uint16(7),
	3977: uint16(aux_sym_WS_token1),
	3978: uint16(aux_sym_request_separator_token1),
	3979: uint16(aux_sym_http_version_token1),
	3980: uint16(anon_sym_LBRACE_LBRACE),
	3981: uint16(anon_sym_LT),
	3982: uint16(anon_sym_AT2),
	3983: uint16(2),
	3984: uint16(362),
	3985: uint16(6),
	3986: uint16(aux_sym_WORD_CHAR_token1),
	3987: uint16(aux_sym_PUNCTUATION_token1),
	3988: uint16(aux_sym_COMMENT_PREFIX_token1),
	3989: uint16(sym_method),
	3990: uint16(sym_header_entity),
	3991: uint16(aux_sym__blank_line_token1),
	3992: uint16(360),
	3993: uint16(7),
	3995: uint16(aux_sym_WS_token1),
	3996: uint16(aux_sym_request_separator_token1),
	3997: uint16(aux_sym_http_version_token1),
	3998: uint16(anon_sym_LBRACE_LBRACE),
	3999: uint16(anon_sym_LT),
	4000: uint16(anon_sym_AT2),
	4001: uint16(2),
	4002: uint16(557),
	4003: uint16(4),
	4004: uint16(aux_sym_WORD_CHAR_token1),
	4005: uint16(aux_sym_PUNCTUATION_token1),
	4006: uint16(aux_sym_COMMENT_PREFIX_token1),
	4007: uint16(aux_sym__blank_line_token1),
	4008: uint16(555),
	4009: uint16(8),
	4011: uint16(aux_sym_WS_token1),
	4012: uint16(aux_sym_request_separator_token1),
	4013: uint16(sym_method),
	4014: uint16(aux_sym_http_version_token1),
	4015: uint16(anon_sym_LBRACE_LBRACE),
	4016: uint16(anon_sym_LT),
	4017: uint16(anon_sym_AT2),
	4018: uint16(2),
	4019: uint16(561),
	4020: uint16(4),
	4021: uint16(aux_sym_WORD_CHAR_token1),
	4022: uint16(aux_sym_PUNCTUATION_token1),
	4023: uint16(aux_sym_COMMENT_PREFIX_token1),
	4024: uint16(aux_sym__blank_line_token1),
	4025: uint16(559),
	4026: uint16(8),
	4028: uint16(aux_sym_WS_token1),
	4029: uint16(aux_sym_request_separator_token1),
	4030: uint16(sym_method),
	4031: uint16(aux_sym_http_version_token1),
	4032: uint16(anon_sym_LBRACE_LBRACE),
	4033: uint16(anon_sym_LT),
	4034: uint16(anon_sym_AT2),
	4035: uint16(2),
	4036: uint16(565),
	4037: uint16(4),
	4038: uint16(aux_sym_WORD_CHAR_token1),
	4039: uint16(aux_sym_PUNCTUATION_token1),
	4040: uint16(aux_sym_COMMENT_PREFIX_token1),
	4041: uint16(aux_sym__blank_line_token1),
	4042: uint16(563),
	4043: uint16(8),
	4045: uint16(aux_sym_WS_token1),
	4046: uint16(aux_sym_request_separator_token1),
	4047: uint16(sym_method),
	4048: uint16(aux_sym_http_version_token1),
	4049: uint16(anon_sym_LBRACE_LBRACE),
	4050: uint16(anon_sym_LT),
	4051: uint16(anon_sym_AT2),
	4052: uint16(2),
	4053: uint16(569),
	4054: uint16(4),
	4055: uint16(aux_sym_WORD_CHAR_token1),
	4056: uint16(aux_sym_PUNCTUATION_token1),
	4057: uint16(aux_sym_COMMENT_PREFIX_token1),
	4058: uint16(aux_sym__blank_line_token1),
	4059: uint16(567),
	4060: uint16(8),
	4062: uint16(aux_sym_WS_token1),
	4063: uint16(aux_sym_request_separator_token1),
	4064: uint16(sym_method),
	4065: uint16(aux_sym_http_version_token1),
	4066: uint16(anon_sym_LBRACE_LBRACE),
	4067: uint16(anon_sym_LT),
	4068: uint16(anon_sym_AT2),
	4069: uint16(2),
	4070: uint16(354),
	4071: uint16(4),
	4072: uint16(aux_sym_WORD_CHAR_token1),
	4073: uint16(aux_sym_PUNCTUATION_token1),
	4074: uint16(aux_sym_COMMENT_PREFIX_token1),
	4075: uint16(aux_sym__blank_line_token1),
	4076: uint16(352),
	4077: uint16(8),
	4079: uint16(aux_sym_WS_token1),
	4080: uint16(aux_sym_request_separator_token1),
	4081: uint16(sym_method),
	4082: uint16(aux_sym_http_version_token1),
	4083: uint16(anon_sym_LBRACE_LBRACE),
	4084: uint16(anon_sym_LT),
	4085: uint16(anon_sym_AT2),
	4086: uint16(2),
	4087: uint16(362),
	4088: uint16(4),
	4089: uint16(aux_sym_WORD_CHAR_token1),
	4090: uint16(aux_sym_PUNCTUATION_token1),
	4091: uint16(aux_sym_COMMENT_PREFIX_token1),
	4092: uint16(aux_sym__blank_line_token1),
	4093: uint16(360),
	4094: uint16(8),
	4096: uint16(aux_sym_WS_token1),
	4097: uint16(aux_sym_request_separator_token1),
	4098: uint16(sym_method),
	4099: uint16(aux_sym_http_version_token1),
	4100: uint16(anon_sym_LBRACE_LBRACE),
	4101: uint16(anon_sym_LT),
	4102: uint16(anon_sym_AT2),
	4103: uint16(2),
	4104: uint16(573),
	4105: uint16(4),
	4106: uint16(aux_sym_WORD_CHAR_token1),
	4107: uint16(aux_sym_PUNCTUATION_token1),
	4108: uint16(aux_sym_COMMENT_PREFIX_token1),
	4109: uint16(aux_sym__blank_line_token1),
	4110: uint16(571),
	4111: uint16(8),
	4113: uint16(aux_sym_WS_token1),
	4114: uint16(aux_sym_request_separator_token1),
	4115: uint16(sym_method),
	4116: uint16(aux_sym_http_version_token1),
	4117: uint16(anon_sym_LBRACE_LBRACE),
	4118: uint16(anon_sym_LT),
	4119: uint16(anon_sym_AT2),
	4120: uint16(2),
	4121: uint16(577),
	4122: uint16(4),
	4123: uint16(aux_sym_WORD_CHAR_token1),
	4124: uint16(aux_sym_PUNCTUATION_token1),
	4125: uint16(aux_sym_COMMENT_PREFIX_token1),
	4126: uint16(aux_sym__blank_line_token1),
	4127: uint16(575),
	4128: uint16(8),
	4130: uint16(aux_sym_WS_token1),
	4131: uint16(aux_sym_request_separator_token1),
	4132: uint16(sym_method),
	4133: uint16(aux_sym_http_version_token1),
	4134: uint16(anon_sym_LBRACE_LBRACE),
	4135: uint16(anon_sym_LT),
	4136: uint16(anon_sym_AT2),
	4137: uint16(2),
	4138: uint16(581),
	4139: uint16(4),
	4140: uint16(aux_sym_WORD_CHAR_token1),
	4141: uint16(aux_sym_PUNCTUATION_token1),
	4142: uint16(aux_sym_COMMENT_PREFIX_token1),
	4143: uint16(aux_sym__blank_line_token1),
	4144: uint16(579),
	4145: uint16(8),
	4147: uint16(aux_sym_WS_token1),
	4148: uint16(aux_sym_request_separator_token1),
	4149: uint16(sym_method),
	4150: uint16(aux_sym_http_version_token1),
	4151: uint16(anon_sym_LBRACE_LBRACE),
	4152: uint16(anon_sym_LT),
	4153: uint16(anon_sym_AT2),
	4154: uint16(2),
	4155: uint16(402),
	4156: uint16(4),
	4157: uint16(aux_sym_WORD_CHAR_token1),
	4158: uint16(aux_sym_PUNCTUATION_token1),
	4159: uint16(aux_sym_COMMENT_PREFIX_token1),
	4160: uint16(aux_sym__blank_line_token1),
	4161: uint16(400),
	4162: uint16(8),
	4164: uint16(aux_sym_WS_token1),
	4165: uint16(aux_sym_request_separator_token1),
	4166: uint16(sym_method),
	4167: uint16(aux_sym_http_version_token1),
	4168: uint16(anon_sym_LBRACE_LBRACE),
	4169: uint16(anon_sym_LT),
	4170: uint16(anon_sym_AT2),
	4171: uint16(2),
	4172: uint16(320),
	4173: uint16(4),
	4174: uint16(aux_sym_WORD_CHAR_token1),
	4175: uint16(aux_sym_PUNCTUATION_token1),
	4176: uint16(aux_sym_COMMENT_PREFIX_token1),
	4177: uint16(aux_sym__blank_line_token1),
	4178: uint16(318),
	4179: uint16(8),
	4181: uint16(aux_sym_WS_token1),
	4182: uint16(aux_sym_request_separator_token1),
	4183: uint16(sym_method),
	4184: uint16(aux_sym_http_version_token1),
	4185: uint16(anon_sym_LBRACE_LBRACE),
	4186: uint16(anon_sym_LT),
	4187: uint16(anon_sym_AT2),
	4188: uint16(2),
	4189: uint16(585),
	4190: uint16(4),
	4191: uint16(aux_sym_WORD_CHAR_token1),
	4192: uint16(aux_sym_PUNCTUATION_token1),
	4193: uint16(aux_sym_COMMENT_PREFIX_token1),
	4194: uint16(aux_sym__blank_line_token1),
	4195: uint16(583),
	4196: uint16(8),
	4198: uint16(aux_sym_WS_token1),
	4199: uint16(aux_sym_request_separator_token1),
	4200: uint16(sym_method),
	4201: uint16(aux_sym_http_version_token1),
	4202: uint16(anon_sym_LBRACE_LBRACE),
	4203: uint16(anon_sym_LT),
	4204: uint16(anon_sym_AT2),
	4205: uint16(2),
	4206: uint16(358),
	4207: uint16(4),
	4208: uint16(aux_sym_WORD_CHAR_token1),
	4209: uint16(aux_sym_PUNCTUATION_token1),
	4210: uint16(aux_sym_COMMENT_PREFIX_token1),
	4211: uint16(aux_sym__blank_line_token1),
	4212: uint16(356),
	4213: uint16(8),
	4215: uint16(aux_sym_WS_token1),
	4216: uint16(aux_sym_request_separator_token1),
	4217: uint16(sym_method),
	4218: uint16(aux_sym_http_version_token1),
	4219: uint16(anon_sym_LBRACE_LBRACE),
	4220: uint16(anon_sym_LT),
	4221: uint16(anon_sym_AT2),
	4222: uint16(2),
	4223: uint16(589),
	4224: uint16(4),
	4225: uint16(aux_sym_WORD_CHAR_token1),
	4226: uint16(aux_sym_PUNCTUATION_token1),
	4227: uint16(aux_sym_COMMENT_PREFIX_token1),
	4228: uint16(aux_sym__blank_line_token1),
	4229: uint16(587),
	4230: uint16(8),
	4232: uint16(aux_sym_WS_token1),
	4233: uint16(aux_sym_request_separator_token1),
	4234: uint16(sym_method),
	4235: uint16(aux_sym_http_version_token1),
	4236: uint16(anon_sym_LBRACE_LBRACE),
	4237: uint16(anon_sym_LT),
	4238: uint16(anon_sym_AT2),
	4239: uint16(2),
	4240: uint16(593),
	4241: uint16(4),
	4242: uint16(aux_sym_WORD_CHAR_token1),
	4243: uint16(aux_sym_PUNCTUATION_token1),
	4244: uint16(aux_sym_COMMENT_PREFIX_token1),
	4245: uint16(aux_sym__blank_line_token1),
	4246: uint16(591),
	4247: uint16(8),
	4249: uint16(aux_sym_WS_token1),
	4250: uint16(aux_sym_request_separator_token1),
	4251: uint16(sym_method),
	4252: uint16(aux_sym_http_version_token1),
	4253: uint16(anon_sym_LBRACE_LBRACE),
	4254: uint16(anon_sym_LT),
	4255: uint16(anon_sym_AT2),
	4256: uint16(2),
	4257: uint16(329),
	4258: uint16(4),
	4259: uint16(aux_sym_WORD_CHAR_token1),
	4260: uint16(aux_sym_PUNCTUATION_token1),
	4261: uint16(aux_sym_COMMENT_PREFIX_token1),
	4262: uint16(aux_sym__blank_line_token1),
	4263: uint16(327),
	4264: uint16(8),
	4266: uint16(aux_sym_WS_token1),
	4267: uint16(aux_sym_request_separator_token1),
	4268: uint16(sym_method),
	4269: uint16(aux_sym_http_version_token1),
	4270: uint16(anon_sym_LBRACE_LBRACE),
	4271: uint16(anon_sym_LT),
	4272: uint16(anon_sym_AT2),
	4273: uint16(2),
	4274: uint16(597),
	4275: uint16(4),
	4276: uint16(aux_sym_WORD_CHAR_token1),
	4277: uint16(aux_sym_PUNCTUATION_token1),
	4278: uint16(aux_sym_COMMENT_PREFIX_token1),
	4279: uint16(aux_sym__blank_line_token1),
	4280: uint16(595),
	4281: uint16(8),
	4283: uint16(aux_sym_WS_token1),
	4284: uint16(aux_sym_request_separator_token1),
	4285: uint16(sym_method),
	4286: uint16(aux_sym_http_version_token1),
	4287: uint16(anon_sym_LBRACE_LBRACE),
	4288: uint16(anon_sym_LT),
	4289: uint16(anon_sym_AT2),
	4290: uint16(2),
	4291: uint16(601),
	4292: uint16(4),
	4293: uint16(aux_sym_WORD_CHAR_token1),
	4294: uint16(aux_sym_PUNCTUATION_token1),
	4295: uint16(aux_sym_COMMENT_PREFIX_token1),
	4296: uint16(aux_sym__blank_line_token1),
	4297: uint16(599),
	4298: uint16(8),
	4300: uint16(aux_sym_WS_token1),
	4301: uint16(aux_sym_request_separator_token1),
	4302: uint16(sym_method),
	4303: uint16(aux_sym_http_version_token1),
	4304: uint16(anon_sym_LBRACE_LBRACE),
	4305: uint16(anon_sym_LT),
	4306: uint16(anon_sym_AT2),
	4307: uint16(2),
	4308: uint16(605),
	4309: uint16(4),
	4310: uint16(aux_sym_WORD_CHAR_token1),
	4311: uint16(aux_sym_PUNCTUATION_token1),
	4312: uint16(aux_sym_COMMENT_PREFIX_token1),
	4313: uint16(aux_sym__blank_line_token1),
	4314: uint16(603),
	4315: uint16(8),
	4317: uint16(aux_sym_WS_token1),
	4318: uint16(aux_sym_request_separator_token1),
	4319: uint16(sym_method),
	4320: uint16(aux_sym_http_version_token1),
	4321: uint16(anon_sym_LBRACE_LBRACE),
	4322: uint16(anon_sym_LT),
	4323: uint16(anon_sym_AT2),
	4324: uint16(2),
	4325: uint16(333),
	4326: uint16(4),
	4327: uint16(aux_sym_WORD_CHAR_token1),
	4328: uint16(aux_sym_PUNCTUATION_token1),
	4329: uint16(aux_sym_COMMENT_PREFIX_token1),
	4330: uint16(aux_sym__blank_line_token1),
	4331: uint16(331),
	4332: uint16(8),
	4334: uint16(aux_sym_WS_token1),
	4335: uint16(aux_sym_request_separator_token1),
	4336: uint16(sym_method),
	4337: uint16(aux_sym_http_version_token1),
	4338: uint16(anon_sym_LBRACE_LBRACE),
	4339: uint16(anon_sym_LT),
	4340: uint16(anon_sym_AT2),
	4341: uint16(7),
	4342: uint16(17),
	4343: uint16(1),
	4344: uint16(anon_sym_LBRACE_LBRACE),
	4345: uint16(607),
	4346: uint16(1),
	4347: uint16(aux_sym_WORD_CHAR_token1),
	4348: uint16(611),
	4349: uint16(1),
	4350: uint16(aux_sym_NL_token1),
	4351: uint16(613),
	4352: uint16(1),
	4353: uint16(aux_sym__var_comment_token1),
	4354: uint16(211),
	4355: uint16(1),
	4356: uint16(sym_value),
	4357: uint16(609),
	4358: uint16(2),
	4359: uint16(aux_sym_PUNCTUATION_token1),
	4360: uint16(aux_sym_WS_token1),
	4361: uint16(145),
	4362: uint16(2),
	4363: uint16(sym_variable),
	4364: uint16(aux_sym_value_repeat1),
	4365: uint16(7),
	4366: uint16(17),
	4367: uint16(1),
	4368: uint16(anon_sym_LBRACE_LBRACE),
	4369: uint16(607),
	4370: uint16(1),
	4371: uint16(aux_sym_WORD_CHAR_token1),
	4372: uint16(615),
	4373: uint16(1),
	4374: uint16(aux_sym_NL_token1),
	4375: uint16(617),
	4376: uint16(1),
	4377: uint16(aux_sym__var_comment_token1),
	4378: uint16(231),
	4379: uint16(1),
	4380: uint16(sym_value),
	4381: uint16(609),
	4382: uint16(2),
	4383: uint16(aux_sym_PUNCTUATION_token1),
	4384: uint16(aux_sym_WS_token1),
	4385: uint16(145),
	4386: uint16(2),
	4387: uint16(sym_variable),
	4388: uint16(aux_sym_value_repeat1),
	4389: uint16(6),
	4390: uint16(621),
	4391: uint16(1),
	4392: uint16(aux_sym_PUNCTUATION_token1),
	4393: uint16(623),
	4394: uint16(1),
	4395: uint16(anon_sym_LBRACE_LBRACE),
	4396: uint16(625),
	4397: uint16(1),
	4398: uint16(anon_sym_LBRACE_PERCENT),
	4399: uint16(619),
	4400: uint16(2),
	4401: uint16(aux_sym_WORD_CHAR_token1),
	4402: uint16(aux_sym_path_token1),
	4403: uint16(150),
	4404: uint16(2),
	4405: uint16(sym_variable),
	4406: uint16(aux_sym_path_repeat1),
	4407: uint16(210),
	4408: uint16(2),
	4409: uint16(sym_script),
	4410: uint16(sym_path),
	4411: uint16(6),
	4412: uint16(621),
	4413: uint16(1),
	4414: uint16(aux_sym_PUNCTUATION_token1),
	4415: uint16(623),
	4416: uint16(1),
	4417: uint16(anon_sym_LBRACE_LBRACE),
	4418: uint16(625),
	4419: uint16(1),
	4420: uint16(anon_sym_LBRACE_PERCENT),
	4421: uint16(619),
	4422: uint16(2),
	4423: uint16(aux_sym_WORD_CHAR_token1),
	4424: uint16(aux_sym_path_token1),
	4425: uint16(150),
	4426: uint16(2),
	4427: uint16(sym_variable),
	4428: uint16(aux_sym_path_repeat1),
	4429: uint16(221),
	4430: uint16(2),
	4431: uint16(sym_script),
	4432: uint16(sym_path),
	4433: uint16(7),
	4434: uint16(17),
	4435: uint16(1),
	4436: uint16(anon_sym_LBRACE_LBRACE),
	4437: uint16(607),
	4438: uint16(1),
	4439: uint16(aux_sym_WORD_CHAR_token1),
	4440: uint16(627),
	4441: uint16(1),
	4442: uint16(aux_sym_NL_token1),
	4443: uint16(629),
	4444: uint16(1),
	4445: uint16(aux_sym__var_comment_token1),
	4446: uint16(228),
	4447: uint16(1),
	4448: uint16(sym_value),
	4449: uint16(609),
	4450: uint16(2),
	4451: uint16(aux_sym_PUNCTUATION_token1),
	4452: uint16(aux_sym_WS_token1),
	4453: uint16(145),
	4454: uint16(2),
	4455: uint16(sym_variable),
	4456: uint16(aux_sym_value_repeat1),
	4457: uint16(6),
	4458: uint16(17),
	4459: uint16(1),
	4460: uint16(anon_sym_LBRACE_LBRACE),
	4461: uint16(609),
	4462: uint16(1),
	4463: uint16(aux_sym_PUNCTUATION_token1),
	4464: uint16(611),
	4465: uint16(1),
	4466: uint16(aux_sym_NL_token1),
	4467: uint16(211),
	4468: uint16(1),
	4469: uint16(sym_value),
	4470: uint16(607),
	4471: uint16(2),
	4472: uint16(aux_sym_WORD_CHAR_token1),
	4473: uint16(aux_sym_WS_token1),
	4474: uint16(145),
	4475: uint16(2),
	4476: uint16(sym_variable),
	4477: uint16(aux_sym_value_repeat1),
	4478: uint16(6),
	4479: uint16(17),
	4480: uint16(1),
	4481: uint16(anon_sym_LBRACE_LBRACE),
	4482: uint16(607),
	4483: uint16(1),
	4484: uint16(aux_sym_WORD_CHAR_token1),
	4485: uint16(631),
	4486: uint16(1),
	4487: uint16(aux_sym__var_comment_token1),
	4488: uint16(238),
	4489: uint16(1),
	4490: uint16(sym_value),
	4491: uint16(609),
	4492: uint16(2),
	4493: uint16(aux_sym_PUNCTUATION_token1),
	4494: uint16(aux_sym_WS_token1),
	4495: uint16(145),
	4496: uint16(2),
	4497: uint16(sym_variable),
	4498: uint16(aux_sym_value_repeat1),
	4499: uint16(6),
	4500: uint16(17),
	4501: uint16(1),
	4502: uint16(anon_sym_LBRACE_LBRACE),
	4503: uint16(607),
	4504: uint16(1),
	4505: uint16(aux_sym_WORD_CHAR_token1),
	4506: uint16(633),
	4507: uint16(1),
	4508: uint16(aux_sym__var_comment_token1),
	4509: uint16(246),
	4510: uint16(1),
	4511: uint16(sym_value),
	4512: uint16(609),
	4513: uint16(2),
	4514: uint16(aux_sym_PUNCTUATION_token1),
	4515: uint16(aux_sym_WS_token1),
	4516: uint16(145),
	4517: uint16(2),
	4518: uint16(sym_variable),
	4519: uint16(aux_sym_value_repeat1),
	4520: uint16(6),
	4521: uint16(17),
	4522: uint16(1),
	4523: uint16(anon_sym_LBRACE_LBRACE),
	4524: uint16(607),
	4525: uint16(1),
	4526: uint16(aux_sym_WORD_CHAR_token1),
	4527: uint16(635),
	4528: uint16(1),
	4529: uint16(aux_sym__var_comment_token1),
	4530: uint16(234),
	4531: uint16(1),
	4532: uint16(sym_value),
	4533: uint16(609),
	4534: uint16(2),
	4535: uint16(aux_sym_PUNCTUATION_token1),
	4536: uint16(aux_sym_WS_token1),
	4537: uint16(145),
	4538: uint16(2),
	4539: uint16(sym_variable),
	4540: uint16(aux_sym_value_repeat1),
	4541: uint16(6),
	4542: uint16(17),
	4543: uint16(1),
	4544: uint16(anon_sym_LBRACE_LBRACE),
	4545: uint16(609),
	4546: uint16(1),
	4547: uint16(aux_sym_PUNCTUATION_token1),
	4548: uint16(637),
	4549: uint16(1),
	4550: uint16(aux_sym_NL_token1),
	4551: uint16(240),
	4552: uint16(1),
	4553: uint16(sym_value),
	4554: uint16(607),
	4555: uint16(2),
	4556: uint16(aux_sym_WORD_CHAR_token1),
	4557: uint16(aux_sym_WS_token1),
	4558: uint16(145),
	4559: uint16(2),
	4560: uint16(sym_variable),
	4561: uint16(aux_sym_value_repeat1),
	4562: uint16(6),
	4563: uint16(17),
	4564: uint16(1),
	4565: uint16(anon_sym_LBRACE_LBRACE),
	4566: uint16(609),
	4567: uint16(1),
	4568: uint16(aux_sym_PUNCTUATION_token1),
	4569: uint16(639),
	4570: uint16(1),
	4571: uint16(aux_sym_NL_token1),
	4572: uint16(223),
	4573: uint16(1),
	4574: uint16(sym_value),
	4575: uint16(607),
	4576: uint16(2),
	4577: uint16(aux_sym_WORD_CHAR_token1),
	4578: uint16(aux_sym_WS_token1),
	4579: uint16(145),
	4580: uint16(2),
	4581: uint16(sym_variable),
	4582: uint16(aux_sym_value_repeat1),
	4583: uint16(6),
	4584: uint16(17),
	4585: uint16(1),
	4586: uint16(anon_sym_LBRACE_LBRACE),
	4587: uint16(607),
	4588: uint16(1),
	4589: uint16(aux_sym_WORD_CHAR_token1),
	4590: uint16(641),
	4591: uint16(1),
	4592: uint16(aux_sym__var_comment_token1),
	4593: uint16(232),
	4594: uint16(1),
	4595: uint16(sym_value),
	4596: uint16(609),
	4597: uint16(2),
	4598: uint16(aux_sym_PUNCTUATION_token1),
	4599: uint16(aux_sym_WS_token1),
	4600: uint16(145),
	4601: uint16(2),
	4602: uint16(sym_variable),
	4603: uint16(aux_sym_value_repeat1),
	4604: uint16(6),
	4605: uint16(645),
	4606: uint16(1),
	4607: uint16(aux_sym_PUNCTUATION_token1),
	4608: uint16(647),
	4609: uint16(1),
	4610: uint16(aux_sym_WS_token1),
	4611: uint16(649),
	4612: uint16(1),
	4613: uint16(aux_sym_NL_token1),
	4614: uint16(651),
	4615: uint16(1),
	4616: uint16(anon_sym_LBRACE_LBRACE),
	4617: uint16(643),
	4618: uint16(2),
	4619: uint16(aux_sym_WORD_CHAR_token1),
	4620: uint16(aux_sym_path_token1),
	4621: uint16(130),
	4622: uint16(2),
	4623: uint16(sym_variable),
	4624: uint16(aux_sym_path_repeat1),
	4625: uint16(6),
	4626: uint16(656),
	4627: uint16(1),
	4628: uint16(aux_sym_PUNCTUATION_token1),
	4629: uint16(659),
	4630: uint16(1),
	4631: uint16(aux_sym_WS_token1),
	4632: uint16(661),
	4633: uint16(1),
	4634: uint16(aux_sym_NL_token1),
	4635: uint16(663),
	4636: uint16(1),
	4637: uint16(anon_sym_LBRACE_LBRACE),
	4638: uint16(653),
	4639: uint16(2),
	4640: uint16(aux_sym_WORD_CHAR_token1),
	4641: uint16(aux_sym_path_token1),
	4642: uint16(130),
	4643: uint16(2),
	4644: uint16(sym_variable),
	4645: uint16(aux_sym_path_repeat1),
	4646: uint16(6),
	4647: uint16(17),
	4648: uint16(1),
	4649: uint16(anon_sym_LBRACE_LBRACE),
	4650: uint16(607),
	4651: uint16(1),
	4652: uint16(aux_sym_WORD_CHAR_token1),
	4653: uint16(666),
	4654: uint16(1),
	4655: uint16(aux_sym__var_comment_token1),
	4656: uint16(226),
	4657: uint16(1),
	4658: uint16(sym_value),
	4659: uint16(609),
	4660: uint16(2),
	4661: uint16(aux_sym_PUNCTUATION_token1),
	4662: uint16(aux_sym_WS_token1),
	4663: uint16(145),
	4664: uint16(2),
	4665: uint16(sym_variable),
	4666: uint16(aux_sym_value_repeat1),
	4667: uint16(7),
	4668: uint16(17),
	4669: uint16(1),
	4670: uint16(anon_sym_LBRACE_LBRACE),
	4671: uint16(668),
	4672: uint16(1),
	4673: uint16(aux_sym_WORD_CHAR_token1),
	4674: uint16(670),
	4675: uint16(1),
	4676: uint16(aux_sym_PUNCTUATION_token1),
	4677: uint16(672),
	4678: uint16(1),
	4679: uint16(aux_sym_WS_token1),
	4680: uint16(675),
	4681: uint16(1),
	4682: uint16(aux_sym_NL_token1),
	4683: uint16(174),
	4684: uint16(1),
	4685: uint16(aux_sym_target_url_repeat1),
	4686: uint16(134),
	4687: uint16(2),
	4688: uint16(aux_sym__target_url_line),
	4689: uint16(sym_variable),
	4690: uint16(6),
	4691: uint16(17),
	4692: uint16(1),
	4693: uint16(anon_sym_LBRACE_LBRACE),
	4694: uint16(607),
	4695: uint16(1),
	4696: uint16(aux_sym_WORD_CHAR_token1),
	4697: uint16(678),
	4698: uint16(1),
	4699: uint16(aux_sym__var_comment_token1),
	4700: uint16(243),
	4701: uint16(1),
	4702: uint16(sym_value),
	4703: uint16(609),
	4704: uint16(2),
	4705: uint16(aux_sym_PUNCTUATION_token1),
	4706: uint16(aux_sym_WS_token1),
	4707: uint16(145),
	4708: uint16(2),
	4709: uint16(sym_variable),
	4710: uint16(aux_sym_value_repeat1),
	4711: uint16(5),
	4712: uint16(683),
	4713: uint16(1),
	4714: uint16(aux_sym_PUNCTUATION_token1),
	4715: uint16(686),
	4716: uint16(1),
	4717: uint16(aux_sym_NL_token1),
	4718: uint16(688),
	4719: uint16(1),
	4720: uint16(anon_sym_LBRACE_LBRACE),
	4721: uint16(680),
	4722: uint16(2),
	4723: uint16(aux_sym_WORD_CHAR_token1),
	4724: uint16(aux_sym_WS_token1),
	4725: uint16(134),
	4726: uint16(2),
	4727: uint16(aux_sym__target_url_line),
	4728: uint16(sym_variable),
	4729: uint16(6),
	4730: uint16(17),
	4731: uint16(1),
	4732: uint16(anon_sym_LBRACE_LBRACE),
	4733: uint16(668),
	4734: uint16(1),
	4735: uint16(aux_sym_WORD_CHAR_token1),
	4736: uint16(670),
	4737: uint16(1),
	4738: uint16(aux_sym_PUNCTUATION_token1),
	4739: uint16(691),
	4740: uint16(1),
	4741: uint16(aux_sym_WS_token1),
	4742: uint16(694),
	4743: uint16(1),
	4744: uint16(aux_sym_NL_token1),
	4745: uint16(134),
	4746: uint16(2),
	4747: uint16(aux_sym__target_url_line),
	4748: uint16(sym_variable),
	4749: uint16(5),
	4750: uint16(17),
	4751: uint16(1),
	4752: uint16(anon_sym_LBRACE_LBRACE),
	4753: uint16(609),
	4754: uint16(1),
	4755: uint16(aux_sym_PUNCTUATION_token1),
	4756: uint16(247),
	4757: uint16(1),
	4758: uint16(sym_value),
	4759: uint16(607),
	4760: uint16(2),
	4761: uint16(aux_sym_WORD_CHAR_token1),
	4762: uint16(aux_sym_WS_token1),
	4763: uint16(145),
	4764: uint16(2),
	4765: uint16(sym_variable),
	4766: uint16(aux_sym_value_repeat1),
	4767: uint16(5),
	4768: uint16(621),
	4769: uint16(1),
	4770: uint16(aux_sym_PUNCTUATION_token1),
	4771: uint16(623),
	4772: uint16(1),
	4773: uint16(anon_sym_LBRACE_LBRACE),
	4774: uint16(209),
	4775: uint16(1),
	4776: uint16(sym_path),
	4777: uint16(619),
	4778: uint16(2),
	4779: uint16(aux_sym_WORD_CHAR_token1),
	4780: uint16(aux_sym_path_token1),
	4781: uint16(150),
	4782: uint16(2),
	4783: uint16(sym_variable),
	4784: uint16(aux_sym_path_repeat1),
	4785: uint16(5),
	4786: uint16(17),
	4787: uint16(1),
	4788: uint16(anon_sym_LBRACE_LBRACE),
	4789: uint16(609),
	4790: uint16(1),
	4791: uint16(aux_sym_PUNCTUATION_token1),
	4792: uint16(245),
	4793: uint16(1),
	4794: uint16(sym_value),
	4795: uint16(607),
	4796: uint16(2),
	4797: uint16(aux_sym_WORD_CHAR_token1),
	4798: uint16(aux_sym_WS_token1),
	4799: uint16(145),
	4800: uint16(2),
	4801: uint16(sym_variable),
	4802: uint16(aux_sym_value_repeat1),
	4803: uint16(5),
	4804: uint16(651),
	4805: uint16(1),
	4806: uint16(anon_sym_LBRACE_LBRACE),
	4807: uint16(698),
	4808: uint16(1),
	4809: uint16(aux_sym_PUNCTUATION_token1),
	4810: uint16(204),
	4811: uint16(1),
	4812: uint16(sym_path),
	4813: uint16(696),
	4814: uint16(2),
	4815: uint16(aux_sym_WORD_CHAR_token1),
	4816: uint16(aux_sym_path_token1),
	4817: uint16(129),
	4818: uint16(2),
	4819: uint16(sym_variable),
	4820: uint16(aux_sym_path_repeat1),
	4821: uint16(5),
	4822: uint16(17),
	4823: uint16(1),
	4824: uint16(anon_sym_LBRACE_LBRACE),
	4825: uint16(609),
	4826: uint16(1),
	4827: uint16(aux_sym_PUNCTUATION_token1),
	4828: uint16(244),
	4829: uint16(1),
	4830: uint16(sym_value),
	4831: uint16(607),
	4832: uint16(2),
	4833: uint16(aux_sym_WORD_CHAR_token1),
	4834: uint16(aux_sym_WS_token1),
	4835: uint16(145),
	4836: uint16(2),
	4837: uint16(sym_variable),
	4838: uint16(aux_sym_value_repeat1),
	4839: uint16(5),
	4840: uint16(661),
	4841: uint16(1),
	4842: uint16(aux_sym_pre_request_script_token1),
	4843: uint16(703),
	4844: uint16(1),
	4845: uint16(aux_sym_PUNCTUATION_token1),
	4846: uint16(706),
	4847: uint16(1),
	4848: uint16(anon_sym_LBRACE_LBRACE),
	4849: uint16(700),
	4850: uint16(2),
	4851: uint16(aux_sym_WORD_CHAR_token1),
	4852: uint16(aux_sym_path_token1),
	4853: uint16(141),
	4854: uint16(2),
	4855: uint16(sym_variable),
	4856: uint16(aux_sym_path_repeat1),
	4857: uint16(5),
	4858: uint16(17),
	4859: uint16(1),
	4860: uint16(anon_sym_LBRACE_LBRACE),
	4861: uint16(609),
	4862: uint16(1),
	4863: uint16(aux_sym_PUNCTUATION_token1),
	4864: uint16(236),
	4865: uint16(1),
	4866: uint16(sym_value),
	4867: uint16(607),
	4868: uint16(2),
	4869: uint16(aux_sym_WORD_CHAR_token1),
	4870: uint16(aux_sym_WS_token1),
	4871: uint16(145),
	4872: uint16(2),
	4873: uint16(sym_variable),
	4874: uint16(aux_sym_value_repeat1),
	4875: uint16(5),
	4876: uint16(17),
	4877: uint16(1),
	4878: uint16(anon_sym_LBRACE_LBRACE),
	4879: uint16(609),
	4880: uint16(1),
	4881: uint16(aux_sym_PUNCTUATION_token1),
	4882: uint16(222),
	4883: uint16(1),
	4884: uint16(sym_value),
	4885: uint16(607),
	4886: uint16(2),
	4887: uint16(aux_sym_WORD_CHAR_token1),
	4888: uint16(aux_sym_WS_token1),
	4889: uint16(145),
	4890: uint16(2),
	4891: uint16(sym_variable),
	4892: uint16(aux_sym_value_repeat1),
	4893: uint16(5),
	4894: uint16(17),
	4895: uint16(1),
	4896: uint16(anon_sym_LBRACE_LBRACE),
	4897: uint16(609),
	4898: uint16(1),
	4899: uint16(aux_sym_PUNCTUATION_token1),
	4900: uint16(238),
	4901: uint16(1),
	4902: uint16(sym_value),
	4903: uint16(607),
	4904: uint16(2),
	4905: uint16(aux_sym_WORD_CHAR_token1),
	4906: uint16(aux_sym_WS_token1),
	4907: uint16(145),
	4908: uint16(2),
	4909: uint16(sym_variable),
	4910: uint16(aux_sym_value_repeat1),
	4911: uint16(5),
	4912: uint16(17),
	4913: uint16(1),
	4914: uint16(anon_sym_LBRACE_LBRACE),
	4915: uint16(711),
	4916: uint16(1),
	4917: uint16(aux_sym_PUNCTUATION_token1),
	4918: uint16(713),
	4919: uint16(1),
	4920: uint16(aux_sym_NL_token1),
	4921: uint16(709),
	4922: uint16(2),
	4923: uint16(aux_sym_WORD_CHAR_token1),
	4924: uint16(aux_sym_WS_token1),
	4925: uint16(149),
	4926: uint16(2),
	4927: uint16(sym_variable),
	4928: uint16(aux_sym_value_repeat1),
	4929: uint16(5),
	4930: uint16(5),
	4931: uint16(1),
	4932: uint16(aux_sym_PUNCTUATION_token1),
	4933: uint16(17),
	4934: uint16(1),
	4935: uint16(anon_sym_LBRACE_LBRACE),
	4936: uint16(199),
	4937: uint16(1),
	4938: uint16(sym_target_url),
	4939: uint16(715),
	4940: uint16(2),
	4941: uint16(aux_sym_WORD_CHAR_token1),
	4942: uint16(aux_sym_WS_token1),
	4943: uint16(132),
	4944: uint16(2),
	4945: uint16(aux_sym__target_url_line),
	4946: uint16(sym_variable),
	4947: uint16(5),
	4948: uint16(651),
	4949: uint16(1),
	4950: uint16(anon_sym_LBRACE_LBRACE),
	4951: uint16(698),
	4952: uint16(1),
	4953: uint16(aux_sym_PUNCTUATION_token1),
	4954: uint16(196),
	4955: uint16(1),
	4956: uint16(sym_path),
	4957: uint16(696),
	4958: uint16(2),
	4959: uint16(aux_sym_WORD_CHAR_token1),
	4960: uint16(aux_sym_path_token1),
	4961: uint16(129),
	4962: uint16(2),
	4963: uint16(sym_variable),
	4964: uint16(aux_sym_path_repeat1),
	4965: uint16(5),
	4966: uint16(651),
	4967: uint16(1),
	4968: uint16(anon_sym_LBRACE_LBRACE),
	4969: uint16(698),
	4970: uint16(1),
	4971: uint16(aux_sym_PUNCTUATION_token1),
	4972: uint16(191),
	4973: uint16(1),
	4974: uint16(sym_path),
	4975: uint16(696),
	4976: uint16(2),
	4977: uint16(aux_sym_WORD_CHAR_token1),
	4978: uint16(aux_sym_path_token1),
	4979: uint16(129),
	4980: uint16(2),
	4981: uint16(sym_variable),
	4982: uint16(aux_sym_path_repeat1),
	4983: uint16(5),
	4984: uint16(720),
	4985: uint16(1),
	4986: uint16(aux_sym_PUNCTUATION_token1),
	4987: uint16(723),
	4988: uint16(1),
	4989: uint16(aux_sym_NL_token1),
	4990: uint16(725),
	4991: uint16(1),
	4992: uint16(anon_sym_LBRACE_LBRACE),
	4993: uint16(717),
	4994: uint16(2),
	4995: uint16(aux_sym_WORD_CHAR_token1),
	4996: uint16(aux_sym_WS_token1),
	4997: uint16(149),
	4998: uint16(2),
	4999: uint16(sym_variable),
	5000: uint16(aux_sym_value_repeat1),
	5001: uint16(5),
	5002: uint16(623),
	5003: uint16(1),
	5004: uint16(anon_sym_LBRACE_LBRACE),
	5005: uint16(649),
	5006: uint16(1),
	5007: uint16(aux_sym_pre_request_script_token1),
	5008: uint16(730),
	5009: uint16(1),
	5010: uint16(aux_sym_PUNCTUATION_token1),
	5011: uint16(728),
	5012: uint16(2),
	5013: uint16(aux_sym_WORD_CHAR_token1),
	5014: uint16(aux_sym_path_token1),
	5015: uint16(141),
	5016: uint16(2),
	5017: uint16(sym_variable),
	5018: uint16(aux_sym_path_repeat1),
	5019: uint16(2),
	5020: uint16(734),
	5021: uint16(2),
	5022: uint16(aux_sym_PUNCTUATION_token1),
	5023: uint16(aux_sym_NL_token1),
	5024: uint16(732),
	5025: uint16(4),
	5026: uint16(aux_sym_WORD_CHAR_token1),
	5027: uint16(aux_sym_WS_token1),
	5028: uint16(anon_sym_LBRACE_LBRACE),
	5029: uint16(aux_sym_path_token1),
	5030: uint16(2),
	5031: uint16(738),
	5032: uint16(2),
	5033: uint16(aux_sym_PUNCTUATION_token1),
	5034: uint16(aux_sym_NL_token1),
	5035: uint16(736),
	5036: uint16(4),
	5037: uint16(aux_sym_WORD_CHAR_token1),
	5038: uint16(aux_sym_WS_token1),
	5039: uint16(anon_sym_LBRACE_LBRACE),
	5040: uint16(aux_sym_path_token1),
	5041: uint16(2),
	5042: uint16(742),
	5043: uint16(2),
	5044: uint16(aux_sym_PUNCTUATION_token1),
	5045: uint16(aux_sym_NL_token1),
	5046: uint16(740),
	5047: uint16(4),
	5048: uint16(aux_sym_WORD_CHAR_token1),
	5049: uint16(aux_sym_WS_token1),
	5050: uint16(anon_sym_LBRACE_LBRACE),
	5051: uint16(aux_sym_path_token1),
	5052: uint16(2),
	5053: uint16(746),
	5054: uint16(2),
	5055: uint16(aux_sym_PUNCTUATION_token1),
	5056: uint16(aux_sym_NL_token1),
	5057: uint16(744),
	5058: uint16(4),
	5059: uint16(aux_sym_WORD_CHAR_token1),
	5060: uint16(aux_sym_WS_token1),
	5061: uint16(anon_sym_LBRACE_LBRACE),
	5062: uint16(aux_sym_path_token1),
	5063: uint16(3),
	5064: uint16(752),
	5065: uint16(1),
	5066: uint16(aux_sym__blank_line_token1),
	5067: uint16(750),
	5068: uint16(2),
	5069: uint16(aux_sym_PUNCTUATION_token1),
	5070: uint16(aux_sym_NL_token1),
	5071: uint16(748),
	5072: uint16(3),
	5073: uint16(aux_sym_WORD_CHAR_token1),
	5074: uint16(aux_sym_WS_token1),
	5075: uint16(anon_sym_LBRACE_LBRACE),
	5076: uint16(4),
	5077: uint16(17),
	5078: uint16(1),
	5079: uint16(anon_sym_LBRACE_LBRACE),
	5080: uint16(756),
	5081: uint16(1),
	5082: uint16(aux_sym_PUNCTUATION_token1),
	5083: uint16(754),
	5084: uint16(2),
	5085: uint16(aux_sym_WORD_CHAR_token1),
	5086: uint16(aux_sym_WS_token1),
	5087: uint16(135),
	5088: uint16(2),
	5089: uint16(aux_sym__target_url_line),
	5090: uint16(sym_variable),
	5091: uint16(2),
	5092: uint16(738),
	5093: uint16(2),
	5094: uint16(aux_sym_PUNCTUATION_token1),
	5095: uint16(aux_sym_pre_request_script_token1),
	5096: uint16(736),
	5097: uint16(3),
	5098: uint16(aux_sym_WORD_CHAR_token1),
	5099: uint16(anon_sym_LBRACE_LBRACE),
	5100: uint16(aux_sym_path_token1),
	5101: uint16(2),
	5102: uint16(742),
	5103: uint16(2),
	5104: uint16(aux_sym_PUNCTUATION_token1),
	5105: uint16(aux_sym_pre_request_script_token1),
	5106: uint16(740),
	5107: uint16(3),
	5108: uint16(aux_sym_WORD_CHAR_token1),
	5109: uint16(anon_sym_LBRACE_LBRACE),
	5110: uint16(aux_sym_path_token1),
	5111: uint16(2),
	5112: uint16(746),
	5113: uint16(2),
	5114: uint16(aux_sym_PUNCTUATION_token1),
	5115: uint16(aux_sym_pre_request_script_token1),
	5116: uint16(744),
	5117: uint16(3),
	5118: uint16(aux_sym_WORD_CHAR_token1),
	5119: uint16(anon_sym_LBRACE_LBRACE),
	5120: uint16(aux_sym_path_token1),
	5121: uint16(2),
	5122: uint16(734),
	5123: uint16(2),
	5124: uint16(aux_sym_PUNCTUATION_token1),
	5125: uint16(aux_sym_pre_request_script_token1),
	5126: uint16(732),
	5127: uint16(3),
	5128: uint16(aux_sym_WORD_CHAR_token1),
	5129: uint16(anon_sym_LBRACE_LBRACE),
	5130: uint16(aux_sym_path_token1),
	5131: uint16(2),
	5132: uint16(738),
	5133: uint16(2),
	5134: uint16(aux_sym_PUNCTUATION_token1),
	5135: uint16(aux_sym_NL_token1),
	5136: uint16(736),
	5137: uint16(3),
	5138: uint16(aux_sym_WORD_CHAR_token1),
	5139: uint16(aux_sym_WS_token1),
	5140: uint16(anon_sym_LBRACE_LBRACE),
	5141: uint16(2),
	5142: uint16(742),
	5143: uint16(2),
	5144: uint16(aux_sym_PUNCTUATION_token1),
	5145: uint16(aux_sym_NL_token1),
	5146: uint16(740),
	5147: uint16(3),
	5148: uint16(aux_sym_WORD_CHAR_token1),
	5149: uint16(aux_sym_WS_token1),
	5150: uint16(anon_sym_LBRACE_LBRACE),
	5151: uint16(2),
	5152: uint16(746),
	5153: uint16(2),
	5154: uint16(aux_sym_PUNCTUATION_token1),
	5155: uint16(aux_sym_NL_token1),
	5156: uint16(744),
	5157: uint16(3),
	5158: uint16(aux_sym_WORD_CHAR_token1),
	5159: uint16(aux_sym_WS_token1),
	5160: uint16(anon_sym_LBRACE_LBRACE),
	5161: uint16(2),
	5162: uint16(734),
	5163: uint16(2),
	5164: uint16(aux_sym_PUNCTUATION_token1),
	5165: uint16(aux_sym_NL_token1),
	5166: uint16(732),
	5167: uint16(3),
	5168: uint16(aux_sym_WORD_CHAR_token1),
	5169: uint16(aux_sym_WS_token1),
	5170: uint16(anon_sym_LBRACE_LBRACE),
	5171: uint16(3),
	5172: uint16(758),
	5173: uint16(1),
	5174: uint16(aux_sym_LINE_TAIL_token1),
	5175: uint16(760),
	5176: uint16(1),
	5177: uint16(anon_sym_PERCENT_RBRACE),
	5178: uint16(177),
	5179: uint16(1),
	5180: uint16(aux_sym_script_repeat1),
	5181: uint16(3),
	5182: uint16(762),
	5183: uint16(1),
	5184: uint16(aux_sym_LINE_TAIL_token1),
	5185: uint16(765),
	5186: uint16(1),
	5187: uint16(anon_sym_PERCENT_RBRACE),
	5188: uint16(166),
	5189: uint16(1),
	5190: uint16(aux_sym_script_repeat1),
	5191: uint16(3),
	5192: uint16(325),
	5193: uint16(1),
	5194: uint16(aux_sym__raw_body_token1),
	5195: uint16(767),
	5196: uint16(1),
	5197: uint16(aux_sym_COMMENT_PREFIX_token1),
	5198: uint16(66),
	5199: uint16(1),
	5200: uint16(sym__raw_body),
	5201: uint16(3),
	5202: uint16(325),
	5203: uint16(1),
	5204: uint16(aux_sym__raw_body_token1),
	5205: uint16(767),
	5206: uint16(1),
	5207: uint16(aux_sym_COMMENT_PREFIX_token1),
	5208: uint16(63),
	5209: uint16(1),
	5210: uint16(sym__raw_body),
	5211: uint16(3),
	5212: uint16(309),
	5213: uint16(1),
	5214: uint16(aux_sym__raw_body_token1),
	5215: uint16(769),
	5216: uint16(1),
	5217: uint16(aux_sym_COMMENT_PREFIX_token1),
	5218: uint16(49),
	5219: uint16(1),
	5220: uint16(sym__raw_body),
	5221: uint16(2),
	5222: uint16(773),
	5223: uint16(1),
	5224: uint16(aux_sym_NL_token1),
	5225: uint16(771),
	5226: uint16(2),
	5227: uint16(aux_sym_WS_token1),
	5228: uint16(anon_sym_EQ),
	5229: uint16(2),
	5230: uint16(777),
	5231: uint16(1),
	5232: uint16(aux_sym_NL_token1),
	5233: uint16(775),
	5234: uint16(2),
	5235: uint16(aux_sym_WS_token1),
	5236: uint16(anon_sym_EQ),
	5237: uint16(2),
	5238: uint16(781),
	5239: uint16(1),
	5240: uint16(aux_sym_NL_token1),
	5241: uint16(779),
	5242: uint16(2),
	5243: uint16(aux_sym_WS_token1),
	5244: uint16(anon_sym_EQ),
	5245: uint16(2),
	5246: uint16(785),
	5247: uint16(1),
	5248: uint16(aux_sym_NL_token1),
	5249: uint16(783),
	5250: uint16(2),
	5251: uint16(aux_sym_WS_token1),
	5252: uint16(anon_sym_EQ),
	5253: uint16(3),
	5254: uint16(787),
	5255: uint16(1),
	5256: uint16(aux_sym_WS_token1),
	5257: uint16(789),
	5258: uint16(1),
	5259: uint16(aux_sym_NL_token1),
	5260: uint16(175),
	5261: uint16(1),
	5262: uint16(aux_sym_target_url_repeat1),
	5263: uint16(3),
	5264: uint16(792),
	5265: uint16(1),
	5266: uint16(aux_sym_WS_token1),
	5267: uint16(794),
	5268: uint16(1),
	5269: uint16(aux_sym_NL_token1),
	5270: uint16(175),
	5271: uint16(1),
	5272: uint16(aux_sym_target_url_repeat1),
	5273: uint16(3),
	5274: uint16(325),
	5275: uint16(1),
	5276: uint16(aux_sym__raw_body_token1),
	5277: uint16(767),
	5278: uint16(1),
	5279: uint16(aux_sym_COMMENT_PREFIX_token1),
	5280: uint16(54),
	5281: uint16(1),
	5282: uint16(sym__raw_body),
	5283: uint16(3),
	5284: uint16(797),
	5285: uint16(1),
	5286: uint16(aux_sym_LINE_TAIL_token1),
	5287: uint16(799),
	5288: uint16(1),
	5289: uint16(anon_sym_PERCENT_RBRACE),
	5290: uint16(166),
	5291: uint16(1),
	5292: uint16(aux_sym_script_repeat1),
	5293: uint16(2),
	5294: uint16(801),
	5295: uint16(1),
	5296: uint16(aux_sym_LINE_TAIL_token1),
	5297: uint16(803),
	5298: uint16(1),
	5299: uint16(anon_sym_AT),
	5300: uint16(2),
	5301: uint16(805),
	5302: uint16(1),
	5303: uint16(anon_sym_AT),
	5304: uint16(807),
	5305: uint16(1),
	5306: uint16(sym__not_comment),
	5307: uint16(2),
	5308: uint16(809),
	5309: uint16(1),
	5310: uint16(aux_sym_WS_token1),
	5311: uint16(811),
	5312: uint16(1),
	5313: uint16(sym_identifier),
	5314: uint16(2),
	5315: uint16(15),
	5316: uint16(1),
	5317: uint16(aux_sym_http_version_token1),
	5318: uint16(252),
	5319: uint16(1),
	5320: uint16(sym_http_version),
	5321: uint16(2),
	5322: uint16(813),
	5323: uint16(1),
	5324: uint16(aux_sym_WS_token1),
	5325: uint16(815),
	5326: uint16(1),
	5327: uint16(anon_sym_COLON),
	5328: uint16(2),
	5329: uint16(817),
	5330: uint16(1),
	5331: uint16(aux_sym_WS_token1),
	5332: uint16(819),
	5333: uint16(1),
	5334: uint16(sym_identifier),
	5335: uint16(2),
	5336: uint16(821),
	5337: uint16(1),
	5338: uint16(aux_sym_WS_token1),
	5339: uint16(823),
	5340: uint16(1),
	5341: uint16(sym_identifier),
	5342: uint16(2),
	5343: uint16(825),
	5344: uint16(1),
	5345: uint16(aux_sym_WS_token1),
	5346: uint16(827),
	5347: uint16(1),
	5348: uint16(anon_sym_EQ),
	5349: uint16(2),
	5350: uint16(829),
	5351: uint16(1),
	5352: uint16(aux_sym_WS_token1),
	5353: uint16(831),
	5354: uint16(1),
	5355: uint16(anon_sym_RBRACE_RBRACE),
	5356: uint16(2),
	5357: uint16(833),
	5358: uint16(1),
	5359: uint16(aux_sym_WS_token1),
	5360: uint16(835),
	5361: uint16(1),
	5362: uint16(anon_sym_RBRACE_RBRACE),
	5363: uint16(2),
	5364: uint16(837),
	5365: uint16(1),
	5366: uint16(aux_sym_LINE_TAIL_token1),
	5367: uint16(839),
	5368: uint16(1),
	5369: uint16(anon_sym_AT),
	5370: uint16(2),
	5371: uint16(841),
	5372: uint16(1),
	5373: uint16(aux_sym_WS_token1),
	5374: uint16(843),
	5375: uint16(1),
	5376: uint16(anon_sym_RBRACE_RBRACE),
	5377: uint16(2),
	5378: uint16(845),
	5379: uint16(1),
	5380: uint16(aux_sym_WS_token1),
	5381: uint16(847),
	5382: uint16(1),
	5383: uint16(anon_sym_RBRACE_RBRACE),
	5384: uint16(1),
	5385: uint16(849),
	5386: uint16(2),
	5387: uint16(aux_sym_WS_token1),
	5388: uint16(aux_sym_NL_token1),
	5389: uint16(2),
	5390: uint16(851),
	5391: uint16(1),
	5392: uint16(aux_sym_WS_token1),
	5393: uint16(853),
	5394: uint16(1),
	5395: uint16(anon_sym_RBRACE_RBRACE),
	5396: uint16(2),
	5397: uint16(855),
	5398: uint16(1),
	5399: uint16(aux_sym_LINE_TAIL_token1),
	5400: uint16(857),
	5401: uint16(1),
	5402: uint16(anon_sym_AT),
	5403: uint16(1),
	5404: uint16(859),
	5405: uint16(2),
	5406: uint16(aux_sym_WS_token1),
	5407: uint16(aux_sym_NL_token1),
	5408: uint16(2),
	5409: uint16(861),
	5410: uint16(1),
	5411: uint16(aux_sym_WS_token1),
	5412: uint16(863),
	5413: uint16(1),
	5414: uint16(sym_identifier),
	5415: uint16(1),
	5416: uint16(865),
	5417: uint16(2),
	5418: uint16(aux_sym_WS_token1),
	5419: uint16(aux_sym_NL_token1),
	5420: uint16(2),
	5421: uint16(867),
	5422: uint16(1),
	5423: uint16(aux_sym_NL_token1),
	5424: uint16(869),
	5425: uint16(1),
	5426: uint16(sym_status_text),
	5427: uint16(1),
	5428: uint16(871),
	5429: uint16(2),
	5430: uint16(aux_sym_WS_token1),
	5431: uint16(aux_sym_NL_token1),
	5432: uint16(2),
	5433: uint16(873),
	5434: uint16(1),
	5435: uint16(aux_sym_WS_token1),
	5436: uint16(875),
	5437: uint16(1),
	5438: uint16(aux_sym_NL_token1),
	5439: uint16(2),
	5440: uint16(877),
	5441: uint16(1),
	5442: uint16(aux_sym_WS_token1),
	5443: uint16(879),
	5444: uint16(1),
	5445: uint16(aux_sym_NL_token1),
	5446: uint16(2),
	5447: uint16(881),
	5448: uint16(1),
	5449: uint16(aux_sym_WS_token1),
	5450: uint16(883),
	5451: uint16(1),
	5452: uint16(anon_sym_AT2),
	5453: uint16(2),
	5454: uint16(15),
	5455: uint16(1),
	5456: uint16(aux_sym_http_version_token1),
	5457: uint16(224),
	5458: uint16(1),
	5459: uint16(sym_http_version),
	5460: uint16(2),
	5461: uint16(885),
	5462: uint16(1),
	5463: uint16(aux_sym_WS_token1),
	5464: uint16(887),
	5465: uint16(1),
	5466: uint16(anon_sym_RBRACE_RBRACE),
	5467: uint16(1),
	5468: uint16(889),
	5469: uint16(2),
	5470: uint16(aux_sym_WS_token1),
	5471: uint16(aux_sym_NL_token1),
	5472: uint16(1),
	5473: uint16(891),
	5474: uint16(1),
	5475: uint16(aux_sym_NL_token1),
	5476: uint16(1),
	5477: uint16(893),
	5478: uint16(1),
	5479: uint16(anon_sym_RBRACE_RBRACE),
	5480: uint16(1),
	5481: uint16(895),
	5482: uint16(1),
	5483: uint16(aux_sym_multipart_form_data_token1),
	5484: uint16(1),
	5485: uint16(897),
	5486: uint16(1),
	5487: uint16(aux_sym_WS_token1),
	5488: uint16(1),
	5489: uint16(899),
	5490: uint16(1),
	5491: uint16(aux_sym_pre_request_script_token1),
	5492: uint16(1),
	5493: uint16(901),
	5494: uint16(1),
	5495: uint16(aux_sym_pre_request_script_token1),
	5496: uint16(1),
	5497: uint16(903),
	5498: uint16(1),
	5499: uint16(aux_sym_NL_token1),
	5500: uint16(1),
	5501: uint16(905),
	5502: uint16(1),
	5503: uint16(aux_sym_WS_token1),
	5504: uint16(1),
	5505: uint16(907),
	5506: uint16(1),
	5507: uint16(aux_sym_WS_token1),
	5508: uint16(1),
	5509: uint16(909),
	5510: uint16(1),
	5511: uint16(aux_sym_WS_token1),
	5512: uint16(1),
	5513: uint16(911),
	5514: uint16(1),
	5515: uint16(anon_sym_COLON),
	5516: uint16(1),
	5517: uint16(913),
	5518: uint16(1),
	5520: uint16(1),
	5521: uint16(915),
	5522: uint16(1),
	5523: uint16(sym_identifier),
	5524: uint16(1),
	5525: uint16(917),
	5526: uint16(1),
	5527: uint16(anon_sym_RBRACE_RBRACE),
	5528: uint16(1),
	5529: uint16(919),
	5530: uint16(1),
	5531: uint16(aux_sym_WS_token1),
	5532: uint16(1),
	5533: uint16(921),
	5534: uint16(1),
	5535: uint16(aux_sym_NL_token1),
	5536: uint16(1),
	5537: uint16(923),
	5538: uint16(1),
	5539: uint16(aux_sym_pre_request_script_token1),
	5540: uint16(1),
	5541: uint16(925),
	5542: uint16(1),
	5543: uint16(aux_sym_NL_token1),
	5544: uint16(1),
	5545: uint16(927),
	5546: uint16(1),
	5547: uint16(aux_sym_NL_token1),
	5548: uint16(1),
	5549: uint16(929),
	5550: uint16(1),
	5551: uint16(aux_sym_NL_token1),
	5552: uint16(1),
	5553: uint16(931),
	5554: uint16(1),
	5555: uint16(aux_sym_WS_token1),
	5556: uint16(1),
	5557: uint16(933),
	5558: uint16(1),
	5559: uint16(aux_sym_NL_token1),
	5560: uint16(1),
	5561: uint16(935),
	5562: uint16(1),
	5563: uint16(sym_identifier),
	5564: uint16(1),
	5565: uint16(937),
	5566: uint16(1),
	5567: uint16(aux_sym_NL_token1),
	5568: uint16(1),
	5569: uint16(939),
	5570: uint16(1),
	5571: uint16(aux_sym_pre_request_script_token1),
	5572: uint16(1),
	5573: uint16(941),
	5574: uint16(1),
	5575: uint16(aux_sym_pre_request_script_token1),
	5576: uint16(1),
	5577: uint16(943),
	5578: uint16(1),
	5579: uint16(aux_sym_NL_token1),
	5580: uint16(1),
	5581: uint16(945),
	5582: uint16(1),
	5583: uint16(aux_sym_NL_token1),
	5584: uint16(1),
	5585: uint16(947),
	5586: uint16(1),
	5587: uint16(anon_sym_RBRACE_RBRACE),
	5588: uint16(1),
	5589: uint16(949),
	5590: uint16(1),
	5591: uint16(aux_sym_NL_token1),
	5592: uint16(1),
	5593: uint16(951),
	5594: uint16(1),
	5595: uint16(sym__not_comment),
	5596: uint16(1),
	5597: uint16(953),
	5598: uint16(1),
	5599: uint16(aux_sym_NL_token1),
	5600: uint16(1),
	5601: uint16(955),
	5602: uint16(1),
	5603: uint16(sym_identifier),
	5604: uint16(1),
	5605: uint16(957),
	5606: uint16(1),
	5607: uint16(aux_sym_NL_token1),
	5608: uint16(1),
	5609: uint16(959),
	5610: uint16(1),
	5611: uint16(anon_sym_EQ),
	5612: uint16(1),
	5613: uint16(961),
	5614: uint16(1),
	5615: uint16(aux_sym_NL_token1),
	5616: uint16(1),
	5617: uint16(963),
	5618: uint16(1),
	5619: uint16(anon_sym_RBRACE_RBRACE),
	5620: uint16(1),
	5621: uint16(965),
	5622: uint16(1),
	5623: uint16(anon_sym_RBRACE_RBRACE),
	5624: uint16(1),
	5625: uint16(967),
	5626: uint16(1),
	5627: uint16(aux_sym_NL_token1),
	5628: uint16(1),
	5629: uint16(969),
	5630: uint16(1),
	5631: uint16(aux_sym_NL_token1),
	5632: uint16(1),
	5633: uint16(971),
	5634: uint16(1),
	5635: uint16(aux_sym_NL_token1),
	5636: uint16(1),
	5637: uint16(973),
	5638: uint16(1),
	5639: uint16(aux_sym_NL_token1),
	5640: uint16(1),
	5641: uint16(975),
	5642: uint16(1),
	5643: uint16(aux_sym_NL_token1),
	5644: uint16(1),
	5645: uint16(977),
	5646: uint16(1),
	5647: uint16(aux_sym_NL_token1),
	5648: uint16(1),
	5649: uint16(979),
	5650: uint16(1),
	5651: uint16(sym_identifier),
	5652: uint16(1),
	5653: uint16(981),
	5654: uint16(1),
	5655: uint16(sym_identifier),
	5656: uint16(1),
	5657: uint16(983),
	5658: uint16(1),
	5659: uint16(aux_sym_WS_token1),
	5660: uint16(1),
	5661: uint16(985),
	5662: uint16(1),
	5663: uint16(aux_sym_NL_token1),
	5664: uint16(1),
	5665: uint16(987),
	5666: uint16(1),
	5667: uint16(anon_sym_RBRACE_RBRACE),
	5668: uint16(1),
	5669: uint16(989),
	5670: uint16(1),
	5671: uint16(sym_identifier),
	5672: uint16(1),
	5673: uint16(991),
	5674: uint16(1),
	5675: uint16(sym_identifier),
	5676: uint16(1),
	5677: uint16(993),
	5678: uint16(1),
	5679: uint16(aux_sym__blank_line_token1),
	5680: uint16(1),
	5681: uint16(995),
	5682: uint16(1),
	5683: uint16(sym_status_code),
	5684: uint16(1),
	5685: uint16(997),
	5686: uint16(1),
	5687: uint16(sym_identifier),
	5688: uint16(1),
	5689: uint16(999),
	5690: uint16(1),
	5691: uint16(aux_sym_WS_token1),
	5692: uint16(1),
	5693: uint16(1001),
	5694: uint16(1),
	5695: uint16(sym__not_comment),
}

var ts_small_parse_table_map = [259]uint32_t{
	1:   uint32(76),
	2:   uint32(152),
	3:   uint32(228),
	4:   uint32(304),
	5:   uint32(380),
	6:   uint32(456),
	7:   uint32(532),
	8:   uint32(608),
	9:   uint32(684),
	10:  uint32(760),
	11:  uint32(836),
	12:  uint32(912),
	13:  uint32(980),
	14:  uint32(1048),
	15:  uint32(1116),
	16:  uint32(1184),
	17:  uint32(1252),
	18:  uint32(1320),
	19:  uint32(1388),
	20:  uint32(1456),
	21:  uint32(1524),
	22:  uint32(1592),
	23:  uint32(1660),
	24:  uint32(1728),
	25:  uint32(1796),
	26:  uint32(1869),
	27:  uint32(1918),
	28:  uint32(1991),
	29:  uint32(2040),
	30:  uint32(2089),
	31:  uint32(2153),
	32:  uint32(2217),
	33:  uint32(2252),
	34:  uint32(2287),
	35:  uint32(2315),
	36:  uint32(2349),
	37:  uint32(2377),
	38:  uint32(2405),
	39:  uint32(2433),
	40:  uint32(2467),
	41:  uint32(2501),
	42:  uint32(2535),
	43:  uint32(2563),
	44:  uint32(2591),
	45:  uint32(2619),
	46:  uint32(2651),
	47:  uint32(2685),
	48:  uint32(2712),
	49:  uint32(2739),
	50:  uint32(2766),
	51:  uint32(2792),
	52:  uint32(2818),
	53:  uint32(2844),
	54:  uint32(2870),
	55:  uint32(2896),
	56:  uint32(2922),
	57:  uint32(2948),
	58:  uint32(2974),
	59:  uint32(3000),
	60:  uint32(3026),
	61:  uint32(3052),
	62:  uint32(3078),
	63:  uint32(3104),
	64:  uint32(3130),
	65:  uint32(3156),
	66:  uint32(3182),
	67:  uint32(3208),
	68:  uint32(3251),
	69:  uint32(3294),
	70:  uint32(3337),
	71:  uint32(3380),
	72:  uint32(3423),
	73:  uint32(3466),
	74:  uint32(3509),
	75:  uint32(3552),
	76:  uint32(3587),
	77:  uint32(3621),
	78:  uint32(3655),
	79:  uint32(3689),
	80:  uint32(3723),
	81:  uint32(3749),
	82:  uint32(3767),
	83:  uint32(3785),
	84:  uint32(3803),
	85:  uint32(3821),
	86:  uint32(3839),
	87:  uint32(3857),
	88:  uint32(3875),
	89:  uint32(3893),
	90:  uint32(3911),
	91:  uint32(3929),
	92:  uint32(3947),
	93:  uint32(3965),
	94:  uint32(3983),
	95:  uint32(4001),
	96:  uint32(4018),
	97:  uint32(4035),
	98:  uint32(4052),
	99:  uint32(4069),
	100: uint32(4086),
	101: uint32(4103),
	102: uint32(4120),
	103: uint32(4137),
	104: uint32(4154),
	105: uint32(4171),
	106: uint32(4188),
	107: uint32(4205),
	108: uint32(4222),
	109: uint32(4239),
	110: uint32(4256),
	111: uint32(4273),
	112: uint32(4290),
	113: uint32(4307),
	114: uint32(4324),
	115: uint32(4341),
	116: uint32(4365),
	117: uint32(4389),
	118: uint32(4411),
	119: uint32(4433),
	120: uint32(4457),
	121: uint32(4478),
	122: uint32(4499),
	123: uint32(4520),
	124: uint32(4541),
	125: uint32(4562),
	126: uint32(4583),
	127: uint32(4604),
	128: uint32(4625),
	129: uint32(4646),
	130: uint32(4667),
	131: uint32(4690),
	132: uint32(4711),
	133: uint32(4729),
	134: uint32(4749),
	135: uint32(4767),
	136: uint32(4785),
	137: uint32(4803),
	138: uint32(4821),
	139: uint32(4839),
	140: uint32(4857),
	141: uint32(4875),
	142: uint32(4893),
	143: uint32(4911),
	144: uint32(4929),
	145: uint32(4947),
	146: uint32(4965),
	147: uint32(4983),
	148: uint32(5001),
	149: uint32(5019),
	150: uint32(5030),
	151: uint32(5041),
	152: uint32(5052),
	153: uint32(5063),
	154: uint32(5076),
	155: uint32(5091),
	156: uint32(5101),
	157: uint32(5111),
	158: uint32(5121),
	159: uint32(5131),
	160: uint32(5141),
	161: uint32(5151),
	162: uint32(5161),
	163: uint32(5171),
	164: uint32(5181),
	165: uint32(5191),
	166: uint32(5201),
	167: uint32(5211),
	168: uint32(5221),
	169: uint32(5229),
	170: uint32(5237),
	171: uint32(5245),
	172: uint32(5253),
	173: uint32(5263),
	174: uint32(5273),
	175: uint32(5283),
	176: uint32(5293),
	177: uint32(5300),
	178: uint32(5307),
	179: uint32(5314),
	180: uint32(5321),
	181: uint32(5328),
	182: uint32(5335),
	183: uint32(5342),
	184: uint32(5349),
	185: uint32(5356),
	186: uint32(5363),
	187: uint32(5370),
	188: uint32(5377),
	189: uint32(5384),
	190: uint32(5389),
	191: uint32(5396),
	192: uint32(5403),
	193: uint32(5408),
	194: uint32(5415),
	195: uint32(5420),
	196: uint32(5427),
	197: uint32(5432),
	198: uint32(5439),
	199: uint32(5446),
	200: uint32(5453),
	201: uint32(5460),
	202: uint32(5467),
	203: uint32(5472),
	204: uint32(5476),
	205: uint32(5480),
	206: uint32(5484),
	207: uint32(5488),
	208: uint32(5492),
	209: uint32(5496),
	210: uint32(5500),
	211: uint32(5504),
	212: uint32(5508),
	213: uint32(5512),
	214: uint32(5516),
	215: uint32(5520),
	216: uint32(5524),
	217: uint32(5528),
	218: uint32(5532),
	219: uint32(5536),
	220: uint32(5540),
	221: uint32(5544),
	222: uint32(5548),
	223: uint32(5552),
	224: uint32(5556),
	225: uint32(5560),
	226: uint32(5564),
	227: uint32(5568),
	228: uint32(5572),
	229: uint32(5576),
	230: uint32(5580),
	231: uint32(5584),
	232: uint32(5588),
	233: uint32(5592),
	234: uint32(5596),
	235: uint32(5600),
	236: uint32(5604),
	237: uint32(5608),
	238: uint32(5612),
	239: uint32(5616),
	240: uint32(5620),
	241: uint32(5624),
	242: uint32(5628),
	243: uint32(5632),
	244: uint32(5636),
	245: uint32(5640),
	246: uint32(5644),
	247: uint32(5648),
	248: uint32(5652),
	249: uint32(5656),
	250: uint32(5660),
	251: uint32(5664),
	252: uint32(5668),
	253: uint32(5672),
	254: uint32(5676),
	255: uint32(5680),
	256: uint32(5684),
	257: uint32(5688),
	258: uint32(5692),
}

var ts_parse_actions = [1003]TSParseActionEntry{
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
		Fsymbol:     uint16(sym_document),
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
		Fstate: uint16(libc.Int32FromInt32(132)),
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
		Fstate: uint16(libc.Int32FromInt32(155)),
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
		Fstate: uint16(libc.Int32FromInt32(193)),
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
		Fstate: uint16(libc.Int32FromInt32(118)),
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
		Fstate: uint16(libc.Int32FromInt32(225)),
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
		Fstate: uint16(libc.Int32FromInt32(194)),
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
		Fstate: uint16(libc.Int32FromInt32(195)),
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
		Fstate: uint16(libc.Int32FromInt32(219)),
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
		Fstate: uint16(libc.Int32FromInt32(217)),
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
		Fstate: uint16(libc.Int32FromInt32(32)),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_request),
		Fproduction_id: uint16(19),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_request),
		Fproduction_id: uint16(19),
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
		Fstate: uint16(libc.Int32FromInt32(256)),
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
		Fstate: uint16(libc.Int32FromInt32(25)),
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
		Fstate: uint16(libc.Int32FromInt32(179)),
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
		Fstate: uint16(libc.Int32FromInt32(213)),
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
		Fstate: uint16(libc.Int32FromInt32(212)),
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
		Fstate: uint16(libc.Int32FromInt32(167)),
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
		Fstate: uint16(libc.Int32FromInt32(168)),
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
		Fstate: uint16(libc.Int32FromInt32(169)),
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
		Fstate: uint16(libc.Int32FromInt32(201)),
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
		Fstate: uint16(libc.Int32FromInt32(207)),
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
		Fstate: uint16(libc.Int32FromInt32(43)),
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
		Fstate: uint16(libc.Int32FromInt32(48)),
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
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_response),
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
		Fcount: uint8(1),
	}})),
	56: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_response),
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
		Fstate: uint16(libc.Int32FromInt32(26)),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_request),
		Fproduction_id: uint16(21),
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
		Fcount: uint8(1),
	}})),
	62: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_request),
		Fproduction_id: uint16(21),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(15)),
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
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_request),
		Fproduction_id: uint16(26),
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
		Fcount: uint8(1),
	}})),
	68: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_request),
		Fproduction_id: uint16(26),
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
		Fstate: uint16(libc.Int32FromInt32(14)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_request),
		Fproduction_id: uint16(13),
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
		Fcount: uint8(1),
	}})),
	74: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_request),
		Fproduction_id: uint16(13),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	78: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_request),
		Fproduction_id: uint16(10),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_request),
		Fproduction_id: uint16(10),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(17)),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_request),
		Fproduction_id: uint16(6),
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
		Fcount: uint8(1),
	}})),
	86: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_request),
		Fproduction_id: uint16(6),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(19)),
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
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_response),
		Fproduction_id: uint16(29),
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
		Fcount: uint8(1),
	}})),
	92: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_response),
		Fproduction_id: uint16(29),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(24)),
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
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_response),
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
		Fcount: uint8(1),
	}})),
	98: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_response),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(23)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	102: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_request),
		Fproduction_id: uint16(14),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_request),
		Fproduction_id: uint16(14),
	})))),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(22)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_request),
		Fproduction_id: uint16(34),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_request),
		Fproduction_id: uint16(34),
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
		Fstate: uint16(libc.Int32FromInt32(21)),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_response),
		Fproduction_id: uint16(38),
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
		Fcount: uint8(1),
	}})),
	116: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_response),
		Fproduction_id: uint16(38),
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
		Fstate: uint16(libc.Int32FromInt32(20)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_request),
		Fproduction_id: uint16(42),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_request),
		Fproduction_id: uint16(42),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(16)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_request),
		Fproduction_id: uint16(39),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_request),
		Fproduction_id: uint16(39),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym___body_repeat2),
		Fproduction_id: uint16(23),
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
		Fcount: uint8(1),
	}})),
	132: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym___body_repeat2),
		Fproduction_id: uint16(23),
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
		Fcount: uint8(2),
	}})),
	134: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym___body_repeat2),
		Fproduction_id: uint16(23),
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
		Fstate:      uint16(libc.Int32FromInt32(16)),
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
		Fcount: uint8(2),
	}})),
	137: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym___body_repeat2),
		Fproduction_id: uint16(23),
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
		Fstate:      uint16(libc.Int32FromInt32(179)),
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
		Fcount: uint8(2),
	}})),
	140: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym___body_repeat2),
		Fproduction_id: uint16(23),
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
		Fstate:      uint16(libc.Int32FromInt32(213)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	143: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym___body_repeat2),
		Fproduction_id: uint16(23),
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
		Fstate:      uint16(libc.Int32FromInt32(212)),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym___body_repeat2),
		Fproduction_id: uint16(23),
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
		Fstate:      uint16(libc.Int32FromInt32(167)),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym___body_repeat2),
		Fproduction_id: uint16(23),
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
		Fstate:      uint16(libc.Int32FromInt32(168)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	152: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym___body_repeat2),
		Fproduction_id: uint16(23),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(169)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym___body_repeat2),
		Fproduction_id: uint16(23),
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
		Fstate:      uint16(libc.Int32FromInt32(201)),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym___body_repeat2),
		Fproduction_id: uint16(23),
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
		Fstate:      uint16(libc.Int32FromInt32(207)),
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
		Fcount: uint8(2),
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
		Fsymbol:        uint16(aux_sym___body_repeat2),
		Fproduction_id: uint16(23),
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
		Fstate:      uint16(libc.Int32FromInt32(43)),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_request),
		Fproduction_id: uint16(24),
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
		Fcount: uint8(1),
	}})),
	166: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_request),
		Fproduction_id: uint16(24),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_request),
		Fproduction_id: uint16(27),
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
		Fcount: uint8(1),
	}})),
	170: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_request),
		Fproduction_id: uint16(27),
	})))),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_request),
		Fproduction_id: uint16(17),
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
		Fcount: uint8(1),
	}})),
	174: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_request),
		Fproduction_id: uint16(17),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(9),
		Fsymbol:        uint16(sym_response),
		Fproduction_id: uint16(48),
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
		Fcount: uint8(1),
	}})),
	178: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(9),
		Fsymbol:        uint16(sym_response),
		Fproduction_id: uint16(48),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	180: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(9),
		Fsymbol:        uint16(sym_request),
		Fproduction_id: uint16(47),
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
		Fcount: uint8(1),
	}})),
	182: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(9),
		Fsymbol:        uint16(sym_request),
		Fproduction_id: uint16(47),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_request),
		Fproduction_id: uint16(30),
	})))),
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
		Fcount: uint8(1),
	}})),
	186: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_request),
		Fproduction_id: uint16(30),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_response),
		Fproduction_id: uint16(44),
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
		Fcount: uint8(1),
	}})),
	190: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_response),
		Fproduction_id: uint16(44),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_response),
		Fproduction_id: uint16(43),
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
		Fcount: uint8(1),
	}})),
	194: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_response),
		Fproduction_id: uint16(43),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_request),
		Fproduction_id: uint16(35),
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
		Fcount: uint8(1),
	}})),
	198: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_request),
		Fproduction_id: uint16(35),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_response),
		Fproduction_id: uint16(37),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_response),
		Fproduction_id: uint16(37),
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
		Fsymbol:      uint16(aux_sym_document_repeat1),
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
		Fsymbol:      uint16(aux_sym_document_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(132)),
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
		Fsymbol:      uint16(aux_sym_document_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(155)),
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
		Fsymbol:      uint16(aux_sym_document_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(193)),
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
		Fsymbol:      uint16(aux_sym_document_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(118)),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_document_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(225)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_document_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(194)),
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
		Fsymbol:      uint16(aux_sym_document_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(195)),
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
		Fsymbol:      uint16(aux_sym_document_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(219)),
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
		Fsymbol:      uint16(aux_sym_document_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(217)),
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
		Fcount: uint8(2),
	}})),
	233: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      uint16(libc.Int32FromInt32(32)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_multipart_form_data),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_multipart_form_data),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(30)),
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
		Fstate: uint16(libc.Int32FromInt32(30)),
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
		Fsymbol:      uint16(sym_document),
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
		Fsymbol:      uint16(sym_multipart_form_data),
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
		Fsymbol:      uint16(sym_multipart_form_data),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fsymbol:      uint16(aux_sym_multipart_form_data_repeat1),
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
		Fsymbol:      uint16(aux_sym_multipart_form_data_repeat1),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_multipart_form_data_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(188)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_multipart_form_data_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(31)),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_multipart_form_data_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(201)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	269: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_multipart_form_data_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(31)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__section_content),
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
		Fcount: uint8(2),
	}})),
	274: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__section_content),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(132)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	277: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__section_content),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(155)),
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
		Fcount: uint8(2),
	}})),
	280: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__section_content),
	})))),
	281: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__section_content),
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
		Fstate: uint16(libc.Int32FromInt32(225)),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__section_content),
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
		Fstate: uint16(libc.Int32FromInt32(194)),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__section_content),
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
		Fstate: uint16(libc.Int32FromInt32(195)),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(sym__section_content),
	})))),
	293: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	294: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	295: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__section_content),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(217)),
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
		Fcount: uint8(2),
	}})),
	298: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__section_content),
	})))),
	299: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	300: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	301: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_section),
	})))),
	302: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	303: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__raw_body),
	})))),
	304: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	305: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__raw_body),
	})))),
	306: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	307: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__raw_body),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(235)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(35)),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__raw_body),
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
		Fcount: uint8(1),
	}})),
	314: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__raw_body),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(sym__raw_body),
	})))),
	317: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	318: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	319: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__plain_comment),
	})))),
	320: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	321: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__plain_comment),
	})))),
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
		Fcount: uint8(2),
	}})),
	323: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__raw_body),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(260)),
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
		Fstate: uint16(libc.Int32FromInt32(37)),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_comment),
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
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_comment),
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
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_comment),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_raw_body),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_raw_body),
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
		Fcount: uint8(2),
	}})),
	340: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_raw_body),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(260)),
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
		Fcount: uint8(2),
	}})),
	343: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__raw_body),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(260)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_raw_body),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_raw_body),
	})))),
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
		Fcount: uint8(2),
	}})),
	350: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_raw_body),
	})))),
	351: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	353: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym__var_comment),
		Fproduction_id: uint16(33),
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
		Fcount: uint8(1),
	}})),
	355: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym__var_comment),
		Fproduction_id: uint16(33),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym__var_comment),
		Fproduction_id: uint16(11),
	})))),
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
		Fcount: uint8(1),
	}})),
	359: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym__var_comment),
		Fproduction_id: uint16(11),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym__var_comment),
		Fproduction_id: uint16(25),
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
		Fcount: uint8(1),
	}})),
	363: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym__var_comment),
		Fproduction_id: uint16(25),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_graphql_body),
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
		Fcount: uint8(1),
	}})),
	367: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_graphql_body),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(176)),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym___body_repeat1),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym___body_repeat1),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym___body_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(256)),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym___body_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(48)),
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
		Fcount:    uint8(1),
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
		Fsymbol:      uint16(sym_graphql_data),
	})))),
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
		Fcount: uint8(1),
	}})),
	383: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_graphql_data),
	})))),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__raw_body),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__raw_body),
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
		Fcount:    uint8(1),
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
		Fsymbol:      uint16(sym_graphql_body),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_graphql_body),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_res_handler_script),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_res_handler_script),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_graphql_json_body),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_graphql_json_body),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__blank_line),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym__blank_line),
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
		Fcount:    uint8(1),
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
		Fsymbol:      uint16(sym__external_body),
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
		Fcount: uint8(1),
	}})),
	407: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__external_body),
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
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(aux_sym___body_repeat2),
		Fproduction_id: uint16(15),
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
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(aux_sym___body_repeat2),
		Fproduction_id: uint16(15),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_raw_body),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_raw_body),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_json_body),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_json_body),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_res_redirect),
		Fproduction_id: uint16(32),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_res_redirect),
		Fproduction_id: uint16(32),
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
		Fsymbol:      uint16(sym_xml_body),
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
		Fsymbol:      uint16(sym_xml_body),
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
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(aux_sym___body_repeat2),
		Fproduction_id: uint16(16),
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
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(aux_sym___body_repeat2),
		Fproduction_id: uint16(16),
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
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_request),
		Fproduction_id: uint16(34),
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
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_request),
		Fproduction_id: uint16(34),
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
		Fstate: uint16(libc.Int32FromInt32(256)),
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
		Fcount: uint8(1),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(178)),
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
		Fcount: uint8(1),
	}})),
	441: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_request),
		Fproduction_id: uint16(13),
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
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_request),
		Fproduction_id: uint16(13),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_request),
		Fproduction_id: uint16(21),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_request),
		Fproduction_id: uint16(21),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_request),
		Fproduction_id: uint16(26),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_request),
		Fproduction_id: uint16(26),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_request),
		Fproduction_id: uint16(10),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_request),
		Fproduction_id: uint16(10),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_request),
		Fproduction_id: uint16(6),
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
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_request),
		Fproduction_id: uint16(6),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_request),
		Fproduction_id: uint16(19),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_request),
		Fproduction_id: uint16(19),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(2)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_request),
		Fproduction_id: uint16(14),
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
		Fcount: uint8(1),
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
		Fsymbol:        uint16(sym_request),
		Fproduction_id: uint16(14),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_request_repeat1),
		Fproduction_id: uint16(18),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_request_repeat1),
		Fproduction_id: uint16(18),
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
		Fcount: uint8(2),
	}})),
	491: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_request_repeat1),
		Fproduction_id: uint16(18),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(178)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	494: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_request_repeat1),
		Fproduction_id: uint16(18),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(182)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_response),
		Fproduction_id: uint16(38),
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
		Fcount: uint8(1),
	}})),
	499: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_response),
		Fproduction_id: uint16(38),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_response),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_response),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(3)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_response),
		Fproduction_id: uint16(29),
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
		Fcount: uint8(1),
	}})),
	511: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_response),
		Fproduction_id: uint16(29),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(10)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_response_repeat1),
		Fproduction_id: uint16(18),
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
		Fcount: uint8(1),
	}})),
	519: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_response_repeat1),
		Fproduction_id: uint16(18),
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
		Fcount: uint8(2),
	}})),
	521: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_response_repeat1),
		Fproduction_id: uint16(18),
	})))),
	522: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(libc.Int32FromInt32(182)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_header),
		Fproduction_id: uint16(45),
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
		Fcount: uint8(1),
	}})),
	526: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_header),
		Fproduction_id: uint16(45),
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
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_header),
		Fproduction_id: uint16(31),
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
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_header),
		Fproduction_id: uint16(31),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(aux_sym_request_repeat1),
		Fproduction_id: uint16(9),
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
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(aux_sym_request_repeat1),
		Fproduction_id: uint16(9),
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
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(aux_sym_response_repeat1),
		Fproduction_id: uint16(9),
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
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(aux_sym_response_repeat1),
		Fproduction_id: uint16(9),
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
		Fsymbol:        uint16(sym_header),
		Fproduction_id: uint16(22),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_header),
		Fproduction_id: uint16(22),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_header),
		Fproduction_id: uint16(22),
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
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_header),
		Fproduction_id: uint16(22),
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
		Fsymbol:        uint16(sym_header),
		Fproduction_id: uint16(22),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_header),
		Fproduction_id: uint16(22),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_header),
		Fproduction_id: uint16(40),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_header),
		Fproduction_id: uint16(40),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_variable_declaration),
		Fproduction_id: uint16(28),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_variable_declaration),
		Fproduction_id: uint16(28),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_variable_declaration),
		Fproduction_id: uint16(36),
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
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_variable_declaration),
		Fproduction_id: uint16(36),
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
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_section),
		Fproduction_id: uint16(2),
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
		Fcount: uint8(1),
	}})),
	566: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_section),
		Fproduction_id: uint16(2),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_request_separator),
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
		Fcount: uint8(1),
	}})),
	570: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_request_separator),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_pre_request_script),
	})))),
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
		Fcount: uint8(1),
	}})),
	574: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_pre_request_script),
	})))),
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
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym__section_content),
		Fproduction_id: uint16(3),
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
		Fcount: uint8(1),
	}})),
	578: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym__section_content),
		Fproduction_id: uint16(3),
	})))),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym__section_content),
		Fproduction_id: uint16(4),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym__section_content),
		Fproduction_id: uint16(4),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_request_separator),
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
		Fcount: uint8(1),
	}})),
	586: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_request_separator),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_request_separator),
		Fproduction_id: uint16(12),
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
		Fcount: uint8(1),
	}})),
	590: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_request_separator),
		Fproduction_id: uint16(12),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_request_separator),
		Fproduction_id: uint16(7),
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
		Fcount: uint8(1),
	}})),
	594: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_request_separator),
		Fproduction_id: uint16(7),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_variable_declaration),
		Fproduction_id: uint16(20),
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
		Fcount: uint8(1),
	}})),
	598: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_variable_declaration),
		Fproduction_id: uint16(20),
	})))),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym__section_content),
		Fproduction_id: uint16(5),
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
		Fcount: uint8(1),
	}})),
	602: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym__section_content),
		Fproduction_id: uint16(5),
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
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_section),
		Fproduction_id: uint16(5),
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
		Fcount: uint8(1),
	}})),
	606: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_section),
		Fproduction_id: uint16(5),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(88)),
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
		Fstate: uint16(libc.Int32FromInt32(126)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(108)),
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
		Fstate: uint16(libc.Int32FromInt32(127)),
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
		Fstate: uint16(libc.Int32FromInt32(150)),
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
		Fstate: uint16(libc.Int32FromInt32(150)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(220)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(89)),
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
		Fstate: uint16(libc.Int32FromInt32(122)),
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
		Fstate: uint16(libc.Int32FromInt32(143)),
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
		Fstate: uint16(libc.Int32FromInt32(136)),
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
		Fstate: uint16(libc.Int32FromInt32(142)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(90)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(100)),
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
		Fstate: uint16(libc.Int32FromInt32(138)),
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
		Fstate: uint16(libc.Int32FromInt32(130)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(130)),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_path),
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
		Fcount: uint8(1),
	}})),
	650: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_path),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	652: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	654: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_path_repeat1),
	})))),
	655: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(2),
	}})),
	657: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_path_repeat1),
	})))),
	658: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_path_repeat1),
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
		Fcount: uint8(1),
	}})),
	662: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_path_repeat1),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_path_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(183)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(144)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(134)),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_target_url),
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
		Fstate: uint16(libc.Int32FromInt32(134)),
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
		Fcount: uint8(2),
	}})),
	676: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_target_url),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(259)),
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
		Fstate: uint16(libc.Int32FromInt32(140)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	681: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__target_url_line),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(134)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym__target_url_line),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(134)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(aux_sym__target_url_line),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym__target_url_line),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(195)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	692: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(aux_sym_target_url_repeat1),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(134)),
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
		Fcount: uint8(1),
	}})),
	695: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(aux_sym_target_url_repeat1),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(129)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	701: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_path_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(141)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	704: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_path_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(141)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_path_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(184)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(149)),
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
		Fcount: uint8(1),
	}})),
	714: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: uint16(libc.Int32FromInt32(132)),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_value_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(149)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	721: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_value_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(149)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(1),
	}})),
	724: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_value_repeat1),
	})))),
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
		Fsymbol:      uint16(aux_sym_value_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(195)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	729: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
	731: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_variable),
		Fproduction_id: uint16(8),
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
		Fcount: uint8(1),
	}})),
	735: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_variable),
		Fproduction_id: uint16(8),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_variable),
		Fproduction_id: uint16(11),
	})))),
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
		Fcount: uint8(1),
	}})),
	739: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_variable),
		Fproduction_id: uint16(11),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_variable),
		Fproduction_id: uint16(8),
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
		Fcount: uint8(1),
	}})),
	743: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_variable),
		Fproduction_id: uint16(8),
	})))),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_variable),
		Fproduction_id: uint16(11),
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
		Fcount: uint8(1),
	}})),
	747: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_variable),
		Fproduction_id: uint16(11),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym__target_url_line),
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
		Fcount: uint8(1),
	}})),
	751: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym__target_url_line),
	})))),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(106)),
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
		Fstate: uint16(libc.Int32FromInt32(135)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(135)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(177)),
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
		Fstate: uint16(libc.Int32FromInt32(229)),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_script_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(166)),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_script_repeat1),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(235)),
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
		Fstate: uint16(libc.Int32FromInt32(125)),
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
		Fstate: uint16(libc.Int32FromInt32(94)),
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
		Fstate: uint16(libc.Int32FromInt32(133)),
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
		Fstate: uint16(libc.Int32FromInt32(56)),
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
		Fstate: uint16(libc.Int32FromInt32(124)),
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
		Fstate: uint16(libc.Int32FromInt32(45)),
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
		Fstate: uint16(libc.Int32FromInt32(128)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_target_url),
	})))),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(sym_target_url),
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
		Fstate: uint16(libc.Int32FromInt32(259)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_target_url_repeat1),
	})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	795: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_target_url_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(259)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(166)),
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
		Fstate: uint16(libc.Int32FromInt32(230)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(84)),
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
		Fstate: uint16(libc.Int32FromInt32(249)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(41)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(148)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(215)),
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
		Fstate: uint16(libc.Int32FromInt32(121)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(255)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(187)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(250)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(192)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(131)),
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
		Fstate: uint16(libc.Int32FromInt32(242)),
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
		Fstate: uint16(libc.Int32FromInt32(154)),
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
		Fstate: uint16(libc.Int32FromInt32(241)),
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
		Fstate: uint16(libc.Int32FromInt32(151)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(36)),
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
		Fstate: uint16(libc.Int32FromInt32(258)),
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
		Fstate: uint16(libc.Int32FromInt32(218)),
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
		Fstate: uint16(libc.Int32FromInt32(164)),
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
		Fstate: uint16(libc.Int32FromInt32(233)),
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
		Fstate: uint16(libc.Int32FromInt32(159)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_external_body),
		Fproduction_id: uint16(41),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(206)),
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
		Fstate: uint16(libc.Int32FromInt32(160)),
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
		Fcount: uint8(1),
	}})),
	856: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(227)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:         uint8(TSParseActionTypeReduce),
		Fchild_count:        uint8(1),
		Fsymbol:             uint16(sym_http_version),
		Fdynamic_precedence: int16(1),
	})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(189)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_external_body),
		Fproduction_id: uint16(46),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(79)),
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
		Fstate: uint16(libc.Int32FromInt32(248)),
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
		Fstate: uint16(libc.Int32FromInt32(40)),
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
		Fstate: uint16(libc.Int32FromInt32(202)),
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
		Fstate: uint16(libc.Int32FromInt32(70)),
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
		Fstate: uint16(libc.Int32FromInt32(181)),
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
		Fstate: uint16(libc.Int32FromInt32(74)),
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
		Fstate: uint16(libc.Int32FromInt32(139)),
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
		Fstate: uint16(libc.Int32FromInt32(180)),
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
		Fstate: uint16(libc.Int32FromInt32(253)),
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
		Fstate: uint16(libc.Int32FromInt32(163)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_external_body),
		Fproduction_id: uint16(32),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(58)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(158)),
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
		Fstate: uint16(libc.Int32FromInt32(28)),
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
		Fstate: uint16(libc.Int32FromInt32(147)),
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
		Fstate: uint16(libc.Int32FromInt32(65)),
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
		Fstate: uint16(libc.Int32FromInt32(53)),
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
		Fstate: uint16(libc.Int32FromInt32(91)),
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
		Fstate: uint16(libc.Int32FromInt32(137)),
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
		Fstate: uint16(libc.Int32FromInt32(119)),
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
		Fstate: uint16(libc.Int32FromInt32(257)),
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
		Fstate: uint16(libc.Int32FromInt32(117)),
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
	914: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f uint8_t
		_ [7]byte
	}{f: uint8(TSParseActionTypeAccept)})),
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
		Fstate: uint16(libc.Int32FromInt32(185)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(120)),
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
		Fstate: uint16(libc.Int32FromInt32(165)),
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
		Fstate: uint16(libc.Int32FromInt32(103)),
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
		Fstate: uint16(libc.Int32FromInt32(98)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(72)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(113)),
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
		Fstate: uint16(libc.Int32FromInt32(173)),
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
		Fstate: uint16(libc.Int32FromInt32(85)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_script),
	})))),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_script),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(123)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	962: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	963: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	964: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	965: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	966: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	967: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	968: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	969: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	970: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	971: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	972: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	973: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	974: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	975: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	976: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	977: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	978: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	979: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	980: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	981: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	982: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	983: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	984: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	985: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	986: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	987: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	988: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	989: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	990: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	991: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	992: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	993: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	994: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	995: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	996: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	997: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	998: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	999: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1000: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1001: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1002: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
}

func tree_sitter_http(tls *libc.TLS) (r uintptr) {
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

var __ccgo_ts1 = "end\x00WORD_CHAR_token1\x00PUNCTUATION_token1\x00WS_token1\x00NL_token1\x00LINE_TAIL_token1\x00COMMENT_PREFIX_token1\x00@\x00=\x00_var_comment_token1\x00request_separator_token1\x00method\x00http_version_token1\x00status_code\x00status_text\x00:\x00{{\x00}}\x00<\x00pre_request_script_token1\x00>\x00{%\x00%}\x00res_redirect_token1\x00xml_body_token1\x00json_body_token1\x00graphql_data_token1\x00graphql_json_body_token1\x00--\x00multipart_form_data_token1\x00multipart_form_data_token2\x00raw_body_token1\x00_raw_body_token1\x00_not_comment\x00header_entity\x00identifier\x00path_token1\x00_blank_line_token1\x00document\x00comment\x00_plain_comment\x00_var_comment\x00request_separator\x00section\x00_section_content\x00http_version\x00_target_url_line\x00target_url\x00response\x00request\x00header\x00variable\x00pre_request_script\x00res_handler_script\x00script\x00res_redirect\x00variable_declaration\x00xml_body\x00json_body\x00graphql_body\x00graphql_data\x00_external_body\x00external_body\x00multipart_form_data\x00raw_body\x00_raw_body\x00path\x00value\x00_blank_line\x00document_repeat1\x00target_url_repeat1\x00__body_repeat1\x00__body_repeat2\x00response_repeat1\x00request_repeat1\x00script_repeat1\x00multipart_form_data_repeat1\x00path_repeat1\x00value_repeat1\x00body\x00name\x00url\x00version\x00"
