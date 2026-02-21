// Code generated for linux/arm64 by 'ccgo /tmp/grammar-gen-3578027329/preprocessed.c -o /tmp/grammar-gen-3578027329/grammar.go', DO NOT EDIT.

//go:build linux && arm64

package grammar_jsdoc

import (
	"reflect"
	"unsafe"

	"modernc.org/libc"
)

var _ reflect.Type
var _ unsafe.Pointer

const _GNU_SOURCE = 1
const _LP64 = 1
const _STDC_PREDEF_H = 1
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
const __BYTE_ORDER__ = "__ORDER_LITTLE_ENDIAN__"
const __CCGO__ = 1
const __CET__ = 3
const __CHAR_BIT__ = 8
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
const __DEC64_EPSILON__ = 1e-15
const __DEC64_MANT_DIG__ = 16
const __DEC64_MAX_EXP__ = 385
const __DEC64_MAX__ = "9.999999999999999E384"
const __DEC64_MIN__ = 1e-383
const __DEC64_SUBNORMAL_MIN__ = 0.000000000000001e-383
const __DECIMAL_BID_FORMAT__ = 1
const __DECIMAL_DIG__ = 21
const __DEC_EVAL_METHOD__ = 2
const __ELF__ = 1
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
const __FLT64X_DECIMAL_DIG__ = 21
const __FLT64X_DENORM_MIN__ = 3.64519953188247460252840593361941982e-4951
const __FLT64X_DIG__ = 18
const __FLT64X_EPSILON__ = 1.08420217248550443400745280086994171e-19
const __FLT64X_HAS_DENORM__ = 1
const __FLT64X_HAS_INFINITY__ = 1
const __FLT64X_HAS_QUIET_NAN__ = 1
const __FLT64X_IS_IEC_60559__ = 1
const __FLT64X_MANT_DIG__ = 64
const __FLT64X_MAX_10_EXP__ = 4932
const __FLT64X_MAX_EXP__ = 16384
const __FLT64X_MAX__ = "1.18973149535723176502126385303097021e+4932"
const __FLT64X_MIN__ = 3.36210314311209350626267781732175260e-4932
const __FLT64X_NORM_MAX__ = "1.18973149535723176502126385303097021e+4932"
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
const __FLT_EVAL_METHOD_TS_18661_3__ = 0
const __FLT_EVAL_METHOD__ = 0
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
const __GCC_HAVE_DWARF2_CFI_ASM = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_1 = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_2 = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_4 = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_8 = 1
const __GCC_IEC_559 = 2
const __GCC_IEC_559_COMPLEX = 2
const __GNUC_EXECUTION_CHARSET_NAME = "UTF-8"
const __GNUC_MINOR__ = 3
const __GNUC_PATCHLEVEL__ = 0
const __GNUC_STDC_INLINE__ = 1
const __GNUC_WIDE_EXECUTION_CHARSET_NAME = "UTF-32LE"
const __GNUC__ = 13
const __GXX_ABI_VERSION = 1018
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
const __INT_FAST16_MAX__ = 0x7fffffffffffffff
const __INT_FAST16_WIDTH__ = 64
const __INT_FAST32_MAX__ = 0x7fffffffffffffff
const __INT_FAST32_WIDTH__ = 64
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
const __LDBL_DECIMAL_DIG__ = 21
const __LDBL_DENORM_MIN__ = 3.64519953188247460252840593361941982e-4951
const __LDBL_DIG__ = 18
const __LDBL_EPSILON__ = 1.08420217248550443400745280086994171e-19
const __LDBL_HAS_DENORM__ = 1
const __LDBL_HAS_INFINITY__ = 1
const __LDBL_HAS_QUIET_NAN__ = 1
const __LDBL_IS_IEC_60559__ = 1
const __LDBL_MANT_DIG__ = 64
const __LDBL_MAX_10_EXP__ = 4932
const __LDBL_MAX_EXP__ = 16384
const __LDBL_MAX__ = "1.18973149535723176502126385303097021e+4932"
const __LDBL_MIN__ = 3.36210314311209350626267781732175260e-4932
const __LDBL_NORM_MAX__ = "1.18973149535723176502126385303097021e+4932"
const __LONG_LONG_MAX__ = 0x7fffffffffffffff
const __LONG_LONG_WIDTH__ = 64
const __LONG_MAX__ = 0x7fffffffffffffff
const __LONG_WIDTH__ = 64
const __LP64__ = 1
const __MMX_WITH_SSE__ = 1
const __MMX__ = 1
const __NO_INLINE__ = 1
const __ORDER_BIG_ENDIAN__ = 4321
const __ORDER_LITTLE_ENDIAN__ = 1234
const __ORDER_PDP_ENDIAN__ = 3412
const __PIC__ = 2
const __PIE__ = 2
const __PRAGMA_REDEFINE_EXTNAME = 1
const __PRETTY_FUNCTION__ = "__func__"
const __PTRDIFF_MAX__ = 0x7fffffffffffffff
const __PTRDIFF_WIDTH__ = 64
const __SCHAR_MAX__ = 0x7f
const __SCHAR_WIDTH__ = 8
const __SEG_FS = 1
const __SEG_GS = 1
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
const __SIZEOF_LONG__ = 8
const __SIZEOF_POINTER__ = 8
const __SIZEOF_PTRDIFF_T__ = 8
const __SIZEOF_SHORT__ = 2
const __SIZEOF_SIZE_T__ = 8
const __SIZEOF_WCHAR_T__ = 4
const __SIZEOF_WINT_T__ = 4
const __SIZE_MAX__ = 0xffffffffffffffff
const __SIZE_WIDTH__ = 64
const __SSE2_MATH__ = 1
const __SSE2__ = 1
const __SSE_MATH__ = 1
const __SSE__ = 1
const __SSP_STRONG__ = 3
const __STDC_HOSTED__ = 1
const __STDC_IEC_559_COMPLEX__ = 1
const __STDC_IEC_559__ = 1
const __STDC_IEC_60559_BFP__ = 201404
const __STDC_IEC_60559_COMPLEX__ = 201404
const __STDC_ISO_10646__ = 201706
const __STDC_UTF_16__ = 1
const __STDC_UTF_32__ = 1
const __STDC_VERSION__ = 201710
const __STDC__ = 1
const __UINT16_MAX__ = 0xffff
const __UINT32_MAX__ = 0xffffffff
const __UINT64_MAX__ = 0xffffffffffffffff
const __UINT8_MAX__ = 0xff
const __UINTMAX_MAX__ = 0xffffffffffffffff
const __UINTPTR_MAX__ = 0xffffffffffffffff
const __UINT_FAST16_MAX__ = 0xffffffffffffffff
const __UINT_FAST32_MAX__ = 0xffffffffffffffff
const __UINT_FAST64_MAX__ = 0xffffffffffffffff
const __UINT_FAST8_MAX__ = 0xff
const __UINT_LEAST16_MAX__ = 0xffff
const __UINT_LEAST32_MAX__ = 0xffffffff
const __UINT_LEAST64_MAX__ = 0xffffffffffffffff
const __UINT_LEAST8_MAX__ = 0xff
const __VERSION__ = "13.3.0"
const __WCHAR_MAX__ = 0x7fffffff
const __WCHAR_TYPE__ = "int"
const __WCHAR_WIDTH__ = 32
const __WINT_MAX__ = 0xffffffff
const __WINT_MIN__ = 0
const __WINT_WIDTH__ = 32
const __amd64 = 1
const __amd64__ = 1
const __code_model_small__ = 1
const __gnu_linux__ = 1
const __k8 = 1
const __k8__ = 1
const __linux = 1
const __linux__ = 1
const __pic__ = 2
const __pie__ = 2
const __restrict_arr = "restrict"
const __unix = 1
const __unix__ = 1
const __x86_64 = 1
const __x86_64__ = 1
const linux = 1
const unix = 1

type __builtin_va_list = uintptr

type __predefined_size_t = uint64

type __predefined_wchar_t = int32

type __predefined_ptrdiff_t = int64

type __u_char = uint8

type __u_short = uint16

type __u_int = uint32

type __u_long = uint64

type __int8_t = int8

type __uint8_t = uint8

type __int16_t = int16

type __uint16_t = uint16

type __int32_t = int32

type __uint32_t = uint32

type __int64_t = int64

type __uint64_t = uint64

type __int_least8_t = int8

type __uint_least8_t = uint8

type __int_least16_t = int16

type __uint_least16_t = uint16

type __int_least32_t = int32

type __uint_least32_t = uint32

type __int_least64_t = int64

type __uint_least64_t = uint64

type __quad_t = int64

type __u_quad_t = uint64

type __intmax_t = int64

type __uintmax_t = uint64

type __dev_t = uint64

type __uid_t = uint32

type __gid_t = uint32

type __ino_t = uint64

type __ino64_t = uint64

type __mode_t = uint32

type __nlink_t = uint64

type __off_t = int64

type __off64_t = int64

type __pid_t = int32

type __fsid_t = struct {
	F__val [2]int32
}

type __clock_t = int64

type __rlim_t = uint64

type __rlim64_t = uint64

type __id_t = uint32

type __time_t = int64

type __useconds_t = uint32

type __suseconds_t = int64

type __suseconds64_t = int64

type __daddr_t = int32

type __key_t = int32

type __clockid_t = int32

type __timer_t = uintptr

type __blksize_t = int64

type __blkcnt_t = int64

type __blkcnt64_t = int64

type __fsblkcnt_t = uint64

type __fsblkcnt64_t = uint64

type __fsfilcnt_t = uint64

type __fsfilcnt64_t = uint64

type __fsword_t = int64

type __ssize_t = int64

type __syscall_slong_t = int64

type __syscall_ulong_t = uint64

type __loff_t = int64

type __caddr_t = uintptr

type __intptr_t = int64

type __socklen_t = uint32

type __sig_atomic_t = int32

type int8_t = int8

type int16_t = int16

type int32_t = int32

type int64_t = int64

type uint8_t = uint8

type uint16_t = uint16

type uint32_t = uint32

type uint64_t = uint64

type int_least8_t = int8

type int_least16_t = int16

type int_least32_t = int32

type int_least64_t = int64

type uint_least8_t = uint8

type uint_least16_t = uint16

type uint_least32_t = uint32

type uint_least64_t = uint64

type int_fast8_t = int8

type int_fast16_t = int64

type int_fast32_t = int64

type int_fast64_t = int64

type uint_fast8_t = uint8

type uint_fast16_t = uint64

type uint_fast32_t = uint64

type uint_fast64_t = uint64

type intptr_t = int64

type uintptr_t = uint64

type intmax_t = int64

type uintmax_t = uint64

type size_t = uint64

type wchar_t = int32

type div_t = struct {
	Fquot int32
	Frem  int32
}

type ldiv_t = struct {
	Fquot int64
	Frem  int64
}

type lldiv_t = struct {
	Fquot int64
	Frem  int64
}

type u_char = uint8

type u_short = uint16

type u_int = uint32

type u_long = uint64

type quad_t = int64

type u_quad_t = uint64

type fsid_t = struct {
	F__val [2]int32
}

type loff_t = int64

type ino_t = uint64

type dev_t = uint64

type gid_t = uint32

type mode_t = uint32

type nlink_t = uint64

type uid_t = uint32

type off_t = int64

type pid_t = int32

type id_t = uint32

type ssize_t = int64

type daddr_t = int32

type caddr_t = uintptr

type key_t = int32

type clock_t = int64

type clockid_t = int32

type time_t = int64

type timer_t = uintptr

type ulong = uint64

type ushort = uint16

type uint1 = uint32

type u_int8_t = uint8

type u_int16_t = uint16

type u_int32_t = uint32

type u_int64_t = uint64

type register_t = int32

func __bswap_16(tls *libc.TLS, __bsx __uint16_t) (r __uint16_t) {
	return libc.X__builtin_bswap16(tls, __bsx)
}

func __bswap_32(tls *libc.TLS, __bsx __uint32_t) (r __uint32_t) {
	return libc.X__builtin_bswap32(tls, __bsx)
}

func __bswap_64(tls *libc.TLS, __bsx __uint64_t) (r __uint64_t) {
	return libc.X__builtin_bswap64(tls, __bsx)
}

func __uint16_identity(tls *libc.TLS, __x __uint16_t) (r __uint16_t) {
	return __x
}

func __uint32_identity(tls *libc.TLS, __x __uint32_t) (r __uint32_t) {
	return __x
}

func __uint64_identity(tls *libc.TLS, __x __uint64_t) (r __uint64_t) {
	return __x
}

type __sigset_t = struct {
	F__val [16]uint64
}

type sigset_t = struct {
	F__val [16]uint64
}

type timeval = struct {
	Ftv_sec  __time_t
	Ftv_usec __suseconds_t
}

type timespec = struct {
	Ftv_sec  __time_t
	Ftv_nsec __syscall_slong_t
}

type suseconds_t = int64

type __fd_mask = int64

type fd_set = struct {
	F__fds_bits [16]__fd_mask
}

type fd_mask = int64

type blksize_t = int64

type blkcnt_t = int64

type fsblkcnt_t = uint64

type fsfilcnt_t = uint64

type __atomic_wide_counter = struct {
	F__value32 [0]struct {
		F__low  uint32
		F__high uint32
	}
	F__value64 uint64
}

type __pthread_list_t = struct {
	F__prev uintptr
	F__next uintptr
}

type __pthread_internal_list = __pthread_list_t

type __pthread_slist_t = struct {
	F__next uintptr
}

type __pthread_internal_slist = __pthread_slist_t

type __pthread_mutex_s = struct {
	F__lock    int32
	F__count   uint32
	F__owner   int32
	F__nusers  uint32
	F__kind    int32
	F__spins   int16
	F__elision int16
	F__list    __pthread_list_t
}

type __pthread_rwlock_arch_t = struct {
	F__readers       uint32
	F__writers       uint32
	F__wrphase_futex uint32
	F__writers_futex uint32
	F__pad3          uint32
	F__pad4          uint32
	F__cur_writer    int32
	F__shared        int32
	F__rwelision     int8
	F__pad1          [7]uint8
	F__pad2          uint64
	F__flags         uint32
}

type __pthread_cond_s = struct {
	F__wseq         __atomic_wide_counter
	F__g1_start     __atomic_wide_counter
	F__g_refs       [2]uint32
	F__g_size       [2]uint32
	F__g1_orig_size uint32
	F__wrefs        uint32
	F__g_signals    [2]uint32
}

type __tss_t = uint32

type __thrd_t = uint64

type __once_flag = struct {
	F__data int32
}

type pthread_t = uint64

type pthread_mutexattr_t = struct {
	F__align [0]int32
	F__size  [4]uint8
}

type pthread_condattr_t = struct {
	F__align [0]int32
	F__size  [4]uint8
}

type pthread_key_t = uint32

type pthread_once_t = int32

type pthread_attr_t = struct {
	F__align [0]int64
	F__size  [56]uint8
}

type pthread_mutex_t = struct {
	F__size  [0][40]uint8
	F__align [0]int64
	F__data  __pthread_mutex_s
}

type pthread_cond_t = struct {
	F__size  [0][48]uint8
	F__align [0]int64
	F__data  __pthread_cond_s
}

