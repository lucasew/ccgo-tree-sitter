// Code generated for linux/arm64 by 'ccgo /tmp/grammar-gen-2977731351/preprocessed.c -o /tmp/grammar-gen-2977731351/grammar.go', DO NOT EDIT.

//go:build linux && arm64

package grammar_regex

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

type ts_symbol_identifiers = int32

const anon_sym_PIPE = 1
const sym_any_character = 2
const anon_sym_CARET = 3
const sym_end_assertion = 4
const sym_boundary_assertion = 5
const sym_non_boundary_assertion = 6
const anon_sym_LPAREN_QMARK = 7
const anon_sym_EQ = 8
const anon_sym_BANG = 9
const anon_sym_RPAREN = 10
const anon_sym_LPAREN_QMARK_LT = 11
const sym_pattern_character = 12
const anon_sym_LBRACK = 13
const anon_sym_DASH = 14
const anon_sym_RBRACK = 15
const anon_sym_LBRACK_COLON = 16
const anon_sym_COLON_RBRACK = 17
const aux_sym_posix_class_name_token1 = 18
const anon_sym_BSLASH_DASH = 19
const sym_class_character = 20
const anon_sym_LPAREN = 21
const anon_sym_LPAREN_QMARKP_LT = 22
const anon_sym_GT = 23
const anon_sym_LPAREN_QMARK_COLON = 24
const anon_sym_COLON = 25
const anon_sym_STAR = 26
const anon_sym_QMARK = 27
const anon_sym_PLUS = 28
const anon_sym_LBRACE = 29
const anon_sym_COMMA = 30
const anon_sym_RBRACE = 31
const anon_sym_BSLASHk = 32
const anon_sym_LT = 33
const anon_sym_LPAREN_QMARKP_EQ = 34
const sym_decimal_escape = 35
const aux_sym_character_class_escape_token1 = 36
const aux_sym_character_class_escape_token2 = 37
const aux_sym_unicode_character_escape_token1 = 38
const aux_sym_unicode_character_escape_token2 = 39
const sym_unicode_property = 40
const aux_sym_control_escape_token1 = 41
const aux_sym_control_escape_token2 = 42
const sym_control_letter_escape = 43
const sym_identity_escape = 44
const sym_group_name = 45
const sym_decimal_digits = 46
const sym_pattern = 47
const sym_alternation = 48
const sym_term = 49
const sym_start_assertion = 50
const sym_lookaround_assertion = 51
const sym__lookahead_assertion = 52
const sym__lookbehind_assertion = 53
const sym_character_class = 54
const sym_posix_character_class = 55
const sym_posix_class_name = 56
const sym_class_range = 57
const sym_anonymous_capturing_group = 58
const sym_named_capturing_group = 59
const sym_non_capturing_group = 60
const sym_inline_flags_group = 61
const sym_flags = 62
const sym_zero_or_more = 63
const sym_one_or_more = 64
const sym_optional = 65
const sym_count_quantifier = 66
const sym_backreference_escape = 67
const sym_named_group_backreference = 68
const sym_character_class_escape = 69
const sym_unicode_character_escape = 70
const sym_unicode_property_value_expression = 71
const sym_control_escape = 72
const aux_sym_alternation_repeat1 = 73
const aux_sym_term_repeat1 = 74
const aux_sym_character_class_repeat1 = 75
const alias_sym_lazy = 76
const alias_sym_unicode_property_name = 77

var ts_symbol_names = [78]uintptr{
	0:  __ccgo_ts,
	1:  __ccgo_ts + 4,
	2:  __ccgo_ts + 6,
	3:  __ccgo_ts + 20,
	4:  __ccgo_ts + 22,
	5:  __ccgo_ts + 36,
	6:  __ccgo_ts + 55,
	7:  __ccgo_ts + 78,
	8:  __ccgo_ts + 81,
	9:  __ccgo_ts + 83,
	10: __ccgo_ts + 85,
	11: __ccgo_ts + 87,
	12: __ccgo_ts + 91,
	13: __ccgo_ts + 109,
	14: __ccgo_ts + 111,
	15: __ccgo_ts + 113,
	16: __ccgo_ts + 115,
	17: __ccgo_ts + 118,
	18: __ccgo_ts + 121,
	19: __ccgo_ts + 145,
	20: __ccgo_ts + 161,
	21: __ccgo_ts + 177,
	22: __ccgo_ts + 179,
	23: __ccgo_ts + 184,
	24: __ccgo_ts + 186,
	25: __ccgo_ts + 190,
	26: __ccgo_ts + 192,
	27: __ccgo_ts + 194,
	28: __ccgo_ts + 196,
	29: __ccgo_ts + 198,
	30: __ccgo_ts + 200,
	31: __ccgo_ts + 202,
	32: __ccgo_ts + 204,
	33: __ccgo_ts + 207,
	34: __ccgo_ts + 209,
	35: __ccgo_ts + 214,
	36: __ccgo_ts + 229,
	37: __ccgo_ts + 259,
	38: __ccgo_ts + 289,
	39: __ccgo_ts + 321,
	40: __ccgo_ts + 353,
	41: __ccgo_ts + 376,
	42: __ccgo_ts + 398,
	43: __ccgo_ts + 420,
	44: __ccgo_ts + 145,
	45: __ccgo_ts + 442,
	46: __ccgo_ts + 453,
	47: __ccgo_ts + 468,
	48: __ccgo_ts + 476,
	49: __ccgo_ts + 488,
	50: __ccgo_ts + 493,
	51: __ccgo_ts + 509,
	52: __ccgo_ts + 530,
	53: __ccgo_ts + 551,
	54: __ccgo_ts + 573,
	55: __ccgo_ts + 589,
	56: __ccgo_ts + 611,
	57: __ccgo_ts + 628,
	58: __ccgo_ts + 640,
	59: __ccgo_ts + 666,
	60: __ccgo_ts + 688,
	61: __ccgo_ts + 708,
	62: __ccgo_ts + 727,
	63: __ccgo_ts + 733,
	64: __ccgo_ts + 746,
	65: __ccgo_ts + 758,
	66: __ccgo_ts + 767,
	67: __ccgo_ts + 784,
	68: __ccgo_ts + 805,
	69: __ccgo_ts + 831,
	70: __ccgo_ts + 854,
	71: __ccgo_ts + 879,
	72: __ccgo_ts + 913,
	73: __ccgo_ts + 928,
	74: __ccgo_ts + 948,
	75: __ccgo_ts + 961,
	76: __ccgo_ts + 985,
	77: __ccgo_ts + 990,
}

var ts_symbol_map = [78]TSSymbol{
	1:  uint16(anon_sym_PIPE),
	2:  uint16(sym_any_character),
	3:  uint16(anon_sym_CARET),
	4:  uint16(sym_end_assertion),
	5:  uint16(sym_boundary_assertion),
	6:  uint16(sym_non_boundary_assertion),
	7:  uint16(anon_sym_LPAREN_QMARK),
	8:  uint16(anon_sym_EQ),
	9:  uint16(anon_sym_BANG),
	10: uint16(anon_sym_RPAREN),
	11: uint16(anon_sym_LPAREN_QMARK_LT),
	12: uint16(sym_pattern_character),
	13: uint16(anon_sym_LBRACK),
	14: uint16(anon_sym_DASH),
	15: uint16(anon_sym_RBRACK),
	16: uint16(anon_sym_LBRACK_COLON),
	17: uint16(anon_sym_COLON_RBRACK),
	18: uint16(aux_sym_posix_class_name_token1),
	19: uint16(sym_identity_escape),
	20: uint16(sym_class_character),
	21: uint16(anon_sym_LPAREN),
	22: uint16(anon_sym_LPAREN_QMARKP_LT),
	23: uint16(anon_sym_GT),
	24: uint16(anon_sym_LPAREN_QMARK_COLON),
	25: uint16(anon_sym_COLON),
	26: uint16(anon_sym_STAR),
	27: uint16(anon_sym_QMARK),
	28: uint16(anon_sym_PLUS),
	29: uint16(anon_sym_LBRACE),
	30: uint16(anon_sym_COMMA),
	31: uint16(anon_sym_RBRACE),
	32: uint16(anon_sym_BSLASHk),
	33: uint16(anon_sym_LT),
	34: uint16(anon_sym_LPAREN_QMARKP_EQ),
	35: uint16(sym_decimal_escape),
	36: uint16(aux_sym_character_class_escape_token1),
	37: uint16(aux_sym_character_class_escape_token2),
	38: uint16(aux_sym_unicode_character_escape_token1),
	39: uint16(aux_sym_unicode_character_escape_token2),
	40: uint16(sym_unicode_property),
	41: uint16(aux_sym_control_escape_token1),
	42: uint16(aux_sym_control_escape_token2),
	43: uint16(sym_control_letter_escape),
	44: uint16(sym_identity_escape),
	45: uint16(sym_group_name),
	46: uint16(sym_decimal_digits),
	47: uint16(sym_pattern),
	48: uint16(sym_alternation),
	49: uint16(sym_term),
	50: uint16(sym_start_assertion),
	51: uint16(sym_lookaround_assertion),
	52: uint16(sym__lookahead_assertion),
	53: uint16(sym__lookbehind_assertion),
	54: uint16(sym_character_class),
	55: uint16(sym_posix_character_class),
	56: uint16(sym_posix_class_name),
	57: uint16(sym_class_range),
	58: uint16(sym_anonymous_capturing_group),
	59: uint16(sym_named_capturing_group),
	60: uint16(sym_non_capturing_group),
	61: uint16(sym_inline_flags_group),
	62: uint16(sym_flags),
	63: uint16(sym_zero_or_more),
	64: uint16(sym_one_or_more),
	65: uint16(sym_optional),
	66: uint16(sym_count_quantifier),
	67: uint16(sym_backreference_escape),
	68: uint16(sym_named_group_backreference),
	69: uint16(sym_character_class_escape),
	70: uint16(sym_unicode_character_escape),
	71: uint16(sym_unicode_property_value_expression),
	72: uint16(sym_control_escape),
	73: uint16(aux_sym_alternation_repeat1),
	74: uint16(aux_sym_term_repeat1),
	75: uint16(aux_sym_character_class_repeat1),
	76: uint16(alias_sym_lazy),
	77: uint16(alias_sym_unicode_property_name),
}

var ts_symbol_metadata = [78]TSSymbolMetadata{
	0: {
		Fnamed: libc.BoolUint8(1 != 0),
	},
	1: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	2: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	3: {
		Fvisible: libc.BoolUint8(1 != 0),
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
		Fnamed:   libc.BoolUint8(1 != 0),
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
	},
	17: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	18: {},
	19: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	20: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	21: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	22: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	23: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	24: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	25: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	26: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	27: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	28: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	29: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	30: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	31: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	32: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	33: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	34: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	35: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	36: {},
	37: {},
	38: {},
	39: {},
	40: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	41: {},
	42: {},
	43: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	44: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	45: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	46: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	47: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	48: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	49: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	50: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	51: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	52: {
		Fnamed: libc.BoolUint8(1 != 0),
	},
	53: {
		Fnamed: libc.BoolUint8(1 != 0),
	},
	54: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	55: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	56: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	57: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	58: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	59: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	60: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	61: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	62: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	63: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	64: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	65: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	66: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	67: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	68: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	69: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	70: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	71: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	72: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	73: {},
	74: {},
	75: {},
	76: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	77: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
}

var ts_alias_sequences = [15][7]TSSymbol{
	0: {},
	1: {
		1: uint16(alias_sym_lazy),
	},
	2: {
		1: uint16(sym_class_character),
	},
	3: {
		2: uint16(sym_class_character),
	},
	4: {
		0: uint16(sym_class_character),
		2: uint16(sym_class_character),
	},
	5: {
		1: uint16(sym_class_character),
		2: uint16(sym_class_character),
	},
	6: {
		0: uint16(sym_class_character),
	},
	7: {
		3: uint16(alias_sym_lazy),
	},
	8: {
		2: uint16(sym_class_character),
		3: uint16(sym_class_character),
	},
	9: {
		3: uint16(sym_class_character),
	},
	10: {
		1: uint16(sym_class_character),
		3: uint16(sym_class_character),
	},
	11: {
		0: uint16(alias_sym_unicode_property_name),
	},
	12: {
		4: uint16(alias_sym_lazy),
	},
	13: {
		2: uint16(sym_class_character),
		4: uint16(sym_class_character),
	},
	14: {
		5: uint16(alias_sym_lazy),
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
	40:  uint16(40),
	41:  uint16(41),
	42:  uint16(42),
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
	72:  uint16(27),
	73:  uint16(73),
	74:  uint16(14),
	75:  uint16(15),
	76:  uint16(22),
	77:  uint16(34),
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
	94:  uint16(92),
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
	106: uint16(96),
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
	132: uint16(115),
	133: uint16(116),
	134: uint16(134),
	135: uint16(135),
	136: uint16(111),
}

func ts_lex(tls *libc.TLS, lexer uintptr, state TSStateId) (r uint8) {
	var eof, result, skip uint8
	var i, i1, i2, i3 uint32_t
	var lookahead int32_t
	_, _, _, _, _, _, _, _ = eof, i, i1, i2, i3, lookahead, result, skip
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
			state = uint16(34)
			goto next_state
		}
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(0)
			goto next_state
		}
		if lookahead == int32('\r') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(29)
			goto next_state
		}
		if lookahead == int32('!') {
			state = uint16(43)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(38)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(59)
			goto next_state
		}
		if lookahead == int32(')') {
			state = uint16(44)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(65)
			goto next_state
		}
		if lookahead == int32('+') {
			state = uint16(67)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(69)
			goto next_state
		}
		if lookahead == int32('-') {
			state = uint16(48)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(36)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(64)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(72)
			goto next_state
		}
		if lookahead == int32('=') {
			state = uint16(42)
			goto next_state
		}
		if lookahead == int32('>') {
			state = uint16(61)
			goto next_state
		}
		if lookahead == int32('?') {
			state = uint16(66)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(47)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(10)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(49)
			goto next_state
		}
		if lookahead == int32('^') {
			state = uint16(37)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(68)
			goto next_state
		}
		if lookahead == int32('|') {
			state = uint16(35)
			goto next_state
		}
		if lookahead == int32('}') {
			state = uint16(70)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(46)
			goto next_state
		}
		return result
	case int32(1):
		if lookahead == int32('\n') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('\r') {
			state = uint16(54)
			goto next_state
		}
		if lookahead == int32('-') {
			state = uint16(48)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(58)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(11)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(49)
			goto next_state
		}
		if lookahead == int32('^') {
			state = uint16(37)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(54)
			goto next_state
		}
		return result
	case int32(2):
		if lookahead == int32('\n') {
			state = uint16(56)
			goto next_state
		}
		if lookahead == int32('\r') {
			state = uint16(54)
			goto next_state
		}
		if lookahead == int32('-') {
			state = uint16(48)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(58)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(11)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(49)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(54)
			goto next_state
		}
		return result
	case int32(3):
		if lookahead == int32('\n') {
			state = uint16(57)
			goto next_state
		}
		if lookahead == int32('\r') {
			state = uint16(54)
			goto next_state
		}
		if lookahead == int32('-') {
			state = uint16(48)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(49)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(54)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(5)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(5)
			goto next_state
		}
		if lookahead == int32('\r') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(4)
			goto next_state
		}
		if lookahead == int32('!') {
			state = uint16(43)
			goto next_state
		}
		if lookahead == int32(')') {
			state = uint16(44)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(69)
			goto next_state
		}
		if lookahead == int32('-') {
			state = uint16(48)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(63)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(72)
			goto next_state
		}
		if lookahead == int32('=') {
			state = uint16(42)
			goto next_state
		}
		if lookahead == int32('>') {
			state = uint16(61)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(49)
			goto next_state
		}
		if lookahead == int32('}') {
			state = uint16(70)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(88)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(7)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(7)
			goto next_state
		}
		if lookahead == int32('\r') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(6)
			goto next_state
		}
		if lookahead == int32('!') {
			state = uint16(43)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('=') {
			state = uint16(42)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(87)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(9)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(9)
			goto next_state
		}
		if lookahead == int32('\r') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(8)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(79)
			goto next_state
		}
		return result
	case int32(10):
		i = uint32(0)
		for {
			if !(uint64(i) < libc.Uint64FromInt64(84)/libc.Uint64FromInt64(2)) {
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
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(74)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(83)
			goto next_state
		}
		return result
	case int32(11):
		i1 = uint32(0)
		for {
			if !(uint64(i1) < libc.Uint64FromInt64(76)/libc.Uint64FromInt64(2)) {
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
		if lookahead != 0 && (lookahead < int32('0') || int32('9') < lookahead) && lookahead != int32('k') {
			state = uint16(83)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead == int32('<') {
			state = uint16(60)
			goto next_state
		}
		if lookahead == int32('=') {
			state = uint16(73)
			goto next_state
		}
		return result
	case int32(13):
		i2 = uint32(0)
		for {
			if !(uint64(i2) < libc.Uint64FromInt64(80)/libc.Uint64FromInt64(2)) {
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
			state = uint16(74)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(83)
			goto next_state
		}
		return result
	case int32(14):
		if lookahead == int32(']') {
			state = uint16(51)
			goto next_state
		}
		return result
	case int32(15):
		i3 = uint32(0)
		for {
			if !(uint64(i3) < libc.Uint64FromInt64(68)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token3[i3]) == lookahead {
				state = map_token3[i3+uint32(1)]
				goto next_state
			}
			goto _4
		_4:
			;
			i3 = i3 + uint32(2)
		}
		return result
	case int32(16):
		if lookahead == int32('{') {
			state = uint16(23)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(28)
			goto next_state
		}
		return result
	case int32(17):
		if lookahead == int32('}') {
			state = uint16(78)
			goto next_state
		}
		return result
	case int32(18):
		if lookahead == int32('}') {
			state = uint16(78)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(17)
			goto next_state
		}
		return result
	case int32(19):
		if lookahead == int32('}') {
			state = uint16(78)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(18)
			goto next_state
		}
		return result
	case int32(20):
		if lookahead == int32('}') {
			state = uint16(78)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(19)
			goto next_state
		}
		return result
	case int32(21):
		if lookahead == int32('}') {
			state = uint16(78)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(20)
			goto next_state
		}
		return result
	case int32(22):
		if lookahead == int32('}') {
			state = uint16(78)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(21)
			goto next_state
		}
		return result
	case int32(23):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(22)
			goto next_state
		}
		return result
	case int32(24):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(81)
			goto next_state
		}
		return result
	case int32(25):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(77)
			goto next_state
		}
		return result
	case int32(26):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(24)
			goto next_state
		}
		return result
	case int32(27):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(25)
			goto next_state
		}
		return result
	case int32(28):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(27)
			goto next_state
		}
		return result
	case int32(29):
		if eof != 0 {
			state = uint16(34)
			goto next_state
		}
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(0)
			goto next_state
		}
		return result
	case int32(30):
		if eof != 0 {
			state = uint16(34)
			goto next_state
		}
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(31)
			goto next_state
		}
		return result
	case int32(31):
		if eof != 0 {
			state = uint16(34)
			goto next_state
		}
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(31)
			goto next_state
		}
		if lookahead == int32('\r') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(30)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(38)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(59)
			goto next_state
		}
		if lookahead == int32(')') {
			state = uint16(44)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(36)
			goto next_state
		}
		if lookahead == int32('?') {
			state = uint16(66)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(47)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(13)
			goto next_state
		}
		if lookahead == int32('^') {
			state = uint16(37)
			goto next_state
		}
		if lookahead == int32('|') {
			state = uint16(35)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('(') || int32('+') < lookahead) && (lookahead < int32('[') || int32('^') < lookahead) {
			state = uint16(46)
			goto next_state
		}
		return result
	case int32(32):
		if eof != 0 {
			state = uint16(34)
			goto next_state
		}
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(33):
		if eof != 0 {
			state = uint16(34)
			goto next_state
		}
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(33)
			goto next_state
		}
		if lookahead == int32('\r') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(32)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(38)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(59)
			goto next_state
		}
		if lookahead == int32(')') {
			state = uint16(44)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(65)
			goto next_state
		}
		if lookahead == int32('+') {
			state = uint16(67)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(36)
			goto next_state
		}
		if lookahead == int32('?') {
			state = uint16(66)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(47)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(13)
			goto next_state
		}
		if lookahead == int32('^') {
			state = uint16(37)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(68)
			goto next_state
		}
		if lookahead == int32('|') {
			state = uint16(35)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('[') || int32('^') < lookahead) {
			state = uint16(46)
			goto next_state
		}
		return result
	case int32(34):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(0)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(35):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PIPE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(36):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_any_character)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(37):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_CARET)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(38):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_end_assertion)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(39):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_boundary_assertion)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(40):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_non_boundary_assertion)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(41):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LPAREN_QMARK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(':') {
			state = uint16(62)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(45)
			goto next_state
		}
		if lookahead == int32('P') {
			state = uint16(12)
			goto next_state
		}
		return result
	case int32(42):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(43):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BANG)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(44):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(45):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LPAREN_QMARK_LT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(46):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_pattern_character)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(47):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(':') {
			state = uint16(50)
			goto next_state
		}
		return result
	case int32(48):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(49):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(50):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACK_COLON)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(51):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COLON_RBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(52):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_posix_class_name_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(53):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BSLASH_DASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(54):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_class_character)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(55):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_class_character)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('\r') {
			state = uint16(54)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(58)
			goto next_state
		}
		if lookahead == int32('^') {
			state = uint16(37)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('-') && (lookahead < int32('[') || int32('^') < lookahead) {
			state = uint16(54)
			goto next_state
		}
		return result
	case int32(56):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_class_character)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(56)
			goto next_state
		}
		if lookahead == int32('\r') {
			state = uint16(54)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(58)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('-') && (lookahead < int32('[') || int32(']') < lookahead) {
			state = uint16(54)
			goto next_state
		}
		return result
	case int32(57):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_class_character)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(57)
			goto next_state
		}
		if lookahead == int32('\r') {
			state = uint16(54)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('-') && lookahead != int32('\\') && lookahead != int32(']') {
			state = uint16(54)
			goto next_state
		}
		return result
	case int32(58):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_class_character)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(':') {
			state = uint16(50)
			goto next_state
		}
		return result
	case int32(59):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('?') {
			state = uint16(41)
			goto next_state
		}
		return result
	case int32(60):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LPAREN_QMARKP_LT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(61):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_GT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(62):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LPAREN_QMARK_COLON)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(63):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COLON)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(64):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COLON)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(']') {
			state = uint16(51)
			goto next_state
		}
		return result
	case int32(65):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STAR)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(66):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_QMARK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(67):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PLUS)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(68):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(69):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COMMA)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(70):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(71):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BSLASHk)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(72):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(73):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LPAREN_QMARKP_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(74):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_decimal_escape)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(74)
			goto next_state
		}
		return result
	case int32(75):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_character_class_escape_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(76):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_character_class_escape_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(77):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_unicode_character_escape_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(78):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_unicode_character_escape_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(79):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_unicode_property)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(79)
			goto next_state
		}
		return result
	case int32(80):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_control_escape_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(81):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_control_escape_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(82):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_control_letter_escape)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(83):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identity_escape)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(84):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identity_escape)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('{') {
			state = uint16(23)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(28)
			goto next_state
		}
		return result
	case int32(85):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identity_escape)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(24)
			goto next_state
		}
		return result
	case int32(86):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identity_escape)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(82)
			goto next_state
		}
		return result
	case int32(87):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_group_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(87)
			goto next_state
		}
		return result
	case int32(88):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_decimal_digits)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(88)
			goto next_state
		}
		return result
	default:
		return libc.BoolUint8(0 != 0)
	}
	return r
}

