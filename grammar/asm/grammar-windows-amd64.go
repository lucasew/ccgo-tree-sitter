// Code generated for windows/amd64 by 'ccgo -ignore-unsupported-alignment -ignore-unsupported-atomic-sizes -ignore-vector-functions --winapi-no-errno -ignore-link-errors -std=gnu11 -O0 -D__extension__= -D__forceinline=static inline -D__attribute__(...)= -D__declspec(x)= -D__cdecl= -D__stdcall= -D__fastcall= -D__thiscall= -D_cdecl= -D__restrict= -D__restrict__= -D__MINGW_EXTENSION= -D_X86INTRIN_H_INCLUDED -D_X86GPRINTRIN_H_INCLUDED -D_IMMINTRIN_H_INCLUDED -D_MMINTRIN_H_INCLUDED -D_XMMINTRIN_H_INCLUDED -D_EMMINTRIN_H_INCLUDED -D_PMMINTRIN_H_INCLUDED -D_MM3DNOW_H_INCLUDED -Dfunc=func_token -Dinterface=interface_token -Dselect=select_token -Dchan=chan_token -Dgo=go_token -Dmap=map_token -Dpackage=package_token -Dtype=type_token -Dvar=var_token -Dimport=import_token -Ddefer=defer_token -Dfallthrough=fallthrough_token -Drange=range_token -I D:\a\ccgo-tree-sitter\ccgo-tree-sitter\third-party\tree-sitter-asm\src -I D:\a\ccgo-tree-sitter\ccgo-tree-sitter\third-party\tree-sitter-asm -I D:\a\ccgo-tree-sitter\ccgo-tree-sitter\third-party -I C:\Users\runneradmin\.cache\workspaced\sources\github\7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2\lib\include -I C:\Users\runneradmin\.cache\workspaced\sources\github\7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2\lib\src D:\a\ccgo-tree-sitter\ccgo-tree-sitter\third-party\tree-sitter-asm\src\parser.c -o grammar.go', DO NOT EDIT.

//go:build windows && amd64

package grammar_asm

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
const FIELD_COUNT = 6
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
const LARGE_STATE_COUNT = 3
const MAX_ALIAS_SEQUENCE_LENGTH = 7
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
const SYMBOL_COUNT = 65
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

const anon_sym_LF = 1
const anon_sym_COMMA = 2
const anon_sym_COLON = 3
const anon_sym_LPAREN = 4
const anon_sym_RPAREN = 5
const anon_sym_label = 6
const anon_sym_const = 7
const anon_sym_LBRACE = 8
const anon_sym_DASH = 9
const anon_sym_RBRACE = 10
const anon_sym_byte = 11
const anon_sym_word = 12
const anon_sym_dword = 13
const anon_sym_qword = 14
const anon_sym_ptr = 15
const anon_sym_LBRACK = 16
const anon_sym_PLUS = 17
const anon_sym_RBRACK = 18
const anon_sym_STAR = 19
const anon_sym_rel = 20
const anon_sym_BANG = 21
const anon_sym_SLASH = 22
const anon_sym_PERCENT = 23
const anon_sym_PIPE = 24
const anon_sym_CARET = 25
const anon_sym_AMP = 26
const anon_sym_POUND = 27
const aux_sym_int_token1 = 28
const aux_sym_int_token2 = 29
const sym_float = 30
const aux_sym_string_token1 = 31
const aux_sym_string_token2 = 32
const sym_word = 33
const sym__reg = 34
const sym_address = 35
const sym_meta_ident = 36
const sym__ident = 37
const aux_sym_line_comment_token1 = 38
const aux_sym_line_comment_token2 = 39
const sym_block_comment = 40
const sym_program = 41
const sym__item = 42
const sym_meta = 43
const sym_label = 44
const sym_const = 45
const sym_instruction = 46
const sym__expr = 47
const sym_list = 48
const sym_ptr = 49
const sym__tc_expr = 50
const sym_tc_infix = 51
const sym_int = 52
const sym_string = 53
const sym_reg = 54
const sym_ident = 55
const sym_line_comment = 56
const aux_sym_program_repeat1 = 57
const aux_sym_program_repeat2 = 58
const aux_sym_meta_repeat1 = 59
const aux_sym_meta_repeat2 = 60
const aux_sym_meta_repeat3 = 61
const aux_sym_instruction_repeat1 = 62
const aux_sym_instruction_repeat2 = 63
const aux_sym_list_repeat1 = 64

var ts_symbol_names = [65]uintptr{
	0:  __ccgo_ts,
	1:  __ccgo_ts + 4,
	2:  __ccgo_ts + 6,
	3:  __ccgo_ts + 8,
	4:  __ccgo_ts + 10,
	5:  __ccgo_ts + 12,
	6:  __ccgo_ts + 14,
	7:  __ccgo_ts + 20,
	8:  __ccgo_ts + 26,
	9:  __ccgo_ts + 28,
	10: __ccgo_ts + 30,
	11: __ccgo_ts + 32,
	12: __ccgo_ts + 37,
	13: __ccgo_ts + 42,
	14: __ccgo_ts + 48,
	15: __ccgo_ts + 54,
	16: __ccgo_ts + 58,
	17: __ccgo_ts + 60,
	18: __ccgo_ts + 62,
	19: __ccgo_ts + 64,
	20: __ccgo_ts + 66,
	21: __ccgo_ts + 70,
	22: __ccgo_ts + 72,
	23: __ccgo_ts + 74,
	24: __ccgo_ts + 76,
	25: __ccgo_ts + 78,
	26: __ccgo_ts + 80,
	27: __ccgo_ts + 82,
	28: __ccgo_ts + 84,
	29: __ccgo_ts + 95,
	30: __ccgo_ts + 106,
	31: __ccgo_ts + 112,
	32: __ccgo_ts + 126,
	33: __ccgo_ts + 37,
	34: __ccgo_ts + 140,
	35: __ccgo_ts + 145,
	36: __ccgo_ts + 153,
	37: __ccgo_ts + 164,
	38: __ccgo_ts + 171,
	39: __ccgo_ts + 191,
	40: __ccgo_ts + 211,
	41: __ccgo_ts + 225,
	42: __ccgo_ts + 233,
	43: __ccgo_ts + 239,
	44: __ccgo_ts + 14,
	45: __ccgo_ts + 20,
	46: __ccgo_ts + 244,
	47: __ccgo_ts + 256,
	48: __ccgo_ts + 262,
	49: __ccgo_ts + 54,
	50: __ccgo_ts + 267,
	51: __ccgo_ts + 276,
	52: __ccgo_ts + 285,
	53: __ccgo_ts + 289,
	54: __ccgo_ts + 296,
	55: __ccgo_ts + 300,
	56: __ccgo_ts + 306,
	57: __ccgo_ts + 319,
	58: __ccgo_ts + 335,
	59: __ccgo_ts + 351,
	60: __ccgo_ts + 364,
	61: __ccgo_ts + 377,
	62: __ccgo_ts + 390,
	63: __ccgo_ts + 410,
	64: __ccgo_ts + 430,
}

var ts_symbol_map = [65]TSSymbol{
	1:  uint16(anon_sym_LF),
	2:  uint16(anon_sym_COMMA),
	3:  uint16(anon_sym_COLON),
	4:  uint16(anon_sym_LPAREN),
	5:  uint16(anon_sym_RPAREN),
	6:  uint16(anon_sym_label),
	7:  uint16(anon_sym_const),
	8:  uint16(anon_sym_LBRACE),
	9:  uint16(anon_sym_DASH),
	10: uint16(anon_sym_RBRACE),
	11: uint16(anon_sym_byte),
	12: uint16(anon_sym_word),
	13: uint16(anon_sym_dword),
	14: uint16(anon_sym_qword),
	15: uint16(anon_sym_ptr),
	16: uint16(anon_sym_LBRACK),
	17: uint16(anon_sym_PLUS),
	18: uint16(anon_sym_RBRACK),
	19: uint16(anon_sym_STAR),
	20: uint16(anon_sym_rel),
	21: uint16(anon_sym_BANG),
	22: uint16(anon_sym_SLASH),
	23: uint16(anon_sym_PERCENT),
	24: uint16(anon_sym_PIPE),
	25: uint16(anon_sym_CARET),
	26: uint16(anon_sym_AMP),
	27: uint16(anon_sym_POUND),
	28: uint16(aux_sym_int_token1),
	29: uint16(aux_sym_int_token2),
	30: uint16(sym_float),
	31: uint16(aux_sym_string_token1),
	32: uint16(aux_sym_string_token2),
	33: uint16(sym_word),
	34: uint16(sym__reg),
	35: uint16(sym_address),
	36: uint16(sym_meta_ident),
	37: uint16(sym__ident),
	38: uint16(aux_sym_line_comment_token1),
	39: uint16(aux_sym_line_comment_token2),
	40: uint16(sym_block_comment),
	41: uint16(sym_program),
	42: uint16(sym__item),
	43: uint16(sym_meta),
	44: uint16(sym_label),
	45: uint16(sym_const),
	46: uint16(sym_instruction),
	47: uint16(sym__expr),
	48: uint16(sym_list),
	49: uint16(sym_ptr),
	50: uint16(sym__tc_expr),
	51: uint16(sym_tc_infix),
	52: uint16(sym_int),
	53: uint16(sym_string),
	54: uint16(sym_reg),
	55: uint16(sym_ident),
	56: uint16(sym_line_comment),
	57: uint16(aux_sym_program_repeat1),
	58: uint16(aux_sym_program_repeat2),
	59: uint16(aux_sym_meta_repeat1),
	60: uint16(aux_sym_meta_repeat2),
	61: uint16(aux_sym_meta_repeat3),
	62: uint16(aux_sym_instruction_repeat1),
	63: uint16(aux_sym_instruction_repeat2),
	64: uint16(aux_sym_list_repeat1),
}

var ts_symbol_metadata = [65]TSSymbolMetadata{
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
	28: {},
	29: {},
	30: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	31: {},
	32: {},
	33: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	34: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	35: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	36: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	37: {
		Fnamed: libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
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
	56: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	57: {},
	58: {},
	59: {},
	60: {},
	61: {},
	62: {},
	63: {},
	64: {},
}

type ts_field_identifiers = int32

const field_kind = 1
const field_lhs = 2
const field_name = 3
const field_op = 4
const field_rhs = 5
const field_value = 6

var ts_field_names = [7]uintptr{
	0: libc.UintptrFromInt32(0),
	1: __ccgo_ts + 443,
	2: __ccgo_ts + 448,
	3: __ccgo_ts + 452,
	4: __ccgo_ts + 457,
	5: __ccgo_ts + 460,
	6: __ccgo_ts + 464,
}

var ts_field_map_slices = [6]TSFieldMapSlice{
	1: {
		Flength: uint16(1),
	},
	2: {
		Findex:  uint16(1),
		Flength: uint16(1),
	},
	4: {
		Findex:  uint16(2),
		Flength: uint16(2),
	},
	5: {
		Findex:  uint16(4),
		Flength: uint16(3),
	},
}

var ts_field_map_entries = [7]TSFieldMapEntry{
	0: {
		Ffield_id: uint16(field_kind),
	},
	1: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
	},
	2: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
	},
	3: {
		Ffield_id:    uint16(field_value),
		Fchild_index: uint8(2),
	},
	4: {
		Ffield_id: uint16(field_lhs),
	},
	5: {
		Ffield_id:    uint16(field_op),
		Fchild_index: uint8(1),
	},
	6: {
		Ffield_id:    uint16(field_rhs),
		Fchild_index: uint8(2),
	},
}

var ts_alias_sequences = [6][7]TSSymbol{
	0: {},
	3: {
		0: uint16(sym_ident),
	},
}

var ts_non_terminal_alias_map = [1]uint16_t{}

var ts_primary_state_ids = [137]TSStateId{
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
	29:  uint16(26),
	30:  uint16(25),
	31:  uint16(24),
	32:  uint16(32),
	33:  uint16(27),
	34:  uint16(28),
	35:  uint16(9),
	36:  uint16(8),
	37:  uint16(37),
	38:  uint16(38),
	39:  uint16(6),
	40:  uint16(11),
	41:  uint16(41),
	42:  uint16(12),
	43:  uint16(17),
	44:  uint16(19),
	45:  uint16(16),
	46:  uint16(46),
	47:  uint16(13),
	48:  uint16(14),
	49:  uint16(15),
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
	134: uint16(125),
	135: uint16(135),
	136: uint16(136),
}

func ts_lex(tls *libc.TLS, lexer uintptr, state TSStateId) (r uint8) {
	var eof, result, skip uint8
	var i, i1, i2, i3, i4, i5, i6 uint32_t
	var lookahead int32_t
	_, _, _, _, _, _, _, _, _, _, _ = eof, i, i1, i2, i3, i4, i5, i6, lookahead, result, skip
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
			if !(uint64(i) < libc.Uint64FromInt64(136)/libc.Uint64FromInt64(2)) {
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
		if lookahead == int32('\t') || lookahead == int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(0)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(77)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(1):
		if lookahead == int32('"') {
			state = uint16(93)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(1)
			goto next_state
		}
		return result
	case int32(2):
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
		if lookahead == int32('\t') || lookahead == int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(2)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(86)
			goto next_state
		}
		return result
	case int32(3):
		if lookahead == int32('#') {
			state = uint16(65)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(21)
			goto next_state
		}
		if lookahead == int32('-') {
			state = uint16(8)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(10)
			goto next_state
		}
		if lookahead == int32('0') {
			state = uint16(67)
			goto next_state
		}
		if lookahead == int32(';') {
			state = uint16(154)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(5)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead == int32('#') {
			state = uint16(65)
			goto next_state
		}
		if lookahead == int32('%') {
			state = uint16(23)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(147)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(10)
			goto next_state
		}
		if lookahead == int32(';') {
			state = uint16(154)
			goto next_state
		}
		if lookahead == int32('$') || lookahead == int32('=') {
			state = uint16(24)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(4)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') {
			state = uint16(140)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(139)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead == int32('#') {
			state = uint16(65)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(10)
			goto next_state
		}
		if lookahead == int32(';') {
			state = uint16(154)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(5)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead == int32('$') {
			state = uint16(20)
			goto next_state
		}
		if lookahead == int32('0') {
			state = uint16(84)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(86)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead == int32('$') {
			state = uint16(20)
			goto next_state
		}
		if lookahead == int32('0') {
			state = uint16(82)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(83)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead == int32('$') {
			state = uint16(21)
			goto next_state
		}
		if lookahead == int32('0') {
			state = uint16(67)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead == int32('\'') {
			state = uint16(94)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(9)
			goto next_state
		}
		return result
	case int32(10):
		if lookahead == int32('*') {
			state = uint16(12)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(154)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead == int32('*') {
			state = uint16(11)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(155)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(12)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead == int32('*') {
			state = uint16(11)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(12)
			goto next_state
		}
		return result
	case int32(13):
		if lookahead == int32('e') {
			state = uint16(14)
			goto next_state
		}
		return result
	case int32(14):
		if lookahead == int32('l') {
			state = uint16(56)
			goto next_state
		}
		return result
	case int32(15):
		if lookahead == int32('r') {
			state = uint16(50)
			goto next_state
		}
		return result
	case int32(16):
		if lookahead == int32('t') {
			state = uint16(15)
			goto next_state
		}
		return result
	case int32(17):
		if lookahead == int32('0') || lookahead == int32('1') {
			state = uint16(85)
			goto next_state
		}
		return result
	case int32(18):
		if lookahead == int32('0') || lookahead == int32('1') {
			state = uint16(68)
			goto next_state
		}
		return result
	case int32(19):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(87)
			goto next_state
		}
		if int32('G') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('g') <= lookahead && lookahead <= int32('z') {
			state = uint16(144)
			goto next_state
		}
		return result
	case int32(20):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(88)
			goto next_state
		}
		return result
	case int32(21):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(70)
			goto next_state
		}
		return result
	case int32(22):
		if lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(146)
			goto next_state
		}
		return result
	case int32(23):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(143)
			goto next_state
		}
		return result
	case int32(24):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(144)
			goto next_state
		}
		return result
	case int32(25):
		if eof != 0 {
			state = uint16(30)
			goto next_state
		}
		i2 = uint32(0)
		for {
			if !(uint64(i2) < libc.Uint64FromInt64(84)/libc.Uint64FromInt64(2)) {
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
		if lookahead == int32('\t') || lookahead == int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(25)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(80)
			goto next_state
		}
		if int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(139)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(26):
		if eof != 0 {
			state = uint16(30)
			goto next_state
		}
		i3 = uint32(0)
		for {
			if !(uint64(i3) < libc.Uint64FromInt64(52)/libc.Uint64FromInt64(2)) {
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
		if lookahead == int32('\t') || lookahead == int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(26)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(80)
			goto next_state
		}
		if int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(139)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(27):
		if eof != 0 {
			state = uint16(30)
			goto next_state
		}
		i4 = uint32(0)
		for {
			if !(uint64(i4) < libc.Uint64FromInt64(76)/libc.Uint64FromInt64(2)) {
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
		if lookahead == int32('\t') || lookahead == int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(27)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(72)
			goto next_state
		}
		if int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(139)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(28):
		if eof != 0 {
			state = uint16(30)
			goto next_state
		}
		i5 = uint32(0)
		for {
			if !(uint64(i5) < libc.Uint64FromInt64(76)/libc.Uint64FromInt64(2)) {
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
		if lookahead == int32('\t') || lookahead == int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(28)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(29):
		if eof != 0 {
			state = uint16(30)
			goto next_state
		}
		i6 = uint32(0)
		for {
			if !(uint64(i6) < libc.Uint64FromInt64(40)/libc.Uint64FromInt64(2)) {
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
		if lookahead == int32('\t') || lookahead == int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(29)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') {
			state = uint16(142)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(141)
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
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LF)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(32):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COMMA)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(33):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COLON)
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
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_label)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(37):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_const)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(38):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(39):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(40):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('$') {
			state = uint16(20)
			goto next_state
		}
		if lookahead == int32('0') {
			state = uint16(84)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(86)
			goto next_state
		}
		return result
	case int32(41):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(42):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_byte)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') {
			state = uint16(140)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(139)
			goto next_state
		}
		return result
	case int32(43):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_byte)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(44):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') {
			state = uint16(140)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(139)
			goto next_state
		}
		return result
	case int32(45):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(46):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_dword)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') {
			state = uint16(140)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(139)
			goto next_state
		}
		return result
	case int32(47):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_dword)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(48):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_qword)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') {
			state = uint16(140)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(139)
			goto next_state
		}
		return result
	case int32(49):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_qword)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(50):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ptr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(51):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ptr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(52):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(53):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PLUS)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(54):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(55):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STAR)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(56):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_rel)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(57):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_rel)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(58):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BANG)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(59):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SLASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('*') {
			state = uint16(12)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(154)
			goto next_state
		}
		return result
	case int32(60):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PERCENT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(61):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PERCENT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(143)
			goto next_state
		}
		return result
	case int32(62):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PIPE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(63):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_CARET)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(64):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AMP)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(65):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUND)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(66):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUND)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(153)
			goto next_state
		}
		return result
	case int32(67):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_int_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('b') {
			state = uint16(18)
			goto next_state
		}
		if lookahead == int32('x') {
			state = uint16(21)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || lookahead == int32('_') {
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(68):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_int_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('0') || lookahead == int32('1') || lookahead == int32('_') {
			state = uint16(68)
			goto next_state
		}
		return result
	case int32(69):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_int_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || lookahead == int32('_') {
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(70):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_int_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(70)
			goto next_state
		}
		return result
	case int32(71):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_int_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if lookahead == int32('_') {
			state = uint16(77)
			goto next_state
		}
		if lookahead == int32('b') {
			state = uint16(136)
			goto next_state
		}
		if lookahead == int32('x') {
			state = uint16(137)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(72)
			goto next_state
		}
		if int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(139)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(72):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_int_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if lookahead == int32('_') {
			state = uint16(77)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(72)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') {
			state = uint16(140)
			goto next_state
		}
		if int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(139)
			goto next_state
		}
		return result
	case int32(73):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_int_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if lookahead == int32('_') {
			state = uint16(75)
			goto next_state
		}
		if lookahead == int32('0') || lookahead == int32('1') {
			state = uint16(73)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') {
			state = uint16(140)
			goto next_state
		}
		if int32('2') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(139)
			goto next_state
		}
		return result
	case int32(74):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_int_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if lookahead == int32('b') {
			state = uint16(135)
			goto next_state
		}
		if lookahead == int32('x') {
			state = uint16(138)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || lookahead == int32('_') {
			state = uint16(77)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(75):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_int_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if lookahead == int32('0') || lookahead == int32('1') || lookahead == int32('_') {
			state = uint16(75)
			goto next_state
		}
		if int32('2') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(76):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_int_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('F') || lookahead == int32('_') {
			state = uint16(78)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(76)
			goto next_state
		}
		if int32('G') <= lookahead && lookahead <= int32('Z') {
			state = uint16(140)
			goto next_state
		}
		if int32('g') <= lookahead && lookahead <= int32('z') {
			state = uint16(139)
			goto next_state
		}
		return result
	case int32(77):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_int_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || lookahead == int32('_') {
			state = uint16(77)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(78):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_int_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(78)
			goto next_state
		}
		if int32('G') <= lookahead && lookahead <= int32('Z') || int32('g') <= lookahead && lookahead <= int32('z') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(79):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_int_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(89)
			goto next_state
		}
		if lookahead == int32('_') {
			state = uint16(81)
			goto next_state
		}
		if lookahead == int32('b') {
			state = uint16(136)
			goto next_state
		}
		if lookahead == int32('x') {
			state = uint16(137)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(80)
			goto next_state
		}
		if int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(139)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(80):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_int_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(89)
			goto next_state
		}
		if lookahead == int32('_') {
			state = uint16(81)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(80)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') {
			state = uint16(140)
			goto next_state
		}
		if int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(139)
			goto next_state
		}
		return result
	case int32(81):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_int_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(89)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || lookahead == int32('_') {
			state = uint16(81)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(82):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_int_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(90)
			goto next_state
		}
		if lookahead == int32('b') {
			state = uint16(17)
			goto next_state
		}
		if lookahead == int32('x') {
			state = uint16(20)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || lookahead == int32('_') {
			state = uint16(83)
			goto next_state
		}
		return result
	case int32(83):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_int_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(90)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || lookahead == int32('_') {
			state = uint16(83)
			goto next_state
		}
		return result
	case int32(84):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_int_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('b') {
			state = uint16(17)
			goto next_state
		}
		if lookahead == int32('x') {
			state = uint16(20)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || lookahead == int32('_') {
			state = uint16(86)
			goto next_state
		}
		return result
	case int32(85):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_int_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('0') || lookahead == int32('1') || lookahead == int32('_') {
			state = uint16(85)
			goto next_state
		}
		return result
	case int32(86):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_int_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || lookahead == int32('_') {
			state = uint16(86)
			goto next_state
		}
		return result
	case int32(87):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_int_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(87)
			goto next_state
		}
		if int32('G') <= lookahead && lookahead <= int32('Z') || int32('g') <= lookahead && lookahead <= int32('z') {
			state = uint16(144)
			goto next_state
		}
		return result
	case int32(88):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_int_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(88)
			goto next_state
		}
		return result
	case int32(89):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_float)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(91)
			goto next_state
		}
		if lookahead == int32('.') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(148)
			goto next_state
		}
		return result
	case int32(90):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_float)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(92)
			goto next_state
		}
		return result
	case int32(91):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_float)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || lookahead == int32('_') {
			state = uint16(91)
			goto next_state
		}
		if lookahead == int32('.') || int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(148)
			goto next_state
		}
		return result
	case int32(92):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_float)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || lookahead == int32('_') {
			state = uint16(92)
			goto next_state
		}
		return result
	case int32(93):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(94):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(95):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if lookahead == int32('a') {
			state = uint16(96)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('b') <= lookahead && lookahead <= int32('z') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(96):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if lookahead == int32('b') {
			state = uint16(106)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(97):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if lookahead == int32('d') {
			state = uint16(45)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(98):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if lookahead == int32('d') {
			state = uint16(47)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(99):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if lookahead == int32('d') {
			state = uint16(49)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(100):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if lookahead == int32('d') {
			state = uint16(44)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') {
			state = uint16(140)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(139)
			goto next_state
		}
		return result
	case int32(101):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if lookahead == int32('d') {
			state = uint16(46)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') {
			state = uint16(140)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(139)
			goto next_state
		}
		return result
	case int32(102):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if lookahead == int32('d') {
			state = uint16(48)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') {
			state = uint16(140)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(139)
			goto next_state
		}
		return result
	case int32(103):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(107)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(104):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(43)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(105):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(42)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') {
			state = uint16(140)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(139)
			goto next_state
		}
		return result
	case int32(106):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(108)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(107):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if lookahead == int32('l') {
			state = uint16(57)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(108):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if lookahead == int32('l') {
			state = uint16(36)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(109):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(124)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(110):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(109)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(111):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(119)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') {
			state = uint16(140)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(139)
			goto next_state
		}
		return result
	case int32(112):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(118)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(113):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(120)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(114):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(122)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(115):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(121)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') {
			state = uint16(140)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(139)
			goto next_state
		}
		return result
	case int32(116):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(123)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') {
			state = uint16(140)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(139)
			goto next_state
		}
		return result
	case int32(117):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(51)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(118):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(97)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(119):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(100)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') {
			state = uint16(140)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(139)
			goto next_state
		}
		return result
	case int32(120):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(98)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(121):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(101)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') {
			state = uint16(140)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(139)
			goto next_state
		}
		return result
	case int32(122):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(99)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(123):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(102)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') {
			state = uint16(140)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(139)
			goto next_state
		}
		return result
	case int32(124):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(126)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(125):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(117)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(126):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(37)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(127):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(105)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') {
			state = uint16(140)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(139)
			goto next_state
		}
		return result
	case int32(128):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(104)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(129):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if lookahead == int32('w') {
			state = uint16(113)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(130):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if lookahead == int32('w') {
			state = uint16(114)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(131):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if lookahead == int32('w') {
			state = uint16(115)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') {
			state = uint16(140)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(139)
			goto next_state
		}
		return result
	case int32(132):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if lookahead == int32('w') {
			state = uint16(116)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') {
			state = uint16(140)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(139)
			goto next_state
		}
		return result
	case int32(133):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if lookahead == int32('y') {
			state = uint16(127)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') {
			state = uint16(140)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(139)
			goto next_state
		}
		return result
	case int32(134):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if lookahead == int32('y') {
			state = uint16(128)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(135):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if lookahead == int32('0') || lookahead == int32('1') {
			state = uint16(75)
			goto next_state
		}
		if int32('2') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(136):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if lookahead == int32('0') || lookahead == int32('1') {
			state = uint16(73)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') {
			state = uint16(140)
			goto next_state
		}
		if int32('2') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(139)
			goto next_state
		}
		return result
	case int32(137):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('F') {
			state = uint16(78)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(76)
			goto next_state
		}
		if int32('g') <= lookahead && lookahead <= int32('z') {
			state = uint16(139)
			goto next_state
		}
		if int32('G') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(138):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(78)
			goto next_state
		}
		if int32('G') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('g') <= lookahead && lookahead <= int32('z') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(139):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') {
			state = uint16(140)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(139)
			goto next_state
		}
		return result
	case int32(140):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(148)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(141):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') {
			state = uint16(142)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(142):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(142)
			goto next_state
		}
		return result
	case int32(143):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__reg)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(143)
			goto next_state
		}
		return result
	case int32(144):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_address)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(144)
			goto next_state
		}
		return result
	case int32(145):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_meta_ident)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(145)
			goto next_state
		}
		if lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') {
			state = uint16(148)
			goto next_state
		}
		return result
	case int32(146):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_meta_ident)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(146)
			goto next_state
		}
		return result
	case int32(147):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__ident)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(145)
			goto next_state
		}
		if lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') {
			state = uint16(148)
			goto next_state
		}
		return result
	case int32(148):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__ident)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(148)
			goto next_state
		}
		return result
	case int32(149):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_line_comment_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(12)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(149)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(153)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(150)
			goto next_state
		}
		return result
	case int32(150):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_line_comment_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(12)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(149)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(150)
			goto next_state
		}
		return result
	case int32(151):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_line_comment_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(66)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(152)
			goto next_state
		}
		if lookahead == int32(';') {
			state = uint16(153)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32('\r') || lookahead == int32(' ') {
			state = uint16(151)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\t') && lookahead != int32('\n') {
			state = uint16(153)
			goto next_state
		}
		return result
	case int32(152):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_line_comment_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('*') {
			state = uint16(150)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(153)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(153)
			goto next_state
		}
		return result
	case int32(153):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_line_comment_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(153)
			goto next_state
		}
		return result
	case int32(154):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_line_comment_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(154)
			goto next_state
		}
		return result
	case int32(155):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_block_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var map_token = [68]uint16_t{
	0:  uint16('\n'),
	1:  uint16(31),
	2:  uint16('!'),
	3:  uint16(58),
	4:  uint16('"'),
	5:  uint16(1),
	6:  uint16('#'),
	7:  uint16(65),
	8:  uint16('$'),
	9:  uint16(19),
	10: uint16('%'),
	11: uint16(60),
	12: uint16('&'),
	13: uint16(64),
	14: uint16('\''),
	15: uint16(9),
	16: uint16('('),
	17: uint16(34),
	18: uint16(')'),
	19: uint16(35),
	20: uint16('*'),
	21: uint16(55),
	22: uint16('+'),
	23: uint16(53),
	24: uint16(','),
	25: uint16(32),
	26: uint16('-'),
	27: uint16(40),
	28: uint16('.'),
	29: uint16(147),
	30: uint16('/'),
	31: uint16(59),
	32: uint16('0'),
	33: uint16(74),
	34: uint16(':'),
	35: uint16(33),
	36: uint16(';'),
	37: uint16(154),
	38: uint16('='),
	39: uint16(24),
	40: uint16('['),
	41: uint16(52),
	42: uint16(']'),
	43: uint16(54),
	44: uint16('^'),
	45: uint16(63),
	46: uint16('b'),
	47: uint16(134),
	48: uint16('c'),
	49: uint16(110),
	50: uint16('d'),
	51: uint16(129),
	52: uint16('l'),
	53: uint16(95),
	54: uint16('p'),
	55: uint16(125),
	56: uint16('q'),
	57: uint16(130),
	58: uint16('r'),
	59: uint16(103),
	60: uint16('w'),
	61: uint16(112),
	62: uint16('{'),
	63: uint16(38),
	64: uint16('|'),
	65: uint16(62),
	66: uint16('}'),
	67: uint16(41),
}

