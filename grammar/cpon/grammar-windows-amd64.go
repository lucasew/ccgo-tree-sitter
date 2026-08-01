// Code generated for windows/amd64 by 'ccgo -ignore-unsupported-alignment -ignore-unsupported-atomic-sizes -ignore-vector-functions --winapi-no-errno -ignore-link-errors -std=gnu11 -O0 -D__extension__= -D__forceinline=static inline -D__attribute__(...)= -D__declspec(x)= -D__cdecl= -D__stdcall= -D__fastcall= -D__thiscall= -D_cdecl= -D__restrict= -D__restrict__= -D__MINGW_EXTENSION= -D_X86INTRIN_H_INCLUDED -D_X86GPRINTRIN_H_INCLUDED -D_IMMINTRIN_H_INCLUDED -D_MMINTRIN_H_INCLUDED -D_XMMINTRIN_H_INCLUDED -D_EMMINTRIN_H_INCLUDED -D_PMMINTRIN_H_INCLUDED -D_MM3DNOW_H_INCLUDED -Dfunc=func_token -Dinterface=interface_token -Dselect=select_token -Dchan=chan_token -Dgo=go_token -Dmap=map_token -Dpackage=package_token -Dtype=type_token -Dvar=var_token -Dimport=import_token -Ddefer=defer_token -Dfallthrough=fallthrough_token -Drange=range_token -I D:\a\ccgo-tree-sitter\ccgo-tree-sitter\third-party\tree-sitter-cpon\src -I D:\a\ccgo-tree-sitter\ccgo-tree-sitter\third-party\tree-sitter-cpon -I D:\a\ccgo-tree-sitter\ccgo-tree-sitter\third-party -I C:\Users\runneradmin\.cache\workspaced\sources\github\7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2\lib\include -I C:\Users\runneradmin\.cache\workspaced\sources\github\7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2\lib\src D:\a\ccgo-tree-sitter\ccgo-tree-sitter\third-party\tree-sitter-cpon\src\parser.c -o grammar.go', DO NOT EDIT.

//go:build windows && amd64

package grammar_cpon

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
const FIELD_COUNT = 2
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
const LARGE_STATE_COUNT = 12
const MAX_ALIAS_SEQUENCE_LENGTH = 6
const MB_LEN_MAX = 5
const MINGW_HAS_DDK_H = 1
const MINGW_HAS_SECURE_API = 1
const PATH_MAX = 260
const PRODUCTION_ID_COUNT = 2
const PTRDIFF_MAX = "INT64_MAX"
const PTRDIFF_MIN = "INT64_MIN"
const RAND_MAX = 0x7fff
const SIG_ATOMIC_MAX = "INT32_MAX"
const SIG_ATOMIC_MIN = "INT32_MIN"
const SIZE_MAX = "UINT64_MAX"
const SSIZE_MAX = "_I64_MAX"
const STATE_COUNT = 84
const STRUNCATE = 80
const SYMBOL_COUNT = 46
const TOKEN_COUNT = 25
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

const anon_sym_LT = 1
const anon_sym_COMMA = 2
const anon_sym_GT = 3
const anon_sym_COLON = 4
const anon_sym_LBRACE = 5
const anon_sym_RBRACE = 6
const anon_sym_i = 7
const anon_sym_LBRACK = 8
const anon_sym_RBRACK = 9
const anon_sym_DQUOTE = 10
const sym_string_content = 11
const aux_sym__escape_sequence_token1 = 12
const sym_escape_sequence = 13
const sym_number = 14
const sym_float = 15
const anon_sym_d = 16
const aux_sym_datetime_token1 = 17
const anon_sym_true = 18
const anon_sym_false = 19
const sym_null = 20
const anon_sym_x = 21
const aux_sym_hex_blob_token1 = 22
const anon_sym_b = 23
const sym_comment = 24
const sym_document = 25
const sym__value = 26
const sym_meta_map = 27
const sym_meta_pair = 28
const sym_map = 29
const sym_pair = 30
const sym_imap = 31
const sym_ipair = 32
const sym_array = 33
const sym_string = 34
const sym__escape_sequence = 35
const sym_datetime = 36
const sym_boolean = 37
const sym_hex_blob = 38
const sym_esc_blob = 39
const aux_sym_meta_map_repeat1 = 40
const aux_sym_meta_map_repeat2 = 41
const aux_sym_map_repeat1 = 42
const aux_sym_imap_repeat1 = 43
const aux_sym_array_repeat1 = 44
const aux_sym_string_repeat1 = 45

var ts_symbol_names = [46]uintptr{
	0:  __ccgo_ts,
	1:  __ccgo_ts + 4,
	2:  __ccgo_ts + 6,
	3:  __ccgo_ts + 8,
	4:  __ccgo_ts + 10,
	5:  __ccgo_ts + 12,
	6:  __ccgo_ts + 14,
	7:  __ccgo_ts + 16,
	8:  __ccgo_ts + 18,
	9:  __ccgo_ts + 20,
	10: __ccgo_ts + 22,
	11: __ccgo_ts + 24,
	12: __ccgo_ts + 39,
	13: __ccgo_ts + 63,
	14: __ccgo_ts + 79,
	15: __ccgo_ts + 86,
	16: __ccgo_ts + 92,
	17: __ccgo_ts + 94,
	18: __ccgo_ts + 110,
	19: __ccgo_ts + 115,
	20: __ccgo_ts + 121,
	21: __ccgo_ts + 126,
	22: __ccgo_ts + 128,
	23: __ccgo_ts + 144,
	24: __ccgo_ts + 146,
	25: __ccgo_ts + 154,
	26: __ccgo_ts + 163,
	27: __ccgo_ts + 170,
	28: __ccgo_ts + 179,
	29: __ccgo_ts + 189,
	30: __ccgo_ts + 193,
	31: __ccgo_ts + 198,
	32: __ccgo_ts + 203,
	33: __ccgo_ts + 209,
	34: __ccgo_ts + 215,
	35: __ccgo_ts + 222,
	36: __ccgo_ts + 239,
	37: __ccgo_ts + 248,
	38: __ccgo_ts + 256,
	39: __ccgo_ts + 265,
	40: __ccgo_ts + 274,
	41: __ccgo_ts + 291,
	42: __ccgo_ts + 308,
	43: __ccgo_ts + 320,
	44: __ccgo_ts + 333,
	45: __ccgo_ts + 347,
}

var ts_symbol_map = [46]TSSymbol{
	1:  uint16(anon_sym_LT),
	2:  uint16(anon_sym_COMMA),
	3:  uint16(anon_sym_GT),
	4:  uint16(anon_sym_COLON),
	5:  uint16(anon_sym_LBRACE),
	6:  uint16(anon_sym_RBRACE),
	7:  uint16(anon_sym_i),
	8:  uint16(anon_sym_LBRACK),
	9:  uint16(anon_sym_RBRACK),
	10: uint16(anon_sym_DQUOTE),
	11: uint16(sym_string_content),
	12: uint16(aux_sym__escape_sequence_token1),
	13: uint16(sym_escape_sequence),
	14: uint16(sym_number),
	15: uint16(sym_float),
	16: uint16(anon_sym_d),
	17: uint16(aux_sym_datetime_token1),
	18: uint16(anon_sym_true),
	19: uint16(anon_sym_false),
	20: uint16(sym_null),
	21: uint16(anon_sym_x),
	22: uint16(aux_sym_hex_blob_token1),
	23: uint16(anon_sym_b),
	24: uint16(sym_comment),
	25: uint16(sym_document),
	26: uint16(sym__value),
	27: uint16(sym_meta_map),
	28: uint16(sym_meta_pair),
	29: uint16(sym_map),
	30: uint16(sym_pair),
	31: uint16(sym_imap),
	32: uint16(sym_ipair),
	33: uint16(sym_array),
	34: uint16(sym_string),
	35: uint16(sym__escape_sequence),
	36: uint16(sym_datetime),
	37: uint16(sym_boolean),
	38: uint16(sym_hex_blob),
	39: uint16(sym_esc_blob),
	40: uint16(aux_sym_meta_map_repeat1),
	41: uint16(aux_sym_meta_map_repeat2),
	42: uint16(aux_sym_map_repeat1),
	43: uint16(aux_sym_imap_repeat1),
	44: uint16(aux_sym_array_repeat1),
	45: uint16(aux_sym_string_repeat1),
}

var ts_symbol_metadata = [46]TSSymbolMetadata{
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
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	16: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	17: {},
	18: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	19: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	20: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	21: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	22: {},
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
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	40: {},
	41: {},
	42: {},
	43: {},
	44: {},
	45: {},
}

type ts_field_identifiers = int32

const field_key = 1
const field_value = 2

var ts_field_names = [3]uintptr{
	0: libc.UintptrFromInt32(0),
	1: __ccgo_ts + 362,
	2: __ccgo_ts + 366,
}

var ts_field_map_slices = [2]TSFieldMapSlice{
	1: {
		Flength: uint16(2),
	},
}

var ts_field_map_entries = [2]TSFieldMapEntry{
	0: {
		Ffield_id: uint16(field_key),
	},
	1: {
		Ffield_id:    uint16(field_value),
		Fchild_index: uint8(2),
	},
}

var ts_alias_sequences = [2][6]TSSymbol{}

var ts_non_terminal_alias_map = [1]uint16_t{}

var ts_primary_state_ids = [84]TSStateId{
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
	54: uint16(54),
	55: uint16(55),
	56: uint16(56),
	57: uint16(57),
	58: uint16(58),
	59: uint16(59),
	60: uint16(60),
	61: uint16(61),
	62: uint16(62),
	63: uint16(63),
	64: uint16(64),
	65: uint16(65),
	66: uint16(66),
	67: uint16(67),
	68: uint16(68),
	69: uint16(69),
	70: uint16(70),
	71: uint16(71),
	72: uint16(72),
	73: uint16(73),
	74: uint16(74),
	75: uint16(75),
	76: uint16(76),
	77: uint16(77),
	78: uint16(78),
	79: uint16(79),
	80: uint16(80),
	81: uint16(81),
	82: uint16(82),
	83: uint16(83),
}