var map_token = [42]uint16_t{
	0:  uint16('-'),
	1:  uint16(53),
	2:  uint16('B'),
	3:  uint16(40),
	4:  uint16('b'),
	5:  uint16(39),
	6:  uint16('c'),
	7:  uint16(86),
	8:  uint16('k'),
	9:  uint16(71),
	10: uint16('u'),
	11: uint16(84),
	12: uint16('x'),
	13: uint16(85),
	14: uint16('P'),
	15: uint16(76),
	16: uint16('p'),
	17: uint16(76),
	18: uint16('0'),
	19: uint16(80),
	20: uint16('f'),
	21: uint16(80),
	22: uint16('n'),
	23: uint16(80),
	24: uint16('r'),
	25: uint16(80),
	26: uint16('t'),
	27: uint16(80),
	28: uint16('v'),
	29: uint16(80),
	30: uint16('D'),
	31: uint16(75),
	32: uint16('S'),
	33: uint16(75),
	34: uint16('W'),
	35: uint16(75),
	36: uint16('d'),
	37: uint16(75),
	38: uint16('s'),
	39: uint16(75),
	40: uint16('w'),
	41: uint16(75),
}

var map_token1 = [38]uint16_t{
	0:  uint16('-'),
	1:  uint16(53),
	2:  uint16('c'),
	3:  uint16(86),
	4:  uint16('u'),
	5:  uint16(84),
	6:  uint16('x'),
	7:  uint16(85),
	8:  uint16('P'),
	9:  uint16(76),
	10: uint16('p'),
	11: uint16(76),
	12: uint16('D'),
	13: uint16(75),
	14: uint16('S'),
	15: uint16(75),
	16: uint16('W'),
	17: uint16(75),
	18: uint16('d'),
	19: uint16(75),
	20: uint16('s'),
	21: uint16(75),
	22: uint16('w'),
	23: uint16(75),
	24: uint16('0'),
	25: uint16(80),
	26: uint16('b'),
	27: uint16(80),
	28: uint16('f'),
	29: uint16(80),
	30: uint16('n'),
	31: uint16(80),
	32: uint16('r'),
	33: uint16(80),
	34: uint16('t'),
	35: uint16(80),
	36: uint16('v'),
	37: uint16(80),
}

var map_token2 = [40]uint16_t{
	0:  uint16('B'),
	1:  uint16(40),
	2:  uint16('b'),
	3:  uint16(39),
	4:  uint16('c'),
	5:  uint16(86),
	6:  uint16('k'),
	7:  uint16(71),
	8:  uint16('u'),
	9:  uint16(84),
	10: uint16('x'),
	11: uint16(85),
	12: uint16('P'),
	13: uint16(76),
	14: uint16('p'),
	15: uint16(76),
	16: uint16('0'),
	17: uint16(80),
	18: uint16('f'),
	19: uint16(80),
	20: uint16('n'),
	21: uint16(80),
	22: uint16('r'),
	23: uint16(80),
	24: uint16('t'),
	25: uint16(80),
	26: uint16('v'),
	27: uint16(80),
	28: uint16('D'),
	29: uint16(75),
	30: uint16('S'),
	31: uint16(75),
	32: uint16('W'),
	33: uint16(75),
	34: uint16('d'),
	35: uint16(75),
	36: uint16('s'),
	37: uint16(75),
	38: uint16('w'),
	39: uint16(75),
}

var map_token3 = [34]uint16_t{
	0:  uint16('u'),
	1:  uint16(16),
	2:  uint16('x'),
	3:  uint16(26),
	4:  uint16('P'),
	5:  uint16(76),
	6:  uint16('p'),
	7:  uint16(76),
	8:  uint16('D'),
	9:  uint16(75),
	10: uint16('S'),
	11: uint16(75),
	12: uint16('W'),
	13: uint16(75),
	14: uint16('d'),
	15: uint16(75),
	16: uint16('s'),
	17: uint16(75),
	18: uint16('w'),
	19: uint16(75),
	20: uint16('0'),
	21: uint16(80),
	22: uint16('b'),
	23: uint16(80),
	24: uint16('f'),
	25: uint16(80),
	26: uint16('n'),
	27: uint16(80),
	28: uint16('r'),
	29: uint16(80),
	30: uint16('t'),
	31: uint16(80),
	32: uint16('v'),
	33: uint16(80),
}

var ts_lex_modes = [137]TSLexerMode{
	0: {},
	1: {
		Flex_state: uint16(31),
	},
	2: {
		Flex_state: uint16(31),
	},
	3: {
		Flex_state: uint16(31),
	},
	4: {
		Flex_state: uint16(31),
	},
	5: {
		Flex_state: uint16(31),
	},
	6: {
		Flex_state: uint16(31),
	},
	7: {
		Flex_state: uint16(31),
	},
	8: {
		Flex_state: uint16(31),
	},
	9: {
		Flex_state: uint16(31),
	},
	10: {
		Flex_state: uint16(31),
	},
	11: {
		Flex_state: uint16(31),
	},
	12: {
		Flex_state: uint16(31),
	},
	13: {
		Flex_state: uint16(33),
	},
	14: {
		Flex_state: uint16(33),
	},
	15: {
		Flex_state: uint16(33),
	},
	16: {
		Flex_state: uint16(33),
	},
	17: {
		Flex_state: uint16(33),
	},
	18: {
		Flex_state: uint16(33),
	},
	19: {
		Flex_state: uint16(33),
	},
	20: {
		Flex_state: uint16(33),
	},
	21: {
		Flex_state: uint16(33),
	},
	22: {
		Flex_state: uint16(33),
	},
	23: {
		Flex_state: uint16(33),
	},
	24: {
		Flex_state: uint16(33),
	},
	25: {
		Flex_state: uint16(33),
	},
	26: {
		Flex_state: uint16(33),
	},
	27: {
		Flex_state: uint16(33),
	},
	28: {
		Flex_state: uint16(33),
	},
	29: {
		Flex_state: uint16(33),
	},
	30: {
		Flex_state: uint16(33),
	},
	31: {
		Flex_state: uint16(33),
	},
	32: {
		Flex_state: uint16(33),
	},
	33: {
		Flex_state: uint16(33),
	},
	34: {
		Flex_state: uint16(33),
	},
	35: {
		Flex_state: uint16(33),
	},
	36: {
		Flex_state: uint16(33),
	},
	37: {
		Flex_state: uint16(33),
	},
	38: {
		Flex_state: uint16(33),
	},
	39: {
		Flex_state: uint16(33),
	},
	40: {
		Flex_state: uint16(33),
	},
	41: {
		Flex_state: uint16(33),
	},
	42: {
		Flex_state: uint16(33),
	},
	43: {
		Flex_state: uint16(33),
	},
	44: {
		Flex_state: uint16(33),
	},
	45: {
		Flex_state: uint16(31),
	},
	46: {
		Flex_state: uint16(31),
	},
	47: {
		Flex_state: uint16(31),
	},
	48: {
		Flex_state: uint16(31),
	},
	49: {
		Flex_state: uint16(31),
	},
	50: {
		Flex_state: uint16(31),
	},
	51: {
		Flex_state: uint16(31),
	},
	52: {
		Flex_state: uint16(31),
	},
	53: {
		Flex_state: uint16(31),
	},
	54: {
		Flex_state: uint16(31),
	},
	55: {
		Flex_state: uint16(31),
	},
	56: {
		Flex_state: uint16(31),
	},
	57: {
		Flex_state: uint16(31),
	},
	58: {
		Flex_state: uint16(1),
	},
	59: {
		Flex_state: uint16(2),
	},
	60: {
		Flex_state: uint16(2),
	},
	61: {
		Flex_state: uint16(2),
	},
	62: {
		Flex_state: uint16(2),
	},
	63: {
		Flex_state: uint16(2),
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
		Flex_state: uint16(2),
	},
	69: {
		Flex_state: uint16(2),
	},
	70: {
		Flex_state: uint16(2),
	},
	71: {
		Flex_state: uint16(2),
	},
	72: {
		Flex_state: uint16(2),
	},
	73: {
		Flex_state: uint16(2),
	},
	74: {
		Flex_state: uint16(2),
	},
	75: {
		Flex_state: uint16(2),
	},
	76: {
		Flex_state: uint16(2),
	},
	77: {
		Flex_state: uint16(2),
	},
	78: {
		Flex_state: uint16(3),
	},
	79: {
		Flex_state: uint16(3),
	},
	80: {
		Flex_state: uint16(3),
	},
	81: {
		Flex_state: uint16(3),
	},
	82: {
		Flex_state: uint16(5),
	},
	83: {},
	84: {},
	85: {},
	86: {},
	87: {
		Flex_state: uint16(5),
	},
	88: {
		Flex_state: uint16(7),
	},
	89: {},
	90: {
		Flex_state: uint16(5),
	},
	91: {},
	92: {
		Flex_state: uint16(9),
	},
	93: {
		Flex_state: uint16(5),
	},
	94: {
		Flex_state: uint16(9),
	},
	95: {
		Flex_state: uint16(5),
	},
	96: {
		Flex_state: uint16(5),
	},
	97: {
		Flex_state: uint16(5),
	},
	98: {
		Flex_state: uint16(5),
	},
	99: {
		Flex_state: uint16(5),
	},
	100: {
		Flex_state: uint16(5),
	},
	101: {
		Flex_state: uint16(5),
	},
	102: {
		Flex_state: uint16(5),
	},
	103: {
		Flex_state: uint16(5),
	},
	104: {
		Flex_state: uint16(5),
	},
	105: {
		Flex_state: uint16(5),
	},
	106: {
		Flex_state: uint16(5),
	},
	107: {
		Flex_state: uint16(5),
	},
	108: {},
	109: {},
	110: {},
	111: {},
	112: {
		Flex_state: uint16(5),
	},
	113: {
		Flex_state: uint16(7),
	},
	114: {},
	115: {
		Flex_state: uint16(7),
	},
	116: {
		Flex_state: uint16(5),
	},
	117: {
		Flex_state: uint16(5),
	},
	118: {
		Flex_state: uint16(5),
	},
	119: {
		Flex_state: uint16(5),
	},
	120: {
		Flex_state: uint16(5),
	},
	121: {
		Flex_state: uint16(5),
	},
	122: {},
	123: {},
	124: {
		Flex_state: uint16(7),
	},
	125: {
		Flex_state: uint16(9),
	},
	126: {},
	127: {},
	128: {},
	129: {},
	130: {
		Flex_state: uint16(5),
	},
	131: {
		Flex_state: uint16(7),
	},
	132: {
		Flex_state: uint16(7),
	},
	133: {
		Flex_state: uint16(5),
	},
	134: {
		Flex_state: uint16(5),
	},
	135: {
		Flex_state: uint16(7),
	},
	136: {},
}

var ts_parse_table = [13][76]uint16_t{
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
		19: uint16(1),
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
		36: uint16(1),
		37: uint16(1),
		38: uint16(1),
		39: uint16(1),
		41: uint16(1),
		42: uint16(1),
		43: uint16(1),
		44: uint16(1),
	},
	1: {
		1:  uint16(3),
		2:  uint16(5),
		3:  uint16(7),
		4:  uint16(5),
		5:  uint16(5),
		6:  uint16(5),
		7:  uint16(9),
		11: uint16(11),
		12: uint16(5),
		13: uint16(13),
		16: uint16(15),
		21: uint16(17),
		22: uint16(19),
		24: uint16(21),
		32: uint16(23),
		34: uint16(25),
		35: uint16(5),
		36: uint16(27),
		37: uint16(29),
		38: uint16(31),
		39: uint16(31),
		41: uint16(33),
		42: uint16(35),
		43: uint16(5),
		44: uint16(37),
		47: uint16(129),
		48: uint16(91),
		49: uint16(85),
		50: uint16(13),
		51: uint16(13),
		52: uint16(16),
		53: uint16(16),
		54: uint16(13),
		55: uint16(13),
		58: uint16(13),
		59: uint16(13),
		60: uint16(13),
		61: uint16(13),
		67: uint16(13),
		68: uint16(13),
		69: uint16(13),
		70: uint16(14),
		72: uint16(13),
		73: uint16(86),
		74: uint16(11),
	},
	2: {
		1:  uint16(3),
		2:  uint16(5),
		3:  uint16(7),
		4:  uint16(5),
		5:  uint16(5),
		6:  uint16(5),
		7:  uint16(9),
		11: uint16(11),
		12: uint16(5),
		13: uint16(13),
		16: uint16(15),
		21: uint16(17),
		22: uint16(19),
		24: uint16(21),
		32: uint16(23),
		34: uint16(25),
		35: uint16(5),
		36: uint16(27),
		37: uint16(29),
		38: uint16(31),
		39: uint16(31),
		41: uint16(33),
		42: uint16(35),
		43: uint16(5),
		44: uint16(37),
		47: uint16(126),
		48: uint16(91),
		49: uint16(85),
		50: uint16(13),
		51: uint16(13),
		52: uint16(16),
		53: uint16(16),
		54: uint16(13),
		55: uint16(13),
		58: uint16(13),
		59: uint16(13),
		60: uint16(13),
		61: uint16(13),
		67: uint16(13),
		68: uint16(13),
		69: uint16(13),
		70: uint16(14),
		72: uint16(13),
		73: uint16(86),
		74: uint16(11),
	},
	3: {
		1:  uint16(3),
		2:  uint16(5),
		3:  uint16(7),
		4:  uint16(5),
		5:  uint16(5),
		6:  uint16(5),
		7:  uint16(9),
		11: uint16(11),
		12: uint16(5),
		13: uint16(13),
		16: uint16(15),
		21: uint16(17),
		22: uint16(19),
		24: uint16(21),
		32: uint16(23),
		34: uint16(25),
		35: uint16(5),
		36: uint16(27),
		37: uint16(29),
		38: uint16(31),
		39: uint16(31),
		41: uint16(33),
		42: uint16(35),
		43: uint16(5),
		44: uint16(37),
		47: uint16(108),
		48: uint16(91),
		49: uint16(85),
		50: uint16(13),
		51: uint16(13),
		52: uint16(16),
		53: uint16(16),
		54: uint16(13),
		55: uint16(13),
		58: uint16(13),
		59: uint16(13),
		60: uint16(13),
		61: uint16(13),
		67: uint16(13),
		68: uint16(13),
		69: uint16(13),
		70: uint16(14),
		72: uint16(13),
		73: uint16(86),
		74: uint16(11),
	},
	4: {
		1:  uint16(3),
		2:  uint16(5),
		3:  uint16(7),
		4:  uint16(5),
		5:  uint16(5),
		6:  uint16(5),
		7:  uint16(9),
		11: uint16(11),
		12: uint16(5),
		13: uint16(13),
		16: uint16(15),
		21: uint16(17),
		22: uint16(19),
		24: uint16(21),
		32: uint16(23),
		34: uint16(25),
		35: uint16(5),
		36: uint16(27),
		37: uint16(29),
		38: uint16(31),
		39: uint16(31),
		41: uint16(33),
		42: uint16(35),
		43: uint16(5),
		44: uint16(37),
		47: uint16(110),
		48: uint16(91),
		49: uint16(85),
		50: uint16(13),
		51: uint16(13),
		52: uint16(16),
		53: uint16(16),
		54: uint16(13),
		55: uint16(13),
		58: uint16(13),
		59: uint16(13),
		60: uint16(13),
		61: uint16(13),
		67: uint16(13),
		68: uint16(13),
		69: uint16(13),
		70: uint16(14),
		72: uint16(13),
		73: uint16(86),
		74: uint16(11),
	},
	5: {
		1:  uint16(3),
		2:  uint16(5),
		3:  uint16(7),
		4:  uint16(5),
		5:  uint16(5),
		6:  uint16(5),
		7:  uint16(9),
		11: uint16(11),
		12: uint16(5),
		13: uint16(13),
		16: uint16(15),
		21: uint16(17),
		22: uint16(19),
		24: uint16(21),
		32: uint16(23),
		34: uint16(25),
		35: uint16(5),
		36: uint16(27),
		37: uint16(29),
		38: uint16(31),
		39: uint16(31),
		41: uint16(33),
		42: uint16(35),
		43: uint16(5),
		44: uint16(37),
		47: uint16(127),
		48: uint16(91),
		49: uint16(85),
		50: uint16(13),
		51: uint16(13),
		52: uint16(16),
		53: uint16(16),
		54: uint16(13),
		55: uint16(13),
		58: uint16(13),
		59: uint16(13),
		60: uint16(13),
		61: uint16(13),
		67: uint16(13),
		68: uint16(13),
		69: uint16(13),
		70: uint16(14),
		72: uint16(13),
		73: uint16(86),
		74: uint16(11),
	},
	6: {
		1:  uint16(3),
		2:  uint16(5),
		3:  uint16(7),
		4:  uint16(5),
		5:  uint16(5),
		6:  uint16(5),
		7:  uint16(9),
		11: uint16(11),
		12: uint16(5),
		13: uint16(13),
		16: uint16(15),
		21: uint16(17),
		22: uint16(19),
		24: uint16(21),
		32: uint16(23),
		34: uint16(25),
		35: uint16(5),
		36: uint16(27),
		37: uint16(29),
		38: uint16(31),
		39: uint16(31),
		41: uint16(33),
		42: uint16(35),
		43: uint16(5),
		44: uint16(37),
		47: uint16(109),
		48: uint16(91),
		49: uint16(85),
		50: uint16(13),
		51: uint16(13),
		52: uint16(16),
		53: uint16(16),
		54: uint16(13),
		55: uint16(13),
		58: uint16(13),
		59: uint16(13),
		60: uint16(13),
		61: uint16(13),
		67: uint16(13),
		68: uint16(13),
		69: uint16(13),
		70: uint16(14),
		72: uint16(13),
		73: uint16(86),
		74: uint16(11),
	},
	7: {
		1:  uint16(3),
		2:  uint16(5),
		3:  uint16(7),
		4:  uint16(5),
		5:  uint16(5),
		6:  uint16(5),
		7:  uint16(9),
		11: uint16(11),
		12: uint16(5),
		13: uint16(13),
		16: uint16(15),
		21: uint16(17),
		22: uint16(19),
		24: uint16(21),
		32: uint16(23),
		34: uint16(25),
		35: uint16(5),
		36: uint16(27),
		37: uint16(29),
		38: uint16(31),
		39: uint16(31),
		41: uint16(33),
		42: uint16(35),
		43: uint16(5),
		44: uint16(37),
		47: uint16(122),
		48: uint16(91),
		49: uint16(85),
		50: uint16(13),
		51: uint16(13),
		52: uint16(16),
		53: uint16(16),
		54: uint16(13),
		55: uint16(13),
		58: uint16(13),
		59: uint16(13),
		60: uint16(13),
		61: uint16(13),
		67: uint16(13),
		68: uint16(13),
		69: uint16(13),
		70: uint16(14),
		72: uint16(13),
		73: uint16(86),
		74: uint16(11),
	},
	8: {
		1:  uint16(3),
		2:  uint16(5),
		3:  uint16(7),
		4:  uint16(5),
		5:  uint16(5),
		6:  uint16(5),
		7:  uint16(9),
		11: uint16(11),
		12: uint16(5),
		13: uint16(13),
		16: uint16(15),
		21: uint16(17),
		22: uint16(19),
		24: uint16(21),
		32: uint16(23),
		34: uint16(25),
		35: uint16(5),
		36: uint16(27),
		37: uint16(29),
		38: uint16(31),
		39: uint16(31),
		41: uint16(33),
		42: uint16(35),
		43: uint16(5),
		44: uint16(37),
		47: uint16(128),
		48: uint16(91),
		49: uint16(85),
		50: uint16(13),
		51: uint16(13),
		52: uint16(16),
		53: uint16(16),
		54: uint16(13),
		55: uint16(13),
		58: uint16(13),
		59: uint16(13),
		60: uint16(13),
		61: uint16(13),
		67: uint16(13),
		68: uint16(13),
		69: uint16(13),
		70: uint16(14),
		72: uint16(13),
		73: uint16(86),
		74: uint16(11),
	},
	9: {
		1:  uint16(3),
		2:  uint16(5),
		3:  uint16(7),
		4:  uint16(5),
		5:  uint16(5),
		6:  uint16(5),
		7:  uint16(9),
		11: uint16(11),
		12: uint16(5),
		13: uint16(13),
		16: uint16(15),
		21: uint16(17),
		22: uint16(19),
		24: uint16(21),
		32: uint16(23),
		34: uint16(25),
		35: uint16(5),
		36: uint16(27),
		37: uint16(29),
		38: uint16(31),
		39: uint16(31),
		41: uint16(33),
		42: uint16(35),
		43: uint16(5),
		44: uint16(37),
		47: uint16(123),
		48: uint16(91),
		49: uint16(85),
		50: uint16(13),
		51: uint16(13),
		52: uint16(16),
		53: uint16(16),
		54: uint16(13),
		55: uint16(13),
		58: uint16(13),
		59: uint16(13),
		60: uint16(13),
		61: uint16(13),
		67: uint16(13),
		68: uint16(13),
		69: uint16(13),
		70: uint16(14),
		72: uint16(13),
		73: uint16(86),
		74: uint16(11),
	},
	10: {
		0:  uint16(39),
		1:  uint16(39),
		2:  uint16(5),
		3:  uint16(7),
		4:  uint16(5),
		5:  uint16(5),
		6:  uint16(5),
		7:  uint16(9),
		10: uint16(39),
		11: uint16(11),
		12: uint16(5),
		13: uint16(13),
		16: uint16(15),
		21: uint16(17),
		22: uint16(19),
		24: uint16(21),
		32: uint16(23),
		34: uint16(25),
		35: uint16(5),
		36: uint16(27),
		37: uint16(29),
		38: uint16(31),
		39: uint16(31),
		41: uint16(33),
		42: uint16(35),
		43: uint16(5),
		44: uint16(37),
		49: uint16(89),
		50: uint16(13),
		51: uint16(13),
		52: uint16(16),
		53: uint16(16),
		54: uint16(13),
		55: uint16(13),
		58: uint16(13),
		59: uint16(13),
		60: uint16(13),
		61: uint16(13),
		67: uint16(13),
		68: uint16(13),
		69: uint16(13),
		70: uint16(14),
		72: uint16(13),
		74: uint16(11),
	},
	11: {
		0:  uint16(41),
		1:  uint16(41),
		2:  uint16(5),
		3:  uint16(7),
		4:  uint16(5),
		5:  uint16(5),
		6:  uint16(5),
		7:  uint16(9),
		10: uint16(41),
		11: uint16(11),
		12: uint16(5),
		13: uint16(13),
		16: uint16(15),
		21: uint16(17),
		22: uint16(19),
		24: uint16(21),
		32: uint16(23),
		34: uint16(25),
		35: uint16(5),
		36: uint16(27),
		37: uint16(29),
		38: uint16(31),
		39: uint16(31),
		41: uint16(33),
		42: uint16(35),
		43: uint16(5),
		44: uint16(37),
		50: uint16(13),
		51: uint16(13),
		52: uint16(16),
		53: uint16(16),
		54: uint16(13),
		55: uint16(13),
		58: uint16(13),
		59: uint16(13),
		60: uint16(13),
		61: uint16(13),
		67: uint16(13),
		68: uint16(13),
		69: uint16(13),
		70: uint16(14),
		72: uint16(13),
		74: uint16(12),
	},
	12: {
		0:  uint16(43),
		1:  uint16(43),
		2:  uint16(45),
		3:  uint16(48),
		4:  uint16(45),
		5:  uint16(45),
		6:  uint16(45),
		7:  uint16(51),
		10: uint16(43),
		11: uint16(54),
		12: uint16(45),
		13: uint16(57),
		16: uint16(60),
		21: uint16(63),
		22: uint16(66),
		24: uint16(69),
		32: uint16(72),
		34: uint16(75),
		35: uint16(45),
		36: uint16(78),
		37: uint16(81),
		38: uint16(84),
		39: uint16(84),
		41: uint16(87),
		42: uint16(90),
		43: uint16(45),
		44: uint16(93),
		50: uint16(13),
		51: uint16(13),
		52: uint16(16),
		53: uint16(16),
		54: uint16(13),
		55: uint16(13),
		58: uint16(13),
		59: uint16(13),
		60: uint16(13),
		61: uint16(13),
		67: uint16(13),
		68: uint16(13),
		69: uint16(13),
		70: uint16(14),
		72: uint16(13),
		74: uint16(12),
	},
}