var map_token1 = [16]uint16_t{
	0:  uint16('#'),
	1:  uint16(65),
	2:  uint16('$'),
	3:  uint16(20),
	4:  uint16('-'),
	5:  uint16(6),
	6:  uint16('/'),
	7:  uint16(10),
	8:  uint16('0'),
	9:  uint16(84),
	10: uint16(';'),
	11: uint16(154),
	12: uint16('p'),
	13: uint16(16),
	14: uint16('r'),
	15: uint16(13),
}

var map_token2 = [42]uint16_t{
	0:  uint16('\n'),
	1:  uint16(31),
	2:  uint16('"'),
	3:  uint16(1),
	4:  uint16('#'),
	5:  uint16(65),
	6:  uint16('$'),
	7:  uint16(19),
	8:  uint16('%'),
	9:  uint16(23),
	10: uint16('\''),
	11: uint16(9),
	12: uint16('('),
	13: uint16(34),
	14: uint16('*'),
	15: uint16(55),
	16: uint16('-'),
	17: uint16(7),
	18: uint16('.'),
	19: uint16(147),
	20: uint16('/'),
	21: uint16(10),
	22: uint16('0'),
	23: uint16(79),
	24: uint16(':'),
	25: uint16(33),
	26: uint16(';'),
	27: uint16(154),
	28: uint16('='),
	29: uint16(24),
	30: uint16('['),
	31: uint16(52),
	32: uint16('b'),
	33: uint16(133),
	34: uint16('d'),
	35: uint16(131),
	36: uint16('q'),
	37: uint16(132),
	38: uint16('w'),
	39: uint16(111),
	40: uint16('{'),
	41: uint16(38),
}

var map_token3 = [26]uint16_t{
	0:  uint16('\n'),
	1:  uint16(31),
	2:  uint16('"'),
	3:  uint16(1),
	4:  uint16('#'),
	5:  uint16(65),
	6:  uint16('$'),
	7:  uint16(19),
	8:  uint16('%'),
	9:  uint16(23),
	10: uint16('\''),
	11: uint16(9),
	12: uint16('-'),
	13: uint16(7),
	14: uint16('.'),
	15: uint16(147),
	16: uint16('/'),
	17: uint16(10),
	18: uint16('0'),
	19: uint16(79),
	20: uint16(':'),
	21: uint16(33),
	22: uint16(';'),
	23: uint16(154),
	24: uint16('='),
	25: uint16(24),
}

var map_token4 = [38]uint16_t{
	0:  uint16('\n'),
	1:  uint16(31),
	2:  uint16('"'),
	3:  uint16(1),
	4:  uint16('#'),
	5:  uint16(65),
	6:  uint16('$'),
	7:  uint16(19),
	8:  uint16('%'),
	9:  uint16(61),
	10: uint16('&'),
	11: uint16(64),
	12: uint16('\''),
	13: uint16(9),
	14: uint16('('),
	15: uint16(34),
	16: uint16('*'),
	17: uint16(55),
	18: uint16('+'),
	19: uint16(53),
	20: uint16(','),
	21: uint16(32),
	22: uint16('-'),
	23: uint16(40),
	24: uint16('.'),
	25: uint16(147),
	26: uint16('/'),
	27: uint16(59),
	28: uint16('0'),
	29: uint16(71),
	30: uint16(';'),
	31: uint16(154),
	32: uint16('='),
	33: uint16(24),
	34: uint16('^'),
	35: uint16(63),
	36: uint16('|'),
	37: uint16(62),
}

var map_token5 = [38]uint16_t{
	0:  uint16('\n'),
	1:  uint16(31),
	2:  uint16('#'),
	3:  uint16(65),
	4:  uint16('%'),
	5:  uint16(60),
	6:  uint16('&'),
	7:  uint16(64),
	8:  uint16('('),
	9:  uint16(34),
	10: uint16(')'),
	11: uint16(35),
	12: uint16('*'),
	13: uint16(55),
	14: uint16('+'),
	15: uint16(53),
	16: uint16(','),
	17: uint16(32),
	18: uint16('-'),
	19: uint16(39),
	20: uint16('.'),
	21: uint16(147),
	22: uint16('/'),
	23: uint16(59),
	24: uint16(';'),
	25: uint16(154),
	26: uint16(']'),
	27: uint16(54),
	28: uint16('^'),
	29: uint16(63),
	30: uint16('c'),
	31: uint16(110),
	32: uint16('l'),
	33: uint16(95),
	34: uint16('|'),
	35: uint16(62),
	36: uint16('}'),
	37: uint16(41),
}

var map_token6 = [20]uint16_t{
	0:  uint16('\n'),
	1:  uint16(31),
	2:  uint16('#'),
	3:  uint16(65),
	4:  uint16('%'),
	5:  uint16(23),
	6:  uint16('('),
	7:  uint16(34),
	8:  uint16('.'),
	9:  uint16(22),
	10: uint16('/'),
	11: uint16(10),
	12: uint16(';'),
	13: uint16(154),
	14: uint16('}'),
	15: uint16(41),
	16: uint16('$'),
	17: uint16(24),
	18: uint16('='),
	19: uint16(24),
}

var ts_lex_modes = [137]TSLexMode{
	0: {},
	1: {
		Flex_state: uint16(28),
	},
	2: {
		Flex_state: uint16(25),
	},
	3: {
		Flex_state: uint16(25),
	},
	4: {
		Flex_state: uint16(25),
	},
	5: {
		Flex_state: uint16(25),
	},
	6: {
		Flex_state: uint16(27),
	},
	7: {
		Flex_state: uint16(27),
	},
	8: {
		Flex_state: uint16(27),
	},
	9: {
		Flex_state: uint16(27),
	},
	10: {
		Flex_state: uint16(27),
	},
	11: {
		Flex_state: uint16(27),
	},
	12: {
		Flex_state: uint16(27),
	},
	13: {
		Flex_state: uint16(27),
	},
	14: {
		Flex_state: uint16(27),
	},
	15: {
		Flex_state: uint16(27),
	},
	16: {
		Flex_state: uint16(27),
	},
	17: {
		Flex_state: uint16(27),
	},
	18: {
		Flex_state: uint16(27),
	},
	19: {
		Flex_state: uint16(27),
	},
	20: {
		Flex_state: uint16(27),
	},
	21: {
		Flex_state: uint16(27),
	},
	22: {
		Flex_state: uint16(26),
	},
	23: {
		Flex_state: uint16(26),
	},
	24: {
		Flex_state: uint16(27),
	},
	25: {
		Flex_state: uint16(27),
	},
	26: {
		Flex_state: uint16(27),
	},
	27: {
		Flex_state: uint16(27),
	},
	28: {
		Flex_state: uint16(27),
	},
	29: {
		Flex_state: uint16(27),
	},
	30: {
		Flex_state: uint16(27),
	},
	31: {
		Flex_state: uint16(27),
	},
	32: {
		Flex_state: uint16(27),
	},
	33: {
		Flex_state: uint16(27),
	},
	34: {
		Flex_state: uint16(27),
	},
	35: {
		Flex_state: uint16(28),
	},
	36: {
		Flex_state: uint16(28),
	},
	37: {
		Flex_state: uint16(28),
	},
	38: {
		Flex_state: uint16(28),
	},
	39: {
		Flex_state: uint16(28),
	},
	40: {
		Flex_state: uint16(28),
	},
	41: {
		Flex_state: uint16(28),
	},
	42: {
		Flex_state: uint16(28),
	},
	43: {
		Flex_state: uint16(28),
	},
	44: {
		Flex_state: uint16(28),
	},
	45: {
		Flex_state: uint16(28),
	},
	46: {
		Flex_state: uint16(28),
	},
	47: {
		Flex_state: uint16(28),
	},
	48: {
		Flex_state: uint16(28),
	},
	49: {
		Flex_state: uint16(28),
	},
	50: {
		Flex_state: uint16(27),
	},
	51: {
		Flex_state: uint16(27),
	},
	52: {
		Flex_state: uint16(28),
	},
	53: {
		Flex_state: uint16(28),
	},
	54: {
		Flex_state: uint16(4),
	},
	55: {
		Flex_state: uint16(4),
	},
	56: {
		Flex_state: uint16(29),
	},
	57: {
		Flex_state: uint16(29),
	},
	58: {
		Flex_state: uint16(29),
	},
	59: {
		Flex_state: uint16(29),
	},
	60: {
		Flex_state: uint16(29),
	},
	61: {},
	62: {
		Flex_state: uint16(28),
	},
	63: {
		Flex_state: uint16(28),
	},
	64: {},
	65: {},
	66: {},
	67: {
		Flex_state: uint16(29),
	},
	68: {},
	69: {},
	70: {},
	71: {},
	72: {
		Flex_state: uint16(29),
	},
	73: {
		Flex_state: uint16(29),
	},
	74: {},
	75: {
		Flex_state: uint16(29),
	},
	76: {
		Flex_state: uint16(28),
	},
	77: {},
	78: {},
	79: {
		Flex_state: uint16(29),
	},
	80: {},
	81: {},
	82: {
		Flex_state: uint16(28),
	},
	83:  {},
	84:  {},
	85:  {},
	86:  {},
	87:  {},
	88:  {},
	89:  {},
	90:  {},
	91:  {},
	92:  {},
	93:  {},
	94:  {},
	95:  {},
	96:  {},
	97:  {},
	98:  {},
	99:  {},
	100: {},
	101: {},
	102: {
		Flex_state: uint16(28),
	},
	103: {},
	104: {
		Flex_state: uint16(28),
	},
	105: {
		Flex_state: uint16(2),
	},
	106: {},
	107: {},
	108: {},
	109: {
		Flex_state: uint16(2),
	},
	110: {},
	111: {},
	112: {},
	113: {},
	114: {},
	115: {
		Flex_state: uint16(2),
	},
	116: {
		Flex_state: uint16(151),
	},
	117: {},
	118: {},
	119: {},
	120: {},
	121: {},
	122: {},
	123: {},
	124: {
		Flex_state: uint16(25),
	},
	125: {
		Flex_state: uint16(3),
	},
	126: {
		Flex_state: uint16(2),
	},
	127: {
		Flex_state: uint16(2),
	},
	128: {},
	129: {},
	130: {},
	131: {},
	132: {
		Flex_state: uint16(29),
	},
	133: {
		Flex_state: uint16(29),
	},
	134: {
		Flex_state: uint16(3),
	},
	135: {
		Flex_state: uint16(-libc.Int32FromInt32(1)),
	},
	136: {
		Flex_state: uint16(-libc.Int32FromInt32(1)),
	},
}

var ts_parse_table = [3][65]uint16_t{
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
		27: uint16(3),
		29: uint16(1),
		31: uint16(1),
		32: uint16(1),
		33: uint16(1),
		35: uint16(1),
		36: uint16(1),
		37: uint16(1),
		39: uint16(5),
		40: uint16(7),
	},
	1: {
		0:  uint16(9),
		6:  uint16(11),
		7:  uint16(13),
		27: uint16(3),
		33: uint16(15),
		36: uint16(17),
		37: uint16(19),
		39: uint16(5),
		40: uint16(7),
		41: uint16(130),
		42: uint16(81),
		43: uint16(114),
		44: uint16(114),
		45: uint16(114),
		46: uint16(114),
		56: uint16(1),
	},
	2: {
		0:  uint16(21),
		1:  uint16(21),
		3:  uint16(23),
		4:  uint16(25),
		8:  uint16(27),
		11: uint16(29),
		12: uint16(29),
		13: uint16(29),
		14: uint16(29),
		16: uint16(31),
		19: uint16(33),
		27: uint16(35),
		29: uint16(37),
		30: uint16(39),
		31: uint16(41),
		32: uint16(41),
		33: uint16(43),
		34: uint16(43),
		35: uint16(43),
		36: uint16(45),
		37: uint16(45),
		39: uint16(5),
		40: uint16(7),
		47: uint16(69),
		48: uint16(100),
		49: uint16(100),
		50: uint16(18),
		51: uint16(17),
		52: uint16(7),
		53: uint16(10),
		54: uint16(11),
		55: uint16(10),
		56: uint16(2),
		63: uint16(20),
	},
}