func ts_lex(tls *libc.TLS, lexer uintptr, state TSStateId) (r uint8) {
	var eof, result, skip uint8
	var i, i1 uint32_t
	var lookahead int32_t
	_, _, _, _, _, _ = eof, i, i1, lookahead, result, skip
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
			state = uint16(59)
			goto next_state
		}
		i = uint32(0)
		for {
			if !(uint64(i) < libc.Uint64FromInt64(88)/libc.Uint64FromInt64(2)) {
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
			state = uint16(58)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(81)
			goto next_state
		}
		return result
	case int32(1):
		if lookahead == int32('"') {
			state = uint16(69)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(71)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(24)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(70)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(74)
			goto next_state
		}
		return result
	case int32(2):
		if lookahead == int32('*') {
			state = uint16(4)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(99)
			goto next_state
		}
		return result
	case int32(3):
		if lookahead == int32('*') {
			state = uint16(3)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(97)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(4)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead == int32('*') {
			state = uint16(3)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(4)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead == int32('-') {
			state = uint16(9)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead == int32('-') {
			state = uint16(10)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead == int32('.') {
			state = uint16(39)
			goto next_state
		}
		if lookahead == int32('0') {
			state = uint16(82)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(82)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead == int32('/') {
			state = uint16(2)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(8)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(49)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead == int32('0') {
			state = uint16(38)
			goto next_state
		}
		if lookahead == int32('1') {
			state = uint16(32)
			goto next_state
		}
		return result
	case int32(10):
		if lookahead == int32('0') {
			state = uint16(36)
			goto next_state
		}
		if lookahead == int32('3') {
			state = uint16(30)
			goto next_state
		}
		if lookahead == int32('1') || lookahead == int32('2') {
			state = uint16(42)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead == int32('0') {
			state = uint16(37)
			goto next_state
		}
		if lookahead == int32('1') {
			state = uint16(31)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead == int32('0') {
			state = uint16(43)
			goto next_state
		}
		if lookahead == int32('1') {
			state = uint16(43)
			goto next_state
		}
		if lookahead == int32('2') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(13):
		if lookahead == int32(':') {
			state = uint16(35)
			goto next_state
		}
		return result
	case int32(14):
		if lookahead == int32(':') {
			state = uint16(34)
			goto next_state
		}
		return result
	case int32(15):
		if lookahead == int32('T') {
			state = uint16(12)
			goto next_state
		}
		return result
	case int32(16):
		if lookahead == int32('a') {
			state = uint16(19)
			goto next_state
		}
		return result
	case int32(17):
		if lookahead == int32('e') {
			state = uint16(91)
			goto next_state
		}
		return result
	case int32(18):
		if lookahead == int32('e') {
			state = uint16(92)
			goto next_state
		}
		return result
	case int32(19):
		if lookahead == int32('l') {
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(20):
		if lookahead == int32('l') {
			state = uint16(93)
			goto next_state
		}
		return result
	case int32(21):
		if lookahead == int32('l') {
			state = uint16(20)
			goto next_state
		}
		return result
	case int32(22):
		if lookahead == int32('r') {
			state = uint16(25)
			goto next_state
		}
		return result
	case int32(23):
		if lookahead == int32('s') {
			state = uint16(18)
			goto next_state
		}
		return result
	case int32(24):
		if lookahead == int32('u') {
			state = uint16(27)
			goto next_state
		}
		if lookahead == int32('x') {
			state = uint16(57)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('7') {
			state = uint16(76)
			goto next_state
		}
		if lookahead == int32('"') || lookahead == int32('\'') || lookahead == int32('?') || lookahead == int32('\\') || lookahead == int32('a') || lookahead == int32('b') || lookahead == int32('f') || lookahead == int32('n') || lookahead == int32('r') || int32('t') <= lookahead && lookahead <= int32('v') {
			state = uint16(77)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(75)
			goto next_state
		}
		return result
	case int32(25):
		if lookahead == int32('u') {
			state = uint16(17)
			goto next_state
		}
		return result
	case int32(26):
		if lookahead == int32('u') {
			state = uint16(21)
			goto next_state
		}
		return result
	case int32(27):
		if lookahead == int32('{') {
			state = uint16(54)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(56)
			goto next_state
		}
		return result
	case int32(28):
		if lookahead == int32('}') {
			state = uint16(77)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(28)
			goto next_state
		}
		return result
	case int32(29):
		if lookahead == int32('+') || lookahead == int32('-') {
			state = uint16(40)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(85)
			goto next_state
		}
		return result
	case int32(30):
		if lookahead == int32('0') || lookahead == int32('1') {
			state = uint16(15)
			goto next_state
		}
		return result
	case int32(31):
		if int32('0') <= lookahead && lookahead <= int32('2') {
			state = uint16(90)
			goto next_state
		}
		return result
	case int32(32):
		if int32('0') <= lookahead && lookahead <= int32('2') {
			state = uint16(6)
			goto next_state
		}
		return result
	case int32(33):
		if int32('0') <= lookahead && lookahead <= int32('3') {
			state = uint16(13)
			goto next_state
		}
		return result
	case int32(34):
		if int32('0') <= lookahead && lookahead <= int32('5') {
			state = uint16(44)
			goto next_state
		}
		return result
	case int32(35):
		if int32('0') <= lookahead && lookahead <= int32('5') {
			state = uint16(51)
			goto next_state
		}
		return result
	case int32(36):
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(15)
			goto next_state
		}
		return result
	case int32(37):
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(90)
			goto next_state
		}
		return result
	case int32(38):
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(6)
			goto next_state
		}
		return result
	case int32(39):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(84)
			goto next_state
		}
		return result
	case int32(40):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(85)
			goto next_state
		}
		return result
	case int32(41):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(5)
			goto next_state
		}
		return result
	case int32(42):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(15)
			goto next_state
		}
		return result
	case int32(43):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(13)
			goto next_state
		}
		return result
	case int32(44):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(88)
			goto next_state
		}
		return result
	case int32(45):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(87)
			goto next_state
		}
		return result
	case int32(46):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(89)
			goto next_state
		}
		return result
	case int32(47):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(41)
			goto next_state
		}
		return result
	case int32(48):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(46)
			goto next_state
		}
		return result
	case int32(49):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(47)
			goto next_state
		}
		return result
	case int32(50):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(48)
			goto next_state
		}
		return result
	case int32(51):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(14)
			goto next_state
		}
		return result
	case int32(52):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(77)
			goto next_state
		}
		return result
	case int32(53):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(83)
			goto next_state
		}
		return result
	case int32(54):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(28)
			goto next_state
		}
		return result
	case int32(55):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(95)
			goto next_state
		}
		return result
	case int32(56):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(57)
			goto next_state
		}
		return result
	case int32(57):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(58):
		if eof != 0 {
			state = uint16(59)
			goto next_state
		}
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
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(58)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(81)
			goto next_state
		}
		return result
	case int32(59):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(ts_builtin_sym_end)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(60):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(61):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COMMA)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(62):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_GT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(63):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COLON)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(64):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(65):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(66):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_i)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(67):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(68):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(69):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(70):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string_content)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('"') {
			state = uint16(69)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(71)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(70)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\\') {
			state = uint16(74)
			goto next_state
		}
		return result
	case int32(71):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string_content)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('*') {
			state = uint16(73)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(98)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('"') && lookahead != int32('\\') {
			state = uint16(74)
			goto next_state
		}
		return result
	case int32(72):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string_content)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('*') {
			state = uint16(72)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(97)
			goto next_state
		}
		if lookahead == int32('"') || lookahead == int32('\\') {
			state = uint16(4)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(73)
			goto next_state
		}
		return result
	case int32(73):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string_content)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('*') {
			state = uint16(72)
			goto next_state
		}
		if lookahead == int32('"') || lookahead == int32('\\') {
			state = uint16(4)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(73)
			goto next_state
		}
		return result
	case int32(74):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string_content)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('"') && lookahead != int32('\\') {
			state = uint16(74)
			goto next_state
		}
		return result
	case int32(75):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__escape_sequence_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(76):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__escape_sequence_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('7') {
			state = uint16(78)
			goto next_state
		}
		return result
	case int32(77):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_escape_sequence)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(78):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_escape_sequence)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('7') {
			state = uint16(77)
			goto next_state
		}
		return result
	case int32(79):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(80):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(39)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(79)
			goto next_state
		}
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(29)
			goto next_state
		}
		if lookahead == int32('X') || lookahead == int32('x') {
			state = uint16(53)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(81)
			goto next_state
		}
		return result
	case int32(81):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(39)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(79)
			goto next_state
		}
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(29)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(81)
			goto next_state
		}
		return result
	case int32(82):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(39)
			goto next_state
		}
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(29)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(82)
			goto next_state
		}
		return result
	case int32(83):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('u') {
			state = uint16(79)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(83)
			goto next_state
		}
		return result
	case int32(84):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_float)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(29)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(84)
			goto next_state
		}
		return result
	case int32(85):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_float)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(85)
			goto next_state
		}
		return result
	case int32(86):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_d)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(87):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_datetime_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(88):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_datetime_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(50)
			goto next_state
		}
		if lookahead == int32('Z') {
			state = uint16(87)
			goto next_state
		}
		if lookahead == int32('+') || lookahead == int32('-') {
			state = uint16(11)
			goto next_state
		}
		return result
	case int32(89):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_datetime_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('Z') {
			state = uint16(87)
			goto next_state
		}
		if lookahead == int32('+') || lookahead == int32('-') {
			state = uint16(11)
			goto next_state
		}
		return result
	case int32(90):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_datetime_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('5') {
			state = uint16(45)
			goto next_state
		}
		return result
	case int32(91):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_true)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(92):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_false)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(93):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_null)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(94):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_x)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(95):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_hex_blob_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(55)
			goto next_state
		}
		return result
	case int32(96):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_b)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(97):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(98):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('"') || lookahead == int32('\\') {
			state = uint16(99)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(98)
			goto next_state
		}
		return result
	case int32(99):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(99)
			goto next_state
		}
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var map_token = [44]uint16_t{
	0:  uint16('"'),
	1:  uint16(69),
	2:  uint16(','),
	3:  uint16(61),
	4:  uint16('.'),
	5:  uint16(39),
	6:  uint16('/'),
	7:  uint16(2),
	8:  uint16('0'),
	9:  uint16(80),
	10: uint16(':'),
	11: uint16(63),
	12: uint16('<'),
	13: uint16(60),
	14: uint16('>'),
	15: uint16(62),
	16: uint16('['),
	17: uint16(67),
	18: uint16('\\'),
	19: uint16(24),
	20: uint16(']'),
	21: uint16(68),
	22: uint16('b'),
	23: uint16(96),
	24: uint16('d'),
	25: uint16(86),
	26: uint16('f'),
	27: uint16(16),
	28: uint16('i'),
	29: uint16(66),
	30: uint16('n'),
	31: uint16(26),
	32: uint16('t'),
	33: uint16(22),
	34: uint16('x'),
	35: uint16(94),
	36: uint16('{'),
	37: uint16(64),
	38: uint16('}'),
	39: uint16(65),
	40: uint16('+'),
	41: uint16(7),
	42: uint16('-'),
	43: uint16(7),
}

var map_token1 = [42]uint16_t{
	0:  uint16('"'),
	1:  uint16(69),
	2:  uint16(','),
	3:  uint16(61),
	4:  uint16('.'),
	5:  uint16(39),
	6:  uint16('/'),
	7:  uint16(2),
	8:  uint16('0'),
	9:  uint16(80),
	10: uint16(':'),
	11: uint16(63),
	12: uint16('<'),
	13: uint16(60),
	14: uint16('>'),
	15: uint16(62),
	16: uint16('['),
	17: uint16(67),
	18: uint16(']'),
	19: uint16(68),
	20: uint16('b'),
	21: uint16(96),
	22: uint16('d'),
	23: uint16(86),
	24: uint16('f'),
	25: uint16(16),
	26: uint16('i'),
	27: uint16(66),
	28: uint16('n'),
	29: uint16(26),
	30: uint16('t'),
	31: uint16(22),
	32: uint16('x'),
	33: uint16(94),
	34: uint16('{'),
	35: uint16(64),
	36: uint16('}'),
	37: uint16(65),
	38: uint16('+'),
	39: uint16(7),
	40: uint16('-'),
	41: uint16(7),
}