var ts_small_parse_table = [2635]uint16_t{
	0:    uint16(7),
	1:    uint16(100),
	2:    uint16(1),
	3:    uint16(anon_sym_STAR),
	4:    uint16(102),
	5:    uint16(1),
	6:    uint16(anon_sym_QMARK),
	7:    uint16(104),
	8:    uint16(1),
	9:    uint16(anon_sym_PLUS),
	10:   uint16(106),
	11:   uint16(1),
	12:   uint16(anon_sym_LBRACE),
	13:   uint16(57),
	14:   uint16(4),
	15:   uint16(sym_zero_or_more),
	16:   uint16(sym_one_or_more),
	17:   uint16(sym_optional),
	18:   uint16(sym_count_quantifier),
	19:   uint16(98),
	20:   uint16(6),
	21:   uint16(anon_sym_LPAREN_QMARK),
	22:   uint16(sym_pattern_character),
	23:   uint16(anon_sym_LBRACK),
	24:   uint16(anon_sym_LPAREN),
	25:   uint16(aux_sym_control_escape_token1),
	26:   uint16(sym_identity_escape),
	27:   uint16(96),
	28:   uint16(21),
	30:   uint16(anon_sym_PIPE),
	31:   uint16(sym_any_character),
	32:   uint16(anon_sym_CARET),
	33:   uint16(sym_end_assertion),
	34:   uint16(sym_boundary_assertion),
	35:   uint16(sym_non_boundary_assertion),
	36:   uint16(anon_sym_RPAREN),
	37:   uint16(anon_sym_LPAREN_QMARK_LT),
	38:   uint16(anon_sym_LBRACK_COLON),
	39:   uint16(anon_sym_LPAREN_QMARKP_LT),
	40:   uint16(anon_sym_LPAREN_QMARK_COLON),
	41:   uint16(anon_sym_BSLASHk),
	42:   uint16(anon_sym_LPAREN_QMARKP_EQ),
	43:   uint16(sym_decimal_escape),
	44:   uint16(aux_sym_character_class_escape_token1),
	45:   uint16(aux_sym_character_class_escape_token2),
	46:   uint16(aux_sym_unicode_character_escape_token1),
	47:   uint16(aux_sym_unicode_character_escape_token2),
	48:   uint16(aux_sym_control_escape_token2),
	49:   uint16(sym_control_letter_escape),
	50:   uint16(2),
	51:   uint16(110),
	52:   uint16(6),
	53:   uint16(anon_sym_LPAREN_QMARK),
	54:   uint16(sym_pattern_character),
	55:   uint16(anon_sym_LBRACK),
	56:   uint16(anon_sym_LPAREN),
	57:   uint16(aux_sym_control_escape_token1),
	58:   uint16(sym_identity_escape),
	59:   uint16(108),
	60:   uint16(25),
	62:   uint16(anon_sym_PIPE),
	63:   uint16(sym_any_character),
	64:   uint16(anon_sym_CARET),
	65:   uint16(sym_end_assertion),
	66:   uint16(sym_boundary_assertion),
	67:   uint16(sym_non_boundary_assertion),
	68:   uint16(anon_sym_RPAREN),
	69:   uint16(anon_sym_LPAREN_QMARK_LT),
	70:   uint16(anon_sym_LBRACK_COLON),
	71:   uint16(anon_sym_LPAREN_QMARKP_LT),
	72:   uint16(anon_sym_LPAREN_QMARK_COLON),
	73:   uint16(anon_sym_STAR),
	74:   uint16(anon_sym_QMARK),
	75:   uint16(anon_sym_PLUS),
	76:   uint16(anon_sym_LBRACE),
	77:   uint16(anon_sym_BSLASHk),
	78:   uint16(anon_sym_LPAREN_QMARKP_EQ),
	79:   uint16(sym_decimal_escape),
	80:   uint16(aux_sym_character_class_escape_token1),
	81:   uint16(aux_sym_character_class_escape_token2),
	82:   uint16(aux_sym_unicode_character_escape_token1),
	83:   uint16(aux_sym_unicode_character_escape_token2),
	84:   uint16(aux_sym_control_escape_token2),
	85:   uint16(sym_control_letter_escape),
	86:   uint16(2),
	87:   uint16(114),
	88:   uint16(6),
	89:   uint16(anon_sym_LPAREN_QMARK),
	90:   uint16(sym_pattern_character),
	91:   uint16(anon_sym_LBRACK),
	92:   uint16(anon_sym_LPAREN),
	93:   uint16(aux_sym_control_escape_token1),
	94:   uint16(sym_identity_escape),
	95:   uint16(112),
	96:   uint16(25),
	98:   uint16(anon_sym_PIPE),
	99:   uint16(sym_any_character),
	100:  uint16(anon_sym_CARET),
	101:  uint16(sym_end_assertion),
	102:  uint16(sym_boundary_assertion),
	103:  uint16(sym_non_boundary_assertion),
	104:  uint16(anon_sym_RPAREN),
	105:  uint16(anon_sym_LPAREN_QMARK_LT),
	106:  uint16(anon_sym_LBRACK_COLON),
	107:  uint16(anon_sym_LPAREN_QMARKP_LT),
	108:  uint16(anon_sym_LPAREN_QMARK_COLON),
	109:  uint16(anon_sym_STAR),
	110:  uint16(anon_sym_QMARK),
	111:  uint16(anon_sym_PLUS),
	112:  uint16(anon_sym_LBRACE),
	113:  uint16(anon_sym_BSLASHk),
	114:  uint16(anon_sym_LPAREN_QMARKP_EQ),
	115:  uint16(sym_decimal_escape),
	116:  uint16(aux_sym_character_class_escape_token1),
	117:  uint16(aux_sym_character_class_escape_token2),
	118:  uint16(aux_sym_unicode_character_escape_token1),
	119:  uint16(aux_sym_unicode_character_escape_token2),
	120:  uint16(aux_sym_control_escape_token2),
	121:  uint16(sym_control_letter_escape),
	122:  uint16(2),
	123:  uint16(118),
	124:  uint16(6),
	125:  uint16(anon_sym_LPAREN_QMARK),
	126:  uint16(sym_pattern_character),
	127:  uint16(anon_sym_LBRACK),
	128:  uint16(anon_sym_LPAREN),
	129:  uint16(aux_sym_control_escape_token1),
	130:  uint16(sym_identity_escape),
	131:  uint16(116),
	132:  uint16(25),
	134:  uint16(anon_sym_PIPE),
	135:  uint16(sym_any_character),
	136:  uint16(anon_sym_CARET),
	137:  uint16(sym_end_assertion),
	138:  uint16(sym_boundary_assertion),
	139:  uint16(sym_non_boundary_assertion),
	140:  uint16(anon_sym_RPAREN),
	141:  uint16(anon_sym_LPAREN_QMARK_LT),
	142:  uint16(anon_sym_LBRACK_COLON),
	143:  uint16(anon_sym_LPAREN_QMARKP_LT),
	144:  uint16(anon_sym_LPAREN_QMARK_COLON),
	145:  uint16(anon_sym_STAR),
	146:  uint16(anon_sym_QMARK),
	147:  uint16(anon_sym_PLUS),
	148:  uint16(anon_sym_LBRACE),
	149:  uint16(anon_sym_BSLASHk),
	150:  uint16(anon_sym_LPAREN_QMARKP_EQ),
	151:  uint16(sym_decimal_escape),
	152:  uint16(aux_sym_character_class_escape_token1),
	153:  uint16(aux_sym_character_class_escape_token2),
	154:  uint16(aux_sym_unicode_character_escape_token1),
	155:  uint16(aux_sym_unicode_character_escape_token2),
	156:  uint16(aux_sym_control_escape_token2),
	157:  uint16(sym_control_letter_escape),
	158:  uint16(2),
	159:  uint16(122),
	160:  uint16(6),
	161:  uint16(anon_sym_LPAREN_QMARK),
	162:  uint16(sym_pattern_character),
	163:  uint16(anon_sym_LBRACK),
	164:  uint16(anon_sym_LPAREN),
	165:  uint16(aux_sym_control_escape_token1),
	166:  uint16(sym_identity_escape),
	167:  uint16(120),
	168:  uint16(25),
	170:  uint16(anon_sym_PIPE),
	171:  uint16(sym_any_character),
	172:  uint16(anon_sym_CARET),
	173:  uint16(sym_end_assertion),
	174:  uint16(sym_boundary_assertion),
	175:  uint16(sym_non_boundary_assertion),
	176:  uint16(anon_sym_RPAREN),
	177:  uint16(anon_sym_LPAREN_QMARK_LT),
	178:  uint16(anon_sym_LBRACK_COLON),
	179:  uint16(anon_sym_LPAREN_QMARKP_LT),
	180:  uint16(anon_sym_LPAREN_QMARK_COLON),
	181:  uint16(anon_sym_STAR),
	182:  uint16(anon_sym_QMARK),
	183:  uint16(anon_sym_PLUS),
	184:  uint16(anon_sym_LBRACE),
	185:  uint16(anon_sym_BSLASHk),
	186:  uint16(anon_sym_LPAREN_QMARKP_EQ),
	187:  uint16(sym_decimal_escape),
	188:  uint16(aux_sym_character_class_escape_token1),
	189:  uint16(aux_sym_character_class_escape_token2),
	190:  uint16(aux_sym_unicode_character_escape_token1),
	191:  uint16(aux_sym_unicode_character_escape_token2),
	192:  uint16(aux_sym_control_escape_token2),
	193:  uint16(sym_control_letter_escape),
	194:  uint16(2),
	195:  uint16(126),
	196:  uint16(6),
	197:  uint16(anon_sym_LPAREN_QMARK),
	198:  uint16(sym_pattern_character),
	199:  uint16(anon_sym_LBRACK),
	200:  uint16(anon_sym_LPAREN),
	201:  uint16(aux_sym_control_escape_token1),
	202:  uint16(sym_identity_escape),
	203:  uint16(124),
	204:  uint16(25),
	206:  uint16(anon_sym_PIPE),
	207:  uint16(sym_any_character),
	208:  uint16(anon_sym_CARET),
	209:  uint16(sym_end_assertion),
	210:  uint16(sym_boundary_assertion),
	211:  uint16(sym_non_boundary_assertion),
	212:  uint16(anon_sym_RPAREN),
	213:  uint16(anon_sym_LPAREN_QMARK_LT),
	214:  uint16(anon_sym_LBRACK_COLON),
	215:  uint16(anon_sym_LPAREN_QMARKP_LT),
	216:  uint16(anon_sym_LPAREN_QMARK_COLON),
	217:  uint16(anon_sym_STAR),
	218:  uint16(anon_sym_QMARK),
	219:  uint16(anon_sym_PLUS),
	220:  uint16(anon_sym_LBRACE),
	221:  uint16(anon_sym_BSLASHk),
	222:  uint16(anon_sym_LPAREN_QMARKP_EQ),
	223:  uint16(sym_decimal_escape),
	224:  uint16(aux_sym_character_class_escape_token1),
	225:  uint16(aux_sym_character_class_escape_token2),
	226:  uint16(aux_sym_unicode_character_escape_token1),
	227:  uint16(aux_sym_unicode_character_escape_token2),
	228:  uint16(aux_sym_control_escape_token2),
	229:  uint16(sym_control_letter_escape),
	230:  uint16(2),
	231:  uint16(130),
	232:  uint16(6),
	233:  uint16(anon_sym_LPAREN_QMARK),
	234:  uint16(sym_pattern_character),
	235:  uint16(anon_sym_LBRACK),
	236:  uint16(anon_sym_LPAREN),
	237:  uint16(aux_sym_control_escape_token1),
	238:  uint16(sym_identity_escape),
	239:  uint16(128),
	240:  uint16(25),
	242:  uint16(anon_sym_PIPE),
	243:  uint16(sym_any_character),
	244:  uint16(anon_sym_CARET),
	245:  uint16(sym_end_assertion),
	246:  uint16(sym_boundary_assertion),
	247:  uint16(sym_non_boundary_assertion),
	248:  uint16(anon_sym_RPAREN),
	249:  uint16(anon_sym_LPAREN_QMARK_LT),
	250:  uint16(anon_sym_LBRACK_COLON),
	251:  uint16(anon_sym_LPAREN_QMARKP_LT),
	252:  uint16(anon_sym_LPAREN_QMARK_COLON),
	253:  uint16(anon_sym_STAR),
	254:  uint16(anon_sym_QMARK),
	255:  uint16(anon_sym_PLUS),
	256:  uint16(anon_sym_LBRACE),
	257:  uint16(anon_sym_BSLASHk),
	258:  uint16(anon_sym_LPAREN_QMARKP_EQ),
	259:  uint16(sym_decimal_escape),
	260:  uint16(aux_sym_character_class_escape_token1),
	261:  uint16(aux_sym_character_class_escape_token2),
	262:  uint16(aux_sym_unicode_character_escape_token1),
	263:  uint16(aux_sym_unicode_character_escape_token2),
	264:  uint16(aux_sym_control_escape_token2),
	265:  uint16(sym_control_letter_escape),
	266:  uint16(2),
	267:  uint16(134),
	268:  uint16(6),
	269:  uint16(anon_sym_LPAREN_QMARK),
	270:  uint16(sym_pattern_character),
	271:  uint16(anon_sym_LBRACK),
	272:  uint16(anon_sym_LPAREN),
	273:  uint16(aux_sym_control_escape_token1),
	274:  uint16(sym_identity_escape),
	275:  uint16(132),
	276:  uint16(25),
	278:  uint16(anon_sym_PIPE),
	279:  uint16(sym_any_character),
	280:  uint16(anon_sym_CARET),
	281:  uint16(sym_end_assertion),
	282:  uint16(sym_boundary_assertion),
	283:  uint16(sym_non_boundary_assertion),
	284:  uint16(anon_sym_RPAREN),
	285:  uint16(anon_sym_LPAREN_QMARK_LT),
	286:  uint16(anon_sym_LBRACK_COLON),
	287:  uint16(anon_sym_LPAREN_QMARKP_LT),
	288:  uint16(anon_sym_LPAREN_QMARK_COLON),
	289:  uint16(anon_sym_STAR),
	290:  uint16(anon_sym_QMARK),
	291:  uint16(anon_sym_PLUS),
	292:  uint16(anon_sym_LBRACE),
	293:  uint16(anon_sym_BSLASHk),
	294:  uint16(anon_sym_LPAREN_QMARKP_EQ),
	295:  uint16(sym_decimal_escape),
	296:  uint16(aux_sym_character_class_escape_token1),
	297:  uint16(aux_sym_character_class_escape_token2),
	298:  uint16(aux_sym_unicode_character_escape_token1),
	299:  uint16(aux_sym_unicode_character_escape_token2),
	300:  uint16(aux_sym_control_escape_token2),
	301:  uint16(sym_control_letter_escape),
	302:  uint16(2),
	303:  uint16(138),
	304:  uint16(6),
	305:  uint16(anon_sym_LPAREN_QMARK),
	306:  uint16(sym_pattern_character),
	307:  uint16(anon_sym_LBRACK),
	308:  uint16(anon_sym_LPAREN),
	309:  uint16(aux_sym_control_escape_token1),
	310:  uint16(sym_identity_escape),
	311:  uint16(136),
	312:  uint16(25),
	314:  uint16(anon_sym_PIPE),
	315:  uint16(sym_any_character),
	316:  uint16(anon_sym_CARET),
	317:  uint16(sym_end_assertion),
	318:  uint16(sym_boundary_assertion),
	319:  uint16(sym_non_boundary_assertion),
	320:  uint16(anon_sym_RPAREN),
	321:  uint16(anon_sym_LPAREN_QMARK_LT),
	322:  uint16(anon_sym_LBRACK_COLON),
	323:  uint16(anon_sym_LPAREN_QMARKP_LT),
	324:  uint16(anon_sym_LPAREN_QMARK_COLON),
	325:  uint16(anon_sym_STAR),
	326:  uint16(anon_sym_QMARK),
	327:  uint16(anon_sym_PLUS),
	328:  uint16(anon_sym_LBRACE),
	329:  uint16(anon_sym_BSLASHk),
	330:  uint16(anon_sym_LPAREN_QMARKP_EQ),
	331:  uint16(sym_decimal_escape),
	332:  uint16(aux_sym_character_class_escape_token1),
	333:  uint16(aux_sym_character_class_escape_token2),
	334:  uint16(aux_sym_unicode_character_escape_token1),
	335:  uint16(aux_sym_unicode_character_escape_token2),
	336:  uint16(aux_sym_control_escape_token2),
	337:  uint16(sym_control_letter_escape),
	338:  uint16(2),
	339:  uint16(142),
	340:  uint16(6),
	341:  uint16(anon_sym_LPAREN_QMARK),
	342:  uint16(sym_pattern_character),
	343:  uint16(anon_sym_LBRACK),
	344:  uint16(anon_sym_LPAREN),
	345:  uint16(aux_sym_control_escape_token1),
	346:  uint16(sym_identity_escape),
	347:  uint16(140),
	348:  uint16(25),
	350:  uint16(anon_sym_PIPE),
	351:  uint16(sym_any_character),
	352:  uint16(anon_sym_CARET),
	353:  uint16(sym_end_assertion),
	354:  uint16(sym_boundary_assertion),
	355:  uint16(sym_non_boundary_assertion),
	356:  uint16(anon_sym_RPAREN),
	357:  uint16(anon_sym_LPAREN_QMARK_LT),
	358:  uint16(anon_sym_LBRACK_COLON),
	359:  uint16(anon_sym_LPAREN_QMARKP_LT),
	360:  uint16(anon_sym_LPAREN_QMARK_COLON),
	361:  uint16(anon_sym_STAR),
	362:  uint16(anon_sym_QMARK),
	363:  uint16(anon_sym_PLUS),
	364:  uint16(anon_sym_LBRACE),
	365:  uint16(anon_sym_BSLASHk),
	366:  uint16(anon_sym_LPAREN_QMARKP_EQ),
	367:  uint16(sym_decimal_escape),
	368:  uint16(aux_sym_character_class_escape_token1),
	369:  uint16(aux_sym_character_class_escape_token2),
	370:  uint16(aux_sym_unicode_character_escape_token1),
	371:  uint16(aux_sym_unicode_character_escape_token2),
	372:  uint16(aux_sym_control_escape_token2),
	373:  uint16(sym_control_letter_escape),
	374:  uint16(2),
	375:  uint16(146),
	376:  uint16(6),
	377:  uint16(anon_sym_LPAREN_QMARK),
	378:  uint16(sym_pattern_character),
	379:  uint16(anon_sym_LBRACK),
	380:  uint16(anon_sym_LPAREN),
	381:  uint16(aux_sym_control_escape_token1),
	382:  uint16(sym_identity_escape),
	383:  uint16(144),
	384:  uint16(25),
	386:  uint16(anon_sym_PIPE),
	387:  uint16(sym_any_character),
	388:  uint16(anon_sym_CARET),
	389:  uint16(sym_end_assertion),
	390:  uint16(sym_boundary_assertion),
	391:  uint16(sym_non_boundary_assertion),
	392:  uint16(anon_sym_RPAREN),
	393:  uint16(anon_sym_LPAREN_QMARK_LT),
	394:  uint16(anon_sym_LBRACK_COLON),
	395:  uint16(anon_sym_LPAREN_QMARKP_LT),
	396:  uint16(anon_sym_LPAREN_QMARK_COLON),
	397:  uint16(anon_sym_STAR),
	398:  uint16(anon_sym_QMARK),
	399:  uint16(anon_sym_PLUS),
	400:  uint16(anon_sym_LBRACE),
	401:  uint16(anon_sym_BSLASHk),
	402:  uint16(anon_sym_LPAREN_QMARKP_EQ),
	403:  uint16(sym_decimal_escape),
	404:  uint16(aux_sym_character_class_escape_token1),
	405:  uint16(aux_sym_character_class_escape_token2),
	406:  uint16(aux_sym_unicode_character_escape_token1),
	407:  uint16(aux_sym_unicode_character_escape_token2),
	408:  uint16(aux_sym_control_escape_token2),
	409:  uint16(sym_control_letter_escape),
	410:  uint16(2),
	411:  uint16(150),
	412:  uint16(6),
	413:  uint16(anon_sym_LPAREN_QMARK),
	414:  uint16(sym_pattern_character),
	415:  uint16(anon_sym_LBRACK),
	416:  uint16(anon_sym_LPAREN),
	417:  uint16(aux_sym_control_escape_token1),
	418:  uint16(sym_identity_escape),
	419:  uint16(148),
	420:  uint16(25),
	422:  uint16(anon_sym_PIPE),
	423:  uint16(sym_any_character),
	424:  uint16(anon_sym_CARET),
	425:  uint16(sym_end_assertion),
	426:  uint16(sym_boundary_assertion),
	427:  uint16(sym_non_boundary_assertion),
	428:  uint16(anon_sym_RPAREN),
	429:  uint16(anon_sym_LPAREN_QMARK_LT),
	430:  uint16(anon_sym_LBRACK_COLON),
	431:  uint16(anon_sym_LPAREN_QMARKP_LT),
	432:  uint16(anon_sym_LPAREN_QMARK_COLON),
	433:  uint16(anon_sym_STAR),
	434:  uint16(anon_sym_QMARK),
	435:  uint16(anon_sym_PLUS),
	436:  uint16(anon_sym_LBRACE),
	437:  uint16(anon_sym_BSLASHk),
	438:  uint16(anon_sym_LPAREN_QMARKP_EQ),
	439:  uint16(sym_decimal_escape),
	440:  uint16(aux_sym_character_class_escape_token1),
	441:  uint16(aux_sym_character_class_escape_token2),
	442:  uint16(aux_sym_unicode_character_escape_token1),
	443:  uint16(aux_sym_unicode_character_escape_token2),
	444:  uint16(aux_sym_control_escape_token2),
	445:  uint16(sym_control_letter_escape),
	446:  uint16(2),
	447:  uint16(154),
	448:  uint16(6),
	449:  uint16(anon_sym_LPAREN_QMARK),
	450:  uint16(sym_pattern_character),
	451:  uint16(anon_sym_LBRACK),
	452:  uint16(anon_sym_LPAREN),
	453:  uint16(aux_sym_control_escape_token1),
	454:  uint16(sym_identity_escape),
	455:  uint16(152),
	456:  uint16(25),
	458:  uint16(anon_sym_PIPE),
	459:  uint16(sym_any_character),
	460:  uint16(anon_sym_CARET),
	461:  uint16(sym_end_assertion),
	462:  uint16(sym_boundary_assertion),
	463:  uint16(sym_non_boundary_assertion),
	464:  uint16(anon_sym_RPAREN),
	465:  uint16(anon_sym_LPAREN_QMARK_LT),
	466:  uint16(anon_sym_LBRACK_COLON),
	467:  uint16(anon_sym_LPAREN_QMARKP_LT),
	468:  uint16(anon_sym_LPAREN_QMARK_COLON),
	469:  uint16(anon_sym_STAR),
	470:  uint16(anon_sym_QMARK),
	471:  uint16(anon_sym_PLUS),
	472:  uint16(anon_sym_LBRACE),
	473:  uint16(anon_sym_BSLASHk),
	474:  uint16(anon_sym_LPAREN_QMARKP_EQ),
	475:  uint16(sym_decimal_escape),
	476:  uint16(aux_sym_character_class_escape_token1),
	477:  uint16(aux_sym_character_class_escape_token2),
	478:  uint16(aux_sym_unicode_character_escape_token1),
	479:  uint16(aux_sym_unicode_character_escape_token2),
	480:  uint16(aux_sym_control_escape_token2),
	481:  uint16(sym_control_letter_escape),
	482:  uint16(2),
	483:  uint16(158),
	484:  uint16(6),
	485:  uint16(anon_sym_LPAREN_QMARK),
	486:  uint16(sym_pattern_character),
	487:  uint16(anon_sym_LBRACK),
	488:  uint16(anon_sym_LPAREN),
	489:  uint16(aux_sym_control_escape_token1),
	490:  uint16(sym_identity_escape),
	491:  uint16(156),
	492:  uint16(25),
	494:  uint16(anon_sym_PIPE),
	495:  uint16(sym_any_character),
	496:  uint16(anon_sym_CARET),
	497:  uint16(sym_end_assertion),
	498:  uint16(sym_boundary_assertion),
	499:  uint16(sym_non_boundary_assertion),
	500:  uint16(anon_sym_RPAREN),
	501:  uint16(anon_sym_LPAREN_QMARK_LT),
	502:  uint16(anon_sym_LBRACK_COLON),
	503:  uint16(anon_sym_LPAREN_QMARKP_LT),
	504:  uint16(anon_sym_LPAREN_QMARK_COLON),
	505:  uint16(anon_sym_STAR),
	506:  uint16(anon_sym_QMARK),
	507:  uint16(anon_sym_PLUS),
	508:  uint16(anon_sym_LBRACE),
	509:  uint16(anon_sym_BSLASHk),
	510:  uint16(anon_sym_LPAREN_QMARKP_EQ),
	511:  uint16(sym_decimal_escape),
	512:  uint16(aux_sym_character_class_escape_token1),
	513:  uint16(aux_sym_character_class_escape_token2),
	514:  uint16(aux_sym_unicode_character_escape_token1),
	515:  uint16(aux_sym_unicode_character_escape_token2),
	516:  uint16(aux_sym_control_escape_token2),
	517:  uint16(sym_control_letter_escape),
	518:  uint16(2),
	519:  uint16(162),
	520:  uint16(6),
	521:  uint16(anon_sym_LPAREN_QMARK),
	522:  uint16(sym_pattern_character),
	523:  uint16(anon_sym_LBRACK),
	524:  uint16(anon_sym_LPAREN),
	525:  uint16(aux_sym_control_escape_token1),
	526:  uint16(sym_identity_escape),
	527:  uint16(160),
	528:  uint16(25),
	530:  uint16(anon_sym_PIPE),
	531:  uint16(sym_any_character),
	532:  uint16(anon_sym_CARET),
	533:  uint16(sym_end_assertion),
	534:  uint16(sym_boundary_assertion),
	535:  uint16(sym_non_boundary_assertion),
	536:  uint16(anon_sym_RPAREN),
	537:  uint16(anon_sym_LPAREN_QMARK_LT),
	538:  uint16(anon_sym_LBRACK_COLON),
	539:  uint16(anon_sym_LPAREN_QMARKP_LT),
	540:  uint16(anon_sym_LPAREN_QMARK_COLON),
	541:  uint16(anon_sym_STAR),
	542:  uint16(anon_sym_QMARK),
	543:  uint16(anon_sym_PLUS),
	544:  uint16(anon_sym_LBRACE),
	545:  uint16(anon_sym_BSLASHk),
	546:  uint16(anon_sym_LPAREN_QMARKP_EQ),
	547:  uint16(sym_decimal_escape),
	548:  uint16(aux_sym_character_class_escape_token1),
	549:  uint16(aux_sym_character_class_escape_token2),
	550:  uint16(aux_sym_unicode_character_escape_token1),
	551:  uint16(aux_sym_unicode_character_escape_token2),
	552:  uint16(aux_sym_control_escape_token2),
	553:  uint16(sym_control_letter_escape),
	554:  uint16(2),
	555:  uint16(166),
	556:  uint16(6),
	557:  uint16(anon_sym_LPAREN_QMARK),
	558:  uint16(sym_pattern_character),
	559:  uint16(anon_sym_LBRACK),
	560:  uint16(anon_sym_LPAREN),
	561:  uint16(aux_sym_control_escape_token1),
	562:  uint16(sym_identity_escape),
	563:  uint16(164),
	564:  uint16(25),
	566:  uint16(anon_sym_PIPE),
	567:  uint16(sym_any_character),
	568:  uint16(anon_sym_CARET),
	569:  uint16(sym_end_assertion),
	570:  uint16(sym_boundary_assertion),
	571:  uint16(sym_non_boundary_assertion),
	572:  uint16(anon_sym_RPAREN),
	573:  uint16(anon_sym_LPAREN_QMARK_LT),
	574:  uint16(anon_sym_LBRACK_COLON),
	575:  uint16(anon_sym_LPAREN_QMARKP_LT),
	576:  uint16(anon_sym_LPAREN_QMARK_COLON),
	577:  uint16(anon_sym_STAR),
	578:  uint16(anon_sym_QMARK),
	579:  uint16(anon_sym_PLUS),
	580:  uint16(anon_sym_LBRACE),
	581:  uint16(anon_sym_BSLASHk),
	582:  uint16(anon_sym_LPAREN_QMARKP_EQ),
	583:  uint16(sym_decimal_escape),
	584:  uint16(aux_sym_character_class_escape_token1),
	585:  uint16(aux_sym_character_class_escape_token2),
	586:  uint16(aux_sym_unicode_character_escape_token1),
	587:  uint16(aux_sym_unicode_character_escape_token2),
	588:  uint16(aux_sym_control_escape_token2),
	589:  uint16(sym_control_letter_escape),
	590:  uint16(2),
	591:  uint16(170),
	592:  uint16(6),
	593:  uint16(anon_sym_LPAREN_QMARK),
	594:  uint16(sym_pattern_character),
	595:  uint16(anon_sym_LBRACK),
	596:  uint16(anon_sym_LPAREN),
	597:  uint16(aux_sym_control_escape_token1),
	598:  uint16(sym_identity_escape),
	599:  uint16(168),
	600:  uint16(25),
	602:  uint16(anon_sym_PIPE),
	603:  uint16(sym_any_character),
	604:  uint16(anon_sym_CARET),
	605:  uint16(sym_end_assertion),
	606:  uint16(sym_boundary_assertion),
	607:  uint16(sym_non_boundary_assertion),
	608:  uint16(anon_sym_RPAREN),
	609:  uint16(anon_sym_LPAREN_QMARK_LT),
	610:  uint16(anon_sym_LBRACK_COLON),
	611:  uint16(anon_sym_LPAREN_QMARKP_LT),
	612:  uint16(anon_sym_LPAREN_QMARK_COLON),
	613:  uint16(anon_sym_STAR),
	614:  uint16(anon_sym_QMARK),
	615:  uint16(anon_sym_PLUS),
	616:  uint16(anon_sym_LBRACE),
	617:  uint16(anon_sym_BSLASHk),
	618:  uint16(anon_sym_LPAREN_QMARKP_EQ),
	619:  uint16(sym_decimal_escape),
	620:  uint16(aux_sym_character_class_escape_token1),
	621:  uint16(aux_sym_character_class_escape_token2),
	622:  uint16(aux_sym_unicode_character_escape_token1),
	623:  uint16(aux_sym_unicode_character_escape_token2),
	624:  uint16(aux_sym_control_escape_token2),
	625:  uint16(sym_control_letter_escape),
	626:  uint16(2),
	627:  uint16(174),
	628:  uint16(6),
	629:  uint16(anon_sym_LPAREN_QMARK),
	630:  uint16(sym_pattern_character),
	631:  uint16(anon_sym_LBRACK),
	632:  uint16(anon_sym_LPAREN),
	633:  uint16(aux_sym_control_escape_token1),
	634:  uint16(sym_identity_escape),
	635:  uint16(172),
	636:  uint16(25),
	638:  uint16(anon_sym_PIPE),
	639:  uint16(sym_any_character),
	640:  uint16(anon_sym_CARET),
	641:  uint16(sym_end_assertion),
	642:  uint16(sym_boundary_assertion),
	643:  uint16(sym_non_boundary_assertion),
	644:  uint16(anon_sym_RPAREN),
	645:  uint16(anon_sym_LPAREN_QMARK_LT),
	646:  uint16(anon_sym_LBRACK_COLON),
	647:  uint16(anon_sym_LPAREN_QMARKP_LT),
	648:  uint16(anon_sym_LPAREN_QMARK_COLON),
	649:  uint16(anon_sym_STAR),
	650:  uint16(anon_sym_QMARK),
	651:  uint16(anon_sym_PLUS),
	652:  uint16(anon_sym_LBRACE),
	653:  uint16(anon_sym_BSLASHk),
	654:  uint16(anon_sym_LPAREN_QMARKP_EQ),
	655:  uint16(sym_decimal_escape),
	656:  uint16(aux_sym_character_class_escape_token1),
	657:  uint16(aux_sym_character_class_escape_token2),
	658:  uint16(aux_sym_unicode_character_escape_token1),
	659:  uint16(aux_sym_unicode_character_escape_token2),
	660:  uint16(aux_sym_control_escape_token2),
	661:  uint16(sym_control_letter_escape),
	662:  uint16(2),
	663:  uint16(178),
	664:  uint16(6),
	665:  uint16(anon_sym_LPAREN_QMARK),
	666:  uint16(sym_pattern_character),
	667:  uint16(anon_sym_LBRACK),
	668:  uint16(anon_sym_LPAREN),
	669:  uint16(aux_sym_control_escape_token1),
	670:  uint16(sym_identity_escape),
	671:  uint16(176),
	672:  uint16(25),
	674:  uint16(anon_sym_PIPE),
	675:  uint16(sym_any_character),
	676:  uint16(anon_sym_CARET),
	677:  uint16(sym_end_assertion),
	678:  uint16(sym_boundary_assertion),
	679:  uint16(sym_non_boundary_assertion),
	680:  uint16(anon_sym_RPAREN),
	681:  uint16(anon_sym_LPAREN_QMARK_LT),
	682:  uint16(anon_sym_LBRACK_COLON),
	683:  uint16(anon_sym_LPAREN_QMARKP_LT),
	684:  uint16(anon_sym_LPAREN_QMARK_COLON),
	685:  uint16(anon_sym_STAR),
	686:  uint16(anon_sym_QMARK),
	687:  uint16(anon_sym_PLUS),
	688:  uint16(anon_sym_LBRACE),
	689:  uint16(anon_sym_BSLASHk),
	690:  uint16(anon_sym_LPAREN_QMARKP_EQ),
	691:  uint16(sym_decimal_escape),
	692:  uint16(aux_sym_character_class_escape_token1),
	693:  uint16(aux_sym_character_class_escape_token2),
	694:  uint16(aux_sym_unicode_character_escape_token1),
	695:  uint16(aux_sym_unicode_character_escape_token2),
	696:  uint16(aux_sym_control_escape_token2),
	697:  uint16(sym_control_letter_escape),
	698:  uint16(2),
	699:  uint16(182),
	700:  uint16(6),
	701:  uint16(anon_sym_LPAREN_QMARK),
	702:  uint16(sym_pattern_character),
	703:  uint16(anon_sym_LBRACK),
	704:  uint16(anon_sym_LPAREN),
	705:  uint16(aux_sym_control_escape_token1),
	706:  uint16(sym_identity_escape),
	707:  uint16(180),
	708:  uint16(25),
	710:  uint16(anon_sym_PIPE),
	711:  uint16(sym_any_character),
	712:  uint16(anon_sym_CARET),
	713:  uint16(sym_end_assertion),
	714:  uint16(sym_boundary_assertion),
	715:  uint16(sym_non_boundary_assertion),
	716:  uint16(anon_sym_RPAREN),
	717:  uint16(anon_sym_LPAREN_QMARK_LT),
	718:  uint16(anon_sym_LBRACK_COLON),
	719:  uint16(anon_sym_LPAREN_QMARKP_LT),
	720:  uint16(anon_sym_LPAREN_QMARK_COLON),
	721:  uint16(anon_sym_STAR),
	722:  uint16(anon_sym_QMARK),
	723:  uint16(anon_sym_PLUS),
	724:  uint16(anon_sym_LBRACE),
	725:  uint16(anon_sym_BSLASHk),
	726:  uint16(anon_sym_LPAREN_QMARKP_EQ),
	727:  uint16(sym_decimal_escape),
	728:  uint16(aux_sym_character_class_escape_token1),
	729:  uint16(aux_sym_character_class_escape_token2),
	730:  uint16(aux_sym_unicode_character_escape_token1),
	731:  uint16(aux_sym_unicode_character_escape_token2),
	732:  uint16(aux_sym_control_escape_token2),
	733:  uint16(sym_control_letter_escape),
	734:  uint16(2),
	735:  uint16(186),
	736:  uint16(6),
	737:  uint16(anon_sym_LPAREN_QMARK),
	738:  uint16(sym_pattern_character),
	739:  uint16(anon_sym_LBRACK),
	740:  uint16(anon_sym_LPAREN),
	741:  uint16(aux_sym_control_escape_token1),
	742:  uint16(sym_identity_escape),
	743:  uint16(184),
	744:  uint16(25),
	746:  uint16(anon_sym_PIPE),
	747:  uint16(sym_any_character),
	748:  uint16(anon_sym_CARET),
	749:  uint16(sym_end_assertion),
	750:  uint16(sym_boundary_assertion),
	751:  uint16(sym_non_boundary_assertion),
	752:  uint16(anon_sym_RPAREN),
	753:  uint16(anon_sym_LPAREN_QMARK_LT),
	754:  uint16(anon_sym_LBRACK_COLON),
	755:  uint16(anon_sym_LPAREN_QMARKP_LT),
	756:  uint16(anon_sym_LPAREN_QMARK_COLON),
	757:  uint16(anon_sym_STAR),
	758:  uint16(anon_sym_QMARK),
	759:  uint16(anon_sym_PLUS),
	760:  uint16(anon_sym_LBRACE),
	761:  uint16(anon_sym_BSLASHk),
	762:  uint16(anon_sym_LPAREN_QMARKP_EQ),
	763:  uint16(sym_decimal_escape),
	764:  uint16(aux_sym_character_class_escape_token1),
	765:  uint16(aux_sym_character_class_escape_token2),
	766:  uint16(aux_sym_unicode_character_escape_token1),
	767:  uint16(aux_sym_unicode_character_escape_token2),
	768:  uint16(aux_sym_control_escape_token2),
	769:  uint16(sym_control_letter_escape),
	770:  uint16(2),
	771:  uint16(190),
	772:  uint16(6),
	773:  uint16(anon_sym_LPAREN_QMARK),
	774:  uint16(sym_pattern_character),
	775:  uint16(anon_sym_LBRACK),
	776:  uint16(anon_sym_LPAREN),
	777:  uint16(aux_sym_control_escape_token1),
	778:  uint16(sym_identity_escape),
	779:  uint16(188),
	780:  uint16(25),
	782:  uint16(anon_sym_PIPE),
	783:  uint16(sym_any_character),
	784:  uint16(anon_sym_CARET),
	785:  uint16(sym_end_assertion),
	786:  uint16(sym_boundary_assertion),
	787:  uint16(sym_non_boundary_assertion),
	788:  uint16(anon_sym_RPAREN),
	789:  uint16(anon_sym_LPAREN_QMARK_LT),
	790:  uint16(anon_sym_LBRACK_COLON),
	791:  uint16(anon_sym_LPAREN_QMARKP_LT),
	792:  uint16(anon_sym_LPAREN_QMARK_COLON),
	793:  uint16(anon_sym_STAR),
	794:  uint16(anon_sym_QMARK),
	795:  uint16(anon_sym_PLUS),
	796:  uint16(anon_sym_LBRACE),
	797:  uint16(anon_sym_BSLASHk),
	798:  uint16(anon_sym_LPAREN_QMARKP_EQ),
	799:  uint16(sym_decimal_escape),
	800:  uint16(aux_sym_character_class_escape_token1),
	801:  uint16(aux_sym_character_class_escape_token2),
	802:  uint16(aux_sym_unicode_character_escape_token1),
	803:  uint16(aux_sym_unicode_character_escape_token2),
	804:  uint16(aux_sym_control_escape_token2),
	805:  uint16(sym_control_letter_escape),
	806:  uint16(2),
	807:  uint16(194),
	808:  uint16(6),
	809:  uint16(anon_sym_LPAREN_QMARK),
	810:  uint16(sym_pattern_character),
	811:  uint16(anon_sym_LBRACK),
	812:  uint16(anon_sym_LPAREN),
	813:  uint16(aux_sym_control_escape_token1),
	814:  uint16(sym_identity_escape),
	815:  uint16(192),
	816:  uint16(25),
	818:  uint16(anon_sym_PIPE),
	819:  uint16(sym_any_character),
	820:  uint16(anon_sym_CARET),
	821:  uint16(sym_end_assertion),
	822:  uint16(sym_boundary_assertion),
	823:  uint16(sym_non_boundary_assertion),
	824:  uint16(anon_sym_RPAREN),
	825:  uint16(anon_sym_LPAREN_QMARK_LT),
	826:  uint16(anon_sym_LBRACK_COLON),
	827:  uint16(anon_sym_LPAREN_QMARKP_LT),
	828:  uint16(anon_sym_LPAREN_QMARK_COLON),
	829:  uint16(anon_sym_STAR),
	830:  uint16(anon_sym_QMARK),
	831:  uint16(anon_sym_PLUS),
	832:  uint16(anon_sym_LBRACE),
	833:  uint16(anon_sym_BSLASHk),
	834:  uint16(anon_sym_LPAREN_QMARKP_EQ),
	835:  uint16(sym_decimal_escape),
	836:  uint16(aux_sym_character_class_escape_token1),
	837:  uint16(aux_sym_character_class_escape_token2),
	838:  uint16(aux_sym_unicode_character_escape_token1),
	839:  uint16(aux_sym_unicode_character_escape_token2),
	840:  uint16(aux_sym_control_escape_token2),
	841:  uint16(sym_control_letter_escape),
	842:  uint16(2),
	843:  uint16(198),
	844:  uint16(6),
	845:  uint16(anon_sym_LPAREN_QMARK),
	846:  uint16(sym_pattern_character),
	847:  uint16(anon_sym_LBRACK),
	848:  uint16(anon_sym_LPAREN),
	849:  uint16(aux_sym_control_escape_token1),
	850:  uint16(sym_identity_escape),
	851:  uint16(196),
	852:  uint16(25),
	854:  uint16(anon_sym_PIPE),
	855:  uint16(sym_any_character),
	856:  uint16(anon_sym_CARET),
	857:  uint16(sym_end_assertion),
	858:  uint16(sym_boundary_assertion),
	859:  uint16(sym_non_boundary_assertion),
	860:  uint16(anon_sym_RPAREN),
	861:  uint16(anon_sym_LPAREN_QMARK_LT),
	862:  uint16(anon_sym_LBRACK_COLON),
	863:  uint16(anon_sym_LPAREN_QMARKP_LT),
	864:  uint16(anon_sym_LPAREN_QMARK_COLON),
	865:  uint16(anon_sym_STAR),
	866:  uint16(anon_sym_QMARK),
	867:  uint16(anon_sym_PLUS),
	868:  uint16(anon_sym_LBRACE),
	869:  uint16(anon_sym_BSLASHk),
	870:  uint16(anon_sym_LPAREN_QMARKP_EQ),
	871:  uint16(sym_decimal_escape),
	872:  uint16(aux_sym_character_class_escape_token1),
	873:  uint16(aux_sym_character_class_escape_token2),
	874:  uint16(aux_sym_unicode_character_escape_token1),
	875:  uint16(aux_sym_unicode_character_escape_token2),
	876:  uint16(aux_sym_control_escape_token2),
	877:  uint16(sym_control_letter_escape),
	878:  uint16(2),
	879:  uint16(202),
	880:  uint16(6),
	881:  uint16(anon_sym_LPAREN_QMARK),
	882:  uint16(sym_pattern_character),
	883:  uint16(anon_sym_LBRACK),
	884:  uint16(anon_sym_LPAREN),
	885:  uint16(aux_sym_control_escape_token1),
	886:  uint16(sym_identity_escape),
	887:  uint16(200),
	888:  uint16(25),
	890:  uint16(anon_sym_PIPE),
	891:  uint16(sym_any_character),
	892:  uint16(anon_sym_CARET),
	893:  uint16(sym_end_assertion),
	894:  uint16(sym_boundary_assertion),
	895:  uint16(sym_non_boundary_assertion),
	896:  uint16(anon_sym_RPAREN),
	897:  uint16(anon_sym_LPAREN_QMARK_LT),
	898:  uint16(anon_sym_LBRACK_COLON),
	899:  uint16(anon_sym_LPAREN_QMARKP_LT),
	900:  uint16(anon_sym_LPAREN_QMARK_COLON),
	901:  uint16(anon_sym_STAR),
	902:  uint16(anon_sym_QMARK),
	903:  uint16(anon_sym_PLUS),
	904:  uint16(anon_sym_LBRACE),
	905:  uint16(anon_sym_BSLASHk),
	906:  uint16(anon_sym_LPAREN_QMARKP_EQ),
	907:  uint16(sym_decimal_escape),
	908:  uint16(aux_sym_character_class_escape_token1),
	909:  uint16(aux_sym_character_class_escape_token2),
	910:  uint16(aux_sym_unicode_character_escape_token1),
	911:  uint16(aux_sym_unicode_character_escape_token2),
	912:  uint16(aux_sym_control_escape_token2),
	913:  uint16(sym_control_letter_escape),
	914:  uint16(2),
	915:  uint16(206),
	916:  uint16(6),
	917:  uint16(anon_sym_LPAREN_QMARK),
	918:  uint16(sym_pattern_character),
	919:  uint16(anon_sym_LBRACK),
	920:  uint16(anon_sym_LPAREN),
	921:  uint16(aux_sym_control_escape_token1),
	922:  uint16(sym_identity_escape),
	923:  uint16(204),
	924:  uint16(25),
	926:  uint16(anon_sym_PIPE),
	927:  uint16(sym_any_character),
	928:  uint16(anon_sym_CARET),
	929:  uint16(sym_end_assertion),
	930:  uint16(sym_boundary_assertion),
	931:  uint16(sym_non_boundary_assertion),
	932:  uint16(anon_sym_RPAREN),
	933:  uint16(anon_sym_LPAREN_QMARK_LT),
	934:  uint16(anon_sym_LBRACK_COLON),
	935:  uint16(anon_sym_LPAREN_QMARKP_LT),
	936:  uint16(anon_sym_LPAREN_QMARK_COLON),
	937:  uint16(anon_sym_STAR),
	938:  uint16(anon_sym_QMARK),
	939:  uint16(anon_sym_PLUS),
	940:  uint16(anon_sym_LBRACE),
	941:  uint16(anon_sym_BSLASHk),
	942:  uint16(anon_sym_LPAREN_QMARKP_EQ),
	943:  uint16(sym_decimal_escape),
	944:  uint16(aux_sym_character_class_escape_token1),
	945:  uint16(aux_sym_character_class_escape_token2),
	946:  uint16(aux_sym_unicode_character_escape_token1),
	947:  uint16(aux_sym_unicode_character_escape_token2),
	948:  uint16(aux_sym_control_escape_token2),
	949:  uint16(sym_control_letter_escape),
	950:  uint16(2),
	951:  uint16(210),
	952:  uint16(6),
	953:  uint16(anon_sym_LPAREN_QMARK),
	954:  uint16(sym_pattern_character),
	955:  uint16(anon_sym_LBRACK),
	956:  uint16(anon_sym_LPAREN),
	957:  uint16(aux_sym_control_escape_token1),
	958:  uint16(sym_identity_escape),
	959:  uint16(208),
	960:  uint16(25),
	962:  uint16(anon_sym_PIPE),
	963:  uint16(sym_any_character),
	964:  uint16(anon_sym_CARET),
	965:  uint16(sym_end_assertion),
	966:  uint16(sym_boundary_assertion),
	967:  uint16(sym_non_boundary_assertion),
	968:  uint16(anon_sym_RPAREN),
	969:  uint16(anon_sym_LPAREN_QMARK_LT),
	970:  uint16(anon_sym_LBRACK_COLON),
	971:  uint16(anon_sym_LPAREN_QMARKP_LT),
	972:  uint16(anon_sym_LPAREN_QMARK_COLON),
	973:  uint16(anon_sym_STAR),
	974:  uint16(anon_sym_QMARK),
	975:  uint16(anon_sym_PLUS),
	976:  uint16(anon_sym_LBRACE),
	977:  uint16(anon_sym_BSLASHk),
	978:  uint16(anon_sym_LPAREN_QMARKP_EQ),
	979:  uint16(sym_decimal_escape),
	980:  uint16(aux_sym_character_class_escape_token1),
	981:  uint16(aux_sym_character_class_escape_token2),
	982:  uint16(aux_sym_unicode_character_escape_token1),
	983:  uint16(aux_sym_unicode_character_escape_token2),
	984:  uint16(aux_sym_control_escape_token2),
	985:  uint16(sym_control_letter_escape),
	986:  uint16(2),
	987:  uint16(214),
	988:  uint16(6),
	989:  uint16(anon_sym_LPAREN_QMARK),
	990:  uint16(sym_pattern_character),
	991:  uint16(anon_sym_LBRACK),
	992:  uint16(anon_sym_LPAREN),
	993:  uint16(aux_sym_control_escape_token1),
	994:  uint16(sym_identity_escape),
	995:  uint16(212),
	996:  uint16(25),
	998:  uint16(anon_sym_PIPE),
	999:  uint16(sym_any_character),
	1000: uint16(anon_sym_CARET),
	1001: uint16(sym_end_assertion),
	1002: uint16(sym_boundary_assertion),
	1003: uint16(sym_non_boundary_assertion),
	1004: uint16(anon_sym_RPAREN),
	1005: uint16(anon_sym_LPAREN_QMARK_LT),
	1006: uint16(anon_sym_LBRACK_COLON),
	1007: uint16(anon_sym_LPAREN_QMARKP_LT),
	1008: uint16(anon_sym_LPAREN_QMARK_COLON),
	1009: uint16(anon_sym_STAR),
	1010: uint16(anon_sym_QMARK),
	1011: uint16(anon_sym_PLUS),
	1012: uint16(anon_sym_LBRACE),
	1013: uint16(anon_sym_BSLASHk),
	1014: uint16(anon_sym_LPAREN_QMARKP_EQ),
	1015: uint16(sym_decimal_escape),
	1016: uint16(aux_sym_character_class_escape_token1),
	1017: uint16(aux_sym_character_class_escape_token2),
	1018: uint16(aux_sym_unicode_character_escape_token1),
	1019: uint16(aux_sym_unicode_character_escape_token2),
	1020: uint16(aux_sym_control_escape_token2),
	1021: uint16(sym_control_letter_escape),
	1022: uint16(2),
	1023: uint16(218),
	1024: uint16(6),
	1025: uint16(anon_sym_LPAREN_QMARK),
	1026: uint16(sym_pattern_character),
	1027: uint16(anon_sym_LBRACK),
	1028: uint16(anon_sym_LPAREN),
	1029: uint16(aux_sym_control_escape_token1),
	1030: uint16(sym_identity_escape),
	1031: uint16(216),
	1032: uint16(25),
	1034: uint16(anon_sym_PIPE),
	1035: uint16(sym_any_character),
	1036: uint16(anon_sym_CARET),
	1037: uint16(sym_end_assertion),
	1038: uint16(sym_boundary_assertion),
	1039: uint16(sym_non_boundary_assertion),
	1040: uint16(anon_sym_RPAREN),
	1041: uint16(anon_sym_LPAREN_QMARK_LT),
	1042: uint16(anon_sym_LBRACK_COLON),
	1043: uint16(anon_sym_LPAREN_QMARKP_LT),
	1044: uint16(anon_sym_LPAREN_QMARK_COLON),
	1045: uint16(anon_sym_STAR),
	1046: uint16(anon_sym_QMARK),
	1047: uint16(anon_sym_PLUS),
	1048: uint16(anon_sym_LBRACE),
	1049: uint16(anon_sym_BSLASHk),
	1050: uint16(anon_sym_LPAREN_QMARKP_EQ),
	1051: uint16(sym_decimal_escape),
	1052: uint16(aux_sym_character_class_escape_token1),
	1053: uint16(aux_sym_character_class_escape_token2),
	1054: uint16(aux_sym_unicode_character_escape_token1),
	1055: uint16(aux_sym_unicode_character_escape_token2),
	1056: uint16(aux_sym_control_escape_token2),
	1057: uint16(sym_control_letter_escape),
	1058: uint16(2),
	1059: uint16(222),
	1060: uint16(6),
	1061: uint16(anon_sym_LPAREN_QMARK),
	1062: uint16(sym_pattern_character),
	1063: uint16(anon_sym_LBRACK),
	1064: uint16(anon_sym_LPAREN),
	1065: uint16(aux_sym_control_escape_token1),
	1066: uint16(sym_identity_escape),
	1067: uint16(220),
	1068: uint16(25),
	1070: uint16(anon_sym_PIPE),
	1071: uint16(sym_any_character),
	1072: uint16(anon_sym_CARET),
	1073: uint16(sym_end_assertion),
	1074: uint16(sym_boundary_assertion),
	1075: uint16(sym_non_boundary_assertion),
	1076: uint16(anon_sym_RPAREN),
	1077: uint16(anon_sym_LPAREN_QMARK_LT),
	1078: uint16(anon_sym_LBRACK_COLON),
	1079: uint16(anon_sym_LPAREN_QMARKP_LT),
	1080: uint16(anon_sym_LPAREN_QMARK_COLON),
	1081: uint16(anon_sym_STAR),
	1082: uint16(anon_sym_QMARK),
	1083: uint16(anon_sym_PLUS),
	1084: uint16(anon_sym_LBRACE),
	1085: uint16(anon_sym_BSLASHk),
	1086: uint16(anon_sym_LPAREN_QMARKP_EQ),
	1087: uint16(sym_decimal_escape),
	1088: uint16(aux_sym_character_class_escape_token1),
	1089: uint16(aux_sym_character_class_escape_token2),
	1090: uint16(aux_sym_unicode_character_escape_token1),
	1091: uint16(aux_sym_unicode_character_escape_token2),
	1092: uint16(aux_sym_control_escape_token2),
	1093: uint16(sym_control_letter_escape),
	1094: uint16(2),
	1095: uint16(226),
	1096: uint16(6),
	1097: uint16(anon_sym_LPAREN_QMARK),
	1098: uint16(sym_pattern_character),
	1099: uint16(anon_sym_LBRACK),
	1100: uint16(anon_sym_LPAREN),
	1101: uint16(aux_sym_control_escape_token1),
	1102: uint16(sym_identity_escape),
	1103: uint16(224),
	1104: uint16(25),
	1106: uint16(anon_sym_PIPE),
	1107: uint16(sym_any_character),
	1108: uint16(anon_sym_CARET),
	1109: uint16(sym_end_assertion),
	1110: uint16(sym_boundary_assertion),
	1111: uint16(sym_non_boundary_assertion),
	1112: uint16(anon_sym_RPAREN),
	1113: uint16(anon_sym_LPAREN_QMARK_LT),
	1114: uint16(anon_sym_LBRACK_COLON),
	1115: uint16(anon_sym_LPAREN_QMARKP_LT),
	1116: uint16(anon_sym_LPAREN_QMARK_COLON),
	1117: uint16(anon_sym_STAR),
	1118: uint16(anon_sym_QMARK),
	1119: uint16(anon_sym_PLUS),
	1120: uint16(anon_sym_LBRACE),
	1121: uint16(anon_sym_BSLASHk),
	1122: uint16(anon_sym_LPAREN_QMARKP_EQ),
	1123: uint16(sym_decimal_escape),
	1124: uint16(aux_sym_character_class_escape_token1),
	1125: uint16(aux_sym_character_class_escape_token2),
	1126: uint16(aux_sym_unicode_character_escape_token1),
	1127: uint16(aux_sym_unicode_character_escape_token2),
	1128: uint16(aux_sym_control_escape_token2),
	1129: uint16(sym_control_letter_escape),
	1130: uint16(2),
	1131: uint16(230),
	1132: uint16(6),
	1133: uint16(anon_sym_LPAREN_QMARK),
	1134: uint16(sym_pattern_character),
	1135: uint16(anon_sym_LBRACK),
	1136: uint16(anon_sym_LPAREN),
	1137: uint16(aux_sym_control_escape_token1),
	1138: uint16(sym_identity_escape),
	1139: uint16(228),
	1140: uint16(25),
	1142: uint16(anon_sym_PIPE),
	1143: uint16(sym_any_character),
	1144: uint16(anon_sym_CARET),
	1145: uint16(sym_end_assertion),
	1146: uint16(sym_boundary_assertion),
	1147: uint16(sym_non_boundary_assertion),
	1148: uint16(anon_sym_RPAREN),
	1149: uint16(anon_sym_LPAREN_QMARK_LT),
	1150: uint16(anon_sym_LBRACK_COLON),
	1151: uint16(anon_sym_LPAREN_QMARKP_LT),
	1152: uint16(anon_sym_LPAREN_QMARK_COLON),
	1153: uint16(anon_sym_STAR),
	1154: uint16(anon_sym_QMARK),
	1155: uint16(anon_sym_PLUS),
	1156: uint16(anon_sym_LBRACE),
	1157: uint16(anon_sym_BSLASHk),
	1158: uint16(anon_sym_LPAREN_QMARKP_EQ),
	1159: uint16(sym_decimal_escape),
	1160: uint16(aux_sym_character_class_escape_token1),
	1161: uint16(aux_sym_character_class_escape_token2),
	1162: uint16(aux_sym_unicode_character_escape_token1),
	1163: uint16(aux_sym_unicode_character_escape_token2),
	1164: uint16(aux_sym_control_escape_token2),
	1165: uint16(sym_control_letter_escape),
	1166: uint16(3),
	1167: uint16(236),
	1168: uint16(1),
	1169: uint16(anon_sym_QMARK),
	1170: uint16(234),
	1171: uint16(5),
	1172: uint16(anon_sym_LPAREN_QMARK),
	1173: uint16(anon_sym_LBRACK),
	1174: uint16(anon_sym_LPAREN),
	1175: uint16(aux_sym_control_escape_token1),
	1176: uint16(sym_identity_escape),
	1177: uint16(232),
	1178: uint16(22),
	1180: uint16(anon_sym_PIPE),
	1181: uint16(sym_any_character),
	1182: uint16(anon_sym_CARET),
	1183: uint16(sym_end_assertion),
	1184: uint16(sym_boundary_assertion),
	1185: uint16(sym_non_boundary_assertion),
	1186: uint16(anon_sym_RPAREN),
	1187: uint16(anon_sym_LPAREN_QMARK_LT),
	1188: uint16(sym_pattern_character),
	1189: uint16(anon_sym_LBRACK_COLON),
	1190: uint16(anon_sym_LPAREN_QMARKP_LT),
	1191: uint16(anon_sym_LPAREN_QMARK_COLON),
	1192: uint16(anon_sym_BSLASHk),
	1193: uint16(anon_sym_LPAREN_QMARKP_EQ),
	1194: uint16(sym_decimal_escape),
	1195: uint16(aux_sym_character_class_escape_token1),
	1196: uint16(aux_sym_character_class_escape_token2),
	1197: uint16(aux_sym_unicode_character_escape_token1),
	1198: uint16(aux_sym_unicode_character_escape_token2),
	1199: uint16(aux_sym_control_escape_token2),
	1200: uint16(sym_control_letter_escape),
	1201: uint16(3),
	1202: uint16(242),
	1203: uint16(1),
	1204: uint16(anon_sym_QMARK),
	1205: uint16(240),
	1206: uint16(5),
	1207: uint16(anon_sym_LPAREN_QMARK),
	1208: uint16(anon_sym_LBRACK),
	1209: uint16(anon_sym_LPAREN),
	1210: uint16(aux_sym_control_escape_token1),
	1211: uint16(sym_identity_escape),
	1212: uint16(238),
	1213: uint16(22),
	1215: uint16(anon_sym_PIPE),
	1216: uint16(sym_any_character),
	1217: uint16(anon_sym_CARET),
	1218: uint16(sym_end_assertion),
	1219: uint16(sym_boundary_assertion),
	1220: uint16(sym_non_boundary_assertion),
	1221: uint16(anon_sym_RPAREN),
	1222: uint16(anon_sym_LPAREN_QMARK_LT),
	1223: uint16(sym_pattern_character),
	1224: uint16(anon_sym_LBRACK_COLON),
	1225: uint16(anon_sym_LPAREN_QMARKP_LT),
	1226: uint16(anon_sym_LPAREN_QMARK_COLON),
	1227: uint16(anon_sym_BSLASHk),
	1228: uint16(anon_sym_LPAREN_QMARKP_EQ),
	1229: uint16(sym_decimal_escape),
	1230: uint16(aux_sym_character_class_escape_token1),
	1231: uint16(aux_sym_character_class_escape_token2),
	1232: uint16(aux_sym_unicode_character_escape_token1),
	1233: uint16(aux_sym_unicode_character_escape_token2),
	1234: uint16(aux_sym_control_escape_token2),
	1235: uint16(sym_control_letter_escape),
	1236: uint16(3),
	1237: uint16(248),
	1238: uint16(1),
	1239: uint16(anon_sym_QMARK),
	1240: uint16(246),
	1241: uint16(5),
	1242: uint16(anon_sym_LPAREN_QMARK),
	1243: uint16(anon_sym_LBRACK),
	1244: uint16(anon_sym_LPAREN),
	1245: uint16(aux_sym_control_escape_token1),
	1246: uint16(sym_identity_escape),
	1247: uint16(244),
	1248: uint16(22),
	1250: uint16(anon_sym_PIPE),
	1251: uint16(sym_any_character),
	1252: uint16(anon_sym_CARET),
	1253: uint16(sym_end_assertion),
	1254: uint16(sym_boundary_assertion),
	1255: uint16(sym_non_boundary_assertion),
	1256: uint16(anon_sym_RPAREN),
	1257: uint16(anon_sym_LPAREN_QMARK_LT),
	1258: uint16(sym_pattern_character),
	1259: uint16(anon_sym_LBRACK_COLON),
	1260: uint16(anon_sym_LPAREN_QMARKP_LT),
	1261: uint16(anon_sym_LPAREN_QMARK_COLON),
	1262: uint16(anon_sym_BSLASHk),
	1263: uint16(anon_sym_LPAREN_QMARKP_EQ),
	1264: uint16(sym_decimal_escape),
	1265: uint16(aux_sym_character_class_escape_token1),
	1266: uint16(aux_sym_character_class_escape_token2),
	1267: uint16(aux_sym_unicode_character_escape_token1),
	1268: uint16(aux_sym_unicode_character_escape_token2),
	1269: uint16(aux_sym_control_escape_token2),
	1270: uint16(sym_control_letter_escape),
	1271: uint16(3),
	1272: uint16(254),
	1273: uint16(1),
	1274: uint16(anon_sym_QMARK),
	1275: uint16(252),
	1276: uint16(5),
	1277: uint16(anon_sym_LPAREN_QMARK),
	1278: uint16(anon_sym_LBRACK),
	1279: uint16(anon_sym_LPAREN),
	1280: uint16(aux_sym_control_escape_token1),
	1281: uint16(sym_identity_escape),
	1282: uint16(250),
	1283: uint16(22),
	1285: uint16(anon_sym_PIPE),
	1286: uint16(sym_any_character),
	1287: uint16(anon_sym_CARET),
	1288: uint16(sym_end_assertion),
	1289: uint16(sym_boundary_assertion),
	1290: uint16(sym_non_boundary_assertion),
	1291: uint16(anon_sym_RPAREN),
	1292: uint16(anon_sym_LPAREN_QMARK_LT),
	1293: uint16(sym_pattern_character),
	1294: uint16(anon_sym_LBRACK_COLON),
	1295: uint16(anon_sym_LPAREN_QMARKP_LT),
	1296: uint16(anon_sym_LPAREN_QMARK_COLON),
	1297: uint16(anon_sym_BSLASHk),
	1298: uint16(anon_sym_LPAREN_QMARKP_EQ),
	1299: uint16(sym_decimal_escape),
	1300: uint16(aux_sym_character_class_escape_token1),
	1301: uint16(aux_sym_character_class_escape_token2),
	1302: uint16(aux_sym_unicode_character_escape_token1),
	1303: uint16(aux_sym_unicode_character_escape_token2),
	1304: uint16(aux_sym_control_escape_token2),
	1305: uint16(sym_control_letter_escape),
	1306: uint16(3),
	1307: uint16(260),
	1308: uint16(1),
	1309: uint16(anon_sym_QMARK),
	1310: uint16(258),
	1311: uint16(5),
	1312: uint16(anon_sym_LPAREN_QMARK),
	1313: uint16(anon_sym_LBRACK),
	1314: uint16(anon_sym_LPAREN),
	1315: uint16(aux_sym_control_escape_token1),
	1316: uint16(sym_identity_escape),
	1317: uint16(256),
	1318: uint16(22),
	1320: uint16(anon_sym_PIPE),
	1321: uint16(sym_any_character),
	1322: uint16(anon_sym_CARET),
	1323: uint16(sym_end_assertion),
	1324: uint16(sym_boundary_assertion),
	1325: uint16(sym_non_boundary_assertion),
	1326: uint16(anon_sym_RPAREN),
	1327: uint16(anon_sym_LPAREN_QMARK_LT),
	1328: uint16(sym_pattern_character),
	1329: uint16(anon_sym_LBRACK_COLON),
	1330: uint16(anon_sym_LPAREN_QMARKP_LT),
	1331: uint16(anon_sym_LPAREN_QMARK_COLON),
	1332: uint16(anon_sym_BSLASHk),
	1333: uint16(anon_sym_LPAREN_QMARKP_EQ),
	1334: uint16(sym_decimal_escape),
	1335: uint16(aux_sym_character_class_escape_token1),
	1336: uint16(aux_sym_character_class_escape_token2),
	1337: uint16(aux_sym_unicode_character_escape_token1),
	1338: uint16(aux_sym_unicode_character_escape_token2),
	1339: uint16(aux_sym_control_escape_token2),
	1340: uint16(sym_control_letter_escape),
	1341: uint16(3),
	1342: uint16(266),
	1343: uint16(1),
	1344: uint16(anon_sym_QMARK),
	1345: uint16(264),
	1346: uint16(5),
	1347: uint16(anon_sym_LPAREN_QMARK),
	1348: uint16(anon_sym_LBRACK),
	1349: uint16(anon_sym_LPAREN),
	1350: uint16(aux_sym_control_escape_token1),
	1351: uint16(sym_identity_escape),
	1352: uint16(262),
	1353: uint16(22),
	1355: uint16(anon_sym_PIPE),
	1356: uint16(sym_any_character),
	1357: uint16(anon_sym_CARET),
	1358: uint16(sym_end_assertion),
	1359: uint16(sym_boundary_assertion),
	1360: uint16(sym_non_boundary_assertion),
	1361: uint16(anon_sym_RPAREN),
	1362: uint16(anon_sym_LPAREN_QMARK_LT),
	1363: uint16(sym_pattern_character),
	1364: uint16(anon_sym_LBRACK_COLON),
	1365: uint16(anon_sym_LPAREN_QMARKP_LT),
	1366: uint16(anon_sym_LPAREN_QMARK_COLON),
	1367: uint16(anon_sym_BSLASHk),
	1368: uint16(anon_sym_LPAREN_QMARKP_EQ),
	1369: uint16(sym_decimal_escape),
	1370: uint16(aux_sym_character_class_escape_token1),
	1371: uint16(aux_sym_character_class_escape_token2),
	1372: uint16(aux_sym_unicode_character_escape_token1),
	1373: uint16(aux_sym_unicode_character_escape_token2),
	1374: uint16(aux_sym_control_escape_token2),
	1375: uint16(sym_control_letter_escape),
	1376: uint16(2),
	1377: uint16(270),
	1378: uint16(5),
	1379: uint16(anon_sym_LPAREN_QMARK),
	1380: uint16(anon_sym_LBRACK),
	1381: uint16(anon_sym_LPAREN),
	1382: uint16(aux_sym_control_escape_token1),
	1383: uint16(sym_identity_escape),
	1384: uint16(268),
	1385: uint16(22),
	1387: uint16(anon_sym_PIPE),
	1388: uint16(sym_any_character),
	1389: uint16(anon_sym_CARET),
	1390: uint16(sym_end_assertion),
	1391: uint16(sym_boundary_assertion),
	1392: uint16(sym_non_boundary_assertion),
	1393: uint16(anon_sym_RPAREN),
	1394: uint16(anon_sym_LPAREN_QMARK_LT),
	1395: uint16(sym_pattern_character),
	1396: uint16(anon_sym_LBRACK_COLON),
	1397: uint16(anon_sym_LPAREN_QMARKP_LT),
	1398: uint16(anon_sym_LPAREN_QMARK_COLON),
	1399: uint16(anon_sym_BSLASHk),
	1400: uint16(anon_sym_LPAREN_QMARKP_EQ),
	1401: uint16(sym_decimal_escape),
	1402: uint16(aux_sym_character_class_escape_token1),
	1403: uint16(aux_sym_character_class_escape_token2),
	1404: uint16(aux_sym_unicode_character_escape_token1),
	1405: uint16(aux_sym_unicode_character_escape_token2),
	1406: uint16(aux_sym_control_escape_token2),
	1407: uint16(sym_control_letter_escape),
	1408: uint16(2),
	1409: uint16(274),
	1410: uint16(5),
	1411: uint16(anon_sym_LPAREN_QMARK),
	1412: uint16(anon_sym_LBRACK),
	1413: uint16(anon_sym_LPAREN),
	1414: uint16(aux_sym_control_escape_token1),
	1415: uint16(sym_identity_escape),
	1416: uint16(272),
	1417: uint16(22),
	1419: uint16(anon_sym_PIPE),
	1420: uint16(sym_any_character),
	1421: uint16(anon_sym_CARET),
	1422: uint16(sym_end_assertion),
	1423: uint16(sym_boundary_assertion),
	1424: uint16(sym_non_boundary_assertion),
	1425: uint16(anon_sym_RPAREN),
	1426: uint16(anon_sym_LPAREN_QMARK_LT),
	1427: uint16(sym_pattern_character),
	1428: uint16(anon_sym_LBRACK_COLON),
	1429: uint16(anon_sym_LPAREN_QMARKP_LT),
	1430: uint16(anon_sym_LPAREN_QMARK_COLON),
	1431: uint16(anon_sym_BSLASHk),
	1432: uint16(anon_sym_LPAREN_QMARKP_EQ),
	1433: uint16(sym_decimal_escape),
	1434: uint16(aux_sym_character_class_escape_token1),
	1435: uint16(aux_sym_character_class_escape_token2),
	1436: uint16(aux_sym_unicode_character_escape_token1),
	1437: uint16(aux_sym_unicode_character_escape_token2),
	1438: uint16(aux_sym_control_escape_token2),
	1439: uint16(sym_control_letter_escape),
	1440: uint16(2),
	1441: uint16(278),
	1442: uint16(5),
	1443: uint16(anon_sym_LPAREN_QMARK),
	1444: uint16(anon_sym_LBRACK),
	1445: uint16(anon_sym_LPAREN),
	1446: uint16(aux_sym_control_escape_token1),
	1447: uint16(sym_identity_escape),
	1448: uint16(276),
	1449: uint16(22),
	1451: uint16(anon_sym_PIPE),
	1452: uint16(sym_any_character),
	1453: uint16(anon_sym_CARET),
	1454: uint16(sym_end_assertion),
	1455: uint16(sym_boundary_assertion),
	1456: uint16(sym_non_boundary_assertion),
	1457: uint16(anon_sym_RPAREN),
	1458: uint16(anon_sym_LPAREN_QMARK_LT),
	1459: uint16(sym_pattern_character),
	1460: uint16(anon_sym_LBRACK_COLON),
	1461: uint16(anon_sym_LPAREN_QMARKP_LT),
	1462: uint16(anon_sym_LPAREN_QMARK_COLON),
	1463: uint16(anon_sym_BSLASHk),
	1464: uint16(anon_sym_LPAREN_QMARKP_EQ),
	1465: uint16(sym_decimal_escape),
	1466: uint16(aux_sym_character_class_escape_token1),
	1467: uint16(aux_sym_character_class_escape_token2),
	1468: uint16(aux_sym_unicode_character_escape_token1),
	1469: uint16(aux_sym_unicode_character_escape_token2),
	1470: uint16(aux_sym_control_escape_token2),
	1471: uint16(sym_control_letter_escape),
	1472: uint16(2),
	1473: uint16(282),
	1474: uint16(5),
	1475: uint16(anon_sym_LPAREN_QMARK),
	1476: uint16(anon_sym_LBRACK),
	1477: uint16(anon_sym_LPAREN),
	1478: uint16(aux_sym_control_escape_token1),
	1479: uint16(sym_identity_escape),
	1480: uint16(280),
	1481: uint16(22),
	1483: uint16(anon_sym_PIPE),
	1484: uint16(sym_any_character),
	1485: uint16(anon_sym_CARET),
	1486: uint16(sym_end_assertion),
	1487: uint16(sym_boundary_assertion),
	1488: uint16(sym_non_boundary_assertion),
	1489: uint16(anon_sym_RPAREN),
	1490: uint16(anon_sym_LPAREN_QMARK_LT),
	1491: uint16(sym_pattern_character),
	1492: uint16(anon_sym_LBRACK_COLON),
	1493: uint16(anon_sym_LPAREN_QMARKP_LT),
	1494: uint16(anon_sym_LPAREN_QMARK_COLON),
	1495: uint16(anon_sym_BSLASHk),
	1496: uint16(anon_sym_LPAREN_QMARKP_EQ),
	1497: uint16(sym_decimal_escape),
	1498: uint16(aux_sym_character_class_escape_token1),
	1499: uint16(aux_sym_character_class_escape_token2),
	1500: uint16(aux_sym_unicode_character_escape_token1),
	1501: uint16(aux_sym_unicode_character_escape_token2),
	1502: uint16(aux_sym_control_escape_token2),
	1503: uint16(sym_control_letter_escape),
	1504: uint16(2),
	1505: uint16(286),
	1506: uint16(5),
	1507: uint16(anon_sym_LPAREN_QMARK),
	1508: uint16(anon_sym_LBRACK),
	1509: uint16(anon_sym_LPAREN),
	1510: uint16(aux_sym_control_escape_token1),
	1511: uint16(sym_identity_escape),
	1512: uint16(284),
	1513: uint16(22),
	1515: uint16(anon_sym_PIPE),
	1516: uint16(sym_any_character),
	1517: uint16(anon_sym_CARET),
	1518: uint16(sym_end_assertion),
	1519: uint16(sym_boundary_assertion),
	1520: uint16(sym_non_boundary_assertion),
	1521: uint16(anon_sym_RPAREN),
	1522: uint16(anon_sym_LPAREN_QMARK_LT),
	1523: uint16(sym_pattern_character),
	1524: uint16(anon_sym_LBRACK_COLON),
	1525: uint16(anon_sym_LPAREN_QMARKP_LT),
	1526: uint16(anon_sym_LPAREN_QMARK_COLON),
	1527: uint16(anon_sym_BSLASHk),
	1528: uint16(anon_sym_LPAREN_QMARKP_EQ),
	1529: uint16(sym_decimal_escape),
	1530: uint16(aux_sym_character_class_escape_token1),
	1531: uint16(aux_sym_character_class_escape_token2),
	1532: uint16(aux_sym_unicode_character_escape_token1),
	1533: uint16(aux_sym_unicode_character_escape_token2),
	1534: uint16(aux_sym_control_escape_token2),
	1535: uint16(sym_control_letter_escape),
	1536: uint16(2),
	1537: uint16(290),
	1538: uint16(5),
	1539: uint16(anon_sym_LPAREN_QMARK),
	1540: uint16(anon_sym_LBRACK),
	1541: uint16(anon_sym_LPAREN),
	1542: uint16(aux_sym_control_escape_token1),
	1543: uint16(sym_identity_escape),
	1544: uint16(288),
	1545: uint16(22),
	1547: uint16(anon_sym_PIPE),
	1548: uint16(sym_any_character),
	1549: uint16(anon_sym_CARET),
	1550: uint16(sym_end_assertion),
	1551: uint16(sym_boundary_assertion),
	1552: uint16(sym_non_boundary_assertion),
	1553: uint16(anon_sym_RPAREN),
	1554: uint16(anon_sym_LPAREN_QMARK_LT),
	1555: uint16(sym_pattern_character),
	1556: uint16(anon_sym_LBRACK_COLON),
	1557: uint16(anon_sym_LPAREN_QMARKP_LT),
	1558: uint16(anon_sym_LPAREN_QMARK_COLON),
	1559: uint16(anon_sym_BSLASHk),
	1560: uint16(anon_sym_LPAREN_QMARKP_EQ),
	1561: uint16(sym_decimal_escape),
	1562: uint16(aux_sym_character_class_escape_token1),
	1563: uint16(aux_sym_character_class_escape_token2),
	1564: uint16(aux_sym_unicode_character_escape_token1),
	1565: uint16(aux_sym_unicode_character_escape_token2),
	1566: uint16(aux_sym_control_escape_token2),
	1567: uint16(sym_control_letter_escape),
	1568: uint16(2),
	1569: uint16(292),
	1570: uint16(5),
	1571: uint16(anon_sym_LPAREN_QMARK),
	1572: uint16(anon_sym_LBRACK),
	1573: uint16(anon_sym_LPAREN),
	1574: uint16(aux_sym_control_escape_token1),
	1575: uint16(sym_identity_escape),
	1576: uint16(43),
	1577: uint16(22),
	1579: uint16(anon_sym_PIPE),
	1580: uint16(sym_any_character),
	1581: uint16(anon_sym_CARET),
	1582: uint16(sym_end_assertion),
	1583: uint16(sym_boundary_assertion),
	1584: uint16(sym_non_boundary_assertion),
	1585: uint16(anon_sym_RPAREN),
	1586: uint16(anon_sym_LPAREN_QMARK_LT),
	1587: uint16(sym_pattern_character),
	1588: uint16(anon_sym_LBRACK_COLON),
	1589: uint16(anon_sym_LPAREN_QMARKP_LT),
	1590: uint16(anon_sym_LPAREN_QMARK_COLON),
	1591: uint16(anon_sym_BSLASHk),
	1592: uint16(anon_sym_LPAREN_QMARKP_EQ),
	1593: uint16(sym_decimal_escape),
	1594: uint16(aux_sym_character_class_escape_token1),
	1595: uint16(aux_sym_character_class_escape_token2),
	1596: uint16(aux_sym_unicode_character_escape_token1),
	1597: uint16(aux_sym_unicode_character_escape_token2),
	1598: uint16(aux_sym_control_escape_token2),
	1599: uint16(sym_control_letter_escape),
	1600: uint16(13),
	1601: uint16(294),
	1602: uint16(1),
	1603: uint16(anon_sym_CARET),
	1604: uint16(296),
	1605: uint16(1),
	1606: uint16(anon_sym_DASH),
	1607: uint16(298),
	1608: uint16(1),
	1609: uint16(anon_sym_RBRACK),
	1610: uint16(300),
	1611: uint16(1),
	1612: uint16(anon_sym_LBRACK_COLON),
	1613: uint16(304),
	1614: uint16(1),
	1615: uint16(sym_class_character),
	1616: uint16(306),
	1617: uint16(1),
	1618: uint16(aux_sym_character_class_escape_token1),
	1619: uint16(308),
	1620: uint16(1),
	1621: uint16(aux_sym_character_class_escape_token2),
	1622: uint16(74),
	1623: uint16(1),
	1624: uint16(sym_unicode_character_escape),
	1625: uint16(310),
	1626: uint16(2),
	1627: uint16(aux_sym_unicode_character_escape_token1),
	1628: uint16(aux_sym_unicode_character_escape_token2),
	1629: uint16(312),
	1630: uint16(2),
	1631: uint16(aux_sym_control_escape_token1),
	1632: uint16(aux_sym_control_escape_token2),
	1633: uint16(69),
	1634: uint16(2),
	1635: uint16(sym_character_class_escape),
	1636: uint16(sym_control_escape),
	1637: uint16(302),
	1638: uint16(3),
	1639: uint16(anon_sym_BSLASH_DASH),
	1640: uint16(sym_control_letter_escape),
	1641: uint16(sym_identity_escape),
	1642: uint16(60),
	1643: uint16(3),
	1644: uint16(sym_posix_character_class),
	1645: uint16(sym_class_range),
	1646: uint16(aux_sym_character_class_repeat1),
	1647: uint16(12),
	1648: uint16(300),
	1649: uint16(1),
	1650: uint16(anon_sym_LBRACK_COLON),
	1651: uint16(304),
	1652: uint16(1),
	1653: uint16(sym_class_character),
	1654: uint16(306),
	1655: uint16(1),
	1656: uint16(aux_sym_character_class_escape_token1),
	1657: uint16(308),
	1658: uint16(1),
	1659: uint16(aux_sym_character_class_escape_token2),
	1660: uint16(314),
	1661: uint16(1),
	1662: uint16(anon_sym_DASH),
	1663: uint16(316),
	1664: uint16(1),
	1665: uint16(anon_sym_RBRACK),
	1666: uint16(74),
	1667: uint16(1),
	1668: uint16(sym_unicode_character_escape),
	1669: uint16(310),
	1670: uint16(2),
	1671: uint16(aux_sym_unicode_character_escape_token1),
	1672: uint16(aux_sym_unicode_character_escape_token2),
	1673: uint16(312),
	1674: uint16(2),
	1675: uint16(aux_sym_control_escape_token1),
	1676: uint16(aux_sym_control_escape_token2),
	1677: uint16(69),
	1678: uint16(2),
	1679: uint16(sym_character_class_escape),
	1680: uint16(sym_control_escape),
	1681: uint16(318),
	1682: uint16(3),
	1683: uint16(anon_sym_BSLASH_DASH),
	1684: uint16(sym_control_letter_escape),
	1685: uint16(sym_identity_escape),
	1686: uint16(61),
	1687: uint16(3),
	1688: uint16(sym_posix_character_class),
	1689: uint16(sym_class_range),
	1690: uint16(aux_sym_character_class_repeat1),
	1691: uint16(12),
	1692: uint16(300),
	1693: uint16(1),
	1694: uint16(anon_sym_LBRACK_COLON),
	1695: uint16(304),
	1696: uint16(1),
	1697: uint16(sym_class_character),
	1698: uint16(306),
	1699: uint16(1),
	1700: uint16(aux_sym_character_class_escape_token1),
	1701: uint16(308),
	1702: uint16(1),
	1703: uint16(aux_sym_character_class_escape_token2),
	1704: uint16(320),
	1705: uint16(1),
	1706: uint16(anon_sym_DASH),
	1707: uint16(322),
	1708: uint16(1),
	1709: uint16(anon_sym_RBRACK),
	1710: uint16(74),
	1711: uint16(1),
	1712: uint16(sym_unicode_character_escape),
	1713: uint16(310),
	1714: uint16(2),
	1715: uint16(aux_sym_unicode_character_escape_token1),
	1716: uint16(aux_sym_unicode_character_escape_token2),
	1717: uint16(312),
	1718: uint16(2),
	1719: uint16(aux_sym_control_escape_token1),
	1720: uint16(aux_sym_control_escape_token2),
	1721: uint16(69),
	1722: uint16(2),
	1723: uint16(sym_character_class_escape),
	1724: uint16(sym_control_escape),
	1725: uint16(318),
	1726: uint16(3),
	1727: uint16(anon_sym_BSLASH_DASH),
	1728: uint16(sym_control_letter_escape),
	1729: uint16(sym_identity_escape),
	1730: uint16(61),
	1731: uint16(3),
	1732: uint16(sym_posix_character_class),
	1733: uint16(sym_class_range),
	1734: uint16(aux_sym_character_class_repeat1),
	1735: uint16(12),
	1736: uint16(324),
	1737: uint16(1),
	1738: uint16(anon_sym_DASH),
	1739: uint16(327),
	1740: uint16(1),
	1741: uint16(anon_sym_RBRACK),
	1742: uint16(329),
	1743: uint16(1),
	1744: uint16(anon_sym_LBRACK_COLON),
	1745: uint16(335),
	1746: uint16(1),
	1747: uint16(sym_class_character),
	1748: uint16(338),
	1749: uint16(1),
	1750: uint16(aux_sym_character_class_escape_token1),
	1751: uint16(341),
	1752: uint16(1),
	1753: uint16(aux_sym_character_class_escape_token2),
	1754: uint16(74),
	1755: uint16(1),
	1756: uint16(sym_unicode_character_escape),
	1757: uint16(344),
	1758: uint16(2),
	1759: uint16(aux_sym_unicode_character_escape_token1),
	1760: uint16(aux_sym_unicode_character_escape_token2),
	1761: uint16(347),
	1762: uint16(2),
	1763: uint16(aux_sym_control_escape_token1),
	1764: uint16(aux_sym_control_escape_token2),
	1765: uint16(69),
	1766: uint16(2),
	1767: uint16(sym_character_class_escape),
	1768: uint16(sym_control_escape),
	1769: uint16(332),
	1770: uint16(3),
	1771: uint16(anon_sym_BSLASH_DASH),
	1772: uint16(sym_control_letter_escape),
	1773: uint16(sym_identity_escape),
	1774: uint16(61),
	1775: uint16(3),
	1776: uint16(sym_posix_character_class),
	1777: uint16(sym_class_range),
	1778: uint16(aux_sym_character_class_repeat1),
	1779: uint16(12),
	1780: uint16(300),
	1781: uint16(1),
	1782: uint16(anon_sym_LBRACK_COLON),
	1783: uint16(304),
	1784: uint16(1),
	1785: uint16(sym_class_character),
	1786: uint16(306),
	1787: uint16(1),
	1788: uint16(aux_sym_character_class_escape_token1),
	1789: uint16(308),
	1790: uint16(1),
	1791: uint16(aux_sym_character_class_escape_token2),
	1792: uint16(322),
	1793: uint16(1),
	1794: uint16(anon_sym_RBRACK),
	1795: uint16(350),
	1796: uint16(1),
	1797: uint16(anon_sym_DASH),
	1798: uint16(74),
	1799: uint16(1),
	1800: uint16(sym_unicode_character_escape),
	1801: uint16(310),
	1802: uint16(2),
	1803: uint16(aux_sym_unicode_character_escape_token1),
	1804: uint16(aux_sym_unicode_character_escape_token2),
	1805: uint16(312),
	1806: uint16(2),
	1807: uint16(aux_sym_control_escape_token1),
	1808: uint16(aux_sym_control_escape_token2),
	1809: uint16(69),
	1810: uint16(2),
	1811: uint16(sym_character_class_escape),
	1812: uint16(sym_control_escape),
	1813: uint16(352),
	1814: uint16(3),
	1815: uint16(anon_sym_BSLASH_DASH),
	1816: uint16(sym_control_letter_escape),
	1817: uint16(sym_identity_escape),
	1818: uint16(66),
	1819: uint16(3),
	1820: uint16(sym_posix_character_class),
	1821: uint16(sym_class_range),
	1822: uint16(aux_sym_character_class_repeat1),
	1823: uint16(12),
	1824: uint16(300),
	1825: uint16(1),
	1826: uint16(anon_sym_LBRACK_COLON),
	1827: uint16(304),
	1828: uint16(1),
	1829: uint16(sym_class_character),
	1830: uint16(306),
	1831: uint16(1),
	1832: uint16(aux_sym_character_class_escape_token1),
	1833: uint16(308),
	1834: uint16(1),
	1835: uint16(aux_sym_character_class_escape_token2),
	1836: uint16(354),
	1837: uint16(1),
	1838: uint16(anon_sym_DASH),
	1839: uint16(356),
	1840: uint16(1),
	1841: uint16(anon_sym_RBRACK),
	1842: uint16(74),
	1843: uint16(1),
	1844: uint16(sym_unicode_character_escape),
	1845: uint16(310),
	1846: uint16(2),
	1847: uint16(aux_sym_unicode_character_escape_token1),
	1848: uint16(aux_sym_unicode_character_escape_token2),
	1849: uint16(312),
	1850: uint16(2),
	1851: uint16(aux_sym_control_escape_token1),
	1852: uint16(aux_sym_control_escape_token2),
	1853: uint16(69),
	1854: uint16(2),
	1855: uint16(sym_character_class_escape),
	1856: uint16(sym_control_escape),
	1857: uint16(358),
	1858: uint16(3),
	1859: uint16(anon_sym_BSLASH_DASH),
	1860: uint16(sym_control_letter_escape),
	1861: uint16(sym_identity_escape),
	1862: uint16(65),
	1863: uint16(3),
	1864: uint16(sym_posix_character_class),
	1865: uint16(sym_class_range),
	1866: uint16(aux_sym_character_class_repeat1),
	1867: uint16(12),
	1868: uint16(300),
	1869: uint16(1),
	1870: uint16(anon_sym_LBRACK_COLON),
	1871: uint16(304),
	1872: uint16(1),
	1873: uint16(sym_class_character),
	1874: uint16(306),
	1875: uint16(1),
	1876: uint16(aux_sym_character_class_escape_token1),
	1877: uint16(308),
	1878: uint16(1),
	1879: uint16(aux_sym_character_class_escape_token2),
	1880: uint16(360),
	1881: uint16(1),
	1882: uint16(anon_sym_DASH),
	1883: uint16(362),
	1884: uint16(1),
	1885: uint16(anon_sym_RBRACK),
	1886: uint16(74),
	1887: uint16(1),
	1888: uint16(sym_unicode_character_escape),
	1889: uint16(310),
	1890: uint16(2),
	1891: uint16(aux_sym_unicode_character_escape_token1),
	1892: uint16(aux_sym_unicode_character_escape_token2),
	1893: uint16(312),
	1894: uint16(2),
	1895: uint16(aux_sym_control_escape_token1),
	1896: uint16(aux_sym_control_escape_token2),
	1897: uint16(69),
	1898: uint16(2),
	1899: uint16(sym_character_class_escape),
	1900: uint16(sym_control_escape),
	1901: uint16(364),
	1902: uint16(3),
	1903: uint16(anon_sym_BSLASH_DASH),
	1904: uint16(sym_control_letter_escape),
	1905: uint16(sym_identity_escape),
	1906: uint16(59),
	1907: uint16(3),
	1908: uint16(sym_posix_character_class),
	1909: uint16(sym_class_range),
	1910: uint16(aux_sym_character_class_repeat1),
	1911: uint16(12),
	1912: uint16(300),
	1913: uint16(1),
	1914: uint16(anon_sym_LBRACK_COLON),
	1915: uint16(304),
	1916: uint16(1),
	1917: uint16(sym_class_character),
	1918: uint16(306),
	1919: uint16(1),
	1920: uint16(aux_sym_character_class_escape_token1),
	1921: uint16(308),
	1922: uint16(1),
	1923: uint16(aux_sym_character_class_escape_token2),
	1924: uint16(366),
	1925: uint16(1),
	1926: uint16(anon_sym_DASH),
	1927: uint16(368),
	1928: uint16(1),
	1929: uint16(anon_sym_RBRACK),
	1930: uint16(74),
	1931: uint16(1),
	1932: uint16(sym_unicode_character_escape),
	1933: uint16(310),
	1934: uint16(2),
	1935: uint16(aux_sym_unicode_character_escape_token1),
	1936: uint16(aux_sym_unicode_character_escape_token2),
	1937: uint16(312),
	1938: uint16(2),
	1939: uint16(aux_sym_control_escape_token1),
	1940: uint16(aux_sym_control_escape_token2),
	1941: uint16(69),
	1942: uint16(2),
	1943: uint16(sym_character_class_escape),
	1944: uint16(sym_control_escape),
	1945: uint16(318),
	1946: uint16(3),
	1947: uint16(anon_sym_BSLASH_DASH),
	1948: uint16(sym_control_letter_escape),
	1949: uint16(sym_identity_escape),
	1950: uint16(61),
	1951: uint16(3),
	1952: uint16(sym_posix_character_class),
	1953: uint16(sym_class_range),
	1954: uint16(aux_sym_character_class_repeat1),
	1955: uint16(12),
	1956: uint16(300),
	1957: uint16(1),
	1958: uint16(anon_sym_LBRACK_COLON),
	1959: uint16(304),
	1960: uint16(1),
	1961: uint16(sym_class_character),
	1962: uint16(306),
	1963: uint16(1),
	1964: uint16(aux_sym_character_class_escape_token1),
	1965: uint16(308),
	1966: uint16(1),
	1967: uint16(aux_sym_character_class_escape_token2),
	1968: uint16(370),
	1969: uint16(1),
	1970: uint16(anon_sym_DASH),
	1971: uint16(372),
	1972: uint16(1),
	1973: uint16(anon_sym_RBRACK),
	1974: uint16(74),
	1975: uint16(1),
	1976: uint16(sym_unicode_character_escape),
	1977: uint16(310),
	1978: uint16(2),
	1979: uint16(aux_sym_unicode_character_escape_token1),
	1980: uint16(aux_sym_unicode_character_escape_token2),
	1981: uint16(312),
	1982: uint16(2),
	1983: uint16(aux_sym_control_escape_token1),
	1984: uint16(aux_sym_control_escape_token2),
	1985: uint16(69),
	1986: uint16(2),
	1987: uint16(sym_character_class_escape),
	1988: uint16(sym_control_escape),
	1989: uint16(318),
	1990: uint16(3),
	1991: uint16(anon_sym_BSLASH_DASH),
	1992: uint16(sym_control_letter_escape),
	1993: uint16(sym_identity_escape),
	1994: uint16(61),
	1995: uint16(3),
	1996: uint16(sym_posix_character_class),
	1997: uint16(sym_class_range),
	1998: uint16(aux_sym_character_class_repeat1),
	1999: uint16(9),
	2000: uint16(306),
	2001: uint16(1),
	2002: uint16(aux_sym_character_class_escape_token1),
	2003: uint16(308),
	2004: uint16(1),
	2005: uint16(aux_sym_character_class_escape_token2),
	2006: uint16(374),
	2007: uint16(1),
	2008: uint16(anon_sym_DASH),
	2009: uint16(378),
	2010: uint16(1),
	2011: uint16(sym_class_character),
	2012: uint16(74),
	2013: uint16(1),
	2014: uint16(sym_unicode_character_escape),
	2015: uint16(310),
	2016: uint16(2),
	2017: uint16(aux_sym_unicode_character_escape_token1),
	2018: uint16(aux_sym_unicode_character_escape_token2),
	2019: uint16(312),
	2020: uint16(2),
	2021: uint16(aux_sym_control_escape_token1),
	2022: uint16(aux_sym_control_escape_token2),
	2023: uint16(68),
	2024: uint16(2),
	2025: uint16(sym_character_class_escape),
	2026: uint16(sym_control_escape),
	2027: uint16(376),
	2028: uint16(5),
	2029: uint16(anon_sym_RBRACK),
	2030: uint16(anon_sym_LBRACK_COLON),
	2031: uint16(anon_sym_BSLASH_DASH),
	2032: uint16(sym_control_letter_escape),
	2033: uint16(sym_identity_escape),
	2034: uint16(1),
	2035: uint16(380),
	2036: uint16(13),
	2037: uint16(anon_sym_DASH),
	2038: uint16(anon_sym_RBRACK),
	2039: uint16(anon_sym_LBRACK_COLON),
	2040: uint16(anon_sym_BSLASH_DASH),
	2041: uint16(sym_class_character),
	2042: uint16(aux_sym_character_class_escape_token1),
	2043: uint16(aux_sym_character_class_escape_token2),
	2044: uint16(aux_sym_unicode_character_escape_token1),
	2045: uint16(aux_sym_unicode_character_escape_token2),
	2046: uint16(aux_sym_control_escape_token1),
	2047: uint16(aux_sym_control_escape_token2),
	2048: uint16(sym_control_letter_escape),
	2049: uint16(sym_identity_escape),
	2050: uint16(2),
	2051: uint16(382),
	2052: uint16(1),
	2053: uint16(anon_sym_DASH),
	2054: uint16(385),
	2055: uint16(12),
	2056: uint16(anon_sym_RBRACK),
	2057: uint16(anon_sym_LBRACK_COLON),
	2058: uint16(anon_sym_BSLASH_DASH),
	2059: uint16(sym_class_character),
	2060: uint16(aux_sym_character_class_escape_token1),
	2061: uint16(aux_sym_character_class_escape_token2),
	2062: uint16(aux_sym_unicode_character_escape_token1),
	2063: uint16(aux_sym_unicode_character_escape_token2),
	2064: uint16(aux_sym_control_escape_token1),
	2065: uint16(aux_sym_control_escape_token2),
	2066: uint16(sym_control_letter_escape),
	2067: uint16(sym_identity_escape),
	2068: uint16(1),
	2069: uint16(387),
	2070: uint16(13),
	2071: uint16(anon_sym_DASH),
	2072: uint16(anon_sym_RBRACK),
	2073: uint16(anon_sym_LBRACK_COLON),
	2074: uint16(anon_sym_BSLASH_DASH),
	2075: uint16(sym_class_character),
	2076: uint16(aux_sym_character_class_escape_token1),
	2077: uint16(aux_sym_character_class_escape_token2),
	2078: uint16(aux_sym_unicode_character_escape_token1),
	2079: uint16(aux_sym_unicode_character_escape_token2),
	2080: uint16(aux_sym_control_escape_token1),
	2081: uint16(aux_sym_control_escape_token2),
	2082: uint16(sym_control_letter_escape),
	2083: uint16(sym_identity_escape),
	2084: uint16(1),
	2085: uint16(389),
	2086: uint16(13),
	2087: uint16(anon_sym_DASH),
	2088: uint16(anon_sym_RBRACK),
	2089: uint16(anon_sym_LBRACK_COLON),
	2090: uint16(anon_sym_BSLASH_DASH),
	2091: uint16(sym_class_character),
	2092: uint16(aux_sym_character_class_escape_token1),
	2093: uint16(aux_sym_character_class_escape_token2),
	2094: uint16(aux_sym_unicode_character_escape_token1),
	2095: uint16(aux_sym_unicode_character_escape_token2),
	2096: uint16(aux_sym_control_escape_token1),
	2097: uint16(aux_sym_control_escape_token2),
	2098: uint16(sym_control_letter_escape),
	2099: uint16(sym_identity_escape),
	2100: uint16(1),
	2101: uint16(162),
	2102: uint16(13),
	2103: uint16(anon_sym_DASH),
	2104: uint16(anon_sym_RBRACK),
	2105: uint16(anon_sym_LBRACK_COLON),
	2106: uint16(anon_sym_BSLASH_DASH),
	2107: uint16(sym_class_character),
	2108: uint16(aux_sym_character_class_escape_token1),
	2109: uint16(aux_sym_character_class_escape_token2),
	2110: uint16(aux_sym_unicode_character_escape_token1),
	2111: uint16(aux_sym_unicode_character_escape_token2),
	2112: uint16(aux_sym_control_escape_token1),
	2113: uint16(aux_sym_control_escape_token2),
	2114: uint16(sym_control_letter_escape),
	2115: uint16(sym_identity_escape),
	2116: uint16(1),
	2117: uint16(376),
	2118: uint16(13),
	2119: uint16(anon_sym_DASH),
	2120: uint16(anon_sym_RBRACK),
	2121: uint16(anon_sym_LBRACK_COLON),
	2122: uint16(anon_sym_BSLASH_DASH),
	2123: uint16(sym_class_character),
	2124: uint16(aux_sym_character_class_escape_token1),
	2125: uint16(aux_sym_character_class_escape_token2),
	2126: uint16(aux_sym_unicode_character_escape_token1),
	2127: uint16(aux_sym_unicode_character_escape_token2),
	2128: uint16(aux_sym_control_escape_token1),
	2129: uint16(aux_sym_control_escape_token2),
	2130: uint16(sym_control_letter_escape),
	2131: uint16(sym_identity_escape),
	2132: uint16(1),
	2133: uint16(110),
	2134: uint16(13),
	2135: uint16(anon_sym_DASH),
	2136: uint16(anon_sym_RBRACK),
	2137: uint16(anon_sym_LBRACK_COLON),
	2138: uint16(anon_sym_BSLASH_DASH),
	2139: uint16(sym_class_character),
	2140: uint16(aux_sym_character_class_escape_token1),
	2141: uint16(aux_sym_character_class_escape_token2),
	2142: uint16(aux_sym_unicode_character_escape_token1),
	2143: uint16(aux_sym_unicode_character_escape_token2),
	2144: uint16(aux_sym_control_escape_token1),
	2145: uint16(aux_sym_control_escape_token2),
	2146: uint16(sym_control_letter_escape),
	2147: uint16(sym_identity_escape),
	2148: uint16(1),
	2149: uint16(114),
	2150: uint16(13),
	2151: uint16(anon_sym_DASH),
	2152: uint16(anon_sym_RBRACK),
	2153: uint16(anon_sym_LBRACK_COLON),
	2154: uint16(anon_sym_BSLASH_DASH),
	2155: uint16(sym_class_character),
	2156: uint16(aux_sym_character_class_escape_token1),
	2157: uint16(aux_sym_character_class_escape_token2),
	2158: uint16(aux_sym_unicode_character_escape_token1),
	2159: uint16(aux_sym_unicode_character_escape_token2),
	2160: uint16(aux_sym_control_escape_token1),
	2161: uint16(aux_sym_control_escape_token2),
	2162: uint16(sym_control_letter_escape),
	2163: uint16(sym_identity_escape),
	2164: uint16(1),
	2165: uint16(142),
	2166: uint16(13),
	2167: uint16(anon_sym_DASH),
	2168: uint16(anon_sym_RBRACK),
	2169: uint16(anon_sym_LBRACK_COLON),
	2170: uint16(anon_sym_BSLASH_DASH),
	2171: uint16(sym_class_character),
	2172: uint16(aux_sym_character_class_escape_token1),
	2173: uint16(aux_sym_character_class_escape_token2),
	2174: uint16(aux_sym_unicode_character_escape_token1),
	2175: uint16(aux_sym_unicode_character_escape_token2),
	2176: uint16(aux_sym_control_escape_token1),
	2177: uint16(aux_sym_control_escape_token2),
	2178: uint16(sym_control_letter_escape),
	2179: uint16(sym_identity_escape),
	2180: uint16(1),
	2181: uint16(190),
	2182: uint16(13),
	2183: uint16(anon_sym_DASH),
	2184: uint16(anon_sym_RBRACK),
	2185: uint16(anon_sym_LBRACK_COLON),
	2186: uint16(anon_sym_BSLASH_DASH),
	2187: uint16(sym_class_character),
	2188: uint16(aux_sym_character_class_escape_token1),
	2189: uint16(aux_sym_character_class_escape_token2),
	2190: uint16(aux_sym_unicode_character_escape_token1),
	2191: uint16(aux_sym_unicode_character_escape_token2),
	2192: uint16(aux_sym_control_escape_token1),
	2193: uint16(aux_sym_control_escape_token2),
	2194: uint16(sym_control_letter_escape),
	2195: uint16(sym_identity_escape),
	2196: uint16(9),
	2197: uint16(306),
	2198: uint16(1),
	2199: uint16(aux_sym_character_class_escape_token1),
	2200: uint16(308),
	2201: uint16(1),
	2202: uint16(aux_sym_character_class_escape_token2),
	2203: uint16(391),
	2204: uint16(1),
	2205: uint16(anon_sym_DASH),
	2206: uint16(393),
	2207: uint16(1),
	2208: uint16(anon_sym_RBRACK),
	2209: uint16(395),
	2210: uint16(1),
	2211: uint16(sym_class_character),
	2212: uint16(74),
	2213: uint16(1),
	2214: uint16(sym_unicode_character_escape),
	2215: uint16(310),
	2216: uint16(2),
	2217: uint16(aux_sym_unicode_character_escape_token1),
	2218: uint16(aux_sym_unicode_character_escape_token2),
	2219: uint16(312),
	2220: uint16(2),
	2221: uint16(aux_sym_control_escape_token1),
	2222: uint16(aux_sym_control_escape_token2),
	2223: uint16(68),
	2224: uint16(2),
	2225: uint16(sym_character_class_escape),
	2226: uint16(sym_control_escape),
	2227: uint16(9),
	2228: uint16(306),
	2229: uint16(1),
	2230: uint16(aux_sym_character_class_escape_token1),
	2231: uint16(308),
	2232: uint16(1),
	2233: uint16(aux_sym_character_class_escape_token2),
	2234: uint16(391),
	2235: uint16(1),
	2236: uint16(anon_sym_DASH),
	2237: uint16(395),
	2238: uint16(1),
	2239: uint16(sym_class_character),
	2240: uint16(397),
	2241: uint16(1),
	2242: uint16(anon_sym_RBRACK),
	2243: uint16(74),
	2244: uint16(1),
	2245: uint16(sym_unicode_character_escape),
	2246: uint16(310),
	2247: uint16(2),
	2248: uint16(aux_sym_unicode_character_escape_token1),
	2249: uint16(aux_sym_unicode_character_escape_token2),
	2250: uint16(312),
	2251: uint16(2),
	2252: uint16(aux_sym_control_escape_token1),
	2253: uint16(aux_sym_control_escape_token2),
	2254: uint16(68),
	2255: uint16(2),
	2256: uint16(sym_character_class_escape),
	2257: uint16(sym_control_escape),
	2258: uint16(8),
	2259: uint16(306),
	2260: uint16(1),
	2261: uint16(aux_sym_character_class_escape_token1),
	2262: uint16(308),
	2263: uint16(1),
	2264: uint16(aux_sym_character_class_escape_token2),
	2265: uint16(399),
	2266: uint16(1),
	2267: uint16(anon_sym_DASH),
	2268: uint16(401),
	2269: uint16(1),
	2270: uint16(sym_class_character),
	2271: uint16(74),
	2272: uint16(1),
	2273: uint16(sym_unicode_character_escape),
	2274: uint16(310),
	2275: uint16(2),
	2276: uint16(aux_sym_unicode_character_escape_token1),
	2277: uint16(aux_sym_unicode_character_escape_token2),
	2278: uint16(312),
	2279: uint16(2),
	2280: uint16(aux_sym_control_escape_token1),
	2281: uint16(aux_sym_control_escape_token2),
	2282: uint16(71),
	2283: uint16(2),
	2284: uint16(sym_character_class_escape),
	2285: uint16(sym_control_escape),
	2286: uint16(8),
	2287: uint16(306),
	2288: uint16(1),
	2289: uint16(aux_sym_character_class_escape_token1),
	2290: uint16(308),
	2291: uint16(1),
	2292: uint16(aux_sym_character_class_escape_token2),
	2293: uint16(374),
	2294: uint16(1),
	2295: uint16(anon_sym_DASH),
	2296: uint16(395),
	2297: uint16(1),
	2298: uint16(sym_class_character),
	2299: uint16(74),
	2300: uint16(1),
	2301: uint16(sym_unicode_character_escape),
	2302: uint16(310),
	2303: uint16(2),
	2304: uint16(aux_sym_unicode_character_escape_token1),
	2305: uint16(aux_sym_unicode_character_escape_token2),
	2306: uint16(312),
	2307: uint16(2),
	2308: uint16(aux_sym_control_escape_token1),
	2309: uint16(aux_sym_control_escape_token2),
	2310: uint16(68),
	2311: uint16(2),
	2312: uint16(sym_character_class_escape),
	2313: uint16(sym_control_escape),
	2314: uint16(4),
	2315: uint16(405),
	2316: uint16(1),
	2317: uint16(anon_sym_DASH),
	2318: uint16(407),
	2319: uint16(1),
	2320: uint16(aux_sym_posix_class_name_token1),
	2321: uint16(87),
	2322: uint16(1),
	2323: uint16(sym_flags),
	2324: uint16(403),
	2325: uint16(2),
	2326: uint16(anon_sym_EQ),
	2327: uint16(anon_sym_BANG),
	2328: uint16(3),
	2329: uint16(411),
	2330: uint16(1),
	2331: uint16(anon_sym_PIPE),
	2332: uint16(83),
	2333: uint16(1),
	2334: uint16(aux_sym_alternation_repeat1),
	2335: uint16(409),
	2336: uint16(2),
	2338: uint16(anon_sym_RPAREN),
	2339: uint16(3),
	2340: uint16(3),
	2341: uint16(1),
	2342: uint16(anon_sym_PIPE),
	2343: uint16(83),
	2344: uint16(1),
	2345: uint16(aux_sym_alternation_repeat1),
	2346: uint16(414),
	2347: uint16(2),
	2349: uint16(anon_sym_RPAREN),
	2350: uint16(3),
	2351: uint16(3),
	2352: uint16(1),
	2353: uint16(anon_sym_PIPE),
	2354: uint16(84),
	2355: uint16(1),
	2356: uint16(aux_sym_alternation_repeat1),
	2357: uint16(416),
	2358: uint16(2),
	2360: uint16(anon_sym_RPAREN),
	2361: uint16(3),
	2362: uint16(3),
	2363: uint16(1),
	2364: uint16(anon_sym_PIPE),
	2365: uint16(83),
	2366: uint16(1),
	2367: uint16(aux_sym_alternation_repeat1),
	2368: uint16(418),
	2369: uint16(2),
	2371: uint16(anon_sym_RPAREN),
	2372: uint16(3),
	2373: uint16(420),
	2374: uint16(1),
	2375: uint16(anon_sym_RPAREN),
	2376: uint16(422),
	2377: uint16(1),
	2378: uint16(anon_sym_DASH),
	2379: uint16(424),
	2380: uint16(1),
	2381: uint16(anon_sym_COLON),
	2382: uint16(2),
	2383: uint16(428),
	2384: uint16(1),
	2385: uint16(sym_group_name),
	2386: uint16(426),
	2387: uint16(2),
	2388: uint16(anon_sym_EQ),
	2389: uint16(anon_sym_BANG),
	2390: uint16(1),
	2391: uint16(409),
	2392: uint16(3),
	2394: uint16(anon_sym_PIPE),
	2395: uint16(anon_sym_RPAREN),
	2396: uint16(1),
	2397: uint16(430),
	2398: uint16(3),
	2399: uint16(anon_sym_RPAREN),
	2400: uint16(anon_sym_DASH),
	2401: uint16(anon_sym_COLON),
	2402: uint16(1),
	2403: uint16(416),
	2404: uint16(2),
	2406: uint16(anon_sym_RPAREN),
	2407: uint16(2),
	2408: uint16(432),
	2409: uint16(1),
	2410: uint16(sym_unicode_property),
	2411: uint16(133),
	2412: uint16(1),
	2413: uint16(sym_unicode_property_value_expression),
	2414: uint16(2),
	2415: uint16(434),
	2416: uint16(1),
	2417: uint16(anon_sym_DASH),
	2418: uint16(436),
	2419: uint16(1),
	2420: uint16(anon_sym_RBRACK),
	2421: uint16(2),
	2422: uint16(432),
	2423: uint16(1),
	2424: uint16(sym_unicode_property),
	2425: uint16(116),
	2426: uint16(1),
	2427: uint16(sym_unicode_property_value_expression),
	2428: uint16(2),
	2429: uint16(434),
	2430: uint16(1),
	2431: uint16(anon_sym_DASH),
	2432: uint16(438),
	2433: uint16(1),
	2434: uint16(anon_sym_RBRACK),
	2435: uint16(2),
	2436: uint16(440),
	2437: uint16(1),
	2438: uint16(aux_sym_posix_class_name_token1),
	2439: uint16(115),
	2440: uint16(1),
	2441: uint16(sym_posix_class_name),
	2442: uint16(2),
	2443: uint16(407),
	2444: uint16(1),
	2445: uint16(aux_sym_posix_class_name_token1),
	2446: uint16(104),
	2447: uint16(1),
	2448: uint16(sym_flags),
	2449: uint16(2),
	2450: uint16(442),
	2451: uint16(1),
	2452: uint16(anon_sym_COMMA),
	2453: uint16(444),
	2454: uint16(1),
	2455: uint16(sym_decimal_digits),
	2456: uint16(2),
	2457: uint16(434),
	2458: uint16(1),
	2459: uint16(anon_sym_DASH),
	2460: uint16(446),
	2461: uint16(1),
	2462: uint16(anon_sym_RBRACK),
	2463: uint16(2),
	2464: uint16(448),
	2465: uint16(1),
	2466: uint16(anon_sym_EQ),
	2467: uint16(450),
	2468: uint16(1),
	2469: uint16(anon_sym_RBRACE),
	2470: uint16(2),
	2471: uint16(452),
	2472: uint16(1),
	2473: uint16(anon_sym_RBRACE),
	2474: uint16(454),
	2475: uint16(1),
	2476: uint16(sym_decimal_digits),
	2477: uint16(2),
	2478: uint16(456),
	2479: uint16(1),
	2480: uint16(anon_sym_COMMA),
	2481: uint16(458),
	2482: uint16(1),
	2483: uint16(anon_sym_RBRACE),
	2484: uint16(2),
	2485: uint16(460),
	2486: uint16(1),
	2487: uint16(anon_sym_RPAREN),
	2488: uint16(462),
	2489: uint16(1),
	2490: uint16(anon_sym_COLON),
	2491: uint16(2),
	2492: uint16(464),
	2493: uint16(1),
	2494: uint16(anon_sym_RPAREN),
	2495: uint16(466),
	2496: uint16(1),
	2497: uint16(anon_sym_COLON),
	2498: uint16(2),
	2499: uint16(407),
	2500: uint16(1),
	2501: uint16(aux_sym_posix_class_name_token1),
	2502: uint16(103),
	2503: uint16(1),
	2504: uint16(sym_flags),
	2505: uint16(2),
	2506: uint16(440),
	2507: uint16(1),
	2508: uint16(aux_sym_posix_class_name_token1),
	2509: uint16(132),
	2510: uint16(1),
	2511: uint16(sym_posix_class_name),
	2512: uint16(2),
	2513: uint16(434),
	2514: uint16(1),
	2515: uint16(anon_sym_DASH),
	2516: uint16(468),
	2517: uint16(1),
	2518: uint16(anon_sym_RBRACK),
	2519: uint16(1),
	2520: uint16(470),
	2521: uint16(1),
	2522: uint16(anon_sym_RPAREN),
	2523: uint16(1),
	2524: uint16(472),
	2525: uint16(1),
	2526: uint16(anon_sym_RPAREN),
	2527: uint16(1),
	2528: uint16(474),
	2529: uint16(1),
	2530: uint16(anon_sym_RPAREN),
	2531: uint16(1),
	2532: uint16(476),
	2533: uint16(1),
	2534: uint16(anon_sym_LBRACE),
	2535: uint16(1),
	2536: uint16(478),
	2537: uint16(1),
	2538: uint16(anon_sym_GT),
	2539: uint16(1),
	2540: uint16(480),
	2541: uint16(1),
	2542: uint16(sym_group_name),
	2543: uint16(1),
	2544: uint16(482),
	2545: uint16(1),
	2546: uint16(anon_sym_RPAREN),
	2547: uint16(1),
	2548: uint16(484),
	2549: uint16(1),
	2550: uint16(anon_sym_COLON_RBRACK),
	2551: uint16(1),
	2552: uint16(486),
	2553: uint16(1),
	2554: uint16(anon_sym_RBRACE),
	2555: uint16(1),
	2556: uint16(452),
	2557: uint16(1),
	2558: uint16(anon_sym_RBRACE),
	2559: uint16(1),
	2560: uint16(434),
	2561: uint16(1),
	2562: uint16(anon_sym_DASH),
	2563: uint16(1),
	2564: uint16(488),
	2565: uint16(1),
	2566: uint16(anon_sym_RBRACE),
	2567: uint16(1),
	2568: uint16(490),
	2569: uint16(1),
	2570: uint16(sym_decimal_digits),
	2571: uint16(1),
	2572: uint16(492),
	2573: uint16(1),
	2574: uint16(anon_sym_GT),
	2575: uint16(1),
	2576: uint16(494),
	2577: uint16(1),
	2578: uint16(anon_sym_RPAREN),
	2579: uint16(1),
	2580: uint16(496),
	2581: uint16(1),
	2582: uint16(anon_sym_RPAREN),
	2583: uint16(1),
	2584: uint16(498),
	2585: uint16(1),
	2586: uint16(sym_group_name),
	2587: uint16(1),
	2588: uint16(500),
	2589: uint16(1),
	2590: uint16(sym_unicode_property),
	2591: uint16(1),
	2592: uint16(460),
	2593: uint16(1),
	2594: uint16(anon_sym_RPAREN),
	2595: uint16(1),
	2596: uint16(502),
	2597: uint16(1),
	2598: uint16(anon_sym_RPAREN),
	2599: uint16(1),
	2600: uint16(504),
	2601: uint16(1),
	2602: uint16(anon_sym_RPAREN),
	2603: uint16(1),
	2604: uint16(506),
	2605: uint16(1),
	2607: uint16(1),
	2608: uint16(508),
	2609: uint16(1),
	2610: uint16(anon_sym_LT),
	2611: uint16(1),
	2612: uint16(510),
	2613: uint16(1),
	2614: uint16(anon_sym_COLON_RBRACK),
	2615: uint16(1),
	2616: uint16(512),
	2617: uint16(1),
	2618: uint16(anon_sym_COLON_RBRACK),
	2619: uint16(1),
	2620: uint16(514),
	2621: uint16(1),
	2622: uint16(anon_sym_RBRACE),
	2623: uint16(1),
	2624: uint16(516),
	2625: uint16(1),
	2626: uint16(anon_sym_RBRACE),
	2627: uint16(1),
	2628: uint16(428),
	2629: uint16(1),
	2630: uint16(sym_group_name),
	2631: uint16(1),
	2632: uint16(518),
	2633: uint16(1),
	2634: uint16(anon_sym_LBRACE),
}