var ts_small_parse_table = [3601]uint16_t{
	0:    uint16(19),
	1:    uint16(5),
	2:    uint16(1),
	3:    uint16(aux_sym_line_comment_token2),
	4:    uint16(7),
	5:    uint16(1),
	6:    uint16(sym_block_comment),
	7:    uint16(25),
	8:    uint16(1),
	9:    uint16(anon_sym_LPAREN),
	10:   uint16(27),
	11:   uint16(1),
	12:   uint16(anon_sym_LBRACE),
	13:   uint16(31),
	14:   uint16(1),
	15:   uint16(anon_sym_LBRACK),
	16:   uint16(33),
	17:   uint16(1),
	18:   uint16(anon_sym_STAR),
	19:   uint16(39),
	20:   uint16(1),
	21:   uint16(sym_float),
	22:   uint16(49),
	23:   uint16(1),
	24:   uint16(anon_sym_POUND),
	25:   uint16(51),
	26:   uint16(1),
	27:   uint16(aux_sym_int_token2),
	28:   uint16(3),
	29:   uint16(1),
	30:   uint16(sym_line_comment),
	31:   uint16(40),
	32:   uint16(1),
	33:   uint16(sym_reg),
	34:   uint16(78),
	35:   uint16(1),
	36:   uint16(sym_int),
	37:   uint16(91),
	38:   uint16(1),
	39:   uint16(sym__expr),
	40:   uint16(47),
	41:   uint16(2),
	43:   uint16(anon_sym_LF),
	44:   uint16(53),
	45:   uint16(2),
	46:   uint16(aux_sym_string_token1),
	47:   uint16(aux_sym_string_token2),
	48:   uint16(57),
	49:   uint16(2),
	50:   uint16(sym_meta_ident),
	51:   uint16(sym__ident),
	52:   uint16(55),
	53:   uint16(3),
	54:   uint16(sym_word),
	55:   uint16(sym__reg),
	56:   uint16(sym_address),
	57:   uint16(29),
	58:   uint16(4),
	59:   uint16(anon_sym_byte),
	60:   uint16(anon_sym_word),
	61:   uint16(anon_sym_dword),
	62:   uint16(anon_sym_qword),
	63:   uint16(100),
	64:   uint16(4),
	65:   uint16(sym_list),
	66:   uint16(sym_ptr),
	67:   uint16(sym_string),
	68:   uint16(sym_ident),
	69:   uint16(19),
	70:   uint16(5),
	71:   uint16(1),
	72:   uint16(aux_sym_line_comment_token2),
	73:   uint16(7),
	74:   uint16(1),
	75:   uint16(sym_block_comment),
	76:   uint16(25),
	77:   uint16(1),
	78:   uint16(anon_sym_LPAREN),
	79:   uint16(27),
	80:   uint16(1),
	81:   uint16(anon_sym_LBRACE),
	82:   uint16(31),
	83:   uint16(1),
	84:   uint16(anon_sym_LBRACK),
	85:   uint16(33),
	86:   uint16(1),
	87:   uint16(anon_sym_STAR),
	88:   uint16(39),
	89:   uint16(1),
	90:   uint16(sym_float),
	91:   uint16(49),
	92:   uint16(1),
	93:   uint16(anon_sym_POUND),
	94:   uint16(51),
	95:   uint16(1),
	96:   uint16(aux_sym_int_token2),
	97:   uint16(4),
	98:   uint16(1),
	99:   uint16(sym_line_comment),
	100:  uint16(40),
	101:  uint16(1),
	102:  uint16(sym_reg),
	103:  uint16(78),
	104:  uint16(1),
	105:  uint16(sym_int),
	106:  uint16(91),
	107:  uint16(1),
	108:  uint16(sym__expr),
	109:  uint16(53),
	110:  uint16(2),
	111:  uint16(aux_sym_string_token1),
	112:  uint16(aux_sym_string_token2),
	113:  uint16(57),
	114:  uint16(2),
	115:  uint16(sym_meta_ident),
	116:  uint16(sym__ident),
	117:  uint16(59),
	118:  uint16(2),
	120:  uint16(anon_sym_LF),
	121:  uint16(55),
	122:  uint16(3),
	123:  uint16(sym_word),
	124:  uint16(sym__reg),
	125:  uint16(sym_address),
	126:  uint16(29),
	127:  uint16(4),
	128:  uint16(anon_sym_byte),
	129:  uint16(anon_sym_word),
	130:  uint16(anon_sym_dword),
	131:  uint16(anon_sym_qword),
	132:  uint16(100),
	133:  uint16(4),
	134:  uint16(sym_list),
	135:  uint16(sym_ptr),
	136:  uint16(sym_string),
	137:  uint16(sym_ident),
	138:  uint16(18),
	139:  uint16(5),
	140:  uint16(1),
	141:  uint16(aux_sym_line_comment_token2),
	142:  uint16(7),
	143:  uint16(1),
	144:  uint16(sym_block_comment),
	145:  uint16(25),
	146:  uint16(1),
	147:  uint16(anon_sym_LPAREN),
	148:  uint16(27),
	149:  uint16(1),
	150:  uint16(anon_sym_LBRACE),
	151:  uint16(31),
	152:  uint16(1),
	153:  uint16(anon_sym_LBRACK),
	154:  uint16(33),
	155:  uint16(1),
	156:  uint16(anon_sym_STAR),
	157:  uint16(39),
	158:  uint16(1),
	159:  uint16(sym_float),
	160:  uint16(49),
	161:  uint16(1),
	162:  uint16(anon_sym_POUND),
	163:  uint16(51),
	164:  uint16(1),
	165:  uint16(aux_sym_int_token2),
	166:  uint16(5),
	167:  uint16(1),
	168:  uint16(sym_line_comment),
	169:  uint16(40),
	170:  uint16(1),
	171:  uint16(sym_reg),
	172:  uint16(78),
	173:  uint16(1),
	174:  uint16(sym_int),
	175:  uint16(91),
	176:  uint16(1),
	177:  uint16(sym__expr),
	178:  uint16(53),
	179:  uint16(2),
	180:  uint16(aux_sym_string_token1),
	181:  uint16(aux_sym_string_token2),
	182:  uint16(57),
	183:  uint16(2),
	184:  uint16(sym_meta_ident),
	185:  uint16(sym__ident),
	186:  uint16(55),
	187:  uint16(3),
	188:  uint16(sym_word),
	189:  uint16(sym__reg),
	190:  uint16(sym_address),
	191:  uint16(29),
	192:  uint16(4),
	193:  uint16(anon_sym_byte),
	194:  uint16(anon_sym_word),
	195:  uint16(anon_sym_dword),
	196:  uint16(anon_sym_qword),
	197:  uint16(100),
	198:  uint16(4),
	199:  uint16(sym_list),
	200:  uint16(sym_ptr),
	201:  uint16(sym_string),
	202:  uint16(sym_ident),
	203:  uint16(5),
	204:  uint16(5),
	205:  uint16(1),
	206:  uint16(aux_sym_line_comment_token2),
	207:  uint16(7),
	208:  uint16(1),
	209:  uint16(sym_block_comment),
	210:  uint16(6),
	211:  uint16(1),
	212:  uint16(sym_line_comment),
	213:  uint16(63),
	214:  uint16(9),
	215:  uint16(anon_sym_DASH),
	216:  uint16(anon_sym_SLASH),
	217:  uint16(anon_sym_PERCENT),
	218:  uint16(aux_sym_int_token2),
	219:  uint16(sym_word),
	220:  uint16(sym__reg),
	221:  uint16(sym_address),
	222:  uint16(sym_meta_ident),
	223:  uint16(sym__ident),
	224:  uint16(61),
	225:  uint16(12),
	227:  uint16(anon_sym_LF),
	228:  uint16(anon_sym_COMMA),
	229:  uint16(anon_sym_LPAREN),
	230:  uint16(anon_sym_PLUS),
	231:  uint16(anon_sym_STAR),
	232:  uint16(anon_sym_PIPE),
	233:  uint16(anon_sym_CARET),
	234:  uint16(anon_sym_AMP),
	235:  uint16(anon_sym_POUND),
	236:  uint16(aux_sym_string_token1),
	237:  uint16(aux_sym_string_token2),
	238:  uint16(8),
	239:  uint16(5),
	240:  uint16(1),
	241:  uint16(aux_sym_line_comment_token2),
	242:  uint16(7),
	243:  uint16(1),
	244:  uint16(sym_block_comment),
	245:  uint16(68),
	246:  uint16(1),
	247:  uint16(anon_sym_COMMA),
	248:  uint16(70),
	249:  uint16(1),
	250:  uint16(anon_sym_LPAREN),
	251:  uint16(7),
	252:  uint16(1),
	253:  uint16(sym_line_comment),
	254:  uint16(65),
	255:  uint16(2),
	257:  uint16(anon_sym_LF),
	258:  uint16(74),
	259:  uint16(8),
	260:  uint16(anon_sym_PLUS),
	261:  uint16(anon_sym_STAR),
	262:  uint16(anon_sym_PIPE),
	263:  uint16(anon_sym_CARET),
	264:  uint16(anon_sym_AMP),
	265:  uint16(anon_sym_POUND),
	266:  uint16(aux_sym_string_token1),
	267:  uint16(aux_sym_string_token2),
	268:  uint16(72),
	269:  uint16(9),
	270:  uint16(anon_sym_DASH),
	271:  uint16(anon_sym_SLASH),
	272:  uint16(anon_sym_PERCENT),
	273:  uint16(aux_sym_int_token2),
	274:  uint16(sym_word),
	275:  uint16(sym__reg),
	276:  uint16(sym_address),
	277:  uint16(sym_meta_ident),
	278:  uint16(sym__ident),
	279:  uint16(5),
	280:  uint16(5),
	281:  uint16(1),
	282:  uint16(aux_sym_line_comment_token2),
	283:  uint16(7),
	284:  uint16(1),
	285:  uint16(sym_block_comment),
	286:  uint16(8),
	287:  uint16(1),
	288:  uint16(sym_line_comment),
	289:  uint16(78),
	290:  uint16(9),
	291:  uint16(anon_sym_DASH),
	292:  uint16(anon_sym_SLASH),
	293:  uint16(anon_sym_PERCENT),
	294:  uint16(aux_sym_int_token2),
	295:  uint16(sym_word),
	296:  uint16(sym__reg),
	297:  uint16(sym_address),
	298:  uint16(sym_meta_ident),
	299:  uint16(sym__ident),
	300:  uint16(76),
	301:  uint16(12),
	303:  uint16(anon_sym_LF),
	304:  uint16(anon_sym_COMMA),
	305:  uint16(anon_sym_LPAREN),
	306:  uint16(anon_sym_PLUS),
	307:  uint16(anon_sym_STAR),
	308:  uint16(anon_sym_PIPE),
	309:  uint16(anon_sym_CARET),
	310:  uint16(anon_sym_AMP),
	311:  uint16(anon_sym_POUND),
	312:  uint16(aux_sym_string_token1),
	313:  uint16(aux_sym_string_token2),
	314:  uint16(5),
	315:  uint16(5),
	316:  uint16(1),
	317:  uint16(aux_sym_line_comment_token2),
	318:  uint16(7),
	319:  uint16(1),
	320:  uint16(sym_block_comment),
	321:  uint16(9),
	322:  uint16(1),
	323:  uint16(sym_line_comment),
	324:  uint16(82),
	325:  uint16(9),
	326:  uint16(anon_sym_DASH),
	327:  uint16(anon_sym_SLASH),
	328:  uint16(anon_sym_PERCENT),
	329:  uint16(aux_sym_int_token2),
	330:  uint16(sym_word),
	331:  uint16(sym__reg),
	332:  uint16(sym_address),
	333:  uint16(sym_meta_ident),
	334:  uint16(sym__ident),
	335:  uint16(80),
	336:  uint16(11),
	338:  uint16(anon_sym_LF),
	339:  uint16(anon_sym_COMMA),
	340:  uint16(anon_sym_PLUS),
	341:  uint16(anon_sym_STAR),
	342:  uint16(anon_sym_PIPE),
	343:  uint16(anon_sym_CARET),
	344:  uint16(anon_sym_AMP),
	345:  uint16(anon_sym_POUND),
	346:  uint16(aux_sym_string_token1),
	347:  uint16(aux_sym_string_token2),
	348:  uint16(7),
	349:  uint16(5),
	350:  uint16(1),
	351:  uint16(aux_sym_line_comment_token2),
	352:  uint16(7),
	353:  uint16(1),
	354:  uint16(sym_block_comment),
	355:  uint16(68),
	356:  uint16(1),
	357:  uint16(anon_sym_COMMA),
	358:  uint16(10),
	359:  uint16(1),
	360:  uint16(sym_line_comment),
	361:  uint16(65),
	362:  uint16(2),
	364:  uint16(anon_sym_LF),
	365:  uint16(74),
	366:  uint16(8),
	367:  uint16(anon_sym_PLUS),
	368:  uint16(anon_sym_STAR),
	369:  uint16(anon_sym_PIPE),
	370:  uint16(anon_sym_CARET),
	371:  uint16(anon_sym_AMP),
	372:  uint16(anon_sym_POUND),
	373:  uint16(aux_sym_string_token1),
	374:  uint16(aux_sym_string_token2),
	375:  uint16(72),
	376:  uint16(9),
	377:  uint16(anon_sym_DASH),
	378:  uint16(anon_sym_SLASH),
	379:  uint16(anon_sym_PERCENT),
	380:  uint16(aux_sym_int_token2),
	381:  uint16(sym_word),
	382:  uint16(sym__reg),
	383:  uint16(sym_address),
	384:  uint16(sym_meta_ident),
	385:  uint16(sym__ident),
	386:  uint16(5),
	387:  uint16(5),
	388:  uint16(1),
	389:  uint16(aux_sym_line_comment_token2),
	390:  uint16(7),
	391:  uint16(1),
	392:  uint16(sym_block_comment),
	393:  uint16(11),
	394:  uint16(1),
	395:  uint16(sym_line_comment),
	396:  uint16(86),
	397:  uint16(9),
	398:  uint16(anon_sym_DASH),
	399:  uint16(anon_sym_SLASH),
	400:  uint16(anon_sym_PERCENT),
	401:  uint16(aux_sym_int_token2),
	402:  uint16(sym_word),
	403:  uint16(sym__reg),
	404:  uint16(sym_address),
	405:  uint16(sym_meta_ident),
	406:  uint16(sym__ident),
	407:  uint16(84),
	408:  uint16(11),
	410:  uint16(anon_sym_LF),
	411:  uint16(anon_sym_COMMA),
	412:  uint16(anon_sym_PLUS),
	413:  uint16(anon_sym_STAR),
	414:  uint16(anon_sym_PIPE),
	415:  uint16(anon_sym_CARET),
	416:  uint16(anon_sym_AMP),
	417:  uint16(anon_sym_POUND),
	418:  uint16(aux_sym_string_token1),
	419:  uint16(aux_sym_string_token2),
	420:  uint16(5),
	421:  uint16(5),
	422:  uint16(1),
	423:  uint16(aux_sym_line_comment_token2),
	424:  uint16(7),
	425:  uint16(1),
	426:  uint16(sym_block_comment),
	427:  uint16(12),
	428:  uint16(1),
	429:  uint16(sym_line_comment),
	430:  uint16(90),
	431:  uint16(9),
	432:  uint16(anon_sym_DASH),
	433:  uint16(anon_sym_SLASH),
	434:  uint16(anon_sym_PERCENT),
	435:  uint16(aux_sym_int_token2),
	436:  uint16(sym_word),
	437:  uint16(sym__reg),
	438:  uint16(sym_address),
	439:  uint16(sym_meta_ident),
	440:  uint16(sym__ident),
	441:  uint16(88),
	442:  uint16(11),
	444:  uint16(anon_sym_LF),
	445:  uint16(anon_sym_COMMA),
	446:  uint16(anon_sym_PLUS),
	447:  uint16(anon_sym_STAR),
	448:  uint16(anon_sym_PIPE),
	449:  uint16(anon_sym_CARET),
	450:  uint16(anon_sym_AMP),
	451:  uint16(anon_sym_POUND),
	452:  uint16(aux_sym_string_token1),
	453:  uint16(aux_sym_string_token2),
	454:  uint16(8),
	455:  uint16(5),
	456:  uint16(1),
	457:  uint16(aux_sym_line_comment_token2),
	458:  uint16(7),
	459:  uint16(1),
	460:  uint16(sym_block_comment),
	461:  uint16(96),
	462:  uint16(1),
	463:  uint16(anon_sym_PIPE),
	464:  uint16(98),
	465:  uint16(1),
	466:  uint16(anon_sym_CARET),
	467:  uint16(100),
	468:  uint16(1),
	469:  uint16(anon_sym_AMP),
	470:  uint16(13),
	471:  uint16(1),
	472:  uint16(sym_line_comment),
	473:  uint16(92),
	474:  uint16(7),
	476:  uint16(anon_sym_LF),
	477:  uint16(anon_sym_PLUS),
	478:  uint16(anon_sym_STAR),
	479:  uint16(anon_sym_POUND),
	480:  uint16(aux_sym_string_token1),
	481:  uint16(aux_sym_string_token2),
	482:  uint16(94),
	483:  uint16(9),
	484:  uint16(anon_sym_DASH),
	485:  uint16(anon_sym_SLASH),
	486:  uint16(anon_sym_PERCENT),
	487:  uint16(aux_sym_int_token2),
	488:  uint16(sym_word),
	489:  uint16(sym__reg),
	490:  uint16(sym_address),
	491:  uint16(sym_meta_ident),
	492:  uint16(sym__ident),
	493:  uint16(10),
	494:  uint16(5),
	495:  uint16(1),
	496:  uint16(aux_sym_line_comment_token2),
	497:  uint16(7),
	498:  uint16(1),
	499:  uint16(sym_block_comment),
	500:  uint16(96),
	501:  uint16(1),
	502:  uint16(anon_sym_PIPE),
	503:  uint16(98),
	504:  uint16(1),
	505:  uint16(anon_sym_CARET),
	506:  uint16(100),
	507:  uint16(1),
	508:  uint16(anon_sym_AMP),
	509:  uint16(102),
	510:  uint16(1),
	511:  uint16(anon_sym_STAR),
	512:  uint16(14),
	513:  uint16(1),
	514:  uint16(sym_line_comment),
	515:  uint16(104),
	516:  uint16(2),
	517:  uint16(anon_sym_SLASH),
	518:  uint16(anon_sym_PERCENT),
	519:  uint16(92),
	520:  uint16(6),
	522:  uint16(anon_sym_LF),
	523:  uint16(anon_sym_PLUS),
	524:  uint16(anon_sym_POUND),
	525:  uint16(aux_sym_string_token1),
	526:  uint16(aux_sym_string_token2),
	527:  uint16(94),
	528:  uint16(7),
	529:  uint16(anon_sym_DASH),
	530:  uint16(aux_sym_int_token2),
	531:  uint16(sym_word),
	532:  uint16(sym__reg),
	533:  uint16(sym_address),
	534:  uint16(sym_meta_ident),
	535:  uint16(sym__ident),
	536:  uint16(5),
	537:  uint16(5),
	538:  uint16(1),
	539:  uint16(aux_sym_line_comment_token2),
	540:  uint16(7),
	541:  uint16(1),
	542:  uint16(sym_block_comment),
	543:  uint16(15),
	544:  uint16(1),
	545:  uint16(sym_line_comment),
	546:  uint16(94),
	547:  uint16(9),
	548:  uint16(anon_sym_DASH),
	549:  uint16(anon_sym_SLASH),
	550:  uint16(anon_sym_PERCENT),
	551:  uint16(aux_sym_int_token2),
	552:  uint16(sym_word),
	553:  uint16(sym__reg),
	554:  uint16(sym_address),
	555:  uint16(sym_meta_ident),
	556:  uint16(sym__ident),
	557:  uint16(92),
	558:  uint16(10),
	560:  uint16(anon_sym_LF),
	561:  uint16(anon_sym_PLUS),
	562:  uint16(anon_sym_STAR),
	563:  uint16(anon_sym_PIPE),
	564:  uint16(anon_sym_CARET),
	565:  uint16(anon_sym_AMP),
	566:  uint16(anon_sym_POUND),
	567:  uint16(aux_sym_string_token1),
	568:  uint16(aux_sym_string_token2),
	569:  uint16(7),
	570:  uint16(5),
	571:  uint16(1),
	572:  uint16(aux_sym_line_comment_token2),
	573:  uint16(7),
	574:  uint16(1),
	575:  uint16(sym_block_comment),
	576:  uint16(98),
	577:  uint16(1),
	578:  uint16(anon_sym_CARET),
	579:  uint16(100),
	580:  uint16(1),
	581:  uint16(anon_sym_AMP),
	582:  uint16(16),
	583:  uint16(1),
	584:  uint16(sym_line_comment),
	585:  uint16(92),
	586:  uint16(8),
	588:  uint16(anon_sym_LF),
	589:  uint16(anon_sym_PLUS),
	590:  uint16(anon_sym_STAR),
	591:  uint16(anon_sym_PIPE),
	592:  uint16(anon_sym_POUND),
	593:  uint16(aux_sym_string_token1),
	594:  uint16(aux_sym_string_token2),
	595:  uint16(94),
	596:  uint16(9),
	597:  uint16(anon_sym_DASH),
	598:  uint16(anon_sym_SLASH),
	599:  uint16(anon_sym_PERCENT),
	600:  uint16(aux_sym_int_token2),
	601:  uint16(sym_word),
	602:  uint16(sym__reg),
	603:  uint16(sym_address),
	604:  uint16(sym_meta_ident),
	605:  uint16(sym__ident),
	606:  uint16(5),
	607:  uint16(5),
	608:  uint16(1),
	609:  uint16(aux_sym_line_comment_token2),
	610:  uint16(7),
	611:  uint16(1),
	612:  uint16(sym_block_comment),
	613:  uint16(17),
	614:  uint16(1),
	615:  uint16(sym_line_comment),
	616:  uint16(72),
	617:  uint16(9),
	618:  uint16(anon_sym_DASH),
	619:  uint16(anon_sym_SLASH),
	620:  uint16(anon_sym_PERCENT),
	621:  uint16(aux_sym_int_token2),
	622:  uint16(sym_word),
	623:  uint16(sym__reg),
	624:  uint16(sym_address),
	625:  uint16(sym_meta_ident),
	626:  uint16(sym__ident),
	627:  uint16(74),
	628:  uint16(10),
	630:  uint16(anon_sym_LF),
	631:  uint16(anon_sym_PLUS),
	632:  uint16(anon_sym_STAR),
	633:  uint16(anon_sym_PIPE),
	634:  uint16(anon_sym_CARET),
	635:  uint16(anon_sym_AMP),
	636:  uint16(anon_sym_POUND),
	637:  uint16(aux_sym_string_token1),
	638:  uint16(aux_sym_string_token2),
	639:  uint16(12),
	640:  uint16(5),
	641:  uint16(1),
	642:  uint16(aux_sym_line_comment_token2),
	643:  uint16(7),
	644:  uint16(1),
	645:  uint16(sym_block_comment),
	646:  uint16(96),
	647:  uint16(1),
	648:  uint16(anon_sym_PIPE),
	649:  uint16(98),
	650:  uint16(1),
	651:  uint16(anon_sym_CARET),
	652:  uint16(100),
	653:  uint16(1),
	654:  uint16(anon_sym_AMP),
	655:  uint16(102),
	656:  uint16(1),
	657:  uint16(anon_sym_STAR),
	658:  uint16(108),
	659:  uint16(1),
	660:  uint16(anon_sym_DASH),
	661:  uint16(110),
	662:  uint16(1),
	663:  uint16(anon_sym_PLUS),
	664:  uint16(18),
	665:  uint16(1),
	666:  uint16(sym_line_comment),
	667:  uint16(104),
	668:  uint16(2),
	669:  uint16(anon_sym_SLASH),
	670:  uint16(anon_sym_PERCENT),
	671:  uint16(106),
	672:  uint16(5),
	674:  uint16(anon_sym_LF),
	675:  uint16(anon_sym_POUND),
	676:  uint16(aux_sym_string_token1),
	677:  uint16(aux_sym_string_token2),
	678:  uint16(112),
	679:  uint16(6),
	680:  uint16(aux_sym_int_token2),
	681:  uint16(sym_word),
	682:  uint16(sym__reg),
	683:  uint16(sym_address),
	684:  uint16(sym_meta_ident),
	685:  uint16(sym__ident),
	686:  uint16(6),
	687:  uint16(5),
	688:  uint16(1),
	689:  uint16(aux_sym_line_comment_token2),
	690:  uint16(7),
	691:  uint16(1),
	692:  uint16(sym_block_comment),
	693:  uint16(100),
	694:  uint16(1),
	695:  uint16(anon_sym_AMP),
	696:  uint16(19),
	697:  uint16(1),
	698:  uint16(sym_line_comment),
	699:  uint16(92),
	700:  uint16(9),
	702:  uint16(anon_sym_LF),
	703:  uint16(anon_sym_PLUS),
	704:  uint16(anon_sym_STAR),
	705:  uint16(anon_sym_PIPE),
	706:  uint16(anon_sym_CARET),
	707:  uint16(anon_sym_POUND),
	708:  uint16(aux_sym_string_token1),
	709:  uint16(aux_sym_string_token2),
	710:  uint16(94),
	711:  uint16(9),
	712:  uint16(anon_sym_DASH),
	713:  uint16(anon_sym_SLASH),
	714:  uint16(anon_sym_PERCENT),
	715:  uint16(aux_sym_int_token2),
	716:  uint16(sym_word),
	717:  uint16(sym__reg),
	718:  uint16(sym_address),
	719:  uint16(sym_meta_ident),
	720:  uint16(sym__ident),
	721:  uint16(13),
	722:  uint16(5),
	723:  uint16(1),
	724:  uint16(aux_sym_line_comment_token2),
	725:  uint16(7),
	726:  uint16(1),
	727:  uint16(sym_block_comment),
	728:  uint16(35),
	729:  uint16(1),
	730:  uint16(anon_sym_POUND),
	731:  uint16(37),
	732:  uint16(1),
	733:  uint16(aux_sym_int_token2),
	734:  uint16(11),
	735:  uint16(1),
	736:  uint16(sym_reg),
	737:  uint16(18),
	738:  uint16(1),
	739:  uint16(sym__tc_expr),
	740:  uint16(20),
	741:  uint16(1),
	742:  uint16(sym_line_comment),
	743:  uint16(21),
	744:  uint16(1),
	745:  uint16(aux_sym_instruction_repeat2),
	746:  uint16(41),
	747:  uint16(2),
	748:  uint16(aux_sym_string_token1),
	749:  uint16(aux_sym_string_token2),
	750:  uint16(45),
	751:  uint16(2),
	752:  uint16(sym_meta_ident),
	753:  uint16(sym__ident),
	754:  uint16(114),
	755:  uint16(2),
	757:  uint16(anon_sym_LF),
	758:  uint16(43),
	759:  uint16(3),
	760:  uint16(sym_word),
	761:  uint16(sym__reg),
	762:  uint16(sym_address),
	763:  uint16(17),
	764:  uint16(4),
	765:  uint16(sym_tc_infix),
	766:  uint16(sym_int),
	767:  uint16(sym_string),
	768:  uint16(sym_ident),
	769:  uint16(12),
	770:  uint16(5),
	771:  uint16(1),
	772:  uint16(aux_sym_line_comment_token2),
	773:  uint16(7),
	774:  uint16(1),
	775:  uint16(sym_block_comment),
	776:  uint16(118),
	777:  uint16(1),
	778:  uint16(anon_sym_POUND),
	779:  uint16(121),
	780:  uint16(1),
	781:  uint16(aux_sym_int_token2),
	782:  uint16(11),
	783:  uint16(1),
	784:  uint16(sym_reg),
	785:  uint16(18),
	786:  uint16(1),
	787:  uint16(sym__tc_expr),
	788:  uint16(116),
	789:  uint16(2),
	791:  uint16(anon_sym_LF),
	792:  uint16(124),
	793:  uint16(2),
	794:  uint16(aux_sym_string_token1),
	795:  uint16(aux_sym_string_token2),
	796:  uint16(130),
	797:  uint16(2),
	798:  uint16(sym_meta_ident),
	799:  uint16(sym__ident),
	800:  uint16(21),
	801:  uint16(2),
	802:  uint16(sym_line_comment),
	803:  uint16(aux_sym_instruction_repeat2),
	804:  uint16(127),
	805:  uint16(3),
	806:  uint16(sym_word),
	807:  uint16(sym__reg),
	808:  uint16(sym_address),
	809:  uint16(17),
	810:  uint16(4),
	811:  uint16(sym_tc_infix),
	812:  uint16(sym_int),
	813:  uint16(sym_string),
	814:  uint16(sym_ident),
	815:  uint16(15),
	816:  uint16(5),
	817:  uint16(1),
	818:  uint16(aux_sym_line_comment_token2),
	819:  uint16(7),
	820:  uint16(1),
	821:  uint16(sym_block_comment),
	822:  uint16(49),
	823:  uint16(1),
	824:  uint16(anon_sym_POUND),
	825:  uint16(51),
	826:  uint16(1),
	827:  uint16(aux_sym_int_token2),
	828:  uint16(135),
	829:  uint16(1),
	830:  uint16(anon_sym_COLON),
	831:  uint16(137),
	832:  uint16(1),
	833:  uint16(sym_float),
	834:  uint16(22),
	835:  uint16(1),
	836:  uint16(sym_line_comment),
	837:  uint16(40),
	838:  uint16(1),
	839:  uint16(sym_reg),
	840:  uint16(65),
	841:  uint16(1),
	842:  uint16(sym_string),
	843:  uint16(74),
	844:  uint16(1),
	845:  uint16(sym_int),
	846:  uint16(106),
	847:  uint16(1),
	848:  uint16(sym_ident),
	849:  uint16(53),
	850:  uint16(2),
	851:  uint16(aux_sym_string_token1),
	852:  uint16(aux_sym_string_token2),
	853:  uint16(57),
	854:  uint16(2),
	855:  uint16(sym_meta_ident),
	856:  uint16(sym__ident),
	857:  uint16(133),
	858:  uint16(2),
	860:  uint16(anon_sym_LF),
	861:  uint16(55),
	862:  uint16(3),
	863:  uint16(sym_word),
	864:  uint16(sym__reg),
	865:  uint16(sym_address),
	866:  uint16(14),
	867:  uint16(5),
	868:  uint16(1),
	869:  uint16(aux_sym_line_comment_token2),
	870:  uint16(7),
	871:  uint16(1),
	872:  uint16(sym_block_comment),
	873:  uint16(49),
	874:  uint16(1),
	875:  uint16(anon_sym_POUND),
	876:  uint16(51),
	877:  uint16(1),
	878:  uint16(aux_sym_int_token2),
	879:  uint16(137),
	880:  uint16(1),
	881:  uint16(sym_float),
	882:  uint16(23),
	883:  uint16(1),
	884:  uint16(sym_line_comment),
	885:  uint16(40),
	886:  uint16(1),
	887:  uint16(sym_reg),
	888:  uint16(65),
	889:  uint16(1),
	890:  uint16(sym_string),
	891:  uint16(74),
	892:  uint16(1),
	893:  uint16(sym_int),
	894:  uint16(106),
	895:  uint16(1),
	896:  uint16(sym_ident),
	897:  uint16(53),
	898:  uint16(2),
	899:  uint16(aux_sym_string_token1),
	900:  uint16(aux_sym_string_token2),
	901:  uint16(57),
	902:  uint16(2),
	903:  uint16(sym_meta_ident),
	904:  uint16(sym__ident),
	905:  uint16(133),
	906:  uint16(2),
	908:  uint16(anon_sym_LF),
	909:  uint16(55),
	910:  uint16(3),
	911:  uint16(sym_word),
	912:  uint16(sym__reg),
	913:  uint16(sym_address),
	914:  uint16(11),
	915:  uint16(5),
	916:  uint16(1),
	917:  uint16(aux_sym_line_comment_token2),
	918:  uint16(7),
	919:  uint16(1),
	920:  uint16(sym_block_comment),
	921:  uint16(49),
	922:  uint16(1),
	923:  uint16(anon_sym_POUND),
	924:  uint16(51),
	925:  uint16(1),
	926:  uint16(aux_sym_int_token2),
	927:  uint16(24),
	928:  uint16(1),
	929:  uint16(sym_line_comment),
	930:  uint16(40),
	931:  uint16(1),
	932:  uint16(sym_reg),
	933:  uint16(49),
	934:  uint16(1),
	935:  uint16(sym__tc_expr),
	936:  uint16(53),
	937:  uint16(2),
	938:  uint16(aux_sym_string_token1),
	939:  uint16(aux_sym_string_token2),
	940:  uint16(57),
	941:  uint16(2),
	942:  uint16(sym_meta_ident),
	943:  uint16(sym__ident),
	944:  uint16(55),
	945:  uint16(3),
	946:  uint16(sym_word),
	947:  uint16(sym__reg),
	948:  uint16(sym_address),
	949:  uint16(43),
	950:  uint16(4),
	951:  uint16(sym_tc_infix),
	952:  uint16(sym_int),
	953:  uint16(sym_string),
	954:  uint16(sym_ident),
	955:  uint16(11),
	956:  uint16(5),
	957:  uint16(1),
	958:  uint16(aux_sym_line_comment_token2),
	959:  uint16(7),
	960:  uint16(1),
	961:  uint16(sym_block_comment),
	962:  uint16(35),
	963:  uint16(1),
	964:  uint16(anon_sym_POUND),
	965:  uint16(37),
	966:  uint16(1),
	967:  uint16(aux_sym_int_token2),
	968:  uint16(11),
	969:  uint16(1),
	970:  uint16(sym_reg),
	971:  uint16(14),
	972:  uint16(1),
	973:  uint16(sym__tc_expr),
	974:  uint16(25),
	975:  uint16(1),
	976:  uint16(sym_line_comment),
	977:  uint16(41),
	978:  uint16(2),
	979:  uint16(aux_sym_string_token1),
	980:  uint16(aux_sym_string_token2),
	981:  uint16(45),
	982:  uint16(2),
	983:  uint16(sym_meta_ident),
	984:  uint16(sym__ident),
	985:  uint16(43),
	986:  uint16(3),
	987:  uint16(sym_word),
	988:  uint16(sym__reg),
	989:  uint16(sym_address),
	990:  uint16(17),
	991:  uint16(4),
	992:  uint16(sym_tc_infix),
	993:  uint16(sym_int),
	994:  uint16(sym_string),
	995:  uint16(sym_ident),
	996:  uint16(11),
	997:  uint16(5),
	998:  uint16(1),
	999:  uint16(aux_sym_line_comment_token2),
	1000: uint16(7),
	1001: uint16(1),
	1002: uint16(sym_block_comment),
	1003: uint16(49),
	1004: uint16(1),
	1005: uint16(anon_sym_POUND),
	1006: uint16(51),
	1007: uint16(1),
	1008: uint16(aux_sym_int_token2),
	1009: uint16(26),
	1010: uint16(1),
	1011: uint16(sym_line_comment),
	1012: uint16(40),
	1013: uint16(1),
	1014: uint16(sym_reg),
	1015: uint16(44),
	1016: uint16(1),
	1017: uint16(sym__tc_expr),
	1018: uint16(53),
	1019: uint16(2),
	1020: uint16(aux_sym_string_token1),
	1021: uint16(aux_sym_string_token2),
	1022: uint16(57),
	1023: uint16(2),
	1024: uint16(sym_meta_ident),
	1025: uint16(sym__ident),
	1026: uint16(55),
	1027: uint16(3),
	1028: uint16(sym_word),
	1029: uint16(sym__reg),
	1030: uint16(sym_address),
	1031: uint16(43),
	1032: uint16(4),
	1033: uint16(sym_tc_infix),
	1034: uint16(sym_int),
	1035: uint16(sym_string),
	1036: uint16(sym_ident),
	1037: uint16(11),
	1038: uint16(5),
	1039: uint16(1),
	1040: uint16(aux_sym_line_comment_token2),
	1041: uint16(7),
	1042: uint16(1),
	1043: uint16(sym_block_comment),
	1044: uint16(35),
	1045: uint16(1),
	1046: uint16(anon_sym_POUND),
	1047: uint16(37),
	1048: uint16(1),
	1049: uint16(aux_sym_int_token2),
	1050: uint16(11),
	1051: uint16(1),
	1052: uint16(sym_reg),
	1053: uint16(13),
	1054: uint16(1),
	1055: uint16(sym__tc_expr),
	1056: uint16(27),
	1057: uint16(1),
	1058: uint16(sym_line_comment),
	1059: uint16(41),
	1060: uint16(2),
	1061: uint16(aux_sym_string_token1),
	1062: uint16(aux_sym_string_token2),
	1063: uint16(45),
	1064: uint16(2),
	1065: uint16(sym_meta_ident),
	1066: uint16(sym__ident),
	1067: uint16(43),
	1068: uint16(3),
	1069: uint16(sym_word),
	1070: uint16(sym__reg),
	1071: uint16(sym_address),
	1072: uint16(17),
	1073: uint16(4),
	1074: uint16(sym_tc_infix),
	1075: uint16(sym_int),
	1076: uint16(sym_string),
	1077: uint16(sym_ident),
	1078: uint16(11),
	1079: uint16(5),
	1080: uint16(1),
	1081: uint16(aux_sym_line_comment_token2),
	1082: uint16(7),
	1083: uint16(1),
	1084: uint16(sym_block_comment),
	1085: uint16(35),
	1086: uint16(1),
	1087: uint16(anon_sym_POUND),
	1088: uint16(37),
	1089: uint16(1),
	1090: uint16(aux_sym_int_token2),
	1091: uint16(11),
	1092: uint16(1),
	1093: uint16(sym_reg),
	1094: uint16(16),
	1095: uint16(1),
	1096: uint16(sym__tc_expr),
	1097: uint16(28),
	1098: uint16(1),
	1099: uint16(sym_line_comment),
	1100: uint16(41),
	1101: uint16(2),
	1102: uint16(aux_sym_string_token1),
	1103: uint16(aux_sym_string_token2),
	1104: uint16(45),
	1105: uint16(2),
	1106: uint16(sym_meta_ident),
	1107: uint16(sym__ident),
	1108: uint16(43),
	1109: uint16(3),
	1110: uint16(sym_word),
	1111: uint16(sym__reg),
	1112: uint16(sym_address),
	1113: uint16(17),
	1114: uint16(4),
	1115: uint16(sym_tc_infix),
	1116: uint16(sym_int),
	1117: uint16(sym_string),
	1118: uint16(sym_ident),
	1119: uint16(11),
	1120: uint16(5),
	1121: uint16(1),
	1122: uint16(aux_sym_line_comment_token2),
	1123: uint16(7),
	1124: uint16(1),
	1125: uint16(sym_block_comment),
	1126: uint16(35),
	1127: uint16(1),
	1128: uint16(anon_sym_POUND),
	1129: uint16(37),
	1130: uint16(1),
	1131: uint16(aux_sym_int_token2),
	1132: uint16(11),
	1133: uint16(1),
	1134: uint16(sym_reg),
	1135: uint16(19),
	1136: uint16(1),
	1137: uint16(sym__tc_expr),
	1138: uint16(29),
	1139: uint16(1),
	1140: uint16(sym_line_comment),
	1141: uint16(41),
	1142: uint16(2),
	1143: uint16(aux_sym_string_token1),
	1144: uint16(aux_sym_string_token2),
	1145: uint16(45),
	1146: uint16(2),
	1147: uint16(sym_meta_ident),
	1148: uint16(sym__ident),
	1149: uint16(43),
	1150: uint16(3),
	1151: uint16(sym_word),
	1152: uint16(sym__reg),
	1153: uint16(sym_address),
	1154: uint16(17),
	1155: uint16(4),
	1156: uint16(sym_tc_infix),
	1157: uint16(sym_int),
	1158: uint16(sym_string),
	1159: uint16(sym_ident),
	1160: uint16(11),
	1161: uint16(5),
	1162: uint16(1),
	1163: uint16(aux_sym_line_comment_token2),
	1164: uint16(7),
	1165: uint16(1),
	1166: uint16(sym_block_comment),
	1167: uint16(49),
	1168: uint16(1),
	1169: uint16(anon_sym_POUND),
	1170: uint16(51),
	1171: uint16(1),
	1172: uint16(aux_sym_int_token2),
	1173: uint16(30),
	1174: uint16(1),
	1175: uint16(sym_line_comment),
	1176: uint16(40),
	1177: uint16(1),
	1178: uint16(sym_reg),
	1179: uint16(48),
	1180: uint16(1),
	1181: uint16(sym__tc_expr),
	1182: uint16(53),
	1183: uint16(2),
	1184: uint16(aux_sym_string_token1),
	1185: uint16(aux_sym_string_token2),
	1186: uint16(57),
	1187: uint16(2),
	1188: uint16(sym_meta_ident),
	1189: uint16(sym__ident),
	1190: uint16(55),
	1191: uint16(3),
	1192: uint16(sym_word),
	1193: uint16(sym__reg),
	1194: uint16(sym_address),
	1195: uint16(43),
	1196: uint16(4),
	1197: uint16(sym_tc_infix),
	1198: uint16(sym_int),
	1199: uint16(sym_string),
	1200: uint16(sym_ident),
	1201: uint16(11),
	1202: uint16(5),
	1203: uint16(1),
	1204: uint16(aux_sym_line_comment_token2),
	1205: uint16(7),
	1206: uint16(1),
	1207: uint16(sym_block_comment),
	1208: uint16(35),
	1209: uint16(1),
	1210: uint16(anon_sym_POUND),
	1211: uint16(37),
	1212: uint16(1),
	1213: uint16(aux_sym_int_token2),
	1214: uint16(11),
	1215: uint16(1),
	1216: uint16(sym_reg),
	1217: uint16(15),
	1218: uint16(1),
	1219: uint16(sym__tc_expr),
	1220: uint16(31),
	1221: uint16(1),
	1222: uint16(sym_line_comment),
	1223: uint16(41),
	1224: uint16(2),
	1225: uint16(aux_sym_string_token1),
	1226: uint16(aux_sym_string_token2),
	1227: uint16(45),
	1228: uint16(2),
	1229: uint16(sym_meta_ident),
	1230: uint16(sym__ident),
	1231: uint16(43),
	1232: uint16(3),
	1233: uint16(sym_word),
	1234: uint16(sym__reg),
	1235: uint16(sym_address),
	1236: uint16(17),
	1237: uint16(4),
	1238: uint16(sym_tc_infix),
	1239: uint16(sym_int),
	1240: uint16(sym_string),
	1241: uint16(sym_ident),
	1242: uint16(11),
	1243: uint16(5),
	1244: uint16(1),
	1245: uint16(aux_sym_line_comment_token2),
	1246: uint16(7),
	1247: uint16(1),
	1248: uint16(sym_block_comment),
	1249: uint16(49),
	1250: uint16(1),
	1251: uint16(anon_sym_POUND),
	1252: uint16(51),
	1253: uint16(1),
	1254: uint16(aux_sym_int_token2),
	1255: uint16(32),
	1256: uint16(1),
	1257: uint16(sym_line_comment),
	1258: uint16(40),
	1259: uint16(1),
	1260: uint16(sym_reg),
	1261: uint16(46),
	1262: uint16(1),
	1263: uint16(sym__tc_expr),
	1264: uint16(53),
	1265: uint16(2),
	1266: uint16(aux_sym_string_token1),
	1267: uint16(aux_sym_string_token2),
	1268: uint16(57),
	1269: uint16(2),
	1270: uint16(sym_meta_ident),
	1271: uint16(sym__ident),
	1272: uint16(55),
	1273: uint16(3),
	1274: uint16(sym_word),
	1275: uint16(sym__reg),
	1276: uint16(sym_address),
	1277: uint16(43),
	1278: uint16(4),
	1279: uint16(sym_tc_infix),
	1280: uint16(sym_int),
	1281: uint16(sym_string),
	1282: uint16(sym_ident),
	1283: uint16(11),
	1284: uint16(5),
	1285: uint16(1),
	1286: uint16(aux_sym_line_comment_token2),
	1287: uint16(7),
	1288: uint16(1),
	1289: uint16(sym_block_comment),
	1290: uint16(49),
	1291: uint16(1),
	1292: uint16(anon_sym_POUND),
	1293: uint16(51),
	1294: uint16(1),
	1295: uint16(aux_sym_int_token2),
	1296: uint16(33),
	1297: uint16(1),
	1298: uint16(sym_line_comment),
	1299: uint16(40),
	1300: uint16(1),
	1301: uint16(sym_reg),
	1302: uint16(47),
	1303: uint16(1),
	1304: uint16(sym__tc_expr),
	1305: uint16(53),
	1306: uint16(2),
	1307: uint16(aux_sym_string_token1),
	1308: uint16(aux_sym_string_token2),
	1309: uint16(57),
	1310: uint16(2),
	1311: uint16(sym_meta_ident),
	1312: uint16(sym__ident),
	1313: uint16(55),
	1314: uint16(3),
	1315: uint16(sym_word),
	1316: uint16(sym__reg),
	1317: uint16(sym_address),
	1318: uint16(43),
	1319: uint16(4),
	1320: uint16(sym_tc_infix),
	1321: uint16(sym_int),
	1322: uint16(sym_string),
	1323: uint16(sym_ident),
	1324: uint16(11),
	1325: uint16(5),
	1326: uint16(1),
	1327: uint16(aux_sym_line_comment_token2),
	1328: uint16(7),
	1329: uint16(1),
	1330: uint16(sym_block_comment),
	1331: uint16(49),
	1332: uint16(1),
	1333: uint16(anon_sym_POUND),
	1334: uint16(51),
	1335: uint16(1),
	1336: uint16(aux_sym_int_token2),
	1337: uint16(34),
	1338: uint16(1),
	1339: uint16(sym_line_comment),
	1340: uint16(40),
	1341: uint16(1),
	1342: uint16(sym_reg),
	1343: uint16(45),
	1344: uint16(1),
	1345: uint16(sym__tc_expr),
	1346: uint16(53),
	1347: uint16(2),
	1348: uint16(aux_sym_string_token1),
	1349: uint16(aux_sym_string_token2),
	1350: uint16(57),
	1351: uint16(2),
	1352: uint16(sym_meta_ident),
	1353: uint16(sym__ident),
	1354: uint16(55),
	1355: uint16(3),
	1356: uint16(sym_word),
	1357: uint16(sym__reg),
	1358: uint16(sym_address),
	1359: uint16(43),
	1360: uint16(4),
	1361: uint16(sym_tc_infix),
	1362: uint16(sym_int),
	1363: uint16(sym_string),
	1364: uint16(sym_ident),
	1365: uint16(6),
	1366: uint16(3),
	1367: uint16(1),
	1368: uint16(anon_sym_POUND),
	1369: uint16(5),
	1370: uint16(1),
	1371: uint16(aux_sym_line_comment_token2),
	1372: uint16(7),
	1373: uint16(1),
	1374: uint16(sym_block_comment),
	1375: uint16(82),
	1376: uint16(1),
	1377: uint16(anon_sym_SLASH),
	1378: uint16(35),
	1379: uint16(1),
	1380: uint16(sym_line_comment),
	1381: uint16(80),
	1382: uint16(13),
	1384: uint16(anon_sym_LF),
	1385: uint16(anon_sym_COMMA),
	1386: uint16(anon_sym_RPAREN),
	1387: uint16(anon_sym_DASH),
	1388: uint16(anon_sym_RBRACE),
	1389: uint16(anon_sym_PLUS),
	1390: uint16(anon_sym_RBRACK),
	1391: uint16(anon_sym_STAR),
	1392: uint16(anon_sym_PERCENT),
	1393: uint16(anon_sym_PIPE),
	1394: uint16(anon_sym_CARET),
	1395: uint16(anon_sym_AMP),
	1396: uint16(6),
	1397: uint16(3),
	1398: uint16(1),
	1399: uint16(anon_sym_POUND),
	1400: uint16(5),
	1401: uint16(1),
	1402: uint16(aux_sym_line_comment_token2),
	1403: uint16(7),
	1404: uint16(1),
	1405: uint16(sym_block_comment),
	1406: uint16(78),
	1407: uint16(1),
	1408: uint16(anon_sym_SLASH),
	1409: uint16(36),
	1410: uint16(1),
	1411: uint16(sym_line_comment),
	1412: uint16(76),
	1413: uint16(12),
	1415: uint16(anon_sym_LF),
	1416: uint16(anon_sym_COMMA),
	1417: uint16(anon_sym_LPAREN),
	1418: uint16(anon_sym_DASH),
	1419: uint16(anon_sym_PLUS),
	1420: uint16(anon_sym_RBRACK),
	1421: uint16(anon_sym_STAR),
	1422: uint16(anon_sym_PERCENT),
	1423: uint16(anon_sym_PIPE),
	1424: uint16(anon_sym_CARET),
	1425: uint16(anon_sym_AMP),
	1426: uint16(14),
	1427: uint16(3),
	1428: uint16(1),
	1429: uint16(anon_sym_POUND),
	1430: uint16(5),
	1431: uint16(1),
	1432: uint16(aux_sym_line_comment_token2),
	1433: uint16(7),
	1434: uint16(1),
	1435: uint16(sym_block_comment),
	1436: uint16(11),
	1437: uint16(1),
	1438: uint16(anon_sym_label),
	1439: uint16(13),
	1440: uint16(1),
	1441: uint16(anon_sym_const),
	1442: uint16(15),
	1443: uint16(1),
	1444: uint16(sym_word),
	1445: uint16(17),
	1446: uint16(1),
	1447: uint16(sym_meta_ident),
	1448: uint16(19),
	1449: uint16(1),
	1450: uint16(sym__ident),
	1451: uint16(139),
	1452: uint16(1),
	1454: uint16(141),
	1455: uint16(1),
	1456: uint16(anon_sym_LF),
	1457: uint16(37),
	1458: uint16(1),
	1459: uint16(sym_line_comment),
	1460: uint16(52),
	1461: uint16(1),
	1462: uint16(aux_sym_program_repeat1),
	1463: uint16(110),
	1464: uint16(1),
	1465: uint16(sym__item),
	1466: uint16(114),
	1467: uint16(4),
	1468: uint16(sym_meta),
	1469: uint16(sym_label),
	1470: uint16(sym_const),
	1471: uint16(sym_instruction),
	1472: uint16(14),
	1473: uint16(3),
	1474: uint16(1),
	1475: uint16(anon_sym_POUND),
	1476: uint16(5),
	1477: uint16(1),
	1478: uint16(aux_sym_line_comment_token2),
	1479: uint16(7),
	1480: uint16(1),
	1481: uint16(sym_block_comment),
	1482: uint16(11),
	1483: uint16(1),
	1484: uint16(anon_sym_label),
	1485: uint16(13),
	1486: uint16(1),
	1487: uint16(anon_sym_const),
	1488: uint16(15),
	1489: uint16(1),
	1490: uint16(sym_word),
	1491: uint16(17),
	1492: uint16(1),
	1493: uint16(sym_meta_ident),
	1494: uint16(19),
	1495: uint16(1),
	1496: uint16(sym__ident),
	1497: uint16(141),
	1498: uint16(1),
	1499: uint16(anon_sym_LF),
	1500: uint16(143),
	1501: uint16(1),
	1503: uint16(38),
	1504: uint16(1),
	1505: uint16(sym_line_comment),
	1506: uint16(52),
	1507: uint16(1),
	1508: uint16(aux_sym_program_repeat1),
	1509: uint16(110),
	1510: uint16(1),
	1511: uint16(sym__item),
	1512: uint16(114),
	1513: uint16(4),
	1514: uint16(sym_meta),
	1515: uint16(sym_label),
	1516: uint16(sym_const),
	1517: uint16(sym_instruction),
	1518: uint16(6),
	1519: uint16(3),
	1520: uint16(1),
	1521: uint16(anon_sym_POUND),
	1522: uint16(5),
	1523: uint16(1),
	1524: uint16(aux_sym_line_comment_token2),
	1525: uint16(7),
	1526: uint16(1),
	1527: uint16(sym_block_comment),
	1528: uint16(63),
	1529: uint16(1),
	1530: uint16(anon_sym_SLASH),
	1531: uint16(39),
	1532: uint16(1),
	1533: uint16(sym_line_comment),
	1534: uint16(61),
	1535: uint16(12),
	1537: uint16(anon_sym_LF),
	1538: uint16(anon_sym_COMMA),
	1539: uint16(anon_sym_LPAREN),
	1540: uint16(anon_sym_DASH),
	1541: uint16(anon_sym_PLUS),
	1542: uint16(anon_sym_RBRACK),
	1543: uint16(anon_sym_STAR),
	1544: uint16(anon_sym_PERCENT),
	1545: uint16(anon_sym_PIPE),
	1546: uint16(anon_sym_CARET),
	1547: uint16(anon_sym_AMP),
	1548: uint16(6),
	1549: uint16(3),
	1550: uint16(1),
	1551: uint16(anon_sym_POUND),
	1552: uint16(5),
	1553: uint16(1),
	1554: uint16(aux_sym_line_comment_token2),
	1555: uint16(7),
	1556: uint16(1),
	1557: uint16(sym_block_comment),
	1558: uint16(86),
	1559: uint16(1),
	1560: uint16(anon_sym_SLASH),
	1561: uint16(40),
	1562: uint16(1),
	1563: uint16(sym_line_comment),
	1564: uint16(84),
	1565: uint16(12),
	1567: uint16(anon_sym_LF),
	1568: uint16(anon_sym_COMMA),
	1569: uint16(anon_sym_RPAREN),
	1570: uint16(anon_sym_DASH),
	1571: uint16(anon_sym_PLUS),
	1572: uint16(anon_sym_RBRACK),
	1573: uint16(anon_sym_STAR),
	1574: uint16(anon_sym_PERCENT),
	1575: uint16(anon_sym_PIPE),
	1576: uint16(anon_sym_CARET),
	1577: uint16(anon_sym_AMP),
	1578: uint16(13),
	1579: uint16(3),
	1580: uint16(1),
	1581: uint16(anon_sym_POUND),
	1582: uint16(5),
	1583: uint16(1),
	1584: uint16(aux_sym_line_comment_token2),
	1585: uint16(7),
	1586: uint16(1),
	1587: uint16(sym_block_comment),
	1588: uint16(11),
	1589: uint16(1),
	1590: uint16(anon_sym_label),
	1591: uint16(13),
	1592: uint16(1),
	1593: uint16(anon_sym_const),
	1594: uint16(15),
	1595: uint16(1),
	1596: uint16(sym_word),
	1597: uint16(17),
	1598: uint16(1),
	1599: uint16(sym_meta_ident),
	1600: uint16(19),
	1601: uint16(1),
	1602: uint16(sym__ident),
	1603: uint16(141),
	1604: uint16(1),
	1605: uint16(anon_sym_LF),
	1606: uint16(41),
	1607: uint16(1),
	1608: uint16(sym_line_comment),
	1609: uint16(52),
	1610: uint16(1),
	1611: uint16(aux_sym_program_repeat1),
	1612: uint16(110),
	1613: uint16(1),
	1614: uint16(sym__item),
	1615: uint16(114),
	1616: uint16(4),
	1617: uint16(sym_meta),
	1618: uint16(sym_label),
	1619: uint16(sym_const),
	1620: uint16(sym_instruction),
	1621: uint16(6),
	1622: uint16(3),
	1623: uint16(1),
	1624: uint16(anon_sym_POUND),
	1625: uint16(5),
	1626: uint16(1),
	1627: uint16(aux_sym_line_comment_token2),
	1628: uint16(7),
	1629: uint16(1),
	1630: uint16(sym_block_comment),
	1631: uint16(90),
	1632: uint16(1),
	1633: uint16(anon_sym_SLASH),
	1634: uint16(42),
	1635: uint16(1),
	1636: uint16(sym_line_comment),
	1637: uint16(88),
	1638: uint16(10),
	1640: uint16(anon_sym_LF),
	1641: uint16(anon_sym_COMMA),
	1642: uint16(anon_sym_DASH),
	1643: uint16(anon_sym_PLUS),
	1644: uint16(anon_sym_STAR),
	1645: uint16(anon_sym_PERCENT),
	1646: uint16(anon_sym_PIPE),
	1647: uint16(anon_sym_CARET),
	1648: uint16(anon_sym_AMP),
	1649: uint16(6),
	1650: uint16(3),
	1651: uint16(1),
	1652: uint16(anon_sym_POUND),
	1653: uint16(5),
	1654: uint16(1),
	1655: uint16(aux_sym_line_comment_token2),
	1656: uint16(7),
	1657: uint16(1),
	1658: uint16(sym_block_comment),
	1659: uint16(72),
	1660: uint16(1),
	1661: uint16(anon_sym_SLASH),
	1662: uint16(43),
	1663: uint16(1),
	1664: uint16(sym_line_comment),
	1665: uint16(74),
	1666: uint16(9),
	1668: uint16(anon_sym_LF),
	1669: uint16(anon_sym_DASH),
	1670: uint16(anon_sym_PLUS),
	1671: uint16(anon_sym_STAR),
	1672: uint16(anon_sym_PERCENT),
	1673: uint16(anon_sym_PIPE),
	1674: uint16(anon_sym_CARET),
	1675: uint16(anon_sym_AMP),
	1676: uint16(7),
	1677: uint16(3),
	1678: uint16(1),
	1679: uint16(anon_sym_POUND),
	1680: uint16(5),
	1681: uint16(1),
	1682: uint16(aux_sym_line_comment_token2),
	1683: uint16(7),
	1684: uint16(1),
	1685: uint16(sym_block_comment),
	1686: uint16(94),
	1687: uint16(1),
	1688: uint16(anon_sym_SLASH),
	1689: uint16(145),
	1690: uint16(1),
	1691: uint16(anon_sym_AMP),
	1692: uint16(44),
	1693: uint16(1),
	1694: uint16(sym_line_comment),
	1695: uint16(92),
	1696: uint16(8),
	1698: uint16(anon_sym_LF),
	1699: uint16(anon_sym_DASH),
	1700: uint16(anon_sym_PLUS),
	1701: uint16(anon_sym_STAR),
	1702: uint16(anon_sym_PERCENT),
	1703: uint16(anon_sym_PIPE),
	1704: uint16(anon_sym_CARET),
	1705: uint16(8),
	1706: uint16(3),
	1707: uint16(1),
	1708: uint16(anon_sym_POUND),
	1709: uint16(5),
	1710: uint16(1),
	1711: uint16(aux_sym_line_comment_token2),
	1712: uint16(7),
	1713: uint16(1),
	1714: uint16(sym_block_comment),
	1715: uint16(94),
	1716: uint16(1),
	1717: uint16(anon_sym_SLASH),
	1718: uint16(145),
	1719: uint16(1),
	1720: uint16(anon_sym_AMP),
	1721: uint16(147),
	1722: uint16(1),
	1723: uint16(anon_sym_CARET),
	1724: uint16(45),
	1725: uint16(1),
	1726: uint16(sym_line_comment),
	1727: uint16(92),
	1728: uint16(7),
	1730: uint16(anon_sym_LF),
	1731: uint16(anon_sym_DASH),
	1732: uint16(anon_sym_PLUS),
	1733: uint16(anon_sym_STAR),
	1734: uint16(anon_sym_PERCENT),
	1735: uint16(anon_sym_PIPE),
	1736: uint16(11),
	1737: uint16(3),
	1738: uint16(1),
	1739: uint16(anon_sym_POUND),
	1740: uint16(5),
	1741: uint16(1),
	1742: uint16(aux_sym_line_comment_token2),
	1743: uint16(7),
	1744: uint16(1),
	1745: uint16(sym_block_comment),
	1746: uint16(145),
	1747: uint16(1),
	1748: uint16(anon_sym_AMP),
	1749: uint16(147),
	1750: uint16(1),
	1751: uint16(anon_sym_CARET),
	1752: uint16(155),
	1753: uint16(1),
	1754: uint16(anon_sym_SLASH),
	1755: uint16(157),
	1756: uint16(1),
	1757: uint16(anon_sym_PIPE),
	1758: uint16(46),
	1759: uint16(1),
	1760: uint16(sym_line_comment),
	1761: uint16(149),
	1762: uint16(2),
	1764: uint16(anon_sym_LF),
	1765: uint16(151),
	1766: uint16(2),
	1767: uint16(anon_sym_DASH),
	1768: uint16(anon_sym_PLUS),
	1769: uint16(153),
	1770: uint16(2),
	1771: uint16(anon_sym_STAR),
	1772: uint16(anon_sym_PERCENT),
	1773: uint16(9),
	1774: uint16(3),
	1775: uint16(1),
	1776: uint16(anon_sym_POUND),
	1777: uint16(5),
	1778: uint16(1),
	1779: uint16(aux_sym_line_comment_token2),
	1780: uint16(7),
	1781: uint16(1),
	1782: uint16(sym_block_comment),
	1783: uint16(94),
	1784: uint16(1),
	1785: uint16(anon_sym_SLASH),
	1786: uint16(145),
	1787: uint16(1),
	1788: uint16(anon_sym_AMP),
	1789: uint16(147),
	1790: uint16(1),
	1791: uint16(anon_sym_CARET),
	1792: uint16(157),
	1793: uint16(1),
	1794: uint16(anon_sym_PIPE),
	1795: uint16(47),
	1796: uint16(1),
	1797: uint16(sym_line_comment),
	1798: uint16(92),
	1799: uint16(6),
	1801: uint16(anon_sym_LF),
	1802: uint16(anon_sym_DASH),
	1803: uint16(anon_sym_PLUS),
	1804: uint16(anon_sym_STAR),
	1805: uint16(anon_sym_PERCENT),
	1806: uint16(10),
	1807: uint16(3),
	1808: uint16(1),
	1809: uint16(anon_sym_POUND),
	1810: uint16(5),
	1811: uint16(1),
	1812: uint16(aux_sym_line_comment_token2),
	1813: uint16(7),
	1814: uint16(1),
	1815: uint16(sym_block_comment),
	1816: uint16(145),
	1817: uint16(1),
	1818: uint16(anon_sym_AMP),
	1819: uint16(147),
	1820: uint16(1),
	1821: uint16(anon_sym_CARET),
	1822: uint16(155),
	1823: uint16(1),
	1824: uint16(anon_sym_SLASH),
	1825: uint16(157),
	1826: uint16(1),
	1827: uint16(anon_sym_PIPE),
	1828: uint16(48),
	1829: uint16(1),
	1830: uint16(sym_line_comment),
	1831: uint16(153),
	1832: uint16(2),
	1833: uint16(anon_sym_STAR),
	1834: uint16(anon_sym_PERCENT),
	1835: uint16(92),
	1836: uint16(4),
	1838: uint16(anon_sym_LF),
	1839: uint16(anon_sym_DASH),
	1840: uint16(anon_sym_PLUS),
	1841: uint16(6),
	1842: uint16(3),
	1843: uint16(1),
	1844: uint16(anon_sym_POUND),
	1845: uint16(5),
	1846: uint16(1),
	1847: uint16(aux_sym_line_comment_token2),
	1848: uint16(7),
	1849: uint16(1),
	1850: uint16(sym_block_comment),
	1851: uint16(94),
	1852: uint16(1),
	1853: uint16(anon_sym_SLASH),
	1854: uint16(49),
	1855: uint16(1),
	1856: uint16(sym_line_comment),
	1857: uint16(92),
	1858: uint16(9),
	1860: uint16(anon_sym_LF),
	1861: uint16(anon_sym_DASH),
	1862: uint16(anon_sym_PLUS),
	1863: uint16(anon_sym_STAR),
	1864: uint16(anon_sym_PERCENT),
	1865: uint16(anon_sym_PIPE),
	1866: uint16(anon_sym_CARET),
	1867: uint16(anon_sym_AMP),
	1868: uint16(9),
	1869: uint16(5),
	1870: uint16(1),
	1871: uint16(aux_sym_line_comment_token2),
	1872: uint16(7),
	1873: uint16(1),
	1874: uint16(sym_block_comment),
	1875: uint16(49),
	1876: uint16(1),
	1877: uint16(anon_sym_POUND),
	1878: uint16(51),
	1879: uint16(1),
	1880: uint16(aux_sym_int_token2),
	1881: uint16(40),
	1882: uint16(1),
	1883: uint16(sym_reg),
	1884: uint16(50),
	1885: uint16(1),
	1886: uint16(sym_line_comment),
	1887: uint16(57),
	1888: uint16(2),
	1889: uint16(sym_meta_ident),
	1890: uint16(sym__ident),
	1891: uint16(122),
	1892: uint16(2),
	1893: uint16(sym_int),
	1894: uint16(sym_ident),
	1895: uint16(55),
	1896: uint16(3),
	1897: uint16(sym_word),
	1898: uint16(sym__reg),
	1899: uint16(sym_address),
	1900: uint16(9),
	1901: uint16(5),
	1902: uint16(1),
	1903: uint16(aux_sym_line_comment_token2),
	1904: uint16(7),
	1905: uint16(1),
	1906: uint16(sym_block_comment),
	1907: uint16(49),
	1908: uint16(1),
	1909: uint16(anon_sym_POUND),
	1910: uint16(51),
	1911: uint16(1),
	1912: uint16(aux_sym_int_token2),
	1913: uint16(40),
	1914: uint16(1),
	1915: uint16(sym_reg),
	1916: uint16(51),
	1917: uint16(1),
	1918: uint16(sym_line_comment),
	1919: uint16(57),
	1920: uint16(2),
	1921: uint16(sym_meta_ident),
	1922: uint16(sym__ident),
	1923: uint16(117),
	1924: uint16(2),
	1925: uint16(sym_int),
	1926: uint16(sym_ident),
	1927: uint16(55),
	1928: uint16(3),
	1929: uint16(sym_word),
	1930: uint16(sym__reg),
	1931: uint16(sym_address),
	1932: uint16(7),
	1933: uint16(3),
	1934: uint16(1),
	1935: uint16(anon_sym_POUND),
	1936: uint16(5),
	1937: uint16(1),
	1938: uint16(aux_sym_line_comment_token2),
	1939: uint16(7),
	1940: uint16(1),
	1941: uint16(sym_block_comment),
	1942: uint16(159),
	1943: uint16(1),
	1945: uint16(161),
	1946: uint16(1),
	1947: uint16(anon_sym_LF),
	1948: uint16(52),
	1949: uint16(2),
	1950: uint16(sym_line_comment),
	1951: uint16(aux_sym_program_repeat1),
	1952: uint16(164),
	1953: uint16(5),
	1954: uint16(anon_sym_label),
	1955: uint16(anon_sym_const),
	1956: uint16(sym_word),
	1957: uint16(sym_meta_ident),
	1958: uint16(sym__ident),
	1959: uint16(6),
	1960: uint16(3),
	1961: uint16(1),
	1962: uint16(anon_sym_POUND),
	1963: uint16(5),
	1964: uint16(1),
	1965: uint16(aux_sym_line_comment_token2),
	1966: uint16(7),
	1967: uint16(1),
	1968: uint16(sym_block_comment),
	1969: uint16(53),
	1970: uint16(1),
	1971: uint16(sym_line_comment),
	1972: uint16(166),
	1973: uint16(2),
	1975: uint16(anon_sym_LF),
	1976: uint16(168),
	1977: uint16(5),
	1978: uint16(anon_sym_label),
	1979: uint16(anon_sym_const),
	1980: uint16(sym_word),
	1981: uint16(sym_meta_ident),
	1982: uint16(sym__ident),
	1983: uint16(9),
	1984: uint16(3),
	1985: uint16(1),
	1986: uint16(anon_sym_POUND),
	1987: uint16(5),
	1988: uint16(1),
	1989: uint16(aux_sym_line_comment_token2),
	1990: uint16(7),
	1991: uint16(1),
	1992: uint16(sym_block_comment),
	1993: uint16(170),
	1994: uint16(1),
	1995: uint16(sym_address),
	1996: uint16(40),
	1997: uint16(1),
	1998: uint16(sym_reg),
	1999: uint16(54),
	2000: uint16(1),
	2001: uint16(sym_line_comment),
	2002: uint16(128),
	2003: uint16(1),
	2004: uint16(sym_ident),
	2005: uint16(55),
	2006: uint16(2),
	2007: uint16(sym_word),
	2008: uint16(sym__reg),
	2009: uint16(57),
	2010: uint16(2),
	2011: uint16(sym_meta_ident),
	2012: uint16(sym__ident),
	2013: uint16(9),
	2014: uint16(3),
	2015: uint16(1),
	2016: uint16(anon_sym_POUND),
	2017: uint16(5),
	2018: uint16(1),
	2019: uint16(aux_sym_line_comment_token2),
	2020: uint16(7),
	2021: uint16(1),
	2022: uint16(sym_block_comment),
	2023: uint16(170),
	2024: uint16(1),
	2025: uint16(sym_address),
	2026: uint16(40),
	2027: uint16(1),
	2028: uint16(sym_reg),
	2029: uint16(55),
	2030: uint16(1),
	2031: uint16(sym_line_comment),
	2032: uint16(129),
	2033: uint16(1),
	2034: uint16(sym_ident),
	2035: uint16(55),
	2036: uint16(2),
	2037: uint16(sym_word),
	2038: uint16(sym__reg),
	2039: uint16(57),
	2040: uint16(2),
	2041: uint16(sym_meta_ident),
	2042: uint16(sym__ident),
	2043: uint16(8),
	2044: uint16(3),
	2045: uint16(1),
	2046: uint16(anon_sym_POUND),
	2047: uint16(5),
	2048: uint16(1),
	2049: uint16(aux_sym_line_comment_token2),
	2050: uint16(7),
	2051: uint16(1),
	2052: uint16(sym_block_comment),
	2053: uint16(55),
	2054: uint16(1),
	2055: uint16(sym__reg),
	2056: uint16(172),
	2057: uint16(1),
	2058: uint16(anon_sym_RBRACE),
	2059: uint16(56),
	2060: uint16(1),
	2061: uint16(sym_line_comment),
	2062: uint16(82),
	2063: uint16(1),
	2064: uint16(sym_reg),
	2065: uint16(170),
	2066: uint16(2),
	2067: uint16(sym_word),
	2068: uint16(sym_address),
	2069: uint16(8),
	2070: uint16(3),
	2071: uint16(1),
	2072: uint16(anon_sym_POUND),
	2073: uint16(5),
	2074: uint16(1),
	2075: uint16(aux_sym_line_comment_token2),
	2076: uint16(7),
	2077: uint16(1),
	2078: uint16(sym_block_comment),
	2079: uint16(55),
	2080: uint16(1),
	2081: uint16(sym__reg),
	2082: uint16(174),
	2083: uint16(1),
	2084: uint16(anon_sym_RBRACE),
	2085: uint16(57),
	2086: uint16(1),
	2087: uint16(sym_line_comment),
	2088: uint16(104),
	2089: uint16(1),
	2090: uint16(sym_reg),
	2091: uint16(170),
	2092: uint16(2),
	2093: uint16(sym_word),
	2094: uint16(sym_address),
	2095: uint16(8),
	2096: uint16(3),
	2097: uint16(1),
	2098: uint16(anon_sym_POUND),
	2099: uint16(5),
	2100: uint16(1),
	2101: uint16(aux_sym_line_comment_token2),
	2102: uint16(7),
	2103: uint16(1),
	2104: uint16(sym_block_comment),
	2105: uint16(178),
	2106: uint16(1),
	2107: uint16(anon_sym_LPAREN),
	2108: uint16(180),
	2109: uint16(1),
	2110: uint16(sym_meta_ident),
	2111: uint16(58),
	2112: uint16(1),
	2113: uint16(sym_line_comment),
	2114: uint16(108),
	2115: uint16(1),
	2116: uint16(sym_meta),
	2117: uint16(176),
	2118: uint16(2),
	2120: uint16(anon_sym_LF),
	2121: uint16(8),
	2122: uint16(3),
	2123: uint16(1),
	2124: uint16(anon_sym_POUND),
	2125: uint16(5),
	2126: uint16(1),
	2127: uint16(aux_sym_line_comment_token2),
	2128: uint16(7),
	2129: uint16(1),
	2130: uint16(sym_block_comment),
	2131: uint16(55),
	2132: uint16(1),
	2133: uint16(sym__reg),
	2134: uint16(182),
	2135: uint16(1),
	2136: uint16(anon_sym_RBRACE),
	2137: uint16(59),
	2138: uint16(1),
	2139: uint16(sym_line_comment),
	2140: uint16(104),
	2141: uint16(1),
	2142: uint16(sym_reg),
	2143: uint16(170),
	2144: uint16(2),
	2145: uint16(sym_word),
	2146: uint16(sym_address),
	2147: uint16(8),
	2148: uint16(3),
	2149: uint16(1),
	2150: uint16(anon_sym_POUND),
	2151: uint16(5),
	2152: uint16(1),
	2153: uint16(aux_sym_line_comment_token2),
	2154: uint16(7),
	2155: uint16(1),
	2156: uint16(sym_block_comment),
	2157: uint16(180),
	2158: uint16(1),
	2159: uint16(sym_meta_ident),
	2160: uint16(186),
	2161: uint16(1),
	2162: uint16(anon_sym_LPAREN),
	2163: uint16(60),
	2164: uint16(1),
	2165: uint16(sym_line_comment),
	2166: uint16(111),
	2167: uint16(1),
	2168: uint16(sym_meta),
	2169: uint16(184),
	2170: uint16(2),
	2172: uint16(anon_sym_LF),
	2173: uint16(7),
	2174: uint16(3),
	2175: uint16(1),
	2176: uint16(anon_sym_POUND),
	2177: uint16(5),
	2178: uint16(1),
	2179: uint16(aux_sym_line_comment_token2),
	2180: uint16(7),
	2181: uint16(1),
	2182: uint16(sym_block_comment),
	2183: uint16(188),
	2184: uint16(1),
	2186: uint16(190),
	2187: uint16(1),
	2188: uint16(anon_sym_LF),
	2189: uint16(41),
	2190: uint16(1),
	2191: uint16(aux_sym_program_repeat1),
	2192: uint16(61),
	2193: uint16(2),
	2194: uint16(sym_line_comment),
	2195: uint16(aux_sym_program_repeat2),
	2196: uint16(8),
	2197: uint16(3),
	2198: uint16(1),
	2199: uint16(anon_sym_POUND),
	2200: uint16(5),
	2201: uint16(1),
	2202: uint16(aux_sym_line_comment_token2),
	2203: uint16(7),
	2204: uint16(1),
	2205: uint16(sym_block_comment),
	2206: uint16(182),
	2207: uint16(1),
	2208: uint16(anon_sym_RBRACE),
	2209: uint16(193),
	2210: uint16(1),
	2211: uint16(anon_sym_COMMA),
	2212: uint16(195),
	2213: uint16(1),
	2214: uint16(anon_sym_DASH),
	2215: uint16(62),
	2216: uint16(1),
	2217: uint16(sym_line_comment),
	2218: uint16(63),
	2219: uint16(1),
	2220: uint16(aux_sym_list_repeat1),
	2221: uint16(6),
	2222: uint16(3),
	2223: uint16(1),
	2224: uint16(anon_sym_POUND),
	2225: uint16(5),
	2226: uint16(1),
	2227: uint16(aux_sym_line_comment_token2),
	2228: uint16(7),
	2229: uint16(1),
	2230: uint16(sym_block_comment),
	2231: uint16(200),
	2232: uint16(1),
	2233: uint16(anon_sym_RBRACE),
	2234: uint16(197),
	2235: uint16(2),
	2236: uint16(anon_sym_COMMA),
	2237: uint16(anon_sym_DASH),
	2238: uint16(63),
	2239: uint16(2),
	2240: uint16(sym_line_comment),
	2241: uint16(aux_sym_list_repeat1),
	2242: uint16(7),
	2243: uint16(3),
	2244: uint16(1),
	2245: uint16(anon_sym_POUND),
	2246: uint16(5),
	2247: uint16(1),
	2248: uint16(aux_sym_line_comment_token2),
	2249: uint16(7),
	2250: uint16(1),
	2251: uint16(sym_block_comment),
	2252: uint16(204),
	2253: uint16(1),
	2254: uint16(anon_sym_COMMA),
	2255: uint16(64),
	2256: uint16(1),
	2257: uint16(sym_line_comment),
	2258: uint16(86),
	2259: uint16(1),
	2260: uint16(aux_sym_meta_repeat2),
	2261: uint16(202),
	2262: uint16(2),
	2264: uint16(anon_sym_LF),
	2265: uint16(7),
	2266: uint16(3),
	2267: uint16(1),
	2268: uint16(anon_sym_POUND),
	2269: uint16(5),
	2270: uint16(1),
	2271: uint16(aux_sym_line_comment_token2),
	2272: uint16(7),
	2273: uint16(1),
	2274: uint16(sym_block_comment),
	2275: uint16(208),
	2276: uint16(1),
	2277: uint16(anon_sym_COMMA),
	2278: uint16(65),
	2279: uint16(1),
	2280: uint16(sym_line_comment),
	2281: uint16(68),
	2282: uint16(1),
	2283: uint16(aux_sym_meta_repeat3),
	2284: uint16(206),
	2285: uint16(2),
	2287: uint16(anon_sym_LF),
	2288: uint16(7),
	2289: uint16(3),
	2290: uint16(1),
	2291: uint16(anon_sym_POUND),
	2292: uint16(5),
	2293: uint16(1),
	2294: uint16(aux_sym_line_comment_token2),
	2295: uint16(7),
	2296: uint16(1),
	2297: uint16(sym_block_comment),
	2298: uint16(210),
	2299: uint16(1),
	2300: uint16(anon_sym_COMMA),
	2301: uint16(66),
	2302: uint16(1),
	2303: uint16(sym_line_comment),
	2304: uint16(87),
	2305: uint16(1),
	2306: uint16(aux_sym_meta_repeat1),
	2307: uint16(202),
	2308: uint16(2),
	2310: uint16(anon_sym_LF),
	2311: uint16(7),
	2312: uint16(3),
	2313: uint16(1),
	2314: uint16(anon_sym_POUND),
	2315: uint16(5),
	2316: uint16(1),
	2317: uint16(aux_sym_line_comment_token2),
	2318: uint16(7),
	2319: uint16(1),
	2320: uint16(sym_block_comment),
	2321: uint16(55),
	2322: uint16(1),
	2323: uint16(sym__reg),
	2324: uint16(67),
	2325: uint16(1),
	2326: uint16(sym_line_comment),
	2327: uint16(121),
	2328: uint16(1),
	2329: uint16(sym_reg),
	2330: uint16(170),
	2331: uint16(2),
	2332: uint16(sym_word),
	2333: uint16(sym_address),
	2334: uint16(7),
	2335: uint16(3),
	2336: uint16(1),
	2337: uint16(anon_sym_POUND),
	2338: uint16(5),
	2339: uint16(1),
	2340: uint16(aux_sym_line_comment_token2),
	2341: uint16(7),
	2342: uint16(1),
	2343: uint16(sym_block_comment),
	2344: uint16(208),
	2345: uint16(1),
	2346: uint16(anon_sym_COMMA),
	2347: uint16(68),
	2348: uint16(1),
	2349: uint16(sym_line_comment),
	2350: uint16(84),
	2351: uint16(1),
	2352: uint16(aux_sym_meta_repeat3),
	2353: uint16(202),
	2354: uint16(2),
	2356: uint16(anon_sym_LF),
	2357: uint16(7),
	2358: uint16(3),
	2359: uint16(1),
	2360: uint16(anon_sym_POUND),
	2361: uint16(5),
	2362: uint16(1),
	2363: uint16(aux_sym_line_comment_token2),
	2364: uint16(7),
	2365: uint16(1),
	2366: uint16(sym_block_comment),
	2367: uint16(212),
	2368: uint16(1),
	2369: uint16(anon_sym_COMMA),
	2370: uint16(69),
	2371: uint16(1),
	2372: uint16(sym_line_comment),
	2373: uint16(71),
	2374: uint16(1),
	2375: uint16(aux_sym_instruction_repeat1),
	2376: uint16(114),
	2377: uint16(2),
	2379: uint16(anon_sym_LF),
	2380: uint16(7),
	2381: uint16(3),
	2382: uint16(1),
	2383: uint16(anon_sym_POUND),
	2384: uint16(5),
	2385: uint16(1),
	2386: uint16(aux_sym_line_comment_token2),
	2387: uint16(7),
	2388: uint16(1),
	2389: uint16(sym_block_comment),
	2390: uint16(204),
	2391: uint16(1),
	2392: uint16(anon_sym_COMMA),
	2393: uint16(64),
	2394: uint16(1),
	2395: uint16(aux_sym_meta_repeat2),
	2396: uint16(70),
	2397: uint16(1),
	2398: uint16(sym_line_comment),
	2399: uint16(206),
	2400: uint16(2),
	2402: uint16(anon_sym_LF),
	2403: uint16(7),
	2404: uint16(3),
	2405: uint16(1),
	2406: uint16(anon_sym_POUND),
	2407: uint16(5),
	2408: uint16(1),
	2409: uint16(aux_sym_line_comment_token2),
	2410: uint16(7),
	2411: uint16(1),
	2412: uint16(sym_block_comment),
	2413: uint16(214),
	2414: uint16(1),
	2415: uint16(anon_sym_COMMA),
	2416: uint16(71),
	2417: uint16(1),
	2418: uint16(sym_line_comment),
	2419: uint16(80),
	2420: uint16(1),
	2421: uint16(aux_sym_instruction_repeat1),
	2422: uint16(47),
	2423: uint16(2),
	2425: uint16(anon_sym_LF),
	2426: uint16(7),
	2427: uint16(3),
	2428: uint16(1),
	2429: uint16(anon_sym_POUND),
	2430: uint16(5),
	2431: uint16(1),
	2432: uint16(aux_sym_line_comment_token2),
	2433: uint16(7),
	2434: uint16(1),
	2435: uint16(sym_block_comment),
	2436: uint16(55),
	2437: uint16(1),
	2438: uint16(sym__reg),
	2439: uint16(72),
	2440: uint16(1),
	2441: uint16(sym_line_comment),
	2442: uint16(120),
	2443: uint16(1),
	2444: uint16(sym_reg),
	2445: uint16(170),
	2446: uint16(2),
	2447: uint16(sym_word),
	2448: uint16(sym_address),
	2449: uint16(7),
	2450: uint16(3),
	2451: uint16(1),
	2452: uint16(anon_sym_POUND),
	2453: uint16(5),
	2454: uint16(1),
	2455: uint16(aux_sym_line_comment_token2),
	2456: uint16(7),
	2457: uint16(1),
	2458: uint16(sym_block_comment),
	2459: uint16(55),
	2460: uint16(1),
	2461: uint16(sym__reg),
	2462: uint16(73),
	2463: uint16(1),
	2464: uint16(sym_line_comment),
	2465: uint16(104),
	2466: uint16(1),
	2467: uint16(sym_reg),
	2468: uint16(170),
	2469: uint16(2),
	2470: uint16(sym_word),
	2471: uint16(sym_address),
	2472: uint16(7),
	2473: uint16(3),
	2474: uint16(1),
	2475: uint16(anon_sym_POUND),
	2476: uint16(5),
	2477: uint16(1),
	2478: uint16(aux_sym_line_comment_token2),
	2479: uint16(7),
	2480: uint16(1),
	2481: uint16(sym_block_comment),
	2482: uint16(210),
	2483: uint16(1),
	2484: uint16(anon_sym_COMMA),
	2485: uint16(66),
	2486: uint16(1),
	2487: uint16(aux_sym_meta_repeat1),
	2488: uint16(74),
	2489: uint16(1),
	2490: uint16(sym_line_comment),
	2491: uint16(206),
	2492: uint16(2),
	2494: uint16(anon_sym_LF),
	2495: uint16(7),
	2496: uint16(3),
	2497: uint16(1),
	2498: uint16(anon_sym_POUND),
	2499: uint16(5),
	2500: uint16(1),
	2501: uint16(aux_sym_line_comment_token2),
	2502: uint16(7),
	2503: uint16(1),
	2504: uint16(sym_block_comment),
	2505: uint16(55),
	2506: uint16(1),
	2507: uint16(sym__reg),
	2508: uint16(75),
	2509: uint16(1),
	2510: uint16(sym_line_comment),
	2511: uint16(102),
	2512: uint16(1),
	2513: uint16(sym_reg),
	2514: uint16(170),
	2515: uint16(2),
	2516: uint16(sym_word),
	2517: uint16(sym_address),
	2518: uint16(7),
	2519: uint16(3),
	2520: uint16(1),
	2521: uint16(anon_sym_POUND),
	2522: uint16(5),
	2523: uint16(1),
	2524: uint16(aux_sym_line_comment_token2),
	2525: uint16(7),
	2526: uint16(1),
	2527: uint16(sym_block_comment),
	2528: uint16(216),
	2529: uint16(1),
	2530: uint16(anon_sym_COMMA),
	2531: uint16(220),
	2532: uint16(1),
	2533: uint16(anon_sym_RBRACK),
	2534: uint16(76),
	2535: uint16(1),
	2536: uint16(sym_line_comment),
	2537: uint16(218),
	2538: uint16(2),
	2539: uint16(anon_sym_DASH),
	2540: uint16(anon_sym_PLUS),
	2541: uint16(6),
	2542: uint16(3),
	2543: uint16(1),
	2544: uint16(anon_sym_POUND),
	2545: uint16(5),
	2546: uint16(1),
	2547: uint16(aux_sym_line_comment_token2),
	2548: uint16(7),
	2549: uint16(1),
	2550: uint16(sym_block_comment),
	2551: uint16(224),
	2552: uint16(1),
	2553: uint16(anon_sym_BANG),
	2554: uint16(77),
	2555: uint16(1),
	2556: uint16(sym_line_comment),
	2557: uint16(222),
	2558: uint16(3),
	2560: uint16(anon_sym_LF),
	2561: uint16(anon_sym_COMMA),
	2562: uint16(6),
	2563: uint16(3),
	2564: uint16(1),
	2565: uint16(anon_sym_POUND),
	2566: uint16(5),
	2567: uint16(1),
	2568: uint16(aux_sym_line_comment_token2),
	2569: uint16(7),
	2570: uint16(1),
	2571: uint16(sym_block_comment),
	2572: uint16(70),
	2573: uint16(1),
	2574: uint16(anon_sym_LPAREN),
	2575: uint16(78),
	2576: uint16(1),
	2577: uint16(sym_line_comment),
	2578: uint16(68),
	2579: uint16(3),
	2581: uint16(anon_sym_LF),
	2582: uint16(anon_sym_COMMA),
	2583: uint16(7),
	2584: uint16(3),
	2585: uint16(1),
	2586: uint16(anon_sym_POUND),
	2587: uint16(5),
	2588: uint16(1),
	2589: uint16(aux_sym_line_comment_token2),
	2590: uint16(7),
	2591: uint16(1),
	2592: uint16(sym_block_comment),
	2593: uint16(55),
	2594: uint16(1),
	2595: uint16(sym__reg),
	2596: uint16(76),
	2597: uint16(1),
	2598: uint16(sym_reg),
	2599: uint16(79),
	2600: uint16(1),
	2601: uint16(sym_line_comment),
	2602: uint16(170),
	2603: uint16(2),
	2604: uint16(sym_word),
	2605: uint16(sym_address),
	2606: uint16(6),
	2607: uint16(3),
	2608: uint16(1),
	2609: uint16(anon_sym_POUND),
	2610: uint16(5),
	2611: uint16(1),
	2612: uint16(aux_sym_line_comment_token2),
	2613: uint16(7),
	2614: uint16(1),
	2615: uint16(sym_block_comment),
	2616: uint16(228),
	2617: uint16(1),
	2618: uint16(anon_sym_COMMA),
	2619: uint16(226),
	2620: uint16(2),
	2622: uint16(anon_sym_LF),
	2623: uint16(80),
	2624: uint16(2),
	2625: uint16(sym_line_comment),
	2626: uint16(aux_sym_instruction_repeat1),
	2627: uint16(8),
	2628: uint16(3),
	2629: uint16(1),
	2630: uint16(anon_sym_POUND),
	2631: uint16(5),
	2632: uint16(1),
	2633: uint16(aux_sym_line_comment_token2),
	2634: uint16(7),
	2635: uint16(1),
	2636: uint16(sym_block_comment),
	2637: uint16(141),
	2638: uint16(1),
	2639: uint16(anon_sym_LF),
	2640: uint16(231),
	2641: uint16(1),
	2643: uint16(37),
	2644: uint16(1),
	2645: uint16(aux_sym_program_repeat1),
	2646: uint16(81),
	2647: uint16(1),
	2648: uint16(sym_line_comment),
	2649: uint16(85),
	2650: uint16(1),
	2651: uint16(aux_sym_program_repeat2),
	2652: uint16(8),
	2653: uint16(3),
	2654: uint16(1),
	2655: uint16(anon_sym_POUND),
	2656: uint16(5),
	2657: uint16(1),
	2658: uint16(aux_sym_line_comment_token2),
	2659: uint16(7),
	2660: uint16(1),
	2661: uint16(sym_block_comment),
	2662: uint16(195),
	2663: uint16(1),
	2664: uint16(anon_sym_DASH),
	2665: uint16(233),
	2666: uint16(1),
	2667: uint16(anon_sym_COMMA),
	2668: uint16(235),
	2669: uint16(1),
	2670: uint16(anon_sym_RBRACE),
	2671: uint16(62),
	2672: uint16(1),
	2673: uint16(aux_sym_list_repeat1),
	2674: uint16(82),
	2675: uint16(1),
	2676: uint16(sym_line_comment),
	2677: uint16(6),
	2678: uint16(3),
	2679: uint16(1),
	2680: uint16(anon_sym_POUND),
	2681: uint16(5),
	2682: uint16(1),
	2683: uint16(aux_sym_line_comment_token2),
	2684: uint16(7),
	2685: uint16(1),
	2686: uint16(sym_block_comment),
	2687: uint16(239),
	2688: uint16(1),
	2689: uint16(anon_sym_BANG),
	2690: uint16(83),
	2691: uint16(1),
	2692: uint16(sym_line_comment),
	2693: uint16(237),
	2694: uint16(3),
	2696: uint16(anon_sym_LF),
	2697: uint16(anon_sym_COMMA),
	2698: uint16(6),
	2699: uint16(3),
	2700: uint16(1),
	2701: uint16(anon_sym_POUND),
	2702: uint16(5),
	2703: uint16(1),
	2704: uint16(aux_sym_line_comment_token2),
	2705: uint16(7),
	2706: uint16(1),
	2707: uint16(sym_block_comment),
	2708: uint16(243),
	2709: uint16(1),
	2710: uint16(anon_sym_COMMA),
	2711: uint16(241),
	2712: uint16(2),
	2714: uint16(anon_sym_LF),
	2715: uint16(84),
	2716: uint16(2),
	2717: uint16(sym_line_comment),
	2718: uint16(aux_sym_meta_repeat3),
	2719: uint16(8),
	2720: uint16(3),
	2721: uint16(1),
	2722: uint16(anon_sym_POUND),
	2723: uint16(5),
	2724: uint16(1),
	2725: uint16(aux_sym_line_comment_token2),
	2726: uint16(7),
	2727: uint16(1),
	2728: uint16(sym_block_comment),
	2729: uint16(139),
	2730: uint16(1),
	2732: uint16(141),
	2733: uint16(1),
	2734: uint16(anon_sym_LF),
	2735: uint16(38),
	2736: uint16(1),
	2737: uint16(aux_sym_program_repeat1),
	2738: uint16(61),
	2739: uint16(1),
	2740: uint16(aux_sym_program_repeat2),
	2741: uint16(85),
	2742: uint16(1),
	2743: uint16(sym_line_comment),
	2744: uint16(6),
	2745: uint16(3),
	2746: uint16(1),
	2747: uint16(anon_sym_POUND),
	2748: uint16(5),
	2749: uint16(1),
	2750: uint16(aux_sym_line_comment_token2),
	2751: uint16(7),
	2752: uint16(1),
	2753: uint16(sym_block_comment),
	2754: uint16(248),
	2755: uint16(1),
	2756: uint16(anon_sym_COMMA),
	2757: uint16(246),
	2758: uint16(2),
	2760: uint16(anon_sym_LF),
	2761: uint16(86),
	2762: uint16(2),
	2763: uint16(sym_line_comment),
	2764: uint16(aux_sym_meta_repeat2),
	2765: uint16(6),
	2766: uint16(3),
	2767: uint16(1),
	2768: uint16(anon_sym_POUND),
	2769: uint16(5),
	2770: uint16(1),
	2771: uint16(aux_sym_line_comment_token2),
	2772: uint16(7),
	2773: uint16(1),
	2774: uint16(sym_block_comment),
	2775: uint16(253),
	2776: uint16(1),
	2777: uint16(anon_sym_COMMA),
	2778: uint16(251),
	2779: uint16(2),
	2781: uint16(anon_sym_LF),
	2782: uint16(87),
	2783: uint16(2),
	2784: uint16(sym_line_comment),
	2785: uint16(aux_sym_meta_repeat1),
	2786: uint16(6),
	2787: uint16(3),
	2788: uint16(1),
	2789: uint16(anon_sym_POUND),
	2790: uint16(5),
	2791: uint16(1),
	2792: uint16(aux_sym_line_comment_token2),
	2793: uint16(7),
	2794: uint16(1),
	2795: uint16(sym_block_comment),
	2796: uint16(88),
	2797: uint16(1),
	2798: uint16(sym_line_comment),
	2799: uint16(99),
	2800: uint16(1),
	2801: uint16(sym_string),
	2802: uint16(53),
	2803: uint16(2),
	2804: uint16(aux_sym_string_token1),
	2805: uint16(aux_sym_string_token2),
	2806: uint16(5),
	2807: uint16(3),
	2808: uint16(1),
	2809: uint16(anon_sym_POUND),
	2810: uint16(5),
	2811: uint16(1),
	2812: uint16(aux_sym_line_comment_token2),
	2813: uint16(7),
	2814: uint16(1),
	2815: uint16(sym_block_comment),
	2816: uint16(89),
	2817: uint16(1),
	2818: uint16(sym_line_comment),
	2819: uint16(256),
	2820: uint16(3),
	2822: uint16(anon_sym_LF),
	2823: uint16(anon_sym_COMMA),
	2824: uint16(5),
	2825: uint16(3),
	2826: uint16(1),
	2827: uint16(anon_sym_POUND),
	2828: uint16(5),
	2829: uint16(1),
	2830: uint16(aux_sym_line_comment_token2),
	2831: uint16(7),
	2832: uint16(1),
	2833: uint16(sym_block_comment),
	2834: uint16(90),
	2835: uint16(1),
	2836: uint16(sym_line_comment),
	2837: uint16(258),
	2838: uint16(3),
	2840: uint16(anon_sym_LF),
	2841: uint16(anon_sym_COMMA),
	2842: uint16(5),
	2843: uint16(3),
	2844: uint16(1),
	2845: uint16(anon_sym_POUND),
	2846: uint16(5),
	2847: uint16(1),
	2848: uint16(aux_sym_line_comment_token2),
	2849: uint16(7),
	2850: uint16(1),
	2851: uint16(sym_block_comment),
	2852: uint16(91),
	2853: uint16(1),
	2854: uint16(sym_line_comment),
	2855: uint16(226),
	2856: uint16(3),
	2858: uint16(anon_sym_LF),
	2859: uint16(anon_sym_COMMA),
	2860: uint16(5),
	2861: uint16(3),
	2862: uint16(1),
	2863: uint16(anon_sym_POUND),
	2864: uint16(5),
	2865: uint16(1),
	2866: uint16(aux_sym_line_comment_token2),
	2867: uint16(7),
	2868: uint16(1),
	2869: uint16(sym_block_comment),
	2870: uint16(92),
	2871: uint16(1),
	2872: uint16(sym_line_comment),
	2873: uint16(222),
	2874: uint16(3),
	2876: uint16(anon_sym_LF),
	2877: uint16(anon_sym_COMMA),
	2878: uint16(5),
	2879: uint16(3),
	2880: uint16(1),
	2881: uint16(anon_sym_POUND),
	2882: uint16(5),
	2883: uint16(1),
	2884: uint16(aux_sym_line_comment_token2),
	2885: uint16(7),
	2886: uint16(1),
	2887: uint16(sym_block_comment),
	2888: uint16(93),
	2889: uint16(1),
	2890: uint16(sym_line_comment),
	2891: uint16(237),
	2892: uint16(3),
	2894: uint16(anon_sym_LF),
	2895: uint16(anon_sym_COMMA),
	2896: uint16(5),
	2897: uint16(3),
	2898: uint16(1),
	2899: uint16(anon_sym_POUND),
	2900: uint16(5),
	2901: uint16(1),
	2902: uint16(aux_sym_line_comment_token2),
	2903: uint16(7),
	2904: uint16(1),
	2905: uint16(sym_block_comment),
	2906: uint16(94),
	2907: uint16(1),
	2908: uint16(sym_line_comment),
	2909: uint16(260),
	2910: uint16(3),
	2912: uint16(anon_sym_LF),
	2913: uint16(anon_sym_COMMA),
	2914: uint16(5),
	2915: uint16(3),
	2916: uint16(1),
	2917: uint16(anon_sym_POUND),
	2918: uint16(5),
	2919: uint16(1),
	2920: uint16(aux_sym_line_comment_token2),
	2921: uint16(7),
	2922: uint16(1),
	2923: uint16(sym_block_comment),
	2924: uint16(95),
	2925: uint16(1),
	2926: uint16(sym_line_comment),
	2927: uint16(246),
	2928: uint16(3),
	2930: uint16(anon_sym_LF),
	2931: uint16(anon_sym_COMMA),
	2932: uint16(5),
	2933: uint16(3),
	2934: uint16(1),
	2935: uint16(anon_sym_POUND),
	2936: uint16(5),
	2937: uint16(1),
	2938: uint16(aux_sym_line_comment_token2),
	2939: uint16(7),
	2940: uint16(1),
	2941: uint16(sym_block_comment),
	2942: uint16(96),
	2943: uint16(1),
	2944: uint16(sym_line_comment),
	2945: uint16(262),
	2946: uint16(3),
	2948: uint16(anon_sym_LF),
	2949: uint16(anon_sym_COMMA),
	2950: uint16(5),
	2951: uint16(3),
	2952: uint16(1),
	2953: uint16(anon_sym_POUND),
	2954: uint16(5),
	2955: uint16(1),
	2956: uint16(aux_sym_line_comment_token2),
	2957: uint16(7),
	2958: uint16(1),
	2959: uint16(sym_block_comment),
	2960: uint16(97),
	2961: uint16(1),
	2962: uint16(sym_line_comment),
	2963: uint16(251),
	2964: uint16(3),
	2966: uint16(anon_sym_LF),
	2967: uint16(anon_sym_COMMA),
	2968: uint16(5),
	2969: uint16(3),
	2970: uint16(1),
	2971: uint16(anon_sym_POUND),
	2972: uint16(5),
	2973: uint16(1),
	2974: uint16(aux_sym_line_comment_token2),
	2975: uint16(7),
	2976: uint16(1),
	2977: uint16(sym_block_comment),
	2978: uint16(98),
	2979: uint16(1),
	2980: uint16(sym_line_comment),
	2981: uint16(264),
	2982: uint16(3),
	2984: uint16(anon_sym_LF),
	2985: uint16(anon_sym_COMMA),
	2986: uint16(5),
	2987: uint16(3),
	2988: uint16(1),
	2989: uint16(anon_sym_POUND),
	2990: uint16(5),
	2991: uint16(1),
	2992: uint16(aux_sym_line_comment_token2),
	2993: uint16(7),
	2994: uint16(1),
	2995: uint16(sym_block_comment),
	2996: uint16(99),
	2997: uint16(1),
	2998: uint16(sym_line_comment),
	2999: uint16(241),
	3000: uint16(3),
	3002: uint16(anon_sym_LF),
	3003: uint16(anon_sym_COMMA),
	3004: uint16(5),
	3005: uint16(3),
	3006: uint16(1),
	3007: uint16(anon_sym_POUND),
	3008: uint16(5),
	3009: uint16(1),
	3010: uint16(aux_sym_line_comment_token2),
	3011: uint16(7),
	3012: uint16(1),
	3013: uint16(sym_block_comment),
	3014: uint16(100),
	3015: uint16(1),
	3016: uint16(sym_line_comment),
	3017: uint16(68),
	3018: uint16(3),
	3020: uint16(anon_sym_LF),
	3021: uint16(anon_sym_COMMA),
	3022: uint16(5),
	3023: uint16(3),
	3024: uint16(1),
	3025: uint16(anon_sym_POUND),
	3026: uint16(5),
	3027: uint16(1),
	3028: uint16(aux_sym_line_comment_token2),
	3029: uint16(7),
	3030: uint16(1),
	3031: uint16(sym_block_comment),
	3032: uint16(101),
	3033: uint16(1),
	3034: uint16(sym_line_comment),
	3035: uint16(266),
	3036: uint16(3),
	3038: uint16(anon_sym_LF),
	3039: uint16(anon_sym_COMMA),
	3040: uint16(6),
	3041: uint16(3),
	3042: uint16(1),
	3043: uint16(anon_sym_POUND),
	3044: uint16(5),
	3045: uint16(1),
	3046: uint16(aux_sym_line_comment_token2),
	3047: uint16(7),
	3048: uint16(1),
	3049: uint16(sym_block_comment),
	3050: uint16(270),
	3051: uint16(1),
	3052: uint16(anon_sym_RBRACK),
	3053: uint16(102),
	3054: uint16(1),
	3055: uint16(sym_line_comment),
	3056: uint16(268),
	3057: uint16(2),
	3058: uint16(anon_sym_DASH),
	3059: uint16(anon_sym_PLUS),
	3060: uint16(5),
	3061: uint16(3),
	3062: uint16(1),
	3063: uint16(anon_sym_POUND),
	3064: uint16(5),
	3065: uint16(1),
	3066: uint16(aux_sym_line_comment_token2),
	3067: uint16(7),
	3068: uint16(1),
	3069: uint16(sym_block_comment),
	3070: uint16(103),
	3071: uint16(1),
	3072: uint16(sym_line_comment),
	3073: uint16(272),
	3074: uint16(3),
	3076: uint16(anon_sym_LF),
	3077: uint16(anon_sym_COMMA),
	3078: uint16(5),
	3079: uint16(3),
	3080: uint16(1),
	3081: uint16(anon_sym_POUND),
	3082: uint16(5),
	3083: uint16(1),
	3084: uint16(aux_sym_line_comment_token2),
	3085: uint16(7),
	3086: uint16(1),
	3087: uint16(sym_block_comment),
	3088: uint16(104),
	3089: uint16(1),
	3090: uint16(sym_line_comment),
	3091: uint16(200),
	3092: uint16(3),
	3093: uint16(anon_sym_COMMA),
	3094: uint16(anon_sym_DASH),
	3095: uint16(anon_sym_RBRACE),
	3096: uint16(6),
	3097: uint16(5),
	3098: uint16(1),
	3099: uint16(aux_sym_line_comment_token2),
	3100: uint16(7),
	3101: uint16(1),
	3102: uint16(sym_block_comment),
	3103: uint16(49),
	3104: uint16(1),
	3105: uint16(anon_sym_POUND),
	3106: uint16(274),
	3107: uint16(1),
	3108: uint16(aux_sym_int_token2),
	3109: uint16(97),
	3110: uint16(1),
	3111: uint16(sym_int),
	3112: uint16(105),
	3113: uint16(1),
	3114: uint16(sym_line_comment),
	3115: uint16(5),
	3116: uint16(3),
	3117: uint16(1),
	3118: uint16(anon_sym_POUND),
	3119: uint16(5),
	3120: uint16(1),
	3121: uint16(aux_sym_line_comment_token2),
	3122: uint16(7),
	3123: uint16(1),
	3124: uint16(sym_block_comment),
	3125: uint16(106),
	3126: uint16(1),
	3127: uint16(sym_line_comment),
	3128: uint16(206),
	3129: uint16(2),
	3131: uint16(anon_sym_LF),
	3132: uint16(5),
	3133: uint16(3),
	3134: uint16(1),
	3135: uint16(anon_sym_POUND),
	3136: uint16(5),
	3137: uint16(1),
	3138: uint16(aux_sym_line_comment_token2),
	3139: uint16(7),
	3140: uint16(1),
	3141: uint16(sym_block_comment),
	3142: uint16(107),
	3143: uint16(1),
	3144: uint16(sym_line_comment),
	3145: uint16(276),
	3146: uint16(2),
	3148: uint16(anon_sym_LF),
	3149: uint16(5),
	3150: uint16(3),
	3151: uint16(1),
	3152: uint16(anon_sym_POUND),
	3153: uint16(5),
	3154: uint16(1),
	3155: uint16(aux_sym_line_comment_token2),
	3156: uint16(7),
	3157: uint16(1),
	3158: uint16(sym_block_comment),
	3159: uint16(108),
	3160: uint16(1),
	3161: uint16(sym_line_comment),
	3162: uint16(278),
	3163: uint16(2),
	3165: uint16(anon_sym_LF),
	3166: uint16(6),
	3167: uint16(5),
	3168: uint16(1),
	3169: uint16(aux_sym_line_comment_token2),
	3170: uint16(7),
	3171: uint16(1),
	3172: uint16(sym_block_comment),
	3173: uint16(49),
	3174: uint16(1),
	3175: uint16(anon_sym_POUND),
	3176: uint16(274),
	3177: uint16(1),
	3178: uint16(aux_sym_int_token2),
	3179: uint16(109),
	3180: uint16(1),
	3181: uint16(sym_line_comment),
	3182: uint16(123),
	3183: uint16(1),
	3184: uint16(sym_int),
	3185: uint16(5),
	3186: uint16(3),
	3187: uint16(1),
	3188: uint16(anon_sym_POUND),
	3189: uint16(5),
	3190: uint16(1),
	3191: uint16(aux_sym_line_comment_token2),
	3192: uint16(7),
	3193: uint16(1),
	3194: uint16(sym_block_comment),
	3195: uint16(110),
	3196: uint16(1),
	3197: uint16(sym_line_comment),
	3198: uint16(188),
	3199: uint16(2),
	3201: uint16(anon_sym_LF),
	3202: uint16(5),
	3203: uint16(3),
	3204: uint16(1),
	3205: uint16(anon_sym_POUND),
	3206: uint16(5),
	3207: uint16(1),
	3208: uint16(aux_sym_line_comment_token2),
	3209: uint16(7),
	3210: uint16(1),
	3211: uint16(sym_block_comment),
	3212: uint16(111),
	3213: uint16(1),
	3214: uint16(sym_line_comment),
	3215: uint16(280),
	3216: uint16(2),
	3218: uint16(anon_sym_LF),
	3219: uint16(5),
	3220: uint16(3),
	3221: uint16(1),
	3222: uint16(anon_sym_POUND),
	3223: uint16(5),
	3224: uint16(1),
	3225: uint16(aux_sym_line_comment_token2),
	3226: uint16(7),
	3227: uint16(1),
	3228: uint16(sym_block_comment),
	3229: uint16(112),
	3230: uint16(1),
	3231: uint16(sym_line_comment),
	3232: uint16(282),
	3233: uint16(2),
	3235: uint16(anon_sym_LF),
	3236: uint16(5),
	3237: uint16(3),
	3238: uint16(1),
	3239: uint16(anon_sym_POUND),
	3240: uint16(5),
	3241: uint16(1),
	3242: uint16(aux_sym_line_comment_token2),
	3243: uint16(7),
	3244: uint16(1),
	3245: uint16(sym_block_comment),
	3246: uint16(113),
	3247: uint16(1),
	3248: uint16(sym_line_comment),
	3249: uint16(284),
	3250: uint16(2),
	3252: uint16(anon_sym_LF),
	3253: uint16(5),
	3254: uint16(3),
	3255: uint16(1),
	3256: uint16(anon_sym_POUND),
	3257: uint16(5),
	3258: uint16(1),
	3259: uint16(aux_sym_line_comment_token2),
	3260: uint16(7),
	3261: uint16(1),
	3262: uint16(sym_block_comment),
	3263: uint16(114),
	3264: uint16(1),
	3265: uint16(sym_line_comment),
	3266: uint16(286),
	3267: uint16(2),
	3269: uint16(anon_sym_LF),
	3270: uint16(6),
	3271: uint16(5),
	3272: uint16(1),
	3273: uint16(aux_sym_line_comment_token2),
	3274: uint16(7),
	3275: uint16(1),
	3276: uint16(sym_block_comment),
	3277: uint16(49),
	3278: uint16(1),
	3279: uint16(anon_sym_POUND),
	3280: uint16(274),
	3281: uint16(1),
	3282: uint16(aux_sym_int_token2),
	3283: uint16(115),
	3284: uint16(1),
	3285: uint16(sym_line_comment),
	3286: uint16(122),
	3287: uint16(1),
	3288: uint16(sym_int),
	3289: uint16(5),
	3290: uint16(288),
	3291: uint16(1),
	3292: uint16(anon_sym_POUND),
	3293: uint16(290),
	3294: uint16(1),
	3295: uint16(aux_sym_line_comment_token1),
	3296: uint16(292),
	3297: uint16(1),
	3298: uint16(aux_sym_line_comment_token2),
	3299: uint16(294),
	3300: uint16(1),
	3301: uint16(sym_block_comment),
	3302: uint16(116),
	3303: uint16(1),
	3304: uint16(sym_line_comment),
	3305: uint16(5),
	3306: uint16(3),
	3307: uint16(1),
	3308: uint16(anon_sym_POUND),
	3309: uint16(5),
	3310: uint16(1),
	3311: uint16(aux_sym_line_comment_token2),
	3312: uint16(7),
	3313: uint16(1),
	3314: uint16(sym_block_comment),
	3315: uint16(296),
	3316: uint16(1),
	3317: uint16(anon_sym_RBRACK),
	3318: uint16(117),
	3319: uint16(1),
	3320: uint16(sym_line_comment),
	3321: uint16(5),
	3322: uint16(3),
	3323: uint16(1),
	3324: uint16(anon_sym_POUND),
	3325: uint16(5),
	3326: uint16(1),
	3327: uint16(aux_sym_line_comment_token2),
	3328: uint16(7),
	3329: uint16(1),
	3330: uint16(sym_block_comment),
	3331: uint16(298),
	3332: uint16(1),
	3333: uint16(anon_sym_LBRACK),
	3334: uint16(118),
	3335: uint16(1),
	3336: uint16(sym_line_comment),
	3337: uint16(5),
	3338: uint16(3),
	3339: uint16(1),
	3340: uint16(anon_sym_POUND),
	3341: uint16(5),
	3342: uint16(1),
	3343: uint16(aux_sym_line_comment_token2),
	3344: uint16(7),
	3345: uint16(1),
	3346: uint16(sym_block_comment),
	3347: uint16(300),
	3348: uint16(1),
	3349: uint16(anon_sym_LBRACK),
	3350: uint16(119),
	3351: uint16(1),
	3352: uint16(sym_line_comment),
	3353: uint16(5),
	3354: uint16(3),
	3355: uint16(1),
	3356: uint16(anon_sym_POUND),
	3357: uint16(5),
	3358: uint16(1),
	3359: uint16(aux_sym_line_comment_token2),
	3360: uint16(7),
	3361: uint16(1),
	3362: uint16(sym_block_comment),
	3363: uint16(224),
	3364: uint16(1),
	3365: uint16(anon_sym_RPAREN),
	3366: uint16(120),
	3367: uint16(1),
	3368: uint16(sym_line_comment),
	3369: uint16(5),
	3370: uint16(3),
	3371: uint16(1),
	3372: uint16(anon_sym_POUND),
	3373: uint16(5),
	3374: uint16(1),
	3375: uint16(aux_sym_line_comment_token2),
	3376: uint16(7),
	3377: uint16(1),
	3378: uint16(sym_block_comment),
	3379: uint16(302),
	3380: uint16(1),
	3381: uint16(anon_sym_RPAREN),
	3382: uint16(121),
	3383: uint16(1),
	3384: uint16(sym_line_comment),
	3385: uint16(5),
	3386: uint16(3),
	3387: uint16(1),
	3388: uint16(anon_sym_POUND),
	3389: uint16(5),
	3390: uint16(1),
	3391: uint16(aux_sym_line_comment_token2),
	3392: uint16(7),
	3393: uint16(1),
	3394: uint16(sym_block_comment),
	3395: uint16(270),
	3396: uint16(1),
	3397: uint16(anon_sym_RBRACK),
	3398: uint16(122),
	3399: uint16(1),
	3400: uint16(sym_line_comment),
	3401: uint16(5),
	3402: uint16(3),
	3403: uint16(1),
	3404: uint16(anon_sym_POUND),
	3405: uint16(5),
	3406: uint16(1),
	3407: uint16(aux_sym_line_comment_token2),
	3408: uint16(7),
	3409: uint16(1),
	3410: uint16(sym_block_comment),
	3411: uint16(304),
	3412: uint16(1),
	3413: uint16(anon_sym_RBRACK),
	3414: uint16(123),
	3415: uint16(1),
	3416: uint16(sym_line_comment),
	3417: uint16(5),
	3418: uint16(3),
	3419: uint16(1),
	3420: uint16(anon_sym_POUND),
	3421: uint16(5),
	3422: uint16(1),
	3423: uint16(aux_sym_line_comment_token2),
	3424: uint16(7),
	3425: uint16(1),
	3426: uint16(sym_block_comment),
	3427: uint16(306),
	3428: uint16(1),
	3429: uint16(sym_float),
	3430: uint16(124),
	3431: uint16(1),
	3432: uint16(sym_line_comment),
	3433: uint16(5),
	3434: uint16(3),
	3435: uint16(1),
	3436: uint16(anon_sym_POUND),
	3437: uint16(5),
	3438: uint16(1),
	3439: uint16(aux_sym_line_comment_token2),
	3440: uint16(7),
	3441: uint16(1),
	3442: uint16(sym_block_comment),
	3443: uint16(308),
	3444: uint16(1),
	3445: uint16(aux_sym_int_token1),
	3446: uint16(125),
	3447: uint16(1),
	3448: uint16(sym_line_comment),
	3449: uint16(5),
	3450: uint16(3),
	3451: uint16(1),
	3452: uint16(anon_sym_POUND),
	3453: uint16(5),
	3454: uint16(1),
	3455: uint16(aux_sym_line_comment_token2),
	3456: uint16(7),
	3457: uint16(1),
	3458: uint16(sym_block_comment),
	3459: uint16(310),
	3460: uint16(1),
	3461: uint16(anon_sym_rel),
	3462: uint16(126),
	3463: uint16(1),
	3464: uint16(sym_line_comment),
	3465: uint16(5),
	3466: uint16(3),
	3467: uint16(1),
	3468: uint16(anon_sym_POUND),
	3469: uint16(5),
	3470: uint16(1),
	3471: uint16(aux_sym_line_comment_token2),
	3472: uint16(7),
	3473: uint16(1),
	3474: uint16(sym_block_comment),
	3475: uint16(312),
	3476: uint16(1),
	3477: uint16(anon_sym_ptr),
	3478: uint16(127),
	3479: uint16(1),
	3480: uint16(sym_line_comment),
	3481: uint16(5),
	3482: uint16(3),
	3483: uint16(1),
	3484: uint16(anon_sym_POUND),
	3485: uint16(5),
	3486: uint16(1),
	3487: uint16(aux_sym_line_comment_token2),
	3488: uint16(7),
	3489: uint16(1),
	3490: uint16(sym_block_comment),
	3491: uint16(314),
	3492: uint16(1),
	3493: uint16(anon_sym_RPAREN),
	3494: uint16(128),
	3495: uint16(1),
	3496: uint16(sym_line_comment),
	3497: uint16(5),
	3498: uint16(3),
	3499: uint16(1),
	3500: uint16(anon_sym_POUND),
	3501: uint16(5),
	3502: uint16(1),
	3503: uint16(aux_sym_line_comment_token2),
	3504: uint16(7),
	3505: uint16(1),
	3506: uint16(sym_block_comment),
	3507: uint16(316),
	3508: uint16(1),
	3509: uint16(anon_sym_RPAREN),
	3510: uint16(129),
	3511: uint16(1),
	3512: uint16(sym_line_comment),
	3513: uint16(5),
	3514: uint16(3),
	3515: uint16(1),
	3516: uint16(anon_sym_POUND),
	3517: uint16(5),
	3518: uint16(1),
	3519: uint16(aux_sym_line_comment_token2),
	3520: uint16(7),
	3521: uint16(1),
	3522: uint16(sym_block_comment),
	3523: uint16(318),
	3524: uint16(1),
	3526: uint16(130),
	3527: uint16(1),
	3528: uint16(sym_line_comment),
	3529: uint16(5),
	3530: uint16(3),
	3531: uint16(1),
	3532: uint16(anon_sym_POUND),
	3533: uint16(5),
	3534: uint16(1),
	3535: uint16(aux_sym_line_comment_token2),
	3536: uint16(7),
	3537: uint16(1),
	3538: uint16(sym_block_comment),
	3539: uint16(23),
	3540: uint16(1),
	3541: uint16(anon_sym_COLON),
	3542: uint16(131),
	3543: uint16(1),
	3544: uint16(sym_line_comment),
	3545: uint16(5),
	3546: uint16(3),
	3547: uint16(1),
	3548: uint16(anon_sym_POUND),
	3549: uint16(5),
	3550: uint16(1),
	3551: uint16(aux_sym_line_comment_token2),
	3552: uint16(7),
	3553: uint16(1),
	3554: uint16(sym_block_comment),
	3555: uint16(320),
	3556: uint16(1),
	3557: uint16(sym_word),
	3558: uint16(132),
	3559: uint16(1),
	3560: uint16(sym_line_comment),
	3561: uint16(5),
	3562: uint16(3),
	3563: uint16(1),
	3564: uint16(anon_sym_POUND),
	3565: uint16(5),
	3566: uint16(1),
	3567: uint16(aux_sym_line_comment_token2),
	3568: uint16(7),
	3569: uint16(1),
	3570: uint16(sym_block_comment),
	3571: uint16(322),
	3572: uint16(1),
	3573: uint16(sym_word),
	3574: uint16(133),
	3575: uint16(1),
	3576: uint16(sym_line_comment),
	3577: uint16(5),
	3578: uint16(3),
	3579: uint16(1),
	3580: uint16(anon_sym_POUND),
	3581: uint16(5),
	3582: uint16(1),
	3583: uint16(aux_sym_line_comment_token2),
	3584: uint16(7),
	3585: uint16(1),
	3586: uint16(sym_block_comment),
	3587: uint16(324),
	3588: uint16(1),
	3589: uint16(aux_sym_int_token1),
	3590: uint16(134),
	3591: uint16(1),
	3592: uint16(sym_line_comment),
	3593: uint16(1),
	3594: uint16(326),
	3595: uint16(1),
	3597: uint16(1),
	3598: uint16(328),
	3599: uint16(1),
}