type pthread_rwlock_t = struct {
	F__size  [0][56]uint8
	F__align [0]int64
	F__data  __pthread_rwlock_arch_t
}

type pthread_rwlockattr_t = struct {
	F__align [0]int64
	F__size  [8]uint8
}

type pthread_spinlock_t = int32

type pthread_barrier_t = struct {
	F__align [0]int64
	F__size  [32]uint8
}

type pthread_barrierattr_t = struct {
	F__align [0]int32
	F__size  [4]uint8
}

type random_data = struct {
	Ffptr      uintptr
	Frptr      uintptr
	Fstate     uintptr
	Frand_type int32
	Frand_deg  int32
	Frand_sep  int32
	Fend_ptr   uintptr
}

type drand48_data = struct {
	F__x     [3]uint16
	F__old_x [3]uint16
	F__c     uint16
	F__init  uint16
	F__a     uint64
}

type __compar_fn_t = uintptr

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

func set_contains(tls *libc.TLS, ranges uintptr, len1 uint32_t, lookahead int32_t) (r uint8) {
	var half_size, index, mid_index, size uint32_t
	var range_token, range_token1 uintptr
	_, _, _, _, _, _ = half_size, index, mid_index, range_token, range_token1, size
	index = uint32(0)
	size = len1 - index
	for size > uint32(1) {
		half_size = size / uint32(2)
		mid_index = index + half_size
		range_token = ranges + uintptr(mid_index)*8
		if lookahead >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && lookahead <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
			return libc.BoolUint8(1 != 0)
		} else {
			if lookahead > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				index = mid_index
			}
		}
		size = size - half_size
	}
	range_token1 = ranges + uintptr(index)*8
	return libc.BoolUint8(lookahead >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && lookahead <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
}

type TokenType = int32

const TYPE_TOKEN = 0

func tree_sitter_jsdoc_external_scanner_create(tls *libc.TLS) (r uintptr) {
	return libc.UintptrFromInt32(0)
}

func tree_sitter_jsdoc_external_scanner_destroy(tls *libc.TLS, payload uintptr) {
}

func tree_sitter_jsdoc_external_scanner_serialize(tls *libc.TLS, payload uintptr, buffer uintptr) (r uint32) {
	return uint32(0)
}

func tree_sitter_jsdoc_external_scanner_deserialize(tls *libc.TLS, payload uintptr, buffer uintptr, length uint32) {
}

func scan_for_type(tls *libc.TLS, lexer uintptr) (r uint8) {
	var stack int32
	_ = stack
	stack = 0
	for int32(1) != 0 {
		if (*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Feof})))(tls, lexer) != 0 {
			return libc.BoolUint8(0 != 0)
		}
		switch (*TSLexer)(unsafe.Pointer(lexer)).Flookahead {
		case int32('{'):
			stack = stack + 1
		case int32('}'):
			stack = stack - 1
			if stack == -int32(1) {
				return libc.BoolUint8(1 != 0)
			}
		case int32('\n'):
			fallthrough
		case int32('\000'):
			return libc.BoolUint8(0 != 0)
		default:
		}
		(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(0 != 0))
	}
	return r
}

func tree_sitter_jsdoc_external_scanner_scan(tls *libc.TLS, payload uintptr, lexer uintptr, valid_symbols uintptr) (r uint8) {
	if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(TYPE_TOKEN))) != 0 && scan_for_type(tls, lexer) != 0 {
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(TYPE_TOKEN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return libc.BoolUint8(1 != 0)
	}
	return libc.BoolUint8(0 != 0)
}

type ts_symbol_identifiers = int32

const anon_sym_LBRACE = 1
const anon_sym_RBRACE = 2
const sym__inline_tag_false_positive = 3
const sym_tag_name_with_argument = 4
const sym_tag_name_with_type = 5
const sym_tag_name = 6
const anon_sym_COLON = 7
const anon_sym_SLASH = 8
const anon_sym_DOT = 9
const anon_sym_POUND = 10
const anon_sym_TILDE = 11
const anon_sym_LBRACK = 12
const anon_sym_COMMA = 13
const anon_sym_RBRACK = 14
const anon_sym_BQUOTE_BQUOTE_BQUOTE = 15
const sym_code_block_language = 16
const sym_code_block_line = 17
const anon_sym_EQ = 18
const sym_identifier = 19
const sym_number = 20
const sym__text = 21
const anon_sym_SLASH2 = 22
const anon_sym_STAR = 23
const sym_type = 24
const sym_document = 25
const sym_description = 26
const sym_tag = 27
const sym_inline_tag = 28
const sym_expression = 29
const sym_qualified_expression = 30
const sym_path_expression = 31
const sym_member_expression = 32
const sym_array_expression = 33
const sym_code_block = 34
const sym_optional_identifier = 35
const sym__begin = 36
const sym__end = 37
const aux_sym_document_repeat1 = 38
const aux_sym_description_repeat1 = 39
const aux_sym_array_expression_repeat1 = 40
const aux_sym_code_block_repeat1 = 41
const aux_sym__begin_repeat1 = 42

var ts_symbol_names = [43]uintptr{
	0:  __ccgo_ts,
	1:  __ccgo_ts + 4,
	2:  __ccgo_ts + 6,
	3:  __ccgo_ts + 8,
	4:  __ccgo_ts + 35,
	5:  __ccgo_ts + 35,
	6:  __ccgo_ts + 35,
	7:  __ccgo_ts + 44,
	8:  __ccgo_ts + 46,
	9:  __ccgo_ts + 48,
	10: __ccgo_ts + 50,
	11: __ccgo_ts + 52,
	12: __ccgo_ts + 54,
	13: __ccgo_ts + 56,
	14: __ccgo_ts + 58,
	15: __ccgo_ts + 60,
	16: __ccgo_ts + 64,
	17: __ccgo_ts + 84,
	18: __ccgo_ts + 100,
	19: __ccgo_ts + 102,
	20: __ccgo_ts + 113,
	21: __ccgo_ts + 120,
	22: __ccgo_ts + 46,
	23: __ccgo_ts + 126,
	24: __ccgo_ts + 128,
	25: __ccgo_ts + 133,
	26: __ccgo_ts + 142,
	27: __ccgo_ts + 154,
	28: __ccgo_ts + 158,
	29: __ccgo_ts + 169,
	30: __ccgo_ts + 180,
	31: __ccgo_ts + 201,
	32: __ccgo_ts + 217,
	33: __ccgo_ts + 235,
	34: __ccgo_ts + 252,
	35: __ccgo_ts + 263,
	36: __ccgo_ts + 283,
	37: __ccgo_ts + 290,
	38: __ccgo_ts + 295,
	39: __ccgo_ts + 312,
	40: __ccgo_ts + 332,
	41: __ccgo_ts + 357,
	42: __ccgo_ts + 376,
}

var ts_symbol_map = [43]TSSymbol{
	1:  uint16(anon_sym_LBRACE),
	2:  uint16(anon_sym_RBRACE),
	3:  uint16(sym__inline_tag_false_positive),
	4:  uint16(sym_tag_name),
	5:  uint16(sym_tag_name),
	6:  uint16(sym_tag_name),
	7:  uint16(anon_sym_COLON),
	8:  uint16(anon_sym_SLASH),
	9:  uint16(anon_sym_DOT),
	10: uint16(anon_sym_POUND),
	11: uint16(anon_sym_TILDE),
	12: uint16(anon_sym_LBRACK),
	13: uint16(anon_sym_COMMA),
	14: uint16(anon_sym_RBRACK),
	15: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	16: uint16(sym_code_block_language),
	17: uint16(sym_code_block_line),
	18: uint16(anon_sym_EQ),
	19: uint16(sym_identifier),
	20: uint16(sym_number),
	21: uint16(sym__text),
	22: uint16(anon_sym_SLASH),
	23: uint16(anon_sym_STAR),
	24: uint16(sym_type),
	25: uint16(sym_document),
	26: uint16(sym_description),
	27: uint16(sym_tag),
	28: uint16(sym_inline_tag),
	29: uint16(sym_expression),
	30: uint16(sym_qualified_expression),
	31: uint16(sym_path_expression),
	32: uint16(sym_member_expression),
	33: uint16(sym_array_expression),
	34: uint16(sym_code_block),
	35: uint16(sym_optional_identifier),
	36: uint16(sym__begin),
	37: uint16(sym__end),
	38: uint16(aux_sym_document_repeat1),
	39: uint16(aux_sym_description_repeat1),
	40: uint16(aux_sym_array_expression_repeat1),
	41: uint16(aux_sym_code_block_repeat1),
	42: uint16(aux_sym__begin_repeat1),
}

var ts_symbol_metadata = [43]TSSymbolMetadata{
	0: {
		Fnamed: libc.BoolUint8(1 != 0),
	},
	1: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	2: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	3: {
		Fnamed: libc.BoolUint8(1 != 0),
	},
	4: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	5: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	6: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	7: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	8: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	9: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	10: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	11: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	12: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	13: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	14: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	15: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	16: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	17: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	18: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	19: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	20: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	21: {
		Fnamed: libc.BoolUint8(1 != 0),
	},
	22: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	23: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	24: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	25: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	26: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	27: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	28: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	29: {
		Fnamed:     libc.BoolUint8(1 != 0),
		Fsupertype: libc.BoolUint8(1 != 0),
	},
	30: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	31: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	32: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	33: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	34: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	35: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	36: {
		Fnamed: libc.BoolUint8(1 != 0),
	},
	37: {
		Fnamed: libc.BoolUint8(1 != 0),
	},
	38: {},
	39: {},
	40: {},
	41: {},
	42: {},
}

type ts_field_identifiers = int32

const field_value = 1

var ts_field_names = [2]uintptr{
	0: libc.UintptrFromInt32(0),
	1: __ccgo_ts + 391,
}

var ts_field_map_slices = [2]TSMapSlice{
	1: {
		Flength: uint16(1),
	},
}

var ts_field_map_entries = [1]TSFieldMapEntry{
	0: {
		Ffield_id:    uint16(field_value),
		Fchild_index: uint8(3),
	},
}

var ts_alias_sequences = [2][6]TSSymbol{}

var ts_non_terminal_alias_map = [1]uint16_t{}

var ts_primary_state_ids = [99]TSStateId{
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
	14: uint16(13),
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
	38: uint16(5),
	39: uint16(34),
	40: uint16(8),
	41: uint16(10),
	42: uint16(42),
	43: uint16(43),
	44: uint16(6),
	45: uint16(45),
	46: uint16(46),
	47: uint16(47),
	48: uint16(48),
	49: uint16(49),
	50: uint16(48),
	51: uint16(11),
	52: uint16(12),
	53: uint16(21),
	54: uint16(26),
	55: uint16(55),
	56: uint16(36),
	57: uint16(17),
	58: uint16(20),
	59: uint16(16),
	60: uint16(18),
	61: uint16(33),
	62: uint16(22),
	63: uint16(37),
	64: uint16(64),
	65: uint16(65),
	66: uint16(66),
	67: uint16(67),
	68: uint16(68),
	69: uint16(69),
	70: uint16(69),
	71: uint16(71),
	72: uint16(71),
	73: uint16(73),
	74: uint16(74),
	75: uint16(75),
	76: uint16(76),
	77: uint16(77),
	78: uint16(78),
	79: uint16(73),
	80: uint16(78),
	81: uint16(76),
	82: uint16(75),
	83: uint16(83),
	84: uint16(83),
	85: uint16(85),
	86: uint16(86),
	87: uint16(87),
	88: uint16(88),
	89: uint16(89),
	90: uint16(90),
	91: uint16(91),
	92: uint16(92),
	93: uint16(93),
	94: uint16(85),
	95: uint16(95),
	96: uint16(96),
	97: uint16(96),
	98: uint16(92),
}

var ts_supertype_symbols = [1]TSSymbol{
	0: uint16(sym_expression),
}

var ts_supertype_map_slices = [30]TSMapSlice{
	29: {
		Flength: uint16(6),
	},
}

var ts_supertype_map_entries = [6]TSSymbol{
	0: uint16(sym_array_expression),
	1: uint16(sym_identifier),
	2: uint16(sym_member_expression),
	3: uint16(sym_number),
	4: uint16(sym_path_expression),
	5: uint16(sym_qualified_expression),
}