var ts_lex_modes = [84]TSLexMode{
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
	40: {},
	41: {},
	42: {},
	43: {},
	44: {},
	45: {},
	46: {},
	47: {},
	48: {},
	49: {},
	50: {},
	51: {
		Flex_state: uint16(1),
	},
	52: {
		Flex_state: uint16(1),
	},
	53: {},
	54: {},
	55: {},
	56: {
		Flex_state: uint16(1),
	},
	57: {},
	58: {},
	59: {
		Flex_state: uint16(1),
	},
	60: {
		Flex_state: uint16(1),
	},
	61: {},
	62: {},
	63: {},
	64: {},
	65: {},
	66: {},
	67: {},
	68: {},
	69: {},
	70: {},
	71: {},
	72: {
		Flex_state: uint16(95),
	},
	73: {
		Flex_state: uint16(8),
	},
	74: {},
	75: {},
	76: {},
	77: {},
	78: {},
	79: {},
	80: {},
	81: {},
	82: {},
	83: {},
}

var ts_parse_table = [12][46]uint16_t{
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
		12: uint16(1),
		13: uint16(1),
		14: uint16(1),
		15: uint16(1),
		16: uint16(1),
		18: uint16(1),
		19: uint16(1),
		20: uint16(1),
		21: uint16(1),
		23: uint16(1),
		24: uint16(3),
	},
	1: {
		1:  uint16(5),
		5:  uint16(7),
		7:  uint16(9),
		8:  uint16(11),
		10: uint16(13),
		14: uint16(15),
		15: uint16(15),
		16: uint16(17),
		18: uint16(19),
		19: uint16(19),
		20: uint16(21),
		21: uint16(23),
		23: uint16(25),
		24: uint16(3),
		25: uint16(82),
		26: uint16(80),
		27: uint16(12),
		29: uint16(80),
		31: uint16(80),
		33: uint16(80),
		34: uint16(80),
		36: uint16(80),
		37: uint16(80),
		38: uint16(80),
		39: uint16(80),
	},
	2: {
		1:  uint16(27),
		2:  uint16(30),
		5:  uint16(33),
		7:  uint16(36),
		8:  uint16(39),
		9:  uint16(42),
		10: uint16(44),
		14: uint16(47),
		15: uint16(47),
		16: uint16(50),
		18: uint16(53),
		19: uint16(53),
		20: uint16(56),
		21: uint16(59),
		23: uint16(62),
		24: uint16(3),
		26: uint16(2),
		27: uint16(12),
		29: uint16(2),
		31: uint16(2),
		33: uint16(2),
		34: uint16(2),
		36: uint16(2),
		37: uint16(2),
		38: uint16(2),
		39: uint16(2),
		40: uint16(7),
		44: uint16(2),
	},
	3: {
		1:  uint16(5),
		2:  uint16(65),
		5:  uint16(7),
		7:  uint16(9),
		8:  uint16(11),
		9:  uint16(67),
		10: uint16(13),
		14: uint16(69),
		15: uint16(69),
		16: uint16(17),
		18: uint16(19),
		19: uint16(19),
		20: uint16(71),
		21: uint16(23),
		23: uint16(25),
		24: uint16(3),
		26: uint16(2),
		27: uint16(12),
		29: uint16(2),
		31: uint16(2),
		33: uint16(2),
		34: uint16(2),
		36: uint16(2),
		37: uint16(2),
		38: uint16(2),
		39: uint16(2),
		40: uint16(5),
		44: uint16(2),
	},
	4: {
		1:  uint16(5),
		2:  uint16(73),
		5:  uint16(7),
		7:  uint16(9),
		8:  uint16(11),
		9:  uint16(75),
		10: uint16(13),
		14: uint16(77),
		15: uint16(77),
		16: uint16(17),
		18: uint16(19),
		19: uint16(19),
		20: uint16(79),
		21: uint16(23),
		23: uint16(25),
		24: uint16(3),
		26: uint16(3),
		27: uint16(12),
		29: uint16(3),
		31: uint16(3),
		33: uint16(3),
		34: uint16(3),
		36: uint16(3),
		37: uint16(3),
		38: uint16(3),
		39: uint16(3),
		40: uint16(6),
		44: uint16(3),
	},
	5: {
		1:  uint16(5),
		2:  uint16(81),
		5:  uint16(7),
		7:  uint16(9),
		8:  uint16(11),
		9:  uint16(83),
		10: uint16(13),
		14: uint16(85),
		15: uint16(85),
		16: uint16(17),
		18: uint16(19),
		19: uint16(19),
		20: uint16(87),
		21: uint16(23),
		23: uint16(25),
		24: uint16(3),
		26: uint16(34),
		27: uint16(12),
		29: uint16(34),
		31: uint16(34),
		33: uint16(34),
		34: uint16(34),
		36: uint16(34),
		37: uint16(34),
		38: uint16(34),
		39: uint16(34),
		40: uint16(24),
	},
	6: {
		1:  uint16(5),
		2:  uint16(81),
		5:  uint16(7),
		7:  uint16(9),
		8:  uint16(11),
		9:  uint16(67),
		10: uint16(13),
		14: uint16(85),
		15: uint16(85),
		16: uint16(17),
		18: uint16(19),
		19: uint16(19),
		20: uint16(87),
		21: uint16(23),
		23: uint16(25),
		24: uint16(3),
		26: uint16(34),
		27: uint16(12),
		29: uint16(34),
		31: uint16(34),
		33: uint16(34),
		34: uint16(34),
		36: uint16(34),
		37: uint16(34),
		38: uint16(34),
		39: uint16(34),
		40: uint16(24),
	},
	7: {
		1:  uint16(5),
		2:  uint16(81),
		5:  uint16(7),
		7:  uint16(9),
		8:  uint16(11),
		10: uint16(13),
		14: uint16(85),
		15: uint16(85),
		16: uint16(17),
		18: uint16(19),
		19: uint16(19),
		20: uint16(87),
		21: uint16(23),
		23: uint16(25),
		24: uint16(3),
		26: uint16(34),
		27: uint16(12),
		29: uint16(34),
		31: uint16(34),
		33: uint16(34),
		34: uint16(34),
		36: uint16(34),
		37: uint16(34),
		38: uint16(34),
		39: uint16(34),
		40: uint16(24),
	},
	8: {
		1:  uint16(5),
		5:  uint16(7),
		7:  uint16(9),
		8:  uint16(11),
		9:  uint16(89),
		10: uint16(13),
		14: uint16(91),
		15: uint16(91),
		16: uint16(17),
		18: uint16(19),
		19: uint16(19),
		20: uint16(93),
		21: uint16(23),
		23: uint16(25),
		24: uint16(3),
		26: uint16(4),
		27: uint16(12),
		29: uint16(4),
		31: uint16(4),
		33: uint16(4),
		34: uint16(4),
		36: uint16(4),
		37: uint16(4),
		38: uint16(4),
		39: uint16(4),
	},
	9: {
		1:  uint16(5),
		5:  uint16(7),
		7:  uint16(9),
		8:  uint16(11),
		10: uint16(13),
		14: uint16(95),
		15: uint16(95),
		16: uint16(17),
		18: uint16(19),
		19: uint16(19),
		20: uint16(97),
		21: uint16(23),
		23: uint16(25),
		24: uint16(3),
		26: uint16(68),
		27: uint16(12),
		29: uint16(68),
		31: uint16(68),
		33: uint16(68),
		34: uint16(68),
		36: uint16(68),
		37: uint16(68),
		38: uint16(68),
		39: uint16(68),
	},
	10: {
		1:  uint16(5),
		5:  uint16(7),
		7:  uint16(9),
		8:  uint16(11),
		10: uint16(13),
		14: uint16(99),
		15: uint16(99),
		16: uint16(17),
		18: uint16(19),
		19: uint16(19),
		20: uint16(101),
		21: uint16(23),
		23: uint16(25),
		24: uint16(3),
		26: uint16(70),
		27: uint16(12),
		29: uint16(70),
		31: uint16(70),
		33: uint16(70),
		34: uint16(70),
		36: uint16(70),
		37: uint16(70),
		38: uint16(70),
		39: uint16(70),
	},
	11: {
		1:  uint16(5),
		5:  uint16(7),
		7:  uint16(9),
		8:  uint16(11),
		10: uint16(13),
		14: uint16(103),
		15: uint16(103),
		16: uint16(17),
		18: uint16(19),
		19: uint16(19),
		20: uint16(105),
		21: uint16(23),
		23: uint16(25),
		24: uint16(3),
		26: uint16(62),
		27: uint16(12),
		29: uint16(62),
		31: uint16(62),
		33: uint16(62),
		34: uint16(62),
		36: uint16(62),
		37: uint16(62),
		38: uint16(62),
		39: uint16(62),
	},
}

