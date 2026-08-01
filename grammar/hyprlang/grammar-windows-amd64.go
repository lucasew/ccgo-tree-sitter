// Code generated for windows/amd64 by 'ccgo -ignore-unsupported-alignment -ignore-unsupported-atomic-sizes -ignore-vector-functions --winapi-no-errno -ignore-link-errors -std=gnu11 -O0 -D__extension__= -D__forceinline=static inline -D__attribute__(...)= -D__declspec(x)= -D__cdecl= -D__stdcall= -D__fastcall= -D__thiscall= -D_cdecl= -D__restrict= -D__restrict__= -D__MINGW_EXTENSION= -D_X86INTRIN_H_INCLUDED -D_X86GPRINTRIN_H_INCLUDED -D_IMMINTRIN_H_INCLUDED -D_MMINTRIN_H_INCLUDED -D_XMMINTRIN_H_INCLUDED -D_EMMINTRIN_H_INCLUDED -D_PMMINTRIN_H_INCLUDED -D_MM3DNOW_H_INCLUDED -Dfunc=func_token -Dinterface=interface_token -Dselect=select_token -Dchan=chan_token -Dgo=go_token -Dmap=map_token -Dpackage=package_token -Dtype=type_token -Dvar=var_token -Dimport=import_token -Ddefer=defer_token -Dfallthrough=fallthrough_token -Drange=range_token -I D:\a\ccgo-tree-sitter\ccgo-tree-sitter\third-party\tree-sitter-hyprlang\src -I D:\a\ccgo-tree-sitter\ccgo-tree-sitter\third-party\tree-sitter-hyprlang -I D:\a\ccgo-tree-sitter\ccgo-tree-sitter\third-party -I C:\Users\runneradmin\.cache\workspaced\sources\github\7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2\lib\include -I C:\Users\runneradmin\.cache\workspaced\sources\github\7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2\lib\src D:\a\ccgo-tree-sitter\ccgo-tree-sitter\third-party\tree-sitter-hyprlang\src\parser.c -o grammar.go', DO NOT EDIT.

//go:build windows && amd64

package grammar_hyprlang

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
const MAX_ALIAS_SEQUENCE_LENGTH = 8
const MAX_RESERVED_WORD_SET_SIZE = 0
const MB_LEN_MAX = 5
const MINGW_HAS_DDK_H = 1
const MINGW_HAS_SECURE_API = 1
const PATH_MAX = 260
const PRODUCTION_ID_COUNT = 6
const PTRDIFF_MAX = "INT64_MAX"
const PTRDIFF_MIN = "INT64_MIN"
const RAND_MAX = 0x7fff
const SIG_ATOMIC_MAX = "INT32_MAX"
const SIG_ATOMIC_MIN = "INT32_MIN"
const SIZE_MAX = "UINT64_MAX"
const SSIZE_MAX = "_I64_MAX"
const STATE_COUNT = 137
const STRUNCATE = 80
const SUPERTYPE_COUNT = 0
const SYMBOL_COUNT = 94
const TOKEN_COUNT = 58
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

const sym_string = 1
const anon_sym_EQ = 2
const anon_sym_COLON = 3
const anon_sym_LBRACE = 4
const anon_sym_RBRACE = 5
const anon_sym_source = 6
const anon_sym_LBRACK = 7
const anon_sym_SEMI = 8
const anon_sym_RBRACK = 9
const anon_sym_exec_DASHonce = 10
const anon_sym_exec = 11
const anon_sym_execr_DASHonce = 12
const anon_sym_execr = 13
const anon_sym_exec_DASHshutdown = 14
const anon_sym_true = 15
const anon_sym_false = 16
const anon_sym_on = 17
const anon_sym_off = 18
const anon_sym_yes = 19
const anon_sym_no = 20
const anon_sym_PLUS = 21
const anon_sym_DASH = 22
const aux_sym_number_token1 = 23
const anon_sym_x = 24
const anon_sym_rgb = 25
const anon_sym_rgba = 26
const anon_sym_LPAREN = 27
const anon_sym_RPAREN = 28
const anon_sym_COMMA = 29
const anon_sym_AT = 30
const sym_hex = 31
const aux_sym_angle_token1 = 32
const anon_sym_deg = 33
const anon_sym_SHIFT = 34
const anon_sym_CAPS = 35
const anon_sym_CTRL = 36
const anon_sym_CONTROL = 37
const anon_sym_ALT = 38
const anon_sym_ALT_L = 39
const anon_sym_MOD2 = 40
const anon_sym_MOD3 = 41
const anon_sym_SUPER = 42
const anon_sym_WIN = 43
const anon_sym_LOGO = 44
const anon_sym_MOD4 = 45
const anon_sym_MOD5 = 46
const anon_sym_TAB = 47
const sym_string_literal = 48
const sym_name = 49
const sym_device_name = 50
const anon_sym_DOLLAR = 51
const aux_sym_variable_token1 = 52
const anon_sym_0 = 53
const sym__window_rule_argument = 54
const anon_sym_LF = 55
const anon_sym_POUND = 56
const aux_sym_comment_token1 = 57
const sym_configuration = 58
const sym_declaration = 59
const sym_assignment = 60
const sym_keyword = 61
const sym_section = 62
const sym_source = 63
const sym_arguments = 64
const sym_window_rule = 65
const sym_rules = 66
const sym_exec = 67
const sym__value = 68
const sym_boolean = 69
const sym_number = 70
const sym_vec2 = 71
const sym_color = 72
const sym_legacy_hex = 73
const sym_rgb = 74
const sym_gradient = 75
const sym_number_tuple = 76
const sym_display = 77
const sym_position = 78
const sym_angle = 79
const sym_mod = 80
const sym_keys = 81
const sym_params = 82
const sym_variable = 83
const sym__zero = 84
const sym__linebreak = 85
const sym_comment = 86
const aux_sym_configuration_repeat1 = 87
const aux_sym_section_repeat1 = 88
const aux_sym_arguments_repeat1 = 89
const aux_sym_rules_repeat1 = 90
const aux_sym_gradient_repeat1 = 91
const aux_sym_number_tuple_repeat1 = 92
const aux_sym_params_repeat1 = 93

var ts_symbol_names = [94]uintptr{
	0:  __ccgo_ts,
	1:  __ccgo_ts + 4,
	2:  __ccgo_ts + 11,
	3:  __ccgo_ts + 13,
	4:  __ccgo_ts + 15,
	5:  __ccgo_ts + 17,
	6:  __ccgo_ts + 19,
	7:  __ccgo_ts + 26,
	8:  __ccgo_ts + 28,
	9:  __ccgo_ts + 30,
	10: __ccgo_ts + 32,
	11: __ccgo_ts + 42,
	12: __ccgo_ts + 47,
	13: __ccgo_ts + 58,
	14: __ccgo_ts + 64,
	15: __ccgo_ts + 78,
	16: __ccgo_ts + 83,
	17: __ccgo_ts + 89,
	18: __ccgo_ts + 92,
	19: __ccgo_ts + 96,
	20: __ccgo_ts + 100,
	21: __ccgo_ts + 103,
	22: __ccgo_ts + 105,
	23: __ccgo_ts + 107,
	24: __ccgo_ts + 121,
	25: __ccgo_ts + 123,
	26: __ccgo_ts + 127,
	27: __ccgo_ts + 132,
	28: __ccgo_ts + 134,
	29: __ccgo_ts + 136,
	30: __ccgo_ts + 138,
	31: __ccgo_ts + 140,
	32: __ccgo_ts + 144,
	33: __ccgo_ts + 157,
	34: __ccgo_ts + 161,
	35: __ccgo_ts + 167,
	36: __ccgo_ts + 172,
	37: __ccgo_ts + 177,
	38: __ccgo_ts + 185,
	39: __ccgo_ts + 189,
	40: __ccgo_ts + 195,
	41: __ccgo_ts + 200,
	42: __ccgo_ts + 205,
	43: __ccgo_ts + 211,
	44: __ccgo_ts + 215,
	45: __ccgo_ts + 220,
	46: __ccgo_ts + 225,
	47: __ccgo_ts + 230,
	48: __ccgo_ts + 234,
	49: __ccgo_ts + 249,
	50: __ccgo_ts + 254,
	51: __ccgo_ts + 266,
	52: __ccgo_ts + 268,
	53: __ccgo_ts + 284,
	54: __ccgo_ts + 4,
	55: __ccgo_ts + 286,
	56: __ccgo_ts + 288,
	57: __ccgo_ts + 290,
	58: __ccgo_ts + 305,
	59: __ccgo_ts + 319,
	60: __ccgo_ts + 331,
	61: __ccgo_ts + 342,
	62: __ccgo_ts + 350,
	63: __ccgo_ts + 19,
	64: __ccgo_ts + 358,
	65: __ccgo_ts + 368,
	66: __ccgo_ts + 380,
	67: __ccgo_ts + 42,
	68: __ccgo_ts + 386,
	69: __ccgo_ts + 393,
	70: __ccgo_ts + 401,
	71: __ccgo_ts + 408,
	72: __ccgo_ts + 413,
	73: __ccgo_ts + 419,
	74: __ccgo_ts + 123,
	75: __ccgo_ts + 430,
	76: __ccgo_ts + 439,
	77: __ccgo_ts + 452,
	78: __ccgo_ts + 460,
	79: __ccgo_ts + 469,
	80: __ccgo_ts + 475,
	81: __ccgo_ts + 479,
	82: __ccgo_ts + 484,
	83: __ccgo_ts + 491,
	84: __ccgo_ts + 500,
	85: __ccgo_ts + 506,
	86: __ccgo_ts + 517,
	87: __ccgo_ts + 525,
	88: __ccgo_ts + 547,
	89: __ccgo_ts + 563,
	90: __ccgo_ts + 581,
	91: __ccgo_ts + 595,
	92: __ccgo_ts + 612,
	93: __ccgo_ts + 633,
}

var ts_symbol_map = [94]TSSymbol{
	1:  uint16(sym_string),
	2:  uint16(anon_sym_EQ),
	3:  uint16(anon_sym_COLON),
	4:  uint16(anon_sym_LBRACE),
	5:  uint16(anon_sym_RBRACE),
	6:  uint16(anon_sym_source),
	7:  uint16(anon_sym_LBRACK),
	8:  uint16(anon_sym_SEMI),
	9:  uint16(anon_sym_RBRACK),
	10: uint16(anon_sym_exec_DASHonce),
	11: uint16(anon_sym_exec),
	12: uint16(anon_sym_execr_DASHonce),
	13: uint16(anon_sym_execr),
	14: uint16(anon_sym_exec_DASHshutdown),
	15: uint16(anon_sym_true),
	16: uint16(anon_sym_false),
	17: uint16(anon_sym_on),
	18: uint16(anon_sym_off),
	19: uint16(anon_sym_yes),
	20: uint16(anon_sym_no),
	21: uint16(anon_sym_PLUS),
	22: uint16(anon_sym_DASH),
	23: uint16(aux_sym_number_token1),
	24: uint16(anon_sym_x),
	25: uint16(anon_sym_rgb),
	26: uint16(anon_sym_rgba),
	27: uint16(anon_sym_LPAREN),
	28: uint16(anon_sym_RPAREN),
	29: uint16(anon_sym_COMMA),
	30: uint16(anon_sym_AT),
	31: uint16(sym_hex),
	32: uint16(aux_sym_angle_token1),
	33: uint16(anon_sym_deg),
	34: uint16(anon_sym_SHIFT),
	35: uint16(anon_sym_CAPS),
	36: uint16(anon_sym_CTRL),
	37: uint16(anon_sym_CONTROL),
	38: uint16(anon_sym_ALT),
	39: uint16(anon_sym_ALT_L),
	40: uint16(anon_sym_MOD2),
	41: uint16(anon_sym_MOD3),
	42: uint16(anon_sym_SUPER),
	43: uint16(anon_sym_WIN),
	44: uint16(anon_sym_LOGO),
	45: uint16(anon_sym_MOD4),
	46: uint16(anon_sym_MOD5),
	47: uint16(anon_sym_TAB),
	48: uint16(sym_string_literal),
	49: uint16(sym_name),
	50: uint16(sym_device_name),
	51: uint16(anon_sym_DOLLAR),
	52: uint16(aux_sym_variable_token1),
	53: uint16(anon_sym_0),
	54: uint16(sym_string),
	55: uint16(anon_sym_LF),
	56: uint16(anon_sym_POUND),
	57: uint16(aux_sym_comment_token1),
	58: uint16(sym_configuration),
	59: uint16(sym_declaration),
	60: uint16(sym_assignment),
	61: uint16(sym_keyword),
	62: uint16(sym_section),
	63: uint16(sym_source),
	64: uint16(sym_arguments),
	65: uint16(sym_window_rule),
	66: uint16(sym_rules),
	67: uint16(sym_exec),
	68: uint16(sym__value),
	69: uint16(sym_boolean),
	70: uint16(sym_number),
	71: uint16(sym_vec2),
	72: uint16(sym_color),
	73: uint16(sym_legacy_hex),
	74: uint16(sym_rgb),
	75: uint16(sym_gradient),
	76: uint16(sym_number_tuple),
	77: uint16(sym_display),
	78: uint16(sym_position),
	79: uint16(sym_angle),
	80: uint16(sym_mod),
	81: uint16(sym_keys),
	82: uint16(sym_params),
	83: uint16(sym_variable),
	84: uint16(sym__zero),
	85: uint16(sym__linebreak),
	86: uint16(sym_comment),
	87: uint16(aux_sym_configuration_repeat1),
	88: uint16(aux_sym_section_repeat1),
	89: uint16(aux_sym_arguments_repeat1),
	90: uint16(aux_sym_rules_repeat1),
	91: uint16(aux_sym_gradient_repeat1),
	92: uint16(aux_sym_number_tuple_repeat1),
	93: uint16(aux_sym_params_repeat1),
}

var ts_symbol_metadata = [94]TSSymbolMetadata{
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
	23: {},
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
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	32: {},
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
	},
	52: {},
	53: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	54: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	55: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	56: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	57: {},
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
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	85: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	86: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	87: {},
	88: {},
	89: {},
	90: {},
	91: {},
	92: {},
	93: {},
}

type ts_field_identifiers = int32

const field_device = 1
const field_keyword = 2
const field_name = 3
const field_value = 4

var ts_field_names = [5]uintptr{
	0: libc.UintptrFromInt32(0),
	1: __ccgo_ts + 648,
	2: __ccgo_ts + 342,
	3: __ccgo_ts + 249,
	4: __ccgo_ts + 655,
}

var ts_field_map_slices = [6]TSMapSlice{
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
}

var ts_field_map_entries = [8]TSFieldMapEntry{
	0: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
	},
	1: {
		Ffield_id: uint16(field_name),
	},
	2: {
		Ffield_id: uint16(field_name),
	},
	3: {
		Ffield_id:    uint16(field_value),
		Fchild_index: uint8(2),
	},
	4: {
		Ffield_id: uint16(field_keyword),
	},
	5: {
		Ffield_id:    uint16(field_value),
		Fchild_index: uint8(2),
	},
	6: {
		Ffield_id:    uint16(field_device),
		Fchild_index: uint8(2),
	},
	7: {
		Ffield_id: uint16(field_name),
	},
}

var ts_alias_sequences = [6][8]TSSymbol{}

var ts_non_terminal_alias_map = [1]uint16_t{}

var ts_primary_state_ids = [137]TSStateId{
	1:   uint16(1),
	2:   uint16(2),
	3:   uint16(2),
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
	40:  uint16(39),
	41:  uint16(35),
	42:  uint16(37),
	43:  uint16(36),
	44:  uint16(44),
	45:  uint16(45),
	46:  uint16(33),
	47:  uint16(30),
	48:  uint16(31),
	49:  uint16(49),
	50:  uint16(50),
	51:  uint16(51),
	52:  uint16(52),
	53:  uint16(53),
	54:  uint16(54),
	55:  uint16(55),
	56:  uint16(56),
	57:  uint16(56),
	58:  uint16(58),
	59:  uint16(59),
	60:  uint16(60),
	61:  uint16(61),
	62:  uint16(62),
	63:  uint16(63),
	64:  uint16(64),
	65:  uint16(15),
	66:  uint16(66),
	67:  uint16(67),
	68:  uint16(68),
	69:  uint16(20),
	70:  uint16(70),
	71:  uint16(27),
	72:  uint16(28),
	73:  uint16(19),
	74:  uint16(66),
	75:  uint16(24),
	76:  uint16(29),
	77:  uint16(25),
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
	100: uint16(44),
	101: uint16(101),
	102: uint16(102),
	103: uint16(102),
	104: uint16(80),
	105: uint16(89),
	106: uint16(93),
	107: uint16(99),
	108: uint16(108),
	109: uint16(97),
	110: uint16(110),
	111: uint16(111),
	112: uint16(101),
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
	127: uint16(124),
	128: uint16(128),
	129: uint16(129),
	130: uint16(130),
	131: uint16(131),
	132: uint16(132),
	133: uint16(133),
	134: uint16(130),
	135: uint16(118),
	136: uint16(136),
}