var ts_small_parse_table_map = [134]uint32_t{
	1:   uint32(69),
	2:   uint32(138),
	3:   uint32(203),
	4:   uint32(238),
	5:   uint32(279),
	6:   uint32(314),
	7:   uint32(348),
	8:   uint32(386),
	9:   uint32(420),
	10:  uint32(454),
	11:  uint32(493),
	12:  uint32(536),
	13:  uint32(569),
	14:  uint32(606),
	15:  uint32(639),
	16:  uint32(686),
	17:  uint32(721),
	18:  uint32(769),
	19:  uint32(815),
	20:  uint32(866),
	21:  uint32(914),
	22:  uint32(955),
	23:  uint32(996),
	24:  uint32(1037),
	25:  uint32(1078),
	26:  uint32(1119),
	27:  uint32(1160),
	28:  uint32(1201),
	29:  uint32(1242),
	30:  uint32(1283),
	31:  uint32(1324),
	32:  uint32(1365),
	33:  uint32(1396),
	34:  uint32(1426),
	35:  uint32(1472),
	36:  uint32(1518),
	37:  uint32(1548),
	38:  uint32(1578),
	39:  uint32(1621),
	40:  uint32(1649),
	41:  uint32(1676),
	42:  uint32(1705),
	43:  uint32(1736),
	44:  uint32(1773),
	45:  uint32(1806),
	46:  uint32(1841),
	47:  uint32(1868),
	48:  uint32(1900),
	49:  uint32(1932),
	50:  uint32(1959),
	51:  uint32(1983),
	52:  uint32(2013),
	53:  uint32(2043),
	54:  uint32(2069),
	55:  uint32(2095),
	56:  uint32(2121),
	57:  uint32(2147),
	58:  uint32(2173),
	59:  uint32(2196),
	60:  uint32(2221),
	61:  uint32(2242),
	62:  uint32(2265),
	63:  uint32(2288),
	64:  uint32(2311),
	65:  uint32(2334),
	66:  uint32(2357),
	67:  uint32(2380),
	68:  uint32(2403),
	69:  uint32(2426),
	70:  uint32(2449),
	71:  uint32(2472),
	72:  uint32(2495),
	73:  uint32(2518),
	74:  uint32(2541),
	75:  uint32(2562),
	76:  uint32(2583),
	77:  uint32(2606),
	78:  uint32(2627),
	79:  uint32(2652),
	80:  uint32(2677),
	81:  uint32(2698),
	82:  uint32(2719),
	83:  uint32(2744),
	84:  uint32(2765),
	85:  uint32(2786),
	86:  uint32(2806),
	87:  uint32(2824),
	88:  uint32(2842),
	89:  uint32(2860),
	90:  uint32(2878),
	91:  uint32(2896),
	92:  uint32(2914),
	93:  uint32(2932),
	94:  uint32(2950),
	95:  uint32(2968),
	96:  uint32(2986),
	97:  uint32(3004),
	98:  uint32(3022),
	99:  uint32(3040),
	100: uint32(3060),
	101: uint32(3078),
	102: uint32(3096),
	103: uint32(3115),
	104: uint32(3132),
	105: uint32(3149),
	106: uint32(3166),
	107: uint32(3185),
	108: uint32(3202),
	109: uint32(3219),
	110: uint32(3236),
	111: uint32(3253),
	112: uint32(3270),
	113: uint32(3289),
	114: uint32(3305),
	115: uint32(3321),
	116: uint32(3337),
	117: uint32(3353),
	118: uint32(3369),
	119: uint32(3385),
	120: uint32(3401),
	121: uint32(3417),
	122: uint32(3433),
	123: uint32(3449),
	124: uint32(3465),
	125: uint32(3481),
	126: uint32(3497),
	127: uint32(3513),
	128: uint32(3529),
	129: uint32(3545),
	130: uint32(3561),
	131: uint32(3577),
	132: uint32(3593),
	133: uint32(3597),
}