var ts_small_parse_table = [1398]uint16_t{
	0:    uint16(12),
	1:    uint16(3),
	2:    uint16(1),
	3:    uint16(sym_comment),
	4:    uint16(7),
	5:    uint16(1),
	6:    uint16(anon_sym_LBRACE),
	7:    uint16(9),
	8:    uint16(1),
	9:    uint16(anon_sym_i),
	10:   uint16(11),
	11:   uint16(1),
	12:   uint16(anon_sym_LBRACK),
	13:   uint16(13),
	14:   uint16(1),
	15:   uint16(anon_sym_DQUOTE),
	16:   uint16(17),
	17:   uint16(1),
	18:   uint16(anon_sym_d),
	19:   uint16(23),
	20:   uint16(1),
	21:   uint16(anon_sym_x),
	22:   uint16(25),
	23:   uint16(1),
	24:   uint16(anon_sym_b),
	25:   uint16(109),
	26:   uint16(1),
	27:   uint16(sym_null),
	28:   uint16(19),
	29:   uint16(2),
	30:   uint16(anon_sym_true),
	31:   uint16(anon_sym_false),
	32:   uint16(107),
	33:   uint16(2),
	34:   uint16(sym_number),
	35:   uint16(sym_float),
	36:   uint16(31),
	37:   uint16(8),
	38:   uint16(sym_map),
	39:   uint16(sym_imap),
	40:   uint16(sym_array),
	41:   uint16(sym_string),
	42:   uint16(sym_datetime),
	43:   uint16(sym_boolean),
	44:   uint16(sym_hex_blob),
	45:   uint16(sym_esc_blob),
	46:   uint16(3),
	47:   uint16(3),
	48:   uint16(1),
	49:   uint16(sym_comment),
	50:   uint16(113),
	51:   uint16(2),
	52:   uint16(sym_number),
	53:   uint16(sym_float),
	54:   uint16(111),
	55:   uint16(17),
	57:   uint16(anon_sym_LT),
	58:   uint16(anon_sym_COMMA),
	59:   uint16(anon_sym_GT),
	60:   uint16(anon_sym_COLON),
	61:   uint16(anon_sym_LBRACE),
	62:   uint16(anon_sym_RBRACE),
	63:   uint16(anon_sym_i),
	64:   uint16(anon_sym_LBRACK),
	65:   uint16(anon_sym_RBRACK),
	66:   uint16(anon_sym_DQUOTE),
	67:   uint16(anon_sym_d),
	68:   uint16(anon_sym_true),
	69:   uint16(anon_sym_false),
	70:   uint16(sym_null),
	71:   uint16(anon_sym_x),
	72:   uint16(anon_sym_b),
	73:   uint16(3),
	74:   uint16(3),
	75:   uint16(1),
	76:   uint16(sym_comment),
	77:   uint16(117),
	78:   uint16(2),
	79:   uint16(sym_number),
	80:   uint16(sym_float),
	81:   uint16(115),
	82:   uint16(17),
	84:   uint16(anon_sym_LT),
	85:   uint16(anon_sym_COMMA),
	86:   uint16(anon_sym_GT),
	87:   uint16(anon_sym_COLON),
	88:   uint16(anon_sym_LBRACE),
	89:   uint16(anon_sym_RBRACE),
	90:   uint16(anon_sym_i),
	91:   uint16(anon_sym_LBRACK),
	92:   uint16(anon_sym_RBRACK),
	93:   uint16(anon_sym_DQUOTE),
	94:   uint16(anon_sym_d),
	95:   uint16(anon_sym_true),
	96:   uint16(anon_sym_false),
	97:   uint16(sym_null),
	98:   uint16(anon_sym_x),
	99:   uint16(anon_sym_b),
	100:  uint16(3),
	101:  uint16(3),
	102:  uint16(1),
	103:  uint16(sym_comment),
	104:  uint16(121),
	105:  uint16(2),
	106:  uint16(sym_number),
	107:  uint16(sym_float),
	108:  uint16(119),
	109:  uint16(16),
	111:  uint16(anon_sym_LT),
	112:  uint16(anon_sym_COMMA),
	113:  uint16(anon_sym_GT),
	114:  uint16(anon_sym_LBRACE),
	115:  uint16(anon_sym_RBRACE),
	116:  uint16(anon_sym_i),
	117:  uint16(anon_sym_LBRACK),
	118:  uint16(anon_sym_RBRACK),
	119:  uint16(anon_sym_DQUOTE),
	120:  uint16(anon_sym_d),
	121:  uint16(anon_sym_true),
	122:  uint16(anon_sym_false),
	123:  uint16(sym_null),
	124:  uint16(anon_sym_x),
	125:  uint16(anon_sym_b),
	126:  uint16(3),
	127:  uint16(3),
	128:  uint16(1),
	129:  uint16(sym_comment),
	130:  uint16(125),
	131:  uint16(2),
	132:  uint16(sym_number),
	133:  uint16(sym_float),
	134:  uint16(123),
	135:  uint16(16),
	137:  uint16(anon_sym_LT),
	138:  uint16(anon_sym_COMMA),
	139:  uint16(anon_sym_GT),
	140:  uint16(anon_sym_LBRACE),
	141:  uint16(anon_sym_RBRACE),
	142:  uint16(anon_sym_i),
	143:  uint16(anon_sym_LBRACK),
	144:  uint16(anon_sym_RBRACK),
	145:  uint16(anon_sym_DQUOTE),
	146:  uint16(anon_sym_d),
	147:  uint16(anon_sym_true),
	148:  uint16(anon_sym_false),
	149:  uint16(sym_null),
	150:  uint16(anon_sym_x),
	151:  uint16(anon_sym_b),
	152:  uint16(3),
	153:  uint16(3),
	154:  uint16(1),
	155:  uint16(sym_comment),
	156:  uint16(129),
	157:  uint16(2),
	158:  uint16(sym_number),
	159:  uint16(sym_float),
	160:  uint16(127),
	161:  uint16(16),
	163:  uint16(anon_sym_LT),
	164:  uint16(anon_sym_COMMA),
	165:  uint16(anon_sym_GT),
	166:  uint16(anon_sym_LBRACE),
	167:  uint16(anon_sym_RBRACE),
	168:  uint16(anon_sym_i),
	169:  uint16(anon_sym_LBRACK),
	170:  uint16(anon_sym_RBRACK),
	171:  uint16(anon_sym_DQUOTE),
	172:  uint16(anon_sym_d),
	173:  uint16(anon_sym_true),
	174:  uint16(anon_sym_false),
	175:  uint16(sym_null),
	176:  uint16(anon_sym_x),
	177:  uint16(anon_sym_b),
	178:  uint16(3),
	179:  uint16(3),
	180:  uint16(1),
	181:  uint16(sym_comment),
	182:  uint16(133),
	183:  uint16(2),
	184:  uint16(sym_number),
	185:  uint16(sym_float),
	186:  uint16(131),
	187:  uint16(16),
	189:  uint16(anon_sym_LT),
	190:  uint16(anon_sym_COMMA),
	191:  uint16(anon_sym_GT),
	192:  uint16(anon_sym_LBRACE),
	193:  uint16(anon_sym_RBRACE),
	194:  uint16(anon_sym_i),
	195:  uint16(anon_sym_LBRACK),
	196:  uint16(anon_sym_RBRACK),
	197:  uint16(anon_sym_DQUOTE),
	198:  uint16(anon_sym_d),
	199:  uint16(anon_sym_true),
	200:  uint16(anon_sym_false),
	201:  uint16(sym_null),
	202:  uint16(anon_sym_x),
	203:  uint16(anon_sym_b),
	204:  uint16(3),
	205:  uint16(3),
	206:  uint16(1),
	207:  uint16(sym_comment),
	208:  uint16(137),
	209:  uint16(2),
	210:  uint16(sym_number),
	211:  uint16(sym_float),
	212:  uint16(135),
	213:  uint16(16),
	215:  uint16(anon_sym_LT),
	216:  uint16(anon_sym_COMMA),
	217:  uint16(anon_sym_GT),
	218:  uint16(anon_sym_LBRACE),
	219:  uint16(anon_sym_RBRACE),
	220:  uint16(anon_sym_i),
	221:  uint16(anon_sym_LBRACK),
	222:  uint16(anon_sym_RBRACK),
	223:  uint16(anon_sym_DQUOTE),
	224:  uint16(anon_sym_d),
	225:  uint16(anon_sym_true),
	226:  uint16(anon_sym_false),
	227:  uint16(sym_null),
	228:  uint16(anon_sym_x),
	229:  uint16(anon_sym_b),
	230:  uint16(3),
	231:  uint16(3),
	232:  uint16(1),
	233:  uint16(sym_comment),
	234:  uint16(141),
	235:  uint16(2),
	236:  uint16(sym_number),
	237:  uint16(sym_float),
	238:  uint16(139),
	239:  uint16(16),
	241:  uint16(anon_sym_LT),
	242:  uint16(anon_sym_COMMA),
	243:  uint16(anon_sym_GT),
	244:  uint16(anon_sym_LBRACE),
	245:  uint16(anon_sym_RBRACE),
	246:  uint16(anon_sym_i),
	247:  uint16(anon_sym_LBRACK),
	248:  uint16(anon_sym_RBRACK),
	249:  uint16(anon_sym_DQUOTE),
	250:  uint16(anon_sym_d),
	251:  uint16(anon_sym_true),
	252:  uint16(anon_sym_false),
	253:  uint16(sym_null),
	254:  uint16(anon_sym_x),
	255:  uint16(anon_sym_b),
	256:  uint16(3),
	257:  uint16(3),
	258:  uint16(1),
	259:  uint16(sym_comment),
	260:  uint16(145),
	261:  uint16(2),
	262:  uint16(sym_number),
	263:  uint16(sym_float),
	264:  uint16(143),
	265:  uint16(16),
	267:  uint16(anon_sym_LT),
	268:  uint16(anon_sym_COMMA),
	269:  uint16(anon_sym_GT),
	270:  uint16(anon_sym_LBRACE),
	271:  uint16(anon_sym_RBRACE),
	272:  uint16(anon_sym_i),
	273:  uint16(anon_sym_LBRACK),
	274:  uint16(anon_sym_RBRACK),
	275:  uint16(anon_sym_DQUOTE),
	276:  uint16(anon_sym_d),
	277:  uint16(anon_sym_true),
	278:  uint16(anon_sym_false),
	279:  uint16(sym_null),
	280:  uint16(anon_sym_x),
	281:  uint16(anon_sym_b),
	282:  uint16(3),
	283:  uint16(3),
	284:  uint16(1),
	285:  uint16(sym_comment),
	286:  uint16(149),
	287:  uint16(2),
	288:  uint16(sym_number),
	289:  uint16(sym_float),
	290:  uint16(147),
	291:  uint16(16),
	293:  uint16(anon_sym_LT),
	294:  uint16(anon_sym_COMMA),
	295:  uint16(anon_sym_GT),
	296:  uint16(anon_sym_LBRACE),
	297:  uint16(anon_sym_RBRACE),
	298:  uint16(anon_sym_i),
	299:  uint16(anon_sym_LBRACK),
	300:  uint16(anon_sym_RBRACK),
	301:  uint16(anon_sym_DQUOTE),
	302:  uint16(anon_sym_d),
	303:  uint16(anon_sym_true),
	304:  uint16(anon_sym_false),
	305:  uint16(sym_null),
	306:  uint16(anon_sym_x),
	307:  uint16(anon_sym_b),
	308:  uint16(3),
	309:  uint16(3),
	310:  uint16(1),
	311:  uint16(sym_comment),
	312:  uint16(153),
	313:  uint16(2),
	314:  uint16(sym_number),
	315:  uint16(sym_float),
	316:  uint16(151),
	317:  uint16(16),
	319:  uint16(anon_sym_LT),
	320:  uint16(anon_sym_COMMA),
	321:  uint16(anon_sym_GT),
	322:  uint16(anon_sym_LBRACE),
	323:  uint16(anon_sym_RBRACE),
	324:  uint16(anon_sym_i),
	325:  uint16(anon_sym_LBRACK),
	326:  uint16(anon_sym_RBRACK),
	327:  uint16(anon_sym_DQUOTE),
	328:  uint16(anon_sym_d),
	329:  uint16(anon_sym_true),
	330:  uint16(anon_sym_false),
	331:  uint16(sym_null),
	332:  uint16(anon_sym_x),
	333:  uint16(anon_sym_b),
	334:  uint16(5),
	335:  uint16(3),
	336:  uint16(1),
	337:  uint16(sym_comment),
	338:  uint16(157),
	339:  uint16(1),
	340:  uint16(anon_sym_COMMA),
	341:  uint16(24),
	342:  uint16(1),
	343:  uint16(aux_sym_meta_map_repeat1),
	344:  uint16(160),
	345:  uint16(2),
	346:  uint16(sym_number),
	347:  uint16(sym_float),
	348:  uint16(155),
	349:  uint16(14),
	350:  uint16(anon_sym_LT),
	351:  uint16(anon_sym_GT),
	352:  uint16(anon_sym_LBRACE),
	353:  uint16(anon_sym_RBRACE),
	354:  uint16(anon_sym_i),
	355:  uint16(anon_sym_LBRACK),
	356:  uint16(anon_sym_RBRACK),
	357:  uint16(anon_sym_DQUOTE),
	358:  uint16(anon_sym_d),
	359:  uint16(anon_sym_true),
	360:  uint16(anon_sym_false),
	361:  uint16(sym_null),
	362:  uint16(anon_sym_x),
	363:  uint16(anon_sym_b),
	364:  uint16(3),
	365:  uint16(3),
	366:  uint16(1),
	367:  uint16(sym_comment),
	368:  uint16(164),
	369:  uint16(2),
	370:  uint16(sym_number),
	371:  uint16(sym_float),
	372:  uint16(162),
	373:  uint16(16),
	375:  uint16(anon_sym_LT),
	376:  uint16(anon_sym_COMMA),
	377:  uint16(anon_sym_GT),
	378:  uint16(anon_sym_LBRACE),
	379:  uint16(anon_sym_RBRACE),
	380:  uint16(anon_sym_i),
	381:  uint16(anon_sym_LBRACK),
	382:  uint16(anon_sym_RBRACK),
	383:  uint16(anon_sym_DQUOTE),
	384:  uint16(anon_sym_d),
	385:  uint16(anon_sym_true),
	386:  uint16(anon_sym_false),
	387:  uint16(sym_null),
	388:  uint16(anon_sym_x),
	389:  uint16(anon_sym_b),
	390:  uint16(3),
	391:  uint16(3),
	392:  uint16(1),
	393:  uint16(sym_comment),
	394:  uint16(168),
	395:  uint16(2),
	396:  uint16(sym_number),
	397:  uint16(sym_float),
	398:  uint16(166),
	399:  uint16(16),
	401:  uint16(anon_sym_LT),
	402:  uint16(anon_sym_COMMA),
	403:  uint16(anon_sym_GT),
	404:  uint16(anon_sym_LBRACE),
	405:  uint16(anon_sym_RBRACE),
	406:  uint16(anon_sym_i),
	407:  uint16(anon_sym_LBRACK),
	408:  uint16(anon_sym_RBRACK),
	409:  uint16(anon_sym_DQUOTE),
	410:  uint16(anon_sym_d),
	411:  uint16(anon_sym_true),
	412:  uint16(anon_sym_false),
	413:  uint16(sym_null),
	414:  uint16(anon_sym_x),
	415:  uint16(anon_sym_b),
	416:  uint16(3),
	417:  uint16(3),
	418:  uint16(1),
	419:  uint16(sym_comment),
	420:  uint16(172),
	421:  uint16(2),
	422:  uint16(sym_number),
	423:  uint16(sym_float),
	424:  uint16(170),
	425:  uint16(16),
	427:  uint16(anon_sym_LT),
	428:  uint16(anon_sym_COMMA),
	429:  uint16(anon_sym_GT),
	430:  uint16(anon_sym_LBRACE),
	431:  uint16(anon_sym_RBRACE),
	432:  uint16(anon_sym_i),
	433:  uint16(anon_sym_LBRACK),
	434:  uint16(anon_sym_RBRACK),
	435:  uint16(anon_sym_DQUOTE),
	436:  uint16(anon_sym_d),
	437:  uint16(anon_sym_true),
	438:  uint16(anon_sym_false),
	439:  uint16(sym_null),
	440:  uint16(anon_sym_x),
	441:  uint16(anon_sym_b),
	442:  uint16(3),
	443:  uint16(3),
	444:  uint16(1),
	445:  uint16(sym_comment),
	446:  uint16(176),
	447:  uint16(2),
	448:  uint16(sym_number),
	449:  uint16(sym_float),
	450:  uint16(174),
	451:  uint16(16),
	453:  uint16(anon_sym_LT),
	454:  uint16(anon_sym_COMMA),
	455:  uint16(anon_sym_GT),
	456:  uint16(anon_sym_LBRACE),
	457:  uint16(anon_sym_RBRACE),
	458:  uint16(anon_sym_i),
	459:  uint16(anon_sym_LBRACK),
	460:  uint16(anon_sym_RBRACK),
	461:  uint16(anon_sym_DQUOTE),
	462:  uint16(anon_sym_d),
	463:  uint16(anon_sym_true),
	464:  uint16(anon_sym_false),
	465:  uint16(sym_null),
	466:  uint16(anon_sym_x),
	467:  uint16(anon_sym_b),
	468:  uint16(3),
	469:  uint16(3),
	470:  uint16(1),
	471:  uint16(sym_comment),
	472:  uint16(180),
	473:  uint16(2),
	474:  uint16(sym_number),
	475:  uint16(sym_float),
	476:  uint16(178),
	477:  uint16(16),
	479:  uint16(anon_sym_LT),
	480:  uint16(anon_sym_COMMA),
	481:  uint16(anon_sym_GT),
	482:  uint16(anon_sym_LBRACE),
	483:  uint16(anon_sym_RBRACE),
	484:  uint16(anon_sym_i),
	485:  uint16(anon_sym_LBRACK),
	486:  uint16(anon_sym_RBRACK),
	487:  uint16(anon_sym_DQUOTE),
	488:  uint16(anon_sym_d),
	489:  uint16(anon_sym_true),
	490:  uint16(anon_sym_false),
	491:  uint16(sym_null),
	492:  uint16(anon_sym_x),
	493:  uint16(anon_sym_b),
	494:  uint16(3),
	495:  uint16(3),
	496:  uint16(1),
	497:  uint16(sym_comment),
	498:  uint16(184),
	499:  uint16(2),
	500:  uint16(sym_number),
	501:  uint16(sym_float),
	502:  uint16(182),
	503:  uint16(16),
	505:  uint16(anon_sym_LT),
	506:  uint16(anon_sym_COMMA),
	507:  uint16(anon_sym_GT),
	508:  uint16(anon_sym_LBRACE),
	509:  uint16(anon_sym_RBRACE),
	510:  uint16(anon_sym_i),
	511:  uint16(anon_sym_LBRACK),
	512:  uint16(anon_sym_RBRACK),
	513:  uint16(anon_sym_DQUOTE),
	514:  uint16(anon_sym_d),
	515:  uint16(anon_sym_true),
	516:  uint16(anon_sym_false),
	517:  uint16(sym_null),
	518:  uint16(anon_sym_x),
	519:  uint16(anon_sym_b),
	520:  uint16(3),
	521:  uint16(3),
	522:  uint16(1),
	523:  uint16(sym_comment),
	524:  uint16(188),
	525:  uint16(2),
	526:  uint16(sym_number),
	527:  uint16(sym_float),
	528:  uint16(186),
	529:  uint16(16),
	531:  uint16(anon_sym_LT),
	532:  uint16(anon_sym_COMMA),
	533:  uint16(anon_sym_GT),
	534:  uint16(anon_sym_LBRACE),
	535:  uint16(anon_sym_RBRACE),
	536:  uint16(anon_sym_i),
	537:  uint16(anon_sym_LBRACK),
	538:  uint16(anon_sym_RBRACK),
	539:  uint16(anon_sym_DQUOTE),
	540:  uint16(anon_sym_d),
	541:  uint16(anon_sym_true),
	542:  uint16(anon_sym_false),
	543:  uint16(sym_null),
	544:  uint16(anon_sym_x),
	545:  uint16(anon_sym_b),
	546:  uint16(3),
	547:  uint16(3),
	548:  uint16(1),
	549:  uint16(sym_comment),
	550:  uint16(192),
	551:  uint16(2),
	552:  uint16(sym_number),
	553:  uint16(sym_float),
	554:  uint16(190),
	555:  uint16(16),
	557:  uint16(anon_sym_LT),
	558:  uint16(anon_sym_COMMA),
	559:  uint16(anon_sym_GT),
	560:  uint16(anon_sym_LBRACE),
	561:  uint16(anon_sym_RBRACE),
	562:  uint16(anon_sym_i),
	563:  uint16(anon_sym_LBRACK),
	564:  uint16(anon_sym_RBRACK),
	565:  uint16(anon_sym_DQUOTE),
	566:  uint16(anon_sym_d),
	567:  uint16(anon_sym_true),
	568:  uint16(anon_sym_false),
	569:  uint16(sym_null),
	570:  uint16(anon_sym_x),
	571:  uint16(anon_sym_b),
	572:  uint16(3),
	573:  uint16(3),
	574:  uint16(1),
	575:  uint16(sym_comment),
	576:  uint16(196),
	577:  uint16(2),
	578:  uint16(sym_number),
	579:  uint16(sym_float),
	580:  uint16(194),
	581:  uint16(16),
	583:  uint16(anon_sym_LT),
	584:  uint16(anon_sym_COMMA),
	585:  uint16(anon_sym_GT),
	586:  uint16(anon_sym_LBRACE),
	587:  uint16(anon_sym_RBRACE),
	588:  uint16(anon_sym_i),
	589:  uint16(anon_sym_LBRACK),
	590:  uint16(anon_sym_RBRACK),
	591:  uint16(anon_sym_DQUOTE),
	592:  uint16(anon_sym_d),
	593:  uint16(anon_sym_true),
	594:  uint16(anon_sym_false),
	595:  uint16(sym_null),
	596:  uint16(anon_sym_x),
	597:  uint16(anon_sym_b),
	598:  uint16(3),
	599:  uint16(3),
	600:  uint16(1),
	601:  uint16(sym_comment),
	602:  uint16(198),
	603:  uint16(2),
	604:  uint16(sym_number),
	605:  uint16(sym_float),
	606:  uint16(42),
	607:  uint16(13),
	608:  uint16(anon_sym_LT),
	609:  uint16(anon_sym_COMMA),
	610:  uint16(anon_sym_LBRACE),
	611:  uint16(anon_sym_i),
	612:  uint16(anon_sym_LBRACK),
	613:  uint16(anon_sym_RBRACK),
	614:  uint16(anon_sym_DQUOTE),
	615:  uint16(anon_sym_d),
	616:  uint16(anon_sym_true),
	617:  uint16(anon_sym_false),
	618:  uint16(sym_null),
	619:  uint16(anon_sym_x),
	620:  uint16(anon_sym_b),
	621:  uint16(3),
	622:  uint16(3),
	623:  uint16(1),
	624:  uint16(sym_comment),
	625:  uint16(202),
	626:  uint16(2),
	627:  uint16(sym_number),
	628:  uint16(sym_float),
	629:  uint16(200),
	630:  uint16(10),
	631:  uint16(anon_sym_LBRACE),
	632:  uint16(anon_sym_i),
	633:  uint16(anon_sym_LBRACK),
	634:  uint16(anon_sym_DQUOTE),
	635:  uint16(anon_sym_d),
	636:  uint16(anon_sym_true),
	637:  uint16(anon_sym_false),
	638:  uint16(sym_null),
	639:  uint16(anon_sym_x),
	640:  uint16(anon_sym_b),
	641:  uint16(3),
	642:  uint16(3),
	643:  uint16(1),
	644:  uint16(sym_comment),
	645:  uint16(206),
	646:  uint16(2),
	647:  uint16(sym_number),
	648:  uint16(sym_float),
	649:  uint16(204),
	650:  uint16(10),
	651:  uint16(anon_sym_LBRACE),
	652:  uint16(anon_sym_i),
	653:  uint16(anon_sym_LBRACK),
	654:  uint16(anon_sym_DQUOTE),
	655:  uint16(anon_sym_d),
	656:  uint16(anon_sym_true),
	657:  uint16(anon_sym_false),
	658:  uint16(sym_null),
	659:  uint16(anon_sym_x),
	660:  uint16(anon_sym_b),
	661:  uint16(3),
	662:  uint16(3),
	663:  uint16(1),
	664:  uint16(sym_comment),
	665:  uint16(210),
	666:  uint16(2),
	667:  uint16(sym_number),
	668:  uint16(sym_float),
	669:  uint16(208),
	670:  uint16(10),
	671:  uint16(anon_sym_LBRACE),
	672:  uint16(anon_sym_i),
	673:  uint16(anon_sym_LBRACK),
	674:  uint16(anon_sym_DQUOTE),
	675:  uint16(anon_sym_d),
	676:  uint16(anon_sym_true),
	677:  uint16(anon_sym_false),
	678:  uint16(sym_null),
	679:  uint16(anon_sym_x),
	680:  uint16(anon_sym_b),
	681:  uint16(3),
	682:  uint16(3),
	683:  uint16(1),
	684:  uint16(sym_comment),
	685:  uint16(214),
	686:  uint16(2),
	687:  uint16(sym_number),
	688:  uint16(sym_float),
	689:  uint16(212),
	690:  uint16(10),
	691:  uint16(anon_sym_LBRACE),
	692:  uint16(anon_sym_i),
	693:  uint16(anon_sym_LBRACK),
	694:  uint16(anon_sym_DQUOTE),
	695:  uint16(anon_sym_d),
	696:  uint16(anon_sym_true),
	697:  uint16(anon_sym_false),
	698:  uint16(sym_null),
	699:  uint16(anon_sym_x),
	700:  uint16(anon_sym_b),
	701:  uint16(8),
	702:  uint16(3),
	703:  uint16(1),
	704:  uint16(sym_comment),
	705:  uint16(13),
	706:  uint16(1),
	707:  uint16(anon_sym_DQUOTE),
	708:  uint16(216),
	709:  uint16(1),
	710:  uint16(anon_sym_COMMA),
	711:  uint16(218),
	712:  uint16(1),
	713:  uint16(anon_sym_GT),
	714:  uint16(43),
	715:  uint16(1),
	716:  uint16(aux_sym_meta_map_repeat1),
	717:  uint16(79),
	718:  uint16(1),
	719:  uint16(sym_string),
	720:  uint16(220),
	721:  uint16(2),
	722:  uint16(sym_number),
	723:  uint16(sym_float),
	724:  uint16(40),
	725:  uint16(2),
	726:  uint16(sym_meta_pair),
	727:  uint16(aux_sym_meta_map_repeat2),
	728:  uint16(8),
	729:  uint16(3),
	730:  uint16(1),
	731:  uint16(sym_comment),
	732:  uint16(222),
	733:  uint16(1),
	734:  uint16(anon_sym_COMMA),
	735:  uint16(225),
	736:  uint16(1),
	737:  uint16(anon_sym_GT),
	738:  uint16(227),
	739:  uint16(1),
	740:  uint16(anon_sym_DQUOTE),
	741:  uint16(50),
	742:  uint16(1),
	743:  uint16(aux_sym_meta_map_repeat1),
	744:  uint16(79),
	745:  uint16(1),
	746:  uint16(sym_string),
	747:  uint16(230),
	748:  uint16(2),
	749:  uint16(sym_number),
	750:  uint16(sym_float),
	751:  uint16(40),
	752:  uint16(2),
	753:  uint16(sym_meta_pair),
	754:  uint16(aux_sym_meta_map_repeat2),
	755:  uint16(8),
	756:  uint16(3),
	757:  uint16(1),
	758:  uint16(sym_comment),
	759:  uint16(13),
	760:  uint16(1),
	761:  uint16(anon_sym_DQUOTE),
	762:  uint16(233),
	763:  uint16(1),
	764:  uint16(anon_sym_COMMA),
	765:  uint16(235),
	766:  uint16(1),
	767:  uint16(anon_sym_GT),
	768:  uint16(42),
	769:  uint16(1),
	770:  uint16(aux_sym_meta_map_repeat1),
	771:  uint16(79),
	772:  uint16(1),
	773:  uint16(sym_string),
	774:  uint16(220),
	775:  uint16(2),
	776:  uint16(sym_number),
	777:  uint16(sym_float),
	778:  uint16(39),
	779:  uint16(2),
	780:  uint16(sym_meta_pair),
	781:  uint16(aux_sym_meta_map_repeat2),
	782:  uint16(8),
	783:  uint16(3),
	784:  uint16(1),
	785:  uint16(sym_comment),
	786:  uint16(13),
	787:  uint16(1),
	788:  uint16(anon_sym_DQUOTE),
	789:  uint16(81),
	790:  uint16(1),
	791:  uint16(anon_sym_COMMA),
	792:  uint16(218),
	793:  uint16(1),
	794:  uint16(anon_sym_GT),
	795:  uint16(24),
	796:  uint16(1),
	797:  uint16(aux_sym_meta_map_repeat1),
	798:  uint16(63),
	799:  uint16(1),
	800:  uint16(sym_meta_pair),
	801:  uint16(79),
	802:  uint16(1),
	803:  uint16(sym_string),
	804:  uint16(220),
	805:  uint16(2),
	806:  uint16(sym_number),
	807:  uint16(sym_float),
	808:  uint16(8),
	809:  uint16(3),
	810:  uint16(1),
	811:  uint16(sym_comment),
	812:  uint16(13),
	813:  uint16(1),
	814:  uint16(anon_sym_DQUOTE),
	815:  uint16(81),
	816:  uint16(1),
	817:  uint16(anon_sym_COMMA),
	818:  uint16(237),
	819:  uint16(1),
	820:  uint16(anon_sym_GT),
	821:  uint16(24),
	822:  uint16(1),
	823:  uint16(aux_sym_meta_map_repeat1),
	824:  uint16(63),
	825:  uint16(1),
	826:  uint16(sym_meta_pair),
	827:  uint16(79),
	828:  uint16(1),
	829:  uint16(sym_string),
	830:  uint16(220),
	831:  uint16(2),
	832:  uint16(sym_number),
	833:  uint16(sym_float),
	834:  uint16(7),
	835:  uint16(3),
	836:  uint16(1),
	837:  uint16(sym_comment),
	838:  uint16(13),
	839:  uint16(1),
	840:  uint16(anon_sym_DQUOTE),
	841:  uint16(239),
	842:  uint16(1),
	843:  uint16(anon_sym_COMMA),
	844:  uint16(241),
	845:  uint16(1),
	846:  uint16(anon_sym_RBRACE),
	847:  uint16(58),
	848:  uint16(1),
	849:  uint16(aux_sym_meta_map_repeat1),
	850:  uint16(76),
	851:  uint16(1),
	852:  uint16(sym_string),
	853:  uint16(48),
	854:  uint16(2),
	855:  uint16(sym_pair),
	856:  uint16(aux_sym_map_repeat1),
	857:  uint16(7),
	858:  uint16(3),
	859:  uint16(1),
	860:  uint16(sym_comment),
	861:  uint16(243),
	862:  uint16(1),
	863:  uint16(anon_sym_COMMA),
	864:  uint16(246),
	865:  uint16(1),
	866:  uint16(anon_sym_RBRACE),
	867:  uint16(248),
	868:  uint16(1),
	869:  uint16(anon_sym_DQUOTE),
	870:  uint16(64),
	871:  uint16(1),
	872:  uint16(aux_sym_meta_map_repeat1),
	873:  uint16(76),
	874:  uint16(1),
	875:  uint16(sym_string),
	876:  uint16(45),
	877:  uint16(2),
	878:  uint16(sym_pair),
	879:  uint16(aux_sym_map_repeat1),
	880:  uint16(6),
	881:  uint16(3),
	882:  uint16(1),
	883:  uint16(sym_comment),
	884:  uint16(251),
	885:  uint16(1),
	886:  uint16(anon_sym_COMMA),
	887:  uint16(253),
	888:  uint16(1),
	889:  uint16(anon_sym_RBRACE),
	890:  uint16(54),
	891:  uint16(1),
	892:  uint16(aux_sym_meta_map_repeat1),
	893:  uint16(255),
	894:  uint16(2),
	895:  uint16(sym_number),
	896:  uint16(sym_float),
	897:  uint16(49),
	898:  uint16(2),
	899:  uint16(sym_ipair),
	900:  uint16(aux_sym_imap_repeat1),
	901:  uint16(6),
	902:  uint16(3),
	903:  uint16(1),
	904:  uint16(sym_comment),
	905:  uint16(257),
	906:  uint16(1),
	907:  uint16(anon_sym_COMMA),
	908:  uint16(259),
	909:  uint16(1),
	910:  uint16(anon_sym_RBRACE),
	911:  uint16(55),
	912:  uint16(1),
	913:  uint16(aux_sym_meta_map_repeat1),
	914:  uint16(255),
	915:  uint16(2),
	916:  uint16(sym_number),
	917:  uint16(sym_float),
	918:  uint16(46),
	919:  uint16(2),
	920:  uint16(sym_ipair),
	921:  uint16(aux_sym_imap_repeat1),
	922:  uint16(7),
	923:  uint16(3),
	924:  uint16(1),
	925:  uint16(sym_comment),
	926:  uint16(13),
	927:  uint16(1),
	928:  uint16(anon_sym_DQUOTE),
	929:  uint16(261),
	930:  uint16(1),
	931:  uint16(anon_sym_COMMA),
	932:  uint16(263),
	933:  uint16(1),
	934:  uint16(anon_sym_RBRACE),
	935:  uint16(53),
	936:  uint16(1),
	937:  uint16(aux_sym_meta_map_repeat1),
	938:  uint16(76),
	939:  uint16(1),
	940:  uint16(sym_string),
	941:  uint16(45),
	942:  uint16(2),
	943:  uint16(sym_pair),
	944:  uint16(aux_sym_map_repeat1),
	945:  uint16(6),
	946:  uint16(3),
	947:  uint16(1),
	948:  uint16(sym_comment),
	949:  uint16(265),
	950:  uint16(1),
	951:  uint16(anon_sym_COMMA),
	952:  uint16(268),
	953:  uint16(1),
	954:  uint16(anon_sym_RBRACE),
	955:  uint16(61),
	956:  uint16(1),
	957:  uint16(aux_sym_meta_map_repeat1),
	958:  uint16(270),
	959:  uint16(2),
	960:  uint16(sym_number),
	961:  uint16(sym_float),
	962:  uint16(49),
	963:  uint16(2),
	964:  uint16(sym_ipair),
	965:  uint16(aux_sym_imap_repeat1),
	966:  uint16(7),
	967:  uint16(3),
	968:  uint16(1),
	969:  uint16(sym_comment),
	970:  uint16(13),
	971:  uint16(1),
	972:  uint16(anon_sym_DQUOTE),
	973:  uint16(81),
	974:  uint16(1),
	975:  uint16(anon_sym_COMMA),
	976:  uint16(24),
	977:  uint16(1),
	978:  uint16(aux_sym_meta_map_repeat1),
	979:  uint16(63),
	980:  uint16(1),
	981:  uint16(sym_meta_pair),
	982:  uint16(79),
	983:  uint16(1),
	984:  uint16(sym_string),
	985:  uint16(220),
	986:  uint16(2),
	987:  uint16(sym_number),
	988:  uint16(sym_float),
	989:  uint16(4),
	990:  uint16(273),
	991:  uint16(1),
	992:  uint16(anon_sym_DQUOTE),
	993:  uint16(277),
	994:  uint16(1),
	995:  uint16(sym_comment),
	996:  uint16(60),
	997:  uint16(2),
	998:  uint16(sym__escape_sequence),
	999:  uint16(aux_sym_string_repeat1),
	1000: uint16(275),
	1001: uint16(3),
	1002: uint16(sym_string_content),
	1003: uint16(aux_sym__escape_sequence_token1),
	1004: uint16(sym_escape_sequence),
	1005: uint16(4),
	1006: uint16(277),
	1007: uint16(1),
	1008: uint16(sym_comment),
	1009: uint16(279),
	1010: uint16(1),
	1011: uint16(anon_sym_DQUOTE),
	1012: uint16(56),
	1013: uint16(2),
	1014: uint16(sym__escape_sequence),
	1015: uint16(aux_sym_string_repeat1),
	1016: uint16(281),
	1017: uint16(3),
	1018: uint16(sym_string_content),
	1019: uint16(aux_sym__escape_sequence_token1),
	1020: uint16(sym_escape_sequence),
	1021: uint16(7),
	1022: uint16(3),
	1023: uint16(1),
	1024: uint16(sym_comment),
	1025: uint16(13),
	1026: uint16(1),
	1027: uint16(anon_sym_DQUOTE),
	1028: uint16(81),
	1029: uint16(1),
	1030: uint16(anon_sym_COMMA),
	1031: uint16(283),
	1032: uint16(1),
	1033: uint16(anon_sym_RBRACE),
	1034: uint16(24),
	1035: uint16(1),
	1036: uint16(aux_sym_meta_map_repeat1),
	1037: uint16(69),
	1038: uint16(1),
	1039: uint16(sym_pair),
	1040: uint16(76),
	1041: uint16(1),
	1042: uint16(sym_string),
	1043: uint16(6),
	1044: uint16(3),
	1045: uint16(1),
	1046: uint16(sym_comment),
	1047: uint16(81),
	1048: uint16(1),
	1049: uint16(anon_sym_COMMA),
	1050: uint16(285),
	1051: uint16(1),
	1052: uint16(anon_sym_RBRACE),
	1053: uint16(24),
	1054: uint16(1),
	1055: uint16(aux_sym_meta_map_repeat1),
	1056: uint16(67),
	1057: uint16(1),
	1058: uint16(sym_ipair),
	1059: uint16(255),
	1060: uint16(2),
	1061: uint16(sym_number),
	1062: uint16(sym_float),
	1063: uint16(6),
	1064: uint16(3),
	1065: uint16(1),
	1066: uint16(sym_comment),
	1067: uint16(81),
	1068: uint16(1),
	1069: uint16(anon_sym_COMMA),
	1070: uint16(253),
	1071: uint16(1),
	1072: uint16(anon_sym_RBRACE),
	1073: uint16(24),
	1074: uint16(1),
	1075: uint16(aux_sym_meta_map_repeat1),
	1076: uint16(67),
	1077: uint16(1),
	1078: uint16(sym_ipair),
	1079: uint16(255),
	1080: uint16(2),
	1081: uint16(sym_number),
	1082: uint16(sym_float),
	1083: uint16(4),
	1084: uint16(277),
	1085: uint16(1),
	1086: uint16(sym_comment),
	1087: uint16(287),
	1088: uint16(1),
	1089: uint16(anon_sym_DQUOTE),
	1090: uint16(56),
	1091: uint16(2),
	1092: uint16(sym__escape_sequence),
	1093: uint16(aux_sym_string_repeat1),
	1094: uint16(289),
	1095: uint16(3),
	1096: uint16(sym_string_content),
	1097: uint16(aux_sym__escape_sequence_token1),
	1098: uint16(sym_escape_sequence),
	1099: uint16(6),
	1100: uint16(3),
	1101: uint16(1),
	1102: uint16(sym_comment),
	1103: uint16(13),
	1104: uint16(1),
	1105: uint16(anon_sym_DQUOTE),
	1106: uint16(292),
	1107: uint16(1),
	1108: uint16(anon_sym_GT),
	1109: uint16(41),
	1110: uint16(1),
	1111: uint16(sym_meta_pair),
	1112: uint16(79),
	1113: uint16(1),
	1114: uint16(sym_string),
	1115: uint16(220),
	1116: uint16(2),
	1117: uint16(sym_number),
	1118: uint16(sym_float),
	1119: uint16(7),
	1120: uint16(3),
	1121: uint16(1),
	1122: uint16(sym_comment),
	1123: uint16(13),
	1124: uint16(1),
	1125: uint16(anon_sym_DQUOTE),
	1126: uint16(81),
	1127: uint16(1),
	1128: uint16(anon_sym_COMMA),
	1129: uint16(263),
	1130: uint16(1),
	1131: uint16(anon_sym_RBRACE),
	1132: uint16(24),
	1133: uint16(1),
	1134: uint16(aux_sym_meta_map_repeat1),
	1135: uint16(69),
	1136: uint16(1),
	1137: uint16(sym_pair),
	1138: uint16(76),
	1139: uint16(1),
	1140: uint16(sym_string),
	1141: uint16(4),
	1142: uint16(277),
	1143: uint16(1),
	1144: uint16(sym_comment),
	1145: uint16(294),
	1146: uint16(1),
	1147: uint16(anon_sym_DQUOTE),
	1148: uint16(52),
	1149: uint16(2),
	1150: uint16(sym__escape_sequence),
	1151: uint16(aux_sym_string_repeat1),
	1152: uint16(296),
	1153: uint16(3),
	1154: uint16(sym_string_content),
	1155: uint16(aux_sym__escape_sequence_token1),
	1156: uint16(sym_escape_sequence),
	1157: uint16(4),
	1158: uint16(277),
	1159: uint16(1),
	1160: uint16(sym_comment),
	1161: uint16(298),
	1162: uint16(1),
	1163: uint16(anon_sym_DQUOTE),
	1164: uint16(56),
	1165: uint16(2),
	1166: uint16(sym__escape_sequence),
	1167: uint16(aux_sym_string_repeat1),
	1168: uint16(281),
	1169: uint16(3),
	1170: uint16(sym_string_content),
	1171: uint16(aux_sym__escape_sequence_token1),
	1172: uint16(sym_escape_sequence),
	1173: uint16(5),
	1174: uint16(3),
	1175: uint16(1),
	1176: uint16(sym_comment),
	1177: uint16(81),
	1178: uint16(1),
	1179: uint16(anon_sym_COMMA),
	1180: uint16(24),
	1181: uint16(1),
	1182: uint16(aux_sym_meta_map_repeat1),
	1183: uint16(67),
	1184: uint16(1),
	1185: uint16(sym_ipair),
	1186: uint16(255),
	1187: uint16(2),
	1188: uint16(sym_number),
	1189: uint16(sym_float),
	1190: uint16(3),
	1191: uint16(3),
	1192: uint16(1),
	1193: uint16(sym_comment),
	1194: uint16(302),
	1195: uint16(2),
	1196: uint16(sym_number),
	1197: uint16(sym_float),
	1198: uint16(300),
	1199: uint16(3),
	1200: uint16(anon_sym_COMMA),
	1201: uint16(anon_sym_GT),
	1202: uint16(anon_sym_DQUOTE),
	1203: uint16(3),
	1204: uint16(3),
	1205: uint16(1),
	1206: uint16(sym_comment),
	1207: uint16(304),
	1208: uint16(2),
	1209: uint16(sym_number),
	1210: uint16(sym_float),
	1211: uint16(225),
	1212: uint16(3),
	1213: uint16(anon_sym_COMMA),
	1214: uint16(anon_sym_GT),
	1215: uint16(anon_sym_DQUOTE),
	1216: uint16(6),
	1217: uint16(3),
	1218: uint16(1),
	1219: uint16(sym_comment),
	1220: uint16(13),
	1221: uint16(1),
	1222: uint16(anon_sym_DQUOTE),
	1223: uint16(81),
	1224: uint16(1),
	1225: uint16(anon_sym_COMMA),
	1226: uint16(24),
	1227: uint16(1),
	1228: uint16(aux_sym_meta_map_repeat1),
	1229: uint16(69),
	1230: uint16(1),
	1231: uint16(sym_pair),
	1232: uint16(76),
	1233: uint16(1),
	1234: uint16(sym_string),
	1235: uint16(4),
	1236: uint16(3),
	1237: uint16(1),
	1238: uint16(sym_comment),
	1239: uint16(306),
	1240: uint16(1),
	1241: uint16(anon_sym_RBRACE),
	1242: uint16(47),
	1243: uint16(1),
	1244: uint16(sym_ipair),
	1245: uint16(255),
	1246: uint16(2),
	1247: uint16(sym_number),
	1248: uint16(sym_float),
	1249: uint16(5),
	1250: uint16(3),
	1251: uint16(1),
	1252: uint16(sym_comment),
	1253: uint16(13),
	1254: uint16(1),
	1255: uint16(anon_sym_DQUOTE),
	1256: uint16(308),
	1257: uint16(1),
	1258: uint16(anon_sym_RBRACE),
	1259: uint16(44),
	1260: uint16(1),
	1261: uint16(sym_pair),
	1262: uint16(76),
	1263: uint16(1),
	1264: uint16(sym_string),
	1265: uint16(3),
	1266: uint16(3),
	1267: uint16(1),
	1268: uint16(sym_comment),
	1269: uint16(268),
	1270: uint16(2),
	1271: uint16(anon_sym_COMMA),
	1272: uint16(anon_sym_RBRACE),
	1273: uint16(310),
	1274: uint16(2),
	1275: uint16(sym_number),
	1276: uint16(sym_float),
	1277: uint16(3),
	1278: uint16(3),
	1279: uint16(1),
	1280: uint16(sym_comment),
	1281: uint16(312),
	1282: uint16(2),
	1283: uint16(anon_sym_COMMA),
	1284: uint16(anon_sym_RBRACE),
	1285: uint16(314),
	1286: uint16(2),
	1287: uint16(sym_number),
	1288: uint16(sym_float),
	1289: uint16(2),
	1290: uint16(3),
	1291: uint16(1),
	1292: uint16(sym_comment),
	1293: uint16(246),
	1294: uint16(3),
	1295: uint16(anon_sym_COMMA),
	1296: uint16(anon_sym_RBRACE),
	1297: uint16(anon_sym_DQUOTE),
	1298: uint16(2),
	1299: uint16(3),
	1300: uint16(1),
	1301: uint16(sym_comment),
	1302: uint16(316),
	1303: uint16(3),
	1304: uint16(anon_sym_COMMA),
	1305: uint16(anon_sym_RBRACE),
	1306: uint16(anon_sym_DQUOTE),
	1307: uint16(2),
	1308: uint16(3),
	1309: uint16(1),
	1310: uint16(sym_comment),
	1311: uint16(318),
	1312: uint16(1),
	1313: uint16(anon_sym_DQUOTE),
	1314: uint16(2),
	1315: uint16(277),
	1316: uint16(1),
	1317: uint16(sym_comment),
	1318: uint16(320),
	1319: uint16(1),
	1320: uint16(aux_sym_hex_blob_token1),
	1321: uint16(2),
	1322: uint16(3),
	1323: uint16(1),
	1324: uint16(sym_comment),
	1325: uint16(322),
	1326: uint16(1),
	1327: uint16(aux_sym_datetime_token1),
	1328: uint16(2),
	1329: uint16(3),
	1330: uint16(1),
	1331: uint16(sym_comment),
	1332: uint16(324),
	1333: uint16(1),
	1334: uint16(anon_sym_DQUOTE),
	1335: uint16(2),
	1336: uint16(3),
	1337: uint16(1),
	1338: uint16(sym_comment),
	1339: uint16(326),
	1340: uint16(1),
	1341: uint16(anon_sym_LBRACE),
	1342: uint16(2),
	1343: uint16(3),
	1344: uint16(1),
	1345: uint16(sym_comment),
	1346: uint16(328),
	1347: uint16(1),
	1348: uint16(anon_sym_COLON),
	1349: uint16(2),
	1350: uint16(3),
	1351: uint16(1),
	1352: uint16(sym_comment),
	1353: uint16(330),
	1354: uint16(1),
	1355: uint16(anon_sym_DQUOTE),
	1356: uint16(2),
	1357: uint16(3),
	1358: uint16(1),
	1359: uint16(sym_comment),
	1360: uint16(332),
	1361: uint16(1),
	1362: uint16(anon_sym_COLON),
	1363: uint16(2),
	1364: uint16(3),
	1365: uint16(1),
	1366: uint16(sym_comment),
	1367: uint16(334),
	1368: uint16(1),
	1369: uint16(anon_sym_COLON),
	1370: uint16(2),
	1371: uint16(3),
	1372: uint16(1),
	1373: uint16(sym_comment),
	1374: uint16(336),
	1375: uint16(1),
	1377: uint16(2),
	1378: uint16(3),
	1379: uint16(1),
	1380: uint16(sym_comment),
	1381: uint16(338),
	1382: uint16(1),
	1383: uint16(anon_sym_DQUOTE),
	1384: uint16(2),
	1385: uint16(3),
	1386: uint16(1),
	1387: uint16(sym_comment),
	1388: uint16(340),
	1389: uint16(1),
	1391: uint16(2),
	1392: uint16(3),
	1393: uint16(1),
	1394: uint16(sym_comment),
	1395: uint16(342),
	1396: uint16(1),
	1397: uint16(anon_sym_DQUOTE),
}