func ts_lex(tls *libc.TLS, lexer uintptr, state TSStateId) (r uint8) {
	var eof, result, skip uint8
	var i, i1, i10, i2, i3, i4, i5, i6, i7, i8, i9 uint32_t
	var lookahead int32_t
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = eof, i, i1, i10, i2, i3, i4, i5, i6, i7, i8, i9, lookahead, result, skip
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
			if !(uint64(i) < libc.Uint64FromInt64(124)/libc.Uint64FromInt64(2)) {
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
			state = uint16(85)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(1):
		i1 = uint32(0)
		for {
			if !(uint64(i1) < libc.Uint64FromInt64(68)/libc.Uint64FromInt64(2)) {
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
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(92)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(2):
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
			state = uint16(92)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(3):
		i3 = uint32(0)
		for {
			if !(uint64(i3) < libc.Uint64FromInt64(36)/libc.Uint64FromInt64(2)) {
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
		if lookahead == int32('\t') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(3)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(92)
			goto next_state
		}
		return result
	case int32(4):
		i4 = uint32(0)
		for {
			if !(uint64(i4) < libc.Uint64FromInt64(60)/libc.Uint64FromInt64(2)) {
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
		if lookahead == int32('\t') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(4)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(106)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead == int32('\n') {
			state = uint16(256)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(257)
			goto next_state
		}
		if lookahead == int32('}') {
			state = uint16(63)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(5)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(246)
			goto next_state
		}
		return result
	case int32(6):
		i5 = uint32(0)
		for {
			if !(uint64(i5) < libc.Uint64FromInt64(56)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token5[i5]) == lookahead {
				state = map_token5[i5+uint32(1)]
				goto next_state
			}
			goto _6
		_6:
			;
			i5 = i5 + uint32(2)
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(92)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\t') && lookahead != int32('\n') {
			state = uint16(222)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead == int32('#') {
			state = uint16(257)
			goto next_state
		}
		if lookahead == int32('+') {
			state = uint16(79)
			goto next_state
		}
		if lookahead == int32('-') {
			state = uint16(81)
			goto next_state
		}
		if lookahead == int32('0') {
			state = uint16(251)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(7)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(91)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(57)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead == int32('#') {
			state = uint16(257)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(66)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			state = uint16(127)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\t') && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead == int32('#') {
			state = uint16(257)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			state = uint16(128)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\t') && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(10):
		if lookahead == int32('#') {
			state = uint16(257)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(10)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(249)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead == int32('#') {
			state = uint16(257)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(11)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(92)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead == int32('#') {
			state = uint16(257)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(12)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(57)
			goto next_state
		}
		return result
	case int32(13):
		if lookahead == int32('#') {
			state = uint16(257)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(13)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(247)
			goto next_state
		}
		return result
	case int32(14):
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(14)
			goto next_state
		}
		return result
	case int32(15):
		if lookahead == int32('#') {
			state = uint16(187)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(14)
			goto next_state
		}
		return result
	case int32(16):
		if lookahead == int32('#') {
			state = uint16(258)
			goto next_state
		}
		if lookahead == int32('+') {
			state = uint16(80)
			goto next_state
		}
		if lookahead == int32('-') {
			state = uint16(82)
			goto next_state
		}
		if lookahead == int32('0') {
			state = uint16(254)
			goto next_state
		}
		if lookahead == int32(';') {
			state = uint16(67)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(68)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(16)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(93)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && (lookahead < int32('+') || int32('-') < lookahead) {
			state = uint16(255)
			goto next_state
		}
		return result
	case int32(17):
		if lookahead == int32('#') {
			state = uint16(189)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(18)
			goto next_state
		}
		return result
	case int32(18):
		if lookahead == int32('#') {
			state = uint16(17)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(18)
			goto next_state
		}
		return result
	case int32(19):
		if lookahead == int32('2') {
			state = uint16(116)
			goto next_state
		}
		if lookahead == int32('3') {
			state = uint16(117)
			goto next_state
		}
		if lookahead == int32('4') {
			state = uint16(121)
			goto next_state
		}
		if lookahead == int32('5') {
			state = uint16(122)
			goto next_state
		}
		return result
	case int32(20):
		if lookahead == int32('A') {
			state = uint16(40)
			goto next_state
		}
		if lookahead == int32('O') {
			state = uint16(35)
			goto next_state
		}
		if lookahead == int32('T') {
			state = uint16(43)
			goto next_state
		}
		return result
	case int32(21):
		if lookahead == int32('A') {
			state = uint16(22)
			goto next_state
		}
		return result
	case int32(22):
		if lookahead == int32('B') {
			state = uint16(123)
			goto next_state
		}
		return result
	case int32(23):
		if lookahead == int32('D') {
			state = uint16(19)
			goto next_state
		}
		return result
	case int32(24):
		if lookahead == int32('E') {
			state = uint16(42)
			goto next_state
		}
		return result
	case int32(25):
		if lookahead == int32('F') {
			state = uint16(46)
			goto next_state
		}
		return result
	case int32(26):
		if lookahead == int32('G') {
			state = uint16(36)
			goto next_state
		}
		return result
	case int32(27):
		if lookahead == int32('H') {
			state = uint16(28)
			goto next_state
		}
		if lookahead == int32('U') {
			state = uint16(41)
			goto next_state
		}
		return result
	case int32(28):
		if lookahead == int32('I') {
			state = uint16(25)
			goto next_state
		}
		return result
	case int32(29):
		if lookahead == int32('I') {
			state = uint16(34)
			goto next_state
		}
		return result
	case int32(30):
		if lookahead == int32('L') {
			state = uint16(110)
			goto next_state
		}
		return result
	case int32(31):
		if lookahead == int32('L') {
			state = uint16(115)
			goto next_state
		}
		return result
	case int32(32):
		if lookahead == int32('L') {
			state = uint16(111)
			goto next_state
		}
		return result
	case int32(33):
		if lookahead == int32('L') {
			state = uint16(48)
			goto next_state
		}
		return result
	case int32(34):
		if lookahead == int32('N') {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(35):
		if lookahead == int32('N') {
			state = uint16(47)
			goto next_state
		}
		return result
	case int32(36):
		if lookahead == int32('O') {
			state = uint16(120)
			goto next_state
		}
		return result
	case int32(37):
		if lookahead == int32('O') {
			state = uint16(26)
			goto next_state
		}
		return result
	case int32(38):
		if lookahead == int32('O') {
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(39):
		if lookahead == int32('O') {
			state = uint16(32)
			goto next_state
		}
		return result
	case int32(40):
		if lookahead == int32('P') {
			state = uint16(45)
			goto next_state
		}
		return result
	case int32(41):
		if lookahead == int32('P') {
			state = uint16(24)
			goto next_state
		}
		return result
	case int32(42):
		if lookahead == int32('R') {
			state = uint16(118)
			goto next_state
		}
		return result
	case int32(43):
		if lookahead == int32('R') {
			state = uint16(30)
			goto next_state
		}
		return result
	case int32(44):
		if lookahead == int32('R') {
			state = uint16(39)
			goto next_state
		}
		return result
	case int32(45):
		if lookahead == int32('S') {
			state = uint16(109)
			goto next_state
		}
		return result
	case int32(46):
		if lookahead == int32('T') {
			state = uint16(108)
			goto next_state
		}
		return result
	case int32(47):
		if lookahead == int32('T') {
			state = uint16(44)
			goto next_state
		}
		return result
	case int32(48):
		if lookahead == int32('T') {
			state = uint16(114)
			goto next_state
		}
		return result
	case int32(49):
		if lookahead == int32('b') {
			state = uint16(95)
			goto next_state
		}
		return result
	case int32(50):
		if lookahead == int32('e') {
			state = uint16(51)
			goto next_state
		}
		return result
	case int32(51):
		if lookahead == int32('g') {
			state = uint16(107)
			goto next_state
		}
		return result
	case int32(52):
		if lookahead == int32('g') {
			state = uint16(49)
			goto next_state
		}
		return result
	case int32(53):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(103)
			goto next_state
		}
		return result
	case int32(54):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(53)
			goto next_state
		}
		return result
	case int32(55):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(54)
			goto next_state
		}
		return result
	case int32(56):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(55)
			goto next_state
		}
		return result
	case int32(57):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(56)
			goto next_state
		}
		return result
	case int32(58):
		if eof != 0 {
			state = uint16(59)
			goto next_state
		}
		i6 = uint32(0)
		for {
			if !(uint64(i6) < libc.Uint64FromInt64(56)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token6[i6]) == lookahead {
				state = map_token6[i6+uint32(1)]
				goto next_state
			}
			goto _7
		_7:
			;
			i6 = i6 + uint32(2)
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(58)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(246)
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
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(61):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COLON)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(62):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(63):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(64):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_source)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(65):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_source)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') || lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(246)
			goto next_state
		}
		return result
	case int32(66):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(67):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SEMI)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(68):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(69):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_exec_DASHonce)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(70):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_exec_DASHonce)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') || lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(246)
			goto next_state
		}
		return result
	case int32(71):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_exec)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') {
			state = uint16(177)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(75)
			goto next_state
		}
		return result
	case int32(72):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_exec)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') {
			state = uint16(237)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(76)
			goto next_state
		}
		if lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(246)
			goto next_state
		}
		return result
	case int32(73):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_execr_DASHonce)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(74):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_execr_DASHonce)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') || lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(246)
			goto next_state
		}
		return result
	case int32(75):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_execr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') {
			state = uint16(179)
			goto next_state
		}
		return result
	case int32(76):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_execr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') {
			state = uint16(239)
			goto next_state
		}
		if lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(246)
			goto next_state
		}
		return result
	case int32(77):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_exec_DASHshutdown)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(78):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_exec_DASHshutdown)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') || lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(246)
			goto next_state
		}
		return result
	case int32(79):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PLUS)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(80):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PLUS)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32(',') && lookahead != int32(';') && lookahead != int32(']') {
			state = uint16(255)
			goto next_state
		}
		return result
	case int32(81):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(82):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32(',') && lookahead != int32(';') && lookahead != int32(']') {
			state = uint16(255)
			goto next_state
		}
		return result
	case int32(83):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_number_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(92)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(92)
			goto next_state
		}
		return result
	case int32(84):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_number_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(92)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(92)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(101)
			goto next_state
		}
		return result
	case int32(85):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_number_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(92)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(83)
			goto next_state
		}
		return result
	case int32(86):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_number_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(92)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(87)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(53)
			goto next_state
		}
		return result
	case int32(87):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_number_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(92)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(88)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(103)
			goto next_state
		}
		return result
	case int32(88):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_number_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(92)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(84)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(102)
			goto next_state
		}
		return result
	case int32(89):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_number_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(92)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(86)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(54)
			goto next_state
		}
		return result
	case int32(90):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_number_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(92)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(89)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(55)
			goto next_state
		}
		return result
	case int32(91):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_number_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(92)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(90)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(56)
			goto next_state
		}
		return result
	case int32(92):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_number_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(92)
			goto next_state
		}
		return result
	case int32(93):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_number_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(93)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32(',') && lookahead != int32(';') && lookahead != int32(']') {
			state = uint16(255)
			goto next_state
		}
		return result
	case int32(94):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_x)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(95):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_rgb)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('a') {
			state = uint16(96)
			goto next_state
		}
		return result
	case int32(96):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_rgba)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(97):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(98):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(99):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COMMA)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(100):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(101):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_hex)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(102):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_hex)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(101)
			goto next_state
		}
		return result
	case int32(103):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_hex)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(102)
			goto next_state
		}
		return result
	case int32(104):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_angle_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(105):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_angle_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(106):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_angle_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(105)
			goto next_state
		}
		return result
	case int32(107):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_deg)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(108):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SHIFT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(109):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_CAPS)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(110):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_CTRL)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(111):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_CONTROL)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(112):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ALT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(142)
			goto next_state
		}
		return result
	case int32(113):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ALT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(202)
			goto next_state
		}
		return result
	case int32(114):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ALT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(31)
			goto next_state
		}
		return result
	case int32(115):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ALT_L)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(116):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_MOD2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(117):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_MOD3)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(118):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SUPER)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(119):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_WIN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(120):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LOGO)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(121):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_MOD4)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(122):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_MOD5)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(123):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_TAB)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(124):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i7 = uint32(0)
		for {
			if !(uint64(i7) < libc.Uint64FromInt64(124)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token7[i7]) == lookahead {
				state = map_token7[i7+uint32(1)]
				goto next_state
			}
			goto _8
		_8:
			;
			i7 = i7 + uint32(2)
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(85)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(125):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i8 = uint32(0)
		for {
			if !(uint64(i8) < libc.Uint64FromInt64(68)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token8[i8]) == lookahead {
				state = map_token8[i8+uint32(1)]
				goto next_state
			}
			goto _9
		_9:
			;
			i8 = i8 + uint32(2)
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(92)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(126):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i9 = uint32(0)
		for {
			if !(uint64(i9) < libc.Uint64FromInt64(68)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token9[i9]) == lookahead {
				state = map_token9[i9+uint32(1)]
				goto next_state
			}
			goto _10
		_10:
			;
			i9 = i9 + uint32(2)
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(92)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(127):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(257)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(66)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			state = uint16(127)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\t') && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(128):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(257)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			state = uint16(128)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\t') && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(129):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('2') {
			state = uint16(116)
			goto next_state
		}
		if lookahead == int32('3') {
			state = uint16(117)
			goto next_state
		}
		if lookahead == int32('4') {
			state = uint16(121)
			goto next_state
		}
		if lookahead == int32('5') {
			state = uint16(122)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(130):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('A') {
			state = uint16(150)
			goto next_state
		}
		if lookahead == int32('O') {
			state = uint16(145)
			goto next_state
		}
		if lookahead == int32('T') {
			state = uint16(153)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(131):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('A') {
			state = uint16(132)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(132):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('B') {
			state = uint16(123)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(133):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('D') {
			state = uint16(129)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(134):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('E') {
			state = uint16(152)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(135):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('F') {
			state = uint16(157)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(136):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('G') {
			state = uint16(148)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(137):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('H') {
			state = uint16(138)
			goto next_state
		}
		if lookahead == int32('U') {
			state = uint16(151)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(138):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('I') {
			state = uint16(135)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(139):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('I') {
			state = uint16(144)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(140):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('L') {
			state = uint16(156)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(141):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('L') {
			state = uint16(110)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(142):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('L') {
			state = uint16(115)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(143):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('L') {
			state = uint16(111)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(144):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('N') {
			state = uint16(119)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(145):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('N') {
			state = uint16(158)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(146):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('O') {
			state = uint16(136)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(147):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('O') {
			state = uint16(133)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(148):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('O') {
			state = uint16(120)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(149):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('O') {
			state = uint16(143)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(150):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('P') {
			state = uint16(155)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(151):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('P') {
			state = uint16(134)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(152):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('R') {
			state = uint16(118)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(153):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('R') {
			state = uint16(141)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(154):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('R') {
			state = uint16(149)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(155):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('S') {
			state = uint16(109)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(156):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('T') {
			state = uint16(112)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(157):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('T') {
			state = uint16(108)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(158):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('T') {
			state = uint16(154)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(159):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('b') {
			state = uint16(95)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(160):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('c') {
			state = uint16(71)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(161):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('c') {
			state = uint16(166)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(162):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('c') {
			state = uint16(167)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(163):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('c') {
			state = uint16(168)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(164):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('d') {
			state = uint16(178)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(165):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(160)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(166):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(64)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(167):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(69)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(168):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(73)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(169):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(171)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(170):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('g') {
			state = uint16(159)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(171):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('g') {
			state = uint16(107)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(172):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('h') {
			state = uint16(183)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(173):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(77)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(174):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(162)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(175):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(163)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(176):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(182)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(177):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(174)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(172)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(178):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(184)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(179):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(175)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(180):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(161)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(181):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(164)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(182):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(180)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(183):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(181)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(184):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('w') {
			state = uint16(173)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(185):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('x') {
			state = uint16(165)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(186):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(14)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(187):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(187)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(187)
			goto next_state
		}
		return result
	case int32(188):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i10 = uint32(0)
		for {
			if !(uint64(i10) < libc.Uint64FromInt64(56)/libc.Uint64FromInt64(2)) {
				break
			}
			if int32(map_token10[i10]) == lookahead {
				state = map_token10[i10+uint32(1)]
				goto next_state
			}
			goto _11
		_11:
			;
			i10 = i10 + uint32(2)
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(92)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\t') && lookahead != int32('\n') {
			state = uint16(222)
			goto next_state
		}
		return result
	case int32(189):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(189)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(189)
			goto next_state
		}
		return result
	case int32(190):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(17)
			goto next_state
		}
		if lookahead == int32('2') {
			state = uint16(116)
			goto next_state
		}
		if lookahead == int32('3') {
			state = uint16(117)
			goto next_state
		}
		if lookahead == int32('4') {
			state = uint16(121)
			goto next_state
		}
		if lookahead == int32('5') {
			state = uint16(122)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(222)
			goto next_state
		}
		return result
	case int32(191):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(17)
			goto next_state
		}
		if lookahead == int32('A') {
			state = uint16(211)
			goto next_state
		}
		if lookahead == int32('O') {
			state = uint16(206)
			goto next_state
		}
		if lookahead == int32('T') {
			state = uint16(214)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(222)
			goto next_state
		}
		return result
	case int32(192):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(17)
			goto next_state
		}
		if lookahead == int32('A') {
			state = uint16(193)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(222)
			goto next_state
		}
		return result
	case int32(193):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(17)
			goto next_state
		}
		if lookahead == int32('B') {
			state = uint16(123)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(222)
			goto next_state
		}
		return result
	case int32(194):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(17)
			goto next_state
		}
		if lookahead == int32('D') {
			state = uint16(190)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(222)
			goto next_state
		}
		return result
	case int32(195):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(17)
			goto next_state
		}
		if lookahead == int32('E') {
			state = uint16(213)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(222)
			goto next_state
		}
		return result
	case int32(196):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(17)
			goto next_state
		}
		if lookahead == int32('F') {
			state = uint16(217)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(222)
			goto next_state
		}
		return result
	case int32(197):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(17)
			goto next_state
		}
		if lookahead == int32('G') {
			state = uint16(207)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(222)
			goto next_state
		}
		return result
	case int32(198):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(17)
			goto next_state
		}
		if lookahead == int32('H') {
			state = uint16(199)
			goto next_state
		}
		if lookahead == int32('U') {
			state = uint16(212)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(222)
			goto next_state
		}
		return result
	case int32(199):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(17)
			goto next_state
		}
		if lookahead == int32('I') {
			state = uint16(196)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(222)
			goto next_state
		}
		return result
	case int32(200):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(17)
			goto next_state
		}
		if lookahead == int32('I') {
			state = uint16(205)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(222)
			goto next_state
		}
		return result
	case int32(201):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(17)
			goto next_state
		}
		if lookahead == int32('L') {
			state = uint16(110)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(222)
			goto next_state
		}
		return result
	case int32(202):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(17)
			goto next_state
		}
		if lookahead == int32('L') {
			state = uint16(115)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(222)
			goto next_state
		}
		return result
	case int32(203):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(17)
			goto next_state
		}
		if lookahead == int32('L') {
			state = uint16(111)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(222)
			goto next_state
		}
		return result
	case int32(204):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(17)
			goto next_state
		}
		if lookahead == int32('L') {
			state = uint16(218)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(222)
			goto next_state
		}
		return result
	case int32(205):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(17)
			goto next_state
		}
		if lookahead == int32('N') {
			state = uint16(119)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(222)
			goto next_state
		}
		return result
	case int32(206):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(17)
			goto next_state
		}
		if lookahead == int32('N') {
			state = uint16(219)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(222)
			goto next_state
		}
		return result
	case int32(207):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(17)
			goto next_state
		}
		if lookahead == int32('O') {
			state = uint16(120)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(222)
			goto next_state
		}
		return result
	case int32(208):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(17)
			goto next_state
		}
		if lookahead == int32('O') {
			state = uint16(197)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(222)
			goto next_state
		}
		return result
	case int32(209):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(17)
			goto next_state
		}
		if lookahead == int32('O') {
			state = uint16(194)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(222)
			goto next_state
		}
		return result
	case int32(210):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(17)
			goto next_state
		}
		if lookahead == int32('O') {
			state = uint16(203)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(222)
			goto next_state
		}
		return result
	case int32(211):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(17)
			goto next_state
		}
		if lookahead == int32('P') {
			state = uint16(216)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(222)
			goto next_state
		}
		return result
	case int32(212):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(17)
			goto next_state
		}
		if lookahead == int32('P') {
			state = uint16(195)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(222)
			goto next_state
		}
		return result
	case int32(213):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(17)
			goto next_state
		}
		if lookahead == int32('R') {
			state = uint16(118)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(222)
			goto next_state
		}
		return result
	case int32(214):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(17)
			goto next_state
		}
		if lookahead == int32('R') {
			state = uint16(201)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(222)
			goto next_state
		}
		return result
	case int32(215):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(17)
			goto next_state
		}
		if lookahead == int32('R') {
			state = uint16(210)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(222)
			goto next_state
		}
		return result
	case int32(216):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(17)
			goto next_state
		}
		if lookahead == int32('S') {
			state = uint16(109)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(222)
			goto next_state
		}
		return result
	case int32(217):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(17)
			goto next_state
		}
		if lookahead == int32('T') {
			state = uint16(108)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(222)
			goto next_state
		}
		return result
	case int32(218):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(17)
			goto next_state
		}
		if lookahead == int32('T') {
			state = uint16(113)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(222)
			goto next_state
		}
		return result
	case int32(219):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(17)
			goto next_state
		}
		if lookahead == int32('T') {
			state = uint16(215)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(222)
			goto next_state
		}
		return result
	case int32(220):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(17)
			goto next_state
		}
		if lookahead == int32('b') {
			state = uint16(95)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(222)
			goto next_state
		}
		return result
	case int32(221):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(17)
			goto next_state
		}
		if lookahead == int32('g') {
			state = uint16(220)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(222)
			goto next_state
		}
		return result
	case int32(222):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(17)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(222)
			goto next_state
		}
		return result
	case int32(223):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('c') {
			state = uint16(72)
			goto next_state
		}
		if lookahead == int32('-') || lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(246)
			goto next_state
		}
		return result
	case int32(224):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('c') {
			state = uint16(229)
			goto next_state
		}
		if lookahead == int32('-') || lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(246)
			goto next_state
		}
		return result
	case int32(225):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('c') {
			state = uint16(230)
			goto next_state
		}
		if lookahead == int32('-') || lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(246)
			goto next_state
		}
		return result
	case int32(226):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('c') {
			state = uint16(231)
			goto next_state
		}
		if lookahead == int32('-') || lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(246)
			goto next_state
		}
		return result
	case int32(227):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('d') {
			state = uint16(238)
			goto next_state
		}
		if lookahead == int32('-') || lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(246)
			goto next_state
		}
		return result
	case int32(228):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(223)
			goto next_state
		}
		if lookahead == int32('-') || lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(246)
			goto next_state
		}
		return result
	case int32(229):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(65)
			goto next_state
		}
		if lookahead == int32('-') || lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(246)
			goto next_state
		}
		return result
	case int32(230):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(70)
			goto next_state
		}
		if lookahead == int32('-') || lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(246)
			goto next_state
		}
		return result
	case int32(231):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(74)
			goto next_state
		}
		if lookahead == int32('-') || lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(246)
			goto next_state
		}
		return result
	case int32(232):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('h') {
			state = uint16(243)
			goto next_state
		}
		if lookahead == int32('-') || lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(246)
			goto next_state
		}
		return result
	case int32(233):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('n') {
			state = uint16(78)
			goto next_state
		}
		if lookahead == int32('-') || lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(246)
			goto next_state
		}
		return result
	case int32(234):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('n') {
			state = uint16(225)
			goto next_state
		}
		if lookahead == int32('-') || lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(246)
			goto next_state
		}
		return result
	case int32(235):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('n') {
			state = uint16(226)
			goto next_state
		}
		if lookahead == int32('-') || lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(246)
			goto next_state
		}
		return result
	case int32(236):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('o') {
			state = uint16(242)
			goto next_state
		}
		if lookahead == int32('-') || lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(246)
			goto next_state
		}
		return result
	case int32(237):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('o') {
			state = uint16(234)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(232)
			goto next_state
		}
		if lookahead == int32('-') || lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(246)
			goto next_state
		}
		return result
	case int32(238):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('o') {
			state = uint16(244)
			goto next_state
		}
		if lookahead == int32('-') || lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(246)
			goto next_state
		}
		return result
	case int32(239):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('o') {
			state = uint16(235)
			goto next_state
		}
		if lookahead == int32('-') || lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(246)
			goto next_state
		}
		return result
	case int32(240):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(224)
			goto next_state
		}
		if lookahead == int32('-') || lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(246)
			goto next_state
		}
		return result
	case int32(241):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('t') {
			state = uint16(227)
			goto next_state
		}
		if lookahead == int32('-') || lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(246)
			goto next_state
		}
		return result
	case int32(242):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('u') {
			state = uint16(240)
			goto next_state
		}
		if lookahead == int32('-') || lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(246)
			goto next_state
		}
		return result
	case int32(243):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('u') {
			state = uint16(241)
			goto next_state
		}
		if lookahead == int32('-') || lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(246)
			goto next_state
		}
		return result
	case int32(244):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('w') {
			state = uint16(233)
			goto next_state
		}
		if lookahead == int32('-') || lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(246)
			goto next_state
		}
		return result
	case int32(245):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('x') {
			state = uint16(228)
			goto next_state
		}
		if lookahead == int32('-') || lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(246)
			goto next_state
		}
		return result
	case int32(246):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') || lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(246)
			goto next_state
		}
		return result
	case int32(247):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_device_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('-') <= lookahead && lookahead <= int32(':') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(247)
			goto next_state
		}
		return result
	case int32(248):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOLLAR)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(249):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_variable_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(249)
			goto next_state
		}
		return result
	case int32(250):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_0)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(92)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(83)
			goto next_state
		}
		return result
	case int32(251):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_0)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(92)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(90)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(56)
			goto next_state
		}
		return result
	case int32(252):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_0)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(105)
			goto next_state
		}
		return result
	case int32(253):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_0)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(92)
			goto next_state
		}
		return result
	case int32(254):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_0)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(93)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32(',') && lookahead != int32(';') && lookahead != int32(']') {
			state = uint16(255)
			goto next_state
		}
		return result
	case int32(255):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__window_rule_argument)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32(',') && lookahead != int32(';') && lookahead != int32(']') {
			state = uint16(255)
			goto next_state
		}
		return result
	case int32(256):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LF)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(257):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUND)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(258):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUND)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32(',') && lookahead != int32(';') && lookahead != int32(']') {
			state = uint16(255)
			goto next_state
		}
		return result
	case int32(259):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUND)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(260):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_comment_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(259)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			state = uint16(260)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\t') && lookahead != int32('\n') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(261):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_comment_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(261)
			goto next_state
		}
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var map_token = [62]uint16_t{
	0:  uint16('\n'),
	1:  uint16(256),
	2:  uint16('#'),
	3:  uint16(257),
	4:  uint16('$'),
	5:  uint16(248),
	6:  uint16('('),
	7:  uint16(97),
	8:  uint16(')'),
	9:  uint16(98),
	10: uint16('+'),
	11: uint16(79),
	12: uint16(','),
	13: uint16(99),
	14: uint16('-'),
	15: uint16(81),
	16: uint16('0'),
	17: uint16(250),
	18: uint16(':'),
	19: uint16(61),
	20: uint16(';'),
	21: uint16(67),
	22: uint16('='),
	23: uint16(60),
	24: uint16('@'),
	25: uint16(100),
	26: uint16('A'),
	27: uint16(140),
	28: uint16('C'),
	29: uint16(130),
	30: uint16('L'),
	31: uint16(146),
	32: uint16('M'),
	33: uint16(147),
	34: uint16('S'),
	35: uint16(137),
	36: uint16('T'),
	37: uint16(131),
	38: uint16('W'),
	39: uint16(139),
	40: uint16('['),
	41: uint16(66),
	42: uint16(']'),
	43: uint16(68),
	44: uint16('d'),
	45: uint16(169),
	46: uint16('e'),
	47: uint16(185),
	48: uint16('r'),
	49: uint16(170),
	50: uint16('s'),
	51: uint16(176),
	52: uint16('x'),
	53: uint16(94),
	54: uint16('{'),
	55: uint16(62),
	56: uint16('}'),
	57: uint16(63),
	58: uint16('\t'),
	59: uint16(124),
	60: uint16(' '),
	61: uint16(124),
}

