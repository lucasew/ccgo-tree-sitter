// Code generated for windows/amd64 by 'ccgo -ignore-unsupported-alignment -ignore-unsupported-atomic-sizes -ignore-vector-functions --winapi-no-errno -ignore-link-errors -std=gnu11 -O0 -D__extension__= -D__forceinline=static inline -D__attribute__(...)= -D__declspec(x)= -D__cdecl= -D__stdcall= -D__fastcall= -D__thiscall= -D_cdecl= -D__restrict= -D__restrict__= -D__MINGW_EXTENSION= -D_X86INTRIN_H_INCLUDED -D_X86GPRINTRIN_H_INCLUDED -D_IMMINTRIN_H_INCLUDED -D_MMINTRIN_H_INCLUDED -D_XMMINTRIN_H_INCLUDED -D_EMMINTRIN_H_INCLUDED -D_PMMINTRIN_H_INCLUDED -D_MM3DNOW_H_INCLUDED -Dfunc=func_token -Dinterface=interface_token -Dselect=select_token -Dchan=chan_token -Dgo=go_token -Dmap=map_token -Dpackage=package_token -Dtype=type_token -Dvar=var_token -Dimport=import_token -Ddefer=defer_token -Dfallthrough=fallthrough_token -Drange=range_token -I D:\a\ccgo-tree-sitter\ccgo-tree-sitter\third-party\tree-sitter-smithy\src -I D:\a\ccgo-tree-sitter\ccgo-tree-sitter\third-party\tree-sitter-smithy -I D:\a\ccgo-tree-sitter\ccgo-tree-sitter\third-party -I C:\Users\runneradmin\.cache\workspaced\sources\github\7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2\lib\include -I C:\Users\runneradmin\.cache\workspaced\sources\github\7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2\lib\src D:\a\ccgo-tree-sitter\ccgo-tree-sitter\third-party\tree-sitter-smithy\src\parser.c -o grammar.go', DO NOT EDIT.

//go:build windows && amd64

package grammar_smithy

import (
	"reflect"
	"unsafe"

	"modernc.org/libc"
)

var _ reflect.Type
var _ unsafe.Pointer

const ALIAS_COUNT = 7
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
const FIELD_COUNT = 0
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
const MAX_ALIAS_SEQUENCE_LENGTH = 5
const MB_LEN_MAX = 5
const MINGW_HAS_DDK_H = 1
const MINGW_HAS_SECURE_API = 1
const PATH_MAX = 260
const PRODUCTION_ID_COUNT = 11
const PTRDIFF_MAX = "INT64_MAX"
const PTRDIFF_MIN = "INT64_MIN"
const RAND_MAX = 0x7fff
const SIG_ATOMIC_MAX = "INT32_MAX"
const SIG_ATOMIC_MIN = "INT32_MIN"
const SIZE_MAX = "UINT64_MAX"
const SSIZE_MAX = "_I64_MAX"
const STATE_COUNT = 350
const STRUNCATE = 80
const SYMBOL_COUNT = 145
const TOKEN_COUNT = 61
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

const anon_sym_DOLLAR = 1
const anon_sym_COLON = 2
const anon_sym_metadata = 3
const anon_sym_EQ = 4
const anon_sym_namespace = 5
const anon_sym_DOT = 6
const anon_sym_use = 7
const anon_sym_POUND = 8
const anon_sym_enum = 9
const anon_sym_intEnum = 10
const anon_sym_LBRACE = 11
const anon_sym_RBRACE = 12
const anon_sym_list = 13
const anon_sym_map = 14
const anon_sym_set = 15
const anon_sym_structure = 16
const anon_sym_union = 17
const anon_sym_service = 18
const anon_sym_operation = 19
const anon_sym_resource = 20
const anon_sym_COMMA = 21
const anon_sym_LBRACK = 22
const anon_sym_RBRACK = 23
const anon_sym_COLON_EQ = 24
const anon_sym_AT = 25
const anon_sym_LPAREN = 26
const anon_sym_RPAREN = 27
const anon_sym_apply = 28
const anon_sym_with = 29
const anon_sym_for = 30
const anon_sym_blob = 31
const anon_sym_boolean = 32
const anon_sym_byte = 33
const anon_sym_document = 34
const anon_sym_double = 35
const anon_sym_float = 36
const anon_sym_integer = 37
const anon_sym_long = 38
const anon_sym_short = 39
const anon_sym_string = 40
const anon_sym_timestamp = 41
const anon_sym_bigInteger = 42
const anon_sym_bigDecimal = 43
const anon_sym_true = 44
const anon_sym_false = 45
const sym_null = 46
const anon_sym_DASH = 47
const aux_sym_number_token1 = 48
const aux_sym_float_token1 = 49
const anon_sym_DQUOTE = 50
const anon_sym_DQUOTE_DQUOTE_DQUOTE = 51
const sym_string_fragment = 52
const aux_sym__multiline_string_fragment_token1 = 53
const aux_sym__multiline_string_fragment_token2 = 54
const aux_sym__escape_sequence_token1 = 55
const sym_escape_sequence = 56
const aux_sym_identifier_token1 = 57
const sym_comment = 58
const anon_sym_SLASH_SLASH_SLASH = 59
const aux_sym_documentation_comment_token1 = 60
const sym_source_file = 61
const sym_control_section = 62
const sym_control_statement = 63
const sym_control_var_name = 64
const sym_metadata_section = 65
const sym_metadata_statement = 66
const sym_shape_section = 67
const sym_namespace_statement = 68
const sym_namespace = 69
const sym__definition = 70
const sym_use_statement = 71
const sym_shape_statement = 72
const sym_shape_body = 73
const sym_absolute_root_shape_id = 74
const sym_root_shape_id = 75
const sym_shape_id_member = 76
const sym_shape_id = 77
const sym_simple_shape_statement = 78
const sym_enum_statement = 79
const sym_enum_members = 80
const sym_enum_member = 81
const sym_list_statement = 82
const sym_map_statement = 83
const sym_set_statement = 84
const sym_structure_statement = 85
const sym_union_statement = 86
const sym_service_statement = 87
const sym_operation_statement = 88
const sym_resource_statement = 89
const sym_shape_members = 90
const sym_shape_member = 91
const sym_shape_member_elided = 92
const sym_operation_body = 93
const sym_operation_member = 94
const sym_operation_errors = 95
const sym_operation_error = 96
const sym_inline_structure = 97
const sym_trait_statement = 98
const sym_trait_body = 99
const sym_trait_body_value = 100
const sym_trait_structure = 101
const sym_apply_statement = 102
const sym_apply_statement_singular = 103
const sym_apply_statement_block = 104
const sym_mixins = 105
const sym_structure_resource = 106
const sym_value_assignment = 107
const sym_node_value = 108
const sym_node_array = 109
const sym_node_object = 110
const sym_node_object_kvp = 111
const sym_node_object_key = 112
const sym_literal = 113
const sym_primitive = 114
const sym_boolean = 115
const sym_number = 116
const sym_float = 117
const sym_string = 118
const sym__string_literal = 119
const sym__multiline_string_literal = 120
const sym__multiline_string_fragment = 121
const sym__escape_sequence = 122
const sym_identifier = 123
const sym__control_identifier = 124
const sym__namespace_identifier = 125
const sym_documentation_comment = 126
const aux_sym_control_section_repeat1 = 127
const aux_sym_metadata_section_repeat1 = 128
const aux_sym_shape_section_repeat1 = 129
const aux_sym_namespace_repeat1 = 130
const aux_sym_shape_statement_repeat1 = 131
const aux_sym_shape_id_repeat1 = 132
const aux_sym_enum_members_repeat1 = 133
const aux_sym_shape_members_repeat1 = 134
const aux_sym_operation_body_repeat1 = 135
const aux_sym_operation_member_repeat1 = 136
const aux_sym_operation_errors_repeat1 = 137
const aux_sym_trait_structure_repeat1 = 138
const aux_sym_mixins_repeat1 = 139
const aux_sym_node_array_repeat1 = 140
const aux_sym_node_object_repeat1 = 141
const aux_sym__string_literal_repeat1 = 142
const aux_sym__multiline_string_literal_repeat1 = 143
const aux_sym__multiline_string_fragment_repeat1 = 144
const alias_sym_enum_field = 145
const alias_sym_field = 146
const alias_sym_key_identifier = 147
const alias_sym_operation_error_field = 148
const alias_sym_operation_field = 149
const alias_sym_trait_node_value = 150
const alias_sym_trait_object_kvp = 151

var ts_symbol_names = [152]uintptr{
	0:   __ccgo_ts,
	1:   __ccgo_ts + 4,
	2:   __ccgo_ts + 6,
	3:   __ccgo_ts + 8,
	4:   __ccgo_ts + 17,
	5:   __ccgo_ts + 19,
	6:   __ccgo_ts + 29,
	7:   __ccgo_ts + 31,
	8:   __ccgo_ts + 35,
	9:   __ccgo_ts + 37,
	10:  __ccgo_ts + 42,
	11:  __ccgo_ts + 50,
	12:  __ccgo_ts + 52,
	13:  __ccgo_ts + 54,
	14:  __ccgo_ts + 59,
	15:  __ccgo_ts + 63,
	16:  __ccgo_ts + 67,
	17:  __ccgo_ts + 77,
	18:  __ccgo_ts + 83,
	19:  __ccgo_ts + 91,
	20:  __ccgo_ts + 101,
	21:  __ccgo_ts + 110,
	22:  __ccgo_ts + 112,
	23:  __ccgo_ts + 114,
	24:  __ccgo_ts + 116,
	25:  __ccgo_ts + 119,
	26:  __ccgo_ts + 121,
	27:  __ccgo_ts + 123,
	28:  __ccgo_ts + 125,
	29:  __ccgo_ts + 131,
	30:  __ccgo_ts + 136,
	31:  __ccgo_ts + 140,
	32:  __ccgo_ts + 145,
	33:  __ccgo_ts + 153,
	34:  __ccgo_ts + 158,
	35:  __ccgo_ts + 167,
	36:  __ccgo_ts + 174,
	37:  __ccgo_ts + 180,
	38:  __ccgo_ts + 188,
	39:  __ccgo_ts + 193,
	40:  __ccgo_ts + 199,
	41:  __ccgo_ts + 206,
	42:  __ccgo_ts + 216,
	43:  __ccgo_ts + 227,
	44:  __ccgo_ts + 238,
	45:  __ccgo_ts + 243,
	46:  __ccgo_ts + 249,
	47:  __ccgo_ts + 254,
	48:  __ccgo_ts + 256,
	49:  __ccgo_ts + 270,
	50:  __ccgo_ts + 283,
	51:  __ccgo_ts + 285,
	52:  __ccgo_ts + 289,
	53:  __ccgo_ts + 305,
	54:  __ccgo_ts + 339,
	55:  __ccgo_ts + 373,
	56:  __ccgo_ts + 397,
	57:  __ccgo_ts + 413,
	58:  __ccgo_ts + 431,
	59:  __ccgo_ts + 439,
	60:  __ccgo_ts + 443,
	61:  __ccgo_ts + 472,
	62:  __ccgo_ts + 484,
	63:  __ccgo_ts + 500,
	64:  __ccgo_ts + 518,
	65:  __ccgo_ts + 535,
	66:  __ccgo_ts + 552,
	67:  __ccgo_ts + 571,
	68:  __ccgo_ts + 585,
	69:  __ccgo_ts + 19,
	70:  __ccgo_ts + 605,
	71:  __ccgo_ts + 617,
	72:  __ccgo_ts + 631,
	73:  __ccgo_ts + 647,
	74:  __ccgo_ts + 658,
	75:  __ccgo_ts + 681,
	76:  __ccgo_ts + 695,
	77:  __ccgo_ts + 711,
	78:  __ccgo_ts + 720,
	79:  __ccgo_ts + 743,
	80:  __ccgo_ts + 758,
	81:  __ccgo_ts + 771,
	82:  __ccgo_ts + 783,
	83:  __ccgo_ts + 798,
	84:  __ccgo_ts + 812,
	85:  __ccgo_ts + 826,
	86:  __ccgo_ts + 846,
	87:  __ccgo_ts + 862,
	88:  __ccgo_ts + 880,
	89:  __ccgo_ts + 900,
	90:  __ccgo_ts + 919,
	91:  __ccgo_ts + 933,
	92:  __ccgo_ts + 946,
	93:  __ccgo_ts + 966,
	94:  __ccgo_ts + 981,
	95:  __ccgo_ts + 998,
	96:  __ccgo_ts + 1015,
	97:  __ccgo_ts + 1031,
	98:  __ccgo_ts + 1048,
	99:  __ccgo_ts + 1064,
	100: __ccgo_ts + 1075,
	101: __ccgo_ts + 1092,
	102: __ccgo_ts + 1108,
	103: __ccgo_ts + 1124,
	104: __ccgo_ts + 1149,
	105: __ccgo_ts + 1171,
	106: __ccgo_ts + 1178,
	107: __ccgo_ts + 1197,
	108: __ccgo_ts + 1214,
	109: __ccgo_ts + 1225,
	110: __ccgo_ts + 1236,
	111: __ccgo_ts + 1248,
	112: __ccgo_ts + 1264,
	113: __ccgo_ts + 1280,
	114: __ccgo_ts + 1288,
	115: __ccgo_ts + 145,
	116: __ccgo_ts + 1298,
	117: __ccgo_ts + 174,
	118: __ccgo_ts + 199,
	119: __ccgo_ts + 1305,
	120: __ccgo_ts + 1321,
	121: __ccgo_ts + 1347,
	122: __ccgo_ts + 1373,
	123: __ccgo_ts + 1390,
	124: __ccgo_ts + 1401,
	125: __ccgo_ts + 1413,
	126: __ccgo_ts + 1435,
	127: __ccgo_ts + 1457,
	128: __ccgo_ts + 1481,
	129: __ccgo_ts + 1506,
	130: __ccgo_ts + 1528,
	131: __ccgo_ts + 1546,
	132: __ccgo_ts + 1570,
	133: __ccgo_ts + 1587,
	134: __ccgo_ts + 1608,
	135: __ccgo_ts + 1630,
	136: __ccgo_ts + 1653,
	137: __ccgo_ts + 1678,
	138: __ccgo_ts + 1703,
	139: __ccgo_ts + 1727,
	140: __ccgo_ts + 1742,
	141: __ccgo_ts + 1761,
	142: __ccgo_ts + 1781,
	143: __ccgo_ts + 1805,
	144: __ccgo_ts + 1839,
	145: __ccgo_ts + 1874,
	146: __ccgo_ts + 1885,
	147: __ccgo_ts + 1891,
	148: __ccgo_ts + 1906,
	149: __ccgo_ts + 1928,
	150: __ccgo_ts + 1944,
	151: __ccgo_ts + 1961,
}

var ts_symbol_map = [152]TSSymbol{
	1:   uint16(anon_sym_DOLLAR),
	2:   uint16(anon_sym_COLON),
	3:   uint16(anon_sym_metadata),
	4:   uint16(anon_sym_EQ),
	5:   uint16(anon_sym_namespace),
	6:   uint16(anon_sym_DOT),
	7:   uint16(anon_sym_use),
	8:   uint16(anon_sym_POUND),
	9:   uint16(anon_sym_enum),
	10:  uint16(anon_sym_intEnum),
	11:  uint16(anon_sym_LBRACE),
	12:  uint16(anon_sym_RBRACE),
	13:  uint16(anon_sym_list),
	14:  uint16(anon_sym_map),
	15:  uint16(anon_sym_set),
	16:  uint16(anon_sym_structure),
	17:  uint16(anon_sym_union),
	18:  uint16(anon_sym_service),
	19:  uint16(anon_sym_operation),
	20:  uint16(anon_sym_resource),
	21:  uint16(anon_sym_COMMA),
	22:  uint16(anon_sym_LBRACK),
	23:  uint16(anon_sym_RBRACK),
	24:  uint16(anon_sym_COLON_EQ),
	25:  uint16(anon_sym_AT),
	26:  uint16(anon_sym_LPAREN),
	27:  uint16(anon_sym_RPAREN),
	28:  uint16(anon_sym_apply),
	29:  uint16(anon_sym_with),
	30:  uint16(anon_sym_for),
	31:  uint16(anon_sym_blob),
	32:  uint16(anon_sym_boolean),
	33:  uint16(anon_sym_byte),
	34:  uint16(anon_sym_document),
	35:  uint16(anon_sym_double),
	36:  uint16(anon_sym_float),
	37:  uint16(anon_sym_integer),
	38:  uint16(anon_sym_long),
	39:  uint16(anon_sym_short),
	40:  uint16(anon_sym_string),
	41:  uint16(anon_sym_timestamp),
	42:  uint16(anon_sym_bigInteger),
	43:  uint16(anon_sym_bigDecimal),
	44:  uint16(anon_sym_true),
	45:  uint16(anon_sym_false),
	46:  uint16(sym_null),
	47:  uint16(anon_sym_DASH),
	48:  uint16(aux_sym_number_token1),
	49:  uint16(aux_sym_float_token1),
	50:  uint16(anon_sym_DQUOTE),
	51:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	52:  uint16(sym_string_fragment),
	53:  uint16(aux_sym__multiline_string_fragment_token1),
	54:  uint16(aux_sym__multiline_string_fragment_token2),
	55:  uint16(aux_sym__escape_sequence_token1),
	56:  uint16(sym_escape_sequence),
	57:  uint16(aux_sym_identifier_token1),
	58:  uint16(sym_comment),
	59:  uint16(anon_sym_SLASH_SLASH_SLASH),
	60:  uint16(aux_sym_documentation_comment_token1),
	61:  uint16(sym_source_file),
	62:  uint16(sym_control_section),
	63:  uint16(sym_control_statement),
	64:  uint16(sym_control_var_name),
	65:  uint16(sym_metadata_section),
	66:  uint16(sym_metadata_statement),
	67:  uint16(sym_shape_section),
	68:  uint16(sym_namespace_statement),
	69:  uint16(sym_namespace),
	70:  uint16(sym__definition),
	71:  uint16(sym_use_statement),
	72:  uint16(sym_shape_statement),
	73:  uint16(sym_shape_body),
	74:  uint16(sym_absolute_root_shape_id),
	75:  uint16(sym_root_shape_id),
	76:  uint16(sym_shape_id_member),
	77:  uint16(sym_shape_id),
	78:  uint16(sym_simple_shape_statement),
	79:  uint16(sym_enum_statement),
	80:  uint16(sym_enum_members),
	81:  uint16(sym_enum_member),
	82:  uint16(sym_list_statement),
	83:  uint16(sym_map_statement),
	84:  uint16(sym_set_statement),
	85:  uint16(sym_structure_statement),
	86:  uint16(sym_union_statement),
	87:  uint16(sym_service_statement),
	88:  uint16(sym_operation_statement),
	89:  uint16(sym_resource_statement),
	90:  uint16(sym_shape_members),
	91:  uint16(sym_shape_member),
	92:  uint16(sym_shape_member_elided),
	93:  uint16(sym_operation_body),
	94:  uint16(sym_operation_member),
	95:  uint16(sym_operation_errors),
	96:  uint16(sym_operation_error),
	97:  uint16(sym_inline_structure),
	98:  uint16(sym_trait_statement),
	99:  uint16(sym_trait_body),
	100: uint16(sym_trait_body_value),
	101: uint16(sym_trait_structure),
	102: uint16(sym_apply_statement),
	103: uint16(sym_apply_statement_singular),
	104: uint16(sym_apply_statement_block),
	105: uint16(sym_mixins),
	106: uint16(sym_structure_resource),
	107: uint16(sym_value_assignment),
	108: uint16(sym_node_value),
	109: uint16(sym_node_array),
	110: uint16(sym_node_object),
	111: uint16(sym_node_object_kvp),
	112: uint16(sym_node_object_key),
	113: uint16(sym_literal),
	114: uint16(sym_primitive),
	115: uint16(sym_boolean),
	116: uint16(sym_number),
	117: uint16(sym_float),
	118: uint16(sym_string),
	119: uint16(sym__string_literal),
	120: uint16(sym__multiline_string_literal),
	121: uint16(sym__multiline_string_fragment),
	122: uint16(sym__escape_sequence),
	123: uint16(sym_identifier),
	124: uint16(sym__control_identifier),
	125: uint16(sym__namespace_identifier),
	126: uint16(sym_documentation_comment),
	127: uint16(aux_sym_control_section_repeat1),
	128: uint16(aux_sym_metadata_section_repeat1),
	129: uint16(aux_sym_shape_section_repeat1),
	130: uint16(aux_sym_namespace_repeat1),
	131: uint16(aux_sym_shape_statement_repeat1),
	132: uint16(aux_sym_shape_id_repeat1),
	133: uint16(aux_sym_enum_members_repeat1),
	134: uint16(aux_sym_shape_members_repeat1),
	135: uint16(aux_sym_operation_body_repeat1),
	136: uint16(aux_sym_operation_member_repeat1),
	137: uint16(aux_sym_operation_errors_repeat1),
	138: uint16(aux_sym_trait_structure_repeat1),
	139: uint16(aux_sym_mixins_repeat1),
	140: uint16(aux_sym_node_array_repeat1),
	141: uint16(aux_sym_node_object_repeat1),
	142: uint16(aux_sym__string_literal_repeat1),
	143: uint16(aux_sym__multiline_string_literal_repeat1),
	144: uint16(aux_sym__multiline_string_fragment_repeat1),
	145: uint16(alias_sym_enum_field),
	146: uint16(alias_sym_field),
	147: uint16(alias_sym_key_identifier),
	148: uint16(alias_sym_operation_error_field),
	149: uint16(alias_sym_operation_field),
	150: uint16(alias_sym_trait_node_value),
	151: uint16(alias_sym_trait_object_kvp),
}

var ts_symbol_metadata = [152]TSSymbolMetadata{
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
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	47: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	48: {},
	49: {},
	50: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	51: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	52: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	53: {},
	54: {},
	55: {},
	56: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	57: {},
	58: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	59: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	60: {},
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
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	96: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	102: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	120: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	121: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	122: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	123: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	124: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	125: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	126: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	127: {},
	128: {},
	129: {},
	130: {},
	131: {},
	132: {},
	133: {},
	134: {},
	135: {},
	136: {},
	137: {},
	138: {},
	139: {},
	140: {},
	141: {},
	142: {},
	143: {},
	144: {},
	145: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	146: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
}

var ts_alias_sequences = [11][5]TSSymbol{
	0: {},
	1: {
		1: uint16(sym__control_identifier),
	},
	2: {
		0: uint16(alias_sym_key_identifier),
	},
	3: {
		0: uint16(alias_sym_enum_field),
	},
	4: {
		0: uint16(alias_sym_trait_node_value),
	},
	5: {
		0: uint16(alias_sym_trait_object_kvp),
	},
	6: {
		1: uint16(alias_sym_enum_field),
	},
	7: {
		0: uint16(alias_sym_field),
	},
	8: {
		0: uint16(alias_sym_operation_field),
	},
	9: {
		1: uint16(alias_sym_field),
	},
	10: {
		0: uint16(alias_sym_operation_error_field),
	},
}

var ts_non_terminal_alias_map = [21]uint16_t{
	0:  uint16(sym_node_value),
	1:  uint16(2),
	2:  uint16(sym_node_value),
	3:  uint16(alias_sym_trait_node_value),
	4:  uint16(sym_node_object_kvp),
	5:  uint16(2),
	6:  uint16(sym_node_object_kvp),
	7:  uint16(alias_sym_trait_object_kvp),
	8:  uint16(sym_string),
	9:  uint16(2),
	10: uint16(sym_string),
	11: uint16(sym__control_identifier),
	12: uint16(sym_identifier),
	13: uint16(6),
	14: uint16(sym_identifier),
	15: uint16(alias_sym_enum_field),
	16: uint16(alias_sym_field),
	17: uint16(alias_sym_key_identifier),
	18: uint16(alias_sym_operation_error_field),
	19: uint16(alias_sym_operation_field),
}

var ts_primary_state_ids = [350]TSStateId{
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
	17:  uint16(16),
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
	29:  uint16(28),
	30:  uint16(30),
	31:  uint16(31),
	32:  uint16(28),
	33:  uint16(33),
	34:  uint16(34),
	35:  uint16(34),
	36:  uint16(34),
	37:  uint16(37),
	38:  uint16(38),
	39:  uint16(39),
	40:  uint16(40),
	41:  uint16(41),
	42:  uint16(42),
	43:  uint16(43),
	44:  uint16(39),
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
	81:  uint16(9),
	82:  uint16(7),
	83:  uint16(5),
	84:  uint16(8),
	85:  uint16(6),
	86:  uint16(11),
	87:  uint16(12),
	88:  uint16(13),
	89:  uint16(14),
	90:  uint16(90),
	91:  uint16(91),
	92:  uint16(92),
	93:  uint16(93),
	94:  uint16(7),
	95:  uint16(95),
	96:  uint16(96),
	97:  uint16(97),
	98:  uint16(98),
	99:  uint16(19),
	100: uint16(5),
	101: uint16(101),
	102: uint16(102),
	103: uint16(103),
	104: uint16(104),
	105: uint16(18),
	106: uint16(106),
	107: uint16(107),
	108: uint16(108),
	109: uint16(8),
	110: uint16(110),
	111: uint16(110),
	112: uint16(112),
	113: uint16(113),
	114: uint16(114),
	115: uint16(110),
	116: uint16(116),
	117: uint16(117),
	118: uint16(113),
	119: uint16(113),
	120: uint16(9),
	121: uint16(12),
	122: uint16(6),
	123: uint16(13),
	124: uint16(124),
	125: uint16(125),
	126: uint16(126),
	127: uint16(14),
	128: uint16(124),
	129: uint16(125),
	130: uint16(11),
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
	142: uint16(139),
	143: uint16(9),
	144: uint16(139),
	145: uint16(145),
	146: uint16(138),
	147: uint16(138),
	148: uint16(148),
	149: uint16(8),
	150: uint16(19),
	151: uint16(151),
	152: uint16(152),
	153: uint16(153),
	154: uint16(97),
	155: uint16(155),
	156: uint16(156),
	157: uint16(157),
	158: uint16(158),
	159: uint16(159),
	160: uint16(160),
	161: uint16(106),
	162: uint16(98),
	163: uint16(163),
	164: uint16(164),
	165: uint16(107),
	166: uint16(157),
	167: uint16(167),
	168: uint16(91),
	169: uint16(91),
	170: uint16(103),
	171: uint16(18),
	172: uint16(92),
	173: uint16(106),
	174: uint16(90),
	175: uint16(107),
	176: uint16(176),
	177: uint16(93),
	178: uint16(108),
	179: uint16(95),
	180: uint16(102),
	181: uint16(157),
	182: uint16(101),
	183: uint16(183),
	184: uint16(96),
	185: uint16(97),
	186: uint16(98),
	187: uint16(187),
	188: uint16(188),
	189: uint16(189),
	190: uint16(190),
	191: uint16(9),
	192: uint16(192),
	193: uint16(193),
	194: uint16(193),
	195: uint16(189),
	196: uint16(189),
	197: uint16(197),
	198: uint16(193),
	199: uint16(8),
	200: uint16(200),
	201: uint16(201),
	202: uint16(202),
	203: uint16(203),
	204: uint16(204),
	205: uint16(205),
	206: uint16(206),
	207: uint16(207),
	208: uint16(92),
	209: uint16(90),
	210: uint16(210),
	211: uint16(211),
	212: uint16(212),
	213: uint16(213),
	214: uint16(214),
	215: uint16(15),
	216: uint16(108),
	217: uint16(217),
	218: uint16(102),
	219: uint16(101),
	220: uint16(220),
	221: uint16(221),
	222: uint16(222),
	223: uint16(26),
	224: uint16(224),
	225: uint16(225),
	226: uint16(96),
	227: uint16(95),
	228: uint16(103),
	229: uint16(93),
	230: uint16(230),
	231: uint16(231),
	232: uint16(232),
	233: uint16(233),
	234: uint16(234),
	235: uint16(235),
	236: uint16(217),
	237: uint16(211),
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
	255: uint16(26),
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
	269: uint16(20),
	270: uint16(270),
	271: uint16(22),
	272: uint16(272),
	273: uint16(21),
	274: uint16(37),
	275: uint16(275),
	276: uint16(276),
	277: uint16(277),
	278: uint16(278),
	279: uint16(279),
	280: uint16(280),
	281: uint16(281),
	282: uint16(282),
	283: uint16(275),
	284: uint16(284),
	285: uint16(285),
	286: uint16(286),
	287: uint16(285),
	288: uint16(284),
	289: uint16(289),
	290: uint16(290),
	291: uint16(291),
	292: uint16(45),
	293: uint16(73),
	294: uint16(294),
	295: uint16(295),
	296: uint16(296),
	297: uint16(297),
	298: uint16(298),
	299: uint16(299),
	300: uint16(275),
	301: uint16(301),
	302: uint16(302),
	303: uint16(303),
	304: uint16(304),
	305: uint16(305),
	306: uint16(306),
	307: uint16(307),
	308: uint16(284),
	309: uint16(285),
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
	332: uint16(328),
	333: uint16(333),
	334: uint16(334),
	335: uint16(335),
	336: uint16(336),
	337: uint16(337),
	338: uint16(336),
	339: uint16(339),
	340: uint16(340),
	341: uint16(341),
	342: uint16(342),
	343: uint16(343),
	344: uint16(344),
	345: uint16(328),
	346: uint16(346),
	347: uint16(347),
	348: uint16(348),
	349: uint16(349),
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
			state = uint16(152)
			goto next_state
		}
		i = uint32(0)
		for {
			if !(uint64(i) < libc.Uint64FromInt64(128)/libc.Uint64FromInt64(2)) {
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
			state = uint16(151)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(206)
			goto next_state
		}
		return result
	case int32(1):
		i1 = uint32(0)
		for {
			if !(uint64(i1) < libc.Uint64FromInt64(52)/libc.Uint64FromInt64(2)) {
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
			state = uint16(1)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(249)
			goto next_state
		}
		return result
	case int32(2):
		i2 = uint32(0)
		for {
			if !(uint64(i2) < libc.Uint64FromInt64(56)/libc.Uint64FromInt64(2)) {
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
			state = uint16(2)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(206)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(249)
			goto next_state
		}
		return result
	case int32(3):
		if lookahead == int32('"') {
			state = uint16(211)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead == int32('"') {
			state = uint16(5)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(176)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(221)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(219)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(220)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(228)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead == int32('"') {
			state = uint16(230)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(6)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead == int32('"') {
			state = uint16(229)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(6)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead == int32('"') {
			state = uint16(209)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(216)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(215)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(133)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(214)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(216)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead == int32('/') {
			state = uint16(250)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead == int32('D') {
			state = uint16(43)
			goto next_state
		}
		if lookahead == int32('I') {
			state = uint16(90)
			goto next_state
		}
		return result
	case int32(10):
		if lookahead == int32('E') {
			state = uint16(91)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(55)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead == int32('a') {
			state = uint16(70)
			goto next_state
		}
		if lookahead == int32('l') {
			state = uint16(94)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(106)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead == int32('a') {
			state = uint16(79)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(71)
			goto next_state
		}
		return result
	case int32(13):
		if lookahead == int32('a') {
			state = uint16(31)
			goto next_state
		}
		return result
	case int32(14):
		if lookahead == int32('a') {
			state = uint16(155)
			goto next_state
		}
		return result
	case int32(15):
		if lookahead == int32('a') {
			state = uint16(100)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(126)
			goto next_state
		}
		return result
	case int32(16):
		if lookahead == int32('a') {
			state = uint16(123)
			goto next_state
		}
		return result
	case int32(17):
		if lookahead == int32('a') {
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(18):
		if lookahead == int32('a') {
			state = uint16(76)
			goto next_state
		}
		return result
	case int32(19):
		if lookahead == int32('a') {
			state = uint16(83)
			goto next_state
		}
		return result
	case int32(20):
		if lookahead == int32('a') {
			state = uint16(131)
			goto next_state
		}
		return result
	case int32(21):
		if lookahead == int32('a') {
			state = uint16(129)
			goto next_state
		}
		return result
	case int32(22):
		if lookahead == int32('a') {
			state = uint16(29)
			goto next_state
		}
		return result
	case int32(23):
		if lookahead == int32('b') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(24):
		if lookahead == int32('b') {
			state = uint16(73)
			goto next_state
		}
		return result
	case int32(25):
		if lookahead == int32('c') {
			state = uint16(63)
			goto next_state
		}
		return result
	case int32(26):
		if lookahead == int32('c') {
			state = uint16(132)
			goto next_state
		}
		return result
	case int32(27):
		if lookahead == int32('c') {
			state = uint16(38)
			goto next_state
		}
		return result
	case int32(28):
		if lookahead == int32('c') {
			state = uint16(39)
			goto next_state
		}
		return result
	case int32(29):
		if lookahead == int32('c') {
			state = uint16(40)
			goto next_state
		}
		return result
	case int32(30):
		if lookahead == int32('c') {
			state = uint16(135)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(24)
			goto next_state
		}
		return result
	case int32(31):
		if lookahead == int32('d') {
			state = uint16(21)
			goto next_state
		}
		return result
	case int32(32):
		if lookahead == int32('e') {
			state = uint16(107)
			goto next_state
		}
		if lookahead == int32('h') {
			state = uint16(97)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(108)
			goto next_state
		}
		return result
	case int32(33):
		if lookahead == int32('e') {
			state = uint16(160)
			goto next_state
		}
		return result
	case int32(34):
		if lookahead == int32('e') {
			state = uint16(188)
			goto next_state
		}
		return result
	case int32(35):
		if lookahead == int32('e') {
			state = uint16(199)
			goto next_state
		}
		return result
	case int32(36):
		if lookahead == int32('e') {
			state = uint16(201)
			goto next_state
		}
		return result
	case int32(37):
		if lookahead == int32('e') {
			state = uint16(190)
			goto next_state
		}
		return result
	case int32(38):
		if lookahead == int32('e') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(39):
		if lookahead == int32('e') {
			state = uint16(173)
			goto next_state
		}
		return result
	case int32(40):
		if lookahead == int32('e') {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(41):
		if lookahead == int32('e') {
			state = uint16(169)
			goto next_state
		}
		return result
	case int32(42):
		if lookahead == int32('e') {
			state = uint16(117)
			goto next_state
		}
		return result
	case int32(43):
		if lookahead == int32('e') {
			state = uint16(25)
			goto next_state
		}
		return result
	case int32(44):
		if lookahead == int32('e') {
			state = uint16(113)
			goto next_state
		}
		return result
	case int32(45):
		if lookahead == int32('e') {
			state = uint16(19)
			goto next_state
		}
		return result
	case int32(46):
		if lookahead == int32('e') {
			state = uint16(115)
			goto next_state
		}
		return result
	case int32(47):
		if lookahead == int32('e') {
			state = uint16(109)
			goto next_state
		}
		return result
	case int32(48):
		if lookahead == int32('e') {
			state = uint16(110)
			goto next_state
		}
		return result
	case int32(49):
		if lookahead == int32('e') {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(50):
		if lookahead == int32('e') {
			state = uint16(89)
			goto next_state
		}
		return result
	case int32(51):
		if lookahead == int32('e') {
			state = uint16(56)
			goto next_state
		}
		return result
	case int32(52):
		if lookahead == int32('g') {
			state = uint16(9)
			goto next_state
		}
		return result
	case int32(53):
		if lookahead == int32('g') {
			state = uint16(193)
			goto next_state
		}
		return result
	case int32(54):
		if lookahead == int32('g') {
			state = uint16(195)
			goto next_state
		}
		return result
	case int32(55):
		if lookahead == int32('g') {
			state = uint16(47)
			goto next_state
		}
		return result
	case int32(56):
		if lookahead == int32('g') {
			state = uint16(48)
			goto next_state
		}
		return result
	case int32(57):
		if lookahead == int32('h') {
			state = uint16(184)
			goto next_state
		}
		return result
	case int32(58):
		if lookahead == int32('i') {
			state = uint16(52)
			goto next_state
		}
		if lookahead == int32('l') {
			state = uint16(93)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(96)
			goto next_state
		}
		if lookahead == int32('y') {
			state = uint16(127)
			goto next_state
		}
		return result
	case int32(59):
		if lookahead == int32('i') {
			state = uint16(116)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(86)
			goto next_state
		}
		return result
	case int32(60):
		if lookahead == int32('i') {
			state = uint16(27)
			goto next_state
		}
		return result
	case int32(61):
		if lookahead == int32('i') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(62):
		if lookahead == int32('i') {
			state = uint16(88)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(26)
			goto next_state
		}
		return result
	case int32(63):
		if lookahead == int32('i') {
			state = uint16(77)
			goto next_state
		}
		return result
	case int32(64):
		if lookahead == int32('i') {
			state = uint16(98)
			goto next_state
		}
		return result
	case int32(65):
		if lookahead == int32('i') {
			state = uint16(99)
			goto next_state
		}
		return result
	case int32(66):
		if lookahead == int32('i') {
			state = uint16(80)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(137)
			goto next_state
		}
		return result
	case int32(67):
		if lookahead == int32('l') {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(68):
		if lookahead == int32('l') {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(69):
		if lookahead == int32('l') {
			state = uint16(198)
			goto next_state
		}
		return result
	case int32(70):
		if lookahead == int32('l') {
			state = uint16(118)
			goto next_state
		}
		return result
	case int32(71):
		if lookahead == int32('l') {
			state = uint16(68)
			goto next_state
		}
		return result
	case int32(72):
		if lookahead == int32('l') {
			state = uint16(45)
			goto next_state
		}
		return result
	case int32(73):
		if lookahead == int32('l') {
			state = uint16(37)
			goto next_state
		}
		return result
	case int32(74):
		if lookahead == int32('m') {
			state = uint16(162)
			goto next_state
		}
		return result
	case int32(75):
		if lookahead == int32('m') {
			state = uint16(163)
			goto next_state
		}
		return result
	case int32(76):
		if lookahead == int32('m') {
			state = uint16(101)
			goto next_state
		}
		return result
	case int32(77):
		if lookahead == int32('m') {
			state = uint16(17)
			goto next_state
		}
		return result
	case int32(78):
		if lookahead == int32('m') {
			state = uint16(50)
			goto next_state
		}
		return result
	case int32(79):
		if lookahead == int32('m') {
			state = uint16(46)
			goto next_state
		}
		return result
	case int32(80):
		if lookahead == int32('m') {
			state = uint16(49)
			goto next_state
		}
		return result
	case int32(81):
		if lookahead == int32('n') {
			state = uint16(134)
			goto next_state
		}
		return result
	case int32(82):
		if lookahead == int32('n') {
			state = uint16(170)
			goto next_state
		}
		return result
	case int32(83):
		if lookahead == int32('n') {
			state = uint16(187)
			goto next_state
		}
		return result
	case int32(84):
		if lookahead == int32('n') {
			state = uint16(172)
			goto next_state
		}
		return result
	case int32(85):
		if lookahead == int32('n') {
			state = uint16(64)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(86):
		if lookahead == int32('n') {
			state = uint16(53)
			goto next_state
		}
		return result
	case int32(87):
		if lookahead == int32('n') {
			state = uint16(120)
			goto next_state
		}
		return result
	case int32(88):
		if lookahead == int32('n') {
			state = uint16(54)
			goto next_state
		}
		return result
	case int32(89):
		if lookahead == int32('n') {
			state = uint16(125)
			goto next_state
		}
		return result
	case int32(90):
		if lookahead == int32('n') {
			state = uint16(130)
			goto next_state
		}
		return result
	case int32(91):
		if lookahead == int32('n') {
			state = uint16(136)
			goto next_state
		}
		return result
	case int32(92):
		if lookahead == int32('o') {
			state = uint16(30)
			goto next_state
		}
		return result
	case int32(93):
		if lookahead == int32('o') {
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(94):
		if lookahead == int32('o') {
			state = uint16(16)
			goto next_state
		}
		return result
	case int32(95):
		if lookahead == int32('o') {
			state = uint16(138)
			goto next_state
		}
		return result
	case int32(96):
		if lookahead == int32('o') {
			state = uint16(72)
			goto next_state
		}
		return result
	case int32(97):
		if lookahead == int32('o') {
			state = uint16(111)
			goto next_state
		}
		return result
	case int32(98):
		if lookahead == int32('o') {
			state = uint16(82)
			goto next_state
		}
		return result
	case int32(99):
		if lookahead == int32('o') {
			state = uint16(84)
			goto next_state
		}
		return result
	case int32(100):
		if lookahead == int32('p') {
			state = uint16(167)
			goto next_state
		}
		return result
	case int32(101):
		if lookahead == int32('p') {
			state = uint16(196)
			goto next_state
		}
		return result
	case int32(102):
		if lookahead == int32('p') {
			state = uint16(44)
			goto next_state
		}
		return result
	case int32(103):
		if lookahead == int32('p') {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(104):
		if lookahead == int32('p') {
			state = uint16(67)
			goto next_state
		}
		return result
	case int32(105):
		if lookahead == int32('p') {
			state = uint16(22)
			goto next_state
		}
		return result
	case int32(106):
		if lookahead == int32('r') {
			state = uint16(185)
			goto next_state
		}
		return result
	case int32(107):
		if lookahead == int32('r') {
			state = uint16(140)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(168)
			goto next_state
		}
		return result
	case int32(108):
		if lookahead == int32('r') {
			state = uint16(62)
			goto next_state
		}
		return result
	case int32(109):
		if lookahead == int32('r') {
			state = uint16(192)
			goto next_state
		}
		return result
	case int32(110):
		if lookahead == int32('r') {
			state = uint16(197)
			goto next_state
		}
		return result
	case int32(111):
		if lookahead == int32('r') {
			state = uint16(124)
			goto next_state
		}
		return result
	case int32(112):
		if lookahead == int32('r') {
			state = uint16(41)
			goto next_state
		}
		return result
	case int32(113):
		if lookahead == int32('r') {
			state = uint16(20)
			goto next_state
		}
		return result
	case int32(114):
		if lookahead == int32('r') {
			state = uint16(28)
			goto next_state
		}
		return result
	case int32(115):
		if lookahead == int32('s') {
			state = uint16(105)
			goto next_state
		}
		return result
	case int32(116):
		if lookahead == int32('s') {
			state = uint16(122)
			goto next_state
		}
		return result
	case int32(117):
		if lookahead == int32('s') {
			state = uint16(95)
			goto next_state
		}
		return result
	case int32(118):
		if lookahead == int32('s') {
			state = uint16(36)
			goto next_state
		}
		return result
	case int32(119):
		if lookahead == int32('s') {
			state = uint16(128)
			goto next_state
		}
		return result
	case int32(120):
		if lookahead == int32('t') {
			state = uint16(10)
			goto next_state
		}
		return result
	case int32(121):
		if lookahead == int32('t') {
			state = uint16(57)
			goto next_state
		}
		return result
	case int32(122):
		if lookahead == int32('t') {
			state = uint16(166)
			goto next_state
		}
		return result
	case int32(123):
		if lookahead == int32('t') {
			state = uint16(191)
			goto next_state
		}
		return result
	case int32(124):
		if lookahead == int32('t') {
			state = uint16(194)
			goto next_state
		}
		return result
	case int32(125):
		if lookahead == int32('t') {
			state = uint16(189)
			goto next_state
		}
		return result
	case int32(126):
		if lookahead == int32('t') {
			state = uint16(13)
			goto next_state
		}
		return result
	case int32(127):
		if lookahead == int32('t') {
			state = uint16(34)
			goto next_state
		}
		return result
	case int32(128):
		if lookahead == int32('t') {
			state = uint16(18)
			goto next_state
		}
		return result
	case int32(129):
		if lookahead == int32('t') {
			state = uint16(14)
			goto next_state
		}
		return result
	case int32(130):
		if lookahead == int32('t') {
			state = uint16(51)
			goto next_state
		}
		return result
	case int32(131):
		if lookahead == int32('t') {
			state = uint16(65)
			goto next_state
		}
		return result
	case int32(132):
		if lookahead == int32('t') {
			state = uint16(139)
			goto next_state
		}
		return result
	case int32(133):
		if lookahead == int32('u') {
			state = uint16(142)
			goto next_state
		}
		if lookahead == int32('x') {
			state = uint16(150)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('7') {
			state = uint16(232)
			goto next_state
		}
		if lookahead == int32('"') || lookahead == int32('\'') || lookahead == int32('?') || lookahead == int32('\\') || lookahead == int32('a') || lookahead == int32('b') || lookahead == int32('f') || lookahead == int32('n') || lookahead == int32('r') || int32('t') <= lookahead && lookahead <= int32('v') {
			state = uint16(235)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(231)
			goto next_state
		}
		return result
	case int32(134):
		if lookahead == int32('u') {
			state = uint16(74)
			goto next_state
		}
		return result
	case int32(135):
		if lookahead == int32('u') {
			state = uint16(78)
			goto next_state
		}
		return result
	case int32(136):
		if lookahead == int32('u') {
			state = uint16(75)
			goto next_state
		}
		return result
	case int32(137):
		if lookahead == int32('u') {
			state = uint16(35)
			goto next_state
		}
		return result
	case int32(138):
		if lookahead == int32('u') {
			state = uint16(114)
			goto next_state
		}
		return result
	case int32(139):
		if lookahead == int32('u') {
			state = uint16(112)
			goto next_state
		}
		return result
	case int32(140):
		if lookahead == int32('v') {
			state = uint16(60)
			goto next_state
		}
		return result
	case int32(141):
		if lookahead == int32('y') {
			state = uint16(183)
			goto next_state
		}
		return result
	case int32(142):
		if lookahead == int32('{') {
			state = uint16(149)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(147)
			goto next_state
		}
		return result
	case int32(143):
		if lookahead == int32('}') {
			state = uint16(235)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(143)
			goto next_state
		}
		return result
	case int32(144):
		if lookahead == int32('+') || lookahead == int32('-') {
			state = uint16(146)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(208)
			goto next_state
		}
		return result
	case int32(145):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(207)
			goto next_state
		}
		return result
	case int32(146):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(208)
			goto next_state
		}
		return result
	case int32(147):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(150)
			goto next_state
		}
		return result
	case int32(148):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(235)
			goto next_state
		}
		return result
	case int32(149):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(143)
			goto next_state
		}
		return result
	case int32(150):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(148)
			goto next_state
		}
		return result
	case int32(151):
		if eof != 0 {
			state = uint16(152)
			goto next_state
		}
		i3 = uint32(0)
		for {
			if !(uint64(i3) < libc.Uint64FromInt64(124)/libc.Uint64FromInt64(2)) {
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
			state = uint16(151)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(206)
			goto next_state
		}
		return result
	case int32(152):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(ts_builtin_sym_end)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(153):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOLLAR)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(154):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COLON)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('=') {
			state = uint16(179)
			goto next_state
		}
		return result
	case int32(155):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_metadata)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(156):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(157):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_namespace)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(158):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(159):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(207)
			goto next_state
		}
		return result
	case int32(160):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_use)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(161):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUND)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(162):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_enum)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(163):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_intEnum)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(164):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(165):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(166):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_list)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(167):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_map)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(168):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_set)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(169):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_structure)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(170):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_union)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(171):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_service)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(172):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_operation)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(173):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_resource)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(174):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COMMA)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(175):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COMMA)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(256)
			goto next_state
		}
		return result
	case int32(176):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COMMA)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('"') {
			state = uint16(228)
			goto next_state
		}
		return result
	case int32(177):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(178):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(179):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COLON_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(180):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(181):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(182):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(183):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_apply)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(184):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_with)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(185):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_for)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(186):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_blob)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(187):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_boolean)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(188):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_byte)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(189):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_document)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(190):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_double)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(191):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_float)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(192):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_integer)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(193):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_long)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(194):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_short)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(195):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(196):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_timestamp)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(197):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_bigInteger)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(198):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_bigDecimal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(199):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_true)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(200):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_true)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(249)
			goto next_state
		}
		return result
	case int32(201):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_false)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(202):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_false)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(249)
			goto next_state
		}
		return result
	case int32(203):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_null)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(204):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_null)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(249)
			goto next_state
		}
		return result
	case int32(205):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(206):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_number_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(145)
			goto next_state
		}
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(144)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(206)
			goto next_state
		}
		return result
	case int32(207):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_float_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(144)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(207)
			goto next_state
		}
		return result
	case int32(208):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_float_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(208)
			goto next_state
		}
		return result
	case int32(209):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(210):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('"') {
			state = uint16(3)
			goto next_state
		}
		return result
	case int32(211):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(212):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string_fragment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(216)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(213)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('"') && lookahead != int32('\\') {
			state = uint16(213)
			goto next_state
		}
		return result
	case int32(213):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string_fragment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(216)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('"') && lookahead != int32('\\') {
			state = uint16(213)
			goto next_state
		}
		return result
	case int32(214):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string_fragment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(',') {
			state = uint16(216)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(215)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(214)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('"') && lookahead != int32('\\') {
			state = uint16(216)
			goto next_state
		}
		return result
	case int32(215):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string_fragment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('/') {
			state = uint16(212)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('"') && lookahead != int32('\\') {
			state = uint16(216)
			goto next_state
		}
		return result
	case int32(216):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string_fragment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('"') && lookahead != int32('\\') {
			state = uint16(216)
			goto next_state
		}
		return result
	case int32(217):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__multiline_string_fragment_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(228)
			goto next_state
		}
		if lookahead == int32('"') {
			state = uint16(251)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(252)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(218)
			goto next_state
		}
		return result
	case int32(218):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__multiline_string_fragment_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(228)
			goto next_state
		}
		if lookahead == int32('"') {
			state = uint16(251)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(218)
			goto next_state
		}
		return result
	case int32(219):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__multiline_string_fragment_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('"') {
			state = uint16(235)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(222)
			goto next_state
		}
		if lookahead == int32('x') {
			state = uint16(227)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('7') {
			state = uint16(233)
			goto next_state
		}
		if lookahead == int32('\'') || lookahead == int32('?') || lookahead == int32('\\') || lookahead == int32('a') || lookahead == int32('b') || lookahead == int32('f') || lookahead == int32('n') || lookahead == int32('r') || int32('t') <= lookahead && lookahead <= int32('v') {
			state = uint16(238)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(234)
			goto next_state
		}
		return result
	case int32(220):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__multiline_string_fragment_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(',') {
			state = uint16(176)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(221)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(220)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('"') {
			state = uint16(228)
			goto next_state
		}
		return result
	case int32(221):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__multiline_string_fragment_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('/') {
			state = uint16(217)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('"') {
			state = uint16(228)
			goto next_state
		}
		return result
	case int32(222):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__multiline_string_fragment_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('{') {
			state = uint16(226)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(224)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('"') {
			state = uint16(228)
			goto next_state
		}
		return result
	case int32(223):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__multiline_string_fragment_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('}') {
			state = uint16(238)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(223)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('"') {
			state = uint16(228)
			goto next_state
		}
		return result
	case int32(224):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__multiline_string_fragment_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(227)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('"') {
			state = uint16(228)
			goto next_state
		}
		return result
	case int32(225):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__multiline_string_fragment_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(238)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('"') {
			state = uint16(228)
			goto next_state
		}
		return result
	case int32(226):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__multiline_string_fragment_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(223)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('"') {
			state = uint16(228)
			goto next_state
		}
		return result
	case int32(227):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__multiline_string_fragment_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(225)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('"') {
			state = uint16(228)
			goto next_state
		}
		return result
	case int32(228):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__multiline_string_fragment_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('"') {
			state = uint16(228)
			goto next_state
		}
		return result
	case int32(229):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__multiline_string_fragment_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(230):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__multiline_string_fragment_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('"') {
			state = uint16(211)
			goto next_state
		}
		return result
	case int32(231):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__escape_sequence_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(232):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__escape_sequence_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('7') {
			state = uint16(236)
			goto next_state
		}
		return result
	case int32(233):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__escape_sequence_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('7') {
			state = uint16(237)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('"') {
			state = uint16(228)
			goto next_state
		}
		return result
	case int32(234):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__escape_sequence_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('"') {
			state = uint16(228)
			goto next_state
		}
		return result
	case int32(235):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_escape_sequence)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(236):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_escape_sequence)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('7') {
			state = uint16(235)
			goto next_state
		}
		return result
	case int32(237):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_escape_sequence)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('7') {
			state = uint16(238)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('"') {
			state = uint16(228)
			goto next_state
		}
		return result
	case int32(238):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_escape_sequence)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('"') {
			state = uint16(228)
			goto next_state
		}
		return result
	case int32(239):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_identifier_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('a') {
			state = uint16(242)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('b') <= lookahead && lookahead <= int32('z') {
			state = uint16(249)
			goto next_state
		}
		return result
	case int32(240):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_identifier_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(200)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(249)
			goto next_state
		}
		return result
	case int32(241):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_identifier_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(202)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(249)
			goto next_state
		}
		return result
	case int32(242):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_identifier_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('l') {
			state = uint16(246)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(249)
			goto next_state
		}
		return result
	case int32(243):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_identifier_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('l') {
			state = uint16(204)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(249)
			goto next_state
		}
		return result
	case int32(244):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_identifier_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('l') {
			state = uint16(243)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(249)
			goto next_state
		}
		return result
	case int32(245):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_identifier_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(247)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(249)
			goto next_state
		}
		return result
	case int32(246):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_identifier_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('s') {
			state = uint16(241)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(249)
			goto next_state
		}
		return result
	case int32(247):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_identifier_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('u') {
			state = uint16(240)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(249)
			goto next_state
		}
		return result
	case int32(248):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_identifier_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('u') {
			state = uint16(244)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(249)
			goto next_state
		}
		return result
	case int32(249):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_identifier_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(249)
			goto next_state
		}
		return result
	case int32(250):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('/') {
			state = uint16(253)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(251)
			goto next_state
		}
		return result
	case int32(251):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(251)
			goto next_state
		}
		return result
	case int32(252):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SLASH_SLASH_SLASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(228)
			goto next_state
		}
		if lookahead == int32('"') {
			state = uint16(251)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(218)
			goto next_state
		}
		return result
	case int32(253):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SLASH_SLASH_SLASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(251)
			goto next_state
		}
		return result
	case int32(254):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_documentation_comment_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(',') {
			state = uint16(175)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(255)
			goto next_state
		}
		if lookahead == int32('\t') || int32(0x0b) <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(254)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) {
			state = uint16(256)
			goto next_state
		}
		return result
	case int32(255):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_documentation_comment_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('/') {
			state = uint16(250)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(256)
			goto next_state
		}
		return result
	case int32(256):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_documentation_comment_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(256)
			goto next_state
		}
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var map_token = [64]uint16_t{
	0:  uint16('"'),
	1:  uint16(210),
	2:  uint16('#'),
	3:  uint16(161),
	4:  uint16('$'),
	5:  uint16(153),
	6:  uint16('('),
	7:  uint16(181),
	8:  uint16(')'),
	9:  uint16(182),
	10: uint16(','),
	11: uint16(174),
	12: uint16('-'),
	13: uint16(205),
	14: uint16('.'),
	15: uint16(159),
	16: uint16('/'),
	17: uint16(8),
	18: uint16(':'),
	19: uint16(154),
	20: uint16('='),
	21: uint16(156),
	22: uint16('@'),
	23: uint16(180),
	24: uint16('['),
	25: uint16(177),
	26: uint16('\\'),
	27: uint16(133),
	28: uint16(']'),
	29: uint16(178),
	30: uint16('a'),
	31: uint16(103),
	32: uint16('b'),
	33: uint16(58),
	34: uint16('d'),
	35: uint16(92),
	36: uint16('e'),
	37: uint16(81),
	38: uint16('f'),
	39: uint16(11),
	40: uint16('i'),
	41: uint16(87),
	42: uint16('l'),
	43: uint16(59),
	44: uint16('m'),
	45: uint16(15),
	46: uint16('n'),
	47: uint16(12),
	48: uint16('o'),
	49: uint16(102),
	50: uint16('r'),
	51: uint16(42),
	52: uint16('s'),
	53: uint16(32),
	54: uint16('t'),
	55: uint16(66),
	56: uint16('u'),
	57: uint16(85),
	58: uint16('w'),
	59: uint16(61),
	60: uint16('{'),
	61: uint16(164),
	62: uint16('}'),
	63: uint16(165),
}

var map_token1 = [26]uint16_t{
	0:  uint16('"'),
	1:  uint16(210),
	2:  uint16('#'),
	3:  uint16(161),
	4:  uint16('$'),
	5:  uint16(153),
	6:  uint16('('),
	7:  uint16(181),
	8:  uint16(')'),
	9:  uint16(182),
	10: uint16(','),
	11: uint16(174),
	12: uint16('.'),
	13: uint16(158),
	14: uint16('/'),
	15: uint16(8),
	16: uint16('='),
	17: uint16(156),
	18: uint16('@'),
	19: uint16(180),
	20: uint16('['),
	21: uint16(177),
	22: uint16(']'),
	23: uint16(178),
	24: uint16('}'),
	25: uint16(165),
}

var map_token2 = [28]uint16_t{
	0:  uint16('"'),
	1:  uint16(210),
	2:  uint16('#'),
	3:  uint16(161),
	4:  uint16('$'),
	5:  uint16(153),
	6:  uint16(')'),
	7:  uint16(182),
	8:  uint16(','),
	9:  uint16(174),
	10: uint16('-'),
	11: uint16(205),
	12: uint16('.'),
	13: uint16(159),
	14: uint16('/'),
	15: uint16(8),
	16: uint16('['),
	17: uint16(177),
	18: uint16(']'),
	19: uint16(178),
	20: uint16('f'),
	21: uint16(239),
	22: uint16('n'),
	23: uint16(248),
	24: uint16('t'),
	25: uint16(245),
	26: uint16('{'),
	27: uint16(164),
}

var map_token3 = [62]uint16_t{
	0:  uint16('"'),
	1:  uint16(210),
	2:  uint16('#'),
	3:  uint16(161),
	4:  uint16('$'),
	5:  uint16(153),
	6:  uint16('('),
	7:  uint16(181),
	8:  uint16(')'),
	9:  uint16(182),
	10: uint16(','),
	11: uint16(174),
	12: uint16('-'),
	13: uint16(205),
	14: uint16('.'),
	15: uint16(159),
	16: uint16('/'),
	17: uint16(8),
	18: uint16(':'),
	19: uint16(154),
	20: uint16('='),
	21: uint16(156),
	22: uint16('@'),
	23: uint16(180),
	24: uint16('['),
	25: uint16(177),
	26: uint16(']'),
	27: uint16(178),
	28: uint16('a'),
	29: uint16(103),
	30: uint16('b'),
	31: uint16(58),
	32: uint16('d'),
	33: uint16(92),
	34: uint16('e'),
	35: uint16(81),
	36: uint16('f'),
	37: uint16(11),
	38: uint16('i'),
	39: uint16(87),
	40: uint16('l'),
	41: uint16(59),
	42: uint16('m'),
	43: uint16(15),
	44: uint16('n'),
	45: uint16(12),
	46: uint16('o'),
	47: uint16(102),
	48: uint16('r'),
	49: uint16(42),
	50: uint16('s'),
	51: uint16(32),
	52: uint16('t'),
	53: uint16(66),
	54: uint16('u'),
	55: uint16(85),
	56: uint16('w'),
	57: uint16(61),
	58: uint16('{'),
	59: uint16(164),
	60: uint16('}'),
	61: uint16(165),
}

var ts_lex_modes = [350]TSLexMode{
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
	16: {
		Flex_state: uint16(2),
	},
	17: {
		Flex_state: uint16(2),
	},
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
	28: {
		Flex_state: uint16(2),
	},
	29: {
		Flex_state: uint16(2),
	},
	30: {},
	31: {
		Flex_state: uint16(2),
	},
	32: {
		Flex_state: uint16(2),
	},
	33: {},
	34: {
		Flex_state: uint16(2),
	},
	35: {
		Flex_state: uint16(2),
	},
	36: {
		Flex_state: uint16(2),
	},
	37: {},
	38: {},
	39: {
		Flex_state: uint16(2),
	},
	40: {},
	41: {},
	42: {},
	43: {},
	44: {
		Flex_state: uint16(2),
	},
	45: {},
	46: {},
	47: {},
	48: {},
	49: {},
	50: {},
	51: {},
	52: {},
	53: {},
	54: {},
	55: {},
	56: {},
	57: {},
	58: {},
	59: {},
	60: {},
	61: {},
	62: {},
	63: {},
	64: {},
	65: {
		Flex_state: uint16(2),
	},
	66: {},
	67: {},
	68: {},
	69: {},
	70: {},
	71: {},
	72: {},
	73: {},
	74: {},
	75: {},
	76: {},
	77: {
		Flex_state: uint16(2),
	},
	78: {},
	79: {
		Flex_state: uint16(2),
	},
	80: {},
	81: {
		Flex_state: uint16(2),
	},
	82: {
		Flex_state: uint16(2),
	},
	83: {
		Flex_state: uint16(2),
	},
	84: {
		Flex_state: uint16(2),
	},
	85: {
		Flex_state: uint16(2),
	},
	86: {
		Flex_state: uint16(2),
	},
	87: {
		Flex_state: uint16(2),
	},
	88: {
		Flex_state: uint16(2),
	},
	89: {
		Flex_state: uint16(2),
	},
	90: {
		Flex_state: uint16(2),
	},
	91: {
		Flex_state: uint16(2),
	},
	92: {
		Flex_state: uint16(2),
	},
	93: {
		Flex_state: uint16(2),
	},
	94: {
		Flex_state: uint16(1),
	},
	95: {
		Flex_state: uint16(2),
	},
	96: {
		Flex_state: uint16(2),
	},
	97: {
		Flex_state: uint16(2),
	},
	98: {
		Flex_state: uint16(2),
	},
	99: {
		Flex_state: uint16(2),
	},
	100: {
		Flex_state: uint16(1),
	},
	101: {
		Flex_state: uint16(2),
	},
	102: {
		Flex_state: uint16(2),
	},
	103: {
		Flex_state: uint16(2),
	},
	104: {
		Flex_state: uint16(2),
	},
	105: {
		Flex_state: uint16(2),
	},
	106: {
		Flex_state: uint16(2),
	},
	107: {
		Flex_state: uint16(2),
	},
	108: {
		Flex_state: uint16(2),
	},
	109: {
		Flex_state: uint16(1),
	},
	110: {
		Flex_state: uint16(1),
	},
	111: {
		Flex_state: uint16(1),
	},
	112: {
		Flex_state: uint16(1),
	},
	113: {
		Flex_state: uint16(1),
	},
	114: {
		Flex_state: uint16(2),
	},
	115: {
		Flex_state: uint16(1),
	},
	116: {
		Flex_state: uint16(1),
	},
	117: {
		Flex_state: uint16(1),
	},
	118: {
		Flex_state: uint16(1),
	},
	119: {
		Flex_state: uint16(1),
	},
	120: {
		Flex_state: uint16(1),
	},
	121: {
		Flex_state: uint16(1),
	},
	122: {
		Flex_state: uint16(1),
	},
	123: {
		Flex_state: uint16(1),
	},
	124: {
		Flex_state: uint16(1),
	},
	125: {
		Flex_state: uint16(1),
	},
	126: {
		Flex_state: uint16(1),
	},
	127: {
		Flex_state: uint16(1),
	},
	128: {
		Flex_state: uint16(1),
	},
	129: {
		Flex_state: uint16(1),
	},
	130: {
		Flex_state: uint16(1),
	},
	131: {
		Flex_state: uint16(1),
	},
	132: {},
	133: {
		Flex_state: uint16(1),
	},
	134: {},
	135: {
		Flex_state: uint16(1),
	},
	136: {},
	137: {
		Flex_state: uint16(4),
	},
	138: {
		Flex_state: uint16(4),
	},
	139: {
		Flex_state: uint16(4),
	},
	140: {
		Flex_state: uint16(1),
	},
	141: {
		Flex_state: uint16(1),
	},
	142: {
		Flex_state: uint16(4),
	},
	143: {
		Flex_state: uint16(1),
	},
	144: {
		Flex_state: uint16(4),
	},
	145: {
		Flex_state: uint16(1),
	},
	146: {
		Flex_state: uint16(4),
	},
	147: {
		Flex_state: uint16(4),
	},
	148: {
		Flex_state: uint16(1),
	},
	149: {
		Flex_state: uint16(1),
	},
	150: {
		Flex_state: uint16(1),
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
	154: {},
	155: {},
	156: {},
	157: {
		Flex_state: uint16(1),
	},
	158: {
		Flex_state: uint16(1),
	},
	159: {
		Flex_state: uint16(1),
	},
	160: {
		Flex_state: uint16(1),
	},
	161: {},
	162: {},
	163: {
		Flex_state: uint16(1),
	},
	164: {
		Flex_state: uint16(1),
	},
	165: {},
	166: {
		Flex_state: uint16(1),
	},
	167: {
		Flex_state: uint16(1),
	},
	168: {},
	169: {
		Flex_state: uint16(1),
	},
	170: {
		Flex_state: uint16(1),
	},
	171: {
		Flex_state: uint16(1),
	},
	172: {
		Flex_state: uint16(1),
	},
	173: {
		Flex_state: uint16(1),
	},
	174: {
		Flex_state: uint16(1),
	},
	175: {
		Flex_state: uint16(1),
	},
	176: {
		Flex_state: uint16(1),
	},
	177: {
		Flex_state: uint16(1),
	},
	178: {
		Flex_state: uint16(1),
	},
	179: {
		Flex_state: uint16(1),
	},
	180: {
		Flex_state: uint16(1),
	},
	181: {
		Flex_state: uint16(1),
	},
	182: {
		Flex_state: uint16(1),
	},
	183: {
		Flex_state: uint16(1),
	},
	184: {
		Flex_state: uint16(1),
	},
	185: {
		Flex_state: uint16(1),
	},
	186: {
		Flex_state: uint16(1),
	},
	187: {
		Flex_state: uint16(4),
	},
	188: {
		Flex_state: uint16(1),
	},
	189: {
		Flex_state: uint16(7),
	},
	190: {
		Flex_state: uint16(1),
	},
	191: {},
	192: {
		Flex_state: uint16(1),
	},
	193: {
		Flex_state: uint16(7),
	},
	194: {
		Flex_state: uint16(7),
	},
	195: {
		Flex_state: uint16(7),
	},
	196: {
		Flex_state: uint16(7),
	},
	197: {
		Flex_state: uint16(4),
	},
	198: {
		Flex_state: uint16(7),
	},
	199: {},
	200: {},
	201: {
		Flex_state: uint16(7),
	},
	202: {
		Flex_state: uint16(1),
	},
	203: {
		Flex_state: uint16(4),
	},
	204: {
		Flex_state: uint16(1),
	},
	205: {},
	206: {
		Flex_state: uint16(4),
	},
	207: {
		Flex_state: uint16(1),
	},
	208: {},
	209: {},
	210: {
		Flex_state: uint16(1),
	},
	211: {
		Flex_state: uint16(4),
	},
	212: {
		Flex_state: uint16(4),
	},
	213: {
		Flex_state: uint16(1),
	},
	214: {
		Flex_state: uint16(4),
	},
	215: {
		Flex_state: uint16(1),
	},
	216: {},
	217: {
		Flex_state: uint16(4),
	},
	218: {},
	219: {},
	220: {
		Flex_state: uint16(1),
	},
	221: {
		Flex_state: uint16(1),
	},
	222: {
		Flex_state: uint16(1),
	},
	223: {
		Flex_state: uint16(1),
	},
	224: {},
	225: {
		Flex_state: uint16(1),
	},
	226: {},
	227: {},
	228: {},
	229: {},
	230: {
		Flex_state: uint16(1),
	},
	231: {},
	232: {
		Flex_state: uint16(1),
	},
	233: {},
	234: {
		Flex_state: uint16(1),
	},
	235: {
		Flex_state: uint16(7),
	},
	236: {
		Flex_state: uint16(7),
	},
	237: {
		Flex_state: uint16(7),
	},
	238: {},
	239: {},
	240: {
		Flex_state: uint16(1),
	},
	241: {},
	242: {
		Flex_state: uint16(1),
	},
	243: {},
	244: {
		Flex_state: uint16(1),
	},
	245: {},
	246: {
		Flex_state: uint16(1),
	},
	247: {},
	248: {
		Flex_state: uint16(1),
	},
	249: {},
	250: {
		Flex_state: uint16(1),
	},
	251: {},
	252: {
		Flex_state: uint16(1),
	},
	253: {
		Flex_state: uint16(1),
	},
	254: {},
	255: {
		Flex_state: uint16(1),
	},
	256: {},
	257: {},
	258: {},
	259: {},
	260: {},
	261: {},
	262: {
		Flex_state: uint16(1),
	},
	263: {
		Flex_state: uint16(1),
	},
	264: {
		Flex_state: uint16(1),
	},
	265: {},
	266: {},
	267: {},
	268: {
		Flex_state: uint16(1),
	},
	269: {
		Flex_state: uint16(1),
	},
	270: {},
	271: {
		Flex_state: uint16(1),
	},
	272: {},
	273: {
		Flex_state: uint16(1),
	},
	274: {
		Flex_state: uint16(1),
	},
	275: {},
	276: {},
	277: {
		Flex_state: uint16(1),
	},
	278: {
		Flex_state: uint16(1),
	},
	279: {},
	280: {},
	281: {},
	282: {
		Flex_state: uint16(1),
	},
	283: {},
	284: {
		Flex_state: uint16(1),
	},
	285: {
		Flex_state: uint16(1),
	},
	286: {
		Flex_state: uint16(1),
	},
	287: {
		Flex_state: uint16(1),
	},
	288: {
		Flex_state: uint16(1),
	},
	289: {},
	290: {},
	291: {
		Flex_state: uint16(1),
	},
	292: {
		Flex_state: uint16(1),
	},
	293: {
		Flex_state: uint16(1),
	},
	294: {
		Flex_state: uint16(1),
	},
	295: {},
	296: {},
	297: {
		Flex_state: uint16(1),
	},
	298: {
		Flex_state: uint16(1),
	},
	299: {
		Flex_state: uint16(1),
	},
	300: {},
	301: {
		Flex_state: uint16(1),
	},
	302: {},
	303: {
		Flex_state: uint16(1),
	},
	304: {
		Flex_state: uint16(1),
	},
	305: {
		Flex_state: uint16(1),
	},
	306: {},
	307: {},
	308: {
		Flex_state: uint16(1),
	},
	309: {
		Flex_state: uint16(1),
	},
	310: {
		Flex_state: uint16(1),
	},
	311: {},
	312: {
		Flex_state: uint16(1),
	},
	313: {},
	314: {
		Flex_state: uint16(1),
	},
	315: {
		Flex_state: uint16(1),
	},
	316: {},
	317: {
		Flex_state: uint16(1),
	},
	318: {
		Flex_state: uint16(1),
	},
	319: {
		Flex_state: uint16(1),
	},
	320: {
		Flex_state: uint16(1),
	},
	321: {
		Flex_state: uint16(1),
	},
	322: {
		Flex_state: uint16(1),
	},
	323: {
		Flex_state: uint16(1),
	},
	324: {},
	325: {
		Flex_state: uint16(1),
	},
	326: {},
	327: {},
	328: {},
	329: {},
	330: {},
	331: {},
	332: {},
	333: {
		Flex_state: uint16(1),
	},
	334: {},
	335: {},
	336: {},
	337: {
		Flex_state: uint16(254),
	},
	338: {},
	339: {},
	340: {},
	341: {},
	342: {},
	343: {},
	344: {},
	345: {},
	346: {},
	347: {},
	348: {},
	349: {
		Flex_state: uint16(-libc.Int32FromInt32(1)),
	},
}

var ts_parse_table = [2][145]uint16_t{
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
		21: uint16(3),
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
		55: uint16(1),
		56: uint16(1),
		58: uint16(5),
		59: uint16(7),
	},
	1: {
		0:   uint16(9),
		1:   uint16(11),
		3:   uint16(13),
		5:   uint16(15),
		21:  uint16(3),
		58:  uint16(5),
		59:  uint16(7),
		61:  uint16(334),
		62:  uint16(136),
		63:  uint16(239),
		64:  uint16(79),
		65:  uint16(245),
		66:  uint16(270),
		67:  uint16(347),
		68:  uint16(2),
		126: uint16(1),
		127: uint16(155),
		128: uint16(224),
	},
}

var ts_small_parse_table = [10927]uint16_t{
	0:     uint16(27),
	1:     uint16(3),
	2:     uint16(1),
	3:     uint16(anon_sym_COMMA),
	4:     uint16(5),
	5:     uint16(1),
	6:     uint16(sym_comment),
	7:     uint16(7),
	8:     uint16(1),
	9:     uint16(anon_sym_SLASH_SLASH_SLASH),
	10:    uint16(17),
	11:    uint16(1),
	13:    uint16(19),
	14:    uint16(1),
	15:    uint16(anon_sym_use),
	16:    uint16(23),
	17:    uint16(1),
	18:    uint16(anon_sym_list),
	19:    uint16(25),
	20:    uint16(1),
	21:    uint16(anon_sym_map),
	22:    uint16(27),
	23:    uint16(1),
	24:    uint16(anon_sym_set),
	25:    uint16(29),
	26:    uint16(1),
	27:    uint16(anon_sym_structure),
	28:    uint16(31),
	29:    uint16(1),
	30:    uint16(anon_sym_union),
	31:    uint16(33),
	32:    uint16(1),
	33:    uint16(anon_sym_service),
	34:    uint16(35),
	35:    uint16(1),
	36:    uint16(anon_sym_operation),
	37:    uint16(37),
	38:    uint16(1),
	39:    uint16(anon_sym_resource),
	40:    uint16(39),
	41:    uint16(1),
	42:    uint16(anon_sym_AT),
	43:    uint16(41),
	44:    uint16(1),
	45:    uint16(anon_sym_apply),
	46:    uint16(2),
	47:    uint16(1),
	48:    uint16(sym_documentation_comment),
	49:    uint16(3),
	50:    uint16(1),
	51:    uint16(aux_sym_shape_section_repeat1),
	52:    uint16(10),
	53:    uint16(1),
	54:    uint16(aux_sym_shape_statement_repeat1),
	55:    uint16(37),
	56:    uint16(1),
	57:    uint16(sym_trait_statement),
	58:    uint16(55),
	59:    uint16(1),
	60:    uint16(sym__definition),
	61:    uint16(57),
	62:    uint16(1),
	63:    uint16(sym_shape_body),
	64:    uint16(277),
	65:    uint16(1),
	66:    uint16(sym_primitive),
	67:    uint16(21),
	68:    uint16(2),
	69:    uint16(anon_sym_enum),
	70:    uint16(anon_sym_intEnum),
	71:    uint16(40),
	72:    uint16(2),
	73:    uint16(sym_apply_statement_singular),
	74:    uint16(sym_apply_statement_block),
	75:    uint16(56),
	76:    uint16(3),
	77:    uint16(sym_use_statement),
	78:    uint16(sym_shape_statement),
	79:    uint16(sym_apply_statement),
	80:    uint16(58),
	81:    uint16(10),
	82:    uint16(sym_simple_shape_statement),
	83:    uint16(sym_enum_statement),
	84:    uint16(sym_list_statement),
	85:    uint16(sym_map_statement),
	86:    uint16(sym_set_statement),
	87:    uint16(sym_structure_statement),
	88:    uint16(sym_union_statement),
	89:    uint16(sym_service_statement),
	90:    uint16(sym_operation_statement),
	91:    uint16(sym_resource_statement),
	92:    uint16(43),
	93:    uint16(13),
	94:    uint16(anon_sym_blob),
	95:    uint16(anon_sym_boolean),
	96:    uint16(anon_sym_byte),
	97:    uint16(anon_sym_document),
	98:    uint16(anon_sym_double),
	99:    uint16(anon_sym_float),
	100:   uint16(anon_sym_integer),
	101:   uint16(anon_sym_long),
	102:   uint16(anon_sym_short),
	103:   uint16(anon_sym_string),
	104:   uint16(anon_sym_timestamp),
	105:   uint16(anon_sym_bigInteger),
	106:   uint16(anon_sym_bigDecimal),
	107:   uint16(27),
	108:   uint16(3),
	109:   uint16(1),
	110:   uint16(anon_sym_COMMA),
	111:   uint16(5),
	112:   uint16(1),
	113:   uint16(sym_comment),
	114:   uint16(7),
	115:   uint16(1),
	116:   uint16(anon_sym_SLASH_SLASH_SLASH),
	117:   uint16(19),
	118:   uint16(1),
	119:   uint16(anon_sym_use),
	120:   uint16(23),
	121:   uint16(1),
	122:   uint16(anon_sym_list),
	123:   uint16(25),
	124:   uint16(1),
	125:   uint16(anon_sym_map),
	126:   uint16(27),
	127:   uint16(1),
	128:   uint16(anon_sym_set),
	129:   uint16(29),
	130:   uint16(1),
	131:   uint16(anon_sym_structure),
	132:   uint16(31),
	133:   uint16(1),
	134:   uint16(anon_sym_union),
	135:   uint16(33),
	136:   uint16(1),
	137:   uint16(anon_sym_service),
	138:   uint16(35),
	139:   uint16(1),
	140:   uint16(anon_sym_operation),
	141:   uint16(37),
	142:   uint16(1),
	143:   uint16(anon_sym_resource),
	144:   uint16(39),
	145:   uint16(1),
	146:   uint16(anon_sym_AT),
	147:   uint16(41),
	148:   uint16(1),
	149:   uint16(anon_sym_apply),
	150:   uint16(45),
	151:   uint16(1),
	153:   uint16(3),
	154:   uint16(1),
	155:   uint16(sym_documentation_comment),
	156:   uint16(4),
	157:   uint16(1),
	158:   uint16(aux_sym_shape_section_repeat1),
	159:   uint16(10),
	160:   uint16(1),
	161:   uint16(aux_sym_shape_statement_repeat1),
	162:   uint16(37),
	163:   uint16(1),
	164:   uint16(sym_trait_statement),
	165:   uint16(55),
	166:   uint16(1),
	167:   uint16(sym__definition),
	168:   uint16(57),
	169:   uint16(1),
	170:   uint16(sym_shape_body),
	171:   uint16(277),
	172:   uint16(1),
	173:   uint16(sym_primitive),
	174:   uint16(21),
	175:   uint16(2),
	176:   uint16(anon_sym_enum),
	177:   uint16(anon_sym_intEnum),
	178:   uint16(40),
	179:   uint16(2),
	180:   uint16(sym_apply_statement_singular),
	181:   uint16(sym_apply_statement_block),
	182:   uint16(56),
	183:   uint16(3),
	184:   uint16(sym_use_statement),
	185:   uint16(sym_shape_statement),
	186:   uint16(sym_apply_statement),
	187:   uint16(58),
	188:   uint16(10),
	189:   uint16(sym_simple_shape_statement),
	190:   uint16(sym_enum_statement),
	191:   uint16(sym_list_statement),
	192:   uint16(sym_map_statement),
	193:   uint16(sym_set_statement),
	194:   uint16(sym_structure_statement),
	195:   uint16(sym_union_statement),
	196:   uint16(sym_service_statement),
	197:   uint16(sym_operation_statement),
	198:   uint16(sym_resource_statement),
	199:   uint16(43),
	200:   uint16(13),
	201:   uint16(anon_sym_blob),
	202:   uint16(anon_sym_boolean),
	203:   uint16(anon_sym_byte),
	204:   uint16(anon_sym_document),
	205:   uint16(anon_sym_double),
	206:   uint16(anon_sym_float),
	207:   uint16(anon_sym_integer),
	208:   uint16(anon_sym_long),
	209:   uint16(anon_sym_short),
	210:   uint16(anon_sym_string),
	211:   uint16(anon_sym_timestamp),
	212:   uint16(anon_sym_bigInteger),
	213:   uint16(anon_sym_bigDecimal),
	214:   uint16(26),
	215:   uint16(3),
	216:   uint16(1),
	217:   uint16(anon_sym_COMMA),
	218:   uint16(5),
	219:   uint16(1),
	220:   uint16(sym_comment),
	221:   uint16(7),
	222:   uint16(1),
	223:   uint16(anon_sym_SLASH_SLASH_SLASH),
	224:   uint16(47),
	225:   uint16(1),
	227:   uint16(49),
	228:   uint16(1),
	229:   uint16(anon_sym_use),
	230:   uint16(55),
	231:   uint16(1),
	232:   uint16(anon_sym_list),
	233:   uint16(58),
	234:   uint16(1),
	235:   uint16(anon_sym_map),
	236:   uint16(61),
	237:   uint16(1),
	238:   uint16(anon_sym_set),
	239:   uint16(64),
	240:   uint16(1),
	241:   uint16(anon_sym_structure),
	242:   uint16(67),
	243:   uint16(1),
	244:   uint16(anon_sym_union),
	245:   uint16(70),
	246:   uint16(1),
	247:   uint16(anon_sym_service),
	248:   uint16(73),
	249:   uint16(1),
	250:   uint16(anon_sym_operation),
	251:   uint16(76),
	252:   uint16(1),
	253:   uint16(anon_sym_resource),
	254:   uint16(79),
	255:   uint16(1),
	256:   uint16(anon_sym_AT),
	257:   uint16(82),
	258:   uint16(1),
	259:   uint16(anon_sym_apply),
	260:   uint16(10),
	261:   uint16(1),
	262:   uint16(aux_sym_shape_statement_repeat1),
	263:   uint16(37),
	264:   uint16(1),
	265:   uint16(sym_trait_statement),
	266:   uint16(55),
	267:   uint16(1),
	268:   uint16(sym__definition),
	269:   uint16(57),
	270:   uint16(1),
	271:   uint16(sym_shape_body),
	272:   uint16(277),
	273:   uint16(1),
	274:   uint16(sym_primitive),
	275:   uint16(52),
	276:   uint16(2),
	277:   uint16(anon_sym_enum),
	278:   uint16(anon_sym_intEnum),
	279:   uint16(4),
	280:   uint16(2),
	281:   uint16(sym_documentation_comment),
	282:   uint16(aux_sym_shape_section_repeat1),
	283:   uint16(40),
	284:   uint16(2),
	285:   uint16(sym_apply_statement_singular),
	286:   uint16(sym_apply_statement_block),
	287:   uint16(56),
	288:   uint16(3),
	289:   uint16(sym_use_statement),
	290:   uint16(sym_shape_statement),
	291:   uint16(sym_apply_statement),
	292:   uint16(58),
	293:   uint16(10),
	294:   uint16(sym_simple_shape_statement),
	295:   uint16(sym_enum_statement),
	296:   uint16(sym_list_statement),
	297:   uint16(sym_map_statement),
	298:   uint16(sym_set_statement),
	299:   uint16(sym_structure_statement),
	300:   uint16(sym_union_statement),
	301:   uint16(sym_service_statement),
	302:   uint16(sym_operation_statement),
	303:   uint16(sym_resource_statement),
	304:   uint16(85),
	305:   uint16(13),
	306:   uint16(anon_sym_blob),
	307:   uint16(anon_sym_boolean),
	308:   uint16(anon_sym_byte),
	309:   uint16(anon_sym_document),
	310:   uint16(anon_sym_double),
	311:   uint16(anon_sym_float),
	312:   uint16(anon_sym_integer),
	313:   uint16(anon_sym_long),
	314:   uint16(anon_sym_short),
	315:   uint16(anon_sym_string),
	316:   uint16(anon_sym_timestamp),
	317:   uint16(anon_sym_bigInteger),
	318:   uint16(anon_sym_bigDecimal),
	319:   uint16(6),
	320:   uint16(3),
	321:   uint16(1),
	322:   uint16(anon_sym_COMMA),
	323:   uint16(5),
	324:   uint16(1),
	325:   uint16(sym_comment),
	326:   uint16(7),
	327:   uint16(1),
	328:   uint16(anon_sym_SLASH_SLASH_SLASH),
	329:   uint16(5),
	330:   uint16(1),
	331:   uint16(sym_documentation_comment),
	332:   uint16(90),
	333:   uint16(2),
	334:   uint16(anon_sym_DOT),
	335:   uint16(anon_sym_POUND),
	336:   uint16(88),
	337:   uint16(37),
	339:   uint16(anon_sym_DOLLAR),
	340:   uint16(anon_sym_COLON),
	341:   uint16(anon_sym_metadata),
	342:   uint16(anon_sym_namespace),
	343:   uint16(anon_sym_use),
	344:   uint16(anon_sym_enum),
	345:   uint16(anon_sym_intEnum),
	346:   uint16(anon_sym_LBRACE),
	347:   uint16(anon_sym_RBRACE),
	348:   uint16(anon_sym_list),
	349:   uint16(anon_sym_map),
	350:   uint16(anon_sym_set),
	351:   uint16(anon_sym_structure),
	352:   uint16(anon_sym_union),
	353:   uint16(anon_sym_service),
	354:   uint16(anon_sym_operation),
	355:   uint16(anon_sym_resource),
	356:   uint16(anon_sym_AT),
	357:   uint16(anon_sym_LPAREN),
	358:   uint16(anon_sym_RPAREN),
	359:   uint16(anon_sym_apply),
	360:   uint16(anon_sym_with),
	361:   uint16(anon_sym_for),
	362:   uint16(anon_sym_blob),
	363:   uint16(anon_sym_boolean),
	364:   uint16(anon_sym_byte),
	365:   uint16(anon_sym_document),
	366:   uint16(anon_sym_double),
	367:   uint16(anon_sym_float),
	368:   uint16(anon_sym_integer),
	369:   uint16(anon_sym_long),
	370:   uint16(anon_sym_short),
	371:   uint16(anon_sym_string),
	372:   uint16(anon_sym_timestamp),
	373:   uint16(anon_sym_bigInteger),
	374:   uint16(anon_sym_bigDecimal),
	375:   uint16(6),
	376:   uint16(3),
	377:   uint16(1),
	378:   uint16(anon_sym_COMMA),
	379:   uint16(5),
	380:   uint16(1),
	381:   uint16(sym_comment),
	382:   uint16(7),
	383:   uint16(1),
	384:   uint16(anon_sym_SLASH_SLASH_SLASH),
	385:   uint16(92),
	386:   uint16(1),
	387:   uint16(anon_sym_COLON),
	388:   uint16(6),
	389:   uint16(1),
	390:   uint16(sym_documentation_comment),
	391:   uint16(88),
	392:   uint16(38),
	394:   uint16(anon_sym_DOLLAR),
	395:   uint16(anon_sym_metadata),
	396:   uint16(anon_sym_EQ),
	397:   uint16(anon_sym_namespace),
	398:   uint16(anon_sym_use),
	399:   uint16(anon_sym_enum),
	400:   uint16(anon_sym_intEnum),
	401:   uint16(anon_sym_LBRACE),
	402:   uint16(anon_sym_RBRACE),
	403:   uint16(anon_sym_list),
	404:   uint16(anon_sym_map),
	405:   uint16(anon_sym_set),
	406:   uint16(anon_sym_structure),
	407:   uint16(anon_sym_union),
	408:   uint16(anon_sym_service),
	409:   uint16(anon_sym_operation),
	410:   uint16(anon_sym_resource),
	411:   uint16(anon_sym_COLON_EQ),
	412:   uint16(anon_sym_AT),
	413:   uint16(anon_sym_LPAREN),
	414:   uint16(anon_sym_RPAREN),
	415:   uint16(anon_sym_apply),
	416:   uint16(anon_sym_with),
	417:   uint16(anon_sym_for),
	418:   uint16(anon_sym_blob),
	419:   uint16(anon_sym_boolean),
	420:   uint16(anon_sym_byte),
	421:   uint16(anon_sym_document),
	422:   uint16(anon_sym_double),
	423:   uint16(anon_sym_float),
	424:   uint16(anon_sym_integer),
	425:   uint16(anon_sym_long),
	426:   uint16(anon_sym_short),
	427:   uint16(anon_sym_string),
	428:   uint16(anon_sym_timestamp),
	429:   uint16(anon_sym_bigInteger),
	430:   uint16(anon_sym_bigDecimal),
	431:   uint16(7),
	432:   uint16(3),
	433:   uint16(1),
	434:   uint16(anon_sym_COMMA),
	435:   uint16(5),
	436:   uint16(1),
	437:   uint16(sym_comment),
	438:   uint16(7),
	439:   uint16(1),
	440:   uint16(anon_sym_SLASH_SLASH_SLASH),
	441:   uint16(96),
	442:   uint16(1),
	443:   uint16(anon_sym_DOLLAR),
	444:   uint16(12),
	445:   uint16(1),
	446:   uint16(sym_shape_id_member),
	447:   uint16(7),
	448:   uint16(2),
	449:   uint16(sym_documentation_comment),
	450:   uint16(aux_sym_shape_id_repeat1),
	451:   uint16(94),
	452:   uint16(35),
	454:   uint16(anon_sym_metadata),
	455:   uint16(anon_sym_namespace),
	456:   uint16(anon_sym_use),
	457:   uint16(anon_sym_enum),
	458:   uint16(anon_sym_intEnum),
	459:   uint16(anon_sym_LBRACE),
	460:   uint16(anon_sym_RBRACE),
	461:   uint16(anon_sym_list),
	462:   uint16(anon_sym_map),
	463:   uint16(anon_sym_set),
	464:   uint16(anon_sym_structure),
	465:   uint16(anon_sym_union),
	466:   uint16(anon_sym_service),
	467:   uint16(anon_sym_operation),
	468:   uint16(anon_sym_resource),
	469:   uint16(anon_sym_AT),
	470:   uint16(anon_sym_LPAREN),
	471:   uint16(anon_sym_RPAREN),
	472:   uint16(anon_sym_apply),
	473:   uint16(anon_sym_with),
	474:   uint16(anon_sym_for),
	475:   uint16(anon_sym_blob),
	476:   uint16(anon_sym_boolean),
	477:   uint16(anon_sym_byte),
	478:   uint16(anon_sym_document),
	479:   uint16(anon_sym_double),
	480:   uint16(anon_sym_float),
	481:   uint16(anon_sym_integer),
	482:   uint16(anon_sym_long),
	483:   uint16(anon_sym_short),
	484:   uint16(anon_sym_string),
	485:   uint16(anon_sym_timestamp),
	486:   uint16(anon_sym_bigInteger),
	487:   uint16(anon_sym_bigDecimal),
	488:   uint16(8),
	489:   uint16(3),
	490:   uint16(1),
	491:   uint16(anon_sym_COMMA),
	492:   uint16(5),
	493:   uint16(1),
	494:   uint16(sym_comment),
	495:   uint16(7),
	496:   uint16(1),
	497:   uint16(anon_sym_SLASH_SLASH_SLASH),
	498:   uint16(101),
	499:   uint16(1),
	500:   uint16(anon_sym_DOLLAR),
	501:   uint16(7),
	502:   uint16(1),
	503:   uint16(aux_sym_shape_id_repeat1),
	504:   uint16(8),
	505:   uint16(1),
	506:   uint16(sym_documentation_comment),
	507:   uint16(12),
	508:   uint16(1),
	509:   uint16(sym_shape_id_member),
	510:   uint16(99),
	511:   uint16(35),
	513:   uint16(anon_sym_metadata),
	514:   uint16(anon_sym_namespace),
	515:   uint16(anon_sym_use),
	516:   uint16(anon_sym_enum),
	517:   uint16(anon_sym_intEnum),
	518:   uint16(anon_sym_LBRACE),
	519:   uint16(anon_sym_RBRACE),
	520:   uint16(anon_sym_list),
	521:   uint16(anon_sym_map),
	522:   uint16(anon_sym_set),
	523:   uint16(anon_sym_structure),
	524:   uint16(anon_sym_union),
	525:   uint16(anon_sym_service),
	526:   uint16(anon_sym_operation),
	527:   uint16(anon_sym_resource),
	528:   uint16(anon_sym_AT),
	529:   uint16(anon_sym_LPAREN),
	530:   uint16(anon_sym_RPAREN),
	531:   uint16(anon_sym_apply),
	532:   uint16(anon_sym_with),
	533:   uint16(anon_sym_for),
	534:   uint16(anon_sym_blob),
	535:   uint16(anon_sym_boolean),
	536:   uint16(anon_sym_byte),
	537:   uint16(anon_sym_document),
	538:   uint16(anon_sym_double),
	539:   uint16(anon_sym_float),
	540:   uint16(anon_sym_integer),
	541:   uint16(anon_sym_long),
	542:   uint16(anon_sym_short),
	543:   uint16(anon_sym_string),
	544:   uint16(anon_sym_timestamp),
	545:   uint16(anon_sym_bigInteger),
	546:   uint16(anon_sym_bigDecimal),
	547:   uint16(8),
	548:   uint16(3),
	549:   uint16(1),
	550:   uint16(anon_sym_COMMA),
	551:   uint16(5),
	552:   uint16(1),
	553:   uint16(sym_comment),
	554:   uint16(7),
	555:   uint16(1),
	556:   uint16(anon_sym_SLASH_SLASH_SLASH),
	557:   uint16(101),
	558:   uint16(1),
	559:   uint16(anon_sym_DOLLAR),
	560:   uint16(8),
	561:   uint16(1),
	562:   uint16(aux_sym_shape_id_repeat1),
	563:   uint16(9),
	564:   uint16(1),
	565:   uint16(sym_documentation_comment),
	566:   uint16(12),
	567:   uint16(1),
	568:   uint16(sym_shape_id_member),
	569:   uint16(103),
	570:   uint16(35),
	572:   uint16(anon_sym_metadata),
	573:   uint16(anon_sym_namespace),
	574:   uint16(anon_sym_use),
	575:   uint16(anon_sym_enum),
	576:   uint16(anon_sym_intEnum),
	577:   uint16(anon_sym_LBRACE),
	578:   uint16(anon_sym_RBRACE),
	579:   uint16(anon_sym_list),
	580:   uint16(anon_sym_map),
	581:   uint16(anon_sym_set),
	582:   uint16(anon_sym_structure),
	583:   uint16(anon_sym_union),
	584:   uint16(anon_sym_service),
	585:   uint16(anon_sym_operation),
	586:   uint16(anon_sym_resource),
	587:   uint16(anon_sym_AT),
	588:   uint16(anon_sym_LPAREN),
	589:   uint16(anon_sym_RPAREN),
	590:   uint16(anon_sym_apply),
	591:   uint16(anon_sym_with),
	592:   uint16(anon_sym_for),
	593:   uint16(anon_sym_blob),
	594:   uint16(anon_sym_boolean),
	595:   uint16(anon_sym_byte),
	596:   uint16(anon_sym_document),
	597:   uint16(anon_sym_double),
	598:   uint16(anon_sym_float),
	599:   uint16(anon_sym_integer),
	600:   uint16(anon_sym_long),
	601:   uint16(anon_sym_short),
	602:   uint16(anon_sym_string),
	603:   uint16(anon_sym_timestamp),
	604:   uint16(anon_sym_bigInteger),
	605:   uint16(anon_sym_bigDecimal),
	606:   uint16(20),
	607:   uint16(3),
	608:   uint16(1),
	609:   uint16(anon_sym_COMMA),
	610:   uint16(5),
	611:   uint16(1),
	612:   uint16(sym_comment),
	613:   uint16(7),
	614:   uint16(1),
	615:   uint16(anon_sym_SLASH_SLASH_SLASH),
	616:   uint16(23),
	617:   uint16(1),
	618:   uint16(anon_sym_list),
	619:   uint16(25),
	620:   uint16(1),
	621:   uint16(anon_sym_map),
	622:   uint16(27),
	623:   uint16(1),
	624:   uint16(anon_sym_set),
	625:   uint16(29),
	626:   uint16(1),
	627:   uint16(anon_sym_structure),
	628:   uint16(31),
	629:   uint16(1),
	630:   uint16(anon_sym_union),
	631:   uint16(33),
	632:   uint16(1),
	633:   uint16(anon_sym_service),
	634:   uint16(35),
	635:   uint16(1),
	636:   uint16(anon_sym_operation),
	637:   uint16(37),
	638:   uint16(1),
	639:   uint16(anon_sym_resource),
	640:   uint16(39),
	641:   uint16(1),
	642:   uint16(anon_sym_AT),
	643:   uint16(10),
	644:   uint16(1),
	645:   uint16(sym_documentation_comment),
	646:   uint16(26),
	647:   uint16(1),
	648:   uint16(aux_sym_shape_statement_repeat1),
	649:   uint16(37),
	650:   uint16(1),
	651:   uint16(sym_trait_statement),
	652:   uint16(80),
	653:   uint16(1),
	654:   uint16(sym_shape_body),
	655:   uint16(277),
	656:   uint16(1),
	657:   uint16(sym_primitive),
	658:   uint16(21),
	659:   uint16(2),
	660:   uint16(anon_sym_enum),
	661:   uint16(anon_sym_intEnum),
	662:   uint16(58),
	663:   uint16(10),
	664:   uint16(sym_simple_shape_statement),
	665:   uint16(sym_enum_statement),
	666:   uint16(sym_list_statement),
	667:   uint16(sym_map_statement),
	668:   uint16(sym_set_statement),
	669:   uint16(sym_structure_statement),
	670:   uint16(sym_union_statement),
	671:   uint16(sym_service_statement),
	672:   uint16(sym_operation_statement),
	673:   uint16(sym_resource_statement),
	674:   uint16(43),
	675:   uint16(13),
	676:   uint16(anon_sym_blob),
	677:   uint16(anon_sym_boolean),
	678:   uint16(anon_sym_byte),
	679:   uint16(anon_sym_document),
	680:   uint16(anon_sym_double),
	681:   uint16(anon_sym_float),
	682:   uint16(anon_sym_integer),
	683:   uint16(anon_sym_long),
	684:   uint16(anon_sym_short),
	685:   uint16(anon_sym_string),
	686:   uint16(anon_sym_timestamp),
	687:   uint16(anon_sym_bigInteger),
	688:   uint16(anon_sym_bigDecimal),
	689:   uint16(5),
	690:   uint16(3),
	691:   uint16(1),
	692:   uint16(anon_sym_COMMA),
	693:   uint16(5),
	694:   uint16(1),
	695:   uint16(sym_comment),
	696:   uint16(7),
	697:   uint16(1),
	698:   uint16(anon_sym_SLASH_SLASH_SLASH),
	699:   uint16(11),
	700:   uint16(1),
	701:   uint16(sym_documentation_comment),
	702:   uint16(105),
	703:   uint16(36),
	705:   uint16(anon_sym_DOLLAR),
	706:   uint16(anon_sym_metadata),
	707:   uint16(anon_sym_namespace),
	708:   uint16(anon_sym_use),
	709:   uint16(anon_sym_enum),
	710:   uint16(anon_sym_intEnum),
	711:   uint16(anon_sym_LBRACE),
	712:   uint16(anon_sym_RBRACE),
	713:   uint16(anon_sym_list),
	714:   uint16(anon_sym_map),
	715:   uint16(anon_sym_set),
	716:   uint16(anon_sym_structure),
	717:   uint16(anon_sym_union),
	718:   uint16(anon_sym_service),
	719:   uint16(anon_sym_operation),
	720:   uint16(anon_sym_resource),
	721:   uint16(anon_sym_AT),
	722:   uint16(anon_sym_LPAREN),
	723:   uint16(anon_sym_RPAREN),
	724:   uint16(anon_sym_apply),
	725:   uint16(anon_sym_with),
	726:   uint16(anon_sym_for),
	727:   uint16(anon_sym_blob),
	728:   uint16(anon_sym_boolean),
	729:   uint16(anon_sym_byte),
	730:   uint16(anon_sym_document),
	731:   uint16(anon_sym_double),
	732:   uint16(anon_sym_float),
	733:   uint16(anon_sym_integer),
	734:   uint16(anon_sym_long),
	735:   uint16(anon_sym_short),
	736:   uint16(anon_sym_string),
	737:   uint16(anon_sym_timestamp),
	738:   uint16(anon_sym_bigInteger),
	739:   uint16(anon_sym_bigDecimal),
	740:   uint16(5),
	741:   uint16(3),
	742:   uint16(1),
	743:   uint16(anon_sym_COMMA),
	744:   uint16(5),
	745:   uint16(1),
	746:   uint16(sym_comment),
	747:   uint16(7),
	748:   uint16(1),
	749:   uint16(anon_sym_SLASH_SLASH_SLASH),
	750:   uint16(12),
	751:   uint16(1),
	752:   uint16(sym_documentation_comment),
	753:   uint16(107),
	754:   uint16(36),
	756:   uint16(anon_sym_DOLLAR),
	757:   uint16(anon_sym_metadata),
	758:   uint16(anon_sym_namespace),
	759:   uint16(anon_sym_use),
	760:   uint16(anon_sym_enum),
	761:   uint16(anon_sym_intEnum),
	762:   uint16(anon_sym_LBRACE),
	763:   uint16(anon_sym_RBRACE),
	764:   uint16(anon_sym_list),
	765:   uint16(anon_sym_map),
	766:   uint16(anon_sym_set),
	767:   uint16(anon_sym_structure),
	768:   uint16(anon_sym_union),
	769:   uint16(anon_sym_service),
	770:   uint16(anon_sym_operation),
	771:   uint16(anon_sym_resource),
	772:   uint16(anon_sym_AT),
	773:   uint16(anon_sym_LPAREN),
	774:   uint16(anon_sym_RPAREN),
	775:   uint16(anon_sym_apply),
	776:   uint16(anon_sym_with),
	777:   uint16(anon_sym_for),
	778:   uint16(anon_sym_blob),
	779:   uint16(anon_sym_boolean),
	780:   uint16(anon_sym_byte),
	781:   uint16(anon_sym_document),
	782:   uint16(anon_sym_double),
	783:   uint16(anon_sym_float),
	784:   uint16(anon_sym_integer),
	785:   uint16(anon_sym_long),
	786:   uint16(anon_sym_short),
	787:   uint16(anon_sym_string),
	788:   uint16(anon_sym_timestamp),
	789:   uint16(anon_sym_bigInteger),
	790:   uint16(anon_sym_bigDecimal),
	791:   uint16(5),
	792:   uint16(3),
	793:   uint16(1),
	794:   uint16(anon_sym_COMMA),
	795:   uint16(5),
	796:   uint16(1),
	797:   uint16(sym_comment),
	798:   uint16(7),
	799:   uint16(1),
	800:   uint16(anon_sym_SLASH_SLASH_SLASH),
	801:   uint16(13),
	802:   uint16(1),
	803:   uint16(sym_documentation_comment),
	804:   uint16(109),
	805:   uint16(36),
	807:   uint16(anon_sym_DOLLAR),
	808:   uint16(anon_sym_metadata),
	809:   uint16(anon_sym_namespace),
	810:   uint16(anon_sym_use),
	811:   uint16(anon_sym_enum),
	812:   uint16(anon_sym_intEnum),
	813:   uint16(anon_sym_LBRACE),
	814:   uint16(anon_sym_RBRACE),
	815:   uint16(anon_sym_list),
	816:   uint16(anon_sym_map),
	817:   uint16(anon_sym_set),
	818:   uint16(anon_sym_structure),
	819:   uint16(anon_sym_union),
	820:   uint16(anon_sym_service),
	821:   uint16(anon_sym_operation),
	822:   uint16(anon_sym_resource),
	823:   uint16(anon_sym_AT),
	824:   uint16(anon_sym_LPAREN),
	825:   uint16(anon_sym_RPAREN),
	826:   uint16(anon_sym_apply),
	827:   uint16(anon_sym_with),
	828:   uint16(anon_sym_for),
	829:   uint16(anon_sym_blob),
	830:   uint16(anon_sym_boolean),
	831:   uint16(anon_sym_byte),
	832:   uint16(anon_sym_document),
	833:   uint16(anon_sym_double),
	834:   uint16(anon_sym_float),
	835:   uint16(anon_sym_integer),
	836:   uint16(anon_sym_long),
	837:   uint16(anon_sym_short),
	838:   uint16(anon_sym_string),
	839:   uint16(anon_sym_timestamp),
	840:   uint16(anon_sym_bigInteger),
	841:   uint16(anon_sym_bigDecimal),
	842:   uint16(5),
	843:   uint16(3),
	844:   uint16(1),
	845:   uint16(anon_sym_COMMA),
	846:   uint16(5),
	847:   uint16(1),
	848:   uint16(sym_comment),
	849:   uint16(7),
	850:   uint16(1),
	851:   uint16(anon_sym_SLASH_SLASH_SLASH),
	852:   uint16(14),
	853:   uint16(1),
	854:   uint16(sym_documentation_comment),
	855:   uint16(111),
	856:   uint16(36),
	858:   uint16(anon_sym_DOLLAR),
	859:   uint16(anon_sym_metadata),
	860:   uint16(anon_sym_namespace),
	861:   uint16(anon_sym_use),
	862:   uint16(anon_sym_enum),
	863:   uint16(anon_sym_intEnum),
	864:   uint16(anon_sym_LBRACE),
	865:   uint16(anon_sym_RBRACE),
	866:   uint16(anon_sym_list),
	867:   uint16(anon_sym_map),
	868:   uint16(anon_sym_set),
	869:   uint16(anon_sym_structure),
	870:   uint16(anon_sym_union),
	871:   uint16(anon_sym_service),
	872:   uint16(anon_sym_operation),
	873:   uint16(anon_sym_resource),
	874:   uint16(anon_sym_AT),
	875:   uint16(anon_sym_LPAREN),
	876:   uint16(anon_sym_RPAREN),
	877:   uint16(anon_sym_apply),
	878:   uint16(anon_sym_with),
	879:   uint16(anon_sym_for),
	880:   uint16(anon_sym_blob),
	881:   uint16(anon_sym_boolean),
	882:   uint16(anon_sym_byte),
	883:   uint16(anon_sym_document),
	884:   uint16(anon_sym_double),
	885:   uint16(anon_sym_float),
	886:   uint16(anon_sym_integer),
	887:   uint16(anon_sym_long),
	888:   uint16(anon_sym_short),
	889:   uint16(anon_sym_string),
	890:   uint16(anon_sym_timestamp),
	891:   uint16(anon_sym_bigInteger),
	892:   uint16(anon_sym_bigDecimal),
	893:   uint16(7),
	894:   uint16(3),
	895:   uint16(1),
	896:   uint16(anon_sym_COMMA),
	897:   uint16(5),
	898:   uint16(1),
	899:   uint16(sym_comment),
	900:   uint16(7),
	901:   uint16(1),
	902:   uint16(anon_sym_SLASH_SLASH_SLASH),
	903:   uint16(115),
	904:   uint16(1),
	905:   uint16(anon_sym_LPAREN),
	906:   uint16(15),
	907:   uint16(1),
	908:   uint16(sym_documentation_comment),
	909:   uint16(21),
	910:   uint16(1),
	911:   uint16(sym_trait_body),
	912:   uint16(113),
	913:   uint16(31),
	915:   uint16(anon_sym_use),
	916:   uint16(anon_sym_enum),
	917:   uint16(anon_sym_intEnum),
	918:   uint16(anon_sym_LBRACE),
	919:   uint16(anon_sym_RBRACE),
	920:   uint16(anon_sym_list),
	921:   uint16(anon_sym_map),
	922:   uint16(anon_sym_set),
	923:   uint16(anon_sym_structure),
	924:   uint16(anon_sym_union),
	925:   uint16(anon_sym_service),
	926:   uint16(anon_sym_operation),
	927:   uint16(anon_sym_resource),
	928:   uint16(anon_sym_AT),
	929:   uint16(anon_sym_apply),
	930:   uint16(anon_sym_with),
	931:   uint16(anon_sym_for),
	932:   uint16(anon_sym_blob),
	933:   uint16(anon_sym_boolean),
	934:   uint16(anon_sym_byte),
	935:   uint16(anon_sym_document),
	936:   uint16(anon_sym_double),
	937:   uint16(anon_sym_float),
	938:   uint16(anon_sym_integer),
	939:   uint16(anon_sym_long),
	940:   uint16(anon_sym_short),
	941:   uint16(anon_sym_string),
	942:   uint16(anon_sym_timestamp),
	943:   uint16(anon_sym_bigInteger),
	944:   uint16(anon_sym_bigDecimal),
	945:   uint16(30),
	946:   uint16(3),
	947:   uint16(1),
	948:   uint16(anon_sym_COMMA),
	949:   uint16(5),
	950:   uint16(1),
	951:   uint16(sym_comment),
	952:   uint16(7),
	953:   uint16(1),
	954:   uint16(anon_sym_SLASH_SLASH_SLASH),
	955:   uint16(117),
	956:   uint16(1),
	957:   uint16(anon_sym_LBRACE),
	958:   uint16(119),
	959:   uint16(1),
	960:   uint16(anon_sym_LBRACK),
	961:   uint16(121),
	962:   uint16(1),
	963:   uint16(anon_sym_RPAREN),
	964:   uint16(125),
	965:   uint16(1),
	966:   uint16(sym_null),
	967:   uint16(127),
	968:   uint16(1),
	969:   uint16(anon_sym_DASH),
	970:   uint16(129),
	971:   uint16(1),
	972:   uint16(aux_sym_number_token1),
	973:   uint16(131),
	974:   uint16(1),
	975:   uint16(aux_sym_float_token1),
	976:   uint16(133),
	977:   uint16(1),
	978:   uint16(anon_sym_DQUOTE),
	979:   uint16(135),
	980:   uint16(1),
	981:   uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	982:   uint16(137),
	983:   uint16(1),
	984:   uint16(aux_sym_identifier_token1),
	985:   uint16(9),
	986:   uint16(1),
	987:   uint16(sym_root_shape_id),
	988:   uint16(11),
	989:   uint16(1),
	990:   uint16(sym_absolute_root_shape_id),
	991:   uint16(16),
	992:   uint16(1),
	993:   uint16(sym_documentation_comment),
	994:   uint16(23),
	995:   uint16(1),
	996:   uint16(sym__namespace_identifier),
	997:   uint16(112),
	998:   uint16(1),
	999:   uint16(aux_sym_trait_structure_repeat1),
	1000:  uint16(253),
	1001:  uint16(1),
	1002:  uint16(sym_node_object_kvp),
	1003:  uint16(267),
	1004:  uint16(1),
	1005:  uint16(sym_identifier),
	1006:  uint16(289),
	1007:  uint16(1),
	1008:  uint16(sym_string),
	1009:  uint16(326),
	1010:  uint16(1),
	1011:  uint16(sym_node_value),
	1012:  uint16(332),
	1013:  uint16(1),
	1014:  uint16(sym_namespace),
	1015:  uint16(338),
	1016:  uint16(1),
	1017:  uint16(sym_trait_body_value),
	1018:  uint16(339),
	1019:  uint16(1),
	1020:  uint16(sym_trait_structure),
	1021:  uint16(340),
	1022:  uint16(1),
	1023:  uint16(sym_node_object_key),
	1024:  uint16(123),
	1025:  uint16(2),
	1026:  uint16(anon_sym_true),
	1027:  uint16(anon_sym_false),
	1028:  uint16(168),
	1029:  uint16(2),
	1030:  uint16(sym__string_literal),
	1031:  uint16(sym__multiline_string_literal),
	1032:  uint16(209),
	1033:  uint16(3),
	1034:  uint16(sym_boolean),
	1035:  uint16(sym_number),
	1036:  uint16(sym_float),
	1037:  uint16(226),
	1038:  uint16(4),
	1039:  uint16(sym_shape_id),
	1040:  uint16(sym_node_array),
	1041:  uint16(sym_node_object),
	1042:  uint16(sym_literal),
	1043:  uint16(30),
	1044:  uint16(3),
	1045:  uint16(1),
	1046:  uint16(anon_sym_COMMA),
	1047:  uint16(5),
	1048:  uint16(1),
	1049:  uint16(sym_comment),
	1050:  uint16(7),
	1051:  uint16(1),
	1052:  uint16(anon_sym_SLASH_SLASH_SLASH),
	1053:  uint16(117),
	1054:  uint16(1),
	1055:  uint16(anon_sym_LBRACE),
	1056:  uint16(119),
	1057:  uint16(1),
	1058:  uint16(anon_sym_LBRACK),
	1059:  uint16(125),
	1060:  uint16(1),
	1061:  uint16(sym_null),
	1062:  uint16(127),
	1063:  uint16(1),
	1064:  uint16(anon_sym_DASH),
	1065:  uint16(129),
	1066:  uint16(1),
	1067:  uint16(aux_sym_number_token1),
	1068:  uint16(131),
	1069:  uint16(1),
	1070:  uint16(aux_sym_float_token1),
	1071:  uint16(133),
	1072:  uint16(1),
	1073:  uint16(anon_sym_DQUOTE),
	1074:  uint16(135),
	1075:  uint16(1),
	1076:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1077:  uint16(137),
	1078:  uint16(1),
	1079:  uint16(aux_sym_identifier_token1),
	1080:  uint16(139),
	1081:  uint16(1),
	1082:  uint16(anon_sym_RPAREN),
	1083:  uint16(9),
	1084:  uint16(1),
	1085:  uint16(sym_root_shape_id),
	1086:  uint16(11),
	1087:  uint16(1),
	1088:  uint16(sym_absolute_root_shape_id),
	1089:  uint16(17),
	1090:  uint16(1),
	1091:  uint16(sym_documentation_comment),
	1092:  uint16(23),
	1093:  uint16(1),
	1094:  uint16(sym__namespace_identifier),
	1095:  uint16(112),
	1096:  uint16(1),
	1097:  uint16(aux_sym_trait_structure_repeat1),
	1098:  uint16(253),
	1099:  uint16(1),
	1100:  uint16(sym_node_object_kvp),
	1101:  uint16(267),
	1102:  uint16(1),
	1103:  uint16(sym_identifier),
	1104:  uint16(289),
	1105:  uint16(1),
	1106:  uint16(sym_string),
	1107:  uint16(326),
	1108:  uint16(1),
	1109:  uint16(sym_node_value),
	1110:  uint16(332),
	1111:  uint16(1),
	1112:  uint16(sym_namespace),
	1113:  uint16(336),
	1114:  uint16(1),
	1115:  uint16(sym_trait_body_value),
	1116:  uint16(339),
	1117:  uint16(1),
	1118:  uint16(sym_trait_structure),
	1119:  uint16(340),
	1120:  uint16(1),
	1121:  uint16(sym_node_object_key),
	1122:  uint16(123),
	1123:  uint16(2),
	1124:  uint16(anon_sym_true),
	1125:  uint16(anon_sym_false),
	1126:  uint16(168),
	1127:  uint16(2),
	1128:  uint16(sym__string_literal),
	1129:  uint16(sym__multiline_string_literal),
	1130:  uint16(209),
	1131:  uint16(3),
	1132:  uint16(sym_boolean),
	1133:  uint16(sym_number),
	1134:  uint16(sym_float),
	1135:  uint16(226),
	1136:  uint16(4),
	1137:  uint16(sym_shape_id),
	1138:  uint16(sym_node_array),
	1139:  uint16(sym_node_object),
	1140:  uint16(sym_literal),
	1141:  uint16(5),
	1142:  uint16(3),
	1143:  uint16(1),
	1144:  uint16(anon_sym_COMMA),
	1145:  uint16(5),
	1146:  uint16(1),
	1147:  uint16(sym_comment),
	1148:  uint16(7),
	1149:  uint16(1),
	1150:  uint16(anon_sym_SLASH_SLASH_SLASH),
	1151:  uint16(18),
	1152:  uint16(1),
	1153:  uint16(sym_documentation_comment),
	1154:  uint16(141),
	1155:  uint16(31),
	1157:  uint16(anon_sym_DOLLAR),
	1158:  uint16(anon_sym_metadata),
	1159:  uint16(anon_sym_namespace),
	1160:  uint16(anon_sym_use),
	1161:  uint16(anon_sym_enum),
	1162:  uint16(anon_sym_intEnum),
	1163:  uint16(anon_sym_list),
	1164:  uint16(anon_sym_map),
	1165:  uint16(anon_sym_set),
	1166:  uint16(anon_sym_structure),
	1167:  uint16(anon_sym_union),
	1168:  uint16(anon_sym_service),
	1169:  uint16(anon_sym_operation),
	1170:  uint16(anon_sym_resource),
	1171:  uint16(anon_sym_AT),
	1172:  uint16(anon_sym_RPAREN),
	1173:  uint16(anon_sym_apply),
	1174:  uint16(anon_sym_blob),
	1175:  uint16(anon_sym_boolean),
	1176:  uint16(anon_sym_byte),
	1177:  uint16(anon_sym_document),
	1178:  uint16(anon_sym_double),
	1179:  uint16(anon_sym_float),
	1180:  uint16(anon_sym_integer),
	1181:  uint16(anon_sym_long),
	1182:  uint16(anon_sym_short),
	1183:  uint16(anon_sym_string),
	1184:  uint16(anon_sym_timestamp),
	1185:  uint16(anon_sym_bigInteger),
	1186:  uint16(anon_sym_bigDecimal),
	1187:  uint16(5),
	1188:  uint16(3),
	1189:  uint16(1),
	1190:  uint16(anon_sym_COMMA),
	1191:  uint16(5),
	1192:  uint16(1),
	1193:  uint16(sym_comment),
	1194:  uint16(7),
	1195:  uint16(1),
	1196:  uint16(anon_sym_SLASH_SLASH_SLASH),
	1197:  uint16(19),
	1198:  uint16(1),
	1199:  uint16(sym_documentation_comment),
	1200:  uint16(143),
	1201:  uint16(31),
	1203:  uint16(anon_sym_DOLLAR),
	1204:  uint16(anon_sym_metadata),
	1205:  uint16(anon_sym_namespace),
	1206:  uint16(anon_sym_use),
	1207:  uint16(anon_sym_enum),
	1208:  uint16(anon_sym_intEnum),
	1209:  uint16(anon_sym_list),
	1210:  uint16(anon_sym_map),
	1211:  uint16(anon_sym_set),
	1212:  uint16(anon_sym_structure),
	1213:  uint16(anon_sym_union),
	1214:  uint16(anon_sym_service),
	1215:  uint16(anon_sym_operation),
	1216:  uint16(anon_sym_resource),
	1217:  uint16(anon_sym_AT),
	1218:  uint16(anon_sym_RPAREN),
	1219:  uint16(anon_sym_apply),
	1220:  uint16(anon_sym_blob),
	1221:  uint16(anon_sym_boolean),
	1222:  uint16(anon_sym_byte),
	1223:  uint16(anon_sym_document),
	1224:  uint16(anon_sym_double),
	1225:  uint16(anon_sym_float),
	1226:  uint16(anon_sym_integer),
	1227:  uint16(anon_sym_long),
	1228:  uint16(anon_sym_short),
	1229:  uint16(anon_sym_string),
	1230:  uint16(anon_sym_timestamp),
	1231:  uint16(anon_sym_bigInteger),
	1232:  uint16(anon_sym_bigDecimal),
	1233:  uint16(5),
	1234:  uint16(3),
	1235:  uint16(1),
	1236:  uint16(anon_sym_COMMA),
	1237:  uint16(5),
	1238:  uint16(1),
	1239:  uint16(sym_comment),
	1240:  uint16(7),
	1241:  uint16(1),
	1242:  uint16(anon_sym_SLASH_SLASH_SLASH),
	1243:  uint16(20),
	1244:  uint16(1),
	1245:  uint16(sym_documentation_comment),
	1246:  uint16(145),
	1247:  uint16(31),
	1249:  uint16(anon_sym_use),
	1250:  uint16(anon_sym_enum),
	1251:  uint16(anon_sym_intEnum),
	1252:  uint16(anon_sym_LBRACE),
	1253:  uint16(anon_sym_RBRACE),
	1254:  uint16(anon_sym_list),
	1255:  uint16(anon_sym_map),
	1256:  uint16(anon_sym_set),
	1257:  uint16(anon_sym_structure),
	1258:  uint16(anon_sym_union),
	1259:  uint16(anon_sym_service),
	1260:  uint16(anon_sym_operation),
	1261:  uint16(anon_sym_resource),
	1262:  uint16(anon_sym_AT),
	1263:  uint16(anon_sym_apply),
	1264:  uint16(anon_sym_with),
	1265:  uint16(anon_sym_for),
	1266:  uint16(anon_sym_blob),
	1267:  uint16(anon_sym_boolean),
	1268:  uint16(anon_sym_byte),
	1269:  uint16(anon_sym_document),
	1270:  uint16(anon_sym_double),
	1271:  uint16(anon_sym_float),
	1272:  uint16(anon_sym_integer),
	1273:  uint16(anon_sym_long),
	1274:  uint16(anon_sym_short),
	1275:  uint16(anon_sym_string),
	1276:  uint16(anon_sym_timestamp),
	1277:  uint16(anon_sym_bigInteger),
	1278:  uint16(anon_sym_bigDecimal),
	1279:  uint16(5),
	1280:  uint16(3),
	1281:  uint16(1),
	1282:  uint16(anon_sym_COMMA),
	1283:  uint16(5),
	1284:  uint16(1),
	1285:  uint16(sym_comment),
	1286:  uint16(7),
	1287:  uint16(1),
	1288:  uint16(anon_sym_SLASH_SLASH_SLASH),
	1289:  uint16(21),
	1290:  uint16(1),
	1291:  uint16(sym_documentation_comment),
	1292:  uint16(147),
	1293:  uint16(31),
	1295:  uint16(anon_sym_use),
	1296:  uint16(anon_sym_enum),
	1297:  uint16(anon_sym_intEnum),
	1298:  uint16(anon_sym_LBRACE),
	1299:  uint16(anon_sym_RBRACE),
	1300:  uint16(anon_sym_list),
	1301:  uint16(anon_sym_map),
	1302:  uint16(anon_sym_set),
	1303:  uint16(anon_sym_structure),
	1304:  uint16(anon_sym_union),
	1305:  uint16(anon_sym_service),
	1306:  uint16(anon_sym_operation),
	1307:  uint16(anon_sym_resource),
	1308:  uint16(anon_sym_AT),
	1309:  uint16(anon_sym_apply),
	1310:  uint16(anon_sym_with),
	1311:  uint16(anon_sym_for),
	1312:  uint16(anon_sym_blob),
	1313:  uint16(anon_sym_boolean),
	1314:  uint16(anon_sym_byte),
	1315:  uint16(anon_sym_document),
	1316:  uint16(anon_sym_double),
	1317:  uint16(anon_sym_float),
	1318:  uint16(anon_sym_integer),
	1319:  uint16(anon_sym_long),
	1320:  uint16(anon_sym_short),
	1321:  uint16(anon_sym_string),
	1322:  uint16(anon_sym_timestamp),
	1323:  uint16(anon_sym_bigInteger),
	1324:  uint16(anon_sym_bigDecimal),
	1325:  uint16(5),
	1326:  uint16(3),
	1327:  uint16(1),
	1328:  uint16(anon_sym_COMMA),
	1329:  uint16(5),
	1330:  uint16(1),
	1331:  uint16(sym_comment),
	1332:  uint16(7),
	1333:  uint16(1),
	1334:  uint16(anon_sym_SLASH_SLASH_SLASH),
	1335:  uint16(22),
	1336:  uint16(1),
	1337:  uint16(sym_documentation_comment),
	1338:  uint16(149),
	1339:  uint16(31),
	1341:  uint16(anon_sym_use),
	1342:  uint16(anon_sym_enum),
	1343:  uint16(anon_sym_intEnum),
	1344:  uint16(anon_sym_LBRACE),
	1345:  uint16(anon_sym_RBRACE),
	1346:  uint16(anon_sym_list),
	1347:  uint16(anon_sym_map),
	1348:  uint16(anon_sym_set),
	1349:  uint16(anon_sym_structure),
	1350:  uint16(anon_sym_union),
	1351:  uint16(anon_sym_service),
	1352:  uint16(anon_sym_operation),
	1353:  uint16(anon_sym_resource),
	1354:  uint16(anon_sym_AT),
	1355:  uint16(anon_sym_apply),
	1356:  uint16(anon_sym_with),
	1357:  uint16(anon_sym_for),
	1358:  uint16(anon_sym_blob),
	1359:  uint16(anon_sym_boolean),
	1360:  uint16(anon_sym_byte),
	1361:  uint16(anon_sym_document),
	1362:  uint16(anon_sym_double),
	1363:  uint16(anon_sym_float),
	1364:  uint16(anon_sym_integer),
	1365:  uint16(anon_sym_long),
	1366:  uint16(anon_sym_short),
	1367:  uint16(anon_sym_string),
	1368:  uint16(anon_sym_timestamp),
	1369:  uint16(anon_sym_bigInteger),
	1370:  uint16(anon_sym_bigDecimal),
	1371:  uint16(7),
	1372:  uint16(3),
	1373:  uint16(1),
	1374:  uint16(anon_sym_COMMA),
	1375:  uint16(5),
	1376:  uint16(1),
	1377:  uint16(sym_comment),
	1378:  uint16(7),
	1379:  uint16(1),
	1380:  uint16(anon_sym_SLASH_SLASH_SLASH),
	1381:  uint16(153),
	1382:  uint16(1),
	1383:  uint16(anon_sym_DOT),
	1384:  uint16(23),
	1385:  uint16(1),
	1386:  uint16(sym_documentation_comment),
	1387:  uint16(25),
	1388:  uint16(1),
	1389:  uint16(aux_sym_namespace_repeat1),
	1390:  uint16(151),
	1391:  uint16(28),
	1393:  uint16(anon_sym_use),
	1394:  uint16(anon_sym_POUND),
	1395:  uint16(anon_sym_enum),
	1396:  uint16(anon_sym_intEnum),
	1397:  uint16(anon_sym_list),
	1398:  uint16(anon_sym_map),
	1399:  uint16(anon_sym_set),
	1400:  uint16(anon_sym_structure),
	1401:  uint16(anon_sym_union),
	1402:  uint16(anon_sym_service),
	1403:  uint16(anon_sym_operation),
	1404:  uint16(anon_sym_resource),
	1405:  uint16(anon_sym_AT),
	1406:  uint16(anon_sym_apply),
	1407:  uint16(anon_sym_blob),
	1408:  uint16(anon_sym_boolean),
	1409:  uint16(anon_sym_byte),
	1410:  uint16(anon_sym_document),
	1411:  uint16(anon_sym_double),
	1412:  uint16(anon_sym_float),
	1413:  uint16(anon_sym_integer),
	1414:  uint16(anon_sym_long),
	1415:  uint16(anon_sym_short),
	1416:  uint16(anon_sym_string),
	1417:  uint16(anon_sym_timestamp),
	1418:  uint16(anon_sym_bigInteger),
	1419:  uint16(anon_sym_bigDecimal),
	1420:  uint16(6),
	1421:  uint16(3),
	1422:  uint16(1),
	1423:  uint16(anon_sym_COMMA),
	1424:  uint16(5),
	1425:  uint16(1),
	1426:  uint16(sym_comment),
	1427:  uint16(7),
	1428:  uint16(1),
	1429:  uint16(anon_sym_SLASH_SLASH_SLASH),
	1430:  uint16(157),
	1431:  uint16(1),
	1432:  uint16(anon_sym_DOT),
	1433:  uint16(24),
	1434:  uint16(2),
	1435:  uint16(sym_documentation_comment),
	1436:  uint16(aux_sym_namespace_repeat1),
	1437:  uint16(155),
	1438:  uint16(28),
	1440:  uint16(anon_sym_use),
	1441:  uint16(anon_sym_POUND),
	1442:  uint16(anon_sym_enum),
	1443:  uint16(anon_sym_intEnum),
	1444:  uint16(anon_sym_list),
	1445:  uint16(anon_sym_map),
	1446:  uint16(anon_sym_set),
	1447:  uint16(anon_sym_structure),
	1448:  uint16(anon_sym_union),
	1449:  uint16(anon_sym_service),
	1450:  uint16(anon_sym_operation),
	1451:  uint16(anon_sym_resource),
	1452:  uint16(anon_sym_AT),
	1453:  uint16(anon_sym_apply),
	1454:  uint16(anon_sym_blob),
	1455:  uint16(anon_sym_boolean),
	1456:  uint16(anon_sym_byte),
	1457:  uint16(anon_sym_document),
	1458:  uint16(anon_sym_double),
	1459:  uint16(anon_sym_float),
	1460:  uint16(anon_sym_integer),
	1461:  uint16(anon_sym_long),
	1462:  uint16(anon_sym_short),
	1463:  uint16(anon_sym_string),
	1464:  uint16(anon_sym_timestamp),
	1465:  uint16(anon_sym_bigInteger),
	1466:  uint16(anon_sym_bigDecimal),
	1467:  uint16(7),
	1468:  uint16(3),
	1469:  uint16(1),
	1470:  uint16(anon_sym_COMMA),
	1471:  uint16(5),
	1472:  uint16(1),
	1473:  uint16(sym_comment),
	1474:  uint16(7),
	1475:  uint16(1),
	1476:  uint16(anon_sym_SLASH_SLASH_SLASH),
	1477:  uint16(153),
	1478:  uint16(1),
	1479:  uint16(anon_sym_DOT),
	1480:  uint16(24),
	1481:  uint16(1),
	1482:  uint16(aux_sym_namespace_repeat1),
	1483:  uint16(25),
	1484:  uint16(1),
	1485:  uint16(sym_documentation_comment),
	1486:  uint16(160),
	1487:  uint16(28),
	1489:  uint16(anon_sym_use),
	1490:  uint16(anon_sym_POUND),
	1491:  uint16(anon_sym_enum),
	1492:  uint16(anon_sym_intEnum),
	1493:  uint16(anon_sym_list),
	1494:  uint16(anon_sym_map),
	1495:  uint16(anon_sym_set),
	1496:  uint16(anon_sym_structure),
	1497:  uint16(anon_sym_union),
	1498:  uint16(anon_sym_service),
	1499:  uint16(anon_sym_operation),
	1500:  uint16(anon_sym_resource),
	1501:  uint16(anon_sym_AT),
	1502:  uint16(anon_sym_apply),
	1503:  uint16(anon_sym_blob),
	1504:  uint16(anon_sym_boolean),
	1505:  uint16(anon_sym_byte),
	1506:  uint16(anon_sym_document),
	1507:  uint16(anon_sym_double),
	1508:  uint16(anon_sym_float),
	1509:  uint16(anon_sym_integer),
	1510:  uint16(anon_sym_long),
	1511:  uint16(anon_sym_short),
	1512:  uint16(anon_sym_string),
	1513:  uint16(anon_sym_timestamp),
	1514:  uint16(anon_sym_bigInteger),
	1515:  uint16(anon_sym_bigDecimal),
	1516:  uint16(7),
	1517:  uint16(3),
	1518:  uint16(1),
	1519:  uint16(anon_sym_COMMA),
	1520:  uint16(5),
	1521:  uint16(1),
	1522:  uint16(sym_comment),
	1523:  uint16(7),
	1524:  uint16(1),
	1525:  uint16(anon_sym_SLASH_SLASH_SLASH),
	1526:  uint16(164),
	1527:  uint16(1),
	1528:  uint16(anon_sym_AT),
	1529:  uint16(37),
	1530:  uint16(1),
	1531:  uint16(sym_trait_statement),
	1532:  uint16(26),
	1533:  uint16(2),
	1534:  uint16(sym_documentation_comment),
	1535:  uint16(aux_sym_shape_statement_repeat1),
	1536:  uint16(162),
	1537:  uint16(27),
	1538:  uint16(anon_sym_enum),
	1539:  uint16(anon_sym_intEnum),
	1540:  uint16(anon_sym_LBRACE),
	1541:  uint16(anon_sym_RBRACE),
	1542:  uint16(anon_sym_list),
	1543:  uint16(anon_sym_map),
	1544:  uint16(anon_sym_set),
	1545:  uint16(anon_sym_structure),
	1546:  uint16(anon_sym_union),
	1547:  uint16(anon_sym_service),
	1548:  uint16(anon_sym_operation),
	1549:  uint16(anon_sym_resource),
	1550:  uint16(anon_sym_with),
	1551:  uint16(anon_sym_for),
	1552:  uint16(anon_sym_blob),
	1553:  uint16(anon_sym_boolean),
	1554:  uint16(anon_sym_byte),
	1555:  uint16(anon_sym_document),
	1556:  uint16(anon_sym_double),
	1557:  uint16(anon_sym_float),
	1558:  uint16(anon_sym_integer),
	1559:  uint16(anon_sym_long),
	1560:  uint16(anon_sym_short),
	1561:  uint16(anon_sym_string),
	1562:  uint16(anon_sym_timestamp),
	1563:  uint16(anon_sym_bigInteger),
	1564:  uint16(anon_sym_bigDecimal),
	1565:  uint16(5),
	1566:  uint16(3),
	1567:  uint16(1),
	1568:  uint16(anon_sym_COMMA),
	1569:  uint16(5),
	1570:  uint16(1),
	1571:  uint16(sym_comment),
	1572:  uint16(7),
	1573:  uint16(1),
	1574:  uint16(anon_sym_SLASH_SLASH_SLASH),
	1575:  uint16(27),
	1576:  uint16(1),
	1577:  uint16(sym_documentation_comment),
	1578:  uint16(155),
	1579:  uint16(29),
	1581:  uint16(anon_sym_DOT),
	1582:  uint16(anon_sym_use),
	1583:  uint16(anon_sym_POUND),
	1584:  uint16(anon_sym_enum),
	1585:  uint16(anon_sym_intEnum),
	1586:  uint16(anon_sym_list),
	1587:  uint16(anon_sym_map),
	1588:  uint16(anon_sym_set),
	1589:  uint16(anon_sym_structure),
	1590:  uint16(anon_sym_union),
	1591:  uint16(anon_sym_service),
	1592:  uint16(anon_sym_operation),
	1593:  uint16(anon_sym_resource),
	1594:  uint16(anon_sym_AT),
	1595:  uint16(anon_sym_apply),
	1596:  uint16(anon_sym_blob),
	1597:  uint16(anon_sym_boolean),
	1598:  uint16(anon_sym_byte),
	1599:  uint16(anon_sym_document),
	1600:  uint16(anon_sym_double),
	1601:  uint16(anon_sym_float),
	1602:  uint16(anon_sym_integer),
	1603:  uint16(anon_sym_long),
	1604:  uint16(anon_sym_short),
	1605:  uint16(anon_sym_string),
	1606:  uint16(anon_sym_timestamp),
	1607:  uint16(anon_sym_bigInteger),
	1608:  uint16(anon_sym_bigDecimal),
	1609:  uint16(24),
	1610:  uint16(3),
	1611:  uint16(1),
	1612:  uint16(anon_sym_COMMA),
	1613:  uint16(5),
	1614:  uint16(1),
	1615:  uint16(sym_comment),
	1616:  uint16(7),
	1617:  uint16(1),
	1618:  uint16(anon_sym_SLASH_SLASH_SLASH),
	1619:  uint16(167),
	1620:  uint16(1),
	1621:  uint16(anon_sym_LBRACE),
	1622:  uint16(169),
	1623:  uint16(1),
	1624:  uint16(anon_sym_LBRACK),
	1625:  uint16(171),
	1626:  uint16(1),
	1627:  uint16(anon_sym_RBRACK),
	1628:  uint16(175),
	1629:  uint16(1),
	1630:  uint16(sym_null),
	1631:  uint16(177),
	1632:  uint16(1),
	1633:  uint16(anon_sym_DASH),
	1634:  uint16(179),
	1635:  uint16(1),
	1636:  uint16(aux_sym_number_token1),
	1637:  uint16(181),
	1638:  uint16(1),
	1639:  uint16(aux_sym_float_token1),
	1640:  uint16(183),
	1641:  uint16(1),
	1642:  uint16(anon_sym_DQUOTE),
	1643:  uint16(185),
	1644:  uint16(1),
	1645:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1646:  uint16(187),
	1647:  uint16(1),
	1648:  uint16(aux_sym_identifier_token1),
	1649:  uint16(23),
	1650:  uint16(1),
	1651:  uint16(sym__namespace_identifier),
	1652:  uint16(28),
	1653:  uint16(1),
	1654:  uint16(sym_documentation_comment),
	1655:  uint16(31),
	1656:  uint16(1),
	1657:  uint16(aux_sym_node_array_repeat1),
	1658:  uint16(81),
	1659:  uint16(1),
	1660:  uint16(sym_root_shape_id),
	1661:  uint16(104),
	1662:  uint16(1),
	1663:  uint16(sym_node_value),
	1664:  uint16(328),
	1665:  uint16(1),
	1666:  uint16(sym_namespace),
	1667:  uint16(173),
	1668:  uint16(2),
	1669:  uint16(anon_sym_true),
	1670:  uint16(anon_sym_false),
	1671:  uint16(86),
	1672:  uint16(2),
	1673:  uint16(sym_absolute_root_shape_id),
	1674:  uint16(sym_identifier),
	1675:  uint16(91),
	1676:  uint16(2),
	1677:  uint16(sym__string_literal),
	1678:  uint16(sym__multiline_string_literal),
	1679:  uint16(90),
	1680:  uint16(4),
	1681:  uint16(sym_boolean),
	1682:  uint16(sym_number),
	1683:  uint16(sym_float),
	1684:  uint16(sym_string),
	1685:  uint16(96),
	1686:  uint16(4),
	1687:  uint16(sym_shape_id),
	1688:  uint16(sym_node_array),
	1689:  uint16(sym_node_object),
	1690:  uint16(sym_literal),
	1691:  uint16(24),
	1692:  uint16(3),
	1693:  uint16(1),
	1694:  uint16(anon_sym_COMMA),
	1695:  uint16(5),
	1696:  uint16(1),
	1697:  uint16(sym_comment),
	1698:  uint16(7),
	1699:  uint16(1),
	1700:  uint16(anon_sym_SLASH_SLASH_SLASH),
	1701:  uint16(167),
	1702:  uint16(1),
	1703:  uint16(anon_sym_LBRACE),
	1704:  uint16(169),
	1705:  uint16(1),
	1706:  uint16(anon_sym_LBRACK),
	1707:  uint16(175),
	1708:  uint16(1),
	1709:  uint16(sym_null),
	1710:  uint16(177),
	1711:  uint16(1),
	1712:  uint16(anon_sym_DASH),
	1713:  uint16(179),
	1714:  uint16(1),
	1715:  uint16(aux_sym_number_token1),
	1716:  uint16(181),
	1717:  uint16(1),
	1718:  uint16(aux_sym_float_token1),
	1719:  uint16(183),
	1720:  uint16(1),
	1721:  uint16(anon_sym_DQUOTE),
	1722:  uint16(185),
	1723:  uint16(1),
	1724:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1725:  uint16(187),
	1726:  uint16(1),
	1727:  uint16(aux_sym_identifier_token1),
	1728:  uint16(189),
	1729:  uint16(1),
	1730:  uint16(anon_sym_RBRACK),
	1731:  uint16(23),
	1732:  uint16(1),
	1733:  uint16(sym__namespace_identifier),
	1734:  uint16(29),
	1735:  uint16(1),
	1736:  uint16(sym_documentation_comment),
	1737:  uint16(31),
	1738:  uint16(1),
	1739:  uint16(aux_sym_node_array_repeat1),
	1740:  uint16(81),
	1741:  uint16(1),
	1742:  uint16(sym_root_shape_id),
	1743:  uint16(104),
	1744:  uint16(1),
	1745:  uint16(sym_node_value),
	1746:  uint16(328),
	1747:  uint16(1),
	1748:  uint16(sym_namespace),
	1749:  uint16(173),
	1750:  uint16(2),
	1751:  uint16(anon_sym_true),
	1752:  uint16(anon_sym_false),
	1753:  uint16(86),
	1754:  uint16(2),
	1755:  uint16(sym_absolute_root_shape_id),
	1756:  uint16(sym_identifier),
	1757:  uint16(91),
	1758:  uint16(2),
	1759:  uint16(sym__string_literal),
	1760:  uint16(sym__multiline_string_literal),
	1761:  uint16(90),
	1762:  uint16(4),
	1763:  uint16(sym_boolean),
	1764:  uint16(sym_number),
	1765:  uint16(sym_float),
	1766:  uint16(sym_string),
	1767:  uint16(96),
	1768:  uint16(4),
	1769:  uint16(sym_shape_id),
	1770:  uint16(sym_node_array),
	1771:  uint16(sym_node_object),
	1772:  uint16(sym_literal),
	1773:  uint16(5),
	1774:  uint16(3),
	1775:  uint16(1),
	1776:  uint16(anon_sym_COMMA),
	1777:  uint16(5),
	1778:  uint16(1),
	1779:  uint16(sym_comment),
	1780:  uint16(7),
	1781:  uint16(1),
	1782:  uint16(anon_sym_SLASH_SLASH_SLASH),
	1783:  uint16(30),
	1784:  uint16(1),
	1785:  uint16(sym_documentation_comment),
	1786:  uint16(90),
	1787:  uint16(29),
	1789:  uint16(anon_sym_DOT),
	1790:  uint16(anon_sym_use),
	1791:  uint16(anon_sym_POUND),
	1792:  uint16(anon_sym_enum),
	1793:  uint16(anon_sym_intEnum),
	1794:  uint16(anon_sym_list),
	1795:  uint16(anon_sym_map),
	1796:  uint16(anon_sym_set),
	1797:  uint16(anon_sym_structure),
	1798:  uint16(anon_sym_union),
	1799:  uint16(anon_sym_service),
	1800:  uint16(anon_sym_operation),
	1801:  uint16(anon_sym_resource),
	1802:  uint16(anon_sym_AT),
	1803:  uint16(anon_sym_apply),
	1804:  uint16(anon_sym_blob),
	1805:  uint16(anon_sym_boolean),
	1806:  uint16(anon_sym_byte),
	1807:  uint16(anon_sym_document),
	1808:  uint16(anon_sym_double),
	1809:  uint16(anon_sym_float),
	1810:  uint16(anon_sym_integer),
	1811:  uint16(anon_sym_long),
	1812:  uint16(anon_sym_short),
	1813:  uint16(anon_sym_string),
	1814:  uint16(anon_sym_timestamp),
	1815:  uint16(anon_sym_bigInteger),
	1816:  uint16(anon_sym_bigDecimal),
	1817:  uint16(23),
	1818:  uint16(3),
	1819:  uint16(1),
	1820:  uint16(anon_sym_COMMA),
	1821:  uint16(5),
	1822:  uint16(1),
	1823:  uint16(sym_comment),
	1824:  uint16(7),
	1825:  uint16(1),
	1826:  uint16(anon_sym_SLASH_SLASH_SLASH),
	1827:  uint16(191),
	1828:  uint16(1),
	1829:  uint16(anon_sym_LBRACE),
	1830:  uint16(194),
	1831:  uint16(1),
	1832:  uint16(anon_sym_LBRACK),
	1833:  uint16(197),
	1834:  uint16(1),
	1835:  uint16(anon_sym_RBRACK),
	1836:  uint16(202),
	1837:  uint16(1),
	1838:  uint16(sym_null),
	1839:  uint16(205),
	1840:  uint16(1),
	1841:  uint16(anon_sym_DASH),
	1842:  uint16(208),
	1843:  uint16(1),
	1844:  uint16(aux_sym_number_token1),
	1845:  uint16(211),
	1846:  uint16(1),
	1847:  uint16(aux_sym_float_token1),
	1848:  uint16(214),
	1849:  uint16(1),
	1850:  uint16(anon_sym_DQUOTE),
	1851:  uint16(217),
	1852:  uint16(1),
	1853:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1854:  uint16(220),
	1855:  uint16(1),
	1856:  uint16(aux_sym_identifier_token1),
	1857:  uint16(23),
	1858:  uint16(1),
	1859:  uint16(sym__namespace_identifier),
	1860:  uint16(81),
	1861:  uint16(1),
	1862:  uint16(sym_root_shape_id),
	1863:  uint16(104),
	1864:  uint16(1),
	1865:  uint16(sym_node_value),
	1866:  uint16(328),
	1867:  uint16(1),
	1868:  uint16(sym_namespace),
	1869:  uint16(199),
	1870:  uint16(2),
	1871:  uint16(anon_sym_true),
	1872:  uint16(anon_sym_false),
	1873:  uint16(31),
	1874:  uint16(2),
	1875:  uint16(sym_documentation_comment),
	1876:  uint16(aux_sym_node_array_repeat1),
	1877:  uint16(86),
	1878:  uint16(2),
	1879:  uint16(sym_absolute_root_shape_id),
	1880:  uint16(sym_identifier),
	1881:  uint16(91),
	1882:  uint16(2),
	1883:  uint16(sym__string_literal),
	1884:  uint16(sym__multiline_string_literal),
	1885:  uint16(90),
	1886:  uint16(4),
	1887:  uint16(sym_boolean),
	1888:  uint16(sym_number),
	1889:  uint16(sym_float),
	1890:  uint16(sym_string),
	1891:  uint16(96),
	1892:  uint16(4),
	1893:  uint16(sym_shape_id),
	1894:  uint16(sym_node_array),
	1895:  uint16(sym_node_object),
	1896:  uint16(sym_literal),
	1897:  uint16(24),
	1898:  uint16(3),
	1899:  uint16(1),
	1900:  uint16(anon_sym_COMMA),
	1901:  uint16(5),
	1902:  uint16(1),
	1903:  uint16(sym_comment),
	1904:  uint16(7),
	1905:  uint16(1),
	1906:  uint16(anon_sym_SLASH_SLASH_SLASH),
	1907:  uint16(167),
	1908:  uint16(1),
	1909:  uint16(anon_sym_LBRACE),
	1910:  uint16(169),
	1911:  uint16(1),
	1912:  uint16(anon_sym_LBRACK),
	1913:  uint16(175),
	1914:  uint16(1),
	1915:  uint16(sym_null),
	1916:  uint16(177),
	1917:  uint16(1),
	1918:  uint16(anon_sym_DASH),
	1919:  uint16(179),
	1920:  uint16(1),
	1921:  uint16(aux_sym_number_token1),
	1922:  uint16(181),
	1923:  uint16(1),
	1924:  uint16(aux_sym_float_token1),
	1925:  uint16(183),
	1926:  uint16(1),
	1927:  uint16(anon_sym_DQUOTE),
	1928:  uint16(185),
	1929:  uint16(1),
	1930:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1931:  uint16(187),
	1932:  uint16(1),
	1933:  uint16(aux_sym_identifier_token1),
	1934:  uint16(223),
	1935:  uint16(1),
	1936:  uint16(anon_sym_RBRACK),
	1937:  uint16(23),
	1938:  uint16(1),
	1939:  uint16(sym__namespace_identifier),
	1940:  uint16(31),
	1941:  uint16(1),
	1942:  uint16(aux_sym_node_array_repeat1),
	1943:  uint16(32),
	1944:  uint16(1),
	1945:  uint16(sym_documentation_comment),
	1946:  uint16(81),
	1947:  uint16(1),
	1948:  uint16(sym_root_shape_id),
	1949:  uint16(104),
	1950:  uint16(1),
	1951:  uint16(sym_node_value),
	1952:  uint16(328),
	1953:  uint16(1),
	1954:  uint16(sym_namespace),
	1955:  uint16(173),
	1956:  uint16(2),
	1957:  uint16(anon_sym_true),
	1958:  uint16(anon_sym_false),
	1959:  uint16(86),
	1960:  uint16(2),
	1961:  uint16(sym_absolute_root_shape_id),
	1962:  uint16(sym_identifier),
	1963:  uint16(91),
	1964:  uint16(2),
	1965:  uint16(sym__string_literal),
	1966:  uint16(sym__multiline_string_literal),
	1967:  uint16(90),
	1968:  uint16(4),
	1969:  uint16(sym_boolean),
	1970:  uint16(sym_number),
	1971:  uint16(sym_float),
	1972:  uint16(sym_string),
	1973:  uint16(96),
	1974:  uint16(4),
	1975:  uint16(sym_shape_id),
	1976:  uint16(sym_node_array),
	1977:  uint16(sym_node_object),
	1978:  uint16(sym_literal),
	1979:  uint16(7),
	1980:  uint16(3),
	1981:  uint16(1),
	1982:  uint16(anon_sym_COMMA),
	1983:  uint16(5),
	1984:  uint16(1),
	1985:  uint16(sym_comment),
	1986:  uint16(7),
	1987:  uint16(1),
	1988:  uint16(anon_sym_SLASH_SLASH_SLASH),
	1989:  uint16(227),
	1990:  uint16(1),
	1991:  uint16(anon_sym_with),
	1992:  uint16(33),
	1993:  uint16(1),
	1994:  uint16(sym_documentation_comment),
	1995:  uint16(42),
	1996:  uint16(1),
	1997:  uint16(sym_mixins),
	1998:  uint16(225),
	1999:  uint16(27),
	2001:  uint16(anon_sym_use),
	2002:  uint16(anon_sym_enum),
	2003:  uint16(anon_sym_intEnum),
	2004:  uint16(anon_sym_list),
	2005:  uint16(anon_sym_map),
	2006:  uint16(anon_sym_set),
	2007:  uint16(anon_sym_structure),
	2008:  uint16(anon_sym_union),
	2009:  uint16(anon_sym_service),
	2010:  uint16(anon_sym_operation),
	2011:  uint16(anon_sym_resource),
	2012:  uint16(anon_sym_AT),
	2013:  uint16(anon_sym_apply),
	2014:  uint16(anon_sym_blob),
	2015:  uint16(anon_sym_boolean),
	2016:  uint16(anon_sym_byte),
	2017:  uint16(anon_sym_document),
	2018:  uint16(anon_sym_double),
	2019:  uint16(anon_sym_float),
	2020:  uint16(anon_sym_integer),
	2021:  uint16(anon_sym_long),
	2022:  uint16(anon_sym_short),
	2023:  uint16(anon_sym_string),
	2024:  uint16(anon_sym_timestamp),
	2025:  uint16(anon_sym_bigInteger),
	2026:  uint16(anon_sym_bigDecimal),
	2027:  uint16(24),
	2028:  uint16(3),
	2029:  uint16(1),
	2030:  uint16(anon_sym_COMMA),
	2031:  uint16(5),
	2032:  uint16(1),
	2033:  uint16(sym_comment),
	2034:  uint16(7),
	2035:  uint16(1),
	2036:  uint16(anon_sym_SLASH_SLASH_SLASH),
	2037:  uint16(167),
	2038:  uint16(1),
	2039:  uint16(anon_sym_LBRACE),
	2040:  uint16(169),
	2041:  uint16(1),
	2042:  uint16(anon_sym_LBRACK),
	2043:  uint16(175),
	2044:  uint16(1),
	2045:  uint16(sym_null),
	2046:  uint16(177),
	2047:  uint16(1),
	2048:  uint16(anon_sym_DASH),
	2049:  uint16(179),
	2050:  uint16(1),
	2051:  uint16(aux_sym_number_token1),
	2052:  uint16(181),
	2053:  uint16(1),
	2054:  uint16(aux_sym_float_token1),
	2055:  uint16(183),
	2056:  uint16(1),
	2057:  uint16(anon_sym_DQUOTE),
	2058:  uint16(185),
	2059:  uint16(1),
	2060:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2061:  uint16(187),
	2062:  uint16(1),
	2063:  uint16(aux_sym_identifier_token1),
	2064:  uint16(229),
	2065:  uint16(1),
	2066:  uint16(anon_sym_RBRACK),
	2067:  uint16(23),
	2068:  uint16(1),
	2069:  uint16(sym__namespace_identifier),
	2070:  uint16(28),
	2071:  uint16(1),
	2072:  uint16(aux_sym_node_array_repeat1),
	2073:  uint16(34),
	2074:  uint16(1),
	2075:  uint16(sym_documentation_comment),
	2076:  uint16(81),
	2077:  uint16(1),
	2078:  uint16(sym_root_shape_id),
	2079:  uint16(104),
	2080:  uint16(1),
	2081:  uint16(sym_node_value),
	2082:  uint16(328),
	2083:  uint16(1),
	2084:  uint16(sym_namespace),
	2085:  uint16(173),
	2086:  uint16(2),
	2087:  uint16(anon_sym_true),
	2088:  uint16(anon_sym_false),
	2089:  uint16(86),
	2090:  uint16(2),
	2091:  uint16(sym_absolute_root_shape_id),
	2092:  uint16(sym_identifier),
	2093:  uint16(91),
	2094:  uint16(2),
	2095:  uint16(sym__string_literal),
	2096:  uint16(sym__multiline_string_literal),
	2097:  uint16(90),
	2098:  uint16(4),
	2099:  uint16(sym_boolean),
	2100:  uint16(sym_number),
	2101:  uint16(sym_float),
	2102:  uint16(sym_string),
	2103:  uint16(96),
	2104:  uint16(4),
	2105:  uint16(sym_shape_id),
	2106:  uint16(sym_node_array),
	2107:  uint16(sym_node_object),
	2108:  uint16(sym_literal),
	2109:  uint16(24),
	2110:  uint16(3),
	2111:  uint16(1),
	2112:  uint16(anon_sym_COMMA),
	2113:  uint16(5),
	2114:  uint16(1),
	2115:  uint16(sym_comment),
	2116:  uint16(7),
	2117:  uint16(1),
	2118:  uint16(anon_sym_SLASH_SLASH_SLASH),
	2119:  uint16(167),
	2120:  uint16(1),
	2121:  uint16(anon_sym_LBRACE),
	2122:  uint16(169),
	2123:  uint16(1),
	2124:  uint16(anon_sym_LBRACK),
	2125:  uint16(175),
	2126:  uint16(1),
	2127:  uint16(sym_null),
	2128:  uint16(177),
	2129:  uint16(1),
	2130:  uint16(anon_sym_DASH),
	2131:  uint16(179),
	2132:  uint16(1),
	2133:  uint16(aux_sym_number_token1),
	2134:  uint16(181),
	2135:  uint16(1),
	2136:  uint16(aux_sym_float_token1),
	2137:  uint16(183),
	2138:  uint16(1),
	2139:  uint16(anon_sym_DQUOTE),
	2140:  uint16(185),
	2141:  uint16(1),
	2142:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2143:  uint16(187),
	2144:  uint16(1),
	2145:  uint16(aux_sym_identifier_token1),
	2146:  uint16(231),
	2147:  uint16(1),
	2148:  uint16(anon_sym_RBRACK),
	2149:  uint16(23),
	2150:  uint16(1),
	2151:  uint16(sym__namespace_identifier),
	2152:  uint16(29),
	2153:  uint16(1),
	2154:  uint16(aux_sym_node_array_repeat1),
	2155:  uint16(35),
	2156:  uint16(1),
	2157:  uint16(sym_documentation_comment),
	2158:  uint16(81),
	2159:  uint16(1),
	2160:  uint16(sym_root_shape_id),
	2161:  uint16(104),
	2162:  uint16(1),
	2163:  uint16(sym_node_value),
	2164:  uint16(328),
	2165:  uint16(1),
	2166:  uint16(sym_namespace),
	2167:  uint16(173),
	2168:  uint16(2),
	2169:  uint16(anon_sym_true),
	2170:  uint16(anon_sym_false),
	2171:  uint16(86),
	2172:  uint16(2),
	2173:  uint16(sym_absolute_root_shape_id),
	2174:  uint16(sym_identifier),
	2175:  uint16(91),
	2176:  uint16(2),
	2177:  uint16(sym__string_literal),
	2178:  uint16(sym__multiline_string_literal),
	2179:  uint16(90),
	2180:  uint16(4),
	2181:  uint16(sym_boolean),
	2182:  uint16(sym_number),
	2183:  uint16(sym_float),
	2184:  uint16(sym_string),
	2185:  uint16(96),
	2186:  uint16(4),
	2187:  uint16(sym_shape_id),
	2188:  uint16(sym_node_array),
	2189:  uint16(sym_node_object),
	2190:  uint16(sym_literal),
	2191:  uint16(24),
	2192:  uint16(3),
	2193:  uint16(1),
	2194:  uint16(anon_sym_COMMA),
	2195:  uint16(5),
	2196:  uint16(1),
	2197:  uint16(sym_comment),
	2198:  uint16(7),
	2199:  uint16(1),
	2200:  uint16(anon_sym_SLASH_SLASH_SLASH),
	2201:  uint16(167),
	2202:  uint16(1),
	2203:  uint16(anon_sym_LBRACE),
	2204:  uint16(169),
	2205:  uint16(1),
	2206:  uint16(anon_sym_LBRACK),
	2207:  uint16(175),
	2208:  uint16(1),
	2209:  uint16(sym_null),
	2210:  uint16(177),
	2211:  uint16(1),
	2212:  uint16(anon_sym_DASH),
	2213:  uint16(179),
	2214:  uint16(1),
	2215:  uint16(aux_sym_number_token1),
	2216:  uint16(181),
	2217:  uint16(1),
	2218:  uint16(aux_sym_float_token1),
	2219:  uint16(183),
	2220:  uint16(1),
	2221:  uint16(anon_sym_DQUOTE),
	2222:  uint16(185),
	2223:  uint16(1),
	2224:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2225:  uint16(187),
	2226:  uint16(1),
	2227:  uint16(aux_sym_identifier_token1),
	2228:  uint16(233),
	2229:  uint16(1),
	2230:  uint16(anon_sym_RBRACK),
	2231:  uint16(23),
	2232:  uint16(1),
	2233:  uint16(sym__namespace_identifier),
	2234:  uint16(32),
	2235:  uint16(1),
	2236:  uint16(aux_sym_node_array_repeat1),
	2237:  uint16(36),
	2238:  uint16(1),
	2239:  uint16(sym_documentation_comment),
	2240:  uint16(81),
	2241:  uint16(1),
	2242:  uint16(sym_root_shape_id),
	2243:  uint16(104),
	2244:  uint16(1),
	2245:  uint16(sym_node_value),
	2246:  uint16(328),
	2247:  uint16(1),
	2248:  uint16(sym_namespace),
	2249:  uint16(173),
	2250:  uint16(2),
	2251:  uint16(anon_sym_true),
	2252:  uint16(anon_sym_false),
	2253:  uint16(86),
	2254:  uint16(2),
	2255:  uint16(sym_absolute_root_shape_id),
	2256:  uint16(sym_identifier),
	2257:  uint16(91),
	2258:  uint16(2),
	2259:  uint16(sym__string_literal),
	2260:  uint16(sym__multiline_string_literal),
	2261:  uint16(90),
	2262:  uint16(4),
	2263:  uint16(sym_boolean),
	2264:  uint16(sym_number),
	2265:  uint16(sym_float),
	2266:  uint16(sym_string),
	2267:  uint16(96),
	2268:  uint16(4),
	2269:  uint16(sym_shape_id),
	2270:  uint16(sym_node_array),
	2271:  uint16(sym_node_object),
	2272:  uint16(sym_literal),
	2273:  uint16(5),
	2274:  uint16(3),
	2275:  uint16(1),
	2276:  uint16(anon_sym_COMMA),
	2277:  uint16(5),
	2278:  uint16(1),
	2279:  uint16(sym_comment),
	2280:  uint16(7),
	2281:  uint16(1),
	2282:  uint16(anon_sym_SLASH_SLASH_SLASH),
	2283:  uint16(37),
	2284:  uint16(1),
	2285:  uint16(sym_documentation_comment),
	2286:  uint16(235),
	2287:  uint16(28),
	2288:  uint16(anon_sym_enum),
	2289:  uint16(anon_sym_intEnum),
	2290:  uint16(anon_sym_LBRACE),
	2291:  uint16(anon_sym_RBRACE),
	2292:  uint16(anon_sym_list),
	2293:  uint16(anon_sym_map),
	2294:  uint16(anon_sym_set),
	2295:  uint16(anon_sym_structure),
	2296:  uint16(anon_sym_union),
	2297:  uint16(anon_sym_service),
	2298:  uint16(anon_sym_operation),
	2299:  uint16(anon_sym_resource),
	2300:  uint16(anon_sym_AT),
	2301:  uint16(anon_sym_with),
	2302:  uint16(anon_sym_for),
	2303:  uint16(anon_sym_blob),
	2304:  uint16(anon_sym_boolean),
	2305:  uint16(anon_sym_byte),
	2306:  uint16(anon_sym_document),
	2307:  uint16(anon_sym_double),
	2308:  uint16(anon_sym_float),
	2309:  uint16(anon_sym_integer),
	2310:  uint16(anon_sym_long),
	2311:  uint16(anon_sym_short),
	2312:  uint16(anon_sym_string),
	2313:  uint16(anon_sym_timestamp),
	2314:  uint16(anon_sym_bigInteger),
	2315:  uint16(anon_sym_bigDecimal),
	2316:  uint16(5),
	2317:  uint16(3),
	2318:  uint16(1),
	2319:  uint16(anon_sym_COMMA),
	2320:  uint16(5),
	2321:  uint16(1),
	2322:  uint16(sym_comment),
	2323:  uint16(7),
	2324:  uint16(1),
	2325:  uint16(anon_sym_SLASH_SLASH_SLASH),
	2326:  uint16(38),
	2327:  uint16(1),
	2328:  uint16(sym_documentation_comment),
	2329:  uint16(237),
	2330:  uint16(28),
	2332:  uint16(anon_sym_use),
	2333:  uint16(anon_sym_enum),
	2334:  uint16(anon_sym_intEnum),
	2335:  uint16(anon_sym_LBRACE),
	2336:  uint16(anon_sym_list),
	2337:  uint16(anon_sym_map),
	2338:  uint16(anon_sym_set),
	2339:  uint16(anon_sym_structure),
	2340:  uint16(anon_sym_union),
	2341:  uint16(anon_sym_service),
	2342:  uint16(anon_sym_operation),
	2343:  uint16(anon_sym_resource),
	2344:  uint16(anon_sym_AT),
	2345:  uint16(anon_sym_apply),
	2346:  uint16(anon_sym_blob),
	2347:  uint16(anon_sym_boolean),
	2348:  uint16(anon_sym_byte),
	2349:  uint16(anon_sym_document),
	2350:  uint16(anon_sym_double),
	2351:  uint16(anon_sym_float),
	2352:  uint16(anon_sym_integer),
	2353:  uint16(anon_sym_long),
	2354:  uint16(anon_sym_short),
	2355:  uint16(anon_sym_string),
	2356:  uint16(anon_sym_timestamp),
	2357:  uint16(anon_sym_bigInteger),
	2358:  uint16(anon_sym_bigDecimal),
	2359:  uint16(22),
	2360:  uint16(3),
	2361:  uint16(1),
	2362:  uint16(anon_sym_COMMA),
	2363:  uint16(5),
	2364:  uint16(1),
	2365:  uint16(sym_comment),
	2366:  uint16(7),
	2367:  uint16(1),
	2368:  uint16(anon_sym_SLASH_SLASH_SLASH),
	2369:  uint16(239),
	2370:  uint16(1),
	2371:  uint16(anon_sym_LBRACE),
	2372:  uint16(241),
	2373:  uint16(1),
	2374:  uint16(anon_sym_LBRACK),
	2375:  uint16(245),
	2376:  uint16(1),
	2377:  uint16(sym_null),
	2378:  uint16(247),
	2379:  uint16(1),
	2380:  uint16(anon_sym_DASH),
	2381:  uint16(249),
	2382:  uint16(1),
	2383:  uint16(aux_sym_number_token1),
	2384:  uint16(251),
	2385:  uint16(1),
	2386:  uint16(aux_sym_float_token1),
	2387:  uint16(253),
	2388:  uint16(1),
	2389:  uint16(anon_sym_DQUOTE),
	2390:  uint16(255),
	2391:  uint16(1),
	2392:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2393:  uint16(257),
	2394:  uint16(1),
	2395:  uint16(aux_sym_identifier_token1),
	2396:  uint16(23),
	2397:  uint16(1),
	2398:  uint16(sym__namespace_identifier),
	2399:  uint16(39),
	2400:  uint16(1),
	2401:  uint16(sym_documentation_comment),
	2402:  uint16(120),
	2403:  uint16(1),
	2404:  uint16(sym_root_shape_id),
	2405:  uint16(232),
	2406:  uint16(1),
	2407:  uint16(sym_node_value),
	2408:  uint16(345),
	2409:  uint16(1),
	2410:  uint16(sym_namespace),
	2411:  uint16(243),
	2412:  uint16(2),
	2413:  uint16(anon_sym_true),
	2414:  uint16(anon_sym_false),
	2415:  uint16(130),
	2416:  uint16(2),
	2417:  uint16(sym_absolute_root_shape_id),
	2418:  uint16(sym_identifier),
	2419:  uint16(169),
	2420:  uint16(2),
	2421:  uint16(sym__string_literal),
	2422:  uint16(sym__multiline_string_literal),
	2423:  uint16(174),
	2424:  uint16(4),
	2425:  uint16(sym_boolean),
	2426:  uint16(sym_number),
	2427:  uint16(sym_float),
	2428:  uint16(sym_string),
	2429:  uint16(184),
	2430:  uint16(4),
	2431:  uint16(sym_shape_id),
	2432:  uint16(sym_node_array),
	2433:  uint16(sym_node_object),
	2434:  uint16(sym_literal),
	2435:  uint16(5),
	2436:  uint16(3),
	2437:  uint16(1),
	2438:  uint16(anon_sym_COMMA),
	2439:  uint16(5),
	2440:  uint16(1),
	2441:  uint16(sym_comment),
	2442:  uint16(7),
	2443:  uint16(1),
	2444:  uint16(anon_sym_SLASH_SLASH_SLASH),
	2445:  uint16(40),
	2446:  uint16(1),
	2447:  uint16(sym_documentation_comment),
	2448:  uint16(259),
	2449:  uint16(27),
	2451:  uint16(anon_sym_use),
	2452:  uint16(anon_sym_enum),
	2453:  uint16(anon_sym_intEnum),
	2454:  uint16(anon_sym_list),
	2455:  uint16(anon_sym_map),
	2456:  uint16(anon_sym_set),
	2457:  uint16(anon_sym_structure),
	2458:  uint16(anon_sym_union),
	2459:  uint16(anon_sym_service),
	2460:  uint16(anon_sym_operation),
	2461:  uint16(anon_sym_resource),
	2462:  uint16(anon_sym_AT),
	2463:  uint16(anon_sym_apply),
	2464:  uint16(anon_sym_blob),
	2465:  uint16(anon_sym_boolean),
	2466:  uint16(anon_sym_byte),
	2467:  uint16(anon_sym_document),
	2468:  uint16(anon_sym_double),
	2469:  uint16(anon_sym_float),
	2470:  uint16(anon_sym_integer),
	2471:  uint16(anon_sym_long),
	2472:  uint16(anon_sym_short),
	2473:  uint16(anon_sym_string),
	2474:  uint16(anon_sym_timestamp),
	2475:  uint16(anon_sym_bigInteger),
	2476:  uint16(anon_sym_bigDecimal),
	2477:  uint16(5),
	2478:  uint16(3),
	2479:  uint16(1),
	2480:  uint16(anon_sym_COMMA),
	2481:  uint16(5),
	2482:  uint16(1),
	2483:  uint16(sym_comment),
	2484:  uint16(7),
	2485:  uint16(1),
	2486:  uint16(anon_sym_SLASH_SLASH_SLASH),
	2487:  uint16(41),
	2488:  uint16(1),
	2489:  uint16(sym_documentation_comment),
	2490:  uint16(261),
	2491:  uint16(27),
	2493:  uint16(anon_sym_use),
	2494:  uint16(anon_sym_enum),
	2495:  uint16(anon_sym_intEnum),
	2496:  uint16(anon_sym_list),
	2497:  uint16(anon_sym_map),
	2498:  uint16(anon_sym_set),
	2499:  uint16(anon_sym_structure),
	2500:  uint16(anon_sym_union),
	2501:  uint16(anon_sym_service),
	2502:  uint16(anon_sym_operation),
	2503:  uint16(anon_sym_resource),
	2504:  uint16(anon_sym_AT),
	2505:  uint16(anon_sym_apply),
	2506:  uint16(anon_sym_blob),
	2507:  uint16(anon_sym_boolean),
	2508:  uint16(anon_sym_byte),
	2509:  uint16(anon_sym_document),
	2510:  uint16(anon_sym_double),
	2511:  uint16(anon_sym_float),
	2512:  uint16(anon_sym_integer),
	2513:  uint16(anon_sym_long),
	2514:  uint16(anon_sym_short),
	2515:  uint16(anon_sym_string),
	2516:  uint16(anon_sym_timestamp),
	2517:  uint16(anon_sym_bigInteger),
	2518:  uint16(anon_sym_bigDecimal),
	2519:  uint16(5),
	2520:  uint16(3),
	2521:  uint16(1),
	2522:  uint16(anon_sym_COMMA),
	2523:  uint16(5),
	2524:  uint16(1),
	2525:  uint16(sym_comment),
	2526:  uint16(7),
	2527:  uint16(1),
	2528:  uint16(anon_sym_SLASH_SLASH_SLASH),
	2529:  uint16(42),
	2530:  uint16(1),
	2531:  uint16(sym_documentation_comment),
	2532:  uint16(263),
	2533:  uint16(27),
	2535:  uint16(anon_sym_use),
	2536:  uint16(anon_sym_enum),
	2537:  uint16(anon_sym_intEnum),
	2538:  uint16(anon_sym_list),
	2539:  uint16(anon_sym_map),
	2540:  uint16(anon_sym_set),
	2541:  uint16(anon_sym_structure),
	2542:  uint16(anon_sym_union),
	2543:  uint16(anon_sym_service),
	2544:  uint16(anon_sym_operation),
	2545:  uint16(anon_sym_resource),
	2546:  uint16(anon_sym_AT),
	2547:  uint16(anon_sym_apply),
	2548:  uint16(anon_sym_blob),
	2549:  uint16(anon_sym_boolean),
	2550:  uint16(anon_sym_byte),
	2551:  uint16(anon_sym_document),
	2552:  uint16(anon_sym_double),
	2553:  uint16(anon_sym_float),
	2554:  uint16(anon_sym_integer),
	2555:  uint16(anon_sym_long),
	2556:  uint16(anon_sym_short),
	2557:  uint16(anon_sym_string),
	2558:  uint16(anon_sym_timestamp),
	2559:  uint16(anon_sym_bigInteger),
	2560:  uint16(anon_sym_bigDecimal),
	2561:  uint16(5),
	2562:  uint16(3),
	2563:  uint16(1),
	2564:  uint16(anon_sym_COMMA),
	2565:  uint16(5),
	2566:  uint16(1),
	2567:  uint16(sym_comment),
	2568:  uint16(7),
	2569:  uint16(1),
	2570:  uint16(anon_sym_SLASH_SLASH_SLASH),
	2571:  uint16(43),
	2572:  uint16(1),
	2573:  uint16(sym_documentation_comment),
	2574:  uint16(265),
	2575:  uint16(27),
	2577:  uint16(anon_sym_use),
	2578:  uint16(anon_sym_enum),
	2579:  uint16(anon_sym_intEnum),
	2580:  uint16(anon_sym_list),
	2581:  uint16(anon_sym_map),
	2582:  uint16(anon_sym_set),
	2583:  uint16(anon_sym_structure),
	2584:  uint16(anon_sym_union),
	2585:  uint16(anon_sym_service),
	2586:  uint16(anon_sym_operation),
	2587:  uint16(anon_sym_resource),
	2588:  uint16(anon_sym_AT),
	2589:  uint16(anon_sym_apply),
	2590:  uint16(anon_sym_blob),
	2591:  uint16(anon_sym_boolean),
	2592:  uint16(anon_sym_byte),
	2593:  uint16(anon_sym_document),
	2594:  uint16(anon_sym_double),
	2595:  uint16(anon_sym_float),
	2596:  uint16(anon_sym_integer),
	2597:  uint16(anon_sym_long),
	2598:  uint16(anon_sym_short),
	2599:  uint16(anon_sym_string),
	2600:  uint16(anon_sym_timestamp),
	2601:  uint16(anon_sym_bigInteger),
	2602:  uint16(anon_sym_bigDecimal),
	2603:  uint16(22),
	2604:  uint16(3),
	2605:  uint16(1),
	2606:  uint16(anon_sym_COMMA),
	2607:  uint16(5),
	2608:  uint16(1),
	2609:  uint16(sym_comment),
	2610:  uint16(7),
	2611:  uint16(1),
	2612:  uint16(anon_sym_SLASH_SLASH_SLASH),
	2613:  uint16(239),
	2614:  uint16(1),
	2615:  uint16(anon_sym_LBRACE),
	2616:  uint16(241),
	2617:  uint16(1),
	2618:  uint16(anon_sym_LBRACK),
	2619:  uint16(245),
	2620:  uint16(1),
	2621:  uint16(sym_null),
	2622:  uint16(247),
	2623:  uint16(1),
	2624:  uint16(anon_sym_DASH),
	2625:  uint16(249),
	2626:  uint16(1),
	2627:  uint16(aux_sym_number_token1),
	2628:  uint16(251),
	2629:  uint16(1),
	2630:  uint16(aux_sym_float_token1),
	2631:  uint16(253),
	2632:  uint16(1),
	2633:  uint16(anon_sym_DQUOTE),
	2634:  uint16(255),
	2635:  uint16(1),
	2636:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2637:  uint16(257),
	2638:  uint16(1),
	2639:  uint16(aux_sym_identifier_token1),
	2640:  uint16(23),
	2641:  uint16(1),
	2642:  uint16(sym__namespace_identifier),
	2643:  uint16(44),
	2644:  uint16(1),
	2645:  uint16(sym_documentation_comment),
	2646:  uint16(143),
	2647:  uint16(1),
	2648:  uint16(sym_root_shape_id),
	2649:  uint16(232),
	2650:  uint16(1),
	2651:  uint16(sym_node_value),
	2652:  uint16(345),
	2653:  uint16(1),
	2654:  uint16(sym_namespace),
	2655:  uint16(243),
	2656:  uint16(2),
	2657:  uint16(anon_sym_true),
	2658:  uint16(anon_sym_false),
	2659:  uint16(130),
	2660:  uint16(2),
	2661:  uint16(sym_absolute_root_shape_id),
	2662:  uint16(sym_identifier),
	2663:  uint16(169),
	2664:  uint16(2),
	2665:  uint16(sym__string_literal),
	2666:  uint16(sym__multiline_string_literal),
	2667:  uint16(174),
	2668:  uint16(4),
	2669:  uint16(sym_boolean),
	2670:  uint16(sym_number),
	2671:  uint16(sym_float),
	2672:  uint16(sym_string),
	2673:  uint16(184),
	2674:  uint16(4),
	2675:  uint16(sym_shape_id),
	2676:  uint16(sym_node_array),
	2677:  uint16(sym_node_object),
	2678:  uint16(sym_literal),
	2679:  uint16(5),
	2680:  uint16(3),
	2681:  uint16(1),
	2682:  uint16(anon_sym_COMMA),
	2683:  uint16(5),
	2684:  uint16(1),
	2685:  uint16(sym_comment),
	2686:  uint16(7),
	2687:  uint16(1),
	2688:  uint16(anon_sym_SLASH_SLASH_SLASH),
	2689:  uint16(45),
	2690:  uint16(1),
	2691:  uint16(sym_documentation_comment),
	2692:  uint16(267),
	2693:  uint16(27),
	2695:  uint16(anon_sym_use),
	2696:  uint16(anon_sym_enum),
	2697:  uint16(anon_sym_intEnum),
	2698:  uint16(anon_sym_list),
	2699:  uint16(anon_sym_map),
	2700:  uint16(anon_sym_set),
	2701:  uint16(anon_sym_structure),
	2702:  uint16(anon_sym_union),
	2703:  uint16(anon_sym_service),
	2704:  uint16(anon_sym_operation),
	2705:  uint16(anon_sym_resource),
	2706:  uint16(anon_sym_AT),
	2707:  uint16(anon_sym_apply),
	2708:  uint16(anon_sym_blob),
	2709:  uint16(anon_sym_boolean),
	2710:  uint16(anon_sym_byte),
	2711:  uint16(anon_sym_document),
	2712:  uint16(anon_sym_double),
	2713:  uint16(anon_sym_float),
	2714:  uint16(anon_sym_integer),
	2715:  uint16(anon_sym_long),
	2716:  uint16(anon_sym_short),
	2717:  uint16(anon_sym_string),
	2718:  uint16(anon_sym_timestamp),
	2719:  uint16(anon_sym_bigInteger),
	2720:  uint16(anon_sym_bigDecimal),
	2721:  uint16(5),
	2722:  uint16(3),
	2723:  uint16(1),
	2724:  uint16(anon_sym_COMMA),
	2725:  uint16(5),
	2726:  uint16(1),
	2727:  uint16(sym_comment),
	2728:  uint16(7),
	2729:  uint16(1),
	2730:  uint16(anon_sym_SLASH_SLASH_SLASH),
	2731:  uint16(46),
	2732:  uint16(1),
	2733:  uint16(sym_documentation_comment),
	2734:  uint16(269),
	2735:  uint16(27),
	2737:  uint16(anon_sym_use),
	2738:  uint16(anon_sym_enum),
	2739:  uint16(anon_sym_intEnum),
	2740:  uint16(anon_sym_list),
	2741:  uint16(anon_sym_map),
	2742:  uint16(anon_sym_set),
	2743:  uint16(anon_sym_structure),
	2744:  uint16(anon_sym_union),
	2745:  uint16(anon_sym_service),
	2746:  uint16(anon_sym_operation),
	2747:  uint16(anon_sym_resource),
	2748:  uint16(anon_sym_AT),
	2749:  uint16(anon_sym_apply),
	2750:  uint16(anon_sym_blob),
	2751:  uint16(anon_sym_boolean),
	2752:  uint16(anon_sym_byte),
	2753:  uint16(anon_sym_document),
	2754:  uint16(anon_sym_double),
	2755:  uint16(anon_sym_float),
	2756:  uint16(anon_sym_integer),
	2757:  uint16(anon_sym_long),
	2758:  uint16(anon_sym_short),
	2759:  uint16(anon_sym_string),
	2760:  uint16(anon_sym_timestamp),
	2761:  uint16(anon_sym_bigInteger),
	2762:  uint16(anon_sym_bigDecimal),
	2763:  uint16(5),
	2764:  uint16(3),
	2765:  uint16(1),
	2766:  uint16(anon_sym_COMMA),
	2767:  uint16(5),
	2768:  uint16(1),
	2769:  uint16(sym_comment),
	2770:  uint16(7),
	2771:  uint16(1),
	2772:  uint16(anon_sym_SLASH_SLASH_SLASH),
	2773:  uint16(47),
	2774:  uint16(1),
	2775:  uint16(sym_documentation_comment),
	2776:  uint16(271),
	2777:  uint16(27),
	2779:  uint16(anon_sym_use),
	2780:  uint16(anon_sym_enum),
	2781:  uint16(anon_sym_intEnum),
	2782:  uint16(anon_sym_list),
	2783:  uint16(anon_sym_map),
	2784:  uint16(anon_sym_set),
	2785:  uint16(anon_sym_structure),
	2786:  uint16(anon_sym_union),
	2787:  uint16(anon_sym_service),
	2788:  uint16(anon_sym_operation),
	2789:  uint16(anon_sym_resource),
	2790:  uint16(anon_sym_AT),
	2791:  uint16(anon_sym_apply),
	2792:  uint16(anon_sym_blob),
	2793:  uint16(anon_sym_boolean),
	2794:  uint16(anon_sym_byte),
	2795:  uint16(anon_sym_document),
	2796:  uint16(anon_sym_double),
	2797:  uint16(anon_sym_float),
	2798:  uint16(anon_sym_integer),
	2799:  uint16(anon_sym_long),
	2800:  uint16(anon_sym_short),
	2801:  uint16(anon_sym_string),
	2802:  uint16(anon_sym_timestamp),
	2803:  uint16(anon_sym_bigInteger),
	2804:  uint16(anon_sym_bigDecimal),
	2805:  uint16(5),
	2806:  uint16(3),
	2807:  uint16(1),
	2808:  uint16(anon_sym_COMMA),
	2809:  uint16(5),
	2810:  uint16(1),
	2811:  uint16(sym_comment),
	2812:  uint16(7),
	2813:  uint16(1),
	2814:  uint16(anon_sym_SLASH_SLASH_SLASH),
	2815:  uint16(48),
	2816:  uint16(1),
	2817:  uint16(sym_documentation_comment),
	2818:  uint16(273),
	2819:  uint16(27),
	2821:  uint16(anon_sym_use),
	2822:  uint16(anon_sym_enum),
	2823:  uint16(anon_sym_intEnum),
	2824:  uint16(anon_sym_list),
	2825:  uint16(anon_sym_map),
	2826:  uint16(anon_sym_set),
	2827:  uint16(anon_sym_structure),
	2828:  uint16(anon_sym_union),
	2829:  uint16(anon_sym_service),
	2830:  uint16(anon_sym_operation),
	2831:  uint16(anon_sym_resource),
	2832:  uint16(anon_sym_AT),
	2833:  uint16(anon_sym_apply),
	2834:  uint16(anon_sym_blob),
	2835:  uint16(anon_sym_boolean),
	2836:  uint16(anon_sym_byte),
	2837:  uint16(anon_sym_document),
	2838:  uint16(anon_sym_double),
	2839:  uint16(anon_sym_float),
	2840:  uint16(anon_sym_integer),
	2841:  uint16(anon_sym_long),
	2842:  uint16(anon_sym_short),
	2843:  uint16(anon_sym_string),
	2844:  uint16(anon_sym_timestamp),
	2845:  uint16(anon_sym_bigInteger),
	2846:  uint16(anon_sym_bigDecimal),
	2847:  uint16(5),
	2848:  uint16(3),
	2849:  uint16(1),
	2850:  uint16(anon_sym_COMMA),
	2851:  uint16(5),
	2852:  uint16(1),
	2853:  uint16(sym_comment),
	2854:  uint16(7),
	2855:  uint16(1),
	2856:  uint16(anon_sym_SLASH_SLASH_SLASH),
	2857:  uint16(49),
	2858:  uint16(1),
	2859:  uint16(sym_documentation_comment),
	2860:  uint16(275),
	2861:  uint16(27),
	2863:  uint16(anon_sym_use),
	2864:  uint16(anon_sym_enum),
	2865:  uint16(anon_sym_intEnum),
	2866:  uint16(anon_sym_list),
	2867:  uint16(anon_sym_map),
	2868:  uint16(anon_sym_set),
	2869:  uint16(anon_sym_structure),
	2870:  uint16(anon_sym_union),
	2871:  uint16(anon_sym_service),
	2872:  uint16(anon_sym_operation),
	2873:  uint16(anon_sym_resource),
	2874:  uint16(anon_sym_AT),
	2875:  uint16(anon_sym_apply),
	2876:  uint16(anon_sym_blob),
	2877:  uint16(anon_sym_boolean),
	2878:  uint16(anon_sym_byte),
	2879:  uint16(anon_sym_document),
	2880:  uint16(anon_sym_double),
	2881:  uint16(anon_sym_float),
	2882:  uint16(anon_sym_integer),
	2883:  uint16(anon_sym_long),
	2884:  uint16(anon_sym_short),
	2885:  uint16(anon_sym_string),
	2886:  uint16(anon_sym_timestamp),
	2887:  uint16(anon_sym_bigInteger),
	2888:  uint16(anon_sym_bigDecimal),
	2889:  uint16(5),
	2890:  uint16(3),
	2891:  uint16(1),
	2892:  uint16(anon_sym_COMMA),
	2893:  uint16(5),
	2894:  uint16(1),
	2895:  uint16(sym_comment),
	2896:  uint16(7),
	2897:  uint16(1),
	2898:  uint16(anon_sym_SLASH_SLASH_SLASH),
	2899:  uint16(50),
	2900:  uint16(1),
	2901:  uint16(sym_documentation_comment),
	2902:  uint16(277),
	2903:  uint16(27),
	2905:  uint16(anon_sym_use),
	2906:  uint16(anon_sym_enum),
	2907:  uint16(anon_sym_intEnum),
	2908:  uint16(anon_sym_list),
	2909:  uint16(anon_sym_map),
	2910:  uint16(anon_sym_set),
	2911:  uint16(anon_sym_structure),
	2912:  uint16(anon_sym_union),
	2913:  uint16(anon_sym_service),
	2914:  uint16(anon_sym_operation),
	2915:  uint16(anon_sym_resource),
	2916:  uint16(anon_sym_AT),
	2917:  uint16(anon_sym_apply),
	2918:  uint16(anon_sym_blob),
	2919:  uint16(anon_sym_boolean),
	2920:  uint16(anon_sym_byte),
	2921:  uint16(anon_sym_document),
	2922:  uint16(anon_sym_double),
	2923:  uint16(anon_sym_float),
	2924:  uint16(anon_sym_integer),
	2925:  uint16(anon_sym_long),
	2926:  uint16(anon_sym_short),
	2927:  uint16(anon_sym_string),
	2928:  uint16(anon_sym_timestamp),
	2929:  uint16(anon_sym_bigInteger),
	2930:  uint16(anon_sym_bigDecimal),
	2931:  uint16(5),
	2932:  uint16(3),
	2933:  uint16(1),
	2934:  uint16(anon_sym_COMMA),
	2935:  uint16(5),
	2936:  uint16(1),
	2937:  uint16(sym_comment),
	2938:  uint16(7),
	2939:  uint16(1),
	2940:  uint16(anon_sym_SLASH_SLASH_SLASH),
	2941:  uint16(51),
	2942:  uint16(1),
	2943:  uint16(sym_documentation_comment),
	2944:  uint16(279),
	2945:  uint16(27),
	2947:  uint16(anon_sym_use),
	2948:  uint16(anon_sym_enum),
	2949:  uint16(anon_sym_intEnum),
	2950:  uint16(anon_sym_list),
	2951:  uint16(anon_sym_map),
	2952:  uint16(anon_sym_set),
	2953:  uint16(anon_sym_structure),
	2954:  uint16(anon_sym_union),
	2955:  uint16(anon_sym_service),
	2956:  uint16(anon_sym_operation),
	2957:  uint16(anon_sym_resource),
	2958:  uint16(anon_sym_AT),
	2959:  uint16(anon_sym_apply),
	2960:  uint16(anon_sym_blob),
	2961:  uint16(anon_sym_boolean),
	2962:  uint16(anon_sym_byte),
	2963:  uint16(anon_sym_document),
	2964:  uint16(anon_sym_double),
	2965:  uint16(anon_sym_float),
	2966:  uint16(anon_sym_integer),
	2967:  uint16(anon_sym_long),
	2968:  uint16(anon_sym_short),
	2969:  uint16(anon_sym_string),
	2970:  uint16(anon_sym_timestamp),
	2971:  uint16(anon_sym_bigInteger),
	2972:  uint16(anon_sym_bigDecimal),
	2973:  uint16(5),
	2974:  uint16(3),
	2975:  uint16(1),
	2976:  uint16(anon_sym_COMMA),
	2977:  uint16(5),
	2978:  uint16(1),
	2979:  uint16(sym_comment),
	2980:  uint16(7),
	2981:  uint16(1),
	2982:  uint16(anon_sym_SLASH_SLASH_SLASH),
	2983:  uint16(52),
	2984:  uint16(1),
	2985:  uint16(sym_documentation_comment),
	2986:  uint16(281),
	2987:  uint16(27),
	2989:  uint16(anon_sym_use),
	2990:  uint16(anon_sym_enum),
	2991:  uint16(anon_sym_intEnum),
	2992:  uint16(anon_sym_list),
	2993:  uint16(anon_sym_map),
	2994:  uint16(anon_sym_set),
	2995:  uint16(anon_sym_structure),
	2996:  uint16(anon_sym_union),
	2997:  uint16(anon_sym_service),
	2998:  uint16(anon_sym_operation),
	2999:  uint16(anon_sym_resource),
	3000:  uint16(anon_sym_AT),
	3001:  uint16(anon_sym_apply),
	3002:  uint16(anon_sym_blob),
	3003:  uint16(anon_sym_boolean),
	3004:  uint16(anon_sym_byte),
	3005:  uint16(anon_sym_document),
	3006:  uint16(anon_sym_double),
	3007:  uint16(anon_sym_float),
	3008:  uint16(anon_sym_integer),
	3009:  uint16(anon_sym_long),
	3010:  uint16(anon_sym_short),
	3011:  uint16(anon_sym_string),
	3012:  uint16(anon_sym_timestamp),
	3013:  uint16(anon_sym_bigInteger),
	3014:  uint16(anon_sym_bigDecimal),
	3015:  uint16(5),
	3016:  uint16(3),
	3017:  uint16(1),
	3018:  uint16(anon_sym_COMMA),
	3019:  uint16(5),
	3020:  uint16(1),
	3021:  uint16(sym_comment),
	3022:  uint16(7),
	3023:  uint16(1),
	3024:  uint16(anon_sym_SLASH_SLASH_SLASH),
	3025:  uint16(53),
	3026:  uint16(1),
	3027:  uint16(sym_documentation_comment),
	3028:  uint16(283),
	3029:  uint16(27),
	3031:  uint16(anon_sym_use),
	3032:  uint16(anon_sym_enum),
	3033:  uint16(anon_sym_intEnum),
	3034:  uint16(anon_sym_list),
	3035:  uint16(anon_sym_map),
	3036:  uint16(anon_sym_set),
	3037:  uint16(anon_sym_structure),
	3038:  uint16(anon_sym_union),
	3039:  uint16(anon_sym_service),
	3040:  uint16(anon_sym_operation),
	3041:  uint16(anon_sym_resource),
	3042:  uint16(anon_sym_AT),
	3043:  uint16(anon_sym_apply),
	3044:  uint16(anon_sym_blob),
	3045:  uint16(anon_sym_boolean),
	3046:  uint16(anon_sym_byte),
	3047:  uint16(anon_sym_document),
	3048:  uint16(anon_sym_double),
	3049:  uint16(anon_sym_float),
	3050:  uint16(anon_sym_integer),
	3051:  uint16(anon_sym_long),
	3052:  uint16(anon_sym_short),
	3053:  uint16(anon_sym_string),
	3054:  uint16(anon_sym_timestamp),
	3055:  uint16(anon_sym_bigInteger),
	3056:  uint16(anon_sym_bigDecimal),
	3057:  uint16(5),
	3058:  uint16(3),
	3059:  uint16(1),
	3060:  uint16(anon_sym_COMMA),
	3061:  uint16(5),
	3062:  uint16(1),
	3063:  uint16(sym_comment),
	3064:  uint16(7),
	3065:  uint16(1),
	3066:  uint16(anon_sym_SLASH_SLASH_SLASH),
	3067:  uint16(54),
	3068:  uint16(1),
	3069:  uint16(sym_documentation_comment),
	3070:  uint16(285),
	3071:  uint16(27),
	3073:  uint16(anon_sym_use),
	3074:  uint16(anon_sym_enum),
	3075:  uint16(anon_sym_intEnum),
	3076:  uint16(anon_sym_list),
	3077:  uint16(anon_sym_map),
	3078:  uint16(anon_sym_set),
	3079:  uint16(anon_sym_structure),
	3080:  uint16(anon_sym_union),
	3081:  uint16(anon_sym_service),
	3082:  uint16(anon_sym_operation),
	3083:  uint16(anon_sym_resource),
	3084:  uint16(anon_sym_AT),
	3085:  uint16(anon_sym_apply),
	3086:  uint16(anon_sym_blob),
	3087:  uint16(anon_sym_boolean),
	3088:  uint16(anon_sym_byte),
	3089:  uint16(anon_sym_document),
	3090:  uint16(anon_sym_double),
	3091:  uint16(anon_sym_float),
	3092:  uint16(anon_sym_integer),
	3093:  uint16(anon_sym_long),
	3094:  uint16(anon_sym_short),
	3095:  uint16(anon_sym_string),
	3096:  uint16(anon_sym_timestamp),
	3097:  uint16(anon_sym_bigInteger),
	3098:  uint16(anon_sym_bigDecimal),
	3099:  uint16(5),
	3100:  uint16(3),
	3101:  uint16(1),
	3102:  uint16(anon_sym_COMMA),
	3103:  uint16(5),
	3104:  uint16(1),
	3105:  uint16(sym_comment),
	3106:  uint16(7),
	3107:  uint16(1),
	3108:  uint16(anon_sym_SLASH_SLASH_SLASH),
	3109:  uint16(55),
	3110:  uint16(1),
	3111:  uint16(sym_documentation_comment),
	3112:  uint16(287),
	3113:  uint16(27),
	3115:  uint16(anon_sym_use),
	3116:  uint16(anon_sym_enum),
	3117:  uint16(anon_sym_intEnum),
	3118:  uint16(anon_sym_list),
	3119:  uint16(anon_sym_map),
	3120:  uint16(anon_sym_set),
	3121:  uint16(anon_sym_structure),
	3122:  uint16(anon_sym_union),
	3123:  uint16(anon_sym_service),
	3124:  uint16(anon_sym_operation),
	3125:  uint16(anon_sym_resource),
	3126:  uint16(anon_sym_AT),
	3127:  uint16(anon_sym_apply),
	3128:  uint16(anon_sym_blob),
	3129:  uint16(anon_sym_boolean),
	3130:  uint16(anon_sym_byte),
	3131:  uint16(anon_sym_document),
	3132:  uint16(anon_sym_double),
	3133:  uint16(anon_sym_float),
	3134:  uint16(anon_sym_integer),
	3135:  uint16(anon_sym_long),
	3136:  uint16(anon_sym_short),
	3137:  uint16(anon_sym_string),
	3138:  uint16(anon_sym_timestamp),
	3139:  uint16(anon_sym_bigInteger),
	3140:  uint16(anon_sym_bigDecimal),
	3141:  uint16(5),
	3142:  uint16(3),
	3143:  uint16(1),
	3144:  uint16(anon_sym_COMMA),
	3145:  uint16(5),
	3146:  uint16(1),
	3147:  uint16(sym_comment),
	3148:  uint16(7),
	3149:  uint16(1),
	3150:  uint16(anon_sym_SLASH_SLASH_SLASH),
	3151:  uint16(56),
	3152:  uint16(1),
	3153:  uint16(sym_documentation_comment),
	3154:  uint16(289),
	3155:  uint16(27),
	3157:  uint16(anon_sym_use),
	3158:  uint16(anon_sym_enum),
	3159:  uint16(anon_sym_intEnum),
	3160:  uint16(anon_sym_list),
	3161:  uint16(anon_sym_map),
	3162:  uint16(anon_sym_set),
	3163:  uint16(anon_sym_structure),
	3164:  uint16(anon_sym_union),
	3165:  uint16(anon_sym_service),
	3166:  uint16(anon_sym_operation),
	3167:  uint16(anon_sym_resource),
	3168:  uint16(anon_sym_AT),
	3169:  uint16(anon_sym_apply),
	3170:  uint16(anon_sym_blob),
	3171:  uint16(anon_sym_boolean),
	3172:  uint16(anon_sym_byte),
	3173:  uint16(anon_sym_document),
	3174:  uint16(anon_sym_double),
	3175:  uint16(anon_sym_float),
	3176:  uint16(anon_sym_integer),
	3177:  uint16(anon_sym_long),
	3178:  uint16(anon_sym_short),
	3179:  uint16(anon_sym_string),
	3180:  uint16(anon_sym_timestamp),
	3181:  uint16(anon_sym_bigInteger),
	3182:  uint16(anon_sym_bigDecimal),
	3183:  uint16(5),
	3184:  uint16(3),
	3185:  uint16(1),
	3186:  uint16(anon_sym_COMMA),
	3187:  uint16(5),
	3188:  uint16(1),
	3189:  uint16(sym_comment),
	3190:  uint16(7),
	3191:  uint16(1),
	3192:  uint16(anon_sym_SLASH_SLASH_SLASH),
	3193:  uint16(57),
	3194:  uint16(1),
	3195:  uint16(sym_documentation_comment),
	3196:  uint16(291),
	3197:  uint16(27),
	3199:  uint16(anon_sym_use),
	3200:  uint16(anon_sym_enum),
	3201:  uint16(anon_sym_intEnum),
	3202:  uint16(anon_sym_list),
	3203:  uint16(anon_sym_map),
	3204:  uint16(anon_sym_set),
	3205:  uint16(anon_sym_structure),
	3206:  uint16(anon_sym_union),
	3207:  uint16(anon_sym_service),
	3208:  uint16(anon_sym_operation),
	3209:  uint16(anon_sym_resource),
	3210:  uint16(anon_sym_AT),
	3211:  uint16(anon_sym_apply),
	3212:  uint16(anon_sym_blob),
	3213:  uint16(anon_sym_boolean),
	3214:  uint16(anon_sym_byte),
	3215:  uint16(anon_sym_document),
	3216:  uint16(anon_sym_double),
	3217:  uint16(anon_sym_float),
	3218:  uint16(anon_sym_integer),
	3219:  uint16(anon_sym_long),
	3220:  uint16(anon_sym_short),
	3221:  uint16(anon_sym_string),
	3222:  uint16(anon_sym_timestamp),
	3223:  uint16(anon_sym_bigInteger),
	3224:  uint16(anon_sym_bigDecimal),
	3225:  uint16(5),
	3226:  uint16(3),
	3227:  uint16(1),
	3228:  uint16(anon_sym_COMMA),
	3229:  uint16(5),
	3230:  uint16(1),
	3231:  uint16(sym_comment),
	3232:  uint16(7),
	3233:  uint16(1),
	3234:  uint16(anon_sym_SLASH_SLASH_SLASH),
	3235:  uint16(58),
	3236:  uint16(1),
	3237:  uint16(sym_documentation_comment),
	3238:  uint16(293),
	3239:  uint16(27),
	3241:  uint16(anon_sym_use),
	3242:  uint16(anon_sym_enum),
	3243:  uint16(anon_sym_intEnum),
	3244:  uint16(anon_sym_list),
	3245:  uint16(anon_sym_map),
	3246:  uint16(anon_sym_set),
	3247:  uint16(anon_sym_structure),
	3248:  uint16(anon_sym_union),
	3249:  uint16(anon_sym_service),
	3250:  uint16(anon_sym_operation),
	3251:  uint16(anon_sym_resource),
	3252:  uint16(anon_sym_AT),
	3253:  uint16(anon_sym_apply),
	3254:  uint16(anon_sym_blob),
	3255:  uint16(anon_sym_boolean),
	3256:  uint16(anon_sym_byte),
	3257:  uint16(anon_sym_document),
	3258:  uint16(anon_sym_double),
	3259:  uint16(anon_sym_float),
	3260:  uint16(anon_sym_integer),
	3261:  uint16(anon_sym_long),
	3262:  uint16(anon_sym_short),
	3263:  uint16(anon_sym_string),
	3264:  uint16(anon_sym_timestamp),
	3265:  uint16(anon_sym_bigInteger),
	3266:  uint16(anon_sym_bigDecimal),
	3267:  uint16(5),
	3268:  uint16(3),
	3269:  uint16(1),
	3270:  uint16(anon_sym_COMMA),
	3271:  uint16(5),
	3272:  uint16(1),
	3273:  uint16(sym_comment),
	3274:  uint16(7),
	3275:  uint16(1),
	3276:  uint16(anon_sym_SLASH_SLASH_SLASH),
	3277:  uint16(59),
	3278:  uint16(1),
	3279:  uint16(sym_documentation_comment),
	3280:  uint16(295),
	3281:  uint16(27),
	3283:  uint16(anon_sym_use),
	3284:  uint16(anon_sym_enum),
	3285:  uint16(anon_sym_intEnum),
	3286:  uint16(anon_sym_list),
	3287:  uint16(anon_sym_map),
	3288:  uint16(anon_sym_set),
	3289:  uint16(anon_sym_structure),
	3290:  uint16(anon_sym_union),
	3291:  uint16(anon_sym_service),
	3292:  uint16(anon_sym_operation),
	3293:  uint16(anon_sym_resource),
	3294:  uint16(anon_sym_AT),
	3295:  uint16(anon_sym_apply),
	3296:  uint16(anon_sym_blob),
	3297:  uint16(anon_sym_boolean),
	3298:  uint16(anon_sym_byte),
	3299:  uint16(anon_sym_document),
	3300:  uint16(anon_sym_double),
	3301:  uint16(anon_sym_float),
	3302:  uint16(anon_sym_integer),
	3303:  uint16(anon_sym_long),
	3304:  uint16(anon_sym_short),
	3305:  uint16(anon_sym_string),
	3306:  uint16(anon_sym_timestamp),
	3307:  uint16(anon_sym_bigInteger),
	3308:  uint16(anon_sym_bigDecimal),
	3309:  uint16(5),
	3310:  uint16(3),
	3311:  uint16(1),
	3312:  uint16(anon_sym_COMMA),
	3313:  uint16(5),
	3314:  uint16(1),
	3315:  uint16(sym_comment),
	3316:  uint16(7),
	3317:  uint16(1),
	3318:  uint16(anon_sym_SLASH_SLASH_SLASH),
	3319:  uint16(60),
	3320:  uint16(1),
	3321:  uint16(sym_documentation_comment),
	3322:  uint16(297),
	3323:  uint16(27),
	3325:  uint16(anon_sym_use),
	3326:  uint16(anon_sym_enum),
	3327:  uint16(anon_sym_intEnum),
	3328:  uint16(anon_sym_list),
	3329:  uint16(anon_sym_map),
	3330:  uint16(anon_sym_set),
	3331:  uint16(anon_sym_structure),
	3332:  uint16(anon_sym_union),
	3333:  uint16(anon_sym_service),
	3334:  uint16(anon_sym_operation),
	3335:  uint16(anon_sym_resource),
	3336:  uint16(anon_sym_AT),
	3337:  uint16(anon_sym_apply),
	3338:  uint16(anon_sym_blob),
	3339:  uint16(anon_sym_boolean),
	3340:  uint16(anon_sym_byte),
	3341:  uint16(anon_sym_document),
	3342:  uint16(anon_sym_double),
	3343:  uint16(anon_sym_float),
	3344:  uint16(anon_sym_integer),
	3345:  uint16(anon_sym_long),
	3346:  uint16(anon_sym_short),
	3347:  uint16(anon_sym_string),
	3348:  uint16(anon_sym_timestamp),
	3349:  uint16(anon_sym_bigInteger),
	3350:  uint16(anon_sym_bigDecimal),
	3351:  uint16(5),
	3352:  uint16(3),
	3353:  uint16(1),
	3354:  uint16(anon_sym_COMMA),
	3355:  uint16(5),
	3356:  uint16(1),
	3357:  uint16(sym_comment),
	3358:  uint16(7),
	3359:  uint16(1),
	3360:  uint16(anon_sym_SLASH_SLASH_SLASH),
	3361:  uint16(61),
	3362:  uint16(1),
	3363:  uint16(sym_documentation_comment),
	3364:  uint16(299),
	3365:  uint16(27),
	3367:  uint16(anon_sym_use),
	3368:  uint16(anon_sym_enum),
	3369:  uint16(anon_sym_intEnum),
	3370:  uint16(anon_sym_list),
	3371:  uint16(anon_sym_map),
	3372:  uint16(anon_sym_set),
	3373:  uint16(anon_sym_structure),
	3374:  uint16(anon_sym_union),
	3375:  uint16(anon_sym_service),
	3376:  uint16(anon_sym_operation),
	3377:  uint16(anon_sym_resource),
	3378:  uint16(anon_sym_AT),
	3379:  uint16(anon_sym_apply),
	3380:  uint16(anon_sym_blob),
	3381:  uint16(anon_sym_boolean),
	3382:  uint16(anon_sym_byte),
	3383:  uint16(anon_sym_document),
	3384:  uint16(anon_sym_double),
	3385:  uint16(anon_sym_float),
	3386:  uint16(anon_sym_integer),
	3387:  uint16(anon_sym_long),
	3388:  uint16(anon_sym_short),
	3389:  uint16(anon_sym_string),
	3390:  uint16(anon_sym_timestamp),
	3391:  uint16(anon_sym_bigInteger),
	3392:  uint16(anon_sym_bigDecimal),
	3393:  uint16(5),
	3394:  uint16(3),
	3395:  uint16(1),
	3396:  uint16(anon_sym_COMMA),
	3397:  uint16(5),
	3398:  uint16(1),
	3399:  uint16(sym_comment),
	3400:  uint16(7),
	3401:  uint16(1),
	3402:  uint16(anon_sym_SLASH_SLASH_SLASH),
	3403:  uint16(62),
	3404:  uint16(1),
	3405:  uint16(sym_documentation_comment),
	3406:  uint16(301),
	3407:  uint16(27),
	3409:  uint16(anon_sym_use),
	3410:  uint16(anon_sym_enum),
	3411:  uint16(anon_sym_intEnum),
	3412:  uint16(anon_sym_list),
	3413:  uint16(anon_sym_map),
	3414:  uint16(anon_sym_set),
	3415:  uint16(anon_sym_structure),
	3416:  uint16(anon_sym_union),
	3417:  uint16(anon_sym_service),
	3418:  uint16(anon_sym_operation),
	3419:  uint16(anon_sym_resource),
	3420:  uint16(anon_sym_AT),
	3421:  uint16(anon_sym_apply),
	3422:  uint16(anon_sym_blob),
	3423:  uint16(anon_sym_boolean),
	3424:  uint16(anon_sym_byte),
	3425:  uint16(anon_sym_document),
	3426:  uint16(anon_sym_double),
	3427:  uint16(anon_sym_float),
	3428:  uint16(anon_sym_integer),
	3429:  uint16(anon_sym_long),
	3430:  uint16(anon_sym_short),
	3431:  uint16(anon_sym_string),
	3432:  uint16(anon_sym_timestamp),
	3433:  uint16(anon_sym_bigInteger),
	3434:  uint16(anon_sym_bigDecimal),
	3435:  uint16(5),
	3436:  uint16(3),
	3437:  uint16(1),
	3438:  uint16(anon_sym_COMMA),
	3439:  uint16(5),
	3440:  uint16(1),
	3441:  uint16(sym_comment),
	3442:  uint16(7),
	3443:  uint16(1),
	3444:  uint16(anon_sym_SLASH_SLASH_SLASH),
	3445:  uint16(63),
	3446:  uint16(1),
	3447:  uint16(sym_documentation_comment),
	3448:  uint16(303),
	3449:  uint16(27),
	3451:  uint16(anon_sym_use),
	3452:  uint16(anon_sym_enum),
	3453:  uint16(anon_sym_intEnum),
	3454:  uint16(anon_sym_list),
	3455:  uint16(anon_sym_map),
	3456:  uint16(anon_sym_set),
	3457:  uint16(anon_sym_structure),
	3458:  uint16(anon_sym_union),
	3459:  uint16(anon_sym_service),
	3460:  uint16(anon_sym_operation),
	3461:  uint16(anon_sym_resource),
	3462:  uint16(anon_sym_AT),
	3463:  uint16(anon_sym_apply),
	3464:  uint16(anon_sym_blob),
	3465:  uint16(anon_sym_boolean),
	3466:  uint16(anon_sym_byte),
	3467:  uint16(anon_sym_document),
	3468:  uint16(anon_sym_double),
	3469:  uint16(anon_sym_float),
	3470:  uint16(anon_sym_integer),
	3471:  uint16(anon_sym_long),
	3472:  uint16(anon_sym_short),
	3473:  uint16(anon_sym_string),
	3474:  uint16(anon_sym_timestamp),
	3475:  uint16(anon_sym_bigInteger),
	3476:  uint16(anon_sym_bigDecimal),
	3477:  uint16(5),
	3478:  uint16(3),
	3479:  uint16(1),
	3480:  uint16(anon_sym_COMMA),
	3481:  uint16(5),
	3482:  uint16(1),
	3483:  uint16(sym_comment),
	3484:  uint16(7),
	3485:  uint16(1),
	3486:  uint16(anon_sym_SLASH_SLASH_SLASH),
	3487:  uint16(64),
	3488:  uint16(1),
	3489:  uint16(sym_documentation_comment),
	3490:  uint16(305),
	3491:  uint16(27),
	3493:  uint16(anon_sym_use),
	3494:  uint16(anon_sym_enum),
	3495:  uint16(anon_sym_intEnum),
	3496:  uint16(anon_sym_list),
	3497:  uint16(anon_sym_map),
	3498:  uint16(anon_sym_set),
	3499:  uint16(anon_sym_structure),
	3500:  uint16(anon_sym_union),
	3501:  uint16(anon_sym_service),
	3502:  uint16(anon_sym_operation),
	3503:  uint16(anon_sym_resource),
	3504:  uint16(anon_sym_AT),
	3505:  uint16(anon_sym_apply),
	3506:  uint16(anon_sym_blob),
	3507:  uint16(anon_sym_boolean),
	3508:  uint16(anon_sym_byte),
	3509:  uint16(anon_sym_document),
	3510:  uint16(anon_sym_double),
	3511:  uint16(anon_sym_float),
	3512:  uint16(anon_sym_integer),
	3513:  uint16(anon_sym_long),
	3514:  uint16(anon_sym_short),
	3515:  uint16(anon_sym_string),
	3516:  uint16(anon_sym_timestamp),
	3517:  uint16(anon_sym_bigInteger),
	3518:  uint16(anon_sym_bigDecimal),
	3519:  uint16(22),
	3520:  uint16(3),
	3521:  uint16(1),
	3522:  uint16(anon_sym_COMMA),
	3523:  uint16(5),
	3524:  uint16(1),
	3525:  uint16(sym_comment),
	3526:  uint16(7),
	3527:  uint16(1),
	3528:  uint16(anon_sym_SLASH_SLASH_SLASH),
	3529:  uint16(239),
	3530:  uint16(1),
	3531:  uint16(anon_sym_LBRACE),
	3532:  uint16(241),
	3533:  uint16(1),
	3534:  uint16(anon_sym_LBRACK),
	3535:  uint16(245),
	3536:  uint16(1),
	3537:  uint16(sym_null),
	3538:  uint16(247),
	3539:  uint16(1),
	3540:  uint16(anon_sym_DASH),
	3541:  uint16(249),
	3542:  uint16(1),
	3543:  uint16(aux_sym_number_token1),
	3544:  uint16(251),
	3545:  uint16(1),
	3546:  uint16(aux_sym_float_token1),
	3547:  uint16(253),
	3548:  uint16(1),
	3549:  uint16(anon_sym_DQUOTE),
	3550:  uint16(255),
	3551:  uint16(1),
	3552:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	3553:  uint16(257),
	3554:  uint16(1),
	3555:  uint16(aux_sym_identifier_token1),
	3556:  uint16(23),
	3557:  uint16(1),
	3558:  uint16(sym__namespace_identifier),
	3559:  uint16(65),
	3560:  uint16(1),
	3561:  uint16(sym_documentation_comment),
	3562:  uint16(120),
	3563:  uint16(1),
	3564:  uint16(sym_root_shape_id),
	3565:  uint16(207),
	3566:  uint16(1),
	3567:  uint16(sym_node_value),
	3568:  uint16(345),
	3569:  uint16(1),
	3570:  uint16(sym_namespace),
	3571:  uint16(243),
	3572:  uint16(2),
	3573:  uint16(anon_sym_true),
	3574:  uint16(anon_sym_false),
	3575:  uint16(130),
	3576:  uint16(2),
	3577:  uint16(sym_absolute_root_shape_id),
	3578:  uint16(sym_identifier),
	3579:  uint16(169),
	3580:  uint16(2),
	3581:  uint16(sym__string_literal),
	3582:  uint16(sym__multiline_string_literal),
	3583:  uint16(174),
	3584:  uint16(4),
	3585:  uint16(sym_boolean),
	3586:  uint16(sym_number),
	3587:  uint16(sym_float),
	3588:  uint16(sym_string),
	3589:  uint16(184),
	3590:  uint16(4),
	3591:  uint16(sym_shape_id),
	3592:  uint16(sym_node_array),
	3593:  uint16(sym_node_object),
	3594:  uint16(sym_literal),
	3595:  uint16(5),
	3596:  uint16(3),
	3597:  uint16(1),
	3598:  uint16(anon_sym_COMMA),
	3599:  uint16(5),
	3600:  uint16(1),
	3601:  uint16(sym_comment),
	3602:  uint16(7),
	3603:  uint16(1),
	3604:  uint16(anon_sym_SLASH_SLASH_SLASH),
	3605:  uint16(66),
	3606:  uint16(1),
	3607:  uint16(sym_documentation_comment),
	3608:  uint16(307),
	3609:  uint16(27),
	3611:  uint16(anon_sym_use),
	3612:  uint16(anon_sym_enum),
	3613:  uint16(anon_sym_intEnum),
	3614:  uint16(anon_sym_list),
	3615:  uint16(anon_sym_map),
	3616:  uint16(anon_sym_set),
	3617:  uint16(anon_sym_structure),
	3618:  uint16(anon_sym_union),
	3619:  uint16(anon_sym_service),
	3620:  uint16(anon_sym_operation),
	3621:  uint16(anon_sym_resource),
	3622:  uint16(anon_sym_AT),
	3623:  uint16(anon_sym_apply),
	3624:  uint16(anon_sym_blob),
	3625:  uint16(anon_sym_boolean),
	3626:  uint16(anon_sym_byte),
	3627:  uint16(anon_sym_document),
	3628:  uint16(anon_sym_double),
	3629:  uint16(anon_sym_float),
	3630:  uint16(anon_sym_integer),
	3631:  uint16(anon_sym_long),
	3632:  uint16(anon_sym_short),
	3633:  uint16(anon_sym_string),
	3634:  uint16(anon_sym_timestamp),
	3635:  uint16(anon_sym_bigInteger),
	3636:  uint16(anon_sym_bigDecimal),
	3637:  uint16(5),
	3638:  uint16(3),
	3639:  uint16(1),
	3640:  uint16(anon_sym_COMMA),
	3641:  uint16(5),
	3642:  uint16(1),
	3643:  uint16(sym_comment),
	3644:  uint16(7),
	3645:  uint16(1),
	3646:  uint16(anon_sym_SLASH_SLASH_SLASH),
	3647:  uint16(67),
	3648:  uint16(1),
	3649:  uint16(sym_documentation_comment),
	3650:  uint16(309),
	3651:  uint16(27),
	3653:  uint16(anon_sym_use),
	3654:  uint16(anon_sym_enum),
	3655:  uint16(anon_sym_intEnum),
	3656:  uint16(anon_sym_list),
	3657:  uint16(anon_sym_map),
	3658:  uint16(anon_sym_set),
	3659:  uint16(anon_sym_structure),
	3660:  uint16(anon_sym_union),
	3661:  uint16(anon_sym_service),
	3662:  uint16(anon_sym_operation),
	3663:  uint16(anon_sym_resource),
	3664:  uint16(anon_sym_AT),
	3665:  uint16(anon_sym_apply),
	3666:  uint16(anon_sym_blob),
	3667:  uint16(anon_sym_boolean),
	3668:  uint16(anon_sym_byte),
	3669:  uint16(anon_sym_document),
	3670:  uint16(anon_sym_double),
	3671:  uint16(anon_sym_float),
	3672:  uint16(anon_sym_integer),
	3673:  uint16(anon_sym_long),
	3674:  uint16(anon_sym_short),
	3675:  uint16(anon_sym_string),
	3676:  uint16(anon_sym_timestamp),
	3677:  uint16(anon_sym_bigInteger),
	3678:  uint16(anon_sym_bigDecimal),
	3679:  uint16(5),
	3680:  uint16(3),
	3681:  uint16(1),
	3682:  uint16(anon_sym_COMMA),
	3683:  uint16(5),
	3684:  uint16(1),
	3685:  uint16(sym_comment),
	3686:  uint16(7),
	3687:  uint16(1),
	3688:  uint16(anon_sym_SLASH_SLASH_SLASH),
	3689:  uint16(68),
	3690:  uint16(1),
	3691:  uint16(sym_documentation_comment),
	3692:  uint16(311),
	3693:  uint16(27),
	3695:  uint16(anon_sym_use),
	3696:  uint16(anon_sym_enum),
	3697:  uint16(anon_sym_intEnum),
	3698:  uint16(anon_sym_list),
	3699:  uint16(anon_sym_map),
	3700:  uint16(anon_sym_set),
	3701:  uint16(anon_sym_structure),
	3702:  uint16(anon_sym_union),
	3703:  uint16(anon_sym_service),
	3704:  uint16(anon_sym_operation),
	3705:  uint16(anon_sym_resource),
	3706:  uint16(anon_sym_AT),
	3707:  uint16(anon_sym_apply),
	3708:  uint16(anon_sym_blob),
	3709:  uint16(anon_sym_boolean),
	3710:  uint16(anon_sym_byte),
	3711:  uint16(anon_sym_document),
	3712:  uint16(anon_sym_double),
	3713:  uint16(anon_sym_float),
	3714:  uint16(anon_sym_integer),
	3715:  uint16(anon_sym_long),
	3716:  uint16(anon_sym_short),
	3717:  uint16(anon_sym_string),
	3718:  uint16(anon_sym_timestamp),
	3719:  uint16(anon_sym_bigInteger),
	3720:  uint16(anon_sym_bigDecimal),
	3721:  uint16(5),
	3722:  uint16(3),
	3723:  uint16(1),
	3724:  uint16(anon_sym_COMMA),
	3725:  uint16(5),
	3726:  uint16(1),
	3727:  uint16(sym_comment),
	3728:  uint16(7),
	3729:  uint16(1),
	3730:  uint16(anon_sym_SLASH_SLASH_SLASH),
	3731:  uint16(69),
	3732:  uint16(1),
	3733:  uint16(sym_documentation_comment),
	3734:  uint16(313),
	3735:  uint16(27),
	3737:  uint16(anon_sym_use),
	3738:  uint16(anon_sym_enum),
	3739:  uint16(anon_sym_intEnum),
	3740:  uint16(anon_sym_list),
	3741:  uint16(anon_sym_map),
	3742:  uint16(anon_sym_set),
	3743:  uint16(anon_sym_structure),
	3744:  uint16(anon_sym_union),
	3745:  uint16(anon_sym_service),
	3746:  uint16(anon_sym_operation),
	3747:  uint16(anon_sym_resource),
	3748:  uint16(anon_sym_AT),
	3749:  uint16(anon_sym_apply),
	3750:  uint16(anon_sym_blob),
	3751:  uint16(anon_sym_boolean),
	3752:  uint16(anon_sym_byte),
	3753:  uint16(anon_sym_document),
	3754:  uint16(anon_sym_double),
	3755:  uint16(anon_sym_float),
	3756:  uint16(anon_sym_integer),
	3757:  uint16(anon_sym_long),
	3758:  uint16(anon_sym_short),
	3759:  uint16(anon_sym_string),
	3760:  uint16(anon_sym_timestamp),
	3761:  uint16(anon_sym_bigInteger),
	3762:  uint16(anon_sym_bigDecimal),
	3763:  uint16(5),
	3764:  uint16(3),
	3765:  uint16(1),
	3766:  uint16(anon_sym_COMMA),
	3767:  uint16(5),
	3768:  uint16(1),
	3769:  uint16(sym_comment),
	3770:  uint16(7),
	3771:  uint16(1),
	3772:  uint16(anon_sym_SLASH_SLASH_SLASH),
	3773:  uint16(70),
	3774:  uint16(1),
	3775:  uint16(sym_documentation_comment),
	3776:  uint16(315),
	3777:  uint16(27),
	3779:  uint16(anon_sym_use),
	3780:  uint16(anon_sym_enum),
	3781:  uint16(anon_sym_intEnum),
	3782:  uint16(anon_sym_list),
	3783:  uint16(anon_sym_map),
	3784:  uint16(anon_sym_set),
	3785:  uint16(anon_sym_structure),
	3786:  uint16(anon_sym_union),
	3787:  uint16(anon_sym_service),
	3788:  uint16(anon_sym_operation),
	3789:  uint16(anon_sym_resource),
	3790:  uint16(anon_sym_AT),
	3791:  uint16(anon_sym_apply),
	3792:  uint16(anon_sym_blob),
	3793:  uint16(anon_sym_boolean),
	3794:  uint16(anon_sym_byte),
	3795:  uint16(anon_sym_document),
	3796:  uint16(anon_sym_double),
	3797:  uint16(anon_sym_float),
	3798:  uint16(anon_sym_integer),
	3799:  uint16(anon_sym_long),
	3800:  uint16(anon_sym_short),
	3801:  uint16(anon_sym_string),
	3802:  uint16(anon_sym_timestamp),
	3803:  uint16(anon_sym_bigInteger),
	3804:  uint16(anon_sym_bigDecimal),
	3805:  uint16(5),
	3806:  uint16(3),
	3807:  uint16(1),
	3808:  uint16(anon_sym_COMMA),
	3809:  uint16(5),
	3810:  uint16(1),
	3811:  uint16(sym_comment),
	3812:  uint16(7),
	3813:  uint16(1),
	3814:  uint16(anon_sym_SLASH_SLASH_SLASH),
	3815:  uint16(71),
	3816:  uint16(1),
	3817:  uint16(sym_documentation_comment),
	3818:  uint16(317),
	3819:  uint16(27),
	3821:  uint16(anon_sym_use),
	3822:  uint16(anon_sym_enum),
	3823:  uint16(anon_sym_intEnum),
	3824:  uint16(anon_sym_list),
	3825:  uint16(anon_sym_map),
	3826:  uint16(anon_sym_set),
	3827:  uint16(anon_sym_structure),
	3828:  uint16(anon_sym_union),
	3829:  uint16(anon_sym_service),
	3830:  uint16(anon_sym_operation),
	3831:  uint16(anon_sym_resource),
	3832:  uint16(anon_sym_AT),
	3833:  uint16(anon_sym_apply),
	3834:  uint16(anon_sym_blob),
	3835:  uint16(anon_sym_boolean),
	3836:  uint16(anon_sym_byte),
	3837:  uint16(anon_sym_document),
	3838:  uint16(anon_sym_double),
	3839:  uint16(anon_sym_float),
	3840:  uint16(anon_sym_integer),
	3841:  uint16(anon_sym_long),
	3842:  uint16(anon_sym_short),
	3843:  uint16(anon_sym_string),
	3844:  uint16(anon_sym_timestamp),
	3845:  uint16(anon_sym_bigInteger),
	3846:  uint16(anon_sym_bigDecimal),
	3847:  uint16(5),
	3848:  uint16(3),
	3849:  uint16(1),
	3850:  uint16(anon_sym_COMMA),
	3851:  uint16(5),
	3852:  uint16(1),
	3853:  uint16(sym_comment),
	3854:  uint16(7),
	3855:  uint16(1),
	3856:  uint16(anon_sym_SLASH_SLASH_SLASH),
	3857:  uint16(72),
	3858:  uint16(1),
	3859:  uint16(sym_documentation_comment),
	3860:  uint16(319),
	3861:  uint16(27),
	3863:  uint16(anon_sym_use),
	3864:  uint16(anon_sym_enum),
	3865:  uint16(anon_sym_intEnum),
	3866:  uint16(anon_sym_list),
	3867:  uint16(anon_sym_map),
	3868:  uint16(anon_sym_set),
	3869:  uint16(anon_sym_structure),
	3870:  uint16(anon_sym_union),
	3871:  uint16(anon_sym_service),
	3872:  uint16(anon_sym_operation),
	3873:  uint16(anon_sym_resource),
	3874:  uint16(anon_sym_AT),
	3875:  uint16(anon_sym_apply),
	3876:  uint16(anon_sym_blob),
	3877:  uint16(anon_sym_boolean),
	3878:  uint16(anon_sym_byte),
	3879:  uint16(anon_sym_document),
	3880:  uint16(anon_sym_double),
	3881:  uint16(anon_sym_float),
	3882:  uint16(anon_sym_integer),
	3883:  uint16(anon_sym_long),
	3884:  uint16(anon_sym_short),
	3885:  uint16(anon_sym_string),
	3886:  uint16(anon_sym_timestamp),
	3887:  uint16(anon_sym_bigInteger),
	3888:  uint16(anon_sym_bigDecimal),
	3889:  uint16(5),
	3890:  uint16(3),
	3891:  uint16(1),
	3892:  uint16(anon_sym_COMMA),
	3893:  uint16(5),
	3894:  uint16(1),
	3895:  uint16(sym_comment),
	3896:  uint16(7),
	3897:  uint16(1),
	3898:  uint16(anon_sym_SLASH_SLASH_SLASH),
	3899:  uint16(73),
	3900:  uint16(1),
	3901:  uint16(sym_documentation_comment),
	3902:  uint16(321),
	3903:  uint16(27),
	3905:  uint16(anon_sym_use),
	3906:  uint16(anon_sym_enum),
	3907:  uint16(anon_sym_intEnum),
	3908:  uint16(anon_sym_list),
	3909:  uint16(anon_sym_map),
	3910:  uint16(anon_sym_set),
	3911:  uint16(anon_sym_structure),
	3912:  uint16(anon_sym_union),
	3913:  uint16(anon_sym_service),
	3914:  uint16(anon_sym_operation),
	3915:  uint16(anon_sym_resource),
	3916:  uint16(anon_sym_AT),
	3917:  uint16(anon_sym_apply),
	3918:  uint16(anon_sym_blob),
	3919:  uint16(anon_sym_boolean),
	3920:  uint16(anon_sym_byte),
	3921:  uint16(anon_sym_document),
	3922:  uint16(anon_sym_double),
	3923:  uint16(anon_sym_float),
	3924:  uint16(anon_sym_integer),
	3925:  uint16(anon_sym_long),
	3926:  uint16(anon_sym_short),
	3927:  uint16(anon_sym_string),
	3928:  uint16(anon_sym_timestamp),
	3929:  uint16(anon_sym_bigInteger),
	3930:  uint16(anon_sym_bigDecimal),
	3931:  uint16(5),
	3932:  uint16(3),
	3933:  uint16(1),
	3934:  uint16(anon_sym_COMMA),
	3935:  uint16(5),
	3936:  uint16(1),
	3937:  uint16(sym_comment),
	3938:  uint16(7),
	3939:  uint16(1),
	3940:  uint16(anon_sym_SLASH_SLASH_SLASH),
	3941:  uint16(74),
	3942:  uint16(1),
	3943:  uint16(sym_documentation_comment),
	3944:  uint16(323),
	3945:  uint16(27),
	3947:  uint16(anon_sym_use),
	3948:  uint16(anon_sym_enum),
	3949:  uint16(anon_sym_intEnum),
	3950:  uint16(anon_sym_list),
	3951:  uint16(anon_sym_map),
	3952:  uint16(anon_sym_set),
	3953:  uint16(anon_sym_structure),
	3954:  uint16(anon_sym_union),
	3955:  uint16(anon_sym_service),
	3956:  uint16(anon_sym_operation),
	3957:  uint16(anon_sym_resource),
	3958:  uint16(anon_sym_AT),
	3959:  uint16(anon_sym_apply),
	3960:  uint16(anon_sym_blob),
	3961:  uint16(anon_sym_boolean),
	3962:  uint16(anon_sym_byte),
	3963:  uint16(anon_sym_document),
	3964:  uint16(anon_sym_double),
	3965:  uint16(anon_sym_float),
	3966:  uint16(anon_sym_integer),
	3967:  uint16(anon_sym_long),
	3968:  uint16(anon_sym_short),
	3969:  uint16(anon_sym_string),
	3970:  uint16(anon_sym_timestamp),
	3971:  uint16(anon_sym_bigInteger),
	3972:  uint16(anon_sym_bigDecimal),
	3973:  uint16(5),
	3974:  uint16(3),
	3975:  uint16(1),
	3976:  uint16(anon_sym_COMMA),
	3977:  uint16(5),
	3978:  uint16(1),
	3979:  uint16(sym_comment),
	3980:  uint16(7),
	3981:  uint16(1),
	3982:  uint16(anon_sym_SLASH_SLASH_SLASH),
	3983:  uint16(75),
	3984:  uint16(1),
	3985:  uint16(sym_documentation_comment),
	3986:  uint16(325),
	3987:  uint16(27),
	3989:  uint16(anon_sym_use),
	3990:  uint16(anon_sym_enum),
	3991:  uint16(anon_sym_intEnum),
	3992:  uint16(anon_sym_list),
	3993:  uint16(anon_sym_map),
	3994:  uint16(anon_sym_set),
	3995:  uint16(anon_sym_structure),
	3996:  uint16(anon_sym_union),
	3997:  uint16(anon_sym_service),
	3998:  uint16(anon_sym_operation),
	3999:  uint16(anon_sym_resource),
	4000:  uint16(anon_sym_AT),
	4001:  uint16(anon_sym_apply),
	4002:  uint16(anon_sym_blob),
	4003:  uint16(anon_sym_boolean),
	4004:  uint16(anon_sym_byte),
	4005:  uint16(anon_sym_document),
	4006:  uint16(anon_sym_double),
	4007:  uint16(anon_sym_float),
	4008:  uint16(anon_sym_integer),
	4009:  uint16(anon_sym_long),
	4010:  uint16(anon_sym_short),
	4011:  uint16(anon_sym_string),
	4012:  uint16(anon_sym_timestamp),
	4013:  uint16(anon_sym_bigInteger),
	4014:  uint16(anon_sym_bigDecimal),
	4015:  uint16(5),
	4016:  uint16(3),
	4017:  uint16(1),
	4018:  uint16(anon_sym_COMMA),
	4019:  uint16(5),
	4020:  uint16(1),
	4021:  uint16(sym_comment),
	4022:  uint16(7),
	4023:  uint16(1),
	4024:  uint16(anon_sym_SLASH_SLASH_SLASH),
	4025:  uint16(76),
	4026:  uint16(1),
	4027:  uint16(sym_documentation_comment),
	4028:  uint16(327),
	4029:  uint16(27),
	4031:  uint16(anon_sym_use),
	4032:  uint16(anon_sym_enum),
	4033:  uint16(anon_sym_intEnum),
	4034:  uint16(anon_sym_list),
	4035:  uint16(anon_sym_map),
	4036:  uint16(anon_sym_set),
	4037:  uint16(anon_sym_structure),
	4038:  uint16(anon_sym_union),
	4039:  uint16(anon_sym_service),
	4040:  uint16(anon_sym_operation),
	4041:  uint16(anon_sym_resource),
	4042:  uint16(anon_sym_AT),
	4043:  uint16(anon_sym_apply),
	4044:  uint16(anon_sym_blob),
	4045:  uint16(anon_sym_boolean),
	4046:  uint16(anon_sym_byte),
	4047:  uint16(anon_sym_document),
	4048:  uint16(anon_sym_double),
	4049:  uint16(anon_sym_float),
	4050:  uint16(anon_sym_integer),
	4051:  uint16(anon_sym_long),
	4052:  uint16(anon_sym_short),
	4053:  uint16(anon_sym_string),
	4054:  uint16(anon_sym_timestamp),
	4055:  uint16(anon_sym_bigInteger),
	4056:  uint16(anon_sym_bigDecimal),
	4057:  uint16(22),
	4058:  uint16(3),
	4059:  uint16(1),
	4060:  uint16(anon_sym_COMMA),
	4061:  uint16(5),
	4062:  uint16(1),
	4063:  uint16(sym_comment),
	4064:  uint16(7),
	4065:  uint16(1),
	4066:  uint16(anon_sym_SLASH_SLASH_SLASH),
	4067:  uint16(117),
	4068:  uint16(1),
	4069:  uint16(anon_sym_LBRACE),
	4070:  uint16(119),
	4071:  uint16(1),
	4072:  uint16(anon_sym_LBRACK),
	4073:  uint16(125),
	4074:  uint16(1),
	4075:  uint16(sym_null),
	4076:  uint16(127),
	4077:  uint16(1),
	4078:  uint16(anon_sym_DASH),
	4079:  uint16(129),
	4080:  uint16(1),
	4081:  uint16(aux_sym_number_token1),
	4082:  uint16(131),
	4083:  uint16(1),
	4084:  uint16(aux_sym_float_token1),
	4085:  uint16(133),
	4086:  uint16(1),
	4087:  uint16(anon_sym_DQUOTE),
	4088:  uint16(135),
	4089:  uint16(1),
	4090:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4091:  uint16(137),
	4092:  uint16(1),
	4093:  uint16(aux_sym_identifier_token1),
	4094:  uint16(9),
	4095:  uint16(1),
	4096:  uint16(sym_root_shape_id),
	4097:  uint16(23),
	4098:  uint16(1),
	4099:  uint16(sym__namespace_identifier),
	4100:  uint16(77),
	4101:  uint16(1),
	4102:  uint16(sym_documentation_comment),
	4103:  uint16(265),
	4104:  uint16(1),
	4105:  uint16(sym_node_value),
	4106:  uint16(332),
	4107:  uint16(1),
	4108:  uint16(sym_namespace),
	4109:  uint16(123),
	4110:  uint16(2),
	4111:  uint16(anon_sym_true),
	4112:  uint16(anon_sym_false),
	4113:  uint16(11),
	4114:  uint16(2),
	4115:  uint16(sym_absolute_root_shape_id),
	4116:  uint16(sym_identifier),
	4117:  uint16(168),
	4118:  uint16(2),
	4119:  uint16(sym__string_literal),
	4120:  uint16(sym__multiline_string_literal),
	4121:  uint16(209),
	4122:  uint16(4),
	4123:  uint16(sym_boolean),
	4124:  uint16(sym_number),
	4125:  uint16(sym_float),
	4126:  uint16(sym_string),
	4127:  uint16(226),
	4128:  uint16(4),
	4129:  uint16(sym_shape_id),
	4130:  uint16(sym_node_array),
	4131:  uint16(sym_node_object),
	4132:  uint16(sym_literal),
	4133:  uint16(5),
	4134:  uint16(3),
	4135:  uint16(1),
	4136:  uint16(anon_sym_COMMA),
	4137:  uint16(5),
	4138:  uint16(1),
	4139:  uint16(sym_comment),
	4140:  uint16(7),
	4141:  uint16(1),
	4142:  uint16(anon_sym_SLASH_SLASH_SLASH),
	4143:  uint16(78),
	4144:  uint16(1),
	4145:  uint16(sym_documentation_comment),
	4146:  uint16(329),
	4147:  uint16(27),
	4149:  uint16(anon_sym_use),
	4150:  uint16(anon_sym_enum),
	4151:  uint16(anon_sym_intEnum),
	4152:  uint16(anon_sym_list),
	4153:  uint16(anon_sym_map),
	4154:  uint16(anon_sym_set),
	4155:  uint16(anon_sym_structure),
	4156:  uint16(anon_sym_union),
	4157:  uint16(anon_sym_service),
	4158:  uint16(anon_sym_operation),
	4159:  uint16(anon_sym_resource),
	4160:  uint16(anon_sym_AT),
	4161:  uint16(anon_sym_apply),
	4162:  uint16(anon_sym_blob),
	4163:  uint16(anon_sym_boolean),
	4164:  uint16(anon_sym_byte),
	4165:  uint16(anon_sym_document),
	4166:  uint16(anon_sym_double),
	4167:  uint16(anon_sym_float),
	4168:  uint16(anon_sym_integer),
	4169:  uint16(anon_sym_long),
	4170:  uint16(anon_sym_short),
	4171:  uint16(anon_sym_string),
	4172:  uint16(anon_sym_timestamp),
	4173:  uint16(anon_sym_bigInteger),
	4174:  uint16(anon_sym_bigDecimal),
	4175:  uint16(22),
	4176:  uint16(3),
	4177:  uint16(1),
	4178:  uint16(anon_sym_COMMA),
	4179:  uint16(5),
	4180:  uint16(1),
	4181:  uint16(sym_comment),
	4182:  uint16(7),
	4183:  uint16(1),
	4184:  uint16(anon_sym_SLASH_SLASH_SLASH),
	4185:  uint16(117),
	4186:  uint16(1),
	4187:  uint16(anon_sym_LBRACE),
	4188:  uint16(119),
	4189:  uint16(1),
	4190:  uint16(anon_sym_LBRACK),
	4191:  uint16(125),
	4192:  uint16(1),
	4193:  uint16(sym_null),
	4194:  uint16(127),
	4195:  uint16(1),
	4196:  uint16(anon_sym_DASH),
	4197:  uint16(129),
	4198:  uint16(1),
	4199:  uint16(aux_sym_number_token1),
	4200:  uint16(131),
	4201:  uint16(1),
	4202:  uint16(aux_sym_float_token1),
	4203:  uint16(133),
	4204:  uint16(1),
	4205:  uint16(anon_sym_DQUOTE),
	4206:  uint16(135),
	4207:  uint16(1),
	4208:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4209:  uint16(137),
	4210:  uint16(1),
	4211:  uint16(aux_sym_identifier_token1),
	4212:  uint16(23),
	4213:  uint16(1),
	4214:  uint16(sym__namespace_identifier),
	4215:  uint16(79),
	4216:  uint16(1),
	4217:  uint16(sym_documentation_comment),
	4218:  uint16(191),
	4219:  uint16(1),
	4220:  uint16(sym_root_shape_id),
	4221:  uint16(243),
	4222:  uint16(1),
	4223:  uint16(sym_node_value),
	4224:  uint16(332),
	4225:  uint16(1),
	4226:  uint16(sym_namespace),
	4227:  uint16(123),
	4228:  uint16(2),
	4229:  uint16(anon_sym_true),
	4230:  uint16(anon_sym_false),
	4231:  uint16(11),
	4232:  uint16(2),
	4233:  uint16(sym_absolute_root_shape_id),
	4234:  uint16(sym_identifier),
	4235:  uint16(168),
	4236:  uint16(2),
	4237:  uint16(sym__string_literal),
	4238:  uint16(sym__multiline_string_literal),
	4239:  uint16(209),
	4240:  uint16(4),
	4241:  uint16(sym_boolean),
	4242:  uint16(sym_number),
	4243:  uint16(sym_float),
	4244:  uint16(sym_string),
	4245:  uint16(226),
	4246:  uint16(4),
	4247:  uint16(sym_shape_id),
	4248:  uint16(sym_node_array),
	4249:  uint16(sym_node_object),
	4250:  uint16(sym_literal),
	4251:  uint16(5),
	4252:  uint16(3),
	4253:  uint16(1),
	4254:  uint16(anon_sym_COMMA),
	4255:  uint16(5),
	4256:  uint16(1),
	4257:  uint16(sym_comment),
	4258:  uint16(7),
	4259:  uint16(1),
	4260:  uint16(anon_sym_SLASH_SLASH_SLASH),
	4261:  uint16(80),
	4262:  uint16(1),
	4263:  uint16(sym_documentation_comment),
	4264:  uint16(331),
	4265:  uint16(27),
	4267:  uint16(anon_sym_use),
	4268:  uint16(anon_sym_enum),
	4269:  uint16(anon_sym_intEnum),
	4270:  uint16(anon_sym_list),
	4271:  uint16(anon_sym_map),
	4272:  uint16(anon_sym_set),
	4273:  uint16(anon_sym_structure),
	4274:  uint16(anon_sym_union),
	4275:  uint16(anon_sym_service),
	4276:  uint16(anon_sym_operation),
	4277:  uint16(anon_sym_resource),
	4278:  uint16(anon_sym_AT),
	4279:  uint16(anon_sym_apply),
	4280:  uint16(anon_sym_blob),
	4281:  uint16(anon_sym_boolean),
	4282:  uint16(anon_sym_byte),
	4283:  uint16(anon_sym_document),
	4284:  uint16(anon_sym_double),
	4285:  uint16(anon_sym_float),
	4286:  uint16(anon_sym_integer),
	4287:  uint16(anon_sym_long),
	4288:  uint16(anon_sym_short),
	4289:  uint16(anon_sym_string),
	4290:  uint16(anon_sym_timestamp),
	4291:  uint16(anon_sym_bigInteger),
	4292:  uint16(anon_sym_bigDecimal),
	4293:  uint16(9),
	4294:  uint16(3),
	4295:  uint16(1),
	4296:  uint16(anon_sym_COMMA),
	4297:  uint16(5),
	4298:  uint16(1),
	4299:  uint16(sym_comment),
	4300:  uint16(7),
	4301:  uint16(1),
	4302:  uint16(anon_sym_SLASH_SLASH_SLASH),
	4303:  uint16(333),
	4304:  uint16(1),
	4305:  uint16(anon_sym_DOLLAR),
	4306:  uint16(81),
	4307:  uint16(1),
	4308:  uint16(sym_documentation_comment),
	4309:  uint16(84),
	4310:  uint16(1),
	4311:  uint16(aux_sym_shape_id_repeat1),
	4312:  uint16(87),
	4313:  uint16(1),
	4314:  uint16(sym_shape_id_member),
	4315:  uint16(103),
	4316:  uint16(5),
	4317:  uint16(anon_sym_LBRACE),
	4318:  uint16(anon_sym_LBRACK),
	4319:  uint16(anon_sym_RBRACK),
	4320:  uint16(anon_sym_DASH),
	4321:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4322:  uint16(335),
	4323:  uint16(7),
	4324:  uint16(anon_sym_true),
	4325:  uint16(anon_sym_false),
	4326:  uint16(sym_null),
	4327:  uint16(aux_sym_number_token1),
	4328:  uint16(aux_sym_float_token1),
	4329:  uint16(anon_sym_DQUOTE),
	4330:  uint16(aux_sym_identifier_token1),
	4331:  uint16(8),
	4332:  uint16(3),
	4333:  uint16(1),
	4334:  uint16(anon_sym_COMMA),
	4335:  uint16(5),
	4336:  uint16(1),
	4337:  uint16(sym_comment),
	4338:  uint16(7),
	4339:  uint16(1),
	4340:  uint16(anon_sym_SLASH_SLASH_SLASH),
	4341:  uint16(337),
	4342:  uint16(1),
	4343:  uint16(anon_sym_DOLLAR),
	4344:  uint16(87),
	4345:  uint16(1),
	4346:  uint16(sym_shape_id_member),
	4347:  uint16(82),
	4348:  uint16(2),
	4349:  uint16(sym_documentation_comment),
	4350:  uint16(aux_sym_shape_id_repeat1),
	4351:  uint16(94),
	4352:  uint16(5),
	4353:  uint16(anon_sym_LBRACE),
	4354:  uint16(anon_sym_LBRACK),
	4355:  uint16(anon_sym_RBRACK),
	4356:  uint16(anon_sym_DASH),
	4357:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4358:  uint16(340),
	4359:  uint16(7),
	4360:  uint16(anon_sym_true),
	4361:  uint16(anon_sym_false),
	4362:  uint16(sym_null),
	4363:  uint16(aux_sym_number_token1),
	4364:  uint16(aux_sym_float_token1),
	4365:  uint16(anon_sym_DQUOTE),
	4366:  uint16(aux_sym_identifier_token1),
	4367:  uint16(8),
	4368:  uint16(3),
	4369:  uint16(1),
	4370:  uint16(anon_sym_COMMA),
	4371:  uint16(5),
	4372:  uint16(1),
	4373:  uint16(sym_comment),
	4374:  uint16(7),
	4375:  uint16(1),
	4376:  uint16(anon_sym_SLASH_SLASH_SLASH),
	4377:  uint16(90),
	4378:  uint16(1),
	4379:  uint16(anon_sym_POUND),
	4380:  uint16(342),
	4381:  uint16(1),
	4382:  uint16(anon_sym_DOT),
	4383:  uint16(83),
	4384:  uint16(1),
	4385:  uint16(sym_documentation_comment),
	4386:  uint16(88),
	4387:  uint16(6),
	4388:  uint16(anon_sym_DOLLAR),
	4389:  uint16(anon_sym_LBRACE),
	4390:  uint16(anon_sym_LBRACK),
	4391:  uint16(anon_sym_RBRACK),
	4392:  uint16(anon_sym_DASH),
	4393:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4394:  uint16(92),
	4395:  uint16(7),
	4396:  uint16(anon_sym_true),
	4397:  uint16(anon_sym_false),
	4398:  uint16(sym_null),
	4399:  uint16(aux_sym_number_token1),
	4400:  uint16(aux_sym_float_token1),
	4401:  uint16(anon_sym_DQUOTE),
	4402:  uint16(aux_sym_identifier_token1),
	4403:  uint16(9),
	4404:  uint16(3),
	4405:  uint16(1),
	4406:  uint16(anon_sym_COMMA),
	4407:  uint16(5),
	4408:  uint16(1),
	4409:  uint16(sym_comment),
	4410:  uint16(7),
	4411:  uint16(1),
	4412:  uint16(anon_sym_SLASH_SLASH_SLASH),
	4413:  uint16(333),
	4414:  uint16(1),
	4415:  uint16(anon_sym_DOLLAR),
	4416:  uint16(82),
	4417:  uint16(1),
	4418:  uint16(aux_sym_shape_id_repeat1),
	4419:  uint16(84),
	4420:  uint16(1),
	4421:  uint16(sym_documentation_comment),
	4422:  uint16(87),
	4423:  uint16(1),
	4424:  uint16(sym_shape_id_member),
	4425:  uint16(99),
	4426:  uint16(5),
	4427:  uint16(anon_sym_LBRACE),
	4428:  uint16(anon_sym_LBRACK),
	4429:  uint16(anon_sym_RBRACK),
	4430:  uint16(anon_sym_DASH),
	4431:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4432:  uint16(344),
	4433:  uint16(7),
	4434:  uint16(anon_sym_true),
	4435:  uint16(anon_sym_false),
	4436:  uint16(sym_null),
	4437:  uint16(aux_sym_number_token1),
	4438:  uint16(aux_sym_float_token1),
	4439:  uint16(anon_sym_DQUOTE),
	4440:  uint16(aux_sym_identifier_token1),
	4441:  uint16(6),
	4442:  uint16(3),
	4443:  uint16(1),
	4444:  uint16(anon_sym_COMMA),
	4445:  uint16(5),
	4446:  uint16(1),
	4447:  uint16(sym_comment),
	4448:  uint16(7),
	4449:  uint16(1),
	4450:  uint16(anon_sym_SLASH_SLASH_SLASH),
	4451:  uint16(85),
	4452:  uint16(1),
	4453:  uint16(sym_documentation_comment),
	4454:  uint16(88),
	4455:  uint16(6),
	4456:  uint16(anon_sym_DOLLAR),
	4457:  uint16(anon_sym_LBRACE),
	4458:  uint16(anon_sym_LBRACK),
	4459:  uint16(anon_sym_RBRACK),
	4460:  uint16(anon_sym_DASH),
	4461:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4462:  uint16(92),
	4463:  uint16(7),
	4464:  uint16(anon_sym_true),
	4465:  uint16(anon_sym_false),
	4466:  uint16(sym_null),
	4467:  uint16(aux_sym_number_token1),
	4468:  uint16(aux_sym_float_token1),
	4469:  uint16(anon_sym_DQUOTE),
	4470:  uint16(aux_sym_identifier_token1),
	4471:  uint16(6),
	4472:  uint16(3),
	4473:  uint16(1),
	4474:  uint16(anon_sym_COMMA),
	4475:  uint16(5),
	4476:  uint16(1),
	4477:  uint16(sym_comment),
	4478:  uint16(7),
	4479:  uint16(1),
	4480:  uint16(anon_sym_SLASH_SLASH_SLASH),
	4481:  uint16(86),
	4482:  uint16(1),
	4483:  uint16(sym_documentation_comment),
	4484:  uint16(105),
	4485:  uint16(6),
	4486:  uint16(anon_sym_DOLLAR),
	4487:  uint16(anon_sym_LBRACE),
	4488:  uint16(anon_sym_LBRACK),
	4489:  uint16(anon_sym_RBRACK),
	4490:  uint16(anon_sym_DASH),
	4491:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4492:  uint16(346),
	4493:  uint16(7),
	4494:  uint16(anon_sym_true),
	4495:  uint16(anon_sym_false),
	4496:  uint16(sym_null),
	4497:  uint16(aux_sym_number_token1),
	4498:  uint16(aux_sym_float_token1),
	4499:  uint16(anon_sym_DQUOTE),
	4500:  uint16(aux_sym_identifier_token1),
	4501:  uint16(6),
	4502:  uint16(3),
	4503:  uint16(1),
	4504:  uint16(anon_sym_COMMA),
	4505:  uint16(5),
	4506:  uint16(1),
	4507:  uint16(sym_comment),
	4508:  uint16(7),
	4509:  uint16(1),
	4510:  uint16(anon_sym_SLASH_SLASH_SLASH),
	4511:  uint16(87),
	4512:  uint16(1),
	4513:  uint16(sym_documentation_comment),
	4514:  uint16(107),
	4515:  uint16(6),
	4516:  uint16(anon_sym_DOLLAR),
	4517:  uint16(anon_sym_LBRACE),
	4518:  uint16(anon_sym_LBRACK),
	4519:  uint16(anon_sym_RBRACK),
	4520:  uint16(anon_sym_DASH),
	4521:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4522:  uint16(348),
	4523:  uint16(7),
	4524:  uint16(anon_sym_true),
	4525:  uint16(anon_sym_false),
	4526:  uint16(sym_null),
	4527:  uint16(aux_sym_number_token1),
	4528:  uint16(aux_sym_float_token1),
	4529:  uint16(anon_sym_DQUOTE),
	4530:  uint16(aux_sym_identifier_token1),
	4531:  uint16(6),
	4532:  uint16(3),
	4533:  uint16(1),
	4534:  uint16(anon_sym_COMMA),
	4535:  uint16(5),
	4536:  uint16(1),
	4537:  uint16(sym_comment),
	4538:  uint16(7),
	4539:  uint16(1),
	4540:  uint16(anon_sym_SLASH_SLASH_SLASH),
	4541:  uint16(88),
	4542:  uint16(1),
	4543:  uint16(sym_documentation_comment),
	4544:  uint16(109),
	4545:  uint16(6),
	4546:  uint16(anon_sym_DOLLAR),
	4547:  uint16(anon_sym_LBRACE),
	4548:  uint16(anon_sym_LBRACK),
	4549:  uint16(anon_sym_RBRACK),
	4550:  uint16(anon_sym_DASH),
	4551:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4552:  uint16(350),
	4553:  uint16(7),
	4554:  uint16(anon_sym_true),
	4555:  uint16(anon_sym_false),
	4556:  uint16(sym_null),
	4557:  uint16(aux_sym_number_token1),
	4558:  uint16(aux_sym_float_token1),
	4559:  uint16(anon_sym_DQUOTE),
	4560:  uint16(aux_sym_identifier_token1),
	4561:  uint16(6),
	4562:  uint16(3),
	4563:  uint16(1),
	4564:  uint16(anon_sym_COMMA),
	4565:  uint16(5),
	4566:  uint16(1),
	4567:  uint16(sym_comment),
	4568:  uint16(7),
	4569:  uint16(1),
	4570:  uint16(anon_sym_SLASH_SLASH_SLASH),
	4571:  uint16(89),
	4572:  uint16(1),
	4573:  uint16(sym_documentation_comment),
	4574:  uint16(111),
	4575:  uint16(6),
	4576:  uint16(anon_sym_DOLLAR),
	4577:  uint16(anon_sym_LBRACE),
	4578:  uint16(anon_sym_LBRACK),
	4579:  uint16(anon_sym_RBRACK),
	4580:  uint16(anon_sym_DASH),
	4581:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4582:  uint16(352),
	4583:  uint16(7),
	4584:  uint16(anon_sym_true),
	4585:  uint16(anon_sym_false),
	4586:  uint16(sym_null),
	4587:  uint16(aux_sym_number_token1),
	4588:  uint16(aux_sym_float_token1),
	4589:  uint16(anon_sym_DQUOTE),
	4590:  uint16(aux_sym_identifier_token1),
	4591:  uint16(6),
	4592:  uint16(3),
	4593:  uint16(1),
	4594:  uint16(anon_sym_COMMA),
	4595:  uint16(5),
	4596:  uint16(1),
	4597:  uint16(sym_comment),
	4598:  uint16(7),
	4599:  uint16(1),
	4600:  uint16(anon_sym_SLASH_SLASH_SLASH),
	4601:  uint16(90),
	4602:  uint16(1),
	4603:  uint16(sym_documentation_comment),
	4604:  uint16(354),
	4605:  uint16(5),
	4606:  uint16(anon_sym_LBRACE),
	4607:  uint16(anon_sym_LBRACK),
	4608:  uint16(anon_sym_RBRACK),
	4609:  uint16(anon_sym_DASH),
	4610:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4611:  uint16(356),
	4612:  uint16(7),
	4613:  uint16(anon_sym_true),
	4614:  uint16(anon_sym_false),
	4615:  uint16(sym_null),
	4616:  uint16(aux_sym_number_token1),
	4617:  uint16(aux_sym_float_token1),
	4618:  uint16(anon_sym_DQUOTE),
	4619:  uint16(aux_sym_identifier_token1),
	4620:  uint16(6),
	4621:  uint16(3),
	4622:  uint16(1),
	4623:  uint16(anon_sym_COMMA),
	4624:  uint16(5),
	4625:  uint16(1),
	4626:  uint16(sym_comment),
	4627:  uint16(7),
	4628:  uint16(1),
	4629:  uint16(anon_sym_SLASH_SLASH_SLASH),
	4630:  uint16(91),
	4631:  uint16(1),
	4632:  uint16(sym_documentation_comment),
	4633:  uint16(358),
	4634:  uint16(5),
	4635:  uint16(anon_sym_LBRACE),
	4636:  uint16(anon_sym_LBRACK),
	4637:  uint16(anon_sym_RBRACK),
	4638:  uint16(anon_sym_DASH),
	4639:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4640:  uint16(360),
	4641:  uint16(7),
	4642:  uint16(anon_sym_true),
	4643:  uint16(anon_sym_false),
	4644:  uint16(sym_null),
	4645:  uint16(aux_sym_number_token1),
	4646:  uint16(aux_sym_float_token1),
	4647:  uint16(anon_sym_DQUOTE),
	4648:  uint16(aux_sym_identifier_token1),
	4649:  uint16(6),
	4650:  uint16(3),
	4651:  uint16(1),
	4652:  uint16(anon_sym_COMMA),
	4653:  uint16(5),
	4654:  uint16(1),
	4655:  uint16(sym_comment),
	4656:  uint16(7),
	4657:  uint16(1),
	4658:  uint16(anon_sym_SLASH_SLASH_SLASH),
	4659:  uint16(92),
	4660:  uint16(1),
	4661:  uint16(sym_documentation_comment),
	4662:  uint16(362),
	4663:  uint16(5),
	4664:  uint16(anon_sym_LBRACE),
	4665:  uint16(anon_sym_LBRACK),
	4666:  uint16(anon_sym_RBRACK),
	4667:  uint16(anon_sym_DASH),
	4668:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4669:  uint16(364),
	4670:  uint16(7),
	4671:  uint16(anon_sym_true),
	4672:  uint16(anon_sym_false),
	4673:  uint16(sym_null),
	4674:  uint16(aux_sym_number_token1),
	4675:  uint16(aux_sym_float_token1),
	4676:  uint16(anon_sym_DQUOTE),
	4677:  uint16(aux_sym_identifier_token1),
	4678:  uint16(6),
	4679:  uint16(3),
	4680:  uint16(1),
	4681:  uint16(anon_sym_COMMA),
	4682:  uint16(5),
	4683:  uint16(1),
	4684:  uint16(sym_comment),
	4685:  uint16(7),
	4686:  uint16(1),
	4687:  uint16(anon_sym_SLASH_SLASH_SLASH),
	4688:  uint16(93),
	4689:  uint16(1),
	4690:  uint16(sym_documentation_comment),
	4691:  uint16(366),
	4692:  uint16(5),
	4693:  uint16(anon_sym_LBRACE),
	4694:  uint16(anon_sym_LBRACK),
	4695:  uint16(anon_sym_RBRACK),
	4696:  uint16(anon_sym_DASH),
	4697:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4698:  uint16(368),
	4699:  uint16(7),
	4700:  uint16(anon_sym_true),
	4701:  uint16(anon_sym_false),
	4702:  uint16(sym_null),
	4703:  uint16(aux_sym_number_token1),
	4704:  uint16(aux_sym_float_token1),
	4705:  uint16(anon_sym_DQUOTE),
	4706:  uint16(aux_sym_identifier_token1),
	4707:  uint16(8),
	4708:  uint16(3),
	4709:  uint16(1),
	4710:  uint16(anon_sym_COMMA),
	4711:  uint16(5),
	4712:  uint16(1),
	4713:  uint16(sym_comment),
	4714:  uint16(7),
	4715:  uint16(1),
	4716:  uint16(anon_sym_SLASH_SLASH_SLASH),
	4717:  uint16(340),
	4718:  uint16(1),
	4719:  uint16(anon_sym_DQUOTE),
	4720:  uint16(370),
	4721:  uint16(1),
	4722:  uint16(anon_sym_DOLLAR),
	4723:  uint16(121),
	4724:  uint16(1),
	4725:  uint16(sym_shape_id_member),
	4726:  uint16(94),
	4727:  uint16(2),
	4728:  uint16(sym_documentation_comment),
	4729:  uint16(aux_sym_shape_id_repeat1),
	4730:  uint16(94),
	4731:  uint16(8),
	4732:  uint16(anon_sym_EQ),
	4733:  uint16(anon_sym_RBRACE),
	4734:  uint16(anon_sym_RBRACK),
	4735:  uint16(anon_sym_AT),
	4736:  uint16(anon_sym_LPAREN),
	4737:  uint16(anon_sym_RPAREN),
	4738:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4739:  uint16(aux_sym_identifier_token1),
	4740:  uint16(6),
	4741:  uint16(3),
	4742:  uint16(1),
	4743:  uint16(anon_sym_COMMA),
	4744:  uint16(5),
	4745:  uint16(1),
	4746:  uint16(sym_comment),
	4747:  uint16(7),
	4748:  uint16(1),
	4749:  uint16(anon_sym_SLASH_SLASH_SLASH),
	4750:  uint16(95),
	4751:  uint16(1),
	4752:  uint16(sym_documentation_comment),
	4753:  uint16(373),
	4754:  uint16(5),
	4755:  uint16(anon_sym_LBRACE),
	4756:  uint16(anon_sym_LBRACK),
	4757:  uint16(anon_sym_RBRACK),
	4758:  uint16(anon_sym_DASH),
	4759:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4760:  uint16(375),
	4761:  uint16(7),
	4762:  uint16(anon_sym_true),
	4763:  uint16(anon_sym_false),
	4764:  uint16(sym_null),
	4765:  uint16(aux_sym_number_token1),
	4766:  uint16(aux_sym_float_token1),
	4767:  uint16(anon_sym_DQUOTE),
	4768:  uint16(aux_sym_identifier_token1),
	4769:  uint16(6),
	4770:  uint16(3),
	4771:  uint16(1),
	4772:  uint16(anon_sym_COMMA),
	4773:  uint16(5),
	4774:  uint16(1),
	4775:  uint16(sym_comment),
	4776:  uint16(7),
	4777:  uint16(1),
	4778:  uint16(anon_sym_SLASH_SLASH_SLASH),
	4779:  uint16(96),
	4780:  uint16(1),
	4781:  uint16(sym_documentation_comment),
	4782:  uint16(377),
	4783:  uint16(5),
	4784:  uint16(anon_sym_LBRACE),
	4785:  uint16(anon_sym_LBRACK),
	4786:  uint16(anon_sym_RBRACK),
	4787:  uint16(anon_sym_DASH),
	4788:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4789:  uint16(379),
	4790:  uint16(7),
	4791:  uint16(anon_sym_true),
	4792:  uint16(anon_sym_false),
	4793:  uint16(sym_null),
	4794:  uint16(aux_sym_number_token1),
	4795:  uint16(aux_sym_float_token1),
	4796:  uint16(anon_sym_DQUOTE),
	4797:  uint16(aux_sym_identifier_token1),
	4798:  uint16(6),
	4799:  uint16(3),
	4800:  uint16(1),
	4801:  uint16(anon_sym_COMMA),
	4802:  uint16(5),
	4803:  uint16(1),
	4804:  uint16(sym_comment),
	4805:  uint16(7),
	4806:  uint16(1),
	4807:  uint16(anon_sym_SLASH_SLASH_SLASH),
	4808:  uint16(97),
	4809:  uint16(1),
	4810:  uint16(sym_documentation_comment),
	4811:  uint16(381),
	4812:  uint16(5),
	4813:  uint16(anon_sym_LBRACE),
	4814:  uint16(anon_sym_LBRACK),
	4815:  uint16(anon_sym_RBRACK),
	4816:  uint16(anon_sym_DASH),
	4817:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4818:  uint16(383),
	4819:  uint16(7),
	4820:  uint16(anon_sym_true),
	4821:  uint16(anon_sym_false),
	4822:  uint16(sym_null),
	4823:  uint16(aux_sym_number_token1),
	4824:  uint16(aux_sym_float_token1),
	4825:  uint16(anon_sym_DQUOTE),
	4826:  uint16(aux_sym_identifier_token1),
	4827:  uint16(6),
	4828:  uint16(3),
	4829:  uint16(1),
	4830:  uint16(anon_sym_COMMA),
	4831:  uint16(5),
	4832:  uint16(1),
	4833:  uint16(sym_comment),
	4834:  uint16(7),
	4835:  uint16(1),
	4836:  uint16(anon_sym_SLASH_SLASH_SLASH),
	4837:  uint16(98),
	4838:  uint16(1),
	4839:  uint16(sym_documentation_comment),
	4840:  uint16(385),
	4841:  uint16(5),
	4842:  uint16(anon_sym_LBRACE),
	4843:  uint16(anon_sym_LBRACK),
	4844:  uint16(anon_sym_RBRACK),
	4845:  uint16(anon_sym_DASH),
	4846:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4847:  uint16(387),
	4848:  uint16(7),
	4849:  uint16(anon_sym_true),
	4850:  uint16(anon_sym_false),
	4851:  uint16(sym_null),
	4852:  uint16(aux_sym_number_token1),
	4853:  uint16(aux_sym_float_token1),
	4854:  uint16(anon_sym_DQUOTE),
	4855:  uint16(aux_sym_identifier_token1),
	4856:  uint16(6),
	4857:  uint16(3),
	4858:  uint16(1),
	4859:  uint16(anon_sym_COMMA),
	4860:  uint16(5),
	4861:  uint16(1),
	4862:  uint16(sym_comment),
	4863:  uint16(7),
	4864:  uint16(1),
	4865:  uint16(anon_sym_SLASH_SLASH_SLASH),
	4866:  uint16(99),
	4867:  uint16(1),
	4868:  uint16(sym_documentation_comment),
	4869:  uint16(143),
	4870:  uint16(5),
	4871:  uint16(anon_sym_LBRACE),
	4872:  uint16(anon_sym_LBRACK),
	4873:  uint16(anon_sym_RBRACK),
	4874:  uint16(anon_sym_DASH),
	4875:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4876:  uint16(389),
	4877:  uint16(7),
	4878:  uint16(anon_sym_true),
	4879:  uint16(anon_sym_false),
	4880:  uint16(sym_null),
	4881:  uint16(aux_sym_number_token1),
	4882:  uint16(aux_sym_float_token1),
	4883:  uint16(anon_sym_DQUOTE),
	4884:  uint16(aux_sym_identifier_token1),
	4885:  uint16(7),
	4886:  uint16(3),
	4887:  uint16(1),
	4888:  uint16(anon_sym_COMMA),
	4889:  uint16(5),
	4890:  uint16(1),
	4891:  uint16(sym_comment),
	4892:  uint16(7),
	4893:  uint16(1),
	4894:  uint16(anon_sym_SLASH_SLASH_SLASH),
	4895:  uint16(92),
	4896:  uint16(1),
	4897:  uint16(anon_sym_DQUOTE),
	4898:  uint16(100),
	4899:  uint16(1),
	4900:  uint16(sym_documentation_comment),
	4901:  uint16(90),
	4902:  uint16(2),
	4903:  uint16(anon_sym_DOT),
	4904:  uint16(anon_sym_POUND),
	4905:  uint16(88),
	4906:  uint16(9),
	4907:  uint16(anon_sym_DOLLAR),
	4908:  uint16(anon_sym_EQ),
	4909:  uint16(anon_sym_RBRACE),
	4910:  uint16(anon_sym_RBRACK),
	4911:  uint16(anon_sym_AT),
	4912:  uint16(anon_sym_LPAREN),
	4913:  uint16(anon_sym_RPAREN),
	4914:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4915:  uint16(aux_sym_identifier_token1),
	4916:  uint16(6),
	4917:  uint16(3),
	4918:  uint16(1),
	4919:  uint16(anon_sym_COMMA),
	4920:  uint16(5),
	4921:  uint16(1),
	4922:  uint16(sym_comment),
	4923:  uint16(7),
	4924:  uint16(1),
	4925:  uint16(anon_sym_SLASH_SLASH_SLASH),
	4926:  uint16(101),
	4927:  uint16(1),
	4928:  uint16(sym_documentation_comment),
	4929:  uint16(391),
	4930:  uint16(5),
	4931:  uint16(anon_sym_LBRACE),
	4932:  uint16(anon_sym_LBRACK),
	4933:  uint16(anon_sym_RBRACK),
	4934:  uint16(anon_sym_DASH),
	4935:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4936:  uint16(393),
	4937:  uint16(7),
	4938:  uint16(anon_sym_true),
	4939:  uint16(anon_sym_false),
	4940:  uint16(sym_null),
	4941:  uint16(aux_sym_number_token1),
	4942:  uint16(aux_sym_float_token1),
	4943:  uint16(anon_sym_DQUOTE),
	4944:  uint16(aux_sym_identifier_token1),
	4945:  uint16(6),
	4946:  uint16(3),
	4947:  uint16(1),
	4948:  uint16(anon_sym_COMMA),
	4949:  uint16(5),
	4950:  uint16(1),
	4951:  uint16(sym_comment),
	4952:  uint16(7),
	4953:  uint16(1),
	4954:  uint16(anon_sym_SLASH_SLASH_SLASH),
	4955:  uint16(102),
	4956:  uint16(1),
	4957:  uint16(sym_documentation_comment),
	4958:  uint16(395),
	4959:  uint16(5),
	4960:  uint16(anon_sym_LBRACE),
	4961:  uint16(anon_sym_LBRACK),
	4962:  uint16(anon_sym_RBRACK),
	4963:  uint16(anon_sym_DASH),
	4964:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4965:  uint16(397),
	4966:  uint16(7),
	4967:  uint16(anon_sym_true),
	4968:  uint16(anon_sym_false),
	4969:  uint16(sym_null),
	4970:  uint16(aux_sym_number_token1),
	4971:  uint16(aux_sym_float_token1),
	4972:  uint16(anon_sym_DQUOTE),
	4973:  uint16(aux_sym_identifier_token1),
	4974:  uint16(6),
	4975:  uint16(3),
	4976:  uint16(1),
	4977:  uint16(anon_sym_COMMA),
	4978:  uint16(5),
	4979:  uint16(1),
	4980:  uint16(sym_comment),
	4981:  uint16(7),
	4982:  uint16(1),
	4983:  uint16(anon_sym_SLASH_SLASH_SLASH),
	4984:  uint16(103),
	4985:  uint16(1),
	4986:  uint16(sym_documentation_comment),
	4987:  uint16(399),
	4988:  uint16(5),
	4989:  uint16(anon_sym_LBRACE),
	4990:  uint16(anon_sym_LBRACK),
	4991:  uint16(anon_sym_RBRACK),
	4992:  uint16(anon_sym_DASH),
	4993:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4994:  uint16(401),
	4995:  uint16(7),
	4996:  uint16(anon_sym_true),
	4997:  uint16(anon_sym_false),
	4998:  uint16(sym_null),
	4999:  uint16(aux_sym_number_token1),
	5000:  uint16(aux_sym_float_token1),
	5001:  uint16(anon_sym_DQUOTE),
	5002:  uint16(aux_sym_identifier_token1),
	5003:  uint16(6),
	5004:  uint16(3),
	5005:  uint16(1),
	5006:  uint16(anon_sym_COMMA),
	5007:  uint16(5),
	5008:  uint16(1),
	5009:  uint16(sym_comment),
	5010:  uint16(7),
	5011:  uint16(1),
	5012:  uint16(anon_sym_SLASH_SLASH_SLASH),
	5013:  uint16(104),
	5014:  uint16(1),
	5015:  uint16(sym_documentation_comment),
	5016:  uint16(403),
	5017:  uint16(5),
	5018:  uint16(anon_sym_LBRACE),
	5019:  uint16(anon_sym_LBRACK),
	5020:  uint16(anon_sym_RBRACK),
	5021:  uint16(anon_sym_DASH),
	5022:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5023:  uint16(405),
	5024:  uint16(7),
	5025:  uint16(anon_sym_true),
	5026:  uint16(anon_sym_false),
	5027:  uint16(sym_null),
	5028:  uint16(aux_sym_number_token1),
	5029:  uint16(aux_sym_float_token1),
	5030:  uint16(anon_sym_DQUOTE),
	5031:  uint16(aux_sym_identifier_token1),
	5032:  uint16(6),
	5033:  uint16(3),
	5034:  uint16(1),
	5035:  uint16(anon_sym_COMMA),
	5036:  uint16(5),
	5037:  uint16(1),
	5038:  uint16(sym_comment),
	5039:  uint16(7),
	5040:  uint16(1),
	5041:  uint16(anon_sym_SLASH_SLASH_SLASH),
	5042:  uint16(105),
	5043:  uint16(1),
	5044:  uint16(sym_documentation_comment),
	5045:  uint16(141),
	5046:  uint16(5),
	5047:  uint16(anon_sym_LBRACE),
	5048:  uint16(anon_sym_LBRACK),
	5049:  uint16(anon_sym_RBRACK),
	5050:  uint16(anon_sym_DASH),
	5051:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5052:  uint16(407),
	5053:  uint16(7),
	5054:  uint16(anon_sym_true),
	5055:  uint16(anon_sym_false),
	5056:  uint16(sym_null),
	5057:  uint16(aux_sym_number_token1),
	5058:  uint16(aux_sym_float_token1),
	5059:  uint16(anon_sym_DQUOTE),
	5060:  uint16(aux_sym_identifier_token1),
	5061:  uint16(6),
	5062:  uint16(3),
	5063:  uint16(1),
	5064:  uint16(anon_sym_COMMA),
	5065:  uint16(5),
	5066:  uint16(1),
	5067:  uint16(sym_comment),
	5068:  uint16(7),
	5069:  uint16(1),
	5070:  uint16(anon_sym_SLASH_SLASH_SLASH),
	5071:  uint16(106),
	5072:  uint16(1),
	5073:  uint16(sym_documentation_comment),
	5074:  uint16(409),
	5075:  uint16(5),
	5076:  uint16(anon_sym_LBRACE),
	5077:  uint16(anon_sym_LBRACK),
	5078:  uint16(anon_sym_RBRACK),
	5079:  uint16(anon_sym_DASH),
	5080:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5081:  uint16(411),
	5082:  uint16(7),
	5083:  uint16(anon_sym_true),
	5084:  uint16(anon_sym_false),
	5085:  uint16(sym_null),
	5086:  uint16(aux_sym_number_token1),
	5087:  uint16(aux_sym_float_token1),
	5088:  uint16(anon_sym_DQUOTE),
	5089:  uint16(aux_sym_identifier_token1),
	5090:  uint16(6),
	5091:  uint16(3),
	5092:  uint16(1),
	5093:  uint16(anon_sym_COMMA),
	5094:  uint16(5),
	5095:  uint16(1),
	5096:  uint16(sym_comment),
	5097:  uint16(7),
	5098:  uint16(1),
	5099:  uint16(anon_sym_SLASH_SLASH_SLASH),
	5100:  uint16(107),
	5101:  uint16(1),
	5102:  uint16(sym_documentation_comment),
	5103:  uint16(413),
	5104:  uint16(5),
	5105:  uint16(anon_sym_LBRACE),
	5106:  uint16(anon_sym_LBRACK),
	5107:  uint16(anon_sym_RBRACK),
	5108:  uint16(anon_sym_DASH),
	5109:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5110:  uint16(415),
	5111:  uint16(7),
	5112:  uint16(anon_sym_true),
	5113:  uint16(anon_sym_false),
	5114:  uint16(sym_null),
	5115:  uint16(aux_sym_number_token1),
	5116:  uint16(aux_sym_float_token1),
	5117:  uint16(anon_sym_DQUOTE),
	5118:  uint16(aux_sym_identifier_token1),
	5119:  uint16(6),
	5120:  uint16(3),
	5121:  uint16(1),
	5122:  uint16(anon_sym_COMMA),
	5123:  uint16(5),
	5124:  uint16(1),
	5125:  uint16(sym_comment),
	5126:  uint16(7),
	5127:  uint16(1),
	5128:  uint16(anon_sym_SLASH_SLASH_SLASH),
	5129:  uint16(108),
	5130:  uint16(1),
	5131:  uint16(sym_documentation_comment),
	5132:  uint16(417),
	5133:  uint16(5),
	5134:  uint16(anon_sym_LBRACE),
	5135:  uint16(anon_sym_LBRACK),
	5136:  uint16(anon_sym_RBRACK),
	5137:  uint16(anon_sym_DASH),
	5138:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5139:  uint16(419),
	5140:  uint16(7),
	5141:  uint16(anon_sym_true),
	5142:  uint16(anon_sym_false),
	5143:  uint16(sym_null),
	5144:  uint16(aux_sym_number_token1),
	5145:  uint16(aux_sym_float_token1),
	5146:  uint16(anon_sym_DQUOTE),
	5147:  uint16(aux_sym_identifier_token1),
	5148:  uint16(9),
	5149:  uint16(3),
	5150:  uint16(1),
	5151:  uint16(anon_sym_COMMA),
	5152:  uint16(5),
	5153:  uint16(1),
	5154:  uint16(sym_comment),
	5155:  uint16(7),
	5156:  uint16(1),
	5157:  uint16(anon_sym_SLASH_SLASH_SLASH),
	5158:  uint16(344),
	5159:  uint16(1),
	5160:  uint16(anon_sym_DQUOTE),
	5161:  uint16(421),
	5162:  uint16(1),
	5163:  uint16(anon_sym_DOLLAR),
	5164:  uint16(94),
	5165:  uint16(1),
	5166:  uint16(aux_sym_shape_id_repeat1),
	5167:  uint16(109),
	5168:  uint16(1),
	5169:  uint16(sym_documentation_comment),
	5170:  uint16(121),
	5171:  uint16(1),
	5172:  uint16(sym_shape_id_member),
	5173:  uint16(99),
	5174:  uint16(7),
	5175:  uint16(anon_sym_RBRACE),
	5176:  uint16(anon_sym_RBRACK),
	5177:  uint16(anon_sym_AT),
	5178:  uint16(anon_sym_LPAREN),
	5179:  uint16(anon_sym_RPAREN),
	5180:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5181:  uint16(aux_sym_identifier_token1),
	5182:  uint16(14),
	5183:  uint16(3),
	5184:  uint16(1),
	5185:  uint16(anon_sym_COMMA),
	5186:  uint16(5),
	5187:  uint16(1),
	5188:  uint16(sym_comment),
	5189:  uint16(7),
	5190:  uint16(1),
	5191:  uint16(anon_sym_SLASH_SLASH_SLASH),
	5192:  uint16(133),
	5193:  uint16(1),
	5194:  uint16(anon_sym_DQUOTE),
	5195:  uint16(135),
	5196:  uint16(1),
	5197:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5198:  uint16(423),
	5199:  uint16(1),
	5200:  uint16(anon_sym_RBRACE),
	5201:  uint16(425),
	5202:  uint16(1),
	5203:  uint16(aux_sym_identifier_token1),
	5204:  uint16(110),
	5205:  uint16(1),
	5206:  uint16(sym_documentation_comment),
	5207:  uint16(117),
	5208:  uint16(1),
	5209:  uint16(aux_sym_node_object_repeat1),
	5210:  uint16(246),
	5211:  uint16(1),
	5212:  uint16(sym_node_object_kvp),
	5213:  uint16(340),
	5214:  uint16(1),
	5215:  uint16(sym_node_object_key),
	5216:  uint16(342),
	5217:  uint16(1),
	5218:  uint16(sym_string),
	5219:  uint16(343),
	5220:  uint16(1),
	5221:  uint16(sym_identifier),
	5222:  uint16(168),
	5223:  uint16(2),
	5224:  uint16(sym__string_literal),
	5225:  uint16(sym__multiline_string_literal),
	5226:  uint16(14),
	5227:  uint16(3),
	5228:  uint16(1),
	5229:  uint16(anon_sym_COMMA),
	5230:  uint16(5),
	5231:  uint16(1),
	5232:  uint16(sym_comment),
	5233:  uint16(7),
	5234:  uint16(1),
	5235:  uint16(anon_sym_SLASH_SLASH_SLASH),
	5236:  uint16(133),
	5237:  uint16(1),
	5238:  uint16(anon_sym_DQUOTE),
	5239:  uint16(135),
	5240:  uint16(1),
	5241:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5242:  uint16(425),
	5243:  uint16(1),
	5244:  uint16(aux_sym_identifier_token1),
	5245:  uint16(427),
	5246:  uint16(1),
	5247:  uint16(anon_sym_RBRACE),
	5248:  uint16(111),
	5249:  uint16(1),
	5250:  uint16(sym_documentation_comment),
	5251:  uint16(117),
	5252:  uint16(1),
	5253:  uint16(aux_sym_node_object_repeat1),
	5254:  uint16(246),
	5255:  uint16(1),
	5256:  uint16(sym_node_object_kvp),
	5257:  uint16(340),
	5258:  uint16(1),
	5259:  uint16(sym_node_object_key),
	5260:  uint16(342),
	5261:  uint16(1),
	5262:  uint16(sym_string),
	5263:  uint16(343),
	5264:  uint16(1),
	5265:  uint16(sym_identifier),
	5266:  uint16(168),
	5267:  uint16(2),
	5268:  uint16(sym__string_literal),
	5269:  uint16(sym__multiline_string_literal),
	5270:  uint16(14),
	5271:  uint16(3),
	5272:  uint16(1),
	5273:  uint16(anon_sym_COMMA),
	5274:  uint16(5),
	5275:  uint16(1),
	5276:  uint16(sym_comment),
	5277:  uint16(7),
	5278:  uint16(1),
	5279:  uint16(anon_sym_SLASH_SLASH_SLASH),
	5280:  uint16(133),
	5281:  uint16(1),
	5282:  uint16(anon_sym_DQUOTE),
	5283:  uint16(135),
	5284:  uint16(1),
	5285:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5286:  uint16(425),
	5287:  uint16(1),
	5288:  uint16(aux_sym_identifier_token1),
	5289:  uint16(429),
	5290:  uint16(1),
	5291:  uint16(anon_sym_RPAREN),
	5292:  uint16(112),
	5293:  uint16(1),
	5294:  uint16(sym_documentation_comment),
	5295:  uint16(116),
	5296:  uint16(1),
	5297:  uint16(aux_sym_trait_structure_repeat1),
	5298:  uint16(253),
	5299:  uint16(1),
	5300:  uint16(sym_node_object_kvp),
	5301:  uint16(340),
	5302:  uint16(1),
	5303:  uint16(sym_node_object_key),
	5304:  uint16(342),
	5305:  uint16(1),
	5306:  uint16(sym_string),
	5307:  uint16(343),
	5308:  uint16(1),
	5309:  uint16(sym_identifier),
	5310:  uint16(168),
	5311:  uint16(2),
	5312:  uint16(sym__string_literal),
	5313:  uint16(sym__multiline_string_literal),
	5314:  uint16(14),
	5315:  uint16(3),
	5316:  uint16(1),
	5317:  uint16(anon_sym_COMMA),
	5318:  uint16(5),
	5319:  uint16(1),
	5320:  uint16(sym_comment),
	5321:  uint16(7),
	5322:  uint16(1),
	5323:  uint16(anon_sym_SLASH_SLASH_SLASH),
	5324:  uint16(133),
	5325:  uint16(1),
	5326:  uint16(anon_sym_DQUOTE),
	5327:  uint16(135),
	5328:  uint16(1),
	5329:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5330:  uint16(425),
	5331:  uint16(1),
	5332:  uint16(aux_sym_identifier_token1),
	5333:  uint16(431),
	5334:  uint16(1),
	5335:  uint16(anon_sym_RBRACE),
	5336:  uint16(111),
	5337:  uint16(1),
	5338:  uint16(aux_sym_node_object_repeat1),
	5339:  uint16(113),
	5340:  uint16(1),
	5341:  uint16(sym_documentation_comment),
	5342:  uint16(246),
	5343:  uint16(1),
	5344:  uint16(sym_node_object_kvp),
	5345:  uint16(340),
	5346:  uint16(1),
	5347:  uint16(sym_node_object_key),
	5348:  uint16(342),
	5349:  uint16(1),
	5350:  uint16(sym_string),
	5351:  uint16(343),
	5352:  uint16(1),
	5353:  uint16(sym_identifier),
	5354:  uint16(168),
	5355:  uint16(2),
	5356:  uint16(sym__string_literal),
	5357:  uint16(sym__multiline_string_literal),
	5358:  uint16(6),
	5359:  uint16(3),
	5360:  uint16(1),
	5361:  uint16(anon_sym_COMMA),
	5362:  uint16(5),
	5363:  uint16(1),
	5364:  uint16(sym_comment),
	5365:  uint16(7),
	5366:  uint16(1),
	5367:  uint16(anon_sym_SLASH_SLASH_SLASH),
	5368:  uint16(114),
	5369:  uint16(1),
	5370:  uint16(sym_documentation_comment),
	5371:  uint16(433),
	5372:  uint16(4),
	5373:  uint16(anon_sym_LBRACE),
	5374:  uint16(anon_sym_LBRACK),
	5375:  uint16(anon_sym_DASH),
	5376:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5377:  uint16(435),
	5378:  uint16(7),
	5379:  uint16(anon_sym_true),
	5380:  uint16(anon_sym_false),
	5381:  uint16(sym_null),
	5382:  uint16(aux_sym_number_token1),
	5383:  uint16(aux_sym_float_token1),
	5384:  uint16(anon_sym_DQUOTE),
	5385:  uint16(aux_sym_identifier_token1),
	5386:  uint16(14),
	5387:  uint16(3),
	5388:  uint16(1),
	5389:  uint16(anon_sym_COMMA),
	5390:  uint16(5),
	5391:  uint16(1),
	5392:  uint16(sym_comment),
	5393:  uint16(7),
	5394:  uint16(1),
	5395:  uint16(anon_sym_SLASH_SLASH_SLASH),
	5396:  uint16(133),
	5397:  uint16(1),
	5398:  uint16(anon_sym_DQUOTE),
	5399:  uint16(135),
	5400:  uint16(1),
	5401:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5402:  uint16(425),
	5403:  uint16(1),
	5404:  uint16(aux_sym_identifier_token1),
	5405:  uint16(437),
	5406:  uint16(1),
	5407:  uint16(anon_sym_RBRACE),
	5408:  uint16(115),
	5409:  uint16(1),
	5410:  uint16(sym_documentation_comment),
	5411:  uint16(117),
	5412:  uint16(1),
	5413:  uint16(aux_sym_node_object_repeat1),
	5414:  uint16(246),
	5415:  uint16(1),
	5416:  uint16(sym_node_object_kvp),
	5417:  uint16(340),
	5418:  uint16(1),
	5419:  uint16(sym_node_object_key),
	5420:  uint16(342),
	5421:  uint16(1),
	5422:  uint16(sym_string),
	5423:  uint16(343),
	5424:  uint16(1),
	5425:  uint16(sym_identifier),
	5426:  uint16(168),
	5427:  uint16(2),
	5428:  uint16(sym__string_literal),
	5429:  uint16(sym__multiline_string_literal),
	5430:  uint16(13),
	5431:  uint16(3),
	5432:  uint16(1),
	5433:  uint16(anon_sym_COMMA),
	5434:  uint16(5),
	5435:  uint16(1),
	5436:  uint16(sym_comment),
	5437:  uint16(7),
	5438:  uint16(1),
	5439:  uint16(anon_sym_SLASH_SLASH_SLASH),
	5440:  uint16(439),
	5441:  uint16(1),
	5442:  uint16(anon_sym_RPAREN),
	5443:  uint16(441),
	5444:  uint16(1),
	5445:  uint16(anon_sym_DQUOTE),
	5446:  uint16(444),
	5447:  uint16(1),
	5448:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5449:  uint16(447),
	5450:  uint16(1),
	5451:  uint16(aux_sym_identifier_token1),
	5452:  uint16(253),
	5453:  uint16(1),
	5454:  uint16(sym_node_object_kvp),
	5455:  uint16(340),
	5456:  uint16(1),
	5457:  uint16(sym_node_object_key),
	5458:  uint16(342),
	5459:  uint16(1),
	5460:  uint16(sym_string),
	5461:  uint16(343),
	5462:  uint16(1),
	5463:  uint16(sym_identifier),
	5464:  uint16(116),
	5465:  uint16(2),
	5466:  uint16(sym_documentation_comment),
	5467:  uint16(aux_sym_trait_structure_repeat1),
	5468:  uint16(168),
	5469:  uint16(2),
	5470:  uint16(sym__string_literal),
	5471:  uint16(sym__multiline_string_literal),
	5472:  uint16(13),
	5473:  uint16(3),
	5474:  uint16(1),
	5475:  uint16(anon_sym_COMMA),
	5476:  uint16(5),
	5477:  uint16(1),
	5478:  uint16(sym_comment),
	5479:  uint16(7),
	5480:  uint16(1),
	5481:  uint16(anon_sym_SLASH_SLASH_SLASH),
	5482:  uint16(450),
	5483:  uint16(1),
	5484:  uint16(anon_sym_RBRACE),
	5485:  uint16(452),
	5486:  uint16(1),
	5487:  uint16(anon_sym_DQUOTE),
	5488:  uint16(455),
	5489:  uint16(1),
	5490:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5491:  uint16(458),
	5492:  uint16(1),
	5493:  uint16(aux_sym_identifier_token1),
	5494:  uint16(246),
	5495:  uint16(1),
	5496:  uint16(sym_node_object_kvp),
	5497:  uint16(340),
	5498:  uint16(1),
	5499:  uint16(sym_node_object_key),
	5500:  uint16(342),
	5501:  uint16(1),
	5502:  uint16(sym_string),
	5503:  uint16(343),
	5504:  uint16(1),
	5505:  uint16(sym_identifier),
	5506:  uint16(117),
	5507:  uint16(2),
	5508:  uint16(sym_documentation_comment),
	5509:  uint16(aux_sym_node_object_repeat1),
	5510:  uint16(168),
	5511:  uint16(2),
	5512:  uint16(sym__string_literal),
	5513:  uint16(sym__multiline_string_literal),
	5514:  uint16(14),
	5515:  uint16(3),
	5516:  uint16(1),
	5517:  uint16(anon_sym_COMMA),
	5518:  uint16(5),
	5519:  uint16(1),
	5520:  uint16(sym_comment),
	5521:  uint16(7),
	5522:  uint16(1),
	5523:  uint16(anon_sym_SLASH_SLASH_SLASH),
	5524:  uint16(133),
	5525:  uint16(1),
	5526:  uint16(anon_sym_DQUOTE),
	5527:  uint16(135),
	5528:  uint16(1),
	5529:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5530:  uint16(425),
	5531:  uint16(1),
	5532:  uint16(aux_sym_identifier_token1),
	5533:  uint16(461),
	5534:  uint16(1),
	5535:  uint16(anon_sym_RBRACE),
	5536:  uint16(115),
	5537:  uint16(1),
	5538:  uint16(aux_sym_node_object_repeat1),
	5539:  uint16(118),
	5540:  uint16(1),
	5541:  uint16(sym_documentation_comment),
	5542:  uint16(246),
	5543:  uint16(1),
	5544:  uint16(sym_node_object_kvp),
	5545:  uint16(340),
	5546:  uint16(1),
	5547:  uint16(sym_node_object_key),
	5548:  uint16(342),
	5549:  uint16(1),
	5550:  uint16(sym_string),
	5551:  uint16(343),
	5552:  uint16(1),
	5553:  uint16(sym_identifier),
	5554:  uint16(168),
	5555:  uint16(2),
	5556:  uint16(sym__string_literal),
	5557:  uint16(sym__multiline_string_literal),
	5558:  uint16(14),
	5559:  uint16(3),
	5560:  uint16(1),
	5561:  uint16(anon_sym_COMMA),
	5562:  uint16(5),
	5563:  uint16(1),
	5564:  uint16(sym_comment),
	5565:  uint16(7),
	5566:  uint16(1),
	5567:  uint16(anon_sym_SLASH_SLASH_SLASH),
	5568:  uint16(133),
	5569:  uint16(1),
	5570:  uint16(anon_sym_DQUOTE),
	5571:  uint16(135),
	5572:  uint16(1),
	5573:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5574:  uint16(425),
	5575:  uint16(1),
	5576:  uint16(aux_sym_identifier_token1),
	5577:  uint16(463),
	5578:  uint16(1),
	5579:  uint16(anon_sym_RBRACE),
	5580:  uint16(110),
	5581:  uint16(1),
	5582:  uint16(aux_sym_node_object_repeat1),
	5583:  uint16(119),
	5584:  uint16(1),
	5585:  uint16(sym_documentation_comment),
	5586:  uint16(246),
	5587:  uint16(1),
	5588:  uint16(sym_node_object_kvp),
	5589:  uint16(340),
	5590:  uint16(1),
	5591:  uint16(sym_node_object_key),
	5592:  uint16(342),
	5593:  uint16(1),
	5594:  uint16(sym_string),
	5595:  uint16(343),
	5596:  uint16(1),
	5597:  uint16(sym_identifier),
	5598:  uint16(168),
	5599:  uint16(2),
	5600:  uint16(sym__string_literal),
	5601:  uint16(sym__multiline_string_literal),
	5602:  uint16(9),
	5603:  uint16(3),
	5604:  uint16(1),
	5605:  uint16(anon_sym_COMMA),
	5606:  uint16(5),
	5607:  uint16(1),
	5608:  uint16(sym_comment),
	5609:  uint16(7),
	5610:  uint16(1),
	5611:  uint16(anon_sym_SLASH_SLASH_SLASH),
	5612:  uint16(335),
	5613:  uint16(1),
	5614:  uint16(anon_sym_DQUOTE),
	5615:  uint16(421),
	5616:  uint16(1),
	5617:  uint16(anon_sym_DOLLAR),
	5618:  uint16(109),
	5619:  uint16(1),
	5620:  uint16(aux_sym_shape_id_repeat1),
	5621:  uint16(120),
	5622:  uint16(1),
	5623:  uint16(sym_documentation_comment),
	5624:  uint16(121),
	5625:  uint16(1),
	5626:  uint16(sym_shape_id_member),
	5627:  uint16(103),
	5628:  uint16(7),
	5629:  uint16(anon_sym_RBRACE),
	5630:  uint16(anon_sym_RBRACK),
	5631:  uint16(anon_sym_AT),
	5632:  uint16(anon_sym_LPAREN),
	5633:  uint16(anon_sym_RPAREN),
	5634:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5635:  uint16(aux_sym_identifier_token1),
	5636:  uint16(6),
	5637:  uint16(3),
	5638:  uint16(1),
	5639:  uint16(anon_sym_COMMA),
	5640:  uint16(5),
	5641:  uint16(1),
	5642:  uint16(sym_comment),
	5643:  uint16(7),
	5644:  uint16(1),
	5645:  uint16(anon_sym_SLASH_SLASH_SLASH),
	5646:  uint16(348),
	5647:  uint16(1),
	5648:  uint16(anon_sym_DQUOTE),
	5649:  uint16(121),
	5650:  uint16(1),
	5651:  uint16(sym_documentation_comment),
	5652:  uint16(107),
	5653:  uint16(9),
	5654:  uint16(anon_sym_DOLLAR),
	5655:  uint16(anon_sym_EQ),
	5656:  uint16(anon_sym_RBRACE),
	5657:  uint16(anon_sym_RBRACK),
	5658:  uint16(anon_sym_AT),
	5659:  uint16(anon_sym_LPAREN),
	5660:  uint16(anon_sym_RPAREN),
	5661:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5662:  uint16(aux_sym_identifier_token1),
	5663:  uint16(6),
	5664:  uint16(3),
	5665:  uint16(1),
	5666:  uint16(anon_sym_COMMA),
	5667:  uint16(5),
	5668:  uint16(1),
	5669:  uint16(sym_comment),
	5670:  uint16(7),
	5671:  uint16(1),
	5672:  uint16(anon_sym_SLASH_SLASH_SLASH),
	5673:  uint16(92),
	5674:  uint16(1),
	5675:  uint16(anon_sym_DQUOTE),
	5676:  uint16(122),
	5677:  uint16(1),
	5678:  uint16(sym_documentation_comment),
	5679:  uint16(88),
	5680:  uint16(9),
	5681:  uint16(anon_sym_DOLLAR),
	5682:  uint16(anon_sym_EQ),
	5683:  uint16(anon_sym_RBRACE),
	5684:  uint16(anon_sym_RBRACK),
	5685:  uint16(anon_sym_AT),
	5686:  uint16(anon_sym_LPAREN),
	5687:  uint16(anon_sym_RPAREN),
	5688:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5689:  uint16(aux_sym_identifier_token1),
	5690:  uint16(6),
	5691:  uint16(3),
	5692:  uint16(1),
	5693:  uint16(anon_sym_COMMA),
	5694:  uint16(5),
	5695:  uint16(1),
	5696:  uint16(sym_comment),
	5697:  uint16(7),
	5698:  uint16(1),
	5699:  uint16(anon_sym_SLASH_SLASH_SLASH),
	5700:  uint16(350),
	5701:  uint16(1),
	5702:  uint16(anon_sym_DQUOTE),
	5703:  uint16(123),
	5704:  uint16(1),
	5705:  uint16(sym_documentation_comment),
	5706:  uint16(109),
	5707:  uint16(9),
	5708:  uint16(anon_sym_DOLLAR),
	5709:  uint16(anon_sym_EQ),
	5710:  uint16(anon_sym_RBRACE),
	5711:  uint16(anon_sym_RBRACK),
	5712:  uint16(anon_sym_AT),
	5713:  uint16(anon_sym_LPAREN),
	5714:  uint16(anon_sym_RPAREN),
	5715:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5716:  uint16(aux_sym_identifier_token1),
	5717:  uint16(14),
	5718:  uint16(3),
	5719:  uint16(1),
	5720:  uint16(anon_sym_COMMA),
	5721:  uint16(5),
	5722:  uint16(1),
	5723:  uint16(sym_comment),
	5724:  uint16(7),
	5725:  uint16(1),
	5726:  uint16(anon_sym_SLASH_SLASH_SLASH),
	5727:  uint16(425),
	5728:  uint16(1),
	5729:  uint16(aux_sym_identifier_token1),
	5730:  uint16(465),
	5731:  uint16(1),
	5732:  uint16(anon_sym_DOLLAR),
	5733:  uint16(467),
	5734:  uint16(1),
	5735:  uint16(anon_sym_RBRACE),
	5736:  uint16(469),
	5737:  uint16(1),
	5738:  uint16(anon_sym_AT),
	5739:  uint16(124),
	5740:  uint16(1),
	5741:  uint16(sym_documentation_comment),
	5742:  uint16(126),
	5743:  uint16(1),
	5744:  uint16(aux_sym_shape_members_repeat1),
	5745:  uint16(176),
	5746:  uint16(1),
	5747:  uint16(aux_sym_shape_statement_repeat1),
	5748:  uint16(192),
	5749:  uint16(1),
	5750:  uint16(sym_shape_member_elided),
	5751:  uint16(244),
	5752:  uint16(1),
	5753:  uint16(sym_shape_member),
	5754:  uint16(274),
	5755:  uint16(1),
	5756:  uint16(sym_trait_statement),
	5757:  uint16(330),
	5758:  uint16(1),
	5759:  uint16(sym_identifier),
	5760:  uint16(14),
	5761:  uint16(3),
	5762:  uint16(1),
	5763:  uint16(anon_sym_COMMA),
	5764:  uint16(5),
	5765:  uint16(1),
	5766:  uint16(sym_comment),
	5767:  uint16(7),
	5768:  uint16(1),
	5769:  uint16(anon_sym_SLASH_SLASH_SLASH),
	5770:  uint16(425),
	5771:  uint16(1),
	5772:  uint16(aux_sym_identifier_token1),
	5773:  uint16(465),
	5774:  uint16(1),
	5775:  uint16(anon_sym_DOLLAR),
	5776:  uint16(469),
	5777:  uint16(1),
	5778:  uint16(anon_sym_AT),
	5779:  uint16(471),
	5780:  uint16(1),
	5781:  uint16(anon_sym_RBRACE),
	5782:  uint16(124),
	5783:  uint16(1),
	5784:  uint16(aux_sym_shape_members_repeat1),
	5785:  uint16(125),
	5786:  uint16(1),
	5787:  uint16(sym_documentation_comment),
	5788:  uint16(176),
	5789:  uint16(1),
	5790:  uint16(aux_sym_shape_statement_repeat1),
	5791:  uint16(192),
	5792:  uint16(1),
	5793:  uint16(sym_shape_member_elided),
	5794:  uint16(244),
	5795:  uint16(1),
	5796:  uint16(sym_shape_member),
	5797:  uint16(274),
	5798:  uint16(1),
	5799:  uint16(sym_trait_statement),
	5800:  uint16(330),
	5801:  uint16(1),
	5802:  uint16(sym_identifier),
	5803:  uint16(13),
	5804:  uint16(3),
	5805:  uint16(1),
	5806:  uint16(anon_sym_COMMA),
	5807:  uint16(5),
	5808:  uint16(1),
	5809:  uint16(sym_comment),
	5810:  uint16(7),
	5811:  uint16(1),
	5812:  uint16(anon_sym_SLASH_SLASH_SLASH),
	5813:  uint16(473),
	5814:  uint16(1),
	5815:  uint16(anon_sym_DOLLAR),
	5816:  uint16(476),
	5817:  uint16(1),
	5818:  uint16(anon_sym_RBRACE),
	5819:  uint16(478),
	5820:  uint16(1),
	5821:  uint16(anon_sym_AT),
	5822:  uint16(481),
	5823:  uint16(1),
	5824:  uint16(aux_sym_identifier_token1),
	5825:  uint16(176),
	5826:  uint16(1),
	5827:  uint16(aux_sym_shape_statement_repeat1),
	5828:  uint16(192),
	5829:  uint16(1),
	5830:  uint16(sym_shape_member_elided),
	5831:  uint16(244),
	5832:  uint16(1),
	5833:  uint16(sym_shape_member),
	5834:  uint16(274),
	5835:  uint16(1),
	5836:  uint16(sym_trait_statement),
	5837:  uint16(330),
	5838:  uint16(1),
	5839:  uint16(sym_identifier),
	5840:  uint16(126),
	5841:  uint16(2),
	5842:  uint16(sym_documentation_comment),
	5843:  uint16(aux_sym_shape_members_repeat1),
	5844:  uint16(6),
	5845:  uint16(3),
	5846:  uint16(1),
	5847:  uint16(anon_sym_COMMA),
	5848:  uint16(5),
	5849:  uint16(1),
	5850:  uint16(sym_comment),
	5851:  uint16(7),
	5852:  uint16(1),
	5853:  uint16(anon_sym_SLASH_SLASH_SLASH),
	5854:  uint16(352),
	5855:  uint16(1),
	5856:  uint16(anon_sym_DQUOTE),
	5857:  uint16(127),
	5858:  uint16(1),
	5859:  uint16(sym_documentation_comment),
	5860:  uint16(111),
	5861:  uint16(9),
	5862:  uint16(anon_sym_DOLLAR),
	5863:  uint16(anon_sym_EQ),
	5864:  uint16(anon_sym_RBRACE),
	5865:  uint16(anon_sym_RBRACK),
	5866:  uint16(anon_sym_AT),
	5867:  uint16(anon_sym_LPAREN),
	5868:  uint16(anon_sym_RPAREN),
	5869:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5870:  uint16(aux_sym_identifier_token1),
	5871:  uint16(14),
	5872:  uint16(3),
	5873:  uint16(1),
	5874:  uint16(anon_sym_COMMA),
	5875:  uint16(5),
	5876:  uint16(1),
	5877:  uint16(sym_comment),
	5878:  uint16(7),
	5879:  uint16(1),
	5880:  uint16(anon_sym_SLASH_SLASH_SLASH),
	5881:  uint16(425),
	5882:  uint16(1),
	5883:  uint16(aux_sym_identifier_token1),
	5884:  uint16(465),
	5885:  uint16(1),
	5886:  uint16(anon_sym_DOLLAR),
	5887:  uint16(469),
	5888:  uint16(1),
	5889:  uint16(anon_sym_AT),
	5890:  uint16(484),
	5891:  uint16(1),
	5892:  uint16(anon_sym_RBRACE),
	5893:  uint16(126),
	5894:  uint16(1),
	5895:  uint16(aux_sym_shape_members_repeat1),
	5896:  uint16(128),
	5897:  uint16(1),
	5898:  uint16(sym_documentation_comment),
	5899:  uint16(176),
	5900:  uint16(1),
	5901:  uint16(aux_sym_shape_statement_repeat1),
	5902:  uint16(192),
	5903:  uint16(1),
	5904:  uint16(sym_shape_member_elided),
	5905:  uint16(244),
	5906:  uint16(1),
	5907:  uint16(sym_shape_member),
	5908:  uint16(274),
	5909:  uint16(1),
	5910:  uint16(sym_trait_statement),
	5911:  uint16(330),
	5912:  uint16(1),
	5913:  uint16(sym_identifier),
	5914:  uint16(14),
	5915:  uint16(3),
	5916:  uint16(1),
	5917:  uint16(anon_sym_COMMA),
	5918:  uint16(5),
	5919:  uint16(1),
	5920:  uint16(sym_comment),
	5921:  uint16(7),
	5922:  uint16(1),
	5923:  uint16(anon_sym_SLASH_SLASH_SLASH),
	5924:  uint16(425),
	5925:  uint16(1),
	5926:  uint16(aux_sym_identifier_token1),
	5927:  uint16(465),
	5928:  uint16(1),
	5929:  uint16(anon_sym_DOLLAR),
	5930:  uint16(469),
	5931:  uint16(1),
	5932:  uint16(anon_sym_AT),
	5933:  uint16(486),
	5934:  uint16(1),
	5935:  uint16(anon_sym_RBRACE),
	5936:  uint16(128),
	5937:  uint16(1),
	5938:  uint16(aux_sym_shape_members_repeat1),
	5939:  uint16(129),
	5940:  uint16(1),
	5941:  uint16(sym_documentation_comment),
	5942:  uint16(176),
	5943:  uint16(1),
	5944:  uint16(aux_sym_shape_statement_repeat1),
	5945:  uint16(192),
	5946:  uint16(1),
	5947:  uint16(sym_shape_member_elided),
	5948:  uint16(244),
	5949:  uint16(1),
	5950:  uint16(sym_shape_member),
	5951:  uint16(274),
	5952:  uint16(1),
	5953:  uint16(sym_trait_statement),
	5954:  uint16(330),
	5955:  uint16(1),
	5956:  uint16(sym_identifier),
	5957:  uint16(6),
	5958:  uint16(3),
	5959:  uint16(1),
	5960:  uint16(anon_sym_COMMA),
	5961:  uint16(5),
	5962:  uint16(1),
	5963:  uint16(sym_comment),
	5964:  uint16(7),
	5965:  uint16(1),
	5966:  uint16(anon_sym_SLASH_SLASH_SLASH),
	5967:  uint16(346),
	5968:  uint16(1),
	5969:  uint16(anon_sym_DQUOTE),
	5970:  uint16(130),
	5971:  uint16(1),
	5972:  uint16(sym_documentation_comment),
	5973:  uint16(105),
	5974:  uint16(9),
	5975:  uint16(anon_sym_DOLLAR),
	5976:  uint16(anon_sym_EQ),
	5977:  uint16(anon_sym_RBRACE),
	5978:  uint16(anon_sym_RBRACK),
	5979:  uint16(anon_sym_AT),
	5980:  uint16(anon_sym_LPAREN),
	5981:  uint16(anon_sym_RPAREN),
	5982:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5983:  uint16(aux_sym_identifier_token1),
	5984:  uint16(12),
	5985:  uint16(3),
	5986:  uint16(1),
	5987:  uint16(anon_sym_COMMA),
	5988:  uint16(5),
	5989:  uint16(1),
	5990:  uint16(sym_comment),
	5991:  uint16(7),
	5992:  uint16(1),
	5993:  uint16(anon_sym_SLASH_SLASH_SLASH),
	5994:  uint16(488),
	5995:  uint16(1),
	5996:  uint16(anon_sym_RBRACK),
	5997:  uint16(490),
	5998:  uint16(1),
	5999:  uint16(aux_sym_identifier_token1),
	6000:  uint16(23),
	6001:  uint16(1),
	6002:  uint16(sym__namespace_identifier),
	6003:  uint16(120),
	6004:  uint16(1),
	6005:  uint16(sym_root_shape_id),
	6006:  uint16(131),
	6007:  uint16(1),
	6008:  uint16(sym_documentation_comment),
	6009:  uint16(133),
	6010:  uint16(1),
	6011:  uint16(aux_sym_mixins_repeat1),
	6012:  uint16(282),
	6013:  uint16(1),
	6014:  uint16(sym_shape_id),
	6015:  uint16(345),
	6016:  uint16(1),
	6017:  uint16(sym_namespace),
	6018:  uint16(130),
	6019:  uint16(2),
	6020:  uint16(sym_absolute_root_shape_id),
	6021:  uint16(sym_identifier),
	6022:  uint16(13),
	6023:  uint16(3),
	6024:  uint16(1),
	6025:  uint16(anon_sym_COMMA),
	6026:  uint16(5),
	6027:  uint16(1),
	6028:  uint16(sym_comment),
	6029:  uint16(7),
	6030:  uint16(1),
	6031:  uint16(anon_sym_SLASH_SLASH_SLASH),
	6032:  uint16(39),
	6033:  uint16(1),
	6034:  uint16(anon_sym_AT),
	6035:  uint16(227),
	6036:  uint16(1),
	6037:  uint16(anon_sym_with),
	6038:  uint16(492),
	6039:  uint16(1),
	6040:  uint16(anon_sym_LBRACE),
	6041:  uint16(494),
	6042:  uint16(1),
	6043:  uint16(anon_sym_for),
	6044:  uint16(26),
	6045:  uint16(1),
	6046:  uint16(aux_sym_shape_statement_repeat1),
	6047:  uint16(37),
	6048:  uint16(1),
	6049:  uint16(sym_trait_statement),
	6050:  uint16(132),
	6051:  uint16(1),
	6052:  uint16(sym_documentation_comment),
	6053:  uint16(247),
	6054:  uint16(1),
	6055:  uint16(sym_structure_resource),
	6056:  uint16(311),
	6057:  uint16(1),
	6058:  uint16(sym_mixins),
	6059:  uint16(314),
	6060:  uint16(1),
	6061:  uint16(sym_shape_members),
	6062:  uint16(11),
	6063:  uint16(3),
	6064:  uint16(1),
	6065:  uint16(anon_sym_COMMA),
	6066:  uint16(5),
	6067:  uint16(1),
	6068:  uint16(sym_comment),
	6069:  uint16(7),
	6070:  uint16(1),
	6071:  uint16(anon_sym_SLASH_SLASH_SLASH),
	6072:  uint16(496),
	6073:  uint16(1),
	6074:  uint16(anon_sym_RBRACK),
	6075:  uint16(498),
	6076:  uint16(1),
	6077:  uint16(aux_sym_identifier_token1),
	6078:  uint16(23),
	6079:  uint16(1),
	6080:  uint16(sym__namespace_identifier),
	6081:  uint16(120),
	6082:  uint16(1),
	6083:  uint16(sym_root_shape_id),
	6084:  uint16(282),
	6085:  uint16(1),
	6086:  uint16(sym_shape_id),
	6087:  uint16(345),
	6088:  uint16(1),
	6089:  uint16(sym_namespace),
	6090:  uint16(130),
	6091:  uint16(2),
	6092:  uint16(sym_absolute_root_shape_id),
	6093:  uint16(sym_identifier),
	6094:  uint16(133),
	6095:  uint16(2),
	6096:  uint16(sym_documentation_comment),
	6097:  uint16(aux_sym_mixins_repeat1),
	6098:  uint16(13),
	6099:  uint16(3),
	6100:  uint16(1),
	6101:  uint16(anon_sym_COMMA),
	6102:  uint16(5),
	6103:  uint16(1),
	6104:  uint16(sym_comment),
	6105:  uint16(7),
	6106:  uint16(1),
	6107:  uint16(anon_sym_SLASH_SLASH_SLASH),
	6108:  uint16(39),
	6109:  uint16(1),
	6110:  uint16(anon_sym_AT),
	6111:  uint16(227),
	6112:  uint16(1),
	6113:  uint16(anon_sym_with),
	6114:  uint16(492),
	6115:  uint16(1),
	6116:  uint16(anon_sym_LBRACE),
	6117:  uint16(494),
	6118:  uint16(1),
	6119:  uint16(anon_sym_for),
	6120:  uint16(37),
	6121:  uint16(1),
	6122:  uint16(sym_trait_statement),
	6123:  uint16(132),
	6124:  uint16(1),
	6125:  uint16(aux_sym_shape_statement_repeat1),
	6126:  uint16(134),
	6127:  uint16(1),
	6128:  uint16(sym_documentation_comment),
	6129:  uint16(238),
	6130:  uint16(1),
	6131:  uint16(sym_structure_resource),
	6132:  uint16(276),
	6133:  uint16(1),
	6134:  uint16(sym_mixins),
	6135:  uint16(298),
	6136:  uint16(1),
	6137:  uint16(sym_shape_members),
	6138:  uint16(11),
	6139:  uint16(3),
	6140:  uint16(1),
	6141:  uint16(anon_sym_COMMA),
	6142:  uint16(5),
	6143:  uint16(1),
	6144:  uint16(sym_comment),
	6145:  uint16(7),
	6146:  uint16(1),
	6147:  uint16(anon_sym_SLASH_SLASH_SLASH),
	6148:  uint16(490),
	6149:  uint16(1),
	6150:  uint16(aux_sym_identifier_token1),
	6151:  uint16(501),
	6152:  uint16(1),
	6153:  uint16(anon_sym_LBRACK),
	6154:  uint16(23),
	6155:  uint16(1),
	6156:  uint16(sym__namespace_identifier),
	6157:  uint16(120),
	6158:  uint16(1),
	6159:  uint16(sym_root_shape_id),
	6160:  uint16(135),
	6161:  uint16(1),
	6162:  uint16(sym_documentation_comment),
	6163:  uint16(291),
	6164:  uint16(1),
	6165:  uint16(sym_shape_id),
	6166:  uint16(345),
	6167:  uint16(1),
	6168:  uint16(sym_namespace),
	6169:  uint16(130),
	6170:  uint16(2),
	6171:  uint16(sym_absolute_root_shape_id),
	6172:  uint16(sym_identifier),
	6173:  uint16(12),
	6174:  uint16(3),
	6175:  uint16(1),
	6176:  uint16(anon_sym_COMMA),
	6177:  uint16(5),
	6178:  uint16(1),
	6179:  uint16(sym_comment),
	6180:  uint16(7),
	6181:  uint16(1),
	6182:  uint16(anon_sym_SLASH_SLASH_SLASH),
	6183:  uint16(13),
	6184:  uint16(1),
	6185:  uint16(anon_sym_metadata),
	6186:  uint16(15),
	6187:  uint16(1),
	6188:  uint16(anon_sym_namespace),
	6189:  uint16(503),
	6190:  uint16(1),
	6192:  uint16(2),
	6193:  uint16(1),
	6194:  uint16(sym_namespace_statement),
	6195:  uint16(136),
	6196:  uint16(1),
	6197:  uint16(sym_documentation_comment),
	6198:  uint16(224),
	6199:  uint16(1),
	6200:  uint16(aux_sym_metadata_section_repeat1),
	6201:  uint16(249),
	6202:  uint16(1),
	6203:  uint16(sym_metadata_section),
	6204:  uint16(270),
	6205:  uint16(1),
	6206:  uint16(sym_metadata_statement),
	6207:  uint16(327),
	6208:  uint16(1),
	6209:  uint16(sym_shape_section),
	6210:  uint16(9),
	6211:  uint16(7),
	6212:  uint16(1),
	6213:  uint16(anon_sym_SLASH_SLASH_SLASH),
	6214:  uint16(505),
	6215:  uint16(1),
	6216:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	6217:  uint16(507),
	6218:  uint16(1),
	6219:  uint16(aux_sym__multiline_string_fragment_token1),
	6220:  uint16(510),
	6221:  uint16(1),
	6222:  uint16(aux_sym__multiline_string_fragment_token2),
	6223:  uint16(513),
	6224:  uint16(1),
	6225:  uint16(aux_sym__escape_sequence_token1),
	6226:  uint16(516),
	6227:  uint16(1),
	6228:  uint16(sym_escape_sequence),
	6229:  uint16(5),
	6230:  uint16(2),
	6231:  uint16(anon_sym_COMMA),
	6232:  uint16(sym_comment),
	6233:  uint16(137),
	6234:  uint16(2),
	6235:  uint16(sym_documentation_comment),
	6236:  uint16(aux_sym__multiline_string_literal_repeat1),
	6237:  uint16(206),
	6238:  uint16(2),
	6239:  uint16(sym__multiline_string_fragment),
	6240:  uint16(sym__escape_sequence),
	6241:  uint16(10),
	6242:  uint16(7),
	6243:  uint16(1),
	6244:  uint16(anon_sym_SLASH_SLASH_SLASH),
	6245:  uint16(519),
	6246:  uint16(1),
	6247:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	6248:  uint16(521),
	6249:  uint16(1),
	6250:  uint16(aux_sym__multiline_string_fragment_token1),
	6251:  uint16(523),
	6252:  uint16(1),
	6253:  uint16(aux_sym__multiline_string_fragment_token2),
	6254:  uint16(525),
	6255:  uint16(1),
	6256:  uint16(aux_sym__escape_sequence_token1),
	6257:  uint16(527),
	6258:  uint16(1),
	6259:  uint16(sym_escape_sequence),
	6260:  uint16(138),
	6261:  uint16(1),
	6262:  uint16(sym_documentation_comment),
	6263:  uint16(144),
	6264:  uint16(1),
	6265:  uint16(aux_sym__multiline_string_literal_repeat1),
	6266:  uint16(5),
	6267:  uint16(2),
	6268:  uint16(anon_sym_COMMA),
	6269:  uint16(sym_comment),
	6270:  uint16(206),
	6271:  uint16(2),
	6272:  uint16(sym__multiline_string_fragment),
	6273:  uint16(sym__escape_sequence),
	6274:  uint16(10),
	6275:  uint16(7),
	6276:  uint16(1),
	6277:  uint16(anon_sym_SLASH_SLASH_SLASH),
	6278:  uint16(521),
	6279:  uint16(1),
	6280:  uint16(aux_sym__multiline_string_fragment_token1),
	6281:  uint16(523),
	6282:  uint16(1),
	6283:  uint16(aux_sym__multiline_string_fragment_token2),
	6284:  uint16(525),
	6285:  uint16(1),
	6286:  uint16(aux_sym__escape_sequence_token1),
	6287:  uint16(527),
	6288:  uint16(1),
	6289:  uint16(sym_escape_sequence),
	6290:  uint16(529),
	6291:  uint16(1),
	6292:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	6293:  uint16(137),
	6294:  uint16(1),
	6295:  uint16(aux_sym__multiline_string_literal_repeat1),
	6296:  uint16(139),
	6297:  uint16(1),
	6298:  uint16(sym_documentation_comment),
	6299:  uint16(5),
	6300:  uint16(2),
	6301:  uint16(anon_sym_COMMA),
	6302:  uint16(sym_comment),
	6303:  uint16(206),
	6304:  uint16(2),
	6305:  uint16(sym__multiline_string_fragment),
	6306:  uint16(sym__escape_sequence),
	6307:  uint16(12),
	6308:  uint16(3),
	6309:  uint16(1),
	6310:  uint16(anon_sym_COMMA),
	6311:  uint16(5),
	6312:  uint16(1),
	6313:  uint16(sym_comment),
	6314:  uint16(7),
	6315:  uint16(1),
	6316:  uint16(anon_sym_SLASH_SLASH_SLASH),
	6317:  uint16(531),
	6318:  uint16(1),
	6319:  uint16(anon_sym_RBRACE),
	6320:  uint16(533),
	6321:  uint16(1),
	6322:  uint16(anon_sym_AT),
	6323:  uint16(535),
	6324:  uint16(1),
	6325:  uint16(aux_sym_identifier_token1),
	6326:  uint16(140),
	6327:  uint16(1),
	6328:  uint16(sym_documentation_comment),
	6329:  uint16(145),
	6330:  uint16(1),
	6331:  uint16(aux_sym_enum_members_repeat1),
	6332:  uint16(213),
	6333:  uint16(1),
	6334:  uint16(sym_identifier),
	6335:  uint16(222),
	6336:  uint16(1),
	6337:  uint16(aux_sym_shape_statement_repeat1),
	6338:  uint16(262),
	6339:  uint16(1),
	6340:  uint16(sym_enum_member),
	6341:  uint16(274),
	6342:  uint16(1),
	6343:  uint16(sym_trait_statement),
	6344:  uint16(11),
	6345:  uint16(3),
	6346:  uint16(1),
	6347:  uint16(anon_sym_COMMA),
	6348:  uint16(5),
	6349:  uint16(1),
	6350:  uint16(sym_comment),
	6351:  uint16(7),
	6352:  uint16(1),
	6353:  uint16(anon_sym_SLASH_SLASH_SLASH),
	6354:  uint16(537),
	6355:  uint16(1),
	6356:  uint16(anon_sym_RBRACE),
	6357:  uint16(539),
	6358:  uint16(1),
	6359:  uint16(anon_sym_AT),
	6360:  uint16(542),
	6361:  uint16(1),
	6362:  uint16(aux_sym_identifier_token1),
	6363:  uint16(213),
	6364:  uint16(1),
	6365:  uint16(sym_identifier),
	6366:  uint16(222),
	6367:  uint16(1),
	6368:  uint16(aux_sym_shape_statement_repeat1),
	6369:  uint16(262),
	6370:  uint16(1),
	6371:  uint16(sym_enum_member),
	6372:  uint16(274),
	6373:  uint16(1),
	6374:  uint16(sym_trait_statement),
	6375:  uint16(141),
	6376:  uint16(2),
	6377:  uint16(sym_documentation_comment),
	6378:  uint16(aux_sym_enum_members_repeat1),
	6379:  uint16(10),
	6380:  uint16(7),
	6381:  uint16(1),
	6382:  uint16(anon_sym_SLASH_SLASH_SLASH),
	6383:  uint16(521),
	6384:  uint16(1),
	6385:  uint16(aux_sym__multiline_string_fragment_token1),
	6386:  uint16(523),
	6387:  uint16(1),
	6388:  uint16(aux_sym__multiline_string_fragment_token2),
	6389:  uint16(525),
	6390:  uint16(1),
	6391:  uint16(aux_sym__escape_sequence_token1),
	6392:  uint16(527),
	6393:  uint16(1),
	6394:  uint16(sym_escape_sequence),
	6395:  uint16(545),
	6396:  uint16(1),
	6397:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	6398:  uint16(137),
	6399:  uint16(1),
	6400:  uint16(aux_sym__multiline_string_literal_repeat1),
	6401:  uint16(142),
	6402:  uint16(1),
	6403:  uint16(sym_documentation_comment),
	6404:  uint16(5),
	6405:  uint16(2),
	6406:  uint16(anon_sym_COMMA),
	6407:  uint16(sym_comment),
	6408:  uint16(206),
	6409:  uint16(2),
	6410:  uint16(sym__multiline_string_fragment),
	6411:  uint16(sym__escape_sequence),
	6412:  uint16(7),
	6413:  uint16(3),
	6414:  uint16(1),
	6415:  uint16(anon_sym_COMMA),
	6416:  uint16(5),
	6417:  uint16(1),
	6418:  uint16(sym_comment),
	6419:  uint16(7),
	6420:  uint16(1),
	6421:  uint16(anon_sym_SLASH_SLASH_SLASH),
	6422:  uint16(121),
	6423:  uint16(1),
	6424:  uint16(sym_shape_id_member),
	6425:  uint16(143),
	6426:  uint16(1),
	6427:  uint16(sym_documentation_comment),
	6428:  uint16(149),
	6429:  uint16(1),
	6430:  uint16(aux_sym_shape_id_repeat1),
	6431:  uint16(103),
	6432:  uint16(6),
	6433:  uint16(anon_sym_DOLLAR),
	6434:  uint16(anon_sym_EQ),
	6435:  uint16(anon_sym_RBRACE),
	6436:  uint16(anon_sym_AT),
	6437:  uint16(anon_sym_LPAREN),
	6438:  uint16(aux_sym_identifier_token1),
	6439:  uint16(10),
	6440:  uint16(7),
	6441:  uint16(1),
	6442:  uint16(anon_sym_SLASH_SLASH_SLASH),
	6443:  uint16(521),
	6444:  uint16(1),
	6445:  uint16(aux_sym__multiline_string_fragment_token1),
	6446:  uint16(523),
	6447:  uint16(1),
	6448:  uint16(aux_sym__multiline_string_fragment_token2),
	6449:  uint16(525),
	6450:  uint16(1),
	6451:  uint16(aux_sym__escape_sequence_token1),
	6452:  uint16(527),
	6453:  uint16(1),
	6454:  uint16(sym_escape_sequence),
	6455:  uint16(547),
	6456:  uint16(1),
	6457:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	6458:  uint16(137),
	6459:  uint16(1),
	6460:  uint16(aux_sym__multiline_string_literal_repeat1),
	6461:  uint16(144),
	6462:  uint16(1),
	6463:  uint16(sym_documentation_comment),
	6464:  uint16(5),
	6465:  uint16(2),
	6466:  uint16(anon_sym_COMMA),
	6467:  uint16(sym_comment),
	6468:  uint16(206),
	6469:  uint16(2),
	6470:  uint16(sym__multiline_string_fragment),
	6471:  uint16(sym__escape_sequence),
	6472:  uint16(12),
	6473:  uint16(3),
	6474:  uint16(1),
	6475:  uint16(anon_sym_COMMA),
	6476:  uint16(5),
	6477:  uint16(1),
	6478:  uint16(sym_comment),
	6479:  uint16(7),
	6480:  uint16(1),
	6481:  uint16(anon_sym_SLASH_SLASH_SLASH),
	6482:  uint16(533),
	6483:  uint16(1),
	6484:  uint16(anon_sym_AT),
	6485:  uint16(535),
	6486:  uint16(1),
	6487:  uint16(aux_sym_identifier_token1),
	6488:  uint16(549),
	6489:  uint16(1),
	6490:  uint16(anon_sym_RBRACE),
	6491:  uint16(141),
	6492:  uint16(1),
	6493:  uint16(aux_sym_enum_members_repeat1),
	6494:  uint16(145),
	6495:  uint16(1),
	6496:  uint16(sym_documentation_comment),
	6497:  uint16(213),
	6498:  uint16(1),
	6499:  uint16(sym_identifier),
	6500:  uint16(222),
	6501:  uint16(1),
	6502:  uint16(aux_sym_shape_statement_repeat1),
	6503:  uint16(262),
	6504:  uint16(1),
	6505:  uint16(sym_enum_member),
	6506:  uint16(274),
	6507:  uint16(1),
	6508:  uint16(sym_trait_statement),
	6509:  uint16(10),
	6510:  uint16(7),
	6511:  uint16(1),
	6512:  uint16(anon_sym_SLASH_SLASH_SLASH),
	6513:  uint16(521),
	6514:  uint16(1),
	6515:  uint16(aux_sym__multiline_string_fragment_token1),
	6516:  uint16(523),
	6517:  uint16(1),
	6518:  uint16(aux_sym__multiline_string_fragment_token2),
	6519:  uint16(525),
	6520:  uint16(1),
	6521:  uint16(aux_sym__escape_sequence_token1),
	6522:  uint16(527),
	6523:  uint16(1),
	6524:  uint16(sym_escape_sequence),
	6525:  uint16(551),
	6526:  uint16(1),
	6527:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	6528:  uint16(139),
	6529:  uint16(1),
	6530:  uint16(aux_sym__multiline_string_literal_repeat1),
	6531:  uint16(146),
	6532:  uint16(1),
	6533:  uint16(sym_documentation_comment),
	6534:  uint16(5),
	6535:  uint16(2),
	6536:  uint16(anon_sym_COMMA),
	6537:  uint16(sym_comment),
	6538:  uint16(206),
	6539:  uint16(2),
	6540:  uint16(sym__multiline_string_fragment),
	6541:  uint16(sym__escape_sequence),
	6542:  uint16(10),
	6543:  uint16(7),
	6544:  uint16(1),
	6545:  uint16(anon_sym_SLASH_SLASH_SLASH),
	6546:  uint16(521),
	6547:  uint16(1),
	6548:  uint16(aux_sym__multiline_string_fragment_token1),
	6549:  uint16(523),
	6550:  uint16(1),
	6551:  uint16(aux_sym__multiline_string_fragment_token2),
	6552:  uint16(525),
	6553:  uint16(1),
	6554:  uint16(aux_sym__escape_sequence_token1),
	6555:  uint16(527),
	6556:  uint16(1),
	6557:  uint16(sym_escape_sequence),
	6558:  uint16(553),
	6559:  uint16(1),
	6560:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	6561:  uint16(142),
	6562:  uint16(1),
	6563:  uint16(aux_sym__multiline_string_literal_repeat1),
	6564:  uint16(147),
	6565:  uint16(1),
	6566:  uint16(sym_documentation_comment),
	6567:  uint16(5),
	6568:  uint16(2),
	6569:  uint16(anon_sym_COMMA),
	6570:  uint16(sym_comment),
	6571:  uint16(206),
	6572:  uint16(2),
	6573:  uint16(sym__multiline_string_fragment),
	6574:  uint16(sym__escape_sequence),
	6575:  uint16(11),
	6576:  uint16(3),
	6577:  uint16(1),
	6578:  uint16(anon_sym_COMMA),
	6579:  uint16(5),
	6580:  uint16(1),
	6581:  uint16(sym_comment),
	6582:  uint16(7),
	6583:  uint16(1),
	6584:  uint16(anon_sym_SLASH_SLASH_SLASH),
	6585:  uint16(490),
	6586:  uint16(1),
	6587:  uint16(aux_sym_identifier_token1),
	6588:  uint16(23),
	6589:  uint16(1),
	6590:  uint16(sym__namespace_identifier),
	6591:  uint16(120),
	6592:  uint16(1),
	6593:  uint16(sym_root_shape_id),
	6594:  uint16(131),
	6595:  uint16(1),
	6596:  uint16(aux_sym_mixins_repeat1),
	6597:  uint16(148),
	6598:  uint16(1),
	6599:  uint16(sym_documentation_comment),
	6600:  uint16(282),
	6601:  uint16(1),
	6602:  uint16(sym_shape_id),
	6603:  uint16(345),
	6604:  uint16(1),
	6605:  uint16(sym_namespace),
	6606:  uint16(130),
	6607:  uint16(2),
	6608:  uint16(sym_absolute_root_shape_id),
	6609:  uint16(sym_identifier),
	6610:  uint16(7),
	6611:  uint16(3),
	6612:  uint16(1),
	6613:  uint16(anon_sym_COMMA),
	6614:  uint16(5),
	6615:  uint16(1),
	6616:  uint16(sym_comment),
	6617:  uint16(7),
	6618:  uint16(1),
	6619:  uint16(anon_sym_SLASH_SLASH_SLASH),
	6620:  uint16(94),
	6621:  uint16(1),
	6622:  uint16(aux_sym_shape_id_repeat1),
	6623:  uint16(121),
	6624:  uint16(1),
	6625:  uint16(sym_shape_id_member),
	6626:  uint16(149),
	6627:  uint16(1),
	6628:  uint16(sym_documentation_comment),
	6629:  uint16(99),
	6630:  uint16(6),
	6631:  uint16(anon_sym_DOLLAR),
	6632:  uint16(anon_sym_EQ),
	6633:  uint16(anon_sym_RBRACE),
	6634:  uint16(anon_sym_AT),
	6635:  uint16(anon_sym_LPAREN),
	6636:  uint16(aux_sym_identifier_token1),
	6637:  uint16(6),
	6638:  uint16(3),
	6639:  uint16(1),
	6640:  uint16(anon_sym_COMMA),
	6641:  uint16(5),
	6642:  uint16(1),
	6643:  uint16(sym_comment),
	6644:  uint16(7),
	6645:  uint16(1),
	6646:  uint16(anon_sym_SLASH_SLASH_SLASH),
	6647:  uint16(389),
	6648:  uint16(1),
	6649:  uint16(anon_sym_DQUOTE),
	6650:  uint16(150),
	6651:  uint16(1),
	6652:  uint16(sym_documentation_comment),
	6653:  uint16(143),
	6654:  uint16(6),
	6655:  uint16(anon_sym_DOLLAR),
	6656:  uint16(anon_sym_RBRACE),
	6657:  uint16(anon_sym_AT),
	6658:  uint16(anon_sym_RPAREN),
	6659:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	6660:  uint16(aux_sym_identifier_token1),
	6661:  uint16(11),
	6662:  uint16(3),
	6663:  uint16(1),
	6664:  uint16(anon_sym_COMMA),
	6665:  uint16(5),
	6666:  uint16(1),
	6667:  uint16(sym_comment),
	6668:  uint16(7),
	6669:  uint16(1),
	6670:  uint16(anon_sym_SLASH_SLASH_SLASH),
	6671:  uint16(425),
	6672:  uint16(1),
	6673:  uint16(aux_sym_identifier_token1),
	6674:  uint16(555),
	6675:  uint16(1),
	6676:  uint16(anon_sym_RBRACE),
	6677:  uint16(151),
	6678:  uint16(1),
	6679:  uint16(sym_documentation_comment),
	6680:  uint16(160),
	6681:  uint16(1),
	6682:  uint16(aux_sym_operation_body_repeat1),
	6683:  uint16(230),
	6684:  uint16(1),
	6685:  uint16(aux_sym_operation_member_repeat1),
	6686:  uint16(266),
	6687:  uint16(1),
	6688:  uint16(sym_identifier),
	6689:  uint16(303),
	6690:  uint16(1),
	6691:  uint16(sym_operation_errors),
	6692:  uint16(304),
	6693:  uint16(1),
	6694:  uint16(sym_operation_member),
	6695:  uint16(9),
	6696:  uint16(3),
	6697:  uint16(1),
	6698:  uint16(anon_sym_COMMA),
	6699:  uint16(5),
	6700:  uint16(1),
	6701:  uint16(sym_comment),
	6702:  uint16(7),
	6703:  uint16(1),
	6704:  uint16(anon_sym_SLASH_SLASH_SLASH),
	6705:  uint16(133),
	6706:  uint16(1),
	6707:  uint16(anon_sym_DQUOTE),
	6708:  uint16(135),
	6709:  uint16(1),
	6710:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	6711:  uint16(425),
	6712:  uint16(1),
	6713:  uint16(aux_sym_identifier_token1),
	6714:  uint16(152),
	6715:  uint16(1),
	6716:  uint16(sym_documentation_comment),
	6717:  uint16(168),
	6718:  uint16(2),
	6719:  uint16(sym__string_literal),
	6720:  uint16(sym__multiline_string_literal),
	6721:  uint16(344),
	6722:  uint16(2),
	6723:  uint16(sym_string),
	6724:  uint16(sym_identifier),
	6725:  uint16(10),
	6726:  uint16(3),
	6727:  uint16(1),
	6728:  uint16(anon_sym_COMMA),
	6729:  uint16(5),
	6730:  uint16(1),
	6731:  uint16(sym_comment),
	6732:  uint16(7),
	6733:  uint16(1),
	6734:  uint16(anon_sym_SLASH_SLASH_SLASH),
	6735:  uint16(557),
	6736:  uint16(1),
	6737:  uint16(aux_sym_identifier_token1),
	6738:  uint16(9),
	6739:  uint16(1),
	6740:  uint16(sym_root_shape_id),
	6741:  uint16(23),
	6742:  uint16(1),
	6743:  uint16(sym__namespace_identifier),
	6744:  uint16(153),
	6745:  uint16(1),
	6746:  uint16(sym_documentation_comment),
	6747:  uint16(307),
	6748:  uint16(1),
	6749:  uint16(sym_shape_id),
	6750:  uint16(332),
	6751:  uint16(1),
	6752:  uint16(sym_namespace),
	6753:  uint16(11),
	6754:  uint16(2),
	6755:  uint16(sym_absolute_root_shape_id),
	6756:  uint16(sym_identifier),
	6757:  uint16(5),
	6758:  uint16(3),
	6759:  uint16(1),
	6760:  uint16(anon_sym_COMMA),
	6761:  uint16(5),
	6762:  uint16(1),
	6763:  uint16(sym_comment),
	6764:  uint16(7),
	6765:  uint16(1),
	6766:  uint16(anon_sym_SLASH_SLASH_SLASH),
	6767:  uint16(154),
	6768:  uint16(1),
	6769:  uint16(sym_documentation_comment),
	6770:  uint16(381),
	6771:  uint16(7),
	6773:  uint16(anon_sym_DOLLAR),
	6774:  uint16(anon_sym_COLON),
	6775:  uint16(anon_sym_metadata),
	6776:  uint16(anon_sym_EQ),
	6777:  uint16(anon_sym_namespace),
	6778:  uint16(anon_sym_RPAREN),
	6779:  uint16(9),
	6780:  uint16(3),
	6781:  uint16(1),
	6782:  uint16(anon_sym_COMMA),
	6783:  uint16(5),
	6784:  uint16(1),
	6785:  uint16(sym_comment),
	6786:  uint16(7),
	6787:  uint16(1),
	6788:  uint16(anon_sym_SLASH_SLASH_SLASH),
	6789:  uint16(11),
	6790:  uint16(1),
	6791:  uint16(anon_sym_DOLLAR),
	6792:  uint16(79),
	6793:  uint16(1),
	6794:  uint16(sym_control_var_name),
	6795:  uint16(155),
	6796:  uint16(1),
	6797:  uint16(sym_documentation_comment),
	6798:  uint16(156),
	6799:  uint16(1),
	6800:  uint16(aux_sym_control_section_repeat1),
	6801:  uint16(239),
	6802:  uint16(1),
	6803:  uint16(sym_control_statement),
	6804:  uint16(559),
	6805:  uint16(3),
	6807:  uint16(anon_sym_metadata),
	6808:  uint16(anon_sym_namespace),
	6809:  uint16(8),
	6810:  uint16(3),
	6811:  uint16(1),
	6812:  uint16(anon_sym_COMMA),
	6813:  uint16(5),
	6814:  uint16(1),
	6815:  uint16(sym_comment),
	6816:  uint16(7),
	6817:  uint16(1),
	6818:  uint16(anon_sym_SLASH_SLASH_SLASH),
	6819:  uint16(563),
	6820:  uint16(1),
	6821:  uint16(anon_sym_DOLLAR),
	6822:  uint16(79),
	6823:  uint16(1),
	6824:  uint16(sym_control_var_name),
	6825:  uint16(239),
	6826:  uint16(1),
	6827:  uint16(sym_control_statement),
	6828:  uint16(156),
	6829:  uint16(2),
	6830:  uint16(sym_documentation_comment),
	6831:  uint16(aux_sym_control_section_repeat1),
	6832:  uint16(561),
	6833:  uint16(3),
	6835:  uint16(anon_sym_metadata),
	6836:  uint16(anon_sym_namespace),
	6837:  uint16(10),
	6838:  uint16(3),
	6839:  uint16(1),
	6840:  uint16(anon_sym_COMMA),
	6841:  uint16(5),
	6842:  uint16(1),
	6843:  uint16(sym_comment),
	6844:  uint16(7),
	6845:  uint16(1),
	6846:  uint16(anon_sym_SLASH_SLASH_SLASH),
	6847:  uint16(490),
	6848:  uint16(1),
	6849:  uint16(aux_sym_identifier_token1),
	6850:  uint16(23),
	6851:  uint16(1),
	6852:  uint16(sym__namespace_identifier),
	6853:  uint16(143),
	6854:  uint16(1),
	6855:  uint16(sym_root_shape_id),
	6856:  uint16(157),
	6857:  uint16(1),
	6858:  uint16(sym_documentation_comment),
	6859:  uint16(215),
	6860:  uint16(1),
	6861:  uint16(sym_shape_id),
	6862:  uint16(345),
	6863:  uint16(1),
	6864:  uint16(sym_namespace),
	6865:  uint16(130),
	6866:  uint16(2),
	6867:  uint16(sym_absolute_root_shape_id),
	6868:  uint16(sym_identifier),
	6869:  uint16(11),
	6870:  uint16(3),
	6871:  uint16(1),
	6872:  uint16(anon_sym_COMMA),
	6873:  uint16(5),
	6874:  uint16(1),
	6875:  uint16(sym_comment),
	6876:  uint16(7),
	6877:  uint16(1),
	6878:  uint16(anon_sym_SLASH_SLASH_SLASH),
	6879:  uint16(425),
	6880:  uint16(1),
	6881:  uint16(aux_sym_identifier_token1),
	6882:  uint16(566),
	6883:  uint16(1),
	6884:  uint16(anon_sym_RBRACE),
	6885:  uint16(151),
	6886:  uint16(1),
	6887:  uint16(aux_sym_operation_body_repeat1),
	6888:  uint16(158),
	6889:  uint16(1),
	6890:  uint16(sym_documentation_comment),
	6891:  uint16(230),
	6892:  uint16(1),
	6893:  uint16(aux_sym_operation_member_repeat1),
	6894:  uint16(266),
	6895:  uint16(1),
	6896:  uint16(sym_identifier),
	6897:  uint16(303),
	6898:  uint16(1),
	6899:  uint16(sym_operation_errors),
	6900:  uint16(304),
	6901:  uint16(1),
	6902:  uint16(sym_operation_member),
	6903:  uint16(10),
	6904:  uint16(3),
	6905:  uint16(1),
	6906:  uint16(anon_sym_COMMA),
	6907:  uint16(5),
	6908:  uint16(1),
	6909:  uint16(sym_comment),
	6910:  uint16(7),
	6911:  uint16(1),
	6912:  uint16(anon_sym_SLASH_SLASH_SLASH),
	6913:  uint16(490),
	6914:  uint16(1),
	6915:  uint16(aux_sym_identifier_token1),
	6916:  uint16(23),
	6917:  uint16(1),
	6918:  uint16(sym__namespace_identifier),
	6919:  uint16(143),
	6920:  uint16(1),
	6921:  uint16(sym_root_shape_id),
	6922:  uint16(159),
	6923:  uint16(1),
	6924:  uint16(sym_documentation_comment),
	6925:  uint16(190),
	6926:  uint16(1),
	6927:  uint16(sym_shape_id),
	6928:  uint16(345),
	6929:  uint16(1),
	6930:  uint16(sym_namespace),
	6931:  uint16(130),
	6932:  uint16(2),
	6933:  uint16(sym_absolute_root_shape_id),
	6934:  uint16(sym_identifier),
	6935:  uint16(10),
	6936:  uint16(3),
	6937:  uint16(1),
	6938:  uint16(anon_sym_COMMA),
	6939:  uint16(5),
	6940:  uint16(1),
	6941:  uint16(sym_comment),
	6942:  uint16(7),
	6943:  uint16(1),
	6944:  uint16(anon_sym_SLASH_SLASH_SLASH),
	6945:  uint16(568),
	6946:  uint16(1),
	6947:  uint16(anon_sym_RBRACE),
	6948:  uint16(570),
	6949:  uint16(1),
	6950:  uint16(aux_sym_identifier_token1),
	6951:  uint16(230),
	6952:  uint16(1),
	6953:  uint16(aux_sym_operation_member_repeat1),
	6954:  uint16(266),
	6955:  uint16(1),
	6956:  uint16(sym_identifier),
	6957:  uint16(303),
	6958:  uint16(1),
	6959:  uint16(sym_operation_errors),
	6960:  uint16(304),
	6961:  uint16(1),
	6962:  uint16(sym_operation_member),
	6963:  uint16(160),
	6964:  uint16(2),
	6965:  uint16(sym_documentation_comment),
	6966:  uint16(aux_sym_operation_body_repeat1),
	6967:  uint16(5),
	6968:  uint16(3),
	6969:  uint16(1),
	6970:  uint16(anon_sym_COMMA),
	6971:  uint16(5),
	6972:  uint16(1),
	6973:  uint16(sym_comment),
	6974:  uint16(7),
	6975:  uint16(1),
	6976:  uint16(anon_sym_SLASH_SLASH_SLASH),
	6977:  uint16(161),
	6978:  uint16(1),
	6979:  uint16(sym_documentation_comment),
	6980:  uint16(409),
	6981:  uint16(7),
	6983:  uint16(anon_sym_DOLLAR),
	6984:  uint16(anon_sym_COLON),
	6985:  uint16(anon_sym_metadata),
	6986:  uint16(anon_sym_EQ),
	6987:  uint16(anon_sym_namespace),
	6988:  uint16(anon_sym_RPAREN),
	6989:  uint16(5),
	6990:  uint16(3),
	6991:  uint16(1),
	6992:  uint16(anon_sym_COMMA),
	6993:  uint16(5),
	6994:  uint16(1),
	6995:  uint16(sym_comment),
	6996:  uint16(7),
	6997:  uint16(1),
	6998:  uint16(anon_sym_SLASH_SLASH_SLASH),
	6999:  uint16(162),
	7000:  uint16(1),
	7001:  uint16(sym_documentation_comment),
	7002:  uint16(385),
	7003:  uint16(7),
	7005:  uint16(anon_sym_DOLLAR),
	7006:  uint16(anon_sym_COLON),
	7007:  uint16(anon_sym_metadata),
	7008:  uint16(anon_sym_EQ),
	7009:  uint16(anon_sym_namespace),
	7010:  uint16(anon_sym_RPAREN),
	7011:  uint16(10),
	7012:  uint16(3),
	7013:  uint16(1),
	7014:  uint16(anon_sym_COMMA),
	7015:  uint16(5),
	7016:  uint16(1),
	7017:  uint16(sym_comment),
	7018:  uint16(7),
	7019:  uint16(1),
	7020:  uint16(anon_sym_SLASH_SLASH_SLASH),
	7021:  uint16(557),
	7022:  uint16(1),
	7023:  uint16(aux_sym_identifier_token1),
	7024:  uint16(9),
	7025:  uint16(1),
	7026:  uint16(sym_root_shape_id),
	7027:  uint16(23),
	7028:  uint16(1),
	7029:  uint16(sym__namespace_identifier),
	7030:  uint16(163),
	7031:  uint16(1),
	7032:  uint16(sym_documentation_comment),
	7033:  uint16(272),
	7034:  uint16(1),
	7035:  uint16(sym_shape_id),
	7036:  uint16(332),
	7037:  uint16(1),
	7038:  uint16(sym_namespace),
	7039:  uint16(11),
	7040:  uint16(2),
	7041:  uint16(sym_absolute_root_shape_id),
	7042:  uint16(sym_identifier),
	7043:  uint16(10),
	7044:  uint16(3),
	7045:  uint16(1),
	7046:  uint16(anon_sym_COMMA),
	7047:  uint16(5),
	7048:  uint16(1),
	7049:  uint16(sym_comment),
	7050:  uint16(7),
	7051:  uint16(1),
	7052:  uint16(anon_sym_SLASH_SLASH_SLASH),
	7053:  uint16(490),
	7054:  uint16(1),
	7055:  uint16(aux_sym_identifier_token1),
	7056:  uint16(23),
	7057:  uint16(1),
	7058:  uint16(sym__namespace_identifier),
	7059:  uint16(143),
	7060:  uint16(1),
	7061:  uint16(sym_root_shape_id),
	7062:  uint16(164),
	7063:  uint16(1),
	7064:  uint16(sym_documentation_comment),
	7065:  uint16(188),
	7066:  uint16(1),
	7067:  uint16(sym_shape_id),
	7068:  uint16(345),
	7069:  uint16(1),
	7070:  uint16(sym_namespace),
	7071:  uint16(130),
	7072:  uint16(2),
	7073:  uint16(sym_absolute_root_shape_id),
	7074:  uint16(sym_identifier),
	7075:  uint16(5),
	7076:  uint16(3),
	7077:  uint16(1),
	7078:  uint16(anon_sym_COMMA),
	7079:  uint16(5),
	7080:  uint16(1),
	7081:  uint16(sym_comment),
	7082:  uint16(7),
	7083:  uint16(1),
	7084:  uint16(anon_sym_SLASH_SLASH_SLASH),
	7085:  uint16(165),
	7086:  uint16(1),
	7087:  uint16(sym_documentation_comment),
	7088:  uint16(413),
	7089:  uint16(7),
	7091:  uint16(anon_sym_DOLLAR),
	7092:  uint16(anon_sym_COLON),
	7093:  uint16(anon_sym_metadata),
	7094:  uint16(anon_sym_EQ),
	7095:  uint16(anon_sym_namespace),
	7096:  uint16(anon_sym_RPAREN),
	7097:  uint16(10),
	7098:  uint16(3),
	7099:  uint16(1),
	7100:  uint16(anon_sym_COMMA),
	7101:  uint16(5),
	7102:  uint16(1),
	7103:  uint16(sym_comment),
	7104:  uint16(7),
	7105:  uint16(1),
	7106:  uint16(anon_sym_SLASH_SLASH_SLASH),
	7107:  uint16(490),
	7108:  uint16(1),
	7109:  uint16(aux_sym_identifier_token1),
	7110:  uint16(23),
	7111:  uint16(1),
	7112:  uint16(sym__namespace_identifier),
	7113:  uint16(120),
	7114:  uint16(1),
	7115:  uint16(sym_root_shape_id),
	7116:  uint16(166),
	7117:  uint16(1),
	7118:  uint16(sym_documentation_comment),
	7119:  uint16(215),
	7120:  uint16(1),
	7121:  uint16(sym_shape_id),
	7122:  uint16(345),
	7123:  uint16(1),
	7124:  uint16(sym_namespace),
	7125:  uint16(130),
	7126:  uint16(2),
	7127:  uint16(sym_absolute_root_shape_id),
	7128:  uint16(sym_identifier),
	7129:  uint16(9),
	7130:  uint16(3),
	7131:  uint16(1),
	7132:  uint16(anon_sym_COMMA),
	7133:  uint16(5),
	7134:  uint16(1),
	7135:  uint16(sym_comment),
	7136:  uint16(7),
	7137:  uint16(1),
	7138:  uint16(anon_sym_SLASH_SLASH_SLASH),
	7139:  uint16(133),
	7140:  uint16(1),
	7141:  uint16(anon_sym_DQUOTE),
	7142:  uint16(135),
	7143:  uint16(1),
	7144:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	7145:  uint16(573),
	7146:  uint16(1),
	7147:  uint16(aux_sym_identifier_token1),
	7148:  uint16(167),
	7149:  uint16(1),
	7150:  uint16(sym_documentation_comment),
	7151:  uint16(168),
	7152:  uint16(2),
	7153:  uint16(sym__string_literal),
	7154:  uint16(sym__multiline_string_literal),
	7155:  uint16(346),
	7156:  uint16(2),
	7157:  uint16(sym_string),
	7158:  uint16(sym__control_identifier),
	7159:  uint16(5),
	7160:  uint16(3),
	7161:  uint16(1),
	7162:  uint16(anon_sym_COMMA),
	7163:  uint16(5),
	7164:  uint16(1),
	7165:  uint16(sym_comment),
	7166:  uint16(7),
	7167:  uint16(1),
	7168:  uint16(anon_sym_SLASH_SLASH_SLASH),
	7169:  uint16(168),
	7170:  uint16(1),
	7171:  uint16(sym_documentation_comment),
	7172:  uint16(358),
	7173:  uint16(7),
	7175:  uint16(anon_sym_DOLLAR),
	7176:  uint16(anon_sym_COLON),
	7177:  uint16(anon_sym_metadata),
	7178:  uint16(anon_sym_EQ),
	7179:  uint16(anon_sym_namespace),
	7180:  uint16(anon_sym_RPAREN),
	7181:  uint16(6),
	7182:  uint16(3),
	7183:  uint16(1),
	7184:  uint16(anon_sym_COMMA),
	7185:  uint16(5),
	7186:  uint16(1),
	7187:  uint16(sym_comment),
	7188:  uint16(7),
	7189:  uint16(1),
	7190:  uint16(anon_sym_SLASH_SLASH_SLASH),
	7191:  uint16(360),
	7192:  uint16(1),
	7193:  uint16(anon_sym_DQUOTE),
	7194:  uint16(169),
	7195:  uint16(1),
	7196:  uint16(sym_documentation_comment),
	7197:  uint16(358),
	7198:  uint16(6),
	7199:  uint16(anon_sym_DOLLAR),
	7200:  uint16(anon_sym_RBRACE),
	7201:  uint16(anon_sym_AT),
	7202:  uint16(anon_sym_RPAREN),
	7203:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	7204:  uint16(aux_sym_identifier_token1),
	7205:  uint16(6),
	7206:  uint16(3),
	7207:  uint16(1),
	7208:  uint16(anon_sym_COMMA),
	7209:  uint16(5),
	7210:  uint16(1),
	7211:  uint16(sym_comment),
	7212:  uint16(7),
	7213:  uint16(1),
	7214:  uint16(anon_sym_SLASH_SLASH_SLASH),
	7215:  uint16(401),
	7216:  uint16(1),
	7217:  uint16(anon_sym_DQUOTE),
	7218:  uint16(170),
	7219:  uint16(1),
	7220:  uint16(sym_documentation_comment),
	7221:  uint16(399),
	7222:  uint16(6),
	7223:  uint16(anon_sym_DOLLAR),
	7224:  uint16(anon_sym_RBRACE),
	7225:  uint16(anon_sym_AT),
	7226:  uint16(anon_sym_RPAREN),
	7227:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	7228:  uint16(aux_sym_identifier_token1),
	7229:  uint16(6),
	7230:  uint16(3),
	7231:  uint16(1),
	7232:  uint16(anon_sym_COMMA),
	7233:  uint16(5),
	7234:  uint16(1),
	7235:  uint16(sym_comment),
	7236:  uint16(7),
	7237:  uint16(1),
	7238:  uint16(anon_sym_SLASH_SLASH_SLASH),
	7239:  uint16(407),
	7240:  uint16(1),
	7241:  uint16(anon_sym_DQUOTE),
	7242:  uint16(171),
	7243:  uint16(1),
	7244:  uint16(sym_documentation_comment),
	7245:  uint16(141),
	7246:  uint16(6),
	7247:  uint16(anon_sym_DOLLAR),
	7248:  uint16(anon_sym_RBRACE),
	7249:  uint16(anon_sym_AT),
	7250:  uint16(anon_sym_RPAREN),
	7251:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	7252:  uint16(aux_sym_identifier_token1),
	7253:  uint16(6),
	7254:  uint16(3),
	7255:  uint16(1),
	7256:  uint16(anon_sym_COMMA),
	7257:  uint16(5),
	7258:  uint16(1),
	7259:  uint16(sym_comment),
	7260:  uint16(7),
	7261:  uint16(1),
	7262:  uint16(anon_sym_SLASH_SLASH_SLASH),
	7263:  uint16(364),
	7264:  uint16(1),
	7265:  uint16(anon_sym_DQUOTE),
	7266:  uint16(172),
	7267:  uint16(1),
	7268:  uint16(sym_documentation_comment),
	7269:  uint16(362),
	7270:  uint16(6),
	7271:  uint16(anon_sym_DOLLAR),
	7272:  uint16(anon_sym_RBRACE),
	7273:  uint16(anon_sym_AT),
	7274:  uint16(anon_sym_RPAREN),
	7275:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	7276:  uint16(aux_sym_identifier_token1),
	7277:  uint16(6),
	7278:  uint16(3),
	7279:  uint16(1),
	7280:  uint16(anon_sym_COMMA),
	7281:  uint16(5),
	7282:  uint16(1),
	7283:  uint16(sym_comment),
	7284:  uint16(7),
	7285:  uint16(1),
	7286:  uint16(anon_sym_SLASH_SLASH_SLASH),
	7287:  uint16(411),
	7288:  uint16(1),
	7289:  uint16(anon_sym_DQUOTE),
	7290:  uint16(173),
	7291:  uint16(1),
	7292:  uint16(sym_documentation_comment),
	7293:  uint16(409),
	7294:  uint16(6),
	7295:  uint16(anon_sym_DOLLAR),
	7296:  uint16(anon_sym_RBRACE),
	7297:  uint16(anon_sym_AT),
	7298:  uint16(anon_sym_RPAREN),
	7299:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	7300:  uint16(aux_sym_identifier_token1),
	7301:  uint16(6),
	7302:  uint16(3),
	7303:  uint16(1),
	7304:  uint16(anon_sym_COMMA),
	7305:  uint16(5),
	7306:  uint16(1),
	7307:  uint16(sym_comment),
	7308:  uint16(7),
	7309:  uint16(1),
	7310:  uint16(anon_sym_SLASH_SLASH_SLASH),
	7311:  uint16(356),
	7312:  uint16(1),
	7313:  uint16(anon_sym_DQUOTE),
	7314:  uint16(174),
	7315:  uint16(1),
	7316:  uint16(sym_documentation_comment),
	7317:  uint16(354),
	7318:  uint16(6),
	7319:  uint16(anon_sym_DOLLAR),
	7320:  uint16(anon_sym_RBRACE),
	7321:  uint16(anon_sym_AT),
	7322:  uint16(anon_sym_RPAREN),
	7323:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	7324:  uint16(aux_sym_identifier_token1),
	7325:  uint16(6),
	7326:  uint16(3),
	7327:  uint16(1),
	7328:  uint16(anon_sym_COMMA),
	7329:  uint16(5),
	7330:  uint16(1),
	7331:  uint16(sym_comment),
	7332:  uint16(7),
	7333:  uint16(1),
	7334:  uint16(anon_sym_SLASH_SLASH_SLASH),
	7335:  uint16(415),
	7336:  uint16(1),
	7337:  uint16(anon_sym_DQUOTE),
	7338:  uint16(175),
	7339:  uint16(1),
	7340:  uint16(sym_documentation_comment),
	7341:  uint16(413),
	7342:  uint16(6),
	7343:  uint16(anon_sym_DOLLAR),
	7344:  uint16(anon_sym_RBRACE),
	7345:  uint16(anon_sym_AT),
	7346:  uint16(anon_sym_RPAREN),
	7347:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	7348:  uint16(aux_sym_identifier_token1),
	7349:  uint16(11),
	7350:  uint16(3),
	7351:  uint16(1),
	7352:  uint16(anon_sym_COMMA),
	7353:  uint16(5),
	7354:  uint16(1),
	7355:  uint16(sym_comment),
	7356:  uint16(7),
	7357:  uint16(1),
	7358:  uint16(anon_sym_SLASH_SLASH_SLASH),
	7359:  uint16(425),
	7360:  uint16(1),
	7361:  uint16(aux_sym_identifier_token1),
	7362:  uint16(465),
	7363:  uint16(1),
	7364:  uint16(anon_sym_DOLLAR),
	7365:  uint16(469),
	7366:  uint16(1),
	7367:  uint16(anon_sym_AT),
	7368:  uint16(176),
	7369:  uint16(1),
	7370:  uint16(sym_documentation_comment),
	7371:  uint16(202),
	7372:  uint16(1),
	7373:  uint16(sym_shape_member_elided),
	7374:  uint16(223),
	7375:  uint16(1),
	7376:  uint16(aux_sym_shape_statement_repeat1),
	7377:  uint16(274),
	7378:  uint16(1),
	7379:  uint16(sym_trait_statement),
	7380:  uint16(348),
	7381:  uint16(1),
	7382:  uint16(sym_identifier),
	7383:  uint16(6),
	7384:  uint16(3),
	7385:  uint16(1),
	7386:  uint16(anon_sym_COMMA),
	7387:  uint16(5),
	7388:  uint16(1),
	7389:  uint16(sym_comment),
	7390:  uint16(7),
	7391:  uint16(1),
	7392:  uint16(anon_sym_SLASH_SLASH_SLASH),
	7393:  uint16(368),
	7394:  uint16(1),
	7395:  uint16(anon_sym_DQUOTE),
	7396:  uint16(177),
	7397:  uint16(1),
	7398:  uint16(sym_documentation_comment),
	7399:  uint16(366),
	7400:  uint16(6),
	7401:  uint16(anon_sym_DOLLAR),
	7402:  uint16(anon_sym_RBRACE),
	7403:  uint16(anon_sym_AT),
	7404:  uint16(anon_sym_RPAREN),
	7405:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	7406:  uint16(aux_sym_identifier_token1),
	7407:  uint16(6),
	7408:  uint16(3),
	7409:  uint16(1),
	7410:  uint16(anon_sym_COMMA),
	7411:  uint16(5),
	7412:  uint16(1),
	7413:  uint16(sym_comment),
	7414:  uint16(7),
	7415:  uint16(1),
	7416:  uint16(anon_sym_SLASH_SLASH_SLASH),
	7417:  uint16(419),
	7418:  uint16(1),
	7419:  uint16(anon_sym_DQUOTE),
	7420:  uint16(178),
	7421:  uint16(1),
	7422:  uint16(sym_documentation_comment),
	7423:  uint16(417),
	7424:  uint16(6),
	7425:  uint16(anon_sym_DOLLAR),
	7426:  uint16(anon_sym_RBRACE),
	7427:  uint16(anon_sym_AT),
	7428:  uint16(anon_sym_RPAREN),
	7429:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	7430:  uint16(aux_sym_identifier_token1),
	7431:  uint16(6),
	7432:  uint16(3),
	7433:  uint16(1),
	7434:  uint16(anon_sym_COMMA),
	7435:  uint16(5),
	7436:  uint16(1),
	7437:  uint16(sym_comment),
	7438:  uint16(7),
	7439:  uint16(1),
	7440:  uint16(anon_sym_SLASH_SLASH_SLASH),
	7441:  uint16(375),
	7442:  uint16(1),
	7443:  uint16(anon_sym_DQUOTE),
	7444:  uint16(179),
	7445:  uint16(1),
	7446:  uint16(sym_documentation_comment),
	7447:  uint16(373),
	7448:  uint16(6),
	7449:  uint16(anon_sym_DOLLAR),
	7450:  uint16(anon_sym_RBRACE),
	7451:  uint16(anon_sym_AT),
	7452:  uint16(anon_sym_RPAREN),
	7453:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	7454:  uint16(aux_sym_identifier_token1),
	7455:  uint16(6),
	7456:  uint16(3),
	7457:  uint16(1),
	7458:  uint16(anon_sym_COMMA),
	7459:  uint16(5),
	7460:  uint16(1),
	7461:  uint16(sym_comment),
	7462:  uint16(7),
	7463:  uint16(1),
	7464:  uint16(anon_sym_SLASH_SLASH_SLASH),
	7465:  uint16(397),
	7466:  uint16(1),
	7467:  uint16(anon_sym_DQUOTE),
	7468:  uint16(180),
	7469:  uint16(1),
	7470:  uint16(sym_documentation_comment),
	7471:  uint16(395),
	7472:  uint16(6),
	7473:  uint16(anon_sym_DOLLAR),
	7474:  uint16(anon_sym_RBRACE),
	7475:  uint16(anon_sym_AT),
	7476:  uint16(anon_sym_RPAREN),
	7477:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	7478:  uint16(aux_sym_identifier_token1),
	7479:  uint16(10),
	7480:  uint16(3),
	7481:  uint16(1),
	7482:  uint16(anon_sym_COMMA),
	7483:  uint16(5),
	7484:  uint16(1),
	7485:  uint16(sym_comment),
	7486:  uint16(7),
	7487:  uint16(1),
	7488:  uint16(anon_sym_SLASH_SLASH_SLASH),
	7489:  uint16(557),
	7490:  uint16(1),
	7491:  uint16(aux_sym_identifier_token1),
	7492:  uint16(9),
	7493:  uint16(1),
	7494:  uint16(sym_root_shape_id),
	7495:  uint16(15),
	7496:  uint16(1),
	7497:  uint16(sym_shape_id),
	7498:  uint16(23),
	7499:  uint16(1),
	7500:  uint16(sym__namespace_identifier),
	7501:  uint16(181),
	7502:  uint16(1),
	7503:  uint16(sym_documentation_comment),
	7504:  uint16(332),
	7505:  uint16(1),
	7506:  uint16(sym_namespace),
	7507:  uint16(11),
	7508:  uint16(2),
	7509:  uint16(sym_absolute_root_shape_id),
	7510:  uint16(sym_identifier),
	7511:  uint16(6),
	7512:  uint16(3),
	7513:  uint16(1),
	7514:  uint16(anon_sym_COMMA),
	7515:  uint16(5),
	7516:  uint16(1),
	7517:  uint16(sym_comment),
	7518:  uint16(7),
	7519:  uint16(1),
	7520:  uint16(anon_sym_SLASH_SLASH_SLASH),
	7521:  uint16(393),
	7522:  uint16(1),
	7523:  uint16(anon_sym_DQUOTE),
	7524:  uint16(182),
	7525:  uint16(1),
	7526:  uint16(sym_documentation_comment),
	7527:  uint16(391),
	7528:  uint16(6),
	7529:  uint16(anon_sym_DOLLAR),
	7530:  uint16(anon_sym_RBRACE),
	7531:  uint16(anon_sym_AT),
	7532:  uint16(anon_sym_RPAREN),
	7533:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	7534:  uint16(aux_sym_identifier_token1),
	7535:  uint16(10),
	7536:  uint16(3),
	7537:  uint16(1),
	7538:  uint16(anon_sym_COMMA),
	7539:  uint16(5),
	7540:  uint16(1),
	7541:  uint16(sym_comment),
	7542:  uint16(7),
	7543:  uint16(1),
	7544:  uint16(anon_sym_SLASH_SLASH_SLASH),
	7545:  uint16(490),
	7546:  uint16(1),
	7547:  uint16(aux_sym_identifier_token1),
	7548:  uint16(23),
	7549:  uint16(1),
	7550:  uint16(sym__namespace_identifier),
	7551:  uint16(120),
	7552:  uint16(1),
	7553:  uint16(sym_root_shape_id),
	7554:  uint16(183),
	7555:  uint16(1),
	7556:  uint16(sym_documentation_comment),
	7557:  uint16(291),
	7558:  uint16(1),
	7559:  uint16(sym_shape_id),
	7560:  uint16(345),
	7561:  uint16(1),
	7562:  uint16(sym_namespace),
	7563:  uint16(130),
	7564:  uint16(2),
	7565:  uint16(sym_absolute_root_shape_id),
	7566:  uint16(sym_identifier),
	7567:  uint16(6),
	7568:  uint16(3),
	7569:  uint16(1),
	7570:  uint16(anon_sym_COMMA),
	7571:  uint16(5),
	7572:  uint16(1),
	7573:  uint16(sym_comment),
	7574:  uint16(7),
	7575:  uint16(1),
	7576:  uint16(anon_sym_SLASH_SLASH_SLASH),
	7577:  uint16(379),
	7578:  uint16(1),
	7579:  uint16(anon_sym_DQUOTE),
	7580:  uint16(184),
	7581:  uint16(1),
	7582:  uint16(sym_documentation_comment),
	7583:  uint16(377),
	7584:  uint16(6),
	7585:  uint16(anon_sym_DOLLAR),
	7586:  uint16(anon_sym_RBRACE),
	7587:  uint16(anon_sym_AT),
	7588:  uint16(anon_sym_RPAREN),
	7589:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	7590:  uint16(aux_sym_identifier_token1),
	7591:  uint16(6),
	7592:  uint16(3),
	7593:  uint16(1),
	7594:  uint16(anon_sym_COMMA),
	7595:  uint16(5),
	7596:  uint16(1),
	7597:  uint16(sym_comment),
	7598:  uint16(7),
	7599:  uint16(1),
	7600:  uint16(anon_sym_SLASH_SLASH_SLASH),
	7601:  uint16(383),
	7602:  uint16(1),
	7603:  uint16(anon_sym_DQUOTE),
	7604:  uint16(185),
	7605:  uint16(1),
	7606:  uint16(sym_documentation_comment),
	7607:  uint16(381),
	7608:  uint16(6),
	7609:  uint16(anon_sym_DOLLAR),
	7610:  uint16(anon_sym_RBRACE),
	7611:  uint16(anon_sym_AT),
	7612:  uint16(anon_sym_RPAREN),
	7613:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	7614:  uint16(aux_sym_identifier_token1),
	7615:  uint16(6),
	7616:  uint16(3),
	7617:  uint16(1),
	7618:  uint16(anon_sym_COMMA),
	7619:  uint16(5),
	7620:  uint16(1),
	7621:  uint16(sym_comment),
	7622:  uint16(7),
	7623:  uint16(1),
	7624:  uint16(anon_sym_SLASH_SLASH_SLASH),
	7625:  uint16(387),
	7626:  uint16(1),
	7627:  uint16(anon_sym_DQUOTE),
	7628:  uint16(186),
	7629:  uint16(1),
	7630:  uint16(sym_documentation_comment),
	7631:  uint16(385),
	7632:  uint16(6),
	7633:  uint16(anon_sym_DOLLAR),
	7634:  uint16(anon_sym_RBRACE),
	7635:  uint16(anon_sym_AT),
	7636:  uint16(anon_sym_RPAREN),
	7637:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	7638:  uint16(aux_sym_identifier_token1),
	7639:  uint16(5),
	7640:  uint16(7),
	7641:  uint16(1),
	7642:  uint16(anon_sym_SLASH_SLASH_SLASH),
	7643:  uint16(577),
	7644:  uint16(1),
	7645:  uint16(aux_sym__multiline_string_fragment_token1),
	7646:  uint16(5),
	7647:  uint16(2),
	7648:  uint16(anon_sym_COMMA),
	7649:  uint16(sym_comment),
	7650:  uint16(187),
	7651:  uint16(2),
	7652:  uint16(sym_documentation_comment),
	7653:  uint16(aux_sym__multiline_string_fragment_repeat1),
	7654:  uint16(575),
	7655:  uint16(4),
	7656:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	7657:  uint16(aux_sym__multiline_string_fragment_token2),
	7658:  uint16(aux_sym__escape_sequence_token1),
	7659:  uint16(sym_escape_sequence),
	7660:  uint16(7),
	7661:  uint16(3),
	7662:  uint16(1),
	7663:  uint16(anon_sym_COMMA),
	7664:  uint16(5),
	7665:  uint16(1),
	7666:  uint16(sym_comment),
	7667:  uint16(7),
	7668:  uint16(1),
	7669:  uint16(anon_sym_SLASH_SLASH_SLASH),
	7670:  uint16(582),
	7671:  uint16(1),
	7672:  uint16(anon_sym_EQ),
	7673:  uint16(188),
	7674:  uint16(1),
	7675:  uint16(sym_documentation_comment),
	7676:  uint16(240),
	7677:  uint16(1),
	7678:  uint16(sym_value_assignment),
	7679:  uint16(580),
	7680:  uint16(4),
	7681:  uint16(anon_sym_DOLLAR),
	7682:  uint16(anon_sym_RBRACE),
	7683:  uint16(anon_sym_AT),
	7684:  uint16(aux_sym_identifier_token1),
	7685:  uint16(9),
	7686:  uint16(7),
	7687:  uint16(1),
	7688:  uint16(anon_sym_SLASH_SLASH_SLASH),
	7689:  uint16(584),
	7690:  uint16(1),
	7691:  uint16(anon_sym_DQUOTE),
	7692:  uint16(586),
	7693:  uint16(1),
	7694:  uint16(sym_string_fragment),
	7695:  uint16(588),
	7696:  uint16(1),
	7697:  uint16(aux_sym__escape_sequence_token1),
	7698:  uint16(590),
	7699:  uint16(1),
	7700:  uint16(sym_escape_sequence),
	7701:  uint16(189),
	7702:  uint16(1),
	7703:  uint16(sym_documentation_comment),
	7704:  uint16(198),
	7705:  uint16(1),
	7706:  uint16(aux_sym__string_literal_repeat1),
	7707:  uint16(235),
	7708:  uint16(1),
	7709:  uint16(sym__escape_sequence),
	7710:  uint16(5),
	7711:  uint16(2),
	7712:  uint16(anon_sym_COMMA),
	7713:  uint16(sym_comment),
	7714:  uint16(7),
	7715:  uint16(3),
	7716:  uint16(1),
	7717:  uint16(anon_sym_COMMA),
	7718:  uint16(5),
	7719:  uint16(1),
	7720:  uint16(sym_comment),
	7721:  uint16(7),
	7722:  uint16(1),
	7723:  uint16(anon_sym_SLASH_SLASH_SLASH),
	7724:  uint16(582),
	7725:  uint16(1),
	7726:  uint16(anon_sym_EQ),
	7727:  uint16(190),
	7728:  uint16(1),
	7729:  uint16(sym_documentation_comment),
	7730:  uint16(248),
	7731:  uint16(1),
	7732:  uint16(sym_value_assignment),
	7733:  uint16(592),
	7734:  uint16(4),
	7735:  uint16(anon_sym_DOLLAR),
	7736:  uint16(anon_sym_RBRACE),
	7737:  uint16(anon_sym_AT),
	7738:  uint16(aux_sym_identifier_token1),
	7739:  uint16(7),
	7740:  uint16(3),
	7741:  uint16(1),
	7742:  uint16(anon_sym_COMMA),
	7743:  uint16(5),
	7744:  uint16(1),
	7745:  uint16(sym_comment),
	7746:  uint16(7),
	7747:  uint16(1),
	7748:  uint16(anon_sym_SLASH_SLASH_SLASH),
	7749:  uint16(12),
	7750:  uint16(1),
	7751:  uint16(sym_shape_id_member),
	7752:  uint16(191),
	7753:  uint16(1),
	7754:  uint16(sym_documentation_comment),
	7755:  uint16(199),
	7756:  uint16(1),
	7757:  uint16(aux_sym_shape_id_repeat1),
	7758:  uint16(103),
	7759:  uint16(4),
	7761:  uint16(anon_sym_DOLLAR),
	7762:  uint16(anon_sym_metadata),
	7763:  uint16(anon_sym_namespace),
	7764:  uint16(7),
	7765:  uint16(3),
	7766:  uint16(1),
	7767:  uint16(anon_sym_COMMA),
	7768:  uint16(5),
	7769:  uint16(1),
	7770:  uint16(sym_comment),
	7771:  uint16(7),
	7772:  uint16(1),
	7773:  uint16(anon_sym_SLASH_SLASH_SLASH),
	7774:  uint16(582),
	7775:  uint16(1),
	7776:  uint16(anon_sym_EQ),
	7777:  uint16(192),
	7778:  uint16(1),
	7779:  uint16(sym_documentation_comment),
	7780:  uint16(252),
	7781:  uint16(1),
	7782:  uint16(sym_value_assignment),
	7783:  uint16(594),
	7784:  uint16(4),
	7785:  uint16(anon_sym_DOLLAR),
	7786:  uint16(anon_sym_RBRACE),
	7787:  uint16(anon_sym_AT),
	7788:  uint16(aux_sym_identifier_token1),
	7789:  uint16(9),
	7790:  uint16(7),
	7791:  uint16(1),
	7792:  uint16(anon_sym_SLASH_SLASH_SLASH),
	7793:  uint16(586),
	7794:  uint16(1),
	7795:  uint16(sym_string_fragment),
	7796:  uint16(588),
	7797:  uint16(1),
	7798:  uint16(aux_sym__escape_sequence_token1),
	7799:  uint16(590),
	7800:  uint16(1),
	7801:  uint16(sym_escape_sequence),
	7802:  uint16(596),
	7803:  uint16(1),
	7804:  uint16(anon_sym_DQUOTE),
	7805:  uint16(193),
	7806:  uint16(1),
	7807:  uint16(sym_documentation_comment),
	7808:  uint16(201),
	7809:  uint16(1),
	7810:  uint16(aux_sym__string_literal_repeat1),
	7811:  uint16(235),
	7812:  uint16(1),
	7813:  uint16(sym__escape_sequence),
	7814:  uint16(5),
	7815:  uint16(2),
	7816:  uint16(anon_sym_COMMA),
	7817:  uint16(sym_comment),
	7818:  uint16(9),
	7819:  uint16(7),
	7820:  uint16(1),
	7821:  uint16(anon_sym_SLASH_SLASH_SLASH),
	7822:  uint16(586),
	7823:  uint16(1),
	7824:  uint16(sym_string_fragment),
	7825:  uint16(588),
	7826:  uint16(1),
	7827:  uint16(aux_sym__escape_sequence_token1),
	7828:  uint16(590),
	7829:  uint16(1),
	7830:  uint16(sym_escape_sequence),
	7831:  uint16(598),
	7832:  uint16(1),
	7833:  uint16(anon_sym_DQUOTE),
	7834:  uint16(194),
	7835:  uint16(1),
	7836:  uint16(sym_documentation_comment),
	7837:  uint16(201),
	7838:  uint16(1),
	7839:  uint16(aux_sym__string_literal_repeat1),
	7840:  uint16(235),
	7841:  uint16(1),
	7842:  uint16(sym__escape_sequence),
	7843:  uint16(5),
	7844:  uint16(2),
	7845:  uint16(anon_sym_COMMA),
	7846:  uint16(sym_comment),
	7847:  uint16(9),
	7848:  uint16(7),
	7849:  uint16(1),
	7850:  uint16(anon_sym_SLASH_SLASH_SLASH),
	7851:  uint16(586),
	7852:  uint16(1),
	7853:  uint16(sym_string_fragment),
	7854:  uint16(588),
	7855:  uint16(1),
	7856:  uint16(aux_sym__escape_sequence_token1),
	7857:  uint16(590),
	7858:  uint16(1),
	7859:  uint16(sym_escape_sequence),
	7860:  uint16(600),
	7861:  uint16(1),
	7862:  uint16(anon_sym_DQUOTE),
	7863:  uint16(194),
	7864:  uint16(1),
	7865:  uint16(aux_sym__string_literal_repeat1),
	7866:  uint16(195),
	7867:  uint16(1),
	7868:  uint16(sym_documentation_comment),
	7869:  uint16(235),
	7870:  uint16(1),
	7871:  uint16(sym__escape_sequence),
	7872:  uint16(5),
	7873:  uint16(2),
	7874:  uint16(anon_sym_COMMA),
	7875:  uint16(sym_comment),
	7876:  uint16(9),
	7877:  uint16(7),
	7878:  uint16(1),
	7879:  uint16(anon_sym_SLASH_SLASH_SLASH),
	7880:  uint16(586),
	7881:  uint16(1),
	7882:  uint16(sym_string_fragment),
	7883:  uint16(588),
	7884:  uint16(1),
	7885:  uint16(aux_sym__escape_sequence_token1),
	7886:  uint16(590),
	7887:  uint16(1),
	7888:  uint16(sym_escape_sequence),
	7889:  uint16(602),
	7890:  uint16(1),
	7891:  uint16(anon_sym_DQUOTE),
	7892:  uint16(193),
	7893:  uint16(1),
	7894:  uint16(aux_sym__string_literal_repeat1),
	7895:  uint16(196),
	7896:  uint16(1),
	7897:  uint16(sym_documentation_comment),
	7898:  uint16(235),
	7899:  uint16(1),
	7900:  uint16(sym__escape_sequence),
	7901:  uint16(5),
	7902:  uint16(2),
	7903:  uint16(anon_sym_COMMA),
	7904:  uint16(sym_comment),
	7905:  uint16(6),
	7906:  uint16(7),
	7907:  uint16(1),
	7908:  uint16(anon_sym_SLASH_SLASH_SLASH),
	7909:  uint16(606),
	7910:  uint16(1),
	7911:  uint16(aux_sym__multiline_string_fragment_token1),
	7912:  uint16(187),
	7913:  uint16(1),
	7914:  uint16(aux_sym__multiline_string_fragment_repeat1),
	7915:  uint16(197),
	7916:  uint16(1),
	7917:  uint16(sym_documentation_comment),
	7918:  uint16(5),
	7919:  uint16(2),
	7920:  uint16(anon_sym_COMMA),
	7921:  uint16(sym_comment),
	7922:  uint16(604),
	7923:  uint16(4),
	7924:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	7925:  uint16(aux_sym__multiline_string_fragment_token2),
	7926:  uint16(aux_sym__escape_sequence_token1),
	7927:  uint16(sym_escape_sequence),
	7928:  uint16(9),
	7929:  uint16(7),
	7930:  uint16(1),
	7931:  uint16(anon_sym_SLASH_SLASH_SLASH),
	7932:  uint16(586),
	7933:  uint16(1),
	7934:  uint16(sym_string_fragment),
	7935:  uint16(588),
	7936:  uint16(1),
	7937:  uint16(aux_sym__escape_sequence_token1),
	7938:  uint16(590),
	7939:  uint16(1),
	7940:  uint16(sym_escape_sequence),
	7941:  uint16(608),
	7942:  uint16(1),
	7943:  uint16(anon_sym_DQUOTE),
	7944:  uint16(198),
	7945:  uint16(1),
	7946:  uint16(sym_documentation_comment),
	7947:  uint16(201),
	7948:  uint16(1),
	7949:  uint16(aux_sym__string_literal_repeat1),
	7950:  uint16(235),
	7951:  uint16(1),
	7952:  uint16(sym__escape_sequence),
	7953:  uint16(5),
	7954:  uint16(2),
	7955:  uint16(anon_sym_COMMA),
	7956:  uint16(sym_comment),
	7957:  uint16(7),
	7958:  uint16(3),
	7959:  uint16(1),
	7960:  uint16(anon_sym_COMMA),
	7961:  uint16(5),
	7962:  uint16(1),
	7963:  uint16(sym_comment),
	7964:  uint16(7),
	7965:  uint16(1),
	7966:  uint16(anon_sym_SLASH_SLASH_SLASH),
	7967:  uint16(7),
	7968:  uint16(1),
	7969:  uint16(aux_sym_shape_id_repeat1),
	7970:  uint16(12),
	7971:  uint16(1),
	7972:  uint16(sym_shape_id_member),
	7973:  uint16(199),
	7974:  uint16(1),
	7975:  uint16(sym_documentation_comment),
	7976:  uint16(99),
	7977:  uint16(4),
	7979:  uint16(anon_sym_DOLLAR),
	7980:  uint16(anon_sym_metadata),
	7981:  uint16(anon_sym_namespace),
	7982:  uint16(10),
	7983:  uint16(3),
	7984:  uint16(1),
	7985:  uint16(anon_sym_COMMA),
	7986:  uint16(5),
	7987:  uint16(1),
	7988:  uint16(sym_comment),
	7989:  uint16(7),
	7990:  uint16(1),
	7991:  uint16(anon_sym_SLASH_SLASH_SLASH),
	7992:  uint16(227),
	7993:  uint16(1),
	7994:  uint16(anon_sym_with),
	7995:  uint16(494),
	7996:  uint16(1),
	7997:  uint16(anon_sym_for),
	7998:  uint16(610),
	7999:  uint16(1),
	8000:  uint16(anon_sym_LBRACE),
	8001:  uint16(50),
	8002:  uint16(1),
	8003:  uint16(sym_shape_members),
	8004:  uint16(200),
	8005:  uint16(1),
	8006:  uint16(sym_documentation_comment),
	8007:  uint16(233),
	8008:  uint16(1),
	8009:  uint16(sym_structure_resource),
	8010:  uint16(290),
	8011:  uint16(1),
	8012:  uint16(sym_mixins),
	8013:  uint16(8),
	8014:  uint16(7),
	8015:  uint16(1),
	8016:  uint16(anon_sym_SLASH_SLASH_SLASH),
	8017:  uint16(612),
	8018:  uint16(1),
	8019:  uint16(anon_sym_DQUOTE),
	8020:  uint16(614),
	8021:  uint16(1),
	8022:  uint16(sym_string_fragment),
	8023:  uint16(617),
	8024:  uint16(1),
	8025:  uint16(aux_sym__escape_sequence_token1),
	8026:  uint16(620),
	8027:  uint16(1),
	8028:  uint16(sym_escape_sequence),
	8029:  uint16(235),
	8030:  uint16(1),
	8031:  uint16(sym__escape_sequence),
	8032:  uint16(5),
	8033:  uint16(2),
	8034:  uint16(anon_sym_COMMA),
	8035:  uint16(sym_comment),
	8036:  uint16(201),
	8037:  uint16(2),
	8038:  uint16(sym_documentation_comment),
	8039:  uint16(aux_sym__string_literal_repeat1),
	8040:  uint16(7),
	8041:  uint16(3),
	8042:  uint16(1),
	8043:  uint16(anon_sym_COMMA),
	8044:  uint16(5),
	8045:  uint16(1),
	8046:  uint16(sym_comment),
	8047:  uint16(7),
	8048:  uint16(1),
	8049:  uint16(anon_sym_SLASH_SLASH_SLASH),
	8050:  uint16(582),
	8051:  uint16(1),
	8052:  uint16(anon_sym_EQ),
	8053:  uint16(202),
	8054:  uint16(1),
	8055:  uint16(sym_documentation_comment),
	8056:  uint16(234),
	8057:  uint16(1),
	8058:  uint16(sym_value_assignment),
	8059:  uint16(623),
	8060:  uint16(4),
	8061:  uint16(anon_sym_DOLLAR),
	8062:  uint16(anon_sym_RBRACE),
	8063:  uint16(anon_sym_AT),
	8064:  uint16(aux_sym_identifier_token1),
	8065:  uint16(6),
	8066:  uint16(7),
	8067:  uint16(1),
	8068:  uint16(anon_sym_SLASH_SLASH_SLASH),
	8069:  uint16(606),
	8070:  uint16(1),
	8071:  uint16(aux_sym__multiline_string_fragment_token1),
	8072:  uint16(197),
	8073:  uint16(1),
	8074:  uint16(aux_sym__multiline_string_fragment_repeat1),
	8075:  uint16(203),
	8076:  uint16(1),
	8077:  uint16(sym_documentation_comment),
	8078:  uint16(5),
	8079:  uint16(2),
	8080:  uint16(anon_sym_COMMA),
	8081:  uint16(sym_comment),
	8082:  uint16(625),
	8083:  uint16(4),
	8084:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	8085:  uint16(aux_sym__multiline_string_fragment_token2),
	8086:  uint16(aux_sym__escape_sequence_token1),
	8087:  uint16(sym_escape_sequence),
	8088:  uint16(5),
	8089:  uint16(3),
	8090:  uint16(1),
	8091:  uint16(anon_sym_COMMA),
	8092:  uint16(5),
	8093:  uint16(1),
	8094:  uint16(sym_comment),
	8095:  uint16(7),
	8096:  uint16(1),
	8097:  uint16(anon_sym_SLASH_SLASH_SLASH),
	8098:  uint16(204),
	8099:  uint16(1),
	8100:  uint16(sym_documentation_comment),
	8101:  uint16(627),
	8102:  uint16(5),
	8103:  uint16(anon_sym_DOLLAR),
	8104:  uint16(anon_sym_EQ),
	8105:  uint16(anon_sym_RBRACE),
	8106:  uint16(anon_sym_AT),
	8107:  uint16(aux_sym_identifier_token1),
	8108:  uint16(7),
	8109:  uint16(3),
	8110:  uint16(1),
	8111:  uint16(anon_sym_COMMA),
	8112:  uint16(5),
	8113:  uint16(1),
	8114:  uint16(sym_comment),
	8115:  uint16(7),
	8116:  uint16(1),
	8117:  uint16(anon_sym_SLASH_SLASH_SLASH),
	8118:  uint16(631),
	8119:  uint16(1),
	8120:  uint16(anon_sym_metadata),
	8121:  uint16(270),
	8122:  uint16(1),
	8123:  uint16(sym_metadata_statement),
	8124:  uint16(629),
	8125:  uint16(2),
	8127:  uint16(anon_sym_namespace),
	8128:  uint16(205),
	8129:  uint16(2),
	8130:  uint16(sym_documentation_comment),
	8131:  uint16(aux_sym_metadata_section_repeat1),
	8132:  uint16(4),
	8133:  uint16(7),
	8134:  uint16(1),
	8135:  uint16(anon_sym_SLASH_SLASH_SLASH),
	8136:  uint16(206),
	8137:  uint16(1),
	8138:  uint16(sym_documentation_comment),
	8139:  uint16(5),
	8140:  uint16(2),
	8141:  uint16(anon_sym_COMMA),
	8142:  uint16(sym_comment),
	8143:  uint16(634),
	8144:  uint16(5),
	8145:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	8146:  uint16(aux_sym__multiline_string_fragment_token1),
	8147:  uint16(aux_sym__multiline_string_fragment_token2),
	8148:  uint16(aux_sym__escape_sequence_token1),
	8149:  uint16(sym_escape_sequence),
	8150:  uint16(6),
	8151:  uint16(3),
	8152:  uint16(1),
	8153:  uint16(anon_sym_COMMA),
	8154:  uint16(5),
	8155:  uint16(1),
	8156:  uint16(sym_comment),
	8157:  uint16(7),
	8158:  uint16(1),
	8159:  uint16(anon_sym_SLASH_SLASH_SLASH),
	8160:  uint16(638),
	8161:  uint16(1),
	8162:  uint16(anon_sym_DQUOTE),
	8163:  uint16(207),
	8164:  uint16(1),
	8165:  uint16(sym_documentation_comment),
	8166:  uint16(636),
	8167:  uint16(4),
	8168:  uint16(anon_sym_RBRACE),
	8169:  uint16(anon_sym_RPAREN),
	8170:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	8171:  uint16(aux_sym_identifier_token1),
	8172:  uint16(5),
	8173:  uint16(3),
	8174:  uint16(1),
	8175:  uint16(anon_sym_COMMA),
	8176:  uint16(5),
	8177:  uint16(1),
	8178:  uint16(sym_comment),
	8179:  uint16(7),
	8180:  uint16(1),
	8181:  uint16(anon_sym_SLASH_SLASH_SLASH),
	8182:  uint16(208),
	8183:  uint16(1),
	8184:  uint16(sym_documentation_comment),
	8185:  uint16(362),
	8186:  uint16(5),
	8188:  uint16(anon_sym_DOLLAR),
	8189:  uint16(anon_sym_metadata),
	8190:  uint16(anon_sym_namespace),
	8191:  uint16(anon_sym_RPAREN),
	8192:  uint16(5),
	8193:  uint16(3),
	8194:  uint16(1),
	8195:  uint16(anon_sym_COMMA),
	8196:  uint16(5),
	8197:  uint16(1),
	8198:  uint16(sym_comment),
	8199:  uint16(7),
	8200:  uint16(1),
	8201:  uint16(anon_sym_SLASH_SLASH_SLASH),
	8202:  uint16(209),
	8203:  uint16(1),
	8204:  uint16(sym_documentation_comment),
	8205:  uint16(354),
	8206:  uint16(5),
	8208:  uint16(anon_sym_DOLLAR),
	8209:  uint16(anon_sym_metadata),
	8210:  uint16(anon_sym_namespace),
	8211:  uint16(anon_sym_RPAREN),
	8212:  uint16(7),
	8213:  uint16(3),
	8214:  uint16(1),
	8215:  uint16(anon_sym_COMMA),
	8216:  uint16(5),
	8217:  uint16(1),
	8218:  uint16(sym_comment),
	8219:  uint16(7),
	8220:  uint16(1),
	8221:  uint16(anon_sym_SLASH_SLASH_SLASH),
	8222:  uint16(640),
	8223:  uint16(1),
	8224:  uint16(anon_sym_EQ),
	8225:  uint16(210),
	8226:  uint16(1),
	8227:  uint16(sym_documentation_comment),
	8228:  uint16(263),
	8229:  uint16(1),
	8230:  uint16(sym_value_assignment),
	8231:  uint16(642),
	8232:  uint16(3),
	8233:  uint16(anon_sym_RBRACE),
	8234:  uint16(anon_sym_AT),
	8235:  uint16(aux_sym_identifier_token1),
	8236:  uint16(4),
	8237:  uint16(7),
	8238:  uint16(1),
	8239:  uint16(anon_sym_SLASH_SLASH_SLASH),
	8240:  uint16(211),
	8241:  uint16(1),
	8242:  uint16(sym_documentation_comment),
	8243:  uint16(5),
	8244:  uint16(2),
	8245:  uint16(anon_sym_COMMA),
	8246:  uint16(sym_comment),
	8247:  uint16(644),
	8248:  uint16(5),
	8249:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	8250:  uint16(aux_sym__multiline_string_fragment_token1),
	8251:  uint16(aux_sym__multiline_string_fragment_token2),
	8252:  uint16(aux_sym__escape_sequence_token1),
	8253:  uint16(sym_escape_sequence),
	8254:  uint16(4),
	8255:  uint16(7),
	8256:  uint16(1),
	8257:  uint16(anon_sym_SLASH_SLASH_SLASH),
	8258:  uint16(212),
	8259:  uint16(1),
	8260:  uint16(sym_documentation_comment),
	8261:  uint16(5),
	8262:  uint16(2),
	8263:  uint16(anon_sym_COMMA),
	8264:  uint16(sym_comment),
	8265:  uint16(646),
	8266:  uint16(5),
	8267:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	8268:  uint16(aux_sym__multiline_string_fragment_token1),
	8269:  uint16(aux_sym__multiline_string_fragment_token2),
	8270:  uint16(aux_sym__escape_sequence_token1),
	8271:  uint16(sym_escape_sequence),
	8272:  uint16(7),
	8273:  uint16(3),
	8274:  uint16(1),
	8275:  uint16(anon_sym_COMMA),
	8276:  uint16(5),
	8277:  uint16(1),
	8278:  uint16(sym_comment),
	8279:  uint16(7),
	8280:  uint16(1),
	8281:  uint16(anon_sym_SLASH_SLASH_SLASH),
	8282:  uint16(640),
	8283:  uint16(1),
	8284:  uint16(anon_sym_EQ),
	8285:  uint16(213),
	8286:  uint16(1),
	8287:  uint16(sym_documentation_comment),
	8288:  uint16(268),
	8289:  uint16(1),
	8290:  uint16(sym_value_assignment),
	8291:  uint16(648),
	8292:  uint16(3),
	8293:  uint16(anon_sym_RBRACE),
	8294:  uint16(anon_sym_AT),
	8295:  uint16(aux_sym_identifier_token1),
	8296:  uint16(4),
	8297:  uint16(7),
	8298:  uint16(1),
	8299:  uint16(anon_sym_SLASH_SLASH_SLASH),
	8300:  uint16(214),
	8301:  uint16(1),
	8302:  uint16(sym_documentation_comment),
	8303:  uint16(5),
	8304:  uint16(2),
	8305:  uint16(anon_sym_COMMA),
	8306:  uint16(sym_comment),
	8307:  uint16(625),
	8308:  uint16(5),
	8309:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	8310:  uint16(aux_sym__multiline_string_fragment_token1),
	8311:  uint16(aux_sym__multiline_string_fragment_token2),
	8312:  uint16(aux_sym__escape_sequence_token1),
	8313:  uint16(sym_escape_sequence),
	8314:  uint16(7),
	8315:  uint16(3),
	8316:  uint16(1),
	8317:  uint16(anon_sym_COMMA),
	8318:  uint16(5),
	8319:  uint16(1),
	8320:  uint16(sym_comment),
	8321:  uint16(7),
	8322:  uint16(1),
	8323:  uint16(anon_sym_SLASH_SLASH_SLASH),
	8324:  uint16(650),
	8325:  uint16(1),
	8326:  uint16(anon_sym_LPAREN),
	8327:  uint16(215),
	8328:  uint16(1),
	8329:  uint16(sym_documentation_comment),
	8330:  uint16(273),
	8331:  uint16(1),
	8332:  uint16(sym_trait_body),
	8333:  uint16(113),
	8334:  uint16(3),
	8335:  uint16(anon_sym_DOLLAR),
	8336:  uint16(anon_sym_AT),
	8337:  uint16(aux_sym_identifier_token1),
	8338:  uint16(5),
	8339:  uint16(3),
	8340:  uint16(1),
	8341:  uint16(anon_sym_COMMA),
	8342:  uint16(5),
	8343:  uint16(1),
	8344:  uint16(sym_comment),
	8345:  uint16(7),
	8346:  uint16(1),
	8347:  uint16(anon_sym_SLASH_SLASH_SLASH),
	8348:  uint16(216),
	8349:  uint16(1),
	8350:  uint16(sym_documentation_comment),
	8351:  uint16(417),
	8352:  uint16(5),
	8354:  uint16(anon_sym_DOLLAR),
	8355:  uint16(anon_sym_metadata),
	8356:  uint16(anon_sym_namespace),
	8357:  uint16(anon_sym_RPAREN),
	8358:  uint16(4),
	8359:  uint16(7),
	8360:  uint16(1),
	8361:  uint16(anon_sym_SLASH_SLASH_SLASH),
	8362:  uint16(217),
	8363:  uint16(1),
	8364:  uint16(sym_documentation_comment),
	8365:  uint16(5),
	8366:  uint16(2),
	8367:  uint16(anon_sym_COMMA),
	8368:  uint16(sym_comment),
	8369:  uint16(644),
	8370:  uint16(5),
	8371:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	8372:  uint16(aux_sym__multiline_string_fragment_token1),
	8373:  uint16(aux_sym__multiline_string_fragment_token2),
	8374:  uint16(aux_sym__escape_sequence_token1),
	8375:  uint16(sym_escape_sequence),
	8376:  uint16(5),
	8377:  uint16(3),
	8378:  uint16(1),
	8379:  uint16(anon_sym_COMMA),
	8380:  uint16(5),
	8381:  uint16(1),
	8382:  uint16(sym_comment),
	8383:  uint16(7),
	8384:  uint16(1),
	8385:  uint16(anon_sym_SLASH_SLASH_SLASH),
	8386:  uint16(218),
	8387:  uint16(1),
	8388:  uint16(sym_documentation_comment),
	8389:  uint16(395),
	8390:  uint16(5),
	8392:  uint16(anon_sym_DOLLAR),
	8393:  uint16(anon_sym_metadata),
	8394:  uint16(anon_sym_namespace),
	8395:  uint16(anon_sym_RPAREN),
	8396:  uint16(5),
	8397:  uint16(3),
	8398:  uint16(1),
	8399:  uint16(anon_sym_COMMA),
	8400:  uint16(5),
	8401:  uint16(1),
	8402:  uint16(sym_comment),
	8403:  uint16(7),
	8404:  uint16(1),
	8405:  uint16(anon_sym_SLASH_SLASH_SLASH),
	8406:  uint16(219),
	8407:  uint16(1),
	8408:  uint16(sym_documentation_comment),
	8409:  uint16(391),
	8410:  uint16(5),
	8412:  uint16(anon_sym_DOLLAR),
	8413:  uint16(anon_sym_metadata),
	8414:  uint16(anon_sym_namespace),
	8415:  uint16(anon_sym_RPAREN),
	8416:  uint16(8),
	8417:  uint16(3),
	8418:  uint16(1),
	8419:  uint16(anon_sym_COMMA),
	8420:  uint16(5),
	8421:  uint16(1),
	8422:  uint16(sym_comment),
	8423:  uint16(7),
	8424:  uint16(1),
	8425:  uint16(anon_sym_SLASH_SLASH_SLASH),
	8426:  uint16(652),
	8427:  uint16(1),
	8428:  uint16(anon_sym_RBRACK),
	8429:  uint16(654),
	8430:  uint16(1),
	8431:  uint16(aux_sym_identifier_token1),
	8432:  uint16(317),
	8433:  uint16(1),
	8434:  uint16(sym_identifier),
	8435:  uint16(320),
	8436:  uint16(1),
	8437:  uint16(sym_operation_error),
	8438:  uint16(220),
	8439:  uint16(2),
	8440:  uint16(sym_documentation_comment),
	8441:  uint16(aux_sym_operation_errors_repeat1),
	8442:  uint16(9),
	8443:  uint16(3),
	8444:  uint16(1),
	8445:  uint16(anon_sym_COMMA),
	8446:  uint16(5),
	8447:  uint16(1),
	8448:  uint16(sym_comment),
	8449:  uint16(7),
	8450:  uint16(1),
	8451:  uint16(anon_sym_SLASH_SLASH_SLASH),
	8452:  uint16(535),
	8453:  uint16(1),
	8454:  uint16(aux_sym_identifier_token1),
	8455:  uint16(657),
	8456:  uint16(1),
	8457:  uint16(anon_sym_RBRACK),
	8458:  uint16(220),
	8459:  uint16(1),
	8460:  uint16(aux_sym_operation_errors_repeat1),
	8461:  uint16(221),
	8462:  uint16(1),
	8463:  uint16(sym_documentation_comment),
	8464:  uint16(317),
	8465:  uint16(1),
	8466:  uint16(sym_identifier),
	8467:  uint16(320),
	8468:  uint16(1),
	8469:  uint16(sym_operation_error),
	8470:  uint16(9),
	8471:  uint16(3),
	8472:  uint16(1),
	8473:  uint16(anon_sym_COMMA),
	8474:  uint16(5),
	8475:  uint16(1),
	8476:  uint16(sym_comment),
	8477:  uint16(7),
	8478:  uint16(1),
	8479:  uint16(anon_sym_SLASH_SLASH_SLASH),
	8480:  uint16(533),
	8481:  uint16(1),
	8482:  uint16(anon_sym_AT),
	8483:  uint16(535),
	8484:  uint16(1),
	8485:  uint16(aux_sym_identifier_token1),
	8486:  uint16(210),
	8487:  uint16(1),
	8488:  uint16(sym_identifier),
	8489:  uint16(222),
	8490:  uint16(1),
	8491:  uint16(sym_documentation_comment),
	8492:  uint16(255),
	8493:  uint16(1),
	8494:  uint16(aux_sym_shape_statement_repeat1),
	8495:  uint16(274),
	8496:  uint16(1),
	8497:  uint16(sym_trait_statement),
	8498:  uint16(7),
	8499:  uint16(3),
	8500:  uint16(1),
	8501:  uint16(anon_sym_COMMA),
	8502:  uint16(5),
	8503:  uint16(1),
	8504:  uint16(sym_comment),
	8505:  uint16(7),
	8506:  uint16(1),
	8507:  uint16(anon_sym_SLASH_SLASH_SLASH),
	8508:  uint16(659),
	8509:  uint16(1),
	8510:  uint16(anon_sym_AT),
	8511:  uint16(274),
	8512:  uint16(1),
	8513:  uint16(sym_trait_statement),
	8514:  uint16(162),
	8515:  uint16(2),
	8516:  uint16(anon_sym_DOLLAR),
	8517:  uint16(aux_sym_identifier_token1),
	8518:  uint16(223),
	8519:  uint16(2),
	8520:  uint16(sym_documentation_comment),
	8521:  uint16(aux_sym_shape_statement_repeat1),
	8522:  uint16(8),
	8523:  uint16(3),
	8524:  uint16(1),
	8525:  uint16(anon_sym_COMMA),
	8526:  uint16(5),
	8527:  uint16(1),
	8528:  uint16(sym_comment),
	8529:  uint16(7),
	8530:  uint16(1),
	8531:  uint16(anon_sym_SLASH_SLASH_SLASH),
	8532:  uint16(13),
	8533:  uint16(1),
	8534:  uint16(anon_sym_metadata),
	8535:  uint16(205),
	8536:  uint16(1),
	8537:  uint16(aux_sym_metadata_section_repeat1),
	8538:  uint16(224),
	8539:  uint16(1),
	8540:  uint16(sym_documentation_comment),
	8541:  uint16(270),
	8542:  uint16(1),
	8543:  uint16(sym_metadata_statement),
	8544:  uint16(662),
	8545:  uint16(2),
	8547:  uint16(anon_sym_namespace),
	8548:  uint16(9),
	8549:  uint16(3),
	8550:  uint16(1),
	8551:  uint16(anon_sym_COMMA),
	8552:  uint16(5),
	8553:  uint16(1),
	8554:  uint16(sym_comment),
	8555:  uint16(7),
	8556:  uint16(1),
	8557:  uint16(anon_sym_SLASH_SLASH_SLASH),
	8558:  uint16(535),
	8559:  uint16(1),
	8560:  uint16(aux_sym_identifier_token1),
	8561:  uint16(664),
	8562:  uint16(1),
	8563:  uint16(anon_sym_RBRACK),
	8564:  uint16(221),
	8565:  uint16(1),
	8566:  uint16(aux_sym_operation_errors_repeat1),
	8567:  uint16(225),
	8568:  uint16(1),
	8569:  uint16(sym_documentation_comment),
	8570:  uint16(317),
	8571:  uint16(1),
	8572:  uint16(sym_identifier),
	8573:  uint16(320),
	8574:  uint16(1),
	8575:  uint16(sym_operation_error),
	8576:  uint16(5),
	8577:  uint16(3),
	8578:  uint16(1),
	8579:  uint16(anon_sym_COMMA),
	8580:  uint16(5),
	8581:  uint16(1),
	8582:  uint16(sym_comment),
	8583:  uint16(7),
	8584:  uint16(1),
	8585:  uint16(anon_sym_SLASH_SLASH_SLASH),
	8586:  uint16(226),
	8587:  uint16(1),
	8588:  uint16(sym_documentation_comment),
	8589:  uint16(377),
	8590:  uint16(5),
	8592:  uint16(anon_sym_DOLLAR),
	8593:  uint16(anon_sym_metadata),
	8594:  uint16(anon_sym_namespace),
	8595:  uint16(anon_sym_RPAREN),
	8596:  uint16(5),
	8597:  uint16(3),
	8598:  uint16(1),
	8599:  uint16(anon_sym_COMMA),
	8600:  uint16(5),
	8601:  uint16(1),
	8602:  uint16(sym_comment),
	8603:  uint16(7),
	8604:  uint16(1),
	8605:  uint16(anon_sym_SLASH_SLASH_SLASH),
	8606:  uint16(227),
	8607:  uint16(1),
	8608:  uint16(sym_documentation_comment),
	8609:  uint16(373),
	8610:  uint16(5),
	8612:  uint16(anon_sym_DOLLAR),
	8613:  uint16(anon_sym_metadata),
	8614:  uint16(anon_sym_namespace),
	8615:  uint16(anon_sym_RPAREN),
	8616:  uint16(5),
	8617:  uint16(3),
	8618:  uint16(1),
	8619:  uint16(anon_sym_COMMA),
	8620:  uint16(5),
	8621:  uint16(1),
	8622:  uint16(sym_comment),
	8623:  uint16(7),
	8624:  uint16(1),
	8625:  uint16(anon_sym_SLASH_SLASH_SLASH),
	8626:  uint16(228),
	8627:  uint16(1),
	8628:  uint16(sym_documentation_comment),
	8629:  uint16(399),
	8630:  uint16(5),
	8632:  uint16(anon_sym_DOLLAR),
	8633:  uint16(anon_sym_metadata),
	8634:  uint16(anon_sym_namespace),
	8635:  uint16(anon_sym_RPAREN),
	8636:  uint16(5),
	8637:  uint16(3),
	8638:  uint16(1),
	8639:  uint16(anon_sym_COMMA),
	8640:  uint16(5),
	8641:  uint16(1),
	8642:  uint16(sym_comment),
	8643:  uint16(7),
	8644:  uint16(1),
	8645:  uint16(anon_sym_SLASH_SLASH_SLASH),
	8646:  uint16(229),
	8647:  uint16(1),
	8648:  uint16(sym_documentation_comment),
	8649:  uint16(366),
	8650:  uint16(5),
	8652:  uint16(anon_sym_DOLLAR),
	8653:  uint16(anon_sym_metadata),
	8654:  uint16(anon_sym_namespace),
	8655:  uint16(anon_sym_RPAREN),
	8656:  uint16(7),
	8657:  uint16(3),
	8658:  uint16(1),
	8659:  uint16(anon_sym_COMMA),
	8660:  uint16(5),
	8661:  uint16(1),
	8662:  uint16(sym_comment),
	8663:  uint16(7),
	8664:  uint16(1),
	8665:  uint16(anon_sym_SLASH_SLASH_SLASH),
	8666:  uint16(230),
	8667:  uint16(1),
	8668:  uint16(sym_documentation_comment),
	8669:  uint16(250),
	8670:  uint16(1),
	8671:  uint16(aux_sym_operation_member_repeat1),
	8672:  uint16(331),
	8673:  uint16(1),
	8674:  uint16(sym_identifier),
	8675:  uint16(666),
	8676:  uint16(2),
	8677:  uint16(anon_sym_RBRACE),
	8678:  uint16(aux_sym_identifier_token1),
	8679:  uint16(8),
	8680:  uint16(3),
	8681:  uint16(1),
	8682:  uint16(anon_sym_COMMA),
	8683:  uint16(5),
	8684:  uint16(1),
	8685:  uint16(sym_comment),
	8686:  uint16(7),
	8687:  uint16(1),
	8688:  uint16(anon_sym_SLASH_SLASH_SLASH),
	8689:  uint16(227),
	8690:  uint16(1),
	8691:  uint16(anon_sym_with),
	8692:  uint16(610),
	8693:  uint16(1),
	8694:  uint16(anon_sym_LBRACE),
	8695:  uint16(53),
	8696:  uint16(1),
	8697:  uint16(sym_shape_members),
	8698:  uint16(231),
	8699:  uint16(1),
	8700:  uint16(sym_documentation_comment),
	8701:  uint16(279),
	8702:  uint16(1),
	8703:  uint16(sym_mixins),
	8704:  uint16(5),
	8705:  uint16(3),
	8706:  uint16(1),
	8707:  uint16(anon_sym_COMMA),
	8708:  uint16(5),
	8709:  uint16(1),
	8710:  uint16(sym_comment),
	8711:  uint16(7),
	8712:  uint16(1),
	8713:  uint16(anon_sym_SLASH_SLASH_SLASH),
	8714:  uint16(232),
	8715:  uint16(1),
	8716:  uint16(sym_documentation_comment),
	8717:  uint16(668),
	8718:  uint16(4),
	8719:  uint16(anon_sym_DOLLAR),
	8720:  uint16(anon_sym_RBRACE),
	8721:  uint16(anon_sym_AT),
	8722:  uint16(aux_sym_identifier_token1),
	8723:  uint16(8),
	8724:  uint16(3),
	8725:  uint16(1),
	8726:  uint16(anon_sym_COMMA),
	8727:  uint16(5),
	8728:  uint16(1),
	8729:  uint16(sym_comment),
	8730:  uint16(7),
	8731:  uint16(1),
	8732:  uint16(anon_sym_SLASH_SLASH_SLASH),
	8733:  uint16(227),
	8734:  uint16(1),
	8735:  uint16(anon_sym_with),
	8736:  uint16(610),
	8737:  uint16(1),
	8738:  uint16(anon_sym_LBRACE),
	8739:  uint16(63),
	8740:  uint16(1),
	8741:  uint16(sym_shape_members),
	8742:  uint16(233),
	8743:  uint16(1),
	8744:  uint16(sym_documentation_comment),
	8745:  uint16(306),
	8746:  uint16(1),
	8747:  uint16(sym_mixins),
	8748:  uint16(5),
	8749:  uint16(3),
	8750:  uint16(1),
	8751:  uint16(anon_sym_COMMA),
	8752:  uint16(5),
	8753:  uint16(1),
	8754:  uint16(sym_comment),
	8755:  uint16(7),
	8756:  uint16(1),
	8757:  uint16(anon_sym_SLASH_SLASH_SLASH),
	8758:  uint16(234),
	8759:  uint16(1),
	8760:  uint16(sym_documentation_comment),
	8761:  uint16(670),
	8762:  uint16(4),
	8763:  uint16(anon_sym_DOLLAR),
	8764:  uint16(anon_sym_RBRACE),
	8765:  uint16(anon_sym_AT),
	8766:  uint16(aux_sym_identifier_token1),
	8767:  uint16(5),
	8768:  uint16(7),
	8769:  uint16(1),
	8770:  uint16(anon_sym_SLASH_SLASH_SLASH),
	8771:  uint16(674),
	8772:  uint16(1),
	8773:  uint16(sym_string_fragment),
	8774:  uint16(235),
	8775:  uint16(1),
	8776:  uint16(sym_documentation_comment),
	8777:  uint16(5),
	8778:  uint16(2),
	8779:  uint16(anon_sym_COMMA),
	8780:  uint16(sym_comment),
	8781:  uint16(672),
	8782:  uint16(3),
	8783:  uint16(anon_sym_DQUOTE),
	8784:  uint16(aux_sym__escape_sequence_token1),
	8785:  uint16(sym_escape_sequence),
	8786:  uint16(5),
	8787:  uint16(7),
	8788:  uint16(1),
	8789:  uint16(anon_sym_SLASH_SLASH_SLASH),
	8790:  uint16(676),
	8791:  uint16(1),
	8792:  uint16(sym_string_fragment),
	8793:  uint16(236),
	8794:  uint16(1),
	8795:  uint16(sym_documentation_comment),
	8796:  uint16(5),
	8797:  uint16(2),
	8798:  uint16(anon_sym_COMMA),
	8799:  uint16(sym_comment),
	8800:  uint16(644),
	8801:  uint16(3),
	8802:  uint16(anon_sym_DQUOTE),
	8803:  uint16(aux_sym__escape_sequence_token1),
	8804:  uint16(sym_escape_sequence),
	8805:  uint16(5),
	8806:  uint16(7),
	8807:  uint16(1),
	8808:  uint16(anon_sym_SLASH_SLASH_SLASH),
	8809:  uint16(676),
	8810:  uint16(1),
	8811:  uint16(sym_string_fragment),
	8812:  uint16(237),
	8813:  uint16(1),
	8814:  uint16(sym_documentation_comment),
	8815:  uint16(5),
	8816:  uint16(2),
	8817:  uint16(anon_sym_COMMA),
	8818:  uint16(sym_comment),
	8819:  uint16(644),
	8820:  uint16(3),
	8821:  uint16(anon_sym_DQUOTE),
	8822:  uint16(aux_sym__escape_sequence_token1),
	8823:  uint16(sym_escape_sequence),
	8824:  uint16(8),
	8825:  uint16(3),
	8826:  uint16(1),
	8827:  uint16(anon_sym_COMMA),
	8828:  uint16(5),
	8829:  uint16(1),
	8830:  uint16(sym_comment),
	8831:  uint16(7),
	8832:  uint16(1),
	8833:  uint16(anon_sym_SLASH_SLASH_SLASH),
	8834:  uint16(227),
	8835:  uint16(1),
	8836:  uint16(anon_sym_with),
	8837:  uint16(492),
	8838:  uint16(1),
	8839:  uint16(anon_sym_LBRACE),
	8840:  uint16(238),
	8841:  uint16(1),
	8842:  uint16(sym_documentation_comment),
	8843:  uint16(311),
	8844:  uint16(1),
	8845:  uint16(sym_mixins),
	8846:  uint16(314),
	8847:  uint16(1),
	8848:  uint16(sym_shape_members),
	8849:  uint16(5),
	8850:  uint16(3),
	8851:  uint16(1),
	8852:  uint16(anon_sym_COMMA),
	8853:  uint16(5),
	8854:  uint16(1),
	8855:  uint16(sym_comment),
	8856:  uint16(7),
	8857:  uint16(1),
	8858:  uint16(anon_sym_SLASH_SLASH_SLASH),
	8859:  uint16(239),
	8860:  uint16(1),
	8861:  uint16(sym_documentation_comment),
	8862:  uint16(678),
	8863:  uint16(4),
	8865:  uint16(anon_sym_DOLLAR),
	8866:  uint16(anon_sym_metadata),
	8867:  uint16(anon_sym_namespace),
	8868:  uint16(5),
	8869:  uint16(3),
	8870:  uint16(1),
	8871:  uint16(anon_sym_COMMA),
	8872:  uint16(5),
	8873:  uint16(1),
	8874:  uint16(sym_comment),
	8875:  uint16(7),
	8876:  uint16(1),
	8877:  uint16(anon_sym_SLASH_SLASH_SLASH),
	8878:  uint16(240),
	8879:  uint16(1),
	8880:  uint16(sym_documentation_comment),
	8881:  uint16(680),
	8882:  uint16(4),
	8883:  uint16(anon_sym_DOLLAR),
	8884:  uint16(anon_sym_RBRACE),
	8885:  uint16(anon_sym_AT),
	8886:  uint16(aux_sym_identifier_token1),
	8887:  uint16(8),
	8888:  uint16(3),
	8889:  uint16(1),
	8890:  uint16(anon_sym_COMMA),
	8891:  uint16(5),
	8892:  uint16(1),
	8893:  uint16(sym_comment),
	8894:  uint16(7),
	8895:  uint16(1),
	8896:  uint16(anon_sym_SLASH_SLASH_SLASH),
	8897:  uint16(39),
	8898:  uint16(1),
	8899:  uint16(anon_sym_AT),
	8900:  uint16(682),
	8901:  uint16(1),
	8902:  uint16(anon_sym_RBRACE),
	8903:  uint16(37),
	8904:  uint16(1),
	8905:  uint16(sym_trait_statement),
	8906:  uint16(241),
	8907:  uint16(1),
	8908:  uint16(sym_documentation_comment),
	8909:  uint16(258),
	8910:  uint16(1),
	8911:  uint16(aux_sym_shape_statement_repeat1),
	8912:  uint16(8),
	8913:  uint16(3),
	8914:  uint16(1),
	8915:  uint16(anon_sym_COMMA),
	8916:  uint16(5),
	8917:  uint16(1),
	8918:  uint16(sym_comment),
	8919:  uint16(7),
	8920:  uint16(1),
	8921:  uint16(anon_sym_SLASH_SLASH_SLASH),
	8922:  uint16(684),
	8923:  uint16(1),
	8924:  uint16(aux_sym_identifier_token1),
	8925:  uint16(23),
	8926:  uint16(1),
	8927:  uint16(sym__namespace_identifier),
	8928:  uint16(78),
	8929:  uint16(1),
	8930:  uint16(sym_absolute_root_shape_id),
	8931:  uint16(242),
	8932:  uint16(1),
	8933:  uint16(sym_documentation_comment),
	8934:  uint16(332),
	8935:  uint16(1),
	8936:  uint16(sym_namespace),
	8937:  uint16(5),
	8938:  uint16(3),
	8939:  uint16(1),
	8940:  uint16(anon_sym_COMMA),
	8941:  uint16(5),
	8942:  uint16(1),
	8943:  uint16(sym_comment),
	8944:  uint16(7),
	8945:  uint16(1),
	8946:  uint16(anon_sym_SLASH_SLASH_SLASH),
	8947:  uint16(243),
	8948:  uint16(1),
	8949:  uint16(sym_documentation_comment),
	8950:  uint16(686),
	8951:  uint16(4),
	8953:  uint16(anon_sym_DOLLAR),
	8954:  uint16(anon_sym_metadata),
	8955:  uint16(anon_sym_namespace),
	8956:  uint16(5),
	8957:  uint16(3),
	8958:  uint16(1),
	8959:  uint16(anon_sym_COMMA),
	8960:  uint16(5),
	8961:  uint16(1),
	8962:  uint16(sym_comment),
	8963:  uint16(7),
	8964:  uint16(1),
	8965:  uint16(anon_sym_SLASH_SLASH_SLASH),
	8966:  uint16(244),
	8967:  uint16(1),
	8968:  uint16(sym_documentation_comment),
	8969:  uint16(688),
	8970:  uint16(4),
	8971:  uint16(anon_sym_DOLLAR),
	8972:  uint16(anon_sym_RBRACE),
	8973:  uint16(anon_sym_AT),
	8974:  uint16(aux_sym_identifier_token1),
	8975:  uint16(8),
	8976:  uint16(3),
	8977:  uint16(1),
	8978:  uint16(anon_sym_COMMA),
	8979:  uint16(5),
	8980:  uint16(1),
	8981:  uint16(sym_comment),
	8982:  uint16(7),
	8983:  uint16(1),
	8984:  uint16(anon_sym_SLASH_SLASH_SLASH),
	8985:  uint16(15),
	8986:  uint16(1),
	8987:  uint16(anon_sym_namespace),
	8988:  uint16(503),
	8989:  uint16(1),
	8991:  uint16(2),
	8992:  uint16(1),
	8993:  uint16(sym_namespace_statement),
	8994:  uint16(245),
	8995:  uint16(1),
	8996:  uint16(sym_documentation_comment),
	8997:  uint16(327),
	8998:  uint16(1),
	8999:  uint16(sym_shape_section),
	9000:  uint16(6),
	9001:  uint16(3),
	9002:  uint16(1),
	9003:  uint16(anon_sym_COMMA),
	9004:  uint16(5),
	9005:  uint16(1),
	9006:  uint16(sym_comment),
	9007:  uint16(7),
	9008:  uint16(1),
	9009:  uint16(anon_sym_SLASH_SLASH_SLASH),
	9010:  uint16(692),
	9011:  uint16(1),
	9012:  uint16(anon_sym_DQUOTE),
	9013:  uint16(246),
	9014:  uint16(1),
	9015:  uint16(sym_documentation_comment),
	9016:  uint16(690),
	9017:  uint16(3),
	9018:  uint16(anon_sym_RBRACE),
	9019:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	9020:  uint16(aux_sym_identifier_token1),
	9021:  uint16(8),
	9022:  uint16(3),
	9023:  uint16(1),
	9024:  uint16(anon_sym_COMMA),
	9025:  uint16(5),
	9026:  uint16(1),
	9027:  uint16(sym_comment),
	9028:  uint16(7),
	9029:  uint16(1),
	9030:  uint16(anon_sym_SLASH_SLASH_SLASH),
	9031:  uint16(227),
	9032:  uint16(1),
	9033:  uint16(anon_sym_with),
	9034:  uint16(492),
	9035:  uint16(1),
	9036:  uint16(anon_sym_LBRACE),
	9037:  uint16(247),
	9038:  uint16(1),
	9039:  uint16(sym_documentation_comment),
	9040:  uint16(302),
	9041:  uint16(1),
	9042:  uint16(sym_mixins),
	9043:  uint16(305),
	9044:  uint16(1),
	9045:  uint16(sym_shape_members),
	9046:  uint16(5),
	9047:  uint16(3),
	9048:  uint16(1),
	9049:  uint16(anon_sym_COMMA),
	9050:  uint16(5),
	9051:  uint16(1),
	9052:  uint16(sym_comment),
	9053:  uint16(7),
	9054:  uint16(1),
	9055:  uint16(anon_sym_SLASH_SLASH_SLASH),
	9056:  uint16(248),
	9057:  uint16(1),
	9058:  uint16(sym_documentation_comment),
	9059:  uint16(694),
	9060:  uint16(4),
	9061:  uint16(anon_sym_DOLLAR),
	9062:  uint16(anon_sym_RBRACE),
	9063:  uint16(anon_sym_AT),
	9064:  uint16(aux_sym_identifier_token1),
	9065:  uint16(8),
	9066:  uint16(3),
	9067:  uint16(1),
	9068:  uint16(anon_sym_COMMA),
	9069:  uint16(5),
	9070:  uint16(1),
	9071:  uint16(sym_comment),
	9072:  uint16(7),
	9073:  uint16(1),
	9074:  uint16(anon_sym_SLASH_SLASH_SLASH),
	9075:  uint16(15),
	9076:  uint16(1),
	9077:  uint16(anon_sym_namespace),
	9078:  uint16(696),
	9079:  uint16(1),
	9081:  uint16(2),
	9082:  uint16(1),
	9083:  uint16(sym_namespace_statement),
	9084:  uint16(249),
	9085:  uint16(1),
	9086:  uint16(sym_documentation_comment),
	9087:  uint16(335),
	9088:  uint16(1),
	9089:  uint16(sym_shape_section),
	9090:  uint16(7),
	9091:  uint16(3),
	9092:  uint16(1),
	9093:  uint16(anon_sym_COMMA),
	9094:  uint16(5),
	9095:  uint16(1),
	9096:  uint16(sym_comment),
	9097:  uint16(7),
	9098:  uint16(1),
	9099:  uint16(anon_sym_SLASH_SLASH_SLASH),
	9100:  uint16(698),
	9101:  uint16(1),
	9102:  uint16(anon_sym_RBRACE),
	9103:  uint16(700),
	9104:  uint16(1),
	9105:  uint16(aux_sym_identifier_token1),
	9106:  uint16(331),
	9107:  uint16(1),
	9108:  uint16(sym_identifier),
	9109:  uint16(250),
	9110:  uint16(2),
	9111:  uint16(sym_documentation_comment),
	9112:  uint16(aux_sym_operation_member_repeat1),
	9113:  uint16(8),
	9114:  uint16(3),
	9115:  uint16(1),
	9116:  uint16(anon_sym_COMMA),
	9117:  uint16(5),
	9118:  uint16(1),
	9119:  uint16(sym_comment),
	9120:  uint16(7),
	9121:  uint16(1),
	9122:  uint16(anon_sym_SLASH_SLASH_SLASH),
	9123:  uint16(227),
	9124:  uint16(1),
	9125:  uint16(anon_sym_with),
	9126:  uint16(703),
	9127:  uint16(1),
	9128:  uint16(anon_sym_LBRACE),
	9129:  uint16(47),
	9130:  uint16(1),
	9131:  uint16(sym_operation_body),
	9132:  uint16(251),
	9133:  uint16(1),
	9134:  uint16(sym_documentation_comment),
	9135:  uint16(313),
	9136:  uint16(1),
	9137:  uint16(sym_mixins),
	9138:  uint16(5),
	9139:  uint16(3),
	9140:  uint16(1),
	9141:  uint16(anon_sym_COMMA),
	9142:  uint16(5),
	9143:  uint16(1),
	9144:  uint16(sym_comment),
	9145:  uint16(7),
	9146:  uint16(1),
	9147:  uint16(anon_sym_SLASH_SLASH_SLASH),
	9148:  uint16(252),
	9149:  uint16(1),
	9150:  uint16(sym_documentation_comment),
	9151:  uint16(623),
	9152:  uint16(4),
	9153:  uint16(anon_sym_DOLLAR),
	9154:  uint16(anon_sym_RBRACE),
	9155:  uint16(anon_sym_AT),
	9156:  uint16(aux_sym_identifier_token1),
	9157:  uint16(6),
	9158:  uint16(3),
	9159:  uint16(1),
	9160:  uint16(anon_sym_COMMA),
	9161:  uint16(5),
	9162:  uint16(1),
	9163:  uint16(sym_comment),
	9164:  uint16(7),
	9165:  uint16(1),
	9166:  uint16(anon_sym_SLASH_SLASH_SLASH),
	9167:  uint16(707),
	9168:  uint16(1),
	9169:  uint16(anon_sym_DQUOTE),
	9170:  uint16(253),
	9171:  uint16(1),
	9172:  uint16(sym_documentation_comment),
	9173:  uint16(705),
	9174:  uint16(3),
	9175:  uint16(anon_sym_RPAREN),
	9176:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	9177:  uint16(aux_sym_identifier_token1),
	9178:  uint16(8),
	9179:  uint16(3),
	9180:  uint16(1),
	9181:  uint16(anon_sym_COMMA),
	9182:  uint16(5),
	9183:  uint16(1),
	9184:  uint16(sym_comment),
	9185:  uint16(7),
	9186:  uint16(1),
	9187:  uint16(anon_sym_SLASH_SLASH_SLASH),
	9188:  uint16(117),
	9189:  uint16(1),
	9190:  uint16(anon_sym_LBRACE),
	9191:  uint16(227),
	9192:  uint16(1),
	9193:  uint16(anon_sym_with),
	9194:  uint16(48),
	9195:  uint16(1),
	9196:  uint16(sym_node_object),
	9197:  uint16(254),
	9198:  uint16(1),
	9199:  uint16(sym_documentation_comment),
	9200:  uint16(296),
	9201:  uint16(1),
	9202:  uint16(sym_mixins),
	9203:  uint16(7),
	9204:  uint16(3),
	9205:  uint16(1),
	9206:  uint16(anon_sym_COMMA),
	9207:  uint16(5),
	9208:  uint16(1),
	9209:  uint16(sym_comment),
	9210:  uint16(7),
	9211:  uint16(1),
	9212:  uint16(anon_sym_SLASH_SLASH_SLASH),
	9213:  uint16(162),
	9214:  uint16(1),
	9215:  uint16(aux_sym_identifier_token1),
	9216:  uint16(709),
	9217:  uint16(1),
	9218:  uint16(anon_sym_AT),
	9219:  uint16(274),
	9220:  uint16(1),
	9221:  uint16(sym_trait_statement),
	9222:  uint16(255),
	9223:  uint16(2),
	9224:  uint16(sym_documentation_comment),
	9225:  uint16(aux_sym_shape_statement_repeat1),
	9226:  uint16(8),
	9227:  uint16(3),
	9228:  uint16(1),
	9229:  uint16(anon_sym_COMMA),
	9230:  uint16(5),
	9231:  uint16(1),
	9232:  uint16(sym_comment),
	9233:  uint16(7),
	9234:  uint16(1),
	9235:  uint16(anon_sym_SLASH_SLASH_SLASH),
	9236:  uint16(227),
	9237:  uint16(1),
	9238:  uint16(anon_sym_with),
	9239:  uint16(712),
	9240:  uint16(1),
	9241:  uint16(anon_sym_LBRACE),
	9242:  uint16(54),
	9243:  uint16(1),
	9244:  uint16(sym_enum_members),
	9245:  uint16(256),
	9246:  uint16(1),
	9247:  uint16(sym_documentation_comment),
	9248:  uint16(316),
	9249:  uint16(1),
	9250:  uint16(sym_mixins),
	9251:  uint16(8),
	9252:  uint16(3),
	9253:  uint16(1),
	9254:  uint16(anon_sym_COMMA),
	9255:  uint16(5),
	9256:  uint16(1),
	9257:  uint16(sym_comment),
	9258:  uint16(7),
	9259:  uint16(1),
	9260:  uint16(anon_sym_SLASH_SLASH_SLASH),
	9261:  uint16(227),
	9262:  uint16(1),
	9263:  uint16(anon_sym_with),
	9264:  uint16(610),
	9265:  uint16(1),
	9266:  uint16(anon_sym_LBRACE),
	9267:  uint16(49),
	9268:  uint16(1),
	9269:  uint16(sym_shape_members),
	9270:  uint16(257),
	9271:  uint16(1),
	9272:  uint16(sym_documentation_comment),
	9273:  uint16(295),
	9274:  uint16(1),
	9275:  uint16(sym_mixins),
	9276:  uint16(8),
	9277:  uint16(3),
	9278:  uint16(1),
	9279:  uint16(anon_sym_COMMA),
	9280:  uint16(5),
	9281:  uint16(1),
	9282:  uint16(sym_comment),
	9283:  uint16(7),
	9284:  uint16(1),
	9285:  uint16(anon_sym_SLASH_SLASH_SLASH),
	9286:  uint16(39),
	9287:  uint16(1),
	9288:  uint16(anon_sym_AT),
	9289:  uint16(714),
	9290:  uint16(1),
	9291:  uint16(anon_sym_RBRACE),
	9292:  uint16(26),
	9293:  uint16(1),
	9294:  uint16(aux_sym_shape_statement_repeat1),
	9295:  uint16(37),
	9296:  uint16(1),
	9297:  uint16(sym_trait_statement),
	9298:  uint16(258),
	9299:  uint16(1),
	9300:  uint16(sym_documentation_comment),
	9301:  uint16(8),
	9302:  uint16(3),
	9303:  uint16(1),
	9304:  uint16(anon_sym_COMMA),
	9305:  uint16(5),
	9306:  uint16(1),
	9307:  uint16(sym_comment),
	9308:  uint16(7),
	9309:  uint16(1),
	9310:  uint16(anon_sym_SLASH_SLASH_SLASH),
	9311:  uint16(227),
	9312:  uint16(1),
	9313:  uint16(anon_sym_with),
	9314:  uint16(610),
	9315:  uint16(1),
	9316:  uint16(anon_sym_LBRACE),
	9317:  uint16(51),
	9318:  uint16(1),
	9319:  uint16(sym_shape_members),
	9320:  uint16(259),
	9321:  uint16(1),
	9322:  uint16(sym_documentation_comment),
	9323:  uint16(281),
	9324:  uint16(1),
	9325:  uint16(sym_mixins),
	9326:  uint16(8),
	9327:  uint16(3),
	9328:  uint16(1),
	9329:  uint16(anon_sym_COMMA),
	9330:  uint16(5),
	9331:  uint16(1),
	9332:  uint16(sym_comment),
	9333:  uint16(7),
	9334:  uint16(1),
	9335:  uint16(anon_sym_SLASH_SLASH_SLASH),
	9336:  uint16(227),
	9337:  uint16(1),
	9338:  uint16(anon_sym_with),
	9339:  uint16(610),
	9340:  uint16(1),
	9341:  uint16(anon_sym_LBRACE),
	9342:  uint16(52),
	9343:  uint16(1),
	9344:  uint16(sym_shape_members),
	9345:  uint16(260),
	9346:  uint16(1),
	9347:  uint16(sym_documentation_comment),
	9348:  uint16(280),
	9349:  uint16(1),
	9350:  uint16(sym_mixins),
	9351:  uint16(8),
	9352:  uint16(3),
	9353:  uint16(1),
	9354:  uint16(anon_sym_COMMA),
	9355:  uint16(5),
	9356:  uint16(1),
	9357:  uint16(sym_comment),
	9358:  uint16(7),
	9359:  uint16(1),
	9360:  uint16(anon_sym_SLASH_SLASH_SLASH),
	9361:  uint16(117),
	9362:  uint16(1),
	9363:  uint16(anon_sym_LBRACE),
	9364:  uint16(227),
	9365:  uint16(1),
	9366:  uint16(anon_sym_with),
	9367:  uint16(46),
	9368:  uint16(1),
	9369:  uint16(sym_node_object),
	9370:  uint16(261),
	9371:  uint16(1),
	9372:  uint16(sym_documentation_comment),
	9373:  uint16(324),
	9374:  uint16(1),
	9375:  uint16(sym_mixins),
	9376:  uint16(5),
	9377:  uint16(3),
	9378:  uint16(1),
	9379:  uint16(anon_sym_COMMA),
	9380:  uint16(5),
	9381:  uint16(1),
	9382:  uint16(sym_comment),
	9383:  uint16(7),
	9384:  uint16(1),
	9385:  uint16(anon_sym_SLASH_SLASH_SLASH),
	9386:  uint16(262),
	9387:  uint16(1),
	9388:  uint16(sym_documentation_comment),
	9389:  uint16(716),
	9390:  uint16(3),
	9391:  uint16(anon_sym_RBRACE),
	9392:  uint16(anon_sym_AT),
	9393:  uint16(aux_sym_identifier_token1),
	9394:  uint16(5),
	9395:  uint16(3),
	9396:  uint16(1),
	9397:  uint16(anon_sym_COMMA),
	9398:  uint16(5),
	9399:  uint16(1),
	9400:  uint16(sym_comment),
	9401:  uint16(7),
	9402:  uint16(1),
	9403:  uint16(anon_sym_SLASH_SLASH_SLASH),
	9404:  uint16(263),
	9405:  uint16(1),
	9406:  uint16(sym_documentation_comment),
	9407:  uint16(718),
	9408:  uint16(3),
	9409:  uint16(anon_sym_RBRACE),
	9410:  uint16(anon_sym_AT),
	9411:  uint16(aux_sym_identifier_token1),
	9412:  uint16(7),
	9413:  uint16(3),
	9414:  uint16(1),
	9415:  uint16(anon_sym_COMMA),
	9416:  uint16(5),
	9417:  uint16(1),
	9418:  uint16(sym_comment),
	9419:  uint16(7),
	9420:  uint16(1),
	9421:  uint16(anon_sym_SLASH_SLASH_SLASH),
	9422:  uint16(684),
	9423:  uint16(1),
	9424:  uint16(aux_sym_identifier_token1),
	9425:  uint16(23),
	9426:  uint16(1),
	9427:  uint16(sym__namespace_identifier),
	9428:  uint16(70),
	9429:  uint16(1),
	9430:  uint16(sym_namespace),
	9431:  uint16(264),
	9432:  uint16(1),
	9433:  uint16(sym_documentation_comment),
	9434:  uint16(5),
	9435:  uint16(3),
	9436:  uint16(1),
	9437:  uint16(anon_sym_COMMA),
	9438:  uint16(5),
	9439:  uint16(1),
	9440:  uint16(sym_comment),
	9441:  uint16(7),
	9442:  uint16(1),
	9443:  uint16(anon_sym_SLASH_SLASH_SLASH),
	9444:  uint16(265),
	9445:  uint16(1),
	9446:  uint16(sym_documentation_comment),
	9447:  uint16(720),
	9448:  uint16(3),
	9450:  uint16(anon_sym_metadata),
	9451:  uint16(anon_sym_namespace),
	9452:  uint16(7),
	9453:  uint16(3),
	9454:  uint16(1),
	9455:  uint16(anon_sym_COMMA),
	9456:  uint16(5),
	9457:  uint16(1),
	9458:  uint16(sym_comment),
	9459:  uint16(7),
	9460:  uint16(1),
	9461:  uint16(anon_sym_SLASH_SLASH_SLASH),
	9462:  uint16(722),
	9463:  uint16(1),
	9464:  uint16(anon_sym_COLON),
	9465:  uint16(724),
	9466:  uint16(1),
	9467:  uint16(anon_sym_COLON_EQ),
	9468:  uint16(266),
	9469:  uint16(1),
	9470:  uint16(sym_documentation_comment),
	9471:  uint16(278),
	9472:  uint16(1),
	9473:  uint16(sym_inline_structure),
	9474:  uint16(6),
	9475:  uint16(3),
	9476:  uint16(1),
	9477:  uint16(anon_sym_COMMA),
	9478:  uint16(5),
	9479:  uint16(1),
	9480:  uint16(sym_comment),
	9481:  uint16(7),
	9482:  uint16(1),
	9483:  uint16(anon_sym_SLASH_SLASH_SLASH),
	9484:  uint16(726),
	9485:  uint16(1),
	9486:  uint16(anon_sym_COLON),
	9487:  uint16(267),
	9488:  uint16(1),
	9489:  uint16(sym_documentation_comment),
	9490:  uint16(105),
	9491:  uint16(2),
	9492:  uint16(anon_sym_DOLLAR),
	9493:  uint16(anon_sym_RPAREN),
	9494:  uint16(5),
	9495:  uint16(3),
	9496:  uint16(1),
	9497:  uint16(anon_sym_COMMA),
	9498:  uint16(5),
	9499:  uint16(1),
	9500:  uint16(sym_comment),
	9501:  uint16(7),
	9502:  uint16(1),
	9503:  uint16(anon_sym_SLASH_SLASH_SLASH),
	9504:  uint16(268),
	9505:  uint16(1),
	9506:  uint16(sym_documentation_comment),
	9507:  uint16(728),
	9508:  uint16(3),
	9509:  uint16(anon_sym_RBRACE),
	9510:  uint16(anon_sym_AT),
	9511:  uint16(aux_sym_identifier_token1),
	9512:  uint16(5),
	9513:  uint16(3),
	9514:  uint16(1),
	9515:  uint16(anon_sym_COMMA),
	9516:  uint16(5),
	9517:  uint16(1),
	9518:  uint16(sym_comment),
	9519:  uint16(7),
	9520:  uint16(1),
	9521:  uint16(anon_sym_SLASH_SLASH_SLASH),
	9522:  uint16(269),
	9523:  uint16(1),
	9524:  uint16(sym_documentation_comment),
	9525:  uint16(145),
	9526:  uint16(3),
	9527:  uint16(anon_sym_DOLLAR),
	9528:  uint16(anon_sym_AT),
	9529:  uint16(aux_sym_identifier_token1),
	9530:  uint16(5),
	9531:  uint16(3),
	9532:  uint16(1),
	9533:  uint16(anon_sym_COMMA),
	9534:  uint16(5),
	9535:  uint16(1),
	9536:  uint16(sym_comment),
	9537:  uint16(7),
	9538:  uint16(1),
	9539:  uint16(anon_sym_SLASH_SLASH_SLASH),
	9540:  uint16(270),
	9541:  uint16(1),
	9542:  uint16(sym_documentation_comment),
	9543:  uint16(730),
	9544:  uint16(3),
	9546:  uint16(anon_sym_metadata),
	9547:  uint16(anon_sym_namespace),
	9548:  uint16(5),
	9549:  uint16(3),
	9550:  uint16(1),
	9551:  uint16(anon_sym_COMMA),
	9552:  uint16(5),
	9553:  uint16(1),
	9554:  uint16(sym_comment),
	9555:  uint16(7),
	9556:  uint16(1),
	9557:  uint16(anon_sym_SLASH_SLASH_SLASH),
	9558:  uint16(271),
	9559:  uint16(1),
	9560:  uint16(sym_documentation_comment),
	9561:  uint16(149),
	9562:  uint16(3),
	9563:  uint16(anon_sym_DOLLAR),
	9564:  uint16(anon_sym_AT),
	9565:  uint16(aux_sym_identifier_token1),
	9566:  uint16(7),
	9567:  uint16(3),
	9568:  uint16(1),
	9569:  uint16(anon_sym_COMMA),
	9570:  uint16(5),
	9571:  uint16(1),
	9572:  uint16(sym_comment),
	9573:  uint16(7),
	9574:  uint16(1),
	9575:  uint16(anon_sym_SLASH_SLASH_SLASH),
	9576:  uint16(39),
	9577:  uint16(1),
	9578:  uint16(anon_sym_AT),
	9579:  uint16(732),
	9580:  uint16(1),
	9581:  uint16(anon_sym_LBRACE),
	9582:  uint16(43),
	9583:  uint16(1),
	9584:  uint16(sym_trait_statement),
	9585:  uint16(272),
	9586:  uint16(1),
	9587:  uint16(sym_documentation_comment),
	9588:  uint16(5),
	9589:  uint16(3),
	9590:  uint16(1),
	9591:  uint16(anon_sym_COMMA),
	9592:  uint16(5),
	9593:  uint16(1),
	9594:  uint16(sym_comment),
	9595:  uint16(7),
	9596:  uint16(1),
	9597:  uint16(anon_sym_SLASH_SLASH_SLASH),
	9598:  uint16(273),
	9599:  uint16(1),
	9600:  uint16(sym_documentation_comment),
	9601:  uint16(147),
	9602:  uint16(3),
	9603:  uint16(anon_sym_DOLLAR),
	9604:  uint16(anon_sym_AT),
	9605:  uint16(aux_sym_identifier_token1),
	9606:  uint16(5),
	9607:  uint16(3),
	9608:  uint16(1),
	9609:  uint16(anon_sym_COMMA),
	9610:  uint16(5),
	9611:  uint16(1),
	9612:  uint16(sym_comment),
	9613:  uint16(7),
	9614:  uint16(1),
	9615:  uint16(anon_sym_SLASH_SLASH_SLASH),
	9616:  uint16(274),
	9617:  uint16(1),
	9618:  uint16(sym_documentation_comment),
	9619:  uint16(235),
	9620:  uint16(3),
	9621:  uint16(anon_sym_DOLLAR),
	9622:  uint16(anon_sym_AT),
	9623:  uint16(aux_sym_identifier_token1),
	9624:  uint16(6),
	9625:  uint16(3),
	9626:  uint16(1),
	9627:  uint16(anon_sym_COMMA),
	9628:  uint16(5),
	9629:  uint16(1),
	9630:  uint16(sym_comment),
	9631:  uint16(7),
	9632:  uint16(1),
	9633:  uint16(anon_sym_SLASH_SLASH_SLASH),
	9634:  uint16(734),
	9635:  uint16(1),
	9636:  uint16(aux_sym_number_token1),
	9637:  uint16(736),
	9638:  uint16(1),
	9639:  uint16(aux_sym_float_token1),
	9640:  uint16(275),
	9641:  uint16(1),
	9642:  uint16(sym_documentation_comment),
	9643:  uint16(6),
	9644:  uint16(3),
	9645:  uint16(1),
	9646:  uint16(anon_sym_COMMA),
	9647:  uint16(5),
	9648:  uint16(1),
	9649:  uint16(sym_comment),
	9650:  uint16(7),
	9651:  uint16(1),
	9652:  uint16(anon_sym_SLASH_SLASH_SLASH),
	9653:  uint16(492),
	9654:  uint16(1),
	9655:  uint16(anon_sym_LBRACE),
	9656:  uint16(276),
	9657:  uint16(1),
	9658:  uint16(sym_documentation_comment),
	9659:  uint16(314),
	9660:  uint16(1),
	9661:  uint16(sym_shape_members),
	9662:  uint16(6),
	9663:  uint16(3),
	9664:  uint16(1),
	9665:  uint16(anon_sym_COMMA),
	9666:  uint16(5),
	9667:  uint16(1),
	9668:  uint16(sym_comment),
	9669:  uint16(7),
	9670:  uint16(1),
	9671:  uint16(anon_sym_SLASH_SLASH_SLASH),
	9672:  uint16(425),
	9673:  uint16(1),
	9674:  uint16(aux_sym_identifier_token1),
	9675:  uint16(33),
	9676:  uint16(1),
	9677:  uint16(sym_identifier),
	9678:  uint16(277),
	9679:  uint16(1),
	9680:  uint16(sym_documentation_comment),
	9681:  uint16(5),
	9682:  uint16(3),
	9683:  uint16(1),
	9684:  uint16(anon_sym_COMMA),
	9685:  uint16(5),
	9686:  uint16(1),
	9687:  uint16(sym_comment),
	9688:  uint16(7),
	9689:  uint16(1),
	9690:  uint16(anon_sym_SLASH_SLASH_SLASH),
	9691:  uint16(278),
	9692:  uint16(1),
	9693:  uint16(sym_documentation_comment),
	9694:  uint16(738),
	9695:  uint16(2),
	9696:  uint16(anon_sym_RBRACE),
	9697:  uint16(aux_sym_identifier_token1),
	9698:  uint16(6),
	9699:  uint16(3),
	9700:  uint16(1),
	9701:  uint16(anon_sym_COMMA),
	9702:  uint16(5),
	9703:  uint16(1),
	9704:  uint16(sym_comment),
	9705:  uint16(7),
	9706:  uint16(1),
	9707:  uint16(anon_sym_SLASH_SLASH_SLASH),
	9708:  uint16(610),
	9709:  uint16(1),
	9710:  uint16(anon_sym_LBRACE),
	9711:  uint16(59),
	9712:  uint16(1),
	9713:  uint16(sym_shape_members),
	9714:  uint16(279),
	9715:  uint16(1),
	9716:  uint16(sym_documentation_comment),
	9717:  uint16(6),
	9718:  uint16(3),
	9719:  uint16(1),
	9720:  uint16(anon_sym_COMMA),
	9721:  uint16(5),
	9722:  uint16(1),
	9723:  uint16(sym_comment),
	9724:  uint16(7),
	9725:  uint16(1),
	9726:  uint16(anon_sym_SLASH_SLASH_SLASH),
	9727:  uint16(610),
	9728:  uint16(1),
	9729:  uint16(anon_sym_LBRACE),
	9730:  uint16(61),
	9731:  uint16(1),
	9732:  uint16(sym_shape_members),
	9733:  uint16(280),
	9734:  uint16(1),
	9735:  uint16(sym_documentation_comment),
	9736:  uint16(6),
	9737:  uint16(3),
	9738:  uint16(1),
	9739:  uint16(anon_sym_COMMA),
	9740:  uint16(5),
	9741:  uint16(1),
	9742:  uint16(sym_comment),
	9743:  uint16(7),
	9744:  uint16(1),
	9745:  uint16(anon_sym_SLASH_SLASH_SLASH),
	9746:  uint16(610),
	9747:  uint16(1),
	9748:  uint16(anon_sym_LBRACE),
	9749:  uint16(62),
	9750:  uint16(1),
	9751:  uint16(sym_shape_members),
	9752:  uint16(281),
	9753:  uint16(1),
	9754:  uint16(sym_documentation_comment),
	9755:  uint16(5),
	9756:  uint16(3),
	9757:  uint16(1),
	9758:  uint16(anon_sym_COMMA),
	9759:  uint16(5),
	9760:  uint16(1),
	9761:  uint16(sym_comment),
	9762:  uint16(7),
	9763:  uint16(1),
	9764:  uint16(anon_sym_SLASH_SLASH_SLASH),
	9765:  uint16(282),
	9766:  uint16(1),
	9767:  uint16(sym_documentation_comment),
	9768:  uint16(740),
	9769:  uint16(2),
	9770:  uint16(anon_sym_RBRACK),
	9771:  uint16(aux_sym_identifier_token1),
	9772:  uint16(6),
	9773:  uint16(3),
	9774:  uint16(1),
	9775:  uint16(anon_sym_COMMA),
	9776:  uint16(5),
	9777:  uint16(1),
	9778:  uint16(sym_comment),
	9779:  uint16(7),
	9780:  uint16(1),
	9781:  uint16(anon_sym_SLASH_SLASH_SLASH),
	9782:  uint16(742),
	9783:  uint16(1),
	9784:  uint16(aux_sym_number_token1),
	9785:  uint16(744),
	9786:  uint16(1),
	9787:  uint16(aux_sym_float_token1),
	9788:  uint16(283),
	9789:  uint16(1),
	9790:  uint16(sym_documentation_comment),
	9791:  uint16(6),
	9792:  uint16(3),
	9793:  uint16(1),
	9794:  uint16(anon_sym_COMMA),
	9795:  uint16(5),
	9796:  uint16(1),
	9797:  uint16(sym_comment),
	9798:  uint16(7),
	9799:  uint16(1),
	9800:  uint16(anon_sym_SLASH_SLASH_SLASH),
	9801:  uint16(746),
	9802:  uint16(1),
	9803:  uint16(aux_sym_identifier_token1),
	9804:  uint16(89),
	9805:  uint16(1),
	9806:  uint16(sym_identifier),
	9807:  uint16(284),
	9808:  uint16(1),
	9809:  uint16(sym_documentation_comment),
	9810:  uint16(6),
	9811:  uint16(3),
	9812:  uint16(1),
	9813:  uint16(anon_sym_COMMA),
	9814:  uint16(5),
	9815:  uint16(1),
	9816:  uint16(sym_comment),
	9817:  uint16(7),
	9818:  uint16(1),
	9819:  uint16(anon_sym_SLASH_SLASH_SLASH),
	9820:  uint16(746),
	9821:  uint16(1),
	9822:  uint16(aux_sym_identifier_token1),
	9823:  uint16(88),
	9824:  uint16(1),
	9825:  uint16(sym_identifier),
	9826:  uint16(285),
	9827:  uint16(1),
	9828:  uint16(sym_documentation_comment),
	9829:  uint16(6),
	9830:  uint16(3),
	9831:  uint16(1),
	9832:  uint16(anon_sym_COMMA),
	9833:  uint16(5),
	9834:  uint16(1),
	9835:  uint16(sym_comment),
	9836:  uint16(7),
	9837:  uint16(1),
	9838:  uint16(anon_sym_SLASH_SLASH_SLASH),
	9839:  uint16(425),
	9840:  uint16(1),
	9841:  uint16(aux_sym_identifier_token1),
	9842:  uint16(261),
	9843:  uint16(1),
	9844:  uint16(sym_identifier),
	9845:  uint16(286),
	9846:  uint16(1),
	9847:  uint16(sym_documentation_comment),
	9848:  uint16(6),
	9849:  uint16(3),
	9850:  uint16(1),
	9851:  uint16(anon_sym_COMMA),
	9852:  uint16(5),
	9853:  uint16(1),
	9854:  uint16(sym_comment),
	9855:  uint16(7),
	9856:  uint16(1),
	9857:  uint16(anon_sym_SLASH_SLASH_SLASH),
	9858:  uint16(425),
	9859:  uint16(1),
	9860:  uint16(aux_sym_identifier_token1),
	9861:  uint16(13),
	9862:  uint16(1),
	9863:  uint16(sym_identifier),
	9864:  uint16(287),
	9865:  uint16(1),
	9866:  uint16(sym_documentation_comment),
	9867:  uint16(6),
	9868:  uint16(3),
	9869:  uint16(1),
	9870:  uint16(anon_sym_COMMA),
	9871:  uint16(5),
	9872:  uint16(1),
	9873:  uint16(sym_comment),
	9874:  uint16(7),
	9875:  uint16(1),
	9876:  uint16(anon_sym_SLASH_SLASH_SLASH),
	9877:  uint16(425),
	9878:  uint16(1),
	9879:  uint16(aux_sym_identifier_token1),
	9880:  uint16(14),
	9881:  uint16(1),
	9882:  uint16(sym_identifier),
	9883:  uint16(288),
	9884:  uint16(1),
	9885:  uint16(sym_documentation_comment),
	9886:  uint16(6),
	9887:  uint16(3),
	9888:  uint16(1),
	9889:  uint16(anon_sym_COMMA),
	9890:  uint16(5),
	9891:  uint16(1),
	9892:  uint16(sym_comment),
	9893:  uint16(7),
	9894:  uint16(1),
	9895:  uint16(anon_sym_SLASH_SLASH_SLASH),
	9896:  uint16(354),
	9897:  uint16(1),
	9898:  uint16(anon_sym_RPAREN),
	9899:  uint16(748),
	9900:  uint16(1),
	9901:  uint16(anon_sym_COLON),
	9902:  uint16(289),
	9903:  uint16(1),
	9904:  uint16(sym_documentation_comment),
	9905:  uint16(6),
	9906:  uint16(3),
	9907:  uint16(1),
	9908:  uint16(anon_sym_COMMA),
	9909:  uint16(5),
	9910:  uint16(1),
	9911:  uint16(sym_comment),
	9912:  uint16(7),
	9913:  uint16(1),
	9914:  uint16(anon_sym_SLASH_SLASH_SLASH),
	9915:  uint16(610),
	9916:  uint16(1),
	9917:  uint16(anon_sym_LBRACE),
	9918:  uint16(63),
	9919:  uint16(1),
	9920:  uint16(sym_shape_members),
	9921:  uint16(290),
	9922:  uint16(1),
	9923:  uint16(sym_documentation_comment),
	9924:  uint16(5),
	9925:  uint16(3),
	9926:  uint16(1),
	9927:  uint16(anon_sym_COMMA),
	9928:  uint16(5),
	9929:  uint16(1),
	9930:  uint16(sym_comment),
	9931:  uint16(7),
	9932:  uint16(1),
	9933:  uint16(anon_sym_SLASH_SLASH_SLASH),
	9934:  uint16(291),
	9935:  uint16(1),
	9936:  uint16(sym_documentation_comment),
	9937:  uint16(750),
	9938:  uint16(2),
	9939:  uint16(anon_sym_RBRACE),
	9940:  uint16(aux_sym_identifier_token1),
	9941:  uint16(5),
	9942:  uint16(3),
	9943:  uint16(1),
	9944:  uint16(anon_sym_COMMA),
	9945:  uint16(5),
	9946:  uint16(1),
	9947:  uint16(sym_comment),
	9948:  uint16(7),
	9949:  uint16(1),
	9950:  uint16(anon_sym_SLASH_SLASH_SLASH),
	9951:  uint16(292),
	9952:  uint16(1),
	9953:  uint16(sym_documentation_comment),
	9954:  uint16(267),
	9955:  uint16(2),
	9956:  uint16(anon_sym_RBRACE),
	9957:  uint16(aux_sym_identifier_token1),
	9958:  uint16(5),
	9959:  uint16(3),
	9960:  uint16(1),
	9961:  uint16(anon_sym_COMMA),
	9962:  uint16(5),
	9963:  uint16(1),
	9964:  uint16(sym_comment),
	9965:  uint16(7),
	9966:  uint16(1),
	9967:  uint16(anon_sym_SLASH_SLASH_SLASH),
	9968:  uint16(293),
	9969:  uint16(1),
	9970:  uint16(sym_documentation_comment),
	9971:  uint16(321),
	9972:  uint16(2),
	9973:  uint16(anon_sym_RBRACE),
	9974:  uint16(aux_sym_identifier_token1),
	9975:  uint16(6),
	9976:  uint16(3),
	9977:  uint16(1),
	9978:  uint16(anon_sym_COMMA),
	9979:  uint16(5),
	9980:  uint16(1),
	9981:  uint16(sym_comment),
	9982:  uint16(7),
	9983:  uint16(1),
	9984:  uint16(anon_sym_SLASH_SLASH_SLASH),
	9985:  uint16(425),
	9986:  uint16(1),
	9987:  uint16(aux_sym_identifier_token1),
	9988:  uint16(251),
	9989:  uint16(1),
	9990:  uint16(sym_identifier),
	9991:  uint16(294),
	9992:  uint16(1),
	9993:  uint16(sym_documentation_comment),
	9994:  uint16(6),
	9995:  uint16(3),
	9996:  uint16(1),
	9997:  uint16(anon_sym_COMMA),
	9998:  uint16(5),
	9999:  uint16(1),
	10000: uint16(sym_comment),
	10001: uint16(7),
	10002: uint16(1),
	10003: uint16(anon_sym_SLASH_SLASH_SLASH),
	10004: uint16(610),
	10005: uint16(1),
	10006: uint16(anon_sym_LBRACE),
	10007: uint16(64),
	10008: uint16(1),
	10009: uint16(sym_shape_members),
	10010: uint16(295),
	10011: uint16(1),
	10012: uint16(sym_documentation_comment),
	10013: uint16(6),
	10014: uint16(3),
	10015: uint16(1),
	10016: uint16(anon_sym_COMMA),
	10017: uint16(5),
	10018: uint16(1),
	10019: uint16(sym_comment),
	10020: uint16(7),
	10021: uint16(1),
	10022: uint16(anon_sym_SLASH_SLASH_SLASH),
	10023: uint16(117),
	10024: uint16(1),
	10025: uint16(anon_sym_LBRACE),
	10026: uint16(66),
	10027: uint16(1),
	10028: uint16(sym_node_object),
	10029: uint16(296),
	10030: uint16(1),
	10031: uint16(sym_documentation_comment),
	10032: uint16(5),
	10033: uint16(3),
	10034: uint16(1),
	10035: uint16(anon_sym_COMMA),
	10036: uint16(5),
	10037: uint16(1),
	10038: uint16(sym_comment),
	10039: uint16(7),
	10040: uint16(1),
	10041: uint16(anon_sym_SLASH_SLASH_SLASH),
	10042: uint16(297),
	10043: uint16(1),
	10044: uint16(sym_documentation_comment),
	10045: uint16(752),
	10046: uint16(2),
	10047: uint16(anon_sym_RBRACE),
	10048: uint16(aux_sym_identifier_token1),
	10049: uint16(5),
	10050: uint16(3),
	10051: uint16(1),
	10052: uint16(anon_sym_COMMA),
	10053: uint16(5),
	10054: uint16(1),
	10055: uint16(sym_comment),
	10056: uint16(7),
	10057: uint16(1),
	10058: uint16(anon_sym_SLASH_SLASH_SLASH),
	10059: uint16(298),
	10060: uint16(1),
	10061: uint16(sym_documentation_comment),
	10062: uint16(754),
	10063: uint16(2),
	10064: uint16(anon_sym_RBRACE),
	10065: uint16(aux_sym_identifier_token1),
	10066: uint16(4),
	10067: uint16(5),
	10068: uint16(1),
	10069: uint16(sym_comment),
	10070: uint16(7),
	10071: uint16(1),
	10072: uint16(anon_sym_SLASH_SLASH_SLASH),
	10073: uint16(299),
	10074: uint16(1),
	10075: uint16(sym_documentation_comment),
	10076: uint16(756),
	10077: uint16(3),
	10078: uint16(anon_sym_RBRACE),
	10079: uint16(anon_sym_COMMA),
	10080: uint16(aux_sym_identifier_token1),
	10081: uint16(6),
	10082: uint16(3),
	10083: uint16(1),
	10084: uint16(anon_sym_COMMA),
	10085: uint16(5),
	10086: uint16(1),
	10087: uint16(sym_comment),
	10088: uint16(7),
	10089: uint16(1),
	10090: uint16(anon_sym_SLASH_SLASH_SLASH),
	10091: uint16(758),
	10092: uint16(1),
	10093: uint16(aux_sym_number_token1),
	10094: uint16(760),
	10095: uint16(1),
	10096: uint16(aux_sym_float_token1),
	10097: uint16(300),
	10098: uint16(1),
	10099: uint16(sym_documentation_comment),
	10100: uint16(6),
	10101: uint16(3),
	10102: uint16(1),
	10103: uint16(anon_sym_COMMA),
	10104: uint16(5),
	10105: uint16(1),
	10106: uint16(sym_comment),
	10107: uint16(7),
	10108: uint16(1),
	10109: uint16(anon_sym_SLASH_SLASH_SLASH),
	10110: uint16(425),
	10111: uint16(1),
	10112: uint16(aux_sym_identifier_token1),
	10113: uint16(254),
	10114: uint16(1),
	10115: uint16(sym_identifier),
	10116: uint16(301),
	10117: uint16(1),
	10118: uint16(sym_documentation_comment),
	10119: uint16(6),
	10120: uint16(3),
	10121: uint16(1),
	10122: uint16(anon_sym_COMMA),
	10123: uint16(5),
	10124: uint16(1),
	10125: uint16(sym_comment),
	10126: uint16(7),
	10127: uint16(1),
	10128: uint16(anon_sym_SLASH_SLASH_SLASH),
	10129: uint16(492),
	10130: uint16(1),
	10131: uint16(anon_sym_LBRACE),
	10132: uint16(297),
	10133: uint16(1),
	10134: uint16(sym_shape_members),
	10135: uint16(302),
	10136: uint16(1),
	10137: uint16(sym_documentation_comment),
	10138: uint16(5),
	10139: uint16(5),
	10140: uint16(1),
	10141: uint16(sym_comment),
	10142: uint16(7),
	10143: uint16(1),
	10144: uint16(anon_sym_SLASH_SLASH_SLASH),
	10145: uint16(762),
	10146: uint16(1),
	10147: uint16(anon_sym_COMMA),
	10148: uint16(303),
	10149: uint16(1),
	10150: uint16(sym_documentation_comment),
	10151: uint16(666),
	10152: uint16(2),
	10153: uint16(anon_sym_RBRACE),
	10154: uint16(aux_sym_identifier_token1),
	10155: uint16(5),
	10156: uint16(3),
	10157: uint16(1),
	10158: uint16(anon_sym_COMMA),
	10159: uint16(5),
	10160: uint16(1),
	10161: uint16(sym_comment),
	10162: uint16(7),
	10163: uint16(1),
	10164: uint16(anon_sym_SLASH_SLASH_SLASH),
	10165: uint16(304),
	10166: uint16(1),
	10167: uint16(sym_documentation_comment),
	10168: uint16(764),
	10169: uint16(2),
	10170: uint16(anon_sym_RBRACE),
	10171: uint16(aux_sym_identifier_token1),
	10172: uint16(5),
	10173: uint16(3),
	10174: uint16(1),
	10175: uint16(anon_sym_COMMA),
	10176: uint16(5),
	10177: uint16(1),
	10178: uint16(sym_comment),
	10179: uint16(7),
	10180: uint16(1),
	10181: uint16(anon_sym_SLASH_SLASH_SLASH),
	10182: uint16(305),
	10183: uint16(1),
	10184: uint16(sym_documentation_comment),
	10185: uint16(766),
	10186: uint16(2),
	10187: uint16(anon_sym_RBRACE),
	10188: uint16(aux_sym_identifier_token1),
	10189: uint16(6),
	10190: uint16(3),
	10191: uint16(1),
	10192: uint16(anon_sym_COMMA),
	10193: uint16(5),
	10194: uint16(1),
	10195: uint16(sym_comment),
	10196: uint16(7),
	10197: uint16(1),
	10198: uint16(anon_sym_SLASH_SLASH_SLASH),
	10199: uint16(610),
	10200: uint16(1),
	10201: uint16(anon_sym_LBRACE),
	10202: uint16(74),
	10203: uint16(1),
	10204: uint16(sym_shape_members),
	10205: uint16(306),
	10206: uint16(1),
	10207: uint16(sym_documentation_comment),
	10208: uint16(5),
	10209: uint16(3),
	10210: uint16(1),
	10211: uint16(anon_sym_COMMA),
	10212: uint16(5),
	10213: uint16(1),
	10214: uint16(sym_comment),
	10215: uint16(7),
	10216: uint16(1),
	10217: uint16(anon_sym_SLASH_SLASH_SLASH),
	10218: uint16(307),
	10219: uint16(1),
	10220: uint16(sym_documentation_comment),
	10221: uint16(768),
	10222: uint16(2),
	10223: uint16(anon_sym_LBRACE),
	10224: uint16(anon_sym_with),
	10225: uint16(6),
	10226: uint16(3),
	10227: uint16(1),
	10228: uint16(anon_sym_COMMA),
	10229: uint16(5),
	10230: uint16(1),
	10231: uint16(sym_comment),
	10232: uint16(7),
	10233: uint16(1),
	10234: uint16(anon_sym_SLASH_SLASH_SLASH),
	10235: uint16(535),
	10236: uint16(1),
	10237: uint16(aux_sym_identifier_token1),
	10238: uint16(127),
	10239: uint16(1),
	10240: uint16(sym_identifier),
	10241: uint16(308),
	10242: uint16(1),
	10243: uint16(sym_documentation_comment),
	10244: uint16(6),
	10245: uint16(3),
	10246: uint16(1),
	10247: uint16(anon_sym_COMMA),
	10248: uint16(5),
	10249: uint16(1),
	10250: uint16(sym_comment),
	10251: uint16(7),
	10252: uint16(1),
	10253: uint16(anon_sym_SLASH_SLASH_SLASH),
	10254: uint16(535),
	10255: uint16(1),
	10256: uint16(aux_sym_identifier_token1),
	10257: uint16(123),
	10258: uint16(1),
	10259: uint16(sym_identifier),
	10260: uint16(309),
	10261: uint16(1),
	10262: uint16(sym_documentation_comment),
	10263: uint16(6),
	10264: uint16(3),
	10265: uint16(1),
	10266: uint16(anon_sym_COMMA),
	10267: uint16(5),
	10268: uint16(1),
	10269: uint16(sym_comment),
	10270: uint16(7),
	10271: uint16(1),
	10272: uint16(anon_sym_SLASH_SLASH_SLASH),
	10273: uint16(684),
	10274: uint16(1),
	10275: uint16(aux_sym_identifier_token1),
	10276: uint16(27),
	10277: uint16(1),
	10278: uint16(sym__namespace_identifier),
	10279: uint16(310),
	10280: uint16(1),
	10281: uint16(sym_documentation_comment),
	10282: uint16(6),
	10283: uint16(3),
	10284: uint16(1),
	10285: uint16(anon_sym_COMMA),
	10286: uint16(5),
	10287: uint16(1),
	10288: uint16(sym_comment),
	10289: uint16(7),
	10290: uint16(1),
	10291: uint16(anon_sym_SLASH_SLASH_SLASH),
	10292: uint16(492),
	10293: uint16(1),
	10294: uint16(anon_sym_LBRACE),
	10295: uint16(305),
	10296: uint16(1),
	10297: uint16(sym_shape_members),
	10298: uint16(311),
	10299: uint16(1),
	10300: uint16(sym_documentation_comment),
	10301: uint16(6),
	10302: uint16(3),
	10303: uint16(1),
	10304: uint16(anon_sym_COMMA),
	10305: uint16(5),
	10306: uint16(1),
	10307: uint16(sym_comment),
	10308: uint16(7),
	10309: uint16(1),
	10310: uint16(anon_sym_SLASH_SLASH_SLASH),
	10311: uint16(425),
	10312: uint16(1),
	10313: uint16(aux_sym_identifier_token1),
	10314: uint16(257),
	10315: uint16(1),
	10316: uint16(sym_identifier),
	10317: uint16(312),
	10318: uint16(1),
	10319: uint16(sym_documentation_comment),
	10320: uint16(6),
	10321: uint16(3),
	10322: uint16(1),
	10323: uint16(anon_sym_COMMA),
	10324: uint16(5),
	10325: uint16(1),
	10326: uint16(sym_comment),
	10327: uint16(7),
	10328: uint16(1),
	10329: uint16(anon_sym_SLASH_SLASH_SLASH),
	10330: uint16(703),
	10331: uint16(1),
	10332: uint16(anon_sym_LBRACE),
	10333: uint16(68),
	10334: uint16(1),
	10335: uint16(sym_operation_body),
	10336: uint16(313),
	10337: uint16(1),
	10338: uint16(sym_documentation_comment),
	10339: uint16(5),
	10340: uint16(3),
	10341: uint16(1),
	10342: uint16(anon_sym_COMMA),
	10343: uint16(5),
	10344: uint16(1),
	10345: uint16(sym_comment),
	10346: uint16(7),
	10347: uint16(1),
	10348: uint16(anon_sym_SLASH_SLASH_SLASH),
	10349: uint16(314),
	10350: uint16(1),
	10351: uint16(sym_documentation_comment),
	10352: uint16(770),
	10353: uint16(2),
	10354: uint16(anon_sym_RBRACE),
	10355: uint16(aux_sym_identifier_token1),
	10356: uint16(6),
	10357: uint16(3),
	10358: uint16(1),
	10359: uint16(anon_sym_COMMA),
	10360: uint16(5),
	10361: uint16(1),
	10362: uint16(sym_comment),
	10363: uint16(7),
	10364: uint16(1),
	10365: uint16(anon_sym_SLASH_SLASH_SLASH),
	10366: uint16(535),
	10367: uint16(1),
	10368: uint16(aux_sym_identifier_token1),
	10369: uint16(204),
	10370: uint16(1),
	10371: uint16(sym_identifier),
	10372: uint16(315),
	10373: uint16(1),
	10374: uint16(sym_documentation_comment),
	10375: uint16(6),
	10376: uint16(3),
	10377: uint16(1),
	10378: uint16(anon_sym_COMMA),
	10379: uint16(5),
	10380: uint16(1),
	10381: uint16(sym_comment),
	10382: uint16(7),
	10383: uint16(1),
	10384: uint16(anon_sym_SLASH_SLASH_SLASH),
	10385: uint16(712),
	10386: uint16(1),
	10387: uint16(anon_sym_LBRACE),
	10388: uint16(60),
	10389: uint16(1),
	10390: uint16(sym_enum_members),
	10391: uint16(316),
	10392: uint16(1),
	10393: uint16(sym_documentation_comment),
	10394: uint16(5),
	10395: uint16(3),
	10396: uint16(1),
	10397: uint16(anon_sym_COMMA),
	10398: uint16(5),
	10399: uint16(1),
	10400: uint16(sym_comment),
	10401: uint16(7),
	10402: uint16(1),
	10403: uint16(anon_sym_SLASH_SLASH_SLASH),
	10404: uint16(317),
	10405: uint16(1),
	10406: uint16(sym_documentation_comment),
	10407: uint16(772),
	10408: uint16(2),
	10409: uint16(anon_sym_RBRACK),
	10410: uint16(aux_sym_identifier_token1),
	10411: uint16(6),
	10412: uint16(3),
	10413: uint16(1),
	10414: uint16(anon_sym_COMMA),
	10415: uint16(5),
	10416: uint16(1),
	10417: uint16(sym_comment),
	10418: uint16(7),
	10419: uint16(1),
	10420: uint16(anon_sym_SLASH_SLASH_SLASH),
	10421: uint16(425),
	10422: uint16(1),
	10423: uint16(aux_sym_identifier_token1),
	10424: uint16(256),
	10425: uint16(1),
	10426: uint16(sym_identifier),
	10427: uint16(318),
	10428: uint16(1),
	10429: uint16(sym_documentation_comment),
	10430: uint16(6),
	10431: uint16(3),
	10432: uint16(1),
	10433: uint16(anon_sym_COMMA),
	10434: uint16(5),
	10435: uint16(1),
	10436: uint16(sym_comment),
	10437: uint16(7),
	10438: uint16(1),
	10439: uint16(anon_sym_SLASH_SLASH_SLASH),
	10440: uint16(425),
	10441: uint16(1),
	10442: uint16(aux_sym_identifier_token1),
	10443: uint16(231),
	10444: uint16(1),
	10445: uint16(sym_identifier),
	10446: uint16(319),
	10447: uint16(1),
	10448: uint16(sym_documentation_comment),
	10449: uint16(5),
	10450: uint16(3),
	10451: uint16(1),
	10452: uint16(anon_sym_COMMA),
	10453: uint16(5),
	10454: uint16(1),
	10455: uint16(sym_comment),
	10456: uint16(7),
	10457: uint16(1),
	10458: uint16(anon_sym_SLASH_SLASH_SLASH),
	10459: uint16(320),
	10460: uint16(1),
	10461: uint16(sym_documentation_comment),
	10462: uint16(774),
	10463: uint16(2),
	10464: uint16(anon_sym_RBRACK),
	10465: uint16(aux_sym_identifier_token1),
	10466: uint16(6),
	10467: uint16(3),
	10468: uint16(1),
	10469: uint16(anon_sym_COMMA),
	10470: uint16(5),
	10471: uint16(1),
	10472: uint16(sym_comment),
	10473: uint16(7),
	10474: uint16(1),
	10475: uint16(anon_sym_SLASH_SLASH_SLASH),
	10476: uint16(425),
	10477: uint16(1),
	10478: uint16(aux_sym_identifier_token1),
	10479: uint16(260),
	10480: uint16(1),
	10481: uint16(sym_identifier),
	10482: uint16(321),
	10483: uint16(1),
	10484: uint16(sym_documentation_comment),
	10485: uint16(6),
	10486: uint16(3),
	10487: uint16(1),
	10488: uint16(anon_sym_COMMA),
	10489: uint16(5),
	10490: uint16(1),
	10491: uint16(sym_comment),
	10492: uint16(7),
	10493: uint16(1),
	10494: uint16(anon_sym_SLASH_SLASH_SLASH),
	10495: uint16(425),
	10496: uint16(1),
	10497: uint16(aux_sym_identifier_token1),
	10498: uint16(259),
	10499: uint16(1),
	10500: uint16(sym_identifier),
	10501: uint16(322),
	10502: uint16(1),
	10503: uint16(sym_documentation_comment),
	10504: uint16(6),
	10505: uint16(3),
	10506: uint16(1),
	10507: uint16(anon_sym_COMMA),
	10508: uint16(5),
	10509: uint16(1),
	10510: uint16(sym_comment),
	10511: uint16(7),
	10512: uint16(1),
	10513: uint16(anon_sym_SLASH_SLASH_SLASH),
	10514: uint16(425),
	10515: uint16(1),
	10516: uint16(aux_sym_identifier_token1),
	10517: uint16(200),
	10518: uint16(1),
	10519: uint16(sym_identifier),
	10520: uint16(323),
	10521: uint16(1),
	10522: uint16(sym_documentation_comment),
	10523: uint16(6),
	10524: uint16(3),
	10525: uint16(1),
	10526: uint16(anon_sym_COMMA),
	10527: uint16(5),
	10528: uint16(1),
	10529: uint16(sym_comment),
	10530: uint16(7),
	10531: uint16(1),
	10532: uint16(anon_sym_SLASH_SLASH_SLASH),
	10533: uint16(117),
	10534: uint16(1),
	10535: uint16(anon_sym_LBRACE),
	10536: uint16(69),
	10537: uint16(1),
	10538: uint16(sym_node_object),
	10539: uint16(324),
	10540: uint16(1),
	10541: uint16(sym_documentation_comment),
	10542: uint16(4),
	10543: uint16(5),
	10544: uint16(1),
	10545: uint16(sym_comment),
	10546: uint16(7),
	10547: uint16(1),
	10548: uint16(anon_sym_SLASH_SLASH_SLASH),
	10549: uint16(325),
	10550: uint16(1),
	10551: uint16(sym_documentation_comment),
	10552: uint16(776),
	10553: uint16(3),
	10554: uint16(anon_sym_RBRACE),
	10555: uint16(anon_sym_COMMA),
	10556: uint16(aux_sym_identifier_token1),
	10557: uint16(5),
	10558: uint16(3),
	10559: uint16(1),
	10560: uint16(anon_sym_COMMA),
	10561: uint16(5),
	10562: uint16(1),
	10563: uint16(sym_comment),
	10564: uint16(7),
	10565: uint16(1),
	10566: uint16(anon_sym_SLASH_SLASH_SLASH),
	10567: uint16(778),
	10568: uint16(1),
	10569: uint16(anon_sym_RPAREN),
	10570: uint16(326),
	10571: uint16(1),
	10572: uint16(sym_documentation_comment),
	10573: uint16(5),
	10574: uint16(3),
	10575: uint16(1),
	10576: uint16(anon_sym_COMMA),
	10577: uint16(5),
	10578: uint16(1),
	10579: uint16(sym_comment),
	10580: uint16(7),
	10581: uint16(1),
	10582: uint16(anon_sym_SLASH_SLASH_SLASH),
	10583: uint16(696),
	10584: uint16(1),
	10586: uint16(327),
	10587: uint16(1),
	10588: uint16(sym_documentation_comment),
	10589: uint16(5),
	10590: uint16(3),
	10591: uint16(1),
	10592: uint16(anon_sym_COMMA),
	10593: uint16(5),
	10594: uint16(1),
	10595: uint16(sym_comment),
	10596: uint16(7),
	10597: uint16(1),
	10598: uint16(anon_sym_SLASH_SLASH_SLASH),
	10599: uint16(780),
	10600: uint16(1),
	10601: uint16(anon_sym_POUND),
	10602: uint16(328),
	10603: uint16(1),
	10604: uint16(sym_documentation_comment),
	10605: uint16(5),
	10606: uint16(3),
	10607: uint16(1),
	10608: uint16(anon_sym_COMMA),
	10609: uint16(5),
	10610: uint16(1),
	10611: uint16(sym_comment),
	10612: uint16(7),
	10613: uint16(1),
	10614: uint16(anon_sym_SLASH_SLASH_SLASH),
	10615: uint16(782),
	10616: uint16(1),
	10617: uint16(anon_sym_COLON),
	10618: uint16(329),
	10619: uint16(1),
	10620: uint16(sym_documentation_comment),
	10621: uint16(5),
	10622: uint16(3),
	10623: uint16(1),
	10624: uint16(anon_sym_COMMA),
	10625: uint16(5),
	10626: uint16(1),
	10627: uint16(sym_comment),
	10628: uint16(7),
	10629: uint16(1),
	10630: uint16(anon_sym_SLASH_SLASH_SLASH),
	10631: uint16(784),
	10632: uint16(1),
	10633: uint16(anon_sym_COLON),
	10634: uint16(330),
	10635: uint16(1),
	10636: uint16(sym_documentation_comment),
	10637: uint16(5),
	10638: uint16(3),
	10639: uint16(1),
	10640: uint16(anon_sym_COMMA),
	10641: uint16(5),
	10642: uint16(1),
	10643: uint16(sym_comment),
	10644: uint16(7),
	10645: uint16(1),
	10646: uint16(anon_sym_SLASH_SLASH_SLASH),
	10647: uint16(786),
	10648: uint16(1),
	10649: uint16(anon_sym_COLON),
	10650: uint16(331),
	10651: uint16(1),
	10652: uint16(sym_documentation_comment),
	10653: uint16(5),
	10654: uint16(3),
	10655: uint16(1),
	10656: uint16(anon_sym_COMMA),
	10657: uint16(5),
	10658: uint16(1),
	10659: uint16(sym_comment),
	10660: uint16(7),
	10661: uint16(1),
	10662: uint16(anon_sym_SLASH_SLASH_SLASH),
	10663: uint16(788),
	10664: uint16(1),
	10665: uint16(anon_sym_POUND),
	10666: uint16(332),
	10667: uint16(1),
	10668: uint16(sym_documentation_comment),
	10669: uint16(5),
	10670: uint16(3),
	10671: uint16(1),
	10672: uint16(anon_sym_COMMA),
	10673: uint16(5),
	10674: uint16(1),
	10675: uint16(sym_comment),
	10676: uint16(7),
	10677: uint16(1),
	10678: uint16(anon_sym_SLASH_SLASH_SLASH),
	10679: uint16(790),
	10680: uint16(1),
	10681: uint16(aux_sym_identifier_token1),
	10682: uint16(333),
	10683: uint16(1),
	10684: uint16(sym_documentation_comment),
	10685: uint16(5),
	10686: uint16(3),
	10687: uint16(1),
	10688: uint16(anon_sym_COMMA),
	10689: uint16(5),
	10690: uint16(1),
	10691: uint16(sym_comment),
	10692: uint16(7),
	10693: uint16(1),
	10694: uint16(anon_sym_SLASH_SLASH_SLASH),
	10695: uint16(792),
	10696: uint16(1),
	10698: uint16(334),
	10699: uint16(1),
	10700: uint16(sym_documentation_comment),
	10701: uint16(5),
	10702: uint16(3),
	10703: uint16(1),
	10704: uint16(anon_sym_COMMA),
	10705: uint16(5),
	10706: uint16(1),
	10707: uint16(sym_comment),
	10708: uint16(7),
	10709: uint16(1),
	10710: uint16(anon_sym_SLASH_SLASH_SLASH),
	10711: uint16(794),
	10712: uint16(1),
	10714: uint16(335),
	10715: uint16(1),
	10716: uint16(sym_documentation_comment),
	10717: uint16(5),
	10718: uint16(3),
	10719: uint16(1),
	10720: uint16(anon_sym_COMMA),
	10721: uint16(5),
	10722: uint16(1),
	10723: uint16(sym_comment),
	10724: uint16(7),
	10725: uint16(1),
	10726: uint16(anon_sym_SLASH_SLASH_SLASH),
	10727: uint16(796),
	10728: uint16(1),
	10729: uint16(anon_sym_RPAREN),
	10730: uint16(336),
	10731: uint16(1),
	10732: uint16(sym_documentation_comment),
	10733: uint16(4),
	10734: uint16(7),
	10735: uint16(1),
	10736: uint16(anon_sym_SLASH_SLASH_SLASH),
	10737: uint16(798),
	10738: uint16(1),
	10739: uint16(aux_sym_documentation_comment_token1),
	10740: uint16(337),
	10741: uint16(1),
	10742: uint16(sym_documentation_comment),
	10743: uint16(5),
	10744: uint16(2),
	10745: uint16(anon_sym_COMMA),
	10746: uint16(sym_comment),
	10747: uint16(5),
	10748: uint16(3),
	10749: uint16(1),
	10750: uint16(anon_sym_COMMA),
	10751: uint16(5),
	10752: uint16(1),
	10753: uint16(sym_comment),
	10754: uint16(7),
	10755: uint16(1),
	10756: uint16(anon_sym_SLASH_SLASH_SLASH),
	10757: uint16(800),
	10758: uint16(1),
	10759: uint16(anon_sym_RPAREN),
	10760: uint16(338),
	10761: uint16(1),
	10762: uint16(sym_documentation_comment),
	10763: uint16(5),
	10764: uint16(3),
	10765: uint16(1),
	10766: uint16(anon_sym_COMMA),
	10767: uint16(5),
	10768: uint16(1),
	10769: uint16(sym_comment),
	10770: uint16(7),
	10771: uint16(1),
	10772: uint16(anon_sym_SLASH_SLASH_SLASH),
	10773: uint16(802),
	10774: uint16(1),
	10775: uint16(anon_sym_RPAREN),
	10776: uint16(339),
	10777: uint16(1),
	10778: uint16(sym_documentation_comment),
	10779: uint16(5),
	10780: uint16(3),
	10781: uint16(1),
	10782: uint16(anon_sym_COMMA),
	10783: uint16(5),
	10784: uint16(1),
	10785: uint16(sym_comment),
	10786: uint16(7),
	10787: uint16(1),
	10788: uint16(anon_sym_SLASH_SLASH_SLASH),
	10789: uint16(804),
	10790: uint16(1),
	10791: uint16(anon_sym_COLON),
	10792: uint16(340),
	10793: uint16(1),
	10794: uint16(sym_documentation_comment),
	10795: uint16(5),
	10796: uint16(3),
	10797: uint16(1),
	10798: uint16(anon_sym_COMMA),
	10799: uint16(5),
	10800: uint16(1),
	10801: uint16(sym_comment),
	10802: uint16(7),
	10803: uint16(1),
	10804: uint16(anon_sym_SLASH_SLASH_SLASH),
	10805: uint16(806),
	10806: uint16(1),
	10807: uint16(anon_sym_LBRACK),
	10808: uint16(341),
	10809: uint16(1),
	10810: uint16(sym_documentation_comment),
	10811: uint16(5),
	10812: uint16(3),
	10813: uint16(1),
	10814: uint16(anon_sym_COMMA),
	10815: uint16(5),
	10816: uint16(1),
	10817: uint16(sym_comment),
	10818: uint16(7),
	10819: uint16(1),
	10820: uint16(anon_sym_SLASH_SLASH_SLASH),
	10821: uint16(748),
	10822: uint16(1),
	10823: uint16(anon_sym_COLON),
	10824: uint16(342),
	10825: uint16(1),
	10826: uint16(sym_documentation_comment),
	10827: uint16(5),
	10828: uint16(3),
	10829: uint16(1),
	10830: uint16(anon_sym_COMMA),
	10831: uint16(5),
	10832: uint16(1),
	10833: uint16(sym_comment),
	10834: uint16(7),
	10835: uint16(1),
	10836: uint16(anon_sym_SLASH_SLASH_SLASH),
	10837: uint16(726),
	10838: uint16(1),
	10839: uint16(anon_sym_COLON),
	10840: uint16(343),
	10841: uint16(1),
	10842: uint16(sym_documentation_comment),
	10843: uint16(5),
	10844: uint16(3),
	10845: uint16(1),
	10846: uint16(anon_sym_COMMA),
	10847: uint16(5),
	10848: uint16(1),
	10849: uint16(sym_comment),
	10850: uint16(7),
	10851: uint16(1),
	10852: uint16(anon_sym_SLASH_SLASH_SLASH),
	10853: uint16(808),
	10854: uint16(1),
	10855: uint16(anon_sym_EQ),
	10856: uint16(344),
	10857: uint16(1),
	10858: uint16(sym_documentation_comment),
	10859: uint16(5),
	10860: uint16(3),
	10861: uint16(1),
	10862: uint16(anon_sym_COMMA),
	10863: uint16(5),
	10864: uint16(1),
	10865: uint16(sym_comment),
	10866: uint16(7),
	10867: uint16(1),
	10868: uint16(anon_sym_SLASH_SLASH_SLASH),
	10869: uint16(810),
	10870: uint16(1),
	10871: uint16(anon_sym_POUND),
	10872: uint16(345),
	10873: uint16(1),
	10874: uint16(sym_documentation_comment),
	10875: uint16(5),
	10876: uint16(3),
	10877: uint16(1),
	10878: uint16(anon_sym_COMMA),
	10879: uint16(5),
	10880: uint16(1),
	10881: uint16(sym_comment),
	10882: uint16(7),
	10883: uint16(1),
	10884: uint16(anon_sym_SLASH_SLASH_SLASH),
	10885: uint16(812),
	10886: uint16(1),
	10887: uint16(anon_sym_COLON),
	10888: uint16(346),
	10889: uint16(1),
	10890: uint16(sym_documentation_comment),
	10891: uint16(5),
	10892: uint16(3),
	10893: uint16(1),
	10894: uint16(anon_sym_COMMA),
	10895: uint16(5),
	10896: uint16(1),
	10897: uint16(sym_comment),
	10898: uint16(7),
	10899: uint16(1),
	10900: uint16(anon_sym_SLASH_SLASH_SLASH),
	10901: uint16(503),
	10902: uint16(1),
	10904: uint16(347),
	10905: uint16(1),
	10906: uint16(sym_documentation_comment),
	10907: uint16(5),
	10908: uint16(3),
	10909: uint16(1),
	10910: uint16(anon_sym_COMMA),
	10911: uint16(5),
	10912: uint16(1),
	10913: uint16(sym_comment),
	10914: uint16(7),
	10915: uint16(1),
	10916: uint16(anon_sym_SLASH_SLASH_SLASH),
	10917: uint16(814),
	10918: uint16(1),
	10919: uint16(anon_sym_COLON),
	10920: uint16(348),
	10921: uint16(1),
	10922: uint16(sym_documentation_comment),
	10923: uint16(1),
	10924: uint16(816),
	10925: uint16(1),
}

var ts_small_parse_table_map = [348]uint32_t{
	1:   uint32(107),
	2:   uint32(214),
	3:   uint32(319),
	4:   uint32(375),
	5:   uint32(431),
	6:   uint32(488),
	7:   uint32(547),
	8:   uint32(606),
	9:   uint32(689),
	10:  uint32(740),
	11:  uint32(791),
	12:  uint32(842),
	13:  uint32(893),
	14:  uint32(945),
	15:  uint32(1043),
	16:  uint32(1141),
	17:  uint32(1187),
	18:  uint32(1233),
	19:  uint32(1279),
	20:  uint32(1325),
	21:  uint32(1371),
	22:  uint32(1420),
	23:  uint32(1467),
	24:  uint32(1516),
	25:  uint32(1565),
	26:  uint32(1609),
	27:  uint32(1691),
	28:  uint32(1773),
	29:  uint32(1817),
	30:  uint32(1897),
	31:  uint32(1979),
	32:  uint32(2027),
	33:  uint32(2109),
	34:  uint32(2191),
	35:  uint32(2273),
	36:  uint32(2316),
	37:  uint32(2359),
	38:  uint32(2435),
	39:  uint32(2477),
	40:  uint32(2519),
	41:  uint32(2561),
	42:  uint32(2603),
	43:  uint32(2679),
	44:  uint32(2721),
	45:  uint32(2763),
	46:  uint32(2805),
	47:  uint32(2847),
	48:  uint32(2889),
	49:  uint32(2931),
	50:  uint32(2973),
	51:  uint32(3015),
	52:  uint32(3057),
	53:  uint32(3099),
	54:  uint32(3141),
	55:  uint32(3183),
	56:  uint32(3225),
	57:  uint32(3267),
	58:  uint32(3309),
	59:  uint32(3351),
	60:  uint32(3393),
	61:  uint32(3435),
	62:  uint32(3477),
	63:  uint32(3519),
	64:  uint32(3595),
	65:  uint32(3637),
	66:  uint32(3679),
	67:  uint32(3721),
	68:  uint32(3763),
	69:  uint32(3805),
	70:  uint32(3847),
	71:  uint32(3889),
	72:  uint32(3931),
	73:  uint32(3973),
	74:  uint32(4015),
	75:  uint32(4057),
	76:  uint32(4133),
	77:  uint32(4175),
	78:  uint32(4251),
	79:  uint32(4293),
	80:  uint32(4331),
	81:  uint32(4367),
	82:  uint32(4403),
	83:  uint32(4441),
	84:  uint32(4471),
	85:  uint32(4501),
	86:  uint32(4531),
	87:  uint32(4561),
	88:  uint32(4591),
	89:  uint32(4620),
	90:  uint32(4649),
	91:  uint32(4678),
	92:  uint32(4707),
	93:  uint32(4740),
	94:  uint32(4769),
	95:  uint32(4798),
	96:  uint32(4827),
	97:  uint32(4856),
	98:  uint32(4885),
	99:  uint32(4916),
	100: uint32(4945),
	101: uint32(4974),
	102: uint32(5003),
	103: uint32(5032),
	104: uint32(5061),
	105: uint32(5090),
	106: uint32(5119),
	107: uint32(5148),
	108: uint32(5182),
	109: uint32(5226),
	110: uint32(5270),
	111: uint32(5314),
	112: uint32(5358),
	113: uint32(5386),
	114: uint32(5430),
	115: uint32(5472),
	116: uint32(5514),
	117: uint32(5558),
	118: uint32(5602),
	119: uint32(5636),
	120: uint32(5663),
	121: uint32(5690),
	122: uint32(5717),
	123: uint32(5760),
	124: uint32(5803),
	125: uint32(5844),
	126: uint32(5871),
	127: uint32(5914),
	128: uint32(5957),
	129: uint32(5984),
	130: uint32(6022),
	131: uint32(6062),
	132: uint32(6098),
	133: uint32(6138),
	134: uint32(6173),
	135: uint32(6210),
	136: uint32(6241),
	137: uint32(6274),
	138: uint32(6307),
	139: uint32(6344),
	140: uint32(6379),
	141: uint32(6412),
	142: uint32(6439),
	143: uint32(6472),
	144: uint32(6509),
	145: uint32(6542),
	146: uint32(6575),
	147: uint32(6610),
	148: uint32(6637),
	149: uint32(6661),
	150: uint32(6695),
	151: uint32(6725),
	152: uint32(6757),
	153: uint32(6779),
	154: uint32(6809),
	155: uint32(6837),
	156: uint32(6869),
	157: uint32(6903),
	158: uint32(6935),
	159: uint32(6967),
	160: uint32(6989),
	161: uint32(7011),
	162: uint32(7043),
	163: uint32(7075),
	164: uint32(7097),
	165: uint32(7129),
	166: uint32(7159),
	167: uint32(7181),
	168: uint32(7205),
	169: uint32(7229),
	170: uint32(7253),
	171: uint32(7277),
	172: uint32(7301),
	173: uint32(7325),
	174: uint32(7349),
	175: uint32(7383),
	176: uint32(7407),
	177: uint32(7431),
	178: uint32(7455),
	179: uint32(7479),
	180: uint32(7511),
	181: uint32(7535),
	182: uint32(7567),
	183: uint32(7591),
	184: uint32(7615),
	185: uint32(7639),
	186: uint32(7660),
	187: uint32(7685),
	188: uint32(7714),
	189: uint32(7739),
	190: uint32(7764),
	191: uint32(7789),
	192: uint32(7818),
	193: uint32(7847),
	194: uint32(7876),
	195: uint32(7905),
	196: uint32(7928),
	197: uint32(7957),
	198: uint32(7982),
	199: uint32(8013),
	200: uint32(8040),
	201: uint32(8065),
	202: uint32(8088),
	203: uint32(8108),
	204: uint32(8132),
	205: uint32(8150),
	206: uint32(8172),
	207: uint32(8192),
	208: uint32(8212),
	209: uint32(8236),
	210: uint32(8254),
	211: uint32(8272),
	212: uint32(8296),
	213: uint32(8314),
	214: uint32(8338),
	215: uint32(8358),
	216: uint32(8376),
	217: uint32(8396),
	218: uint32(8416),
	219: uint32(8442),
	220: uint32(8470),
	221: uint32(8498),
	222: uint32(8522),
	223: uint32(8548),
	224: uint32(8576),
	225: uint32(8596),
	226: uint32(8616),
	227: uint32(8636),
	228: uint32(8656),
	229: uint32(8679),
	230: uint32(8704),
	231: uint32(8723),
	232: uint32(8748),
	233: uint32(8767),
	234: uint32(8786),
	235: uint32(8805),
	236: uint32(8824),
	237: uint32(8849),
	238: uint32(8868),
	239: uint32(8887),
	240: uint32(8912),
	241: uint32(8937),
	242: uint32(8956),
	243: uint32(8975),
	244: uint32(9000),
	245: uint32(9021),
	246: uint32(9046),
	247: uint32(9065),
	248: uint32(9090),
	249: uint32(9113),
	250: uint32(9138),
	251: uint32(9157),
	252: uint32(9178),
	253: uint32(9203),
	254: uint32(9226),
	255: uint32(9251),
	256: uint32(9276),
	257: uint32(9301),
	258: uint32(9326),
	259: uint32(9351),
	260: uint32(9376),
	261: uint32(9394),
	262: uint32(9412),
	263: uint32(9434),
	264: uint32(9452),
	265: uint32(9474),
	266: uint32(9494),
	267: uint32(9512),
	268: uint32(9530),
	269: uint32(9548),
	270: uint32(9566),
	271: uint32(9588),
	272: uint32(9606),
	273: uint32(9624),
	274: uint32(9643),
	275: uint32(9662),
	276: uint32(9681),
	277: uint32(9698),
	278: uint32(9717),
	279: uint32(9736),
	280: uint32(9755),
	281: uint32(9772),
	282: uint32(9791),
	283: uint32(9810),
	284: uint32(9829),
	285: uint32(9848),
	286: uint32(9867),
	287: uint32(9886),
	288: uint32(9905),
	289: uint32(9924),
	290: uint32(9941),
	291: uint32(9958),
	292: uint32(9975),
	293: uint32(9994),
	294: uint32(10013),
	295: uint32(10032),
	296: uint32(10049),
	297: uint32(10066),
	298: uint32(10081),
	299: uint32(10100),
	300: uint32(10119),
	301: uint32(10138),
	302: uint32(10155),
	303: uint32(10172),
	304: uint32(10189),
	305: uint32(10208),
	306: uint32(10225),
	307: uint32(10244),
	308: uint32(10263),
	309: uint32(10282),
	310: uint32(10301),
	311: uint32(10320),
	312: uint32(10339),
	313: uint32(10356),
	314: uint32(10375),
	315: uint32(10394),
	316: uint32(10411),
	317: uint32(10430),
	318: uint32(10449),
	319: uint32(10466),
	320: uint32(10485),
	321: uint32(10504),
	322: uint32(10523),
	323: uint32(10542),
	324: uint32(10557),
	325: uint32(10573),
	326: uint32(10589),
	327: uint32(10605),
	328: uint32(10621),
	329: uint32(10637),
	330: uint32(10653),
	331: uint32(10669),
	332: uint32(10685),
	333: uint32(10701),
	334: uint32(10717),
	335: uint32(10733),
	336: uint32(10747),
	337: uint32(10763),
	338: uint32(10779),
	339: uint32(10795),
	340: uint32(10811),
	341: uint32(10827),
	342: uint32(10843),
	343: uint32(10859),
	344: uint32(10875),
	345: uint32(10891),
	346: uint32(10907),
	347: uint32(10923),
}

var ts_parse_actions = [818]TSParseActionEntry{
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
		Fextra: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(337)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token: uint8(TSParseActionTypeReduce),
		Fsymbol:     uint16(sym_source_file),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(167)),
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
		Fstate: uint16(libc.Int32FromInt32(152)),
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
		Fstate: uint16(libc.Int32FromInt32(264)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_shape_section),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(242)),
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
		Fstate: uint16(libc.Int32FromInt32(318)),
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
		Fstate: uint16(libc.Int32FromInt32(319)),
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
		Fstate: uint16(libc.Int32FromInt32(321)),
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
		Fstate: uint16(libc.Int32FromInt32(322)),
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
		Fstate: uint16(libc.Int32FromInt32(323)),
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
		Fstate: uint16(libc.Int32FromInt32(312)),
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
		Fstate: uint16(libc.Int32FromInt32(301)),
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
		Fstate: uint16(libc.Int32FromInt32(294)),
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
		Fstate: uint16(libc.Int32FromInt32(286)),
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
		Fstate: uint16(libc.Int32FromInt32(181)),
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
		Fstate: uint16(libc.Int32FromInt32(163)),
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
		Fstate: uint16(libc.Int32FromInt32(333)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_shape_section),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_shape_section_repeat1),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_shape_section_repeat1),
	})))),
	51: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	53: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_shape_section_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(318)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_shape_section_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(319)),
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
		Fsymbol:      uint16(aux_sym_shape_section_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(321)),
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
		Fsymbol:      uint16(aux_sym_shape_section_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(322)),
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
		Fsymbol:      uint16(aux_sym_shape_section_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(323)),
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
		Fsymbol:      uint16(aux_sym_shape_section_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(312)),
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
		Fsymbol:      uint16(aux_sym_shape_section_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(301)),
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
		Fsymbol:      uint16(aux_sym_shape_section_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(294)),
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
		Fsymbol:      uint16(aux_sym_shape_section_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(286)),
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
		Fsymbol:      uint16(aux_sym_shape_section_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(181)),
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
		Fsymbol:      uint16(aux_sym_shape_section_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(163)),
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
		Fsymbol:      uint16(aux_sym_shape_section_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(333)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_identifier),
	})))),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__namespace_identifier),
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
		Fcount: uint8(1),
	}})),
	93: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_identifier),
	})))),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_shape_id_repeat1),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_shape_id_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(287)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_shape_id),
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
		Fstate: uint16(libc.Int32FromInt32(287)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	104: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_shape_id),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	106: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_root_shape_id),
	})))),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_shape_id_repeat1),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_shape_id_member),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_absolute_root_shape_id),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_trait_statement),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(113)),
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
		Fstate: uint16(libc.Int32FromInt32(34)),
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
		Fstate: uint16(libc.Int32FromInt32(20)),
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
		Fstate: uint16(libc.Int32FromInt32(208)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(209)),
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
		Fstate: uint16(libc.Int32FromInt32(275)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(229)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(269)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_node_object),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_node_object),
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
		Fcount:    uint8(1),
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
		Fsymbol:      uint16(sym_trait_body),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_trait_statement),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	150: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_trait_body),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_namespace),
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
		Fstate: uint16(libc.Int32FromInt32(310)),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_namespace_repeat1),
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
		Fsymbol:      uint16(aux_sym_namespace_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(310)),
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
		Fcount:    uint8(1),
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
		Fsymbol:      uint16(sym_namespace),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_shape_statement_repeat1),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	165: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_shape_statement_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(181)),
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
		Fstate: uint16(libc.Int32FromInt32(118)),
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
		Fstate: uint16(libc.Int32FromInt32(36)),
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
		Fstate: uint16(libc.Int32FromInt32(228)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(92)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(90)),
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
		Fstate: uint16(libc.Int32FromInt32(300)),
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
		Fstate: uint16(libc.Int32FromInt32(93)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(95)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(196)),
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
		Fstate: uint16(libc.Int32FromInt32(147)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(83)),
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
		Fstate: uint16(libc.Int32FromInt32(170)),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_node_array_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(118)),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_node_array_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(36)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_node_array_repeat1),
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
		Fsymbol:      uint16(aux_sym_node_array_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(92)),
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
		Fsymbol:      uint16(aux_sym_node_array_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(90)),
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
		Fsymbol:      uint16(aux_sym_node_array_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(300)),
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
		Fsymbol:      uint16(aux_sym_node_array_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(93)),
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
		Fsymbol:      uint16(aux_sym_node_array_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(95)),
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
		Fsymbol:      uint16(aux_sym_node_array_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(196)),
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
		Fsymbol:      uint16(aux_sym_node_array_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(147)),
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
		Fsymbol:      uint16(aux_sym_node_array_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(83)),
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
		Fstate: uint16(libc.Int32FromInt32(103)),
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
		Fsymbol:      uint16(sym_simple_shape_statement),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(182)),
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
		Fstate: uint16(libc.Int32FromInt32(101)),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_shape_statement_repeat1),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_mixins),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(35)),
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
		Fstate: uint16(libc.Int32FromInt32(172)),
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
		Fcount: uint8(1),
	}})),
	246: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(189)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_apply_statement),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_enum_members),
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
		Fsymbol:      uint16(sym_simple_shape_statement),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_apply_statement_singular),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_shape_members),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_resource_statement),
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
		Fsymbol:      uint16(sym_operation_statement),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_service_statement),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_union_statement),
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
		Fsymbol:      uint16(sym_structure_statement),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_set_statement),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_map_statement),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_list_statement),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_enum_statement),
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
		Fsymbol:      uint16(aux_sym_shape_section_repeat1),
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
		Fsymbol:      uint16(sym__definition),
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
		Fsymbol:      uint16(sym_shape_statement),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_shape_body),
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
		Fsymbol:      uint16(sym_list_statement),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_enum_statement),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_map_statement),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_set_statement),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_structure_statement),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_union_statement),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_service_statement),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_operation_body),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_operation_statement),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_resource_statement),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_namespace_statement),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_apply_statement_block),
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
		Fsymbol:      uint16(sym_enum_members),
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
		Fsymbol:      uint16(sym_shape_members),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_structure_statement),
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
		Fsymbol:      uint16(sym_operation_body),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_apply_statement_block),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_use_statement),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_shape_statement),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fsymbol:      uint16(sym_shape_id),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_shape_id_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(285)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(1),
	}})),
	341: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_shape_id_repeat1),
	})))),
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
		Fcount: uint8(1),
	}})),
	343: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__namespace_identifier),
	})))),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_shape_id),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_root_shape_id),
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
		Fcount: uint8(1),
	}})),
	349: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_shape_id_repeat1),
	})))),
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
		Fcount: uint8(1),
	}})),
	351: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_shape_id_member),
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
		Fsymbol:      uint16(sym_absolute_root_shape_id),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_literal),
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
		Fcount: uint8(1),
	}})),
	357: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_literal),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_string),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_string),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_boolean),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_boolean),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_number),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_number),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_shape_id_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(309)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_float),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_float),
	})))),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_node_value),
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
		Fcount: uint8(1),
	}})),
	380: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_node_value),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(sym__string_literal),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym__string_literal),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__multiline_string_literal),
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
		Fcount: uint8(1),
	}})),
	388: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__multiline_string_literal),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_node_object),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_node_array),
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
		Fcount: uint8(1),
	}})),
	394: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_node_array),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_number),
	})))),
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
		Fcount: uint8(1),
	}})),
	398: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_number),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_node_array),
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
		Fcount: uint8(1),
	}})),
	402: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_node_array),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_node_array_repeat1),
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
		Fcount: uint8(1),
	}})),
	406: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_node_array_repeat1),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_node_object),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__multiline_string_literal),
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
		Fcount: uint8(1),
	}})),
	412: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__multiline_string_literal),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	414: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__string_literal),
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
		Fcount: uint8(1),
	}})),
	416: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__string_literal),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_float),
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
		Fcount: uint8(1),
	}})),
	420: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_float),
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
		Fstate: uint16(libc.Int32FromInt32(309)),
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
		Fstate: uint16(libc.Int32FromInt32(171)),
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
		Fstate: uint16(libc.Int32FromInt32(6)),
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
		Fstate: uint16(libc.Int32FromInt32(18)),
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
		Fsymbol:      uint16(sym_trait_structure),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_control_var_name),
		Fproduction_id: uint16(1),
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
		Fcount: uint8(1),
	}})),
	436: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_control_var_name),
		Fproduction_id: uint16(1),
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
		Fstate: uint16(libc.Int32FromInt32(105)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_trait_structure_repeat1),
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
		Fcount: uint8(2),
	}})),
	442: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_trait_structure_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(195)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_trait_structure_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(138)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	448: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_trait_structure_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(6)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_node_object_repeat1),
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
		Fcount: uint8(2),
	}})),
	453: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_node_object_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(195)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	456: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_node_object_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(138)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_node_object_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(6)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(99)),
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
		Fstate: uint16(libc.Int32FromInt32(150)),
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
	466: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(293)),
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
		Fstate: uint16(libc.Int32FromInt32(157)),
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
		Fstate: uint16(libc.Int32FromInt32(292)),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_shape_members_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(315)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_shape_members_repeat1),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	479: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_shape_members_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(157)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	482: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_shape_members_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(6)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(73)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(38)),
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
		Fstate: uint16(libc.Int32FromInt32(100)),
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
		Fstate: uint16(libc.Int32FromInt32(125)),
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
		Fstate: uint16(libc.Int32FromInt32(153)),
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
		Fsymbol:      uint16(aux_sym_mixins_repeat1),
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
		Fsymbol:      uint16(aux_sym_mixins_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(100)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
	506: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__multiline_string_literal_repeat1),
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
		Fcount: uint8(2),
	}})),
	508: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__multiline_string_literal_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(214)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	511: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__multiline_string_literal_repeat1),
	})))),
	512: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(libc.Int32FromInt32(203)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount: uint8(2),
	}})),
	514: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__multiline_string_literal_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(217)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	517: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__multiline_string_literal_repeat1),
	})))),
	518: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
	520: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(203)),
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
		Fcount: uint8(1),
	}})),
	526: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(41)),
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
		Fstate: uint16(libc.Int32FromInt32(166)),
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
		Fstate: uint16(libc.Int32FromInt32(122)),
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
		Fsymbol:      uint16(aux_sym_enum_members_repeat1),
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
		Fsymbol:      uint16(aux_sym_enum_members_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(166)),
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
		Fsymbol:      uint16(aux_sym_enum_members_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(122)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(106)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(161)),
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
		Fstate: uint16(libc.Int32FromInt32(72)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(186)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(98)),
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
		Fstate: uint16(libc.Int32FromInt32(75)),
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
		Fstate: uint16(libc.Int32FromInt32(5)),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_control_section),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_control_section_repeat1),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_control_section_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(167)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(67)),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_operation_body_repeat1),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	571: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_operation_body_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(6)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(329)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__multiline_string_fragment_repeat1),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym__multiline_string_fragment_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(212)),
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
		Fcount:    uint8(1),
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
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_shape_member),
		Fproduction_id: uint16(7),
	})))),
	582: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	583: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	584: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	585: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	586: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	587: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	588: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	589: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	590: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	591: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	592: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	593: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_shape_member),
		Fproduction_id: uint16(9),
	})))),
	594: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	595: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_shape_member),
	})))),
	596: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(107)),
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
		Fcount: uint8(1),
	}})),
	599: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(154)),
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
		Fcount: uint8(1),
	}})),
	603: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
	605: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__multiline_string_fragment),
	})))),
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
		Fcount: uint8(1),
	}})),
	607: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(175)),
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
		Fstate: uint16(libc.Int32FromInt32(129)),
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
		Fcount: uint8(1),
	}})),
	613: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__string_literal_repeat1),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	615: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__string_literal_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(235)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	618: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__string_literal_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(236)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym__string_literal_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(237)),
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
		Fsymbol:      uint16(sym_shape_member),
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
		Fcount: uint8(1),
	}})),
	626: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__multiline_string_fragment),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_shape_member_elided),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_metadata_section_repeat1),
	})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	632: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_metadata_section_repeat1),
	})))),
	633: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
	635: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym__multiline_string_literal_repeat1),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_node_object_kvp),
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
		Fcount: uint8(1),
	}})),
	639: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_node_object_kvp),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_enum_member),
		Fproduction_id: uint16(6),
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
		Fcount: uint8(1),
	}})),
	645: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__escape_sequence),
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
		Fcount: uint8(1),
	}})),
	647: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym__multiline_string_fragment_repeat1),
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
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_enum_member),
		Fproduction_id: uint16(3),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_operation_errors_repeat1),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_operation_errors_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(122)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_shape_statement_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(157)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_metadata_section),
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
		Fstate: uint16(libc.Int32FromInt32(325)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_operation_member),
	})))),
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
		Fsymbol:      uint16(sym_value_assignment),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_shape_member),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(aux_sym__string_literal_repeat1),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym__string_literal_repeat1),
	})))),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__escape_sequence),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_control_section_repeat1),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_shape_member),
		Fproduction_id: uint16(7),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(71)),
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
		Fstate: uint16(libc.Int32FromInt32(30)),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_control_statement),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_shape_members_repeat1),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_node_object_repeat1),
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
		Fcount: uint8(1),
	}})),
	693: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_node_object_repeat1),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_shape_member),
		Fproduction_id: uint16(9),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_source_file),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_operation_member_repeat1),
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
		Fsymbol:      uint16(aux_sym_operation_member_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(6)),
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
		Fstate: uint16(libc.Int32FromInt32(158)),
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
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(aux_sym_trait_structure_repeat1),
		Fproduction_id: uint16(5),
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
		Fcount: uint8(1),
	}})),
	708: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(aux_sym_trait_structure_repeat1),
		Fproduction_id: uint16(5),
	})))),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_shape_statement_repeat1),
	})))),
	711: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(140)),
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
		Fstate: uint16(libc.Int32FromInt32(76)),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_enum_members_repeat1),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_enum_member),
		Fproduction_id: uint16(6),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_metadata_statement),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(135)),
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
		Fstate: uint16(libc.Int32FromInt32(134)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_node_object_key),
		Fproduction_id: uint16(2),
	})))),
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
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_enum_member),
		Fproduction_id: uint16(3),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_metadata_section_repeat1),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(216)),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_operation_member),
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
		Fsymbol:      uint16(aux_sym_mixins_repeat1),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
	745: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fsymbol:      uint16(sym_node_object_key),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(aux_sym_operation_member_repeat1),
		Fproduction_id: uint16(8),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_inline_structure),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_inline_structure),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_operation_errors),
		Fproduction_id: uint16(10),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(102)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(108)),
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
		Fstate: uint16(libc.Int32FromInt32(278)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_operation_body_repeat1),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_inline_structure),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_structure_resource),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_inline_structure),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_operation_error),
	})))),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_operation_errors_repeat1),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_operation_errors),
		Fproduction_id: uint16(10),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_trait_body_value),
		Fproduction_id: uint16(4),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__control_identifier),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(164)),
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
		Fstate: uint16(libc.Int32FromInt32(183)),
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
		Fstate: uint16(libc.Int32FromInt32(288)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_primitive),
	})))),
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
	793: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f uint8_t
		_ [7]byte
	}{f: uint8(TSParseActionTypeAccept)})),
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
		Fstate: uint16(libc.Int32FromInt32(271)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(349)),
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
		Fstate: uint16(libc.Int32FromInt32(22)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_trait_body_value),
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
		Fstate: uint16(libc.Int32FromInt32(65)),
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
		Fstate: uint16(libc.Int32FromInt32(148)),
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
		Fstate: uint16(libc.Int32FromInt32(77)),
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
		Fstate: uint16(libc.Int32FromInt32(308)),
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
		Fstate: uint16(libc.Int32FromInt32(114)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_documentation_comment),
	})))),
}

func tree_sitter_smithy(tls *libc.TLS) (r uintptr) {
	return uintptr(unsafe.Pointer(&language))
}

var language = TSLanguage{
	Fversion:                   uint32(LANGUAGE_VERSION),
	Fsymbol_count:              uint32(SYMBOL_COUNT),
	Falias_count:               uint32(ALIAS_COUNT),
	Ftoken_count:               uint32(TOKEN_COUNT),
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

var __ccgo_ts1 = "end\x00$\x00:\x00metadata\x00=\x00namespace\x00.\x00use\x00#\x00enum\x00intEnum\x00{\x00}\x00list\x00map\x00set\x00structure\x00union\x00service\x00operation\x00resource\x00,\x00[\x00]\x00:=\x00@\x00(\x00)\x00apply\x00with\x00for\x00blob\x00boolean\x00byte\x00document\x00double\x00float\x00integer\x00long\x00short\x00string\x00timestamp\x00bigInteger\x00bigDecimal\x00true\x00false\x00null\x00-\x00number_token1\x00float_token1\x00\"\x00\"\"\"\x00string_fragment\x00_multiline_string_fragment_token1\x00_multiline_string_fragment_token2\x00_escape_sequence_token1\x00escape_sequence\x00identifier_token1\x00comment\x00///\x00documentation_comment_token1\x00source_file\x00control_section\x00control_statement\x00control_var_name\x00metadata_section\x00metadata_statement\x00shape_section\x00namespace_statement\x00_definition\x00use_statement\x00shape_statement\x00shape_body\x00absolute_root_shape_id\x00root_shape_id\x00shape_id_member\x00shape_id\x00simple_shape_statement\x00enum_statement\x00enum_members\x00enum_member\x00list_statement\x00map_statement\x00set_statement\x00structure_statement\x00union_statement\x00service_statement\x00operation_statement\x00resource_statement\x00shape_members\x00shape_member\x00shape_member_elided\x00operation_body\x00operation_member\x00operation_errors\x00operation_error\x00inline_structure\x00trait_statement\x00trait_body\x00trait_body_value\x00trait_structure\x00apply_statement\x00apply_statement_singular\x00apply_statement_block\x00mixins\x00structure_resource\x00value_assignment\x00node_value\x00node_array\x00node_object\x00node_object_kvp\x00node_object_key\x00literal\x00primitive\x00number\x00_string_literal\x00_multiline_string_literal\x00multiline_string_fragment\x00_escape_sequence\x00identifier\x00control_key\x00_namespace_identifier\x00documentation_comment\x00control_section_repeat1\x00metadata_section_repeat1\x00shape_section_repeat1\x00namespace_repeat1\x00shape_statement_repeat1\x00shape_id_repeat1\x00enum_members_repeat1\x00shape_members_repeat1\x00operation_body_repeat1\x00operation_member_repeat1\x00operation_errors_repeat1\x00trait_structure_repeat1\x00mixins_repeat1\x00node_array_repeat1\x00node_object_repeat1\x00_string_literal_repeat1\x00_multiline_string_literal_repeat1\x00_multiline_string_fragment_repeat1\x00enum_field\x00field\x00key_identifier\x00operation_error_field\x00operation_field\x00trait_node_value\x00trait_object_kvp\x00"