var ts_small_parse_table_map = [72]uint32_t{
	1:  uint32(46),
	2:  uint32(73),
	3:  uint32(100),
	4:  uint32(126),
	5:  uint32(152),
	6:  uint32(178),
	7:  uint32(204),
	8:  uint32(230),
	9:  uint32(256),
	10: uint32(282),
	11: uint32(308),
	12: uint32(334),
	13: uint32(364),
	14: uint32(390),
	15: uint32(416),
	16: uint32(442),
	17: uint32(468),
	18: uint32(494),
	19: uint32(520),
	20: uint32(546),
	21: uint32(572),
	22: uint32(598),
	23: uint32(621),
	24: uint32(641),
	25: uint32(661),
	26: uint32(681),
	27: uint32(701),
	28: uint32(728),
	29: uint32(755),
	30: uint32(782),
	31: uint32(808),
	32: uint32(834),
	33: uint32(857),
	34: uint32(880),
	35: uint32(901),
	36: uint32(922),
	37: uint32(945),
	38: uint32(966),
	39: uint32(989),
	40: uint32(1005),
	41: uint32(1021),
	42: uint32(1043),
	43: uint32(1063),
	44: uint32(1083),
	45: uint32(1099),
	46: uint32(1119),
	47: uint32(1141),
	48: uint32(1157),
	49: uint32(1173),
	50: uint32(1190),
	51: uint32(1203),
	52: uint32(1216),
	53: uint32(1235),
	54: uint32(1249),
	55: uint32(1265),
	56: uint32(1277),
	57: uint32(1289),
	58: uint32(1298),
	59: uint32(1307),
	60: uint32(1314),
	61: uint32(1321),
	62: uint32(1328),
	63: uint32(1335),
	64: uint32(1342),
	65: uint32(1349),
	66: uint32(1356),
	67: uint32(1363),
	68: uint32(1370),
	69: uint32(1377),
	70: uint32(1384),
	71: uint32(1391),
}