var map_token1 = [34]uint16_t{
	0:  uint16('\n'),
	1:  uint16(256),
	2:  uint16('#'),
	3:  uint16(257),
	4:  uint16('$'),
	5:  uint16(248),
	6:  uint16('+'),
	7:  uint16(79),
	8:  uint16(','),
	9:  uint16(99),
	10: uint16('-'),
	11: uint16(81),
	12: uint16('0'),
	13: uint16(253),
	14: uint16('A'),
	15: uint16(140),
	16: uint16('C'),
	17: uint16(130),
	18: uint16('L'),
	19: uint16(146),
	20: uint16('M'),
	21: uint16(147),
	22: uint16('S'),
	23: uint16(137),
	24: uint16('T'),
	25: uint16(131),
	26: uint16('W'),
	27: uint16(139),
	28: uint16('r'),
	29: uint16(170),
	30: uint16('\t'),
	31: uint16(125),
	32: uint16(' '),
	33: uint16(125),
}

var map_token2 = [34]uint16_t{
	0:  uint16('\n'),
	1:  uint16(256),
	2:  uint16('#'),
	3:  uint16(257),
	4:  uint16('$'),
	5:  uint16(248),
	6:  uint16('+'),
	7:  uint16(79),
	8:  uint16(','),
	9:  uint16(14),
	10: uint16('-'),
	11: uint16(81),
	12: uint16('0'),
	13: uint16(253),
	14: uint16('A'),
	15: uint16(140),
	16: uint16('C'),
	17: uint16(130),
	18: uint16('L'),
	19: uint16(146),
	20: uint16('M'),
	21: uint16(147),
	22: uint16('S'),
	23: uint16(137),
	24: uint16('T'),
	25: uint16(131),
	26: uint16('W'),
	27: uint16(139),
	28: uint16('r'),
	29: uint16(170),
	30: uint16('\t'),
	31: uint16(126),
	32: uint16(' '),
	33: uint16(126),
}

var map_token3 = [18]uint16_t{
	0:  uint16('\n'),
	1:  uint16(256),
	2:  uint16('#'),
	3:  uint16(257),
	4:  uint16(')'),
	5:  uint16(98),
	6:  uint16('+'),
	7:  uint16(79),
	8:  uint16(','),
	9:  uint16(99),
	10: uint16('-'),
	11: uint16(81),
	12: uint16('0'),
	13: uint16(253),
	14: uint16('@'),
	15: uint16(100),
	16: uint16('x'),
	17: uint16(94),
}

var map_token4 = [30]uint16_t{
	0:  uint16('\n'),
	1:  uint16(256),
	2:  uint16('#'),
	3:  uint16(257),
	4:  uint16(','),
	5:  uint16(99),
	6:  uint16('0'),
	7:  uint16(252),
	8:  uint16('='),
	9:  uint16(60),
	10: uint16('A'),
	11: uint16(33),
	12: uint16('C'),
	13: uint16(20),
	14: uint16('L'),
	15: uint16(37),
	16: uint16('M'),
	17: uint16(38),
	18: uint16('S'),
	19: uint16(27),
	20: uint16('T'),
	21: uint16(21),
	22: uint16('W'),
	23: uint16(29),
	24: uint16('d'),
	25: uint16(50),
	26: uint16('r'),
	27: uint16(52),
	28: uint16('x'),
	29: uint16(94),
}

var map_token5 = [28]uint16_t{
	0:  uint16('#'),
	1:  uint16(257),
	2:  uint16('+'),
	3:  uint16(79),
	4:  uint16('-'),
	5:  uint16(81),
	6:  uint16('0'),
	7:  uint16(253),
	8:  uint16('A'),
	9:  uint16(204),
	10: uint16('C'),
	11: uint16(191),
	12: uint16('L'),
	13: uint16(208),
	14: uint16('M'),
	15: uint16(209),
	16: uint16('S'),
	17: uint16(198),
	18: uint16('T'),
	19: uint16(192),
	20: uint16('W'),
	21: uint16(200),
	22: uint16('r'),
	23: uint16(221),
	24: uint16('\t'),
	25: uint16(188),
	26: uint16(' '),
	27: uint16(188),
}

var map_token6 = [28]uint16_t{
	0:  uint16('\n'),
	1:  uint16(256),
	2:  uint16('#'),
	3:  uint16(257),
	4:  uint16('$'),
	5:  uint16(248),
	6:  uint16('('),
	7:  uint16(97),
	8:  uint16(')'),
	9:  uint16(98),
	10: uint16(','),
	11: uint16(99),
	12: uint16(':'),
	13: uint16(61),
	14: uint16(';'),
	15: uint16(67),
	16: uint16('='),
	17: uint16(60),
	18: uint16('@'),
	19: uint16(100),
	20: uint16(']'),
	21: uint16(68),
	22: uint16('e'),
	23: uint16(245),
	24: uint16('s'),
	25: uint16(236),
	26: uint16('{'),
	27: uint16(62),
}

var map_token7 = [62]uint16_t{
	0:  uint16('\n'),
	1:  uint16(256),
	2:  uint16('#'),
	3:  uint16(257),
	4:  uint16('$'),
	5:  uint16(248),
	6:  uint16('('),
	7:  uint16(97),
	8:  uint16(')'),
	9:  uint16(98),
	10: uint16('+'),
	11: uint16(79),
	12: uint16(','),
	13: uint16(99),
	14: uint16('-'),
	15: uint16(81),
	16: uint16('0'),
	17: uint16(250),
	18: uint16(':'),
	19: uint16(61),
	20: uint16(';'),
	21: uint16(67),
	22: uint16('='),
	23: uint16(60),
	24: uint16('@'),
	25: uint16(100),
	26: uint16('A'),
	27: uint16(140),
	28: uint16('C'),
	29: uint16(130),
	30: uint16('L'),
	31: uint16(146),
	32: uint16('M'),
	33: uint16(147),
	34: uint16('S'),
	35: uint16(137),
	36: uint16('T'),
	37: uint16(131),
	38: uint16('W'),
	39: uint16(139),
	40: uint16('['),
	41: uint16(66),
	42: uint16(']'),
	43: uint16(68),
	44: uint16('d'),
	45: uint16(169),
	46: uint16('e'),
	47: uint16(185),
	48: uint16('r'),
	49: uint16(170),
	50: uint16('s'),
	51: uint16(176),
	52: uint16('x'),
	53: uint16(94),
	54: uint16('{'),
	55: uint16(62),
	56: uint16('}'),
	57: uint16(63),
	58: uint16('\t'),
	59: uint16(124),
	60: uint16(' '),
	61: uint16(124),
}

var map_token8 = [34]uint16_t{
	0:  uint16('\n'),
	1:  uint16(256),
	2:  uint16('#'),
	3:  uint16(257),
	4:  uint16('$'),
	5:  uint16(248),
	6:  uint16('+'),
	7:  uint16(79),
	8:  uint16(','),
	9:  uint16(99),
	10: uint16('-'),
	11: uint16(81),
	12: uint16('0'),
	13: uint16(253),
	14: uint16('A'),
	15: uint16(140),
	16: uint16('C'),
	17: uint16(130),
	18: uint16('L'),
	19: uint16(146),
	20: uint16('M'),
	21: uint16(147),
	22: uint16('S'),
	23: uint16(137),
	24: uint16('T'),
	25: uint16(131),
	26: uint16('W'),
	27: uint16(139),
	28: uint16('r'),
	29: uint16(170),
	30: uint16('\t'),
	31: uint16(125),
	32: uint16(' '),
	33: uint16(125),
}

var map_token9 = [34]uint16_t{
	0:  uint16('\n'),
	1:  uint16(256),
	2:  uint16('#'),
	3:  uint16(257),
	4:  uint16('$'),
	5:  uint16(248),
	6:  uint16('+'),
	7:  uint16(79),
	8:  uint16(','),
	9:  uint16(14),
	10: uint16('-'),
	11: uint16(81),
	12: uint16('0'),
	13: uint16(253),
	14: uint16('A'),
	15: uint16(140),
	16: uint16('C'),
	17: uint16(130),
	18: uint16('L'),
	19: uint16(146),
	20: uint16('M'),
	21: uint16(147),
	22: uint16('S'),
	23: uint16(137),
	24: uint16('T'),
	25: uint16(131),
	26: uint16('W'),
	27: uint16(139),
	28: uint16('r'),
	29: uint16(170),
	30: uint16('\t'),
	31: uint16(126),
	32: uint16(' '),
	33: uint16(126),
}

var map_token10 = [28]uint16_t{
	0:  uint16('#'),
	1:  uint16(257),
	2:  uint16('+'),
	3:  uint16(79),
	4:  uint16('-'),
	5:  uint16(81),
	6:  uint16('0'),
	7:  uint16(253),
	8:  uint16('A'),
	9:  uint16(204),
	10: uint16('C'),
	11: uint16(191),
	12: uint16('L'),
	13: uint16(208),
	14: uint16('M'),
	15: uint16(209),
	16: uint16('S'),
	17: uint16(198),
	18: uint16('T'),
	19: uint16(192),
	20: uint16('W'),
	21: uint16(200),
	22: uint16('r'),
	23: uint16(221),
	24: uint16('\t'),
	25: uint16(188),
	26: uint16(' '),
	27: uint16(188),
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
		if lookahead == int32('f') {
			state = uint16(1)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(2)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(3)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(4)
			goto next_state
		}
		if lookahead == int32('y') {
			state = uint16(5)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(0)
			goto next_state
		}
		return result
	case int32(1):
		if lookahead == int32('a') {
			state = uint16(6)
			goto next_state
		}
		return result
	case int32(2):
		if lookahead == int32('o') {
			state = uint16(7)
			goto next_state
		}
		return result
	case int32(3):
		if lookahead == int32('f') {
			state = uint16(8)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(9)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead == int32('r') {
			state = uint16(10)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead == int32('e') {
			state = uint16(11)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead == int32('l') {
			state = uint16(12)
			goto next_state
		}
		return result
	case int32(7):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_no)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(8):
		if lookahead == int32('f') {
			state = uint16(13)
			goto next_state
		}
		return result
	case int32(9):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_on)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(10):
		if lookahead == int32('u') {
			state = uint16(14)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead == int32('s') {
			state = uint16(15)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead == int32('s') {
			state = uint16(16)
			goto next_state
		}
		return result
	case int32(13):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_off)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(14):
		if lookahead == int32('e') {
			state = uint16(17)
			goto next_state
		}
		return result
	case int32(15):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_yes)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(16):
		if lookahead == int32('e') {
			state = uint16(18)
			goto next_state
		}
		return result
	case int32(17):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_true)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(18):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_false)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var ts_lex_modes = [137]TSLexerMode{
	0: {},
	1: {
		Flex_state: uint16(58),
	},
	2: {
		Flex_state: uint16(2),
	},
	3: {
		Flex_state: uint16(2),
	},
	4: {
		Flex_state: uint16(1),
	},
	5: {
		Flex_state: uint16(6),
	},
	6: {
		Flex_state: uint16(58),
	},
	7: {
		Flex_state: uint16(58),
	},
	8: {
		Flex_state: uint16(4),
	},
	9: {
		Flex_state: uint16(4),
	},
	10: {
		Flex_state: uint16(4),
	},
	11: {
		Flex_state: uint16(4),
	},
	12: {
		Flex_state: uint16(4),
	},
	13: {
		Flex_state: uint16(16),
	},
	14: {
		Flex_state: uint16(4),
	},
	15: {
		Flex_state: uint16(58),
	},
	16: {
		Flex_state: uint16(58),
	},
	17: {
		Flex_state: uint16(58),
	},
	18: {
		Flex_state: uint16(58),
	},
	19: {
		Flex_state: uint16(58),
	},
	20: {
		Flex_state: uint16(58),
	},
	21: {
		Flex_state: uint16(58),
	},
	22: {
		Flex_state: uint16(16),
	},
	23: {
		Flex_state: uint16(58),
	},
	24: {
		Flex_state: uint16(58),
	},
	25: {
		Flex_state: uint16(58),
	},
	26: {
		Flex_state: uint16(16),
	},
	27: {
		Flex_state: uint16(58),
	},
	28: {
		Flex_state: uint16(58),
	},
	29: {
		Flex_state: uint16(58),
	},
	30: {
		Flex_state: uint16(3),
	},
	31: {
		Flex_state: uint16(3),
	},
	32: {
		Flex_state: uint16(3),
	},
	33: {
		Flex_state: uint16(3),
	},
	34: {
		Flex_state: uint16(7),
	},
	35: {
		Flex_state: uint16(5),
	},
	36: {
		Flex_state: uint16(5),
	},
	37: {
		Flex_state: uint16(5),
	},
	38: {
		Flex_state: uint16(5),
	},
	39: {
		Flex_state: uint16(5),
	},
	40: {
		Flex_state: uint16(5),
	},
	41: {
		Flex_state: uint16(5),
	},
	42: {
		Flex_state: uint16(5),
	},
	43: {
		Flex_state: uint16(5),
	},
	44: {
		Flex_state: uint16(3),
	},
	45: {
		Flex_state: uint16(16),
	},
	46: {
		Flex_state: uint16(16),
	},
	47: {
		Flex_state: uint16(16),
	},
	48: {
		Flex_state: uint16(16),
	},
	49: {
		Flex_state: uint16(4),
	},
	50: {
		Flex_state: uint16(3),
	},
	51: {
		Flex_state: uint16(4),
	},
	52: {
		Flex_state: uint16(3),
	},
	53: {
		Flex_state: uint16(3),
	},
	54: {
		Flex_state: uint16(4),
	},
	55: {
		Flex_state: uint16(4),
	},
	56: {
		Flex_state: uint16(58),
	},
	57: {
		Flex_state: uint16(58),
	},
	58: {
		Flex_state: uint16(58),
	},
	59: {
		Flex_state: uint16(58),
	},
	60: {
		Flex_state: uint16(8),
	},
	61: {
		Flex_state: uint16(58),
	},
	62: {
		Flex_state: uint16(58),
	},
	63: {
		Flex_state: uint16(58),
	},
	64: {
		Flex_state: uint16(58),
	},
	65: {
		Flex_state: uint16(5),
	},
	66: {
		Flex_state: uint16(58),
	},
	67: {
		Flex_state: uint16(58),
	},
	68: {
		Flex_state: uint16(58),
	},
	69: {
		Flex_state: uint16(5),
	},
	70: {
		Flex_state: uint16(58),
	},
	71: {
		Flex_state: uint16(5),
	},
	72: {
		Flex_state: uint16(5),
	},
	73: {
		Flex_state: uint16(5),
	},
	74: {
		Flex_state: uint16(58),
	},
	75: {
		Flex_state: uint16(5),
	},
	76: {
		Flex_state: uint16(5),
	},
	77: {
		Flex_state: uint16(5),
	},
	78: {
		Flex_state: uint16(58),
	},
	79: {
		Flex_state: uint16(5),
	},
	80: {
		Flex_state: uint16(58),
	},
	81: {
		Flex_state: uint16(58),
	},
	82: {
		Flex_state: uint16(58),
	},
	83: {
		Flex_state: uint16(58),
	},
	84: {
		Flex_state: uint16(58),
	},
	85: {
		Flex_state: uint16(5),
	},
	86: {
		Flex_state: uint16(58),
	},
	87: {
		Flex_state: uint16(58),
	},
	88: {
		Flex_state: uint16(58),
	},
	89: {
		Flex_state: uint16(58),
	},
	90: {
		Flex_state: uint16(58),
	},
	91: {
		Flex_state: uint16(58),
	},
	92: {
		Flex_state: uint16(58),
	},
	93: {
		Flex_state: uint16(58),
	},
	94: {
		Flex_state: uint16(58),
	},
	95: {
		Flex_state: uint16(58),
	},
	96: {
		Flex_state: uint16(58),
	},
	97: {
		Flex_state: uint16(58),
	},
	98: {
		Flex_state: uint16(58),
	},
	99: {
		Flex_state: uint16(58),
	},
	100: {
		Flex_state: uint16(4),
	},
	101: {
		Flex_state: uint16(58),
	},
	102: {
		Flex_state: uint16(58),
	},
	103: {
		Flex_state: uint16(58),
	},
	104: {
		Flex_state: uint16(58),
	},
	105: {
		Flex_state: uint16(58),
	},
	106: {
		Flex_state: uint16(58),
	},
	107: {
		Flex_state: uint16(58),
	},
	108: {
		Flex_state: uint16(58),
	},
	109: {
		Flex_state: uint16(58),
	},
	110: {
		Flex_state: uint16(5),
	},
	111: {
		Flex_state: uint16(58),
	},
	112: {
		Flex_state: uint16(58),
	},
	113: {
		Flex_state: uint16(58),
	},
	114: {
		Flex_state: uint16(9),
	},
	115: {
		Flex_state: uint16(4),
	},
	116: {
		Flex_state: uint16(58),
	},
	117: {
		Flex_state: uint16(10),
	},
	118: {
		Flex_state: uint16(13),
	},
	119: {
		Flex_state: uint16(58),
	},
	120: {
		Flex_state: uint16(4),
	},
	121: {
		Flex_state: uint16(58),
	},
	122: {
		Flex_state: uint16(9),
	},
	123: {
		Flex_state: uint16(260),
	},
	124: {
		Flex_state: uint16(11),
	},
	125: {
		Flex_state: uint16(58),
	},
	126: {
		Flex_state: uint16(9),
	},
	127: {
		Flex_state: uint16(11),
	},
	128: {
		Flex_state: uint16(9),
	},
	129: {
		Flex_state: uint16(12),
	},
	130: {
		Flex_state: uint16(58),
	},
	131: {
		Flex_state: uint16(58),
	},
	132: {
		Flex_state: uint16(58),
	},
	133: {
		Flex_state: uint16(9),
	},
	134: {
		Flex_state: uint16(58),
	},
	135: {
		Flex_state: uint16(13),
	},
	136: {
		Flex_state: uint16(-libc.Int32FromInt32(1)),
	},
}