func ts_lex(tls *libc.TLS, lexer uintptr, state TSStateId) (r uint8) {
	var eof, result, skip uint8
	var i, i1, i2 uint32_t
	var lookahead int32_t
	_, _, _, _, _, _, _ = eof, i, i1, i2, lookahead, result, skip
	result = libc.BoolUint8(0 != 0)
	skip = libc.BoolUint8(0 != 0)
	eof = libc.BoolUint8(0 != 0)
	goto start
	goto next_state
next_state:
	;
	(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, skip)
	goto start
start:
	;
	skip = libc.BoolUint8(0 != 0)
	lookahead = (*TSLexer)(unsafe.Pointer(lexer)).Flookahead
	eof = (*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Feof})))(tls, lexer)
	switch libc.Int32FromUint16(state) {
	case 0:
		if eof != 0 {
			state = uint16(58)
			goto next_state
		}
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(56)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(176)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(179)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(175)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(173)
			goto next_state
		}
		if lookahead == int32('0') {
			state = uint16(211)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(172)
			goto next_state
		}
		if lookahead == int32('=') {
			state = uint16(208)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(42)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(178)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(180)
			goto next_state
		}
		if lookahead == int32('`') {
			state = uint16(240)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(60)
			goto next_state
		}
		if lookahead == int32('}') {
			state = uint16(61)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(177)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(57)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(212)
			goto next_state
		}
		if int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(188)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('*') {
			state = uint16(242)
			goto next_state
		}
		return result
	case int32(1):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(1)
			goto next_state
		}
		if lookahead == int32('*') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(1)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(241)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(243)
			goto next_state
		}
		if lookahead == int32('0') {
			state = uint16(211)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(42)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(178)
			goto next_state
		}
		if lookahead == int32('`') {
			state = uint16(240)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(59)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(1)
			goto next_state
		}
		if int32(0x0b) <= lookahead && lookahead <= int32('\r') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(2)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(212)
			goto next_state
		}
		if lookahead == int32('$') || int32('A') <= lookahead && lookahead <= int32('Z') || int32('_') <= lookahead && lookahead <= int32('z') {
			state = uint16(209)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('}') {
			state = uint16(242)
			goto next_state
		}
		return result
	case int32(2):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(1)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(241)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(243)
			goto next_state
		}
		if lookahead == int32('0') {
			state = uint16(211)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(42)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(178)
			goto next_state
		}
		if lookahead == int32('`') {
			state = uint16(240)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(59)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(2)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(212)
			goto next_state
		}
		if lookahead == int32('$') || int32('A') <= lookahead && lookahead <= int32('Z') || int32('_') <= lookahead && lookahead <= int32('z') {
			state = uint16(209)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('*') && lookahead != int32('}') {
			state = uint16(242)
			goto next_state
		}
		return result
	case int32(3):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(3)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(176)
			goto next_state
		}
		if lookahead == int32('*') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(3)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(174)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(243)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(42)
			goto next_state
		}
		if lookahead == int32('`') {
			state = uint16(240)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(177)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(3)
			goto next_state
		}
		if int32(0x0b) <= lookahead && lookahead <= int32('\r') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(4)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('{') && lookahead != int32('}') && lookahead != int32('~') {
			state = uint16(242)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(3)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(176)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(174)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(243)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(42)
			goto next_state
		}
		if lookahead == int32('`') {
			state = uint16(240)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(177)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(4)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('*') && lookahead != int32('{') && lookahead != int32('}') && lookahead != int32('~') {
			state = uint16(242)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(5)
			goto next_state
		}
		if lookahead == int32('*') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(5)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(243)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(42)
			goto next_state
		}
		if lookahead == int32('`') {
			state = uint16(240)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(60)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(5)
			goto next_state
		}
		if int32(0x0b) <= lookahead && lookahead <= int32('\r') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(6)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('}') {
			state = uint16(242)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(5)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(243)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(42)
			goto next_state
		}
		if lookahead == int32('`') {
			state = uint16(240)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(60)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(6)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('*') && lookahead != int32('}') {
			state = uint16(242)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(7)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(176)
			goto next_state
		}
		if lookahead == int32('*') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(7)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(174)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(243)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(172)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(42)
			goto next_state
		}
		if lookahead == int32('`') {
			state = uint16(240)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(177)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(7)
			goto next_state
		}
		if int32(0x0b) <= lookahead && lookahead <= int32('\r') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(9)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('{') && lookahead != int32('}') && lookahead != int32('~') {
			state = uint16(242)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(7)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(176)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(174)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(173)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(172)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(42)
			goto next_state
		}
		if lookahead == int32('`') {
			state = uint16(240)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(177)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(9)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('*') && lookahead != int32('{') && lookahead != int32('}') && lookahead != int32('~') {
			state = uint16(242)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(7)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(176)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(174)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(243)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(172)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(42)
			goto next_state
		}
		if lookahead == int32('`') {
			state = uint16(240)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(177)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(9)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('*') && lookahead != int32('{') && lookahead != int32('}') && lookahead != int32('~') {
			state = uint16(242)
			goto next_state
		}
		return result
	case int32(10):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(10)
			goto next_state
		}
		if lookahead == int32('*') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(10)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(179)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(47)
			goto next_state
		}
		if lookahead == int32('0') {
			state = uint16(213)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(51)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(178)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(180)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(10)
			goto next_state
		}
		if int32(0x0b) <= lookahead && lookahead <= int32('\r') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(11)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(214)
			goto next_state
		}
		if lookahead == int32('$') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(209)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(10)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(179)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(47)
			goto next_state
		}
		if lookahead == int32('0') {
			state = uint16(213)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(51)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(178)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(180)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(11)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(214)
			goto next_state
		}
		if lookahead == int32('$') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(209)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(13)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(244)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(243)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(42)
			goto next_state
		}
		if lookahead == int32('`') {
			state = uint16(240)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(12)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(242)
			goto next_state
		}
		return result
	case int32(13):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(13)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(245)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(243)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(42)
			goto next_state
		}
		if lookahead == int32('`') {
			state = uint16(240)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(13)
			goto next_state
		}
		if int32(0x0b) <= lookahead && lookahead <= int32('\r') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(12)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(242)
			goto next_state
		}
		return result
	case int32(14):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(176)
			goto next_state
		}
		if lookahead == int32('*') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(179)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(174)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(172)
			goto next_state
		}
		if lookahead == int32('=') {
			state = uint16(208)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(180)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(177)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(14)
			goto next_state
		}
		if int32(0x0b) <= lookahead && lookahead <= int32('\r') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(16)
			goto next_state
		}
		return result
	case int32(15):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(176)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(179)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(174)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(173)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(172)
			goto next_state
		}
		if lookahead == int32('=') {
			state = uint16(208)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(180)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(177)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(16)
			goto next_state
		}
		return result
	case int32(16):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(176)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(179)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(174)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(172)
			goto next_state
		}
		if lookahead == int32('=') {
			state = uint16(208)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(180)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(177)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(16)
			goto next_state
		}
		return result
	case int32(17):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(17)
			goto next_state
		}
		if lookahead == int32('*') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(17)
			goto next_state
		}
		if lookahead == int32('`') {
			state = uint16(240)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(60)
			goto next_state
		}
		if lookahead == int32('}') {
			state = uint16(61)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(17)
			goto next_state
		}
		if int32(0x0b) <= lookahead && lookahead <= int32('\r') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(18)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('@') {
			state = uint16(242)
			goto next_state
		}
		return result
	case int32(18):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(17)
			goto next_state
		}
		if lookahead == int32('`') {
			state = uint16(240)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(60)
			goto next_state
		}
		if lookahead == int32('}') {
			state = uint16(61)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(18)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('*') && lookahead != int32('@') {
			state = uint16(242)
			goto next_state
		}
		return result
	case int32(19):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(20)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(196)
			goto next_state
		}
		if lookahead == int32('`') {
			state = uint16(204)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			state = uint16(196)
			goto next_state
		}
		if int32(0x0b) <= lookahead && lookahead <= int32('\r') {
			state = uint16(198)
			goto next_state
		}
		if int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(187)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(200)
			goto next_state
		}
		return result
	case int32(20):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(20)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(21)
			goto next_state
		}
		if lookahead == int32('`') {
			state = uint16(41)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			state = uint16(21)
			goto next_state
		}
		if int32(0x0b) <= lookahead && lookahead <= int32('\r') {
			state = uint16(23)
			goto next_state
		}
		if int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(193)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(53)
			goto next_state
		}
		return result
	case int32(21):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(20)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(22)
			goto next_state
		}
		if lookahead == int32('`') {
			state = uint16(41)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			state = uint16(22)
			goto next_state
		}
		if int32(0x0b) <= lookahead && lookahead <= int32('\r') {
			state = uint16(24)
			goto next_state
		}
		if int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(195)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(54)
			goto next_state
		}
		return result
	case int32(22):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(20)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(19)
			goto next_state
		}
		if lookahead == int32('`') {
			state = uint16(41)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			state = uint16(19)
			goto next_state
		}
		if int32(0x0b) <= lookahead && lookahead <= int32('\r') {
			state = uint16(26)
			goto next_state
		}
		if int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(186)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(40)
			goto next_state
		}
		return result
	case int32(23):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(20)
			goto next_state
		}
		if lookahead == int32('`') {
			state = uint16(41)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(24)
			goto next_state
		}
		if int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(195)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(54)
			goto next_state
		}
		return result
	case int32(24):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(20)
			goto next_state
		}
		if lookahead == int32('`') {
			state = uint16(41)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(26)
			goto next_state
		}
		if int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(186)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(40)
			goto next_state
		}
		return result
	case int32(25):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(20)
			goto next_state
		}
		if lookahead == int32('`') {
			state = uint16(41)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(23)
			goto next_state
		}
		if int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(193)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(53)
			goto next_state
		}
		return result
	case int32(26):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(20)
			goto next_state
		}
		if lookahead == int32('`') {
			state = uint16(204)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(198)
			goto next_state
		}
		if int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(187)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(200)
			goto next_state
		}
		return result
	case int32(27):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(28)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(197)
			goto next_state
		}
		if lookahead == int32('`') {
			state = uint16(204)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			state = uint16(197)
			goto next_state
		}
		if int32(0x0b) <= lookahead && lookahead <= int32('\r') {
			state = uint16(199)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(200)
			goto next_state
		}
		return result
	case int32(28):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(28)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(29)
			goto next_state
		}
		if lookahead == int32('`') {
			state = uint16(41)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			state = uint16(29)
			goto next_state
		}
		if int32(0x0b) <= lookahead && lookahead <= int32('\r') {
			state = uint16(32)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(53)
			goto next_state
		}
		return result
	case int32(29):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(28)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(30)
			goto next_state
		}
		if lookahead == int32('`') {
			state = uint16(41)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			state = uint16(30)
			goto next_state
		}
		if int32(0x0b) <= lookahead && lookahead <= int32('\r') {
			state = uint16(33)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(54)
			goto next_state
		}
		return result
	case int32(30):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(28)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(27)
			goto next_state
		}
		if lookahead == int32('`') {
			state = uint16(41)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			state = uint16(27)
			goto next_state
		}
		if int32(0x0b) <= lookahead && lookahead <= int32('\r') {
			state = uint16(34)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(40)
			goto next_state
		}
		return result
	case int32(31):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(28)
			goto next_state
		}
		if lookahead == int32('`') {
			state = uint16(41)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(32)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(53)
			goto next_state
		}
		return result
	case int32(32):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(28)
			goto next_state
		}
		if lookahead == int32('`') {
			state = uint16(41)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(33)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(54)
			goto next_state
		}
		return result
	case int32(33):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(28)
			goto next_state
		}
		if lookahead == int32('`') {
			state = uint16(41)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(34)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(40)
			goto next_state
		}
		return result
	case int32(34):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(28)
			goto next_state
		}
		if lookahead == int32('`') {
			state = uint16(204)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(199)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(200)
			goto next_state
		}
		return result
	case int32(35):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(35)
			goto next_state
		}
		if lookahead == int32('*') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(35)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(243)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(42)
			goto next_state
		}
		if lookahead == int32('`') {
			state = uint16(240)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(59)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(35)
			goto next_state
		}
		if int32(0x0b) <= lookahead && lookahead <= int32('\r') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(36)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('}') {
			state = uint16(242)
			goto next_state
		}
		return result
	case int32(36):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(35)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(243)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(42)
			goto next_state
		}
		if lookahead == int32('`') {
			state = uint16(240)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(59)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(36)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('*') && lookahead != int32('}') {
			state = uint16(242)
			goto next_state
		}
		return result
	case int32(37):
		if lookahead == int32('*') {
			state = uint16(239)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(37)
			goto next_state
		}
		return result
	case int32(38):
		if lookahead == int32('`') {
			state = uint16(181)
			goto next_state
		}
		return result
	case int32(39):
		if lookahead == int32('`') {
			state = uint16(205)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(201)
			goto next_state
		}
		return result
	case int32(40):
		if lookahead == int32('`') {
			state = uint16(205)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(41):
		if lookahead == int32('`') {
			state = uint16(38)
			goto next_state
		}
		return result
	case int32(42):
		i = uint32(0)
		for {
			if !(uint64(i) < libc.Uint64FromInt64(48)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token[i]) == lookahead {
				state = map_token[i+uint32(1)]
				goto next_state
			}
			goto _1
		_1:
			;
			i = i + uint32(2)
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('d') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(43):
		if lookahead == int32('+') || lookahead == int32('-') {
			state = uint16(49)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(226)
			goto next_state
		}
		return result
	case int32(44):
		if lookahead == int32('0') || lookahead == int32('1') {
			state = uint16(220)
			goto next_state
		}
		return result
	case int32(45):
		if int32('0') <= lookahead && lookahead <= int32('7') {
			state = uint16(221)
			goto next_state
		}
		return result
	case int32(46):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(214)
			goto next_state
		}
		return result
	case int32(47):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(219)
			goto next_state
		}
		return result
	case int32(48):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(223)
			goto next_state
		}
		return result
	case int32(49):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(226)
			goto next_state
		}
		return result
	case int32(50):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(222)
			goto next_state
		}
		return result
	case int32(51):
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(52):
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('`') {
			state = uint16(55)
			goto next_state
		}
		return result
	case int32(53):
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('`') {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(54):
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('`') {
			state = uint16(39)
			goto next_state
		}
		return result
	case int32(55):
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(205)
			goto next_state
		}
		return result
	case int32(56):
		if eof != 0 {
			state = uint16(58)
			goto next_state
		}
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(56)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(176)
			goto next_state
		}
		if lookahead == int32('*') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(56)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(179)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(175)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(243)
			goto next_state
		}
		if lookahead == int32('0') {
			state = uint16(211)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(172)
			goto next_state
		}
		if lookahead == int32('=') {
			state = uint16(208)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(42)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(178)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(180)
			goto next_state
		}
		if lookahead == int32('`') {
			state = uint16(240)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(60)
			goto next_state
		}
		if lookahead == int32('}') {
			state = uint16(61)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(177)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(56)
			goto next_state
		}
		if int32(0x0b) <= lookahead && lookahead <= int32('\r') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(57)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(212)
			goto next_state
		}
		if int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(188)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(242)
			goto next_state
		}
		return result
	case int32(57):
		if eof != 0 {
			state = uint16(58)
			goto next_state
		}
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(56)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(176)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(179)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(175)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(243)
			goto next_state
		}
		if lookahead == int32('0') {
			state = uint16(211)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(172)
			goto next_state
		}
		if lookahead == int32('=') {
			state = uint16(208)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(42)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(178)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(180)
			goto next_state
		}
		if lookahead == int32('`') {
			state = uint16(240)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(60)
			goto next_state
		}
		if lookahead == int32('}') {
			state = uint16(61)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(177)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(57)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(212)
			goto next_state
		}
		if int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(188)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('*') {
			state = uint16(242)
			goto next_state
		}
		return result
	case int32(58):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(0)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(59):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(60):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('@') && lookahead != int32('}') {
			state = uint16(63)
			goto next_state
		}
		return result
	case int32(61):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(62):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__inline_tag_false_positive)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(63):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__inline_tag_false_positive)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('}') {
			state = uint16(62)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('@') {
			state = uint16(63)
			goto next_state
		}
		return result
	case int32(64):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name_with_argument)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(148)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(65):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name_with_argument)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('s') {
			state = uint16(137)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(66):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name_with_argument)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(67):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name_with_type)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('d') {
			state = uint16(94)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(68):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name_with_type)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('s') {
			state = uint16(69)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(69):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name_with_type)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(70):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('a') {
			state = uint16(113)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(121)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('b') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(71):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('a') {
			state = uint16(159)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('b') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(72):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('a') {
			state = uint16(118)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('b') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(73):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('a') {
			state = uint16(82)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('b') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(74):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('a') {
			state = uint16(149)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(131)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('b') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(75):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('a') {
			state = uint16(112)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('b') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(76):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('a') {
			state = uint16(151)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('b') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(77):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('a') {
			state = uint16(116)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('b') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(78):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('a') {
			state = uint16(85)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('b') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(79):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('b') {
			state = uint16(73)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(80):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('b') {
			state = uint16(98)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(81):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('c') {
			state = uint16(83)
			goto next_state
		}
		if lookahead == int32('l') {
			state = uint16(107)
			goto next_state
		}
		if lookahead == int32('p') {
			state = uint16(104)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(102)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(82):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('c') {
			state = uint16(111)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(83):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('c') {
			state = uint16(96)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(84):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('c') {
			state = uint16(163)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(85):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('c') {
			state = uint16(89)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(86):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('c') {
			state = uint16(160)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(87):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('d') {
			state = uint16(151)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(88):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(157)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(89):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(66)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(90):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(122)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(91):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(65)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(92):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(67)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(93):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(117)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(169)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(94):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(100)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(95):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(125)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(96):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(155)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(97):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(151)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(98):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(141)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(99):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(127)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(100):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('f') {
			state = uint16(66)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(101):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('f') {
			state = uint16(110)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(102):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('g') {
			state = uint16(119)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(103):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('h') {
			state = uint16(106)
			goto next_state
		}
		if lookahead == int32('y') {
			state = uint16(138)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(104):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('i') {
			state = uint16(66)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(105):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('i') {
			state = uint16(150)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(124)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(106):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('i') {
			state = uint16(152)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(129)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(107):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('i') {
			state = uint16(76)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(108):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('i') {
			state = uint16(134)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(109):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('i') {
			state = uint16(153)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(110):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('i') {
			state = uint16(97)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(111):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('k') {
			state = uint16(66)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(112):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('l') {
			state = uint16(66)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(113):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('l') {
			state = uint16(114)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(114):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('l') {
			state = uint16(79)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(115):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('m') {
			state = uint16(140)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(116):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('m') {
			state = uint16(66)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(117):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('m') {
			state = uint16(80)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(118):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('m') {
			state = uint16(91)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(119):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('m') {
			state = uint16(99)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(120):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('n') {
			state = uint16(66)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(121):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('n') {
			state = uint16(154)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(122):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('n') {
			state = uint16(87)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(126)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(123):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('n') {
			state = uint16(68)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(124):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('n') {
			state = uint16(86)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(125):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('n') {
			state = uint16(156)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(126):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('n') {
			state = uint16(75)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(127):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('n') {
			state = uint16(161)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(128):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('o') {
			state = uint16(144)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(129):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('o') {
			state = uint16(167)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(130):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('o') {
			state = uint16(168)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(131):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('o') {
			state = uint16(136)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(132):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('o') {
			state = uint16(147)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(133):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('o') {
			state = uint16(145)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(134):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('o') {
			state = uint16(120)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(135):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('o') {
			state = uint16(141)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(136):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('p') {
			state = uint16(64)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(137):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('p') {
			state = uint16(78)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(138):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('p') {
			state = uint16(92)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(139):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('p') {
			state = uint16(132)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(90)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(140):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('p') {
			state = uint16(133)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(141):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(66)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(142):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(165)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(143):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(130)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(144):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(143)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(145):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(156)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(146):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(123)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(147):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(161)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(148):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(158)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(149):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(77)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(150):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(97)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(151):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('s') {
			state = uint16(66)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(152):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('s') {
			state = uint16(69)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(153):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('s') {
			state = uint16(101)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(154):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('s') {
			state = uint16(162)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(155):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('s') {
			state = uint16(151)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(156):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('t') {
			state = uint16(66)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(157):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('t') {
			state = uint16(164)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(158):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('t') {
			state = uint16(170)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(159):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('t') {
			state = uint16(109)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(160):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('t') {
			state = uint16(108)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(161):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('t') {
			state = uint16(151)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(162):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('t') {
			state = uint16(142)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(163):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('t') {
			state = uint16(135)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(164):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('u') {
			state = uint16(146)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(165):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('u') {
			state = uint16(84)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(166):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('v') {
			state = uint16(95)
			goto next_state
		}
		if lookahead == int32('x') {
			state = uint16(139)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(167):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('w') {
			state = uint16(68)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(168):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('w') {
			state = uint16(151)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(169):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('x') {
			state = uint16(97)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(170):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('y') {
			state = uint16(66)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(171):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(172):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COLON)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(173):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SLASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(174):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(175):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(224)
			goto next_state
		}
		return result
	case int32(176):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUND)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(177):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_TILDE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(178):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(179):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COMMA)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(180):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(181):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(182):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_code_block_language)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('`') {
			state = uint16(205)
			goto next_state
		}
		if int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(185)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(201)
			goto next_state
		}
		return result
	case int32(183):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_code_block_language)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('`') {
			state = uint16(205)
			goto next_state
		}
		if int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(184)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(202)
			goto next_state
		}
		return result
	case int32(184):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_code_block_language)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('`') {
			state = uint16(205)
			goto next_state
		}
		if int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(190)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(206)
			goto next_state
		}
		return result
	case int32(185):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_code_block_language)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('`') {
			state = uint16(205)
			goto next_state
		}
		if int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(192)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(207)
			goto next_state
		}
		return result
	case int32(186):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_code_block_language)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('`') {
			state = uint16(205)
			goto next_state
		}
		if int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(183)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(187):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_code_block_language)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('`') {
			state = uint16(205)
			goto next_state
		}
		if int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(187)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(200)
			goto next_state
		}
		return result
	case int32(188):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_code_block_language)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(188)
			goto next_state
		}
		return result
	case int32(189):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_code_block_language)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(193)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && (lookahead < int32('`') || int32('z') < lookahead) {
			state = uint16(53)
			goto next_state
		}
		return result
	case int32(190):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_code_block_language)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(186)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && (lookahead < int32('`') || int32('z') < lookahead) {
			state = uint16(40)
			goto next_state
		}
		return result
	case int32(191):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_code_block_language)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(194)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && (lookahead < int32('`') || int32('z') < lookahead) {
			state = uint16(55)
			goto next_state
		}
		return result
	case int32(192):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_code_block_language)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(195)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && (lookahead < int32('`') || int32('z') < lookahead) {
			state = uint16(54)
			goto next_state
		}
		return result
	case int32(193):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_code_block_language)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(191)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && (lookahead < int32('`') || int32('z') < lookahead) {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(194):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_code_block_language)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(189)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(205)
			goto next_state
		}
		return result
	case int32(195):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_code_block_language)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(182)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && (lookahead < int32('`') || int32('z') < lookahead) {
			state = uint16(39)
			goto next_state
		}
		return result
	case int32(196):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_code_block_line)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('*') {
			state = uint16(196)
			goto next_state
		}
		if lookahead == int32('`') {
			state = uint16(204)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			state = uint16(196)
			goto next_state
		}
		if int32(0x0b) <= lookahead && lookahead <= int32('\r') {
			state = uint16(198)
			goto next_state
		}
		if int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(187)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) {
			state = uint16(200)
			goto next_state
		}
		return result
	case int32(197):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_code_block_line)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('*') {
			state = uint16(197)
			goto next_state
		}
		if lookahead == int32('`') {
			state = uint16(204)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			state = uint16(197)
			goto next_state
		}
		if int32(0x0b) <= lookahead && lookahead <= int32('\r') {
			state = uint16(199)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) {
			state = uint16(200)
			goto next_state
		}
		return result
	case int32(198):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_code_block_line)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('`') {
			state = uint16(204)
			goto next_state
		}
		if lookahead == int32('\t') || int32(0x0b) <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(198)
			goto next_state
		}
		if int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(187)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) {
			state = uint16(200)
			goto next_state
		}
		return result
	case int32(199):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_code_block_line)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('`') {
			state = uint16(204)
			goto next_state
		}
		if lookahead == int32('\t') || int32(0x0b) <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(199)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) {
			state = uint16(200)
			goto next_state
		}
		return result
	case int32(200):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_code_block_line)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('`') {
			state = uint16(205)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(200)
			goto next_state
		}
		return result
	case int32(201):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_code_block_line)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('`') {
			state = uint16(205)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(207)
			goto next_state
		}
		return result
	case int32(202):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_code_block_line)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('`') {
			state = uint16(205)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(206)
			goto next_state
		}
		return result
	case int32(203):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_code_block_line)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('`') {
			state = uint16(205)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(202)
			goto next_state
		}
		return result
	case int32(204):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_code_block_line)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('`') {
			state = uint16(38)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(53)
			goto next_state
		}
		return result
	case int32(205):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_code_block_line)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('`') {
			state = uint16(53)
			goto next_state
		}
		return result
	case int32(206):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_code_block_line)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('`') {
			state = uint16(40)
			goto next_state
		}
		return result
	case int32(207):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_code_block_line)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('`') {
			state = uint16(54)
			goto next_state
		}
		return result
	case int32(208):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(209):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('$') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(209)
			goto next_state
		}
		return result
	case int32(210):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(211):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i1 = uint32(0)
		for {
			if !(uint64(i1) < libc.Uint64FromInt64(48)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token1[i1]) == lookahead {
				state = map_token1[i1+uint32(1)]
				goto next_state
			}
			goto _2
		_2:
			;
			i1 = i1 + uint32(2)
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(212)
			goto next_state
		}
		return result
	case int32(212):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(227)
			goto next_state
		}
		if lookahead == int32('_') {
			state = uint16(233)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(210)
			goto next_state
		}
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(230)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(212)
			goto next_state
		}
		return result
	case int32(213):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i2 = uint32(0)
		for {
			if !(uint64(i2) < libc.Uint64FromInt64(48)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token2[i2]) == lookahead {
				state = map_token2[i2+uint32(1)]
				goto next_state
			}
			goto _3
		_3:
			;
			i2 = i2 + uint32(2)
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(214)
			goto next_state
		}
		return result
	case int32(214):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(228)
			goto next_state
		}
		if lookahead == int32('_') {
			state = uint16(46)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(210)
			goto next_state
		}
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(43)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(214)
			goto next_state
		}
		return result
	case int32(215):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(235)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(210)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(215)
			goto next_state
		}
		return result
	case int32(216):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(231)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(210)
			goto next_state
		}
		if lookahead == int32('0') || lookahead == int32('1') {
			state = uint16(216)
			goto next_state
		}
		return result
	case int32(217):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(232)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(210)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('7') {
			state = uint16(217)
			goto next_state
		}
		return result
	case int32(218):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(237)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(210)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(218)
			goto next_state
		}
		return result
	case int32(219):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(47)
			goto next_state
		}
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(43)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(219)
			goto next_state
		}
		return result
	case int32(220):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(44)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(210)
			goto next_state
		}
		if lookahead == int32('0') || lookahead == int32('1') {
			state = uint16(220)
			goto next_state
		}
		return result
	case int32(221):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(45)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(210)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('7') {
			state = uint16(221)
			goto next_state
		}
		return result
	case int32(222):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(50)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(210)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(222)
			goto next_state
		}
		return result
	case int32(223):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(48)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(210)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(223)
			goto next_state
		}
		return result
	case int32(224):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(234)
			goto next_state
		}
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(230)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(224)
			goto next_state
		}
		return result
	case int32(225):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(236)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(225)
			goto next_state
		}
		return result
	case int32(226):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(49)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(226)
			goto next_state
		}
		return result
	case int32(227):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(230)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(224)
			goto next_state
		}
		return result
	case int32(228):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(43)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(219)
			goto next_state
		}
		return result
	case int32(229):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('*') {
			state = uint16(239)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(238)
			goto next_state
		}
		if lookahead == int32('`') {
			state = uint16(181)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(238)
			goto next_state
		}
		return result
	case int32(230):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('*') {
			state = uint16(239)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(238)
			goto next_state
		}
		if lookahead == int32('+') || lookahead == int32('-') {
			state = uint16(236)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(225)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(238)
			goto next_state
		}
		return result
	case int32(231):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('*') {
			state = uint16(239)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(238)
			goto next_state
		}
		if lookahead == int32('0') || lookahead == int32('1') {
			state = uint16(216)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(238)
			goto next_state
		}
		return result
	case int32(232):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('*') {
			state = uint16(239)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(238)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('7') {
			state = uint16(217)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(238)
			goto next_state
		}
		return result
	case int32(233):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('*') {
			state = uint16(239)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(238)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(212)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(238)
			goto next_state
		}
		return result
	case int32(234):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('*') {
			state = uint16(239)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(238)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(224)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(238)
			goto next_state
		}
		return result
	case int32(235):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('*') {
			state = uint16(239)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(238)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(215)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(238)
			goto next_state
		}
		return result
	case int32(236):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('*') {
			state = uint16(239)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(238)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(225)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(238)
			goto next_state
		}
		return result
	case int32(237):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('*') {
			state = uint16(239)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(238)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(218)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(238)
			goto next_state
		}
		return result
	case int32(238):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('*') {
			state = uint16(239)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(238)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(238)
			goto next_state
		}
		return result
	case int32(239):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('*') {
			state = uint16(239)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('/') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(37)
			goto next_state
		}
		return result
	case int32(240):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('/') {
			state = uint16(242)
			goto next_state
		}
		if lookahead == int32('`') {
			state = uint16(229)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('*') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(238)
			goto next_state
		}
		return result
	case int32(241):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('/') {
			state = uint16(242)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(224)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('*') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(238)
			goto next_state
		}
		return result
	case int32(242):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('/') {
			state = uint16(242)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('*') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(238)
			goto next_state
		}
		return result
	case int32(243):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SLASH2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(244):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STAR)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(245):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STAR)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('*') {
			state = uint16(245)
			goto next_state
		}
		return result
	default:
		return libc.BoolUint8(0 != 0)
	}
	return r
}

var map_token = [24]uint16_t{
	0:  uint16('a'),
	1:  uint16(81),
	2:  uint16('b'),
	3:  uint16(128),
	4:  uint16('c'),
	5:  uint16(70),
	6:  uint16('e'),
	7:  uint16(166),
	8:  uint16('f'),
	9:  uint16(105),
	10: uint16('i'),
	11: uint16(115),
	12: uint16('m'),
	13: uint16(93),
	14: uint16('n'),
	15: uint16(72),
	16: uint16('p'),
	17: uint16(74),
	18: uint16('r'),
	19: uint16(88),
	20: uint16('s'),
	21: uint16(71),
	22: uint16('t'),
	23: uint16(103),
}

var map_token1 = [24]uint16_t{
	0:  uint16('.'),
	1:  uint16(227),
	2:  uint16('0'),
	3:  uint16(215),
	4:  uint16('_'),
	5:  uint16(235),
	6:  uint16('n'),
	7:  uint16(210),
	8:  uint16('B'),
	9:  uint16(231),
	10: uint16('b'),
	11: uint16(231),
	12: uint16('E'),
	13: uint16(230),
	14: uint16('e'),
	15: uint16(230),
	16: uint16('O'),
	17: uint16(232),
	18: uint16('o'),
	19: uint16(232),
	20: uint16('X'),
	21: uint16(237),
	22: uint16('x'),
	23: uint16(237),
}

var map_token2 = [24]uint16_t{
	0:  uint16('.'),
	1:  uint16(228),
	2:  uint16('0'),
	3:  uint16(223),
	4:  uint16('_'),
	5:  uint16(48),
	6:  uint16('n'),
	7:  uint16(210),
	8:  uint16('B'),
	9:  uint16(44),
	10: uint16('b'),
	11: uint16(44),
	12: uint16('E'),
	13: uint16(43),
	14: uint16('e'),
	15: uint16(43),
	16: uint16('O'),
	17: uint16(45),
	18: uint16('o'),
	19: uint16(45),
	20: uint16('X'),
	21: uint16(50),
	22: uint16('x'),
	23: uint16(50),
}

var ts_lex_modes = [99]TSLexerMode{
	0: {
		Fexternal_lex_state: uint16(1),
	},
	1: {
		Flex_state: uint16(2),
	},
	2: {
		Flex_state: uint16(2),
	},
	3: {
		Flex_state: uint16(2),
	},
	4: {
		Flex_state: uint16(4),
	},
	5: {
		Flex_state: uint16(6),
	},
	6: {
		Flex_state: uint16(8),
	},
	7: {
		Flex_state: uint16(4),
	},
	8: {
		Flex_state: uint16(6),
	},
	9: {
		Flex_state: uint16(6),
	},
	10: {
		Flex_state: uint16(6),
	},
	11: {
		Flex_state: uint16(9),
	},
	12: {
		Flex_state: uint16(4),
	},
	13: {
		Flex_state: uint16(11),
	},
	14: {
		Flex_state: uint16(11),
	},
	15: {
		Flex_state: uint16(11),
	},
	16: {
		Flex_state: uint16(4),
	},
	17: {
		Flex_state: uint16(4),
	},
	18: {
		Flex_state: uint16(4),
	},
	19: {
		Flex_state: uint16(36),
	},
	20: {
		Flex_state: uint16(4),
	},
	21: {
		Flex_state: uint16(4),
	},
	22: {
		Flex_state: uint16(4),
	},
	23: {
		Flex_state: uint16(12),
	},
	24: {
		Flex_state: uint16(6),
	},
	25: {
		Flex_state: uint16(12),
	},
	26: {
		Flex_state: uint16(6),
	},
	27: {
		Flex_state: uint16(2),
	},
	28: {
		Flex_state: uint16(2),
	},
	29: {
		Flex_state: uint16(12),
	},
	30: {
		Flex_state: uint16(6),
	},
	31: {
		Flex_state: uint16(15),
	},
	32: {
		Flex_state: uint16(6),
	},
	33: {
		Flex_state: uint16(6),
	},
	34: {
		Flex_state: uint16(2),
	},
	35: {
		Flex_state: uint16(6),
	},
	36: {
		Flex_state: uint16(6),
	},
	37: {
		Flex_state: uint16(6),
	},
	38: {
		Flex_state: uint16(18),
	},
	39: {
		Flex_state: uint16(2),
	},
	40: {
		Flex_state: uint16(18),
	},
	41: {
		Flex_state: uint16(18),
	},
	42: {
		Flex_state: uint16(2),
	},
	43: {
		Flex_state: uint16(2),
	},
	44: {
		Flex_state: uint16(15),
	},
	45: {
		Flex_state: uint16(2),
	},
	46: {
		Flex_state: uint16(6),
	},
	47: {
		Flex_state: uint16(2),
	},
	48: {
		Flex_state: uint16(15),
	},
	49: {
		Flex_state: uint16(6),
	},
	50: {
		Flex_state: uint16(15),
	},
	51: {
		Flex_state: uint16(15),
	},
	52: {
		Flex_state: uint16(15),
	},
	53: {
		Flex_state: uint16(15),
	},
	54: {
		Flex_state: uint16(18),
	},
	55: {
		Flex_state: uint16(15),
	},
	56: {
		Flex_state: uint16(18),
	},
	57: {
		Flex_state: uint16(15),
	},
	58: {
		Flex_state: uint16(15),
	},
	59: {
		Flex_state: uint16(15),
	},
	60: {
		Flex_state: uint16(15),
	},
	61: {
		Flex_state: uint16(18),
	},
	62: {
		Flex_state: uint16(15),
	},
	63: {
		Flex_state: uint16(18),
	},
	64: {
		Flex_state: uint16(2),
	},
	65: {
		Flex_state: uint16(2),
	},
	66: {
		Flex_state: uint16(2),
	},
	67: {
		Flex_state: uint16(2),
	},
	68: {
		Flex_state: uint16(15),
	},
	69: {
		Flex_state:          uint16(25),
		Fexternal_lex_state: uint16(2),
	},
	70: {
		Flex_state:          uint16(25),
		Fexternal_lex_state: uint16(2),
	},
	71: {
		Flex_state: uint16(18),
	},
	72: {
		Flex_state: uint16(18),
	},
	73: {
		Flex_state:          uint16(31),
		Fexternal_lex_state: uint16(2),
	},
	74: {
		Flex_state:          uint16(31),
		Fexternal_lex_state: uint16(2),
	},
	75: {
		Flex_state: uint16(11),
	},
	76: {
		Flex_state:          uint16(31),
		Fexternal_lex_state: uint16(2),
	},
	77: {
		Flex_state: uint16(11),
	},
	78: {
		Flex_state:          uint16(31),
		Fexternal_lex_state: uint16(2),
	},
	79: {
		Flex_state:          uint16(31),
		Fexternal_lex_state: uint16(2),
	},
	80: {
		Flex_state:          uint16(31),
		Fexternal_lex_state: uint16(2),
	},
	81: {
		Flex_state:          uint16(31),
		Fexternal_lex_state: uint16(2),
	},
	82: {
		Flex_state: uint16(11),
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
		Fexternal_lex_state: uint16(1),
	},
	87: {},
	88: {},
	89: {},
	90: {
		Fexternal_lex_state: uint16(1),
	},
	91: {},
	92: {
		Flex_state: uint16(11),
	},
	93: {},
	94: {
		Flex_state: uint16(2),
	},
	95: {},
	96: {},
	97: {},
	98: {
		Flex_state: uint16(11),
	},
}

var ts_parse_table = [2][43]uint16_t{
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
		18: uint16(1),
		20: uint16(1),
		21: uint16(1),
		22: uint16(1),
		24: uint16(1),
	},
	1: {
		22: uint16(3),
		25: uint16(95),
		36: uint16(9),
	},
}

var ts_small_parse_table = [1340]uint16_t{
	0:    uint16(13),
	1:    uint16(5),
	2:    uint16(1),
	3:    uint16(anon_sym_LBRACE),
	4:    uint16(9),
	5:    uint16(1),
	6:    uint16(anon_sym_LBRACK),
	7:    uint16(11),
	8:    uint16(1),
	9:    uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	10:   uint16(13),
	11:   uint16(1),
	12:   uint16(sym_identifier),
	13:   uint16(15),
	14:   uint16(1),
	15:   uint16(sym_number),
	16:   uint16(17),
	17:   uint16(1),
	18:   uint16(sym__text),
	19:   uint16(19),
	20:   uint16(1),
	21:   uint16(anon_sym_SLASH2),
	22:   uint16(5),
	23:   uint16(1),
	24:   uint16(sym_code_block),
	25:   uint16(7),
	26:   uint16(1),
	27:   uint16(sym_expression),
	28:   uint16(30),
	29:   uint16(1),
	30:   uint16(sym_optional_identifier),
	31:   uint16(67),
	32:   uint16(1),
	33:   uint16(sym_description),
	34:   uint16(7),
	35:   uint16(3),
	36:   uint16(sym_tag_name_with_argument),
	37:   uint16(sym_tag_name_with_type),
	38:   uint16(sym_tag_name),
	39:   uint16(21),
	40:   uint16(4),
	41:   uint16(sym_qualified_expression),
	42:   uint16(sym_path_expression),
	43:   uint16(sym_member_expression),
	44:   uint16(sym_array_expression),
	45:   uint16(12),
	46:   uint16(9),
	47:   uint16(1),
	48:   uint16(anon_sym_LBRACK),
	49:   uint16(11),
	50:   uint16(1),
	51:   uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	52:   uint16(13),
	53:   uint16(1),
	54:   uint16(sym_identifier),
	55:   uint16(15),
	56:   uint16(1),
	57:   uint16(sym_number),
	58:   uint16(17),
	59:   uint16(1),
	60:   uint16(sym__text),
	61:   uint16(23),
	62:   uint16(1),
	63:   uint16(anon_sym_SLASH2),
	64:   uint16(4),
	65:   uint16(1),
	66:   uint16(sym_expression),
	67:   uint16(5),
	68:   uint16(1),
	69:   uint16(sym_code_block),
	70:   uint16(35),
	71:   uint16(1),
	72:   uint16(sym_optional_identifier),
	73:   uint16(64),
	74:   uint16(1),
	75:   uint16(sym_description),
	76:   uint16(21),
	77:   uint16(3),
	78:   uint16(sym_tag_name_with_argument),
	79:   uint16(sym_tag_name_with_type),
	80:   uint16(sym_tag_name),
	81:   uint16(21),
	82:   uint16(4),
	83:   uint16(sym_qualified_expression),
	84:   uint16(sym_path_expression),
	85:   uint16(sym_member_expression),
	86:   uint16(sym_array_expression),
	87:   uint16(7),
	88:   uint16(11),
	89:   uint16(1),
	90:   uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	91:   uint16(17),
	92:   uint16(1),
	93:   uint16(sym__text),
	94:   uint16(29),
	95:   uint16(1),
	96:   uint16(anon_sym_SLASH2),
	97:   uint16(5),
	98:   uint16(1),
	99:   uint16(sym_code_block),
	100:  uint16(65),
	101:  uint16(1),
	102:  uint16(sym_description),
	103:  uint16(25),
	104:  uint16(3),
	105:  uint16(sym_tag_name_with_argument),
	106:  uint16(sym_tag_name_with_type),
	107:  uint16(sym_tag_name),
	108:  uint16(27),
	109:  uint16(3),
	110:  uint16(anon_sym_DOT),
	111:  uint16(anon_sym_POUND),
	112:  uint16(anon_sym_TILDE),
	113:  uint16(7),
	114:  uint16(11),
	115:  uint16(1),
	116:  uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	117:  uint16(31),
	118:  uint16(1),
	119:  uint16(anon_sym_LBRACE),
	120:  uint16(33),
	121:  uint16(1),
	122:  uint16(sym__inline_tag_false_positive),
	123:  uint16(37),
	124:  uint16(1),
	125:  uint16(sym__text),
	126:  uint16(39),
	127:  uint16(1),
	128:  uint16(anon_sym_SLASH2),
	129:  uint16(35),
	130:  uint16(3),
	131:  uint16(sym_tag_name_with_argument),
	132:  uint16(sym_tag_name_with_type),
	133:  uint16(sym_tag_name),
	134:  uint16(8),
	135:  uint16(3),
	136:  uint16(sym_inline_tag),
	137:  uint16(sym_code_block),
	138:  uint16(aux_sym_description_repeat1),
	139:  uint16(4),
	140:  uint16(43),
	141:  uint16(1),
	142:  uint16(anon_sym_COLON),
	143:  uint16(45),
	144:  uint16(1),
	145:  uint16(anon_sym_SLASH),
	146:  uint16(47),
	147:  uint16(4),
	148:  uint16(anon_sym_DOT),
	149:  uint16(anon_sym_POUND),
	150:  uint16(anon_sym_TILDE),
	151:  uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	152:  uint16(41),
	153:  uint16(5),
	154:  uint16(sym_tag_name_with_argument),
	155:  uint16(sym_tag_name_with_type),
	156:  uint16(sym_tag_name),
	157:  uint16(sym__text),
	158:  uint16(anon_sym_SLASH2),
	159:  uint16(7),
	160:  uint16(11),
	161:  uint16(1),
	162:  uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	163:  uint16(17),
	164:  uint16(1),
	165:  uint16(sym__text),
	166:  uint16(51),
	167:  uint16(1),
	168:  uint16(anon_sym_SLASH2),
	169:  uint16(5),
	170:  uint16(1),
	171:  uint16(sym_code_block),
	172:  uint16(66),
	173:  uint16(1),
	174:  uint16(sym_description),
	175:  uint16(27),
	176:  uint16(3),
	177:  uint16(anon_sym_DOT),
	178:  uint16(anon_sym_POUND),
	179:  uint16(anon_sym_TILDE),
	180:  uint16(49),
	181:  uint16(3),
	182:  uint16(sym_tag_name_with_argument),
	183:  uint16(sym_tag_name_with_type),
	184:  uint16(sym_tag_name),
	185:  uint16(7),
	186:  uint16(11),
	187:  uint16(1),
	188:  uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	189:  uint16(31),
	190:  uint16(1),
	191:  uint16(anon_sym_LBRACE),
	192:  uint16(53),
	193:  uint16(1),
	194:  uint16(sym__inline_tag_false_positive),
	195:  uint16(57),
	196:  uint16(1),
	197:  uint16(sym__text),
	198:  uint16(59),
	199:  uint16(1),
	200:  uint16(anon_sym_SLASH2),
	201:  uint16(55),
	202:  uint16(3),
	203:  uint16(sym_tag_name_with_argument),
	204:  uint16(sym_tag_name_with_type),
	205:  uint16(sym_tag_name),
	206:  uint16(10),
	207:  uint16(3),
	208:  uint16(sym_inline_tag),
	209:  uint16(sym_code_block),
	210:  uint16(aux_sym_description_repeat1),
	211:  uint16(10),
	212:  uint16(11),
	213:  uint16(1),
	214:  uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	215:  uint16(17),
	216:  uint16(1),
	217:  uint16(sym__text),
	218:  uint16(61),
	219:  uint16(1),
	220:  uint16(sym_tag_name_with_argument),
	221:  uint16(63),
	222:  uint16(1),
	223:  uint16(sym_tag_name_with_type),
	224:  uint16(65),
	225:  uint16(1),
	226:  uint16(sym_tag_name),
	227:  uint16(67),
	228:  uint16(1),
	229:  uint16(anon_sym_SLASH2),
	230:  uint16(5),
	231:  uint16(1),
	232:  uint16(sym_code_block),
	233:  uint16(43),
	234:  uint16(1),
	235:  uint16(sym_description),
	236:  uint16(88),
	237:  uint16(1),
	238:  uint16(sym__end),
	239:  uint16(45),
	240:  uint16(2),
	241:  uint16(sym_tag),
	242:  uint16(aux_sym_document_repeat1),
	243:  uint16(7),
	244:  uint16(69),
	245:  uint16(1),
	246:  uint16(anon_sym_LBRACE),
	247:  uint16(72),
	248:  uint16(1),
	249:  uint16(sym__inline_tag_false_positive),
	250:  uint16(77),
	251:  uint16(1),
	252:  uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	253:  uint16(80),
	254:  uint16(1),
	255:  uint16(sym__text),
	256:  uint16(83),
	257:  uint16(1),
	258:  uint16(anon_sym_SLASH2),
	259:  uint16(75),
	260:  uint16(3),
	261:  uint16(sym_tag_name_with_argument),
	262:  uint16(sym_tag_name_with_type),
	263:  uint16(sym_tag_name),
	264:  uint16(10),
	265:  uint16(3),
	266:  uint16(sym_inline_tag),
	267:  uint16(sym_code_block),
	268:  uint16(aux_sym_description_repeat1),
	269:  uint16(3),
	270:  uint16(43),
	271:  uint16(1),
	272:  uint16(anon_sym_COLON),
	273:  uint16(85),
	274:  uint16(4),
	275:  uint16(sym_tag_name_with_argument),
	276:  uint16(sym_tag_name_with_type),
	277:  uint16(sym_tag_name),
	278:  uint16(sym__text),
	279:  uint16(87),
	280:  uint16(5),
	281:  uint16(anon_sym_DOT),
	282:  uint16(anon_sym_POUND),
	283:  uint16(anon_sym_TILDE),
	284:  uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	285:  uint16(anon_sym_SLASH2),
	286:  uint16(2),
	287:  uint16(89),
	288:  uint16(4),
	289:  uint16(sym_tag_name_with_argument),
	290:  uint16(sym_tag_name_with_type),
	291:  uint16(sym_tag_name),
	292:  uint16(sym__text),
	293:  uint16(91),
	294:  uint16(5),
	295:  uint16(anon_sym_DOT),
	296:  uint16(anon_sym_POUND),
	297:  uint16(anon_sym_TILDE),
	298:  uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	299:  uint16(anon_sym_SLASH2),
	300:  uint16(6),
	301:  uint16(93),
	302:  uint16(1),
	303:  uint16(anon_sym_LBRACK),
	304:  uint16(95),
	305:  uint16(1),
	306:  uint16(anon_sym_RBRACK),
	307:  uint16(97),
	308:  uint16(1),
	309:  uint16(sym_identifier),
	310:  uint16(99),
	311:  uint16(1),
	312:  uint16(sym_number),
	313:  uint16(48),
	314:  uint16(1),
	315:  uint16(sym_expression),
	316:  uint16(53),
	317:  uint16(4),
	318:  uint16(sym_qualified_expression),
	319:  uint16(sym_path_expression),
	320:  uint16(sym_member_expression),
	321:  uint16(sym_array_expression),
	322:  uint16(6),
	323:  uint16(93),
	324:  uint16(1),
	325:  uint16(anon_sym_LBRACK),
	326:  uint16(97),
	327:  uint16(1),
	328:  uint16(sym_identifier),
	329:  uint16(99),
	330:  uint16(1),
	331:  uint16(sym_number),
	332:  uint16(101),
	333:  uint16(1),
	334:  uint16(anon_sym_RBRACK),
	335:  uint16(50),
	336:  uint16(1),
	337:  uint16(sym_expression),
	338:  uint16(53),
	339:  uint16(4),
	340:  uint16(sym_qualified_expression),
	341:  uint16(sym_path_expression),
	342:  uint16(sym_member_expression),
	343:  uint16(sym_array_expression),
	344:  uint16(6),
	345:  uint16(93),
	346:  uint16(1),
	347:  uint16(anon_sym_LBRACK),
	348:  uint16(95),
	349:  uint16(1),
	350:  uint16(anon_sym_RBRACK),
	351:  uint16(99),
	352:  uint16(1),
	353:  uint16(sym_number),
	354:  uint16(103),
	355:  uint16(1),
	356:  uint16(sym_identifier),
	357:  uint16(48),
	358:  uint16(1),
	359:  uint16(sym_expression),
	360:  uint16(53),
	361:  uint16(4),
	362:  uint16(sym_qualified_expression),
	363:  uint16(sym_path_expression),
	364:  uint16(sym_member_expression),
	365:  uint16(sym_array_expression),
	366:  uint16(2),
	367:  uint16(105),
	368:  uint16(4),
	369:  uint16(sym_tag_name_with_argument),
	370:  uint16(sym_tag_name_with_type),
	371:  uint16(sym_tag_name),
	372:  uint16(sym__text),
	373:  uint16(107),
	374:  uint16(5),
	375:  uint16(anon_sym_DOT),
	376:  uint16(anon_sym_POUND),
	377:  uint16(anon_sym_TILDE),
	378:  uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	379:  uint16(anon_sym_SLASH2),
	380:  uint16(2),
	381:  uint16(109),
	382:  uint16(4),
	383:  uint16(sym_tag_name_with_argument),
	384:  uint16(sym_tag_name_with_type),
	385:  uint16(sym_tag_name),
	386:  uint16(sym__text),
	387:  uint16(111),
	388:  uint16(5),
	389:  uint16(anon_sym_DOT),
	390:  uint16(anon_sym_POUND),
	391:  uint16(anon_sym_TILDE),
	392:  uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	393:  uint16(anon_sym_SLASH2),
	394:  uint16(2),
	395:  uint16(113),
	396:  uint16(4),
	397:  uint16(sym_tag_name_with_argument),
	398:  uint16(sym_tag_name_with_type),
	399:  uint16(sym_tag_name),
	400:  uint16(sym__text),
	401:  uint16(115),
	402:  uint16(5),
	403:  uint16(anon_sym_DOT),
	404:  uint16(anon_sym_POUND),
	405:  uint16(anon_sym_TILDE),
	406:  uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	407:  uint16(anon_sym_SLASH2),
	408:  uint16(7),
	409:  uint16(11),
	410:  uint16(1),
	411:  uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	412:  uint16(17),
	413:  uint16(1),
	414:  uint16(sym__text),
	415:  uint16(19),
	416:  uint16(1),
	417:  uint16(anon_sym_SLASH2),
	418:  uint16(117),
	419:  uint16(1),
	420:  uint16(anon_sym_LBRACE),
	421:  uint16(5),
	422:  uint16(1),
	423:  uint16(sym_code_block),
	424:  uint16(67),
	425:  uint16(1),
	426:  uint16(sym_description),
	427:  uint16(7),
	428:  uint16(3),
	429:  uint16(sym_tag_name_with_argument),
	430:  uint16(sym_tag_name_with_type),
	431:  uint16(sym_tag_name),
	432:  uint16(2),
	433:  uint16(85),
	434:  uint16(4),
	435:  uint16(sym_tag_name_with_argument),
	436:  uint16(sym_tag_name_with_type),
	437:  uint16(sym_tag_name),
	438:  uint16(sym__text),
	439:  uint16(87),
	440:  uint16(5),
	441:  uint16(anon_sym_DOT),
	442:  uint16(anon_sym_POUND),
	443:  uint16(anon_sym_TILDE),
	444:  uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	445:  uint16(anon_sym_SLASH2),
	446:  uint16(2),
	447:  uint16(41),
	448:  uint16(4),
	449:  uint16(sym_tag_name_with_argument),
	450:  uint16(sym_tag_name_with_type),
	451:  uint16(sym_tag_name),
	452:  uint16(sym__text),
	453:  uint16(47),
	454:  uint16(5),
	455:  uint16(anon_sym_DOT),
	456:  uint16(anon_sym_POUND),
	457:  uint16(anon_sym_TILDE),
	458:  uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	459:  uint16(anon_sym_SLASH2),
	460:  uint16(2),
	461:  uint16(119),
	462:  uint16(4),
	463:  uint16(sym_tag_name_with_argument),
	464:  uint16(sym_tag_name_with_type),
	465:  uint16(sym_tag_name),
	466:  uint16(sym__text),
	467:  uint16(121),
	468:  uint16(5),
	469:  uint16(anon_sym_DOT),
	470:  uint16(anon_sym_POUND),
	471:  uint16(anon_sym_TILDE),
	472:  uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	473:  uint16(anon_sym_SLASH2),
	474:  uint16(3),
	475:  uint16(125),
	476:  uint16(1),
	477:  uint16(anon_sym_STAR),
	478:  uint16(29),
	479:  uint16(1),
	480:  uint16(aux_sym__begin_repeat1),
	481:  uint16(123),
	482:  uint16(6),
	483:  uint16(sym_tag_name_with_argument),
	484:  uint16(sym_tag_name_with_type),
	485:  uint16(sym_tag_name),
	486:  uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	487:  uint16(sym__text),
	488:  uint16(anon_sym_SLASH2),
	489:  uint16(6),
	490:  uint16(11),
	491:  uint16(1),
	492:  uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	493:  uint16(17),
	494:  uint16(1),
	495:  uint16(sym__text),
	496:  uint16(19),
	497:  uint16(1),
	498:  uint16(anon_sym_SLASH2),
	499:  uint16(5),
	500:  uint16(1),
	501:  uint16(sym_code_block),
	502:  uint16(67),
	503:  uint16(1),
	504:  uint16(sym_description),
	505:  uint16(7),
	506:  uint16(3),
	507:  uint16(sym_tag_name_with_argument),
	508:  uint16(sym_tag_name_with_type),
	509:  uint16(sym_tag_name),
	510:  uint16(3),
	511:  uint16(129),
	512:  uint16(1),
	513:  uint16(anon_sym_STAR),
	514:  uint16(25),
	515:  uint16(1),
	516:  uint16(aux_sym__begin_repeat1),
	517:  uint16(127),
	518:  uint16(6),
	519:  uint16(sym_tag_name_with_argument),
	520:  uint16(sym_tag_name_with_type),
	521:  uint16(sym_tag_name),
	522:  uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	523:  uint16(sym__text),
	524:  uint16(anon_sym_SLASH2),
	525:  uint16(2),
	526:  uint16(134),
	527:  uint16(3),
	528:  uint16(sym__inline_tag_false_positive),
	529:  uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	530:  uint16(anon_sym_SLASH2),
	531:  uint16(132),
	532:  uint16(5),
	533:  uint16(anon_sym_LBRACE),
	534:  uint16(sym_tag_name_with_argument),
	535:  uint16(sym_tag_name_with_type),
	536:  uint16(sym_tag_name),
	537:  uint16(sym__text),
	538:  uint16(5),
	539:  uint16(93),
	540:  uint16(1),
	541:  uint16(anon_sym_LBRACK),
	542:  uint16(97),
	543:  uint16(1),
	544:  uint16(sym_identifier),
	545:  uint16(99),
	546:  uint16(1),
	547:  uint16(sym_number),
	548:  uint16(68),
	549:  uint16(1),
	550:  uint16(sym_expression),
	551:  uint16(53),
	552:  uint16(4),
	553:  uint16(sym_qualified_expression),
	554:  uint16(sym_path_expression),
	555:  uint16(sym_member_expression),
	556:  uint16(sym_array_expression),
	557:  uint16(5),
	558:  uint16(93),
	559:  uint16(1),
	560:  uint16(anon_sym_LBRACK),
	561:  uint16(97),
	562:  uint16(1),
	563:  uint16(sym_identifier),
	564:  uint16(99),
	565:  uint16(1),
	566:  uint16(sym_number),
	567:  uint16(55),
	568:  uint16(1),
	569:  uint16(sym_expression),
	570:  uint16(53),
	571:  uint16(4),
	572:  uint16(sym_qualified_expression),
	573:  uint16(sym_path_expression),
	574:  uint16(sym_member_expression),
	575:  uint16(sym_array_expression),
	576:  uint16(3),
	577:  uint16(138),
	578:  uint16(1),
	579:  uint16(anon_sym_STAR),
	580:  uint16(25),
	581:  uint16(1),
	582:  uint16(aux_sym__begin_repeat1),
	583:  uint16(136),
	584:  uint16(6),
	585:  uint16(sym_tag_name_with_argument),
	586:  uint16(sym_tag_name_with_type),
	587:  uint16(sym_tag_name),
	588:  uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	589:  uint16(sym__text),
	590:  uint16(anon_sym_SLASH2),
	591:  uint16(6),
	592:  uint16(11),
	593:  uint16(1),
	594:  uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	595:  uint16(17),
	596:  uint16(1),
	597:  uint16(sym__text),
	598:  uint16(51),
	599:  uint16(1),
	600:  uint16(anon_sym_SLASH2),
	601:  uint16(5),
	602:  uint16(1),
	603:  uint16(sym_code_block),
	604:  uint16(66),
	605:  uint16(1),
	606:  uint16(sym_description),
	607:  uint16(49),
	608:  uint16(3),
	609:  uint16(sym_tag_name_with_argument),
	610:  uint16(sym_tag_name_with_type),
	611:  uint16(sym_tag_name),
	612:  uint16(5),
	613:  uint16(140),
	614:  uint16(1),
	615:  uint16(anon_sym_COLON),
	616:  uint16(142),
	617:  uint16(1),
	618:  uint16(anon_sym_SLASH),
	619:  uint16(144),
	620:  uint16(1),
	621:  uint16(anon_sym_RBRACK),
	622:  uint16(146),
	623:  uint16(1),
	624:  uint16(anon_sym_EQ),
	625:  uint16(47),
	626:  uint16(4),
	627:  uint16(anon_sym_DOT),
	628:  uint16(anon_sym_POUND),
	629:  uint16(anon_sym_TILDE),
	630:  uint16(anon_sym_COMMA),
	631:  uint16(6),
	632:  uint16(11),
	633:  uint16(1),
	634:  uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	635:  uint16(17),
	636:  uint16(1),
	637:  uint16(sym__text),
	638:  uint16(23),
	639:  uint16(1),
	640:  uint16(anon_sym_SLASH2),
	641:  uint16(5),
	642:  uint16(1),
	643:  uint16(sym_code_block),
	644:  uint16(64),
	645:  uint16(1),
	646:  uint16(sym_description),
	647:  uint16(21),
	648:  uint16(3),
	649:  uint16(sym_tag_name_with_argument),
	650:  uint16(sym_tag_name_with_type),
	651:  uint16(sym_tag_name),
	652:  uint16(2),
	653:  uint16(150),
	654:  uint16(3),
	655:  uint16(sym__inline_tag_false_positive),
	656:  uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	657:  uint16(anon_sym_SLASH2),
	658:  uint16(148),
	659:  uint16(5),
	660:  uint16(anon_sym_LBRACE),
	661:  uint16(sym_tag_name_with_argument),
	662:  uint16(sym_tag_name_with_type),
	663:  uint16(sym_tag_name),
	664:  uint16(sym__text),
	665:  uint16(5),
	666:  uint16(13),
	667:  uint16(1),
	668:  uint16(sym_identifier),
	669:  uint16(15),
	670:  uint16(1),
	671:  uint16(sym_number),
	672:  uint16(152),
	673:  uint16(1),
	674:  uint16(anon_sym_LBRACK),
	675:  uint16(17),
	676:  uint16(1),
	677:  uint16(sym_expression),
	678:  uint16(21),
	679:  uint16(4),
	680:  uint16(sym_qualified_expression),
	681:  uint16(sym_path_expression),
	682:  uint16(sym_member_expression),
	683:  uint16(sym_array_expression),
	684:  uint16(6),
	685:  uint16(11),
	686:  uint16(1),
	687:  uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	688:  uint16(17),
	689:  uint16(1),
	690:  uint16(sym__text),
	691:  uint16(29),
	692:  uint16(1),
	693:  uint16(anon_sym_SLASH2),
	694:  uint16(5),
	695:  uint16(1),
	696:  uint16(sym_code_block),
	697:  uint16(65),
	698:  uint16(1),
	699:  uint16(sym_description),
	700:  uint16(25),
	701:  uint16(3),
	702:  uint16(sym_tag_name_with_argument),
	703:  uint16(sym_tag_name_with_type),
	704:  uint16(sym_tag_name),
	705:  uint16(2),
	706:  uint16(156),
	707:  uint16(3),
	708:  uint16(sym__inline_tag_false_positive),
	709:  uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	710:  uint16(anon_sym_SLASH2),
	711:  uint16(154),
	712:  uint16(5),
	713:  uint16(anon_sym_LBRACE),
	714:  uint16(sym_tag_name_with_argument),
	715:  uint16(sym_tag_name_with_type),
	716:  uint16(sym_tag_name),
	717:  uint16(sym__text),
	718:  uint16(2),
	719:  uint16(160),
	720:  uint16(3),
	721:  uint16(sym__inline_tag_false_positive),
	722:  uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	723:  uint16(anon_sym_SLASH2),
	724:  uint16(158),
	725:  uint16(5),
	726:  uint16(anon_sym_LBRACE),
	727:  uint16(sym_tag_name_with_argument),
	728:  uint16(sym_tag_name_with_type),
	729:  uint16(sym_tag_name),
	730:  uint16(sym__text),
	731:  uint16(6),
	732:  uint16(39),
	733:  uint16(1),
	734:  uint16(anon_sym_RBRACE),
	735:  uint16(162),
	736:  uint16(1),
	737:  uint16(anon_sym_LBRACE),
	738:  uint16(164),
	739:  uint16(1),
	740:  uint16(sym__inline_tag_false_positive),
	741:  uint16(166),
	742:  uint16(1),
	743:  uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	744:  uint16(168),
	745:  uint16(1),
	746:  uint16(sym__text),
	747:  uint16(40),
	748:  uint16(3),
	749:  uint16(sym_inline_tag),
	750:  uint16(sym_code_block),
	751:  uint16(aux_sym_description_repeat1),
	752:  uint16(5),
	753:  uint16(93),
	754:  uint16(1),
	755:  uint16(anon_sym_LBRACK),
	756:  uint16(97),
	757:  uint16(1),
	758:  uint16(sym_identifier),
	759:  uint16(99),
	760:  uint16(1),
	761:  uint16(sym_number),
	762:  uint16(57),
	763:  uint16(1),
	764:  uint16(sym_expression),
	765:  uint16(53),
	766:  uint16(4),
	767:  uint16(sym_qualified_expression),
	768:  uint16(sym_path_expression),
	769:  uint16(sym_member_expression),
	770:  uint16(sym_array_expression),
	771:  uint16(6),
	772:  uint16(59),
	773:  uint16(1),
	774:  uint16(anon_sym_RBRACE),
	775:  uint16(162),
	776:  uint16(1),
	777:  uint16(anon_sym_LBRACE),
	778:  uint16(166),
	779:  uint16(1),
	780:  uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	781:  uint16(170),
	782:  uint16(1),
	783:  uint16(sym__inline_tag_false_positive),
	784:  uint16(172),
	785:  uint16(1),
	786:  uint16(sym__text),
	787:  uint16(41),
	788:  uint16(3),
	789:  uint16(sym_inline_tag),
	790:  uint16(sym_code_block),
	791:  uint16(aux_sym_description_repeat1),
	792:  uint16(6),
	793:  uint16(83),
	794:  uint16(1),
	795:  uint16(anon_sym_RBRACE),
	796:  uint16(174),
	797:  uint16(1),
	798:  uint16(anon_sym_LBRACE),
	799:  uint16(177),
	800:  uint16(1),
	801:  uint16(sym__inline_tag_false_positive),
	802:  uint16(180),
	803:  uint16(1),
	804:  uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	805:  uint16(183),
	806:  uint16(1),
	807:  uint16(sym__text),
	808:  uint16(41),
	809:  uint16(3),
	810:  uint16(sym_inline_tag),
	811:  uint16(sym_code_block),
	812:  uint16(aux_sym_description_repeat1),
	813:  uint16(6),
	814:  uint16(61),
	815:  uint16(1),
	816:  uint16(sym_tag_name_with_argument),
	817:  uint16(63),
	818:  uint16(1),
	819:  uint16(sym_tag_name_with_type),
	820:  uint16(65),
	821:  uint16(1),
	822:  uint16(sym_tag_name),
	823:  uint16(186),
	824:  uint16(1),
	825:  uint16(anon_sym_SLASH2),
	826:  uint16(93),
	827:  uint16(1),
	828:  uint16(sym__end),
	829:  uint16(47),
	830:  uint16(2),
	831:  uint16(sym_tag),
	832:  uint16(aux_sym_document_repeat1),
	833:  uint16(6),
	834:  uint16(61),
	835:  uint16(1),
	836:  uint16(sym_tag_name_with_argument),
	837:  uint16(63),
	838:  uint16(1),
	839:  uint16(sym_tag_name_with_type),
	840:  uint16(65),
	841:  uint16(1),
	842:  uint16(sym_tag_name),
	843:  uint16(188),
	844:  uint16(1),
	845:  uint16(anon_sym_SLASH2),
	846:  uint16(91),
	847:  uint16(1),
	848:  uint16(sym__end),
	849:  uint16(42),
	850:  uint16(2),
	851:  uint16(sym_tag),
	852:  uint16(aux_sym_document_repeat1),
	853:  uint16(3),
	854:  uint16(140),
	855:  uint16(1),
	856:  uint16(anon_sym_COLON),
	857:  uint16(142),
	858:  uint16(1),
	859:  uint16(anon_sym_SLASH),
	860:  uint16(47),
	861:  uint16(5),
	862:  uint16(anon_sym_DOT),
	863:  uint16(anon_sym_POUND),
	864:  uint16(anon_sym_TILDE),
	865:  uint16(anon_sym_COMMA),
	866:  uint16(anon_sym_RBRACK),
	867:  uint16(6),
	868:  uint16(61),
	869:  uint16(1),
	870:  uint16(sym_tag_name_with_argument),
	871:  uint16(63),
	872:  uint16(1),
	873:  uint16(sym_tag_name_with_type),
	874:  uint16(65),
	875:  uint16(1),
	876:  uint16(sym_tag_name),
	877:  uint16(188),
	878:  uint16(1),
	879:  uint16(anon_sym_SLASH2),
	880:  uint16(91),
	881:  uint16(1),
	882:  uint16(sym__end),
	883:  uint16(47),
	884:  uint16(2),
	885:  uint16(sym_tag),
	886:  uint16(aux_sym_document_repeat1),
	887:  uint16(2),
	888:  uint16(192),
	889:  uint16(2),
	890:  uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	891:  uint16(anon_sym_SLASH2),
	892:  uint16(190),
	893:  uint16(4),
	894:  uint16(sym_tag_name_with_argument),
	895:  uint16(sym_tag_name_with_type),
	896:  uint16(sym_tag_name),
	897:  uint16(sym__text),
	898:  uint16(5),
	899:  uint16(194),
	900:  uint16(1),
	901:  uint16(sym_tag_name_with_argument),
	902:  uint16(197),
	903:  uint16(1),
	904:  uint16(sym_tag_name_with_type),
	905:  uint16(200),
	906:  uint16(1),
	907:  uint16(sym_tag_name),
	908:  uint16(203),
	909:  uint16(1),
	910:  uint16(anon_sym_SLASH2),
	911:  uint16(47),
	912:  uint16(2),
	913:  uint16(sym_tag),
	914:  uint16(aux_sym_document_repeat1),
	915:  uint16(4),
	916:  uint16(207),
	917:  uint16(1),
	918:  uint16(anon_sym_COMMA),
	919:  uint16(209),
	920:  uint16(1),
	921:  uint16(anon_sym_RBRACK),
	922:  uint16(75),
	923:  uint16(1),
	924:  uint16(aux_sym_array_expression_repeat1),
	925:  uint16(205),
	926:  uint16(3),
	927:  uint16(anon_sym_DOT),
	928:  uint16(anon_sym_POUND),
	929:  uint16(anon_sym_TILDE),
	930:  uint16(2),
	931:  uint16(213),
	932:  uint16(2),
	933:  uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	934:  uint16(anon_sym_SLASH2),
	935:  uint16(211),
	936:  uint16(4),
	937:  uint16(sym_tag_name_with_argument),
	938:  uint16(sym_tag_name_with_type),
	939:  uint16(sym_tag_name),
	940:  uint16(sym__text),
	941:  uint16(4),
	942:  uint16(207),
	943:  uint16(1),
	944:  uint16(anon_sym_COMMA),
	945:  uint16(215),
	946:  uint16(1),
	947:  uint16(anon_sym_RBRACK),
	948:  uint16(82),
	949:  uint16(1),
	950:  uint16(aux_sym_array_expression_repeat1),
	951:  uint16(205),
	952:  uint16(3),
	953:  uint16(anon_sym_DOT),
	954:  uint16(anon_sym_POUND),
	955:  uint16(anon_sym_TILDE),
	956:  uint16(2),
	957:  uint16(140),
	958:  uint16(1),
	959:  uint16(anon_sym_COLON),
	960:  uint16(87),
	961:  uint16(5),
	962:  uint16(anon_sym_DOT),
	963:  uint16(anon_sym_POUND),
	964:  uint16(anon_sym_TILDE),
	965:  uint16(anon_sym_COMMA),
	966:  uint16(anon_sym_RBRACK),
	967:  uint16(1),
	968:  uint16(91),
	969:  uint16(5),
	970:  uint16(anon_sym_DOT),
	971:  uint16(anon_sym_POUND),
	972:  uint16(anon_sym_TILDE),
	973:  uint16(anon_sym_COMMA),
	974:  uint16(anon_sym_RBRACK),
	975:  uint16(1),
	976:  uint16(47),
	977:  uint16(5),
	978:  uint16(anon_sym_DOT),
	979:  uint16(anon_sym_POUND),
	980:  uint16(anon_sym_TILDE),
	981:  uint16(anon_sym_COMMA),
	982:  uint16(anon_sym_RBRACK),
	983:  uint16(2),
	984:  uint16(132),
	985:  uint16(2),
	986:  uint16(anon_sym_LBRACE),
	987:  uint16(sym__text),
	988:  uint16(134),
	989:  uint16(3),
	990:  uint16(anon_sym_RBRACE),
	991:  uint16(sym__inline_tag_false_positive),
	992:  uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	993:  uint16(2),
	994:  uint16(217),
	995:  uint16(2),
	996:  uint16(anon_sym_COMMA),
	997:  uint16(anon_sym_RBRACK),
	998:  uint16(205),
	999:  uint16(3),
	1000: uint16(anon_sym_DOT),
	1001: uint16(anon_sym_POUND),
	1002: uint16(anon_sym_TILDE),
	1003: uint16(2),
	1004: uint16(154),
	1005: uint16(2),
	1006: uint16(anon_sym_LBRACE),
	1007: uint16(sym__text),
	1008: uint16(156),
	1009: uint16(3),
	1010: uint16(anon_sym_RBRACE),
	1011: uint16(sym__inline_tag_false_positive),
	1012: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	1013: uint16(1),
	1014: uint16(111),
	1015: uint16(5),
	1016: uint16(anon_sym_DOT),
	1017: uint16(anon_sym_POUND),
	1018: uint16(anon_sym_TILDE),
	1019: uint16(anon_sym_COMMA),
	1020: uint16(anon_sym_RBRACK),
	1021: uint16(1),
	1022: uint16(87),
	1023: uint16(5),
	1024: uint16(anon_sym_DOT),
	1025: uint16(anon_sym_POUND),
	1026: uint16(anon_sym_TILDE),
	1027: uint16(anon_sym_COMMA),
	1028: uint16(anon_sym_RBRACK),
	1029: uint16(1),
	1030: uint16(107),
	1031: uint16(5),
	1032: uint16(anon_sym_DOT),
	1033: uint16(anon_sym_POUND),
	1034: uint16(anon_sym_TILDE),
	1035: uint16(anon_sym_COMMA),
	1036: uint16(anon_sym_RBRACK),
	1037: uint16(1),
	1038: uint16(115),
	1039: uint16(5),
	1040: uint16(anon_sym_DOT),
	1041: uint16(anon_sym_POUND),
	1042: uint16(anon_sym_TILDE),
	1043: uint16(anon_sym_COMMA),
	1044: uint16(anon_sym_RBRACK),
	1045: uint16(2),
	1046: uint16(148),
	1047: uint16(2),
	1048: uint16(anon_sym_LBRACE),
	1049: uint16(sym__text),
	1050: uint16(150),
	1051: uint16(3),
	1052: uint16(anon_sym_RBRACE),
	1053: uint16(sym__inline_tag_false_positive),
	1054: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	1055: uint16(1),
	1056: uint16(121),
	1057: uint16(5),
	1058: uint16(anon_sym_DOT),
	1059: uint16(anon_sym_POUND),
	1060: uint16(anon_sym_TILDE),
	1061: uint16(anon_sym_COMMA),
	1062: uint16(anon_sym_RBRACK),
	1063: uint16(2),
	1064: uint16(158),
	1065: uint16(2),
	1066: uint16(anon_sym_LBRACE),
	1067: uint16(sym__text),
	1068: uint16(160),
	1069: uint16(3),
	1070: uint16(anon_sym_RBRACE),
	1071: uint16(sym__inline_tag_false_positive),
	1072: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	1073: uint16(2),
	1074: uint16(29),
	1075: uint16(1),
	1076: uint16(anon_sym_SLASH2),
	1077: uint16(25),
	1078: uint16(3),
	1079: uint16(sym_tag_name_with_argument),
	1080: uint16(sym_tag_name_with_type),
	1081: uint16(sym_tag_name),
	1082: uint16(2),
	1083: uint16(221),
	1084: uint16(1),
	1085: uint16(anon_sym_SLASH2),
	1086: uint16(219),
	1087: uint16(3),
	1088: uint16(sym_tag_name_with_argument),
	1089: uint16(sym_tag_name_with_type),
	1090: uint16(sym_tag_name),
	1091: uint16(2),
	1092: uint16(225),
	1093: uint16(1),
	1094: uint16(anon_sym_SLASH2),
	1095: uint16(223),
	1096: uint16(3),
	1097: uint16(sym_tag_name_with_argument),
	1098: uint16(sym_tag_name_with_type),
	1099: uint16(sym_tag_name),
	1100: uint16(2),
	1101: uint16(51),
	1102: uint16(1),
	1103: uint16(anon_sym_SLASH2),
	1104: uint16(49),
	1105: uint16(3),
	1106: uint16(sym_tag_name_with_argument),
	1107: uint16(sym_tag_name_with_type),
	1108: uint16(sym_tag_name),
	1109: uint16(2),
	1110: uint16(227),
	1111: uint16(1),
	1112: uint16(anon_sym_RBRACK),
	1113: uint16(205),
	1114: uint16(3),
	1115: uint16(anon_sym_DOT),
	1116: uint16(anon_sym_POUND),
	1117: uint16(anon_sym_TILDE),
	1118: uint16(4),
	1119: uint16(229),
	1120: uint16(1),
	1121: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	1122: uint16(231),
	1123: uint16(1),
	1124: uint16(sym_code_block_language),
	1125: uint16(233),
	1126: uint16(1),
	1127: uint16(sym_code_block_line),
	1128: uint16(76),
	1129: uint16(1),
	1130: uint16(aux_sym_code_block_repeat1),
	1131: uint16(4),
	1132: uint16(235),
	1133: uint16(1),
	1134: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	1135: uint16(237),
	1136: uint16(1),
	1137: uint16(sym_code_block_language),
	1138: uint16(239),
	1139: uint16(1),
	1140: uint16(sym_code_block_line),
	1141: uint16(81),
	1142: uint16(1),
	1143: uint16(aux_sym_code_block_repeat1),
	1144: uint16(4),
	1145: uint16(166),
	1146: uint16(1),
	1147: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	1148: uint16(241),
	1149: uint16(1),
	1150: uint16(sym__text),
	1151: uint16(38),
	1152: uint16(1),
	1153: uint16(sym_code_block),
	1154: uint16(97),
	1155: uint16(1),
	1156: uint16(sym_description),
	1157: uint16(4),
	1158: uint16(166),
	1159: uint16(1),
	1160: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	1161: uint16(241),
	1162: uint16(1),
	1163: uint16(sym__text),
	1164: uint16(38),
	1165: uint16(1),
	1166: uint16(sym_code_block),
	1167: uint16(96),
	1168: uint16(1),
	1169: uint16(sym_description),
	1170: uint16(3),
	1171: uint16(243),
	1172: uint16(1),
	1173: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	1174: uint16(245),
	1175: uint16(1),
	1176: uint16(sym_code_block_line),
	1177: uint16(74),
	1178: uint16(1),
	1179: uint16(aux_sym_code_block_repeat1),
	1180: uint16(3),
	1181: uint16(247),
	1182: uint16(1),
	1183: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	1184: uint16(249),
	1185: uint16(1),
	1186: uint16(sym_code_block_line),
	1187: uint16(74),
	1188: uint16(1),
	1189: uint16(aux_sym_code_block_repeat1),
	1190: uint16(3),
	1191: uint16(207),
	1192: uint16(1),
	1193: uint16(anon_sym_COMMA),
	1194: uint16(252),
	1195: uint16(1),
	1196: uint16(anon_sym_RBRACK),
	1197: uint16(77),
	1198: uint16(1),
	1199: uint16(aux_sym_array_expression_repeat1),
	1200: uint16(3),
	1201: uint16(245),
	1202: uint16(1),
	1203: uint16(sym_code_block_line),
	1204: uint16(254),
	1205: uint16(1),
	1206: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	1207: uint16(74),
	1208: uint16(1),
	1209: uint16(aux_sym_code_block_repeat1),
	1210: uint16(3),
	1211: uint16(217),
	1212: uint16(1),
	1213: uint16(anon_sym_RBRACK),
	1214: uint16(256),
	1215: uint16(1),
	1216: uint16(anon_sym_COMMA),
	1217: uint16(77),
	1218: uint16(1),
	1219: uint16(aux_sym_array_expression_repeat1),
	1220: uint16(3),
	1221: uint16(254),
	1222: uint16(1),
	1223: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	1224: uint16(259),
	1225: uint16(1),
	1226: uint16(sym_code_block_line),
	1227: uint16(73),
	1228: uint16(1),
	1229: uint16(aux_sym_code_block_repeat1),
	1230: uint16(3),
	1231: uint16(245),
	1232: uint16(1),
	1233: uint16(sym_code_block_line),
	1234: uint16(261),
	1235: uint16(1),
	1236: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	1237: uint16(74),
	1238: uint16(1),
	1239: uint16(aux_sym_code_block_repeat1),
	1240: uint16(3),
	1241: uint16(263),
	1242: uint16(1),
	1243: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	1244: uint16(265),
	1245: uint16(1),
	1246: uint16(sym_code_block_line),
	1247: uint16(79),
	1248: uint16(1),
	1249: uint16(aux_sym_code_block_repeat1),
	1250: uint16(3),
	1251: uint16(245),
	1252: uint16(1),
	1253: uint16(sym_code_block_line),
	1254: uint16(263),
	1255: uint16(1),
	1256: uint16(anon_sym_BQUOTE_BQUOTE_BQUOTE),
	1257: uint16(74),
	1258: uint16(1),
	1259: uint16(aux_sym_code_block_repeat1),
	1260: uint16(3),
	1261: uint16(207),
	1262: uint16(1),
	1263: uint16(anon_sym_COMMA),
	1264: uint16(267),
	1265: uint16(1),
	1266: uint16(anon_sym_RBRACK),
	1267: uint16(77),
	1268: uint16(1),
	1269: uint16(aux_sym_array_expression_repeat1),
	1270: uint16(2),
	1271: uint16(269),
	1272: uint16(1),
	1273: uint16(sym_identifier),
	1274: uint16(20),
	1275: uint16(1),
	1276: uint16(sym_qualified_expression),
	1277: uint16(2),
	1278: uint16(271),
	1279: uint16(1),
	1280: uint16(sym_identifier),
	1281: uint16(58),
	1282: uint16(1),
	1283: uint16(sym_qualified_expression),
	1284: uint16(1),
	1285: uint16(273),
	1286: uint16(1),
	1287: uint16(sym_identifier),
	1288: uint16(1),
	1289: uint16(275),
	1290: uint16(1),
	1291: uint16(sym_type),
	1292: uint16(1),
	1293: uint16(277),
	1294: uint16(1),
	1295: uint16(anon_sym_RBRACE),
	1296: uint16(1),
	1297: uint16(279),
	1298: uint16(1),
	1300: uint16(1),
	1301: uint16(281),
	1302: uint16(1),
	1303: uint16(anon_sym_RBRACE),
	1304: uint16(1),
	1305: uint16(283),
	1306: uint16(1),
	1307: uint16(sym_type),
	1308: uint16(1),
	1309: uint16(285),
	1310: uint16(1),
	1312: uint16(1),
	1313: uint16(287),
	1314: uint16(1),
	1315: uint16(sym_tag_name),
	1316: uint16(1),
	1317: uint16(289),
	1318: uint16(1),
	1320: uint16(1),
	1321: uint16(291),
	1322: uint16(1),
	1323: uint16(sym_identifier),
	1324: uint16(1),
	1325: uint16(293),
	1326: uint16(1),
	1328: uint16(1),
	1329: uint16(295),
	1330: uint16(1),
	1331: uint16(anon_sym_RBRACE),
	1332: uint16(1),
	1333: uint16(297),
	1334: uint16(1),
	1335: uint16(anon_sym_RBRACE),
	1336: uint16(1),
	1337: uint16(299),
	1338: uint16(1),
	1339: uint16(sym_tag_name),
}

var ts_small_parse_table_map = [97]uint32_t{
	1:  uint32(45),
	2:  uint32(87),
	3:  uint32(113),
	4:  uint32(139),
	5:  uint32(159),
	6:  uint32(185),
	7:  uint32(211),
	8:  uint32(243),
	9:  uint32(269),
	10: uint32(286),
	11: uint32(300),
	12: uint32(322),
	13: uint32(344),
	14: uint32(366),
	15: uint32(380),
	16: uint32(394),
	17: uint32(408),
	18: uint32(432),
	19: uint32(446),
	20: uint32(460),
	21: uint32(474),
	22: uint32(489),
	23: uint32(510),
	24: uint32(525),
	25: uint32(538),
	26: uint32(557),
	27: uint32(576),
	28: uint32(591),
	29: uint32(612),
	30: uint32(631),
	31: uint32(652),
	32: uint32(665),
	33: uint32(684),
	34: uint32(705),
	35: uint32(718),
	36: uint32(731),
	37: uint32(752),
	38: uint32(771),
	39: uint32(792),
	40: uint32(813),
	41: uint32(833),
	42: uint32(853),
	43: uint32(867),
	44: uint32(887),
	45: uint32(898),
	46: uint32(915),
	47: uint32(930),
	48: uint32(941),
	49: uint32(956),
	50: uint32(967),
	51: uint32(975),
	52: uint32(983),
	53: uint32(993),
	54: uint32(1003),
	55: uint32(1013),
	56: uint32(1021),
	57: uint32(1029),
	58: uint32(1037),
	59: uint32(1045),
	60: uint32(1055),
	61: uint32(1063),
	62: uint32(1073),
	63: uint32(1082),
	64: uint32(1091),
	65: uint32(1100),
	66: uint32(1109),
	67: uint32(1118),
	68: uint32(1131),
	69: uint32(1144),
	70: uint32(1157),
	71: uint32(1170),
	72: uint32(1180),
	73: uint32(1190),
	74: uint32(1200),
	75: uint32(1210),
	76: uint32(1220),
	77: uint32(1230),
	78: uint32(1240),
	79: uint32(1250),
	80: uint32(1260),
	81: uint32(1270),
	82: uint32(1277),
	83: uint32(1284),
	84: uint32(1288),
	85: uint32(1292),
	86: uint32(1296),
	87: uint32(1300),
	88: uint32(1304),
	89: uint32(1308),
	90: uint32(1312),
	91: uint32(1316),
	92: uint32(1320),
	93: uint32(1324),
	94: uint32(1328),
	95: uint32(1332),
	96: uint32(1336),
}

var ts_parse_actions = [301]TSParseActionEntry{
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(23)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(86)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_tag),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(15)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(69)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(6)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(21)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(5)),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	20: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_tag),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_tag),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	24: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_tag),
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
		Fcount: uint8(1),
	}})),
	26: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_tag),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(83)),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	30: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_tag),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(92)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(8)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_description),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(8)),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	40: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_description),
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
		Fsymbol:      uint16(sym_expression),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(34)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(85)),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	48: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_tag),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fsymbol:      uint16(sym_tag),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(10)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_description),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(10)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fsymbol:      uint16(sym_description),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(2)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(19)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(24)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(88)),
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
		Fcount: uint8(2),
	}})),
	70: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_description_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(92)),
		Frepetition: libc.BoolUint8(1 != 0),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fsymbol:      uint16(aux_sym_description_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(10)),
		Frepetition: libc.BoolUint8(1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_description_repeat1),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	78: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_description_repeat1),
	})))),
	79: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(69)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
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
		Fcount: uint8(2),
	}})),
	81: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_description_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(10)),
		Frepetition: libc.BoolUint8(1 != 0),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	84: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_description_repeat1),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_member_expression),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	88: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_member_expression),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_array_expression),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	92: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_array_expression),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(14)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(12)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(44)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(53)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(52)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(31)),
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
		Fcount: uint8(1),
	}})),
	106: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_array_expression),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	108: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_array_expression),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_qualified_expression),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fsymbol:      uint16(sym_qualified_expression),
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
		Fsymbol:      uint16(sym_path_expression),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	116: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_path_expression),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(90)),
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
		Fcount: uint8(1),
	}})),
	120: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_array_expression),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	122: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_array_expression),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__begin),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(29)),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(aux_sym__begin_repeat1),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(1 != 0),
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
		Fsymbol:      uint16(aux_sym__begin_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(25)),
		Frepetition: libc.BoolUint8(1 != 0),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_code_block),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
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
		Fsymbol:      uint16(sym_code_block),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym__begin),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(25)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(39)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(94)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(46)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(27)),
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
		Fcount: uint8(1),
	}})),
	149: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_code_block),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	151: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_code_block),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(13)),
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
		Fcount: uint8(1),
	}})),
	155: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_code_block),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	157: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_code_block),
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
		Fcount: uint8(1),
	}})),
	159: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_inline_tag),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	161: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_inline_tag),
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
		Fcount: uint8(1),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(98)),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	165: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(40)),
	}})))),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(70)),
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
		Fcount: uint8(1),
	}})),
	169: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(40)),
	}})))),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(41)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(41)),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_description_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(98)),
		Frepetition: libc.BoolUint8(1 != 0),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fsymbol:      uint16(aux_sym_description_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(41)),
		Frepetition: libc.BoolUint8(1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(1 != 0),
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
		Fsymbol:      uint16(aux_sym_description_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(70)),
		Frepetition: libc.BoolUint8(1 != 0),
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
		Fsymbol:      uint16(aux_sym_description_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(41)),
		Frepetition: libc.BoolUint8(1 != 0),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(93)),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	189: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(91)),
	}})))),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_optional_identifier),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
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
		Fsymbol:      uint16(sym_optional_identifier),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_document_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(2)),
		Frepetition: libc.BoolUint8(1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_document_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(19)),
		Frepetition: libc.BoolUint8(1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_document_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(24)),
		Frepetition: libc.BoolUint8(1 != 0),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(84)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(28)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(16)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_optional_identifier),
		Fproduction_id: uint16(1),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	214: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_optional_identifier),
		Fproduction_id: uint16(1),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(59)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fsymbol:      uint16(aux_sym_array_expression_repeat1),
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
		Fcount: uint8(1),
	}})),
	220: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_tag),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	222: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_tag),
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
		Fcount: uint8(1),
	}})),
	224: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_tag),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	226: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_tag),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(49)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(26)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(78)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(76)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(54)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(80)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(81)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(38)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(33)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(74)),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_code_block_repeat1),
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
		Fcount: uint8(2),
	}})),
	250: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_code_block_repeat1),
	})))),
	251: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(74)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	253: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(22)),
	}})))),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(36)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	257: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_array_expression_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(28)),
		Frepetition: libc.BoolUint8(1 != 0),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(73)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(61)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(56)),
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
		Fcount: uint8(1),
	}})),
	266: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(79)),
	}})))),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(62)),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(11)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(51)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(18)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(87)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(3)),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	280: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(32)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(89)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fsymbol:      uint16(sym_document),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(71)),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	290: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_document),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(60)),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	294: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f uint8_t
		_ [7]byte
	}{f: uint8(TSParseActionTypeAccept)})),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(63)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	298: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(37)),
	}})))),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	300: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(72)),
	}})))),
}