var ts_parse_actions = [344]TSParseActionEntry{
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
		Fstate: uint16(libc.Int32FromInt32(57)),
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
		Fstate: uint16(libc.Int32FromInt32(66)),
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
		Fstate: uint16(libc.Int32FromInt32(75)),
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
		Fstate: uint16(libc.Int32FromInt32(8)),
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
		Fstate: uint16(libc.Int32FromInt32(51)),
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
		Fstate: uint16(libc.Int32FromInt32(80)),
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
		Fstate: uint16(libc.Int32FromInt32(71)),
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
		Fstate: uint16(libc.Int32FromInt32(26)),
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
		Fstate: uint16(libc.Int32FromInt32(80)),
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
		Fstate: uint16(libc.Int32FromInt32(77)),
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
		Fstate: uint16(libc.Int32FromInt32(74)),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_array_repeat1),
	})))),
	29: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(libc.Int32FromInt32(57)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	30: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	31: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      uint16(libc.Int32FromInt32(7)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_array_repeat1),
	})))),
	35: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(libc.Int32FromInt32(66)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	36: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	37: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      uint16(libc.Int32FromInt32(75)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_array_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(8)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_array_repeat1),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	45: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      uint16(libc.Int32FromInt32(51)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_array_repeat1),
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
		Fsymbol:      uint16(aux_sym_array_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(71)),
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
		Fsymbol:      uint16(aux_sym_array_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(26)),
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
		Fsymbol:      uint16(aux_sym_array_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(2)),
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
		Fsymbol:      uint16(aux_sym_array_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(77)),
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
		Fsymbol:      uint16(aux_sym_array_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(74)),
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
		Fstate: uint16(libc.Int32FromInt32(5)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(2)),
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
		Fstate: uint16(libc.Int32FromInt32(2)),
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
		Fstate: uint16(libc.Int32FromInt32(6)),
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
		Fstate: uint16(libc.Int32FromInt32(27)),
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
		Fstate: uint16(libc.Int32FromInt32(3)),
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
		Fstate: uint16(libc.Int32FromInt32(3)),
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
		Fstate: uint16(libc.Int32FromInt32(24)),
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
		Fstate: uint16(libc.Int32FromInt32(18)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(34)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(15)),
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
		Fstate: uint16(libc.Int32FromInt32(4)),
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
		Fstate: uint16(libc.Int32FromInt32(4)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(68)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(70)),
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
		Fstate: uint16(libc.Int32FromInt32(62)),
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
		Fstate: uint16(libc.Int32FromInt32(62)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(31)),
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
		Fstate: uint16(libc.Int32FromInt32(31)),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_string),
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
		Fcount: uint8(1),
	}})),
	114: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_string),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_array),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_array),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	124: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_datetime),
	})))),
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
		Fcount: uint8(1),
	}})),
	126: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_datetime),
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
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_imap),
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
		Fcount: uint8(1),
	}})),
	130: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_imap),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_array),
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
		Fcount: uint8(1),
	}})),
	134: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_esc_blob),
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
		Fcount: uint8(1),
	}})),
	138: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_esc_blob),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_map),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_map),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_imap),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_imap),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_imap),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_imap),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_map),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_map),
	})))),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_meta_map_repeat1),
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
		Fsymbol:      uint16(aux_sym_meta_map_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(24)),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(aux_sym_meta_map_repeat1),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_map),
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
		Fcount: uint8(1),
	}})),
	165: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_map),
	})))),
	166: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	167: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fcount: uint8(1),
	}})),
	169: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	171: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fcount: uint8(1),
	}})),
	173: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_esc_blob),
	})))),
	176: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	177: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_esc_blob),
	})))),
	178: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	179: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_hex_blob),
	})))),
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
		Fcount: uint8(1),
	}})),
	181: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_hex_blob),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	183: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fcount: uint8(1),
	}})),
	185: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fsymbol:      uint16(sym__value),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym__value),
	})))),
	190: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	191: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_imap),
	})))),
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
		Fcount: uint8(1),
	}})),
	193: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_imap),
	})))),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_map),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_map),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(aux_sym_array_repeat1),
	})))),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_meta_map),
	})))),
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
		Fcount: uint8(1),
	}})),
	203: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_meta_map),
	})))),
	204: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	205: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_meta_map),
	})))),
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
		Fcount: uint8(1),
	}})),
	207: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_meta_map),
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
		Fcount:    uint8(1),
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
		Fsymbol:      uint16(sym_meta_map),
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
		Fcount: uint8(1),
	}})),
	211: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_meta_map),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	213: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_meta_map),
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
		Fcount: uint8(1),
	}})),
	215: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_meta_map),
	})))),
	216: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	217: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	218: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(38)),
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
		Fcount: uint8(1),
	}})),
	221: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	222: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	223: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_meta_map_repeat2),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(50)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_meta_map_repeat2),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_meta_map_repeat2),
	})))),
	229: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	230: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	231: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_meta_map_repeat2),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(79)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(42)),
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
		Fstate: uint16(libc.Int32FromInt32(35)),
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
		Fstate: uint16(libc.Int32FromInt32(36)),
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
		Fstate: uint16(libc.Int32FromInt32(58)),
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
		Fstate: uint16(libc.Int32FromInt32(33)),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_map_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(64)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	247: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_map_repeat1),
	})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	249: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_map_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(51)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(54)),
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
		Fstate: uint16(libc.Int32FromInt32(21)),
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
		Fstate: uint16(libc.Int32FromInt32(78)),
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
		Fstate: uint16(libc.Int32FromInt32(55)),
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
		Fstate: uint16(libc.Int32FromInt32(22)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(20)),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_imap_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(61)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_imap_repeat1),
	})))),
	270: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	271: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_imap_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(78)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(14)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(60)),
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
		Fcount: uint8(1),
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
		Fextra: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(28)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(56)),
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
		Fstate: uint16(libc.Int32FromInt32(25)),
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
		Fstate: uint16(libc.Int32FromInt32(17)),
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
		Fsymbol:      uint16(aux_sym_string_repeat1),
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
		Fcount: uint8(2),
	}})),
	290: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	291: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(libc.Int32FromInt32(56)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	292: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(37)),
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
		Fcount: uint8(1),
	}})),
	295: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	296: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	297: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	298: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(13)),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_meta_pair),
		Fproduction_id: uint16(1),
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
		Fcount: uint8(1),
	}})),
	303: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_meta_pair),
		Fproduction_id: uint16(1),
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
		Fsymbol:      uint16(aux_sym_meta_map_repeat2),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	307: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	308: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	309: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	310: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	311: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_imap_repeat1),
	})))),
	312: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	313: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_ipair),
		Fproduction_id: uint16(1),
	})))),
	314: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	315: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_ipair),
		Fproduction_id: uint16(1),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	317: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_pair),
		Fproduction_id: uint16(1),
	})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(83)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	325: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	326: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	327: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	328: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(10)),
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
		Fstate: uint16(libc.Int32FromInt32(72)),
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
		Fstate: uint16(libc.Int32FromInt32(9)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	335: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fsymbol:      uint16(sym_document),
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
		Fstate: uint16(libc.Int32FromInt32(16)),
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
	341: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f uint8_t
		_ [7]byte
	}{f: uint8(TSParseActionTypeAccept)})),
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
		Fstate: uint16(libc.Int32FromInt32(29)),
	}})))),
}

func tree_sitter_cpon(tls *libc.TLS) (r uintptr) {
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

var __ccgo_ts1 = "end\x00<\x00,\x00>\x00:\x00{\x00}\x00i\x00[\x00]\x00\"\x00string_content\x00_escape_sequence_token1\x00escape_sequence\x00number\x00float\x00d\x00datetime_token1\x00true\x00false\x00null\x00x\x00hex_blob_token1\x00b\x00comment\x00document\x00_value\x00meta_map\x00meta_pair\x00map\x00pair\x00imap\x00ipair\x00array\x00string\x00_escape_sequence\x00datetime\x00boolean\x00hex_blob\x00esc_blob\x00meta_map_repeat1\x00meta_map_repeat2\x00map_repeat1\x00imap_repeat1\x00array_repeat1\x00string_repeat1\x00key\x00value\x00"