var ts_small_parse_table_map = [124]uint32_t{
	1:   uint32(50),
	2:   uint32(86),
	3:   uint32(122),
	4:   uint32(158),
	5:   uint32(194),
	6:   uint32(230),
	7:   uint32(266),
	8:   uint32(302),
	9:   uint32(338),
	10:  uint32(374),
	11:  uint32(410),
	12:  uint32(446),
	13:  uint32(482),
	14:  uint32(518),
	15:  uint32(554),
	16:  uint32(590),
	17:  uint32(626),
	18:  uint32(662),
	19:  uint32(698),
	20:  uint32(734),
	21:  uint32(770),
	22:  uint32(806),
	23:  uint32(842),
	24:  uint32(878),
	25:  uint32(914),
	26:  uint32(950),
	27:  uint32(986),
	28:  uint32(1022),
	29:  uint32(1058),
	30:  uint32(1094),
	31:  uint32(1130),
	32:  uint32(1166),
	33:  uint32(1201),
	34:  uint32(1236),
	35:  uint32(1271),
	36:  uint32(1306),
	37:  uint32(1341),
	38:  uint32(1376),
	39:  uint32(1408),
	40:  uint32(1440),
	41:  uint32(1472),
	42:  uint32(1504),
	43:  uint32(1536),
	44:  uint32(1568),
	45:  uint32(1600),
	46:  uint32(1647),
	47:  uint32(1691),
	48:  uint32(1735),
	49:  uint32(1779),
	50:  uint32(1823),
	51:  uint32(1867),
	52:  uint32(1911),
	53:  uint32(1955),
	54:  uint32(1999),
	55:  uint32(2034),
	56:  uint32(2050),
	57:  uint32(2068),
	58:  uint32(2084),
	59:  uint32(2100),
	60:  uint32(2116),
	61:  uint32(2132),
	62:  uint32(2148),
	63:  uint32(2164),
	64:  uint32(2180),
	65:  uint32(2196),
	66:  uint32(2227),
	67:  uint32(2258),
	68:  uint32(2286),
	69:  uint32(2314),
	70:  uint32(2328),
	71:  uint32(2339),
	72:  uint32(2350),
	73:  uint32(2361),
	74:  uint32(2372),
	75:  uint32(2382),
	76:  uint32(2390),
	77:  uint32(2396),
	78:  uint32(2402),
	79:  uint32(2407),
	80:  uint32(2414),
	81:  uint32(2421),
	82:  uint32(2428),
	83:  uint32(2435),
	84:  uint32(2442),
	85:  uint32(2449),
	86:  uint32(2456),
	87:  uint32(2463),
	88:  uint32(2470),
	89:  uint32(2477),
	90:  uint32(2484),
	91:  uint32(2491),
	92:  uint32(2498),
	93:  uint32(2505),
	94:  uint32(2512),
	95:  uint32(2519),
	96:  uint32(2523),
	97:  uint32(2527),
	98:  uint32(2531),
	99:  uint32(2535),
	100: uint32(2539),
	101: uint32(2543),
	102: uint32(2547),
	103: uint32(2551),
	104: uint32(2555),
	105: uint32(2559),
	106: uint32(2563),
	107: uint32(2567),
	108: uint32(2571),
	109: uint32(2575),
	110: uint32(2579),
	111: uint32(2583),
	112: uint32(2587),
	113: uint32(2591),
	114: uint32(2595),
	115: uint32(2599),
	116: uint32(2603),
	117: uint32(2607),
	118: uint32(2611),
	119: uint32(2615),
	120: uint32(2619),
	121: uint32(2623),
	122: uint32(2627),
	123: uint32(2631),
}