var ts_parse_table = [2][94]uint16_t{
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
		51: uint16(1),
		53: uint16(1),
		55: uint16(1),
		56: uint16(3),
	},
	1: {
		0:  uint16(5),
		6:  uint16(7),
		10: uint16(9),
		11: uint16(9),
		12: uint16(11),
		13: uint16(11),
		14: uint16(11),
		49: uint16(13),
		51: uint16(15),
		55: uint16(17),
		56: uint16(19),
		58: uint16(113),
		59: uint16(18),
		60: uint16(18),
		61: uint16(18),
		62: uint16(18),
		63: uint16(18),
		67: uint16(18),
		83: uint16(116),
		85: uint16(18),
		86: uint16(1),
		87: uint16(6),
	},
}

var ts_small_parse_table = [2531]uint16_t{
	0:    uint16(21),
	1:    uint16(3),
	2:    uint16(1),
	3:    uint16(anon_sym_POUND),
	4:    uint16(21),
	5:    uint16(1),
	6:    uint16(sym_string),
	7:    uint16(27),
	8:    uint16(1),
	9:    uint16(aux_sym_number_token1),
	10:   uint16(33),
	11:   uint16(1),
	12:   uint16(anon_sym_DOLLAR),
	13:   uint16(35),
	14:   uint16(1),
	15:   uint16(anon_sym_0),
	16:   uint16(37),
	17:   uint16(1),
	18:   uint16(anon_sym_LF),
	19:   uint16(2),
	20:   uint16(1),
	21:   uint16(sym_comment),
	22:   uint16(12),
	23:   uint16(1),
	24:   uint16(sym_color),
	25:   uint16(32),
	26:   uint16(1),
	27:   uint16(sym_number),
	28:   uint16(44),
	29:   uint16(1),
	30:   uint16(sym__zero),
	31:   uint16(57),
	32:   uint16(1),
	33:   uint16(sym__value),
	34:   uint16(75),
	35:   uint16(1),
	36:   uint16(sym__linebreak),
	37:   uint16(78),
	38:   uint16(1),
	39:   uint16(sym_position),
	40:   uint16(99),
	41:   uint16(1),
	42:   uint16(sym_params),
	43:   uint16(25),
	44:   uint16(2),
	45:   uint16(anon_sym_PLUS),
	46:   uint16(anon_sym_DASH),
	47:   uint16(29),
	48:   uint16(2),
	49:   uint16(anon_sym_rgb),
	50:   uint16(anon_sym_rgba),
	51:   uint16(9),
	52:   uint16(2),
	53:   uint16(sym_mod),
	54:   uint16(sym_variable),
	55:   uint16(55),
	56:   uint16(2),
	57:   uint16(sym_legacy_hex),
	58:   uint16(sym_rgb),
	59:   uint16(92),
	60:   uint16(5),
	61:   uint16(sym_boolean),
	62:   uint16(sym_vec2),
	63:   uint16(sym_gradient),
	64:   uint16(sym_display),
	65:   uint16(sym_keys),
	66:   uint16(23),
	67:   uint16(6),
	68:   uint16(anon_sym_true),
	69:   uint16(anon_sym_false),
	70:   uint16(anon_sym_on),
	71:   uint16(anon_sym_off),
	72:   uint16(anon_sym_yes),
	73:   uint16(anon_sym_no),
	74:   uint16(31),
	75:   uint16(14),
	76:   uint16(anon_sym_SHIFT),
	77:   uint16(anon_sym_CAPS),
	78:   uint16(anon_sym_CTRL),
	79:   uint16(anon_sym_CONTROL),
	80:   uint16(anon_sym_ALT),
	81:   uint16(anon_sym_ALT_L),
	82:   uint16(anon_sym_MOD2),
	83:   uint16(anon_sym_MOD3),
	84:   uint16(anon_sym_SUPER),
	85:   uint16(anon_sym_WIN),
	86:   uint16(anon_sym_LOGO),
	87:   uint16(anon_sym_MOD4),
	88:   uint16(anon_sym_MOD5),
	89:   uint16(anon_sym_TAB),
	90:   uint16(21),
	91:   uint16(3),
	92:   uint16(1),
	93:   uint16(anon_sym_POUND),
	94:   uint16(21),
	95:   uint16(1),
	96:   uint16(sym_string),
	97:   uint16(27),
	98:   uint16(1),
	99:   uint16(aux_sym_number_token1),
	100:  uint16(33),
	101:  uint16(1),
	102:  uint16(anon_sym_DOLLAR),
	103:  uint16(35),
	104:  uint16(1),
	105:  uint16(anon_sym_0),
	106:  uint16(39),
	107:  uint16(1),
	108:  uint16(anon_sym_LF),
	109:  uint16(3),
	110:  uint16(1),
	111:  uint16(sym_comment),
	112:  uint16(12),
	113:  uint16(1),
	114:  uint16(sym_color),
	115:  uint16(24),
	116:  uint16(1),
	117:  uint16(sym__linebreak),
	118:  uint16(32),
	119:  uint16(1),
	120:  uint16(sym_number),
	121:  uint16(44),
	122:  uint16(1),
	123:  uint16(sym__zero),
	124:  uint16(56),
	125:  uint16(1),
	126:  uint16(sym__value),
	127:  uint16(78),
	128:  uint16(1),
	129:  uint16(sym_position),
	130:  uint16(107),
	131:  uint16(1),
	132:  uint16(sym_params),
	133:  uint16(25),
	134:  uint16(2),
	135:  uint16(anon_sym_PLUS),
	136:  uint16(anon_sym_DASH),
	137:  uint16(29),
	138:  uint16(2),
	139:  uint16(anon_sym_rgb),
	140:  uint16(anon_sym_rgba),
	141:  uint16(9),
	142:  uint16(2),
	143:  uint16(sym_mod),
	144:  uint16(sym_variable),
	145:  uint16(55),
	146:  uint16(2),
	147:  uint16(sym_legacy_hex),
	148:  uint16(sym_rgb),
	149:  uint16(92),
	150:  uint16(5),
	151:  uint16(sym_boolean),
	152:  uint16(sym_vec2),
	153:  uint16(sym_gradient),
	154:  uint16(sym_display),
	155:  uint16(sym_keys),
	156:  uint16(23),
	157:  uint16(6),
	158:  uint16(anon_sym_true),
	159:  uint16(anon_sym_false),
	160:  uint16(anon_sym_on),
	161:  uint16(anon_sym_off),
	162:  uint16(anon_sym_yes),
	163:  uint16(anon_sym_no),
	164:  uint16(31),
	165:  uint16(14),
	166:  uint16(anon_sym_SHIFT),
	167:  uint16(anon_sym_CAPS),
	168:  uint16(anon_sym_CTRL),
	169:  uint16(anon_sym_CONTROL),
	170:  uint16(anon_sym_ALT),
	171:  uint16(anon_sym_ALT_L),
	172:  uint16(anon_sym_MOD2),
	173:  uint16(anon_sym_MOD3),
	174:  uint16(anon_sym_SUPER),
	175:  uint16(anon_sym_WIN),
	176:  uint16(anon_sym_LOGO),
	177:  uint16(anon_sym_MOD4),
	178:  uint16(anon_sym_MOD5),
	179:  uint16(anon_sym_TAB),
	180:  uint16(19),
	181:  uint16(3),
	182:  uint16(1),
	183:  uint16(anon_sym_POUND),
	184:  uint16(21),
	185:  uint16(1),
	186:  uint16(sym_string),
	187:  uint16(27),
	188:  uint16(1),
	189:  uint16(aux_sym_number_token1),
	190:  uint16(33),
	191:  uint16(1),
	192:  uint16(anon_sym_DOLLAR),
	193:  uint16(35),
	194:  uint16(1),
	195:  uint16(anon_sym_0),
	196:  uint16(4),
	197:  uint16(1),
	198:  uint16(sym_comment),
	199:  uint16(12),
	200:  uint16(1),
	201:  uint16(sym_color),
	202:  uint16(32),
	203:  uint16(1),
	204:  uint16(sym_number),
	205:  uint16(44),
	206:  uint16(1),
	207:  uint16(sym__zero),
	208:  uint16(78),
	209:  uint16(1),
	210:  uint16(sym_position),
	211:  uint16(111),
	212:  uint16(1),
	213:  uint16(sym__value),
	214:  uint16(25),
	215:  uint16(2),
	216:  uint16(anon_sym_PLUS),
	217:  uint16(anon_sym_DASH),
	218:  uint16(29),
	219:  uint16(2),
	220:  uint16(anon_sym_rgb),
	221:  uint16(anon_sym_rgba),
	222:  uint16(41),
	223:  uint16(2),
	224:  uint16(anon_sym_COMMA),
	225:  uint16(anon_sym_LF),
	226:  uint16(9),
	227:  uint16(2),
	228:  uint16(sym_mod),
	229:  uint16(sym_variable),
	230:  uint16(55),
	231:  uint16(2),
	232:  uint16(sym_legacy_hex),
	233:  uint16(sym_rgb),
	234:  uint16(92),
	235:  uint16(5),
	236:  uint16(sym_boolean),
	237:  uint16(sym_vec2),
	238:  uint16(sym_gradient),
	239:  uint16(sym_display),
	240:  uint16(sym_keys),
	241:  uint16(23),
	242:  uint16(6),
	243:  uint16(anon_sym_true),
	244:  uint16(anon_sym_false),
	245:  uint16(anon_sym_on),
	246:  uint16(anon_sym_off),
	247:  uint16(anon_sym_yes),
	248:  uint16(anon_sym_no),
	249:  uint16(31),
	250:  uint16(14),
	251:  uint16(anon_sym_SHIFT),
	252:  uint16(anon_sym_CAPS),
	253:  uint16(anon_sym_CTRL),
	254:  uint16(anon_sym_CONTROL),
	255:  uint16(anon_sym_ALT),
	256:  uint16(anon_sym_ALT_L),
	257:  uint16(anon_sym_MOD2),
	258:  uint16(anon_sym_MOD3),
	259:  uint16(anon_sym_SUPER),
	260:  uint16(anon_sym_WIN),
	261:  uint16(anon_sym_LOGO),
	262:  uint16(anon_sym_MOD4),
	263:  uint16(anon_sym_MOD5),
	264:  uint16(anon_sym_TAB),
	265:  uint16(11),
	266:  uint16(3),
	267:  uint16(1),
	268:  uint16(anon_sym_POUND),
	269:  uint16(27),
	270:  uint16(1),
	271:  uint16(aux_sym_number_token1),
	272:  uint16(35),
	273:  uint16(1),
	274:  uint16(anon_sym_0),
	275:  uint16(43),
	276:  uint16(1),
	277:  uint16(sym_string_literal),
	278:  uint16(5),
	279:  uint16(1),
	280:  uint16(sym_comment),
	281:  uint16(100),
	282:  uint16(1),
	283:  uint16(sym__zero),
	284:  uint16(25),
	285:  uint16(2),
	286:  uint16(anon_sym_PLUS),
	287:  uint16(anon_sym_DASH),
	288:  uint16(29),
	289:  uint16(2),
	290:  uint16(anon_sym_rgb),
	291:  uint16(anon_sym_rgba),
	292:  uint16(55),
	293:  uint16(2),
	294:  uint16(sym_legacy_hex),
	295:  uint16(sym_rgb),
	296:  uint16(81),
	297:  uint16(3),
	298:  uint16(sym_number),
	299:  uint16(sym_color),
	300:  uint16(sym_mod),
	301:  uint16(31),
	302:  uint16(14),
	303:  uint16(anon_sym_SHIFT),
	304:  uint16(anon_sym_CAPS),
	305:  uint16(anon_sym_CTRL),
	306:  uint16(anon_sym_CONTROL),
	307:  uint16(anon_sym_ALT),
	308:  uint16(anon_sym_ALT_L),
	309:  uint16(anon_sym_MOD2),
	310:  uint16(anon_sym_MOD3),
	311:  uint16(anon_sym_SUPER),
	312:  uint16(anon_sym_WIN),
	313:  uint16(anon_sym_LOGO),
	314:  uint16(anon_sym_MOD4),
	315:  uint16(anon_sym_MOD5),
	316:  uint16(anon_sym_TAB),
	317:  uint16(12),
	318:  uint16(7),
	319:  uint16(1),
	320:  uint16(anon_sym_source),
	321:  uint16(13),
	322:  uint16(1),
	323:  uint16(sym_name),
	324:  uint16(15),
	325:  uint16(1),
	326:  uint16(anon_sym_DOLLAR),
	327:  uint16(17),
	328:  uint16(1),
	329:  uint16(anon_sym_LF),
	330:  uint16(19),
	331:  uint16(1),
	332:  uint16(anon_sym_POUND),
	333:  uint16(45),
	334:  uint16(1),
	336:  uint16(6),
	337:  uint16(1),
	338:  uint16(sym_comment),
	339:  uint16(7),
	340:  uint16(1),
	341:  uint16(aux_sym_configuration_repeat1),
	342:  uint16(116),
	343:  uint16(1),
	344:  uint16(sym_variable),
	345:  uint16(9),
	346:  uint16(2),
	347:  uint16(anon_sym_exec_DASHonce),
	348:  uint16(anon_sym_exec),
	349:  uint16(11),
	350:  uint16(3),
	351:  uint16(anon_sym_execr_DASHonce),
	352:  uint16(anon_sym_execr),
	353:  uint16(anon_sym_exec_DASHshutdown),
	354:  uint16(18),
	355:  uint16(7),
	356:  uint16(sym_declaration),
	357:  uint16(sym_assignment),
	358:  uint16(sym_keyword),
	359:  uint16(sym_section),
	360:  uint16(sym_source),
	361:  uint16(sym_exec),
	362:  uint16(sym__linebreak),
	363:  uint16(11),
	364:  uint16(19),
	365:  uint16(1),
	366:  uint16(anon_sym_POUND),
	367:  uint16(47),
	368:  uint16(1),
	370:  uint16(49),
	371:  uint16(1),
	372:  uint16(anon_sym_source),
	373:  uint16(58),
	374:  uint16(1),
	375:  uint16(sym_name),
	376:  uint16(61),
	377:  uint16(1),
	378:  uint16(anon_sym_DOLLAR),
	379:  uint16(64),
	380:  uint16(1),
	381:  uint16(anon_sym_LF),
	382:  uint16(116),
	383:  uint16(1),
	384:  uint16(sym_variable),
	385:  uint16(52),
	386:  uint16(2),
	387:  uint16(anon_sym_exec_DASHonce),
	388:  uint16(anon_sym_exec),
	389:  uint16(7),
	390:  uint16(2),
	391:  uint16(sym_comment),
	392:  uint16(aux_sym_configuration_repeat1),
	393:  uint16(55),
	394:  uint16(3),
	395:  uint16(anon_sym_execr_DASHonce),
	396:  uint16(anon_sym_execr),
	397:  uint16(anon_sym_exec_DASHshutdown),
	398:  uint16(18),
	399:  uint16(7),
	400:  uint16(sym_declaration),
	401:  uint16(sym_assignment),
	402:  uint16(sym_keyword),
	403:  uint16(sym_section),
	404:  uint16(sym_source),
	405:  uint16(sym_exec),
	406:  uint16(sym__linebreak),
	407:  uint16(4),
	408:  uint16(19),
	409:  uint16(1),
	410:  uint16(anon_sym_POUND),
	411:  uint16(69),
	412:  uint16(1),
	413:  uint16(anon_sym_ALT),
	414:  uint16(8),
	415:  uint16(1),
	416:  uint16(sym_comment),
	417:  uint16(67),
	418:  uint16(16),
	419:  uint16(anon_sym_EQ),
	420:  uint16(anon_sym_COMMA),
	421:  uint16(anon_sym_SHIFT),
	422:  uint16(anon_sym_CAPS),
	423:  uint16(anon_sym_CTRL),
	424:  uint16(anon_sym_CONTROL),
	425:  uint16(anon_sym_ALT_L),
	426:  uint16(anon_sym_MOD2),
	427:  uint16(anon_sym_MOD3),
	428:  uint16(anon_sym_SUPER),
	429:  uint16(anon_sym_WIN),
	430:  uint16(anon_sym_LOGO),
	431:  uint16(anon_sym_MOD4),
	432:  uint16(anon_sym_MOD5),
	433:  uint16(anon_sym_TAB),
	434:  uint16(anon_sym_LF),
	435:  uint16(6),
	436:  uint16(19),
	437:  uint16(1),
	438:  uint16(anon_sym_POUND),
	439:  uint16(31),
	440:  uint16(1),
	441:  uint16(anon_sym_ALT),
	442:  uint16(9),
	443:  uint16(1),
	444:  uint16(sym_comment),
	445:  uint16(98),
	446:  uint16(1),
	447:  uint16(sym_mod),
	448:  uint16(71),
	449:  uint16(2),
	450:  uint16(anon_sym_COMMA),
	451:  uint16(anon_sym_LF),
	452:  uint16(73),
	453:  uint16(13),
	454:  uint16(anon_sym_SHIFT),
	455:  uint16(anon_sym_CAPS),
	456:  uint16(anon_sym_CTRL),
	457:  uint16(anon_sym_CONTROL),
	458:  uint16(anon_sym_ALT_L),
	459:  uint16(anon_sym_MOD2),
	460:  uint16(anon_sym_MOD3),
	461:  uint16(anon_sym_SUPER),
	462:  uint16(anon_sym_WIN),
	463:  uint16(anon_sym_LOGO),
	464:  uint16(anon_sym_MOD4),
	465:  uint16(anon_sym_MOD5),
	466:  uint16(anon_sym_TAB),
	467:  uint16(4),
	468:  uint16(19),
	469:  uint16(1),
	470:  uint16(anon_sym_POUND),
	471:  uint16(77),
	472:  uint16(1),
	473:  uint16(anon_sym_ALT),
	474:  uint16(10),
	475:  uint16(1),
	476:  uint16(sym_comment),
	477:  uint16(75),
	478:  uint16(15),
	479:  uint16(anon_sym_COMMA),
	480:  uint16(anon_sym_SHIFT),
	481:  uint16(anon_sym_CAPS),
	482:  uint16(anon_sym_CTRL),
	483:  uint16(anon_sym_CONTROL),
	484:  uint16(anon_sym_ALT_L),
	485:  uint16(anon_sym_MOD2),
	486:  uint16(anon_sym_MOD3),
	487:  uint16(anon_sym_SUPER),
	488:  uint16(anon_sym_WIN),
	489:  uint16(anon_sym_LOGO),
	490:  uint16(anon_sym_MOD4),
	491:  uint16(anon_sym_MOD5),
	492:  uint16(anon_sym_TAB),
	493:  uint16(anon_sym_LF),
	494:  uint16(12),
	495:  uint16(19),
	496:  uint16(1),
	497:  uint16(anon_sym_POUND),
	498:  uint16(29),
	499:  uint16(1),
	500:  uint16(anon_sym_rgb),
	501:  uint16(35),
	502:  uint16(1),
	503:  uint16(anon_sym_0),
	504:  uint16(79),
	505:  uint16(1),
	506:  uint16(anon_sym_rgba),
	507:  uint16(83),
	508:  uint16(1),
	509:  uint16(aux_sym_angle_token1),
	510:  uint16(11),
	511:  uint16(1),
	512:  uint16(sym_comment),
	513:  uint16(14),
	514:  uint16(1),
	515:  uint16(aux_sym_gradient_repeat1),
	516:  uint16(54),
	517:  uint16(1),
	518:  uint16(sym_color),
	519:  uint16(84),
	520:  uint16(1),
	521:  uint16(sym_angle),
	522:  uint16(120),
	523:  uint16(1),
	524:  uint16(sym__zero),
	525:  uint16(81),
	526:  uint16(2),
	527:  uint16(anon_sym_COMMA),
	528:  uint16(anon_sym_LF),
	529:  uint16(55),
	530:  uint16(2),
	531:  uint16(sym_legacy_hex),
	532:  uint16(sym_rgb),
	533:  uint16(12),
	534:  uint16(19),
	535:  uint16(1),
	536:  uint16(anon_sym_POUND),
	537:  uint16(29),
	538:  uint16(1),
	539:  uint16(anon_sym_rgb),
	540:  uint16(35),
	541:  uint16(1),
	542:  uint16(anon_sym_0),
	543:  uint16(79),
	544:  uint16(1),
	545:  uint16(anon_sym_rgba),
	546:  uint16(83),
	547:  uint16(1),
	548:  uint16(aux_sym_angle_token1),
	549:  uint16(11),
	550:  uint16(1),
	551:  uint16(aux_sym_gradient_repeat1),
	552:  uint16(12),
	553:  uint16(1),
	554:  uint16(sym_comment),
	555:  uint16(54),
	556:  uint16(1),
	557:  uint16(sym_color),
	558:  uint16(96),
	559:  uint16(1),
	560:  uint16(sym_angle),
	561:  uint16(120),
	562:  uint16(1),
	563:  uint16(sym__zero),
	564:  uint16(71),
	565:  uint16(2),
	566:  uint16(anon_sym_COMMA),
	567:  uint16(anon_sym_LF),
	568:  uint16(55),
	569:  uint16(2),
	570:  uint16(sym_legacy_hex),
	571:  uint16(sym_rgb),
	572:  uint16(11),
	573:  uint16(3),
	574:  uint16(1),
	575:  uint16(anon_sym_POUND),
	576:  uint16(89),
	577:  uint16(1),
	578:  uint16(aux_sym_number_token1),
	579:  uint16(91),
	580:  uint16(1),
	581:  uint16(anon_sym_0),
	582:  uint16(93),
	583:  uint16(1),
	584:  uint16(sym__window_rule_argument),
	585:  uint16(13),
	586:  uint16(1),
	587:  uint16(sym_comment),
	588:  uint16(22),
	589:  uint16(1),
	590:  uint16(aux_sym_arguments_repeat1),
	591:  uint16(45),
	592:  uint16(1),
	593:  uint16(sym_number),
	594:  uint16(47),
	595:  uint16(1),
	596:  uint16(sym__zero),
	597:  uint16(108),
	598:  uint16(1),
	599:  uint16(sym_arguments),
	600:  uint16(85),
	601:  uint16(2),
	602:  uint16(anon_sym_SEMI),
	603:  uint16(anon_sym_RBRACK),
	604:  uint16(87),
	605:  uint16(2),
	606:  uint16(anon_sym_PLUS),
	607:  uint16(anon_sym_DASH),
	608:  uint16(10),
	609:  uint16(19),
	610:  uint16(1),
	611:  uint16(anon_sym_POUND),
	612:  uint16(95),
	613:  uint16(1),
	614:  uint16(anon_sym_rgb),
	615:  uint16(98),
	616:  uint16(1),
	617:  uint16(anon_sym_rgba),
	618:  uint16(103),
	619:  uint16(1),
	620:  uint16(aux_sym_angle_token1),
	621:  uint16(105),
	622:  uint16(1),
	623:  uint16(anon_sym_0),
	624:  uint16(54),
	625:  uint16(1),
	626:  uint16(sym_color),
	627:  uint16(120),
	628:  uint16(1),
	629:  uint16(sym__zero),
	630:  uint16(101),
	631:  uint16(2),
	632:  uint16(anon_sym_COMMA),
	633:  uint16(anon_sym_LF),
	634:  uint16(14),
	635:  uint16(2),
	636:  uint16(sym_comment),
	637:  uint16(aux_sym_gradient_repeat1),
	638:  uint16(55),
	639:  uint16(2),
	640:  uint16(sym_legacy_hex),
	641:  uint16(sym_rgb),
	642:  uint16(4),
	643:  uint16(19),
	644:  uint16(1),
	645:  uint16(anon_sym_POUND),
	646:  uint16(15),
	647:  uint16(1),
	648:  uint16(sym_comment),
	649:  uint16(108),
	650:  uint16(3),
	652:  uint16(anon_sym_DOLLAR),
	653:  uint16(anon_sym_LF),
	654:  uint16(110),
	655:  uint16(7),
	656:  uint16(anon_sym_source),
	657:  uint16(anon_sym_exec_DASHonce),
	658:  uint16(anon_sym_exec),
	659:  uint16(anon_sym_execr_DASHonce),
	660:  uint16(anon_sym_execr),
	661:  uint16(anon_sym_exec_DASHshutdown),
	662:  uint16(sym_name),
	663:  uint16(4),
	664:  uint16(19),
	665:  uint16(1),
	666:  uint16(anon_sym_POUND),
	667:  uint16(16),
	668:  uint16(1),
	669:  uint16(sym_comment),
	670:  uint16(112),
	671:  uint16(3),
	673:  uint16(anon_sym_DOLLAR),
	674:  uint16(anon_sym_LF),
	675:  uint16(114),
	676:  uint16(7),
	677:  uint16(anon_sym_source),
	678:  uint16(anon_sym_exec_DASHonce),
	679:  uint16(anon_sym_exec),
	680:  uint16(anon_sym_execr_DASHonce),
	681:  uint16(anon_sym_execr),
	682:  uint16(anon_sym_exec_DASHshutdown),
	683:  uint16(sym_name),
	684:  uint16(4),
	685:  uint16(19),
	686:  uint16(1),
	687:  uint16(anon_sym_POUND),
	688:  uint16(17),
	689:  uint16(1),
	690:  uint16(sym_comment),
	691:  uint16(116),
	692:  uint16(3),
	694:  uint16(anon_sym_DOLLAR),
	695:  uint16(anon_sym_LF),
	696:  uint16(118),
	697:  uint16(7),
	698:  uint16(anon_sym_source),
	699:  uint16(anon_sym_exec_DASHonce),
	700:  uint16(anon_sym_exec),
	701:  uint16(anon_sym_execr_DASHonce),
	702:  uint16(anon_sym_execr),
	703:  uint16(anon_sym_exec_DASHshutdown),
	704:  uint16(sym_name),
	705:  uint16(4),
	706:  uint16(19),
	707:  uint16(1),
	708:  uint16(anon_sym_POUND),
	709:  uint16(18),
	710:  uint16(1),
	711:  uint16(sym_comment),
	712:  uint16(120),
	713:  uint16(3),
	715:  uint16(anon_sym_DOLLAR),
	716:  uint16(anon_sym_LF),
	717:  uint16(122),
	718:  uint16(7),
	719:  uint16(anon_sym_source),
	720:  uint16(anon_sym_exec_DASHonce),
	721:  uint16(anon_sym_exec),
	722:  uint16(anon_sym_execr_DASHonce),
	723:  uint16(anon_sym_execr),
	724:  uint16(anon_sym_exec_DASHshutdown),
	725:  uint16(sym_name),
	726:  uint16(4),
	727:  uint16(19),
	728:  uint16(1),
	729:  uint16(anon_sym_POUND),
	730:  uint16(19),
	731:  uint16(1),
	732:  uint16(sym_comment),
	733:  uint16(124),
	734:  uint16(3),
	736:  uint16(anon_sym_DOLLAR),
	737:  uint16(anon_sym_LF),
	738:  uint16(126),
	739:  uint16(7),
	740:  uint16(anon_sym_source),
	741:  uint16(anon_sym_exec_DASHonce),
	742:  uint16(anon_sym_exec),
	743:  uint16(anon_sym_execr_DASHonce),
	744:  uint16(anon_sym_execr),
	745:  uint16(anon_sym_exec_DASHshutdown),
	746:  uint16(sym_name),
	747:  uint16(4),
	748:  uint16(19),
	749:  uint16(1),
	750:  uint16(anon_sym_POUND),
	751:  uint16(20),
	752:  uint16(1),
	753:  uint16(sym_comment),
	754:  uint16(128),
	755:  uint16(3),
	757:  uint16(anon_sym_DOLLAR),
	758:  uint16(anon_sym_LF),
	759:  uint16(130),
	760:  uint16(7),
	761:  uint16(anon_sym_source),
	762:  uint16(anon_sym_exec_DASHonce),
	763:  uint16(anon_sym_exec),
	764:  uint16(anon_sym_execr_DASHonce),
	765:  uint16(anon_sym_execr),
	766:  uint16(anon_sym_exec_DASHshutdown),
	767:  uint16(sym_name),
	768:  uint16(4),
	769:  uint16(19),
	770:  uint16(1),
	771:  uint16(anon_sym_POUND),
	772:  uint16(21),
	773:  uint16(1),
	774:  uint16(sym_comment),
	775:  uint16(132),
	776:  uint16(3),
	778:  uint16(anon_sym_DOLLAR),
	779:  uint16(anon_sym_LF),
	780:  uint16(134),
	781:  uint16(7),
	782:  uint16(anon_sym_source),
	783:  uint16(anon_sym_exec_DASHonce),
	784:  uint16(anon_sym_exec),
	785:  uint16(anon_sym_execr_DASHonce),
	786:  uint16(anon_sym_execr),
	787:  uint16(anon_sym_exec_DASHshutdown),
	788:  uint16(sym_name),
	789:  uint16(10),
	790:  uint16(3),
	791:  uint16(1),
	792:  uint16(anon_sym_POUND),
	793:  uint16(89),
	794:  uint16(1),
	795:  uint16(aux_sym_number_token1),
	796:  uint16(91),
	797:  uint16(1),
	798:  uint16(anon_sym_0),
	799:  uint16(93),
	800:  uint16(1),
	801:  uint16(sym__window_rule_argument),
	802:  uint16(22),
	803:  uint16(1),
	804:  uint16(sym_comment),
	805:  uint16(26),
	806:  uint16(1),
	807:  uint16(aux_sym_arguments_repeat1),
	808:  uint16(45),
	809:  uint16(1),
	810:  uint16(sym_number),
	811:  uint16(47),
	812:  uint16(1),
	813:  uint16(sym__zero),
	814:  uint16(87),
	815:  uint16(2),
	816:  uint16(anon_sym_PLUS),
	817:  uint16(anon_sym_DASH),
	818:  uint16(136),
	819:  uint16(2),
	820:  uint16(anon_sym_SEMI),
	821:  uint16(anon_sym_RBRACK),
	822:  uint16(4),
	823:  uint16(19),
	824:  uint16(1),
	825:  uint16(anon_sym_POUND),
	826:  uint16(23),
	827:  uint16(1),
	828:  uint16(sym_comment),
	829:  uint16(138),
	830:  uint16(3),
	832:  uint16(anon_sym_DOLLAR),
	833:  uint16(anon_sym_LF),
	834:  uint16(140),
	835:  uint16(7),
	836:  uint16(anon_sym_source),
	837:  uint16(anon_sym_exec_DASHonce),
	838:  uint16(anon_sym_exec),
	839:  uint16(anon_sym_execr_DASHonce),
	840:  uint16(anon_sym_execr),
	841:  uint16(anon_sym_exec_DASHshutdown),
	842:  uint16(sym_name),
	843:  uint16(4),
	844:  uint16(19),
	845:  uint16(1),
	846:  uint16(anon_sym_POUND),
	847:  uint16(24),
	848:  uint16(1),
	849:  uint16(sym_comment),
	850:  uint16(142),
	851:  uint16(3),
	853:  uint16(anon_sym_DOLLAR),
	854:  uint16(anon_sym_LF),
	855:  uint16(144),
	856:  uint16(7),
	857:  uint16(anon_sym_source),
	858:  uint16(anon_sym_exec_DASHonce),
	859:  uint16(anon_sym_exec),
	860:  uint16(anon_sym_execr_DASHonce),
	861:  uint16(anon_sym_execr),
	862:  uint16(anon_sym_exec_DASHshutdown),
	863:  uint16(sym_name),
	864:  uint16(4),
	865:  uint16(19),
	866:  uint16(1),
	867:  uint16(anon_sym_POUND),
	868:  uint16(25),
	869:  uint16(1),
	870:  uint16(sym_comment),
	871:  uint16(146),
	872:  uint16(3),
	874:  uint16(anon_sym_DOLLAR),
	875:  uint16(anon_sym_LF),
	876:  uint16(148),
	877:  uint16(7),
	878:  uint16(anon_sym_source),
	879:  uint16(anon_sym_exec_DASHonce),
	880:  uint16(anon_sym_exec),
	881:  uint16(anon_sym_execr_DASHonce),
	882:  uint16(anon_sym_execr),
	883:  uint16(anon_sym_exec_DASHshutdown),
	884:  uint16(sym_name),
	885:  uint16(9),
	886:  uint16(3),
	887:  uint16(1),
	888:  uint16(anon_sym_POUND),
	889:  uint16(155),
	890:  uint16(1),
	891:  uint16(aux_sym_number_token1),
	892:  uint16(158),
	893:  uint16(1),
	894:  uint16(anon_sym_0),
	895:  uint16(161),
	896:  uint16(1),
	897:  uint16(sym__window_rule_argument),
	898:  uint16(45),
	899:  uint16(1),
	900:  uint16(sym_number),
	901:  uint16(47),
	902:  uint16(1),
	903:  uint16(sym__zero),
	904:  uint16(150),
	905:  uint16(2),
	906:  uint16(anon_sym_SEMI),
	907:  uint16(anon_sym_RBRACK),
	908:  uint16(152),
	909:  uint16(2),
	910:  uint16(anon_sym_PLUS),
	911:  uint16(anon_sym_DASH),
	912:  uint16(26),
	913:  uint16(2),
	914:  uint16(sym_comment),
	915:  uint16(aux_sym_arguments_repeat1),
	916:  uint16(4),
	917:  uint16(19),
	918:  uint16(1),
	919:  uint16(anon_sym_POUND),
	920:  uint16(27),
	921:  uint16(1),
	922:  uint16(sym_comment),
	923:  uint16(164),
	924:  uint16(3),
	926:  uint16(anon_sym_DOLLAR),
	927:  uint16(anon_sym_LF),
	928:  uint16(166),
	929:  uint16(7),
	930:  uint16(anon_sym_source),
	931:  uint16(anon_sym_exec_DASHonce),
	932:  uint16(anon_sym_exec),
	933:  uint16(anon_sym_execr_DASHonce),
	934:  uint16(anon_sym_execr),
	935:  uint16(anon_sym_exec_DASHshutdown),
	936:  uint16(sym_name),
	937:  uint16(4),
	938:  uint16(19),
	939:  uint16(1),
	940:  uint16(anon_sym_POUND),
	941:  uint16(28),
	942:  uint16(1),
	943:  uint16(sym_comment),
	944:  uint16(168),
	945:  uint16(3),
	947:  uint16(anon_sym_DOLLAR),
	948:  uint16(anon_sym_LF),
	949:  uint16(170),
	950:  uint16(7),
	951:  uint16(anon_sym_source),
	952:  uint16(anon_sym_exec_DASHonce),
	953:  uint16(anon_sym_exec),
	954:  uint16(anon_sym_execr_DASHonce),
	955:  uint16(anon_sym_execr),
	956:  uint16(anon_sym_exec_DASHshutdown),
	957:  uint16(sym_name),
	958:  uint16(4),
	959:  uint16(19),
	960:  uint16(1),
	961:  uint16(anon_sym_POUND),
	962:  uint16(29),
	963:  uint16(1),
	964:  uint16(sym_comment),
	965:  uint16(172),
	966:  uint16(3),
	968:  uint16(anon_sym_DOLLAR),
	969:  uint16(anon_sym_LF),
	970:  uint16(174),
	971:  uint16(7),
	972:  uint16(anon_sym_source),
	973:  uint16(anon_sym_exec_DASHonce),
	974:  uint16(anon_sym_exec),
	975:  uint16(anon_sym_execr_DASHonce),
	976:  uint16(anon_sym_execr),
	977:  uint16(anon_sym_exec_DASHshutdown),
	978:  uint16(sym_name),
	979:  uint16(4),
	980:  uint16(19),
	981:  uint16(1),
	982:  uint16(anon_sym_POUND),
	983:  uint16(30),
	984:  uint16(1),
	985:  uint16(sym_comment),
	986:  uint16(178),
	987:  uint16(2),
	988:  uint16(aux_sym_number_token1),
	989:  uint16(anon_sym_0),
	990:  uint16(176),
	991:  uint16(7),
	992:  uint16(anon_sym_PLUS),
	993:  uint16(anon_sym_DASH),
	994:  uint16(anon_sym_x),
	995:  uint16(anon_sym_RPAREN),
	996:  uint16(anon_sym_COMMA),
	997:  uint16(anon_sym_AT),
	998:  uint16(anon_sym_LF),
	999:  uint16(4),
	1000: uint16(19),
	1001: uint16(1),
	1002: uint16(anon_sym_POUND),
	1003: uint16(31),
	1004: uint16(1),
	1005: uint16(sym_comment),
	1006: uint16(182),
	1007: uint16(2),
	1008: uint16(aux_sym_number_token1),
	1009: uint16(anon_sym_0),
	1010: uint16(180),
	1011: uint16(7),
	1012: uint16(anon_sym_PLUS),
	1013: uint16(anon_sym_DASH),
	1014: uint16(anon_sym_x),
	1015: uint16(anon_sym_RPAREN),
	1016: uint16(anon_sym_COMMA),
	1017: uint16(anon_sym_AT),
	1018: uint16(anon_sym_LF),
	1019: uint16(9),
	1020: uint16(19),
	1021: uint16(1),
	1022: uint16(anon_sym_POUND),
	1023: uint16(27),
	1024: uint16(1),
	1025: uint16(aux_sym_number_token1),
	1026: uint16(35),
	1027: uint16(1),
	1028: uint16(anon_sym_0),
	1029: uint16(186),
	1030: uint16(1),
	1031: uint16(anon_sym_x),
	1032: uint16(30),
	1033: uint16(1),
	1034: uint16(sym__zero),
	1035: uint16(32),
	1036: uint16(1),
	1037: uint16(sym_comment),
	1038: uint16(94),
	1039: uint16(1),
	1040: uint16(sym_number),
	1041: uint16(71),
	1042: uint16(2),
	1043: uint16(anon_sym_COMMA),
	1044: uint16(anon_sym_LF),
	1045: uint16(184),
	1046: uint16(2),
	1047: uint16(anon_sym_PLUS),
	1048: uint16(anon_sym_DASH),
	1049: uint16(4),
	1050: uint16(19),
	1051: uint16(1),
	1052: uint16(anon_sym_POUND),
	1053: uint16(33),
	1054: uint16(1),
	1055: uint16(sym_comment),
	1056: uint16(190),
	1057: uint16(2),
	1058: uint16(aux_sym_number_token1),
	1059: uint16(anon_sym_0),
	1060: uint16(188),
	1061: uint16(7),
	1062: uint16(anon_sym_PLUS),
	1063: uint16(anon_sym_DASH),
	1064: uint16(anon_sym_x),
	1065: uint16(anon_sym_RPAREN),
	1066: uint16(anon_sym_COMMA),
	1067: uint16(anon_sym_AT),
	1068: uint16(anon_sym_LF),
	1069: uint16(9),
	1070: uint16(19),
	1071: uint16(1),
	1072: uint16(anon_sym_POUND),
	1073: uint16(27),
	1074: uint16(1),
	1075: uint16(aux_sym_number_token1),
	1076: uint16(35),
	1077: uint16(1),
	1078: uint16(anon_sym_0),
	1079: uint16(192),
	1080: uint16(1),
	1081: uint16(sym_hex),
	1082: uint16(30),
	1083: uint16(1),
	1084: uint16(sym__zero),
	1085: uint16(34),
	1086: uint16(1),
	1087: uint16(sym_comment),
	1088: uint16(67),
	1089: uint16(1),
	1090: uint16(sym_number),
	1091: uint16(131),
	1092: uint16(1),
	1093: uint16(sym_number_tuple),
	1094: uint16(184),
	1095: uint16(2),
	1096: uint16(anon_sym_PLUS),
	1097: uint16(anon_sym_DASH),
	1098: uint16(7),
	1099: uint16(19),
	1100: uint16(1),
	1101: uint16(anon_sym_POUND),
	1102: uint16(194),
	1103: uint16(1),
	1104: uint16(anon_sym_RBRACE),
	1105: uint16(196),
	1106: uint16(1),
	1107: uint16(sym_name),
	1108: uint16(198),
	1109: uint16(1),
	1110: uint16(anon_sym_LF),
	1111: uint16(35),
	1112: uint16(1),
	1113: uint16(sym_comment),
	1114: uint16(38),
	1115: uint16(1),
	1116: uint16(aux_sym_section_repeat1),
	1117: uint16(79),
	1118: uint16(4),
	1119: uint16(sym_assignment),
	1120: uint16(sym_keyword),
	1121: uint16(sym_section),
	1122: uint16(sym__linebreak),
	1123: uint16(7),
	1124: uint16(19),
	1125: uint16(1),
	1126: uint16(anon_sym_POUND),
	1127: uint16(196),
	1128: uint16(1),
	1129: uint16(sym_name),
	1130: uint16(198),
	1131: uint16(1),
	1132: uint16(anon_sym_LF),
	1133: uint16(200),
	1134: uint16(1),
	1135: uint16(anon_sym_RBRACE),
	1136: uint16(36),
	1137: uint16(1),
	1138: uint16(sym_comment),
	1139: uint16(37),
	1140: uint16(1),
	1141: uint16(aux_sym_section_repeat1),
	1142: uint16(79),
	1143: uint16(4),
	1144: uint16(sym_assignment),
	1145: uint16(sym_keyword),
	1146: uint16(sym_section),
	1147: uint16(sym__linebreak),
	1148: uint16(7),
	1149: uint16(19),
	1150: uint16(1),
	1151: uint16(anon_sym_POUND),
	1152: uint16(196),
	1153: uint16(1),
	1154: uint16(sym_name),
	1155: uint16(198),
	1156: uint16(1),
	1157: uint16(anon_sym_LF),
	1158: uint16(202),
	1159: uint16(1),
	1160: uint16(anon_sym_RBRACE),
	1161: uint16(37),
	1162: uint16(1),
	1163: uint16(sym_comment),
	1164: uint16(38),
	1165: uint16(1),
	1166: uint16(aux_sym_section_repeat1),
	1167: uint16(79),
	1168: uint16(4),
	1169: uint16(sym_assignment),
	1170: uint16(sym_keyword),
	1171: uint16(sym_section),
	1172: uint16(sym__linebreak),
	1173: uint16(6),
	1174: uint16(19),
	1175: uint16(1),
	1176: uint16(anon_sym_POUND),
	1177: uint16(204),
	1178: uint16(1),
	1179: uint16(anon_sym_RBRACE),
	1180: uint16(206),
	1181: uint16(1),
	1182: uint16(sym_name),
	1183: uint16(209),
	1184: uint16(1),
	1185: uint16(anon_sym_LF),
	1186: uint16(38),
	1187: uint16(2),
	1188: uint16(sym_comment),
	1189: uint16(aux_sym_section_repeat1),
	1190: uint16(79),
	1191: uint16(4),
	1192: uint16(sym_assignment),
	1193: uint16(sym_keyword),
	1194: uint16(sym_section),
	1195: uint16(sym__linebreak),
	1196: uint16(7),
	1197: uint16(19),
	1198: uint16(1),
	1199: uint16(anon_sym_POUND),
	1200: uint16(196),
	1201: uint16(1),
	1202: uint16(sym_name),
	1203: uint16(198),
	1204: uint16(1),
	1205: uint16(anon_sym_LF),
	1206: uint16(212),
	1207: uint16(1),
	1208: uint16(anon_sym_RBRACE),
	1209: uint16(35),
	1210: uint16(1),
	1211: uint16(aux_sym_section_repeat1),
	1212: uint16(39),
	1213: uint16(1),
	1214: uint16(sym_comment),
	1215: uint16(79),
	1216: uint16(4),
	1217: uint16(sym_assignment),
	1218: uint16(sym_keyword),
	1219: uint16(sym_section),
	1220: uint16(sym__linebreak),
	1221: uint16(7),
	1222: uint16(19),
	1223: uint16(1),
	1224: uint16(anon_sym_POUND),
	1225: uint16(196),
	1226: uint16(1),
	1227: uint16(sym_name),
	1228: uint16(198),
	1229: uint16(1),
	1230: uint16(anon_sym_LF),
	1231: uint16(214),
	1232: uint16(1),
	1233: uint16(anon_sym_RBRACE),
	1234: uint16(40),
	1235: uint16(1),
	1236: uint16(sym_comment),
	1237: uint16(41),
	1238: uint16(1),
	1239: uint16(aux_sym_section_repeat1),
	1240: uint16(79),
	1241: uint16(4),
	1242: uint16(sym_assignment),
	1243: uint16(sym_keyword),
	1244: uint16(sym_section),
	1245: uint16(sym__linebreak),
	1246: uint16(7),
	1247: uint16(19),
	1248: uint16(1),
	1249: uint16(anon_sym_POUND),
	1250: uint16(196),
	1251: uint16(1),
	1252: uint16(sym_name),
	1253: uint16(198),
	1254: uint16(1),
	1255: uint16(anon_sym_LF),
	1256: uint16(216),
	1257: uint16(1),
	1258: uint16(anon_sym_RBRACE),
	1259: uint16(38),
	1260: uint16(1),
	1261: uint16(aux_sym_section_repeat1),
	1262: uint16(41),
	1263: uint16(1),
	1264: uint16(sym_comment),
	1265: uint16(79),
	1266: uint16(4),
	1267: uint16(sym_assignment),
	1268: uint16(sym_keyword),
	1269: uint16(sym_section),
	1270: uint16(sym__linebreak),
	1271: uint16(7),
	1272: uint16(19),
	1273: uint16(1),
	1274: uint16(anon_sym_POUND),
	1275: uint16(196),
	1276: uint16(1),
	1277: uint16(sym_name),
	1278: uint16(198),
	1279: uint16(1),
	1280: uint16(anon_sym_LF),
	1281: uint16(218),
	1282: uint16(1),
	1283: uint16(anon_sym_RBRACE),
	1284: uint16(38),
	1285: uint16(1),
	1286: uint16(aux_sym_section_repeat1),
	1287: uint16(42),
	1288: uint16(1),
	1289: uint16(sym_comment),
	1290: uint16(79),
	1291: uint16(4),
	1292: uint16(sym_assignment),
	1293: uint16(sym_keyword),
	1294: uint16(sym_section),
	1295: uint16(sym__linebreak),
	1296: uint16(7),
	1297: uint16(19),
	1298: uint16(1),
	1299: uint16(anon_sym_POUND),
	1300: uint16(196),
	1301: uint16(1),
	1302: uint16(sym_name),
	1303: uint16(198),
	1304: uint16(1),
	1305: uint16(anon_sym_LF),
	1306: uint16(220),
	1307: uint16(1),
	1308: uint16(anon_sym_RBRACE),
	1309: uint16(42),
	1310: uint16(1),
	1311: uint16(aux_sym_section_repeat1),
	1312: uint16(43),
	1313: uint16(1),
	1314: uint16(sym_comment),
	1315: uint16(79),
	1316: uint16(4),
	1317: uint16(sym_assignment),
	1318: uint16(sym_keyword),
	1319: uint16(sym_section),
	1320: uint16(sym__linebreak),
	1321: uint16(5),
	1322: uint16(19),
	1323: uint16(1),
	1324: uint16(anon_sym_POUND),
	1325: uint16(222),
	1326: uint16(1),
	1327: uint16(anon_sym_x),
	1328: uint16(44),
	1329: uint16(1),
	1330: uint16(sym_comment),
	1331: uint16(178),
	1332: uint16(2),
	1333: uint16(aux_sym_number_token1),
	1334: uint16(anon_sym_0),
	1335: uint16(176),
	1336: uint16(4),
	1337: uint16(anon_sym_PLUS),
	1338: uint16(anon_sym_DASH),
	1339: uint16(anon_sym_COMMA),
	1340: uint16(anon_sym_LF),
	1341: uint16(4),
	1342: uint16(3),
	1343: uint16(1),
	1344: uint16(anon_sym_POUND),
	1345: uint16(45),
	1346: uint16(1),
	1347: uint16(sym_comment),
	1348: uint16(225),
	1349: uint16(2),
	1350: uint16(anon_sym_SEMI),
	1351: uint16(anon_sym_RBRACK),
	1352: uint16(227),
	1353: uint16(5),
	1354: uint16(anon_sym_PLUS),
	1355: uint16(anon_sym_DASH),
	1356: uint16(aux_sym_number_token1),
	1357: uint16(anon_sym_0),
	1358: uint16(sym__window_rule_argument),
	1359: uint16(4),
	1360: uint16(3),
	1361: uint16(1),
	1362: uint16(anon_sym_POUND),
	1363: uint16(46),
	1364: uint16(1),
	1365: uint16(sym_comment),
	1366: uint16(188),
	1367: uint16(2),
	1368: uint16(anon_sym_SEMI),
	1369: uint16(anon_sym_RBRACK),
	1370: uint16(190),
	1371: uint16(5),
	1372: uint16(anon_sym_PLUS),
	1373: uint16(anon_sym_DASH),
	1374: uint16(aux_sym_number_token1),
	1375: uint16(anon_sym_0),
	1376: uint16(sym__window_rule_argument),
	1377: uint16(4),
	1378: uint16(3),
	1379: uint16(1),
	1380: uint16(anon_sym_POUND),
	1381: uint16(47),
	1382: uint16(1),
	1383: uint16(sym_comment),
	1384: uint16(176),
	1385: uint16(2),
	1386: uint16(anon_sym_SEMI),
	1387: uint16(anon_sym_RBRACK),
	1388: uint16(178),
	1389: uint16(5),
	1390: uint16(anon_sym_PLUS),
	1391: uint16(anon_sym_DASH),
	1392: uint16(aux_sym_number_token1),
	1393: uint16(anon_sym_0),
	1394: uint16(sym__window_rule_argument),
	1395: uint16(4),
	1396: uint16(3),
	1397: uint16(1),
	1398: uint16(anon_sym_POUND),
	1399: uint16(48),
	1400: uint16(1),
	1401: uint16(sym_comment),
	1402: uint16(180),
	1403: uint16(2),
	1404: uint16(anon_sym_SEMI),
	1405: uint16(anon_sym_RBRACK),
	1406: uint16(182),
	1407: uint16(5),
	1408: uint16(anon_sym_PLUS),
	1409: uint16(anon_sym_DASH),
	1410: uint16(aux_sym_number_token1),
	1411: uint16(anon_sym_0),
	1412: uint16(sym__window_rule_argument),
	1413: uint16(4),
	1414: uint16(19),
	1415: uint16(1),
	1416: uint16(anon_sym_POUND),
	1417: uint16(49),
	1418: uint16(1),
	1419: uint16(sym_comment),
	1420: uint16(229),
	1421: uint16(3),
	1422: uint16(anon_sym_rgb),
	1423: uint16(aux_sym_angle_token1),
	1424: uint16(anon_sym_0),
	1425: uint16(231),
	1426: uint16(3),
	1427: uint16(anon_sym_rgba),
	1428: uint16(anon_sym_COMMA),
	1429: uint16(anon_sym_LF),
	1430: uint16(7),
	1431: uint16(19),
	1432: uint16(1),
	1433: uint16(anon_sym_POUND),
	1434: uint16(27),
	1435: uint16(1),
	1436: uint16(aux_sym_number_token1),
	1437: uint16(35),
	1438: uint16(1),
	1439: uint16(anon_sym_0),
	1440: uint16(30),
	1441: uint16(1),
	1442: uint16(sym__zero),
	1443: uint16(50),
	1444: uint16(1),
	1445: uint16(sym_comment),
	1446: uint16(70),
	1447: uint16(1),
	1448: uint16(sym_number),
	1449: uint16(184),
	1450: uint16(2),
	1451: uint16(anon_sym_PLUS),
	1452: uint16(anon_sym_DASH),
	1453: uint16(4),
	1454: uint16(19),
	1455: uint16(1),
	1456: uint16(anon_sym_POUND),
	1457: uint16(51),
	1458: uint16(1),
	1459: uint16(sym_comment),
	1460: uint16(233),
	1461: uint16(3),
	1462: uint16(anon_sym_rgb),
	1463: uint16(aux_sym_angle_token1),
	1464: uint16(anon_sym_0),
	1465: uint16(235),
	1466: uint16(3),
	1467: uint16(anon_sym_rgba),
	1468: uint16(anon_sym_COMMA),
	1469: uint16(anon_sym_LF),
	1470: uint16(7),
	1471: uint16(19),
	1472: uint16(1),
	1473: uint16(anon_sym_POUND),
	1474: uint16(27),
	1475: uint16(1),
	1476: uint16(aux_sym_number_token1),
	1477: uint16(35),
	1478: uint16(1),
	1479: uint16(anon_sym_0),
	1480: uint16(30),
	1481: uint16(1),
	1482: uint16(sym__zero),
	1483: uint16(52),
	1484: uint16(1),
	1485: uint16(sym_comment),
	1486: uint16(95),
	1487: uint16(1),
	1488: uint16(sym_number),
	1489: uint16(184),
	1490: uint16(2),
	1491: uint16(anon_sym_PLUS),
	1492: uint16(anon_sym_DASH),
	1493: uint16(7),
	1494: uint16(19),
	1495: uint16(1),
	1496: uint16(anon_sym_POUND),
	1497: uint16(27),
	1498: uint16(1),
	1499: uint16(aux_sym_number_token1),
	1500: uint16(35),
	1501: uint16(1),
	1502: uint16(anon_sym_0),
	1503: uint16(30),
	1504: uint16(1),
	1505: uint16(sym__zero),
	1506: uint16(53),
	1507: uint16(1),
	1508: uint16(sym_comment),
	1509: uint16(86),
	1510: uint16(1),
	1511: uint16(sym_number),
	1512: uint16(184),
	1513: uint16(2),
	1514: uint16(anon_sym_PLUS),
	1515: uint16(anon_sym_DASH),
	1516: uint16(4),
	1517: uint16(19),
	1518: uint16(1),
	1519: uint16(anon_sym_POUND),
	1520: uint16(54),
	1521: uint16(1),
	1522: uint16(sym_comment),
	1523: uint16(237),
	1524: uint16(3),
	1525: uint16(anon_sym_rgb),
	1526: uint16(aux_sym_angle_token1),
	1527: uint16(anon_sym_0),
	1528: uint16(239),
	1529: uint16(3),
	1530: uint16(anon_sym_rgba),
	1531: uint16(anon_sym_COMMA),
	1532: uint16(anon_sym_LF),
	1533: uint16(4),
	1534: uint16(19),
	1535: uint16(1),
	1536: uint16(anon_sym_POUND),
	1537: uint16(55),
	1538: uint16(1),
	1539: uint16(sym_comment),
	1540: uint16(241),
	1541: uint16(3),
	1542: uint16(anon_sym_rgb),
	1543: uint16(aux_sym_angle_token1),
	1544: uint16(anon_sym_0),
	1545: uint16(243),
	1546: uint16(3),
	1547: uint16(anon_sym_rgba),
	1548: uint16(anon_sym_COMMA),
	1549: uint16(anon_sym_LF),
	1550: uint16(6),
	1551: uint16(17),
	1552: uint16(1),
	1553: uint16(anon_sym_LF),
	1554: uint16(19),
	1555: uint16(1),
	1556: uint16(anon_sym_POUND),
	1557: uint16(245),
	1558: uint16(1),
	1559: uint16(anon_sym_COMMA),
	1560: uint16(19),
	1561: uint16(1),
	1562: uint16(sym__linebreak),
	1563: uint16(56),
	1564: uint16(1),
	1565: uint16(sym_comment),
	1566: uint16(63),
	1567: uint16(1),
	1568: uint16(aux_sym_params_repeat1),
	1569: uint16(6),
	1570: uint16(19),
	1571: uint16(1),
	1572: uint16(anon_sym_POUND),
	1573: uint16(198),
	1574: uint16(1),
	1575: uint16(anon_sym_LF),
	1576: uint16(245),
	1577: uint16(1),
	1578: uint16(anon_sym_COMMA),
	1579: uint16(57),
	1580: uint16(1),
	1581: uint16(sym_comment),
	1582: uint16(63),
	1583: uint16(1),
	1584: uint16(aux_sym_params_repeat1),
	1585: uint16(73),
	1586: uint16(1),
	1587: uint16(sym__linebreak),
	1588: uint16(4),
	1589: uint16(19),
	1590: uint16(1),
	1591: uint16(anon_sym_POUND),
	1592: uint16(247),
	1593: uint16(1),
	1594: uint16(anon_sym_SEMI),
	1595: uint16(250),
	1596: uint16(1),
	1597: uint16(anon_sym_RBRACK),
	1598: uint16(58),
	1599: uint16(2),
	1600: uint16(sym_comment),
	1601: uint16(aux_sym_rules_repeat1),
	1602: uint16(5),
	1603: uint16(19),
	1604: uint16(1),
	1605: uint16(anon_sym_POUND),
	1606: uint16(252),
	1607: uint16(1),
	1608: uint16(anon_sym_RPAREN),
	1609: uint16(254),
	1610: uint16(1),
	1611: uint16(anon_sym_COMMA),
	1612: uint16(59),
	1613: uint16(1),
	1614: uint16(sym_comment),
	1615: uint16(62),
	1616: uint16(1),
	1617: uint16(aux_sym_number_tuple_repeat1),
	1618: uint16(5),
	1619: uint16(3),
	1620: uint16(1),
	1621: uint16(anon_sym_POUND),
	1622: uint16(256),
	1623: uint16(1),
	1624: uint16(sym_string),
	1625: uint16(258),
	1626: uint16(1),
	1627: uint16(anon_sym_LBRACK),
	1628: uint16(60),
	1629: uint16(1),
	1630: uint16(sym_comment),
	1631: uint16(133),
	1632: uint16(1),
	1633: uint16(sym_rules),
	1634: uint16(5),
	1635: uint16(19),
	1636: uint16(1),
	1637: uint16(anon_sym_POUND),
	1638: uint16(260),
	1639: uint16(1),
	1640: uint16(anon_sym_SEMI),
	1641: uint16(262),
	1642: uint16(1),
	1643: uint16(anon_sym_RBRACK),
	1644: uint16(61),
	1645: uint16(1),
	1646: uint16(sym_comment),
	1647: uint16(64),
	1648: uint16(1),
	1649: uint16(aux_sym_rules_repeat1),
	1650: uint16(4),
	1651: uint16(19),
	1652: uint16(1),
	1653: uint16(anon_sym_POUND),
	1654: uint16(264),
	1655: uint16(1),
	1656: uint16(anon_sym_RPAREN),
	1657: uint16(266),
	1658: uint16(1),
	1659: uint16(anon_sym_COMMA),
	1660: uint16(62),
	1661: uint16(2),
	1662: uint16(sym_comment),
	1663: uint16(aux_sym_number_tuple_repeat1),
	1664: uint16(5),
	1665: uint16(19),
	1666: uint16(1),
	1667: uint16(anon_sym_POUND),
	1668: uint16(245),
	1669: uint16(1),
	1670: uint16(anon_sym_COMMA),
	1671: uint16(269),
	1672: uint16(1),
	1673: uint16(anon_sym_LF),
	1674: uint16(63),
	1675: uint16(1),
	1676: uint16(sym_comment),
	1677: uint16(68),
	1678: uint16(1),
	1679: uint16(aux_sym_params_repeat1),
	1680: uint16(5),
	1681: uint16(19),
	1682: uint16(1),
	1683: uint16(anon_sym_POUND),
	1684: uint16(260),
	1685: uint16(1),
	1686: uint16(anon_sym_SEMI),
	1687: uint16(271),
	1688: uint16(1),
	1689: uint16(anon_sym_RBRACK),
	1690: uint16(58),
	1691: uint16(1),
	1692: uint16(aux_sym_rules_repeat1),
	1693: uint16(64),
	1694: uint16(1),
	1695: uint16(sym_comment),
	1696: uint16(3),
	1697: uint16(19),
	1698: uint16(1),
	1699: uint16(anon_sym_POUND),
	1700: uint16(65),
	1701: uint16(1),
	1702: uint16(sym_comment),
	1703: uint16(108),
	1704: uint16(3),
	1705: uint16(anon_sym_RBRACE),
	1706: uint16(sym_name),
	1707: uint16(anon_sym_LF),
	1708: uint16(5),
	1709: uint16(19),
	1710: uint16(1),
	1711: uint16(anon_sym_POUND),
	1712: uint16(273),
	1713: uint16(1),
	1714: uint16(anon_sym_EQ),
	1715: uint16(275),
	1716: uint16(1),
	1717: uint16(anon_sym_COLON),
	1718: uint16(277),
	1719: uint16(1),
	1720: uint16(anon_sym_LBRACE),
	1721: uint16(66),
	1722: uint16(1),
	1723: uint16(sym_comment),
	1724: uint16(5),
	1725: uint16(19),
	1726: uint16(1),
	1727: uint16(anon_sym_POUND),
	1728: uint16(254),
	1729: uint16(1),
	1730: uint16(anon_sym_COMMA),
	1731: uint16(279),
	1732: uint16(1),
	1733: uint16(anon_sym_RPAREN),
	1734: uint16(59),
	1735: uint16(1),
	1736: uint16(aux_sym_number_tuple_repeat1),
	1737: uint16(67),
	1738: uint16(1),
	1739: uint16(sym_comment),
	1740: uint16(4),
	1741: uint16(19),
	1742: uint16(1),
	1743: uint16(anon_sym_POUND),
	1744: uint16(281),
	1745: uint16(1),
	1746: uint16(anon_sym_COMMA),
	1747: uint16(284),
	1748: uint16(1),
	1749: uint16(anon_sym_LF),
	1750: uint16(68),
	1751: uint16(2),
	1752: uint16(sym_comment),
	1753: uint16(aux_sym_params_repeat1),
	1754: uint16(3),
	1755: uint16(19),
	1756: uint16(1),
	1757: uint16(anon_sym_POUND),
	1758: uint16(69),
	1759: uint16(1),
	1760: uint16(sym_comment),
	1761: uint16(128),
	1762: uint16(3),
	1763: uint16(anon_sym_RBRACE),
	1764: uint16(sym_name),
	1765: uint16(anon_sym_LF),
	1766: uint16(3),
	1767: uint16(19),
	1768: uint16(1),
	1769: uint16(anon_sym_POUND),
	1770: uint16(70),
	1771: uint16(1),
	1772: uint16(sym_comment),
	1773: uint16(286),
	1774: uint16(3),
	1775: uint16(anon_sym_COMMA),
	1776: uint16(anon_sym_AT),
	1777: uint16(anon_sym_LF),
	1778: uint16(3),
	1779: uint16(19),
	1780: uint16(1),
	1781: uint16(anon_sym_POUND),
	1782: uint16(71),
	1783: uint16(1),
	1784: uint16(sym_comment),
	1785: uint16(164),
	1786: uint16(3),
	1787: uint16(anon_sym_RBRACE),
	1788: uint16(sym_name),
	1789: uint16(anon_sym_LF),
	1790: uint16(3),
	1791: uint16(19),
	1792: uint16(1),
	1793: uint16(anon_sym_POUND),
	1794: uint16(72),
	1795: uint16(1),
	1796: uint16(sym_comment),
	1797: uint16(168),
	1798: uint16(3),
	1799: uint16(anon_sym_RBRACE),
	1800: uint16(sym_name),
	1801: uint16(anon_sym_LF),
	1802: uint16(3),
	1803: uint16(19),
	1804: uint16(1),
	1805: uint16(anon_sym_POUND),
	1806: uint16(73),
	1807: uint16(1),
	1808: uint16(sym_comment),
	1809: uint16(124),
	1810: uint16(3),
	1811: uint16(anon_sym_RBRACE),
	1812: uint16(sym_name),
	1813: uint16(anon_sym_LF),
	1814: uint16(5),
	1815: uint16(19),
	1816: uint16(1),
	1817: uint16(anon_sym_POUND),
	1818: uint16(288),
	1819: uint16(1),
	1820: uint16(anon_sym_EQ),
	1821: uint16(290),
	1822: uint16(1),
	1823: uint16(anon_sym_COLON),
	1824: uint16(292),
	1825: uint16(1),
	1826: uint16(anon_sym_LBRACE),
	1827: uint16(74),
	1828: uint16(1),
	1829: uint16(sym_comment),
	1830: uint16(3),
	1831: uint16(19),
	1832: uint16(1),
	1833: uint16(anon_sym_POUND),
	1834: uint16(75),
	1835: uint16(1),
	1836: uint16(sym_comment),
	1837: uint16(142),
	1838: uint16(3),
	1839: uint16(anon_sym_RBRACE),
	1840: uint16(sym_name),
	1841: uint16(anon_sym_LF),
	1842: uint16(3),
	1843: uint16(19),
	1844: uint16(1),
	1845: uint16(anon_sym_POUND),
	1846: uint16(76),
	1847: uint16(1),
	1848: uint16(sym_comment),
	1849: uint16(172),
	1850: uint16(3),
	1851: uint16(anon_sym_RBRACE),
	1852: uint16(sym_name),
	1853: uint16(anon_sym_LF),
	1854: uint16(3),
	1855: uint16(19),
	1856: uint16(1),
	1857: uint16(anon_sym_POUND),
	1858: uint16(77),
	1859: uint16(1),
	1860: uint16(sym_comment),
	1861: uint16(146),
	1862: uint16(3),
	1863: uint16(anon_sym_RBRACE),
	1864: uint16(sym_name),
	1865: uint16(anon_sym_LF),
	1866: uint16(4),
	1867: uint16(19),
	1868: uint16(1),
	1869: uint16(anon_sym_POUND),
	1870: uint16(294),
	1871: uint16(1),
	1872: uint16(anon_sym_AT),
	1873: uint16(78),
	1874: uint16(1),
	1875: uint16(sym_comment),
	1876: uint16(71),
	1877: uint16(2),
	1878: uint16(anon_sym_COMMA),
	1879: uint16(anon_sym_LF),
	1880: uint16(3),
	1881: uint16(19),
	1882: uint16(1),
	1883: uint16(anon_sym_POUND),
	1884: uint16(79),
	1885: uint16(1),
	1886: uint16(sym_comment),
	1887: uint16(296),
	1888: uint16(3),
	1889: uint16(anon_sym_RBRACE),
	1890: uint16(sym_name),
	1891: uint16(anon_sym_LF),
	1892: uint16(4),
	1893: uint16(19),
	1894: uint16(1),
	1895: uint16(anon_sym_POUND),
	1896: uint16(198),
	1897: uint16(1),
	1898: uint16(anon_sym_LF),
	1899: uint16(40),
	1900: uint16(1),
	1901: uint16(sym__linebreak),
	1902: uint16(80),
	1903: uint16(1),
	1904: uint16(sym_comment),
	1905: uint16(4),
	1906: uint16(17),
	1907: uint16(1),
	1908: uint16(anon_sym_LF),
	1909: uint16(19),
	1910: uint16(1),
	1911: uint16(anon_sym_POUND),
	1912: uint16(21),
	1913: uint16(1),
	1914: uint16(sym__linebreak),
	1915: uint16(81),
	1916: uint16(1),
	1917: uint16(sym_comment),
	1918: uint16(4),
	1919: uint16(17),
	1920: uint16(1),
	1921: uint16(anon_sym_LF),
	1922: uint16(19),
	1923: uint16(1),
	1924: uint16(anon_sym_POUND),
	1925: uint16(16),
	1926: uint16(1),
	1927: uint16(sym__linebreak),
	1928: uint16(82),
	1929: uint16(1),
	1930: uint16(sym_comment),
	1931: uint16(3),
	1932: uint16(19),
	1933: uint16(1),
	1934: uint16(anon_sym_POUND),
	1935: uint16(83),
	1936: uint16(1),
	1937: uint16(sym_comment),
	1938: uint16(298),
	1939: uint16(2),
	1940: uint16(anon_sym_COMMA),
	1941: uint16(anon_sym_LF),
	1942: uint16(3),
	1943: uint16(19),
	1944: uint16(1),
	1945: uint16(anon_sym_POUND),
	1946: uint16(84),
	1947: uint16(1),
	1948: uint16(sym_comment),
	1949: uint16(300),
	1950: uint16(2),
	1951: uint16(anon_sym_COMMA),
	1952: uint16(anon_sym_LF),
	1953: uint16(4),
	1954: uint16(19),
	1955: uint16(1),
	1956: uint16(anon_sym_POUND),
	1957: uint16(302),
	1958: uint16(1),
	1959: uint16(sym_name),
	1960: uint16(61),
	1961: uint16(1),
	1962: uint16(sym_window_rule),
	1963: uint16(85),
	1964: uint16(1),
	1965: uint16(sym_comment),
	1966: uint16(3),
	1967: uint16(19),
	1968: uint16(1),
	1969: uint16(anon_sym_POUND),
	1970: uint16(86),
	1971: uint16(1),
	1972: uint16(sym_comment),
	1973: uint16(304),
	1974: uint16(2),
	1975: uint16(anon_sym_COMMA),
	1976: uint16(anon_sym_LF),
	1977: uint16(4),
	1978: uint16(17),
	1979: uint16(1),
	1980: uint16(anon_sym_LF),
	1981: uint16(19),
	1982: uint16(1),
	1983: uint16(anon_sym_POUND),
	1984: uint16(17),
	1985: uint16(1),
	1986: uint16(sym__linebreak),
	1987: uint16(87),
	1988: uint16(1),
	1989: uint16(sym_comment),
	1990: uint16(4),
	1991: uint16(17),
	1992: uint16(1),
	1993: uint16(anon_sym_LF),
	1994: uint16(19),
	1995: uint16(1),
	1996: uint16(anon_sym_POUND),
	1997: uint16(23),
	1998: uint16(1),
	1999: uint16(sym__linebreak),
	2000: uint16(88),
	2001: uint16(1),
	2002: uint16(sym_comment),
	2003: uint16(4),
	2004: uint16(17),
	2005: uint16(1),
	2006: uint16(anon_sym_LF),
	2007: uint16(19),
	2008: uint16(1),
	2009: uint16(anon_sym_POUND),
	2010: uint16(27),
	2011: uint16(1),
	2012: uint16(sym__linebreak),
	2013: uint16(89),
	2014: uint16(1),
	2015: uint16(sym_comment),
	2016: uint16(3),
	2017: uint16(19),
	2018: uint16(1),
	2019: uint16(anon_sym_POUND),
	2020: uint16(90),
	2021: uint16(1),
	2022: uint16(sym_comment),
	2023: uint16(306),
	2024: uint16(2),
	2025: uint16(anon_sym_COMMA),
	2026: uint16(anon_sym_LF),
	2027: uint16(3),
	2028: uint16(19),
	2029: uint16(1),
	2030: uint16(anon_sym_POUND),
	2031: uint16(91),
	2032: uint16(1),
	2033: uint16(sym_comment),
	2034: uint16(250),
	2035: uint16(2),
	2036: uint16(anon_sym_SEMI),
	2037: uint16(anon_sym_RBRACK),
	2038: uint16(3),
	2039: uint16(19),
	2040: uint16(1),
	2041: uint16(anon_sym_POUND),
	2042: uint16(92),
	2043: uint16(1),
	2044: uint16(sym_comment),
	2045: uint16(71),
	2046: uint16(2),
	2047: uint16(anon_sym_COMMA),
	2048: uint16(anon_sym_LF),
	2049: uint16(4),
	2050: uint16(17),
	2051: uint16(1),
	2052: uint16(anon_sym_LF),
	2053: uint16(19),
	2054: uint16(1),
	2055: uint16(anon_sym_POUND),
	2056: uint16(28),
	2057: uint16(1),
	2058: uint16(sym__linebreak),
	2059: uint16(93),
	2060: uint16(1),
	2061: uint16(sym_comment),
	2062: uint16(3),
	2063: uint16(19),
	2064: uint16(1),
	2065: uint16(anon_sym_POUND),
	2066: uint16(94),
	2067: uint16(1),
	2068: uint16(sym_comment),
	2069: uint16(308),
	2070: uint16(2),
	2071: uint16(anon_sym_COMMA),
	2072: uint16(anon_sym_LF),
	2073: uint16(3),
	2074: uint16(19),
	2075: uint16(1),
	2076: uint16(anon_sym_POUND),
	2077: uint16(95),
	2078: uint16(1),
	2079: uint16(sym_comment),
	2080: uint16(264),
	2081: uint16(2),
	2082: uint16(anon_sym_RPAREN),
	2083: uint16(anon_sym_COMMA),
	2084: uint16(3),
	2085: uint16(19),
	2086: uint16(1),
	2087: uint16(anon_sym_POUND),
	2088: uint16(96),
	2089: uint16(1),
	2090: uint16(sym_comment),
	2091: uint16(81),
	2092: uint16(2),
	2093: uint16(anon_sym_COMMA),
	2094: uint16(anon_sym_LF),
	2095: uint16(4),
	2096: uint16(17),
	2097: uint16(1),
	2098: uint16(anon_sym_LF),
	2099: uint16(19),
	2100: uint16(1),
	2101: uint16(anon_sym_POUND),
	2102: uint16(29),
	2103: uint16(1),
	2104: uint16(sym__linebreak),
	2105: uint16(97),
	2106: uint16(1),
	2107: uint16(sym_comment),
	2108: uint16(3),
	2109: uint16(19),
	2110: uint16(1),
	2111: uint16(anon_sym_POUND),
	2112: uint16(98),
	2113: uint16(1),
	2114: uint16(sym_comment),
	2115: uint16(310),
	2116: uint16(2),
	2117: uint16(anon_sym_COMMA),
	2118: uint16(anon_sym_LF),
	2119: uint16(4),
	2120: uint16(19),
	2121: uint16(1),
	2122: uint16(anon_sym_POUND),
	2123: uint16(198),
	2124: uint16(1),
	2125: uint16(anon_sym_LF),
	2126: uint16(69),
	2127: uint16(1),
	2128: uint16(sym__linebreak),
	2129: uint16(99),
	2130: uint16(1),
	2131: uint16(sym_comment),
	2132: uint16(4),
	2133: uint16(19),
	2134: uint16(1),
	2135: uint16(anon_sym_POUND),
	2136: uint16(176),
	2137: uint16(1),
	2138: uint16(anon_sym_LF),
	2139: uint16(312),
	2140: uint16(1),
	2141: uint16(anon_sym_x),
	2142: uint16(100),
	2143: uint16(1),
	2144: uint16(sym_comment),
	2145: uint16(4),
	2146: uint16(19),
	2147: uint16(1),
	2148: uint16(anon_sym_POUND),
	2149: uint16(198),
	2150: uint16(1),
	2151: uint16(anon_sym_LF),
	2152: uint16(36),
	2153: uint16(1),
	2154: uint16(sym__linebreak),
	2155: uint16(101),
	2156: uint16(1),
	2157: uint16(sym_comment),
	2158: uint16(4),
	2159: uint16(17),
	2160: uint16(1),
	2161: uint16(anon_sym_LF),
	2162: uint16(19),
	2163: uint16(1),
	2164: uint16(anon_sym_POUND),
	2165: uint16(25),
	2166: uint16(1),
	2167: uint16(sym__linebreak),
	2168: uint16(102),
	2169: uint16(1),
	2170: uint16(sym_comment),
	2171: uint16(4),
	2172: uint16(19),
	2173: uint16(1),
	2174: uint16(anon_sym_POUND),
	2175: uint16(198),
	2176: uint16(1),
	2177: uint16(anon_sym_LF),
	2178: uint16(77),
	2179: uint16(1),
	2180: uint16(sym__linebreak),
	2181: uint16(103),
	2182: uint16(1),
	2183: uint16(sym_comment),
	2184: uint16(4),
	2185: uint16(19),
	2186: uint16(1),
	2187: uint16(anon_sym_POUND),
	2188: uint16(198),
	2189: uint16(1),
	2190: uint16(anon_sym_LF),
	2191: uint16(39),
	2192: uint16(1),
	2193: uint16(sym__linebreak),
	2194: uint16(104),
	2195: uint16(1),
	2196: uint16(sym_comment),
	2197: uint16(4),
	2198: uint16(19),
	2199: uint16(1),
	2200: uint16(anon_sym_POUND),
	2201: uint16(198),
	2202: uint16(1),
	2203: uint16(anon_sym_LF),
	2204: uint16(71),
	2205: uint16(1),
	2206: uint16(sym__linebreak),
	2207: uint16(105),
	2208: uint16(1),
	2209: uint16(sym_comment),
	2210: uint16(4),
	2211: uint16(19),
	2212: uint16(1),
	2213: uint16(anon_sym_POUND),
	2214: uint16(198),
	2215: uint16(1),
	2216: uint16(anon_sym_LF),
	2217: uint16(72),
	2218: uint16(1),
	2219: uint16(sym__linebreak),
	2220: uint16(106),
	2221: uint16(1),
	2222: uint16(sym_comment),
	2223: uint16(4),
	2224: uint16(17),
	2225: uint16(1),
	2226: uint16(anon_sym_LF),
	2227: uint16(19),
	2228: uint16(1),
	2229: uint16(anon_sym_POUND),
	2230: uint16(20),
	2231: uint16(1),
	2232: uint16(sym__linebreak),
	2233: uint16(107),
	2234: uint16(1),
	2235: uint16(sym_comment),
	2236: uint16(3),
	2237: uint16(19),
	2238: uint16(1),
	2239: uint16(anon_sym_POUND),
	2240: uint16(108),
	2241: uint16(1),
	2242: uint16(sym_comment),
	2243: uint16(314),
	2244: uint16(2),
	2245: uint16(anon_sym_SEMI),
	2246: uint16(anon_sym_RBRACK),
	2247: uint16(4),
	2248: uint16(19),
	2249: uint16(1),
	2250: uint16(anon_sym_POUND),
	2251: uint16(198),
	2252: uint16(1),
	2253: uint16(anon_sym_LF),
	2254: uint16(76),
	2255: uint16(1),
	2256: uint16(sym__linebreak),
	2257: uint16(109),
	2258: uint16(1),
	2259: uint16(sym_comment),
	2260: uint16(4),
	2261: uint16(19),
	2262: uint16(1),
	2263: uint16(anon_sym_POUND),
	2264: uint16(302),
	2265: uint16(1),
	2266: uint16(sym_name),
	2267: uint16(91),
	2268: uint16(1),
	2269: uint16(sym_window_rule),
	2270: uint16(110),
	2271: uint16(1),
	2272: uint16(sym_comment),
	2273: uint16(3),
	2274: uint16(19),
	2275: uint16(1),
	2276: uint16(anon_sym_POUND),
	2277: uint16(111),
	2278: uint16(1),
	2279: uint16(sym_comment),
	2280: uint16(284),
	2281: uint16(2),
	2282: uint16(anon_sym_COMMA),
	2283: uint16(anon_sym_LF),
	2284: uint16(4),
	2285: uint16(19),
	2286: uint16(1),
	2287: uint16(anon_sym_POUND),
	2288: uint16(198),
	2289: uint16(1),
	2290: uint16(anon_sym_LF),
	2291: uint16(43),
	2292: uint16(1),
	2293: uint16(sym__linebreak),
	2294: uint16(112),
	2295: uint16(1),
	2296: uint16(sym_comment),
	2297: uint16(3),
	2298: uint16(19),
	2299: uint16(1),
	2300: uint16(anon_sym_POUND),
	2301: uint16(316),
	2302: uint16(1),
	2304: uint16(113),
	2305: uint16(1),
	2306: uint16(sym_comment),
	2307: uint16(3),
	2308: uint16(3),
	2309: uint16(1),
	2310: uint16(anon_sym_POUND),
	2311: uint16(318),
	2312: uint16(1),
	2313: uint16(sym_string),
	2314: uint16(114),
	2315: uint16(1),
	2316: uint16(sym_comment),
	2317: uint16(3),
	2318: uint16(19),
	2319: uint16(1),
	2320: uint16(anon_sym_POUND),
	2321: uint16(320),
	2322: uint16(1),
	2323: uint16(anon_sym_deg),
	2324: uint16(115),
	2325: uint16(1),
	2326: uint16(sym_comment),
	2327: uint16(3),
	2328: uint16(19),
	2329: uint16(1),
	2330: uint16(anon_sym_POUND),
	2331: uint16(322),
	2332: uint16(1),
	2333: uint16(anon_sym_EQ),
	2334: uint16(116),
	2335: uint16(1),
	2336: uint16(sym_comment),
	2337: uint16(3),
	2338: uint16(19),
	2339: uint16(1),
	2340: uint16(anon_sym_POUND),
	2341: uint16(324),
	2342: uint16(1),
	2343: uint16(aux_sym_variable_token1),
	2344: uint16(117),
	2345: uint16(1),
	2346: uint16(sym_comment),
	2347: uint16(3),
	2348: uint16(19),
	2349: uint16(1),
	2350: uint16(anon_sym_POUND),
	2351: uint16(326),
	2352: uint16(1),
	2353: uint16(sym_device_name),
	2354: uint16(118),
	2355: uint16(1),
	2356: uint16(sym_comment),
	2357: uint16(3),
	2358: uint16(19),
	2359: uint16(1),
	2360: uint16(anon_sym_POUND),
	2361: uint16(328),
	2362: uint16(1),
	2363: uint16(anon_sym_LPAREN),
	2364: uint16(119),
	2365: uint16(1),
	2366: uint16(sym_comment),
	2367: uint16(3),
	2368: uint16(19),
	2369: uint16(1),
	2370: uint16(anon_sym_POUND),
	2371: uint16(312),
	2372: uint16(1),
	2373: uint16(anon_sym_x),
	2374: uint16(120),
	2375: uint16(1),
	2376: uint16(sym_comment),
	2377: uint16(3),
	2378: uint16(19),
	2379: uint16(1),
	2380: uint16(anon_sym_POUND),
	2381: uint16(330),
	2382: uint16(1),
	2383: uint16(anon_sym_EQ),
	2384: uint16(121),
	2385: uint16(1),
	2386: uint16(sym_comment),
	2387: uint16(3),
	2388: uint16(3),
	2389: uint16(1),
	2390: uint16(anon_sym_POUND),
	2391: uint16(332),
	2392: uint16(1),
	2393: uint16(sym_string),
	2394: uint16(122),
	2395: uint16(1),
	2396: uint16(sym_comment),
	2397: uint16(3),
	2398: uint16(3),
	2399: uint16(1),
	2400: uint16(anon_sym_POUND),
	2401: uint16(334),
	2402: uint16(1),
	2403: uint16(aux_sym_comment_token1),
	2404: uint16(123),
	2405: uint16(1),
	2406: uint16(sym_comment),
	2407: uint16(3),
	2408: uint16(19),
	2409: uint16(1),
	2410: uint16(anon_sym_POUND),
	2411: uint16(336),
	2412: uint16(1),
	2413: uint16(aux_sym_number_token1),
	2414: uint16(124),
	2415: uint16(1),
	2416: uint16(sym_comment),
	2417: uint16(3),
	2418: uint16(19),
	2419: uint16(1),
	2420: uint16(anon_sym_POUND),
	2421: uint16(338),
	2422: uint16(1),
	2423: uint16(anon_sym_EQ),
	2424: uint16(125),
	2425: uint16(1),
	2426: uint16(sym_comment),
	2427: uint16(3),
	2428: uint16(3),
	2429: uint16(1),
	2430: uint16(anon_sym_POUND),
	2431: uint16(256),
	2432: uint16(1),
	2433: uint16(sym_string),
	2434: uint16(126),
	2435: uint16(1),
	2436: uint16(sym_comment),
	2437: uint16(3),
	2438: uint16(19),
	2439: uint16(1),
	2440: uint16(anon_sym_POUND),
	2441: uint16(340),
	2442: uint16(1),
	2443: uint16(aux_sym_number_token1),
	2444: uint16(127),
	2445: uint16(1),
	2446: uint16(sym_comment),
	2447: uint16(3),
	2448: uint16(3),
	2449: uint16(1),
	2450: uint16(anon_sym_POUND),
	2451: uint16(342),
	2452: uint16(1),
	2453: uint16(sym_string),
	2454: uint16(128),
	2455: uint16(1),
	2456: uint16(sym_comment),
	2457: uint16(3),
	2458: uint16(19),
	2459: uint16(1),
	2460: uint16(anon_sym_POUND),
	2461: uint16(344),
	2462: uint16(1),
	2463: uint16(sym_hex),
	2464: uint16(129),
	2465: uint16(1),
	2466: uint16(sym_comment),
	2467: uint16(3),
	2468: uint16(19),
	2469: uint16(1),
	2470: uint16(anon_sym_POUND),
	2471: uint16(346),
	2472: uint16(1),
	2473: uint16(anon_sym_LBRACE),
	2474: uint16(130),
	2475: uint16(1),
	2476: uint16(sym_comment),
	2477: uint16(3),
	2478: uint16(19),
	2479: uint16(1),
	2480: uint16(anon_sym_POUND),
	2481: uint16(348),
	2482: uint16(1),
	2483: uint16(anon_sym_RPAREN),
	2484: uint16(131),
	2485: uint16(1),
	2486: uint16(sym_comment),
	2487: uint16(3),
	2488: uint16(19),
	2489: uint16(1),
	2490: uint16(anon_sym_POUND),
	2491: uint16(350),
	2492: uint16(1),
	2493: uint16(anon_sym_EQ),
	2494: uint16(132),
	2495: uint16(1),
	2496: uint16(sym_comment),
	2497: uint16(3),
	2498: uint16(3),
	2499: uint16(1),
	2500: uint16(anon_sym_POUND),
	2501: uint16(352),
	2502: uint16(1),
	2503: uint16(sym_string),
	2504: uint16(133),
	2505: uint16(1),
	2506: uint16(sym_comment),
	2507: uint16(3),
	2508: uint16(19),
	2509: uint16(1),
	2510: uint16(anon_sym_POUND),
	2511: uint16(354),
	2512: uint16(1),
	2513: uint16(anon_sym_LBRACE),
	2514: uint16(134),
	2515: uint16(1),
	2516: uint16(sym_comment),
	2517: uint16(3),
	2518: uint16(19),
	2519: uint16(1),
	2520: uint16(anon_sym_POUND),
	2521: uint16(356),
	2522: uint16(1),
	2523: uint16(sym_device_name),
	2524: uint16(135),
	2525: uint16(1),
	2526: uint16(sym_comment),
	2527: uint16(1),
	2528: uint16(358),
	2529: uint16(1),
}

