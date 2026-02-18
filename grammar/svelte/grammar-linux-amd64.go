// Code generated for linux/amd64 by 'ccgo /tmp/grammar-gen-2388767627/preprocessed.c -o /tmp/grammar-gen-2388767627/grammar.go', DO NOT EDIT.

//go:build linux && amd64

package svelte

import (
	"reflect"
	"unsafe"

	"modernc.org/libc"
)

var _ reflect.Type
var _ unsafe.Pointer

const _FORTIFY_SOURCE = 3
const _GNU_SOURCE = 1
const _LIBCPP_HARDENING_MODE = "_LIBCPP_HARDENING_MODE_EXTENSIVE"
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
const __BITINT_MAXWIDTH__ = 65535
const __BYTE_ORDER__ = "__ORDER_LITTLE_ENDIAN__"
const __CCGO__ = 1
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
const __DECIMAL_DIG__ = 17
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
const __GNUC__ = 14
const __GXX_ABI_VERSION = 1019
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
const __LONG_DOUBLE_64__ = 1
const __LONG_LONG_MAX__ = 0x7fffffffffffffff
const __LONG_LONG_WIDTH__ = 64
const __LONG_MAX__ = 0x7fffffffffffffff
const __LONG_WIDTH__ = 64
const __LP64__ = 1
const __MMX_WITH_SSE__ = 1
const __MMX__ = 1
const __OPTIMIZE__ = 1
const __ORDER_BIG_ENDIAN__ = 4321
const __ORDER_LITTLE_ENDIAN__ = 1234
const __ORDER_PDP_ENDIAN__ = 3412
const __PIC__ = 2
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
const __VERSION__ = "14.3.0"
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

type size_t = uint64

type __gnuc_va_list = uintptr

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

type __mbstate_t = struct {
	F__count int32
	F__value struct {
		F__wchb [0][4]int8
		F__wch  uint32
	}
}

type __fpos_t = struct {
	F__pos   __off_t
	F__state __mbstate_t
}

type _G_fpos_t = __fpos_t

type __fpos64_t = struct {
	F__pos   __off64_t
	F__state __mbstate_t
}

type _G_fpos64_t = __fpos64_t

type _IO_FILE = struct {
	F_flags          int32
	F_IO_read_ptr    uintptr
	F_IO_read_end    uintptr
	F_IO_read_base   uintptr
	F_IO_write_base  uintptr
	F_IO_write_ptr   uintptr
	F_IO_write_end   uintptr
	F_IO_buf_base    uintptr
	F_IO_buf_end     uintptr
	F_IO_save_base   uintptr
	F_IO_backup_base uintptr
	F_IO_save_end    uintptr
	F_markers        uintptr
	F_chain          uintptr
	F_fileno         int32
	F_flags2         int32
	F_old_offset     __off_t
	F_cur_column     uint16
	F_vtable_offset  int8
	F_shortbuf       [1]int8
	F_lock           uintptr
	F_offset         __off64_t
	F_codecvt        uintptr
	F_wide_data      uintptr
	F_freeres_list   uintptr
	F_freeres_buf    uintptr
	F_prevchain      uintptr
	F_mode           int32
	F_unused2        [20]int8
}

type __FILE = struct {
	F_flags          int32
	F_IO_read_ptr    uintptr
	F_IO_read_end    uintptr
	F_IO_read_base   uintptr
	F_IO_write_base  uintptr
	F_IO_write_ptr   uintptr
	F_IO_write_end   uintptr
	F_IO_buf_base    uintptr
	F_IO_buf_end     uintptr
	F_IO_save_base   uintptr
	F_IO_backup_base uintptr
	F_IO_save_end    uintptr
	F_markers        uintptr
	F_chain          uintptr
	F_fileno         int32
	F_flags2         int32
	F_old_offset     __off_t
	F_cur_column     uint16
	F_vtable_offset  int8
	F_shortbuf       [1]int8
	F_lock           uintptr
	F_offset         __off64_t
	F_codecvt        uintptr
	F_wide_data      uintptr
	F_freeres_list   uintptr
	F_freeres_buf    uintptr
	F_prevchain      uintptr
	F_mode           int32
	F_unused2        [20]int8
}

type FILE = struct {
	F_flags          int32
	F_IO_read_ptr    uintptr
	F_IO_read_end    uintptr
	F_IO_read_base   uintptr
	F_IO_write_base  uintptr
	F_IO_write_ptr   uintptr
	F_IO_write_end   uintptr
	F_IO_buf_base    uintptr
	F_IO_buf_end     uintptr
	F_IO_save_base   uintptr
	F_IO_backup_base uintptr
	F_IO_save_end    uintptr
	F_markers        uintptr
	F_chain          uintptr
	F_fileno         int32
	F_flags2         int32
	F_old_offset     __off_t
	F_cur_column     uint16
	F_vtable_offset  int8
	F_shortbuf       [1]int8
	F_lock           uintptr
	F_offset         __off64_t
	F_codecvt        uintptr
	F_wide_data      uintptr
	F_freeres_list   uintptr
	F_freeres_buf    uintptr
	F_prevchain      uintptr
	F_mode           int32
	F_unused2        [20]int8
}

type _IO_lock_t = struct{}

type cookie_io_functions_t = struct {
	Fread   uintptr
	Fwrite  uintptr
	Fseek   uintptr
	Fclose1 uintptr
}

type _IO_cookie_io_functions_t = cookie_io_functions_t

type va_list = uintptr

type off_t = int64

type ssize_t = int64

type fpos_t = struct {
	F__pos   __off_t
	F__state __mbstate_t
}

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

type pid_t = int32

type id_t = uint32

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

type int8_t = int8

type int16_t = int16

type int32_t = int32

type int64_t = int64

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
	F__wseq                 __atomic_wide_counter
	F__g1_start             __atomic_wide_counter
	F__g_size               [2]uint32
	F__g1_orig_size         uint32
	F__wrefs                uint32
	F__g_signals            [2]uint32
	F__unused_initialized_1 uint32
	F__unused_initialized_2 uint32
}

type __tss_t = uint32

type __thrd_t = uint64

type __once_flag = struct {
	F__data int32
}

type pthread_t = uint64

type pthread_mutexattr_t = struct {
	F__align [0]int32
	F__size  [4]int8
}

type pthread_condattr_t = struct {
	F__align [0]int32
	F__size  [4]int8
}

type pthread_key_t = uint32

type pthread_once_t = int32

type pthread_attr_t = struct {
	F__align [0]int64
	F__size  [56]int8
}

type pthread_mutex_t = struct {
	F__size  [0][40]int8
	F__align [0]int64
	F__data  __pthread_mutex_s
}

type pthread_cond_t = struct {
	F__size  [0][48]int8
	F__align [0]int64
	F__data  __pthread_cond_s
}

type pthread_rwlock_t = struct {
	F__size  [0][56]int8
	F__align [0]int64
	F__data  __pthread_rwlock_arch_t
}

type pthread_rwlockattr_t = struct {
	F__align [0]int64
	F__size  [8]int8
}

type pthread_spinlock_t = int32

type pthread_barrier_t = struct {
	F__align [0]int64
	F__size  [32]int8
}

type pthread_barrierattr_t = struct {
	F__align [0]int32
	F__size  [4]int8
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

type __locale_struct = struct {
	F__locales       [13]uintptr
	F__ctype_b       uintptr
	F__ctype_tolower uintptr
	F__ctype_toupper uintptr
	F__names         [13]uintptr
}

type __locale_t = uintptr

type locale_t = uintptr

type Array = struct {
	Fcontents uintptr
	Fsize     uint32_t
	Fcapacity uint32_t
}

func _array__delete(tls *libc.TLS, self uintptr) {
	if (*Array)(unsafe.Pointer(self)).Fcontents != 0 {
		libc.Xfree(tls, (*Array)(unsafe.Pointer(self)).Fcontents)
		(*Array)(unsafe.Pointer(self)).Fcontents = libc.UintptrFromInt32(0)
		(*Array)(unsafe.Pointer(self)).Fsize = uint32(0)
		(*Array)(unsafe.Pointer(self)).Fcapacity = uint32(0)
	}
}

func _array__erase(tls *libc.TLS, self uintptr, element_size size_t, index uint32_t) {
	var contents uintptr
	_ = contents
	_ = libc.Uint64FromInt64(4)
	{
		if !(index < (*Array)(unsafe.Pointer(self)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts, __ccgo_ts+19, int32(174), uintptr(unsafe.Pointer(&__func__)))
		}
	}
	contents = (*Array)(unsafe.Pointer(self)).Fcontents
	libc.Xmemmove(tls, contents+uintptr(uint64(index)*element_size), contents+uintptr(uint64(index+libc.Uint32FromInt32(1))*element_size), uint64((*Array)(unsafe.Pointer(self)).Fsize-index-libc.Uint32FromInt32(1))*element_size)
	(*Array)(unsafe.Pointer(self)).Fsize = (*Array)(unsafe.Pointer(self)).Fsize - 1
}

var __func__ = [14]int8{'_', 'a', 'r', 'r', 'a', 'y', '_', '_', 'e', 'r', 'a', 's', 'e'}

func _array__reserve(tls *libc.TLS, self uintptr, element_size size_t, new_capacity uint32_t) {
	if new_capacity > (*Array)(unsafe.Pointer(self)).Fcapacity {
		if (*Array)(unsafe.Pointer(self)).Fcontents != 0 {
			(*Array)(unsafe.Pointer(self)).Fcontents = libc.Xrealloc(tls, (*Array)(unsafe.Pointer(self)).Fcontents, uint64(new_capacity)*element_size)
		} else {
			(*Array)(unsafe.Pointer(self)).Fcontents = libc.Xmalloc(tls, uint64(new_capacity)*element_size)
		}
		(*Array)(unsafe.Pointer(self)).Fcapacity = new_capacity
	}
}

func _array__assign(tls *libc.TLS, self uintptr, other uintptr, element_size size_t) {
	_array__reserve(tls, self, element_size, (*Array)(unsafe.Pointer(other)).Fsize)
	(*Array)(unsafe.Pointer(self)).Fsize = (*Array)(unsafe.Pointer(other)).Fsize
	libc.Xmemcpy(tls, (*Array)(unsafe.Pointer(self)).Fcontents, (*Array)(unsafe.Pointer(other)).Fcontents, uint64((*Array)(unsafe.Pointer(self)).Fsize)*element_size)
}

func _array__swap(tls *libc.TLS, self uintptr, other uintptr) {
	var swap Array
	_ = swap
	swap = *(*Array)(unsafe.Pointer(other))
	*(*Array)(unsafe.Pointer(other)) = *(*Array)(unsafe.Pointer(self))
	*(*Array)(unsafe.Pointer(self)) = swap
}

func _array__grow(tls *libc.TLS, self uintptr, count uint32_t, element_size size_t) {
	var new_capacity, new_size uint32_t
	_, _ = new_capacity, new_size
	new_size = (*Array)(unsafe.Pointer(self)).Fsize + count
	if new_size > (*Array)(unsafe.Pointer(self)).Fcapacity {
		new_capacity = (*Array)(unsafe.Pointer(self)).Fcapacity * uint32(2)
		if new_capacity < uint32(8) {
			new_capacity = uint32(8)
		}
		if new_capacity < new_size {
			new_capacity = new_size
		}
		_array__reserve(tls, self, element_size, new_capacity)
	}
}

func _array__splice(tls *libc.TLS, self uintptr, element_size size_t, index uint32_t, old_count uint32_t, new_count uint32_t, elements uintptr) {
	var contents uintptr
	var new_end, new_size, old_end uint32_t
	_, _, _, _ = contents, new_end, new_size, old_end
	new_size = (*Array)(unsafe.Pointer(self)).Fsize + new_count - old_count
	old_end = index + old_count
	new_end = index + new_count
	_ = libc.Uint64FromInt64(4)
	{
		if !(old_end <= (*Array)(unsafe.Pointer(self)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+74, __ccgo_ts+19, int32(225), uintptr(unsafe.Pointer(&__func__1)))
		}
	}
	_array__reserve(tls, self, element_size, new_size)
	contents = (*Array)(unsafe.Pointer(self)).Fcontents
	if (*Array)(unsafe.Pointer(self)).Fsize > old_end {
		libc.Xmemmove(tls, contents+uintptr(uint64(new_end)*element_size), contents+uintptr(uint64(old_end)*element_size), uint64((*Array)(unsafe.Pointer(self)).Fsize-old_end)*element_size)
	}
	if new_count > uint32(0) {
		if elements != 0 {
			libc.Xmemcpy(tls, contents+uintptr(uint64(index)*element_size), elements, uint64(new_count)*element_size)
		} else {
			libc.Xmemset(tls, contents+uintptr(uint64(index)*element_size), 0, uint64(new_count)*element_size)
		}
	}
	*(*uint32_t)(unsafe.Pointer(self + 8)) += new_count - old_count
}

var __func__1 = [15]int8{'_', 'a', 'r', 'r', 'a', 'y', '_', '_', 's', 'p', 'l', 'i', 'c', 'e'}

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
	Ftype1           TagType
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
		Ftag_name: [16]int8{'b', 'a', 's', 'e', 'f', 'o', 'n', 't'},
		Ftag_type: int32(BASEFONT),
	},
	3: {
		Ftag_name: [16]int8{'b', 'g', 's', 'o', 'u', 'n', 'd'},
		Ftag_type: int32(BGSOUND),
	},
	4: {
		Ftag_name: [16]int8{'b', 'r'},
		Ftag_type: int32(BR),
	},
	5: {
		Ftag_name: [16]int8{'c', 'o', 'l'},
		Ftag_type: int32(COL),
	},
	6: {
		Ftag_name: [16]int8{'c', 'o', 'm', 'm', 'a', 'n', 'd'},
		Ftag_type: int32(COMMAND),
	},
	7: {
		Ftag_name: [16]int8{'e', 'm', 'b', 'e', 'd'},
		Ftag_type: int32(EMBED),
	},
	8: {
		Ftag_name: [16]int8{'f', 'r', 'a', 'm', 'e'},
		Ftag_type: int32(FRAME),
	},
	9: {
		Ftag_name: [16]int8{'h', 'r'},
		Ftag_type: int32(HR),
	},
	10: {
		Ftag_name: [16]int8{'i', 'm', 'a', 'g', 'e'},
		Ftag_type: int32(IMAGE),
	},
	11: {
		Ftag_name: [16]int8{'i', 'm', 'g'},
		Ftag_type: int32(IMG),
	},
	12: {
		Ftag_name: [16]int8{'i', 'n', 'p', 'u', 't'},
		Ftag_type: int32(INPUT),
	},
	13: {
		Ftag_name: [16]int8{'i', 's', 'i', 'n', 'd', 'e', 'x'},
		Ftag_type: int32(ISINDEX),
	},
	14: {
		Ftag_name: [16]int8{'k', 'e', 'y', 'g', 'e', 'n'},
		Ftag_type: int32(KEYGEN),
	},
	15: {
		Ftag_name: [16]int8{'l', 'i', 'n', 'k'},
		Ftag_type: int32(LINK),
	},
	16: {
		Ftag_name: [16]int8{'m', 'e', 'n', 'u', 'i', 't', 'e', 'm'},
		Ftag_type: int32(MENUITEM),
	},
	17: {
		Ftag_name: [16]int8{'m', 'e', 't', 'a'},
		Ftag_type: int32(META),
	},
	18: {
		Ftag_name: [16]int8{'n', 'e', 'x', 't', 'i', 'd'},
		Ftag_type: int32(NEXTID),
	},
	19: {
		Ftag_name: [16]int8{'p', 'a', 'r', 'a', 'm'},
		Ftag_type: int32(PARAM),
	},
	20: {
		Ftag_name: [16]int8{'s', 'o', 'u', 'r', 'c', 'e'},
		Ftag_type: int32(SOURCE),
	},
	21: {
		Ftag_name: [16]int8{'t', 'r', 'a', 'c', 'k'},
		Ftag_type: int32(TRACK),
	},
	22: {
		Ftag_name: [16]int8{'w', 'b', 'r'},
		Ftag_type: int32(WBR),
	},
	23: {
		Ftag_name: [16]int8{'a'},
		Ftag_type: int32(A),
	},
	24: {
		Ftag_name: [16]int8{'a', 'b', 'b', 'r'},
		Ftag_type: int32(ABBR),
	},
	25: {
		Ftag_name: [16]int8{'a', 'd', 'd', 'r', 'e', 's', 's'},
		Ftag_type: int32(ADDRESS),
	},
	26: {
		Ftag_name: [16]int8{'a', 'r', 't', 'i', 'c', 'l', 'e'},
		Ftag_type: int32(ARTICLE),
	},
	27: {
		Ftag_name: [16]int8{'a', 's', 'i', 'd', 'e'},
		Ftag_type: int32(ASIDE),
	},
	28: {
		Ftag_name: [16]int8{'a', 'u', 'd', 'i', 'o'},
		Ftag_type: int32(AUDIO),
	},
	29: {
		Ftag_name: [16]int8{'b'},
		Ftag_type: int32(B),
	},
	30: {
		Ftag_name: [16]int8{'b', 'd', 'i'},
		Ftag_type: int32(BDI),
	},
	31: {
		Ftag_name: [16]int8{'b', 'd', 'o'},
		Ftag_type: int32(BDO),
	},
	32: {
		Ftag_name: [16]int8{'b', 'l', 'o', 'c', 'k', 'q', 'u', 'o', 't', 'e'},
		Ftag_type: int32(BLOCKQUOTE),
	},
	33: {
		Ftag_name: [16]int8{'b', 'o', 'd', 'y'},
		Ftag_type: int32(BODY),
	},
	34: {
		Ftag_name: [16]int8{'b', 'u', 't', 't', 'o', 'n'},
		Ftag_type: int32(BUTTON),
	},
	35: {
		Ftag_name: [16]int8{'c', 'a', 'n', 'v', 'a', 's'},
		Ftag_type: int32(CANVAS),
	},
	36: {
		Ftag_name: [16]int8{'c', 'a', 'p', 't', 'i', 'o', 'n'},
		Ftag_type: int32(CAPTION),
	},
	37: {
		Ftag_name: [16]int8{'c', 'i', 't', 'e'},
		Ftag_type: int32(CITE),
	},
	38: {
		Ftag_name: [16]int8{'c', 'o', 'd', 'e'},
		Ftag_type: int32(CODE),
	},
	39: {
		Ftag_name: [16]int8{'c', 'o', 'l', 'g', 'r', 'o', 'u', 'p'},
		Ftag_type: int32(COLGROUP),
	},
	40: {
		Ftag_name: [16]int8{'d', 'a', 't', 'a'},
		Ftag_type: int32(DATA),
	},
	41: {
		Ftag_name: [16]int8{'d', 'a', 't', 'a', 'l', 'i', 's', 't'},
		Ftag_type: int32(DATALIST),
	},
	42: {
		Ftag_name: [16]int8{'d', 'd'},
		Ftag_type: int32(DD),
	},
	43: {
		Ftag_name: [16]int8{'d', 'e', 'l'},
		Ftag_type: int32(DEL),
	},
	44: {
		Ftag_name: [16]int8{'d', 'e', 't', 'a', 'i', 'l', 's'},
		Ftag_type: int32(DETAILS),
	},
	45: {
		Ftag_name: [16]int8{'d', 'f', 'n'},
		Ftag_type: int32(DFN),
	},
	46: {
		Ftag_name: [16]int8{'d', 'i', 'a', 'l', 'o', 'g'},
		Ftag_type: int32(DIALOG),
	},
	47: {
		Ftag_name: [16]int8{'d', 'i', 'v'},
		Ftag_type: int32(DIV),
	},
	48: {
		Ftag_name: [16]int8{'d', 'l'},
		Ftag_type: int32(DL),
	},
	49: {
		Ftag_name: [16]int8{'d', 't'},
		Ftag_type: int32(DT),
	},
	50: {
		Ftag_name: [16]int8{'e', 'm'},
		Ftag_type: int32(EM),
	},
	51: {
		Ftag_name: [16]int8{'f', 'i', 'e', 'l', 'd', 's', 'e', 't'},
		Ftag_type: int32(FIELDSET),
	},
	52: {
		Ftag_name: [16]int8{'f', 'i', 'g', 'c', 'a', 'p', 't', 'i', 'o', 'n'},
		Ftag_type: int32(FIGCAPTION),
	},
	53: {
		Ftag_name: [16]int8{'f', 'i', 'g', 'u', 'r', 'e'},
		Ftag_type: int32(FIGURE),
	},
	54: {
		Ftag_name: [16]int8{'f', 'o', 'o', 't', 'e', 'r'},
		Ftag_type: int32(FOOTER),
	},
	55: {
		Ftag_name: [16]int8{'f', 'o', 'r', 'm'},
		Ftag_type: int32(FORM),
	},
	56: {
		Ftag_name: [16]int8{'h', '1'},
		Ftag_type: int32(H1),
	},
	57: {
		Ftag_name: [16]int8{'h', '2'},
		Ftag_type: int32(H2),
	},
	58: {
		Ftag_name: [16]int8{'h', '3'},
		Ftag_type: int32(H3),
	},
	59: {
		Ftag_name: [16]int8{'h', '4'},
		Ftag_type: int32(H4),
	},
	60: {
		Ftag_name: [16]int8{'h', '5'},
		Ftag_type: int32(H5),
	},
	61: {
		Ftag_name: [16]int8{'h', '6'},
		Ftag_type: int32(H6),
	},
	62: {
		Ftag_name: [16]int8{'h', 'e', 'a', 'd'},
		Ftag_type: int32(HEAD),
	},
	63: {
		Ftag_name: [16]int8{'h', 'e', 'a', 'd', 'e', 'r'},
		Ftag_type: int32(HEADER),
	},
	64: {
		Ftag_name: [16]int8{'h', 'g', 'r', 'o', 'u', 'p'},
		Ftag_type: int32(HGROUP),
	},
	65: {
		Ftag_name: [16]int8{'h', 't', 'm', 'l'},
		Ftag_type: int32(HTML),
	},
	66: {
		Ftag_name: [16]int8{'i'},
		Ftag_type: int32(I),
	},
	67: {
		Ftag_name: [16]int8{'i', 'f', 'r', 'a', 'm', 'e'},
		Ftag_type: int32(IFRAME),
	},
	68: {
		Ftag_name: [16]int8{'i', 'n', 's'},
		Ftag_type: int32(INS),
	},
	69: {
		Ftag_name: [16]int8{'k', 'b', 'd'},
		Ftag_type: int32(KBD),
	},
	70: {
		Ftag_name: [16]int8{'l', 'a', 'b', 'e', 'l'},
		Ftag_type: int32(LABEL),
	},
	71: {
		Ftag_name: [16]int8{'l', 'e', 'g', 'e', 'n', 'd'},
		Ftag_type: int32(LEGEND),
	},
	72: {
		Ftag_name: [16]int8{'l', 'i'},
		Ftag_type: int32(LI),
	},
	73: {
		Ftag_name: [16]int8{'m', 'a', 'i', 'n'},
		Ftag_type: int32(MAIN),
	},
	74: {
		Ftag_name: [16]int8{'m', 'a', 'p'},
		Ftag_type: int32(MAP),
	},
	75: {
		Ftag_name: [16]int8{'m', 'a', 'r', 'k'},
		Ftag_type: int32(MARK),
	},
	76: {
		Ftag_name: [16]int8{'m', 'a', 't', 'h'},
		Ftag_type: int32(MATH),
	},
	77: {
		Ftag_name: [16]int8{'m', 'e', 'n', 'u'},
		Ftag_type: int32(MENU),
	},
	78: {
		Ftag_name: [16]int8{'m', 'e', 't', 'e', 'r'},
		Ftag_type: int32(METER),
	},
	79: {
		Ftag_name: [16]int8{'n', 'a', 'v'},
		Ftag_type: int32(NAV),
	},
	80: {
		Ftag_name: [16]int8{'n', 'o', 's', 'c', 'r', 'i', 'p', 't'},
		Ftag_type: int32(NOSCRIPT),
	},
	81: {
		Ftag_name: [16]int8{'o', 'b', 'j', 'e', 'c', 't'},
		Ftag_type: int32(OBJECT),
	},
	82: {
		Ftag_name: [16]int8{'o', 'l'},
		Ftag_type: int32(OL),
	},
	83: {
		Ftag_name: [16]int8{'o', 'p', 't', 'g', 'r', 'o', 'u', 'p'},
		Ftag_type: int32(OPTGROUP),
	},
	84: {
		Ftag_name: [16]int8{'o', 'p', 't', 'i', 'o', 'n'},
		Ftag_type: int32(OPTION),
	},
	85: {
		Ftag_name: [16]int8{'o', 'u', 't', 'p', 'u', 't'},
		Ftag_type: int32(OUTPUT),
	},
	86: {
		Ftag_name: [16]int8{'p'},
		Ftag_type: int32(P),
	},
	87: {
		Ftag_name: [16]int8{'p', 'i', 'c', 't', 'u', 'r', 'e'},
		Ftag_type: int32(PICTURE),
	},
	88: {
		Ftag_name: [16]int8{'p', 'r', 'e'},
		Ftag_type: int32(PRE),
	},
	89: {
		Ftag_name: [16]int8{'p', 'r', 'o', 'g', 'r', 'e', 's', 's'},
		Ftag_type: int32(PROGRESS),
	},
	90: {
		Ftag_name: [16]int8{'q'},
		Ftag_type: int32(Q),
	},
	91: {
		Ftag_name: [16]int8{'r', 'b'},
		Ftag_type: int32(RB),
	},
	92: {
		Ftag_name: [16]int8{'r', 'p'},
		Ftag_type: int32(RP),
	},
	93: {
		Ftag_name: [16]int8{'r', 't'},
		Ftag_type: int32(RT),
	},
	94: {
		Ftag_name: [16]int8{'r', 't', 'c'},
		Ftag_type: int32(RTC),
	},
	95: {
		Ftag_name: [16]int8{'r', 'u', 'b', 'y'},
		Ftag_type: int32(RUBY),
	},
	96: {
		Ftag_name: [16]int8{'s'},
		Ftag_type: int32(S),
	},
	97: {
		Ftag_name: [16]int8{'s', 'a', 'm', 'p'},
		Ftag_type: int32(SAMP),
	},
	98: {
		Ftag_name: [16]int8{'s', 'c', 'r', 'i', 'p', 't'},
		Ftag_type: int32(SCRIPT),
	},
	99: {
		Ftag_name: [16]int8{'s', 'e', 'c', 't', 'i', 'o', 'n'},
		Ftag_type: int32(SECTION),
	},
	100: {
		Ftag_name: [16]int8{'s', 'e', 'l', 'e', 'c', 't'},
		Ftag_type: int32(SELECT),
	},
	101: {
		Ftag_name: [16]int8{'s', 'l', 'o', 't'},
		Ftag_type: int32(SLOT),
	},
	102: {
		Ftag_name: [16]int8{'s', 'm', 'a', 'l', 'l'},
		Ftag_type: int32(SMALL),
	},
	103: {
		Ftag_name: [16]int8{'s', 'p', 'a', 'n'},
		Ftag_type: int32(SPAN),
	},
	104: {
		Ftag_name: [16]int8{'s', 't', 'r', 'o', 'n', 'g'},
		Ftag_type: int32(STRONG),
	},
	105: {
		Ftag_name: [16]int8{'s', 't', 'y', 'l', 'e'},
		Ftag_type: int32(STYLE),
	},
	106: {
		Ftag_name: [16]int8{'s', 'u', 'b'},
		Ftag_type: int32(SUB),
	},
	107: {
		Ftag_name: [16]int8{'s', 'u', 'm', 'm', 'a', 'r', 'y'},
		Ftag_type: int32(SUMMARY),
	},
	108: {
		Ftag_name: [16]int8{'s', 'u', 'p'},
		Ftag_type: int32(SUP),
	},
	109: {
		Ftag_name: [16]int8{'s', 'v', 'g'},
		Ftag_type: int32(SVG),
	},
	110: {
		Ftag_name: [16]int8{'t', 'a', 'b', 'l', 'e'},
		Ftag_type: int32(TABLE),
	},
	111: {
		Ftag_name: [16]int8{'t', 'b', 'o', 'd', 'y'},
		Ftag_type: int32(TBODY),
	},
	112: {
		Ftag_name: [16]int8{'t', 'd'},
		Ftag_type: int32(TD),
	},
	113: {
		Ftag_name: [16]int8{'t', 'e', 'm', 'p', 'l', 'a', 't', 'e'},
		Ftag_type: int32(TEMPLATE),
	},
	114: {
		Ftag_name: [16]int8{'t', 'e', 'x', 't', 'a', 'r', 'e', 'a'},
		Ftag_type: int32(TEXTAREA),
	},
	115: {
		Ftag_name: [16]int8{'t', 'f', 'o', 'o', 't'},
		Ftag_type: int32(TFOOT),
	},
	116: {
		Ftag_name: [16]int8{'t', 'h'},
		Ftag_type: int32(TH),
	},
	117: {
		Ftag_name: [16]int8{'t', 'h', 'e', 'a', 'd'},
		Ftag_type: int32(THEAD),
	},
	118: {
		Ftag_name: [16]int8{'t', 'i', 'm', 'e'},
		Ftag_type: int32(TIME),
	},
	119: {
		Ftag_name: [16]int8{'t', 'i', 't', 'l', 'e'},
		Ftag_type: int32(TITLE),
	},
	120: {
		Ftag_name: [16]int8{'t', 'r'},
		Ftag_type: int32(TR),
	},
	121: {
		Ftag_name: [16]int8{'u'},
		Ftag_type: int32(U),
	},
	122: {
		Ftag_name: [16]int8{'u', 'l'},
		Ftag_type: int32(UL),
	},
	123: {
		Ftag_name: [16]int8{'v', 'a', 'r'},
		Ftag_type: int32(VAR),
	},
	124: {
		Ftag_name: [16]int8{'v', 'i', 'd', 'e', 'o'},
		Ftag_type: int32(VIDEO),
	},
	125: {
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

func tag_new(tls *libc.TLS) (r Tag) {
	var tag Tag
	_ = tag
	tag.Ftype1 = int32(END_)
	tag.Fcustom_tag_name = String{}
	return tag
}

func tag_for_name(tls *libc.TLS, _name String) (r Tag) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*String)(unsafe.Pointer(bp)) = _name
	var tag Tag
	_ = tag
	tag = tag_new(tls)
	tag.Ftype1 = tag_type_for_name(tls, bp)
	if tag.Ftype1 == int32(CUSTOM) {
		tag.Fcustom_tag_name = *(*String)(unsafe.Pointer(bp))
	} else {
		_array__delete(tls, bp)
	}
	return tag
}

func tag_free(tls *libc.TLS, self uintptr) {
	if (*Tag)(unsafe.Pointer(self)).Ftype1 == int32(CUSTOM) {
		_array__delete(tls, self+8)
	}
}

func tag_is_void(tls *libc.TLS, self uintptr) (r uint8) {
	return libc.BoolUint8((*Tag)(unsafe.Pointer(self)).Ftype1 < int32(END_OF_VOID_TAGS))
}

func tag_eq(tls *libc.TLS, self uintptr, other uintptr) (r uint8) {
	if (*Tag)(unsafe.Pointer(self)).Ftype1 != (*Tag)(unsafe.Pointer(other)).Ftype1 {
		return libc.BoolUint8(0 != 0)
	}
	if (*Tag)(unsafe.Pointer(self)).Ftype1 == int32(CUSTOM) {
		if (*Tag)(unsafe.Pointer(self)).Fcustom_tag_name.Fsize != (*Tag)(unsafe.Pointer(other)).Fcustom_tag_name.Fsize {
			return libc.BoolUint8(0 != 0)
		}
		if libc.Xmemcmp(tls, (*Tag)(unsafe.Pointer(self)).Fcustom_tag_name.Fcontents, (*Tag)(unsafe.Pointer(other)).Fcustom_tag_name.Fcontents, uint64((*Tag)(unsafe.Pointer(self)).Fcustom_tag_name.Fsize)) != 0 {
			return libc.BoolUint8(0 != 0)
		}
	}
	return libc.BoolUint8(1 != 0)
}

func tag_can_contain(tls *libc.TLS, self uintptr, other uintptr) (r uint8) {
	var child TagType
	var i int32
	_, _ = child, i
	child = (*Tag)(unsafe.Pointer(other)).Ftype1
	switch (*Tag)(unsafe.Pointer(self)).Ftype1 {
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
		return libc.BoolUint8(0 != 0)
	}
	goto _16
_16:
	;
	i = i + 1
	goto _17
_15:
	;
	return libc.BoolUint8(1 != 0)
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
	return libc.BoolUint8(1 != 0)
_14:
	;
	return r
}

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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}
	Ftype1 [0]uint8_t
	Fshift struct {
		Ftype1      uint8_t
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

func set_contains(tls *libc.TLS, ranges uintptr, len1 uint32_t, lookahead int32_t) (r uint8) {
	var half_size, index, mid_index, size uint32_t
	var range1, range11 uintptr
	_, _, _, _, _, _ = half_size, index, mid_index, range1, range11, size
	index = uint32(0)
	size = len1 - index
	for size > uint32(1) {
		half_size = size / uint32(2)
		mid_index = index + half_size
		range1 = ranges + uintptr(mid_index)*8
		if lookahead >= (*TSCharacterRange)(unsafe.Pointer(range1)).Fstart && lookahead <= (*TSCharacterRange)(unsafe.Pointer(range1)).Fend {
			return libc.BoolUint8(1 != 0)
		} else {
			if lookahead > (*TSCharacterRange)(unsafe.Pointer(range1)).Fend {
				index = mid_index
			}
		}
		size = size - half_size
	}
	range11 = ranges + uintptr(index)*8
	return libc.BoolUint8(lookahead >= (*TSCharacterRange)(unsafe.Pointer(range11)).Fstart && lookahead <= (*TSCharacterRange)(unsafe.Pointer(range11)).Fend)
}

type wint_t = uint32

type wctype_t = uint64

const __ISwupper = 0
const __ISwlower = 1
const __ISwalpha = 2
const __ISwdigit = 3
const __ISwxdigit = 4
const __ISwspace = 5
const __ISwprint = 6
const __ISwgraph = 7
const __ISwblank = 8
const __ISwcntrl = 9
const __ISwpunct = 10
const __ISwalnum = 11
const _ISwupper = 16777216
const _ISwlower = 33554432
const _ISwalpha = 67108864
const _ISwdigit = 134217728
const _ISwxdigit = 268435456
const _ISwspace = 536870912
const _ISwprint = 1073741824
const _ISwgraph = -2147483648
const _ISwblank = 65536
const _ISwcntrl = 131072
const _ISwpunct = 262144
const _ISwalnum = 524288

type wctrans_t = uintptr

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
const SVELTE_RAW_TEXT = 9
const SVELTE_RAW_TEXT_EACH = 10
const SVELTE_RAW_TEXT_SNIPPET_ARGUMENTS = 11
const AT = 12
const HASH = 13
const SLASH = 14
const COLON = 15

type Scanner = struct {
	Ftags struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	}
}

func advance(tls *libc.TLS, lexer uintptr) {
	(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(0 != 0))
}

func skip(tls *libc.TLS, lexer uintptr) {
	(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(1 != 0))
}

func serialize(tls *libc.TLS, scanner uintptr, buffer uintptr) (r uint32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var name_length, size, v1 uint32
	var tag Tag
	var _ /* serialized_tag_count at bp+2 */ uint16_t
	var _ /* tag_count at bp+0 */ uint16_t
	_, _, _, _ = name_length, size, tag, v1
	if (*Scanner)(unsafe.Pointer(scanner)).Ftags.Fsize > libc.Uint32FromInt32(libc.Int32FromInt32(65535)) {
		v1 = libc.Uint32FromInt32(libc.Int32FromInt32(65535))
	} else {
		v1 = (*Scanner)(unsafe.Pointer(scanner)).Ftags.Fsize
	}
	*(*uint16_t)(unsafe.Pointer(bp)) = uint16(v1)
	*(*uint16_t)(unsafe.Pointer(bp + 2)) = uint16(0)
	size = uint32(2)
	libc.Xmemcpy(tls, buffer+uintptr(size), bp, uint64(2))
	size = uint32(uint64(size) + libc.Uint64FromInt64(2))
	for {
		if !(libc.Int32FromUint16(*(*uint16_t)(unsafe.Pointer(bp + 2))) < libc.Int32FromUint16(*(*uint16_t)(unsafe.Pointer(bp)))) {
			break
		}
		tag = *(*Tag)(unsafe.Pointer((*Scanner)(unsafe.Pointer(scanner)).Ftags.Fcontents + uintptr(*(*uint16_t)(unsafe.Pointer(bp + 2)))*24))
		if tag.Ftype1 == int32(CUSTOM) {
			name_length = tag.Fcustom_tag_name.Fsize
			if name_length > libc.Uint32FromInt32(libc.Int32FromInt32(255)) {
				name_length = libc.Uint32FromInt32(libc.Int32FromInt32(255))
			}
			if size+uint32(2)+name_length >= uint32(1024) {
				break
			}
			v1 = size
			size = size + 1
			*(*int8)(unsafe.Pointer(buffer + uintptr(v1))) = int8(tag.Ftype1)
			v1 = size
			size = size + 1
			*(*int8)(unsafe.Pointer(buffer + uintptr(v1))) = libc.Int8FromUint32(name_length)
			libc.Xstrncpy(tls, buffer+uintptr(size), tag.Fcustom_tag_name.Fcontents, uint64(name_length))
			size = size + name_length
		} else {
			if size+uint32(1) >= uint32(1024) {
				break
			}
			v1 = size
			size = size + 1
			*(*int8)(unsafe.Pointer(buffer + uintptr(v1))) = int8(tag.Ftype1)
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
	var i, iter, size, v3 uint32
	var name_length uint16_t
	var v5 uint32_t
	var v6 uintptr
	var _ /* serialized_tag_count at bp+2 */ uint16_t
	var _ /* tag at bp+8 */ Tag
	var _ /* tag_count at bp+0 */ uint16_t
	_, _, _, _, _, _, _ = i, iter, name_length, size, v3, v5, v6
	i = uint32(0)
	for {
		if !(i < (*Scanner)(unsafe.Pointer(scanner)).Ftags.Fsize) {
			break
		}
		tag_free(tls, (*Scanner)(unsafe.Pointer(scanner)).Ftags.Fcontents+uintptr(i)*24)
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
		_array__reserve(tls, scanner, libc.Uint64FromInt64(24), uint32(*(*uint16_t)(unsafe.Pointer(bp))))
		if libc.Int32FromUint16(*(*uint16_t)(unsafe.Pointer(bp))) > 0 {
			iter = uint32(0)
			iter = uint32(0)
			for {
				if !(iter < uint32(*(*uint16_t)(unsafe.Pointer(bp + 2)))) {
					break
				}
				*(*Tag)(unsafe.Pointer(bp + 8)) = tag_new(tls)
				v3 = size
				size = size + 1
				(*(*Tag)(unsafe.Pointer(bp + 8))).Ftype1 = int32(*(*int8)(unsafe.Pointer(buffer + uintptr(v3))))
				if (*(*Tag)(unsafe.Pointer(bp + 8))).Ftype1 == int32(CUSTOM) {
					v3 = size
					size = size + 1
					name_length = uint16(libc.Uint8FromInt8(*(*int8)(unsafe.Pointer(buffer + uintptr(v3)))))
					_array__reserve(tls, bp+8+8, libc.Uint64FromInt64(1), uint32(name_length))
					(*(*Tag)(unsafe.Pointer(bp + 8))).Fcustom_tag_name.Fsize = uint32(name_length)
					libc.Xmemcpy(tls, (*(*Tag)(unsafe.Pointer(bp + 8))).Fcustom_tag_name.Fcontents, buffer+uintptr(size), uint64(name_length))
					size = size + uint32(name_length)
				}
				_array__grow(tls, scanner, uint32(1), libc.Uint64FromInt64(24))
				v6 = scanner + 8
				v5 = *(*uint32_t)(unsafe.Pointer(v6))
				*(*uint32_t)(unsafe.Pointer(v6)) = *(*uint32_t)(unsafe.Pointer(v6)) + 1
				*(*Tag)(unsafe.Pointer((*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(scanner)).Fcontents + uintptr(v5)*24)) = *(*Tag)(unsafe.Pointer(bp + 8))
				goto _2
			_2:
				;
				iter = iter + 1
			}
			for {
				if !(iter < uint32(*(*uint16_t)(unsafe.Pointer(bp)))) {
					break
				}
				_array__grow(tls, scanner, uint32(1), libc.Uint64FromInt64(24))
				v6 = scanner + 8
				v5 = *(*uint32_t)(unsafe.Pointer(v6))
				*(*uint32_t)(unsafe.Pointer(v6)) = *(*uint32_t)(unsafe.Pointer(v6)) + 1
				*(*Tag)(unsafe.Pointer((*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(scanner)).Fcontents + uintptr(v5)*24)) = tag_new(tls)
				goto _7
			_7:
				;
				iter = iter + 1
			}
		}
	}
}

func scan_tag_name(tls *libc.TLS, lexer uintptr) (r String) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var v1 uint32_t
	var v2 uintptr
	var _ /* tag_name at bp+0 */ String
	_, _ = v1, v2
	*(*String)(unsafe.Pointer(bp)) = String{}
	for libc.Xiswalnum(tls, libc.Uint32FromInt32((*TSLexer)(unsafe.Pointer(lexer)).Flookahead)) != 0 || (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('-') || (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32(':') || (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('.') {
		_array__grow(tls, bp, uint32(1), libc.Uint64FromInt64(1))
		v2 = bp + 8
		v1 = *(*uint32_t)(unsafe.Pointer(v2))
		*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) + 1
		*(*int8)(unsafe.Pointer((*String)(unsafe.Pointer(bp)).Fcontents + uintptr(v1))) = int8((*TSLexer)(unsafe.Pointer(lexer)).Flookahead)
		advance(tls, lexer)
	}
	return *(*String)(unsafe.Pointer(bp))
}

func scan_comment(tls *libc.TLS, lexer uintptr) (r uint8) {
	var dashes uint32
	_ = dashes
	if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32('-') {
		return libc.BoolUint8(0 != 0)
	}
	advance(tls, lexer)
	if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32('-') {
		return libc.BoolUint8(0 != 0)
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
				return libc.BoolUint8(1 != 0)
			}
			dashes = uint32(0)
		default:
			dashes = uint32(0)
		}
		advance(tls, lexer)
	}
	return libc.BoolUint8(0 != 0)
}

func scan_javascript_block_comment(tls *libc.TLS, lexer uintptr) (r uint8) {
	if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32('*') {
		return libc.BoolUint8(0 != 0)
	}
	advance(tls, lexer)
	for (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != 0 {
		switch (*TSLexer)(unsafe.Pointer(lexer)).Flookahead {
		case int32('*'):
			advance(tls, lexer)
			if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('/') {
				advance(tls, lexer)
				return libc.BoolUint8(1 != 0)
			}
		default:
			advance(tls, lexer)
		}
	}
	return libc.BoolUint8(0 != 0)
}

func scan_javascript_line_comment(tls *libc.TLS, lexer uintptr) (r uint8) {
	if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32('/') {
		return libc.BoolUint8(0 != 0)
	}
	advance(tls, lexer)
	for (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != 0 {
		switch (*TSLexer)(unsafe.Pointer(lexer)).Flookahead {
		case int32('\n'):
			fallthrough
		case int32('\r'):
			advance(tls, lexer)
			return libc.BoolUint8(1 != 0)
		default:
			advance(tls, lexer)
		}
	}
	return libc.BoolUint8(0 != 0)
}

func scan_javascript_balanced_brace(tls *libc.TLS, lexer uintptr) (r uint8) {
	var brace_level uint8_t
	_ = brace_level
	if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32('{') {
		return libc.BoolUint8(0 != 0)
	}
	brace_level = uint8(0)
	advance(tls, lexer)
	for (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != 0 {
		switch (*TSLexer)(unsafe.Pointer(lexer)).Flookahead {
		case int32('`'):
			scan_javascript_template_string(tls, lexer)
		case int32('\\'):
			advance(tls, lexer)
			advance(tls, lexer)
		case int32('\''):
			fallthrough
		case int32('"'):
			scan_javascript_quoted_string(tls, lexer, (*TSLexer)(unsafe.Pointer(lexer)).Flookahead)
		case int32('{'):
			brace_level = brace_level + 1
			advance(tls, lexer)
		case int32('}'):
			advance(tls, lexer)
			if libc.Int32FromUint8(brace_level) == 0 {
				return libc.BoolUint8(1 != 0)
			}
			brace_level = brace_level - 1
		default:
			advance(tls, lexer)
		}
	}
	return libc.BoolUint8(0 != 0)
}

func scan_javascript_quoted_string(tls *libc.TLS, lexer uintptr, delimiter int32_t) (r uint8) {
	if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != delimiter {
		return libc.BoolUint8(0 != 0)
	}
	advance(tls, lexer)
	for (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != 0 {
		switch (*TSLexer)(unsafe.Pointer(lexer)).Flookahead {
		case int32('\\'):
			advance(tls, lexer)
			advance(tls, lexer)
		default:
			if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == delimiter {
				advance(tls, lexer)
				return libc.BoolUint8(1 != 0)
			}
			advance(tls, lexer)
		}
	}
	return libc.BoolUint8(0 != 0)
}

func scan_javascript_template_string(tls *libc.TLS, lexer uintptr) (r uint8) {
	if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32('`') {
		return libc.BoolUint8(0 != 0)
	}
	advance(tls, lexer)
	for (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != 0 {
		switch (*TSLexer)(unsafe.Pointer(lexer)).Flookahead {
		case int32('$'):
			advance(tls, lexer)
			if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('{') {
				scan_javascript_balanced_brace(tls, lexer)
			}
		case int32('\\'):
			advance(tls, lexer)
			advance(tls, lexer)
		case int32('`'):
			advance(tls, lexer)
			return libc.BoolUint8(1 != 0)
		default:
			advance(tls, lexer)
		}
	}
	return libc.BoolUint8(0 != 0)
}

func scan_raw_text(tls *libc.TLS, scanner uintptr, lexer uintptr) (r uint8) {
	var delimiter_index uint32
	var end_delimiter, v1 uintptr
	_, _, _ = delimiter_index, end_delimiter, v1
	if (*Scanner)(unsafe.Pointer(scanner)).Ftags.Fsize == uint32(0) {
		return libc.BoolUint8(0 != 0)
	}
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
	_ = libc.Uint64FromInt64(4)
	{
		if !((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(scanner)).Fsize-libc.Uint32FromInt32(1) < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(scanner)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+96, __ccgo_ts+160, int32(319), uintptr(unsafe.Pointer(&__func__2)))
		}
	}
	if (*Tag)(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(scanner)).Fcontents+uintptr((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(scanner)).Fsize-uint32(1))*24)).Ftype1 == int32(SCRIPT) {
		v1 = __ccgo_ts + 199
	} else {
		v1 = __ccgo_ts + 208
	}
	end_delimiter = v1
	delimiter_index = uint32(0)
	for (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != 0 {
		if int32(libc.Int8FromUint32(libc.Xtowupper(tls, libc.Uint32FromInt32((*TSLexer)(unsafe.Pointer(lexer)).Flookahead)))) == int32(*(*int8)(unsafe.Pointer(end_delimiter + uintptr(delimiter_index)))) {
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
	return libc.BoolUint8(1 != 0)
}

var __func__2 = [14]int8{'s', 'c', 'a', 'n', '_', 'r', 'a', 'w', '_', 't', 'e', 'x', 't'}

func scan_svelte_raw_text_snippet(tls *libc.TLS, lexer uintptr) (r uint8) {
	var advanced_once uint8
	var paren_level uint8_t
	_, _ = advanced_once, paren_level
	for libc.Xiswspace(tls, libc.Uint32FromInt32((*TSLexer)(unsafe.Pointer(lexer)).Flookahead)) != 0 {
		skip(tls, lexer)
	}
	(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(SVELTE_RAW_TEXT_SNIPPET_ARGUMENTS)
	paren_level = uint8(0)
	advanced_once = libc.BoolUint8(0 != 0)
	for !((*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Feof})))(tls, lexer) != 0) {
		switch (*TSLexer)(unsafe.Pointer(lexer)).Flookahead {
		case int32('/'):
			advance(tls, lexer)
			if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('*') {
				scan_javascript_block_comment(tls, lexer)
			} else {
				if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('/') {
					scan_javascript_line_comment(tls, lexer)
				}
			}
		case int32('\\'):
			advance(tls, lexer)
		case int32('"'):
			fallthrough
		case int32('\''):
			scan_javascript_quoted_string(tls, lexer, (*TSLexer)(unsafe.Pointer(lexer)).Flookahead)
		case int32('`'):
			scan_javascript_template_string(tls, lexer)
		case int32(')'):
			if libc.Int32FromUint8(paren_level) == 0 {
				(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
				return advanced_once
			}
			advance(tls, lexer)
			paren_level = paren_level - 1
		case int32('('):
			advance(tls, lexer)
			paren_level = paren_level + 1
		default:
			advance(tls, lexer)
			break
		}
		advanced_once = libc.BoolUint8(1 != 0)
	}
	return libc.BoolUint8(0 != 0)
}

func scan_svelte_raw_text(tls *libc.TLS, lexer uintptr, valid_symbols uintptr) (r uint8) {
	var advanced_once, has_sigil uint8
	var brace_level uint8_t
	var v1 int32
	_, _, _, _ = advanced_once, brace_level, has_sigil, v1
	for libc.Xiswspace(tls, libc.Uint32FromInt32((*TSLexer)(unsafe.Pointer(lexer)).Flookahead)) != 0 {
		skip(tls, lexer)
	}
	if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('@') && *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(AT))) != 0 || (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('#') && *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(HASH))) != 0 || (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32(':') && *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(COLON))) != 0 {
		return libc.BoolUint8(0 != 0)
	}
	has_sigil = libc.BoolUint8((*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('@') || (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('#') || (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32(':'))
	if has_sigil != 0 {
		return libc.BoolUint8(0 != 0)
	}
	advanced_once = libc.BoolUint8(0 != 0)
	if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('/') && *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(SLASH))) != 0 {
		advance(tls, lexer)
		if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('*') {
			return scan_javascript_block_comment(tls, lexer)
		}
		if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32('/') {
			return libc.BoolUint8(0 != 0)
		}
		advanced_once = libc.BoolUint8(1 != 0)
	}
	if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(SVELTE_RAW_TEXT_EACH))) != 0 {
		v1 = int32(SVELTE_RAW_TEXT_EACH)
	} else {
		v1 = int32(SVELTE_RAW_TEXT)
	}
	(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = libc.Uint16FromInt32(v1)
	brace_level = uint8(0)
	for !((*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Feof})))(tls, lexer) != 0) {
		switch (*TSLexer)(unsafe.Pointer(lexer)).Flookahead {
		case int32('/'):
			advance(tls, lexer)
			advanced_once = libc.BoolUint8(1 != 0)
			if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('*') {
				scan_javascript_block_comment(tls, lexer)
			} else {
				if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('/') {
					scan_javascript_line_comment(tls, lexer)
				}
			}
		case int32('\\'):
			advance(tls, lexer)
			advanced_once = libc.BoolUint8(1 != 0)
		case int32('"'):
			fallthrough
		case int32('\''):
			scan_javascript_quoted_string(tls, lexer, (*TSLexer)(unsafe.Pointer(lexer)).Flookahead)
			advanced_once = libc.BoolUint8(1 != 0)
		case int32('`'):
			scan_javascript_template_string(tls, lexer)
			advanced_once = libc.BoolUint8(1 != 0)
		case int32('}'):
			if libc.Int32FromUint8(brace_level) == 0 {
				(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
				return advanced_once
			}
			advance(tls, lexer)
			brace_level = brace_level - 1
			advanced_once = libc.BoolUint8(1 != 0)
		case int32('{'):
			advance(tls, lexer)
			brace_level = brace_level + 1
			advanced_once = libc.BoolUint8(1 != 0)
		case int32('a'):
			if libc.Int32FromUint16((*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol) == int32(SVELTE_RAW_TEXT_EACH) {
				(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
				advance(tls, lexer)
				advanced_once = libc.BoolUint8(1 != 0)
				if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('s') {
					advance(tls, lexer)
					if libc.Xiswspace(tls, libc.Uint32FromInt32((*TSLexer)(unsafe.Pointer(lexer)).Flookahead)) != 0 {
						return advanced_once
					}
				}
			} else {
				advance(tls, lexer)
				advanced_once = libc.BoolUint8(1 != 0)
			}
		default:
			advance(tls, lexer)
			advanced_once = libc.BoolUint8(1 != 0)
			break
		}
	}
	return libc.BoolUint8(0 != 0)
}

func pop_tag(tls *libc.TLS, scanner uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var v1 uint32_t
	var v2 uintptr
	var _ /* popped_tag at bp+0 */ Tag
	_, _ = v1, v2
	v2 = scanner + 8
	*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) - 1
	v1 = *(*uint32_t)(unsafe.Pointer(v2))
	*(*Tag)(unsafe.Pointer(bp)) = *(*Tag)(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(scanner)).Fcontents + uintptr(v1)*24))
	tag_free(tls, bp)
}

func scan_implicit_end_tag(tls *libc.TLS, scanner uintptr, lexer uintptr) (r uint8) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var i uint32
	var is_closing_tag uint8
	var parent, v1 uintptr
	var v2 bool
	var _ /* next_tag at bp+16 */ Tag
	var _ /* tag_name at bp+0 */ String
	_, _, _, _, _ = i, is_closing_tag, parent, v1, v2
	if (*Scanner)(unsafe.Pointer(scanner)).Ftags.Fsize == uint32(0) {
		v1 = libc.UintptrFromInt32(0)
	} else {
		_ = libc.Uint64FromInt64(4)
		{
			if !((*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(scanner)).Fsize-libc.Uint32FromInt32(1) < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(scanner)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+96, __ccgo_ts+160, int32(517), uintptr(unsafe.Pointer(&__func__3)))
			}
		}
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
	is_closing_tag = libc.BoolUint8(0 != 0)
	if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('/') {
		is_closing_tag = libc.BoolUint8(1 != 0)
		advance(tls, lexer)
	} else {
		if parent != 0 && tag_is_void(tls, parent) != 0 {
			pop_tag(tls, scanner)
			(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(IMPLICIT_END_TAG)
			return libc.BoolUint8(1 != 0)
		}
	}
	*(*String)(unsafe.Pointer(bp)) = scan_tag_name(tls, lexer)
	if (*(*String)(unsafe.Pointer(bp))).Fsize == uint32(0) && !((*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Feof})))(tls, lexer) != 0) {
		_array__delete(tls, bp)
		return libc.BoolUint8(0 != 0)
	}
	*(*Tag)(unsafe.Pointer(bp + 16)) = tag_for_name(tls, *(*String)(unsafe.Pointer(bp)))
	if is_closing_tag != 0 {
		if v2 = (*Scanner)(unsafe.Pointer(scanner)).Ftags.Fsize > uint32(0); v2 {
			_ = libc.Uint64FromInt64(4)
			{
				if !((*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(scanner)).Fsize-libc.Uint32FromInt32(1) < (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(scanner)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+96, __ccgo_ts+160, int32(541), uintptr(unsafe.Pointer(&__func__3)))
				}
			}
		}
		if v2 && tag_eq(tls, (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(scanner)).Fcontents+uintptr((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(scanner)).Fsize-uint32(1))*24, bp+16) != 0 {
			tag_free(tls, bp+16)
			return libc.BoolUint8(0 != 0)
		}
		i = (*Scanner)(unsafe.Pointer(scanner)).Ftags.Fsize
		for {
			if !(i > uint32(0)) {
				break
			}
			if (*(*Tag)(unsafe.Pointer((*Scanner)(unsafe.Pointer(scanner)).Ftags.Fcontents + uintptr(i-uint32(1))*24))).Ftype1 == (*(*Tag)(unsafe.Pointer(bp + 16))).Ftype1 {
				pop_tag(tls, scanner)
				(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(IMPLICIT_END_TAG)
				tag_free(tls, bp+16)
				return libc.BoolUint8(1 != 0)
			}
			goto _3
		_3:
			;
			i = i - 1
		}
	} else {
		if parent != 0 && (!(tag_can_contain(tls, parent, bp+16) != 0) || ((*Tag)(unsafe.Pointer(parent)).Ftype1 == int32(HTML) || (*Tag)(unsafe.Pointer(parent)).Ftype1 == int32(HEAD) || (*Tag)(unsafe.Pointer(parent)).Ftype1 == int32(BODY)) && (*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Feof})))(tls, lexer) != 0) {
			pop_tag(tls, scanner)
			(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(IMPLICIT_END_TAG)
			tag_free(tls, bp+16)
			return libc.BoolUint8(1 != 0)
		}
	}
	tag_free(tls, bp+16)
	return libc.BoolUint8(0 != 0)
}

var __func__3 = [22]int8{'s', 'c', 'a', 'n', '_', 'i', 'm', 'p', 'l', 'i', 'c', 'i', 't', '_', 'e', 'n', 'd', '_', 't', 'a', 'g'}

func scan_start_tag_name(tls *libc.TLS, scanner uintptr, lexer uintptr) (r uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var tag Tag
	var v1 uint32_t
	var v2 uintptr
	var _ /* tag_name at bp+0 */ String
	_, _, _ = tag, v1, v2
	*(*String)(unsafe.Pointer(bp)) = scan_tag_name(tls, lexer)
	if (*(*String)(unsafe.Pointer(bp))).Fsize == uint32(0) {
		_array__delete(tls, bp)
		return libc.BoolUint8(0 != 0)
	}
	tag = tag_for_name(tls, *(*String)(unsafe.Pointer(bp)))
	_array__grow(tls, scanner, uint32(1), libc.Uint64FromInt64(24))
	v2 = scanner + 8
	v1 = *(*uint32_t)(unsafe.Pointer(v2))
	*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) + 1
	*(*Tag)(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(scanner)).Fcontents + uintptr(v1)*24)) = tag
	switch tag.Ftype1 {
	case int32(SCRIPT):
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(SCRIPT_START_TAG_NAME)
	case int32(STYLE):
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(STYLE_START_TAG_NAME)
	default:
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(START_TAG_NAME)
		break
	}
	return libc.BoolUint8(1 != 0)
}

func scan_end_tag_name(tls *libc.TLS, scanner uintptr, lexer uintptr) (r uint8) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var v1 bool
	var _ /* tag at bp+16 */ Tag
	var _ /* tag_name at bp+0 */ String
	_ = v1
	*(*String)(unsafe.Pointer(bp)) = scan_tag_name(tls, lexer)
	if (*(*String)(unsafe.Pointer(bp))).Fsize == uint32(0) {
		_array__delete(tls, bp)
		return libc.BoolUint8(0 != 0)
	}
	*(*Tag)(unsafe.Pointer(bp + 16)) = tag_for_name(tls, *(*String)(unsafe.Pointer(bp)))
	if v1 = (*Scanner)(unsafe.Pointer(scanner)).Ftags.Fsize > uint32(0); v1 {
		_ = libc.Uint64FromInt64(4)
		{
			if !((*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(scanner)).Fsize-libc.Uint32FromInt32(1) < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(scanner)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+96, __ccgo_ts+160, int32(603), uintptr(unsafe.Pointer(&__func__4)))
			}
		}
	}
	if v1 && tag_eq(tls, (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(scanner)).Fcontents+uintptr((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(scanner)).Fsize-uint32(1))*24, bp+16) != 0 {
		pop_tag(tls, scanner)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(END_TAG_NAME)
	} else {
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(ERRONEOUS_END_TAG_NAME)
	}
	tag_free(tls, bp+16)
	return libc.BoolUint8(1 != 0)
}

var __func__4 = [18]int8{'s', 'c', 'a', 'n', '_', 'e', 'n', 'd', '_', 't', 'a', 'g', '_', 'n', 'a', 'm', 'e'}

func scan_self_closing_tag_delimiter(tls *libc.TLS, scanner uintptr, lexer uintptr) (r uint8) {
	advance(tls, lexer)
	if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('>') {
		advance(tls, lexer)
		if (*Scanner)(unsafe.Pointer(scanner)).Ftags.Fsize > uint32(0) {
			pop_tag(tls, scanner)
			(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(SELF_CLOSING_TAG_DELIMITER)
		}
		return libc.BoolUint8(1 != 0)
	}
	return libc.BoolUint8(0 != 0)
}

func scan(tls *libc.TLS, scanner uintptr, lexer uintptr, valid_symbols uintptr) (r uint8) {
	var v1 int32
	_ = v1
	if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(RAW_TEXT))) != 0 && !(*(*uint8)(unsafe.Pointer(valid_symbols + uintptr(START_TAG_NAME))) != 0) && !(*(*uint8)(unsafe.Pointer(valid_symbols + uintptr(END_TAG_NAME))) != 0) {
		return scan_raw_text(tls, scanner, lexer)
	}
	if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(SVELTE_RAW_TEXT_SNIPPET_ARGUMENTS))) != 0 {
		return scan_svelte_raw_text_snippet(tls, lexer)
	}
	if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(SVELTE_RAW_TEXT))) != 0 || *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(SVELTE_RAW_TEXT_EACH))) != 0 {
		return scan_svelte_raw_text(tls, lexer, valid_symbols)
	}
	for libc.Xiswspace(tls, libc.Uint32FromInt32((*TSLexer)(unsafe.Pointer(lexer)).Flookahead)) != 0 {
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
	case int32('{'):
		fallthrough
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
				v1 = libc.Int32FromUint8(scan_start_tag_name(tls, scanner, lexer))
			} else {
				v1 = libc.Int32FromUint8(scan_end_tag_name(tls, scanner, lexer))
			}
			return libc.Uint8FromInt32(libc.BoolInt32(v1 != 0))
		}
	}
	return libc.BoolUint8(0 != 0)
}

func tree_sitter_svelte_external_scanner_create(tls *libc.TLS) (r uintptr) {
	var scanner uintptr
	_ = scanner
	scanner = libc.Xcalloc(tls, uint64(1), uint64(16))
	return scanner
}

func tree_sitter_svelte_external_scanner_scan(tls *libc.TLS, payload uintptr, lexer uintptr, valid_symbols uintptr) (r uint8) {
	var scanner uintptr
	_ = scanner
	scanner = payload
	return scan(tls, scanner, lexer, valid_symbols)
}

func tree_sitter_svelte_external_scanner_serialize(tls *libc.TLS, payload uintptr, buffer uintptr) (r uint32) {
	var scanner uintptr
	_ = scanner
	scanner = payload
	return serialize(tls, scanner, buffer)
}

func tree_sitter_svelte_external_scanner_deserialize(tls *libc.TLS, payload uintptr, buffer uintptr, length uint32) {
	var scanner uintptr
	_ = scanner
	scanner = payload
	deserialize(tls, scanner, buffer, length)
}

func tree_sitter_svelte_external_scanner_destroy(tls *libc.TLS, payload uintptr) {
	var i uint32
	var scanner uintptr
	_, _ = i, scanner
	scanner = payload
	i = uint32(0)
	for {
		if !(i < (*Scanner)(unsafe.Pointer(scanner)).Ftags.Fsize) {
			break
		}
		_ = libc.Uint64FromInt64(4)
		{
			if !(i < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(scanner)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+216, __ccgo_ts+160, int32(705), uintptr(unsafe.Pointer(&__func__5)))
			}
		}
		tag_free(tls, (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(scanner)).Fcontents+uintptr(i)*24)
		goto _1
	_1:
		;
		i = i + 1
	}
	_array__delete(tls, scanner)
	libc.Xfree(tls, scanner)
}

var __func__5 = [44]int8{'t', 'r', 'e', 'e', '_', 's', 'i', 't', 't', 'e', 'r', '_', 's', 'v', 'e', 'l', 't', 'e', '_', 'e', 'x', 't', 'e', 'r', 'n', 'a', 'l', '_', 's', 'c', 'a', 'n', 'n', 'e', 'r', '_', 'd', 'e', 's', 't', 'r', 'o', 'y'}

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
const anon_sym_DQUOTE = 13
const sym_text = 14
const aux_sym__single_quoted_attribute_value_token1 = 15
const aux_sym__double_quoted_attribute_value_token1 = 16
const anon_sym_POUND = 17
const anon_sym_if = 18
const anon_sym_LBRACE = 19
const anon_sym_RBRACE = 20
const anon_sym_COLON = 21
const anon_sym_elseif = 22
const anon_sym_else = 23
const anon_sym_SLASH = 24
const anon_sym_each = 25
const anon_sym_as = 26
const anon_sym_await = 27
const anon_sym_then = 28
const anon_sym_catch = 29
const anon_sym_key = 30
const anon_sym_snippet = 31
const aux_sym_snippet_start_token1 = 32
const anon_sym_LPAREN = 33
const anon_sym_RPAREN = 34
const aux_sym__tag_value_token1 = 35
const anon_sym_AT = 36
const anon_sym_html = 37
const anon_sym_const = 38
const anon_sym_debug = 39
const anon_sym_render = 40
const sym__start_tag_name = 41
const sym__script_start_tag_name = 42
const sym__style_start_tag_name = 43
const sym__end_tag_name = 44
const sym_erroneous_end_tag_name = 45
const sym__implicit_end_tag = 46
const sym_raw_text = 47
const sym_comment = 48
const sym_svelte_raw_text = 49
const sym_svelte_raw_text_each = 50
const sym_svelte_raw_text_snippet_arguments = 51
const sym_document = 52
const sym_doctype = 53
const sym__node = 54
const sym_element = 55
const sym_script_element = 56
const sym_style_element = 57
const sym_start_tag = 58
const sym_script_start_tag = 59
const sym_style_start_tag = 60
const sym_self_closing_tag = 61
const sym_end_tag = 62
const sym_erroneous_end_tag = 63
const sym_attribute = 64
const sym_quoted_attribute_value = 65
const aux_sym__single_quoted_attribute_value = 66
const aux_sym__double_quoted_attribute_value = 67
const sym_if_statement = 68
const sym__if_start_tag = 69
const sym_if_start = 70
const sym__else_if_tag = 71
const sym_else_if_start = 72
const sym_else_if_block = 73
const sym__else_tag = 74
const sym_else_start = 75
const sym_else_block = 76
const sym__if_end_tag = 77
const sym_if_end = 78
const sym_each_statement = 79
const sym__each_start_tag = 80
const sym_each_start = 81
const sym__each_end_tag = 82
const sym_each_end = 83
const sym_await_statement = 84
const sym__await_start_tag = 85
const sym_await_start = 86
const sym__then_tag = 87
const sym_then_start = 88
const sym_then_block = 89
const sym__catch_tag = 90
const sym_catch_start = 91
const sym_catch_block = 92
const sym__await_end_tag = 93
const sym_await_end = 94
const sym_key_statement = 95
const sym__key_start_tag = 96
const sym_key_start = 97
const sym__key_end_tag = 98
const sym_key_end = 99
const sym_snippet_statement = 100
const sym__snippet_start_tag = 101
const sym_snippet_start = 102
const sym__snippet_end_tag = 103
const sym_snippet_end = 104
const sym_expression = 105
const sym__tag_value = 106
const sym__html_tag = 107
const sym_html_tag = 108
const sym__const_tag = 109
const sym_const_tag = 110
const sym__debug_tag = 111
const sym_debug_tag = 112
const sym__render_tag = 113
const sym_render_tag = 114
const aux_sym_document_repeat1 = 115
const aux_sym_start_tag_repeat1 = 116
const aux_sym_if_statement_repeat1 = 117

var ts_symbol_names = [118]uintptr{
	0:   __ccgo_ts + 255,
	1:   __ccgo_ts + 259,
	2:   __ccgo_ts + 262,
	3:   __ccgo_ts + 277,
	4:   __ccgo_ts + 279,
	5:   __ccgo_ts + 287,
	6:   __ccgo_ts + 289,
	7:   __ccgo_ts + 292,
	8:   __ccgo_ts + 295,
	9:   __ccgo_ts + 297,
	10:  __ccgo_ts + 312,
	11:  __ccgo_ts + 328,
	12:  __ccgo_ts + 335,
	13:  __ccgo_ts + 337,
	14:  __ccgo_ts + 339,
	15:  __ccgo_ts + 344,
	16:  __ccgo_ts + 382,
	17:  __ccgo_ts + 420,
	18:  __ccgo_ts + 422,
	19:  __ccgo_ts + 425,
	20:  __ccgo_ts + 427,
	21:  __ccgo_ts + 429,
	22:  __ccgo_ts + 431,
	23:  __ccgo_ts + 439,
	24:  __ccgo_ts + 444,
	25:  __ccgo_ts + 446,
	26:  __ccgo_ts + 451,
	27:  __ccgo_ts + 454,
	28:  __ccgo_ts + 460,
	29:  __ccgo_ts + 465,
	30:  __ccgo_ts + 471,
	31:  __ccgo_ts + 475,
	32:  __ccgo_ts + 483,
	33:  __ccgo_ts + 496,
	34:  __ccgo_ts + 498,
	35:  __ccgo_ts + 500,
	36:  __ccgo_ts + 518,
	37:  __ccgo_ts + 520,
	38:  __ccgo_ts + 525,
	39:  __ccgo_ts + 531,
	40:  __ccgo_ts + 537,
	41:  __ccgo_ts + 544,
	42:  __ccgo_ts + 544,
	43:  __ccgo_ts + 544,
	44:  __ccgo_ts + 544,
	45:  __ccgo_ts + 553,
	46:  __ccgo_ts + 576,
	47:  __ccgo_ts + 594,
	48:  __ccgo_ts + 603,
	49:  __ccgo_ts + 611,
	50:  __ccgo_ts + 611,
	51:  __ccgo_ts + 611,
	52:  __ccgo_ts + 627,
	53:  __ccgo_ts + 279,
	54:  __ccgo_ts + 636,
	55:  __ccgo_ts + 642,
	56:  __ccgo_ts + 650,
	57:  __ccgo_ts + 665,
	58:  __ccgo_ts + 679,
	59:  __ccgo_ts + 679,
	60:  __ccgo_ts + 679,
	61:  __ccgo_ts + 689,
	62:  __ccgo_ts + 706,
	63:  __ccgo_ts + 714,
	64:  __ccgo_ts + 732,
	65:  __ccgo_ts + 742,
	66:  __ccgo_ts + 765,
	67:  __ccgo_ts + 796,
	68:  __ccgo_ts + 827,
	69:  __ccgo_ts + 840,
	70:  __ccgo_ts + 856,
	71:  __ccgo_ts + 865,
	72:  __ccgo_ts + 875,
	73:  __ccgo_ts + 889,
	74:  __ccgo_ts + 865,
	75:  __ccgo_ts + 903,
	76:  __ccgo_ts + 914,
	77:  __ccgo_ts + 925,
	78:  __ccgo_ts + 939,
	79:  __ccgo_ts + 946,
	80:  __ccgo_ts + 840,
	81:  __ccgo_ts + 961,
	82:  __ccgo_ts + 925,
	83:  __ccgo_ts + 972,
	84:  __ccgo_ts + 981,
	85:  __ccgo_ts + 840,
	86:  __ccgo_ts + 997,
	87:  __ccgo_ts + 865,
	88:  __ccgo_ts + 1009,
	89:  __ccgo_ts + 1020,
	90:  __ccgo_ts + 865,
	91:  __ccgo_ts + 1031,
	92:  __ccgo_ts + 1043,
	93:  __ccgo_ts + 925,
	94:  __ccgo_ts + 1055,
	95:  __ccgo_ts + 1065,
	96:  __ccgo_ts + 840,
	97:  __ccgo_ts + 1079,
	98:  __ccgo_ts + 925,
	99:  __ccgo_ts + 1089,
	100: __ccgo_ts + 1097,
	101: __ccgo_ts + 840,
	102: __ccgo_ts + 1115,
	103: __ccgo_ts + 925,
	104: __ccgo_ts + 1129,
	105: __ccgo_ts + 1141,
	106: __ccgo_ts + 1152,
	107: __ccgo_ts + 1163,
	108: __ccgo_ts + 1178,
	109: __ccgo_ts + 1163,
	110: __ccgo_ts + 1187,
	111: __ccgo_ts + 1163,
	112: __ccgo_ts + 1197,
	113: __ccgo_ts + 1163,
	114: __ccgo_ts + 1207,
	115: __ccgo_ts + 1218,
	116: __ccgo_ts + 1235,
	117: __ccgo_ts + 1253,
}

var ts_symbol_map = [118]TSSymbol{
	1:   uint16(anon_sym_LT_BANG),
	2:   uint16(aux_sym_doctype_token1),
	3:   uint16(anon_sym_GT),
	4:   uint16(sym__doctype),
	5:   uint16(anon_sym_LT),
	6:   uint16(anon_sym_SLASH_GT),
	7:   uint16(anon_sym_LT_SLASH),
	8:   uint16(anon_sym_EQ),
	9:   uint16(sym_attribute_name),
	10:  uint16(sym_attribute_value),
	11:  uint16(sym_entity),
	12:  uint16(anon_sym_SQUOTE),
	13:  uint16(anon_sym_DQUOTE),
	14:  uint16(sym_text),
	15:  uint16(aux_sym__single_quoted_attribute_value_token1),
	16:  uint16(aux_sym__double_quoted_attribute_value_token1),
	17:  uint16(anon_sym_POUND),
	18:  uint16(anon_sym_if),
	19:  uint16(anon_sym_LBRACE),
	20:  uint16(anon_sym_RBRACE),
	21:  uint16(anon_sym_COLON),
	22:  uint16(anon_sym_elseif),
	23:  uint16(anon_sym_else),
	24:  uint16(anon_sym_SLASH),
	25:  uint16(anon_sym_each),
	26:  uint16(anon_sym_as),
	27:  uint16(anon_sym_await),
	28:  uint16(anon_sym_then),
	29:  uint16(anon_sym_catch),
	30:  uint16(anon_sym_key),
	31:  uint16(anon_sym_snippet),
	32:  uint16(aux_sym_snippet_start_token1),
	33:  uint16(anon_sym_LPAREN),
	34:  uint16(anon_sym_RPAREN),
	35:  uint16(aux_sym__tag_value_token1),
	36:  uint16(anon_sym_AT),
	37:  uint16(anon_sym_html),
	38:  uint16(anon_sym_const),
	39:  uint16(anon_sym_debug),
	40:  uint16(anon_sym_render),
	41:  uint16(sym__start_tag_name),
	42:  uint16(sym__start_tag_name),
	43:  uint16(sym__start_tag_name),
	44:  uint16(sym__start_tag_name),
	45:  uint16(sym_erroneous_end_tag_name),
	46:  uint16(sym__implicit_end_tag),
	47:  uint16(sym_raw_text),
	48:  uint16(sym_comment),
	49:  uint16(sym_svelte_raw_text),
	50:  uint16(sym_svelte_raw_text),
	51:  uint16(sym_svelte_raw_text),
	52:  uint16(sym_document),
	53:  uint16(sym_doctype),
	54:  uint16(sym__node),
	55:  uint16(sym_element),
	56:  uint16(sym_script_element),
	57:  uint16(sym_style_element),
	58:  uint16(sym_start_tag),
	59:  uint16(sym_start_tag),
	60:  uint16(sym_start_tag),
	61:  uint16(sym_self_closing_tag),
	62:  uint16(sym_end_tag),
	63:  uint16(sym_erroneous_end_tag),
	64:  uint16(sym_attribute),
	65:  uint16(sym_quoted_attribute_value),
	66:  uint16(aux_sym__single_quoted_attribute_value),
	67:  uint16(aux_sym__double_quoted_attribute_value),
	68:  uint16(sym_if_statement),
	69:  uint16(sym__if_start_tag),
	70:  uint16(sym_if_start),
	71:  uint16(sym__else_if_tag),
	72:  uint16(sym_else_if_start),
	73:  uint16(sym_else_if_block),
	74:  uint16(sym__else_if_tag),
	75:  uint16(sym_else_start),
	76:  uint16(sym_else_block),
	77:  uint16(sym__if_end_tag),
	78:  uint16(sym_if_end),
	79:  uint16(sym_each_statement),
	80:  uint16(sym__if_start_tag),
	81:  uint16(sym_each_start),
	82:  uint16(sym__if_end_tag),
	83:  uint16(sym_each_end),
	84:  uint16(sym_await_statement),
	85:  uint16(sym__if_start_tag),
	86:  uint16(sym_await_start),
	87:  uint16(sym__else_if_tag),
	88:  uint16(sym_then_start),
	89:  uint16(sym_then_block),
	90:  uint16(sym__else_if_tag),
	91:  uint16(sym_catch_start),
	92:  uint16(sym_catch_block),
	93:  uint16(sym__if_end_tag),
	94:  uint16(sym_await_end),
	95:  uint16(sym_key_statement),
	96:  uint16(sym__if_start_tag),
	97:  uint16(sym_key_start),
	98:  uint16(sym__if_end_tag),
	99:  uint16(sym_key_end),
	100: uint16(sym_snippet_statement),
	101: uint16(sym__if_start_tag),
	102: uint16(sym_snippet_start),
	103: uint16(sym__if_end_tag),
	104: uint16(sym_snippet_end),
	105: uint16(sym_expression),
	106: uint16(sym__tag_value),
	107: uint16(sym__html_tag),
	108: uint16(sym_html_tag),
	109: uint16(sym__html_tag),
	110: uint16(sym_const_tag),
	111: uint16(sym__html_tag),
	112: uint16(sym_debug_tag),
	113: uint16(sym__html_tag),
	114: uint16(sym_render_tag),
	115: uint16(aux_sym_document_repeat1),
	116: uint16(aux_sym_start_tag_repeat1),
	117: uint16(aux_sym_if_statement_repeat1),
}

var ts_symbol_metadata = [118]TSSymbolMetadata{
	0: {
		Fnamed: libc.BoolUint8(1 != 0),
	},
	1: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	2: {},
	3: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	4: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	5: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	6: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	7: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	8: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	9: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	10: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	11: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	12: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	13: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	14: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	15: {},
	16: {},
	17: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	18: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	19: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	20: {
		Fvisible: libc.BoolUint8(1 != 0),
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
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	33: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	34: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	35: {},
	36: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	37: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	38: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	39: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	40: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	41: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	42: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
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
		Fnamed: libc.BoolUint8(1 != 0),
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
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	53: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	54: {
		Fnamed:     libc.BoolUint8(1 != 0),
		Fsupertype: libc.BoolUint8(1 != 0),
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
	66: {},
	67: {},
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
	73: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	74: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	75: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	76: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	77: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	78: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	79: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	80: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	81: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	82: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	83: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	84: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	85: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	86: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	87: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	88: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	89: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	90: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	91: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	92: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	93: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	94: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	95: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	96: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	97: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	98: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	99: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	100: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	101: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	102: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	103: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	104: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	105: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	106: {
		Fnamed: libc.BoolUint8(1 != 0),
	},
	107: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	108: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	109: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	110: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	111: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	112: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	113: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	114: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	115: {},
	116: {},
	117: {},
}

type ts_field_identifiers = int32

const field_condition = 1
const field_identifier = 2
const field_parameter = 3
const field_tag = 4

var ts_field_names = [5]uintptr{
	0: libc.UintptrFromInt32(0),
	1: __ccgo_ts + 1274,
	2: __ccgo_ts + 1284,
	3: __ccgo_ts + 1295,
	4: __ccgo_ts + 1305,
}

var ts_field_map_slices = [8]TSFieldMapSlice{
	2: {
		Flength: uint16(1),
	},
	3: {
		Findex:  uint16(1),
		Flength: uint16(2),
	},
	4: {
		Findex:  uint16(3),
		Flength: uint16(2),
	},
	5: {
		Findex:  uint16(5),
		Flength: uint16(1),
	},
	6: {
		Findex:  uint16(6),
		Flength: uint16(3),
	},
}

var ts_field_map_entries = [9]TSFieldMapEntry{
	0: {
		Ffield_id:    uint16(field_tag),
		Fchild_index: uint8(1),
	},
	1: {
		Ffield_id:    uint16(field_condition),
		Fchild_index: uint8(2),
	},
	2: {
		Ffield_id:    uint16(field_tag),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(1 != 0),
	},
	3: {
		Ffield_id:    uint16(field_identifier),
		Fchild_index: uint8(2),
	},
	4: {
		Ffield_id:    uint16(field_tag),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(1 != 0),
	},
	5: {
		Ffield_id:    uint16(field_tag),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(1 != 0),
	},
	6: {
		Ffield_id:    uint16(field_identifier),
		Fchild_index: uint8(2),
	},
	7: {
		Ffield_id:    uint16(field_parameter),
		Fchild_index: uint8(4),
	},
	8: {
		Ffield_id:    uint16(field_tag),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(1 != 0),
	},
}

var ts_alias_sequences = [8][7]TSSymbol{
	0: {},
	1: {
		0: uint16(sym_attribute_name),
	},
	7: {
		1: uint16(sym_attribute_value),
	},
}

var ts_non_terminal_alias_map = [13]uint16_t{
	0:  uint16(aux_sym__single_quoted_attribute_value),
	1:  uint16(2),
	2:  uint16(aux_sym__single_quoted_attribute_value),
	3:  uint16(sym_attribute_value),
	4:  uint16(aux_sym__double_quoted_attribute_value),
	5:  uint16(2),
	6:  uint16(aux_sym__double_quoted_attribute_value),
	7:  uint16(sym_attribute_value),
	8:  uint16(sym_expression),
	9:  uint16(2),
	10: uint16(sym_expression),
	11: uint16(sym_attribute_name),
}

var ts_primary_state_ids = [345]TSStateId{
	1:   uint16(1),
	2:   uint16(2),
	3:   uint16(3),
	4:   uint16(3),
	5:   uint16(2),
	6:   uint16(6),
	7:   uint16(7),
	8:   uint16(6),
	9:   uint16(7),
	10:  uint16(10),
	11:  uint16(11),
	12:  uint16(11),
	13:  uint16(10),
	14:  uint16(14),
	15:  uint16(15),
	16:  uint16(14),
	17:  uint16(15),
	18:  uint16(18),
	19:  uint16(19),
	20:  uint16(20),
	21:  uint16(18),
	22:  uint16(22),
	23:  uint16(23),
	24:  uint16(24),
	25:  uint16(25),
	26:  uint16(19),
	27:  uint16(23),
	28:  uint16(28),
	29:  uint16(29),
	30:  uint16(20),
	31:  uint16(31),
	32:  uint16(31),
	33:  uint16(29),
	34:  uint16(22),
	35:  uint16(24),
	36:  uint16(25),
	37:  uint16(37),
	38:  uint16(38),
	39:  uint16(39),
	40:  uint16(40),
	41:  uint16(41),
	42:  uint16(42),
	43:  uint16(43),
	44:  uint16(40),
	45:  uint16(45),
	46:  uint16(45),
	47:  uint16(47),
	48:  uint16(48),
	49:  uint16(48),
	50:  uint16(47),
	51:  uint16(51),
	52:  uint16(51),
	53:  uint16(53),
	54:  uint16(54),
	55:  uint16(55),
	56:  uint16(53),
	57:  uint16(55),
	58:  uint16(54),
	59:  uint16(59),
	60:  uint16(59),
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
	100: uint16(61),
	101: uint16(101),
	102: uint16(102),
	103: uint16(103),
	104: uint16(104),
	105: uint16(105),
	106: uint16(66),
	107: uint16(62),
	108: uint16(99),
	109: uint16(105),
	110: uint16(67),
	111: uint16(68),
	112: uint16(69),
	113: uint16(70),
	114: uint16(71),
	115: uint16(72),
	116: uint16(74),
	117: uint16(75),
	118: uint16(76),
	119: uint16(77),
	120: uint16(78),
	121: uint16(80),
	122: uint16(82),
	123: uint16(83),
	124: uint16(84),
	125: uint16(85),
	126: uint16(86),
	127: uint16(87),
	128: uint16(88),
	129: uint16(89),
	130: uint16(90),
	131: uint16(91),
	132: uint16(92),
	133: uint16(93),
	134: uint16(94),
	135: uint16(95),
	136: uint16(96),
	137: uint16(97),
	138: uint16(101),
	139: uint16(102),
	140: uint16(98),
	141: uint16(141),
	142: uint16(142),
	143: uint16(103),
	144: uint16(142),
	145: uint16(65),
	146: uint16(73),
	147: uint16(104),
	148: uint16(141),
	149: uint16(64),
	150: uint16(150),
	151: uint16(151),
	152: uint16(152),
	153: uint16(153),
	154: uint16(154),
	155: uint16(155),
	156: uint16(156),
	157: uint16(157),
	158: uint16(158),
	159: uint16(159),
	160: uint16(160),
	161: uint16(161),
	162: uint16(162),
	163: uint16(81),
	164: uint16(164),
	165: uint16(165),
	166: uint16(164),
	167: uint16(167),
	168: uint16(168),
	169: uint16(169),
	170: uint16(170),
	171: uint16(171),
	172: uint16(172),
	173: uint16(173),
	174: uint16(174),
	175: uint16(175),
	176: uint16(176),
	177: uint16(177),
	178: uint16(176),
	179: uint16(179),
	180: uint16(177),
	181: uint16(179),
	182: uint16(171),
	183: uint16(172),
	184: uint16(68),
	185: uint16(185),
	186: uint16(186),
	187: uint16(187),
	188: uint16(188),
	189: uint16(189),
	190: uint16(190),
	191: uint16(191),
	192: uint16(190),
	193: uint16(189),
	194: uint16(175),
	195: uint16(185),
	196: uint16(196),
	197: uint16(197),
	198: uint16(196),
	199: uint16(188),
	200: uint16(197),
	201: uint16(186),
	202: uint16(202),
	203: uint16(68),
	204: uint16(68),
	205: uint16(68),
	206: uint16(206),
	207: uint16(207),
	208: uint16(208),
	209: uint16(209),
	210: uint16(206),
	211: uint16(207),
	212: uint16(209),
	213: uint16(213),
	214: uint16(214),
	215: uint16(215),
	216: uint16(216),
	217: uint16(216),
	218: uint16(218),
	219: uint16(219),
	220: uint16(214),
	221: uint16(221),
	222: uint16(222),
	223: uint16(223),
	224: uint16(219),
	225: uint16(225),
	226: uint16(226),
	227: uint16(227),
	228: uint16(228),
	229: uint16(229),
	230: uint16(225),
	231: uint16(231),
	232: uint16(218),
	233: uint16(233),
	234: uint16(234),
	235: uint16(235),
	236: uint16(236),
	237: uint16(223),
	238: uint16(238),
	239: uint16(239),
	240: uint16(240),
	241: uint16(241),
	242: uint16(242),
	243: uint16(231),
	244: uint16(240),
	245: uint16(245),
	246: uint16(213),
	247: uint16(215),
	248: uint16(235),
	249: uint16(249),
	250: uint16(250),
	251: uint16(251),
	252: uint16(252),
	253: uint16(253),
	254: uint16(254),
	255: uint16(255),
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
	269: uint16(269),
	270: uint16(270),
	271: uint16(271),
	272: uint16(272),
	273: uint16(249),
	274: uint16(274),
	275: uint16(275),
	276: uint16(276),
	277: uint16(277),
	278: uint16(278),
	279: uint16(279),
	280: uint16(280),
	281: uint16(281),
	282: uint16(282),
	283: uint16(283),
	284: uint16(284),
	285: uint16(285),
	286: uint16(286),
	287: uint16(259),
	288: uint16(288),
	289: uint16(289),
	290: uint16(290),
	291: uint16(252),
	292: uint16(270),
	293: uint16(275),
	294: uint16(294),
	295: uint16(295),
	296: uint16(296),
	297: uint16(297),
	298: uint16(298),
	299: uint16(299),
	300: uint16(300),
	301: uint16(301),
	302: uint16(302),
	303: uint16(303),
	304: uint16(304),
	305: uint16(305),
	306: uint16(306),
	307: uint16(258),
	308: uint16(268),
	309: uint16(309),
	310: uint16(310),
	311: uint16(311),
	312: uint16(299),
	313: uint16(313),
	314: uint16(314),
	315: uint16(315),
	316: uint16(316),
	317: uint16(249),
	318: uint16(249),
	319: uint16(249),
	320: uint16(249),
	321: uint16(321),
	322: uint16(269),
	323: uint16(289),
	324: uint16(301),
	325: uint16(315),
	326: uint16(321),
	327: uint16(327),
	328: uint16(253),
	329: uint16(329),
	330: uint16(263),
	331: uint16(331),
	332: uint16(290),
	333: uint16(303),
	334: uint16(272),
	335: uint16(327),
	336: uint16(277),
	337: uint16(277),
	338: uint16(277),
	339: uint16(274),
	340: uint16(286),
	341: uint16(341),
	342: uint16(342),
	343: uint16(343),
	344: uint16(344),
}

var sym_attribute_name_character_set_1 = [9]TSCharacterRange{
	0: {
		Fend: int32(0x08),
	},
	1: {
		Fstart: int32(0x0e),
		Fend:   int32(0x1f),
	},
	2: {
		Fstart: int32('!'),
		Fend:   int32('!'),
	},
	3: {
		Fstart: int32('#'),
		Fend:   int32('&'),
	},
	4: {
		Fstart: int32('('),
		Fend:   int32('.'),
	},
	5: {
		Fstart: int32('0'),
		Fend:   int32(';'),
	},
	6: {
		Fstart: int32('?'),
		Fend:   int32('z'),
	},
	7: {
		Fstart: int32('|'),
		Fend:   int32('|'),
	},
	8: {
		Fstart: int32('~'),
		Fend:   int32(0x10ffff),
	},
}

func ts_lex(tls *libc.TLS, lexer uintptr, state TSStateId) (r uint8) {
	var eof, result, skip uint8
	var i uint32_t
	var lookahead int32_t
	_, _, _, _, _ = eof, i, lookahead, result, skip
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
			state = uint16(65)
			goto next_state
		}
		i = uint32(0)
		for {
			if !(uint64(i) < libc.Uint64FromInt64(100)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map1[i]) == lookahead {
				state = map1[i+uint32(1)]
				goto next_state
			}
			goto _1
		_1:
			;
			i = i + uint32(2)
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(144)
			goto next_state
		}
		return result
	case int32(1):
		if lookahead == int32('"') {
			state = uint16(118)
			goto next_state
		}
		if lookahead == int32('\'') {
			state = uint16(117)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(126)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(144)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('<') || int32('>') < lookahead) && lookahead != int32('}') {
			state = uint16(76)
			goto next_state
		}
		return result
	case int32(2):
		if lookahead == int32('"') {
			state = uint16(118)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(126)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(122)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(123)
			goto next_state
		}
		return result
	case int32(3):
		if lookahead == int32('#') {
			state = uint16(124)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(132)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(128)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(145)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(144)
			goto next_state
		}
		if lookahead == int32('$') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead == int32('#') {
			state = uint16(59)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(116)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead == int32('\'') {
			state = uint16(117)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(126)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(120)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead == int32('/') {
			state = uint16(7)
			goto next_state
		}
		if lookahead == int32('=') {
			state = uint16(74)
			goto next_state
		}
		if lookahead == int32('>') {
			state = uint16(69)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(126)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(144)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('"') && lookahead != int32('\'') && (lookahead < int32('<') || int32('>') < lookahead) && lookahead != int32('}') {
			state = uint16(75)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead == int32('>') {
			state = uint16(72)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead == int32('a') {
			state = uint16(12)
			goto next_state
		}
		if lookahead == int32('l') {
			state = uint16(44)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead == int32('a') {
			state = uint16(51)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(39)
			goto next_state
		}
		return result
	case int32(10):
		if lookahead == int32('a') {
			state = uint16(32)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead == int32('b') {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead == int32('c') {
			state = uint16(27)
			goto next_state
		}
		return result
	case int32(13):
		if lookahead == int32('c') {
			state = uint16(28)
			goto next_state
		}
		return result
	case int32(14):
		if lookahead == int32('d') {
			state = uint16(18)
			goto next_state
		}
		return result
	case int32(15):
		if lookahead == int32('e') {
			state = uint16(11)
			goto next_state
		}
		if lookahead == int32('O') || lookahead == int32('o') {
			state = uint16(54)
			goto next_state
		}
		return result
	case int32(16):
		if lookahead == int32('e') {
			state = uint16(53)
			goto next_state
		}
		return result
	case int32(17):
		if lookahead == int32('e') {
			state = uint16(131)
			goto next_state
		}
		return result
	case int32(18):
		if lookahead == int32('e') {
			state = uint16(42)
			goto next_state
		}
		return result
	case int32(19):
		if lookahead == int32('e') {
			state = uint16(130)
			goto next_state
		}
		return result
	case int32(20):
		if lookahead == int32('e') {
			state = uint16(34)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(144)
			goto next_state
		}
		return result
	case int32(21):
		if lookahead == int32('e') {
			state = uint16(37)
			goto next_state
		}
		return result
	case int32(22):
		if lookahead == int32('e') {
			state = uint16(38)
			goto next_state
		}
		return result
	case int32(23):
		if lookahead == int32('e') {
			state = uint16(50)
			goto next_state
		}
		return result
	case int32(24):
		if lookahead == int32('f') {
			state = uint16(125)
			goto next_state
		}
		return result
	case int32(25):
		if lookahead == int32('f') {
			state = uint16(129)
			goto next_state
		}
		return result
	case int32(26):
		if lookahead == int32('g') {
			state = uint16(148)
			goto next_state
		}
		return result
	case int32(27):
		if lookahead == int32('h') {
			state = uint16(134)
			goto next_state
		}
		return result
	case int32(28):
		if lookahead == int32('h') {
			state = uint16(138)
			goto next_state
		}
		return result
	case int32(29):
		if lookahead == int32('h') {
			state = uint16(22)
			goto next_state
		}
		return result
	case int32(30):
		if lookahead == int32('i') {
			state = uint16(40)
			goto next_state
		}
		return result
	case int32(31):
		if lookahead == int32('i') {
			state = uint16(25)
			goto next_state
		}
		return result
	case int32(32):
		if lookahead == int32('i') {
			state = uint16(48)
			goto next_state
		}
		return result
	case int32(33):
		if lookahead == int32('l') {
			state = uint16(146)
			goto next_state
		}
		return result
	case int32(34):
		if lookahead == int32('l') {
			state = uint16(46)
			goto next_state
		}
		return result
	case int32(35):
		if lookahead == int32('m') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(36):
		if lookahead == int32('n') {
			state = uint16(30)
			goto next_state
		}
		return result
	case int32(37):
		if lookahead == int32('n') {
			state = uint16(14)
			goto next_state
		}
		return result
	case int32(38):
		if lookahead == int32('n') {
			state = uint16(137)
			goto next_state
		}
		return result
	case int32(39):
		if lookahead == int32('n') {
			state = uint16(45)
			goto next_state
		}
		return result
	case int32(40):
		if lookahead == int32('p') {
			state = uint16(41)
			goto next_state
		}
		return result
	case int32(41):
		if lookahead == int32('p') {
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(42):
		if lookahead == int32('r') {
			state = uint16(149)
			goto next_state
		}
		return result
	case int32(43):
		if lookahead == int32('s') {
			state = uint16(135)
			goto next_state
		}
		if lookahead == int32('w') {
			state = uint16(10)
			goto next_state
		}
		return result
	case int32(44):
		if lookahead == int32('s') {
			state = uint16(17)
			goto next_state
		}
		return result
	case int32(45):
		if lookahead == int32('s') {
			state = uint16(49)
			goto next_state
		}
		return result
	case int32(46):
		if lookahead == int32('s') {
			state = uint16(19)
			goto next_state
		}
		return result
	case int32(47):
		if lookahead == int32('t') {
			state = uint16(35)
			goto next_state
		}
		return result
	case int32(48):
		if lookahead == int32('t') {
			state = uint16(136)
			goto next_state
		}
		return result
	case int32(49):
		if lookahead == int32('t') {
			state = uint16(147)
			goto next_state
		}
		return result
	case int32(50):
		if lookahead == int32('t') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(51):
		if lookahead == int32('t') {
			state = uint16(13)
			goto next_state
		}
		return result
	case int32(52):
		if lookahead == int32('u') {
			state = uint16(26)
			goto next_state
		}
		return result
	case int32(53):
		if lookahead == int32('y') {
			state = uint16(139)
			goto next_state
		}
		return result
	case int32(54):
		if lookahead == int32('C') || lookahead == int32('c') {
			state = uint16(58)
			goto next_state
		}
		return result
	case int32(55):
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(70)
			goto next_state
		}
		return result
	case int32(56):
		if lookahead == int32('O') || lookahead == int32('o') {
			state = uint16(54)
			goto next_state
		}
		return result
	case int32(57):
		if lookahead == int32('P') || lookahead == int32('p') {
			state = uint16(55)
			goto next_state
		}
		return result
	case int32(58):
		if lookahead == int32('T') || lookahead == int32('t') {
			state = uint16(60)
			goto next_state
		}
		return result
	case int32(59):
		if lookahead == int32('X') || lookahead == int32('x') {
			state = uint16(63)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(82)
			goto next_state
		}
		return result
	case int32(60):
		if lookahead == int32('Y') || lookahead == int32('y') {
			state = uint16(57)
			goto next_state
		}
		return result
	case int32(61):
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(61)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('&') && lookahead != int32('<') && lookahead != int32('>') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(62):
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(67)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('>') {
			state = uint16(68)
			goto next_state
		}
		return result
	case int32(63):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(87)
			goto next_state
		}
		return result
	case int32(64):
		if eof != 0 {
			state = uint16(65)
			goto next_state
		}
		if lookahead == int32('&') {
			state = uint16(4)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(71)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(126)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(144)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('>') && lookahead != int32('}') {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(65):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(0)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(66):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT_BANG)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(67):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_doctype_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(67)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('>') {
			state = uint16(68)
			goto next_state
		}
		return result
	case int32(68):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_doctype_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('>') {
			state = uint16(68)
			goto next_state
		}
		return result
	case int32(69):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_GT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(70):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__doctype)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(71):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('!') {
			state = uint16(66)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(73)
			goto next_state
		}
		return result
	case int32(72):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SLASH_GT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(73):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT_SLASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(74):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(75):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_attribute_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && set_contains(tls, uintptr(unsafe.Pointer(&sym_attribute_name_character_set_1)), uint32(9), lookahead) != 0 {
			state = uint16(75)
			goto next_state
		}
		return result
	case int32(76):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_attribute_value)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(eof != 0) && set_contains(tls, uintptr(unsafe.Pointer(&sym_attribute_name_character_set_1)), uint32(9), lookahead) != 0 || lookahead == int32('/') {
			state = uint16(76)
			goto next_state
		}
		return result
	case int32(77):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(78):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(77)
			goto next_state
		}
		return result
	case int32(79):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(77)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(78)
			goto next_state
		}
		return result
	case int32(80):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(77)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(79)
			goto next_state
		}
		return result
	case int32(81):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(77)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(80)
			goto next_state
		}
		return result
	case int32(82):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(77)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(81)
			goto next_state
		}
		return result
	case int32(83):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(77)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(78)
			goto next_state
		}
		return result
	case int32(84):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(77)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(83)
			goto next_state
		}
		return result
	case int32(85):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(77)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(84)
			goto next_state
		}
		return result
	case int32(86):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(77)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(85)
			goto next_state
		}
		return result
	case int32(87):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(77)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(86)
			goto next_state
		}
		return result
	case int32(88):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(77)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(78)
			goto next_state
		}
		return result
	case int32(89):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(77)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(88)
			goto next_state
		}
		return result
	case int32(90):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(77)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(89)
			goto next_state
		}
		return result
	case int32(91):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(77)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(90)
			goto next_state
		}
		return result
	case int32(92):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(77)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(91)
			goto next_state
		}
		return result
	case int32(93):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(77)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(92)
			goto next_state
		}
		return result
	case int32(94):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(77)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(93)
			goto next_state
		}
		return result
	case int32(95):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(77)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(94)
			goto next_state
		}
		return result
	case int32(96):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(77)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(95)
			goto next_state
		}
		return result
	case int32(97):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(77)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(96)
			goto next_state
		}
		return result
	case int32(98):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(77)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(97)
			goto next_state
		}
		return result
	case int32(99):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(77)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(98)
			goto next_state
		}
		return result
	case int32(100):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(77)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(99)
			goto next_state
		}
		return result
	case int32(101):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(77)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(100)
			goto next_state
		}
		return result
	case int32(102):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(77)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(101)
			goto next_state
		}
		return result
	case int32(103):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(77)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(102)
			goto next_state
		}
		return result
	case int32(104):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(77)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(103)
			goto next_state
		}
		return result
	case int32(105):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(77)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(106):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(77)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(105)
			goto next_state
		}
		return result
	case int32(107):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(77)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(106)
			goto next_state
		}
		return result
	case int32(108):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(77)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(107)
			goto next_state
		}
		return result
	case int32(109):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(77)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(108)
			goto next_state
		}
		return result
	case int32(110):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(77)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(109)
			goto next_state
		}
		return result
	case int32(111):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(77)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(110)
			goto next_state
		}
		return result
	case int32(112):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(77)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(111)
			goto next_state
		}
		return result
	case int32(113):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(77)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(112)
			goto next_state
		}
		return result
	case int32(114):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(77)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(113)
			goto next_state
		}
		return result
	case int32(115):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(77)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(114)
			goto next_state
		}
		return result
	case int32(116):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(77)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(115)
			goto next_state
		}
		return result
	case int32(117):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(118):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(119):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(61)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('&') && lookahead != int32('<') && lookahead != int32('>') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(120):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__single_quoted_attribute_value_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(120)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\'') && lookahead != int32('{') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(121):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__single_quoted_attribute_value_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\'') && lookahead != int32('{') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(122):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__double_quoted_attribute_value_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(122)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('"') && lookahead != int32('{') {
			state = uint16(123)
			goto next_state
		}
		return result
	case int32(123):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__double_quoted_attribute_value_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('"') && lookahead != int32('{') {
			state = uint16(123)
			goto next_state
		}
		return result
	case int32(124):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUND)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(125):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_if)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(126):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(127):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(128):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COLON)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(129):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_elseif)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(130):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_else)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(131):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_else)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(' ') {
			state = uint16(31)
			goto next_state
		}
		return result
	case int32(132):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SLASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(133):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SLASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('>') {
			state = uint16(72)
			goto next_state
		}
		return result
	case int32(134):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_each)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(135):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_as)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(136):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_await)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(137):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_then)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(138):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_catch)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(139):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_key)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(140):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_snippet)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(141):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_snippet_start_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(142):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(143):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(144):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__tag_value_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(144)
			goto next_state
		}
		return result
	case int32(145):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(146):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_html)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(147):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_const)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(148):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_debug)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(149):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_render)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	default:
		return libc.BoolUint8(0 != 0)
	}
	return r
}

var map1 = [50]uint16_t{
	0:  uint16('"'),
	1:  uint16(118),
	2:  uint16('#'),
	3:  uint16(124),
	4:  uint16('&'),
	5:  uint16(4),
	6:  uint16('\''),
	7:  uint16(117),
	8:  uint16('('),
	9:  uint16(142),
	10: uint16(')'),
	11: uint16(143),
	12: uint16('/'),
	13: uint16(133),
	14: uint16(':'),
	15: uint16(128),
	16: uint16('<'),
	17: uint16(71),
	18: uint16('='),
	19: uint16(74),
	20: uint16('>'),
	21: uint16(69),
	22: uint16('@'),
	23: uint16(145),
	24: uint16('D'),
	25: uint16(56),
	26: uint16('a'),
	27: uint16(43),
	28: uint16('c'),
	29: uint16(9),
	30: uint16('d'),
	31: uint16(15),
	32: uint16('e'),
	33: uint16(8),
	34: uint16('h'),
	35: uint16(47),
	36: uint16('i'),
	37: uint16(24),
	38: uint16('k'),
	39: uint16(16),
	40: uint16('r'),
	41: uint16(21),
	42: uint16('s'),
	43: uint16(36),
	44: uint16('t'),
	45: uint16(29),
	46: uint16('{'),
	47: uint16(126),
	48: uint16('}'),
	49: uint16(127),
}

var ts_lex_modes = [345]TSLexMode{
	0: {
		Fexternal_lex_state: uint16(1),
	},
	1: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	2: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	3: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	4: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	5: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	6: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	7: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	8: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	9: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	10: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	11: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	12: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	13: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	14: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(3),
	},
	15: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(3),
	},
	16: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(3),
	},
	17: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(3),
	},
	18: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	19: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	20: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	21: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	22: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	23: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	24: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	25: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	26: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	27: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(3),
	},
	28: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	29: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	30: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	31: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	32: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	33: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	34: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	35: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	36: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	37: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	38: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	39: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	40: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	41: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	42: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	43: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	44: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	45: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	46: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	47: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(4),
	},
	48: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(4),
	},
	49: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(4),
	},
	50: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(4),
	},
	51: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(4),
	},
	52: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(4),
	},
	53: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(5),
	},
	54: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(5),
	},
	55: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(5),
	},
	56: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(5),
	},
	57: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(5),
	},
	58: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(5),
	},
	59: {
		Fexternal_lex_state: uint16(6),
	},
	60: {
		Fexternal_lex_state: uint16(6),
	},
	61: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	62: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	63: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(3),
	},
	64: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	65: {
		Flex_state:          uint16(6),
		Fexternal_lex_state: uint16(7),
	},
	66: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	67: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	68: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	69: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	70: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	71: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	72: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	73: {
		Fexternal_lex_state: uint16(2),
	},
	74: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	75: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	76: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	77: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	78: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	79: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(3),
	},
	80: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	81: {
		Flex_state:          uint16(6),
		Fexternal_lex_state: uint16(7),
	},
	82: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	83: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	84: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	85: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	86: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	87: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	88: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	89: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	90: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	91: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	92: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	93: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	94: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	95: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	96: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	97: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	98: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(3),
	},
	99: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(3),
	},
	100: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(3),
	},
	101: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(3),
	},
	102: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(3),
	},
	103: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(3),
	},
	104: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(3),
	},
	105: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(3),
	},
	106: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(3),
	},
	107: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(3),
	},
	108: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	109: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	110: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(3),
	},
	111: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(3),
	},
	112: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(3),
	},
	113: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(3),
	},
	114: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(3),
	},
	115: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(3),
	},
	116: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(3),
	},
	117: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(3),
	},
	118: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(3),
	},
	119: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(3),
	},
	120: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(3),
	},
	121: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(3),
	},
	122: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(3),
	},
	123: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(3),
	},
	124: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(3),
	},
	125: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(3),
	},
	126: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(3),
	},
	127: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(3),
	},
	128: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(3),
	},
	129: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(3),
	},
	130: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(3),
	},
	131: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(3),
	},
	132: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(3),
	},
	133: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(3),
	},
	134: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(3),
	},
	135: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(3),
	},
	136: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(3),
	},
	137: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(3),
	},
	138: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	139: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	140: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	141: {
		Flex_state:          uint16(6),
		Fexternal_lex_state: uint16(7),
	},
	142: {
		Fexternal_lex_state: uint16(2),
	},
	143: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	144: {
		Fexternal_lex_state: uint16(2),
	},
	145: {
		Flex_state:          uint16(6),
		Fexternal_lex_state: uint16(7),
	},
	146: {
		Fexternal_lex_state: uint16(2),
	},
	147: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	148: {
		Flex_state:          uint16(6),
		Fexternal_lex_state: uint16(7),
	},
	149: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(3),
	},
	150: {
		Flex_state:          uint16(6),
		Fexternal_lex_state: uint16(2),
	},
	151: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	152: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	153: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	154: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	155: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	156: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	157: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	158: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	159: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	160: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	161: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	162: {
		Flex_state:          uint16(6),
		Fexternal_lex_state: uint16(2),
	},
	163: {
		Flex_state:          uint16(6),
		Fexternal_lex_state: uint16(2),
	},
	164: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(2),
	},
	165: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	166: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(2),
	},
	167: {
		Flex_state:          uint16(6),
		Fexternal_lex_state: uint16(2),
	},
	168: {
		Flex_state:          uint16(6),
		Fexternal_lex_state: uint16(2),
	},
	169: {
		Flex_state:          uint16(64),
		Fexternal_lex_state: uint16(2),
	},
	170: {
		Fexternal_lex_state: uint16(2),
	},
	171: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(2),
	},
	172: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(2),
	},
	173: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(2),
	},
	174: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(2),
	},
	175: {
		Flex_state:          uint16(6),
		Fexternal_lex_state: uint16(7),
	},
	176: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(8),
	},
	177: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(2),
	},
	178: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(8),
	},
	179: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(2),
	},
	180: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(2),
	},
	181: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(2),
	},
	182: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(2),
	},
	183: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(2),
	},
	184: {
		Flex_state:          uint16(6),
		Fexternal_lex_state: uint16(7),
	},
	185: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(8),
	},
	186: {
		Flex_state:          uint16(6),
		Fexternal_lex_state: uint16(7),
	},
	187: {
		Fexternal_lex_state: uint16(2),
	},
	188: {
		Flex_state:          uint16(6),
		Fexternal_lex_state: uint16(7),
	},
	189: {
		Fexternal_lex_state: uint16(2),
	},
	190: {
		Fexternal_lex_state: uint16(2),
	},
	191: {
		Fexternal_lex_state: uint16(2),
	},
	192: {
		Fexternal_lex_state: uint16(2),
	},
	193: {
		Fexternal_lex_state: uint16(2),
	},
	194: {
		Flex_state:          uint16(6),
		Fexternal_lex_state: uint16(2),
	},
	195: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(8),
	},
	196: {
		Flex_state:          uint16(6),
		Fexternal_lex_state: uint16(7),
	},
	197: {
		Flex_state:          uint16(6),
		Fexternal_lex_state: uint16(7),
	},
	198: {
		Flex_state:          uint16(6),
		Fexternal_lex_state: uint16(2),
	},
	199: {
		Flex_state:          uint16(6),
		Fexternal_lex_state: uint16(2),
	},
	200: {
		Flex_state:          uint16(6),
		Fexternal_lex_state: uint16(2),
	},
	201: {
		Flex_state:          uint16(6),
		Fexternal_lex_state: uint16(2),
	},
	202: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(2),
	},
	203: {
		Flex_state:          uint16(6),
		Fexternal_lex_state: uint16(2),
	},
	204: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(2),
	},
	205: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(2),
	},
	206: {
		Fexternal_lex_state: uint16(9),
	},
	207: {
		Fexternal_lex_state: uint16(10),
	},
	208: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(2),
	},
	209: {
		Fexternal_lex_state: uint16(9),
	},
	210: {
		Fexternal_lex_state: uint16(9),
	},
	211: {
		Fexternal_lex_state: uint16(10),
	},
	212: {
		Fexternal_lex_state: uint16(9),
	},
	213: {
		Fexternal_lex_state: uint16(2),
	},
	214: {
		Fexternal_lex_state: uint16(2),
	},
	215: {
		Fexternal_lex_state: uint16(2),
	},
	216: {
		Fexternal_lex_state: uint16(2),
	},
	217: {
		Fexternal_lex_state: uint16(2),
	},
	218: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(11),
	},
	219: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(11),
	},
	220: {
		Fexternal_lex_state: uint16(2),
	},
	221: {
		Fexternal_lex_state: uint16(12),
	},
	222: {
		Fexternal_lex_state: uint16(12),
	},
	223: {
		Fexternal_lex_state: uint16(2),
	},
	224: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(11),
	},
	225: {
		Fexternal_lex_state: uint16(2),
	},
	226: {
		Fexternal_lex_state: uint16(9),
	},
	227: {
		Fexternal_lex_state: uint16(13),
	},
	228: {
		Fexternal_lex_state: uint16(9),
	},
	229: {
		Fexternal_lex_state: uint16(9),
	},
	230: {
		Fexternal_lex_state: uint16(2),
	},
	231: {
		Fexternal_lex_state: uint16(2),
	},
	232: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(11),
	},
	233: {
		Fexternal_lex_state: uint16(2),
	},
	234: {
		Fexternal_lex_state: uint16(14),
	},
	235: {
		Fexternal_lex_state: uint16(2),
	},
	236: {
		Fexternal_lex_state: uint16(13),
	},
	237: {
		Fexternal_lex_state: uint16(2),
	},
	238: {
		Fexternal_lex_state: uint16(15),
	},
	239: {
		Fexternal_lex_state: uint16(13),
	},
	240: {
		Fexternal_lex_state: uint16(16),
	},
	241: {
		Fexternal_lex_state: uint16(2),
	},
	242: {
		Fexternal_lex_state: uint16(9),
	},
	243: {
		Fexternal_lex_state: uint16(2),
	},
	244: {
		Fexternal_lex_state: uint16(16),
	},
	245: {
		Fexternal_lex_state: uint16(13),
	},
	246: {
		Fexternal_lex_state: uint16(2),
	},
	247: {
		Fexternal_lex_state: uint16(2),
	},
	248: {
		Fexternal_lex_state: uint16(2),
	},
	249: {
		Fexternal_lex_state: uint16(2),
	},
	250: {
		Fexternal_lex_state: uint16(13),
	},
	251: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(2),
	},
	252: {
		Fexternal_lex_state: uint16(2),
	},
	253: {
		Fexternal_lex_state: uint16(2),
	},
	254: {
		Fexternal_lex_state: uint16(2),
	},
	255: {
		Fexternal_lex_state: uint16(2),
	},
	256: {
		Fexternal_lex_state: uint16(2),
	},
	257: {
		Fexternal_lex_state: uint16(2),
	},
	258: {
		Fexternal_lex_state: uint16(2),
	},
	259: {
		Fexternal_lex_state: uint16(2),
	},
	260: {
		Fexternal_lex_state: uint16(2),
	},
	261: {
		Fexternal_lex_state: uint16(2),
	},
	262: {
		Fexternal_lex_state: uint16(2),
	},
	263: {
		Fexternal_lex_state: uint16(17),
	},
	264: {
		Fexternal_lex_state: uint16(13),
	},
	265: {
		Fexternal_lex_state: uint16(2),
	},
	266: {
		Fexternal_lex_state: uint16(2),
	},
	267: {
		Fexternal_lex_state: uint16(2),
	},
	268: {
		Fexternal_lex_state: uint16(2),
	},
	269: {
		Fexternal_lex_state: uint16(18),
	},
	270: {
		Fexternal_lex_state: uint16(2),
	},
	271: {
		Fexternal_lex_state: uint16(13),
	},
	272: {
		Fexternal_lex_state: uint16(2),
	},
	273: {
		Fexternal_lex_state: uint16(2),
	},
	274: {
		Fexternal_lex_state: uint16(2),
	},
	275: {
		Fexternal_lex_state: uint16(2),
	},
	276: {
		Fexternal_lex_state: uint16(2),
	},
	277: {
		Fexternal_lex_state: uint16(13),
	},
	278: {
		Fexternal_lex_state: uint16(2),
	},
	279: {
		Fexternal_lex_state: uint16(2),
	},
	280: {
		Fexternal_lex_state: uint16(13),
	},
	281: {
		Fexternal_lex_state: uint16(2),
	},
	282: {
		Fexternal_lex_state: uint16(2),
	},
	283: {
		Fexternal_lex_state: uint16(2),
	},
	284: {
		Fexternal_lex_state: uint16(13),
	},
	285: {
		Fexternal_lex_state: uint16(2),
	},
	286: {
		Fexternal_lex_state: uint16(2),
	},
	287: {
		Fexternal_lex_state: uint16(2),
	},
	288: {
		Fexternal_lex_state: uint16(13),
	},
	289: {
		Fexternal_lex_state: uint16(2),
	},
	290: {
		Fexternal_lex_state: uint16(2),
	},
	291: {
		Fexternal_lex_state: uint16(2),
	},
	292: {
		Fexternal_lex_state: uint16(2),
	},
	293: {
		Fexternal_lex_state: uint16(2),
	},
	294: {
		Fexternal_lex_state: uint16(13),
	},
	295: {
		Fexternal_lex_state: uint16(13),
	},
	296: {
		Fexternal_lex_state: uint16(13),
	},
	297: {
		Fexternal_lex_state: uint16(2),
	},
	298: {
		Fexternal_lex_state: uint16(2),
	},
	299: {
		Fexternal_lex_state: uint16(2),
	},
	300: {
		Fexternal_lex_state: uint16(2),
	},
	301: {
		Flex_state:          uint16(62),
		Fexternal_lex_state: uint16(2),
	},
	302: {
		Fexternal_lex_state: uint16(13),
	},
	303: {
		Fexternal_lex_state: uint16(2),
	},
	304: {
		Fexternal_lex_state: uint16(2),
	},
	305: {
		Fexternal_lex_state: uint16(2),
	},
	306: {
		Fexternal_lex_state: uint16(2),
	},
	307: {
		Fexternal_lex_state: uint16(2),
	},
	308: {
		Fexternal_lex_state: uint16(2),
	},
	309: {
		Fexternal_lex_state: uint16(2),
	},
	310: {
		Flex_state:          uint16(20),
		Fexternal_lex_state: uint16(2),
	},
	311: {
		Fexternal_lex_state: uint16(2),
	},
	312: {
		Fexternal_lex_state: uint16(2),
	},
	313: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(2),
	},
	314: {
		Fexternal_lex_state: uint16(2),
	},
	315: {
		Fexternal_lex_state: uint16(2),
	},
	316: {
		Fexternal_lex_state: uint16(2),
	},
	317: {
		Fexternal_lex_state: uint16(2),
	},
	318: {
		Fexternal_lex_state: uint16(2),
	},
	319: {
		Fexternal_lex_state: uint16(2),
	},
	320: {
		Fexternal_lex_state: uint16(2),
	},
	321: {
		Fexternal_lex_state: uint16(2),
	},
	322: {
		Fexternal_lex_state: uint16(18),
	},
	323: {
		Fexternal_lex_state: uint16(2),
	},
	324: {
		Flex_state:          uint16(62),
		Fexternal_lex_state: uint16(2),
	},
	325: {
		Fexternal_lex_state: uint16(2),
	},
	326: {
		Fexternal_lex_state: uint16(2),
	},
	327: {
		Fexternal_lex_state: uint16(2),
	},
	328: {
		Fexternal_lex_state: uint16(2),
	},
	329: {
		Fexternal_lex_state: uint16(2),
	},
	330: {
		Fexternal_lex_state: uint16(17),
	},
	331: {
		Fexternal_lex_state: uint16(2),
	},
	332: {
		Fexternal_lex_state: uint16(2),
	},
	333: {
		Fexternal_lex_state: uint16(2),
	},
	334: {
		Fexternal_lex_state: uint16(2),
	},
	335: {
		Fexternal_lex_state: uint16(2),
	},
	336: {
		Fexternal_lex_state: uint16(13),
	},
	337: {
		Fexternal_lex_state: uint16(13),
	},
	338: {
		Fexternal_lex_state: uint16(13),
	},
	339: {
		Fexternal_lex_state: uint16(2),
	},
	340: {
		Fexternal_lex_state: uint16(2),
	},
	341: {
		Fexternal_lex_state: uint16(2),
	},
	342: {
		Fexternal_lex_state: uint16(2),
	},
	343: {
		Fexternal_lex_state: uint16(2),
	},
	344: {
		Fexternal_lex_state: uint16(2),
	},
}

var ts_parse_table = [2][118]uint16_t{
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
		13: uint16(1),
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
		31: uint16(1),
		33: uint16(1),
		34: uint16(1),
		35: uint16(3),
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
		48: uint16(3),
		49: uint16(1),
		50: uint16(1),
		51: uint16(1),
	},
	1: {
		0:   uint16(5),
		1:   uint16(7),
		5:   uint16(9),
		7:   uint16(11),
		11:  uint16(13),
		14:  uint16(13),
		19:  uint16(15),
		35:  uint16(3),
		48:  uint16(3),
		52:  uint16(281),
		53:  uint16(140),
		54:  uint16(28),
		55:  uint16(140),
		56:  uint16(140),
		57:  uint16(140),
		58:  uint16(15),
		59:  uint16(210),
		60:  uint16(209),
		61:  uint16(108),
		63:  uint16(140),
		68:  uint16(140),
		70:  uint16(3),
		79:  uint16(140),
		81:  uint16(10),
		84:  uint16(140),
		86:  uint16(6),
		95:  uint16(140),
		97:  uint16(33),
		100: uint16(140),
		102: uint16(20),
		105: uint16(140),
		108: uint16(140),
		110: uint16(140),
		112: uint16(140),
		114: uint16(140),
		115: uint16(28),
	},
}

var ts_small_parse_table = [7720]uint16_t{
	0:    uint16(22),
	1:    uint16(7),
	2:    uint16(1),
	3:    uint16(anon_sym_LT_BANG),
	4:    uint16(9),
	5:    uint16(1),
	6:    uint16(anon_sym_LT),
	7:    uint16(11),
	8:    uint16(1),
	9:    uint16(anon_sym_LT_SLASH),
	10:   uint16(17),
	11:   uint16(1),
	12:   uint16(anon_sym_LBRACE),
	13:   uint16(3),
	14:   uint16(1),
	15:   uint16(sym_if_start),
	16:   uint16(6),
	17:   uint16(1),
	18:   uint16(sym_await_start),
	19:   uint16(10),
	20:   uint16(1),
	21:   uint16(sym_each_start),
	22:   uint16(15),
	23:   uint16(1),
	24:   uint16(sym_start_tag),
	25:   uint16(20),
	26:   uint16(1),
	27:   uint16(sym_snippet_start),
	28:   uint16(33),
	29:   uint16(1),
	30:   uint16(sym_key_start),
	31:   uint16(37),
	32:   uint16(1),
	33:   uint16(sym_else_if_start),
	34:   uint16(40),
	35:   uint16(1),
	36:   uint16(sym_else_start),
	37:   uint16(72),
	38:   uint16(1),
	39:   uint16(sym_if_end),
	40:   uint16(108),
	41:   uint16(1),
	42:   uint16(sym_self_closing_tag),
	43:   uint16(209),
	44:   uint16(1),
	45:   uint16(sym_style_start_tag),
	46:   uint16(210),
	47:   uint16(1),
	48:   uint16(sym_script_start_tag),
	49:   uint16(230),
	50:   uint16(1),
	51:   uint16(sym_else_block),
	52:   uint16(3),
	53:   uint16(2),
	54:   uint16(sym_comment),
	55:   uint16(aux_sym__tag_value_token1),
	56:   uint16(13),
	57:   uint16(2),
	58:   uint16(sym_entity),
	59:   uint16(sym_text),
	60:   uint16(23),
	61:   uint16(2),
	62:   uint16(sym__node),
	63:   uint16(aux_sym_document_repeat1),
	64:   uint16(73),
	65:   uint16(2),
	66:   uint16(sym_else_if_block),
	67:   uint16(aux_sym_if_statement_repeat1),
	68:   uint16(140),
	69:   uint16(15),
	70:   uint16(sym_doctype),
	71:   uint16(sym_element),
	72:   uint16(sym_script_element),
	73:   uint16(sym_style_element),
	74:   uint16(sym_erroneous_end_tag),
	75:   uint16(sym_if_statement),
	76:   uint16(sym_each_statement),
	77:   uint16(sym_await_statement),
	78:   uint16(sym_key_statement),
	79:   uint16(sym_snippet_statement),
	80:   uint16(sym_expression),
	81:   uint16(sym_html_tag),
	82:   uint16(sym_const_tag),
	83:   uint16(sym_debug_tag),
	84:   uint16(sym_render_tag),
	85:   uint16(22),
	86:   uint16(7),
	87:   uint16(1),
	88:   uint16(anon_sym_LT_BANG),
	89:   uint16(9),
	90:   uint16(1),
	91:   uint16(anon_sym_LT),
	92:   uint16(11),
	93:   uint16(1),
	94:   uint16(anon_sym_LT_SLASH),
	95:   uint16(17),
	96:   uint16(1),
	97:   uint16(anon_sym_LBRACE),
	98:   uint16(3),
	99:   uint16(1),
	100:  uint16(sym_if_start),
	101:  uint16(6),
	102:  uint16(1),
	103:  uint16(sym_await_start),
	104:  uint16(10),
	105:  uint16(1),
	106:  uint16(sym_each_start),
	107:  uint16(15),
	108:  uint16(1),
	109:  uint16(sym_start_tag),
	110:  uint16(20),
	111:  uint16(1),
	112:  uint16(sym_snippet_start),
	113:  uint16(33),
	114:  uint16(1),
	115:  uint16(sym_key_start),
	116:  uint16(37),
	117:  uint16(1),
	118:  uint16(sym_else_if_start),
	119:  uint16(40),
	120:  uint16(1),
	121:  uint16(sym_else_start),
	122:  uint16(108),
	123:  uint16(1),
	124:  uint16(sym_self_closing_tag),
	125:  uint16(143),
	126:  uint16(1),
	127:  uint16(sym_if_end),
	128:  uint16(209),
	129:  uint16(1),
	130:  uint16(sym_style_start_tag),
	131:  uint16(210),
	132:  uint16(1),
	133:  uint16(sym_script_start_tag),
	134:  uint16(217),
	135:  uint16(1),
	136:  uint16(sym_else_block),
	137:  uint16(3),
	138:  uint16(2),
	139:  uint16(sym_comment),
	140:  uint16(aux_sym__tag_value_token1),
	141:  uint16(13),
	142:  uint16(2),
	143:  uint16(sym_entity),
	144:  uint16(sym_text),
	145:  uint16(2),
	146:  uint16(2),
	147:  uint16(sym__node),
	148:  uint16(aux_sym_document_repeat1),
	149:  uint16(144),
	150:  uint16(2),
	151:  uint16(sym_else_if_block),
	152:  uint16(aux_sym_if_statement_repeat1),
	153:  uint16(140),
	154:  uint16(15),
	155:  uint16(sym_doctype),
	156:  uint16(sym_element),
	157:  uint16(sym_script_element),
	158:  uint16(sym_style_element),
	159:  uint16(sym_erroneous_end_tag),
	160:  uint16(sym_if_statement),
	161:  uint16(sym_each_statement),
	162:  uint16(sym_await_statement),
	163:  uint16(sym_key_statement),
	164:  uint16(sym_snippet_statement),
	165:  uint16(sym_expression),
	166:  uint16(sym_html_tag),
	167:  uint16(sym_const_tag),
	168:  uint16(sym_debug_tag),
	169:  uint16(sym_render_tag),
	170:  uint16(22),
	171:  uint16(7),
	172:  uint16(1),
	173:  uint16(anon_sym_LT_BANG),
	174:  uint16(9),
	175:  uint16(1),
	176:  uint16(anon_sym_LT),
	177:  uint16(11),
	178:  uint16(1),
	179:  uint16(anon_sym_LT_SLASH),
	180:  uint16(19),
	181:  uint16(1),
	182:  uint16(anon_sym_LBRACE),
	183:  uint16(3),
	184:  uint16(1),
	185:  uint16(sym_if_start),
	186:  uint16(6),
	187:  uint16(1),
	188:  uint16(sym_await_start),
	189:  uint16(10),
	190:  uint16(1),
	191:  uint16(sym_each_start),
	192:  uint16(15),
	193:  uint16(1),
	194:  uint16(sym_start_tag),
	195:  uint16(20),
	196:  uint16(1),
	197:  uint16(sym_snippet_start),
	198:  uint16(33),
	199:  uint16(1),
	200:  uint16(sym_key_start),
	201:  uint16(37),
	202:  uint16(1),
	203:  uint16(sym_else_if_start),
	204:  uint16(40),
	205:  uint16(1),
	206:  uint16(sym_else_start),
	207:  uint16(103),
	208:  uint16(1),
	209:  uint16(sym_if_end),
	210:  uint16(108),
	211:  uint16(1),
	212:  uint16(sym_self_closing_tag),
	213:  uint16(209),
	214:  uint16(1),
	215:  uint16(sym_style_start_tag),
	216:  uint16(210),
	217:  uint16(1),
	218:  uint16(sym_script_start_tag),
	219:  uint16(216),
	220:  uint16(1),
	221:  uint16(sym_else_block),
	222:  uint16(3),
	223:  uint16(2),
	224:  uint16(sym_comment),
	225:  uint16(aux_sym__tag_value_token1),
	226:  uint16(13),
	227:  uint16(2),
	228:  uint16(sym_entity),
	229:  uint16(sym_text),
	230:  uint16(5),
	231:  uint16(2),
	232:  uint16(sym__node),
	233:  uint16(aux_sym_document_repeat1),
	234:  uint16(142),
	235:  uint16(2),
	236:  uint16(sym_else_if_block),
	237:  uint16(aux_sym_if_statement_repeat1),
	238:  uint16(140),
	239:  uint16(15),
	240:  uint16(sym_doctype),
	241:  uint16(sym_element),
	242:  uint16(sym_script_element),
	243:  uint16(sym_style_element),
	244:  uint16(sym_erroneous_end_tag),
	245:  uint16(sym_if_statement),
	246:  uint16(sym_each_statement),
	247:  uint16(sym_await_statement),
	248:  uint16(sym_key_statement),
	249:  uint16(sym_snippet_statement),
	250:  uint16(sym_expression),
	251:  uint16(sym_html_tag),
	252:  uint16(sym_const_tag),
	253:  uint16(sym_debug_tag),
	254:  uint16(sym_render_tag),
	255:  uint16(22),
	256:  uint16(7),
	257:  uint16(1),
	258:  uint16(anon_sym_LT_BANG),
	259:  uint16(9),
	260:  uint16(1),
	261:  uint16(anon_sym_LT),
	262:  uint16(11),
	263:  uint16(1),
	264:  uint16(anon_sym_LT_SLASH),
	265:  uint16(19),
	266:  uint16(1),
	267:  uint16(anon_sym_LBRACE),
	268:  uint16(3),
	269:  uint16(1),
	270:  uint16(sym_if_start),
	271:  uint16(6),
	272:  uint16(1),
	273:  uint16(sym_await_start),
	274:  uint16(10),
	275:  uint16(1),
	276:  uint16(sym_each_start),
	277:  uint16(15),
	278:  uint16(1),
	279:  uint16(sym_start_tag),
	280:  uint16(20),
	281:  uint16(1),
	282:  uint16(sym_snippet_start),
	283:  uint16(33),
	284:  uint16(1),
	285:  uint16(sym_key_start),
	286:  uint16(37),
	287:  uint16(1),
	288:  uint16(sym_else_if_start),
	289:  uint16(40),
	290:  uint16(1),
	291:  uint16(sym_else_start),
	292:  uint16(108),
	293:  uint16(1),
	294:  uint16(sym_self_closing_tag),
	295:  uint16(115),
	296:  uint16(1),
	297:  uint16(sym_if_end),
	298:  uint16(209),
	299:  uint16(1),
	300:  uint16(sym_style_start_tag),
	301:  uint16(210),
	302:  uint16(1),
	303:  uint16(sym_script_start_tag),
	304:  uint16(225),
	305:  uint16(1),
	306:  uint16(sym_else_block),
	307:  uint16(3),
	308:  uint16(2),
	309:  uint16(sym_comment),
	310:  uint16(aux_sym__tag_value_token1),
	311:  uint16(13),
	312:  uint16(2),
	313:  uint16(sym_entity),
	314:  uint16(sym_text),
	315:  uint16(23),
	316:  uint16(2),
	317:  uint16(sym__node),
	318:  uint16(aux_sym_document_repeat1),
	319:  uint16(146),
	320:  uint16(2),
	321:  uint16(sym_else_if_block),
	322:  uint16(aux_sym_if_statement_repeat1),
	323:  uint16(140),
	324:  uint16(15),
	325:  uint16(sym_doctype),
	326:  uint16(sym_element),
	327:  uint16(sym_script_element),
	328:  uint16(sym_style_element),
	329:  uint16(sym_erroneous_end_tag),
	330:  uint16(sym_if_statement),
	331:  uint16(sym_each_statement),
	332:  uint16(sym_await_statement),
	333:  uint16(sym_key_statement),
	334:  uint16(sym_snippet_statement),
	335:  uint16(sym_expression),
	336:  uint16(sym_html_tag),
	337:  uint16(sym_const_tag),
	338:  uint16(sym_debug_tag),
	339:  uint16(sym_render_tag),
	340:  uint16(22),
	341:  uint16(7),
	342:  uint16(1),
	343:  uint16(anon_sym_LT_BANG),
	344:  uint16(9),
	345:  uint16(1),
	346:  uint16(anon_sym_LT),
	347:  uint16(11),
	348:  uint16(1),
	349:  uint16(anon_sym_LT_SLASH),
	350:  uint16(21),
	351:  uint16(1),
	352:  uint16(anon_sym_LBRACE),
	353:  uint16(3),
	354:  uint16(1),
	355:  uint16(sym_if_start),
	356:  uint16(6),
	357:  uint16(1),
	358:  uint16(sym_await_start),
	359:  uint16(10),
	360:  uint16(1),
	361:  uint16(sym_each_start),
	362:  uint16(15),
	363:  uint16(1),
	364:  uint16(sym_start_tag),
	365:  uint16(20),
	366:  uint16(1),
	367:  uint16(sym_snippet_start),
	368:  uint16(33),
	369:  uint16(1),
	370:  uint16(sym_key_start),
	371:  uint16(38),
	372:  uint16(1),
	373:  uint16(sym_catch_start),
	374:  uint16(43),
	375:  uint16(1),
	376:  uint16(sym_then_start),
	377:  uint16(108),
	378:  uint16(1),
	379:  uint16(sym_self_closing_tag),
	380:  uint16(109),
	381:  uint16(1),
	382:  uint16(sym_await_end),
	383:  uint16(190),
	384:  uint16(1),
	385:  uint16(sym_then_block),
	386:  uint16(209),
	387:  uint16(1),
	388:  uint16(sym_style_start_tag),
	389:  uint16(210),
	390:  uint16(1),
	391:  uint16(sym_script_start_tag),
	392:  uint16(246),
	393:  uint16(1),
	394:  uint16(sym_catch_block),
	395:  uint16(3),
	396:  uint16(2),
	397:  uint16(sym_comment),
	398:  uint16(aux_sym__tag_value_token1),
	399:  uint16(13),
	400:  uint16(2),
	401:  uint16(sym_entity),
	402:  uint16(sym_text),
	403:  uint16(7),
	404:  uint16(2),
	405:  uint16(sym__node),
	406:  uint16(aux_sym_document_repeat1),
	407:  uint16(140),
	408:  uint16(15),
	409:  uint16(sym_doctype),
	410:  uint16(sym_element),
	411:  uint16(sym_script_element),
	412:  uint16(sym_style_element),
	413:  uint16(sym_erroneous_end_tag),
	414:  uint16(sym_if_statement),
	415:  uint16(sym_each_statement),
	416:  uint16(sym_await_statement),
	417:  uint16(sym_key_statement),
	418:  uint16(sym_snippet_statement),
	419:  uint16(sym_expression),
	420:  uint16(sym_html_tag),
	421:  uint16(sym_const_tag),
	422:  uint16(sym_debug_tag),
	423:  uint16(sym_render_tag),
	424:  uint16(22),
	425:  uint16(7),
	426:  uint16(1),
	427:  uint16(anon_sym_LT_BANG),
	428:  uint16(9),
	429:  uint16(1),
	430:  uint16(anon_sym_LT),
	431:  uint16(11),
	432:  uint16(1),
	433:  uint16(anon_sym_LT_SLASH),
	434:  uint16(21),
	435:  uint16(1),
	436:  uint16(anon_sym_LBRACE),
	437:  uint16(3),
	438:  uint16(1),
	439:  uint16(sym_if_start),
	440:  uint16(6),
	441:  uint16(1),
	442:  uint16(sym_await_start),
	443:  uint16(10),
	444:  uint16(1),
	445:  uint16(sym_each_start),
	446:  uint16(15),
	447:  uint16(1),
	448:  uint16(sym_start_tag),
	449:  uint16(20),
	450:  uint16(1),
	451:  uint16(sym_snippet_start),
	452:  uint16(33),
	453:  uint16(1),
	454:  uint16(sym_key_start),
	455:  uint16(38),
	456:  uint16(1),
	457:  uint16(sym_catch_start),
	458:  uint16(43),
	459:  uint16(1),
	460:  uint16(sym_then_start),
	461:  uint16(75),
	462:  uint16(1),
	463:  uint16(sym_await_end),
	464:  uint16(108),
	465:  uint16(1),
	466:  uint16(sym_self_closing_tag),
	467:  uint16(189),
	468:  uint16(1),
	469:  uint16(sym_then_block),
	470:  uint16(209),
	471:  uint16(1),
	472:  uint16(sym_style_start_tag),
	473:  uint16(210),
	474:  uint16(1),
	475:  uint16(sym_script_start_tag),
	476:  uint16(243),
	477:  uint16(1),
	478:  uint16(sym_catch_block),
	479:  uint16(3),
	480:  uint16(2),
	481:  uint16(sym_comment),
	482:  uint16(aux_sym__tag_value_token1),
	483:  uint16(13),
	484:  uint16(2),
	485:  uint16(sym_entity),
	486:  uint16(sym_text),
	487:  uint16(23),
	488:  uint16(2),
	489:  uint16(sym__node),
	490:  uint16(aux_sym_document_repeat1),
	491:  uint16(140),
	492:  uint16(15),
	493:  uint16(sym_doctype),
	494:  uint16(sym_element),
	495:  uint16(sym_script_element),
	496:  uint16(sym_style_element),
	497:  uint16(sym_erroneous_end_tag),
	498:  uint16(sym_if_statement),
	499:  uint16(sym_each_statement),
	500:  uint16(sym_await_statement),
	501:  uint16(sym_key_statement),
	502:  uint16(sym_snippet_statement),
	503:  uint16(sym_expression),
	504:  uint16(sym_html_tag),
	505:  uint16(sym_const_tag),
	506:  uint16(sym_debug_tag),
	507:  uint16(sym_render_tag),
	508:  uint16(22),
	509:  uint16(7),
	510:  uint16(1),
	511:  uint16(anon_sym_LT_BANG),
	512:  uint16(9),
	513:  uint16(1),
	514:  uint16(anon_sym_LT),
	515:  uint16(11),
	516:  uint16(1),
	517:  uint16(anon_sym_LT_SLASH),
	518:  uint16(23),
	519:  uint16(1),
	520:  uint16(anon_sym_LBRACE),
	521:  uint16(3),
	522:  uint16(1),
	523:  uint16(sym_if_start),
	524:  uint16(6),
	525:  uint16(1),
	526:  uint16(sym_await_start),
	527:  uint16(10),
	528:  uint16(1),
	529:  uint16(sym_each_start),
	530:  uint16(15),
	531:  uint16(1),
	532:  uint16(sym_start_tag),
	533:  uint16(20),
	534:  uint16(1),
	535:  uint16(sym_snippet_start),
	536:  uint16(33),
	537:  uint16(1),
	538:  uint16(sym_key_start),
	539:  uint16(38),
	540:  uint16(1),
	541:  uint16(sym_catch_start),
	542:  uint16(43),
	543:  uint16(1),
	544:  uint16(sym_then_start),
	545:  uint16(105),
	546:  uint16(1),
	547:  uint16(sym_await_end),
	548:  uint16(108),
	549:  uint16(1),
	550:  uint16(sym_self_closing_tag),
	551:  uint16(192),
	552:  uint16(1),
	553:  uint16(sym_then_block),
	554:  uint16(209),
	555:  uint16(1),
	556:  uint16(sym_style_start_tag),
	557:  uint16(210),
	558:  uint16(1),
	559:  uint16(sym_script_start_tag),
	560:  uint16(213),
	561:  uint16(1),
	562:  uint16(sym_catch_block),
	563:  uint16(3),
	564:  uint16(2),
	565:  uint16(sym_comment),
	566:  uint16(aux_sym__tag_value_token1),
	567:  uint16(13),
	568:  uint16(2),
	569:  uint16(sym_entity),
	570:  uint16(sym_text),
	571:  uint16(9),
	572:  uint16(2),
	573:  uint16(sym__node),
	574:  uint16(aux_sym_document_repeat1),
	575:  uint16(140),
	576:  uint16(15),
	577:  uint16(sym_doctype),
	578:  uint16(sym_element),
	579:  uint16(sym_script_element),
	580:  uint16(sym_style_element),
	581:  uint16(sym_erroneous_end_tag),
	582:  uint16(sym_if_statement),
	583:  uint16(sym_each_statement),
	584:  uint16(sym_await_statement),
	585:  uint16(sym_key_statement),
	586:  uint16(sym_snippet_statement),
	587:  uint16(sym_expression),
	588:  uint16(sym_html_tag),
	589:  uint16(sym_const_tag),
	590:  uint16(sym_debug_tag),
	591:  uint16(sym_render_tag),
	592:  uint16(22),
	593:  uint16(7),
	594:  uint16(1),
	595:  uint16(anon_sym_LT_BANG),
	596:  uint16(9),
	597:  uint16(1),
	598:  uint16(anon_sym_LT),
	599:  uint16(11),
	600:  uint16(1),
	601:  uint16(anon_sym_LT_SLASH),
	602:  uint16(23),
	603:  uint16(1),
	604:  uint16(anon_sym_LBRACE),
	605:  uint16(3),
	606:  uint16(1),
	607:  uint16(sym_if_start),
	608:  uint16(6),
	609:  uint16(1),
	610:  uint16(sym_await_start),
	611:  uint16(10),
	612:  uint16(1),
	613:  uint16(sym_each_start),
	614:  uint16(15),
	615:  uint16(1),
	616:  uint16(sym_start_tag),
	617:  uint16(20),
	618:  uint16(1),
	619:  uint16(sym_snippet_start),
	620:  uint16(33),
	621:  uint16(1),
	622:  uint16(sym_key_start),
	623:  uint16(38),
	624:  uint16(1),
	625:  uint16(sym_catch_start),
	626:  uint16(43),
	627:  uint16(1),
	628:  uint16(sym_then_start),
	629:  uint16(108),
	630:  uint16(1),
	631:  uint16(sym_self_closing_tag),
	632:  uint16(117),
	633:  uint16(1),
	634:  uint16(sym_await_end),
	635:  uint16(193),
	636:  uint16(1),
	637:  uint16(sym_then_block),
	638:  uint16(209),
	639:  uint16(1),
	640:  uint16(sym_style_start_tag),
	641:  uint16(210),
	642:  uint16(1),
	643:  uint16(sym_script_start_tag),
	644:  uint16(231),
	645:  uint16(1),
	646:  uint16(sym_catch_block),
	647:  uint16(3),
	648:  uint16(2),
	649:  uint16(sym_comment),
	650:  uint16(aux_sym__tag_value_token1),
	651:  uint16(13),
	652:  uint16(2),
	653:  uint16(sym_entity),
	654:  uint16(sym_text),
	655:  uint16(23),
	656:  uint16(2),
	657:  uint16(sym__node),
	658:  uint16(aux_sym_document_repeat1),
	659:  uint16(140),
	660:  uint16(15),
	661:  uint16(sym_doctype),
	662:  uint16(sym_element),
	663:  uint16(sym_script_element),
	664:  uint16(sym_style_element),
	665:  uint16(sym_erroneous_end_tag),
	666:  uint16(sym_if_statement),
	667:  uint16(sym_each_statement),
	668:  uint16(sym_await_statement),
	669:  uint16(sym_key_statement),
	670:  uint16(sym_snippet_statement),
	671:  uint16(sym_expression),
	672:  uint16(sym_html_tag),
	673:  uint16(sym_const_tag),
	674:  uint16(sym_debug_tag),
	675:  uint16(sym_render_tag),
	676:  uint16(20),
	677:  uint16(7),
	678:  uint16(1),
	679:  uint16(anon_sym_LT_BANG),
	680:  uint16(9),
	681:  uint16(1),
	682:  uint16(anon_sym_LT),
	683:  uint16(11),
	684:  uint16(1),
	685:  uint16(anon_sym_LT_SLASH),
	686:  uint16(25),
	687:  uint16(1),
	688:  uint16(anon_sym_LBRACE),
	689:  uint16(3),
	690:  uint16(1),
	691:  uint16(sym_if_start),
	692:  uint16(6),
	693:  uint16(1),
	694:  uint16(sym_await_start),
	695:  uint16(10),
	696:  uint16(1),
	697:  uint16(sym_each_start),
	698:  uint16(15),
	699:  uint16(1),
	700:  uint16(sym_start_tag),
	701:  uint16(20),
	702:  uint16(1),
	703:  uint16(sym_snippet_start),
	704:  uint16(31),
	705:  uint16(1),
	706:  uint16(sym_else_block),
	707:  uint16(33),
	708:  uint16(1),
	709:  uint16(sym_key_start),
	710:  uint16(44),
	711:  uint16(1),
	712:  uint16(sym_else_start),
	713:  uint16(108),
	714:  uint16(1),
	715:  uint16(sym_self_closing_tag),
	716:  uint16(147),
	717:  uint16(1),
	718:  uint16(sym_each_end),
	719:  uint16(209),
	720:  uint16(1),
	721:  uint16(sym_style_start_tag),
	722:  uint16(210),
	723:  uint16(1),
	724:  uint16(sym_script_start_tag),
	725:  uint16(3),
	726:  uint16(2),
	727:  uint16(sym_comment),
	728:  uint16(aux_sym__tag_value_token1),
	729:  uint16(13),
	730:  uint16(2),
	731:  uint16(sym_entity),
	732:  uint16(sym_text),
	733:  uint16(11),
	734:  uint16(2),
	735:  uint16(sym__node),
	736:  uint16(aux_sym_document_repeat1),
	737:  uint16(140),
	738:  uint16(15),
	739:  uint16(sym_doctype),
	740:  uint16(sym_element),
	741:  uint16(sym_script_element),
	742:  uint16(sym_style_element),
	743:  uint16(sym_erroneous_end_tag),
	744:  uint16(sym_if_statement),
	745:  uint16(sym_each_statement),
	746:  uint16(sym_await_statement),
	747:  uint16(sym_key_statement),
	748:  uint16(sym_snippet_statement),
	749:  uint16(sym_expression),
	750:  uint16(sym_html_tag),
	751:  uint16(sym_const_tag),
	752:  uint16(sym_debug_tag),
	753:  uint16(sym_render_tag),
	754:  uint16(20),
	755:  uint16(7),
	756:  uint16(1),
	757:  uint16(anon_sym_LT_BANG),
	758:  uint16(9),
	759:  uint16(1),
	760:  uint16(anon_sym_LT),
	761:  uint16(11),
	762:  uint16(1),
	763:  uint16(anon_sym_LT_SLASH),
	764:  uint16(25),
	765:  uint16(1),
	766:  uint16(anon_sym_LBRACE),
	767:  uint16(3),
	768:  uint16(1),
	769:  uint16(sym_if_start),
	770:  uint16(6),
	771:  uint16(1),
	772:  uint16(sym_await_start),
	773:  uint16(10),
	774:  uint16(1),
	775:  uint16(sym_each_start),
	776:  uint16(15),
	777:  uint16(1),
	778:  uint16(sym_start_tag),
	779:  uint16(20),
	780:  uint16(1),
	781:  uint16(sym_snippet_start),
	782:  uint16(25),
	783:  uint16(1),
	784:  uint16(sym_else_block),
	785:  uint16(33),
	786:  uint16(1),
	787:  uint16(sym_key_start),
	788:  uint16(44),
	789:  uint16(1),
	790:  uint16(sym_else_start),
	791:  uint16(74),
	792:  uint16(1),
	793:  uint16(sym_each_end),
	794:  uint16(108),
	795:  uint16(1),
	796:  uint16(sym_self_closing_tag),
	797:  uint16(209),
	798:  uint16(1),
	799:  uint16(sym_style_start_tag),
	800:  uint16(210),
	801:  uint16(1),
	802:  uint16(sym_script_start_tag),
	803:  uint16(3),
	804:  uint16(2),
	805:  uint16(sym_comment),
	806:  uint16(aux_sym__tag_value_token1),
	807:  uint16(13),
	808:  uint16(2),
	809:  uint16(sym_entity),
	810:  uint16(sym_text),
	811:  uint16(23),
	812:  uint16(2),
	813:  uint16(sym__node),
	814:  uint16(aux_sym_document_repeat1),
	815:  uint16(140),
	816:  uint16(15),
	817:  uint16(sym_doctype),
	818:  uint16(sym_element),
	819:  uint16(sym_script_element),
	820:  uint16(sym_style_element),
	821:  uint16(sym_erroneous_end_tag),
	822:  uint16(sym_if_statement),
	823:  uint16(sym_each_statement),
	824:  uint16(sym_await_statement),
	825:  uint16(sym_key_statement),
	826:  uint16(sym_snippet_statement),
	827:  uint16(sym_expression),
	828:  uint16(sym_html_tag),
	829:  uint16(sym_const_tag),
	830:  uint16(sym_debug_tag),
	831:  uint16(sym_render_tag),
	832:  uint16(20),
	833:  uint16(7),
	834:  uint16(1),
	835:  uint16(anon_sym_LT_BANG),
	836:  uint16(9),
	837:  uint16(1),
	838:  uint16(anon_sym_LT),
	839:  uint16(11),
	840:  uint16(1),
	841:  uint16(anon_sym_LT_SLASH),
	842:  uint16(27),
	843:  uint16(1),
	844:  uint16(anon_sym_LBRACE),
	845:  uint16(3),
	846:  uint16(1),
	847:  uint16(sym_if_start),
	848:  uint16(6),
	849:  uint16(1),
	850:  uint16(sym_await_start),
	851:  uint16(10),
	852:  uint16(1),
	853:  uint16(sym_each_start),
	854:  uint16(15),
	855:  uint16(1),
	856:  uint16(sym_start_tag),
	857:  uint16(20),
	858:  uint16(1),
	859:  uint16(sym_snippet_start),
	860:  uint16(33),
	861:  uint16(1),
	862:  uint16(sym_key_start),
	863:  uint16(36),
	864:  uint16(1),
	865:  uint16(sym_else_block),
	866:  uint16(44),
	867:  uint16(1),
	868:  uint16(sym_else_start),
	869:  uint16(108),
	870:  uint16(1),
	871:  uint16(sym_self_closing_tag),
	872:  uint16(116),
	873:  uint16(1),
	874:  uint16(sym_each_end),
	875:  uint16(209),
	876:  uint16(1),
	877:  uint16(sym_style_start_tag),
	878:  uint16(210),
	879:  uint16(1),
	880:  uint16(sym_script_start_tag),
	881:  uint16(3),
	882:  uint16(2),
	883:  uint16(sym_comment),
	884:  uint16(aux_sym__tag_value_token1),
	885:  uint16(13),
	886:  uint16(2),
	887:  uint16(sym_entity),
	888:  uint16(sym_text),
	889:  uint16(23),
	890:  uint16(2),
	891:  uint16(sym__node),
	892:  uint16(aux_sym_document_repeat1),
	893:  uint16(140),
	894:  uint16(15),
	895:  uint16(sym_doctype),
	896:  uint16(sym_element),
	897:  uint16(sym_script_element),
	898:  uint16(sym_style_element),
	899:  uint16(sym_erroneous_end_tag),
	900:  uint16(sym_if_statement),
	901:  uint16(sym_each_statement),
	902:  uint16(sym_await_statement),
	903:  uint16(sym_key_statement),
	904:  uint16(sym_snippet_statement),
	905:  uint16(sym_expression),
	906:  uint16(sym_html_tag),
	907:  uint16(sym_const_tag),
	908:  uint16(sym_debug_tag),
	909:  uint16(sym_render_tag),
	910:  uint16(20),
	911:  uint16(7),
	912:  uint16(1),
	913:  uint16(anon_sym_LT_BANG),
	914:  uint16(9),
	915:  uint16(1),
	916:  uint16(anon_sym_LT),
	917:  uint16(11),
	918:  uint16(1),
	919:  uint16(anon_sym_LT_SLASH),
	920:  uint16(27),
	921:  uint16(1),
	922:  uint16(anon_sym_LBRACE),
	923:  uint16(3),
	924:  uint16(1),
	925:  uint16(sym_if_start),
	926:  uint16(6),
	927:  uint16(1),
	928:  uint16(sym_await_start),
	929:  uint16(10),
	930:  uint16(1),
	931:  uint16(sym_each_start),
	932:  uint16(15),
	933:  uint16(1),
	934:  uint16(sym_start_tag),
	935:  uint16(20),
	936:  uint16(1),
	937:  uint16(sym_snippet_start),
	938:  uint16(32),
	939:  uint16(1),
	940:  uint16(sym_else_block),
	941:  uint16(33),
	942:  uint16(1),
	943:  uint16(sym_key_start),
	944:  uint16(44),
	945:  uint16(1),
	946:  uint16(sym_else_start),
	947:  uint16(104),
	948:  uint16(1),
	949:  uint16(sym_each_end),
	950:  uint16(108),
	951:  uint16(1),
	952:  uint16(sym_self_closing_tag),
	953:  uint16(209),
	954:  uint16(1),
	955:  uint16(sym_style_start_tag),
	956:  uint16(210),
	957:  uint16(1),
	958:  uint16(sym_script_start_tag),
	959:  uint16(3),
	960:  uint16(2),
	961:  uint16(sym_comment),
	962:  uint16(aux_sym__tag_value_token1),
	963:  uint16(13),
	964:  uint16(2),
	965:  uint16(sym_entity),
	966:  uint16(sym_text),
	967:  uint16(12),
	968:  uint16(2),
	969:  uint16(sym__node),
	970:  uint16(aux_sym_document_repeat1),
	971:  uint16(140),
	972:  uint16(15),
	973:  uint16(sym_doctype),
	974:  uint16(sym_element),
	975:  uint16(sym_script_element),
	976:  uint16(sym_style_element),
	977:  uint16(sym_erroneous_end_tag),
	978:  uint16(sym_if_statement),
	979:  uint16(sym_each_statement),
	980:  uint16(sym_await_statement),
	981:  uint16(sym_key_statement),
	982:  uint16(sym_snippet_statement),
	983:  uint16(sym_expression),
	984:  uint16(sym_html_tag),
	985:  uint16(sym_const_tag),
	986:  uint16(sym_debug_tag),
	987:  uint16(sym_render_tag),
	988:  uint16(19),
	989:  uint16(29),
	990:  uint16(1),
	991:  uint16(anon_sym_LT_BANG),
	992:  uint16(31),
	993:  uint16(1),
	994:  uint16(anon_sym_LT),
	995:  uint16(33),
	996:  uint16(1),
	997:  uint16(anon_sym_LT_SLASH),
	998:  uint16(37),
	999:  uint16(1),
	1000: uint16(anon_sym_LBRACE),
	1001: uint16(39),
	1002: uint16(1),
	1003: uint16(sym__implicit_end_tag),
	1004: uint16(4),
	1005: uint16(1),
	1006: uint16(sym_if_start),
	1007: uint16(8),
	1008: uint16(1),
	1009: uint16(sym_await_start),
	1010: uint16(13),
	1011: uint16(1),
	1012: uint16(sym_each_start),
	1013: uint16(17),
	1014: uint16(1),
	1015: uint16(sym_start_tag),
	1016: uint16(29),
	1017: uint16(1),
	1018: uint16(sym_key_start),
	1019: uint16(30),
	1020: uint16(1),
	1021: uint16(sym_snippet_start),
	1022: uint16(69),
	1023: uint16(1),
	1024: uint16(sym_end_tag),
	1025: uint16(99),
	1026: uint16(1),
	1027: uint16(sym_self_closing_tag),
	1028: uint16(206),
	1029: uint16(1),
	1030: uint16(sym_script_start_tag),
	1031: uint16(212),
	1032: uint16(1),
	1033: uint16(sym_style_start_tag),
	1034: uint16(3),
	1035: uint16(2),
	1036: uint16(sym_comment),
	1037: uint16(aux_sym__tag_value_token1),
	1038: uint16(35),
	1039: uint16(2),
	1040: uint16(sym_entity),
	1041: uint16(sym_text),
	1042: uint16(27),
	1043: uint16(2),
	1044: uint16(sym__node),
	1045: uint16(aux_sym_document_repeat1),
	1046: uint16(98),
	1047: uint16(15),
	1048: uint16(sym_doctype),
	1049: uint16(sym_element),
	1050: uint16(sym_script_element),
	1051: uint16(sym_style_element),
	1052: uint16(sym_erroneous_end_tag),
	1053: uint16(sym_if_statement),
	1054: uint16(sym_each_statement),
	1055: uint16(sym_await_statement),
	1056: uint16(sym_key_statement),
	1057: uint16(sym_snippet_statement),
	1058: uint16(sym_expression),
	1059: uint16(sym_html_tag),
	1060: uint16(sym_const_tag),
	1061: uint16(sym_debug_tag),
	1062: uint16(sym_render_tag),
	1063: uint16(19),
	1064: uint16(29),
	1065: uint16(1),
	1066: uint16(anon_sym_LT_BANG),
	1067: uint16(31),
	1068: uint16(1),
	1069: uint16(anon_sym_LT),
	1070: uint16(33),
	1071: uint16(1),
	1072: uint16(anon_sym_LT_SLASH),
	1073: uint16(37),
	1074: uint16(1),
	1075: uint16(anon_sym_LBRACE),
	1076: uint16(41),
	1077: uint16(1),
	1078: uint16(sym__implicit_end_tag),
	1079: uint16(4),
	1080: uint16(1),
	1081: uint16(sym_if_start),
	1082: uint16(8),
	1083: uint16(1),
	1084: uint16(sym_await_start),
	1085: uint16(13),
	1086: uint16(1),
	1087: uint16(sym_each_start),
	1088: uint16(17),
	1089: uint16(1),
	1090: uint16(sym_start_tag),
	1091: uint16(29),
	1092: uint16(1),
	1093: uint16(sym_key_start),
	1094: uint16(30),
	1095: uint16(1),
	1096: uint16(sym_snippet_start),
	1097: uint16(61),
	1098: uint16(1),
	1099: uint16(sym_end_tag),
	1100: uint16(99),
	1101: uint16(1),
	1102: uint16(sym_self_closing_tag),
	1103: uint16(206),
	1104: uint16(1),
	1105: uint16(sym_script_start_tag),
	1106: uint16(212),
	1107: uint16(1),
	1108: uint16(sym_style_start_tag),
	1109: uint16(3),
	1110: uint16(2),
	1111: uint16(sym_comment),
	1112: uint16(aux_sym__tag_value_token1),
	1113: uint16(35),
	1114: uint16(2),
	1115: uint16(sym_entity),
	1116: uint16(sym_text),
	1117: uint16(14),
	1118: uint16(2),
	1119: uint16(sym__node),
	1120: uint16(aux_sym_document_repeat1),
	1121: uint16(98),
	1122: uint16(15),
	1123: uint16(sym_doctype),
	1124: uint16(sym_element),
	1125: uint16(sym_script_element),
	1126: uint16(sym_style_element),
	1127: uint16(sym_erroneous_end_tag),
	1128: uint16(sym_if_statement),
	1129: uint16(sym_each_statement),
	1130: uint16(sym_await_statement),
	1131: uint16(sym_key_statement),
	1132: uint16(sym_snippet_statement),
	1133: uint16(sym_expression),
	1134: uint16(sym_html_tag),
	1135: uint16(sym_const_tag),
	1136: uint16(sym_debug_tag),
	1137: uint16(sym_render_tag),
	1138: uint16(19),
	1139: uint16(29),
	1140: uint16(1),
	1141: uint16(anon_sym_LT_BANG),
	1142: uint16(31),
	1143: uint16(1),
	1144: uint16(anon_sym_LT),
	1145: uint16(37),
	1146: uint16(1),
	1147: uint16(anon_sym_LBRACE),
	1148: uint16(43),
	1149: uint16(1),
	1150: uint16(anon_sym_LT_SLASH),
	1151: uint16(45),
	1152: uint16(1),
	1153: uint16(sym__implicit_end_tag),
	1154: uint16(4),
	1155: uint16(1),
	1156: uint16(sym_if_start),
	1157: uint16(8),
	1158: uint16(1),
	1159: uint16(sym_await_start),
	1160: uint16(13),
	1161: uint16(1),
	1162: uint16(sym_each_start),
	1163: uint16(17),
	1164: uint16(1),
	1165: uint16(sym_start_tag),
	1166: uint16(29),
	1167: uint16(1),
	1168: uint16(sym_key_start),
	1169: uint16(30),
	1170: uint16(1),
	1171: uint16(sym_snippet_start),
	1172: uint16(99),
	1173: uint16(1),
	1174: uint16(sym_self_closing_tag),
	1175: uint16(112),
	1176: uint16(1),
	1177: uint16(sym_end_tag),
	1178: uint16(206),
	1179: uint16(1),
	1180: uint16(sym_script_start_tag),
	1181: uint16(212),
	1182: uint16(1),
	1183: uint16(sym_style_start_tag),
	1184: uint16(3),
	1185: uint16(2),
	1186: uint16(sym_comment),
	1187: uint16(aux_sym__tag_value_token1),
	1188: uint16(35),
	1189: uint16(2),
	1190: uint16(sym_entity),
	1191: uint16(sym_text),
	1192: uint16(27),
	1193: uint16(2),
	1194: uint16(sym__node),
	1195: uint16(aux_sym_document_repeat1),
	1196: uint16(98),
	1197: uint16(15),
	1198: uint16(sym_doctype),
	1199: uint16(sym_element),
	1200: uint16(sym_script_element),
	1201: uint16(sym_style_element),
	1202: uint16(sym_erroneous_end_tag),
	1203: uint16(sym_if_statement),
	1204: uint16(sym_each_statement),
	1205: uint16(sym_await_statement),
	1206: uint16(sym_key_statement),
	1207: uint16(sym_snippet_statement),
	1208: uint16(sym_expression),
	1209: uint16(sym_html_tag),
	1210: uint16(sym_const_tag),
	1211: uint16(sym_debug_tag),
	1212: uint16(sym_render_tag),
	1213: uint16(19),
	1214: uint16(29),
	1215: uint16(1),
	1216: uint16(anon_sym_LT_BANG),
	1217: uint16(31),
	1218: uint16(1),
	1219: uint16(anon_sym_LT),
	1220: uint16(37),
	1221: uint16(1),
	1222: uint16(anon_sym_LBRACE),
	1223: uint16(43),
	1224: uint16(1),
	1225: uint16(anon_sym_LT_SLASH),
	1226: uint16(47),
	1227: uint16(1),
	1228: uint16(sym__implicit_end_tag),
	1229: uint16(4),
	1230: uint16(1),
	1231: uint16(sym_if_start),
	1232: uint16(8),
	1233: uint16(1),
	1234: uint16(sym_await_start),
	1235: uint16(13),
	1236: uint16(1),
	1237: uint16(sym_each_start),
	1238: uint16(17),
	1239: uint16(1),
	1240: uint16(sym_start_tag),
	1241: uint16(29),
	1242: uint16(1),
	1243: uint16(sym_key_start),
	1244: uint16(30),
	1245: uint16(1),
	1246: uint16(sym_snippet_start),
	1247: uint16(99),
	1248: uint16(1),
	1249: uint16(sym_self_closing_tag),
	1250: uint16(100),
	1251: uint16(1),
	1252: uint16(sym_end_tag),
	1253: uint16(206),
	1254: uint16(1),
	1255: uint16(sym_script_start_tag),
	1256: uint16(212),
	1257: uint16(1),
	1258: uint16(sym_style_start_tag),
	1259: uint16(3),
	1260: uint16(2),
	1261: uint16(sym_comment),
	1262: uint16(aux_sym__tag_value_token1),
	1263: uint16(35),
	1264: uint16(2),
	1265: uint16(sym_entity),
	1266: uint16(sym_text),
	1267: uint16(16),
	1268: uint16(2),
	1269: uint16(sym__node),
	1270: uint16(aux_sym_document_repeat1),
	1271: uint16(98),
	1272: uint16(15),
	1273: uint16(sym_doctype),
	1274: uint16(sym_element),
	1275: uint16(sym_script_element),
	1276: uint16(sym_style_element),
	1277: uint16(sym_erroneous_end_tag),
	1278: uint16(sym_if_statement),
	1279: uint16(sym_each_statement),
	1280: uint16(sym_await_statement),
	1281: uint16(sym_key_statement),
	1282: uint16(sym_snippet_statement),
	1283: uint16(sym_expression),
	1284: uint16(sym_html_tag),
	1285: uint16(sym_const_tag),
	1286: uint16(sym_debug_tag),
	1287: uint16(sym_render_tag),
	1288: uint16(18),
	1289: uint16(7),
	1290: uint16(1),
	1291: uint16(anon_sym_LT_BANG),
	1292: uint16(9),
	1293: uint16(1),
	1294: uint16(anon_sym_LT),
	1295: uint16(11),
	1296: uint16(1),
	1297: uint16(anon_sym_LT_SLASH),
	1298: uint16(49),
	1299: uint16(1),
	1300: uint16(anon_sym_LBRACE),
	1301: uint16(3),
	1302: uint16(1),
	1303: uint16(sym_if_start),
	1304: uint16(6),
	1305: uint16(1),
	1306: uint16(sym_await_start),
	1307: uint16(10),
	1308: uint16(1),
	1309: uint16(sym_each_start),
	1310: uint16(15),
	1311: uint16(1),
	1312: uint16(sym_start_tag),
	1313: uint16(20),
	1314: uint16(1),
	1315: uint16(sym_snippet_start),
	1316: uint16(33),
	1317: uint16(1),
	1318: uint16(sym_key_start),
	1319: uint16(108),
	1320: uint16(1),
	1321: uint16(sym_self_closing_tag),
	1322: uint16(118),
	1323: uint16(1),
	1324: uint16(sym_key_end),
	1325: uint16(209),
	1326: uint16(1),
	1327: uint16(sym_style_start_tag),
	1328: uint16(210),
	1329: uint16(1),
	1330: uint16(sym_script_start_tag),
	1331: uint16(3),
	1332: uint16(2),
	1333: uint16(sym_comment),
	1334: uint16(aux_sym__tag_value_token1),
	1335: uint16(13),
	1336: uint16(2),
	1337: uint16(sym_entity),
	1338: uint16(sym_text),
	1339: uint16(23),
	1340: uint16(2),
	1341: uint16(sym__node),
	1342: uint16(aux_sym_document_repeat1),
	1343: uint16(140),
	1344: uint16(15),
	1345: uint16(sym_doctype),
	1346: uint16(sym_element),
	1347: uint16(sym_script_element),
	1348: uint16(sym_style_element),
	1349: uint16(sym_erroneous_end_tag),
	1350: uint16(sym_if_statement),
	1351: uint16(sym_each_statement),
	1352: uint16(sym_await_statement),
	1353: uint16(sym_key_statement),
	1354: uint16(sym_snippet_statement),
	1355: uint16(sym_expression),
	1356: uint16(sym_html_tag),
	1357: uint16(sym_const_tag),
	1358: uint16(sym_debug_tag),
	1359: uint16(sym_render_tag),
	1360: uint16(18),
	1361: uint16(7),
	1362: uint16(1),
	1363: uint16(anon_sym_LT_BANG),
	1364: uint16(9),
	1365: uint16(1),
	1366: uint16(anon_sym_LT),
	1367: uint16(11),
	1368: uint16(1),
	1369: uint16(anon_sym_LT_SLASH),
	1370: uint16(51),
	1371: uint16(1),
	1372: uint16(anon_sym_LBRACE),
	1373: uint16(3),
	1374: uint16(1),
	1375: uint16(sym_if_start),
	1376: uint16(6),
	1377: uint16(1),
	1378: uint16(sym_await_start),
	1379: uint16(10),
	1380: uint16(1),
	1381: uint16(sym_each_start),
	1382: uint16(15),
	1383: uint16(1),
	1384: uint16(sym_start_tag),
	1385: uint16(20),
	1386: uint16(1),
	1387: uint16(sym_snippet_start),
	1388: uint16(33),
	1389: uint16(1),
	1390: uint16(sym_key_start),
	1391: uint16(108),
	1392: uint16(1),
	1393: uint16(sym_self_closing_tag),
	1394: uint16(136),
	1395: uint16(1),
	1396: uint16(sym_each_end),
	1397: uint16(209),
	1398: uint16(1),
	1399: uint16(sym_style_start_tag),
	1400: uint16(210),
	1401: uint16(1),
	1402: uint16(sym_script_start_tag),
	1403: uint16(3),
	1404: uint16(2),
	1405: uint16(sym_comment),
	1406: uint16(aux_sym__tag_value_token1),
	1407: uint16(13),
	1408: uint16(2),
	1409: uint16(sym_entity),
	1410: uint16(sym_text),
	1411: uint16(23),
	1412: uint16(2),
	1413: uint16(sym__node),
	1414: uint16(aux_sym_document_repeat1),
	1415: uint16(140),
	1416: uint16(15),
	1417: uint16(sym_doctype),
	1418: uint16(sym_element),
	1419: uint16(sym_script_element),
	1420: uint16(sym_style_element),
	1421: uint16(sym_erroneous_end_tag),
	1422: uint16(sym_if_statement),
	1423: uint16(sym_each_statement),
	1424: uint16(sym_await_statement),
	1425: uint16(sym_key_statement),
	1426: uint16(sym_snippet_statement),
	1427: uint16(sym_expression),
	1428: uint16(sym_html_tag),
	1429: uint16(sym_const_tag),
	1430: uint16(sym_debug_tag),
	1431: uint16(sym_render_tag),
	1432: uint16(18),
	1433: uint16(7),
	1434: uint16(1),
	1435: uint16(anon_sym_LT_BANG),
	1436: uint16(9),
	1437: uint16(1),
	1438: uint16(anon_sym_LT),
	1439: uint16(11),
	1440: uint16(1),
	1441: uint16(anon_sym_LT_SLASH),
	1442: uint16(53),
	1443: uint16(1),
	1444: uint16(anon_sym_LBRACE),
	1445: uint16(3),
	1446: uint16(1),
	1447: uint16(sym_if_start),
	1448: uint16(6),
	1449: uint16(1),
	1450: uint16(sym_await_start),
	1451: uint16(10),
	1452: uint16(1),
	1453: uint16(sym_each_start),
	1454: uint16(15),
	1455: uint16(1),
	1456: uint16(sym_start_tag),
	1457: uint16(20),
	1458: uint16(1),
	1459: uint16(sym_snippet_start),
	1460: uint16(33),
	1461: uint16(1),
	1462: uint16(sym_key_start),
	1463: uint16(62),
	1464: uint16(1),
	1465: uint16(sym_snippet_end),
	1466: uint16(108),
	1467: uint16(1),
	1468: uint16(sym_self_closing_tag),
	1469: uint16(209),
	1470: uint16(1),
	1471: uint16(sym_style_start_tag),
	1472: uint16(210),
	1473: uint16(1),
	1474: uint16(sym_script_start_tag),
	1475: uint16(3),
	1476: uint16(2),
	1477: uint16(sym_comment),
	1478: uint16(aux_sym__tag_value_token1),
	1479: uint16(13),
	1480: uint16(2),
	1481: uint16(sym_entity),
	1482: uint16(sym_text),
	1483: uint16(22),
	1484: uint16(2),
	1485: uint16(sym__node),
	1486: uint16(aux_sym_document_repeat1),
	1487: uint16(140),
	1488: uint16(15),
	1489: uint16(sym_doctype),
	1490: uint16(sym_element),
	1491: uint16(sym_script_element),
	1492: uint16(sym_style_element),
	1493: uint16(sym_erroneous_end_tag),
	1494: uint16(sym_if_statement),
	1495: uint16(sym_each_statement),
	1496: uint16(sym_await_statement),
	1497: uint16(sym_key_statement),
	1498: uint16(sym_snippet_statement),
	1499: uint16(sym_expression),
	1500: uint16(sym_html_tag),
	1501: uint16(sym_const_tag),
	1502: uint16(sym_debug_tag),
	1503: uint16(sym_render_tag),
	1504: uint16(18),
	1505: uint16(7),
	1506: uint16(1),
	1507: uint16(anon_sym_LT_BANG),
	1508: uint16(9),
	1509: uint16(1),
	1510: uint16(anon_sym_LT),
	1511: uint16(11),
	1512: uint16(1),
	1513: uint16(anon_sym_LT_SLASH),
	1514: uint16(55),
	1515: uint16(1),
	1516: uint16(anon_sym_LBRACE),
	1517: uint16(3),
	1518: uint16(1),
	1519: uint16(sym_if_start),
	1520: uint16(6),
	1521: uint16(1),
	1522: uint16(sym_await_start),
	1523: uint16(10),
	1524: uint16(1),
	1525: uint16(sym_each_start),
	1526: uint16(15),
	1527: uint16(1),
	1528: uint16(sym_start_tag),
	1529: uint16(20),
	1530: uint16(1),
	1531: uint16(sym_snippet_start),
	1532: uint16(33),
	1533: uint16(1),
	1534: uint16(sym_key_start),
	1535: uint16(76),
	1536: uint16(1),
	1537: uint16(sym_key_end),
	1538: uint16(108),
	1539: uint16(1),
	1540: uint16(sym_self_closing_tag),
	1541: uint16(209),
	1542: uint16(1),
	1543: uint16(sym_style_start_tag),
	1544: uint16(210),
	1545: uint16(1),
	1546: uint16(sym_script_start_tag),
	1547: uint16(3),
	1548: uint16(2),
	1549: uint16(sym_comment),
	1550: uint16(aux_sym__tag_value_token1),
	1551: uint16(13),
	1552: uint16(2),
	1553: uint16(sym_entity),
	1554: uint16(sym_text),
	1555: uint16(23),
	1556: uint16(2),
	1557: uint16(sym__node),
	1558: uint16(aux_sym_document_repeat1),
	1559: uint16(140),
	1560: uint16(15),
	1561: uint16(sym_doctype),
	1562: uint16(sym_element),
	1563: uint16(sym_script_element),
	1564: uint16(sym_style_element),
	1565: uint16(sym_erroneous_end_tag),
	1566: uint16(sym_if_statement),
	1567: uint16(sym_each_statement),
	1568: uint16(sym_await_statement),
	1569: uint16(sym_key_statement),
	1570: uint16(sym_snippet_statement),
	1571: uint16(sym_expression),
	1572: uint16(sym_html_tag),
	1573: uint16(sym_const_tag),
	1574: uint16(sym_debug_tag),
	1575: uint16(sym_render_tag),
	1576: uint16(18),
	1577: uint16(7),
	1578: uint16(1),
	1579: uint16(anon_sym_LT_BANG),
	1580: uint16(9),
	1581: uint16(1),
	1582: uint16(anon_sym_LT),
	1583: uint16(11),
	1584: uint16(1),
	1585: uint16(anon_sym_LT_SLASH),
	1586: uint16(53),
	1587: uint16(1),
	1588: uint16(anon_sym_LBRACE),
	1589: uint16(3),
	1590: uint16(1),
	1591: uint16(sym_if_start),
	1592: uint16(6),
	1593: uint16(1),
	1594: uint16(sym_await_start),
	1595: uint16(10),
	1596: uint16(1),
	1597: uint16(sym_each_start),
	1598: uint16(15),
	1599: uint16(1),
	1600: uint16(sym_start_tag),
	1601: uint16(20),
	1602: uint16(1),
	1603: uint16(sym_snippet_start),
	1604: uint16(33),
	1605: uint16(1),
	1606: uint16(sym_key_start),
	1607: uint16(77),
	1608: uint16(1),
	1609: uint16(sym_snippet_end),
	1610: uint16(108),
	1611: uint16(1),
	1612: uint16(sym_self_closing_tag),
	1613: uint16(209),
	1614: uint16(1),
	1615: uint16(sym_style_start_tag),
	1616: uint16(210),
	1617: uint16(1),
	1618: uint16(sym_script_start_tag),
	1619: uint16(3),
	1620: uint16(2),
	1621: uint16(sym_comment),
	1622: uint16(aux_sym__tag_value_token1),
	1623: uint16(13),
	1624: uint16(2),
	1625: uint16(sym_entity),
	1626: uint16(sym_text),
	1627: uint16(23),
	1628: uint16(2),
	1629: uint16(sym__node),
	1630: uint16(aux_sym_document_repeat1),
	1631: uint16(140),
	1632: uint16(15),
	1633: uint16(sym_doctype),
	1634: uint16(sym_element),
	1635: uint16(sym_script_element),
	1636: uint16(sym_style_element),
	1637: uint16(sym_erroneous_end_tag),
	1638: uint16(sym_if_statement),
	1639: uint16(sym_each_statement),
	1640: uint16(sym_await_statement),
	1641: uint16(sym_key_statement),
	1642: uint16(sym_snippet_statement),
	1643: uint16(sym_expression),
	1644: uint16(sym_html_tag),
	1645: uint16(sym_const_tag),
	1646: uint16(sym_debug_tag),
	1647: uint16(sym_render_tag),
	1648: uint16(18),
	1649: uint16(57),
	1650: uint16(1),
	1652: uint16(59),
	1653: uint16(1),
	1654: uint16(anon_sym_LT_BANG),
	1655: uint16(62),
	1656: uint16(1),
	1657: uint16(anon_sym_LT),
	1658: uint16(65),
	1659: uint16(1),
	1660: uint16(anon_sym_LT_SLASH),
	1661: uint16(71),
	1662: uint16(1),
	1663: uint16(anon_sym_LBRACE),
	1664: uint16(3),
	1665: uint16(1),
	1666: uint16(sym_if_start),
	1667: uint16(6),
	1668: uint16(1),
	1669: uint16(sym_await_start),
	1670: uint16(10),
	1671: uint16(1),
	1672: uint16(sym_each_start),
	1673: uint16(15),
	1674: uint16(1),
	1675: uint16(sym_start_tag),
	1676: uint16(20),
	1677: uint16(1),
	1678: uint16(sym_snippet_start),
	1679: uint16(33),
	1680: uint16(1),
	1681: uint16(sym_key_start),
	1682: uint16(108),
	1683: uint16(1),
	1684: uint16(sym_self_closing_tag),
	1685: uint16(209),
	1686: uint16(1),
	1687: uint16(sym_style_start_tag),
	1688: uint16(210),
	1689: uint16(1),
	1690: uint16(sym_script_start_tag),
	1691: uint16(3),
	1692: uint16(2),
	1693: uint16(sym_comment),
	1694: uint16(aux_sym__tag_value_token1),
	1695: uint16(68),
	1696: uint16(2),
	1697: uint16(sym_entity),
	1698: uint16(sym_text),
	1699: uint16(23),
	1700: uint16(2),
	1701: uint16(sym__node),
	1702: uint16(aux_sym_document_repeat1),
	1703: uint16(140),
	1704: uint16(15),
	1705: uint16(sym_doctype),
	1706: uint16(sym_element),
	1707: uint16(sym_script_element),
	1708: uint16(sym_style_element),
	1709: uint16(sym_erroneous_end_tag),
	1710: uint16(sym_if_statement),
	1711: uint16(sym_each_statement),
	1712: uint16(sym_await_statement),
	1713: uint16(sym_key_statement),
	1714: uint16(sym_snippet_statement),
	1715: uint16(sym_expression),
	1716: uint16(sym_html_tag),
	1717: uint16(sym_const_tag),
	1718: uint16(sym_debug_tag),
	1719: uint16(sym_render_tag),
	1720: uint16(18),
	1721: uint16(7),
	1722: uint16(1),
	1723: uint16(anon_sym_LT_BANG),
	1724: uint16(9),
	1725: uint16(1),
	1726: uint16(anon_sym_LT),
	1727: uint16(11),
	1728: uint16(1),
	1729: uint16(anon_sym_LT_SLASH),
	1730: uint16(74),
	1731: uint16(1),
	1732: uint16(anon_sym_LBRACE),
	1733: uint16(3),
	1734: uint16(1),
	1735: uint16(sym_if_start),
	1736: uint16(6),
	1737: uint16(1),
	1738: uint16(sym_await_start),
	1739: uint16(10),
	1740: uint16(1),
	1741: uint16(sym_each_start),
	1742: uint16(15),
	1743: uint16(1),
	1744: uint16(sym_start_tag),
	1745: uint16(20),
	1746: uint16(1),
	1747: uint16(sym_snippet_start),
	1748: uint16(33),
	1749: uint16(1),
	1750: uint16(sym_key_start),
	1751: uint16(90),
	1752: uint16(1),
	1753: uint16(sym_each_end),
	1754: uint16(108),
	1755: uint16(1),
	1756: uint16(sym_self_closing_tag),
	1757: uint16(209),
	1758: uint16(1),
	1759: uint16(sym_style_start_tag),
	1760: uint16(210),
	1761: uint16(1),
	1762: uint16(sym_script_start_tag),
	1763: uint16(3),
	1764: uint16(2),
	1765: uint16(sym_comment),
	1766: uint16(aux_sym__tag_value_token1),
	1767: uint16(13),
	1768: uint16(2),
	1769: uint16(sym_entity),
	1770: uint16(sym_text),
	1771: uint16(23),
	1772: uint16(2),
	1773: uint16(sym__node),
	1774: uint16(aux_sym_document_repeat1),
	1775: uint16(140),
	1776: uint16(15),
	1777: uint16(sym_doctype),
	1778: uint16(sym_element),
	1779: uint16(sym_script_element),
	1780: uint16(sym_style_element),
	1781: uint16(sym_erroneous_end_tag),
	1782: uint16(sym_if_statement),
	1783: uint16(sym_each_statement),
	1784: uint16(sym_await_statement),
	1785: uint16(sym_key_statement),
	1786: uint16(sym_snippet_statement),
	1787: uint16(sym_expression),
	1788: uint16(sym_html_tag),
	1789: uint16(sym_const_tag),
	1790: uint16(sym_debug_tag),
	1791: uint16(sym_render_tag),
	1792: uint16(18),
	1793: uint16(7),
	1794: uint16(1),
	1795: uint16(anon_sym_LT_BANG),
	1796: uint16(9),
	1797: uint16(1),
	1798: uint16(anon_sym_LT),
	1799: uint16(11),
	1800: uint16(1),
	1801: uint16(anon_sym_LT_SLASH),
	1802: uint16(74),
	1803: uint16(1),
	1804: uint16(anon_sym_LBRACE),
	1805: uint16(3),
	1806: uint16(1),
	1807: uint16(sym_if_start),
	1808: uint16(6),
	1809: uint16(1),
	1810: uint16(sym_await_start),
	1811: uint16(10),
	1812: uint16(1),
	1813: uint16(sym_each_start),
	1814: uint16(15),
	1815: uint16(1),
	1816: uint16(sym_start_tag),
	1817: uint16(20),
	1818: uint16(1),
	1819: uint16(sym_snippet_start),
	1820: uint16(33),
	1821: uint16(1),
	1822: uint16(sym_key_start),
	1823: uint16(90),
	1824: uint16(1),
	1825: uint16(sym_each_end),
	1826: uint16(108),
	1827: uint16(1),
	1828: uint16(sym_self_closing_tag),
	1829: uint16(209),
	1830: uint16(1),
	1831: uint16(sym_style_start_tag),
	1832: uint16(210),
	1833: uint16(1),
	1834: uint16(sym_script_start_tag),
	1835: uint16(3),
	1836: uint16(2),
	1837: uint16(sym_comment),
	1838: uint16(aux_sym__tag_value_token1),
	1839: uint16(13),
	1840: uint16(2),
	1841: uint16(sym_entity),
	1842: uint16(sym_text),
	1843: uint16(26),
	1844: uint16(2),
	1845: uint16(sym__node),
	1846: uint16(aux_sym_document_repeat1),
	1847: uint16(140),
	1848: uint16(15),
	1849: uint16(sym_doctype),
	1850: uint16(sym_element),
	1851: uint16(sym_script_element),
	1852: uint16(sym_style_element),
	1853: uint16(sym_erroneous_end_tag),
	1854: uint16(sym_if_statement),
	1855: uint16(sym_each_statement),
	1856: uint16(sym_await_statement),
	1857: uint16(sym_key_statement),
	1858: uint16(sym_snippet_statement),
	1859: uint16(sym_expression),
	1860: uint16(sym_html_tag),
	1861: uint16(sym_const_tag),
	1862: uint16(sym_debug_tag),
	1863: uint16(sym_render_tag),
	1864: uint16(18),
	1865: uint16(7),
	1866: uint16(1),
	1867: uint16(anon_sym_LT_BANG),
	1868: uint16(9),
	1869: uint16(1),
	1870: uint16(anon_sym_LT),
	1871: uint16(11),
	1872: uint16(1),
	1873: uint16(anon_sym_LT_SLASH),
	1874: uint16(74),
	1875: uint16(1),
	1876: uint16(anon_sym_LBRACE),
	1877: uint16(3),
	1878: uint16(1),
	1879: uint16(sym_if_start),
	1880: uint16(6),
	1881: uint16(1),
	1882: uint16(sym_await_start),
	1883: uint16(10),
	1884: uint16(1),
	1885: uint16(sym_each_start),
	1886: uint16(15),
	1887: uint16(1),
	1888: uint16(sym_start_tag),
	1889: uint16(20),
	1890: uint16(1),
	1891: uint16(sym_snippet_start),
	1892: uint16(33),
	1893: uint16(1),
	1894: uint16(sym_key_start),
	1895: uint16(96),
	1896: uint16(1),
	1897: uint16(sym_each_end),
	1898: uint16(108),
	1899: uint16(1),
	1900: uint16(sym_self_closing_tag),
	1901: uint16(209),
	1902: uint16(1),
	1903: uint16(sym_style_start_tag),
	1904: uint16(210),
	1905: uint16(1),
	1906: uint16(sym_script_start_tag),
	1907: uint16(3),
	1908: uint16(2),
	1909: uint16(sym_comment),
	1910: uint16(aux_sym__tag_value_token1),
	1911: uint16(13),
	1912: uint16(2),
	1913: uint16(sym_entity),
	1914: uint16(sym_text),
	1915: uint16(23),
	1916: uint16(2),
	1917: uint16(sym__node),
	1918: uint16(aux_sym_document_repeat1),
	1919: uint16(140),
	1920: uint16(15),
	1921: uint16(sym_doctype),
	1922: uint16(sym_element),
	1923: uint16(sym_script_element),
	1924: uint16(sym_style_element),
	1925: uint16(sym_erroneous_end_tag),
	1926: uint16(sym_if_statement),
	1927: uint16(sym_each_statement),
	1928: uint16(sym_await_statement),
	1929: uint16(sym_key_statement),
	1930: uint16(sym_snippet_statement),
	1931: uint16(sym_expression),
	1932: uint16(sym_html_tag),
	1933: uint16(sym_const_tag),
	1934: uint16(sym_debug_tag),
	1935: uint16(sym_render_tag),
	1936: uint16(18),
	1937: uint16(57),
	1938: uint16(1),
	1939: uint16(sym__implicit_end_tag),
	1940: uint16(76),
	1941: uint16(1),
	1942: uint16(anon_sym_LT_BANG),
	1943: uint16(79),
	1944: uint16(1),
	1945: uint16(anon_sym_LT),
	1946: uint16(82),
	1947: uint16(1),
	1948: uint16(anon_sym_LT_SLASH),
	1949: uint16(88),
	1950: uint16(1),
	1951: uint16(anon_sym_LBRACE),
	1952: uint16(4),
	1953: uint16(1),
	1954: uint16(sym_if_start),
	1955: uint16(8),
	1956: uint16(1),
	1957: uint16(sym_await_start),
	1958: uint16(13),
	1959: uint16(1),
	1960: uint16(sym_each_start),
	1961: uint16(17),
	1962: uint16(1),
	1963: uint16(sym_start_tag),
	1964: uint16(29),
	1965: uint16(1),
	1966: uint16(sym_key_start),
	1967: uint16(30),
	1968: uint16(1),
	1969: uint16(sym_snippet_start),
	1970: uint16(99),
	1971: uint16(1),
	1972: uint16(sym_self_closing_tag),
	1973: uint16(206),
	1974: uint16(1),
	1975: uint16(sym_script_start_tag),
	1976: uint16(212),
	1977: uint16(1),
	1978: uint16(sym_style_start_tag),
	1979: uint16(3),
	1980: uint16(2),
	1981: uint16(sym_comment),
	1982: uint16(aux_sym__tag_value_token1),
	1983: uint16(85),
	1984: uint16(2),
	1985: uint16(sym_entity),
	1986: uint16(sym_text),
	1987: uint16(27),
	1988: uint16(2),
	1989: uint16(sym__node),
	1990: uint16(aux_sym_document_repeat1),
	1991: uint16(98),
	1992: uint16(15),
	1993: uint16(sym_doctype),
	1994: uint16(sym_element),
	1995: uint16(sym_script_element),
	1996: uint16(sym_style_element),
	1997: uint16(sym_erroneous_end_tag),
	1998: uint16(sym_if_statement),
	1999: uint16(sym_each_statement),
	2000: uint16(sym_await_statement),
	2001: uint16(sym_key_statement),
	2002: uint16(sym_snippet_statement),
	2003: uint16(sym_expression),
	2004: uint16(sym_html_tag),
	2005: uint16(sym_const_tag),
	2006: uint16(sym_debug_tag),
	2007: uint16(sym_render_tag),
	2008: uint16(18),
	2009: uint16(7),
	2010: uint16(1),
	2011: uint16(anon_sym_LT_BANG),
	2012: uint16(9),
	2013: uint16(1),
	2014: uint16(anon_sym_LT),
	2015: uint16(11),
	2016: uint16(1),
	2017: uint16(anon_sym_LT_SLASH),
	2018: uint16(15),
	2019: uint16(1),
	2020: uint16(anon_sym_LBRACE),
	2021: uint16(91),
	2022: uint16(1),
	2024: uint16(3),
	2025: uint16(1),
	2026: uint16(sym_if_start),
	2027: uint16(6),
	2028: uint16(1),
	2029: uint16(sym_await_start),
	2030: uint16(10),
	2031: uint16(1),
	2032: uint16(sym_each_start),
	2033: uint16(15),
	2034: uint16(1),
	2035: uint16(sym_start_tag),
	2036: uint16(20),
	2037: uint16(1),
	2038: uint16(sym_snippet_start),
	2039: uint16(33),
	2040: uint16(1),
	2041: uint16(sym_key_start),
	2042: uint16(108),
	2043: uint16(1),
	2044: uint16(sym_self_closing_tag),
	2045: uint16(209),
	2046: uint16(1),
	2047: uint16(sym_style_start_tag),
	2048: uint16(210),
	2049: uint16(1),
	2050: uint16(sym_script_start_tag),
	2051: uint16(3),
	2052: uint16(2),
	2053: uint16(sym_comment),
	2054: uint16(aux_sym__tag_value_token1),
	2055: uint16(13),
	2056: uint16(2),
	2057: uint16(sym_entity),
	2058: uint16(sym_text),
	2059: uint16(23),
	2060: uint16(2),
	2061: uint16(sym__node),
	2062: uint16(aux_sym_document_repeat1),
	2063: uint16(140),
	2064: uint16(15),
	2065: uint16(sym_doctype),
	2066: uint16(sym_element),
	2067: uint16(sym_script_element),
	2068: uint16(sym_style_element),
	2069: uint16(sym_erroneous_end_tag),
	2070: uint16(sym_if_statement),
	2071: uint16(sym_each_statement),
	2072: uint16(sym_await_statement),
	2073: uint16(sym_key_statement),
	2074: uint16(sym_snippet_statement),
	2075: uint16(sym_expression),
	2076: uint16(sym_html_tag),
	2077: uint16(sym_const_tag),
	2078: uint16(sym_debug_tag),
	2079: uint16(sym_render_tag),
	2080: uint16(18),
	2081: uint16(7),
	2082: uint16(1),
	2083: uint16(anon_sym_LT_BANG),
	2084: uint16(9),
	2085: uint16(1),
	2086: uint16(anon_sym_LT),
	2087: uint16(11),
	2088: uint16(1),
	2089: uint16(anon_sym_LT_SLASH),
	2090: uint16(49),
	2091: uint16(1),
	2092: uint16(anon_sym_LBRACE),
	2093: uint16(3),
	2094: uint16(1),
	2095: uint16(sym_if_start),
	2096: uint16(6),
	2097: uint16(1),
	2098: uint16(sym_await_start),
	2099: uint16(10),
	2100: uint16(1),
	2101: uint16(sym_each_start),
	2102: uint16(15),
	2103: uint16(1),
	2104: uint16(sym_start_tag),
	2105: uint16(20),
	2106: uint16(1),
	2107: uint16(sym_snippet_start),
	2108: uint16(33),
	2109: uint16(1),
	2110: uint16(sym_key_start),
	2111: uint16(106),
	2112: uint16(1),
	2113: uint16(sym_key_end),
	2114: uint16(108),
	2115: uint16(1),
	2116: uint16(sym_self_closing_tag),
	2117: uint16(209),
	2118: uint16(1),
	2119: uint16(sym_style_start_tag),
	2120: uint16(210),
	2121: uint16(1),
	2122: uint16(sym_script_start_tag),
	2123: uint16(3),
	2124: uint16(2),
	2125: uint16(sym_comment),
	2126: uint16(aux_sym__tag_value_token1),
	2127: uint16(13),
	2128: uint16(2),
	2129: uint16(sym_entity),
	2130: uint16(sym_text),
	2131: uint16(18),
	2132: uint16(2),
	2133: uint16(sym__node),
	2134: uint16(aux_sym_document_repeat1),
	2135: uint16(140),
	2136: uint16(15),
	2137: uint16(sym_doctype),
	2138: uint16(sym_element),
	2139: uint16(sym_script_element),
	2140: uint16(sym_style_element),
	2141: uint16(sym_erroneous_end_tag),
	2142: uint16(sym_if_statement),
	2143: uint16(sym_each_statement),
	2144: uint16(sym_await_statement),
	2145: uint16(sym_key_statement),
	2146: uint16(sym_snippet_statement),
	2147: uint16(sym_expression),
	2148: uint16(sym_html_tag),
	2149: uint16(sym_const_tag),
	2150: uint16(sym_debug_tag),
	2151: uint16(sym_render_tag),
	2152: uint16(18),
	2153: uint16(7),
	2154: uint16(1),
	2155: uint16(anon_sym_LT_BANG),
	2156: uint16(9),
	2157: uint16(1),
	2158: uint16(anon_sym_LT),
	2159: uint16(11),
	2160: uint16(1),
	2161: uint16(anon_sym_LT_SLASH),
	2162: uint16(93),
	2163: uint16(1),
	2164: uint16(anon_sym_LBRACE),
	2165: uint16(3),
	2166: uint16(1),
	2167: uint16(sym_if_start),
	2168: uint16(6),
	2169: uint16(1),
	2170: uint16(sym_await_start),
	2171: uint16(10),
	2172: uint16(1),
	2173: uint16(sym_each_start),
	2174: uint16(15),
	2175: uint16(1),
	2176: uint16(sym_start_tag),
	2177: uint16(20),
	2178: uint16(1),
	2179: uint16(sym_snippet_start),
	2180: uint16(33),
	2181: uint16(1),
	2182: uint16(sym_key_start),
	2183: uint16(107),
	2184: uint16(1),
	2185: uint16(sym_snippet_end),
	2186: uint16(108),
	2187: uint16(1),
	2188: uint16(sym_self_closing_tag),
	2189: uint16(209),
	2190: uint16(1),
	2191: uint16(sym_style_start_tag),
	2192: uint16(210),
	2193: uint16(1),
	2194: uint16(sym_script_start_tag),
	2195: uint16(3),
	2196: uint16(2),
	2197: uint16(sym_comment),
	2198: uint16(aux_sym__tag_value_token1),
	2199: uint16(13),
	2200: uint16(2),
	2201: uint16(sym_entity),
	2202: uint16(sym_text),
	2203: uint16(34),
	2204: uint16(2),
	2205: uint16(sym__node),
	2206: uint16(aux_sym_document_repeat1),
	2207: uint16(140),
	2208: uint16(15),
	2209: uint16(sym_doctype),
	2210: uint16(sym_element),
	2211: uint16(sym_script_element),
	2212: uint16(sym_style_element),
	2213: uint16(sym_erroneous_end_tag),
	2214: uint16(sym_if_statement),
	2215: uint16(sym_each_statement),
	2216: uint16(sym_await_statement),
	2217: uint16(sym_key_statement),
	2218: uint16(sym_snippet_statement),
	2219: uint16(sym_expression),
	2220: uint16(sym_html_tag),
	2221: uint16(sym_const_tag),
	2222: uint16(sym_debug_tag),
	2223: uint16(sym_render_tag),
	2224: uint16(18),
	2225: uint16(7),
	2226: uint16(1),
	2227: uint16(anon_sym_LT_BANG),
	2228: uint16(9),
	2229: uint16(1),
	2230: uint16(anon_sym_LT),
	2231: uint16(11),
	2232: uint16(1),
	2233: uint16(anon_sym_LT_SLASH),
	2234: uint16(74),
	2235: uint16(1),
	2236: uint16(anon_sym_LBRACE),
	2237: uint16(3),
	2238: uint16(1),
	2239: uint16(sym_if_start),
	2240: uint16(6),
	2241: uint16(1),
	2242: uint16(sym_await_start),
	2243: uint16(10),
	2244: uint16(1),
	2245: uint16(sym_each_start),
	2246: uint16(15),
	2247: uint16(1),
	2248: uint16(sym_start_tag),
	2249: uint16(20),
	2250: uint16(1),
	2251: uint16(sym_snippet_start),
	2252: uint16(33),
	2253: uint16(1),
	2254: uint16(sym_key_start),
	2255: uint16(74),
	2256: uint16(1),
	2257: uint16(sym_each_end),
	2258: uint16(108),
	2259: uint16(1),
	2260: uint16(sym_self_closing_tag),
	2261: uint16(209),
	2262: uint16(1),
	2263: uint16(sym_style_start_tag),
	2264: uint16(210),
	2265: uint16(1),
	2266: uint16(sym_script_start_tag),
	2267: uint16(3),
	2268: uint16(2),
	2269: uint16(sym_comment),
	2270: uint16(aux_sym__tag_value_token1),
	2271: uint16(13),
	2272: uint16(2),
	2273: uint16(sym_entity),
	2274: uint16(sym_text),
	2275: uint16(24),
	2276: uint16(2),
	2277: uint16(sym__node),
	2278: uint16(aux_sym_document_repeat1),
	2279: uint16(140),
	2280: uint16(15),
	2281: uint16(sym_doctype),
	2282: uint16(sym_element),
	2283: uint16(sym_script_element),
	2284: uint16(sym_style_element),
	2285: uint16(sym_erroneous_end_tag),
	2286: uint16(sym_if_statement),
	2287: uint16(sym_each_statement),
	2288: uint16(sym_await_statement),
	2289: uint16(sym_key_statement),
	2290: uint16(sym_snippet_statement),
	2291: uint16(sym_expression),
	2292: uint16(sym_html_tag),
	2293: uint16(sym_const_tag),
	2294: uint16(sym_debug_tag),
	2295: uint16(sym_render_tag),
	2296: uint16(18),
	2297: uint16(7),
	2298: uint16(1),
	2299: uint16(anon_sym_LT_BANG),
	2300: uint16(9),
	2301: uint16(1),
	2302: uint16(anon_sym_LT),
	2303: uint16(11),
	2304: uint16(1),
	2305: uint16(anon_sym_LT_SLASH),
	2306: uint16(51),
	2307: uint16(1),
	2308: uint16(anon_sym_LBRACE),
	2309: uint16(3),
	2310: uint16(1),
	2311: uint16(sym_if_start),
	2312: uint16(6),
	2313: uint16(1),
	2314: uint16(sym_await_start),
	2315: uint16(10),
	2316: uint16(1),
	2317: uint16(sym_each_start),
	2318: uint16(15),
	2319: uint16(1),
	2320: uint16(sym_start_tag),
	2321: uint16(20),
	2322: uint16(1),
	2323: uint16(sym_snippet_start),
	2324: uint16(33),
	2325: uint16(1),
	2326: uint16(sym_key_start),
	2327: uint16(108),
	2328: uint16(1),
	2329: uint16(sym_self_closing_tag),
	2330: uint16(116),
	2331: uint16(1),
	2332: uint16(sym_each_end),
	2333: uint16(209),
	2334: uint16(1),
	2335: uint16(sym_style_start_tag),
	2336: uint16(210),
	2337: uint16(1),
	2338: uint16(sym_script_start_tag),
	2339: uint16(3),
	2340: uint16(2),
	2341: uint16(sym_comment),
	2342: uint16(aux_sym__tag_value_token1),
	2343: uint16(13),
	2344: uint16(2),
	2345: uint16(sym_entity),
	2346: uint16(sym_text),
	2347: uint16(35),
	2348: uint16(2),
	2349: uint16(sym__node),
	2350: uint16(aux_sym_document_repeat1),
	2351: uint16(140),
	2352: uint16(15),
	2353: uint16(sym_doctype),
	2354: uint16(sym_element),
	2355: uint16(sym_script_element),
	2356: uint16(sym_style_element),
	2357: uint16(sym_erroneous_end_tag),
	2358: uint16(sym_if_statement),
	2359: uint16(sym_each_statement),
	2360: uint16(sym_await_statement),
	2361: uint16(sym_key_statement),
	2362: uint16(sym_snippet_statement),
	2363: uint16(sym_expression),
	2364: uint16(sym_html_tag),
	2365: uint16(sym_const_tag),
	2366: uint16(sym_debug_tag),
	2367: uint16(sym_render_tag),
	2368: uint16(18),
	2369: uint16(7),
	2370: uint16(1),
	2371: uint16(anon_sym_LT_BANG),
	2372: uint16(9),
	2373: uint16(1),
	2374: uint16(anon_sym_LT),
	2375: uint16(11),
	2376: uint16(1),
	2377: uint16(anon_sym_LT_SLASH),
	2378: uint16(55),
	2379: uint16(1),
	2380: uint16(anon_sym_LBRACE),
	2381: uint16(3),
	2382: uint16(1),
	2383: uint16(sym_if_start),
	2384: uint16(6),
	2385: uint16(1),
	2386: uint16(sym_await_start),
	2387: uint16(10),
	2388: uint16(1),
	2389: uint16(sym_each_start),
	2390: uint16(15),
	2391: uint16(1),
	2392: uint16(sym_start_tag),
	2393: uint16(20),
	2394: uint16(1),
	2395: uint16(sym_snippet_start),
	2396: uint16(33),
	2397: uint16(1),
	2398: uint16(sym_key_start),
	2399: uint16(66),
	2400: uint16(1),
	2401: uint16(sym_key_end),
	2402: uint16(108),
	2403: uint16(1),
	2404: uint16(sym_self_closing_tag),
	2405: uint16(209),
	2406: uint16(1),
	2407: uint16(sym_style_start_tag),
	2408: uint16(210),
	2409: uint16(1),
	2410: uint16(sym_script_start_tag),
	2411: uint16(3),
	2412: uint16(2),
	2413: uint16(sym_comment),
	2414: uint16(aux_sym__tag_value_token1),
	2415: uint16(13),
	2416: uint16(2),
	2417: uint16(sym_entity),
	2418: uint16(sym_text),
	2419: uint16(21),
	2420: uint16(2),
	2421: uint16(sym__node),
	2422: uint16(aux_sym_document_repeat1),
	2423: uint16(140),
	2424: uint16(15),
	2425: uint16(sym_doctype),
	2426: uint16(sym_element),
	2427: uint16(sym_script_element),
	2428: uint16(sym_style_element),
	2429: uint16(sym_erroneous_end_tag),
	2430: uint16(sym_if_statement),
	2431: uint16(sym_each_statement),
	2432: uint16(sym_await_statement),
	2433: uint16(sym_key_statement),
	2434: uint16(sym_snippet_statement),
	2435: uint16(sym_expression),
	2436: uint16(sym_html_tag),
	2437: uint16(sym_const_tag),
	2438: uint16(sym_debug_tag),
	2439: uint16(sym_render_tag),
	2440: uint16(18),
	2441: uint16(7),
	2442: uint16(1),
	2443: uint16(anon_sym_LT_BANG),
	2444: uint16(9),
	2445: uint16(1),
	2446: uint16(anon_sym_LT),
	2447: uint16(11),
	2448: uint16(1),
	2449: uint16(anon_sym_LT_SLASH),
	2450: uint16(93),
	2451: uint16(1),
	2452: uint16(anon_sym_LBRACE),
	2453: uint16(3),
	2454: uint16(1),
	2455: uint16(sym_if_start),
	2456: uint16(6),
	2457: uint16(1),
	2458: uint16(sym_await_start),
	2459: uint16(10),
	2460: uint16(1),
	2461: uint16(sym_each_start),
	2462: uint16(15),
	2463: uint16(1),
	2464: uint16(sym_start_tag),
	2465: uint16(20),
	2466: uint16(1),
	2467: uint16(sym_snippet_start),
	2468: uint16(33),
	2469: uint16(1),
	2470: uint16(sym_key_start),
	2471: uint16(108),
	2472: uint16(1),
	2473: uint16(sym_self_closing_tag),
	2474: uint16(119),
	2475: uint16(1),
	2476: uint16(sym_snippet_end),
	2477: uint16(209),
	2478: uint16(1),
	2479: uint16(sym_style_start_tag),
	2480: uint16(210),
	2481: uint16(1),
	2482: uint16(sym_script_start_tag),
	2483: uint16(3),
	2484: uint16(2),
	2485: uint16(sym_comment),
	2486: uint16(aux_sym__tag_value_token1),
	2487: uint16(13),
	2488: uint16(2),
	2489: uint16(sym_entity),
	2490: uint16(sym_text),
	2491: uint16(23),
	2492: uint16(2),
	2493: uint16(sym__node),
	2494: uint16(aux_sym_document_repeat1),
	2495: uint16(140),
	2496: uint16(15),
	2497: uint16(sym_doctype),
	2498: uint16(sym_element),
	2499: uint16(sym_script_element),
	2500: uint16(sym_style_element),
	2501: uint16(sym_erroneous_end_tag),
	2502: uint16(sym_if_statement),
	2503: uint16(sym_each_statement),
	2504: uint16(sym_await_statement),
	2505: uint16(sym_key_statement),
	2506: uint16(sym_snippet_statement),
	2507: uint16(sym_expression),
	2508: uint16(sym_html_tag),
	2509: uint16(sym_const_tag),
	2510: uint16(sym_debug_tag),
	2511: uint16(sym_render_tag),
	2512: uint16(18),
	2513: uint16(7),
	2514: uint16(1),
	2515: uint16(anon_sym_LT_BANG),
	2516: uint16(9),
	2517: uint16(1),
	2518: uint16(anon_sym_LT),
	2519: uint16(11),
	2520: uint16(1),
	2521: uint16(anon_sym_LT_SLASH),
	2522: uint16(51),
	2523: uint16(1),
	2524: uint16(anon_sym_LBRACE),
	2525: uint16(3),
	2526: uint16(1),
	2527: uint16(sym_if_start),
	2528: uint16(6),
	2529: uint16(1),
	2530: uint16(sym_await_start),
	2531: uint16(10),
	2532: uint16(1),
	2533: uint16(sym_each_start),
	2534: uint16(15),
	2535: uint16(1),
	2536: uint16(sym_start_tag),
	2537: uint16(20),
	2538: uint16(1),
	2539: uint16(sym_snippet_start),
	2540: uint16(33),
	2541: uint16(1),
	2542: uint16(sym_key_start),
	2543: uint16(108),
	2544: uint16(1),
	2545: uint16(sym_self_closing_tag),
	2546: uint16(130),
	2547: uint16(1),
	2548: uint16(sym_each_end),
	2549: uint16(209),
	2550: uint16(1),
	2551: uint16(sym_style_start_tag),
	2552: uint16(210),
	2553: uint16(1),
	2554: uint16(sym_script_start_tag),
	2555: uint16(3),
	2556: uint16(2),
	2557: uint16(sym_comment),
	2558: uint16(aux_sym__tag_value_token1),
	2559: uint16(13),
	2560: uint16(2),
	2561: uint16(sym_entity),
	2562: uint16(sym_text),
	2563: uint16(23),
	2564: uint16(2),
	2565: uint16(sym__node),
	2566: uint16(aux_sym_document_repeat1),
	2567: uint16(140),
	2568: uint16(15),
	2569: uint16(sym_doctype),
	2570: uint16(sym_element),
	2571: uint16(sym_script_element),
	2572: uint16(sym_style_element),
	2573: uint16(sym_erroneous_end_tag),
	2574: uint16(sym_if_statement),
	2575: uint16(sym_each_statement),
	2576: uint16(sym_await_statement),
	2577: uint16(sym_key_statement),
	2578: uint16(sym_snippet_statement),
	2579: uint16(sym_expression),
	2580: uint16(sym_html_tag),
	2581: uint16(sym_const_tag),
	2582: uint16(sym_debug_tag),
	2583: uint16(sym_render_tag),
	2584: uint16(18),
	2585: uint16(7),
	2586: uint16(1),
	2587: uint16(anon_sym_LT_BANG),
	2588: uint16(9),
	2589: uint16(1),
	2590: uint16(anon_sym_LT),
	2591: uint16(11),
	2592: uint16(1),
	2593: uint16(anon_sym_LT_SLASH),
	2594: uint16(51),
	2595: uint16(1),
	2596: uint16(anon_sym_LBRACE),
	2597: uint16(3),
	2598: uint16(1),
	2599: uint16(sym_if_start),
	2600: uint16(6),
	2601: uint16(1),
	2602: uint16(sym_await_start),
	2603: uint16(10),
	2604: uint16(1),
	2605: uint16(sym_each_start),
	2606: uint16(15),
	2607: uint16(1),
	2608: uint16(sym_start_tag),
	2609: uint16(20),
	2610: uint16(1),
	2611: uint16(sym_snippet_start),
	2612: uint16(33),
	2613: uint16(1),
	2614: uint16(sym_key_start),
	2615: uint16(108),
	2616: uint16(1),
	2617: uint16(sym_self_closing_tag),
	2618: uint16(130),
	2619: uint16(1),
	2620: uint16(sym_each_end),
	2621: uint16(209),
	2622: uint16(1),
	2623: uint16(sym_style_start_tag),
	2624: uint16(210),
	2625: uint16(1),
	2626: uint16(sym_script_start_tag),
	2627: uint16(3),
	2628: uint16(2),
	2629: uint16(sym_comment),
	2630: uint16(aux_sym__tag_value_token1),
	2631: uint16(13),
	2632: uint16(2),
	2633: uint16(sym_entity),
	2634: uint16(sym_text),
	2635: uint16(19),
	2636: uint16(2),
	2637: uint16(sym__node),
	2638: uint16(aux_sym_document_repeat1),
	2639: uint16(140),
	2640: uint16(15),
	2641: uint16(sym_doctype),
	2642: uint16(sym_element),
	2643: uint16(sym_script_element),
	2644: uint16(sym_style_element),
	2645: uint16(sym_erroneous_end_tag),
	2646: uint16(sym_if_statement),
	2647: uint16(sym_each_statement),
	2648: uint16(sym_await_statement),
	2649: uint16(sym_key_statement),
	2650: uint16(sym_snippet_statement),
	2651: uint16(sym_expression),
	2652: uint16(sym_html_tag),
	2653: uint16(sym_const_tag),
	2654: uint16(sym_debug_tag),
	2655: uint16(sym_render_tag),
	2656: uint16(17),
	2657: uint16(7),
	2658: uint16(1),
	2659: uint16(anon_sym_LT_BANG),
	2660: uint16(9),
	2661: uint16(1),
	2662: uint16(anon_sym_LT),
	2663: uint16(11),
	2664: uint16(1),
	2665: uint16(anon_sym_LT_SLASH),
	2666: uint16(95),
	2667: uint16(1),
	2668: uint16(anon_sym_LBRACE),
	2669: uint16(3),
	2670: uint16(1),
	2671: uint16(sym_if_start),
	2672: uint16(6),
	2673: uint16(1),
	2674: uint16(sym_await_start),
	2675: uint16(10),
	2676: uint16(1),
	2677: uint16(sym_each_start),
	2678: uint16(15),
	2679: uint16(1),
	2680: uint16(sym_start_tag),
	2681: uint16(20),
	2682: uint16(1),
	2683: uint16(sym_snippet_start),
	2684: uint16(33),
	2685: uint16(1),
	2686: uint16(sym_key_start),
	2687: uint16(108),
	2688: uint16(1),
	2689: uint16(sym_self_closing_tag),
	2690: uint16(209),
	2691: uint16(1),
	2692: uint16(sym_style_start_tag),
	2693: uint16(210),
	2694: uint16(1),
	2695: uint16(sym_script_start_tag),
	2696: uint16(3),
	2697: uint16(2),
	2698: uint16(sym_comment),
	2699: uint16(aux_sym__tag_value_token1),
	2700: uint16(13),
	2701: uint16(2),
	2702: uint16(sym_entity),
	2703: uint16(sym_text),
	2704: uint16(42),
	2705: uint16(2),
	2706: uint16(sym__node),
	2707: uint16(aux_sym_document_repeat1),
	2708: uint16(140),
	2709: uint16(15),
	2710: uint16(sym_doctype),
	2711: uint16(sym_element),
	2712: uint16(sym_script_element),
	2713: uint16(sym_style_element),
	2714: uint16(sym_erroneous_end_tag),
	2715: uint16(sym_if_statement),
	2716: uint16(sym_each_statement),
	2717: uint16(sym_await_statement),
	2718: uint16(sym_key_statement),
	2719: uint16(sym_snippet_statement),
	2720: uint16(sym_expression),
	2721: uint16(sym_html_tag),
	2722: uint16(sym_const_tag),
	2723: uint16(sym_debug_tag),
	2724: uint16(sym_render_tag),
	2725: uint16(17),
	2726: uint16(7),
	2727: uint16(1),
	2728: uint16(anon_sym_LT_BANG),
	2729: uint16(9),
	2730: uint16(1),
	2731: uint16(anon_sym_LT),
	2732: uint16(11),
	2733: uint16(1),
	2734: uint16(anon_sym_LT_SLASH),
	2735: uint16(98),
	2736: uint16(1),
	2737: uint16(anon_sym_LBRACE),
	2738: uint16(3),
	2739: uint16(1),
	2740: uint16(sym_if_start),
	2741: uint16(6),
	2742: uint16(1),
	2743: uint16(sym_await_start),
	2744: uint16(10),
	2745: uint16(1),
	2746: uint16(sym_each_start),
	2747: uint16(15),
	2748: uint16(1),
	2749: uint16(sym_start_tag),
	2750: uint16(20),
	2751: uint16(1),
	2752: uint16(sym_snippet_start),
	2753: uint16(33),
	2754: uint16(1),
	2755: uint16(sym_key_start),
	2756: uint16(108),
	2757: uint16(1),
	2758: uint16(sym_self_closing_tag),
	2759: uint16(209),
	2760: uint16(1),
	2761: uint16(sym_style_start_tag),
	2762: uint16(210),
	2763: uint16(1),
	2764: uint16(sym_script_start_tag),
	2765: uint16(3),
	2766: uint16(2),
	2767: uint16(sym_comment),
	2768: uint16(aux_sym__tag_value_token1),
	2769: uint16(13),
	2770: uint16(2),
	2771: uint16(sym_entity),
	2772: uint16(sym_text),
	2773: uint16(41),
	2774: uint16(2),
	2775: uint16(sym__node),
	2776: uint16(aux_sym_document_repeat1),
	2777: uint16(140),
	2778: uint16(15),
	2779: uint16(sym_doctype),
	2780: uint16(sym_element),
	2781: uint16(sym_script_element),
	2782: uint16(sym_style_element),
	2783: uint16(sym_erroneous_end_tag),
	2784: uint16(sym_if_statement),
	2785: uint16(sym_each_statement),
	2786: uint16(sym_await_statement),
	2787: uint16(sym_key_statement),
	2788: uint16(sym_snippet_statement),
	2789: uint16(sym_expression),
	2790: uint16(sym_html_tag),
	2791: uint16(sym_const_tag),
	2792: uint16(sym_debug_tag),
	2793: uint16(sym_render_tag),
	2794: uint16(17),
	2795: uint16(7),
	2796: uint16(1),
	2797: uint16(anon_sym_LT_BANG),
	2798: uint16(9),
	2799: uint16(1),
	2800: uint16(anon_sym_LT),
	2801: uint16(11),
	2802: uint16(1),
	2803: uint16(anon_sym_LT_SLASH),
	2804: uint16(101),
	2805: uint16(1),
	2806: uint16(anon_sym_LBRACE),
	2807: uint16(3),
	2808: uint16(1),
	2809: uint16(sym_if_start),
	2810: uint16(6),
	2811: uint16(1),
	2812: uint16(sym_await_start),
	2813: uint16(10),
	2814: uint16(1),
	2815: uint16(sym_each_start),
	2816: uint16(15),
	2817: uint16(1),
	2818: uint16(sym_start_tag),
	2819: uint16(20),
	2820: uint16(1),
	2821: uint16(sym_snippet_start),
	2822: uint16(33),
	2823: uint16(1),
	2824: uint16(sym_key_start),
	2825: uint16(108),
	2826: uint16(1),
	2827: uint16(sym_self_closing_tag),
	2828: uint16(209),
	2829: uint16(1),
	2830: uint16(sym_style_start_tag),
	2831: uint16(210),
	2832: uint16(1),
	2833: uint16(sym_script_start_tag),
	2834: uint16(3),
	2835: uint16(2),
	2836: uint16(sym_comment),
	2837: uint16(aux_sym__tag_value_token1),
	2838: uint16(13),
	2839: uint16(2),
	2840: uint16(sym_entity),
	2841: uint16(sym_text),
	2842: uint16(23),
	2843: uint16(2),
	2844: uint16(sym__node),
	2845: uint16(aux_sym_document_repeat1),
	2846: uint16(140),
	2847: uint16(15),
	2848: uint16(sym_doctype),
	2849: uint16(sym_element),
	2850: uint16(sym_script_element),
	2851: uint16(sym_style_element),
	2852: uint16(sym_erroneous_end_tag),
	2853: uint16(sym_if_statement),
	2854: uint16(sym_each_statement),
	2855: uint16(sym_await_statement),
	2856: uint16(sym_key_statement),
	2857: uint16(sym_snippet_statement),
	2858: uint16(sym_expression),
	2859: uint16(sym_html_tag),
	2860: uint16(sym_const_tag),
	2861: uint16(sym_debug_tag),
	2862: uint16(sym_render_tag),
	2863: uint16(17),
	2864: uint16(7),
	2865: uint16(1),
	2866: uint16(anon_sym_LT_BANG),
	2867: uint16(9),
	2868: uint16(1),
	2869: uint16(anon_sym_LT),
	2870: uint16(11),
	2871: uint16(1),
	2872: uint16(anon_sym_LT_SLASH),
	2873: uint16(104),
	2874: uint16(1),
	2875: uint16(anon_sym_LBRACE),
	2876: uint16(3),
	2877: uint16(1),
	2878: uint16(sym_if_start),
	2879: uint16(6),
	2880: uint16(1),
	2881: uint16(sym_await_start),
	2882: uint16(10),
	2883: uint16(1),
	2884: uint16(sym_each_start),
	2885: uint16(15),
	2886: uint16(1),
	2887: uint16(sym_start_tag),
	2888: uint16(20),
	2889: uint16(1),
	2890: uint16(sym_snippet_start),
	2891: uint16(33),
	2892: uint16(1),
	2893: uint16(sym_key_start),
	2894: uint16(108),
	2895: uint16(1),
	2896: uint16(sym_self_closing_tag),
	2897: uint16(209),
	2898: uint16(1),
	2899: uint16(sym_style_start_tag),
	2900: uint16(210),
	2901: uint16(1),
	2902: uint16(sym_script_start_tag),
	2903: uint16(3),
	2904: uint16(2),
	2905: uint16(sym_comment),
	2906: uint16(aux_sym__tag_value_token1),
	2907: uint16(13),
	2908: uint16(2),
	2909: uint16(sym_entity),
	2910: uint16(sym_text),
	2911: uint16(46),
	2912: uint16(2),
	2913: uint16(sym__node),
	2914: uint16(aux_sym_document_repeat1),
	2915: uint16(140),
	2916: uint16(15),
	2917: uint16(sym_doctype),
	2918: uint16(sym_element),
	2919: uint16(sym_script_element),
	2920: uint16(sym_style_element),
	2921: uint16(sym_erroneous_end_tag),
	2922: uint16(sym_if_statement),
	2923: uint16(sym_each_statement),
	2924: uint16(sym_await_statement),
	2925: uint16(sym_key_statement),
	2926: uint16(sym_snippet_statement),
	2927: uint16(sym_expression),
	2928: uint16(sym_html_tag),
	2929: uint16(sym_const_tag),
	2930: uint16(sym_debug_tag),
	2931: uint16(sym_render_tag),
	2932: uint16(17),
	2933: uint16(7),
	2934: uint16(1),
	2935: uint16(anon_sym_LT_BANG),
	2936: uint16(9),
	2937: uint16(1),
	2938: uint16(anon_sym_LT),
	2939: uint16(11),
	2940: uint16(1),
	2941: uint16(anon_sym_LT_SLASH),
	2942: uint16(107),
	2943: uint16(1),
	2944: uint16(anon_sym_LBRACE),
	2945: uint16(3),
	2946: uint16(1),
	2947: uint16(sym_if_start),
	2948: uint16(6),
	2949: uint16(1),
	2950: uint16(sym_await_start),
	2951: uint16(10),
	2952: uint16(1),
	2953: uint16(sym_each_start),
	2954: uint16(15),
	2955: uint16(1),
	2956: uint16(sym_start_tag),
	2957: uint16(20),
	2958: uint16(1),
	2959: uint16(sym_snippet_start),
	2960: uint16(33),
	2961: uint16(1),
	2962: uint16(sym_key_start),
	2963: uint16(108),
	2964: uint16(1),
	2965: uint16(sym_self_closing_tag),
	2966: uint16(209),
	2967: uint16(1),
	2968: uint16(sym_style_start_tag),
	2969: uint16(210),
	2970: uint16(1),
	2971: uint16(sym_script_start_tag),
	2972: uint16(3),
	2973: uint16(2),
	2974: uint16(sym_comment),
	2975: uint16(aux_sym__tag_value_token1),
	2976: uint16(13),
	2977: uint16(2),
	2978: uint16(sym_entity),
	2979: uint16(sym_text),
	2980: uint16(23),
	2981: uint16(2),
	2982: uint16(sym__node),
	2983: uint16(aux_sym_document_repeat1),
	2984: uint16(140),
	2985: uint16(15),
	2986: uint16(sym_doctype),
	2987: uint16(sym_element),
	2988: uint16(sym_script_element),
	2989: uint16(sym_style_element),
	2990: uint16(sym_erroneous_end_tag),
	2991: uint16(sym_if_statement),
	2992: uint16(sym_each_statement),
	2993: uint16(sym_await_statement),
	2994: uint16(sym_key_statement),
	2995: uint16(sym_snippet_statement),
	2996: uint16(sym_expression),
	2997: uint16(sym_html_tag),
	2998: uint16(sym_const_tag),
	2999: uint16(sym_debug_tag),
	3000: uint16(sym_render_tag),
	3001: uint16(17),
	3002: uint16(7),
	3003: uint16(1),
	3004: uint16(anon_sym_LT_BANG),
	3005: uint16(9),
	3006: uint16(1),
	3007: uint16(anon_sym_LT),
	3008: uint16(11),
	3009: uint16(1),
	3010: uint16(anon_sym_LT_SLASH),
	3011: uint16(110),
	3012: uint16(1),
	3013: uint16(anon_sym_LBRACE),
	3014: uint16(3),
	3015: uint16(1),
	3016: uint16(sym_if_start),
	3017: uint16(6),
	3018: uint16(1),
	3019: uint16(sym_await_start),
	3020: uint16(10),
	3021: uint16(1),
	3022: uint16(sym_each_start),
	3023: uint16(15),
	3024: uint16(1),
	3025: uint16(sym_start_tag),
	3026: uint16(20),
	3027: uint16(1),
	3028: uint16(sym_snippet_start),
	3029: uint16(33),
	3030: uint16(1),
	3031: uint16(sym_key_start),
	3032: uint16(108),
	3033: uint16(1),
	3034: uint16(sym_self_closing_tag),
	3035: uint16(209),
	3036: uint16(1),
	3037: uint16(sym_style_start_tag),
	3038: uint16(210),
	3039: uint16(1),
	3040: uint16(sym_script_start_tag),
	3041: uint16(3),
	3042: uint16(2),
	3043: uint16(sym_comment),
	3044: uint16(aux_sym__tag_value_token1),
	3045: uint16(13),
	3046: uint16(2),
	3047: uint16(sym_entity),
	3048: uint16(sym_text),
	3049: uint16(23),
	3050: uint16(2),
	3051: uint16(sym__node),
	3052: uint16(aux_sym_document_repeat1),
	3053: uint16(140),
	3054: uint16(15),
	3055: uint16(sym_doctype),
	3056: uint16(sym_element),
	3057: uint16(sym_script_element),
	3058: uint16(sym_style_element),
	3059: uint16(sym_erroneous_end_tag),
	3060: uint16(sym_if_statement),
	3061: uint16(sym_each_statement),
	3062: uint16(sym_await_statement),
	3063: uint16(sym_key_statement),
	3064: uint16(sym_snippet_statement),
	3065: uint16(sym_expression),
	3066: uint16(sym_html_tag),
	3067: uint16(sym_const_tag),
	3068: uint16(sym_debug_tag),
	3069: uint16(sym_render_tag),
	3070: uint16(17),
	3071: uint16(7),
	3072: uint16(1),
	3073: uint16(anon_sym_LT_BANG),
	3074: uint16(9),
	3075: uint16(1),
	3076: uint16(anon_sym_LT),
	3077: uint16(11),
	3078: uint16(1),
	3079: uint16(anon_sym_LT_SLASH),
	3080: uint16(113),
	3081: uint16(1),
	3082: uint16(anon_sym_LBRACE),
	3083: uint16(3),
	3084: uint16(1),
	3085: uint16(sym_if_start),
	3086: uint16(6),
	3087: uint16(1),
	3088: uint16(sym_await_start),
	3089: uint16(10),
	3090: uint16(1),
	3091: uint16(sym_each_start),
	3092: uint16(15),
	3093: uint16(1),
	3094: uint16(sym_start_tag),
	3095: uint16(20),
	3096: uint16(1),
	3097: uint16(sym_snippet_start),
	3098: uint16(33),
	3099: uint16(1),
	3100: uint16(sym_key_start),
	3101: uint16(108),
	3102: uint16(1),
	3103: uint16(sym_self_closing_tag),
	3104: uint16(209),
	3105: uint16(1),
	3106: uint16(sym_style_start_tag),
	3107: uint16(210),
	3108: uint16(1),
	3109: uint16(sym_script_start_tag),
	3110: uint16(3),
	3111: uint16(2),
	3112: uint16(sym_comment),
	3113: uint16(aux_sym__tag_value_token1),
	3114: uint16(13),
	3115: uint16(2),
	3116: uint16(sym_entity),
	3117: uint16(sym_text),
	3118: uint16(39),
	3119: uint16(2),
	3120: uint16(sym__node),
	3121: uint16(aux_sym_document_repeat1),
	3122: uint16(140),
	3123: uint16(15),
	3124: uint16(sym_doctype),
	3125: uint16(sym_element),
	3126: uint16(sym_script_element),
	3127: uint16(sym_style_element),
	3128: uint16(sym_erroneous_end_tag),
	3129: uint16(sym_if_statement),
	3130: uint16(sym_each_statement),
	3131: uint16(sym_await_statement),
	3132: uint16(sym_key_statement),
	3133: uint16(sym_snippet_statement),
	3134: uint16(sym_expression),
	3135: uint16(sym_html_tag),
	3136: uint16(sym_const_tag),
	3137: uint16(sym_debug_tag),
	3138: uint16(sym_render_tag),
	3139: uint16(17),
	3140: uint16(104),
	3141: uint16(1),
	3142: uint16(anon_sym_LBRACE),
	3143: uint16(116),
	3144: uint16(1),
	3145: uint16(anon_sym_LT_BANG),
	3146: uint16(119),
	3147: uint16(1),
	3148: uint16(anon_sym_LT),
	3149: uint16(122),
	3150: uint16(1),
	3151: uint16(anon_sym_LT_SLASH),
	3152: uint16(3),
	3153: uint16(1),
	3154: uint16(sym_if_start),
	3155: uint16(6),
	3156: uint16(1),
	3157: uint16(sym_await_start),
	3158: uint16(10),
	3159: uint16(1),
	3160: uint16(sym_each_start),
	3161: uint16(15),
	3162: uint16(1),
	3163: uint16(sym_start_tag),
	3164: uint16(20),
	3165: uint16(1),
	3166: uint16(sym_snippet_start),
	3167: uint16(33),
	3168: uint16(1),
	3169: uint16(sym_key_start),
	3170: uint16(108),
	3171: uint16(1),
	3172: uint16(sym_self_closing_tag),
	3173: uint16(209),
	3174: uint16(1),
	3175: uint16(sym_style_start_tag),
	3176: uint16(210),
	3177: uint16(1),
	3178: uint16(sym_script_start_tag),
	3179: uint16(3),
	3180: uint16(2),
	3181: uint16(sym_comment),
	3182: uint16(aux_sym__tag_value_token1),
	3183: uint16(125),
	3184: uint16(2),
	3185: uint16(sym_entity),
	3186: uint16(sym_text),
	3187: uint16(45),
	3188: uint16(2),
	3189: uint16(sym__node),
	3190: uint16(aux_sym_document_repeat1),
	3191: uint16(140),
	3192: uint16(15),
	3193: uint16(sym_doctype),
	3194: uint16(sym_element),
	3195: uint16(sym_script_element),
	3196: uint16(sym_style_element),
	3197: uint16(sym_erroneous_end_tag),
	3198: uint16(sym_if_statement),
	3199: uint16(sym_each_statement),
	3200: uint16(sym_await_statement),
	3201: uint16(sym_key_statement),
	3202: uint16(sym_snippet_statement),
	3203: uint16(sym_expression),
	3204: uint16(sym_html_tag),
	3205: uint16(sym_const_tag),
	3206: uint16(sym_debug_tag),
	3207: uint16(sym_render_tag),
	3208: uint16(17),
	3209: uint16(128),
	3210: uint16(1),
	3211: uint16(anon_sym_LT_BANG),
	3212: uint16(131),
	3213: uint16(1),
	3214: uint16(anon_sym_LT),
	3215: uint16(134),
	3216: uint16(1),
	3217: uint16(anon_sym_LT_SLASH),
	3218: uint16(140),
	3219: uint16(1),
	3220: uint16(anon_sym_LBRACE),
	3221: uint16(3),
	3222: uint16(1),
	3223: uint16(sym_if_start),
	3224: uint16(6),
	3225: uint16(1),
	3226: uint16(sym_await_start),
	3227: uint16(10),
	3228: uint16(1),
	3229: uint16(sym_each_start),
	3230: uint16(15),
	3231: uint16(1),
	3232: uint16(sym_start_tag),
	3233: uint16(20),
	3234: uint16(1),
	3235: uint16(sym_snippet_start),
	3236: uint16(33),
	3237: uint16(1),
	3238: uint16(sym_key_start),
	3239: uint16(108),
	3240: uint16(1),
	3241: uint16(sym_self_closing_tag),
	3242: uint16(209),
	3243: uint16(1),
	3244: uint16(sym_style_start_tag),
	3245: uint16(210),
	3246: uint16(1),
	3247: uint16(sym_script_start_tag),
	3248: uint16(3),
	3249: uint16(2),
	3250: uint16(sym_comment),
	3251: uint16(aux_sym__tag_value_token1),
	3252: uint16(137),
	3253: uint16(2),
	3254: uint16(sym_entity),
	3255: uint16(sym_text),
	3256: uint16(23),
	3257: uint16(2),
	3258: uint16(sym__node),
	3259: uint16(aux_sym_document_repeat1),
	3260: uint16(140),
	3261: uint16(15),
	3262: uint16(sym_doctype),
	3263: uint16(sym_element),
	3264: uint16(sym_script_element),
	3265: uint16(sym_style_element),
	3266: uint16(sym_erroneous_end_tag),
	3267: uint16(sym_if_statement),
	3268: uint16(sym_each_statement),
	3269: uint16(sym_await_statement),
	3270: uint16(sym_key_statement),
	3271: uint16(sym_snippet_statement),
	3272: uint16(sym_expression),
	3273: uint16(sym_html_tag),
	3274: uint16(sym_const_tag),
	3275: uint16(sym_debug_tag),
	3276: uint16(sym_render_tag),
	3277: uint16(17),
	3278: uint16(7),
	3279: uint16(1),
	3280: uint16(anon_sym_LT_BANG),
	3281: uint16(9),
	3282: uint16(1),
	3283: uint16(anon_sym_LT),
	3284: uint16(11),
	3285: uint16(1),
	3286: uint16(anon_sym_LT_SLASH),
	3287: uint16(140),
	3288: uint16(1),
	3289: uint16(anon_sym_LBRACE),
	3290: uint16(3),
	3291: uint16(1),
	3292: uint16(sym_if_start),
	3293: uint16(6),
	3294: uint16(1),
	3295: uint16(sym_await_start),
	3296: uint16(10),
	3297: uint16(1),
	3298: uint16(sym_each_start),
	3299: uint16(15),
	3300: uint16(1),
	3301: uint16(sym_start_tag),
	3302: uint16(20),
	3303: uint16(1),
	3304: uint16(sym_snippet_start),
	3305: uint16(33),
	3306: uint16(1),
	3307: uint16(sym_key_start),
	3308: uint16(108),
	3309: uint16(1),
	3310: uint16(sym_self_closing_tag),
	3311: uint16(209),
	3312: uint16(1),
	3313: uint16(sym_style_start_tag),
	3314: uint16(210),
	3315: uint16(1),
	3316: uint16(sym_script_start_tag),
	3317: uint16(3),
	3318: uint16(2),
	3319: uint16(sym_comment),
	3320: uint16(aux_sym__tag_value_token1),
	3321: uint16(13),
	3322: uint16(2),
	3323: uint16(sym_entity),
	3324: uint16(sym_text),
	3325: uint16(23),
	3326: uint16(2),
	3327: uint16(sym__node),
	3328: uint16(aux_sym_document_repeat1),
	3329: uint16(140),
	3330: uint16(15),
	3331: uint16(sym_doctype),
	3332: uint16(sym_element),
	3333: uint16(sym_script_element),
	3334: uint16(sym_style_element),
	3335: uint16(sym_erroneous_end_tag),
	3336: uint16(sym_if_statement),
	3337: uint16(sym_each_statement),
	3338: uint16(sym_await_statement),
	3339: uint16(sym_key_statement),
	3340: uint16(sym_snippet_statement),
	3341: uint16(sym_expression),
	3342: uint16(sym_html_tag),
	3343: uint16(sym_const_tag),
	3344: uint16(sym_debug_tag),
	3345: uint16(sym_render_tag),
	3346: uint16(18),
	3347: uint16(143),
	3348: uint16(1),
	3349: uint16(anon_sym_POUND),
	3350: uint16(145),
	3351: uint16(1),
	3352: uint16(anon_sym_COLON),
	3353: uint16(147),
	3354: uint16(1),
	3355: uint16(anon_sym_SLASH),
	3356: uint16(149),
	3357: uint16(1),
	3358: uint16(anon_sym_AT),
	3359: uint16(151),
	3360: uint16(1),
	3361: uint16(sym_svelte_raw_text),
	3362: uint16(222),
	3363: uint16(1),
	3364: uint16(sym__each_start_tag),
	3365: uint16(236),
	3366: uint16(1),
	3367: uint16(sym__then_tag),
	3368: uint16(239),
	3369: uint16(1),
	3370: uint16(sym__catch_tag),
	3371: uint16(253),
	3372: uint16(1),
	3373: uint16(sym__render_tag),
	3374: uint16(271),
	3375: uint16(1),
	3376: uint16(sym__if_start_tag),
	3377: uint16(295),
	3378: uint16(1),
	3379: uint16(sym__await_start_tag),
	3380: uint16(296),
	3381: uint16(1),
	3382: uint16(sym__key_start_tag),
	3383: uint16(313),
	3384: uint16(1),
	3385: uint16(sym__snippet_start_tag),
	3386: uint16(315),
	3387: uint16(1),
	3388: uint16(sym__html_tag),
	3389: uint16(321),
	3390: uint16(1),
	3391: uint16(sym__const_tag),
	3392: uint16(333),
	3393: uint16(1),
	3394: uint16(sym__await_end_tag),
	3395: uint16(335),
	3396: uint16(1),
	3397: uint16(sym__debug_tag),
	3398: uint16(3),
	3399: uint16(2),
	3400: uint16(sym_comment),
	3401: uint16(aux_sym__tag_value_token1),
	3402: uint16(18),
	3403: uint16(143),
	3404: uint16(1),
	3405: uint16(anon_sym_POUND),
	3406: uint16(149),
	3407: uint16(1),
	3408: uint16(anon_sym_AT),
	3409: uint16(151),
	3410: uint16(1),
	3411: uint16(sym_svelte_raw_text),
	3412: uint16(153),
	3413: uint16(1),
	3414: uint16(anon_sym_COLON),
	3415: uint16(155),
	3416: uint16(1),
	3417: uint16(anon_sym_SLASH),
	3418: uint16(222),
	3419: uint16(1),
	3420: uint16(sym__each_start_tag),
	3421: uint16(253),
	3422: uint16(1),
	3423: uint16(sym__render_tag),
	3424: uint16(271),
	3425: uint16(1),
	3426: uint16(sym__if_start_tag),
	3427: uint16(284),
	3428: uint16(1),
	3429: uint16(sym__else_if_tag),
	3430: uint16(285),
	3431: uint16(1),
	3432: uint16(sym__else_tag),
	3433: uint16(286),
	3434: uint16(1),
	3435: uint16(sym__if_end_tag),
	3436: uint16(295),
	3437: uint16(1),
	3438: uint16(sym__await_start_tag),
	3439: uint16(296),
	3440: uint16(1),
	3441: uint16(sym__key_start_tag),
	3442: uint16(313),
	3443: uint16(1),
	3444: uint16(sym__snippet_start_tag),
	3445: uint16(315),
	3446: uint16(1),
	3447: uint16(sym__html_tag),
	3448: uint16(321),
	3449: uint16(1),
	3450: uint16(sym__const_tag),
	3451: uint16(335),
	3452: uint16(1),
	3453: uint16(sym__debug_tag),
	3454: uint16(3),
	3455: uint16(2),
	3456: uint16(sym_comment),
	3457: uint16(aux_sym__tag_value_token1),
	3458: uint16(18),
	3459: uint16(143),
	3460: uint16(1),
	3461: uint16(anon_sym_POUND),
	3462: uint16(149),
	3463: uint16(1),
	3464: uint16(anon_sym_AT),
	3465: uint16(151),
	3466: uint16(1),
	3467: uint16(sym_svelte_raw_text),
	3468: uint16(153),
	3469: uint16(1),
	3470: uint16(anon_sym_COLON),
	3471: uint16(155),
	3472: uint16(1),
	3473: uint16(anon_sym_SLASH),
	3474: uint16(222),
	3475: uint16(1),
	3476: uint16(sym__each_start_tag),
	3477: uint16(253),
	3478: uint16(1),
	3479: uint16(sym__render_tag),
	3480: uint16(271),
	3481: uint16(1),
	3482: uint16(sym__if_start_tag),
	3483: uint16(284),
	3484: uint16(1),
	3485: uint16(sym__else_if_tag),
	3486: uint16(285),
	3487: uint16(1),
	3488: uint16(sym__else_tag),
	3489: uint16(295),
	3490: uint16(1),
	3491: uint16(sym__await_start_tag),
	3492: uint16(296),
	3493: uint16(1),
	3494: uint16(sym__key_start_tag),
	3495: uint16(313),
	3496: uint16(1),
	3497: uint16(sym__snippet_start_tag),
	3498: uint16(315),
	3499: uint16(1),
	3500: uint16(sym__html_tag),
	3501: uint16(321),
	3502: uint16(1),
	3503: uint16(sym__const_tag),
	3504: uint16(335),
	3505: uint16(1),
	3506: uint16(sym__debug_tag),
	3507: uint16(340),
	3508: uint16(1),
	3509: uint16(sym__if_end_tag),
	3510: uint16(3),
	3511: uint16(2),
	3512: uint16(sym_comment),
	3513: uint16(aux_sym__tag_value_token1),
	3514: uint16(18),
	3515: uint16(143),
	3516: uint16(1),
	3517: uint16(anon_sym_POUND),
	3518: uint16(145),
	3519: uint16(1),
	3520: uint16(anon_sym_COLON),
	3521: uint16(147),
	3522: uint16(1),
	3523: uint16(anon_sym_SLASH),
	3524: uint16(149),
	3525: uint16(1),
	3526: uint16(anon_sym_AT),
	3527: uint16(151),
	3528: uint16(1),
	3529: uint16(sym_svelte_raw_text),
	3530: uint16(222),
	3531: uint16(1),
	3532: uint16(sym__each_start_tag),
	3533: uint16(236),
	3534: uint16(1),
	3535: uint16(sym__then_tag),
	3536: uint16(239),
	3537: uint16(1),
	3538: uint16(sym__catch_tag),
	3539: uint16(253),
	3540: uint16(1),
	3541: uint16(sym__render_tag),
	3542: uint16(271),
	3543: uint16(1),
	3544: uint16(sym__if_start_tag),
	3545: uint16(295),
	3546: uint16(1),
	3547: uint16(sym__await_start_tag),
	3548: uint16(296),
	3549: uint16(1),
	3550: uint16(sym__key_start_tag),
	3551: uint16(303),
	3552: uint16(1),
	3553: uint16(sym__await_end_tag),
	3554: uint16(313),
	3555: uint16(1),
	3556: uint16(sym__snippet_start_tag),
	3557: uint16(315),
	3558: uint16(1),
	3559: uint16(sym__html_tag),
	3560: uint16(321),
	3561: uint16(1),
	3562: uint16(sym__const_tag),
	3563: uint16(335),
	3564: uint16(1),
	3565: uint16(sym__debug_tag),
	3566: uint16(3),
	3567: uint16(2),
	3568: uint16(sym_comment),
	3569: uint16(aux_sym__tag_value_token1),
	3570: uint16(17),
	3571: uint16(143),
	3572: uint16(1),
	3573: uint16(anon_sym_POUND),
	3574: uint16(149),
	3575: uint16(1),
	3576: uint16(anon_sym_AT),
	3577: uint16(151),
	3578: uint16(1),
	3579: uint16(sym_svelte_raw_text),
	3580: uint16(157),
	3581: uint16(1),
	3582: uint16(anon_sym_COLON),
	3583: uint16(159),
	3584: uint16(1),
	3585: uint16(anon_sym_SLASH),
	3586: uint16(222),
	3587: uint16(1),
	3588: uint16(sym__each_start_tag),
	3589: uint16(253),
	3590: uint16(1),
	3591: uint16(sym__render_tag),
	3592: uint16(271),
	3593: uint16(1),
	3594: uint16(sym__if_start_tag),
	3595: uint16(285),
	3596: uint16(1),
	3597: uint16(sym__else_tag),
	3598: uint16(295),
	3599: uint16(1),
	3600: uint16(sym__await_start_tag),
	3601: uint16(296),
	3602: uint16(1),
	3603: uint16(sym__key_start_tag),
	3604: uint16(312),
	3605: uint16(1),
	3606: uint16(sym__each_end_tag),
	3607: uint16(313),
	3608: uint16(1),
	3609: uint16(sym__snippet_start_tag),
	3610: uint16(315),
	3611: uint16(1),
	3612: uint16(sym__html_tag),
	3613: uint16(321),
	3614: uint16(1),
	3615: uint16(sym__const_tag),
	3616: uint16(335),
	3617: uint16(1),
	3618: uint16(sym__debug_tag),
	3619: uint16(3),
	3620: uint16(2),
	3621: uint16(sym_comment),
	3622: uint16(aux_sym__tag_value_token1),
	3623: uint16(17),
	3624: uint16(143),
	3625: uint16(1),
	3626: uint16(anon_sym_POUND),
	3627: uint16(149),
	3628: uint16(1),
	3629: uint16(anon_sym_AT),
	3630: uint16(151),
	3631: uint16(1),
	3632: uint16(sym_svelte_raw_text),
	3633: uint16(157),
	3634: uint16(1),
	3635: uint16(anon_sym_COLON),
	3636: uint16(159),
	3637: uint16(1),
	3638: uint16(anon_sym_SLASH),
	3639: uint16(222),
	3640: uint16(1),
	3641: uint16(sym__each_start_tag),
	3642: uint16(253),
	3643: uint16(1),
	3644: uint16(sym__render_tag),
	3645: uint16(271),
	3646: uint16(1),
	3647: uint16(sym__if_start_tag),
	3648: uint16(285),
	3649: uint16(1),
	3650: uint16(sym__else_tag),
	3651: uint16(295),
	3652: uint16(1),
	3653: uint16(sym__await_start_tag),
	3654: uint16(296),
	3655: uint16(1),
	3656: uint16(sym__key_start_tag),
	3657: uint16(299),
	3658: uint16(1),
	3659: uint16(sym__each_end_tag),
	3660: uint16(313),
	3661: uint16(1),
	3662: uint16(sym__snippet_start_tag),
	3663: uint16(315),
	3664: uint16(1),
	3665: uint16(sym__html_tag),
	3666: uint16(321),
	3667: uint16(1),
	3668: uint16(sym__const_tag),
	3669: uint16(335),
	3670: uint16(1),
	3671: uint16(sym__debug_tag),
	3672: uint16(3),
	3673: uint16(2),
	3674: uint16(sym_comment),
	3675: uint16(aux_sym__tag_value_token1),
	3676: uint16(15),
	3677: uint16(143),
	3678: uint16(1),
	3679: uint16(anon_sym_POUND),
	3680: uint16(149),
	3681: uint16(1),
	3682: uint16(anon_sym_AT),
	3683: uint16(151),
	3684: uint16(1),
	3685: uint16(sym_svelte_raw_text),
	3686: uint16(159),
	3687: uint16(1),
	3688: uint16(anon_sym_SLASH),
	3689: uint16(222),
	3690: uint16(1),
	3691: uint16(sym__each_start_tag),
	3692: uint16(253),
	3693: uint16(1),
	3694: uint16(sym__render_tag),
	3695: uint16(271),
	3696: uint16(1),
	3697: uint16(sym__if_start_tag),
	3698: uint16(295),
	3699: uint16(1),
	3700: uint16(sym__await_start_tag),
	3701: uint16(296),
	3702: uint16(1),
	3703: uint16(sym__key_start_tag),
	3704: uint16(312),
	3705: uint16(1),
	3706: uint16(sym__each_end_tag),
	3707: uint16(313),
	3708: uint16(1),
	3709: uint16(sym__snippet_start_tag),
	3710: uint16(315),
	3711: uint16(1),
	3712: uint16(sym__html_tag),
	3713: uint16(321),
	3714: uint16(1),
	3715: uint16(sym__const_tag),
	3716: uint16(335),
	3717: uint16(1),
	3718: uint16(sym__debug_tag),
	3719: uint16(3),
	3720: uint16(2),
	3721: uint16(sym_comment),
	3722: uint16(aux_sym__tag_value_token1),
	3723: uint16(15),
	3724: uint16(143),
	3725: uint16(1),
	3726: uint16(anon_sym_POUND),
	3727: uint16(149),
	3728: uint16(1),
	3729: uint16(anon_sym_AT),
	3730: uint16(151),
	3731: uint16(1),
	3732: uint16(sym_svelte_raw_text),
	3733: uint16(161),
	3734: uint16(1),
	3735: uint16(anon_sym_SLASH),
	3736: uint16(222),
	3737: uint16(1),
	3738: uint16(sym__each_start_tag),
	3739: uint16(253),
	3740: uint16(1),
	3741: uint16(sym__render_tag),
	3742: uint16(268),
	3743: uint16(1),
	3744: uint16(sym__snippet_end_tag),
	3745: uint16(271),
	3746: uint16(1),
	3747: uint16(sym__if_start_tag),
	3748: uint16(295),
	3749: uint16(1),
	3750: uint16(sym__await_start_tag),
	3751: uint16(296),
	3752: uint16(1),
	3753: uint16(sym__key_start_tag),
	3754: uint16(313),
	3755: uint16(1),
	3756: uint16(sym__snippet_start_tag),
	3757: uint16(315),
	3758: uint16(1),
	3759: uint16(sym__html_tag),
	3760: uint16(321),
	3761: uint16(1),
	3762: uint16(sym__const_tag),
	3763: uint16(335),
	3764: uint16(1),
	3765: uint16(sym__debug_tag),
	3766: uint16(3),
	3767: uint16(2),
	3768: uint16(sym_comment),
	3769: uint16(aux_sym__tag_value_token1),
	3770: uint16(15),
	3771: uint16(143),
	3772: uint16(1),
	3773: uint16(anon_sym_POUND),
	3774: uint16(149),
	3775: uint16(1),
	3776: uint16(anon_sym_AT),
	3777: uint16(151),
	3778: uint16(1),
	3779: uint16(sym_svelte_raw_text),
	3780: uint16(163),
	3781: uint16(1),
	3782: uint16(anon_sym_SLASH),
	3783: uint16(222),
	3784: uint16(1),
	3785: uint16(sym__each_start_tag),
	3786: uint16(253),
	3787: uint16(1),
	3788: uint16(sym__render_tag),
	3789: uint16(258),
	3790: uint16(1),
	3791: uint16(sym__key_end_tag),
	3792: uint16(271),
	3793: uint16(1),
	3794: uint16(sym__if_start_tag),
	3795: uint16(295),
	3796: uint16(1),
	3797: uint16(sym__await_start_tag),
	3798: uint16(296),
	3799: uint16(1),
	3800: uint16(sym__key_start_tag),
	3801: uint16(313),
	3802: uint16(1),
	3803: uint16(sym__snippet_start_tag),
	3804: uint16(315),
	3805: uint16(1),
	3806: uint16(sym__html_tag),
	3807: uint16(321),
	3808: uint16(1),
	3809: uint16(sym__const_tag),
	3810: uint16(335),
	3811: uint16(1),
	3812: uint16(sym__debug_tag),
	3813: uint16(3),
	3814: uint16(2),
	3815: uint16(sym_comment),
	3816: uint16(aux_sym__tag_value_token1),
	3817: uint16(15),
	3818: uint16(143),
	3819: uint16(1),
	3820: uint16(anon_sym_POUND),
	3821: uint16(149),
	3822: uint16(1),
	3823: uint16(anon_sym_AT),
	3824: uint16(151),
	3825: uint16(1),
	3826: uint16(sym_svelte_raw_text),
	3827: uint16(159),
	3828: uint16(1),
	3829: uint16(anon_sym_SLASH),
	3830: uint16(222),
	3831: uint16(1),
	3832: uint16(sym__each_start_tag),
	3833: uint16(253),
	3834: uint16(1),
	3835: uint16(sym__render_tag),
	3836: uint16(271),
	3837: uint16(1),
	3838: uint16(sym__if_start_tag),
	3839: uint16(295),
	3840: uint16(1),
	3841: uint16(sym__await_start_tag),
	3842: uint16(296),
	3843: uint16(1),
	3844: uint16(sym__key_start_tag),
	3845: uint16(299),
	3846: uint16(1),
	3847: uint16(sym__each_end_tag),
	3848: uint16(313),
	3849: uint16(1),
	3850: uint16(sym__snippet_start_tag),
	3851: uint16(315),
	3852: uint16(1),
	3853: uint16(sym__html_tag),
	3854: uint16(321),
	3855: uint16(1),
	3856: uint16(sym__const_tag),
	3857: uint16(335),
	3858: uint16(1),
	3859: uint16(sym__debug_tag),
	3860: uint16(3),
	3861: uint16(2),
	3862: uint16(sym_comment),
	3863: uint16(aux_sym__tag_value_token1),
	3864: uint16(15),
	3865: uint16(143),
	3866: uint16(1),
	3867: uint16(anon_sym_POUND),
	3868: uint16(149),
	3869: uint16(1),
	3870: uint16(anon_sym_AT),
	3871: uint16(151),
	3872: uint16(1),
	3873: uint16(sym_svelte_raw_text),
	3874: uint16(163),
	3875: uint16(1),
	3876: uint16(anon_sym_SLASH),
	3877: uint16(222),
	3878: uint16(1),
	3879: uint16(sym__each_start_tag),
	3880: uint16(253),
	3881: uint16(1),
	3882: uint16(sym__render_tag),
	3883: uint16(271),
	3884: uint16(1),
	3885: uint16(sym__if_start_tag),
	3886: uint16(295),
	3887: uint16(1),
	3888: uint16(sym__await_start_tag),
	3889: uint16(296),
	3890: uint16(1),
	3891: uint16(sym__key_start_tag),
	3892: uint16(307),
	3893: uint16(1),
	3894: uint16(sym__key_end_tag),
	3895: uint16(313),
	3896: uint16(1),
	3897: uint16(sym__snippet_start_tag),
	3898: uint16(315),
	3899: uint16(1),
	3900: uint16(sym__html_tag),
	3901: uint16(321),
	3902: uint16(1),
	3903: uint16(sym__const_tag),
	3904: uint16(335),
	3905: uint16(1),
	3906: uint16(sym__debug_tag),
	3907: uint16(3),
	3908: uint16(2),
	3909: uint16(sym_comment),
	3910: uint16(aux_sym__tag_value_token1),
	3911: uint16(15),
	3912: uint16(143),
	3913: uint16(1),
	3914: uint16(anon_sym_POUND),
	3915: uint16(149),
	3916: uint16(1),
	3917: uint16(anon_sym_AT),
	3918: uint16(151),
	3919: uint16(1),
	3920: uint16(sym_svelte_raw_text),
	3921: uint16(161),
	3922: uint16(1),
	3923: uint16(anon_sym_SLASH),
	3924: uint16(222),
	3925: uint16(1),
	3926: uint16(sym__each_start_tag),
	3927: uint16(253),
	3928: uint16(1),
	3929: uint16(sym__render_tag),
	3930: uint16(271),
	3931: uint16(1),
	3932: uint16(sym__if_start_tag),
	3933: uint16(295),
	3934: uint16(1),
	3935: uint16(sym__await_start_tag),
	3936: uint16(296),
	3937: uint16(1),
	3938: uint16(sym__key_start_tag),
	3939: uint16(308),
	3940: uint16(1),
	3941: uint16(sym__snippet_end_tag),
	3942: uint16(313),
	3943: uint16(1),
	3944: uint16(sym__snippet_start_tag),
	3945: uint16(315),
	3946: uint16(1),
	3947: uint16(sym__html_tag),
	3948: uint16(321),
	3949: uint16(1),
	3950: uint16(sym__const_tag),
	3951: uint16(335),
	3952: uint16(1),
	3953: uint16(sym__debug_tag),
	3954: uint16(3),
	3955: uint16(2),
	3956: uint16(sym_comment),
	3957: uint16(aux_sym__tag_value_token1),
	3958: uint16(13),
	3959: uint16(143),
	3960: uint16(1),
	3961: uint16(anon_sym_POUND),
	3962: uint16(149),
	3963: uint16(1),
	3964: uint16(anon_sym_AT),
	3965: uint16(151),
	3966: uint16(1),
	3967: uint16(sym_svelte_raw_text),
	3968: uint16(222),
	3969: uint16(1),
	3970: uint16(sym__each_start_tag),
	3971: uint16(253),
	3972: uint16(1),
	3973: uint16(sym__render_tag),
	3974: uint16(271),
	3975: uint16(1),
	3976: uint16(sym__if_start_tag),
	3977: uint16(295),
	3978: uint16(1),
	3979: uint16(sym__await_start_tag),
	3980: uint16(296),
	3981: uint16(1),
	3982: uint16(sym__key_start_tag),
	3983: uint16(313),
	3984: uint16(1),
	3985: uint16(sym__snippet_start_tag),
	3986: uint16(315),
	3987: uint16(1),
	3988: uint16(sym__html_tag),
	3989: uint16(321),
	3990: uint16(1),
	3991: uint16(sym__const_tag),
	3992: uint16(335),
	3993: uint16(1),
	3994: uint16(sym__debug_tag),
	3995: uint16(3),
	3996: uint16(2),
	3997: uint16(sym_comment),
	3998: uint16(aux_sym__tag_value_token1),
	3999: uint16(13),
	4000: uint16(143),
	4001: uint16(1),
	4002: uint16(anon_sym_POUND),
	4003: uint16(149),
	4004: uint16(1),
	4005: uint16(anon_sym_AT),
	4006: uint16(165),
	4007: uint16(1),
	4008: uint16(sym_svelte_raw_text),
	4009: uint16(222),
	4010: uint16(1),
	4011: uint16(sym__each_start_tag),
	4012: uint16(271),
	4013: uint16(1),
	4014: uint16(sym__if_start_tag),
	4015: uint16(295),
	4016: uint16(1),
	4017: uint16(sym__await_start_tag),
	4018: uint16(296),
	4019: uint16(1),
	4020: uint16(sym__key_start_tag),
	4021: uint16(313),
	4022: uint16(1),
	4023: uint16(sym__snippet_start_tag),
	4024: uint16(325),
	4025: uint16(1),
	4026: uint16(sym__html_tag),
	4027: uint16(326),
	4028: uint16(1),
	4029: uint16(sym__const_tag),
	4030: uint16(327),
	4031: uint16(1),
	4032: uint16(sym__debug_tag),
	4033: uint16(328),
	4034: uint16(1),
	4035: uint16(sym__render_tag),
	4036: uint16(3),
	4037: uint16(2),
	4038: uint16(sym_comment),
	4039: uint16(aux_sym__tag_value_token1),
	4040: uint16(3),
	4041: uint16(169),
	4042: uint16(1),
	4043: uint16(anon_sym_LT),
	4044: uint16(3),
	4045: uint16(2),
	4046: uint16(sym_comment),
	4047: uint16(aux_sym__tag_value_token1),
	4048: uint16(167),
	4049: uint16(6),
	4051: uint16(anon_sym_LT_BANG),
	4052: uint16(anon_sym_LT_SLASH),
	4053: uint16(sym_entity),
	4054: uint16(sym_text),
	4055: uint16(anon_sym_LBRACE),
	4056: uint16(3),
	4057: uint16(173),
	4058: uint16(1),
	4059: uint16(anon_sym_LT),
	4060: uint16(3),
	4061: uint16(2),
	4062: uint16(sym_comment),
	4063: uint16(aux_sym__tag_value_token1),
	4064: uint16(171),
	4065: uint16(6),
	4067: uint16(anon_sym_LT_BANG),
	4068: uint16(anon_sym_LT_SLASH),
	4069: uint16(sym_entity),
	4070: uint16(sym_text),
	4071: uint16(anon_sym_LBRACE),
	4072: uint16(3),
	4073: uint16(177),
	4074: uint16(1),
	4075: uint16(anon_sym_LT),
	4076: uint16(3),
	4077: uint16(2),
	4078: uint16(sym_comment),
	4079: uint16(aux_sym__tag_value_token1),
	4080: uint16(175),
	4081: uint16(6),
	4082: uint16(sym__implicit_end_tag),
	4083: uint16(anon_sym_LT_BANG),
	4084: uint16(anon_sym_LT_SLASH),
	4085: uint16(sym_entity),
	4086: uint16(sym_text),
	4087: uint16(anon_sym_LBRACE),
	4088: uint16(3),
	4089: uint16(181),
	4090: uint16(1),
	4091: uint16(anon_sym_LT),
	4092: uint16(3),
	4093: uint16(2),
	4094: uint16(sym_comment),
	4095: uint16(aux_sym__tag_value_token1),
	4096: uint16(179),
	4097: uint16(6),
	4099: uint16(anon_sym_LT_BANG),
	4100: uint16(anon_sym_LT_SLASH),
	4101: uint16(sym_entity),
	4102: uint16(sym_text),
	4103: uint16(anon_sym_LBRACE),
	4104: uint16(7),
	4105: uint16(183),
	4106: uint16(1),
	4107: uint16(anon_sym_GT),
	4108: uint16(185),
	4109: uint16(1),
	4110: uint16(anon_sym_SLASH_GT),
	4111: uint16(187),
	4112: uint16(1),
	4113: uint16(sym_attribute_name),
	4114: uint16(189),
	4115: uint16(1),
	4116: uint16(anon_sym_LBRACE),
	4117: uint16(196),
	4118: uint16(1),
	4119: uint16(sym_expression),
	4120: uint16(3),
	4121: uint16(2),
	4122: uint16(sym_comment),
	4123: uint16(aux_sym__tag_value_token1),
	4124: uint16(81),
	4125: uint16(2),
	4126: uint16(sym_attribute),
	4127: uint16(aux_sym_start_tag_repeat1),
	4128: uint16(3),
	4129: uint16(193),
	4130: uint16(1),
	4131: uint16(anon_sym_LT),
	4132: uint16(3),
	4133: uint16(2),
	4134: uint16(sym_comment),
	4135: uint16(aux_sym__tag_value_token1),
	4136: uint16(191),
	4137: uint16(6),
	4139: uint16(anon_sym_LT_BANG),
	4140: uint16(anon_sym_LT_SLASH),
	4141: uint16(sym_entity),
	4142: uint16(sym_text),
	4143: uint16(anon_sym_LBRACE),
	4144: uint16(3),
	4145: uint16(197),
	4146: uint16(1),
	4147: uint16(anon_sym_LT),
	4148: uint16(3),
	4149: uint16(2),
	4150: uint16(sym_comment),
	4151: uint16(aux_sym__tag_value_token1),
	4152: uint16(195),
	4153: uint16(6),
	4155: uint16(anon_sym_LT_BANG),
	4156: uint16(anon_sym_LT_SLASH),
	4157: uint16(sym_entity),
	4158: uint16(sym_text),
	4159: uint16(anon_sym_LBRACE),
	4160: uint16(3),
	4161: uint16(201),
	4162: uint16(1),
	4163: uint16(anon_sym_LT),
	4164: uint16(3),
	4165: uint16(2),
	4166: uint16(sym_comment),
	4167: uint16(aux_sym__tag_value_token1),
	4168: uint16(199),
	4169: uint16(6),
	4171: uint16(anon_sym_LT_BANG),
	4172: uint16(anon_sym_LT_SLASH),
	4173: uint16(sym_entity),
	4174: uint16(sym_text),
	4175: uint16(anon_sym_LBRACE),
	4176: uint16(3),
	4177: uint16(205),
	4178: uint16(1),
	4179: uint16(anon_sym_LT),
	4180: uint16(3),
	4181: uint16(2),
	4182: uint16(sym_comment),
	4183: uint16(aux_sym__tag_value_token1),
	4184: uint16(203),
	4185: uint16(6),
	4187: uint16(anon_sym_LT_BANG),
	4188: uint16(anon_sym_LT_SLASH),
	4189: uint16(sym_entity),
	4190: uint16(sym_text),
	4191: uint16(anon_sym_LBRACE),
	4192: uint16(3),
	4193: uint16(209),
	4194: uint16(1),
	4195: uint16(anon_sym_LT),
	4196: uint16(3),
	4197: uint16(2),
	4198: uint16(sym_comment),
	4199: uint16(aux_sym__tag_value_token1),
	4200: uint16(207),
	4201: uint16(6),
	4203: uint16(anon_sym_LT_BANG),
	4204: uint16(anon_sym_LT_SLASH),
	4205: uint16(sym_entity),
	4206: uint16(sym_text),
	4207: uint16(anon_sym_LBRACE),
	4208: uint16(3),
	4209: uint16(213),
	4210: uint16(1),
	4211: uint16(anon_sym_LT),
	4212: uint16(3),
	4213: uint16(2),
	4214: uint16(sym_comment),
	4215: uint16(aux_sym__tag_value_token1),
	4216: uint16(211),
	4217: uint16(6),
	4219: uint16(anon_sym_LT_BANG),
	4220: uint16(anon_sym_LT_SLASH),
	4221: uint16(sym_entity),
	4222: uint16(sym_text),
	4223: uint16(anon_sym_LBRACE),
	4224: uint16(3),
	4225: uint16(217),
	4226: uint16(1),
	4227: uint16(anon_sym_LT),
	4228: uint16(3),
	4229: uint16(2),
	4230: uint16(sym_comment),
	4231: uint16(aux_sym__tag_value_token1),
	4232: uint16(215),
	4233: uint16(6),
	4235: uint16(anon_sym_LT_BANG),
	4236: uint16(anon_sym_LT_SLASH),
	4237: uint16(sym_entity),
	4238: uint16(sym_text),
	4239: uint16(anon_sym_LBRACE),
	4240: uint16(7),
	4241: uint16(219),
	4242: uint16(1),
	4243: uint16(anon_sym_LBRACE),
	4244: uint16(37),
	4245: uint16(1),
	4246: uint16(sym_else_if_start),
	4247: uint16(40),
	4248: uint16(1),
	4249: uint16(sym_else_start),
	4250: uint16(88),
	4251: uint16(1),
	4252: uint16(sym_if_end),
	4253: uint16(248),
	4254: uint16(1),
	4255: uint16(sym_else_block),
	4256: uint16(3),
	4257: uint16(2),
	4258: uint16(sym_comment),
	4259: uint16(aux_sym__tag_value_token1),
	4260: uint16(187),
	4261: uint16(2),
	4262: uint16(sym_else_if_block),
	4263: uint16(aux_sym_if_statement_repeat1),
	4264: uint16(3),
	4265: uint16(223),
	4266: uint16(1),
	4267: uint16(anon_sym_LT),
	4268: uint16(3),
	4269: uint16(2),
	4270: uint16(sym_comment),
	4271: uint16(aux_sym__tag_value_token1),
	4272: uint16(221),
	4273: uint16(6),
	4275: uint16(anon_sym_LT_BANG),
	4276: uint16(anon_sym_LT_SLASH),
	4277: uint16(sym_entity),
	4278: uint16(sym_text),
	4279: uint16(anon_sym_LBRACE),
	4280: uint16(3),
	4281: uint16(227),
	4282: uint16(1),
	4283: uint16(anon_sym_LT),
	4284: uint16(3),
	4285: uint16(2),
	4286: uint16(sym_comment),
	4287: uint16(aux_sym__tag_value_token1),
	4288: uint16(225),
	4289: uint16(6),
	4291: uint16(anon_sym_LT_BANG),
	4292: uint16(anon_sym_LT_SLASH),
	4293: uint16(sym_entity),
	4294: uint16(sym_text),
	4295: uint16(anon_sym_LBRACE),
	4296: uint16(3),
	4297: uint16(231),
	4298: uint16(1),
	4299: uint16(anon_sym_LT),
	4300: uint16(3),
	4301: uint16(2),
	4302: uint16(sym_comment),
	4303: uint16(aux_sym__tag_value_token1),
	4304: uint16(229),
	4305: uint16(6),
	4307: uint16(anon_sym_LT_BANG),
	4308: uint16(anon_sym_LT_SLASH),
	4309: uint16(sym_entity),
	4310: uint16(sym_text),
	4311: uint16(anon_sym_LBRACE),
	4312: uint16(3),
	4313: uint16(235),
	4314: uint16(1),
	4315: uint16(anon_sym_LT),
	4316: uint16(3),
	4317: uint16(2),
	4318: uint16(sym_comment),
	4319: uint16(aux_sym__tag_value_token1),
	4320: uint16(233),
	4321: uint16(6),
	4323: uint16(anon_sym_LT_BANG),
	4324: uint16(anon_sym_LT_SLASH),
	4325: uint16(sym_entity),
	4326: uint16(sym_text),
	4327: uint16(anon_sym_LBRACE),
	4328: uint16(3),
	4329: uint16(239),
	4330: uint16(1),
	4331: uint16(anon_sym_LT),
	4332: uint16(3),
	4333: uint16(2),
	4334: uint16(sym_comment),
	4335: uint16(aux_sym__tag_value_token1),
	4336: uint16(237),
	4337: uint16(6),
	4339: uint16(anon_sym_LT_BANG),
	4340: uint16(anon_sym_LT_SLASH),
	4341: uint16(sym_entity),
	4342: uint16(sym_text),
	4343: uint16(anon_sym_LBRACE),
	4344: uint16(3),
	4345: uint16(243),
	4346: uint16(1),
	4347: uint16(anon_sym_LT),
	4348: uint16(3),
	4349: uint16(2),
	4350: uint16(sym_comment),
	4351: uint16(aux_sym__tag_value_token1),
	4352: uint16(241),
	4353: uint16(6),
	4354: uint16(sym__implicit_end_tag),
	4355: uint16(anon_sym_LT_BANG),
	4356: uint16(anon_sym_LT_SLASH),
	4357: uint16(sym_entity),
	4358: uint16(sym_text),
	4359: uint16(anon_sym_LBRACE),
	4360: uint16(3),
	4361: uint16(247),
	4362: uint16(1),
	4363: uint16(anon_sym_LT),
	4364: uint16(3),
	4365: uint16(2),
	4366: uint16(sym_comment),
	4367: uint16(aux_sym__tag_value_token1),
	4368: uint16(245),
	4369: uint16(6),
	4371: uint16(anon_sym_LT_BANG),
	4372: uint16(anon_sym_LT_SLASH),
	4373: uint16(sym_entity),
	4374: uint16(sym_text),
	4375: uint16(anon_sym_LBRACE),
	4376: uint16(6),
	4377: uint16(251),
	4378: uint16(1),
	4379: uint16(sym_attribute_name),
	4380: uint16(254),
	4381: uint16(1),
	4382: uint16(anon_sym_LBRACE),
	4383: uint16(196),
	4384: uint16(1),
	4385: uint16(sym_expression),
	4386: uint16(3),
	4387: uint16(2),
	4388: uint16(sym_comment),
	4389: uint16(aux_sym__tag_value_token1),
	4390: uint16(249),
	4391: uint16(2),
	4392: uint16(anon_sym_GT),
	4393: uint16(anon_sym_SLASH_GT),
	4394: uint16(81),
	4395: uint16(2),
	4396: uint16(sym_attribute),
	4397: uint16(aux_sym_start_tag_repeat1),
	4398: uint16(3),
	4399: uint16(259),
	4400: uint16(1),
	4401: uint16(anon_sym_LT),
	4402: uint16(3),
	4403: uint16(2),
	4404: uint16(sym_comment),
	4405: uint16(aux_sym__tag_value_token1),
	4406: uint16(257),
	4407: uint16(6),
	4409: uint16(anon_sym_LT_BANG),
	4410: uint16(anon_sym_LT_SLASH),
	4411: uint16(sym_entity),
	4412: uint16(sym_text),
	4413: uint16(anon_sym_LBRACE),
	4414: uint16(3),
	4415: uint16(263),
	4416: uint16(1),
	4417: uint16(anon_sym_LT),
	4418: uint16(3),
	4419: uint16(2),
	4420: uint16(sym_comment),
	4421: uint16(aux_sym__tag_value_token1),
	4422: uint16(261),
	4423: uint16(6),
	4425: uint16(anon_sym_LT_BANG),
	4426: uint16(anon_sym_LT_SLASH),
	4427: uint16(sym_entity),
	4428: uint16(sym_text),
	4429: uint16(anon_sym_LBRACE),
	4430: uint16(3),
	4431: uint16(267),
	4432: uint16(1),
	4433: uint16(anon_sym_LT),
	4434: uint16(3),
	4435: uint16(2),
	4436: uint16(sym_comment),
	4437: uint16(aux_sym__tag_value_token1),
	4438: uint16(265),
	4439: uint16(6),
	4441: uint16(anon_sym_LT_BANG),
	4442: uint16(anon_sym_LT_SLASH),
	4443: uint16(sym_entity),
	4444: uint16(sym_text),
	4445: uint16(anon_sym_LBRACE),
	4446: uint16(3),
	4447: uint16(271),
	4448: uint16(1),
	4449: uint16(anon_sym_LT),
	4450: uint16(3),
	4451: uint16(2),
	4452: uint16(sym_comment),
	4453: uint16(aux_sym__tag_value_token1),
	4454: uint16(269),
	4455: uint16(6),
	4457: uint16(anon_sym_LT_BANG),
	4458: uint16(anon_sym_LT_SLASH),
	4459: uint16(sym_entity),
	4460: uint16(sym_text),
	4461: uint16(anon_sym_LBRACE),
	4462: uint16(3),
	4463: uint16(275),
	4464: uint16(1),
	4465: uint16(anon_sym_LT),
	4466: uint16(3),
	4467: uint16(2),
	4468: uint16(sym_comment),
	4469: uint16(aux_sym__tag_value_token1),
	4470: uint16(273),
	4471: uint16(6),
	4473: uint16(anon_sym_LT_BANG),
	4474: uint16(anon_sym_LT_SLASH),
	4475: uint16(sym_entity),
	4476: uint16(sym_text),
	4477: uint16(anon_sym_LBRACE),
	4478: uint16(3),
	4479: uint16(279),
	4480: uint16(1),
	4481: uint16(anon_sym_LT),
	4482: uint16(3),
	4483: uint16(2),
	4484: uint16(sym_comment),
	4485: uint16(aux_sym__tag_value_token1),
	4486: uint16(277),
	4487: uint16(6),
	4489: uint16(anon_sym_LT_BANG),
	4490: uint16(anon_sym_LT_SLASH),
	4491: uint16(sym_entity),
	4492: uint16(sym_text),
	4493: uint16(anon_sym_LBRACE),
	4494: uint16(3),
	4495: uint16(283),
	4496: uint16(1),
	4497: uint16(anon_sym_LT),
	4498: uint16(3),
	4499: uint16(2),
	4500: uint16(sym_comment),
	4501: uint16(aux_sym__tag_value_token1),
	4502: uint16(281),
	4503: uint16(6),
	4505: uint16(anon_sym_LT_BANG),
	4506: uint16(anon_sym_LT_SLASH),
	4507: uint16(sym_entity),
	4508: uint16(sym_text),
	4509: uint16(anon_sym_LBRACE),
	4510: uint16(3),
	4511: uint16(287),
	4512: uint16(1),
	4513: uint16(anon_sym_LT),
	4514: uint16(3),
	4515: uint16(2),
	4516: uint16(sym_comment),
	4517: uint16(aux_sym__tag_value_token1),
	4518: uint16(285),
	4519: uint16(6),
	4521: uint16(anon_sym_LT_BANG),
	4522: uint16(anon_sym_LT_SLASH),
	4523: uint16(sym_entity),
	4524: uint16(sym_text),
	4525: uint16(anon_sym_LBRACE),
	4526: uint16(3),
	4527: uint16(291),
	4528: uint16(1),
	4529: uint16(anon_sym_LT),
	4530: uint16(3),
	4531: uint16(2),
	4532: uint16(sym_comment),
	4533: uint16(aux_sym__tag_value_token1),
	4534: uint16(289),
	4535: uint16(6),
	4537: uint16(anon_sym_LT_BANG),
	4538: uint16(anon_sym_LT_SLASH),
	4539: uint16(sym_entity),
	4540: uint16(sym_text),
	4541: uint16(anon_sym_LBRACE),
	4542: uint16(3),
	4543: uint16(295),
	4544: uint16(1),
	4545: uint16(anon_sym_LT),
	4546: uint16(3),
	4547: uint16(2),
	4548: uint16(sym_comment),
	4549: uint16(aux_sym__tag_value_token1),
	4550: uint16(293),
	4551: uint16(6),
	4553: uint16(anon_sym_LT_BANG),
	4554: uint16(anon_sym_LT_SLASH),
	4555: uint16(sym_entity),
	4556: uint16(sym_text),
	4557: uint16(anon_sym_LBRACE),
	4558: uint16(3),
	4559: uint16(299),
	4560: uint16(1),
	4561: uint16(anon_sym_LT),
	4562: uint16(3),
	4563: uint16(2),
	4564: uint16(sym_comment),
	4565: uint16(aux_sym__tag_value_token1),
	4566: uint16(297),
	4567: uint16(6),
	4569: uint16(anon_sym_LT_BANG),
	4570: uint16(anon_sym_LT_SLASH),
	4571: uint16(sym_entity),
	4572: uint16(sym_text),
	4573: uint16(anon_sym_LBRACE),
	4574: uint16(3),
	4575: uint16(303),
	4576: uint16(1),
	4577: uint16(anon_sym_LT),
	4578: uint16(3),
	4579: uint16(2),
	4580: uint16(sym_comment),
	4581: uint16(aux_sym__tag_value_token1),
	4582: uint16(301),
	4583: uint16(6),
	4585: uint16(anon_sym_LT_BANG),
	4586: uint16(anon_sym_LT_SLASH),
	4587: uint16(sym_entity),
	4588: uint16(sym_text),
	4589: uint16(anon_sym_LBRACE),
	4590: uint16(3),
	4591: uint16(307),
	4592: uint16(1),
	4593: uint16(anon_sym_LT),
	4594: uint16(3),
	4595: uint16(2),
	4596: uint16(sym_comment),
	4597: uint16(aux_sym__tag_value_token1),
	4598: uint16(305),
	4599: uint16(6),
	4601: uint16(anon_sym_LT_BANG),
	4602: uint16(anon_sym_LT_SLASH),
	4603: uint16(sym_entity),
	4604: uint16(sym_text),
	4605: uint16(anon_sym_LBRACE),
	4606: uint16(3),
	4607: uint16(311),
	4608: uint16(1),
	4609: uint16(anon_sym_LT),
	4610: uint16(3),
	4611: uint16(2),
	4612: uint16(sym_comment),
	4613: uint16(aux_sym__tag_value_token1),
	4614: uint16(309),
	4615: uint16(6),
	4617: uint16(anon_sym_LT_BANG),
	4618: uint16(anon_sym_LT_SLASH),
	4619: uint16(sym_entity),
	4620: uint16(sym_text),
	4621: uint16(anon_sym_LBRACE),
	4622: uint16(3),
	4623: uint16(315),
	4624: uint16(1),
	4625: uint16(anon_sym_LT),
	4626: uint16(3),
	4627: uint16(2),
	4628: uint16(sym_comment),
	4629: uint16(aux_sym__tag_value_token1),
	4630: uint16(313),
	4631: uint16(6),
	4633: uint16(anon_sym_LT_BANG),
	4634: uint16(anon_sym_LT_SLASH),
	4635: uint16(sym_entity),
	4636: uint16(sym_text),
	4637: uint16(anon_sym_LBRACE),
	4638: uint16(3),
	4639: uint16(319),
	4640: uint16(1),
	4641: uint16(anon_sym_LT),
	4642: uint16(3),
	4643: uint16(2),
	4644: uint16(sym_comment),
	4645: uint16(aux_sym__tag_value_token1),
	4646: uint16(317),
	4647: uint16(6),
	4649: uint16(anon_sym_LT_BANG),
	4650: uint16(anon_sym_LT_SLASH),
	4651: uint16(sym_entity),
	4652: uint16(sym_text),
	4653: uint16(anon_sym_LBRACE),
	4654: uint16(3),
	4655: uint16(323),
	4656: uint16(1),
	4657: uint16(anon_sym_LT),
	4658: uint16(3),
	4659: uint16(2),
	4660: uint16(sym_comment),
	4661: uint16(aux_sym__tag_value_token1),
	4662: uint16(321),
	4663: uint16(6),
	4664: uint16(sym__implicit_end_tag),
	4665: uint16(anon_sym_LT_BANG),
	4666: uint16(anon_sym_LT_SLASH),
	4667: uint16(sym_entity),
	4668: uint16(sym_text),
	4669: uint16(anon_sym_LBRACE),
	4670: uint16(3),
	4671: uint16(327),
	4672: uint16(1),
	4673: uint16(anon_sym_LT),
	4674: uint16(3),
	4675: uint16(2),
	4676: uint16(sym_comment),
	4677: uint16(aux_sym__tag_value_token1),
	4678: uint16(325),
	4679: uint16(6),
	4680: uint16(sym__implicit_end_tag),
	4681: uint16(anon_sym_LT_BANG),
	4682: uint16(anon_sym_LT_SLASH),
	4683: uint16(sym_entity),
	4684: uint16(sym_text),
	4685: uint16(anon_sym_LBRACE),
	4686: uint16(3),
	4687: uint16(169),
	4688: uint16(1),
	4689: uint16(anon_sym_LT),
	4690: uint16(3),
	4691: uint16(2),
	4692: uint16(sym_comment),
	4693: uint16(aux_sym__tag_value_token1),
	4694: uint16(167),
	4695: uint16(6),
	4696: uint16(sym__implicit_end_tag),
	4697: uint16(anon_sym_LT_BANG),
	4698: uint16(anon_sym_LT_SLASH),
	4699: uint16(sym_entity),
	4700: uint16(sym_text),
	4701: uint16(anon_sym_LBRACE),
	4702: uint16(3),
	4703: uint16(331),
	4704: uint16(1),
	4705: uint16(anon_sym_LT),
	4706: uint16(3),
	4707: uint16(2),
	4708: uint16(sym_comment),
	4709: uint16(aux_sym__tag_value_token1),
	4710: uint16(329),
	4711: uint16(6),
	4712: uint16(sym__implicit_end_tag),
	4713: uint16(anon_sym_LT_BANG),
	4714: uint16(anon_sym_LT_SLASH),
	4715: uint16(sym_entity),
	4716: uint16(sym_text),
	4717: uint16(anon_sym_LBRACE),
	4718: uint16(3),
	4719: uint16(335),
	4720: uint16(1),
	4721: uint16(anon_sym_LT),
	4722: uint16(3),
	4723: uint16(2),
	4724: uint16(sym_comment),
	4725: uint16(aux_sym__tag_value_token1),
	4726: uint16(333),
	4727: uint16(6),
	4728: uint16(sym__implicit_end_tag),
	4729: uint16(anon_sym_LT_BANG),
	4730: uint16(anon_sym_LT_SLASH),
	4731: uint16(sym_entity),
	4732: uint16(sym_text),
	4733: uint16(anon_sym_LBRACE),
	4734: uint16(3),
	4735: uint16(339),
	4736: uint16(1),
	4737: uint16(anon_sym_LT),
	4738: uint16(3),
	4739: uint16(2),
	4740: uint16(sym_comment),
	4741: uint16(aux_sym__tag_value_token1),
	4742: uint16(337),
	4743: uint16(6),
	4744: uint16(sym__implicit_end_tag),
	4745: uint16(anon_sym_LT_BANG),
	4746: uint16(anon_sym_LT_SLASH),
	4747: uint16(sym_entity),
	4748: uint16(sym_text),
	4749: uint16(anon_sym_LBRACE),
	4750: uint16(3),
	4751: uint16(343),
	4752: uint16(1),
	4753: uint16(anon_sym_LT),
	4754: uint16(3),
	4755: uint16(2),
	4756: uint16(sym_comment),
	4757: uint16(aux_sym__tag_value_token1),
	4758: uint16(341),
	4759: uint16(6),
	4760: uint16(sym__implicit_end_tag),
	4761: uint16(anon_sym_LT_BANG),
	4762: uint16(anon_sym_LT_SLASH),
	4763: uint16(sym_entity),
	4764: uint16(sym_text),
	4765: uint16(anon_sym_LBRACE),
	4766: uint16(3),
	4767: uint16(347),
	4768: uint16(1),
	4769: uint16(anon_sym_LT),
	4770: uint16(3),
	4771: uint16(2),
	4772: uint16(sym_comment),
	4773: uint16(aux_sym__tag_value_token1),
	4774: uint16(345),
	4775: uint16(6),
	4776: uint16(sym__implicit_end_tag),
	4777: uint16(anon_sym_LT_BANG),
	4778: uint16(anon_sym_LT_SLASH),
	4779: uint16(sym_entity),
	4780: uint16(sym_text),
	4781: uint16(anon_sym_LBRACE),
	4782: uint16(3),
	4783: uint16(193),
	4784: uint16(1),
	4785: uint16(anon_sym_LT),
	4786: uint16(3),
	4787: uint16(2),
	4788: uint16(sym_comment),
	4789: uint16(aux_sym__tag_value_token1),
	4790: uint16(191),
	4791: uint16(6),
	4792: uint16(sym__implicit_end_tag),
	4793: uint16(anon_sym_LT_BANG),
	4794: uint16(anon_sym_LT_SLASH),
	4795: uint16(sym_entity),
	4796: uint16(sym_text),
	4797: uint16(anon_sym_LBRACE),
	4798: uint16(3),
	4799: uint16(173),
	4800: uint16(1),
	4801: uint16(anon_sym_LT),
	4802: uint16(3),
	4803: uint16(2),
	4804: uint16(sym_comment),
	4805: uint16(aux_sym__tag_value_token1),
	4806: uint16(171),
	4807: uint16(6),
	4808: uint16(sym__implicit_end_tag),
	4809: uint16(anon_sym_LT_BANG),
	4810: uint16(anon_sym_LT_SLASH),
	4811: uint16(sym_entity),
	4812: uint16(sym_text),
	4813: uint16(anon_sym_LBRACE),
	4814: uint16(3),
	4815: uint16(327),
	4816: uint16(1),
	4817: uint16(anon_sym_LT),
	4818: uint16(3),
	4819: uint16(2),
	4820: uint16(sym_comment),
	4821: uint16(aux_sym__tag_value_token1),
	4822: uint16(325),
	4823: uint16(6),
	4825: uint16(anon_sym_LT_BANG),
	4826: uint16(anon_sym_LT_SLASH),
	4827: uint16(sym_entity),
	4828: uint16(sym_text),
	4829: uint16(anon_sym_LBRACE),
	4830: uint16(3),
	4831: uint16(347),
	4832: uint16(1),
	4833: uint16(anon_sym_LT),
	4834: uint16(3),
	4835: uint16(2),
	4836: uint16(sym_comment),
	4837: uint16(aux_sym__tag_value_token1),
	4838: uint16(345),
	4839: uint16(6),
	4841: uint16(anon_sym_LT_BANG),
	4842: uint16(anon_sym_LT_SLASH),
	4843: uint16(sym_entity),
	4844: uint16(sym_text),
	4845: uint16(anon_sym_LBRACE),
	4846: uint16(3),
	4847: uint16(197),
	4848: uint16(1),
	4849: uint16(anon_sym_LT),
	4850: uint16(3),
	4851: uint16(2),
	4852: uint16(sym_comment),
	4853: uint16(aux_sym__tag_value_token1),
	4854: uint16(195),
	4855: uint16(6),
	4856: uint16(sym__implicit_end_tag),
	4857: uint16(anon_sym_LT_BANG),
	4858: uint16(anon_sym_LT_SLASH),
	4859: uint16(sym_entity),
	4860: uint16(sym_text),
	4861: uint16(anon_sym_LBRACE),
	4862: uint16(3),
	4863: uint16(201),
	4864: uint16(1),
	4865: uint16(anon_sym_LT),
	4866: uint16(3),
	4867: uint16(2),
	4868: uint16(sym_comment),
	4869: uint16(aux_sym__tag_value_token1),
	4870: uint16(199),
	4871: uint16(6),
	4872: uint16(sym__implicit_end_tag),
	4873: uint16(anon_sym_LT_BANG),
	4874: uint16(anon_sym_LT_SLASH),
	4875: uint16(sym_entity),
	4876: uint16(sym_text),
	4877: uint16(anon_sym_LBRACE),
	4878: uint16(3),
	4879: uint16(205),
	4880: uint16(1),
	4881: uint16(anon_sym_LT),
	4882: uint16(3),
	4883: uint16(2),
	4884: uint16(sym_comment),
	4885: uint16(aux_sym__tag_value_token1),
	4886: uint16(203),
	4887: uint16(6),
	4888: uint16(sym__implicit_end_tag),
	4889: uint16(anon_sym_LT_BANG),
	4890: uint16(anon_sym_LT_SLASH),
	4891: uint16(sym_entity),
	4892: uint16(sym_text),
	4893: uint16(anon_sym_LBRACE),
	4894: uint16(3),
	4895: uint16(209),
	4896: uint16(1),
	4897: uint16(anon_sym_LT),
	4898: uint16(3),
	4899: uint16(2),
	4900: uint16(sym_comment),
	4901: uint16(aux_sym__tag_value_token1),
	4902: uint16(207),
	4903: uint16(6),
	4904: uint16(sym__implicit_end_tag),
	4905: uint16(anon_sym_LT_BANG),
	4906: uint16(anon_sym_LT_SLASH),
	4907: uint16(sym_entity),
	4908: uint16(sym_text),
	4909: uint16(anon_sym_LBRACE),
	4910: uint16(3),
	4911: uint16(213),
	4912: uint16(1),
	4913: uint16(anon_sym_LT),
	4914: uint16(3),
	4915: uint16(2),
	4916: uint16(sym_comment),
	4917: uint16(aux_sym__tag_value_token1),
	4918: uint16(211),
	4919: uint16(6),
	4920: uint16(sym__implicit_end_tag),
	4921: uint16(anon_sym_LT_BANG),
	4922: uint16(anon_sym_LT_SLASH),
	4923: uint16(sym_entity),
	4924: uint16(sym_text),
	4925: uint16(anon_sym_LBRACE),
	4926: uint16(3),
	4927: uint16(217),
	4928: uint16(1),
	4929: uint16(anon_sym_LT),
	4930: uint16(3),
	4931: uint16(2),
	4932: uint16(sym_comment),
	4933: uint16(aux_sym__tag_value_token1),
	4934: uint16(215),
	4935: uint16(6),
	4936: uint16(sym__implicit_end_tag),
	4937: uint16(anon_sym_LT_BANG),
	4938: uint16(anon_sym_LT_SLASH),
	4939: uint16(sym_entity),
	4940: uint16(sym_text),
	4941: uint16(anon_sym_LBRACE),
	4942: uint16(3),
	4943: uint16(223),
	4944: uint16(1),
	4945: uint16(anon_sym_LT),
	4946: uint16(3),
	4947: uint16(2),
	4948: uint16(sym_comment),
	4949: uint16(aux_sym__tag_value_token1),
	4950: uint16(221),
	4951: uint16(6),
	4952: uint16(sym__implicit_end_tag),
	4953: uint16(anon_sym_LT_BANG),
	4954: uint16(anon_sym_LT_SLASH),
	4955: uint16(sym_entity),
	4956: uint16(sym_text),
	4957: uint16(anon_sym_LBRACE),
	4958: uint16(3),
	4959: uint16(227),
	4960: uint16(1),
	4961: uint16(anon_sym_LT),
	4962: uint16(3),
	4963: uint16(2),
	4964: uint16(sym_comment),
	4965: uint16(aux_sym__tag_value_token1),
	4966: uint16(225),
	4967: uint16(6),
	4968: uint16(sym__implicit_end_tag),
	4969: uint16(anon_sym_LT_BANG),
	4970: uint16(anon_sym_LT_SLASH),
	4971: uint16(sym_entity),
	4972: uint16(sym_text),
	4973: uint16(anon_sym_LBRACE),
	4974: uint16(3),
	4975: uint16(231),
	4976: uint16(1),
	4977: uint16(anon_sym_LT),
	4978: uint16(3),
	4979: uint16(2),
	4980: uint16(sym_comment),
	4981: uint16(aux_sym__tag_value_token1),
	4982: uint16(229),
	4983: uint16(6),
	4984: uint16(sym__implicit_end_tag),
	4985: uint16(anon_sym_LT_BANG),
	4986: uint16(anon_sym_LT_SLASH),
	4987: uint16(sym_entity),
	4988: uint16(sym_text),
	4989: uint16(anon_sym_LBRACE),
	4990: uint16(3),
	4991: uint16(235),
	4992: uint16(1),
	4993: uint16(anon_sym_LT),
	4994: uint16(3),
	4995: uint16(2),
	4996: uint16(sym_comment),
	4997: uint16(aux_sym__tag_value_token1),
	4998: uint16(233),
	4999: uint16(6),
	5000: uint16(sym__implicit_end_tag),
	5001: uint16(anon_sym_LT_BANG),
	5002: uint16(anon_sym_LT_SLASH),
	5003: uint16(sym_entity),
	5004: uint16(sym_text),
	5005: uint16(anon_sym_LBRACE),
	5006: uint16(3),
	5007: uint16(239),
	5008: uint16(1),
	5009: uint16(anon_sym_LT),
	5010: uint16(3),
	5011: uint16(2),
	5012: uint16(sym_comment),
	5013: uint16(aux_sym__tag_value_token1),
	5014: uint16(237),
	5015: uint16(6),
	5016: uint16(sym__implicit_end_tag),
	5017: uint16(anon_sym_LT_BANG),
	5018: uint16(anon_sym_LT_SLASH),
	5019: uint16(sym_entity),
	5020: uint16(sym_text),
	5021: uint16(anon_sym_LBRACE),
	5022: uint16(3),
	5023: uint16(247),
	5024: uint16(1),
	5025: uint16(anon_sym_LT),
	5026: uint16(3),
	5027: uint16(2),
	5028: uint16(sym_comment),
	5029: uint16(aux_sym__tag_value_token1),
	5030: uint16(245),
	5031: uint16(6),
	5032: uint16(sym__implicit_end_tag),
	5033: uint16(anon_sym_LT_BANG),
	5034: uint16(anon_sym_LT_SLASH),
	5035: uint16(sym_entity),
	5036: uint16(sym_text),
	5037: uint16(anon_sym_LBRACE),
	5038: uint16(3),
	5039: uint16(259),
	5040: uint16(1),
	5041: uint16(anon_sym_LT),
	5042: uint16(3),
	5043: uint16(2),
	5044: uint16(sym_comment),
	5045: uint16(aux_sym__tag_value_token1),
	5046: uint16(257),
	5047: uint16(6),
	5048: uint16(sym__implicit_end_tag),
	5049: uint16(anon_sym_LT_BANG),
	5050: uint16(anon_sym_LT_SLASH),
	5051: uint16(sym_entity),
	5052: uint16(sym_text),
	5053: uint16(anon_sym_LBRACE),
	5054: uint16(3),
	5055: uint16(263),
	5056: uint16(1),
	5057: uint16(anon_sym_LT),
	5058: uint16(3),
	5059: uint16(2),
	5060: uint16(sym_comment),
	5061: uint16(aux_sym__tag_value_token1),
	5062: uint16(261),
	5063: uint16(6),
	5064: uint16(sym__implicit_end_tag),
	5065: uint16(anon_sym_LT_BANG),
	5066: uint16(anon_sym_LT_SLASH),
	5067: uint16(sym_entity),
	5068: uint16(sym_text),
	5069: uint16(anon_sym_LBRACE),
	5070: uint16(3),
	5071: uint16(267),
	5072: uint16(1),
	5073: uint16(anon_sym_LT),
	5074: uint16(3),
	5075: uint16(2),
	5076: uint16(sym_comment),
	5077: uint16(aux_sym__tag_value_token1),
	5078: uint16(265),
	5079: uint16(6),
	5080: uint16(sym__implicit_end_tag),
	5081: uint16(anon_sym_LT_BANG),
	5082: uint16(anon_sym_LT_SLASH),
	5083: uint16(sym_entity),
	5084: uint16(sym_text),
	5085: uint16(anon_sym_LBRACE),
	5086: uint16(3),
	5087: uint16(271),
	5088: uint16(1),
	5089: uint16(anon_sym_LT),
	5090: uint16(3),
	5091: uint16(2),
	5092: uint16(sym_comment),
	5093: uint16(aux_sym__tag_value_token1),
	5094: uint16(269),
	5095: uint16(6),
	5096: uint16(sym__implicit_end_tag),
	5097: uint16(anon_sym_LT_BANG),
	5098: uint16(anon_sym_LT_SLASH),
	5099: uint16(sym_entity),
	5100: uint16(sym_text),
	5101: uint16(anon_sym_LBRACE),
	5102: uint16(3),
	5103: uint16(275),
	5104: uint16(1),
	5105: uint16(anon_sym_LT),
	5106: uint16(3),
	5107: uint16(2),
	5108: uint16(sym_comment),
	5109: uint16(aux_sym__tag_value_token1),
	5110: uint16(273),
	5111: uint16(6),
	5112: uint16(sym__implicit_end_tag),
	5113: uint16(anon_sym_LT_BANG),
	5114: uint16(anon_sym_LT_SLASH),
	5115: uint16(sym_entity),
	5116: uint16(sym_text),
	5117: uint16(anon_sym_LBRACE),
	5118: uint16(3),
	5119: uint16(279),
	5120: uint16(1),
	5121: uint16(anon_sym_LT),
	5122: uint16(3),
	5123: uint16(2),
	5124: uint16(sym_comment),
	5125: uint16(aux_sym__tag_value_token1),
	5126: uint16(277),
	5127: uint16(6),
	5128: uint16(sym__implicit_end_tag),
	5129: uint16(anon_sym_LT_BANG),
	5130: uint16(anon_sym_LT_SLASH),
	5131: uint16(sym_entity),
	5132: uint16(sym_text),
	5133: uint16(anon_sym_LBRACE),
	5134: uint16(3),
	5135: uint16(283),
	5136: uint16(1),
	5137: uint16(anon_sym_LT),
	5138: uint16(3),
	5139: uint16(2),
	5140: uint16(sym_comment),
	5141: uint16(aux_sym__tag_value_token1),
	5142: uint16(281),
	5143: uint16(6),
	5144: uint16(sym__implicit_end_tag),
	5145: uint16(anon_sym_LT_BANG),
	5146: uint16(anon_sym_LT_SLASH),
	5147: uint16(sym_entity),
	5148: uint16(sym_text),
	5149: uint16(anon_sym_LBRACE),
	5150: uint16(3),
	5151: uint16(287),
	5152: uint16(1),
	5153: uint16(anon_sym_LT),
	5154: uint16(3),
	5155: uint16(2),
	5156: uint16(sym_comment),
	5157: uint16(aux_sym__tag_value_token1),
	5158: uint16(285),
	5159: uint16(6),
	5160: uint16(sym__implicit_end_tag),
	5161: uint16(anon_sym_LT_BANG),
	5162: uint16(anon_sym_LT_SLASH),
	5163: uint16(sym_entity),
	5164: uint16(sym_text),
	5165: uint16(anon_sym_LBRACE),
	5166: uint16(3),
	5167: uint16(291),
	5168: uint16(1),
	5169: uint16(anon_sym_LT),
	5170: uint16(3),
	5171: uint16(2),
	5172: uint16(sym_comment),
	5173: uint16(aux_sym__tag_value_token1),
	5174: uint16(289),
	5175: uint16(6),
	5176: uint16(sym__implicit_end_tag),
	5177: uint16(anon_sym_LT_BANG),
	5178: uint16(anon_sym_LT_SLASH),
	5179: uint16(sym_entity),
	5180: uint16(sym_text),
	5181: uint16(anon_sym_LBRACE),
	5182: uint16(3),
	5183: uint16(295),
	5184: uint16(1),
	5185: uint16(anon_sym_LT),
	5186: uint16(3),
	5187: uint16(2),
	5188: uint16(sym_comment),
	5189: uint16(aux_sym__tag_value_token1),
	5190: uint16(293),
	5191: uint16(6),
	5192: uint16(sym__implicit_end_tag),
	5193: uint16(anon_sym_LT_BANG),
	5194: uint16(anon_sym_LT_SLASH),
	5195: uint16(sym_entity),
	5196: uint16(sym_text),
	5197: uint16(anon_sym_LBRACE),
	5198: uint16(3),
	5199: uint16(299),
	5200: uint16(1),
	5201: uint16(anon_sym_LT),
	5202: uint16(3),
	5203: uint16(2),
	5204: uint16(sym_comment),
	5205: uint16(aux_sym__tag_value_token1),
	5206: uint16(297),
	5207: uint16(6),
	5208: uint16(sym__implicit_end_tag),
	5209: uint16(anon_sym_LT_BANG),
	5210: uint16(anon_sym_LT_SLASH),
	5211: uint16(sym_entity),
	5212: uint16(sym_text),
	5213: uint16(anon_sym_LBRACE),
	5214: uint16(3),
	5215: uint16(303),
	5216: uint16(1),
	5217: uint16(anon_sym_LT),
	5218: uint16(3),
	5219: uint16(2),
	5220: uint16(sym_comment),
	5221: uint16(aux_sym__tag_value_token1),
	5222: uint16(301),
	5223: uint16(6),
	5224: uint16(sym__implicit_end_tag),
	5225: uint16(anon_sym_LT_BANG),
	5226: uint16(anon_sym_LT_SLASH),
	5227: uint16(sym_entity),
	5228: uint16(sym_text),
	5229: uint16(anon_sym_LBRACE),
	5230: uint16(3),
	5231: uint16(307),
	5232: uint16(1),
	5233: uint16(anon_sym_LT),
	5234: uint16(3),
	5235: uint16(2),
	5236: uint16(sym_comment),
	5237: uint16(aux_sym__tag_value_token1),
	5238: uint16(305),
	5239: uint16(6),
	5240: uint16(sym__implicit_end_tag),
	5241: uint16(anon_sym_LT_BANG),
	5242: uint16(anon_sym_LT_SLASH),
	5243: uint16(sym_entity),
	5244: uint16(sym_text),
	5245: uint16(anon_sym_LBRACE),
	5246: uint16(3),
	5247: uint16(311),
	5248: uint16(1),
	5249: uint16(anon_sym_LT),
	5250: uint16(3),
	5251: uint16(2),
	5252: uint16(sym_comment),
	5253: uint16(aux_sym__tag_value_token1),
	5254: uint16(309),
	5255: uint16(6),
	5256: uint16(sym__implicit_end_tag),
	5257: uint16(anon_sym_LT_BANG),
	5258: uint16(anon_sym_LT_SLASH),
	5259: uint16(sym_entity),
	5260: uint16(sym_text),
	5261: uint16(anon_sym_LBRACE),
	5262: uint16(3),
	5263: uint16(315),
	5264: uint16(1),
	5265: uint16(anon_sym_LT),
	5266: uint16(3),
	5267: uint16(2),
	5268: uint16(sym_comment),
	5269: uint16(aux_sym__tag_value_token1),
	5270: uint16(313),
	5271: uint16(6),
	5272: uint16(sym__implicit_end_tag),
	5273: uint16(anon_sym_LT_BANG),
	5274: uint16(anon_sym_LT_SLASH),
	5275: uint16(sym_entity),
	5276: uint16(sym_text),
	5277: uint16(anon_sym_LBRACE),
	5278: uint16(3),
	5279: uint16(319),
	5280: uint16(1),
	5281: uint16(anon_sym_LT),
	5282: uint16(3),
	5283: uint16(2),
	5284: uint16(sym_comment),
	5285: uint16(aux_sym__tag_value_token1),
	5286: uint16(317),
	5287: uint16(6),
	5288: uint16(sym__implicit_end_tag),
	5289: uint16(anon_sym_LT_BANG),
	5290: uint16(anon_sym_LT_SLASH),
	5291: uint16(sym_entity),
	5292: uint16(sym_text),
	5293: uint16(anon_sym_LBRACE),
	5294: uint16(3),
	5295: uint16(331),
	5296: uint16(1),
	5297: uint16(anon_sym_LT),
	5298: uint16(3),
	5299: uint16(2),
	5300: uint16(sym_comment),
	5301: uint16(aux_sym__tag_value_token1),
	5302: uint16(329),
	5303: uint16(6),
	5305: uint16(anon_sym_LT_BANG),
	5306: uint16(anon_sym_LT_SLASH),
	5307: uint16(sym_entity),
	5308: uint16(sym_text),
	5309: uint16(anon_sym_LBRACE),
	5310: uint16(3),
	5311: uint16(335),
	5312: uint16(1),
	5313: uint16(anon_sym_LT),
	5314: uint16(3),
	5315: uint16(2),
	5316: uint16(sym_comment),
	5317: uint16(aux_sym__tag_value_token1),
	5318: uint16(333),
	5319: uint16(6),
	5321: uint16(anon_sym_LT_BANG),
	5322: uint16(anon_sym_LT_SLASH),
	5323: uint16(sym_entity),
	5324: uint16(sym_text),
	5325: uint16(anon_sym_LBRACE),
	5326: uint16(3),
	5327: uint16(323),
	5328: uint16(1),
	5329: uint16(anon_sym_LT),
	5330: uint16(3),
	5331: uint16(2),
	5332: uint16(sym_comment),
	5333: uint16(aux_sym__tag_value_token1),
	5334: uint16(321),
	5335: uint16(6),
	5337: uint16(anon_sym_LT_BANG),
	5338: uint16(anon_sym_LT_SLASH),
	5339: uint16(sym_entity),
	5340: uint16(sym_text),
	5341: uint16(anon_sym_LBRACE),
	5342: uint16(7),
	5343: uint16(187),
	5344: uint16(1),
	5345: uint16(sym_attribute_name),
	5346: uint16(189),
	5347: uint16(1),
	5348: uint16(anon_sym_LBRACE),
	5349: uint16(349),
	5350: uint16(1),
	5351: uint16(anon_sym_GT),
	5352: uint16(351),
	5353: uint16(1),
	5354: uint16(anon_sym_SLASH_GT),
	5355: uint16(196),
	5356: uint16(1),
	5357: uint16(sym_expression),
	5358: uint16(3),
	5359: uint16(2),
	5360: uint16(sym_comment),
	5361: uint16(aux_sym__tag_value_token1),
	5362: uint16(145),
	5363: uint16(2),
	5364: uint16(sym_attribute),
	5365: uint16(aux_sym_start_tag_repeat1),
	5366: uint16(7),
	5367: uint16(353),
	5368: uint16(1),
	5369: uint16(anon_sym_LBRACE),
	5370: uint16(37),
	5371: uint16(1),
	5372: uint16(sym_else_if_start),
	5373: uint16(40),
	5374: uint16(1),
	5375: uint16(sym_else_start),
	5376: uint16(115),
	5377: uint16(1),
	5378: uint16(sym_if_end),
	5379: uint16(225),
	5380: uint16(1),
	5381: uint16(sym_else_block),
	5382: uint16(3),
	5383: uint16(2),
	5384: uint16(sym_comment),
	5385: uint16(aux_sym__tag_value_token1),
	5386: uint16(187),
	5387: uint16(2),
	5388: uint16(sym_else_if_block),
	5389: uint16(aux_sym_if_statement_repeat1),
	5390: uint16(3),
	5391: uint16(339),
	5392: uint16(1),
	5393: uint16(anon_sym_LT),
	5394: uint16(3),
	5395: uint16(2),
	5396: uint16(sym_comment),
	5397: uint16(aux_sym__tag_value_token1),
	5398: uint16(337),
	5399: uint16(6),
	5401: uint16(anon_sym_LT_BANG),
	5402: uint16(anon_sym_LT_SLASH),
	5403: uint16(sym_entity),
	5404: uint16(sym_text),
	5405: uint16(anon_sym_LBRACE),
	5406: uint16(7),
	5407: uint16(219),
	5408: uint16(1),
	5409: uint16(anon_sym_LBRACE),
	5410: uint16(37),
	5411: uint16(1),
	5412: uint16(sym_else_if_start),
	5413: uint16(40),
	5414: uint16(1),
	5415: uint16(sym_else_start),
	5416: uint16(72),
	5417: uint16(1),
	5418: uint16(sym_if_end),
	5419: uint16(230),
	5420: uint16(1),
	5421: uint16(sym_else_block),
	5422: uint16(3),
	5423: uint16(2),
	5424: uint16(sym_comment),
	5425: uint16(aux_sym__tag_value_token1),
	5426: uint16(187),
	5427: uint16(2),
	5428: uint16(sym_else_if_block),
	5429: uint16(aux_sym_if_statement_repeat1),
	5430: uint16(7),
	5431: uint16(183),
	5432: uint16(1),
	5433: uint16(anon_sym_GT),
	5434: uint16(187),
	5435: uint16(1),
	5436: uint16(sym_attribute_name),
	5437: uint16(189),
	5438: uint16(1),
	5439: uint16(anon_sym_LBRACE),
	5440: uint16(355),
	5441: uint16(1),
	5442: uint16(anon_sym_SLASH_GT),
	5443: uint16(196),
	5444: uint16(1),
	5445: uint16(sym_expression),
	5446: uint16(3),
	5447: uint16(2),
	5448: uint16(sym_comment),
	5449: uint16(aux_sym__tag_value_token1),
	5450: uint16(81),
	5451: uint16(2),
	5452: uint16(sym_attribute),
	5453: uint16(aux_sym_start_tag_repeat1),
	5454: uint16(7),
	5455: uint16(353),
	5456: uint16(1),
	5457: uint16(anon_sym_LBRACE),
	5458: uint16(37),
	5459: uint16(1),
	5460: uint16(sym_else_if_start),
	5461: uint16(40),
	5462: uint16(1),
	5463: uint16(sym_else_start),
	5464: uint16(128),
	5465: uint16(1),
	5466: uint16(sym_if_end),
	5467: uint16(235),
	5468: uint16(1),
	5469: uint16(sym_else_block),
	5470: uint16(3),
	5471: uint16(2),
	5472: uint16(sym_comment),
	5473: uint16(aux_sym__tag_value_token1),
	5474: uint16(187),
	5475: uint16(2),
	5476: uint16(sym_else_if_block),
	5477: uint16(aux_sym_if_statement_repeat1),
	5478: uint16(3),
	5479: uint16(343),
	5480: uint16(1),
	5481: uint16(anon_sym_LT),
	5482: uint16(3),
	5483: uint16(2),
	5484: uint16(sym_comment),
	5485: uint16(aux_sym__tag_value_token1),
	5486: uint16(341),
	5487: uint16(6),
	5489: uint16(anon_sym_LT_BANG),
	5490: uint16(anon_sym_LT_SLASH),
	5491: uint16(sym_entity),
	5492: uint16(sym_text),
	5493: uint16(anon_sym_LBRACE),
	5494: uint16(7),
	5495: uint16(187),
	5496: uint16(1),
	5497: uint16(sym_attribute_name),
	5498: uint16(189),
	5499: uint16(1),
	5500: uint16(anon_sym_LBRACE),
	5501: uint16(349),
	5502: uint16(1),
	5503: uint16(anon_sym_GT),
	5504: uint16(357),
	5505: uint16(1),
	5506: uint16(anon_sym_SLASH_GT),
	5507: uint16(196),
	5508: uint16(1),
	5509: uint16(sym_expression),
	5510: uint16(3),
	5511: uint16(2),
	5512: uint16(sym_comment),
	5513: uint16(aux_sym__tag_value_token1),
	5514: uint16(65),
	5515: uint16(2),
	5516: uint16(sym_attribute),
	5517: uint16(aux_sym_start_tag_repeat1),
	5518: uint16(3),
	5519: uint16(181),
	5520: uint16(1),
	5521: uint16(anon_sym_LT),
	5522: uint16(3),
	5523: uint16(2),
	5524: uint16(sym_comment),
	5525: uint16(aux_sym__tag_value_token1),
	5526: uint16(179),
	5527: uint16(6),
	5528: uint16(sym__implicit_end_tag),
	5529: uint16(anon_sym_LT_BANG),
	5530: uint16(anon_sym_LT_SLASH),
	5531: uint16(sym_entity),
	5532: uint16(sym_text),
	5533: uint16(anon_sym_LBRACE),
	5534: uint16(6),
	5535: uint16(359),
	5536: uint16(1),
	5537: uint16(anon_sym_GT),
	5538: uint16(361),
	5539: uint16(1),
	5540: uint16(sym_attribute_name),
	5541: uint16(363),
	5542: uint16(1),
	5543: uint16(anon_sym_LBRACE),
	5544: uint16(198),
	5545: uint16(1),
	5546: uint16(sym_expression),
	5547: uint16(3),
	5548: uint16(2),
	5549: uint16(sym_comment),
	5550: uint16(aux_sym__tag_value_token1),
	5551: uint16(163),
	5552: uint16(2),
	5553: uint16(sym_attribute),
	5554: uint16(aux_sym_start_tag_repeat1),
	5555: uint16(3),
	5556: uint16(367),
	5557: uint16(1),
	5558: uint16(anon_sym_LT),
	5559: uint16(3),
	5560: uint16(2),
	5561: uint16(sym_comment),
	5562: uint16(aux_sym__tag_value_token1),
	5563: uint16(365),
	5564: uint16(5),
	5565: uint16(anon_sym_LT_BANG),
	5566: uint16(anon_sym_LT_SLASH),
	5567: uint16(sym_entity),
	5568: uint16(sym_text),
	5569: uint16(anon_sym_LBRACE),
	5570: uint16(3),
	5571: uint16(371),
	5572: uint16(1),
	5573: uint16(anon_sym_LT),
	5574: uint16(3),
	5575: uint16(2),
	5576: uint16(sym_comment),
	5577: uint16(aux_sym__tag_value_token1),
	5578: uint16(369),
	5579: uint16(5),
	5580: uint16(anon_sym_LT_BANG),
	5581: uint16(anon_sym_LT_SLASH),
	5582: uint16(sym_entity),
	5583: uint16(sym_text),
	5584: uint16(anon_sym_LBRACE),
	5585: uint16(3),
	5586: uint16(375),
	5587: uint16(1),
	5588: uint16(anon_sym_LT),
	5589: uint16(3),
	5590: uint16(2),
	5591: uint16(sym_comment),
	5592: uint16(aux_sym__tag_value_token1),
	5593: uint16(373),
	5594: uint16(5),
	5595: uint16(anon_sym_LT_BANG),
	5596: uint16(anon_sym_LT_SLASH),
	5597: uint16(sym_entity),
	5598: uint16(sym_text),
	5599: uint16(anon_sym_LBRACE),
	5600: uint16(3),
	5601: uint16(379),
	5602: uint16(1),
	5603: uint16(anon_sym_LT),
	5604: uint16(3),
	5605: uint16(2),
	5606: uint16(sym_comment),
	5607: uint16(aux_sym__tag_value_token1),
	5608: uint16(377),
	5609: uint16(5),
	5610: uint16(anon_sym_LT_BANG),
	5611: uint16(anon_sym_LT_SLASH),
	5612: uint16(sym_entity),
	5613: uint16(sym_text),
	5614: uint16(anon_sym_LBRACE),
	5615: uint16(3),
	5616: uint16(383),
	5617: uint16(1),
	5618: uint16(anon_sym_LT),
	5619: uint16(3),
	5620: uint16(2),
	5621: uint16(sym_comment),
	5622: uint16(aux_sym__tag_value_token1),
	5623: uint16(381),
	5624: uint16(5),
	5625: uint16(anon_sym_LT_BANG),
	5626: uint16(anon_sym_LT_SLASH),
	5627: uint16(sym_entity),
	5628: uint16(sym_text),
	5629: uint16(anon_sym_LBRACE),
	5630: uint16(3),
	5631: uint16(387),
	5632: uint16(1),
	5633: uint16(anon_sym_LT),
	5634: uint16(3),
	5635: uint16(2),
	5636: uint16(sym_comment),
	5637: uint16(aux_sym__tag_value_token1),
	5638: uint16(385),
	5639: uint16(5),
	5640: uint16(anon_sym_LT_BANG),
	5641: uint16(anon_sym_LT_SLASH),
	5642: uint16(sym_entity),
	5643: uint16(sym_text),
	5644: uint16(anon_sym_LBRACE),
	5645: uint16(3),
	5646: uint16(391),
	5647: uint16(1),
	5648: uint16(anon_sym_LT),
	5649: uint16(3),
	5650: uint16(2),
	5651: uint16(sym_comment),
	5652: uint16(aux_sym__tag_value_token1),
	5653: uint16(389),
	5654: uint16(5),
	5655: uint16(anon_sym_LT_BANG),
	5656: uint16(anon_sym_LT_SLASH),
	5657: uint16(sym_entity),
	5658: uint16(sym_text),
	5659: uint16(anon_sym_LBRACE),
	5660: uint16(3),
	5661: uint16(395),
	5662: uint16(1),
	5663: uint16(anon_sym_LT),
	5664: uint16(3),
	5665: uint16(2),
	5666: uint16(sym_comment),
	5667: uint16(aux_sym__tag_value_token1),
	5668: uint16(393),
	5669: uint16(5),
	5670: uint16(anon_sym_LT_BANG),
	5671: uint16(anon_sym_LT_SLASH),
	5672: uint16(sym_entity),
	5673: uint16(sym_text),
	5674: uint16(anon_sym_LBRACE),
	5675: uint16(3),
	5676: uint16(399),
	5677: uint16(1),
	5678: uint16(anon_sym_LT),
	5679: uint16(3),
	5680: uint16(2),
	5681: uint16(sym_comment),
	5682: uint16(aux_sym__tag_value_token1),
	5683: uint16(397),
	5684: uint16(5),
	5685: uint16(anon_sym_LT_BANG),
	5686: uint16(anon_sym_LT_SLASH),
	5687: uint16(sym_entity),
	5688: uint16(sym_text),
	5689: uint16(anon_sym_LBRACE),
	5690: uint16(3),
	5691: uint16(403),
	5692: uint16(1),
	5693: uint16(anon_sym_LT),
	5694: uint16(3),
	5695: uint16(2),
	5696: uint16(sym_comment),
	5697: uint16(aux_sym__tag_value_token1),
	5698: uint16(401),
	5699: uint16(5),
	5700: uint16(anon_sym_LT_BANG),
	5701: uint16(anon_sym_LT_SLASH),
	5702: uint16(sym_entity),
	5703: uint16(sym_text),
	5704: uint16(anon_sym_LBRACE),
	5705: uint16(3),
	5706: uint16(407),
	5707: uint16(1),
	5708: uint16(anon_sym_LT),
	5709: uint16(3),
	5710: uint16(2),
	5711: uint16(sym_comment),
	5712: uint16(aux_sym__tag_value_token1),
	5713: uint16(405),
	5714: uint16(5),
	5715: uint16(anon_sym_LT_BANG),
	5716: uint16(anon_sym_LT_SLASH),
	5717: uint16(sym_entity),
	5718: uint16(sym_text),
	5719: uint16(anon_sym_LBRACE),
	5720: uint16(6),
	5721: uint16(361),
	5722: uint16(1),
	5723: uint16(sym_attribute_name),
	5724: uint16(363),
	5725: uint16(1),
	5726: uint16(anon_sym_LBRACE),
	5727: uint16(409),
	5728: uint16(1),
	5729: uint16(anon_sym_GT),
	5730: uint16(198),
	5731: uint16(1),
	5732: uint16(sym_expression),
	5733: uint16(3),
	5734: uint16(2),
	5735: uint16(sym_comment),
	5736: uint16(aux_sym__tag_value_token1),
	5737: uint16(163),
	5738: uint16(2),
	5739: uint16(sym_attribute),
	5740: uint16(aux_sym_start_tag_repeat1),
	5741: uint16(6),
	5742: uint16(249),
	5743: uint16(1),
	5744: uint16(anon_sym_GT),
	5745: uint16(411),
	5746: uint16(1),
	5747: uint16(sym_attribute_name),
	5748: uint16(414),
	5749: uint16(1),
	5750: uint16(anon_sym_LBRACE),
	5751: uint16(198),
	5752: uint16(1),
	5753: uint16(sym_expression),
	5754: uint16(3),
	5755: uint16(2),
	5756: uint16(sym_comment),
	5757: uint16(aux_sym__tag_value_token1),
	5758: uint16(163),
	5759: uint16(2),
	5760: uint16(sym_attribute),
	5761: uint16(aux_sym_start_tag_repeat1),
	5762: uint16(6),
	5763: uint16(189),
	5764: uint16(1),
	5765: uint16(anon_sym_LBRACE),
	5766: uint16(417),
	5767: uint16(1),
	5768: uint16(sym_attribute_value),
	5769: uint16(419),
	5770: uint16(1),
	5771: uint16(anon_sym_SQUOTE),
	5772: uint16(421),
	5773: uint16(1),
	5774: uint16(anon_sym_DQUOTE),
	5775: uint16(3),
	5776: uint16(2),
	5777: uint16(sym_comment),
	5778: uint16(aux_sym__tag_value_token1),
	5779: uint16(188),
	5780: uint16(2),
	5781: uint16(sym_quoted_attribute_value),
	5782: uint16(sym_expression),
	5783: uint16(3),
	5784: uint16(425),
	5785: uint16(1),
	5786: uint16(anon_sym_LT),
	5787: uint16(3),
	5788: uint16(2),
	5789: uint16(sym_comment),
	5790: uint16(aux_sym__tag_value_token1),
	5791: uint16(423),
	5792: uint16(5),
	5793: uint16(anon_sym_LT_BANG),
	5794: uint16(anon_sym_LT_SLASH),
	5795: uint16(sym_entity),
	5796: uint16(sym_text),
	5797: uint16(anon_sym_LBRACE),
	5798: uint16(6),
	5799: uint16(363),
	5800: uint16(1),
	5801: uint16(anon_sym_LBRACE),
	5802: uint16(427),
	5803: uint16(1),
	5804: uint16(sym_attribute_value),
	5805: uint16(429),
	5806: uint16(1),
	5807: uint16(anon_sym_SQUOTE),
	5808: uint16(431),
	5809: uint16(1),
	5810: uint16(anon_sym_DQUOTE),
	5811: uint16(3),
	5812: uint16(2),
	5813: uint16(sym_comment),
	5814: uint16(aux_sym__tag_value_token1),
	5815: uint16(199),
	5816: uint16(2),
	5817: uint16(sym_quoted_attribute_value),
	5818: uint16(sym_expression),
	5819: uint16(6),
	5820: uint16(361),
	5821: uint16(1),
	5822: uint16(sym_attribute_name),
	5823: uint16(363),
	5824: uint16(1),
	5825: uint16(anon_sym_LBRACE),
	5826: uint16(433),
	5827: uint16(1),
	5828: uint16(anon_sym_GT),
	5829: uint16(198),
	5830: uint16(1),
	5831: uint16(sym_expression),
	5832: uint16(3),
	5833: uint16(2),
	5834: uint16(sym_comment),
	5835: uint16(aux_sym__tag_value_token1),
	5836: uint16(162),
	5837: uint16(2),
	5838: uint16(sym_attribute),
	5839: uint16(aux_sym_start_tag_repeat1),
	5840: uint16(6),
	5841: uint16(361),
	5842: uint16(1),
	5843: uint16(sym_attribute_name),
	5844: uint16(363),
	5845: uint16(1),
	5846: uint16(anon_sym_LBRACE),
	5847: uint16(435),
	5848: uint16(1),
	5849: uint16(anon_sym_GT),
	5850: uint16(198),
	5851: uint16(1),
	5852: uint16(sym_expression),
	5853: uint16(3),
	5854: uint16(2),
	5855: uint16(sym_comment),
	5856: uint16(aux_sym__tag_value_token1),
	5857: uint16(150),
	5858: uint16(2),
	5859: uint16(sym_attribute),
	5860: uint16(aux_sym_start_tag_repeat1),
	5861: uint16(3),
	5862: uint16(439),
	5863: uint16(1),
	5864: uint16(anon_sym_LT),
	5865: uint16(3),
	5866: uint16(2),
	5867: uint16(sym_comment),
	5868: uint16(aux_sym__tag_value_token1),
	5869: uint16(437),
	5870: uint16(5),
	5871: uint16(anon_sym_LT_BANG),
	5872: uint16(anon_sym_LT_SLASH),
	5873: uint16(sym_entity),
	5874: uint16(sym_text),
	5875: uint16(anon_sym_LBRACE),
	5876: uint16(6),
	5877: uint16(441),
	5878: uint16(1),
	5879: uint16(anon_sym_if),
	5880: uint16(443),
	5881: uint16(1),
	5882: uint16(anon_sym_each),
	5883: uint16(445),
	5884: uint16(1),
	5885: uint16(anon_sym_await),
	5886: uint16(447),
	5887: uint16(1),
	5888: uint16(anon_sym_key),
	5889: uint16(449),
	5890: uint16(1),
	5891: uint16(anon_sym_snippet),
	5892: uint16(3),
	5893: uint16(2),
	5894: uint16(sym_comment),
	5895: uint16(aux_sym__tag_value_token1),
	5896: uint16(7),
	5897: uint16(3),
	5898: uint16(1),
	5899: uint16(sym_comment),
	5900: uint16(451),
	5901: uint16(1),
	5902: uint16(anon_sym_SQUOTE),
	5903: uint16(453),
	5904: uint16(1),
	5905: uint16(aux_sym__single_quoted_attribute_value_token1),
	5906: uint16(455),
	5907: uint16(1),
	5908: uint16(anon_sym_LBRACE),
	5909: uint16(457),
	5910: uint16(1),
	5911: uint16(aux_sym__tag_value_token1),
	5912: uint16(173),
	5913: uint16(1),
	5914: uint16(aux_sym__single_quoted_attribute_value),
	5915: uint16(202),
	5916: uint16(1),
	5917: uint16(sym_expression),
	5918: uint16(7),
	5919: uint16(3),
	5920: uint16(1),
	5921: uint16(sym_comment),
	5922: uint16(451),
	5923: uint16(1),
	5924: uint16(anon_sym_DQUOTE),
	5925: uint16(457),
	5926: uint16(1),
	5927: uint16(aux_sym__tag_value_token1),
	5928: uint16(459),
	5929: uint16(1),
	5930: uint16(aux_sym__double_quoted_attribute_value_token1),
	5931: uint16(461),
	5932: uint16(1),
	5933: uint16(anon_sym_LBRACE),
	5934: uint16(174),
	5935: uint16(1),
	5936: uint16(aux_sym__double_quoted_attribute_value),
	5937: uint16(208),
	5938: uint16(1),
	5939: uint16(sym_expression),
	5940: uint16(7),
	5941: uint16(3),
	5942: uint16(1),
	5943: uint16(sym_comment),
	5944: uint16(457),
	5945: uint16(1),
	5946: uint16(aux_sym__tag_value_token1),
	5947: uint16(463),
	5948: uint16(1),
	5949: uint16(anon_sym_SQUOTE),
	5950: uint16(465),
	5951: uint16(1),
	5952: uint16(aux_sym__single_quoted_attribute_value_token1),
	5953: uint16(468),
	5954: uint16(1),
	5955: uint16(anon_sym_LBRACE),
	5956: uint16(173),
	5957: uint16(1),
	5958: uint16(aux_sym__single_quoted_attribute_value),
	5959: uint16(202),
	5960: uint16(1),
	5961: uint16(sym_expression),
	5962: uint16(7),
	5963: uint16(3),
	5964: uint16(1),
	5965: uint16(sym_comment),
	5966: uint16(457),
	5967: uint16(1),
	5968: uint16(aux_sym__tag_value_token1),
	5969: uint16(471),
	5970: uint16(1),
	5971: uint16(anon_sym_DQUOTE),
	5972: uint16(473),
	5973: uint16(1),
	5974: uint16(aux_sym__double_quoted_attribute_value_token1),
	5975: uint16(476),
	5976: uint16(1),
	5977: uint16(anon_sym_LBRACE),
	5978: uint16(174),
	5979: uint16(1),
	5980: uint16(aux_sym__double_quoted_attribute_value),
	5981: uint16(208),
	5982: uint16(1),
	5983: uint16(sym_expression),
	5984: uint16(3),
	5985: uint16(481),
	5986: uint16(1),
	5987: uint16(anon_sym_EQ),
	5988: uint16(3),
	5989: uint16(2),
	5990: uint16(sym_comment),
	5991: uint16(aux_sym__tag_value_token1),
	5992: uint16(479),
	5993: uint16(4),
	5994: uint16(anon_sym_GT),
	5995: uint16(anon_sym_SLASH_GT),
	5996: uint16(sym_attribute_name),
	5997: uint16(anon_sym_LBRACE),
	5998: uint16(6),
	5999: uint16(153),
	6000: uint16(1),
	6001: uint16(anon_sym_COLON),
	6002: uint16(155),
	6003: uint16(1),
	6004: uint16(anon_sym_SLASH),
	6005: uint16(284),
	6006: uint16(1),
	6007: uint16(sym__else_if_tag),
	6008: uint16(285),
	6009: uint16(1),
	6010: uint16(sym__else_tag),
	6011: uint16(286),
	6012: uint16(1),
	6013: uint16(sym__if_end_tag),
	6014: uint16(3),
	6015: uint16(2),
	6016: uint16(sym_comment),
	6017: uint16(aux_sym__tag_value_token1),
	6018: uint16(7),
	6019: uint16(3),
	6020: uint16(1),
	6021: uint16(sym_comment),
	6022: uint16(453),
	6023: uint16(1),
	6024: uint16(aux_sym__single_quoted_attribute_value_token1),
	6025: uint16(455),
	6026: uint16(1),
	6027: uint16(anon_sym_LBRACE),
	6028: uint16(457),
	6029: uint16(1),
	6030: uint16(aux_sym__tag_value_token1),
	6031: uint16(483),
	6032: uint16(1),
	6033: uint16(anon_sym_SQUOTE),
	6034: uint16(171),
	6035: uint16(1),
	6036: uint16(aux_sym__single_quoted_attribute_value),
	6037: uint16(202),
	6038: uint16(1),
	6039: uint16(sym_expression),
	6040: uint16(6),
	6041: uint16(153),
	6042: uint16(1),
	6043: uint16(anon_sym_COLON),
	6044: uint16(155),
	6045: uint16(1),
	6046: uint16(anon_sym_SLASH),
	6047: uint16(284),
	6048: uint16(1),
	6049: uint16(sym__else_if_tag),
	6050: uint16(285),
	6051: uint16(1),
	6052: uint16(sym__else_tag),
	6053: uint16(340),
	6054: uint16(1),
	6055: uint16(sym__if_end_tag),
	6056: uint16(3),
	6057: uint16(2),
	6058: uint16(sym_comment),
	6059: uint16(aux_sym__tag_value_token1),
	6060: uint16(7),
	6061: uint16(3),
	6062: uint16(1),
	6063: uint16(sym_comment),
	6064: uint16(457),
	6065: uint16(1),
	6066: uint16(aux_sym__tag_value_token1),
	6067: uint16(459),
	6068: uint16(1),
	6069: uint16(aux_sym__double_quoted_attribute_value_token1),
	6070: uint16(461),
	6071: uint16(1),
	6072: uint16(anon_sym_LBRACE),
	6073: uint16(483),
	6074: uint16(1),
	6075: uint16(anon_sym_DQUOTE),
	6076: uint16(172),
	6077: uint16(1),
	6078: uint16(aux_sym__double_quoted_attribute_value),
	6079: uint16(208),
	6080: uint16(1),
	6081: uint16(sym_expression),
	6082: uint16(7),
	6083: uint16(3),
	6084: uint16(1),
	6085: uint16(sym_comment),
	6086: uint16(453),
	6087: uint16(1),
	6088: uint16(aux_sym__single_quoted_attribute_value_token1),
	6089: uint16(455),
	6090: uint16(1),
	6091: uint16(anon_sym_LBRACE),
	6092: uint16(457),
	6093: uint16(1),
	6094: uint16(aux_sym__tag_value_token1),
	6095: uint16(485),
	6096: uint16(1),
	6097: uint16(anon_sym_SQUOTE),
	6098: uint16(182),
	6099: uint16(1),
	6100: uint16(aux_sym__single_quoted_attribute_value),
	6101: uint16(202),
	6102: uint16(1),
	6103: uint16(sym_expression),
	6104: uint16(7),
	6105: uint16(3),
	6106: uint16(1),
	6107: uint16(sym_comment),
	6108: uint16(457),
	6109: uint16(1),
	6110: uint16(aux_sym__tag_value_token1),
	6111: uint16(459),
	6112: uint16(1),
	6113: uint16(aux_sym__double_quoted_attribute_value_token1),
	6114: uint16(461),
	6115: uint16(1),
	6116: uint16(anon_sym_LBRACE),
	6117: uint16(485),
	6118: uint16(1),
	6119: uint16(anon_sym_DQUOTE),
	6120: uint16(183),
	6121: uint16(1),
	6122: uint16(aux_sym__double_quoted_attribute_value),
	6123: uint16(208),
	6124: uint16(1),
	6125: uint16(sym_expression),
	6126: uint16(7),
	6127: uint16(3),
	6128: uint16(1),
	6129: uint16(sym_comment),
	6130: uint16(453),
	6131: uint16(1),
	6132: uint16(aux_sym__single_quoted_attribute_value_token1),
	6133: uint16(455),
	6134: uint16(1),
	6135: uint16(anon_sym_LBRACE),
	6136: uint16(457),
	6137: uint16(1),
	6138: uint16(aux_sym__tag_value_token1),
	6139: uint16(487),
	6140: uint16(1),
	6141: uint16(anon_sym_SQUOTE),
	6142: uint16(173),
	6143: uint16(1),
	6144: uint16(aux_sym__single_quoted_attribute_value),
	6145: uint16(202),
	6146: uint16(1),
	6147: uint16(sym_expression),
	6148: uint16(7),
	6149: uint16(3),
	6150: uint16(1),
	6151: uint16(sym_comment),
	6152: uint16(457),
	6153: uint16(1),
	6154: uint16(aux_sym__tag_value_token1),
	6155: uint16(459),
	6156: uint16(1),
	6157: uint16(aux_sym__double_quoted_attribute_value_token1),
	6158: uint16(461),
	6159: uint16(1),
	6160: uint16(anon_sym_LBRACE),
	6161: uint16(487),
	6162: uint16(1),
	6163: uint16(anon_sym_DQUOTE),
	6164: uint16(174),
	6165: uint16(1),
	6166: uint16(aux_sym__double_quoted_attribute_value),
	6167: uint16(208),
	6168: uint16(1),
	6169: uint16(sym_expression),
	6170: uint16(2),
	6171: uint16(3),
	6172: uint16(2),
	6173: uint16(sym_comment),
	6174: uint16(aux_sym__tag_value_token1),
	6175: uint16(199),
	6176: uint16(4),
	6177: uint16(anon_sym_GT),
	6178: uint16(anon_sym_SLASH_GT),
	6179: uint16(sym_attribute_name),
	6180: uint16(anon_sym_LBRACE),
	6181: uint16(5),
	6182: uint16(147),
	6183: uint16(1),
	6184: uint16(anon_sym_SLASH),
	6185: uint16(489),
	6186: uint16(1),
	6187: uint16(anon_sym_COLON),
	6188: uint16(239),
	6189: uint16(1),
	6190: uint16(sym__catch_tag),
	6191: uint16(333),
	6192: uint16(1),
	6193: uint16(sym__await_end_tag),
	6194: uint16(3),
	6195: uint16(2),
	6196: uint16(sym_comment),
	6197: uint16(aux_sym__tag_value_token1),
	6198: uint16(2),
	6199: uint16(3),
	6200: uint16(2),
	6201: uint16(sym_comment),
	6202: uint16(aux_sym__tag_value_token1),
	6203: uint16(491),
	6204: uint16(4),
	6205: uint16(anon_sym_GT),
	6206: uint16(anon_sym_SLASH_GT),
	6207: uint16(sym_attribute_name),
	6208: uint16(anon_sym_LBRACE),
	6209: uint16(4),
	6210: uint16(493),
	6211: uint16(1),
	6212: uint16(anon_sym_LBRACE),
	6213: uint16(37),
	6214: uint16(1),
	6215: uint16(sym_else_if_start),
	6216: uint16(3),
	6217: uint16(2),
	6218: uint16(sym_comment),
	6219: uint16(aux_sym__tag_value_token1),
	6220: uint16(187),
	6221: uint16(2),
	6222: uint16(sym_else_if_block),
	6223: uint16(aux_sym_if_statement_repeat1),
	6224: uint16(2),
	6225: uint16(3),
	6226: uint16(2),
	6227: uint16(sym_comment),
	6228: uint16(aux_sym__tag_value_token1),
	6229: uint16(496),
	6230: uint16(4),
	6231: uint16(anon_sym_GT),
	6232: uint16(anon_sym_SLASH_GT),
	6233: uint16(sym_attribute_name),
	6234: uint16(anon_sym_LBRACE),
	6235: uint16(5),
	6236: uint16(498),
	6237: uint16(1),
	6238: uint16(anon_sym_LBRACE),
	6239: uint16(38),
	6240: uint16(1),
	6241: uint16(sym_catch_start),
	6242: uint16(92),
	6243: uint16(1),
	6244: uint16(sym_await_end),
	6245: uint16(223),
	6246: uint16(1),
	6247: uint16(sym_catch_block),
	6248: uint16(3),
	6249: uint16(2),
	6250: uint16(sym_comment),
	6251: uint16(aux_sym__tag_value_token1),
	6252: uint16(5),
	6253: uint16(498),
	6254: uint16(1),
	6255: uint16(anon_sym_LBRACE),
	6256: uint16(38),
	6257: uint16(1),
	6258: uint16(sym_catch_start),
	6259: uint16(75),
	6260: uint16(1),
	6261: uint16(sym_await_end),
	6262: uint16(243),
	6263: uint16(1),
	6264: uint16(sym_catch_block),
	6265: uint16(3),
	6266: uint16(2),
	6267: uint16(sym_comment),
	6268: uint16(aux_sym__tag_value_token1),
	6269: uint16(5),
	6270: uint16(500),
	6271: uint16(1),
	6272: uint16(anon_sym_html),
	6273: uint16(502),
	6274: uint16(1),
	6275: uint16(anon_sym_const),
	6276: uint16(504),
	6277: uint16(1),
	6278: uint16(anon_sym_debug),
	6279: uint16(506),
	6280: uint16(1),
	6281: uint16(anon_sym_render),
	6282: uint16(3),
	6283: uint16(2),
	6284: uint16(sym_comment),
	6285: uint16(aux_sym__tag_value_token1),
	6286: uint16(5),
	6287: uint16(508),
	6288: uint16(1),
	6289: uint16(anon_sym_LBRACE),
	6290: uint16(38),
	6291: uint16(1),
	6292: uint16(sym_catch_start),
	6293: uint16(117),
	6294: uint16(1),
	6295: uint16(sym_await_end),
	6296: uint16(231),
	6297: uint16(1),
	6298: uint16(sym_catch_block),
	6299: uint16(3),
	6300: uint16(2),
	6301: uint16(sym_comment),
	6302: uint16(aux_sym__tag_value_token1),
	6303: uint16(5),
	6304: uint16(508),
	6305: uint16(1),
	6306: uint16(anon_sym_LBRACE),
	6307: uint16(38),
	6308: uint16(1),
	6309: uint16(sym_catch_start),
	6310: uint16(132),
	6311: uint16(1),
	6312: uint16(sym_await_end),
	6313: uint16(237),
	6314: uint16(1),
	6315: uint16(sym_catch_block),
	6316: uint16(3),
	6317: uint16(2),
	6318: uint16(sym_comment),
	6319: uint16(aux_sym__tag_value_token1),
	6320: uint16(3),
	6321: uint16(510),
	6322: uint16(1),
	6323: uint16(anon_sym_EQ),
	6324: uint16(3),
	6325: uint16(2),
	6326: uint16(sym_comment),
	6327: uint16(aux_sym__tag_value_token1),
	6328: uint16(479),
	6329: uint16(3),
	6330: uint16(anon_sym_GT),
	6331: uint16(sym_attribute_name),
	6332: uint16(anon_sym_LBRACE),
	6333: uint16(5),
	6334: uint16(147),
	6335: uint16(1),
	6336: uint16(anon_sym_SLASH),
	6337: uint16(489),
	6338: uint16(1),
	6339: uint16(anon_sym_COLON),
	6340: uint16(239),
	6341: uint16(1),
	6342: uint16(sym__catch_tag),
	6343: uint16(303),
	6344: uint16(1),
	6345: uint16(sym__await_end_tag),
	6346: uint16(3),
	6347: uint16(2),
	6348: uint16(sym_comment),
	6349: uint16(aux_sym__tag_value_token1),
	6350: uint16(2),
	6351: uint16(3),
	6352: uint16(2),
	6353: uint16(sym_comment),
	6354: uint16(aux_sym__tag_value_token1),
	6355: uint16(512),
	6356: uint16(4),
	6357: uint16(anon_sym_GT),
	6358: uint16(anon_sym_SLASH_GT),
	6359: uint16(sym_attribute_name),
	6360: uint16(anon_sym_LBRACE),
	6361: uint16(2),
	6362: uint16(3),
	6363: uint16(2),
	6364: uint16(sym_comment),
	6365: uint16(aux_sym__tag_value_token1),
	6366: uint16(514),
	6367: uint16(4),
	6368: uint16(anon_sym_GT),
	6369: uint16(anon_sym_SLASH_GT),
	6370: uint16(sym_attribute_name),
	6371: uint16(anon_sym_LBRACE),
	6372: uint16(2),
	6373: uint16(3),
	6374: uint16(2),
	6375: uint16(sym_comment),
	6376: uint16(aux_sym__tag_value_token1),
	6377: uint16(512),
	6378: uint16(3),
	6379: uint16(anon_sym_GT),
	6380: uint16(sym_attribute_name),
	6381: uint16(anon_sym_LBRACE),
	6382: uint16(2),
	6383: uint16(3),
	6384: uint16(2),
	6385: uint16(sym_comment),
	6386: uint16(aux_sym__tag_value_token1),
	6387: uint16(496),
	6388: uint16(3),
	6389: uint16(anon_sym_GT),
	6390: uint16(sym_attribute_name),
	6391: uint16(anon_sym_LBRACE),
	6392: uint16(2),
	6393: uint16(3),
	6394: uint16(2),
	6395: uint16(sym_comment),
	6396: uint16(aux_sym__tag_value_token1),
	6397: uint16(514),
	6398: uint16(3),
	6399: uint16(anon_sym_GT),
	6400: uint16(sym_attribute_name),
	6401: uint16(anon_sym_LBRACE),
	6402: uint16(2),
	6403: uint16(3),
	6404: uint16(2),
	6405: uint16(sym_comment),
	6406: uint16(aux_sym__tag_value_token1),
	6407: uint16(491),
	6408: uint16(3),
	6409: uint16(anon_sym_GT),
	6410: uint16(sym_attribute_name),
	6411: uint16(anon_sym_LBRACE),
	6412: uint16(3),
	6413: uint16(3),
	6414: uint16(1),
	6415: uint16(sym_comment),
	6416: uint16(457),
	6417: uint16(1),
	6418: uint16(aux_sym__tag_value_token1),
	6419: uint16(516),
	6420: uint16(3),
	6421: uint16(anon_sym_SQUOTE),
	6422: uint16(aux_sym__single_quoted_attribute_value_token1),
	6423: uint16(anon_sym_LBRACE),
	6424: uint16(2),
	6425: uint16(3),
	6426: uint16(2),
	6427: uint16(sym_comment),
	6428: uint16(aux_sym__tag_value_token1),
	6429: uint16(199),
	6430: uint16(3),
	6431: uint16(anon_sym_GT),
	6432: uint16(sym_attribute_name),
	6433: uint16(anon_sym_LBRACE),
	6434: uint16(3),
	6435: uint16(3),
	6436: uint16(1),
	6437: uint16(sym_comment),
	6438: uint16(457),
	6439: uint16(1),
	6440: uint16(aux_sym__tag_value_token1),
	6441: uint16(199),
	6442: uint16(3),
	6443: uint16(anon_sym_SQUOTE),
	6444: uint16(aux_sym__single_quoted_attribute_value_token1),
	6445: uint16(anon_sym_LBRACE),
	6446: uint16(3),
	6447: uint16(3),
	6448: uint16(1),
	6449: uint16(sym_comment),
	6450: uint16(457),
	6451: uint16(1),
	6452: uint16(aux_sym__tag_value_token1),
	6453: uint16(199),
	6454: uint16(3),
	6455: uint16(anon_sym_DQUOTE),
	6456: uint16(aux_sym__double_quoted_attribute_value_token1),
	6457: uint16(anon_sym_LBRACE),
	6458: uint16(4),
	6459: uint16(518),
	6460: uint16(1),
	6461: uint16(anon_sym_LT_SLASH),
	6462: uint16(520),
	6463: uint16(1),
	6464: uint16(sym_raw_text),
	6465: uint16(101),
	6466: uint16(1),
	6467: uint16(sym_end_tag),
	6468: uint16(3),
	6469: uint16(2),
	6470: uint16(sym_comment),
	6471: uint16(aux_sym__tag_value_token1),
	6472: uint16(4),
	6473: uint16(522),
	6474: uint16(1),
	6475: uint16(sym__start_tag_name),
	6476: uint16(524),
	6477: uint16(1),
	6478: uint16(sym__script_start_tag_name),
	6479: uint16(526),
	6480: uint16(1),
	6481: uint16(sym__style_start_tag_name),
	6482: uint16(3),
	6483: uint16(2),
	6484: uint16(sym_comment),
	6485: uint16(aux_sym__tag_value_token1),
	6486: uint16(3),
	6487: uint16(3),
	6488: uint16(1),
	6489: uint16(sym_comment),
	6490: uint16(457),
	6491: uint16(1),
	6492: uint16(aux_sym__tag_value_token1),
	6493: uint16(528),
	6494: uint16(3),
	6495: uint16(anon_sym_DQUOTE),
	6496: uint16(aux_sym__double_quoted_attribute_value_token1),
	6497: uint16(anon_sym_LBRACE),
	6498: uint16(4),
	6499: uint16(530),
	6500: uint16(1),
	6501: uint16(anon_sym_LT_SLASH),
	6502: uint16(532),
	6503: uint16(1),
	6504: uint16(sym_raw_text),
	6505: uint16(139),
	6506: uint16(1),
	6507: uint16(sym_end_tag),
	6508: uint16(3),
	6509: uint16(2),
	6510: uint16(sym_comment),
	6511: uint16(aux_sym__tag_value_token1),
	6512: uint16(4),
	6513: uint16(530),
	6514: uint16(1),
	6515: uint16(anon_sym_LT_SLASH),
	6516: uint16(534),
	6517: uint16(1),
	6518: uint16(sym_raw_text),
	6519: uint16(138),
	6520: uint16(1),
	6521: uint16(sym_end_tag),
	6522: uint16(3),
	6523: uint16(2),
	6524: uint16(sym_comment),
	6525: uint16(aux_sym__tag_value_token1),
	6526: uint16(4),
	6527: uint16(524),
	6528: uint16(1),
	6529: uint16(sym__script_start_tag_name),
	6530: uint16(526),
	6531: uint16(1),
	6532: uint16(sym__style_start_tag_name),
	6533: uint16(536),
	6534: uint16(1),
	6535: uint16(sym__start_tag_name),
	6536: uint16(3),
	6537: uint16(2),
	6538: uint16(sym_comment),
	6539: uint16(aux_sym__tag_value_token1),
	6540: uint16(4),
	6541: uint16(518),
	6542: uint16(1),
	6543: uint16(anon_sym_LT_SLASH),
	6544: uint16(538),
	6545: uint16(1),
	6546: uint16(sym_raw_text),
	6547: uint16(102),
	6548: uint16(1),
	6549: uint16(sym_end_tag),
	6550: uint16(3),
	6551: uint16(2),
	6552: uint16(sym_comment),
	6553: uint16(aux_sym__tag_value_token1),
	6554: uint16(3),
	6555: uint16(540),
	6556: uint16(1),
	6557: uint16(anon_sym_LBRACE),
	6558: uint16(117),
	6559: uint16(1),
	6560: uint16(sym_await_end),
	6561: uint16(3),
	6562: uint16(2),
	6563: uint16(sym_comment),
	6564: uint16(aux_sym__tag_value_token1),
	6565: uint16(3),
	6566: uint16(518),
	6567: uint16(1),
	6568: uint16(anon_sym_LT_SLASH),
	6569: uint16(113),
	6570: uint16(1),
	6571: uint16(sym_end_tag),
	6572: uint16(3),
	6573: uint16(2),
	6574: uint16(sym_comment),
	6575: uint16(aux_sym__tag_value_token1),
	6576: uint16(3),
	6577: uint16(518),
	6578: uint16(1),
	6579: uint16(anon_sym_LT_SLASH),
	6580: uint16(114),
	6581: uint16(1),
	6582: uint16(sym_end_tag),
	6583: uint16(3),
	6584: uint16(2),
	6585: uint16(sym_comment),
	6586: uint16(aux_sym__tag_value_token1),
	6587: uint16(3),
	6588: uint16(542),
	6589: uint16(1),
	6590: uint16(anon_sym_LBRACE),
	6591: uint16(115),
	6592: uint16(1),
	6593: uint16(sym_if_end),
	6594: uint16(3),
	6595: uint16(2),
	6596: uint16(sym_comment),
	6597: uint16(aux_sym__tag_value_token1),
	6598: uint16(3),
	6599: uint16(544),
	6600: uint16(1),
	6601: uint16(anon_sym_LBRACE),
	6602: uint16(72),
	6603: uint16(1),
	6604: uint16(sym_if_end),
	6605: uint16(3),
	6606: uint16(2),
	6607: uint16(sym_comment),
	6608: uint16(aux_sym__tag_value_token1),
	6609: uint16(3),
	6610: uint16(147),
	6611: uint16(1),
	6612: uint16(anon_sym_SLASH),
	6613: uint16(333),
	6614: uint16(1),
	6615: uint16(sym__await_end_tag),
	6616: uint16(3),
	6617: uint16(2),
	6618: uint16(sym_comment),
	6619: uint16(aux_sym__tag_value_token1),
	6620: uint16(3),
	6621: uint16(155),
	6622: uint16(1),
	6623: uint16(anon_sym_SLASH),
	6624: uint16(286),
	6625: uint16(1),
	6626: uint16(sym__if_end_tag),
	6627: uint16(3),
	6628: uint16(2),
	6629: uint16(sym_comment),
	6630: uint16(aux_sym__tag_value_token1),
	6631: uint16(3),
	6632: uint16(530),
	6633: uint16(1),
	6634: uint16(anon_sym_LT_SLASH),
	6635: uint16(70),
	6636: uint16(1),
	6637: uint16(sym_end_tag),
	6638: uint16(3),
	6639: uint16(2),
	6640: uint16(sym_comment),
	6641: uint16(aux_sym__tag_value_token1),
	6642: uint16(2),
	6643: uint16(3),
	6644: uint16(2),
	6645: uint16(sym_comment),
	6646: uint16(aux_sym__tag_value_token1),
	6647: uint16(546),
	6648: uint16(2),
	6649: uint16(sym_svelte_raw_text),
	6650: uint16(sym_svelte_raw_text_each),
	6651: uint16(3),
	6652: uint16(548),
	6653: uint16(1),
	6654: uint16(sym_svelte_raw_text),
	6655: uint16(550),
	6656: uint16(1),
	6657: uint16(sym_svelte_raw_text_each),
	6658: uint16(3),
	6659: uint16(2),
	6660: uint16(sym_comment),
	6661: uint16(aux_sym__tag_value_token1),
	6662: uint16(3),
	6663: uint16(552),
	6664: uint16(1),
	6665: uint16(anon_sym_LBRACE),
	6666: uint16(97),
	6667: uint16(1),
	6668: uint16(sym_await_end),
	6669: uint16(3),
	6670: uint16(2),
	6671: uint16(sym_comment),
	6672: uint16(aux_sym__tag_value_token1),
	6673: uint16(3),
	6674: uint16(155),
	6675: uint16(1),
	6676: uint16(anon_sym_SLASH),
	6677: uint16(340),
	6678: uint16(1),
	6679: uint16(sym__if_end_tag),
	6680: uint16(3),
	6681: uint16(2),
	6682: uint16(sym_comment),
	6683: uint16(aux_sym__tag_value_token1),
	6684: uint16(3),
	6685: uint16(542),
	6686: uint16(1),
	6687: uint16(anon_sym_LBRACE),
	6688: uint16(128),
	6689: uint16(1),
	6690: uint16(sym_if_end),
	6691: uint16(3),
	6692: uint16(2),
	6693: uint16(sym_comment),
	6694: uint16(aux_sym__tag_value_token1),
	6695: uint16(2),
	6696: uint16(3),
	6697: uint16(2),
	6698: uint16(sym_comment),
	6699: uint16(aux_sym__tag_value_token1),
	6700: uint16(554),
	6701: uint16(2),
	6702: uint16(sym_raw_text),
	6703: uint16(anon_sym_LT_SLASH),
	6704: uint16(2),
	6705: uint16(3),
	6706: uint16(2),
	6707: uint16(sym_comment),
	6708: uint16(aux_sym__tag_value_token1),
	6709: uint16(556),
	6710: uint16(2),
	6711: uint16(sym_svelte_raw_text),
	6712: uint16(anon_sym_RBRACE),
	6713: uint16(2),
	6714: uint16(3),
	6715: uint16(2),
	6716: uint16(sym_comment),
	6717: uint16(aux_sym__tag_value_token1),
	6718: uint16(558),
	6719: uint16(2),
	6720: uint16(sym_raw_text),
	6721: uint16(anon_sym_LT_SLASH),
	6722: uint16(2),
	6723: uint16(3),
	6724: uint16(2),
	6725: uint16(sym_comment),
	6726: uint16(aux_sym__tag_value_token1),
	6727: uint16(560),
	6728: uint16(2),
	6729: uint16(sym_raw_text),
	6730: uint16(anon_sym_LT_SLASH),
	6731: uint16(3),
	6732: uint16(544),
	6733: uint16(1),
	6734: uint16(anon_sym_LBRACE),
	6735: uint16(88),
	6736: uint16(1),
	6737: uint16(sym_if_end),
	6738: uint16(3),
	6739: uint16(2),
	6740: uint16(sym_comment),
	6741: uint16(aux_sym__tag_value_token1),
	6742: uint16(3),
	6743: uint16(540),
	6744: uint16(1),
	6745: uint16(anon_sym_LBRACE),
	6746: uint16(132),
	6747: uint16(1),
	6748: uint16(sym_await_end),
	6749: uint16(3),
	6750: uint16(2),
	6751: uint16(sym_comment),
	6752: uint16(aux_sym__tag_value_token1),
	6753: uint16(3),
	6754: uint16(147),
	6755: uint16(1),
	6756: uint16(anon_sym_SLASH),
	6757: uint16(303),
	6758: uint16(1),
	6759: uint16(sym__await_end_tag),
	6760: uint16(3),
	6761: uint16(2),
	6762: uint16(sym_comment),
	6763: uint16(aux_sym__tag_value_token1),
	6764: uint16(3),
	6765: uint16(562),
	6766: uint16(1),
	6767: uint16(anon_sym_then),
	6768: uint16(564),
	6769: uint16(1),
	6770: uint16(anon_sym_catch),
	6771: uint16(3),
	6772: uint16(2),
	6773: uint16(sym_comment),
	6774: uint16(aux_sym__tag_value_token1),
	6775: uint16(3),
	6776: uint16(566),
	6777: uint16(1),
	6778: uint16(anon_sym_COLON),
	6779: uint16(284),
	6780: uint16(1),
	6781: uint16(sym__else_if_tag),
	6782: uint16(3),
	6783: uint16(2),
	6784: uint16(sym_comment),
	6785: uint16(aux_sym__tag_value_token1),
	6786: uint16(3),
	6787: uint16(542),
	6788: uint16(1),
	6789: uint16(anon_sym_LBRACE),
	6790: uint16(135),
	6791: uint16(1),
	6792: uint16(sym_if_end),
	6793: uint16(3),
	6794: uint16(2),
	6795: uint16(sym_comment),
	6796: uint16(aux_sym__tag_value_token1),
	6797: uint16(3),
	6798: uint16(568),
	6799: uint16(1),
	6800: uint16(anon_sym_RBRACE),
	6801: uint16(570),
	6802: uint16(1),
	6803: uint16(sym_svelte_raw_text),
	6804: uint16(3),
	6805: uint16(2),
	6806: uint16(sym_comment),
	6807: uint16(aux_sym__tag_value_token1),
	6808: uint16(3),
	6809: uint16(540),
	6810: uint16(1),
	6811: uint16(anon_sym_LBRACE),
	6812: uint16(137),
	6813: uint16(1),
	6814: uint16(sym_await_end),
	6815: uint16(3),
	6816: uint16(2),
	6817: uint16(sym_comment),
	6818: uint16(aux_sym__tag_value_token1),
	6819: uint16(3),
	6820: uint16(572),
	6821: uint16(1),
	6822: uint16(anon_sym_RPAREN),
	6823: uint16(574),
	6824: uint16(1),
	6825: uint16(sym_svelte_raw_text_snippet_arguments),
	6826: uint16(3),
	6827: uint16(2),
	6828: uint16(sym_comment),
	6829: uint16(aux_sym__tag_value_token1),
	6830: uint16(3),
	6831: uint16(576),
	6832: uint16(1),
	6833: uint16(anon_sym_RBRACE),
	6834: uint16(578),
	6835: uint16(1),
	6836: uint16(sym_svelte_raw_text),
	6837: uint16(3),
	6838: uint16(2),
	6839: uint16(sym_comment),
	6840: uint16(aux_sym__tag_value_token1),
	6841: uint16(3),
	6842: uint16(580),
	6843: uint16(1),
	6844: uint16(sym__end_tag_name),
	6845: uint16(582),
	6846: uint16(1),
	6847: uint16(sym_erroneous_end_tag_name),
	6848: uint16(3),
	6849: uint16(2),
	6850: uint16(sym_comment),
	6851: uint16(aux_sym__tag_value_token1),
	6852: uint16(3),
	6853: uint16(584),
	6854: uint16(1),
	6855: uint16(anon_sym_elseif),
	6856: uint16(586),
	6857: uint16(1),
	6858: uint16(anon_sym_else),
	6859: uint16(3),
	6860: uint16(2),
	6861: uint16(sym_comment),
	6862: uint16(aux_sym__tag_value_token1),
	6863: uint16(2),
	6864: uint16(3),
	6865: uint16(2),
	6866: uint16(sym_comment),
	6867: uint16(aux_sym__tag_value_token1),
	6868: uint16(588),
	6869: uint16(2),
	6870: uint16(sym_raw_text),
	6871: uint16(anon_sym_LT_SLASH),
	6872: uint16(3),
	6873: uint16(552),
	6874: uint16(1),
	6875: uint16(anon_sym_LBRACE),
	6876: uint16(92),
	6877: uint16(1),
	6878: uint16(sym_await_end),
	6879: uint16(3),
	6880: uint16(2),
	6881: uint16(sym_comment),
	6882: uint16(aux_sym__tag_value_token1),
	6883: uint16(3),
	6884: uint16(582),
	6885: uint16(1),
	6886: uint16(sym_erroneous_end_tag_name),
	6887: uint16(590),
	6888: uint16(1),
	6889: uint16(sym__end_tag_name),
	6890: uint16(3),
	6891: uint16(2),
	6892: uint16(sym_comment),
	6893: uint16(aux_sym__tag_value_token1),
	6894: uint16(2),
	6895: uint16(3),
	6896: uint16(2),
	6897: uint16(sym_comment),
	6898: uint16(aux_sym__tag_value_token1),
	6899: uint16(592),
	6900: uint16(2),
	6901: uint16(sym_svelte_raw_text),
	6902: uint16(anon_sym_RBRACE),
	6903: uint16(3),
	6904: uint16(552),
	6905: uint16(1),
	6906: uint16(anon_sym_LBRACE),
	6907: uint16(75),
	6908: uint16(1),
	6909: uint16(sym_await_end),
	6910: uint16(3),
	6911: uint16(2),
	6912: uint16(sym_comment),
	6913: uint16(aux_sym__tag_value_token1),
	6914: uint16(3),
	6915: uint16(530),
	6916: uint16(1),
	6917: uint16(anon_sym_LT_SLASH),
	6918: uint16(71),
	6919: uint16(1),
	6920: uint16(sym_end_tag),
	6921: uint16(3),
	6922: uint16(2),
	6923: uint16(sym_comment),
	6924: uint16(aux_sym__tag_value_token1),
	6925: uint16(3),
	6926: uint16(544),
	6927: uint16(1),
	6928: uint16(anon_sym_LBRACE),
	6929: uint16(95),
	6930: uint16(1),
	6931: uint16(sym_if_end),
	6932: uint16(3),
	6933: uint16(2),
	6934: uint16(sym_comment),
	6935: uint16(aux_sym__tag_value_token1),
	6936: uint16(2),
	6937: uint16(594),
	6938: uint16(1),
	6939: uint16(anon_sym_RBRACE),
	6940: uint16(3),
	6941: uint16(2),
	6942: uint16(sym_comment),
	6943: uint16(aux_sym__tag_value_token1),
	6944: uint16(2),
	6945: uint16(596),
	6946: uint16(1),
	6947: uint16(sym_svelte_raw_text),
	6948: uint16(3),
	6949: uint16(2),
	6950: uint16(sym_comment),
	6951: uint16(aux_sym__tag_value_token1),
	6952: uint16(2),
	6953: uint16(598),
	6954: uint16(1),
	6955: uint16(aux_sym_snippet_start_token1),
	6956: uint16(3),
	6957: uint16(2),
	6958: uint16(sym_comment),
	6959: uint16(aux_sym__tag_value_token1),
	6960: uint16(2),
	6961: uint16(600),
	6962: uint16(1),
	6963: uint16(anon_sym_RBRACE),
	6964: uint16(3),
	6965: uint16(2),
	6966: uint16(sym_comment),
	6967: uint16(aux_sym__tag_value_token1),
	6968: uint16(3),
	6969: uint16(3),
	6970: uint16(1),
	6971: uint16(sym_comment),
	6972: uint16(602),
	6973: uint16(1),
	6974: uint16(aux_sym__tag_value_token1),
	6975: uint16(270),
	6976: uint16(1),
	6977: uint16(sym__tag_value),
	6978: uint16(2),
	6979: uint16(604),
	6980: uint16(1),
	6981: uint16(anon_sym_RBRACE),
	6982: uint16(3),
	6983: uint16(2),
	6984: uint16(sym_comment),
	6985: uint16(aux_sym__tag_value_token1),
	6986: uint16(2),
	6987: uint16(606),
	6988: uint16(1),
	6989: uint16(anon_sym_RBRACE),
	6990: uint16(3),
	6991: uint16(2),
	6992: uint16(sym_comment),
	6993: uint16(aux_sym__tag_value_token1),
	6994: uint16(2),
	6995: uint16(608),
	6996: uint16(1),
	6997: uint16(anon_sym_as),
	6998: uint16(3),
	6999: uint16(2),
	7000: uint16(sym_comment),
	7001: uint16(aux_sym__tag_value_token1),
	7002: uint16(2),
	7003: uint16(610),
	7004: uint16(1),
	7005: uint16(anon_sym_key),
	7006: uint16(3),
	7007: uint16(2),
	7008: uint16(sym_comment),
	7009: uint16(aux_sym__tag_value_token1),
	7010: uint16(2),
	7011: uint16(612),
	7012: uint16(1),
	7013: uint16(anon_sym_RBRACE),
	7014: uint16(3),
	7015: uint16(2),
	7016: uint16(sym_comment),
	7017: uint16(aux_sym__tag_value_token1),
	7018: uint16(2),
	7019: uint16(614),
	7020: uint16(1),
	7021: uint16(anon_sym_GT),
	7022: uint16(3),
	7023: uint16(2),
	7024: uint16(sym_comment),
	7025: uint16(aux_sym__tag_value_token1),
	7026: uint16(2),
	7027: uint16(616),
	7028: uint16(1),
	7029: uint16(anon_sym_RBRACE),
	7030: uint16(3),
	7031: uint16(2),
	7032: uint16(sym_comment),
	7033: uint16(aux_sym__tag_value_token1),
	7034: uint16(2),
	7035: uint16(618),
	7036: uint16(1),
	7037: uint16(anon_sym_RBRACE),
	7038: uint16(3),
	7039: uint16(2),
	7040: uint16(sym_comment),
	7041: uint16(aux_sym__tag_value_token1),
	7042: uint16(2),
	7043: uint16(620),
	7044: uint16(1),
	7045: uint16(anon_sym_RBRACE),
	7046: uint16(3),
	7047: uint16(2),
	7048: uint16(sym_comment),
	7049: uint16(aux_sym__tag_value_token1),
	7050: uint16(2),
	7051: uint16(580),
	7052: uint16(1),
	7053: uint16(sym__end_tag_name),
	7054: uint16(3),
	7055: uint16(2),
	7056: uint16(sym_comment),
	7057: uint16(aux_sym__tag_value_token1),
	7058: uint16(2),
	7059: uint16(622),
	7060: uint16(1),
	7061: uint16(sym_svelte_raw_text),
	7062: uint16(3),
	7063: uint16(2),
	7064: uint16(sym_comment),
	7065: uint16(aux_sym__tag_value_token1),
	7066: uint16(2),
	7067: uint16(624),
	7068: uint16(1),
	7069: uint16(anon_sym_RBRACE),
	7070: uint16(3),
	7071: uint16(2),
	7072: uint16(sym_comment),
	7073: uint16(aux_sym__tag_value_token1),
	7074: uint16(2),
	7075: uint16(626),
	7076: uint16(1),
	7077: uint16(anon_sym_snippet),
	7078: uint16(3),
	7079: uint16(2),
	7080: uint16(sym_comment),
	7081: uint16(aux_sym__tag_value_token1),
	7082: uint16(2),
	7083: uint16(628),
	7084: uint16(1),
	7085: uint16(anon_sym_RBRACE),
	7086: uint16(3),
	7087: uint16(2),
	7088: uint16(sym_comment),
	7089: uint16(aux_sym__tag_value_token1),
	7090: uint16(2),
	7091: uint16(630),
	7092: uint16(1),
	7093: uint16(anon_sym_RBRACE),
	7094: uint16(3),
	7095: uint16(2),
	7096: uint16(sym_comment),
	7097: uint16(aux_sym__tag_value_token1),
	7098: uint16(2),
	7099: uint16(632),
	7100: uint16(1),
	7101: uint16(sym_erroneous_end_tag_name),
	7102: uint16(3),
	7103: uint16(2),
	7104: uint16(sym_comment),
	7105: uint16(aux_sym__tag_value_token1),
	7106: uint16(2),
	7107: uint16(634),
	7108: uint16(1),
	7109: uint16(anon_sym_RBRACE),
	7110: uint16(3),
	7111: uint16(2),
	7112: uint16(sym_comment),
	7113: uint16(aux_sym__tag_value_token1),
	7114: uint16(2),
	7115: uint16(636),
	7116: uint16(1),
	7117: uint16(sym_svelte_raw_text),
	7118: uint16(3),
	7119: uint16(2),
	7120: uint16(sym_comment),
	7121: uint16(aux_sym__tag_value_token1),
	7122: uint16(2),
	7123: uint16(638),
	7124: uint16(1),
	7125: uint16(anon_sym_GT),
	7126: uint16(3),
	7127: uint16(2),
	7128: uint16(sym_comment),
	7129: uint16(aux_sym__tag_value_token1),
	7130: uint16(2),
	7131: uint16(640),
	7132: uint16(1),
	7133: uint16(anon_sym_RBRACE),
	7134: uint16(3),
	7135: uint16(2),
	7136: uint16(sym_comment),
	7137: uint16(aux_sym__tag_value_token1),
	7138: uint16(2),
	7139: uint16(642),
	7140: uint16(1),
	7141: uint16(sym__doctype),
	7142: uint16(3),
	7143: uint16(2),
	7144: uint16(sym_comment),
	7145: uint16(aux_sym__tag_value_token1),
	7146: uint16(2),
	7147: uint16(644),
	7148: uint16(1),
	7149: uint16(anon_sym_GT),
	7150: uint16(3),
	7151: uint16(2),
	7152: uint16(sym_comment),
	7153: uint16(aux_sym__tag_value_token1),
	7154: uint16(2),
	7155: uint16(646),
	7156: uint16(1),
	7157: uint16(anon_sym_RBRACE),
	7158: uint16(3),
	7159: uint16(2),
	7160: uint16(sym_comment),
	7161: uint16(aux_sym__tag_value_token1),
	7162: uint16(2),
	7163: uint16(648),
	7164: uint16(1),
	7165: uint16(sym_svelte_raw_text),
	7166: uint16(3),
	7167: uint16(2),
	7168: uint16(sym_comment),
	7169: uint16(aux_sym__tag_value_token1),
	7170: uint16(2),
	7171: uint16(564),
	7172: uint16(1),
	7173: uint16(anon_sym_catch),
	7174: uint16(3),
	7175: uint16(2),
	7176: uint16(sym_comment),
	7177: uint16(aux_sym__tag_value_token1),
	7178: uint16(2),
	7179: uint16(650),
	7180: uint16(1),
	7181: uint16(anon_sym_RBRACE),
	7182: uint16(3),
	7183: uint16(2),
	7184: uint16(sym_comment),
	7185: uint16(aux_sym__tag_value_token1),
	7186: uint16(2),
	7187: uint16(652),
	7188: uint16(1),
	7189: uint16(sym_svelte_raw_text),
	7190: uint16(3),
	7191: uint16(2),
	7192: uint16(sym_comment),
	7193: uint16(aux_sym__tag_value_token1),
	7194: uint16(2),
	7195: uint16(654),
	7196: uint16(1),
	7198: uint16(3),
	7199: uint16(2),
	7200: uint16(sym_comment),
	7201: uint16(aux_sym__tag_value_token1),
	7202: uint16(2),
	7203: uint16(656),
	7204: uint16(1),
	7205: uint16(anon_sym_RBRACE),
	7206: uint16(3),
	7207: uint16(2),
	7208: uint16(sym_comment),
	7209: uint16(aux_sym__tag_value_token1),
	7210: uint16(2),
	7211: uint16(658),
	7212: uint16(1),
	7213: uint16(anon_sym_if),
	7214: uint16(3),
	7215: uint16(2),
	7216: uint16(sym_comment),
	7217: uint16(aux_sym__tag_value_token1),
	7218: uint16(2),
	7219: uint16(660),
	7220: uint16(1),
	7221: uint16(sym_svelte_raw_text),
	7222: uint16(3),
	7223: uint16(2),
	7224: uint16(sym_comment),
	7225: uint16(aux_sym__tag_value_token1),
	7226: uint16(2),
	7227: uint16(662),
	7228: uint16(1),
	7229: uint16(anon_sym_RBRACE),
	7230: uint16(3),
	7231: uint16(2),
	7232: uint16(sym_comment),
	7233: uint16(aux_sym__tag_value_token1),
	7234: uint16(2),
	7235: uint16(664),
	7236: uint16(1),
	7237: uint16(anon_sym_RBRACE),
	7238: uint16(3),
	7239: uint16(2),
	7240: uint16(sym_comment),
	7241: uint16(aux_sym__tag_value_token1),
	7242: uint16(2),
	7243: uint16(666),
	7244: uint16(1),
	7245: uint16(anon_sym_GT),
	7246: uint16(3),
	7247: uint16(2),
	7248: uint16(sym_comment),
	7249: uint16(aux_sym__tag_value_token1),
	7250: uint16(2),
	7251: uint16(668),
	7252: uint16(1),
	7253: uint16(sym_svelte_raw_text),
	7254: uint16(3),
	7255: uint16(2),
	7256: uint16(sym_comment),
	7257: uint16(aux_sym__tag_value_token1),
	7258: uint16(2),
	7259: uint16(670),
	7260: uint16(1),
	7261: uint16(anon_sym_RBRACE),
	7262: uint16(3),
	7263: uint16(2),
	7264: uint16(sym_comment),
	7265: uint16(aux_sym__tag_value_token1),
	7266: uint16(2),
	7267: uint16(672),
	7268: uint16(1),
	7269: uint16(anon_sym_RBRACE),
	7270: uint16(3),
	7271: uint16(2),
	7272: uint16(sym_comment),
	7273: uint16(aux_sym__tag_value_token1),
	7274: uint16(2),
	7275: uint16(674),
	7276: uint16(1),
	7277: uint16(anon_sym_RBRACE),
	7278: uint16(3),
	7279: uint16(2),
	7280: uint16(sym_comment),
	7281: uint16(aux_sym__tag_value_token1),
	7282: uint16(2),
	7283: uint16(676),
	7284: uint16(1),
	7285: uint16(anon_sym_RBRACE),
	7286: uint16(3),
	7287: uint16(2),
	7288: uint16(sym_comment),
	7289: uint16(aux_sym__tag_value_token1),
	7290: uint16(2),
	7291: uint16(678),
	7292: uint16(1),
	7293: uint16(anon_sym_GT),
	7294: uint16(3),
	7295: uint16(2),
	7296: uint16(sym_comment),
	7297: uint16(aux_sym__tag_value_token1),
	7298: uint16(2),
	7299: uint16(680),
	7300: uint16(1),
	7301: uint16(sym_svelte_raw_text),
	7302: uint16(3),
	7303: uint16(2),
	7304: uint16(sym_comment),
	7305: uint16(aux_sym__tag_value_token1),
	7306: uint16(2),
	7307: uint16(682),
	7308: uint16(1),
	7309: uint16(sym_svelte_raw_text),
	7310: uint16(3),
	7311: uint16(2),
	7312: uint16(sym_comment),
	7313: uint16(aux_sym__tag_value_token1),
	7314: uint16(2),
	7315: uint16(684),
	7316: uint16(1),
	7317: uint16(sym_svelte_raw_text),
	7318: uint16(3),
	7319: uint16(2),
	7320: uint16(sym_comment),
	7321: uint16(aux_sym__tag_value_token1),
	7322: uint16(2),
	7323: uint16(686),
	7324: uint16(1),
	7325: uint16(anon_sym_RBRACE),
	7326: uint16(3),
	7327: uint16(2),
	7328: uint16(sym_comment),
	7329: uint16(aux_sym__tag_value_token1),
	7330: uint16(2),
	7331: uint16(688),
	7332: uint16(1),
	7333: uint16(anon_sym_LPAREN),
	7334: uint16(3),
	7335: uint16(2),
	7336: uint16(sym_comment),
	7337: uint16(aux_sym__tag_value_token1),
	7338: uint16(2),
	7339: uint16(690),
	7340: uint16(1),
	7341: uint16(anon_sym_RBRACE),
	7342: uint16(3),
	7343: uint16(2),
	7344: uint16(sym_comment),
	7345: uint16(aux_sym__tag_value_token1),
	7346: uint16(2),
	7347: uint16(692),
	7348: uint16(1),
	7349: uint16(anon_sym_RBRACE),
	7350: uint16(3),
	7351: uint16(2),
	7352: uint16(sym_comment),
	7353: uint16(aux_sym__tag_value_token1),
	7354: uint16(3),
	7355: uint16(3),
	7356: uint16(1),
	7357: uint16(sym_comment),
	7358: uint16(457),
	7359: uint16(1),
	7360: uint16(aux_sym__tag_value_token1),
	7361: uint16(694),
	7362: uint16(1),
	7363: uint16(aux_sym_doctype_token1),
	7364: uint16(2),
	7365: uint16(696),
	7366: uint16(1),
	7367: uint16(sym_svelte_raw_text),
	7368: uint16(3),
	7369: uint16(2),
	7370: uint16(sym_comment),
	7371: uint16(aux_sym__tag_value_token1),
	7372: uint16(2),
	7373: uint16(698),
	7374: uint16(1),
	7375: uint16(anon_sym_RBRACE),
	7376: uint16(3),
	7377: uint16(2),
	7378: uint16(sym_comment),
	7379: uint16(aux_sym__tag_value_token1),
	7380: uint16(2),
	7381: uint16(700),
	7382: uint16(1),
	7383: uint16(anon_sym_RBRACE),
	7384: uint16(3),
	7385: uint16(2),
	7386: uint16(sym_comment),
	7387: uint16(aux_sym__tag_value_token1),
	7388: uint16(2),
	7389: uint16(702),
	7390: uint16(1),
	7391: uint16(anon_sym_RBRACE),
	7392: uint16(3),
	7393: uint16(2),
	7394: uint16(sym_comment),
	7395: uint16(aux_sym__tag_value_token1),
	7396: uint16(2),
	7397: uint16(704),
	7398: uint16(1),
	7399: uint16(anon_sym_RBRACE),
	7400: uint16(3),
	7401: uint16(2),
	7402: uint16(sym_comment),
	7403: uint16(aux_sym__tag_value_token1),
	7404: uint16(2),
	7405: uint16(706),
	7406: uint16(1),
	7407: uint16(anon_sym_RBRACE),
	7408: uint16(3),
	7409: uint16(2),
	7410: uint16(sym_comment),
	7411: uint16(aux_sym__tag_value_token1),
	7412: uint16(2),
	7413: uint16(708),
	7414: uint16(1),
	7415: uint16(anon_sym_RBRACE),
	7416: uint16(3),
	7417: uint16(2),
	7418: uint16(sym_comment),
	7419: uint16(aux_sym__tag_value_token1),
	7420: uint16(2),
	7421: uint16(710),
	7422: uint16(1),
	7423: uint16(anon_sym_RPAREN),
	7424: uint16(3),
	7425: uint16(2),
	7426: uint16(sym_comment),
	7427: uint16(aux_sym__tag_value_token1),
	7428: uint16(2),
	7429: uint16(712),
	7430: uint16(1),
	7431: uint16(anon_sym_else),
	7432: uint16(3),
	7433: uint16(2),
	7434: uint16(sym_comment),
	7435: uint16(aux_sym__tag_value_token1),
	7436: uint16(2),
	7437: uint16(714),
	7438: uint16(1),
	7439: uint16(anon_sym_each),
	7440: uint16(3),
	7441: uint16(2),
	7442: uint16(sym_comment),
	7443: uint16(aux_sym__tag_value_token1),
	7444: uint16(2),
	7445: uint16(716),
	7446: uint16(1),
	7447: uint16(anon_sym_RBRACE),
	7448: uint16(3),
	7449: uint16(2),
	7450: uint16(sym_comment),
	7451: uint16(aux_sym__tag_value_token1),
	7452: uint16(2),
	7453: uint16(718),
	7454: uint16(1),
	7455: uint16(aux_sym_snippet_start_token1),
	7456: uint16(3),
	7457: uint16(2),
	7458: uint16(sym_comment),
	7459: uint16(aux_sym__tag_value_token1),
	7460: uint16(2),
	7461: uint16(720),
	7462: uint16(1),
	7463: uint16(anon_sym_RBRACE),
	7464: uint16(3),
	7465: uint16(2),
	7466: uint16(sym_comment),
	7467: uint16(aux_sym__tag_value_token1),
	7468: uint16(3),
	7469: uint16(3),
	7470: uint16(1),
	7471: uint16(sym_comment),
	7472: uint16(602),
	7473: uint16(1),
	7474: uint16(aux_sym__tag_value_token1),
	7475: uint16(323),
	7476: uint16(1),
	7477: uint16(sym__tag_value),
	7478: uint16(2),
	7479: uint16(722),
	7480: uint16(1),
	7481: uint16(anon_sym_RBRACE),
	7482: uint16(3),
	7483: uint16(2),
	7484: uint16(sym_comment),
	7485: uint16(aux_sym__tag_value_token1),
	7486: uint16(2),
	7487: uint16(724),
	7488: uint16(1),
	7489: uint16(anon_sym_RBRACE),
	7490: uint16(3),
	7491: uint16(2),
	7492: uint16(sym_comment),
	7493: uint16(aux_sym__tag_value_token1),
	7494: uint16(2),
	7495: uint16(726),
	7496: uint16(1),
	7497: uint16(anon_sym_RBRACE),
	7498: uint16(3),
	7499: uint16(2),
	7500: uint16(sym_comment),
	7501: uint16(aux_sym__tag_value_token1),
	7502: uint16(2),
	7503: uint16(728),
	7504: uint16(1),
	7505: uint16(anon_sym_RBRACE),
	7506: uint16(3),
	7507: uint16(2),
	7508: uint16(sym_comment),
	7509: uint16(aux_sym__tag_value_token1),
	7510: uint16(2),
	7511: uint16(730),
	7512: uint16(1),
	7513: uint16(anon_sym_RBRACE),
	7514: uint16(3),
	7515: uint16(2),
	7516: uint16(sym_comment),
	7517: uint16(aux_sym__tag_value_token1),
	7518: uint16(3),
	7519: uint16(3),
	7520: uint16(1),
	7521: uint16(sym_comment),
	7522: uint16(602),
	7523: uint16(1),
	7524: uint16(aux_sym__tag_value_token1),
	7525: uint16(332),
	7526: uint16(1),
	7527: uint16(sym__tag_value),
	7528: uint16(2),
	7529: uint16(582),
	7530: uint16(1),
	7531: uint16(sym_erroneous_end_tag_name),
	7532: uint16(3),
	7533: uint16(2),
	7534: uint16(sym_comment),
	7535: uint16(aux_sym__tag_value_token1),
	7536: uint16(2),
	7537: uint16(732),
	7538: uint16(1),
	7539: uint16(anon_sym_RBRACE),
	7540: uint16(3),
	7541: uint16(2),
	7542: uint16(sym_comment),
	7543: uint16(aux_sym__tag_value_token1),
	7544: uint16(3),
	7545: uint16(3),
	7546: uint16(1),
	7547: uint16(sym_comment),
	7548: uint16(457),
	7549: uint16(1),
	7550: uint16(aux_sym__tag_value_token1),
	7551: uint16(734),
	7552: uint16(1),
	7553: uint16(aux_sym_doctype_token1),
	7554: uint16(3),
	7555: uint16(3),
	7556: uint16(1),
	7557: uint16(sym_comment),
	7558: uint16(602),
	7559: uint16(1),
	7560: uint16(aux_sym__tag_value_token1),
	7561: uint16(289),
	7562: uint16(1),
	7563: uint16(sym__tag_value),
	7564: uint16(3),
	7565: uint16(3),
	7566: uint16(1),
	7567: uint16(sym_comment),
	7568: uint16(602),
	7569: uint16(1),
	7570: uint16(aux_sym__tag_value_token1),
	7571: uint16(290),
	7572: uint16(1),
	7573: uint16(sym__tag_value),
	7574: uint16(3),
	7575: uint16(3),
	7576: uint16(1),
	7577: uint16(sym_comment),
	7578: uint16(602),
	7579: uint16(1),
	7580: uint16(aux_sym__tag_value_token1),
	7581: uint16(291),
	7582: uint16(1),
	7583: uint16(sym__tag_value),
	7584: uint16(3),
	7585: uint16(3),
	7586: uint16(1),
	7587: uint16(sym_comment),
	7588: uint16(602),
	7589: uint16(1),
	7590: uint16(aux_sym__tag_value_token1),
	7591: uint16(292),
	7592: uint16(1),
	7593: uint16(sym__tag_value),
	7594: uint16(2),
	7595: uint16(736),
	7596: uint16(1),
	7597: uint16(anon_sym_await),
	7598: uint16(3),
	7599: uint16(2),
	7600: uint16(sym_comment),
	7601: uint16(aux_sym__tag_value_token1),
	7602: uint16(2),
	7603: uint16(590),
	7604: uint16(1),
	7605: uint16(sym__end_tag_name),
	7606: uint16(3),
	7607: uint16(2),
	7608: uint16(sym_comment),
	7609: uint16(aux_sym__tag_value_token1),
	7610: uint16(2),
	7611: uint16(584),
	7612: uint16(1),
	7613: uint16(anon_sym_elseif),
	7614: uint16(3),
	7615: uint16(2),
	7616: uint16(sym_comment),
	7617: uint16(aux_sym__tag_value_token1),
	7618: uint16(2),
	7619: uint16(738),
	7620: uint16(1),
	7621: uint16(anon_sym_RBRACE),
	7622: uint16(3),
	7623: uint16(2),
	7624: uint16(sym_comment),
	7625: uint16(aux_sym__tag_value_token1),
	7626: uint16(2),
	7627: uint16(740),
	7628: uint16(1),
	7629: uint16(anon_sym_RBRACE),
	7630: uint16(3),
	7631: uint16(2),
	7632: uint16(sym_comment),
	7633: uint16(aux_sym__tag_value_token1),
	7634: uint16(2),
	7635: uint16(742),
	7636: uint16(1),
	7637: uint16(anon_sym_GT),
	7638: uint16(3),
	7639: uint16(2),
	7640: uint16(sym_comment),
	7641: uint16(aux_sym__tag_value_token1),
	7642: uint16(3),
	7643: uint16(3),
	7644: uint16(1),
	7645: uint16(sym_comment),
	7646: uint16(602),
	7647: uint16(1),
	7648: uint16(aux_sym__tag_value_token1),
	7649: uint16(252),
	7650: uint16(1),
	7651: uint16(sym__tag_value),
	7652: uint16(2),
	7653: uint16(744),
	7654: uint16(1),
	7655: uint16(sym_svelte_raw_text),
	7656: uint16(3),
	7657: uint16(2),
	7658: uint16(sym_comment),
	7659: uint16(aux_sym__tag_value_token1),
	7660: uint16(2),
	7661: uint16(746),
	7662: uint16(1),
	7663: uint16(sym_svelte_raw_text),
	7664: uint16(3),
	7665: uint16(2),
	7666: uint16(sym_comment),
	7667: uint16(aux_sym__tag_value_token1),
	7668: uint16(2),
	7669: uint16(748),
	7670: uint16(1),
	7671: uint16(sym_svelte_raw_text),
	7672: uint16(3),
	7673: uint16(2),
	7674: uint16(sym_comment),
	7675: uint16(aux_sym__tag_value_token1),
	7676: uint16(2),
	7677: uint16(750),
	7678: uint16(1),
	7679: uint16(sym__doctype),
	7680: uint16(3),
	7681: uint16(2),
	7682: uint16(sym_comment),
	7683: uint16(aux_sym__tag_value_token1),
	7684: uint16(2),
	7685: uint16(752),
	7686: uint16(1),
	7687: uint16(anon_sym_RBRACE),
	7688: uint16(3),
	7689: uint16(2),
	7690: uint16(sym_comment),
	7691: uint16(aux_sym__tag_value_token1),
	7692: uint16(2),
	7693: uint16(3),
	7694: uint16(1),
	7695: uint16(sym_comment),
	7696: uint16(754),
	7697: uint16(1),
	7698: uint16(aux_sym__tag_value_token1),
	7699: uint16(2),
	7700: uint16(3),
	7701: uint16(1),
	7702: uint16(sym_comment),
	7703: uint16(756),
	7704: uint16(1),
	7705: uint16(aux_sym__tag_value_token1),
	7706: uint16(2),
	7707: uint16(3),
	7708: uint16(1),
	7709: uint16(sym_comment),
	7710: uint16(758),
	7711: uint16(1),
	7712: uint16(aux_sym__tag_value_token1),
	7713: uint16(2),
	7714: uint16(3),
	7715: uint16(1),
	7716: uint16(sym_comment),
	7717: uint16(760),
	7718: uint16(1),
	7719: uint16(aux_sym__tag_value_token1),
}

var ts_small_parse_table_map = [343]uint32_t{
	1:   uint32(85),
	2:   uint32(170),
	3:   uint32(255),
	4:   uint32(340),
	5:   uint32(424),
	6:   uint32(508),
	7:   uint32(592),
	8:   uint32(676),
	9:   uint32(754),
	10:  uint32(832),
	11:  uint32(910),
	12:  uint32(988),
	13:  uint32(1063),
	14:  uint32(1138),
	15:  uint32(1213),
	16:  uint32(1288),
	17:  uint32(1360),
	18:  uint32(1432),
	19:  uint32(1504),
	20:  uint32(1576),
	21:  uint32(1648),
	22:  uint32(1720),
	23:  uint32(1792),
	24:  uint32(1864),
	25:  uint32(1936),
	26:  uint32(2008),
	27:  uint32(2080),
	28:  uint32(2152),
	29:  uint32(2224),
	30:  uint32(2296),
	31:  uint32(2368),
	32:  uint32(2440),
	33:  uint32(2512),
	34:  uint32(2584),
	35:  uint32(2656),
	36:  uint32(2725),
	37:  uint32(2794),
	38:  uint32(2863),
	39:  uint32(2932),
	40:  uint32(3001),
	41:  uint32(3070),
	42:  uint32(3139),
	43:  uint32(3208),
	44:  uint32(3277),
	45:  uint32(3346),
	46:  uint32(3402),
	47:  uint32(3458),
	48:  uint32(3514),
	49:  uint32(3570),
	50:  uint32(3623),
	51:  uint32(3676),
	52:  uint32(3723),
	53:  uint32(3770),
	54:  uint32(3817),
	55:  uint32(3864),
	56:  uint32(3911),
	57:  uint32(3958),
	58:  uint32(3999),
	59:  uint32(4040),
	60:  uint32(4056),
	61:  uint32(4072),
	62:  uint32(4088),
	63:  uint32(4104),
	64:  uint32(4128),
	65:  uint32(4144),
	66:  uint32(4160),
	67:  uint32(4176),
	68:  uint32(4192),
	69:  uint32(4208),
	70:  uint32(4224),
	71:  uint32(4240),
	72:  uint32(4264),
	73:  uint32(4280),
	74:  uint32(4296),
	75:  uint32(4312),
	76:  uint32(4328),
	77:  uint32(4344),
	78:  uint32(4360),
	79:  uint32(4376),
	80:  uint32(4398),
	81:  uint32(4414),
	82:  uint32(4430),
	83:  uint32(4446),
	84:  uint32(4462),
	85:  uint32(4478),
	86:  uint32(4494),
	87:  uint32(4510),
	88:  uint32(4526),
	89:  uint32(4542),
	90:  uint32(4558),
	91:  uint32(4574),
	92:  uint32(4590),
	93:  uint32(4606),
	94:  uint32(4622),
	95:  uint32(4638),
	96:  uint32(4654),
	97:  uint32(4670),
	98:  uint32(4686),
	99:  uint32(4702),
	100: uint32(4718),
	101: uint32(4734),
	102: uint32(4750),
	103: uint32(4766),
	104: uint32(4782),
	105: uint32(4798),
	106: uint32(4814),
	107: uint32(4830),
	108: uint32(4846),
	109: uint32(4862),
	110: uint32(4878),
	111: uint32(4894),
	112: uint32(4910),
	113: uint32(4926),
	114: uint32(4942),
	115: uint32(4958),
	116: uint32(4974),
	117: uint32(4990),
	118: uint32(5006),
	119: uint32(5022),
	120: uint32(5038),
	121: uint32(5054),
	122: uint32(5070),
	123: uint32(5086),
	124: uint32(5102),
	125: uint32(5118),
	126: uint32(5134),
	127: uint32(5150),
	128: uint32(5166),
	129: uint32(5182),
	130: uint32(5198),
	131: uint32(5214),
	132: uint32(5230),
	133: uint32(5246),
	134: uint32(5262),
	135: uint32(5278),
	136: uint32(5294),
	137: uint32(5310),
	138: uint32(5326),
	139: uint32(5342),
	140: uint32(5366),
	141: uint32(5390),
	142: uint32(5406),
	143: uint32(5430),
	144: uint32(5454),
	145: uint32(5478),
	146: uint32(5494),
	147: uint32(5518),
	148: uint32(5534),
	149: uint32(5555),
	150: uint32(5570),
	151: uint32(5585),
	152: uint32(5600),
	153: uint32(5615),
	154: uint32(5630),
	155: uint32(5645),
	156: uint32(5660),
	157: uint32(5675),
	158: uint32(5690),
	159: uint32(5705),
	160: uint32(5720),
	161: uint32(5741),
	162: uint32(5762),
	163: uint32(5783),
	164: uint32(5798),
	165: uint32(5819),
	166: uint32(5840),
	167: uint32(5861),
	168: uint32(5876),
	169: uint32(5896),
	170: uint32(5918),
	171: uint32(5940),
	172: uint32(5962),
	173: uint32(5984),
	174: uint32(5998),
	175: uint32(6018),
	176: uint32(6040),
	177: uint32(6060),
	178: uint32(6082),
	179: uint32(6104),
	180: uint32(6126),
	181: uint32(6148),
	182: uint32(6170),
	183: uint32(6181),
	184: uint32(6198),
	185: uint32(6209),
	186: uint32(6224),
	187: uint32(6235),
	188: uint32(6252),
	189: uint32(6269),
	190: uint32(6286),
	191: uint32(6303),
	192: uint32(6320),
	193: uint32(6333),
	194: uint32(6350),
	195: uint32(6361),
	196: uint32(6372),
	197: uint32(6382),
	198: uint32(6392),
	199: uint32(6402),
	200: uint32(6412),
	201: uint32(6424),
	202: uint32(6434),
	203: uint32(6446),
	204: uint32(6458),
	205: uint32(6472),
	206: uint32(6486),
	207: uint32(6498),
	208: uint32(6512),
	209: uint32(6526),
	210: uint32(6540),
	211: uint32(6554),
	212: uint32(6565),
	213: uint32(6576),
	214: uint32(6587),
	215: uint32(6598),
	216: uint32(6609),
	217: uint32(6620),
	218: uint32(6631),
	219: uint32(6642),
	220: uint32(6651),
	221: uint32(6662),
	222: uint32(6673),
	223: uint32(6684),
	224: uint32(6695),
	225: uint32(6704),
	226: uint32(6713),
	227: uint32(6722),
	228: uint32(6731),
	229: uint32(6742),
	230: uint32(6753),
	231: uint32(6764),
	232: uint32(6775),
	233: uint32(6786),
	234: uint32(6797),
	235: uint32(6808),
	236: uint32(6819),
	237: uint32(6830),
	238: uint32(6841),
	239: uint32(6852),
	240: uint32(6863),
	241: uint32(6872),
	242: uint32(6883),
	243: uint32(6894),
	244: uint32(6903),
	245: uint32(6914),
	246: uint32(6925),
	247: uint32(6936),
	248: uint32(6944),
	249: uint32(6952),
	250: uint32(6960),
	251: uint32(6968),
	252: uint32(6978),
	253: uint32(6986),
	254: uint32(6994),
	255: uint32(7002),
	256: uint32(7010),
	257: uint32(7018),
	258: uint32(7026),
	259: uint32(7034),
	260: uint32(7042),
	261: uint32(7050),
	262: uint32(7058),
	263: uint32(7066),
	264: uint32(7074),
	265: uint32(7082),
	266: uint32(7090),
	267: uint32(7098),
	268: uint32(7106),
	269: uint32(7114),
	270: uint32(7122),
	271: uint32(7130),
	272: uint32(7138),
	273: uint32(7146),
	274: uint32(7154),
	275: uint32(7162),
	276: uint32(7170),
	277: uint32(7178),
	278: uint32(7186),
	279: uint32(7194),
	280: uint32(7202),
	281: uint32(7210),
	282: uint32(7218),
	283: uint32(7226),
	284: uint32(7234),
	285: uint32(7242),
	286: uint32(7250),
	287: uint32(7258),
	288: uint32(7266),
	289: uint32(7274),
	290: uint32(7282),
	291: uint32(7290),
	292: uint32(7298),
	293: uint32(7306),
	294: uint32(7314),
	295: uint32(7322),
	296: uint32(7330),
	297: uint32(7338),
	298: uint32(7346),
	299: uint32(7354),
	300: uint32(7364),
	301: uint32(7372),
	302: uint32(7380),
	303: uint32(7388),
	304: uint32(7396),
	305: uint32(7404),
	306: uint32(7412),
	307: uint32(7420),
	308: uint32(7428),
	309: uint32(7436),
	310: uint32(7444),
	311: uint32(7452),
	312: uint32(7460),
	313: uint32(7468),
	314: uint32(7478),
	315: uint32(7486),
	316: uint32(7494),
	317: uint32(7502),
	318: uint32(7510),
	319: uint32(7518),
	320: uint32(7528),
	321: uint32(7536),
	322: uint32(7544),
	323: uint32(7554),
	324: uint32(7564),
	325: uint32(7574),
	326: uint32(7584),
	327: uint32(7594),
	328: uint32(7602),
	329: uint32(7610),
	330: uint32(7618),
	331: uint32(7626),
	332: uint32(7634),
	333: uint32(7642),
	334: uint32(7652),
	335: uint32(7660),
	336: uint32(7668),
	337: uint32(7676),
	338: uint32(7684),
	339: uint32(7692),
	340: uint32(7699),
	341: uint32(7706),
	342: uint32(7713),
}

var ts_parse_actions = [762]TSParseActionEntry{
	0: {},
	1: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
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
		_ [0]byte
		f struct {
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
		_ [32]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fextra: libc.BoolUint8(1 != 0),
	}})))),
	5: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:  uint8(TSParseActionTypeReduce),
		Fsymbol: uint16(sym_document),
	})))),
	7: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
		_ [64]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(274)),
	}})))),
	9: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
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
		_ [80]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(207)),
	}})))),
	11: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
		_ [96]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(269)),
	}})))),
	13: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
		_ [112]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(140)),
	}})))),
	15: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
		_ [128]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(59)),
	}})))),
	17: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	18: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [144]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(48)),
	}})))),
	19: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
		_ [160]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(49)),
	}})))),
	21: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
		_ [176]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(47)),
	}})))),
	23: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
		_ [192]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(50)),
	}})))),
	25: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
		_ [208]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(51)),
	}})))),
	27: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
		_ [224]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(52)),
	}})))),
	29: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
		_ [240]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(339)),
	}})))),
	31: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
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
		_ [256]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(211)),
	}})))),
	33: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
		_ [272]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(240)),
	}})))),
	35: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
		_ [288]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(98)),
	}})))),
	37: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	38: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [304]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(60)),
	}})))),
	39: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
		_ [320]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(69)),
	}})))),
	41: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
		_ [336]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(61)),
	}})))),
	43: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
		_ [352]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(244)),
	}})))),
	45: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
		_ [368]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(112)),
	}})))),
	47: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
		_ [384]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(100)),
	}})))),
	49: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	50: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [400]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(57)),
	}})))),
	51: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
		_ [416]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(56)),
	}})))),
	53: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
		_ [432]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(54)),
	}})))),
	55: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	56: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [448]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(55)),
	}})))),
	57: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	58: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_document_repeat1),
	})))),
	59: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
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
	60: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_document_repeat1),
	})))),
	61: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [488]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(274)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	62: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
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
	63: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_document_repeat1),
	})))),
	64: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [512]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(207)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	65: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
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
	66: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_document_repeat1),
	})))),
	67: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [536]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(269)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	68: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
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
	69: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_document_repeat1),
	})))),
	70: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [560]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(140)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	71: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
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
	72: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_document_repeat1),
	})))),
	73: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [584]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(59)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	74: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	75: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [600]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(53)),
	}})))),
	76: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
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
	77: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_document_repeat1),
	})))),
	78: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [624]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(339)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	79: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
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
	80: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_document_repeat1),
	})))),
	81: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [648]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(211)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	82: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
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
	83: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_document_repeat1),
	})))),
	84: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [672]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(322)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	85: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
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
	86: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_document_repeat1),
	})))),
	87: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [696]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(98)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	88: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
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
	89: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_document_repeat1),
	})))),
	90: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [720]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(60)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	91: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_document),
	})))),
	93: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
		_ [752]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(58)),
	}})))),
	95: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
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
	96: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_else_if_block),
	})))),
	97: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [776]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(59)),
	}})))),
	98: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
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
	99: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_catch_block),
	})))),
	100: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [800]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(59)),
	}})))),
	101: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
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
	102: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_then_block),
	})))),
	103: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [824]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(59)),
	}})))),
	104: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
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
	105: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_else_block),
	})))),
	106: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [848]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(59)),
	}})))),
	107: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
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
	108: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_catch_block),
	})))),
	109: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [872]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(59)),
	}})))),
	110: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
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
	111: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_else_if_block),
	})))),
	112: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [896]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(59)),
	}})))),
	113: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
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
	114: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_then_block),
	})))),
	115: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [920]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(59)),
	}})))),
	116: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
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
	117: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_else_block),
	})))),
	118: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [944]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(274)),
	}})))),
	119: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
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
	120: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_else_block),
	})))),
	121: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [968]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(207)),
	}})))),
	122: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
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
	123: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_else_block),
	})))),
	124: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [992]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(269)),
	}})))),
	125: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
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
	126: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_else_block),
	})))),
	127: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [1016]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(140)),
	}})))),
	128: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
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
	129: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_else_block),
	})))),
	130: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [1040]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(274)),
	}})))),
	131: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
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
	132: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_else_block),
	})))),
	133: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [1064]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(207)),
	}})))),
	134: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
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
	135: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_else_block),
	})))),
	136: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [1088]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(269)),
	}})))),
	137: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
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
	138: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_else_block),
	})))),
	139: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [1112]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(140)),
	}})))),
	140: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
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
	141: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_else_block),
	})))),
	142: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [1136]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(59)),
	}})))),
	143: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	144: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [1152]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(170)),
	}})))),
	145: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	146: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [1168]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(233)),
	}})))),
	147: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	148: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [1184]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(329)),
	}})))),
	149: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	150: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [1200]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(191)),
	}})))),
	151: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	152: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [1216]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(249)),
	}})))),
	153: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	154: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [1232]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(241)),
	}})))),
	155: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	156: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [1248]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(283)),
	}})))),
	157: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	158: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [1264]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(310)),
	}})))),
	159: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	160: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [1280]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(311)),
	}})))),
	161: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	162: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [1296]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(266)),
	}})))),
	163: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	164: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [1312]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(257)),
	}})))),
	165: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	166: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [1328]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(273)),
	}})))),
	167: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	168: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_element),
	})))),
	169: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_element),
	})))),
	171: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	172: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_snippet_statement),
	})))),
	173: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_snippet_statement),
	})))),
	175: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	176: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_start_tag),
	})))),
	177: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_start_tag),
	})))),
	179: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	180: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_self_closing_tag),
	})))),
	181: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_self_closing_tag),
	})))),
	183: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	184: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [1472]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(79)),
	}})))),
	185: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	186: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [1488]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(80)),
	}})))),
	187: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	188: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [1504]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(175)),
	}})))),
	189: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	190: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [1520]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(277)),
	}})))),
	191: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	192: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_key_statement),
	})))),
	193: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_key_statement),
	})))),
	195: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	196: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_erroneous_end_tag),
	})))),
	197: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_erroneous_end_tag),
	})))),
	199: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	200: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_expression),
	})))),
	201: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_expression),
	})))),
	203: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_element),
	})))),
	205: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_element),
	})))),
	207: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_script_element),
	})))),
	209: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_script_element),
	})))),
	211: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	212: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_style_element),
	})))),
	213: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_style_element),
	})))),
	215: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_if_statement),
	})))),
	217: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_if_statement),
	})))),
	219: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	220: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [1760]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(176)),
	}})))),
	221: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_each_statement),
	})))),
	223: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_each_statement),
	})))),
	225: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_await_statement),
	})))),
	227: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_await_statement),
	})))),
	229: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	230: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_key_statement),
	})))),
	231: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_key_statement),
	})))),
	233: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	234: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_snippet_statement),
	})))),
	235: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_snippet_statement),
	})))),
	237: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	238: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_doctype),
	})))),
	239: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_doctype),
	})))),
	241: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	242: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_start_tag),
	})))),
	243: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_start_tag),
	})))),
	245: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	246: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_self_closing_tag),
	})))),
	247: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_self_closing_tag),
	})))),
	249: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	250: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_start_tag_repeat1),
	})))),
	251: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
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
	252: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_start_tag_repeat1),
	})))),
	253: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [2024]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(175)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	254: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
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
	255: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_start_tag_repeat1),
	})))),
	256: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [2048]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(277)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	257: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	258: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_html_tag),
		Fproduction_id: uint16(5),
	})))),
	259: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_html_tag),
		Fproduction_id: uint16(5),
	})))),
	261: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	262: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_const_tag),
		Fproduction_id: uint16(5),
	})))),
	263: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_const_tag),
		Fproduction_id: uint16(5),
	})))),
	265: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	266: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_debug_tag),
		Fproduction_id: uint16(5),
	})))),
	267: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_debug_tag),
		Fproduction_id: uint16(5),
	})))),
	269: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_render_tag),
		Fproduction_id: uint16(5),
	})))),
	271: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_render_tag),
		Fproduction_id: uint16(5),
	})))),
	273: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_end_tag),
	})))),
	275: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_end_tag),
	})))),
	277: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_if_end),
		Fproduction_id: uint16(5),
	})))),
	279: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_if_end),
		Fproduction_id: uint16(5),
	})))),
	281: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_if_statement),
	})))),
	283: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_if_statement),
	})))),
	285: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_each_end),
		Fproduction_id: uint16(5),
	})))),
	287: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_each_end),
		Fproduction_id: uint16(5),
	})))),
	289: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_each_statement),
	})))),
	291: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_each_statement),
	})))),
	293: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	294: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_await_end),
		Fproduction_id: uint16(5),
	})))),
	295: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_await_end),
		Fproduction_id: uint16(5),
	})))),
	297: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_await_statement),
	})))),
	299: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_await_statement),
	})))),
	301: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	302: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_key_end),
		Fproduction_id: uint16(5),
	})))),
	303: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_key_end),
		Fproduction_id: uint16(5),
	})))),
	305: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	306: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_snippet_end),
		Fproduction_id: uint16(5),
	})))),
	307: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_snippet_end),
		Fproduction_id: uint16(5),
	})))),
	309: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	310: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_if_statement),
	})))),
	311: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_if_statement),
	})))),
	313: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	314: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_each_statement),
	})))),
	315: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_each_statement),
	})))),
	317: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	318: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_await_statement),
	})))),
	319: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_await_statement),
	})))),
	321: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	322: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__node),
	})))),
	323: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__node),
	})))),
	325: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	326: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_element),
	})))),
	327: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_element),
	})))),
	329: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	330: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_script_element),
	})))),
	331: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_script_element),
	})))),
	333: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	334: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_style_element),
	})))),
	335: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_style_element),
	})))),
	337: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	338: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_if_statement),
	})))),
	339: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_if_statement),
	})))),
	341: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	342: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_each_statement),
	})))),
	343: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_each_statement),
	})))),
	345: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	346: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_await_statement),
	})))),
	347: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_await_statement),
	})))),
	349: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	350: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [2800]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(63)),
	}})))),
	351: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	352: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [2816]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(149)),
	}})))),
	353: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	354: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [2832]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(178)),
	}})))),
	355: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	356: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [2848]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(121)),
	}})))),
	357: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	358: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [2864]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(64)),
	}})))),
	359: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	360: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [2880]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(228)),
	}})))),
	361: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	362: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [2896]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(194)),
	}})))),
	363: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	364: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [2912]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(336)),
	}})))),
	365: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	366: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_await_start),
		Fproduction_id: uint16(5),
	})))),
	367: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	368: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_await_start),
		Fproduction_id: uint16(5),
	})))),
	369: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	370: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_key_start),
		Fproduction_id: uint16(5),
	})))),
	371: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	372: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_key_start),
		Fproduction_id: uint16(5),
	})))),
	373: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	374: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_else_start),
		Fproduction_id: uint16(5),
	})))),
	375: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_else_start),
		Fproduction_id: uint16(5),
	})))),
	377: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	378: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_then_start),
		Fproduction_id: uint16(5),
	})))),
	379: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_then_start),
		Fproduction_id: uint16(5),
	})))),
	381: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	382: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_catch_start),
		Fproduction_id: uint16(5),
	})))),
	383: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_catch_start),
		Fproduction_id: uint16(5),
	})))),
	385: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	386: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_else_if_start),
		Fproduction_id: uint16(3),
	})))),
	387: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_else_if_start),
		Fproduction_id: uint16(3),
	})))),
	389: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	390: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_then_start),
		Fproduction_id: uint16(5),
	})))),
	391: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_then_start),
		Fproduction_id: uint16(5),
	})))),
	393: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	394: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_catch_start),
		Fproduction_id: uint16(5),
	})))),
	395: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	396: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_catch_start),
		Fproduction_id: uint16(5),
	})))),
	397: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	398: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_each_start),
		Fproduction_id: uint16(6),
	})))),
	399: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_each_start),
		Fproduction_id: uint16(6),
	})))),
	401: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_snippet_start),
		Fproduction_id: uint16(5),
	})))),
	403: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	404: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_snippet_start),
		Fproduction_id: uint16(5),
	})))),
	405: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_snippet_start),
		Fproduction_id: uint16(5),
	})))),
	407: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_snippet_start),
		Fproduction_id: uint16(5),
	})))),
	409: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
		_ [3280]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(226)),
	}})))),
	411: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_start_tag_repeat1),
	})))),
	413: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [3304]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(194)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	414: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
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
	415: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_start_tag_repeat1),
	})))),
	416: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [3328]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(336)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	417: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	418: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [3344]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(188)),
	}})))),
	419: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	420: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [3360]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(177)),
	}})))),
	421: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	422: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [3376]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(179)),
	}})))),
	423: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	424: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_if_start),
		Fproduction_id: uint16(3),
	})))),
	425: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	426: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_if_start),
		Fproduction_id: uint16(3),
	})))),
	427: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	428: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [3424]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(199)),
	}})))),
	429: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	430: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [3440]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(180)),
	}})))),
	431: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	432: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [3456]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(181)),
	}})))),
	433: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	434: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [3472]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(242)),
	}})))),
	435: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	436: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [3488]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(229)),
	}})))),
	437: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	438: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_each_start),
		Fproduction_id: uint16(4),
	})))),
	439: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	440: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_each_start),
		Fproduction_id: uint16(4),
	})))),
	441: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	442: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [3536]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(280)),
	}})))),
	443: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	444: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [3552]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(221)),
	}})))),
	445: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	446: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [3568]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(288)),
	}})))),
	447: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	448: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [3584]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(250)),
	}})))),
	449: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	450: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [3600]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(251)),
	}})))),
	451: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	452: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [3616]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(186)),
	}})))),
	453: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	454: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [3632]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(202)),
	}})))),
	455: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	456: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [3648]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(337)),
	}})))),
	457: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	458: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [3664]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fextra: libc.BoolUint8(1 != 0),
	}})))),
	459: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	460: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [3680]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(208)),
	}})))),
	461: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	462: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [3696]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(338)),
	}})))),
	463: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	464: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__single_quoted_attribute_value),
	})))),
	465: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
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
	466: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__single_quoted_attribute_value),
	})))),
	467: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [3736]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(202)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	468: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
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
	469: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__single_quoted_attribute_value),
	})))),
	470: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [3760]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(337)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	471: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	472: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__double_quoted_attribute_value),
	})))),
	473: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
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
	474: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__double_quoted_attribute_value),
	})))),
	475: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [3800]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(208)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	476: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
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
	477: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__double_quoted_attribute_value),
	})))),
	478: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [3824]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(338)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	479: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	480: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_attribute),
	})))),
	481: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	482: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [3856]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(164)),
	}})))),
	483: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	484: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [3872]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(197)),
	}})))),
	485: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	486: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [3888]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(200)),
	}})))),
	487: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	488: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [3904]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(201)),
	}})))),
	489: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	490: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [3920]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(278)),
	}})))),
	491: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	492: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_quoted_attribute_value),
		Fproduction_id: uint16(7),
	})))),
	493: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
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
	494: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_if_statement_repeat1),
	})))),
	495: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [3960]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(234)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	496: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_attribute),
	})))),
	498: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
		_ [3992]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(185)),
	}})))),
	500: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
		_ [4008]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(341)),
	}})))),
	502: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
		_ [4024]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(343)),
	}})))),
	504: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
		_ [4040]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(344)),
	}})))),
	506: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	507: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [4056]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(342)),
	}})))),
	508: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
		_ [4072]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(195)),
	}})))),
	510: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
		_ [4088]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(166)),
	}})))),
	512: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_attribute),
		Fproduction_id: uint16(1),
	})))),
	514: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_quoted_attribute_value),
	})))),
	516: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym__single_quoted_attribute_value),
	})))),
	518: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
		_ [4152]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(330)),
	}})))),
	520: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	521: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [4168]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(214)),
	}})))),
	522: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	523: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [4184]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(148)),
	}})))),
	524: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	525: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [4200]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(167)),
	}})))),
	526: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	527: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [4216]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(168)),
	}})))),
	528: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	529: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym__double_quoted_attribute_value),
	})))),
	530: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	531: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [4248]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(263)),
	}})))),
	532: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	533: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [4264]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(247)),
	}})))),
	534: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	535: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [4280]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(220)),
	}})))),
	536: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	537: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [4296]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(141)),
	}})))),
	538: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	539: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [4312]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(215)),
	}})))),
	540: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	541: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [4328]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(232)),
	}})))),
	542: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	543: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [4344]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(224)),
	}})))),
	544: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	545: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [4360]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(219)),
	}})))),
	546: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	547: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym__each_start_tag),
		Fproduction_id: uint16(2),
	})))),
	548: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	549: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [4392]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(255)),
	}})))),
	550: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	551: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [4408]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(256)),
	}})))),
	552: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	553: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [4424]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(218)),
	}})))),
	554: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	555: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_script_start_tag),
	})))),
	556: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	557: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym__catch_tag),
		Fproduction_id: uint16(2),
	})))),
	558: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	559: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_style_start_tag),
	})))),
	560: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	561: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_style_start_tag),
	})))),
	562: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	563: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [4504]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(245)),
	}})))),
	564: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	565: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [4520]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(227)),
	}})))),
	566: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	567: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [4536]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(331)),
	}})))),
	568: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	569: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [4552]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(154)),
	}})))),
	570: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	571: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [4568]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(262)),
	}})))),
	572: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	573: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [4584]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(314)),
	}})))),
	574: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	575: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [4600]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(309)),
	}})))),
	576: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	577: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [4616]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(155)),
	}})))),
	578: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	579: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [4632]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(265)),
	}})))),
	580: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	581: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [4648]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(275)),
	}})))),
	582: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	583: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [4664]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(272)),
	}})))),
	584: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	585: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [4680]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(264)),
	}})))),
	586: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	587: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [4696]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(267)),
	}})))),
	588: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	589: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_script_start_tag),
	})))),
	590: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	591: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [4728]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(293)),
	}})))),
	592: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	593: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym__then_tag),
		Fproduction_id: uint16(2),
	})))),
	594: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	595: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [4760]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(68)),
	}})))),
	596: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	597: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym__key_start_tag),
		Fproduction_id: uint16(2),
	})))),
	598: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	599: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym__snippet_start_tag),
		Fproduction_id: uint16(2),
	})))),
	600: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	601: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [4808]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(84)),
	}})))),
	602: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	603: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [4824]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(302)),
	}})))),
	604: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	605: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym__each_end_tag),
		Fproduction_id: uint16(2),
	})))),
	606: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	607: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [4856]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(169)),
	}})))),
	608: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	609: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [4872]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(294)),
	}})))),
	610: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	611: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [4888]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(297)),
	}})))),
	612: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	613: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [4904]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(93)),
	}})))),
	614: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	615: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [4920]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(78)),
	}})))),
	616: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	617: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [4936]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(161)),
	}})))),
	618: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	619: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym__await_end_tag),
		Fproduction_id: uint16(2),
	})))),
	620: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	621: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [4968]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(157)),
	}})))),
	622: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	623: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym__else_if_tag),
		Fproduction_id: uint16(2),
	})))),
	624: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	625: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [5000]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(158)),
	}})))),
	626: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	627: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [5016]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(300)),
	}})))),
	628: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	629: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym__else_tag),
		Fproduction_id: uint16(2),
	})))),
	630: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	631: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [5048]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(94)),
	}})))),
	632: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	633: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [5064]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(334)),
	}})))),
	634: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	635: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [5080]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(85)),
	}})))),
	636: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	637: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [5096]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(316)),
	}})))),
	638: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	639: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [5112]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(110)),
	}})))),
	640: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	641: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [5128]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(111)),
	}})))),
	642: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	643: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [5144]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(301)),
	}})))),
	644: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	645: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [5160]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(86)),
	}})))),
	646: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	647: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [5176]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(151)),
	}})))),
	648: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	649: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [5192]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(317)),
	}})))),
	650: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	651: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [5208]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(152)),
	}})))),
	652: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	653: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym__if_start_tag),
		Fproduction_id: uint16(2),
	})))),
	654: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	655: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f uint8_t
		_ [7]byte
	}{f: uint8(TSParseActionTypeAccept)})),
	656: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	657: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym__if_end_tag),
		Fproduction_id: uint16(2),
	})))),
	658: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	659: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [5272]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(282)),
	}})))),
	660: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	661: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [5288]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(304)),
	}})))),
	662: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	663: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [5304]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(153)),
	}})))),
	664: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	665: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [5320]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(87)),
	}})))),
	666: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	667: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [5336]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(120)),
	}})))),
	668: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	669: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym__await_start_tag),
		Fproduction_id: uint16(2),
	})))),
	670: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	671: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [5368]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(122)),
	}})))),
	672: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	673: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [5384]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(123)),
	}})))),
	674: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	675: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [5400]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(124)),
	}})))),
	676: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	677: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [5416]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(125)),
	}})))),
	678: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	679: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [5432]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(126)),
	}})))),
	680: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	681: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [5448]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(306)),
	}})))),
	682: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	683: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [5464]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(276)),
	}})))),
	684: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	685: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [5480]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(279)),
	}})))),
	686: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	687: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym__key_end_tag),
		Fproduction_id: uint16(2),
	})))),
	688: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	689: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [5512]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(238)),
	}})))),
	690: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	691: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [5528]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(129)),
	}})))),
	692: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	693: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym__snippet_end_tag),
		Fproduction_id: uint16(2),
	})))),
	694: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	695: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [5560]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(259)),
	}})))),
	696: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	697: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [5576]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(305)),
	}})))),
	698: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	699: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [5592]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(131)),
	}})))),
	700: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	701: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [5608]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(156)),
	}})))),
	702: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	703: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:       uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__tag_value),
	})))),
	704: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	705: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [5640]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(159)),
	}})))),
	706: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	707: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [5656]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(133)),
	}})))),
	708: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	709: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [5672]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(134)),
	}})))),
	710: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	711: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [5688]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(260)),
	}})))),
	712: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	713: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [5704]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(267)),
	}})))),
	714: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	715: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [5720]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(254)),
	}})))),
	716: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	717: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [5736]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(89)),
	}})))),
	718: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	719: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [5752]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(298)),
	}})))),
	720: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	721: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [5768]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(160)),
	}})))),
	722: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	723: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [5784]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(165)),
	}})))),
	724: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	725: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [5800]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(184)),
	}})))),
	726: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	727: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [5816]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(203)),
	}})))),
	728: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	729: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [5832]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(204)),
	}})))),
	730: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	731: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [5848]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(205)),
	}})))),
	732: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	733: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [5864]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(82)),
	}})))),
	734: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	735: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [5880]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(287)),
	}})))),
	736: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	737: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [5896]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(261)),
	}})))),
	738: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	739: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [5912]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(83)),
	}})))),
	740: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	741: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [5928]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(91)),
	}})))),
	742: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	743: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [5944]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(67)),
	}})))),
	744: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	745: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [5960]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(318)),
	}})))),
	746: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	747: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [5976]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(319)),
	}})))),
	748: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	749: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [5992]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(320)),
	}})))),
	750: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	751: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [6008]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(324)),
	}})))),
	752: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	753: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		_ [6024]byte
		f struct {
			Ftype1      uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
	}{f: struct {
		Ftype1      uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(127)),
	}})))),
	754: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	755: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym__html_tag),
		Fproduction_id: uint16(2),
	})))),
	756: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	757: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym__render_tag),
		Fproduction_id: uint16(2),
	})))),
	758: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	759: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym__const_tag),
		Fproduction_id: uint16(2),
	})))),
	760: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		_ [0]byte
		f struct {
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
	761: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype1              uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype1:         uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym__debug_tag),
		Fproduction_id: uint16(2),
	})))),
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
const ts_external_token_svelte_raw_text = 9
const ts_external_token_svelte_raw_text_each = 10
const ts_external_token_svelte_raw_text_snippet_arguments = 11
const ts_external_token_AT = 12
const ts_external_token_POUND = 13
const ts_external_token_SLASH = 14
const ts_external_token_COLON = 15

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
	9:  uint16(sym_svelte_raw_text),
	10: uint16(sym_svelte_raw_text_each),
	11: uint16(sym_svelte_raw_text_snippet_arguments),
	12: uint16(anon_sym_AT),
	13: uint16(anon_sym_POUND),
	14: uint16(anon_sym_SLASH),
	15: uint16(anon_sym_COLON),
}

var ts_external_scanner_states = [19][16]uint8{
	1: {
		0:  libc.BoolUint8(1 != 0),
		1:  libc.BoolUint8(1 != 0),
		2:  libc.BoolUint8(1 != 0),
		3:  libc.BoolUint8(1 != 0),
		4:  libc.BoolUint8(1 != 0),
		5:  libc.BoolUint8(1 != 0),
		6:  libc.BoolUint8(1 != 0),
		7:  libc.BoolUint8(1 != 0),
		8:  libc.BoolUint8(1 != 0),
		9:  libc.BoolUint8(1 != 0),
		10: libc.BoolUint8(1 != 0),
		11: libc.BoolUint8(1 != 0),
		12: libc.BoolUint8(1 != 0),
		13: libc.BoolUint8(1 != 0),
		14: libc.BoolUint8(1 != 0),
		15: libc.BoolUint8(1 != 0),
	},
	2: {
		8: libc.BoolUint8(1 != 0),
	},
	3: {
		6: libc.BoolUint8(1 != 0),
		8: libc.BoolUint8(1 != 0),
	},
	4: {
		8:  libc.BoolUint8(1 != 0),
		9:  libc.BoolUint8(1 != 0),
		12: libc.BoolUint8(1 != 0),
		13: libc.BoolUint8(1 != 0),
		14: libc.BoolUint8(1 != 0),
		15: libc.BoolUint8(1 != 0),
	},
	5: {
		8:  libc.BoolUint8(1 != 0),
		9:  libc.BoolUint8(1 != 0),
		12: libc.BoolUint8(1 != 0),
		13: libc.BoolUint8(1 != 0),
		14: libc.BoolUint8(1 != 0),
	},
	6: {
		8:  libc.BoolUint8(1 != 0),
		9:  libc.BoolUint8(1 != 0),
		12: libc.BoolUint8(1 != 0),
		13: libc.BoolUint8(1 != 0),
	},
	7: {
		5: libc.BoolUint8(1 != 0),
		8: libc.BoolUint8(1 != 0),
	},
	8: {
		8:  libc.BoolUint8(1 != 0),
		14: libc.BoolUint8(1 != 0),
		15: libc.BoolUint8(1 != 0),
	},
	9: {
		7: libc.BoolUint8(1 != 0),
		8: libc.BoolUint8(1 != 0),
	},
	10: {
		0: libc.BoolUint8(1 != 0),
		1: libc.BoolUint8(1 != 0),
		2: libc.BoolUint8(1 != 0),
		8: libc.BoolUint8(1 != 0),
	},
	11: {
		8:  libc.BoolUint8(1 != 0),
		14: libc.BoolUint8(1 != 0),
	},
	12: {
		8:  libc.BoolUint8(1 != 0),
		9:  libc.BoolUint8(1 != 0),
		10: libc.BoolUint8(1 != 0),
	},
	13: {
		8: libc.BoolUint8(1 != 0),
		9: libc.BoolUint8(1 != 0),
	},
	14: {
		8:  libc.BoolUint8(1 != 0),
		15: libc.BoolUint8(1 != 0),
	},
	15: {
		8:  libc.BoolUint8(1 != 0),
		11: libc.BoolUint8(1 != 0),
	},
	16: {
		3: libc.BoolUint8(1 != 0),
		4: libc.BoolUint8(1 != 0),
		8: libc.BoolUint8(1 != 0),
	},
	17: {
		3: libc.BoolUint8(1 != 0),
		8: libc.BoolUint8(1 != 0),
	},
	18: {
		4: libc.BoolUint8(1 != 0),
		8: libc.BoolUint8(1 != 0),
	},
}

func tree_sitter_svelte(tls *libc.TLS) (r uintptr) {
	return uintptr(unsafe.Pointer(&language))
}

var language = TSLanguage{
	Fversion:                   uint32(14),
	Fsymbol_count:              uint32(118),
	Ftoken_count:               uint32(52),
	Fexternal_token_count:      uint32(16),
	Fstate_count:               uint32(345),
	Flarge_state_count:         uint32(2),
	Fproduction_id_count:       uint32(8),
	Ffield_count:               uint32(4),
	Fmax_alias_sequence_length: uint16(7),
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
	Fprimary_state_ids: uintptr(unsafe.Pointer(&ts_primary_state_ids)),
}

func init() {
	p := unsafe.Pointer(&language)
	*(*uintptr)(unsafe.Add(p, 144)) = __ccgo_fp(ts_lex)
	*(*uintptr)(unsafe.Add(p, 184)) = __ccgo_fp(tree_sitter_svelte_external_scanner_create)
	*(*uintptr)(unsafe.Add(p, 192)) = __ccgo_fp(tree_sitter_svelte_external_scanner_destroy)
	*(*uintptr)(unsafe.Add(p, 200)) = __ccgo_fp(tree_sitter_svelte_external_scanner_scan)
	*(*uintptr)(unsafe.Add(p, 208)) = __ccgo_fp(tree_sitter_svelte_external_scanner_serialize)
	*(*uintptr)(unsafe.Add(p, 216)) = __ccgo_fp(tree_sitter_svelte_external_scanner_deserialize)
}

func __ccgo_fp(f interface{}) uintptr {
	type iface [2]uintptr
	return (*iface)(unsafe.Pointer(&f))[1]
}

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = "index < self->size\x00third-party/tree-sitter-svelte/src/tree_sitter/array.h\x00old_end <= self->size\x00(uint32_t)((&scanner->tags)->size - 1) < (&scanner->tags)->size\x00/tmp/grammar-gen-2388767627/combined.c\x00</SCRIPT\x00</STYLE\x00(uint32_t)(i) < (&scanner->tags)->size\x00end\x00<!\x00doctype_token1\x00>\x00doctype\x00<\x00/>\x00</\x00=\x00attribute_name\x00attribute_value\x00entity\x00'\x00\"\x00text\x00_single_quoted_attribute_value_token1\x00_double_quoted_attribute_value_token1\x00#\x00if\x00{\x00}\x00:\x00else if\x00else\x00/\x00each\x00as\x00await\x00then\x00catch\x00key\x00snippet\x00snippet_name\x00(\x00)\x00_tag_value_token1\x00@\x00html\x00const\x00debug\x00render\x00tag_name\x00erroneous_end_tag_name\x00_implicit_end_tag\x00raw_text\x00comment\x00svelte_raw_text\x00document\x00_node\x00element\x00script_element\x00style_element\x00start_tag\x00self_closing_tag\x00end_tag\x00erroneous_end_tag\x00attribute\x00quoted_attribute_value\x00_single_quoted_attribute_value\x00_double_quoted_attribute_value\x00if_statement\x00block_start_tag\x00if_start\x00block_tag\x00else_if_start\x00else_if_block\x00else_start\x00else_block\x00block_end_tag\x00if_end\x00each_statement\x00each_start\x00each_end\x00await_statement\x00await_start\x00then_start\x00then_block\x00catch_start\x00catch_block\x00await_end\x00key_statement\x00key_start\x00key_end\x00snippet_statement\x00snippet_start\x00snippet_end\x00expression\x00_tag_value\x00expression_tag\x00html_tag\x00const_tag\x00debug_tag\x00render_tag\x00document_repeat1\x00start_tag_repeat1\x00if_statement_repeat1\x00condition\x00identifier\x00parameter\x00tag\x00"
