// Code generated for windows/amd64 by 'ccgo -ignore-unsupported-alignment -ignore-unsupported-atomic-sizes -ignore-vector-functions --winapi-no-errno -ignore-link-errors -std=gnu11 -O0 -D__extension__= -D__forceinline=static inline -D__attribute__(...)= -D__declspec(x)= -D__cdecl= -D__stdcall= -D__fastcall= -D__thiscall= -D_cdecl= -D__restrict= -D__restrict__= -D__MINGW_EXTENSION= -D_X86INTRIN_H_INCLUDED -D_X86GPRINTRIN_H_INCLUDED -D_IMMINTRIN_H_INCLUDED -D_MMINTRIN_H_INCLUDED -D_XMMINTRIN_H_INCLUDED -D_EMMINTRIN_H_INCLUDED -D_PMMINTRIN_H_INCLUDED -D_MM3DNOW_H_INCLUDED -Dfunc=func_token -Dinterface=interface_token -Dselect=select_token -Dchan=chan_token -Dgo=go_token -Dmap=map_token -Dpackage=package_token -Dtype=type_token -Dvar=var_token -Dimport=import_token -Ddefer=defer_token -Dfallthrough=fallthrough_token -Drange=range_token -I D:\a\ccgo-tree-sitter\ccgo-tree-sitter\third-party\tree-sitter-cyberchef\src -I D:\a\ccgo-tree-sitter\ccgo-tree-sitter\third-party\tree-sitter-cyberchef -I D:\a\ccgo-tree-sitter\ccgo-tree-sitter\third-party -I C:\Users\runneradmin\.cache\workspaced\sources\github\7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2\lib\include -I C:\Users\runneradmin\.cache\workspaced\sources\github\7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2\lib\src D:\a\ccgo-tree-sitter\ccgo-tree-sitter\third-party\tree-sitter-cyberchef\src\parser.c -o grammar.go', DO NOT EDIT.

//go:build windows && amd64

package grammar_cyberchef

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
const FIELD_COUNT = 4
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
const MAX_ALIAS_SEQUENCE_LENGTH = 7
const MAX_RESERVED_WORD_SET_SIZE = 0
const MB_LEN_MAX = 5
const MINGW_HAS_DDK_H = 1
const MINGW_HAS_SECURE_API = 1
const PATH_MAX = 260
const PRODUCTION_ID_COUNT = 13
const PTRDIFF_MAX = "INT64_MAX"
const PTRDIFF_MIN = "INT64_MIN"
const RAND_MAX = 0x7fff
const SIG_ATOMIC_MAX = "INT32_MAX"
const SIG_ATOMIC_MIN = "INT32_MIN"
const SIZE_MAX = "UINT64_MAX"
const SSIZE_MAX = "_I64_MAX"
const STATE_COUNT = 48
const STRUNCATE = 80
const SUPERTYPE_COUNT = 0
const SYMBOL_COUNT = 25
const TOKEN_COUNT = 14
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

const aux_sym_recipe_token1 = 1
const anon_sym_LPAREN = 2
const anon_sym_RPAREN = 3
const sym_name = 4
const sym_string = 5
const sym_number = 6
const sym_boolean = 7
const anon_sym_LBRACE = 8
const anon_sym_COMMA = 9
const anon_sym_RBRACE = 10
const anon_sym_COLON = 11
const anon_sym_SLASHdisabled = 12
const anon_sym_SLASHbreakpoint = 13
const sym_recipe = 14
const sym_operation = 15
const sym_object = 16
const sym_arguments = 17
const sym__argument = 18
const sym__pair = 19
const sym__disabled = 20
const sym__breakpoint = 21
const aux_sym_recipe_repeat1 = 22
const aux_sym_object_repeat1 = 23
const aux_sym_arguments_repeat1 = 24