type ts_external_scanner_symbol_identifiers = int32

const ts_external_token_type = 0
const ts_external_token_code_block_line = 1

var ts_external_scanner_symbol_map = [2]TSSymbol{
	0: uint16(sym_type),
	1: uint16(sym_code_block_line),
}

var ts_external_scanner_states = [3][2]uint8{
	1: {
		0: libc.BoolUint8(1 != 0),
	},
	2: {
		1: libc.BoolUint8(1 != 0),
	},
}

func tree_sitter_jsdoc(tls *libc.TLS) (r uintptr) {
	return uintptr(unsafe.Pointer(&language))
}

var language = TSLanguage{
	Fabi_version:               uint32(15),
	Fsymbol_count:              uint32(43),
	Ftoken_count:               uint32(25),
	Fexternal_token_count:      uint32(2),
	Fstate_count:               uint32(99),
	Flarge_state_count:         uint32(2),
	Fproduction_id_count:       uint32(2),
	Ffield_count:               uint32(1),
	Fmax_alias_sequence_length: uint16(6),
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
	Fname:                  __ccgo_ts + 397,
	Fsupertype_count:       uint32(1),
	Fsupertype_symbols:     uintptr(unsafe.Pointer(&ts_supertype_symbols)),
	Fsupertype_map_slices:  uintptr(unsafe.Pointer(&ts_supertype_map_slices)),
	Fsupertype_map_entries: uintptr(unsafe.Pointer(&ts_supertype_map_entries)),
	Fmetadata: TSLanguageMetadata{
		Fminor_version: uint8(25),
	},
}