var ts_small_parse_table_map = [135]uint32_t{
	1:   uint32(90),
	2:   uint32(180),
	3:   uint32(265),
	4:   uint32(317),
	5:   uint32(363),
	6:   uint32(407),
	7:   uint32(435),
	8:   uint32(467),
	9:   uint32(494),
	10:  uint32(533),
	11:  uint32(572),
	12:  uint32(608),
	13:  uint32(642),
	14:  uint32(663),
	15:  uint32(684),
	16:  uint32(705),
	17:  uint32(726),
	18:  uint32(747),
	19:  uint32(768),
	20:  uint32(789),
	21:  uint32(822),
	22:  uint32(843),
	23:  uint32(864),
	24:  uint32(885),
	25:  uint32(916),
	26:  uint32(937),
	27:  uint32(958),
	28:  uint32(979),
	29:  uint32(999),
	30:  uint32(1019),
	31:  uint32(1049),
	32:  uint32(1069),
	33:  uint32(1098),
	34:  uint32(1123),
	35:  uint32(1148),
	36:  uint32(1173),
	37:  uint32(1196),
	38:  uint32(1221),
	39:  uint32(1246),
	40:  uint32(1271),
	41:  uint32(1296),
	42:  uint32(1321),
	43:  uint32(1341),
	44:  uint32(1359),
	45:  uint32(1377),
	46:  uint32(1395),
	47:  uint32(1413),
	48:  uint32(1430),
	49:  uint32(1453),
	50:  uint32(1470),
	51:  uint32(1493),
	52:  uint32(1516),
	53:  uint32(1533),
	54:  uint32(1550),
	55:  uint32(1569),
	56:  uint32(1588),
	57:  uint32(1602),
	58:  uint32(1618),
	59:  uint32(1634),
	60:  uint32(1650),
	61:  uint32(1664),
	62:  uint32(1680),
	63:  uint32(1696),
	64:  uint32(1708),
	65:  uint32(1724),
	66:  uint32(1740),
	67:  uint32(1754),
	68:  uint32(1766),
	69:  uint32(1778),
	70:  uint32(1790),
	71:  uint32(1802),
	72:  uint32(1814),
	73:  uint32(1830),
	74:  uint32(1842),
	75:  uint32(1854),
	76:  uint32(1866),
	77:  uint32(1880),
	78:  uint32(1892),
	79:  uint32(1905),
	80:  uint32(1918),
	81:  uint32(1931),
	82:  uint32(1942),
	83:  uint32(1953),
	84:  uint32(1966),
	85:  uint32(1977),
	86:  uint32(1990),
	87:  uint32(2003),
	88:  uint32(2016),
	89:  uint32(2027),
	90:  uint32(2038),
	91:  uint32(2049),
	92:  uint32(2062),
	93:  uint32(2073),
	94:  uint32(2084),
	95:  uint32(2095),
	96:  uint32(2108),
	97:  uint32(2119),
	98:  uint32(2132),
	99:  uint32(2145),
	100: uint32(2158),
	101: uint32(2171),
	102: uint32(2184),
	103: uint32(2197),
	104: uint32(2210),
	105: uint32(2223),
	106: uint32(2236),
	107: uint32(2247),
	108: uint32(2260),
	109: uint32(2273),
	110: uint32(2284),
	111: uint32(2297),
	112: uint32(2307),
	113: uint32(2317),
	114: uint32(2327),
	115: uint32(2337),
	116: uint32(2347),
	117: uint32(2357),
	118: uint32(2367),
	119: uint32(2377),
	120: uint32(2387),
	121: uint32(2397),
	122: uint32(2407),
	123: uint32(2417),
	124: uint32(2427),
	125: uint32(2437),
	126: uint32(2447),
	127: uint32(2457),
	128: uint32(2467),
	129: uint32(2477),
	130: uint32(2487),
	131: uint32(2497),
	132: uint32(2507),
	133: uint32(2517),
	134: uint32(2527),
}