var ts_symbol_names = [25]uintptr{
	0:  __ccgo_ts,
	1:  __ccgo_ts + 4,
	2:  __ccgo_ts + 18,
	3:  __ccgo_ts + 20,
	4:  __ccgo_ts + 22,
	5:  __ccgo_ts + 27,
	6:  __ccgo_ts + 34,
	7:  __ccgo_ts + 41,
	8:  __ccgo_ts + 49,
	9:  __ccgo_ts + 51,
	10: __ccgo_ts + 53,
	11: __ccgo_ts + 55,
	12: __ccgo_ts + 57,
	13: __ccgo_ts + 67,
	14: __ccgo_ts + 79,
	15: __ccgo_ts + 86,
	16: __ccgo_ts + 96,
	17: __ccgo_ts + 103,
	18: __ccgo_ts + 113,
	19: __ccgo_ts + 123,
	20: __ccgo_ts + 129,
	21: __ccgo_ts + 139,
	22: __ccgo_ts + 151,
	23: __ccgo_ts + 166,
	24: __ccgo_ts + 181,
}

var ts_symbol_map = [25]TSSymbol{
	1:  uint16(aux_sym_recipe_token1),
	2:  uint16(anon_sym_LPAREN),
	3:  uint16(anon_sym_RPAREN),
	4:  uint16(sym_name),
	5:  uint16(sym_string),
	6:  uint16(sym_number),
	7:  uint16(sym_boolean),
	8:  uint16(anon_sym_LBRACE),
	9:  uint16(anon_sym_COMMA),
	10: uint16(anon_sym_RBRACE),
	11: uint16(anon_sym_COLON),
	12: uint16(anon_sym_SLASHdisabled),
	13: uint16(anon_sym_SLASHbreakpoint),
	14: uint16(sym_recipe),
	15: uint16(sym_operation),
	16: uint16(sym_object),
	17: uint16(sym_arguments),
	18: uint16(sym__argument),
	19: uint16(sym__pair),
	20: uint16(sym__disabled),
	21: uint16(sym__breakpoint),
	22: uint16(aux_sym_recipe_repeat1),
	23: uint16(aux_sym_object_repeat1),
	24: uint16(aux_sym_arguments_repeat1),
}

var ts_symbol_metadata = [25]TSSymbolMetadata{
	0: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	1: {},
	2: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	3: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	4: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	5: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	6: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	7: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
	19: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	20: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	21: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	22: {},
	23: {},
	24: {},
}

type ts_field_identifiers = int32

const field_breakpoint = 1
const field_disabled = 2
const field_key = 3
const field_value = 4

var ts_field_names = [5]uintptr{
	0: libc.UintptrFromInt32(0),
	1: __ccgo_ts + 199,
	2: __ccgo_ts + 210,
	3: __ccgo_ts + 219,
	4: __ccgo_ts + 223,
}

var ts_field_map_slices = [13]TSMapSlice{
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
		Flength: uint16(1),
	},
	5: {
		Findex:  uint16(5),
		Flength: uint16(1),
	},
	6: {
		Findex:  uint16(6),
		Flength: uint16(2),
	},
	7: {
		Findex:  uint16(8),
		Flength: uint16(4),
	},
	8: {
		Findex:  uint16(12),
		Flength: uint16(4),
	},
	9: {
		Findex:  uint16(16),
		Flength: uint16(1),
	},
	10: {
		Findex:  uint16(17),
		Flength: uint16(1),
	},
	11: {
		Findex:  uint16(18),
		Flength: uint16(2),
	},
	12: {
		Findex:  uint16(20),
		Flength: uint16(2),
	},
}