var ts_parse_actions = [520]TSParseActionEntry{
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(10)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(13)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(17)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(82)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(88)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(58)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(96)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(135)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(4)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(130)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(124)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(14)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(111)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(27)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(15)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(15)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(13)),
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
		Fsymbol:      uint16(aux_sym_alternation_repeat1),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fsymbol:      uint16(sym_term),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_term_repeat1),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(1 != 0),
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
		Fsymbol:      uint16(aux_sym_term_repeat1),
	})))),
	47: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(13)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	48: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	49: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_term_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(17)),
		Frepetition: libc.BoolUint8(1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_term_repeat1),
	})))),
	53: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(82)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	54: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	55: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_term_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(88)),
		Frepetition: libc.BoolUint8(1 != 0),
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
		Fcount: uint8(2),
	}})),
	58: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_term_repeat1),
	})))),
	59: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(58)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	60: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	61: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_term_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(96)),
		Frepetition: libc.BoolUint8(1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_term_repeat1),
	})))),
	65: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(5)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	66: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	67: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_term_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(135)),
		Frepetition: libc.BoolUint8(1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(1 != 0),
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
		Fsymbol:      uint16(aux_sym_term_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(4)),
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
		Fsymbol:      uint16(aux_sym_term_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(130)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(1 != 0),
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
		Fsymbol:      uint16(aux_sym_term_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(124)),
		Frepetition: libc.BoolUint8(1 != 0),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fsymbol:      uint16(aux_sym_term_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(14)),
		Frepetition: libc.BoolUint8(1 != 0),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fsymbol:      uint16(aux_sym_term_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(111)),
		Frepetition: libc.BoolUint8(1 != 0),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fsymbol:      uint16(aux_sym_term_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(27)),
		Frepetition: libc.BoolUint8(1 != 0),
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
		Fcount: uint8(2),
	}})),
	88: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_term_repeat1),
	})))),
	89: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(15)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	91: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_term_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(15)),
		Frepetition: libc.BoolUint8(1 != 0),
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
		Fcount: uint8(2),
	}})),
	94: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_term_repeat1),
	})))),
	95: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(13)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	97: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_term_repeat1),
	})))),
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
		Fcount: uint8(1),
	}})),
	99: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_term_repeat1),
	})))),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(50)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(46)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(47)),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(98)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fsymbol:      uint16(sym_character_class_escape),
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
		Fsymbol:      uint16(sym_character_class_escape),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fsymbol:      uint16(sym_control_escape),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_control_escape),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	117: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_lookaround_assertion),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_lookaround_assertion),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fsymbol:      uint16(sym_start_assertion),
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
		Fsymbol:      uint16(sym_start_assertion),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fsymbol:      uint16(sym_character_class),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_character_class),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	129: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_inline_flags_group),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_inline_flags_group),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	133: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_character_class),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_character_class),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	137: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_character_class),
		Fproduction_id: uint16(2),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_character_class),
		Fproduction_id: uint16(2),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	141: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_posix_character_class),
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
		Fsymbol:      uint16(sym_posix_character_class),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fsymbol:      uint16(sym_anonymous_capturing_group),
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
		Fsymbol:      uint16(sym_anonymous_capturing_group),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fsymbol:      uint16(sym_non_capturing_group),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_non_capturing_group),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_named_group_backreference),
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
		Fsymbol:      uint16(sym_named_group_backreference),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym__lookahead_assertion),
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
		Fsymbol:      uint16(sym__lookahead_assertion),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_unicode_character_escape),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_unicode_character_escape),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	165: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym__lookbehind_assertion),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym__lookbehind_assertion),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	169: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_character_class),
		Fproduction_id: uint16(3),
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
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_character_class),
		Fproduction_id: uint16(3),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	173: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_character_class),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_character_class),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	177: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_character_class),
		Fproduction_id: uint16(5),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_character_class),
		Fproduction_id: uint16(5),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	181: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_character_class),
		Fproduction_id: uint16(2),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_character_class),
		Fproduction_id: uint16(2),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fsymbol:      uint16(sym_backreference_escape),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_backreference_escape),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	189: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_character_class_escape),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_character_class_escape),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_inline_flags_group),
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
		Fcount: uint8(1),
	}})),
	195: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_inline_flags_group),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	197: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_named_capturing_group),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_named_capturing_group),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	201: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_character_class),
		Fproduction_id: uint16(8),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_character_class),
		Fproduction_id: uint16(8),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	205: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_character_class),
		Fproduction_id: uint16(3),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_character_class),
		Fproduction_id: uint16(3),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	209: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_character_class),
		Fproduction_id: uint16(9),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_character_class),
		Fproduction_id: uint16(9),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	213: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_character_class),
		Fproduction_id: uint16(10),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_character_class),
		Fproduction_id: uint16(10),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	217: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_inline_flags_group),
	})))),
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
		Fcount: uint8(1),
	}})),
	219: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_inline_flags_group),
	})))),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	221: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_character_class),
		Fproduction_id: uint16(13),
	})))),
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
		Fcount: uint8(1),
	}})),
	223: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_character_class),
		Fproduction_id: uint16(13),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	225: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_inline_flags_group),
	})))),
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
		Fcount: uint8(1),
	}})),
	227: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_inline_flags_group),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	229: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_inline_flags_group),
	})))),
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
		Fcount: uint8(1),
	}})),
	231: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_inline_flags_group),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	233: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_count_quantifier),
	})))),
	234: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	235: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_count_quantifier),
	})))),
	236: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
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
	237: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	238: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
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
	239: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_optional),
	})))),
	240: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	241: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_optional),
	})))),
	242: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
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
	243: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(55)),
	}})))),
	244: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
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
	245: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_one_or_more),
	})))),
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
		Fcount: uint8(1),
	}})),
	247: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_one_or_more),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(56)),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	251: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_count_quantifier),
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
		Fcount: uint8(1),
	}})),
	253: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_count_quantifier),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(53)),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_count_quantifier),
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
		Fcount: uint8(1),
	}})),
	259: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_count_quantifier),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(51)),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	263: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_zero_or_more),
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
		Fcount: uint8(1),
	}})),
	265: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_zero_or_more),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(54)),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	269: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_count_quantifier),
		Fproduction_id: uint16(12),
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
		Fcount: uint8(1),
	}})),
	271: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_count_quantifier),
		Fproduction_id: uint16(12),
	})))),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	273: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_count_quantifier),
		Fproduction_id: uint16(14),
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
		Fcount: uint8(1),
	}})),
	275: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_count_quantifier),
		Fproduction_id: uint16(14),
	})))),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	277: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_count_quantifier),
		Fproduction_id: uint16(7),
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
		Fcount: uint8(1),
	}})),
	279: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_count_quantifier),
		Fproduction_id: uint16(7),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	281: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_zero_or_more),
		Fproduction_id: uint16(1),
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
		Fcount: uint8(1),
	}})),
	283: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_zero_or_more),
		Fproduction_id: uint16(1),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fsymbol:        uint16(sym_optional),
		Fproduction_id: uint16(1),
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
		Fcount: uint8(1),
	}})),
	287: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_optional),
		Fproduction_id: uint16(1),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	289: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_one_or_more),
		Fproduction_id: uint16(1),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_one_or_more),
		Fproduction_id: uint16(1),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_term_repeat1),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(62)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(64)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(18)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(106)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(60)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(69)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(74)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(136)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(75)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(93)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(32)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(61)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(95)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(20)),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_character_class_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(118)),
		Frepetition: libc.BoolUint8(1 != 0),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_character_class_repeat1),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_character_class_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(106)),
		Frepetition: libc.BoolUint8(1 != 0),
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
		Fcount: uint8(2),
	}})),
	333: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_character_class_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(61)),
		Frepetition: libc.BoolUint8(1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_character_class_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(69)),
		Frepetition: libc.BoolUint8(1 != 0),
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
		Fcount: uint8(2),
	}})),
	339: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_character_class_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(74)),
		Frepetition: libc.BoolUint8(1 != 0),
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
		Fcount: uint8(2),
	}})),
	342: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_character_class_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(136)),
		Frepetition: libc.BoolUint8(1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_character_class_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(72)),
		Frepetition: libc.BoolUint8(1 != 0),
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
		Fcount: uint8(2),
	}})),
	348: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_character_class_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(75)),
		Frepetition: libc.BoolUint8(1 != 0),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(63)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(66)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(78)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(29)),
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
		Fcount: uint8(1),
	}})),
	359: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(65)),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(99)),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
	371: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(107)),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(30)),
	}})))),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(73)),
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
		Fcount: uint8(1),
	}})),
	377: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_class_range),
		Fproduction_id: uint16(4),
	})))),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(68)),
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
		Fcount: uint8(1),
	}})),
	381: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_class_range),
		Fproduction_id: uint16(6),
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
		Fcount: uint8(2),
	}})),
	383: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_character_class_repeat1),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(80)),
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
		Fcount: uint8(1),
	}})),
	386: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_character_class_repeat1),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_class_range),
		Fproduction_id: uint16(3),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_class_range),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(67)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(37)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(68)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(31)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(70)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(71)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(7)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(97)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(90)),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	410: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_alternation_repeat1),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	412: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_alternation_repeat1),
	})))),
	413: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	415: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_alternation),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	417: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_pattern),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	419: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_alternation),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	421: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	423: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(105)),
	}})))),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	425: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	427: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	429: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(121)),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	431: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_flags),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	433: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(100)),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	435: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(40)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(29)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(131)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(120)),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	445: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(102)),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	447: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(42)),
	}})))),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(125)),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	451: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_unicode_property_value_expression),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	453: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(134)),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(101)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	459: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(48)),
	}})))),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(35)),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	463: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(9)),
	}})))),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	465: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(6)),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	469: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	471: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(41)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(24)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(94)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(33)),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	481: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(112)),
	}})))),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(25)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(22)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(34)),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	489: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_unicode_property_value_expression),
		Fproduction_id: uint16(11),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(117)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(8)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(26)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(43)),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	499: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(114)),
	}})))),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(119)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(23)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(36)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	507: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f uint8_t
		_ [7]byte
	}{f: uint8(TSParseActionTypeAccept)})),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(113)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	511: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_posix_class_name),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(76)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(77)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(45)),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	519: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
}