var ts_parse_actions = [330]TSParseActionEntry{
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
		Fstate: uint16(libc.Int32FromInt32(116)),
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
		Fstate: uint16(libc.Int32FromInt32(136)),
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
		Fextra: libc.BoolUint8(true1 != 0),
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
		Fsymbol:     uint16(sym_program),
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
		Fstate: uint16(libc.Int32FromInt32(133)),
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
		Fstate: uint16(libc.Int32FromInt32(132)),
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
		Fstate: uint16(libc.Int32FromInt32(2)),
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
		Fstate: uint16(libc.Int32FromInt32(22)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(131)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_instruction),
		Fproduction_id: uint16(1),
	})))),
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
		Fstate: uint16(libc.Int32FromInt32(58)),
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
		Fstate: uint16(libc.Int32FromInt32(67)),
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
		Fstate: uint16(libc.Int32FromInt32(56)),
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
		Fstate: uint16(libc.Int32FromInt32(127)),
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
		Fstate: uint16(libc.Int32FromInt32(79)),
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
		Fstate: uint16(libc.Int32FromInt32(126)),
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
		Fstate: uint16(libc.Int32FromInt32(125)),
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
		Fstate: uint16(libc.Int32FromInt32(8)),
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
		Fstate: uint16(libc.Int32FromInt32(100)),
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
		Fstate: uint16(libc.Int32FromInt32(12)),
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
		Fstate: uint16(libc.Int32FromInt32(9)),
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
		Fstate: uint16(libc.Int32FromInt32(11)),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_instruction),
		Fproduction_id: uint16(1),
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
		Fstate: uint16(libc.Int32FromInt32(134)),
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
		Fstate: uint16(libc.Int32FromInt32(36)),
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
		Fstate: uint16(libc.Int32FromInt32(42)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(35)),
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
		Fstate: uint16(libc.Int32FromInt32(40)),
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
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_instruction),
		Fproduction_id: uint16(1),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_int),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_int),
	})))),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__expr),
	})))),
	67: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__tc_expr),
	})))),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__expr),
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
		Fstate: uint16(libc.Int32FromInt32(72)),
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
		Fcount: uint8(1),
	}})),
	73: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__tc_expr),
	})))),
	74: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	75: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__tc_expr),
	})))),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_int),
	})))),
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
		Fcount: uint8(1),
	}})),
	79: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_int),
	})))),
	80: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	81: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_reg),
	})))),
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
		Fcount: uint8(1),
	}})),
	83: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_reg),
	})))),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_ident),
	})))),
	86: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	87: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_ident),
	})))),
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
		Fsymbol:      uint16(sym_string),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_string),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	93: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_tc_infix),
		Fproduction_id: uint16(5),
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
		Fcount: uint8(1),
	}})),
	95: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_tc_infix),
		Fproduction_id: uint16(5),
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
		Fstate: uint16(libc.Int32FromInt32(28)),
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
		Fstate: uint16(libc.Int32FromInt32(29)),
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
		Fstate: uint16(libc.Int32FromInt32(31)),
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
		Fstate: uint16(libc.Int32FromInt32(27)),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(27)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_instruction_repeat2),
	})))),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(25)),
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
		Fstate: uint16(libc.Int32FromInt32(25)),
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
		Fcount: uint8(1),
	}})),
	113: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_instruction_repeat2),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	115: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_instruction),
		Fproduction_id: uint16(1),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_instruction_repeat2),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	119: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_instruction_repeat2),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(125)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_instruction_repeat2),
	})))),
	123: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	125: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_instruction_repeat2),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(12)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_instruction_repeat2),
	})))),
	129: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(2),
	}})),
	131: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_instruction_repeat2),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(11)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_meta),
		Fproduction_id: uint16(1),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(70)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_program),
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
		Fstate: uint16(libc.Int32FromInt32(53)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_program),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(26)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_const),
		Fproduction_id: uint16(4),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(33)),
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
		Fstate: uint16(libc.Int32FromInt32(33)),
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
		Fstate: uint16(libc.Int32FromInt32(34)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_program_repeat1),
	})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_program_repeat1),
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
		Fstate:      uint16(libc.Int32FromInt32(53)),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_program_repeat1),
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
		Fsymbol:      uint16(aux_sym_program_repeat1),
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
		Fsymbol:      uint16(aux_sym_program_repeat1),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_label),
		Fproduction_id: uint16(3),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fsymbol:      uint16(sym_label),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_program_repeat2),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_program_repeat2),
	})))),
	192: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(libc.Int32FromInt32(53)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fstate: uint16(libc.Int32FromInt32(57)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_list_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(73)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_list_repeat1),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	203: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_meta),
		Fproduction_id: uint16(1),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_meta),
		Fproduction_id: uint16(1),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(3)),
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
		Fstate: uint16(libc.Int32FromInt32(4)),
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
		Fstate: uint16(libc.Int32FromInt32(109)),
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
		Fstate: uint16(libc.Int32FromInt32(50)),
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
		Fstate: uint16(libc.Int32FromInt32(77)),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_ptr),
	})))),
	224: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(96)),
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
		Fcount:    uint8(1),
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
		Fsymbol:      uint16(aux_sym_instruction_repeat1),
	})))),
	228: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	229: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_instruction_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(5)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_program),
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
		Fstate: uint16(libc.Int32FromInt32(59)),
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
		Fstate: uint16(libc.Int32FromInt32(89)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_ptr),
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
		Fstate: uint16(libc.Int32FromInt32(90)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_meta_repeat3),
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
		Fsymbol:      uint16(aux_sym_meta_repeat3),
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
		Fstate:      uint16(libc.Int32FromInt32(88)),
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
		Fsymbol:      uint16(aux_sym_meta_repeat2),
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
		Fsymbol:      uint16(aux_sym_meta_repeat2),
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
		Fstate:      uint16(libc.Int32FromInt32(124)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_meta_repeat1),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	254: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_meta_repeat1),
	})))),
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
		Fstate:      uint16(libc.Int32FromInt32(105)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	257: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_list),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	259: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_ptr),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_list),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_ptr),
	})))),
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
		Fsymbol:      uint16(sym_list),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_ptr),
	})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	271: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	272: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	273: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	274: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(libc.Int32FromInt32(36)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	277: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_label),
		Fproduction_id: uint16(3),
	})))),
	278: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	279: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_label),
		Fproduction_id: uint16(3),
	})))),
	280: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	281: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_label),
	})))),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_label),
	})))),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_label),
		Fproduction_id: uint16(2),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__item),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(116)),
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
		Fcount: uint8(1),
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
		Fcount: uint8(1),
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
		Fstate: uint16(libc.Int32FromInt32(136)),
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
		Fextra: libc.BoolUint8(true1 != 0),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(92)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(95)),
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
		Fstate: uint16(libc.Int32FromInt32(6)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	311: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(libc.Int32FromInt32(119)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	319: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f uint8_t
		_ [7]byte
	}{f: uint8(TSParseActionTypeAccept)})),
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
		Fstate: uint16(libc.Int32FromInt32(32)),
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
		Fstate: uint16(libc.Int32FromInt32(113)),
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
		Fstate: uint16(libc.Int32FromInt32(39)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_line_comment),
	})))),
}

func tree_sitter_asm(tls *libc.TLS) (r uintptr) {
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

var __ccgo_ts1 = "end\x00\n\x00,\x00:\x00(\x00)\x00label\x00const\x00{\x00-\x00}\x00byte\x00word\x00dword\x00qword\x00ptr\x00[\x00+\x00]\x00*\x00rel\x00!\x00/\x00%\x00|\x00^\x00&\x00#\x00int_token1\x00int_token2\x00float\x00string_token1\x00string_token2\x00_reg\x00address\x00meta_ident\x00_ident\x00line_comment_token1\x00line_comment_token2\x00block_comment\x00program\x00_item\x00meta\x00instruction\x00_expr\x00list\x00_tc_expr\x00tc_infix\x00int\x00string\x00reg\x00ident\x00line_comment\x00program_repeat1\x00program_repeat2\x00meta_repeat1\x00meta_repeat2\x00meta_repeat3\x00instruction_repeat1\x00instruction_repeat2\x00list_repeat1\x00kind\x00lhs\x00name\x00op\x00rhs\x00value\x00"