var ts_field_map_entries = [22]TSFieldMapEntry{
	0: {
		Ffield_id: uint16(field_disabled),
	},
	1: {
		Ffield_id: uint16(field_breakpoint),
	},
	2: {
		Ffield_id:    uint16(field_key),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	3: {
		Ffield_id:    uint16(field_value),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	4: {
		Ffield_id:    uint16(field_disabled),
		Fchild_index: uint8(2),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	5: {
		Ffield_id:    uint16(field_breakpoint),
		Fchild_index: uint8(2),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	6: {
		Ffield_id: uint16(field_key),
	},
	7: {
		Ffield_id:    uint16(field_value),
		Fchild_index: uint8(2),
	},
	8: {
		Ffield_id:    uint16(field_key),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	9: {
		Ffield_id:    uint16(field_key),
		Fchild_index: uint8(2),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	10: {
		Ffield_id:    uint16(field_value),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	11: {
		Ffield_id:    uint16(field_value),
		Fchild_index: uint8(2),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	12: {
		Ffield_id:  uint16(field_key),
		Finherited: libc.BoolUint8(true1 != 0),
	},
	13: {
		Ffield_id:    uint16(field_key),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	14: {
		Ffield_id:  uint16(field_value),
		Finherited: libc.BoolUint8(true1 != 0),
	},
	15: {
		Ffield_id:    uint16(field_value),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	16: {
		Ffield_id:    uint16(field_disabled),
		Fchild_index: uint8(3),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	17: {
		Ffield_id:    uint16(field_breakpoint),
		Fchild_index: uint8(3),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	18: {
		Ffield_id:    uint16(field_breakpoint),
		Fchild_index: uint8(3),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	19: {
		Ffield_id:    uint16(field_disabled),
		Fchild_index: uint8(2),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	20: {
		Ffield_id:    uint16(field_breakpoint),
		Fchild_index: uint8(4),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	21: {
		Ffield_id:    uint16(field_disabled),
		Fchild_index: uint8(3),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
}

var ts_alias_sequences = [13][7]TSSymbol{}

var ts_non_terminal_alias_map = [1]uint16_t{}

var ts_primary_state_ids = [48]TSStateId{
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
}

func ts_lex(tls *libc.TLS, lexer uintptr, state TSStateId) (r uint8) {
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
		if eof != 0 {
			state = uint16(30)
			goto next_state
		}
		i = uint32(0)
		for {
			if !(uint64(i) < libc.Uint64FromInt64(44)/libc.Uint64FromInt64(2)) {
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
			state = uint16(0)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(37)
			goto next_state
		}
		return result
	case int32(1):
		if lookahead == int32('\n') {
			state = uint16(31)
			goto next_state
		}
		if lookahead == int32('\r') {
			state = uint16(1)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\f') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(1)
			goto next_state
		}
		if int32('/') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(34)
			goto next_state
		}
		return result
	case int32(2):
		if lookahead == int32('\'') {
			state = uint16(35)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(3)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(2)
			goto next_state
		}
		return result
	case int32(3):
		if lookahead == int32('\'') {
			state = uint16(36)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(3)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(2)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead == int32('a') {
			state = uint16(16)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead == int32('a') {
			state = uint16(15)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead == int32('a') {
			state = uint16(8)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead == int32('b') {
			state = uint16(22)
			goto next_state
		}
		if lookahead == int32('d') {
			state = uint16(13)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead == int32('b') {
			state = uint16(17)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead == int32('d') {
			state = uint16(44)
			goto next_state
		}
		return result
	case int32(10):
		if lookahead == int32('e') {
			state = uint16(39)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead == int32('e') {
			state = uint16(9)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead == int32('e') {
			state = uint16(5)
			goto next_state
		}
		return result
	case int32(13):
		if lookahead == int32('i') {
			state = uint16(24)
			goto next_state
		}
		return result
	case int32(14):
		if lookahead == int32('i') {
			state = uint16(18)
			goto next_state
		}
		return result
	case int32(15):
		if lookahead == int32('k') {
			state = uint16(20)
			goto next_state
		}
		return result
	case int32(16):
		if lookahead == int32('l') {
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(17):
		if lookahead == int32('l') {
			state = uint16(11)
			goto next_state
		}
		return result
	case int32(18):
		if lookahead == int32('n') {
			state = uint16(25)
			goto next_state
		}
		return result
	case int32(19):
		if lookahead == int32('o') {
			state = uint16(14)
			goto next_state
		}
		return result
	case int32(20):
		if lookahead == int32('p') {
			state = uint16(19)
			goto next_state
		}
		return result
	case int32(21):
		if lookahead == int32('r') {
			state = uint16(26)
			goto next_state
		}
		return result
	case int32(22):
		if lookahead == int32('r') {
			state = uint16(12)
			goto next_state
		}
		return result
	case int32(23):
		if lookahead == int32('s') {
			state = uint16(10)
			goto next_state
		}
		return result
	case int32(24):
		if lookahead == int32('s') {
			state = uint16(6)
			goto next_state
		}
		return result
	case int32(25):
		if lookahead == int32('t') {
			state = uint16(45)
			goto next_state
		}
		return result
	case int32(26):
		if lookahead == int32('u') {
			state = uint16(10)
			goto next_state
		}
		return result
	case int32(27):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(37)
			goto next_state
		}
		return result
	case int32(28):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(38)
			goto next_state
		}
		return result
	case int32(29):
		if eof != 0 {
			state = uint16(30)
			goto next_state
		}
		if lookahead == int32('\n') {
			state = uint16(31)
			goto next_state
		}
		if lookahead == int32('\r') {
			state = uint16(1)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\f') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(29)
			goto next_state
		}
		if int32('/') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(34)
			goto next_state
		}
		return result
	case int32(30):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(ts_builtin_sym_end)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(31):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_recipe_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(31)
			goto next_state
		}
		if lookahead == int32('\r') {
			state = uint16(1)
			goto next_state
		}
		return result
	case int32(32):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(33):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(34):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('/') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(34)
			goto next_state
		}
		return result
	case int32(35):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(36):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\'') {
			state = uint16(35)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(3)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(2)
			goto next_state
		}
		return result
	case int32(37):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(28)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(37)
			goto next_state
		}
		return result
	case int32(38):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(38)
			goto next_state
		}
		return result
	case int32(39):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_boolean)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(40):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(41):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COMMA)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(42):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(43):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COLON)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(44):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SLASHdisabled)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(45):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SLASHbreakpoint)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var map_token = [22]uint16_t{
	0:  uint16('\''),
	1:  uint16(2),
	2:  uint16('('),
	3:  uint16(32),
	4:  uint16(')'),
	5:  uint16(33),
	6:  uint16(','),
	7:  uint16(41),
	8:  uint16('-'),
	9:  uint16(27),
	10: uint16('/'),
	11: uint16(7),
	12: uint16(':'),
	13: uint16(43),
	14: uint16('f'),
	15: uint16(4),
	16: uint16('t'),
	17: uint16(21),
	18: uint16('{'),
	19: uint16(40),
	20: uint16('}'),
	21: uint16(42),
}

var ts_lex_modes = [48]TSLexerMode{
	0: {},
	1: {
		Flex_state: uint16(29),
	},
	2: {},
	3: {},
	4: {},
	5: {
		Flex_state: uint16(29),
	},
	6: {},
	7: {
		Flex_state: uint16(29),
	},
	8:  {},
	9:  {},
	10: {},
	11: {},
	12: {},
	13: {},
	14: {},
	15: {
		Flex_state: uint16(29),
	},
	16: {},
	17: {},
	18: {},
	19: {
		Flex_state: uint16(29),
	},
	20: {
		Flex_state: uint16(29),
	},
	21: {
		Flex_state: uint16(29),
	},
	22: {},
	23: {
		Flex_state: uint16(29),
	},
	24: {
		Flex_state: uint16(29),
	},
	25: {
		Flex_state: uint16(29),
	},
	26: {
		Flex_state: uint16(29),
	},
	27: {},
	28: {},
	29: {},
	30: {},
	31: {},
	32: {},
	33: {},
	34: {
		Flex_state: uint16(29),
	},
	35: {},
	36: {
		Flex_state: uint16(29),
	},
	37: {},
	38: {
		Flex_state: uint16(29),
	},
	39: {},
	40: {},
	41: {
		Flex_state: uint16(29),
	},
	42: {
		Flex_state: uint16(29),
	},
	43: {
		Flex_state: uint16(29),
	},
	44: {},
	45: {},
	46: {
		Flex_state: uint16(29),
	},
	47: {
		Flex_state: uint16(29),
	},
}

var ts_parse_table = [2][25]uint16_t{
	0: {
		0:  uint16(1),
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
	},
	1: {
		0:  uint16(3),
		1:  uint16(5),
		4:  uint16(7),
		14: uint16(39),
		15: uint16(7),
		22: uint16(7),
	},
}

var ts_small_parse_table = [367]uint16_t{
	0:   uint16(9),
	1:   uint16(9),
	2:   uint16(1),
	3:   uint16(anon_sym_RPAREN),
	4:   uint16(13),
	5:   uint16(1),
	6:   uint16(anon_sym_LBRACE),
	7:   uint16(15),
	8:   uint16(1),
	9:   uint16(anon_sym_SLASHdisabled),
	10:  uint16(17),
	11:  uint16(1),
	12:  uint16(anon_sym_SLASHbreakpoint),
	13:  uint16(8),
	14:  uint16(1),
	15:  uint16(sym_arguments),
	16:  uint16(22),
	17:  uint16(1),
	18:  uint16(sym__disabled),
	19:  uint16(44),
	20:  uint16(1),
	21:  uint16(sym__breakpoint),
	22:  uint16(6),
	23:  uint16(2),
	24:  uint16(sym_object),
	25:  uint16(sym__argument),
	26:  uint16(11),
	27:  uint16(3),
	28:  uint16(sym_string),
	29:  uint16(sym_number),
	30:  uint16(sym_boolean),
	31:  uint16(3),
	32:  uint16(13),
	33:  uint16(1),
	34:  uint16(anon_sym_LBRACE),
	35:  uint16(13),
	36:  uint16(2),
	37:  uint16(sym_object),
	38:  uint16(sym__argument),
	39:  uint16(19),
	40:  uint16(3),
	41:  uint16(sym_string),
	42:  uint16(sym_number),
	43:  uint16(sym_boolean),
	44:  uint16(3),
	45:  uint16(13),
	46:  uint16(1),
	47:  uint16(anon_sym_LBRACE),
	48:  uint16(29),
	49:  uint16(2),
	50:  uint16(sym_object),
	51:  uint16(sym__argument),
	52:  uint16(21),
	53:  uint16(3),
	54:  uint16(sym_string),
	55:  uint16(sym_number),
	56:  uint16(sym_boolean),
	57:  uint16(4),
	58:  uint16(23),
	59:  uint16(1),
	61:  uint16(25),
	62:  uint16(1),
	63:  uint16(aux_sym_recipe_token1),
	64:  uint16(28),
	65:  uint16(1),
	66:  uint16(sym_name),
	67:  uint16(5),
	68:  uint16(2),
	69:  uint16(sym_operation),
	70:  uint16(aux_sym_recipe_repeat1),
	71:  uint16(3),
	72:  uint16(33),
	73:  uint16(1),
	74:  uint16(anon_sym_COMMA),
	75:  uint16(9),
	76:  uint16(1),
	77:  uint16(aux_sym_arguments_repeat1),
	78:  uint16(31),
	79:  uint16(3),
	80:  uint16(anon_sym_RPAREN),
	81:  uint16(anon_sym_SLASHdisabled),
	82:  uint16(anon_sym_SLASHbreakpoint),
	83:  uint16(4),
	84:  uint16(7),
	85:  uint16(1),
	86:  uint16(sym_name),
	87:  uint16(35),
	88:  uint16(1),
	90:  uint16(37),
	91:  uint16(1),
	92:  uint16(aux_sym_recipe_token1),
	93:  uint16(5),
	94:  uint16(2),
	95:  uint16(sym_operation),
	96:  uint16(aux_sym_recipe_repeat1),
	97:  uint16(5),
	98:  uint16(15),
	99:  uint16(1),
	100: uint16(anon_sym_SLASHdisabled),
	101: uint16(17),
	102: uint16(1),
	103: uint16(anon_sym_SLASHbreakpoint),
	104: uint16(39),
	105: uint16(1),
	106: uint16(anon_sym_RPAREN),
	107: uint16(17),
	108: uint16(1),
	109: uint16(sym__disabled),
	110: uint16(35),
	111: uint16(1),
	112: uint16(sym__breakpoint),
	113: uint16(3),
	114: uint16(33),
	115: uint16(1),
	116: uint16(anon_sym_COMMA),
	117: uint16(12),
	118: uint16(1),
	119: uint16(aux_sym_arguments_repeat1),
	120: uint16(41),
	121: uint16(3),
	122: uint16(anon_sym_RPAREN),
	123: uint16(anon_sym_SLASHdisabled),
	124: uint16(anon_sym_SLASHbreakpoint),
	125: uint16(1),
	126: uint16(43),
	127: uint16(5),
	128: uint16(anon_sym_RPAREN),
	129: uint16(anon_sym_COMMA),
	130: uint16(anon_sym_RBRACE),
	131: uint16(anon_sym_SLASHdisabled),
	132: uint16(anon_sym_SLASHbreakpoint),
	133: uint16(1),
	134: uint16(45),
	135: uint16(5),
	136: uint16(anon_sym_RPAREN),
	137: uint16(anon_sym_COMMA),
	138: uint16(anon_sym_RBRACE),
	139: uint16(anon_sym_SLASHdisabled),
	140: uint16(anon_sym_SLASHbreakpoint),
	141: uint16(3),
	142: uint16(49),
	143: uint16(1),
	144: uint16(anon_sym_COMMA),
	145: uint16(12),
	146: uint16(1),
	147: uint16(aux_sym_arguments_repeat1),
	148: uint16(47),
	149: uint16(3),
	150: uint16(anon_sym_RPAREN),
	151: uint16(anon_sym_SLASHdisabled),
	152: uint16(anon_sym_SLASHbreakpoint),
	153: uint16(1),
	154: uint16(47),
	155: uint16(4),
	156: uint16(anon_sym_RPAREN),
	157: uint16(anon_sym_COMMA),
	158: uint16(anon_sym_SLASHdisabled),
	159: uint16(anon_sym_SLASHbreakpoint),
	160: uint16(3),
	161: uint16(52),
	162: uint16(1),
	163: uint16(anon_sym_COMMA),
	164: uint16(55),
	165: uint16(1),
	166: uint16(anon_sym_RBRACE),
	167: uint16(14),
	168: uint16(1),
	169: uint16(aux_sym_object_repeat1),
	170: uint16(2),
	171: uint16(59),
	172: uint16(1),
	173: uint16(sym_name),
	174: uint16(57),
	175: uint16(2),
	177: uint16(aux_sym_recipe_token1),
	178: uint16(3),
	179: uint16(61),
	180: uint16(1),
	181: uint16(anon_sym_COMMA),
	182: uint16(63),
	183: uint16(1),
	184: uint16(anon_sym_RBRACE),
	185: uint16(18),
	186: uint16(1),
	187: uint16(aux_sym_object_repeat1),
	188: uint16(3),
	189: uint16(17),
	190: uint16(1),
	191: uint16(anon_sym_SLASHbreakpoint),
	192: uint16(65),
	193: uint16(1),
	194: uint16(anon_sym_RPAREN),
	195: uint16(37),
	196: uint16(1),
	197: uint16(sym__breakpoint),
	198: uint16(3),
	199: uint16(61),
	200: uint16(1),
	201: uint16(anon_sym_COMMA),
	202: uint16(67),
	203: uint16(1),
	204: uint16(anon_sym_RBRACE),
	205: uint16(14),
	206: uint16(1),
	207: uint16(aux_sym_object_repeat1),
	208: uint16(2),
	209: uint16(71),
	210: uint16(1),
	211: uint16(sym_name),
	212: uint16(69),
	213: uint16(2),
	215: uint16(aux_sym_recipe_token1),
	216: uint16(2),
	217: uint16(75),
	218: uint16(1),
	219: uint16(sym_name),
	220: uint16(73),
	221: uint16(2),
	223: uint16(aux_sym_recipe_token1),
	224: uint16(2),
	225: uint16(79),
	226: uint16(1),
	227: uint16(sym_name),
	228: uint16(77),
	229: uint16(2),
	231: uint16(aux_sym_recipe_token1),
	232: uint16(3),
	233: uint16(17),
	234: uint16(1),
	235: uint16(anon_sym_SLASHbreakpoint),
	236: uint16(81),
	237: uint16(1),
	238: uint16(anon_sym_RPAREN),
	239: uint16(45),
	240: uint16(1),
	241: uint16(sym__breakpoint),
	242: uint16(2),
	243: uint16(85),
	244: uint16(1),
	245: uint16(sym_name),
	246: uint16(83),
	247: uint16(2),
	249: uint16(aux_sym_recipe_token1),
	250: uint16(2),
	251: uint16(89),
	252: uint16(1),
	253: uint16(sym_name),
	254: uint16(87),
	255: uint16(2),
	257: uint16(aux_sym_recipe_token1),
	258: uint16(2),
	259: uint16(93),
	260: uint16(1),
	261: uint16(sym_name),
	262: uint16(91),
	263: uint16(2),
	265: uint16(aux_sym_recipe_token1),
	266: uint16(2),
	267: uint16(97),
	268: uint16(1),
	269: uint16(sym_name),
	270: uint16(95),
	271: uint16(2),
	273: uint16(aux_sym_recipe_token1),
	274: uint16(1),
	275: uint16(99),
	276: uint16(2),
	277: uint16(anon_sym_COMMA),
	278: uint16(anon_sym_RBRACE),
	279: uint16(2),
	280: uint16(101),
	281: uint16(1),
	282: uint16(sym_string),
	283: uint16(16),
	284: uint16(1),
	285: uint16(sym__pair),
	286: uint16(1),
	287: uint16(103),
	288: uint16(2),
	289: uint16(anon_sym_COMMA),
	290: uint16(anon_sym_RBRACE),
	291: uint16(1),
	292: uint16(105),
	293: uint16(2),
	294: uint16(anon_sym_RPAREN),
	295: uint16(anon_sym_SLASHbreakpoint),
	296: uint16(2),
	297: uint16(101),
	298: uint16(1),
	299: uint16(sym_string),
	300: uint16(27),
	301: uint16(1),
	302: uint16(sym__pair),
	303: uint16(1),
	304: uint16(107),
	305: uint16(1),
	306: uint16(anon_sym_COLON),
	307: uint16(1),
	308: uint16(109),
	309: uint16(1),
	310: uint16(anon_sym_RPAREN),
	311: uint16(1),
	312: uint16(111),
	313: uint16(1),
	314: uint16(aux_sym_recipe_token1),
	315: uint16(1),
	316: uint16(113),
	317: uint16(1),
	318: uint16(anon_sym_RPAREN),
	319: uint16(1),
	320: uint16(115),
	321: uint16(1),
	322: uint16(aux_sym_recipe_token1),
	323: uint16(1),
	324: uint16(117),
	325: uint16(1),
	326: uint16(anon_sym_RPAREN),
	327: uint16(1),
	328: uint16(119),
	329: uint16(1),
	330: uint16(aux_sym_recipe_token1),
	331: uint16(1),
	332: uint16(121),
	333: uint16(1),
	335: uint16(1),
	336: uint16(123),
	337: uint16(1),
	338: uint16(anon_sym_LPAREN),
	339: uint16(1),
	340: uint16(125),
	341: uint16(1),
	342: uint16(aux_sym_recipe_token1),
	343: uint16(1),
	344: uint16(127),
	345: uint16(1),
	346: uint16(aux_sym_recipe_token1),
	347: uint16(1),
	348: uint16(129),
	349: uint16(1),
	350: uint16(aux_sym_recipe_token1),
	351: uint16(1),
	352: uint16(131),
	353: uint16(1),
	354: uint16(anon_sym_RPAREN),
	355: uint16(1),
	356: uint16(133),
	357: uint16(1),
	358: uint16(anon_sym_RPAREN),
	359: uint16(1),
	360: uint16(135),
	361: uint16(1),
	362: uint16(aux_sym_recipe_token1),
	363: uint16(1),
	364: uint16(137),
	365: uint16(1),
	366: uint16(aux_sym_recipe_token1),
}

var ts_small_parse_table_map = [46]uint32_t{
	1:  uint32(31),
	2:  uint32(44),
	3:  uint32(57),
	4:  uint32(71),
	5:  uint32(83),
	6:  uint32(97),
	7:  uint32(113),
	8:  uint32(125),
	9:  uint32(133),
	10: uint32(141),
	11: uint32(153),
	12: uint32(160),
	13: uint32(170),
	14: uint32(178),
	15: uint32(188),
	16: uint32(198),
	17: uint32(208),
	18: uint32(216),
	19: uint32(224),
	20: uint32(232),
	21: uint32(242),
	22: uint32(250),
	23: uint32(258),
	24: uint32(266),
	25: uint32(274),
	26: uint32(279),
	27: uint32(286),
	28: uint32(291),
	29: uint32(296),
	30: uint32(303),
	31: uint32(307),
	32: uint32(311),
	33: uint32(315),
	34: uint32(319),
	35: uint32(323),
	36: uint32(327),
	37: uint32(331),
	38: uint32(335),
	39: uint32(339),
	40: uint32(343),
	41: uint32(347),
	42: uint32(351),
	43: uint32(355),
	44: uint32(359),
	45: uint32(363),
}

var ts_parse_actions = [139]TSParseActionEntry{
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
		Fsymbol:     uint16(sym_recipe),
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
		Fstate: uint16(libc.Int32FromInt32(40)),
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
		Fstate: uint16(libc.Int32FromInt32(41)),
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
		Fstate: uint16(libc.Int32FromInt32(28)),
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
		Fstate: uint16(libc.Int32FromInt32(30)),
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
		Fstate: uint16(libc.Int32FromInt32(33)),
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
		Fstate: uint16(libc.Int32FromInt32(13)),
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
		Fstate: uint16(libc.Int32FromInt32(29)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_recipe_repeat1),
	})))),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_recipe_repeat1),
	})))),
	27: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(libc.Int32FromInt32(5)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	28: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	29: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_recipe_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(40)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_arguments),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_recipe),
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
		Fstate: uint16(libc.Int32FromInt32(5)),
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
		Fstate: uint16(libc.Int32FromInt32(34)),
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
		Fsymbol:      uint16(sym_arguments),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_object),
		Fproduction_id: uint16(3),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_object),
		Fproduction_id: uint16(7),
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
		Fsymbol:      uint16(aux_sym_arguments_repeat1),
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
		Fsymbol:      uint16(aux_sym_arguments_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(3)),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_object_repeat1),
		Fproduction_id: uint16(8),
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
		Fstate:      uint16(libc.Int32FromInt32(31)),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_object_repeat1),
		Fproduction_id: uint16(8),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	58: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_operation),
		Fproduction_id: uint16(12),
	})))),
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
		Fcount: uint8(1),
	}})),
	60: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_operation),
		Fproduction_id: uint16(12),
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
		Fstate: uint16(libc.Int32FromInt32(31)),
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
		Fstate: uint16(libc.Int32FromInt32(10)),
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
		Fstate: uint16(libc.Int32FromInt32(36)),
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
		Fstate: uint16(libc.Int32FromInt32(11)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_operation),
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
		Fcount: uint8(1),
	}})),
	72: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_operation),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:        uint16(sym_operation),
		Fproduction_id: uint16(4),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_operation),
		Fproduction_id: uint16(4),
	})))),
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
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_operation),
		Fproduction_id: uint16(5),
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
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_operation),
		Fproduction_id: uint16(5),
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
		Fstate: uint16(libc.Int32FromInt32(43)),
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
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_operation),
		Fproduction_id: uint16(9),
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
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_operation),
		Fproduction_id: uint16(9),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_operation),
		Fproduction_id: uint16(10),
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
		Fcount: uint8(1),
	}})),
	90: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_operation),
		Fproduction_id: uint16(10),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	92: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_operation),
		Fproduction_id: uint16(11),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_operation),
		Fproduction_id: uint16(11),
	})))),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_operation),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_operation),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_object_repeat1),
		Fproduction_id: uint16(3),
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
		Fstate: uint16(libc.Int32FromInt32(32)),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym__pair),
		Fproduction_id: uint16(6),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym__disabled),
		Fproduction_id: uint16(1),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym__breakpoint),
		Fproduction_id: uint16(2),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(38)),
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
		Fstate: uint16(libc.Int32FromInt32(23)),
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
		Fstate: uint16(libc.Int32FromInt32(46)),
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
		Fstate: uint16(libc.Int32FromInt32(24)),
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
	122: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f uint8_t
		_ [7]byte
	}{f: uint8(TSParseActionTypeAccept)})),
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
		Fstate: uint16(libc.Int32FromInt32(2)),
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
		Fstate: uint16(libc.Int32FromInt32(26)),
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
		Fstate: uint16(libc.Int32FromInt32(25)),
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
		Fstate: uint16(libc.Int32FromInt32(20)),
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
		Fstate: uint16(libc.Int32FromInt32(47)),
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
		Fstate: uint16(libc.Int32FromInt32(42)),
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
		Fstate: uint16(libc.Int32FromInt32(15)),
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
		Fstate: uint16(libc.Int32FromInt32(21)),
	}})))),
}

func tree_sitter_cyberchef(tls *libc.TLS) (r uintptr) {
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
	Fname:                      __ccgo_ts + 229,
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

var __ccgo_ts1 = "end\x00recipe_token1\x00(\x00)\x00name\x00string\x00number\x00boolean\x00{\x00,\x00}\x00:\x00/disabled\x00/breakpoint\x00recipe\x00operation\x00object\x00arguments\x00_argument\x00_pair\x00_disabled\x00_breakpoint\x00recipe_repeat1\x00object_repeat1\x00arguments_repeat1\x00breakpoint\x00disabled\x00key\x00value\x00cyberchef\x00"