func tree_sitter_regex(tls *libc.TLS) (r uintptr) {
	return uintptr(unsafe.Pointer(&language))
}

var language = TSLanguage{
	Fabi_version:               uint32(15),
	Fsymbol_count:              uint32(76),
	Falias_count:               uint32(2),
	Ftoken_count:               uint32(47),
	Fstate_count:               uint32(137),
	Flarge_state_count:         uint32(13),
	Fproduction_id_count:       uint32(15),
	Fmax_alias_sequence_length: uint16(7),
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
	Fname:                      __ccgo_ts + 1012,
	Fmetadata: TSLanguageMetadata{
		Fminor_version: uint8(25),
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

var __ccgo_ts1 = "end\x00|\x00any_character\x00^\x00end_assertion\x00boundary_assertion\x00non_boundary_assertion\x00(?\x00=\x00!\x00)\x00(?<\x00pattern_character\x00[\x00-\x00]\x00[:\x00:]\x00posix_class_name_token1\x00identity_escape\x00class_character\x00(\x00(?P<\x00>\x00(?:\x00:\x00*\x00?\x00+\x00{\x00,\x00}\x00\\k\x00<\x00(?P=\x00decimal_escape\x00character_class_escape_token1\x00character_class_escape_token2\x00unicode_character_escape_token1\x00unicode_character_escape_token2\x00unicode_property_value\x00control_escape_token1\x00control_escape_token2\x00control_letter_escape\x00group_name\x00decimal_digits\x00pattern\x00alternation\x00term\x00start_assertion\x00lookaround_assertion\x00_lookahead_assertion\x00_lookbehind_assertion\x00character_class\x00posix_character_class\x00posix_class_name\x00class_range\x00anonymous_capturing_group\x00named_capturing_group\x00non_capturing_group\x00inline_flags_group\x00flags\x00zero_or_more\x00one_or_more\x00optional\x00count_quantifier\x00backreference_escape\x00named_group_backreference\x00character_class_escape\x00unicode_character_escape\x00unicode_property_value_expression\x00control_escape\x00alternation_repeat1\x00term_repeat1\x00character_class_repeat1\x00lazy\x00unicode_property_name\x00regex\x00"