var ts_parse_actions = [360]TSParseActionEntry{
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
		Fstate: uint16(libc.Int32FromInt32(123)),
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
		Fsymbol:     uint16(sym_configuration),
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
		Fstate: uint16(libc.Int32FromInt32(125)),
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
		Fstate: uint16(libc.Int32FromInt32(132)),
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
		Fstate: uint16(libc.Int32FromInt32(121)),
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
		Fstate: uint16(libc.Int32FromInt32(66)),
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
		Fstate: uint16(libc.Int32FromInt32(117)),
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
		Fstate: uint16(libc.Int32FromInt32(15)),
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
		Fstate: uint16(libc.Int32FromInt32(123)),
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
		Fstate: uint16(libc.Int32FromInt32(92)),
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
		Fstate: uint16(libc.Int32FromInt32(90)),
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
		Fstate: uint16(libc.Int32FromInt32(124)),
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
		Fstate: uint16(libc.Int32FromInt32(30)),
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
		Fstate: uint16(libc.Int32FromInt32(119)),
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
		Fstate: uint16(libc.Int32FromInt32(10)),
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
		Fstate: uint16(libc.Int32FromInt32(117)),
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
		Fstate: uint16(libc.Int32FromInt32(33)),
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
		Fstate: uint16(libc.Int32FromInt32(65)),
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
		Fstate: uint16(libc.Int32FromInt32(15)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_params_repeat1),
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
		Fstate: uint16(libc.Int32FromInt32(81)),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_configuration),
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
		Fsymbol:      uint16(aux_sym_configuration_repeat1),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_configuration_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(125)),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_configuration_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(132)),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_configuration_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(121)),
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
		Fsymbol:      uint16(aux_sym_configuration_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(66)),
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
		Fsymbol:      uint16(aux_sym_configuration_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(117)),
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
		Fsymbol:      uint16(aux_sym_configuration_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(15)),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_variable),
		Fproduction_id: uint16(1),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_variable),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__value),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fsymbol:      uint16(sym_mod),
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
		Fcount: uint8(1),
	}})),
	78: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_mod),
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
		Fstate: uint16(libc.Int32FromInt32(119)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_gradient),
	})))),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(115)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_window_rule),
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
		Fstate: uint16(libc.Int32FromInt32(127)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(47)),
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
		Fstate: uint16(libc.Int32FromInt32(46)),
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
		Fstate: uint16(libc.Int32FromInt32(45)),
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
		Fsymbol:      uint16(aux_sym_gradient_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(119)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_gradient_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(119)),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_gradient_repeat1),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_gradient_repeat1),
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
		Fsymbol:      uint16(aux_sym_gradient_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(33)),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__linebreak),
	})))),
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
		Fcount: uint8(1),
	}})),
	111: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__linebreak),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_source),
	})))),
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
		Fcount: uint8(1),
	}})),
	115: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_source),
	})))),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_exec),
	})))),
	118: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	119: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_exec),
	})))),
	120: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	121: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_configuration_repeat1),
	})))),
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
		Fcount: uint8(1),
	}})),
	123: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_configuration_repeat1),
	})))),
	124: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	125: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_assignment),
		Fproduction_id: uint16(3),
	})))),
	126: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	127: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_assignment),
		Fproduction_id: uint16(3),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	129: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_keyword),
		Fproduction_id: uint16(4),
	})))),
	130: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	131: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_keyword),
		Fproduction_id: uint16(4),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	133: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_declaration),
		Fproduction_id: uint16(3),
	})))),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_declaration),
		Fproduction_id: uint16(3),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_arguments),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	139: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_exec),
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
		Fcount: uint8(1),
	}})),
	141: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_exec),
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
		Fcount:    uint8(1),
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
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_assignment),
		Fproduction_id: uint16(2),
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
		Fcount: uint8(1),
	}})),
	145: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_assignment),
		Fproduction_id: uint16(2),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	147: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_section),
		Fproduction_id: uint16(2),
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
		Fcount: uint8(1),
	}})),
	149: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_section),
		Fproduction_id: uint16(2),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_arguments_repeat1),
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
		Fcount: uint8(2),
	}})),
	153: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      uint16(libc.Int32FromInt32(127)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_arguments_repeat1),
	})))),
	157: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(2),
	}})),
	159: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      uint16(libc.Int32FromInt32(46)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	162: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	163: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(libc.Int32FromInt32(45)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	165: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_section),
		Fproduction_id: uint16(2),
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
		Fcount: uint8(1),
	}})),
	167: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_section),
		Fproduction_id: uint16(2),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	169: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_section),
		Fproduction_id: uint16(5),
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
		Fcount: uint8(1),
	}})),
	171: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_section),
		Fproduction_id: uint16(5),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_section),
		Fproduction_id: uint16(5),
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
		Fcount: uint8(1),
	}})),
	175: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_section),
		Fproduction_id: uint16(5),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	177: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fcount: uint8(1),
	}})),
	179: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_number),
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
		Fsymbol:      uint16(sym_number),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(124)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__zero),
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
		Fcount: uint8(1),
	}})),
	191: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__zero),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(74)),
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
		Fstate: uint16(libc.Int32FromInt32(65)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(93)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	203: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_section_repeat1),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_section_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(74)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_section_repeat1),
	})))),
	211: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(libc.Int32FromInt32(65)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	215: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(105)),
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
		Fstate: uint16(libc.Int32FromInt32(109)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(106)),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_number),
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
		Fstate: uint16(libc.Int32FromInt32(129)),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_arguments_repeat1),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_arguments_repeat1),
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
		Fsymbol:      uint16(sym_legacy_hex),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_legacy_hex),
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
		Fsymbol:      uint16(sym_rgb),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_rgb),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_gradient_repeat1),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_gradient_repeat1),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_color),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_color),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_rules_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(110)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	251: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_rules_repeat1),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	253: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_number_tuple),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(52)),
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
		Fstate: uint16(libc.Int32FromInt32(87)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(85)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(110)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	263: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	264: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	265: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_number_tuple_repeat1),
	})))),
	266: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	267: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_number_tuple_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(52)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(sym_params),
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
		Fstate: uint16(libc.Int32FromInt32(3)),
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
		Fstate: uint16(libc.Int32FromInt32(118)),
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
		Fstate: uint16(libc.Int32FromInt32(104)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_number_tuple),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_params_repeat1),
	})))),
	283: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	284: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	285: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_params_repeat1),
	})))),
	286: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	287: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_position),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	289: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	290: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(135)),
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
		Fstate: uint16(libc.Int32FromInt32(80)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(libc.Int32FromInt32(53)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	297: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_section_repeat1),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	299: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_angle),
	})))),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_gradient),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	305: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_display),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_vec2),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(sym_keys),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	315: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_window_rule),
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
	317: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f uint8_t
		_ [7]byte
	}{f: uint8(TSParseActionTypeAccept)})),
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
		Fcount: uint8(1),
	}})),
	319: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_rules),
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
		Fstate: uint16(libc.Int32FromInt32(5)),
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
		Fstate: uint16(libc.Int32FromInt32(8)),
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
		Fstate: uint16(libc.Int32FromInt32(130)),
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
		Fstate: uint16(libc.Int32FromInt32(34)),
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
		Fstate: uint16(libc.Int32FromInt32(126)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(82)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(122)),
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
		Fstate: uint16(libc.Int32FromInt32(48)),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_rules),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	345: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	347: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(51)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(112)),
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
		Fstate: uint16(libc.Int32FromInt32(134)),
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
		Fsymbol:      uint16(sym_comment),
	})))),
}