func init() {
	p := unsafe.Pointer(&language)
	*(*uintptr)(unsafe.Add(p, 144)) = __ccgo_fp(ts_lex)
	*(*uintptr)(unsafe.Add(p, 184)) = __ccgo_fp(tree_sitter_jsdoc_external_scanner_create)
	*(*uintptr)(unsafe.Add(p, 192)) = __ccgo_fp(tree_sitter_jsdoc_external_scanner_destroy)
	*(*uintptr)(unsafe.Add(p, 200)) = __ccgo_fp(tree_sitter_jsdoc_external_scanner_scan)
	*(*uintptr)(unsafe.Add(p, 208)) = __ccgo_fp(tree_sitter_jsdoc_external_scanner_serialize)
	*(*uintptr)(unsafe.Add(p, 216)) = __ccgo_fp(tree_sitter_jsdoc_external_scanner_deserialize)
}

func __ccgo_fp(f interface{}) uintptr {
	type iface [2]uintptr
	return (*iface)(unsafe.Pointer(&f))[1]
}

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = "end\x00{\x00}\x00_inline_tag_false_positive\x00tag_name\x00:\x00/\x00.\x00#\x00~\x00[\x00,\x00]\x00```\x00code_block_language\x00code_block_line\x00=\x00identifier\x00number\x00_text\x00*\x00type\x00document\x00description\x00tag\x00inline_tag\x00expression\x00qualified_expression\x00path_expression\x00member_expression\x00array_expression\x00code_block\x00optional_identifier\x00_begin\x00_end\x00document_repeat1\x00description_repeat1\x00array_expression_repeat1\x00code_block_repeat1\x00_begin_repeat1\x00value\x00jsdoc\x00"