func tree_sitter_hyprlang(tls *libc.TLS) (r uintptr) {
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
	Fkeyword_capture_token:     uint16(sym_string),
	Fprimary_state_ids:         uintptr(unsafe.Pointer(&ts_primary_state_ids)),
	Fname:                      __ccgo_ts + 661,
	Fmetadata: TSLanguageMetadata{
		Fmajor_version: uint8(3),
		Fminor_version: uint8(1),
	},
}

func init() {
	p := unsafe.Pointer(&language)
	*(*uintptr)(unsafe.Add(p, 144)) = __ccgo_fp(ts_lex)
	*(*uintptr)(unsafe.Add(p, 152)) = __ccgo_fp(ts_lex_keywords)
}

func __ccgo_fp(f interface{}) uintptr {
	type iface [2]uintptr
	return (*iface)(unsafe.Pointer(&f))[1]
}

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = "end\x00string\x00=\x00:\x00{\x00}\x00source\x00[\x00;\x00]\x00exec-once\x00exec\x00execr-once\x00execr\x00exec-shutdown\x00true\x00false\x00on\x00off\x00yes\x00no\x00+\x00-\x00number_token1\x00x\x00rgb\x00rgba\x00(\x00)\x00,\x00@\x00hex\x00angle_token1\x00deg\x00SHIFT\x00CAPS\x00CTRL\x00CONTROL\x00ALT\x00ALT_L\x00MOD2\x00MOD3\x00SUPER\x00WIN\x00LOGO\x00MOD4\x00MOD5\x00TAB\x00string_literal\x00name\x00device_name\x00$\x00variable_token1\x000\x00\n\x00#\x00comment_token1\x00configuration\x00declaration\x00assignment\x00keyword\x00section\x00arguments\x00window_rule\x00rules\x00_value\x00boolean\x00number\x00vec2\x00color\x00legacy_hex\x00gradient\x00number_tuple\x00display\x00position\x00angle\x00mod\x00keys\x00params\x00variable\x00_zero\x00_linebreak\x00comment\x00configuration_repeat1\x00section_repeat1\x00arguments_repeat1\x00rules_repeat1\x00gradient_repeat1\x00number_tuple_repeat1\x00params_repeat1\x00device\x00value\x00hyprlang\x00"
